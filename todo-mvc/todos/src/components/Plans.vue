<template>
    <div class="root">
        <div class="pageTitle">
            todos
        </div>
        <div class="addPlanArea">
            <a-input v-model="planContent" placeholder="What needs to be done?" size="large" @pressEnter="addPlan(planContent)">
                <a-icon slot="prefix" type="down" @click="checkAll()"/>
            </a-input>
            <a-list v-if="plans != null && plans.length != 0" item-layout="horizontal" :data-source="plans">
                <a-list-item v-if="(status == 'all') || (status == 'active' && plan.isCompleted == false) || (status == 'completed' && plan.isCompleted == true)" id="listItem" class="item" slot="renderItem" slot-scope="plan, index"  @mouseenter="showDeleteButton(index)" @mouseleave="hideDeleteButton()" @dblclick="switchInputBox()">
                    <a-checkbox v-if="!(isEditable && index==current)" :checked=plan.isCompleted @change="changeCompletedStatusOrContent($event, plan._id, plan)">{{ plan.content }}</a-checkbox>
                    <div v-if="(isEditable && current==index)">
                        <a-input v-model="plan.content" style="width: 400px;" @pressEnter="changeCompletedStatusOrContent(null, plan._id, plan)"></a-input>
                    </div>
                    <span v-show="seen && index==current" style="padding-right:10px">
                        <a-icon slot="prefix" type="close" @click="deleteOne(plan._id)"/>
                    </span>
                </a-list-item>
            </a-list>
            <div class="underItemsArea">
                <a-row>
                    <a-col :span="6">
                        <div style="padding-top:10px; padding-left: 10px; float: left">{{ activeCount }} items left</div>
                    </a-col>
                    <a-col :span="12">
                        <div>
                            <a-button size="small" style="float: left; margin-top: 10px" @click="showPlans('all')">All</a-button>
                            <a-button size="small" style="margin-top: 10px" @click="showPlans('active')">Active</a-button>
                            <a-button size="small" style="float: right; margin-top: 10px" @click="showPlans('completed')">Completed</a-button>
                        </div>
                    </a-col>
                    <a-col :span="6">
                        <div>
                            <a-button size="small" style="float: right; margin:10px 5px;" @click="clearCompletedPlans()">Clear completed</a-button>
                        </div>
                    </a-col>
                </a-row>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'Plans',
    data() {
      return {
          seen: false,
          current: '',
          planContent: '',
          plans: [],
          ids: [],
          status: 'all',
          activeCount: 0,
          isEditable: false,
      }
    },
    methods: {
        //添加一条数据
        addPlan(content) {
            var _this = this
            axios.post('/plans', {
                'content': content,
                'isCompleted': false,
            }).then(function(response) {
                if(response.data.code == 1) {
                    _this.planContent = ''
                    _this.getPlans()
                }
                if(response.data.code == 0) {
                    _this.$message.error(response.data.message)
                }
            })
        },
        //删除一条数据
        deleteOne(id) {
            var _this = this
            axios.delete('/plans/'+id).then(function(response){
                if(response.data.code == 1) {
                    _this.getPlans()
                }
                if(response.data.code == 0) {
                    _this.$message.error(response.data.message)
                }
            })
        },
        //清除已完成的
        clearCompletedPlans() {
            var _this = this
            axios.post('/plans/batchDelete').then(function(response) {
                if(response.data.code == 1) {
                    _this.getPlans()
                }
                if(response.data.code == 0) {
                    _this.$message.error(response.data.message)
                }
            })
        },
        //鼠标移动到一条数据显示删除按钮
        showDeleteButton(index) {
            this.seen = true
            this.current = index
        },
        hideDeleteButton() {
            this.seen = false
            this.isEditable = false
        },
        //修改完成状态
        changeCompletedStatusOrContent(e, id, plan) {
            var _this = this
            if(e != null && e.target.checked == true) {
                this.activeCount--
            }
            if(e != null && e.target.checked == false) {
                this.activeCount++
            }
            //发起请求
            if(e != null) {
                plan.isCompleted = e.target.checked
            } 
            axios.put('/plans/'+id, plan).then(function(response) {
                if(response.data.code == 1) {
                    _this.getPlans()
                    _this.isEditable = false
                }
                if(response.data.code == 0) {
                    _this.$message.error(response.data.message)
                }
            })
        },
        //标记全部为已完成或未完成
        checkAll() {
            var _this = this
            axios.post('/plans/batchUpdate', { 'activeCount': _this.activeCount} ).then(function(response) {
                if(response.data.code == 1) {
                    _this.getPlans()
                }
                if(response.data.code == 0) {
                     _this.$message.error(response.data.message)
                }
            })
        },
        //用于列表下方按钮控制显示已完成的、未完成的或全部
        showPlans(status) {
            this.status = status
        },
        //双击变成输入框
        switchInputBox() {
            this.isEditable = true
        },
        //查询所有
        getPlans() {
            var _this = this
            this.activeCount = 0
            axios.get('/plans').then(function(response) {
                if(response.data.code == 1) {
                    _this.plans = response.data.data
                    //计算 active plan 的数量
                    for(var i = 0; i < _this.plans.length; i++) {
                        if(_this.plans[i].isCompleted == false) {
                            _this.activeCount++
                        }
                    }
                }
                if(response.data.code == 0) {
                    _this.$router.push({ path: '/' })
                }
            })
        }
    },
    created() {
       this.getPlans()
    },
}
</script>

<style scoped>
    .root{
        width: 100%;
        height: 100%;
        position: absolute;
        background-color: #f5f5f5
    }
    .pageTitle{
        font-size: 100px;
        color: lightpink
    }
    .addPlanArea{
        width: 600px;
        margin: auto;
        border: 1px solid seashell;
    }
    .item{
        background-color: white;
        padding-left: 10px
    }
    .underItemsArea{
        width: 600px;
        height: 40px;
        border: 1px solid seashell;
        background-color: white;
    }
</style>