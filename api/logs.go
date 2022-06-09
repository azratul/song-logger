package api

import (
	"cron/config"
	"cron/gdrive"
	logs "cron/log"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func WriteLog() {
	res, err := http.Get(config.Conf.IcecastURL)
	if err != nil {
		log.Error(err)
	}

	defer res.Body.Close()

	var radio config.Radio

	if err := json.NewDecoder(res.Body).Decode(&radio); err != nil {
		log.Error(err)
		return
	}

	song := fmt.Sprintf("%s-%s", radio.Icestats.Source.Artist, radio.Icestats.Source.Title)
	songBase64 := base64.StdEncoding.EncodeToString([]byte(song))
	lastsong, _ := ioutil.ReadFile(config.Conf.LastSong)

	if len(lastsong) > 0 {
		tmp := string(lastsong[:len(lastsong)-1])

		if tmp == songBase64 {
			log.Println("OK")
			return
		}
	}

	LastSong(config.Conf.LastSong, songBase64)

	l := logs.Log{
		Filename: config.Conf.Logs.Filename,
		Artist:   radio.Icestats.Source.Artist,
		Song:     radio.Icestats.Source.Title,
		Time:     time.Now(),
		Hour:     time.Now().Format("15:04"),
		Day:      time.Now().Format("02"),
		Month:    time.Now().Format("01"),
		Year:     time.Now().Format("2006"),
	}

	err = l.Save()
	if err != nil {
		log.Error(err)
		return
	}

	log.Println("OK")
}

func UploadLog() {
	filename := config.Conf.Logs.Filename
	mime := config.Conf.Logs.MimeType
	var gd gdrive.Client
	gd.ClientEmail = config.Conf.GDrive.ClientEmail
	gd.PrivateKey = config.Conf.GDrive.PrivateKey
	gd.Parents = config.Conf.GDrive.Parents

	err := gd.Upload(filename, mime)
	if err != nil {
		log.Error(err)
		return
	}

	log.Println("OK")
}

func GetLogs() {
	var gd gdrive.Client
	gd.ClientEmail = config.Conf.GDrive.ClientEmail
	gd.PrivateKey = config.Conf.GDrive.PrivateKey
	r, err := gd.Get()
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Printf("%+v", r)

	log.Println("OK")
}

func DeleteLogs() {
	/*
		type Files struct {
			ID []string `json:"file_id"`
		}

		var files Files

		if err := c.ShouldBindJSON(&files); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}

		for _, x := range files.ID {
			var gd gdrive.Client
			gd.ClientEmail = config.Conf.GDrive.ClientEmail
			gd.PrivateKey = config.Conf.GDrive.PrivateKey
			err := gd.Delete(x)
			if err != nil {
				log.Error(err)
				return
			}
		}
	*/

	log.Println("OK")
}

func LastSong(file, song string) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(f, song)
	if err != nil {
		log.Error(err)
	}
}
