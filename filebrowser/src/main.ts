import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from '@/App.vue'
import router from '@/router'
import vuetify from '@/plugins/vuetify'
import axios from 'axios'
import VueAxios from 'vue-axios'
// 全局样式
import '@less/global.less'
import '@mdi/font/css/materialdesignicons.css'

const app = createApp(App)
app.use(createPinia()) // 启用 Pinia
app.use(router)
app.use(vuetify)
app.use(VueAxios, axios)
app.provide('axios', app.config.globalProperties.axios)
app.mount('#app')
