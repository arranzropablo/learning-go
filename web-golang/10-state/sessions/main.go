package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//to "create a session" we set a uuid in the cookie using a uuid generator
	//we can link this uuids with users to preserve a session

	cookie, err := req.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//this is useful if using https
			// Secure: true,
			//this means you can't get this cookie through javascript
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
