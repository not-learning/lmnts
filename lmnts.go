package lmnts

import (//"fmt"
	"cmp"
	"slices"
)

// TODO relative sizes
// TODO: think: 0 for 0 size, -1 for autosize
type size struct{ w, h float32 }

func (el *Lmnt) Size() (w, h float32) {
	return el.w, el.h
}

func (el *Lmnt) SetSize(w, h float32) {
	el.w, el.h = w, h
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
	row   bool
	//name  string
	total *size
	kids  []*Lmnt
}

func newLayout() *layout {
	return &layout{total: &size{}}
}

func (el *Lmnt) SetRow() { el.row = true }

// ### The Lmnt ###
type Lmnt struct {
	Name string
	*size
	*rect
	//*Font
	*layout
}

func New() *Lmnt {
	return &Lmnt{
		size:   &size{},
		rect:   &rect{&point{}, &point{}},
		layout: newLayout(),
	}
}

func (el *Lmnt) Add(lls ...*Lmnt) {
	for _, v := range lls {
		el.kids = append(el.kids, v)
	}
}

func (el *Lmnt) Del(lls ...*Lmnt) {
	for _, v := range lls {
		el.kids = slices.DeleteFunc(el.kids, func(l *Lmnt) bool {
			return v == l
		})
	}
}

func (el *Lmnt) Clear() {
	el.kids = []*Lmnt{}
}

func (el *Lmnt) WalkDown(fn func(*Lmnt)) {
	fn(el)
	for _, ll := range el.kids {
		ll.WalkDown(fn)
	}
}

func (el *Lmnt) WalkUp(fn func(*Lmnt)) {
	for _, ll := range el.kids {
		ll.WalkUp(fn)
	}
	fn(el)
}

// ### The Mechanism
func setTotals(el *Lmnt) {
	el.total.w, el.total.h = 0, 0
	for _, ll := range el.kids {
		mw, mh := myMax(ll.w, ll.total.w), myMax(ll.h, ll.total.h)
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
		if done {
			break
		}
	}
	for i := range sl {
		if sl[i] == 0 {
			sl[i] = delta
		}
	}
	return sl
}

func (el *Lmnt) setClm() {
	fs := myMax(el.p2.y - el.p1.y, el.total.h)
	n := float32(len(el.kids))
	sl := make([]float32, len(el.kids))
	st := make([]float32, len(el.kids))
	for i, v := range el.kids {
		if v.h > 0 {
			sl[i] = v.h
			fs -= v.h
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
		v.p2.x = v.p1.x + cmp.Or(v.w, w)
		v.p1.y = dy
		dy += hl[i]
		v.p2.y = dy
	}
}

func (el *Lmnt) setRow() {
	fs := myMax(el.p2.x - el.p1.x, el.total.w)
	n := float32(len(el.kids))
	sl := make([]float32, len(el.kids))
	st := make([]float32, len(el.kids))
	for i, v := range el.kids {
		if v.w > 0 {
			sl[i] = v.w
			fs -= v.w
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
		v.p2.y = v.p1.y + cmp.Or(v.h, h)

		v.p1.x = dx
		dx += wl[i]
		v.p2.x = dx
	}
}

func setRects(el *Lmnt) {
	if el.kids == nil { return }
	if el.row { el.setRow() } else { el.setClm() }
}

func (el *Lmnt) DoAll() {
	el.WalkUp(setTotals)
	el.WalkDown(setRects)
}
