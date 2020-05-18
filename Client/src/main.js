import Vue from 'vue'

import router from "./router";
import store from "./store/index.js";
import '@babel/polyfill'
import 'mutationobserver-shim'
import './plugins/bootstrap-vue'
import App from './App.vue'

Vue.config.productionTip = false

Vue.mixin({
  methods: {
    notification: function (title, message, delay, append) {
      this.$bvToast.toast(message, {
        title: title,
        toaster: "b-toaster-top-center",
        autoHideDelay: delay,
        appendToast: append
      });
    },
  }
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
