package format

import (
	_gd "bytes"
	_a "fmt"
	_cg "github.com/arcpd/msword/common/logger"
	_e "io"
	_g "math"
	_cf "strconv"
	_b "strings"
	_f "time"
)

// String returns the string formatted according to the type.  In format strings
// this is the fourth item, where '@' is used as a placeholder for text.
func String(v string, f string) string {
	_ce := Parse(f)
	var _bb Format
	if len(_ce) == 1 {
		_bb = _ce[0]
	} else if len(_ce) == 4 {
		_bb = _ce[3]
	}
	_afc := false
	for _, _ed := range _bb.Whole {
		if _ed.Type == FmtTypeText {
			_afc = true
		}
	}
	if !_afc {
		return v
	}
	_aee := _gd.Buffer{}
	for _, _ea := range _bb.Whole {
		switch _ea.Type {
		case FmtTypeLiteral:
			_aee.WriteByte(_ea.Literal)
		case FmtTypeText:
			_aee.WriteString(v)
		}
	}
	return _aee.String()
}
func (_eg FmtType) String() string {
	if _eg >= FmtType(len(_d)-1) {
		return _a.Sprintf("F\u006d\u0074\u0054\u0079\u0070\u0065\u0028\u0025\u0064\u0029", _eg)
	}
	return _ac[_d[_eg]:_d[_eg+1]]
}

const _ae = 1e-10

func _dg(_gff float64, _fa Format, _gac bool) string {
	if _fa._ca {
		return NumberGeneric(_gff)
	}
	_cab := make([]byte, 0, 20)
	_caf := _g.Signbit(_gff)
	_dcd := _g.Abs(_gff)
	_de := int64(0)
	_fga := int64(0)
	if _fa.IsExponential {
		for _dcd >= 10 {
			_fga++
			_dcd /= 10
		}
		for _dcd < 1 {
			_fga--
			_dcd *= 10
		}
	} else if _fa._acc {
		_dcd *= 100
	} else if _fa._bf {
		if _fa._bg == 0 {
			_ded := _g.Pow(10, float64(_fa._cfg))
			_deg, _ff := 1.0, 1.0
			_ = _deg
			for _caa := 1.0; _caa < _ded; _caa++ {
				_, _ge := _g.Modf(_dcd * float64(_caa))
				if _ge < _ff {
					_ff = _ge
					_deg = _caa
					if _ge == 0 {
						break
					}
				}
			}
			_fa._bg = int64(_deg)
		}
		_de = int64(_dcd*float64(_fa._bg) + 0.5)
		if len(_fa.Whole) > 0 && _de > _fa._bg {
			_de = int64(_dcd*float64(_fa._bg)) % _fa._bg
			_dcd -= float64(_de) / float64(_fa._bg)
		} else {
			_dcd -= float64(_de) / float64(_fa._bg)
			if _g.Abs(_dcd) < 1 {
				_ead := true
				for _, _gfc := range _fa.Whole {
					if _gfc.Type == FmtTypeDigitOpt {
						continue
					}
					if _gfc.Type == FmtTypeLiteral && _gfc.Literal == ' ' {
						continue
					}
					_ead = false
				}
				if _ead {
					_fa.Whole = nil
				}
			}
		}
	}
	_ddc := 1
	for _, _bgc := range _fa.Fractional {
		if _bgc.Type == FmtTypeDigit || _bgc.Type == FmtTypeDigitOpt {
			_ddc++
		}
	}
	_dcd += 5 * _g.Pow10(-_ddc)
	_ad, _cad := _g.Modf(_dcd)
	_cab = append(_cab, _bbc(_ad, _gff, _fa)...)
	_cab = append(_cab, _bgf(_cad, _gff, _fa)...)
	_cab = append(_cab, _ccg(_fga, _fa)...)
	if _fa._bf {
		_cab = _cf.AppendInt(_cab, _de, 10)
		_cab = append(_cab, '/')
		_cab = _cf.AppendInt(_cab, _fa._bg, 10)
	}
	if !_gac && _caf {
		return "\u002d" + string(_cab)
	}
	return string(_cab)
}

// Token is a format token in the Excel format string.
type Token struct {
	Type     FmtType
	Literal  byte
	DateTime string
}

