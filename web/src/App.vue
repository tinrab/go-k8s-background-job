<template>
  <div id="app">
    <h1>Hi</h1>
    <button v-on:click="add">Add</button>
  </div>
</template>

<script>
export default {
  methods: {
    add() {
      this.$socket.sendObj({ x: 42 });
    },
  },
};
</script>

<style lang="scss">
</style>
