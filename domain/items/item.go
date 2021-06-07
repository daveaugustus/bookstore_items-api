package items

type Item struct {
	Id                string      `json:"id"`
	Seller            string      `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Pictures          []Picture   `json:"picture"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	URL string `json:"url"`
}
