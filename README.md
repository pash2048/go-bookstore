# go-bookstore
Simple book store using go &amp; mysql

## start project
Simply go to ```cmd/main``` with your terminal then start project with ```go run .``` or ```go buil .``` then ```./main``` in linux machine or ```main.exe``` in windows machine

## setup MYSQL
You can start mysql using docker with this command:
```docker run --name go-bookstore-mysql -e MYSQL_ROOT_PASSWORD=ROOT_PASSWORD -e MYSQL_DATABASE=DATABASE_NAME -p 3306:3306 --rm -d mysql```
and repalce ```ROOT_PASSWORD``` and ```DATABASE_NAME``` with your password and DB name.

then in ```pkg/config/app.go``` at line 13 change mysql connection paramters

## change port
You can Change http port in ```cmd/main/main.go``` at line 16
