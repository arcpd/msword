package core_properties

import (
	_b "encoding/xml"
	_db "fmt"
	_bc "time"

	_c "github.com/arcpd/msword"
	_e "github.com/arcpd/msword/common/logger"
)

func (_cd *CT_CoreProperties) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
_dge:
	for {
		_bb, _gc := d.Token()
		if _gc != nil {
			return _gc
		}
		switch _df := _bb.(type) {
		case _b.StartElement:
			switch _df.Name {
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:
				_cd.Category = new(string)
				if _fg := d.DecodeElement(_cd.Category, &_df); _fg != nil {
					return _fg
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:
				_cd.ContentStatus = new(string)
				if _ed := d.DecodeElement(_cd.ContentStatus, &_df); _ed != nil {
					return _ed
				}
			case _b.Name{Space: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", Local: "\u0063r\u0065\u0061\u0074\u0065\u0064"}:
				_cd.Created = new(_c.XSDAny)
				if _bcd := d.DecodeElement(_cd.Created, &_df); _bcd != nil {
					return _bcd
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0063r\u0065\u0061\u0074\u006f\u0072"}:
				_cd.Creator = new(_c.XSDAny)
				if _fd := d.DecodeElement(_cd.Creator, &_df); _fd != nil {
					return _fd
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:
				_cd.Description = new(_c.XSDAny)
				if _ceb := d.DecodeElement(_cd.Description, &_df); _ceb != nil {
					return _ceb
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:
				_cd.Identifier = new(_c.XSDAny)
				if _aa := d.DecodeElement(_cd.Identifier, &_df); _aa != nil {
					return _aa
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:
				_cd.Keywords = NewCT_Keywords()
				if _ac := d.DecodeElement(_cd.Keywords, &_df); _ac != nil {
					return _ac
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:
				_cd.Language = new(_c.XSDAny)
				if _ae := d.DecodeElement(_cd.Language, &_df); _ae != nil {
					return _ae
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:
				_cd.LastModifiedBy = new(string)
				if _eb := d.DecodeElement(_cd.LastModifiedBy, &_df); _eb != nil {
					return _eb
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:
				_cd.LastPrinted = new(_bc.Time)
				if _gaf := d.DecodeElement(_cd.LastPrinted, &_df); _gaf != nil {
					return _gaf
				}
			case _b.Name{Space: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", Local: "\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:
				_cd.Modified = new(_c.XSDAny)
				if _ee := d.DecodeElement(_cd.Modified, &_df); _ee != nil {
					return _ee
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:
				_cd.Revision = new(string)
				if _af := d.DecodeElement(_cd.Revision, &_df); _af != nil {
					return _af
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0073u\u0062\u006a\u0065\u0063\u0074"}:
				_cd.Subject = new(_c.XSDAny)
				if _dd := d.DecodeElement(_cd.Subject, &_df); _dd != nil {
					return _dd
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0074\u0069\u0074l\u0065"}:
				_cd.Title = new(_c.XSDAny)
				if _aaa := d.DecodeElement(_cd.Title, &_df); _aaa != nil {
					return _aaa
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0076e\u0072\u0073\u0069\u006f\u006e"}:
				_cd.Version = new(string)
				if _cac := d.DecodeElement(_cd.Version, &_df); _cac != nil {
					return _cac
				}
			default:
				_e.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u0020\u0025\u0076", _df.Name)
				if _eeb := d.Skip(); _eeb != nil {
					return _eeb
				}
			}
		case _b.EndElement:
			break _dge
		case _b.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_CoreProperties and its children, prefixing error messages with path
func (_cbc *CT_CoreProperties) ValidateWithPath(path string) error {
	if _cbc.Keywords != nil {
		if _ec := _cbc.Keywords.ValidateWithPath(path + "\u002fK\u0065\u0079\u0077\u006f\u0072\u0064s"); _ec != nil {
			return _ec
		}
	}
	return nil
}

func (_edc *CT_Keyword) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for _, _afa := range start.Attr {
		if _afa.Name.Space == "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065" && _afa.Name.Local == "\u006c\u0061\u006e\u0067" {
			_bg, _cdd := _afa.Value, error(nil)
			if _cdd != nil {
				return _cdd
			}
			_edc.LangAttr = &_bg
			continue
		}
	}
	for {
		_gae, _bdf := d.Token()
		if _bdf != nil {
			return _db.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u003a\u0020%\u0073", _bdf)
		}
		if _fa, _dc := _gae.(_b.CharData); _dc {
			_edc.Content = string(_fa)
		}
		if _dfg, _fdb := _gae.(_b.EndElement); _fdb && _dfg.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ace *CT_Keywords) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	if _ace.LangAttr != nil {
		start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"}, Value: _db.Sprintf("\u0025\u0076", *_ace.LangAttr)})
	}
	e.EncodeToken(start)
	if _ace.Value != nil {
		_bbf := _b.StartElement{Name: _b.Name{Local: "\u0063\u0070\u003a\u0076\u0061\u006c\u0075\u0065"}}
		for _, _fb := range _ace.Value {
			e.EncodeElement(_fb, _bbf)
		}
	}
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_Keyword and its children
func (_bba *CT_Keyword) Validate() error {
	return _bba.ValidateWithPath("\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064")
}

func NewCT_Keyword() *CT_Keyword { _dad := &CT_Keyword{}; return _dad }

type CT_Keywords struct {
	LangAttr *string
	Value    []*CT_Keyword
}

// ValidateWithPath validates the CoreProperties and its children, prefixing error messages with path
func (_cba *CoreProperties) ValidateWithPath(path string) error {
	if _afb := _cba.CT_CoreProperties.ValidateWithPath(path); _afb != nil {
		return _afb
	}
	return nil
}

// Validate validates the CoreProperties and its children
func (_cc *CoreProperties) Validate() error {
	return _cc.ValidateWithPath("\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073")
}

// Validate validates the CT_Keywords and its children
func (_dfb *CT_Keywords) Validate() error {
	return _dfb.ValidateWithPath("C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073")
}

func (_gg *CT_Keywords) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	for _, _ece := range start.Attr {
		if _ece.Name.Space == "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065" && _ece.Name.Local == "\u006c\u0061\u006e\u0067" {
			_afe, _gd := _ece.Value, error(nil)
			if _gd != nil {
				return _gd
			}
			_gg.LangAttr = &_afe
			continue
		}
	}
_cag:
	for {
		_aac, _bgf := d.Token()
		if _bgf != nil {
			return _bgf
		}
		switch _dbc := _aac.(type) {
		case _b.StartElement:
			switch _dbc.Name {
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0076\u0061\u006cu\u0065"}:
				_gafd := NewCT_Keyword()
				if _ge := d.DecodeElement(_gafd, &_dbc); _ge != nil {
					return _ge
				}
				_gg.Value = append(_gg.Value, _gafd)
			default:
				_e.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073\u0020\u0025\u0076", _dbc.Name)
				if _ef := d.Skip(); _ef != nil {
					return _ef
				}
			}
		case _b.EndElement:
			break _cag
		case _b.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Keywords and its children, prefixing error messages with path
func (_aag *CT_Keywords) ValidateWithPath(path string) error {
	for _cbb, _gf := range _aag.Value {
		if _bgc := _gf.ValidateWithPath(_db.Sprintf("\u0025\u0073\u002fV\u0061\u006c\u0075\u0065\u005b\u0025\u0064\u005d", path, _cbb)); _bgc != nil {
			return _bgc
		}
	}
	return nil
}

type CT_Keyword struct {
	LangAttr *string
	Content  string
}

// ValidateWithPath validates the CT_Keyword and its children, prefixing error messages with path
func (_aca *CT_Keyword) ValidateWithPath(path string) error { return nil }

type CoreProperties struct{ CT_CoreProperties }

func NewCT_Keywords() *CT_Keywords { _daf := &CT_Keywords{}; return _daf }

func NewCoreProperties() *CoreProperties {
	_bga := &CoreProperties{}
	_bga.CT_CoreProperties = *NewCT_CoreProperties()
	return _bga
}

// Validate validates the CT_CoreProperties and its children
func (_cg *CT_CoreProperties) Validate() error {
	return _cg.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073")
}

func (_bca *CT_Keyword) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	if _bca.LangAttr != nil {
		start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"}, Value: _db.Sprintf("\u0025\u0076", *_bca.LangAttr)})
	}
	e.EncodeElement(_bca.Content, start)
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

func (_dg *CT_CoreProperties) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	e.EncodeToken(start)
	if _dg.Category != nil {
		_a := _b.StartElement{Name: _b.Name{Local: "c\u0070\u003a\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}}
		_c.AddPreserveSpaceAttr(&_a, *_dg.Category)
		e.EncodeElement(_dg.Category, _a)
	}
	if _dg.ContentStatus != nil {
		_f := _b.StartElement{Name: _b.Name{Local: "\u0063\u0070:\u0063\u006f\u006et\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}}
		_c.AddPreserveSpaceAttr(&_f, *_dg.ContentStatus)
		e.EncodeElement(_dg.ContentStatus, _f)
	}
	if _dg.Created != nil {
		_ba := _b.StartElement{Name: _b.Name{Local: "\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064"}}
		e.EncodeElement(_dg.Created, _ba)
	}
	if _dg.Creator != nil {
		_ag := _b.StartElement{Name: _b.Name{Local: "\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}}
		e.EncodeElement(_dg.Creator, _ag)
	}
	if _dg.Description != nil {
		_g := _b.StartElement{Name: _b.Name{Local: "\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}}
		e.EncodeElement(_dg.Description, _g)
	}
	if _dg.Identifier != nil {
		_ce := _b.StartElement{Name: _b.Name{Local: "\u0064\u0063\u003a\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}}
		e.EncodeElement(_dg.Identifier, _ce)
	}
	if _dg.Keywords != nil {
		_ga := _b.StartElement{Name: _b.Name{Local: "c\u0070\u003a\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}}
		e.EncodeElement(_dg.Keywords, _ga)
	}
	if _dg.Language != nil {
		_cb := _b.StartElement{Name: _b.Name{Local: "d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}}
		e.EncodeElement(_dg.Language, _cb)
	}
	if _dg.LastModifiedBy != nil {
		_bf := _b.StartElement{Name: _b.Name{Local: "\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}}
		_c.AddPreserveSpaceAttr(&_bf, *_dg.LastModifiedBy)
		e.EncodeElement(_dg.LastModifiedBy, _bf)
	}
	if _dg.LastPrinted != nil {
		_be := _b.StartElement{Name: _b.Name{Local: "\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u0050\u0072i\u006e\u0074\u0065\u0064"}}
		e.EncodeElement(_dg.LastPrinted, _be)
	}
	if _dg.Modified != nil {
		_fe := _b.StartElement{Name: _b.Name{Local: "\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}}
		e.EncodeElement(_dg.Modified, _fe)
	}
	if _dg.Revision != nil {
		_ca := _b.StartElement{Name: _b.Name{Local: "c\u0070\u003a\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}}
		_c.AddPreserveSpaceAttr(&_ca, *_dg.Revision)
		e.EncodeElement(_dg.Revision, _ca)
	}
	if _dg.Subject != nil {
		_cf := _b.StartElement{Name: _b.Name{Local: "\u0064\u0063\u003a\u0073\u0075\u0062\u006a\u0065\u0063\u0074"}}
		e.EncodeElement(_dg.Subject, _cf)
	}
	if _dg.Title != nil {
		_fc := _b.StartElement{Name: _b.Name{Local: "\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}}
		e.EncodeElement(_dg.Title, _fc)
	}
	if _dg.Version != nil {
		_eg := _b.StartElement{Name: _b.Name{Local: "\u0063\u0070\u003a\u0076\u0065\u0072\u0073\u0069\u006f\u006e"}}
		_c.AddPreserveSpaceAttr(&_eg, *_dg.Version)
		e.EncodeElement(_dg.Version, _eg)
	}
	e.EncodeToken(_b.EndElement{Name: start.Name})
	return nil
}

func NewCT_CoreProperties() *CT_CoreProperties { _da := &CT_CoreProperties{}; return _da }

type CT_CoreProperties struct {
	Category       *string
	ContentStatus  *string
	Created        *_c.XSDAny
	Creator        *_c.XSDAny
	Description    *_c.XSDAny
	Identifier     *_c.XSDAny
	Keywords       *CT_Keywords
	Language       *_c.XSDAny
	LastModifiedBy *string
	LastPrinted    *_bc.Time
	Modified       *_c.XSDAny
	Revision       *string
	Subject        *_c.XSDAny
	Title          *_c.XSDAny
	Version        *string
}

func (_bag *CoreProperties) UnmarshalXML(d *_b.Decoder, start _b.StartElement) error {
	_bag.CT_CoreProperties = *NewCT_CoreProperties()
_dca:
	for {
		_eba, _ad := d.Token()
		if _ad != nil {
			return _ad
		}
		switch _aae := _eba.(type) {
		case _b.StartElement:
			switch _aae.Name {
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:
				_bag.Category = new(string)
				if _dbcb := d.DecodeElement(_bag.Category, &_aae); _dbcb != nil {
					return _dbcb
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:
				_bag.ContentStatus = new(string)
				if _ceg := d.DecodeElement(_bag.ContentStatus, &_aae); _ceg != nil {
					return _ceg
				}
			case _b.Name{Space: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", Local: "\u0063r\u0065\u0061\u0074\u0065\u0064"}:
				_bag.Created = new(_c.XSDAny)
				if _ff := d.DecodeElement(_bag.Created, &_aae); _ff != nil {
					return _ff
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0063r\u0065\u0061\u0074\u006f\u0072"}:
				_bag.Creator = new(_c.XSDAny)
				if _aad := d.DecodeElement(_bag.Creator, &_aae); _aad != nil {
					return _aad
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:
				_bag.Description = new(_c.XSDAny)
				if _daa := d.DecodeElement(_bag.Description, &_aae); _daa != nil {
					return _daa
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:
				_bag.Identifier = new(_c.XSDAny)
				if _gdd := d.DecodeElement(_bag.Identifier, &_aae); _gdd != nil {
					return _gdd
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:
				_bag.Keywords = NewCT_Keywords()
				if _gaa := d.DecodeElement(_bag.Keywords, &_aae); _gaa != nil {
					return _gaa
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:
				_bag.Language = new(_c.XSDAny)
				if _eca := d.DecodeElement(_bag.Language, &_aae); _eca != nil {
					return _eca
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:
				_bag.LastModifiedBy = new(string)
				if _eea := d.DecodeElement(_bag.LastModifiedBy, &_aae); _eea != nil {
					return _eea
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:
				_bag.LastPrinted = new(_bc.Time)
				if _eec := d.DecodeElement(_bag.LastPrinted, &_aae); _eec != nil {
					return _eec
				}
			case _b.Name{Space: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", Local: "\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:
				_bag.Modified = new(_c.XSDAny)
				if _cff := d.DecodeElement(_bag.Modified, &_aae); _cff != nil {
					return _cff
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:
				_bag.Revision = new(string)
				if _eda := d.DecodeElement(_bag.Revision, &_aae); _eda != nil {
					return _eda
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0073u\u0062\u006a\u0065\u0063\u0074"}:
				_bag.Subject = new(_c.XSDAny)
				if _ea := d.DecodeElement(_bag.Subject, &_aae); _ea != nil {
					return _ea
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0074\u0069\u0074l\u0065"}:
				_bag.Title = new(_c.XSDAny)
				if _fcb := d.DecodeElement(_bag.Title, &_aae); _fcb != nil {
					return _fcb
				}
			case _b.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", Local: "\u0076e\u0072\u0073\u0069\u006f\u006e"}:
				_bag.Version = new(string)
				if _cda := d.DecodeElement(_bag.Version, &_aae); _cda != nil {
					return _cda
				}
			default:
				_e.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072t\u0069e\u0073\u0020\u0025\u0076", _aae.Name)
				if _edb := d.Skip(); _edb != nil {
					return _edb
				}
			}
		case _b.EndElement:
			break _dca
		case _b.CharData:
		}
	}
	return nil
}

func (_bbaa *CoreProperties) MarshalXML(e *_b.Encoder, start _b.StartElement) error {
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0063\u0070"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"}, Value: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"})
	start.Attr = append(start.Attr, _b.Attr{Name: _b.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0063\u0070\u003a\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073"
	return _bbaa.CT_CoreProperties.MarshalXML(e, start)
}

func init() {
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", "\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073", NewCT_CoreProperties)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", "C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073", NewCT_Keywords)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", "\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064", NewCT_Keyword)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073", "\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073", NewCoreProperties)
}
