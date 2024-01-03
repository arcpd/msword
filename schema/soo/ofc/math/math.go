package math

import (
	_g "encoding/xml"
	_d "fmt"
	_be "strconv"

	_c "github.com/arcpd/msword"
	_f "github.com/arcpd/msword/common/logger"
	_dc "github.com/arcpd/msword/schema/soo/ofc/sharedTypes"
)

// Validate validates the CT_D and its children
func (_bcf *CT_D) Validate() error { return _bcf.ValidateWithPath("\u0043\u0054\u005f\u0044") }

func (_fcac *CT_SPrePr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_bdfeg:
	for {
		_cgcf, _eacf := d.Token()
		if _eacf != nil {
			return _eacf
		}
		switch _bgbbc := _cgcf.(type) {
		case _g.StartElement:
			switch _bgbbc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_fcac.CtrlPr = NewCT_CtrlPr()
				if _bcbc := d.DecodeElement(_fcac.CtrlPr, &_bgbbc); _bcbc != nil {
					return _bcbc
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0053\u0050\u0072e\u0050\u0072 \u0025\u0076", _bgbbc.Name)
				if _bcbd := d.Skip(); _bcbd != nil {
					return _bcbd
				}
			}
		case _g.EndElement:
			break _bdfeg
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OMath and its children
func (_dcac *CT_OMath) Validate() error {
	return _dcac.ValidateWithPath("\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068")
}

type CT_OMathPara struct {
	OMathParaPr *CT_OMathParaPr
	OMath       []*CT_OMath
}

func (_bafec ST_Style) Validate() error { return _bafec.ValidateWithPath("") }

func (_febb *CT_LimUpp) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _febb.LimUppPr != nil {
		_bcb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u0055\u0070\u0070\u0050\u0072"}}
		e.EncodeElement(_febb.LimUppPr, _bcb)
	}
	_adff := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_febb.E, _adff)
	_adfc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006ci\u006d"}}
	e.EncodeElement(_febb.Lim, _adfc)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_gecg *CT_Rad) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_gecg.Deg = NewCT_OMathArg()
	_gecg.E = NewCT_OMathArg()
_bfebb:
	for {
		_gbafe, _dabg := d.Token()
		if _dabg != nil {
			return _dabg
		}
		switch _bedb := _gbafe.(type) {
		case _g.StartElement:
			switch _bedb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064P\u0072"}:
				_gecg.RadPr = NewCT_RadPr()
				if _ebbg := d.DecodeElement(_gecg.RadPr, &_bedb); _ebbg != nil {
					return _ebbg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0065\u0067"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0065\u0067"}:
				if _babe := d.DecodeElement(_gecg.Deg, &_bedb); _babe != nil {
					return _babe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _aeaa := d.DecodeElement(_gecg.E, &_bedb); _aeaa != nil {
					return _aeaa
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0061\u0064\u0020\u0025\u0076", _bedb.Name)
				if _egffc := d.Skip(); _egffc != nil {
					return _egffc
				}
			}
		case _g.EndElement:
			break _bfebb
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SPre and its children
func (_fbec *CT_SPre) Validate() error {
	return _fbec.ValidateWithPath("\u0043T\u005f\u0053\u0050\u0072\u0065")
}

func (_cdba *CT_FPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cdba.Type != nil {
		_aga := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0074\u0079\u0070\u0065"}}
		e.EncodeElement(_cdba.Type, _aga)
	}
	if _cdba.CtrlPr != nil {
		_fca := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_cdba.CtrlPr, _fca)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_dgeccd ST_Jc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_dgeccd.String(), start)
}

func NewCT_BorderBoxPr() *CT_BorderBoxPr { _ddfc := &CT_BorderBoxPr{}; return _ddfc }

func (_dbef *CT_LimLoc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_dbef.ValAttr = ST_LimLoc(1)
	for _, _dfec := range start.Attr {
		if _dfec.Name.Local == "\u0076\u0061\u006c" {
			_dbef.ValAttr.UnmarshalXMLAttr(_dfec)
			continue
		}
	}
	for {
		_eggg, _fcc := d.Token()
		if _fcc != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u004c\u0069\u006dL\u006f\u0063\u003a\u0020\u0025\u0073", _fcc)
		}
		if _fabe, _ffb := _eggg.(_g.EndElement); _ffb && _fabe.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_M and its children
func (_fadf *CT_M) Validate() error { return _fadf.ValidateWithPath("\u0043\u0054\u005f\u004d") }

type CT_R struct {
	RPr    *CT_RPR
	Choice []*CT_RChoice
}

// ValidateWithPath validates the CT_SSubSupPr and its children, prefixing error messages with path
func (_eccgc *CT_SSubSupPr) ValidateWithPath(path string) error {
	if _eccgc.AlnScr != nil {
		if _cdge := _eccgc.AlnScr.ValidateWithPath(path + "\u002fA\u006c\u006e\u0053\u0063\u0072"); _cdge != nil {
			return _cdge
		}
	}
	if _eccgc.CtrlPr != nil {
		if _fcbc := _eccgc.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _fcbc != nil {
			return _fcbc
		}
	}
	return nil
}

func (_dbdca *CT_RChoice) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gggc:
	for {
		_fagf, _gfea := d.Token()
		if _gfea != nil {
			return _gfea
		}
		switch _ggfb := _fagf.(type) {
		case _g.StartElement:
			switch _ggfb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0074"}:
				_dgbae := NewCT_Text()
				if _cada := d.DecodeElement(_dgbae, &_ggfb); _cada != nil {
					return _cada
				}
				_dbdca.T = append(_dbdca.T, _dgbae)
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fR\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _ggfb.Name)
				if _gdcg := d.Skip(); _gdcg != nil {
					return _gdcg
				}
			}
		case _g.EndElement:
			break _gggc
		case _g.CharData:
		}
	}
	return nil
}

func (_dabf ST_LimLoc) ValidateWithPath(path string) error {
	switch _dabf {
	case 0, 1, 2:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dabf))
	}
	return nil
}

type CT_String struct{ ValAttr *string }

type CT_EqArrPr struct {
	BaseJc  *CT_YAlign
	MaxDist *CT_OnOff
	ObjDist *CT_OnOff
	RSpRule *CT_SpacingRule
	RSp     *CT_UnSignedInteger
	CtrlPr  *CT_CtrlPr
}

// Validate validates the CT_XAlign and its children
func (_gbcfb *CT_XAlign) Validate() error {
	return _gbcfb.ValidateWithPath("\u0043T\u005f\u0058\u0041\u006c\u0069\u0067n")
}

func (_acfff ST_FType) Validate() error { return _acfff.ValidateWithPath("") }

// ValidateWithPath validates the EG_OMathMathElements and its children, prefixing error messages with path
func (_dgdgc *EG_OMathMathElements) ValidateWithPath(path string) error {
	if _dgdgc.Acc != nil {
		if _fbfad := _dgdgc.Acc.ValidateWithPath(path + "\u002f\u0041\u0063\u0063"); _fbfad != nil {
			return _fbfad
		}
	}
	if _dgdgc.Bar != nil {
		if _gcfge := _dgdgc.Bar.ValidateWithPath(path + "\u002f\u0042\u0061\u0072"); _gcfge != nil {
			return _gcfge
		}
	}
	if _dgdgc.Box != nil {
		if _aadae := _dgdgc.Box.ValidateWithPath(path + "\u002f\u0042\u006f\u0078"); _aadae != nil {
			return _aadae
		}
	}
	if _dgdgc.BorderBox != nil {
		if _efbfd := _dgdgc.BorderBox.ValidateWithPath(path + "\u002f\u0042\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078"); _efbfd != nil {
			return _efbfd
		}
	}
	if _dgdgc.D != nil {
		if _afdeg := _dgdgc.D.ValidateWithPath(path + "\u002f\u0044"); _afdeg != nil {
			return _afdeg
		}
	}
	if _dgdgc.EqArr != nil {
		if _gfegg := _dgdgc.EqArr.ValidateWithPath(path + "\u002f\u0045\u0071\u0041\u0072\u0072"); _gfegg != nil {
			return _gfegg
		}
	}
	if _dgdgc.F != nil {
		if _fefa := _dgdgc.F.ValidateWithPath(path + "\u002f\u0046"); _fefa != nil {
			return _fefa
		}
	}
	if _dgdgc.Func != nil {
		if _eecg := _dgdgc.Func.ValidateWithPath(path + "\u002f\u0046\u0075n\u0063"); _eecg != nil {
			return _eecg
		}
	}
	if _dgdgc.GroupChr != nil {
		if _ccfb := _dgdgc.GroupChr.ValidateWithPath(path + "\u002fG\u0072\u006f\u0075\u0070\u0043\u0068r"); _ccfb != nil {
			return _ccfb
		}
	}
	if _dgdgc.LimLow != nil {
		if _fdbea := _dgdgc.LimLow.ValidateWithPath(path + "\u002fL\u0069\u006d\u004c\u006f\u0077"); _fdbea != nil {
			return _fdbea
		}
	}
	if _dgdgc.LimUpp != nil {
		if _abab := _dgdgc.LimUpp.ValidateWithPath(path + "\u002fL\u0069\u006d\u0055\u0070\u0070"); _abab != nil {
			return _abab
		}
	}
	if _dgdgc.M != nil {
		if _bbebb := _dgdgc.M.ValidateWithPath(path + "\u002f\u004d"); _bbebb != nil {
			return _bbebb
		}
	}
	if _dgdgc.Nary != nil {
		if _dbefd := _dgdgc.Nary.ValidateWithPath(path + "\u002f\u004e\u0061r\u0079"); _dbefd != nil {
			return _dbefd
		}
	}
	if _dgdgc.Phant != nil {
		if _fegg := _dgdgc.Phant.ValidateWithPath(path + "\u002f\u0050\u0068\u0061\u006e\u0074"); _fegg != nil {
			return _fegg
		}
	}
	if _dgdgc.Rad != nil {
		if _efea := _dgdgc.Rad.ValidateWithPath(path + "\u002f\u0052\u0061\u0064"); _efea != nil {
			return _efea
		}
	}
	if _dgdgc.SPre != nil {
		if _adcgf := _dgdgc.SPre.ValidateWithPath(path + "\u002f\u0053\u0050r\u0065"); _adcgf != nil {
			return _adcgf
		}
	}
	if _dgdgc.SSub != nil {
		if _dcfae := _dgdgc.SSub.ValidateWithPath(path + "\u002f\u0053\u0053u\u0062"); _dcfae != nil {
			return _dcfae
		}
	}
	if _dgdgc.SSubSup != nil {
		if _fcff := _dgdgc.SSubSup.ValidateWithPath(path + "\u002f\u0053\u0053\u0075\u0062\u0053\u0075\u0070"); _fcff != nil {
			return _fcff
		}
	}
	if _dgdgc.SSup != nil {
		if _eaebe := _dgdgc.SSup.ValidateWithPath(path + "\u002f\u0053\u0053u\u0070"); _eaebe != nil {
			return _eaebe
		}
	}
	if _dgdgc.R != nil {
		if _defd := _dgdgc.R.ValidateWithPath(path + "\u002f\u0052"); _defd != nil {
			return _defd
		}
	}
	return nil
}

type CT_TwipsMeasure struct{ ValAttr _dc.ST_TwipsMeasure }

// ValidateWithPath validates the CT_Integer255 and its children, prefixing error messages with path
func (_egf *CT_Integer255) ValidateWithPath(path string) error {
	if _egf.ValAttr < 1 {
		return _d.Errorf("%\u0073\u002f\u006d\u002e\u0056\u0061l\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0031\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _egf.ValAttr)
	}
	if _egf.ValAttr > 255 {
		return _d.Errorf("\u0025\u0073/\u006d\u002e\u0056\u0061l\u0041\u0074t\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u0062e\u0020\u003c\u003d\u0020\u0032\u0035\u0035\u0020\u0028\u0068\u0061\u0076e\u0020\u0025\u0076\u0029", path, _egf.ValAttr)
	}
	return nil
}

type CT_Script struct{ ValAttr ST_Script }

func (_af *CT_Bar) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _af.BarPr != nil {
		_dbd := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0062\u0061\u0072\u0050\u0072"}}
		e.EncodeElement(_af.BarPr, _dbd)
	}
	_ecg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_af.E, _ecg)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_LimLoc struct{ ValAttr ST_LimLoc }

// Validate validates the CT_MathPrChoice and its children
func (_bbee *CT_MathPrChoice) Validate() error {
	return _bbee.ValidateWithPath("\u0043T\u005fM\u0061\u0074\u0068\u0050\u0072\u0043\u0068\u006f\u0069\u0063\u0065")
}

// Validate validates the CT_LimUppPr and its children
func (_adccf *CT_LimUppPr) Validate() error {
	return _adccf.ValidateWithPath("C\u0054\u005f\u004c\u0069\u006d\u0055\u0070\u0070\u0050\u0072")
}

// ValidateWithPath validates the CT_DPr and its children, prefixing error messages with path
func (_bfbe *CT_DPr) ValidateWithPath(path string) error {
	if _bfbe.BegChr != nil {
		if _efaad := _bfbe.BegChr.ValidateWithPath(path + "\u002fB\u0065\u0067\u0043\u0068\u0072"); _efaad != nil {
			return _efaad
		}
	}
	if _bfbe.SepChr != nil {
		if _bafe := _bfbe.SepChr.ValidateWithPath(path + "\u002fS\u0065\u0070\u0043\u0068\u0072"); _bafe != nil {
			return _bafe
		}
	}
	if _bfbe.EndChr != nil {
		if _cad := _bfbe.EndChr.ValidateWithPath(path + "\u002fE\u006e\u0064\u0043\u0068\u0072"); _cad != nil {
			return _cad
		}
	}
	if _bfbe.Grow != nil {
		if _fcg := _bfbe.Grow.ValidateWithPath(path + "\u002f\u0047\u0072o\u0077"); _fcg != nil {
			return _fcg
		}
	}
	if _bfbe.Shp != nil {
		if _eea := _bfbe.Shp.ValidateWithPath(path + "\u002f\u0053\u0068\u0070"); _eea != nil {
			return _eea
		}
	}
	if _bfbe.CtrlPr != nil {
		if _eee := _bfbe.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _eee != nil {
			return _eee
		}
	}
	return nil
}

func (_ffga *CT_Shp) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ffga.ValAttr = ST_Shp(1)
	for _, _febf := range start.Attr {
		if _febf.Name.Local == "\u0076\u0061\u006c" {
			_ffga.ValAttr.UnmarshalXMLAttr(_febf)
			continue
		}
	}
	for {
		_fdef, _aefgf := d.Token()
		if _aefgf != nil {
			return _d.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0053h\u0070\u003a\u0020\u0025\u0073", _aefgf)
		}
		if _cdbbc, _dccg := _fdef.(_g.EndElement); _dccg && _cdbbc.Name == start.Name {
			break
		}
	}
	return nil
}

type CT_Rad struct {
	RadPr *CT_RadPr
	Deg   *CT_OMathArg
	E     *CT_OMathArg
}

func (_fgeg *CT_GroupChr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fgeg.GroupChrPr != nil {
		_bdc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0067r\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072"}}
		e.EncodeElement(_fgeg.GroupChrPr, _bdc)
	}
	_bbc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_fgeg.E, _bbc)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_GroupChr() *CT_GroupChr { _cfae := &CT_GroupChr{}; _cfae.E = NewCT_OMathArg(); return _cfae }

