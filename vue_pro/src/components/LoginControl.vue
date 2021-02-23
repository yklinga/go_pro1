<template>
  <div>
    <div class="unlogin" v-if="!isLogin">
      <v-btn text color="shades white" @click="handleSignIn">sign in</v-btn>
      <v-btn text color="shades white" @click="handleSignUp">sign up</v-btn>
    </div>
    <div v-else>
      <v-btn text color="shades white" @click="handleSignOut">sign out</v-btn>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LoginControl',

  data() {
    return {
    }
  },

  methods: {
    handleSignIn() {
      this.$router.push({ path: '/signin' });
    },
    handleSignUp() {
      this.$router.push({ path: '/signup' });
    },
    handleSignOut() {
      localStorage.removeItem('token');
    }
  },
  computed: {
    isLogin() {
      if (localStorage.getItem('token')) {
        return true;
      }
      return false;
    }
  }
}
</script>
