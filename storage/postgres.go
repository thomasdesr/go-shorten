package storage

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Postgres struct {
	dbx *sqlx.DB
}

func NewPostgres(connectURL string) (*Postgres, error) {
	db, err := sqlx.Open("postgres", connectURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a DB connector")
	}

	// Retry connecting up to 10 times
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			return &Postgres{dbx: db}, nil
		}

		time.Sleep(time.Second)
	}

	return nil, errors.Wrap(err, "failed to connect to DB")

}

var loadQuery = `
	SELECT
		regexp_replace($1, l.link, u.url)
	FROM
		urls u
	JOIN
		links l
			ON l.urlID = u.id
	WHERE
		$1 ~ ('^' || l.link || '$')
`

func (p *Postgres) Load(ctx context.Context, rawShort string) (string, error) {
	short, err := postgresSanitizeShort(rawShort)
	if err != nil {
		return "", err
	}

	var s struct {
		Url string
		ID  int
	}
	switch err := p.dbx.GetContext(ctx, &s, loadQuery, short); err {
	case nil:
		// Short found, log access
		if err := p.accessEvent(ctx, s.ID); err != nil {
			log.Printf("Error logging access event: %s", err)
		}
	case sql.ErrNoRows:
		fuzzyMatchedShort, err := p.loadFuzzyMatch(ctx, short)
		if err != nil {
			// Found something similar, pass that back
			return fuzzyMatchedShort, err
		}

		// No fuzzy match found
		return "", ErrShortNotSet
	default:
		return "", errors.Wrap(err, "load from DB failed")
	}

	return s.Url, nil
}

func (p *Postgres) accessEvent(ctx context.Context, link_id int) error {
	const accessEventQuery = `
		INSERT INTO
			links_usage(linkID)
		SELECT
			l.id
		FROM
			links l
		WHERE
			l.id = $1
		ON CONFLICT(linkID, day)
			DO UPDATE
				SET hit_count = links_usage.hit_count + 1;
	`

	if _, err := p.dbx.ExecContext(ctx, accessEventQuery, link_id); err != nil {
		return errors.Wrap(err, "load from DB failed")
	}
	return nil
}

func (p *Postgres) loadFuzzyMatch(ctx context.Context, short string) (string, error) {
	const fuzzyMatchQuery = `
		SELECT
			l.link
		FROM
			links l
		WHERE
				difference(l.link, $1) > 2
			AND levenshtein(l.link, $1) < 5
		ORDER BY levenshtein(l.link, $1)
		LIMIT    1
	`

	var fuzzyMatchedShort string
	switch err := p.dbx.GetContext(ctx, &fuzzyMatchedShort, fuzzyMatchQuery, short); err {
	case nil:
		// Found a fuzzy match
		return fuzzyMatchedShort, ErrFuzzyMatchFound
	case sql.ErrNoRows:
		// Didn't find a good enough match, no answer
		return "", nil
	default:
		return "", errors.Wrap(err, "load from DB for fuzzyMatch failed")
	}
}

var saveURLQuery = `
	INSERT INTO urls(url)
		VALUES (:url)
		ON CONFLICT DO NOTHING;
`

var saveLinkQuery = `
	WITH url_id AS (
		SELECT
			id
		FROM
			urls
		WHERE
			url = :url
	)

	INSERT INTO
		links (link, urlID)
	VALUES
		(:link, (SELECT id FROM url_id))
	ON CONFLICT (link)
		DO UPDATE
			SET urlID = (SELECT id FROM url_id)
			WHERE links.link = :link
	;
`

func saveLink(ctx context.Context, dbx *sqlx.DB, short string, url string) error {
	tx, err := dbx.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}

	if _, err := tx.NamedExecContext(
		ctx,
		saveURLQuery,
		&struct{ URL string }{url},
	); err != nil {
		return errors.Wrap(err, "failed to insert url")
	}

	if _, err := tx.NamedExecContext(ctx,
		saveLinkQuery,
		&struct {
			Link string
			URL  string
		}{short, url},
	); err != nil {
		return errors.Wrap(err, "failed to insert short")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "SaveName transaction failed")
	}

	return nil
}

func (p *Postgres) SaveName(ctx context.Context, rawShort string, url string) error {
	short, err := postgresSanitizeShort(rawShort)
	if err != nil {
		return err
	}
	if _, err := validateURL(url); err != nil {
		return err
	}

	return saveLink(ctx, p.dbx, short, url)
}

func (p *Postgres) Search(ctx context.Context, searchTerm string) ([]SearchResult, error) {
	const setLimitQuery = `
		SELECT set_limit(0.2)
	` // Sets the upper limit for the `%` operator

	const searchQuery = `
		WITH url_matches AS (
			SELECT l.link, u.url, similarity(u.url, $1) AS sml
			FROM urls u
			JOIN links l
			ON u.id = l.urlId
			WHERE u.url % $1
		),
		link_matches AS (
			SELECT l.link, u.url, similarity(l.link, $1) AS sml
			FROM links l
			JOIN urls u
			ON l.urlId = u.id
			WHERE l.link % $1
		),
		union_matches AS (
			SELECT *
			FROM url_matches
			UNION ALL
			SELECT *
			FROM link_matches
		)
	
		SELECT link, url
		FROM union_matches
		GROUP BY link, url
		ORDER BY sum(sml) DESC
	`

	if _, err := p.dbx.ExecContext(ctx, setLimitQuery); err != nil {
		return nil, err
	}

	var results []SearchResult
	switch err := p.dbx.SelectContext(ctx, &results, searchQuery, searchTerm); err {
	case nil:
		return results, nil
	default:
		return nil, errors.Wrap(err, "load from DB failed")
	}
}

func (p *Postgres) TopNForPeriod(ctx context.Context, n int, days int) ([]TopNResult, error) {
	const getTopLinksForPeriodQuery = `
		SELECT
			l.link,
			sum(lu.hit_count) as hitCount
		FROM
			links l
		JOIN
			links_usage lu
				ON l.id = lu.linkID
		WHERE
			lu.day >= CURRENT_DATE - $2::integer
		GROUP BY
			l.id
		ORDER BY
			hitCount DESC
		LIMIT
			$1
	`

	var results []TopNResult
	if err := p.dbx.SelectContext(
		ctx,
		&results,
		getTopLinksForPeriodQuery,
		n,
		days,
	); err != nil {
		return nil, err
	}

	return results, nil
}

func postgresSanitizeShort(raw string) (string, error) {
	if _, err := regexp.Compile(raw); err != nil {
		return "", errors.Wrap(err, "unable to validate regex")
	}

	return raw, nil
}
