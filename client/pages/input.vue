<template>
  <div class="hero">
    <div class="hero-body">
      <input v-model="content" type="text" style="height: 100px; width: 10%">
      <button style="height: 100px;" @click="post">
        カンペ投稿
      </button>
      <p class="title">
        {{ result }}
      </p>
    </div>
  </div>
</template>

<script>
import {w3cwebsocket} from 'websocket'
const W3cWebsocket = w3cwebsocket

export default {
  data () {
    return {
      content: '',
      result: '',
      socket: ''
    }
  },
  methods: {
    post () {
      this.socket.send(this.content)
      this.result = this.content
      this.content = ''
    }
  },
  created(){
    this.socket = new W3cWebsocket('ws://' + "canning-zeno0119.herokuapp.com" + "/ws")
  }
}
</script>

<style scoped>

</style>
