package word

import (
	_g "encoding/xml"
	_d "fmt"
	_f "github.com/arcpd/msword"
	_a "strconv"
)

func (_bg *Bordertop) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bg.CT_Border = *NewCT_Border()
	for _, _ca := range start.Attr {
		if _ca.Name.Local == "\u0074\u0079\u0070\u0065" {
			_bg.TypeAttr.UnmarshalXMLAttr(_ca)
			continue
		}
		if _ca.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_ed, _cff := _a.ParseUint(_ca.Value, 10, 32)
			if _cff != nil {
				return _cff
			}
			_cfg := uint32(_ed)
			_bg.WidthAttr = &_cfg
			continue
		}
		if _ca.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_bg.ShadowAttr.UnmarshalXMLAttr(_ca)
			continue
		}
	}
	for {
		_eg, _cbc := d.Token()
		if _cbc != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020B\u006f\u0072\u0064\u0065\u0072t\u006f\u0070\u003a\u0020\u0025\u0073", _cbc)
		}
		if _daa, _bd := _eg.(_g.EndElement); _bd && _daa.Name == start.Name {
			break
		}
	}
	return nil
}

const (
	ST_WrapSideUnset   ST_WrapSide = 0
	ST_WrapSideBoth    ST_WrapSide = 1
	ST_WrapSideLeft    ST_WrapSide = 2
	ST_WrapSideRight   ST_WrapSide = 3
	ST_WrapSideLargest ST_WrapSide = 4
)
const (
	ST_BorderTypeUnset                 ST_BorderType = 0
	ST_BorderTypeNone                  ST_BorderType = 1
	ST_BorderTypeSingle                ST_BorderType = 2
	ST_BorderTypeThick                 ST_BorderType = 3
	ST_BorderTypeDouble                ST_BorderType = 4
	ST_BorderTypeHairline              ST_BorderType = 5
	ST_BorderTypeDot                   ST_BorderType = 6
	ST_BorderTypeDash                  ST_BorderType = 7
	ST_BorderTypeDotDash               ST_BorderType = 8
	ST_BorderTypeDashDotDot            ST_BorderType = 9
	ST_BorderTypeTriple                ST_BorderType = 10
	ST_BorderTypeThinThickSmall        ST_BorderType = 11
	ST_BorderTypeThickThinSmall        ST_BorderType = 12
	ST_BorderTypeThickBetweenThinSmall ST_BorderType = 13
	ST_BorderTypeThinThick             ST_BorderType = 14
	ST_BorderTypeThickThin             ST_BorderType = 15
	ST_BorderTypeThickBetweenThin      ST_BorderType = 16
	ST_BorderTypeThinThickLarge        ST_BorderType = 17
	ST_BorderTypeThickThinLarge        ST_BorderType = 18
	ST_BorderTypeThickBetweenThinLarge ST_BorderType = 19
	ST_BorderTypeWave                  ST_BorderType = 20
	ST_BorderTypeDoubleWave            ST_BorderType = 21
	ST_BorderTypeDashedSmall           ST_BorderType = 22
	ST_BorderTypeDashDotStroked        ST_BorderType = 23
	ST_BorderTypeThreeDEmboss          ST_BorderType = 24
	ST_BorderTypeThreeDEngrave         ST_BorderType = 25
	ST_BorderTypeHTMLOutset            ST_BorderType = 26
	ST_BorderTypeHTMLInset             ST_BorderType = 27
)

func (_gde *Borderright) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_gde.CT_Border = *NewCT_Border()
	for _, _fca := range start.Attr {
		if _fca.Name.Local == "\u0074\u0079\u0070\u0065" {
			_gde.TypeAttr.UnmarshalXMLAttr(_fca)
			continue
		}
		if _fca.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_acd, _da := _a.ParseUint(_fca.Value, 10, 32)
			if _da != nil {
				return _da
			}
			_adf := uint32(_acd)
			_gde.WidthAttr = &_adf
			continue
		}
		if _fca.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_gde.ShadowAttr.UnmarshalXMLAttr(_fca)
			continue
		}
	}
	for {
		_fa, _ff := d.Token()
		if _ff != nil {
			return _d.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0042\u006f\u0072\u0064\u0065\u0072\u0072\u0069\u0067\u0068\u0074\u003a\u0020\u0025\u0073", _ff)
		}
		if _fcag, _cga := _fa.(_g.EndElement); _cga && _fcag.Name == start.Name {
			break
		}
	}
	return nil
}
func (_ced *Borderright) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "b\u006f\u0072\u0064\u0065\u0072\u0072\u0069\u0067\u0068\u0074"
	return _ced.CT_Border.MarshalXML(e, start)
}
func NewBorderright() *Borderright {
	_fbe := &Borderright{}
	_fbe.CT_Border = *NewCT_Border()
	return _fbe
}
func (_dcaf ST_WrapType) ValidateWithPath(path string) error {
	switch _dcaf {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dcaf))
	}
	return nil
}
func (_cea ST_VerticalAnchor) ValidateWithPath(path string) error {
	switch _cea {
	case 0, 1, 2, 3, 4:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cea))
	}
	return nil
}
func (_bfa ST_WrapType) String() string {
	switch _bfa {
	case 0:
		return ""
	case 1:
		return "\u0074\u006f\u0070A\u006e\u0064\u0042\u006f\u0074\u0074\u006f\u006d"
	case 2:
		return "\u0073\u0071\u0075\u0061\u0072\u0065"
	case 3:
		return "\u006e\u006f\u006e\u0065"
	case 4:
		return "\u0074\u0069\u0067h\u0074"
	case 5:
		return "\u0074h\u0072\u006f\u0075\u0067\u0068"
	}
	return ""
}
func NewBordertop() *Bordertop { _aeg := &Bordertop{}; _aeg.CT_Border = *NewCT_Border(); return _aeg }
func (_gd *Borderbottom) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_gd.CT_Border = *NewCT_Border()
	for _, _ge := range start.Attr {
		if _ge.Name.Local == "\u0074\u0079\u0070\u0065" {
			_gd.TypeAttr.UnmarshalXMLAttr(_ge)
			continue
		}
		if _ge.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_ab, _gc := _a.ParseUint(_ge.Value, 10, 32)
			if _gc != nil {
				return _gc
			}
			_cg := uint32(_ab)
			_gd.WidthAttr = &_cg
			continue
		}
		if _ge.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_gd.ShadowAttr.UnmarshalXMLAttr(_ge)
			continue
		}
	}
	for {
		_e, _df := d.Token()
		if _df != nil {
			return _d.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0042\u006f\u0072\u0064\u0065r\u0062\u006f\u0074\u0074\u006f\u006d\u003a\u0020\u0025\u0073", _df)
		}
		if _eb, _ae := _e.(_g.EndElement); _ae && _eb.Name == start.Name {
			break
		}
	}
	return nil
}

