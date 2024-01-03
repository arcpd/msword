package relationships

import (
	_g "encoding/xml"
	_a "fmt"

	_fg "github.com/arcpd/msword"
	_f "github.com/arcpd/msword/common/logger"
)

func (_fa *CT_Relationships) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fa.Relationship != nil {
		_edc := _g.StartElement{Name: _g.Name{Local: "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}}
		for _, _ee := range _fa.Relationship {
			e.EncodeElement(_ee, _edc)
		}
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_dc *ST_TargetMode) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_dc = 0
	case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":
		*_dc = 1
	case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":
		*_dc = 2
	}
	return nil
}

func NewRelationship() *Relationship {
	_cd := &Relationship{}
	_cd.CT_Relationship = *NewCT_Relationship()
	return _cd
}

func (_fff *CT_Relationship) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _bb := range start.Attr {
		if _bb.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065" {
			_fff.TargetModeAttr.UnmarshalXMLAttr(_bb)
			continue
		}
		if _bb.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074" {
			_ae, _fc := _bb.Value, error(nil)
			if _fc != nil {
				return _fc
			}
			_fff.TargetAttr = _ae
			continue
		}
		if _bb.Name.Local == "\u0054\u0079\u0070\u0065" {
			_da, _c := _bb.Value, error(nil)
			if _c != nil {
				return _c
			}
			_fff.TypeAttr = _da
			continue
		}
		if _bb.Name.Local == "\u0049\u0064" {
			_dg, _eg := _bb.Value, error(nil)
			if _eg != nil {
				return _eg
			}
			_fff.IdAttr = _dg
			continue
		}
	}
	for {
		_db, _bd := d.Token()
		if _bd != nil {
			return _a.Errorf("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069o\u006e\u0073\u0068i\u0070:\u0020\u0025\u0073", _bd)
		}
		if _fgc, _dba := _db.(_g.CharData); _dba {
			_fff.Content = string(_fgc)
		}
		if _ed, _ad := _db.(_g.EndElement); _ad && _ed.Name == start.Name {
			break
		}
	}
	return nil
}

func (_fgg *Relationship) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_fgg.CT_Relationship = *NewCT_Relationship()
	for _, _dde := range start.Attr {
		if _dde.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065" {
			_fgg.TargetModeAttr.UnmarshalXMLAttr(_dde)
			continue
		}
		if _dde.Name.Local == "\u0054\u0061\u0072\u0067\u0065\u0074" {
			_dfg, _fdf := _dde.Value, error(nil)
			if _fdf != nil {
				return _fdf
			}
			_fgg.TargetAttr = _dfg
			continue
		}
		if _dde.Name.Local == "\u0054\u0079\u0070\u0065" {
			_fce, _ea := _dde.Value, error(nil)
			if _ea != nil {
				return _ea
			}
			_fgg.TypeAttr = _fce
			continue
		}
		if _dde.Name.Local == "\u0049\u0064" {
			_egc, _bc := _dde.Value, error(nil)
			if _bc != nil {
				return _bc
			}
			_fgg.IdAttr = _egc
			continue
		}
	}
	for {
		_cdd, _cg := d.Token()
		if _cg != nil {
			return _a.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0052\u0065\u006c\u0061\u0074i\u006f\u006e\u0073\u0068\u0069\u0070\u003a\u0020\u0025\u0073", _cg)
		}
		if _ga, _ffe := _cdd.(_g.EndElement); _ffe && _ga.Name == start.Name {
			break
		}
	}
	return nil
}

func (_e *CT_Relationship) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _e.TargetModeAttr != ST_TargetModeUnset {
		_ff, _fd := _e.TargetModeAttr.MarshalXMLAttr(_g.Name{Local: "\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"})
		if _fd != nil {
			return _fd
		}
		start.Attr = append(start.Attr, _ff)
	}
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0054\u0061\u0072\u0067\u0065\u0074"}, Value: _a.Sprintf("\u0025\u0076", _e.TargetAttr)})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0054\u0079\u0070\u0065"}, Value: _a.Sprintf("\u0025\u0076", _e.TypeAttr)})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0049\u0064"}, Value: _a.Sprintf("\u0025\u0076", _e.IdAttr)})
	e.EncodeElement(_e.Content, start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the Relationships and its children
func (_bce *Relationships) Validate() error {
	return _bce.ValidateWithPath("\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073")
}

// Validate validates the CT_Relationships and its children
func (_dd *CT_Relationships) Validate() error {
	return _dd.ValidateWithPath("\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073")
}

type CT_Relationship struct {
	TargetModeAttr ST_TargetMode
	TargetAttr     string
	TypeAttr       string
	IdAttr         string
	Content        string
}

const (
	ST_TargetModeUnset    ST_TargetMode = 0
	ST_TargetModeExternal ST_TargetMode = 1
	ST_TargetModeInternal ST_TargetMode = 2
)

type ST_TargetMode byte

func (_bca ST_TargetMode) String() string {
	switch _bca {
	case 0:
		return ""
	case 1:
		return "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c"
	case 2:
		return "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c"
	}
	return ""
}

// ValidateWithPath validates the Relationships and its children, prefixing error messages with path
func (_gbf *Relationships) ValidateWithPath(path string) error {
	if _afb := _gbf.CT_Relationships.ValidateWithPath(path); _afb != nil {
		return _afb
	}
	return nil
}

func (_gf *Relationships) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"
	return _gf.CT_Relationships.MarshalXML(e, start)
}

