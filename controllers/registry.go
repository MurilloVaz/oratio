package controllers

import (
	"io/ioutil"
	"net/http"

	e "github.com/MurilloVaz/oratio/entities"
	dto "github.com/MurilloVaz/oratio/entities/dto"
	g "github.com/MurilloVaz/oratio/extensions/guid"
	json "github.com/MurilloVaz/oratio/extensions/json"
	s "github.com/MurilloVaz/oratio/services"
	si "github.com/MurilloVaz/oratio/services/interfaces"
	"gopkg.in/mgo.v2/bson"
)

type RegistryController struct {
	urlService si.IUrlService
}

func (c *RegistryController) Post(w http.ResponseWriter, r *http.Request) {
	newGUID, err := g.UrlGuid()

	if err != nil {
		w.WriteHeader(501)
		return
	}

	var urlDto dto.UrlPost
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = json.FromJSON(&urlDto, body)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = c.urlService.InsertOne(&e.Url{Url: urlDto.Url, Id: bson.ObjectId(newGUID)})

	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Write([]byte(newGUID))
}

func (c *RegistryController) GetAll(w http.ResponseWriter, r *http.Request) {
	valueFromDb, err := c.urlService.GetAll()

	if err != nil {
		w.WriteHeader(400)
		return
	}

	jsonValue, err := json.ToJSON(valueFromDb)

	if err != nil {
		w.WriteHeader(501)
		return
	}

	w.Write(jsonValue)
}

func NewRegistryController(urlService *s.UrlService) *RegistryController {
	return &RegistryController{urlService}
}
