package elements

import (
	_ee "encoding/xml"
	_c "fmt"

	_f "github.com/arcpd/msword"
	_a "github.com/arcpd/msword/common/logger"
)

func (_dg *SimpleLiteral) UnmarshalXML(d *_ee.Decoder, start _ee.StartElement) error {
	for {
		_dga, _ef := d.Token()
		if _ef != nil {
			return _c.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s", _ef)
		}
		if _ebe, _gga := _dga.(_ee.EndElement); _gga && _ebe.Name == start.Name {
			break
		}
	}
	return nil
}

func (_fa *Any) MarshalXML(e *_ee.Encoder, start _ee.StartElement) error {
	return _fa.SimpleLiteral.MarshalXML(e, start)
}

func NewAny() *Any { _b := &Any{}; _b.SimpleLiteral = *NewSimpleLiteral(); return _b }

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_eg *ElementContainer) ValidateWithPath(path string) error {
	for _eb, _g := range _eg.Choice {
		if _fb := _g.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _eb)); _fb != nil {
			return _fb
		}
	}
	return nil
}

// Validate validates the Any and its children
func (_bc *Any) Validate() error { return _bc.ValidateWithPath("\u0041\u006e\u0079") }

// Validate validates the ElementContainer and its children
func (_abe *ElementContainer) Validate() error {
	return _abe.ValidateWithPath("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072")
}

// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_agd *ElementsGroupChoice) ValidateWithPath(path string) error {
	for _fcc, _bge := range _agd.Any {
		if _ce := _bge.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d", path, _fcc)); _ce != nil {
			return _ce
		}
	}
	return nil
}

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_bf *Any) ValidateWithPath(path string) error {
	if _ed := _bf.SimpleLiteral.ValidateWithPath(path); _ed != nil {
		return _ed
	}
	return nil
}

type ElementsGroup struct{ Choice []*ElementsGroupChoice }

func NewSimpleLiteral() *SimpleLiteral { _ga := &SimpleLiteral{}; return _ga }

type Any struct{ SimpleLiteral }

func (_bbc *SimpleLiteral) MarshalXML(e *_ee.Encoder, start _ee.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(_ee.EndElement{Name: start.Name})
	return nil
}

func (_dd *ElementsGroup) MarshalXML(e *_ee.Encoder, start _ee.StartElement) error {
	if _dd.Choice != nil {
		for _, _fc := range _dd.Choice {
			_fc.MarshalXML(e, _ee.StartElement{})
		}
	}
	return nil
}

type SimpleLiteral struct{}

// Validate validates the ElementsGroup and its children
func (_db *ElementsGroup) Validate() error {
	return _db.ValidateWithPath("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070")
}

func NewElementsGroupChoice() *ElementsGroupChoice { _ea := &ElementsGroupChoice{}; return _ea }

func NewElementContainer() *ElementContainer { _ac := &ElementContainer{}; return _ac }

type ElementsGroupChoice struct{ Any []*Any }

func (_acg *ElementsGroupChoice) UnmarshalXML(d *_ee.Decoder, start _ee.StartElement) error {
_fcd:
	for {
		_fdd, _gdfa := d.Token()
		if _gdfa != nil {
			return _gdfa
		}
		switch _ae := _fdd.(type) {
		case _ee.StartElement:
			switch _ae.Name {
			case _ee.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_fac := NewAny()
				if _gg := d.DecodeElement(_fac, &_ae); _gg != nil {
					return _gg
				}
				_acg.Any = append(_acg.Any, _fac)
			default:
				_a.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _ae.Name)
				if _acf := d.Skip(); _acf != nil {
					return _acf
				}
			}
		case _ee.EndElement:
			break _fcd
		case _ee.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_cag *ElementsGroup) ValidateWithPath(path string) error {
	for _gdf, _ag := range _cag.Choice {
		if _bfa := _ag.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _gdf)); _bfa != nil {
			return _bfa
		}
	}
	return nil
}

type ElementContainer struct{ Choice []*ElementsGroupChoice }

// Validate validates the ElementsGroupChoice and its children
func (_fbf *ElementsGroupChoice) Validate() error {
	return _fbf.ValidateWithPath("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065")
}

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_agb *SimpleLiteral) ValidateWithPath(path string) error { return nil }

func (_ad *Any) UnmarshalXML(d *_ee.Decoder, start _ee.StartElement) error {
	_ad.SimpleLiteral = *NewSimpleLiteral()
	for {
		_ca, _adc := d.Token()
		if _adc != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073", _adc)
		}
		if _ff, _ec := _ca.(_ee.EndElement); _ec && _ff.Name == start.Name {
			break
		}
	}
	return nil
}

func (_d *ElementContainer) UnmarshalXML(d *_ee.Decoder, start _ee.StartElement) error {
_aa:
	for {
		_be, _fg := d.Token()
		if _fg != nil {
			return _fg
		}
		switch _bd := _be.(type) {
		case _ee.StartElement:
			switch _bd.Name {
			case _ee.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_cc := NewElementsGroupChoice()
				if _ab := d.DecodeElement(&_cc.Any, &_bd); _ab != nil {
					return _ab
				}
				_d.Choice = append(_d.Choice, _cc)
			default:
				_a.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v", _bd.Name)
				if _ba := d.Skip(); _ba != nil {
					return _ba
				}
			}
		case _ee.EndElement:
			break _aa
		case _ee.CharData:
		}
	}
	return nil
}

func (_cd *ElementsGroupChoice) MarshalXML(e *_ee.Encoder, start _ee.StartElement) error {
	if _cd.Any != nil {
		_fd := _ee.StartElement{Name: _ee.Name{Local: "\u0064\u0063\u003a\u0061\u006e\u0079"}}
		for _, _cdc := range _cd.Any {
			e.EncodeElement(_cdc, _fd)
		}
	}
	return nil
}

func (_ece *ElementContainer) MarshalXML(e *_ee.Encoder, start _ee.StartElement) error {
	start.Name.Local = "\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072"
	e.EncodeToken(start)
	if _ece.Choice != nil {
		for _, _ffd := range _ece.Choice {
			_ffd.MarshalXML(e, _ee.StartElement{})
		}
	}
	e.EncodeToken(_ee.EndElement{Name: start.Name})
	return nil
}

func NewElementsGroup() *ElementsGroup { _ge := &ElementsGroup{}; return _ge }

func (_gd *ElementsGroup) UnmarshalXML(d *_ee.Decoder, start _ee.StartElement) error {
_edf:
	for {
		_abeb, _ddc := d.Token()
		if _ddc != nil {
			return _ddc
		}
		switch _bb := _abeb.(type) {
		case _ee.StartElement:
			switch _bb.Name {
			case _ee.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_fce := NewElementsGroupChoice()
				if _eda := d.DecodeElement(&_fce.Any, &_bb); _eda != nil {
					return _eda
				}
				_gd.Choice = append(_gd.Choice, _fce)
			default:
				_a.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076", _bb.Name)
				if _ccg := d.Skip(); _ccg != nil {
					return _ccg
				}
			}
		case _ee.EndElement:
			break _edf
		case _ee.CharData:
		}
	}
	return nil
}

// Validate validates the SimpleLiteral and its children
func (_aaf *SimpleLiteral) Validate() error {
	return _aaf.ValidateWithPath("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c")
}

func init() {
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", "\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c", NewSimpleLiteral)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", "\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072", NewElementContainer)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", "\u0061\u006e\u0079", NewAny)
	_f.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", "\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070", NewElementsGroup)
}
