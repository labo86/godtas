package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func DecodeAsJson(w *httptest.ResponseRecorder, values interface{}) error {
	result := w.Result()

	if result.StatusCode != http.StatusOK {
		content, _ := ioutil.ReadAll(result.Body)
		return fmt.Errorf("http request deberia ser exitosa pero es %d : %v", result.StatusCode, string(content))
	}

	if err := json.NewDecoder(result.Body).Decode(values); err != nil {
		return fmt.Errorf("http no se pudo leer como json: %v", err)
	}
	return nil
}

func AssertStatusCode(r *http.Response, expectedStatusCode int) error {
	if r.StatusCode != expectedStatusCode {
		content, _ := ioutil.ReadAll(r.Body)
		return fmt.Errorf("el codigo esperado es %d pero fue %d : %q", expectedStatusCode, r.StatusCode, string(content))
	}
	return nil
}

func AssertContent(r *http.Response, expectedContent string) error {
	content, _ := ioutil.ReadAll(r.Body)
	contentStr := string(content)

	if contentStr != expectedContent {
		return fmt.Errorf("la respuesta esperada es %q pero fue : %q", expectedContent, content)
	}
	return nil
}
