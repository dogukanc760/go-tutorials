package main

import (
	"net/http"

	helper "./helpers"
)

func main() {

	mux := http.NewServeMux()

	uName, email, pwd, pwdConfirm := "", "", "", ""

	// Signup
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("passwordConfirm")

		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		if uNameCheck && pwd == pwdConfirm && emailCheck {
			// Commit the database
		}
		//  for key, value := r.Form {
		// 	fmt.Printf("%s = %s\n", key, value)
		//  }
	})

	// Login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(":8080", mux)

}
