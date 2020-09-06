import Vue from 'vue'
import VueRouter from 'vue-router'
import vuetify from './plugins/vuetify';
import App from './App.vue'

Vue.config.productionTip = process.env.NODE_ENV == 'production'
Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import('./views/Home.vue')
    },
    {
        path: '/blog/login',
        name: 'Login',
        component: () => import('./views/Login.vue')
    },
    {
        path: '/blog/about',
        name: 'About',
        component: () => import('./views/About.vue')
    },
    {
        path: '/admin/create',
        name: 'Write',
        component: () => import('./views/WriteBlog.vue')
    },
    {
        path: '/admin/edit',
        name: 'Edit',
        component: () => import('./views/EditBlog.vue')
    },
    {
        path: '/blog/article/:id',
        name: 'Article',
        component: () => import('./views/Article.vue')
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})
export default router

new Vue({
    router,
    vuetify,
    render: h => h(App),
    created() {
        this.$vuetify.theme.dark = true
    }
}).$mount('#app')
