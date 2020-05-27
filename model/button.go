package model

type Button struct {
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	URL       string    `json:"url"`
	SubButton []*Button `json:"sub_button"`
}

type Buttons struct {
	Button []*Button `json:"button"`
}