type ST_WrapType byte

func (_efa *ST_BorderType) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bae, _gec := d.Token()
	if _gec != nil {
		return _gec
	}
	if _cedc, _eec := _bae.(_g.EndElement); _eec && _cedc.Name == start.Name {
		*_efa = 1
		return nil
	}
	if _cbba, _beg := _bae.(_g.CharData); !_beg {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bae)
	} else {
		switch string(_cbba) {
		case "":
			*_efa = 0
		case "\u006e\u006f\u006e\u0065":
			*_efa = 1
		case "\u0073\u0069\u006e\u0067\u006c\u0065":
			*_efa = 2
		case "\u0074\u0068\u0069c\u006b":
			*_efa = 3
		case "\u0064\u006f\u0075\u0062\u006c\u0065":
			*_efa = 4
		case "\u0068\u0061\u0069\u0072\u006c\u0069\u006e\u0065":
			*_efa = 5
		case "\u0064\u006f\u0074":
			*_efa = 6
		case "\u0064\u0061\u0073\u0068":
			*_efa = 7
		case "\u0064o\u0074\u0044\u0061\u0073\u0068":
			*_efa = 8
		case "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0044\u006f\u0074":
			*_efa = 9
		case "\u0074\u0072\u0069\u0070\u006c\u0065":
			*_efa = 10
		case "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bS\u006d\u0061\u006c\u006c":
			*_efa = 11
		case "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eS\u006d\u0061\u006c\u006c":
			*_efa = 12
		case "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u0053\u006d\u0061\u006c\u006c":
			*_efa = 13
		case "\u0074h\u0069\u006e\u0054\u0068\u0069\u0063k":
			*_efa = 14
		case "\u0074h\u0069\u0063\u006b\u0054\u0068\u0069n":
			*_efa = 15
		case "\u0074\u0068i\u0063\u006b\u0042e\u0074\u0077\u0065\u0065\u006e\u0054\u0068\u0069\u006e":
			*_efa = 16
		case "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bL\u0061\u0072\u0067\u0065":
			*_efa = 17
		case "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eL\u0061\u0072\u0067\u0065":
			*_efa = 18
		case "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u004c\u0061\u0072\u0067\u0065":
			*_efa = 19
		case "\u0077\u0061\u0076\u0065":
			*_efa = 20
		case "\u0064\u006f\u0075\u0062\u006c\u0065\u0057\u0061\u0076\u0065":
			*_efa = 21
		case "d\u0061\u0073\u0068\u0065\u0064\u0053\u006d\u0061\u006c\u006c":
			*_efa = 22
		case "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0053\u0074r\u006f\u006b\u0065\u0064":
			*_efa = 23
		case "\u0074\u0068\u0072e\u0065\u0044\u0045\u006d\u0062\u006f\u0073\u0073":
			*_efa = 24
		case "\u0074\u0068\u0072\u0065\u0065\u0044\u0045\u006e\u0067\u0072\u0061\u0076\u0065":
			*_efa = 25
		case "\u0048\u0054\u004d\u004c\u004f\u0075\u0074\u0073\u0065\u0074":
			*_efa = 26
		case "\u0048T\u004d\u004c\u0049\u006e\u0073\u0065t":
			*_efa = 27
		}
	}
	_bae, _gec = d.Token()
	if _gec != nil {
		return _gec
	}
	if _eacc, _bebd := _bae.(_g.EndElement); _bebd && _eacc.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bae)
}
func (_ddg ST_BorderType) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_bcg := _g.Attr{}
	_bcg.Name = name
	switch _ddg {
	case ST_BorderTypeUnset:
		_bcg.Value = ""
	case ST_BorderTypeNone:
		_bcg.Value = "\u006e\u006f\u006e\u0065"
	case ST_BorderTypeSingle:
		_bcg.Value = "\u0073\u0069\u006e\u0067\u006c\u0065"
	case ST_BorderTypeThick:
		_bcg.Value = "\u0074\u0068\u0069c\u006b"
	case ST_BorderTypeDouble:
		_bcg.Value = "\u0064\u006f\u0075\u0062\u006c\u0065"
	case ST_BorderTypeHairline:
		_bcg.Value = "\u0068\u0061\u0069\u0072\u006c\u0069\u006e\u0065"
	case ST_BorderTypeDot:
		_bcg.Value = "\u0064\u006f\u0074"
	case ST_BorderTypeDash:
		_bcg.Value = "\u0064\u0061\u0073\u0068"
	case ST_BorderTypeDotDash:
		_bcg.Value = "\u0064o\u0074\u0044\u0061\u0073\u0068"
	case ST_BorderTypeDashDotDot:
		_bcg.Value = "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0044\u006f\u0074"
	case ST_BorderTypeTriple:
		_bcg.Value = "\u0074\u0072\u0069\u0070\u006c\u0065"
	case ST_BorderTypeThinThickSmall:
		_bcg.Value = "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bS\u006d\u0061\u006c\u006c"
	case ST_BorderTypeThickThinSmall:
		_bcg.Value = "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eS\u006d\u0061\u006c\u006c"
	case ST_BorderTypeThickBetweenThinSmall:
		_bcg.Value = "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u0053\u006d\u0061\u006c\u006c"
	case ST_BorderTypeThinThick:
		_bcg.Value = "\u0074h\u0069\u006e\u0054\u0068\u0069\u0063k"
	case ST_BorderTypeThickThin:
		_bcg.Value = "\u0074h\u0069\u0063\u006b\u0054\u0068\u0069n"
	case ST_BorderTypeThickBetweenThin:
		_bcg.Value = "\u0074\u0068i\u0063\u006b\u0042e\u0074\u0077\u0065\u0065\u006e\u0054\u0068\u0069\u006e"
	case ST_BorderTypeThinThickLarge:
		_bcg.Value = "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bL\u0061\u0072\u0067\u0065"
	case ST_BorderTypeThickThinLarge:
		_bcg.Value = "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eL\u0061\u0072\u0067\u0065"
	case ST_BorderTypeThickBetweenThinLarge:
		_bcg.Value = "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u004c\u0061\u0072\u0067\u0065"
	case ST_BorderTypeWave:
		_bcg.Value = "\u0077\u0061\u0076\u0065"
	case ST_BorderTypeDoubleWave:
		_bcg.Value = "\u0064\u006f\u0075\u0062\u006c\u0065\u0057\u0061\u0076\u0065"
	case ST_BorderTypeDashedSmall:
		_bcg.Value = "d\u0061\u0073\u0068\u0065\u0064\u0053\u006d\u0061\u006c\u006c"
	case ST_BorderTypeDashDotStroked:
		_bcg.Value = "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0053\u0074r\u006f\u006b\u0065\u0064"
	case ST_BorderTypeThreeDEmboss:
		_bcg.Value = "\u0074\u0068\u0072e\u0065\u0044\u0045\u006d\u0062\u006f\u0073\u0073"
	case ST_BorderTypeThreeDEngrave:
		_bcg.Value = "\u0074\u0068\u0072\u0065\u0065\u0044\u0045\u006e\u0067\u0072\u0061\u0076\u0065"
	case ST_BorderTypeHTMLOutset:
		_bcg.Value = "\u0048\u0054\u004d\u004c\u004f\u0075\u0074\u0073\u0065\u0074"
	case ST_BorderTypeHTMLInset:
		_bcg.Value = "\u0048T\u004d\u004c\u0049\u006e\u0073\u0065t"
	}
	return _bcg, nil
}

