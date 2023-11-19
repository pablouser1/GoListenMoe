package models

type Song struct {
	ID       int64         `json:"id"`
	Title    string        `json:"title"`
	Sources  []interface{} `json:"sources"`
	Artists  []Album       `json:"artists"`
	Albums   []Album       `json:"albums"`
	Duration int64         `json:"duration"`
}
