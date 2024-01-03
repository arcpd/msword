package algo

import _g "strconv"

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess(lhs, rhs string) bool {
	_a, _b := 0, 0
	for _a < len(lhs) && _b < len(rhs) {
		_gf := lhs[_a]
		_bb := rhs[_b]
		_gaf := _ga(_gf)
		_d := _ga(_bb)
		switch {
		case _gaf && !_d:
			return true
		case !_gaf && _d:
			return false
		case !_gaf && !_d:
			if _gf != _bb {
				return _gf < _bb
			}
			_a++
			_b++
		default:
			_cf := _a + 1
			_eb := _b + 1
			for _cf < len(lhs) && _ga(lhs[_cf]) {
				_cf++
			}
			for _eb < len(rhs) && _ga(rhs[_eb]) {
				_eb++
			}
			_gc, _ := _g.ParseUint(lhs[_a:_cf], 10, 64)
			_bf, _ := _g.ParseUint(rhs[_a:_eb], 10, 64)
			if _gc != _bf {
				return _gc < _bf
			}
			_a = _cf
			_b = _eb
		}
	}
	return len(lhs) < len(rhs)
}
func RepeatString(s string, cnt int) string {
	if cnt <= 0 {
		return ""
	}
	_aa := make([]byte, len(s)*cnt)
	_aae := []byte(s)
	for _bfa := 0; _bfa < cnt; _bfa++ {
		copy(_aa[_bfa:], _aae)
	}
	return string(_aa)
}

func _ga(_e byte) bool { return _e >= '0' && _e <= '9' }
