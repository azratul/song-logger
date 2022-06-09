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
	err = w.Write([]string{log.Year, log.Month, log.Day, log.Hour, log.Artist, log.Song})
	if err != nil {
		return err
	}

	w.Flush()

	return nil
}
