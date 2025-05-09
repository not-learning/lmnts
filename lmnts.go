package lmnts

import (
	"cmp"
	"log"
	"slices"
)

//const startW, startH float32 = 450, 1000
var scrW, scrH int
var ratW, ratH float32 = 1, 1

// TODO relative sizes
// TODO: think: 0 for 0 size, -1 for autosize
type size struct{ w, h float32 }

type sizeRel struct{ wr, hr float32 }

func (el *Lmnt) Size() (w, h float32) {
	return el.w, el.h
}

func (el *Lmnt) SetSize(w, h float32) {
	el.w, el.h = w, h
	//el.wr, el.hr = w * ratW, h * ratH
}

type point struct{ x, y float32 }

type rect struct {
	p1 *point
	p2 *point
}

func (el *Lmnt) Rect() (x1, y1, x2, y2 float32) {
	return el.p1.x, el.p1.y, el.p2.x, el.p2.y
}

func (el *Lmnt) SetRect(x1, y1, x2, y2 float32) {
	el.p1.x, el.p1.y = x1, y1
	el.p2.x, el.p2.y = x2, y2
}

type layout struct {
	row bool
	//name  string
	total *size
	kids  []*Lmnt
}

func newLayout() *layout {
	return &layout{ total: &size{} }
}

func (el *Lmnt) Row() bool { return el.row }

func (el *Lmnt) SetRow() { el.row = true }

// ### The Lmnt ###
type Lmnt struct {
	Name string
	*size
	*sizeRel
	*rect
	*layout
}

func New() *Lmnt {
	return &Lmnt{
		size:   &size{},
		sizeRel:   &sizeRel{},
		rect:   &rect{&point{}, &point{}},
		layout: newLayout(),
	}
}

func (el *Lmnt) Update(w, h int, rW, rH float32) {
	if scrW == w && scrH == h { return }
	scrW, scrH = w, h
	ratW, ratH = rW, rH
	el.DoAll()
}

func (el *Lmnt) Add(lls ...*Lmnt) {
	el.kids = append(el.kids, lls...)
}

func (el *Lmnt) AddN(n int, lls ...*Lmnt) {
	l := len(el.kids)
	if n > l {
		log.Println("AddN: index out of bounds")
	}
	if n < 0 { n += l+1 }
	if n < 0 {
		log.Println("AddN: index out of bounds")
	}
	el.kids = slices.Insert(el.kids, n, lls...)
}//*/

// todo: might be slow?
func (el *Lmnt) IsAdded(ll *Lmnt) bool {
	if el.kids == nil { return false }
	for _, v := range el.kids {
		if v == ll { return true }
	}
	return false
}

func (el *Lmnt) Del(ll *Lmnt) {
	el.kids = slices.DeleteFunc(el.kids, func(l *Lmnt) bool {
		return ll == l
	})
}//*/

func (el *Lmnt) Clear() {
	//a := []*Lmnt{}
	el.kids = nil
}

func (el *Lmnt) WalkDown(fn func(*Lmnt)) {
	// runtime.Breakpoint()
	fn(el)
	for _, ll := range el.kids {
		ll.WalkDown(fn)
	}
}

func (el *Lmnt) WalkUp(fn func(*Lmnt)) {
	// runtime.Breakpoint()
	for _, ll := range el.kids {
		ll.WalkUp(fn)
	}
	fn(el)
}

// ### The Mechanism
func setRelSize(el *Lmnt) {
	el.wr = el.w * ratW
	el.hr = el.h * ratW
}

func setTotalSizes(el *Lmnt) {
	el.total.w, el.total.h = 0, 0
	for _, ll := range el.kids {
		mw := myMax(ll.wr, ll.total.w)
		mh := myMax(ll.hr, ll.total.h)
		if el.row {
			el.total.w += mw
			el.total.h = myMax(mh, el.total.h)
		} else {
			el.total.w = myMax(mw, el.total.w)
			el.total.h += mh
		}
	}
}

func sizesList(fs, n float32, sl, st []float32) []float32 {
	delta := fs / n
	for {
		done := true
		if done { break }

		for i := range st {
			if st[i] > delta {
				sl[i] = st[i]
				fs -= st[i]
				st[i] = 0
				if n > 1 { n-- } // todo: is this really necessary?
				delta = fs / n
				done = false
			}
		}
	}
	for i := range sl {
		if sl[i] == 0 {
			sl[i] = delta
		}
	}
	return sl
}

func (el *Lmnt) setClmRect() {
	fs := myMax(el.p2.y-el.p1.y, el.total.h)
	n := float32(len(el.kids))
	sl := make([]float32, len(el.kids))
	st := make([]float32, len(el.kids))
	for i, v := range el.kids {
		if v.hr > 0 {
			sl[i] = v.hr
			fs -= v.hr
			if n > 1 { n-- } // todo: is this really necessary?
		} else {
			st[i] = v.total.h
		}
	}

	hl := sizesList(fs, n, sl, st)
	w := myMax(el.p2.x-el.p1.x, el.total.w)
	dy := el.p1.y
	for i, v := range el.kids {
		v.p1.x = el.p1.x
		v.p2.x = v.p1.x + cmp.Or(v.wr, w)
		v.p1.y = dy
		dy += hl[i]
		v.p2.y = dy
	}
}

func (el *Lmnt) setRowRect() {
	fs := myMax(el.p2.x-el.p1.x, el.total.w)
	n := float32(len(el.kids))
	sl := make([]float32, len(el.kids))
	st := make([]float32, len(el.kids))
	for i, v := range el.kids {
		if v.wr > 0 {
			sl[i] = v.wr
			fs -= v.wr
			if n > 1 { n-- } // todo: is this really necessary?
		} else {
			st[i] = v.total.w
		}
	}

	wl := sizesList(fs, n, sl, st)
	h := myMax(el.p2.y-el.p1.y, el.total.h)
	dx := el.p1.x

	for i, v := range el.kids {
		v.p1.y = el.p1.y
		v.p2.y = v.p1.y + cmp.Or(v.hr, h)

		v.p1.x = dx
		dx += wl[i]
		v.p2.x = dx
	}
}

func setRects(el *Lmnt) {
	if el.kids == nil {
		return
	}
	if el.row {
		el.setRowRect()
	} else {
		el.setClmRect()
	}
}

func (el *Lmnt) DoAll() {
	el.WalkDown(setRelSize)
	el.WalkUp(setTotalSizes)
	el.WalkDown(setRects)
}
