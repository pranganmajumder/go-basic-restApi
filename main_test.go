package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnSingleUser(t *testing.T) {
	DBInit()
	tests := []struct {
		Method            string
		Id                string
		ExpectedStatusCode int
	}{
		// TODO: Add test cases.
		{Method: "GET", Id: "1" , ExpectedStatusCode: http.StatusOK},
		{Method: "GET", Id: "2" , ExpectedStatusCode: http.StatusOK},
		{Method: "GET", Id: "3" , ExpectedStatusCode: http.StatusNoContent},
		{Method: "GET", Id: "4" , ExpectedStatusCode: http.StatusNoContent},
		{Method: "GET", Id: "5" , ExpectedStatusCode: http.StatusNoContent},
	}
	for index, tt := range tests {
		request , err := http.NewRequest(tt.Method , "http://localhost:8080/user/userId?id="+tt.Id, nil)
		fmt.Println("request:   ", request)
		if err != nil{
			t.Fatalf("unabale to create any request : %v", err)
		}
		response := httptest.NewRecorder()
		ReturnSingleUser(response , request)
		if res := response.Result(); res.StatusCode != tt.ExpectedStatusCode{
			t.Errorf("Case %v: expected %v got %v", index, tt.ExpectedStatusCode , res.Status)
		}else{
			fmt.Println("Matched with ")
		}
	}
}