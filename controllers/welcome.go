package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/durban89/wiki/config"

	"github.com/durban89/wiki/session"
	// memory session provider
	_ "github.com/durban89/wiki/session/providers/memory"
)

var appSession *session.Manager

// WelcomeLogin 欢迎登录页
func WelcomeLogin(w http.ResponseWriter, r *http.Request) {
	session, err := appSession.SessionStart(w, r)
	if err != nil {
		fmt.Fprintf(w, "session error")
		return
	}

	count := session.Get("count")
	if count == nil {
		session.Set("count", 1)
	} else {
		session.Set("count", count.(int)+1)
	}

	t, err := template.ParseFiles(config.TemplateDir + "/login.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, session.Get("count"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func init() {
	var err error
	appSession, err = session.GetManager("memory", "sessionid", 3600)
	if err != nil {
		fmt.Println(err)
		return
	}

	go appSession.SessionGC()
}
