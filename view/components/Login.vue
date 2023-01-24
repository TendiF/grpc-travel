<template>
  <div>
    <input type="text" placeholder="username" v-model="username"/>
    <input type="password" placeholder="password" v-model="password"/>
    <button @click="doLogin">Login</button>

  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { userLoginRequest, userLoginResponse } from '../../proto/users_pb'


export default Vue.extend({
  name: 'Login',
  data: () => {
    return {
      username : '',
      password: ''
    }
  },
  mounted(){
  },
  methods: {
    async doLogin(){
      if(!this.username || !this.password){
        alert('fill username & password')
        return
      }
      let request = new userLoginRequest()
      request.setUsername(this.username)
      request.setPassword(this.password)
      let response = await <Promise<userLoginResponse>>this.$store.dispatch('login', request)
      if(response){
        this.$store.commit('setToken', response.getToken())
        this.$router.push('/customers')
      }
    }
  }
})
</script>
