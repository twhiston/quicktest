package main

import (
	"encoding/json"
	"github.com/owtorg/wrouter"
	"log"
	"net/http"
	"os"
)

type mainController struct {
}

func (m *mainController) Action(h *http.Request, w http.ResponseWriter) {
	profile := make(map[string]string, 2)
	profile["name"] = "test"
	profile["version"] = "1"

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	log.Println("Jenkins Test is Running")
	router := wrouter.NewRouter()
	router.AddController(new(mainController))
	router.PrintRoutes(os.Stdout)
	http.ListenAndServe(":8020", router)
}

func PrintSomething(something string) {
	log.Println("Printing:", something)
}
