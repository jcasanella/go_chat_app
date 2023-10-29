# Chat App

## Docker postgres

```
cd docker
docker-compose up -d
```

This spawn a `postgres` instance with the db `chat_development`

## How to run a Migration

### Migrate installation
```
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ xenial main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

### Create a migration file

```
migrate create -ext sql -dir db/migrations -seq create_users_table
migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/chat_development?sslmode=disable" -verbose up
migrate -path db/migrations -database "postgresql://postgres:mysecretpassword@localhost:5432/chat_app" -verbose up
```

### How to connect to Postgres and list the tables

* Connect to Postgres

```
docker exec -it docker-db-1 psql -h localhost -U postgres
```

* List the schemas

```
postgres=# \l
                                                    List of databases
       Name       |  Owner   | Encoding |  Collate   |   Ctype    | ICU Locale | Locale Provider |   Access privileges
------------------+----------+----------+------------+------------+------------+-----------------+-----------------------
 chat_development | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            |
 postgres         | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            |
 template0        | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            | =c/postgres          +
                  |          |          |            |            |            |                 | postgres=CTc/postgres
 template1        | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            | =c/postgres          +
                  |          |          |            |            |            |                 | postgres=CTc/postgres
(4 rows)
```

* Connect into Schema **chat_development**
```
postgres=# \c chat_development
You are now connected to database "chat_development" as user "postgres".
```

* List tables
```
chat_development=# \dt
               List of relations
 Schema |       Name        | Type  |  Owner
--------+-------------------+-------+----------
 public | schema_migrations | table | postgres
 public | users             | table | postgres
(2 rows)

```



## References

* Migrations: https://github.com/golang-migrate/migrate
* How to Encrypt Passwords in Golang: https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go