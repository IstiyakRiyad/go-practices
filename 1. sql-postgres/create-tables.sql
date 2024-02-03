-- Create Database
CREATE DATABASE recordings;

-- Insert into the database
-- \c recordings;

-- Delete table if already exists
DROP TABLE IF EXISTS album;

-- Create Table. Here id should be SERIAL for postgres
CREATE TABLE album (
  id         SERIAL PRIMARY KEY,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL
);

-- Insert some data to the table
INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);

