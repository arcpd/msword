package sharedTypes

import (
	_e "encoding/xml"
	_ea "fmt"
	_ec "regexp"
)

func (_dce *ST_TrueFalseBlank) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_cgg, _ebb := d.Token()
	if _ebb != nil {
		return _ebb
	}
	if _bfde, _dea := _cgg.(_e.EndElement); _dea && _bfde.Name == start.Name {
		*_dce = 1
		return nil
	}
	if _beg, _gf := _cgg.(_e.CharData); !_gf {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cgg)
	} else {
		switch string(_beg) {
		case "":
			*_dce = 0
		case "\u0074":
			*_dce = 1
		case "\u0066":
			*_dce = 2
		case "\u0074\u0072\u0075\u0065":
			*_dce = 3
		case "\u0066\u0061\u006cs\u0065":
			*_dce = 4
		case "\u0054\u0072\u0075\u0065":
			*_dce = 6
		case "\u0046\u0061\u006cs\u0065":
			*_dce = 7
		}
	}
	_cgg, _ebb = d.Token()
	if _ebb != nil {
		return _ebb
	}
	if _fbfb, _ega := _cgg.(_e.EndElement); _ega && _fbfb.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cgg)
}

func (_eg ST_CalendarType) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_dg := _e.Attr{}
	_dg.Name = name
	switch _eg {
	case ST_CalendarTypeUnset:
		_dg.Value = ""
	case ST_CalendarTypeGregorian:
		_dg.Value = "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n"
	case ST_CalendarTypeGregorianUs:
		_dg.Value = "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073"
	case ST_CalendarTypeGregorianMeFrench:
		_dg.Value = "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068"
	case ST_CalendarTypeGregorianArabic:
		_dg.Value = "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063"
	case ST_CalendarTypeHijri:
		_dg.Value = "\u0068\u0069\u006ar\u0069"
	case ST_CalendarTypeHebrew:
		_dg.Value = "\u0068\u0065\u0062\u0072\u0065\u0077"
	case ST_CalendarTypeTaiwan:
		_dg.Value = "\u0074\u0061\u0069\u0077\u0061\u006e"
	case ST_CalendarTypeJapan:
		_dg.Value = "\u006a\u0061\u0070a\u006e"
	case ST_CalendarTypeThai:
		_dg.Value = "\u0074\u0068\u0061\u0069"
	case ST_CalendarTypeKorea:
		_dg.Value = "\u006b\u006f\u0072e\u0061"
	case ST_CalendarTypeSaka:
		_dg.Value = "\u0073\u0061\u006b\u0061"
	case ST_CalendarTypeGregorianXlitEnglish:
		_dg.Value = "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068"
	case ST_CalendarTypeGregorianXlitFrench:
		_dg.Value = "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068"
	case ST_CalendarTypeNone:
		_dg.Value = "\u006e\u006f\u006e\u0065"
	}
	return _dg, nil
}

const (
	ST_TrueFalseBlankUnset  ST_TrueFalseBlank = 0
	ST_TrueFalseBlankT      ST_TrueFalseBlank = 1
	ST_TrueFalseBlankF      ST_TrueFalseBlank = 2
	ST_TrueFalseBlankTrue   ST_TrueFalseBlank = 3
	ST_TrueFalseBlankFalse  ST_TrueFalseBlank = 4
	ST_TrueFalseBlankTrue_  ST_TrueFalseBlank = 6
	ST_TrueFalseBlankFalse_ ST_TrueFalseBlank = 7
)

func (_bfb ST_AlgType) Validate() error { return _bfb.ValidateWithPath("") }

var ST_PositiveFixedPercentagePatternRe = _ec.MustCompile(ST_PositiveFixedPercentagePattern)

func (_deee *ST_XAlign) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_eccf, _ggb := d.Token()
	if _ggb != nil {
		return _ggb
	}
	if _bgf, _bdc := _eccf.(_e.EndElement); _bdc && _bgf.Name == start.Name {
		*_deee = 1
		return nil
	}
	if _bfe, _dgd := _eccf.(_e.CharData); !_dgd {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eccf)
	} else {
		switch string(_bfe) {
		case "":
			*_deee = 0
		case "\u006c\u0065\u0066\u0074":
			*_deee = 1
		case "\u0063\u0065\u006e\u0074\u0065\u0072":
			*_deee = 2
		case "\u0072\u0069\u0067h\u0074":
			*_deee = 3
		case "\u0069\u006e\u0073\u0069\u0064\u0065":
			*_deee = 4
		case "\u006fu\u0074\u0073\u0069\u0064\u0065":
			*_deee = 5
		}
	}
	_eccf, _ggb = d.Token()
	if _ggb != nil {
		return _ggb
	}
	if _bfbe, _dae := _eccf.(_e.EndElement); _dae && _bfbe.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eccf)
}

func _a(_de bool) uint8 {
	if _de {
		return 1
	}
	return 0
}

func (_dbbe *ST_YAlign) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_dbbe = 0
	case "\u0069\u006e\u006c\u0069\u006e\u0065":
		*_dbbe = 1
	case "\u0074\u006f\u0070":
		*_dbbe = 2
	case "\u0063\u0065\u006e\u0074\u0065\u0072":
		*_dbbe = 3
	case "\u0062\u006f\u0074\u0074\u006f\u006d":
		*_dbbe = 4
	case "\u0069\u006e\u0073\u0069\u0064\u0065":
		*_dbbe = 5
	case "\u006fu\u0074\u0073\u0069\u0064\u0065":
		*_dbbe = 6
	}
	return nil
}

