package mergesort

func MergeSort(array []float64) []float64 {
	if len(array) <= 1 {
		_c := make([]float64, len(array))
		copy(_c, array)
		return _c
	}
	_f := len(array) / 2
	_e := MergeSort(array[:_f])
	_gc := MergeSort(array[_f:])
	_ec := make([]float64, len(array))
	_fc := 0
	_cb := 0
	_fb := 0
	for _cb < len(_e) && _fb < len(_gc) {
		if _e[_cb] <= _gc[_fb] {
			_ec[_fc] = _e[_cb]
			_cb++
		} else {
			_ec[_fc] = _gc[_fb]
			_fb++
		}
		_fc++
	}
	for _cb < len(_e) {
		_ec[_fc] = _e[_cb]
		_cb++
		_fc++
	}
	for _fb < len(_gc) {
		_ec[_fc] = _gc[_fb]
		_fb++
		_fc++
	}
	return _ec
}
