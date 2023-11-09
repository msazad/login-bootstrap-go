package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/cookiejar"

	"github.com/google/uuid"
)

var tmpl *template.Template

var dbSessions = make(map[string]string)
var dbUsers = make(map[string]user)

type user struct {
	Name     string
	Username string
	Password string
}
type errors struct {
	UserError string
	PassError string
	FullError string
}

var errorval errors

func init() {
	template.Must(template.ParseGlob())
	dbUsers["azad@gmail.com"]=user{"azad","azad@gmail.com","1212"}
}
func main() {
	fmt.Println("Server is running in port :8080")
	http.HandleFunc("/", loginHandler)
}
//loginHandler function
func loginHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err == nil {
		if _, ok := dbSessions[cookie.Value]; ok {
			http.Redirect(w, req, "/home", http.StatusSeeOther)
		}
	}
	if req.Method == http.MethodPost {
		uname := req.FormValue("username")
		pass := req.FormValue("password")
		// check username
		if _, ok := dbUsers[uname]; !ok {
			http.Redirect(w, req, "/", http.StatusSeeOther)
			errorval.UserError = "Username Error"
			return
		}
		//check Password
		if pass != dbUsers[uname].Password {
			http.Redirect(w, req, "/", http.StatusSeeOther)
			errorval.PassError = "Password Error"
			return
		}
		//if Password matches
		if pass==dbUsers[uname].Password{
			//Create Cookie
			uid:=uuid.NewString()
			cookie=&http.Cookie{
				Name: "session",
				Value: uid,
			}
			http.SetCookie(w,cookie)
			dbSessions[cookie.Value]=uname
			http.Redirect(w,req,"/home",http.StatusSeeOther)
		}
	}
	tmpl.ExecuteTemplate(w,"login.html",errorval)
}

//signupHandler function