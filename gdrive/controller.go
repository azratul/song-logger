package gdrive

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"

	drive "google.golang.org/api/drive/v3"
)

func (c *Client) Upload(filename, mime string) error {
	today := time.Now().Format("2006-01-02")
	client := c.ServiceAccount()

	srv, err := drive.New(client)
	if err != nil {
		log.Error(err)
		return err
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Error(err)
		return err
	}
	fileInf, err := file.Stat()
	if err != nil {
		log.Error(err)
		return err
	}
	defer file.Close()
	basename := filepath.Base(filename)
	ext := filepath.Ext(basename)
	name := basename[:len(basename)-len(ext)]
	newfilename := fmt.Sprintf("%s-%s%s", name, today, ext)

	f := &drive.File{Name: newfilename, Parents: c.Parents}
	_, err = srv.Files.
		Create(f).
		ResumableMedia(context.Background(), file, fileInf.Size(), mime).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (c *Client) Get() ([]*drive.File, error) {
	client := c.ServiceAccount()

	srv, err := drive.New(client)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	r, err := srv.Files.List().PageSize(10).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Error(fmt.Sprintf("Unable to retrieve files: %v", err))
		return nil, err
	}

	if len(r.Files) == 0 {
		log.Info("No files found.")
		return nil, errors.New("No Files found")
	}

	return r.Files, nil
}

func (c *Client) Delete(fileID string) error {
	client := c.ServiceAccount()

	srv, err := drive.New(client)
	if err != nil {
		log.Error(err)
		return err
	}

	err = srv.Files.Delete(fileID).Do()
	if err != nil {
		log.Error(fmt.Sprintf("An error occurred while deleting a file: %v\n", err))
		return err
	}

	return nil
}

func (c *Client) ServiceAccount() *http.Client {
	conf := &jwt.Config{
		Email:      c.ClientEmail,
		PrivateKey: []byte(c.PrivateKey),
		Scopes: []string{
			drive.DriveScope,
		},
		TokenURL: google.JWTTokenURL,
	}

	return conf.Client(oauth2.NoContext)
}
