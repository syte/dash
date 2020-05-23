package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func GetPwd() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func main() {
	mux := http.NewServeMux()
	pwd := GetPwd()

	var port int
	var certFile string
	var keyFile string

	flag.IntVar(&port, "port", 8080, "Specified Port")
	flag.StringVar(&certFile, "cert", "", "File Path to Certificate")
	flag.StringVar(&keyFile, "key", "", "Path to Key File")

	flag.Parse()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		path := path.Join(pwd, r.URL.Path)
		contents, err := ioutil.ReadFile(path)

		fmt.Printf("Now serving request for %s from %s", r.URL.Path, r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Write(contents)
	})

	host := fmt.Sprintf("localhost:%d", port)

	if certFile == "" || keyFile == "" {
		http.ListenAndServe(host, mux)
	} else {
		http.ListenAndServeTLS(host, certFile, keyFile, mux)
	}
}
