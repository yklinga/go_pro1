import Vue from 'vue'
import Vuex from 'vuex'
import * as loginApi from '../api/login'
Vue.use(Vuex)

export default new Vuex.Store({
  // 数据状态
  state: {
    token: '',
    userinfo: {}
  },
  // 赋值方法
  mutations: {
    // 设置token
    setToken: (state, token) => {
      if (token) {
        localStorage.setItem('token', token);
      } else {
        localStorage.removeItem('token');
      }
      state.token = token;
    },
    setUserInfo: (state, userinfo) => state.userinfo = userinfo,
  },
  // 
  actions: {
    async getUserInfo(commit) {
      const result = await loginApi.getUserInfo();
      if (!result) return;
      const { code, data } = result;
      if (code === 200) {
        commit('setUserInfo', data);
        commit('setToken', data.token);
      }
    }
  },
  modules: {
  }
})
