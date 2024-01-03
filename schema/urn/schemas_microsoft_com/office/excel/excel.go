package excel

import (
	_g "encoding/xml"
	_b "fmt"
	_gf "github.com/arcpd/msword"
	_bg "github.com/arcpd/msword/common/logger"
	_c "github.com/arcpd/msword/schema/soo/ofc/sharedTypes"
)

func (_cgea *ST_ObjectType) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_cgea = 0
	case "\u0042\u0075\u0074\u0074\u006f\u006e":
		*_cgea = 1
	case "\u0043\u0068\u0065\u0063\u006b\u0062\u006f\u0078":
		*_cgea = 2
	case "\u0044\u0069\u0061\u006c\u006f\u0067":
		*_cgea = 3
	case "\u0044\u0072\u006f\u0070":
		*_cgea = 4
	case "\u0045\u0064\u0069\u0074":
		*_cgea = 5
	case "\u0047\u0042\u006f\u0078":
		*_cgea = 6
	case "\u004c\u0061\u0062e\u006c":
		*_cgea = 7
	case "\u004c\u0069\u006ee\u0041":
		*_cgea = 8
	case "\u004c\u0069\u0073\u0074":
		*_cgea = 9
	case "\u004d\u006f\u0076i\u0065":
		*_cgea = 10
	case "\u004e\u006f\u0074\u0065":
		*_cgea = 11
	case "\u0050\u0069\u0063\u0074":
		*_cgea = 12
	case "\u0052\u0061\u0064i\u006f":
		*_cgea = 13
	case "\u0052\u0065\u0063t\u0041":
		*_cgea = 14
	case "\u0053\u0063\u0072\u006f\u006c\u006c":
		*_cgea = 15
	case "\u0053\u0070\u0069\u006e":
		*_cgea = 16
	case "\u0053\u0068\u0061p\u0065":
		*_cgea = 17
	case "\u0047\u0072\u006fu\u0070":
		*_cgea = 18
	case "\u0052\u0065\u0063\u0074":
		*_cgea = 19
	}
	return nil
}
func (_egff *ST_ObjectType) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_aaf, _dbfc := d.Token()
	if _dbfc != nil {
		return _dbfc
	}
	if _egc, _afgg := _aaf.(_g.EndElement); _afgg && _egc.Name == start.Name {
		*_egff = 1
		return nil
	}
	if _ccf, _bgdc := _aaf.(_g.CharData); !_bgdc {
		return _b.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _aaf)
	} else {
		switch string(_ccf) {
		case "":
			*_egff = 0
		case "\u0042\u0075\u0074\u0074\u006f\u006e":
			*_egff = 1
		case "\u0043\u0068\u0065\u0063\u006b\u0062\u006f\u0078":
			*_egff = 2
		case "\u0044\u0069\u0061\u006c\u006f\u0067":
			*_egff = 3
		case "\u0044\u0072\u006f\u0070":
			*_egff = 4
		case "\u0045\u0064\u0069\u0074":
			*_egff = 5
		case "\u0047\u0042\u006f\u0078":
			*_egff = 6
		case "\u004c\u0061\u0062e\u006c":
			*_egff = 7
		case "\u004c\u0069\u006ee\u0041":
			*_egff = 8
		case "\u004c\u0069\u0073\u0074":
			*_egff = 9
		case "\u004d\u006f\u0076i\u0065":
			*_egff = 10
		case "\u004e\u006f\u0074\u0065":
			*_egff = 11
		case "\u0050\u0069\u0063\u0074":
			*_egff = 12
		case "\u0052\u0061\u0064i\u006f":
			*_egff = 13
		case "\u0052\u0065\u0063t\u0041":
			*_egff = 14
		case "\u0053\u0063\u0072\u006f\u006c\u006c":
			*_egff = 15
		case "\u0053\u0070\u0069\u006e":
			*_egff = 16
		case "\u0053\u0068\u0061p\u0065":
			*_egff = 17
		case "\u0047\u0072\u006fu\u0070":
			*_egff = 18
		case "\u0052\u0065\u0063\u0074":
			*_egff = 19
		}
	}
	_aaf, _dbfc = d.Token()
	if _dbfc != nil {
		return _dbfc
	}
	if _dgf, _fgad := _aaf.(_g.EndElement); _fgad && _dgf.Name == start.Name {
		return nil
	}
	return _b.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _aaf)
}

type ClientData struct{ CT_ClientData }

func (_egg ST_ObjectType) ValidateWithPath(path string) error {
	switch _egg {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		return _b.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_egg))
	}
	return nil
}
func (_dbb ST_ObjectType) String() string {
	switch _dbb {
	case 0:
		return ""
	case 1:
		return "\u0042\u0075\u0074\u0074\u006f\u006e"
	case 2:
		return "\u0043\u0068\u0065\u0063\u006b\u0062\u006f\u0078"
	case 3:
		return "\u0044\u0069\u0061\u006c\u006f\u0067"
	case 4:
		return "\u0044\u0072\u006f\u0070"
	case 5:
		return "\u0045\u0064\u0069\u0074"
	case 6:
		return "\u0047\u0042\u006f\u0078"
	case 7:
		return "\u004c\u0061\u0062e\u006c"
	case 8:
		return "\u004c\u0069\u006ee\u0041"
	case 9:
		return "\u004c\u0069\u0073\u0074"
	case 10:
		return "\u004d\u006f\u0076i\u0065"
	case 11:
		return "\u004e\u006f\u0074\u0065"
	case 12:
		return "\u0050\u0069\u0063\u0074"
	case 13:
		return "\u0052\u0061\u0064i\u006f"
	case 14:
		return "\u0052\u0065\u0063t\u0041"
	case 15:
		return "\u0053\u0063\u0072\u006f\u006c\u006c"
	case 16:
		return "\u0053\u0070\u0069\u006e"
	case 17:
		return "\u0053\u0068\u0061p\u0065"
	case 18:
		return "\u0047\u0072\u006fu\u0070"
	case 19:
		return "\u0052\u0065\u0063\u0074"
	}
	return ""
}

// Validate validates the ClientData and its children
func (_dbc *ClientData) Validate() error {
	return _dbc.ValidateWithPath("\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061")
}
func (_bbfd ST_ObjectType) Validate() error { return _bbfd.ValidateWithPath("") }
func (_bcd ST_ObjectType) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_bcd.String(), start)
}
func (_dea ST_ObjectType) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_eag := _g.Attr{}
	_eag.Name = name
	switch _dea {
	case ST_ObjectTypeUnset:
		_eag.Value = ""
	case ST_ObjectTypeButton:
		_eag.Value = "\u0042\u0075\u0074\u0074\u006f\u006e"
	case ST_ObjectTypeCheckbox:
		_eag.Value = "\u0043\u0068\u0065\u0063\u006b\u0062\u006f\u0078"
	case ST_ObjectTypeDialog:
		_eag.Value = "\u0044\u0069\u0061\u006c\u006f\u0067"
	case ST_ObjectTypeDrop:
		_eag.Value = "\u0044\u0072\u006f\u0070"
	case ST_ObjectTypeEdit:
		_eag.Value = "\u0045\u0064\u0069\u0074"
	case ST_ObjectTypeGBox:
		_eag.Value = "\u0047\u0042\u006f\u0078"
	case ST_ObjectTypeLabel:
		_eag.Value = "\u004c\u0061\u0062e\u006c"
	case ST_ObjectTypeLineA:
		_eag.Value = "\u004c\u0069\u006ee\u0041"
	case ST_ObjectTypeList:
		_eag.Value = "\u004c\u0069\u0073\u0074"
	case ST_ObjectTypeMovie:
		_eag.Value = "\u004d\u006f\u0076i\u0065"
	case ST_ObjectTypeNote:
		_eag.Value = "\u004e\u006f\u0074\u0065"
	case ST_ObjectTypePict:
		_eag.Value = "\u0050\u0069\u0063\u0074"
	case ST_ObjectTypeRadio:
		_eag.Value = "\u0052\u0061\u0064i\u006f"
	case ST_ObjectTypeRectA:
		_eag.Value = "\u0052\u0065\u0063t\u0041"
	case ST_ObjectTypeScroll:
		_eag.Value = "\u0053\u0063\u0072\u006f\u006c\u006c"
	case ST_ObjectTypeSpin:
		_eag.Value = "\u0053\u0070\u0069\u006e"
	case ST_ObjectTypeShape:
		_eag.Value = "\u0053\u0068\u0061p\u0065"
	case ST_ObjectTypeGroup:
		_eag.Value = "\u0047\u0072\u006fu\u0070"
	case ST_ObjectTypeRect:
		_eag.Value = "\u0052\u0065\u0063\u0074"
	}
	return _eag, nil
}