func _bbc(_ab, _ba float64, _bca Format) []byte {
	if len(_bca.Whole) == 0 {
		return nil
	}
	_cc := _f.Date(1899, 12, 30, 0, 0, 0, 0, _f.UTC)
	_bcf := _cc.Add(_f.Duration(_ba * float64(24*_f.Hour)))
	_bcf = _faf(_bcf)
	_fac := _cf.AppendFloat(nil, _ab, 'f', -1, 64)
	_faca := make([]byte, 0, len(_fac))
	_gb := 0
	_eb := 1
_egc:
	for _gfd := len(_bca.Whole) - 1; _gfd >= 0; _gfd-- {
		_gbe := len(_fac) - 1 - _gb
		_gdd := _bca.Whole[_gfd]
		switch _gdd.Type {
		case FmtTypeDigit:
			if _gbe >= 0 {
				_faca = append(_faca, _fac[_gbe])
				_gb++
				_eb = _gfd
			} else {
				_faca = append(_faca, '0')
			}
		case FmtTypeDigitOpt:
			if _gbe >= 0 {
				_faca = append(_faca, _fac[_gbe])
				_gb++
				_eb = _gfd
			} else {
				for _fc := _gfd; _fc >= 0; _fc-- {
					_cfd := _bca.Whole[_fc]
					if _cfd.Type == FmtTypeLiteral {
						_faca = append(_faca, _cfd.Literal)
					}
				}
				break _egc
			}
		case FmtTypeDollar:
			for _dbf := _gb; _dbf < len(_fac); _dbf++ {
				_faca = append(_faca, _fac[len(_fac)-1-_dbf])
				_gb++
			}
			_faca = append(_faca, '$')
		case FmtTypeComma:
			if !_bca._cgg {
				_faca = append(_faca, ',')
			}
		case FmtTypeLiteral:
			_faca = append(_faca, _gdd.Literal)
		case FmtTypeDate:
			_faca = append(_faca, _ga(_ggd(_bcf, _gdd.DateTime))...)
		case FmtTypeTime:
			_faca = append(_faca, _ga(_dcc(_bcf, _ba, _gdd.DateTime))...)
		default:
			_cg.Log.Debug("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0077\u0068\u006f\u006c\u0065 \u0025\u0076", _gdd)
		}
	}
	_dbc := _ga(_faca)
	if _gb < len(_fac) && (_gb != 0 || _bca._af) {
		_fae := len(_fac) - _gb
		_geg := make([]byte, len(_dbc)+_fae)
		copy(_geg, _dbc[0:_eb])
		copy(_geg[_eb:], _fac[0:])
		copy(_geg[_eb+_fae:], _dbc[_eb:])
		_dbc = _geg
	}
	if _bca._cgg {
		_egg := _gd.Buffer{}
		_dde := 0
		for _abf := len(_dbc) - 1; _abf >= 0; _abf-- {
			if !(_dbc[_abf] >= '0' && _dbc[_abf] <= '9') {
				_dde++
			} else {
				break
			}
		}
		for _gegf := 0; _gegf < len(_dbc); _gegf++ {
			_dda := (len(_dbc) - _gegf - _dde)
			if _dda%3 == 0 && _dda != 0 && _gegf != 0 {
				_egg.WriteByte(',')
			}
			_egg.WriteByte(_dbc[_gegf])
		}
		_dbc = _egg.Bytes()
	}
	return _dbc
}
func _faf(_dafd _f.Time) _f.Time {
	_dafd = _dafd.UTC()
	return _f.Date(_dafd.Year(), _dafd.Month(), _dafd.Day(), _dafd.Hour(), _dafd.Minute(), _dafd.Second(), _dafd.Nanosecond(), _f.Local)
}
func _faa(_bgd float64) string {
	_gbbe := _cf.FormatFloat(_bgd, 'E', -1, 64)
	_dgg := _cf.FormatFloat(_bgd, 'E', 5, 64)
	if len(_gbbe) < len(_dgg) {
		return _cf.FormatFloat(_bgd, 'E', 2, 64)
	}
	return _dgg
}
func _ggd(_bcb _f.Time, _acf string) []byte {
	_eba := []byte{}
	_bea := 0
	for _gegb := 0; _gegb < len(_acf); _gegb++ {
		var _agf string
		if _acf[_gegb] == '/' {
			_agf = string(_acf[_bea:_gegb])
			_bea = _gegb + 1
		} else if _gegb == len(_acf)-1 {
			_agf = string(_acf[_bea : _gegb+1])
		} else {
			continue
		}
		switch _agf {
		case "\u0079\u0079":
			_eba = _bcb.AppendFormat(_eba, "\u0030\u0036")
		case "\u0079\u0079\u0079\u0079":
			_eba = _bcb.AppendFormat(_eba, "\u0032\u0030\u0030\u0036")
		case "\u006d":
			_eba = _bcb.AppendFormat(_eba, "\u0031")
		case "\u006d\u006d":
			_eba = _bcb.AppendFormat(_eba, "\u0030\u0031")
		case "\u006d\u006d\u006d":
			_eba = _bcb.AppendFormat(_eba, "\u004a\u0061\u006e")
		case "\u006d\u006d\u006d\u006d":
			_eba = _bcb.AppendFormat(_eba, "\u004aa\u006e\u0075\u0061\u0072\u0079")
		case "\u006d\u006d\u006dm\u006d":
			switch _bcb.Month() {
			case _f.January, _f.July, _f.June:
				_eba = append(_eba, 'J')
			case _f.February:
				_eba = append(_eba, 'M')
			case _f.March, _f.May:
				_eba = append(_eba, 'M')
			case _f.April, _f.August:
				_eba = append(_eba, 'A')
			case _f.September:
				_eba = append(_eba, 'S')
			case _f.October:
				_eba = append(_eba, 'O')
			case _f.November:
				_eba = append(_eba, 'N')
			case _f.December:
				_eba = append(_eba, 'D')
			}
		case "\u0064":
			_eba = _bcb.AppendFormat(_eba, "\u0032")
		case "\u0064\u0064":
			_eba = _bcb.AppendFormat(_eba, "\u0030\u0032")
		case "\u0064\u0064\u0064":
			_eba = _bcb.AppendFormat(_eba, "\u004d\u006f\u006e")
		case "\u0064\u0064\u0064\u0064":
			_eba = _bcb.AppendFormat(_eba, "\u004d\u006f\u006e\u0064\u0061\u0079")
		default:
			_cg.Log.Debug("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0064\u0061\u0074\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073", _agf)
		}
		if _acf[_gegb] == '/' {
			_eba = append(_eba, '/')
		}
	}
	return _eba
}
func IsNumber(data string) (_cca bool) {
	_bac, _bbf, _aff := 0, 0, len(data)
	_ef := len(data)
	_gebb, _cebg, _bee := 0, 0, 0
	_ = _cebg
	_ = _bee
	_ = _gebb
	{
		_bac = _cce
		_gebb = 0
		_cebg = 0
		_bee = 0
	}
	{
		if _bbf == _aff {
			goto _bfe
		}
		switch _bac {
		case 0:
			goto _ebad
		case 1:
			goto _befa
		case 2:
			goto _cfec
		case 3:
			goto _gfde
		case 4:
			goto _ecb
		case 5:
			goto _fd
		case 6:
			goto _gba
		case 7:
			goto _fba
		}
		goto _fgeg
	_dgb:
		_cebg = _bbf
		_bbf--
		{
			_cca = false
		}
		goto _gad
	_bef:
		_cebg = _bbf
		_bbf--
		{
			_cca = _cebg == len(data)
		}
		goto _gad
	_fb:
		_cebg = _bbf
		_bbf--
		{
			_cca = _cebg == len(data)
		}
		goto _gad
	_ebaa:
		switch _bee {
		case 2:
			{
				_bbf = (_cebg) - 1
				_cca = _cebg == len(data)
			}
		case 3:
			{
				_bbf = (_cebg) - 1
				_cca = false
			}
		}
		goto _gad
	_gad:
		_gebb = 0
		if _bbf++; _bbf == _aff {
			goto _ege
		}
	_ebad:
		_gebb = _bbf
		switch data[_bbf] {
		case 43:
			goto _fgd
		case 45:
			goto _fgd
		}
		if 48 <= data[_bbf] && data[_bbf] <= 57 {
			goto _ceg
		}
		goto _dbe
	_dbe:
		if _bbf++; _bbf == _aff {
			goto _ddcc
		}
	_befa:
		goto _dbe
	_fgd:
		if _bbf++; _bbf == _aff {
			goto _dgf
		}
	_cfec:
		if 48 <= data[_bbf] && data[_bbf] <= 57 {
			goto _ceg
		}
		goto _dbe
	_ceg:
		if _bbf++; _bbf == _aff {
			goto _ffb
		}
	_gfde:
		if data[_bbf] == 46 {
			goto _ec
		}
		if 48 <= data[_bbf] && data[_bbf] <= 57 {
			goto _ceg
		}
		goto _dbe
	_ec:
		if _bbf++; _bbf == _aff {
			goto _dcg
		}
	_ecb:
		if 48 <= data[_bbf] && data[_bbf] <= 57 {
			goto _ffc
		}
		goto _dbe
	_ffc:
		if _bbf++; _bbf == _aff {
			goto _cec
		}
	_fd:
		if data[_bbf] == 69 {
			goto _add
		}
		if 48 <= data[_bbf] && data[_bbf] <= 57 {
			goto _ffc
		}
		goto _dbe
	_add:
		if _bbf++; _bbf == _aff {
			goto _ddb
		}
	_gba:
		switch data[_bbf] {
		case 43:
			goto _gbbg
		case 45:
			goto _gbbg
		}
		goto _dbe
	_gbbg:
		_cebg = _bbf + 1
		_bee = 3
		goto _afd
	_afe:
		_cebg = _bbf + 1
		_bee = 2
		goto _afd
	_afd:
		if _bbf++; _bbf == _aff {
			goto _fdc
		}
	_fba:
		if 48 <= data[_bbf] && data[_bbf] <= 57 {
			goto _afe
		}
		goto _dbe
	_fgeg:
	_ege:
		_bac = 0
		goto _bfe
	_ddcc:
		_bac = 1
		goto _bfe
	_dgf:
		_bac = 2
		goto _bfe
	_ffb:
		_bac = 3
		goto _bfe
	_dcg:
		_bac = 4
		goto _bfe
	_cec:
		_bac = 5
		goto _bfe
	_ddb:
		_bac = 6
		goto _bfe
	_fdc:
		_bac = 7
		goto _bfe
	_bfe:
		{
		}
		if _bbf == _ef {
			switch _bac {
			case 1:
				goto _dgb
			case 2:
				goto _dgb
			case 3:
				goto _bef
			case 4:
				goto _dgb
			case 5:
				goto _fb
			case 6:
				goto _dgb
			case 7:
				goto _ebaa
			}
		}
	}
	if _bac == _eag {
		return false
	}
	return
}

const _dff int = 34

