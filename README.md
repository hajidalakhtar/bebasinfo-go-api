# 62teknologi-senior-backend-test-Muhammad-Hajid-Al-Akhtar


## Description

This repository is created as a solution for the 62teknologi-senior-backend-test.

### How To Run This Project

> Make sure to modify `config.json` first

#### Run the Applications

```bash

# Move to the directory called "workspace"
$ cd workspace

# Clone the repository into your workspace
$ git clone https://github.com/hajidalakhtar/62teknologi-senior-backend-test-muhammad-hajid-al-akhtar.git

# Navigate to the project directory
$ cd 62teknologi-senior-backend-test-muhammad-hajid-al-akhtar

# Install the dependencies and tidy up the go.mod file
$ go mod tidy

# Start the server 
$ make serve
```



### Example
Endpoints:

* `[GET] /token/generate` 
* `[GET] /business_search/search`
* `[POST] /business_search`
* `[PUT] /business_search/:id`
* `[DEL] /business_search/:id`

>`[GET] /token/generate` is used to generate a JWT Token.

Default port 3000.




