const mongoose = require("mongoose");

mongoose
  .connect(process.env.MONGODB_URL)
  .catch((error) => console.log("mongodb connection error", error));
