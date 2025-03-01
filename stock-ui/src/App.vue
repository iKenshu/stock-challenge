<template>
  <div id="app">
    <StockList />
  </div>
</template>

<script>
import StockList from './components/StockList.vue'

export default {
  name: 'App',
  components: {
    StockList
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
