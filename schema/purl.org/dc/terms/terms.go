package terms

import (
	_af "encoding/xml"
	_c "fmt"

	_f "github.com/arcpd/msword"
	_e "github.com/arcpd/msword/common/logger"
	_cb "github.com/arcpd/msword/schema/purl.org/dc/elements"
)

func (_ebe *ISO639_2) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

type RFC1766 struct{}

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_egff *MESH) ValidateWithPath(path string) error { return nil }

func (_cf *DDC) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_gf, _cd := d.Token()
		if _cd != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0044\u0044\u0043\u003a\u0020\u0025\u0073", _cd)
		}
		if _da, _de := _gf.(_af.EndElement); _de && _da.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_bb *ElementsAndRefinementsGroupChoice) ValidateWithPath(path string) error {
	for _ba, _eed := range _bb.Any {
		if _bg := _eed.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d", path, _ba)); _bg != nil {
			return _bg
		}
	}
	return nil
}

func (_gba *TGN) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_cdde, _ace := d.Token()
		if _ace != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0054\u0047\u004e\u003a\u0020\u0025\u0073", _ace)
		}
		if _ebaa, _adf := _cdde.(_af.EndElement); _adf && _ebaa.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_abb *ISO639_2) ValidateWithPath(path string) error { return nil }

type ISO3166 struct{}

func (_cag *ElementsAndRefinementsGroupChoice) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
_cge:
	for {
		_ga, _aef := d.Token()
		if _aef != nil {
			return _aef
		}
		switch _fdd := _ga.(type) {
		case _af.StartElement:
			switch _fdd.Name {
			case _af.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_aec := _cb.NewAny()
				if _dbca := d.DecodeElement(_aec, &_fdd); _dbca != nil {
					return _dbca
				}
				_cag.Any = append(_cag.Any, _aec)
			default:
				_e.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0041\u006ed\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006fu\u0070\u0043\u0068o\u0069\u0063\u0065\u0020\u0025\u0076", _fdd.Name)
				if _gde := d.Skip(); _gde != nil {
					return _gde
				}
			}
		case _af.EndElement:
			break _cge
		case _af.CharData:
		}
	}
	return nil
}

type ElementsAndRefinementsGroupChoice struct{ Any []*_cb.Any }

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_cc *Box) ValidateWithPath(path string) error { return nil }

func NewDCMIType() *DCMIType { _eb := &DCMIType{}; return _eb }

func (_eaba *LCSH) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u004c\u0043\u0053\u0048"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func NewUDC() *UDC { _efa := &UDC{}; return _efa }

type DDC struct{}

func (_bdf *ElementsAndRefinementsGroup) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
_add:
	for {
		_bfa, _cec := d.Token()
		if _cec != nil {
			return _cec
		}
		switch _bda := _bfa.(type) {
		case _af.StartElement:
			switch _bda.Name {
			case _af.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_dab := NewElementsAndRefinementsGroupChoice()
				if _ggd := d.DecodeElement(&_dab.Any, &_bda); _ggd != nil {
					return _ggd
				}
				_bdf.Choice = append(_bdf.Choice, _dab)
			default:
				_e.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006ce\u006d\u0065\u006e\u0074\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0020\u0025\u0076", _bda.Name)
				if _dgg := d.Skip(); _dgg != nil {
					return _dgg
				}
			}
		case _af.EndElement:
			break _add
		case _af.CharData:
		}
	}
	return nil
}

func (_eg *ElementOrRefinementContainer) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072"
	e.EncodeToken(start)
	if _eg.Choice != nil {
		for _, _agg := range _eg.Choice {
			_agg.MarshalXML(e, _af.StartElement{})
		}
	}
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_eac *LCC) ValidateWithPath(path string) error { return nil }

func (_fff *ISO639_2) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_fb, _eba := d.Token()
		if _eba != nil {
			return _c.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0049\u0053\u004f6\u0033\u0039\u005f\u0032: \u0025\u0073", _eba)
		}
		if _fa, _bfd := _fb.(_af.EndElement); _bfd && _fa.Name == start.Name {
			break
		}
	}
	return nil
}

