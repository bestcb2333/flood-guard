import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'element-plus/dist/index.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(ElementPlus, { locale: zhCn })
Object.keys(ElementPlusIconsVue).forEach(key => {
  app.component(key, (ElementPlusIconsVue as Record<string, any>)[key])
})

app.use(createPinia())
app.use(router)

app.mount('#app')