func (_gacb *Lexer) Lex(r _e.Reader) {
	_addg, _cff, _dcgc := 0, 0, 0
	_dggd := -1
	_dce, _dffb, _ddd := 0, 0, 0
	_ = _dffb
	_ = _ddd
	_cafe := 1
	_ = _cafe
	_bbca := make([]byte, 4096)
	_gdg := false
	for !_gdg {
		_agb := 0
		if _dce > 0 {
			_agb = _cff - _dce
		}
		_cff = 0
		_fgee, _cgd := r.Read(_bbca[_agb:])
		if _fgee == 0 || _cgd != nil {
			_gdg = true
		}
		_dcgc = _fgee + _agb
		if _dcgc < len(_bbca) {
			_dggd = _dcgc
		}
		{
			_addg = _dff
			_dce = 0
			_dffb = 0
			_ddd = 0
		}
		{
			if _cff == _dcgc {
				goto _bdcc
			}
			switch _addg {
			case 34:
				goto _dgfb
			case 35:
				goto _dbdd
			case 0:
				goto _bfd
			case 36:
				goto _adg
			case 37:
				goto _eef
			case 1:
				goto _gef
			case 2:
				goto _dfeg
			case 38:
				goto _bede
			case 3:
				goto _bedf
			case 4:
				goto _acg
			case 39:
				goto _dbfa
			case 5:
				goto _dddd
			case 6:
				goto _ade
			case 7:
				goto _aca
			case 8:
				goto _abg
			case 40:
				goto _agda
			case 9:
				goto _dgcf
			case 41:
				goto _ddcd
			case 10:
				goto _dbae
			case 42:
				goto _abfc
			case 11:
				goto _gca
			case 43:
				goto _cbea
			case 44:
				goto _baab
			case 45:
				goto _fgda
			case 12:
				goto _ceccb
			case 46:
				goto _dgfg
			case 13:
				goto _def
			case 14:
				goto _gcg
			case 15:
				goto _ggf
			case 16:
				goto _eea
			case 47:
				goto _ecbd
			case 17:
				goto _afg
			case 48:
				goto _gcga
			case 18:
				goto _aac
			case 19:
				goto _bedeb
			case 20:
				goto _gdf
			case 49:
				goto _bab
			case 50:
				goto _gdb
			case 21:
				goto _bcd
			case 22:
				goto _dbdf
			case 23:
				goto _aga
			case 24:
				goto _dga
			case 25:
				goto _fdcf
			case 51:
				goto _dcaa
			case 26:
				goto _bdd
			case 52:
				goto _egfe
			case 53:
				goto _dbeb
			case 54:
				goto _dccd
			case 55:
				goto _cabf
			case 56:
				goto _fdg
			case 57:
				goto _ebg
			case 27:
				goto _ecbe
			case 28:
				goto _cbfa
			case 29:
				goto _eaca
			case 30:
				goto _afb
			case 31:
				goto _gbcb
			case 58:
				goto _fadf
			case 32:
				goto _fgeb
			case 59:
				goto _aeac
			case 33:
				goto _ebaf
			case 60:
				goto _ffdd
			case 61:
				goto _eade
			case 62:
				goto _dbgd
			}
			goto _bgg
		_bfc:
			switch _ddd {
			case 2:
				{
					_cff = (_dffb) - 1
					_gacb._bga.AddToken(FmtTypeDigit, nil)
				}
			case 3:
				{
					_cff = (_dffb) - 1
					_gacb._bga.AddToken(FmtTypeDigitOpt, nil)
				}
			case 5:
				{
					_cff = (_dffb) - 1
				}
			case 8:
				{
					_cff = (_dffb) - 1
					_gacb._bga.AddToken(FmtTypePercent, nil)
				}
			case 13:
				{
					_cff = (_dffb) - 1
					_gacb._bga.AddToken(FmtTypeFraction, _bbca[_dce:_dffb])
				}
			case 14:
				{
					_cff = (_dffb) - 1
					_gacb._bga.AddToken(FmtTypeDate, _bbca[_dce:_dffb])
				}
			case 15:
				{
					_cff = (_dffb) - 1
					_gacb._bga.AddToken(FmtTypeTime, _bbca[_dce:_dffb])
				}
			case 16:
				{
					_cff = (_dffb) - 1
					_gacb._bga.AddToken(FmtTypeTime, _bbca[_dce:_dffb])
				}
			case 18:
				{
					_cff = (_dffb) - 1
				}
			case 20:
				{
					_cff = (_dffb) - 1
					_gacb._bga.AddToken(FmtTypeLiteral, _bbca[_dce:_dffb])
				}
			case 21:
				{
					_cff = (_dffb) - 1
					_gacb._bga.AddToken(FmtTypeLiteral, _bbca[_dce+1:_dffb-1])
				}
			}
			goto _bebb
		_fe:
			_cff = (_dffb) - 1
			{
				_gacb._bga.AddToken(FmtTypeFraction, _bbca[_dce:_dffb])
			}
			goto _bebb
		_bag:
			_cff = (_dffb) - 1
			{
				_gacb._bga.AddToken(FmtTypeDigitOpt, nil)
			}
			goto _bebb
		_geda:
			_dffb = _cff + 1
			{
				_gacb._bga.AddToken(FmtTypeDigitOptThousands, nil)
			}
			goto _bebb
		_caac:
			_cff = (_dffb) - 1
			{
				_gacb._bga.AddToken(FmtTypePercent, nil)
			}
			goto _bebb
		_eac:
			_cff = (_dffb) - 1
			{
				_gacb._bga.AddToken(FmtTypeDate, _bbca[_dce:_dffb])
			}
			goto _bebb
		_agd:
			_cff = (_dffb) - 1
			{
				_gacb._bga.AddToken(FmtTypeDigit, nil)
			}
			goto _bebb
		_fdcb:
			_cff = (_dffb) - 1
			{
				_gacb._bga.AddToken(FmtTypeTime, _bbca[_dce:_dffb])
			}
			goto _bebb
		_ced:
			_cff = (_dffb) - 1
			{
				_gacb._bga.AddToken(FmtTypeLiteral, _bbca[_dce:_dffb])
			}
			goto _bebb
		_bfg:
			_dffb = _cff + 1
			{
				_gacb._bga._ca = true
			}
			goto _bebb
		_cgdb:
			_dffb = _cff + 1
			{
				_gacb._bga.AddToken(FmtTypeLiteral, _bbca[_dce:_dffb])
			}
			goto _bebb
		_faag:
			_dffb = _cff + 1
			{
				_gacb._bga.AddToken(FmtTypeDollar, nil)
			}
			goto _bebb
		_egce:
			_dffb = _cff + 1
			{
				_gacb._bga.AddToken(FmtTypeComma, nil)
			}
			goto _bebb
		_egeg:
			_dffb = _cff + 1
			{
				_gacb._bga.AddToken(FmtTypeDecimal, nil)
			}
			goto _bebb
		_dbfb:
			_dffb = _cff + 1
			{
				_gacb.nextFmt()
			}
			goto _bebb
		_bdg:
			_dffb = _cff + 1
			{
				_gacb._bga.AddToken(FmtTypeText, nil)
			}
			goto _bebb
		_bdf:
			_dffb = _cff + 1
			{
				_gacb._bga.AddToken(FmtTypeUnderscore, nil)
			}
			goto _bebb
		_efb:
			_dffb = _cff
			_cff--
			{
				_gacb._bga.AddToken(FmtTypeLiteral, _bbca[_dce:_dffb])
			}
			goto _bebb
		_aa:
			_dffb = _cff
			_cff--
			{
				_gacb._bga.AddToken(FmtTypeLiteral, _bbca[_dce+1:_dffb-1])
			}
			goto _bebb
		_egb:
			_dffb = _cff
			_cff--
			{
				_gacb._bga.AddToken(FmtTypeDigitOpt, nil)
			}
			goto _bebb
		_ecaf:
			_dffb = _cff
			_cff--
			{
				_gacb._bga.AddToken(FmtTypeFraction, _bbca[_dce:_dffb])
			}
			goto _bebb
		_abc:
			_dffb = _cff
			_cff--
			{
				_gacb._bga.AddToken(FmtTypePercent, nil)
			}
			goto _bebb
		_dab:
			_dffb = _cff
			_cff--
			{
				_gacb._bga.AddToken(FmtTypeDate, _bbca[_dce:_dffb])
			}
			goto _bebb
		_gacf:
			_dffb = _cff
			_cff--
			{
				_gacb._bga.AddToken(FmtTypeDigit, nil)
			}
			goto _bebb
		_daa:
			_dffb = _cff
			_cff--
			{
				_gacb._bga.AddToken(FmtTypeTime, _bbca[_dce:_dffb])
			}
			goto _bebb
		_dffbf:
			_dffb = _cff
			_cff--
			{
			}
			goto _bebb
		_fee:
			_dffb = _cff + 1
			{
				_gacb._bga.IsExponential = true
			}
			goto _bebb
		_daf:
			_dffb = _cff + 1
			{
				_gacb._bga.AddToken(FmtTypeLiteral, _bbca[_dce+1:_dffb])
			}
			goto _bebb
		_bebb:
			_dce = 0
			if _cff++; _cff == _dcgc {
				goto _fdb
			}
		_dgfb:
			_dce = _cff
			switch _bbca[_cff] {
			case 34:
				goto _cga
			case 35:
				goto _gee
			case 36:
				goto _faag
			case 37:
				goto _cbg
			case 44:
				goto _egce
			case 46:
				goto _egeg
			case 47:
				goto _fdce
			case 48:
				goto _gfff
			case 58:
				goto _aceb
			case 59:
				goto _dbfb
			case 63:
				goto _ecf
			case 64:
				goto _bdg
			case 65:
				goto _ega
			case 69:
				goto _ddbd
			case 71:
				goto _aaf
			case 91:
				goto _adda
			case 92:
				goto _cace
			case 95:
				goto _bdf
			case 100:
				goto _fdce
			case 104:
				goto _aceb
			case 109:
				goto _ecfb
			case 115:
				goto _agbb
			case 121:
				goto _gbbeg
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _edge
			}
			goto _cgdb
		_cga:
			_dffb = _cff + 1
			_ddd = 20
			goto _dge
		_dge:
			if _cff++; _cff == _dcgc {
				goto _eggf
			}
		_dbdd:
			if _bbca[_cff] == 34 {
				goto _egeb
			}
			goto _cbc
		_cbc:
			if _cff++; _cff == _dcgc {
				goto _eec
			}
		_bfd:
			if _bbca[_cff] == 34 {
				goto _egeb
			}
			goto _cbc
		_egeb:
			_dffb = _cff + 1
			_ddd = 21
			goto _eeb
		_eeb:
			if _cff++; _cff == _dcgc {
				goto _daeg
			}
		_adg:
			if _bbca[_cff] == 34 {
				goto _cbc
			}
			goto _aa
		_gee:
			_dffb = _cff + 1
			_ddd = 3
			goto _bae
		_bae:
			if _cff++; _cff == _dcgc {
				goto _beab
			}
		_eef:
			switch _bbca[_cff] {
			case 35:
				goto _dcge
			case 37:
				goto _dcge
			case 44:
				goto _bfee
			case 47:
				goto _abeg
			case 48:
				goto _dcge
			case 63:
				goto _dcge
			}
			goto _egb
		_dcge:
			if _cff++; _cff == _dcgc {
				goto _eda
			}
		_gef:
			switch _bbca[_cff] {
			case 35:
				goto _dcge
			case 37:
				goto _dcge
			case 47:
				goto _abeg
			case 48:
				goto _dcge
			case 63:
				goto _dcge
			}
			goto _bfc
		_abeg:
			if _cff++; _cff == _dcgc {
				goto _ddbdf
			}
		_dfeg:
			switch _bbca[_cff] {
			case 35:
				goto _ddda
			case 37:
				goto _ddg
			case 48:
				goto _cbb
			case 63:
				goto _ddda
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _gccd
			}
			goto _bfc
		_ddda:
			_dffb = _cff + 1
			goto _cdc
		_cdc:
			if _cff++; _cff == _dcgc {
				goto _cffc
			}
		_bede:
			switch _bbca[_cff] {
			case 35:
				goto _ddda
			case 37:
				goto _ddda
			case 44:
				goto _ddda
			case 46:
				goto _ddda
			case 48:
				goto _ddda
			case 63:
				goto _ddda
			case 65:
				goto _cge
			}
			goto _ecaf
		_cge:
			if _cff++; _cff == _dcgc {
				goto _ddec
			}
		_bedf:
			switch _bbca[_cff] {
			case 47:
				goto _bbga
			case 77:
				goto _cgb
			}
			goto _fe
		_bbga:
			if _cff++; _cff == _dcgc {
				goto _bgaa
			}
		_acg:
			if _bbca[_cff] == 80 {
				goto _aea
			}
			goto _fe
		_aea:
			_dffb = _cff + 1
			goto _cabb
		_cabb:
			if _cff++; _cff == _dcgc {
				goto _fcdd
			}
		_dbfa:
			if _bbca[_cff] == 65 {
				goto _cge
			}
			goto _ecaf
		_cgb:
			if _cff++; _cff == _dcgc {
				goto _baac
			}
		_dddd:
			if _bbca[_cff] == 47 {
				goto _dgc
			}
			goto _fe
		_dgc:
			if _cff++; _cff == _dcgc {
				goto _fff
			}
		_ade:
			if _bbca[_cff] == 80 {
				goto _ceda
			}
			goto _fe
		_ceda:
			if _cff++; _cff == _dcgc {
				goto _feg
			}
		_aca:
			if _bbca[_cff] == 77 {
				goto _aea
			}
			goto _fe
		_ddg:
			if _cff++; _cff == _dcgc {
				goto _gffe
			}
		_abg:
			switch _bbca[_cff] {
			case 35:
				goto _cbfb
			case 37:
				goto _ddf
			case 63:
				goto _cbfb
			}
			if 48 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _eaf
			}
			goto _bfc
		_cbfb:
			_dffb = _cff + 1
			goto _edg
		_edg:
			if _cff++; _cff == _dcgc {
				goto _bddc
			}
		_agda:
			switch _bbca[_cff] {
			case 35:
				goto _ddda
			case 37:
				goto _fca
			case 44:
				goto _ddda
			case 46:
				goto _ddda
			case 48:
				goto _ddda
			case 63:
				goto _ddda
			case 65:
				goto _cge
			}
			goto _ecaf
		_fca:
			if _cff++; _cff == _dcgc {
				goto _adgb
			}
		_dgcf:
			switch _bbca[_cff] {
			case 35:
				goto _ada
			case 44:
				goto _ada
			case 46:
				goto _ada
			case 48:
				goto _ada
			case 63:
				goto _ada
			}
			goto _fe
		_ada:
			_dffb = _cff + 1
			goto _gdcf
		_gdcf:
			if _cff++; _cff == _dcgc {
				goto _ebgb
			}
		_ddcd:
			switch _bbca[_cff] {
			case 35:
				goto _ada
			case 44:
				goto _ada
			case 46:
				goto _ada
			case 48:
				goto _ada
			case 63:
				goto _ada
			case 65:
				goto _cge
			}
			goto _ecaf
		_ddf:
			if _cff++; _cff == _dcgc {
				goto _ebe
			}
		_dbae:
			if _bbca[_cff] == 37 {
				goto _ddf
			}
			if 48 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _eaf
			}
			goto _bfc
		_eaf:
			_dffb = _cff + 1
			_ddd = 13
			goto _cbe
		_cbe:
			if _cff++; _cff == _dcgc {
				goto _ddab
			}
		_abfc:
			switch _bbca[_cff] {
			case 35:
				goto _ddda
			case 37:
				goto _cecc
			case 44:
				goto _ddda
			case 46:
				goto _ddda
			case 48:
				goto _bbfcg
			case 63:
				goto _ddda
			case 65:
				goto _cge
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _eaf
			}
			goto _ecaf
		_cecc:
			if _cff++; _cff == _dcgc {
				goto _egea
			}
		_gca:
			switch _bbca[_cff] {
			case 35:
				goto _ada
			case 37:
				goto _ddf
			case 44:
				goto _ada
			case 46:
				goto _ada
			case 63:
				goto _ada
			}
			if 48 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _eaf
			}
			goto _fe
		_bbfcg:
			_dffb = _cff + 1
			goto _bbfd
		_bbfd:
			if _cff++; _cff == _dcgc {
				goto _egbd
			}
		_cbea:
			switch _bbca[_cff] {
			case 35:
				goto _ddda
			case 37:
				goto _bbfcg
			case 44:
				goto _ddda
			case 46:
				goto _ddda
			case 48:
				goto _bbfcg
			case 63:
				goto _ddda
			case 65:
				goto _cge
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _eaf
			}
			goto _ecaf
		_cbb:
			_dffb = _cff + 1
			goto _ffd
		_ffd:
			if _cff++; _cff == _dcgc {
				goto _dag
			}
		_baab:
			switch _bbca[_cff] {
			case 35:
				goto _ddda
			case 37:
				goto _bbfcg
			case 44:
				goto _ddda
			case 46:
				goto _ddda
			case 48:
				goto _cbb
			case 63:
				goto _ddda
			case 65:
				goto _cge
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _gccd
			}
			goto _ecaf
		_gccd:
			_dffb = _cff + 1
			goto _aed
		_aed:
			if _cff++; _cff == _dcgc {
				goto _cdb
			}
		_fgda:
			switch _bbca[_cff] {
			case 35:
				goto _ddda
			case 37:
				goto _eaf
			case 44:
				goto _ddda
			case 46:
				goto _ddda
			case 48:
				goto _cbb
			case 63:
				goto _ddda
			case 65:
				goto _cge
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _gccd
			}
			goto _ecaf
		_bfee:
			if _cff++; _cff == _dcgc {
				goto _fde
			}
		_ceccb:
			if _bbca[_cff] == 35 {
				goto _geda
			}
			goto _bag
		_cbg:
			_dffb = _cff + 1
			_ddd = 8
			goto _bfca
		_bfca:
			if _cff++; _cff == _dcgc {
				goto _dfc
			}
		_dgfg:
			switch _bbca[_cff] {
			case 35:
				goto _dgfbd
			case 37:
				goto _gfee
			case 48:
				goto _ecg
			case 63:
				goto _dgfbd
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _aba
			}
			goto _abc
		_dgfbd:
			if _cff++; _cff == _dcgc {
				goto _dgfgf
			}
		_def:
			switch _bbca[_cff] {
			case 35:
				goto _dgfbd
			case 47:
				goto _abeg
			case 48:
				goto _dgfbd
			case 63:
				goto _dgfbd
			}
			goto _caac
		_gfee:
			if _cff++; _cff == _dcgc {
				goto _bda
			}
		_gcg:
			if _bbca[_cff] == 37 {
				goto _gfee
			}
			if 48 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _aba
			}
			goto _bfc
		_aba:
			if _cff++; _cff == _dcgc {
				goto _efd
			}
		_ggf:
			switch _bbca[_cff] {
			case 37:
				goto _gfee
			case 47:
				goto _abeg
			}
			if 48 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _aba
			}
			goto _bfc
		_ecg:
			if _cff++; _cff == _dcgc {
				goto _eagb
			}
		_eea:
			switch _bbca[_cff] {
			case 35:
				goto _dgfbd
			case 37:
				goto _gfee
			case 47:
				goto _abeg
			case 48:
				goto _ecg
			case 63:
				goto _dgfbd
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _aba
			}
			goto _caac
		_fdce:
			_dffb = _cff + 1
			goto _fad
		_fad:
			if _cff++; _cff == _dcgc {
				goto _aeg
			}
		_ecbd:
			switch _bbca[_cff] {
			case 47:
				goto _fdce
			case 100:
				goto _fdce
			case 109:
				goto _fdce
			case 121:
				goto _bdge
			}
			goto _dab
		_bdge:
			if _cff++; _cff == _dcgc {
				goto _abeb
			}
		_afg:
			if _bbca[_cff] == 121 {
				goto _fdce
			}
			goto _eac
		_gfff:
			_dffb = _cff + 1
			_ddd = 2
			goto _fbg
		_fbg:
			if _cff++; _cff == _dcgc {
				goto _efdc
			}
		_gcga:
			switch _bbca[_cff] {
			case 35:
				goto _dcge
			case 37:
				goto _aeba
			case 47:
				goto _abeg
			case 48:
				goto _eafe
			case 63:
				goto _dcge
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _ebdd
			}
			goto _gacf
		_aeba:
			if _cff++; _cff == _dcgc {
				goto _afa
			}
		_aac:
			switch _bbca[_cff] {
			case 35:
				goto _dcge
			case 37:
				goto _aeba
			case 47:
				goto _abeg
			case 48:
				goto _aeba
			case 63:
				goto _dcge
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _aba
			}
			goto _agd
		_eafe:
			if _cff++; _cff == _dcgc {
				goto _bbbc
			}
		_bedeb:
			switch _bbca[_cff] {
			case 35:
				goto _dcge
			case 37:
				goto _aeba
			case 47:
				goto _abeg
			case 48:
				goto _eafe
			case 63:
				goto _dcge
			}
			if 49 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _ebdd
			}
			goto _agd
		_ebdd:
			if _cff++; _cff == _dcgc {
				goto _eee
			}
		_gdf:
			switch _bbca[_cff] {
			case 37:
				goto _aba
			case 47:
				goto _abeg
			}
			if 48 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _ebdd
			}
			goto _bfc
		_edge:
			_dffb = _cff + 1
			_ddd = 20
			goto _cdcf
		_cdcf:
			if _cff++; _cff == _dcgc {
				goto _ebab
			}
		_bab:
			switch _bbca[_cff] {
			case 37:
				goto _aba
			case 47:
				goto _abeg
			}
			if 48 <= _bbca[_cff] && _bbca[_cff] <= 57 {
				goto _ebdd
			}
			goto _efb
		_aceb:
			_dffb = _cff + 1
			_ddd = 15
			goto _ecgd
		_ecgd:
			if _cff++; _cff == _dcgc {
				goto _ggg
			}
		_gdb:
			switch _bbca[_cff] {
			case 58:
				goto _aceb
			case 65:
				goto _abgc
			case 104:
				goto _aceb
			case 109:
				goto _aceb
			case 115:
				goto _agbb
			}
			goto _daa
		_abgc:
			if _cff++; _cff == _dcgc {
				goto _fegf
			}
		_bcd:
			switch _bbca[_cff] {
			case 47:
				goto _bgca
			case 77:
				goto _efg
			}
			goto _bfc
		_bgca:
			if _cff++; _cff == _dcgc {
				goto _cbcf
			}
		_dbdf:
			if _bbca[_cff] == 80 {
				goto _aceb
			}
			goto _bfc
		_efg:
			if _cff++; _cff == _dcgc {
				goto _gbbd
			}
		_aga:
			if _bbca[_cff] == 47 {
				goto _gcb
			}
			goto _bfc
		_gcb:
			if _cff++; _cff == _dcgc {
				goto _gaa
			}
		_dga:
			if _bbca[_cff] == 80 {
				goto _dcad
			}
			goto _bfc
		_dcad:
			if _cff++; _cff == _dcgc {
				goto _dbgc
			}
		_fdcf:
			if _bbca[_cff] == 77 {
				goto _aceb
			}
			goto _bfc
		_agbb:
			_dffb = _cff + 1
			_ddd = 15
			goto _eeba
		_eeba:
			if _cff++; _cff == _dcgc {
				goto _gda
			}
		_dcaa:
			switch _bbca[_cff] {
			case 46:
				goto _baed
			case 58:
				goto _aceb
			case 65:
				goto _abgc
			case 104:
				goto _aceb
			case 109:
				goto _aceb
			case 115:
				goto _agbb
			}
			goto _daa
		_baed:
			if _cff++; _cff == _dcgc {
				goto _egag
			}
		_bdd:
			if _bbca[_cff] == 48 {
				goto _cafd
			}
			goto _fdcb
		_cafd:
			_dffb = _cff + 1
			_ddd = 15
			goto _gcee
		_gcee:
			if _cff++; _cff == _dcgc {
				goto _cggd
			}
		_egfe:
			switch _bbca[_cff] {
			case 48:
				goto _agbd
			case 58:
				goto _aceb
			case 65:
				goto _abgc
			case 104:
				goto _aceb
			case 109:
				goto _aceb
			case 115:
				goto _agbb
			}
			goto _daa
		_agbd:
			_dffb = _cff + 1
			_ddd = 15
			goto _cde
		_cde:
			if _cff++; _cff == _dcgc {
				goto _dddg
			}
		_dbeb:
			switch _bbca[_cff] {
			case 48:
				goto _aceb
			case 58:
				goto _aceb
			case 65:
				goto _abgc
			case 104:
				goto _aceb
			case 109:
				goto _aceb
			case 115:
				goto _agbb
			}
			goto _daa
		_ecf:
			_dffb = _cff + 1
			_ddd = 5
			goto _bbe
		_bbe:
			if _cff++; _cff == _dcgc {
				goto _bfdc
			}
		_dccd:
			switch _bbca[_cff] {
			case 35:
				goto _dcge
			case 37:
				goto _dcge
			case 47:
				goto _abeg
			case 48:
				goto _dcge
			case 63:
				goto _dcge
			}
			goto _dffbf
		_ega:
			_dffb = _cff + 1
			_ddd = 20
			goto _agdg
		_agdg:
			if _cff++; _cff == _dcgc {
				goto _aeeg
			}
		_cabf:
			switch _bbca[_cff] {
			case 47:
				goto _bgca
			case 77:
				goto _efg
			}
			goto _efb
		_ddbd:
			if _cff++; _cff == _dcgc {
				goto _dggg
			}
		_fdg:
			switch _bbca[_cff] {
			case 43:
				goto _fee
			case 45:
				goto _fee
			}
			goto _efb
		_aaf:
			_dffb = _cff + 1
			goto _dec
		_dec:
			if _cff++; _cff == _dcgc {
				goto _agbg
			}
		_ebg:
			if _bbca[_cff] == 101 {
				goto _fdcbe
			}
			goto _efb
		_fdcbe:
			if _cff++; _cff == _dcgc {
				goto _dfee
			}
		_ecbe:
			if _bbca[_cff] == 110 {
				goto _bcbd
			}
			goto _ced
		_bcbd:
			if _cff++; _cff == _dcgc {
				goto _ffa
			}
		_cbfa:
			if _bbca[_cff] == 101 {
				goto _dega
			}
			goto _ced
		_dega:
			if _cff++; _cff == _dcgc {
				goto _fdf
			}
		_eaca:
			if _bbca[_cff] == 114 {
				goto _gcab
			}
			goto _ced
		_gcab:
			if _cff++; _cff == _dcgc {
				goto _ffg
			}
		_afb:
			if _bbca[_cff] == 97 {
				goto _gbaf
			}
			goto _ced
		_gbaf:
			if _cff++; _cff == _dcgc {
				goto _beee
			}
		_gbcb:
			if _bbca[_cff] == 108 {
				goto _bfg
			}
			goto _ced
		_adda:
			_dffb = _cff + 1
			_ddd = 20
			goto _fcd
		_fcd:
			if _cff++; _cff == _dcgc {
				goto _faed
			}
		_fadf:
			switch _bbca[_cff] {
			case 104:
				goto _eae
			case 109:
				goto _eae
			case 115:
				goto _eae
			}
			goto _ffcf
		_ffcf:
			if _cff++; _cff == _dcgc {
				goto _dee
			}
		_fgeb:
			if _bbca[_cff] == 93 {
				goto _dfa
			}
			goto _ffcf
		_dfa:
			_dffb = _cff + 1
			_ddd = 18
			goto _fec
		_bec:
			_dffb = _cff + 1
			_ddd = 16
			goto _fec
		_fec:
			if _cff++; _cff == _dcgc {
				goto _cgag
			}
		_aeac:
			if _bbca[_cff] == 93 {
				goto _dfa
			}
			goto _ffcf
		_eae:
			if _cff++; _cff == _dcgc {
				goto _dgfa
			}
		_ebaf:
			if _bbca[_cff] == 93 {
				goto _bec
			}
			goto _ffcf
		_cace:
			if _cff++; _cff == _dcgc {
				goto _dfg
			}
		_ffdd:
			goto _daf
		_ecfb:
			_dffb = _cff + 1
			_ddd = 14
			goto _cgad
		_cgad:
			if _cff++; _cff == _dcgc {
				goto _agfa
			}
		_eade:
			switch _bbca[_cff] {
			case 47:
				goto _fdce
			case 58:
				goto _aceb
			case 65:
				goto _abgc
			case 100:
				goto _fdce
			case 104:
				goto _aceb
			case 109:
				goto _ecfb
			case 115:
				goto _agbb
			case 121:
				goto _bdge
			}
			goto _dab
		_gbbeg:
			if _cff++; _cff == _dcgc {
				goto _edf
			}
		_dbgd:
			if _bbca[_cff] == 121 {
				goto _fdce
			}
			goto _efb
		_bgg:
		_fdb:
			_addg = 34
			goto _bdcc
		_eggf:
			_addg = 35
			goto _bdcc
		_eec:
			_addg = 0
			goto _bdcc
		_daeg:
			_addg = 36
			goto _bdcc
		_beab:
			_addg = 37
			goto _bdcc
		_eda:
			_addg = 1
			goto _bdcc
		_ddbdf:
			_addg = 2
			goto _bdcc
		_cffc:
			_addg = 38
			goto _bdcc
		_ddec:
			_addg = 3
			goto _bdcc
		_bgaa:
			_addg = 4
			goto _bdcc
		_fcdd:
			_addg = 39
			goto _bdcc
		_baac:
			_addg = 5
			goto _bdcc
		_fff:
			_addg = 6
			goto _bdcc
		_feg:
			_addg = 7
			goto _bdcc
		_gffe:
			_addg = 8
			goto _bdcc
		_bddc:
			_addg = 40
			goto _bdcc
		_adgb:
			_addg = 9
			goto _bdcc
		_ebgb:
			_addg = 41
			goto _bdcc
		_ebe:
			_addg = 10
			goto _bdcc
		_ddab:
			_addg = 42
			goto _bdcc
		_egea:
			_addg = 11
			goto _bdcc
		_egbd:
			_addg = 43
			goto _bdcc
		_dag:
			_addg = 44
			goto _bdcc
		_cdb:
			_addg = 45
			goto _bdcc
		_fde:
			_addg = 12
			goto _bdcc
		_dfc:
			_addg = 46
			goto _bdcc
		_dgfgf:
			_addg = 13
			goto _bdcc
		_bda:
			_addg = 14
			goto _bdcc
		_efd:
			_addg = 15
			goto _bdcc
		_eagb:
			_addg = 16
			goto _bdcc
		_aeg:
			_addg = 47
			goto _bdcc
		_abeb:
			_addg = 17
			goto _bdcc
		_efdc:
			_addg = 48
			goto _bdcc
		_afa:
			_addg = 18
			goto _bdcc
		_bbbc:
			_addg = 19
			goto _bdcc
		_eee:
			_addg = 20
			goto _bdcc
		_ebab:
			_addg = 49
			goto _bdcc
		_ggg:
			_addg = 50
			goto _bdcc
		_fegf:
			_addg = 21
			goto _bdcc
		_cbcf:
			_addg = 22
			goto _bdcc
		_gbbd:
			_addg = 23
			goto _bdcc
		_gaa:
			_addg = 24
			goto _bdcc
		_dbgc:
			_addg = 25
			goto _bdcc
		_gda:
			_addg = 51
			goto _bdcc
		_egag:
			_addg = 26
			goto _bdcc
		_cggd:
			_addg = 52
			goto _bdcc
		_dddg:
			_addg = 53
			goto _bdcc
		_bfdc:
			_addg = 54
			goto _bdcc
		_aeeg:
			_addg = 55
			goto _bdcc
		_dggg:
			_addg = 56
			goto _bdcc
		_agbg:
			_addg = 57
			goto _bdcc
		_dfee:
			_addg = 27
			goto _bdcc
		_ffa:
			_addg = 28
			goto _bdcc
		_fdf:
			_addg = 29
			goto _bdcc
		_ffg:
			_addg = 30
			goto _bdcc
		_beee:
			_addg = 31
			goto _bdcc
		_faed:
			_addg = 58
			goto _bdcc
		_dee:
			_addg = 32
			goto _bdcc
		_cgag:
			_addg = 59
			goto _bdcc
		_dgfa:
			_addg = 33
			goto _bdcc
		_dfg:
			_addg = 60
			goto _bdcc
		_agfa:
			_addg = 61
			goto _bdcc
		_edf:
			_addg = 62
			goto _bdcc
		_bdcc:
			{
			}
			if _cff == _dggd {
				switch _addg {
				case 35:
					goto _efb
				case 0:
					goto _bfc
				case 36:
					goto _aa
				case 37:
					goto _egb
				case 1:
					goto _bfc
				case 2:
					goto _bfc
				case 38:
					goto _ecaf
				case 3:
					goto _fe
				case 4:
					goto _fe
				case 39:
					goto _ecaf
				case 5:
					goto _fe
				case 6:
					goto _fe
				case 7:
					goto _fe
				case 8:
					goto _bfc
				case 40:
					goto _ecaf
				case 9:
					goto _fe
				case 41:
					goto _ecaf
				case 10:
					goto _bfc
				case 42:
					goto _ecaf
				case 11:
					goto _fe
				case 43:
					goto _ecaf
				case 44:
					goto _ecaf
				case 45:
					goto _ecaf
				case 12:
					goto _bag
				case 46:
					goto _abc
				case 13:
					goto _caac
				case 14:
					goto _bfc
				case 15:
					goto _bfc
				case 16:
					goto _caac
				case 47:
					goto _dab
				case 17:
					goto _eac
				case 48:
					goto _gacf
				case 18:
					goto _agd
				case 19:
					goto _agd
				case 20:
					goto _bfc
				case 49:
					goto _efb
				case 50:
					goto _daa
				case 21:
					goto _bfc
				case 22:
					goto _bfc
				case 23:
					goto _bfc
				case 24:
					goto _bfc
				case 25:
					goto _bfc
				case 51:
					goto _daa
				case 26:
					goto _fdcb
				case 52:
					goto _daa
				case 53:
					goto _daa
				case 54:
					goto _dffbf
				case 55:
					goto _efb
				case 56:
					goto _efb
				case 57:
					goto _efb
				case 27:
					goto _ced
				case 28:
					goto _ced
				case 29:
					goto _ced
				case 30:
					goto _ced
				case 31:
					goto _ced
				case 58:
					goto _efb
				case 32:
					goto _bfc
				case 59:
					goto _bfc
				case 33:
					goto _ced
				case 60:
					goto _efb
				case 61:
					goto _dab
				case 62:
					goto _efb
				}
			}
		}
		if _dce > 0 {
			copy(_bbca[0:], _bbca[_dce:])
		}
	}
	_ = _dggd
	if _addg == _eag {
		_cg.Log.Debug("\u0066o\u0072m\u0061\u0074\u0020\u0070\u0061r\u0073\u0065 \u0065\u0072\u0072\u006f\u0072")
	}
}

