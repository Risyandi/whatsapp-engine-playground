const baseUrl = process.env.VUE_APP_BASE_API_URL;

import axios from "axios";

export function sendMessage(id, receiver, text) {
  return axios.post(`${baseUrl}/message/${id}`, { receiver, text });
}
