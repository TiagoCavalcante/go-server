# go-server
A server in Golang using fasthttp

## run
* run using local installed Golang:
  * install Golang
  * put the HTML files in the folder `docs`
  * execute the following command: `export PORT=xxxx`
  * execute the following command: `go run src/main.go`
* run using Docker
  * install Docker
  * put the HTML files in the folder `docs`
  * edit the `ports` in `docker-compose.yml`
  * exit environment variable `PORT` in `docker-compose.yml`
  * execute the following command: `docker-compose up --build && docker-compose down`