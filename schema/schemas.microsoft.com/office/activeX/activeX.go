package activeX

import (
	_f "encoding/xml"
	_a "fmt"
	_bc "github.com/arcpd/msword"
	_b "github.com/arcpd/msword/common/logger"
)

type CT_OcxPrChoice struct {
	Font    *CT_Font
	Picture *CT_Picture
}

type ST_Persistence byte

func NewCT_Picture() *CT_Picture { _deg := &CT_Picture{}; return _deg }

func (_e *CT_Font) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	if _e.PersistenceAttr != ST_PersistenceUnset {
		_ec, _ae := _e.PersistenceAttr.MarshalXMLAttr(_f.Name{Local: "\u0061\u0078\u003a\u0070\u0065\u0072\u0073\u0069\u0073t\u0065\u006e\u0063\u0065"})
		if _ae != nil {
			return _ae
		}
		start.Attr = append(start.Attr, _ec)
	}
	if _e.IdAttr != nil {
		start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0072\u003a\u0069\u0064"}, Value: _a.Sprintf("\u0025\u0076", *_e.IdAttr)})
	}
	e.EncodeToken(start)
	if _e.OcxPr != nil {
		_c := _f.StartElement{Name: _f.Name{Local: "\u0061\u0078\u003a\u006f\u0063\u0078\u0050\u0072"}}
		for _, _cd := range _e.OcxPr {
			e.EncodeElement(_cd, _c)
		}
	}
	e.EncodeToken(_f.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the Ocx and its children, prefixing error messages with path
func (_egb *Ocx) ValidateWithPath(path string) error {
	if _ccc := _egb.CT_Ocx.ValidateWithPath(path); _ccc != nil {
		return _ccc
	}
	return nil
}

type CT_Picture struct{ IdAttr *string }

type CT_Ocx struct {
	ClassidAttr     string
	LicenseAttr     *string
	IdAttr          *string
	PersistenceAttr ST_Persistence
	OcxPr           []*CT_OcxPr
}

type CT_Font struct {
	PersistenceAttr ST_Persistence
	IdAttr          *string
	OcxPr           []*CT_OcxPr
}

func (_ba *CT_Font) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	for _, _db := range start.Attr {
		if _db.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _db.Name.Local == "\u0069\u0064" || _db.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _db.Name.Local == "\u0069\u0064" {
			_bca, _ed := _db.Value, error(nil)
			if _ed != nil {
				return _ed
			}
			_ba.IdAttr = &_bca
			continue
		}
		if _db.Name.Local == "p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065" {
			_ba.PersistenceAttr.UnmarshalXMLAttr(_db)
			continue
		}
	}
_fa:
	for {
		_eca, _cdd := d.Token()
		if _cdd != nil {
			return _cdd
		}
		switch _g := _eca.(type) {
		case _f.StartElement:
			switch _g.Name {
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", Local: "\u006f\u0063\u0078P\u0072"}:
				_bg := NewCT_OcxPr()
				if _eab := d.DecodeElement(_bg, &_g); _eab != nil {
					return _eab
				}
				_ba.OcxPr = append(_ba.OcxPr, _bg)
			default:
				_b.Log.Debug("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0046\u006f\u006e\u0074\u0020\u0025\u0076", _g.Name)
				if _ee := d.Skip(); _ee != nil {
					return _ee
				}
			}
		case _f.EndElement:
			break _fa
		case _f.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_OcxPrChoice and its children, prefixing error messages with path
func (_agg *CT_OcxPrChoice) ValidateWithPath(path string) error {
	if _agg.Font != nil {
		if _cba := _agg.Font.ValidateWithPath(path + "\u002f\u0046\u006fn\u0074"); _cba != nil {
			return _cba
		}
	}
	if _agg.Picture != nil {
		if _feda := _agg.Picture.ValidateWithPath(path + "\u002f\u0050\u0069\u0063\u0074\u0075\u0072\u0065"); _feda != nil {
			return _feda
		}
	}
	return nil
}

type Ocx struct{ CT_Ocx }

// Validate validates the CT_OcxPrChoice and its children
func (_fbd *CT_OcxPrChoice) Validate() error {
	return _fbd.ValidateWithPath("\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0043h\u006f\u0069\u0063\u0065")
}

// Validate validates the CT_OcxPr and its children
func (_ge *CT_OcxPr) Validate() error {
	return _ge.ValidateWithPath("\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072")
}

func (_afa ST_Persistence) Validate() error { return _afa.ValidateWithPath("") }

func NewCT_OcxPr() *CT_OcxPr { _eg := &CT_OcxPr{}; return _eg }

func NewCT_Font() *CT_Font { _fe := &CT_Font{}; return _fe }

func (_fea *CT_OcxPrChoice) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
_ddb:
	for {
		_ebg, _bb := d.Token()
		if _bb != nil {
			return _bb
		}
		switch _cf := _ebg.(type) {
		case _f.StartElement:
			switch _cf.Name {
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", Local: "\u0066\u006f\u006e\u0074"}:
				_fea.Font = NewCT_Font()
				if _ebf := d.DecodeElement(_fea.Font, &_cf); _ebf != nil {
					return _ebf
				}
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", Local: "\u0070i\u0063\u0074\u0075\u0072\u0065"}:
				_fea.Picture = NewCT_Picture()
				if _fae := d.DecodeElement(_fea.Picture, &_cf); _fae != nil {
					return _fae
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0043\u0068o\u0069c\u0065\u0020\u0025\u0076", _cf.Name)
				if _cad := d.Skip(); _cad != nil {
					return _cad
				}
			}
		case _f.EndElement:
			break _ddb
		case _f.CharData:
		}
	}
	return nil
}

func NewCT_OcxPrChoice() *CT_OcxPrChoice { _cb := &CT_OcxPrChoice{}; return _cb }

// Validate validates the CT_Ocx and its children
func (_dbd *CT_Ocx) Validate() error {
	return _dbd.ValidateWithPath("\u0043\u0054\u005f\u004f\u0063\u0078")
}

// ValidateWithPath validates the CT_OcxPr and its children, prefixing error messages with path
func (_ddg *CT_OcxPr) ValidateWithPath(path string) error {
	if _ddg.Choice != nil {
		if _egcc := _ddg.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _egcc != nil {
			return _egcc
		}
	}
	return nil
}

// Validate validates the Ocx and its children
func (_cfcd *Ocx) Validate() error { return _cfcd.ValidateWithPath("\u004f\u0063\u0078") }

func NewOcx() *Ocx { _gad := &Ocx{}; _gad.CT_Ocx = *NewCT_Ocx(); return _gad }

func (_be *CT_Ocx) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_be.PersistenceAttr = ST_Persistence(1)
	for _, _df := range start.Attr {
		if _df.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _df.Name.Local == "\u0069\u0064" || _df.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _df.Name.Local == "\u0069\u0064" {
			_fd, _bd := _df.Value, error(nil)
			if _bd != nil {
				return _bd
			}
			_be.IdAttr = &_fd
			continue
		}
		if _df.Name.Local == "\u0063l\u0061\u0073\u0073\u0069\u0064" {
			_ad, _cg := _df.Value, error(nil)
			if _cg != nil {
				return _cg
			}
			_be.ClassidAttr = _ad
			continue
		}
		if _df.Name.Local == "\u006ci\u0063\u0065\u006e\u0073\u0065" {
			_adb, _ga := _df.Value, error(nil)
			if _ga != nil {
				return _ga
			}
			_be.LicenseAttr = &_adb
			continue
		}
		if _df.Name.Local == "p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065" {
			_be.PersistenceAttr.UnmarshalXMLAttr(_df)
			continue
		}
	}
_abe:
	for {
		_eacb, _eee := d.Token()
		if _eee != nil {
			return _eee
		}
		switch _gaa := _eacb.(type) {
		case _f.StartElement:
			switch _gaa.Name {
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", Local: "\u006f\u0063\u0078P\u0072"}:
				_cc := NewCT_OcxPr()
				if _fef := d.DecodeElement(_cc, &_gaa); _fef != nil {
					return _fef
				}
				_be.OcxPr = append(_be.OcxPr, _cc)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0020\u0025\u0076", _gaa.Name)
				if _gf := d.Skip(); _gf != nil {
					return _gf
				}
			}
		case _f.EndElement:
			break _abe
		case _f.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Font and its children
func (_fed *CT_Font) Validate() error {
	return _fed.ValidateWithPath("\u0043T\u005f\u0046\u006f\u006e\u0074")
}

func (_def *ST_Persistence) UnmarshalXMLAttr(attr _f.Attr) error {
	switch attr.Value {
	case "":
		*_def = 0
	case "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067":
		*_def = 1
	case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d":
		*_def = 2
	case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074":
		*_def = 3
	case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065":
		*_def = 4
	}
	return nil
}

type CT_OcxPr struct {
	NameAttr  string
	ValueAttr *string
	Choice    *CT_OcxPrChoice
}

func (_gdb ST_Persistence) ValidateWithPath(path string) error {
	switch _gdb {
	case 0, 1, 2, 3, 4:
	default:
		return _a.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gdb))
	}
	return nil
}

const (
	ST_PersistenceUnset              ST_Persistence = 0
	ST_PersistencePersistPropertyBag ST_Persistence = 1
	ST_PersistencePersistStream      ST_Persistence = 2
	ST_PersistencePersistStreamInit  ST_Persistence = 3
	ST_PersistencePersistStorage     ST_Persistence = 4
)

// ValidateWithPath validates the CT_Font and its children, prefixing error messages with path
func (_bcad *CT_Font) ValidateWithPath(path string) error {
	if _bgb := _bcad.PersistenceAttr.ValidateWithPath(path + "\u002f\u0050e\u0072\u0073\u0069s\u0074\u0065\u006e\u0063\u0065\u0041\u0074\u0074\u0072"); _bgb != nil {
		return _bgb
	}
	for _eac, _eed := range _bcad.OcxPr {
		if _fc := _eed.ValidateWithPath(_a.Sprintf("\u0025\u0073\u002fO\u0063\u0078\u0050\u0072\u005b\u0025\u0064\u005d", path, _eac)); _fc != nil {
			return _fc
		}
	}
	return nil
}

func (_defc ST_Persistence) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	return e.EncodeElement(_defc.String(), start)
}

func (_adg *CT_OcxPrChoice) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	if _adg.Font != nil {
		_ebe := _f.StartElement{Name: _f.Name{Local: "\u0061x\u003a\u0066\u006f\u006e\u0074"}}
		e.EncodeElement(_adg.Font, _ebe)
	}
	if _adg.Picture != nil {
		_bf := _f.StartElement{Name: _f.Name{Local: "\u0061\u0078\u003a\u0070\u0069\u0063\u0074\u0075\u0072\u0065"}}
		e.EncodeElement(_adg.Picture, _bf)
	}
	return nil
}

func (_dgf ST_Persistence) String() string {
	switch _dgf {
	case 0:
		return ""
	case 1:
		return "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067"
	case 2:
		return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d"
	case 3:
		return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074"
	case 4:
		return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065"
	}
	return ""
}

func (_bbe *CT_Picture) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	for _, _bcag := range start.Attr {
		if _bcag.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _bcag.Name.Local == "\u0069\u0064" || _bcag.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _bcag.Name.Local == "\u0069\u0064" {
			_gg, _ada := _bcag.Value, error(nil)
			if _ada != nil {
				return _ada
			}
			_bbe.IdAttr = &_gg
			continue
		}
	}
	for {
		_ef, _dc := d.Token()
		if _dc != nil {
			return _a.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065\u003a\u0020%\u0073", _dc)
		}
		if _gceg, _ddd := _ef.(_f.EndElement); _ddd && _gceg.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ff ST_Persistence) MarshalXMLAttr(name _f.Name) (_f.Attr, error) {
	_gbg := _f.Attr{}
	_gbg.Name = name
	switch _ff {
	case ST_PersistenceUnset:
		_gbg.Value = ""
	case ST_PersistencePersistPropertyBag:
		_gbg.Value = "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067"
	case ST_PersistencePersistStream:
		_gbg.Value = "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d"
	case ST_PersistencePersistStreamInit:
		_gbg.Value = "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074"
	case ST_PersistencePersistStorage:
		_gbg.Value = "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065"
	}
	return _gbg, nil
}

func (_ddgg *Ocx) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058"})
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0061\u0078"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058"})
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0061\u0078\u003a\u006f\u0063\u0078"
	return _ddgg.CT_Ocx.MarshalXML(e, start)
}

