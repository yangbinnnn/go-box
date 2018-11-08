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
      <button v-on:click="login">登录</button>
      <button v-on:click="info">个人信息</button>
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
axios.defaults.baseURL = (process.env.NODE_ENV !== 'production') ? 'http://127.0.0.1:8000' : ''

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
      }).catch(error => (
        this.consoleMsg = error.response.data
      ))
    },
    login (email) {
      axios.post('/api/user/login', {
        'email': this.email,
        'name': this.name
      }).then(response => {
        this.consoleMsg = response.data
        setTimeout(this.info, 3)
      }).catch(error => {
        this.consoleMsg = error.response.data
      })
    },
    info (email) {
      axios.get('/api/user/info')
        .then(response => {
          let user = response.data.user
          this.name = user.name
          this.email = user.email
          this.registerTime = user.registerTime
          this.consoleMsg = response.data
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
  .user, .userInfo {
    height: 120px;
    background-color: gray;
  }

  .console {
    height: 120px;
    border: solid black;
  }
</style>
