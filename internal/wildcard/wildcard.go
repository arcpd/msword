package wildcard

func Match(pattern, name string) (_df bool) {
	if pattern == "" {
		return name == pattern
	}
	if pattern == "\u002a" {
		return true
	}
	_cc := make([]rune, 0, len(name))
	_cf := make([]rune, 0, len(pattern))
	for _, _ae := range name {
		_cc = append(_cc, _ae)
	}
	for _, _ef := range pattern {
		_cf = append(_cf, _ef)
	}
	_be := false
	return _f(_cc, _cf, _be)
}

func Index(pattern, name string) (_aa int) {
	if pattern == "" || pattern == "\u002a" {
		return 0
	}
	_dc := make([]rune, 0, len(name))
	_bb := make([]rune, 0, len(pattern))
	for _, _ec := range name {
		_dc = append(_dc, _ec)
	}
	for _, _gg := range pattern {
		_bb = append(_bb, _gg)
	}
	return _ff(_dc, _bb, 0)
}

func _f(_cag, _da []rune, _gc bool) bool {
	for len(_da) > 0 {
		switch _da[0] {
		default:
			if len(_cag) == 0 || _cag[0] != _da[0] {
				return false
			}
		case '?':
			if len(_cag) == 0 && !_gc {
				return false
			}
		case '*':
			return _f(_cag, _da[1:], _gc) || (len(_cag) > 0 && _f(_cag[1:], _da, _gc))
		}
		_cag = _cag[1:]
		_da = _da[1:]
	}
	return len(_cag) == 0 && len(_da) == 0
}

func _ff(_db, _bf []rune, _eg int) int {
	for len(_bf) > 0 {
		switch _bf[0] {
		default:
			if len(_db) == 0 {
				return -1
			}
			if _db[0] != _bf[0] {
				return _ff(_db[1:], _bf, _eg+1)
			}
		case '?':
			if len(_db) == 0 {
				return -1
			}
		case '*':
			if len(_db) == 0 {
				return -1
			}
			_ee := _ff(_db, _bf[1:], _eg)
			if _ee != -1 {
				return _eg
			} else {
				_ee = _ff(_db[1:], _bf, _eg)
				if _ee != -1 {
					return _eg
				} else {
					return -1
				}
			}
		}
		_db = _db[1:]
		_bf = _bf[1:]
	}
	return _eg
}

func MatchSimple(pattern, name string) bool {
	if pattern == "" {
		return name == pattern
	}
	if pattern == "\u002a" {
		return true
	}
	_b := make([]rune, 0, len(name))
	_e := make([]rune, 0, len(pattern))
	for _, _ab := range name {
		_b = append(_b, _ab)
	}
	for _, _d := range pattern {
		_e = append(_e, _d)
	}
	_bc := true
	return _f(_b, _e, _bc)
}
