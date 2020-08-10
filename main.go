package main

import (
	"fmt"
	"net/http"

	"github.com/MurilloVaz/oratio/controllers"
	"github.com/MurilloVaz/oratio/data"
	"github.com/MurilloVaz/oratio/extensions/configuration"
	"github.com/MurilloVaz/oratio/extensions/logger"
	"github.com/MurilloVaz/oratio/repositories"
	"github.com/MurilloVaz/oratio/services"
	"github.com/gorilla/mux"
)

func main() {
	configuration.SetConfiguration()
	logger.SetLogger()

	r := mux.NewRouter()
	r.HandleFunc("/api/registry", func(w http.ResponseWriter, r *http.Request) {
		controller := initializeRegistryController()
		controller.GetAll(w, r)
	}).Methods("GET")
	r.HandleFunc("/api/registry", func(w http.ResponseWriter, r *http.Request) {
		controller := initializeRegistryController()
		controller.Post(w, r)
	}).Methods("POST")
	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller := initializeRedirectController()
		controller.Redirect(w, r)
	}).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", configuration.Global.Port), r)
}

func initializeRegistryController() *controllers.RegistryController {
	ctx := data.NewDataContext()
	repo := repositories.NewUrlRepository(&ctx)
	service := services.NewUrlService(repo)
	controller := controllers.NewRegistryController(service)
	return controller
}

func initializeRedirectController() *controllers.RedirectController {
	ctx := data.NewDataContext()
	repo := repositories.NewUrlRepository(&ctx)
	service := services.NewUrlService(repo)
	controller := controllers.NewRedirectController(service)
	return controller
}
