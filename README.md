# Blog
I am designing this as a single tenant blogging solution that will meet my needs.
It will be backed by PostgreSQL and deployed to Heroku.

## Local Development not using docker-compose

#### Database Setup (Not mounting volume at this point)
```shell script
docker run --name postgres -e POSTGRES_USER=blog -e POSTGRES_PASSWORD=blog-development -e POSTGRES_DB=blog -p "5432:5432" -d postgres:12
```

#### Database Migration (MacOS) - github.com/golang-migrate/migrate/tree/master/cmd/migrate
Install Migration Command Line Tool
```shell script
brew install golang-migrate
```
Create Migration (assumes you are running in project directory)
```shell script
migrate create -ext sql -dir ./database-migrations -seq create_articles_table
```

General Notes 
Running Go Locally
```shell script
yarn --cwd ./frontend/ build && go run main.go
```

Docker
```shell script
docker build -t blog .
docker run --name blog -d blog
```

Vue Setup Take so far
```shell script
vue create frontend
vue add router
vue add vuetify
yarn add axios
yarn --cwd ./frontend/ build
```