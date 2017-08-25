package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/maikonformigari/golang-azure-test/contato"

	log "github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang-test by Sogitec in Bitbuckt - porta %s", os.Getenv("HTTP_PLATFORM_PORT"))
}
func main() {
	log.Println("testando...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/contato", contato.ContatoHandler)
	http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
