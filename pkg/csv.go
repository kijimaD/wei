package wei

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type Entry struct {
	config *config
	weight float64
}

func NewEntry(config *config, weight float64) *Entry {
	return &Entry{
		config: config,
		weight: weight,
	}
}

func (e *Entry) Record() error {
	// CSVファイルがない場合は作る
	if _, err := os.Stat(e.config.CsvPath); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(e.config.CsvPath)
		defer f.Close()
		if err != nil {
			return err
		}
	}
	f, err := os.OpenFile(e.config.CsvPath, os.O_APPEND|os.O_WRONLY, 0600) // 追加モード
	if err != nil {
		return err
	}
	defer f.Close()
	if err != nil {
		return err
	}
	err = e.Write(f)
	if err != nil {
		return err
	}
	return nil
}

// エントリを書き込む
func (e *Entry) Write(w io.Writer) error {
	csv := csv.NewWriter(w)
	datestr := fmt.Sprint(time.Now().Format("20060102"))
	err := csv.Write([]string{datestr, fmt.Sprintf("%.2f", e.weight)})
	if err != nil {
		return err
	}
	csv.Flush() // Writeだと内部バッファに書き込まれるだけなのでFlush()が必要

	return nil
}
