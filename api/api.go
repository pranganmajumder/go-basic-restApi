//https://www.youtube.com/watch?v=W5b64DXeP0o
//https://tutorialedge.net/golang/creating-restful-api-with-golang/

package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pranganmajumder/go-basic-restApi/auth"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Varsity    string `json:"varsity"`
	Occupation string `json:"occupation"`
}

var Users []User

func parseID(request *http.Request) string {
	params := mux.Vars(request)
	ID := params["id"]
	if len(ID) > 0 {
		return ID
	}

	values := request.URL.Query()
	if val, ok := values["id"]; ok && len(val) > 0 {
		return val[0]
	}
	return ""
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome to homepage")
	fmt.Println("Endpoint Hit: homepage")

}

func ReturnAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit:  returnAllUser")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Users); err != nil {
		log.Println(err)
	}
}

func ReturnSingleUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("This is our ReturnSingleUser function")
	//vars := mux.Vars(r)
	//key := vars["id"]

	key := parseID(r)
	// Loop over all of our Users
	// if the user.Id equals the key we pass in return the article encoded as JSON
	for _, user := range Users {
		if user.Id == key {
			w.WriteHeader(http.StatusOK)
			if err:= json.NewEncoder(w).Encode(user) ; err != nil{
					fmt.Println("Error occurred : " , err)
			}
			return
		}
	}
	// if the user Id is not found in the list return a bad request
	// https://stackoverflow.com/questions/40096750/how-to-set-http-status-code-on-http-responsewriter
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("400 - User not found in the list"))
}

//func BasicAuthentication(header string)bool  {
//	b := strings.Split(header , " ")
//	//fmt.Println("B   =  ", b[1])
//	decoded , _ := base64.StdEncoding.DecodeString(b[1])
//	//fmt.Println("encoded  :  " , string(encoded))
//	name := strings.Split(string(decoded), ":")
//	if name[0] == "prangan" && name[1] == "1234"{
//		return true
//	}
//	return false
//}

// Endpoint : /user/id
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("created  NewUser")
	// fmt.Fprintf(w, "%v", string(reqBody)) // print request body in responseBody

	vars := mux.Vars(r)
	id := vars["id"]
	//key := parseID(r)
	// check if ID is in the list or not
	for _, u := range Users {
		if u.Id == id {
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
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user) // send the json encoded format
}

// (PUT) Endpoint : /user/id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, u := range Users {
		if u.Id == id {
			w.WriteHeader(http.StatusOK)
			Users[index] = user
			err := json.NewEncoder(w).Encode(user) // you'll see the json user data in the response body
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}

	// if ID is not found return an error
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - Requested User is not present in the list , Update won't work"))

}

// (DELETE) Endpoint : /user/id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the user we wish to delete
	id := vars["id"]

	for index, user := range Users {
		if user.Id == id {
			Users = append(Users[:index], Users[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user) // you'll see the deleted json data in the response body
			return
		}
	}
	// if user Id is not present in the list , return a error code
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - User not found in the list"))
}

func HandleRequests(port string) {
	DBInit()
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage).Methods("GET")

	myRouter.HandleFunc("/users", auth.MiddlewareAuth(ReturnAllUser)).Methods("GET")

	// Ordering is important here! This has to be defined before the other '/user' endpoint
	myRouter.HandleFunc("/user/{id}", auth.MiddlewareAuth(CreateNewUser)).Methods("POST")
	myRouter.HandleFunc("/user/{id}", auth.MiddlewareAuth(UpdateUser)).Methods("PUT")
	myRouter.HandleFunc("/user/{id}", auth.MiddlewareAuth(DeleteUser)).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", auth.MiddlewareAuth(ReturnSingleUser)).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func DBInit() {
	Users = []User{
		{Id: "1", Name: "Prangan", Varsity: "CoU", Occupation: "Student"},
		{Id: "2", Name: "Sakib", Varsity: "NSU", Occupation: "software Engineer"},
	}
}

//func main() {
//	DBInit()
//	HandleRequests()
//}
