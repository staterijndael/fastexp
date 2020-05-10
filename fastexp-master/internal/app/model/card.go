package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Card struct {
	Name      string
	ShortDesc string
	FullDesc  string
}

// Validate ...
func (c *Card) ValidateCard() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.ShortDesc, validation.Required, validation.Length(10, 104)),
		validation.Field(&c.FullDesc, validation.Required, validation.Length(10, 579)),
	)
}
