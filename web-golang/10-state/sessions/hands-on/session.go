package main

import (
	"net/http"
)

func getUser(r *http.Request) user{
	var u user

	c, err := r.Cookie("sid")
	if err == nil {
		if s, ok := sessions[c.Value]; ok{
			u = users[s]
		}
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool{
	isLoggedIn := false
	if c, err := r.Cookie("sid"); err == nil {
		if s, ok := sessions[c.Value]; ok {
			_, ok := users[s]
			isLoggedIn = ok
		}
	}
	return isLoggedIn
}

//We can also store the last login time in the session and clean them after some time
//func cleanSessions() {
//	for k, v := range dbSessions {
//		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
//			delete(dbSessions, k)
//		}
//	}
//	dbSessionsCleaned = time.Now()
//}