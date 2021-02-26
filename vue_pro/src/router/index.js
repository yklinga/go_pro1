import Vue from 'vue'
import VueRouter from 'vue-router'
import index from '../views/index.vue'
import loginRouter from '../views/loginAbout/router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'index',
    component: index
  },
  ...loginRouter,
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
