package main


import (

	// fmt allows us to print to our command line to debug our program.
	"fmt"

	//os allows us to access the operating systems current environment.
	"os"

	//godotenv allows us to load environment variables from a file into our current operating system's environment.
	"github.com/joho/godotenv"

	/*
		http gives us access to all networking libaries associated with http.
		Theoretically this is the only library we need to create a backend service, but gorilla/mux makes it a LOT easier.
	*/
	"net/http"

	//gorilla/mux builds on top of net/http to make it easier to create endpoints.
	"github.com/gorilla/mux"

	//json lets us take the json bodies from incoming requests and then decode them into a Golang object.
	"encoding/json"

	//strconv lets us convert any primitive data type (integer, string, boolean) into its string equivalent.
	//This allows us to convert our local data types to string when we try and output them in our HTTP responses
	"strconv"
)

/*
	We want to create a "blueprint" of a rectangle so that any function that wants to create a rectangle knows how.
	Every rectangle blueprint has its length and width. Both are type integer.
	the "json" tags at the end lets the "encoding/json" library know that any json keys
*/
type Rectangle struct {
	Length int `json:"length"`
	Width int `json:"Width"`
}

func helloWorld(response http.ResponseWriter, request *http.Request) {
	
	//Obtain the query paramters with keys "username" and "password"
	username := request.URL.Query().Get("username")
	password := request.URL.Query().Get("password")

	
	//Print the values of the query parameter to the response ouput
	//If we sent out request from insomnia, these are the values we will see
	fmt.Fprintf(response, username +"\n")
	fmt.Fprintf(response, password +"\n")


	//Take the JSON request body in the request we received and populate our empty Rectangle with the values

	//create an empty rectangle
	rectangle := Rectangle{}

	//create a new json decoder that will allow us to decode the request body
	jsonDecoder := json.NewDecoder(request.Body)

	//use our decoder to decode the contents of the request body into our rectangle
	err := jsonDecoder.Decode(&rectangle)

	//Check if we got an error. If err != nil, that function returned an error
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	//print out our rectangle dimensions to the response output
	//If we sent out request from insomnia, these are the values we will see
	fmt.Fprintf(response, strconv.Itoa(rectangle.Length) + "\n")
	fmt.Fprintf(response, strconv.Itoa(rectangle.Width) + "\n")


	//Obtain the the cookie "access_token", this gives us all the chararcteristics of the cookie
	cookie, err := request.Cookie("access_token")

	//check if obtaining the cookie returned an error
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	
	//Get the value of the cookie we just obtained and print it out to the response output
	accessToken := cookie.Value
	fmt.Fprintln(response, accessToken)

	return
}


func main() {

	//get command line arguments
	args := os.Args

	fmt.Println(args)

	//load the environment variables from the .env file into our program's environment
	err :=  godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	//Get the "SECRET_TOKEN" variable and store it in a local Golang variable
	secretToken := os.Getenv("SECRET_TOKEN")
	fmt.Println(secretToken)

	//Create a router that matches urls to their specific functions
	router := mux.NewRouter()


	//Forward every request with url "localhost:8080/helloworld" to the helloWorld function. The request method must be GET.
	router.HandleFunc("/helloworld", helloWorld).Methods(http.MethodGet)

	//Begin listening for requests
	http.ListenAndServe(":8080", router)



}