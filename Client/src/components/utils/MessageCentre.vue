<template>
  <div id="MessageCentre"></div>
</template>

<script>
import ConfigAPI from "@/api/config.js";

export default {
  name: "MessageCentre",
  created() {
    this.motd();
  },
  methods: {
    motd: function() {
      console.log("displaying message of the day");

      var globalConfig = ConfigAPI.getGlobalConfig();
      var broadcasterConfig = ConfigAPI.getBroadcasterConfig();

      var motd = {};
      if (globalConfig.message) {
        console.log("displaying GLOBAL message of the day");

        motd = globalConfig.message;
        this.notification(motd.title, motd.msg, motd.ts, false);
      }
      if (broadcasterConfig.motd) {
        motd = broadcasterConfig.motd;
        if (motd.msg == "") {
          return;
        }

        this.notification("Message Of The Day", motd.msg, motd.ts, true);
      }
    },
  },
};
</script>
