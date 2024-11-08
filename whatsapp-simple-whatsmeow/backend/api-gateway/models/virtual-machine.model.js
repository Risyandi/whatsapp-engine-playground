const mongoose = require("mongoose");

const VirtualMachineSchema = new mongoose.Schema({
  ipPublic: {
    type: String,
    required: true,
  },
  ipPrivate: {
    type: String,
    required: true,
  },
  currentUsage: {
    type: Number,
    required: true,
  },
});

VirtualMachineSchema.index({ currentUsage: 1 });

module.exports = mongoose.model("virtual-machines", VirtualMachineSchema);
