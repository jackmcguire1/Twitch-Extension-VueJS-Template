import { SET_AUTH } from "@/store/mutations";
import store from "@/store/index.js";

var Twitch = window.Twitch.ext;

Twitch.onAuthorized(async auth => {
  const { channelId, token, userId } = auth;

  store.commit(SET_AUTH, {
    channelId,
    token,
    userId
  });

});
