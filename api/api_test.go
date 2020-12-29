package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnSingleUser(t *testing.T) {
	DBInit()

	tests := []struct {
		Method             string
		Id                 string
		ExpectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			Method:             "GET",
			Id:                 "1",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method:             "GET",
			Id:                 "2",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method:             "GET",
			Id:                 "3",
			ExpectedStatusCode: http.StatusNoContent,
		},
	}
	for index, tt := range tests {
		req, err := http.NewRequest(tt.Method, "http://localhost:8080/user/userId?id="+tt.Id, nil)
		if err != nil {
			t.Fatalf("test request failed for testCase  : %v ", index+1)
		}
		res := httptest.NewRecorder()
		ReturnSingleUser(res, req)
		if r := res.Result(); r.StatusCode != tt.ExpectedStatusCode {
			fmt.Println("r.StatusCode  :  ", r.StatusCode)
			fmt.Printf("Failed to TestReturnSingleUser for testcase %v ", index+1)
		} //else{
		//fmt.Printf("Passed TestReturnSingleUser for testcase %v\n ", index+1)
		//}
	}
}

func TestReturnAllUser(t *testing.T) {
	tests := []struct {
		Method             string
		ExpectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			Method:             "GET",
			ExpectedStatusCode: http.StatusOK,
		},
	}
	for index, tt := range tests {
		req, err := http.NewRequest(tt.Method, "http://localhost:8080/users", nil)
		if err != nil {
			t.Fatalf("test request failed for testCase : %v ", index+1)
		}
		res := httptest.NewRecorder()
		ReturnAllUser(res, req)
		if r := res.Result(); r.StatusCode != tt.ExpectedStatusCode {
			fmt.Println("r.StatusCode :  ", r.StatusCode)
			fmt.Printf("Failed to TestReturnAllUser for test case : %v ", index+1)
		}
	}
}

func TestCreateNewUser(t *testing.T) {
	tests := []struct {
		Method             string
		Person             User
		ExpectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			Method: "POST",
			Person: User{
				Id:         "3",
				Name:       "Dipika Padukone",
				Varsity:    "North South University",
				Occupation: "Modelling",
			},
			ExpectedStatusCode: http.StatusOK,
		},
	}
	for index, tt := range tests {
		byteData, _ := json.Marshal(tt.Person)

		req, err := http.NewRequest(tt.Method, "http://localhost:8080/api/user/userId?id="+tt.Person.Id, bytes.NewReader(byteData))
		if err != nil {
			t.Fatalf("unable to create any request : %v", err)
		}

		res := httptest.NewRecorder()
		CreateNewUser(res, req)
		if r := res.Result(); r.StatusCode != tt.ExpectedStatusCode {
			fmt.Printf("Test Failed to Create a New User for test case : %v ", index+1)
		}
	}
}

func TestUpdateUser(t *testing.T) {
	DBInit()
	tests := []struct {
		Method             string
		Person             User
		Url                string
		ExpectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			Method: "PUT",
			Url:    "/user/%s",
			Person: User{
				Id:         "1",
				Name:       "Amitav Baccan",
				Varsity:    "Honululu University",
				Occupation: "Actor",
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method: "PUT",
			Url:    "/user/%s",
			Person: User{
				Id:         "4",
				Name:       "Emruz Hosasin",
				Varsity:    "CUET",
				Occupation: "Software Engineer",
			},
			ExpectedStatusCode: http.StatusBadRequest,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", UpdateUser).Methods(http.MethodPut)

	for _, tt := range tests {
		byteData, _ := json.Marshal(tt.Person)
		url := fmt.Sprintf(tt.Url, tt.Person.Id)

		req, err := http.NewRequest(tt.Method, url, bytes.NewReader(byteData))
		if err != nil {
			t.Fatalf("unable to create any request  : %v", err)
		}

		vars := make(map[string]string)
		vars["id"] = tt.Person.Id
		req = mux.SetURLVars(req, vars)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, tt.ExpectedStatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		Method             string
		Id                 string
		Url                string
		ExpectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			Method:             "DELETE",
			Id:                 "1",
			Url:                "/user/%s",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method:             "DELETE",
			Id:                 "20",
			Url:                "/user/%s",
			ExpectedStatusCode: http.StatusBadRequest,
		},
	}
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", DeleteUser).Methods(http.MethodDelete)

	for _, tt := range tests {
		byteData, _ := json.Marshal(tt.Id)
		url := fmt.Sprintf(tt.Url, tt.Id)

		req, err := http.NewRequest(tt.Method, url, bytes.NewReader(byteData))
		if err != nil {
			t.Fatalf("unable to create any request  : %v", err)
		}

		vars := make(map[string]string)
		vars["id"] = tt.Id
		req = mux.SetURLVars(req, vars)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, res.Result().StatusCode, tt.ExpectedStatusCode)
	}
}
