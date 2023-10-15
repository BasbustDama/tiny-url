package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/BasbustDama/tiny-url/internal/errors"
)

type registerUsecase interface {
	Create(url string) (shortUrl string, err error)
}

func registerHandler(usecase registerUsecase) http.HandlerFunc {
	type requestBody struct {
		Url string `json:"url"`
	}

	type responseBody struct {
		ShortUrl string `json:"shortUrl"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, methodNotAllowedMsg, http.StatusMethodNotAllowed)
			return
		}

		if r.Header.Get(contentType) != jsonContentType {
			http.Error(w, contentTypeHeaderMsg, http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		var request requestBody
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, jsonParseMsg, http.StatusBadRequest)
			return
		}

		if !validateUrl(request.Url) {
			slog.Info("URL is not valid", slog.String("url", request.Url))
			http.Error(w, invalidUrlMsg, http.StatusBadRequest)
			return
		}

		shortUrl, err := usecase.Create(request.Url)
		if err != nil {
			registerGlobalError(w, err)
			return
		}

		responseBody, err := json.Marshal(responseBody{ShortUrl: shortUrl})
		if err != nil {
			registerGlobalError(w, errors.ErrorInternal)
			return
		}

		w.Header().Add(contentType, jsonContentType)
		w.Write(responseBody)
		w.WriteHeader(http.StatusOK)
	}
}
