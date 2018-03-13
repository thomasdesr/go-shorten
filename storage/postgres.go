package storage

import (
	"context"
	"database/sql"
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
	SELECT u.url
	  FROM urls u
	  JOIN links l
		ON l.urlID = u.id
	 WHERE l.link = $1;
`

var fuzzyMatchQuery = `
	WITH top_link AS (
		SELECT l.link
		FROM links l
		WHERE difference(l.link, $1) > 2
		AND levenshtein(l.link, $1) < 5
		ORDER BY levenshtein(l.link, $1)
		LIMIT 1
	)

	SELECT l.link
	FROM urls u
	JOIN top_link l
	ON l.urlID = u.id;
`

func (p *Postgres) Load(ctx context.Context, rawShort string) (string, error) {
	short, err := sanitizeShort(rawShort)
	if err != nil {
		return "", err
	}

	var url string
	switch err := p.dbx.GetContext(ctx, &url, loadQuery, short); err {
	case nil:
		// Short found, serve this
		return url, nil
	case sql.ErrNoRows:
		fuzzyMatchedShort, err := p.loadFuzzyMatch(ctx, short)
		if err != nil {
			return "", err
		}

		if len(fuzzyMatchedShort) == 0 {
			// No fuzzy match found
			return "", ErrShortNotSet
		}
		return fuzzyMatchedShort, err
	default:
		return "", errors.Wrap(err, "load from DB failed")
	}
}

func (p *Postgres) loadFuzzyMatch(ctx context.Context, short string) (string, error) {
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
		SELECT id
		FROM urls
		WHERE url = :url
	)

	INSERT INTO links (link, urlID)
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
	short, err := sanitizeShort(rawShort)
	if err != nil {
		return err
	}
	if _, err := validateURL(url); err != nil {
		return err
	}

	return saveLink(ctx, p.dbx, short, url)
}

var setLimitQuery = `
	SELECT set_limit(0.2);
`

var searchQuery = `
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
	ORDER BY sum(sml) DESC;
`

type SearchResult struct {
	Link string
	URL string
}

func (p *Postgres) Search(ctx context.Context, searchTerm string) ([]SearchResult, error) {
	term, err := sanitizeShort(searchTerm)
	if err != nil {
		return nil, err
	}

	if _, err := p.dbx.ExecContext(ctx, setLimitQuery); err != nil {
		return nil, err
	}

	var results []SearchResult
	switch err := p.dbx.SelectContext(ctx, &results, searchQuery, term); err {
	case nil:
		return results, nil
	case sql.ErrNoRows:
		return nil, ErrNoResults
	default:
		return nil, errors.Wrap(err, "load from DB failed")
	}
}
