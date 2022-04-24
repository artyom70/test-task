package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type api struct {
	hc      *http.Client
	baseURL string
}

func NewAPI(hc *http.Client, baseURL string) *api {
	return &api{
		hc:      hc,
		baseURL: baseURL,
	}
}

func (s *api) GetPrice(ctx context.Context, fsyms []string, tsyms []string) (PriceResponse, error) {
	apiURL, err := url.Parse(fmt.Sprintf("%s/data/pricemultifull", s.baseURL))
	if err != nil {
		return PriceResponse{}, fmt.Errorf("couldn't parse base_url, error: %v", err)
	}
	values := apiURL.Query()
	values.Add("fsyms", strings.Join(fsyms, ","))
	values.Add("tsyms", strings.Join(tsyms, ","))

	apiURL.RawQuery = values.Encode()
	resp, err := s.hc.Get(apiURL.String())
	if err != nil {
		return PriceResponse{}, fmt.Errorf("couldn't get price, error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return PriceResponse{}, fmt.Errorf("couldn't get price from api, status_code=%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	res := PriceResponse{Data: map[string]map[string]PriceData{}}
	if err := json.Unmarshal(body, &res); err != nil {
		return PriceResponse{}, fmt.Errorf("couldn't decode response body, error: %v", err)
	}

	return res, nil

}
