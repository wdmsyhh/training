<template>
  <div class="root">
      <div class="registerFormArea">
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
                        <td>确认密码：</td>
                        <td><a-input-password v-model="confirmPassword"></a-input-password></td>
                    </tr>
                    <tr>
                        <td></td>
                        <td>
                            <a-button style="float: left;" type="primary" @click="register()">提交</a-button>
                            <a-button size="small" @click="gotoLoginPage()">返回登录页面</a-button>
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
          confirmPassword: '',
      }
    },
    methods: {
        gotoLoginPage() {
            this.$router.push({ path: '/' })
        },
        register() {
            if(this.password != this.confirmPassword) {
                this.$message.error('两次输入密码不一致！')
                return
            }
            var _this = this
            axios.post('/user/register',{
                'userName': _this.userName,
                'password': _this.password,
            }).then(function(response) {
                if(response.data.code == 1) {
                    _this.$message.info('注册成功，请登录')
                    _this.$router.push({ path: '/' })
                }
                if(response.data.code == 0) {
                    _this.$message.error(response.data.message)
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
    .registerFormArea{
        height: 300px;
        margin-top: 260px;
    }
    .root{
        width: 100%;
        height: 100%;
        position: absolute;
        background-color: darkcyan
    }
</style>