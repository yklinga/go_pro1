<template>
  <div class="signup">
    <v-container>
      <v-card>
        <v-card-text>
          <v-form ref="form" v-model="valid">
            <v-text-field
              v-model="params.username"
              :rules="rules.username"
              :counter="10"
              label="username"
              required
            ></v-text-field>

            <v-text-field
              v-model="params.password"
              :rules="rules.password"
              type="password"
              label="password"
              required
            ></v-text-field>

            <v-text-field
              v-model="params.telephone"
              :rules="rules.telephone"
              :counter="11"
              label="telephone"
              required
            ></v-text-field>

            <v-btn
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="submit"
            >
              Sign up
            </v-btn>

            <v-btn color="error" class="mr-4" @click="reset"> Clear </v-btn>
          </v-form>
        </v-card-text>
      </v-card>
    </v-container>
  </div>
</template>

<script>
import * as loginApi from '../../api/login'
export default {
  name: 'signup',
  data: () => ({
    valid: true,
    params: {
      username: '',
      password: '',
      telephone: '',
    },
    rules: {
      username: [
        v => !!v || 'username is required',
        v => (v && v.length <= 10) || 'username must be less than 10 characters',
      ],
      password: [
        v => !!v || 'password is required',
        v => (v && v.length >= 6) || 'password must be more than 6 characters',
      ],
      telephone: [
        v => !!v || 'telephone is required',
        v => (v && v.length == 11) || 'telephone must be equal 11 characters',
      ]
    }
  }),
  methods: {
    submit() {
      this.$refs.form.validate();
      if (this.valid) {
        console.log('submit');
        this.login(this.params)
        loginApi.register(this.params).then(res => {
          console.log(res, 'res');
        }).catch(err => {
          console.log(err, 'err');
        })
      } else {
        console.log('form error');
      }
    },
    reset() {
      this.$refs.form.reset()
    }
  }
}
</script>