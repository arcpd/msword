package reference

import (
	_g "errors"
	_ee "fmt"
	_a "github.com/arcpd/msword/spreadsheet/update"
	_ge "regexp"
	_fd "strconv"
	_f "strings"
)

var _cfg = _ge.MustCompile("^\u005b\u0061\u002d\u007aA-\u005a]\u0028\u005b\u0061\u002d\u007aA\u002d\u005a\u005d\u003f\u0029\u0024")

// ParseCellReference parses a cell reference of the form 'A10' and splits it
// into column/row segments.
func ParseCellReference(s string) (CellReference, error) {
	s = _f.TrimSpace(s)
	if len(s) < 2 {
		return CellReference{}, _g.New("\u0063\u0065\u006c\u006c\u0020\u0072\u0065\u0066e\u0072\u0065\u006ece\u0020\u006d\u0075\u0073\u0074\u0020h\u0061\u0076\u0065\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0077o\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074e\u0072\u0073")
	}
	_d := CellReference{}
	_ac, _dg, _b := _dd(s)
	if _b != nil {
		return CellReference{}, _b
	}
	if _ac != "" {
		_d.SheetName = _ac
	}
	if s[0] == '$' {
		_d.AbsoluteColumn = true
		_dg = _dg[1:]
	}
	_eba := -1
_c:
	for _ff := 0; _ff < len(_dg); _ff++ {
		switch {
		case _dg[_ff] >= '0' && _dg[_ff] <= '9' || _dg[_ff] == '$':
			_eba = _ff
			break _c
		}
	}
	switch _eba {
	case 0:
		return CellReference{}, _ee.Errorf("\u006e\u006f\u0020\u006cet\u0074\u0065\u0072\u0020\u0070\u0072\u0065\u0066\u0069\u0078\u0020\u0069\u006e\u0020%\u0073", _dg)
	case -1:
		return CellReference{}, _ee.Errorf("\u006eo\u0020d\u0069\u0067\u0069\u0074\u0073\u0020\u0069\u006e\u0020\u0025\u0073", _dg)
	}
	_d.Column = _dg[0:_eba]
	if _dg[_eba] == '$' {
		_d.AbsoluteRow = true
		_eba++
	}
	_d.ColumnIdx = ColumnToIndex(_d.Column)
	_ae, _b := _fd.ParseUint(_dg[_eba:], 10, 32)
	if _b != nil {
		return CellReference{}, _ee.Errorf("e\u0072\u0072\u006f\u0072 p\u0061r\u0073\u0069\u006e\u0067\u0020r\u006f\u0077\u003a\u0020\u0025\u0073", _b)
	}
	if _ae == 0 {
		return CellReference{}, _ee.Errorf("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0072\u006f\u0077\u003a \u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u0030")
	}
	_d.RowIdx = uint32(_ae)
	return _d, nil
}

// String returns a string representation of CellReference.
func (_eb CellReference) String() string {
	_ed := make([]byte, 0, 4)
	if _eb.AbsoluteColumn {
		_ed = append(_ed, '$')
	}
	_ed = append(_ed, _eb.Column...)
	if _eb.AbsoluteRow {
		_ed = append(_ed, '$')
	}
	_ed = _fd.AppendInt(_ed, int64(_eb.RowIdx), 10)
	return string(_ed)
}

// ParseRangeReference splits a range reference of the form "A1:B5" into its
// components.
func ParseRangeReference(s string) (_fa, _fe CellReference, _ce error) {
	_gad, _df, _ce := _dd(s)
	if _ce != nil {
		return CellReference{}, CellReference{}, _ce
	}
	_eca := _f.Split(_df, "\u003a")
	if len(_eca) != 2 {
		return CellReference{}, CellReference{}, _g.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074")
	}
	if _gad != "" {
		_eca[0] = _gad + "\u0021" + _eca[0]
		_eca[1] = _gad + "\u0021" + _eca[1]
	}
	_gb, _ce := ParseCellReference(_eca[0])
	if _ce != nil {
		return CellReference{}, CellReference{}, _ce
	}
	_ad, _ce := ParseCellReference(_eca[1])
	if _ce != nil {
		return CellReference{}, CellReference{}, _ce
	}
	return _gb, _ad, nil
}

// ColumnReference is a parsed reference to a column.  Input is of the form 'A',
// '$C', etc.
type ColumnReference struct {
	ColumnIdx      uint32
	Column         string
	AbsoluteColumn bool
	SheetName      string
}

// String returns a string representation of ColumnReference.
func (_cf ColumnReference) String() string {
	_cg := make([]byte, 0, 4)
	if _cf.AbsoluteColumn {
		_cg = append(_cg, '$')
	}
	_cg = append(_cg, _cf.Column...)
	return string(_cg)
}

