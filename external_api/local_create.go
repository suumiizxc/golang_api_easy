package external_api

import (
	base64Upload "github.com/heliojuniorkroger/golang-base64-upload"
)

func CreateLocal(content string) {

	err := base64Upload.Upload("kitten123.jpg", content)
	if err != nil {
		panic(err)
	}
}