// FmtType is the type of a format token.
//
//go:generate stringer -type=FmtType
type FmtType byte

// Format is a parsed number format.
type Format struct {
	Whole         []Token
	Fractional    []Token
	Exponent      []Token
	IsExponential bool
	_bf           bool
	_acc          bool
	_ca           bool
	_cgg          bool
	_egf          bool
	_af           bool
	_bg           int64
	_cfg          int
}

func _dcf(_dcdc int64) int64 {
	if _dcdc < 0 {
		return -_dcdc
	}
	return _dcdc
}
func Parse(s string) []Format {
	_cdg := Lexer{}
	_cdg.Lex(_b.NewReader(s))
	_cdg._dfe = append(_cdg._dfe, _cdg._bga)
	return _cdg._dfe
}
func _bgf(_aef, _gfe float64, _geb Format) []byte {
	if len(_geb.Fractional) == 0 {
		return nil
	}
	_cfe := _cf.AppendFloat(nil, _aef, 'f', -1, 64)
	if len(_cfe) > 2 {
		_cfe = _cfe[2:]
	} else {
		_cfe = nil
	}
	_cbf := make([]byte, 0, len(_cfe))
	_cbf = append(_cbf, '.')
	_gg := 0
_bdc:
	for _ged := 0; _ged < len(_geb.Fractional); _ged++ {
		_eadc := _ged
		_deb := _geb.Fractional[_ged]
		switch _deb.Type {
		case FmtTypeDigit:
			if _eadc < len(_cfe) {
				_cbf = append(_cbf, _cfe[_eadc])
				_gg++
			} else {
				_cbf = append(_cbf, '0')
			}
		case FmtTypeDigitOpt:
			if _eadc >= 0 {
				_cbf = append(_cbf, _cfe[_eadc])
				_gg++
			} else {
				break _bdc
			}
		case FmtTypeLiteral:
			_cbf = append(_cbf, _deb.Literal)
		default:
			_cg.Log.Debug("\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020f\u0072\u0061\u0063\u0074\u0069\u006f\u006ea\u006c\u0020\u0025\u0076", _deb)
		}
	}
	return _cbf
}
func _dcc(_age _f.Time, _ebd float64, _fge string) []byte {
	_da := []byte{}
	_ccf := 0
	for _adc := 0; _adc < len(_fge); _adc++ {
		var _fgab string
		if _fge[_adc] == ':' {
			_fgab = string(_fge[_ccf:_adc])
			_ccf = _adc + 1
		} else if _adc == len(_fge)-1 {
			_fgab = string(_fge[_ccf : _adc+1])
		} else {
			continue
		}
		switch _fgab {
		case "\u0064":
			_da = _age.AppendFormat(_da, "\u0032")
		case "\u0068":
			_da = _age.AppendFormat(_da, "\u0033")
		case "\u0068\u0068":
			_da = _age.AppendFormat(_da, "\u0031\u0035")
		case "\u006d":
			_da = _age.AppendFormat(_da, "\u0034")
		case "\u006d\u006d":
			_da = _age.AppendFormat(_da, "\u0030\u0034")
		case "\u0073":
			_da = _age.Round(_f.Second).AppendFormat(_da, "\u0035")
		case "\u0073\u002e\u0030":
			_da = _age.Round(_f.Second/10).AppendFormat(_da, "\u0035\u002e\u0030")
		case "\u0073\u002e\u0030\u0030":
			_da = _age.Round(_f.Second/100).AppendFormat(_da, "\u0035\u002e\u0030\u0030")
		case "\u0073\u002e\u00300\u0030":
			_da = _age.Round(_f.Second/1000).AppendFormat(_da, "\u0035\u002e\u00300\u0030")
		case "\u0073\u0073":
			_da = _age.Round(_f.Second).AppendFormat(_da, "\u0030\u0035")
		case "\u0073\u0073\u002e\u0030":
			_da = _age.Round(_f.Second/10).AppendFormat(_da, "\u0030\u0035\u002e\u0030")
		case "\u0073\u0073\u002e0\u0030":
			_da = _age.Round(_f.Second/100).AppendFormat(_da, "\u0030\u0035\u002e0\u0030")
		case "\u0073\u0073\u002e\u0030\u0030\u0030":
			_da = _age.Round(_f.Second/1000).AppendFormat(_da, "\u0030\u0035\u002e\u0030\u0030\u0030")
		case "\u0041\u004d\u002fP\u004d":
			_da = _age.AppendFormat(_da, "\u0050\u004d")
		case "\u005b\u0068\u005d":
			_da = _cf.AppendInt(_da, int64(_ebd*24), 10)
		case "\u005b\u006d\u005d":
			_da = _cf.AppendInt(_da, int64(_ebd*24*60), 10)
		case "\u005b\u0073\u005d":
			_da = _cf.AppendInt(_da, int64(_ebd*24*60*60), 10)
		case "":
		default:
			_cg.Log.Debug("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0069\u006d\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073", _fgab)
		}
		if _fge[_adc] == ':' {
			_da = append(_da, ':')
		}
	}
	return _da
}

