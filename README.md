# shping
make a point api for e-commerce application

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
