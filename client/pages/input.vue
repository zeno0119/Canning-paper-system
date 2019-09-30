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
import axios from 'axios'

export default {
  data () {
    return {
      content: '',
      result: '',
      address: 'https://canning-zeno0119.herokuapp.com'
    }
  },
  methods: {
    post () {
      const param = new URLSearchParams()
      param.append('content', this.content)
      axios.post(this.address + '/api/post', param)
        .then((res) => {
          this.result = res.data.content
        })
    }
  }
}
</script>

<style scoped>

</style>