func NewElementsAndRefinementsGroup() *ElementsAndRefinementsGroup {
	_ad := &ElementsAndRefinementsGroup{}
	return _ad
}

type Period struct{}

func (_gag *Point) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0050\u006f\u0069n\u0074"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

// Validate validates the ISO639_2 and its children
func (_fffe *ISO639_2) Validate() error {
	return _fffe.ValidateWithPath("\u0049\u0053\u004f\u0036\u0033\u0039\u005f\u0032")
}

// Validate validates the Point and its children
func (_fe *Point) Validate() error { return _fe.ValidateWithPath("\u0050\u006f\u0069n\u0074") }

type LCSH struct{}

// Validate validates the Period and its children
func (_abe *Period) Validate() error {
	return _abe.ValidateWithPath("\u0050\u0065\u0072\u0069\u006f\u0064")
}

func NewLCC() *LCC { _abd := &LCC{}; return _abd }

type Point struct{}

func (_eafc *ISO3166) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_fcgb, _ge := d.Token()
		if _ge != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0049\u0053\u004f\u0033\u0031\u0036\u0036\u003a\u0020\u0025\u0073", _ge)
		}
		if _caa, _afbb := _fcgb.(_af.EndElement); _afbb && _caa.Name == start.Name {
			break
		}
	}
	return nil
}

func (_afb *Box) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_bd, _ag := d.Token()
		if _ag != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0042\u006f\u0078\u003a\u0020\u0025\u0073", _ag)
		}
		if _bf, _bfc := _bd.(_af.EndElement); _bfc && _bf.Name == start.Name {
			break
		}
	}
	return nil
}

func (_cgf *Period) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_dgf, _dfd := d.Token()
		if _dfd != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0050e\u0072\u0069o\u0064\u003a\u0020\u0025\u0073", _dfd)
		}
		if _ccgc, _aefb := _dgf.(_af.EndElement); _aefb && _ccgc.Name == start.Name {
			break
		}
	}
	return nil
}

func NewMESH() *MESH { _afg := &MESH{}; return _afg }

// Validate validates the LCSH and its children
func (_egfb *LCSH) Validate() error { return _egfb.ValidateWithPath("\u004c\u0043\u0053\u0048") }

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_gfc *DDC) ValidateWithPath(path string) error { return nil }

type DCMIType struct{}

func NewLCSH() *LCSH { _bdc := &LCSH{}; return _bdc }

type IMT struct{}

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_gbaa *TGN) ValidateWithPath(path string) error { return nil }

func NewRFC3066() *RFC3066 { _gad := &RFC3066{}; return _gad }

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_abc *LCSH) ValidateWithPath(path string) error { return nil }

func (_gb *ElementOrRefinementContainer) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
_dg:
	for {
		_df, _ee := d.Token()
		if _ee != nil {
			return _ee
		}
		switch _cda := _df.(type) {
		case _af.StartElement:
			switch _cda.Name {
			case _af.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f", Local: "\u0061\u006e\u0079"}:
				_fc := NewElementsAndRefinementsGroupChoice()
				if _ebb := d.DecodeElement(&_fc.Any, &_cda); _ebb != nil {
					return _ebb
				}
				_gb.Choice = append(_gb.Choice, _fc)
			default:
				_e.Log.Debug("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020E\u006c\u0065\u006d\u0065\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065n\u0074\u0043on\u0074\u0061\u0069n\u0065\u0072\u0020\u0025\u0076", _cda.Name)
				if _gbc := d.Skip(); _gbc != nil {
					return _gbc
				}
			}
		case _af.EndElement:
			break _dg
		case _af.CharData:
		}
	}
	return nil
}

// Validate validates the W3CDTF and its children
func (_gadc *W3CDTF) Validate() error {
	return _gadc.ValidateWithPath("\u0057\u0033\u0043\u0044\u0054\u0046")
}

