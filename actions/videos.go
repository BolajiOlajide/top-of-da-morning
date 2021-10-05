package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/BolajiOlajide/top-of-da-morning/constants"
	"github.com/BolajiOlajide/top-of-da-morning/models"
)

func loadAllVideos() ([]models.TikTokVideo, error) {
	file, err := os.Open(constants.VideoJSONFilePath)

	if err != nil {
		return nil, err
	}

	transformedFile, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	videos := []models.TikTokVideo{}
	err = json.Unmarshal(transformedFile, &videos)

	return videos, err
}

// PickRandomVideo pick a video at random to be tweeted
func pickRandomVideo() (models.TikTokVideo, error) {
	videos, err := loadAllVideos()

	if err != nil {
		return models.TikTokVideo{}, fmt.Errorf("An error occurred while loading the videos from the JSON. %s", err)
	}

	seedSrc := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seedSrc)

	randomIndex := r.Intn(len(videos))

	return videos[randomIndex], nil
}
