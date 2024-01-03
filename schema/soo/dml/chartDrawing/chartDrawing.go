package chartDrawing

import (
	_ed "encoding/xml"
	_b "fmt"
	_a "strconv"

	_bd "github.com/arcpd/msword"
	_d "github.com/arcpd/msword/common/logger"
	_c "github.com/arcpd/msword/schema/soo/dml"
)

// Validate validates the CT_Connector and its children
func (_aa *CT_Connector) Validate() error {
	return _aa.ValidateWithPath("\u0043\u0054\u005fC\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072")
}

// Validate validates the CT_Shape and its children
func (_cef *CT_Shape) Validate() error {
	return _cef.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065")
}

func (_ged *CT_Shape) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	if _ged.MacroAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _b.Sprintf("\u0025\u0076", *_ged.MacroAttr)})
	}
	if _ged.TextlinkAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u0074\u0065\u0078\u0074\u006c\u0069\u006e\u006b"}, Value: _b.Sprintf("\u0025\u0076", *_ged.TextlinkAttr)})
	}
	if _ged.FLocksTextAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u0066\u004c\u006f\u0063\u006b\u0073\u0054\u0065\u0078\u0074"}, Value: _b.Sprintf("\u0025\u0064", _bgf(*_ged.FLocksTextAttr))})
	}
	if _ged.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _b.Sprintf("\u0025\u0064", _bgf(*_ged.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_cccf := _ed.StartElement{Name: _ed.Name{Local: "\u006e\u0076\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_ged.NvSpPr, _cccf)
	_gfc := _ed.StartElement{Name: _ed.Name{Local: "\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_ged.SpPr, _gfc)
	if _ged.Style != nil {
		_dgg := _ed.StartElement{Name: _ed.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_ged.Style, _dgg)
	}
	if _ged.TxBody != nil {
		_ccbg := _ed.StartElement{Name: _ed.Name{Local: "\u0074\u0078\u0042\u006f\u0064\u0079"}}
		e.EncodeElement(_ged.TxBody, _ccbg)
	}
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func NewCT_ShapeNonVisual() *CT_ShapeNonVisual {
	_cdea := &CT_ShapeNonVisual{}
	_cdea.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_cdea.CNvSpPr = _c.NewCT_NonVisualDrawingShapeProps()
	return _cdea
}

func (_fed *CT_Marker) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_fed.X = 0.0
	_fed.Y = 0.0
_ggg:
	for {
		_ffgb, _gecb := d.Token()
		if _gecb != nil {
			return _gecb
		}
		switch _eba := _ffgb.(type) {
		case _ed.StartElement:
			switch _eba.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0078"}:
				if _gfef := d.DecodeElement(&_fed.X, &_eba); _gfef != nil {
					return _gfef
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0079"}:
				if _eca := d.DecodeElement(&_fed.Y, &_eba); _eca != nil {
					return _eca
				}
			default:
				_d.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004d\u0061\u0072k\u0065\u0072 \u0025\u0076", _eba.Name)
				if _dbe := d.Skip(); _dbe != nil {
					return _dbe
				}
			}
		case _ed.EndElement:
			break _ggg
		case _ed.CharData:
		}
	}
	return nil
}

func (_ecc *CT_ConnectorNonVisual) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_ecc.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_ecc.CNvCxnSpPr = _c.NewCT_NonVisualConnectorProperties()
_bde:
	for {
		_geb, _eb := d.Token()
		if _eb != nil {
			return _eb
		}
		switch _fcf := _geb.(type) {
		case _ed.StartElement:
			switch _fcf.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _gbgc := d.DecodeElement(_ecc.CNvPr, &_fcf); _gbgc != nil {
					return _gbgc
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}:
				if _dge := d.DecodeElement(_ecc.CNvCxnSpPr, &_fcf); _dge != nil {
					return _dge
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075n\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u006f\u006e C\u0054\u005f\u0043\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _fcf.Name)
				if _dac := d.Skip(); _dac != nil {
					return _dac
				}
			}
		case _ed.EndElement:
			break _bde
		case _ed.CharData:
		}
	}
	return nil
}

func (_dcb *CT_Picture) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	if _dcb.MacroAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _b.Sprintf("\u0025\u0076", *_dcb.MacroAttr)})
	}
	if _dcb.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _b.Sprintf("\u0025\u0064", _bgf(*_dcb.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_aff := _ed.StartElement{Name: _ed.Name{Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_dcb.NvPicPr, _aff)
	_ada := _ed.StartElement{Name: _ed.Name{Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}}
	e.EncodeElement(_dcb.BlipFill, _ada)
	_fefd := _ed.StartElement{Name: _ed.Name{Local: "\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_dcb.SpPr, _fefd)
	if _dcb.Style != nil {
		_cfcb := _ed.StartElement{Name: _ed.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_dcb.Style, _cfcb)
	}
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

type CT_AbsSizeAnchor struct {
	From   *CT_Marker
	Ext    *_c.CT_PositiveSize2D
	Choice *EG_ObjectChoicesChoice
}

func NewCT_PictureNonVisual() *CT_PictureNonVisual {
	_bgee := &CT_PictureNonVisual{}
	_bgee.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_bgee.CNvPicPr = _c.NewCT_NonVisualPictureProperties()
	return _bgee
}

// ValidateWithPath validates the CT_GroupShape and its children, prefixing error messages with path
func (_bece *CT_GroupShape) ValidateWithPath(path string) error {
	if _fad := _bece.NvGrpSpPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _fad != nil {
		return _fad
	}
	if _dca := _bece.GrpSpPr.ValidateWithPath(path + "\u002f\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _dca != nil {
		return _dca
	}
	for _gac, _afd := range _bece.Choice {
		if _cdef := _afd.ValidateWithPath(_b.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _gac)); _cdef != nil {
			return _cdef
		}
	}
	return nil
}

func (_gcd *CT_GroupShapeChoice) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	if _gcd.Sp != nil {
		_fbbc := _ed.StartElement{Name: _ed.Name{Local: "\u0073\u0070"}}
		for _, _aae := range _gcd.Sp {
			e.EncodeElement(_aae, _fbbc)
		}
	}
	if _gcd.GrpSp != nil {
		_ebc := _ed.StartElement{Name: _ed.Name{Local: "\u0067\u0072\u0070S\u0070"}}
		for _, _ece := range _gcd.GrpSp {
			e.EncodeElement(_ece, _ebc)
		}
	}
	if _gcd.GraphicFrame != nil {
		_bbg := _ed.StartElement{Name: _ed.Name{Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}}
		for _, _baa := range _gcd.GraphicFrame {
			e.EncodeElement(_baa, _bbg)
		}
	}
	if _gcd.CxnSp != nil {
		_eeb := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u0078\u006eS\u0070"}}
		for _, _bafa := range _gcd.CxnSp {
			e.EncodeElement(_bafa, _eeb)
		}
	}
	if _gcd.Pic != nil {
		_eacc := _ed.StartElement{Name: _ed.Name{Local: "\u0070\u0069\u0063"}}
		for _, _bae := range _gcd.Pic {
			e.EncodeElement(_bae, _eacc)
		}
	}
	return nil
}

// Validate validates the CT_ShapeNonVisual and its children
func (_gbgdd *CT_ShapeNonVisual) Validate() error {
	return _gbgdd.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c")
}

func NewEG_ObjectChoicesChoice() *EG_ObjectChoicesChoice {
	_feg := &EG_ObjectChoicesChoice{}
	return _feg
}

func NewCT_Shape() *CT_Shape {
	_dad := &CT_Shape{}
	_dad.NvSpPr = NewCT_ShapeNonVisual()
	_dad.SpPr = _c.NewCT_ShapeProperties()
	return _dad
}

func NewCT_Drawing() *CT_Drawing { _aec := &CT_Drawing{}; return _aec }

func (_feb *CT_GraphicFrameNonVisual) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	e.EncodeToken(start)
	_fbb := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_feb.CNvPr, _fbb)
	_eae := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"}}
	e.EncodeElement(_feb.CNvGraphicFramePr, _eae)
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func (_cc *CT_AbsSizeAnchor) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	e.EncodeToken(start)
	_fb := _ed.StartElement{Name: _ed.Name{Local: "\u0066\u0072\u006f\u006d"}}
	e.EncodeElement(_cc.From, _fb)
	_dg := _ed.StartElement{Name: _ed.Name{Local: "\u0065\u0078\u0074"}}
	e.EncodeElement(_cc.Ext, _dg)
	if _cc.Choice != nil {
		_cc.Choice.MarshalXML(e, _ed.StartElement{})
	}
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func (_aeb *CT_GroupShapeNonVisual) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_aeb.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_aeb.CNvGrpSpPr = _c.NewCT_NonVisualGroupDrawingShapeProps()
_caa:
	for {
		_aegg, _gea := d.Token()
		if _gea != nil {
			return _gea
		}
		switch _gaf := _aegg.(type) {
		case _ed.StartElement:
			switch _gaf.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _gfe := d.DecodeElement(_aeb.CNvPr, &_gaf); _gfe != nil {
					return _gfe
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}:
				if _dea := d.DecodeElement(_aeb.CNvGrpSpPr, &_gaf); _dea != nil {
					return _dea
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070p\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0047\u0072\u006f\u0075p\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _gaf.Name)
				if _gde := d.Skip(); _gde != nil {
					return _gde
				}
			}
		case _ed.EndElement:
			break _caa
		case _ed.CharData:
		}
	}
	return nil
}

