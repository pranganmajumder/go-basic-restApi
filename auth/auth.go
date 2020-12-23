package auth

import (
	"fmt"
	"net/http"
)


////Sabik Middleware

//func BasicAuthentication(next http.HandlerFunc) http.HandlerFunc {
//	return func(response http.ResponseWriter, request *http.Request) {
//		//user := os.Getenv("username")
//		//pass := os.Getenv("password")
//		user := "prangan"
//		pass := "1234"
//
//		username, password, authOK := request.BasicAuth()
//		fmt.Println("Userrrrr : ", user)
//		fmt.Println("Passsssss :  ", pass)
//
//		if authOK == false {
//			http.Error(response, "Not authorized", http.StatusUnauthorized)
//			return
//		}
//
//		if username != user || password != pass {
//			http.Error(response, "Not authorized", http.StatusUnauthorized)
//			return
//		}
//
//		next.ServeHTTP(response, request)
//	}
//}







func MiddlewareAuth(original http.HandlerFunc) func(http.ResponseWriter, *http.Request)  {
	return func(res http.ResponseWriter, req *http.Request) {
		user := "prangan"
		pass := "1234"
		username, password, authOk := req.BasicAuth()
		if authOk == false{
			http.Error(res, "Access Denied" , http.StatusUnauthorized)
			return
		}
		if username != user || password != pass{
			http.Error(res, "Access Denied" , http.StatusUnauthorized)
			return
		}
		fmt.Println("Middle Auth running")
		original.ServeHTTP(res, req)

	}
}