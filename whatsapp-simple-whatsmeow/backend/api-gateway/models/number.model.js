const mongoose = require("mongoose");

const NumberSchema = new mongoose.Schema({
  isConnected: {
    type: Boolean,
    required: true,
  },
  name: {
    type: String,
    required: false,
  },
  phoneNumber: {
    type: String,
    required: false,
  },
  jid: {
    type: String,
    required: false,
  },
  qrcode: {
    type: String,
    required: false,
  },
  virtualMachineId: {
    type: mongoose.Types.ObjectId,
    required: true,
  },
  createdAt: {
    type: Date,
    required: true,
  },
});

NumberSchema.index({ virtualMachineId: 1 });
NumberSchema.index({ createdAt: -1 });

module.exports = mongoose.model("numbers", NumberSchema);
