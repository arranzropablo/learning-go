package main

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}
	//if u.Role != "admin" {
	//	http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
	//	return
	//}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	switch r.Method {

		case http.MethodPost:
			//retrieve data
			un := r.FormValue("username")
			p := r.FormValue("password")
			f := r.FormValue("firstname")
			l := r.FormValue("lastname")

			//validate data
			if _, ok := users[un]; ok {
				http.Error(w, "Username already taken", http.StatusPreconditionFailed)
				return
			}

			//generate session
			sid, _ := uuid.NewV4()
			c := &http.Cookie{
				Name: "sid",
				Value: sid.String(),

				//If we want a session to expire over time we can set maxage
				MaxAge: 10,
			}

			//We want to encrypt the password to store it
			//The higher the cost the harder is to encrypt it, and the harder to decrypt
			password, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			//save session and data
			http.SetCookie(w, c)
			sessions[sid.String()] = un
			users[un] = user{un, password, f, l}

			//redirect to homepage
			http.Redirect(w, r, "/", http.StatusSeeOther)

		case http.MethodGet:
			fallthrough
		default:
			tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	switch r.Method {

		case http.MethodPost:
			//retrieve data
			un := r.FormValue("username")
			p := r.FormValue("password")

			//validate data
			if _, ok := users[un]; !ok {
				http.Error(w, "Username does not exist", http.StatusPreconditionFailed)
				return
			}

			//compare passwords to log in
			if wrong := bcrypt.CompareHashAndPassword(users[un].Password, []byte(p)) != nil; wrong {
				http.Error(w, "Wrong password", http.StatusUnauthorized)
				return
			}

			//generate session
			sid, _ := uuid.NewV4()
			c := &http.Cookie{
				Name: "sid",
				Value: sid.String(),
				MaxAge: 10,
			}
			//save session and data
			http.SetCookie(w, c)
			sessions[sid.String()] = un

			//redirect to homepage
			http.Redirect(w, r, "/", http.StatusSeeOther)

		case http.MethodGet:
			fallthrough
		default:
			tpl.ExecuteTemplate(w, "login.gohtml", nil)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, err := r.Cookie("sid")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	//delete cookie from map      redirect a /
	delete(sessions, c.Value)

	//expire cookie
	c.MaxAge = -1
	http.SetCookie(w, c)

	//redirect to /
	http.Redirect(w, r, "/", http.StatusSeeOther)

}