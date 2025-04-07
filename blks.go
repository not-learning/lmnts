package lmnts

// ### Zero (or unset) size is autosize.

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
