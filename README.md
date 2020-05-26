# Blog (William Molsbee)
I am developing/designing this as a single tenant blogging solution that will meet my needs.
I am deploying it to Heroku and backing it by its free PostgreSQL add-on.
Eventually this will move to a different hosted solution after development has concluded.
The articles for the blog will be wrote in markdown format, and the UI will render it as HTML.

App: https://william-molsbee.herokuapp.com/

## API
The REST API will have Basic Authentication for some endpoints documented below.

| HTTP METHOD | Secured | URI                      |
|-------------|:-------:|:-------------------------|
| GET         | No      | /api/articles            |
| GET         | No      | /api/articles/:articleID |
| POST        | Yes     | /api/articles            |
| PUT         | Yes     | /api/articles/:articleID |

## Local Development

### Database Migrations 
Database migrations are bing managed by golang-migrate.
There is a CLI for creating files in the migration directory (MacOS) - github.com/golang-migrate/migrate/tree/master/cmd/migrate

Install CLI (MacOS)
```shell script
brew install golang-migrate
```

Create Migration Example (assumes you are running in project directory)
```shell script
migrate create -ext sql -dir ./database-migrations -seq create_articles_table
```

### Deploying application with Docker Compose
```shell script
docker-compose build
docker-compose up
docker-compose down
```

### Running Locally without Docker Compose
Setup Postgres database for application to consume (example using docker without mounting external volume)
```shell script
docker run --name postgres -e POSTGRES_USER=blogger -e POSTGRES_PASSWORD=password -e POSTGRES_DB=blog -p "5432:5432" -d postgres:12
```
Build UI Assets and run go code.  You can run yarn build at any time to recompile UI Assets, and the go project will serve the new assets.
```shell script
yarn --cwd ./frontend/ build && go run main.go
```

#### Notes on project setup
Vue Setup
```shell script
vue create frontend
vue add router
vue add vuetify
yarn add axios
yarn --cwd ./frontend/ build
```

Heroku Commands
```shell script
heroku container:login
heroku container:push web -a william-molsbee
heroku container:release web -a william-molsbee 
```

Docker
```shell script
docker build -t blog .
docker run --name blog -d blog
```


Cleanup docker-compose
Create Make File to automate some things and stuff like heroku deployment.