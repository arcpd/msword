package content_types

import (
	_c "encoding/xml"
	_g "fmt"
	_e "regexp"

	_b "github.com/arcpd/msword"
	_gf "github.com/arcpd/msword/common/logger"
)

type Override struct{ CT_Override }

func NewCT_Default() *CT_Default {
	_ag := &CT_Default{}
	_ag.ExtensionAttr = "\u0078\u006d\u006c"
	_ag.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	return _ag
}

var ST_ExtensionPatternRe = _e.MustCompile(ST_ExtensionPattern)

func NewCT_Override() *CT_Override {
	_cde := &CT_Override{}
	_cde.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	return _cde
}

func NewCT_Types() *CT_Types { _fgg := &CT_Types{}; return _fgg }

func (_cb *Override) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	return _cb.CT_Override.MarshalXML(e, start)
}

func NewDefault() *Default { _dc := &Default{}; _dc.CT_Default = *NewCT_Default(); return _dc }

func (_af *CT_Override) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_af.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	for _, _d := range start.Attr {
		if _d.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_fc, _cc := _d.Value, error(nil)
			if _cc != nil {
				return _cc
			}
			_af.ContentTypeAttr = _fc
			continue
		}
		if _d.Name.Local == "\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065" {
			_gb, _ce := _d.Value, error(nil)
			if _ce != nil {
				return _ce
			}
			_af.PartNameAttr = _gb
			continue
		}
	}
	for {
		_dd, _bgb := d.Token()
		if _bgb != nil {
			return _g.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025\u0073", _bgb)
		}
		if _bge, _gc := _dd.(_c.EndElement); _gc && _bge.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ea *CT_Default) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"}, Value: _g.Sprintf("\u0025\u0076", _ea.ExtensionAttr)})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"}, Value: _g.Sprintf("\u0025\u0076", _ea.ContentTypeAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

func NewOverride() *Override { _bdc := &Override{}; _bdc.CT_Override = *NewCT_Override(); return _bdc }

// Validate validates the CT_Override and its children
func (_ab *CT_Override) Validate() error {
	return _ab.ValidateWithPath("C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065")
}

// Validate validates the Default and its children
func (_acf *Default) Validate() error {
	return _acf.ValidateWithPath("\u0044e\u0066\u0061\u0075\u006c\u0074")
}

type CT_Default struct {
	ExtensionAttr   string
	ContentTypeAttr string
}

func (_ge *CT_Types) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	e.EncodeToken(start)
	if _ge.Default != nil {
		_geg := _c.StartElement{Name: _c.Name{Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}}
		for _, _gef := range _ge.Default {
			e.EncodeElement(_gef, _geg)
		}
	}
	if _ge.Override != nil {
		_fda := _c.StartElement{Name: _c.Name{Local: "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}}
		for _, _cg := range _ge.Override {
			e.EncodeElement(_cg, _fda)
		}
	}
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

func (_ebc *Types) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s"})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0054\u0079\u0070e\u0073"
	return _ebc.CT_Types.MarshalXML(e, start)
}

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (_gega *CT_Types) ValidateWithPath(path string) error {
	for _gg, _daf := range _gega.Default {
		if _eba := _daf.ValidateWithPath(_g.Sprintf("\u0025\u0073\u002f\u0044\u0065\u0066\u0061\u0075\u006ct\u005b\u0025\u0064\u005d", path, _gg)); _eba != nil {
			return _eba
		}
	}
	for _ddga, _be := range _gega.Override {
		if _ccb := _be.ValidateWithPath(_g.Sprintf("\u0025s\u002fO\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u005b\u0025\u0064\u005d", path, _ddga)); _ccb != nil {
			return _ccb
		}
	}
	return nil
}

const ST_ExtensionPattern = "\u0028\u005b\u0021\u0024\u0026\u0027\\\u0028\u005c\u0029\u005c\u002a\\\u002b\u002c\u003a\u003d\u005d\u007c(\u0025\u005b\u0030\u002d\u0039a\u002d\u0066\u0041\u002d\u0046\u005d\u005b\u0030\u002d\u0039\u0061\u002df\u0041\u002d\u0046\u005d\u0029\u007c\u005b\u003a\u0040\u005d\u007c\u005b\u0061\u002d\u007aA\u002d\u005a\u0030\u002d\u0039\u005c\u002d\u005f~\u005d\u0029\u002b"

const ST_ContentTypePattern = "\u005e\\\u0070{\u004c\u0061\u0074\u0069\u006e\u007d\u002b\u002f\u002e\u002a\u0024"

func (_agc *CT_Default) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_agc.ExtensionAttr = "\u0078\u006d\u006c"
	_agc.ContentTypeAttr = "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c"
	for _, _ee := range start.Attr {
		if _ee.Name.Local == "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn" {
			_bg, _cd := _ee.Value, error(nil)
			if _cd != nil {
				return _cd
			}
			_agc.ExtensionAttr = _bg
			continue
		}
		if _ee.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_bc, _fg := _ee.Value, error(nil)
			if _fg != nil {
				return _fg
			}
			_agc.ContentTypeAttr = _bc
			continue
		}
	}
	for {
		_ed, _ad := d.Token()
		if _ad != nil {
			return _g.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020%\u0073", _ad)
		}
		if _bd, _eea := _ed.(_c.EndElement); _eea && _bd.Name == start.Name {
			break
		}
	}
	return nil
}