type ST_XAlign byte

func (_aa *ST_AlgClass) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_aa = 0
	case "\u0068\u0061\u0073\u0068":
		*_aa = 1
	case "\u0063\u0075\u0073\u0074\u006f\u006d":
		*_aa = 2
	}
	return nil
}

func (_fdb *ST_ConformanceClass) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_fdb = 0
	case "\u0073\u0074\u0072\u0069\u0063\u0074":
		*_fdb = 1
	case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":
		*_fdb = 2
	}
	return nil
}

// ST_TwipsMeasure is a union type
type ST_TwipsMeasure struct {
	ST_UnsignedDecimalNumber    *uint64
	ST_PositiveUniversalMeasure *string
}

func (_beb *ST_VerticalAlignRun) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_fabe, _cefe := d.Token()
	if _cefe != nil {
		return _cefe
	}
	if _ceg, _ag := _fabe.(_e.EndElement); _ag && _ceg.Name == start.Name {
		*_beb = 1
		return nil
	}
	if _gcf, _dcc := _fabe.(_e.CharData); !_dcc {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fabe)
	} else {
		switch string(_gcf) {
		case "":
			*_beb = 0
		case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":
			*_beb = 1
		case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":
			*_beb = 2
		case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":
			*_beb = 3
		}
	}
	_fabe, _cefe = d.Token()
	if _cefe != nil {
		return _cefe
	}
	if _fcb, _dee := _fabe.(_e.EndElement); _dee && _fcb.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fabe)
}

var ST_GuidPatternRe = _ec.MustCompile(ST_GuidPattern)

func (_ab ST_AlgType) ValidateWithPath(path string) error {
	switch _ab {
	case 0, 1, 2:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ab))
	}
	return nil
}

func (_db *ST_CryptProv) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_db = 0
	case "\u0072\u0073\u0061\u0041\u0045\u0053":
		*_db = 1
	case "\u0072s\u0061\u0046\u0075\u006c\u006c":
		*_db = 2
	case "\u0063\u0075\u0073\u0074\u006f\u006d":
		*_db = 3
	}
	return nil
}

const ST_PositivePercentagePattern = "\u005b0\u002d9\u005d\u002b\u0028\u005c\u002e[\u0030\u002d9\u005d\u002b\u0029\u003f\u0025"

func (_da ST_CryptProv) Validate() error { return _da.ValidateWithPath("") }

// ST_OnOff is a union type
type ST_OnOff struct {
	Bool      *bool
	ST_OnOff1 ST_OnOff1
}

func (_abe *ST_TrueFalseBlank) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_abe = 0
	case "\u0074":
		*_abe = 1
	case "\u0066":
		*_abe = 2
	case "\u0074\u0072\u0075\u0065":
		*_abe = 3
	case "\u0066\u0061\u006cs\u0065":
		*_abe = 4
	case "\u0054\u0072\u0075\u0065":
		*_abe = 6
	case "\u0046\u0061\u006cs\u0065":
		*_abe = 7
	}
	return nil
}

func (_ecb ST_ConformanceClass) String() string {
	switch _ecb {
	case 0:
		return ""
	case 1:
		return "\u0073\u0074\u0072\u0069\u0063\u0074"
	case 2:
		return "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c"
	}
	return ""
}

func (_b ST_OnOff) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	e.EncodeToken(start)
	if _b.Bool != nil {
		e.EncodeToken(_e.CharData(_ea.Sprintf("\u0025\u0064", _a(*_b.Bool))))
	}
	if _b.ST_OnOff1 != ST_OnOff1Unset {
		e.EncodeToken(_e.CharData(_b.ST_OnOff1.String()))
	}
	return e.EncodeToken(_e.EndElement{Name: start.Name})
}

func (_cae ST_ConformanceClass) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_dcca := _e.Attr{}
	_dcca.Name = name
	switch _cae {
	case ST_ConformanceClassUnset:
		_dcca.Value = ""
	case ST_ConformanceClassStrict:
		_dcca.Value = "\u0073\u0074\u0072\u0069\u0063\u0074"
	case ST_ConformanceClassTransitional:
		_dcca.Value = "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c"
	}
	return _dcca, nil
}

func (_dge ST_CalendarType) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_dge.String(), start)
}

var ST_UniversalMeasurePatternRe = _ec.MustCompile(ST_UniversalMeasurePattern)

const ST_FixedPercentagePattern = "-\u003f\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039]\u003f\u0029\u0029\u0028\u005c\u002e\u005b\u0030\u002d\u0039][\u0030\u002d\u0039]\u003f)\u003f\u0025"

func (_dcg ST_CalendarType) ValidateWithPath(path string) error {
	switch _dcg {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dcg))
	}
	return nil
}

func (_bbc ST_VerticalAlignRun) ValidateWithPath(path string) error {
	switch _bbc {
	case 0, 1, 2, 3:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bbc))
	}
	return nil
}

func (_ddc ST_TrueFalse) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_ddc.String(), start)
}

func (_eec ST_VerticalAlignRun) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_eeb := _e.Attr{}
	_eeb.Name = name
	switch _eec {
	case ST_VerticalAlignRunUnset:
		_eeb.Value = ""
	case ST_VerticalAlignRunBaseline:
		_eeb.Value = "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065"
	case ST_VerticalAlignRunSuperscript:
		_eeb.Value = "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074"
	case ST_VerticalAlignRunSubscript:
		_eeb.Value = "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t"
	}
	return _eeb, nil
}

