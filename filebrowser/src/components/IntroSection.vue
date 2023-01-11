<template>
  <div class="section cc-store-home-wrap">
    <div class="intro-header">
      <section class="hero-heading-left wf-section">
        <div class="main-intro">
          <div class="hero-wrapper">
            <div class="hero-split">
              <img
                src="@img/CreatorSuite.png"
                loading="lazy"
                width="600"
                sizes="(max-width: 479px) 100vw, (max-width: 767px) 92vw, (max-width: 991px) 84vw, 40vw"
                alt=""
              />
              <br />
              <p class="paragraph">filebrowser —— A tool for CreatorSuite.</p>
              <div class="paragraph-bigger cc-bigger-light">
                // Manage/Encrypt/Decrypt/Publish you data.<br />// A
                applications to encrypt or decrypt your data. <br />// It can be
                used to encrypt/decrypt files<br />// backups or even large
                object.<br />
              </div>
              <a href="/" class="button w-inline-block">
                <div>Explore</div>
              </a>
            </div>
            <!-- <div class="hero-split">
              <img
                src="@img/DrawKit-Vector-Illustration-Team-Work-3.png"
                loading="lazy"
                width="359"
                sizes="(max-width: 479px) 100vw, 359px"
                alt=""
                class="shadow-two"
              />
            </div> -->
          </div>
        </div>
      </section>
    </div>
  </div>
  <div class="intro-detail">
    <div class="step w-container">
      <h3 class="heading">
        <strong class="bold-text">What do I do?</strong>
      </h3>
      <div class="wf-section">
        <div class="text-block">
          I offer three steps to help you publish and manage your encrypted
          content<br />‍
        </div>
      </div>
      <div class="w-row">
        <div class="gradient w-col w-col-3">
          <h5 class="heading-2"><strong>01 Browser file</strong></h5>
          <p class="paragraph-2">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse
            varius enim in eros elementum tristique.<br />
          </p>
          <a href="/" class="link">Text Link</a>
        </div>
        <div class="gradient w-col w-col-3">
          <h5 class="heading-2"><strong>02 Generate link </strong></h5>
          <p class="paragraph-2">
            Duis cursus, mi quis viverra ornare, eros dolor interdum nulla, ut
            commodo diam libero vitae erat.‍<br />
          </p>
          <a href="/" class="link">Text Link</a>
        </div>
        <div class="gradient w-col w-col-3">
          <h5 class="heading-2"><strong>03 Publish</strong></h5>
          <p class="paragraph-2">
            Aenean faucibus nibh et justo cursus id rutrum lorem imperdiet. Nunc
            ut sem vitae risus tristique posuere.<br />
          </p>
          <a href="/" class="link">Text Link</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import creatorsuite from '@img/CreatorSuite.png'
// srcset="@img/BROWSER-p-500.png 500w, @img/BROWSER.png 555w"
</script>

<style lang="less" scoped>
.section {
  margin-right: 0px;
  margin-left: 0px;
  // background-color: transparent;
  // background-color: #000000;
  // height: 80px;
}

.section.cc-store-home-wrap {
  background-clip: border-box;
  -webkit-text-fill-color: inherit;
  mix-blend-mode: normal;
}

.intro-header {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  height: 400px;
  margin-top: 0px;
  margin-bottom: 0px;
  padding-top: 0px;
  -webkit-justify-content: space-around;
  -ms-flex-pack: distribute;
  justify-content: space-around;
  -webkit-flex-wrap: wrap;
  -ms-flex-wrap: wrap;
  flex-wrap: wrap;
  -webkit-box-align: start;
  -webkit-align-items: flex-start;
  -ms-flex-align: start;
  align-items: flex-start;
  -webkit-align-content: space-around;
  -ms-flex-line-pack: distribute;
  align-content: space-around;
  background-color: transparent;
  color: #fff;
  // background-color: #54cde2;
}

.hero-heading-left {
  position: relative;
  padding: 80px 30px;
  border: 1px none #000;
  background-color: transparent;
}

.main-intro {
  // width: 100%;
  max-width: 940px;
  margin-right: auto;
  margin-left: auto;
  // max-width: 1200px;
}

// .hero-wrapper {
//   display: -webkit-box;
//   display: -webkit-flex;
//   display: -ms-flexbox;
//   display: flex;
//   -webkit-box-pack: justify;
//   -webkit-justify-content: space-between;
//   -ms-flex-pack: justify;
//   justify-content: space-between;
//   -webkit-box-align: center;
//   -webkit-align-items: center;
//   -ms-flex-align: center;
//   align-items: center;
// }

