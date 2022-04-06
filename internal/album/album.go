package album

import (
	"encoding/json"
	"fmt"
	"github.com/ingkiller/hackernews/internal/client"
	"github.com/ingkiller/hackernews/internal/photo"
	"io/ioutil"
	"sync"
)

type Album struct {
	Id             int
	UserId         int
	Title          string
	NumberOfPhotos int
}

func GetAlbumByUserId(userId int) []Album {
	var result []Album
	albumUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/albums?userId=", userId)
	resp, err := client.MakeReq(albumUrl)
	if err != nil {
		fmt.Print("NewRequest: %v", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &result)

	ch := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(result))
		for j := 0; j < len(result); j++ {
			go func(a Album) {
				defer wg.Done()
				ch <- len(photo.GetPhotosByAlbumId(a.Id))
			}(result[j])
		}
		wg.Wait()
		close(ch)
	}()
	for i := 0; i < len(ch); i++ {
		result[i].NumberOfPhotos = ch[i]
	}

	return result
}