func (_dba *ST_AlgType) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_fbc, _bgb := d.Token()
	if _bgb != nil {
		return _bgb
	}
	if _fag, _dgb := _fbc.(_e.EndElement); _dgb && _fag.Name == start.Name {
		*_dba = 1
		return nil
	}
	if _dag, _cde := _fbc.(_e.CharData); !_cde {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fbc)
	} else {
		switch string(_dag) {
		case "":
			*_dba = 0
		case "\u0074y\u0070\u0065\u0041\u006e\u0079":
			*_dba = 1
		case "\u0063\u0075\u0073\u0074\u006f\u006d":
			*_dba = 2
		}
	}
	_fbc, _bgb = d.Token()
	if _bgb != nil {
		return _bgb
	}
	if _bff, _fda := _fbc.(_e.EndElement); _fda && _bff.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fbc)
}

func (_aefc *ST_XAlign) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_aefc = 0
	case "\u006c\u0065\u0066\u0074":
		*_aefc = 1
	case "\u0063\u0065\u006e\u0074\u0065\u0072":
		*_aefc = 2
	case "\u0072\u0069\u0067h\u0074":
		*_aefc = 3
	case "\u0069\u006e\u0073\u0069\u0064\u0065":
		*_aefc = 4
	case "\u006fu\u0074\u0073\u0069\u0064\u0065":
		*_aefc = 5
	}
	return nil
}

type ST_AlgType byte

var ST_FixedPercentagePatternRe = _ec.MustCompile(ST_FixedPercentagePattern)

func (_bac *ST_ConformanceClass) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_cega, _deg := d.Token()
	if _deg != nil {
		return _deg
	}
	if _dcb, _gce := _cega.(_e.EndElement); _gce && _dcb.Name == start.Name {
		*_bac = 1
		return nil
	}
	if _fagb, _bade := _cega.(_e.CharData); !_bade {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cega)
	} else {
		switch string(_fagb) {
		case "":
			*_bac = 0
		case "\u0073\u0074\u0072\u0069\u0063\u0074":
			*_bac = 1
		case "\u0074\u0072\u0061n\u0073\u0069\u0074\u0069\u006f\u006e\u0061\u006c":
			*_bac = 2
		}
	}
	_cega, _deg = d.Token()
	if _deg != nil {
		return _deg
	}
	if _eac, _ccb := _cega.(_e.EndElement); _ccb && _eac.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cega)
}

const ST_UniversalMeasurePattern = "\u002d\u003f\u005b\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e\u005b\u0030\u002d\u0039\u005d\u002b\u0029\u003f\u0028\u006d\u006d\u007c\u0063m\u007c\u0069\u006e\u007c\u0070t\u007c\u0070c\u007c\u0070\u0069\u0029"

type ST_AlgClass byte

func (_bg *ST_CryptProv) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_ba, _df := d.Token()
	if _df != nil {
		return _df
	}
	if _fg, _ac := _ba.(_e.EndElement); _ac && _fg.Name == start.Name {
		*_bg = 1
		return nil
	}
	if _fbf, _gba := _ba.(_e.CharData); !_gba {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ba)
	} else {
		switch string(_fbf) {
		case "":
			*_bg = 0
		case "\u0072\u0073\u0061\u0041\u0045\u0053":
			*_bg = 1
		case "\u0072s\u0061\u0046\u0075\u006c\u006c":
			*_bg = 2
		case "\u0063\u0075\u0073\u0074\u006f\u006d":
			*_bg = 3
		}
	}
	_ba, _df = d.Token()
	if _df != nil {
		return _df
	}
	if _bbg, _be := _ba.(_e.EndElement); _be && _bbg.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ba)
}

const ST_PositiveFixedPercentagePattern = "\u0028\u0028\u0031\u0030\u0030\u0029\u007c\u0028\u005b\u0030\u002d\u0039\u005d\u005b\u0030\u002d\u0039\u005d\u003f\u0029\u0029\u0028\u005c\u002e[\u0030\u002d\u0039\u005d\u005b0\u002d\u0039]\u003f\u0029\u003f\u0025"

func (_cag ST_CalendarType) String() string {
	switch _cag {
	case 0:
		return ""
	case 1:
		return "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n"
	case 2:
		return "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073"
	case 3:
		return "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068"
	case 4:
		return "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063"
	case 5:
		return "\u0068\u0069\u006ar\u0069"
	case 6:
		return "\u0068\u0065\u0062\u0072\u0065\u0077"
	case 7:
		return "\u0074\u0061\u0069\u0077\u0061\u006e"
	case 8:
		return "\u006a\u0061\u0070a\u006e"
	case 9:
		return "\u0074\u0068\u0061\u0069"
	case 10:
		return "\u006b\u006f\u0072e\u0061"
	case 11:
		return "\u0073\u0061\u006b\u0061"
	case 12:
		return "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068"
	case 13:
		return "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068"
	case 14:
		return "\u006e\u006f\u006e\u0065"
	}
	return ""
}

var ST_PositiveUniversalMeasurePatternRe = _ec.MustCompile(ST_PositiveUniversalMeasurePattern)

func (_ef ST_AlgClass) String() string {
	switch _ef {
	case 0:
		return ""
	case 1:
		return "\u0068\u0061\u0073\u0068"
	case 2:
		return "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return ""
}

func (_fdc ST_XAlign) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_fdc.String(), start)
}

