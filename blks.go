package lmnts

// ### Zero (or not set) size is autosize.
// TODO Gaps and friends arrange their kids?
// TODO think: upon Adding any Pad, Size (if smaller) makes no sense.
// TODO make tests

import "math"

func (el *Lmnt) GapsBetween(size float32, lls ...*Lmnt) {
	for i, k := range lls {
		el.Add(k)
		if i == len(lls)-1 { continue }
		gap := New()
		gap.Name = "bgap"
		el.Add(gap)
		if el.row { gap.w = size } else { gap.h = size }
	}
}//*/

// todo: test!
func (el *Lmnt) GapsAround(size float32, lls ...*Lmnt) {
	gap1, gap2 := New(), New()
	gap1.Name, gap2.Name = "agap", "agap"
	if el.row {
		gap1.w, gap2.w = size, size
	} else {
		gap1.h, gap2.h = size, size
	}
	el.Add(gap1)
	el.GapsBetween(size, lls...)
	el.Add(gap2)
}

// todo 0 size for auto-arrange?
// todo center last row?
func (el *Lmnt) Grid(num int, gap float32, lls ...*Lmnt) {
	if len(el.kids) != 0 { return }
	n := len(lls) / num
	if n*num < len(lls) { n++ }

	in := make([]*Lmnt, n)
	for i := range n {
		in[i] = New()
		in[i].Name = "grid"
		if !el.row { in[i].SetRow() }
		a := int(math.Min(float64((i+1)*num), float64(len(lls))))
		in[i].GapsBetween(gap, lls[i*num:a]...)
	}
	el.GapsBetween(gap, in...)
}

func (el *Lmnt) AddT(size float32, ll *Lmnt) {
	pad := New()
	pad.h = size
	if !el.row {
		el.Add(pad, ll)
	} else {
		inL := New()
		inL.Add(pad, ll)
		el.Add(inL)
	}
}

func (el *Lmnt) AddB(size float32, ll *Lmnt) {
	pad := New()
	pad.h = size
	if !el.row {
		el.Add(ll, pad)
	} else {
		inL := New()
		inL.Add(ll, pad)
		el.Add(inL)
	}
}

func (el *Lmnt) AddL(size float32, ll *Lmnt) {
	pad := New()
	pad.Name = "padl"
	pad.w = size
	if el.row {
		el.Add(pad, ll)
	} else {
		inL := New()
		inL.SetRow()
		inL.Add(pad, ll)
		el.Add(inL)
	}
}

func (el *Lmnt) AddR(size float32, ll *Lmnt) {
	pad := New()
	pad.w = size
	if el.row {
		el.Add(ll, pad)
	} else {
		inL := New()
		inL.SetRow()
		inL.Add(ll, pad)
		el.Add(inL)
	}
}

func (el *Lmnt) AddTB(tS, bS float32, ll *Lmnt) {
	padT, padB := New(), New()
	padT.h, padB.h = tS, bS
	if !el.row {
		el.Add(padT, ll, padB)
	} else {
		inL := New()
		inL.Add(padT, ll, padB)
		el.Add(inL)
	}
}

func (el *Lmnt) AddLR(lS, rS float32, ll *Lmnt) {
	padL, padR := New(), New()
	padL.Name, padR.Name = "lr", "lr"
	padL.w, padR.w = lS, rS
	if el.row {
		el.Add(padL, ll, padR)
	} else {
		inL := New()
		inL.SetRow()
		inL.Add(padL, ll, padR)
		el.Add(inL)
	}
}

func (el *Lmnt) AddTBLR(tSize, bSize, lSize, rSize float32, ll *Lmnt) {
	t, b, c, l, r := New(), New(), New(), New(), New()
	t.Name, b.Name, c.Name, l.Name, r.Name = "t", "b", "c", "l", "r"
	t.h, b.h, l.w, r.w = tSize, bSize, lSize, rSize
	if !el.row {
		el.Add(t, c, b)
		c.SetRow()
		c.Add(l, ll, r)
	} else {
		el.Add(l, c, r)
		c.Add(t, ll, b)
	}
}//*/

//func (el *Lmnt) ScrollList(gap, lls ...*Lmnt) {}
