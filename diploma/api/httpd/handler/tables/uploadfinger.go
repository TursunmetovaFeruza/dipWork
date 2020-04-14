package tables

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	db := PostSQLConfig()

	file, _ := c.FormFile("file")
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer src.Close()
	byt := make([]byte, file.Size)
	count, err := src.Read(byt)
	if err != nil {
		log.Fatal(err)
	}
	date := getDate(src, byt, count)
	fmt.Println(date)
	finger, er := getFingerPrints(src, byt, count)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(finger)
	defer db.Close()
}

func getDate(f multipart.File, byt []byte, count int) string {
	offs := strings.IndexAny(string(byt[:count]), "123456789")
	date := make([]byte, 19)
	count, err := f.ReadAt(date, int64(offs))
	if err != nil {
		fmt.Println(err)
	}
	return string(date[:count])
}

func getFingerPrints(f multipart.File, byt []byte, count int) (map[int][]byte, error) {
	fingerprints := make(map[int][]byte)
	for i := 1; i < 128; i++ {
		id := "#" + strconv.Itoa(i)
		offs := strings.LastIndex(string(byt[:count]), id)
		if offs == -1 {
			return fingerprints, errors.New("done")
		}
		finger := make([]byte, 1028)
		count, _ := f.ReadAt(finger, int64(offs))
		v := finger[:count]
		fingerprints[i] = v
	}
	return fingerprints, nil
}
