export default {
  setAuth({ commit }, payload) {
    commit("setAuth", payload);
  },
  setReleaseState({ commit }, payload) {
    commit("setReleaseState", payload);
  }
};
