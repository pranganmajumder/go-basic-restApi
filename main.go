//https://www.youtube.com/watch?v=W5b64DXeP0o
//https://tutorialedge.net/golang/creating-restful-api-with-golang/


package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Varsity string `json:"varsity"`
	Occupation string `json:"occupation"`
}
var Users []User

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w , "Wellcome to homepage")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit:  returnAllArticles")
	json.NewEncoder(w).Encode(Users)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Users
	// if the user.Id equals the key we pass in return the article encoded as JSON
	var flag = false
	for _, user := range Users{
		if user.Id == key{
			json.NewEncoder(w).Encode(user)
			flag = true
		}
	}
	// if the user Id is not found in the list return a bad request
	// https://stackoverflow.com/questions/40096750/how-to-set-http-status-code-on-http-responsewriter
	if flag == false{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte ("400 - User Not Found in the list"))
	}
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody , _ := ioutil.ReadAll(r.Body)
	fmt.Println("createdNewUser" )
	// fmt.Fprintf(w, "%v", string(reqBody)) // print request body in responseBody

	vars := mux.Vars(r)
	id := vars["id"]
	// check if ID is in the list or not
	for _, u := range Users{
		if u.Id == id{
			fmt.Println("User ID matched")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - User ID already contains in the list. Please Change your ID"))
			return
		}
	}

	// now unmarshal this into a new User struct
	var user User
	json.Unmarshal(reqBody, &user)


	// append this to our Users array
	// update our global Users array to include our new User
	Users = append(Users, user)
	json.NewEncoder(w).Encode(user) // send the json encoded format
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the user we wish to delete
	id := vars["id"]

	var flag = false
	for index, user := range Users{
		if user.Id == id{
			Users = append(Users[:index] , Users[index+1:]...)
			flag = true
		}
	}
	// if user Id is not present in the list , return a error code
	if flag == false{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - User not found in the list"))
	}

}



func handleRequests()  {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/users", returnAllUser).Methods("GET")

	// Ordering is important here! This has to be defined before the other '/user' endpoint
	myRouter.HandleFunc("/user/{id}", createNewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", returnSingleUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",myRouter))
}

func main() {
	Users = []User{
		User{Id:         "1", Name:       "Prangan", Varsity:    "CoU", Occupation: "Student"},
		User{Id:         "2", Name:       "Sakib", Varsity:    "NSU", Occupation: "software Engineer"},
	}
	handleRequests()
}