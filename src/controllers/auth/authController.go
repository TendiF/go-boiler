package authController

import (
  "net/http"
  // jwt "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  // "fmt"
  "../../database"
  "../../utils"
  _ "github.com/lib/pq"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
	Id int
	Username string `form:"username"`
	Email string `form:"email"`
	EmailUsername string `form:"email_username"`
	Name string `form:"name"`
	Password string `form:"password"`
	IdRole int `form:"id_role"`
}

func RegisterUser(c *gin.Context) {
  // Bind variable
  var user User
	err := c.Bind(&user)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
			"status":  http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
			"error": nil,
		})
	}

  // ==== validate form ====
  errorMessage := map[string]string{}
  if (user.Username == "" || len(user.Username) < 6){
    errorMessage["username"] = "Harap isi username, minimal 6 karakter"
  }
  if (user.Password == "" || len(user.Password) < 8){
    errorMessage["password"] = "Harap isi password, minimal 8 karakter"
  }
  if (user.Email == "" || !utils.CheckEmail(user.Email)){
    errorMessage["email"] = "Harap isi email dengan benar"
  }
  if (user.Name == "" || len(user.Name) < 3){
    errorMessage["name"] = "Harap isi nama, minimal 3 karakter"
  }
  if (user.IdRole == 0){
    errorMessage["id_role"] = "Harap isi role user"
  }
  if (len(errorMessage) > 0) {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "status":  http.StatusBadRequest,
      "message":  http.StatusText(http.StatusBadRequest),
      "error":  errorMessage,
    })
    return
  }
  // ==== end validate form ====

  // pre INSERT
  bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
  hashPassword := string(bytes)

  // database query
  db := database.InitDB()
  const id_role = 4
  const query = `
    INSERT INTO public.users (username, email, password, name, id_role)
    VALUES ($1, $2, $3, $4, $5)`
  _, error := db.Exec(query, user.Username, user.Email, hashPassword, user.Name, id_role)
  if (error != nil){
    panic(error)
    return
  } else {
    user.Password = ""
    c.JSON(http.StatusOK, gin.H{
      "success": true,
      "status": http.StatusOK,
      "message": "Registrasi berhasil",
      "error":  nil,
      "data": user,
    })
    return
  }
}

func LoginUser(c *gin.Context) {
  // Bind variable
  var user User
	err := c.Bind(&user)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
			"status":  http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
			"error": nil,
		})
	}

  // ==== validate form ====
  errorMessage := map[string]string{}
  if (user.EmailUsername == "" || (!utils.CheckEmail(user.EmailUsername) && len(user.EmailUsername) < 6)){
    errorMessage["email_username"] = "Harap isi email / username dengan benar"
  }
  if (user.Password == "" || len(user.Password) < 8){
    errorMessage["password"] = "Harap isi password, minimal 8 karakter"
  }
  if (len(errorMessage) > 0) {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "status":  http.StatusBadRequest,
      "message":  http.StatusText(http.StatusBadRequest),
      "error":  errorMessage,
    })
    return
  }
  // ==== end validate form ====

  // database query
  db := database.InitDB()
  const id_role = 4
  query := `
    SELECT id, email, password, username FROM public.users WHERE email = '` + user.EmailUsername + `' OR username = '` + user.EmailUsername +`'  LIMIT 1`
  rows, error := db.Query(query)
  if (error != nil){
    c.JSON(http.StatusOK, gin.H{
      "success": false,
      "status": http.StatusOK,
      "message": "Email / username tidak ditemukan",
      "error":  nil,
      "data": nil,
    })
    return
  } else {
    for rows.Next() {
  		row := User{}
      rows.Scan(
  			&row.Id,
  			&row.Email,
  			&row.Password,
  			&row.Username,
  		)
      err := bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(user.Password))
      // wrong password
      if err != nil {
        c.JSON(http.StatusOK, gin.H{
          "success": false,
          "status": http.StatusOK,
          "message": "Kombinasi email / username dengan password salah",
          "error":  nil,
          "data": nil,
        })
  			return
  		}
      // success
  		c.JSON(http.StatusOK, gin.H{
        "success": true,
        "status": http.StatusOK,
        "message": "Berhasil masuk",
        "error":  nil,
        "data": row,
      })
  	}
  }
}