func (_dd *CT_OcxPr) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0061x\u003a\u006e\u0061\u006d\u0065"}, Value: _a.Sprintf("\u0025\u0076", _dd.NameAttr)})
	if _dd.ValueAttr != nil {
		start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0061\u0078\u003a\u0076\u0061\u006c\u0075\u0065"}, Value: _a.Sprintf("\u0025\u0076", *_dd.ValueAttr)})
	}
	e.EncodeToken(start)
	if _dd.Choice != nil {
		_dd.Choice.MarshalXML(e, _f.StartElement{})
	}
	e.EncodeToken(_f.EndElement{Name: start.Name})
	return nil
}

func (_fcc *CT_Ocx) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0061\u0078\u003a\u0063\u006c\u0061\u0073\u0073\u0069\u0064"}, Value: _a.Sprintf("\u0025\u0076", _fcc.ClassidAttr)})
	if _fcc.LicenseAttr != nil {
		start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0061\u0078\u003a\u006c\u0069\u0063\u0065\u006e\u0073\u0065"}, Value: _a.Sprintf("\u0025\u0076", *_fcc.LicenseAttr)})
	}
	if _fcc.IdAttr != nil {
		start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0072\u003a\u0069\u0064"}, Value: _a.Sprintf("\u0025\u0076", *_fcc.IdAttr)})
	}
	_ecd, _gc := _fcc.PersistenceAttr.MarshalXMLAttr(_f.Name{Local: "\u0061\u0078\u003a\u0070\u0065\u0072\u0073\u0069\u0073t\u0065\u006e\u0063\u0065"})
	if _gc != nil {
		return _gc
	}
	start.Attr = append(start.Attr, _ecd)
	e.EncodeToken(start)
	if _fcc.OcxPr != nil {
		_bad := _f.StartElement{Name: _f.Name{Local: "\u0061\u0078\u003a\u006f\u0063\u0078\u0050\u0072"}}
		for _, _ag := range _fcc.OcxPr {
			e.EncodeElement(_ag, _bad)
		}
	}
	e.EncodeToken(_f.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Ocx and its children, prefixing error messages with path
func (_fad *CT_Ocx) ValidateWithPath(path string) error {
	if _fad.PersistenceAttr == ST_PersistenceUnset {
		return _a.Errorf("\u0025\u0073\u002f\u0050\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063e\u0041\u0074\u0074\u0072\u0020\u0069s\u0020\u0061\u0020\u006d\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020f\u0069\u0065\u006c\u0064", path)
	}
	if _beb := _fad.PersistenceAttr.ValidateWithPath(path + "\u002f\u0050e\u0072\u0073\u0069s\u0074\u0065\u006e\u0063\u0065\u0041\u0074\u0074\u0072"); _beb != nil {
		return _beb
	}
	for _gb, _ca := range _fad.OcxPr {
		if _cgd := _ca.ValidateWithPath(_a.Sprintf("\u0025\u0073\u002fO\u0063\u0078\u0050\u0072\u005b\u0025\u0064\u005d", path, _gb)); _cgd != nil {
			return _cgd
		}
	}
	return nil
}

func NewCT_Ocx() *CT_Ocx { _edd := &CT_Ocx{}; _edd.PersistenceAttr = ST_Persistence(1); return _edd }

// Validate validates the CT_Picture and its children
func (_cfd *CT_Picture) Validate() error {
	return _cfd.ValidateWithPath("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065")
}

func (_cbc *ST_Persistence) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_bda, _cbe := d.Token()
	if _cbe != nil {
		return _cbe
	}
	if _gcea, _gag := _bda.(_f.EndElement); _gag && _gcea.Name == start.Name {
		*_cbc = 1
		return nil
	}
	if _dca, _dgb := _bda.(_f.CharData); !_dgb {
		return _a.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bda)
	} else {
		switch string(_dca) {
		case "":
			*_cbc = 0
		case "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067":
			*_cbc = 1
		case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d":
			*_cbc = 2
		case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074":
			*_cbc = 3
		case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065":
			*_cbc = 4
		}
	}
	_bda, _cbe = d.Token()
	if _cbe != nil {
		return _cbe
	}
	if _ffb, _gfg := _bda.(_f.EndElement); _gfg && _ffb.Name == start.Name {
		return nil
	}
	return _a.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bda)
}

