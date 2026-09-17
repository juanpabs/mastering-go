package main

import (
	"fmt"
)

type MultipleLiveStream struct {
	CallSign   string
	Thumbnails []Thumbnail
}

type Thumbnail struct {
	URL string
}

func main() {
	fmt.Println("Hello, World!")
	formattedMultipleLiveStream := []MultipleLiveStream{
		{
			CallSign:   "KABC",
			Thumbnails: []Thumbnail{},
		},
	}
	var imageUrl string
	if len(formattedMultipleLiveStream[0].Thumbnails) != 0 {
		imageUrl = formattedMultipleLiveStream[0].Thumbnails[0].URL
	} else {
		imageUrl = ""
	}

	if imageUrl == "" {
		imageUrl, _ = getStationPlaceholderImageURL(formattedMultipleLiveStream[0].CallSign)
	}

	fmt.Println(imageUrl)
}

func getStationPlaceholderImageURL(callsign string) (string, error) {

	return "placeHolderImage", nil
}