const (
	ST_XAlignUnset   ST_XAlign = 0
	ST_XAlignLeft    ST_XAlign = 1
	ST_XAlignCenter  ST_XAlign = 2
	ST_XAlignRight   ST_XAlign = 3
	ST_XAlignInside  ST_XAlign = 4
	ST_XAlignOutside ST_XAlign = 5
)

func (_eccd ST_XAlign) Validate() error { return _eccd.ValidateWithPath("") }

func (_fd ST_OnOff) String() string {
	if _fd.Bool != nil {
		return _ea.Sprintf("\u0025\u0076", *_fd.Bool)
	}
	if _fd.ST_OnOff1 != ST_OnOff1Unset {
		return _fd.ST_OnOff1.String()
	}
	return ""
}

func (_fcd ST_ConformanceClass) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_fcd.String(), start)
}

func (_ga *ST_TrueFalse) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_gga, _fac := d.Token()
	if _fac != nil {
		return _fac
	}
	if _dbb, _ffc := _gga.(_e.EndElement); _ffc && _dbb.Name == start.Name {
		*_ga = 1
		return nil
	}
	if _gdf, _bfbf := _gga.(_e.CharData); !_bfbf {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gga)
	} else {
		switch string(_gdf) {
		case "":
			*_ga = 0
		case "\u0074":
			*_ga = 1
		case "\u0066":
			*_ga = 2
		case "\u0074\u0072\u0075\u0065":
			*_ga = 3
		case "\u0066\u0061\u006cs\u0065":
			*_ga = 4
		}
	}
	_gga, _fac = d.Token()
	if _fac != nil {
		return _fac
	}
	if _eee, _dbg := _gga.(_e.EndElement); _dbg && _eee.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gga)
}

func (_ed *ST_TwipsMeasure) Validate() error { return _ed.ValidateWithPath("") }

func (_ceb *ST_VerticalAlignRun) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_ceb = 0
	case "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065":
		*_ceb = 1
	case "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074":
		*_ceb = 2
	case "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t":
		*_ceb = 3
	}
	return nil
}

type ST_CryptProv byte

func (_dbd ST_TrueFalse) Validate() error { return _dbd.ValidateWithPath("") }

func (_faef ST_YAlign) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_gbc := _e.Attr{}
	_gbc.Name = name
	switch _faef {
	case ST_YAlignUnset:
		_gbc.Value = ""
	case ST_YAlignInline:
		_gbc.Value = "\u0069\u006e\u006c\u0069\u006e\u0065"
	case ST_YAlignTop:
		_gbc.Value = "\u0074\u006f\u0070"
	case ST_YAlignCenter:
		_gbc.Value = "\u0063\u0065\u006e\u0074\u0065\u0072"
	case ST_YAlignBottom:
		_gbc.Value = "\u0062\u006f\u0074\u0074\u006f\u006d"
	case ST_YAlignInside:
		_gbc.Value = "\u0069\u006e\u0073\u0069\u0064\u0065"
	case ST_YAlignOutside:
		_gbc.Value = "\u006fu\u0074\u0073\u0069\u0064\u0065"
	}
	return _gbc, nil
}

func (_egb ST_CryptProv) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_ae := _e.Attr{}
	_ae.Name = name
	switch _egb {
	case ST_CryptProvUnset:
		_ae.Value = ""
	case ST_CryptProvRsaAES:
		_ae.Value = "\u0072\u0073\u0061\u0041\u0045\u0053"
	case ST_CryptProvRsaFull:
		_ae.Value = "\u0072s\u0061\u0046\u0075\u006c\u006c"
	case ST_CryptProvCustom:
		_ae.Value = "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return _ae, nil
}

func (_dbab ST_XAlign) ValidateWithPath(path string) error {
	switch _dbab {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dbab))
	}
	return nil
}

type ST_VerticalAlignRun byte

func (_bga *ST_OnOff1) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_ege, _fgg := d.Token()
	if _fgg != nil {
		return _fgg
	}
	if _bec, _gcc := _ege.(_e.EndElement); _gcc && _bec.Name == start.Name {
		*_bga = 1
		return nil
	}
	if _cb, _ce := _ege.(_e.CharData); !_ce {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ege)
	} else {
		switch string(_cb) {
		case "":
			*_bga = 0
		case "\u006f\u006e":
			*_bga = 1
		case "\u006f\u0066\u0066":
			*_bga = 2
		}
	}
	_ege, _fgg = d.Token()
	if _fgg != nil {
		return _fgg
	}
	if _ff, _aeb := _ege.(_e.EndElement); _aeb && _ff.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ege)
}

func (_gc *ST_AlgType) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_gc = 0
	case "\u0074y\u0070\u0065\u0041\u006e\u0079":
		*_gc = 1
	case "\u0063\u0075\u0073\u0074\u006f\u006d":
		*_gc = 2
	}
	return nil
}

func (_fab ST_AlgType) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_fab.String(), start)
}

func (_cd *ST_CalendarType) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_cd = 0
	case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":
		*_cd = 1
	case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":
		*_cd = 2
	case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":
		*_cd = 3
	case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":
		*_cd = 4
	case "\u0068\u0069\u006ar\u0069":
		*_cd = 5
	case "\u0068\u0065\u0062\u0072\u0065\u0077":
		*_cd = 6
	case "\u0074\u0061\u0069\u0077\u0061\u006e":
		*_cd = 7
	case "\u006a\u0061\u0070a\u006e":
		*_cd = 8
	case "\u0074\u0068\u0061\u0069":
		*_cd = 9
	case "\u006b\u006f\u0072e\u0061":
		*_cd = 10
	case "\u0073\u0061\u006b\u0061":
		*_cd = 11
	case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":
		*_cd = 12
	case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":
		*_cd = 13
	case "\u006e\u006f\u006e\u0065":
		*_cd = 14
	}
	return nil
}

