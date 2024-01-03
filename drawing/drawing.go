package drawing

import (
	_c "github.com/arcpd/msword"
	_d "github.com/arcpd/msword/color"
	_a "github.com/arcpd/msword/measurement"
	_e "github.com/arcpd/msword/schema/soo/dml"
)

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_g LineProperties) SetWidth(w _a.Distance) { _g._f.WAttr = _c.Int32(int32(w / _a.EMU)) }

func (_bfg ShapeProperties) SetNoFill() {
	_bfg.clearFill()
	_bfg._ea.NoFill = _e.NewCT_NoFillProperties()
}

const (
	LineJoinRound LineJoin = iota
	LineJoinBevel
	LineJoinMiter
)

// SetBulletFont controls the font for the bullet character.
func (_eb ParagraphProperties) SetBulletFont(f string) {
	if f == "" {
		_eb._ab.BuFont = nil
	} else {
		_eb._ab.BuFont = _e.NewCT_TextFont()
		_eb._ab.BuFont.TypefaceAttr = f
	}
}

// SetText sets the run's text contents.
func (_dc Run) SetText(s string) {
	_dc._agfa.Br = nil
	_dc._agfa.Fld = nil
	if _dc._agfa.R == nil {
		_dc._agfa.R = _e.NewCT_RegularTextRun()
	}
	_dc._agfa.R.T = s
}

// X returns the inner wrapped XML type.
func (_ggd Run) X() *_e.EG_TextRun { return _ggd._agfa }

// LineJoin is the type of line join
type LineJoin byte

// MakeRun constructs a new Run wrapper.
func MakeRun(x *_e.EG_TextRun) Run { return Run{x} }

func (_ec ShapeProperties) ensureXfrm() {
	if _ec._ea.Xfrm == nil {
		_ec._ea.Xfrm = _e.NewCT_Transform2D()
	}
}

// AddBreak adds a new line break to a paragraph.
func (_gg Paragraph) AddBreak() {
	_fb := _e.NewEG_TextRun()
	_fb.Br = _e.NewCT_TextLineBreak()
	_gg._eg.EG_TextRun = append(_gg._eg.EG_TextRun, _fb)
}

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties(x *_e.CT_TextCharacterProperties) RunProperties { return RunProperties{x} }

// AddRun adds a new run to a paragraph.
func (_da Paragraph) AddRun() Run {
	_ae := MakeRun(_e.NewEG_TextRun())
	_da._eg.EG_TextRun = append(_da._eg.EG_TextRun, _ae.X())
	return _ae
}

// SetHeight sets the height of the shape.
func (_ac ShapeProperties) SetHeight(h _a.Distance) {
	_ac.ensureXfrm()
	if _ac._ea.Xfrm.Ext == nil {
		_ac._ea.Xfrm.Ext = _e.NewCT_PositiveSize2D()
	}
	_ac._ea.Xfrm.Ext.CyAttr = int64(h / _a.EMU)
}

// SetFlipVertical controls if the shape is flipped vertically.
func (_gee ShapeProperties) SetFlipVertical(b bool) {
	_gee.ensureXfrm()
	if !b {
		_gee._ea.Xfrm.FlipVAttr = nil
	} else {
		_gee._ea.Xfrm.FlipVAttr = _c.Bool(true)
	}
}

// SetFont controls the font of a run.
func (_ebc RunProperties) SetFont(s string) {
	_ebc._def.Latin = _e.NewCT_TextFont()
	_ebc._def.Latin.TypefaceAttr = s
}

// SetGeometry sets the shape type of the shape
func (_cf ShapeProperties) SetGeometry(g _e.ST_ShapeType) {
	if _cf._ea.PrstGeom == nil {
		_cf._ea.PrstGeom = _e.NewCT_PresetGeometry2D()
	}
	_cf._ea.PrstGeom.PrstAttr = g
}

