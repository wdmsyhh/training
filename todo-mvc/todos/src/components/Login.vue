<template>
  <div class="root">
      <div class="loginFormArea">
          <a-row>
              <a-col :span="6"></a-col>
              <a-col :span="12">
                <table style="margin:auto;">
                    <tr>
                        <td>用户名：</td>
                        <td><a-input v-model="userName"></a-input></td>
                    </tr>
                    <tr>
                        <td>密码：</td>
                        <td><a-input-password v-model="password"></a-input-password></td>
                    </tr>
                    <tr>
                        <td></td>
                        <td>
                            <a-button style="float: left;" type="primary" @click="login()">登录</a-button>
                            <a-button size="small" @click="gotoRegisterPage()">注册</a-button>
                        </td>
                    </tr>
                </table>
              </a-col>
              <a-col :span="6"></a-col>
          </a-row>
      </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
     data() {
      return {
          userName: '',
          password: '',
      }
    },
    methods: {
        gotoRegisterPage() {
            this.$router.push({ path: '/register' })
        },
        login() {
            var _this = this
            axios.post('/user/login',{
                'userName': _this.userName,
                'password': _this.password,
            }).then(function(response) {
                if(response.data.data != null) {
                    //跳转到首页
                    _this.$router.push({ path: '/plans' })
                } else {
                    _this.$message.error('请确认用户名密码是否正确！')
                }
            })
        }
    }
}
</script>

<style scoped>
    table{
        border-collapse: separate; 
        border-spacing: 15px;
        width: 400px;
        height: 200px;
    }
    .loginFormArea{
        height: 300px;
        margin-top: 260px;
        background-color: cadetblue
    }
    .root{
        width: 100%;
        height: 100%;
        position: absolute;
        background-color: darkcyan
    }
</style>