// ValidateWithPath validates the CT_R and its children, prefixing error messages with path
func (_ffca *CT_R) ValidateWithPath(path string) error {
	if _ffca.RPr != nil {
		if _bcbb := _ffca.RPr.ValidateWithPath(path + "\u002f\u0052\u0050\u0072"); _bcbb != nil {
			return _bcbb
		}
	}
	for _fcaf, _adea := range _ffca.Choice {
		if _ebcg := _adea.ValidateWithPath(_d.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _fcaf)); _ebcg != nil {
			return _ebcg
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Script and its children, prefixing error messages with path
func (_aegcg *CT_Script) ValidateWithPath(path string) error {
	if _fecc := _aegcg.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _fecc != nil {
		return _fecc
	}
	return nil
}

// ValidateWithPath validates the CT_GroupChrPr and its children, prefixing error messages with path
func (_bfe *CT_GroupChrPr) ValidateWithPath(path string) error {
	if _bfe.Chr != nil {
		if _agdd := _bfe.Chr.ValidateWithPath(path + "\u002f\u0043\u0068\u0072"); _agdd != nil {
			return _agdd
		}
	}
	if _bfe.Pos != nil {
		if _gged := _bfe.Pos.ValidateWithPath(path + "\u002f\u0050\u006f\u0073"); _gged != nil {
			return _gged
		}
	}
	if _bfe.VertJc != nil {
		if _feg := _bfe.VertJc.ValidateWithPath(path + "\u002fV\u0065\u0072\u0074\u004a\u0063"); _feg != nil {
			return _feg
		}
	}
	if _bfe.CtrlPr != nil {
		if _caea := _bfe.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _caea != nil {
			return _caea
		}
	}
	return nil
}

func (_cbdc *CT_SSub) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cbdc.SSubPr != nil {
		_fbcb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0053\u0075\u0062\u0050\u0072"}}
		e.EncodeElement(_cbdc.SSubPr, _fbcb)
	}
	_efgd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_cbdc.E, _efgd)
	_dgce := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0062"}}
	e.EncodeElement(_cbdc.Sub, _dgce)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_OMathArgPr and its children
func (_ddaa *CT_OMathArgPr) Validate() error {
	return _ddaa.ValidateWithPath("\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067\u0050\u0072")
}

// Validate validates the CT_BreakBin and its children
func (_fcd *CT_BreakBin) Validate() error {
	return _fcd.ValidateWithPath("C\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042\u0069\u006e")
}

// ValidateWithPath validates the CT_Text and its children, prefixing error messages with path
func (_dabd *CT_Text) ValidateWithPath(path string) error { return nil }

type CT_MPr struct {
	BaseJc  *CT_YAlign
	PlcHide *CT_OnOff
	RSpRule *CT_SpacingRule
	CGpRule *CT_SpacingRule
	RSp     *CT_UnSignedInteger
	CSp     *CT_UnSignedInteger
	CGp     *CT_UnSignedInteger
	Mcs     *CT_MCS
	CtrlPr  *CT_CtrlPr
}

func (_fe *CT_AccPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ec:
	for {
		_dd, _dcg := d.Token()
		if _dcg != nil {
			return _dcg
		}
		switch _ef := _dd.(type) {
		case _g.StartElement:
			switch _ef.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0068\u0072"}:
				_fe.Chr = NewCT_Char()
				if _gdd := d.DecodeElement(_fe.Chr, &_ef); _gdd != nil {
					return _gdd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_fe.CtrlPr = NewCT_CtrlPr()
				if _dba := d.DecodeElement(_fe.CtrlPr, &_ef); _dba != nil {
					return _dba
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0041\u0063\u0063\u0050\u0072\u0020\u0025\u0076", _ef.Name)
				if _dbc := d.Skip(); _dbc != nil {
					return _dbc
				}
			}
		case _g.EndElement:
			break _ec
		case _g.CharData:
		}
	}
	return nil
}

func (_ccac *CT_SSubSupPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ccac.AlnScr != nil {
		_gegeg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0061\u006c\u006e\u0053\u0063\u0072"}}
		e.EncodeElement(_ccac.AlnScr, _gegeg)
	}
	if _ccac.CtrlPr != nil {
		_gedg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_ccac.CtrlPr, _gedg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_acg *CT_CtrlPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for {
		_facc, _abc := d.Token()
		if _abc != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0043\u0074\u0072l\u0050\u0072\u003a\u0020\u0025\u0073", _abc)
		}
		if _ddg, _bfc := _facc.(_g.EndElement); _bfc && _ddg.Name == start.Name {
			break
		}
	}
	return nil
}

func NewCT_GroupChrPr() *CT_GroupChrPr { _agdf := &CT_GroupChrPr{}; return _agdf }

func NewCT_YAlign() *CT_YAlign {
	_cdaf := &CT_YAlign{}
	_cdaf.ValAttr = _dc.ST_YAlign(1)
	return _cdaf
}

func (_dcc *CT_Acc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dcc.AccPr != nil {
		_ce := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0061\u0063\u0063\u0050\u0072"}}
		e.EncodeElement(_dcc.AccPr, _ce)
	}
	_db := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_dcc.E, _db)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Shp and its children, prefixing error messages with path
func (_gbffe *CT_Shp) ValidateWithPath(path string) error {
	if _gbffe.ValAttr == ST_ShpUnset {
		return _d.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _gdcd := _gbffe.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _gdcd != nil {
		return _gdcd
	}
	return nil
}

type CT_BarPr struct {
	Pos    *CT_TopBot
	CtrlPr *CT_CtrlPr
}

func (_afgdb ST_TopBot) Validate() error { return _afgdb.ValidateWithPath("") }

// Validate validates the CT_EqArr and its children
func (_gbf *CT_EqArr) Validate() error {
	return _gbf.ValidateWithPath("\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072")
}

func (_adae *CT_R) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _adae.RPr != nil {
		_bccd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072P\u0072"}}
		e.EncodeElement(_adae.RPr, _bccd)
	}
	if _adae.Choice != nil {
		for _, _adab := range _adae.Choice {
			_adab.MarshalXML(e, _g.StartElement{})
		}
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_BreakBinSub struct{ ValAttr ST_BreakBinSub }

func NewCT_Style() *CT_Style { _decea := &CT_Style{}; return _decea }

type CT_Bar struct {
	BarPr *CT_BarPr
	E     *CT_OMathArg
}

func NewCT_CtrlPr() *CT_CtrlPr { _ada := &CT_CtrlPr{}; return _ada }

func (_ecbbb *ST_Shp) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_ecbbb = 0
	case "\u0063\u0065\u006e\u0074\u0065\u0072\u0065\u0064":
		*_ecbbb = 1
	case "\u006d\u0061\u0074c\u0068":
		*_ecbbb = 2
	}
	return nil
}

func (_abdf *CT_PhantPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fggc:
	for {
		_ecbc, _acdd := d.Token()
		if _acdd != nil {
			return _acdd
		}
		switch _gbcfe := _ecbc.(type) {
		case _g.StartElement:
			switch _gbcfe.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0068\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0068\u006f\u0077"}:
				_abdf.Show = NewCT_OnOff()
				if _dbgc := d.DecodeElement(_abdf.Show, &_gbcfe); _dbgc != nil {
					return _dbgc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u007ae\u0072\u006f\u0057\u0069\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u007ae\u0072\u006f\u0057\u0069\u0064"}:
				_abdf.ZeroWid = NewCT_OnOff()
				if _gdggb := d.DecodeElement(_abdf.ZeroWid, &_gbcfe); _gdggb != nil {
					return _gdggb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u007ae\u0072\u006f\u0041\u0073\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u007ae\u0072\u006f\u0041\u0073\u0063"}:
				_abdf.ZeroAsc = NewCT_OnOff()
				if _dead := d.DecodeElement(_abdf.ZeroAsc, &_gbcfe); _dead != nil {
					return _dead
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u007a\u0065\u0072\u006f\u0044\u0065\u0073\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u007a\u0065\u0072\u006f\u0044\u0065\u0073\u0063"}:
				_abdf.ZeroDesc = NewCT_OnOff()
				if _agef := d.DecodeElement(_abdf.ZeroDesc, &_gbcfe); _agef != nil {
					return _agef
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0074\u0072\u0061\u006e\u0073\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0074\u0072\u0061\u006e\u0073\u0070"}:
				_abdf.Transp = NewCT_OnOff()
				if _eaea := d.DecodeElement(_abdf.Transp, &_gbcfe); _eaea != nil {
					return _eaea
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_abdf.CtrlPr = NewCT_CtrlPr()
				if _fbaec := d.DecodeElement(_abdf.CtrlPr, &_gbcfe); _fbaec != nil {
					return _fbaec
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0068\u0061\u006e\u0074\u0050\u0072\u0020\u0025\u0076", _gbcfe.Name)
				if _cfecb := d.Skip(); _cfecb != nil {
					return _cfecb
				}
			}
		case _g.EndElement:
			break _fggc
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Text and its children
func (_egcfg *CT_Text) Validate() error {
	return _egcfg.ValidateWithPath("\u0043T\u005f\u0054\u0065\u0078\u0074")
}

// ValidateWithPath validates the CT_LimLowPr and its children, prefixing error messages with path
func (_cfbf *CT_LimLowPr) ValidateWithPath(path string) error {
	if _cfbf.CtrlPr != nil {
		if _ffec := _cfbf.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _ffec != nil {
			return _ffec
		}
	}
	return nil
}

func (_feecd *CT_TwipsMeasure) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _agcd := range start.Attr {
		if _agcd.Name.Local == "\u0076\u0061\u006c" {
			_bafgd, _aebd := ParseUnionST_TwipsMeasure(_agcd.Value)
			if _aebd != nil {
				return _aebd
			}
			_feecd.ValAttr = _bafgd
			continue
		}
	}
	for {
		_ecbfa, _aadc := d.Token()
		if _aadc != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0043\u0054\u005f\u0054\u0077\u0069\u0070\u0073\u004de\u0061\u0073\u0075r\u0065:\u0020\u0025\u0073", _aadc)
		}
		if _edad, _gdce := _ecbfa.(_g.EndElement); _gdce && _edad.Name == start.Name {
			break
		}
	}
	return nil
}

const (
	ST_BreakBinSubUnset ST_BreakBinSub = 0
	ST_BreakBinSub__    ST_BreakBinSub = 1
	ST_BreakBinSub___   ST_BreakBinSub = 2
	ST_BreakBinSub____  ST_BreakBinSub = 3
)

func (_addf *CT_TopBot) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_addf.ValAttr = ST_TopBot(1)
	for _, _dgbb := range start.Attr {
		if _dgbb.Name.Local == "\u0076\u0061\u006c" {
			_addf.ValAttr.UnmarshalXMLAttr(_dgbb)
			continue
		}
	}
	for {
		_eaeae, _efgf := d.Token()
		if _efgf != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0054\u006f\u0070B\u006f\u0074\u003a\u0020\u0025\u0073", _efgf)
		}
		if _ggdcb, _ddca := _eaeae.(_g.EndElement); _ddca && _ggdcb.Name == start.Name {
			break
		}
	}
	return nil
}

func (_eaf *CT_DPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_faed:
	for {
		_gbc, _bdf := d.Token()
		if _bdf != nil {
			return _bdf
		}
		switch _gdf := _gbc.(type) {
		case _g.StartElement:
			switch _gdf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0065\u0067\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0065\u0067\u0043\u0068\u0072"}:
				_eaf.BegChr = NewCT_Char()
				if _eab := d.DecodeElement(_eaf.BegChr, &_gdf); _eab != nil {
					return _eab
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0065\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0065\u0070\u0043\u0068\u0072"}:
				_eaf.SepChr = NewCT_Char()
				if _dcag := d.DecodeElement(_eaf.SepChr, &_gdf); _dcag != nil {
					return _dcag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u006e\u0064\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u006e\u0064\u0043\u0068\u0072"}:
				_eaf.EndChr = NewCT_Char()
				if _geb := d.DecodeElement(_eaf.EndChr, &_gdf); _geb != nil {
					return _geb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0077"}:
				_eaf.Grow = NewCT_OnOff()
				if _fdbb := d.DecodeElement(_eaf.Grow, &_gdf); _fdbb != nil {
					return _fdbb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0068\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0068\u0070"}:
				_eaf.Shp = NewCT_Shp()
				if _daa := d.DecodeElement(_eaf.Shp, &_gdf); _daa != nil {
					return _daa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_eaf.CtrlPr = NewCT_CtrlPr()
				if _aeg := d.DecodeElement(_eaf.CtrlPr, &_gdf); _aeg != nil {
					return _aeg
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0044\u0050\u0072\u0020\u0025\u0076", _gdf.Name)
				if _aefb := d.Skip(); _aefb != nil {
					return _aefb
				}
			}
		case _g.EndElement:
			break _faed
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ScriptStyle and its children
func (_afaab *EG_ScriptStyle) Validate() error {
	return _afaab.ValidateWithPath("\u0045\u0047\u005f\u0053\u0063\u0072\u0069\u0070\u0074S\u0074\u0079\u006c\u0065")
}

func NewCT_BoxPr() *CT_BoxPr { _fdb := &CT_BoxPr{}; return _fdb }

// ValidateWithPath validates the CT_SPre and its children, prefixing error messages with path
func (_egfbg *CT_SPre) ValidateWithPath(path string) error {
	if _egfbg.SPrePr != nil {
		if _aaef := _egfbg.SPrePr.ValidateWithPath(path + "\u002fS\u0050\u0072\u0065\u0050\u0072"); _aaef != nil {
			return _aaef
		}
	}
	if _edbc := _egfbg.Sub.ValidateWithPath(path + "\u002f\u0053\u0075\u0062"); _edbc != nil {
		return _edbc
	}
	if _fgdd := _egfbg.Sup.ValidateWithPath(path + "\u002f\u0053\u0075\u0070"); _fgdd != nil {
		return _fgdd
	}
	if _eeea := _egfbg.E.ValidateWithPath(path + "\u002f\u0045"); _eeea != nil {
		return _eeea
	}
	return nil
}

type CT_Integer2 struct{ ValAttr int64 }

func NewEG_OMathElements() *EG_OMathElements { _ggedc := &EG_OMathElements{}; return _ggedc }

// ValidateWithPath validates the CT_LimUpp and its children, prefixing error messages with path
func (_gdgf *CT_LimUpp) ValidateWithPath(path string) error {
	if _gdgf.LimUppPr != nil {
		if _acc := _gdgf.LimUppPr.ValidateWithPath(path + "\u002fL\u0069\u006d\u0055\u0070\u0070\u0050r"); _acc != nil {
			return _acc
		}
	}
	if _gdfeb := _gdgf.E.ValidateWithPath(path + "\u002f\u0045"); _gdfeb != nil {
		return _gdfeb
	}
	if _aagd := _gdgf.Lim.ValidateWithPath(path + "\u002f\u004c\u0069\u006d"); _aagd != nil {
		return _aagd
	}
	return nil
}

func (_fgbe *CT_LimUppPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_faeg:
	for {
		_aefg, _ebeg := d.Token()
		if _ebeg != nil {
			return _ebeg
		}
		switch _cgeg := _aefg.(type) {
		case _g.StartElement:
			switch _cgeg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_fgbe.CtrlPr = NewCT_CtrlPr()
				if _bgc := d.DecodeElement(_fgbe.CtrlPr, &_cgeg); _bgc != nil {
					return _bgc
				}
			default:
				_f.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004c\u0069\u006d\u0055\u0070\u0070\u0050\u0072\u0020\u0025\u0076", _cgeg.Name)
				if _eddd := d.Skip(); _eddd != nil {
					return _eddd
				}
			}
		case _g.EndElement:
			break _faeg
		case _g.CharData:
		}
	}
	return nil
}

func (_dgdad ST_TopBot) ValidateWithPath(path string) error {
	switch _dgdad {
	case 0, 1, 2:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dgdad))
	}
	return nil
}

// ValidateWithPath validates the CT_SSubPr and its children, prefixing error messages with path
func (_dgbf *CT_SSubPr) ValidateWithPath(path string) error {
	if _dgbf.CtrlPr != nil {
		if _fdee := _dgbf.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _fdee != nil {
			return _fdee
		}
	}
	return nil
}

func (_cfdgb *CT_RPR) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_beeb:
	for {
		_gfdc, _ffcd := d.Token()
		if _ffcd != nil {
			return _ffcd
		}
		switch _ddcg := _gfdc.(type) {
		case _g.StartElement:
			switch _ddcg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u0074"}:
				_cfdgb.Lit = NewCT_OnOff()
				if _cbac := d.DecodeElement(_cfdgb.Lit, &_ddcg); _cbac != nil {
					return _cbac
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u006f\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u006f\u0072"}:
				_cfdgb.Choice = NewCT_RPRChoice()
				if _adfge := d.DecodeElement(&_cfdgb.Choice.Nor, &_ddcg); _adfge != nil {
					return _adfge
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0072\u006b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0072\u006b"}:
				_cfdgb.Brk = NewCT_ManualBreak()
				if _gdcf := d.DecodeElement(_cfdgb.Brk, &_ddcg); _gdcf != nil {
					return _gdcf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u006c\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u006c\u006e"}:
				_cfdgb.Aln = NewCT_OnOff()
				if _acff := d.DecodeElement(_cfdgb.Aln, &_ddcg); _acff != nil {
					return _acff
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0050\u0052\u0020\u0025\u0076", _ddcg.Name)
				if _fdfe := d.Skip(); _fdfe != nil {
					return _fdfe
				}
			}
		case _g.EndElement:
			break _beeb
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the OMathPara and its children, prefixing error messages with path
func (_ddabe *OMathPara) ValidateWithPath(path string) error {
	if _aee := _ddabe.CT_OMathPara.ValidateWithPath(path); _aee != nil {
		return _aee
	}
	return nil
}

// Validate validates the CT_AccPr and its children
func (_gef *CT_AccPr) Validate() error {
	return _gef.ValidateWithPath("\u0043\u0054\u005f\u0041\u0063\u0063\u0050\u0072")
}

type CT_SSupPr struct{ CtrlPr *CT_CtrlPr }

func (_ffab *CT_MPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ffab.BaseJc != nil {
		_cdbac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u0061\u0073\u0065\u004a\u0063"}}
		e.EncodeElement(_ffab.BaseJc, _cdbac)
	}
	if _ffab.PlcHide != nil {
		_gfa := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0070\u006c\u0063\u0048\u0069\u0064e"}}
		e.EncodeElement(_ffab.PlcHide, _gfa)
	}
	if _ffab.RSpRule != nil {
		_dgdg := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0072\u0053\u0070\u0052\u0075\u006ce"}}
		e.EncodeElement(_ffab.RSpRule, _dgdg)
	}
	if _ffab.CGpRule != nil {
		_cgf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0063\u0047\u0070\u0052\u0075\u006ce"}}
		e.EncodeElement(_ffab.CGpRule, _cgf)
	}
	if _ffab.RSp != nil {
		_gaga := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072S\u0070"}}
		e.EncodeElement(_ffab.RSp, _gaga)
	}
	if _ffab.CSp != nil {
		_begd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063S\u0070"}}
		e.EncodeElement(_ffab.CSp, _begd)
	}
	if _ffab.CGp != nil {
		_efbf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063G\u0070"}}
		e.EncodeElement(_ffab.CGp, _efbf)
	}
	if _ffab.Mcs != nil {
		_aagc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006dc\u0073"}}
		e.EncodeElement(_ffab.Mcs, _aagc)
	}
	if _ffab.CtrlPr != nil {
		_fccc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_ffab.CtrlPr, _fccc)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_SSubSupPr() *CT_SSubSupPr { _eadb := &CT_SSubSupPr{}; return _eadb }

// Validate validates the CT_OMathJc and its children
func (_caagd *CT_OMathJc) Validate() error {
	return _caagd.ValidateWithPath("\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u004a\u0063")
}

// Validate validates the CT_CtrlPr and its children
func (_adb *CT_CtrlPr) Validate() error {
	return _adb.ValidateWithPath("\u0043T\u005f\u0043\u0074\u0072\u006c\u0050r")
}

// ValidateWithPath validates the CT_AccPr and its children, prefixing error messages with path
func (_bg *CT_AccPr) ValidateWithPath(path string) error {
	if _bg.Chr != nil {
		if _da := _bg.Chr.ValidateWithPath(path + "\u002f\u0043\u0068\u0072"); _da != nil {
			return _da
		}
	}
	if _bg.CtrlPr != nil {
		if _gdg := _bg.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _gdg != nil {
			return _gdg
		}
	}
	return nil
}

// Validate validates the CT_RPRChoice and its children
func (_fagfb *CT_RPRChoice) Validate() error {
	return _fagfb.ValidateWithPath("\u0043\u0054\u005fR\u0050\u0052\u0043\u0068\u006f\u0069\u0063\u0065")
}

func (_cffe ST_Shp) String() string {
	switch _cffe {
	case 0:
		return ""
	case 1:
		return "\u0063\u0065\u006e\u0074\u0065\u0072\u0065\u0064"
	case 2:
		return "\u006d\u0061\u0074c\u0068"
	}
	return ""
}

func (_cc *CT_AccPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cc.Chr != nil {
		_bf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063h\u0072"}}
		e.EncodeElement(_cc.Chr, _bf)
	}
	if _cc.CtrlPr != nil {
		_gd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_cc.CtrlPr, _gd)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_GroupChrPr and its children
func (_gebf *CT_GroupChrPr) Validate() error {
	return _gebf.ValidateWithPath("\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072")
}

func (_geggb *ST_LimLoc) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_geggb = 0
	case "\u0075\u006e\u0064\u004f\u0076\u0072":
		*_geggb = 1
	case "\u0073\u0075\u0062\u0053\u0075\u0070":
		*_geggb = 2
	}
	return nil
}

// Validate validates the CT_TopBot and its children
func (_abdb *CT_TopBot) Validate() error {
	return _abdb.ValidateWithPath("\u0043T\u005f\u0054\u006f\u0070\u0042\u006ft")
}

func (_feda *CT_OMath) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ccec:
	for {
		_cbefb, _ggba := d.Token()
		if _ggba != nil {
			return _ggba
		}
		switch _eaef := _cbefb.(type) {
		case _g.StartElement:
			switch _eaef.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_bfeg := NewEG_OMathMathElements()
				_bfeg.Acc = NewCT_Acc()
				if _agafc := d.DecodeElement(_bfeg.Acc, &_eaef); _agafc != nil {
					return _agafc
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _bfeg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_bcda := NewEG_OMathMathElements()
				_bcda.Bar = NewCT_Bar()
				if _gae := d.DecodeElement(_bcda.Bar, &_eaef); _gae != nil {
					return _gae
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _bcda)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_afgda := NewEG_OMathMathElements()
				_afgda.Box = NewCT_Box()
				if _abfg := d.DecodeElement(_afgda.Box, &_eaef); _abfg != nil {
					return _abfg
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _afgda)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_dea := NewEG_OMathMathElements()
				_dea.BorderBox = NewCT_BorderBox()
				if _feed := d.DecodeElement(_dea.BorderBox, &_eaef); _feed != nil {
					return _feed
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _dea)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_caaf := NewEG_OMathMathElements()
				_caaf.D = NewCT_D()
				if _deba := d.DecodeElement(_caaf.D, &_eaef); _deba != nil {
					return _deba
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _caaf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_acca := NewEG_OMathMathElements()
				_acca.EqArr = NewCT_EqArr()
				if _dbbc := d.DecodeElement(_acca.EqArr, &_eaef); _dbbc != nil {
					return _dbbc
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _acca)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_efaeg := NewEG_OMathMathElements()
				_efaeg.F = NewCT_F()
				if _age := d.DecodeElement(_efaeg.F, &_eaef); _age != nil {
					return _age
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _efaeg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_fgcf := NewEG_OMathMathElements()
				_fgcf.Func = NewCT_Func()
				if _defb := d.DecodeElement(_fgcf.Func, &_eaef); _defb != nil {
					return _defb
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _fgcf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_gadd := NewEG_OMathMathElements()
				_gadd.GroupChr = NewCT_GroupChr()
				if _gfdeb := d.DecodeElement(_gadd.GroupChr, &_eaef); _gfdeb != nil {
					return _gfdeb
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _gadd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_gfdb := NewEG_OMathMathElements()
				_gfdb.LimLow = NewCT_LimLow()
				if _bcgb := d.DecodeElement(_gfdb.LimLow, &_eaef); _bcgb != nil {
					return _bcgb
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _gfdb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_eafb := NewEG_OMathMathElements()
				_eafb.LimUpp = NewCT_LimUpp()
				if _aegb := d.DecodeElement(_eafb.LimUpp, &_eaef); _aegb != nil {
					return _aegb
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _eafb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_edbbb := NewEG_OMathMathElements()
				_edbbb.M = NewCT_M()
				if _ecgf := d.DecodeElement(_edbbb.M, &_eaef); _ecgf != nil {
					return _ecgf
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _edbbb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_ecgd := NewEG_OMathMathElements()
				_ecgd.Nary = NewCT_Nary()
				if _ageb := d.DecodeElement(_ecgd.Nary, &_eaef); _ageb != nil {
					return _ageb
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _ecgd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_eed := NewEG_OMathMathElements()
				_eed.Phant = NewCT_Phant()
				if _ecgfd := d.DecodeElement(_eed.Phant, &_eaef); _ecgfd != nil {
					return _ecgfd
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _eed)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_gcdcb := NewEG_OMathMathElements()
				_gcdcb.Rad = NewCT_Rad()
				if _ebed := d.DecodeElement(_gcdcb.Rad, &_eaef); _ebed != nil {
					return _ebed
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _gcdcb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_babf := NewEG_OMathMathElements()
				_babf.SPre = NewCT_SPre()
				if _afdb := d.DecodeElement(_babf.SPre, &_eaef); _afdb != nil {
					return _afdb
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _babf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_gfca := NewEG_OMathMathElements()
				_gfca.SSub = NewCT_SSub()
				if _cggf := d.DecodeElement(_gfca.SSub, &_eaef); _cggf != nil {
					return _cggf
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _gfca)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_gbg := NewEG_OMathMathElements()
				_gbg.SSubSup = NewCT_SSubSup()
				if _gcda := d.DecodeElement(_gbg.SSubSup, &_eaef); _gcda != nil {
					return _gcda
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _gbg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_fcfc := NewEG_OMathMathElements()
				_fcfc.SSup = NewCT_SSup()
				if _daab := d.DecodeElement(_fcfc.SSup, &_eaef); _daab != nil {
					return _daab
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _fcfc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_ffac := NewEG_OMathMathElements()
				_ffac.R = NewCT_R()
				if _dgbe := d.DecodeElement(_ffac.R, &_eaef); _dgbe != nil {
					return _dgbe
				}
				_feda.EG_OMathMathElements = append(_feda.EG_OMathMathElements, _ffac)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0020\u0025\u0076", _eaef.Name)
				if _edgf := d.Skip(); _edgf != nil {
					return _edgf
				}
			}
		case _g.EndElement:
			break _ccec
		case _g.CharData:
		}
	}
	return nil
}

func (_bed *CT_BarPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _bed.Pos != nil {
		_dbe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0070o\u0073"}}
		e.EncodeElement(_bed.Pos, _dbe)
	}
	if _bed.CtrlPr != nil {
		_ggd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_bed.CtrlPr, _ggd)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_ManualBreak() *CT_ManualBreak { _egff := &CT_ManualBreak{}; return _egff }

// Validate validates the CT_LimLow and its children
func (_cged *CT_LimLow) Validate() error {
	return _cged.ValidateWithPath("\u0043T\u005f\u004c\u0069\u006d\u004c\u006fw")
}

func (_fab *CT_BorderBox) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fab.BorderBoxPr != nil {
		_gegf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078\u0050\u0072"}}
		e.EncodeElement(_fab.BorderBoxPr, _gegf)
	}
	_gc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_fab.E, _gc)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_dffbe *CT_SSubSup) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dffbe.SSubSupPr != nil {
		_edbbbf := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0073\u0053\u0075\u0062\u0053\u0075\u0070\u0050\u0072"}}
		e.EncodeElement(_dffbe.SSubSupPr, _edbbbf)
	}
	_cbf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_dffbe.E, _cbf)
	_gcfe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0062"}}
	e.EncodeElement(_dffbe.Sub, _gcfe)
	_dafe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0070"}}
	e.EncodeElement(_dffbe.Sup, _dafe)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the EG_ScriptStyle and its children, prefixing error messages with path
func (_bbed *EG_ScriptStyle) ValidateWithPath(path string) error {
	if _bbed.Scr != nil {
		if _ebdbf := _bbed.Scr.ValidateWithPath(path + "\u002f\u0053\u0063\u0072"); _ebdbf != nil {
			return _ebdbf
		}
	}
	if _bbed.Sty != nil {
		if _fafa := _bbed.Sty.ValidateWithPath(path + "\u002f\u0053\u0074\u0079"); _fafa != nil {
			return _fafa
		}
	}
	return nil
}

func (_gcf *CT_EqArr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gcf.EqArrPr != nil {
		_ccba := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0065\u0071\u0041\u0072\u0072\u0050r"}}
		e.EncodeElement(_gcf.EqArrPr, _ccba)
	}
	_gdgb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	for _, _cfag := range _gcf.E {
		e.EncodeElement(_cfag, _gdgb)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_cgafd ST_Style) ValidateWithPath(path string) error {
	switch _cgafd {
	case 0, 1, 2, 3, 4:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cgafd))
	}
	return nil
}

func NewCT_Shp() *CT_Shp { _cbdf := &CT_Shp{}; _cbdf.ValAttr = ST_Shp(1); return _cbdf }

type CT_Text struct {
	SpaceAttr *string
	Content   string
}

// ValidateWithPath validates the CT_OMath and its children, prefixing error messages with path
func (_bdfd *CT_OMath) ValidateWithPath(path string) error {
	for _bbcde, _ffc := range _bdfd.EG_OMathMathElements {
		if _gdb := _ffc.ValidateWithPath(_d.Sprintf("%\u0073\u002f\u0045\u0047\u005f\u004fM\u0061\u0074\u0068\u004d\u0061\u0074\u0068\u0045\u006ce\u006d\u0065\u006et\u0073[\u0025\u0064\u005d", path, _bbcde)); _gdb != nil {
			return _gdb
		}
	}
	return nil
}

func (_egb ST_BreakBinSub) String() string {
	switch _egb {
	case 0:
		return ""
	case 1:
		return "\u002d\u002d"
	case 2:
		return "\u002d\u002b"
	case 3:
		return "\u002b\u002d"
	}
	return ""
}

func (_aecga *EG_OMathMathElements) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _aecga.Acc != nil {
		_ccab := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0061c\u0063"}}
		e.EncodeElement(_aecga.Acc, _ccab)
	}
	if _aecga.Bar != nil {
		_cgfb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062a\u0072"}}
		e.EncodeElement(_aecga.Bar, _cgfb)
	}
	if _aecga.Box != nil {
		_dcbb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062o\u0078"}}
		e.EncodeElement(_aecga.Box, _dcbb)
	}
	if _aecga.BorderBox != nil {
		_dabc := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0062\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078"}}
		e.EncodeElement(_aecga.BorderBox, _dabc)
	}
	if _aecga.D != nil {
		_bdce := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064"}}
		e.EncodeElement(_aecga.D, _bdce)
	}
	if _aecga.EqArr != nil {
		_dffc := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0065\u0071\u0041\u0072\u0072"}}
		e.EncodeElement(_aecga.EqArr, _dffc)
	}
	if _aecga.F != nil {
		_bcag := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0066"}}
		e.EncodeElement(_aecga.F, _bcag)
	}
	if _aecga.Func != nil {
		_aeaba := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0066\u0075\u006e\u0063"}}
		e.EncodeElement(_aecga.Func, _aeaba)
	}
	if _aecga.GroupChr != nil {
		_fgec := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}}
		e.EncodeElement(_aecga.GroupChr, _fgec)
	}
	if _aecga.LimLow != nil {
		_dcfa := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u004c\u006f\u0077"}}
		e.EncodeElement(_aecga.LimLow, _dcfa)
	}
	if _aecga.LimUpp != nil {
		_dbff := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u0055\u0070\u0070"}}
		e.EncodeElement(_aecga.LimUpp, _dbff)
	}
	if _aecga.M != nil {
		_bddb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d"}}
		e.EncodeElement(_aecga.M, _bddb)
	}
	if _aecga.Nary != nil {
		_fgef := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006e\u0061\u0072\u0079"}}
		e.EncodeElement(_aecga.Nary, _fgef)
	}
	if _aecga.Phant != nil {
		_bbeb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0070\u0068\u0061\u006e\u0074"}}
		e.EncodeElement(_aecga.Phant, _bbeb)
	}
	if _aecga.Rad != nil {
		_edgc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072a\u0064"}}
		e.EncodeElement(_aecga.Rad, _edgc)
	}
	if _aecga.SPre != nil {
		_ebaea := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0050\u0072\u0065"}}
		e.EncodeElement(_aecga.SPre, _ebaea)
	}
	if _aecga.SSub != nil {
		_bbad := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0053\u0075\u0062"}}
		e.EncodeElement(_aecga.SSub, _bbad)
	}
	if _aecga.SSubSup != nil {
		_fdbbf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0053\u0075\u0062\u0053\u0075p"}}
		e.EncodeElement(_aecga.SSubSup, _fdbbf)
	}
	if _aecga.SSup != nil {
		_cagg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0053\u0075\u0070"}}
		e.EncodeElement(_aecga.SSup, _cagg)
	}
	if _aecga.R != nil {
		_fcde := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072"}}
		e.EncodeElement(_aecga.R, _fcde)
	}
	return nil
}

func (_cegba *ST_BreakBin) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cebdb, _gfcg := d.Token()
	if _gfcg != nil {
		return _gfcg
	}
	if _bbbg, _bbdd := _cebdb.(_g.EndElement); _bbdd && _bbbg.Name == start.Name {
		*_cegba = 1
		return nil
	}
	if _eege, _fdbbg := _cebdb.(_g.CharData); !_fdbbg {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cebdb)
	} else {
		switch string(_eege) {
		case "":
			*_cegba = 0
		case "\u0062\u0065\u0066\u006f\u0072\u0065":
			*_cegba = 1
		case "\u0061\u0066\u0074e\u0072":
			*_cegba = 2
		case "\u0072\u0065\u0070\u0065\u0061\u0074":
			*_cegba = 3
		}
	}
	_cebdb, _gfcg = d.Token()
	if _gfcg != nil {
		return _gfcg
	}
	if _ebfg, _fbbf := _cebdb.(_g.EndElement); _fbbf && _ebfg.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cebdb)
}

// ValidateWithPath validates the CT_BreakBinSub and its children, prefixing error messages with path
func (_aaed *CT_BreakBinSub) ValidateWithPath(path string) error {
	if _cff := _aaed.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _cff != nil {
		return _cff
	}
	return nil
}

// Validate validates the CT_BreakBinSub and its children
func (_ddc *CT_BreakBinSub) Validate() error {
	return _ddc.ValidateWithPath("\u0043\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042i\u006e\u0053\u0075\u0062")
}

func (_ebeee *CT_SSubSupPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fdcbd:
	for {
		_daee, _ccbg := d.Token()
		if _ccbg != nil {
			return _ccbg
		}
		switch _deea := _daee.(type) {
		case _g.StartElement:
			switch _deea.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u006c\u006e\u0053\u0063\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u006c\u006e\u0053\u0063\u0072"}:
				_ebeee.AlnScr = NewCT_OnOff()
				if _fggae := d.DecodeElement(_ebeee.AlnScr, &_deea); _fggae != nil {
					return _fggae
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_ebeee.CtrlPr = NewCT_CtrlPr()
				if _beada := d.DecodeElement(_ebeee.CtrlPr, &_deea); _beada != nil {
					return _beada
				}
			default:
				_f.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_S\u0053\u0075b\u0053\u0075\u0070\u0050\u0072\u0020\u0025\u0076", _deea.Name)
				if _ecgfg := d.Skip(); _ecgfg != nil {
					return _ecgfg
				}
			}
		case _g.EndElement:
			break _fdcbd
		case _g.CharData:
		}
	}
	return nil
}

func NewCT_LimLowPr() *CT_LimLowPr { _dafce := &CT_LimLowPr{}; return _dafce }

// Validate validates the CT_Char and its children
func (_agf *CT_Char) Validate() error {
	return _agf.ValidateWithPath("\u0043T\u005f\u0043\u0068\u0061\u0072")
}

func (_ffee *CT_SSubPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ebce:
	for {
		_dega, _abag := d.Token()
		if _abag != nil {
			return _abag
		}
		switch _ebac := _dega.(type) {
		case _g.StartElement:
			switch _ebac.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_ffee.CtrlPr = NewCT_CtrlPr()
				if _dedd := d.DecodeElement(_ffee.CtrlPr, &_ebac); _dedd != nil {
					return _dedd
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0053\u0053\u0075b\u0050\u0072 \u0025\u0076", _ebac.Name)
				if _bffb := d.Skip(); _bffb != nil {
					return _bffb
				}
			}
		case _g.EndElement:
			break _ebce
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the MathPr and its children
func (_gaeb *MathPr) Validate() error {
	return _gaeb.ValidateWithPath("\u004d\u0061\u0074\u0068\u0050\u0072")
}

// ValidateWithPath validates the CT_BarPr and its children, prefixing error messages with path
func (_edf *CT_BarPr) ValidateWithPath(path string) error {
	if _edf.Pos != nil {
		if _df := _edf.Pos.ValidateWithPath(path + "\u002f\u0050\u006f\u0073"); _df != nil {
			return _df
		}
	}
	if _edf.CtrlPr != nil {
		if _bfb := _edf.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _bfb != nil {
			return _bfb
		}
	}
	return nil
}

func NewCT_M() *CT_M { _cdbf := &CT_M{}; return _cdbf }

// ValidateWithPath validates the CT_D and its children, prefixing error messages with path
func (_debb *CT_D) ValidateWithPath(path string) error {
	if _debb.DPr != nil {
		if _bga := _debb.DPr.ValidateWithPath(path + "\u002f\u0044\u0050\u0072"); _bga != nil {
			return _bga
		}
	}
	for _fdgc, _bcg := range _debb.E {
		if _edb := _bcg.ValidateWithPath(_d.Sprintf("\u0025\u0073\u002f\u0045\u005b\u0025\u0064\u005d", path, _fdgc)); _edb != nil {
			return _edb
		}
	}
	return nil
}

func (_agfbf *CT_UnSignedInteger) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _d.Sprintf("\u0025\u0076", _agfbf.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_XAlign() *CT_XAlign {
	_bdee := &CT_XAlign{}
	_bdee.ValAttr = _dc.ST_XAlign(1)
	return _bdee
}

type CT_CtrlPr struct{}

func (_eaebc *CT_LimLowPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _eaebc.CtrlPr != nil {
		_bcfg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_eaebc.CtrlPr, _bcfg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_faccf *CT_OnOff) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _fggg := range start.Attr {
		if _fggg.Name.Local == "\u0076\u0061\u006c" {
			_dcgac, _egeb := ParseUnionST_OnOff(_fggg.Value)
			if _egeb != nil {
				return _egeb
			}
			_faccf.ValAttr = &_dcgac
			continue
		}
	}
	for {
		_abb, _cfgb := d.Token()
		if _cfgb != nil {
			return _d.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fO\u006e\u004f\u0066\u0066: \u0025\u0073", _cfgb)
		}
		if _ebg, _bdga := _abb.(_g.EndElement); _bdga && _ebg.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ceff *CT_FType) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_bafc, _cgcb := _ceff.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _cgcb != nil {
		return _cgcb
	}
	start.Attr = append(start.Attr, _bafc)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_MathPr and its children
func (_faa *CT_MathPr) Validate() error {
	return _faa.ValidateWithPath("\u0043T\u005f\u004d\u0061\u0074\u0068\u0050r")
}

func (_bab *CT_MCPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gada:
	for {
		_gbac, _daff := d.Token()
		if _daff != nil {
			return _daff
		}
		switch _fbf := _gbac.(type) {
		case _g.StartElement:
			switch _fbf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u006f\u0075n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u006f\u0075n\u0074"}:
				_bab.Count = NewCT_Integer255()
				if _ade := d.DecodeElement(_bab.Count, &_fbf); _ade != nil {
					return _ade
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0063\u004a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0063\u004a\u0063"}:
				_bab.McJc = NewCT_XAlign()
				if _adaf := d.DecodeElement(_bab.McJc, &_fbf); _adaf != nil {
					return _adaf
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u004d\u0043\u0050\u0072\u0020\u0025\u0076", _fbf.Name)
				if _dgdc := d.Skip(); _dgdc != nil {
					return _dgdc
				}
			}
		case _g.EndElement:
			break _gada
		case _g.CharData:
		}
	}
	return nil
}

func NewCT_R() *CT_R { _gaag := &CT_R{}; return _gaag }

// Validate validates the CT_ManualBreak and its children
func (_bbec *CT_ManualBreak) Validate() error {
	return _bbec.ValidateWithPath("\u0043\u0054\u005f\u004d\u0061\u006e\u0075\u0061\u006cB\u0072\u0065\u0061\u006b")
}

// ValidateWithPath validates the CT_NaryPr and its children, prefixing error messages with path
func (_egga *CT_NaryPr) ValidateWithPath(path string) error {
	if _egga.Chr != nil {
		if _dada := _egga.Chr.ValidateWithPath(path + "\u002f\u0043\u0068\u0072"); _dada != nil {
			return _dada
		}
	}
	if _egga.LimLoc != nil {
		if _cfda := _egga.LimLoc.ValidateWithPath(path + "\u002fL\u0069\u006d\u004c\u006f\u0063"); _cfda != nil {
			return _cfda
		}
	}
	if _egga.Grow != nil {
		if _fdeb := _egga.Grow.ValidateWithPath(path + "\u002f\u0047\u0072o\u0077"); _fdeb != nil {
			return _fdeb
		}
	}
	if _egga.SubHide != nil {
		if _gebb := _egga.SubHide.ValidateWithPath(path + "\u002f\u0053\u0075\u0062\u0048\u0069\u0064\u0065"); _gebb != nil {
			return _gebb
		}
	}
	if _egga.SupHide != nil {
		if _bece := _egga.SupHide.ValidateWithPath(path + "\u002f\u0053\u0075\u0070\u0048\u0069\u0064\u0065"); _bece != nil {
			return _bece
		}
	}
	if _egga.CtrlPr != nil {
		if _faf := _egga.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _faf != nil {
			return _faf
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Box and its children, prefixing error messages with path
func (_cbec *CT_Box) ValidateWithPath(path string) error {
	if _cbec.BoxPr != nil {
		if _afb := _cbec.BoxPr.ValidateWithPath(path + "\u002f\u0042\u006f\u0078\u0050\u0072"); _afb != nil {
			return _afb
		}
	}
	if _fed := _cbec.E.ValidateWithPath(path + "\u002f\u0045"); _fed != nil {
		return _fed
	}
	return nil
}

// ValidateWithPath validates the CT_MathPr and its children, prefixing error messages with path
func (_adca *CT_MathPr) ValidateWithPath(path string) error {
	if _adca.MathFont != nil {
		if _eecf := _adca.MathFont.ValidateWithPath(path + "\u002fM\u0061\u0074\u0068\u0046\u006f\u006et"); _eecf != nil {
			return _eecf
		}
	}
	if _adca.BrkBin != nil {
		if _aedaf := _adca.BrkBin.ValidateWithPath(path + "\u002fB\u0072\u006b\u0042\u0069\u006e"); _aedaf != nil {
			return _aedaf
		}
	}
	if _adca.BrkBinSub != nil {
		if _gfgb := _adca.BrkBinSub.ValidateWithPath(path + "\u002f\u0042\u0072\u006b\u0042\u0069\u006e\u0053\u0075\u0062"); _gfgb != nil {
			return _gfgb
		}
	}
	if _adca.SmallFrac != nil {
		if _eecd := _adca.SmallFrac.ValidateWithPath(path + "\u002f\u0053\u006d\u0061\u006c\u006c\u0046\u0072\u0061\u0063"); _eecd != nil {
			return _eecd
		}
	}
	if _adca.DispDef != nil {
		if _bdeb := _adca.DispDef.ValidateWithPath(path + "\u002f\u0044\u0069\u0073\u0070\u0044\u0065\u0066"); _bdeb != nil {
			return _bdeb
		}
	}
	if _adca.LMargin != nil {
		if _gege := _adca.LMargin.ValidateWithPath(path + "\u002f\u004c\u004d\u0061\u0072\u0067\u0069\u006e"); _gege != nil {
			return _gege
		}
	}
	if _adca.RMargin != nil {
		if _gbbd := _adca.RMargin.ValidateWithPath(path + "\u002f\u0052\u004d\u0061\u0072\u0067\u0069\u006e"); _gbbd != nil {
			return _gbbd
		}
	}
	if _adca.DefJc != nil {
		if _fbde := _adca.DefJc.ValidateWithPath(path + "\u002f\u0044\u0065\u0066\u004a\u0063"); _fbde != nil {
			return _fbde
		}
	}
	if _adca.PreSp != nil {
		if _abdd := _adca.PreSp.ValidateWithPath(path + "\u002f\u0050\u0072\u0065\u0053\u0070"); _abdd != nil {
			return _abdd
		}
	}
	if _adca.PostSp != nil {
		if _bead := _adca.PostSp.ValidateWithPath(path + "\u002fP\u006f\u0073\u0074\u0053\u0070"); _bead != nil {
			return _bead
		}
	}
	if _adca.InterSp != nil {
		if _ecdg := _adca.InterSp.ValidateWithPath(path + "\u002f\u0049\u006e\u0074\u0065\u0072\u0053\u0070"); _ecdg != nil {
			return _ecdg
		}
	}
	if _adca.IntraSp != nil {
		if _fade := _adca.IntraSp.ValidateWithPath(path + "\u002f\u0049\u006e\u0074\u0072\u0061\u0053\u0070"); _fade != nil {
			return _fade
		}
	}
	if _adca.Choice != nil {
		if _acbc := _adca.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _acbc != nil {
			return _acbc
		}
	}
	if _adca.IntLim != nil {
		if _fegc := _adca.IntLim.ValidateWithPath(path + "\u002fI\u006e\u0074\u004c\u0069\u006d"); _fegc != nil {
			return _fegc
		}
	}
	if _adca.NaryLim != nil {
		if _dceg := _adca.NaryLim.ValidateWithPath(path + "\u002f\u004e\u0061\u0072\u0079\u004c\u0069\u006d"); _dceg != nil {
			return _dceg
		}
	}
	return nil
}

// ValidateWithPath validates the CT_FPr and its children, prefixing error messages with path
func (_edfd *CT_FPr) ValidateWithPath(path string) error {
	if _edfd.Type != nil {
		if _ddb := _edfd.Type.ValidateWithPath(path + "\u002f\u0054\u0079p\u0065"); _ddb != nil {
			return _ddb
		}
	}
	if _edfd.CtrlPr != nil {
		if _bdfe := _edfd.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _bdfe != nil {
			return _bdfe
		}
	}
	return nil
}

type EG_OMathMathElements struct {
	Acc       *CT_Acc
	Bar       *CT_Bar
	Box       *CT_Box
	BorderBox *CT_BorderBox
	D         *CT_D
	EqArr     *CT_EqArr
	F         *CT_F
	Func      *CT_Func
	GroupChr  *CT_GroupChr
	LimLow    *CT_LimLow
	LimUpp    *CT_LimUpp
	M         *CT_M
	Nary      *CT_Nary
	Phant     *CT_Phant
	Rad       *CT_Rad
	SPre      *CT_SPre
	SSub      *CT_SSub
	SSubSup   *CT_SSubSup
	SSup      *CT_SSup
	R         *CT_R
}

func (_facde *ST_FType) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bgcaf, _fbaedd := d.Token()
	if _fbaedd != nil {
		return _fbaedd
	}
	if _accc, _efcg := _bgcaf.(_g.EndElement); _efcg && _accc.Name == start.Name {
		*_facde = 1
		return nil
	}
	if _fgbd, _fbfgb := _bgcaf.(_g.CharData); !_fbfgb {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bgcaf)
	} else {
		switch string(_fgbd) {
		case "":
			*_facde = 0
		case "\u0062\u0061\u0072":
			*_facde = 1
		case "\u0073\u006b\u0077":
			*_facde = 2
		case "\u006c\u0069\u006e":
			*_facde = 3
		case "\u006e\u006f\u0042a\u0072":
			*_facde = 4
		}
	}
	_bgcaf, _fbaedd = d.Token()
	if _fbaedd != nil {
		return _fbaedd
	}
	if _dgecc, _bgdge := _bgcaf.(_g.EndElement); _bgdge && _dgecc.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bgcaf)
}

type CT_Char struct{ ValAttr string }

func NewCT_EqArrPr() *CT_EqArrPr { _dfgf := &CT_EqArrPr{}; return _dfgf }

type CT_SSup struct {
	SSupPr *CT_SSupPr
	E      *CT_OMathArg
	Sup    *CT_OMathArg
}

func (_fcf *CT_LimLowPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_afae:
	for {
		_bafff, _baea := d.Token()
		if _baea != nil {
			return _baea
		}
		switch _ceffc := _bafff.(type) {
		case _g.StartElement:
			switch _ceffc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_fcf.CtrlPr = NewCT_CtrlPr()
				if _adf := d.DecodeElement(_fcf.CtrlPr, &_ceffc); _adf != nil {
					return _adf
				}
			default:
				_f.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004c\u0069\u006d\u004c\u006f\u0077\u0050\u0072\u0020\u0025\u0076", _ceffc.Name)
				if _abd := d.Skip(); _abd != nil {
					return _abd
				}
			}
		case _g.EndElement:
			break _afae
		case _g.CharData:
		}
	}
	return nil
}

func (_dbbf *CT_GroupChr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_dbbf.E = NewCT_OMathArg()
_cfdc:
	for {
		_cgge, _cbed := d.Token()
		if _cbed != nil {
			return _cbed
		}
		switch _febg := _cgge.(type) {
		case _g.StartElement:
			switch _febg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072"}:
				_dbbf.GroupChrPr = NewCT_GroupChrPr()
				if _egg := d.DecodeElement(_dbbf.GroupChrPr, &_febg); _egg != nil {
					return _egg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _fcbg := d.DecodeElement(_dbbf.E, &_febg); _fcbg != nil {
					return _fcbg
				}
			default:
				_f.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0020\u0025\u0076", _febg.Name)
				if _fcbgc := d.Skip(); _fcbgc != nil {
					return _fcbgc
				}
			}
		case _g.EndElement:
			break _cfdc
		case _g.CharData:
		}
	}
	return nil
}

func (_gaeac *CT_SSupPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_acce:
	for {
		_accf, _bffg := d.Token()
		if _bffg != nil {
			return _bffg
		}
		switch _dffd := _accf.(type) {
		case _g.StartElement:
			switch _dffd.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_gaeac.CtrlPr = NewCT_CtrlPr()
				if _cddd := d.DecodeElement(_gaeac.CtrlPr, &_dffd); _cddd != nil {
					return _cddd
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0053\u0053\u0075p\u0050\u0072 \u0025\u0076", _dffd.Name)
				if _daeed := d.Skip(); _daeed != nil {
					return _daeed
				}
			}
		case _g.EndElement:
			break _acce
		case _g.CharData:
		}
	}
	return nil
}

type CT_BoxPr struct {
	OpEmu   *CT_OnOff
	NoBreak *CT_OnOff
	Diff    *CT_OnOff
	Brk     *CT_ManualBreak
	Aln     *CT_OnOff
	CtrlPr  *CT_CtrlPr
}

// ValidateWithPath validates the CT_LimLow and its children, prefixing error messages with path
func (_fdbe *CT_LimLow) ValidateWithPath(path string) error {
	if _fdbe.LimLowPr != nil {
		if _cdcc := _fdbe.LimLowPr.ValidateWithPath(path + "\u002fL\u0069\u006d\u004c\u006f\u0077\u0050r"); _cdcc != nil {
			return _cdcc
		}
	}
	if _geeg := _fdbe.E.ValidateWithPath(path + "\u002f\u0045"); _geeg != nil {
		return _geeg
	}
	if _ceec := _fdbe.Lim.ValidateWithPath(path + "\u002f\u004c\u0069\u006d"); _ceec != nil {
		return _ceec
	}
	return nil
}

func NewCT_EqArr() *CT_EqArr { _adbb := &CT_EqArr{}; return _adbb }

func (_fcad *CT_Integer2) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _d.Sprintf("\u0025\u0076", _fcad.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_SpacingRule and its children
func (_bgda *CT_SpacingRule) Validate() error {
	return _bgda.ValidateWithPath("\u0043\u0054\u005f\u0053\u0070\u0061\u0063\u0069\u006eg\u0052\u0075\u006c\u0065")
}

func (_deed *CT_MCS) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	_dgbg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0063"}}
	for _, _edddf := range _deed.Mc {
		e.EncodeElement(_edddf, _dgbg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_MathPr struct {
	MathFont  *CT_String
	BrkBin    *CT_BreakBin
	BrkBinSub *CT_BreakBinSub
	SmallFrac *CT_OnOff
	DispDef   *CT_OnOff
	LMargin   *CT_TwipsMeasure
	RMargin   *CT_TwipsMeasure
	DefJc     *CT_OMathJc
	PreSp     *CT_TwipsMeasure
	PostSp    *CT_TwipsMeasure
	InterSp   *CT_TwipsMeasure
	IntraSp   *CT_TwipsMeasure
	Choice    *CT_MathPrChoice
	IntLim    *CT_LimLoc
	NaryLim   *CT_LimLoc
}

type CT_FType struct{ ValAttr ST_FType }

func (_aggg *CT_OMathParaPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_baeb:
	for {
		_fgfg, _fgbb := d.Token()
		if _fgbb != nil {
			return _fgbb
		}
		switch _ggebg := _fgfg.(type) {
		case _g.StartElement:
			switch _ggebg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006a\u0063"}:
				_aggg.Jc = NewCT_OMathJc()
				if _agab := d.DecodeElement(_aggg.Jc, &_ggebg); _agab != nil {
					return _agab
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0050\u0061r\u0061P\u0072\u0020\u0025\u0076", _ggebg.Name)
				if _beef := d.Skip(); _beef != nil {
					return _beef
				}
			}
		case _g.EndElement:
			break _baeb
		case _g.CharData:
		}
	}
	return nil
}

type CT_GroupChrPr struct {
	Chr    *CT_Char
	Pos    *CT_TopBot
	VertJc *CT_TopBot
	CtrlPr *CT_CtrlPr
}

// ValidateWithPath validates the CT_OMathPara and its children, prefixing error messages with path
func (_bfea *CT_OMathPara) ValidateWithPath(path string) error {
	if _bfea.OMathParaPr != nil {
		if _edab := _bfea.OMathParaPr.ValidateWithPath(path + "\u002f\u004f\u004da\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"); _edab != nil {
			return _edab
		}
	}
	for _gadc, _gbge := range _bfea.OMath {
		if _fcab := _gbge.ValidateWithPath(_d.Sprintf("\u0025\u0073\u002fO\u004d\u0061\u0074\u0068\u005b\u0025\u0064\u005d", path, _gadc)); _fcab != nil {
			return _fcab
		}
	}
	return nil
}

// ValidateWithPath validates the CT_MCPr and its children, prefixing error messages with path
func (_addg *CT_MCPr) ValidateWithPath(path string) error {
	if _addg.Count != nil {
		if _dcga := _addg.Count.ValidateWithPath(path + "\u002f\u0043\u006f\u0075\u006e\u0074"); _dcga != nil {
			return _dcga
		}
	}
	if _addg.McJc != nil {
		if _gbec := _addg.McJc.ValidateWithPath(path + "\u002f\u004d\u0063J\u0063"); _gbec != nil {
			return _gbec
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Char and its children, prefixing error messages with path
func (_cdcb *CT_Char) ValidateWithPath(path string) error { return nil }

func (_fbfd *CT_XAlign) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_dfgbg, _ddff := _fbfd.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _ddff != nil {
		return _ddff
	}
	start.Attr = append(start.Attr, _dfgbg)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the OMath and its children, prefixing error messages with path
func (_eddaa *OMath) ValidateWithPath(path string) error {
	if _aggga := _eddaa.CT_OMath.ValidateWithPath(path); _aggga != nil {
		return _aggga
	}
	return nil
}

func NewCT_Rad() *CT_Rad {
	_ggef := &CT_Rad{}
	_ggef.Deg = NewCT_OMathArg()
	_ggef.E = NewCT_OMathArg()
	return _ggef
}

func (_geac *CT_Script) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _eccd := range start.Attr {
		if _eccd.Name.Local == "\u0076\u0061\u006c" {
			_geac.ValAttr.UnmarshalXMLAttr(_eccd)
			continue
		}
	}
	for {
		_beab, _daae := d.Token()
		if _daae != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0072i\u0070\u0074\u003a\u0020\u0025\u0073", _daae)
		}
		if _dgege, _egcf := _beab.(_g.EndElement); _egcf && _dgege.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_YAlign and its children
func (_eacg *CT_YAlign) Validate() error {
	return _eacg.ValidateWithPath("\u0043T\u005f\u0059\u0041\u006c\u0069\u0067n")
}

// ValidateWithPath validates the CT_Acc and its children, prefixing error messages with path
func (_geg *CT_Acc) ValidateWithPath(path string) error {
	if _geg.AccPr != nil {
		if _fa := _geg.AccPr.ValidateWithPath(path + "\u002f\u0041\u0063\u0063\u0050\u0072"); _fa != nil {
			return _fa
		}
	}
	if _ca := _geg.E.ValidateWithPath(path + "\u002f\u0045"); _ca != nil {
		return _ca
	}
	return nil
}

func (_cbebf ST_Style) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_cbebf.String(), start)
}

func (_cabgd *MathPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006d"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0073\u0063\u0068\u0065m\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0068\u0061\u0072e\u0064\u0054\u0079\u0070\u0065\u0073"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0077"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065s\u0073i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u00306\u002fm\u0061\u0069n"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006d\u003a\u006d\u0061\u0074\u0068\u0050\u0072"
	return _cabgd.CT_MathPr.MarshalXML(e, start)
}

func (_ccbad *EG_ScriptStyle) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cccea:
	for {
		_ffba, _bfaec := d.Token()
		if _bfaec != nil {
			return _bfaec
		}
		switch _gafd := _ffba.(type) {
		case _g.StartElement:
			switch _gafd.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0063\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0063\u0072"}:
				_ccbad.Scr = NewCT_Script()
				if _cfdgd := d.DecodeElement(_ccbad.Scr, &_gafd); _cfdgd != nil {
					return _cfdgd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0074\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0074\u0079"}:
				_ccbad.Sty = NewCT_Style()
				if _adg := d.DecodeElement(_ccbad.Sty, &_gafd); _adg != nil {
					return _adg
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0045\u0047\u005f\u0053\u0063\u0072\u0069\u0070\u0074\u0053t\u0079l\u0065\u0020\u0025\u0076", _gafd.Name)
				if _bcdd := d.Skip(); _bcdd != nil {
					return _bcdd
				}
			}
		case _g.EndElement:
			break _cccea
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SSub and its children
func (_gebd *CT_SSub) Validate() error {
	return _gebd.ValidateWithPath("\u0043T\u005f\u0053\u0053\u0075\u0062")
}

func NewCT_RChoice() *CT_RChoice { _afbc := &CT_RChoice{}; return _afbc }

func (_de *CT_BarPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fag:
	for {
		_ggg, _ag := d.Token()
		if _ag != nil {
			return _ag
		}
		switch _ecb := _ggg.(type) {
		case _g.StartElement:
			switch _ecb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u006f\u0073"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u006f\u0073"}:
				_de.Pos = NewCT_TopBot()
				if _ddf := d.DecodeElement(_de.Pos, &_ecb); _ddf != nil {
					return _ddf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_de.CtrlPr = NewCT_CtrlPr()
				if _ee := d.DecodeElement(_de.CtrlPr, &_ecb); _ee != nil {
					return _ee
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u0061\u0072\u0050\u0072\u0020\u0025\u0076", _ecb.Name)
				if _eg := d.Skip(); _eg != nil {
					return _eg
				}
			}
		case _g.EndElement:
			break _fag
		case _g.CharData:
		}
	}
	return nil
}

func (_befb *CT_SPrePr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _befb.CtrlPr != nil {
		_ged := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_befb.CtrlPr, _ged)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_OMathParaPr() *CT_OMathParaPr { _ggdc := &CT_OMathParaPr{}; return _ggdc }

func NewCT_Integer2() *CT_Integer2 { _fadb := &CT_Integer2{}; _fadb.ValAttr = -2; return _fadb }

func (_babb *CT_OnOff) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _babb.ValAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _d.Sprintf("\u0025\u0076", *_babb.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_bgdba *ST_TopBot) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bddbd, _adac := d.Token()
	if _adac != nil {
		return _adac
	}
	if _gfac, _ggded := _bddbd.(_g.EndElement); _ggded && _gfac.Name == start.Name {
		*_bgdba = 1
		return nil
	}
	if _gebbb, _geed := _bddbd.(_g.CharData); !_geed {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bddbd)
	} else {
		switch string(_gebbb) {
		case "":
			*_bgdba = 0
		case "\u0074\u006f\u0070":
			*_bgdba = 1
		case "\u0062\u006f\u0074":
			*_bgdba = 2
		}
	}
	_bddbd, _adac = d.Token()
	if _adac != nil {
		return _adac
	}
	if _gfgd, _eaeab := _bddbd.(_g.EndElement); _eaeab && _gfgd.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bddbd)
}

func (_gfcd *CT_SSub) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_gfcd.E = NewCT_OMathArg()
	_gfcd.Sub = NewCT_OMathArg()
_dbdb:
	for {
		_bdae, _fbb := d.Token()
		if _fbb != nil {
			return _fbb
		}
		switch _bgab := _bdae.(type) {
		case _g.StartElement:
			switch _bgab.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062\u0050\u0072"}:
				_gfcd.SSubPr = NewCT_SSubPr()
				if _ffad := d.DecodeElement(_gfcd.SSubPr, &_bgab); _ffad != nil {
					return _ffad
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _ebag := d.DecodeElement(_gfcd.E, &_bgab); _ebag != nil {
					return _ebag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0062"}:
				if _fgeba := d.DecodeElement(_gfcd.Sub, &_bgab); _fgeba != nil {
					return _fgeba
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0053\u0053\u0075\u0062\u0020\u0025\u0076", _bgab.Name)
				if _ecff := d.Skip(); _ecff != nil {
					return _ecff
				}
			}
		case _g.EndElement:
			break _dbdb
		case _g.CharData:
		}
	}
	return nil
}

func (_gdbb ST_BreakBinSub) ValidateWithPath(path string) error {
	switch _gdbb {
	case 0, 1, 2, 3:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gdbb))
	}
	return nil
}

type ST_BreakBin byte

// Validate validates the CT_PhantPr and its children
func (_fddd *CT_PhantPr) Validate() error {
	return _fddd.ValidateWithPath("\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074\u0050\u0072")
}

func (_agbe *CT_Style) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _agbe.ValAttr != ST_StyleUnset {
		_caab, _bdbab := _agbe.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _bdbab != nil {
			return _bdbab
		}
		start.Attr = append(start.Attr, _caab)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_abef *CT_GroupChrPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _abef.Chr != nil {
		_aceb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063h\u0072"}}
		e.EncodeElement(_abef.Chr, _aceb)
	}
	if _abef.Pos != nil {
		_gdfg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0070o\u0073"}}
		e.EncodeElement(_abef.Pos, _gdfg)
	}
	if _abef.VertJc != nil {
		_ccc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0076\u0065\u0072\u0074\u004a\u0063"}}
		e.EncodeElement(_abef.VertJc, _ccc)
	}
	if _abef.CtrlPr != nil {
		_deca := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_abef.CtrlPr, _deca)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_Style struct{ ValAttr ST_Style }

func (_adfb ST_Script) Validate() error { return _adfb.ValidateWithPath("") }

func NewCT_SSup() *CT_SSup {
	_bffe := &CT_SSup{}
	_bffe.E = NewCT_OMathArg()
	_bffe.Sup = NewCT_OMathArg()
	return _bffe
}

type CT_ManualBreak struct{ AlnAtAttr *int64 }

func NewCT_SSupPr() *CT_SSupPr { _gcdg := &CT_SSupPr{}; return _gcdg }

// Validate validates the CT_Integer2 and its children
func (_cbeg *CT_Integer2) Validate() error {
	return _cbeg.ValidateWithPath("C\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032")
}

func (_gcad *CT_MCPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gcad.Count != nil {
		_cda := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0063\u006f\u0075\u006e\u0074"}}
		e.EncodeElement(_gcad.Count, _cda)
	}
	if _gcad.McJc != nil {
		_cdca := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0063\u004a\u0063"}}
		e.EncodeElement(_gcad.McJc, _cdca)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_cgefa ST_FType) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_beeeaa := _g.Attr{}
	_beeeaa.Name = name
	switch _cgefa {
	case ST_FTypeUnset:
		_beeeaa.Value = ""
	case ST_FTypeBar:
		_beeeaa.Value = "\u0062\u0061\u0072"
	case ST_FTypeSkw:
		_beeeaa.Value = "\u0073\u006b\u0077"
	case ST_FTypeLin:
		_beeeaa.Value = "\u006c\u0069\u006e"
	case ST_FTypeNoBar:
		_beeeaa.Value = "\u006e\u006f\u0042a\u0072"
	}
	return _beeeaa, nil
}

// Validate validates the CT_OMathPara and its children
func (_baed *CT_OMathPara) Validate() error {
	return _baed.ValidateWithPath("\u0043\u0054\u005fO\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061")
}

func NewCT_Nary() *CT_Nary {
	_cbeb := &CT_Nary{}
	_cbeb.Sub = NewCT_OMathArg()
	_cbeb.Sup = NewCT_OMathArg()
	_cbeb.E = NewCT_OMathArg()
	return _cbeb
}

func (_fcef *CT_OMathParaPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fcef.Jc != nil {
		_gagac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006a\u0063"}}
		e.EncodeElement(_fcef.Jc, _gagac)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_bcab *CT_TwipsMeasure) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _d.Sprintf("\u0025\u0076", _bcab.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_daf *CT_BorderBoxPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _daf.HideTop != nil {
		_gff := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0068\u0069\u0064\u0065\u0054\u006fp"}}
		e.EncodeElement(_daf.HideTop, _gff)
	}
	if _daf.HideBot != nil {
		_cbe := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0068\u0069\u0064\u0065\u0042\u006ft"}}
		e.EncodeElement(_daf.HideBot, _cbe)
	}
	if _daf.HideLeft != nil {
		_cdb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0068\u0069\u0064\u0065\u004c\u0065\u0066\u0074"}}
		e.EncodeElement(_daf.HideLeft, _cdb)
	}
	if _daf.HideRight != nil {
		_ggde := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0068\u0069\u0064\u0065\u0052\u0069\u0067\u0068\u0074"}}
		e.EncodeElement(_daf.HideRight, _ggde)
	}
	if _daf.StrikeH != nil {
		_fdf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0074\u0072\u0069\u006b\u0065H"}}
		e.EncodeElement(_daf.StrikeH, _fdf)
	}
	if _daf.StrikeV != nil {
		_adc := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0074\u0072\u0069\u006b\u0065V"}}
		e.EncodeElement(_daf.StrikeV, _adc)
	}
	if _daf.StrikeBLTR != nil {
		_ccg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073t\u0072\u0069\u006b\u0065\u0042\u004c\u0054\u0052"}}
		e.EncodeElement(_daf.StrikeBLTR, _ccg)
	}
	if _daf.StrikeTLBR != nil {
		_bc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073t\u0072\u0069\u006b\u0065\u0054\u004c\u0042\u0052"}}
		e.EncodeElement(_daf.StrikeTLBR, _bc)
	}
	if _daf.CtrlPr != nil {
		_fg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_daf.CtrlPr, _fg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_Box and its children
func (_dab *CT_Box) Validate() error {
	return _dab.ValidateWithPath("\u0043\u0054\u005f\u0042\u006f\u0078")
}

func (_fedd *CT_Integer2) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_fedd.ValAttr = -2
	for _, _ecd := range start.Attr {
		if _ecd.Name.Local == "\u0076\u0061\u006c" {
			_ebfa, _afd := _be.ParseInt(_ecd.Value, 10, 64)
			if _afd != nil {
				return _afd
			}
			_fedd.ValAttr = _ebfa
			continue
		}
	}
	for {
		_eaa, _cadf := d.Token()
		if _cadf != nil {
			return _d.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032\u003a\u0020\u0025\u0073", _cadf)
		}
		if _cbef, _aeda := _eaa.(_g.EndElement); _aeda && _cbef.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_UnSignedInteger and its children
func (_ccacg *CT_UnSignedInteger) Validate() error {
	return _ccacg.ValidateWithPath("\u0043T\u005fU\u006e\u0053\u0069\u0067\u006ee\u0064\u0049n\u0074\u0065\u0067\u0065\u0072")
}

func (_bddd *CT_SpacingRule) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _d.Sprintf("\u0025\u0076", _bddd.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_MathPrChoice() *CT_MathPrChoice { _fdad := &CT_MathPrChoice{}; return _fdad }

type CT_LimLowPr struct{ CtrlPr *CT_CtrlPr }

func (_dedb *CT_YAlign) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_ecec, _aecc := _dedb.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _aecc != nil {
		return _aecc
	}
	start.Attr = append(start.Attr, _ecec)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func ParseUnionST_TwipsMeasure(s string) (_dc.ST_TwipsMeasure, error) {
	_gaggc := _dc.ST_TwipsMeasure{}
	if _dc.ST_PositiveUniversalMeasurePatternRe.MatchString(s) {
		_gaggc.ST_PositiveUniversalMeasure = &s
	} else {
		_cbab, _faagf := _be.ParseFloat(s, 64)
		if _faagf != nil {
			return _gaggc, _d.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0025\u0073\u0020\u0061\u0073\u0020\u0075\u0069\u006e\u0074\u003a\u0020%\u0073", s, _faagf)
		}
		_gaggc.ST_UnsignedDecimalNumber = _c.Uint64(uint64(_cbab))
	}
	return _gaggc, nil
}

type CT_SSubPr struct{ CtrlPr *CT_CtrlPr }

// Validate validates the CT_MPr and its children
func (_gbcf *CT_MPr) Validate() error {
	return _gbcf.ValidateWithPath("\u0043\u0054\u005f\u004d\u0050\u0072")
}

func (_gac *CT_DPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gac.BegChr != nil {
		_baf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u0065\u0067\u0043\u0068\u0072"}}
		e.EncodeElement(_gac.BegChr, _baf)
	}
	if _gac.SepChr != nil {
		_bdg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0065\u0070\u0043\u0068\u0072"}}
		e.EncodeElement(_gac.SepChr, _bdg)
	}
	if _gac.EndChr != nil {
		_gfefa := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065\u006e\u0064\u0043\u0068\u0072"}}
		e.EncodeElement(_gac.EndChr, _gfefa)
	}
	if _gac.Grow != nil {
		_fff := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0067\u0072\u006f\u0077"}}
		e.EncodeElement(_gac.Grow, _fff)
	}
	if _gac.Shp != nil {
		_gfg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073h\u0070"}}
		e.EncodeElement(_gac.Shp, _gfg)
	}
	if _gac.CtrlPr != nil {
		_ceag := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_gac.CtrlPr, _ceag)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_Shp and its children
func (_ecdd *CT_Shp) Validate() error {
	return _ecdd.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0070")
}

func (_efca *CT_ManualBreak) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _bafffd := range start.Attr {
		if _bafffd.Name.Local == "\u0061\u006c\u006eA\u0074" {
			_cfce, _fagd := _be.ParseInt(_bafffd.Value, 10, 64)
			if _fagd != nil {
				return _fagd
			}
			_efca.AlnAtAttr = &_cfce
			continue
		}
	}
	for {
		_cde, _fgbg := d.Token()
		if _fgbg != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fM\u0061\u006e\u0075\u0061\u006c\u0042\u0072\u0065\u0061\u006b:\u0020\u0025\u0073", _fgbg)
		}
		if _bbae, _cag := _cde.(_g.EndElement); _cag && _bbae.Name == start.Name {
			break
		}
	}
	return nil
}

func NewCT_MR() *CT_MR { _dbefa := &CT_MR{}; return _dbefa }

type CT_BreakBin struct{ ValAttr ST_BreakBin }

func NewEG_ScriptStyle() *EG_ScriptStyle { _gfdbe := &EG_ScriptStyle{}; return _gfdbe }

// Validate validates the CT_OMathArg and its children
func (_caeb *CT_OMathArg) Validate() error {
	return _caeb.ValidateWithPath("C\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067")
}

func (_ga *CT_Bar) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ga.E = NewCT_OMathArg()
_cb:
	for {
		_gag, _gf := d.Token()
		if _gf != nil {
			return _gf
		}
		switch _fb := _gag.(type) {
		case _g.StartElement:
			switch _fb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072P\u0072"}:
				_ga.BarPr = NewCT_BarPr()
				if _dcf := d.DecodeElement(_ga.BarPr, &_fb); _dcf != nil {
					return _dcf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _afg := d.DecodeElement(_ga.E, &_fb); _afg != nil {
					return _afg
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u0061\u0072\u0020\u0025\u0076", _fb.Name)
				if _fc := d.Skip(); _fc != nil {
					return _fc
				}
			}
		case _g.EndElement:
			break _cb
		case _g.CharData:
		}
	}
	return nil
}

func (_eaaf *CT_LimUpp) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_eaaf.E = NewCT_OMathArg()
	_eaaf.Lim = NewCT_OMathArg()
_baeg:
	for {
		_ceca, _eaeg := d.Token()
		if _eaeg != nil {
			return _eaeg
		}
		switch _deg := _ceca.(type) {
		case _g.StartElement:
			switch _deg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070\u0050\u0072"}:
				_eaaf.LimUppPr = NewCT_LimUppPr()
				if _aaga := d.DecodeElement(_eaaf.LimUppPr, &_deg); _aaga != nil {
					return _aaga
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _debg := d.DecodeElement(_eaaf.E, &_deg); _debg != nil {
					return _debg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d"}:
				if _eebc := d.DecodeElement(_eaaf.Lim, &_deg); _eebc != nil {
					return _eebc
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004c\u0069\u006dU\u0070\u0070 \u0025\u0076", _deg.Name)
				if _degd := d.Skip(); _degd != nil {
					return _degd
				}
			}
		case _g.EndElement:
			break _baeg
		case _g.CharData:
		}
	}
	return nil
}

type CT_DPr struct {
	BegChr *CT_Char
	SepChr *CT_Char
	EndChr *CT_Char
	Grow   *CT_OnOff
	Shp    *CT_Shp
	CtrlPr *CT_CtrlPr
}

func (_ddebb *CT_SPre) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ddebb.SPrePr != nil {
		_dcad := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0050\u0072\u0065\u0050\u0072"}}
		e.EncodeElement(_ddebb.SPrePr, _dcad)
	}
	_cdef := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0062"}}
	e.EncodeElement(_ddebb.Sub, _cdef)
	_dfac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0070"}}
	e.EncodeElement(_ddebb.Sup, _dfac)
	_bfaf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_ddebb.E, _bfaf)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_cebg *CT_M) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_dcd:
	for {
		_bea, _fdca := d.Token()
		if _fdca != nil {
			return _fdca
		}
		switch _cfe := _bea.(type) {
		case _g.StartElement:
			switch _cfe.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0050\u0072"}:
				_cebg.MPr = NewCT_MPr()
				if _facf := d.DecodeElement(_cebg.MPr, &_cfe); _facf != nil {
					return _facf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0072"}:
				_afbe := NewCT_MR()
				if _fdae := d.DecodeElement(_afbe, &_cfe); _fdae != nil {
					return _fdae
				}
				_cebg.Mr = append(_cebg.Mr, _afbe)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u004d\u0020\u0025\u0076", _cfe.Name)
				if _edfb := d.Skip(); _edfb != nil {
					return _edfb
				}
			}
		case _g.EndElement:
			break _dcd
		case _g.CharData:
		}
	}
	return nil
}

func (_bdbe *CT_Nary) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _bdbe.NaryPr != nil {
		_fcgc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006e\u0061\u0072\u0079\u0050\u0072"}}
		e.EncodeElement(_bdbe.NaryPr, _fcgc)
	}
	_bfebf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0062"}}
	e.EncodeElement(_bdbe.Sub, _bfebf)
	_acag := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0070"}}
	e.EncodeElement(_bdbe.Sup, _acag)
	_cfdb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_bdbe.E, _cfdb)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_ffgb *CT_MathPrChoice) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _ffgb.WrapIndent != nil {
		_bfce := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0077r\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}}
		e.EncodeElement(_ffgb.WrapIndent, _bfce)
	}
	if _ffgb.WrapRight != nil {
		_cbc := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0077\u0072\u0061\u0070\u0052\u0069\u0067\u0068\u0074"}}
		e.EncodeElement(_ffgb.WrapRight, _cbc)
	}
	return nil
}

func NewCT_D() *CT_D { _cge := &CT_D{}; return _cge }

// ValidateWithPath validates the CT_PhantPr and its children, prefixing error messages with path
func (_dfef *CT_PhantPr) ValidateWithPath(path string) error {
	if _dfef.Show != nil {
		if _bcde := _dfef.Show.ValidateWithPath(path + "\u002f\u0053\u0068o\u0077"); _bcde != nil {
			return _bcde
		}
	}
	if _dfef.ZeroWid != nil {
		if _eag := _dfef.ZeroWid.ValidateWithPath(path + "\u002f\u005a\u0065\u0072\u006f\u0057\u0069\u0064"); _eag != nil {
			return _eag
		}
	}
	if _dfef.ZeroAsc != nil {
		if _fegeb := _dfef.ZeroAsc.ValidateWithPath(path + "\u002f\u005a\u0065\u0072\u006f\u0041\u0073\u0063"); _fegeb != nil {
			return _fegeb
		}
	}
	if _dfef.ZeroDesc != nil {
		if _adde := _dfef.ZeroDesc.ValidateWithPath(path + "\u002fZ\u0065\u0072\u006f\u0044\u0065\u0073c"); _adde != nil {
			return _adde
		}
	}
	if _dfef.Transp != nil {
		if _cgbb := _dfef.Transp.ValidateWithPath(path + "\u002fT\u0072\u0061\u006e\u0073\u0070"); _cgbb != nil {
			return _cgbb
		}
	}
	if _dfef.CtrlPr != nil {
		if _deag := _dfef.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _deag != nil {
			return _deag
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Bar and its children, prefixing error messages with path
func (_fbd *CT_Bar) ValidateWithPath(path string) error {
	if _fbd.BarPr != nil {
		if _bgb := _fbd.BarPr.ValidateWithPath(path + "\u002f\u0042\u0061\u0072\u0050\u0072"); _bgb != nil {
			return _bgb
		}
	}
	if _dag := _fbd.E.ValidateWithPath(path + "\u002f\u0045"); _dag != nil {
		return _dag
	}
	return nil
}

func (_cabbf ST_BreakBin) Validate() error { return _cabbf.ValidateWithPath("") }

// ValidateWithPath validates the CT_OMathArgPr and its children, prefixing error messages with path
func (_abfb *CT_OMathArgPr) ValidateWithPath(path string) error {
	if _abfb.ArgSz != nil {
		if _gfce := _abfb.ArgSz.ValidateWithPath(path + "\u002f\u0041\u0072\u0067\u0053\u007a"); _gfce != nil {
			return _gfce
		}
	}
	return nil
}

type CT_SpacingRule struct{ ValAttr int64 }

func (_adcd *CT_Phant) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_adcd.E = NewCT_OMathArg()
_bded:
	for {
		_agdce, _ceffd := d.Token()
		if _ceffd != nil {
			return _ceffd
		}
		switch _gbgb := _agdce.(type) {
		case _g.StartElement:
			switch _gbgb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070h\u0061\u006e\u0074\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070h\u0061\u006e\u0074\u0050\u0072"}:
				_adcd.PhantPr = NewCT_PhantPr()
				if _gage := d.DecodeElement(_adcd.PhantPr, &_gbgb); _gage != nil {
					return _gage
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _bgbbf := d.DecodeElement(_adcd.E, &_gbgb); _bgbbf != nil {
					return _bgbbf
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074\u0020\u0025\u0076", _gbgb.Name)
				if _ebcb := d.Skip(); _ebcb != nil {
					return _ebcb
				}
			}
		case _g.EndElement:
			break _bded
		case _g.CharData:
		}
	}
	return nil
}

func (_ddeb *CT_RChoice) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _ddeb.T != nil {
		_aaab := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0074"}}
		for _, _gfeb := range _ddeb.T {
			e.EncodeElement(_gfeb, _aaab)
		}
	}
	return nil
}

type CT_LimUpp struct {
	LimUppPr *CT_LimUppPr
	E        *CT_OMathArg
	Lim      *CT_OMathArg
}

func NewCT_OMathJc() *CT_OMathJc { _ecf := &CT_OMathJc{}; return _ecf }

func (_dae *CT_F) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dae.FPr != nil {
		_facb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0066P\u0072"}}
		e.EncodeElement(_dae.FPr, _facb)
	}
	_bdfb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006eu\u006d"}}
	e.EncodeElement(_dae.Num, _bdfb)
	_eabf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064e\u006e"}}
	e.EncodeElement(_dae.Den, _eabf)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_FType and its children
func (_cfff *CT_FType) Validate() error {
	return _cfff.ValidateWithPath("\u0043\u0054\u005f\u0046\u0054\u0079\u0070\u0065")
}

func (_gea *CT_Func) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gea.FuncPr != nil {
		_gec := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0066\u0075\u006e\u0063\u0050\u0072"}}
		e.EncodeElement(_gea.FuncPr, _gec)
	}
	_acd := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0066\u004e\u0061\u006d\u0065"}}
	e.EncodeElement(_gea.FName, _acd)
	_bb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_gea.E, _bb)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_cdcdb *ST_Jc) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_cdcdb = 0
	case "\u006c\u0065\u0066\u0074":
		*_cdcdb = 1
	case "\u0072\u0069\u0067h\u0074":
		*_cdcdb = 2
	case "\u0063\u0065\u006e\u0074\u0065\u0072":
		*_cdcdb = 3
	case "c\u0065\u006e\u0074\u0065\u0072\u0047\u0072\u006f\u0075\u0070":
		*_cdcdb = 4
	}
	return nil
}

// Validate validates the CT_MCPr and its children
func (_gcc *CT_MCPr) Validate() error {
	return _gcc.ValidateWithPath("\u0043T\u005f\u004d\u0043\u0050\u0072")
}

type CT_Integer255 struct{ ValAttr int64 }

func (_gbfg *CT_OMathPara) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_debgb:
	for {
		_ggag, _egc := d.Token()
		if _egc != nil {
			return _egc
		}
		switch _dbgg := _ggag.(type) {
		case _g.StartElement:
			switch _dbgg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "o\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "o\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}:
				_gbfg.OMathParaPr = NewCT_OMathParaPr()
				if _fafd := d.DecodeElement(_gbfg.OMathParaPr, &_dbgg); _fafd != nil {
					return _fafd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006f\u004d\u0061t\u0068"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006f\u004d\u0061t\u0068"}:
				_babgb := NewCT_OMath()
				if _feaf := d.DecodeElement(_babgb, &_dbgg); _feaf != nil {
					return _feaf
				}
				_gbfg.OMath = append(_gbfg.OMath, _babgb)
			default:
				_f.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_O\u004d\u0061t\u0068\u0050\u0061\u0072\u0061\u0020\u0025\u0076", _dbgg.Name)
				if _cadfcb := d.Skip(); _cadfcb != nil {
					return _cadfcb
				}
			}
		case _g.EndElement:
			break _debgb
		case _g.CharData:
		}
	}
	return nil
}

func (_fabf *CT_D) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fee:
	for {
		_aac, _acea := d.Token()
		if _acea != nil {
			return _acea
		}
		switch _ega := _aac.(type) {
		case _g.StartElement:
			switch _ega.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0050\u0072"}:
				_fabf.DPr = NewCT_DPr()
				if _edd := d.DecodeElement(_fabf.DPr, &_ega); _edd != nil {
					return _edd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				_bce := NewCT_OMathArg()
				if _dac := d.DecodeElement(_bce, &_ega); _dac != nil {
					return _dac
				}
				_fabf.E = append(_fabf.E, _bce)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u0044\u0020\u0025\u0076", _ega.Name)
				if _efb := d.Skip(); _efb != nil {
					return _efb
				}
			}
		case _g.EndElement:
			break _fee
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Script and its children
func (_dddc *CT_Script) Validate() error {
	return _dddc.ValidateWithPath("\u0043T\u005f\u0053\u0063\u0072\u0069\u0070t")
}

func (_cgad *OMathPara) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cgad.CT_OMathPara = *NewCT_OMathPara()
_caed:
	for {
		_egcag, _ceee := d.Token()
		if _ceee != nil {
			return _ceee
		}
		switch _cgee := _egcag.(type) {
		case _g.StartElement:
			switch _cgee.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "o\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "o\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}:
				_cgad.OMathParaPr = NewCT_OMathParaPr()
				if _caafd := d.DecodeElement(_cgad.OMathParaPr, &_cgee); _caafd != nil {
					return _caafd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006f\u004d\u0061t\u0068"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006f\u004d\u0061t\u0068"}:
				_ffcf := NewCT_OMath()
				if _fddbe := d.DecodeElement(_ffcf, &_cgee); _fddbe != nil {
					return _fddbe
				}
				_cgad.OMath = append(_cgad.OMath, _ffcf)
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u004f\u004d\u0061\u0074\u0068\u0050a\u0072\u0061 \u0025\u0076", _cgee.Name)
				if _addfb := d.Skip(); _addfb != nil {
					return _addfb
				}
			}
		case _g.EndElement:
			break _caed
		case _g.CharData:
		}
	}
	return nil
}

func (_bceac ST_Script) String() string {
	switch _bceac {
	case 0:
		return ""
	case 1:
		return "\u0072\u006f\u006da\u006e"
	case 2:
		return "\u0073\u0063\u0072\u0069\u0070\u0074"
	case 3:
		return "\u0066r\u0061\u006b\u0074\u0075\u0072"
	case 4:
		return "\u0064\u006f\u0075\u0062\u006c\u0065\u002d\u0073\u0074\u0072\u0075\u0063\u006b"
	case 5:
		return "\u0073\u0061\u006e\u0073\u002d\u0073\u0065\u0072\u0069\u0066"
	case 6:
		return "\u006do\u006e\u006f\u0073\u0070\u0061\u0063e"
	}
	return ""
}

// Validate validates the CT_SSupPr and its children
func (_fage *CT_SSupPr) Validate() error {
	return _fage.ValidateWithPath("\u0043T\u005f\u0053\u0053\u0075\u0070\u0050r")
}

func NewCT_MPr() *CT_MPr { _fbfg := &CT_MPr{}; return _fbfg }

type CT_EqArr struct {
	EqArrPr *CT_EqArrPr
	E       []*CT_OMathArg
}

func (_bdbf *CT_Style) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _fgff := range start.Attr {
		if _fgff.Name.Local == "\u0076\u0061\u006c" {
			_bdbf.ValAttr.UnmarshalXMLAttr(_fgff)
			continue
		}
	}
	for {
		_aafe, _ccbb := d.Token()
		if _ccbb != nil {
			return _d.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fS\u0074\u0079\u006c\u0065: \u0025\u0073", _ccbb)
		}
		if _bbbb, _ccae := _aafe.(_g.EndElement); _ccae && _bbbb.Name == start.Name {
			break
		}
	}
	return nil
}

func (_aeaae *CT_SSubSup) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_aeaae.E = NewCT_OMathArg()
	_aeaae.Sub = NewCT_OMathArg()
	_aeaae.Sup = NewCT_OMathArg()
_bdff:
	for {
		_eaead, _bcbbb := d.Token()
		if _bcbbb != nil {
			return _bcbbb
		}
		switch _deac := _eaead.(type) {
		case _g.StartElement:
			switch _deac.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070\u0050r"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070\u0050r"}:
				_aeaae.SSubSupPr = NewCT_SSubSupPr()
				if _gdgbe := d.DecodeElement(_aeaae.SSubSupPr, &_deac); _gdgbe != nil {
					return _gdgbe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _dcdg := d.DecodeElement(_aeaae.E, &_deac); _dcdg != nil {
					return _dcdg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0062"}:
				if _dcgf := d.DecodeElement(_aeaae.Sub, &_deac); _dcgf != nil {
					return _dcgf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0070"}:
				if _bfae := d.DecodeElement(_aeaae.Sup, &_deac); _bfae != nil {
					return _bfae
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fS\u0053\u0075\u0062\u0053\u0075\u0070\u0020\u0025\u0076", _deac.Name)
				if _bfcd := d.Skip(); _bfcd != nil {
					return _bfcd
				}
			}
		case _g.EndElement:
			break _bdff
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_GroupChr and its children, prefixing error messages with path
func (_fgdc *CT_GroupChr) ValidateWithPath(path string) error {
	if _fgdc.GroupChrPr != nil {
		if _bfbb := _fgdc.GroupChrPr.ValidateWithPath(path + "/\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072"); _bfbb != nil {
			return _bfbb
		}
	}
	if _dge := _fgdc.E.ValidateWithPath(path + "\u002f\u0045"); _dge != nil {
		return _dge
	}
	return nil
}

func (_egfe *CT_Text) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _egfe.SpaceAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u003a\u0073\u0070\u0061\u0063e"}, Value: _d.Sprintf("\u0025\u0076", *_egfe.SpaceAttr)})
	}
	e.EncodeElement(_egfe.Content, start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_LimUppPr() *CT_LimUppPr { _dbfc := &CT_LimUppPr{}; return _dbfc }

func NewMathPr() *MathPr { _gbfd := &MathPr{}; _gbfd.CT_MathPr = *NewCT_MathPr(); return _gbfd }

// Validate validates the CT_LimLowPr and its children
func (_caag *CT_LimLowPr) Validate() error {
	return _caag.ValidateWithPath("C\u0054\u005f\u004c\u0069\u006d\u004c\u006f\u0077\u0050\u0072")
}

func NewCT_OnOff() *CT_OnOff { _eecfc := &CT_OnOff{}; return _eecfc }

func (_eccfd *ST_Jc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_gacb, _cabcf := d.Token()
	if _cabcf != nil {
		return _cabcf
	}
	if _efaf, _fdcf := _gacb.(_g.EndElement); _fdcf && _efaf.Name == start.Name {
		*_eccfd = 1
		return nil
	}
	if _fdbf, _gcac := _gacb.(_g.CharData); !_gcac {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gacb)
	} else {
		switch string(_fdbf) {
		case "":
			*_eccfd = 0
		case "\u006c\u0065\u0066\u0074":
			*_eccfd = 1
		case "\u0072\u0069\u0067h\u0074":
			*_eccfd = 2
		case "\u0063\u0065\u006e\u0074\u0065\u0072":
			*_eccfd = 3
		case "c\u0065\u006e\u0074\u0065\u0072\u0047\u0072\u006f\u0075\u0070":
			*_eccfd = 4
		}
	}
	_gacb, _cabcf = d.Token()
	if _cabcf != nil {
		return _cabcf
	}
	if _bgcfd, _bbagb := _gacb.(_g.EndElement); _bbagb && _bgcfd.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gacb)
}

func (_eaga *CT_SSupPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _eaga.CtrlPr != nil {
		_eefcc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_eaga.CtrlPr, _eefcc)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_F and its children, prefixing error messages with path
func (_ggga *CT_F) ValidateWithPath(path string) error {
	if _ggga.FPr != nil {
		if _cgd := _ggga.FPr.ValidateWithPath(path + "\u002f\u0046\u0050\u0072"); _cgd != nil {
			return _cgd
		}
	}
	if _bdeg := _ggga.Num.ValidateWithPath(path + "\u002f\u004e\u0075\u006d"); _bdeg != nil {
		return _bdeg
	}
	if _gfff := _ggga.Den.ValidateWithPath(path + "\u002f\u0044\u0065\u006e"); _gfff != nil {
		return _gfff
	}
	return nil
}

func (_dadc *CT_YAlign) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_dadc.ValAttr = _dc.ST_YAlign(1)
	for _, _cgcd := range start.Attr {
		if _cgcd.Name.Local == "\u0076\u0061\u006c" {
			_dadc.ValAttr.UnmarshalXMLAttr(_cgcd)
			continue
		}
	}
	for {
		_ccag, _edddd := d.Token()
		if _edddd != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0059\u0041\u006ci\u0067\u006e\u003a\u0020\u0025\u0073", _edddd)
		}
		if _efdg, _gcgd := _ccag.(_g.EndElement); _gcgd && _efdg.Name == start.Name {
			break
		}
	}
	return nil
}

func NewCT_BreakBin() *CT_BreakBin { _bae := &CT_BreakBin{}; return _bae }

// ValidateWithPath validates the CT_FType and its children, prefixing error messages with path
func (_afec *CT_FType) ValidateWithPath(path string) error {
	if _afec.ValAttr == ST_FTypeUnset {
		return _d.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _fddb := _afec.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _fddb != nil {
		return _fddb
	}
	return nil
}

type CT_OnOff struct{ ValAttr *_dc.ST_OnOff }

// ValidateWithPath validates the CT_RadPr and its children, prefixing error messages with path
func (_bbf *CT_RadPr) ValidateWithPath(path string) error {
	if _bbf.DegHide != nil {
		if _dbab := _bbf.DegHide.ValidateWithPath(path + "\u002f\u0044\u0065\u0067\u0048\u0069\u0064\u0065"); _dbab != nil {
			return _dbab
		}
	}
	if _bbf.CtrlPr != nil {
		if _fbdg := _bbf.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _fbdg != nil {
			return _fbdg
		}
	}
	return nil
}

// Validate validates the CT_RadPr and its children
func (_bgf *CT_RadPr) Validate() error {
	return _bgf.ValidateWithPath("\u0043\u0054\u005f\u0052\u0061\u0064\u0050\u0072")
}

type CT_PhantPr struct {
	Show     *CT_OnOff
	ZeroWid  *CT_OnOff
	ZeroAsc  *CT_OnOff
	ZeroDesc *CT_OnOff
	Transp   *CT_OnOff
	CtrlPr   *CT_CtrlPr
}

// ValidateWithPath validates the CT_String and its children, prefixing error messages with path
func (_afdd *CT_String) ValidateWithPath(path string) error { return nil }

func (_eaaff *CT_MPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_bfeb:
	for {
		_efg, _cfec := d.Token()
		if _cfec != nil {
			return _cfec
		}
		switch _efgc := _efg.(type) {
		case _g.StartElement:
			switch _efgc.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0073\u0065\u004a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0073\u0065\u004a\u0063"}:
				_eaaff.BaseJc = NewCT_YAlign()
				if _fba := d.DecodeElement(_eaaff.BaseJc, &_efgc); _fba != nil {
					return _fba
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070l\u0063\u0048\u0069\u0064\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070l\u0063\u0048\u0069\u0064\u0065"}:
				_eaaff.PlcHide = NewCT_OnOff()
				if _bec := d.DecodeElement(_eaaff.PlcHide, &_efgc); _bec != nil {
					return _bec
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072S\u0070\u0052\u0075\u006c\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072S\u0070\u0052\u0075\u006c\u0065"}:
				_eaaff.RSpRule = NewCT_SpacingRule()
				if _fde := d.DecodeElement(_eaaff.RSpRule, &_efgc); _fde != nil {
					return _fde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063G\u0070\u0052\u0075\u006c\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063G\u0070\u0052\u0075\u006c\u0065"}:
				_eaaff.CGpRule = NewCT_SpacingRule()
				if _cfdg := d.DecodeElement(_eaaff.CGpRule, &_efgc); _cfdg != nil {
					return _cfdg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0053\u0070"}:
				_eaaff.RSp = NewCT_UnSignedInteger()
				if _bgbd := d.DecodeElement(_eaaff.RSp, &_efgc); _bgbd != nil {
					return _bgbd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0053\u0070"}:
				_eaaff.CSp = NewCT_UnSignedInteger()
				if _eeaa := d.DecodeElement(_eaaff.CSp, &_efgc); _eeaa != nil {
					return _eeaa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0047\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0047\u0070"}:
				_eaaff.CGp = NewCT_UnSignedInteger()
				if _bgg := d.DecodeElement(_eaaff.CGp, &_efgc); _bgg != nil {
					return _bgg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0063\u0073"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0063\u0073"}:
				_eaaff.Mcs = NewCT_MCS()
				if _efae := d.DecodeElement(_eaaff.Mcs, &_efgc); _efae != nil {
					return _efae
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_eaaff.CtrlPr = NewCT_CtrlPr()
				if _fcfg := d.DecodeElement(_eaaff.CtrlPr, &_efgc); _fcfg != nil {
					return _fcfg
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004d\u0050\u0072\u0020\u0025\u0076", _efgc.Name)
				if _aagb := d.Skip(); _aagb != nil {
					return _aagb
				}
			}
		case _g.EndElement:
			break _bfeb
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_YAlign and its children, prefixing error messages with path
func (_bacg *CT_YAlign) ValidateWithPath(path string) error {
	if _bacg.ValAttr == _dc.ST_YAlignUnset {
		return _d.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _bcba := _bacg.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _bcba != nil {
		return _bcba
	}
	return nil
}

func (_faef *CT_OMathArgPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _faef.ArgSz != nil {
		_aadab := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0061\u0072\u0067\u0053\u007a"}}
		e.EncodeElement(_faef.ArgSz, _aadab)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_TwipsMeasure and its children, prefixing error messages with path
func (_cfbb *CT_TwipsMeasure) ValidateWithPath(path string) error {
	if _aabe := _cfbb.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _aabe != nil {
		return _aabe
	}
	return nil
}

type CT_MCPr struct {
	Count *CT_Integer255
	McJc  *CT_XAlign
}

func (_dbde *CT_OMathJc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _cdcdd := range start.Attr {
		if _cdcdd.Name.Local == "\u0076\u0061\u006c" {
			_dbde.ValAttr.UnmarshalXMLAttr(_cdcdd)
			continue
		}
	}
	for {
		_cgdf, _eceeb := d.Token()
		if _eceeb != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u004a\u0063\u003a\u0020%\u0073", _eceeb)
		}
		if _gbbe, _agfb := _cgdf.(_g.EndElement); _agfb && _gbbe.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OMathParaPr and its children
func (_eeac *CT_OMathParaPr) Validate() error {
	return _eeac.ValidateWithPath("\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0050a\u0072\u0061\u0050\u0072")
}

func (_effg *CT_OMathArg) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ccecg:
	for {
		_agee, _aegf := d.Token()
		if _aegf != nil {
			return _aegf
		}
		switch _fdeg := _agee.(type) {
		case _g.StartElement:
			switch _fdeg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0072\u0067P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0072\u0067P\u0072"}:
				_effg.ArgPr = NewCT_OMathArgPr()
				if _bgcf := d.DecodeElement(_effg.ArgPr, &_fdeg); _bgcf != nil {
					return _bgcf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_ggeb := NewEG_OMathMathElements()
				_ggeb.Acc = NewCT_Acc()
				if _dcfb := d.DecodeElement(_ggeb.Acc, &_fdeg); _dcfb != nil {
					return _dcfb
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _ggeb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_fbge := NewEG_OMathMathElements()
				_fbge.Bar = NewCT_Bar()
				if _faaf := d.DecodeElement(_fbge.Bar, &_fdeg); _faaf != nil {
					return _faaf
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _fbge)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_afaa := NewEG_OMathMathElements()
				_afaa.Box = NewCT_Box()
				if _gcdaa := d.DecodeElement(_afaa.Box, &_fdeg); _gcdaa != nil {
					return _gcdaa
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _afaa)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_cbge := NewEG_OMathMathElements()
				_cbge.BorderBox = NewCT_BorderBox()
				if _dbg := d.DecodeElement(_cbge.BorderBox, &_fdeg); _dbg != nil {
					return _dbg
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _cbge)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_ebdb := NewEG_OMathMathElements()
				_ebdb.D = NewCT_D()
				if _fdba := d.DecodeElement(_ebdb.D, &_fdeg); _fdba != nil {
					return _fdba
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _ebdb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_gdba := NewEG_OMathMathElements()
				_gdba.EqArr = NewCT_EqArr()
				if _adccb := d.DecodeElement(_gdba.EqArr, &_fdeg); _adccb != nil {
					return _adccb
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _gdba)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_gaea := NewEG_OMathMathElements()
				_gaea.F = NewCT_F()
				if _ddcb := d.DecodeElement(_gaea.F, &_fdeg); _ddcb != nil {
					return _ddcb
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _gaea)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_cdbg := NewEG_OMathMathElements()
				_cdbg.Func = NewCT_Func()
				if _agad := d.DecodeElement(_cdbg.Func, &_fdeg); _agad != nil {
					return _agad
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _cdbg)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_febe := NewEG_OMathMathElements()
				_febe.GroupChr = NewCT_GroupChr()
				if _aada := d.DecodeElement(_febe.GroupChr, &_fdeg); _aada != nil {
					return _aada
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _febe)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_bbag := NewEG_OMathMathElements()
				_bbag.LimLow = NewCT_LimLow()
				if _adbd := d.DecodeElement(_bbag.LimLow, &_fdeg); _adbd != nil {
					return _adbd
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _bbag)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_gde := NewEG_OMathMathElements()
				_gde.LimUpp = NewCT_LimUpp()
				if _cegf := d.DecodeElement(_gde.LimUpp, &_fdeg); _cegf != nil {
					return _cegf
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _gde)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_fbcc := NewEG_OMathMathElements()
				_fbcc.M = NewCT_M()
				if _cgbc := d.DecodeElement(_fbcc.M, &_fdeg); _cgbc != nil {
					return _cgbc
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _fbcc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_cfffc := NewEG_OMathMathElements()
				_cfffc.Nary = NewCT_Nary()
				if _fbea := d.DecodeElement(_cfffc.Nary, &_fdeg); _fbea != nil {
					return _fbea
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _cfffc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_fbae := NewEG_OMathMathElements()
				_fbae.Phant = NewCT_Phant()
				if _acaf := d.DecodeElement(_fbae.Phant, &_fdeg); _acaf != nil {
					return _acaf
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _fbae)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_ebcc := NewEG_OMathMathElements()
				_ebcc.Rad = NewCT_Rad()
				if _bda := d.DecodeElement(_ebcc.Rad, &_fdeg); _bda != nil {
					return _bda
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _ebcc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_aafb := NewEG_OMathMathElements()
				_aafb.SPre = NewCT_SPre()
				if _fbce := d.DecodeElement(_aafb.SPre, &_fdeg); _fbce != nil {
					return _fbce
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _aafb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_dagae := NewEG_OMathMathElements()
				_dagae.SSub = NewCT_SSub()
				if _aecf := d.DecodeElement(_dagae.SSub, &_fdeg); _aecf != nil {
					return _aecf
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _dagae)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_baag := NewEG_OMathMathElements()
				_baag.SSubSup = NewCT_SSubSup()
				if _fcaa := d.DecodeElement(_baag.SSubSup, &_fdeg); _fcaa != nil {
					return _fcaa
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _baag)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_gbff := NewEG_OMathMathElements()
				_gbff.SSup = NewCT_SSup()
				if _dggb := d.DecodeElement(_gbff.SSup, &_fdeg); _dggb != nil {
					return _dggb
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _gbff)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_eebd := NewEG_OMathMathElements()
				_eebd.R = NewCT_R()
				if _bcec := d.DecodeElement(_eebd.R, &_fdeg); _bcec != nil {
					return _bcec
				}
				_effg.EG_OMathMathElements = append(_effg.EG_OMathMathElements, _eebd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_effg.CtrlPr = NewCT_CtrlPr()
				if _gagaa := d.DecodeElement(_effg.CtrlPr, &_fdeg); _gagaa != nil {
					return _gagaa
				}
			default:
				_f.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067\u0020\u0025\u0076", _fdeg.Name)
				if _gaec := d.Skip(); _gaec != nil {
					return _gaec
				}
			}
		case _g.EndElement:
			break _ccecg
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Phant and its children, prefixing error messages with path
func (_ebba *CT_Phant) ValidateWithPath(path string) error {
	if _ebba.PhantPr != nil {
		if _eefc := _ebba.PhantPr.ValidateWithPath(path + "\u002f\u0050\u0068\u0061\u006e\u0074\u0050\u0072"); _eefc != nil {
			return _eefc
		}
	}
	if _dgec := _ebba.E.ValidateWithPath(path + "\u002f\u0045"); _dgec != nil {
		return _dgec
	}
	return nil
}

type ST_Style byte

func NewCT_Text() *CT_Text { _fcgd := &CT_Text{}; return _fcgd }

// Validate validates the CT_MCS and its children
func (_ddee *CT_MCS) Validate() error {
	return _ddee.ValidateWithPath("\u0043\u0054\u005f\u004d\u0043\u0053")
}

// Validate validates the CT_R and its children
func (_dcaa *CT_R) Validate() error { return _dcaa.ValidateWithPath("\u0043\u0054\u005f\u0052") }

// Validate validates the EG_OMathElements and its children
func (_cfcda *EG_OMathElements) Validate() error {
	return _cfcda.ValidateWithPath("\u0045\u0047_\u004f\u004d\u0061t\u0068\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073")
}

// ValidateWithPath validates the CT_UnSignedInteger and its children, prefixing error messages with path
func (_ebaac *CT_UnSignedInteger) ValidateWithPath(path string) error { return nil }

func NewCT_FuncPr() *CT_FuncPr { _aba := &CT_FuncPr{}; return _aba }

func (_agg *CT_MathPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cfba:
	for {
		_aefe, _fcea := d.Token()
		if _fcea != nil {
			return _fcea
		}
		switch _eaegg := _aefe.(type) {
		case _g.StartElement:
			switch _eaegg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}:
				_agg.MathFont = NewCT_String()
				if _gadfd := d.DecodeElement(_agg.MathFont, &_eaegg); _gadfd != nil {
					return _gadfd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0072\u006b\u0042\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0072\u006b\u0042\u0069\u006e"}:
				_agg.BrkBin = NewCT_BreakBin()
				if _fgc := d.DecodeElement(_agg.BrkBin, &_eaegg); _fgc != nil {
					return _fgc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062r\u006b\u0042\u0069\u006e\u0053\u0075b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062r\u006b\u0042\u0069\u006e\u0053\u0075b"}:
				_agg.BrkBinSub = NewCT_BreakBinSub()
				if _afdf := d.DecodeElement(_agg.BrkBinSub, &_eaegg); _afdf != nil {
					return _afdf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073m\u0061\u006c\u006c\u0046\u0072\u0061c"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073m\u0061\u006c\u006c\u0046\u0072\u0061c"}:
				_agg.SmallFrac = NewCT_OnOff()
				if _afgd := d.DecodeElement(_agg.SmallFrac, &_eaegg); _afgd != nil {
					return _afgd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064i\u0073\u0070\u0044\u0065\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064i\u0073\u0070\u0044\u0065\u0066"}:
				_agg.DispDef = NewCT_OnOff()
				if _afbd := d.DecodeElement(_agg.DispDef, &_eaegg); _afbd != nil {
					return _afbd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006cM\u0061\u0072\u0067\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006cM\u0061\u0072\u0067\u0069\u006e"}:
				_agg.LMargin = NewCT_TwipsMeasure()
				if _gcde := d.DecodeElement(_agg.LMargin, &_eaegg); _gcde != nil {
					return _gcde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072M\u0061\u0072\u0067\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072M\u0061\u0072\u0067\u0069\u006e"}:
				_agg.RMargin = NewCT_TwipsMeasure()
				if _dbcg := d.DecodeElement(_agg.RMargin, &_eaegg); _dbcg != nil {
					return _dbcg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0065\u0066J\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0065\u0066J\u0063"}:
				_agg.DefJc = NewCT_OMathJc()
				if _afde := d.DecodeElement(_agg.DefJc, &_eaegg); _afde != nil {
					return _afde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0072\u0065S\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0072\u0065S\u0070"}:
				_agg.PreSp = NewCT_TwipsMeasure()
				if _abga := d.DecodeElement(_agg.PreSp, &_eaegg); _abga != nil {
					return _abga
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u006f\u0073\u0074\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u006f\u0073\u0074\u0053\u0070"}:
				_agg.PostSp = NewCT_TwipsMeasure()
				if _ccdg := d.DecodeElement(_agg.PostSp, &_eaegg); _ccdg != nil {
					return _ccdg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069n\u0074\u0065\u0072\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069n\u0074\u0065\u0072\u0053\u0070"}:
				_agg.InterSp = NewCT_TwipsMeasure()
				if _edcee := d.DecodeElement(_agg.InterSp, &_eaegg); _edcee != nil {
					return _edcee
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069n\u0074\u0072\u0061\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069n\u0074\u0072\u0061\u0053\u0070"}:
				_agg.IntraSp = NewCT_TwipsMeasure()
				if _gfde := d.DecodeElement(_agg.IntraSp, &_eaegg); _gfde != nil {
					return _gfde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}:
				_agg.Choice = NewCT_MathPrChoice()
				if _fef := d.DecodeElement(&_agg.Choice.WrapIndent, &_eaegg); _fef != nil {
					return _fef
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}:
				_agg.Choice = NewCT_MathPrChoice()
				if _acbfe := d.DecodeElement(&_agg.Choice.WrapRight, &_eaegg); _acbfe != nil {
					return _acbfe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069\u006e\u0074\u004c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069\u006e\u0074\u004c\u0069\u006d"}:
				_agg.IntLim = NewCT_LimLoc()
				if _bebf := d.DecodeElement(_agg.IntLim, &_eaegg); _bebf != nil {
					return _bebf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006ea\u0072\u0079\u004c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006ea\u0072\u0079\u004c\u0069\u006d"}:
				_agg.NaryLim = NewCT_LimLoc()
				if _beca := d.DecodeElement(_agg.NaryLim, &_eaegg); _beca != nil {
					return _beca
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004d\u0061\u0074h\u0050\u0072 \u0025\u0076", _eaegg.Name)
				if _adfg := d.Skip(); _adfg != nil {
					return _adfg
				}
			}
		case _g.EndElement:
			break _cfba
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Func and its children, prefixing error messages with path
func (_fada *CT_Func) ValidateWithPath(path string) error {
	if _fada.FuncPr != nil {
		if _eadg := _fada.FuncPr.ValidateWithPath(path + "\u002fF\u0075\u006e\u0063\u0050\u0072"); _eadg != nil {
			return _eadg
		}
	}
	if _ebf := _fada.FName.ValidateWithPath(path + "\u002f\u0046\u004e\u0061\u006d\u0065"); _ebf != nil {
		return _ebf
	}
	if _ggeg := _fada.E.ValidateWithPath(path + "\u002f\u0045"); _ggeg != nil {
		return _ggeg
	}
	return nil
}

func (_eaefd ST_LimLoc) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_dcgcb := _g.Attr{}
	_dcgcb.Name = name
	switch _eaefd {
	case ST_LimLocUnset:
		_dcgcb.Value = ""
	case ST_LimLocUndOvr:
		_dcgcb.Value = "\u0075\u006e\u0064\u004f\u0076\u0072"
	case ST_LimLocSubSup:
		_dcgcb.Value = "\u0073\u0075\u0062\u0053\u0075\u0070"
	}
	return _dcgcb, nil
}

type CT_SSubSup struct {
	SSubSupPr *CT_SSubSupPr
	E         *CT_OMathArg
	Sub       *CT_OMathArg
	Sup       *CT_OMathArg
}

func NewCT_OMathArg() *CT_OMathArg { _gbeeac := &CT_OMathArg{}; return _gbeeac }

func NewCT_LimUpp() *CT_LimUpp {
	_bgbb := &CT_LimUpp{}
	_bgbb.E = NewCT_OMathArg()
	_bgbb.Lim = NewCT_OMathArg()
	return _bgbb
}

func NewCT_OMathArgPr() *CT_OMathArgPr { _ebdf := &CT_OMathArgPr{}; return _ebdf }

// Validate validates the CT_BorderBox and its children
func (_caa *CT_BorderBox) Validate() error {
	return _caa.ValidateWithPath("\u0043\u0054\u005fB\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078")
}

// ValidateWithPath validates the CT_M and its children, prefixing error messages with path
func (_bcca *CT_M) ValidateWithPath(path string) error {
	if _bcca.MPr != nil {
		if _daac := _bcca.MPr.ValidateWithPath(path + "\u002f\u004d\u0050\u0072"); _daac != nil {
			return _daac
		}
	}
	for _aeaf, _edea := range _bcca.Mr {
		if _aece := _edea.ValidateWithPath(_d.Sprintf("\u0025s\u002f\u004d\u0072\u005b\u0025\u0064]", path, _aeaf)); _aece != nil {
			return _aece
		}
	}
	return nil
}

func (_geeb *CT_String) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _geag := range start.Attr {
		if _geag.Name.Local == "\u0076\u0061\u006c" {
			_adccfga, _ffeg := _geag.Value, error(nil)
			if _ffeg != nil {
				return _ffeg
			}
			_geeb.ValAttr = &_adccfga
			continue
		}
	}
	for {
		_ddae, _dbae := d.Token()
		if _dbae != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0074\u0072i\u006e\u0067\u003a\u0020\u0025\u0073", _dbae)
		}
		if _ebdc, _fcgb := _ddae.(_g.EndElement); _fcgb && _ebdc.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OnOff and its children
func (_fdfd *CT_OnOff) Validate() error {
	return _fdfd.ValidateWithPath("\u0043\u0054\u005f\u004f\u006e\u004f\u0066\u0066")
}

func NewCT_MCS() *CT_MCS { _cce := &CT_MCS{}; return _cce }

type CT_RChoice struct{ T []*CT_Text }

// ValidateWithPath validates the CT_Style and its children, prefixing error messages with path
func (_agc *CT_Style) ValidateWithPath(path string) error {
	if _egca := _agc.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _egca != nil {
		return _egca
	}
	return nil
}

func (_ecca ST_FType) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_ecca.String(), start)
}

func (_facg *CT_Nary) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_facg.Sub = NewCT_OMathArg()
	_facg.Sup = NewCT_OMathArg()
	_facg.E = NewCT_OMathArg()
_bebe:
	for {
		_ecaec, _afca := d.Token()
		if _afca != nil {
			return _afca
		}
		switch _dfb := _ecaec.(type) {
		case _g.StartElement:
			switch _dfb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079\u0050\u0072"}:
				_facg.NaryPr = NewCT_NaryPr()
				if _cga := d.DecodeElement(_facg.NaryPr, &_dfb); _cga != nil {
					return _cga
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0062"}:
				if _efe := d.DecodeElement(_facg.Sub, &_dfb); _efe != nil {
					return _efe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0070"}:
				if _bfba := d.DecodeElement(_facg.Sup, &_dfb); _bfba != nil {
					return _bfba
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _aaae := d.DecodeElement(_facg.E, &_dfb); _aaae != nil {
					return _aaae
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u004e\u0061\u0072\u0079\u0020\u0025\u0076", _dfb.Name)
				if _fecb := d.Skip(); _fecb != nil {
					return _fecb
				}
			}
		case _g.EndElement:
			break _bebe
		case _g.CharData:
		}
	}
	return nil
}

func NewCT_Bar() *CT_Bar { _ae := &CT_Bar{}; _ae.E = NewCT_OMathArg(); return _ae }

func (_abg *CT_MCS) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gdab:
	for {
		_aage, _eeg := d.Token()
		if _eeg != nil {
			return _eeg
		}
		switch _ggbd := _aage.(type) {
		case _g.StartElement:
			switch _ggbd.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0063"}:
				_dgef := NewCT_MC()
				if _cdac := d.DecodeElement(_dgef, &_ggbd); _cdac != nil {
					return _cdac
				}
				_abg.Mc = append(_abg.Mc, _dgef)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004d\u0043\u0053\u0020\u0025\u0076", _ggbd.Name)
				if _bba := d.Skip(); _bba != nil {
					return _bba
				}
			}
		case _g.EndElement:
			break _gdab
		case _g.CharData:
		}
	}
	return nil
}

type EG_OMathElements struct{ EG_OMathMathElements []*EG_OMathMathElements }

func (_ababf *ST_Script) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_ababf = 0
	case "\u0072\u006f\u006da\u006e":
		*_ababf = 1
	case "\u0073\u0063\u0072\u0069\u0070\u0074":
		*_ababf = 2
	case "\u0066r\u0061\u006b\u0074\u0075\u0072":
		*_ababf = 3
	case "\u0064\u006f\u0075\u0062\u006c\u0065\u002d\u0073\u0074\u0072\u0075\u0063\u006b":
		*_ababf = 4
	case "\u0073\u0061\u006e\u0073\u002d\u0073\u0065\u0072\u0069\u0066":
		*_ababf = 5
	case "\u006do\u006e\u006f\u0073\u0070\u0061\u0063e":
		*_ababf = 6
	}
	return nil
}

// Validate validates the OMathPara and its children
func (_afgg *OMathPara) Validate() error {
	return _afgg.ValidateWithPath("\u004fM\u0061\u0074\u0068\u0050\u0061\u0072a")
}

// ValidateWithPath validates the CT_SPrePr and its children, prefixing error messages with path
func (_dcec *CT_SPrePr) ValidateWithPath(path string) error {
	if _dcec.CtrlPr != nil {
		if _cgda := _dcec.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _cgda != nil {
			return _cgda
		}
	}
	return nil
}

func NewCT_SpacingRule() *CT_SpacingRule { _ceeg := &CT_SpacingRule{}; _ceeg.ValAttr = 0; return _ceeg }

func (_gca *CT_D) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _gca.DPr != nil {
		_gegg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064P\u0072"}}
		e.EncodeElement(_gca.DPr, _gegg)
	}
	_ffd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	for _, _cefg := range _gca.E {
		e.EncodeElement(_cefg, _ffd)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_dfa *CT_Box) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dfa.BoxPr != nil {
		_cgb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0062\u006f\u0078\u0050\u0072"}}
		e.EncodeElement(_dfa.BoxPr, _cgb)
	}
	_bedd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_dfa.E, _bedd)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_dacfd *CT_ManualBreak) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _dacfd.AlnAtAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d:\u0061\u006c\u006e\u0041\u0074"}, Value: _d.Sprintf("\u0025\u0076", *_dacfd.AlnAtAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_Phant() *CT_Phant { _cbbfc := &CT_Phant{}; _cbbfc.E = NewCT_OMathArg(); return _cbbfc }

func NewCT_FPr() *CT_FPr { _agd := &CT_FPr{}; return _agd }

func (_ddgea *EG_OMathElements) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_eecdc:
	for {
		_ebgc, _abfgc := d.Token()
		if _abfgc != nil {
			return _abfgc
		}
		switch _bceg := _ebgc.(type) {
		case _g.StartElement:
			switch _bceg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_bbca := NewEG_OMathMathElements()
				_bbca.Acc = NewCT_Acc()
				if _eaaa := d.DecodeElement(_bbca.Acc, &_bceg); _eaaa != nil {
					return _eaaa
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _bbca)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_beeea := NewEG_OMathMathElements()
				_beeea.Bar = NewCT_Bar()
				if _fbfdb := d.DecodeElement(_beeea.Bar, &_bceg); _fbfdb != nil {
					return _fbfdb
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _beeea)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_cgaf := NewEG_OMathMathElements()
				_cgaf.Box = NewCT_Box()
				if _fcca := d.DecodeElement(_cgaf.Box, &_bceg); _fcca != nil {
					return _fcca
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _cgaf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_bgabd := NewEG_OMathMathElements()
				_bgabd.BorderBox = NewCT_BorderBox()
				if _egedg := d.DecodeElement(_bgabd.BorderBox, &_bceg); _egedg != nil {
					return _egedg
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _bgabd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_gebde := NewEG_OMathMathElements()
				_gebde.D = NewCT_D()
				if _dgga := d.DecodeElement(_gebde.D, &_bceg); _dgga != nil {
					return _dgga
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _gebde)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_eddf := NewEG_OMathMathElements()
				_eddf.EqArr = NewCT_EqArr()
				if _acdg := d.DecodeElement(_eddf.EqArr, &_bceg); _acdg != nil {
					return _acdg
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _eddf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_geddd := NewEG_OMathMathElements()
				_geddd.F = NewCT_F()
				if _baeag := d.DecodeElement(_geddd.F, &_bceg); _baeag != nil {
					return _baeag
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _geddd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_bgca := NewEG_OMathMathElements()
				_bgca.Func = NewCT_Func()
				if _aceg := d.DecodeElement(_bgca.Func, &_bceg); _aceg != nil {
					return _aceg
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _bgca)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_dabb := NewEG_OMathMathElements()
				_dabb.GroupChr = NewCT_GroupChr()
				if _ececc := d.DecodeElement(_dabb.GroupChr, &_bceg); _ececc != nil {
					return _ececc
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _dabb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_cffgd := NewEG_OMathMathElements()
				_cffgd.LimLow = NewCT_LimLow()
				if _fgaa := d.DecodeElement(_cffgd.LimLow, &_bceg); _fgaa != nil {
					return _fgaa
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _cffgd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_gfbc := NewEG_OMathMathElements()
				_gfbc.LimUpp = NewCT_LimUpp()
				if _ecgg := d.DecodeElement(_gfbc.LimUpp, &_bceg); _ecgg != nil {
					return _ecgg
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _gfbc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_dfcb := NewEG_OMathMathElements()
				_dfcb.M = NewCT_M()
				if _fegcf := d.DecodeElement(_dfcb.M, &_bceg); _fegcf != nil {
					return _fegcf
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _dfcb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_gfege := NewEG_OMathMathElements()
				_gfege.Nary = NewCT_Nary()
				if _fbbc := d.DecodeElement(_gfege.Nary, &_bceg); _fbbc != nil {
					return _fbbc
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _gfege)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_ffgaf := NewEG_OMathMathElements()
				_ffgaf.Phant = NewCT_Phant()
				if _dbdeb := d.DecodeElement(_ffgaf.Phant, &_bceg); _dbdeb != nil {
					return _dbdeb
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _ffgaf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_dfbc := NewEG_OMathMathElements()
				_dfbc.Rad = NewCT_Rad()
				if _cbfc := d.DecodeElement(_dfbc.Rad, &_bceg); _cbfc != nil {
					return _cbfc
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _dfbc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_adbeb := NewEG_OMathMathElements()
				_adbeb.SPre = NewCT_SPre()
				if _fefg := d.DecodeElement(_adbeb.SPre, &_bceg); _fefg != nil {
					return _fefg
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _adbeb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_eefa := NewEG_OMathMathElements()
				_eefa.SSub = NewCT_SSub()
				if _cdbe := d.DecodeElement(_eefa.SSub, &_bceg); _cdbe != nil {
					return _cdbe
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _eefa)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_accfd := NewEG_OMathMathElements()
				_accfd.SSubSup = NewCT_SSubSup()
				if _gbgd := d.DecodeElement(_accfd.SSubSup, &_bceg); _gbgd != nil {
					return _gbgd
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _accfd)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_dcdf := NewEG_OMathMathElements()
				_dcdf.SSup = NewCT_SSup()
				if _ecdae := d.DecodeElement(_dcdf.SSup, &_bceg); _ecdae != nil {
					return _ecdae
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _dcdf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_edef := NewEG_OMathMathElements()
				_edef.R = NewCT_R()
				if _fdacc := d.DecodeElement(_edef.R, &_bceg); _fdacc != nil {
					return _fdacc
				}
				_ddgea.EG_OMathMathElements = append(_ddgea.EG_OMathMathElements, _edef)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u0047\u005f\u004f\u004d\u0061\u0074\u0068\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0020\u0025v", _bceg.Name)
				if _eccf := d.Skip(); _eccf != nil {
					return _eccf
				}
			}
		case _g.EndElement:
			break _eecdc
		case _g.CharData:
		}
	}
	return nil
}

func (_gdgdf ST_Shp) Validate() error { return _gdgdf.ValidateWithPath("") }

type CT_OMathArg struct {
	ArgPr                *CT_OMathArgPr
	EG_OMathMathElements []*EG_OMathMathElements
	CtrlPr               *CT_CtrlPr
}

func (_afga *CT_BreakBinSub) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _cac := range start.Attr {
		if _cac.Name.Local == "\u0076\u0061\u006c" {
			_afga.ValAttr.UnmarshalXMLAttr(_cac)
			continue
		}
	}
	for {
		_eae, _bcd := d.Token()
		if _bcd != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fB\u0072\u0065\u0061\u006b\u0042\u0069\u006e\u0053\u0075\u0062:\u0020\u0025\u0073", _bcd)
		}
		if _fagc, _ded := _eae.(_g.EndElement); _ded && _fagc.Name == start.Name {
			break
		}
	}
	return nil
}

func NewOMathPara() *OMathPara {
	_faag := &OMathPara{}
	_faag.CT_OMathPara = *NewCT_OMathPara()
	return _faag
}

func NewCT_SSubSup() *CT_SSubSup {
	_bdebg := &CT_SSubSup{}
	_bdebg.E = NewCT_OMathArg()
	_bdebg.Sub = NewCT_OMathArg()
	_bdebg.Sup = NewCT_OMathArg()
	return _bdebg
}

// ValidateWithPath validates the CT_Nary and its children, prefixing error messages with path
func (_cdg *CT_Nary) ValidateWithPath(path string) error {
	if _cdg.NaryPr != nil {
		if _cegb := _cdg.NaryPr.ValidateWithPath(path + "\u002fN\u0061\u0072\u0079\u0050\u0072"); _cegb != nil {
			return _cegb
		}
	}
	if _gbaf := _cdg.Sub.ValidateWithPath(path + "\u002f\u0053\u0075\u0062"); _gbaf != nil {
		return _gbaf
	}
	if _ceecg := _cdg.Sup.ValidateWithPath(path + "\u002f\u0053\u0075\u0070"); _ceecg != nil {
		return _ceecg
	}
	if _eggb := _cdg.E.ValidateWithPath(path + "\u002f\u0045"); _eggb != nil {
		return _eggb
	}
	return nil
}

// Validate validates the CT_FPr and its children
func (_affd *CT_FPr) Validate() error {
	return _affd.ValidateWithPath("\u0043\u0054\u005f\u0046\u0050\u0072")
}

func (_abdff *CT_RadPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _abdff.DegHide != nil {
		_cgbe := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0064\u0065\u0067\u0048\u0069\u0064e"}}
		e.EncodeElement(_abdff.DegHide, _cgbe)
	}
	if _abdff.CtrlPr != nil {
		_aeab := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_abdff.CtrlPr, _aeab)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_bcgdg *ST_BreakBinSub) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_bcgdg = 0
	case "\u002d\u002d":
		*_bcgdg = 1
	case "\u002d\u002b":
		*_bcgdg = 2
	case "\u002b\u002d":
		*_bcgdg = 3
	}
	return nil
}

func NewCT_AccPr() *CT_AccPr { _ad := &CT_AccPr{}; return _ad }

func (_cfb *CT_BoxPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cfb.OpEmu != nil {
		_add := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006f\u0070\u0045\u006d\u0075"}}
		e.EncodeElement(_cfb.OpEmu, _add)
	}
	if _cfb.NoBreak != nil {
		_aaa := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006e\u006f\u0042\u0072\u0065\u0061k"}}
		e.EncodeElement(_cfb.NoBreak, _aaa)
	}
	if _cfb.Diff != nil {
		_fac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064\u0069\u0066\u0066"}}
		e.EncodeElement(_cfb.Diff, _fac)
	}
	if _cfb.Brk != nil {
		_cgcc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062r\u006b"}}
		e.EncodeElement(_cfb.Brk, _cgcc)
	}
	if _cfb.Aln != nil {
		_gad := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0061l\u006e"}}
		e.EncodeElement(_cfb.Aln, _gad)
	}
	if _cfb.CtrlPr != nil {
		_ffa := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_cfb.CtrlPr, _ffa)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_dfga *ST_FType) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_dfga = 0
	case "\u0062\u0061\u0072":
		*_dfga = 1
	case "\u0073\u006b\u0077":
		*_dfga = 2
	case "\u006c\u0069\u006e":
		*_dfga = 3
	case "\u006e\u006f\u0042a\u0072":
		*_dfga = 4
	}
	return nil
}

func NewCT_SPre() *CT_SPre {
	_dbfa := &CT_SPre{}
	_dbfa.Sub = NewCT_OMathArg()
	_dbfa.Sup = NewCT_OMathArg()
	_dbfa.E = NewCT_OMathArg()
	return _dbfa
}

// ValidateWithPath validates the CT_TopBot and its children, prefixing error messages with path
func (_cgdb *CT_TopBot) ValidateWithPath(path string) error {
	if _cgdb.ValAttr == ST_TopBotUnset {
		return _d.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _fcgg := _cgdb.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _fcgg != nil {
		return _fcgg
	}
	return nil
}

func NewCT_Box() *CT_Box { _gbe := &CT_Box{}; _gbe.E = NewCT_OMathArg(); return _gbe }

type CT_AccPr struct {
	Chr    *CT_Char
	CtrlPr *CT_CtrlPr
}

type CT_XAlign struct{ ValAttr _dc.ST_XAlign }

func (_ddd *CT_Char) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _d.Sprintf("\u0025\u0076", _ddd.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_Acc struct {
	AccPr *CT_AccPr
	E     *CT_OMathArg
}

type CT_RPRChoice struct{ Nor *CT_OnOff }

// ValidateWithPath validates the MathPr and its children, prefixing error messages with path
func (_fdff *MathPr) ValidateWithPath(path string) error {
	if _ccge := _fdff.CT_MathPr.ValidateWithPath(path); _ccge != nil {
		return _ccge
	}
	return nil
}

const (
	ST_FTypeUnset ST_FType = 0
	ST_FTypeBar   ST_FType = 1
	ST_FTypeSkw   ST_FType = 2
	ST_FTypeLin   ST_FType = 3
	ST_FTypeNoBar ST_FType = 4
)

// Validate validates the CT_Integer255 and its children
func (_dce *CT_Integer255) Validate() error {
	return _dce.ValidateWithPath("\u0043\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032\u0035\u0035")
}

type CT_UnSignedInteger struct{ ValAttr uint32 }

type CT_MC struct{ McPr *CT_MCPr }

// ValidateWithPath validates the CT_MPr and its children, prefixing error messages with path
func (_ccee *CT_MPr) ValidateWithPath(path string) error {
	if _ccee.BaseJc != nil {
		if _cbgg := _ccee.BaseJc.ValidateWithPath(path + "\u002fB\u0061\u0073\u0065\u004a\u0063"); _cbgg != nil {
			return _cbgg
		}
	}
	if _ccee.PlcHide != nil {
		if _eccg := _ccee.PlcHide.ValidateWithPath(path + "\u002f\u0050\u006c\u0063\u0048\u0069\u0064\u0065"); _eccg != nil {
			return _eccg
		}
	}
	if _ccee.RSpRule != nil {
		if _adef := _ccee.RSpRule.ValidateWithPath(path + "\u002f\u0052\u0053\u0070\u0052\u0075\u006c\u0065"); _adef != nil {
			return _adef
		}
	}
	if _ccee.CGpRule != nil {
		if _faegc := _ccee.CGpRule.ValidateWithPath(path + "\u002f\u0043\u0047\u0070\u0052\u0075\u006c\u0065"); _faegc != nil {
			return _faegc
		}
	}
	if _ccee.RSp != nil {
		if _dgeg := _ccee.RSp.ValidateWithPath(path + "\u002f\u0052\u0053\u0070"); _dgeg != nil {
			return _dgeg
		}
	}
	if _ccee.CSp != nil {
		if _agdc := _ccee.CSp.ValidateWithPath(path + "\u002f\u0043\u0053\u0070"); _agdc != nil {
			return _agdc
		}
	}
	if _ccee.CGp != nil {
		if _gecb := _ccee.CGp.ValidateWithPath(path + "\u002f\u0043\u0047\u0070"); _gecb != nil {
			return _gecb
		}
	}
	if _ccee.Mcs != nil {
		if _bbcd := _ccee.Mcs.ValidateWithPath(path + "\u002f\u004d\u0063\u0073"); _bbcd != nil {
			return _bbcd
		}
	}
	if _ccee.CtrlPr != nil {
		if _dfda := _ccee.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _dfda != nil {
			return _dfda
		}
	}
	return nil
}

// Validate validates the CT_Phant and its children
func (_bdd *CT_Phant) Validate() error {
	return _bdd.ValidateWithPath("\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074")
}

type CT_MCS struct{ Mc []*CT_MC }

// Validate validates the CT_SSup and its children
func (_edga *CT_SSup) Validate() error {
	return _edga.ValidateWithPath("\u0043T\u005f\u0053\u0053\u0075\u0070")
}

func NewCT_TopBot() *CT_TopBot { _cedc := &CT_TopBot{}; _cedc.ValAttr = ST_TopBot(1); return _cedc }

type CT_LimUppPr struct{ CtrlPr *CT_CtrlPr }

type CT_SPre struct {
	SPrePr *CT_SPrePr
	Sub    *CT_OMathArg
	Sup    *CT_OMathArg
	E      *CT_OMathArg
}

// ValidateWithPath validates the CT_CtrlPr and its children, prefixing error messages with path
func (_cef *CT_CtrlPr) ValidateWithPath(path string) error { return nil }

type CT_BorderBox struct {
	BorderBoxPr *CT_BorderBoxPr
	E           *CT_OMathArg
}

const (
	ST_TopBotUnset ST_TopBot = 0
	ST_TopBotTop   ST_TopBot = 1
	ST_TopBotBot   ST_TopBot = 2
)

type CT_Phant struct {
	PhantPr *CT_PhantPr
	E       *CT_OMathArg
}

func (_dbdg *CT_OMathPara) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dbdg.OMathParaPr != nil {
		_efgg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006f\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061\u0050\u0072"}}
		e.EncodeElement(_dbdg.OMathParaPr, _efgg)
	}
	_cgab := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006f\u004d\u0061\u0074\u0068"}}
	for _, _dfaa := range _dbdg.OMath {
		e.EncodeElement(_dfaa, _cgab)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the EG_OMathMathElements and its children
func (_eafe *EG_OMathMathElements) Validate() error {
	return _eafe.ValidateWithPath("E\u0047_\u004f\u004d\u0061\u0074\u0068\u004d\u0061\u0074h\u0045\u006c\u0065\u006den\u0074\u0073")
}

func (_aaea ST_LimLoc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_aaea.String(), start)
}

func (_cadd *CT_FPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gcdc:
	for {
		_cdcbg, _beda := d.Token()
		if _beda != nil {
			return _beda
		}
		switch _ggdec := _cdcbg.(type) {
		case _g.StartElement:
			switch _ggdec.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0074\u0079\u0070\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0074\u0079\u0070\u0065"}:
				_cadd.Type = NewCT_FType()
				if _fad := d.DecodeElement(_cadd.Type, &_ggdec); _fad != nil {
					return _fad
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_cadd.CtrlPr = NewCT_CtrlPr()
				if _dec := d.DecodeElement(_cadd.CtrlPr, &_ggdec); _dec != nil {
					return _dec
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0046\u0050\u0072\u0020\u0025\u0076", _ggdec.Name)
				if _eba := d.Skip(); _eba != nil {
					return _eba
				}
			}
		case _g.EndElement:
			break _gcdc
		case _g.CharData:
		}
	}
	return nil
}

const (
	ST_BreakBinUnset  ST_BreakBin = 0
	ST_BreakBinBefore ST_BreakBin = 1
	ST_BreakBinAfter  ST_BreakBin = 2
	ST_BreakBinRepeat ST_BreakBin = 3
)

func (_ccdf *CT_M) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ccdf.MPr != nil {
		_aabb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006dP\u0072"}}
		e.EncodeElement(_ccdf.MPr, _aabb)
	}
	_cadfe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0072"}}
	for _, _acda := range _ccdf.Mr {
		e.EncodeElement(_acda, _cadfe)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_M struct {
	MPr *CT_MPr
	Mr  []*CT_MR
}

// ValidateWithPath validates the CT_OnOff and its children, prefixing error messages with path
func (_fdcb *CT_OnOff) ValidateWithPath(path string) error {
	if _fdcb.ValAttr != nil {
		if _agac := _fdcb.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _agac != nil {
			return _agac
		}
	}
	return nil
}

type CT_BorderBoxPr struct {
	HideTop    *CT_OnOff
	HideBot    *CT_OnOff
	HideLeft   *CT_OnOff
	HideRight  *CT_OnOff
	StrikeH    *CT_OnOff
	StrikeV    *CT_OnOff
	StrikeBLTR *CT_OnOff
	StrikeTLBR *CT_OnOff
	CtrlPr     *CT_CtrlPr
}

const (
	ST_ScriptUnset         ST_Script = 0
	ST_ScriptRoman         ST_Script = 1
	ST_ScriptScript        ST_Script = 2
	ST_ScriptFraktur       ST_Script = 3
	ST_ScriptDouble_struck ST_Script = 4
	ST_ScriptSans_serif    ST_Script = 5
	ST_ScriptMonospace     ST_Script = 6
)

func (_dbga ST_Jc) ValidateWithPath(path string) error {
	switch _dbga {
	case 0, 1, 2, 3, 4:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dbga))
	}
	return nil
}

func (_dbeag *CT_GroupChrPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cgcbb:
	for {
		_agdb, _gee := d.Token()
		if _gee != nil {
			return _gee
		}
		switch _gdfe := _agdb.(type) {
		case _g.StartElement:
			switch _gdfe.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0068\u0072"}:
				_dbeag.Chr = NewCT_Char()
				if _acdc := d.DecodeElement(_dbeag.Chr, &_gdfe); _acdc != nil {
					return _acdc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u006f\u0073"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u006f\u0073"}:
				_dbeag.Pos = NewCT_TopBot()
				if _bbe := d.DecodeElement(_dbeag.Pos, &_gdfe); _bbe != nil {
					return _bbe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0076\u0065\u0072\u0074\u004a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0076\u0065\u0072\u0074\u004a\u0063"}:
				_dbeag.VertJc = NewCT_TopBot()
				if _gfeg := d.DecodeElement(_dbeag.VertJc, &_gdfe); _gfeg != nil {
					return _gfeg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_dbeag.CtrlPr = NewCT_CtrlPr()
				if _fbc := d.DecodeElement(_dbeag.CtrlPr, &_gdfe); _fbc != nil {
					return _fbc
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047r\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072 \u0025\u0076", _gdfe.Name)
				if _bfcf := d.Skip(); _bfcf != nil {
					return _bfcf
				}
			}
		case _g.EndElement:
			break _cgcbb
		case _g.CharData:
		}
	}
	return nil
}

type OMathPara struct{ CT_OMathPara }

// Validate validates the CT_Style and its children
func (_gagg *CT_Style) Validate() error {
	return _gagg.ValidateWithPath("\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065")
}

// Validate validates the CT_BorderBoxPr and its children
func (_dgg *CT_BorderBoxPr) Validate() error {
	return _dgg.ValidateWithPath("\u0043\u0054\u005f\u0042\u006f\u0072\u0064\u0065\u0072B\u006f\u0078\u0050\u0072")
}

// Validate validates the OMath and its children
func (_fbead *OMath) Validate() error { return _fbead.ValidateWithPath("\u004f\u004d\u0061t\u0068") }

func (_fecf ST_TopBot) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_fecf.String(), start)
}

func (_gfga *CT_LimLoc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_ecda, _aacc := _gfga.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _aacc != nil {
		return _aacc
	}
	start.Attr = append(start.Attr, _ecda)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_faege ST_Jc) String() string {
	switch _faege {
	case 0:
		return ""
	case 1:
		return "\u006c\u0065\u0066\u0074"
	case 2:
		return "\u0072\u0069\u0067h\u0074"
	case 3:
		return "\u0063\u0065\u006e\u0074\u0065\u0072"
	case 4:
		return "c\u0065\u006e\u0074\u0065\u0072\u0047\u0072\u006f\u0075\u0070"
	}
	return ""
}

func (_ccbgc ST_Shp) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_gabg := _g.Attr{}
	_gabg.Name = name
	switch _ccbgc {
	case ST_ShpUnset:
		_gabg.Value = ""
	case ST_ShpCentered:
		_gabg.Value = "\u0063\u0065\u006e\u0074\u0065\u0072\u0065\u0064"
	case ST_ShpMatch:
		_gabg.Value = "\u006d\u0061\u0074c\u0068"
	}
	return _gabg, nil
}

// ValidateWithPath validates the CT_Integer2 and its children, prefixing error messages with path
func (_eaeb *CT_Integer2) ValidateWithPath(path string) error {
	if _eaeb.ValAttr < -2 {
		return _d.Errorf("\u0025\u0073/m\u002e\u0056\u0061l\u0041\u0074\u0074\u0072 mu\u0073t \u0062\u0065\u0020\u003e\u003d\u0020\u002d2 \u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _eaeb.ValAttr)
	}
	if _eaeb.ValAttr > 2 {
		return _d.Errorf("%\u0073\u002f\u006d\u002e\u0056\u0061l\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003c=\u0020\u0032\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _eaeb.ValAttr)
	}
	return nil
}

func (_bgfe ST_LimLoc) String() string {
	switch _bgfe {
	case 0:
		return ""
	case 1:
		return "\u0075\u006e\u0064\u004f\u0076\u0072"
	case 2:
		return "\u0073\u0075\u0062\u0053\u0075\u0070"
	}
	return ""
}

func (_edeaf *CT_MC) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_bff:
	for {
		_eda, _bgba := d.Token()
		if _bgba != nil {
			return _bgba
		}
		switch _dgfb := _eda.(type) {
		case _g.StartElement:
			switch _dgfb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0063\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0063\u0050\u0072"}:
				_edeaf.McPr = NewCT_MCPr()
				if _edbb := d.DecodeElement(_edeaf.McPr, &_dgfb); _edbb != nil {
					return _edbb
				}
			default:
				_f.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0043\u0054\u005fM\u0043 \u0025\u0076", _dgfb.Name)
				if _cbg := d.Skip(); _cbg != nil {
					return _cbg
				}
			}
		case _g.EndElement:
			break _bff
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SSubSup and its children
func (_gcbd *CT_SSubSup) Validate() error {
	return _gcbd.ValidateWithPath("\u0043\u0054\u005f\u0053\u0053\u0075\u0062\u0053\u0075\u0070")
}

func (_eeec *OMath) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006d"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0073\u0063\u0068\u0065m\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0068\u0061\u0072e\u0064\u0054\u0079\u0070\u0065\u0073"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0077"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065s\u0073i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u00306\u002fm\u0061\u0069n"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006d:\u006f\u004d\u0061\u0074\u0068"
	return _eeec.CT_OMath.MarshalXML(e, start)
}

// Validate validates the CT_TwipsMeasure and its children
func (_degb *CT_TwipsMeasure) Validate() error {
	return _degb.ValidateWithPath("\u0043T\u005fT\u0077\u0069\u0070\u0073\u004d\u0065\u0061\u0073\u0075\u0072\u0065")
}

type CT_OMathParaPr struct{ Jc *CT_OMathJc }

func (_fagca ST_BreakBin) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_gfbd := _g.Attr{}
	_gfbd.Name = name
	switch _fagca {
	case ST_BreakBinUnset:
		_gfbd.Value = ""
	case ST_BreakBinBefore:
		_gfbd.Value = "\u0062\u0065\u0066\u006f\u0072\u0065"
	case ST_BreakBinAfter:
		_gfbd.Value = "\u0061\u0066\u0074e\u0072"
	case ST_BreakBinRepeat:
		_gfbd.Value = "\u0072\u0065\u0070\u0065\u0061\u0074"
	}
	return _gfbd, nil
}

func NewCT_OMathPara() *CT_OMathPara { _fgg := &CT_OMathPara{}; return _fgg }

// Validate validates the CT_RPR and its children
func (_feca *CT_RPR) Validate() error {
	return _feca.ValidateWithPath("\u0043\u0054\u005f\u0052\u0050\u0052")
}

// Validate validates the CT_FuncPr and its children
func (_eeag *CT_FuncPr) Validate() error {
	return _eeag.ValidateWithPath("\u0043T\u005f\u0046\u0075\u006e\u0063\u0050r")
}

type CT_TopBot struct{ ValAttr ST_TopBot }

func NewCT_Char() *CT_Char { _beb := &CT_Char{}; return _beb }

func (_eec *CT_LimUppPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _eec.CtrlPr != nil {
		_bbb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_eec.CtrlPr, _bbb)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_EqArrPr and its children, prefixing error messages with path
func (_aag *CT_EqArrPr) ValidateWithPath(path string) error {
	if _aag.BaseJc != nil {
		if _bgd := _aag.BaseJc.ValidateWithPath(path + "\u002fB\u0061\u0073\u0065\u004a\u0063"); _bgd != nil {
			return _bgd
		}
	}
	if _aag.MaxDist != nil {
		if _ddfgg := _aag.MaxDist.ValidateWithPath(path + "\u002f\u004d\u0061\u0078\u0044\u0069\u0073\u0074"); _ddfgg != nil {
			return _ddfgg
		}
	}
	if _aag.ObjDist != nil {
		if _baa := _aag.ObjDist.ValidateWithPath(path + "\u002f\u004f\u0062\u006a\u0044\u0069\u0073\u0074"); _baa != nil {
			return _baa
		}
	}
	if _aag.RSpRule != nil {
		if _cca := _aag.RSpRule.ValidateWithPath(path + "\u002f\u0052\u0053\u0070\u0052\u0075\u006c\u0065"); _cca != nil {
			return _cca
		}
	}
	if _aag.RSp != nil {
		if _dafc := _aag.RSp.ValidateWithPath(path + "\u002f\u0052\u0053\u0070"); _dafc != nil {
			return _dafc
		}
	}
	if _aag.CtrlPr != nil {
		if _bfda := _aag.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _bfda != nil {
			return _bfda
		}
	}
	return nil
}

// ValidateWithPath validates the CT_FuncPr and its children, prefixing error messages with path
func (_ebd *CT_FuncPr) ValidateWithPath(path string) error {
	if _ebd.CtrlPr != nil {
		if _fec := _ebd.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _fec != nil {
			return _fec
		}
	}
	return nil
}

func (_fd *CT_Acc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_fd.E = NewCT_OMathArg()
_ge:
	for {
		_fdd, _e := d.Token()
		if _e != nil {
			return _e
		}
		switch _a := _fdd.(type) {
		case _g.StartElement:
			switch _a.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063P\u0072"}:
				_fd.AccPr = NewCT_AccPr()
				if _dcb := d.DecodeElement(_fd.AccPr, &_a); _dcb != nil {
					return _dcb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _cd := d.DecodeElement(_fd.E, &_a); _cd != nil {
					return _cd
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0041\u0063\u0063\u0020\u0025\u0076", _a.Name)
				if _ab := d.Skip(); _ab != nil {
					return _ab
				}
			}
		case _g.EndElement:
			break _ge
		case _g.CharData:
		}
	}
	return nil
}

const (
	ST_LimLocUnset  ST_LimLoc = 0
	ST_LimLocUndOvr ST_LimLoc = 1
	ST_LimLocSubSup ST_LimLoc = 2
)

func NewCT_TwipsMeasure() *CT_TwipsMeasure { _decad := &CT_TwipsMeasure{}; return _decad }

func NewCT_UnSignedInteger() *CT_UnSignedInteger { _fgdb := &CT_UnSignedInteger{}; return _fgdb }

type CT_RPR struct {
	Lit    *CT_OnOff
	Choice *CT_RPRChoice
	Brk    *CT_ManualBreak
	Aln    *CT_OnOff
}

func NewCT_String() *CT_String { _eadda := &CT_String{}; return _eadda }

// ValidateWithPath validates the CT_MR and its children, prefixing error messages with path
func (_bgdb *CT_MR) ValidateWithPath(path string) error {
	for _acba, _ccef := range _bgdb.E {
		if _gfc := _ccef.ValidateWithPath(_d.Sprintf("\u0025\u0073\u002f\u0045\u005b\u0025\u0064\u005d", path, _acba)); _gfc != nil {
			return _gfc
		}
	}
	return nil
}

func (_bbfc *ST_Script) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_abeg, _agda := d.Token()
	if _agda != nil {
		return _agda
	}
	if _cdbd, _eedae := _abeg.(_g.EndElement); _eedae && _cdbd.Name == start.Name {
		*_bbfc = 1
		return nil
	}
	if _acddc, _edbe := _abeg.(_g.CharData); !_edbe {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _abeg)
	} else {
		switch string(_acddc) {
		case "":
			*_bbfc = 0
		case "\u0072\u006f\u006da\u006e":
			*_bbfc = 1
		case "\u0073\u0063\u0072\u0069\u0070\u0074":
			*_bbfc = 2
		case "\u0066r\u0061\u006b\u0074\u0075\u0072":
			*_bbfc = 3
		case "\u0064\u006f\u0075\u0062\u006c\u0065\u002d\u0073\u0074\u0072\u0075\u0063\u006b":
			*_bbfc = 4
		case "\u0073\u0061\u006e\u0073\u002d\u0073\u0065\u0072\u0069\u0066":
			*_bbfc = 5
		case "\u006do\u006e\u006f\u0073\u0070\u0061\u0063e":
			*_bbfc = 6
		}
	}
	_abeg, _agda = d.Token()
	if _agda != nil {
		return _agda
	}
	if _abad, _ebfe := _abeg.(_g.EndElement); _ebfe && _abad.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _abeg)
}

type EG_ScriptStyle struct {
	Scr *CT_Script
	Sty *CT_Style
}

const (
	ST_JcUnset       ST_Jc = 0
	ST_JcLeft        ST_Jc = 1
	ST_JcRight       ST_Jc = 2
	ST_JcCenter      ST_Jc = 3
	ST_JcCenterGroup ST_Jc = 4
)

func NewCT_Func() *CT_Func {
	_edce := &CT_Func{}
	_edce.FName = NewCT_OMathArg()
	_edce.E = NewCT_OMathArg()
	return _edce
}

func (_cebde *ST_Style) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_abcbc, _cdee := d.Token()
	if _cdee != nil {
		return _cdee
	}
	if _aced, _fdec := _abcbc.(_g.EndElement); _fdec && _aced.Name == start.Name {
		*_cebde = 1
		return nil
	}
	if _adbg, _aedd := _abcbc.(_g.CharData); !_aedd {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _abcbc)
	} else {
		switch string(_adbg) {
		case "":
			*_cebde = 0
		case "\u0070":
			*_cebde = 1
		case "\u0062":
			*_cebde = 2
		case "\u0069":
			*_cebde = 3
		case "\u0062\u0069":
			*_cebde = 4
		}
	}
	_abcbc, _cdee = d.Token()
	if _cdee != nil {
		return _cdee
	}
	if _egfba, _dgbafa := _abcbc.(_g.EndElement); _dgbafa && _egfba.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _abcbc)
}

func NewCT_SSub() *CT_SSub {
	_ecbe := &CT_SSub{}
	_ecbe.E = NewCT_OMathArg()
	_ecbe.Sub = NewCT_OMathArg()
	return _ecbe
}

func (_cgdc *EG_OMathElements) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _cgdc.EG_OMathMathElements != nil {
		for _, _ecbbd := range _cgdc.EG_OMathMathElements {
			_ecbbd.MarshalXML(e, _g.StartElement{})
		}
	}
	return nil
}

func (_cgfc ST_Style) String() string {
	switch _cgfc {
	case 0:
		return ""
	case 1:
		return "\u0070"
	case 2:
		return "\u0062"
	case 3:
		return "\u0069"
	case 4:
		return "\u0062\u0069"
	}
	return ""
}

// ValidateWithPath validates the CT_MC and its children, prefixing error messages with path
func (_egfbf *CT_MC) ValidateWithPath(path string) error {
	if _egfbf.McPr != nil {
		if _cdcd := _egfbf.McPr.ValidateWithPath(path + "\u002f\u004d\u0063P\u0072"); _cdcd != nil {
			return _cdcd
		}
	}
	return nil
}

func (_dgcba ST_Jc) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_bgdd := _g.Attr{}
	_bgdd.Name = name
	switch _dgcba {
	case ST_JcUnset:
		_bgdd.Value = ""
	case ST_JcLeft:
		_bgdd.Value = "\u006c\u0065\u0066\u0074"
	case ST_JcRight:
		_bgdd.Value = "\u0072\u0069\u0067h\u0074"
	case ST_JcCenter:
		_bgdd.Value = "\u0063\u0065\u006e\u0074\u0065\u0072"
	case ST_JcCenterGroup:
		_bgdd.Value = "c\u0065\u006e\u0074\u0065\u0072\u0047\u0072\u006f\u0075\u0070"
	}
	return _bgdd, nil
}

func (_gbea *CT_Box) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_gbea.E = NewCT_OMathArg()
_ecba:
	for {
		_gcg, _eff := d.Token()
		if _eff != nil {
			return _eff
		}
		switch _bdb := _gcg.(type) {
		case _g.StartElement:
			switch _bdb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078P\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078P\u0072"}:
				_gbea.BoxPr = NewCT_BoxPr()
				if _fce := d.DecodeElement(_gbea.BoxPr, &_bdb); _fce != nil {
					return _fce
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _dga := d.DecodeElement(_gbea.E, &_bdb); _dga != nil {
					return _dga
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u006f\u0078\u0020\u0025\u0076", _bdb.Name)
				if _gfe := d.Skip(); _gfe != nil {
					return _gfe
				}
			}
		case _g.EndElement:
			break _ecba
		case _g.CharData:
		}
	}
	return nil
}

func (_cega *ST_LimLoc) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_agbdf, _cebd := d.Token()
	if _cebd != nil {
		return _cebd
	}
	if _aega, _gbbcd := _agbdf.(_g.EndElement); _gbbcd && _aega.Name == start.Name {
		*_cega = 1
		return nil
	}
	if _cfee, _beff := _agbdf.(_g.CharData); !_beff {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _agbdf)
	} else {
		switch string(_cfee) {
		case "":
			*_cega = 0
		case "\u0075\u006e\u0064\u004f\u0076\u0072":
			*_cega = 1
		case "\u0073\u0075\u0062\u0053\u0075\u0070":
			*_cega = 2
		}
	}
	_agbdf, _cebd = d.Token()
	if _cebd != nil {
		return _cebd
	}
	if _bfdb, _gdge := _agbdf.(_g.EndElement); _gdge && _bfdb.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _agbdf)
}

func (_cgaba ST_Shp) ValidateWithPath(path string) error {
	switch _cgaba {
	case 0, 1, 2:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cgaba))
	}
	return nil
}

func NewCT_MathPr() *CT_MathPr { _gcfg := &CT_MathPr{}; return _gcfg }

func NewCT_RPR() *CT_RPR { _bddg := &CT_RPR{}; return _bddg }

func (_daga *CT_Func) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_daga.FName = NewCT_OMathArg()
	_daga.E = NewCT_OMathArg()
_bdbd:
	for {
		_gbb, _fga := d.Token()
		if _fga != nil {
			return _fga
		}
		switch _gaa := _gbb.(type) {
		case _g.StartElement:
			switch _gaa.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063\u0050\u0072"}:
				_daga.FuncPr = NewCT_FuncPr()
				if _dafb := d.DecodeElement(_daga.FuncPr, &_gaa); _dafb != nil {
					return _dafb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u004e\u0061m\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u004e\u0061m\u0065"}:
				if _gffb := d.DecodeElement(_daga.FName, &_gaa); _gffb != nil {
					return _gffb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _gbbf := d.DecodeElement(_daga.E, &_gaa); _gbbf != nil {
					return _gbbf
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0046\u0075\u006e\u0063\u0020\u0025\u0076", _gaa.Name)
				if _ecaf := d.Skip(); _ecaf != nil {
					return _ecaf
				}
			}
		case _g.EndElement:
			break _bdbd
		case _g.CharData:
		}
	}
	return nil
}

func (_cgef *CT_EqArrPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cgef.BaseJc != nil {
		_dbdc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u0061\u0073\u0065\u004a\u0063"}}
		e.EncodeElement(_cgef.BaseJc, _dbdc)
	}
	if _cgef.MaxDist != nil {
		_dfc := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006d\u0061\u0078\u0044\u0069\u0073t"}}
		e.EncodeElement(_cgef.MaxDist, _dfc)
	}
	if _cgef.ObjDist != nil {
		_eca := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006f\u0062\u006a\u0044\u0069\u0073t"}}
		e.EncodeElement(_cgef.ObjDist, _eca)
	}
	if _cgef.RSpRule != nil {
		_abfd := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0072\u0053\u0070\u0052\u0075\u006ce"}}
		e.EncodeElement(_cgef.RSpRule, _abfd)
	}
	if _cgef.RSp != nil {
		_ead := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0072S\u0070"}}
		e.EncodeElement(_cgef.RSp, _ead)
	}
	if _cgef.CtrlPr != nil {
		_aeb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_cgef.CtrlPr, _aeb)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_SSubPr() *CT_SSubPr { _fbed := &CT_SSubPr{}; return _fbed }

func (_bfac *CT_MathPrChoice) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ceeb:
	for {
		_eged, _bbaf := d.Token()
		if _bbaf != nil {
			return _bbaf
		}
		switch _abffe := _eged.(type) {
		case _g.StartElement:
			switch _abffe.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}:
				_bfac.WrapIndent = NewCT_TwipsMeasure()
				if _gcfc := d.DecodeElement(_bfac.WrapIndent, &_abffe); _gcfc != nil {
					return _gcfc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}:
				_bfac.WrapRight = NewCT_OnOff()
				if _ebegg := d.DecodeElement(_bfac.WrapRight, &_abffe); _ebegg != nil {
					return _ebegg
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004d\u0061\u0074h\u0050\u0072\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _abffe.Name)
				if _dcfg := d.Skip(); _dcfg != nil {
					return _dcfg
				}
			}
		case _g.EndElement:
			break _ceeb
		case _g.CharData:
		}
	}
	return nil
}

type CT_Shp struct{ ValAttr ST_Shp }

func (_affe ST_Script) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_dfdc := _g.Attr{}
	_dfdc.Name = name
	switch _affe {
	case ST_ScriptUnset:
		_dfdc.Value = ""
	case ST_ScriptRoman:
		_dfdc.Value = "\u0072\u006f\u006da\u006e"
	case ST_ScriptScript:
		_dfdc.Value = "\u0073\u0063\u0072\u0069\u0070\u0074"
	case ST_ScriptFraktur:
		_dfdc.Value = "\u0066r\u0061\u006b\u0074\u0075\u0072"
	case ST_ScriptDouble_struck:
		_dfdc.Value = "\u0064\u006f\u0075\u0062\u006c\u0065\u002d\u0073\u0074\u0072\u0075\u0063\u006b"
	case ST_ScriptSans_serif:
		_dfdc.Value = "\u0073\u0061\u006e\u0073\u002d\u0073\u0065\u0072\u0069\u0066"
	case ST_ScriptMonospace:
		_dfdc.Value = "\u006do\u006e\u006f\u0073\u0070\u0061\u0063e"
	}
	return _dfdc, nil
}

func (_eef *CT_CtrlPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_BreakBin and its children, prefixing error messages with path
func (_gda *CT_BreakBin) ValidateWithPath(path string) error {
	if _fdc := _gda.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _fdc != nil {
		return _fdc
	}
	return nil
}

type ST_Script byte

// ValidateWithPath validates the CT_MathPrChoice and its children, prefixing error messages with path
func (_bgac *CT_MathPrChoice) ValidateWithPath(path string) error {
	if _bgac.WrapIndent != nil {
		if _cfgc := _bgac.WrapIndent.ValidateWithPath(path + "/\u0057\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"); _cfgc != nil {
			return _cfgc
		}
	}
	if _bgac.WrapRight != nil {
		if _bdgg := _bgac.WrapRight.ValidateWithPath(path + "\u002f\u0057\u0072\u0061\u0070\u0052\u0069\u0067\u0068\u0074"); _bdgg != nil {
			return _bdgg
		}
	}
	return nil
}

// Validate validates the CT_SPrePr and its children
func (_gab *CT_SPrePr) Validate() error {
	return _gab.ValidateWithPath("\u0043T\u005f\u0053\u0050\u0072\u0065\u0050r")
}

func (_cgeff *ST_Shp) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cdbc, _gbffb := d.Token()
	if _gbffb != nil {
		return _gbffb
	}
	if _fgfff, _dfdad := _cdbc.(_g.EndElement); _dfdad && _fgfff.Name == start.Name {
		*_cgeff = 1
		return nil
	}
	if _gfgaa, _dgab := _cdbc.(_g.CharData); !_dgab {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cdbc)
	} else {
		switch string(_gfgaa) {
		case "":
			*_cgeff = 0
		case "\u0063\u0065\u006e\u0074\u0065\u0072\u0065\u0064":
			*_cgeff = 1
		case "\u006d\u0061\u0074c\u0068":
			*_cgeff = 2
		}
	}
	_cdbc, _gbffb = d.Token()
	if _gbffb != nil {
		return _gbffb
	}
	if _fdfaf, _bfde := _cdbc.(_g.EndElement); _bfde && _fdfaf.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cdbc)
}

func (_ecbb *CT_RPRChoice) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cefd:
	for {
		_dfaf, _cefa := d.Token()
		if _cefa != nil {
			return _cefa
		}
		switch _gbgea := _dfaf.(type) {
		case _g.StartElement:
			switch _gbgea.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u006f\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u006f\u0072"}:
				_ecbb.Nor = NewCT_OnOff()
				if _fdaa := d.DecodeElement(_ecbb.Nor, &_gbgea); _fdaa != nil {
					return _fdaa
				}
			default:
				_f.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_R\u0050\u0052C\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _gbgea.Name)
				if _eddda := d.Skip(); _eddda != nil {
					return _eddda
				}
			}
		case _g.EndElement:
			break _cefd
		case _g.CharData:
		}
	}
	return nil
}

func NewCT_NaryPr() *CT_NaryPr { _gbd := &CT_NaryPr{}; return _gbd }

// ValidateWithPath validates the CT_RPRChoice and its children, prefixing error messages with path
func (_ggfd *CT_RPRChoice) ValidateWithPath(path string) error {
	if _ggfd.Nor != nil {
		if _egagd := _ggfd.Nor.ValidateWithPath(path + "\u002f\u004e\u006f\u0072"); _egagd != nil {
			return _egagd
		}
	}
	return nil
}

func NewCT_SPrePr() *CT_SPrePr { _cgefd := &CT_SPrePr{}; return _cgefd }

func (_ecad *ST_TopBot) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_ecad = 0
	case "\u0074\u006f\u0070":
		*_ecad = 1
	case "\u0062\u006f\u0074":
		*_ecad = 2
	}
	return nil
}

func (_ddge *CT_MathPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ddge.MathFont != nil {
		_bbab := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}}
		e.EncodeElement(_ddge.MathFont, _bbab)
	}
	if _ddge.BrkBin != nil {
		_eegf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062\u0072\u006b\u0042\u0069\u006e"}}
		e.EncodeElement(_ddge.BrkBin, _eegf)
	}
	if _ddge.BrkBinSub != nil {
		_babg := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0062\u0072\u006b\u0042\u0069\u006e\u0053\u0075\u0062"}}
		e.EncodeElement(_ddge.BrkBinSub, _babg)
	}
	if _ddge.SmallFrac != nil {
		_dgfe := _g.StartElement{Name: _g.Name{Local: "m\u003a\u0073\u006d\u0061\u006c\u006c\u0046\u0072\u0061\u0063"}}
		e.EncodeElement(_ddge.SmallFrac, _dgfe)
	}
	if _ddge.DispDef != nil {
		_acbf := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0064\u0069\u0073\u0070\u0044\u0065f"}}
		e.EncodeElement(_ddge.DispDef, _acbf)
	}
	if _ddge.LMargin != nil {
		_cgfd := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006c\u004d\u0061\u0072\u0067\u0069n"}}
		e.EncodeElement(_ddge.LMargin, _cgfd)
	}
	if _ddge.RMargin != nil {
		_ebfc := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0072\u004d\u0061\u0072\u0067\u0069n"}}
		e.EncodeElement(_ddge.RMargin, _ebfc)
	}
	if _ddge.DefJc != nil {
		_bef := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0064\u0065\u0066\u004a\u0063"}}
		e.EncodeElement(_ddge.DefJc, _bef)
	}
	if _ddge.PreSp != nil {
		_dece := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0070\u0072\u0065\u0053\u0070"}}
		e.EncodeElement(_ddge.PreSp, _dece)
	}
	if _ddge.PostSp != nil {
		_ebb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0070\u006f\u0073\u0074\u0053\u0070"}}
		e.EncodeElement(_ddge.PostSp, _ebb)
	}
	if _ddge.InterSp != nil {
		_fdfg := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0069\u006e\u0074\u0065\u0072\u0053p"}}
		e.EncodeElement(_ddge.InterSp, _fdfg)
	}
	if _ddge.IntraSp != nil {
		_bgga := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0069\u006e\u0074\u0072\u0061\u0053p"}}
		e.EncodeElement(_ddge.IntraSp, _bgga)
	}
	if _ddge.Choice != nil {
		_ddge.Choice.MarshalXML(e, _g.StartElement{})
	}
	if _ddge.IntLim != nil {
		_bgdc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0069\u006e\u0074\u004c\u0069\u006d"}}
		e.EncodeElement(_ddge.IntLim, _bgdc)
	}
	if _ddge.NaryLim != nil {
		_aabc := _g.StartElement{Name: _g.Name{Local: "\u006d:\u006e\u0061\u0072\u0079\u004c\u0069m"}}
		e.EncodeElement(_ddge.NaryLim, _aabc)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_Script() *CT_Script { _eaddb := &CT_Script{}; return _eaddb }

func (_efggd *CT_PhantPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _efggd.Show != nil {
		_gdabe := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0068\u006f\u0077"}}
		e.EncodeElement(_efggd.Show, _gdabe)
	}
	if _efggd.ZeroWid != nil {
		_cgedb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u007a\u0065\u0072\u006f\u0057\u0069d"}}
		e.EncodeElement(_efggd.ZeroWid, _cgedb)
	}
	if _efggd.ZeroAsc != nil {
		_ecdag := _g.StartElement{Name: _g.Name{Local: "\u006d:\u007a\u0065\u0072\u006f\u0041\u0073c"}}
		e.EncodeElement(_efggd.ZeroAsc, _ecdag)
	}
	if _efggd.ZeroDesc != nil {
		_aegc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u007a\u0065\u0072\u006f\u0044\u0065\u0073\u0063"}}
		e.EncodeElement(_efggd.ZeroDesc, _aegc)
	}
	if _efggd.Transp != nil {
		_acf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0074\u0072\u0061\u006e\u0073\u0070"}}
		e.EncodeElement(_efggd.Transp, _acf)
	}
	if _efggd.CtrlPr != nil {
		_dade := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_efggd.CtrlPr, _dade)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_SSup and its children, prefixing error messages with path
func (_gfaf *CT_SSup) ValidateWithPath(path string) error {
	if _gfaf.SSupPr != nil {
		if _gcgf := _gfaf.SSupPr.ValidateWithPath(path + "\u002fS\u0053\u0075\u0070\u0050\u0072"); _gcgf != nil {
			return _gcgf
		}
	}
	if _eebde := _gfaf.E.ValidateWithPath(path + "\u002f\u0045"); _eebde != nil {
		return _eebde
	}
	if _feec := _gfaf.Sup.ValidateWithPath(path + "\u002f\u0053\u0075\u0070"); _feec != nil {
		return _feec
	}
	return nil
}

type CT_OMathJc struct{ ValAttr ST_Jc }

func (_dcgc *CT_RPRChoice) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _dcgc.Nor != nil {
		_bffd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006eo\u0072"}}
		e.EncodeElement(_dcgc.Nor, _bffd)
	}
	return nil
}

func (_gcd *CT_BreakBinSub) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _gcd.ValAttr != ST_BreakBinSubUnset {
		_ecge, _efd := _gcd.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _efd != nil {
			return _efd
		}
		start.Attr = append(start.Attr, _ecge)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_Func struct {
	FuncPr *CT_FuncPr
	FName  *CT_OMathArg
	E      *CT_OMathArg
}

type ST_Shp byte

func (_dgda *CT_MC) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dgda.McPr != nil {
		_gbee := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006d\u0063\u0050\u0072"}}
		e.EncodeElement(_dgda.McPr, _gbee)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_BarPr() *CT_BarPr { _aef := &CT_BarPr{}; return _aef }

func (_cdda *MathPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cdda.CT_MathPr = *NewCT_MathPr()
_bgdg:
	for {
		_efdb, _aagbe := d.Token()
		if _aagbe != nil {
			return _aagbe
		}
		switch _bfcec := _efdb.(type) {
		case _g.StartElement:
			switch _bfcec.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d\u0061\u0074\u0068\u0046\u006f\u006e\u0074"}:
				_cdda.MathFont = NewCT_String()
				if _fcdce := d.DecodeElement(_cdda.MathFont, &_bfcec); _fcdce != nil {
					return _fcdce
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0072\u006b\u0042\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0072\u006b\u0042\u0069\u006e"}:
				_cdda.BrkBin = NewCT_BreakBin()
				if _aege := d.DecodeElement(_cdda.BrkBin, &_bfcec); _aege != nil {
					return _aege
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062r\u006b\u0042\u0069\u006e\u0053\u0075b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062r\u006b\u0042\u0069\u006e\u0053\u0075b"}:
				_cdda.BrkBinSub = NewCT_BreakBinSub()
				if _aggc := d.DecodeElement(_cdda.BrkBinSub, &_bfcec); _aggc != nil {
					return _aggc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073m\u0061\u006c\u006c\u0046\u0072\u0061c"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073m\u0061\u006c\u006c\u0046\u0072\u0061c"}:
				_cdda.SmallFrac = NewCT_OnOff()
				if _cded := d.DecodeElement(_cdda.SmallFrac, &_bfcec); _cded != nil {
					return _cded
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064i\u0073\u0070\u0044\u0065\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064i\u0073\u0070\u0044\u0065\u0066"}:
				_cdda.DispDef = NewCT_OnOff()
				if _faeb := d.DecodeElement(_cdda.DispDef, &_bfcec); _faeb != nil {
					return _faeb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006cM\u0061\u0072\u0067\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006cM\u0061\u0072\u0067\u0069\u006e"}:
				_cdda.LMargin = NewCT_TwipsMeasure()
				if _abgag := d.DecodeElement(_cdda.LMargin, &_bfcec); _abgag != nil {
					return _abgag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072M\u0061\u0072\u0067\u0069\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072M\u0061\u0072\u0067\u0069\u006e"}:
				_cdda.RMargin = NewCT_TwipsMeasure()
				if _ebfdb := d.DecodeElement(_cdda.RMargin, &_bfcec); _ebfdb != nil {
					return _ebfdb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0065\u0066J\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0065\u0066J\u0063"}:
				_cdda.DefJc = NewCT_OMathJc()
				if _cbebe := d.DecodeElement(_cdda.DefJc, &_bfcec); _cbebe != nil {
					return _cbebe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0072\u0065S\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0072\u0065S\u0070"}:
				_cdda.PreSp = NewCT_TwipsMeasure()
				if _babde := d.DecodeElement(_cdda.PreSp, &_bfcec); _babde != nil {
					return _babde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u006f\u0073\u0074\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u006f\u0073\u0074\u0053\u0070"}:
				_cdda.PostSp = NewCT_TwipsMeasure()
				if _bbd := d.DecodeElement(_cdda.PostSp, &_bfcec); _bbd != nil {
					return _bbd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069n\u0074\u0065\u0072\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069n\u0074\u0065\u0072\u0053\u0070"}:
				_cdda.InterSp = NewCT_TwipsMeasure()
				if _bbg := d.DecodeElement(_cdda.InterSp, &_bfcec); _bbg != nil {
					return _bbg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069n\u0074\u0072\u0061\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069n\u0074\u0072\u0061\u0053\u0070"}:
				_cdda.IntraSp = NewCT_TwipsMeasure()
				if _afgb := d.DecodeElement(_cdda.IntraSp, &_bfcec); _afgb != nil {
					return _afgb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077\u0072\u0061\u0070\u0049\u006e\u0064\u0065\u006e\u0074"}:
				_cdda.Choice = NewCT_MathPrChoice()
				if _ffge := d.DecodeElement(&_cdda.Choice.WrapIndent, &_bfcec); _ffge != nil {
					return _ffge
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0077r\u0061\u0070\u0052\u0069\u0067\u0068t"}:
				_cdda.Choice = NewCT_MathPrChoice()
				if _aabg := d.DecodeElement(&_cdda.Choice.WrapRight, &_bfcec); _aabg != nil {
					return _aabg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0069\u006e\u0074\u004c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0069\u006e\u0074\u004c\u0069\u006d"}:
				_cdda.IntLim = NewCT_LimLoc()
				if _dcfgb := d.DecodeElement(_cdda.IntLim, &_bfcec); _dcfgb != nil {
					return _dcfgb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006ea\u0072\u0079\u004c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006ea\u0072\u0079\u004c\u0069\u006d"}:
				_cdda.NaryLim = NewCT_LimLoc()
				if _gcfgb := d.DecodeElement(_cdda.NaryLim, &_bfcec); _gcfgb != nil {
					return _gcfgb
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u004d\u0061\u0074\u0068\u0050\u0072\u0020\u0025\u0076", _bfcec.Name)
				if _bfeba := d.Skip(); _bfeba != nil {
					return _bfeba
				}
			}
		case _g.EndElement:
			break _bgdg
		case _g.CharData:
		}
	}
	return nil
}

type CT_GroupChr struct {
	GroupChrPr *CT_GroupChrPr
	E          *CT_OMathArg
}

func (_fdbg *CT_String) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _fdbg.ValAttr != nil {
		start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _d.Sprintf("\u0025\u0076", *_fdbg.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_YAlign struct{ ValAttr _dc.ST_YAlign }

func (_bbdf ST_BreakBin) String() string {
	switch _bbdf {
	case 0:
		return ""
	case 1:
		return "\u0062\u0065\u0066\u006f\u0072\u0065"
	case 2:
		return "\u0061\u0066\u0074e\u0072"
	case 3:
		return "\u0072\u0065\u0070\u0065\u0061\u0074"
	}
	return ""
}

func (_ddad ST_Script) ValidateWithPath(path string) error {
	switch _ddad {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ddad))
	}
	return nil
}

func (_edgb ST_LimLoc) Validate() error { return _edgb.ValidateWithPath("") }

// ValidateWithPath validates the CT_SSubSup and its children, prefixing error messages with path
func (_efcf *CT_SSubSup) ValidateWithPath(path string) error {
	if _efcf.SSubSupPr != nil {
		if _eadee := _efcf.SSubSupPr.ValidateWithPath(path + "\u002f\u0053\u0053\u0075\u0062\u0053\u0075\u0070\u0050\u0072"); _eadee != nil {
			return _eadee
		}
	}
	if _babee := _efcf.E.ValidateWithPath(path + "\u002f\u0045"); _babee != nil {
		return _babee
	}
	if _egfg := _efcf.Sub.ValidateWithPath(path + "\u002f\u0053\u0075\u0062"); _egfg != nil {
		return _egfg
	}
	if _ccca := _efcf.Sup.ValidateWithPath(path + "\u002f\u0053\u0075\u0070"); _ccca != nil {
		return _ccca
	}
	return nil
}

// Validate validates the CT_SSubSupPr and its children
func (_agebf *CT_SSubSupPr) Validate() error {
	return _agebf.ValidateWithPath("\u0043\u0054\u005fS\u0053\u0075\u0062\u0053\u0075\u0070\u0050\u0072")
}

type CT_SPrePr struct{ CtrlPr *CT_CtrlPr }

// ValidateWithPath validates the CT_ManualBreak and its children, prefixing error messages with path
func (_cba *CT_ManualBreak) ValidateWithPath(path string) error {
	if _cba.AlnAtAttr != nil {
		if *_cba.AlnAtAttr < 1 {
			return _d.Errorf("\u0025\u0073/\u006d\u002e\u0041\u006cn\u0041\u0074A\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u003e\u003d\u0020\u0031\u0020\u0028\u0068\u0061\u0076e\u0020\u0025\u0076\u0029", path, *_cba.AlnAtAttr)
		}
		if *_cba.AlnAtAttr > 255 {
			return _d.Errorf("\u0025\u0073\u002f\u006d\u002e\u0041\u006c\u006e\u0041\u0074\u0041\u0074\u0074r\u0020\u006d\u0075\u0073\u0074\u0020b\u0065\u0020\u003c\u003d\u0020\u0032\u0035\u0035\u0020\u0028\u0068\u0061\u0076e\u0020\u0025\u0076\u0029", path, *_cba.AlnAtAttr)
		}
	}
	return nil
}

func NewCT_F() *CT_F {
	_gacf := &CT_F{}
	_gacf.Num = NewCT_OMathArg()
	_gacf.Den = NewCT_OMathArg()
	return _gacf
}

func (_dgcc *CT_SPre) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_dgcc.Sub = NewCT_OMathArg()
	_dgcc.Sup = NewCT_OMathArg()
	_dgcc.E = NewCT_OMathArg()
_efdd:
	for {
		_afbeg, _fgga := d.Token()
		if _fgga != nil {
			return _fgga
		}
		switch _beddg := _afbeg.(type) {
		case _g.StartElement:
			switch _beddg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065\u0050\u0072"}:
				_dgcc.SPrePr = NewCT_SPrePr()
				if _bdcbb := d.DecodeElement(_dgcc.SPrePr, &_beddg); _bdcbb != nil {
					return _bdcbb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0062"}:
				if _dafa := d.DecodeElement(_dgcc.Sub, &_beddg); _dafa != nil {
					return _dafa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0070"}:
				if _ccfa := d.DecodeElement(_dgcc.Sup, &_beddg); _ccfa != nil {
					return _ccfa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _bcee := d.DecodeElement(_dgcc.E, &_beddg); _bcee != nil {
					return _bcee
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0053\u0050\u0072\u0065\u0020\u0025\u0076", _beddg.Name)
				if _fdac := d.Skip(); _fdac != nil {
					return _fdac
				}
			}
		case _g.EndElement:
			break _efdd
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BoxPr and its children
func (_fbg *CT_BoxPr) Validate() error {
	return _fbg.ValidateWithPath("\u0043\u0054\u005f\u0042\u006f\u0078\u0050\u0072")
}

// ValidateWithPath validates the CT_RChoice and its children, prefixing error messages with path
func (_eadec *CT_RChoice) ValidateWithPath(path string) error {
	for _dadaf, _beae := range _eadec.T {
		if _cgff := _beae.ValidateWithPath(_d.Sprintf("\u0025\u0073\u002f\u0054\u005b\u0025\u0064\u005d", path, _dadaf)); _cgff != nil {
			return _cgff
		}
	}
	return nil
}

func (_eaac *OMathPara) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006d"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0073\u0063\u0068\u0065m\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0068\u0061\u0072e\u0064\u0054\u0079\u0070\u0065\u0073"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0077"}, Value: "ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065s\u0073i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u00306\u002fm\u0061\u0069n"})
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "m\u003a\u006f\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061"
	return _eaac.CT_OMathPara.MarshalXML(e, start)
}

func (_dbbce ST_Script) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_dbbce.String(), start)
}

type CT_MathPrChoice struct {
	WrapIndent *CT_TwipsMeasure
	WrapRight  *CT_OnOff
}

// ValidateWithPath validates the EG_OMathElements and its children, prefixing error messages with path
func (_egee *EG_OMathElements) ValidateWithPath(path string) error {
	for _eedd, _dgbbc := range _egee.EG_OMathMathElements {
		if _aggd := _dgbbc.ValidateWithPath(_d.Sprintf("%\u0073\u002f\u0045\u0047\u005f\u004fM\u0061\u0074\u0068\u004d\u0061\u0074\u0068\u0045\u006ce\u006d\u0065\u006et\u0073[\u0025\u0064\u005d", path, _eedd)); _aggd != nil {
			return _aggd
		}
	}
	return nil
}

type CT_OMathArgPr struct{ ArgSz *CT_Integer2 }

func (_dfg *CT_BorderBoxPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cgg:
	for {
		_ege, _eeb := d.Token()
		if _eeb != nil {
			return _eeb
		}
		switch _cae := _ege.(type) {
		case _g.StartElement:
			switch _cae.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0068i\u0064\u0065\u0054\u006f\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0068i\u0064\u0065\u0054\u006f\u0070"}:
				_dfg.HideTop = NewCT_OnOff()
				if _cfd := d.DecodeElement(_dfg.HideTop, &_cae); _cfd != nil {
					return _cfd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0068i\u0064\u0065\u0042\u006f\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0068i\u0064\u0065\u0042\u006f\u0074"}:
				_dfg.HideBot = NewCT_OnOff()
				if _dgf := d.DecodeElement(_dfg.HideBot, &_cae); _dgf != nil {
					return _dgf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0068\u0069\u0064\u0065\u004c\u0065\u0066\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0068\u0069\u0064\u0065\u004c\u0065\u0066\u0074"}:
				_dfg.HideLeft = NewCT_OnOff()
				if _cgc := d.DecodeElement(_dfg.HideLeft, &_cae); _cgc != nil {
					return _cgc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0068i\u0064\u0065\u0052\u0069\u0067\u0068t"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0068i\u0064\u0065\u0052\u0069\u0067\u0068t"}:
				_dfg.HideRight = NewCT_OnOff()
				if _bfa := d.DecodeElement(_dfg.HideRight, &_cae); _bfa != nil {
					return _bfa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073t\u0072\u0069\u006b\u0065\u0048"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073t\u0072\u0069\u006b\u0065\u0048"}:
				_dfg.StrikeH = NewCT_OnOff()
				if _efc := d.DecodeElement(_dfg.StrikeH, &_cae); _efc != nil {
					return _efc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073t\u0072\u0069\u006b\u0065\u0056"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073t\u0072\u0069\u006b\u0065\u0056"}:
				_dfg.StrikeV = NewCT_OnOff()
				if _feb := d.DecodeElement(_dfg.StrikeV, &_cae); _feb != nil {
					return _feb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0074\u0072\u0069\u006b\u0065\u0042\u004c\u0054\u0052"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0074\u0072\u0069\u006b\u0065\u0042\u004c\u0054\u0052"}:
				_dfg.StrikeBLTR = NewCT_OnOff()
				if _efa := d.DecodeElement(_dfg.StrikeBLTR, &_cae); _efa != nil {
					return _efa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0074\u0072\u0069\u006b\u0065\u0054\u004c\u0042\u0052"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0074\u0072\u0069\u006b\u0065\u0054\u004c\u0042\u0052"}:
				_dfg.StrikeTLBR = NewCT_OnOff()
				if _dbea := d.DecodeElement(_dfg.StrikeTLBR, &_cae); _dbea != nil {
					return _dbea
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_dfg.CtrlPr = NewCT_CtrlPr()
				if _ced := d.DecodeElement(_dfg.CtrlPr, &_cae); _ced != nil {
					return _ced
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u006f\u0072\u0064\u0065\u0072\u0042o\u0078P\u0072\u0020\u0025\u0076", _cae.Name)
				if _dca := d.Skip(); _dca != nil {
					return _dca
				}
			}
		case _g.EndElement:
			break _cgg
		case _g.CharData:
		}
	}
	return nil
}

const (
	ST_StyleUnset ST_Style = 0
	ST_StyleP     ST_Style = 1
	ST_StyleB     ST_Style = 2
	ST_StyleI     ST_Style = 3
	ST_StyleBi    ST_Style = 4
)

// ValidateWithPath validates the CT_XAlign and its children, prefixing error messages with path
func (_cceb *CT_XAlign) ValidateWithPath(path string) error {
	if _cceb.ValAttr == _dc.ST_XAlignUnset {
		return _d.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _gdad := _cceb.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _gdad != nil {
		return _gdad
	}
	return nil
}

// ValidateWithPath validates the CT_OMathJc and its children, prefixing error messages with path
func (_cbbf *CT_OMathJc) ValidateWithPath(path string) error {
	if _eeaae := _cbbf.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _eeaae != nil {
		return _eeaae
	}
	return nil
}

type CT_SSubSupPr struct {
	AlnScr *CT_OnOff
	CtrlPr *CT_CtrlPr
}

func (_cegfg ST_Jc) Validate() error { return _cegfg.ValidateWithPath("") }

// Validate validates the CT_Bar and its children
func (_gg *CT_Bar) Validate() error {
	return _gg.ValidateWithPath("\u0043\u0054\u005f\u0042\u0061\u0072")
}

func (_aab *CT_BreakBin) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _ace := range start.Attr {
		if _ace.Name.Local == "\u0076\u0061\u006c" {
			_aab.ValAttr.UnmarshalXMLAttr(_ace)
			continue
		}
	}
	for {
		_fdg, _gagf := d.Token()
		if _gagf != nil {
			return _d.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042\u0069\u006e\u003a\u0020\u0025\u0073", _gagf)
		}
		if _dgad, _efaa := _fdg.(_g.EndElement); _efaa && _dgad.Name == start.Name {
			break
		}
	}
	return nil
}

func (_eaba ST_TopBot) String() string {
	switch _eaba {
	case 0:
		return ""
	case 1:
		return "\u0074\u006f\u0070"
	case 2:
		return "\u0062\u006f\u0074"
	}
	return ""
}

// ValidateWithPath validates the CT_OMathArg and its children, prefixing error messages with path
func (_aebc *CT_OMathArg) ValidateWithPath(path string) error {
	if _aebc.ArgPr != nil {
		if _bebef := _aebc.ArgPr.ValidateWithPath(path + "\u002f\u0041\u0072\u0067\u0050\u0072"); _bebef != nil {
			return _bebef
		}
	}
	for _daba, _dgbaf := range _aebc.EG_OMathMathElements {
		if _dgfbb := _dgbaf.ValidateWithPath(_d.Sprintf("%\u0073\u002f\u0045\u0047\u005f\u004fM\u0061\u0074\u0068\u004d\u0061\u0074\u0068\u0045\u006ce\u006d\u0065\u006et\u0073[\u0025\u0064\u005d", path, _daba)); _dgfbb != nil {
			return _dgfbb
		}
	}
	if _aebc.CtrlPr != nil {
		if _dgc := _aebc.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _dgc != nil {
			return _dgc
		}
	}
	return nil
}

func NewCT_MC() *CT_MC { _ceg := &CT_MC{}; return _ceg }

func NewCT_Integer255() *CT_Integer255 { _bac := &CT_Integer255{}; _bac.ValAttr = 1; return _bac }

// Validate validates the CT_Rad and its children
func (_gcdac *CT_Rad) Validate() error {
	return _gcdac.ValidateWithPath("\u0043\u0054\u005f\u0052\u0061\u0064")
}

func (_gdac *CT_NaryPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_fbfa:
	for {
		_dggd, _eade := d.Token()
		if _eade != nil {
			return _eade
		}
		switch _dcea := _dggd.(type) {
		case _g.StartElement:
			switch _dcea.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0068\u0072"}:
				_gdac.Chr = NewCT_Char()
				if _aedb := d.DecodeElement(_gdac.Chr, &_dcea); _aedb != nil {
					return _aedb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0063"}:
				_gdac.LimLoc = NewCT_LimLoc()
				if _fgcg := d.DecodeElement(_gdac.LimLoc, &_dcea); _fgcg != nil {
					return _fgcg
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0077"}:
				_gdac.Grow = NewCT_OnOff()
				if _bgde := d.DecodeElement(_gdac.Grow, &_dcea); _bgde != nil {
					return _bgde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073u\u0062\u0048\u0069\u0064\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073u\u0062\u0048\u0069\u0064\u0065"}:
				_gdac.SubHide = NewCT_OnOff()
				if _bfff := d.DecodeElement(_gdac.SubHide, &_dcea); _bfff != nil {
					return _bfff
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073u\u0070\u0048\u0069\u0064\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073u\u0070\u0048\u0069\u0064\u0065"}:
				_gdac.SupHide = NewCT_OnOff()
				if _eead := d.DecodeElement(_gdac.SupHide, &_dcea); _eead != nil {
					return _eead
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_gdac.CtrlPr = NewCT_CtrlPr()
				if _acbb := d.DecodeElement(_gdac.CtrlPr, &_dcea); _acbb != nil {
					return _acbb
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004e\u0061\u0072y\u0050\u0072 \u0025\u0076", _dcea.Name)
				if _bafg := d.Skip(); _bafg != nil {
					return _bafg
				}
			}
		case _g.EndElement:
			break _fbfa
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BarPr and its children
func (_ed *CT_BarPr) Validate() error {
	return _ed.ValidateWithPath("\u0043\u0054\u005f\u0042\u0061\u0072\u0050\u0072")
}

func (_dff *CT_MR) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	_ecdc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	for _, _babd := range _dff.E {
		e.EncodeElement(_babd, _ecdc)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type ST_Jc byte

func (_fdde *CT_Phant) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fdde.PhantPr != nil {
		_acefa := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0070\u0068\u0061\u006e\u0074\u0050r"}}
		e.EncodeElement(_fdde.PhantPr, _acefa)
	}
	_addb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_fdde.E, _addb)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_fcbf ST_FType) ValidateWithPath(path string) error {
	switch _fcbf {
	case 0, 1, 2, 3, 4:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fcbf))
	}
	return nil
}

type OMath struct{ CT_OMath }

// Validate validates the CT_F and its children
func (_acgc *CT_F) Validate() error { return _acgc.ValidateWithPath("\u0043\u0054\u005f\u0046") }

func NewCT_RPRChoice() *CT_RPRChoice { _caec := &CT_RPRChoice{}; return _caec }

type CT_NaryPr struct {
	Chr     *CT_Char
	LimLoc  *CT_LimLoc
	Grow    *CT_OnOff
	SubHide *CT_OnOff
	SupHide *CT_OnOff
	CtrlPr  *CT_CtrlPr
}

func (_abe *CT_BoxPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_abf:
	for {
		_ffe, _fabg := d.Token()
		if _fabg != nil {
			return _fabg
		}
		switch _ccb := _ffe.(type) {
		case _g.StartElement:
			switch _ccb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006f\u0070\u0045m\u0075"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006f\u0070\u0045m\u0075"}:
				_abe.OpEmu = NewCT_OnOff()
				if _egd := d.DecodeElement(_abe.OpEmu, &_ccb); _egd != nil {
					return _egd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006eo\u0042\u0072\u0065\u0061\u006b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006eo\u0042\u0072\u0065\u0061\u006b"}:
				_abe.NoBreak = NewCT_OnOff()
				if _cea := d.DecodeElement(_abe.NoBreak, &_ccb); _cea != nil {
					return _cea
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0069\u0066\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0069\u0066\u0066"}:
				_abe.Diff = NewCT_OnOff()
				if _afba := d.DecodeElement(_abe.Diff, &_ccb); _afba != nil {
					return _afba
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0072\u006b"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0072\u006b"}:
				_abe.Brk = NewCT_ManualBreak()
				if _aae := d.DecodeElement(_abe.Brk, &_ccb); _aae != nil {
					return _aae
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u006c\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u006c\u006e"}:
				_abe.Aln = NewCT_OnOff()
				if _ecbd := d.DecodeElement(_abe.Aln, &_ccb); _ecbd != nil {
					return _ecbd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_abe.CtrlPr = NewCT_CtrlPr()
				if _cfa := d.DecodeElement(_abe.CtrlPr, &_ccb); _cfa != nil {
					return _cfa
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0042\u006f\u0078\u0050\u0072\u0020\u0025\u0076", _ccb.Name)
				if _ba := d.Skip(); _ba != nil {
					return _ba
				}
			}
		case _g.EndElement:
			break _abf
		case _g.CharData:
		}
	}
	return nil
}

func (_fcdc *CT_EqArrPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cebb:
	for {
		_cffg, _ggb := d.Token()
		if _ggb != nil {
			return _ggb
		}
		switch _dbf := _cffg.(type) {
		case _g.StartElement:
			switch _dbf.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0073\u0065\u004a\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0073\u0065\u004a\u0063"}:
				_fcdc.BaseJc = NewCT_YAlign()
				if _ggf := d.DecodeElement(_fcdc.BaseJc, &_dbf); _ggf != nil {
					return _ggf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006da\u0078\u0044\u0069\u0073\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006da\u0078\u0044\u0069\u0073\u0074"}:
				_fcdc.MaxDist = NewCT_OnOff()
				if _gaf := d.DecodeElement(_fcdc.MaxDist, &_dbf); _gaf != nil {
					return _gaf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006fb\u006a\u0044\u0069\u0073\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006fb\u006a\u0044\u0069\u0073\u0074"}:
				_fcdc.ObjDist = NewCT_OnOff()
				if _ecae := d.DecodeElement(_fcdc.ObjDist, &_dbf); _ecae != nil {
					return _ecae
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072S\u0070\u0052\u0075\u006c\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072S\u0070\u0052\u0075\u006c\u0065"}:
				_fcdc.RSpRule = NewCT_SpacingRule()
				if _fgf := d.DecodeElement(_fcdc.RSpRule, &_dbf); _fgf != nil {
					return _fgf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0053\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0053\u0070"}:
				_fcdc.RSp = NewCT_UnSignedInteger()
				if _aff := d.DecodeElement(_fcdc.RSp, &_dbf); _aff != nil {
					return _aff
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_fcdc.CtrlPr = NewCT_CtrlPr()
				if _aaf := d.DecodeElement(_fcdc.CtrlPr, &_dbf); _aaf != nil {
					return _aaf
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fE\u0071\u0041\u0072\u0072\u0050\u0072\u0020\u0025\u0076", _dbf.Name)
				if _cfcd := d.Skip(); _cfcd != nil {
					return _cfcd
				}
			}
		case _g.EndElement:
			break _cebb
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_MCS and its children, prefixing error messages with path
func (_dddb *CT_MCS) ValidateWithPath(path string) error {
	for _fdcd, _fbdf := range _dddb.Mc {
		if _cadfc := _fbdf.ValidateWithPath(_d.Sprintf("\u0025s\u002f\u004d\u0063\u005b\u0025\u0064]", path, _fdcd)); _cadfc != nil {
			return _cadfc
		}
	}
	return nil
}

type MathPr struct{ CT_MathPr }

type ST_TopBot byte

type CT_Box struct {
	BoxPr *CT_BoxPr
	E     *CT_OMathArg
}

func (_cabg *EG_OMathMathElements) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_faecf:
	for {
		_adeb, _gdgd := d.Token()
		if _gdgd != nil {
			return _gdgd
		}
		switch _fgfb := _adeb.(type) {
		case _g.StartElement:
			switch _fgfb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_cabg.Acc = NewCT_Acc()
				if _faea := d.DecodeElement(_cabg.Acc, &_fgfb); _faea != nil {
					return _faea
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_cabg.Bar = NewCT_Bar()
				if _fbfe := d.DecodeElement(_cabg.Bar, &_fgfb); _fbfe != nil {
					return _fbfe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_cabg.Box = NewCT_Box()
				if _ddcd := d.DecodeElement(_cabg.Box, &_fgfb); _ddcd != nil {
					return _ddcd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_cabg.BorderBox = NewCT_BorderBox()
				if _ceae := d.DecodeElement(_cabg.BorderBox, &_fgfb); _ceae != nil {
					return _ceae
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_cabg.D = NewCT_D()
				if _gacc := d.DecodeElement(_cabg.D, &_fgfb); _gacc != nil {
					return _gacc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_cabg.EqArr = NewCT_EqArr()
				if _caac := d.DecodeElement(_cabg.EqArr, &_fgfb); _caac != nil {
					return _caac
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_cabg.F = NewCT_F()
				if _efag := d.DecodeElement(_cabg.F, &_fgfb); _efag != nil {
					return _efag
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_cabg.Func = NewCT_Func()
				if _dgcce := d.DecodeElement(_cabg.Func, &_fgfb); _dgcce != nil {
					return _dgcce
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_cabg.GroupChr = NewCT_GroupChr()
				if _fgcd := d.DecodeElement(_cabg.GroupChr, &_fgfb); _fgcd != nil {
					return _fgcd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_cabg.LimLow = NewCT_LimLow()
				if _ffbf := d.DecodeElement(_cabg.LimLow, &_fgfb); _ffbf != nil {
					return _ffbf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_cabg.LimUpp = NewCT_LimUpp()
				if _fgde := d.DecodeElement(_cabg.LimUpp, &_fgfb); _fgde != nil {
					return _fgde
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_cabg.M = NewCT_M()
				if _aaca := d.DecodeElement(_cabg.M, &_fgfb); _aaca != nil {
					return _aaca
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_cabg.Nary = NewCT_Nary()
				if _cffb := d.DecodeElement(_cabg.Nary, &_fgfb); _cffb != nil {
					return _cffb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_cabg.Phant = NewCT_Phant()
				if _eaddf := d.DecodeElement(_cabg.Phant, &_fgfb); _eaddf != nil {
					return _eaddf
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_cabg.Rad = NewCT_Rad()
				if _caeed := d.DecodeElement(_cabg.Rad, &_fgfb); _caeed != nil {
					return _caeed
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_cabg.SPre = NewCT_SPre()
				if _gaad := d.DecodeElement(_cabg.SPre, &_fgfb); _gaad != nil {
					return _gaad
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_cabg.SSub = NewCT_SSub()
				if _fgbgd := d.DecodeElement(_cabg.SSub, &_fgfb); _fgbgd != nil {
					return _fgbgd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_cabg.SSubSup = NewCT_SSubSup()
				if _ddfb := d.DecodeElement(_cabg.SSubSup, &_fgfb); _ddfb != nil {
					return _ddfb
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_cabg.SSup = NewCT_SSup()
				if _abfbd := d.DecodeElement(_cabg.SSup, &_fgfb); _abfbd != nil {
					return _abfbd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_cabg.R = NewCT_R()
				if _affb := d.DecodeElement(_cabg.R, &_fgfb); _affb != nil {
					return _affb
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070o\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006de\u006et \u006f\u006e\u0020\u0045\u0047\u005f\u004f\u004d\u0061\u0074\u0068\u004d\u0061\u0074\u0068\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0020\u0025\u0076", _fgfb.Name)
				if _gbdg := d.Skip(); _gbdg != nil {
					return _gbdg
				}
			}
		case _g.EndElement:
			break _faecf
		case _g.CharData:
		}
	}
	return nil
}

func (_ccce *CT_LimLow) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ccce.LimLowPr != nil {
		_aeac := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u004c\u006f\u0077\u0050\u0072"}}
		e.EncodeElement(_ccce.LimLowPr, _aeac)
	}
	_ecaef := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_ccce.E, _ecaef)
	_faca := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006ci\u006d"}}
	e.EncodeElement(_ccce.Lim, _faca)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_OMath struct{ EG_OMathMathElements []*EG_OMathMathElements }

// Validate validates the CT_MC and its children
func (_bdba *CT_MC) Validate() error { return _bdba.ValidateWithPath("\u0043\u0054\u005fM\u0043") }

// Validate validates the CT_GroupChr and its children
func (_aaba *CT_GroupChr) Validate() error {
	return _aaba.ValidateWithPath("C\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072")
}

func NewCT_OMath() *CT_OMath { _gbeea := &CT_OMath{}; return _gbeea }

func (_ggebc ST_BreakBinSub) Validate() error { return _ggebc.ValidateWithPath("") }

func (_ccaa ST_BreakBin) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_ccaa.String(), start)
}

func (_ebfd *CT_SSup) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ebfd.E = NewCT_OMathArg()
	_ebfd.Sup = NewCT_OMathArg()
_eeed:
	for {
		_dafef, _ecab := d.Token()
		if _ecab != nil {
			return _ecab
		}
		switch _debgbd := _dafef.(type) {
		case _g.StartElement:
			switch _debgbd.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070\u0050\u0072"}:
				_ebfd.SSupPr = NewCT_SSupPr()
				if _gccd := d.DecodeElement(_ebfd.SSupPr, &_debgbd); _gccd != nil {
					return _gccd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _egfa := d.DecodeElement(_ebfd.E, &_debgbd); _egfa != nil {
					return _egfa
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0075\u0070"}:
				if _cbce := d.DecodeElement(_ebfd.Sup, &_debgbd); _cbce != nil {
					return _cbce
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0053\u0053\u0075\u0070\u0020\u0025\u0076", _debgbd.Name)
				if _cfbfb := d.Skip(); _cfbfb != nil {
					return _cfbfb
				}
			}
		case _g.EndElement:
			break _eeed
		case _g.CharData:
		}
	}
	return nil
}

func (_edca *ST_Style) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_edca = 0
	case "\u0070":
		*_edca = 1
	case "\u0062":
		*_edca = 2
	case "\u0069":
		*_edca = 3
	case "\u0062\u0069":
		*_edca = 4
	}
	return nil
}

func (_cabf *CT_UnSignedInteger) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _cddb := range start.Attr {
		if _cddb.Name.Local == "\u0076\u0061\u006c" {
			_gfdf, _bgaef := _be.ParseUint(_cddb.Value, 10, 32)
			if _bgaef != nil {
				return _bgaef
			}
			_cabf.ValAttr = uint32(_gfdf)
			continue
		}
	}
	for {
		_ebeeef, _aagf := d.Token()
		if _aagf != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054_\u0055\u006e\u0053\u0069\u0067\u006e\u0065d\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u003a\u0020\u0025\u0073", _aagf)
		}
		if _ebeb, _ddab := _ebeeef.(_g.EndElement); _ddab && _ebeb.Name == start.Name {
			break
		}
	}
	return nil
}

type CT_D struct {
	DPr *CT_DPr
	E   []*CT_OMathArg
}

func (_efbc *CT_EqArr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_afa:
	for {
		_dbeg, _bde := d.Token()
		if _bde != nil {
			return _bde
		}
		switch _bee := _dbeg.(type) {
		case _g.StartElement:
			switch _bee.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065q\u0041\u0072\u0072\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065q\u0041\u0072\u0072\u0050\u0072"}:
				_efbc.EqArrPr = NewCT_EqArrPr()
				if _dfe := d.DecodeElement(_efbc.EqArrPr, &_bee); _dfe != nil {
					return _dfe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				_cdbb := NewCT_OMathArg()
				if _bag := d.DecodeElement(_cdbb, &_bee); _bag != nil {
					return _bag
				}
				_efbc.E = append(_efbc.E, _cdbb)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072\u0020\u0025\u0076", _bee.Name)
				if _bfd := d.Skip(); _bfd != nil {
					return _bfd
				}
			}
		case _g.EndElement:
			break _afa
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Rad and its children, prefixing error messages with path
func (_fdcce *CT_Rad) ValidateWithPath(path string) error {
	if _fdcce.RadPr != nil {
		if _ebaa := _fdcce.RadPr.ValidateWithPath(path + "\u002f\u0052\u0061\u0064\u0050\u0072"); _ebaa != nil {
			return _ebaa
		}
	}
	if _gdcb := _fdcce.Deg.ValidateWithPath(path + "\u002f\u0044\u0065\u0067"); _gdcb != nil {
		return _gdcb
	}
	if _gfb := _fdcce.E.ValidateWithPath(path + "\u002f\u0045"); _gfb != nil {
		return _gfb
	}
	return nil
}

// Validate validates the CT_String and its children
func (_aecg *CT_String) Validate() error {
	return _aecg.ValidateWithPath("\u0043T\u005f\u0053\u0074\u0072\u0069\u006eg")
}

type CT_MR struct{ E []*CT_OMathArg }

func ParseUnionST_OnOff(s string) (_dc.ST_OnOff, error) { return _dc.ParseUnionST_OnOff(s) }

func (_cabc *CT_SSup) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _cabc.SSupPr != nil {
		_dafg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073\u0053\u0075\u0070\u0050\u0072"}}
		e.EncodeElement(_cabc.SSupPr, _dafg)
	}
	_cdd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_cabc.E, _cdd)
	_cdgf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073u\u0070"}}
	e.EncodeElement(_cabc.Sup, _cdgf)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_edda *CT_Integer255) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_edda.ValAttr = 1
	for _, _addd := range start.Attr {
		if _addd.Name.Local == "\u0076\u0061\u006c" {
			_ddfa, _dde := _be.ParseInt(_addd.Value, 10, 64)
			if _dde != nil {
				return _dde
			}
			_edda.ValAttr = _ddfa
			continue
		}
	}
	for {
		_egag, _bgae := d.Token()
		if _bgae != nil {
			return _d.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0043\u0054\u005f\u0049n\u0074e\u0067e\u0072\u0032\u0035\u0035\u003a\u0020\u0025s", _bgae)
		}
		if _aec, _dacf := _egag.(_g.EndElement); _dacf && _aec.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Func and its children
func (_caaa *CT_Func) Validate() error {
	return _caaa.ValidateWithPath("\u0043T\u005f\u0046\u0075\u006e\u0063")
}

type CT_FuncPr struct{ CtrlPr *CT_CtrlPr }

func (_dbfd ST_TopBot) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_agfc := _g.Attr{}
	_agfc.Name = name
	switch _dbfd {
	case ST_TopBotUnset:
		_agfc.Value = ""
	case ST_TopBotTop:
		_agfc.Value = "\u0074\u006f\u0070"
	case ST_TopBotBot:
		_agfc.Value = "\u0062\u006f\u0074"
	}
	return _agfc, nil
}

func (_efec *CT_SSubPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _efec.CtrlPr != nil {
		_edaf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_efec.CtrlPr, _edaf)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_RChoice and its children
func (_ffae *CT_RChoice) Validate() error {
	return _ffae.ValidateWithPath("\u0043\u0054\u005f\u0052\u0043\u0068\u006f\u0069\u0063\u0065")
}

func (_befba *CT_Text) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _bgaeb := range start.Attr {
		if _bgaeb.Name.Space == "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065" && _bgaeb.Name.Local == "\u0073\u0070\u0061c\u0065" {
			_cagd, _bbff := _bgaeb.Value, error(nil)
			if _bbff != nil {
				return _bbff
			}
			_befba.SpaceAttr = &_cagd
			continue
		}
	}
	for {
		_abfe, _bace := d.Token()
		if _bace != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0054\u0065\u0078\u0074\u003a\u0020\u0025\u0073", _bace)
		}
		if _ddec, _afcc := _abfe.(_g.CharData); _afcc {
			_befba.Content = string(_ddec)
		}
		if _gbcb, _facd := _abfe.(_g.EndElement); _facd && _gbcb.Name == start.Name {
			break
		}
	}
	return nil
}

func (_debc *CT_RadPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cgfec:
	for {
		_gbcd, _fdce := d.Token()
		if _fdce != nil {
			return _fdce
		}
		switch _fabgg := _gbcd.(type) {
		case _g.StartElement:
			switch _fabgg.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064e\u0067\u0048\u0069\u0064\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064e\u0067\u0048\u0069\u0064\u0065"}:
				_debc.DegHide = NewCT_OnOff()
				if _dedc := d.DecodeElement(_debc.DegHide, &_fabgg); _dedc != nil {
					return _dedc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_debc.CtrlPr = NewCT_CtrlPr()
				if _cbd := d.DecodeElement(_debc.CtrlPr, &_fabgg); _cbd != nil {
					return _cbd
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0061\u0064\u0050\u0072\u0020\u0025\u0076", _fabgg.Name)
				if _bcgd := d.Skip(); _bcgd != nil {
					return _bcgd
				}
			}
		case _g.EndElement:
			break _cgfec
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_LimLoc and its children, prefixing error messages with path
func (_cdbbf *CT_LimLoc) ValidateWithPath(path string) error {
	if _cdbbf.ValAttr == ST_LimLocUnset {
		return _d.Errorf("\u0025\u0073\u002fV\u0061\u006c\u0041\u0074t\u0072\u0020\u0069\u0073\u0020\u0061\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _aace := _cdbbf.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _aace != nil {
		return _aace
	}
	return nil
}

func NewCT_LimLoc() *CT_LimLoc { _gga := &CT_LimLoc{}; _gga.ValAttr = ST_LimLoc(1); return _gga }

func (_bege *OMath) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bege.CT_OMath = *NewCT_OMath()
_bcce:
	for {
		_fcdg, _cbaa := d.Token()
		if _cbaa != nil {
			return _cbaa
		}
		switch _edafa := _fcdg.(type) {
		case _g.StartElement:
			switch _edafa.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0063\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0063\u0063"}:
				_baab := NewEG_OMathMathElements()
				_baab.Acc = NewCT_Acc()
				if _abaf := d.DecodeElement(_baab.Acc, &_edafa); _abaf != nil {
					return _abaf
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _baab)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u0061\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u0061\u0072"}:
				_acbag := NewEG_OMathMathElements()
				_acbag.Bar = NewCT_Bar()
				if _dfdd := d.DecodeElement(_acbag.Bar, &_edafa); _dfdd != nil {
					return _dfdd
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _acbag)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062\u006f\u0078"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062\u006f\u0078"}:
				_aaccb := NewEG_OMathMathElements()
				_aaccb.Box = NewCT_Box()
				if _cadc := d.DecodeElement(_aaccb.Box, &_edafa); _cadc != nil {
					return _cadc
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _aaccb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0062o\u0072\u0064\u0065\u0072\u0042\u006fx"}:
				_dbabc := NewEG_OMathMathElements()
				_dbabc.BorderBox = NewCT_BorderBox()
				if _cegc := d.DecodeElement(_dbabc.BorderBox, &_edafa); _cegc != nil {
					return _cegc
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _dbabc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064"}:
				_gbfff := NewEG_OMathMathElements()
				_gbfff.D = NewCT_D()
				if _cecd := d.DecodeElement(_gbfff.D, &_edafa); _cecd != nil {
					return _cecd
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _gbfff)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065\u0071\u0041r\u0072"}:
				_gbbc := NewEG_OMathMathElements()
				_gbbc.EqArr = NewCT_EqArr()
				if _gbbdb := d.DecodeElement(_gbbc.EqArr, &_edafa); _gbbdb != nil {
					return _gbbdb
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _gbbc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066"}:
				_fgcc := NewEG_OMathMathElements()
				_fgcc.F = NewCT_F()
				if _gcgc := d.DecodeElement(_fgcc.F, &_edafa); _gcgc != nil {
					return _gcgc
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _fgcc)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0075\u006e\u0063"}:
				_degf := NewEG_OMathMathElements()
				_degf.Func = NewCT_Func()
				if _ccegd := d.DecodeElement(_degf.Func, &_edafa); _ccegd != nil {
					return _ccegd
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _degf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0067\u0072\u006f\u0075\u0070\u0043\u0068\u0072"}:
				_bdgf := NewEG_OMathMathElements()
				_bdgf.GroupChr = NewCT_GroupChr()
				if _ecgfb := d.DecodeElement(_bdgf.GroupChr, &_edafa); _ecgfb != nil {
					return _ecgfb
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _bdgf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077"}:
				_fccac := NewEG_OMathMathElements()
				_fccac.LimLow = NewCT_LimLow()
				if _ddgg := d.DecodeElement(_fccac.LimLow, &_edafa); _ddgg != nil {
					return _ddgg
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _fccac)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u0055\u0070\u0070"}:
				_ffecf := NewEG_OMathMathElements()
				_ffecf.LimUpp = NewCT_LimUpp()
				if _aacd := d.DecodeElement(_ffecf.LimUpp, &_edafa); _aacd != nil {
					return _aacd
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _ffecf)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006d"}:
				_abed := NewEG_OMathMathElements()
				_abed.M = NewCT_M()
				if _cefaf := d.DecodeElement(_abed.M, &_edafa); _cefaf != nil {
					return _cefaf
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _abed)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0061\u0072\u0079"}:
				_beaa := NewEG_OMathMathElements()
				_beaa.Nary = NewCT_Nary()
				if _bafef := d.DecodeElement(_beaa.Nary, &_edafa); _bafef != nil {
					return _bafef
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _beaa)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0070\u0068\u0061n\u0074"}:
				_cedb := NewEG_OMathMathElements()
				_cedb.Phant = NewCT_Phant()
				if _dbdbb := d.DecodeElement(_cedb.Phant, &_edafa); _dbdbb != nil {
					return _dbdbb
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _cedb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0061\u0064"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0061\u0064"}:
				_agca := NewEG_OMathMathElements()
				_agca.Rad = NewCT_Rad()
				if _dfgc := d.DecodeElement(_agca.Rad, &_edafa); _dfgc != nil {
					return _dfgc
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _agca)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0050\u0072\u0065"}:
				_egebb := NewEG_OMathMathElements()
				_egebb.SPre = NewCT_SPre()
				if _bgcd := d.DecodeElement(_egebb.SPre, &_edafa); _bgcd != nil {
					return _bgcd
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _egebb)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0062"}:
				_fcdca := NewEG_OMathMathElements()
				_fcdca.SSub = NewCT_SSub()
				if _cacbe := d.DecodeElement(_fcdca.SSub, &_edafa); _cacbe != nil {
					return _cacbe
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _fcdca)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073S\u0075\u0062\u0053\u0075\u0070"}:
				_fbaed := NewEG_OMathMathElements()
				_fbaed.SSubSup = NewCT_SSubSup()
				if _ddbe := d.DecodeElement(_fbaed.SSubSup, &_edafa); _ddbe != nil {
					return _ddbe
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _fbaed)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0073\u0053\u0075\u0070"}:
				_dfca := NewEG_OMathMathElements()
				_dfca.SSup = NewCT_SSup()
				if _ddbc := d.DecodeElement(_dfca.SSup, &_edafa); _ddbc != nil {
					return _ddbc
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _dfca)
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072"}:
				_gaadc := NewEG_OMathMathElements()
				_gaadc.R = NewCT_R()
				if _eace := d.DecodeElement(_gaadc.R, &_edafa); _eace != nil {
					return _eace
				}
				_bege.EG_OMathMathElements = append(_bege.EG_OMathMathElements, _gaadc)
			default:
				_f.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u004f\u004d\u0061t\u0068 \u0025\u0076", _edafa.Name)
				if _gcfd := d.Skip(); _gcfd != nil {
					return _gcfd
				}
			}
		case _g.EndElement:
			break _bcce
		case _g.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_SSupPr and its children, prefixing error messages with path
func (_cacb *CT_SSupPr) ValidateWithPath(path string) error {
	if _cacb.CtrlPr != nil {
		if _bgee := _cacb.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _bgee != nil {
			return _bgee
		}
	}
	return nil
}

type CT_FPr struct {
	Type   *CT_FType
	CtrlPr *CT_CtrlPr
}

// ValidateWithPath validates the CT_EqArr and its children, prefixing error messages with path
func (_adcg *CT_EqArr) ValidateWithPath(path string) error {
	if _adcg.EqArrPr != nil {
		if _edc := _adcg.EqArrPr.ValidateWithPath(path + "\u002f\u0045\u0071\u0041\u0072\u0072\u0050\u0072"); _edc != nil {
			return _edc
		}
	}
	for _fffc, _ecbf := range _adcg.E {
		if _ddfg := _ecbf.ValidateWithPath(_d.Sprintf("\u0025\u0073\u002f\u0045\u005b\u0025\u0064\u005d", path, _fffc)); _ddfg != nil {
			return _ddfg
		}
	}
	return nil
}

// Validate validates the CT_LimLoc and its children
func (_aecd *CT_LimLoc) Validate() error {
	return _aecd.ValidateWithPath("\u0043T\u005f\u004c\u0069\u006d\u004c\u006fc")
}

// ValidateWithPath validates the CT_OMathParaPr and its children, prefixing error messages with path
func (_fadeb *CT_OMathParaPr) ValidateWithPath(path string) error {
	if _fadeb.Jc != nil {
		if _cacc := _fadeb.Jc.ValidateWithPath(path + "\u002f\u004a\u0063"); _cacc != nil {
			return _cacc
		}
	}
	return nil
}

func (_ece *CT_FuncPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _ece.CtrlPr != nil {
		_dda := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_ece.CtrlPr, _dda)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_Acc() *CT_Acc { _cf := &CT_Acc{}; _cf.E = NewCT_OMathArg(); return _cf }

func NewEG_OMathMathElements() *EG_OMathMathElements { _ebbe := &EG_OMathMathElements{}; return _ebbe }

func (_bdgb *CT_F) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_bdgb.Num = NewCT_OMathArg()
	_bdgb.Den = NewCT_OMathArg()
_baff:
	for {
		_ecc, _gadf := d.Token()
		if _gadf != nil {
			return _gadf
		}
		switch _dfgb := _ecc.(type) {
		case _g.StartElement:
			switch _dfgb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0066\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0066\u0050\u0072"}:
				_bdgb.FPr = NewCT_FPr()
				if _bca := d.DecodeElement(_bdgb.FPr, &_dfgb); _bca != nil {
					return _bca
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006e\u0075\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006e\u0075\u006d"}:
				if _gdgc := d.DecodeElement(_bdgb.Num, &_dfgb); _gdgc != nil {
					return _gdgc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0064\u0065\u006e"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0064\u0065\u006e"}:
				if _fcb := d.DecodeElement(_bdgb.Den, &_dfgb); _fcb != nil {
					return _fcb
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u0046\u0020\u0025\u0076", _dfgb.Name)
				if _dgd := d.Skip(); _dgd != nil {
					return _dgd
				}
			}
		case _g.EndElement:
			break _baff
		case _g.CharData:
		}
	}
	return nil
}

type CT_SSub struct {
	SSubPr *CT_SSubPr
	E      *CT_OMathArg
	Sub    *CT_OMathArg
}

func (_cfdf *CT_FuncPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_cace:
	for {
		_dee, _gcge := d.Token()
		if _gcge != nil {
			return _gcge
		}
		switch _beee := _dee.(type) {
		case _g.StartElement:
			switch _beee.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0063\u0074\u0072\u006c\u0050\u0072"}:
				_cfdf.CtrlPr = NewCT_CtrlPr()
				if _bcc := d.DecodeElement(_cfdf.CtrlPr, &_beee); _bcc != nil {
					return _bcc
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0046\u0075\u006ec\u0050\u0072 \u0025\u0076", _beee.Name)
				if _cfg := d.Skip(); _cfg != nil {
					return _cfg
				}
			}
		case _g.EndElement:
			break _cace
		case _g.CharData:
		}
	}
	return nil
}

func (_abgc *CT_NaryPr) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _abgc.Chr != nil {
		_fabfc := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063h\u0072"}}
		e.EncodeElement(_abgc.Chr, _fabfc)
	}
	if _abgc.LimLoc != nil {
		_agb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006c\u0069\u006d\u004c\u006f\u0063"}}
		e.EncodeElement(_abgc.LimLoc, _agb)
	}
	if _abgc.Grow != nil {
		_ebee := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0067\u0072\u006f\u0077"}}
		e.EncodeElement(_abgc.Grow, _ebee)
	}
	if _abgc.SubHide != nil {
		_gfae := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0075\u0062\u0048\u0069\u0064e"}}
		e.EncodeElement(_abgc.SubHide, _gfae)
	}
	if _abgc.SupHide != nil {
		_dad := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0073\u0075\u0070\u0048\u0069\u0064e"}}
		e.EncodeElement(_abgc.SupHide, _dad)
	}
	if _abgc.CtrlPr != nil {
		_cbb := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_abgc.CtrlPr, _cbb)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_BorderBox and its children, prefixing error messages with path
func (_gb *CT_BorderBox) ValidateWithPath(path string) error {
	if _gb.BorderBoxPr != nil {
		if _bd := _gb.BorderBoxPr.ValidateWithPath(path + "\u002f\u0042\u006fr\u0064\u0065\u0072\u0042\u006f\u0078\u0050\u0072"); _bd != nil {
			return _bd
		}
	}
	if _dg := _gb.E.ValidateWithPath(path + "\u002f\u0045"); _dg != nil {
		return _dg
	}
	return nil
}

func (_cbfb *CT_Script) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _cbfb.ValAttr != ST_ScriptUnset {
		_dfcc, _adce := _cbfb.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _adce != nil {
			return _adce
		}
		start.Attr = append(start.Attr, _dfcc)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_bcea *CT_R) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_ccbc:
	for {
		_afbee, _becf := d.Token()
		if _becf != nil {
			return _becf
		}
		switch _cefe := _afbee.(type) {
		case _g.StartElement:
			switch _cefe.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0072\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0072\u0050\u0072"}:
				_bcea.RPr = NewCT_RPR()
				if _ccgd := d.DecodeElement(_bcea.RPr, &_cefe); _ccgd != nil {
					return _ccgd
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0074"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0074"}:
				_cgfe := NewCT_RChoice()
				if _fabd := d.DecodeElement(&_cgfe.T, &_cefe); _fabd != nil {
					return _fabd
				}
				_bcea.Choice = append(_bcea.Choice, _cgfe)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u0052\u0020\u0025\u0076", _cefe.Name)
				if _dggc := d.Skip(); _dggc != nil {
					return _dggc
				}
			}
		case _g.EndElement:
			break _ccbc
		case _g.CharData:
		}
	}
	return nil
}

func NewCT_FType() *CT_FType { _dbb := &CT_FType{}; _dbb.ValAttr = ST_FType(1); return _dbb }

func (_eggc ST_Style) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_fddf := _g.Attr{}
	_fddf.Name = name
	switch _eggc {
	case ST_StyleUnset:
		_fddf.Value = ""
	case ST_StyleP:
		_fddf.Value = "\u0070"
	case ST_StyleB:
		_fddf.Value = "\u0062"
	case ST_StyleI:
		_fddf.Value = "\u0069"
	case ST_StyleBi:
		_fddf.Value = "\u0062\u0069"
	}
	return _fddf, nil
}

type ST_BreakBinSub byte

func (_eacga *ST_BreakBinSub) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_dgegf, _aadb := d.Token()
	if _aadb != nil {
		return _aadb
	}
	if _acfd, _eacgb := _dgegf.(_g.EndElement); _eacgb && _acfd.Name == start.Name {
		*_eacga = 1
		return nil
	}
	if _gfdg, _dcfc := _dgegf.(_g.CharData); !_dcfc {
		return _d.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dgegf)
	} else {
		switch string(_gfdg) {
		case "":
			*_eacga = 0
		case "\u002d\u002d":
			*_eacga = 1
		case "\u002d\u002b":
			*_eacga = 2
		case "\u002b\u002d":
			*_eacga = 3
		}
	}
	_dgegf, _aadb = d.Token()
	if _aadb != nil {
		return _aadb
	}
	if _gaaf, _bgbe := _dgegf.(_g.EndElement); _bgbe && _gaaf.Name == start.Name {
		return nil
	}
	return _d.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dgegf)
}

func (_cdfg *CT_Char) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	for _, _bcdb := range start.Attr {
		if _bcdb.Name.Local == "\u0076\u0061\u006c" {
			_fgd, _dgb := _bcdb.Value, error(nil)
			if _dgb != nil {
				return _dgb
			}
			_cdfg.ValAttr = _fgd
			continue
		}
	}
	for {
		_acef, _fge := d.Token()
		if _fge != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0043\u0068\u0061\u0072\u003a\u0020\u0025\u0073", _fge)
		}
		if _ggdd, _bge := _acef.(_g.EndElement); _bge && _ggdd.Name == start.Name {
			break
		}
	}
	return nil
}

func (_cege ST_Shp) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_cege.String(), start)
}

func (_afe *CT_BorderBox) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_afe.E = NewCT_OMathArg()
_ea:
	for {
		_def, _ac := d.Token()
		if _ac != nil {
			return _ac
		}
		switch _deb := _def.(type) {
		case _g.StartElement:
			switch _deb.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "b\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "b\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078\u0050\u0072"}:
				_afe.BorderBoxPr = NewCT_BorderBoxPr()
				if _fbe := d.DecodeElement(_afe.BorderBoxPr, &_deb); _fbe != nil {
					return _fbe
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _gcb := d.DecodeElement(_afe.E, &_deb); _gcb != nil {
					return _gcb
				}
			default:
				_f.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_B\u006f\u0072d\u0065\u0072\u0042\u006f\u0078\u0020\u0025\u0076", _deb.Name)
				if _ceb := d.Skip(); _ceb != nil {
					return _ceb
				}
			}
		case _g.EndElement:
			break _ea
		case _g.CharData:
		}
	}
	return nil
}

func (_aea *CT_BreakBin) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _aea.ValAttr != ST_BreakBinUnset {
		_gefc, _gdc := _aea.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _gdc != nil {
			return _gdc
		}
		start.Attr = append(start.Attr, _gefc)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

type CT_RadPr struct {
	DegHide *CT_OnOff
	CtrlPr  *CT_CtrlPr
}

func (_gbcc ST_BreakBinSub) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	return e.EncodeElement(_gbcc.String(), start)
}

// Validate validates the CT_NaryPr and its children
func (_ffgf *CT_NaryPr) Validate() error {
	return _ffgf.ValidateWithPath("\u0043T\u005f\u004e\u0061\u0072\u0079\u0050r")
}

func NewOMath() *OMath { _ceagd := &OMath{}; _ceagd.CT_OMath = *NewCT_OMath(); return _ceagd }

// Validate validates the CT_Acc and its children
func (_beg *CT_Acc) Validate() error {
	return _beg.ValidateWithPath("\u0043\u0054\u005f\u0041\u0063\u0063")
}

// ValidateWithPath validates the CT_BorderBoxPr and its children, prefixing error messages with path
func (_eb *CT_BorderBoxPr) ValidateWithPath(path string) error {
	if _eb.HideTop != nil {
		if _fda := _eb.HideTop.ValidateWithPath(path + "\u002f\u0048\u0069\u0064\u0065\u0054\u006f\u0070"); _fda != nil {
			return _fda
		}
	}
	if _eb.HideBot != nil {
		if _eac := _eb.HideBot.ValidateWithPath(path + "\u002f\u0048\u0069\u0064\u0065\u0042\u006f\u0074"); _eac != nil {
			return _eac
		}
	}
	if _eb.HideLeft != nil {
		if _gge := _eb.HideLeft.ValidateWithPath(path + "\u002fH\u0069\u0064\u0065\u004c\u0065\u0066t"); _gge != nil {
			return _gge
		}
	}
	if _eb.HideRight != nil {
		if _cdf := _eb.HideRight.ValidateWithPath(path + "\u002f\u0048\u0069\u0064\u0065\u0052\u0069\u0067\u0068\u0074"); _cdf != nil {
			return _cdf
		}
	}
	if _eb.StrikeH != nil {
		if _ccgc := _eb.StrikeH.ValidateWithPath(path + "\u002f\u0053\u0074\u0072\u0069\u006b\u0065\u0048"); _ccgc != nil {
			return _ccgc
		}
	}
	if _eb.StrikeV != nil {
		if _gfd := _eb.StrikeV.ValidateWithPath(path + "\u002f\u0053\u0074\u0072\u0069\u006b\u0065\u0056"); _gfd != nil {
			return _gfd
		}
	}
	if _eb.StrikeBLTR != nil {
		if _ff := _eb.StrikeBLTR.ValidateWithPath(path + "/\u0053\u0074\u0072\u0069\u006b\u0065\u0042\u004c\u0054\u0052"); _ff != nil {
			return _ff
		}
	}
	if _eb.StrikeTLBR != nil {
		if _fgb := _eb.StrikeTLBR.ValidateWithPath(path + "/\u0053\u0074\u0072\u0069\u006b\u0065\u0054\u004c\u0042\u0052"); _fgb != nil {
			return _fgb
		}
	}
	if _eb.CtrlPr != nil {
		if _aa := _eb.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _aa != nil {
			return _aa
		}
	}
	return nil
}

func NewCT_LimLow() *CT_LimLow {
	_eadd := &CT_LimLow{}
	_eadd.E = NewCT_OMathArg()
	_eadd.Lim = NewCT_OMathArg()
	return _eadd
}

// Validate validates the CT_DPr and its children
func (_aad *CT_DPr) Validate() error {
	return _aad.ValidateWithPath("\u0043\u0054\u005f\u0044\u0050\u0072")
}

func (_bddbb *ST_BreakBin) UnmarshalXMLAttr(attr _g.Attr) error {
	switch attr.Value {
	case "":
		*_bddbb = 0
	case "\u0062\u0065\u0066\u006f\u0072\u0065":
		*_bddbb = 1
	case "\u0061\u0066\u0074e\u0072":
		*_bddbb = 2
	case "\u0072\u0065\u0070\u0065\u0061\u0074":
		*_bddbb = 3
	}
	return nil
}

func NewCT_MCPr() *CT_MCPr { _agaf := &CT_MCPr{}; return _agaf }

// ValidateWithPath validates the CT_RPR and its children, prefixing error messages with path
func (_ffag *CT_RPR) ValidateWithPath(path string) error {
	if _ffag.Lit != nil {
		if _agge := _ffag.Lit.ValidateWithPath(path + "\u002f\u004c\u0069\u0074"); _agge != nil {
			return _agge
		}
	}
	if _ffag.Choice != nil {
		if _abcb := _ffag.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _abcb != nil {
			return _abcb
		}
	}
	if _ffag.Brk != nil {
		if _cfdd := _ffag.Brk.ValidateWithPath(path + "\u002f\u0042\u0072\u006b"); _cfdd != nil {
			return _cfdd
		}
	}
	if _ffag.Aln != nil {
		if _cbbd := _ffag.Aln.ValidateWithPath(path + "\u002f\u0041\u006c\u006e"); _cbbd != nil {
			return _cbbd
		}
	}
	return nil
}

func (_gggd *CT_XAlign) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_gggd.ValAttr = _dc.ST_XAlign(1)
	for _, _fbbe := range start.Attr {
		if _fbbe.Name.Local == "\u0076\u0061\u006c" {
			_gggd.ValAttr.UnmarshalXMLAttr(_fbbe)
			continue
		}
	}
	for {
		_fafg, _gaece := d.Token()
		if _gaece != nil {
			return _d.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0058\u0041\u006ci\u0067\u006e\u003a\u0020\u0025\u0073", _gaece)
		}
		if _gdgge, _ccbe := _fafg.(_g.EndElement); _ccbe && _gdgge.Name == start.Name {
			break
		}
	}
	return nil
}

func NewCT_DPr() *CT_DPr { _ede := &CT_DPr{}; return _ede }

func (_cgefc *CT_TopBot) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_dbba, _afcg := _cgefc.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _afcg != nil {
		return _afcg
	}
	start.Attr = append(start.Attr, _dbba)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func NewCT_BorderBox() *CT_BorderBox { _cg := &CT_BorderBox{}; _cg.E = NewCT_OMathArg(); return _cg }

func (_bfca ST_FType) String() string {
	switch _bfca {
	case 0:
		return ""
	case 1:
		return "\u0062\u0061\u0072"
	case 2:
		return "\u0073\u006b\u0077"
	case 3:
		return "\u006c\u0069\u006e"
	case 4:
		return "\u006e\u006f\u0042a\u0072"
	}
	return ""
}

// ValidateWithPath validates the CT_BoxPr and its children, prefixing error messages with path
func (_aed *CT_BoxPr) ValidateWithPath(path string) error {
	if _aed.OpEmu != nil {
		if _eebf := _aed.OpEmu.ValidateWithPath(path + "\u002f\u004f\u0070\u0045\u006d\u0075"); _eebf != nil {
			return _eebf
		}
	}
	if _aed.NoBreak != nil {
		if _aca := _aed.NoBreak.ValidateWithPath(path + "\u002f\u004e\u006f\u0042\u0072\u0065\u0061\u006b"); _aca != nil {
			return _aca
		}
	}
	if _aed.Diff != nil {
		if _ebe := _aed.Diff.ValidateWithPath(path + "\u002f\u0044\u0069f\u0066"); _ebe != nil {
			return _ebe
		}
	}
	if _aed.Brk != nil {
		if _dfd := _aed.Brk.ValidateWithPath(path + "\u002f\u0042\u0072\u006b"); _dfd != nil {
			return _dfd
		}
	}
	if _aed.Aln != nil {
		if _gfef := _aed.Aln.ValidateWithPath(path + "\u002f\u0041\u006c\u006e"); _gfef != nil {
			return _gfef
		}
	}
	if _aed.CtrlPr != nil {
		if _cfc := _aed.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _cfc != nil {
			return _cfc
		}
	}
	return nil
}

func (_cee *CT_LimLow) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cee.E = NewCT_OMathArg()
	_cee.Lim = NewCT_OMathArg()
_fdfa:
	for {
		_afc, _dagf := d.Token()
		if _dagf != nil {
			return _dagf
		}
		switch _gcdd := _afc.(type) {
		case _g.StartElement:
			switch _gcdd.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077\u0050\u0072"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d\u004c\u006f\u0077\u0050\u0072"}:
				_cee.LimLowPr = NewCT_LimLowPr()
				if _cec := d.DecodeElement(_cee.LimLowPr, &_gcdd); _cec != nil {
					return _cec
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				if _fbgc := d.DecodeElement(_cee.E, &_gcdd); _fbgc != nil {
					return _fbgc
				}
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u006c\u0069\u006d"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u006c\u0069\u006d"}:
				if _dcfe := d.DecodeElement(_cee.Lim, &_gcdd); _dcfe != nil {
					return _dcfe
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004c\u0069\u006dL\u006f\u0077 \u0025\u0076", _gcdd.Name)
				if _bdcb := d.Skip(); _bdcb != nil {
					return _bdcb
				}
			}
		case _g.EndElement:
			break _fdfa
		case _g.CharData:
		}
	}
	return nil
}

func NewCT_RadPr() *CT_RadPr { _fdaea := &CT_RadPr{}; return _fdaea }

func (_gece ST_BreakBin) ValidateWithPath(path string) error {
	switch _gece {
	case 0, 1, 2, 3:
	default:
		return _d.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gece))
	}
	return nil
}

type CT_F struct {
	FPr *CT_FPr
	Num *CT_OMathArg
	Den *CT_OMathArg
}

func (_adcae *EG_ScriptStyle) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Name.Local = "\u006d\u003aE\u0047\u005f\u0053c\u0072\u0069\u0070\u0074\u0053\u0074\u0079\u006c\u0065"
	if _adcae.Scr != nil {
		_bfbba := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073c\u0072"}}
		e.EncodeElement(_adcae.Scr, _bfbba)
	}
	if _adcae.Sty != nil {
		_agfbd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0073t\u0079"}}
		e.EncodeElement(_adcae.Sty, _agfbd)
	}
	return nil
}

func (_fgab *CT_OMathArgPr) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_baeaa:
	for {
		_fgeb, _bdef := d.Token()
		if _bdef != nil {
			return _bdef
		}
		switch _fdda := _fgeb.(type) {
		case _g.StartElement:
			switch _fdda.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0061\u0072\u0067S\u007a"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0061\u0072\u0067S\u007a"}:
				_fgab.ArgSz = NewCT_Integer2()
				if _cabb := d.DecodeElement(_fgab.ArgSz, &_fdda); _cabb != nil {
					return _cabb
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004fM\u0061\u0074\u0068\u0041\u0072\u0067\u0050\u0072 \u0025\u0076", _fdda.Name)
				if _bfcg := d.Skip(); _bfcg != nil {
					return _bfcg
				}
			}
		case _g.EndElement:
			break _baeaa
		case _g.CharData:
		}
	}
	return nil
}

func NewCT_PhantPr() *CT_PhantPr { _ggbc := &CT_PhantPr{}; return _ggbc }

type ST_LimLoc byte

type CT_Nary struct {
	NaryPr *CT_NaryPr
	Sub    *CT_OMathArg
	Sup    *CT_OMathArg
	E      *CT_OMathArg
}

// Validate validates the CT_LimUpp and its children
func (_dbbff *CT_LimUpp) Validate() error {
	return _dbbff.ValidateWithPath("\u0043T\u005f\u004c\u0069\u006d\u0055\u0070p")
}

// ValidateWithPath validates the CT_SpacingRule and its children, prefixing error messages with path
func (_fecca *CT_SpacingRule) ValidateWithPath(path string) error {
	if _fecca.ValAttr < 0 {
		return _d.Errorf("%\u0073\u002f\u006d\u002e\u0056\u0061l\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _fecca.ValAttr)
	}
	if _fecca.ValAttr > 4 {
		return _d.Errorf("%\u0073\u002f\u006d\u002e\u0056\u0061l\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003c=\u0020\u0034\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _fecca.ValAttr)
	}
	return nil
}

type CT_LimLow struct {
	LimLowPr *CT_LimLowPr
	E        *CT_OMathArg
	Lim      *CT_OMathArg
}

func (_dgcb *CT_RPR) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _dgcb.Lit != nil {
		_cdbbg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u006ci\u0074"}}
		e.EncodeElement(_dgcb.Lit, _cdbbg)
	}
	if _dgcb.Choice != nil {
		_dgcb.Choice.MarshalXML(e, _g.StartElement{})
	}
	if _dgcb.Brk != nil {
		_bbbd := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0062r\u006b"}}
		e.EncodeElement(_dgcb.Brk, _bbbd)
	}
	if _dgcb.Aln != nil {
		_bfbg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0061l\u006e"}}
		e.EncodeElement(_dgcb.Aln, _bfbg)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_gfgf ST_BreakBinSub) MarshalXMLAttr(name _g.Name) (_g.Attr, error) {
	_edcb := _g.Attr{}
	_edcb.Name = name
	switch _gfgf {
	case ST_BreakBinSubUnset:
		_edcb.Value = ""
	case ST_BreakBinSub__:
		_edcb.Value = "\u002d\u002d"
	case ST_BreakBinSub___:
		_edcb.Value = "\u002d\u002b"
	case ST_BreakBinSub____:
		_edcb.Value = "\u002b\u002d"
	}
	return _edcb, nil
}

func NewCT_BreakBinSub() *CT_BreakBinSub { _fae := &CT_BreakBinSub{}; return _fae }

// Validate validates the CT_Nary and its children
func (_adbe *CT_Nary) Validate() error {
	return _adbe.ValidateWithPath("\u0043T\u005f\u004e\u0061\u0072\u0079")
}

func (_ffg *CT_FType) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_ffg.ValAttr = ST_FType(1)
	for _, _gba := range start.Attr {
		if _gba.Name.Local == "\u0076\u0061\u006c" {
			_ffg.ValAttr.UnmarshalXMLAttr(_gba)
			continue
		}
	}
	for {
		_fcba, _ebea := d.Token()
		if _ebea != nil {
			return _d.Errorf("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fF\u0054\u0079\u0070\u0065: \u0025\u0073", _ebea)
		}
		if _gdgg, _adcc := _fcba.(_g.EndElement); _adcc && _gdgg.Name == start.Name {
			break
		}
	}
	return nil
}

func (_abff *CT_Integer255) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	start.Attr = append(start.Attr, _g.Attr{Name: _g.Name{Local: "\u006d\u003a\u0076a\u006c"}, Value: _d.Sprintf("\u0025\u0076", _abff.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_cadac *CT_SpacingRule) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
	_cadac.ValAttr = 0
	for _, _fbgeb := range start.Attr {
		if _fbgeb.Name.Local == "\u0076\u0061\u006c" {
			_baac, _addba := _be.ParseInt(_fbgeb.Value, 10, 64)
			if _addba != nil {
				return _addba
			}
			_cadac.ValAttr = _baac
			continue
		}
	}
	for {
		_ddcf, _bgea := d.Token()
		if _bgea != nil {
			return _d.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fS\u0070\u0061\u0063\u0069\u006e\u0067\u0052\u0075\u006c\u0065:\u0020\u0025\u0073", _bgea)
		}
		if _fdgf, _fbba := _ddcf.(_g.EndElement); _fbba && _fdgf.Name == start.Name {
			break
		}
	}
	return nil
}

func (_agbd *CT_Rad) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _agbd.RadPr != nil {
		_cceg := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0072\u0061\u0064\u0050\u0072"}}
		e.EncodeElement(_agbd.RadPr, _cceg)
	}
	_bfg := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0064e\u0067"}}
	e.EncodeElement(_agbd.Deg, _bfg)
	_fedaf := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0065"}}
	e.EncodeElement(_agbd.E, _fedaf)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

const (
	ST_ShpUnset    ST_Shp = 0
	ST_ShpCentered ST_Shp = 1
	ST_ShpMatch    ST_Shp = 2
)

// Validate validates the CT_SSubPr and its children
func (_ccfe *CT_SSubPr) Validate() error {
	return _ccfe.ValidateWithPath("\u0043T\u005f\u0053\u0053\u0075\u0062\u0050r")
}

type ST_FType byte

func (_adccfg *CT_OMath) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _adccfg.EG_OMathMathElements != nil {
		for _, _eeca := range _adccfg.EG_OMathMathElements {
			_eeca.MarshalXML(e, _g.StartElement{})
		}
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_cfed *CT_Shp) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	_dbbb, _dfcd := _cfed.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
	if _dfcd != nil {
		return _dfcd
	}
	start.Attr = append(start.Attr, _dbbb)
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

func (_cab *CT_MR) UnmarshalXML(d *_g.Decoder, start _g.StartElement) error {
_gade:
	for {
		_ebc, _acb := d.Token()
		if _acb != nil {
			return _acb
		}
		switch _aedag := _ebc.(type) {
		case _g.StartElement:
			switch _aedag.Name {
			case _g.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", Local: "\u0065"}, _g.Name{Space: "\u0068\u0074t\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u006d\u0061\u0074\u0068", Local: "\u0065"}:
				_edg := NewCT_OMathArg()
				if _fbaf := d.DecodeElement(_edg, &_aedag); _fbaf != nil {
					return _fbaf
				}
				_cab.E = append(_cab.E, _edg)
			default:
				_f.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0043\u0054\u005fM\u0052 \u0025\u0076", _aedag.Name)
				if _acac := d.Skip(); _acac != nil {
					return _acac
				}
			}
		case _g.EndElement:
			break _gade
		case _g.CharData:
		}
	}
	return nil
}

// Validate validates the CT_EqArrPr and its children
func (_dgba *CT_EqArrPr) Validate() error {
	return _dgba.ValidateWithPath("\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072\u0050\u0072")
}

// Validate validates the CT_MR and its children
func (_cfece *CT_MR) Validate() error { return _cfece.ValidateWithPath("\u0043\u0054\u005fM\u0052") }

func (_fege *CT_OMathArg) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	e.EncodeToken(start)
	if _fege.ArgPr != nil {
		_ebbb := _g.StartElement{Name: _g.Name{Local: "\u006d:\u0061\u0072\u0067\u0050\u0072"}}
		e.EncodeElement(_fege.ArgPr, _ebbb)
	}
	if _fege.EG_OMathMathElements != nil {
		for _, _edfe := range _fege.EG_OMathMathElements {
			_edfe.MarshalXML(e, _g.StartElement{})
		}
	}
	if _fege.CtrlPr != nil {
		_ccea := _g.StartElement{Name: _g.Name{Local: "\u006d\u003a\u0063\u0074\u0072\u006c\u0050\u0072"}}
		e.EncodeElement(_fege.CtrlPr, _ccea)
	}
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_LimUppPr and its children, prefixing error messages with path
func (_dggg *CT_LimUppPr) ValidateWithPath(path string) error {
	if _dggg.CtrlPr != nil {
		if _egfb := _dggg.CtrlPr.ValidateWithPath(path + "\u002fC\u0074\u0072\u006c\u0050\u0072"); _egfb != nil {
			return _egfb
		}
	}
	return nil
}

func (_ccf *CT_OMathJc) MarshalXML(e *_g.Encoder, start _g.StartElement) error {
	if _ccf.ValAttr != ST_JcUnset {
		_baeaaf, _fbdb := _ccf.ValAttr.MarshalXMLAttr(_g.Name{Local: "\u006d\u003a\u0076a\u006c"})
		if _fbdb != nil {
			return _fbdb
		}
		start.Attr = append(start.Attr, _baeaaf)
	}
	e.EncodeToken(start)
	e.EncodeToken(_g.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_SSub and its children, prefixing error messages with path
func (_befg *CT_SSub) ValidateWithPath(path string) error {
	if _befg.SSubPr != nil {
		if _ebbc := _befg.SSubPr.ValidateWithPath(path + "\u002fS\u0053\u0075\u0062\u0050\u0072"); _ebbc != nil {
			return _ebbc
		}
	}
	if _eadf := _befg.E.ValidateWithPath(path + "\u002f\u0045"); _eadf != nil {
		return _eadf
	}
	if _faec := _befg.Sub.ValidateWithPath(path + "\u002f\u0053\u0075\u0062"); _faec != nil {
		return _faec
	}
	return nil
}

func init() {
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032\u0035\u0035", NewCT_Integer255)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u0049\u006e\u0074\u0065\u0067\u0065\u0072\u0032", NewCT_Integer2)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0053\u0070\u0061\u0063\u0069\u006eg\u0052\u0075\u006c\u0065", NewCT_SpacingRule)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005fU\u006e\u0053\u0069\u0067\u006ee\u0064\u0049n\u0074\u0065\u0067\u0065\u0072", NewCT_UnSignedInteger)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0043\u0068\u0061\u0072", NewCT_Char)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u006e\u004f\u0066\u0066", NewCT_OnOff)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0074\u0072\u0069\u006eg", NewCT_String)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0058\u0041\u006c\u0069\u0067n", NewCT_XAlign)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0059\u0041\u006c\u0069\u0067n", NewCT_YAlign)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0053\u0068\u0070", NewCT_Shp)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0046\u0054\u0079\u0070\u0065", NewCT_FType)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004c\u0069\u006d\u004c\u006fc", NewCT_LimLoc)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0054\u006f\u0070\u0042\u006ft", NewCT_TopBot)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0063\u0072\u0069\u0070t", NewCT_Script)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065", NewCT_Style)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004d\u0061\u006e\u0075\u0061\u006cB\u0072\u0065\u0061\u006b", NewCT_ManualBreak)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0052\u0050\u0052", NewCT_RPR)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0054\u0065\u0078\u0074", NewCT_Text)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0052", NewCT_R)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0043\u0074\u0072\u006c\u0050r", NewCT_CtrlPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0041\u0063\u0063\u0050\u0072", NewCT_AccPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0041\u0063\u0063", NewCT_Acc)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u0061\u0072\u0050\u0072", NewCT_BarPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u0061\u0072", NewCT_Bar)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u006f\u0078\u0050\u0072", NewCT_BoxPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u006f\u0078", NewCT_Box)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u006f\u0072\u0064\u0065\u0072B\u006f\u0078\u0050\u0072", NewCT_BorderBoxPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fB\u006f\u0072\u0064\u0065\u0072\u0042\u006f\u0078", NewCT_BorderBox)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0044\u0050\u0072", NewCT_DPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0044", NewCT_D)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072\u0050\u0072", NewCT_EqArrPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0045\u0071\u0041\u0072\u0072", NewCT_EqArr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0046\u0050\u0072", NewCT_FPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0046", NewCT_F)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0046\u0075\u006e\u0063\u0050r", NewCT_FuncPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0046\u0075\u006e\u0063", NewCT_Func)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072\u0050\u0072", NewCT_GroupChrPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u0072", NewCT_GroupChr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u004c\u0069\u006d\u004c\u006f\u0077\u0050\u0072", NewCT_LimLowPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004c\u0069\u006d\u004c\u006fw", NewCT_LimLow)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u004c\u0069\u006d\u0055\u0070\u0070\u0050\u0072", NewCT_LimUppPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004c\u0069\u006d\u0055\u0070p", NewCT_LimUpp)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004d\u0043\u0050\u0072", NewCT_MCPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fM\u0043", NewCT_MC)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004d\u0043\u0053", NewCT_MCS)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004d\u0050\u0072", NewCT_MPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fM\u0052", NewCT_MR)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004d", NewCT_M)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004e\u0061\u0072\u0079\u0050r", NewCT_NaryPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004e\u0061\u0072\u0079", NewCT_Nary)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074\u0050\u0072", NewCT_PhantPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0050\u0068\u0061\u006e\u0074", NewCT_Phant)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0052\u0061\u0064\u0050\u0072", NewCT_RadPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0052\u0061\u0064", NewCT_Rad)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0050\u0072\u0065\u0050r", NewCT_SPrePr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0050\u0072\u0065", NewCT_SPre)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0053\u0075\u0062\u0050r", NewCT_SSubPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0053\u0075\u0062", NewCT_SSub)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fS\u0053\u0075\u0062\u0053\u0075\u0070\u0050\u0072", NewCT_SSubSupPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0053\u0053\u0075\u0062\u0053\u0075\u0070", NewCT_SSubSup)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0053\u0075\u0070\u0050r", NewCT_SSupPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u0053\u0053\u0075\u0070", NewCT_SSup)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067\u0050\u0072", NewCT_OMathArgPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0041\u0072\u0067", NewCT_OMathArg)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u004a\u0063", NewCT_OMathJc)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068\u0050a\u0072\u0061\u0050\u0072", NewCT_OMathParaPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005fT\u0077\u0069\u0070\u0073\u004d\u0065\u0061\u0073\u0075\u0072\u0065", NewCT_TwipsMeasure)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "C\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042\u0069\u006e", NewCT_BreakBin)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0042i\u006e\u0053\u0075\u0062", NewCT_BreakBinSub)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043T\u005f\u004d\u0061\u0074\u0068\u0050r", NewCT_MathPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005fO\u004d\u0061\u0074\u0068\u0050\u0061\u0072\u0061", NewCT_OMathPara)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0043\u0054\u005f\u004f\u004d\u0061\u0074\u0068", NewCT_OMath)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u006d\u0061\u0074\u0068\u0050\u0072", NewMathPr)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u006fM\u0061\u0074\u0068\u0050\u0061\u0072a", NewOMathPara)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u006f\u004d\u0061t\u0068", NewOMath)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0045\u0047\u005f\u0053\u0063\u0072\u0069\u0070\u0074S\u0074\u0079\u006c\u0065", NewEG_ScriptStyle)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "E\u0047_\u004f\u004d\u0061\u0074\u0068\u004d\u0061\u0074h\u0045\u006c\u0065\u006den\u0074\u0073", NewEG_OMathMathElements)
	_c.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075m\u0065\u006e\u0074\u002f\u0032\u00300\u0036\u002f\u006da\u0074\u0068", "\u0045\u0047_\u004f\u004d\u0061t\u0068\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073", NewEG_OMathElements)
}