// NumberGeneric formats the number with the generic format which attemps to
// mimic Excel's general formatting.
func NumberGeneric(v float64) string {
	if _g.Abs(v) >= _fg || _g.Abs(v) <= _ae && v != 0 {
		return _faa(v)
	}
	_gce := make([]byte, 0, 15)
	_gce = _cf.AppendFloat(_gce, v, 'f', -1, 64)
	if len(_gce) > 11 {
		_ee := _gce[11] - '0'
		if _ee >= 5 && _ee <= 9 {
			_gce[10]++
			_gce = _gce[0:11]
			_gce = _gbc(_gce)
		}
		_gce = _gce[0:11]
	} else if len(_gce) == 11 {
		if _gce[len(_gce)-1] == '9' {
			_gce[len(_gce)-1]++
			_gce = _gbc(_gce)
		}
	}
	_gce = _eadda(_gce)
	return string(_gce)
}
func _gbc(_baa []byte) []byte {
	for _beb := len(_baa) - 1; _beb > 0; _beb-- {
		if _baa[_beb] == '9'+1 {
			_baa[_beb] = '0'
			if _baa[_beb-1] == '.' {
				_beb--
			}
			_baa[_beb-1]++
		}
	}
	if _baa[0] == '9'+1 {
		_baa[0] = '0'
		copy(_baa[1:], _baa[0:])
		_baa[0] = '1'
	}
	return _baa
}

