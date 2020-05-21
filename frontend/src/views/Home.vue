<template>
  <div class="home">
    <div v-if="loading">Loading...</div>
    <div v-else v-for="article in articles" v-bind:key="article">
      <h2>
        <router-link tag="div" to="/details">{{ article.title }}</router-link>
      </h2>

      <!-- Figure out -->
      {{ article.author }}
      <blockquote>
        {{ article.description }}
      </blockquote>
    </div>
  </div>
</template>

<script>
  import axios from 'axios'
  export default {
    name: "Home",
    data() {
      return {
        articles: [],
        loading: true,
        errored: false
      }
    },
    mounted() {
      axios
              .get('http://localhost:8080/articles')
              .then(response => {
                this.articles = response.data
              })
              .catch(error => {
                console.log(error)
                this.errored = true
              })
              .finally(() => this.loading = false)
    }
  }
</script>