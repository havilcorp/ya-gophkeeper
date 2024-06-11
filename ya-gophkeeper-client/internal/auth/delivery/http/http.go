package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"ya-gophkeeper-client/internal/auth/entity"
	"ya-gophkeeper-client/internal/config"

	"github.com/sirupsen/logrus"
)

type handler struct {
	conf   *config.Config
	client *http.Client
}

func New(conf *config.Config, client *http.Client) *handler {
	return &handler{
		conf:   conf,
		client: client,
	}
}

func (h *handler) Login(dto *entity.LoginDto) (string, error) {
	data, err := json.Marshal(dto)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	body := bytes.NewReader(data)
	contentType := "application/json"
	resp, err := h.client.Post(fmt.Sprintf("%s/auth/login", h.conf.AddressHttp), contentType, body)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("bad request")
	}
	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	jsonToken := entity.TokenDto{}
	if err := dec.Decode(&jsonToken); err != nil {
		return "", err
	}
	return jsonToken.Token, nil
}

func (h *handler) Registration(dto *entity.LoginDto) (string, error) {
	data, err := json.Marshal(dto)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	body := bytes.NewReader(data)
	contentType := "application/json"
	resp, err := h.client.Post(fmt.Sprintf("%s/auth/registration", h.conf.AddressHttp), contentType, body)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("bad request")
	}
	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	jsonToken := entity.TokenDto{}
	if err := dec.Decode(&jsonToken); err != nil {
		return "", err
	}
	return jsonToken.Token, nil
}
