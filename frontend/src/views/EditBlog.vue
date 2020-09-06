<template>
  <div class="edit">
    <v-row>
      <v-select>Select Blog</v-select>
    </v-row>
    <v-row>
      <v-text-field label="Title" v-bind="blog_title"></v-text-field>
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
      <v-btn @click="submit">Save Article</v-btn>
    </v-row>
  </div>
</template>

<script>
import showdown from 'showdown'

let converter = new showdown.Converter()
export default {
  name: 'EditBlog',
  data() {
    return {
      blog_title: "",
      blog_markdown: "",
      blog_html: "",
      show_html: false
    }
  },
  methods: {
    write: function () {
      this.show_html = false
    },
    showHTML: function () {
      console.log(this)
      console.log(this.blog_markdown)
      this.blog_html = converter.makeHtml(this.blog_markdown)
      this.show_html = true
    },
    submit: function () {
      // Implement
    }
  }
}
</script>