package formatutils

import (
	_b "fmt"
	_g "github.com/arcpd/msword/schema/soo/wml"
	_be "strconv"
	_a "strings"
)

func _ge(_bg string) (_ee []string) {
	for _df := 0; _df < len(_bg)-2; _df++ {
		if string(_bg[_df]) == "\u0025" {
			if !_a.Contains(string(_bg[_df+2:]), "\u0025") {
				if _df == 0 {
					_ee = append(_ee, _b.Sprintf("\u0025\u0073\u0025\u0073\u0025\u0073", string(_bg[_df]), string(_bg[_df+1]), string(_bg[_df+2:])))
				} else {
					_ee = append(_ee, _b.Sprintf("\u0025\u0073\u0025\u0073\u0025\u0073\u0025\u0073", string(_bg[_df-1]), string(_bg[_df]), string(_bg[_df+1]), string(_bg[_df+2:])))
				}
			} else {
				_ee = append(_ee, _b.Sprintf("\u0025\u0073\u0025\u0073\u0025\u0073", string(_bg[_df]), string(_bg[_df+1]), string(_bg[_df+2])))
			}
		}
	}
	return
}

func _ebd(_fbf int64, _fag *_g.CT_NumFmt) (_bb string) {
	if _fag == nil {
		return
	}
	_fc := _fag.ValAttr
	switch _fc {
	case _g.ST_NumberFormatNone:
		_bb = ""
	case _g.ST_NumberFormatDecimal:
		_bb = _be.Itoa(int(_fbf))
	case _g.ST_NumberFormatDecimalZero:
		_bb = _be.Itoa(int(_fbf))
		if _fbf < 10 {
			_bb = "\u0030" + _bb
		}
	case _g.ST_NumberFormatUpperRoman:
		var (
			_gec = _fbf % 10
			_dfd = (_fbf % 100) / 10
			_fae = (_fbf % 1000) / 100
			_bcb = _fbf / 1000
		)
		_bb = _bc[_bcb] + _eg[_fae] + _c[_dfd] + _dgg[_gec]
	case _g.ST_NumberFormatLowerRoman:
		var (
			_bbg  = _fbf % 10
			_ga   = (_fbf % 100) / 10
			_bgb  = (_fbf % 1000) / 100
			_ebdd = _fbf / 1000
		)
		_bb = _bc[_ebdd] + _eg[_bgb] + _c[_ga] + _dgg[_bbg]
		_bb = _a.ToLower(_bb)
	case _g.ST_NumberFormatUpperLetter:
		_dfa := _fbf % 780
		if _dfa == 0 {
			_dfa = 780
		}
		_ef := (_dfa - 1) / 26
		_gdb := (_dfa - 1) % 26
		_gca := _cf[_ef+_gdb]
		_bb = string(_gca)
	case _g.ST_NumberFormatLowerLetter:
		_egc := _fbf % 780
		if _egc == 0 {
			_egc = 780
		}
		_bef := (_egc - 1) / 26
		_ab := (_egc - 1) % 26
		_dd := _cf[_bef+_ab]
		_bb = _a.ToLower(string(_dd))
	default:
		_bb = ""
	}
	return
}

func FormatNumberingText(currentNumber int64, ilvl int64, lvlText string, numFmt *_g.CT_NumFmt, levelNumbers map[int64]int64) string {
	_gb := _ge(lvlText)
	_d := _ebd(currentNumber, numFmt)
	_fb := int64(0)
	for _e, _dg := range _gb {
		_ff := _b.Sprintf("\u0025\u0025\u0025\u0064", _e+1)
		if len(_gb) == 1 {
			_ff = _b.Sprintf("\u0025\u0025\u0025\u0064", ilvl+1)
			_gb[_e] = _a.Replace(_dg, _ff, _d, 1)
			break
		}
		if ilvl > 0 && ilvl > _fb && _e < (len(_gb)-1) {
			_gc := _ebd(levelNumbers[_fb], numFmt)
			_gb[_e] = _a.Replace(_dg, _ff, _gc, 1)
			_fb++
		} else {
			_gb[_e] = _a.Replace(_dg, _ff, _d, 1)
		}
	}
	return _a.Join(_gb, "")
}

func StringToNumbers(str string) (int, bool) {
	_db := 0
	_fcc := false
	for _, _ffe := range []byte(str) {
		_ffe -= '0'
		if _ffe > 9 {
			continue
		}
		_db = _db*10 + int(_ffe)
		_fcc = true
	}
	return _db, _fcc
}

