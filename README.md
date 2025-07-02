# TravelRequestPortal

Proof of concept project: corporate travel request portal

Employees can submit and view their travel requests, while admins can approve and reject the requests of the employees and also add any remarks. A company can have multiple admin accounts managing their employees. 

## Prerequisites

- npm
- psql
- nodejs

## Database setup

Create the postgresql database from `./backend/src/database/db.sql` script

```sh
psql -f ./backend/src/database/db.sql
```

## Node.js setup

Clone the project

Generate a jwt secret, and add a `.env` file in `./backend`

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

## Vue.js setup

Clone the project

Start up frontend server

```sh
npm install
npm run dev
```

Frontend is available at http://localhost:5173
