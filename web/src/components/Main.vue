<template>
  <div class="main">
    <h1> {{ msg }} </h1>
    <p>Ping response: {{ resp }}</p>
    <p v-if="isConnected">We're connected to the ws server!</p>
    <p>WebSocket Message from server: {{ socketMessage }}</p>

    <div class="user">
      Email: <input type="email" name="email" id="email" v-model="email">
      <br>
      Name: <input type="text" name="name" id="name" v-model="name">
      <br>
      <button v-on:click="register">注册</button>
    </div>

    <div class="userInfo">
      <p>Email: {{ email }}</p>
      <p>Name: {{ name }}</p>
      <p>RegisterTime: {{ registerTime }}</p>
    </div>

    <div class="console">
      <p> {{ consoleMsg }} </p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      msg: 'go-box web',
      resp: 'waiting...',
      socketMessage: 'waiting...',
      isConnected: false,

      email: '',
      name: '',
      registerTime: '',
      consoleMsg: ''
    }
  },
  mounted () {
    this.ping()
  },
  methods: {
    ping () {
      axios.get('/api/ping').then(response => (
        this.resp = response.data
      )).catch(error => console.log(error))
    },
    register () {
      axios.post('/api/user/register', {
        'email': this.email,
        'name': this.name
      }).then(response => {
        this.consoleMsg = response.data
        setTimeout(this.userInfo, 3, [this.email])
      }).catch(error => (
        this.consoleMsg = error.response.data
      ))
    },
    userInfo (email) {
      axios.get('/api/user/info?email=' + email)
        .then(response => {
          this.name = response.data.name
          this.email = response.data.email
          this.registerTime = response.data.registerTime
        })
        .catch(error => (
          this.consoleMsg = error.response.data
        ))
    }
  },
  sockets: {
    onopen () {
      this.isConnected = true
    },
    onclose () {
      this.isConnected = false
    },
    onmessage (msg) {
      this.socketMessage = msg.data
    }
  }
}
</script>

<style scoped>
  .main {
    width: 960px;
    margin: 0 auto;
  }

  .user .userInfo {
    width: 960px;
    height: 480px;
    background-color: gray;
    border-width: 2px;
    border-color: gray;
  }
</style>
