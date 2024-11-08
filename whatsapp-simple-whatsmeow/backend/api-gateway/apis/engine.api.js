const axios = require("axios");

async function authorize(ip, id) {
  try {
    const { data } = await axios.post(`http://${ip}:5002/authorize/${id}`);
    return data;
  } catch (error) {
    throw new Error(error?.response?.data?.errorMessage || error?.message);
  }
}

async function removeAuthorize(ip, id) {
  try {
    const { data } = await axios.delete(`http://${ip}:5002/authorize/${id}`);
    return data;
  } catch (error) {
    throw new Error(error?.response?.data?.errorMessage || error?.message);
  }
}

async function sendMessage(ip, id, receiver, text) {
  try {
    const { data } = await axios.post(`http://${ip}:5002/message/${id}`, {
      receiver,
      text,
    });

    return data;
  } catch (error) {
    console.log(error?.response?.data);
    throw new Error(error?.response?.data?.errorMessage || error?.message);
  }
}

module.exports = {
  authorize,
  removeAuthorize,
  sendMessage,
};