// ParseColumnRangeReference splits a range reference of the form "A:B" into its
// components.
func ParseColumnRangeReference(s string) (_eef, _cd ColumnReference, _fc error) {
	_ag := ""
	_aag := _f.Split(s, "\u0021")
	if len(_aag) == 2 {
		_ag = _aag[0]
		s = _aag[1]
	}
	_cee := _f.Split(s, "\u003a")
	if len(_cee) != 2 {
		return ColumnReference{}, ColumnReference{}, _g.New("i\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006eg\u0065\u0020\u0066\u006frm\u0061\u0074")
	}
	if _ag != "" {
		_cee[0] = _ag + "\u0021" + _cee[0]
		_cee[1] = _ag + "\u0021" + _cee[1]
	}
	_db, _fc := ParseColumnReference(_cee[0])
	if _fc != nil {
		return ColumnReference{}, ColumnReference{}, _fc
	}
	_aae, _fc := ParseColumnReference(_cee[1])
	if _fc != nil {
		return ColumnReference{}, ColumnReference{}, _fc
	}
	return _db, _aae, nil
}

// ColumnToIndex maps a column to a zero based index (e.g. A = 0, B = 1, AA = 26)
func ColumnToIndex(col string) uint32 {
	col = _f.ToUpper(col)
	_ec := uint32(0)
	for _, _eg := range col {
		_ec *= 26
		_ec += uint32(_eg - 'A' + 1)
	}
	return _ec - 1
}

// ParseColumnReference parses a column reference of the form 'Sheet1!A' and splits it
// into sheet name and column segments.
func ParseColumnReference(s string) (ColumnReference, error) {
	s = _f.TrimSpace(s)
	if len(s) < 1 {
		return ColumnReference{}, _g.New("\u0063\u006f\u006c\u0075\u006d\u006e \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0063\u0068a\u0072a\u0063\u0074\u0065\u0072")
	}
	_bf := ColumnReference{}
	_dga, _bc, _cb := _dd(s)
	if _cb != nil {
		return ColumnReference{}, _cb
	}
	if _dga != "" {
		_bf.SheetName = _dga
	}
	if _bc[0] == '$' {
		_bf.AbsoluteColumn = true
		_bc = _bc[1:]
	}
	if !_cfg.MatchString(_bc) {
		return ColumnReference{}, _g.New("\u0063\u006f\u006c\u0075\u006dn\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u006d\u0075s\u0074\u0020\u0062\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0041\u0020\u0061\u006e\u0064\u0020\u005a\u005a")
	}
	_bf.Column = _bc
	_bf.ColumnIdx = ColumnToIndex(_bf.Column)
	return _bf, nil
}

// IndexToColumn maps a column number to a column name (e.g. 0 = A, 1 = B, 26 = AA)
func IndexToColumn(col uint32) string {
	var _gd [64 + 1]byte
	_fg := len(_gd)
	_aa := col
	const _ab = 26
	for _aa >= _ab {
		_fg--
		_gcge := _aa / _ab
		_gd[_fg] = byte('A' + uint(_aa-_gcge*_ab))
		_aa = _gcge - 1
	}
	_fg--
	_gd[_fg] = byte('A' + uint(_aa))
	return string(_gd[_fg:])
}

// CellReference is a parsed reference to a cell.  Input is of the form 'A1',
// '$C$2', etc.
type CellReference struct {
	RowIdx         uint32
	ColumnIdx      uint32
	Column         string
	AbsoluteColumn bool
	AbsoluteRow    bool
	SheetName      string
}

func _dd(_edd string) (string, string, error) {
	_bb := ""
	_gf := _f.LastIndex(_edd, "\u0021")
	if _gf > -1 {
		_bb = _edd[:_gf]
		_edd = _edd[_gf+1:]
		if _bb == "" {
			return "", "", _g.New("\u0049n\u0076a\u006c\u0069\u0064\u0020\u0073h\u0065\u0065t\u0020\u006e\u0061\u006d\u0065")
		}
	}
	return _bb, _edd, nil
}

// Update updates reference to point one of the neighboring cells with respect to the update type after removing a row/column.
func (_ga *CellReference) Update(updateType _a.UpdateAction) *CellReference {
	switch updateType {
	case _a.UpdateActionRemoveColumn:
		_gc := _ga
		_gc.ColumnIdx = _ga.ColumnIdx - 1
		_gc.Column = IndexToColumn(_gc.ColumnIdx)
		return _gc
	default:
		return _ga
	}
}

// Update updates reference to point one of the neighboring columns with respect to the update type after removing a row/column.
func (_gcg *ColumnReference) Update(updateType _a.UpdateAction) *ColumnReference {
	switch updateType {
	case _a.UpdateActionRemoveColumn:
		_da := _gcg
		_da.ColumnIdx = _gcg.ColumnIdx - 1
		_da.Column = IndexToColumn(_da.ColumnIdx)
		return _da
	default:
		return _gcg
	}
}
