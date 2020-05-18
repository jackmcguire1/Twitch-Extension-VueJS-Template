import * as M from "./mutations";
import actions from "./actions";

import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export const mutations = {
  [M.SET_AUTH](state, payload) {
    state.isAuthed = true;
    state.token = payload.token;
    state.channelId = payload.channelId;
    state.userId = payload.userId;
  },
  [M.SET_RELEASE_STATE](state, releaseState) {
    state.releaseState = releaseState;
    switch (releaseState) {
      case "released":
        // set dynamic config here if being viewed in released mode i.e. public
        break;
      default:
        break;
    }
  }
};

const store = new Vuex.Store({
  state: {
    isAuthed: false,
    token: null,
    channelId: -1,
    userId: -1,
    releaseState: "",
  },
  actions,
  mutations
});

export default store;
