# Golang App

This Golang based application exposes APIs to add, query and delete employee data from Postgres database.

## APIs available

* GET - Ping to validate app - `http://localhost:4000/ping`
* POST - Add an Employee - `http://localhost:4000/crud/`
Body

```json
{
	"id": "1",
	"name": "User1"
}
```
* GET - Get all Employees - `http://localhost:4000/crud/`

Response

```json
{
    "data": [
        {
            "id": "2",
            "name": "User2"
        },
        {
            "id": "1",
            "name": "User1"
        }
    ]
}
```

* DELETE - Delete Employee by ID - `http://localhost:4000/crud/1`
