DROP INDEX idx_short_url_version ON urls;
CREATE UNIQUE INDEX short_url ON urls(short_url);