<template>
  <div class="article">
    <div v-if="loading">Loading...</div>
    <div v-else>
      <div class="overline mb-4">{{ createdDate }}</div>
      <h2>{{ title }}</h2>
      <div v-html="content"></div>
    </div>
  </div>
</template>

<script>
import showdown from 'showdown'
import axios from 'axios'

let converter = new showdown.Converter()
export default {
  name: "Article",
  data() {
    return {
      createdDate: "",
      title: "",
      content: "",
      loading: true,
      errored: false
    }
  },
  mounted() {
    axios
        .get('/api/articles/' + this.$route.params.id)
        .then(response => {
          this.title = response.data.title;
          this.createdDate = response.data.createdDate;
          this.content = converter.makeHtml(response.data.content);
        })
        .catch(error => {
          console.log(error)
          this.errored = true
        })
        .finally(() => this.loading = false)
  }
}
</script>