CREATE TYPE customer_status AS ENUM ('active', 'nonactive');

CREATE TABLE IF NOT EXISTS customers(
    id serial PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    phone_number INT NULL,
    status customer_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
 