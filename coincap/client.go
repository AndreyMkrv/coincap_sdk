package coincap

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &LoggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetAssets() ([]Asset, error) {
	resp, error := c.client.Get(URL)

	if error != nil {
		log.Fatal(error)
	}

	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Fatal(error)
	}

	var r assetsResponse
	if error = json.Unmarshal(body, &r); error != nil {
		log.Fatal(error)
	}

	return r.Assets, nil
}
