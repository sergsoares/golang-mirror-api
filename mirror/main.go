package main

import (
	"fmt"
	"log"
	"net"
 	"net/http"
	"encoding/json"
	"os"
)

func main() {
	log.Print("Initialized")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
        "RemoteAddr": r.RemoteAddr,
        "LocalIp": fmt.Sprintf("%s",GetOutboundIP()),
		"Method" : r.Method,
		"URL" : fmt.Sprintf("%s",r.URL),
		"Proto" : r.Proto,	
		"Host" : r.Host,
		"Version" : GetVersion(),
    }

	for k, v := range r.Header {
		response["Header[" + k + "]"] = fmt.Sprintf("%s",v)
	}

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		response[k] = fmt.Sprintf("%s",v)
	}
	
    jsonStr, err := json.MarshalIndent(response, "", "    ")
   
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    } else {
        fmt.Fprintf(w,string(jsonStr))
		log.Println(string(jsonStr))
    }
}

func GetVersion() string {
	const envKey = "DD_VERSION"

	val, ok := os.LookupEnv(envKey)

	if !ok {
		return "Without Version"
	} else {
		return val
	}
}

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}
