package core

import "errors"

// Image loading errors
var (
	ErrLoadGameImage    = errors.New("failed to load game background image")
	ErrInvalidGameImage = errors.New("invalid game image")
	ErrConvertImage     = errors.New("failed to convert game image")
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
