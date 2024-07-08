# Japhy Backend Test in Golang

## Technical Stack
- Go
- Docker
- MySQL

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/products/docker-desktop/)
- [Git](https://git-scm.com/downloads)

## Tasks
you are a backend developer in a pet food startup. A new functionality will be implemented, we want to be able to easily manage the breeds of dogs and cats that are registered in the database in our back office,
you must implement a CRUD api to manage the breeds of pets. The breeds are stored in a CSV file located at `./breeds.csv`. 
the aim of this test is to demonstrate backend development skills using the Go programming language. The application implements a simple REST API for managing resources
you are free to take initiatives and make improvements to the codebase.
Have fun and good luck!

### you need to implement the following tasks:
- create a new table in the core database to store the breeds of pets, to do this you must create a new migration file in the `database_actions` directory.
- store the breeds of pets in the database (list of breeds are on `breeds.csv`).
- implement CRUD functionality for the breeds resource (GET, POST, PUT, DELETE).
- search functionality to filter breeds based on pet characteristics (weight and species).


## Installation

1. Fork the project repository
2. Copy the `.env.example` file to `.env`
3. Build the application `docker compose build`
4. Run docker compose to start the application `docker compose up -d`
5. Once the application is up and running, you can access the REST API at http://localhost:50010. Use tools like Postman or curl to interact with the API.
6. `curl -v http://localhost:50010/health` to ensure your application is running.
7. send us the link to your repository with the api.



# Developer feedback

Hello, I'm writing this note to explain my work.

## Modification of the configuration

As I was unable to run the project according to the procedure, I took the liberty of making a few changes:

### ARM / AMD

My environment is under AMD64 architecture, so I created the `docker-compose-amd64.yml` file in which I changed the `platform` to make it compatible with my machine.

### mysqld.sock error

For some reason, the command to start the mysql service with `/entrypoint.sh mysqld` was causing a persistent error on my environment:

```
mysql-test    | '/var/lib/mysql/mysql.sock' -> '/var/run/mysqld/mysqld.sock'
mysql-test    | mysql: [Warning] Using a password on the command line interface can be insecure.
mysql-test    | ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/var/run/mysqld/mysqld.sock' (2)
mysql-test exited with code 1
```

I decided not to override the `mysql:8.0.36-oracle` image entrypoint to let it initialise as expected and to move the `cors` database creation script into the `.docker-compose/mysql/init.sql` file.

This file is mapped to the `/docker-entrypoint-initdb.d/` folder (line: `.docker-compose/mysql/init.sql:/docker-entrypoint-initdb.d/init.sql:ro` in the docker-compose) and as specified in the image doc: all `.sql` files in the folder are executed the first time the container is launched.


I've included all these changes in the original `docker-compose.yml` and I hope everything will work as expected, but I don't have an ARM machine to try.

## Installation

I have not made a Go binary, the installation has not changed:

1. Build the application `docker compose build`
2. Run docker compose to start the application `docker compose up -d`
3. Once the application is up and running, the routes can be tested using Swagger at http://localhost:50010/swagger/index.html or the classic tools.

I have published the Postman documentation at this link: https://documenter.getpostman.com/view/36773640/2sA3e1BqHf

Best regards,

guillrak

