package quote

import (
	"encoding/json"
	"net/http"
	"quotesapi/pkg/resp"
	"strconv"
)

type QuoteHandler struct {
	repo *QuotesRepository
}

func NewQuoteHandler(router *http.ServeMux) {
	handler := &QuoteHandler{
		repo: NewQuoteRepository(),
	}
	router.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetQuote(w, r)
		case http.MethodPost:
			handler.Create(w, r)
		}
	})
	router.HandleFunc("GET /quotes/random", handler.GetRand())
	router.HandleFunc("DELETE /quotes/{id}", handler.Delete())
}

func (handler *QuoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload QuoteCreateRequest
	err := json.NewDecoder(r.Body).Decode(&payload)

	// Ошибка декодирования, если одного из полей нет
	if err != nil {
		resp.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Ошибка, если поле автора цитаты пустое (у цитаты должен быть автор)
	if payload.Author == "" {
		resp.Json(w, "ERROR : у цитаты должен быть автор", http.StatusBadRequest)
		return
	}

	// Ошибка, если поле цитаты пустое (цитата без содержания не цитата)
	if payload.Quote == "" {
		resp.Json(w, "ERROR : у цитаты должна быть цитата", http.StatusBadRequest)
		return
	}

	responce := handler.repo.createQuote(payload)
	resp.Json(w, responce, http.StatusOK)
	return
}

func (handler *QuoteHandler) GetQuote(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Query().Has("author"):
		handler.GetFiltered(w, r)
	default:
		handler.GetAll(w, r)
	}

}

func (handler *QuoteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	responce := handler.repo.getAllQuotes()
	resp.Json(w, responce, http.StatusOK)
}

func (handler *QuoteHandler) GetRand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responce, err := handler.repo.getRandomQuote()
		if err != nil {
			resp.Json(w, responce, http.StatusInternalServerError)
		}
		resp.Json(w, responce, http.StatusOK)
	}
}

func (handler *QuoteHandler) GetFiltered(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author == "" {
		resp.Json(w, "Не указан автор в query-параметрах", http.StatusBadRequest)
		return
	}

	responce := handler.repo.getAuthorQuotes(author)
	if responce == nil {
		resp.Json(w, "У данного автора нет цитат", http.StatusNotFound)
		return
	}

	resp.Json(w, responce, http.StatusOK)
}

func (handler *QuoteHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			resp.Json(w, "Неверно указан ID", http.StatusBadRequest)
			return
		}
		err = handler.repo.deleteQuote(uint(id))
		if err != nil {
			resp.Json(w, "Цитата не найдена", http.StatusNotFound)
			return
		}
		resp.Json(w, "Цитата успешно удалена", http.StatusNoContent)
	}
}
