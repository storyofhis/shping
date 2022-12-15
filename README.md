# Shping with Go (Golang) REST API
make a point api for e-commerce application

you can [see](https://shping-production.up.railway.app/v1/)

Used libraries:
- [gin](https://github.com/gin-gonic)
- [gorm](https://gorm.io/docs/)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv?tab=doc)
- [validator](github.com/go-playground/validator/v10)

## Run Locally
Create `.env` at root, i.e.
```
DATABASE_URL=postgresql://${{ PGUSER }}:${{ PGPASSWORD }}@${{ PGHOST }}:${{ PGPORT }}/${{ PGDATABASE }}
PGHOST=your_local
PGPORT=5432
PGUSER=postgres
PGPASSWORD=your_password
PGDATABASE=your_db_name
PORT=8080
```

Setup Db after create database in your postgres
```
CREATE TYPE role AS ENUM ('admin', 'customer')
```

Run 
```
go run cmd/main.go
```

Routes for API
> all of. local routes can change with this [link](https://shping-production.up.railway.app/v1/)
- `user` routes
  - `POST` for register
  ```
  localhost:8080/v1/users/register
  https://shping-production.up.railway.app/v1/users/register
  ```
  - `POST` for login
  ```
  localhost:8080/v1/users/login
  https://shping-production.up.railway.app/v1/users/login
  ```
  - `PATCH` for edit topup user
  ```
  localhost:8080/v1/users/topup
  https://shping-production.up.railway.app/v1/users/topup
  ```
  
- `category` routes
  - `POST` for create category
  ```
  localhost:8080/v1/categories
  https://shping-production.up.railway.app/v1/categories
  ```
  - `GET` for get all category
  ```
  localhost:8080/v1/categories
  https://shping-production.up.railway.app/v1/categories
  ```
  - `PATCH` for edit category
  ```
  localhost:8080/v1/categories/:categoryId
  https://shping-production.up.railway.app/v1/categories/:categoryId
  ```
  - `DELETE`for delete category
  ```
  localhost:8080/v1/categories/:categoryId
  https://shping-production.up.railway.app/v1/categories/:categoryId
  ```
  
- `product` routes
  - `POST` for create product
  ```
  localhost:8080/v1/products
  https://shping-production.up.railway.app/v1/products
  ```
  - `GET` for get all products
  ```
  localhost:8080/v1/products
  https://shping-production.up.railway.app/v1/products
  ```
  - `PUT` for edit product
  ```
  localhost:8080/v1/products/:productId
  https://shping-production.up.railway.app/v1/products/:productId
  ```
  - `DELETE` for delete product
  ```
  localhost:8080/v1/products/:productId
  https://shping-production.up.railway.app/v1/products/:productId
  ```
  
- `transaction` routes
  - `POST` for create transaction
  ```
  localhost:8080/v1/transactions
  https://shping-production.up.railway.app/v1/transactions
  ```
  - `GET` for get all transaction
  ```
  localhost:8080/v1/transactions/my-transactions
  https://shping-production.up.railway.app/v1/transactions/my-transactions
  ```
  - `GET` for get all transaction with users response
  ```
  localhost:8080/v1/transactions/user-transactions
  https://shping-production.up.railway.app/v1/transactions/user-transactions
  ```

Jobdesk Member
- `MAULA IZZA AZIZI` (GLNG-KS04-020) :
  ```
  Initialize Project
  User Endpoints: Register and Login
  Transactions Endpoints: Get My Transactions and Get User Transactions
  Deployment
  ```
- `HEZKYA NATANAEL RAMLI` (GLNG-KS04-008) :
  ```
  Endpoint User Topup
  All Products Endpoint
  Endpoint Create Transaction
  Unit Test
  Mocking
  ```
- `MUHAMAD RESTU FADILLAH` (GLNG-KS04-002) :
  ```
  All Categories EndPoint
  Postman Collection
  ```
