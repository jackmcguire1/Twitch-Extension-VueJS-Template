<template>
  <div class="panel">
    <ExtHeader title="Panel"></ExtHeader>
    <MessageCentre></MessageCentre>

    <b-overlay :show="showOverlay" rounded="sm">
      <b-input-group prepend="Followers">
        <b-input-group-append>
          <div class="input-group-text">
            <strong> {{ total.toLocaleString() }} </strong>
          </div>
          <b-button @click="followChannel()">
            Follow
          </b-button>
        </b-input-group-append>
      </b-input-group>

      <b-table
        stripped
        hover
        :items="followers"
        :busy="isBusy"
        :fields="fields"
        outlined
        responsive="sm"
      >
        <template v-slot:table-busy>
          <div class="text-center text-danger my-2">
            <b-spinner class="align-middle"></b-spinner>
            <strong>Loading...</strong>
          </div>
        </template>
      </b-table>
    </b-overlay>
  </div>
</template>

<script>
import ExtHeader from "@/components/utils/ExtHeader.vue";
import MessageCentre from "@/components/utils/MessageCentre.vue";
import FollowersAPI from "@/api/followers.js";

export default {
  name: "Panel",
  data() {
    return {
      total: 0,
      followers: [],
      isBusy: true,
      showOverlay: true,
      fields: [
        {
          key: "username",
          sortable: true,
        }
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

    this.getFollowers();
  },
  activated: function() {
    // eslint-disable-next-line no-undef
    ga("set", "page", "/" + this.page);
    // eslint-disable-next-line no-undef
    ga("send", "pageview");
  },
  methods: {
    getFollowers: function() {
      FollowersAPI.getFollowers()
        .then((resp) => {
          console.log("got followers from API", resp);
          // got data, stop showing overaly
          this.total = resp.data.total;
          this.showOverlay = false;

          // show table
          this.followers = resp.data.followers;
          this.isBusy = false;
        })
        .catch((error) => {
          console.log(error);
        });
    },
    followChannel: function() {
      FollowersAPI.followChannel();
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