// Validate validates the Anchorlock and its children
func (_ggb *Anchorlock) Validate() error {
	return _ggb.ValidateWithPath("\u0041\u006e\u0063\u0068\u006f\u0072\u006c\u006f\u0063\u006b")
}
func (_caf ST_BorderShadow) ValidateWithPath(path string) error {
	switch _caf {
	case 0, 1, 2, 3, 4:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_caf))
	}
	return nil
}

type ST_VerticalAnchor byte

// ValidateWithPath validates the CT_Border and its children, prefixing error messages with path
func (_cba *CT_Border) ValidateWithPath(path string) error {
	if _eff := _cba.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _eff != nil {
		return _eff
	}
	if _geg := _cba.ShadowAttr.ValidateWithPath(path + "/\u0053\u0068\u0061\u0064\u006f\u0077\u0041\u0074\u0074\u0072"); _geg != nil {
		return _geg
	}
	return nil
}

// Validate validates the CT_AnchorLock and its children
func (_ceb *CT_AnchorLock) Validate() error {
	return _ceb.ValidateWithPath("\u0043\u0054\u005f\u0041\u006e\u0063\u0068\u006f\u0072\u004c\u006f\u0063\u006b")
}
func NewBorderleft() *Borderleft             { _ce := &Borderleft{}; _ce.CT_Border = *NewCT_Border(); return _ce }
func (_egb ST_BorderShadow) Validate() error { return _egb.ValidateWithPath("") }
func (_ege *ST_BorderShadow) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cac, _ede := d.Token()
	if _ede != nil {
		return _ede
	}
	if _ccf, _effg := _cac.(_g.EndElement); _effg && _ccf.Name == start.Name {
		*_ege = 1
		return nil
	}
	if _cd, _gcge := _cac.(_g.CharData); !_gcge {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cac)
	} else {
		switch string(_cd) {
		case "":
			*_ege = 0
		case "\u0074":
			*_ege = 1
		case "\u0074\u0072\u0075\u0065":
			*_ege = 2
		case "\u0066":
			*_ege = 3
		case "\u0066\u0061\u006cs\u0065":
			*_ege = 4
		}
	}
	_cac, _ede = d.Token()
	if _ede != nil {
		return _ede
	}
	if _gee, _dca := _cac.(_g.EndElement); _dca && _gee.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cac)
}

// Validate validates the CT_Wrap and its children
func (_acc *CT_Wrap) Validate() error {
	return _acc.ValidateWithPath("\u0043T\u005f\u0057\u0072\u0061\u0070")
}

type CT_AnchorLock struct{}

func (_aa *Anchorlock) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0061\u006e\u0063\u0068\u006f\u0072\u006c\u006f\u0063\u006b"
	return _aa.CT_AnchorLock.MarshalXML(e, start)
}
func (_bba ST_HorizontalAnchor) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_bba.String(), start)
}
func (_ddae ST_VerticalAnchor) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_bda := _g.Attr{}
	_bda.Name = name
	switch _ddae {
	case ST_VerticalAnchorUnset:
		_bda.Value = ""
	case ST_VerticalAnchorMargin:
		_bda.Value = "\u006d\u0061\u0072\u0067\u0069\u006e"
	case ST_VerticalAnchorPage:
		_bda.Value = "\u0070\u0061\u0067\u0065"
	case ST_VerticalAnchorText:
		_bda.Value = "\u0074\u0065\u0078\u0074"
	case ST_VerticalAnchorLine:
		_bda.Value = "\u006c\u0069\u006e\u0065"
	}
	return _bda, nil
}

type Borderleft struct{ CT_Border }

