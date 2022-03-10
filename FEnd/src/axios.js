import axios from "axios";

const instance = axios.create({
  // baseURL: "http://localhost:11000",
  // server bri:
  // baseURL: "http://172.18.132.107:11000",
  // server sv 1:
  baseURL: "http://165.22.55.132:8009"
  //baseURL: "http://165.22.55.132:8009",
  // server sv 2:
  // baseURL: "http://192.168.2.166:11000",
  // baseURL: "http://165.22.55.132:8011",
});

export default instance;