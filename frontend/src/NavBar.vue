<script setup lang="ts">
import { reactive, ref } from 'vue'
import {useI18n} from 'vue-i18n';
import { useRouter } from 'vue-router'
import useSessionStore from './stores/session';
import usePersistedStore from './stores/persisted';
import {apiAxios} from './axios';

const routes = useRouter().options.routes

const isCollapse = ref(false)
const {t} = useI18n({messages: {
  zh: {
    dashboard: '仪表盘',
    resources: '救援资源管理',
    histories: '历史降水数据',
    sensors: '传感器管理',
    events: '内涝事件管理',
    users: '用户管理',
    wikis: '知识库',
    deepseek: 'DeepSeek',
    settings: '系统设置',
    loginSignup: '登录/注册',
    logout: '退出登录',
  },
}})

const session = useSessionStore()
const persisted = usePersistedStore()

const isDialogOpen = ref(false)
const currentTab = ref('login')

const loginForm = reactive({
  username: '',
  password: '',
})

async function login() {
  try {
    persisted.token = await apiAxios.post<any, string>('/login', loginForm)
    await session.loadUser()
    isDialogOpen.value = false
  } catch {}
}

const signupForm = reactive({
  email: '',
  authcode: '',
  username: '',
  password: '',
})

async function signup() {
  try {
    persisted.token = await apiAxios.post<any, string>('/signup', signupForm)
    await session.loadUser()
    isDialogOpen.value = false
  } catch {}
}

const retrieveForm = reactive({
  email: '',
  authcode: '',
  password: '',
})

async function retrieve() {
  try {
    persisted.token = await apiAxios.post<any, string>('/retrieve', retrieveForm)
    await session.loadUser()
    isDialogOpen.value = false
  } catch {}
}

async function logout() {
  persisted.token = null
  session.user = null
}
</script>

<template>
  <el-menu
    class="min-h-full flex flex-col bg-transparent border-0 gap-4 px-2"
    router
    :collapse="isCollapse"
    :default-active="'/' + $route.path.split('/')[1]"
  >

    <el-menu-item class="font-bold text-2xl mt-4">
      <span>
        FloodGuard
      </span>
    </el-menu-item>

    <template v-for="route in routes" :key="route.path">
      <el-menu-item v-if="route.meta?.icon" :index="route.path" class="rounded-full">
        <el-icon>
          <component :is="route.meta.icon" />
        </el-icon>
        <span>{{t(route.meta.name as string)}}</span>
      </el-menu-item>
    </template>

    <el-menu-item class="mt-auto">
      <el-button v-if="session.user" type="primary" @click="logout" round>
        {{t('logout')}}
      </el-button>
      <el-button v-else type="primary" @click="isDialogOpen=true" round>
        {{t('loginSignup')}}
      </el-button>
    </el-menu-item>

    <el-menu-item @click="isCollapse = !isCollapse">
      <el-icon>
        <Hide v-if="isCollapse" />
        <View v-else />
      </el-icon>
      <template #title>
        {{ $t('navbar.hide') }}
      </template>
    </el-menu-item>
  </el-menu>
  <el-dialog v-model="isDialogOpen" title="欢迎使用Climate Sentinel">
    <el-tabs v-model="currentTab">
      <el-tab-pane label="登录" name="login">
        <el-form :model="loginForm" label-width="auto">
          <el-form-item label="用户名">
            <el-input v-model="loginForm.username" />
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="loginForm.password" />
          </el-form-item>
          <el-form-item>
            <el-button class="ms-2" type="primary" @click="login">
              登录
            </el-button>
            <el-button @click="isDialogOpen=false">
              返回
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="注册" name="signup">
        <el-form :model="signupForm" label-width="auto">
          <el-form-item label="邮箱">
            <el-input v-model="signupForm.email" />
          </el-form-item>
          <el-form-item label="邮箱验证码">
            <el-input v-model="signupForm.authcode" />
          </el-form-item>
          <el-form-item label="用户名">
            <el-input v-model="signupForm.username" />
          </el-form-item>
          <el-form-item label="密码">
            <el-input v-model="signupForm.password" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" class="ms-auto" @click="signup">
              注册
            </el-button>
            <el-button @click="isDialogOpen=false">
              返回
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="找回密码" name="retrieve">
        <el-form :model="retrieveForm" label-width="auto">
          <el-form-item label="邮箱">
            <el-input v-model="retrieveForm.email" />
          </el-form-item>
          <el-form-item label="验证码">
            <el-input v-model="retrieveForm.authcode" />
          </el-form-item>
          <el-form-item label="新密码">
            <el-input v-model="retrieveForm.password" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="retrieve" class="ms-auto">
              找回密码
            </el-button>
            <el-button @click="isDialogOpen=false">
              返回
            </el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
    </el-tabs>
  </el-dialog>
</template>
