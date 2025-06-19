DROP DATABASE IF EXISTS requests_db;
CREATE DATABASE requests_db;
\c requests_db;

CREATE TYPE REQUEST_STATUS AS ENUM('approved', 'rejected', 'pending');

CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    is_admin BOOLEAN NOT NULL,
    last_login_at TIMESTAMPTZ NOT NULL,
    company_id INT NOT NULL,
    UNIQUE(email),
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

CREATE TABLE requests (
    id SERIAL PRIMARY KEY,
    start_location TEXT NOT NULL,
    destination TEXT NOT NULL,
    date_time TIMESTAMPTZ NOT NULL,
    purpose TEXT NOT NULL,
    status REQUEST_STATUS NOT NULL,
    admin_remarks TEXT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
