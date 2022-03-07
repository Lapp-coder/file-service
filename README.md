# File-service

[![Go](https://img.shields.io/badge/go-1.17-blue)](https://golang.org/doc/go1.17)
[![Tagj](https://img.shields.io/badge/release-1.1.0-success)](https://github.com/Lapp-coder/file-service/releases)

***

### File service - storage of your files on a remote server
### Technologies/tools
* Go 1.17
* Fiber web-framework
* MinIO
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
    ```
* #### Run application in docker:
  ```
  $ docker-compose up --build
  ```