func NewCT_GraphicFrame() *CT_GraphicFrame {
	_gd := &CT_GraphicFrame{}
	_gd.NvGraphicFramePr = NewCT_GraphicFrameNonVisual()
	_gd.Xfrm = _c.NewCT_Transform2D()
	_gd.Graphic = _c.NewGraphic()
	return _gd
}

// ValidateWithPath validates the CT_ConnectorNonVisual and its children, prefixing error messages with path
func (_cgf *CT_ConnectorNonVisual) ValidateWithPath(path string) error {
	if _acd := _cgf.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _acd != nil {
		return _acd
	}
	if _fe := _cgf.CNvCxnSpPr.ValidateWithPath(path + "/\u0043\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"); _fe != nil {
		return _fe
	}
	return nil
}

type EG_ObjectChoicesChoice struct {
	Sp           *CT_Shape
	GrpSp        *CT_GroupShape
	GraphicFrame *CT_GraphicFrame
	CxnSp        *CT_Connector
	Pic          *CT_Picture
}

// Validate validates the CT_GroupShapeNonVisual and its children
func (_afb *CT_GroupShapeNonVisual) Validate() error {
	return _afb.ValidateWithPath("\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075a\u006c")
}

// ValidateWithPath validates the CT_GroupShapeNonVisual and its children, prefixing error messages with path
func (_bge *CT_GroupShapeNonVisual) ValidateWithPath(path string) error {
	if _bfe := _bge.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _bfe != nil {
		return _bfe
	}
	if _ebg := _bge.CNvGrpSpPr.ValidateWithPath(path + "/\u0043\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _ebg != nil {
		return _ebg
	}
	return nil
}