func (_gea ST_CalendarType) Validate() error { return _gea.ValidateWithPath("") }

func (_ggg ST_VerticalAlignRun) String() string {
	switch _ggg {
	case 0:
		return ""
	case 1:
		return "\u0062\u0061\u0073\u0065\u006c\u0069\u006e\u0065"
	case 2:
		return "s\u0075\u0070\u0065\u0072\u0073\u0063\u0072\u0069\u0070\u0074"
	case 3:
		return "\u0073u\u0062\u0073\u0063\u0072\u0069\u0070t"
	}
	return ""
}

const (
	ST_OnOff1Unset ST_OnOff1 = 0
	ST_OnOff1On    ST_OnOff1 = 1
	ST_OnOff1Off   ST_OnOff1 = 2
)

func (_cbf ST_VerticalAlignRun) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_cbf.String(), start)
}

func (_fb *ST_TwipsMeasure) ValidateWithPath(path string) error {
	_fba := []string{}
	if _fb.ST_UnsignedDecimalNumber != nil {
		_fba = append(_fba, "\u0053T\u005f\u0055\u006e\u0073\u0069\u0067\u006e\u0065\u0064\u0044\u0065c\u0069\u006d\u0061\u006c\u004e\u0075\u006d\u0062\u0065\u0072")
	}
	if _fb.ST_PositiveUniversalMeasure != nil {
		_fba = append(_fba, "S\u0054\u005f\u0050\u006f\u0073\u0069t\u0069\u0076\u0065\u0055\u006e\u0069\u0076\u0065\u0072s\u0061\u006c\u004de\u0061s\u0075\u0072\u0065")
	}
	if len(_fba) > 1 {
		return _ea.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _fba)
	}
	return nil
}

func (_ebc ST_YAlign) ValidateWithPath(path string) error {
	switch _ebc {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ebc))
	}
	return nil
}

type ST_TrueFalseBlank byte

type ST_TrueFalse byte

func (_dd ST_TwipsMeasure) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	e.EncodeToken(start)
	if _dd.ST_UnsignedDecimalNumber != nil {
		e.EncodeToken(_e.CharData(_ea.Sprintf("\u0025\u0064", *_dd.ST_UnsignedDecimalNumber)))
	}
	if _dd.ST_PositiveUniversalMeasure != nil {
		e.EncodeToken(_e.CharData(*_dd.ST_PositiveUniversalMeasure))
	}
	return e.EncodeToken(_e.EndElement{Name: start.Name})
}

func (_faee ST_OnOff1) ValidateWithPath(path string) error {
	switch _faee {
	case 0, 1, 2:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_faee))
	}
	return nil
}

func (_cga ST_CryptProv) String() string {
	switch _cga {
	case 0:
		return ""
	case 1:
		return "\u0072\u0073\u0061\u0041\u0045\u0053"
	case 2:
		return "\u0072s\u0061\u0046\u0075\u006c\u006c"
	case 3:
		return "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return ""
}

type ST_ConformanceClass byte

const ST_PositiveUniversalMeasurePattern = "\u005b\u0030-9\u005d\u002b\u0028\\\u002e\u005b\u0030\u002d9]+\u0029?(\u006d\u006d\u007c\u0063\u006d\u007c\u0069n|\u0070\u0074\u007c\u0070\u0063\u007c\u0070i\u0029"

func (_ecfe *ST_YAlign) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_fcc, _ebe := d.Token()
	if _ebe != nil {
		return _ebe
	}
	if _baa, _cc := _fcc.(_e.EndElement); _cc && _baa.Name == start.Name {
		*_ecfe = 1
		return nil
	}
	if _acb, _cda := _fcc.(_e.CharData); !_cda {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fcc)
	} else {
		switch string(_acb) {
		case "":
			*_ecfe = 0
		case "\u0069\u006e\u006c\u0069\u006e\u0065":
			*_ecfe = 1
		case "\u0074\u006f\u0070":
			*_ecfe = 2
		case "\u0063\u0065\u006e\u0074\u0065\u0072":
			*_ecfe = 3
		case "\u0062\u006f\u0074\u0074\u006f\u006d":
			*_ecfe = 4
		case "\u0069\u006e\u0073\u0069\u0064\u0065":
			*_ecfe = 5
		case "\u006fu\u0074\u0073\u0069\u0064\u0065":
			*_ecfe = 6
		}
	}
	_fcc, _ebe = d.Token()
	if _ebe != nil {
		return _ebe
	}
	if _adb, _ggf := _fcc.(_e.EndElement); _ggf && _adb.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fcc)
}

func (_bge ST_AlgType) String() string {
	switch _bge {
	case 0:
		return ""
	case 1:
		return "\u0074y\u0070\u0065\u0041\u006e\u0079"
	case 2:
		return "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return ""
}

func (_c *ST_OnOff) ValidateWithPath(path string) error {
	_cf := []string{}
	if _c.Bool != nil {
		_cf = append(_cf, "\u0042\u006f\u006f\u006c")
	}
	if _c.ST_OnOff1 != ST_OnOff1Unset {
		_cf = append(_cf, "\u0053T\u005f\u004f\u006e\u004f\u0066\u00661")
	}
	if len(_cf) > 1 {
		return _ea.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _cf)
	}
	return nil
}

