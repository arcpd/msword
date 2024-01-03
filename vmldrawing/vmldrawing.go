package vmldrawing

import (
	_d "encoding/xml"
	_cb "fmt"
	_eg "github.com/arcpd/msword"
	_cbb "github.com/arcpd/msword/common/logger"
	_de "github.com/arcpd/msword/schema/soo/ofc/sharedTypes"
	_a "github.com/arcpd/msword/schema/urn/schemas_microsoft_com/office/excel"
	_ae "github.com/arcpd/msword/schema/urn/schemas_microsoft_com/vml"
	_e "strconv"
	_f "strings"
)

func NewContainer() *Container { return &Container{} }
func (_b *Container) MarshalXML(e *_d.Encoder, start _d.StartElement) error {
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0076"}, Value: "\u0075\u0072n\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d:v\u006d\u006c"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u006f"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006di\u0063\u0072\u006f\u0073\u006f\u0066t\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u006ff\u0066\u0069\u0063\u0065"})
	start.Attr = append(start.Attr, _d.Attr{Name: _d.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078"}, Value: "\u0075\u0072\u006e\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002d\u006d\u0069\u0063\u0072\u006f\u0073\u006f\u0066\u0074\u002d\u0063\u006fm\u003a\u006f\u0066\u0066\u0069c\u0065\u003ae\u0078\u0063\u0065\u006c"})
	start.Name.Local = "\u0078\u006d\u006c"
	e.EncodeToken(start)
	if _b.Layout != nil {
		_dea := _d.StartElement{Name: _d.Name{Local: "\u006f\u003a\u0073\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074"}}
		e.EncodeElement(_b.Layout, _dea)
	}
	if _b.ShapeType != nil {
		_ba := _d.StartElement{Name: _d.Name{Local: "v\u003a\u0073\u0068\u0061\u0070\u0065\u0074\u0079\u0070\u0065"}}
		e.EncodeElement(_b.ShapeType, _ba)
	}
	for _, _ff := range _b.Shape {
		_fff := _d.StartElement{Name: _d.Name{Local: "\u0076:\u0073\u0068\u0061\u0070\u0065"}}
		e.EncodeElement(_ff, _fff)
	}
	return e.EncodeToken(_d.EndElement{Name: start.Name})
}

// ToString generate string of TextpathStyle.
func (_cea *TextpathStyle) String() string {
	_deaf := ""
	_deaf += _cb.Sprintf("\u0066o\u006et\u002d\u0066\u0061\u006d\u0069\u006c\u0079\u003a\u0025\u0073\u003b", _cea._eag)
	_deaf += _cb.Sprintf("\u0066o\u006et\u002d\u0073\u0069\u007a\u0065\u003a\u0025\u0064\u0070\u0074\u003b", _cea._fbb)
	if _cea._cga {
		_deaf += _cb.Sprintf("\u0066o\u006et\u002d\u0073\u0074\u0079\u006ce\u003a\u0069t\u0061\u006c\u0069\u0063\u003b")
	}
	if _cea._da {
		_deaf += _cb.Sprintf("\u0066\u006f\u006e\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074\u003ab\u006f\u006c\u0064\u003b")
	}
	return _deaf
}

// TextpathStyle is style attribute of element v:textpath.
type TextpathStyle struct {
	_eag string
	_fbb int64
	_da  bool
	_cga bool
}

// Right get right attribute of shape style.
func (_bae *ShapeStyle) Right() float64 { return _bae._aee }

// Width return width of shape.
func (_ade *ShapeStyle) Width() float64 { return _ade._gcg }

// SetFontSize sets text's fontSize.
func (_gdf *TextpathStyle) SetFontSize(fontSize int64) { _gdf._fbb = fontSize }

// Margins get margin top, left, bottom, and right of shape style.
func (_bb *ShapeStyle) Margins() (float64, float64, float64, float64) {
	return _bb._ad, _bb._fg, _bb._ea, _bb._ab
}

// CreateFormula creates F element for typeFormulas.
func CreateFormula(s string) *_ae.CT_F { _cde := _ae.NewCT_F(); _cde.EqnAttr = &s; return _cde }

// Left get left attribute of shape style.
func (_fe *ShapeStyle) Left() float64 { return _fe._gg }

// SetHeight set height of shape.
func (_dgd *ShapeStyle) SetHeight(height float64) { _dgd._deac = height }

// SetBold sets text to bold.
func (_dedg *TextpathStyle) SetBold(bold bool) { _dedg._da = bold }

// IsItalic returns true if text is italic.
func (_be *TextpathStyle) IsItalic() bool { return _be._cga }

// Height return height of shape.
func (_caa *ShapeStyle) Height() float64 { return _caa._deac }

type Container struct {
	Layout    *_ae.OfcShapelayout
	ShapeType *_ae.Shapetype
	Shape     []*_ae.Shape
}

// IsBold returns true if text is bold.
func (_cbf *TextpathStyle) IsBold() bool { return _cbf._da }

// ShapeStyle is style attribute of v:shape element.
type ShapeStyle struct {
	_af   string
	_ad   float64
	_fg   float64
	_ea   float64
	_ab   float64
	_aac  float64
	_gg   float64
	_ca   float64
	_aee  float64
	_gcg  float64
	_deac float64
	_cf   int64
	_bac  string
	_ee   string
	_fb   string
	_ef   string
}

// NewCommentDrawing constructs a new comment drawing.
func NewCommentDrawing() *Container {
	_g := NewContainer()
	_g.Layout = _ae.NewOfcShapelayout()
	_g.Layout.ExtAttr = _ae.ST_ExtEdit
	_g.Layout.Idmap = _ae.NewOfcCT_IdMap()
	_g.Layout.Idmap.DataAttr = _eg.String("\u0031")
	_g.Layout.Idmap.ExtAttr = _ae.ST_ExtEdit
	_g.ShapeType = _ae.NewShapetype()
	_g.ShapeType.IdAttr = _eg.String("_\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032")
	_g.ShapeType.CoordsizeAttr = _eg.String("2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030")
	_g.ShapeType.SptAttr = _eg.Float32(202)
	_g.ShapeType.PathAttr = _eg.String("\u006d\u0030\u002c0l\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u00321\u00360\u0030,\u00321\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030\u002c\u0030\u0078\u0065")
	_gf := _ae.NewEG_ShapeElements()
	_g.ShapeType.EG_ShapeElements = append(_g.ShapeType.EG_ShapeElements, _gf)
	_gf.Path = _ae.NewPath()
	_gf.Path.GradientshapeokAttr = _de.ST_TrueFalseT
	_gf.Path.ConnecttypeAttr = _ae.OfcST_ConnectTypeRect
	return _g
}

// NewCommentShape creates a new comment shape for a given cell index.  The
// indices here are zero based.
func NewCommentShape(col, row int64) *_ae.Shape {
	_gc := _ae.NewShape()
	_gc.IdAttr = _eg.String(_cb.Sprintf("\u0063\u0073\u005f\u0025\u0064\u005f\u0025\u0064", col, row))
	_gc.TypeAttr = _eg.String("\u0023\u005f\u00780\u0030\u0030\u0030\u005f\u0074\u0032\u0030\u0032")
	_gc.StyleAttr = _eg.String("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0061\u0062\u0073\u006f\u006cu\u0074\u0065\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074:\u0038\u0030\u0070\u0074;\u006d\u0061\u0072\u0067\u0069n-\u0074o\u0070\u003a\u0032pt\u003b\u0077\u0069\u0064\u0074\u0068\u003a1\u0030\u0034\u0070\u0074\u003b\u0068\u0065\u0069\u0067\u0068\u0074\u003a\u0037\u0036\u0070\u0074\u003b\u007a\u002d\u0069\u006e\u0064\u0065x\u003a\u0031\u003bv\u0069\u0073\u0069\u0062\u0069\u006c\u0069t\u0079\u003a\u0068\u0069\u0064\u0064\u0065\u006e")
	_gc.FillcolorAttr = _eg.String("\u0023f\u0062\u0066\u0036\u0064\u0036")
	_gc.StrokecolorAttr = _eg.String("\u0023e\u0064\u0065\u0061\u0061\u0031")
	_dc := _ae.NewEG_ShapeElements()
	_dc.Fill = _ae.NewFill()
	_dc.Fill.Color2Attr = _eg.String("\u0023f\u0062\u0066\u0065\u0038\u0032")
	_dc.Fill.AngleAttr = _eg.Float64(-180)
	_dc.Fill.TypeAttr = _ae.ST_FillTypeGradient
	_dc.Fill.Fill = _ae.NewOfcFill()
	_dc.Fill.Fill.ExtAttr = _ae.ST_ExtView
	_dc.Fill.Fill.TypeAttr = _ae.OfcST_FillTypeGradientUnscaled
	_gc.EG_ShapeElements = append(_gc.EG_ShapeElements, _dc)
	_df := _ae.NewEG_ShapeElements()
	_df.Shadow = _ae.NewShadow()
	_df.Shadow.OnAttr = _de.ST_TrueFalseT
	_df.Shadow.ObscuredAttr = _de.ST_TrueFalseT
	_gc.EG_ShapeElements = append(_gc.EG_ShapeElements, _df)
	_aa := _ae.NewEG_ShapeElements()
	_aa.Path = _ae.NewPath()
	_aa.Path.ConnecttypeAttr = _ae.OfcST_ConnectTypeNone
	_gc.EG_ShapeElements = append(_gc.EG_ShapeElements, _aa)
	_ec := _ae.NewEG_ShapeElements()
	_ec.Textbox = _ae.NewTextbox()
	_ec.Textbox.StyleAttr = _eg.String("\u006d\u0073\u006f\u002ddi\u0072\u0065\u0063\u0074\u0069\u006f\u006e\u002d\u0061\u006c\u0074\u003a\u0061\u0075t\u006f")
	_gc.EG_ShapeElements = append(_gc.EG_ShapeElements, _ec)
	_cg := _ae.NewEG_ShapeElements()
	_cg.ClientData = _a.NewClientData()
	_cg.ClientData.ObjectTypeAttr = _a.ST_ObjectTypeNote
	_cg.ClientData.MoveWithCells = _de.ST_TrueFalseBlankT
	_cg.ClientData.SizeWithCells = _de.ST_TrueFalseBlankT
	_cg.ClientData.Anchor = _eg.String("\u0031,\u0020\u0031\u0035\u002c\u0020\u0030\u002c\u0020\u0032\u002c\u00202\u002c\u0020\u0035\u0034\u002c\u0020\u0035\u002c\u0020\u0033")
	_cg.ClientData.AutoFill = _de.ST_TrueFalseBlankFalse
	_cg.ClientData.Row = _eg.Int64(row)
	_cg.ClientData.Column = _eg.Int64(col)
	_gc.EG_ShapeElements = append(_gc.EG_ShapeElements, _cg)
	return _gc
}

// MSOPositionVerticalRelative get `mso-position-vertical-relative` attribute of shape style.
func (_dcd *ShapeStyle) MSOPositionVerticalRelative() string { return _dcd._ef }

// FontSize returns fontSize of the text.
func (_edd *TextpathStyle) FontSize() int64 { return _edd._fbb }

// NewTextpathStyle accept value of string style attribute of element v:textpath and format it to generate TextpathStyle.
func NewTextpathStyle(style string) TextpathStyle {
	_egc := TextpathStyle{_eag: "\u0022C\u0061\u006c\u0069\u0062\u0072\u0069\"", _fbb: 44, _da: false, _cga: false}
	_adc := _f.Split(style, "\u003b")
	for _, _cef := range _adc {
		_ac := _f.Split(_cef, "\u003a")
		if len(_ac) != 2 {
			continue
		}
		switch _ac[0] {
		case "f\u006f\u006e\u0074\u002d\u0066\u0061\u006d\u0069\u006c\u0079":
			_egc._eag = _ac[1]
			break
		case "\u0066o\u006e\u0074\u002d\u0073\u0069\u007ae":
			_egc._fbb, _ = _e.ParseInt(_f.ReplaceAll(_ac[1], "\u0070\u0074", ""), 10, 64)
			break
		case "f\u006f\u006e\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074":
			_egc._da = _ac[1] == "\u0062\u006f\u006c\u0064"
			break
		case "\u0066\u006f\u006e\u0074\u002d\u0073\u0074\u0079\u006c\u0065":
			_egc._cga = _ac[1] == "\u0069\u0074\u0061\u006c\u0069\u0063"
			break
		}
	}
	return _egc
}
func (_ed *Container) UnmarshalXML(d *_d.Decoder, start _d.StartElement) error {
	_ed.Shape = nil
_egb:
	for {
		_dg, _bd := d.Token()
		if _bd != nil {
			return _bd
		}
		switch _eb := _dg.(type) {
		case _d.StartElement:
			switch _eb.Name.Local {
			case "s\u0068\u0061\u0070\u0065\u006c\u0061\u0079\u006f\u0075\u0074":
				_ed.Layout = _ae.NewOfcShapelayout()
				if _bf := d.DecodeElement(_ed.Layout, &_eb); _bf != nil {
					return _bf
				}
			case "\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e":
				_ed.ShapeType = _ae.NewShapetype()
				if _cc := d.DecodeElement(_ed.ShapeType, &_eb); _cc != nil {
					return _cc
				}
			case "\u0073\u0068\u0061p\u0065":
				_dga := _ae.NewShape()
				if _ffd := d.DecodeElement(_dga, &_eb); _ffd != nil {
					return _ffd
				}
				_ed.Shape = append(_ed.Shape, _dga)
			}
		case _d.EndElement:
			break _egb
		}
	}
	return nil
}

// SetItalic sets text to italic.
func (_gge *TextpathStyle) SetItalic(italic bool) { _gge._cga = italic }

// MSOPositionHorizontalRelative get `mso-position-horizontal-relative` attribute of shape style.
func (_gcf *ShapeStyle) MSOPositionHorizontalRelative() string { return _gcf._ee }

const (
	ShapeStylePositionAbsolute = "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065"
	ShapeStylePositionRelative = "\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065"
)

// Bottom get bottom attribute of shape style.
func (_ebb *ShapeStyle) Bottom() float64 { return _ebb._ca }

// ToString formatting ShapeStyle to string.
func (_ggba *ShapeStyle) String() string {
	_cd := ""
	_cd += _cb.Sprintf("\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u003a\u0025\u0073\u003b", _ggba._af)
	_cd += _cb.Sprintf("\u006da\u0072g\u0069\u006e\u002d\u006c\u0065\u0066\u0074\u003a\u0025\u0064\u003b", int64(_ggba._fg))
	_cd += _cb.Sprintf("\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006fp\u003a\u0025\u0064\u003b", int64(_ggba._ad))
	_cd += _cb.Sprintf("w\u0069\u0064\u0074\u0068\u003a\u0025\u0064\u0070\u0074\u003b", int64(_ggba._gcg))
	_cd += _cb.Sprintf("\u0068\u0065\u0069g\u0068\u0074\u003a\u0025\u0064\u0070\u0074\u003b", int64(_ggba._deac))
	_cd += _cb.Sprintf("z\u002d\u0069\u006e\u0064\u0065\u0078\u003a\u0025\u0064\u003b", _ggba._cf)
	_cd += _cb.Sprintf("m\u0073\u006f\u002d\u0070\u006f\u0073i\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069z\u006f\u006e\u0074a\u006c:\u0025\u0073\u003b", _ggba._bac)
	_cd += _cb.Sprintf("\u006d\u0073o-\u0070\u006f\u0073i\u0074\u0069\u006f\u006e-ho\u0072iz\u006f\u006e\u0074\u0061\u006c\u002d\u0072el\u0061\u0074\u0069\u0076\u0065\u003a\u0025s\u003b", _ggba._ee)
	_cd += _cb.Sprintf("\u006ds\u006f\u002d\u0070\u006fs\u0069\u0074\u0069\u006f\u006e-\u0076e\u0072t\u0069\u0063\u0061\u006c\u003a\u0025\u0073;", _ggba._fb)
	_cd += _cb.Sprintf("\u006d\u0073\u006f-p\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0076e\u0072t\u0069c\u0061l\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065\u003a\u0025\u0073\u003b", _ggba._ef)
	return _cd
}

// SetFontFamily sets text's fontFamily.
func (_dgg *TextpathStyle) SetFontFamily(fontFamily string) { _dgg._eag = fontFamily }

// NewShapeStyle accept value of string style attribute in v:shape and format it to generate ShapeStyle.
func NewShapeStyle(style string) ShapeStyle {
	_dfa := ShapeStyle{_gcg: 0, _deac: 0}
	_ggb := _f.Split(style, "\u003b")
	for _, _gb := range _ggb {
		_bc := _f.Split(_gb, "\u003a")
		if len(_bc) != 2 {
			continue
		}
		var _gfc error
		switch _bc[0] {
		case "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e":
			_dfa._af = _bc[1]
			break
		case "\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006f\u0070":
			_dfa._ad, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "m\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074":
			_dfa._fg, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0062\u006f\u0074\u0074\u006f\u006d":
			_dfa._ea, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "\u006d\u0061\u0072g\u0069\u006e\u002d\u0072\u0069\u0067\u0068\u0074":
			_dfa._ab, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "\u0074\u006f\u0070":
			_dfa._aac, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "\u006c\u0065\u0066\u0074":
			_dfa._gg, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "\u0062\u006f\u0074\u0074\u006f\u006d":
			_dfa._ca, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "\u0072\u0069\u0067h\u0074":
			_dfa._aee, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "\u0077\u0069\u0064t\u0068":
			_dfa._gcg, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "\u0068\u0065\u0069\u0067\u0068\u0074":
			_dfa._deac, _gfc = _e.ParseFloat(_f.ReplaceAll(_bc[1], "\u0070\u0074", ""), 64)
			break
		case "\u007a-\u0069\u006e\u0064\u0065\u0078":
			_dfa._cf, _gfc = _e.ParseInt(_bc[1], 10, 64)
			break
		case "\u006d\u0073\u006f-p\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c":
			_dfa._bac = _bc[1]
			break
		case "\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006fr\u0069z\u006f\u006e\u0074\u0061\u006c\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065":
			_dfa._ee = _bc[1]
			break
		case "m\u0073\u006f\u002d\u0070os\u0069t\u0069\u006f\u006e\u002d\u0076e\u0072\u0074\u0069\u0063\u0061\u006c":
			_dfa._fb = _bc[1]
			break
		case "\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069o\u006e\u002d\u0076\u0065\u0072\u0074\u0069c\u0061\u006c\u002d\u0072\u0065\u006c\u0061\u0074\u0069\u0076\u0065":
			_dfa._ef = _bc[1]
			break
		}
		if _gfc != nil {
			_cbb.Log.Debug("\u0055n\u0061\u0062l\u0065\u0020\u0074o\u0020\u0070\u0061\u0072\u0073\u0065\u0020s\u0074\u0079\u006c\u0065\u0020\u0061t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0073 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076", _bc[0], _bc[1])
		}
	}
	return _dfa
}

// Top get top attribute of shape style.
func (_fd *ShapeStyle) Top() float64 { return _fd._aac }

// FontFamily returns fontFamily of the text.
func (_gd *TextpathStyle) FontFamily() string { return _gd._eag }

// Position get position attribute of shape style.
func (_baa *ShapeStyle) Position() string { return _baa._af }

// SetWidth set width of shape.
func (_ce *ShapeStyle) SetWidth(width float64) { _ce._gcg = width }
