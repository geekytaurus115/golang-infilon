### Starting the Task

#### Project Structure

```
golang-infilon
│
├── cmd/
│   └── main.go
│
├── handler/
│   ├── config.go
│   ├── controllers.go
│   ├── database.go
│   ├── model.go
│   ├── query.go
│   └── routes.go
│
└── go.mod
|__ GolangTaskInfilon.postman_collection.json

```

##### Explanation

- cmd/main.go starting point
- config.go: contains the configuration for this project, like db creds
- controllers.go: contains the functions with business logic
- database.go: contains the connection logic for the database
- model.go: contains the struct for table schema
- query.go: contains the queries for fetching and writing to the database
- route.go: contains the routes

- GolangTaskInfilon.postman_collection.json: contains postman tested collection

#### Install Packages

```
go get -u github.com/gin-gonic/gin
go get -u github.com/go-sql-driver/mysql
go get github.com/go-sql-driver/mysql

```

#### MySQL Commands

```
create database cetec;
show databases;
use cetec;
```

##### Create Person Table

```
CREATE TABLE person (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    age INT
);

SHOW TABLES;

INSERT INTO person (name, age) VALUES
    ('mike', 31),
    ('John', 20),
    ('Joseph', 20);

SELECT * FROM person;

```

##### Create Phone Table

```

CREATE TABLE phone (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    number VARCHAR(255),
    person_id INT,
    FOREIGN KEY (person_id) REFERENCES person(id)
);

INSERT INTO phone (id, person_id, number) VALUES
    (1, 1, '444-444-4444'),
    (8, 2, '123-444-7777'),
    (3, 3, '445-222-1234');

SELECT * FROM phone;

```

##### Create Address Table

```

CREATE TABLE address (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    city VARCHAR(255),
    state VARCHAR(255),
    street1 VARCHAR(255),
    street2 VARCHAR(255),
    zip_code VARCHAR(255)
);

INSERT INTO address (id, city, state, street1, street2, zip_code) VALUES
    (1, 'Eugene', 'OR', '111 Main St', '', '98765'),
    (2, 'Sacramento', 'CA', '432 First St', 'Apt 1', '22221'),
    (3, 'Austin', 'TX', '213 South 1st St', '', '78704');

SELECT * FROM address;

```

##### Create Address_Join Table

```
CREATE TABLE address_join (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    person_id INT,
    address_id INT,
    FOREIGN KEY (person_id) REFERENCES person(id),
    FOREIGN KEY (address_id) REFERENCES address(id)
);


INSERT INTO address_join (id, person_id, address_id) VALUES
    (1, 1, 3),
    (2, 2, 1),
    (3, 3, 2);

SELECT * FROM address_join;
```

#### Start Application

```
cd cmd
go run main.go
```