func (_fdb *Wrap) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_fdb.CT_Wrap = *NewCT_Wrap()
	for _, _abb := range start.Attr {
		if _abb.Name.Local == "\u0074\u0079\u0070\u0065" {
			_fdb.TypeAttr.UnmarshalXMLAttr(_abb)
			continue
		}
		if _abb.Name.Local == "\u0073\u0069\u0064\u0065" {
			_fdb.SideAttr.UnmarshalXMLAttr(_abb)
			continue
		}
		if _abb.Name.Local == "\u0061n\u0063\u0068\u006f\u0072\u0078" {
			_fdb.AnchorxAttr.UnmarshalXMLAttr(_abb)
			continue
		}
		if _abb.Name.Local == "\u0061n\u0063\u0068\u006f\u0072\u0079" {
			_fdb.AnchoryAttr.UnmarshalXMLAttr(_abb)
			continue
		}
	}
	for {
		_ebg, _dac := d.Token()
		if _dac != nil {
			return _d.Errorf("\u0070\u0061r\u0073\u0069\u006eg\u0020\u0057\u0072\u0061\u0070\u003a\u0020\u0025\u0073", _dac)
		}
		if _bdg, _fgb := _ebg.(_g.EndElement); _fgb && _bdg.Name == start.Name {
			break
		}
	}
	return nil
}
func (_abg *ST_BorderType) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_abg = 0
	case "\u006e\u006f\u006e\u0065":
		*_abg = 1
	case "\u0073\u0069\u006e\u0067\u006c\u0065":
		*_abg = 2
	case "\u0074\u0068\u0069c\u006b":
		*_abg = 3
	case "\u0064\u006f\u0075\u0062\u006c\u0065":
		*_abg = 4
	case "\u0068\u0061\u0069\u0072\u006c\u0069\u006e\u0065":
		*_abg = 5
	case "\u0064\u006f\u0074":
		*_abg = 6
	case "\u0064\u0061\u0073\u0068":
		*_abg = 7
	case "\u0064o\u0074\u0044\u0061\u0073\u0068":
		*_abg = 8
	case "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0044\u006f\u0074":
		*_abg = 9
	case "\u0074\u0072\u0069\u0070\u006c\u0065":
		*_abg = 10
	case "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bS\u006d\u0061\u006c\u006c":
		*_abg = 11
	case "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eS\u006d\u0061\u006c\u006c":
		*_abg = 12
	case "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u0053\u006d\u0061\u006c\u006c":
		*_abg = 13
	case "\u0074h\u0069\u006e\u0054\u0068\u0069\u0063k":
		*_abg = 14
	case "\u0074h\u0069\u0063\u006b\u0054\u0068\u0069n":
		*_abg = 15
	case "\u0074\u0068i\u0063\u006b\u0042e\u0074\u0077\u0065\u0065\u006e\u0054\u0068\u0069\u006e":
		*_abg = 16
	case "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bL\u0061\u0072\u0067\u0065":
		*_abg = 17
	case "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eL\u0061\u0072\u0067\u0065":
		*_abg = 18
	case "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u004c\u0061\u0072\u0067\u0065":
		*_abg = 19
	case "\u0077\u0061\u0076\u0065":
		*_abg = 20
	case "\u0064\u006f\u0075\u0062\u006c\u0065\u0057\u0061\u0076\u0065":
		*_abg = 21
	case "d\u0061\u0073\u0068\u0065\u0064\u0053\u006d\u0061\u006c\u006c":
		*_abg = 22
	case "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0053\u0074r\u006f\u006b\u0065\u0064":
		*_abg = 23
	case "\u0074\u0068\u0072e\u0065\u0044\u0045\u006d\u0062\u006f\u0073\u0073":
		*_abg = 24
	case "\u0074\u0068\u0072\u0065\u0065\u0044\u0045\u006e\u0067\u0072\u0061\u0076\u0065":
		*_abg = 25
	case "\u0048\u0054\u004d\u004c\u004f\u0075\u0074\u0073\u0065\u0074":
		*_abg = 26
	case "\u0048T\u004d\u004c\u0049\u006e\u0073\u0065t":
		*_abg = 27
	}
	return nil
}

// Validate validates the Borderleft and its children
func (_afa *Borderleft) Validate() error {
	return _afa.ValidateWithPath("\u0042\u006f\u0072\u0064\u0065\u0072\u006c\u0065\u0066\u0074")
}
func NewAnchorlock() *Anchorlock {
	_c := &Anchorlock{}
	_c.CT_AnchorLock = *NewCT_AnchorLock()
	return _c
}
func (_eadg *ST_WrapType) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ebe, _gab := d.Token()
	if _gab != nil {
		return _gab
	}
	if _afc, _dbgg := _ebe.(_g.EndElement); _dbgg && _afc.Name == start.Name {
		*_eadg = 1
		return nil
	}
	if _bgdc, _gac := _ebe.(_g.CharData); !_gac {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ebe)
	} else {
		switch string(_bgdc) {
		case "":
			*_eadg = 0
		case "\u0074\u006f\u0070A\u006e\u0064\u0042\u006f\u0074\u0074\u006f\u006d":
			*_eadg = 1
		case "\u0073\u0071\u0075\u0061\u0072\u0065":
			*_eadg = 2
		case "\u006e\u006f\u006e\u0065":
			*_eadg = 3
		case "\u0074\u0069\u0067h\u0074":
			*_eadg = 4
		case "\u0074h\u0072\u006f\u0075\u0067\u0068":
			*_eadg = 5
		}
	}
	_ebe, _gab = d.Token()
	if _gab != nil {
		return _gab
	}
	if _gag, _gbec := _ebe.(_g.EndElement); _gbec && _gag.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ebe)
}

type Anchorlock struct{ CT_AnchorLock }

func (_dae ST_BorderType) String() string {
	switch _dae {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0073\u0069\u006e\u0067\u006c\u0065"
	case 3:
		return "\u0074\u0068\u0069c\u006b"
	case 4:
		return "\u0064\u006f\u0075\u0062\u006c\u0065"
	case 5:
		return "\u0068\u0061\u0069\u0072\u006c\u0069\u006e\u0065"
	case 6:
		return "\u0064\u006f\u0074"
	case 7:
		return "\u0064\u0061\u0073\u0068"
	case 8:
		return "\u0064o\u0074\u0044\u0061\u0073\u0068"
	case 9:
		return "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0044\u006f\u0074"
	case 10:
		return "\u0074\u0072\u0069\u0070\u006c\u0065"
	case 11:
		return "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bS\u006d\u0061\u006c\u006c"
	case 12:
		return "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eS\u006d\u0061\u006c\u006c"
	case 13:
		return "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u0053\u006d\u0061\u006c\u006c"
	case 14:
		return "\u0074h\u0069\u006e\u0054\u0068\u0069\u0063k"
	case 15:
		return "\u0074h\u0069\u0063\u006b\u0054\u0068\u0069n"
	case 16:
		return "\u0074\u0068i\u0063\u006b\u0042e\u0074\u0077\u0065\u0065\u006e\u0054\u0068\u0069\u006e"
	case 17:
		return "\u0074\u0068\u0069\u006e\u0054\u0068\u0069\u0063\u006bL\u0061\u0072\u0067\u0065"
	case 18:
		return "\u0074\u0068\u0069\u0063\u006b\u0054\u0068\u0069\u006eL\u0061\u0072\u0067\u0065"
	case 19:
		return "t\u0068\u0069\u0063\u006bBe\u0074w\u0065\u0065\u006e\u0054\u0068i\u006e\u004c\u0061\u0072\u0067\u0065"
	case 20:
		return "\u0077\u0061\u0076\u0065"
	case 21:
		return "\u0064\u006f\u0075\u0062\u006c\u0065\u0057\u0061\u0076\u0065"
	case 22:
		return "d\u0061\u0073\u0068\u0065\u0064\u0053\u006d\u0061\u006c\u006c"
	case 23:
		return "\u0064\u0061\u0073\u0068\u0044\u006f\u0074\u0053\u0074r\u006f\u006b\u0065\u0064"
	case 24:
		return "\u0074\u0068\u0072e\u0065\u0044\u0045\u006d\u0062\u006f\u0073\u0073"
	case 25:
		return "\u0074\u0068\u0072\u0065\u0065\u0044\u0045\u006e\u0067\u0072\u0061\u0076\u0065"
	case 26:
		return "\u0048\u0054\u004d\u004c\u004f\u0075\u0074\u0073\u0065\u0074"
	case 27:
		return "\u0048T\u004d\u004c\u0049\u006e\u0073\u0065t"
	}
	return ""
}

