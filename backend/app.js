const express = require("express");
const cors = require("cors");
const authRouter = require("./src/routes/auth.route");
const cookieParser = require("cookie-parser");
const resourceRouter = require("./src/routes/resource.route");

const app = express();

const corsOptions = {
  origin: true,
  credentials: true,
};
app.use(cors(corsOptions));

app.use(cookieParser());

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use("/api/auth", authRouter);
app.use("/api", resourceRouter);

app.listen(3000, () => {
  console.log("Server is running on port 3000");
});
