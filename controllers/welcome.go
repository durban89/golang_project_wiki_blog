package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/durban89/wiki/config"

	"github.com/durban89/wiki/session"
	// memory session provider
	_ "github.com/durban89/wiki/session/providers/memory"
)

// Server Server
type Server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

// XMLServers XMLServers
type XMLServers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Servers     []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}

// SessionManager Session管理器
var SessionManager *session.Manager

// WelcomeProcessXML WelcomeProcessXML
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

// WelcomeXML WelcomeXML
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

func init() {
	log.Println("init welcome")
	var err error
	SessionManager, err = session.GetManager("memory", "sessionid", 3600)
	if err != nil {
		fmt.Println(err)
		return
	}

	go SessionManager.SessionGC()
}
