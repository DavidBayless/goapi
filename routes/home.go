package routes

import (
	"fmt"
	"goapi/db"
	"net/http"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("MOAR UNICORNZ!")
	// var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	var database = db.OpenConnection()
	users := db.GetAllUsers(database)
	writer.Header().Set("Content-Type", "application/json")
	// encoded, err := json.Marshal(users)
	writer.Write(users)
}
