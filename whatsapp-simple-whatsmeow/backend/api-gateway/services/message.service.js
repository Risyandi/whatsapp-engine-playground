const engineApi = require("../apis/engine.api");
const numberDatasource = require("../datasources/number.datasource");
const virtualMachineDatasource = require("../datasources/virtual-machine.datasource");

async function send(id, receiver, text) {
  const number = await numberDatasource.getOne(id);
  if (!number) {
    throw new Error("number not found");
  }

  if (!number.isConnected) {
    throw new Error("number disconnected, please connect first");
  }

  const virtualMachine = await virtualMachineDatasource.getById(
    number.virtualMachineId
  );
  if (!virtualMachine) {
    throw new Error("virtual machine not found");
  }

  return await engineApi.sendMessage(
    virtualMachine.ipPrivate,
    id,
    receiver,
    text
  );
}

module.exports = {
  send,
};