// Validate validates the CT_Relationship and its children
func (_gc *CT_Relationship) Validate() error {
	return _gc.ValidateWithPath("\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070")
}

func (_aec ST_TargetMode) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_cc := _g.Attr{}
	_cc.Name = name
	switch _aec {
	case ST_TargetModeUnset:
		_cc.Value = ""
	case ST_TargetModeExternal:
		_cc.Value = "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c"
	case ST_TargetModeInternal:
		_cc.Value = "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c"
	}
	return _cc, nil
}

func (_afg *Relationship) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return _afg.CT_Relationship.MarshalXML(e, start)
}

func (_fda ST_TargetMode) ValidateWithPath(path string) error {
	switch _fda {
	case 0, 1, 2:
	default:
		return _a.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fda))
	}
	return nil
}

func NewRelationships() *Relationships {
	_bcg := &Relationships{}
	_bcg.CT_Relationships = *NewCT_Relationships()
	return _bcg
}

// Validate validates the Relationship and its children
func (_dgb *Relationship) Validate() error {
	return _dgb.ValidateWithPath("\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070")
}

type CT_Relationships struct{ Relationship []*Relationship }

func (_bae *Relationships) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bae.CT_Relationships = *NewCT_Relationships()
_cbd:
	for {
		_beg, _eab := d.Token()
		if _eab != nil {
			return _eab
		}
		switch _aab := _beg.(type) {
		case _g.StartElement:
			switch _aab.Name {
			case _g.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", Local: "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:
				_bdc := NewRelationship()
				if _adbg := d.DecodeElement(_bdc, &_aab); _adbg != nil {
					return _adbg
				}
				_bae.Relationship = append(_bae.Relationship, _bdc)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0052\u0065\u006c\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073 \u0025\u0076", _aab.Name)
				if _gba := d.Skip(); _gba != nil {
					return _gba
				}
			}
		case _g.EndElement:
			break _cbd
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the Relationship and its children, prefixing error messages with path
func (_egg *Relationship) ValidateWithPath(path string) error {
	if _cbb := _egg.CT_Relationship.ValidateWithPath(path); _cbb != nil {
		return _cbb
	}
	return nil
}

// ValidateWithPath validates the CT_Relationship and its children, prefixing error messages with path
func (_ega *CT_Relationship) ValidateWithPath(path string) error {
	if _cb := _ega.TargetModeAttr.ValidateWithPath(path + "\u002fT\u0061r\u0067\u0065\u0074\u004d\u006f\u0064\u0065\u0041\u0074\u0074\u0072"); _cb != nil {
		return _cb
	}
	return nil
}

func (_ag ST_TargetMode) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_ag.String(), start)
}

func NewCT_Relationship() *CT_Relationship { _b := &CT_Relationship{}; return _b }

func (_baa *ST_TargetMode) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_aeb, _cbe := d.Token()
	if _cbe != nil {
		return _cbe
	}
	if _dab, _cgb := _aeb.(_g.EndElement); _cgb && _dab.Name == start.Name {
		*_baa = 1
		return nil
	}
	if _gcb, _egcf := _aeb.(_g.CharData); !_egcf {
		return _a.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _aeb)
	} else {
		switch string(_gcb) {
		case "":
			*_baa = 0
		case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":
			*_baa = 1
		case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":
			*_baa = 2
		}
	}
	_aeb, _cbe = d.Token()
	if _cbe != nil {
		return _cbe
	}
	if _dgbc, _bbc := _aeb.(_g.EndElement); _bbc && _dgbc.Name == start.Name {
		return nil
	}
	return _a.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _aeb)
}

// ValidateWithPath validates the CT_Relationships and its children, prefixing error messages with path
func (_ba *CT_Relationships) ValidateWithPath(path string) error {
	for _ge, _df := range _ba.Relationship {
		if _ef := _df.ValidateWithPath(_a.Sprintf("\u0025\u0073\u002f\u0052el\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u005b\u0025\u0064\u005d", path, _ge)); _ef != nil {
			return _ef
		}
	}
	return nil
}

func NewCT_Relationships() *CT_Relationships { _bbf := &CT_Relationships{}; return _bbf }

func (_fed ST_TargetMode) Validate() error { return _fed.ValidateWithPath("") }

type Relationship struct{ CT_Relationship }

func (_fe *CT_Relationships) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_eb:
	for {
		_aa, _adb := d.Token()
		if _adb != nil {
			return _adb
		}
		switch _fgca := _aa.(type) {
		case _g.StartElement:
			switch _fgca.Name {
			case _g.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", Local: "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:
				_gd := NewRelationship()
				if _de := d.DecodeElement(_gd, &_fgca); _de != nil {
					return _de
				}
				_fe.Relationship = append(_fe.Relationship, _gd)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u0020\u0025v", _fgca.Name)
				if _af := d.Skip(); _af != nil {
					return _af
				}
			}
		case _g.EndElement:
			break _eb
		case _g.CharData:
		}
	}
	return nil
}

type Relationships struct{ CT_Relationships }

func init() {
	_fg.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073", NewCT_Relationships)
	_fg.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070", NewCT_Relationship)
	_fg.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073", NewRelationships)
	_fg.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s", "\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070", NewRelationship)
}