const (
	ST_ObjectTypeUnset    ST_ObjectType = 0
	ST_ObjectTypeButton   ST_ObjectType = 1
	ST_ObjectTypeCheckbox ST_ObjectType = 2
	ST_ObjectTypeDialog   ST_ObjectType = 3
	ST_ObjectTypeDrop     ST_ObjectType = 4
	ST_ObjectTypeEdit     ST_ObjectType = 5
	ST_ObjectTypeGBox     ST_ObjectType = 6
	ST_ObjectTypeLabel    ST_ObjectType = 7
	ST_ObjectTypeLineA    ST_ObjectType = 8
	ST_ObjectTypeList     ST_ObjectType = 9
	ST_ObjectTypeMovie    ST_ObjectType = 10
	ST_ObjectTypeNote     ST_ObjectType = 11
	ST_ObjectTypePict     ST_ObjectType = 12
	ST_ObjectTypeRadio    ST_ObjectType = 13
	ST_ObjectTypeRectA    ST_ObjectType = 14
	ST_ObjectTypeScroll   ST_ObjectType = 15
	ST_ObjectTypeSpin     ST_ObjectType = 16
	ST_ObjectTypeShape    ST_ObjectType = 17
	ST_ObjectTypeGroup    ST_ObjectType = 18
	ST_ObjectTypeRect     ST_ObjectType = 19
)

func NewClientData() *ClientData {
	_egb := &ClientData{}
	_egb.CT_ClientData = *NewCT_ClientData()
	return _egb
}
func (_dcf *ClientData) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_dcf.CT_ClientData = *NewCT_ClientData()
	for _, _afef := range start.Attr {
		if _afef.Name.Local == "\u004f\u0062\u006a\u0065\u0063\u0074\u0054\u0079\u0070\u0065" {
			_dcf.ObjectTypeAttr.UnmarshalXMLAttr(_afef)
			continue
		}
	}