const (
	ST_CryptProvUnset   ST_CryptProv = 0
	ST_CryptProvRsaAES  ST_CryptProv = 1
	ST_CryptProvRsaFull ST_CryptProv = 2
	ST_CryptProvCustom  ST_CryptProv = 3
)

func (_cfb *ST_AlgClass) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_bda, _gb := d.Token()
	if _gb != nil {
		return _gb
	}
	if _fae, _bc := _bda.(_e.EndElement); _bc && _fae.Name == start.Name {
		*_cfb = 1
		return nil
	}
	if _ad, _bf := _bda.(_e.CharData); !_bf {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bda)
	} else {
		switch string(_ad) {
		case "":
			*_cfb = 0
		case "\u0068\u0061\u0073\u0068":
			*_cfb = 1
		case "\u0063\u0075\u0073\u0074\u006f\u006d":
			*_cfb = 2
		}
	}
	_bda, _gb = d.Token()
	if _gb != nil {
		return _gb
	}
	if _bcg, _cg := _bda.(_e.EndElement); _cg && _bcg.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bda)
}

func ParseUnionST_OnOff(s string) (ST_OnOff, error) {
	_fc := ST_OnOff{}
	switch s {
	case "\u0074\u0072\u0075\u0065", "\u0031", "\u006f\u006e":
		_bd := true
		_fc.Bool = &_bd
	default:
		_dc := false
		_fc.Bool = &_dc
	}
	return _fc, nil
}

const (
	ST_ConformanceClassUnset        ST_ConformanceClass = 0
	ST_ConformanceClassStrict       ST_ConformanceClass = 1
	ST_ConformanceClassTransitional ST_ConformanceClass = 2
)

type ST_YAlign byte

func (_efb ST_AlgClass) Validate() error { return _efb.ValidateWithPath("") }

func (_eddd ST_TrueFalseBlank) ValidateWithPath(path string) error {
	switch _eddd {
	case 0, 1, 2, 3, 4, 6, 7:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_eddd))
	}
	return nil
}

func (_ecc ST_TwipsMeasure) String() string {
	if _ecc.ST_UnsignedDecimalNumber != nil {
		return _ea.Sprintf("\u0025\u0076", *_ecc.ST_UnsignedDecimalNumber)
	}
	if _ecc.ST_PositiveUniversalMeasure != nil {
		return _ea.Sprintf("\u0025\u0076", *_ecc.ST_PositiveUniversalMeasure)
	}
	return ""
}

func (_ecf ST_TrueFalseBlank) Validate() error { return _ecf.ValidateWithPath("") }

func (_egg ST_VerticalAlignRun) Validate() error { return _egg.ValidateWithPath("") }

func (_dfd *ST_TrueFalse) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_dfd = 0
	case "\u0074":
		*_dfd = 1
	case "\u0066":
		*_dfd = 2
	case "\u0074\u0072\u0075\u0065":
		*_dfd = 3
	case "\u0066\u0061\u006cs\u0065":
		*_dfd = 4
	}
	return nil
}

func (_fabd ST_OnOff1) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_fabd.String(), start)
}

var ST_PositivePercentagePatternRe = _ec.MustCompile(ST_PositivePercentagePattern)

const (
	ST_TrueFalseUnset ST_TrueFalse = 0
	ST_TrueFalseT     ST_TrueFalse = 1
	ST_TrueFalseF     ST_TrueFalse = 2
	ST_TrueFalseTrue  ST_TrueFalse = 3
	ST_TrueFalseFalse ST_TrueFalse = 4
)

func (_d *ST_OnOff) Validate() error { return _d.ValidateWithPath("") }

func (_abb *ST_OnOff1) UnmarshalXMLAttr(attr _e.Attr) error {
	switch attr.Value {
	case "":
		*_abb = 0
	case "\u006f\u006e":
		*_abb = 1
	case "\u006f\u0066\u0066":
		*_abb = 2
	}
	return nil
}

const ST_PercentagePattern = "-\u003f[\u0030\u002d\u0039\u005d\u002b\u0028\u005c\u002e[\u0030\u002d\u0039\u005d+)\u003f\u0025"

func (_g *ST_CalendarType) UnmarshalXML(d *_e.Decoder, start _e.StartElement) error {
	_ge, _gec := d.Token()
	if _gec != nil {
		return _gec
	}
	if _ee, _gd := _ge.(_e.EndElement); _gd && _ee.Name == start.Name {
		*_g = 1
		return nil
	}
	if _ddg, _eb := _ge.(_e.CharData); !_eb {
		return _ea.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ge)
	} else {
		switch string(_ddg) {
		case "":
			*_g = 0
		case "\u0067r\u0065\u0067\u006f\u0072\u0069\u0061n":
			*_g = 1
		case "g\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u0055\u0073":
			*_g = 2
		case "\u0067\u0072\u0065\u0067\u006f\u0072\u0069\u0061\u006e\u004d\u0065\u0046r\u0065\u006e\u0063\u0068":
			*_g = 3
		case "\u0067r\u0065g\u006f\u0072\u0069\u0061\u006e\u0041\u0072\u0061\u0062\u0069\u0063":
			*_g = 4
		case "\u0068\u0069\u006ar\u0069":
			*_g = 5
		case "\u0068\u0065\u0062\u0072\u0065\u0077":
			*_g = 6
		case "\u0074\u0061\u0069\u0077\u0061\u006e":
			*_g = 7
		case "\u006a\u0061\u0070a\u006e":
			*_g = 8
		case "\u0074\u0068\u0061\u0069":
			*_g = 9
		case "\u006b\u006f\u0072e\u0061":
			*_g = 10
		case "\u0073\u0061\u006b\u0061":
			*_g = 11
		case "g\u0072e\u0067\u006f\u0072\u0069\u0061\u006e\u0058\u006ci\u0074\u0045\u006e\u0067li\u0073\u0068":
			*_g = 12
		case "\u0067\u0072\u0065\u0067or\u0069\u0061\u006e\u0058\u006c\u0069\u0074\u0046\u0072\u0065\u006e\u0063\u0068":
			*_g = 13
		case "\u006e\u006f\u006e\u0065":
			*_g = 14
		}
	}
	_ge, _gec = d.Token()
	if _gec != nil {
		return _gec
	}
	if _gdg, _ca := _ge.(_e.EndElement); _ca && _gdg.Name == start.Name {
		return nil
	}
	return _ea.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ge)
}

