### Database Setup
```
docker pull mysql:8.4
docker run --name ecomm-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:8.4
```
Pull and run the mysql image in a detached mode (`-d`) with port-forwarding and an environment variable.
```
docker exec -i ecomm-mysql mysql -uroot -ppassword <<< "CREATE DATABASE ecomm;"
```
Make the `CREATE DATABASE` query in the mysql container with `docker exec`.
```
docker run -it --rm --network host --volume ./db:/db migrate/migrate:v4.17.0 -path=/db/migrations -database "mysql://root:password@tcp(localhost:3306)/ecomm" up
```
Apply the migration files in `db/migrations` by running the [golang-migrate](https://github.com/golang-migrate/migrate) image with the `up` command.

### Run the Go Apps
```
go run cmd/ecomm-grpc/main.go
go run cmd/ecomm-api/main.go
```