_fbed:
	for {
		_fdd, _fcac := d.Token()
		if _fcac != nil {
			return _fcac
		}
		switch _dab := _fdd.(type) {
		case _g.StartElement:
			switch _dab.Name {
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u006f\u0076\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}:
				_dcf.MoveWithCells = _c.ST_TrueFalseBlankUnset
				if _agb := d.DecodeElement(&_dcf.MoveWithCells, &_dab); _agb != nil {
					return _agb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0069\u007a\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}:
				_dcf.SizeWithCells = _c.ST_TrueFalseBlankUnset
				if _bgge := d.DecodeElement(&_dcf.SizeWithCells, &_dab); _bgge != nil {
					return _bgge
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_dcf.Anchor = new(string)
				if _bgdd := d.DecodeElement(_dcf.Anchor, &_dab); _bgdd != nil {
					return _bgdd
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u006f\u0063\u006b\u0065\u0064"}:
				_dcf.Locked = _c.ST_TrueFalseBlankUnset
				if _cbdg := d.DecodeElement(&_dcf.Locked, &_dab); _cbdg != nil {
					return _cbdg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "D\u0065\u0066\u0061\u0075\u006c\u0074\u0053\u0069\u007a\u0065"}:
				_dcf.DefaultSize = _c.ST_TrueFalseBlankUnset
				if _ceg := d.DecodeElement(&_dcf.DefaultSize, &_dab); _ceg != nil {
					return _ceg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "P\u0072\u0069\u006e\u0074\u004f\u0062\u006a\u0065\u0063\u0074"}:
				_dcf.PrintObject = _c.ST_TrueFalseBlankUnset
				if _fcba := d.DecodeElement(&_dcf.PrintObject, &_dab); _fcba != nil {
					return _fcba
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0069\u0073\u0061\u0062\u006c\u0065\u0064"}:
				_dcf.Disabled = _c.ST_TrueFalseBlankUnset
				if _fecf := d.DecodeElement(&_dcf.Disabled, &_dab); _fecf != nil {
					return _fecf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u0046\u0069\u006c\u006c"}:
				_dcf.AutoFill = _c.ST_TrueFalseBlankUnset
				if _edc := d.DecodeElement(&_dcf.AutoFill, &_dab); _edc != nil {
					return _edc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u004c\u0069\u006e\u0065"}:
				_dcf.AutoLine = _c.ST_TrueFalseBlankUnset
				if _gede := d.DecodeElement(&_dcf.AutoLine, &_dab); _gede != nil {
					return _gede
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u0050\u0069\u0063\u0074"}:
				_dcf.AutoPict = _c.ST_TrueFalseBlankUnset
				if _acea := d.DecodeElement(&_dcf.AutoPict, &_dab); _acea != nil {
					return _acea
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u004d\u0061\u0063\u0072o"}:
				_dcf.FmlaMacro = new(string)
				if _bdf := d.DecodeElement(_dcf.FmlaMacro, &_dab); _bdf != nil {
					return _bdf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0054\u0065\u0078\u0074\u0048\u0041\u006c\u0069\u0067\u006e"}:
				_dcf.TextHAlign = new(string)
				if _ffa := d.DecodeElement(_dcf.TextHAlign, &_dab); _ffa != nil {
					return _ffa
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0054\u0065\u0078\u0074\u0056\u0041\u006c\u0069\u0067\u006e"}:
				_dcf.TextVAlign = new(string)
				if _gag := d.DecodeElement(_dcf.TextVAlign, &_dab); _gag != nil {
					return _gag
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u006f\u0063\u006b\u0054\u0065\u0078\u0074"}:
				_dcf.LockText = _c.ST_TrueFalseBlankUnset
				if _cfge := d.DecodeElement(&_dcf.LockText, &_dab); _cfge != nil {
					return _cfge
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004au\u0073\u0074\u004c\u0061\u0073\u0074X"}:
				_dcf.JustLastX = _c.ST_TrueFalseBlankUnset
				if _eaa := d.DecodeElement(&_dcf.JustLastX, &_dab); _eaa != nil {
					return _eaa
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0065\u0063\u0072\u0065\u0074\u0045\u0064\u0069\u0074"}:
				_dcf.SecretEdit = _c.ST_TrueFalseBlankUnset
				if _bcad := d.DecodeElement(&_dcf.SecretEdit, &_dab); _bcad != nil {
					return _bcad
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_dcf.Default = _c.ST_TrueFalseBlankUnset
				if _dfe := d.DecodeElement(&_dcf.Default, &_dab); _dfe != nil {
					return _dfe
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0048\u0065\u006c\u0070"}:
				_dcf.Help = _c.ST_TrueFalseBlankUnset
				if _fccf := d.DecodeElement(&_dcf.Help, &_dab); _fccf != nil {
					return _fccf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0061\u006e\u0063\u0065\u006c"}:
				_dcf.Cancel = _c.ST_TrueFalseBlankUnset
				if _fgb := d.DecodeElement(&_dcf.Cancel, &_dab); _fgb != nil {
					return _fgb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044i\u0073\u006d\u0069\u0073\u0073"}:
				_dcf.Dismiss = _c.ST_TrueFalseBlankUnset
				if _fcad := d.DecodeElement(&_dcf.Dismiss, &_dab); _fcad != nil {
					return _fcad
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0063\u0063e\u006c"}:
				_dcf.Accel = new(int64)
				if _ead := d.DecodeElement(_dcf.Accel, &_dab); _ead != nil {
					return _ead
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0063\u0063\u0065\u006c\u0032"}:
				_dcf.Accel2 = new(int64)
				if _aaee := d.DecodeElement(_dcf.Accel2, &_dab); _aaee != nil {
					return _aaee
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052\u006f\u0077"}:
				_dcf.Row = new(int64)
				if _agd := d.DecodeElement(_dcf.Row, &_dab); _agd != nil {
					return _agd
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u006f\u006c\u0075\u006d\u006e"}:
				_dcf.Column = new(int64)
				if _afdd := d.DecodeElement(_dcf.Column, &_dab); _afdd != nil {
					return _afdd
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056i\u0073\u0069\u0062\u006c\u0065"}:
				_dcf.Visible = _c.ST_TrueFalseBlankUnset
				if _egf := d.DecodeElement(&_dcf.Visible, &_dab); _egf != nil {
					return _egf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052o\u0077\u0048\u0069\u0064\u0064\u0065n"}:
				_dcf.RowHidden = _c.ST_TrueFalseBlankUnset
				if _dfb := d.DecodeElement(&_dcf.RowHidden, &_dab); _dfb != nil {
					return _dfb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043o\u006c\u0048\u0069\u0064\u0064\u0065n"}:
				_dcf.ColHidden = _c.ST_TrueFalseBlankUnset
				if _ddb := d.DecodeElement(&_dcf.ColHidden, &_dab); _ddb != nil {
					return _ddb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0054\u0045\u0064\u0069\u0074"}:
				_dcf.VTEdit = new(int64)
				if _ede := d.DecodeElement(_dcf.VTEdit, &_dab); _ede != nil {
					return _ede
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004du\u006c\u0074\u0069\u004c\u0069\u006ee"}:
				_dcf.MultiLine = _c.ST_TrueFalseBlankUnset
				if _cgc := d.DecodeElement(&_dcf.MultiLine, &_dab); _cgc != nil {
					return _cgc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056S\u0063\u0072\u006f\u006c\u006c"}:
				_dcf.VScroll = _c.ST_TrueFalseBlankUnset
				if _fabe := d.DecodeElement(&_dcf.VScroll, &_dab); _fabe != nil {
					return _fabe
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0061\u006c\u0069\u0064\u0049\u0064\u0073"}:
				_dcf.ValidIds = _c.ST_TrueFalseBlankUnset
				if _aeb := d.DecodeElement(&_dcf.ValidIds, &_dab); _aeb != nil {
					return _aeb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u0052\u0061\u006e\u0067e"}:
				_dcf.FmlaRange = new(string)
				if _eec := d.DecodeElement(_dcf.FmlaRange, &_dab); _eec != nil {
					return _eec
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0057\u0069\u0064\u0074\u0068\u004d\u0069\u006e"}:
				_dcf.WidthMin = new(int64)
				if _aebc := d.DecodeElement(_dcf.WidthMin, &_dab); _aebc != nil {
					return _aebc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0065\u006c"}:
				_dcf.Sel = new(int64)
				if _efda := d.DecodeElement(_dcf.Sel, &_dab); _efda != nil {
					return _efda
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004eo\u0054\u0068\u0072\u0065\u0065\u00442"}:
				_dcf.NoThreeD2 = _c.ST_TrueFalseBlankUnset
				if _cba := d.DecodeElement(&_dcf.NoThreeD2, &_dab); _cba != nil {
					return _cba
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053e\u006c\u0054\u0079\u0070\u0065"}:
				_dcf.SelType = new(string)
				if _eed := d.DecodeElement(_dcf.SelType, &_dab); _eed != nil {
					return _eed
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0075\u006c\u0074\u0069\u0053\u0065\u006c"}:
				_dcf.MultiSel = new(string)
				if _fcbd := d.DecodeElement(_dcf.MultiSel, &_dab); _fcbd != nil {
					return _fcbd
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u0043\u0054"}:
				_dcf.LCT = new(string)
				if _cbc := d.DecodeElement(_dcf.LCT, &_dab); _cbc != nil {
					return _cbc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u0069\u0073\u0074\u0049\u0074\u0065\u006d"}:
				_dcf.ListItem = new(string)
				if _ggg := d.DecodeElement(_dcf.ListItem, &_dab); _ggg != nil {
					return _ggg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044r\u006f\u0070\u0053\u0074\u0079\u006ce"}:
				_dcf.DropStyle = new(string)
				if _ggge := d.DecodeElement(_dcf.DropStyle, &_dab); _ggge != nil {
					return _ggge
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043o\u006c\u006f\u0072\u0065\u0064"}:
				_dcf.Colored = _c.ST_TrueFalseBlankUnset
				if _bae := d.DecodeElement(&_dcf.Colored, &_dab); _bae != nil {
					return _bae
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044r\u006f\u0070\u004c\u0069\u006e\u0065s"}:
				_dcf.DropLines = new(int64)
				if _dfcb := d.DecodeElement(_dcf.DropLines, &_dab); _dfcb != nil {
					return _dfcb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043h\u0065\u0063\u006b\u0065\u0064"}:
				_dcf.Checked = new(int64)
				if _agg := d.DecodeElement(_dcf.Checked, &_dab); _agg != nil {
					return _agg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u004c\u0069\u006e\u006b"}:
				_dcf.FmlaLink = new(string)
				if _def := d.DecodeElement(_dcf.FmlaLink, &_dab); _def != nil {
					return _def
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u0050\u0069\u0063\u0074"}:
				_dcf.FmlaPict = new(string)
				if _dgb := d.DecodeElement(_dcf.FmlaPict, &_dab); _dgb != nil {
					return _dgb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044"}:
				_dcf.NoThreeD = _c.ST_TrueFalseBlankUnset
				if _baf := d.DecodeElement(&_dcf.NoThreeD, &_dab); _baf != nil {
					return _baf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "F\u0069\u0072\u0073\u0074\u0042\u0075\u0074\u0074\u006f\u006e"}:
				_dcf.FirstButton = _c.ST_TrueFalseBlankUnset
				if _bda := d.DecodeElement(&_dcf.FirstButton, &_dab); _bda != nil {
					return _bda
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u0047\u0072\u006f\u0075p"}:
				_dcf.FmlaGroup = new(string)
				if _bfc := d.DecodeElement(_dcf.FmlaGroup, &_dab); _bfc != nil {
					return _bfc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0061\u006c"}:
				_dcf.Val = new(int64)
				if _decf := d.DecodeElement(_dcf.Val, &_dab); _decf != nil {
					return _decf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0069\u006e"}:
				_dcf.Min = new(int64)
				if _ffc := d.DecodeElement(_dcf.Min, &_dab); _ffc != nil {
					return _ffc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0061\u0078"}:
				_dcf.Max = new(int64)
				if _cadc := d.DecodeElement(_dcf.Max, &_dab); _cadc != nil {
					return _cadc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0049\u006e\u0063"}:
				_dcf.Inc = new(int64)
				if _ece := d.DecodeElement(_dcf.Inc, &_dab); _ece != nil {
					return _ece
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0050\u0061\u0067\u0065"}:
				_dcf.Page = new(int64)
				if _eceb := d.DecodeElement(_dcf.Page, &_dab); _eceb != nil {
					return _eceb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0048\u006f\u0072i\u007a"}:
				_dcf.Horiz = _c.ST_TrueFalseBlankUnset
				if _eac := d.DecodeElement(&_dcf.Horiz, &_dab); _eac != nil {
					return _eac
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0078"}:
				_dcf.Dx = new(int64)
				if _dfde := d.DecodeElement(_dcf.Dx, &_dab); _dfde != nil {
					return _dfde
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0061\u0070\u004f\u0043\u0058"}:
				_dcf.MapOCX = _c.ST_TrueFalseBlankUnset
				if _bdg := d.DecodeElement(&_dcf.MapOCX, &_dab); _bdg != nil {
					return _bdg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0046"}:
				var _edf string
				if _fgg := d.DecodeElement(&_edf, &_dab); _fgg != nil {
					return _fgg
				}
				_dcf.CF = append(_dcf.CF, _edf)
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0061\u006d\u0065\u0072\u0061"}:
				_dcf.Camera = _c.ST_TrueFalseBlankUnset
				if _dba := d.DecodeElement(&_dcf.Camera, &_dab); _dba != nil {
					return _dba
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052\u0065\u0063a\u006c\u0063\u0041\u006c\u0077\u0061\u0079\u0073"}:
				_dcf.RecalcAlways = _c.ST_TrueFalseBlankUnset
				if _fabeb := d.DecodeElement(&_dcf.RecalcAlways, &_dab); _fabeb != nil {
					return _fabeb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041u\u0074\u006f\u0053\u0063\u0061\u006ce"}:
				_dcf.AutoScale = _c.ST_TrueFalseBlankUnset
				if _fac := d.DecodeElement(&_dcf.AutoScale, &_dab); _fac != nil {
					return _fac
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0044\u0045"}:
				_dcf.DDE = _c.ST_TrueFalseBlankUnset
				if _afga := d.DecodeElement(&_dcf.DDE, &_dab); _afga != nil {
					return _afga
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0055\u0049\u004fb\u006a"}:
				_dcf.UIObj = _c.ST_TrueFalseBlankUnset
				if _cgca := d.DecodeElement(&_dcf.UIObj, &_dab); _cgca != nil {
					return _cgca
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u0054\u0065\u0078\u0074"}:
				_dcf.ScriptText = new(string)
				if _dca := d.DecodeElement(_dcf.ScriptText, &_dab); _dca != nil {
					return _dca
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u0045\u0078\u0074e\u006e\u0064\u0065\u0064"}:
				_dcf.ScriptExtended = new(string)
				if _cc := d.DecodeElement(_dcf.ScriptExtended, &_dab); _cc != nil {
					return _cc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"}:
				_dcf.ScriptLanguage = new(uint32)
				if _gaca := d.DecodeElement(_dcf.ScriptLanguage, &_dab); _gaca != nil {
					return _gaca
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"}:
				_dcf.ScriptLocation = new(uint32)
				if _ffb := d.DecodeElement(_dcf.ScriptLocation, &_dab); _ffb != nil {
					return _ffb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u0054\u0078\u0062\u0078"}:
				_dcf.FmlaTxbx = new(string)
				if _ebg := d.DecodeElement(_dcf.FmlaTxbx, &_dab); _ebg != nil {
					return _ebg
				}
			default:
				_bg.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u006c\u0069e\u006e\u0074\u0044\u0061\u0074\u0061\u0020\u0025\u0076", _dab.Name)
				if _dgd := d.Skip(); _dgd != nil {
					return _dgd
				}
			}
		case _g.EndElement:
			break _fbed
		case _g.CharData:
		}
	}
	return nil
}
func (_cdd *CT_ClientData) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cdd.ObjectTypeAttr = ST_ObjectType(1)
	for _, _eea := range start.Attr {
		if _eea.Name.Local == "\u004f\u0062\u006a\u0065\u0063\u0074\u0054\u0079\u0070\u0065" {
			_cdd.ObjectTypeAttr.UnmarshalXMLAttr(_eea)
			continue
		}
	}
_db:
	for {
		_daf, _gffd := d.Token()
		if _gffd != nil {
			return _gffd
		}
		switch _bgb := _daf.(type) {
		case _g.StartElement:
			switch _bgb.Name {
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u006f\u0076\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}:
				_cdd.MoveWithCells = _c.ST_TrueFalseBlankUnset
				if _fgeb := d.DecodeElement(&_cdd.MoveWithCells, &_bgb); _fgeb != nil {
					return _fgeb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0069\u007a\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}:
				_cdd.SizeWithCells = _c.ST_TrueFalseBlankUnset
				if _efe := d.DecodeElement(&_cdd.SizeWithCells, &_bgb); _efe != nil {
					return _efe
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_cdd.Anchor = new(string)
				if _dfa := d.DecodeElement(_cdd.Anchor, &_bgb); _dfa != nil {
					return _dfa
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u006f\u0063\u006b\u0065\u0064"}:
				_cdd.Locked = _c.ST_TrueFalseBlankUnset
				if _acd := d.DecodeElement(&_cdd.Locked, &_bgb); _acd != nil {
					return _acd
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "D\u0065\u0066\u0061\u0075\u006c\u0074\u0053\u0069\u007a\u0065"}:
				_cdd.DefaultSize = _c.ST_TrueFalseBlankUnset
				if _ce := d.DecodeElement(&_cdd.DefaultSize, &_bgb); _ce != nil {
					return _ce
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "P\u0072\u0069\u006e\u0074\u004f\u0062\u006a\u0065\u0063\u0074"}:
				_cdd.PrintObject = _c.ST_TrueFalseBlankUnset
				if _gcg := d.DecodeElement(&_cdd.PrintObject, &_bgb); _gcg != nil {
					return _gcg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0069\u0073\u0061\u0062\u006c\u0065\u0064"}:
				_cdd.Disabled = _c.ST_TrueFalseBlankUnset
				if _aae := d.DecodeElement(&_cdd.Disabled, &_bgb); _aae != nil {
					return _aae
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u0046\u0069\u006c\u006c"}:
				_cdd.AutoFill = _c.ST_TrueFalseBlankUnset
				if _fab := d.DecodeElement(&_cdd.AutoFill, &_bgb); _fab != nil {
					return _fab
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u004c\u0069\u006e\u0065"}:
				_cdd.AutoLine = _c.ST_TrueFalseBlankUnset
				if _gfc := d.DecodeElement(&_cdd.AutoLine, &_bgb); _gfc != nil {
					return _gfc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0075\u0074\u006f\u0050\u0069\u0063\u0074"}:
				_cdd.AutoPict = _c.ST_TrueFalseBlankUnset
				if _ad := d.DecodeElement(&_cdd.AutoPict, &_bgb); _ad != nil {
					return _ad
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u004d\u0061\u0063\u0072o"}:
				_cdd.FmlaMacro = new(string)
				if _acb := d.DecodeElement(_cdd.FmlaMacro, &_bgb); _acb != nil {
					return _acb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0054\u0065\u0078\u0074\u0048\u0041\u006c\u0069\u0067\u006e"}:
				_cdd.TextHAlign = new(string)
				if _ecb := d.DecodeElement(_cdd.TextHAlign, &_bgb); _ecb != nil {
					return _ecb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0054\u0065\u0078\u0074\u0056\u0041\u006c\u0069\u0067\u006e"}:
				_cdd.TextVAlign = new(string)
				if _ed := d.DecodeElement(_cdd.TextVAlign, &_bgb); _ed != nil {
					return _ed
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u006f\u0063\u006b\u0054\u0065\u0078\u0074"}:
				_cdd.LockText = _c.ST_TrueFalseBlankUnset
				if _gec := d.DecodeElement(&_cdd.LockText, &_bgb); _gec != nil {
					return _gec
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004au\u0073\u0074\u004c\u0061\u0073\u0074X"}:
				_cdd.JustLastX = _c.ST_TrueFalseBlankUnset
				if _ag := d.DecodeElement(&_cdd.JustLastX, &_bgb); _ag != nil {
					return _ag
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0065\u0063\u0072\u0065\u0074\u0045\u0064\u0069\u0074"}:
				_cdd.SecretEdit = _c.ST_TrueFalseBlankUnset
				if _eeb := d.DecodeElement(&_cdd.SecretEdit, &_bgb); _eeb != nil {
					return _eeb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_cdd.Default = _c.ST_TrueFalseBlankUnset
				if _afg := d.DecodeElement(&_cdd.Default, &_bgb); _afg != nil {
					return _afg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0048\u0065\u006c\u0070"}:
				_cdd.Help = _c.ST_TrueFalseBlankUnset
				if _gea := d.DecodeElement(&_cdd.Help, &_bgb); _gea != nil {
					return _gea
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0061\u006e\u0063\u0065\u006c"}:
				_cdd.Cancel = _c.ST_TrueFalseBlankUnset
				if _gbc := d.DecodeElement(&_cdd.Cancel, &_bgb); _gbc != nil {
					return _gbc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044i\u0073\u006d\u0069\u0073\u0073"}:
				_cdd.Dismiss = _c.ST_TrueFalseBlankUnset
				if _fcca := d.DecodeElement(&_cdd.Dismiss, &_bgb); _fcca != nil {
					return _fcca
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0063\u0063e\u006c"}:
				_cdd.Accel = new(int64)
				if _dbg := d.DecodeElement(_cdd.Accel, &_bgb); _dbg != nil {
					return _dbg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041\u0063\u0063\u0065\u006c\u0032"}:
				_cdd.Accel2 = new(int64)
				if _dcc := d.DecodeElement(_cdd.Accel2, &_bgb); _dcc != nil {
					return _dcc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052\u006f\u0077"}:
				_cdd.Row = new(int64)
				if _eccg := d.DecodeElement(_cdd.Row, &_bgb); _eccg != nil {
					return _eccg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u006f\u006c\u0075\u006d\u006e"}:
				_cdd.Column = new(int64)
				if _afe := d.DecodeElement(_cdd.Column, &_bgb); _afe != nil {
					return _afe
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056i\u0073\u0069\u0062\u006c\u0065"}:
				_cdd.Visible = _c.ST_TrueFalseBlankUnset
				if _acg := d.DecodeElement(&_cdd.Visible, &_bgb); _acg != nil {
					return _acg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052o\u0077\u0048\u0069\u0064\u0064\u0065n"}:
				_cdd.RowHidden = _c.ST_TrueFalseBlankUnset
				if _cbg := d.DecodeElement(&_cdd.RowHidden, &_bgb); _cbg != nil {
					return _cbg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043o\u006c\u0048\u0069\u0064\u0064\u0065n"}:
				_cdd.ColHidden = _c.ST_TrueFalseBlankUnset
				if _bce := d.DecodeElement(&_cdd.ColHidden, &_bgb); _bce != nil {
					return _bce
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0054\u0045\u0064\u0069\u0074"}:
				_cdd.VTEdit = new(int64)
				if _afgd := d.DecodeElement(_cdd.VTEdit, &_bgb); _afgd != nil {
					return _afgd
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004du\u006c\u0074\u0069\u004c\u0069\u006ee"}:
				_cdd.MultiLine = _c.ST_TrueFalseBlankUnset
				if _fcb := d.DecodeElement(&_cdd.MultiLine, &_bgb); _fcb != nil {
					return _fcb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056S\u0063\u0072\u006f\u006c\u006c"}:
				_cdd.VScroll = _c.ST_TrueFalseBlankUnset
				if _ceb := d.DecodeElement(&_cdd.VScroll, &_bgb); _ceb != nil {
					return _ceb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0061\u006c\u0069\u0064\u0049\u0064\u0073"}:
				_cdd.ValidIds = _c.ST_TrueFalseBlankUnset
				if _fce := d.DecodeElement(&_cdd.ValidIds, &_bgb); _fce != nil {
					return _fce
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u0052\u0061\u006e\u0067e"}:
				_cdd.FmlaRange = new(string)
				if _bdca := d.DecodeElement(_cdd.FmlaRange, &_bgb); _bdca != nil {
					return _bdca
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0057\u0069\u0064\u0074\u0068\u004d\u0069\u006e"}:
				_cdd.WidthMin = new(int64)
				if _dce := d.DecodeElement(_cdd.WidthMin, &_bgb); _dce != nil {
					return _dce
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0065\u006c"}:
				_cdd.Sel = new(int64)
				if _ff := d.DecodeElement(_cdd.Sel, &_bgb); _ff != nil {
					return _ff
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004eo\u0054\u0068\u0072\u0065\u0065\u00442"}:
				_cdd.NoThreeD2 = _c.ST_TrueFalseBlankUnset
				if _bcf := d.DecodeElement(&_cdd.NoThreeD2, &_bgb); _bcf != nil {
					return _bcf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053e\u006c\u0054\u0079\u0070\u0065"}:
				_cdd.SelType = new(string)
				if _aff := d.DecodeElement(_cdd.SelType, &_bgb); _aff != nil {
					return _aff
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0075\u006c\u0074\u0069\u0053\u0065\u006c"}:
				_cdd.MultiSel = new(string)
				if _gbg := d.DecodeElement(_cdd.MultiSel, &_bgb); _gbg != nil {
					return _gbg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u0043\u0054"}:
				_cdd.LCT = new(string)
				if _gcd := d.DecodeElement(_cdd.LCT, &_bgb); _gcd != nil {
					return _gcd
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004c\u0069\u0073\u0074\u0049\u0074\u0065\u006d"}:
				_cdd.ListItem = new(string)
				if _eae := d.DecodeElement(_cdd.ListItem, &_bgb); _eae != nil {
					return _eae
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044r\u006f\u0070\u0053\u0074\u0079\u006ce"}:
				_cdd.DropStyle = new(string)
				if _gcbc := d.DecodeElement(_cdd.DropStyle, &_bgb); _gcbc != nil {
					return _gcbc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043o\u006c\u006f\u0072\u0065\u0064"}:
				_cdd.Colored = _c.ST_TrueFalseBlankUnset
				if _ada := d.DecodeElement(&_cdd.Colored, &_bgb); _ada != nil {
					return _ada
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044r\u006f\u0070\u004c\u0069\u006e\u0065s"}:
				_cdd.DropLines = new(int64)
				if _aag := d.DecodeElement(_cdd.DropLines, &_bgb); _aag != nil {
					return _aag
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043h\u0065\u0063\u006b\u0065\u0064"}:
				_cdd.Checked = new(int64)
				if _ade := d.DecodeElement(_cdd.Checked, &_bgb); _ade != nil {
					return _ade
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u004c\u0069\u006e\u006b"}:
				_cdd.FmlaLink = new(string)
				if _bge := d.DecodeElement(_cdd.FmlaLink, &_bgb); _bge != nil {
					return _bge
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u0050\u0069\u0063\u0074"}:
				_cdd.FmlaPict = new(string)
				if _fae := d.DecodeElement(_cdd.FmlaPict, &_bgb); _fae != nil {
					return _fae
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044"}:
				_cdd.NoThreeD = _c.ST_TrueFalseBlankUnset
				if _gadg := d.DecodeElement(&_cdd.NoThreeD, &_bgb); _gadg != nil {
					return _gadg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "F\u0069\u0072\u0073\u0074\u0042\u0075\u0074\u0074\u006f\u006e"}:
				_cdd.FirstButton = _c.ST_TrueFalseBlankUnset
				if _efd := d.DecodeElement(&_cdd.FirstButton, &_bgb); _efd != nil {
					return _efd
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046m\u006c\u0061\u0047\u0072\u006f\u0075p"}:
				_cdd.FmlaGroup = new(string)
				if _bcfg := d.DecodeElement(_cdd.FmlaGroup, &_bgb); _bcfg != nil {
					return _bcfg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0056\u0061\u006c"}:
				_cdd.Val = new(int64)
				if _gce := d.DecodeElement(_cdd.Val, &_bgb); _gce != nil {
					return _gce
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0069\u006e"}:
				_cdd.Min = new(int64)
				if _gdf := d.DecodeElement(_cdd.Min, &_bgb); _gdf != nil {
					return _gdf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0061\u0078"}:
				_cdd.Max = new(int64)
				if _cfb := d.DecodeElement(_cdd.Max, &_bgb); _cfb != nil {
					return _cfb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0049\u006e\u0063"}:
				_cdd.Inc = new(int64)
				if _gg := d.DecodeElement(_cdd.Inc, &_bgb); _gg != nil {
					return _gg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0050\u0061\u0067\u0065"}:
				_cdd.Page = new(int64)
				if _cfg := d.DecodeElement(_cdd.Page, &_bgb); _cfg != nil {
					return _cfg
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0048\u006f\u0072i\u007a"}:
				_cdd.Horiz = _c.ST_TrueFalseBlankUnset
				if _adc := d.DecodeElement(&_cdd.Horiz, &_bgb); _adc != nil {
					return _adc
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0078"}:
				_cdd.Dx = new(int64)
				if _fca := d.DecodeElement(_cdd.Dx, &_bgb); _fca != nil {
					return _fca
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u004d\u0061\u0070\u004f\u0043\u0058"}:
				_cdd.MapOCX = _c.ST_TrueFalseBlankUnset
				if _bbf := d.DecodeElement(&_cdd.MapOCX, &_bgb); _bbf != nil {
					return _bbf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0046"}:
				var _afd string
				if _dcd := d.DecodeElement(&_afd, &_bgb); _dcd != nil {
					return _dcd
				}
				_cdd.CF = append(_cdd.CF, _afd)
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0043\u0061\u006d\u0065\u0072\u0061"}:
				_cdd.Camera = _c.ST_TrueFalseBlankUnset
				if _cddf := d.DecodeElement(&_cdd.Camera, &_bgb); _cddf != nil {
					return _cddf
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0052\u0065\u0063a\u006c\u0063\u0041\u006c\u0077\u0061\u0079\u0073"}:
				_cdd.RecalcAlways = _c.ST_TrueFalseBlankUnset
				if _aage := d.DecodeElement(&_cdd.RecalcAlways, &_bgb); _aage != nil {
					return _aage
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0041u\u0074\u006f\u0053\u0063\u0061\u006ce"}:
				_cdd.AutoScale = _c.ST_TrueFalseBlankUnset
				if _bed := d.DecodeElement(&_cdd.AutoScale, &_bgb); _bed != nil {
					return _bed
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0044\u0044\u0045"}:
				_cdd.DDE = _c.ST_TrueFalseBlankUnset
				if _gcbd := d.DecodeElement(&_cdd.DDE, &_bgb); _gcbd != nil {
					return _gcbd
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0055\u0049\u004fb\u006a"}:
				_cdd.UIObj = _c.ST_TrueFalseBlankUnset
				if _baa := d.DecodeElement(&_cdd.UIObj, &_bgb); _baa != nil {
					return _baa
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u0054\u0065\u0078\u0074"}:
				_cdd.ScriptText = new(string)
				if _edb := d.DecodeElement(_cdd.ScriptText, &_bgb); _edb != nil {
					return _edb
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u0045\u0078\u0074e\u006e\u0064\u0065\u0064"}:
				_cdd.ScriptExtended = new(string)
				if _cfe := d.DecodeElement(_cdd.ScriptExtended, &_bgb); _cfe != nil {
					return _cfe
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"}:
				_cdd.ScriptLanguage = new(uint32)
				if _efa := d.DecodeElement(_cdd.ScriptLanguage, &_bgb); _efa != nil {
					return _efa
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0053\u0063\u0072\u0069\u0070\u0074\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"}:
				_cdd.ScriptLocation = new(uint32)
				if _ege := d.DecodeElement(_cdd.ScriptLocation, &_bgb); _ege != nil {
					return _ege
				}
			case _g.Name{Space: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", Local: "\u0046\u006d\u006c\u0061\u0054\u0078\u0062\u0078"}:
				_cdd.FmlaTxbx = new(string)
				if _afgf := d.DecodeElement(_cdd.FmlaTxbx, &_bgb); _afgf != nil {
					return _afgf
				}
			default:
				_bg.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043l\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061 \u0025\u0076", _bgb.Name)
				if _bgd := d.Skip(); _bgd != nil {
					return _bgd
				}
			}
		case _g.EndElement:
			break _db
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ClientData and its children
func (_bcea *CT_ClientData) Validate() error {
	return _bcea.ValidateWithPath("\u0043\u0054\u005f\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061")
}

type CT_ClientData struct {
	ObjectTypeAttr ST_ObjectType
	MoveWithCells  _c.ST_TrueFalseBlank
	SizeWithCells  _c.ST_TrueFalseBlank
	Anchor         *string
	Locked         _c.ST_TrueFalseBlank
	DefaultSize    _c.ST_TrueFalseBlank
	PrintObject    _c.ST_TrueFalseBlank
	Disabled       _c.ST_TrueFalseBlank
	AutoFill       _c.ST_TrueFalseBlank
	AutoLine       _c.ST_TrueFalseBlank
	AutoPict       _c.ST_TrueFalseBlank
	FmlaMacro      *string
	TextHAlign     *string
	TextVAlign     *string
	LockText       _c.ST_TrueFalseBlank
	JustLastX      _c.ST_TrueFalseBlank
	SecretEdit     _c.ST_TrueFalseBlank
	Default        _c.ST_TrueFalseBlank
	Help           _c.ST_TrueFalseBlank
	Cancel         _c.ST_TrueFalseBlank
	Dismiss        _c.ST_TrueFalseBlank
	Accel          *int64
	Accel2         *int64
	Row            *int64
	Column         *int64
	Visible        _c.ST_TrueFalseBlank
	RowHidden      _c.ST_TrueFalseBlank
	ColHidden      _c.ST_TrueFalseBlank
	VTEdit         *int64
	MultiLine      _c.ST_TrueFalseBlank
	VScroll        _c.ST_TrueFalseBlank
	ValidIds       _c.ST_TrueFalseBlank
	FmlaRange      *string
	WidthMin       *int64
	Sel            *int64
	NoThreeD2      _c.ST_TrueFalseBlank
	SelType        *string
	MultiSel       *string
	LCT            *string
	ListItem       *string
	DropStyle      *string
	Colored        _c.ST_TrueFalseBlank
	DropLines      *int64
	Checked        *int64
	FmlaLink       *string
	FmlaPict       *string
	NoThreeD       _c.ST_TrueFalseBlank
	FirstButton    _c.ST_TrueFalseBlank
	FmlaGroup      *string
	Val            *int64
	Min            *int64
	Max            *int64
	Inc            *int64
	Page           *int64
	Horiz          _c.ST_TrueFalseBlank
	Dx             *int64
	MapOCX         _c.ST_TrueFalseBlank
	CF             []string
	Camera         _c.ST_TrueFalseBlank
	RecalcAlways   _c.ST_TrueFalseBlank
	AutoScale      _c.ST_TrueFalseBlank
	DDE            _c.ST_TrueFalseBlank
	UIObj          _c.ST_TrueFalseBlank
	ScriptText     *string
	ScriptExtended *string
	ScriptLanguage *uint32
	ScriptLocation *uint32
	FmlaTxbx       *string
}

// ValidateWithPath validates the ClientData and its children, prefixing error messages with path
func (_cbdf *ClientData) ValidateWithPath(path string) error {
	if _ebc := _cbdf.CT_ClientData.ValidateWithPath(path); _ebc != nil {
		return _ebc
	}
	return nil
}
func (_ecf *ClientData) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0078\u003a\u0043l\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"
	return _ecf.CT_ClientData.MarshalXML(e, start)
}

type ST_ObjectType byte

func (_gb *CT_ClientData) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_fg, _e := _gb.ObjectTypeAttr.MarshalXMLAttr(_g.Name{Local: "\u004f\u0062\u006a\u0065\u0063\u0074\u0054\u0079\u0070\u0065"})
	if _e != nil {
		return _e
	}
	start.Attr = append(start.Attr, _fg)
	e.EncodeToken(start)
	if _gb.MoveWithCells != _c.ST_TrueFalseBlankUnset {
		_eg := _g.StartElement{Name: _g.Name{Local: "\u0078:\u004do\u0076\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}}
		e.EncodeElement(_gb.MoveWithCells, _eg)
	}
	if _gb.SizeWithCells != _c.ST_TrueFalseBlankUnset {
		_bgc := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0053i\u007a\u0065\u0057\u0069\u0074\u0068\u0043\u0065\u006c\u006c\u0073"}}
		e.EncodeElement(_gb.SizeWithCells, _bgc)
	}
	if _gb.Anchor != nil {
		_ea := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0041\u006e\u0063\u0068\u006f\u0072"}}
		_gf.AddPreserveSpaceAttr(&_ea, *_gb.Anchor)
		e.EncodeElement(_gb.Anchor, _ea)
	}
	if _gb.Locked != _c.ST_TrueFalseBlankUnset {
		_a := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u004c\u006f\u0063\u006b\u0065\u0064"}}
		e.EncodeElement(_gb.Locked, _a)
	}
	if _gb.DefaultSize != _c.ST_TrueFalseBlankUnset {
		_cb := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u0053\u0069\u007a\u0065"}}
		e.EncodeElement(_gb.DefaultSize, _cb)
	}
	if _gb.PrintObject != _c.ST_TrueFalseBlankUnset {
		_d := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0050\u0072\u0069\u006e\u0074\u004f\u0062\u006a\u0065\u0063\u0074"}}
		e.EncodeElement(_gb.PrintObject, _d)
	}
	if _gb.Disabled != _c.ST_TrueFalseBlankUnset {
		_ae := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0044\u0069\u0073\u0061\u0062\u006c\u0065\u0064"}}
		e.EncodeElement(_gb.Disabled, _ae)
	}
	if _gb.AutoFill != _c.ST_TrueFalseBlankUnset {
		_cd := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0041\u0075\u0074\u006f\u0046\u0069\u006c\u006c"}}
		e.EncodeElement(_gb.AutoFill, _cd)
	}
	if _gb.AutoLine != _c.ST_TrueFalseBlankUnset {
		_ga := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0041\u0075\u0074\u006f\u004c\u0069\u006e\u0065"}}
		e.EncodeElement(_gb.AutoLine, _ga)
	}
	if _gb.AutoPict != _c.ST_TrueFalseBlankUnset {
		_gff := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0041\u0075\u0074\u006f\u0050\u0069\u0063\u0074"}}
		e.EncodeElement(_gb.AutoPict, _gff)
	}
	if _gb.FmlaMacro != nil {
		_fc := _g.StartElement{Name: _g.Name{Local: "x\u003a\u0046\u006d\u006c\u0061\u004d\u0061\u0063\u0072\u006f"}}
		_gf.AddPreserveSpaceAttr(&_fc, *_gb.FmlaMacro)
		e.EncodeElement(_gb.FmlaMacro, _fc)
	}
	if _gb.TextHAlign != nil {
		_cbd := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0054e\u0078\u0074\u0048\u0041\u006c\u0069\u0067\u006e"}}
		_gf.AddPreserveSpaceAttr(&_cbd, *_gb.TextHAlign)
		e.EncodeElement(_gb.TextHAlign, _cbd)
	}
	if _gb.TextVAlign != nil {
		_be := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0054e\u0078\u0074\u0056\u0041\u006c\u0069\u0067\u006e"}}
		_gf.AddPreserveSpaceAttr(&_be, *_gb.TextVAlign)
		e.EncodeElement(_gb.TextVAlign, _be)
	}
	if _gb.LockText != _c.ST_TrueFalseBlankUnset {
		_ca := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u004c\u006f\u0063\u006b\u0054\u0065\u0078\u0074"}}
		e.EncodeElement(_gb.LockText, _ca)
	}
	if _gb.JustLastX != _c.ST_TrueFalseBlankUnset {
		_fe := _g.StartElement{Name: _g.Name{Local: "x\u003a\u004a\u0075\u0073\u0074\u004c\u0061\u0073\u0074\u0058"}}
		e.EncodeElement(_gb.JustLastX, _fe)
	}
	if _gb.SecretEdit != _c.ST_TrueFalseBlankUnset {
		_fge := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0053e\u0063\u0072\u0065\u0074\u0045\u0064\u0069\u0074"}}
		e.EncodeElement(_gb.SecretEdit, _fge)
	}
	if _gb.Default != _c.ST_TrueFalseBlankUnset {
		_dc := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0044\u0065\u0066\u0061\u0075\u006ct"}}
		e.EncodeElement(_gb.Default, _dc)
	}
	if _gb.Help != _c.ST_TrueFalseBlankUnset {
		_da := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0048\u0065\u006c\u0070"}}
		e.EncodeElement(_gb.Help, _da)
	}
	if _gb.Cancel != _c.ST_TrueFalseBlankUnset {
		_gbf := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0043\u0061\u006e\u0063\u0065\u006c"}}
		e.EncodeElement(_gb.Cancel, _gbf)
	}
	if _gb.Dismiss != _c.ST_TrueFalseBlankUnset {
		_gca := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0044\u0069\u0073\u006d\u0069\u0073s"}}
		e.EncodeElement(_gb.Dismiss, _gca)
	}
	if _gb.Accel != nil {
		_gaf := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0041\u0063\u0063\u0065\u006c"}}
		e.EncodeElement(_gb.Accel, _gaf)
	}
	if _gb.Accel2 != nil {
		_daa := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0041\u0063\u0063\u0065\u006c\u0032"}}
		e.EncodeElement(_gb.Accel2, _daa)
	}
	if _gb.Row != nil {
		_egd := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0052o\u0077"}}
		e.EncodeElement(_gb.Row, _egd)
	}
	if _gb.Column != nil {
		_aeg := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0043\u006f\u006c\u0075\u006d\u006e"}}
		e.EncodeElement(_gb.Column, _aeg)
	}
	if _gb.Visible != _c.ST_TrueFalseBlankUnset {
		_gad := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0056\u0069\u0073\u0069\u0062\u006ce"}}
		e.EncodeElement(_gb.Visible, _gad)
	}
	if _gb.RowHidden != _c.ST_TrueFalseBlankUnset {
		_daaa := _g.StartElement{Name: _g.Name{Local: "x\u003a\u0052\u006f\u0077\u0048\u0069\u0064\u0064\u0065\u006e"}}
		e.EncodeElement(_gb.RowHidden, _daaa)
	}
	if _gb.ColHidden != _c.ST_TrueFalseBlankUnset {
		_ee := _g.StartElement{Name: _g.Name{Local: "x\u003a\u0043\u006f\u006c\u0048\u0069\u0064\u0064\u0065\u006e"}}
		e.EncodeElement(_gb.ColHidden, _ee)
	}
	if _gb.VTEdit != nil {
		_cg := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0056\u0054\u0045\u0064\u0069\u0074"}}
		e.EncodeElement(_gb.VTEdit, _cg)
	}
	if _gb.MultiLine != _c.ST_TrueFalseBlankUnset {
		_aa := _g.StartElement{Name: _g.Name{Local: "x\u003a\u004d\u0075\u006c\u0074\u0069\u004c\u0069\u006e\u0065"}}
		e.EncodeElement(_gb.MultiLine, _aa)
	}
	if _gb.VScroll != _c.ST_TrueFalseBlankUnset {
		_gafe := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0056\u0053\u0063\u0072\u006f\u006cl"}}
		e.EncodeElement(_gb.VScroll, _gafe)
	}
	if _gb.ValidIds != _c.ST_TrueFalseBlankUnset {
		_aea := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0056\u0061\u006c\u0069\u0064\u0049\u0064\u0073"}}
		e.EncodeElement(_gb.ValidIds, _aea)
	}
	if _gb.FmlaRange != nil {
		_fb := _g.StartElement{Name: _g.Name{Local: "x\u003a\u0046\u006d\u006c\u0061\u0052\u0061\u006e\u0067\u0065"}}
		_gf.AddPreserveSpaceAttr(&_fb, *_gb.FmlaRange)
		e.EncodeElement(_gb.FmlaRange, _fb)
	}
	if _gb.WidthMin != nil {
		_bea := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0057\u0069\u0064\u0074\u0068\u004d\u0069\u006e"}}
		e.EncodeElement(_gb.WidthMin, _bea)
	}
	if _gb.Sel != nil {
		_gcb := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0053e\u006c"}}
		e.EncodeElement(_gb.Sel, _gcb)
	}
	if _gb.NoThreeD2 != _c.ST_TrueFalseBlankUnset {
		_bgf := _g.StartElement{Name: _g.Name{Local: "x\u003a\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044\u0032"}}
		e.EncodeElement(_gb.NoThreeD2, _bgf)
	}
	if _gb.SelType != nil {
		_gaa := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0053\u0065\u006c\u0054\u0079\u0070e"}}
		_gf.AddPreserveSpaceAttr(&_gaa, *_gb.SelType)
		e.EncodeElement(_gb.SelType, _gaa)
	}
	if _gb.MultiSel != nil {
		_cf := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u004d\u0075\u006c\u0074\u0069\u0053\u0065\u006c"}}
		_gf.AddPreserveSpaceAttr(&_cf, *_gb.MultiSel)
		e.EncodeElement(_gb.MultiSel, _cf)
	}
	if _gb.LCT != nil {
		_ef := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u004cC\u0054"}}
		_gf.AddPreserveSpaceAttr(&_ef, *_gb.LCT)
		e.EncodeElement(_gb.LCT, _ef)
	}
	if _gb.ListItem != nil {
		_fa := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u004c\u0069\u0073\u0074\u0049\u0074\u0065\u006d"}}
		_gf.AddPreserveSpaceAttr(&_fa, *_gb.ListItem)
		e.EncodeElement(_gb.ListItem, _fa)
	}
	if _gb.DropStyle != nil {
		_fea := _g.StartElement{Name: _g.Name{Local: "x\u003a\u0044\u0072\u006f\u0070\u0053\u0074\u0079\u006c\u0065"}}
		_gf.AddPreserveSpaceAttr(&_fea, *_gb.DropStyle)
		e.EncodeElement(_gb.DropStyle, _fea)
	}
	if _gb.Colored != _c.ST_TrueFalseBlankUnset {
		_feg := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0043\u006f\u006c\u006f\u0072\u0065d"}}
		e.EncodeElement(_gb.Colored, _feg)
	}
	if _gb.DropLines != nil {
		_bec := _g.StartElement{Name: _g.Name{Local: "x\u003a\u0044\u0072\u006f\u0070\u004c\u0069\u006e\u0065\u0073"}}
		e.EncodeElement(_gb.DropLines, _bec)
	}
	if _gb.Checked != nil {
		_bd := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0043\u0068\u0065\u0063\u006b\u0065d"}}
		e.EncodeElement(_gb.Checked, _bd)
	}
	if _gb.FmlaLink != nil {
		_de := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0046\u006d\u006c\u0061\u004c\u0069\u006e\u006b"}}
		_gf.AddPreserveSpaceAttr(&_de, *_gb.FmlaLink)
		e.EncodeElement(_gb.FmlaLink, _de)
	}
	if _gb.FmlaPict != nil {
		_efb := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0046\u006d\u006c\u0061\u0050\u0069\u0063\u0074"}}
		_gf.AddPreserveSpaceAttr(&_efb, *_gb.FmlaPict)
		e.EncodeElement(_gb.FmlaPict, _efb)
	}
	if _gb.NoThreeD != _c.ST_TrueFalseBlankUnset {
		_fbe := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044"}}
		e.EncodeElement(_gb.NoThreeD, _fbe)
	}
	if _gb.FirstButton != _c.ST_TrueFalseBlankUnset {
		_gd := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0046\u0069\u0072\u0073\u0074\u0042\u0075\u0074\u0074\u006f\u006e"}}
		e.EncodeElement(_gb.FirstButton, _gd)
	}
	if _gb.FmlaGroup != nil {
		_ec := _g.StartElement{Name: _g.Name{Local: "x\u003a\u0046\u006d\u006c\u0061\u0047\u0072\u006f\u0075\u0070"}}
		_gf.AddPreserveSpaceAttr(&_ec, *_gb.FmlaGroup)
		e.EncodeElement(_gb.FmlaGroup, _ec)
	}
	if _gb.Val != nil {
		_eb := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0056a\u006c"}}
		e.EncodeElement(_gb.Val, _eb)
	}
	if _gb.Min != nil {
		_cfc := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u004di\u006e"}}
		e.EncodeElement(_gb.Min, _cfc)
	}
	if _gb.Max != nil {
		_bdc := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u004da\u0078"}}
		e.EncodeElement(_gb.Max, _bdc)
	}
	if _gb.Inc != nil {
		_fec := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0049n\u0063"}}
		e.EncodeElement(_gb.Inc, _fec)
	}
	if _gb.Page != nil {
		_af := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0050\u0061\u0067\u0065"}}
		e.EncodeElement(_gb.Page, _af)
	}
	if _gb.Horiz != _c.ST_TrueFalseBlankUnset {
		_deg := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0048\u006f\u0072\u0069\u007a"}}
		e.EncodeElement(_gb.Horiz, _deg)
	}
	if _gb.Dx != nil {
		_cda := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0044\u0078"}}
		e.EncodeElement(_gb.Dx, _cda)
	}
	if _gb.MapOCX != _c.ST_TrueFalseBlankUnset {
		_df := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u004d\u0061\u0070\u004f\u0043\u0058"}}
		e.EncodeElement(_gb.MapOCX, _df)
	}
	if _gb.CF != nil {
		_ecc := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0043\u0046"}}
		for _, _bb := range _gb.CF {
			e.EncodeElement(_bb, _ecc)
		}
	}
	if _gb.Camera != _c.ST_TrueFalseBlankUnset {
		_gaag := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0043\u0061\u006d\u0065\u0072\u0061"}}
		e.EncodeElement(_gb.Camera, _gaag)
	}
	if _gb.RecalcAlways != _c.ST_TrueFalseBlankUnset {
		_bc := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0052\u0065\u0063\u0061\u006c\u0063\u0041l\u0077\u0061\u0079\u0073"}}
		e.EncodeElement(_gb.RecalcAlways, _bc)
	}
	if _gb.AutoScale != _c.ST_TrueFalseBlankUnset {
		_bcb := _g.StartElement{Name: _g.Name{Local: "x\u003a\u0041\u0075\u0074\u006f\u0053\u0063\u0061\u006c\u0065"}}
		e.EncodeElement(_gb.AutoScale, _bcb)
	}
	if _gb.DDE != _c.ST_TrueFalseBlankUnset {
		_ac := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0044D\u0045"}}
		e.EncodeElement(_gb.DDE, _ac)
	}
	if _gb.UIObj != _c.ST_TrueFalseBlankUnset {
		_ge := _g.StartElement{Name: _g.Name{Local: "\u0078:\u0055\u0049\u004f\u0062\u006a"}}
		e.EncodeElement(_gb.UIObj, _ge)
	}
	if _gb.ScriptText != nil {
		_fad := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0053c\u0072\u0069\u0070\u0074\u0054\u0065\u0078\u0074"}}
		_gf.AddPreserveSpaceAttr(&_fad, *_gb.ScriptText)
		e.EncodeElement(_gb.ScriptText, _fad)
	}
	if _gb.ScriptExtended != nil {
		_dfd := _g.StartElement{Name: _g.Name{Local: "\u0078\u003aS\u0063\u0072\u0069p\u0074\u0045\u0078\u0074\u0065\u006e\u0064\u0065\u0064"}}
		_gf.AddPreserveSpaceAttr(&_dfd, *_gb.ScriptExtended)
		e.EncodeElement(_gb.ScriptExtended, _dfd)
	}
	if _gb.ScriptLanguage != nil {
		_fcc := _g.StartElement{Name: _g.Name{Local: "\u0078\u003aS\u0063\u0072\u0069p\u0074\u004c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}}
		e.EncodeElement(_gb.ScriptLanguage, _fcc)
	}
	if _gb.ScriptLocation != nil {
		_bf := _g.StartElement{Name: _g.Name{Local: "\u0078\u003aS\u0063\u0072\u0069p\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"}}
		e.EncodeElement(_gb.ScriptLocation, _bf)
	}
	if _gb.FmlaTxbx != nil {
		_dec := _g.StartElement{Name: _g.Name{Local: "\u0078\u003a\u0046\u006d\u006c\u0061\u0054\u0078\u0062\u0078"}}
		_gf.AddPreserveSpaceAttr(&_dec, *_gb.FmlaTxbx)
		e.EncodeElement(_gb.FmlaTxbx, _dec)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func NewCT_ClientData() *CT_ClientData {
	_gc := &CT_ClientData{}
	_gc.ObjectTypeAttr = ST_ObjectType(1)
	return _gc
}

// ValidateWithPath validates the CT_ClientData and its children, prefixing error messages with path
func (_fd *CT_ClientData) ValidateWithPath(path string) error {
	if _fd.ObjectTypeAttr == ST_ObjectTypeUnset {
		return _b.Errorf("\u0025\u0073\u002f\u004f\u0062\u006a\u0065\u0063\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u0069\u0073\u0020\u0061\u0020\u006da\u006e\u0064\u0061\u0074\u006fr\u0079\u0020f\u0069\u0065\u006c\u0064", path)
	}
	if _faf := _fd.ObjectTypeAttr.ValidateWithPath(path + "\u002fO\u0062j\u0065\u0063\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _faf != nil {
		return _faf
	}
	if _ged := _fd.MoveWithCells.ValidateWithPath(path + "\u002f\u004d\u006f\u0076\u0065\u0057\u0069\u0074\u0068C\u0065\u006c\u006c\u0073"); _ged != nil {
		return _ged
	}
	if _dfc := _fd.SizeWithCells.ValidateWithPath(path + "\u002f\u0053\u0069\u007a\u0065\u0057\u0069\u0074\u0068C\u0065\u006c\u006c\u0073"); _dfc != nil {
		return _dfc
	}
	if _ggd := _fd.Locked.ValidateWithPath(path + "\u002fL\u006f\u0063\u006b\u0065\u0064"); _ggd != nil {
		return _ggd
	}
	if _fbd := _fd.DefaultSize.ValidateWithPath(path + "\u002f\u0044\u0065f\u0061\u0075\u006c\u0074\u0053\u0069\u007a\u0065"); _fbd != nil {
		return _fbd
	}
	if _fga := _fd.PrintObject.ValidateWithPath(path + "\u002f\u0050\u0072i\u006e\u0074\u004f\u0062\u006a\u0065\u0063\u0074"); _fga != nil {
		return _fga
	}
	if _dbe := _fd.Disabled.ValidateWithPath(path + "\u002fD\u0069\u0073\u0061\u0062\u006c\u0065d"); _dbe != nil {
		return _dbe
	}
	if _dg := _fd.AutoFill.ValidateWithPath(path + "\u002fA\u0075\u0074\u006f\u0046\u0069\u006cl"); _dg != nil {
		return _dg
	}
	if _cge := _fd.AutoLine.ValidateWithPath(path + "\u002fA\u0075\u0074\u006f\u004c\u0069\u006ee"); _cge != nil {
		return _cge
	}
	if _gfd := _fd.AutoPict.ValidateWithPath(path + "\u002fA\u0075\u0074\u006f\u0050\u0069\u0063t"); _gfd != nil {
		return _gfd
	}
	if _gedd := _fd.LockText.ValidateWithPath(path + "\u002fL\u006f\u0063\u006b\u0054\u0065\u0078t"); _gedd != nil {
		return _gedd
	}
	if _gbgg := _fd.JustLastX.ValidateWithPath(path + "\u002f\u004a\u0075\u0073\u0074\u004c\u0061\u0073\u0074\u0058"); _gbgg != nil {
		return _gbgg
	}
	if _feb := _fd.SecretEdit.ValidateWithPath(path + "/\u0053\u0065\u0063\u0072\u0065\u0074\u0045\u0064\u0069\u0074"); _feb != nil {
		return _feb
	}
	if _dee := _fd.Default.ValidateWithPath(path + "\u002f\u0044\u0065\u0066\u0061\u0075\u006c\u0074"); _dee != nil {
		return _dee
	}
	if _bca := _fd.Help.ValidateWithPath(path + "\u002f\u0048\u0065l\u0070"); _bca != nil {
		return _bca
	}
	if _cdaf := _fd.Cancel.ValidateWithPath(path + "\u002fC\u0061\u006e\u0063\u0065\u006c"); _cdaf != nil {
		return _cdaf
	}
	if _efdf := _fd.Dismiss.ValidateWithPath(path + "\u002f\u0044\u0069\u0073\u006d\u0069\u0073\u0073"); _efdf != nil {
		return _efdf
	}
	if _gbca := _fd.Visible.ValidateWithPath(path + "\u002f\u0056\u0069\u0073\u0069\u0062\u006c\u0065"); _gbca != nil {
		return _gbca
	}
	if _cad := _fd.RowHidden.ValidateWithPath(path + "\u002f\u0052\u006f\u0077\u0048\u0069\u0064\u0064\u0065\u006e"); _cad != nil {
		return _cad
	}
	if _dd := _fd.ColHidden.ValidateWithPath(path + "\u002f\u0043\u006f\u006c\u0048\u0069\u0064\u0064\u0065\u006e"); _dd != nil {
		return _dd
	}
	if _gcbg := _fd.MultiLine.ValidateWithPath(path + "\u002f\u004d\u0075\u006c\u0074\u0069\u004c\u0069\u006e\u0065"); _gcbg != nil {
		return _gcbg
	}
	if _aaa := _fd.VScroll.ValidateWithPath(path + "\u002f\u0056\u0053\u0063\u0072\u006f\u006c\u006c"); _aaa != nil {
		return _aaa
	}
	if _gac := _fd.ValidIds.ValidateWithPath(path + "\u002fV\u0061\u006c\u0069\u0064\u0049\u0064s"); _gac != nil {
		return _gac
	}
	if _agf := _fd.NoThreeD2.ValidateWithPath(path + "\u002f\u004e\u006f\u0054\u0068\u0072\u0065\u0065\u0044\u0032"); _agf != nil {
		return _agf
	}
	if _egde := _fd.Colored.ValidateWithPath(path + "\u002f\u0043\u006f\u006c\u006f\u0072\u0065\u0064"); _egde != nil {
		return _egde
	}
	if _ddf := _fd.NoThreeD.ValidateWithPath(path + "\u002fN\u006f\u0054\u0068\u0072\u0065\u0065D"); _ddf != nil {
		return _ddf
	}
	if _bgg := _fd.FirstButton.ValidateWithPath(path + "\u002f\u0046\u0069r\u0073\u0074\u0042\u0075\u0074\u0074\u006f\u006e"); _bgg != nil {
		return _bgg
	}
	if _dge := _fd.Horiz.ValidateWithPath(path + "\u002f\u0048\u006f\u0072\u0069\u007a"); _dge != nil {
		return _dge
	}
	if _fff := _fd.MapOCX.ValidateWithPath(path + "\u002fM\u0061\u0070\u004f\u0043\u0058"); _fff != nil {
		return _fff
	}
	if _bfg := _fd.Camera.ValidateWithPath(path + "\u002fC\u0061\u006d\u0065\u0072\u0061"); _bfg != nil {
		return _bfg
	}
	if _dfcg := _fd.RecalcAlways.ValidateWithPath(path + "\u002f\u0052\u0065\u0063\u0061\u006c\u0063\u0041\u006c\u0077\u0061\u0079\u0073"); _dfcg != nil {
		return _dfcg
	}
	if _bfga := _fd.AutoScale.ValidateWithPath(path + "\u002f\u0041\u0075\u0074\u006f\u0053\u0063\u0061\u006c\u0065"); _bfga != nil {
		return _bfga
	}
	if _ace := _fd.DDE.ValidateWithPath(path + "\u002f\u0044\u0044\u0045"); _ace != nil {
		return _ace
	}
	if _dbf := _fd.UIObj.ValidateWithPath(path + "\u002f\u0055\u0049\u004f\u0062\u006a"); _dbf != nil {
		return _dbf
	}
	return nil
}
func init() {
	_gf.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", "\u0043\u0054\u005f\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061", NewCT_ClientData)
	_gf.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c", "\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061", NewClientData)
}
