package main

import (
	"cron/api"
	"cron/config"
	"flag"

	log "github.com/sirupsen/logrus"

	"os"
)

func main() {
	if os.Getenv("LASTSONG") == "" ||
		os.Getenv("LOGS_FILENAME") == "" ||
		os.Getenv("LOGS_MIMETYPE") == "" ||
		os.Getenv("GDRIVE_PARENTS") == "" ||
		os.Getenv("GDRIVE_CLIENTEMAIL") == "" ||
		os.Getenv("GDRIVE_PRIVATEKEY") == "" {
		log.Println("All env vars must be set!!!")
		return
	}

	config.Conf.LastSong = os.Getenv("LASTSONG")
	config.Conf.Logs.Filename = os.Getenv("LOGS_FILENAME")
	config.Conf.Logs.MimeType = os.Getenv("LOGS_MIMETYPE")
	config.Conf.GDrive.ClientEmail = os.Getenv("GDRIVE_CLIENTEMAIL")
	config.Conf.GDrive.PrivateKey = os.Getenv("GDRIVE_PRIVATEKEY")
	config.Conf.GDrive.Parents = []string{os.Getenv("GDRIVE_PARENTS")}

	write := flag.Bool("write", false, "Write to a File")
	upload := flag.Bool("upload", false, "Upload to Google Drive")

	flag.Parse()

	if *write {
		api.WriteLog()
	}

	if *upload {
		api.UploadLog()
	}
	//api.GetLogs
	//api.DeleteLogs
}
