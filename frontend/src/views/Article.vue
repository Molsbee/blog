<template>
    <div class="article">
        <div v-if="loading">Loading...</div>
        <div v-else>
            <div class="overline mb-4">{{ article.createdDate }}</div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios'
    export default {
        name: "Article",
        data() {
            return {
                article: null,
                loading: true,
                errored: false
            }
        },
        mounted() {
            console.log(this.$route.params.id);
            axios
                .get('/api/articles/' + this.$route.params.id)
                .then(response => {
                    this.article = response.data
                })
                .catch(error => {
                    console.log(error)
                    this.errored = true
                })
                .finally(() => this.loading = false)
        }
    }
</script>