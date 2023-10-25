package wei

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

const input = "weight.csv"

func Plot() {
	var firstdate time.Time
	var lastdate time.Time
	randomPoints := func(n int) plotter.XYs {
		pts := make(plotter.XYs, n)

		f, err := os.Open(input)
		if err != nil {
			log.Fatal(err)
		}
		r := csv.NewReader(f)
		i := 0
		for {
			record, err := r.Read()
			if i == 0 {
				firstdate, err = time.Parse("20060102", record[0])
			}
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
			date, err := time.Parse("20060102", record[0])
			if err != nil {
				log.Fatal(err)
			}

			pts[i].X = float64(date.Unix())
			pts[i].Y = weight
			lastdate, err = time.Parse("20060102", record[0])
			i++
		}
		return pts
	}

	n := 5
	data := randomPoints(n)

	p := plot.New()
	p.Title.Text = "Time Series"
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}
	p.Y.Label.Text = "Kg"

	p.X.Min = float64(lastdate.Unix())
	p.X.Max = float64(firstdate.Unix())
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

	err = p.Save(30*vg.Centimeter, 20*vg.Centimeter, "timeseries.png")
	if err != nil {
		log.Panic(err)
	}
}