func (_bb ST_AlgClass) ValidateWithPath(path string) error {
	switch _bb {
	case 0, 1, 2:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bb))
	}
	return nil
}

func (_bad ST_CryptProv) ValidateWithPath(path string) error {
	switch _bad {
	case 0, 1, 2, 3:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bad))
	}
	return nil
}

func (_aefd ST_YAlign) Validate() error { return _aefd.ValidateWithPath("") }

func (_eab ST_XAlign) String() string {
	switch _eab {
	case 0:
		return ""
	case 1:
		return "\u006c\u0065\u0066\u0074"
	case 2:
		return "\u0063\u0065\u006e\u0074\u0065\u0072"
	case 3:
		return "\u0072\u0069\u0067h\u0074"
	case 4:
		return "\u0069\u006e\u0073\u0069\u0064\u0065"
	case 5:
		return "\u006fu\u0074\u0073\u0069\u0064\u0065"
	}
	return ""
}

func (_cef ST_TrueFalse) ValidateWithPath(path string) error {
	switch _cef {
	case 0, 1, 2, 3, 4:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cef))
	}
	return nil
}

func (_ffa ST_YAlign) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_ffa.String(), start)
}

func (_fa ST_AlgClass) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_geg := _e.Attr{}
	_geg.Name = name
	switch _fa {
	case ST_AlgClassUnset:
		_geg.Value = ""
	case ST_AlgClassHash:
		_geg.Value = "\u0068\u0061\u0073\u0068"
	case ST_AlgClassCustom:
		_geg.Value = "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return _geg, nil
}

const (
	ST_VerticalAlignRunUnset       ST_VerticalAlignRun = 0
	ST_VerticalAlignRunBaseline    ST_VerticalAlignRun = 1
	ST_VerticalAlignRunSuperscript ST_VerticalAlignRun = 2
	ST_VerticalAlignRunSubscript   ST_VerticalAlignRun = 3
)

const ST_GuidPattern = "\u005c\u007b\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b\u0038\u007d\u002d\u005b\u0030\u002d9\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030-\u0039\u0041\u002d\u0046\u005d\u007b\u0034\u007d\u002d\u005b\u0030\u002d\u0039\u0041\u002d\u0046\u005d\u007b4\u007d\u002d\u005b\u0030\u002d\u0039A\u002d\u0046]\u007b\u00312\u007d\\\u007d"

type ST_CalendarType byte

func (_cdd ST_OnOff1) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_bdd := _e.Attr{}
	_bdd.Name = name
	switch _cdd {
	case ST_OnOff1Unset:
		_bdd.Value = ""
	case ST_OnOff1On:
		_bdd.Value = "\u006f\u006e"
	case ST_OnOff1Off:
		_bdd.Value = "\u006f\u0066\u0066"
	}
	return _bdd, nil
}

const (
	ST_YAlignUnset   ST_YAlign = 0
	ST_YAlignInline  ST_YAlign = 1
	ST_YAlignTop     ST_YAlign = 2
	ST_YAlignCenter  ST_YAlign = 3
	ST_YAlignBottom  ST_YAlign = 4
	ST_YAlignInside  ST_YAlign = 5
	ST_YAlignOutside ST_YAlign = 6
)

func (_dfg ST_TrueFalseBlank) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_dfg.String(), start)
}

func (_gbe ST_TrueFalse) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_aad := _e.Attr{}
	_aad.Name = name
	switch _gbe {
	case ST_TrueFalseUnset:
		_aad.Value = ""
	case ST_TrueFalseT:
		_aad.Value = "\u0074"
	case ST_TrueFalseF:
		_aad.Value = "\u0066"
	case ST_TrueFalseTrue:
		_aad.Value = "\u0074\u0072\u0075\u0065"
	case ST_TrueFalseFalse:
		_aad.Value = "\u0066\u0061\u006cs\u0065"
	}
	return _aad, nil
}

func (_gbb ST_TrueFalseBlank) String() string {
	switch _gbb {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u0066"
	case 3:
		return "\u0074\u0072\u0075\u0065"
	case 4:
		return "\u0066\u0061\u006cs\u0065"
	case 6:
		return "\u0054\u0072\u0075\u0065"
	case 7:
		return "\u0046\u0061\u006cs\u0065"
	}
	return ""
}

func (_af ST_AlgClass) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_af.String(), start)
}

