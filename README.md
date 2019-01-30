# Go-simple-CRUD
My Playground for learning Go Programming Language by making a simple CRUD using MySQL database

## Installation
1. Open localhost/PHPMyAdmin or any MySQL database management tool that you are using
2. Create a new database "simple"
3. Run the command below
``` bash
#clone repo
git clone https://github.com/HendraPB/Go-simple-CRUD.git

# change your directory to project directory
cd Go-simple-CRUD

# install project package
go get -u github.com/go-sql-driver/mysql
go get -u github.com/astaxie/beego
go get -u github.com/beego/bee

# Running migrations
bee migrate -driver=mysql -conn="<dbUser>:<dbPass>@tcp(<dbHost>:<dbPort>)/<dbName>"
# example
bee migrate -driver=mysql -conn="root:@tcp(127.0.0.1:3306)/simple"

# run project on localhost:8080
go run main.go
```