type CT_Wrap struct {
	TypeAttr    ST_WrapType
	SideAttr    ST_WrapSide
	AnchorxAttr ST_HorizontalAnchor
	AnchoryAttr ST_VerticalAnchor
}

func (_cfc ST_BorderShadow) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_gdd := _g.Attr{}
	_gdd.Name = name
	switch _cfc {
	case ST_BorderShadowUnset:
		_gdd.Value = ""
	case ST_BorderShadowT:
		_gdd.Value = "\u0074"
	case ST_BorderShadowTrue:
		_gdd.Value = "\u0074\u0072\u0075\u0065"
	case ST_BorderShadowF:
		_gdd.Value = "\u0066"
	case ST_BorderShadowFalse:
		_gdd.Value = "\u0066\u0061\u006cs\u0065"
	}
	return _gdd, nil
}

type ST_BorderType byte

func (_fdc *ST_WrapSide) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_eeff, _dde := d.Token()
	if _dde != nil {
		return _dde
	}
	if _daea, _fdf := _eeff.(_g.EndElement); _fdf && _daea.Name == start.Name {
		*_fdc = 1
		return nil
	}
	if _acg, _gebf := _eeff.(_g.CharData); !_gebf {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eeff)
	} else {
		switch string(_acg) {
		case "":
			*_fdc = 0
		case "\u0062\u006f\u0074\u0068":
			*_fdc = 1
		case "\u006c\u0065\u0066\u0074":
			*_fdc = 2
		case "\u0072\u0069\u0067h\u0074":
			*_fdc = 3
		case "\u006ca\u0072\u0067\u0065\u0073\u0074":
			*_fdc = 4
		}
	}
	_eeff, _dde = d.Token()
	if _dde != nil {
		return _dde
	}
	if _egeg, _egef := _eeff.(_g.EndElement); _egef && _egeg.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eeff)
}

// ValidateWithPath validates the Anchorlock and its children, prefixing error messages with path
func (_af *Anchorlock) ValidateWithPath(path string) error {
	if _ac := _af.CT_AnchorLock.ValidateWithPath(path); _ac != nil {
		return _ac
	}
	return nil
}

// Validate validates the Borderright and its children
func (_fcb *Borderright) Validate() error {
	return _fcb.ValidateWithPath("B\u006f\u0072\u0064\u0065\u0072\u0072\u0069\u0067\u0068\u0074")
}
func (_cafg ST_WrapSide) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_efc := _g.Attr{}
	_efc.Name = name
	switch _cafg {
	case ST_WrapSideUnset:
		_efc.Value = ""
	case ST_WrapSideBoth:
		_efc.Value = "\u0062\u006f\u0074\u0068"
	case ST_WrapSideLeft:
		_efc.Value = "\u006c\u0065\u0066\u0074"
	case ST_WrapSideRight:
		_efc.Value = "\u0072\u0069\u0067h\u0074"
	case ST_WrapSideLargest:
		_efc.Value = "\u006ca\u0072\u0067\u0065\u0073\u0074"
	}
	return _efc, nil
}
func (_cace ST_WrapType) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_cace.String(), start)
}
func (_cb *Anchorlock) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cb.CT_AnchorLock = *NewCT_AnchorLock()
	for {
		_fe, _ag := d.Token()
		if _ag != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0041\u006e\u0063\u0068\u006f\u0072\u006c\u006f\u0063\u006b\u003a\u0020%\u0073", _ag)
		}
		if _dc, _gg := _fe.(_g.EndElement); _gg && _dc.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the Borderbottom and its children, prefixing error messages with path
func (_ec *Borderbottom) ValidateWithPath(path string) error {
	if _ecc := _ec.CT_Border.ValidateWithPath(path); _ecc != nil {
		return _ecc
	}
	return nil
}

// ValidateWithPath validates the Borderleft and its children, prefixing error messages with path
func (_cf *Borderleft) ValidateWithPath(path string) error {
	if _ga := _cf.CT_Border.ValidateWithPath(path); _ga != nil {
		return _ga
	}
	return nil
}
func (_bca ST_BorderShadow) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_bca.String(), start)
}
func (_dbb ST_BorderShadow) String() string {
	switch _dbb {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u0074\u0072\u0075\u0065"
	case 3:
		return "\u0066"
	case 4:
		return "\u0066\u0061\u006cs\u0065"
	}
	return ""
}

type ST_HorizontalAnchor byte

func (_dce *Bordertop) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0062o\u0072\u0064\u0065\u0072\u0074\u006fp"
	return _dce.CT_Border.MarshalXML(e, start)
}
func NewWrap() *Wrap { _gbc := &Wrap{}; _gbc.CT_Wrap = *NewCT_Wrap(); return _gbc }

type CT_Border struct {
	TypeAttr   ST_BorderType
	WidthAttr  *uint32
	ShadowAttr ST_BorderShadow
}
type ST_BorderShadow byte

func NewCT_Border() *CT_Border { _dec := &CT_Border{}; return _dec }
func (_dbe ST_WrapType) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_aef := _g.Attr{}
	_aef.Name = name
	switch _dbe {
	case ST_WrapTypeUnset:
		_aef.Value = ""
	case ST_WrapTypeTopAndBottom:
		_aef.Value = "\u0074\u006f\u0070A\u006e\u0064\u0042\u006f\u0074\u0074\u006f\u006d"
	case ST_WrapTypeSquare:
		_aef.Value = "\u0073\u0071\u0075\u0061\u0072\u0065"
	case ST_WrapTypeNone:
		_aef.Value = "\u006e\u006f\u006e\u0065"
	case ST_WrapTypeTight:
		_aef.Value = "\u0074\u0069\u0067h\u0074"
	case ST_WrapTypeThrough:
		_aef.Value = "\u0074h\u0072\u006f\u0075\u0067\u0068"
	}
	return _aef, nil
}
func (_ggag ST_WrapType) Validate() error { return _ggag.ValidateWithPath("") }
func NewCT_AnchorLock() *CT_AnchorLock    { _fcc := &CT_AnchorLock{}; return _fcc }
func (_gcg *CT_Border) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _gcg.TypeAttr != ST_BorderTypeUnset {
		_fd, _bb := _gcg.TypeAttr.MarshalXMLAttr(_g.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _bb != nil {
			return _bb
		}
		start.Attr = append(start.Attr, _fd)
	}
	if _gcg.WidthAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0077\u0069\u0064t\u0068"}, Value: _d.Sprintf("\u0025\u0076", *_gcg.WidthAttr)})
	}
	if _gcg.ShadowAttr != ST_BorderShadowUnset {
		_fee, _gdf := _gcg.ShadowAttr.MarshalXMLAttr(_g.Name{Local: "\u0073\u0068\u0061\u0064\u006f\u0077"})
		if _gdf != nil {
			return _gdf
		}
		start.Attr = append(start.Attr, _fee)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_bfb *ST_WrapSide) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_bfb = 0
	case "\u0062\u006f\u0074\u0068":
		*_bfb = 1
	case "\u006c\u0065\u0066\u0074":
		*_bfb = 2
	case "\u0072\u0069\u0067h\u0074":
		*_bfb = 3
	case "\u006ca\u0072\u0067\u0065\u0073\u0074":
		*_bfb = 4
	}
	return nil
}

