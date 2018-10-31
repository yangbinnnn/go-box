<template>
  <div class="main">
    <h1> {{ msg }} </h1>
    <p>Ping response: {{ resp }}</p>
    <p v-if="isConnected">We're connected to the ws server!</p>
    <p>WebSocket Message from server: {{ socketMessage }}</p>
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
      isConnected: false
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
</style>
