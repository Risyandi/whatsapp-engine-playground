const engineApi = require("../apis/engine.api");
const numberDatasource = require("../datasources/number.datasource");
const virtualMachineDatasource = require("../datasources/virtual-machine.datasource");

async function get() {
  return await numberDatasource.get();
}

async function add() {
  const virtualMachine = await virtualMachineDatasource.getEmpty();
  if (!virtualMachine) {
    throw new Error("there are no empty virtual machines");
  }

  const newNumber = await numberDatasource.add(virtualMachine._id);

  await engineApi.authorize(virtualMachine.ipPrivate, newNumber._id);
  await virtualMachineDatasource.incCurrentUsage(virtualMachine._id);

  return newNumber;
}

async function getOne(id) {
  const number = await numberDatasource.getOne(id);
  if (!number) {
    throw new Error("number not found");
  }

  if (!number.isConnected) {
    try {
      const virtualMachine = await virtualMachineDatasource.getById(
        number.virtualMachineId
      );
      await engineApi.authorize(virtualMachine.ipPrivate, id);
    } catch (error) {
      console.error("error authorize number in getOne", id, error?.message);
    }
  }

  return number;
}

async function remove(id) {
  const number = await numberDatasource.getOne(id);
  if (!number) {
    throw new Error("number not found");
  }

  const virtualMachine = await virtualMachineDatasource.getById(
    number.virtualMachineId
  );
  if (!virtualMachine) {
    throw new Error("virtual machine not found");
  }

  await engineApi.removeAuthorize(virtualMachine.ipPrivate, id);
  await virtualMachineDatasource.decCurrentUsage(virtualMachine._id);
  await numberDatasource.remove(id);
}

module.exports = {
  get,
  add,
  getOne,
  remove,
};
