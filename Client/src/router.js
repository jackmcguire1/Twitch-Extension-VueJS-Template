import Vue from "vue";
import Router from "vue-router";
import Panel from "@/components/page-panel/Panel.vue";
import Config from "@/components/page-config/Config.vue";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "*/panel.html",
      name: "panel",
      component: Panel,
      props: route => ({
        locale: route.query.locale,
        page: "panel"
      })
    },
    {
      path: "*/config.html",
      component: Config,
      props: route => ({
        locale: route.query.locale,
        page: "config",
      })
    },
    {
      path: "*/liveconfig.html",
      component: Config,
      props: route => ({
        locale: route.query.locale,
        page: "liveconfig"
      })
    },
  ],
  mode: "history"
});
