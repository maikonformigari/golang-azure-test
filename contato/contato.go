package contato

import (
	"fmt"
	"net/http"
)

func ContatoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fale conosco")
}
