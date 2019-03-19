package database
import (
  // "database/sql"
  "fmt"
  // "log"
  // _ "github.com/lib/pq"
)

func ConnectDB() {
  fmt.Println("nice")
  // db, err := sql.Open("postgres",
	// 	"host=database user=postgres password=asdasd dbname=postgres sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
  //   rows, error := db.Query(`
	// 	SELECT
	// 		id,
	// 		name
	// 	FROM users`)
  //   if (error != nil) {
  //     fmt.Println("check")
  //     panic(error)
  //   } else {
  //     defer rows.Close()
  //     for rows.Next() {
  //       var id int
  //       var name string
  //       err = rows.Scan(&id, &name)
  //       fmt.Println(name)
  //       fmt.Println(id)
  //     }
  //   }
  // }
  //
	// defer db.Close()
}
