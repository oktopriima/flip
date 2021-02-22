/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22:25
**/

package sbig_api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/oktopriima/flip/helper/configurations"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Client struct {
	BaseURL string `json:"base_url"`
	Token   string `json:"token"`
}

func NewClient(cfg configurations.Config) (*Client, error) {
	var client Client

	c := cfg.GetStringMap("sbig-api")
	jsonb, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(jsonb, &client); err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err)
	}

	return &client, nil
}

func (c *Client) CreateDisbursement(request CreateDisbursementRequest) (*Response, error) {
	payload, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	body := strings.NewReader(payload.Encode())

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing url. err: %v", err)
	}

	u.Path = path.Join(u.Path, "disburse")

	r, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %v", err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.GenerateAuthorization()))

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to perform http request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("an error occured: %v", resp.Status)
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %v", err)
	}

	var response Response
	err = json.Unmarshal(bodyResp, &response)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return &response, nil
}

func (c *Client) FindDisbursement(request FindDisbursementRequest) (*Response, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing url. err: %v", err)
	}

	u.Path = path.Join(u.Path, fmt.Sprintf("disburse/%d", request.TransactionID))

	r, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %v", err)
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", c.GenerateAuthorization()))

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to perform http request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("an error occured: %v", resp.Status)
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %v", err)
	}

	var response Response
	err = json.Unmarshal(bodyResp, &response)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	return &response, nil
}

func (c *Client) GenerateAuthorization() string {
	auth := fmt.Sprintf("%s:", c.Token)
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