const _eag int = -1

func _ccg(_cgcb int64, _abe Format) []byte {
	if !_abe.IsExponential || len(_abe.Exponent) == 0 {
		return nil
	}
	_debd := _cf.AppendInt(nil, _dcf(_cgcb), 10)
	_cacf := make([]byte, 0, len(_debd)+2)
	_cacf = append(_cacf, 'E')
	if _cgcb >= 0 {
		_cacf = append(_cacf, '+')
	} else {
		_cacf = append(_cacf, '-')
		_cgcb *= -1
	}
	_gde := 0
_eadd:
	for _gbb := len(_abe.Exponent) - 1; _gbb >= 0; _gbb-- {
		_bedg := len(_debd) - 1 - _gde
		_gab := _abe.Exponent[_gbb]
		switch _gab.Type {
		case FmtTypeDigit:
			if _bedg >= 0 {
				_cacf = append(_cacf, _debd[_bedg])
				_gde++
			} else {
				_cacf = append(_cacf, '0')
			}
		case FmtTypeDigitOpt:
			if _bedg >= 0 {
				_cacf = append(_cacf, _debd[_bedg])
				_gde++
			} else {
				for _cd := _gbb; _cd >= 0; _cd-- {
					_gdc := _abe.Exponent[_cd]
					if _gdc.Type == FmtTypeLiteral {
						_cacf = append(_cacf, _gdc.Literal)
					}
				}
				break _eadd
			}
		case FmtTypeLiteral:
			_cacf = append(_cacf, _gab.Literal)
		default:
			_cg.Log.Debug("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u0065\u0078p\u0020\u0025\u0076", _gab)
		}
	}
	if _gde < len(_debd) {
		_cacf = append(_cacf, _debd[len(_debd)-_gde-1:_gde-1]...)
	}
	_ga(_cacf[2:])
	return _cacf
}