.hero-wrapper {
  margin-bottom: -10px;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
  -webkit-flex-direction: column;
  -ms-flex-direction: column;
  flex-direction: column;
  -webkit-box-pack: start;
  -webkit-justify-content: flex-start;
  -ms-flex-pack: start;
  justify-content: flex-start;
  -webkit-box-align: start;
  -webkit-align-items: flex-start;
  -ms-flex-align: start;
  align-items: flex-start;
}

.hero-split {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  max-width: 100%;
  margin-right: 0px;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
  -webkit-flex-direction: column;
  -ms-flex-direction: column;
  flex-direction: column;
  -webkit-box-pack: start;
  -webkit-justify-content: flex-start;
  -ms-flex-pack: start;
  justify-content: flex-start;
  -webkit-box-align: start;
  -webkit-align-items: flex-start;
  -ms-flex-align: start;
  align-items: flex-start;
  // background-color: #54cde2;
}

.paragraph {
  margin-top: 10px;
  color: #fff;
  font-size: 16px;
  font-weight: 500;
  text-transform: uppercase;
}

.paragraph-bigger {
  margin-bottom: 10px;
  opacity: 1;
  font-size: 20px;
  line-height: 34px;
  font-weight: 400;
  // font-size: 16px;
  // line-height: 28px;
}

.paragraph-bigger.cc-bigger-light {
  opacity: 0.6;
  font-family: Inconsolata, monospace;
  font-size: 14px;
  text-align: left;
  text-transform: none;
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

.shadow-two {
  position: static;
  -webkit-box-flex: 0;
  -webkit-flex: 0 auto;
  -ms-flex: 0 auto;
  flex: 0 auto;
  box-shadow: 0 4px 24px 0 rgba(150, 163, 181, 0.08);
}

.intro-detail {
  background-color: #000000;
  margin-bottom: 100px;
}

// .step {
//   margin-top: -95px;
// }

.w-container {
  margin-left: 32%;
  margin-right: auto;
  max-width: 1000px;
  background-color: transparent;
}

.w-container:before,
.w-container:after {
  content: ' ';
  display: table;
  grid-column-start: 1;
  grid-row-start: 1;
  grid-column-end: 2;
  grid-row-end: 2;
}

.w-container:after {
  clear: both;
}

// .w-container .w-row {
//   margin-left: -10px;
//   margin-right: -10px;
// }

.heading {
  color: #ffb134;
}

.heading-2 {
  color: #fff;
}

.paragraph-2 {
  color: #fff;
}

.bold-text {
  position: static;
  left: auto;
  display: block;
  overflow: visible;
  margin-left: auto;
  text-align: left;
  -o-object-fit: fill;
  object-fit: fill;
}

.text-block {
  margin-left: 0px;
  color: #fff;
  font-size: 24px;
}

.w-row:before,
.w-row:after {
  content: ' ';
  display: table;
  grid-column-start: 1;
  grid-row-start: 1;
  grid-column-end: 2;
  grid-row-end: 2;
}

.w-row:after {
  clear: both;
}

.w-row .w-row {
  margin-left: 0;
  margin-right: 0;
}

.w-col {
  position: relative;
  float: left;
  width: 100%;
  min-height: 1px;
  padding-left: 10px;
  padding-right: 10px;
  padding-bottom: 30px;
  padding-top: 30px;
  background-color: transparent;
}

.w-col .w-col {
  padding-left: 0;
  padding-right: 0;
}

.w-col-1 {
  width: 8.33333333%;
}

.w-col-2 {
  width: 16.66666667%;
}

.w-col-3 {
  width: 25%;
}

.w-col-4 {
  width: 33.33333333%;
}

.w-col-5 {
  width: 41.66666667%;
}

.w-col-6 {
  width: 50%;
}

.w-col-7 {
  width: 58.33333333%;
}

.w-col-8 {
  width: 66.66666667%;
}

.w-col-9 {
  width: 75%;
}

.w-col-10 {
  width: 83.33333333%;
}

.w-col-11 {
  width: 91.66666667%;
}

.w-col-12 {
  width: 100%;
}

.gradient {
  border: 3px solid #000;
}

.gradient:hover {
  background-image: url('@img/explore.png');
  background-position: 0px 0px;
  background-size: cover;
  color: #000;
}

.link {
  color: #fff;
}
</style>
