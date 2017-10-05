package routes

import (
	"fmt"
	"net/http"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("MOAR UNICORNZ!")
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonStr)
}
