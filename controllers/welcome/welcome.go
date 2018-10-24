package welcome

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/durban89/wiki/config"

	"github.com/durban89/wiki/session"
	// memory session provider
	_ "github.com/durban89/wiki/session/providers/memory"
)

type Server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type XMLServers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Servers     []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}

var appSession *session.Manager

func WelcomeProcessXML(w http.ResponseWriter, r *http.Request) {
	v := &XMLServers{
		Version: "1",
	}

	v.Servers = append(v.Servers, Server{
		ServerName: "Shanghai_VPN",
		ServerIP:   "127.0.0.1",
	})

	v.Servers = append(v.Servers, Server{
		ServerName: "Beijing_VPN",
		ServerIP:   "127.0.0.2",
	})

	output, err := xml.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)

}

func WelcomeXML(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(config.TemplateDir + "/server.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	v := XMLServers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(w, v)
}

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
