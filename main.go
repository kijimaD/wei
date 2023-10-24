package main

import (
	"encoding/csv"
	"image/color"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	xticks := plot.TimeTicks{Format: "2006-01-02\n15:04"}

	// randomPoints returns some random x, y points
	// with some interesting kind of trend.
	randomPoints := func(n int) plotter.XYs {
		const (
			month = 1
			day   = 1
			hour  = 1
			min   = 1
			sec   = 1
			nsec  = 1
		)
		pts := make(plotter.XYs, n)

		f, err := os.Open("weight.csv")
		if err != nil {
			log.Fatal(err)
		}
		r := csv.NewReader(f)
		i := 0
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			weight, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				log.Fatal(err)
			}
			date := time.Date(2020+i, month, day, hour, min, sec, nsec, time.UTC).Unix()
			pts[i].X = float64(date)
			pts[i].Y = weight
			i++
		}
		return pts
	}

	n := 5
	data := randomPoints(n)

	p := plot.New()
	p.Title.Text = "Time Series"
	p.X.Tick.Marker = xticks
	p.Y.Label.Text = "Kg"

	p.X.Min = float64(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC).Unix())
	p.X.Max = float64(time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC).Unix())
	p.Y.Min = 40
	p.Y.Max = 80
	p.Add(plotter.NewGrid())

	line, points, err := plotter.NewLinePoints(data)
	if err != nil {
		log.Panic(err)
	}
	line.Color = color.RGBA{G: 255, A: 255}
	points.Shape = draw.CircleGlyph{}
	points.Color = color.RGBA{R: 255, A: 255}

	p.Add(line, points)

	err = p.Save(30*vg.Centimeter, 5*vg.Centimeter, "timeseries.png")
	if err != nil {
		log.Panic(err)
	}
}
