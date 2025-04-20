package kea_ca_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/klegul/kea-ca-client/model"
	"github.com/mitchellh/mapstructure"
	"io"
	"net/http"
	"strconv"
)

const ContentType string = "application/json"

type KeaCaClient struct {
	httpClient *http.Client
	url        string
}

type KeaCaClientConfig struct {
	Host string
	Port int
}

func New(config KeaCaClientConfig) *KeaCaClient {
	url := "http://" + config.Host + ":" + strconv.Itoa(config.Port)
	return &KeaCaClient{httpClient: &http.Client{}, url: url}
}

func (k *KeaCaClient) doRequest(arguments interface{}, result *model.Response) error {
	bodyStr, err := json.Marshal(arguments)
	if err != nil {
		return err
	}

	response, err := k.httpClient.Post(k.url, ContentType, bytes.NewBuffer(bodyStr))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	responseBody = responseBody[2 : len(responseBody)-2]

	if err := json.Unmarshal(responseBody, result); err != nil {
		return err
	}
	return nil
}

func (k *KeaCaClient) Lease4GetAll(arguments model.Lease4GetAllArguments) (model.Lease4GetAllResponse, error) {
	command := model.Command{
		Command:   "lease4-get-all",
		Arguments: arguments,
		Service:   []string{"dhcp4"},
	}

	response := model.Response{}
	err := k.doRequest(command, &response)
	if err != nil {
		return model.Lease4GetAllResponse{}, err
	}

	if response.Result != 0 {
		return model.Lease4GetAllResponse{}, errors.New(response.Text)
	}

	result := model.Lease4GetAllResponse{}
	if err = mapstructure.Decode(response.Arguments, &result); err != nil {
		return model.Lease4GetAllResponse{}, err
	}
	return result, nil
}
