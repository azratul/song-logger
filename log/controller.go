package log

import (
	"encoding/csv"
	"os"
)

func (log *Log) Save() error {
	f, err := os.OpenFile(log.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = '|'
	//err = w.Write([]string{log.Year, log.Month, log.Day, log.Hour, log.Artist, log.Song})
	err = w.Write([]string{log.Day, log.Month, log.Year, log.Song, log.Artist, log.Artist, log.Hour})
	if err != nil {
		return err
	}

	w.Flush()

	return nil
}