// GetPosition gets the position of the shape in EMU.
func (_bd ShapeProperties) GetPosition() (int64, int64) {
	_bd.ensureXfrm()
	if _bd._ea.Xfrm.Off == nil {
		_bd._ea.Xfrm.Off = _e.NewCT_Point2D()
	}
	return *_bd._ea.Xfrm.Off.XAttr.ST_CoordinateUnqualified, *_bd._ea.Xfrm.Off.YAttr.ST_CoordinateUnqualified
}

// X returns the inner wrapped XML type.
func (_ff Paragraph) X() *_e.CT_TextParagraph { return _ff._eg }

// SetNumbered controls if bullets are numbered or not.
func (_bcg ParagraphProperties) SetNumbered(scheme _e.ST_TextAutonumberScheme) {
	if scheme == _e.ST_TextAutonumberSchemeUnset {
		_bcg._ab.BuAutoNum = nil
	} else {
		_bcg._ab.BuAutoNum = _e.NewCT_TextAutonumberBullet()
		_bcg._ab.BuAutoNum.TypeAttr = scheme
	}
}

func (_ga LineProperties) clearFill() {
	_ga._f.NoFill = nil
	_ga._f.GradFill = nil
	_ga._f.SolidFill = nil
	_ga._f.PattFill = nil
}

func (_ge LineProperties) SetNoFill() {
	_ge.clearFill()
	_ge._f.NoFill = _e.NewCT_NoFillProperties()
}

// X returns the inner wrapped XML type.
func (_fg ShapeProperties) X() *_e.CT_ShapeProperties { return _fg._ea }

// SetSize sets the font size of the run text
func (_bf RunProperties) SetSize(sz _a.Distance) {
	_bf._def.SzAttr = _c.Int32(int32(sz / _a.HundredthPoint))
}

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph(x *_e.CT_TextParagraph) Paragraph { return Paragraph{x} }

// Properties returns the run's properties.
func (_de Run) Properties() RunProperties {
	if _de._agfa.R == nil {
		_de._agfa.R = _e.NewCT_RegularTextRun()
	}
	if _de._agfa.R.RPr == nil {
		_de._agfa.R.RPr = _e.NewCT_TextCharacterProperties()
	}
	return RunProperties{_de._agfa.R.RPr}
}

// RunProperties controls the run properties.
type RunProperties struct {
	_def *_e.CT_TextCharacterProperties
}

// Properties returns the paragraph properties.
func (_ed Paragraph) Properties() ParagraphProperties {
	if _ed._eg.PPr == nil {
		_ed._eg.PPr = _e.NewCT_TextParagraphProperties()
	}
	return MakeParagraphProperties(_ed._eg.PPr)
}

// SetSolidFill controls the text color of a run.
func (_fba RunProperties) SetSolidFill(c _d.Color) {
	_fba._def.NoFill = nil
	_fba._def.BlipFill = nil
	_fba._def.GradFill = nil
	_fba._def.GrpFill = nil
	_fba._def.PattFill = nil
	_fba._def.SolidFill = _e.NewCT_SolidColorFillProperties()
	_fba._def.SolidFill.SrgbClr = _e.NewCT_SRgbColor()
	_fba._def.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

func MakeShapeProperties(x *_e.CT_ShapeProperties) ShapeProperties { return ShapeProperties{x} }

// SetJoin sets the line join style.
func (_aa LineProperties) SetJoin(e LineJoin) {
	_aa._f.Round = nil
	_aa._f.Miter = nil
	_aa._f.Bevel = nil
	switch e {
	case LineJoinRound:
		_aa._f.Round = _e.NewCT_LineJoinRound()
	case LineJoinBevel:
		_aa._f.Bevel = _e.NewCT_LineJoinBevel()
	case LineJoinMiter:
		_aa._f.Miter = _e.NewCT_LineJoinMiterProperties()
	}
}

type ShapeProperties struct{ _ea *_e.CT_ShapeProperties }

type LineProperties struct{ _f *_e.CT_LineProperties }

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct {
	_ab *_e.CT_TextParagraphProperties
}

func (_gge ShapeProperties) LineProperties() LineProperties {
	if _gge._ea.Ln == nil {
		_gge._ea.Ln = _e.NewCT_LineProperties()
	}
	return LineProperties{_gge._ea.Ln}
}

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_cb ShapeProperties) SetFlipHorizontal(b bool) {
	_cb.ensureXfrm()
	if !b {
		_cb._ea.Xfrm.FlipHAttr = nil
	} else {
		_cb._ea.Xfrm.FlipHAttr = _c.Bool(true)
	}
}

