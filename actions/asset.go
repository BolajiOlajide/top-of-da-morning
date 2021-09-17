package actions

import (
	"os"

	"github.com/BolajiOlajide/top-of-da-morning/constants"
)

func loadAsset() (*os.File, error) {
	videoFileName := "top_of_da_morning.mov"
	assetPath := constants.AssetsDirectoryPath + videoFileName
	file, err := os.Open(assetPath)

	if err != nil {
		return nil, err
	}

	return file, nil
}
