<template>
  <div class="ExtHeader">
    <!-- example nav bar -->
    <b-navbar toggleable="lg" type="dark" variant="info">
      <!-- header -->
      <b-navbar-brand href="#">
        <img
          src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/95/Vue.js_Logo_2.svg/1200px-Vue.js_Logo_2.svg.png"
          width="50"
          height="50"
          alt="VueJS"
        />
        {{ title }}
      </b-navbar-brand>
    </b-navbar>
  </div>
</template>

<script>
export default {
  name: "ExtHeader",
  props: {
    title: String,
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
