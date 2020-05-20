<template>
  <div id="app">
    <router-view v-if="isAuthed"/>
  </div>
</template>

<script>
import Vuex from 'vuex'

export default {
  name: 'App',
  computed: {
    ...Vuex.mapState(['isAuthed']),
    canFetchInitialState () {
      return this.isAuthed
    }
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
