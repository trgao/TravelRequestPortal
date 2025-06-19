require("dotenv").config();
const { Pool } = require("pg");

const pool = new Pool({
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  host: "localhost",
  port: process.env.DB_PORT,
  database: "requests_db",
});

module.exports = {
  query: (text, params) => pool.query(text, params),
  client: async () => await pool.connect(),
};
