# book-management-system
Book Management System implemented using GO

## Dev Setup 
### Docker command to run mysql container and create books database
```bash
docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql
docker exec -it some-mysql bash
mysql -u root -p
CREATE DATABASE books;
```

### Export env variables
```bash
export DB_URL=127.0.0.1:3306
export DB_USERNAME=root
export DB_PASSWORD=my-secret-pw
```

### Run service
```bash
go run cmd/main.go
```

## Postman Collection
Find the postman collection of API's under test/ directory