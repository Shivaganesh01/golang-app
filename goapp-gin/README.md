Save Video

curl -X POST -d '{"title":"Video-1", "description":"videos", "url":"https://www.youtube.com/embed/sDJLQMZzzM4", "creator":{"firstname":"Shivaganesh", "lastname":"B", "age":10, "email":"sg@gmail.com"}}' http://localhost:8080/api/video -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiU2hpdmFnYW5lc2giLCJpc0FkbWluIjp0cnVlLCJleHAiOjE2MTQ2OTM2OTgsImlhdCI6MTYxNDQzNDQ5OCwiaXNzIjoic2hpdmFnYW5lc2hAZ21haWwuY29tIn0.Qz1tmZz-UAKUAP-Yu1fMzGhlGm_7h9pCWWpZgXZmeB4"


Find all

curl localhost:8080/api/v1/videos -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiU2hpdmFnYW5lc2giLCJpc0FkbWluIjp0cnVlLCJleHAiOjE2MTQ2OTM2OTgsImlhdCI6MTYxNDQzNDQ5OCwiaXNzIjoic2hpdmFnYW5lc2hAZ21haWwuY29tIn0.Qz1tmZz-UAKUAP-Yu1fMzGhlGm_7h9pCWWpZgXZmeB4"

Login:
curl -X POST -d '{"username":"Shivaganesh", "password":"password"} localhost:8080/login


3. Data binding and Validation


GORM - Object Relation Mapping Using Golang
It supports associations, hooks, Pre-loading, transactions, composite primary key, SQL Builder, auto migrations, pugins based on GORM callbacks, GORM supports SQLite, MySQL, Postgres and SQL Server

Swagger
Open source tools that can help in design, build, document and consume RESt APIs.
