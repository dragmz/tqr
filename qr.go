package tqr

import (
	"github.com/yeqown/go-qrcode/v2"
)

const (
	BW = "▄"
	BB = " "
	WB = "▀"
	WW = "█"
)

type Qr struct {
	Width  int
	Height int
	Values []bool
}

func (q *Qr) Write(mat qrcode.Matrix) error {
	values := make([]bool, mat.Width()*mat.Height())

	i := 0
	mat.Iterate(qrcode.IterDirection_ROW, func(x, y int, s qrcode.QRValue) {
		values[i] = s.IsSet()
		i++
	})

	q.Width = mat.Width()
	q.Height = mat.Height()
	q.Values = values

	return nil
}

func (q *Qr) Close() error {
	return nil
}

func (q *Qr) String() string {
	var s string

	s += " "

	for i := 0; i < q.Width+2; i++ {
		s += BW
	}

	s += "\n "

	for y := 0; y < q.Height; y += 2 {
		for x := 0; x < q.Width; x++ {
			i := y*q.Width + x
			v := q.Values[i]

			var nv bool
			hn := y < q.Height-1
			if hn {
				nv = q.Values[i+q.Width]
			} else {
				hn = true
			}

			if x == 0 {
				s += WW
			}

			if v {
				if nv {
					s += BB
				} else {
					s += BW
				}
			} else {
				if nv {
					s += WB
				} else {
					s += WW
				}
			}

			if x == q.Width-1 {
				s += WW
			}
		}

		s += "\n "
	}

	return s
}

func New(value string) *Qr {
	q := &Qr{}

	qrc, _ := qrcode.New(value)
	qrc.Save(q)

	return q
}
