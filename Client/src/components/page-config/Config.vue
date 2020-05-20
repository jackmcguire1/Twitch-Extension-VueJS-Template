<template>
  <div class="config">
    <ExtHeader title="Configuration"></ExtHeader>

    <div>
      <b-jumbotron
        header="BootstrapVue"
        lead="Bootstrap v4 Components for Vue.js 2"
      >
        <p>For more information visit website</p>
      </b-jumbotron>
    </div>

    <div>
      <b-form-group
        id="fieldset-1"
        description="Let us know your name."
        label="Enter your name"
        label-for="input-1"
        :invalid-feedback="invalidFeedback"
        :valid-feedback="validFeedback"
        :state="state"
      >
        <b-form-input
          id="input-1"
          v-model="name"
          :state="state"
          trim
        ></b-form-input>
      </b-form-group>
    </div>
  </div>
</template>

<script>
import ExtHeader from "@/components/utils/ExtHeader.vue";

export default {
  name: "config",
  computed: {
    state() {
      return this.name.length >= 4 ? true : false;
    },
    invalidFeedback() {
      if (this.name.length > 4) {
        return "";
      } else if (this.name.length > 0) {
        return "Enter at least 4 characters";
      } else {
        return "Please enter something";
      }
    },
    validFeedback() {
      return this.state === true ? "Thank you" : "";
    },
  },
  data() {
    return {
      name: "",
    };
  },
  props: {
    locale: String,
    page: String,
  },
  mounted() {
    if (this.locale) {
      // eslint-disable-next-line no-undef
      ga("set", "language", this.locale);
    }
  },
  activated: function() {
    // eslint-disable-next-line no-undef
    ga("set", "page", "/" + this.page);
    // eslint-disable-next-line no-undef
    ga("send", "pageview");
  },
  components: {
    ExtHeader,
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
