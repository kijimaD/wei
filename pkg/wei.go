package wei

import (
	"encoding/csv"
	"image/color"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

const input = "weight.csv"
const output = "timeseries.png"

type Wei struct {
	firstdate time.Time
	lastdate  time.Time
	count     int
	min       float64
	max       float64
}

func New() *Wei {
	return &Wei{}
}

func (w *Wei) Plot() {
	w.analyze()

	p := plot.New()
	p.Title.Text = "Time Series"
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}
	p.Y.Label.Text = "Kg"

	p.X.Min = float64(w.lastdate.Unix())
	p.X.Max = float64(w.firstdate.Unix())
	p.Y.Min = w.min
	p.Y.Max = w.max
	p.Add(plotter.NewGrid())

	data := w.getPoints()
	line, points, err := plotter.NewLinePoints(data)
	if err != nil {
		log.Panic(err)
	}
	line.Color = color.RGBA{G: 255, A: 255}
	points.Shape = draw.CircleGlyph{}
	points.Color = color.RGBA{R: 255, A: 255}

	p.Add(line, points)

	err = p.Save(30*vg.Centimeter, 20*vg.Centimeter, output)
	if err != nil {
		log.Panic(err)
	}
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

// なめて値を入れる
func (w *Wei) analyze() {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	i := 0
	var weights []float64
	for {
		record, err := r.Read()
		if i == 0 {
			date, err := time.Parse("20060102", record[0])
			if err != nil {
				log.Fatal(err)
			}
			w.firstdate = date
		}
		if err == io.EOF {
			break
		}
		w.lastdate, err = time.Parse("20060102", record[0])
		weight, err := strconv.ParseFloat(record[1], 64)
		weights = append(weights, weight)
		i++
	}
	w.min = slices.Min(weights)
	w.max = slices.Max(weights)
	w.count = i
}