// SetBold controls the bolding of a run.
func (_gb RunProperties) SetBold(b bool) { _gb._def.BAttr = _c.Bool(b) }

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties(x *_e.CT_TextParagraphProperties) ParagraphProperties {
	return ParagraphProperties{x}
}

func (_dab ShapeProperties) SetSolidFill(c _d.Color) {
	_dab.clearFill()
	_dab._ea.SolidFill = _e.NewCT_SolidColorFillProperties()
	_dab._ea.SolidFill.SrgbClr = _e.NewCT_SRgbColor()
	_dab._ea.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// SetWidth sets the width of the shape.
func (_bfc ShapeProperties) SetWidth(w _a.Distance) {
	_bfc.ensureXfrm()
	if _bfc._ea.Xfrm.Ext == nil {
		_bfc._ea.Xfrm.Ext = _e.NewCT_PositiveSize2D()
	}
	_bfc._ea.Xfrm.Ext.CxAttr = int64(w / _a.EMU)
}

// SetBulletChar sets the bullet character for the paragraph.
func (_be ParagraphProperties) SetBulletChar(c string) {
	if c == "" {
		_be._ab.BuChar = nil
	} else {
		_be._ab.BuChar = _e.NewCT_TextCharBullet()
		_be._ab.BuChar.CharAttr = c
	}
}

func (_gf ShapeProperties) clearFill() {
	_gf._ea.NoFill = nil
	_gf._ea.BlipFill = nil
	_gf._ea.GradFill = nil
	_gf._ea.GrpFill = nil
	_gf._ea.SolidFill = nil
	_gf._ea.PattFill = nil
}

// SetSize sets the width and height of the shape.
func (_ef ShapeProperties) SetSize(w, h _a.Distance) { _ef.SetWidth(w); _ef.SetHeight(h) }

func (_ag LineProperties) SetSolidFill(c _d.Color) {
	_ag.clearFill()
	_ag._f.SolidFill = _e.NewCT_SolidColorFillProperties()
	_ag._f.SolidFill.SrgbClr = _e.NewCT_SRgbColor()
	_ag._f.SolidFill.SrgbClr.ValAttr = *c.AsRGBString()
}

// Run is a run within a paragraph.
type Run struct{ _agfa *_e.EG_TextRun }

// X returns the inner wrapped XML type.
func (_bc ParagraphProperties) X() *_e.CT_TextParagraphProperties { return _bc._ab }

// Paragraph is a paragraph within a document.
type Paragraph struct{ _eg *_e.CT_TextParagraph }

// SetPosition sets the position of the shape.
func (_ecc ShapeProperties) SetPosition(x, y _a.Distance) {
	_ecc.ensureXfrm()
	if _ecc._ea.Xfrm.Off == nil {
		_ecc._ea.Xfrm.Off = _e.NewCT_Point2D()
	}
	_ecc._ea.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _c.Int64(int64(x / _a.EMU))
	_ecc._ea.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _c.Int64(int64(y / _a.EMU))
}

// SetLevel sets the level of indentation of a paragraph.
func (_gga ParagraphProperties) SetLevel(idx int32) { _gga._ab.LvlAttr = _c.Int32(idx) }

// SetAlign controls the paragraph alignment
func (_agf ParagraphProperties) SetAlign(a _e.ST_TextAlignType) { _agf._ab.AlgnAttr = a }

// X returns the inner wrapped XML type.
func (_ca LineProperties) X() *_e.CT_LineProperties { return _ca._f }
