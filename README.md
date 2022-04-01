# Go CRUD API Banking App

Banking App CRUD API with Go Language using gorilla/mux and Postman

#### Required: Go, and Postman (or curl, insomnia, etc)

----
1. Open Terminal and go to project directory
2. Type `go mod init example.com/go-crud-bankingapp`
3. Install gorilla/mux, type `go get "github.com/gorilla/mux"`
4. Type `go run main.go`
5. Open Postman, and try to send some request

---
#### URI
> localhost:8080/customers

#### Requests
| Method | Endpoint   | Description |
|--------|------------|-------------|
| GET    | /customers | Get All Customers            |
| GET    | /customers/{id}   | Get Customers by ID            |
| POST   | /customers       | Create Account            |
| PUT    | /customers/{id}           | Update Account            |
| DELETE | /customers/{id}           | Delete Account by ID            |

```azure
{
    "id": "",
    "customer": {
        "account_number": ,
        "first_name": "",
        "last_name": "",
        "sex": ""
    },
    "balance": 
}
```