// ValidateWithPath validates the CT_Wrap and its children, prefixing error messages with path
func (_bee *CT_Wrap) ValidateWithPath(path string) error {
	if _efd := _bee.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _efd != nil {
		return _efd
	}
	if _bed := _bee.SideAttr.ValidateWithPath(path + "\u002fS\u0069\u0064\u0065\u0041\u0074\u0074r"); _bed != nil {
		return _bed
	}
	if _dd := _bee.AnchorxAttr.ValidateWithPath(path + "\u002f\u0041\u006ec\u0068\u006f\u0072\u0078\u0041\u0074\u0074\u0072"); _dd != nil {
		return _dd
	}
	if _aagd := _bee.AnchoryAttr.ValidateWithPath(path + "\u002f\u0041\u006ec\u0068\u006f\u0072\u0079\u0041\u0074\u0074\u0072"); _aagd != nil {
		return _aagd
	}
	return nil
}

const (
	ST_HorizontalAnchorUnset  ST_HorizontalAnchor = 0
	ST_HorizontalAnchorMargin ST_HorizontalAnchor = 1
	ST_HorizontalAnchorPage   ST_HorizontalAnchor = 2
	ST_HorizontalAnchorText   ST_HorizontalAnchor = 3
	ST_HorizontalAnchorChar   ST_HorizontalAnchor = 4
)

func (_ccdf *ST_BorderShadow) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_ccdf = 0
	case "\u0074":
		*_ccdf = 1
	case "\u0074\u0072\u0075\u0065":
		*_ccdf = 2
	case "\u0066":
		*_ccdf = 3
	case "\u0066\u0061\u006cs\u0065":
		*_ccdf = 4
	}
	return nil
}
func (_bcb ST_BorderType) ValidateWithPath(path string) error {
	switch _bcb {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bcb))
	}
	return nil
}
func NewBorderbottom() *Borderbottom {
	_de := &Borderbottom{}
	_de.CT_Border = *NewCT_Border()
	return _de
}
func (_gaaf *CT_Wrap) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _gaaf.TypeAttr != ST_WrapTypeUnset {
		_ccbc, _cbb := _gaaf.TypeAttr.MarshalXMLAttr(_g.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _cbb != nil {
			return _cbb
		}
		start.Attr = append(start.Attr, _ccbc)
	}
	if _gaaf.SideAttr != ST_WrapSideUnset {
		_bfgf, _ffg := _gaaf.SideAttr.MarshalXMLAttr(_g.Name{Local: "\u0073\u0069\u0064\u0065"})
		if _ffg != nil {
			return _ffg
		}
		start.Attr = append(start.Attr, _bfgf)
	}
	if _gaaf.AnchorxAttr != ST_HorizontalAnchorUnset {
		_eef, _fbc := _gaaf.AnchorxAttr.MarshalXMLAttr(_g.Name{Local: "\u0061n\u0063\u0068\u006f\u0072\u0078"})
		if _fbc != nil {
			return _fbc
		}
		start.Attr = append(start.Attr, _eef)
	}
	if _gaaf.AnchoryAttr != ST_VerticalAnchorUnset {
		_ead, _adfg := _gaaf.AnchoryAttr.MarshalXMLAttr(_g.Name{Local: "\u0061n\u0063\u0068\u006f\u0072\u0079"})
		if _adfg != nil {
			return _adfg
		}
		start.Attr = append(start.Attr, _ead)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_fag ST_HorizontalAnchor) Validate() error { return _fag.ValidateWithPath("") }

// ValidateWithPath validates the Wrap and its children, prefixing error messages with path
func (_efdg *Wrap) ValidateWithPath(path string) error {
	if _fce := _efdg.CT_Wrap.ValidateWithPath(path); _fce != nil {
		return _fce
	}
	return nil
}

const (
	ST_VerticalAnchorUnset  ST_VerticalAnchor = 0
	ST_VerticalAnchorMargin ST_VerticalAnchor = 1
	ST_VerticalAnchorPage   ST_VerticalAnchor = 2
	ST_VerticalAnchorText   ST_VerticalAnchor = 3
	ST_VerticalAnchorLine   ST_VerticalAnchor = 4
)

func (_dag *CT_Border) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _gbe := range start.Attr {
		if _gbe.Name.Local == "\u0074\u0079\u0070\u0065" {
			_dag.TypeAttr.UnmarshalXMLAttr(_gbe)
			continue
		}
		if _gbe.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_ccad, _ged := _a.ParseUint(_gbe.Value, 10, 32)
			if _ged != nil {
				return _ged
			}
			_gge := uint32(_ccad)
			_dag.WidthAttr = &_gge
			continue
		}
		if _gbe.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_dag.ShadowAttr.UnmarshalXMLAttr(_gbe)
			continue
		}
	}
	for {
		_aee, _ccb := d.Token()
		if _ccb != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0042\u006f\u0072d\u0065\u0072\u003a\u0020\u0025\u0073", _ccb)
		}
		if _fg, _ccd := _aee.(_g.EndElement); _ccd && _fg.Name == start.Name {
			break
		}
	}
	return nil
}
func (_gbd ST_WrapSide) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_gbd.String(), start)
}
func (_ef *CT_AnchorLock) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for {
		_cag, _ee := d.Token()
		if _ee != nil {
			return _d.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0043\u0054\u005f\u0041n\u0063h\u006fr\u004c\u006f\u0063\u006b\u003a\u0020\u0025s", _ee)
		}
		if _ecd, _bfe := _cag.(_g.EndElement); _bfe && _ecd.Name == start.Name {
			break
		}
	}
	return nil
}

