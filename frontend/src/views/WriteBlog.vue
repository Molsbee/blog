<template>
  <div class="create">
    <v-row>
      <h3>Create a new blog page</h3>
    </v-row>
    <v-row>
      <v-text-field label="Title" v-model="title"></v-text-field>
    </v-row>
    <v-row>
      <v-text-field label="Author" v-model="author"></v-text-field>
      <v-spacer></v-spacer>
      <v-checkbox label="Publish" v-model="publish"></v-checkbox>
    </v-row>
    <v-row>
      <v-tabs>
        <v-tab @click="write">Write</v-tab>
        <v-tab @click="showHTML">Preview</v-tab>
      </v-tabs>
    </v-row>
    <v-row v-show="!show_html" style="height: 100%">
      <v-textarea
          outlined
          name="article"
          v-model="blog_markdown"
          rows="18"
      ></v-textarea>
    </v-row>
    <v-row v-show="show_html" style="height: 100%">
      <div v-html="blog_html"></div>
    </v-row>
    <v-row>
      <v-spacer></v-spacer>
      <v-card-text v-show="errored" class="error--text">Failed to save article</v-card-text>
      <v-btn @click="submit">Save Article</v-btn>
    </v-row>
  </div>
</template>

<script>
import showdown from 'showdown'
import axios from 'axios'

let converter = new showdown.Converter()
export default {
  name: 'WriteBlog',
  data() {
    return {
      title: "",
      author: "",
      blog_markdown: "",
      blog_html: "",
      show_html: false,
      errored: false,
      publish: false
    }
  },
  methods: {
    write: function () {
      this.show_html = false
    },
    showHTML: function () {
      this.blog_html = converter.makeHtml(this.blog_markdown)
      this.show_html = true
    },
    submit: function () {
      axios
          .post("/articles", {
            title: this.title,
            content: this.blog_markdown,
            author: this.author,
            published: this.publish
          })
          .then(() => {
            this.$router.push('/')
          })
          .catch(error => {
            console.log(error)
            this.errored = true
          })
    }
  }
}
</script>