func NewElementOrRefinementContainer() *ElementOrRefinementContainer {
	_eab := &ElementOrRefinementContainer{}
	return _eab
}

func (_gagb *URI) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0055\u0052\u0049"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func NewRFC1766() *RFC1766 { _caga := &RFC1766{}; return _caga }

func NewDDC() *DDC { _ebg := &DDC{}; return _ebg }

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_cgc *RFC1766) ValidateWithPath(path string) error { return nil }

func (_fffg *UDC) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_ggba, _gec := d.Token()
		if _gec != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0044\u0043\u003a\u0020\u0025\u0073", _gec)
		}
		if _bfgf, _ecgg := _ggba.(_af.EndElement); _ecgg && _bfgf.Name == start.Name {
			break
		}
	}
	return nil
}

func (_aefa *Period) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0050\u0065\u0072\u0069\u006f\u0064"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func (_ae *ElementsAndRefinementsGroupChoice) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	if _ae.Any != nil {
		_ccg := _af.StartElement{Name: _af.Name{Local: "\u0064\u0063\u003a\u0061\u006e\u0079"}}
		for _, _egg := range _ae.Any {
			e.EncodeElement(_egg, _ccg)
		}
	}
	return nil
}

func NewIMT() *IMT { _aaf := &IMT{}; return _aaf }

func NewURI() *URI { _gdf := &URI{}; return _gdf }

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_ce *DCMIType) ValidateWithPath(path string) error { return nil }

func (_fcg *ISO3166) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0049S\u004f\u0033\u0031\u0036\u0036"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

type W3CDTF struct{}

func (_ecgf *RFC1766) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_aad, _gca := d.Token()
		if _gca != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0031\u0037\u0036\u0036\u003a\u0020\u0025\u0073", _gca)
		}
		if _bfab, _gage := _aad.(_af.EndElement); _gage && _bfab.Name == start.Name {
			break
		}
	}
	return nil
}

func NewISO3166() *ISO3166 { _bfg := &ISO3166{}; return _bfg }

func (_cce *ElementsAndRefinementsGroup) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	if _cce.Choice != nil {
		for _, _aa := range _cce.Choice {
			_aa.MarshalXML(e, _af.StartElement{})
		}
	}
	return nil
}

func (_gagg *W3CDTF) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_cfd, _cbd := d.Token()
		if _cbd != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u00573\u0043\u0044T\u0046\u003a\u0020\u0025\u0073", _cbd)
		}
		if _gdg, _bef := _cfd.(_af.EndElement); _bef && _gdg.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_cgb *Point) ValidateWithPath(path string) error { return nil }

func (_ed *MESH) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_dea, _dabg := d.Token()
		if _dabg != nil {
			return _c.Errorf("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004d\u0045\u0053\u0048\u003a\u0020\u0025\u0073", _dabg)
		}
		if _fab, _fg := _dea.(_af.EndElement); _fg && _fab.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the ElementsAndRefinementsGroup and its children
func (_eeg *ElementsAndRefinementsGroup) Validate() error {
	return _eeg.ValidateWithPath("E\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070")
}

func (_bgf *TGN) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0054\u0047\u004e"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func (_agd *Point) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_fgc, _faa := d.Token()
		if _faa != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0050\u006f\u0069\u006et\u003a\u0020\u0025\u0073", _faa)
		}
		if _efc, _dgge := _fgc.(_af.EndElement); _dgge && _efc.Name == start.Name {
			break
		}
	}
	return nil
}

type Box struct{}

func (_bcg *DCMIType) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_g, _cg := d.Token()
		if _cg != nil {
			return _c.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0044\u0043\u004dI\u0054\u0079\u0070\u0065: \u0025\u0073", _cg)
		}
		if _ac, _ff := _g.(_af.EndElement); _ff && _ac.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_ege *ISO3166) ValidateWithPath(path string) error { return nil }

func NewBox() *Box { _b := &Box{}; return _b }

