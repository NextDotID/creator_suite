<template>
  <div
    data-collapse="medium"
    data-animation="default"
    data-duration="400"
    data-easing="ease"
    data-easing2="ease"
    role="banner"
    class="navigation w-nav"
    id="header"
  >
    <div class="navigation-items">
      <a href="/">
        <img
          :src="browser"
          width="65"
          alt=""
          sizes="(max-width: 479px) 100vw, 65px"
          class="logo-image"
        />
      </a>

      <div class="navigation-wrap">
        <nav role="navigation" class="navigation-items w-nav-menu">
          <router-link to="/" class="navigation-item w-nav-link">
            ABOUT
          </router-link>
          <router-link to="/browser" class="navigation-item w-nav-link">
            CONTENT
          </router-link>
          <router-link to="/" class="navigation-item w-nav-link">
            CONTACT
          </router-link>
          <!-- <v-btn class="navbutton"> Connect Wallet </v-btn> -->
          <a href="/" class="button w-inline-block">
            <div>Connect Wallet</div>
          </a>
        </nav>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import browser from '@img/BROWSER.png'
</script>

<style lang="less" scoped>
#header {
  background-color: #000000;
  // background-color: #202020;
  height: 60px;
}

.navbutton {
  background-image: url('@img/explore.png');
  background-position: 0px 0px;
  background-size: cover;
  margin-left: 10px;
}

.dialog-encrypt {
  margin-left: 10%;
  width: 70%;
}

.navigation {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  padding: 12px 25px;
  -webkit-box-align: center;
  -webkit-align-items: center;
  -ms-flex-align: center;
  align-items: center;
  border-style: none none solid;
  border-width: 1px;
  border-color: #46d0e1;
  background-color: #fff;
  padding-right: 30px;
  padding-left: 30px;
}

.navigation-items {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  -webkit-box-pack: justify;
  -webkit-justify-content: space-between;
  -ms-flex-pack: justify;
  justify-content: space-between;
  -webkit-box-align: center;
  -webkit-align-items: center;
  -ms-flex-align: center;
  align-items: center;
  -webkit-box-flex: 1;
  -webkit-flex: 1;
  -ms-flex: 1;
  flex: 1;
}

.navigation-item {
  // padding: 12px 25px;
  -webkit-transition: background-color 400ms ease, opacity 400ms ease,
    color 400ms ease;
  transition: background-color 400ms ease, opacity 400ms ease, color 400ms ease;
  text-align: center;
  font-size: 12px;
}

.navigation-item:hover {
  background-image: url('@img/explore.png');
  background-position: 0px 0px;
  background-size: cover;
  opacity: 0.9;
}

.navigation-item:active {
  opacity: 0.8;
}

.navigation-wrap {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  -webkit-box-align: center;
  -webkit-align-items: center;
  -ms-flex-align: center;
  align-items: center;
}

.w-nav-menu {
  position: relative;
  float: right;
}

.w-nav-link {
  position: relative;
  display: inline-block;
  vertical-align: top;
  text-decoration: none;
  color: #ffffff;
  padding: 14px;
  text-align: left;
  margin-left: auto;
  margin-right: auto;
}

.w-nav {
  position: relative;
  background: #000000;
  z-index: 1000;
}
.w-nav:before,
.w-nav:after {
  content: ' ';
  display: table;
  grid-column-start: 1;
  grid-row-start: 1;
  grid-column-end: 2;
  grid-row-end: 2;
  clear: both;
}

.button {
  padding: 12px 25px;
  border-radius: 0px;
  background-color: #5ccbe3;
  background-image: url('@img/explore.png');
  background-position: 0px 0px;
  background-size: cover;
  -webkit-transition: background-color 400ms ease, opacity 400ms ease,
    color 400ms ease;
  transition: background-color 400ms ease, opacity 400ms ease, color 400ms ease;
  color: #fff;
  font-size: 12px;
  line-height: 20px;
  letter-spacing: 2px;
  text-decoration: none;
  text-transform: uppercase;
}

.button:hover {
  background-color: #54cde2;
  background-image: none;
  color: #fff;
}

.button:active {
  background-color: #cb87f4;
}

.button {
  -webkit-box-pack: center;
  -webkit-justify-content: center;
  -ms-flex-pack: center;
  justify-content: center;
}
.w-inline-block {
  max-width: 100%;
  display: inline-block;
}

.logo-image {
  display: block;
}
</style>
