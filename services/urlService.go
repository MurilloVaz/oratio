package services

import (
	e "github.com/MurilloVaz/oratio/entities"
	"github.com/MurilloVaz/oratio/extensions/logger"
	r "github.com/MurilloVaz/oratio/repositories"
	ri "github.com/MurilloVaz/oratio/repositories/interfaces"
)

type UrlService struct {
	urlRepo ri.IUrlRepository
}

func (s *UrlService) InsertOne(url *e.Url) error {
	defer s.urlRepo.Disconnect()

	err := s.urlRepo.InsertOne(url)

	if err != nil {
		logger.Database(err)
		return err
	}

	return nil
}

func (s *UrlService) FindOne(id *string) (*e.Url, error) {
	defer s.urlRepo.Disconnect()

	url, err := s.urlRepo.FindOne(id)

	if err != nil {
		logger.Unexpected(err)
		return nil, err
	}

	return url, nil
}

func (s *UrlService) GetAll() ([]*e.Url, error) {
	defer s.urlRepo.Disconnect()

	url, err := s.urlRepo.GetAll()

	if err != nil {
		logger.Database(err)
		return nil, err
	}

	return url, nil
}

func NewUrlService(repo *r.UrlRepository) *UrlService {
	return &UrlService{repo}
}
