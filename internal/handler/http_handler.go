package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/leontofel/link-shortener/internal/service"
)

type HTTPHandler struct {
	svc *service.ShortenerService
}

func NewHTTPHandler(svc *service.ShortenerService) *HTTPHandler {
	return &HTTPHandler{svc: svc}
}

func (h *HTTPHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var data map[string]string
	_ = json.Unmarshal(body, &data)
	original := data["url"]

	code, err := h.svc.Shorten(original)
	if err != nil {
		http.Error(w, "Failed to shorten", 500)
		return
	}


    host := os.Getenv("LINK_CHECK_HOST")
    if host == "" {
        host = "8080"
    }

	resp := map[string]string{"short_url": host + ":8080/" + code}
	json.NewEncoder(w).Encode(resp)
}

func (h *HTTPHandler) ResolveURL(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		http.NotFound(w, r)
		return
	}

	url, err := h.svc.Resolve(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
