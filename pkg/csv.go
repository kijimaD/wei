package wei

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// 書き込み先と書き込む内容を指定する
// すでに存在する場合と、ない場合の処理を分ける?

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
	const output = "hello"

	// ファイルがない場合は作る
	if _, err := os.Stat("hello"); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(output)
		defer f.Close()
		if err != nil {
			return err
		}
		fmt.Println("created!")
	}
	f, err := os.OpenFile(output, os.O_APPEND|os.O_WRONLY, 0600) // 追加モード
	if err != nil {
		log.Fatal(err)
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
	// TODO: 日付、重さに変更する
	err := csv.Write([]string{"hello", "world"})
	if err != nil {
		return err
	}
	csv.Flush() // Writeだと内部バッファに書き込まれるだけなのでFlush()が必要

	return nil
}
