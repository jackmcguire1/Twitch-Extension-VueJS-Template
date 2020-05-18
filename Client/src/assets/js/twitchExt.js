import { SET_AUTH } from "@/store/mutations";
import store from "@/store/index.js";

var Twitch = window.Twitch.ext;
var ga = window.ga;
if (!Twitch.viewer.id) {
  ga("set", "userId", Twitch.viewer.opaqueId);
} else {
  ga("set", "userId", Twitch.viewer.id);
}

// Prevent duplicate listeners due to HMR
Twitch.unlisten("broadcast", listenCb);
Twitch.listen("broadcast", listenCb);

Twitch.unlisten("global", listenCb);
Twitch.listen("global", listenCb);

Twitch.bits.onTransactionCancelled(bitsTxCancelled);
Twitch.bits.onTransactionComplete(bitsTxCompleted);

function bitsTxCancelled() {
  console.log("found bits transaction cancelled", txObj)
  ga("send", {
    hitType: "event",
    eventCategory: "bits",
    eventAction: "processed",
    eventLabel: "bits-tx-cancelled"
  });
}

function bitsTxCompleted(txObj) {
  console.log("found bits transaction", txObj)
  ga("send", {
    hitType: "event",
    eventCategory: "bits",
    eventAction: "processed",
    eventLabel: "bits-tx-cancelled"
  });
}

Twitch.onAuthorized(async auth => {
  const parts = auth.token.split(".");
  const { channelId, token, userId } = auth;

  store.commit(SET_AUTH, {
    channelId,
    token,
    userId
  });

});

function listenCb(target, contentType, message) {
  const { type, data } = JSON.parse(message);

  console.log(message, type, data)

  ga("send", {
    hitType: "event",
    eventCategory: "pubsub-event",
    eventAction: "processed",
    eventLabel: type
  });
}