func (_gbg *CT_Connector) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	if _gbg.MacroAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _b.Sprintf("\u0025\u0076", *_gbg.MacroAttr)})
	}
	if _gbg.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _b.Sprintf("\u0025\u0064", _bgf(*_gbg.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_fbf := _ed.StartElement{Name: _ed.Name{Local: "\u006ev\u0043\u0078\u006e\u0053\u0070\u0050r"}}
	e.EncodeElement(_gbg.NvCxnSpPr, _fbf)
	_cg := _ed.StartElement{Name: _ed.Name{Local: "\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_gbg.SpPr, _cg)
	if _gbg.Style != nil {
		_ea := _ed.StartElement{Name: _ed.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_gbg.Style, _ea)
	}
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the EG_ObjectChoicesChoice and its children, prefixing error messages with path
func (_fgc *EG_ObjectChoicesChoice) ValidateWithPath(path string) error {
	if _fgc.Sp != nil {
		if _dffb := _fgc.Sp.ValidateWithPath(path + "\u002f\u0053\u0070"); _dffb != nil {
			return _dffb
		}
	}
	if _fgc.GrpSp != nil {
		if _cacd := _fgc.GrpSp.ValidateWithPath(path + "\u002f\u0047\u0072\u0070\u0053\u0070"); _cacd != nil {
			return _cacd
		}
	}
	if _fgc.GraphicFrame != nil {
		if _baff := _fgc.GraphicFrame.ValidateWithPath(path + "\u002f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"); _baff != nil {
			return _baff
		}
	}
	if _fgc.CxnSp != nil {
		if _aab := _fgc.CxnSp.ValidateWithPath(path + "\u002f\u0043\u0078\u006e\u0053\u0070"); _aab != nil {
			return _aab
		}
	}
	if _fgc.Pic != nil {
		if _cfgc := _fgc.Pic.ValidateWithPath(path + "\u002f\u0050\u0069\u0063"); _cfgc != nil {
			return _cfgc
		}
	}
	return nil
}

func (_fbde *CT_GraphicFrame) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_fbde.NvGraphicFramePr = NewCT_GraphicFrameNonVisual()
	_fbde.Xfrm = _c.NewCT_Transform2D()
	_fbde.Graphic = _c.NewGraphic()
	for _, _cfc := range start.Attr {
		if _cfc.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_ba, _cgd := _cfc.Value, error(nil)
			if _cgd != nil {
				return _cgd
			}
			_fbde.MacroAttr = &_ba
			continue
		}
		if _cfc.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_bcc, _gba := _a.ParseBool(_cfc.Value)
			if _gba != nil {
				return _gba
			}
			_fbde.FPublishedAttr = &_bcc
			continue
		}
	}
_efe:
	for {
		_gcg, _fec := d.Token()
		if _fec != nil {
			return _fec
		}
		switch _bdc := _gcg.(type) {
		case _ed.StartElement:
			switch _bdc.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006e\u0076G\u0072\u0061\u0070h\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0050\u0072"}:
				if _eee := d.DecodeElement(_fbde.NvGraphicFramePr, &_bdc); _eee != nil {
					return _eee
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0078\u0066\u0072\u006d"}:
				if _ccce := d.DecodeElement(_fbde.Xfrm, &_bdc); _ccce != nil {
					return _ccce
				}
			case _ed.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0067r\u0061\u0070\u0068\u0069\u0063"}, _ed.Name{Space: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072g\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u006d\u0061\u0069\u006e", Local: "\u0067r\u0061\u0070\u0068\u0069\u0063"}:
				if _fea := d.DecodeElement(_fbde.Graphic, &_bdc); _fea != nil {
					return _fea
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0020\u0025\u0076", _bdc.Name)
				if _egc := d.Skip(); _egc != nil {
					return _egc
				}
			}
		case _ed.EndElement:
			break _efe
		case _ed.CharData:
		}
	}
	return nil
}

func (_fbc *CT_Drawing) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	start.Name.Local = "\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067"
	e.EncodeToken(start)
	if _fbc.EG_Anchor != nil {
		for _, _cda := range _fbc.EG_Anchor {
			_cda.MarshalXML(e, _ed.StartElement{})
		}
	}
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func (_abd *CT_GroupShape) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	e.EncodeToken(start)
	_gga := _ed.StartElement{Name: _ed.Name{Local: "\u006ev\u0047\u0072\u0070\u0053\u0070\u0050r"}}
	e.EncodeElement(_abd.NvGrpSpPr, _gga)
	_af := _ed.StartElement{Name: _ed.Name{Local: "\u0067r\u0070\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_abd.GrpSpPr, _af)
	if _abd.Choice != nil {
		for _, _ccg := range _abd.Choice {
			_ccg.MarshalXML(e, _ed.StartElement{})
		}
	}
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func NewCT_Connector() *CT_Connector {
	_ce := &CT_Connector{}
	_ce.NvCxnSpPr = NewCT_ConnectorNonVisual()
	_ce.SpPr = _c.NewCT_ShapeProperties()
	return _ce
}

// Validate validates the CT_Picture and its children
func (_bcd *CT_Picture) Validate() error {
	return _bcd.ValidateWithPath("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065")
}

// Validate validates the EG_ObjectChoicesChoice and its children
func (_abf *EG_ObjectChoicesChoice) Validate() error {
	return _abf.ValidateWithPath("\u0045\u0047\u005f\u004fbj\u0065\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0043\u0068\u006f\u0069c\u0065")
}

func (_aca *CT_GraphicFrame) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	if _aca.MacroAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _b.Sprintf("\u0025\u0076", *_aca.MacroAttr)})
	}
	if _aca.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _ed.Attr{Name: _ed.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _b.Sprintf("\u0025\u0064", _bgf(*_aca.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_cdg := _ed.StartElement{Name: _ed.Name{Local: "\u006e\u0076G\u0072\u0061\u0070h\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0050\u0072"}}
	e.EncodeElement(_aca.NvGraphicFramePr, _cdg)
	_geba := _ed.StartElement{Name: _ed.Name{Local: "\u0078\u0066\u0072\u006d"}}
	e.EncodeElement(_aca.Xfrm, _geba)
	_ff := _ed.StartElement{Name: _ed.Name{Local: "\u0061:\u0067\u0072\u0061\u0070\u0068\u0069c"}}
	_ff.Attr = append(_ff.Attr, _ed.Attr{Name: _ed.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	e.EncodeElement(_aca.Graphic, _ff)
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (_bdba *CT_PictureNonVisual) ValidateWithPath(path string) error {
	if _edb := _bdba.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _edb != nil {
		return _edb
	}
	if _cbd := _bdba.CNvPicPr.ValidateWithPath(path + "\u002fC\u004e\u0076\u0050\u0069\u0063\u0050r"); _cbd != nil {
		return _cbd
	}
	return nil
}

type CT_GroupShape struct {
	NvGrpSpPr *CT_GroupShapeNonVisual
	GrpSpPr   *_c.CT_GroupShapeProperties
	Choice    []*CT_GroupShapeChoice
}

// Validate validates the CT_Marker and its children
func (_bddb *CT_Marker) Validate() error {
	return _bddb.ValidateWithPath("\u0043T\u005f\u004d\u0061\u0072\u006b\u0065r")
}

func _bgf(_facf bool) uint8 {
	if _facf {
		return 1
	}
	return 0
}

// Validate validates the CT_GraphicFrame and its children
func (_fdd *CT_GraphicFrame) Validate() error {
	return _fdd.ValidateWithPath("\u0043T\u005fG\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065")
}

type CT_GraphicFrameNonVisual struct {
	CNvPr             *_c.CT_NonVisualDrawingProps
	CNvGraphicFramePr *_c.CT_NonVisualGraphicFrameProperties
}

func (_dd *CT_Connector) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_dd.NvCxnSpPr = NewCT_ConnectorNonVisual()
	_dd.SpPr = _c.NewCT_ShapeProperties()
	for _, _ef := range start.Attr {
		if _ef.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_ec, _db := _ef.Value, error(nil)
			if _db != nil {
				return _db
			}
			_dd.MacroAttr = &_ec
			continue
		}
		if _ef.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_da, _dfc := _a.ParseBool(_ef.Value)
			if _dfc != nil {
				return _dfc
			}
			_dd.FPublishedAttr = &_da
			continue
		}
	}
_bf:
	for {
		_aege, _ca := d.Token()
		if _ca != nil {
			return _ca
		}
		switch _cce := _aege.(type) {
		case _ed.StartElement:
			switch _cce.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0043\u0078\u006e\u0053\u0070\u0050r"}:
				if _dc := d.DecodeElement(_dd.NvCxnSpPr, &_cce); _dc != nil {
					return _dc
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _cb := d.DecodeElement(_dd.SpPr, &_cce); _cb != nil {
					return _cb
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_dd.Style = _c.NewCT_ShapeStyle()
				if _be := d.DecodeElement(_dd.Style, &_cce); _be != nil {
					return _be
				}
			default:
				_d.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_C\u006f\u006en\u0065\u0063\u0074\u006f\u0072\u0020\u0025\u0076", _cce.Name)
				if _dbc := d.Skip(); _dbc != nil {
					return _dbc
				}
			}
		case _ed.EndElement:
			break _bf
		case _ed.CharData:
		}
	}
	return nil
}

type CT_Marker struct {
	X float64
	Y float64
}

// ValidateWithPath validates the CT_ShapeNonVisual and its children, prefixing error messages with path
func (_egeaa *CT_ShapeNonVisual) ValidateWithPath(path string) error {
	if _gdfb := _egeaa.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _gdfb != nil {
		return _gdfb
	}
	if _feef := _egeaa.CNvSpPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0053\u0070\u0050\u0072"); _feef != nil {
		return _feef
	}
	return nil
}

// Validate validates the CT_AbsSizeAnchor and its children
func (_gc *CT_AbsSizeAnchor) Validate() error {
	return _gc.ValidateWithPath("\u0043\u0054_\u0041\u0062\u0073S\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072")
}

func (_bgcg *EG_ObjectChoicesChoice) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
_cged:
	for {
		_bbee, _gafb := d.Token()
		if _gafb != nil {
			return _gafb
		}
		switch _bac := _bbee.(type) {
		case _ed.StartElement:
			switch _bac.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_bgcg.Sp = NewCT_Shape()
				if _gag := d.DecodeElement(_bgcg.Sp, &_bac); _gag != nil {
					return _gag
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_bgcg.GrpSp = NewCT_GroupShape()
				if _bab := d.DecodeElement(_bgcg.GrpSp, &_bac); _bab != nil {
					return _bab
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_bgcg.GraphicFrame = NewCT_GraphicFrame()
				if _efcf := d.DecodeElement(_bgcg.GraphicFrame, &_bac); _efcf != nil {
					return _efcf
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_bgcg.CxnSp = NewCT_Connector()
				if _dfea := d.DecodeElement(_bgcg.CxnSp, &_bac); _dfea != nil {
					return _dfea
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_bgcg.Pic = NewCT_Picture()
				if _egef := d.DecodeElement(_bgcg.Pic, &_bac); _egef != nil {
					return _egef
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070p\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045G\u005f\u004f\u0062\u006a\u0065c\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _bac.Name)
				if _fca := d.Skip(); _fca != nil {
					return _fca
				}
			}
		case _ed.EndElement:
			break _cged
		case _ed.CharData:
		}
	}
	return nil
}

func (_ffc *CT_Marker) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	e.EncodeToken(start)
	_ceb := _ed.StartElement{Name: _ed.Name{Local: "\u0078"}}
	e.EncodeElement(_ffc.X, _ceb)
	_efca := _ed.StartElement{Name: _ed.Name{Local: "\u0079"}}
	e.EncodeElement(_ffc.Y, _efca)
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func (_gbgd *CT_PictureNonVisual) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	e.EncodeToken(start)
	_bda := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_gbgd.CNvPr, _bda)
	_dga := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_gbgd.CNvPicPr, _dga)
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func NewCT_GroupShapeNonVisual() *CT_GroupShapeNonVisual {
	_gbga := &CT_GroupShapeNonVisual{}
	_gbga.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_gbga.CNvGrpSpPr = _c.NewCT_NonVisualGroupDrawingShapeProps()
	return _gbga
}

type EG_Anchor struct {
	RelSizeAnchor *CT_RelSizeAnchor
	AbsSizeAnchor *CT_AbsSizeAnchor
}

func (_daa *EG_ObjectChoices) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	if _daa.Choice != nil {
		_daa.Choice.MarshalXML(e, _ed.StartElement{})
	}
	return nil
}

func NewCT_GraphicFrameNonVisual() *CT_GraphicFrameNonVisual {
	_ade := &CT_GraphicFrameNonVisual{}
	_ade.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_ade.CNvGraphicFramePr = _c.NewCT_NonVisualGraphicFrameProperties()
	return _ade
}

type CT_RelSizeAnchor struct {
	From   *CT_Marker
	To     *CT_Marker
	Choice *EG_ObjectChoicesChoice
}

func (_ecaa *EG_Anchor) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	if _ecaa.RelSizeAnchor != nil {
		_age := _ed.StartElement{Name: _ed.Name{Local: "\u0072\u0065\u006c\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_ecaa.RelSizeAnchor, _age)
	}
	if _ecaa.AbsSizeAnchor != nil {
		_fgeg := _ed.StartElement{Name: _ed.Name{Local: "\u0061\u0062\u0073\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_ecaa.AbsSizeAnchor, _fgeg)
	}
	return nil
}

// Validate validates the CT_ConnectorNonVisual and its children
func (_ad *CT_ConnectorNonVisual) Validate() error {
	return _ad.ValidateWithPath("C\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u004e\u006fn\u0056\u0069\u0073\u0075\u0061\u006c")
}

func (_fgb *CT_GroupShape) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_fgb.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	_fgb.GrpSpPr = _c.NewCT_GroupShapeProperties()
_ggaa:
	for {
		_gdb, _cfg := d.Token()
		if _cfg != nil {
			return _cfg
		}
		switch _cee := _gdb.(type) {
		case _ed.StartElement:
			switch _cee.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0047\u0072\u0070\u0053\u0070\u0050r"}:
				if _ega := d.DecodeElement(_fgb.NvGrpSpPr, &_cee); _ega != nil {
					return _ega
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067r\u0070\u0053\u0070\u0050\u0072"}:
				if _eace := d.DecodeElement(_fgb.GrpSpPr, &_cee); _eace != nil {
					return _eace
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_bfg := NewCT_GroupShapeChoice()
				if _cdc := d.DecodeElement(&_bfg.Sp, &_cee); _cdc != nil {
					return _cdc
				}
				_fgb.Choice = append(_fgb.Choice, _bfg)
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_aed := NewCT_GroupShapeChoice()
				if _gec := d.DecodeElement(&_aed.GrpSp, &_cee); _gec != nil {
					return _gec
				}
				_fgb.Choice = append(_fgb.Choice, _aed)
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_gda := NewCT_GroupShapeChoice()
				if _baf := d.DecodeElement(&_gda.GraphicFrame, &_cee); _baf != nil {
					return _baf
				}
				_fgb.Choice = append(_fgb.Choice, _gda)
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_cgeb := NewCT_GroupShapeChoice()
				if _egaa := d.DecodeElement(&_cgeb.CxnSp, &_cee); _egaa != nil {
					return _egaa
				}
				_fgb.Choice = append(_fgb.Choice, _cgeb)
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_eec := NewCT_GroupShapeChoice()
				if _dce := d.DecodeElement(&_eec.Pic, &_cee); _dce != nil {
					return _dce
				}
				_fgb.Choice = append(_fgb.Choice, _eec)
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047r\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065 \u0025\u0076", _cee.Name)
				if _dgb := d.Skip(); _dgb != nil {
					return _dgb
				}
			}
		case _ed.EndElement:
			break _ggaa
		case _ed.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Drawing and its children, prefixing error messages with path
func (_gf *CT_Drawing) ValidateWithPath(path string) error {
	for _geg, _ga := range _gf.EG_Anchor {
		if _gff := _ga.ValidateWithPath(_b.Sprintf("\u0025\u0073/\u0045\u0047\u005fA\u006e\u0063\u0068\u006f\u0072\u005b\u0025\u0064\u005d", path, _geg)); _gff != nil {
			return _gff
		}
	}
	return nil
}

func NewCT_Marker() *CT_Marker { _ebd := &CT_Marker{}; _ebd.X = 0.0; _ebd.Y = 0.0; return _ebd }

type CT_GroupShapeNonVisual struct {
	CNvPr      *_c.CT_NonVisualDrawingProps
	CNvGrpSpPr *_c.CT_NonVisualGroupDrawingShapeProps
}

// Validate validates the CT_GroupShape and its children
func (_bec *CT_GroupShape) Validate() error {
	return _bec.ValidateWithPath("\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065")
}

func (_fcfb *CT_ShapeNonVisual) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_fcfb.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_fcfb.CNvSpPr = _c.NewCT_NonVisualDrawingShapeProps()
_feaf:
	for {
		_eecg, _egec := d.Token()
		if _egec != nil {
			return _egec
		}
		switch _bfge := _eecg.(type) {
		case _ed.StartElement:
			switch _bfge.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _befa := d.DecodeElement(_fcfb.CNvPr, &_bfge); _befa != nil {
					return _befa
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063N\u0076\u0053\u0070\u0050\u0072"}:
				if _bbb := d.DecodeElement(_fcfb.CNvSpPr, &_bfge); _bbb != nil {
					return _bbb
				}
			default:
				_d.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _bfge.Name)
				if _bbba := d.Skip(); _bbba != nil {
					return _bbba
				}
			}
		case _ed.EndElement:
			break _feaf
		case _ed.CharData:
		}
	}
	return nil
}

type EG_ObjectChoices struct{ Choice *EG_ObjectChoicesChoice }

func (_eg *CT_ConnectorNonVisual) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	e.EncodeToken(start)
	_ge := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_eg.CNvPr, _ge)
	_fccg := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_eg.CNvCxnSpPr, _fccg)
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func (_ddfe *CT_Shape) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_ddfe.NvSpPr = NewCT_ShapeNonVisual()
	_ddfe.SpPr = _c.NewCT_ShapeProperties()
	for _, _gdf := range start.Attr {
		if _gdf.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_dcga, _ffed := _gdf.Value, error(nil)
			if _ffed != nil {
				return _ffed
			}
			_ddfe.MacroAttr = &_dcga
			continue
		}
		if _gdf.Name.Local == "\u0074\u0065\u0078\u0074\u006c\u0069\u006e\u006b" {
			_dag, _bebb := _gdf.Value, error(nil)
			if _bebb != nil {
				return _bebb
			}
			_ddfe.TextlinkAttr = &_dag
			continue
		}
		if _gdf.Name.Local == "\u0066\u004c\u006f\u0063\u006b\u0073\u0054\u0065\u0078\u0074" {
			_cea, _dadc := _a.ParseBool(_gdf.Value)
			if _dadc != nil {
				return _dadc
			}
			_ddfe.FLocksTextAttr = &_cea
			continue
		}
		if _gdf.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_egcc, _bce := _a.ParseBool(_gdf.Value)
			if _bce != nil {
				return _bce
			}
			_ddfe.FPublishedAttr = &_egcc
			continue
		}
	}
_efcd:
	for {
		_ffgd, _ede := d.Token()
		if _ede != nil {
			return _ede
		}
		switch _abe := _ffgd.(type) {
		case _ed.StartElement:
			switch _abe.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006e\u0076\u0053\u0070\u0050\u0072"}:
				if _bddg := d.DecodeElement(_ddfe.NvSpPr, &_abe); _bddg != nil {
					return _bddg
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _ebgb := d.DecodeElement(_ddfe.SpPr, &_abe); _ebgb != nil {
					return _ebgb
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_ddfe.Style = _c.NewCT_ShapeStyle()
				if _fae := d.DecodeElement(_ddfe.Style, &_abe); _fae != nil {
					return _fae
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0078\u0042\u006f\u0064\u0079"}:
				_ddfe.TxBody = _c.NewCT_TextBody()
				if _eccd := d.DecodeElement(_ddfe.TxBody, &_abe); _eccd != nil {
					return _eccd
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u0020\u0025\u0076", _abe.Name)
				if _eagd := d.Skip(); _eagd != nil {
					return _eagd
				}
			}
		case _ed.EndElement:
			break _efcd
		case _ed.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the EG_ObjectChoices and its children, prefixing error messages with path
func (_abc *EG_ObjectChoices) ValidateWithPath(path string) error {
	if _abc.Choice != nil {
		if _dgdd := _abc.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _dgdd != nil {
			return _dgdd
		}
	}
	return nil
}

func NewCT_GroupShape() *CT_GroupShape {
	_fgf := &CT_GroupShape{}
	_fgf.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	_fgf.GrpSpPr = _c.NewCT_GroupShapeProperties()
	return _fgf
}

func (_becee *CT_RelSizeAnchor) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_becee.From = NewCT_Marker()
	_becee.To = NewCT_Marker()
_aad:
	for {
		_eaea, _cfb := d.Token()
		if _cfb != nil {
			return _cfb
		}
		switch _ebed := _eaea.(type) {
		case _ed.StartElement:
			switch _ebed.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}:
				if _fefe := d.DecodeElement(_becee.From, &_ebed); _fefe != nil {
					return _fefe
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u006f"}:
				if _bfgg := d.DecodeElement(_becee.To, &_ebed); _bfgg != nil {
					return _bfgg
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_becee.Choice = NewEG_ObjectChoicesChoice()
				if _ecf := d.DecodeElement(&_becee.Choice.Sp, &_ebed); _ecf != nil {
					return _ecf
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_becee.Choice = NewEG_ObjectChoicesChoice()
				if _fbgf := d.DecodeElement(&_becee.Choice.GrpSp, &_ebed); _fbgf != nil {
					return _fbgf
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_becee.Choice = NewEG_ObjectChoicesChoice()
				if _bbe := d.DecodeElement(&_becee.Choice.GraphicFrame, &_ebed); _bbe != nil {
					return _bbe
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_becee.Choice = NewEG_ObjectChoicesChoice()
				if _ggb := d.DecodeElement(&_becee.Choice.CxnSp, &_ebed); _ggb != nil {
					return _ggb
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_becee.Choice = NewEG_ObjectChoicesChoice()
				if _bbea := d.DecodeElement(&_becee.Choice.Pic, &_ebed); _bbea != nil {
					return _bbea
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0065\u006c\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025v", _ebed.Name)
				if _gdd := d.Skip(); _gdd != nil {
					return _gdd
				}
			}
		case _ed.EndElement:
			break _aad
		case _ed.CharData:
		}
	}
	return nil
}

func NewCT_RelSizeAnchor() *CT_RelSizeAnchor {
	_dbcg := &CT_RelSizeAnchor{}
	_dbcg.From = NewCT_Marker()
	_dbcg.To = NewCT_Marker()
	return _dbcg
}

func NewEG_ObjectChoices() *EG_ObjectChoices { _eebe := &EG_ObjectChoices{}; return _eebe }

// Validate validates the CT_RelSizeAnchor and its children
func (_cbec *CT_RelSizeAnchor) Validate() error {
	return _cbec.ValidateWithPath("\u0043\u0054_\u0052\u0065\u006cS\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072")
}

// ValidateWithPath validates the CT_Shape and its children, prefixing error messages with path
func (_bafag *CT_Shape) ValidateWithPath(path string) error {
	if _dbef := _bafag.NvSpPr.ValidateWithPath(path + "\u002fN\u0076\u0053\u0070\u0050\u0072"); _dbef != nil {
		return _dbef
	}
	if _efd := _bafag.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _efd != nil {
		return _efd
	}
	if _bafag.Style != nil {
		if _abgd := _bafag.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _abgd != nil {
			return _abgd
		}
	}
	if _bafag.TxBody != nil {
		if _cag := _bafag.TxBody.ValidateWithPath(path + "\u002fT\u0078\u0042\u006f\u0064\u0079"); _cag != nil {
			return _cag
		}
	}
	return nil
}

// ValidateWithPath validates the CT_GraphicFrame and its children, prefixing error messages with path
func (_ccf *CT_GraphicFrame) ValidateWithPath(path string) error {
	if _eac := _ccf.NvGraphicFramePr.ValidateWithPath(path + "\u002f\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"); _eac != nil {
		return _eac
	}
	if _fbg := _ccf.Xfrm.ValidateWithPath(path + "\u002f\u0058\u0066r\u006d"); _fbg != nil {
		return _fbg
	}
	if _dff := _ccf.Graphic.ValidateWithPath(path + "\u002f\u0047\u0072\u0061\u0070\u0068\u0069\u0063"); _dff != nil {
		return _dff
	}
	return nil
}

type CT_Shape struct {
	MacroAttr      *string
	TextlinkAttr   *string
	FLocksTextAttr *bool
	FPublishedAttr *bool
	NvSpPr         *CT_ShapeNonVisual
	SpPr           *_c.CT_ShapeProperties
	Style          *_c.CT_ShapeStyle
	TxBody         *_c.CT_TextBody
}

func (_fcd *CT_GraphicFrameNonVisual) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_fcd.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_fcd.CNvGraphicFramePr = _c.NewCT_NonVisualGraphicFrameProperties()
_bdd:
	for {
		_cfe, _gbb := d.Token()
		if _gbb != nil {
			return _gbb
		}
		switch _cde := _cfe.(type) {
		case _ed.StartElement:
			switch _cde.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _agcd := d.DecodeElement(_fcd.CNvPr, &_cde); _agcd != nil {
					return _agcd
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"}:
				if _beb := d.DecodeElement(_fcd.CNvGraphicFramePr, &_cde); _beb != nil {
					return _beb
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c\u0020%\u0076", _cde.Name)
				if _dbag := d.Skip(); _dbag != nil {
					return _dbag
				}
			}
		case _ed.EndElement:
			break _bdd
		case _ed.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupShapeChoice and its children
func (_bgc *CT_GroupShapeChoice) Validate() error {
	return _bgc.ValidateWithPath("\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u0043\u0068\u006f\u0069\u0063\u0065")
}

func (_dcdd *EG_ObjectChoices) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
_bffc:
	for {
		_fccf, _eed := d.Token()
		if _eed != nil {
			return _eed
		}
		switch _bgd := _fccf.(type) {
		case _ed.StartElement:
			switch _bgd.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_dcdd.Choice = NewEG_ObjectChoicesChoice()
				if _cdf := d.DecodeElement(&_dcdd.Choice.Sp, &_bgd); _cdf != nil {
					return _cdf
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_dcdd.Choice = NewEG_ObjectChoicesChoice()
				if _gbe := d.DecodeElement(&_dcdd.Choice.GrpSp, &_bgd); _gbe != nil {
					return _gbe
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_dcdd.Choice = NewEG_ObjectChoicesChoice()
				if _cccg := d.DecodeElement(&_dcdd.Choice.GraphicFrame, &_bgd); _cccg != nil {
					return _cccg
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_dcdd.Choice = NewEG_ObjectChoicesChoice()
				if _dacb := d.DecodeElement(&_dcdd.Choice.CxnSp, &_bgd); _dacb != nil {
					return _dacb
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_dcdd.Choice = NewEG_ObjectChoicesChoice()
				if _efcaa := d.DecodeElement(&_dcdd.Choice.Pic, &_bgd); _efcaa != nil {
					return _efcaa
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u0047\u005f\u004f\u0062\u006a\u0065\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0020\u0025v", _bgd.Name)
				if _fff := d.Skip(); _fff != nil {
					return _fff
				}
			}
		case _ed.EndElement:
			break _bffc
		case _ed.CharData:
		}
	}
	return nil
}

func NewCT_Picture() *CT_Picture {
	_dfcc := &CT_Picture{}
	_dfcc.NvPicPr = NewCT_PictureNonVisual()
	_dfcc.BlipFill = _c.NewCT_BlipFillProperties()
	_dfcc.SpPr = _c.NewCT_ShapeProperties()
	return _dfcc
}

// ValidateWithPath validates the CT_GroupShapeChoice and its children, prefixing error messages with path
func (_baee *CT_GroupShapeChoice) ValidateWithPath(path string) error {
	for _dbca, _bfa := range _baee.Sp {
		if _ffd := _bfa.ValidateWithPath(_b.Sprintf("\u0025s\u002f\u0053\u0070\u005b\u0025\u0064]", path, _dbca)); _ffd != nil {
			return _ffd
		}
	}
	for _dgd, _cgc := range _baee.GrpSp {
		if _gcee := _cgc.ValidateWithPath(_b.Sprintf("\u0025\u0073\u002fG\u0072\u0070\u0053\u0070\u005b\u0025\u0064\u005d", path, _dgd)); _gcee != nil {
			return _gcee
		}
	}
	for _dcg, _dfe := range _baee.GraphicFrame {
		if _dgc := _dfe.ValidateWithPath(_b.Sprintf("\u0025\u0073\u002f\u0047ra\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u005b\u0025\u0064\u005d", path, _dcg)); _dgc != nil {
			return _dgc
		}
	}
	for _dgeb, _fef := range _baee.CxnSp {
		if _cgb := _fef.ValidateWithPath(_b.Sprintf("\u0025\u0073\u002fC\u0078\u006e\u0053\u0070\u005b\u0025\u0064\u005d", path, _dgeb)); _cgb != nil {
			return _cgb
		}
	}
	for _aag, _gbgcc := range _baee.Pic {
		if _agd := _gbgcc.ValidateWithPath(_b.Sprintf("\u0025\u0073\u002f\u0050\u0069\u0063\u005b\u0025\u0064\u005d", path, _aag)); _agd != nil {
			return _agd
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Connector and its children, prefixing error messages with path
func (_ag *CT_Connector) ValidateWithPath(path string) error {
	if _dcf := _ag.NvCxnSpPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"); _dcf != nil {
		return _dcf
	}
	if _ccc := _ag.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _ccc != nil {
		return _ccc
	}
	if _ag.Style != nil {
		if _fcc := _ag.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _fcc != nil {
			return _fcc
		}
	}
	return nil
}

type CT_ShapeNonVisual struct {
	CNvPr   *_c.CT_NonVisualDrawingProps
	CNvSpPr *_c.CT_NonVisualDrawingShapeProps
}

func NewCT_AbsSizeAnchor() *CT_AbsSizeAnchor {
	_f := &CT_AbsSizeAnchor{}
	_f.From = NewCT_Marker()
	_f.Ext = _c.NewCT_PositiveSize2D()
	return _f
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_efb *CT_Picture) ValidateWithPath(path string) error {
	if _eff := _efb.NvPicPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0050\u0069\u0063\u0050\u0072"); _eff != nil {
		return _eff
	}
	if _fbbcf := _efb.BlipFill.ValidateWithPath(path + "\u002fB\u006c\u0069\u0070\u0046\u0069\u006cl"); _fbbcf != nil {
		return _fbbcf
	}
	if _cbb := _efb.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _cbb != nil {
		return _cbb
	}
	if _efb.Style != nil {
		if _dfef := _efb.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _dfef != nil {
			return _dfef
		}
	}
	return nil
}

func NewCT_ConnectorNonVisual() *CT_ConnectorNonVisual {
	_cbe := &CT_ConnectorNonVisual{}
	_cbe.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_cbe.CNvCxnSpPr = _c.NewCT_NonVisualConnectorProperties()
	return _cbe
}

// Validate validates the CT_Drawing and its children
func (_agc *CT_Drawing) Validate() error {
	return _agc.ValidateWithPath("\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067")
}

func (_eeg *CT_ShapeNonVisual) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	e.EncodeToken(start)
	_cfgb := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_eeg.CNvPr, _cfgb)
	_fge := _ed.StartElement{Name: _ed.Name{Local: "\u0063N\u0076\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_eeg.CNvSpPr, _fge)
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_AbsSizeAnchor and its children, prefixing error messages with path
func (_ggc *CT_AbsSizeAnchor) ValidateWithPath(path string) error {
	if _gb := _ggc.From.ValidateWithPath(path + "\u002f\u0046\u0072o\u006d"); _gb != nil {
		return _gb
	}
	if _abb := _ggc.Ext.ValidateWithPath(path + "\u002f\u0045\u0078\u0074"); _abb != nil {
		return _abb
	}
	if _ggc.Choice != nil {
		if _fd := _ggc.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _fd != nil {
			return _fd
		}
	}
	return nil
}

type CT_GraphicFrame struct {
	MacroAttr        *string
	FPublishedAttr   *bool
	NvGraphicFramePr *CT_GraphicFrameNonVisual
	Xfrm             *_c.CT_Transform2D
	Graphic          *_c.Graphic
}

func NewCT_GroupShapeChoice() *CT_GroupShapeChoice { _ggaf := &CT_GroupShapeChoice{}; return _ggaf }

// Validate validates the EG_Anchor and its children
func (_ecga *EG_Anchor) Validate() error {
	return _ecga.ValidateWithPath("\u0045G\u005f\u0041\u006e\u0063\u0068\u006fr")
}

type CT_Connector struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvCxnSpPr      *CT_ConnectorNonVisual
	SpPr           *_c.CT_ShapeProperties
	Style          *_c.CT_ShapeStyle
}

type CT_Drawing struct{ EG_Anchor []*EG_Anchor }

// Validate validates the EG_ObjectChoices and its children
func (_gcc *EG_ObjectChoices) Validate() error {
	return _gcc.ValidateWithPath("\u0045\u0047_\u004f\u0062\u006ae\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073")
}

func (_cbbb *EG_ObjectChoicesChoice) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	if _cbbb.Sp != nil {
		_bcb := _ed.StartElement{Name: _ed.Name{Local: "\u0073\u0070"}}
		e.EncodeElement(_cbbb.Sp, _bcb)
	}
	if _cbbb.GrpSp != nil {
		_beea := _ed.StartElement{Name: _ed.Name{Local: "\u0067\u0072\u0070S\u0070"}}
		e.EncodeElement(_cbbb.GrpSp, _beea)
	}
	if _cbbb.GraphicFrame != nil {
		_cac := _ed.StartElement{Name: _ed.Name{Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}}
		e.EncodeElement(_cbbb.GraphicFrame, _cac)
	}
	if _cbbb.CxnSp != nil {
		_aebc := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u0078\u006eS\u0070"}}
		e.EncodeElement(_cbbb.CxnSp, _aebc)
	}
	if _cbbb.Pic != nil {
		_dab := _ed.StartElement{Name: _ed.Name{Local: "\u0070\u0069\u0063"}}
		e.EncodeElement(_cbbb.Pic, _dab)
	}
	return nil
}

func (_gfd *EG_Anchor) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
_aacc:
	for {
		_adcc, _bed := d.Token()
		if _bed != nil {
			return _bed
		}
		switch _efec := _adcc.(type) {
		case _ed.StartElement:
			switch _efec.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u0065\u006c\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_gfd.RelSizeAnchor = NewCT_RelSizeAnchor()
				if _fbfe := d.DecodeElement(_gfd.RelSizeAnchor, &_efec); _fbfe != nil {
					return _fbfe
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_gfd.AbsSizeAnchor = NewCT_AbsSizeAnchor()
				if _effa := d.DecodeElement(_gfd.AbsSizeAnchor, &_efec); _effa != nil {
					return _effa
				}
			default:
				_d.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0045\u0047\u005f\u0041\u006e\u0063h\u006f\u0072 \u0025\u0076", _efec.Name)
				if _ebdc := d.Skip(); _ebdc != nil {
					return _ebdc
				}
			}
		case _ed.EndElement:
			break _aacc
		case _ed.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_RelSizeAnchor and its children, prefixing error messages with path
func (_ddee *CT_RelSizeAnchor) ValidateWithPath(path string) error {
	if _bdg := _ddee.From.ValidateWithPath(path + "\u002f\u0046\u0072o\u006d"); _bdg != nil {
		return _bdg
	}
	if _bdff := _ddee.To.ValidateWithPath(path + "\u002f\u0054\u006f"); _bdff != nil {
		return _bdff
	}
	if _ddee.Choice != nil {
		if _aee := _ddee.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _aee != nil {
			return _aee
		}
	}
	return nil
}

type CT_PictureNonVisual struct {
	CNvPr    *_c.CT_NonVisualDrawingProps
	CNvPicPr *_c.CT_NonVisualPictureProperties
}

func (_ccb *CT_Picture) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_ccb.NvPicPr = NewCT_PictureNonVisual()
	_ccb.BlipFill = _c.NewCT_BlipFillProperties()
	_ccb.SpPr = _c.NewCT_ShapeProperties()
	for _, _ebdb := range start.Attr {
		if _ebdb.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_aagg, _bfac := _ebdb.Value, error(nil)
			if _bfac != nil {
				return _bfac
			}
			_ccb.MacroAttr = &_aagg
			continue
		}
		if _ebdb.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_bbdc, _bdcg := _a.ParseBool(_ebdb.Value)
			if _bdcg != nil {
				return _bdcg
			}
			_ccb.FPublishedAttr = &_bbdc
			continue
		}
	}
_afag:
	for {
		_bad, _afc := d.Token()
		if _afc != nil {
			return _afc
		}
		switch _cbeg := _bad.(type) {
		case _ed.StartElement:
			switch _cbeg.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _ggae := d.DecodeElement(_ccb.NvPicPr, &_cbeg); _ggae != nil {
					return _ggae
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _fece := d.DecodeElement(_ccb.BlipFill, &_cbeg); _fece != nil {
					return _fece
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _dbac := d.DecodeElement(_ccb.SpPr, &_cbeg); _dbac != nil {
					return _dbac
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_ccb.Style = _c.NewCT_ShapeStyle()
				if _bff := d.DecodeElement(_ccb.Style, &_cbeg); _bff != nil {
					return _bff
				}
			default:
				_d.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0069\u0063\u0074\u0075\u0072\u0065\u0020\u0025\u0076", _cbeg.Name)
				if _abbb := d.Skip(); _abbb != nil {
					return _abbb
				}
			}
		case _ed.EndElement:
			break _afag
		case _ed.CharData:
		}
	}
	return nil
}

type CT_ConnectorNonVisual struct {
	CNvPr      *_c.CT_NonVisualDrawingProps
	CNvCxnSpPr *_c.CT_NonVisualConnectorProperties
}

type CT_Picture struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvPicPr        *CT_PictureNonVisual
	BlipFill       *_c.CT_BlipFillProperties
	SpPr           *_c.CT_ShapeProperties
	Style          *_c.CT_ShapeStyle
}

func (_dfd *CT_Drawing) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
_fbd:
	for {
		_ebe, _bc := d.Token()
		if _bc != nil {
			return _bc
		}
		switch _de := _ebe.(type) {
		case _ed.StartElement:
			switch _de.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u0065\u006c\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_cf := NewEG_Anchor()
				_cf.RelSizeAnchor = NewCT_RelSizeAnchor()
				if _egg := d.DecodeElement(_cf.RelSizeAnchor, &_de); _egg != nil {
					return _egg
				}
				_dfd.EG_Anchor = append(_dfd.EG_Anchor, _cf)
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_gce := NewEG_Anchor()
				_gce.AbsSizeAnchor = NewCT_AbsSizeAnchor()
				if _dba := d.DecodeElement(_gce.AbsSizeAnchor, &_de); _dba != nil {
					return _dba
				}
				_dfd.EG_Anchor = append(_dfd.EG_Anchor, _gce)
			default:
				_d.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fD\u0072\u0061\u0077\u0069\u006e\u0067\u0020\u0025\u0076", _de.Name)
				if _dfg := d.Skip(); _dfg != nil {
					return _dfg
				}
			}
		case _ed.EndElement:
			break _fbd
		case _ed.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PictureNonVisual and its children
func (_eag *CT_PictureNonVisual) Validate() error {
	return _eag.ValidateWithPath("\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}

func (_dbd *CT_RelSizeAnchor) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	e.EncodeToken(start)
	_deg := _ed.StartElement{Name: _ed.Name{Local: "\u0066\u0072\u006f\u006d"}}
	e.EncodeElement(_dbd.From, _deg)
	_cba := _ed.StartElement{Name: _ed.Name{Local: "\u0074\u006f"}}
	e.EncodeElement(_dbd.To, _cba)
	if _dbd.Choice != nil {
		_dbd.Choice.MarshalXML(e, _ed.StartElement{})
	}
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func NewEG_Anchor() *EG_Anchor { _face := &EG_Anchor{}; return _face }

// ValidateWithPath validates the EG_Anchor and its children, prefixing error messages with path
func (_cdge *EG_Anchor) ValidateWithPath(path string) error {
	if _cdge.RelSizeAnchor != nil {
		if _begd := _cdge.RelSizeAnchor.ValidateWithPath(path + "\u002f\u0052\u0065\u006c\u0053\u0069\u007a\u0065\u0041n\u0063\u0068\u006f\u0072"); _begd != nil {
			return _begd
		}
	}
	if _cdge.AbsSizeAnchor != nil {
		if _agg := _cdge.AbsSizeAnchor.ValidateWithPath(path + "\u002f\u0041\u0062\u0073\u0053\u0069\u007a\u0065\u0041n\u0063\u0068\u006f\u0072"); _agg != nil {
			return _agg
		}
	}
	return nil
}

func (_abg *CT_GroupShapeNonVisual) MarshalXML(e *_ed.Encoder, start _ed.StartElement) error {
	e.EncodeToken(start)
	_bga := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u004e\u0076P\u0072"}}
	e.EncodeElement(_abg.CNvPr, _bga)
	_ddf := _ed.StartElement{Name: _ed.Name{Local: "\u0063\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_abg.CNvGrpSpPr, _ddf)
	e.EncodeToken(_ed.EndElement{Name: start.Name})
	return nil
}

func (_ae *CT_AbsSizeAnchor) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_ae.From = NewCT_Marker()
	_ae.Ext = _c.NewCT_PositiveSize2D()
_df:
	for {
		_ab, _ac := d.Token()
		if _ac != nil {
			return _ac
		}
		switch _g := _ab.(type) {
		case _ed.StartElement:
			switch _g.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}:
				if _fa := d.DecodeElement(_ae.From, &_g); _fa != nil {
					return _fa
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}:
				if _fg := d.DecodeElement(_ae.Ext, &_g); _fg != nil {
					return _fg
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_ae.Choice = NewEG_ObjectChoicesChoice()
				if _bb := d.DecodeElement(&_ae.Choice.Sp, &_g); _bb != nil {
					return _bb
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_ae.Choice = NewEG_ObjectChoicesChoice()
				if _fc := d.DecodeElement(&_ae.Choice.GrpSp, &_g); _fc != nil {
					return _fc
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_ae.Choice = NewEG_ObjectChoicesChoice()
				if _cd := d.DecodeElement(&_ae.Choice.GraphicFrame, &_g); _cd != nil {
					return _cd
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_ae.Choice = NewEG_ObjectChoicesChoice()
				if _gg := d.DecodeElement(&_ae.Choice.CxnSp, &_g); _gg != nil {
					return _gg
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_ae.Choice = NewEG_ObjectChoicesChoice()
				if _aeg := d.DecodeElement(&_ae.Choice.Pic, &_g); _aeg != nil {
					return _aeg
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0041\u0062\u0073\u0053\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025v", _g.Name)
				if _fab := d.Skip(); _fab != nil {
					return _fab
				}
			}
		case _ed.EndElement:
			break _df
		case _ed.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GraphicFrameNonVisual and its children
func (_bdee *CT_GraphicFrameNonVisual) Validate() error {
	return _bdee.ValidateWithPath("\u0043T\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061m\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}

// ValidateWithPath validates the CT_Marker and its children, prefixing error messages with path
func (_fde *CT_Marker) ValidateWithPath(path string) error {
	if _fde.X < 0.0 {
		return _b.Errorf("\u0025\u0073\u002fm\u002e\u0058\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0030\u002e\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _fde.X)
	}
	if _fde.X > 1.0 {
		return _b.Errorf("\u0025\u0073\u002fm\u002e\u0058\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003c=\u0020\u0031\u002e\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _fde.X)
	}
	if _fde.Y < 0.0 {
		return _b.Errorf("\u0025\u0073\u002fm\u002e\u0059\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0030\u002e\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _fde.Y)
	}
	if _fde.Y > 1.0 {
		return _b.Errorf("\u0025\u0073\u002fm\u002e\u0059\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003c=\u0020\u0031\u002e\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _fde.Y)
	}
	return nil
}

type CT_GroupShapeChoice struct {
	Sp           []*CT_Shape
	GrpSp        []*CT_GroupShape
	GraphicFrame []*CT_GraphicFrame
	CxnSp        []*CT_Connector
	Pic          []*CT_Picture
}

// ValidateWithPath validates the CT_GraphicFrameNonVisual and its children, prefixing error messages with path
func (_ffg *CT_GraphicFrameNonVisual) ValidateWithPath(path string) error {
	if _bbd := _ffg.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _bbd != nil {
		return _bbd
	}
	if _ege := _ffg.CNvGraphicFramePr.ValidateWithPath(path + "\u002fC\u004ev\u0047\u0072\u0061\u0070\u0068i\u0063\u0046r\u0061\u006d\u0065\u0050\u0072"); _ege != nil {
		return _ege
	}
	return nil
}

func (_ebf *CT_PictureNonVisual) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
	_ebf.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_ebf.CNvPicPr = _c.NewCT_NonVisualPictureProperties()
_bdb:
	for {
		_eaa, _cga := d.Token()
		if _cga != nil {
			return _cga
		}
		switch _deae := _eaa.(type) {
		case _ed.StartElement:
			switch _deae.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _ffe := d.DecodeElement(_ebf.CNvPr, &_deae); _ffe != nil {
					return _ffe
				}
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}:
				if _baec := d.DecodeElement(_ebf.CNvPicPr, &_deae); _baec != nil {
					return _baec
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065No\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _deae.Name)
				if _fdc := d.Skip(); _fdc != nil {
					return _fdc
				}
			}
		case _ed.EndElement:
			break _bdb
		case _ed.CharData:
		}
	}
	return nil
}

func (_efa *CT_GroupShapeChoice) UnmarshalXML(d *_ed.Decoder, start _ed.StartElement) error {
_aac:
	for {
		_cfed, _ccgb := d.Token()
		if _ccgb != nil {
			return _ccgb
		}
		switch _gdaf := _cfed.(type) {
		case _ed.StartElement:
			switch _gdaf.Name {
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_dee := NewCT_Shape()
				if _bdf := d.DecodeElement(_dee, &_gdaf); _bdf != nil {
					return _bdf
				}
				_efa.Sp = append(_efa.Sp, _dee)
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_fac := NewCT_GroupShape()
				if _cfgf := d.DecodeElement(_fac, &_gdaf); _cfgf != nil {
					return _cfgf
				}
				_efa.GrpSp = append(_efa.GrpSp, _fac)
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_afa := NewCT_GraphicFrame()
				if _bfd := d.DecodeElement(_afa, &_gdaf); _bfd != nil {
					return _bfd
				}
				_efa.GraphicFrame = append(_efa.GraphicFrame, _afa)
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_beg := NewCT_Connector()
				if _ecg := d.DecodeElement(_beg, &_gdaf); _ecg != nil {
					return _ecg
				}
				_efa.CxnSp = append(_efa.CxnSp, _beg)
			case _ed.Name{Space: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_dde := NewCT_Picture()
				if _dcd := d.DecodeElement(_dde, &_gdaf); _dcd != nil {
					return _dcd
				}
				_efa.Pic = append(_efa.Pic, _dde)
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068ap\u0065\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _gdaf.Name)
				if _adf := d.Skip(); _adf != nil {
					return _adf
				}
			}
		case _ed.EndElement:
			break _aac
		case _ed.CharData:
		}
	}
	return nil
}

func init() {
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c", NewCT_ShapeNonVisual)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065", NewCT_Shape)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "C\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u004e\u006fn\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_ConnectorNonVisual)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005fC\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072", NewCT_Connector)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_PictureNonVisual)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065", NewCT_Picture)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043T\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061m\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_GraphicFrameNonVisual)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043T\u005fG\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065", NewCT_GraphicFrame)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075a\u006c", NewCT_GroupShapeNonVisual)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065", NewCT_GroupShape)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043T\u005f\u004d\u0061\u0072\u006b\u0065r", NewCT_Marker)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054_\u0052\u0065\u006cS\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072", NewCT_RelSizeAnchor)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054_\u0041\u0062\u0073S\u0069\u007a\u0065\u0041\u006e\u0063\u0068\u006f\u0072", NewCT_AbsSizeAnchor)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067", NewCT_Drawing)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0045\u0047_\u004f\u0062\u006ae\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073", NewEG_ObjectChoices)
	_bd.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", "\u0045G\u005f\u0041\u006e\u0063\u0068\u006fr", NewEG_Anchor)
}
