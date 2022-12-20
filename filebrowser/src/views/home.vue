<template>
  <div class="header">
    <Navigation></Navigation>
  </div>
  <div class="intro">
    <IntroSection></IntroSection>
  </div>
  <div class="footer">
    <Footer></Footer>
  </div>
</template>

<script setup lang="ts">
import Navigation from '@cp/Navigation.vue'
import IntroSection from '@cp/IntroSection.vue'
import Footer from '@/components/Footer.vue'
</script>

<style lang="less" scoped>
.header {
  // height: 4%;
  display: flex;
  flex-direction: column;
}

.intro {
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: black;
}

.footer {
  height: 100%;
  // width: 100%;
  display: flex;
}
</style>
