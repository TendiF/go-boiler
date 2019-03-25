package database
import (
  "database/sql"
  "fmt"
  "log"
  _ "github.com/lib/pq"
)

const (
    dbhost = "DBHOST"
    dbport = "DBPORT"
    dbuser = "DBUSER"
    dbpass = "DBPASS"
    dbname = "DBNAME"
)

func InitDB() *sql.DB{
  config := dbConfig()
  var err error
  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
      "password=%s dbname=%s sslmode=disable",
      config[dbhost], config[dbport],
      config[dbuser], config[dbpass], config[dbname])
  var db *sql.DB
  db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
      panic(err)
  }
  err = db.Ping()
  if err != nil {
      panic(err)
  }
  fmt.Println("Successfully connected!")
  return db
}

func dbConfig() map[string]string {
    conf := make(map[string]string)
    conf[dbhost] = "database"
    conf[dbport] = "5432"
    conf[dbuser] = "postgres"
    conf[dbpass] = "asdasd"
    conf[dbname] = "platinum"
    return conf
}

func ConnectDB() {
  db, err := sql.Open("postgres",
		"host=database user=postgres password=asdasd dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	} else {
    rows, error := db.Query(`
		SELECT
			id,
			name
		FROM users`)
    if (error != nil) {
      fmt.Println("check")
      panic(error)
    } else {
      defer rows.Close()
      for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        fmt.Println(name)
        fmt.Println(id)
      }
    }
  }

	defer db.Close()
}
