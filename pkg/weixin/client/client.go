// Copyright 2022-2023 The pmsg Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	httpClient "github.com/lenye/pmsg/pkg/http/client"
)

// CheckHttpResponseStatusCode 检查HTTP响应状态码
func CheckHttpResponseStatusCode(method, url string, statusCode int) error {
	if statusCode/100 != 2 {
		return fmt.Errorf("%w; http response status code: %v, %s %s", httpClient.ErrRequest, statusCode, method, url)
	}
	return nil
}

func GetJSON(url string, respBody any) (http.Header, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w; %s %s, %v", httpClient.ErrRequest, http.MethodGet, url, err)
	}
	defer resp.Body.Close()

	if err := CheckHttpResponseStatusCode(http.MethodGet, url, resp.StatusCode); err != nil {
		return nil, err
	}

	if respBody == nil {
		return resp.Header, nil
	}

	return resp.Header, json.NewDecoder(resp.Body).Decode(respBody)
}

// PostJSON http post json
func PostJSON(url string, reqBody, respBody any) (http.Header, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Post(url, httpClient.HdrValContentTypeJson, bytes.NewReader(buf.Bytes()))
	if err != nil {
		return nil, fmt.Errorf("%w; %s %s, %v", httpClient.ErrRequest, http.MethodPost, url, err)
	}
	defer resp.Body.Close()

	if err := CheckHttpResponseStatusCode(http.MethodPost, url, resp.StatusCode); err != nil {
		return nil, err
	}

	if respBody == nil {
		return resp.Header, nil
	}

	return resp.Header, json.NewDecoder(resp.Body).Decode(respBody)
}

func PostFileJSON(url, fieldName, fileName string, respBody any) (http.Header, error) {
	resp, err := httpClient.PostFile(url, fieldName, fileName)
	if err != nil {
		return nil, fmt.Errorf("%w; %s %s, %v", httpClient.ErrRequest, http.MethodPost, url, err)
	}
	defer resp.Body.Close()

	if err := CheckHttpResponseStatusCode(http.MethodPost, url, resp.StatusCode); err != nil {
		return nil, err
	}

	if respBody == nil {
		return resp.Header, nil
	}

	return resp.Header, json.NewDecoder(resp.Body).Decode(respBody)
}
