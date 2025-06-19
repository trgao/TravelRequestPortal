# TravelRequestPortal

POC Project: Corporate Travel Request Portal

## Prerequisites

- npm
- psql
- nodejs

## Database setup

Create the database from `db.sql` script

```sh
psql -f db.sql
```

# Node.js setup

Clone the project
Generate a jwt secret, and add a `.env` file

```sh
DB_USER=<postgres_username>
DB_PASSWORD=<postgres_password>
DB_PORT=<postgres_localhost_port>
JWT_SECRET=<jwt_secret>
```

Start up backend server

```sh
npm install
node app.js
```

Backend is available at http://localhost:3000

# Vue.js setup

Clone the project
Start up frontend server

```sh
npm install
npm run dev
```

Frontend is available at http://localhost:5173
