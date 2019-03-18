if [ ! -d "/go/src/github.com" ]; then
  echo 'go get'
  go get -u github.com/gin-gonic/gin
fi
go run /go/src/app/router.go
