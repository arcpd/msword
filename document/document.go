package document

import (
	_a "archive/zip"
	_fed "bytes"
	_bg "errors"
	_g "fmt"
	_b "github.com/arcpd/msword"
	_bd "github.com/arcpd/msword/color"
	_ffa "github.com/arcpd/msword/common"
	_ae "github.com/arcpd/msword/common/axcontrol"
	_c "github.com/arcpd/msword/common/logger"
	_be "github.com/arcpd/msword/common/tempstorage"
	_ad "github.com/arcpd/msword/internal/formatutils"
	_bc "github.com/arcpd/msword/measurement"
	_gd "github.com/arcpd/msword/schema/schemas.microsoft.com/office/activeX"
	_aa "github.com/arcpd/msword/schema/soo/dml"
	_ee "github.com/arcpd/msword/schema/soo/dml/chart"
	_ba "github.com/arcpd/msword/schema/soo/dml/picture"
	_ca "github.com/arcpd/msword/schema/soo/ofc/sharedTypes"
	_gf "github.com/arcpd/msword/schema/soo/pkg/relationships"
	_db "github.com/arcpd/msword/schema/soo/wml"
	_eb "github.com/arcpd/msword/schema/urn/schemas_microsoft_com/vml"
	_ge "github.com/arcpd/msword/vmldrawing"
	_ed "github.com/arcpd/msword/zippkg"
	_dg "image"
	_fa "image/jpeg"
	_fg "io"
	_f "math/rand"
	_fe "os"
	_fc "path/filepath"
	_ff "regexp"
	_eg "strings"
	_e "unicode"
)

// DrawingInfo is used for keep information about a drawing wrapping a textbox where the text is located.
type DrawingInfo struct {
	Drawing *_db.CT_Drawing
	Width   int64
	Height  int64
}

// NewAnchorDrawWrapOptions return anchor drawing options property.
func NewAnchorDrawWrapOptions() *AnchorDrawWrapOptions {
	_feb := &AnchorDrawWrapOptions{}
	if !_feb._dbg {
		_fb, _dgb := _ffce()
		_feb._ega = _fb
		_feb._bb = _dgb
	}
	return _feb
}

// SetFirstLineIndent controls the first line indent of the paragraph.
func (_fbbae ParagraphStyleProperties) SetFirstLineIndent(m _bc.Distance) {
	if _fbbae._fagf.Ind == nil {
		_fbbae._fagf.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_fbbae._fagf.Ind.FirstLineAttr = nil
	} else {
		_fbbae._fagf.Ind.FirstLineAttr = &_ca.ST_TwipsMeasure{}
		_fbbae._fagf.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(m / _bc.Twips))
	}
}

// Margins allows controlling individual cell margins.
func (_aed CellProperties) Margins() CellMargins {
	if _aed._beea.TcMar == nil {
		_aed._beea.TcMar = _db.NewCT_TcMar()
	}
	return CellMargins{_aed._beea.TcMar}
}

// RunProperties controls run styling properties.
type RunProperties struct{ _cadb *_db.CT_RPr }

// Settings controls the document settings.
type Settings struct{ _ccgc *_db.Settings }

// TableProperties returns the table style properties.
func (_fbca Style) TableProperties() TableStyleProperties {
	if _fbca._eedcd.TblPr == nil {
		_fbca._eedcd.TblPr = _db.NewCT_TblPrBase()
	}
	return TableStyleProperties{_fbca._eedcd.TblPr}
}

// SetProperties sets the `RunProperties` to the `Run` element
func (_gccf Run) SetProperties(rp RunProperties) {
	if rp._cadb == nil {
		_gccf._afec.RPr = _db.NewCT_RPr()
	} else {
		_adbg := &_db.CT_RPr{}
		if rp._cadb.RStyle != nil {
			_adbg.RStyle = &_db.CT_String{ValAttr: rp._cadb.RStyle.ValAttr}
		}
		if rp._cadb.RFonts != nil {
			_adbg.RFonts = &_db.CT_Fonts{HintAttr: rp._cadb.RFonts.HintAttr, AsciiAttr: rp._cadb.RFonts.AsciiAttr, HAnsiAttr: rp._cadb.RFonts.HAnsiAttr, EastAsiaAttr: rp._cadb.RFonts.EastAsiaAttr, CsAttr: rp._cadb.RFonts.CsAttr, AsciiThemeAttr: rp._cadb.RFonts.AsciiThemeAttr, HAnsiThemeAttr: rp._cadb.RFonts.HAnsiThemeAttr, EastAsiaThemeAttr: rp._cadb.RFonts.EastAsiaThemeAttr, CsthemeAttr: rp._cadb.RFonts.CsthemeAttr}
		}
		if rp._cadb.B != nil {
			_adbg.B = &_db.CT_OnOff{}
			if rp._cadb.B.ValAttr != nil {
				_fgcae := *rp._cadb.B.ValAttr
				_adbg.B.ValAttr = &_fgcae
			}
		}
		if rp._cadb.BCs != nil {
			_adbg.BCs = &_db.CT_OnOff{}
			if rp._cadb.BCs.ValAttr != nil {
				_gged := *rp._cadb.BCs.ValAttr
				_adbg.BCs.ValAttr = &_gged
			}
		}
		if rp._cadb.I != nil {
			_adbg.I = &_db.CT_OnOff{}
			if rp._cadb.I.ValAttr != nil {
				_ffgd := *rp._cadb.I.ValAttr
				_adbg.I.ValAttr = &_ffgd
			}
		}
		if rp._cadb.ICs != nil {
			_adbg.ICs = &_db.CT_OnOff{}
			if rp._cadb.ICs.ValAttr != nil {
				_cdffe := *rp._cadb.ICs.ValAttr
				_adbg.ICs.ValAttr = &_cdffe
			}
		}
		if rp._cadb.Caps != nil {
			_adbg.Caps = &_db.CT_OnOff{}
			if rp._cadb.Caps.ValAttr != nil {
				_aabfa := *rp._cadb.Caps.ValAttr
				_adbg.Caps.ValAttr = &_aabfa
			}
		}
		if rp._cadb.SmallCaps != nil {
			_adbg.SmallCaps = &_db.CT_OnOff{}
			if rp._cadb.SmallCaps.ValAttr != nil {
				_cecdb := *rp._cadb.SmallCaps.ValAttr
				_adbg.SmallCaps.ValAttr = &_cecdb
			}
		}
		if rp._cadb.Strike != nil {
			_adbg.Strike = &_db.CT_OnOff{}
			if rp._cadb.Strike.ValAttr != nil {
				_cfeca := *rp._cadb.Strike.ValAttr
				_adbg.Strike.ValAttr = &_cfeca
			}
		}
		if rp._cadb.Dstrike != nil {
			_adbg.Dstrike = &_db.CT_OnOff{}
			if rp._cadb.Dstrike.ValAttr != nil {
				_cdabb := *rp._cadb.Dstrike.ValAttr
				_adbg.Dstrike.ValAttr = &_cdabb
			}
		}
		if rp._cadb.Outline != nil {
			_adbg.Outline = &_db.CT_OnOff{}
			if rp._cadb.Outline.ValAttr != nil {
				_efce := *rp._cadb.Outline.ValAttr
				_adbg.Outline.ValAttr = &_efce
			}
		}
		if rp._cadb.Shadow != nil {
			_adbg.Shadow = &_db.CT_OnOff{}
			if rp._cadb.Shadow.ValAttr != nil {
				_cfcc := *rp._cadb.Shadow.ValAttr
				_adbg.Shadow.ValAttr = &_cfcc
			}
		}
		if rp._cadb.Emboss != nil {
			_adbg.Emboss = &_db.CT_OnOff{}
			if rp._cadb.Emboss.ValAttr != nil {
				_bdbab := *rp._cadb.Emboss.ValAttr
				_adbg.Emboss.ValAttr = &_bdbab
			}
		}
		if rp._cadb.Imprint != nil {
			_adbg.Imprint = &_db.CT_OnOff{}
			if rp._cadb.Imprint.ValAttr != nil {
				_befc := *rp._cadb.Imprint.ValAttr
				_adbg.Imprint.ValAttr = &_befc
			}
		}
		if rp._cadb.NoProof != nil {
			_adbg.NoProof = &_db.CT_OnOff{}
			if rp._cadb.NoProof.ValAttr != nil {
				_cgad := *rp._cadb.NoProof.ValAttr
				_adbg.NoProof.ValAttr = &_cgad
			}
		}
		if rp._cadb.SnapToGrid != nil {
			_adbg.SnapToGrid = &_db.CT_OnOff{}
			if rp._cadb.SnapToGrid.ValAttr != nil {
				_fbdfe := *rp._cadb.SnapToGrid.ValAttr
				_adbg.SnapToGrid.ValAttr = &_fbdfe
			}
		}
		if rp._cadb.Vanish != nil {
			_adbg.Vanish = &_db.CT_OnOff{}
			if rp._cadb.Vanish.ValAttr != nil {
				_gecge := *rp._cadb.Vanish.ValAttr
				_adbg.Vanish.ValAttr = &_gecge
			}
		}
		if rp._cadb.WebHidden != nil {
			_adbg.WebHidden = &_db.CT_OnOff{}
			if rp._cadb.WebHidden.ValAttr != nil {
				_edge := *rp._cadb.WebHidden.ValAttr
				_adbg.WebHidden.ValAttr = &_edge
			}
		}
		if rp._cadb.Color != nil {
			_adbg.Color = &_db.CT_Color{ValAttr: rp._cadb.Color.ValAttr, ThemeColorAttr: rp._cadb.Color.ThemeColorAttr, ThemeTintAttr: rp._cadb.Color.ThemeTintAttr, ThemeShadeAttr: rp._cadb.Color.ThemeShadeAttr}
		}
		if rp._cadb.Spacing != nil {
			_adbg.Spacing = &_db.CT_SignedTwipsMeasure{ValAttr: rp._cadb.Spacing.ValAttr}
		}
		if rp._cadb.W != nil {
			_adbg.W = &_db.CT_TextScale{ValAttr: rp._cadb.W.ValAttr}
		}
		if rp._cadb.Kern != nil {
			_adbg.Kern = &_db.CT_HpsMeasure{ValAttr: rp._cadb.Kern.ValAttr}
		}
		if rp._cadb.Position != nil {
			_adbg.Position = &_db.CT_SignedHpsMeasure{ValAttr: rp._cadb.Position.ValAttr}
		}
		if rp._cadb.Sz != nil {
			_adbg.Sz = &_db.CT_HpsMeasure{ValAttr: rp._cadb.Sz.ValAttr}
		}
		if rp._cadb.SzCs != nil {
			_adbg.SzCs = &_db.CT_HpsMeasure{ValAttr: rp._cadb.SzCs.ValAttr}
		}
		if rp._cadb.Highlight != nil {
			_adbg.Highlight = &_db.CT_Highlight{ValAttr: rp._cadb.Highlight.ValAttr}
		}
		if rp._cadb.U != nil {
			_adbg.U = &_db.CT_Underline{ValAttr: rp._cadb.U.ValAttr, ColorAttr: rp._cadb.U.ColorAttr, ThemeColorAttr: rp._cadb.U.ThemeColorAttr, ThemeTintAttr: rp._cadb.U.ThemeTintAttr, ThemeShadeAttr: rp._cadb.U.ThemeShadeAttr}
		}
		if rp._cadb.Effect != nil {
			_adbg.Effect = &_db.CT_TextEffect{ValAttr: rp._cadb.Effect.ValAttr}
		}
		if rp._cadb.Bdr != nil {
			_adbg.Bdr = &_db.CT_Border{ValAttr: rp._cadb.Bdr.ValAttr, ColorAttr: rp._cadb.Bdr.ColorAttr, ThemeColorAttr: rp._cadb.Bdr.ThemeColorAttr, ThemeTintAttr: rp._cadb.Bdr.ThemeTintAttr, ThemeShadeAttr: rp._cadb.Bdr.ThemeShadeAttr, SzAttr: rp._cadb.Bdr.SzAttr, SpaceAttr: rp._cadb.Bdr.SpaceAttr, ShadowAttr: rp._cadb.Bdr.ShadowAttr, FrameAttr: rp._cadb.Bdr.FrameAttr}
		}
		if rp._cadb.Shd != nil {
			_adbg.Shd = &_db.CT_Shd{ValAttr: rp._cadb.Shd.ValAttr, ColorAttr: rp._cadb.Shd.ColorAttr, ThemeColorAttr: rp._cadb.Shd.ThemeColorAttr, ThemeTintAttr: rp._cadb.Shd.ThemeTintAttr, ThemeShadeAttr: rp._cadb.Shd.ThemeShadeAttr, FillAttr: rp._cadb.Shd.FillAttr, ThemeFillAttr: rp._cadb.Shd.ThemeFillAttr, ThemeFillTintAttr: rp._cadb.Shd.ThemeFillTintAttr, ThemeFillShadeAttr: rp._cadb.Shd.ThemeFillShadeAttr}
		}
		if rp._cadb.FitText != nil {
			_adbg.FitText = &_db.CT_FitText{ValAttr: rp._cadb.FitText.ValAttr, IdAttr: rp._cadb.FitText.IdAttr}
		}
		if rp._cadb.VertAlign != nil {
			_adbg.VertAlign = &_db.CT_VerticalAlignRun{ValAttr: rp._cadb.VertAlign.ValAttr}
		}
		if rp._cadb.Rtl != nil {
			_adbg.Rtl = &_db.CT_OnOff{ValAttr: rp._cadb.Rtl.ValAttr}
		}
		if rp._cadb.Cs != nil {
			_adbg.Cs = &_db.CT_OnOff{ValAttr: rp._cadb.Cs.ValAttr}
		}
		if rp._cadb.Em != nil {
			_adbg.Em = &_db.CT_Em{ValAttr: rp._cadb.Em.ValAttr}
		}
		if rp._cadb.Lang != nil {
			_adbg.Lang = &_db.CT_Language{ValAttr: rp._cadb.Lang.ValAttr, EastAsiaAttr: rp._cadb.Lang.EastAsiaAttr, BidiAttr: rp._cadb.Lang.BidiAttr}
		}
		if rp._cadb.EastAsianLayout != nil {
			_adbg.EastAsianLayout = &_db.CT_EastAsianLayout{IdAttr: rp._cadb.EastAsianLayout.IdAttr, CombineAttr: rp._cadb.EastAsianLayout.CombineAttr, CombineBracketsAttr: rp._cadb.EastAsianLayout.CombineBracketsAttr, VertAttr: rp._cadb.EastAsianLayout.VertAttr, VertCompressAttr: rp._cadb.EastAsianLayout.VertCompressAttr}
		}
		if rp._cadb.SpecVanish != nil {
			_adbg.SpecVanish = &_db.CT_OnOff{ValAttr: rp._cadb.SpecVanish.ValAttr}
		}
		if rp._cadb.OMath != nil {
			_adbg.OMath = &_db.CT_OnOff{ValAttr: rp._cadb.OMath.ValAttr}
		}
		_gccf._afec.RPr = _adbg
	}
}
func (_gebe *Document) appendTable(_eafb *Paragraph, _fdab Table, _bebb bool) Table {
	_gge := _gebe._adcb.Body
	_dafc := _db.NewEG_BlockLevelElts()
	_gebe._adcb.Body.EG_BlockLevelElts = append(_gebe._adcb.Body.EG_BlockLevelElts, _dafc)
	_bcf := _db.NewEG_ContentBlockContent()
	_dafc.EG_ContentBlockContent = append(_dafc.EG_ContentBlockContent, _bcf)
	if _eafb != nil {
		_acf := _eafb.X()
		for _eedd, _bed := range _gge.EG_BlockLevelElts {
			for _, _afac := range _bed.EG_ContentBlockContent {
				for _bead, _ddca := range _bcf.P {
					if _ddca == _acf {
						_fgege := _fdab.X()
						_ebed := _db.NewEG_BlockLevelElts()
						_bce := _db.NewEG_ContentBlockContent()
						_ebed.EG_ContentBlockContent = append(_ebed.EG_ContentBlockContent, _bce)
						_bce.Tbl = append(_bce.Tbl, _fgege)
						_gge.EG_BlockLevelElts = append(_gge.EG_BlockLevelElts, nil)
						if _bebb {
							copy(_gge.EG_BlockLevelElts[_eedd+1:], _gge.EG_BlockLevelElts[_eedd:])
							_gge.EG_BlockLevelElts[_eedd] = _ebed
							if _bead != 0 {
								_adca := _db.NewEG_BlockLevelElts()
								_cdc := _db.NewEG_ContentBlockContent()
								_adca.EG_ContentBlockContent = append(_adca.EG_ContentBlockContent, _cdc)
								_cdc.P = _afac.P[:_bead]
								_gge.EG_BlockLevelElts = append(_gge.EG_BlockLevelElts, nil)
								copy(_gge.EG_BlockLevelElts[_eedd+1:], _gge.EG_BlockLevelElts[_eedd:])
								_gge.EG_BlockLevelElts[_eedd] = _adca
							}
							_afac.P = _afac.P[_bead:]
						} else {
							copy(_gge.EG_BlockLevelElts[_eedd+2:], _gge.EG_BlockLevelElts[_eedd+1:])
							_gge.EG_BlockLevelElts[_eedd+1] = _ebed
							if _bead != len(_afac.P)-1 {
								_ggbe := _db.NewEG_BlockLevelElts()
								_fbb := _db.NewEG_ContentBlockContent()
								_ggbe.EG_ContentBlockContent = append(_ggbe.EG_ContentBlockContent, _fbb)
								_fbb.P = _afac.P[_bead+1:]
								_gge.EG_BlockLevelElts = append(_gge.EG_BlockLevelElts, nil)
								copy(_gge.EG_BlockLevelElts[_eedd+3:], _gge.EG_BlockLevelElts[_eedd+2:])
								_gge.EG_BlockLevelElts[_eedd+2] = _ggbe
							}
							_afac.P = _afac.P[:_bead+1]
						}
						break
					}
				}
				for _, _gaaf := range _afac.Tbl {
					_cce := _gaed(_gaaf, _acf, _bebb)
					if _cce != nil {
						break
					}
				}
			}
		}
	} else {
		_bcf.Tbl = append(_bcf.Tbl, _fdab.X())
	}
	return Table{_gebe, _fdab.X()}
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (_ebeeb ParagraphProperties) SetKeepOnOnePage(b bool) {
	if !b {
		_ebeeb._aedc.KeepLines = nil
	} else {
		_ebeeb._aedc.KeepLines = _db.NewCT_OnOff()
	}
}

// SetEastAsiaTheme sets the font East Asia Theme.
func (_cegf Fonts) SetEastAsiaTheme(t _db.ST_Theme) { _cegf._bdge.EastAsiaThemeAttr = t }

// VerticalAlign returns the value of paragraph vertical align.
func (_dddbf ParagraphProperties) VerticalAlignment() _ca.ST_VerticalAlignRun {
	if _bcgea := _dddbf._aedc.RPr.VertAlign; _bcgea != nil {
		return _bcgea.ValAttr
	}
	return 0
}
func (_gfdc Paragraph) addBeginFldChar(_egbb string) *_db.CT_FFData {
	_ddebd := _gfdc.addFldChar()
	_ddebd.FldCharTypeAttr = _db.ST_FldCharTypeBegin
	_ddebd.FfData = _db.NewCT_FFData()
	_aadae := _db.NewCT_FFName()
	_aadae.ValAttr = &_egbb
	_ddebd.FfData.Name = []*_db.CT_FFName{_aadae}
	return _ddebd.FfData
}

// SetFontFamily sets the Ascii & HAnsi fonly family for a run.
func (_ddbc RunProperties) SetFontFamily(family string) {
	if _ddbc._cadb.RFonts == nil {
		_ddbc._cadb.RFonts = _db.NewCT_Fonts()
	}
	_ddbc._cadb.RFonts.AsciiAttr = _b.String(family)
	_ddbc._cadb.RFonts.HAnsiAttr = _b.String(family)
	_ddbc._cadb.RFonts.EastAsiaAttr = _b.String(family)
}

// FormFieldType is the type of the form field.
//
//go:generate stringer -type=FormFieldType
type FormFieldType byte

// Footers returns the footers defined in the document.
func (_ecb *Document) Footers() []Footer {
	_fbg := []Footer{}
	for _, _fbd := range _ecb._abc {
		_fbg = append(_fbg, Footer{_ecb, _fbd})
	}
	return _fbg
}

// AddTable adds a new table to the document body.
func (_cbae *Document) AddTable() Table {
	_caac := _db.NewEG_BlockLevelElts()
	_cbae._adcb.Body.EG_BlockLevelElts = append(_cbae._adcb.Body.EG_BlockLevelElts, _caac)
	_ddb := _db.NewEG_ContentBlockContent()
	_caac.EG_ContentBlockContent = append(_caac.EG_ContentBlockContent, _ddb)
	_acbg := _db.NewCT_Tbl()
	_ddb.Tbl = append(_ddb.Tbl, _acbg)
	return Table{_cbae, _acbg}
}

// SetBasedOn sets the style that this style is based on.
func (_ccbaf Style) SetBasedOn(name string) {
	if name == "" {
		_ccbaf._eedcd.BasedOn = nil
	} else {
		_ccbaf._eedcd.BasedOn = _db.NewCT_String()
		_ccbaf._eedcd.BasedOn.ValAttr = name
	}
}

// TableInfo is used for keep information about a table, a row and a cell where the text is located.
type TableInfo struct {
	Table    *_db.CT_Tbl
	Row      *_db.CT_Row
	Cell     *_db.CT_Tc
	RowIndex int
	ColIndex int
}

// InsertStyle insert style to styles.
func (_egdd Styles) InsertStyle(ss Style) { _egdd._cdgg.Style = append(_egdd._cdgg.Style, ss.X()) }

// ItalicValue returns the precise nature of the italic setting (unset, off or on).
func (_eeda RunProperties) ItalicValue() OnOffValue { return _egeb(_eeda._cadb.I) }

// X returns the inner wrapped XML type.
func (_efbd NumberingDefinition) X() *_db.CT_AbstractNum { return _efbd._bbgd }

// SetFirstColumn controls the conditional formatting for the first column in a table.
func (_fcgff TableLook) SetFirstColumn(on bool) {
	if !on {
		_fcgff._febce.FirstColumnAttr = &_ca.ST_OnOff{}
		_fcgff._febce.FirstColumnAttr.ST_OnOff1 = _ca.ST_OnOff1Off
	} else {
		_fcgff._febce.FirstColumnAttr = &_ca.ST_OnOff{}
		_fcgff._febce.FirstColumnAttr.ST_OnOff1 = _ca.ST_OnOff1On
	}
}

// Paragraphs returns all of the paragraphs in the document body including tables.
func (_dgab *Document) Paragraphs() []Paragraph {
	_cdd := []Paragraph{}
	if _dgab._adcb.Body == nil {
		return nil
	}
	for _, _gecg := range _dgab._adcb.Body.EG_BlockLevelElts {
		for _, _deg := range _gecg.EG_ContentBlockContent {
			for _, _abg := range _deg.P {
				_cdd = append(_cdd, Paragraph{_dgab, _abg})
			}
		}
	}
	for _, _fbff := range _dgab.Tables() {
		for _, _dfb := range _fbff.Rows() {
			for _, _acee := range _dfb.Cells() {
				_cdd = append(_cdd, _acee.Paragraphs()...)
			}
		}
	}
	return _cdd
}

// Properties returns the cell properties.
func (_dbf Cell) Properties() CellProperties {
	if _dbf._caf.TcPr == nil {
		_dbf._caf.TcPr = _db.NewCT_TcPr()
	}
	return CellProperties{_dbf._caf.TcPr}
}
func (_cffcg Paragraph) addSeparateFldChar() *_db.CT_FldChar {
	_eedf := _cffcg.addFldChar()
	_eedf.FldCharTypeAttr = _db.ST_FldCharTypeSeparate
	return _eedf
}

// Index returns the index of the footer within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (_aacf Footer) Index() int {
	for _gcbb, _gcea := range _aacf._bcge._abc {
		if _gcea == _aacf._eeed {
			return _gcbb
		}
	}
	return -1
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (_ccf CellBorders) SetInsideHorizontal(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_ccf._bee.InsideH = _db.NewCT_Border()
	_ebca(_ccf._bee.InsideH, t, c, thickness)
}

// SetUnderline controls underline for a run style.
func (_cffcc RunProperties) SetUnderline(style _db.ST_Underline, c _bd.Color) {
	if style == _db.ST_UnderlineUnset {
		_cffcc._cadb.U = nil
	} else {
		_cffcc._cadb.U = _db.NewCT_Underline()
		_cffcc._cadb.U.ColorAttr = &_db.ST_HexColor{}
		_cffcc._cadb.U.ColorAttr.ST_HexColorRGB = c.AsRGBString()
		_cffcc._cadb.U.ValAttr = style
	}
}
func _cbfe(_bdade *_db.CT_P, _bgdc *_db.CT_Hyperlink, _edde *TableInfo, _cfdf *DrawingInfo, _fabb []*_db.EG_ContentRunContent) []TextItem {
	_aecd := []TextItem{}
	for _, _fccd := range _fabb {
		if _acafb := _fccd.Sdt; _acafb != nil {
			if _gedga := _acafb.SdtContent; _gedga != nil {
				_aecd = append(_aecd, _cbfe(_bdade, _bgdc, _edde, _cfdf, _gedga.EG_ContentRunContent)...)
			}
		}
		if _ddfa := _fccd.R; _ddfa != nil {
			_afff := _fed.NewBuffer([]byte{})
			for _, _edfg := range _ddfa.EG_RunInnerContent {
				if _edfg.Br != nil {
					_afff.WriteString("\u000a")
				}
				if _edfg.Tab != nil {
					_afff.WriteString("\u0009")
				}
				if _edfg.T != nil && _edfg.T.Content != "" {
					_afff.WriteString(_edfg.T.Content)
				}
				if _edfg.Pict != nil && len(_edfg.Pict.Any) > 0 {
					for _, _aeafc := range _edfg.Pict.Any {
						if _gbgg, _dbda := _aeafc.(*_eb.Shape); _dbda {
							for _, _cddc := range _gbgg.EG_ShapeElements {
								if _dbcf := _cddc.Textbox; _dbcf != nil {
									if _dbcf.TxbxContent != nil {
										_aecd = append(_aecd, _cccg(_dbcf.TxbxContent.EG_ContentBlockContent, nil)...)
									}
								}
							}
						}
					}
				}
			}
			_aecd = append(_aecd, TextItem{Text: _afff.String(), DrawingInfo: _cfdf, Paragraph: _bdade, Hyperlink: _bgdc, Run: _ddfa, TableInfo: _edde})
			for _, _ebfa := range _ddfa.Extra {
				if _cddcd, _ccecbe := _ebfa.(*_db.AlternateContentRun); _ccecbe {
					_gdag := &DrawingInfo{Drawing: _cddcd.Choice.Drawing}
					for _, _cgbf := range _gdag.Drawing.Anchor {
						for _, _aegd := range _cgbf.Graphic.GraphicData.Any {
							if _bgec, _edbb := _aegd.(*_db.WdWsp); _edbb {
								if _bgec.WChoice != nil {
									if _fff := _bgec.SpPr; _fff != nil {
										if _ecgcd := _fff.Xfrm; _ecgcd != nil {
											if _afbec := _ecgcd.Ext; _afbec != nil {
												_gdag.Width = _afbec.CxAttr
												_gdag.Height = _afbec.CyAttr
											}
										}
									}
									for _, _dce := range _bgec.WChoice.Txbx.TxbxContent.EG_ContentBlockContent {
										_aecd = append(_aecd, _eefc(_dce.P, _edde, _gdag)...)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return _aecd
}

// Fonts returns the style's Fonts.
func (_bbbb RunProperties) Fonts() Fonts {
	if _bbbb._cadb.RFonts == nil {
		_bbbb._cadb.RFonts = _db.NewCT_Fonts()
	}
	return Fonts{_bbbb._cadb.RFonts}
}

// SetChecked marks a FormFieldTypeCheckBox as checked or unchecked.
func (_cagc FormField) SetChecked(b bool) {
	if _cagc._fbcee.CheckBox == nil {
		return
	}
	if !b {
		_cagc._fbcee.CheckBox.Checked = nil
	} else {
		_cagc._fbcee.CheckBox.Checked = _db.NewCT_OnOff()
	}
}

// AddTab adds tab to a run and can be used with the the Paragraph's tab stops.
func (_face Run) AddTab() { _bcbd := _face.newIC(); _bcbd.Tab = _db.NewCT_Empty() }

// SetEnabled marks a FormField as enabled or disabled.
func (_gddb FormField) SetEnabled(enabled bool) {
	_gegg := _db.NewCT_OnOff()
	_gegg.ValAttr = &_ca.ST_OnOff{Bool: &enabled}
	_gddb._fbcee.Enabled = []*_db.CT_OnOff{_gegg}
}

// AddParagraph adds a paragraph to the footer.
func (_bgea Footer) AddParagraph() Paragraph {
	_effef := _db.NewEG_ContentBlockContent()
	_bgea._eeed.EG_ContentBlockContent = append(_bgea._eeed.EG_ContentBlockContent, _effef)
	_bdbg := _db.NewCT_P()
	_effef.P = append(_effef.P, _bdbg)
	return Paragraph{_bgea._bcge, _bdbg}
}

// SetStartIndent controls the start indentation.
func (_eegd ParagraphProperties) SetStartIndent(m _bc.Distance) {
	if _eegd._aedc.Ind == nil {
		_eegd._aedc.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_eegd._aedc.Ind.StartAttr = nil
	} else {
		_eegd._aedc.Ind.StartAttr = &_db.ST_SignedTwipsMeasure{}
		_eegd._aedc.Ind.StartAttr.Int64 = _b.Int64(int64(m / _bc.Twips))
	}
}

// GetFooter gets a section Footer for given type
func (_dafce Section) GetFooter(t _db.ST_HdrFtr) (Footer, bool) {
	for _, _fdgf := range _dafce._ggcfb.EG_HdrFtrReferences {
		if _fdgf.FooterReference.TypeAttr == t {
			for _, _ccffg := range _dafce._bagegb.Footers() {
				_gfbaa := _dafce._bagegb._dfc.FindRIDForN(_ccffg.Index(), _b.FooterType)
				if _gfbaa == _fdgf.FooterReference.IdAttr {
					return _ccffg, true
				}
			}
		}
	}
	return Footer{}, false
}

// SetVerticalAlignment sets the vertical alignment of content within a table cell.
func (_gag CellProperties) SetVerticalAlignment(align _db.ST_VerticalJc) {
	if align == _db.ST_VerticalJcUnset {
		_gag._beea.VAlign = nil
	} else {
		_gag._beea.VAlign = _db.NewCT_VerticalJc()
		_gag._beea.VAlign.ValAttr = align
	}
}

// DoubleStrike returns true if paragraph is double striked.
func (_adefc ParagraphProperties) DoubleStrike() bool { return _ebdd(_adefc._aedc.RPr.Dstrike) }

// Definitions returns the defined numbering definitions.
func (_fcae Numbering) Definitions() []NumberingDefinition {
	_begga := []NumberingDefinition{}
	if _fcae._gebaa != nil {
		for _, _aaaa := range _fcae._gebaa.AbstractNum {
			_begga = append(_begga, NumberingDefinition{_aaaa})
		}
	}
	return _begga
}

// SetStyle sets the style of a paragraph and is identical to setting it on the
// paragraph's Properties()
func (_bbad Paragraph) SetStyle(s string) {
	_bbad.ensurePPr()
	if s == "" {
		_bbad._dcfbd.PPr.PStyle = nil
	} else {
		_bbad._dcfbd.PPr.PStyle = _db.NewCT_String()
		_bbad._dcfbd.PPr.PStyle.ValAttr = s
	}
}

// X returns the inner wrapped XML type.
func (_baedce ParagraphProperties) X() *_db.CT_PPr { return _baedce._aedc }

// RemoveParagraph removes a paragraph from a footer.
func (_ceae Header) RemoveParagraph(p Paragraph) {
	for _, _gacbe := range _ceae._ecead.EG_ContentBlockContent {
		for _cced, _eace := range _gacbe.P {
			if _eace == p._dcfbd {
				copy(_gacbe.P[_cced:], _gacbe.P[_cced+1:])
				_gacbe.P = _gacbe.P[0 : len(_gacbe.P)-1]
				return
			}
		}
	}
}
func (_gabb *WatermarkPicture) getShapeType() *_b.XSDAny {
	return _gabb.getInnerElement("\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e")
}
func (_cfdee Document) mergeFields() []mergeFieldInfo {
	_bfac := []Paragraph{}
	_bbea := []mergeFieldInfo{}
	for _, _ecce := range _cfdee.Tables() {
		for _, _eegaa := range _ecce.Rows() {
			for _, _ebgg := range _eegaa.Cells() {
				_bfac = append(_bfac, _ebgg.Paragraphs()...)
			}
		}
	}
	_bfac = append(_bfac, _cfdee.Paragraphs()...)
	for _, _geacb := range _cfdee.Headers() {
		_bfac = append(_bfac, _geacb.Paragraphs()...)
		for _, _gbaa := range _geacb.Tables() {
			for _, _ccgb := range _gbaa.Rows() {
				for _, _adgag := range _ccgb.Cells() {
					_bfac = append(_bfac, _adgag.Paragraphs()...)
				}
			}
		}
	}
	for _, _gfdb := range _cfdee.Footers() {
		_bfac = append(_bfac, _gfdb.Paragraphs()...)
		for _, _afgg := range _gfdb.Tables() {
			for _, _fbeb := range _afgg.Rows() {
				for _, _geege := range _fbeb.Cells() {
					_bfac = append(_bfac, _geege.Paragraphs()...)
				}
			}
		}
	}
	for _, _dcbac := range _bfac {
		_gcad := _dcbac.Runs()
		_egbcb := -1
		_babc := -1
		_fgdcd := -1
		_faebf := mergeFieldInfo{}
		for _, _ebcd := range _dcbac._dcfbd.EG_PContent {
			for _, _fffe := range _ebcd.FldSimple {
				if _eg.Contains(_fffe.InstrAttr, "\u004d\u0045\u0052\u0047\u0045\u0046\u0049\u0045\u004c\u0044") {
					_bgbgb := _aadb(_fffe.InstrAttr)
					_bgbgb._fcaba = true
					_bgbgb._fbfg = _dcbac
					_bgbgb._bcef = _ebcd
					_bbea = append(_bbea, _bgbgb)
				}
			}
		}
		for _efeb := 0; _efeb < len(_gcad); _efeb++ {
			_gdafa := _gcad[_efeb]
			for _, _aaba := range _gdafa.X().EG_RunInnerContent {
				if _aaba.FldChar != nil {
					switch _aaba.FldChar.FldCharTypeAttr {
					case _db.ST_FldCharTypeBegin:
						_egbcb = _efeb
					case _db.ST_FldCharTypeSeparate:
						_babc = _efeb
					case _db.ST_FldCharTypeEnd:
						_fgdcd = _efeb
						if _faebf._gcfd != "" {
							_faebf._fbfg = _dcbac
							_faebf._cbff = _egbcb
							_faebf._eged = _fgdcd
							_faebf._gbef = _babc
							_bbea = append(_bbea, _faebf)
						}
						_egbcb = -1
						_babc = -1
						_fgdcd = -1
						_faebf = mergeFieldInfo{}
					}
				} else if _aaba.InstrText != nil && _eg.Contains(_aaba.InstrText.Content, "\u004d\u0045\u0052\u0047\u0045\u0046\u0049\u0045\u004c\u0044") {
					if _egbcb != -1 && _fgdcd == -1 {
						_faebf = _aadb(_aaba.InstrText.Content)
					}
				}
			}
		}
	}
	return _bbea
}

// SetCSTheme sets the font complex script theme.
func (_febbf Fonts) SetCSTheme(t _db.ST_Theme) { _febbf._bdge.CsthemeAttr = t }

// SetStartIndent controls the start indent of the paragraph.
func (_ceade ParagraphStyleProperties) SetStartIndent(m _bc.Distance) {
	if _ceade._fagf.Ind == nil {
		_ceade._fagf.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_ceade._fagf.Ind.StartAttr = nil
	} else {
		_ceade._fagf.Ind.StartAttr = &_db.ST_SignedTwipsMeasure{}
		_ceade._fagf.Ind.StartAttr.Int64 = _b.Int64(int64(m / _bc.Twips))
	}
}

// Fonts allows manipulating a style or run's fonts.
type Fonts struct{ _bdge *_db.CT_Fonts }

// GetDocRelTargetByID returns TargetAttr of document relationship given its IdAttr.
func (_edf *Document) GetDocRelTargetByID(idAttr string) string {
	for _, _afag := range _edf._dfc.X().Relationship {
		if _afag.IdAttr == idAttr {
			return _afag.TargetAttr
		}
	}
	return ""
}

// SetFollowImageShape sets wrapPath to follow image shape,
// if nil return wrapPath that follow image size.
func (_bba AnchorDrawWrapOptions) SetFollowImageShape(val bool) {
	_bba._dbg = val
	if !val {
		_df, _age := _ffce()
		_bba._ega = _df
		_bba._bb = _age
	}
}

// NumberingDefinition defines a numbering definition for a list of pragraphs.
type NumberingDefinition struct{ _bbgd *_db.CT_AbstractNum }

// Text returns text from the document as one string separated with line breaks.
func (_fageg *DocText) Text() string {
	_ecebd := _fed.NewBuffer([]byte{})
	for _, _gddf := range _fageg.Items {
		if _gddf.Text != "" {
			_ecebd.WriteString(_gddf.Text)
			_ecebd.WriteString("\u000a")
		}
	}
	return _ecebd.String()
}

// IsBold returns true if the run has been set to bold.
func (_fdff RunProperties) IsBold() bool { return _fdff.BoldValue() == OnOffValueOn }

// GetImageObjByRelId returns a common.Image with the associated relation ID in the
// document.
func (_cbec *Document) GetImageObjByRelId(relId string) (_ffa.Image, error) {
	_cbab := _cbec._dfc.GetTargetByRelIdAndType(relId, _b.ImageType)
	if _cbab == "" {
		for _, _affba := range _cbec._gcdc {
			_cbab = _affba.GetTargetByRelIdAndType(relId, _b.ImageType)
		}
	}
	if _cbab == "" {
		for _, _effe := range _cbec._adg {
			_cbab = _effe.GetTargetByRelIdAndType(relId, _b.ImageType)
		}
	}
	return _cbec.DocBase.GetImageBytesByTarget(_cbab)
}
func (_ebdb Footnote) id() int64 { return _ebdb._cefg.IdAttr }

type listItemInfo struct {
	FromStyle      *Style
	FromParagraph  *Paragraph
	AbstractNumId  *int64
	NumberingLevel *NumberingLevel
}

// AddParagraph adds a paragraph to the endnote.
func (_cgdb Endnote) AddParagraph() Paragraph {
	_debe := _db.NewEG_ContentBlockContent()
	_cagg := len(_cgdb._eeabdb.EG_BlockLevelElts[0].EG_ContentBlockContent)
	_cgdb._eeabdb.EG_BlockLevelElts[0].EG_ContentBlockContent = append(_cgdb._eeabdb.EG_BlockLevelElts[0].EG_ContentBlockContent, _debe)
	_cede := _db.NewCT_P()
	var _gdec *_db.CT_String
	if _cagg != 0 {
		_ebge := len(_cgdb._eeabdb.EG_BlockLevelElts[0].EG_ContentBlockContent[_cagg-1].P)
		_gdec = _cgdb._eeabdb.EG_BlockLevelElts[0].EG_ContentBlockContent[_cagg-1].P[_ebge-1].PPr.PStyle
	} else {
		_gdec = _db.NewCT_String()
		_gdec.ValAttr = "\u0045n\u0064\u006e\u006f\u0074\u0065"
	}
	_debe.P = append(_debe.P, _cede)
	_ffbdf := Paragraph{_cgdb._efb, _cede}
	_ffbdf._dcfbd.PPr = _db.NewCT_PPr()
	_ffbdf._dcfbd.PPr.PStyle = _gdec
	_ffbdf._dcfbd.PPr.RPr = _db.NewCT_ParaRPr()
	return _ffbdf
}

// SetBottom sets the cell bottom margin
func (_ef CellMargins) SetBottom(d _bc.Distance) {
	_ef._caff.Bottom = _db.NewCT_TblWidth()
	_eeb(_ef._caff.Bottom, d)
}
func (_adaa *Document) tables(_efdf *_db.EG_ContentBlockContent) []Table {
	_agdb := []Table{}
	for _, _bab := range _efdf.Tbl {
		_agdb = append(_agdb, Table{_adaa, _bab})
		for _, _ddf := range _bab.EG_ContentRowContent {
			for _, _bae := range _ddf.Tr {
				for _, _fbgf := range _bae.EG_ContentCellContent {
					for _, _ddg := range _fbgf.Tc {
						for _, _gdc := range _ddg.EG_BlockLevelElts {
							for _, _fda := range _gdc.EG_ContentBlockContent {
								for _, _baef := range _adaa.tables(_fda) {
									_agdb = append(_agdb, _baef)
								}
							}
						}
					}
				}
			}
		}
	}
	return _agdb
}

// X returns the inner wrapped XML type.
func (_caffg *Document) X() *_db.Document { return _caffg._adcb }

// TableStyleProperties are table properties as defined in a style.
type TableStyleProperties struct{ _dagce *_db.CT_TblPrBase }

// SetCalcOnExit marks if a FormField should be CalcOnExit or not.
func (_cdca FormField) SetCalcOnExit(calcOnExit bool) {
	_fagge := _db.NewCT_OnOff()
	_fagge.ValAttr = &_ca.ST_OnOff{Bool: &calcOnExit}
	_cdca._fbcee.CalcOnExit = []*_db.CT_OnOff{_fagge}
}

// SetHangingIndent controls the hanging indent of the paragraph.
func (_ccfbd ParagraphStyleProperties) SetHangingIndent(m _bc.Distance) {
	if _ccfbd._fagf.Ind == nil {
		_ccfbd._fagf.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_ccfbd._fagf.Ind.HangingAttr = nil
	} else {
		_ccfbd._fagf.Ind.HangingAttr = &_ca.ST_TwipsMeasure{}
		_ccfbd._fagf.Ind.HangingAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(m / _bc.Twips))
	}
}

// SetName marks sets a name attribute for a FormField.
func (_dbec FormField) SetName(name string) {
	_fffg := _db.NewCT_FFName()
	_fffg.ValAttr = &name
	_dbec._fbcee.Name = []*_db.CT_FFName{_fffg}
}

// SearchStylesById returns style by its id.
func (_bbcbc Styles) SearchStyleById(id string) (Style, bool) {
	for _, _edefd := range _bbcbc._cdgg.Style {
		if _edefd.StyleIdAttr != nil {
			if *_edefd.StyleIdAttr == id {
				return Style{_edefd}, true
			}
		}
	}
	return Style{}, false
}

// PutNodeAfter put node to position after relativeTo.
func (_gcedg *Document) PutNodeAfter(relativeTo, node Node) { _gcedg.putNode(relativeTo, node, false) }

// GetImage returns the ImageRef associated with an InlineDrawing.
func (_acba InlineDrawing) GetImage() (_ffa.ImageRef, bool) {
	_cgdab := _acba._adb.Graphic.GraphicData.Any
	if len(_cgdab) > 0 {
		_afade, _aeacb := _cgdab[0].(*_ba.Pic)
		if _aeacb {
			if _afade.BlipFill != nil && _afade.BlipFill.Blip != nil && _afade.BlipFill.Blip.EmbedAttr != nil {
				return _acba._gfgg.GetImageByRelID(*_afade.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return _ffa.ImageRef{}, false
}

// X returns the inner wrapped XML type.
func (_cfda Header) X() *_db.Hdr { return _cfda._ecead }

// ExtractFromFooter returns text from the document footer as an array of TextItems.
func ExtractFromFooter(footer *_db.Ftr) []TextItem { return _cccg(footer.EG_ContentBlockContent, nil) }

// SetAfterSpacing sets spacing below paragraph.
func (_eeaefg Paragraph) SetAfterSpacing(d _bc.Distance) {
	_eeaefg.ensurePPr()
	if _eeaefg._dcfbd.PPr.Spacing == nil {
		_eeaefg._dcfbd.PPr.Spacing = _db.NewCT_Spacing()
	}
	_caea := _eeaefg._dcfbd.PPr.Spacing
	_caea.AfterAttr = &_ca.ST_TwipsMeasure{}
	_caea.AfterAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(d / _bc.Twips))
}

// AddParagraph adds a paragraph to the table cell.
func (_fbe Cell) AddParagraph() Paragraph {
	_fcc := _db.NewEG_BlockLevelElts()
	_fbe._caf.EG_BlockLevelElts = append(_fbe._caf.EG_BlockLevelElts, _fcc)
	_afa := _db.NewEG_ContentBlockContent()
	_fcc.EG_ContentBlockContent = append(_fcc.EG_ContentBlockContent, _afa)
	_ec := _db.NewCT_P()
	_afa.P = append(_afa.P, _ec)
	return Paragraph{_fbe._cec, _ec}
}
func (_ddbea *Document) insertNumberingFromStyleProperties(_gaagb Numbering, _gdge ParagraphStyleProperties) {
	_beeabe := _gdge.NumId()
	_gde := int64(-1)
	if _beeabe > -1 {
		for _, _fabf := range _gaagb._gebaa.Num {
			if _fabf.NumIdAttr == _beeabe {
				if _fabf.AbstractNumId != nil {
					_gde = _fabf.AbstractNumId.ValAttr
					_fdced := false
					for _, _geea := range _ddbea.Numbering._gebaa.Num {
						if _geea.NumIdAttr == _beeabe {
							_fdced = true
							break
						}
					}
					if !_fdced {
						_ddbea.Numbering._gebaa.Num = append(_ddbea.Numbering._gebaa.Num, _fabf)
					}
					break
				}
			}
		}
		for _, _gcb := range _gaagb._gebaa.AbstractNum {
			if _gcb.AbstractNumIdAttr == _gde {
				_adaf := false
				for _, _fbdc := range _ddbea.Numbering._gebaa.AbstractNum {
					if _fbdc.AbstractNumIdAttr == _gde {
						_adaf = true
						break
					}
				}
				if !_adaf {
					_ddbea.Numbering._gebaa.AbstractNum = append(_ddbea.Numbering._gebaa.AbstractNum, _gcb)
				}
				break
			}
		}
	}
}

// Borders returns the ParagraphBorders for setting-up border on paragraph.
func (_bgcc Paragraph) Borders() ParagraphBorders {
	_bgcc.ensurePPr()
	if _bgcc._dcfbd.PPr.PBdr == nil {
		_bgcc._dcfbd.PPr.PBdr = _db.NewCT_PBdr()
	}
	return ParagraphBorders{_bgcc._cfge, _bgcc._dcfbd.PPr.PBdr}
}

// SetAlignment positions an anchored image via alignment.  Offset is
// incompatible with SetOffset, whichever is called last is applied.
func (_bde AnchoredDrawing) SetAlignment(h _db.WdST_AlignH, v _db.WdST_AlignV) {
	_bde.SetHAlignment(h)
	_bde.SetVAlignment(v)
}
func _deaa(_beba *_db.CT_Tbl, _edbf, _ebef map[int64]int64) {
	for _, _fgcc := range _beba.EG_ContentRowContent {
		for _, _acbdd := range _fgcc.Tr {
			for _, _effdc := range _acbdd.EG_ContentCellContent {
				for _, _cceb := range _effdc.Tc {
					for _, _gced := range _cceb.EG_BlockLevelElts {
						for _, _dcfeb := range _gced.EG_ContentBlockContent {
							for _, _dacc := range _dcfeb.P {
								_ceaf(_dacc, _edbf, _ebef)
							}
							for _, _bfegg := range _dcfeb.Tbl {
								_deaa(_bfegg, _edbf, _ebef)
							}
						}
					}
				}
			}
		}
	}
}

// SetWidthPercent sets the cell to a width percentage.
func (_acd CellProperties) SetWidthPercent(pct float64) {
	_acd._beea.TcW = _db.NewCT_TblWidth()
	_acd._beea.TcW.TypeAttr = _db.ST_TblWidthPct
	_acd._beea.TcW.WAttr = &_db.ST_MeasurementOrPercent{}
	_acd._beea.TcW.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_acd._beea.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(pct * 50))
}

// Style is a style within the styles.xml file.
type Style struct{ _eedcd *_db.CT_Style }

// Text return node and its child text,
func (_adgf *Node) Text() string {
	_aabg := _fed.NewBuffer([]byte{})
	switch _befa := _adgf.X().(type) {
	case *Paragraph:
		for _, _bgcb := range _befa.Runs() {
			if _bgcb.Text() != "" {
				_aabg.WriteString(_bgcb.Text())
				_aabg.WriteString("\u000a")
			}
		}
	}
	for _, _fbebd := range _adgf.Children {
		_aabg.WriteString(_fbebd.Text())
	}
	return _aabg.String()
}

// SetUpdateFieldsOnOpen controls if fields are recalculated upon opening the
// document. This is useful for things like a table of contents as the library
// only adds the field code and relies on Word/LibreOffice to actually compute
// the content.
func (_fcafe Settings) SetUpdateFieldsOnOpen(b bool) {
	if !b {
		_fcafe._ccgc.UpdateFields = nil
	} else {
		_fcafe._ccgc.UpdateFields = _db.NewCT_OnOff()
	}
}

// NewWatermarkPicture generates new WatermarkPicture.
func NewWatermarkPicture() WatermarkPicture {
	_bfaf := _eb.NewShapetype()
	_aggda := _eb.NewEG_ShapeElements()
	_aggda.Formulas = _aaffd()
	_aggda.Path = _cgfe()
	_aggda.Lock = _cegge()
	_bfaf.EG_ShapeElements = []*_eb.EG_ShapeElements{_aggda}
	var (
		_dedd  = "\u005f\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0037\u0035"
		_dcbc  = "2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030"
		_fecg  = float32(75.0)
		_ffdag = "\u006d\u0040\u0034\u00405l\u0040\u0034\u0040\u0031\u0031\u0040\u0039\u0040\u0031\u0031\u0040\u0039\u0040\u0035x\u0065"
	)
	_bfaf.IdAttr = &_dedd
	_bfaf.CoordsizeAttr = &_dcbc
	_bfaf.SptAttr = &_fecg
	_bfaf.PreferrelativeAttr = _ca.ST_TrueFalseTrue
	_bfaf.PathAttr = &_ffdag
	_bfaf.FilledAttr = _ca.ST_TrueFalseFalse
	_bfaf.StrokedAttr = _ca.ST_TrueFalseFalse
	_aedbg := _eb.NewShape()
	_aaacf := _eb.NewEG_ShapeElements()
	_aaacf.Imagedata = _gaeed()
	_aedbg.EG_ShapeElements = []*_eb.EG_ShapeElements{_aaacf}
	var (
		_cbbc  = "\u0057\u006f\u0072\u0064\u0050\u0069\u0063\u0074\u0075\u0072e\u0057\u0061\u0074\u0065\u0072\u006d\u0061r\u006b\u0031\u0036\u0033\u0032\u0033\u0031\u0036\u0035\u0039\u0035"
		_bbbg  = "\u005f\u0078\u00300\u0030\u0030\u005f\u0073\u0032\u0030\u0035\u0031"
		_edfd  = "#\u005f\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0037\u0035"
		_dbcb  = ""
		_cagga = "\u0070os\u0069t\u0069o\u006e\u003a\u0061\u0062\u0073\u006fl\u0075\u0074\u0065\u003bm\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074\u003a\u0030\u003bma\u0072\u0067\u0069\u006e\u002d\u0074\u006f\u0070\u003a\u0030\u003b\u0077\u0069\u0064\u0074\u0068\u003a\u0030\u0070\u0074;\u0068e\u0069\u0067\u0068\u0074\u003a\u0030\u0070\u0074\u003b\u007a\u002d\u0069\u006ed\u0065\u0078:\u002d\u0032\u00351\u0036\u0035\u0038\u0032\u0034\u0030\u003b\u006d\u0073o-\u0070\u006f\u0073i\u0074\u0069\u006f\u006e-\u0068\u006f\u0072\u0069\u007a\u006fn\u0074\u0061l\u003a\u0063\u0065\u006e\u0074\u0065\u0072\u003bm\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u002drela\u0074\u0069\u0076\u0065\u003a\u006d\u0061\u0072\u0067\u0069\u006e\u003b\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0076\u0065\u0072t\u0069c\u0061l\u003a\u0063\u0065\u006e\u0074\u0065\u0072\u003b\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e-\u0076\u0065r\u0074\u0069c\u0061l\u002d\u0072\u0065\u006c\u0061\u0074i\u0076\u0065\u003a\u006d\u0061\u0072\u0067\u0069\u006e"
	)
	_aedbg.IdAttr = &_cbbc
	_aedbg.SpidAttr = &_bbbg
	_aedbg.TypeAttr = &_edfd
	_aedbg.AltAttr = &_dbcb
	_aedbg.StyleAttr = &_cagga
	_aedbg.AllowincellAttr = _ca.ST_TrueFalseFalse
	_bfbbb := _db.NewCT_Picture()
	_bfbbb.Any = []_b.Any{_bfaf, _aedbg}
	return WatermarkPicture{_dgdd: _bfbbb, _febge: _aedbg, _baac: _bfaf}
}

// X returns the inner wrapped XML type.
func (_ceaa NumberingLevel) X() *_db.CT_Lvl { return _ceaa._cfed }

// DocRels returns document relationship.
func (_cgb *Document) DocRels() _ffa.Relationships { return _cgb._dfc }

// SetTop sets the top border to a specified type, color and thickness.
func (_gda CellBorders) SetTop(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_gda._bee.Top = _db.NewCT_Border()
	_ebca(_gda._bee.Top, t, c, thickness)
}
func _egeb(_caab *_db.CT_OnOff) OnOffValue {
	if _caab == nil {
		return OnOffValueUnset
	}
	if _caab.ValAttr != nil && _caab.ValAttr.Bool != nil && *_caab.ValAttr.Bool == false {
		return OnOffValueOff
	}
	return OnOffValueOn
}
func (_cecf *Document) putNode(_bcfc, _cbaa Node, _acdf bool) bool {
	_cecf.insertImageFromNode(_cbaa)
	_cecf.insertStyleFromNode(_cbaa)
	switch _dfga := _cbaa._bdcf.(type) {
	case *Paragraph:
		if _edgb, _eefa := _bcfc.X().(*Paragraph); _eefa {
			_cecf.appendParagraph(_edgb, *_dfga, _acdf)
			return true
		} else {
			for _, _badd := range _bcfc.Children {
				if _egaag := _cecf.putNode(_badd, _cbaa, _acdf); _egaag {
					break
				}
			}
		}
	case *Table:
		if _eaea, _beec := _bcfc.X().(*Paragraph); _beec {
			_fefde := _cecf.appendTable(_eaea, *_dfga, _acdf)
			_fefde._cgffe = _dfga._cgffe
			return true
		} else {
			for _, _ecdee := range _bcfc.Children {
				if _dddd := _cecf.putNode(_ecdee, _cbaa, _acdf); _dddd {
					break
				}
			}
		}
	}
	return false
}

// AddRow adds a row to a table.
func (_fbbbb Table) AddRow() Row {
	_dafbfe := _db.NewEG_ContentRowContent()
	_fbbbb._cgffe.EG_ContentRowContent = append(_fbbbb._cgffe.EG_ContentRowContent, _dafbfe)
	_fgfa := _db.NewCT_Row()
	_dafbfe.Tr = append(_dafbfe.Tr, _fgfa)
	return Row{_fbbbb._ffcf, _fgfa}
}

// Clear clears all content within a footer
func (_dgae Footer) Clear() { _dgae._eeed.EG_ContentBlockContent = nil }

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_bfaa ParagraphProperties) SizeMeasure() string {
	if _bgccd := _bfaa._aedc.RPr.Sz; _bgccd != nil {
		_agbg := _bgccd.ValAttr
		if _agbg.ST_PositiveUniversalMeasure != nil {
			return *_agbg.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// Font returns the name of paragraph font family.
func (_defdf ParagraphProperties) Font() string {
	if _fefe := _defdf._aedc.RPr.RFonts; _fefe != nil {
		if _fefe.AsciiAttr != nil {
			return *_fefe.AsciiAttr
		} else if _fefe.HAnsiAttr != nil {
			return *_fefe.HAnsiAttr
		} else if _fefe.CsAttr != nil {
			return *_fefe.CsAttr
		}
	}
	return ""
}
func (_cdg *Document) InsertTableBefore(relativeTo Paragraph) Table {
	return _cdg.insertTable(relativeTo, true)
}

// SetWidth sets the cell width to a specified width.
func (_adc CellProperties) SetWidth(d _bc.Distance) {
	_adc._beea.TcW = _db.NewCT_TblWidth()
	_adc._beea.TcW.TypeAttr = _db.ST_TblWidthDxa
	_adc._beea.TcW.WAttr = &_db.ST_MeasurementOrPercent{}
	_adc._beea.TcW.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_adc._beea.TcW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(d / _bc.Twips))
}

// AnchoredDrawing is an absolutely positioned image within a document page.
type AnchoredDrawing struct {
	_geb *Document
	_cg  *_db.WdAnchor
}

func _aefaa() *_eb.Handles {
	_dbbc := _eb.NewHandles()
	_ffdc := _eb.NewCT_H()
	_dfdfe := "\u0023\u0030\u002c\u0062\u006f\u0074\u0074\u006f\u006dR\u0069\u0067\u0068\u0074"
	_ffdc.PositionAttr = &_dfdfe
	_ceff := "\u0036\u0036\u0032\u0039\u002c\u0031\u0034\u0039\u0037\u0031"
	_ffdc.XrangeAttr = &_ceff
	_dbbc.H = []*_eb.CT_H{_ffdc}
	return _dbbc
}
func _ceaf(_fdcb *_db.CT_P, _aaad, _dddg map[int64]int64) {
	for _, _ebag := range _fdcb.EG_PContent {
		for _, _cccf := range _ebag.EG_ContentRunContent {
			if _cccf.R != nil {
				for _, _febb := range _cccf.R.EG_RunInnerContent {
					_gggg := _febb.EndnoteReference
					if _gggg != nil && _gggg.IdAttr > 0 {
						if _affb, _aebba := _dddg[_gggg.IdAttr]; _aebba {
							_gggg.IdAttr = _affb
						}
					}
					_gbcg := _febb.FootnoteReference
					if _gbcg != nil && _gbcg.IdAttr > 0 {
						if _cffc, _aggd := _aaad[_gbcg.IdAttr]; _aggd {
							_gbcg.IdAttr = _cffc
						}
					}
				}
			}
		}
	}
}

// SetName sets the name of the bookmark. This is the name that is used to
// reference the bookmark from hyperlinks.
func (_dgc Bookmark) SetName(name string) { _dgc._ac.NameAttr = name }

// Nodes contains slice of Node element.
type Nodes struct{ _dcdf []Node }

// TableConditionalFormatting returns a conditional formatting object of a given
// type.  Calling this method repeatedly will return the same object.
func (_fgbc Style) TableConditionalFormatting(typ _db.ST_TblStyleOverrideType) TableConditionalFormatting {
	for _, _acddf := range _fgbc._eedcd.TblStylePr {
		if _acddf.TypeAttr == typ {
			return TableConditionalFormatting{_acddf}
		}
	}
	_cdde := _db.NewCT_TblStylePr()
	_cdde.TypeAttr = typ
	_fgbc._eedcd.TblStylePr = append(_fgbc._eedcd.TblStylePr, _cdde)
	return TableConditionalFormatting{_cdde}
}
func (_egbdf Endnote) id() int64 { return _egbdf._eeabdb.IdAttr }
func _ffgfb() *_eb.Formulas {
	_dacgg := _eb.NewFormulas()
	_dacgg.F = []*_eb.CT_F{_ge.CreateFormula("\u0073\u0075\u006d\u0020\u0023\u0030\u0020\u0030\u00201\u0030\u0038\u0030\u0030"), _ge.CreateFormula("p\u0072\u006f\u0064\u0020\u0023\u0030\u0020\u0032\u0020\u0031"), _ge.CreateFormula("\u0073\u0075\u006d\u0020\u0032\u0031\u0036\u0030\u0030 \u0030\u0020\u0040\u0031"), _ge.CreateFormula("\u0073\u0075\u006d\u0020\u0030\u0020\u0030\u0020\u0040\u0032"), _ge.CreateFormula("\u0073\u0075\u006d\u0020\u0032\u0031\u0036\u0030\u0030 \u0030\u0020\u0040\u0033"), _ge.CreateFormula("\u0069\u0066\u0020\u0040\u0030\u0020\u0040\u0033\u0020\u0030"), _ge.CreateFormula("\u0069\u0066\u0020\u0040\u0030\u0020\u0032\u0031\u00360\u0030\u0020\u0040\u0031"), _ge.CreateFormula("\u0069\u0066\u0020\u0040\u0030\u0020\u0030\u0020\u0040\u0032"), _ge.CreateFormula("\u0069\u0066\u0020\u0040\u0030\u0020\u0040\u0034\u00202\u0031\u0036\u0030\u0030"), _ge.CreateFormula("\u006di\u0064\u0020\u0040\u0035\u0020\u00406"), _ge.CreateFormula("\u006di\u0064\u0020\u0040\u0038\u0020\u00405"), _ge.CreateFormula("\u006di\u0064\u0020\u0040\u0037\u0020\u00408"), _ge.CreateFormula("\u006di\u0064\u0020\u0040\u0036\u0020\u00407"), _ge.CreateFormula("s\u0075\u006d\u0020\u0040\u0036\u0020\u0030\u0020\u0040\u0035")}
	return _dacgg
}
func _aefa(_egcb *_aa.CT_Blip, _fddd map[string]string) {
	if _egcb.EmbedAttr != nil {
		if _edag, _gefc := _fddd[*_egcb.EmbedAttr]; _gefc {
			*_egcb.EmbedAttr = _edag
		}
	}
}

// PutNodeBefore put node to position before relativeTo.
func (_agab *Document) PutNodeBefore(relativeTo, node Node) { _agab.putNode(relativeTo, node, true) }

// NewWatermarkText generates a new WatermarkText.
func NewWatermarkText() WatermarkText {
	_dcae := _eb.NewShapetype()
	_fcdac := _eb.NewEG_ShapeElements()
	_fcdac.Formulas = _ffgfb()
	_fcdac.Path = _ecdeg()
	_fcdac.Textpath = _cfgc()
	_fcdac.Handles = _aefaa()
	_fcdac.Lock = _fffb()
	_dcae.EG_ShapeElements = []*_eb.EG_ShapeElements{_fcdac}
	var (
		_fgee   = "_\u0078\u0030\u0030\u0030\u0030\u005f\u0074\u0031\u0033\u0036"
		_gdefa  = "2\u0031\u0036\u0030\u0030\u002c\u0032\u0031\u0036\u0030\u0030"
		_ecdff  = float32(136.0)
		_cfeceg = "\u0031\u0030\u00380\u0030"
		_cfabe  = "m\u0040\u0037\u002c\u006c\u0040\u0038,\u006d\u0040\u0035\u002c\u0032\u0031\u0036\u0030\u0030l\u0040\u0036\u002c2\u00316\u0030\u0030\u0065"
	)
	_dcae.IdAttr = &_fgee
	_dcae.CoordsizeAttr = &_gdefa
	_dcae.SptAttr = &_ecdff
	_dcae.AdjAttr = &_cfeceg
	_dcae.PathAttr = &_cfabe
	_eefd := _eb.NewShape()
	_ceag := _eb.NewEG_ShapeElements()
	_ceag.Textpath = _dagg()
	_eefd.EG_ShapeElements = []*_eb.EG_ShapeElements{_ceag}
	var (
		_decd   = "\u0050\u006f\u0077\u0065\u0072\u0050l\u0075\u0073\u0057\u0061\u0074\u0065\u0072\u004d\u0061\u0072\u006b\u004f\u0062j\u0065\u0063\u0074\u0031\u0033\u0036\u00380\u0030\u0038\u0038\u0036"
		_fbgbgc = "\u005f\u0078\u00300\u0030\u0030\u005f\u0073\u0032\u0030\u0035\u0031"
		_feed   = "\u0023\u005f\u00780\u0030\u0030\u0030\u005f\u0074\u0031\u0033\u0036"
		_dfde   = ""
		_dfda   = "\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u003a\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065f\u0074:\u0030\u003b\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074o\u0070\u003a\u0030\u003b\u0077\u0069\u0064\u0074\u0068\u003a\u0034\u0036\u0038\u0070\u0074;\u0068\u0065\u0069\u0067\u0068\u0074\u003a\u0032\u0033\u0034\u0070\u0074\u003b\u007a\u002d\u0069\u006ede\u0078\u003a\u002d\u0032\u0035\u0031\u0036\u0035\u0031\u0030\u0037\u0032\u003b\u006d\u0073\u006f\u002d\u0077\u0072\u0061\u0070\u002d\u0065\u0064\u0069\u0074\u0065\u0064\u003a\u0066\u003b\u006d\u0073\u006f\u002d\u0077\u0069\u0064\u0074\u0068\u002d\u0070\u0065\u0072\u0063\u0065\u006e\u0074\u003a\u0030\u003b\u006d\u0073\u006f\u002d\u0068\u0065\u0069\u0067h\u0074-p\u0065\u0072\u0063\u0065\u006et\u003a\u0030\u003b\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006fn\u0074\u0061\u006c\u003a\u0063\u0065\u006e\u0074\u0065\u0072\u003b\u006d\u0073\u006f\u002d\u0070o\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u002d\u0072\u0065l\u0061\u0074\u0069\u0076\u0065:\u006d\u0061\u0072\u0067\u0069n\u003b\u006d\u0073o\u002d\u0070\u006f\u0073\u0069\u0074\u0069o\u006e-\u0076\u0065\u0072\u0074\u0069\u0063\u0061\u006c\u003a\u0063\u0065\u006e\u0074\u0065\u0072\u003b\u006d\u0073\u006f\u002d\u0070\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u002d\u0076\u0065r\u0074\u0069\u0063\u0061\u006c\u002d\u0072e\u006c\u0061\u0074i\u0076\u0065\u003a\u006d\u0061\u0072\u0067\u0069\u006e\u003b\u006d\u0073\u006f\u002d\u0077\u0069\u0064\u0074\u0068\u002d\u0070\u0065\u0072\u0063e\u006e\u0074\u003a\u0030\u003b\u006d\u0073\u006f\u002dh\u0065\u0069\u0067\u0068t\u002dp\u0065\u0072\u0063\u0065\u006et\u003a0"
		_ggbcg  = "\u0073\u0069\u006c\u0076\u0065\u0072"
	)
	_eefd.IdAttr = &_decd
	_eefd.SpidAttr = &_fbgbgc
	_eefd.TypeAttr = &_feed
	_eefd.AltAttr = &_dfde
	_eefd.StyleAttr = &_dfda
	_eefd.AllowincellAttr = _ca.ST_TrueFalseFalse
	_eefd.FillcolorAttr = &_ggbcg
	_eefd.StrokedAttr = _ca.ST_TrueFalseFalse
	_eddcc := _db.NewCT_Picture()
	_eddcc.Any = []_b.Any{_dcae, _eefd}
	return WatermarkText{_gdgad: _eddcc, _bgdbe: _eefd, _ecgff: _dcae}
}

// Clear clears the styes.
func (_agce Styles) Clear() {
	_agce._cdgg.DocDefaults = nil
	_agce._cdgg.LatentStyles = nil
	_agce._cdgg.Style = nil
}

// ExtractFromHeader returns text from the document header as an array of TextItems.
func ExtractFromHeader(header *_db.Hdr) []TextItem { return _cccg(header.EG_ContentBlockContent, nil) }

// Properties returns the run properties.
func (_bgcca Run) Properties() RunProperties {
	if _bgcca._afec.RPr == nil {
		_bgcca._afec.RPr = _db.NewCT_RPr()
	}
	return RunProperties{_bgcca._afec.RPr}
}

// SetConformance sets conformance attribute of the document
// as one of these values from github.com/arcpd/msword/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (_gadeg Document) SetConformance(conformanceAttr _ca.ST_ConformanceClass) {
	_gadeg._adcb.ConformanceAttr = conformanceAttr
}
func (_fbf *chart) X() *_ee.ChartSpace { return _fbf._bbd }

// SetShapeStyle sets style to the element v:shape in watermark.
func (_eggb *WatermarkPicture) SetShapeStyle(shapeStyle _ge.ShapeStyle) {
	if _eggb._febge != nil {
		_fbge := shapeStyle.String()
		_eggb._febge.StyleAttr = &_fbge
	}
}

// RemoveParagraph removes a paragraph from a document.
func (_dca *Document) RemoveParagraph(p Paragraph) {
	if _dca._adcb.Body == nil {
		return
	}
	for _, _bfb := range _dca._adcb.Body.EG_BlockLevelElts {
		for _, _dbfe := range _bfb.EG_ContentBlockContent {
			for _fgfd, _bfc := range _dbfe.P {
				if _bfc == p._dcfbd {
					copy(_dbfe.P[_fgfd:], _dbfe.P[_fgfd+1:])
					_dbfe.P = _dbfe.P[0 : len(_dbfe.P)-1]
					return
				}
			}
			if _dbfe.Sdt != nil && _dbfe.Sdt.SdtContent != nil && _dbfe.Sdt.SdtContent.P != nil {
				for _fddf, _ded := range _dbfe.Sdt.SdtContent.P {
					if _ded == p._dcfbd {
						copy(_dbfe.P[_fddf:], _dbfe.P[_fddf+1:])
						_dbfe.P = _dbfe.P[0 : len(_dbfe.P)-1]
						return
					}
				}
			}
		}
	}
	for _, _eedg := range _dca.Tables() {
		for _, _gac := range _eedg.Rows() {
			for _, _gfdf := range _gac.Cells() {
				for _, _deff := range _gfdf._caf.EG_BlockLevelElts {
					for _, _add := range _deff.EG_ContentBlockContent {
						for _aeda, _efed := range _add.P {
							if _efed == p._dcfbd {
								copy(_add.P[_aeda:], _add.P[_aeda+1:])
								_add.P = _add.P[0 : len(_add.P)-1]
								return
							}
						}
					}
				}
			}
		}
	}
	for _, _ecg := range _dca.Headers() {
		_ecg.RemoveParagraph(p)
	}
	for _, _fcf := range _dca.Footers() {
		_fcf.RemoveParagraph(p)
	}
}

// Runs returns all of the runs in a paragraph.
func (_dcdfg Paragraph) Runs() []Run {
	_cgfaf := []Run{}
	for _, _fbacb := range _dcdfg._dcfbd.EG_PContent {
		if _fbacb.Hyperlink != nil && _fbacb.Hyperlink.EG_ContentRunContent != nil {
			for _, _bdef := range _fbacb.Hyperlink.EG_ContentRunContent {
				if _bdef.R != nil {
					_cgfaf = append(_cgfaf, Run{_dcdfg._cfge, _bdef.R})
				}
			}
		}
		for _, _gdgcb := range _fbacb.EG_ContentRunContent {
			if _gdgcb.R != nil {
				_cgfaf = append(_cgfaf, Run{_dcdfg._cfge, _gdgcb.R})
			}
			if _gdgcb.Sdt != nil && _gdgcb.Sdt.SdtContent != nil {
				for _, _cdfgb := range _gdgcb.Sdt.SdtContent.EG_ContentRunContent {
					if _cdfgb.R != nil {
						_cgfaf = append(_cgfaf, Run{_dcdfg._cfge, _cdfgb.R})
					}
				}
			}
		}
	}
	return _cgfaf
}

// ReplaceTextByRegexp replace the text within node using regexp expression.
func (_dgbe *Node) ReplaceTextByRegexp(rgx *_ff.Regexp, newText string) {
	switch _cgee := _dgbe.X().(type) {
	case *Paragraph:
		for _, _acedg := range _cgee.Runs() {
			for _, _ebffg := range _acedg._afec.EG_RunInnerContent {
				if _ebffg.T != nil {
					_debf := _ebffg.T.Content
					_debf = rgx.ReplaceAllString(_debf, newText)
					_ebffg.T.Content = _debf
				}
			}
		}
	}
	for _, _bgaf := range _dgbe.Children {
		_bgaf.ReplaceTextByRegexp(rgx, newText)
	}
}

// Append appends a document d0 to a document d1. All settings, headers and footers remain the same as in the document d0 if they exist there, otherwise they are taken from the d1.
func (_gaag *Document) Append(d1orig *Document) error {
	_cabc, _gdgb := d1orig.Copy()
	if _gdgb != nil {
		return _gdgb
	}
	_gaag.DocBase = _gaag.DocBase.Append(_cabc.DocBase)
	if _cabc._adcb.ConformanceAttr != _ca.ST_ConformanceClassStrict {
		_gaag._adcb.ConformanceAttr = _cabc._adcb.ConformanceAttr
	}
	_abbc := _gaag._dfc.X().Relationship
	_gfff := _cabc._dfc.X().Relationship
	_gcaec := _cabc._adcb.Body
	_fged := map[string]string{}
	_effa := map[int64]int64{}
	_adef := map[int64]int64{}
	for _, _dge := range _gfff {
		_dded := true
		_egbd := _dge.IdAttr
		_fgab := _dge.TargetAttr
		_fgcfb := _dge.TypeAttr
		_dafbf := _fgcfb == _b.ImageType
		_ccbab := _fgcfb == _b.HyperLinkType
		var _cafb string
		for _, _acfd := range _abbc {
			if _acfd.TypeAttr == _fgcfb && _acfd.TargetAttr == _fgab {
				_dded = false
				_cafb = _acfd.IdAttr
				break
			}
		}
		if _dafbf {
			_baedf := "\u0077\u006f\u0072d\u002f" + _fgab
			for _, _addg := range _cabc.DocBase.Images {
				if _addg.Target() == _baedf {
					_acda, _dedg := _ffa.ImageFromStorage(_addg.Path())
					if _dedg != nil {
						return _dedg
					}
					_eedc, _dedg := _gaag.AddImage(_acda)
					if _dedg != nil {
						return _dedg
					}
					_cafb = _eedc.RelID()
					break
				}
			}
		} else if _dded {
			if _ccbab {
				_afbd := _gaag._dfc.AddHyperlink(_fgab)
				_cafb = _ffa.Relationship(_afbd).ID()
			} else {
				_eacc := _gaag._dfc.AddRelationship(_fgab, _fgcfb)
				_cafb = _eacc.X().IdAttr
			}
		}
		if _egbd != _cafb {
			_fged[_egbd] = _cafb
		}
	}
	if _gcaec.SectPr != nil {
		for _, _agda := range _gcaec.SectPr.EG_HdrFtrReferences {
			if _agda.HeaderReference != nil {
				if _ecec, _beeag := _fged[_agda.HeaderReference.IdAttr]; _beeag {
					_agda.HeaderReference.IdAttr = _ecec
					_gaag._gcdc = append(_gaag._gcdc, _ffa.NewRelationships())
				}
			} else if _agda.FooterReference != nil {
				if _gccb, _agdae := _fged[_agda.FooterReference.IdAttr]; _agdae {
					_agda.FooterReference.IdAttr = _gccb
					_gaag._adg = append(_gaag._adg, _ffa.NewRelationships())
				}
			}
		}
	}
	_aef, _cbe := _gaag._acdc, _cabc._acdc
	if _aef != nil {
		if _cbe != nil {
			if _aef.Endnote != nil {
				if _cbe.Endnote != nil {
					_gfeea := int64(len(_aef.Endnote) + 1)
					for _, _cffbb := range _cbe.Endnote {
						_afbda := _cffbb.IdAttr
						if _afbda > 0 {
							_cffbb.IdAttr = _gfeea
							_aef.Endnote = append(_aef.Endnote, _cffbb)
							_adef[_afbda] = _gfeea
							_gfeea++
						}
					}
				}
			} else {
				_aef.Endnote = _cbe.Endnote
			}
		}
	} else if _cbe != nil {
		_aef = _cbe
	}
	_gaag._acdc = _aef
	_cdbgd, _acac := _gaag._ffcef, _cabc._ffcef
	if _cdbgd != nil {
		if _acac != nil {
			if _cdbgd.Footnote != nil {
				if _acac.Footnote != nil {
					_dcfeg := int64(len(_cdbgd.Footnote) + 1)
					for _, _fada := range _acac.Footnote {
						_bfbe := _fada.IdAttr
						if _bfbe > 0 {
							_fada.IdAttr = _dcfeg
							_cdbgd.Footnote = append(_cdbgd.Footnote, _fada)
							_effa[_bfbe] = _dcfeg
							_dcfeg++
						}
					}
				}
			} else {
				_cdbgd.Footnote = _acac.Footnote
			}
		}
	} else if _acac != nil {
		_cdbgd = _acac
	}
	_gaag._ffcef = _cdbgd
	for _, _bfga := range _gcaec.EG_BlockLevelElts {
		for _, _aega := range _bfga.EG_ContentBlockContent {
			for _, _eddc := range _aega.P {
				_fdda(_eddc, _fged)
				_ffgb(_eddc, _fged)
				_ceaf(_eddc, _effa, _adef)
			}
			for _, _bbca := range _aega.Tbl {
				_bcd(_bbca, _fged)
				_fced(_bbca, _fged)
				_deaa(_bbca, _effa, _adef)
			}
		}
	}
	_gaag._adcb.Body.EG_BlockLevelElts = append(_gaag._adcb.Body.EG_BlockLevelElts, _cabc._adcb.Body.EG_BlockLevelElts...)
	if _gaag._adcb.Body.SectPr == nil {
		_gaag._adcb.Body.SectPr = _cabc._adcb.Body.SectPr
	} else {
		var _dec, _gggb bool
		for _, _bdgc := range _gaag._adcb.Body.SectPr.EG_HdrFtrReferences {
			if _bdgc.HeaderReference != nil {
				_dec = true
			} else if _bdgc.FooterReference != nil {
				_gggb = true
			}
		}
		if !_dec {
			for _, _gaaa := range _cabc._adcb.Body.SectPr.EG_HdrFtrReferences {
				if _gaaa.HeaderReference != nil {
					_gaag._adcb.Body.SectPr.EG_HdrFtrReferences = append(_gaag._adcb.Body.SectPr.EG_HdrFtrReferences, _gaaa)
					break
				}
			}
		}
		if !_gggb {
			for _, _gccd := range _cabc._adcb.Body.SectPr.EG_HdrFtrReferences {
				if _gccd.FooterReference != nil {
					_gaag._adcb.Body.SectPr.EG_HdrFtrReferences = append(_gaag._adcb.Body.SectPr.EG_HdrFtrReferences, _gccd)
					break
				}
			}
		}
		if _gaag._adcb.Body.SectPr.Cols == nil && _cabc._adcb.Body.SectPr.Cols != nil {
			_gaag._adcb.Body.SectPr.Cols = _cabc._adcb.Body.SectPr.Cols
		}
	}
	_ceebe := _gaag.Numbering._gebaa
	_gfgc := _cabc.Numbering._gebaa
	if _ceebe == nil && _gfgc != nil {
		_ceebe = _gfgc
	}
	_gaag.Numbering._gebaa = _ceebe
	if _gaag.Styles._cdgg == nil && _cabc.Styles._cdgg != nil {
		_gaag.Styles._cdgg = _cabc.Styles._cdgg
	}
	_gaag._cdf = append(_gaag._cdf, _cabc._cdf...)
	_gaag._aede = append(_gaag._aede, _cabc._aede...)
	if len(_gaag._adf) == 0 {
		_gaag._adf = _cabc._adf
	}
	if len(_gaag._abc) == 0 {
		_gaag._abc = _cabc._abc
	}
	_bbb := _gaag._eaf
	_gbc := _cabc._eaf
	if _bbb != nil {
		if _gbc != nil {
			if _bbb.Divs != nil {
				if _gbc.Divs != nil {
					_bbb.Divs.Div = append(_bbb.Divs.Div, _gbc.Divs.Div...)
				}
			} else {
				_bbb.Divs = _gbc.Divs
			}
		}
		_bbb.Frameset = nil
	} else if _gbc != nil {
		_bbb = _gbc
		_bbb.Frameset = nil
	}
	_gaag._eaf = _bbb
	_dcff := _gaag._begg
	_ecgg := _cabc._begg
	if _dcff != nil {
		if _ecgg != nil {
			if _dcff.Font != nil {
				if _ecgg.Font != nil {
					for _, _cffae := range _ecgg.Font {
						_cdbcb := true
						for _, _efdc := range _dcff.Font {
							if _efdc.NameAttr == _cffae.NameAttr {
								_cdbcb = false
								break
							}
						}
						if _cdbcb {
							_dcff.Font = append(_dcff.Font, _cffae)
						}
					}
				}
			} else {
				_dcff.Font = _ecgg.Font
			}
		}
	} else if _ecgg != nil {
		_dcff = _ecgg
	}
	_gaag._begg = _dcff
	return nil
}

// SetSize sets size attribute for a FormFieldTypeCheckBox in pt.
func (_ccgdb FormField) SetSize(size uint64) {
	size *= 2
	if _ccgdb._fbcee.CheckBox != nil {
		_ccgdb._fbcee.CheckBox.Choice = _db.NewCT_FFCheckBoxChoice()
		_ccgdb._fbcee.CheckBox.Choice.Size = _db.NewCT_HpsMeasure()
		_ccgdb._fbcee.CheckBox.Choice.Size.ValAttr = _db.ST_HpsMeasure{ST_UnsignedDecimalNumber: &size}
	}
}

// SetPageMargins sets the page margins for a section
func (_eaffd Section) SetPageMargins(top, right, bottom, left, header, footer, gutter _bc.Distance) {
	_cgbe := _db.NewCT_PageMar()
	_cgbe.TopAttr.Int64 = _b.Int64(int64(top / _bc.Twips))
	_cgbe.BottomAttr.Int64 = _b.Int64(int64(bottom / _bc.Twips))
	_cgbe.RightAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(right / _bc.Twips))
	_cgbe.LeftAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(left / _bc.Twips))
	_cgbe.HeaderAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(header / _bc.Twips))
	_cgbe.FooterAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(footer / _bc.Twips))
	_cgbe.GutterAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(gutter / _bc.Twips))
	_eaffd._ggcfb.PgMar = _cgbe
}

// SetPrimaryStyle marks the style as a primary style.
func (_cgff Style) SetPrimaryStyle(b bool) {
	if b {
		_cgff._eedcd.QFormat = _db.NewCT_OnOff()
	} else {
		_cgff._eedcd.QFormat = nil
	}
}

// AddTabStop adds a tab stop to the paragraph.
func (_ddgfb ParagraphStyleProperties) AddTabStop(position _bc.Distance, justificaton _db.ST_TabJc, leader _db.ST_TabTlc) {
	if _ddgfb._fagf.Tabs == nil {
		_ddgfb._fagf.Tabs = _db.NewCT_Tabs()
	}
	_fedgb := _db.NewCT_TabStop()
	_fedgb.LeaderAttr = leader
	_fedgb.ValAttr = justificaton
	_fedgb.PosAttr.Int64 = _b.Int64(int64(position / _bc.Twips))
	_ddgfb._fagf.Tabs.Tab = append(_ddgfb._fagf.Tabs.Tab, _fedgb)
}

// SetWidthPercent sets the table to a width percentage.
func (_daag TableProperties) SetWidthPercent(pct float64) {
	_daag._faec.TblW = _db.NewCT_TblWidth()
	_daag._faec.TblW.TypeAttr = _db.ST_TblWidthPct
	_daag._faec.TblW.WAttr = &_db.ST_MeasurementOrPercent{}
	_daag._faec.TblW.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_daag._faec.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(pct * 50))
}

// AddText adds text to a run.
func (_dgbc Run) AddText(s string) {
	_ddbbf := _db.NewEG_RunInnerContent()
	_dgbc._afec.EG_RunInnerContent = append(_dgbc._afec.EG_RunInnerContent, _ddbbf)
	_ddbbf.T = _db.NewCT_Text()
	if _b.NeedsSpacePreserve(s) {
		_badcf := "\u0070\u0072\u0065\u0073\u0065\u0072\u0076\u0065"
		_ddbbf.T.SpaceAttr = &_badcf
	}
	_ddbbf.T.Content = s
}

// ComplexSizeValue returns the value of run font size for complex fonts in points.
func (_fgede RunProperties) ComplexSizeValue() float64 {
	if _deaag := _fgede._cadb.SzCs; _deaag != nil {
		_fedc := _deaag.ValAttr
		if _fedc.ST_UnsignedDecimalNumber != nil {
			return float64(*_fedc.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// AddHyperLink adds a new hyperlink to a parapgraph.
func (_ddfb Paragraph) AddHyperLink() HyperLink {
	_bbfd := _db.NewEG_PContent()
	_ddfb._dcfbd.EG_PContent = append(_ddfb._dcfbd.EG_PContent, _bbfd)
	_bbfd.Hyperlink = _db.NewCT_Hyperlink()
	return HyperLink{_ddfb._cfge, _bbfd.Hyperlink}
}

// SetHorizontalBanding controls the conditional formatting for horizontal banding.
func (_dbegd TableLook) SetHorizontalBanding(on bool) {
	if !on {
		_dbegd._febce.NoHBandAttr = &_ca.ST_OnOff{}
		_dbegd._febce.NoHBandAttr.ST_OnOff1 = _ca.ST_OnOff1On
	} else {
		_dbegd._febce.NoHBandAttr = &_ca.ST_OnOff{}
		_dbegd._febce.NoHBandAttr.ST_OnOff1 = _ca.ST_OnOff1Off
	}
}

// SetOutlineLvl sets outline level of paragraph.
func (_bfccc Paragraph) SetOutlineLvl(lvl int64) {
	_bfccc.ensurePPr()
	if _bfccc._dcfbd.PPr.OutlineLvl == nil {
		_bfccc._dcfbd.PPr.OutlineLvl = _db.NewCT_DecimalNumber()
	}
	_becf := lvl - 1
	_bfccc._dcfbd.PPr.OutlineLvl.ValAttr = _becf
}

// SetStyle sets the style of a paragraph.
func (_bgccc ParagraphProperties) SetStyle(s string) {
	if s == "" {
		_bgccc._aedc.PStyle = nil
	} else {
		_bgccc._aedc.PStyle = _db.NewCT_String()
		_bgccc._aedc.PStyle.ValAttr = s
	}
}

// AddField adds a field (automatically computed text) to the document.
func (_efad Run) AddField(code string) { _efad.AddFieldWithFormatting(code, "", true) }

// X returns the inner wrapped XML type.
func (_cabec Paragraph) X() *_db.CT_P { return _cabec._dcfbd }

// X returns the inner wrapped XML type.
func (_fgg AnchoredDrawing) X() *_db.WdAnchor { return _fgg._cg }

// Endnotes returns the endnotes defined in the document.
func (_gdf *Document) Endnotes() []Endnote {
	_dfdf := []Endnote{}
	for _, _fbc := range _gdf._acdc.CT_Endnotes.Endnote {
		_dfdf = append(_dfdf, Endnote{_gdf, _fbc})
	}
	return _dfdf
}

// NewNumbering constructs a new numbering.
func NewNumbering() Numbering { _ggfdb := _db.NewNumbering(); return Numbering{_ggfdb} }

// IsFootnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (_gacad Run) IsFootnote() (bool, int64) {
	if _gacad._afec.EG_RunInnerContent != nil {
		if _gacad._afec.EG_RunInnerContent[0].FootnoteReference != nil {
			return true, _gacad._afec.EG_RunInnerContent[0].FootnoteReference.IdAttr
		}
	}
	return false, 0
}

// SetAll sets all of the borders to a given value.
func (_bbc CellBorders) SetAll(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_bbc.SetBottom(t, c, thickness)
	_bbc.SetLeft(t, c, thickness)
	_bbc.SetRight(t, c, thickness)
	_bbc.SetTop(t, c, thickness)
	_bbc.SetInsideHorizontal(t, c, thickness)
	_bbc.SetInsideVertical(t, c, thickness)
}

// GetColor returns the color.Color object representing the run color.
func (_badf RunProperties) GetColor() _bd.Color {
	if _gdga := _badf._cadb.Color; _gdga != nil {
		_bfgb := _gdga.ValAttr
		if _bfgb.ST_HexColorRGB != nil {
			return _bd.FromHex(*_bfgb.ST_HexColorRGB)
		}
	}
	return _bd.Color{}
}
func (_gcef Paragraph) insertRun(_fcgfe Run, _cegfg bool) Run {
	for _, _dfecf := range _gcef._dcfbd.EG_PContent {
		for _fadf, _bdeba := range _dfecf.EG_ContentRunContent {
			if _bdeba.R == _fcgfe.X() {
				_abeea := _db.NewCT_R()
				_dfecf.EG_ContentRunContent = append(_dfecf.EG_ContentRunContent, nil)
				if _cegfg {
					copy(_dfecf.EG_ContentRunContent[_fadf+1:], _dfecf.EG_ContentRunContent[_fadf:])
					_dfecf.EG_ContentRunContent[_fadf] = _db.NewEG_ContentRunContent()
					_dfecf.EG_ContentRunContent[_fadf].R = _abeea
				} else {
					copy(_dfecf.EG_ContentRunContent[_fadf+2:], _dfecf.EG_ContentRunContent[_fadf+1:])
					_dfecf.EG_ContentRunContent[_fadf+1] = _db.NewEG_ContentRunContent()
					_dfecf.EG_ContentRunContent[_fadf+1].R = _abeea
				}
				return Run{_gcef._cfge, _abeea}
			}
			if _bdeba.Sdt != nil && _bdeba.Sdt.SdtContent != nil {
				for _, _eagb := range _bdeba.Sdt.SdtContent.EG_ContentRunContent {
					if _eagb.R == _fcgfe.X() {
						_badb := _db.NewCT_R()
						_bdeba.Sdt.SdtContent.EG_ContentRunContent = append(_bdeba.Sdt.SdtContent.EG_ContentRunContent, nil)
						if _cegfg {
							copy(_bdeba.Sdt.SdtContent.EG_ContentRunContent[_fadf+1:], _bdeba.Sdt.SdtContent.EG_ContentRunContent[_fadf:])
							_bdeba.Sdt.SdtContent.EG_ContentRunContent[_fadf] = _db.NewEG_ContentRunContent()
							_bdeba.Sdt.SdtContent.EG_ContentRunContent[_fadf].R = _badb
						} else {
							copy(_bdeba.Sdt.SdtContent.EG_ContentRunContent[_fadf+2:], _bdeba.Sdt.SdtContent.EG_ContentRunContent[_fadf+1:])
							_bdeba.Sdt.SdtContent.EG_ContentRunContent[_fadf+1] = _db.NewEG_ContentRunContent()
							_bdeba.Sdt.SdtContent.EG_ContentRunContent[_fadf+1].R = _badb
						}
						return Run{_gcef._cfge, _badb}
					}
				}
			}
		}
	}
	return _gcef.AddRun()
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (_gdfeg RunProperties) RStyle() string {
	if _gdfeg._cadb.RStyle != nil {
		return _gdfeg._cadb.RStyle.ValAttr
	}
	return ""
}

// SetSpacing sets the spacing that comes before and after the paragraph.
func (_cffaf ParagraphStyleProperties) SetSpacing(before, after _bc.Distance) {
	if _cffaf._fagf.Spacing == nil {
		_cffaf._fagf.Spacing = _db.NewCT_Spacing()
	}
	if before == _bc.Zero {
		_cffaf._fagf.Spacing.BeforeAttr = nil
	} else {
		_cffaf._fagf.Spacing.BeforeAttr = &_ca.ST_TwipsMeasure{}
		_cffaf._fagf.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(before / _bc.Twips))
	}
	if after == _bc.Zero {
		_cffaf._fagf.Spacing.AfterAttr = nil
	} else {
		_cffaf._fagf.Spacing.AfterAttr = &_ca.ST_TwipsMeasure{}
		_cffaf._fagf.Spacing.AfterAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(after / _bc.Twips))
	}
}

// X returns the inner wrapped XML type.
func (_cfac ParagraphStyleProperties) X() *_db.CT_PPrGeneral { return _cfac._fagf }

// Numbering is the document wide numbering styles contained in numbering.xml.
type Numbering struct{ _gebaa *_db.Numbering }

func (_gaaee Paragraph) addFldChar() *_db.CT_FldChar {
	_afaf := _gaaee.AddRun()
	_eeagg := _afaf.X()
	_dafbe := _db.NewEG_RunInnerContent()
	_cgead := _db.NewCT_FldChar()
	_dafbe.FldChar = _cgead
	_eeagg.EG_RunInnerContent = append(_eeagg.EG_RunInnerContent, _dafbe)
	return _cgead
}

// SetHAlignment sets the horizontal alignment for an anchored image.
func (_ccg AnchoredDrawing) SetHAlignment(h _db.WdST_AlignH) {
	_ccg._cg.PositionH.Choice = &_db.WdCT_PosHChoice{}
	_ccg._cg.PositionH.Choice.Align = h
}

// SetWrapPathLineTo sets wrapPath lineTo value.
func (_faa AnchorDrawWrapOptions) SetWrapPathLineTo(coordinates []*_aa.CT_Point2D) {
	_faa._bb = coordinates
}

// InsertParagraphBefore adds a new empty paragraph before the relativeTo
// paragraph.
func (_feff *Document) InsertParagraphBefore(relativeTo Paragraph) Paragraph {
	return _feff.insertParagraph(relativeTo, true)
}

// FindNodeByStyleId return slice of node base on style id.
func (_geee *Nodes) FindNodeByStyleId(styleId string) []Node {
	_ddfd := []Node{}
	for _, _feee := range _geee._dcdf {
		switch _ddee := _feee._bdcf.(type) {
		case *Paragraph:
			if _ddee != nil && _ddee.Style() == styleId {
				_ddfd = append(_ddfd, _feee)
			}
		case *Table:
			if _ddee != nil && _ddee.Style() == styleId {
				_ddfd = append(_ddfd, _feee)
			}
		}
		_gacae := Nodes{_dcdf: _feee.Children}
		_ddfd = append(_ddfd, _gacae.FindNodeByStyleId(styleId)...)
	}
	return _ddfd
}

// AppendNode append node to document element.
func (_cgea *Document) AppendNode(node Node) {
	_cgea.insertImageFromNode(node)
	_cgea.insertStyleFromNode(node)
	for _, _gaae := range node.Children {
		_cgea.insertImageFromNode(_gaae)
		_cgea.insertStyleFromNode(_gaae)
	}
	switch _faee := node.X().(type) {
	case *Paragraph:
		_cgea.appendParagraph(nil, *_faee, false)
	case *Table:
		_cgea.appendTable(nil, *_faee, false)
	}
	if node._cffcb != nil {
		if node._cffcb._cdf != nil {
			if _geac := _cgea._dfc.FindRIDForN(0, _b.ThemeType); _geac == "" {
				if _fadgc := node._cffcb._dfc.FindRIDForN(0, _b.ThemeType); _fadgc != "" {
					_cgea._cdf = append(_cgea._cdf, node._cffcb._cdf...)
					_abd := node._cffcb._dfc.GetTargetByRelId(_fadgc)
					_cgea.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_abd, "\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e/\u0076\u006e\u0064.\u006f\u0070e\u006e\u0078\u006d\u006c\u0066\u006fr\u006dat\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0074\u0068\u0065\u006d\u0065\u002b\u0078\u006d\u006c")
					_cgea._dfc.AddRelationship(_abd, _b.ThemeType)
				}
			}
		}
		_cffg := _cgea._begg
		_aadcf := node._cffcb._begg
		if _cffg != nil {
			if _aadcf != nil {
				if _cffg.Font != nil {
					if _aadcf.Font != nil {
						for _, _bgbe := range _aadcf.Font {
							_gada := true
							for _, _fdaa := range _cffg.Font {
								if _fdaa.NameAttr == _bgbe.NameAttr {
									_gada = false
									break
								}
							}
							if _gada {
								_cffg.Font = append(_cffg.Font, _bgbe)
							}
						}
					}
				} else {
					_cffg.Font = _aadcf.Font
				}
			}
		} else if _aadcf != nil {
			_cffg = _aadcf
		}
		_cgea._begg = _cffg
		if _gcdcf := _cgea._dfc.FindRIDForN(0, _b.FontTableType); _gcdcf == "" {
			_cgea.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072d/\u0066\u006f\u006e\u0074\u0054\u0061\u0062\u006c\u0065\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069c\u0061\u0074\u0069\u006f\u006e\u002f\u0076n\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063e\u0073\u0073\u0069\u006e\u0067\u006d\u006c\u002e\u0066\u006f\u006e\u0074T\u0061\u0062\u006c\u0065\u002b\u0078m\u006c")
			_cgea._dfc.AddRelationship("\u0066\u006f\u006e\u0074\u0054\u0061\u0062\u006c\u0065\u002e\u0078\u006d\u006c", _b.FontTableType)
		}
	}
}

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (_dbcdd RunProperties) CharacterSpacingMeasure() string {
	if _edadb := _dbcdd._cadb.Spacing; _edadb != nil {
		_gfce := _edadb.ValAttr
		if _gfce.ST_UniversalMeasure != nil {
			return *_gfce.ST_UniversalMeasure
		}
	}
	return ""
}
func (_gabef Paragraph) addFldCharsForField(_bcde, _bgef string) FormField {
	_ebfdc := _gabef.addBeginFldChar(_bcde)
	_ffdg := FormField{_fbcee: _ebfdc}
	_abdg := _gabef._cfge.Bookmarks()
	_eafde := int64(len(_abdg))
	if _bcde != "" {
		_gabef.addStartBookmark(_eafde, _bcde)
	}
	_gabef.addInstrText(_bgef)
	_gabef.addSeparateFldChar()
	if _bgef == "\u0046\u004f\u0052\u004d\u0054\u0045\u0058\u0054" {
		_dagfcb := _gabef.AddRun()
		_degbd := _db.NewEG_RunInnerContent()
		_dagfcb._afec.EG_RunInnerContent = []*_db.EG_RunInnerContent{_degbd}
		_ffdg._efgb = _degbd
	}
	_gabef.addEndFldChar()
	if _bcde != "" {
		_gabef.addEndBookmark(_eafde)
	}
	return _ffdg
}

// AddCell adds a cell to a row and returns it
func (_fcdd Row) AddCell() Cell {
	_edbda := _db.NewEG_ContentCellContent()
	_fcdd._bageg.EG_ContentCellContent = append(_fcdd._bageg.EG_ContentCellContent, _edbda)
	_cggg := _db.NewCT_Tc()
	_edbda.Tc = append(_edbda.Tc, _cggg)
	return Cell{_fcdd._eefba, _cggg}
}

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// document.
func (_ddda *Document) GetImageByRelID(relID string) (_ffa.ImageRef, bool) {
	_ceeb := _ddda._dfc.GetTargetByRelId(relID)
	_aebcc := ""
	for _, _fgeac := range _ddda._gcdc {
		if _aebcc != "" {
			break
		}
		_aebcc = _fgeac.GetTargetByRelId(relID)
	}
	for _, _feec := range _ddda.Images {
		if _feec.RelID() == relID {
			return _feec, true
		}
		if _ceeb != "" {
			_debc := _eg.Replace(_feec.Target(), "\u0077\u006f\u0072d\u002f", "", 1)
			if _debc == _ceeb {
				if _feec.RelID() == "" {
					_feec.SetRelID(relID)
				}
				return _feec, true
			}
		}
		if _aebcc != "" {
			_faed := _eg.Replace(_feec.Target(), "\u0077\u006f\u0072d\u002f", "", 1)
			if _faed == _aebcc {
				if _feec.RelID() == "" {
					_feec.SetRelID(relID)
				}
				return _feec, true
			}
		}
	}
	return _ffa.ImageRef{}, false
}

// SetOrigin sets the origin of the image.  It defaults to ST_RelFromHPage and
// ST_RelFromVPage
func (_ebd AnchoredDrawing) SetOrigin(h _db.WdST_RelFromH, v _db.WdST_RelFromV) {
	_ebd._cg.PositionH.RelativeFromAttr = h
	_ebd._cg.PositionV.RelativeFromAttr = v
}

// Type returns the type of the field.
func (_dadg FormField) Type() FormFieldType {
	if _dadg._fbcee.TextInput != nil {
		return FormFieldTypeText
	} else if _dadg._fbcee.CheckBox != nil {
		return FormFieldTypeCheckBox
	} else if _dadg._fbcee.DdList != nil {
		return FormFieldTypeDropDown
	}
	return FormFieldTypeUnknown
}

const (
	OnOffValueUnset OnOffValue = iota
	OnOffValueOff
	OnOffValueOn
)

// AnchorDrawWrapOptions is options to set
// wrapPolygon for wrap text through and tight.
type AnchorDrawWrapOptions struct {
	_dbg bool
	_ega *_aa.CT_Point2D
	_bb  []*_aa.CT_Point2D
}

// SetWrapPathStart sets wrapPath start value.
func (_dgg AnchorDrawWrapOptions) SetWrapPathStart(coordinate *_aa.CT_Point2D) {
	_dgg._ega = coordinate
}
func (_ggff *WatermarkPicture) findNode(_defa *_b.XSDAny, _cdabg string) *_b.XSDAny {
	for _, _fdegf := range _defa.Nodes {
		if _fdegf.XMLName.Local == _cdabg {
			return _fdegf
		}
	}
	return nil
}

// Tables returns the tables defined in the document.
func (_fafc *Document) Tables() []Table {
	_caeb := []Table{}
	if _fafc._adcb.Body == nil {
		return nil
	}
	for _, _gcc := range _fafc._adcb.Body.EG_BlockLevelElts {
		for _, _fdad := range _gcc.EG_ContentBlockContent {
			for _, _efc := range _fafc.tables(_fdad) {
				_caeb = append(_caeb, _efc)
			}
		}
	}
	return _caeb
}

// X returns the inner wrapped XML type.
func (_cddb InlineDrawing) X() *_db.WdInline { return _cddb._adb }

// Cell is a table cell within a document (not a spreadsheet)
type Cell struct {
	_cec *Document
	_caf *_db.CT_Tc
}

// SetAfterAuto controls if spacing after a paragraph is automatically determined.
func (_ddfbd ParagraphSpacing) SetAfterAuto(b bool) {
	if b {
		_ddfbd._ffeb.AfterAutospacingAttr = &_ca.ST_OnOff{}
		_ddfbd._ffeb.AfterAutospacingAttr.Bool = _b.Bool(true)
	} else {
		_ddfbd._ffeb.AfterAutospacingAttr = nil
	}
}

// X returns the inner wml.CT_TblBorders
func (_egfg TableBorders) X() *_db.CT_TblBorders { return _egfg._fcdb }

// Caps returns true if paragraph font is capitalized.
func (_fbgag ParagraphProperties) Caps() bool { return _ebdd(_fbgag._aedc.RPr.Caps) }

// SetAlignment set alignment of paragraph.
func (_efggb Paragraph) SetAlignment(alignment _db.ST_Jc) {
	_efggb.ensurePPr()
	if _efggb._dcfbd.PPr.Jc == nil {
		_efggb._dcfbd.PPr.Jc = _db.NewCT_Jc()
	}
	_efggb._dcfbd.PPr.Jc.ValAttr = alignment
}

// Bookmark is a bookmarked location within a document that can be referenced
// with a hyperlink.
type Bookmark struct{ _ac *_db.CT_Bookmark }

// GetSize return the size of anchor on the page.
func (_dc AnchoredDrawing) GetSize() (_ce, _bea int64) {
	return _dc._cg.Extent.CxAttr, _dc._cg.Extent.CyAttr
}

// SetWindowControl controls if the first or last line of the paragraph is
// allowed to display on a separate page.
func (_fead ParagraphProperties) SetWindowControl(b bool) {
	if !b {
		_fead._aedc.WidowControl = nil
	} else {
		_fead._aedc.WidowControl = _db.NewCT_OnOff()
	}
}

// RunProperties returns the run properties controlling text formatting within the table.
func (_cfbba TableConditionalFormatting) RunProperties() RunProperties {
	if _cfbba._adeac.RPr == nil {
		_cfbba._adeac.RPr = _db.NewCT_RPr()
	}
	return RunProperties{_cfbba._adeac.RPr}
}

// Caps returns true if run font is capitalized.
func (_cgef RunProperties) Caps() bool { return _ebdd(_cgef._cadb.Caps) }

// IsEndnote returns a bool based on whether the run has a
// footnote or not. Returns both a bool as to whether it has
// a footnote as well as the ID of the footnote.
func (_gfagc Run) IsEndnote() (bool, int64) {
	if _gfagc._afec.EG_RunInnerContent != nil {
		if _gfagc._afec.EG_RunInnerContent[0].EndnoteReference != nil {
			return true, _gfagc._afec.EG_RunInnerContent[0].EndnoteReference.IdAttr
		}
	}
	return false, 0
}

// SetRight sets the right border to a specified type, color and thickness.
func (_cecc TableBorders) SetRight(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_cecc._fcdb.Right = _db.NewCT_Border()
	_ebca(_cecc._fcdb.Right, t, c, thickness)
}

// Imprint returns true if run imprint is on.
func (_egbe RunProperties) Imprint() bool { return _ebdd(_egbe._cadb.Imprint) }

// SetWidthAuto sets the the table width to automatic.
func (_bacbd TableProperties) SetWidthAuto() {
	_bacbd._faec.TblW = _db.NewCT_TblWidth()
	_bacbd._faec.TblW.TypeAttr = _db.ST_TblWidthAuto
}

// X returns the inner wrapped XML type.
func (_ab Cell) X() *_db.CT_Tc { return _ab._caf }

// GetColor returns the color.Color object representing the run color.
func (_egbda ParagraphProperties) GetColor() _bd.Color {
	if _gccc := _egbda._aedc.RPr.Color; _gccc != nil {
		_adgd := _gccc.ValAttr
		if _adgd.ST_HexColorRGB != nil {
			return _bd.FromHex(*_adgd.ST_HexColorRGB)
		}
	}
	return _bd.Color{}
}

// SetCharacterSpacing sets the run's Character Spacing Adjustment.
func (_bbaeb RunProperties) SetCharacterSpacing(size _bc.Distance) {
	_bbaeb._cadb.Spacing = _db.NewCT_SignedTwipsMeasure()
	_bbaeb._cadb.Spacing.ValAttr.Int64 = _b.Int64(int64(size / _bc.Twips))
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_dgdg ParagraphProperties) ComplexSizeMeasure() string {
	if _fedac := _dgdg._aedc.RPr.SzCs; _fedac != nil {
		_baee := _fedac.ValAttr
		if _baee.ST_PositiveUniversalMeasure != nil {
			return *_baee.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// RemoveEndnote removes a endnote from both the paragraph and the document
// the requested endnote must be anchored on the paragraph being referenced.
func (_adec Paragraph) RemoveEndnote(id int64) {
	_gfbe := _adec._cfge._acdc
	var _fcb int
	for _ggda, _ccgbb := range _gfbe.CT_Endnotes.Endnote {
		if _ccgbb.IdAttr == id {
			_fcb = _ggda
		}
	}
	_fcb = 0
	_gfbe.CT_Endnotes.Endnote[_fcb] = nil
	_gfbe.CT_Endnotes.Endnote[_fcb] = _gfbe.CT_Endnotes.Endnote[len(_gfbe.CT_Endnotes.Endnote)-1]
	_gfbe.CT_Endnotes.Endnote = _gfbe.CT_Endnotes.Endnote[:len(_gfbe.CT_Endnotes.Endnote)-1]
	var _bfae Run
	for _, _dccgg := range _adec.Runs() {
		if _affgb, _gcgg := _dccgg.IsEndnote(); _affgb {
			if _gcgg == id {
				_bfae = _dccgg
			}
		}
	}
	_adec.RemoveRun(_bfae)
}

// X returns the inner wrapped XML type.
func (_beab CellProperties) X() *_db.CT_TcPr { return _beab._beea }

// Control returns an *axcontrol.Control object contained in the run or the nil value in case of no controls.
func (_aggc Run) Control() *_ae.Control {
	if _eggd := _aggc._afec.EG_RunInnerContent; _eggd != nil {
		if _cded := _eggd[0].Object; _cded != nil {
			if _cbcg := _cded.Choice; _cbcg != nil {
				if _eedfb := _cbcg.Control; _eedfb != nil {
					if _eedfb.IdAttr != nil {
						_cgde := _aggc._gacf.GetDocRelTargetByID(*_eedfb.IdAttr)
						for _, _afdeg := range _aggc._gacf._aede {
							if _cgde == _afdeg.TargetAttr {
								return _afdeg
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// SetVerticalAlignment controls the vertical alignment of the run, this is used
// to control if text is superscript/subscript.
func (_befag RunProperties) SetVerticalAlignment(v _ca.ST_VerticalAlignRun) {
	if v == _ca.ST_VerticalAlignRunUnset {
		_befag._cadb.VertAlign = nil
	} else {
		_befag._cadb.VertAlign = _db.NewCT_VerticalAlignRun()
		_befag._cadb.VertAlign.ValAttr = v
	}
}

// Paragraphs returns the paragraphs defined in an endnote.
func (_gfdfb Endnote) Paragraphs() []Paragraph {
	_gaea := []Paragraph{}
	for _, _badda := range _gfdfb.content() {
		for _, _adae := range _badda.P {
			_gaea = append(_gaea, Paragraph{_gfdfb._efb, _adae})
		}
	}
	return _gaea
}
func _cfgc() *_eb.Textpath {
	_gbff := _eb.NewTextpath()
	_gbff.OnAttr = _ca.ST_TrueFalseTrue
	_gbff.FitshapeAttr = _ca.ST_TrueFalseTrue
	return _gbff
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (_bcgd TableBorders) SetBottom(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_bcgd._fcdb.Bottom = _db.NewCT_Border()
	_ebca(_bcgd._fcdb.Bottom, t, c, thickness)
}

type mergeFieldInfo struct {
	_gcfd               string
	_fggfa              string
	_bedf               string
	_adedf              bool
	_edece              bool
	_gdcee              bool
	_afbb               bool
	_fbfg               Paragraph
	_cbff, _gbef, _eged int
	_bcef               *_db.EG_PContent
	_fcaba              bool
}

// SetLeft sets the left border to a specified type, color and thickness.
func (_dba CellBorders) SetLeft(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_dba._bee.Left = _db.NewCT_Border()
	_ebca(_dba._bee.Left, t, c, thickness)
}

// RStyle returns the name of character style.
// It is defined here http://officeopenxml.com/WPstyleCharStyles.php
func (_bfag ParagraphProperties) RStyle() string {
	if _bfag._aedc.RPr.RStyle != nil {
		return _bfag._aedc.RPr.RStyle.ValAttr
	}
	return ""
}

// TextWithOptions extract text with options.
func (_gafd *DocText) TextWithOptions(options ExtractTextOptions) string {
	_cebbc := make(map[int64]map[int64]int64, 0)
	_cgbfc := _fed.NewBuffer([]byte{})
	_afbc := int64(0)
	_bedb := int64(0)
	_fdbe := int64(0)
	for _dbbg, _eddca := range _gafd.Items {
		_dgff := false
		if _eddca.Text != "" {
			if _dbbg > 0 {
				if _eddca.Paragraph != _gafd.Items[_dbbg-1].Paragraph {
					_dgff = true
				}
				if !options.RunsOnNewLine && _dgff {
					_cgbfc.WriteString("\u000a")
				} else if options.RunsOnNewLine {
					_cgbfc.WriteString("\u000a")
				}
			} else {
				_dgff = true
			}
			if options.WithNumbering {
				if _dgff {
					for _, _abge := range _gafd._bbaf {
						if _abge.FromParagraph == nil {
							continue
						}
						if _abge.FromParagraph.X() == _eddca.Paragraph {
							if _feaa := _abge.NumberingLevel.X(); _feaa != nil {
								if _abge.AbstractNumId != nil && _gafd._bgfd[*_abge.AbstractNumId][_feaa.IlvlAttr] > 0 {
									if _, _edcd := _cebbc[*_abge.AbstractNumId]; _edcd {
										if _, _adfg := _cebbc[*_abge.AbstractNumId][_feaa.IlvlAttr]; _adfg {
											_cebbc[*_abge.AbstractNumId][_feaa.IlvlAttr]++
										} else {
											_cebbc[*_abge.AbstractNumId][_feaa.IlvlAttr] = 1
										}
									} else {
										_cebbc[*_abge.AbstractNumId] = map[int64]int64{_feaa.IlvlAttr: 1}
									}
									if _afbc == _abge.NumberingLevel.X().IlvlAttr && _feaa.IlvlAttr > 0 {
										_bedb++
									} else {
										_bedb = _cebbc[*_abge.AbstractNumId][_feaa.IlvlAttr]
										if _feaa.IlvlAttr > _afbc && _fdbe == *_abge.AbstractNumId {
											_bedb = 1
										}
									}
									_fbac := ""
									if _feaa.LvlText.ValAttr != nil {
										_fbac = *_feaa.LvlText.ValAttr
									}
									_bafb := _ad.FormatNumberingText(_bedb, _feaa.IlvlAttr, _fbac, _feaa.NumFmt, _cebbc[*_abge.AbstractNumId])
									_cgbfc.WriteString(_bafb)
									_gafd._bgfd[*_abge.AbstractNumId][_feaa.IlvlAttr]--
									_afbc = _feaa.IlvlAttr
									_fdbe = *_abge.AbstractNumId
									if options.NumberingIndent != "" {
										_cgbfc.WriteString(options.NumberingIndent)
									}
								}
							}
							break
						}
					}
				}
			}
			_cgbfc.WriteString(_eddca.Text)
		}
	}
	return _cgbfc.String()
}

// AddDrawingInline adds an inline drawing from an ImageRef.
func (_dfafg Run) AddDrawingInline(img _ffa.ImageRef) (InlineDrawing, error) {
	_cbbde := _dfafg.newIC()
	_cbbde.Drawing = _db.NewCT_Drawing()
	_bfba := _db.NewWdInline()
	_eagd := InlineDrawing{_dfafg._gacf, _bfba}
	_bfba.CNvGraphicFramePr = _aa.NewCT_NonVisualGraphicFrameProperties()
	_cbbde.Drawing.Inline = append(_cbbde.Drawing.Inline, _bfba)
	_bfba.Graphic = _aa.NewGraphic()
	_bfba.Graphic.GraphicData = _aa.NewCT_GraphicalObjectData()
	_bfba.Graphic.GraphicData.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"
	_bfba.DistTAttr = _b.Uint32(0)
	_bfba.DistLAttr = _b.Uint32(0)
	_bfba.DistBAttr = _b.Uint32(0)
	_bfba.DistRAttr = _b.Uint32(0)
	_bfba.Extent.CxAttr = int64(float64(img.Size().X*_bc.Pixel72) / _bc.EMU)
	_bfba.Extent.CyAttr = int64(float64(img.Size().Y*_bc.Pixel72) / _bc.EMU)
	_aafda := 0x7FFFFFFF & _f.Uint32()
	_bfba.DocPr.IdAttr = _aafda
	_fbec := _ba.NewPic()
	_fbec.NvPicPr.CNvPr.IdAttr = _aafda
	_ceab := img.RelID()
	if _ceab == "" {
		return _eagd, _bg.New("\u0063\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020\u0074\u006f\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0020\u0072\u0065l\u0061\u0074\u0069o\u006e\u0073")
	}
	_bfba.Graphic.GraphicData.Any = append(_bfba.Graphic.GraphicData.Any, _fbec)
	_fbec.BlipFill = _aa.NewCT_BlipFillProperties()
	_fbec.BlipFill.Blip = _aa.NewCT_Blip()
	_fbec.BlipFill.Blip.EmbedAttr = &_ceab
	_fbec.BlipFill.Stretch = _aa.NewCT_StretchInfoProperties()
	_fbec.BlipFill.Stretch.FillRect = _aa.NewCT_RelativeRect()
	_fbec.SpPr = _aa.NewCT_ShapeProperties()
	_fbec.SpPr.Xfrm = _aa.NewCT_Transform2D()
	_fbec.SpPr.Xfrm.Off = _aa.NewCT_Point2D()
	_fbec.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _b.Int64(0)
	_fbec.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _b.Int64(0)
	_fbec.SpPr.Xfrm.Ext = _aa.NewCT_PositiveSize2D()
	_fbec.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * _bc.Point)
	_fbec.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * _bc.Point)
	_fbec.SpPr.PrstGeom = _aa.NewCT_PresetGeometry2D()
	_fbec.SpPr.PrstGeom.PrstAttr = _aa.ST_ShapeTypeRect
	return _eagd, nil
}

// AddDrawingAnchored adds an anchored (floating) drawing from an ImageRef.
func (_ebfec Run) AddDrawingAnchored(img _ffa.ImageRef) (AnchoredDrawing, error) {
	_efbda := _ebfec.newIC()
	_efbda.Drawing = _db.NewCT_Drawing()
	_cdea := _db.NewWdAnchor()
	_gbeeb := AnchoredDrawing{_ebfec._gacf, _cdea}
	_cdea.SimplePosAttr = _b.Bool(false)
	_cdea.AllowOverlapAttr = true
	_cdea.CNvGraphicFramePr = _aa.NewCT_NonVisualGraphicFrameProperties()
	_efbda.Drawing.Anchor = append(_efbda.Drawing.Anchor, _cdea)
	_cdea.Graphic = _aa.NewGraphic()
	_cdea.Graphic.GraphicData = _aa.NewCT_GraphicalObjectData()
	_cdea.Graphic.GraphicData.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"
	_cdea.SimplePos.XAttr.ST_CoordinateUnqualified = _b.Int64(0)
	_cdea.SimplePos.YAttr.ST_CoordinateUnqualified = _b.Int64(0)
	_cdea.PositionH.RelativeFromAttr = _db.WdST_RelFromHPage
	_cdea.PositionH.Choice = &_db.WdCT_PosHChoice{}
	_cdea.PositionH.Choice.PosOffset = _b.Int32(0)
	_cdea.PositionV.RelativeFromAttr = _db.WdST_RelFromVPage
	_cdea.PositionV.Choice = &_db.WdCT_PosVChoice{}
	_cdea.PositionV.Choice.PosOffset = _b.Int32(0)
	_cdea.Extent.CxAttr = int64(float64(img.Size().X*_bc.Pixel72) / _bc.EMU)
	_cdea.Extent.CyAttr = int64(float64(img.Size().Y*_bc.Pixel72) / _bc.EMU)
	_cdea.Choice = &_db.WdEG_WrapTypeChoice{}
	_cdea.Choice.WrapSquare = _db.NewWdCT_WrapSquare()
	_cdea.Choice.WrapSquare.WrapTextAttr = _db.WdST_WrapTextBothSides
	_gegce := 0x7FFFFFFF & _f.Uint32()
	_cdea.DocPr.IdAttr = _gegce
	_dfba := _ba.NewPic()
	_dfba.NvPicPr.CNvPr.IdAttr = _gegce
	_gcda := img.RelID()
	if _gcda == "" {
		return _gbeeb, _bg.New("\u0063\u006f\u0075\u006c\u0064\u006e\u0027\u0074\u0020\u0066\u0069\u006e\u0064\u0020\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065\u0020\u0074\u006f\u0020\u0069\u006d\u0061g\u0065\u0020\u0077\u0069\u0074\u0068\u0069\u006e\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u0020\u0072\u0065l\u0061\u0074\u0069o\u006e\u0073")
	}
	_cdea.Graphic.GraphicData.Any = append(_cdea.Graphic.GraphicData.Any, _dfba)
	_dfba.BlipFill = _aa.NewCT_BlipFillProperties()
	_dfba.BlipFill.Blip = _aa.NewCT_Blip()
	_dfba.BlipFill.Blip.EmbedAttr = &_gcda
	_dfba.BlipFill.Stretch = _aa.NewCT_StretchInfoProperties()
	_dfba.BlipFill.Stretch.FillRect = _aa.NewCT_RelativeRect()
	_dfba.SpPr = _aa.NewCT_ShapeProperties()
	_dfba.SpPr.Xfrm = _aa.NewCT_Transform2D()
	_dfba.SpPr.Xfrm.Off = _aa.NewCT_Point2D()
	_dfba.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _b.Int64(0)
	_dfba.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _b.Int64(0)
	_dfba.SpPr.Xfrm.Ext = _aa.NewCT_PositiveSize2D()
	_dfba.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * _bc.Point)
	_dfba.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * _bc.Point)
	_dfba.SpPr.PrstGeom = _aa.NewCT_PresetGeometry2D()
	_dfba.SpPr.PrstGeom.PrstAttr = _aa.ST_ShapeTypeRect
	return _gbeeb, nil
}

// SetTop sets the top border to a specified type, color and thickness.
func (_ceca TableBorders) SetTop(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_ceca._fcdb.Top = _db.NewCT_Border()
	_ebca(_ceca._fcdb.Top, t, c, thickness)
}

// Shadow returns true if run shadow is on.
func (_gdbg RunProperties) Shadow() bool { return _ebdd(_gdbg._cadb.Shadow) }

// Value returns the tring value of a FormFieldTypeText or FormFieldTypeDropDown.
func (_baga FormField) Value() string {
	if _baga._fbcee.TextInput != nil && _baga._efgb.T != nil {
		return _baga._efgb.T.Content
	} else if _baga._fbcee.DdList != nil && _baga._fbcee.DdList.Result != nil {
		_ddab := _baga.PossibleValues()
		_efcfg := int(_baga._fbcee.DdList.Result.ValAttr)
		if _efcfg < len(_ddab) {
			return _ddab[_efcfg]
		}
	} else if _baga._fbcee.CheckBox != nil {
		if _baga.IsChecked() {
			return "\u0074\u0072\u0075\u0065"
		}
		return "\u0066\u0061\u006cs\u0065"
	}
	return ""
}

// SetEmboss sets the run to embossed text.
func (_aagb RunProperties) SetEmboss(b bool) {
	if !b {
		_aagb._cadb.Emboss = nil
	} else {
		_aagb._cadb.Emboss = _db.NewCT_OnOff()
	}
}

// AddEndnote will create a new endnote and attach it to the Paragraph in the
// location at the end of the previous run (endnotes create their own run within
// the paragraph. The text given to the function is simply a convenience helper,
// paragraphs and runs can always be added to the text of the endnote later.
func (_acbda Paragraph) AddEndnote(text string) Endnote {
	var _cega int64
	if _acbda._cfge.HasEndnotes() {
		for _, _ebfeg := range _acbda._cfge.Endnotes() {
			if _ebfeg.id() > _cega {
				_cega = _ebfeg.id()
			}
		}
		_cega++
	} else {
		_cega = 0
		_acbda._cfge._acdc = &_db.Endnotes{}
	}
	_gegc := _db.NewCT_FtnEdn()
	_fceb := _db.NewCT_FtnEdnRef()
	_fceb.IdAttr = _cega
	_acbda._cfge._acdc.CT_Endnotes.Endnote = append(_acbda._cfge._acdc.CT_Endnotes.Endnote, _gegc)
	_gceb := _acbda.AddRun()
	_fgdb := _gceb.Properties()
	_fgdb.SetStyle("\u0045\u006e\u0064\u006e\u006f\u0074\u0065\u0041\u006e\u0063\u0068\u006f\u0072")
	_gceb._afec.EG_RunInnerContent = []*_db.EG_RunInnerContent{_db.NewEG_RunInnerContent()}
	_gceb._afec.EG_RunInnerContent[0].EndnoteReference = _fceb
	_fafga := Endnote{_acbda._cfge, _gegc}
	_fafga._eeabdb.IdAttr = _cega
	_fafga._eeabdb.EG_BlockLevelElts = []*_db.EG_BlockLevelElts{_db.NewEG_BlockLevelElts()}
	_bfedf := _fafga.AddParagraph()
	_bfedf.Properties().SetStyle("\u0045n\u0064\u006e\u006f\u0074\u0065")
	_bfedf._dcfbd.PPr.RPr = _db.NewCT_ParaRPr()
	_gfbf := _bfedf.AddRun()
	_gfbf.AddTab()
	_gfbf.AddText(text)
	return _fafga
}

// Numbering return numbering that being use by paragraph.
func (_dcbgc Paragraph) Numbering() Numbering {
	_dcbgc.ensurePPr()
	_ffbb := NewNumbering()
	if _dcbgc._dcfbd.PPr.NumPr != nil {
		_gfbdg := int64(-1)
		_ebaaa := int64(-1)
		if _dcbgc._dcfbd.PPr.NumPr.NumId != nil {
			_gfbdg = _dcbgc._dcfbd.PPr.NumPr.NumId.ValAttr
		}
		for _, _acacg := range _dcbgc._cfge.Numbering._gebaa.Num {
			if _gfbdg < 0 {
				break
			}
			if _acacg.NumIdAttr == _gfbdg {
				if _acacg.AbstractNumId != nil {
					_ebaaa = _acacg.AbstractNumId.ValAttr
					_ffbb._gebaa.Num = append(_ffbb._gebaa.Num, _acacg)
					break
				}
			}
		}
		for _, _egceg := range _dcbgc._cfge.Numbering._gebaa.AbstractNum {
			if _ebaaa < 0 {
				break
			}
			if _egceg.AbstractNumIdAttr == _ebaaa {
				_ffbb._gebaa.AbstractNum = append(_ffbb._gebaa.AbstractNum, _egceg)
				break
			}
		}
	}
	return _ffbb
}

// AddSection adds a new document section with an optional section break.  If t
// is ST_SectionMarkUnset, then no break will be inserted.
func (_fbeee ParagraphProperties) AddSection(t _db.ST_SectionMark) Section {
	_fbeee._aedc.SectPr = _db.NewCT_SectPr()
	if t != _db.ST_SectionMarkUnset {
		_fbeee._aedc.SectPr.Type = _db.NewCT_SectType()
		_fbeee._aedc.SectPr.Type.ValAttr = t
	}
	return Section{_fbeee._efedg, _fbeee._aedc.SectPr}
}

// Style return the table style.
func (_bdbec Table) Style() string {
	if _bdbec._cgffe.TblPr != nil && _bdbec._cgffe.TblPr.TblStyle != nil {
		return _bdbec._cgffe.TblPr.TblStyle.ValAttr
	}
	return ""
}
func (_aafg *chart) RelId() string { return _aafg._dfd }

// SetAlignment sets the paragraph alignment
func (_dbfb NumberingLevel) SetAlignment(j _db.ST_Jc) {
	if j == _db.ST_JcUnset {
		_dbfb._cfed.LvlJc = nil
	} else {
		_dbfb._cfed.LvlJc = _db.NewCT_Jc()
		_dbfb._cfed.LvlJc.ValAttr = j
	}
}

// AddBreak adds a line break to a run.
func (_ggga Run) AddBreak() { _agge := _ggga.newIC(); _agge.Br = _db.NewCT_Br() }

// SetLeftIndent controls the left indent of the paragraph.
func (_bbgg ParagraphStyleProperties) SetLeftIndent(m _bc.Distance) {
	if _bbgg._fagf.Ind == nil {
		_bbgg._fagf.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_bbgg._fagf.Ind.LeftAttr = nil
	} else {
		_bbgg._fagf.Ind.LeftAttr = &_db.ST_SignedTwipsMeasure{}
		_bbgg._fagf.Ind.LeftAttr.Int64 = _b.Int64(int64(m / _bc.Twips))
	}
}

var _fbgb = false

func _beb(_aaf *_db.CT_TblWidth, _aba float64) {
	_aaf.TypeAttr = _db.ST_TblWidthPct
	_aaf.WAttr = &_db.ST_MeasurementOrPercent{}
	_aaf.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_aaf.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(_aba * 50))
}

// SetPictureWashout set washout to watermark picture.
func (_abcfg *WatermarkPicture) SetPictureWashout(isWashout bool) {
	if _abcfg._febge != nil {
		_abbce := _abcfg._febge.EG_ShapeElements
		if len(_abbce) > 0 && _abbce[0].Imagedata != nil {
			if isWashout {
				_efgga := "\u0031\u0039\u0036\u0036\u0031\u0066"
				_eeaag := "\u0032\u0032\u0039\u0033\u0038\u0066"
				_abbce[0].Imagedata.GainAttr = &_efgga
				_abbce[0].Imagedata.BlacklevelAttr = &_eeaag
			}
		}
	}
}

// X return element of Node as interface, can be either *Paragraph, *Table and Run.
func (_cabcc *Node) X() interface{} { return _cabcc._bdcf }

// HasEndnotes returns a bool based on the presence or abscence of endnotes within
// the document.
func (_dcge *Document) HasEndnotes() bool { return _dcge._acdc != nil }

// SetTextWrapSquare sets the text wrap to square with a given wrap type.
func (_eda AnchoredDrawing) SetTextWrapSquare(t _db.WdST_WrapText) {
	_eda._cg.Choice = &_db.WdEG_WrapTypeChoice{}
	_eda._cg.Choice.WrapSquare = _db.NewWdCT_WrapSquare()
	_eda._cg.Choice.WrapSquare.WrapTextAttr = t
}
func _egfc(_defb *Document, _ccgf []*_db.CT_P, _eded *TableInfo, _ebagc *DrawingInfo) []Node {
	_fdg := []Node{}
	for _, _cabee := range _ccgf {
		_aegdf := Paragraph{_defb, _cabee}
		_ffbf := Node{_cffcb: _defb, _bdcf: &_aegdf}
		if _dccgf, _fbbg := _defb.Styles.SearchStyleById(_aegdf.Style()); _fbbg {
			_ffbf.Style = _dccgf
		}
		for _, _aaca := range _aegdf.Runs() {
			_ffbf.Children = append(_ffbf.Children, Node{_cffcb: _defb, _bdcf: _aaca, AnchoredDrawings: _aaca.DrawingAnchored(), InlineDrawings: _aaca.DrawingInline()})
		}
		_fdg = append(_fdg, _ffbf)
	}
	return _fdg
}

// GetTargetByRelId returns a target path with the associated relation ID in the
// document.
func (_ecdg *Document) GetTargetByRelId(idAttr string) string {
	return _ecdg._dfc.GetTargetByRelId(idAttr)
}
func (_afaag *WatermarkText) findNode(_adade *_b.XSDAny, _dcbd string) *_b.XSDAny {
	for _, _dcfce := range _adade.Nodes {
		if _dcfce.XMLName.Local == _dcbd {
			return _dcfce
		}
	}
	return nil
}

// NewTableWidth returns a newly intialized TableWidth
func NewTableWidth() TableWidth { return TableWidth{_db.NewCT_TblWidth()} }
func _aadb(_dfbd string) mergeFieldInfo {
	_dcbg := []string{}
	_ggeef := _fed.Buffer{}
	_dgfa := -1
	for _afae, _badgf := range _dfbd {
		switch _badgf {
		case ' ':
			if _ggeef.Len() != 0 {
				_dcbg = append(_dcbg, _ggeef.String())
			}
			_ggeef.Reset()
		case '"':
			if _dgfa != -1 {
				_dcbg = append(_dcbg, _dfbd[_dgfa+1:_afae])
				_dgfa = -1
			} else {
				_dgfa = _afae
			}
		default:
			_ggeef.WriteRune(_badgf)
		}
	}
	if _ggeef.Len() != 0 {
		_dcbg = append(_dcbg, _ggeef.String())
	}
	_ffee := mergeFieldInfo{}
	for _bdfd := 0; _bdfd < len(_dcbg)-1; _bdfd++ {
		_fgec := _dcbg[_bdfd]
		switch _fgec {
		case "\u004d\u0045\u0052\u0047\u0045\u0046\u0049\u0045\u004c\u0044":
			_ffee._gcfd = _dcbg[_bdfd+1]
			_bdfd++
		case "\u005c\u0066":
			_ffee._fggfa = _dcbg[_bdfd+1]
			_bdfd++
		case "\u005c\u0062":
			_ffee._bedf = _dcbg[_bdfd+1]
			_bdfd++
		case "\u005c\u002a":
			switch _dcbg[_bdfd+1] {
			case "\u0055\u0070\u0070e\u0072":
				_ffee._adedf = true
			case "\u004c\u006f\u0077e\u0072":
				_ffee._edece = true
			case "\u0043\u0061\u0070\u0073":
				_ffee._afbb = true
			case "\u0046\u0069\u0072\u0073\u0074\u0043\u0061\u0070":
				_ffee._gdcee = true
			}
			_bdfd++
		}
	}
	return _ffee
}

// TableBorders allows manipulation of borders on a table.
type TableBorders struct{ _fcdb *_db.CT_TblBorders }

// SetLeftPct sets the cell left margin
func (_aac CellMargins) SetLeftPct(pct float64) {
	_aac._caff.Left = _db.NewCT_TblWidth()
	_beb(_aac._caff.Left, pct)
}

// Borders allows controlling individual cell borders.
func (_cf CellProperties) Borders() CellBorders {
	if _cf._beea.TcBorders == nil {
		_cf._beea.TcBorders = _db.NewCT_TcBorders()
	}
	return CellBorders{_cf._beea.TcBorders}
}

// TableWidth controls width values in table settings.
type TableWidth struct{ _bagbe *_db.CT_TblWidth }

// SetAlignment controls the paragraph alignment
func (_daafg ParagraphStyleProperties) SetAlignment(align _db.ST_Jc) {
	if align == _db.ST_JcUnset {
		_daafg._fagf.Jc = nil
	} else {
		_daafg._fagf.Jc = _db.NewCT_Jc()
		_daafg._fagf.Jc.ValAttr = align
	}
}

// SetStyle sets the table style name.
func (_faecb TableProperties) SetStyle(name string) {
	if name == "" {
		_faecb._faec.TblStyle = nil
	} else {
		_faecb._faec.TblStyle = _db.NewCT_String()
		_faecb._faec.TblStyle.ValAttr = name
	}
}

// SetLeft sets the cell left margin
func (_caa CellMargins) SetLeft(d _bc.Distance) {
	_caa._caff.Left = _db.NewCT_TblWidth()
	_eeb(_caa._caff.Left, d)
}

// GetText returns text in the watermark.
func (_ebac *WatermarkText) GetText() string {
	_abga := _ebac.getShape()
	if _ebac._bgdbe != nil {
		_ffced := _ebac._bgdbe.EG_ShapeElements
		if len(_ffced) > 0 && _ffced[0].Textpath != nil {
			return *_ffced[0].Textpath.StringAttr
		}
	} else {
		_fbfbg := _ebac.findNode(_abga, "\u0074\u0065\u0078\u0074\u0070\u0061\u0074\u0068")
		for _, _daba := range _fbfbg.Attrs {
			if _daba.Name.Local == "\u0073\u0074\u0072\u0069\u006e\u0067" {
				return _daba.Value
			}
		}
	}
	return ""
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (_egfec ParagraphProperties) Style() string {
	if _egfec._aedc.PStyle != nil {
		return _egfec._aedc.PStyle.ValAttr
	}
	return ""
}
func (_egd *Document) InsertTableAfter(relativeTo Paragraph) Table {
	return _egd.insertTable(relativeTo, false)
}

// DrawingInline return a slice of InlineDrawings.
func (_fccbc Run) DrawingInline() []InlineDrawing {
	_ccga := []InlineDrawing{}
	for _, _abac := range _fccbc._afec.EG_RunInnerContent {
		if _abac.Drawing == nil {
			continue
		}
		for _, _dfgd := range _abac.Drawing.Inline {
			_ccga = append(_ccga, InlineDrawing{_fccbc._gacf, _dfgd})
		}
	}
	return _ccga
}
func _eefc(_bade []*_db.CT_P, _aebe *TableInfo, _bafge *DrawingInfo) []TextItem {
	_cbcd := []TextItem{}
	for _, _bgfa := range _bade {
		_cbcd = append(_cbcd, _ffgfd(_bgfa, nil, _aebe, _bafge, _bgfa.EG_PContent)...)
	}
	return _cbcd
}

// Cells returns the cells defined in the table.
func (_cdbf Row) Cells() []Cell {
	_cdbb := []Cell{}
	for _, _cacac := range _cdbf._bageg.EG_ContentCellContent {
		for _, _abfg := range _cacac.Tc {
			_cdbb = append(_cdbb, Cell{_cdbf._eefba, _abfg})
		}
		if _cacac.Sdt != nil && _cacac.Sdt.SdtContent != nil {
			for _, _dgceg := range _cacac.Sdt.SdtContent.Tc {
				_cdbb = append(_cdbb, Cell{_cdbf._eefba, _dgceg})
			}
		}
	}
	return _cdbb
}

// ParagraphProperties returns the paragraph style properties.
func (_cdcd Style) ParagraphProperties() ParagraphStyleProperties {
	if _cdcd._eedcd.PPr == nil {
		_cdcd._eedcd.PPr = _db.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_cdcd._eedcd.PPr}
}

// SetRight sets the right border to a specified type, color and thickness.
func (_aeb CellBorders) SetRight(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_aeb._bee.Right = _db.NewCT_Border()
	_ebca(_aeb._bee.Right, t, c, thickness)
}

// Properties returns the paragraph properties.
func (_fcdf Paragraph) Properties() ParagraphProperties {
	_fcdf.ensurePPr()
	return ParagraphProperties{_fcdf._cfge, _fcdf._dcfbd.PPr}
}

// SetLineSpacing controls the line spacing of the paragraph.
func (_fedd ParagraphStyleProperties) SetLineSpacing(m _bc.Distance, rule _db.ST_LineSpacingRule) {
	if _fedd._fagf.Spacing == nil {
		_fedd._fagf.Spacing = _db.NewCT_Spacing()
	}
	if rule == _db.ST_LineSpacingRuleUnset {
		_fedd._fagf.Spacing.LineRuleAttr = _db.ST_LineSpacingRuleUnset
		_fedd._fagf.Spacing.LineAttr = nil
	} else {
		_fedd._fagf.Spacing.LineRuleAttr = rule
		_fedd._fagf.Spacing.LineAttr = &_db.ST_SignedTwipsMeasure{}
		_fedd._fagf.Spacing.LineAttr.Int64 = _b.Int64(int64(m / _bc.Twips))
	}
}

// SetASCIITheme sets the font ASCII Theme.
func (_cdeg Fonts) SetASCIITheme(t _db.ST_Theme) { _cdeg._bdge.AsciiThemeAttr = t }

// UnderlineColor returns the hex color value of run underline.
func (_adgcg RunProperties) UnderlineColor() string {
	if _cfee := _adgcg._cadb.U; _cfee != nil {
		_cfefd := _cfee.ColorAttr
		if _cfefd != nil && _cfefd.ST_HexColorRGB != nil {
			return *_cfefd.ST_HexColorRGB
		}
	}
	return ""
}

// SetLayoutInCell sets the layoutInCell attribute of anchor.
func (_gfgb AnchoredDrawing) SetLayoutInCell(val bool) { _gfgb._cg.LayoutInCellAttr = val }
func (_gdeb Endnote) content() []*_db.EG_ContentBlockContent {
	var _fafgb []*_db.EG_ContentBlockContent
	for _, _efab := range _gdeb._eeabdb.EG_BlockLevelElts {
		_fafgb = append(_fafgb, _efab.EG_ContentBlockContent...)
	}
	return _fafgb
}
func _fffb() *_eb.OfcLock {
	_cccfe := _eb.NewOfcLock()
	_cccfe.ExtAttr = _eb.ST_ExtEdit
	_cccfe.TextAttr = _ca.ST_TrueFalseTrue
	_cccfe.ShapetypeAttr = _ca.ST_TrueFalseTrue
	return _cccfe
}

// RemoveParagraph removes a paragraph from the footnote.
func (_addgf Footnote) RemoveParagraph(p Paragraph) {
	for _, _dad := range _addgf.content() {
		for _abae, _bfde := range _dad.P {
			if _bfde == p._dcfbd {
				copy(_dad.P[_abae:], _dad.P[_abae+1:])
				_dad.P = _dad.P[0 : len(_dad.P)-1]
				return
			}
		}
	}
}

// Outline returns true if run outline is on.
func (_dbdc RunProperties) Outline() bool { return _ebdd(_dbdc._cadb.Outline) }

// SetValue sets the width value.
func (_bdce TableWidth) SetValue(m _bc.Distance) {
	_bdce._bagbe.WAttr = &_db.ST_MeasurementOrPercent{}
	_bdce._bagbe.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_bdce._bagbe.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(m / _bc.Twips))
	_bdce._bagbe.TypeAttr = _db.ST_TblWidthDxa
}
func _cegge() *_eb.OfcLock {
	_deed := _eb.NewOfcLock()
	_deed.ExtAttr = _eb.ST_ExtEdit
	_deed.AspectratioAttr = _ca.ST_TrueFalseTrue
	return _deed
}

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (_fadcg TableProperties) SetCellSpacingPercent(pct float64) {
	_fadcg._faec.TblCellSpacing = _db.NewCT_TblWidth()
	_fadcg._faec.TblCellSpacing.TypeAttr = _db.ST_TblWidthPct
	_fadcg._faec.TblCellSpacing.WAttr = &_db.ST_MeasurementOrPercent{}
	_fadcg._faec.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_fadcg._faec.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(pct * 50))
}

// Properties returns the numbering level paragraph properties.
func (_afea NumberingLevel) Properties() ParagraphStyleProperties {
	if _afea._cfed.PPr == nil {
		_afea._cfed.PPr = _db.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_afea._cfed.PPr}
}

// SetFormat sets the numbering format.
func (_cafbc NumberingLevel) SetFormat(f _db.ST_NumberFormat) {
	if _cafbc._cfed.NumFmt == nil {
		_cafbc._cfed.NumFmt = _db.NewCT_NumFmt()
	}
	_cafbc._cfed.NumFmt.ValAttr = f
}

// StyleID returns the style ID.
func (_defc Style) StyleID() string {
	if _defc._eedcd.StyleIdAttr == nil {
		return ""
	}
	return *_defc._eedcd.StyleIdAttr
}

// AddRun adds a run of text to a hyperlink. This is the text that will be linked.
func (_fafe HyperLink) AddRun() Run {
	_abef := _db.NewEG_ContentRunContent()
	_fafe._cfcf.EG_ContentRunContent = append(_fafe._cfcf.EG_ContentRunContent, _abef)
	_cadc := _db.NewCT_R()
	_abef.R = _cadc
	return Run{_fafe._deac, _cadc}
}

// InsertRowBefore inserts a row before another row
func (_cbabb Table) InsertRowBefore(r Row) Row {
	for _bbfef, _gdee := range _cbabb._cgffe.EG_ContentRowContent {
		if len(_gdee.Tr) > 0 && r.X() == _gdee.Tr[0] {
			_abfdg := _db.NewEG_ContentRowContent()
			_cbabb._cgffe.EG_ContentRowContent = append(_cbabb._cgffe.EG_ContentRowContent, nil)
			copy(_cbabb._cgffe.EG_ContentRowContent[_bbfef+1:], _cbabb._cgffe.EG_ContentRowContent[_bbfef:])
			_cbabb._cgffe.EG_ContentRowContent[_bbfef] = _abfdg
			_dggabc := _db.NewCT_Row()
			_abfdg.Tr = append(_abfdg.Tr, _dggabc)
			return Row{_cbabb._ffcf, _dggabc}
		}
	}
	return _cbabb.AddRow()
}

// Header is a header for a document section.
type Header struct {
	_bgac  *Document
	_ecead *_db.Hdr
}

// AddRun adds a run to a paragraph.
func (_becg Paragraph) AddRun() Run {
	_ccfc := _db.NewEG_PContent()
	_becg._dcfbd.EG_PContent = append(_becg._dcfbd.EG_PContent, _ccfc)
	_aage := _db.NewEG_ContentRunContent()
	_ccfc.EG_ContentRunContent = append(_ccfc.EG_ContentRunContent, _aage)
	_bdec := _db.NewCT_R()
	_aage.R = _bdec
	return Run{_becg._cfge, _bdec}
}

// ReplaceText replace text inside node.
func (_ebdg *Nodes) ReplaceText(oldText, newText string) {
	for _, _bdff := range _ebdg._dcdf {
		_bdff.ReplaceText(oldText, newText)
	}
}
func _bdgfg(_dccd *Document, _cgbg []*_db.EG_ContentBlockContent, _bfcg *TableInfo) []Node {
	_ageg := []Node{}
	for _, _eaegd := range _cgbg {
		if _ccbda := _eaegd.Sdt; _ccbda != nil {
			if _eeae := _ccbda.SdtContent; _eeae != nil {
				_ageg = append(_ageg, _egfc(_dccd, _eeae.P, _bfcg, nil)...)
			}
		}
		_ageg = append(_ageg, _egfc(_dccd, _eaegd.P, _bfcg, nil)...)
		for _, _gfdg := range _eaegd.Tbl {
			_cdgea := Table{_dccd, _gfdg}
			_dggg, _ := _dccd.Styles.SearchStyleById(_cdgea.Style())
			_facf := []Node{}
			for _abgca, _gecb := range _gfdg.EG_ContentRowContent {
				for _, _edce := range _gecb.Tr {
					for _gdda, _daaf := range _edce.EG_ContentCellContent {
						for _, _efff := range _daaf.Tc {
							_ddafb := &TableInfo{Table: _gfdg, Row: _edce, Cell: _efff, RowIndex: _abgca, ColIndex: _gdda}
							for _, _gbd := range _efff.EG_BlockLevelElts {
								_facf = append(_facf, _bdgfg(_dccd, _gbd.EG_ContentBlockContent, _ddafb)...)
							}
						}
					}
				}
			}
			_ageg = append(_ageg, Node{_cffcb: _dccd, _bdcf: &_cdgea, Style: _dggg, Children: _facf})
		}
	}
	return _ageg
}

// X returns the inner wrapped XML type.
func (_gfcad Footnote) X() *_db.CT_FtnEdn { return _gfcad._cefg }

// SetMultiLevelType sets the multilevel type.
func (_geeag NumberingDefinition) SetMultiLevelType(t _db.ST_MultiLevelType) {
	if t == _db.ST_MultiLevelTypeUnset {
		_geeag._bbgd.MultiLevelType = nil
	} else {
		_geeag._bbgd.MultiLevelType = _db.NewCT_MultiLevelType()
		_geeag._bbgd.MultiLevelType.ValAttr = t
	}
}

// SetTopPct sets the cell top margin
func (_gfb CellMargins) SetTopPct(pct float64) {
	_gfb._caff.Top = _db.NewCT_TblWidth()
	_beb(_gfb._caff.Top, pct)
}

// InsertRunBefore inserts a run in the paragraph before the relative run.
func (_ccgdg Paragraph) InsertRunBefore(relativeTo Run) Run {
	return _ccgdg.insertRun(relativeTo, true)
}
func _ecdeg() *_eb.Path {
	_ebad := _eb.NewPath()
	_ebad.TextpathokAttr = _ca.ST_TrueFalseTrue
	_ebad.ConnecttypeAttr = _eb.OfcST_ConnectTypeCustom
	_aggge := "\u0040\u0039\u002c0;\u0040\u0031\u0030\u002c\u0031\u0030\u0038\u0030\u0030;\u00401\u0031,\u00321\u0036\u0030\u0030\u003b\u0040\u0031\u0032\u002c\u0031\u0030\u0038\u0030\u0030"
	_ebad.ConnectlocsAttr = &_aggge
	_bbabc := "\u0032\u0037\u0030,\u0031\u0038\u0030\u002c\u0039\u0030\u002c\u0030"
	_ebad.ConnectanglesAttr = &_bbabc
	return _ebad
}

// X returns the inner wrapped XML type.
func (_cb Bookmark) X() *_db.CT_Bookmark { return _cb._ac }

// GetShapeStyle returns string style of the shape in watermark and format it to ShapeStyle.
func (_acdea *WatermarkPicture) GetShapeStyle() _ge.ShapeStyle {
	if _acdea._febge != nil && _acdea._febge.StyleAttr != nil {
		return _ge.NewShapeStyle(*_acdea._febge.StyleAttr)
	}
	return _ge.NewShapeStyle("")
}

// X returns the inner wrapped XML type.
func (_ebfef TableLook) X() *_db.CT_TblLook { return _ebfef._febce }

// Bold returns true if paragraph font is bold.
func (_gdebf ParagraphProperties) Bold() bool {
	_cdcfea := _gdebf._aedc.RPr
	return _ebdd(_cdcfea.B) || _ebdd(_cdcfea.BCs)
}

// SetTextWrapTopAndBottom sets the text wrap to top and bottom.
func (_ced AnchoredDrawing) SetTextWrapTopAndBottom() {
	_ced._cg.Choice = &_db.WdEG_WrapTypeChoice{}
	_ced._cg.Choice.WrapTopAndBottom = _db.NewWdCT_WrapTopBottom()
	_ced._cg.LayoutInCellAttr = true
	_ced._cg.AllowOverlapAttr = true
}

// X returns the inner wrapped XML type.
func (_abbcf Settings) X() *_db.Settings { return _abbcf._ccgc }

// RemoveMailMerge removes any mail merge settings
func (_dcca Settings) RemoveMailMerge() { _dcca._ccgc.MailMerge = nil }

// SetTextWrapInFrontOfText sets the text wrap to in front of text.
func (_ffe AnchoredDrawing) SetTextWrapInFrontOfText() {
	_ffe._cg.Choice = &_db.WdEG_WrapTypeChoice{}
	_ffe._cg.Choice.WrapNone = _db.NewWdCT_WrapNone()
	_ffe._cg.BehindDocAttr = false
	_ffe._cg.LayoutInCellAttr = true
	_ffe._cg.AllowOverlapAttr = true
}

// IsItalic returns true if the run has been set to italics.
func (_effaeb RunProperties) IsItalic() bool { return _effaeb.ItalicValue() == OnOffValueOn }

// X returns the inner wrapped XML type.
func (_ddcd Row) X() *_db.CT_Row { return _ddcd._bageg }

// SetDoubleStrikeThrough sets the run to double strike-through.
func (_babg RunProperties) SetDoubleStrikeThrough(b bool) {
	if !b {
		_babg._cadb.Dstrike = nil
	} else {
		_babg._cadb.Dstrike = _db.NewCT_OnOff()
	}
}

// InsertRunAfter inserts a run in the paragraph after the relative run.
func (_efae Paragraph) InsertRunAfter(relativeTo Run) Run { return _efae.insertRun(relativeTo, false) }

// IgnoreSpaceBetweenParagraphOfSameStyle sets contextual spacing.
func (_eaae Paragraph) IgnoreSpaceBetweenParagraphOfSameStyle() {
	_eaae.ensurePPr()
	if _eaae._dcfbd.PPr.ContextualSpacing == nil {
		_eaae._dcfbd.PPr.ContextualSpacing = _db.NewCT_OnOff()
	}
	_eaae._dcfbd.PPr.ContextualSpacing.ValAttr = &_ca.ST_OnOff{ST_OnOff1: _ca.ST_OnOff1On}
}

// Emboss returns true if paragraph emboss is on.
func (_bdac ParagraphProperties) Emboss() bool { return _ebdd(_bdac._aedc.RPr.Emboss) }

// Endnote is an individual endnote reference within the document.
type Endnote struct {
	_efb    *Document
	_eeabdb *_db.CT_FtnEdn
}

// SetShading controls the cell shading.
func (_bbg CellProperties) SetShading(shd _db.ST_Shd, foreground, fill _bd.Color) {
	if shd == _db.ST_ShdUnset {
		_bbg._beea.Shd = nil
	} else {
		_bbg._beea.Shd = _db.NewCT_Shd()
		_bbg._beea.Shd.ValAttr = shd
		_bbg._beea.Shd.ColorAttr = &_db.ST_HexColor{}
		if foreground.IsAuto() {
			_bbg._beea.Shd.ColorAttr.ST_HexColorAuto = _db.ST_HexColorAutoAuto
		} else {
			_bbg._beea.Shd.ColorAttr.ST_HexColorRGB = foreground.AsRGBString()
		}
		_bbg._beea.Shd.FillAttr = &_db.ST_HexColor{}
		if fill.IsAuto() {
			_bbg._beea.Shd.FillAttr.ST_HexColorAuto = _db.ST_HexColorAutoAuto
		} else {
			_bbg._beea.Shd.FillAttr.ST_HexColorRGB = fill.AsRGBString()
		}
	}
}

// SetBetween sets the between border to a specified type, color and thickness between paragraph.
func (_bacea ParagraphBorders) SetBetween(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_bacea._afaaa.Between = _db.NewCT_Border()
	_gbge(_bacea._afaaa.Between, t, c, thickness)
}

// SetColor sets a specific color or auto.
func (_fac Color) SetColor(v _bd.Color) {
	if v.IsAuto() {
		_fac._fcd.ValAttr.ST_HexColorAuto = _db.ST_HexColorAutoAuto
		_fac._fcd.ValAttr.ST_HexColorRGB = nil
	} else {
		_fac._fcd.ValAttr.ST_HexColorAuto = _db.ST_HexColorAutoUnset
		_fac._fcd.ValAttr.ST_HexColorRGB = v.AsRGBString()
	}
}

// CharacterSpacingMeasure returns paragraph characters spacing with its measure which can be mm, cm, in, pt, pc or pi.
func (_cdgdc ParagraphProperties) CharacterSpacingMeasure() string {
	if _fgcff := _cdgdc._aedc.RPr.Spacing; _fgcff != nil {
		_gddac := _fgcff.ValAttr
		if _gddac.ST_UniversalMeasure != nil {
			return *_gddac.ST_UniversalMeasure
		}
	}
	return ""
}

// SetSize sets the font size for a run.
func (_dbgg RunProperties) SetSize(size _bc.Distance) {
	_dbgg._cadb.Sz = _db.NewCT_HpsMeasure()
	_dbgg._cadb.Sz.ValAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(size / _bc.HalfPoint))
	_dbgg._cadb.SzCs = _db.NewCT_HpsMeasure()
	_dbgg._cadb.SzCs.ValAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(size / _bc.HalfPoint))
}

// SetSpacing sets the spacing that comes before and after the paragraph.
// Deprecated: See Spacing() instead which allows finer control.
func (_dgad ParagraphProperties) SetSpacing(before, after _bc.Distance) {
	if _dgad._aedc.Spacing == nil {
		_dgad._aedc.Spacing = _db.NewCT_Spacing()
	}
	_dgad._aedc.Spacing.BeforeAttr = &_ca.ST_TwipsMeasure{}
	_dgad._aedc.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(before / _bc.Twips))
	_dgad._aedc.Spacing.AfterAttr = &_ca.ST_TwipsMeasure{}
	_dgad._aedc.Spacing.AfterAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(after / _bc.Twips))
}

// Paragraphs returns the paragraphs defined in a footer.
func (_aggb Footer) Paragraphs() []Paragraph {
	_fdcc := []Paragraph{}
	for _, _dbfg := range _aggb._eeed.EG_ContentBlockContent {
		for _, _cdag := range _dbfg.P {
			_fdcc = append(_fdcc, Paragraph{_aggb._bcge, _cdag})
		}
	}
	for _, _beca := range _aggb.Tables() {
		for _, _bfbd := range _beca.Rows() {
			for _, _ccad := range _bfbd.Cells() {
				_fdcc = append(_fdcc, _ccad.Paragraphs()...)
			}
		}
	}
	return _fdcc
}

// CharacterSpacingValue returns the value of run's characters spacing in twips (1/20 of point).
func (_bbfa RunProperties) CharacterSpacingValue() int64 {
	if _edbe := _bbfa._cadb.Spacing; _edbe != nil {
		_ebcbf := _edbe.ValAttr
		if _ebcbf.Int64 != nil {
			return *_ebcbf.Int64
		}
	}
	return int64(0)
}

// BoldValue returns the precise nature of the bold setting (unset, off or on).
func (_caccf RunProperties) BoldValue() OnOffValue { return _egeb(_caccf._cadb.B) }

// FindNodeByCondition return node based on condition function,
// if wholeElements is true, its will extract childs as next node elements.
func (_dcfb *Nodes) FindNodeByCondition(f func(_fbda *Node) bool, wholeElements bool) []Node {
	_babb := []Node{}
	for _, _bged := range _dcfb._dcdf {
		if f(&_bged) {
			_babb = append(_babb, _bged)
		}
		if wholeElements {
			_dddaff := Nodes{_dcdf: _bged.Children}
			_babb = append(_babb, _dddaff.FindNodeByCondition(f, wholeElements)...)
		}
	}
	return _babb
}

// SetItalic sets the run to italic.
func (_gbec RunProperties) SetItalic(b bool) {
	if !b {
		_gbec._cadb.I = nil
		_gbec._cadb.ICs = nil
	} else {
		_gbec._cadb.I = _db.NewCT_OnOff()
		_gbec._cadb.ICs = _db.NewCT_OnOff()
	}
}

// ReplaceTextByRegexp replace text inside node using regexp.
func (_ccebg *Nodes) ReplaceTextByRegexp(expr *_ff.Regexp, newText string) {
	for _, _acg := range _ccebg._dcdf {
		_acg.ReplaceTextByRegexp(expr, newText)
	}
}

// BodySection returns the default body section used for all preceding
// paragraphs until the previous Section. If there is no previous sections, the
// body section applies to the entire document.
func (_bff *Document) BodySection() Section {
	if _bff._adcb.Body.SectPr == nil {
		_bff._adcb.Body.SectPr = _db.NewCT_SectPr()
	}
	return Section{_bff, _bff._adcb.Body.SectPr}
}

// Close closes the document, removing any temporary files that might have been
// created when opening a document.
func (_fefd *Document) Close() error {
	if _fefd.TmpPath != "" {
		return _be.RemoveAll(_fefd.TmpPath)
	}
	return nil
}
func _cgfe() *_eb.Path {
	_dcfc := _eb.NewPath()
	_dcfc.ExtrusionokAttr = _ca.ST_TrueFalseTrue
	_dcfc.GradientshapeokAttr = _ca.ST_TrueFalseTrue
	_dcfc.ConnecttypeAttr = _eb.OfcST_ConnectTypeRect
	return _dcfc
}

// GetEffect returns the effect of the run.
func (_eddcd RunProperties) GetEffect() _db.ST_TextEffect {
	if _eddcd._cadb.Effect == nil {
		return _db.ST_TextEffectUnset
	}
	return _eddcd._cadb.Effect.ValAttr
}

// Clear removes all of the content from within a run.
func (_bacg Run) Clear() { _bacg._afec.EG_RunInnerContent = nil }

// SetYOffset sets the Y offset for an image relative to the origin.
func (_cd AnchoredDrawing) SetYOffset(y _bc.Distance) {
	_cd._cg.PositionV.Choice = &_db.WdCT_PosVChoice{}
	_cd._cg.PositionV.Choice.PosOffset = _b.Int32(int32(y / _bc.EMU))
}

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (_cbd CellBorders) SetInsideVertical(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_cbd._bee.InsideV = _db.NewCT_Border()
	_ebca(_cbd._bee.InsideV, t, c, thickness)
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (_fcfg TableStyleProperties) SetCellSpacingAuto() {
	_fcfg._dagce.TblCellSpacing = _db.NewCT_TblWidth()
	_fcfg._dagce.TblCellSpacing.TypeAttr = _db.ST_TblWidthAuto
}

// X returns the inner wrapped XML type.
func (_fbce Footer) X() *_db.Ftr { return _fbce._eeed }

// AddParagraph adds a new paragraph to the document body.
func (_efdd *Document) AddParagraph() Paragraph {
	_gee := _db.NewEG_BlockLevelElts()
	_efdd._adcb.Body.EG_BlockLevelElts = append(_efdd._adcb.Body.EG_BlockLevelElts, _gee)
	_febd := _db.NewEG_ContentBlockContent()
	_gee.EG_ContentBlockContent = append(_gee.EG_ContentBlockContent, _febd)
	_dfee := _db.NewCT_P()
	_febd.P = append(_febd.P, _dfee)
	return Paragraph{_efdd, _dfee}
}

// Bookmarks returns all of the bookmarks defined in the document.
func (_gfea Document) Bookmarks() []Bookmark {
	if _gfea._adcb.Body == nil {
		return nil
	}
	_cdgb := []Bookmark{}
	for _, _dcfd := range _gfea._adcb.Body.EG_BlockLevelElts {
		for _, _cdbc := range _dcfd.EG_ContentBlockContent {
			for _, _bcgb := range _efdfc(_cdbc) {
				_cdgb = append(_cdgb, _bcgb)
			}
		}
	}
	return _cdgb
}

// SetAll sets all of the borders to a given value.
func (_ceaba TableBorders) SetAll(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_ceaba.SetBottom(t, c, thickness)
	_ceaba.SetLeft(t, c, thickness)
	_ceaba.SetRight(t, c, thickness)
	_ceaba.SetTop(t, c, thickness)
	_ceaba.SetInsideHorizontal(t, c, thickness)
	_ceaba.SetInsideVertical(t, c, thickness)
}

// SetSemiHidden controls if the style is hidden in the UI.
func (_gega Style) SetSemiHidden(b bool) {
	if b {
		_gega._eedcd.SemiHidden = _db.NewCT_OnOff()
	} else {
		_gega._eedcd.SemiHidden = nil
	}
}

// SetInsideVertical sets the interior vertical borders to a specified type, color and thickness.
func (_cgaf TableBorders) SetInsideVertical(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_cgaf._fcdb.InsideV = _db.NewCT_Border()
	_ebca(_cgaf._fcdb.InsideV, t, c, thickness)
}

// SetLayout controls the table layout. wml.ST_TblLayoutTypeAutofit corresponds
// to "Automatically resize to fit contents" being checked, while
// wml.ST_TblLayoutTypeFixed corresponds to it being unchecked.
func (_efdg TableProperties) SetLayout(l _db.ST_TblLayoutType) {
	if l == _db.ST_TblLayoutTypeUnset || l == _db.ST_TblLayoutTypeAutofit {
		_efdg._faec.TblLayout = nil
	} else {
		_efdg._faec.TblLayout = _db.NewCT_TblLayoutType()
		_efdg._faec.TblLayout.TypeAttr = l
	}
}

// SetLeft sets the left border to a specified type, color and thickness.
func (_ecbb ParagraphBorders) SetLeft(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_ecbb._afaaa.Left = _db.NewCT_Border()
	_gbge(_ecbb._afaaa.Left, t, c, thickness)
}

// Styles is the document wide styles contained in styles.xml.
type Styles struct{ _cdgg *_db.Styles }

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (_afaec Paragraph) SetFirstLineIndent(m _bc.Distance) {
	_afaec.ensurePPr()
	_ebga := _afaec._dcfbd.PPr
	if _ebga.Ind == nil {
		_ebga.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_ebga.Ind.FirstLineAttr = nil
	} else {
		_ebga.Ind.FirstLineAttr = &_ca.ST_TwipsMeasure{}
		_ebga.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(m / _bc.Twips))
	}
}

// HasFootnotes returns a bool based on the presence or abscence of footnotes within
// the document.
func (_dbfc *Document) HasFootnotes() bool { return _dbfc._ffcef != nil }

// X returns the inner wrapped XML type.
func (_afgae TableProperties) X() *_db.CT_TblPr { return _afgae._faec }
func (_gfec *Document) validateBookmarks() error {
	_bagb := make(map[string]struct{})
	for _, _dege := range _gfec.Bookmarks() {
		if _, _fbad := _bagb[_dege.Name()]; _fbad {
			return _g.Errorf("d\u0075\u0070\u006c\u0069\u0063\u0061t\u0065\u0020\u0062\u006f\u006f\u006b\u006d\u0061\u0072k\u0020\u0025\u0073 \u0066o\u0075\u006e\u0064", _dege.Name())
		}
		_bagb[_dege.Name()] = struct{}{}
	}
	return nil
}

// SetOffset sets the offset of the image relative to the origin, which by
// default this is the top-left corner of the page. Offset is incompatible with
// SetAlignment, whichever is called last is applied.
func (_cae AnchoredDrawing) SetOffset(x, y _bc.Distance) { _cae.SetXOffset(x); _cae.SetYOffset(y) }

// AddWatermarkPicture adds new watermark picture to document.
func (_faag *Document) AddWatermarkPicture(imageRef _ffa.ImageRef) WatermarkPicture {
	var _ggfd []Header
	if _degf, _beagf := _faag.BodySection().GetHeader(_db.ST_HdrFtrDefault); _beagf {
		_ggfd = append(_ggfd, _degf)
	}
	if _cdbg, _beage := _faag.BodySection().GetHeader(_db.ST_HdrFtrEven); _beage {
		_ggfd = append(_ggfd, _cdbg)
	}
	if _efcf, _gdfd := _faag.BodySection().GetHeader(_db.ST_HdrFtrFirst); _gdfd {
		_ggfd = append(_ggfd, _efcf)
	}
	if len(_ggfd) < 1 {
		_eebf := _faag.AddHeader()
		_faag.BodySection().SetHeader(_eebf, _db.ST_HdrFtrDefault)
		_ggfd = append(_ggfd, _eebf)
	}
	var _dbbe error
	_afaa := NewWatermarkPicture()
	for _, _aeg := range _ggfd {
		imageRef, _dbbe = _aeg.AddImageRef(imageRef)
		if _dbbe != nil {
			return WatermarkPicture{}
		}
		_ecgd := _aeg.Paragraphs()
		if len(_ecgd) < 1 {
			_edec := _aeg.AddParagraph()
			_edec.AddRun().AddText("")
		}
		for _, _fcde := range _aeg.X().EG_ContentBlockContent {
			for _, _ccecb := range _fcde.P {
				for _, _eeabd := range _ccecb.EG_PContent {
					for _, _fadg := range _eeabd.EG_ContentRunContent {
						if _fadg.R == nil {
							continue
						}
						for _, _bacc := range _fadg.R.EG_RunInnerContent {
							_bacc.Pict = _afaa._dgdd
							break
						}
					}
				}
			}
		}
	}
	_afaa.SetPicture(imageRef)
	return _afaa
}

// X returns the inner wrapped XML type.
func (_acad Endnote) X() *_db.CT_FtnEdn { return _acad._eeabdb }

// CellBorders are the borders for an individual
type CellBorders struct{ _bee *_db.CT_TcBorders }

// X returns the inner wrapped XML type.
func (_gfgaa Style) X() *_db.CT_Style { return _gfgaa._eedcd }

// Strike returns true if paragraph is striked.
func (_dgffd ParagraphProperties) Strike() bool { return _ebdd(_dgffd._aedc.RPr.Strike) }

// InsertRowAfter inserts a row after another row
func (_efge Table) InsertRowAfter(r Row) Row {
	for _dabg, _eacg := range _efge._cgffe.EG_ContentRowContent {
		if len(_eacg.Tr) > 0 && r.X() == _eacg.Tr[0] {
			_eefaf := _db.NewEG_ContentRowContent()
			if len(_efge._cgffe.EG_ContentRowContent) < _dabg+2 {
				return _efge.AddRow()
			}
			_efge._cgffe.EG_ContentRowContent = append(_efge._cgffe.EG_ContentRowContent, nil)
			copy(_efge._cgffe.EG_ContentRowContent[_dabg+2:], _efge._cgffe.EG_ContentRowContent[_dabg+1:])
			_efge._cgffe.EG_ContentRowContent[_dabg+1] = _eefaf
			_gdcg := _db.NewCT_Row()
			_eefaf.Tr = append(_eefaf.Tr, _gdcg)
			return Row{_efge._ffcf, _gdcg}
		}
	}
	return _efge.AddRow()
}

// SetBefore sets the spacing that comes before the paragraph.
func (_fbcc ParagraphSpacing) SetBefore(before _bc.Distance) {
	_fbcc._ffeb.BeforeAttr = &_ca.ST_TwipsMeasure{}
	_fbcc._ffeb.BeforeAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(before / _bc.Twips))
}

// Outline returns true if paragraph outline is on.
func (_bfbb ParagraphProperties) Outline() bool { return _ebdd(_bfbb._aedc.RPr.Outline) }

// SetTop sets the cell top margin
func (_eeg CellMargins) SetTop(d _bc.Distance) {
	_eeg._caff.Top = _db.NewCT_TblWidth()
	_eeb(_eeg._caff.Top, d)
}

// AddFooter creates a Footer associated with the document, but doesn't add it
// to the document for display.
func (_dbc *Document) AddFooter() Footer {
	_bec := _db.NewFtr()
	_dbc._abc = append(_dbc._abc, _bec)
	_abcb := _g.Sprintf("\u0066\u006f\u006ft\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", len(_dbc._abc))
	_dbc._dfc.AddRelationship(_abcb, _b.FooterType)
	_dbc.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_abcb, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0066\u006f\u006f\u0074e\u0072\u002b\u0078\u006d\u006c")
	_dbc._adg = append(_dbc._adg, _ffa.NewRelationships())
	return Footer{_dbc, _bec}
}

// SizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_gfabg RunProperties) SizeMeasure() string {
	if _gbdd := _gfabg._cadb.Sz; _gbdd != nil {
		_eaeag := _gbdd.ValAttr
		if _eaeag.ST_PositiveUniversalMeasure != nil {
			return *_eaeag.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

var _dfcd = [...]uint8{0, 20, 37, 58, 79}

// SetLastRow controls the conditional formatting for the last row in a table.
// This is called the 'Total' row within Word.
func (_bbdcb TableLook) SetLastRow(on bool) {
	if !on {
		_bbdcb._febce.LastRowAttr = &_ca.ST_OnOff{}
		_bbdcb._febce.LastRowAttr.ST_OnOff1 = _ca.ST_OnOff1Off
	} else {
		_bbdcb._febce.LastRowAttr = &_ca.ST_OnOff{}
		_bbdcb._febce.LastRowAttr.ST_OnOff1 = _ca.ST_OnOff1On
	}
}

// AddBookmark adds a bookmark to a document that can then be used from a hyperlink. Name is a document
// unique name that identifies the bookmark so it can be referenced from hyperlinks.
func (_gadca Paragraph) AddBookmark(name string) Bookmark {
	_gcada := _db.NewEG_PContent()
	_gbda := _db.NewEG_ContentRunContent()
	_gcada.EG_ContentRunContent = append(_gcada.EG_ContentRunContent, _gbda)
	_eceab := _db.NewEG_RunLevelElts()
	_gbda.EG_RunLevelElts = append(_gbda.EG_RunLevelElts, _eceab)
	_efgg := _db.NewEG_RangeMarkupElements()
	_bffa := _db.NewCT_Bookmark()
	_efgg.BookmarkStart = _bffa
	_eceab.EG_RangeMarkupElements = append(_eceab.EG_RangeMarkupElements, _efgg)
	_efgg = _db.NewEG_RangeMarkupElements()
	_efgg.BookmarkEnd = _db.NewCT_MarkupRange()
	_eceab.EG_RangeMarkupElements = append(_eceab.EG_RangeMarkupElements, _efgg)
	_gadca._dcfbd.EG_PContent = append(_gadca._dcfbd.EG_PContent, _gcada)
	_efddd := Bookmark{_bffa}
	_efddd.SetName(name)
	return _efddd
}

// SetStyle sets the font size.
func (_fegd RunProperties) SetStyle(style string) {
	if style == "" {
		_fegd._cadb.RStyle = nil
	} else {
		_fegd._cadb.RStyle = _db.NewCT_String()
		_fegd._cadb.RStyle.ValAttr = style
	}
}

// DoubleStrike returns true if run is double striked.
func (_bfgae RunProperties) DoubleStrike() bool { return _ebdd(_bfgae._cadb.Dstrike) }

// Style returns the style for a paragraph, or an empty string if it is unset.
func (_deffa Paragraph) Style() string {
	if _deffa._dcfbd.PPr != nil && _deffa._dcfbd.PPr.PStyle != nil {
		return _deffa._dcfbd.PPr.PStyle.ValAttr
	}
	return ""
}

// SetRightIndent controls right indent of paragraph.
func (_dbfa Paragraph) SetRightIndent(m _bc.Distance) {
	_dbfa.ensurePPr()
	_febdc := _dbfa._dcfbd.PPr
	if _febdc.Ind == nil {
		_febdc.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_febdc.Ind.RightAttr = nil
	} else {
		_febdc.Ind.RightAttr = &_db.ST_SignedTwipsMeasure{}
		_febdc.Ind.RightAttr.Int64 = _b.Int64(int64(m / _bc.Twips))
	}
}

// GetNumberingLevelByIds returns a NumberingLevel by its NumId and LevelId attributes
// or an empty one if not found.
func (_gedf *Document) GetNumberingLevelByIds(numId, levelId int64) NumberingLevel {
	for _, _afbe := range _gedf.Numbering._gebaa.Num {
		if _afbe != nil && _afbe.NumIdAttr == numId {
			_bafa := _afbe.AbstractNumId.ValAttr
			for _, _agb := range _gedf.Numbering._gebaa.AbstractNum {
				if _agb.AbstractNumIdAttr == _bafa {
					if _agb.NumStyleLink != nil && len(_agb.Lvl) == 0 {
						if _aeaf, _gbac := _gedf.Styles.SearchStyleById(_agb.NumStyleLink.ValAttr); _gbac {
							if _aeaf.ParagraphProperties().NumId() > -1 {
								return _gedf.GetNumberingLevelByIds(_aeaf.ParagraphProperties().NumId(), levelId)
							}
						}
					}
					for _, _gdcb := range _agb.Lvl {
						if _gdcb.IlvlAttr == levelId {
							return NumberingLevel{_gdcb}
						}
					}
				}
			}
		}
	}
	return NumberingLevel{}
}

// CellProperties returns the cell properties.
func (_eeafa TableConditionalFormatting) CellProperties() CellProperties {
	if _eeafa._adeac.TcPr == nil {
		_eeafa._adeac.TcPr = _db.NewCT_TcPr()
	}
	return CellProperties{_eeafa._adeac.TcPr}
}

// X returns the inner wrapped XML type.
func (_ecgf Table) X() *_db.CT_Tbl { return _ecgf._cgffe }

// SetTextStyleBold set text style of watermark to bold.
func (_bfbbf *WatermarkText) SetTextStyleBold(value bool) {
	if _bfbbf._bgdbe != nil {
		_ebcc := _bfbbf.GetStyle()
		_ebcc.SetBold(value)
		_bfbbf.SetStyle(_ebcc)
	}
}

// X returns the inner wrapped type
func (_dbgf CellBorders) X() *_db.CT_TcBorders { return _dbgf._bee }

// IsChecked returns true if a FormFieldTypeCheckBox is checked.
func (_agfd FormField) IsChecked() bool {
	if _agfd._fbcee.CheckBox == nil {
		return false
	}
	if _agfd._fbcee.CheckBox.Checked != nil {
		return true
	}
	return false
}
func _gab(_daeg _fg.ReaderAt, _eede int64, _affd string) (*Document, error) {
	const _cfg = "\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0052\u0065\u0061\u0064"
	// if !_fd.GetLicenseKey().IsLicensed() && !_fbgb {
	// 	_g.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
	// 	_g.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
	// 	return nil, _bg.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	// }
	_baae := New()
	_baae.Numbering._gebaa = nil
	if len(_affd) > 0 {
		_baae._afc = _affd
	} else {
		// _beac, _afdd := _fd.GenRefId("\u0064\u0072")
		// if _afdd != nil {
		// 	_c.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _afdd)
		// 	return nil, _afdd
		// }
		// _baae._afc = _beac
	}
	// if _cagd := _fd.Track(_baae._afc, _cfg); _cagd != nil {
	// 	_c.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cagd)
	// 	return nil, _cagd
	// }
	_dcgf, _gfee := _be.TempDir("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065-\u0064\u006f\u0063\u0078")
	if _gfee != nil {
		return nil, _gfee
	}
	_baae.TmpPath = _dcgf
	_afdda, _gfee := _a.NewReader(_daeg, _eede)
	if _gfee != nil {
		return nil, _g.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u007a\u0069\u0070\u003a\u0020\u0025\u0073", _gfee)
	}
	_dfa := []*_a.File{}
	_dfa = append(_dfa, _afdda.File...)
	_gaf := false
	for _, _cead := range _dfa {
		if _cead.FileHeader.Name == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
			_gaf = true
			break
		}
	}
	if _gaf {
		_baae.CreateCustomProperties()
	}
	_gaac := _baae._adcb.ConformanceAttr
	_fage := _ed.DecodeMap{}
	_fage.SetOnNewRelationshipFunc(_baae.onNewRelationship)
	_fage.AddTarget(_b.ContentTypesFilename, _baae.ContentTypes.X(), "", 0)
	_fage.AddTarget(_b.BaseRelsFilename, _baae.Rels.X(), "", 0)
	if _gbgc := _fage.Decode(_dfa); _gbgc != nil {
		return nil, _gbgc
	}
	_baae._adcb.ConformanceAttr = _gaac
	for _, _cfga := range _dfa {
		if _cfga == nil {
			continue
		}
		if _eeacg := _baae.AddExtraFileFromZip(_cfga); _eeacg != nil {
			return nil, _eeacg
		}
	}
	if _gaf {
		_eee := false
		for _, _bbfg := range _baae.Rels.X().Relationship {
			if _bbfg.TargetAttr == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
				_eee = true
				break
			}
		}
		if !_eee {
			_baae.AddCustomRelationships()
		}
	}
	return _baae, nil
}
func _gaeed() *_eb.Imagedata {
	_ffde := _eb.NewImagedata()
	_ggcfa := "\u0072\u0049\u0064\u0031"
	_daac := "\u0057A\u0054\u0045\u0052\u004d\u0041\u0052K"
	_ffde.IdAttr = &_ggcfa
	_ffde.TitleAttr = &_daac
	return _ffde
}

// SetBeforeAuto controls if spacing before a paragraph is automatically determined.
func (_ddcg ParagraphSpacing) SetBeforeAuto(b bool) {
	if b {
		_ddcg._ffeb.BeforeAutospacingAttr = &_ca.ST_OnOff{}
		_ddcg._ffeb.BeforeAutospacingAttr.Bool = _b.Bool(true)
	} else {
		_ddcg._ffeb.BeforeAutospacingAttr = nil
	}
}

// X returns the inner wrapped XML type.
func (_ccbcc Styles) X() *_db.Styles { return _ccbcc._cdgg }

// WatermarkText is watermark text within the document.
type WatermarkText struct {
	_gdgad *_db.CT_Picture
	_aadgd *_ge.TextpathStyle
	_bgdbe *_eb.Shape
	_ecgff *_eb.Shapetype
}

// RightToLeft returns true if paragraph text goes from right to left.
func (_adba ParagraphProperties) RightToLeft() bool { return _ebdd(_adba._aedc.RPr.Rtl) }

// Clear resets the numbering.
func (_adgb Numbering) Clear() {
	_adgb._gebaa.AbstractNum = nil
	_adgb._gebaa.Num = nil
	_adgb._gebaa.NumIdMacAtCleanup = nil
	_adgb._gebaa.NumPicBullet = nil
}

// FindNodeByText return node based on matched text and return a slice of node.
func (_cdcb *Nodes) FindNodeByRegexp(r *_ff.Regexp) []Node {
	_bbba := []Node{}
	for _, _cdbcbd := range _cdcb._dcdf {
		if r.MatchString(_cdbcbd.Text()) {
			_bbba = append(_bbba, _cdbcbd)
		}
		_gecea := Nodes{_dcdf: _cdbcbd.Children}
		_bbba = append(_bbba, _gecea.FindNodeByRegexp(r)...)
	}
	return _bbba
}

// SetWidth sets the table with to a specified width.
func (_gfcee TableProperties) SetWidth(d _bc.Distance) {
	_gfcee._faec.TblW = _db.NewCT_TblWidth()
	_gfcee._faec.TblW.TypeAttr = _db.ST_TblWidthDxa
	_gfcee._faec.TblW.WAttr = &_db.ST_MeasurementOrPercent{}
	_gfcee._faec.TblW.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_gfcee._faec.TblW.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(d / _bc.Twips))
}
func _gbge(_eaaf *_db.CT_Border, _edcga _db.ST_Border, _cadf _bd.Color, _baea _bc.Distance) {
	_eaaf.ValAttr = _edcga
	_eaaf.ColorAttr = &_db.ST_HexColor{}
	if _cadf.IsAuto() {
		_eaaf.ColorAttr.ST_HexColorAuto = _db.ST_HexColorAutoAuto
	} else {
		_eaaf.ColorAttr.ST_HexColorRGB = _cadf.AsRGBString()
	}
	if _baea != _bc.Zero {
		_eaaf.SzAttr = _b.Uint64(uint64(_baea / _bc.Point * 8))
	}
}

// SetOutline sets the run to outlined text.
func (_ecaf RunProperties) SetOutline(b bool) {
	if !b {
		_ecaf._cadb.Outline = nil
	} else {
		_ecaf._cadb.Outline = _db.NewCT_OnOff()
	}
}

// CellMargins are the margins for an individual cell.
type CellMargins struct{ _caff *_db.CT_TcMar }

// Nodes return the document's element as nodes.
func (_ecfae *Document) Nodes() Nodes {
	_fabc := []Node{}
	for _, _cbf := range _ecfae._adcb.Body.EG_BlockLevelElts {
		_fabc = append(_fabc, _bdgfg(_ecfae, _cbf.EG_ContentBlockContent, nil)...)
	}
	if _ecfae._adcb.Body.SectPr != nil {
		_fabc = append(_fabc, Node{_bdcf: _ecfae._adcb.Body.SectPr})
	}
	_ccag := Nodes{_dcdf: _fabc}
	return _ccag
}

// Paragraphs returns the paragraphs within a structured document tag.
func (_cfbc StructuredDocumentTag) Paragraphs() []Paragraph {
	if _cfbc._eddg.SdtContent == nil {
		return nil
	}
	_ebefe := []Paragraph{}
	for _, _bdaec := range _cfbc._eddg.SdtContent.P {
		_ebefe = append(_ebefe, Paragraph{_cfbc._ccedc, _bdaec})
	}
	return _ebefe
}

// InsertParagraphAfter adds a new empty paragraph after the relativeTo
// paragraph.
func (_cddd *Document) InsertParagraphAfter(relativeTo Paragraph) Paragraph {
	return _cddd.insertParagraph(relativeTo, false)
}

// SetTop sets the top border to a specified type, color and thickness.
func (_faca ParagraphBorders) SetTop(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_faca._afaaa.Top = _db.NewCT_Border()
	_gbge(_faca._afaaa.Top, t, c, thickness)
}

// NumberingLevel is the definition for numbering for a particular level within
// a NumberingDefinition.
type NumberingLevel struct{ _cfed *_db.CT_Lvl }

// GetRightToLeft returns true if the run text is displayed from right to left.
func (_fagdf RunProperties) GetRightToLeft() bool { return _ebdd(_fagdf._cadb.Rtl) }

// Borders allows manipulation of the table borders.
func (_cged TableStyleProperties) Borders() TableBorders {
	if _cged._dagce.TblBorders == nil {
		_cged._dagce.TblBorders = _db.NewCT_TblBorders()
	}
	return TableBorders{_cged._dagce.TblBorders}
}

// SetAllCaps sets the run to all caps.
func (_dffa RunProperties) SetAllCaps(b bool) {
	if !b {
		_dffa._cadb.Caps = nil
	} else {
		_dffa._cadb.Caps = _db.NewCT_OnOff()
	}
}
func (_agbd *Document) insertStyleFromNode(_cafc Node) {
	if _cafc.Style.X() != nil {
		if _, _gacd := _agbd.Styles.SearchStyleById(_cafc.Style.StyleID()); !_gacd {
			_agbd.Styles.InsertStyle(_cafc.Style)
			_gffe := _cafc.Style.ParagraphProperties()
			_agbd.insertNumberingFromStyleProperties(_cafc._cffcb.Numbering, _gffe)
		}
	}
}

// SetUISortOrder controls the order the style is displayed in the UI.
func (_fgda Style) SetUISortOrder(order int) {
	_fgda._eedcd.UiPriority = _db.NewCT_DecimalNumber()
	_fgda._eedcd.UiPriority.ValAttr = int64(order)
}
func _ebca(_dced *_db.CT_Border, _ebgd _db.ST_Border, _fefea _bd.Color, _eebfc _bc.Distance) {
	_dced.ValAttr = _ebgd
	_dced.ColorAttr = &_db.ST_HexColor{}
	if _fefea.IsAuto() {
		_dced.ColorAttr.ST_HexColorAuto = _db.ST_HexColorAutoAuto
	} else {
		_dced.ColorAttr.ST_HexColorRGB = _fefea.AsRGBString()
	}
	if _eebfc != _bc.Zero {
		_dced.SzAttr = _b.Uint64(uint64(_eebfc / _bc.Point * 8))
	}
}

// SetImprint sets the run to imprinted text.
func (_affab RunProperties) SetImprint(b bool) {
	if !b {
		_affab._cadb.Imprint = nil
	} else {
		_affab._cadb.Imprint = _db.NewCT_OnOff()
	}
}
func (_adfc Paragraph) addEndFldChar() *_db.CT_FldChar {
	_fgcd := _adfc.addFldChar()
	_fgcd.FldCharTypeAttr = _db.ST_FldCharTypeEnd
	return _fgcd
}

// Bold returns true if run font is bold.
func (_aecf RunProperties) Bold() bool {
	_ccbb := _aecf._cadb
	return _ebdd(_ccbb.B) || _ebdd(_ccbb.BCs)
}

// OnOffValue represents an on/off value that can also be unset
type OnOffValue byte

// SetTargetBookmark sets the bookmark target of the hyperlink.
func (_ggbc HyperLink) SetTargetBookmark(bm Bookmark) {
	_ggbc._cfcf.AnchorAttr = _b.String(bm.Name())
	_ggbc._cfcf.IdAttr = nil
}

// SetRightPct sets the cell right margin
func (_abb CellMargins) SetRightPct(pct float64) {
	_abb._caff.Right = _db.NewCT_TblWidth()
	_beb(_abb._caff.Right, pct)
}

// X returns the inner wrapped XML type.
func (_edef Numbering) X() *_db.Numbering { return _edef._gebaa }
func _ffce() (*_aa.CT_Point2D, []*_aa.CT_Point2D) {
	var (
		_fdf int64 = 0
		_bf  int64 = 21600
	)
	_bgc := _aa.ST_Coordinate{ST_CoordinateUnqualified: &_fdf, ST_UniversalMeasure: nil}
	_aee := _aa.ST_Coordinate{ST_CoordinateUnqualified: &_bf, ST_UniversalMeasure: nil}
	_dcb := _aa.NewCT_Point2D()
	_dcb.XAttr = _bgc
	_dcb.YAttr = _bgc
	_gad := []*_aa.CT_Point2D{&_aa.CT_Point2D{XAttr: _bgc, YAttr: _aee}, &_aa.CT_Point2D{XAttr: _aee, YAttr: _aee}, &_aa.CT_Point2D{XAttr: _aee, YAttr: _bgc}, _dcb}
	return _dcb, _gad
}

// SetStyle sets style to the text in watermark.
func (_dccbd *WatermarkText) SetStyle(style _ge.TextpathStyle) {
	_dbag := _dccbd.getShape()
	if _dccbd._bgdbe != nil {
		_caebg := _dccbd._bgdbe.EG_ShapeElements
		if len(_caebg) > 0 && _caebg[0].Textpath != nil {
			var _adffe = style.String()
			_caebg[0].Textpath.StyleAttr = &_adffe
		}
		return
	}
	_ggfa := _dccbd.findNode(_dbag, "\u0074\u0065\u0078\u0074\u0070\u0061\u0074\u0068")
	for _aece, _abdd := range _ggfa.Attrs {
		if _abdd.Name.Local == "\u0073\u0074\u0079l\u0065" {
			_ggfa.Attrs[_aece].Value = style.String()
		}
	}
}
func (_cdffa Paragraph) addInstrText(_abgcd string) *_db.CT_Text {
	_eaaec := _cdffa.AddRun()
	_ceda := _eaaec.X()
	_ccfg := _db.NewEG_RunInnerContent()
	_gadgc := _db.NewCT_Text()
	_gbdc := "\u0070\u0072\u0065\u0073\u0065\u0072\u0076\u0065"
	_gadgc.SpaceAttr = &_gbdc
	_gadgc.Content = "\u0020" + _abgcd + "\u0020"
	_ccfg.InstrText = _gadgc
	_ceda.EG_RunInnerContent = append(_ceda.EG_RunInnerContent, _ccfg)
	return _gadgc
}

// SetTextStyleItalic set text style of watermark to italic.
func (_fefff *WatermarkText) SetTextStyleItalic(value bool) {
	if _fefff._bgdbe != nil {
		_fgdff := _fefff.GetStyle()
		_fgdff.SetItalic(value)
		_fefff.SetStyle(_fgdff)
	}
}

// RowProperties are the properties for a row within a table
type RowProperties struct{ _caaf *_db.CT_TrPr }

const _afe = "\u0046\u006f\u0072\u006d\u0046\u0069\u0065l\u0064\u0054\u0079\u0070\u0065\u0055\u006e\u006b\u006e\u006f\u0077\u006e\u0046\u006fr\u006dF\u0069\u0065\u006c\u0064\u0054\u0079p\u0065\u0054\u0065\u0078\u0074\u0046\u006fr\u006d\u0046\u0069\u0065\u006c\u0064\u0054\u0079\u0070\u0065\u0043\u0068\u0065\u0063\u006b\u0042\u006f\u0078\u0046\u006f\u0072\u006d\u0046i\u0065\u006c\u0064\u0054\u0079\u0070\u0065\u0044\u0072\u006f\u0070\u0044\u006fw\u006e"

// Document is a text document that can be written out in the OOXML .docx
// format. It can be opened from a file on disk and modified, or created from
// scratch.
type Document struct {
	_ffa.DocBase
	_adcb     *_db.Document
	Settings  Settings
	Numbering Numbering
	Styles    Styles
	_adf      []*_db.Hdr
	_gcdc     []_ffa.Relationships
	_abc      []*_db.Ftr
	_adg      []_ffa.Relationships
	_dfc      _ffa.Relationships
	_cdf      []*_aa.Theme
	_eaf      *_db.WebSettings
	_begg     *_db.Fonts
	_acdc     *_db.Endnotes
	_ffcef    *_db.Footnotes
	_aede     []*_ae.Control
	_beag     []*chart
	_afc      string
}

// ParagraphStyleProperties is the styling information for a paragraph.
type ParagraphStyleProperties struct{ _fagf *_db.CT_PPrGeneral }

// Row is a row within a table within a document.
type Row struct {
	_eefba *Document
	_bageg *_db.CT_Row
}

func (_cag *chart) Target() string { return _cag._ddc }

// SetHeadingLevel sets a heading level and style based on the level to a
// paragraph.  The default styles for a new github.com/arcpd/msword document support headings
// from level 1 to 8.
func (_cbgbf ParagraphProperties) SetHeadingLevel(idx int) {
	_cbgbf.SetStyle(_g.Sprintf("\u0048e\u0061\u0064\u0069\u006e\u0067\u0025d", idx))
	if _cbgbf._aedc.NumPr == nil {
		_cbgbf._aedc.NumPr = _db.NewCT_NumPr()
	}
	_cbgbf._aedc.NumPr.Ilvl = _db.NewCT_DecimalNumber()
	_cbgbf._aedc.NumPr.Ilvl.ValAttr = int64(idx)
}

// InitializeDefault constructs a default numbering.
func (_bbab Numbering) InitializeDefault() {
	_bgbdd := _db.NewCT_AbstractNum()
	_bgbdd.MultiLevelType = _db.NewCT_MultiLevelType()
	_bgbdd.MultiLevelType.ValAttr = _db.ST_MultiLevelTypeHybridMultilevel
	_bbab._gebaa.AbstractNum = append(_bbab._gebaa.AbstractNum, _bgbdd)
	_bgbdd.AbstractNumIdAttr = 1
	const _gbfdg = 720
	const _beeg = 720
	const _abdb = 360
	for _cbfa := 0; _cbfa < 9; _cbfa++ {
		_dgbge := _db.NewCT_Lvl()
		_dgbge.IlvlAttr = int64(_cbfa)
		_dgbge.Start = _db.NewCT_DecimalNumber()
		_dgbge.Start.ValAttr = 1
		_dgbge.NumFmt = _db.NewCT_NumFmt()
		_dgbge.NumFmt.ValAttr = _db.ST_NumberFormatBullet
		_dgbge.Suff = _db.NewCT_LevelSuffix()
		_dgbge.Suff.ValAttr = _db.ST_LevelSuffixNothing
		_dgbge.LvlText = _db.NewCT_LevelText()
		_dgbge.LvlText.ValAttr = _b.String("\uf0b7")
		_dgbge.LvlJc = _db.NewCT_Jc()
		_dgbge.LvlJc.ValAttr = _db.ST_JcLeft
		_dgbge.RPr = _db.NewCT_RPr()
		_dgbge.RPr.RFonts = _db.NewCT_Fonts()
		_dgbge.RPr.RFonts.AsciiAttr = _b.String("\u0053\u0079\u006d\u0062\u006f\u006c")
		_dgbge.RPr.RFonts.HAnsiAttr = _b.String("\u0053\u0079\u006d\u0062\u006f\u006c")
		_dgbge.RPr.RFonts.HintAttr = _db.ST_HintDefault
		_dgbge.PPr = _db.NewCT_PPrGeneral()
		_dffeg := int64(_cbfa*_beeg + _gbfdg)
		_dgbge.PPr.Ind = _db.NewCT_Ind()
		_dgbge.PPr.Ind.LeftAttr = &_db.ST_SignedTwipsMeasure{}
		_dgbge.PPr.Ind.LeftAttr.Int64 = _b.Int64(_dffeg)
		_dgbge.PPr.Ind.HangingAttr = &_ca.ST_TwipsMeasure{}
		_dgbge.PPr.Ind.HangingAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(_abdb))
		_bgbdd.Lvl = append(_bgbdd.Lvl, _dgbge)
	}
	_ceac := _db.NewCT_Num()
	_ceac.NumIdAttr = 1
	_ceac.AbstractNumId = _db.NewCT_DecimalNumber()
	_ceac.AbstractNumId.ValAttr = 1
	_bbab._gebaa.Num = append(_bbab._gebaa.Num, _ceac)
}
func _eeb(_beg *_db.CT_TblWidth, _gcd _bc.Distance) {
	_beg.TypeAttr = _db.ST_TblWidthDxa
	_beg.WAttr = &_db.ST_MeasurementOrPercent{}
	_beg.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_beg.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(_gcd / _bc.Dxa))
}

// SetTableIndent sets the Table Indent from the Leading Margin
func (_fgge TableStyleProperties) SetTableIndent(ind _bc.Distance) {
	_fgge._dagce.TblInd = _db.NewCT_TblWidth()
	_fgge._dagce.TblInd.TypeAttr = _db.ST_TblWidthDxa
	_fgge._dagce.TblInd.WAttr = &_db.ST_MeasurementOrPercent{}
	_fgge._dagce.TblInd.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_fgge._dagce.TblInd.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(ind / _bc.Dxa))
}

// SetKeepWithNext controls if this paragraph should be kept with the next.
func (_dcaa ParagraphProperties) SetKeepWithNext(b bool) {
	if !b {
		_dcaa._aedc.KeepNext = nil
	} else {
		_dcaa._aedc.KeepNext = _db.NewCT_OnOff()
	}
}

// SetLeftIndent controls left indent of paragraph.
func (_feba Paragraph) SetLeftIndent(m _bc.Distance) {
	_feba.ensurePPr()
	_edcda := _feba._dcfbd.PPr
	if _edcda.Ind == nil {
		_edcda.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_edcda.Ind.LeftAttr = nil
	} else {
		_edcda.Ind.LeftAttr = &_db.ST_SignedTwipsMeasure{}
		_edcda.Ind.LeftAttr.Int64 = _b.Int64(int64(m / _bc.Twips))
	}
}

// SetTextWrapThrough sets the text wrap to through with a give wrap type.
func (_da AnchoredDrawing) SetTextWrapThrough(option *AnchorDrawWrapOptions) {
	_da._cg.Choice = &_db.WdEG_WrapTypeChoice{}
	_da._cg.Choice.WrapThrough = _db.NewWdCT_WrapThrough()
	_da._cg.Choice.WrapThrough.WrapTextAttr = _db.WdST_WrapTextBothSides
	_gfd := false
	_da._cg.Choice.WrapThrough.WrapPolygon.EditedAttr = &_gfd
	if option == nil {
		option = NewAnchorDrawWrapOptions()
	}
	_da._cg.Choice.WrapThrough.WrapPolygon.Start = option.GetWrapPathStart()
	_da._cg.Choice.WrapThrough.WrapPolygon.LineTo = option.GetWrapPathLineTo()
	_da._cg.LayoutInCellAttr = true
	_da._cg.AllowOverlapAttr = true
}

// EastAsiaFont returns the name of run font family for East Asia.
func (_cacd RunProperties) EastAsiaFont() string {
	if _afdegb := _cacd._cadb.RFonts; _afdegb != nil {
		if _afdegb.EastAsiaAttr != nil {
			return *_afdegb.EastAsiaAttr
		}
	}
	return ""
}

// ClearColor clears the text color.
func (_ceegg RunProperties) ClearColor() { _ceegg._cadb.Color = nil }

// ParagraphSpacing controls the spacing for a paragraph and its lines.
type ParagraphSpacing struct{ _ffeb *_db.CT_Spacing }

// SetSmallCaps sets the run to small caps.
func (_abeef RunProperties) SetSmallCaps(b bool) {
	if !b {
		_abeef._cadb.SmallCaps = nil
	} else {
		_abeef._cadb.SmallCaps = _db.NewCT_OnOff()
	}
}

// HyperLink is a link within a document.
type HyperLink struct {
	_deac *Document
	_cfcf *_db.CT_Hyperlink
}

// SetBottomPct sets the cell bottom margin
func (_ea CellMargins) SetBottomPct(pct float64) {
	_ea._caff.Bottom = _db.NewCT_TblWidth()
	_beb(_ea._caff.Bottom, pct)
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (_bbed ParagraphBorders) SetBottom(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_bbed._afaaa.Bottom = _db.NewCT_Border()
	_gbge(_bbed._afaaa.Bottom, t, c, thickness)
}

// StructuredDocumentTags returns the structured document tags in the document
// which are commonly used in document templates.
func (_bcg *Document) StructuredDocumentTags() []StructuredDocumentTag {
	_ebfd := []StructuredDocumentTag{}
	for _, _gacb := range _bcg._adcb.Body.EG_BlockLevelElts {
		for _, _fggf := range _gacb.EG_ContentBlockContent {
			if _fggf.Sdt != nil {
				_ebfd = append(_ebfd, StructuredDocumentTag{_bcg, _fggf.Sdt})
			}
		}
	}
	return _ebfd
}

const (
	FieldCurrentPage   = "\u0050\u0041\u0047\u0045"
	FieldNumberOfPages = "\u004e\u0055\u004d\u0050\u0041\u0047\u0045\u0053"
	FieldDate          = "\u0044\u0041\u0054\u0045"
	FieldCreateDate    = "\u0043\u0052\u0045\u0041\u0054\u0045\u0044\u0041\u0054\u0045"
	FieldEditTime      = "\u0045\u0044\u0049\u0054\u0054\u0049\u004d\u0045"
	FieldPrintDate     = "\u0050R\u0049\u004e\u0054\u0044\u0041\u0054E"
	FieldSaveDate      = "\u0053\u0041\u0056\u0045\u0044\u0041\u0054\u0045"
	FieldTIme          = "\u0054\u0049\u004d\u0045"
	FieldTOC           = "\u0054\u004f\u0043"
)

type chart struct {
	_bbd *_ee.ChartSpace
	_dfd string
	_ddc string
}

// ParagraphStyles returns only the paragraph styles.
func (_dggab Styles) ParagraphStyles() []Style {
	_fcge := []Style{}
	for _, _cegb := range _dggab._cdgg.Style {
		if _cegb.TypeAttr != _db.ST_StyleTypeParagraph {
			continue
		}
		_fcge = append(_fcge, Style{_cegb})
	}
	return _fcge
}

// AddDropdownList adds dropdown list form field to the paragraph and returns it.
func (_gddbg Paragraph) AddDropdownList(name string) FormField {
	_cfef := _gddbg.addFldCharsForField(name, "\u0046\u004f\u0052M\u0044\u0052\u004f\u0050\u0044\u004f\u0057\u004e")
	_cfef._fbcee.DdList = _db.NewCT_FFDDList()
	return _cfef
}

// VerticalAlign returns the value of run vertical align.
func (_fbcb RunProperties) VerticalAlignment() _ca.ST_VerticalAlignRun {
	if _gaacg := _fbcb._cadb.VertAlign; _gaacg != nil {
		return _gaacg.ValAttr
	}
	return 0
}

// AddFieldWithFormatting adds a field (automatically computed text) to the
// document with field specifc formatting.
func (_agad Run) AddFieldWithFormatting(code string, fmt string, isDirty bool) {
	_dfac := _agad.newIC()
	_dfac.FldChar = _db.NewCT_FldChar()
	_dfac.FldChar.FldCharTypeAttr = _db.ST_FldCharTypeBegin
	if isDirty {
		_dfac.FldChar.DirtyAttr = &_ca.ST_OnOff{}
		_dfac.FldChar.DirtyAttr.Bool = _b.Bool(true)
	}
	_dfac = _agad.newIC()
	_dfac.InstrText = _db.NewCT_Text()
	if fmt != "" {
		_dfac.InstrText.Content = code + "\u0020" + fmt
	} else {
		_dfac.InstrText.Content = code
	}
	_dfac = _agad.newIC()
	_dfac.FldChar = _db.NewCT_FldChar()
	_dfac.FldChar.FldCharTypeAttr = _db.ST_FldCharTypeEnd
}

// Open opens and reads a document from a file (.docx).
func Open(filename string) (*Document, error) {
	_fdce, _afdc := _fe.Open(filename)
	if _afdc != nil {
		return nil, _g.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _afdc)
	}
	defer _fdce.Close()
	_ccfd, _afdc := _fe.Stat(filename)
	if _afdc != nil {
		return nil, _g.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _afdc)
	}
	_ = _ccfd
	return Read(_fdce, _ccfd.Size())
}

// SetHangingIndent controls special indent of paragraph.
func (_caag Paragraph) SetHangingIndent(m _bc.Distance) {
	_caag.ensurePPr()
	_dgbeb := _caag._dcfbd.PPr
	if _dgbeb.Ind == nil {
		_dgbeb.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_dgbeb.Ind.HangingAttr = nil
	} else {
		_dgbeb.Ind.HangingAttr = &_ca.ST_TwipsMeasure{}
		_dgbeb.Ind.HangingAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(m / _bc.Twips))
	}
}

// StructuredDocumentTag are a tagged bit of content in a document.
type StructuredDocumentTag struct {
	_ccedc *Document
	_eddg  *_db.CT_SdtBlock
}

// X returns the inner wrapped XML type.
func (_gbab Run) X() *_db.CT_R { return _gbab._afec }

// NewSettings constructs a new empty Settings
func NewSettings() Settings {
	_bdefa := _db.NewSettings()
	_bdefa.Compat = _db.NewCT_Compat()
	_cfege := _db.NewCT_CompatSetting()
	_cfege.NameAttr = _b.String("\u0063\u006f\u006d\u0070\u0061\u0074\u0069\u0062\u0069\u006c\u0069\u0074y\u004d\u006f\u0064\u0065")
	_cfege.UriAttr = _b.String("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006fff\u0069\u0063\u0065/\u0077o\u0072\u0064")
	_cfege.ValAttr = _b.String("\u0031\u0035")
	_bdefa.Compat.CompatSetting = append(_bdefa.Compat.CompatSetting, _cfege)
	return Settings{_bdefa}
}

// SetBold sets the run to bold.
func (_afdegf RunProperties) SetBold(b bool) {
	if !b {
		_afdegf._cadb.B = nil
		_afdegf._cadb.BCs = nil
	} else {
		_afdegf._cadb.B = _db.NewCT_OnOff()
		_afdegf._cadb.BCs = _db.NewCT_OnOff()
	}
}
func (_aeac FormFieldType) String() string {
	if _aeac >= FormFieldType(len(_dfcd)-1) {
		return _g.Sprintf("\u0046\u006f\u0072\u006d\u0046\u0069\u0065\u006c\u0064\u0054\u0079\u0070e\u0028\u0025\u0064\u0029", _aeac)
	}
	return _afe[_dfcd[_aeac]:_dfcd[_aeac+1]]
}

// SetContextualSpacing controls whether to Ignore Spacing Above and Below When
// Using Identical Styles
func (_ebcf ParagraphStyleProperties) SetContextualSpacing(b bool) {
	if !b {
		_ebcf._fagf.ContextualSpacing = nil
	} else {
		_ebcf._fagf.ContextualSpacing = _db.NewCT_OnOff()
	}
}

// SetColumnBandSize sets the number of Columns in the column band
func (_cfff TableStyleProperties) SetColumnBandSize(cols int64) {
	_cfff._dagce.TblStyleColBandSize = _db.NewCT_DecimalNumber()
	_cfff._dagce.TblStyleColBandSize.ValAttr = cols
}
func (_aggbb Run) newIC() *_db.EG_RunInnerContent {
	_fege := _db.NewEG_RunInnerContent()
	_aggbb._afec.EG_RunInnerContent = append(_aggbb._afec.EG_RunInnerContent, _fege)
	return _fege
}
func (_aec *Document) validateTableCells() error {
	for _, _dda := range _aec._adcb.Body.EG_BlockLevelElts {
		for _, _dccg := range _dda.EG_ContentBlockContent {
			for _, _cdcc := range _dccg.Tbl {
				for _, _ffda := range _cdcc.EG_ContentRowContent {
					for _, _fbfb := range _ffda.Tr {
						_fbgcf := false
						for _, _efa := range _fbfb.EG_ContentCellContent {
							_dbcd := false
							for _, _gcae := range _efa.Tc {
								_fbgcf = true
								for _, _gade := range _gcae.EG_BlockLevelElts {
									for _, _ggf := range _gade.EG_ContentBlockContent {
										if len(_ggf.P) > 0 {
											_dbcd = true
											break
										}
									}
								}
							}
							if !_dbcd {
								return _bg.New("t\u0061\u0062\u006c\u0065\u0020\u0063e\u006c\u006c\u0020\u006d\u0075\u0073t\u0020\u0063\u006f\u006e\u0074\u0061\u0069n\u0020\u0061\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061p\u0068")
							}
						}
						if !_fbgcf {
							return _bg.New("\u0074\u0061b\u006c\u0065\u0020\u0072\u006f\u0077\u0020\u006d\u0075\u0073\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0020ce\u006c\u006c")
						}
					}
				}
			}
		}
	}
	return nil
}

// SetHANSITheme sets the font H ANSI Theme.
func (_gfecc Fonts) SetHANSITheme(t _db.ST_Theme) { _gfecc._bdge.HAnsiThemeAttr = t }

// SetVAlignment sets the vertical alignment for an anchored image.
func (_ebe AnchoredDrawing) SetVAlignment(v _db.WdST_AlignV) {
	_ebe._cg.PositionV.Choice = &_db.WdCT_PosVChoice{}
	_ebe._cg.PositionV.Choice.Align = v
}

// SetNextStyle sets the style that the next paragraph will use.
func (_geeeb Style) SetNextStyle(name string) {
	if name == "" {
		_geeeb._eedcd.Next = nil
	} else {
		_geeeb._eedcd.Next = _db.NewCT_String()
		_geeeb._eedcd.Next.ValAttr = name
	}
}

// Pict returns the pict object.
func (_abbb *WatermarkText) Pict() *_db.CT_Picture { return _abbb._gdgad }

// SetEffect sets a text effect on the run.
func (_eadc RunProperties) SetEffect(e _db.ST_TextEffect) {
	if e == _db.ST_TextEffectUnset {
		_eadc._cadb.Effect = nil
	} else {
		_eadc._cadb.Effect = _db.NewCT_TextEffect()
		_eadc._cadb.Effect.ValAttr = e
	}
}

// ExtractText returns text from the document as a DocText object.
func (_daae *Document) ExtractText() *DocText {
	_gfga := []TextItem{}
	for _, _fabfd := range _daae._adcb.Body.EG_BlockLevelElts {
		_gfga = append(_gfga, _cccg(_fabfd.EG_ContentBlockContent, nil)...)
	}
	var _dfae []listItemInfo
	_dacac := _daae.Paragraphs()
	for _, _ddddg := range _dacac {
		_deafa := _fcdc(_daae, _ddddg)
		_dfae = append(_dfae, _deafa)
	}
	_fgef := _fbbf(_daae)
	return &DocText{Items: _gfga, _bbaf: _dfae, _bgfd: _fgef}
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (_deeb TableProperties) SetCellSpacingAuto() {
	_deeb._faec.TblCellSpacing = _db.NewCT_TblWidth()
	_deeb._faec.TblCellSpacing.TypeAttr = _db.ST_TblWidthAuto
}

// Strike returns true if run is striked.
func (_daee RunProperties) Strike() bool { return _ebdd(_daee._cadb.Strike) }

// Footnote is an individual footnote reference within the document.
type Footnote struct {
	_abca *Document
	_cefg *_db.CT_FtnEdn
}

// SetAlignment sets the alignment of a table within the page.
func (_fbebe TableProperties) SetAlignment(align _db.ST_JcTable) {
	if align == _db.ST_JcTableUnset {
		_fbebe._faec.Jc = nil
	} else {
		_fbebe._faec.Jc = _db.NewCT_JcTable()
		_fbebe._faec.Jc.ValAttr = align
	}
}
func (_cddcdf Styles) initializeStyleDefaults() {
	_fcgb := _cddcdf.AddStyle("\u004e\u006f\u0072\u006d\u0061\u006c", _db.ST_StyleTypeParagraph, true)
	_fcgb.SetName("\u004e\u006f\u0072\u006d\u0061\u006c")
	_fcgb.SetPrimaryStyle(true)
	_dged := _cddcdf.AddStyle("D\u0065f\u0061\u0075\u006c\u0074\u0050\u0061\u0072\u0061g\u0072\u0061\u0070\u0068Fo\u006e\u0074", _db.ST_StyleTypeCharacter, true)
	_dged.SetName("\u0044\u0065\u0066\u0061ul\u0074\u0020\u0050\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0020\u0046\u006fn\u0074")
	_dged.SetUISortOrder(1)
	_dged.SetSemiHidden(true)
	_dged.SetUnhideWhenUsed(true)
	_faccf := _cddcdf.AddStyle("\u0054i\u0074\u006c\u0065\u0043\u0068\u0061r", _db.ST_StyleTypeCharacter, false)
	_faccf.SetName("\u0054\u0069\u0074\u006c\u0065\u0020\u0043\u0068\u0061\u0072")
	_faccf.SetBasedOn(_dged.StyleID())
	_faccf.SetLinkedStyle("\u0054\u0069\u0074l\u0065")
	_faccf.SetUISortOrder(10)
	_faccf.RunProperties().Fonts().SetASCIITheme(_db.ST_ThemeMajorAscii)
	_faccf.RunProperties().Fonts().SetEastAsiaTheme(_db.ST_ThemeMajorEastAsia)
	_faccf.RunProperties().Fonts().SetHANSITheme(_db.ST_ThemeMajorHAnsi)
	_faccf.RunProperties().Fonts().SetCSTheme(_db.ST_ThemeMajorBidi)
	_faccf.RunProperties().SetSize(28 * _bc.Point)
	_faccf.RunProperties().SetKerning(14 * _bc.Point)
	_faccf.RunProperties().SetCharacterSpacing(-10 * _bc.Twips)
	_cfegg := _cddcdf.AddStyle("\u0054\u0069\u0074l\u0065", _db.ST_StyleTypeParagraph, false)
	_cfegg.SetName("\u0054\u0069\u0074l\u0065")
	_cfegg.SetBasedOn(_fcgb.StyleID())
	_cfegg.SetNextStyle(_fcgb.StyleID())
	_cfegg.SetLinkedStyle(_faccf.StyleID())
	_cfegg.SetUISortOrder(10)
	_cfegg.SetPrimaryStyle(true)
	_cfegg.ParagraphProperties().SetContextualSpacing(true)
	_cfegg.RunProperties().Fonts().SetASCIITheme(_db.ST_ThemeMajorAscii)
	_cfegg.RunProperties().Fonts().SetEastAsiaTheme(_db.ST_ThemeMajorEastAsia)
	_cfegg.RunProperties().Fonts().SetHANSITheme(_db.ST_ThemeMajorHAnsi)
	_cfegg.RunProperties().Fonts().SetCSTheme(_db.ST_ThemeMajorBidi)
	_cfegg.RunProperties().SetSize(28 * _bc.Point)
	_cfegg.RunProperties().SetKerning(14 * _bc.Point)
	_cfegg.RunProperties().SetCharacterSpacing(-10 * _bc.Twips)
	_cgcag := _cddcdf.AddStyle("T\u0061\u0062\u006c\u0065\u004e\u006f\u0072\u006d\u0061\u006c", _db.ST_StyleTypeTable, false)
	_cgcag.SetName("\u004e\u006f\u0072m\u0061\u006c\u0020\u0054\u0061\u0062\u006c\u0065")
	_cgcag.SetUISortOrder(99)
	_cgcag.SetSemiHidden(true)
	_cgcag.SetUnhideWhenUsed(true)
	_cgcag.X().TblPr = _db.NewCT_TblPrBase()
	_ebbe := NewTableWidth()
	_cgcag.X().TblPr.TblInd = _ebbe.X()
	_ebbe.SetValue(0 * _bc.Dxa)
	_cgcag.X().TblPr.TblCellMar = _db.NewCT_TblCellMar()
	_ebbe = NewTableWidth()
	_cgcag.X().TblPr.TblCellMar.Top = _ebbe.X()
	_ebbe.SetValue(0 * _bc.Dxa)
	_ebbe = NewTableWidth()
	_cgcag.X().TblPr.TblCellMar.Bottom = _ebbe.X()
	_ebbe.SetValue(0 * _bc.Dxa)
	_ebbe = NewTableWidth()
	_cgcag.X().TblPr.TblCellMar.Left = _ebbe.X()
	_ebbe.SetValue(108 * _bc.Dxa)
	_ebbe = NewTableWidth()
	_cgcag.X().TblPr.TblCellMar.Right = _ebbe.X()
	_ebbe.SetValue(108 * _bc.Dxa)
	_fcdaf := _cddcdf.AddStyle("\u004e\u006f\u004c\u0069\u0073\u0074", _db.ST_StyleTypeNumbering, false)
	_fcdaf.SetName("\u004eo\u0020\u004c\u0069\u0073\u0074")
	_fcdaf.SetUISortOrder(1)
	_fcdaf.SetSemiHidden(true)
	_fcdaf.SetUnhideWhenUsed(true)
	_cfefdf := []_bc.Distance{16, 13, 12, 11, 11, 11, 11, 11, 11}
	_acde := []_bc.Distance{240, 40, 40, 40, 40, 40, 40, 40, 40}
	for _ecda := 0; _ecda < 9; _ecda++ {
		_aeadf := _g.Sprintf("\u0048e\u0061\u0064\u0069\u006e\u0067\u0025d", _ecda+1)
		_aedea := _cddcdf.AddStyle(_aeadf+"\u0043\u0068\u0061\u0072", _db.ST_StyleTypeCharacter, false)
		_aedea.SetName(_g.Sprintf("\u0048e\u0061d\u0069\u006e\u0067\u0020\u0025\u0064\u0020\u0043\u0068\u0061\u0072", _ecda+1))
		_aedea.SetBasedOn(_dged.StyleID())
		_aedea.SetLinkedStyle(_aeadf)
		_aedea.SetUISortOrder(9 + _ecda)
		_aedea.RunProperties().SetSize(_cfefdf[_ecda] * _bc.Point)
		_dcdga := _cddcdf.AddStyle(_aeadf, _db.ST_StyleTypeParagraph, false)
		_dcdga.SetName(_g.Sprintf("\u0068\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0025\u0064", _ecda+1))
		_dcdga.SetNextStyle(_fcgb.StyleID())
		_dcdga.SetLinkedStyle(_dcdga.StyleID())
		_dcdga.SetUISortOrder(9 + _ecda)
		_dcdga.SetPrimaryStyle(true)
		_dcdga.ParagraphProperties().SetKeepNext(true)
		_dcdga.ParagraphProperties().SetSpacing(_acde[_ecda]*_bc.Twips, 0)
		_dcdga.ParagraphProperties().SetOutlineLevel(_ecda)
		_dcdga.RunProperties().SetSize(_cfefdf[_ecda] * _bc.Point)
	}
}

// RemoveFootnote removes a footnote from both the paragraph and the document
// the requested footnote must be anchored on the paragraph being referenced.
func (_aedb Paragraph) RemoveFootnote(id int64) {
	_debee := _aedb._cfge._ffcef
	var _cdgfd int
	for _ddfdg, _aeafe := range _debee.CT_Footnotes.Footnote {
		if _aeafe.IdAttr == id {
			_cdgfd = _ddfdg
		}
	}
	_cdgfd = 0
	_debee.CT_Footnotes.Footnote[_cdgfd] = nil
	_debee.CT_Footnotes.Footnote[_cdgfd] = _debee.CT_Footnotes.Footnote[len(_debee.CT_Footnotes.Footnote)-1]
	_debee.CT_Footnotes.Footnote = _debee.CT_Footnotes.Footnote[:len(_debee.CT_Footnotes.Footnote)-1]
	var _fbbge Run
	for _, _faeea := range _aedb.Runs() {
		if _bbgaf, _bbfc := _faeea.IsFootnote(); _bbgaf {
			if _bbfc == id {
				_fbbge = _faeea
			}
		}
	}
	_aedb.RemoveRun(_fbbge)
}

// GetChartSpaceByRelId returns a *crt.ChartSpace with the associated relation ID in the
// document.
func (_gebf *Document) GetChartSpaceByRelId(relId string) *_ee.ChartSpace {
	_fdag := _gebf._dfc.GetTargetByRelId(relId)
	for _, _fcded := range _gebf._beag {
		if _fdag == _fcded.Target() {
			return _fcded._bbd
		}
	}
	return nil
}

// X returns the inner wrapped XML type.
func (_edc Color) X() *_db.CT_Color { return _edc._fcd }

// X returns the inner wrapped XML type.
func (_aead RunProperties) X() *_db.CT_RPr { return _aead._cadb }

// AddLevel adds a new numbering level to a NumberingDefinition.
func (_cgfa NumberingDefinition) AddLevel() NumberingLevel {
	_ddac := _db.NewCT_Lvl()
	_ddac.Start = &_db.CT_DecimalNumber{ValAttr: 1}
	_ddac.IlvlAttr = int64(len(_cgfa._bbgd.Lvl))
	_cgfa._bbgd.Lvl = append(_cgfa._bbgd.Lvl, _ddac)
	return NumberingLevel{_ddac}
}

// Tables returns the tables defined in the header.
func (_eeff Header) Tables() []Table {
	_afga := []Table{}
	if _eeff._ecead == nil {
		return nil
	}
	for _, _cdgff := range _eeff._ecead.EG_ContentBlockContent {
		for _, _fbee := range _eeff._bgac.tables(_cdgff) {
			_afga = append(_afga, _fbee)
		}
	}
	return _afga
}

// AddPageBreak adds a page break to a run.
func (_deggd Run) AddPageBreak() {
	_dddc := _deggd.newIC()
	_dddc.Br = _db.NewCT_Br()
	_dddc.Br.TypeAttr = _db.ST_BrTypePage
}

// Footer is a footer for a document section.
type Footer struct {
	_bcge *Document
	_eeed *_db.Ftr
}

// GetHeader gets a section Header for given type t [ST_HdrFtrDefault, ST_HdrFtrEven, ST_HdrFtrFirst]
func (_decf Section) GetHeader(t _db.ST_HdrFtr) (Header, bool) {
	for _, _cegg := range _decf._ggcfb.EG_HdrFtrReferences {
		if _cegg.HeaderReference.TypeAttr == t {
			for _, _eabf := range _decf._bagegb.Headers() {
				_fbgba := _decf._bagegb._dfc.FindRIDForN(_eabf.Index(), _b.HeaderType)
				if _fbgba == _cegg.HeaderReference.IdAttr {
					return _eabf, true
				}
			}
		}
	}
	return Header{}, false
}

// Headers returns the headers defined in the document.
func (_ffd *Document) Headers() []Header {
	_eec := []Header{}
	for _, _gb := range _ffd._adf {
		_eec = append(_eec, Header{_ffd, _gb})
	}
	return _eec
}
func _cccg(_ggcac []*_db.EG_ContentBlockContent, _aefb *TableInfo) []TextItem {
	_ddaf := []TextItem{}
	for _, _adad := range _ggcac {
		if _aae := _adad.Sdt; _aae != nil {
			if _cgfd := _aae.SdtContent; _cgfd != nil {
				_ddaf = append(_ddaf, _eefc(_cgfd.P, _aefb, nil)...)
			}
		}
		_ddaf = append(_ddaf, _eefc(_adad.P, _aefb, nil)...)
		for _, _gcff := range _adad.Tbl {
			for _dcdg, _bfed := range _gcff.EG_ContentRowContent {
				for _, _gdgc := range _bfed.Tr {
					for _ddec, _baba := range _gdgc.EG_ContentCellContent {
						for _, _cgda := range _baba.Tc {
							_gbcgf := &TableInfo{Table: _gcff, Row: _gdgc, Cell: _cgda, RowIndex: _dcdg, ColIndex: _ddec}
							for _, _dfeee := range _cgda.EG_BlockLevelElts {
								_ddaf = append(_ddaf, _cccg(_dfeee.EG_ContentBlockContent, _gbcgf)...)
							}
						}
					}
				}
			}
		}
	}
	return _ddaf
}

// SetLineSpacing sets the spacing between lines in a paragraph.
func (_gadcb ParagraphSpacing) SetLineSpacing(d _bc.Distance, rule _db.ST_LineSpacingRule) {
	if rule == _db.ST_LineSpacingRuleUnset {
		_gadcb._ffeb.LineRuleAttr = _db.ST_LineSpacingRuleUnset
		_gadcb._ffeb.LineAttr = nil
	} else {
		_gadcb._ffeb.LineRuleAttr = rule
		_gadcb._ffeb.LineAttr = &_db.ST_SignedTwipsMeasure{}
		_gadcb._ffeb.LineAttr.Int64 = _b.Int64(int64(d / _bc.Twips))
	}
}

// SetVerticalBanding controls the conditional formatting for vertical banding.
func (_dbegb TableLook) SetVerticalBanding(on bool) {
	if !on {
		_dbegb._febce.NoVBandAttr = &_ca.ST_OnOff{}
		_dbegb._febce.NoVBandAttr.ST_OnOff1 = _ca.ST_OnOff1On
	} else {
		_dbegb._febce.NoVBandAttr = &_ca.ST_OnOff{}
		_dbegb._febce.NoVBandAttr.ST_OnOff1 = _ca.ST_OnOff1Off
	}
}

// Pict returns the pict object.
func (_becaf *WatermarkPicture) Pict() *_db.CT_Picture { return _becaf._dgdd }

// Emboss returns true if run emboss is on.
func (_gfebb RunProperties) Emboss() bool { return _ebdd(_gfebb._cadb.Emboss) }

// SetVerticalMerge controls the vertical merging of cells.
func (_bga CellProperties) SetVerticalMerge(mergeVal _db.ST_Merge) {
	if mergeVal == _db.ST_MergeUnset {
		_bga._beea.VMerge = nil
	} else {
		_bga._beea.VMerge = _db.NewCT_VMerge()
		_bga._beea.VMerge.ValAttr = mergeVal
	}
}

// SetXOffset sets the X offset for an image relative to the origin.
func (_gfg AnchoredDrawing) SetXOffset(x _bc.Distance) {
	_gfg._cg.PositionH.Choice = &_db.WdCT_PosHChoice{}
	_gfg._cg.PositionH.Choice.PosOffset = _b.Int32(int32(x / _bc.EMU))
}
func _gaed(_bdgf *_db.CT_Tbl, _dff *_db.CT_P, _efe bool) *_db.CT_Tbl {
	for _, _ebc := range _bdgf.EG_ContentRowContent {
		for _, _ece := range _ebc.Tr {
			for _, _febc := range _ece.EG_ContentCellContent {
				for _, _fgeg := range _febc.Tc {
					for _ccfe, _gfc := range _fgeg.EG_BlockLevelElts {
						for _, _cdgd := range _gfc.EG_ContentBlockContent {
							for _gcged, _eba := range _cdgd.P {
								if _eba == _dff {
									_ffg := _db.NewEG_BlockLevelElts()
									_ace := _db.NewEG_ContentBlockContent()
									_ffg.EG_ContentBlockContent = append(_ffg.EG_ContentBlockContent, _ace)
									_febe := _db.NewCT_Tbl()
									_ace.Tbl = append(_ace.Tbl, _febe)
									_fgeg.EG_BlockLevelElts = append(_fgeg.EG_BlockLevelElts, nil)
									if _efe {
										copy(_fgeg.EG_BlockLevelElts[_ccfe+1:], _fgeg.EG_BlockLevelElts[_ccfe:])
										_fgeg.EG_BlockLevelElts[_ccfe] = _ffg
										if _gcged != 0 {
											_fegf := _db.NewEG_BlockLevelElts()
											_gaa := _db.NewEG_ContentBlockContent()
											_fegf.EG_ContentBlockContent = append(_fegf.EG_ContentBlockContent, _gaa)
											_gaa.P = _cdgd.P[:_gcged]
											_fgeg.EG_BlockLevelElts = append(_fgeg.EG_BlockLevelElts, nil)
											copy(_fgeg.EG_BlockLevelElts[_ccfe+1:], _fgeg.EG_BlockLevelElts[_ccfe:])
											_fgeg.EG_BlockLevelElts[_ccfe] = _fegf
										}
										_cdgd.P = _cdgd.P[_gcged:]
									} else {
										copy(_fgeg.EG_BlockLevelElts[_ccfe+2:], _fgeg.EG_BlockLevelElts[_ccfe+1:])
										_fgeg.EG_BlockLevelElts[_ccfe+1] = _ffg
										if _gcged != len(_cdgd.P)-1 {
											_agd := _db.NewEG_BlockLevelElts()
											_aeea := _db.NewEG_ContentBlockContent()
											_agd.EG_ContentBlockContent = append(_agd.EG_ContentBlockContent, _aeea)
											_aeea.P = _cdgd.P[_gcged+1:]
											_fgeg.EG_BlockLevelElts = append(_fgeg.EG_BlockLevelElts, nil)
											copy(_fgeg.EG_BlockLevelElts[_ccfe+3:], _fgeg.EG_BlockLevelElts[_ccfe+2:])
											_fgeg.EG_BlockLevelElts[_ccfe+2] = _agd
										} else {
											_dgcfg := _db.NewEG_BlockLevelElts()
											_eed := _db.NewEG_ContentBlockContent()
											_dgcfg.EG_ContentBlockContent = append(_dgcfg.EG_ContentBlockContent, _eed)
											_eed.P = []*_db.CT_P{_db.NewCT_P()}
											_fgeg.EG_BlockLevelElts = append(_fgeg.EG_BlockLevelElts, nil)
											copy(_fgeg.EG_BlockLevelElts[_ccfe+3:], _fgeg.EG_BlockLevelElts[_ccfe+2:])
											_fgeg.EG_BlockLevelElts[_ccfe+2] = _dgcfg
										}
										_cdgd.P = _cdgd.P[:_gcged+1]
									}
									return _febe
								}
							}
							for _, _fedg := range _cdgd.Tbl {
								_gce := _gaed(_fedg, _dff, _efe)
								if _gce != nil {
									return _gce
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// SetFirstRow controls the conditional formatting for the first row in a table.
func (_aegda TableLook) SetFirstRow(on bool) {
	if !on {
		_aegda._febce.FirstRowAttr = &_ca.ST_OnOff{}
		_aegda._febce.FirstRowAttr.ST_OnOff1 = _ca.ST_OnOff1Off
	} else {
		_aegda._febce.FirstRowAttr = &_ca.ST_OnOff{}
		_aegda._febce.FirstRowAttr.ST_OnOff1 = _ca.ST_OnOff1On
	}
}

// SetLeft sets the left border to a specified type, color and thickness.
func (_edaefe TableBorders) SetLeft(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_edaefe._fcdb.Left = _db.NewCT_Border()
	_ebca(_edaefe._fcdb.Left, t, c, thickness)
}

// GetWrapPathStart return wrapPath start value.
func (_dcg AnchorDrawWrapOptions) GetWrapPathStart() *_aa.CT_Point2D { return _dcg._ega }

// AddParagraph adds a paragraph to the header.
func (_eaed Header) AddParagraph() Paragraph {
	_bcdg := _db.NewEG_ContentBlockContent()
	_eaed._ecead.EG_ContentBlockContent = append(_eaed._ecead.EG_ContentBlockContent, _bcdg)
	_adcaf := _db.NewCT_P()
	_bcdg.P = append(_bcdg.P, _adcaf)
	return Paragraph{_eaed._bgac, _adcaf}
}

// SetHighlight highlights text in a specified color.
func (_acfed RunProperties) SetHighlight(c _db.ST_HighlightColor) {
	_acfed._cadb.Highlight = _db.NewCT_Highlight()
	_acfed._cadb.Highlight.ValAttr = c
}

// SetTargetByRef sets the URL target of the hyperlink and is more efficient if a link
// destination will be used many times.
func (_bdba HyperLink) SetTargetByRef(link _ffa.Hyperlink) {
	_bdba._cfcf.IdAttr = _b.String(_ffa.Relationship(link).ID())
	_bdba._cfcf.AnchorAttr = nil
}

// SetBeforeSpacing sets spacing above paragraph.
func (_eeaef Paragraph) SetBeforeSpacing(d _bc.Distance) {
	_eeaef.ensurePPr()
	if _eeaef._dcfbd.PPr.Spacing == nil {
		_eeaef._dcfbd.PPr.Spacing = _db.NewCT_Spacing()
	}
	_ggged := _eeaef._dcfbd.PPr.Spacing
	_ggged.BeforeAttr = &_ca.ST_TwipsMeasure{}
	_ggged.BeforeAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(d / _bc.Twips))
}

// Borders allows manipulation of the table borders.
func (_dggae TableProperties) Borders() TableBorders {
	if _dggae._faec.TblBorders == nil {
		_dggae._faec.TblBorders = _db.NewCT_TblBorders()
	}
	return TableBorders{_dggae._faec.TblBorders}
}

// SetRight sets the right border to a specified type, color and thickness.
func (_dbged ParagraphBorders) SetRight(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_dbged._afaaa.Right = _db.NewCT_Border()
	_gbge(_dbged._afaaa.Right, t, c, thickness)
}

// ComplexSizeValue returns the value of paragraph font size for complex fonts in points.
func (_geaea ParagraphProperties) ComplexSizeValue() float64 {
	if _adbe := _geaea._aedc.RPr.SzCs; _adbe != nil {
		_faccd := _adbe.ValAttr
		if _faccd.ST_UnsignedDecimalNumber != nil {
			return float64(*_faccd.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// SetKeepNext controls if the paragraph is kept with the next paragraph.
func (_dcgec ParagraphStyleProperties) SetKeepNext(b bool) {
	if !b {
		_dcgec._fagf.KeepNext = nil
	} else {
		_dcgec._fagf.KeepNext = _db.NewCT_OnOff()
	}
}
func (_bedc Styles) initializeDocDefaults() {
	_bedc._cdgg.DocDefaults = _db.NewCT_DocDefaults()
	_bedc._cdgg.DocDefaults.RPrDefault = _db.NewCT_RPrDefault()
	_bedc._cdgg.DocDefaults.RPrDefault.RPr = _db.NewCT_RPr()
	_gagb := RunProperties{_bedc._cdgg.DocDefaults.RPrDefault.RPr}
	_gagb.SetSize(12 * _bc.Point)
	_gagb.Fonts().SetASCIITheme(_db.ST_ThemeMajorAscii)
	_gagb.Fonts().SetEastAsiaTheme(_db.ST_ThemeMajorEastAsia)
	_gagb.Fonts().SetHANSITheme(_db.ST_ThemeMajorHAnsi)
	_gagb.Fonts().SetCSTheme(_db.ST_ThemeMajorBidi)
	_gagb.X().Lang = _db.NewCT_Language()
	_gagb.X().Lang.ValAttr = _b.String("\u0065\u006e\u002dU\u0053")
	_gagb.X().Lang.EastAsiaAttr = _b.String("\u0065\u006e\u002dU\u0053")
	_gagb.X().Lang.BidiAttr = _b.String("\u0061\u0072\u002dS\u0041")
	_bedc._cdgg.DocDefaults.PPrDefault = _db.NewCT_PPrDefault()
}

// WatermarkPicture is watermark picture within document.
type WatermarkPicture struct {
	_dgdd  *_db.CT_Picture
	_bddb  *_ge.ShapeStyle
	_febge *_eb.Shape
	_baac  *_eb.Shapetype
}

// RightToLeft returns true if run text goes from right to left.
func (_aeaa RunProperties) RightToLeft() bool { return _ebdd(_aeaa._cadb.Rtl) }

// SetSize sets the size of the displayed image on the page.
func (_ag AnchoredDrawing) SetSize(w, h _bc.Distance) {
	_ag._cg.Extent.CxAttr = int64(float64(w*_bc.Pixel72) / _bc.EMU)
	_ag._cg.Extent.CyAttr = int64(float64(h*_bc.Pixel72) / _bc.EMU)
}

// AddTextInput adds text input form field to the paragraph and returns it.
func (_dfbg Paragraph) AddTextInput(name string) FormField {
	_ddbdc := _dfbg.addFldCharsForField(name, "\u0046\u004f\u0052\u004d\u0054\u0045\u0058\u0054")
	_ddbdc._fbcee.TextInput = _db.NewCT_FFTextInput()
	return _ddbdc
}
func _fdda(_dggbb *_db.CT_P, _efcbd map[string]string) {
	for _, _bebg := range _dggbb.EG_PContent {
		for _, _bece := range _bebg.EG_ContentRunContent {
			if _bece.R != nil {
				for _, _acaf := range _bece.R.EG_RunInnerContent {
					_dfef := _acaf.Drawing
					if _dfef != nil {
						for _, _fggad := range _dfef.Anchor {
							for _, _fbea := range _fggad.Graphic.GraphicData.Any {
								switch _ddde := _fbea.(type) {
								case *_ba.Pic:
									if _ddde.BlipFill != nil && _ddde.BlipFill.Blip != nil {
										_aefa(_ddde.BlipFill.Blip, _efcbd)
									}
								default:
								}
							}
						}
						for _, _daad := range _dfef.Inline {
							for _, _ggcg := range _daad.Graphic.GraphicData.Any {
								switch _fefdb := _ggcg.(type) {
								case *_ba.Pic:
									if _fefdb.BlipFill != nil && _fefdb.BlipFill.Blip != nil {
										_aefa(_fefdb.BlipFill.Blip, _efcbd)
									}
								default:
								}
							}
						}
					}
				}
			}
		}
	}
}

// Spacing returns the paragraph spacing settings.
func (_cafce ParagraphProperties) Spacing() ParagraphSpacing {
	if _cafce._aedc.Spacing == nil {
		_cafce._aedc.Spacing = _db.NewCT_Spacing()
	}
	return ParagraphSpacing{_cafce._aedc.Spacing}
}
func (_cbbd *Document) insertImageFromNode(_baaeb Node) {
	for _, _cga := range _baaeb.AnchoredDrawings {
		if _agcd, _gaga := _cga.GetImage(); _gaga {
			_gbfd, _dbce := _ffa.ImageFromFile(_agcd.Path())
			if _dbce != nil {
				_c.Log.Debug("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063r\u0065\u0061\u0074\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u0073", _dbce)
			}
			_bace, _dbce := _cbbd.AddImage(_gbfd)
			if _dbce != nil {
				_c.Log.Debug("u\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020i\u006d\u0061\u0067\u0065\u0020\u0074\u006f \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020%\u0073", _dbce)
			}
			_abag := _cbbd._dfc.GetByRelId(_bace.RelID())
			_abag.SetID(_agcd.RelID())
		}
	}
	for _, _gedfd := range _baaeb.InlineDrawings {
		if _cabe, _cdcg := _gedfd.GetImage(); _cdcg {
			_gbfe, _gbfg := _ffa.ImageFromFile(_cabe.Path())
			if _gbfg != nil {
				_c.Log.Debug("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063r\u0065\u0061\u0074\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u0073", _gbfg)
			}
			_agag, _gbfg := _cbbd.AddImage(_gbfe)
			if _gbfg != nil {
				_c.Log.Debug("u\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020i\u006d\u0061\u0067\u0065\u0020\u0074\u006f \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020%\u0073", _gbfg)
			}
			_bbcg := _cbbd._dfc.GetByRelId(_agag.RelID())
			_bbcg.SetID(_cabe.RelID())
		}
	}
}

// SetCellSpacing sets the cell spacing within a table.
func (_dgbebf TableProperties) SetCellSpacing(m _bc.Distance) {
	_dgbebf._faec.TblCellSpacing = _db.NewCT_TblWidth()
	_dgbebf._faec.TblCellSpacing.TypeAttr = _db.ST_TblWidthDxa
	_dgbebf._faec.TblCellSpacing.WAttr = &_db.ST_MeasurementOrPercent{}
	_dgbebf._faec.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_dgbebf._faec.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(m / _bc.Dxa))
}
func _fcdc(_debd *Document, _bebbf Paragraph) listItemInfo {
	if _debd.Numbering.X() == nil {
		return listItemInfo{}
	}
	if len(_debd.Numbering.Definitions()) < 1 {
		return listItemInfo{}
	}
	_dccc := _gfag(_bebbf)
	if _dccc == nil {
		return listItemInfo{}
	}
	_gddc := _debd.GetNumberingLevelByIds(_dccc.NumId.ValAttr, _dccc.Ilvl.ValAttr)
	if _dddb := _gddc.X(); _dddb == nil {
		return listItemInfo{}
	}
	_fdbb := int64(0)
	for _, _acbf := range _debd.Numbering._gebaa.Num {
		if _acbf != nil && _acbf.NumIdAttr == _dccc.NumId.ValAttr {
			_fdbb = _acbf.AbstractNumId.ValAttr
		}
	}
	return listItemInfo{FromParagraph: &_bebbf, AbstractNumId: &_fdbb, NumberingLevel: &_gddc}
}

// Styles returns all styles.
func (_ccaggf Styles) Styles() []Style {
	_fcad := []Style{}
	for _, _abfd := range _ccaggf._cdgg.Style {
		_fcad = append(_fcad, Style{_abfd})
	}
	return _fcad
}

// SetAllowOverlapAttr sets the allowOverlap attribute of anchor.
func (_bag AnchoredDrawing) SetAllowOverlapAttr(val bool) { _bag._cg.AllowOverlapAttr = val }

// NewStyles constructs a new empty Styles
func NewStyles() Styles { return Styles{_db.NewStyles()} }

// X returns the internally wrapped *wml.CT_SectPr.
func (_cdgee Section) X() *_db.CT_SectPr { return _cdgee._ggcfb }

// Footnote returns the footnote based on the ID; this can be used nicely with
// the run.IsFootnote() functionality.
func (_bac *Document) Footnote(id int64) Footnote {
	for _, _fcfce := range _bac.Footnotes() {
		if _fcfce.id() == id {
			return _fcfce
		}
	}
	return Footnote{}
}

// AddWatermarkText adds new watermark text to the document.
func (_cagda *Document) AddWatermarkText(text string) WatermarkText {
	var _cabb []Header
	if _bfd, _bfeg := _cagda.BodySection().GetHeader(_db.ST_HdrFtrDefault); _bfeg {
		_cabb = append(_cabb, _bfd)
	}
	if _afcd, _deeg := _cagda.BodySection().GetHeader(_db.ST_HdrFtrEven); _deeg {
		_cabb = append(_cabb, _afcd)
	}
	if _ddbb, _bfge := _cagda.BodySection().GetHeader(_db.ST_HdrFtrFirst); _bfge {
		_cabb = append(_cabb, _ddbb)
	}
	if len(_cabb) < 1 {
		_daec := _cagda.AddHeader()
		_cagda.BodySection().SetHeader(_daec, _db.ST_HdrFtrDefault)
		_cabb = append(_cabb, _daec)
	}
	_ecdb := NewWatermarkText()
	for _, _beee := range _cabb {
		_cfc := _beee.Paragraphs()
		if len(_cfc) < 1 {
			_dea := _beee.AddParagraph()
			_dea.AddRun().AddText("")
		}
		for _, _dafb := range _beee.X().EG_ContentBlockContent {
			for _, _ecfe := range _dafb.P {
				for _, _fgga := range _ecfe.EG_PContent {
					for _, _eae := range _fgga.EG_ContentRunContent {
						if _eae.R == nil {
							continue
						}
						for _, _dcac := range _eae.R.EG_RunInnerContent {
							_dcac.Pict = _ecdb._gdgad
							break
						}
					}
				}
			}
		}
	}
	_ecdb.SetText(text)
	return _ecdb
}

// Italic returns true if run font is italic.
func (_cfae RunProperties) Italic() bool {
	_caga := _cfae._cadb
	return _ebdd(_caga.I) || _ebdd(_caga.ICs)
}

// SetSize sets the size of the displayed image on the page.
func (_eaec InlineDrawing) SetSize(w, h _bc.Distance) {
	_eaec._adb.Extent.CxAttr = int64(float64(w*_bc.Pixel72) / _bc.EMU)
	_eaec._adb.Extent.CyAttr = int64(float64(h*_bc.Pixel72) / _bc.EMU)
}

// X returns the inner wrapped XML type.
func (_edfbg HyperLink) X() *_db.CT_Hyperlink { return _edfbg._cfcf }

// Validate validates the structure and in cases where it't possible, the ranges
// of elements within a document. A validation error dones't mean that the
// document won't work in MS Word or LibreOffice, but it's worth checking into.
func (_cde *Document) Validate() error {
	if _cde == nil || _cde._adcb == nil {
		return _bg.New("\u0064o\u0063\u0075m\u0065\u006e\u0074\u0020n\u006f\u0074\u0020i\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0065d \u0063\u006f\u0072r\u0065\u0063t\u006c\u0079\u002c\u0020\u006e\u0069l\u0020\u0062a\u0073\u0065")
	}
	for _, _gcde := range []func() error{_cde.validateTableCells, _cde.validateBookmarks} {
		if _feeg := _gcde(); _feeg != nil {
			return _feeg
		}
	}
	if _bafd := _cde._adcb.Validate(); _bafd != nil {
		return _bafd
	}
	return nil
}
func (_bddca Paragraph) ensurePPr() {
	if _bddca._dcfbd.PPr == nil {
		_bddca._dcfbd.PPr = _db.NewCT_PPr()
	}
}

// SaveToFile writes the document out to a file.
func (_ead *Document) SaveToFile(path string) error {
	_eada, _fec := _fe.Create(path)
	if _fec != nil {
		return _fec
	}
	defer _eada.Close()
	return _ead.Save(_eada)
}

// SetBehindDoc sets the behindDoc attribute of anchor.
func (_ga AnchoredDrawing) SetBehindDoc(val bool) { _ga._cg.BehindDocAttr = val }

// SetHeight allows controlling the height of a row within a table.
func (_fedacb RowProperties) SetHeight(ht _bc.Distance, rule _db.ST_HeightRule) {
	if rule == _db.ST_HeightRuleUnset {
		_fedacb._caaf.TrHeight = nil
	} else {
		_ebbc := _db.NewCT_Height()
		_ebbc.HRuleAttr = rule
		_ebbc.ValAttr = &_ca.ST_TwipsMeasure{}
		_ebbc.ValAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(ht / _bc.Twips))
		_fedacb._caaf.TrHeight = []*_db.CT_Height{_ebbc}
	}
}
func _ffgfd(_cfdce *_db.CT_P, _adcg *_db.CT_Hyperlink, _dcfa *TableInfo, _defeb *DrawingInfo, _fbgbg []*_db.EG_PContent) []TextItem {
	if len(_fbgbg) == 0 {
		return []TextItem{TextItem{Text: "", DrawingInfo: _defeb, Paragraph: _cfdce, Hyperlink: _adcg, Run: nil, TableInfo: _dcfa}}
	}
	_cdfd := []TextItem{}
	for _, _gdaf := range _fbgbg {
		for _, _gacc := range _gdaf.FldSimple {
			if _gacc != nil {
				_cdfd = append(_cdfd, _ffgfd(_cfdce, _adcg, _dcfa, _defeb, _gacc.EG_PContent)...)
			}
		}
		if _dbcg := _gdaf.Hyperlink; _dbcg != nil {
			_cdfd = append(_cdfd, _cbfe(_cfdce, _dbcg, _dcfa, _defeb, _dbcg.EG_ContentRunContent)...)
		}
		_cdfd = append(_cdfd, _cbfe(_cfdce, nil, _dcfa, _defeb, _gdaf.EG_ContentRunContent)...)
	}
	return _cdfd
}

// SetInsideHorizontal sets the interior horizontal borders to a specified type, color and thickness.
func (_fcff TableBorders) SetInsideHorizontal(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_fcff._fcdb.InsideH = _db.NewCT_Border()
	_ebca(_fcff._fcdb.InsideH, t, c, thickness)
}

// TableProperties are the properties for a table within a document
type TableProperties struct{ _faec *_db.CT_TblPr }

// SetPossibleValues sets possible values for a FormFieldTypeDropDown.
func (_fdadd FormField) SetPossibleValues(values []string) {
	if _fdadd._fbcee.DdList != nil {
		for _, _cbfg := range values {
			_ddgg := _db.NewCT_String()
			_ddgg.ValAttr = _cbfg
			_fdadd._fbcee.DdList.ListEntry = append(_fdadd._fbcee.DdList.ListEntry, _ddgg)
		}
	}
}

// Name returns the name of the bookmark whcih is the document unique ID that
// identifies the bookmark.
func (_bbf Bookmark) Name() string { return _bbf._ac.NameAttr }

// SetRowBandSize sets the number of Rows in the row band
func (_ggggg TableStyleProperties) SetRowBandSize(rows int64) {
	_ggggg._dagce.TblStyleRowBandSize = _db.NewCT_DecimalNumber()
	_ggggg._dagce.TblStyleRowBandSize.ValAttr = rows
}

// Underline returns the type of run underline.
func (_ecgce RunProperties) Underline() _db.ST_Underline {
	if _eeagge := _ecgce._cadb.U; _eeagge != nil {
		return _eeagge.ValAttr
	}
	return 0
}

// GetHighlight returns the HighlightColor.
func (_bdagf RunProperties) GetHighlight() _db.ST_HighlightColor {
	if _bdagf._cadb.Highlight != nil {
		return _bdagf._cadb.Highlight.ValAttr
	}
	return _db.ST_HighlightColorNone
}

// AddParagraph adds a paragraph to the footnote.
func (_aefbc Footnote) AddParagraph() Paragraph {
	_cecde := _db.NewEG_ContentBlockContent()
	_ccagg := len(_aefbc._cefg.EG_BlockLevelElts[0].EG_ContentBlockContent)
	_aefbc._cefg.EG_BlockLevelElts[0].EG_ContentBlockContent = append(_aefbc._cefg.EG_BlockLevelElts[0].EG_ContentBlockContent, _cecde)
	_bfdag := _db.NewCT_P()
	var _becae *_db.CT_String
	if _ccagg != 0 {
		_ddba := len(_aefbc._cefg.EG_BlockLevelElts[0].EG_ContentBlockContent[_ccagg-1].P)
		_becae = _aefbc._cefg.EG_BlockLevelElts[0].EG_ContentBlockContent[_ccagg-1].P[_ddba-1].PPr.PStyle
	} else {
		_becae = _db.NewCT_String()
		_becae.ValAttr = "\u0046\u006f\u006f\u0074\u006e\u006f\u0074\u0065"
	}
	_cecde.P = append(_cecde.P, _bfdag)
	_egad := Paragraph{_aefbc._abca, _bfdag}
	_egad._dcfbd.PPr = _db.NewCT_PPr()
	_egad._dcfbd.PPr.PStyle = _becae
	_egad._dcfbd.PPr.RPr = _db.NewCT_ParaRPr()
	return _egad
}

// SetStartPct sets the cell start margin
func (_bfa CellMargins) SetStartPct(pct float64) {
	_bfa._caff.Start = _db.NewCT_TblWidth()
	_beb(_bfa._caff.Start, pct)
}

// Font returns the name of run font family.
func (_dedf RunProperties) Font() string {
	if _agcc := _dedf._cadb.RFonts; _agcc != nil {
		if _agcc.AsciiAttr != nil {
			return *_agcc.AsciiAttr
		} else if _agcc.HAnsiAttr != nil {
			return *_agcc.HAnsiAttr
		} else if _agcc.CsAttr != nil {
			return *_agcc.CsAttr
		}
	}
	return ""
}
func (_dacg *Document) save(_daf _fg.Writer, _dgd string) error {
	const _cbc = "\u0064o\u0063u\u006d\u0065\u006e\u0074\u003a\u0064\u002e\u0053\u0061\u0076\u0065"
	if _ecd := _dacg._adcb.Validate(); _ecd != nil {
		_c.Log.Warning("\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u0020i\u006e\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0073", _ecd)
	}
	_cff := _b.DocTypeDocument
	// if !_fd.GetLicenseKey().IsLicensed() && !_fbgb {
	// 	_g.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
	// 	_g.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
	// 	return _bg.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	// }
	// if len(_dacg._afc) == 0 {
	// 	if len(_dgd) > 0 {
	// 		_dacg._afc = _dgd
	// 	} else {
	// 		_def, _cba := _fd.GenRefId("\u0064\u0077")
	// 		if _cba != nil {
	// 			_c.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cba)
	// 			return _cba
	// 		}
	// 		_dacg._afc = _def
	// 	}
	// }
	// if _ggba := _fd.Track(_dacg._afc, _cbc); _ggba != nil {
	// 	_c.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _ggba)
	// 	return _ggba
	// }
	_fggb := _a.NewWriter(_daf)
	defer _fggb.Close()
	if _bad := _ed.MarshalXML(_fggb, _b.BaseRelsFilename, _dacg.Rels.X()); _bad != nil {
		return _bad
	}
	if _bdb := _ed.MarshalXMLByType(_fggb, _cff, _b.ExtendedPropertiesType, _dacg.AppProperties.X()); _bdb != nil {
		return _bdb
	}
	if _fbdd := _ed.MarshalXMLByType(_fggb, _cff, _b.CorePropertiesType, _dacg.CoreProperties.X()); _fbdd != nil {
		return _fbdd
	}
	if _dacg.CustomProperties.X() != nil {
		if _agc := _ed.MarshalXMLByType(_fggb, _cff, _b.CustomPropertiesType, _dacg.CustomProperties.X()); _agc != nil {
			return _agc
		}
	}
	if _dacg.Thumbnail != nil {
		_eeaf, _acb := _fggb.Create("\u0064\u006f\u0063Pr\u006f\u0070\u0073\u002f\u0074\u0068\u0075\u006d\u0062\u006e\u0061\u0069\u006c\u002e\u006a\u0070\u0065\u0067")
		if _acb != nil {
			return _acb
		}
		if _afd := _fa.Encode(_eeaf, _dacg.Thumbnail, nil); _afd != nil {
			return _afd
		}
	}
	if _dae := _ed.MarshalXMLByType(_fggb, _cff, _b.SettingsType, _dacg.Settings.X()); _dae != nil {
		return _dae
	}
	_ade := _b.AbsoluteFilename(_cff, _b.OfficeDocumentType, 0)
	if _dfg := _ed.MarshalXML(_fggb, _ade, _dacg._adcb); _dfg != nil {
		return _dfg
	}
	if _feg := _ed.MarshalXML(_fggb, _ed.RelationsPathFor(_ade), _dacg._dfc.X()); _feg != nil {
		return _feg
	}
	if _dacg.Numbering.X() != nil {
		if _cef := _ed.MarshalXMLByType(_fggb, _cff, _b.NumberingType, _dacg.Numbering.X()); _cef != nil {
			return _cef
		}
	}
	if _dabd := _ed.MarshalXMLByType(_fggb, _cff, _b.StylesType, _dacg.Styles.X()); _dabd != nil {
		return _dabd
	}
	if _dacg._eaf != nil {
		if _cdfg := _ed.MarshalXMLByType(_fggb, _cff, _b.WebSettingsType, _dacg._eaf); _cdfg != nil {
			return _cdfg
		}
	}
	if _dacg._begg != nil {
		if _efd := _ed.MarshalXMLByType(_fggb, _cff, _b.FontTableType, _dacg._begg); _efd != nil {
			return _efd
		}
	}
	if _dacg._acdc != nil {
		if _bda := _ed.MarshalXMLByType(_fggb, _cff, _b.EndNotesType, _dacg._acdc); _bda != nil {
			return _bda
		}
	}
	if _dacg._ffcef != nil {
		if _fedf := _ed.MarshalXMLByType(_fggb, _cff, _b.FootNotesType, _dacg._ffcef); _fedf != nil {
			return _fedf
		}
	}
	for _dfe, _dcc := range _dacg._cdf {
		if _edaf := _ed.MarshalXMLByTypeIndex(_fggb, _cff, _b.ThemeType, _dfe+1, _dcc); _edaf != nil {
			return _edaf
		}
	}
	for _ffdf, _faf := range _dacg._aede {
		_ebf, _cbce := _faf.ExportToByteArray()
		if _cbce != nil {
			return _cbce
		}
		_cgd := "\u0077\u006f\u0072d\u002f" + _faf.TargetAttr[:len(_faf.TargetAttr)-4] + "\u002e\u0062\u0069\u006e"
		if _gae := _ed.AddFileFromBytes(_fggb, _cgd, _ebf); _gae != nil {
			return _gae
		}
		if _ebff := _ed.MarshalXMLByTypeIndex(_fggb, _cff, _b.ControlType, _ffdf+1, _faf.Ocx); _ebff != nil {
			return _ebff
		}
	}
	for _deb, _aad := range _dacg._adf {
		_dbaa := _b.AbsoluteFilename(_cff, _b.HeaderType, _deb+1)
		if _dcd := _ed.MarshalXML(_fggb, _dbaa, _aad); _dcd != nil {
			return _dcd
		}
		if !_dacg._gcdc[_deb].IsEmpty() {
			_ed.MarshalXML(_fggb, _ed.RelationsPathFor(_dbaa), _dacg._gcdc[_deb].X())
		}
	}
	for _ged, _gcg := range _dacg._abc {
		_defd := _b.AbsoluteFilename(_cff, _b.FooterType, _ged+1)
		if _fdb := _ed.MarshalXMLByTypeIndex(_fggb, _cff, _b.FooterType, _ged+1, _gcg); _fdb != nil {
			return _fdb
		}
		if !_dacg._adg[_ged].IsEmpty() {
			_ed.MarshalXML(_fggb, _ed.RelationsPathFor(_defd), _dacg._adg[_ged].X())
		}
	}
	for _egba, _bdee := range _dacg.Images {
		if _egf := _ffa.AddImageToZip(_fggb, _bdee, _egba+1, _b.DocTypeDocument); _egf != nil {
			return _egf
		}
	}
	for _acc, _gadb := range _dacg._beag {
		_cbgc := _b.AbsoluteFilename(_cff, _b.ChartType, _acc+1)
		_ed.MarshalXML(_fggb, _cbgc, _gadb._bbd)
	}
	if _gba := _ed.MarshalXML(_fggb, _b.ContentTypesFilename, _dacg.ContentTypes.X()); _gba != nil {
		return _gba
	}
	if _fge := _dacg.WriteExtraFiles(_fggb); _fge != nil {
		return _fge
	}
	return _fggb.Close()
}

// MergeFields returns the list of all mail merge fields found in the document.
func (_cgcbd Document) MergeFields() []string {
	_fefc := map[string]struct{}{}
	for _, _ebddb := range _cgcbd.mergeFields() {
		_fefc[_ebddb._gcfd] = struct{}{}
	}
	_febdf := []string{}
	for _fdeb := range _fefc {
		_febdf = append(_febdf, _fdeb)
	}
	return _febdf
}

// Tables returns the tables defined in the footer.
func (_ccfa Footer) Tables() []Table {
	_aebec := []Table{}
	if _ccfa._eeed == nil {
		return nil
	}
	for _, _gegd := range _ccfa._eeed.EG_ContentBlockContent {
		for _, _cgdc := range _ccfa._bcge.tables(_gegd) {
			_aebec = append(_aebec, _cgdc)
		}
	}
	return _aebec
}

// SetColor sets the text color.
func (_ccdc RunProperties) SetColor(c _bd.Color) {
	_ccdc._cadb.Color = _db.NewCT_Color()
	_ccdc._cadb.Color.ValAttr.ST_HexColorRGB = c.AsRGBString()
}

// UnderlineColor returns the hex color value of paragraph underline.
func (_bgcbc ParagraphProperties) UnderlineColor() string {
	if _cgba := _bgcbc._aedc.RPr.U; _cgba != nil {
		_ggec := _cgba.ColorAttr
		if _ggec != nil && _ggec.ST_HexColorRGB != nil {
			return *_ggec.ST_HexColorRGB
		}
	}
	return ""
}

// SizeValue returns the value of paragraph font size in points.
func (_geeb ParagraphProperties) SizeValue() float64 {
	if _gaebb := _geeb._aedc.RPr.Sz; _gaebb != nil {
		_bfbcg := _gaebb.ValAttr
		if _bfbcg.ST_UnsignedDecimalNumber != nil {
			return float64(*_bfbcg.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// SetValue sets the value of a FormFieldTypeText or FormFieldTypeDropDown. For
// FormFieldTypeDropDown, the value must be one of the fields possible values.
func (_ggee FormField) SetValue(v string) {
	if _ggee._fbcee.DdList != nil {
		for _fecd, _ccaf := range _ggee.PossibleValues() {
			if _ccaf == v {
				_ggee._fbcee.DdList.Result = _db.NewCT_DecimalNumber()
				_ggee._fbcee.DdList.Result.ValAttr = int64(_fecd)
				break
			}
		}
	} else if _ggee._fbcee.TextInput != nil {
		_ggee._efgb.T = _db.NewCT_Text()
		_ggee._efgb.T.Content = v
	}
}

// RemoveParagraph removes a paragraph from a footer.
func (_cccge Footer) RemoveParagraph(p Paragraph) {
	for _, _bfdb := range _cccge._eeed.EG_ContentBlockContent {
		for _cggc, _cgca := range _bfdb.P {
			if _cgca == p._dcfbd {
				copy(_bfdb.P[_cggc:], _bfdb.P[_cggc+1:])
				_bfdb.P = _bfdb.P[0 : len(_bfdb.P)-1]
				return
			}
		}
	}
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_ebfed Header) AddImage(i _ffa.Image) (_ffa.ImageRef, error) {
	var _edaab _ffa.Relationships
	for _fegb, _agaba := range _ebfed._bgac._adf {
		if _agaba == _ebfed._ecead {
			_edaab = _ebfed._bgac._gcdc[_fegb]
		}
	}
	_fgcag := _ffa.MakeImageRef(i, &_ebfed._bgac.DocBase, _edaab)
	if i.Data == nil && i.Path == "" {
		return _fgcag, _bg.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _fgcag, _bg.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _fgcag, _bg.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	_ebfed._bgac.Images = append(_ebfed._bgac.Images, _fgcag)
	_aacff := _g.Sprintf("\u006d\u0065d\u0069\u0061\u002fi\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", len(_ebfed._bgac.Images), i.Format)
	_geae := _edaab.AddRelationship(_aacff, _b.ImageType)
	_fgcag.SetRelID(_geae.X().IdAttr)
	return _fgcag, nil
}

// SetName sets the name of the style.
func (_edaae Style) SetName(name string) {
	_edaae._eedcd.Name = _db.NewCT_String()
	_edaae._eedcd.Name.ValAttr = name
}

// Paragraph is a paragraph within a document.
type Paragraph struct {
	_cfge  *Document
	_dcfbd *_db.CT_P
}

// Themes returns document's themes.
func (_dffc *Document) Themes() []*_aa.Theme { return _dffc._cdf }

// SetTextWrapNone unsets text wrapping so the image can float on top of the
// text. When used in conjunction with X/Y Offset relative to the page it can be
// used to place a logo at the top of a page at an absolute position that
// doesn't interfere with text.
func (_gea AnchoredDrawing) SetTextWrapNone() {
	_gea._cg.Choice = &_db.WdEG_WrapTypeChoice{}
	_gea._cg.Choice.WrapNone = _db.NewWdCT_WrapNone()
}

// SetStrikeThrough sets the run to strike-through.
func (_gefb RunProperties) SetStrikeThrough(b bool) {
	if !b {
		_gefb._cadb.Strike = nil
	} else {
		_gefb._cadb.Strike = _db.NewCT_OnOff()
	}
}

// ParagraphProperties are the properties for a paragraph.
type ParagraphProperties struct {
	_efedg *Document
	_aedc  *_db.CT_PPr
}

const (
	FormFieldTypeUnknown FormFieldType = iota
	FormFieldTypeText
	FormFieldTypeCheckBox
	FormFieldTypeDropDown
)

// GetKerning returns the kerning (character spacing) of a run
func (_ggde RunProperties) GetKerning() _bc.Distance {
	if _ggde._cadb.Kern != nil {
		return _bc.Distance(float64(*_ggde._cadb.Kern.ValAttr.ST_UnsignedDecimalNumber) * _bc.HalfPoint)
	}
	return 0
}

// Paragraphs returns the paragraphs defined in a header.
func (_fabba Header) Paragraphs() []Paragraph {
	_gafb := []Paragraph{}
	for _, _faged := range _fabba._ecead.EG_ContentBlockContent {
		for _, _dgga := range _faged.P {
			_gafb = append(_gafb, Paragraph{_fabba._bgac, _dgga})
		}
	}
	for _, _ggfe := range _fabba.Tables() {
		for _, _ccceg := range _ggfe.Rows() {
			for _, _ccdg := range _ccceg.Cells() {
				_gafb = append(_gafb, _ccdg.Paragraphs()...)
			}
		}
	}
	return _gafb
}

// SetBottom sets the bottom border to a specified type, color and thickness.
func (_ggb CellBorders) SetBottom(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_ggb._bee.Bottom = _db.NewCT_Border()
	_ebca(_ggb._bee.Bottom, t, c, thickness)
}

// X returns the inner wrapped XML type.
func (_gbace TableWidth) X() *_db.CT_TblWidth { return _gbace._bagbe }

// RunProperties returns the run style properties.
func (_gbbgc Style) RunProperties() RunProperties {
	if _gbbgc._eedcd.RPr == nil {
		_gbbgc._eedcd.RPr = _db.NewCT_RPr()
	}
	return RunProperties{_gbbgc._eedcd.RPr}
}

// Name returns the name of the field.
func (_faeb FormField) Name() string { return *_faeb._fbcee.Name[0].ValAttr }
func _efdfc(_edbd *_db.EG_ContentBlockContent) []Bookmark {
	_gfbdf := []Bookmark{}
	for _, _eefb := range _edbd.P {
		for _, _cdcca := range _eefb.EG_PContent {
			for _, _ecae := range _cdcca.EG_ContentRunContent {
				for _, _aafd := range _ecae.EG_RunLevelElts {
					for _, _cafag := range _aafd.EG_RangeMarkupElements {
						if _cafag.BookmarkStart != nil {
							_gfbdf = append(_gfbdf, Bookmark{_cafag.BookmarkStart})
						}
					}
				}
			}
		}
	}
	for _, _bca := range _edbd.EG_RunLevelElts {
		for _, _caae := range _bca.EG_RangeMarkupElements {
			if _caae.BookmarkStart != nil {
				_gfbdf = append(_gfbdf, Bookmark{_caae.BookmarkStart})
			}
		}
	}
	for _, _efde := range _edbd.Tbl {
		for _, _dgbg := range _efde.EG_ContentRowContent {
			for _, _cfeg := range _dgbg.Tr {
				for _, _feca := range _cfeg.EG_ContentCellContent {
					for _, _dfge := range _feca.Tc {
						for _, _cefa := range _dfge.EG_BlockLevelElts {
							for _, _afgf := range _cefa.EG_ContentBlockContent {
								for _, _bbcf := range _efdfc(_afgf) {
									_gfbdf = append(_gfbdf, _bbcf)
								}
							}
						}
					}
				}
			}
		}
	}
	return _gfbdf
}

// MultiLevelType returns the multilevel type, or ST_MultiLevelTypeUnset if not set.
func (_fdga NumberingDefinition) MultiLevelType() _db.ST_MultiLevelType {
	if _fdga._bbgd.MultiLevelType != nil {
		return _fdga._bbgd.MultiLevelType.ValAttr
	} else {
		return _db.ST_MultiLevelTypeUnset
	}
}

// DocText is an array of extracted text items which has some methods for representing extracted text.
type DocText struct {
	Items []TextItem
	_bbaf []listItemInfo
	_bgfd map[int64]map[int64]int64
}

// SetUnhideWhenUsed controls if a semi hidden style becomes visible when used.
func (_affaa Style) SetUnhideWhenUsed(b bool) {
	if b {
		_affaa._eedcd.UnhideWhenUsed = _db.NewCT_OnOff()
	} else {
		_affaa._eedcd.UnhideWhenUsed = nil
	}
}
func (_adeb Footnote) content() []*_db.EG_ContentBlockContent {
	var _aaae []*_db.EG_ContentBlockContent
	for _, _geaf := range _adeb._cefg.EG_BlockLevelElts {
		_aaae = append(_aaae, _geaf.EG_ContentBlockContent...)
	}
	return _aaae
}

// AddTable adds a table to the table cell.
func (_fdd Cell) AddTable() Table {
	_gdd := _db.NewEG_BlockLevelElts()
	_fdd._caf.EG_BlockLevelElts = append(_fdd._caf.EG_BlockLevelElts, _gdd)
	_gec := _db.NewEG_ContentBlockContent()
	_gdd.EG_ContentBlockContent = append(_gdd.EG_ContentBlockContent, _gec)
	_gdb := _db.NewCT_Tbl()
	_gec.Tbl = append(_gec.Tbl, _gdb)
	return Table{_fdd._cec, _gdb}
}

// InitializeDefault constructs the default styles.
func (_daaff Styles) InitializeDefault() {
	_daaff.initializeDocDefaults()
	_daaff.initializeStyleDefaults()
}

// Section return paragraph properties section value.
func (_gbdab ParagraphProperties) Section() (Section, bool) {
	if _gbdab._aedc.SectPr != nil {
		return Section{_gbdab._efedg, _gbdab._aedc.SectPr}, true
	}
	return Section{}, false
}

// Name returns the name of the style if set.
func (_ggeae Style) Name() string {
	if _ggeae._eedcd.Name == nil {
		return ""
	}
	return _ggeae._eedcd.Name.ValAttr
}

// SetAfterLineSpacing sets spacing below paragraph in line units.
func (_abffb Paragraph) SetAfterLineSpacing(d _bc.Distance) {
	_abffb.ensurePPr()
	if _abffb._dcfbd.PPr.Spacing == nil {
		_abffb._dcfbd.PPr.Spacing = _db.NewCT_Spacing()
	}
	_ebeda := _abffb._dcfbd.PPr.Spacing
	_ebeda.AfterLinesAttr = _b.Int64(int64(d / _bc.Twips))
}

// Copy makes a deep copy of the document by saving and reading it back.
// It can be useful to avoid sharing common data between two documents.
func (_bbeg *Document) Copy() (*Document, error) {
	_aaab := _fed.NewBuffer([]byte{})
	_bafg := _bbeg.save(_aaab, _bbeg._afc)
	if _bafg != nil {
		return nil, _bafg
	}
	_fcg := _aaab.Bytes()
	_baff := _fed.NewReader(_fcg)
	return _gab(_baff, int64(_baff.Len()), _bbeg._afc)
}

// SearchStyleByName return style by its name.
func (_dbbf Styles) SearchStyleByName(name string) (Style, bool) {
	for _, _ebde := range _dbbf._cdgg.Style {
		if _ebde.Name != nil {
			if _ebde.Name.ValAttr == name {
				return Style{_ebde}, true
			}
		}
	}
	return Style{}, false
}

// SetPageSizeAndOrientation sets the page size and orientation for a section.
func (_gbea Section) SetPageSizeAndOrientation(w, h _bc.Distance, orientation _db.ST_PageOrientation) {
	if _gbea._ggcfb.PgSz == nil {
		_gbea._ggcfb.PgSz = _db.NewCT_PageSz()
	}
	_gbea._ggcfb.PgSz.OrientAttr = orientation
	if orientation == _db.ST_PageOrientationLandscape {
		_gbea._ggcfb.PgSz.WAttr = &_ca.ST_TwipsMeasure{}
		_gbea._ggcfb.PgSz.WAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(h / _bc.Twips))
		_gbea._ggcfb.PgSz.HAttr = &_ca.ST_TwipsMeasure{}
		_gbea._ggcfb.PgSz.HAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(w / _bc.Twips))
	} else {
		_gbea._ggcfb.PgSz.WAttr = &_ca.ST_TwipsMeasure{}
		_gbea._ggcfb.PgSz.WAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(w / _bc.Twips))
		_gbea._ggcfb.PgSz.HAttr = &_ca.ST_TwipsMeasure{}
		_gbea._ggcfb.PgSz.HAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(h / _bc.Twips))
	}
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (_ddace ParagraphStyleProperties) SetKeepOnOnePage(b bool) {
	if !b {
		_ddace._fagf.KeepLines = nil
	} else {
		_ddace._fagf.KeepLines = _db.NewCT_OnOff()
	}
}
func _fced(_fcfde *_db.CT_Tbl, _ccfb map[string]string) {
	for _, _edad := range _fcfde.EG_ContentRowContent {
		for _, _fbgaf := range _edad.Tr {
			for _, _ffcb := range _fbgaf.EG_ContentCellContent {
				for _, _befd := range _ffcb.Tc {
					for _, _fgca := range _befd.EG_BlockLevelElts {
						for _, _gdgbb := range _fgca.EG_ContentBlockContent {
							for _, _efdfg := range _gdgbb.P {
								_ffgb(_efdfg, _ccfb)
							}
							for _, _febg := range _gdgbb.Tbl {
								_fced(_febg, _ccfb)
							}
						}
					}
				}
			}
		}
	}
}

// Properties returns the row properties.
func (_cdabf Row) Properties() RowProperties {
	if _cdabf._bageg.TrPr == nil {
		_cdabf._bageg.TrPr = _db.NewCT_TrPr()
	}
	return RowProperties{_cdabf._bageg.TrPr}
}

// Properties returns the table properties.
func (_edaef Table) Properties() TableProperties {
	if _edaef._cgffe.TblPr == nil {
		_edaef._cgffe.TblPr = _db.NewCT_TblPr()
	}
	return TableProperties{_edaef._cgffe.TblPr}
}

// SetShadow sets the run to shadowed text.
func (_gaaad RunProperties) SetShadow(b bool) {
	if !b {
		_gaaad._cadb.Shadow = nil
	} else {
		_gaaad._cadb.Shadow = _db.NewCT_OnOff()
	}
}

// Section is the beginning of a new section.
type Section struct {
	_bagegb *Document
	_ggcfb  *_db.CT_SectPr
}

// SetOutlineLevel sets the outline level of this style.
func (_acgf ParagraphStyleProperties) SetOutlineLevel(lvl int) {
	_acgf._fagf.OutlineLvl = _db.NewCT_DecimalNumber()
	_acgf._fagf.OutlineLvl.ValAttr = int64(lvl)
}

// Footnotes returns the footnotes defined in the document.
func (_ccbc *Document) Footnotes() []Footnote {
	_ccgd := []Footnote{}
	for _, _effd := range _ccbc._ffcef.CT_Footnotes.Footnote {
		_ccgd = append(_ccgd, Footnote{_ccbc, _effd})
	}
	return _ccgd
}

// SetLinkedStyle sets the style that this style is linked to.
func (_cffce Style) SetLinkedStyle(name string) {
	if name == "" {
		_cffce._eedcd.Link = nil
	} else {
		_cffce._eedcd.Link = _db.NewCT_String()
		_cffce._eedcd.Link.ValAttr = name
	}
}

// Save writes the document to an io.Writer in the Zip package format.
func (_fba *Document) Save(w _fg.Writer) error { return _fba.save(w, _fba._afc) }

// SetKerning sets the run's font kerning.
func (_eegdg RunProperties) SetKerning(size _bc.Distance) {
	_eegdg._cadb.Kern = _db.NewCT_HpsMeasure()
	_eegdg._cadb.Kern.ValAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(size / _bc.HalfPoint))
}

// CellProperties are a table cells properties within a document.
type CellProperties struct{ _beea *_db.CT_TcPr }

// SetPicture sets the watermark picture.
func (_dbebe *WatermarkPicture) SetPicture(imageRef _ffa.ImageRef) {
	_cbdbd := imageRef.RelID()
	_afcgc := _dbebe.getShape()
	if _dbebe._febge != nil {
		_agbda := _dbebe._febge.EG_ShapeElements
		if len(_agbda) > 0 && _agbda[0].Imagedata != nil {
			_agbda[0].Imagedata.IdAttr = &_cbdbd
		}
	} else {
		_fcbb := _dbebe.findNode(_afcgc, "\u0069m\u0061\u0067\u0065\u0064\u0061\u0074a")
		for _aeeef, _abgcb := range _fcbb.Attrs {
			if _abgcb.Name.Local == "\u0069\u0064" {
				_fcbb.Attrs[_aeeef].Value = _cbdbd
			}
		}
	}
}
func _aaffd() *_eb.Formulas {
	_eeeef := _eb.NewFormulas()
	_eeeef.F = []*_eb.CT_F{_ge.CreateFormula("\u0069\u0066 \u006c\u0069\u006e\u0065\u0044\u0072\u0061\u0077\u006e\u0020\u0070\u0069\u0078\u0065\u006c\u004c\u0069\u006e\u0065\u0057\u0069\u0064th\u0020\u0030"), _ge.CreateFormula("\u0073\u0075\u006d\u0020\u0040\u0030\u0020\u0031\u0020\u0030"), _ge.CreateFormula("\u0073\u0075\u006d\u0020\u0030\u0020\u0030\u0020\u0040\u0031"), _ge.CreateFormula("p\u0072\u006f\u0064\u0020\u0040\u0032\u0020\u0031\u0020\u0032"), _ge.CreateFormula("\u0070r\u006f\u0064\u0020\u0040\u0033\u0020\u0032\u0031\u0036\u0030\u0030 \u0070\u0069\u0078\u0065\u006c\u0057\u0069\u0064\u0074\u0068"), _ge.CreateFormula("\u0070r\u006f\u0064\u0020\u00403\u0020\u0032\u0031\u0036\u00300\u0020p\u0069x\u0065\u006c\u0048\u0065\u0069\u0067\u0068t"), _ge.CreateFormula("\u0073\u0075\u006d\u0020\u0040\u0030\u0020\u0030\u0020\u0031"), _ge.CreateFormula("p\u0072\u006f\u0064\u0020\u0040\u0036\u0020\u0031\u0020\u0032"), _ge.CreateFormula("\u0070r\u006f\u0064\u0020\u0040\u0037\u0020\u0032\u0031\u0036\u0030\u0030 \u0070\u0069\u0078\u0065\u006c\u0057\u0069\u0064\u0074\u0068"), _ge.CreateFormula("\u0073\u0075\u006d\u0020\u0040\u0038\u0020\u0032\u00316\u0030\u0030\u0020\u0030"), _ge.CreateFormula("\u0070r\u006f\u0064\u0020\u00407\u0020\u0032\u0031\u0036\u00300\u0020p\u0069x\u0065\u006c\u0048\u0065\u0069\u0067\u0068t"), _ge.CreateFormula("\u0073u\u006d \u0040\u0031\u0030\u0020\u0032\u0031\u0036\u0030\u0030\u0020\u0030")}
	return _eeeef
}
func (_ggce *WatermarkPicture) getShape() *_b.XSDAny {
	return _ggce.getInnerElement("\u0073\u0068\u0061p\u0065")
}
func (_caccg *WatermarkPicture) getShapeImagedata() *_b.XSDAny {
	return _caccg.getInnerElement("\u0069m\u0061\u0067\u0065\u0064\u0061\u0074a")
}

// Color controls the run or styles color.
type Color struct{ _fcd *_db.CT_Color }

// AddHeader creates a header associated with the document, but doesn't add it
// to the document for display.
func (_gage *Document) AddHeader() Header {
	_cbda := _db.NewHdr()
	_gage._adf = append(_gage._adf, _cbda)
	_ada := _g.Sprintf("\u0068\u0065\u0061d\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", len(_gage._adf))
	_gage._dfc.AddRelationship(_ada, _b.HeaderType)
	_gage.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_ada, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0068\u0065\u0061\u0064e\u0072\u002b\u0078\u006d\u006c")
	_gage._gcdc = append(_gage._gcdc, _ffa.NewRelationships())
	return Header{_gage, _cbda}
}

// Index returns the index of the header within the document.  This is used to
// form its zip packaged filename as well as to match it with its relationship
// ID.
func (_cdegb Header) Index() int {
	for _dbdgg, _aeaff := range _cdegb._bgac._adf {
		if _aeaff == _cdegb._ecead {
			return _dbdgg
		}
	}
	return -1
}

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (_cdga ParagraphProperties) SetFirstLineIndent(m _bc.Distance) {
	if _cdga._aedc.Ind == nil {
		_cdga._aedc.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_cdga._aedc.Ind.FirstLineAttr = nil
	} else {
		_cdga._aedc.Ind.FirstLineAttr = &_ca.ST_TwipsMeasure{}
		_cdga._aedc.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(m / _bc.Twips))
	}
}

// AddStyle adds a new empty style, if styleID is already exists, it will return the style.
func (_ecdfd Styles) AddStyle(styleID string, t _db.ST_StyleType, isDefault bool) Style {
	if _eeea, _aeec := _ecdfd.SearchStyleById(styleID); _aeec {
		return _eeea
	}
	_bdagb := _db.NewCT_Style()
	_bdagb.TypeAttr = t
	if isDefault {
		_bdagb.DefaultAttr = &_ca.ST_OnOff{}
		_bdagb.DefaultAttr.Bool = _b.Bool(isDefault)
	}
	_bdagb.StyleIdAttr = _b.String(styleID)
	_ecdfd._cdgg.Style = append(_ecdfd._cdgg.Style, _bdagb)
	return Style{_bdagb}
}

// SetStrict is a shortcut for document.SetConformance,
// as one of these values from github.com/arcpd/msword/schema/soo/ofc/sharedTypes:
// ST_ConformanceClassUnset, ST_ConformanceClassStrict or ST_ConformanceClassTransitional.
func (_bgbd Document) SetStrict(strict bool) {
	if strict {
		_bgbd._adcb.ConformanceAttr = _ca.ST_ConformanceClassStrict
	} else {
		_bgbd._adcb.ConformanceAttr = _ca.ST_ConformanceClassTransitional
	}
}
func (_eega *Document) insertParagraph(_gfda Paragraph, _cfaf bool) Paragraph {
	if _eega._adcb.Body == nil {
		return _eega.AddParagraph()
	}
	_beaf := _gfda.X()
	for _, _fdadc := range _eega._adcb.Body.EG_BlockLevelElts {
		for _, _edeb := range _fdadc.EG_ContentBlockContent {
			for _beeab, _abf := range _edeb.P {
				if _abf == _beaf {
					_gaee := _db.NewCT_P()
					_edeb.P = append(_edeb.P, nil)
					if _cfaf {
						copy(_edeb.P[_beeab+1:], _edeb.P[_beeab:])
						_edeb.P[_beeab] = _gaee
					} else {
						copy(_edeb.P[_beeab+2:], _edeb.P[_beeab+1:])
						_edeb.P[_beeab+1] = _gaee
					}
					return Paragraph{_eega, _gaee}
				}
			}
			for _, _eage := range _edeb.Tbl {
				for _, _dfec := range _eage.EG_ContentRowContent {
					for _, _ddbd := range _dfec.Tr {
						for _, _cfabg := range _ddbd.EG_ContentCellContent {
							for _, _adff := range _cfabg.Tc {
								for _, _bcec := range _adff.EG_BlockLevelElts {
									for _, _dgf := range _bcec.EG_ContentBlockContent {
										for _faedd, _fddg := range _dgf.P {
											if _fddg == _beaf {
												_degb := _db.NewCT_P()
												_dgf.P = append(_dgf.P, nil)
												if _cfaf {
													copy(_dgf.P[_faedd+1:], _dgf.P[_faedd:])
													_dgf.P[_faedd] = _degb
												} else {
													copy(_dgf.P[_faedd+2:], _dgf.P[_faedd+1:])
													_dgf.P[_faedd+1] = _degb
												}
												return Paragraph{_eega, _degb}
											}
										}
									}
								}
							}
						}
					}
				}
			}
			if _edeb.Sdt != nil && _edeb.Sdt.SdtContent != nil && _edeb.Sdt.SdtContent.P != nil {
				for _aafe, _aadg := range _edeb.Sdt.SdtContent.P {
					if _aadg == _beaf {
						_acdd := _db.NewCT_P()
						_edeb.Sdt.SdtContent.P = append(_edeb.Sdt.SdtContent.P, nil)
						if _cfaf {
							copy(_edeb.Sdt.SdtContent.P[_aafe+1:], _edeb.Sdt.SdtContent.P[_aafe:])
							_edeb.Sdt.SdtContent.P[_aafe] = _acdd
						} else {
							copy(_edeb.Sdt.SdtContent.P[_aafe+2:], _edeb.Sdt.SdtContent.P[_aafe+1:])
							_edeb.Sdt.SdtContent.P[_aafe+1] = _acdd
						}
						return Paragraph{_eega, _acdd}
					}
				}
			}
		}
	}
	return _eega.AddParagraph()
}

// SetNumberingDefinitionByID sets the numbering definition ID directly, which must
// match an ID defined in numbering.xml
func (_aadf Paragraph) SetNumberingDefinitionByID(abstractNumberID int64) {
	_aadf.ensurePPr()
	if _aadf._dcfbd.PPr.NumPr == nil {
		_aadf._dcfbd.PPr.NumPr = _db.NewCT_NumPr()
	}
	_gcgef := _db.NewCT_DecimalNumber()
	_gcgef.ValAttr = int64(abstractNumberID)
	_aadf._dcfbd.PPr.NumPr.NumId = _gcgef
}

// SetNumberingDefinition sets the numbering definition ID via a NumberingDefinition
// defined in numbering.xml
func (_cadg Paragraph) SetNumberingDefinition(nd NumberingDefinition) {
	_cadg.ensurePPr()
	if _cadg._dcfbd.PPr.NumPr == nil {
		_cadg._dcfbd.PPr.NumPr = _db.NewCT_NumPr()
	}
	_gagg := _db.NewCT_DecimalNumber()
	_cedg := int64(-1)
	for _, _fcag := range _cadg._cfge.Numbering._gebaa.Num {
		if _fcag.AbstractNumId != nil && _fcag.AbstractNumId.ValAttr == nd.AbstractNumberID() {
			_cedg = _fcag.NumIdAttr
		}
	}
	if _cedg == -1 {
		_dadc := _db.NewCT_Num()
		_cadg._cfge.Numbering._gebaa.Num = append(_cadg._cfge.Numbering._gebaa.Num, _dadc)
		_dadc.NumIdAttr = int64(len(_cadg._cfge.Numbering._gebaa.Num))
		_dadc.AbstractNumId = _db.NewCT_DecimalNumber()
		_dadc.AbstractNumId.ValAttr = nd.AbstractNumberID()
	}
	_gagg.ValAttr = _cedg
	_cadg._dcfbd.PPr.NumPr.NumId = _gagg
}

// SetName sets the name of the image, visible in the properties of the image
// within Word.
func (_fdc AnchoredDrawing) SetName(name string) {
	_fdc._cg.DocPr.NameAttr = name
	for _, _dga := range _fdc._cg.Graphic.GraphicData.Any {
		if _cc, _eea := _dga.(*_ba.Pic); _eea {
			_cc.NvPicPr.CNvPr.DescrAttr = _b.String(name)
		}
	}
}

// SetRight sets the cell right margin
func (_fae CellMargins) SetRight(d _bc.Distance) {
	_fae._caff.Right = _db.NewCT_TblWidth()
	_eeb(_fae._caff.Right, d)
}
func (_bbcc *WatermarkPicture) getInnerElement(_beadg string) *_b.XSDAny {
	for _, _cacf := range _bbcc._dgdd.Any {
		_bgdcb, _cffe := _cacf.(*_b.XSDAny)
		if _cffe && (_bgdcb.XMLName.Local == _beadg || _bgdcb.XMLName.Local == "\u0076\u003a"+_beadg) {
			return _bgdcb
		}
	}
	return nil
}
func _bcd(_feef *_db.CT_Tbl, _fafg map[string]string) {
	for _, _cffbd := range _feef.EG_ContentRowContent {
		for _, _efgf := range _cffbd.Tr {
			for _, _cebb := range _efgf.EG_ContentCellContent {
				for _, _cebg := range _cebb.Tc {
					for _, _cggb := range _cebg.EG_BlockLevelElts {
						for _, _gaca := range _cggb.EG_ContentBlockContent {
							for _, _ecgc := range _gaca.P {
								_fdda(_ecgc, _fafg)
							}
							for _, _edca := range _gaca.Tbl {
								_bcd(_edca, _fafg)
							}
						}
					}
				}
			}
		}
	}
}

// Endnote returns the endnote based on the ID; this can be used nicely with
// the run.IsEndnote() functionality.
func (_cac *Document) Endnote(id int64) Endnote {
	for _, _dcdd := range _cac.Endnotes() {
		if _dcdd.id() == id {
			return _dcdd
		}
	}
	return Endnote{}
}

// SetPageBreakBefore controls if there is a page break before this paragraph.
func (_cafbe ParagraphProperties) SetPageBreakBefore(b bool) {
	if !b {
		_cafbe._aedc.PageBreakBefore = nil
	} else {
		_cafbe._aedc.PageBreakBefore = _db.NewCT_OnOff()
	}
}
func _ffgb(_afcg *_db.CT_P, _bbae map[string]string) {
	for _, _abab := range _afcg.EG_PContent {
		if _abab.Hyperlink != nil && _abab.Hyperlink.IdAttr != nil {
			if _abff, _eaeg := _bbae[*_abab.Hyperlink.IdAttr]; _eaeg {
				*_abab.Hyperlink.IdAttr = _abff
			}
		}
	}
}

// CharacterSpacingValue returns the value of characters spacing in twips (1/20 of point).
func (_gcefd ParagraphProperties) CharacterSpacingValue() int64 {
	if _fdeg := _gcefd._aedc.RPr.Spacing; _fdeg != nil {
		_caade := _fdeg.ValAttr
		if _caade.Int64 != nil {
			return *_caade.Int64
		}
	}
	return int64(0)
}
func (_beggf *Document) appendParagraph(_baf *Paragraph, _febea Paragraph, _aab bool) Paragraph {
	_gfca := _db.NewEG_BlockLevelElts()
	_beggf._adcb.Body.EG_BlockLevelElts = append(_beggf._adcb.Body.EG_BlockLevelElts, _gfca)
	_fgc := _db.NewEG_ContentBlockContent()
	_gfca.EG_ContentBlockContent = append(_gfca.EG_ContentBlockContent, _fgc)
	if _baf != nil {
		_dbd := _baf.X()
		for _, _afg := range _beggf._adcb.Body.EG_BlockLevelElts {
			for _, _fcfd := range _afg.EG_ContentBlockContent {
				for _cdff, _efcb := range _fcfd.P {
					if _efcb == _dbd {
						_bdc := _febea.X()
						_fcfd.P = append(_fcfd.P, nil)
						if _aab {
							copy(_fcfd.P[_cdff+1:], _fcfd.P[_cdff:])
							_fcfd.P[_cdff] = _bdc
						} else {
							copy(_fcfd.P[_cdff+2:], _fcfd.P[_cdff+1:])
							_fcfd.P[_cdff+1] = _bdc
						}
						break
					}
				}
				for _, _ggca := range _fcfd.Tbl {
					for _, _ebb := range _ggca.EG_ContentRowContent {
						for _, _fbab := range _ebb.Tr {
							for _, _bgd := range _fbab.EG_ContentCellContent {
								for _, _gfe := range _bgd.Tc {
									for _, _fddc := range _gfe.EG_BlockLevelElts {
										for _, _fcda := range _fddc.EG_ContentBlockContent {
											for _ecf, _cffbg := range _fcda.P {
												if _cffbg == _dbd {
													_fga := _febea.X()
													_fcda.P = append(_fcda.P, nil)
													if _aab {
														copy(_fcda.P[_ecf+1:], _fcda.P[_ecf:])
														_fcda.P[_ecf] = _fga
													} else {
														copy(_fcda.P[_ecf+2:], _fcda.P[_ecf+1:])
														_fcda.P[_ecf+1] = _fga
													}
													break
												}
											}
										}
									}
								}
							}
						}
					}
				}
				if _fcfd.Sdt != nil && _fcfd.Sdt.SdtContent != nil && _fcfd.Sdt.SdtContent.P != nil {
					for _gef, _dfdg := range _fcfd.Sdt.SdtContent.P {
						if _dfdg == _dbd {
							_fggfc := _febea.X()
							_fcfd.Sdt.SdtContent.P = append(_fcfd.Sdt.SdtContent.P, nil)
							if _aab {
								copy(_fcfd.Sdt.SdtContent.P[_gef+1:], _fcfd.Sdt.SdtContent.P[_gef:])
								_fcfd.Sdt.SdtContent.P[_gef] = _fggfc
							} else {
								copy(_fcfd.Sdt.SdtContent.P[_gef+2:], _fcfd.Sdt.SdtContent.P[_gef+1:])
								_fcfd.Sdt.SdtContent.P[_gef+1] = _fggfc
							}
							break
						}
					}
				}
			}
		}
	} else {
		_fgc.P = append(_fgc.P, _febea.X())
	}
	_fca := _febea.Properties()
	if _egc, _ffcd := _fca.Section(); _ffcd {
		var (
			_gcdcg map[string]string
			_ffea  map[string]string
		)
		_bfab := _egc.X().EG_HdrFtrReferences
		for _, _edaa := range _bfab {
			if _edaa.HeaderReference != nil {
				_gcdcg = map[string]string{_edaa.HeaderReference.IdAttr: _egc._bagegb._dfc.GetTargetByRelId(_edaa.HeaderReference.IdAttr)}
			}
			if _edaa.FooterReference != nil {
				_ffea = map[string]string{_edaa.FooterReference.IdAttr: _egc._bagegb._dfc.GetTargetByRelId(_edaa.FooterReference.IdAttr)}
			}
		}
		var _agef map[int]_ffa.ImageRef
		for _, _bfe := range _egc._bagegb.Headers() {
			for _bfcc, _baa := range _gcdcg {
				_ebfg := _g.Sprintf("\u0068\u0065\u0061d\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", (_bfe.Index() + 1))
				if _ebfg == _baa {
					_gfa := _g.Sprintf("\u0068\u0065\u0061d\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", _bfe.Index())
					_beggf._adf = append(_beggf._adf, _bfe.X())
					_dcgc := _beggf._dfc.AddRelationship(_gfa, _b.HeaderType)
					_dcgc.SetID(_bfcc)
					_beggf.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_gfa, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0068\u0065\u0061\u0064e\u0072\u002b\u0078\u006d\u006c")
					_beggf._gcdc = append(_beggf._gcdc, _ffa.NewRelationships())
					_adac := _bfe.Paragraphs()
					for _, _cdffc := range _adac {
						for _, _gbe := range _cdffc.Runs() {
							_fcaf := _gbe.DrawingAnchored()
							for _, _bef := range _fcaf {
								if _bbdf, _ggg := _bef.GetImage(); _ggg {
									_agef = map[int]_ffa.ImageRef{_bfe.Index(): _bbdf}
								}
							}
							_abbg := _gbe.DrawingInline()
							for _, _gfgd := range _abbg {
								if _cge, _cbb := _gfgd.GetImage(); _cbb {
									_agef = map[int]_ffa.ImageRef{_bfe.Index(): _cge}
								}
							}
						}
					}
				}
			}
		}
		for _agdf, _cgg := range _agef {
			for _, _fag := range _beggf.Headers() {
				if (_fag.Index() + 1) == _agdf {
					_gfcaa, _gcgb := _ffa.ImageFromFile(_cgg.Path())
					if _gcgb != nil {
						_c.Log.Debug("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063r\u0065\u0061\u0074\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u0073", _gcgb)
					}
					if _, _gcgb = _fag.AddImage(_gfcaa); _gcgb != nil {
						_c.Log.Debug("u\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020i\u006d\u0061\u0067\u0065\u0020\u0074\u006f \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020%\u0073", _gcgb)
					}
				}
				for _, _bbgc := range _fag.Paragraphs() {
					if _efg, _ageb := _egc._bagegb.Styles.SearchStyleById(_bbgc.Style()); _ageb {
						if _, _aea := _beggf.Styles.SearchStyleById(_bbgc.Style()); !_aea {
							_beggf.Styles.InsertStyle(_efg)
						}
					}
				}
			}
		}
		var _ebbd map[int]_ffa.ImageRef
		for _, _fcdab := range _egc._bagegb.Footers() {
			for _deda, _dag := range _ffea {
				_ede := _g.Sprintf("\u0066\u006f\u006ft\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", (_fcdab.Index() + 1))
				if _ede == _dag {
					_aaa := _g.Sprintf("\u0066\u006f\u006ft\u0065\u0072\u0025\u0064\u002e\u0078\u006d\u006c", _fcdab.Index())
					_beggf._abc = append(_beggf._abc, _fcdab.X())
					_dde := _beggf._dfc.AddRelationship(_aaa, _b.FooterType)
					_dde.SetID(_deda)
					_beggf.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072\u0064\u002f"+_aaa, "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0066\u006f\u006f\u0074e\u0072\u002b\u0078\u006d\u006c")
					_beggf._adg = append(_beggf._adg, _ffa.NewRelationships())
					_gbg := _fcdab.Paragraphs()
					for _, _aabc := range _gbg {
						for _, _cbde := range _aabc.Runs() {
							_fad := _cbde.DrawingAnchored()
							for _, _gbf := range _fad {
								if _ddd, _bcff := _gbf.GetImage(); _bcff {
									_ebbd = map[int]_ffa.ImageRef{_fcdab.Index(): _ddd}
								}
							}
							_beda := _cbde.DrawingInline()
							for _, _cbgf := range _beda {
								if _dagf, _ecfa := _cbgf.GetImage(); _ecfa {
									_ebbd = map[int]_ffa.ImageRef{_fcdab.Index(): _dagf}
								}
							}
						}
					}
				}
			}
		}
		for _cdfb, _aded := range _ebbd {
			for _, _acdb := range _beggf.Footers() {
				if (_acdb.Index() + 1) == _cdfb {
					_ebec, _afdf := _ffa.ImageFromFile(_aded.Path())
					if _afdf != nil {
						_c.Log.Debug("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063r\u0065\u0061\u0074\u0065\u0020\u0069\u006d\u0061\u0067\u0065:\u0020\u0025\u0073", _afdf)
					}
					if _, _afdf = _acdb.AddImage(_ebec); _afdf != nil {
						_c.Log.Debug("u\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0061\u0064\u0064\u0020i\u006d\u0061\u0067\u0065\u0020\u0074\u006f \u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020%\u0073", _afdf)
					}
				}
				for _, _deffc := range _acdb.Paragraphs() {
					if _cedc, _aadc := _egc._bagegb.Styles.SearchStyleById(_deffc.Style()); _aadc {
						if _, _gca := _beggf.Styles.SearchStyleById(_deffc.Style()); !_gca {
							_beggf.Styles.InsertStyle(_cedc)
						}
					}
				}
			}
		}
	}
	_dggb := _febea.Numbering()
	_beggf.Numbering._gebaa.AbstractNum = append(_beggf.Numbering._gebaa.AbstractNum, _dggb._gebaa.AbstractNum...)
	_beggf.Numbering._gebaa.Num = append(_beggf.Numbering._gebaa.Num, _dggb._gebaa.Num...)
	return Paragraph{_beggf, _febea.X()}
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_ebffd *Document) AddImage(i _ffa.Image) (_ffa.ImageRef, error) {
	_efac := _ffa.MakeImageRef(i, &_ebffd.DocBase, _ebffd._dfc)
	if i.Data == nil && i.Path == "" {
		return _efac, _bg.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _efac, _bg.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _efac, _bg.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	if i.Path != "" {
		_cfdb := _be.Add(i.Path)
		if _cfdb != nil {
			return _efac, _cfdb
		}
	}
	_ebffd.Images = append(_ebffd.Images, _efac)
	_dcaf := _g.Sprintf("\u006d\u0065d\u0069\u0061\u002fi\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", len(_ebffd.Images), i.Format)
	_fgea := _ebffd._dfc.AddRelationship(_dcaf, _b.ImageType)
	_ebffd.ContentTypes.EnsureDefault("\u0070\u006e\u0067", "\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg")
	_ebffd.ContentTypes.EnsureDefault("\u006a\u0070\u0065\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_ebffd.ContentTypes.EnsureDefault("\u006a\u0070\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_ebffd.ContentTypes.EnsureDefault("\u0077\u006d\u0066", "i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066")
	_ebffd.ContentTypes.EnsureDefault(i.Format, "\u0069\u006d\u0061\u0067\u0065\u002f"+i.Format)
	_efac.SetRelID(_fgea.X().IdAttr)
	_efac.SetTarget(_dcaf)
	return _efac, nil
}

// ReplaceText replace the text inside node.
func (_aaed *Node) ReplaceText(oldText, newText string) {
	switch _cabd := _aaed.X().(type) {
	case *Paragraph:
		for _, _aabfg := range _cabd.Runs() {
			for _, _cgcg := range _aabfg._afec.EG_RunInnerContent {
				if _cgcg.T != nil {
					_feda := _cgcg.T.Content
					_feda = _eg.ReplaceAll(_feda, oldText, newText)
					_cgcg.T.Content = _feda
				}
			}
		}
	}
	for _, _acfe := range _aaed.Children {
		_acfe.ReplaceText(oldText, newText)
	}
}

// SetText sets the text to be used in bullet mode.
func (_bgfe NumberingLevel) SetText(t string) {
	if t == "" {
		_bgfe._cfed.LvlText = nil
	} else {
		_bgfe._cfed.LvlText = _db.NewCT_LevelText()
		_bgfe._cfed.LvlText.ValAttr = _b.String(t)
	}
}

// FormField is a form within a document. It references the document, so changes
// to the form field wil be reflected in the document if it is saved.
type FormField struct {
	_fbcee *_db.CT_FFData
	_efgb  *_db.EG_RunInnerContent
}

// X returns the inner wrapped XML type.
func (_gacg TableStyleProperties) X() *_db.CT_TblPrBase { return _gacg._dagce }

// ExtractTextOptions extraction text options.
type ExtractTextOptions struct {

	// WithNumbering extract numbering elements if set to `true`.
	WithNumbering bool

	// NumberingIndent default value of numbering indent.
	NumberingIndent string

	// RunsOnNewLine write each of runs text on new line if set to `true`.
	RunsOnNewLine bool
}

// Paragraphs returns the paragraphs defined in the cell.
func (_egb Cell) Paragraphs() []Paragraph {
	_gddd := []Paragraph{}
	for _, _bcb := range _egb._caf.EG_BlockLevelElts {
		for _, _afad := range _bcb.EG_ContentBlockContent {
			for _, _dd := range _afad.P {
				_gddd = append(_gddd, Paragraph{_egb._cec, _dd})
			}
		}
	}
	return _gddd
}
func (_efacf *Document) getWatermarkHeaderInnerContentPictures() []*_db.CT_Picture {
	var _fagb []*_db.CT_Picture
	for _, _caef := range _efacf.Headers() {
		for _, _cgc := range _caef.X().EG_ContentBlockContent {
			for _, _gadcc := range _cgc.P {
				for _, _bdde := range _gadcc.EG_PContent {
					for _, _ggge := range _bdde.EG_ContentRunContent {
						if _ggge.R == nil {
							continue
						}
						for _, _gbba := range _ggge.R.EG_RunInnerContent {
							if _gbba.Pict == nil {
								continue
							}
							_gege := false
							for _, _feea := range _gbba.Pict.Any {
								_cefe, _gbae := _feea.(*_b.XSDAny)
								if _gbae && _cefe.XMLName.Local == "\u0073\u0068\u0061p\u0065" {
									_gege = true
								}
							}
							if _gege {
								_fagb = append(_fagb, _gbba.Pict)
							}
						}
					}
				}
			}
		}
	}
	return _fagb
}
func (_fbgc *Document) insertTable(_cfd Paragraph, _ceg bool) Table {
	_bbe := _fbgc._adcb.Body
	if _bbe == nil {
		return _fbgc.AddTable()
	}
	_beagc := _cfd.X()
	for _ggc, _aff := range _bbe.EG_BlockLevelElts {
		for _, _eecd := range _aff.EG_ContentBlockContent {
			for _ecdd, _eag := range _eecd.P {
				if _eag == _beagc {
					_gfde := _db.NewCT_Tbl()
					_cdfe := _db.NewEG_BlockLevelElts()
					_cee := _db.NewEG_ContentBlockContent()
					_cdfe.EG_ContentBlockContent = append(_cdfe.EG_ContentBlockContent, _cee)
					_cee.Tbl = append(_cee.Tbl, _gfde)
					_bbe.EG_BlockLevelElts = append(_bbe.EG_BlockLevelElts, nil)
					if _ceg {
						copy(_bbe.EG_BlockLevelElts[_ggc+1:], _bbe.EG_BlockLevelElts[_ggc:])
						_bbe.EG_BlockLevelElts[_ggc] = _cdfe
						if _ecdd != 0 {
							_fbef := _db.NewEG_BlockLevelElts()
							_bdg := _db.NewEG_ContentBlockContent()
							_fbef.EG_ContentBlockContent = append(_fbef.EG_ContentBlockContent, _bdg)
							_bdg.P = _eecd.P[:_ecdd]
							_bbe.EG_BlockLevelElts = append(_bbe.EG_BlockLevelElts, nil)
							copy(_bbe.EG_BlockLevelElts[_ggc+1:], _bbe.EG_BlockLevelElts[_ggc:])
							_bbe.EG_BlockLevelElts[_ggc] = _fbef
						}
						_eecd.P = _eecd.P[_ecdd:]
					} else {
						copy(_bbe.EG_BlockLevelElts[_ggc+2:], _bbe.EG_BlockLevelElts[_ggc+1:])
						_bbe.EG_BlockLevelElts[_ggc+1] = _cdfe
						if _ecdd != len(_eecd.P)-1 {
							_gfbd := _db.NewEG_BlockLevelElts()
							_gbb := _db.NewEG_ContentBlockContent()
							_gfbd.EG_ContentBlockContent = append(_gfbd.EG_ContentBlockContent, _gbb)
							_gbb.P = _eecd.P[_ecdd+1:]
							_bbe.EG_BlockLevelElts = append(_bbe.EG_BlockLevelElts, nil)
							copy(_bbe.EG_BlockLevelElts[_ggc+3:], _bbe.EG_BlockLevelElts[_ggc+2:])
							_bbe.EG_BlockLevelElts[_ggc+2] = _gfbd
						}
						_eecd.P = _eecd.P[:_ecdd+1]
					}
					return Table{_fbgc, _gfde}
				}
			}
			for _, _cbdb := range _eecd.Tbl {
				_fedb := _gaed(_cbdb, _beagc, _ceg)
				if _fedb != nil {
					return Table{_fbgc, _fedb}
				}
			}
		}
	}
	return _fbgc.AddTable()
}

// Table is a table within a document.
type Table struct {
	_ffcf  *Document
	_cgffe *_db.CT_Tbl
}

// SetThemeShade sets the shade based off the theme color.
func (_dab Color) SetThemeShade(s uint8) {
	_eeab := _g.Sprintf("\u0025\u0030\u0032\u0078", s)
	_dab._fcd.ThemeShadeAttr = &_eeab
}

// SetEndIndent controls the end indentation.
func (_dagc ParagraphProperties) SetEndIndent(m _bc.Distance) {
	if _dagc._aedc.Ind == nil {
		_dagc._aedc.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_dagc._aedc.Ind.EndAttr = nil
	} else {
		_dagc._aedc.Ind.EndAttr = &_db.ST_SignedTwipsMeasure{}
		_dagc._aedc.Ind.EndAttr.Int64 = _b.Int64(int64(m / _bc.Twips))
	}
}

// TableLook is the conditional formatting associated with a table style that
// has been assigned to a table.
type TableLook struct{ _febce *_db.CT_TblLook }

// AddHyperlink adds a hyperlink to a document. Adding the hyperlink to a document
// and setting it on a cell is more efficient than setting hyperlinks directly
// on a cell.
func (_cfcg Document) AddHyperlink(url string) _ffa.Hyperlink { return _cfcg._dfc.AddHyperlink(url) }

// MailMerge finds mail merge fields and replaces them with the text provided.  It also removes
// the mail merge source info from the document settings.
func (_gdac *Document) MailMerge(mergeContent map[string]string) {
	_cfdd := _gdac.mergeFields()
	_cfec := map[Paragraph][]Run{}
	for _, _feegd := range _cfdd {
		_aegac, _eaa := mergeContent[_feegd._gcfd]
		if _eaa {
			if _feegd._adedf {
				_aegac = _eg.ToUpper(_aegac)
			} else if _feegd._edece {
				_aegac = _eg.ToLower(_aegac)
			} else if _feegd._afbb {
				_aegac = _eg.Title(_aegac)
			} else if _feegd._gdcee {
				_gegb := _fed.Buffer{}
				for _fffed, _ebcb := range _aegac {
					if _fffed == 0 {
						_gegb.WriteRune(_e.ToUpper(_ebcb))
					} else {
						_gegb.WriteRune(_ebcb)
					}
				}
				_aegac = _gegb.String()
			}
			if _aegac != "" && _feegd._bedf != "" {
				_aegac = _feegd._bedf + _aegac
			}
			if _aegac != "" && _feegd._fggfa != "" {
				_aegac = _aegac + _feegd._fggfa
			}
		}
		if _feegd._fcaba {
			if len(_feegd._bcef.FldSimple) == 1 && len(_feegd._bcef.FldSimple[0].EG_PContent) == 1 && len(_feegd._bcef.FldSimple[0].EG_PContent[0].EG_ContentRunContent) == 1 {
				_fabca := &_db.EG_ContentRunContent{}
				_fabca.R = _feegd._bcef.FldSimple[0].EG_PContent[0].EG_ContentRunContent[0].R
				_feegd._bcef.FldSimple = nil
				_dgce := Run{_gdac, _fabca.R}
				_dgce.ClearContent()
				_dgce.AddText(_aegac)
				_feegd._bcef.EG_ContentRunContent = append(_feegd._bcef.EG_ContentRunContent, _fabca)
			}
		} else {
			_ebaa := _feegd._fbfg.Runs()
			for _cdcfe := _feegd._cbff; _cdcfe <= _feegd._eged; _cdcfe++ {
				if _cdcfe == _feegd._gbef+1 {
					_ebaa[_cdcfe].ClearContent()
					_ebaa[_cdcfe].AddText(_aegac)
				} else {
					_cfec[_feegd._fbfg] = append(_cfec[_feegd._fbfg], _ebaa[_cdcfe])
				}
			}
		}
	}
	for _daccb, _cdfc := range _cfec {
		for _, _abdc := range _cdfc {
			_daccb.RemoveRun(_abdc)
		}
	}
	_gdac.Settings.RemoveMailMerge()
}

// GetWrapPathLineTo return wrapPath lineTo value.
func (_ceb AnchorDrawWrapOptions) GetWrapPathLineTo() []*_aa.CT_Point2D { return _ceb._bb }

// Clear clears all content within a header
func (_bdag Header) Clear() { _bdag._ecead.EG_ContentBlockContent = nil }

// Levels returns all of the numbering levels defined in the definition.
func (_eafbf NumberingDefinition) Levels() []NumberingLevel {
	_cadd := []NumberingLevel{}
	for _, _gebd := range _eafbf._bbgd.Lvl {
		_cadd = append(_cadd, NumberingLevel{_gebd})
	}
	return _cadd
}

// SizeValue returns the value of run font size in points.
func (_cdcab RunProperties) SizeValue() float64 {
	if _baec := _cdcab._cadb.Sz; _baec != nil {
		_efbb := _baec.ValAttr
		if _efbb.ST_UnsignedDecimalNumber != nil {
			return float64(*_efbb.ST_UnsignedDecimalNumber) / 2
		}
	}
	return 0.0
}

// PossibleValues returns the possible values for a FormFieldTypeDropDown.
func (_cfce FormField) PossibleValues() []string {
	if _cfce._fbcee.DdList == nil {
		return nil
	}
	_bcce := []string{}
	for _, _gbbe := range _cfce._fbcee.DdList.ListEntry {
		if _gbbe == nil {
			continue
		}
		_bcce = append(_bcce, _gbbe.ValAttr)
	}
	return _bcce
}

// FormFields extracts all of the fields from a document.  They can then be
// manipulated via the methods on the field and the document saved.
func (_bbgf *Document) FormFields() []FormField {
	_ffba := []FormField{}
	for _, _ffgf := range _bbgf.Paragraphs() {
		_fabg := _ffgf.Runs()
		for _fgfc, _abcf := range _fabg {
			for _, _faff := range _abcf._afec.EG_RunInnerContent {
				if _faff.FldChar == nil || _faff.FldChar.FfData == nil {
					continue
				}
				if _faff.FldChar.FldCharTypeAttr == _db.ST_FldCharTypeBegin {
					if len(_faff.FldChar.FfData.Name) == 0 || _faff.FldChar.FfData.Name[0].ValAttr == nil {
						continue
					}
					_caad := FormField{_fbcee: _faff.FldChar.FfData}
					if _faff.FldChar.FfData.TextInput != nil {
						for _ddbe := _fgfc + 1; _ddbe < len(_fabg)-1; _ddbe++ {
							if len(_fabg[_ddbe]._afec.EG_RunInnerContent) == 0 {
								continue
							}
							_fde := _fabg[_ddbe]._afec.EG_RunInnerContent[0]
							if _fde.FldChar != nil && _fde.FldChar.FldCharTypeAttr == _db.ST_FldCharTypeSeparate {
								if len(_fabg[_ddbe+1]._afec.EG_RunInnerContent) == 0 {
									continue
								}
								if _fabg[_ddbe+1]._afec.EG_RunInnerContent[0].FldChar == nil {
									_caad._efgb = _fabg[_ddbe+1]._afec.EG_RunInnerContent[0]
									break
								}
							}
						}
					}
					_ffba = append(_ffba, _caad)
				}
			}
		}
	}
	for _, _dagfc := range _bbgf.Headers() {
		for _, _deffe := range _dagfc.Paragraphs() {
			_dcf := _deffe.Runs()
			for _daa, _ecdf := range _dcf {
				for _, _cacc := range _ecdf._afec.EG_RunInnerContent {
					if _cacc.FldChar == nil || _cacc.FldChar.FfData == nil {
						continue
					}
					if _cacc.FldChar.FldCharTypeAttr == _db.ST_FldCharTypeBegin {
						if len(_cacc.FldChar.FfData.Name) == 0 || _cacc.FldChar.FfData.Name[0].ValAttr == nil {
							continue
						}
						_aabf := FormField{_fbcee: _cacc.FldChar.FfData}
						if _cacc.FldChar.FfData.TextInput != nil {
							for _dffd := _daa + 1; _dffd < len(_dcf)-1; _dffd++ {
								if len(_dcf[_dffd]._afec.EG_RunInnerContent) == 0 {
									continue
								}
								_gafc := _dcf[_dffd]._afec.EG_RunInnerContent[0]
								if _gafc.FldChar != nil && _gafc.FldChar.FldCharTypeAttr == _db.ST_FldCharTypeSeparate {
									if len(_dcf[_dffd+1]._afec.EG_RunInnerContent) == 0 {
										continue
									}
									if _dcf[_dffd+1]._afec.EG_RunInnerContent[0].FldChar == nil {
										_aabf._efgb = _dcf[_dffd+1]._afec.EG_RunInnerContent[0]
										break
									}
								}
							}
						}
						_ffba = append(_ffba, _aabf)
					}
				}
			}
		}
	}
	for _, _ccc := range _bbgf.Footers() {
		for _, _cage := range _ccc.Paragraphs() {
			_dcfe := _cage.Runs()
			for _edcg, _ggcf := range _dcfe {
				for _, _dgaf := range _ggcf._afec.EG_RunInnerContent {
					if _dgaf.FldChar == nil || _dgaf.FldChar.FfData == nil {
						continue
					}
					if _dgaf.FldChar.FldCharTypeAttr == _db.ST_FldCharTypeBegin {
						if len(_dgaf.FldChar.FfData.Name) == 0 || _dgaf.FldChar.FfData.Name[0].ValAttr == nil {
							continue
						}
						_aggg := FormField{_fbcee: _dgaf.FldChar.FfData}
						if _dgaf.FldChar.FfData.TextInput != nil {
							for _fcca := _edcg + 1; _fcca < len(_dcfe)-1; _fcca++ {
								if len(_dcfe[_fcca]._afec.EG_RunInnerContent) == 0 {
									continue
								}
								_bedad := _dcfe[_fcca]._afec.EG_RunInnerContent[0]
								if _bedad.FldChar != nil && _bedad.FldChar.FldCharTypeAttr == _db.ST_FldCharTypeSeparate {
									if len(_dcfe[_fcca+1]._afec.EG_RunInnerContent) == 0 {
										continue
									}
									if _dcfe[_fcca+1]._afec.EG_RunInnerContent[0].FldChar == nil {
										_aggg._efgb = _dcfe[_fcca+1]._afec.EG_RunInnerContent[0]
										break
									}
								}
							}
						}
						_ffba = append(_ffba, _aggg)
					}
				}
			}
		}
	}
	return _ffba
}

// SetDefaultValue sets the default value of a FormFieldTypeDropDown. For
// FormFieldTypeDropDown, the value must be one of the fields possible values.
func (_dfea FormField) SetDefaultValue(v string) {
	if _dfea._fbcee.DdList != nil {
		for _ecba, _fgeag := range _dfea.PossibleValues() {
			if _fgeag == v {
				_dfea._fbcee.DdList.Default = _db.NewCT_DecimalNumber()
				_dfea._fbcee.DdList.Default.ValAttr = int64(_ecba)
				break
			}
		}
	}
}

// SetTextWrapTight sets the text wrap to tight with a give wrap type.
func (_ffc AnchoredDrawing) SetTextWrapTight(option *AnchorDrawWrapOptions) {
	_ffc._cg.Choice = &_db.WdEG_WrapTypeChoice{}
	_ffc._cg.Choice.WrapTight = _db.NewWdCT_WrapTight()
	_ffc._cg.Choice.WrapTight.WrapTextAttr = _db.WdST_WrapTextBothSides
	_gc := false
	_ffc._cg.Choice.WrapTight.WrapPolygon.EditedAttr = &_gc
	if option == nil {
		option = NewAnchorDrawWrapOptions()
	}
	_ffc._cg.Choice.WrapTight.WrapPolygon.LineTo = option.GetWrapPathLineTo()
	_ffc._cg.Choice.WrapTight.WrapPolygon.Start = option.GetWrapPathStart()
	_ffc._cg.LayoutInCellAttr = true
	_ffc._cg.AllowOverlapAttr = true
}

// Rows returns the rows defined in the table.
func (_cddf Table) Rows() []Row {
	_fgdf := []Row{}
	for _, _fbbe := range _cddf._cgffe.EG_ContentRowContent {
		for _, _dagaa := range _fbbe.Tr {
			_fgdf = append(_fgdf, Row{_cddf._ffcf, _dagaa})
		}
		if _fbbe.Sdt != nil && _fbbe.Sdt.SdtContent != nil {
			for _, _ebbeb := range _fbbe.Sdt.SdtContent.Tr {
				_fgdf = append(_fgdf, Row{_cddf._ffcf, _ebbeb})
			}
		}
	}
	return _fgdf
}

// SetLastColumn controls the conditional formatting for the last column in a table.
func (_cbgfc TableLook) SetLastColumn(on bool) {
	if !on {
		_cbgfc._febce.LastColumnAttr = &_ca.ST_OnOff{}
		_cbgfc._febce.LastColumnAttr.ST_OnOff1 = _ca.ST_OnOff1Off
	} else {
		_cbgfc._febce.LastColumnAttr = &_ca.ST_OnOff{}
		_cbgfc._febce.LastColumnAttr.ST_OnOff1 = _ca.ST_OnOff1On
	}
}

// // SetBeforeLineSpacing sets spacing above paragraph in line units.
func (_cbbf Paragraph) SetBeforeLineSpacing(d _bc.Distance) {
	_cbbf.ensurePPr()
	if _cbbf._dcfbd.PPr.Spacing == nil {
		_cbbf._dcfbd.PPr.Spacing = _db.NewCT_Spacing()
	}
	_dbdaf := _cbbf._dcfbd.PPr.Spacing
	_dbdaf.BeforeLinesAttr = _b.Int64(int64(d / _bc.Twips))
}
func _gfag(_afgd Paragraph) *_db.CT_NumPr {
	_afgd.ensurePPr()
	if _afgd._dcfbd.PPr.NumPr == nil {
		return nil
	}
	return _afgd._dcfbd.PPr.NumPr
}

// SetHeader sets a section header.
func (_bdga Section) SetHeader(h Header, t _db.ST_HdrFtr) {
	_ecdc := _db.NewEG_HdrFtrReferences()
	_bdga._ggcfb.EG_HdrFtrReferences = append(_bdga._ggcfb.EG_HdrFtrReferences, _ecdc)
	_ecdc.HeaderReference = _db.NewCT_HdrFtrRef()
	_ecdc.HeaderReference.TypeAttr = t
	_afbg := _bdga._bagegb._dfc.FindRIDForN(h.Index(), _b.HeaderType)
	if _afbg == "" {
		_c.Log.Debug("\u0075\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u0068\u0065\u0061\u0064\u0065r \u0049\u0044")
	}
	_ecdc.HeaderReference.IdAttr = _afbg
}

// SetCantSplit set row properties for Can't Split value.
func (_eeaa RowProperties) SetCantSplit(val bool) {
	if !val {
		_eeaa._caaf.CantSplit = nil
	} else {
		_cfdg := _db.NewCT_OnOff()
		_eeaa._caaf.CantSplit = []*_db.CT_OnOff{_cfdg}
	}
}

// AddCheckBox adds checkbox form field to the paragraph and returns it.
func (_bacb Paragraph) AddCheckBox(name string) FormField {
	_ecgeb := _bacb.addFldCharsForField(name, "\u0046\u004f\u0052M\u0043\u0048\u0045\u0043\u004b\u0042\u004f\u0058")
	_ecgeb._fbcee.CheckBox = _db.NewCT_FFCheckBox()
	return _ecgeb
}

// Run is a run of text within a paragraph that shares the same formatting.
type Run struct {
	_gacf *Document
	_afec *_db.CT_R
}

// SetTextWrapBehindText sets the text wrap to behind text.
func (_edd AnchoredDrawing) SetTextWrapBehindText() {
	_edd._cg.Choice = &_db.WdEG_WrapTypeChoice{}
	_edd._cg.Choice.WrapNone = _db.NewWdCT_WrapNone()
	_edd._cg.BehindDocAttr = true
	_edd._cg.LayoutInCellAttr = true
	_edd._cg.AllowOverlapAttr = true
}

// RunProperties returns the RunProperties controlling numbering level font, etc.
func (_gdef NumberingLevel) RunProperties() RunProperties {
	if _gdef._cfed.RPr == nil {
		_gdef._cfed.RPr = _db.NewCT_RPr()
	}
	return RunProperties{_gdef._cfed.RPr}
}

// SetStart sets the cell start margin
func (_bbfe CellMargins) SetStart(d _bc.Distance) {
	_bbfe._caff.Start = _db.NewCT_TblWidth()
	_eeb(_bbfe._caff.Start, d)
}

// ParagraphProperties returns the paragraph properties controlling text formatting within the table.
func (_egfb TableConditionalFormatting) ParagraphProperties() ParagraphStyleProperties {
	if _egfb._adeac.PPr == nil {
		_egfb._adeac.PPr = _db.NewCT_PPrGeneral()
	}
	return ParagraphStyleProperties{_egfb._adeac.PPr}
}

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (_gefgd TableStyleProperties) SetCellSpacingPercent(pct float64) {
	_gefgd._dagce.TblCellSpacing = _db.NewCT_TblWidth()
	_gefgd._dagce.TblCellSpacing.TypeAttr = _db.ST_TblWidthPct
	_gefgd._dagce.TblCellSpacing.WAttr = &_db.ST_MeasurementOrPercent{}
	_gefgd._dagce.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &_db.ST_DecimalNumberOrPercent{}
	_gefgd._dagce.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = _b.Int64(int64(pct * 50))
}

// Type returns the type of the style.
func (_aecfe Style) Type() _db.ST_StyleType { return _aecfe._eedcd.TypeAttr }

// SetAlignment controls the paragraph alignment
func (_ggfg ParagraphProperties) SetAlignment(align _db.ST_Jc) {
	if align == _db.ST_JcUnset {
		_ggfg._aedc.Jc = nil
	} else {
		_ggfg._aedc.Jc = _db.NewCT_Jc()
		_ggfg._aedc.Jc.ValAttr = align
	}
}
func (_gfdd Paragraph) addStartBookmark(_eafc int64, _bbbc string) *_db.CT_Bookmark {
	_ebbb := _db.NewEG_PContent()
	_gfdd._dcfbd.EG_PContent = append(_gfdd._dcfbd.EG_PContent, _ebbb)
	_aeed := _db.NewEG_ContentRunContent()
	_aacae := _db.NewEG_RunLevelElts()
	_ggcgf := _db.NewEG_RangeMarkupElements()
	_eebff := _db.NewCT_Bookmark()
	_eebff.NameAttr = _bbbc
	_eebff.IdAttr = _eafc
	_ggcgf.BookmarkStart = _eebff
	_ebbb.EG_ContentRunContent = append(_ebbb.EG_ContentRunContent, _aeed)
	_aeed.EG_RunLevelElts = append(_aeed.EG_RunLevelElts, _aacae)
	_aacae.EG_RangeMarkupElements = append(_aacae.EG_RangeMarkupElements, _ggcgf)
	return _eebff
}

// TextItem is used for keeping text with references to a paragraph and run or a table, a row and a cell where it is located.
type TextItem struct {
	Text        string
	DrawingInfo *DrawingInfo
	Paragraph   *_db.CT_P
	Hyperlink   *_db.CT_Hyperlink
	Run         *_db.CT_R
	TableInfo   *TableInfo
}

// Shadow returns true if paragraph shadow is on.
func (_facc ParagraphProperties) Shadow() bool { return _ebdd(_facc._aedc.RPr.Shadow) }

// X returns the inner wrapped XML type.
func (_gdfg TableConditionalFormatting) X() *_db.CT_TblStylePr { return _gdfg._adeac }

// SetPictureSize set watermark picture size with given width and height.
func (_gaged *WatermarkPicture) SetPictureSize(width, height int64) {
	if _gaged._febge != nil {
		_acbcb := _gaged.GetShapeStyle()
		_acbcb.SetWidth(float64(width) * _bc.Point)
		_acbcb.SetHeight(float64(height) * _bc.Point)
		_gaged.SetShapeStyle(_acbcb)
	}
}

// Paragraphs returns the paragraphs defined in a footnote.
func (_egcg Footnote) Paragraphs() []Paragraph {
	_dffb := []Paragraph{}
	for _, _eaff := range _egcg.content() {
		for _, _adcge := range _eaff.P {
			_dffb = append(_dffb, Paragraph{_egcg._abca, _adcge})
		}
	}
	return _dffb
}

// SetNumberingLevel sets the numbering level of a paragraph.  If used, then the
// NumberingDefinition must also be set via SetNumberingDefinition or
// SetNumberingDefinitionByID.
func (_aaea Paragraph) SetNumberingLevel(listLevel int) {
	_aaea.ensurePPr()
	if _aaea._dcfbd.PPr.NumPr == nil {
		_aaea._dcfbd.PPr.NumPr = _db.NewCT_NumPr()
	}
	_gfcb := _db.NewCT_DecimalNumber()
	_gfcb.ValAttr = int64(listLevel)
	_aaea._dcfbd.PPr.NumPr.Ilvl = _gfcb
}

// AddTabStop adds a tab stop to the paragraph.  It controls the position of text when using Run.AddTab()
func (_cgec ParagraphProperties) AddTabStop(position _bc.Distance, justificaton _db.ST_TabJc, leader _db.ST_TabTlc) {
	if _cgec._aedc.Tabs == nil {
		_cgec._aedc.Tabs = _db.NewCT_Tabs()
	}
	_aefab := _db.NewCT_TabStop()
	_aefab.LeaderAttr = leader
	_aefab.ValAttr = justificaton
	_aefab.PosAttr.Int64 = _b.Int64(int64(position / _bc.Twips))
	_cgec._aedc.Tabs.Tab = append(_cgec._aedc.Tabs.Tab, _aefab)
}

// InlineDrawing is an inlined image within a run.
type InlineDrawing struct {
	_gfgg *Document
	_adb  *_db.WdInline
}

// SetRightToLeft sets the run text goes from right to left.
func (_bbee RunProperties) SetRightToLeft(b bool) {
	if !b {
		_bbee._cadb.Rtl = nil
	} else {
		_bbee._cadb.Rtl = _db.NewCT_OnOff()
	}
}

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_aceb Footer) AddImage(i _ffa.Image) (_ffa.ImageRef, error) {
	var _agbde _ffa.Relationships
	for _cbcf, _ceeg := range _aceb._bcge._abc {
		if _ceeg == _aceb._eeed {
			_agbde = _aceb._bcge._adg[_cbcf]
		}
	}
	_cdcfa := _ffa.MakeImageRef(i, &_aceb._bcge.DocBase, _agbde)
	if i.Data == nil && i.Path == "" {
		return _cdcfa, _bg.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _cdcfa, _bg.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _cdcfa, _bg.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	_aceb._bcge.Images = append(_aceb._bcge.Images, _cdcfa)
	_cfgb := _g.Sprintf("\u006d\u0065d\u0069\u0061\u002fi\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", len(_aceb._bcge.Images), i.Format)
	_gcdb := _agbde.AddRelationship(_cfgb, _b.ImageType)
	_cdcfa.SetRelID(_gcdb.X().IdAttr)
	return _cdcfa, nil
}

// GetStyleByID returns Style by it's IdAttr.
func (_cgcb *Document) GetStyleByID(id string) Style {
	for _, _aafb := range _cgcb.Styles._cdgg.Style {
		if _aafb.StyleIdAttr != nil && *_aafb.StyleIdAttr == id {
			return Style{_aafb}
		}
	}
	return Style{}
}

// Italic returns true if paragraph font is italic.
func (_caada ParagraphProperties) Italic() bool {
	_eab := _caada._aedc.RPr
	return _ebdd(_eab.I) || _ebdd(_eab.ICs)
}

// FontTable returns document fontTable element.
func (_gdad *Document) FontTable() *_db.Fonts { return _gdad._begg }

// Color returns the style's Color.
func (_ddfg RunProperties) Color() Color {
	if _ddfg._cadb.Color == nil {
		_ddfg._cadb.Color = _db.NewCT_Color()
	}
	return Color{_ddfg._cadb.Color}
}

// Underline returns the type of paragraph underline.
func (_beccd ParagraphProperties) Underline() _db.ST_Underline {
	if _bcgf := _beccd._aedc.RPr.U; _bcgf != nil {
		return _bcgf.ValAttr
	}
	return 0
}
func _fbbf(_afaac *Document) map[int64]map[int64]int64 {
	_ddgbc := _afaac.Paragraphs()
	_dfcc := make(map[int64]map[int64]int64, 0)
	for _, _bbgfa := range _ddgbc {
		_bfgaca := _fcdc(_afaac, _bbgfa)
		if _bfgaca.NumberingLevel != nil && _bfgaca.AbstractNumId != nil {
			_egccc := *_bfgaca.AbstractNumId
			if _, _gdgg := _dfcc[_egccc]; _gdgg {
				if _cbdae := _bfgaca.NumberingLevel.X(); _cbdae != nil {
					if _, _afded := _dfcc[_egccc][_cbdae.IlvlAttr]; _afded {
						_dfcc[_egccc][_cbdae.IlvlAttr]++
					} else {
						_dfcc[_egccc][_cbdae.IlvlAttr] = 1
					}
				}
			} else {
				if _aagf := _bfgaca.NumberingLevel.X(); _aagf != nil {
					_dfcc[_egccc] = map[int64]int64{_aagf.IlvlAttr: 1}
				}
			}
		}
	}
	return _dfcc
}

// SetText sets the watermark text.
func (_gcac *WatermarkText) SetText(text string) {
	_edefb := _gcac.getShape()
	if _gcac._bgdbe != nil {
		_ffeg := _gcac._bgdbe.EG_ShapeElements
		if len(_ffeg) > 0 && _ffeg[0].Textpath != nil {
			_ffeg[0].Textpath.StringAttr = &text
		}
	} else {
		_gcbbb := _gcac.findNode(_edefb, "\u0074\u0065\u0078\u0074\u0070\u0061\u0074\u0068")
		for _bcfa, _dccca := range _gcbbb.Attrs {
			if _dccca.Name.Local == "\u0073\u0074\u0072\u0069\u006e\u0067" {
				_gcbbb.Attrs[_bcfa].Value = text
			}
		}
	}
}

// TableConditionalFormatting controls the conditional formatting within a table
// style.
type TableConditionalFormatting struct{ _adeac *_db.CT_TblStylePr }

// SetAfter sets the spacing that comes after the paragraph.
func (_gebda ParagraphSpacing) SetAfter(after _bc.Distance) {
	_gebda._ffeb.AfterAttr = &_ca.ST_TwipsMeasure{}
	_gebda._ffeb.AfterAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(after / _bc.Twips))
}

// DrawingAnchored returns a slice of AnchoredDrawings.
func (_fcgg Run) DrawingAnchored() []AnchoredDrawing {
	_gadaf := []AnchoredDrawing{}
	for _, _gcgeg := range _fcgg._afec.EG_RunInnerContent {
		if _gcgeg.Drawing == nil {
			continue
		}
		for _, _adda := range _gcgeg.Drawing.Anchor {
			_gadaf = append(_gadaf, AnchoredDrawing{_fcgg._gacf, _adda})
		}
	}
	return _gadaf
}
func (_ccef *WatermarkText) getShape() *_b.XSDAny {
	return _ccef.getInnerElement("\u0073\u0068\u0061p\u0065")
}
func (_ggdeg *WatermarkText) getInnerElement(_dedec string) *_b.XSDAny {
	for _, _fggda := range _ggdeg._gdgad.Any {
		_fefg, _fcadd := _fggda.(*_b.XSDAny)
		if _fcadd && (_fefg.XMLName.Local == _dedec || _fefg.XMLName.Local == "\u0076\u003a"+_dedec) {
			return _fefg
		}
	}
	return nil
}

// SetTarget sets the URL target of the hyperlink.
func (_dgda HyperLink) SetTarget(url string) {
	_fcfa := _dgda._deac.AddHyperlink(url)
	_dgda._cfcf.IdAttr = _b.String(_ffa.Relationship(_fcfa).ID())
	_dgda._cfcf.AnchorAttr = nil
}

// TableLook returns the table look, or conditional formatting applied to a table style.
func (_cdee TableProperties) TableLook() TableLook {
	if _cdee._faec.TblLook == nil {
		_cdee._faec.TblLook = _db.NewCT_TblLook()
	}
	return TableLook{_cdee._faec.TblLook}
}

// EastAsiaFont returns the name of paragraph font family for East Asia.
func (_cgfb ParagraphProperties) EastAsiaFont() string {
	if _gcfe := _cgfb._aedc.RPr.RFonts; _gcfe != nil {
		if _gcfe.EastAsiaAttr != nil {
			return *_gcfe.EastAsiaAttr
		}
	}
	return ""
}
func (_cfece Paragraph) addEndBookmark(_aaaef int64) *_db.CT_MarkupRange {
	_cbgb := _db.NewEG_PContent()
	_cfece._dcfbd.EG_PContent = append(_cfece._dcfbd.EG_PContent, _cbgb)
	_badce := _db.NewEG_ContentRunContent()
	_bffe := _db.NewEG_RunLevelElts()
	_febgg := _db.NewEG_RangeMarkupElements()
	_fgfca := _db.NewCT_MarkupRange()
	_fgfca.IdAttr = _aaaef
	_febgg.BookmarkEnd = _fgfca
	_cbgb.EG_ContentRunContent = append(_cbgb.EG_ContentRunContent, _badce)
	_badce.EG_RunLevelElts = append(_badce.EG_RunLevelElts, _bffe)
	_bffe.EG_RangeMarkupElements = append(_bffe.EG_RangeMarkupElements, _febgg)
	return _fgfca
}

// AddDefinition adds a new numbering definition.
func (_cfbg Numbering) AddDefinition() NumberingDefinition {
	_aefg := _db.NewCT_Num()
	_fgde := int64(1)
	for _, _eece := range _cfbg.Definitions() {
		if _eece.AbstractNumberID() >= _fgde {
			_fgde = _eece.AbstractNumberID() + 1
		}
	}
	_bfbc := int64(1)
	for _, _fggd := range _cfbg.X().Num {
		if _fggd.NumIdAttr >= _bfbc {
			_bfbc = _fggd.NumIdAttr + 1
		}
	}
	_aefg.NumIdAttr = _bfbc
	_aefg.AbstractNumId = _db.NewCT_DecimalNumber()
	_aefg.AbstractNumId.ValAttr = _fgde
	_ecge := _db.NewCT_AbstractNum()
	_ecge.AbstractNumIdAttr = _fgde
	_cfbg._gebaa.AbstractNum = append(_cfbg._gebaa.AbstractNum, _ecge)
	_cfbg._gebaa.Num = append(_cfbg._gebaa.Num, _aefg)
	return NumberingDefinition{_ecge}
}
func (_gaeg *WatermarkText) getShapeType() *_b.XSDAny {
	return _gaeg.getInnerElement("\u0073h\u0061\u0070\u0065\u0074\u0079\u0070e")
}

// FindNodeByStyleName return slice of node base on style name.
func (_fgb *Nodes) FindNodeByStyleName(styleName string) []Node {
	_bfff := []Node{}
	for _, _dddga := range _fgb._dcdf {
		switch _cgbfe := _dddga._bdcf.(type) {
		case *Paragraph:
			if _cgbfe != nil {
				if _baedc, _begf := _dddga._cffcb.Styles.SearchStyleByName(styleName); _begf {
					_edea := _cgbfe.Style()
					if _edea == _baedc.StyleID() {
						_bfff = append(_bfff, _dddga)
					}
				}
			}
		case *Table:
			if _cgbfe != nil {
				if _daga, _dagfa := _dddga._cffcb.Styles.SearchStyleByName(styleName); _dagfa {
					_addec := _cgbfe.Style()
					if _addec == _daga.StyleID() {
						_bfff = append(_bfff, _dddga)
					}
				}
			}
		}
		_ddfaa := Nodes{_dcdf: _dddga.Children}
		_bfff = append(_bfff, _ddfaa.FindNodeByStyleName(styleName)...)
	}
	return _bfff
}

// Node is document element node,
// contains Paragraph or Table element.
type Node struct {
	_cffcb           *Document
	_bdcf            interface{}
	Style            Style
	AnchoredDrawings []AnchoredDrawing
	InlineDrawings   []InlineDrawing
	Children         []Node
}

// SetColumnSpan sets the number of Grid Columns Spanned by the Cell.  This is used
// to give the appearance of merged cells.
func (_cbg CellProperties) SetColumnSpan(cols int) {
	if cols == 0 {
		_cbg._beea.GridSpan = nil
	} else {
		_cbg._beea.GridSpan = _db.NewCT_DecimalNumber()
		_cbg._beea.GridSpan.ValAttr = int64(cols)
	}
}
func _ebdd(_dafe *_db.CT_OnOff) bool { return _dafe != nil }

// GetStyle returns string style of the text in watermark and format it to TextpathStyle.
func (_efcg *WatermarkText) GetStyle() _ge.TextpathStyle {
	_ceea := _efcg.getShape()
	if _efcg._bgdbe != nil {
		_dgcbg := _efcg._bgdbe.EG_ShapeElements
		if len(_dgcbg) > 0 && _dgcbg[0].Textpath != nil {
			return _ge.NewTextpathStyle(*_dgcbg[0].Textpath.StyleAttr)
		}
	} else {
		_cagdg := _efcg.findNode(_ceea, "\u0074\u0065\u0078\u0074\u0070\u0061\u0074\u0068")
		for _, _aedad := range _cagdg.Attrs {
			if _aedad.Name.Local == "\u0073\u0074\u0079l\u0065" {
				return _ge.NewTextpathStyle(_aedad.Value)
			}
		}
	}
	return _ge.NewTextpathStyle("")
}

// AddFootnote will create a new footnote and attach it to the Paragraph in the
// location at the end of the previous run (footnotes create their own run within
// the paragraph). The text given to the function is simply a convenience helper,
// paragraphs and runs can always be added to the text of the footnote later.
func (_egg Paragraph) AddFootnote(text string) Footnote {
	var _bfcb int64
	if _egg._cfge.HasFootnotes() {
		for _, _bcad := range _egg._cfge.Footnotes() {
			if _bcad.id() > _bfcb {
				_bfcb = _bcad.id()
			}
		}
		_bfcb++
	} else {
		_bfcb = 0
		_egg._cfge._ffcef = &_db.Footnotes{}
		_egg._cfge._ffcef.CT_Footnotes = _db.CT_Footnotes{}
		_egg._cfge._ffcef.Footnote = make([]*_db.CT_FtnEdn, 0)
	}
	_egcgg := _db.NewCT_FtnEdn()
	_faae := _db.NewCT_FtnEdnRef()
	_faae.IdAttr = _bfcb
	_egg._cfge._ffcef.CT_Footnotes.Footnote = append(_egg._cfge._ffcef.CT_Footnotes.Footnote, _egcgg)
	_cgfdf := _egg.AddRun()
	_dfcf := _cgfdf.Properties()
	_dfcf.SetStyle("\u0046\u006f\u006f\u0074\u006e\u006f\u0074\u0065\u0041n\u0063\u0068\u006f\u0072")
	_cgfdf._afec.EG_RunInnerContent = []*_db.EG_RunInnerContent{_db.NewEG_RunInnerContent()}
	_cgfdf._afec.EG_RunInnerContent[0].FootnoteReference = _faae
	_bddga := Footnote{_egg._cfge, _egcgg}
	_bddga._cefg.IdAttr = _bfcb
	_bddga._cefg.EG_BlockLevelElts = []*_db.EG_BlockLevelElts{_db.NewEG_BlockLevelElts()}
	_cdad := _bddga.AddParagraph()
	_cdad.Properties().SetStyle("\u0046\u006f\u006f\u0074\u006e\u006f\u0074\u0065")
	_cdad._dcfbd.PPr.RPr = _db.NewCT_ParaRPr()
	_bbfdc := _cdad.AddRun()
	_bbfdc.AddTab()
	_bbfdc.AddText(text)
	return _bddga
}

// ParagraphBorders allows manipulation of borders on a paragraph.
type ParagraphBorders struct {
	_fdccg *Document
	_afaaa *_db.CT_PBdr
}

// FindNodeByText return node based on matched text and return a slice of node.
func (_cbfd *Nodes) FindNodeByText(text string) []Node {
	_bfee := []Node{}
	for _, _dfeb := range _cbfd._dcdf {
		if _eg.TrimSpace(_dfeb.Text()) == text {
			_bfee = append(_bfee, _dfeb)
		}
		_bfgacf := Nodes{_dcdf: _dfeb.Children}
		_bfee = append(_bfee, _bfgacf.FindNodeByText(text)...)
	}
	return _bfee
}

// read reads a document from an io.Reader.
func Read(r _fg.ReaderAt, size int64) (*Document, error) { return _gab(r, size, "") }

// SetWidthAuto sets the the cell width to automatic.
func (_dgcf CellProperties) SetWidthAuto() {
	_dgcf._beea.TcW = _db.NewCT_TblWidth()
	_dgcf._beea.TcW.TypeAttr = _db.ST_TblWidthAuto
}

// New constructs an empty document that content can be added to.
func New() *Document {
	_faad := &Document{_adcb: _db.NewDocument()}
	_faad.ContentTypes = _ffa.NewContentTypes()
	_faad._adcb.Body = _db.NewCT_Body()
	_faad._adcb.ConformanceAttr = _ca.ST_ConformanceClassTransitional
	_faad._dfc = _ffa.NewRelationships()
	_faad.AppProperties = _ffa.NewAppProperties()
	_faad.CoreProperties = _ffa.NewCoreProperties()
	_faad.ContentTypes.AddOverride("\u002fw\u006fr\u0064\u002f\u0064\u006f\u0063u\u006d\u0065n\u0074\u002e\u0078\u006d\u006c", "\u0061p\u0070\u006c\u0069c\u0061\u0074\u0069o\u006e/v\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072dp\u0072o\u0063\u0065\u0073\u0073\u0069\u006eg\u006d\u006c\u002e\u0064\u006fc\u0075\u006d\u0065\u006e\u0074\u002e\u006d\u0061\u0069\u006e\u002bx\u006d\u006c")
	_faad.Settings = NewSettings()
	_faad._dfc.AddRelationship("\u0073\u0065\u0074t\u0069\u006e\u0067\u0073\u002e\u0078\u006d\u006c", _b.SettingsType)
	_faad.ContentTypes.AddOverride("\u002fw\u006fr\u0064\u002f\u0073\u0065\u0074t\u0069\u006eg\u0073\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069o\u006e\u002fv\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006dl\u0066\u006f\u0072\u006da\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c.\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002b\u0078\u006d\u006c")
	_faad.Rels = _ffa.NewRelationships()
	_faad.Rels.AddRelationship(_b.RelativeFilename(_b.DocTypeDocument, "", _b.CorePropertiesType, 0), _b.CorePropertiesType)
	_faad.Rels.AddRelationship("\u0064\u006fc\u0050\u0072\u006fp\u0073\u002f\u0061\u0070\u0070\u002e\u0078\u006d\u006c", _b.ExtendedPropertiesType)
	_faad.Rels.AddRelationship("\u0077\u006f\u0072\u0064\u002f\u0064\u006f\u0063\u0075\u006d\u0065\u006et\u002e\u0078\u006d\u006c", _b.OfficeDocumentType)
	_faad.Numbering = NewNumbering()
	_faad.Numbering.InitializeDefault()
	_faad.ContentTypes.AddOverride("\u002f\u0077\u006f\u0072d/\u006e\u0075\u006d\u0062\u0065\u0072\u0069\u006e\u0067\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069c\u0061\u0074\u0069\u006f\u006e\u002f\u0076n\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063e\u0073\u0073\u0069\u006e\u0067\u006d\u006c\u002e\u006e\u0075\u006d\u0062e\u0072\u0069\u006e\u0067\u002b\u0078m\u006c")
	_faad._dfc.AddRelationship("\u006e\u0075\u006d\u0062\u0065\u0072\u0069\u006e\u0067\u002e\u0078\u006d\u006c", _b.NumberingType)
	_faad.Styles = NewStyles()
	_faad.Styles.InitializeDefault()
	_faad.ContentTypes.AddOverride("\u002f\u0077o\u0072\u0064\u002fs\u0074\u0079\u006c\u0065\u0073\u002e\u0078\u006d\u006c", "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064.\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0077\u006f\u0072\u0064\u0070\u0072\u006f\u0063\u0065\u0073\u0073\u0069n\u0067\u006d\u006c\u002e\u0073\u0074\u0079\u006ce\u0073\u002b\u0078\u006d\u006c")
	_faad._dfc.AddRelationship("\u0073\u0074\u0079\u006c\u0065\u0073\u002e\u0078\u006d\u006c", _b.StylesType)
	_faad._adcb.Body = _db.NewCT_Body()
	return _faad
}

// Text returns the underlying text in the run.
func (_bfec Run) Text() string {
	if len(_bfec._afec.EG_RunInnerContent) == 0 {
		return ""
	}
	_dbeb := _fed.Buffer{}
	for _, _fbbb := range _bfec._afec.EG_RunInnerContent {
		if _fbbb.T != nil {
			_dbeb.WriteString(_fbbb.T.Content)
		}
		if _fbbb.Tab != nil {
			_dbeb.WriteByte('\t')
		}
	}
	return _dbeb.String()
}

// GetImage returns the ImageRef associated with an AnchoredDrawing.
func (_af AnchoredDrawing) GetImage() (_ffa.ImageRef, bool) {
	_gdg := _af._cg.Graphic.GraphicData.Any
	if len(_gdg) > 0 {
		_de, _fgf := _gdg[0].(*_ba.Pic)
		if _fgf {
			if _de.BlipFill != nil && _de.BlipFill.Blip != nil && _de.BlipFill.Blip.EmbedAttr != nil {
				return _af._geb.GetImageByRelID(*_de.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return _ffa.ImageRef{}, false
}

// AbstractNumberID returns the ID that is unique within all numbering
// definitions that is used to assign the definition to a paragraph.
func (_aebbb NumberingDefinition) AbstractNumberID() int64 { return _aebbb._bbgd.AbstractNumIdAttr }
func (_bfda *Document) onNewRelationship(_fdbd *_ed.DecodeMap, _gebec, _fbdf string, _egaa []*_a.File, _dbcdc *_gf.Relationship, _ggbb _ed.Target) error {
	_beeaa := _b.DocTypeDocument
	switch _fbdf {
	case _b.OfficeDocumentType, _b.OfficeDocumentTypeStrict:
		_bfda._adcb = _db.NewDocument()
		_fdbd.AddTarget(_gebec, _bfda._adcb, _fbdf, 0)
		_fdbd.AddTarget(_ed.RelationsPathFor(_gebec), _bfda._dfc.X(), _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.CorePropertiesType:
		_fdbd.AddTarget(_gebec, _bfda.CoreProperties.X(), _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.CustomPropertiesType:
		_fdbd.AddTarget(_gebec, _bfda.CustomProperties.X(), _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.ExtendedPropertiesType, _b.ExtendedPropertiesTypeStrict:
		_fdbd.AddTarget(_gebec, _bfda.AppProperties.X(), _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.ThumbnailType, _b.ThumbnailTypeStrict:
		for _ebg, _ggcfc := range _egaa {
			if _ggcfc == nil {
				continue
			}
			if _ggcfc.Name == _gebec {
				_cggf, _eacb := _ggcfc.Open()
				if _eacb != nil {
					return _g.Errorf("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0068\u0075m\u0062\u006e\u0061i\u006c:\u0020\u0025\u0073", _eacb)
				}
				_bfda.Thumbnail, _, _eacb = _dg.Decode(_cggf)
				_cggf.Close()
				if _eacb != nil {
					return _g.Errorf("\u0065\u0072\u0072\u006fr\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020t\u0068u\u006d\u0062\u006e\u0061\u0069\u006c\u003a \u0025\u0073", _eacb)
				}
				_egaa[_ebg] = nil
			}
		}
	case _b.SettingsType, _b.SettingsTypeStrict:
		_fdbd.AddTarget(_gebec, _bfda.Settings.X(), _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.NumberingType, _b.NumberingTypeStrict:
		_bfda.Numbering = NewNumbering()
		_fdbd.AddTarget(_gebec, _bfda.Numbering.X(), _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.StylesType, _b.StylesTypeStrict:
		_bfda.Styles.Clear()
		_fdbd.AddTarget(_gebec, _bfda.Styles.X(), _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.HeaderType, _b.HeaderTypeStrict:
		_bcc := _db.NewHdr()
		_fdbd.AddTarget(_gebec, _bcc, _fbdf, uint32(len(_bfda._adf)))
		_bfda._adf = append(_bfda._adf, _bcc)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, len(_bfda._adf))
		_badc := _ffa.NewRelationships()
		_fdbd.AddTarget(_ed.RelationsPathFor(_gebec), _badc.X(), _fbdf, 0)
		_bfda._gcdc = append(_bfda._gcdc, _badc)
	case _b.FooterType, _b.FooterTypeStrict:
		_ebeb := _db.NewFtr()
		_fdbd.AddTarget(_gebec, _ebeb, _fbdf, uint32(len(_bfda._abc)))
		_bfda._abc = append(_bfda._abc, _ebeb)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, len(_bfda._abc))
		_ceebc := _ffa.NewRelationships()
		_fdbd.AddTarget(_ed.RelationsPathFor(_gebec), _ceebc.X(), _fbdf, 0)
		_bfda._adg = append(_bfda._adg, _ceebc)
	case _b.ThemeType, _b.ThemeTypeStrict:
		_gdce := _aa.NewTheme()
		_fdbd.AddTarget(_gebec, _gdce, _fbdf, uint32(len(_bfda._cdf)))
		_bfda._cdf = append(_bfda._cdf, _gdce)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, len(_bfda._cdf))
	case _b.WebSettingsType, _b.WebSettingsTypeStrict:
		_bfda._eaf = _db.NewWebSettings()
		_fdbd.AddTarget(_gebec, _bfda._eaf, _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.FontTableType, _b.FontTableTypeStrict:
		_bfda._begg = _db.NewFonts()
		_fdbd.AddTarget(_gebec, _bfda._begg, _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.EndNotesType, _b.EndNotesTypeStrict:
		_bfda._acdc = _db.NewEndnotes()
		_fdbd.AddTarget(_gebec, _bfda._acdc, _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.FootNotesType, _b.FootNotesTypeStrict:
		_bfda._ffcef = _db.NewFootnotes()
		_fdbd.AddTarget(_gebec, _bfda._ffcef, _fbdf, 0)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, 0)
	case _b.ImageType, _b.ImageTypeStrict:
		var _bgg _ffa.ImageRef
		for _dfdfb, _cdae := range _egaa {
			if _cdae == nil {
				continue
			}
			_ggfc := _eg.TrimPrefix(_cdae.Name, "\u0077\u006f\u0072d\u002f")
			if _eadg := _eg.TrimPrefix(_gebec, "\u0077\u006f\u0072d\u002f"); _ggfc == _eadg {
				_edab, _bdaed := _ed.ExtractToDiskTmp(_cdae, _bfda.TmpPath)
				if _bdaed != nil {
					return _bdaed
				}
				_aceed, _bdaed := _ffa.ImageFromStorage(_edab)
				if _bdaed != nil {
					return _bdaed
				}
				_bgg = _ffa.MakeImageRef(_aceed, &_bfda.DocBase, _bfda._dfc)
				_egaa[_dfdfb] = nil
			}
		}
		if _bgg.Format() != "" {
			_gdfa := "\u002e" + _eg.ToLower(_bgg.Format())
			_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, len(_bfda.Images)+1)
			if _gff := _fc.Ext(_dbcdc.TargetAttr); _gff != _gdfa {
				_dbcdc.TargetAttr = _dbcdc.TargetAttr[0:len(_dbcdc.TargetAttr)-len(_gff)] + _gdfa
			}
			_bgg.SetTarget("\u0077\u006f\u0072d\u002f" + _dbcdc.TargetAttr)
			_bfda.Images = append(_bfda.Images, _bgg)
		}
	case _b.ControlType, _b.ControlTypeStrict:
		_ccab := _gd.NewOcx()
		_ccba := _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, len(_bfda._aede)+1)
		_gcdg := "\u0077\u006f\u0072d\u002f" + _ccba[:len(_ccba)-4] + "\u002e\u0062\u0069\u006e"
		for _bgdb, _gdfe := range _egaa {
			if _gdfe == nil {
				continue
			}
			if _gdfe.Name == _gcdg {
				_badg, _fedfg := _ed.ExtractToDiskTmp(_gdfe, _bfda.TmpPath)
				if _fedfg != nil {
					return _fedfg
				}
				_ccce, _fedfg := _ae.ImportFromFile(_badg)
				if _fedfg == nil {
					_ccce.TargetAttr = _ccba
					_ccce.Ocx = _ccab
					_bfda._aede = append(_bfda._aede, _ccce)
					_fdbd.AddTarget(_gebec, _ccab, _fbdf, uint32(len(_bfda._aede)))
					_dbcdc.TargetAttr = _ccba
					_egaa[_bgdb] = nil
				} else {
					_c.Log.Debug("c\u0061\u006e\u006e\u006f\u0074\u0020r\u0065\u0061\u0064\u0020\u0062\u0069\u006e\u0020\u0066i\u006c\u0065\u003a \u0025s\u0020\u0025\u0073", _gcdg, _fedfg.Error())
				}
				break
			}
		}
	case _b.ChartType:
		_gdcc := chart{_bbd: _ee.NewChartSpace()}
		_dfgg := uint32(len(_bfda._beag))
		_fdbd.AddTarget(_gebec, _gdcc._bbd, _fbdf, _dfgg)
		_bfda._beag = append(_bfda._beag, &_gdcc)
		_dbcdc.TargetAttr = _b.RelativeFilename(_beeaa, _ggbb.Typ, _fbdf, len(_bfda._beag))
		_gdcc._ddc = _dbcdc.TargetAttr
	default:
		_c.Log.Debug("\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u0020\u0074\u0067\u0074\u003a\u0020\u0025\u0073", _fbdf, _gebec)
	}
	return nil
}

// X returns the inner wrapped XML type.
func (_dfbb Fonts) X() *_db.CT_Fonts { return _dfbb._bdge }

// RemoveRun removes a child run from a paragraph.
func (_cfaa Paragraph) RemoveRun(r Run) {
	for _, _cfcgc := range _cfaa._dcfbd.EG_PContent {
		for _cfbga, _gcgee := range _cfcgc.EG_ContentRunContent {
			if _gcgee.R == r._afec {
				copy(_cfcgc.EG_ContentRunContent[_cfbga:], _cfcgc.EG_ContentRunContent[_cfbga+1:])
				_cfcgc.EG_ContentRunContent = _cfcgc.EG_ContentRunContent[0 : len(_cfcgc.EG_ContentRunContent)-1]
			}
			if _gcgee.Sdt != nil && _gcgee.Sdt.SdtContent != nil {
				for _gfeb, _fdfgf := range _gcgee.Sdt.SdtContent.EG_ContentRunContent {
					if _fdfgf.R == r._afec {
						copy(_gcgee.Sdt.SdtContent.EG_ContentRunContent[_gfeb:], _gcgee.Sdt.SdtContent.EG_ContentRunContent[_gfeb+1:])
						_gcgee.Sdt.SdtContent.EG_ContentRunContent = _gcgee.Sdt.SdtContent.EG_ContentRunContent[0 : len(_gcgee.Sdt.SdtContent.EG_ContentRunContent)-1]
					}
				}
			}
		}
	}
}

// AddImageRef add ImageRef to header as relationship, returning ImageRef
// that can be used to be placed as header content.
func (_gbee Header) AddImageRef(r _ffa.ImageRef) (_ffa.ImageRef, error) {
	var _cbcc _ffa.Relationships
	for _ddeg, _egdc := range _gbee._bgac._adf {
		if _egdc == _gbee._ecead {
			_cbcc = _gbee._bgac._gcdc[_ddeg]
		}
	}
	_ggfca := _cbcc.AddRelationship(r.Target(), _b.ImageType)
	r.SetRelID(_ggfca.X().IdAttr)
	return r, nil
}

// SetAll sets all of the borders to a given value.
func (_cebfe ParagraphBorders) SetAll(t _db.ST_Border, c _bd.Color, thickness _bc.Distance) {
	_cebfe.SetBottom(t, c, thickness)
	_cebfe.SetLeft(t, c, thickness)
	_cebfe.SetRight(t, c, thickness)
	_cebfe.SetTop(t, c, thickness)
}

// SetThemeColor sets the color from the theme.
func (_fccb Color) SetThemeColor(t _db.ST_ThemeColor) { _fccb._fcd.ThemeColorAttr = t }

// SetLineSpacing sets the spacing between lines in a paragraph.
func (_bbbf Paragraph) SetLineSpacing(d _bc.Distance, rule _db.ST_LineSpacingRule) {
	_bbbf.ensurePPr()
	if _bbbf._dcfbd.PPr.Spacing == nil {
		_bbbf._dcfbd.PPr.Spacing = _db.NewCT_Spacing()
	}
	_fcee := _bbbf._dcfbd.PPr.Spacing
	if rule == _db.ST_LineSpacingRuleUnset {
		_fcee.LineRuleAttr = _db.ST_LineSpacingRuleUnset
		_fcee.LineAttr = nil
	} else {
		_fcee.LineRuleAttr = rule
		_fcee.LineAttr = &_db.ST_SignedTwipsMeasure{}
		_fcee.LineAttr.Int64 = _b.Int64(int64(d / _bc.Twips))
	}
}

// ComplexSizeMeasure returns font with its measure which can be mm, cm, in, pt, pc or pi.
func (_acbaa RunProperties) ComplexSizeMeasure() string {
	if _agebe := _acbaa._cadb.SzCs; _agebe != nil {
		_eefaa := _agebe.ValAttr
		if _eefaa.ST_PositiveUniversalMeasure != nil {
			return *_eefaa.ST_PositiveUniversalMeasure
		}
	}
	return ""
}

// SetHangingIndent controls the indentation of the non-first lines in a paragraph.
func (_baaf ParagraphProperties) SetHangingIndent(m _bc.Distance) {
	if _baaf._aedc.Ind == nil {
		_baaf._aedc.Ind = _db.NewCT_Ind()
	}
	if m == _bc.Zero {
		_baaf._aedc.Ind.HangingAttr = nil
	} else {
		_baaf._aedc.Ind.HangingAttr = &_ca.ST_TwipsMeasure{}
		_baaf._aedc.Ind.HangingAttr.ST_UnsignedDecimalNumber = _b.Uint64(uint64(m / _bc.Twips))
	}
}

// X returns the inner wml.CT_PBdr
func (_fbeg ParagraphBorders) X() *_db.CT_PBdr { return _fbeg._afaaa }
func _dagg() *_eb.Textpath {
	_abgd := _eb.NewTextpath()
	_bcbg := "\u0066\u006f\u006e\u0074\u002d\u0066\u0061\u006d\u0069l\u0079\u003a\u0022\u0043\u0061\u006c\u0069\u0062\u0072\u0069\u0022\u003b\u0066\u006f\u006e\u0074\u002d\u0073\u0069\u007a\u0065\u003a\u00366\u0070\u0074;\u0066\u006fn\u0074\u002d\u0077\u0065\u0069\u0067\u0068\u0074\u003a\u0062\u006f\u006c\u0064;f\u006f\u006e\u0074\u002d\u0073\u0074\u0079\u006c\u0065:\u0069\u0074\u0061\u006c\u0069\u0063"
	_abgd.StyleAttr = &_bcbg
	_eaaea := "\u0041\u0053\u0041\u0050"
	_abgd.StringAttr = &_eaaea
	return _abgd
}

// X return slice of node.
func (_abee *Nodes) X() []Node { return _abee._dcdf }

// NumId return numbering numId that being use by style properties.
func (_ccede ParagraphStyleProperties) NumId() int64 {
	if _ccede._fagf.NumPr != nil {
		if _ccede._fagf.NumPr.NumId != nil {
			return _ccede._fagf.NumPr.NumId.ValAttr
		}
	}
	return -1
}

// RemoveParagraph removes a paragraph from the endnote.
func (_bage Endnote) RemoveParagraph(p Paragraph) {
	for _, _bfgac := range _bage.content() {
		for _gabe, _dcba := range _bfgac.P {
			if _dcba == p._dcfbd {
				copy(_bfgac.P[_gabe:], _bfgac.P[_gabe+1:])
				_bfgac.P = _bfgac.P[0 : len(_bfgac.P)-1]
				return
			}
		}
	}
}

// OpenTemplate opens a document, removing all content so it can be used as a
// template.  Since Word removes unused styles from a document upon save, to
// create a template in Word add a paragraph with every style of interest.  When
// opened with OpenTemplate the document's styles will be available but the
// content will be gone.
func OpenTemplate(filename string) (*Document, error) {
	_fbaf, _eecc := Open(filename)
	if _eecc != nil {
		return nil, _eecc
	}
	_fbaf._adcb.Body = _db.NewCT_Body()
	return _fbaf, nil
}

// SetToolTip sets the tooltip text for a hyperlink.
func (_eaedb HyperLink) SetToolTip(text string) {
	if text == "" {
		_eaedb._cfcf.TooltipAttr = nil
	} else {
		_eaedb._cfcf.TooltipAttr = _b.String(text)
	}
}

// ClearContent clears any content in the run (text, tabs, breaks, etc.)
func (_bgdf Run) ClearContent() { _bgdf._afec.EG_RunInnerContent = nil }

// Clear content of node element.
func (_cgbc *Node) Clear() { _cgbc._bdcf = nil }

// SetFooter sets a section footer.
func (_gbdabc Section) SetFooter(f Footer, t _db.ST_HdrFtr) {
	_cfdga := _db.NewEG_HdrFtrReferences()
	_gbdabc._ggcfb.EG_HdrFtrReferences = append(_gbdabc._ggcfb.EG_HdrFtrReferences, _cfdga)
	_cfdga.FooterReference = _db.NewCT_HdrFtrRef()
	_cfdga.FooterReference.TypeAttr = t
	_fgdeb := _gbdabc._bagegb._dfc.FindRIDForN(f.Index(), _b.FooterType)
	if _fgdeb == "" {
		_c.Log.Debug("\u0075\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0072\u006d\u0069\u006e\u0065\u0020\u0066\u006f\u006f\u0074\u0065r \u0049\u0044")
	}
	_cfdga.FooterReference.IdAttr = _fgdeb
}
