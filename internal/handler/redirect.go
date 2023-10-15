package handler

import "net/http"

type redirectUsecase interface {
	Get(url string) (longUrl string, err error)
}

func redirectHandler(usecase redirectUsecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Path[1:]

		if shortUrl == "" {
			http.Error(w, emptyUrlMsg, http.StatusBadRequest)
			return
		}

		longUrl, err := usecase.Get(shortUrl)
		if err != nil {
			registerGlobalError(w, err)
			return
		}

		http.Redirect(w, r, longUrl, http.StatusFound)
	}
}
