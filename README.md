# File-service

[![Go](https://img.shields.io/badge/go-1.17-blue)](https://golang.org/doc/go1.17)
[![Tagj](https://img.shields.io/badge/tag-1.0.0-success)](https://github.com/Lapp-coder/file-service/tags)

***

### File service - storage of your files on a remote server
### Technologies/tools
* Go 1.17
* Fiber web-framework
* MinIO
* PostgreSQL
* Git
* Docker & docker-compose

### For run this application:
* ```
  $ git clone github.com/Lapp-coder/file-service && cd file-service/
  ```
* #### Create an .env file in the root directory of the project with the following contents:
    ```
    MINIO_ACCESS_KEY=<access-key>
    MINIO_SECRET_KEY=<secret-key>
    POSTGRES_PASSWORD=<password>
    ```
* #### Run application in docker:
  ```
  $ docker-compose up --build
  ```
* #### After launching, apply migrations to the database(must be installed [golang-migrate](https://github.com/golang-migrate/migrate)): 
  ```
  $ export POSTGRES_HOST=<host> \
  POSTGRES_PORT=<port> \
  POSTGRES_USER=<username> \
  POSTGRES_PASSWORD=<password> \
  POSTGRES_DB=<db_name> \
  POSTGRES_USE_SSL=<enable/disable>
  ```
  ```
  $ make migrate-up
  ```
