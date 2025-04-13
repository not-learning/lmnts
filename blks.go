package lmnts

// ### Zero (or not set) size is autosize.
// TODO Gaps and friends arrange their kids?

func (el *Lmnt) GapsBetween(size float32, lls ...*Lmnt) {
	for i, k := range lls {
		el.Add(k)
		if i == len(lls)-1 { continue }
		gap := New()
		//gap.name = "gap"
		el.Add(gap)
		if el.row { gap.w = size } else { gap.h = size }
	}
}//*/

func (el *Lmnt) GapsAround(size float32, lls ...*Lmnt) {
	gap1, gap2 := New(), New()
	if el.row {
		gap1.w, gap2.w = size, size
	} else {
		gap1.h, gap2.h = size, size
	}
	el.Add(gap1)
	el.GapsBetween(size, lls...)
	el.Add(gap2)
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
	pad.w = size
	if el.row {
		el.Add(pad, ll)
	} else {
		inL := New()
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
	padL.w, padR.w = lS, rS
	if el.row {
		el.Add(padL, ll, padR)
	} else {
		inL := New()
		inL.Add(padL, ll, padR)
		el.Add(inL)
	}
}

func (el *Lmnt) AddTBLR(tSize, bSize, lSize, rSize float32, ll *Lmnt) {
	t, b, c, l, r := New(), New(), New(), New(), New()
	//t.Name, b.Name, c.Name, l.Name, r.Name = "t", "b", "c", "l", "r"
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