func (_gac *CT_Picture) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	if _gac.IdAttr != nil {
		start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0072\u003a\u0069\u0064"}, Value: _a.Sprintf("\u0025\u0076", *_gac.IdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_f.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_cab *CT_Picture) ValidateWithPath(path string) error { return nil }

func (_gdd *Ocx) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_gdd.CT_Ocx = *NewCT_Ocx()
	for _, _abb := range start.Attr {
		if _abb.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _abb.Name.Local == "\u0069\u0064" || _abb.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _abb.Name.Local == "\u0069\u0064" {
			_fbb, _cfdd := _abb.Value, error(nil)
			if _cfdd != nil {
				return _cfdd
			}
			_gdd.IdAttr = &_fbb
			continue
		}
		if _abb.Name.Local == "\u0063l\u0061\u0073\u0073\u0069\u0064" {
			_ce, _fead := _abb.Value, error(nil)
			if _fead != nil {
				return _fead
			}
			_gdd.ClassidAttr = _ce
			continue
		}
		if _abb.Name.Local == "\u006ci\u0063\u0065\u006e\u0073\u0065" {
			_gfb, _dg := _abb.Value, error(nil)
			if _dg != nil {
				return _dg
			}
			_gdd.LicenseAttr = &_gfb
			continue
		}
		if _abb.Name.Local == "p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065" {
			_gdd.PersistenceAttr.UnmarshalXMLAttr(_abb)
			continue
		}
	}
_dce:
	for {
		_bec, _bbed := d.Token()
		if _bbed != nil {
			return _bbed
		}
		switch _gde := _bec.(type) {
		case _f.StartElement:
			switch _gde.Name {
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", Local: "\u006f\u0063\u0078P\u0072"}:
				_dgc := NewCT_OcxPr()
				if _cfc := d.DecodeElement(_dgc, &_gde); _cfc != nil {
					return _cfc
				}
				_gdd.OcxPr = append(_gdd.OcxPr, _dgc)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006fn\u0020\u004fc\u0078\u0020\u0025\u0076", _gde.Name)
				if _cgb := d.Skip(); _cgb != nil {
					return _cgb
				}
			}
		case _f.EndElement:
			break _dce
		case _f.CharData:
		}
	}
	return nil
}

