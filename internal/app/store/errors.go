package store

import "errors"

var (
	//  ErrRecordNotFound ...
	ErrRecordNotFound = errors.New("Record not found")
	TagWrongLength    = errors.New("Very long tag")
	TagIsNull         = errors.New("Tag is null")
	WrongUserTag      = errors.New("User in tag not found")
)
