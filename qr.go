package tqr

import (
	"github.com/yeqown/go-qrcode/v2"
)

const (
	bw = "▄"
	bb = " "
	wb = "▀"
	ww = "█"
)

var (
	dark  = [4]string{bw, bb, wb, ww}
	light = [4]string{wb, ww, bw, bb}
)

const (
	io = 0
	ii = 1
	oi = 2
	oo = 3
)

type Qr struct {
	Width  int
	Height int
	Values []bool

	items  [4]string
	invert bool
}

type Option func(q *Qr)

func Invert() Option {
	return func(q *Qr) {
		q.invert = true
	}
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
		if !q.invert {
			s += q.items[io]
		}
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
				s += q.items[oo]
			}

			if v {
				if nv {
					s += q.items[ii]
				} else {
					s += q.items[io]
				}
			} else {
				if nv {
					s += q.items[oi]
				} else {
					s += q.items[oo]
				}
			}

			if x == q.Width-1 {
				s += q.items[oo]
			}
		}

		s += "\n "
	}

	return s
}

func New(value string, opts ...Option) *Qr {
	q := &Qr{}

	for _, opt := range opts {
		opt(q)
	}

	if q.invert {
		q.items = light
	} else {
		q.items = dark
	}

	qrc, _ := qrcode.NewWith(value,
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionLow),
	)
	qrc.Save(q)

	return q
}
