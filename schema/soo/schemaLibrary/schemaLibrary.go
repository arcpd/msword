package schemaLibrary

import (
	_ce "encoding/xml"
	_g "fmt"

	_e "github.com/arcpd/msword"
	_a "github.com/arcpd/msword/common/logger"
)

type CT_SchemaLibrary struct{ Schema []*CT_Schema }

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_fdg *CT_Schema) ValidateWithPath(path string) error { return nil }

func (_aaa *SchemaLibrary) UnmarshalXML(d *_ce.Decoder, start _ce.StartElement) error {
	_aaa.CT_SchemaLibrary = *NewCT_SchemaLibrary()
_efe:
	for {
		_ag, _gee := d.Token()
		if _gee != nil {
			return _gee
		}
		switch _bb := _ag.(type) {
		case _ce.StartElement:
			switch _bb.Name {
			case _ce.Name{Space: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0063\u0068\u0065\u006d\u0061"}:
				_fg := NewCT_Schema()
				if _bfa := d.DecodeElement(_fg, &_bb); _bfa != nil {
					return _bfa
				}
				_aaa.Schema = append(_aaa.Schema, _fg)
			default:
				_a.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0063\u0068\u0065m\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079 \u0025\u0076", _bb.Name)
				if _bbg := d.Skip(); _bbg != nil {
					return _bbg
				}
			}
		case _ce.EndElement:
			break _efe
		case _ce.CharData:
		}
	}
	return nil
}

func (_f *CT_Schema) UnmarshalXML(d *_ce.Decoder, start _ce.StartElement) error {
	for _, _aa := range start.Attr {
		if _aa.Name.Local == "\u0075\u0072\u0069" {
			_cf, _cb := _aa.Value, error(nil)
			if _cb != nil {
				return _cb
			}
			_f.UriAttr = &_cf
			continue
		}
		if _aa.Name.Local == "\u006d\u0061n\u0069\u0066\u0065s\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e" {
			_gb, _de := _aa.Value, error(nil)
			if _de != nil {
				return _de
			}
			_f.ManifestLocationAttr = &_gb
			continue
		}
		if _aa.Name.Local == "\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e" {
			_fd, _ec := _aa.Value, error(nil)
			if _ec != nil {
				return _ec
			}
			_f.SchemaLocationAttr = &_fd
			continue
		}
		if _aa.Name.Local == "\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065" {
			_dg, _b := _aa.Value, error(nil)
			if _b != nil {
				return _b
			}
			_f.SchemaLanguageAttr = &_dg
			continue
		}
	}
	for {
		_cbg, _gf := d.Token()
		if _gf != nil {
			return _g.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0068e\u006d\u0061\u003a\u0020\u0025\u0073", _gf)
		}
		if _da, _ee := _cbg.(_ce.EndElement); _ee && _da.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ed *CT_Schema) MarshalXML(e *_ce.Encoder, start _ce.StartElement) error {
	if _ed.UriAttr != nil {
		start.Attr = append(start.Attr, _ce.Attr{Name: _ce.Name{Local: "\u006d\u0061\u003a\u0075\u0072\u0069"}, Value: _g.Sprintf("\u0025\u0076", *_ed.UriAttr)})
	}
	if _ed.ManifestLocationAttr != nil {
		start.Attr = append(start.Attr, _ce.Attr{Name: _ce.Name{Local: "\u006d\u0061\u003a\u006dan\u0069\u0066\u0065\u0073\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"}, Value: _g.Sprintf("\u0025\u0076", *_ed.ManifestLocationAttr)})
	}
	if _ed.SchemaLocationAttr != nil {
		start.Attr = append(start.Attr, _ce.Attr{Name: _ce.Name{Local: "\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"}, Value: _g.Sprintf("\u0025\u0076", *_ed.SchemaLocationAttr)})
	}
	if _ed.SchemaLanguageAttr != nil {
		start.Attr = append(start.Attr, _ce.Attr{Name: _ce.Name{Local: "\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"}, Value: _g.Sprintf("\u0025\u0076", *_ed.SchemaLanguageAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_ce.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_gg *CT_SchemaLibrary) ValidateWithPath(path string) error {
	for _ge, _fe := range _gg.Schema {
		if _feb := _fe.ValidateWithPath(_g.Sprintf("\u0025\u0073\u002f\u0053\u0063\u0068\u0065\u006d\u0061\u005b\u0025\u0064\u005d", path, _ge)); _feb != nil {
			return _feb
		}
	}
	return nil
}

func NewSchemaLibrary() *SchemaLibrary {
	_feg := &SchemaLibrary{}
	_feg.CT_SchemaLibrary = *NewCT_SchemaLibrary()
	return _feg
}

func (_ef *CT_SchemaLibrary) MarshalXML(e *_ce.Encoder, start _ce.StartElement) error {
	e.EncodeToken(start)
	if _ef.Schema != nil {
		_ac := _ce.StartElement{Name: _ce.Name{Local: "\u006da\u003a\u0073\u0063\u0068\u0065\u006da"}}
		for _, _bde := range _ef.Schema {
			e.EncodeElement(_bde, _ac)
		}
	}
	e.EncodeToken(_ce.EndElement{Name: start.Name})
	return nil
}

type SchemaLibrary struct{ CT_SchemaLibrary }

func NewCT_SchemaLibrary() *CT_SchemaLibrary { _ea := &CT_SchemaLibrary{}; return _ea }

func NewCT_Schema() *CT_Schema { _d := &CT_Schema{}; return _d }

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_bg *SchemaLibrary) ValidateWithPath(path string) error {
	if _fdba := _bg.CT_SchemaLibrary.ValidateWithPath(path); _fdba != nil {
		return _fdba
	}
	return nil
}

func (_fdbf *SchemaLibrary) MarshalXML(e *_ce.Encoder, start _ce.StartElement) error {
	start.Attr = append(start.Attr, _ce.Attr{Name: _ce.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _ce.Attr{Name: _ce.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u006d\u0061"}, Value: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _ce.Attr{Name: _ce.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006d\u0061:\u0073\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079"
	return _fdbf.CT_SchemaLibrary.MarshalXML(e, start)
}

// Validate validates the SchemaLibrary and its children
func (_fb *SchemaLibrary) Validate() error {
	return _fb.ValidateWithPath("\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079")
}

// Validate validates the CT_Schema and its children
func (_dd *CT_Schema) Validate() error {
	return _dd.ValidateWithPath("\u0043T\u005f\u0053\u0063\u0068\u0065\u006da")
}

// Validate validates the CT_SchemaLibrary and its children
func (_fa *CT_SchemaLibrary) Validate() error {
	return _fa.ValidateWithPath("\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079")
}

func (_dc *CT_SchemaLibrary) UnmarshalXML(d *_ce.Decoder, start _ce.StartElement) error {
_ba:
	for {
		_eb, _fdb := d.Token()
		if _fdb != nil {
			return _fdb
		}
		switch _dcb := _eb.(type) {
		case _ce.StartElement:
			switch _dcb.Name {
			case _ce.Name{Space: "\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0063\u0068\u0065\u006d\u0061"}:
				_bc := NewCT_Schema()
				if _bf := d.DecodeElement(_bc, &_dcb); _bf != nil {
					return _bf
				}
				_dc.Schema = append(_dc.Schema, _bc)
			default:
				_a.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u0020\u0025v", _dcb.Name)
				if _dgd := d.Skip(); _dgd != nil {
					return _dgd
				}
			}
		case _ce.EndElement:
			break _ba
		case _ce.CharData:
		}
	}
	return nil
}

type CT_Schema struct {
	UriAttr              *string
	ManifestLocationAttr *string
	SchemaLocationAttr   *string
	SchemaLanguageAttr   *string
}

func init() {
	_e.RegisterConstructor("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", "\u0043T\u005f\u0053\u0063\u0068\u0065\u006da", NewCT_Schema)
	_e.RegisterConstructor("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", "\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079", NewCT_SchemaLibrary)
	_e.RegisterConstructor("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", "\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079", NewSchemaLibrary)
}
