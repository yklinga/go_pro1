<template>
  <div class="signup">
    <v-form ref="form" v-model="valid" lazy-validation>
      <v-text-field
        v-model="params.username"
        :rules="rules.username"
        :counter="10"
        label="username"
        required
      ></v-text-field>

      <v-text-field
        v-model="params.telephone"
        :rules="rules.telephone"
        :counter="11"
        label="telephone"
        required
      ></v-text-field>

      <v-btn color="success" class="mr-4"> Validate </v-btn>

      <v-btn color="error" class="mr-4"> Reset Form </v-btn>

      <v-btn color="warning"> Reset Validation </v-btn>
    </v-form>
  </div>
</template>
<script>
export default {
  name: 'signup',
  data: () => ({
    valid: true,
    params: {
      username: '',
      telephone: '',
    },
    rules: {
      username: [
        v => !!v || 'username is required',
        v => (v && v.length <= 10) || 'username must be less than 10 characters',
      ],
      telephone: [
        v => !!v || 'telephone is required',
        v => (v && v.length <= 11) || 'telephone must be less than 11 characters',
      ]
    }
  }),
  methods: {}
}
</script>