func (_age *RFC3066) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0052F\u0043\u0033\u0030\u0036\u0036"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func (_d *DCMIType) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

// Validate validates the TGN and its children
func (_bac *TGN) Validate() error { return _bac.ValidateWithPath("\u0054\u0047\u004e") }

func (_cee *RFC3066) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_ggb, _bge := d.Token()
		if _bge != nil {
			return _c.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0033\u0030\u0036\u0036\u003a\u0020\u0025\u0073", _bge)
		}
		if _cba, _bdcb := _ggb.(_af.EndElement); _bdcb && _cba.Name == start.Name {
			break
		}
	}
	return nil
}

type ElementOrRefinementContainer struct {
	Choice []*ElementsAndRefinementsGroupChoice
}

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_db *ElementOrRefinementContainer) ValidateWithPath(path string) error {
	for _fd, _dbc := range _db.Choice {
		if _gg := _dbc.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _fd)); _gg != nil {
			return _gg
		}
	}
	return nil
}

func NewISO639_2() *ISO639_2 { _eafce := &ISO639_2{}; return _eafce }

func (_fffd *LCC) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_fbc, _efe := d.Token()
		if _efe != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u004c\u0043\u0043\u003a\u0020\u0025\u0073", _efe)
		}
		if _eeb, _dee := _fbc.(_af.EndElement); _dee && _eeb.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the IMT and its children
func (_gaa *IMT) Validate() error { return _gaa.ValidateWithPath("\u0049\u004d\u0054") }

// Validate validates the LCC and its children
func (_dbd *LCC) Validate() error { return _dbd.ValidateWithPath("\u004c\u0043\u0043") }

func (_gdc *MESH) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u004d\u0045\u0053\u0048"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func (_ecg *IMT) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_cbf, _dff := d.Token()
		if _dff != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0049\u004d\u0054\u003a\u0020\u0025\u0073", _dff)
		}
		if _ece, _gc := _cbf.(_af.EndElement); _gc && _ece.Name == start.Name {
			break
		}
	}
	return nil
}

func NewElementsAndRefinementsGroupChoice() *ElementsAndRefinementsGroupChoice {
	_egf := &ElementsAndRefinementsGroupChoice{}
	return _egf
}

func (_bc *Box) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0042\u006f\u0078"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func (_ea *DDC) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0044\u0044\u0043"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func (_cae *W3CDTF) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0057\u0033\u0043\u0044\u0054\u0046"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func (_cbff *UDC) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0055\u0044\u0043"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

type RFC3066 struct{}

// Validate validates the Box and its children
func (_ca *Box) Validate() error { return _ca.ValidateWithPath("\u0042\u006f\u0078") }

func (_bga *RFC1766) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0052F\u0043\u0031\u0037\u0036\u0036"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func NewW3CDTF() *W3CDTF { _aee := &W3CDTF{}; return _aee }

// Validate validates the RFC3066 and its children
func (_aab *RFC3066) Validate() error {
	return _aab.ValidateWithPath("\u0052F\u0043\u0033\u0030\u0036\u0036")
}

func (_efd *URI) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_gaag, _dgc := d.Token()
		if _dgc != nil {
			return _c.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0052\u0049\u003a\u0020\u0025\u0073", _dgc)
		}
		if _ccad, _efcf := _gaag.(_af.EndElement); _efcf && _ccad.Name == start.Name {
			break
		}
	}
	return nil
}

type ElementsAndRefinementsGroup struct {
	Choice []*ElementsAndRefinementsGroupChoice
}

type TGN struct{}

type UDC struct{}

func NewTGN() *TGN { _cca := &TGN{}; return _cca }

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_dfb *UDC) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_ceg *RFC3066) ValidateWithPath(path string) error { return nil }

type URI struct{}

type MESH struct{}

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_ab *ElementsAndRefinementsGroup) ValidateWithPath(path string) error {
	for _dd, _cga := range _ab.Choice {
		if _cbb := _cga.ValidateWithPath(_c.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _dd)); _cbb != nil {
			return _cbb
		}
	}
	return nil
}

