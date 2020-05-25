<template>
    <v-app-bar app fixed>
        <v-toolbar-title>William Molsbee</v-toolbar-title>
        <v-spacer></v-spacer>
        <!-- TODO: Implement search and improve search feature -->
<!--        <v-btn icon>-->
<!--            <v-icon>mdi-magnify</v-icon>-->
<!--        </v-btn>-->
        <v-menu left bottom>
            <template v-slot:activator="{ on }">
                <v-btn icon v-on="on">
                    <v-icon>mdi-dots-vertical</v-icon>
                </v-btn>
            </template>
            <v-list>
                <v-list-item to="/">
                    <v-list-item-title>Home</v-list-item-title>
                </v-list-item>
                <v-list-item to="/blog/about">
                    <v-list-item-title>About</v-list-item-title>
                </v-list-item>
                <v-list-item v-show="!isAuthenticated" to="/blog/login">
                    <v-list-item-title>Login</v-list-item-title>
                </v-list-item>
                <v-list-item v-show="isAuthenticated" @click="logout">
                    <v-list-item-title>Logout</v-list-item-title>
                </v-list-item>
            </v-list>
        </v-menu>
    </v-app-bar>
</template>

<script>
    import axios from 'axios'

    export default {
        name: 'NavBar',
        data() {
            return {
                isAuthenticated: false,
                username: null
            }
        },
        mounted() {
            this.isAuthenticated = localStorage.getItem("username") != null;
            this.username = localStorage.getItem("username")
            this.$root.$on('LoginEvent', (text) => {
                this.isAuthenticated = true
                this.username = text
            })
        },
        methods: {
            logout: function() {
                axios.get("/logout")
                    .then(() => {
                        localStorage.clear()
                        this.isAuthenticated = false
                        this.username = null
                    })
                    .catch(error => {
                        console.log(error)
                    })
            }
        }
    }
</script>