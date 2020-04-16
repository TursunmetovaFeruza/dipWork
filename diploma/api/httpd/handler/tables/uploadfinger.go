package tables

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	model_finger "newsfeeder/httpd/models/finger"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func UploadFile(c *gin.Context) {
	db := PostSQLConfig()
	inserted := false
	file, _ := c.FormFile("file")
	students := c.PostForm("s")
	var ints []int
	err := json.Unmarshal([]byte("["+students+"]"), &ints)
	if err != nil {
		log.Fatal(err)
	}
	finger, _ := ReadUploadedfile(file)
	for i := 0; i < len(ints); i++ {
		p := model_finger.Finger{}
		rows := db.QueryRow("select finger from finger where finger=$1 or studentid=$2", finger[i+1], ints[i])
		err := rows.Scan(&p.Finger)
		if err == nil {
			inserted = true
			fmt.Println(p.Finger, finger[i+1])
		} else {
			inserted = false
		}
		if !inserted {
			index := 1
			rows, err := db.Query("select id from finger ORDER BY id DESC LIMIT 1")
			repos := []model_finger.Finger{}
			for rows.Next() {
				p := model_finger.Finger{}
				err = rows.Scan(&p.ID)
				if err != nil {
					fmt.Println(err)
				}
				repos = append(repos, p)
			}
			if repos != nil && len(repos) > 0 {
				index = repos[0].ID + 1
			}

			res, err := db.Exec("insert into finger (id, finger, studentid) values ( $1, $2,$3)",
				index, finger[i+1], ints[i])
			if err != nil {
				fmt.Println("create error:", err)
			}
			c.JSON(200, gin.H{
				"message": "Succes",
				"user":    &res,
			})
		}

	}
	defer db.Close()
}

func ReadUploadedfile(file *multipart.FileHeader) (map[int][]byte, string) {

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
	date, err := getDate(src, byt, count)
	finger, er := getFingerPrints(src, byt, count)
	if er != nil {
		fmt.Println(er)
	}
	return finger, date
}
func getDate(f multipart.File, byt []byte, count int) (string, error) {
	offs := strings.IndexAny(string(byt[:count]), "123456789")
	date := make([]byte, 19)
	count, err := f.ReadAt(date, int64(offs))
	if err != nil {
		return string(date[:count]), errors.New("done")
	}
	return string(date[:count]), nil
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

func UploadAttendance(c *gin.Context) {
	db := PostSQLConfig()
	file, _ := c.FormFile("file")
	students := c.PostForm("s")
	id := c.PostForm("id")
	eventId, _ := strconv.Atoi(id)
	var ints []int
	var stid []int
	var attend []int
	var apsent []int
	err := json.Unmarshal([]byte("["+students+"]"), &ints)
	if err != nil {
		log.Fatal(err)
	}
	finger, _ := ReadUploadedfile(file)
	for i := 0; i < len(ints); i++ {
		p := model_finger.Finger{}
		rows := db.QueryRow("select studentid from finger where finger=$1", finger[i+1])
		err := rows.Scan(&p.StudentId)
		if err == nil {
			stid = append(stid, p.StudentId)
		}
	}
	for i := 0; i < len(ints); i++ {
		v := 0
		for j := 0; j < len(stid); j++ {
			if ints[i] == stid[j] {
				v = stid[j]
			}
		}
		if v > 0 {
			attend = append(attend, v)
		} else {
			apsent = append(apsent, ints[i])
		}
	}
	rows, err := db.Query("update attendance set upsent = $1, attend = $2 where id = $3", pq.Array(apsent), pq.Array(attend), eventId)

	if err != nil {
		fmt.Println("create error:", err)
	}
	c.JSON(200, gin.H{
		"message": "Succes",
		"user":    &rows,
	})

	defer db.Close()
}
