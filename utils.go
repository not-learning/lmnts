package lmnts

import "cmp"

func myMax[T cmp.Ordered](nn ...T) T {
	mx := nn[0]
	for _, v := range nn {
		if mx < v {
			mx = v
		}
	}
	return mx
}

func myMin[T cmp.Ordered](nn ...T) T {
	mx := nn[0]
	for _, v := range nn {
		if mx > v {
			mx = v
		}
	}
	return mx
}

func (el *Lmnt) MidF32() (mx, my float32) {
	x1, y1, x2, y2 := el.Rect()
	mx = x1 + (x2-x1)/2
	my = y1 + (y2-y1)/2
	return
}

func (el *Lmnt) MidF64() (float64, float64) {
	mx, my := el.MidF32()
	return float64(mx), float64(my)
}

func (el *Lmnt) MidInt() (int, int) {
	mx, my := el.MidF32()
	return int(mx), int(my)
}
