package models

type Song struct {
	ID       int64         `json:"id"`
	Title    string        `json:"title"`
	Sources  []interface{} `json:"sources"`
	Artists  []Album       `json:"artists"`
	Albums   []Album       `json:"albums"`
	Duration int64         `json:"duration"`
}

type Album struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	NameRomaji *string `json:"nameRomaji"`
	Image      *string `json:"image"`
}
