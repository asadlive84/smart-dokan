BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
   id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
   first_name VARCHAR (250) NOT NULL,
   last_name VARCHAR (250) NOT NULL,
   password VARCHAR (50) NOT NULL,
   phone_number VARCHAR (50) NOT NULL DEFAULT '',
   email VARCHAR (300) UNIQUE NOT NULL
);

COMMIT;