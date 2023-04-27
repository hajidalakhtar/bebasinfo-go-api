# 62teknologi-senior-backend-test-Muhammad-Hajid-Al-Akhtar


## Description

This repository is intended as an answer to 62teknologi-senior-backend-test challenge.


### How To Run This Project

> Make Sure you have run the article.sql in your mysql

Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH.

#### Run the Testing

```bash
$ make tests
```

#### Run the Applications

Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into your workspace
$ git clone https://github.com/bxcodec/go-clean-arch.git

#move to project
$ cd go-clean-arch

# Run the application
$ make up

# The hot reload will running

# Execute the call in another terminal
$ curl localhost:9090/articles
```
