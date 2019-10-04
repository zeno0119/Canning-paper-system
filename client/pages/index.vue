<template>
  <div>
    <div class="hero is-fullheight">
      <div class="hero-body" style="display: block; text-align: center;">
        <p class="title">
          {{ date }}
        </p>

        <p class="title" style="text-align: center; font-size: 100px; vertical-align: center">
          {{ content }}
        </p>
      </div>
    </div>
    <nuxt-link to="input">移動</nuxt-link>
  </div>
</template>

<script>
import {w3cwebsocket} from 'websocket'
import axios from 'axios'
const W3cWebsocket = w3cwebsocket

export default {
  data () {
    return {
      content: '',
      date: '',
      intervalId: undefined,
      socket: undefined,
    }
  },
  created () {
    this.socket = new W3cWebsocket('ws://' + 'localhost:8080' + '/ws')
    this.intervalId = setInterval(() => {
      this.date = new Date()
    }, 1000)

    this.socket.onmessage = (msg) => {
      this.content = msg.data
    }
    axios.get('/api/get')
      .then((res) => {
        this.content = res.data.content
        })
  },
  beforeDestroy () {
    clearInterval(this.intervalId)
  }
}
</script>

<style scoped>

</style>
