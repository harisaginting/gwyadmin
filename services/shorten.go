package services

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"

	model "github.com/harisaginting/gwyn/models"
	httpModel "github.com/harisaginting/gwyn/models/http"
	repo "github.com/harisaginting/gwyn/repositories"
	"github.com/harisaginting/gwyn/utils/helper"
	"github.com/harisaginting/gwyn/utils/jwt/generator"
	"github.com/harisaginting/gwyn/utils/log"
)

type ShortenService interface {
	List(ctx context.Context, res *httpModel.ResponseList) (err error)
	Create(ctx context.Context, req httpModel.RequestCreate) (res httpModel.ResponseCreate, status int, err error)
	Status(ctx context.Context, code string) (res model.Shorten, status int, err error)
	Execute(ctx context.Context, code string) (res model.Shorten, status int, err error)
}

type Shorten struct {
	repo repo.ShortenRepository
}

func (service *Shorten) List(ctx context.Context, res *httpModel.ResponseList) (err error) {
	shortens, err := service.repo.FindAll(ctx)
	if err != nil {
		log.Error(ctx, err)
		return
	}
	res.Items = shortens
	res.Total = len(shortens)
	return
}

func (service *Shorten) Create(ctx context.Context, req httpModel.RequestCreate) (res httpModel.ResponseCreate, status int, err error) {
	status = http.StatusInternalServerError
	req.URL = helper.AdjustUrl(req.URL)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	checkUrl, err := http.NewRequest("GET", req.URL, nil)
	if err != nil {
		log.Error(ctx, err, "Failed initiate request to storageService Service")
		return
	}

	resCheckUrl, err := client.Do(checkUrl)
	if err != nil {
		status = http.StatusBadRequest
		err = errors.New("invalid url host")
		log.Error(ctx, err)
		return
	}
	if !(resCheckUrl.StatusCode >= 200 && resCheckUrl.StatusCode <= 300) {
		status = http.StatusBadRequest
		err = errors.New("url host not found")
		return
	}

	if req.Shortcode == "" {
		for {
			req.Shortcode = generator.GenerateIdentifier()
			check := model.Shorten{Shortcode: req.Shortcode}
			service.repo.Get(ctx, &check)
			if check.ID == 0 {
				break
			}
		}
	} else {
		if !helper.IsMatchRegex(req.Shortcode) {
			err = errors.New("The shortcode fails to meet the following regexp: ^[0-9a-zA-Z_]{6}$.")
			log.Error(ctx, err)
			status = http.StatusUnprocessableEntity
			return
		} else {
			check := model.Shorten{Shortcode: req.Shortcode}
			service.repo.Get(ctx, &check)
			if check.ID != 0 {
				err = errors.New("The desired shortcode is already in use. ")
				status = http.StatusConflict
				return
			}
		}
	}
	shorten, err := service.repo.Create(ctx, req)
	if err != nil {
		log.Error(ctx, err)
		status = http.StatusInternalServerError
		return
	}
	res.Shortcode = shorten.Shortcode
	status = http.StatusCreated
	return
}

func (service *Shorten) Status(ctx context.Context, code string) (res model.Shorten, status int, err error) {
	status = http.StatusInternalServerError
	res.Shortcode = code
	err = service.repo.Get(ctx, &res)
	if err != nil {
		log.Error(ctx, err)
		status = http.StatusInternalServerError
		return
	}
	if res.ID == 0 {
		status = http.StatusNotFound
		err = errors.New("The shortcode cannot be found in the system")
		log.Error(ctx, err)
		return
	}
	status = http.StatusOK
	return
}

func (service *Shorten) Execute(ctx context.Context, code string) (res model.Shorten, status int, err error) {
	status = http.StatusInternalServerError
	res.Shortcode = code
	err = service.repo.Get(ctx, &res)
	if err != nil {
		status = http.StatusInternalServerError
		log.Error(ctx, err)
		return
	}

	if res.ID == 0 {
		status = http.StatusNotFound
		err = errors.New("The shortcode cannot be found in the system")
		log.Error(ctx, err)
		return
	}
	service.repo.Execute(ctx, res)
	status = http.StatusFound
	return
}
