package quote

import (
	"crypto/rand"
	"errors"
	"math/big"
	"slices"
)

type QuotesRepository struct {
	quotes  map[string][]Quote
	authors []string
}

var ID uint = 1

func NewQuoteRepository() *QuotesRepository {
	return &QuotesRepository{
		quotes:  make(map[string][]Quote),
		authors: make([]string, 0),
	}
}

func (qd *QuotesRepository) CreateQuote(q QuoteCreateRequest) Quote {
	createdQuote := Quote{
		ID:     ID,
		Author: q.Author,
		Quote:  q.Quote,
	}
	ID += 1
	qd.quotes[q.Author] = append(qd.quotes[q.Author], createdQuote)
	if !(slices.Contains(qd.authors, q.Author)) {
		qd.authors = append(qd.authors, q.Author)
	}
	return createdQuote
}

func (qd QuotesRepository) GetAllQuotes() []Quote {
	result := make([]Quote, 0)
	// Идем по всем авторам, опуская имя автора как ключ, так как оно есть в структуре цитат
	for _, value := range qd.quotes {
		// Закидываем каждую цитату в лист, каждая цитата содержит в себе текст цитаты, автора и ID
		for _, quote := range value {
			result = append(result, quote)
		}
	}
	return result
}

func (qd QuotesRepository) GetRandomQuote() (Quote, error) {
	// Выбрали случайного автора
	randomAuthorNumber, err := rand.Int(rand.Reader, big.NewInt(int64(len(qd.authors))))
	if err != nil {
		return Quote{}, err
	}
	randomAuthor := qd.authors[randomAuthorNumber.Int64()]

	randomQuoteNumber, err := rand.Int(rand.Reader, big.NewInt(int64(len(qd.quotes[randomAuthor]))))
	if err != nil {
		return Quote{}, err
	}
	randomQuote := qd.quotes[randomAuthor][randomQuoteNumber.Int64()]
	return randomQuote, nil

}

func (qd QuotesRepository) GetAuthorQuotes(author string) []Quote {
	return qd.quotes[author]
}

func (qd *QuotesRepository) DeleteQuote(id uint) error {
	for _, author := range qd.authors {
		for idx, quote := range qd.quotes[author] {
			if quote.ID == id {
				qd.quotes[author] = append(qd.quotes[author][:idx], qd.quotes[author][idx+1:]...)
				return nil
			}
		}
	}
	return errors.New("Нет цитаты с данным ID")
}