const (
	ST_AlgClassUnset  ST_AlgClass = 0
	ST_AlgClassHash   ST_AlgClass = 1
	ST_AlgClassCustom ST_AlgClass = 2
)

func (_fgc ST_TrueFalseBlank) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_dca := _e.Attr{}
	_dca.Name = name
	switch _fgc {
	case ST_TrueFalseBlankUnset:
		_dca.Value = ""
	case ST_TrueFalseBlankT:
		_dca.Value = "\u0074"
	case ST_TrueFalseBlankF:
		_dca.Value = "\u0066"
	case ST_TrueFalseBlankTrue:
		_dca.Value = "\u0074\u0072\u0075\u0065"
	case ST_TrueFalseBlankFalse:
		_dca.Value = "\u0066\u0061\u006cs\u0065"
	case ST_TrueFalseBlankTrue_:
		_dca.Value = "\u0054\u0072\u0075\u0065"
	case ST_TrueFalseBlankFalse_:
		_dca.Value = "\u0046\u0061\u006cs\u0065"
	}
	return _dca, nil
}

func (_dbc ST_YAlign) String() string {
	switch _dbc {
	case 0:
		return ""
	case 1:
		return "\u0069\u006e\u006c\u0069\u006e\u0065"
	case 2:
		return "\u0074\u006f\u0070"
	case 3:
		return "\u0063\u0065\u006e\u0074\u0065\u0072"
	case 4:
		return "\u0062\u006f\u0074\u0074\u006f\u006d"
	case 5:
		return "\u0069\u006e\u0073\u0069\u0064\u0065"
	case 6:
		return "\u006fu\u0074\u0073\u0069\u0064\u0065"
	}
	return ""
}

func (_dfe ST_ConformanceClass) Validate() error { return _dfe.ValidateWithPath("") }

const (
	ST_CalendarTypeUnset                ST_CalendarType = 0
	ST_CalendarTypeGregorian            ST_CalendarType = 1
	ST_CalendarTypeGregorianUs          ST_CalendarType = 2
	ST_CalendarTypeGregorianMeFrench    ST_CalendarType = 3
	ST_CalendarTypeGregorianArabic      ST_CalendarType = 4
	ST_CalendarTypeHijri                ST_CalendarType = 5
	ST_CalendarTypeHebrew               ST_CalendarType = 6
	ST_CalendarTypeTaiwan               ST_CalendarType = 7
	ST_CalendarTypeJapan                ST_CalendarType = 8
	ST_CalendarTypeThai                 ST_CalendarType = 9
	ST_CalendarTypeKorea                ST_CalendarType = 10
	ST_CalendarTypeSaka                 ST_CalendarType = 11
	ST_CalendarTypeGregorianXlitEnglish ST_CalendarType = 12
	ST_CalendarTypeGregorianXlitFrench  ST_CalendarType = 13
	ST_CalendarTypeNone                 ST_CalendarType = 14
)

const (
	ST_AlgTypeUnset   ST_AlgType = 0
	ST_AlgTypeTypeAny ST_AlgType = 1
	ST_AlgTypeCustom  ST_AlgType = 2
)

func (_gg ST_CryptProv) MarshalXML(e *_e.Encoder, start _e.StartElement) error {
	return e.EncodeElement(_gg.String(), start)
}

func (_edd ST_OnOff1) String() string {
	switch _edd {
	case 0:
		return ""
	case 1:
		return "\u006f\u006e"
	case 2:
		return "\u006f\u0066\u0066"
	}
	return ""
}

var ST_PercentagePatternRe = _ec.MustCompile(ST_PercentagePattern)

type ST_OnOff1 byte

func (_gbg ST_ConformanceClass) ValidateWithPath(path string) error {
	switch _gbg {
	case 0, 1, 2:
	default:
		return _ea.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gbg))
	}
	return nil
}

func (_aef ST_XAlign) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_aaa := _e.Attr{}
	_aaa.Name = name
	switch _aef {
	case ST_XAlignUnset:
		_aaa.Value = ""
	case ST_XAlignLeft:
		_aaa.Value = "\u006c\u0065\u0066\u0074"
	case ST_XAlignCenter:
		_aaa.Value = "\u0063\u0065\u006e\u0074\u0065\u0072"
	case ST_XAlignRight:
		_aaa.Value = "\u0072\u0069\u0067h\u0074"
	case ST_XAlignInside:
		_aaa.Value = "\u0069\u006e\u0073\u0069\u0064\u0065"
	case ST_XAlignOutside:
		_aaa.Value = "\u006fu\u0074\u0073\u0069\u0064\u0065"
	}
	return _aaa, nil
}

func (_baf ST_AlgType) MarshalXMLAttr(name _e.Name) (_e.Attr, error) {
	_cgd := _e.Attr{}
	_cgd.Name = name
	switch _baf {
	case ST_AlgTypeUnset:
		_cgd.Value = ""
	case ST_AlgTypeTypeAny:
		_cgd.Value = "\u0074y\u0070\u0065\u0041\u006e\u0079"
	case ST_AlgTypeCustom:
		_cgd.Value = "\u0063\u0075\u0073\u0074\u006f\u006d"
	}
	return _cgd, nil
}

func (_bfd ST_OnOff1) Validate() error { return _bfd.ValidateWithPath("") }

func (_bdg ST_TrueFalse) String() string {
	switch _bdg {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u0066"
	case 3:
		return "\u0074\u0072\u0075\u0065"
	case 4:
		return "\u0066\u0061\u006cs\u0065"
	}
	return ""
}
