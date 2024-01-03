package powerpoint

import (
	_d "encoding/xml"
	_c "fmt"
	_e "github.com/arcpd/msword"
)

func NewCT_Rel() *CT_Rel { _gg := &CT_Rel{}; return _gg }

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_fec *CT_Rel) ValidateWithPath(path string) error { return nil }
func NewIscomment() *Iscomment                          { _gcd := &Iscomment{}; _gcd.CT_Empty = *NewCT_Empty(); return _gcd }
func NewCT_Empty() *CT_Empty                            { _b := &CT_Empty{}; return _b }

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_fae *Iscomment) ValidateWithPath(path string) error {
	if _eg := _fae.CT_Empty.ValidateWithPath(path); _eg != nil {
		return _eg
	}
	return nil
}
func (_dg *CT_Rel) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for _, _ca := range start.Attr {
		if _ca.Name.Local == "\u0069\u0064" {
			_ae, _fb := _ca.Value, error(nil)
			if _fb != nil {
				return _fb
			}
			_dg.IdAttr = &_ae
			continue
		}
	}
	for {
		_gc, _ba := d.Token()
		if _ba != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073", _ba)
		}
		if _eb, _ed := _gc.(_d.EndElement); _ed && _eb.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Iscomment and its children
func (_edg *Iscomment) Validate() error {
	return _edg.ValidateWithPath("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et")
}

type CT_Rel struct{ IdAttr *string }
type Iscomment struct{ CT_Empty }

func (_bg *CT_Empty) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	for {
		_cb, _f := d.Token()
		if _f != nil {
			return _c.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073", _f)
		}
		if _a, _fe := _cb.(_d.EndElement); _fe && _a.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_cg *Textdata) ValidateWithPath(path string) error {
	if _gcdf := _cg.CT_Rel.ValidateWithPath(path); _gcdf != nil {
		return _gcdf
	}
	return nil
}

// Validate validates the CT_Empty and its children
func (_cc *CT_Empty) Validate() error {
	return _cc.ValidateWithPath("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079")
}
func (_cd *CT_Empty) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}

type CT_Empty struct{}

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_ge *CT_Empty) ValidateWithPath(path string) error { return nil }

type Textdata struct{ CT_Rel }

func NewTextdata() *Textdata { _gf := &Textdata{}; _gf.CT_Rel = *NewCT_Rel(); return _gf }

// Validate validates the CT_Rel and its children
func (_ce *CT_Rel) Validate() error {
	return _ce.ValidateWithPath("\u0043\u0054\u005f\u0052\u0065\u006c")
}
func (_gd *Textdata) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_gd.CT_Rel = *NewCT_Rel()
	for _, _dc := range start.Attr {
		if _dc.Name.Local == "\u0069\u0064" {
			_ged, _ad := _dc.Value, error(nil)
			if _ad != nil {
				return _ad
			}
			_gd.IdAttr = &_ged
			continue
		}
	}
	for {
		_ee, _ec := d.Token()
		if _ec != nil {
			return _c.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073", _ec)
		}
		if _egb, _cf := _ee.(_d.EndElement); _cf && _egb.Name == start.Name {
			break
		}
	}
	return nil
}
func (_df *Textdata) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061"
	return _df.CT_Rel.MarshalXML(e, start)
}
func (_fg *CT_Rel) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	if _fg.IdAttr != nil {
		start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0069\u0064"}, Value: _c.Sprintf("\u0025\u0076", *_fg.IdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_d.EndElement{Name: start.Name})
	return nil
}
func (_bgb *Iscomment) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0069s\u0063\u006f\u006d\u006d\u0065\u006et"
	return _bgb.CT_Empty.MarshalXML(e, start)
}
func (_cdb *Iscomment) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_cdb.CT_Empty = *NewCT_Empty()
	for {
		_fa, _ga := d.Token()
		if _ga != nil {
			return _c.Errorf("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073", _ga)
		}
		if _bac, _ef := _fa.(_d.EndElement); _ef && _bac.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Textdata and its children
func (_cfg *Textdata) Validate() error {
	return _cfg.ValidateWithPath("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061")
}
func init() {
	_e.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074", "\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079", NewCT_Empty)
	_e.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074", "\u0043\u0054\u005f\u0052\u0065\u006c", NewCT_Rel)
	_e.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074", "\u0069s\u0063\u006f\u006d\u006d\u0065\u006et", NewIscomment)
	_e.RegisterConstructor("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074", "\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061", NewTextdata)
}
