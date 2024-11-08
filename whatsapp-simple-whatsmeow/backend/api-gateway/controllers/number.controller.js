const numberService = require("../services/number.service");

async function handleGetNumber(req, res) {
  try {
    const numbers = await numberService.get();
    return res.send({ numbers });
  } catch (error) {
    return res.status(500).send({
      errorMessage: error?.message || error,
    });
  }
}

async function handleAddNumber(req, res) {
  try {
    const newNumber = await numberService.add();
    return res.send(newNumber);
  } catch (error) {
    return res.status(500).send({
      errorMessage: error?.message || error,
    });
  }
}

async function handleGetOneNumber(req, res) {
  try {
    const id = req.params.id;
    const number = await numberService.getOne(id);
    return res.send(number);
  } catch (error) {
    return res.status(500).send({
      errorMessage: error?.message || error,
    });
  }
}

async function handleRemoveNumber(req, res) {
  try {
    const id = req.params.id;
    await numberService.remove(id);
    return res.send(null);
  } catch (error) {
    return res.status(500).send({
      errorMessage: error?.message || error,
    });
  }
}

module.exports = {
  handleGetNumber,
  handleAddNumber,
  handleGetOneNumber,
  handleRemoveNumber,
};
