CREATE TABLE urls (
    id  SERIAL  NOT NULL,
    url TEXT    NOT NULL UNIQUE,

    PRIMARY KEY(id)
);

CREATE TABLE links (
    id      SERIAL  NOT NULL,
    link    TEXT    NOT NULL UNIQUE,
    urlID   SERIAL  NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (urlID) REFERENCES urls(id)
);


