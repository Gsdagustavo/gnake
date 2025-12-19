package core

import "errors"

// Image loading errors
var (
	ErrLoadImage    = errors.New("failed to load image")
	ErrInvalidImage = errors.New("invalid image")
	ErrConvertImage = errors.New("failed to convert image")
)

// Image drawing erros
var (
	ErrDrawImage               = errors.New("failed to draw image")
	ErrDrawGameBackgroundImage = errors.New("failed to draw game background image")
)

// IO errors
var (
	ErrReadFS = errors.New("failed to read filesystem")
)
