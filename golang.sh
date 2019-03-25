if [ ! -d "/go/src/github.com" ]; then
  echo 'go get'
  go get -u github.com/gin-gonic/gin
  go get github.com/lib/pq
  go get github.com/dgrijalva/jwt-go
  go get golang.org/x/crypto/bcrypt
  echo 'finish get'
fi
go get github.com/pilu/fresh
# go run /go/src/admin-api/router.go
# /go/bin/gin run /go/src/admin-api/router.go
cd /go/admin-api/ && fresh
