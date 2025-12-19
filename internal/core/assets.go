package core

import (
	"embed"
	"errors"
	"gnake/internal/util"
	"image"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

//go:embed assets/*
var AssetsFS embed.FS

const assetsPath = "assets/"

func LoadAsset(path string) (*ebiten.Image, error) {
	file, err := AssetsFS.Open(assetsPath + path)
	if err != nil {
		return nil, errors.Join(errors.New("failed to open file"), err)
	}
	defer util.DeferFileClose(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, errors.Join(errors.New("failed to decode image"), err)
	}

	if img == nil {
		return nil, errors.New("invalid image")
	}

	ebitenImage, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, errors.Join(errors.New("failed to convert image to ebiten image"), err)
	}

	if ebitenImage == nil {
		return nil, errors.New("invalid ebiten image")
	}

	return ebitenImage, nil
}
