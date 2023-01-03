package middle

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv"
)

type response struct {
	ID      int64  `json:"id.omitempty"`
	Message string `json:"message.omitempty"`
}

func CreateConnection() *sql.DB {
	err := gotdotenv.Load(".env")
	if err != nil {
		log.Fatal("err loading .env file")
	}
	sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)

	}
	fmt.Println("successfully connected to postgres")
	return db
}

func CreateStock() {

}
func GetStock() {

}
func GetAllStock() {

}
func UpdateStock() {

}
func DeleteStock() {

}
