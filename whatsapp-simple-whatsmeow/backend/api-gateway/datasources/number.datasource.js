const mongoose = require("mongoose");
const numberModel = require("../models/number.model");

async function get() {
  return await numberModel.find({});
}

async function add(virtualMachineId) {
  return await numberModel.create({
    isConnected: false,
    name: "",
    phoneNumber: "",
    jid: "",
    qrcode: "",
    virtualMachineId: new mongoose.Types.ObjectId(virtualMachineId),
    createdAt: new Date(),
  });
}

async function getOne(id) {
  return await numberModel.findOne({ _id: new mongoose.Types.ObjectId(id) });
}

async function remove(id) {
  return await numberModel.deleteOne({ _id: new mongoose.Types.ObjectId(id) });
}

module.exports = {
  get,
  add,
  getOne,
  remove,
};
