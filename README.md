# GWYN
## GO Starter ##
[GIN](https://github.com/gin-gonic/gin) go framework router handler
[Testify](https://github.com/stretchr/testify) go test toolkit 
[Mockery](https://github.com/vektra/mockery) auto generator mock interface


### RUN DEVELOPMENT ###
```sh run``` 

### RUN BUILD ####
###### docker ######
```docker compose up -d --build```
open http://localhost:4000/ at your browser


###### local ######
*LINUX*```GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build```
*OSX*  ```GOOS=darwin GOARCH=amd64  CGO_ENABLED=0 go build```
*START-APP* ```./gwyn```

### TEST ###
```go test -v ./api/v1/shorten``` 