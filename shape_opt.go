package arrgo

func (a *Arrf) SameShapeTo(b *Arrf) bool {
	return SameIntSlice(a.shape, b.shape)
}

func Vstack(arrs ...*Arrf) *Arrf {
	for i := range arrs {
		if arrs[i].Ndims() > 2 {
			panic(SHAPE_ERROR)
		}
	}
	if len(arrs) == 0 {
		return nil
	}
	if len(arrs) == 1 {
		return arrs[0].Copy()
	}

	var vlenSum int = 0

	var hlen int
	if arrs[0].Ndims() == 1 {
		hlen = arrs[0].shape[0]
		vlenSum += 1
	} else {
		hlen = arrs[0].shape[1]
		vlenSum += arrs[0].shape[0]
	}
	for i := 1; i < len(arrs); i++ {
		var nextHen int
		if arrs[i].Ndims() == 1 {
			nextHen = arrs[i].shape[0]
			vlenSum += 1
		} else {
			nextHen = arrs[i].shape[1]
			vlenSum += arrs[i].shape[0]
		}
		if hlen != nextHen {
			panic(SHAPE_ERROR)
		}
	}

	data := make([]float64, vlenSum * hlen)
	var offset = 0
	for i := range arrs {
		copy(data[offset:], arrs[i].data)
		offset += len(arrs[i].data)
	}

	return Array(data, vlenSum, hlen)
}

func Hstack(arrs ...*Arrf) *Arrf {
	for i := range arrs {
		if arrs[i].Ndims() > 2 {
			panic(SHAPE_ERROR)
		}
	}
	if len(arrs) == 0 {
		return nil
	}
	if len(arrs) == 1 {
		return arrs[0].Copy()
	}

	var hlenSum int = 0
	var hBlockLens = make([]int, len(arrs))
	var vlen int
	if arrs[0].Ndims() == 1 {
		vlen = 1
		hlenSum += arrs[0].shape[0]
		hBlockLens[0] = arrs[0].shape[0]
	} else {
		vlen = arrs[0].shape[0]
		hlenSum += arrs[0].shape[1]
		hBlockLens[0] = arrs[0].shape[1]
	}
	for i := 1; i < len(arrs); i++ {
		var nextVlen int
		if arrs[i].Ndims() == 1 {
			nextVlen = 1
			hlenSum += arrs[i].shape[0]
			hBlockLens[i] = arrs[i].shape[0]
		} else {
			nextVlen = arrs[i].shape[0]
			hlenSum += arrs[i].shape[1]
			hBlockLens[i] = arrs[i].shape[1]
		}
		if vlen != nextVlen {
			panic(SHAPE_ERROR)
		}
	}

	data := make([]float64, hlenSum*vlen)
	for i := 0; i < vlen; i++ {
		var curPos = 0
		for j := 0; j < len(arrs); j++ {
			copy(data[curPos+i*hlenSum: curPos+i*hlenSum+hBlockLens[j]], arrs[j].data[i*hBlockLens[j]:(i+1)*hBlockLens[j]])
			curPos += hBlockLens[j]
		}
	}

	return Array(data, vlen, hlenSum)
}

