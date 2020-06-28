<template>
    <div id="note">
        <span style="display:none">{{index}}</span><span :style="{color:color}" id="nodecontent">{{content}}</span>
        <button @click="deleteNote" id="deleteButton">x</button>
    </div>
</template>
<script>
export default {
    name: 'note',
    props: ['index','content','color'],
    methods: {
        deleteNote() {
            this.$parent.delNote(this.index)
        }
    }
}
</script>
<style scoped>
#note{
    width: 400px;
    margin: auto
}
#deleteButton{
    float: right;
}
</style>