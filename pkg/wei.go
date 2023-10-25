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

type Wei struct {
	firstdate time.Time
	lastdate  time.Time
	count     int
}

func New() *Wei {
	return &Wei{}
}

func (w *Wei) getPoints() plotter.XYs {
	pts := make(plotter.XYs, w.count)
	f, err := os.Open(input)
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
		date, err := time.Parse("20060102", record[0])
		if err != nil {
			log.Fatal(err)
		}
		pts[i].X = float64(date.Unix())
		pts[i].Y = weight
		i++
	}
	return pts
}

func (w *Wei) analyze() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	i := 0
	for {
		record, err := r.Read()
		if i == 0 {
			w.firstdate, err = time.Parse("20060102", record[0])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		w.lastdate, err = time.Parse("20060102", record[0])
		i++
	}
	w.count = i
}

func (w *Wei) Plot() {
	w.analyze()
	data := w.getPoints()

	p := plot.New()
	p.Title.Text = "Time Series"
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}
	p.Y.Label.Text = "Kg"

	p.X.Min = float64(w.lastdate.Unix())
	p.X.Max = float64(w.firstdate.Unix())
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
