# female-daily

# How To Run API

* Create db on mysql server
```
CREATE DATABASE `femaledaily`;
```

* Run migration with cmd
```
migrate -database 'mysql://root:root@tcp(localhost:3306)/registration?parseTime=true' -path ./db/migrations up
```

* Run App
```
go run .
```

* cURL Example
- Fetch
```
curl --location --request GET 'http://localhost:8080/api/v1/fetch?page=1'
```

- Create
```
curl --location --request POST 'http://localhost:8080/api/v1/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"test123@gmail.com",
    "first_name":"test",
    "last_name":"satu",
    "avatar":"http://test.com"
}'
```

- Update 
```
curl --location --request PUT 'http://localhost:8080/api/v1/user/7' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"test123@gmail.com",
    "first_name":"test",
    "last_name":"satu dua",
    "avatar":"http://test.com"
}'
```

- Delete
```
curl --location --request DELETE 'http://localhost:8080/api/v1/user/7' \
--header 'Authorization: 3cdcnTiBsl'
```

- Get By ID
```
curl --location --request GET 'http://localhost:8080/api/v1/user/1'
```

- Get List
```
curl --location --request GET 'http://localhost:8080/api/v1/user?last_name=Ramos&first_name=Tracey'
```