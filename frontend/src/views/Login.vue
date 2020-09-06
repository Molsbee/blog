<template>
  <div class="login">
    <v-container>
      <v-row align="center" justify="center">
        <v-col cols="12" sm="8" md="4">
          <v-card class="elevation-12">
            <v-toolbar>
              <v-toolbar-title>Login Form</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              <v-form>
                <v-text-field
                    label="Username"
                    name="username"
                    v-model="username"
                    prepend-icon="mdi-account-tie"
                    type="text"
                ></v-text-field>
                <v-text-field
                    label="Password"
                    name="password"
                    v-model="password"
                    prepend-icon="mdi-lock"
                    type="password"
                ></v-text-field>
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-card-text v-show="errored" class="error--text">Failed to login user</v-card-text>
              <v-btn color="primary" @click="submit">Login</v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "Login",
  data() {
    return {
      username: "",
      password: "",
      errored: false
    }
  },
  methods: {
    submit: function () {
      this.errored = false
      let formData = new FormData()
      formData.set("username", this.username)
      formData.set("password", this.password)
      axios
          .post("/login", formData)
          .then(response => {
            // Emit Login Event and Redirect
            this.$root.$emit('LoginEvent', response.data.username)
            this.$router.push('/')
          })
          .catch(error => {
            console.log(error)
            this.errored = true
          })

    }
  }
}
</script>