// Validate validates the MESH and its children
func (_bgb *MESH) Validate() error { return _bgb.ValidateWithPath("\u004d\u0045\u0053\u0048") }

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_cbfg *IMT) ValidateWithPath(path string) error { return nil }

type ISO639_2 struct{}

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_fba *Period) ValidateWithPath(path string) error { return nil }

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_ec *ElementsAndRefinementsGroupChoice) Validate() error {
	return _ec.ValidateWithPath("\u0045\u006c\u0065\u006d\u0065\u006et\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0043h\u006f\u0069\u0063\u0065")
}

func NewPoint() *Point { _beb := &Point{}; return _beb }

// Validate validates the UDC and its children
func (_bed *UDC) Validate() error { return _bed.ValidateWithPath("\u0055\u0044\u0043") }

// Validate validates the DDC and its children
func (_ccd *DDC) Validate() error { return _ccd.ValidateWithPath("\u0044\u0044\u0043") }

func (_geg *LCSH) UnmarshalXML(d *_af.Decoder, start _af.StartElement) error {
	for {
		_be, _cfa := d.Token()
		if _cfa != nil {
			return _c.Errorf("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004c\u0043\u0053\u0048\u003a\u0020\u0025\u0073", _cfa)
		}
		if _cdd, _bbc := _be.(_af.EndElement); _bbc && _cdd.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_bfb *W3CDTF) ValidateWithPath(path string) error { return nil }

// Validate validates the ISO3166 and its children
func (_acf *ISO3166) Validate() error {
	return _acf.ValidateWithPath("\u0049S\u004f\u0033\u0031\u0036\u0036")
}

// Validate validates the ElementOrRefinementContainer and its children
func (_dgb *ElementOrRefinementContainer) Validate() error {
	return _dgb.ValidateWithPath("\u0045\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072")
}

func (_bbg *LCC) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u004c\u0043\u0043"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

func NewPeriod() *Period { _bae := &Period{}; return _bae }

// Validate validates the RFC1766 and its children
func (_fbb *RFC1766) Validate() error {
	return _fbb.ValidateWithPath("\u0052F\u0043\u0031\u0037\u0036\u0036")
}

// Validate validates the DCMIType and its children
func (_gd *DCMIType) Validate() error {
	return _gd.ValidateWithPath("\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065")
}

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_eae *URI) ValidateWithPath(path string) error { return nil }

// Validate validates the URI and its children
func (_eag *URI) Validate() error { return _eag.ValidateWithPath("\u0055\u0052\u0049") }

func (_cged *IMT) MarshalXML(e *_af.Encoder, start _af.StartElement) error {
	start.Name.Local = "\u0049\u004d\u0054"
	e.EncodeToken(start)
	e.EncodeToken(_af.EndElement{Name: start.Name})
	return nil
}

type LCC struct{}

func init() {
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u004c\u0043\u0053\u0048", NewLCSH)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u004d\u0045\u0053\u0048", NewMESH)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0044\u0044\u0043", NewDDC)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u004c\u0043\u0043", NewLCC)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0055\u0044\u0043", NewUDC)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0050\u0065\u0072\u0069\u006f\u0064", NewPeriod)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0057\u0033\u0043\u0044\u0054\u0046", NewW3CDTF)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065", NewDCMIType)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0049\u004d\u0054", NewIMT)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0055\u0052\u0049", NewURI)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032", NewISO639_2)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0052F\u0043\u0031\u0037\u0036\u0036", NewRFC1766)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0052F\u0043\u0033\u0030\u0036\u0036", NewRFC3066)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0050\u006f\u0069n\u0074", NewPoint)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0049S\u004f\u0033\u0031\u0036\u0036", NewISO3166)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0042\u006f\u0078", NewBox)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0054\u0047\u004e", NewTGN)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072", NewElementOrRefinementContainer)
	_f.RegisterConstructor("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/", "e\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070", NewElementsAndRefinementsGroup)
}
