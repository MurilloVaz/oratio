package controllers

import (
	"net/http"

	s "github.com/MurilloVaz/oratio/services"
	si "github.com/MurilloVaz/oratio/services/interfaces"
	"github.com/gorilla/mux"
)

type RedirectController struct {
	urlService si.IUrlService
}

func (s *RedirectController) Redirect(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url, err := s.urlService.FindOne(&id)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	http.Redirect(w, r, url.Url, http.StatusSeeOther)
}

func NewRedirectController(urlService *s.UrlService) *RedirectController {
	return &RedirectController{urlService}
}
