package wei

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

// 書き込み先と書き込む内容を指定する
// すでに存在する場合と、ない場合の処理を分ける?

type Entry struct {
	weight float64
}

func NewEntry(weight float64) *Entry {
	return &Entry{weight: weight}
}

func (e *Entry) Record() error {
	// ファイルがない場合は作る
	if _, err := os.Stat("hello"); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create("hello")
		defer f.Close()
		if err != nil {
			return err
		}
		fmt.Println("created!")
	}
	f, err := os.Open("hello")
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
