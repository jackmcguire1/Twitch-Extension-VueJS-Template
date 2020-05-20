import API from "./api";
import store from "../store/";

function getFollowers() {
  const { channelId } = store.state;

  if (!channelId) {
    throw new Error("Missing required data for fetching initial state.");
  }

  const url = `/followers/${channelId}`
  return API.get(url);
}


function followChannel() {
  const { channelId } = store.state;
  if (!channelId) {
    throw new Error("Missing required data for fetching initial state.");
  }
  
  // eslint-disable-next-line no-undef
  Twitch.ext.actions.followChannel("crazyjack12");
}

export default {
  getFollowers,
  followChannel
};
