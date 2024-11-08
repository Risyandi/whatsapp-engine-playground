const messageService = require("../services/message.service");

async function handleSendMessage(req, res) {
  try {
    const numberId = req.params.numberId;
    const { receiver, text } = req.body;

    const send = await messageService.send(numberId, receiver, text);
    return res.send(send);
  } catch (error) {
    return res.status(500).send({
      errorMessage: error?.message || error,
    });
  }
}

module.exports = {
  handleSendMessage,
};
