<template>
  <div class="header">
    <Navigation></Navigation>
  </div>
  <div>
    <FileBrowser path="storage/"></FileBrowser>
  </div>
  <div class="footer">
    <Footer></Footer>
  </div>
  <!-- <div class="content">
  </div> -->
</template>

<script setup lang="ts">
import FileBrowser from '@cp/FileBrowser.vue'
import Navigation from '@cp/Navigation.vue'
import Footer from '@cp/Footer.vue'
</script>

<style lang="less" scoped>
.content {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: @color-black;
  margin-top: (@margin * 3);
  a {
    color: #42b983;
  }
}
.header {
  // height: 4%;
  display: flex;
  flex-direction: column;
}
.footer {
  height: 100%;
  // width: 100%;
  display: flex;
}
</style>
