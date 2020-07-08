import Vue from 'vue'
import { Button, Input, Icon, List, Checkbox, Col, Row, message } from 'ant-design-vue'
import App from './App.vue'
import axios from 'axios'
import VueRouter from 'vue-router'
axios.defaults.withCredentials=true 


Vue.use(Button)
Vue.use(Input)
Vue.use(Icon)
Vue.use(List)
Vue.use(Checkbox)
Vue.use(Col)
Vue.use(Row)

Vue.config.productionTip = false
Vue.prototype.$axios = axios
Vue.prototype.$message = message
Vue.use(VueRouter)

import Login from './components/Login.vue'
import Plans from './components/Plans.vue'
import Register from './components/Register.vue'


//定义路由
const routes = [
  { path:'/', component: Login},
  { path: '/plans', component: Plans },
  { path: '/register', component: Register },
 ]
 //创建 router 实例，然后传 routes 配置
 const router=new VueRouter({
  routes
 });

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
