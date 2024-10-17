package ibkr

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func logRequest(req *http.Request, verbose bool) {
	log.Printf("---- ibkr http request: %s %s", req.Method, req.URL.String())

	if verbose {
		log.Println("headers:")
		for key, values := range req.Header {
			for _, value := range values {
				log.Printf("  %s: %s", key, value)
			}
		}

		if req.Body != nil {
			bodyBytes, err := io.ReadAll(req.Body)
			if err == nil {
				log.Printf("body: %s", string(bodyBytes))
				req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			} else {
				log.Printf("failed to read request body: %v", err)
			}
		}

		log.Println("---- end outbound request log")
	}
}

func logResponse(resp *http.Response, verbose bool) {
	log.Printf("---- ibkr http response, status: %s", resp.Status)

	if resp.Body != nil {
		if (resp.StatusCode > http.StatusOK) || verbose {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err == nil {
				log.Printf("response body: %s", string(bodyBytes))
				resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			} else {
				log.Printf("failed to read response body: %v", err)
			}

			log.Println("---- end inbound response log")
		}
	}
}

func logValidationErrors(err error) {
	for _, e := range err.(validator.ValidationErrors) {
		log.Println(e.Namespace())
		log.Println(e.Field())
		log.Println(e.StructNamespace())
		log.Println(e.StructField())
		log.Println(e.Tag())
		log.Println(e.ActualTag())
		log.Println(e.Kind())
		log.Println(e.Type())
		log.Println(e.Value())
		log.Println(e.Param())
		log.Println()
	}
}
