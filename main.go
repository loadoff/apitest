package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/rs/cors"
)

type ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/hello", getHello),
		rest.Get("/", getHome),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	handler := cors.Default().Handler(api.MakeHandler())
	log.Fatal(http.ListenAndServe("", handler))
}

func corsHeader(w rest.ResponseWriter, req *rest.Request) {
	ping := ping{http.StatusOK, "ok"}
	w.WriteJson(ping)
}

func getHello(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(map[string]string{
		"Hello": "World",
	})
}

func getHome(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(map[string]string{
		"Home": "ホーム",
	})
}
