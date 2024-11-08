require("dotenv").config();
require("./databases/mongo.database");

const express = require("express");
const bodyParser = require("body-parser");
const cors = require("cors");

const app = express();

const numberController = require("./controllers/number.controller");
const messageController = require("./controllers/message.controller");

app.use(cors());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());
app.use((req, res, next) => {
  console.log(req.method, req.path, new Date());
  next();
});

app.get("/number", numberController.handleGetNumber);
app.post("/number", numberController.handleAddNumber);
app.get("/number/:id", numberController.handleGetOneNumber);
app.delete("/number/:id", numberController.handleRemoveNumber);
app.post("/message/:numberId", messageController.handleSendMessage);

app.listen(5001, () => console.log("backend running at port 5001"));