const (
	ST_WrapTypeUnset        ST_WrapType = 0
	ST_WrapTypeTopAndBottom ST_WrapType = 1
	ST_WrapTypeSquare       ST_WrapType = 2
	ST_WrapTypeNone         ST_WrapType = 3
	ST_WrapTypeTight        ST_WrapType = 4
	ST_WrapTypeThrough      ST_WrapType = 5
)

func (_fgd *CT_Wrap) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _ggg := range start.Attr {
		if _ggg.Name.Local == "\u0074\u0079\u0070\u0065" {
			_fgd.TypeAttr.UnmarshalXMLAttr(_ggg)
			continue
		}
		if _ggg.Name.Local == "\u0073\u0069\u0064\u0065" {
			_fgd.SideAttr.UnmarshalXMLAttr(_ggg)
			continue
		}
		if _ggg.Name.Local == "\u0061n\u0063\u0068\u006f\u0072\u0078" {
			_fgd.AnchorxAttr.UnmarshalXMLAttr(_ggg)
			continue
		}
		if _ggg.Name.Local == "\u0061n\u0063\u0068\u006f\u0072\u0079" {
			_fgd.AnchoryAttr.UnmarshalXMLAttr(_ggg)
			continue
		}
	}
	for {
		_beb, _bc := d.Token()
		if _bc != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0057\u0072\u0061\u0070\u003a\u0020\u0025\u0073", _bc)
		}
		if _gff, _eac := _beb.(_g.EndElement); _eac && _gff.Name == start.Name {
			break
		}
	}
	return nil
}
func (_fed ST_VerticalAnchor) Validate() error { return _fed.ValidateWithPath("") }
func (_bcf *Wrap) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0077\u0072\u0061\u0070"
	return _bcf.CT_Wrap.MarshalXML(e, start)
}
func (_gbg *ST_HorizontalAnchor) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bedd, _acba := d.Token()
	if _acba != nil {
		return _acba
	}
	if _ggbf, _cdc := _bedd.(_g.EndElement); _cdc && _ggbf.Name == start.Name {
		*_gbg = 1
		return nil
	}
	if _cfgb, _fbg := _bedd.(_g.CharData); !_fbg {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bedd)
	} else {
		switch string(_cfgb) {
		case "":
			*_gbg = 0
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			*_gbg = 1
		case "\u0070\u0061\u0067\u0065":
			*_gbg = 2
		case "\u0074\u0065\u0078\u0074":
			*_gbg = 3
		case "\u0063\u0068\u0061\u0072":
			*_gbg = 4
		}
	}
	_bedd, _acba = d.Token()
	if _acba != nil {
		return _acba
	}
	if _dda, _bfag := _bedd.(_g.EndElement); _bfag && _dda.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bedd)
}
func (_gabf ST_HorizontalAnchor) ValidateWithPath(path string) error {
	switch _gabf {
	case 0, 1, 2, 3, 4:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gabf))
	}
	return nil
}

// ValidateWithPath validates the CT_AnchorLock and its children, prefixing error messages with path
func (_dbg *CT_AnchorLock) ValidateWithPath(path string) error { return nil }

type Bordertop struct{ CT_Border }
type ST_WrapSide byte

func (_fagd ST_VerticalAnchor) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_fagd.String(), start)
}
func NewCT_Wrap() *CT_Wrap { _aaa := &CT_Wrap{}; return _aaa }

// Validate validates the Bordertop and its children
func (_ebc *Bordertop) Validate() error {
	return _ebc.ValidateWithPath("\u0042o\u0072\u0064\u0065\u0072\u0074\u006fp")
}

// Validate validates the CT_Border and its children
func (_gaa *CT_Border) Validate() error {
	return _gaa.ValidateWithPath("\u0043T\u005f\u0042\u006f\u0072\u0064\u0065r")
}
func (_gef ST_WrapSide) String() string {
	switch _gef {
	case 0:
		return ""
	case 1:
		return "\u0062\u006f\u0074\u0068"
	case 2:
		return "\u006c\u0065\u0066\u0074"
	case 3:
		return "\u0072\u0069\u0067h\u0074"
	case 4:
		return "\u006ca\u0072\u0067\u0065\u0073\u0074"
	}
	return ""
}
func (_gae ST_VerticalAnchor) String() string {
	switch _gae {
	case 0:
		return ""
	case 1:
		return "\u006d\u0061\u0072\u0067\u0069\u006e"
	case 2:
		return "\u0070\u0061\u0067\u0065"
	case 3:
		return "\u0074\u0065\u0078\u0074"
	case 4:
		return "\u006c\u0069\u006e\u0065"
	}
	return ""
}
func (_cbab *ST_VerticalAnchor) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ddc, _eefb := d.Token()
	if _eefb != nil {
		return _eefb
	}
	if _dcg, _eefe := _ddc.(_g.EndElement); _eefe && _dcg.Name == start.Name {
		*_cbab = 1
		return nil
	}
	if _afde, _ddcg := _ddc.(_g.CharData); !_ddcg {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ddc)
	} else {
		switch string(_afde) {
		case "":
			*_cbab = 0
		case "\u006d\u0061\u0072\u0067\u0069\u006e":
			*_cbab = 1
		case "\u0070\u0061\u0067\u0065":
			*_cbab = 2
		case "\u0074\u0065\u0078\u0074":
			*_cbab = 3
		case "\u006c\u0069\u006e\u0065":
			*_cbab = 4
		}
	}
	_ddc, _eefb = d.Token()
	if _eefb != nil {
		return _eefb
	}
	if _cbcf, _gdc := _ddc.(_g.EndElement); _gdc && _cbcf.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ddc)
}

type Borderbottom struct{ CT_Border }

func (_fb *Borderleft) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0062\u006f\u0072\u0064\u0065\u0072\u006c\u0065\u0066\u0074"
	return _fb.CT_Border.MarshalXML(e, start)
}

