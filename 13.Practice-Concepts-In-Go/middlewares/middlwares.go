package middlewares

import (
	"fmt"
	"net/http"
)


func Middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Executing middleware1()........")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware1() again......")
	})
}

func Middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Executing middleware2()........")
		if r.URL.Path != "/" {
			return 
		}
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware1() again......")
	})
}

func Final(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Executing final()........")
	fmt.Fprintln(w, "Done......")
}
