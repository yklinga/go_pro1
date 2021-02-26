import Vue from 'vue'

export const register = (data) => Vue.axios({
  url: '/auth/register',
  method: 'post',
  data: data
})

export const login = (data) => Vue.axios({
  url: '/auth/login',
  method: 'post',
  data: data
})

export const getUserInfo = () => Vue.axios({
  url: '/auth/userinfo',
  method: 'get'
})