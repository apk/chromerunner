package main

import (
	"flag"
	"os/exec"
	"io/ioutil"
	"net/http"
	"log"
)

var (
	addr       = flag.String("addr", ":3045", "http service address")
	path       = flag.String("path", "chrome", "browser executable")
)

func runchrome(url string) {
		log.Print("URL: ", url)
	cmd := exec.Command(*path, url)
	out, err := cmd.Output()
	if err != nil {
		log.Print("Exec:", err)
	} else {
		s := string(out);
		if s != "" {
			log.Print("Out: ", s);
		}
	}
}

func main() {
	flag.Parse()

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			go runchrome (string(body))
		}
	})

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
