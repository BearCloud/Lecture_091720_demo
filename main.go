package main


import (
	"fmt"
	_"os"
	_"github.com/joho/godotenv"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
)

type Rectangle struct {
	Length int `json:"length"`
	Width int `json:"Width"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	
	//query parameters
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	fmt.Fprintf(w, username +"\n")
	fmt.Fprintf(w, password +"\n")


	//json body
	rectangle := Rectangle{}
	err := json.NewDecoder(r.Body).Decode(&rectangle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, strconv.Itoa(rectangle.Length) + "\n")
	fmt.Fprintf(w, strconv.Itoa(rectangle.Width) + "\n")


	//cookie
	cookie, err := r.Cookie("access_token")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	accessToken := cookie.Value
	fmt.Fprintln(w, accessToken)

	return
}


func main() {

	// //get command line arguments
	// args := os.Args

	// fmt.Println(args)

	// //environment variables
	// err :=  godotenv.Load()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// secretToken := os.Getenv("SECRET_TOKEN")
	// fmt.Println(secretToken)

	//server programming
	router := mux.NewRouter()

	router.HandleFunc("/helloworld", helloWorld).Methods(http.MethodGet, http.MethodPost)

	http.ListenAndServe(":8080", router)



}