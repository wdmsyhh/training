<template>
  <div id="root">
      <div id="pageTitle">Notes</div>
      <div v-for="(note, index) in notes" :key="index">
          <note :index="index" :content="note.text" :color="note.color" style="margin-top:10px"></note>
      </div>
      <div id="addNoteArea">
          <table>
            <tr>
                <td colspan="4"><textarea v-model="notetext" style="width:400px; height:50px"></textarea></td>
            </tr>
            <tr>
                <td><div id="redSelector" @click="setColor('red')"></div></td>
                <td><div id="blueSelector" @click="setColor('blue')"></div></td>
                <td><div id="greenSelector" @click="setColor('green')"></div></td>
                <td><button @click="addNote" id="addID">添加</button></td>
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
          this.notes.splice(i, 1)
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
        left: 40%;
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