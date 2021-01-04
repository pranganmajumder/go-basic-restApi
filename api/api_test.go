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
		ID                 string
		ExpectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			Method:             "GET",
			ID:                 "1",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method:             "GET",
			ID:                 "2",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method:             "GET",
			ID:                 "3",
			ExpectedStatusCode: http.StatusNoContent,
		},
	}
	for index, tt := range tests {
		req, err := http.NewRequest(tt.Method, "http://localhost:8080/user/userId?id="+tt.ID, nil)
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
				ID:         "3",
				Name:       "Dipika Padukone",
				Varsity:    "North South University",
				Occupation: "Modelling",
			},
			ExpectedStatusCode: http.StatusOK,
		},
	}
	for index, tt := range tests {
		byteData, _ := json.Marshal(tt.Person)

		req, err := http.NewRequest(tt.Method, "http://localhost:8080/api/user/userId?id="+tt.Person.ID, bytes.NewReader(byteData))
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
		URL                string
		ExpectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			Method: "PUT",
			URL:    "/user/%s",
			Person: User{
				ID:         "1",
				Name:       "Amitav Baccan",
				Varsity:    "Honululu University",
				Occupation: "Actor",
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method: "PUT",
			URL:    "/user/%s",
			Person: User{
				ID:         "4",
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
		url := fmt.Sprintf(tt.URL, tt.Person.ID)

		req, err := http.NewRequest(tt.Method, url, bytes.NewReader(byteData))
		if err != nil {
			t.Fatalf("unable to create any request  : %v", err)
		}

		vars := make(map[string]string)
		vars["id"] = tt.Person.ID
		req = mux.SetURLVars(req, vars)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, tt.ExpectedStatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		Method             string
		ID                 string
		URL                string
		ExpectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			Method:             "DELETE",
			ID:                 "1",
			URL:                "/user/%s",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method:             "DELETE",
			ID:                 "20",
			URL:                "/user/%s",
			ExpectedStatusCode: http.StatusBadRequest,
		},
	}
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", DeleteUser).Methods(http.MethodDelete)

	for _, tt := range tests {
		byteData, _ := json.Marshal(tt.ID)
		url := fmt.Sprintf(tt.URL, tt.ID)

		req, err := http.NewRequest(tt.Method, url, bytes.NewReader(byteData))
		if err != nil {
			t.Fatalf("unable to create any request  : %v", err)
		}

		vars := make(map[string]string)
		vars["id"] = tt.ID
		req = mux.SetURLVars(req, vars)

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		assert.Equal(t, res.Result().StatusCode, tt.ExpectedStatusCode)
	}
}
