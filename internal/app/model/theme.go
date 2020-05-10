package model

import validation "github.com/go-ozzo/ozzo-validation"

// Theme ...
type Theme struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TagTheme ...
type TagTheme struct {
	ThemeID int    `json:"themeid"`
	Text    string `json:"text"`
}

// ValidateTheme ...
func (th *Theme) ValidateTheme() error {
	return validation.ValidateStruct(
		th,
		validation.Field(&th.Title, validation.Required, validation.Length(3, 50)),
		validation.Field(&th.Description, validation.Length(3, 150)),
	)
}
