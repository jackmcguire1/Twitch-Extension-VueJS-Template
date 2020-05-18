<template>
  <div class="panel">
    <ExtHeader title="Panel"></ExtHeader>
    <MessageCentre></MessageCentre>

    <!-- example table -->
    <div>
      <b-table striped hover :items="items"></b-table>
    </div>
  </div>
</template>

<script>
import ExtHeader from "@/components/utils/ExtHeader.vue";
import MessageCentre from "@/components/utils/MessageCentre.vue";

export default {
  name: "Panel",
  data() {
    return {
      items: [
        { age: 40, first_name: "Dickerson", last_name: "Macdonald" },
        { age: 21, first_name: "Larsen", last_name: "Shaw" },
        { age: 89, first_name: "Geneva", last_name: "Wilson" },
        { age: 38, first_name: "Jami", last_name: "Carney" },
      ],
    };
  },
  props: {
    page: String,
  },
  components: {
    ExtHeader,
    MessageCentre,
  },
  mounted() {
    if (this.locale) {
      // eslint-disable-next-line no-undef
      ga("set", "language", this.locale);
    }
  },
  activated: function() {
    // eslint-disable-next-line no-undef
    ga("set", "page", "/"+this.page);
    // eslint-disable-next-line no-undef
    ga("send", "pageview");
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
