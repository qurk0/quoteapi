package quote

type Quote struct {
	ID     uint
	Author string
	Quote  string
}

type QuoteCreateRequest struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