// Validate validates the Borderbottom and its children
func (_afd *Borderbottom) Validate() error {
	return _afd.ValidateWithPath("\u0042\u006f\u0072d\u0065\u0072\u0062\u006f\u0074\u0074\u006f\u006d")
}
func (_ddeb ST_WrapSide) Validate() error { return _ddeb.ValidateWithPath("") }
func (_baa *ST_WrapType) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_baa = 0
	case "\u0074\u006f\u0070A\u006e\u0064\u0042\u006f\u0074\u0074\u006f\u006d":
		*_baa = 1
	case "\u0073\u0071\u0075\u0061\u0072\u0065":
		*_baa = 2
	case "\u006e\u006f\u006e\u0065":
		*_baa = 3
	case "\u0074\u0069\u0067h\u0074":
		*_baa = 4
	case "\u0074h\u0072\u006f\u0075\u0067\u0068":
		*_baa = 5
	}
	return nil
}
func (_bgd ST_BorderType) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_bgd.String(), start)
}
func (_afce *ST_VerticalAnchor) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_afce = 0
	case "\u006d\u0061\u0072\u0067\u0069\u006e":
		*_afce = 1
	case "\u0070\u0061\u0067\u0065":
		*_afce = 2
	case "\u0074\u0065\u0078\u0074":
		*_afce = 3
	case "\u006c\u0069\u006e\u0065":
		*_afce = 4
	}
	return nil
}

// ValidateWithPath validates the Bordertop and its children, prefixing error messages with path
func (_cca *Bordertop) ValidateWithPath(path string) error {
	if _ea := _cca.CT_Border.ValidateWithPath(path); _ea != nil {
		return _ea
	}
	return nil
}

type Borderright struct{ CT_Border }

func (_fgg ST_HorizontalAnchor) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_fda := _g.Attr{}
	_fda.Name = name
	switch _fgg {
	case ST_HorizontalAnchorUnset:
		_fda.Value = ""
	case ST_HorizontalAnchorMargin:
		_fda.Value = "\u006d\u0061\u0072\u0067\u0069\u006e"
	case ST_HorizontalAnchorPage:
		_fda.Value = "\u0070\u0061\u0067\u0065"
	case ST_HorizontalAnchorText:
		_fda.Value = "\u0074\u0065\u0078\u0074"
	case ST_HorizontalAnchorChar:
		_fda.Value = "\u0063\u0068\u0061\u0072"
	}
	return _fda, nil
}
func (_aag *CT_AnchorLock) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}
func (_gfe ST_WrapSide) ValidateWithPath(path string) error {
	switch _gfe {
	case 0, 1, 2, 3, 4:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gfe))
	}
	return nil
}
func (_abd ST_HorizontalAnchor) String() string {
	switch _abd {
	case 0:
		return ""
	case 1:
		return "\u006d\u0061\u0072\u0067\u0069\u006e"
	case 2:
		return "\u0070\u0061\u0067\u0065"
	case 3:
		return "\u0074\u0065\u0078\u0074"
	case 4:
		return "\u0063\u0068\u0061\u0072"
	}
	return ""
}

type Wrap struct{ CT_Wrap }

func (_fc *Borderleft) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_fc.CT_Border = *NewCT_Border()
	for _, _be := range start.Attr {
		if _be.Name.Local == "\u0074\u0079\u0070\u0065" {
			_fc.TypeAttr.UnmarshalXMLAttr(_be)
			continue
		}
		if _be.Name.Local == "\u0077\u0069\u0064t\u0068" {
			_agf, _db := _a.ParseUint(_be.Value, 10, 32)
			if _db != nil {
				return _db
			}
			_bfg := uint32(_agf)
			_fc.WidthAttr = &_bfg
			continue
		}
		if _be.Name.Local == "\u0073\u0068\u0061\u0064\u006f\u0077" {
			_fc.ShadowAttr.UnmarshalXMLAttr(_be)
			continue
		}
	}
	for {
		_gf, _geb := d.Token()
		if _geb != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0042\u006f\u0072\u0064\u0065\u0072\u006c\u0065\u0066\u0074\u003a\u0020%\u0073", _geb)
		}
		if _dge, _gfb := _gf.(_g.EndElement); _gfb && _dge.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Wrap and its children
func (_gga *Wrap) Validate() error { return _gga.ValidateWithPath("\u0057\u0072\u0061\u0070") }
func (_dg *Borderbottom) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0062\u006f\u0072d\u0065\u0072\u0062\u006f\u0074\u0074\u006f\u006d"
	return _dg.CT_Border.MarshalXML(e, start)
}

// ValidateWithPath validates the Borderright and its children, prefixing error messages with path
func (_gb *Borderright) ValidateWithPath(path string) error {
	if _cc := _gb.CT_Border.ValidateWithPath(path); _cc != nil {
		return _cc
	}
	return nil
}
func (_aeee *ST_HorizontalAnchor) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_aeee = 0
	case "\u006d\u0061\u0072\u0067\u0069\u006e":
		*_aeee = 1
	case "\u0070\u0061\u0067\u0065":
		*_aeee = 2
	case "\u0074\u0065\u0078\u0074":
		*_aeee = 3
	case "\u0063\u0068\u0061\u0072":
		*_aeee = 4
	}
	return nil
}
func (_gggb ST_BorderType) Validate() error { return _gggb.ValidateWithPath("") }

const (
	ST_BorderShadowUnset ST_BorderShadow = 0
	ST_BorderShadowT     ST_BorderShadow = 1
	ST_BorderShadowTrue  ST_BorderShadow = 2
	ST_BorderShadowF     ST_BorderShadow = 3
	ST_BorderShadowFalse ST_BorderShadow = 4
)

func init() {
	_f.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0043T\u005f\u0042\u006f\u0072\u0064\u0065r", NewCT_Border)
	_f.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0043T\u005f\u0057\u0072\u0061\u0070", NewCT_Wrap)
	_f.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0043\u0054\u005f\u0041\u006e\u0063\u0068\u006f\u0072\u004c\u006f\u0063\u006b", NewCT_AnchorLock)
	_f.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0062o\u0072\u0064\u0065\u0072\u0074\u006fp", NewBordertop)
	_f.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0062\u006f\u0072\u0064\u0065\u0072\u006c\u0065\u0066\u0074", NewBorderleft)
	_f.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "b\u006f\u0072\u0064\u0065\u0072\u0072\u0069\u0067\u0068\u0074", NewBorderright)
	_f.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0062\u006f\u0072d\u0065\u0072\u0062\u006f\u0074\u0074\u006f\u006d", NewBorderbottom)
	_f.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0077\u0072\u0061\u0070", NewWrap)
	_f.RegisterConstructor("\u0075\u0072n\u003a\u0073\u0063\u0068e\u006d\u0061s\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006ff\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065:\u0077\u006f\u0072\u0064", "\u0061\u006e\u0063\u0068\u006f\u0072\u006c\u006f\u0063\u006b", NewAnchorlock)
}
