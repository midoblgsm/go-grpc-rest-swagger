package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/mux"
)

func UnmarshalResponse(ctx context.Context, logger *log.Logger, r *http.Response, object interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Printf("ioutil.ReadAll failed %#v", err)
		return err
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		logger.Printf("json.Unmarshal failed%#v", err)
		return err
	}

	return nil
}

func FormatURL(url string, entries ...string) string {
	base := url
	if !strings.HasSuffix(url, "/") {
		base = fmt.Sprintf("%s/", url)
	}
	suffix := ""
	for _, entry := range entries {
		suffix = path.Join(suffix, entry)
	}
	return fmt.Sprintf("%s%s", base, suffix)
}

func HttpExecute(logger *log.Logger, httpClient *http.Client, requestType string, requestURL string, rawPayload interface{}) (*http.Response, error) {

	payload, err := json.MarshalIndent(rawPayload, "", " ")

	if err != nil {
		logger.Printf("Internal error marshalling params %#v", err)
		return nil, err
	}

	request, err := http.NewRequest(requestType, requestURL, bytes.NewBuffer(payload))

	if err != nil {
		logger.Printf("Error in creating request %#v", err)
		return nil, err
	}

	return httpClient.Do(request)
}

func HttpExecuteWithHeader(logger *log.Logger, httpClient *http.Client, requestType string, requestURL string, headers map[string]string) (*http.Response, error) {
	logger.Println("entering-HttpExecuteWithHeader")
	defer logger.Println("exiting-HttpExecuteWithHeader")
	logger.Printf("request URL %s \n", requestURL)
	request, err := http.NewRequest(requestType, requestURL, nil)
	if err != nil {
		logger.Printf("error-in-creating-request-%#v", err)
		return nil, err
	}

	for key, header := range headers {
		request.Header.Set(key, header)
	}

	return httpClient.Do(request)
}

func WriteResponse(w http.ResponseWriter, code int, object interface{}) {
	data, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	fmt.Fprintf(w, string(data))
}

func Unmarshal(r *http.Request, object interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}

	return nil
}

func UnmarshalDataFromRequest(r *http.Request, object interface{}) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func ExtractVarsFromRequest(r *http.Request, varName string) string {
	return mux.Vars(r)[varName]
}

func GetFromHeader(r *http.Request, header string) string {
	return r.Header.Get(header)
}

func ExtractParamsFromURL(r *http.Request, varName string) string {

	keys, ok := r.URL.Query()[varName]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return ""
	}

	return keys[0]
}

type encoding int

const (
	encodePath encoding = 1 + iota

	encodePathSegment

	encodeHost

	encodeZone

	encodeUserPassword

	encodeQueryComponent

	encodeFragment
)
