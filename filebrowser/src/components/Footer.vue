<template>
  <div id="container">
    <div id="contacts">
      <div id="contacts--contact">
        <p>creatorsuite@next.id</p>
        <p>This is footer</p>
      </div>
      <div id="contacts--social">
        <div>
          <a href="/">
            <img
              src="@img/twitter.png"
              width="30"
              alt=""
              sizes="(max-width: 30px) 30vw, 30px"
              class="logo-link w-inline-block w--current"
            />
          </a>
        </div>
        <div>
          <a href="/">
            <img
              src="@img/telegram.png"
              width="30"
              alt=""
              sizes="(max-width: 30px) 30vw, 30px"
              class="logo-link w-inline-block w--current"
            />
          </a>
        </div>
        <div>
          <a href="/">
            <img
              src="@img/nextid.png"
              width="30"
              alt=""
              sizes="(max-width: 30px) 30vw, 30px"
              class="logo-link w-inline-block w--current"
            />
          </a>
        </div>
      </div>
    </div>
    <div id="mentions">
      <p>2022 | Powered by Next.ID</p>
    </div>
  </div>
</template>

<style lang="less" scoped>
body {
  display: flex;
  background-color: #f0f8ff;
  height: 95vh;
  justify-content: center;
  align-items: flex-end;
  // font-family: Open Sans;
}

#container {
  padding-top: 20px;
  margin-top: 0px;
  width: 100vw;
  color: white;
  background-color: #202020;
  display: flex;
  flex-direction: column;
  align-items: center;
}
#contacts {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 100%;
}
#contacts--contact {
  text-align: right;
}
#contacts--social {
  width: 20%;
  display: flex;
  justify-content: space-between;
}
#contacts--social div {
  height: 30px;
  width: 30px;
  background-color: white;
  border-radius: 50%;
}
#mentions p {
  font-size: 0.5em;
}

.section {
  // margin-right: 30px;
  // margin-left: 30px;
  background-color: red;
}

.w-layout-grid {
  display: -ms-grid;
  display: grid;
  grid-auto-columns: 1fr;
  -ms-grid-columns: 1fr 1fr;
  grid-template-columns: 1fr 1fr;
  -ms-grid-rows: auto auto;
  grid-template-rows: auto auto;
  grid-row-gap: 16px;
  grid-column-gap: 16px;
}

.footer {
  padding-top: 120px;
  padding-bottom: 60px;
  justify-items: start;
  -webkit-box-align: start;
  -webkit-align-items: start;
  -ms-flex-align: start;
  align-items: start;
  grid-column-gap: 30px;
  grid-row-gap: 30px;
  grid-template-areas:
    '. . . . .'
    '. . . . .';
  -ms-grid-columns: 2fr 30px 1fr 30px 1fr 30px 1.25fr 30px 1fr;
  grid-template-columns: 2fr 1fr 1fr 1.25fr 1fr;
  -ms-grid-rows: auto 30px auto;
  grid-template-rows: auto auto;
}

.footer-link {
  margin-bottom: 15px;
  opacity: 0.6;
  -webkit-transition: color 200ms ease, opacity 200ms ease;
  transition: color 200ms ease, opacity 200ms ease;
  color: #fff;
  font-size: 14px;
  text-decoration: none;
  background-color: #a7a7a7;
}

.footer-link:hover {
  opacity: 1;
}

.footer-link.w--current {
  color: #fff;
}

.logo-link {
  z-index: 1;
}

.logo-link:hover {
  opacity: 0.8;
}

.logo-link:active {
  opacity: 0.7;
}

.container {
  width: 100%;
  max-width: 1140px;
  margin-right: auto;
  margin-left: auto;
  color: #fff;
}
.container.cc-subscription-form {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  margin-bottom: auto;
  padding-top: 130px;
  padding-bottom: 140px;
  -webkit-box-orient: vertical;
  -webkit-box-direction: normal;
  -webkit-flex-direction: column;
  -ms-flex-direction: column;
  flex-direction: column;
  -webkit-box-pack: center;
  -webkit-justify-content: center;
  -ms-flex-pack: center;
  justify-content: center;
  -webkit-box-align: center;
  -webkit-align-items: center;
  -ms-flex-align: center;
  align-items: center;
  background-color: transparent;
  text-align: center;
}

.container.cc-heading-wrap {
  margin-top: 60px;
  margin-bottom: 80px;
  text-align: center;
}

.container.cc-full-width-form-wrap {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  -webkit-box-align: start;
  -webkit-align-items: flex-start;
  -ms-flex-align: start;
  align-items: flex-start;
}

.w-inline-block {
  max-width: 100%;
  display: inline-block;
}

.w--current {
  opacity: 1;
  font-weight: 600;
  letter-spacing: 1px;
}

.logo-footer {
  margin-bottom: 70px;
}

.links-section-footer {
  margin-top: 10px;
  margin-bottom: 40px;
}

.label {
  margin-bottom: 10px;
  color: #fff;
  font-size: 12px;
  line-height: 20px;
  font-weight: 500;
  letter-spacing: 1px;
  text-transform: uppercase;
}

.label.cc-styleguide-label {
  margin-bottom: 25px;
}

.label.cc-cart-remove-link {
  margin-top: 4px;
  margin-bottom: 5px;
  color: #a7a7a7;
  font-weight: 400;
  text-decoration: none;
}

.label.cc-blog-date {
  margin-top: 20px;
  opacity: 0.6;
}

.made-with-webflow {
  display: -webkit-box;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  margin-top: 70px;
  -webkit-box-align: center;
  -webkit-align-items: center;
  -ms-flex-align: center;
  align-items: center;
  opacity: 0.5;
  -webkit-transition: opacity 400ms ease;
  transition: opacity 400ms ease;
  text-decoration: none;
}

.made-with-webflow:hover {
  opacity: 1;
}

.made-with-webflow:active {
  opacity: 0.8;
}

.webflow-logo-tiny {
  margin-top: -2px;
  margin-right: 8px;
}
.logo-image {
  display: block;
}
</style>
