// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueWebsocket from 'vue-native-websocket'

Vue.config.productionTip = false

Vue.use(VueWebsocket, `//${window.location.host}/ws`, {reconnection: true})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
