<script setup lang="ts">
import CardTitle from '@/components/CardTitle.vue'
import axios from 'axios'
import { SHA256 } from 'crypto-js'
import { reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { FormRules } from 'element-plus'
import { useI18n } from 'vue-i18n'
import type { User } from './tables'
import useSessionStore from './stores/session'
import usePersistedStore from './stores/persisted'
import { apiAxios } from './axios'

const { t } = useI18n()
const session = useSessionStore()
const persisted = usePersistedStore()
const router = useRouter()
const route = useRoute()
const activeTab = ref('login')

// 登录表单
const loginForm = reactive({
  username: '',
  password: '',
})

// 登录校验规则
const loginRules = reactive<FormRules<Record<string, string>>>({
  username: [{ required: true, message: t('login.rules.username'), trigger: 'blur' }],
  password: [{ required: true, message: t('login.rules.password'), trigger: 'blur' }],
})

// 发送登录请求
function login() {
  apiAxios
    .post<string>('/login', {
      username: loginForm.username,
      password: SHA256(loginForm.password).toString(),
    })
    .then((res) => {
      localStorage.setItem('token', res.data)
    })
    .catch((err) => console.log(err))

  apiAxios
    .get<User>('/myinfo')
    .then((res) => {
      session.user = res.data
    })
    .catch((err) => console.log(err))

  router.push(route.params.redirect as string)
}

// 注册表单
const signupForm = reactive({
  username: '',
  password: '',
  rePassword: '',
  emailNumber: '',
  emailCode: '',
  captchaId: '',
  captchaValue: '',
})

// 发送注册请求
function signup() {
  apiAxios
    .post<string>(
      '/signup',
      {
        username: signupForm.username,
        password: SHA256(signupForm.password).toString(),
      },
      {
        headers: {
          'X-Captcha-Id': signupForm.captchaId,
          'X-Captcha-Value': signupForm.captchaValue,
          'X-Email-Number': signupForm.emailNumber,
          'X-Email-Code': signupForm.emailCode,
        },
      },
    )
    .then((res) => {
      localStorage.setItem('token', res.data)
    })
    .catch((err) => console.log(err))

  apiAxios
    .get<User>('/myinfo')
    .then((res) => {
      session.user = res.data
    })
    .catch((err) => console.log(err))

  router.push(route.params.redirect as string)
}

const captchaUrl = ref('')
async function loadCaptcha() {
  try {
    const resp = await axios.get(`${persisted.setting.apiAddr}/captcha`, {
      responseType: 'blob',
    })
    const blob = new Blob([resp.data], { type: resp.headers['Content-Type'] as string })
    captchaUrl.value = URL.createObjectURL(blob)
    signupForm.captchaId = resp.headers['X-Captcha-Id']
  } catch (err) {
    console.log(err)
  }
}

loadCaptcha()

// 找回密码表单
const retrieveForm = reactive({
  email: '',
  authcode: '',
  password: '',
  rePassword: '',
})

// 发送找回密码请求
function retrieve() {
  const { password, ...formData } = retrieveForm
  apiAxios
    .post('/retrieve', {
      ...formData,
      password: SHA256(retrieveForm.password).toString(),
    })
    .catch((err) => console.log(err))
}
</script>

<template>
  <el-row class="h-100">
    <el-col :xs="0" :md="12"> </el-col>
    <el-col :xs="24" :md="12" class="flex-col justify-center align-center">
      <el-card class="w-60">
        <card-title title="user.title" class="mb-4">
          <User />
        </card-title>
        <el-tabs v-model="activeTab">
          <!-- 登录页面 -->
          <el-tab-pane :label="$t('user.login.title')" name="login">
            <el-form :model="loginForm" label-width="auto" :rules="loginRules">
              <el-form-item :label="$t('user.login.username')" prop="username">
                <el-input v-model="loginForm.username" />
              </el-form-item>
              <el-form-item :label="$t('user.login.password')" prop="password">
                <el-input v-model="loginForm.password" />
              </el-form-item>
              <el-form-item>
                <el-button @click="login">{{ $t('user.login.title') }}</el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>

          <!-- 注册页面 -->
          <el-tab-pane :label="$t('user.signup.title')" name="signup">
            <el-form :model="signupForm" label-width="auto">
              <el-form-item :label="$t('user.signup.username')">
                <el-input v-model="signupForm.username" />
              </el-form-item>
              <el-form-item :label="$t('user.signup.email')">
                <el-input v-model="signupForm.emailNumber" />
              </el-form-item>
              <el-form-item :label="$t('user.signup.authcode')">
                <el-input v-model="signupForm.emailCode">
                  <template #append>
                    <el-button @click="apiAxios.get(`/mail?email=${signupForm.emailNumber}`)">
                      {{ $t('user.signup.send') }}
                    </el-button>
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item :label="$t('user.signup.password')">
                <el-input v-model="signupForm.password" />
              </el-form-item>
              <el-form-item :label="$t('user.signup.rePassword')">
                <el-input v-model="signupForm.rePassword" />
              </el-form-item>
              <el-form-item :label="$t('user.signup.captcha')">
                <el-input v-model="signupForm.captchaValue">
                  <template #append>
                    <img
                      :src="captchaUrl"
                      style="height: 30px; width: 80px"
                      :alt="$t('user.signup.captcha')"
                      @click="loadCaptcha"
                    />
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item>
                <el-button @click="signup">{{ $t('user.signup.title') }}</el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>

          <!-- 找回密码页面 -->
          <el-tab-pane :label="$t('user.retrieve.title')" name="retrieve">
            <el-form :model="retrieveForm" label-width="auto">
              <el-form-item :label="$t('user.retrieve.email')">
                <el-input v-model="retrieveForm.email" />
              </el-form-item>
              <el-form-item :label="$t('user.retrieve.authcode')">
                <el-input v-model="retrieveForm.authcode">
                  <template #append>
                    <el-button>{{ $t('user.retrieve.send') }}</el-button>
                  </template>
                </el-input>
              </el-form-item>
              <el-form-item :label="$t('user.retrieve.rePassword')">
                <el-input v-model="retrieveForm.rePassword" />
              </el-form-item>
              <el-form-item :label="$t('user.retrieve.rePassword')">
                <el-input v-model="retrieveForm.rePassword" />
              </el-form-item>
              <el-form-item>
                <el-button @click="retrieve">{{ $t('user.retrieve.title') }}</el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-col>
  </el-row>
</template>

<style scoped>
.w-60 {
  width: 60%;
}
</style>