var (
	_dgg = []string{"", "\u0049", "\u0049\u0049", "\u0049\u0049\u0049", "\u0049\u0056", "\u0056", "\u0056\u0049", "\u0056\u0049\u0049", "\u0056\u0049\u0049\u0049", "\u0049\u0058"}
	_c   = []string{"", "\u0058", "\u0058\u0058", "\u0058\u0058\u0058", "\u0058\u004c", "\u004c", "\u004c\u0058", "\u004c\u0058\u0058", "\u004c\u0058\u0058\u0058", "\u0058\u0043"}
	_eg  = []string{"", "\u0043", "\u0043\u0043", "\u0043\u0043\u0043", "\u0043\u0044", "\u0044", "\u0044\u0043", "\u0044\u0043\u0043", "\u0044\u0043\u0043\u0043", "\u0043\u004d", "\u004d"}
	_bc  = []string{"", "\u004d", "\u004d\u004d", "\u004d\u004d\u004d", "\u004d\u004d\u004d\u004d", "\u004d\u004d\u004dM\u004d", "\u004d\u004d\u004d\u004d\u004d\u004d", "\u004dM\u004d\u004d\u004d\u004d\u004d", "\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d", "\u004dM\u004d\u004d\u004d\u004d\u004d\u004dM", "\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d"}
	_bgc = []string{"\u006f\u006e\u0065", "\u0074\u0077\u006f", "\u0074\u0068\u0072e\u0065", "\u0066\u006f\u0075\u0072", "\u0066\u0069\u0076\u0065", "\u0073\u0069\u0078", "\u0073\u0065\u0076e\u006e", "\u0065\u0069\u0067h\u0074", "\u006e\u0069\u006e\u0065", "\u0074\u0065\u006e", "\u0065\u006c\u0065\u0076\u0065\u006e", "\u0074\u0077\u0065\u006c\u0076\u0065", "\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e", "\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e", "\u0066i\u0066\u0074\u0065\u0065\u006e", "\u0073i\u0078\u0074\u0065\u0065\u006e", "\u0073e\u0076\u0065\u006e\u0074\u0065\u0065n", "\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e", "\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e"}
	_fa  = []string{"\u0074\u0065\u006e", "\u0074\u0077\u0065\u006e\u0074\u0079", "\u0074\u0068\u0069\u0072\u0074\u0079", "\u0066\u006f\u0072t\u0079", "\u0066\u0069\u0066t\u0079", "\u0073\u0069\u0078t\u0079", "\u0073e\u0076\u0065\u006e\u0074\u0079", "\u0065\u0069\u0067\u0068\u0074\u0079", "\u006e\u0069\u006e\u0065\u0074\u0079"}
	_eb  = []string{"\u0066\u0069\u0072s\u0074", "\u0073\u0065\u0063\u006f\u006e\u0064", "\u0074\u0068\u0069r\u0064", "\u0066\u006f\u0075\u0072\u0074\u0068", "\u0066\u0069\u0066t\u0068", "\u0073\u0069\u0078t\u0068", "\u0073e\u0076\u0065\u006e\u0074\u0068", "\u0065\u0069\u0067\u0068\u0074\u0068", "\u006e\u0069\u006et\u0068", "\u0074\u0065\u006et\u0068", "\u0065\u006c\u0065\u0076\u0065\u006e\u0074\u0068", "\u0074w\u0065\u006c\u0066\u0074\u0068", "\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e\u0074\u0068", "\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e\u0074\u0068", "\u0066i\u0066\u0074\u0065\u0065\u006e\u0074h", "\u0073i\u0078\u0074\u0065\u0065\u006e\u0074h", "s\u0065\u0076\u0065\u006e\u0074\u0065\u0065\u006e\u0074\u0068", "\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e\u0074\u0068", "\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e\u0074\u0068"}
	_gd  = []string{"\u0074\u0065\u006et\u0068", "\u0074w\u0065\u006e\u0074\u0069\u0065\u0074h", "\u0074h\u0069\u0072\u0074\u0069\u0065\u0074h", "\u0066\u006f\u0072\u0074\u0069\u0065\u0074\u0068", "\u0066\u0069\u0066\u0074\u0069\u0065\u0074\u0068", "\u0073\u0069\u0078\u0074\u0069\u0065\u0074\u0068", "\u0073\u0065\u0076\u0065\u006e\u0074\u0069\u0065\u0074\u0068", "\u0065i\u0067\u0068\u0074\u0069\u0065\u0074h", "\u006ei\u006e\u0065\u0074\u0069\u0065\u0074h"}
	_cf  = "\u0041\u0042\u0043\u0044\u0045\u0046\u0047\u0048\u0049\u004a\u004bL\u004d\u004e\u004f\u0050\u0051\u0052\u0053\u0054\u0055\u0056W\u0058\u0059\u005a"
)
