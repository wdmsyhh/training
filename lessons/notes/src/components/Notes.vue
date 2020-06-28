<template>
  <div id="root">
      <van-row id="pageTitle">
          <van-col span="8"></van-col>
          <van-col span="8">Notes</van-col>
          <van-col span="8"></van-col>
      </van-row>
      <div v-for="(note, index) in notes" :key="index">
          <note :index="index" :content="note.text" :color="note.color" style="margin-top:10px"></note>
      </div>
      <div id="addNoteArea">
          <table style="margin: auto">
            <tr>
                <td colspan="4">
                <van-field
                    v-model="notetext"
                    rows="2"
                    autosize
                    label="note"
                    type="textarea"
                    maxlength="50"
                    placeholder="请输入 note"
                    show-word-limit
                    style="border: 1px solid grey; width: 400px"/>
                </td>
            </tr>
            <tr>
                <td><div id="redSelector" @click="setColor('red')"></div></td>
                <td><div id="blueSelector" @click="setColor('blue')"></div></td>
                <td><div id="greenSelector" @click="setColor('green')"></div></td>
                <td><van-button type="primary" size="small" @click="addNote">添加</van-button></td>
            </tr>
          </table>
      </div>
  </div>
</template>

<script>

import note from './note.vue'

export default {
  name: 'Notes',
  props: {
    msg: String
  },
  components: {
    note
  },
  data() {
      return {
          notes: [{text: 'note1', color: 'black'},{text: 'note2', color: 'red'}],
          notetext: '',
          color: ''
      }
  },
  methods: {
      addNote() {
          if(this.notetext == '') return
          this.notes.push({text: this.notetext, color: this.color})
      },
      setColor(colorVal) {
          this.color = colorVal
      },
      delNote(i) {
          this.$dialog.confirm({
            title: '删除操作',
            message: '是否确定要删除该 note ?',
            })
            .then(() => {
                this.notes.splice(i, 1)
            })
            .catch(() => {
                return
            });
      }
  },
  watch: {
      notes(){
          this.color = 'black'
      }  
  }
}
</script>

<style scoped>
    #root{
        width: 100%;
        height: 100%;
        position: absolute
    }
    #addNoteArea{
        position: fixed;
        bottom: 50px;
        width: 100%;
    }
    #redSelector{
        width: 30px;
        height: 30px;
        background-color: red;
    }
    #blueSelector{
        width: 30px;
        height: 30px;
        background-color: blue
    }
    #greenSelector{
        width: 30px;
        height: 30px;
        background-color: green
    }
    #pageTitle{
        font-size: 80px;
        background-color: azure
    }
</style>>