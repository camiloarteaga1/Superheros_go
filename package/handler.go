// package/handler.go
package superhero

import (
	"encoding/json"
	"net/http"
)

func GetSuperhero(w http.ResponseWriter, r *http.Request) {
	heroName := r.URL.Query().Get("hero")
	if hero, ok := Superheroes[heroName]; ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hero)
	} else {
		http.Error(w, "Superhero not found", http.StatusNotFound)
	}
}
