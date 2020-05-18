import store from "../store/index.js";
import axios from "axios";

var baseURL = process.env.VUE_APP_ROOT_API_DEV;

const API = initAPI({ baseURL, withCredentials: true });

function initAPI(config) {
  const instance = axios.create(config);
  instance.defaults.headers.common["Content-Type"] = "application/json";
  instance.defaults.headers.common["Accept"] = "application/json";
  store.watch(
    state => state.token,
    token => {
      const header = "Bearer " + token;
      instance.defaults.headers.common["Authorization"] = header;
    }
  );

  // override which api to use depending on the the env of the Twitch extension
  // isit running on dev, testing, released
  store.watch(
    state => state.releaseState,
    releaseState => {
      switch (releaseState) {
        case "released":
          instance.defaults.baseURL = process.env.VUE_APP_ROOT_API_PROD;
          break;
        case "testing":
          instance.defaults.baseURL = process.env.VUE_APP_ROOT_API_DEV;
          break;
        default:
          instance.defaults.baseURL = process.env.VUE_APP_ROOT_API_STAGING;
      }
    }
  );

  return instance;
}

export default API;
