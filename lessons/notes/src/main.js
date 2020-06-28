import Vue from 'vue'
import { Button, Field, Dialog, Col, Row } from 'vant'
import App from './App.vue'

Vue.use(Button)
Vue.use(Field)
Vue.use(Dialog)
Vue.use(Col)
Vue.use(Row)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
