package middle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres-yt/models"
	"log"
	"net/http"
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

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("unable to decode the request body")
	}
	insertID := inspectStock(stock)
	res := response{
		ID:      insertID,
		Message: "stock created successfully",
	}
	json.NewEncoder(w).Encode(res)
}
func GetStock(w http.ResponseWriter, r *http.Request) {

}
func GetAllStock(w http.ResponseWriter, r *http.Request) {

}
func UpdateStock(w http.ResponseWriter, r *http.Request) {

}
func DeleteStock(w http.ResponseWriter, r *http.Request) {

}