type Lexer struct {
	_bga Format
	_dfe []Format
}

func _ga(_df []byte) []byte {
	for _edb := 0; _edb < len(_df)/2; _edb++ {
		_bbg := len(_df) - 1 - _edb
		_df[_edb], _df[_bbg] = _df[_bbg], _df[_edb]
	}
	return _df
}

var _d = [...]uint8{0, 14, 26, 41, 53, 67, 81, 94, 118, 135, 146, 157, 172, 183}

const (
	FmtTypeLiteral FmtType = iota
	FmtTypeDigit
	FmtTypeDigitOpt
	FmtTypeComma
	FmtTypeDecimal
	FmtTypePercent
	FmtTypeDollar
	FmtTypeDigitOptThousands
	FmtTypeUnderscore
	FmtTypeDate
	FmtTypeTime
	FmtTypeFraction
	FmtTypeText
)

// Number is used to format a number with a format string.  If the format
// string is empty, then General number formatting is used which attempts to mimic
// Excel's general formatting.
func Number(v float64, f string) string {
	if f == "" || f == "\u0047e\u006e\u0065\u0072\u0061\u006c" || f == "\u0040" {
		return NumberGeneric(v)
	}
	_dca := Parse(f)
	if len(_dca) == 1 {
		return _dg(v, _dca[0], false)
	} else if len(_dca) > 1 && v < 0 {
		return _dg(v, _dca[1], true)
	} else if len(_dca) > 2 && v == 0 {
		return _dg(v, _dca[2], false)
	}
	return _dg(v, _dca[0], false)
}
func (_eca *Lexer) nextFmt() { _eca._dfe = append(_eca._dfe, _eca._bga); _eca._bga = Format{} }

