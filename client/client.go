package client

import (
	"encoding/json"
	"github.com/hikasami/kodik-api/errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultDomain = "kodikapi.com"
)

// Client хранит настройки для работы с API.
type Client struct {
	Token      string
	BaseURL    string
	HttpClient *http.Client
}

// DefaultClient – глобальный клиент, инициализируется через Init.
var DefaultClient = &Client{}

// NewClient создаёт новый клиент с заданным токеном API и выбором протокола (HTTP или HTTPS).
func NewClient(token string, useHTTPS bool) *Client {
	schema := "http"
	if useHTTPS {
		schema = "https"
	}

	return &Client{
		Token:   token,
		BaseURL: schema + "://" + DefaultDomain,
		HttpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// Init инициализирует глобальный клиент. После вызова этой функции достаточно вызывать функции API, передаётся client.DefaultClient.
func Init(apiToken string, useHTTPS bool) {
	DefaultClient = NewClient(apiToken, useHTTPS)
}

// DoRequest выполняет HTTP-запрос к указанному эндпоинту с заданными параметрами и декодирует ответ в result.
func (c *Client) DoRequest(
	method string,
	endpoint string,
	params map[string]string,
	result interface{},
) error {
	var req *http.Request
	var err error
	fullURL := c.BaseURL + endpoint

	if params == nil {
		params = map[string]string{}
	}
	params["token"] = c.Token

	if strings.ToUpper(method) == "GET" {
		query := url.Values{}
		for key, value := range params {
			query.Set(key, value)
		}

		fullURL = fullURL + "?" + query.Encode()

		req, err = http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return err
		}
	} else if strings.ToUpper(method) == "POST" {
		form := url.Values{}
		for key, value := range params {
			form.Add(key, value)
		}
		req, err = http.NewRequest("POST", fullURL, strings.NewReader(form.Encode()))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		return errors.ErrUnsupportedMethod
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.ErrUnexpectedStatus
	}

	return json.NewDecoder(resp.Body).Decode(result)
}
