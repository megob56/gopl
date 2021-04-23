package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var (
	width, height float64 = 600, 320
	xyscale               = width / 2 / xyrange // pixels per x or y unit
	zscale                = height * 0.4        // pixels per z unit
)

type CornerType int

const (
	middle CornerType = 0 // not the peak or valley of surface
	peak   CornerType = 1 // the peak of surface corner
	valley CornerType = 2 // the valley of surface corner
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// parse width/height
		if err := request.ParseForm(); err != nil {
			log.Printf("parseForm: %v\n", err)
		}
		for k, v := range request.Form {
			var err error

			if k == "width" {
				if width, err = strconv.ParseFloat(v[0], 64); err != nil {
					width = 600
					log.Printf("parseWidth: %v\n", err)
				}
			}

			if k == "height" {
				if height, err = strconv.ParseFloat(v[0], 64); err != nil {
					height = 320
					log.Printf("parseHeight: %v\n", err)
				}
			}
		}

		// need recalculate
		xyscale = width / 2 / xyrange // pixels per x or y unit
		zscale = height * 0.4         // pixels per z unit

		writer.Header().Set("Content-Type", "image/svg+xml")
		surface(writer)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns= 'http://www.w3.org/svg' "+"style= 'stroke: grey; fill: white; stroke-width: 0.7'"+"width='%d' height='%d' > ", int64(width), int64(height))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ct, err := corner(i+j, j)
			if err != nil {
				continue
			}
			bx, by, ct1, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, ct3, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, ct2, err := corner(i+j, j+1)
			if err != nil {
				continue
			}

			var color string
			if ct == peak || ct1 == peak || ct2 == peak || ct3 == peak {
				color = "#f00"
			} else if ct == valley || ct1 == valley || ct2 == valley || ct3 == valley {
				color = "#00f"
			} else {
				color = "grey"
			}

			fmt.Fprintf(w, "<polygon points = '%g, %g %g, %g %g, %g %g, %g' style='stroke: %s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)

		}
	}
	fmt.Println(w, "</svg>")
}

func corner(i, j int) (float64, float64, CornerType, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, ct := f(x, y)

	//check exercise 3.1
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, fmt.Errorf("invalid value")
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, ct, nil
}

//eggbox
// func f(x, y float64) (float64, CornerType) {
// 	return math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12
// }

func f(x, y float64) (float64, CornerType) {
	r := math.Hypot(x, y)
	ct := middle

	if math.Abs(r-math.Tan(r)) < 3 {
		ct = peak
		if 2*(math.Sin(r)-r*math.Cos(r))-r*r*math.Sin(r) > 0 {
			ct = valley
		}
	}

	return math.Sin(r) / r, ct
}
