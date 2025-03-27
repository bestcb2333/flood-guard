import 'element-plus/dist/index.css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import '@/styles/styles.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import App from './App.vue'
import router from './router'
import {createI18n} from 'vue-i18n'

import zh from './i18n/zh'
import en from './i18n/en'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

const i18n = createI18n({
  locale: 'zh',
  fallbackLocale: 'en',
  messages: {
    en: en,
    zh: zh,
  },
})
app.use(i18n)

app.mount('#app')
