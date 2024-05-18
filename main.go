// main.go
package main

import (
	"net/http"

	superhero "github.com/camiloarteaga1/Superheros_go/package"
)

func main() {
	http.HandleFunc("/api/superhero", superhero.GetSuperhero)
	http.ListenAndServe(":8080", nil)
}
