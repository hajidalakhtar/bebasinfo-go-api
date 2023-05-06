# bebasinfo


## Description

This repository is created as a solution for the 62teknologi-senior-backend-test.

### How To Run This Project

> Make sure to modify `config.json` first

#### Run the Applications

```bash
# Clone the repository
$ git clone https://github.com/hajidalakhtar/bebasinfo.git

# Navigate to the project directory
$ cd bebasinfo

# Install the dependencies and tidy up the go.mod file
$ go mod tidy

# Start the server 
$ make serve
```



### Endpoints
>`[GET] /token/generate` is used to generate a JWT Token.

* `[GET] /token/generate` 
* `[GET] /business_search/search`
* `[POST] /business_search`
* `[PUT] /business_search/:id`
* `[DEL] /business_search/:id`


Default port 3000.