// AddToken adds a format token to the format.
func (_dc *Format) AddToken(t FmtType, l []byte) {
	if _dc._egf {
		_dc._egf = false
		return
	}
	switch t {
	case FmtTypeDecimal:
		_dc._af = true
	case FmtTypeUnderscore:
		_dc._egf = true
	case FmtTypeText:
		_dc.Whole = append(_dc.Whole, Token{Type: t})
	case FmtTypeDate, FmtTypeTime:
		_dc.Whole = append(_dc.Whole, Token{Type: t, DateTime: string(l)})
	case FmtTypePercent:
		_dc._acc = true
		t = FmtTypeLiteral
		l = []byte{'%'}
		fallthrough
	case FmtTypeDigitOpt:
		fallthrough
	case FmtTypeLiteral, FmtTypeDigit, FmtTypeDollar, FmtTypeComma:
		if l == nil {
			l = []byte{0}
		}
		for _, _cac := range l {
			if _dc.IsExponential {
				_dc.Exponent = append(_dc.Exponent, Token{Type: t, Literal: _cac})
			} else if !_dc._af {
				_dc.Whole = append(_dc.Whole, Token{Type: t, Literal: _cac})
			} else {
				_dc.Fractional = append(_dc.Fractional, Token{Type: t, Literal: _cac})
			}
		}
	case FmtTypeDigitOptThousands:
		_dc._cgg = true
	case FmtTypeFraction:
		_bd := _b.Split(string(l), "\u002f")
		if len(_bd) == 2 {
			_dc._bf = true
			_dc._bg, _ = _cf.ParseInt(_bd[1], 10, 64)
			for _, _be := range _bd[1] {
				if _be == '?' || _be == '0' {
					_dc._cfg++
				}
			}
		}
	default:
		_cg.Log.Debug("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0070\u0068\u0020\u0074y\u0070\u0065\u0020\u0069\u006e\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0025\u0076", t)
	}
}

const _ac = "\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u004c\u0069\u0074\u0065\u0072a\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0069\u0067\u0069\u0074\u0046\u006d\u0074\u0054y\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0043o\u006d\u006d\u0061\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0065\u0063\u0069\u006da\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065Pe\u0072\u0063e\u006e\u0074\u0046\u006d\u0074\u0054\u0079\u0070e\u0044\u006f\u006c\u006c\u0061\u0072\u0046\u006d\u0074Ty\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0054\u0068\u006f\u0075\u0073\u0061n\u0064\u0073\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0055n\u0064\u0065\u0072\u0073c\u006f\u0072\u0065\u0046\u006d\u0074T\u0079\u0070\u0065\u0044\u0061\u0074\u0065\u0046\u006d\u0074\u0054y\u0070e\u0054\u0069\u006d\u0065\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0046\u0072\u0061\u0063t\u0069\u006f\u006e\u0046\u006dt\u0054\u0079\u0070\u0065\u0054e\u0078\u0074"

func _eadda(_dbg []byte) []byte {
	_aeb := len(_dbg)
	_bdcd := false
	_gfg := false
	for _ceb := len(_dbg) - 1; _ceb >= 0; _ceb-- {
		if _dbg[_ceb] == '0' && !_gfg && !_bdcd {
			_aeb = _ceb
		} else if _dbg[_ceb] == '.' {
			_bdcd = true
		} else {
			_gfg = true
		}
	}
	if _bdcd && _gfg {
		if _dbg[_aeb-1] == '.' {
			_aeb--
		}
		return _dbg[0:_aeb]
	}
	return _dbg
}

const _fg = 1e11
const _bbfc int = 34
const _dba int = 0
const _cfda int = -1
const _gcc int = 34
const _cce int = 0

// Value formats a value as a number or string depending on  if it appears to be
// a number or string.
func Value(v string, f string) string {
	if IsNumber(v) {
		_dd, _ := _cf.ParseFloat(v, 64)
		return Number(_dd, f)
	}
	return String(v, f)
}

const _bcfd int = 0
