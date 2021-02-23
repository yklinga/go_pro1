<template>
  <v-app id="app">
    <v-app-bar app color="deep-purple accent-4">
      <!-- -->
      <h1 style="color: #fff">vueGo</h1>
      <v-spacer></v-spacer>
      <LoginControl />
    </v-app-bar>

    <!-- 根据应用组件来调整你的内容 -->
    <v-main
      ><!-- 给应用提供合适的间距 -->
      <v-container fluid>
        <!-- 如果使用 vue-router -->
        <router-view></router-view>
      </v-container>
    </v-main>

    <v-footer app>
      <!-- -->
    </v-footer>
  </v-app>
</template>

<script>
import LoginControl from "./components/LoginControl"
export default {
  components: {
    LoginControl
  }
}
</script>
<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