func (_fcd *CT_OcxPr) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	for _, _gae := range start.Attr {
		if _gae.Name.Local == "\u006e\u0061\u006d\u0065" {
			_af, _ead := _gae.Value, error(nil)
			if _ead != nil {
				return _ead
			}
			_fcd.NameAttr = _af
			continue
		}
		if _gae.Name.Local == "\u0076\u0061\u006cu\u0065" {
			_gd, _aec := _gae.Value, error(nil)
			if _aec != nil {
				return _aec
			}
			_fcd.ValueAttr = &_gd
			continue
		}
	}
_cgf:
	for {
		_aff, _bce := d.Token()
		if _bce != nil {
			return _bce
		}
		switch _eb := _aff.(type) {
		case _f.StartElement:
			switch _eb.Name {
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", Local: "\u0066\u006f\u006e\u0074"}:
				_fcd.Choice = NewCT_OcxPrChoice()
				if _gce := d.DecodeElement(&_fcd.Choice.Font, &_eb); _gce != nil {
					return _gce
				}
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", Local: "\u0070i\u0063\u0074\u0075\u0072\u0065"}:
				_fcd.Choice = NewCT_OcxPrChoice()
				if _gff := d.DecodeElement(&_fcd.Choice.Picture, &_eb); _gff != nil {
					return _gff
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0020\u0025\u0076", _eb.Name)
				if _egc := d.Skip(); _egc != nil {
					return _egc
				}
			}
		case _f.EndElement:
			break _cgf
		case _f.CharData:
		}
	}
	return nil
}

func init() {
	_bc.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", "\u0043\u0054\u005f\u004f\u0063\u0078", NewCT_Ocx)
	_bc.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", "\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072", NewCT_OcxPr)
	_bc.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", "\u0043T\u005f\u0046\u006f\u006e\u0074", NewCT_Font)
	_bc.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", "\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065", NewCT_Picture)
	_bc.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058", "\u006f\u0063\u0078", NewOcx)
}
