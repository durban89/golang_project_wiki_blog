package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type JsonServer struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type JsonServers struct {
	Servers []JsonServer `json:"servers"`
}

type TestJsonServer struct {
	// ID will not be outputed.
	ID int `json:"-"`

	// ServerName2 will be converted to JSON type.
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// If ServerIP is empty, it will not be outputted.
	ServerIP string `json:"serverIP,omitempty"`
}

func JsonToTest(w http.ResponseWriter, r *http.Request) {
	s := TestJsonServer{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}

	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
}

func Json(w http.ResponseWriter, r *http.Request) {
	var s JsonServers
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func JsonToInterface(w http.ResponseWriter, r *http.Request) {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("Unmarshal error")
		return
	}

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func JsonProcess(w http.ResponseWriter, r *http.Request) {
	var s JsonServers
	s.Servers = append(s.Servers, JsonServer{
		ServerName: "Shanghai_VPN",
		ServerIP:   "127.0.0.1",
	})

	s.Servers = append(s.Servers, JsonServer{
		ServerName: "Beijing_VPN",
		ServerIP:   "127.0.0.2",
	})

	b, err := json.Marshal(s)

	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	fmt.Println(string(b))
}
