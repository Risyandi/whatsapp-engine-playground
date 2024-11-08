const mongoose = require("mongoose");
const virtualMachineModel = require("../models/virtual-machine.model");
const maxInstancePerVM = process.env.MAX_NUMBER_PER_VM;

async function getEmpty() {
  return await virtualMachineModel.findOne({
    currentUsage: { $lt: maxInstancePerVM },
  });
}

async function getById(id) {
  return await virtualMachineModel.findOne({
    _id: new mongoose.Types.ObjectId(id),
  });
}

async function incCurrentUsage(id) {
  await virtualMachineModel.updateOne(
    { _id: new mongoose.Types.ObjectId(id) },
    { $inc: { currentUsage: 1 } }
  );
}

async function decCurrentUsage(id) {
  await virtualMachineModel.updateOne(
    { _id: new mongoose.Types.ObjectId(id) },
    { $inc: { currentUsage: -1 } }
  );
}

module.exports = {
  getEmpty,
  getById,
  incCurrentUsage,
  decCurrentUsage,
};
