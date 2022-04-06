package photo

import (
	"encoding/json"
	"fmt"
	"github.com/ingkiller/hackernews/internal/client"
	"io/ioutil"
)

type Photo struct {
	Id           int
	AlbumId      int
	Title        string
	Url          string
	ThumbnailUrl string
}

func GetPhotosByAlbumId(albumId int) []Photo {
	var result []Photo
	photoUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/photos?albumId=", albumId)
	resp, err := client.MakeReq(photoUrl)
	if err != nil {
		fmt.Print("NewRequest: %v", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &result)
	return result
}
