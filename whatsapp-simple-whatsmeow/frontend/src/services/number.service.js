const baseUrl = process.env.VUE_APP_BASE_API_URL;

import axios from "axios";

export function getNumber() {
  return axios.get(`${baseUrl}/number`);
}

export function getOneNumber(id) {
  return axios.get(`${baseUrl}/number/${id}`);
}

export function addNumber() {
  return axios.post(`${baseUrl}/number`);
}

export function removeNumber(id) {
  return axios.delete(`${baseUrl}/number/${id}`);
}
