package middle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres-yt/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
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
		log.Fatal("unable to decode the request body. %v", err)
	}
	insertID := inspectStock(stock)
	res := response{
		ID:      insertID,
		Message: "stock created successfully",
	}
	json.NewEncoder(w).Encode(res)
}
func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("unable to convert the string into int. %v", err)
	}
	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatal("unable to get stock . %v", err)
	}

}
func GetAllStock(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAllStocks()
	if err != nil {
		log.Fatal("unable to get all the stocks %v ", err)
	}
	json.NewEncoder(w).Encode(stocks)
}
func UpdateStock(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("unable to convert string into int. %v", err)
	}
	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("unable to decode the request body . %v", err)
	}
	updateRows := updateStock(int64(id), stock)
	msg := fmt.Sprintf("stock updated successfully . total rows / records affected %v", updateRows)
	res := response{
		ID:      int64(id),
		message: msg,
	}
	json.NewDecoder(w).Encode(res)
}
func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"])
	if err != nil {
		log.Fatalf("unable to convert string to int. %v", err)
	}
	deleteRows := deleteStock(int64(id))
	msg := fmt.Sprintf("stock deleted successfully.Total rows/records %v", deleteRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}
func insertStock(stock models.Stock) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO STOCKS (name,price,company) VALUES($1,$2,$3)RETURNING`
	var id int64

	db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&oid)
	if err != nil {
		log.Fatalf("unable to execute the query. %v", err)

	}
	fmt.Printf("inserted a single record %v", id)
	return id
}

func getStock(id int64) (models.Stock, error) {

}
func getAllStocks() ([]models.Stock, error) {

}
func updateStock(id int64, stock models.Stock) int64 {

}
func deleteStock(id int64) int64 {

}
