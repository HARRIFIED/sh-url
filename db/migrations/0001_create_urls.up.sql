-- create the urls table
CREATE TABLE IF NOT EXISTS urls (
  id            SERIAL PRIMARY KEY,
  code          VARCHAR(255) UNIQUE NOT NULL,
  original_url  TEXT NOT NULL,
  created_at    BIGINT NOT NULL
);
