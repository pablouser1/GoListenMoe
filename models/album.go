package models

type Album struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	NameRomaji *string `json:"nameRomaji"`
	Image      *string `json:"image"`
}
