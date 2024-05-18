// main.go
package main

import (
	superhero "back-end/package"
	"net/http"
)

func main() {
	http.HandleFunc("/api/superhero", superhero.GetSuperhero)
	http.ListenAndServe(":8080", nil)
}
