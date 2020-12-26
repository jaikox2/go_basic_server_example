package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server run on port: 8080")

	http.HandleFunc("/", indexHandler)
	http.Handle("/admin", allowRoles("admin")(http.HandlerFunc(adminHandler)))
	http.Handle("/staff", allowRoles("staff")(http.HandlerFunc(staffHandler)))
	http.Handle("/staff-admin", allowRoles("staff", "admin")(http.HandlerFunc(adminStaffHandler)))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(err)
	}

}

type middleware func(http.Handler) http.Handler

func allowRoles(roles ...string) middleware {
	allow := make(map[string]struct{})
	for _, role := range roles {
		allow[role] = struct{}{}
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("role")
			if _, ok := allow[reqRole]; !ok {
				http.Error(w, "Fobbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func allowRole(role string) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("role")
			if reqRole != role {
				http.Error(w, "Fobbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func adminStaffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin And Staff Page"))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin Page"))
}

func staffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Staff Page"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Index Page"))
}
