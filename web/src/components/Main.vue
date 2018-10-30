<template>
  <div class="main">
    <h1> {{ msg }} </h1>
    <p>Ping response: {{ resp }}</p>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      msg: 'go-box web',
      resp: '...'
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
  }
}
</script>

<style scoped>
  .main {
    width: 480px;
    margin: 0 auto;
  }
</style>
