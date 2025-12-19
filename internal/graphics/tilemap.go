package graphics

import (
	"encoding/json"
	"errors"
	"os"
)

type TilemapLayerJSON struct {
	Data   []int `json:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

type TylemapJSON struct {
	Layers []TilemapLayerJSON `json:"layers"`
}

func NewTilemapJSON(path string) (*TylemapJSON, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Join(errors.New("failed to read tiles"), err)
	}

	var tilemapJSON TylemapJSON
	err = json.Unmarshal(data, &tilemapJSON)
	if err != nil {
		return nil, errors.Join(errors.New("failed to parse tiles"), err)
	}

	return &tilemapJSON, nil
}
