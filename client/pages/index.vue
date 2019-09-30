<template>
  <div>
    <div class="hero is-fullheight" @click="get">
      <div class="hero-body" style="display: block; text-align: center;">
        <p class="title">
          {{ date }}
        </p>

        <p class="title" style="text-align: center; font-size: 100px; vertical-align: center">
          {{ content }}
        </p>
      </div>
    </div>
    <nuxt-link to="./input">移動</nuxt-link>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      content: '',
      address: 'https://canning-zeno0119.herokuapp.com',
      date: '',
      intervalId: undefined
    }
  },
  created () {
    this.intervalId = setInterval(() => {
      this.date = new Date()
      this.get()
    }, 1000)
  },
  beforeDestroy () {
    clearInterval(this.intervalId)
  },
  methods: {
    get () {
      axios.get(this.address + '/api/get')
        .then((res) => {
          this.content = res.data.content
        })
    }
  }
}
</script>

<style scoped>

</style>
