package welcome

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/durban89/wiki/config"
	"github.com/durban89/wiki/views"

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

// ProcessXML 处理XML
func ProcessXML(w http.ResponseWriter, r *http.Request) {
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

// XML 文件解析
func XML(w http.ResponseWriter, r *http.Request) {
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

// Login 欢迎登录页
func Login(w http.ResponseWriter, r *http.Request) {
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

	views.Render(w, "login.html", session.Get("count"))

	return
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
