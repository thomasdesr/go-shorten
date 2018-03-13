CREATE INDEX links_trgm_idx ON links USING gin (link gin_trgm_ops);
CREATE INDEX urls_trgm_idx ON urls USING gin (url gin_trgm_ops);