func (_afc *Default) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_afc.CT_Default = *NewCT_Default()
	for _, _bf := range start.Attr {
		if _bf.Name.Local == "\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn" {
			_beg, _dbe := _bf.Value, error(nil)
			if _dbe != nil {
				return _dbe
			}
			_afc.ExtensionAttr = _beg
			continue
		}
		if _bf.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_bfg, _ac := _bf.Value, error(nil)
			if _ac != nil {
				return _ac
			}
			_afc.ContentTypeAttr = _bfg
			continue
		}
	}
	for {
		_eg, _aa := d.Token()
		if _aa != nil {
			return _g.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020\u0025\u0073", _aa)
		}
		if _bcd, _bbg := _eg.(_c.EndElement); _bbg && _bcd.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Default and its children, prefixing error messages with path
func (_bgc *CT_Default) ValidateWithPath(path string) error {
	if !ST_ExtensionPatternRe.MatchString(_bgc.ExtensionAttr) {
		return _g.Errorf("\u0025s\u002f\u006d.\u0045\u0078\u0074\u0065n\u0073\u0069\u006fn\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 m\u0061\u0074\u0063h\u0020\u0027%\u0073\u0027\u0020\u0028\u0068\u0061v\u0065\u0020%\u0076\u0029", path, ST_ExtensionPatternRe, _bgc.ExtensionAttr)
	}
	if !ST_ContentTypePatternRe.MatchString(_bgc.ContentTypeAttr) {
		return _g.Errorf("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, ST_ContentTypePatternRe, _bgc.ContentTypeAttr)
	}
	return nil
}

type CT_Types struct {
	Default  []*Default
	Override []*Override
}

// Validate validates the Types and its children
func (_gefa *Types) Validate() error { return _gefa.ValidateWithPath("\u0054\u0079\u0070e\u0073") }

// Validate validates the CT_Types and its children
func (_daa *CT_Types) Validate() error {
	return _daa.ValidateWithPath("\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073")
}

type Types struct{ CT_Types }

func (_ae *Default) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	return _ae.CT_Default.MarshalXML(e, start)
}

func NewTypes() *Types { _ged := &Types{}; _ged.CT_Types = *NewCT_Types(); return _ged }

var ST_ContentTypePatternRe = _e.MustCompile(ST_ContentTypePattern)

// Validate validates the CT_Default and its children
func (_eaa *CT_Default) Validate() error {
	return _eaa.ValidateWithPath("\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074")
}

type CT_Override struct {
	ContentTypeAttr string
	PartNameAttr    string
}

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (_fgf *Override) ValidateWithPath(path string) error {
	if _bgf := _fgf.CT_Override.ValidateWithPath(path); _bgf != nil {
		return _bgf
	}
	return nil
}

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (_cgd *Types) ValidateWithPath(path string) error {
	if _aga := _cgd.CT_Types.ValidateWithPath(path); _aga != nil {
		return _aga
	}
	return nil
}

// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (_aff *CT_Override) ValidateWithPath(path string) error {
	if !ST_ContentTypePatternRe.MatchString(_aff.ContentTypeAttr) {
		return _g.Errorf("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, ST_ContentTypePatternRe, _aff.ContentTypeAttr)
	}
	return nil
}

func (_ca *Override) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_ca.CT_Override = *NewCT_Override()
	for _, _ebg := range start.Attr {
		if _ebg.Name.Local == "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065" {
			_cge, _cgee := _ebg.Value, error(nil)
			if _cgee != nil {
				return _cgee
			}
			_ca.ContentTypeAttr = _cge
			continue
		}
		if _ebg.Name.Local == "\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065" {
			_aec, _edg := _ebg.Value, error(nil)
			if _edg != nil {
				return _edg
			}
			_ca.PartNameAttr = _aec
			continue
		}
	}
	for {
		_feg, _edc := d.Token()
		if _edc != nil {
			return _g.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u004f\u0076\u0065r\u0072\u0069\u0064\u0065: \u0025\u0073", _edc)
		}
		if _acg, _aeg := _feg.(_c.EndElement); _aeg && _acg.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ddg *CT_Types) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
_da:
	for {
		_bcg, _db := d.Token()
		if _db != nil {
			return _db
		}
		switch _bdd := _bcg.(type) {
		case _c.StartElement:
			switch _bdd.Name {
			case _c.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_cgc := NewDefault()
				if _eb := d.DecodeElement(_cgc, &_bdd); _eb != nil {
					return _eb
				}
				_ddg.Default = append(_ddg.Default, _cgc)
			case _c.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:
				_eac := NewOverride()
				if _ddc := d.DecodeElement(_eac, &_bdd); _ddc != nil {
					return _ddc
				}
				_ddg.Override = append(_ddg.Override, _eac)
			default:
				_gf.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073\u0020\u0025\u0076", _bdd.Name)
				if _bgg := d.Skip(); _bgg != nil {
					return _bgg
				}
			}
		case _c.EndElement:
			break _da
		case _c.CharData:
		}
	}
	return nil
}

