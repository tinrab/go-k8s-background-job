import Vue from 'vue';
import VueNativeSock from 'vue-native-websocket';

import App from './App.vue';

Vue.config.productionTip = false;

Vue.use(VueNativeSock, 'ws://localhost:8081/ws', {
  format: 'json',
  reconnection: true,
  reconnectionAttempts: 5,
  reconnectionDelay: 1000,
});

new Vue({
  render: h => h(App),
}).$mount('#app');
