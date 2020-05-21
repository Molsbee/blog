<template>
  <div class="home">
    <div v-if="loading">Loading...</div>
    <v-list-item v-else v-for="article in articles" v-bind:key="article" :href="'/details/' + article.id" >
      <v-list-item-content>
        <div class="overline mb-4">{{ article.createdDate }}</div>
        <v-list-item-title class="headline mb-1">{{ article.title }}</v-list-item-title>
        <v-list-item-subtitle>
          {{article.description}}
        </v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>
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