// Validate validates the Override and its children
func (_begf *Override) Validate() error {
	return _begf.ValidateWithPath("\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065")
}

// ValidateWithPath validates the Default and its children, prefixing error messages with path
func (_bff *Default) ValidateWithPath(path string) error {
	if _cdee := _bff.CT_Default.ValidateWithPath(path); _cdee != nil {
		return _cdee
	}
	return nil
}

func (_gd *Types) UnmarshalXML(d *_c.Decoder, start _c.StartElement) error {
	_gd.CT_Types = *NewCT_Types()
_cga:
	for {
		_dae, _ebcf := d.Token()
		if _ebcf != nil {
			return _ebcf
		}
		switch _ccg := _dae.(type) {
		case _c.StartElement:
			switch _ccg.Name {
			case _c.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u0044e\u0066\u0061\u0075\u006c\u0074"}:
				_df := NewDefault()
				if _bbge := d.DecodeElement(_df, &_ccg); _bbge != nil {
					return _bbge
				}
				_gd.Default = append(_gd.Default, _df)
			case _c.Name{Space: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", Local: "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:
				_caf := NewOverride()
				if _ggd := d.DecodeElement(_caf, &_ccg); _ggd != nil {
					return _ggd
				}
				_gd.Override = append(_gd.Override, _caf)
			default:
				_gf.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0054\u0079\u0070e\u0073 \u0025\u0076", _ccg.Name)
				if _adf := d.Skip(); _adf != nil {
					return _adf
				}
			}
		case _c.EndElement:
			break _cga
		case _c.CharData:
		}
	}
	return nil
}

type Default struct{ CT_Default }

func (_cf *CT_Override) MarshalXML(e *_c.Encoder, start _c.StartElement) error {
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"}, Value: _g.Sprintf("\u0025\u0076", _cf.ContentTypeAttr)})
	start.Attr = append(start.Attr, _c.Attr{Name: _c.Name{Local: "\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"}, Value: _g.Sprintf("\u0025\u0076", _cf.PartNameAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_c.EndElement{Name: start.Name})
	return nil
}

func init() {
	_b.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073", NewCT_Types)
	_b.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074", NewCT_Default)
	_b.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065", NewCT_Override)
	_b.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0054\u0079\u0070e\u0073", NewTypes)
	_b.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u0044e\u0066\u0061\u0075\u006c\u0074", NewDefault)
	_b.RegisterConstructor("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s", "\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065", NewOverride)
}
