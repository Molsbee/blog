# Blog
I am designing this as a single tenant blogging solution that will meet my needs.



## Local Development not using docker-compose

#### Database Setup (Not mounting volume at this point)
```shell script
docker run --name blog-mysql -e MYSQL_ROOT_PASSWORD=blog-development -e MYSQL_DATABASE=blog -p "3306:3306" -d mysql:8.0.17
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