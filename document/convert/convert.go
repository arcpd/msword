package convert

import (
	_d "bytes"
	_ce "errors"
	_da "github.com/arcpd/msword/common/logger"
	_f "github.com/arcpd/msword/common/tempstorage"
	_e "github.com/arcpd/msword/document"
	_a "github.com/arcpd/msword/internal/convertutils"
	_ge "github.com/arcpd/msword/internal/formatutils"
	_eag "github.com/arcpd/msword/measurement"
	_bb "github.com/arcpd/msword/schema/soo/dml"
	_gg "github.com/arcpd/msword/schema/soo/dml/chart"
	_eeb "github.com/arcpd/msword/schema/soo/dml/picture"
	_ee "github.com/arcpd/msword/schema/soo/ofc/sharedTypes"
	_bc "github.com/arcpd/msword/schema/soo/wml"
	_fff "github.com/arcpd/msword/schema/urn/schemas_microsoft_com/vml"
	_ff "github.com/arcpd/msword/vmldrawing"
	_ea "github.com/unidoc/unipdf/v3/creator"
	_gc "github.com/unidoc/unipdf/v3/model"
	_bg "io/ioutil"
	_bf "regexp"
	_b "strconv"
	_g "strings"
)

func (_eaeb *convertContext) addCurrentParagraphToCurrentPage() {
	_eaeb._fbcd._dbg = _eaeb._cecg._bgb + _eaeb._cecg._ac.Top + _eaeb._cecg._egd + _eaeb._cecg._ac.Bottom
	_eaeb._fbcd._fc = append(_eaeb._fbcd._fc, _eaeb._cecg._ebd...)
	_eaeb._fbcd._gec = append(_eaeb._fbcd._gec, _eaeb._cecg._caf...)
	_eaeb._fbcd._ceb = append(_eaeb._fbcd._ceb, _eaeb._cecg._gb...)
	_eaeb._fbcd._ab = append(_eaeb._fbcd._ab, _eaeb._cecg._bfb...)
	_eaeb._fbcd._gf = append(_eaeb._fbcd._gf, _eaeb._cecg._fcd...)
	_eaeb._fbcd._eb = append(_eaeb._fbcd._eb, _eaeb._cecg)
	_eaeb.adjustRightBoundOfLastSpan()
	_eaeb.alignParagraph()
	if len(_eaeb._fbcd._eaf) == 0 && len(_eaeb._cecg._gbg) > 0 {
		_eaeb._fbcd._eeg.Bottom -= _ga
	}
	_eaeb._fbcd._eaf = append(_eaeb._fbcd._eaf, _eaeb._cecg._gbg...)
	_eaeb._fbcd._eeg.Bottom -= _eaeb._cecg._cd
}

// RegisterFont makes a PdfFont accessible for using in converting to PDF.
func RegisterFont(name string, style FontStyle, font *_gc.PdfFont) {
	_a.RegisterFont(name, style, font)
}

func (_ebe *convertContext) addInlineSymbol(_fcfcc *symbol) {
	if len(_ebe._gced._gaa) > 0 {
		_gcdb := _ebe._gced._gaa[len(_ebe._gced._gaa)-1]._fge
		if _gcdb == "\u0020" {
			_ebe.addCurrentWordToParagraph()
			_ebe.newWord()
		}
	}
	_ebe._gced._gaa = append(_ebe._gced._gaa, _fcfcc)
	_fcfcc._cdb = _ebe._gced._bfdc
	_ebe._gced._bfdc += _fcfcc._cdf
	_ebe._gced._gfb = false
	_ebe.adjustHeights(_fcfcc._cgb)
}

// Options contains the options for convert process -
// e.g ProcessFields is when document contains fields
// and the value need to be processed also.
type Options struct {

	// ProcessFields process the document fields if set to true, default is `false`.
	ProcessFields bool

	// EnableFontSubsetting process document with subsetting font to reduce size result.
	// Default value is `true`.
	EnableFontSubsetting bool

	// FontFiles location of fonts for convert process.
	FontFiles []string

	// FontDirectory location of font directory for convert process.
	// This will load all font files inside directoy if set
	// we recommend to use FontFiles for better performance.
	FontDirectory string
}

func _dcad(_ecdg *_e.Document, _fbbge *_bc.CT_TblPr) (*_bc.CT_TblPr, *_bc.CT_PPrGeneral, *_bc.CT_RPr) {
	_adgb := _bc.NewCT_PPrGeneral()
	_dddgc := _bc.NewCT_RPr()
	if _fbbge == nil {
		_fbbge = _bc.NewCT_TblPr()
	} else {
		if _fbbge.TblStyle != nil {
			_fbbge, _adgb, _dddgc = _dccef(_ecdg, _fbbge.TblStyle.ValAttr, _fbbge, _adgb, _dddgc)
		}
	}
	return _fbbge, _adgb, _dddgc
}

func _ebddf(_dfcc int, _gecga bool) string {
	_afdg := _d.NewBuffer([]byte{})
	for _, _gdgee := range _gdff {
		for {
			if _dfcc < _gdgee._bcacd {
				break
			}
			_afdg.WriteString(_gdgee._fbgb)
			_dfcc -= _gdgee._bcacd
		}
	}
	_gcbc := _afdg.String()
	if _gecga {
		_gcbc = _g.ToUpper(_gcbc)
	}
	return _gcbc
}

type convertContext struct {
	_fcgg   *_ea.Creator
	_cffa   *_e.Document
	_bcdegd *_bc.CT_PPrGeneral
	_fbef   *_bc.CT_RPr
	_efff   *_ea.StyledParagraph
	_efceb  []*page
	_fbcd   *page
	_eaebc  *_a.Rectangle
	_cecg   *paragraph
	_bbfa   *line
	_ecbg   *span
	_gced   *word
	_bgeb   *_bc.CT_Hyperlink
	_acgg   *_bc.CT_ParaRPr
	_gffd   *_bc.CT_PPr
	_agae   []note
	_eegda  *prefix
	_ecfg   bool
	_cbae   bool
	_bcdd   bool
	_bcce   float64
	_ggfa   float64
	_cfbg   float64
	_fcfcf  float64
	_dade   bool
	_fbgf   map[int64]map[int64]int64
	_cgaa   map[string]string
	_afff   *Options
	_dceg   []*headerFooterRef
	_gfe    []*paragraph
	_fcbc   []*paragraph
}

var _debg float64

// FontStyle represents a kind of font styling. It can be FontStyle_Regular, FontStyle_Bold, FontStyle_Italic and FontStyle_BoldItalic.
type FontStyle = _a.FontStyle

func _bga(_fedg string) (string, string) {
	_bbdf := _abac.FindStringSubmatch(_fedg)
	if len(_bbdf) < 3 {
		return "", ""
	}
	return _bbdf[1], _bbdf[2]
}

func (_bfae *convertContext) addAbsoluteHeaderFooterTable(_fbad *_bc.CT_Tbl) {
	_cbcb := _fbad.TblGrid
	if _cbcb == nil {
		return
	}
	_egef := len(_cbcb.GridCol)
	if _egef == 0 {
		return
	}
	_dgcd := []float64{}
	_cac := []float64{}
	_dbbfe := 0.0
	for _, _gcfgg := range _cbcb.GridCol {
		_eda := 0.0
		if _gcfgg.WAttr.ST_UnsignedDecimalNumber != nil {
			_eda = _a.PointsFromTwips(int64(*_gcfgg.WAttr.ST_UnsignedDecimalNumber))
		}
		_dgcd = append(_dgcd, _eda)
		_dbbfe += _eda
	}
	for _ccage := 0; _ccage < _egef; _ccage++ {
		_cac = append(_cac, _dgcd[_ccage]/_dbbfe)
	}
	_fcae := _bfae._fcgg.NewTable(_egef)
	_fcae.SetColumnWidths(_cac...)
	_agdg := _bfae._fcgg.NewTable(_egef)
	_agdg.SetColumnWidths(_cac...)
	_fdab, _dadc, _fcdb := _dcad(_bfae._cffa, _fbad.TblPr)
	var _egee []*_bc.CT_TblStylePr
	if _fdab.TblStyle != nil {
		_egee = _bcffb(_bfae._cffa, _fdab.TblStyle.ValAttr)
	}
	_ffga := _gafdd(_fdab.TblW, _bfae._fbcd._eeg.Right-_bfae._fbcd._eeg.Left, 0)
	_fcba := _gafdd(_fdab.TblInd, _bfae._fbcd._eeg.Right-_bfae._fbcd._eeg.Left, 0)
	_ebdd := _bfae._fbcd._eeg.Bottom - _bfae._cecg._bgb
	_bdae := len(_fbad.EG_ContentRowContent)
	for _bcdbf, _gcfad := range _fbad.EG_ContentRowContent {
		if _gcfad == nil {
			continue
		}
		_ccb := _bfae._fcgg.NewTable(_egef)
		_ccb.SetColumnWidths(_cac...)
		if _aae := _gcfad.Tr; len(_aae) > 0 {
			_gggf := _aae[0]
			_acceg := _gggf.TblPrEx
			for _eafg, _eggeg := range _gggf.EG_ContentCellContent {
				if _cace := _eggeg.Tc; len(_cace) > 0 {
					if _fce := _cace[0]; _fce != nil {
						_bfae.addCellToTable(_agdg, _fce, _fdab, _acceg, _bcdbf, _eafg, _bdae, _egef, _egee, _dadc, _fcdb, false)
						_bfae.addCellToTable(_ccb, _fce, _fdab, _acceg, _bcdbf, _eafg, _bdae, _egef, _egee, _dadc, _fcdb, false)
					}
				}
			}
			var _dddf float64
			if _acfe := _gggf.TrPr; _acfe != nil {
				if len(_acfe.TrHeight) != 0 {
					_gdge := _acfe.TrHeight[0]
					if _bacb := _gdge.ValAttr; _bacb != nil {
						if _bacb.ST_UnsignedDecimalNumber != nil {
							_dddf = _a.PointsFromTwips(int64(*_bacb.ST_UnsignedDecimalNumber))
						}
					}
				}
			}
			if _dddf < _ccb.Height() {
				_dddf = _ccb.Height()
			}
			if _dddf < _ffcdc(4) {
				_dddf = _ffcdc(4)
			}
			_agdg.SetRowHeight(_agdg.CurRow(), _dddf)
			_ccb.SetRowHeight(_ccb.CurRow(), _dddf)
			if _ffga == 0 || _ffga > _bfae._fbcd._eeg.Right-_bfae._fbcd._eeg.Left {
				_ffga = _bfae._fbcd._eeg.Right - _bfae._fbcd._eeg.Left
			}
			_bbcd := _a.MakeTempCreator(_ffga, _ffcdc(1000))
			_bbcd.Draw(_agdg)
			if _agdg.Height() >= _ebdd {
				_bfae.addParagraphWithTable(*_fcae, _ffga, _fcba)
				_bfae.newPage()
				*_agdg = *_ccb
				_agdg.SetRowHeight(_agdg.CurRow(), _dddf)
				_ebdd = _bfae._fbcd._eeg.Bottom - _bfae._fbcd._eeg.Top
				_fcae = nil
			} else {
				if _fcae == nil {
					_fcae = _bfae._fcgg.NewTable(_egef)
					_fcae.SetColumnWidths(_cac...)
				}
				*_fcae = *_ccb
			}
		}
	}
	if _fcae != nil {
		_bfae.addParagraphWithTableToHeaderFooter(*_fcae, _ffga, _fcba)
	}
}

func _fcbcd(_ddedc *_ea.Creator, _fcdae *_ea.Block, _cdbg []*paragraph, _ecab, _dfdgc float64) {
	_afac := 0.0
	for _decc, _gfgb := range _cdbg {
		if _decc > 0 {
			_afac += _gfgb._egd
		}
		for _, _egcf := range _gfgb._cc {
			for _, _fefdc := range _egcf._dag {
				for _, _fabe := range _fefdc._bda {
					for _, _fgb := range _fabe._gaa {
						if _fgb._bea != nil {
							_fgb._bea.SetPos(_fabe._cfe+_fgb._cdb, _ecab)
							_fcdae.Draw(_fgb._bea)
						} else if _fgb._cce != nil {
							if _fgb._cce._aef == 0 {
								_fgb._cce._aef = _fabe._cfe + _fgb._cdb
							}
							if _fgb._cce._aga == 0 {
								_fgb._cce._aga = _gfgb._bgb + _egcf._ag
							}
							_dab(_ddedc, _fgb._cce)
						} else {
							_geefa := _ddedc.NewStyledParagraph()
							if _fgb._ffa {
								_fgb._afa = 0
							} else if _fgb._bcb {
								_fgb._afa = 1.2*_egcf._bca - _fgb._cgb
							}
							_cfgfe := _fabe._cfe + _fgb._cdb
							_befg := _ecab + _egcf._ag + _fgb._afa + _afac
							_geefa.SetPos(_cfgfe, _befg)
							var _cabg *_ea.TextChunk
							if _fgb._eae != "" {
								_cabg = _geefa.AddExternalLink(_fgb._fge, _fgb._eae)
							} else {
								_cabg = _geefa.Append(_fgb._fge)
							}
							if _fgb._ba != nil {
								_cabg.Style = *_fgb._ba
							}
							_fcdae.Draw(_geefa)
							if _fgb._ad != nil {
								_baff := _dfdgc + _ecab - _afac + 2.0
								_a.DrawLine(_ddedc, _cfgfe, _baff, _cfgfe+_fgb._cdf, _baff, 1, *_fgb._ad)
							}
						}
					}
				}
			}
		}
		if _gfgb._fdg != nil {
			_cbfc := _ea.NewBlock(_gfgb._fdg._fb, _gfgb._ac.Top+_gfgb._egd+_gfgb._ac.Bottom)
			_cbfc.SetPos(_gfgb._dg, _ecab)
			_cbfc.Draw(_gfgb._fdg._cca)
			_fcdae.Draw(_cbfc)
		}
	}
}

type note struct {
	_bce string
	_ec  []*_bc.EG_BlockLevelElts
	_be  *_ea.Block
}

func _fe(_cdd *_ea.Creator, _agd *image) {
	_agd._gbb.SetPos(_agd._ddc, _agd._bgg)
	_cdd.Draw(_agd._gbb)
}

func _bgfa(_ccgc *_e.Document) map[string]string {
	_bffe := map[string]string{}
	for _, _gcdg := range _ccgc.Paragraphs() {
		for _, _abde := range _gcdg.Runs() {
			for _, _gbec := range _abde.X().EG_RunInnerContent {
				if _ggcaf := _gbec.InstrText; _ggcaf != nil {
					_eeff, _fegc := _bga(_ggcaf.Content)
					if _eeff != "" && _fegc != "" {
						_bffe[_eeff] = _fegc
					}
				}
			}
		}
	}
	return _bffe
}

// ConvertToPdf converts document to PDF file. This package is beta, breaking changes can take place.
func ConvertToPdf(d *_e.Document) *_ea.Creator { return ConvertToPdfWithOptions(d, nil) }

func (_ggge *convertContext) assignPropsToRelativeParagraph(_fddc *_bc.CT_PPr, _gefbf *_ea.StyledParagraph) (float64, float64) {
	_fddc = _cceb(_fddc, _ggge._bcdegd, _ggge._fbef)
	_fgba := 1.1
	if _fddc == nil {
		_gefbf.SetLineHeight(_fgba)
		return 0, 0
	}
	var _cbaea _ea.TextAlignment
	if _fddc.Jc != nil {
		switch _fddc.Jc.ValAttr {
		case _bc.ST_JcRight:
			_cbaea = _ea.TextAlignmentRight
		case _bc.ST_JcCenter:
			_cbaea = _ea.TextAlignmentCenter
		case _bc.ST_JcBoth:
			_cbaea = _ea.TextAlignmentJustify
		case _bc.ST_JcEnd:
			_cbaea = _ea.TextAlignmentRight
		default:
			_cbaea = _ea.TextAlignmentLeft
		}
		_gefbf.SetTextAlignment(_cbaea)
	}
	var _caeb, _efeg, _aaee, _ebad float64
	if _adcg := _fddc.Spacing; _adcg != nil {
		if _fdbff := _adcg.BeforeAttr; _fdbff != nil {
			if _fdbff.ST_UnsignedDecimalNumber != nil {
				_caeb = _a.PointsFromTwips(int64(*_fdbff.ST_UnsignedDecimalNumber))
			}
		}
		if _gfdfe := _adcg.AfterAttr; _gfdfe != nil {
			if _gfdfe.ST_UnsignedDecimalNumber != nil {
				_efeg = _a.PointsFromTwips(int64(*_gfdfe.ST_UnsignedDecimalNumber))
			}
		}
		if _ffcd := _adcg.LineAttr; _ffcd != nil {
			if _ffcd.Int64 != nil {
				_fgba = float64(*_ffcd.Int64 / 240)
			}
		}
	}
	if _fddc.ContextualSpacing != nil && _dgfgag(_fddc.ContextualSpacing) {
		_caeb = 0
		_efeg = 0
	}
	if _ddcd := _fddc.TextAlignment; _ddcd != nil {
		switch _ddcd.ValAttr {
		case _bc.ST_TextAlignmentTop:
			_caeb = (_fgba - (_caeb + _efeg)) * 0.5
		}
	}
	if _cbaea == _ea.TextAlignmentRight && _ebad <= 0 {
		_ebad += 5
	}
	if _caeb > 0 {
		_caeb = _caeb - _fgba/2
	}
	if _efeg > 0 {
		_efeg = _efeg - _fgba/2
	}
	_gefbf.SetLineHeight(_fgba)
	if _dba := _fddc.Ind; _dba != nil {
		if _ecbb := _dba.LeftAttr; _ecbb != nil {
			if _ecbb.Int64 != nil {
				_aaee = _a.PointsFromTwips(*_ecbb.Int64)
			}
		}
		if _bbbg := _dba.RightAttr; _bbbg != nil {
			if _bbbg.Int64 != nil {
				_ebad = _a.PointsFromTwips(*_bbbg.Int64)
			}
		}
	}
	_gefbf.SetMargins(_aaee, _ebad, _caeb, _efeg)
	return _caeb, _aaee
}

func (_bfa *convertContext) addAbsoluteRIC(_efb *_bc.EG_RunInnerContent, _dgfg *_bc.CT_RPr) bool {
	var _bfcb, _bgce bool
	_cdbb := []*symbol{}
	_abc := false
	if _efb == nil {
		if _bfa._eegda != nil {
			_ecbcd := true
			for _, _addb := range _bfa._eegda._ebab {
				if _fgeg, _eegd := _dcd[_addb]; _eegd {
					_bgce = _bfa._eegda._aedd
					_bfa._eegda._ebab = string(rune(_fgeg))
					_ecbcd = false
				}
			}
			_cdbb = _fgef(_bfa._eegda._ebab, "", true, false, _ecbcd)
		}
	} else {
		if _aecg(_efb) {
			return true
		} else if _efb.T != nil && _efb.T.Content != "" {
			_fffb := _efb.T.Content
			if _dgfg != nil && _dgfgag(_dgfg.Caps) {
				_fffb = _g.ToUpper(_fffb)
			}
			if _egc := _bfa._bgeb; _egc != nil && _egc.IdAttr != nil {
				_abc = true
				_cdbb = _fgef(_fffb, _bfa._cffa.GetTargetByRelId(*_egc.IdAttr), false, false, false)
			} else {
				_cdbb = _fgef(_fffb, "", false, false, false)
			}
		} else if _dgbc := _efb.EndnoteReference; _dgbc != nil {
			_caa := _bfa._cffa.BodySection().X()
			_ageg := _dgbc.IdAttr
			_bcab := _ageg
			_fgc := _bc.ST_NumberFormatLowerRoman
			if _gba := _caa.EndnotePr; _gba != nil {
				if _dega := _gba.NumFmt; _dega != nil {
					_fgc = _dega.ValAttr
				}
				if _bcdb := _gba.NumStart; _bcdb != nil {
					_bcab += _bcdb.ValAttr - 1
				}
			}
			_cgge := _ebgb(_bcab, _fgc)
			_abeb := _bfa._cffa.Endnote(_ageg).X()
			if _abeb != nil {
				_bfa._agae = append(_bfa._agae, note{_bce: _cgge, _ec: _abeb.EG_BlockLevelElts})
				_cdbb = _fgef(_cgge, "", true, false, false)
			}
		} else if _ddbg := _efb.FootnoteReference; _ddbg != nil {
			_fegg := _bfa._cffa.BodySection().X()
			_fac := _ddbg.IdAttr
			_bbdb := _fac
			_deed := _bc.ST_NumberFormatDecimal
			if _dff := _fegg.FootnotePr; _dff != nil {
				if _dgff := _dff.NumFmt; _dgff != nil {
					_deed = _dgff.ValAttr
				}
				if _efad := _dff.NumStart; _efad != nil {
					_bbdb += _efad.ValAttr - 1
				}
			}
			_gaef := _ebgb(_bbdb, _deed)
			_daed := _bfa._cffa.Footnote(_fac).X()
			if _daed != nil {
				_cbb := &note{_bce: _gaef, _ec: _daed.EG_BlockLevelElts}
				_caaf := [][]*_bc.EG_ContentBlockContent{}
				for _, _abab := range _daed.EG_BlockLevelElts {
					_caaf = append(_caaf, _abab.EG_ContentBlockContent)
				}
				_fag := &prefix{_ebab: _gaef}
				_ggg, _bdca := _bfa.makePdfBlockFromCBCs(_caaf, _bfa._fbcd._eeg.Right-_bfa._fbcd._eeg.Left, _ffcdc(1000), nil, true, _fag)
				if _bdca != nil {
					_da.Log.Debug("C\u0061\u006e\u006e\u006f\u0074\u0020c\u006f\u006e\u0076\u0065\u0072\u0074\u0020\u0066\u006fo\u0074\u006e\u006ft\u0065:\u0020\u0025\u0073", _bdca)
					return false
				}
				_cbb._be = _ggg
				_bfa._cecg._gbg = append(_bfa._cecg._gbg, _cbb)
				_bfa._cecg._cd += _cbb._be.Height()
				_cdbb = _fgef(_gaef, "", true, false, false)
			}
		} else if _eacb := _efb.InstrText; _eacb != nil {
			_dabc := _bada(_eacb.Content)
			if _dabc != "" {
				_cdbb = _fgef(_bfa._cgaa[_dabc], "", false, false, false)
			}
		} else if _cgd := _efb.Drawing; _cgd != nil {
			for _, _dcdg := range _cgd.Inline {
				if _gdd := _dcdg.Graphic; _gdd != nil {
					if _cde := _gdd.GraphicData; _cde != nil {
						_debf := _dcdg.Extent
						if _debf == nil {
							return false
						}
						_fgec := _eag.FromEMU(_debf.CxAttr)
						_deac := _eag.FromEMU(_debf.CyAttr)
						for _, _ace := range _cde.Any {
							if _debb, _cbe := _ace.(*_eeb.Pic); _cbe {
								_ddg := &symbol{_cgb: _deac, _cdf: _fgec}
								_caba, _eccf := _bfa.makePdfImageFromGraphics(_debb)
								if _eccf != nil {
									_da.Log.Debug("C\u0061\u006e\u006e\u006ft \u0072e\u0061\u0064\u0020\u0069\u006da\u0067\u0065\u003a\u0020\u0025\u0073", _eccf)
								}
								if _caba == nil {
									_ddg._fge = "\u0020"
								} else {
									_caba.Scale(_fgec/_caba.Width(), _deac/_caba.Height())
									_ddg._bea = _caba
									_bfcb = true
								}
								_cdbb = []*symbol{_ddg}
							} else if _daag, _dgde := _ace.(*_gg.Chart); _dgde {
								_gfdf := &symbol{_cgb: _deac, _cdf: _fgec}
								_bcde, _eee := _bfa.makePdfBlockFromChart(_daag, _fgec, _deac)
								if _eee != nil {
									_da.Log.Debug("C\u0061\u006e\u006e\u006ft \u0072e\u0061\u0064\u0020\u0062\u006co\u0063\u006b\u003a\u0020\u0025\u0073", _eee)
								}
								if _bcde == nil {
									_gfdf._fge = "\u0020"
								} else {
									_gfdf._cce = &block{_egb: _bcde}
									_bfcb = true
								}
								_cdbb = []*symbol{_gfdf}
							}
						}
					}
				}
			}
		} else if _fdd := _efb.Pict; _fdd != nil {
			for _, _afd := range _fdd.Any {
				if _cbf, _eca := _afd.(*_fff.Group); _eca {
					if _cbf.Rect != nil {
						for _, _dcg := range _cbf.Rect {
							_fdgd := _ff.NewShapeStyle("")
							if _dcg.StyleAttr != nil {
								_fdgd = _ff.NewShapeStyle(*_dcg.StyleAttr)
							}
							_adgd := _ea.ColorWhite
							if _dcg.FillcolorAttr != nil {
								_adgd = _ea.ColorRGBFromHex(*_dcg.FillcolorAttr)
							}
							_cbg := _a.PointsFromTwips(int64(_fdgd.Width()))
							_dfbd := _a.PointsFromTwips(int64(_fdgd.Height()))
							_fed := _a.PointsFromTwips(int64(_fdgd.Left() - _fdgd.Right()))
							_ggcd := _a.PointsFromTwips(int64(_fdgd.Top() - _fdgd.Bottom()))
							_egca := &borderLine{_ecc: _a.BorderPositionBottom, _dca: _cbg, _adc: _dfbd, _afaf: _adgd}
							_bfa._cecg._gcf = append(_bfa._cecg._gcf, _egca)
							if _fdgd.Position() == _ff.ShapeStylePositionAbsolute {
								_bfa._bbfa._dbb = _bfa._cecg._eg + _fed
								_bfa._bbfa._ag = _ggcd
							}
						}
					}
					if _cbf.Shape != nil {
						for _, _bfeg := range _cbf.Shape {
							_ffcc := _ff.NewShapeStyle("")
							if _bfeg.StyleAttr != nil {
								_ffcc = _ff.NewShapeStyle(*_bfeg.StyleAttr)
							}
							_edff := _a.PointsFromTwips(int64(_ffcc.Width()))
							_bdbf := _a.PointsFromTwips(int64(_ffcc.Height()))
							_fdfa := _a.PointsFromTwips(int64(_ffcc.Left() - _ffcc.Right()))
							_bgff := _a.PointsFromTwips(int64(_ffcc.Top() - _ffcc.Bottom()))
							if _bfeg.EG_ShapeElements != nil {
								for _, _decb := range _bfeg.EG_ShapeElements {
									if _decb.Imagedata != nil {
										_baeg := &symbol{_cgb: _edff, _cdf: _bdbf}
										_cdfd, _abebb := _bfa.makePdfImageFromRelId(_decb.Imagedata.IdAttr)
										if _abebb != nil {
											_da.Log.Debug("C\u0061\u006e\u006e\u006ft \u0072e\u0061\u0064\u0020\u0069\u006da\u0067\u0065\u003a\u0020\u0025\u0073", _abebb)
										}
										if _cdfd == nil {
											_baeg._fge = "\u0020"
										} else {
											_cdfd.Scale(_edff/_cdfd.Width(), _bdbf/_cdfd.Height())
											_cdfd.SetPos(_fdfa, _bgff)
											_baeg._bea = _cdfd
											_bfcb = true
										}
										_cdbb = []*symbol{_baeg}
										if _ffcc.Position() == _ff.ShapeStylePositionAbsolute {
											_bfa._bbfa._dbb = _bfa._cecg._eg + _fdfa
											_bfa._bbfa._ag = _bgff
										}
									}
								}
							}
						}
					}
				}
				if _agdc, _fcfc := _afd.(*_fff.Shape); _fcfc {
					_ead := _ff.NewShapeStyle("")
					if _agdc.StyleAttr != nil {
						_ead = _ff.NewShapeStyle(*_agdc.StyleAttr)
					}
					_gbab := _ea.ColorWhite
					if _agdc.StrokecolorAttr != nil {
						_gbab = _ea.ColorRGBFromHex(*_agdc.StrokecolorAttr)
					}
					if _agdc.FillcolorAttr != nil {
						_gbab = _ea.ColorRGBFromHex(*_agdc.FillcolorAttr)
					}
					_ccd := _ead.Width()
					_dcf := _ead.Height()
					_eace := _a.PointsFromTwips(int64(_ead.Left() - _ead.Right()))
					_cef := _a.PointsFromTwips(int64(_ead.Top() - _ead.Bottom()))
					_ecdd, _gcd, _eec, _egce := _ead.Margins()
					_dda := &borderLine{_ecc: _a.BorderPositionBottom, _dca: _ccd, _adc: _dcf, _afaf: _gbab}
					_bfa._cecg._gcf = append(_bfa._cecg._gcf, _dda)
					_bfa._cecg._ac = &_a.Rectangle{Top: float64(_ecdd), Left: float64(_gcd), Bottom: float64(_eec), Right: float64(_egce)}
					if _ead.Position() == _ff.ShapeStylePositionAbsolute {
						_bfa._bbfa._dbb = _bfa._cecg._eg + _eace + float64(_ead.Left())
						_bfa._bbfa._ag = _cef
					}
					var _bbcg []*symbol
					for _, _fccbb := range _agdc.EG_ShapeElements {
						if _fccbb.Textbox != nil && _fccbb.Textbox.TxbxContent != nil {
							_bdbe, _ := _bfa.makeBlockFromTextboxContent(_fccbb.Textbox.TxbxContent, _ccd, _dcf, nil)
							if _bdbe != nil {
								_ddfb := &symbol{_cgb: _dcf, _cdf: _ccd}
								if _ead.MSOPositionVerticalRelative() == "\u0070\u0061\u0067\u0065" {
									_bdbe._aga = _ecdd
								}
								if _ead.MSOPositionHorizontalRelative() == "\u0070\u0061\u0067\u0065" {
									_bdbe._aef = _gcd
								}
								_ddfb._cce = _bdbe
								_ddfb._fge = "\u0020"
								_bbcg = append(_bbcg, _ddfb)
							}
						}
					}
					if len(_bbcg) > 0 {
						_cdbb = _bbcg
					}
				}
				if _dedg, _dddd := _afd.(*_fff.Line); _dddd {
					_egbf := _ff.NewShapeStyle("")
					if _dedg.StyleAttr != nil {
						_egbf = _ff.NewShapeStyle(*_dedg.StyleAttr)
					}
					_fdda, _ebaa := 0.0, 0.0
					if _dedg.FromAttr != nil {
						_fdda, _ebaa = _baba(*_dedg.FromAttr)
					}
					_cdcff, _beda := _fdda, _ebaa
					if _dedg.ToAttr != nil {
						_cdcff, _beda = _baba(*_dedg.ToAttr)
					}
					_gaggd := _ea.ColorWhite
					if _dedg.StrokecolorAttr != nil {
						_gaggd = _ea.ColorRGBFromHex(*_dedg.StrokecolorAttr)
					}
					_ccdg := _beda - _ebaa
					if _dedg.StrokeweightAttr != nil {
						_ggcg, _cgdc := _b.ParseFloat(_g.ReplaceAll(*_dedg.StrokeweightAttr, "\u0070\u0074", ""), 64)
						if _cgdc != nil {
							_da.Log.Debug("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020\u0070a\u0072\u0073\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006f\u0066\u0020v\u003a\u006c\u0069\u006e\u0065\u0020\u0073\u0074\u0072\u006f\u006b\u0065 w\u0065\u0069\u0067\u0068\u0074\u0020\u0028\u0025\u0073\u0029", _cgdc.Error())
						}
						_ccdg = _ggcg
					}
					_def := &borderLine{_ecc: _a.BorderPositionBottom, _dca: _cdcff - _fdda, _adc: _ccdg, _afaf: _gaggd}
					_bfa._cecg._gcf = append(_bfa._cecg._gcf, _def)
					if _egbf.Position() == _ff.ShapeStylePositionAbsolute {
						_bfa._bbfa._dbb = _bfa._cecg._eg + _fdda
						_bfa._bbfa._ag = _ebaa
					}
				}
			}
		} else if _dgac := _efb.Tab; _dgac != nil {
			_bdge := 0.0
			if _bfa._gffd == nil {
				_bfa._gffd = _bc.NewCT_PPr()
			}
			if _ddcc := _bfa._gffd.Tabs; _ddcc != nil {
				_ddce := _ddcc.Tab[0]
				_bfag := _a.PointsFromTwips(*_ddce.PosAttr.Int64)
				if _ddce.ValAttr != _bc.ST_TabJcEnd && _ddce.ValAttr != _bc.ST_TabJcRight {
					_bfag += _debg
				}
				_bdge = _bfag - _bfa._cecg._af - _bfa._cecg._ac.Left - _bfa._cecg._ac.Right
				_fef := 0.0
				for _, _ffbb := range _bfa._cecg._cc {
					for _, _gdc := range _ffbb._dag {
						for _, _faf := range _gdc._bda {
							for _, _fdef := range _faf._gaa {
								_fef += _fdef._cdf
							}
						}
					}
				}
				_bdge = _bdge - _fef - _bfa._cecg._eg
				if _bdge < _debg {
					_bdge = _debg
				}
			}
			_cdbb = _fgef("\u0009", "", false, false, false)
			_dffb := _cdbb[len(_cdbb)-1]
			_dffb._cdf = _bdge
		} else if _fcfd := _efb.Ptab; _fcfd != nil {
			_gdbe := _bfa._cecg._af + _bfa._cecg._ac.Left
			if _fcfd.RelativeToAttr == _bc.ST_PTabRelativeToIndent {
				_gdbe = _bfa._cecg._af
			}
			_ggda := 0.0
			for _, _baf := range _bfa._cecg._cc {
				for _, _bcbe := range _baf._dag {
					for _, _eadf := range _bcbe._bda {
						for _, _ccag := range _eadf._gaa {
							_ggda += _ccag._cdf
						}
					}
				}
			}
			if _fcfd.AlignmentAttr == _bc.ST_PTabAlignmentCenter {
				_gdbe += (_bfa._cecg._dee - (_bfa._cecg._eg + _bfa._cecg._ac.Left + _bfa._cecg._ac.Right)) / 2
			} else if _fcfd.AlignmentAttr == _bc.ST_PTabAlignmentRight {
				_gdbe += _bfa._cecg._eg + _bfa._cecg._ac.Left + _bfa._cecg._ac.Right + _ggda
			}
			_cdbb = _fgef("\u0009", "", false, false, false)
			_dabb := _cdbb[len(_cdbb)-1]
			_dabb._cdf = _gdbe
		} else if _efb.LastRenderedPageBreak != nil && !_bfa._cecg._fcf {
			_cdbb = append(_cdbb, &symbol{_fa: true})
		}
	}
	var _bfaf _ea.TextStyle
	var _ddfc, _acg bool
	var _acc *_ea.Color
	if !_bfcb {
		_bfaf, _ddfc, _acg, _acc = _bfa.makeRunStyle(_dgfg, false, false, false, _bgce, _abc)
		if _bfaf.Font != nil && (_bfa._afff == nil || (_bfa._afff != nil && _bfa._afff.EnableFontSubsetting)) {
			_bfa._fcgg.EnableFontSubsetting(_bfaf.Font)
		}
	}
	for _, _dgdb := range _cdbb {
		if _dgdb._fa && _bfa._fbcd._dbg > _bfa._eaebc.Top {
			_bfa.addCurrentParagraphToCurrentPage()
			_bfa.newPage()
			_bfa.newParagraph()
			_bfa.determineParagraphBounds()
			_bfa.newLine()
			_bfa.newWord()
			continue
		}
		if _dgdb._bea != nil || _dgdb._cce != nil {
			_bfa.addInlineSymbol(_dgdb)
		} else {
			_dgdb._ba = &_bfaf
			_dgdb._ffa = _ddfc
			_dgdb._bcb = _acg
			_dgdb._ad = _acc
			if _dgdb._fbb {
				_cdbd := *_dgfg
				_cdbd.B = nil
				_cdbd.U = nil
				_cfgc, _, _, _ := _bfa.makeRunStyle(&_cdbd, false, false, false, _bgce, _abc)
				_dgdb._ba = &_cfgc
				_dgdb._ad = nil
			}
			_bfa.addTextSymbol(_dgdb)
		}
	}
	if _bfa._eegda != nil && _bfa._eegda._eaee {
		var _dfe, _dfef float64
		for _, _dbc := range _cdbb {
			_dfe += _dbc._cdf
		}
		_eacg := 0
		_ccab := _bfa._fbcd._eeg.Left
		_fdb := len(_bfa._eegda._aceb)
		if _fdb > 1 && _bfa._eegda._eaee {
			_fdb = len(_bfa._eegda._aceb) - 1
		}
		_gda := _bfa._cecg._af < _dfe
		_bbe := _bfa._bbfa._bfd + _dfe
		for {
			var _ddfe float64
			if _gda || _eacg >= _fdb {
				_ddfe = _debg
			} else {
				_ddfe = _bfa._eegda._aceb[_eacg]
				_eacg++
			}
			_ccab += _ddfe
			if _ccab > _bbe {
				_dfef = _ccab - _bbe
				break
			}
		}
		_bfa.addTextSymbol(&symbol{_fge: "\u0020", _cdf: _dfef})
	}
	return false
}

func _befb(_ffgb, _bdfe *_bc.CT_Border, _aggg bool) *_bc.CT_Border {
	if _aggg {
		return _ffgb
	}
	return _bdfe
}

func (_gbge *convertContext) moveCurrentParagraphToNewPage() {
	_gbge.newPage()
	_gde := _gbge._cecg._bgb - _gbge._fbcd._dbg
	_gbge._cecg._bgb -= _gde
	for _, _feaa := range _gbge._cecg._fcd {
		_feaa._gac.Translate(0, -_gde)
	}
	for _, _fbgc := range _gbge._cecg._gb {
		_fbgc._aga -= _gde
	}
	for _, _gfdfc := range _gbge._cecg._bfb {
		_gfdfc._aga -= _gde
	}
	for _, _ebae := range _gbge._cecg._ebd {
		_ebae._bgg -= _gde
	}
	for _, _dgaf := range _gbge._cecg._caf {
		_dgaf._bgg -= _gde
	}
}

func (_cfea *convertContext) addAbsoluteEGPC(_acff []*_bc.EG_PContent, _cee *_bc.CT_PPr) bool {
	_cgf := len(_acff)
	for _, _fgd := range _acff {
		for _, _aag := range _fgd.FldSimple {
			if _aag != nil {
				_cfea.addAbsoluteEGPC(_aag.EG_PContent, _cee)
			}
		}
		if _ffc := _fgd.Hyperlink; _ffc != nil {
			_cfea._bgeb = _ffc
			_cfea.addAbsoluteCRC(_ffc.EG_ContentRunContent, _cee)
		}
		_cfea._bgeb = nil
		if _cfea.addAbsoluteCRC(_fgd.EG_ContentRunContent, _cee) {
			if _cgf > 1 {
				_cfea.moveCurrentParagraphToNewPage()
				continue
			} else {
				return true
			}
		}
		_cgf--
	}
	return false
}

func (_ggcgc *convertContext) makePdfBlockFromCBCs(_bbcf [][]*_bc.EG_ContentBlockContent, _aabf, _ggefg float64, _aaad *_a.Rectangle, _abb bool, _affff *prefix) (*_ea.Block, error) {
	if _aaad == nil {
		_aaad = &_a.Rectangle{}
	}
	_dagg := &_a.Rectangle{Top: _aaad.Top, Bottom: _ggefg - _aaad.Bottom, Left: _aaad.Left, Right: _aabf - _aaad.Right}
	_acbf := _a.MakeTempCreator(_aabf, _ggefg)
	_deeag := &convertContext{_fcgg: _acbf, _cffa: _ggcgc._cffa, _eaebc: _dagg, _eegda: _affff}
	for _, _bbgga := range _bbcf {
		_deeag.addAbsoluteCBCs(_bbgga, nil)
	}
	if _abb {
		_efaf := 0.0
		for _, _cfef := range _deeag._efceb {
			for _, _fefgc := range _cfef._eb {
				_efaf += (_fefgc._egd + _fefgc._ac.Top + _fefgc._ac.Bottom)
			}
		}
		_dagg.Bottom = _efaf - _aaad.Bottom
		_acbf = _a.MakeTempCreator(_aabf, _efaf)
		_deeag = &convertContext{_fcgg: _acbf, _cffa: _ggcgc._cffa, _eaebc: _dagg, _eegda: _affff}
		for _, _efde := range _bbcf {
			_deeag.addAbsoluteCBCs(_efde, nil)
		}
	}
	_deeag.alignSymbolsVertically()
	_deeag._fcgg.NewPage()
	_deeag.drawPage(_deeag._efceb[len(_deeag._efceb)-1])
	return _a.MakeBlockFromCreator(_acbf)
}

type image struct {
	_gbb *_ea.Image
	_ddc float64
	_bgg float64
}

func (_ddf *convertContext) addAnchorBlocks(_bec []*_bc.EG_PContent) {
	for _, _fcg := range _bec {
		for _, _aggd := range _fcg.EG_ContentRunContent {
			if _dcc := _aggd.R; _dcc != nil {
				for _, _bcf := range _dcc.EG_RunInnerContent {
					if _dac := _bcf.Drawing; _dac != nil {
						for _, _aff := range _dac.Anchor {
							var _geda, _eac, _fffc, _gdg float64
							_dbfb, _fde := _aff.PositionH, _aff.PositionV
							if _bae := _dbfb.Choice; _bae != nil {
								if _bae.PosOffset != nil {
									_geda = _eag.FromEMU(int64(*_bae.PosOffset))
								}
							}
							if _eegf := _fde.Choice; _eegf != nil {
								if _eegf.PosOffset != nil {
									_eac = _eag.FromEMU(int64(*_eegf.PosOffset))
								}
							}
							if _deg := _aff.Extent; _deg != nil {
								_gdg = _eag.FromEMU(_deg.CxAttr)
								_fffc = _eag.FromEMU(_deg.CyAttr)
							}
							_fab := _ddf._eaebc.Top + _eac
							_aecb := _fab + _fffc
							_bfcg := _ddf._cecg._eg + _geda
							_cecb := _bfcg + _gdg
							_ddeb := _eac + _fffc
							if _ddeb > _ddf._cecg._cg {
								_ddf._cecg._cg = _ddeb
							}
							if !_aff.AllowOverlapAttr {
								_ddf._cecg._egd = _fffc
							}
							if _aff.Choice != nil && _aff.Choice.WrapNone == nil {
								_ddf._cecg._fcd = append(_ddf._cecg._fcd, &zoneToSkip{_gac: &_a.Rectangle{Top: _fab, Bottom: _aecb, Left: _bfcg, Right: _cecb}, _aba: _aff.Choice})
							}
							if _agdd := _aff.Graphic; _agdd != nil {
								if _cfead := _agdd.GraphicData; _cfead != nil {
									for _, _fabc := range _cfead.Any {
										if _dgae, _dddg := _fabc.(*_eeb.Pic); _dddg {
											_dadf, _agee := _ddf.makePdfImageFromGraphics(_dgae)
											if _agee != nil {
												_da.Log.Debug("C\u0061\u006e\u006e\u006ft \u0072e\u0061\u0064\u0020\u0069\u006da\u0067\u0065\u003a\u0020\u0025\u0073", _agee)
											}
											_bgga := false
											if _dgae.SpPr != nil && _dgae.SpPr.Xfrm != nil {
												if _eagf := _dgae.SpPr.Xfrm.Ext; _eagf != nil {
													_bgga = true
												}
											}
											if _dadf != nil {
												if !_bgga {
													_dadf.Scale(_gdg/_dadf.Width(), _fffc/_dadf.Height())
												} else {
													_dadf.ScaleToWidth(_gdg)
												}
												_agga := &image{_gbb: _dadf, _ddc: _bfcg, _bgg: _fab}
												if _aff.BehindDocAttr {
													_ddf._cecg._caf = append(_ddf._cecg._caf, _agga)
												} else {
													_ddf._cecg._ebd = append(_ddf._cecg._ebd, _agga)
												}
											}
										} else if _baee, _gbbb := _fabc.(*_gg.Chart); _gbbb {
											_edg, _eacc := _ddf.makePdfBlockFromChart(_baee, _gdg, _fffc)
											if _eacc != nil {
												_da.Log.Debug("C\u0061\u006e\u006e\u006ft \u0072e\u0061\u0064\u0020\u0062\u006co\u0063\u006b\u003a\u0020\u0025\u0073", _eacc)
											}
											if _edg != nil {
												_bcd := &block{_egb: _edg, _aef: _bfcg, _aga: _fab}
												if _aff.BehindDocAttr {
													_ddf._cecg._bfb = append(_ddf._cecg._bfb, _bcd)
												} else {
													_ddf._cecg._gb = append(_ddf._cecg._gb, _bcd)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (_acbc *convertContext) makePdfImageFromRelId(_ceec *string) (*_ea.Image, error) {
	if _ceec != nil {
		_ffgae, _ggdf := _acbc._cffa.GetImageObjByRelId(*_ceec)
		if _ggdf != nil {
			return nil, _ggdf
		}
		_aefdd, _ggdf := _f.Open(_ffgae.Path)
		if _ggdf != nil {
			return nil, _ggdf
		}
		_agdgd, _ggdf := _bg.ReadAll(_aefdd)
		if _ggdf != nil {
			return nil, _ggdf
		}
		_aacd, _ggdf := _acbc._fcgg.NewImageFromData(_agdgd)
		if _ggdf != nil {
			return nil, _ggdf
		}
		return _aacd, nil
	}
	return nil, nil
}

type word struct {
	_gaa  []*symbol
	_cfe  float64
	_bfdc float64
	_gfb  bool
}

func _bada(_cada string) string {
	_bcgfa := _bgdc.FindStringSubmatch(_cada)
	if len(_bcgfa) < 2 {
		return ""
	}
	return _bcgfa[1]
}

func _dgfgag(_bbeg *_bc.CT_OnOff) bool {
	if _bbeg != nil {
		if _ebfg := _bbeg.ValAttr; _ebfg != nil {
			if _adbb := _ebfg.Bool; _adbb != nil {
				return *_adbb
			}
			return _ebfg.ST_OnOff1 == _ee.ST_OnOff1On
		}
		return true
	}
	return false
}

func _aecg(_cbea *_bc.EG_RunInnerContent) bool {
	if _abcgg := _cbea.Br; _abcgg != nil {
		return _abcgg.TypeAttr == _bc.ST_BrTypePage
	}
	return false
}

type span struct {
	_bd  float64
	_cf  float64
	_bda []*word
}

var _dcd = map[int32]int32{61623: 8226, 61607: 8226, 61558: 9830, 61656: 8594, 61692: 8730}

func _ffcdc(_dffg float64) float64 { return _dffg * _eag.Millimeter }

func _efdg(_dggc *_bc.CT_ParaRPr, _caec *_bc.CT_RPr) *_bc.CT_ParaRPr {
	if _caec == nil {
		return _dggc
	}
	if _dggc == nil {
		_dggc = _bc.NewCT_ParaRPr()
		if _caec.B != nil {
			_dggc.B = _caec.B
		}
		if _caec.BCs != nil {
			_dggc.BCs = _caec.BCs
		}
		if _caec.I != nil {
			_dggc.I = _caec.I
		}
		if _caec.ICs != nil {
			_dggc.ICs = _caec.ICs
		}
		if _caec.U != nil {
			_dggc.U = _caec.U
		}
		if _caec.Color != nil {
			_dggc.Color = _caec.Color
		}
		return _dggc
	}
	if _dggc.B != _caec.B {
		_dggc.B = _caec.B
	}
	if _dggc.BCs != _caec.BCs {
		_dggc.BCs = _caec.BCs
	}
	if _dggc.I != _caec.I {
		_dggc.I = _caec.I
	}
	if _dggc.ICs != _caec.ICs {
		_dggc.ICs = _caec.ICs
	}
	if _dggc.U != _caec.U {
		_dggc.U = _caec.U
	}
	if _dggc.Color != _caec.Color {
		_dggc.Color = _caec.Color
	}
	return _dggc
}

var _dadd = _ffcdc(2.5)

func (_eea *convertContext) addAbsoluteCBCs(_bfbd []*_bc.EG_ContentBlockContent, _ece []*_bc.EG_ContentBlockContent) {
	_bcad := ""
	_deb := false
	for _, _bff := range _ece {
		if len(_bff.P) < 1 {
			_deb = true
			break
		}
		for _, _feb := range _bff.P {
			if len(_feb.EG_PContent) == 0 {
				break
			}
			if _feb.PPr != nil && _feb.PPr.PStyle != nil {
				_bcad = _feb.PPr.PStyle.ValAttr
				break
			}
		}
	}
	for _, _ffg := range _bfbd {
		for _, _dbe := range _ffg.P {
			_eea.newParagraph()
			if _dbe.PPr != nil && _dbe.PPr.PStyle == nil {
				_cecd := _eea._cffa.Styles.ParagraphStyles()
				for _, _gfc := range _cecd {
					if _fcda := _gfc.X().DefaultAttr; _fcda != nil {
						if _fdf := _fcda.Bool; _fdf != nil && *_fdf {
							_dbe.PPr = _cceb(_dbe.PPr, _gfc.X().PPr, _gfc.X().RPr)
						}
						if _gae := _fcda.ST_OnOff1; _gae == _ee.ST_OnOff1On {
							_dbe.PPr = _cceb(_dbe.PPr, _gfc.X().PPr, _gfc.X().RPr)
						}
						break
					}
				}
			}
			_bfe, _dea := _eea.combinePPrWithStyles(_dbe.PPr)
			if _dea != nil {
				_eea._eegda = _dea
			}
			if _dbe.PPr != nil && _dbe.PPr.PStyle != nil {
				if _dbe.PPr.PStyle.ValAttr != _bcad {
					_dbe.PPr.ContextualSpacing = nil
				}
			}
			if _bfe != nil && _bfe.SectPr != nil {
				for _, _cda := range _bfe.SectPr.EG_HdrFtrReferences {
					if _bad := _cda.HeaderReference; _bad != nil {
						_cab := &headerFooterRef{_eadd: true, _fgcc: _bad.IdAttr, _gdae: _bad.TypeAttr}
						_eea._fbcd._eba = append(_eea._fbcd._eba, _cab)
					}
					if _febf := _cda.FooterReference; _febf != nil {
						_cdcf := &headerFooterRef{_fcfcfa: true, _fgcc: _febf.IdAttr, _gdae: _febf.TypeAttr}
						_eea._fbcd._eba = append(_eea._fbcd._eba, _cdcf)
					}
				}
				if !_deb && (_bfe.SectPr.Type == nil || (_bfe.SectPr.Type != nil && _bfe.SectPr.Type.ValAttr != _bc.ST_SectionMarkContinuous)) && _dea == nil && !_dgfgag(_bfe.WidowControl) {
					_eea.newPage()
					continue
				}
				if len(_dbe.EG_PContent) < 1 {
					continue
				}
			}
			_eea.assignPropsToAbsoluteParagraph(_bfe, _eea._cecg)
			_eea.determineParagraphBounds()
			_eea.newLine()
			_eea.newWord()
			_aa := _dbe.EG_PContent
			if len(_aa) == 0 {
				_eea.addEmptyLine()
			} else {
				if _eea.addAbsoluteEGPC(_aa, _bfe) {
					_eea.newPage()
					continue
				}
				if _eea.currentParagraphOverflowsCurrentPage() {
					_eea.moveCurrentParagraphToNewPage()
				}
				_eea.addAnchorBlocks(_aa)
				_eea.addAnchorExtra(_aa)
				_eea.addCurrentWordToParagraph()
			}
			_eea.addCurrentParagraphToCurrentPage()
		}
		for _, _fbd := range _ffg.Tbl {
			if _eea._cecg == nil {
				_eea.newParagraph()
				_eea.determineParagraphBounds()
				_eea.newLine()
				_eea.newWord()
			}
			_eea.addAbsoluteTable(_fbd)
		}
	}
}

func (_egge *convertContext) addParagraphWithTableToHeaderFooter(_gdcg _ea.Table, _fafd, _ebbb float64) {
	_egge.newParagraph()
	_egge._cecg._ac = &_a.Rectangle{Top: _ffcdc(2), Bottom: _ffcdc(2), Left: 0, Right: 0}
	_egge._cecg._fdg = &tableWrapper{_cca: &_gdcg, _fb: _fafd}
	_egge._cecg._af = _ebbb
	_egge._cecg._egd = _gdcg.Height()
	_egge.determineParagraphBounds()
	if _egge._cbae {
		_egge.addCurrentParagraphHeaderToCurrentPage()
	} else if _egge._bcdd {
		_egge.addCurrentParagraphFooterToCurrentPage()
	}
}

func (_fbf *convertContext) adjustRightBoundOfLastSpan() {
	_fgce := _fbf._ecbg._cf
	_cafc := _fbf._bbfa._ag + _fbf._cecg._bgb
	_gdab := _cafc + _fbf._bbfa._bca
	for _, _afae := range _fbf._fbcd._gf {
		if ((_cafc > _afae._gac.Top && _cafc < _afae._gac.Bottom) || (_gdab > _afae._gac.Top && _cafc < _afae._gac.Bottom)) && (_fgce > _afae._gac.Left) {
			_fgce = _afae._gac.Left
		}
	}
	_fbf._ecbg._cf = _fgce
}

var _gefg = map[string]uint16{"\u0061\u0061": 0x1404, "\u0061\u0061\u002dD\u004a": 0x1000, "\u0061\u0061\u002dE\u0052": 0x1000, "\u0061\u0061\u002dE\u0054": 0x1000, "\u0061\u0066": 0x0036, "\u0061\u0066\u002dN\u0041": 0x1000, "\u0061\u0066\u002dZ\u0041": 0x0436, "\u0061\u0067\u0071": 0x1000, "\u0061\u0067\u0071\u002d\u0043\u004d": 0x1000, "\u0061\u006b": 0x1000, "\u0061\u006b\u002dG\u0048": 0x1000, "\u0073\u0071": 0x001C, "\u0073\u0071\u002dA\u004c": 0x041C, "\u0073\u0071\u002dM\u004b": 0x1000, "\u0067\u0073\u0077": 0x0084, "\u0067\u0073\u0077\u002d\u0046\u0052": 0x0484, "\u0067\u0073\u0077\u002d\u004c\u0049": 0x1000, "\u0067\u0073\u0077\u002d\u0043\u0048": 0x1000, "\u0061\u006d": 0x005E, "\u0061\u006d\u002dE\u0054": 0x045E, "\u0061\u0072": 0x0001, "\u0061\u0072\u002dD\u005a": 0x1401, "\u0061\u0072\u002dT\u0044": 0x1000, "\u0061\u0072\u002dK\u004d": 0x1000, "\u0061\u0072\u002dD\u004a": 0x1000, "\u0061\u0072\u002dE\u0047": 0x0c01, "\u0061\u0072\u002dE\u0052": 0x1000, "\u0061\u0072\u002dI\u0051": 0x0801, "\u0061\u0072\u002dI\u004c": 0x1000, "\u0061\u0072\u002dJ\u004f": 0x2C01, "\u0061\u0072\u002dK\u0057": 0x3401, "\u0061\u0072\u002dL\u0042": 0x3001, "\u0061\u0072\u002dL\u0059": 0x1001, "\u0061\u0072\u002dM\u0052": 0x1000, "\u0061\u0072\u002dM\u0041": 0x1801, "\u0061\u0072\u002dO\u004d": 0x2001, "\u0061\u0072\u002dP\u0053": 0x1000, "\u0061\u0072\u002dQ\u0041": 0x4001, "\u0061\u0072\u002dS\u0041": 0x0401, "\u0061\u0072\u002dS\u004f": 0x1000, "\u0061\u0072\u002dS\u0053": 0x1000, "\u0061\u0072\u002dS\u0044": 0x1000, "\u0061\u0072\u002dS\u0059": 0x2801, "\u0061\u0072\u002dT\u004e": 0x1C01, "\u0061\u0072\u002dA\u0045": 0x3801, "\u0061\u0072\u002d\u0030\u0030\u0031": 0x1000, "\u0061\u0072\u002dY\u0045": 0x2401, "\u0068\u0079": 0x002B, "\u0068\u0079\u002dA\u004d": 0x042B, "\u0061\u0073": 0x004D, "\u0061\u0073\u002dI\u004e": 0x044D, "\u0061\u0073\u0074": 0x1000, "\u0061\u0073\u0074\u002d\u0045\u0053": 0x1000, "\u0061\u0073\u0061": 0x1000, "\u0061\u0073\u0061\u002d\u0054\u005a": 0x1000, "\u0061z\u002d\u0043\u0079\u0072\u006c": 0x742C, "\u0061\u007a\u002d\u0043\u0079\u0072\u006c\u002d\u0041\u005a": 0x082C, "\u0061\u007a": 0x002C, "\u0061z\u002d\u004c\u0061\u0074\u006e": 0x782C, "\u0061\u007a\u002d\u004c\u0061\u0074\u006e\u002d\u0041\u005a": 0x042C, "\u006b\u0073\u0066": 0x1000, "\u006b\u0073\u0066\u002d\u0043\u004d": 0x1000, "\u0062\u006d": 0x1000, "\u0062\u006d\u002d\u004c\u0061\u0074\u006e\u002d\u004d\u004c": 0x1000, "\u0062\u006e": 0x0045, "\u0062\u006e\u002dB\u0044": 0x0845, "\u0062\u006e\u002dI\u004e": 0x0445, "\u0062\u0061\u0073": 0x1000, "\u0062\u0061\u0073\u002d\u0043\u004d": 0x1000, "\u0062\u0061": 0x006D, "\u0062\u0061\u002dR\u0055": 0x046D, "\u0065\u0075": 0x002D, "\u0065\u0075\u002dE\u0053": 0x042D, "\u0062\u0065": 0x0023, "\u0062\u0065\u002dB\u0059": 0x0423, "\u0062\u0065\u006d": 0x1000, "\u0062\u0065\u006d\u002d\u005a\u004d": 0x1000, "\u0062\u0065\u007a": 0x1000, "\u0062\u0065\u007a\u002d\u0054\u005a": 0x1000, "\u0062\u0079\u006e": 0x1000, "\u0062\u0079\u006e\u002d\u0045\u0052": 0x1000, "\u0062\u0072\u0078": 0x1000, "\u0062\u0072\u0078\u002d\u0049\u004e": 0x1000, "\u0062s\u002d\u0043\u0072\u0079\u006c": 0x6414, "\u0062\u0073\u002d\u0043\u0079\u0072\u006c\u002d\u0042\u0041": 0x201A, "\u0062s\u002d\u004c\u0061\u0074\u006e": 0x681A, "\u0062\u0073": 0x7814, "\u0062\u0073\u002d\u004c\u0061\u0074\u006e\u002d\u0042\u0041": 0x141A, "\u0062\u0072": 0x007E, "\u0062\u0072\u002dF\u0052": 0x047E, "\u0062\u0067": 0x0002, "\u0062\u0067\u002dB\u0047": 0x0402, "\u006d\u0079": 0x0055, "\u006d\u0079\u002dM\u004d": 0x0455, "\u0063\u0061": 0x0003, "\u0063\u0061\u002dA\u0044": 0x1000, "\u0063\u0061\u002dF\u0052": 0x1000, "\u0063\u0061\u002dI\u0054": 0x1000, "\u0063\u0061\u002dE\u0053": 0x0403, "\u0063\u0065\u0062": 0x1000, "\u0063\u0065\u0062\u002d\u004c\u0061\u0074\u006e": 0x1000, "c\u0065\u0062\u002d\u004c\u0061\u0074\u006e\u002d\u0050\u0048": 0x1000, "t\u007a\u006d\u002d\u0041\u0072\u0061\u0062\u002d\u004d\u0041": 0x045F, "t\u006d\u007a\u002d\u004c\u0061\u0074\u006e\u002d\u004d\u0041": 0x1000, "\u006b\u0075": 0x0092, "\u006bu\u002d\u0041\u0072\u0061\u0062": 0x7c92, "\u006b\u0075\u002d\u0041\u0072\u0061\u0062\u002d\u0049\u0051": 0x0492, "\u0063\u0063\u0070": 0x1000, "\u0063\u0063\u0070\u002d\u0043\u0061\u006b\u006d": 0x1000, "c\u0063\u0070\u002d\u0043\u0061\u006b\u006d\u002d\u0049\u004e": 0x1000, "\u0063\u0065\u002dR\u0055": 0x1000, "\u0063\u0068\u0072": 0x005C, "\u0063\u0068\u0072\u002d\u0043\u0068\u0065\u0072": 0x7c5c, "c\u0068\u0072\u002d\u0043\u0068\u0065\u0072\u002d\u0055\u0053": 0x045C, "\u0063\u0067\u0067": 0x1000, "\u0063\u0067\u0067\u002d\u0055\u0047": 0x1000, "\u007ah\u002d\u0048\u0061\u006e\u0073": 0x0004, "\u007a\u0068": 0x7804, "\u007a\u0068\u002dC\u004e": 0x0804, "\u007a\u0068\u002dS\u0047": 0x1004, "\u007ah\u002d\u0048\u0061\u006e\u0074": 0x7C04, "\u007a\u0068\u002dH\u004b": 0x0C04, "\u007a\u0068\u002dM\u004f": 0x1404, "\u007a\u0068\u002dT\u0057": 0x0404, "\u0063\u0075\u002dR\u0055": 0x1000, "\u0073\u0077\u0063": 0x1000, "\u0073\u0077\u0063\u002d\u0043\u0044": 0x1000, "\u006b\u0077": 0x1000, "\u006b\u0077\u002dG\u0042": 0x1000, "\u0063\u006f": 0x0083, "\u0063\u006f\u002dF\u0052": 0x0483, "\u0068\u0072": 0x001A, "\u0068\u0072\u002dH\u0052": 0x041A, "\u0068\u0072\u002dB\u0041": 0x101A, "\u0063\u0073": 0x0005, "\u0063\u0073\u002dC\u005a": 0x0405, "\u0064\u0061": 0x0006, "\u0064\u0061\u002dD\u004b": 0x0406, "\u0064\u0061\u002dG\u004c": 0x1000, "\u0070\u0072\u0073": 0x008C, "\u0070\u0072\u0073\u002d\u0041\u0046": 0x048C, "\u0064\u0076": 0x0065, "\u0064\u0076\u002dM\u0056": 0x0465, "\u0064\u0075\u0061": 0x1000, "\u0064\u0075\u0061\u002d\u0043\u004d": 0x1000, "\u006e\u006c": 0x0013, "\u006e\u006c\u002dA\u0057": 0x1000, "\u006e\u006c\u002dB\u0045": 0x0813, "\u006e\u006c\u002dB\u0051": 0x1000, "\u006e\u006c\u002dC\u0057": 0x1000, "\u006e\u006c\u002dN\u004c": 0x0413, "\u006e\u006c\u002dS\u0058": 0x1000, "\u006e\u006c\u002dS\u0052": 0x1000, "\u0064\u007a": 0x1000, "\u0064\u007a\u002dB\u0054": 0x0C51, "\u0065\u0062\u0075": 0x1000, "\u0065\u0062\u0075\u002d\u004b\u0045": 0x1000, "\u0065\u006e\u002dA\u0053": 0x1000, "\u0065\u006e\u002dA\u0049": 0x1000, "\u0065\u006e\u002dA\u0047": 0x1000, "\u0065\u006e\u002dA\u0055": 0x0C09, "\u0065\u006e\u002dA\u0054": 0x1000, "\u0065\u006e\u002dB\u0053": 0x1000, "\u0065\u006e\u002dB\u0042": 0x1000, "\u0065\u006e\u002dB\u0045": 0x1000, "\u0065\u006e\u002dB\u005a": 0x2809, "\u0065\u006e\u002dB\u004d": 0x1000, "\u0065\u006e\u002dB\u0057": 0x1000, "\u0065\u006e\u002dI\u004f": 0x1000, "\u0065\u006e\u002dV\u0047": 0x1000, "\u0065\u006e\u002dB\u0049": 0x1000, "\u0065\u006e\u002dC\u004d": 0x1000, "\u0065\u006e\u002dC\u0041": 0x1009, "\u0065\u006e\u002d\u0030\u0032\u0039": 0x2409, "\u0065\u006e\u002dK\u0059": 0x1000, "\u0065\u006e\u002dC\u0058": 0x1000, "\u0065\u006e\u002dC\u0043": 0x1000, "\u0065\u006e\u002dC\u004b": 0x1000, "\u0065\u006e\u002dC\u0059": 0x1000, "\u0065\u006e\u002dD\u004b": 0x1000, "\u0065\u006e\u002dD\u004d": 0x1000, "\u0065\u006e\u002dE\u0052": 0x1000, "\u0065\u006e\u002d\u0031\u0035\u0030": 0x1000, "\u0065\u006e\u002dF\u004b": 0x1000, "\u0065\u006e\u002dF\u0049": 0x1000, "\u0065\u006e\u002dF\u004a": 0x1000, "\u0065\u006e\u002dG\u004d": 0x1000, "\u0065\u006e\u002dD\u0045": 0x1000, "\u0065\u006e\u002dG\u0048": 0x1000, "\u0065\u006e\u002dG\u0049": 0x1000, "\u0065\u006e\u002dG\u0044": 0x1000, "\u0065\u006e\u002dG\u0055": 0x1000, "\u0065\u006e\u002dG\u0047": 0x1000, "\u0065\u006e\u002dG\u0059": 0x1000, "\u0065\u006e\u002dH\u004b": 0x3C09, "\u0065\u006e\u002dI\u004e": 0x4009, "\u0065\u006e\u002dI\u004d": 0x1000, "\u0065\u006e\u002dI\u004c": 0x1000, "\u0065\u006e\u002dJ\u004d": 0x2009, "\u0065\u006e\u002dJ\u0045": 0x1000, "\u0065\u006e\u002dK\u0045": 0x1000, "\u0065\u006e\u002dK\u0049": 0x1000, "\u0065\u006e\u002dL\u0053": 0x1000, "\u0065\u006e\u002dL\u0052": 0x1000, "\u0065\u006e\u002dM\u004f": 0x1000, "\u0065\u006e\u002dM\u0047": 0x1000, "\u0065\u006e\u002dM\u0057": 0x1000, "\u0065\u006e\u002dM\u0059": 0x4409, "\u0065\u006e\u002dM\u0054": 0x1000, "\u0065\u006e\u002dM\u0048": 0x1000, "\u0065\u006e\u002dM\u0055": 0x1000, "\u0065\u006e\u002dF\u004d": 0x1000, "\u0065\u006e\u002dM\u0053": 0x1000, "\u0065\u006e\u002dN\u0041": 0x1000, "\u0065\u006e\u002dN\u0052": 0x1000, "\u0065\u006e\u002dN\u004c": 0x1000, "\u0065\u006e\u002dN\u005a": 0x1409, "\u0065\u006e\u002dN\u0047": 0x1000, "\u0065\u006e\u002dN\u0055": 0x1000, "\u0065\u006e\u002dN\u0046": 0x1000, "\u0065\u006e\u002dM\u0050": 0x1000, "\u0065\u006e\u002dP\u004b": 0x1000, "\u0065\u006e\u002dP\u0057": 0x1000, "\u0065\u006e\u002dP\u0047": 0x1000, "\u0065\u006e\u002dP\u004e": 0x1000, "\u0065\u006e\u002dP\u0052": 0x1000, "\u0065\u006e\u002dP\u0048": 0x3409, "\u0065\u006e\u002dR\u0057": 0x1000, "\u0065\u006e\u002dK\u004e": 0x1000, "\u0065\u006e\u002dL\u0043": 0x1000, "\u0065\u006e\u002dV\u0043": 0x1000, "\u0065\u006e\u002dW\u0053": 0x1000, "\u0065\u006e\u002dS\u0043": 0x1000, "\u0065\u006e\u002dS\u004c": 0x1000, "\u0065\u006e\u002dS\u0047": 0x4809, "\u0065\u006e\u002dS\u0058": 0x1000, "\u0065\u006e\u002dS\u0049": 0x1000, "\u0065\u006e\u002dS\u0042": 0x1000, "\u0065\u006e\u002dZ\u0041": 0x1C09, "\u0065\u006e\u002dS\u0053": 0x1000, "\u0065\u006e\u002dS\u0048": 0x1000, "\u0065\u006e\u002dS\u0044": 0x1000, "\u0065\u006e\u002dS\u005a": 0x1000, "\u0065\u006e\u002dS\u0045": 0x1000, "\u0065\u006e\u002dC\u0048": 0x1000, "\u0065\u006e\u002dT\u005a": 0x1000, "\u0065\u006e\u002dT\u004b": 0x1000, "\u0065\u006e\u002dT\u004f": 0x1000, "\u0065\u006e\u002dT\u0054": 0x2c09, "\u0065\u006e\u002dT\u0043": 0x1000, "\u0065\u006e\u002dT\u0056": 0x1000, "\u0065\u006e\u002dU\u0047": 0x1000, "\u0065\u006e\u002dA\u0045": 0x4C09, "\u0065\u006e\u002dG\u0042": 0x0809, "\u0065\u006e\u002dU\u0053": 0x0409, "\u0065\u006e\u002dU\u004d": 0x1000, "\u0065\u006e\u002dV\u0049": 0x1000, "\u0065\u006e\u002dV\u0055": 0x1000, "\u0065\u006e\u002d\u0030\u0030\u0031": 0x1000, "\u0065\u006e\u002dZ\u004d": 0x1000, "\u0065\u006e\u002dZ\u0057": 0x3009, "\u0065\u006f": 0x1000, "\u0065\u006f\u002d\u0030\u0030\u0031": 0x1000, "\u0065\u0074": 0x0025, "\u0065\u0074\u002dE\u0045": 0x0425, "\u0065\u0065": 0x1000, "\u0065\u0065\u002dG\u0048": 0x1000, "\u0065\u0065\u002dT\u0047": 0x1000, "\u0065\u0077\u006f": 0x1000, "\u0065\u0077\u006f\u002d\u0043\u004d": 0x1000, "\u0066\u006f": 0x0038, "\u0066\u006f\u002dD\u004b": 0x1000, "\u0066\u006f\u002dF\u004f": 0x0438, "\u0066\u0069\u006c": 0x0064, "\u0066\u0069\u006c\u002d\u0050\u0048": 0x0464, "\u0066\u0069": 0x000B, "\u0066\u0069\u002dF\u0049": 0x040B, "\u0066\u0072": 0x000C, "\u0066\u0072\u002dD\u005a": 0x1000, "\u0066\u0072\u002dB\u0045": 0x080C, "\u0066\u0072\u002dB\u004a": 0x1000, "\u0066\u0072\u002dB\u0046": 0x1000, "\u0066\u0072\u002dB\u0049": 0x1000, "\u0066\u0072\u002dC\u004d": 0x2c0C, "\u0066\u0072\u002dC\u0041": 0x0c0C, "\u0066\u0072\u002d\u0030\u0032\u0039": 0x1C0C, "\u0066\u0072\u002dC\u0046": 0x1000, "\u0066\u0072\u002dT\u0044": 0x1000, "\u0066\u0072\u002dK\u004d": 0x1000, "\u0066\u0072\u002dC\u0047": 0x1000, "\u0066\u0072\u002dC\u0044": 0x240C, "\u0066\u0072\u002dC\u0049": 0x300C, "\u0066\u0072\u002dD\u004a": 0x1000, "\u0066\u0072\u002dG\u0051": 0x1000, "\u0066\u0072\u002dF\u0052": 0x040C, "\u0066\u0072\u002dG\u0046": 0x1000, "\u0066\u0072\u002dP\u0046": 0x1000, "\u0066\u0072\u002dG\u0041": 0x1000, "\u0066\u0072\u002dG\u0050": 0x1000, "\u0066\u0072\u002dG\u004e": 0x1000, "\u0066\u0072\u002dH\u0054": 0x3c0C, "\u0066\u0072\u002dL\u0055": 0x140C, "\u0066\u0072\u002dM\u0047": 0x1000, "\u0066\u0072\u002dM\u004c": 0x340C, "\u0066\u0072\u002dM\u0051": 0x1000, "\u0066\u0072\u002dM\u0052": 0x1000, "\u0066\u0072\u002dM\u0055": 0x1000, "\u0066\u0072\u002dY\u0054": 0x1000, "\u0066\u0072\u002dM\u0041": 0x380C, "\u0066\u0072\u002dN\u0043": 0x1000, "\u0066\u0072\u002dN\u0045": 0x1000, "\u0066\u0072\u002dM\u0043": 0x180C, "\u0066\u0072\u002dR\u0045": 0x200C, "\u0066\u0072\u002dR\u0057": 0x1000, "\u0066\u0072\u002dB\u004c": 0x1000, "\u0066\u0072\u002dM\u0046": 0x1000, "\u0066\u0072\u002dP\u004d": 0x1000, "\u0066\u0072\u002dS\u004e": 0x280C, "\u0066\u0072\u002dS\u0043": 0x1000, "\u0066\u0072\u002dC\u0048": 0x100C, "\u0066\u0072\u002dS\u0059": 0x1000, "\u0066\u0072\u002dT\u0047": 0x1000, "\u0066\u0072\u002dT\u004e": 0x1000, "\u0066\u0072\u002dV\u0055": 0x1000, "\u0066\u0072\u002dW\u0046": 0x1000, "\u0066\u0079": 0x0062, "\u0066\u0079\u002dN\u004c": 0x0462, "\u0066\u0075\u0072": 0x1000, "\u0066\u0075\u0072\u002d\u0049\u0054": 0x1000, "\u0066\u0066": 0x0067, "\u0066f\u002d\u004c\u0061\u0074\u006e": 0x7C67, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u0042\u0046": 0x1000, "\u0066\u0066\u002dC\u004d": 0x1000, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u0043\u004d": 0x1000, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u0047\u004d": 0x1000, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u0047\u0048": 0x1000, "\u0066\u0066\u002dG\u004e": 0x1000, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u0047\u004e": 0x1000, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u0047\u0057": 0x1000, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u004c\u0052": 0x1000, "\u0066\u0066\u002dM\u0052": 0x1000, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u004d\u0052": 0x1000, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u004e\u0045": 0x1000, "\u0066\u0066\u002dN\u0047": 0x0467, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u004e\u0047": 0x0467, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u0053\u004e": 0x0867, "\u0066\u0066\u002d\u004c\u0061\u0074\u006e\u002d\u0053\u004c": 0x1000, "\u0067\u006c": 0x0056, "\u0067\u006c\u002dE\u0053": 0x0456, "\u006c\u0067": 0x1000, "\u006c\u0067\u002dU\u0047": 0x1000, "\u006b\u0061": 0x0037, "\u006b\u0061\u002dG\u0045": 0x0437, "\u0064\u0065": 0x0007, "\u0064\u0065\u002dA\u0054": 0x0C07, "\u0064\u0065\u002dB\u0045": 0x1000, "\u0064\u0065\u002dD\u0045": 0x0407, "\u0064\u0065\u002dI\u0054": 0x1000, "\u0064\u0065\u002dL\u0049": 0x1407, "\u0064\u0065\u002dL\u0055": 0x1007, "\u0064\u0065\u002dC\u0048": 0x0807, "\u0065\u006c": 0x0008, "\u0065\u006c\u002dC\u0059": 0x1000, "\u0065\u006c\u002dG\u0052": 0x0408, "\u006b\u006c": 0x006F, "\u006b\u006c\u002dG\u004c": 0x046F, "\u0067\u006e": 0x0074, "\u0067\u006e\u002dP\u0059": 0x0474, "\u0067\u0075": 0x0047, "\u0067\u0075\u002dI\u004e": 0x0447, "\u0067\u0075\u007a": 0x1000, "\u0067\u0075\u007a\u002d\u004b\u0045": 0x1000, "\u0068\u0061": 0x0068, "\u0068a\u002d\u004c\u0061\u0074\u006e": 0x7C68, "\u0068\u0061\u002d\u004c\u0061\u0074\u006e\u002d\u0047\u0048": 0x1000, "\u0068\u0061\u002d\u004c\u0061\u0074\u006e\u002d\u004e\u0045": 0x1000, "\u0068\u0061\u002d\u004c\u0061\u0074\u006e\u002d\u004e\u0047": 0x0468, "\u0068\u0061\u0077": 0x0075, "\u0068\u0061\u0077\u002d\u0055\u0053": 0x0475, "\u0068\u0065": 0x000D, "\u0068\u0065\u002dI\u004c": 0x040D, "\u0068\u0069": 0x0039, "\u0068\u0069\u002dI\u004e": 0x0439, "\u0068\u0075": 0x000E, "\u0068\u0075\u002dH\u0055": 0x040E, "\u0069\u0073": 0x000F, "\u0069\u0073\u002dI\u0053": 0x040F, "\u0069\u0067": 0x0070, "\u0069\u0067\u002dN\u0047": 0x0470, "\u0069\u0064": 0x0021, "\u0069\u0064\u002dI\u0044": 0x0421, "\u0069\u0061": 0x1000, "\u0069\u0061\u002dF\u0052": 0x1000, "\u0069\u0061\u002d\u0030\u0030\u0031": 0x1000, "\u0069\u0075": 0x005D, "\u0069u\u002d\u004c\u0061\u0074\u006e": 0x7C5D, "\u0069\u0075\u002d\u004c\u0061\u0074\u006e\u002d\u0043\u0041": 0x085D, "\u0069u\u002d\u0043\u0061\u006e\u0073": 0x785D, "\u0067\u0061": 0x003C, "\u0067\u0061\u002dI\u0045": 0x083C, "\u0069\u0074": 0x0010, "\u0069\u0074\u002dI\u0054": 0x0410, "\u0069\u0074\u002dS\u004d": 0x1000, "\u0069\u0074\u002dC\u0048": 0x0810, "\u0069\u0074\u002dV\u0041": 0x1000, "\u006a\u0061": 0x0011, "\u006a\u0061\u002dJ\u0050": 0x0411, "\u006a\u0076": 0x1000, "\u006av\u002d\u004c\u0061\u0074\u006e": 0x1000, "\u006a\u0076\u002d\u004c\u0061\u0074\u006e\u002d\u0049\u0044": 0x1000, "\u0064\u0079\u006f": 0x1000, "\u0064\u0079\u006f\u002d\u0053\u004e": 0x1000, "\u006b\u0065\u0061": 0x1000, "\u006b\u0065\u0061\u002d\u0043\u0056": 0x1000, "\u006b\u0061\u0062": 0x1000, "\u006b\u0061\u0062\u002d\u0044\u005a": 0x1000, "\u006b\u006b\u006a": 0x1000, "\u006b\u006b\u006a\u002d\u0043\u004d": 0x1000, "\u006b\u006c\u006e": 0x1000, "\u006b\u006c\u006e\u002d\u004b\u0045": 0x1000, "\u006b\u0061\u006d": 0x1000, "\u006b\u0061\u006d\u002d\u004b\u0045": 0x1000, "\u006b\u006e": 0x004B, "\u006b\u006e\u002dI\u004e": 0x044B, "\u006b\u0072\u002d\u004c\u0061\u0074\u006e\u002d\u004e\u0047": 0x0471, "\u006b\u0073": 0x0060, "\u006bs\u002d\u0041\u0072\u0061\u0062": 0x0460, "\u006b\u0073\u002d\u0041\u0072\u0061\u0062\u002d\u0049\u004e": 0x1000, "\u006b\u0073\u002d\u0044\u0065\u0076\u0061\u002d\u0049\u004e": 0x0860, "\u006b\u006b": 0x003F, "\u006b\u006b\u002dK\u005a": 0x043F, "\u006b\u006d": 0x0053, "\u006b\u006d\u002dK\u0048": 0x0453, "\u0071\u0075\u0063": 0x0086, "q\u0075\u0063\u002d\u004c\u0061\u0074\u006e\u002d\u0047\u0054": 0x0486, "\u006b\u0069": 0x1000, "\u006b\u0069\u002dK\u0045": 0x1000, "\u0072\u0077": 0x0087, "\u0072\u0077\u002dR\u0057": 0x0487, "\u0073\u0077\u002dK\u0045": 0x0441, "\u0073\u0077\u002dT\u005a": 0x1000, "\u0073\u0077\u002dU\u0047": 0x1000, "\u006b\u006f\u006b": 0x0057, "\u006b\u006f\u006b\u002d\u0049\u004e": 0x0457, "\u006b\u006f": 0x0012, "\u006b\u006f\u002dK\u0052": 0x0412, "\u006b\u006f\u002dK\u0050": 0x1000, "\u006b\u0068\u0071": 0x1000, "\u006b\u0068\u0071\u002d\u004d\u004c": 0x1000, "\u0073\u0065\u0073": 0x1000, "\u0073\u0065\u0073\u002d\u004d\u004c": 0x1000, "\u006e\u006d\u0067": 0x1000, "\u006e\u006d\u0067\u002d\u0043\u004d": 0x1000, "\u006b\u0079": 0x0040, "\u006b\u0079\u002dK\u0047": 0x0440, "\u006b\u0075\u002d\u0041\u0072\u0061\u0062\u002d\u0049\u0052": 0x1000, "\u006c\u006b\u0074": 0x1000, "\u006c\u006b\u0074\u002d\u0055\u0053": 0x1000, "\u006c\u0061\u0067": 0x1000, "\u006c\u0061\u0067\u002d\u0054\u005a": 0x1000, "\u006c\u006f": 0x0054, "\u006c\u006f\u002dL\u0041": 0x0454, "\u006c\u0061\u002dV\u0041": 0x0476, "\u006c\u0076": 0x0026, "\u006c\u0076\u002dL\u0056": 0x0426, "\u006c\u006e": 0x1000, "\u006c\u006e\u002dA\u004f": 0x1000, "\u006c\u006e\u002dC\u0046": 0x1000, "\u006c\u006e\u002dC\u0047": 0x1000, "\u006c\u006e\u002dC\u0044": 0x1000, "\u006c\u0074": 0x0027, "\u006c\u0074\u002dL\u0054": 0x0427, "\u006e\u0064\u0073": 0x1000, "\u006e\u0064\u0073\u002d\u0044\u0045": 0x1000, "\u006e\u0064\u0073\u002d\u004e\u004c": 0x1000, "\u0064\u0073\u0062": 0x7C2E, "\u0064\u0073\u0062\u002d\u0044\u0045": 0x082E, "\u006c\u0075": 0x1000, "\u006c\u0075\u002dC\u0044": 0x1000, "\u006c\u0075\u006f": 0x1000, "\u006c\u0075\u006f\u002d\u004b\u0045": 0x1000, "\u006c\u0062": 0x006E, "\u006c\u0062\u002dL\u0055": 0x046E, "\u006c\u0075\u0079": 0x1000, "\u006c\u0075\u0079\u002d\u004b\u0045": 0x1000, "\u006d\u006b": 0x002F, "\u006d\u006b\u002dM\u004b": 0x042F, "\u006a\u006d\u0063": 0x1000, "\u006a\u006d\u0063\u002d\u0054\u005a": 0x1000, "\u006d\u0067\u0068": 0x1000, "\u006d\u0067\u0068\u002d\u004d\u005a": 0x1000, "\u006b\u0064\u0065": 0x1000, "\u006b\u0064\u0065\u002d\u0054\u005a": 0x1000, "\u006d\u0067": 0x1000, "\u006d\u0067\u002dM\u0047": 0x1000, "\u006d\u0073": 0x003E, "\u006d\u0073\u002dB\u004e": 0x083E, "\u006d\u0073\u002dM\u0059": 0x043E, "\u006d\u006c": 0x004C, "\u006d\u006c\u002dI\u004e": 0x044C, "\u006d\u0074": 0x003A, "\u006d\u0074\u002dM\u0054": 0x043A, "\u0067\u0076": 0x1000, "\u0067\u0076\u002dI\u004d": 0x1000, "\u006d\u0069": 0x0081, "\u006d\u0069\u002dN\u005a": 0x0481, "\u0061\u0072\u006e": 0x007A, "\u0061\u0072\u006e\u002d\u0043\u004c": 0x047A, "\u006d\u0072": 0x004E, "\u006d\u0072\u002dI\u004e": 0x044E, "\u006d\u0061\u0073": 0x1000, "\u006d\u0061\u0073\u002d\u004b\u0045": 0x1000, "\u006d\u0061\u0073\u002d\u0054\u005a": 0x1000, "\u006d\u007a\u006e\u002d\u0049\u0052": 0x1000, "\u006d\u0065\u0072": 0x1000, "\u006d\u0065\u0072\u002d\u004b\u0045": 0x1000, "\u006d\u0067\u006f": 0x1000, "\u006d\u0067\u006f\u002d\u0043\u004d": 0x1000, "\u006d\u006f\u0068": 0x007C, "\u006d\u006f\u0068\u002d\u0043\u0041": 0x047C, "\u006d\u006e": 0x0050, "\u006dn\u002d\u0043\u0079\u0072\u006c": 0x7850, "\u006d\u006e\u002dM\u004e": 0x0450, "\u006dn\u002d\u004d\u006f\u006e\u0067": 0x7C50, "\u006d\u006e\u002d\u004d\u006f\u006e\u0067\u002d\u0043\u004e": 0x0850, "\u006d\u006e\u002d\u004d\u006f\u006e\u0067\u002d\u004d\u004e": 0x0C50, "\u006d\u0066\u0065": 0x1000, "\u006d\u0066\u0065\u002d\u004d\u0055": 0x1000, "\u006d\u0075\u0061": 0x1000, "\u006d\u0075\u0061\u002d\u0043\u004d": 0x1000, "\u006e\u0071\u006f": 0x1000, "\u006e\u0071\u006f\u002d\u0047\u004e": 0x1000, "\u006e\u0061\u0071": 0x1000, "\u006e\u0061\u0071\u002d\u004e\u0041": 0x1000, "\u006e\u0065": 0x0061, "\u006e\u0065\u002dI\u004e": 0x0861, "\u006e\u0065\u002dN\u0050": 0x0461, "\u006e\u006e\u0068": 0x1000, "\u006e\u006e\u0068\u002d\u0043\u004d": 0x1000, "\u006a\u0067\u006f": 0x1000, "\u006a\u0067\u006f\u002d\u0043\u004d": 0x1000, "\u006c\u0072\u0063\u002d\u0049\u0051": 0x1000, "\u006c\u0072\u0063\u002d\u0049\u0052": 0x1000, "\u006e\u0064": 0x1000, "\u006e\u0064\u002dZ\u0057": 0x1000, "\u006e\u006f": 0x0014, "\u006e\u0062": 0x7C14, "\u006e\u0062\u002dN\u004f": 0x0414, "\u006e\u006e": 0x7814, "\u006e\u006e\u002dN\u004f": 0x0814, "\u006e\u0062\u002dS\u004a": 0x1000, "\u006e\u0075\u0073": 0x1000, "\u006e\u0075\u0073\u002d\u0053\u0044": 0x1000, "\u006e\u0075\u0073\u002d\u0053\u0053": 0x1000, "\u006e\u0079\u006e": 0x1000, "\u006e\u0079\u006e\u002d\u0055\u0047": 0x1000, "\u006f\u0063": 0x0082, "\u006f\u0063\u002dF\u0052": 0x0482, "\u006f\u0072": 0x0048, "\u006f\u0072\u002dI\u004e": 0x0448, "\u006f\u006d": 0x0072, "\u006f\u006d\u002dE\u0054": 0x0472, "\u006f\u006d\u002dK\u0045": 0x1000, "\u006f\u0073": 0x1000, "\u006f\u0073\u002dG\u0045": 0x1000, "\u006f\u0073\u002dR\u0055": 0x1000, "\u0070\u0073": 0x0063, "\u0070\u0073\u002dA\u0046": 0x0463, "\u0070\u0073\u002dP\u004b": 0x1000, "\u0066\u0061": 0x0029, "\u0066\u0061\u002dA\u0046": 0x1000, "\u0066\u0061\u002dI\u0052": 0x0429, "\u0070\u006c": 0x0015, "\u0070\u006c\u002dP\u004c": 0x0415, "\u0070\u0074": 0x0016, "\u0070\u0074\u002dA\u004f": 0x1000, "\u0070\u0074\u002dB\u0052": 0x0416, "\u0070\u0074\u002dC\u0056": 0x1000, "\u0070\u0074\u002dG\u0051": 0x1000, "\u0070\u0074\u002dG\u0057": 0x1000, "\u0070\u0074\u002dL\u0055": 0x1000, "\u0070\u0074\u002dM\u004f": 0x1000, "\u0070\u0074\u002dM\u005a": 0x1000, "\u0070\u0074\u002dP\u0054": 0x0816, "\u0070\u0074\u002dS\u0054": 0x1000, "\u0070\u0074\u002dC\u0048": 0x1000, "\u0070\u0074\u002dT\u004c": 0x1000, "\u0070r\u0067\u002d\u0030\u0030\u0031": 0x1000, "\u0071p\u0073\u002d\u0070\u006c\u006f\u0063a": 0x05FE, "\u0071\u0070\u0073\u002d\u0070\u006c\u006f\u0063": 0x0501, "\u0071p\u0073\u002d\u0070\u006c\u006f\u0063m": 0x09FF, "\u0070\u0061": 0x0046, "\u0070a\u002d\u0041\u0072\u0061\u0062": 0x7C46, "\u0070\u0061\u002dI\u004e": 0x0446, "\u0070\u0061\u002d\u0041\u0072\u0061\u0062\u002d\u0050\u004b": 0x0846, "\u0071\u0075\u007a": 0x006B, "\u0071\u0075\u007a\u002d\u0042\u004f": 0x046B, "\u0071\u0075\u007a\u002d\u0045\u0043": 0x086B, "\u0071\u0075\u007a\u002d\u0050\u0045": 0x0C6B, "\u006b\u0073\u0068": 0x1000, "\u006b\u0073\u0068\u002d\u0044\u0045": 0x1000, "\u0072\u006f": 0x0018, "\u0072\u006f\u002dM\u0044": 0x0818, "\u0072\u006f\u002dR\u004f": 0x0418, "\u0072\u006d": 0x0017, "\u0072\u006d\u002dC\u0048": 0x0417, "\u0072\u006f\u0066": 0x1000, "\u0072\u006f\u0066\u002d\u0054\u005a": 0x1000, "\u0072\u006e": 0x1000, "\u0072\u006e\u002dB\u0049": 0x1000, "\u0072\u0075\u002dB\u0059": 0x1000, "\u0072\u0075\u002dK\u005a": 0x1000, "\u0072\u0075\u002dK\u0047": 0x1000, "\u0072\u0075\u002dM\u0044": 0x0819, "\u0072\u0075\u002dR\u0055": 0x0419, "\u0072\u0075\u002dU\u0041": 0x1000, "\u0072\u0077\u006b": 0x1000, "\u0072\u0077\u006b\u002d\u0054\u005a": 0x1000, "\u0073\u0073\u0079": 0x1000, "\u0073\u0073\u0079\u002d\u0045\u0052": 0x1000, "\u0073\u0061\u0068": 0x0085, "\u0073\u0061\u0068\u002d\u0052\u0055": 0x0485, "\u0073\u0061\u0071": 0x1000, "\u0073\u0061\u0071\u002d\u004b\u0045": 0x1000, "\u0073\u006d\u006e": 0x703B, "\u0073\u006d\u006e\u002d\u0046\u0049": 0x243B, "\u0073\u006d\u006a": 0x7C3B, "\u0073\u006d\u006a\u002d\u004e\u004f": 0x103B, "\u0073\u0065": 0x003B, "\u0073\u0065\u002dF\u0049": 0x0C3B, "\u0073\u0065\u002dN\u004f": 0x043B, "\u0073\u0065\u002dS\u0045": 0x083B, "\u0073\u006d\u0073": 0x743B, "\u0073\u006d\u0073\u002d\u0046\u0049": 0x203B, "\u0073\u006d\u0061": 0x783B, "\u0073\u006d\u0061\u002d\u004e\u004f": 0x183B, "\u0073\u006d\u0061\u002d\u0053\u0045": 0x1C3B, "\u0073\u0067": 0x1000, "\u0073\u0067\u002dC\u0046": 0x1000, "\u0073\u0062\u0070": 0x1000, "\u0073\u0062\u0070\u002d\u0054\u005a": 0x1000, "\u0073\u0061": 0x004F, "\u0073\u0061\u002dI\u004e": 0x044F, "\u0067\u0064": 0x0091, "\u0067\u0064\u002dG\u0042": 0x0491, "\u0073\u0065\u0068": 0x1000, "\u0073\u0065\u0068\u002d\u004d\u005a": 0x1000, "\u0073r\u002d\u0043\u0079\u0072\u006c": 0x6C1A, "\u0073\u0072\u002d\u0043\u0079\u0072\u006c\u002d\u0042\u0041": 0x1C1A, "\u0073\u0072\u002d\u0043\u0079\u0072\u006c\u002d\u004d\u0045": 0x301A, "\u0073\u0072\u002d\u0043\u0079\u0072\u006c\u002d\u0052\u0053": 0x281A, "\u0073\u0072\u002d\u0043\u0079\u0072\u006c\u002d\u0043\u0053": 0x0C1A, "\u0073r\u002d\u004c\u0061\u0074\u006e": 0x701A, "\u0073\u0072": 0x7C1A, "\u0073\u0072\u002d\u004c\u0061\u0074\u006e\u002d\u0042\u0041": 0x181A, "\u0073\u0072\u002d\u004c\u0061\u0074\u006e\u002d\u004d\u0045": 0x2c1A, "\u0073\u0072\u002d\u004c\u0061\u0074\u006e\u002d\u0052\u0053": 0x241A, "\u0073\u0072\u002d\u004c\u0061\u0074\u006e\u002d\u0043\u0053": 0x081A, "\u006e\u0073\u006f": 0x006C, "\u006e\u0073\u006f\u002d\u005a\u0041": 0x046C, "\u0074\u006e": 0x0032, "\u0074\u006e\u002dB\u0057": 0x0832, "\u0074\u006e\u002dZ\u0041": 0x0432, "\u006b\u0073\u0062": 0x1000, "\u006b\u0073\u0062\u002d\u0054\u005a": 0x1000, "\u0073\u006e": 0x1000, "\u0073n\u002d\u004c\u0061\u0074\u006e": 0x1000, "\u0073\u006e\u002d\u004c\u0061\u0074\u006e\u002d\u005a\u0057": 0x1000, "\u0073\u0064": 0x0059, "\u0073d\u002d\u0041\u0072\u0061\u0062": 0x7C59, "\u0073\u0064\u002d\u0041\u0072\u0061\u0062\u002d\u0050\u004b": 0x0859, "\u0073\u0069": 0x005B, "\u0073\u0069\u002dL\u004b": 0x045B, "\u0073\u006b": 0x001B, "\u0073\u006b\u002dS\u004b": 0x041B, "\u0073\u006c": 0x0024, "\u0073\u006c\u002dS\u0049": 0x0424, "\u0078\u006f\u0067": 0x1000, "\u0078\u006f\u0067\u002d\u0055\u0047": 0x1000, "\u0073\u006f": 0x0077, "\u0073\u006f\u002dD\u004a": 0x1000, "\u0073\u006f\u002dE\u0054": 0x1000, "\u0073\u006f\u002dK\u0045": 0x1000, "\u0073\u006f\u002dS\u004f": 0x0477, "\u0073\u0074": 0x0030, "\u0073\u0074\u002dZ\u0041": 0x0430, "\u006e\u0072": 0x1000, "\u006e\u0072\u002dZ\u0041": 0x1000, "\u0073\u0074\u002dL\u0053": 0x1000, "\u0065\u0073": 0x000A, "\u0065\u0073\u002dA\u0052": 0x2C0A, "\u0065\u0073\u002dB\u005a": 0x1000, "\u0065\u0073\u002dV\u0045": 0x200A, "\u0065\u0073\u002dB\u004f": 0x400A, "\u0065\u0073\u002dB\u0052": 0x1000, "\u0065\u0073\u002dC\u004c": 0x340A, "\u0065\u0073\u002dC\u004f": 0x240A, "\u0065\u0073\u002dC\u0052": 0x140A, "\u0065\u0073\u002dC\u0055": 0x5c0A, "\u0065\u0073\u002dD\u004f": 0x1c0A, "\u0065\u0073\u002dE\u0043": 0x300A, "\u0065\u0073\u002dS\u0056": 0x440A, "\u0065\u0073\u002dG\u0051": 0x1000, "\u0065\u0073\u002dG\u0054": 0x100A, "\u0065\u0073\u002dH\u004e": 0x480A, "\u0065\u0073\u002d\u0034\u0031\u0039": 0x580A, "\u0065\u0073\u002dM\u0058": 0x080A, "\u0065\u0073\u002dN\u0049": 0x4C0A, "\u0065\u0073\u002dP\u0041": 0x180A, "\u0065\u0073\u002dP\u0059": 0x3C0A, "\u0065\u0073\u002dP\u0045": 0x280A, "\u0065\u0073\u002dP\u0048": 0x1000, "\u0065\u0073\u002dP\u0052": 0x500A, "\u0065\u0073\u002dE\u0053\u005f\u0074\u0072\u0061\u0064\u006e\u006c": 0x040A, "\u0065\u0073\u002dE\u0053": 0x0c0A, "\u0065\u0073\u002dU\u0053": 0x540A, "\u0065\u0073\u002dU\u0059": 0x390A, "\u007a\u0067\u0068": 0x1000, "z\u0067\u0068\u002d\u0054\u0066\u006e\u0067\u002d\u004d\u0041": 0x1000, "\u007a\u0067\u0068\u002d\u0054\u0066\u006e\u0067": 0x1000, "\u0073\u0073": 0x1000, "\u0073\u0073\u002dZ\u0041": 0x1000, "\u0073\u0073\u002dS\u005a": 0x1000, "\u0073\u0076": 0x001D, "\u0073\u0076\u002dA\u0058": 0x1000, "\u0073\u0076\u002dF\u0049": 0x081D, "\u0073\u0076\u002dS\u0045": 0x041D, "\u0073\u0079\u0072": 0x005A, "\u0073\u0079\u0072\u002d\u0053\u0059": 0x045A, "\u0073\u0068\u0069": 0x1000, "\u0073\u0068\u0069\u002d\u0054\u0066\u006e\u0067": 0x1000, "s\u0068\u0069\u002d\u0054\u0066\u006e\u0067\u002d\u004d\u0041": 0x1000, "\u0073\u0068\u0069\u002d\u004c\u0061\u0074\u006e": 0x1000, "s\u0068\u0069\u002d\u004c\u0061\u0074\u006e\u002d\u004d\u0041": 0x1000, "\u0064\u0061\u0076": 0x1000, "\u0064\u0061\u0076\u002d\u004b\u0045": 0x1000, "\u0074\u0067": 0x0028, "\u0074g\u002d\u0043\u0079\u0072\u006c": 0x7C28, "\u0074\u0067\u002d\u0043\u0079\u0072\u006c\u002d\u0054\u006a": 0x0428, "\u0074\u007a\u006d": 0x005F, "\u0074\u007a\u006d\u002d\u004c\u0061\u0074\u006e": 0x7C5F, "t\u007a\u006d\u002d\u004c\u0061\u0074\u006e\u002d\u0044\u005a": 0x085F, "\u0074\u0061": 0x0049, "\u0074\u0061\u002dI\u004e": 0x0449, "\u0074\u0061\u002dM\u0059": 0x1000, "\u0074\u0061\u002dS\u0047": 0x1000, "\u0074\u0061\u002dL\u004b": 0x0849, "\u0074\u0077\u0071": 0x1000, "\u0074\u0077\u0071\u002d\u004e\u0045": 0x1000, "\u0074\u0074": 0x0044, "\u0074\u0074\u002dR\u0055": 0x0444, "\u0074\u0065": 0x004A, "\u0074\u0065\u002dI\u004e": 0x044A, "\u0074\u0065\u006f": 0x1000, "\u0074\u0065\u006f\u002d\u004b\u0045": 0x1000, "\u0074\u0065\u006f\u002d\u0055\u0047": 0x1000, "\u0074\u0068": 0x001E, "\u0074\u0068\u002dT\u0048": 0x041E, "\u0062\u006f": 0x0051, "\u0062\u006f\u002dI\u004e": 0x1000, "\u0062\u006f\u002dC\u004e": 0x0451, "\u0074\u0069\u0067": 0x1000, "\u0074\u0069\u0067\u002d\u0045\u0052": 0x1000, "\u0074\u0069": 0x0073, "\u0074\u0069\u002dE\u0052": 0x0873, "\u0074\u0069\u002dE\u0054": 0x0473, "\u0074\u006f": 0x1000, "\u0074\u006f\u002dT\u004f": 0x1000, "\u0074\u0073": 0x0031, "\u0074\u0073\u002dZ\u0041": 0x0431, "\u0074\u0072": 0x001F, "\u0074\u0072\u002dC\u0059": 0x1000, "\u0074\u0072\u002dT\u0052": 0x041F, "\u0074\u006b": 0x0042, "\u0074\u006b\u002dT\u004d": 0x0442, "\u0075\u006b": 0x0022, "\u0075\u006b\u002dU\u0041": 0x0422, "\u0068\u0073\u0062": 0x002E, "\u0068\u0073\u0062\u002d\u0044\u0045": 0x042E, "\u0075\u0072": 0x0020, "\u0075\u0072\u002dI\u004e": 0x0820, "\u0075\u0067": 0x0080, "\u0075\u0067\u002dC\u004e": 0x0480, "\u0075z\u002d\u0041\u0072\u0061\u0062": 0x1000, "\u0075\u007a\u002d\u0041\u0072\u0061\u0062\u002d\u0041\u0046": 0x1000, "\u0075z\u002d\u0043\u0079\u0072\u006c": 0x7843, "\u0075\u007a\u002d\u0043\u0079\u0072\u006c\u002d\u0055\u005a": 0x0843, "\u0075\u007a": 0x0043, "\u0075z\u002d\u004c\u0061\u0074\u006e": 0x7C43, "\u0075\u007a\u002d\u004c\u0061\u0074\u006e\u002d\u0055\u005a": 0x0443, "\u0076\u0061\u0069": 0x1000, "\u0076\u0061\u0069\u002d\u0056\u0061\u0069\u0069": 0x1000, "v\u0061\u0069\u002d\u0056\u0061\u0069\u0069\u002d\u004c\u0052": 0x1000, "v\u0061\u0069\u002d\u004c\u0061\u0074\u006e\u002d\u004c\u0052": 0x1000, "\u0076\u0061\u0069\u002d\u004c\u0061\u0074\u006e": 0x1000, "\u0063\u0061\u002d\u0045\u0053\u002d\u0076\u0061\u006ce\u006e\u0063\u0069\u0061": 0x0803, "\u0076\u0065": 0x0033, "\u0076\u0065\u002dZ\u0041": 0x0433, "\u0076\u0069": 0x002A, "\u0076\u0069\u002dV\u004e": 0x042A, "\u0076\u006f": 0x1000, "\u0076\u006f\u002d\u0030\u0030\u0031": 0x1000, "\u0076\u0075\u006e": 0x1000, "\u0076\u0075\u006e\u002d\u0054\u005a": 0x1000, "\u0077\u0061\u0065": 0x1000, "\u0077\u0061\u0065\u002d\u0043\u0048": 0x1000, "\u0063\u0079": 0x0052, "\u0063\u0079\u002dG\u0042": 0x0452, "\u0077\u0061\u006c": 0x1000, "\u0077\u0061\u006c\u002d\u0045\u0054": 0x1000, "\u0077\u006f": 0x0088, "\u0077\u006f\u002dS\u004e": 0x0488, "\u0078\u0068": 0x0034, "\u0078\u0068\u002dZ\u0041": 0x0434, "\u0079\u0061\u0076": 0x1000, "\u0079\u0061\u0076\u002d\u0043\u004d": 0x1000, "\u0069\u0069": 0x0078, "\u0069\u0069\u002dC\u004e": 0x0478, "\u0079\u0069\u002d\u0030\u0030\u0031": 0x043D, "\u0079\u006f": 0x006A, "\u0079\u006f\u002dB\u004a": 0x1000, "\u0079\u006f\u002dN\u0047": 0x046A, "\u0064\u006a\u0065": 0x1000, "\u0064\u006a\u0065\u002d\u004e\u0045": 0x1000, "\u007a\u0075": 0x0035, "\u007a\u0075\u006c\u0075": 0x0435}

func (_dge *convertContext) newSpan() {
	_egda := &span{_bd: _dge._bbfa._dbb, _cf: _dge._bbfa._fdc}
	_dge._ecbg = _egda
	_dge._bbfa._dag = append(_dge._bbfa._dag, _egda)
}

func (_dedc *convertContext) makePdfImageFromGraphics(_ebdc *_eeb.Pic) (*_ea.Image, error) {
	if _gcae := _ebdc.BlipFill; _gcae != nil {
		if _abge := _gcae.Blip; _abge != nil {
			if _fgda := _abge.EmbedAttr; _fgda != nil {
				_addbd, _bccc := _dedc._cffa.GetImageObjByRelId(*_fgda)
				if _bccc != nil {
					return nil, _bccc
				}
				_cbcd, _bccc := _f.Open(_addbd.Path)
				if _bccc != nil {
					return nil, _bccc
				}
				_fddb, _bccc := _bg.ReadAll(_cbcd)
				if _bccc != nil {
					return nil, _bccc
				}
				_cead, _bccc := _dedc._fcgg.NewImageFromData(_fddb)
				if _bccc != nil {
					return nil, _bccc
				}
				return _cead, nil
			}
		}
	}
	return nil, nil
}

func (_efgc *convertContext) addCurrentParagraphFooterToCurrentPage() {
	_efgc.alignParagraph()
	_efgc._fbcd._cb = append(_efgc._fbcd._cb, _efgc._cecg)
}

const (
	_daga = "\u006di\u006e\u006f\u0072\u0046\u006f\u006et"
	_gfef = "\u006da\u006a\u006f\u0072\u0046\u006f\u006et"
	_dgca = "\u006d\u0061\u006a\u006f\u0072\u0045\u0061\u0073\u0074\u0041\u0073\u0069a\u0046\u006f\u006e\u0074"
	_acag = "\u006d\u0069\u006e\u006f\u0072\u0045\u0061\u0073\u0074\u0041\u0073\u0069a\u0046\u006f\u006e\u0074"
)

func _dab(_gd *_ea.Creator, _ecb *block) {
	_ecb._egb.SetPos(_ecb._aef, _ecb._aga)
	_gd.Draw(_ecb._egb)
	if _ecb._dgg {
		_a.DrawRectangle(_gd, &_a.Rectangle{Top: _ecb._aga, Bottom: _ecb._aga + _ecb._egb.Height(), Left: _ecb._aef, Right: _ecb._aef + _ecb._egb.Width()}, _ecb._ebg, _ecb._dc)
	}
}

func (_cfgfa *convertContext) assignPropsToAbsoluteParagraph(_dcea *_bc.CT_PPr, _afe *paragraph) (float64, float64) {
	_cfgfa._gffd = _dcea
	_dcea = _cceb(_dcea, _cfgfa._bcdegd, _cfgfa._fbef)
	_efbf := 12.4
	if _dcea == nil {
		return 0, 0
	}
	if _faca := _dcea.RPr; _faca != nil {
		_ffce := _gdeg(_faca.Sz, _faca.SzCs)
		if _efbf <= _ffce {
			_efbf = _ffce
		} else {
			_efbf = _ffce * _db
		}
		_afe._ef = _efbf
	}
	if _dcea.Jc != nil {
		switch _dcea.Jc.ValAttr {
		case _bc.ST_JcRight:
			_afe._ae = _ea.TextAlignmentRight
		case _bc.ST_JcCenter:
			_afe._ae = _ea.TextAlignmentCenter
		case _bc.ST_JcBoth:
			_afe._ae = _ea.TextAlignmentJustify
		case _bc.ST_JcEnd:
			_afe._ae = _ea.TextAlignmentRight
		default:
			_afe._ae = _ea.TextAlignmentLeft
		}
	}
	var _fbbd, _bgeg, _aeca, _gbbc, _caed float64
	if _cdabb := _dcea.Spacing; _cdabb != nil {
		if _gaaa := _cdabb.BeforeAttr; _gaaa != nil {
			if _gaaa.ST_UnsignedDecimalNumber != nil {
				_fbbd = _a.PointsFromTwips(int64(*_gaaa.ST_UnsignedDecimalNumber))
			}
		}
		if _gacd := _cdabb.AfterAttr; _gacd != nil {
			if _gacd.ST_UnsignedDecimalNumber != nil {
				_bgeg = _a.PointsFromTwips(int64(*_gacd.ST_UnsignedDecimalNumber))
			}
		}
		if _dbbf := _cdabb.LineAttr; _dbbf != nil {
			if _dbbf.Int64 != nil && *_dbbf.Int64 != 0 {
				if _cbaf := float64(*_dbbf.Int64) / 20; _cbaf > _efbf {
					_efbf = _cbaf
				}
			}
		}
	}
	if _dcea.ContextualSpacing != nil && _dgfgag(_dcea.ContextualSpacing) {
		_fbbd = 0
		_bgeg = 0
	}
	_afe._ef = _efbf
	if _afab := _dcea.Ind; _afab != nil {
		if _dfgf := _afab.FirstLineAttr; _dfgf != nil {
			if _dfgf.ST_UnsignedDecimalNumber != nil {
				_caed = _a.PointsFromTwips(int64(*_dfgf.ST_UnsignedDecimalNumber))
			}
		}
		if _bef := _afab.HangingAttr; _bef != nil {
			if _bef.ST_UnsignedDecimalNumber != nil {
				_caed -= _a.PointsFromTwips(int64(*_bef.ST_UnsignedDecimalNumber))
			}
		}
		if _fbfa := _afab.LeftAttr; _fbfa != nil {
			if _fbfa.Int64 != nil {
				_aeca = _a.PointsFromTwips(*_fbfa.Int64)
			}
		}
		if _dbde := _afab.RightAttr; _dbde != nil {
			if _dbde.Int64 != nil {
				_gbbc = _a.PointsFromTwips(*_dbde.Int64)
			}
		}
	}
	if _dcea.PBdr != nil {
		_fee := _cfgfa._fbcd._eeg.Right - _cfgfa._fbcd._eeg.Left
		_ecea := _cfgfa._fbcd._eeg.Bottom - _cfgfa._fbcd._eeg.Top
		if _decdc := _dcea.PBdr.Top; _decdc != nil {
			_gefb := 0.0
			if _abff := _decdc.SzAttr; _abff != nil {
				_gefb = float64(*_abff) * _fee / 4
			}
			_edea := 0.0
			if _addg := _decdc.SpaceAttr; _addg != nil {
				_edea = float64(*_addg) * _eag.Pixel72
			}
			var _eaaf _ea.Color
			if _deea := _decdc.ColorAttr; _deea != nil {
				if _afba := _deea.ST_HexColorAuto; _afba == _bc.ST_HexColorAutoAuto {
					_eaaf = _ea.ColorBlack
				}
				if _eega := _deea.ST_HexColorRGB; _eega != nil {
					_eaaf = _ea.ColorRGBFromHex("\u0023" + *_eega)
				}
			}
			_defc := &borderLine{_dca: _gefb, _ecc: _a.BorderPositionTop, _adc: 1.0, _afaf: _eaaf, _ed: _edea}
			_afe._gcf = append(_afe._gcf, _defc)
		}
		if _fbbg := _dcea.PBdr.Left; _fbbg != nil {
			_eefa := 0.0
			if _bafa := _fbbg.SzAttr; _bafa != nil {
				_eefa = float64(*_bafa) * _ecea / 4
			}
			_geff := 0.0
			if _bffa := _fbbg.SpaceAttr; _bffa != nil {
				_geff = float64(*_bffa) * _eag.Pixel72
			}
			var _fbbb _ea.Color
			if _daagf := _fbbg.ColorAttr; _daagf != nil {
				if _dbgb := _daagf.ST_HexColorAuto; _dbgb == _bc.ST_HexColorAutoAuto {
					_fbbb = _ea.ColorBlack
				}
				if _bcc := _daagf.ST_HexColorRGB; _bcc != nil {
					_fbbb = _ea.ColorRGBFromHex("\u0023" + *_bcc)
				}
			}
			_gcfaa := &borderLine{_ecc: _a.BorderPositionLeft, _dca: 1.0, _adc: _eefa, _afaf: _fbbb, _ed: _geff}
			_afe._gcf = append(_afe._gcf, _gcfaa)
		}
		if _addgf := _dcea.PBdr.Right; _addgf != nil {
			_gdf := 0.0
			if _bbbe := _addgf.SzAttr; _bbbe != nil {
				_gdf = float64(*_bbbe) * _ecea / 4
			}
			_eccfe := 0.0
			if _egac := _addgf.SpaceAttr; _egac != nil {
				_eccfe = float64(*_egac) * _eag.Pixel72
			}
			var _acdb _ea.Color
			if _ggeb := _addgf.ColorAttr; _ggeb != nil {
				if _cfc := _ggeb.ST_HexColorAuto; _cfc == _bc.ST_HexColorAutoAuto {
					_acdb = _ea.ColorBlack
				}
				if _bcabe := _ggeb.ST_HexColorRGB; _bcabe != nil {
					_acdb = _ea.ColorRGBFromHex("\u0023" + *_bcabe)
				}
			}
			_gbed := &borderLine{_ecc: _a.BorderPositionRight, _dca: 1.0, _adc: _gdf, _afaf: _acdb, _ed: _eccfe}
			_afe._gcf = append(_afe._gcf, _gbed)
		}
		if _acgaa := _dcea.PBdr.Bottom; _acgaa != nil {
			_dafab := 0.0
			if _eegdd := _acgaa.SzAttr; _eegdd != nil {
				_dafab = float64(*_eegdd) * _fee / 4
			}
			_bcgfg := 0.0
			if _bffgf := _acgaa.SpaceAttr; _bffgf != nil {
				_bcgfg = float64(*_bffgf) * _eag.Pixel72
			}
			var _fbae _ea.Color
			if _bebe := _acgaa.ColorAttr; _bebe != nil {
				if _fda := _bebe.ST_HexColorAuto; _fda == _bc.ST_HexColorAutoAuto {
					_fbae = _ea.ColorBlack
				}
				if _afdd := _bebe.ST_HexColorRGB; _afdd != nil {
					_fbae = _ea.ColorRGBFromHex("\u0023" + *_afdd)
				}
			}
			_gbdb := &borderLine{_ecc: _a.BorderPositionBottom, _dca: _dafab, _adc: 1.0, _afaf: _fbae, _ed: _bcgfg}
			_afe._gcf = append(_afe._gcf, _gbdb)
		}
	}
	_dgfga := _cfgfa._fbcd._eb
	if len(_dgfga) > 0 {
		_fbbd -= _dgfga[len(_dgfga)-1]._ac.Bottom
		if _fbbd < 0 {
			_fbbd = 0
		}
	} else {
		_fbbd -= _dd
	}
	_afe._ac = &_a.Rectangle{Top: _fbbd, Bottom: _bgeg, Left: _aeca, Right: _gbbc}
	_afe._af = _caed
	return _fbbd, _aeca
}

// ConvertToPdfWithOptions convert the document to PDF with file given options.
func ConvertToPdfWithOptions(d *_e.Document, opts *Options) *_ea.Creator {
	var _cbc map[string]string
	if opts != nil {
		if opts.ProcessFields {
			_cbc = _bgfa(d)
		}
		if len(opts.FontFiles) > 0 {
			_geef := _a.RegisterFontsFromFiles(opts.FontFiles)
			if _geef != nil {
				_da.Log.Debug("\u0046\u0061\u0069\u006c t\u006f\u0020\u006c\u006f\u0061\u0064\u0020\u0066\u006f\u006e\u0074\u0073\u003a\u0020%\u0076", opts.FontDirectory)
			}
		}
		if opts.FontDirectory != "" {
			_afdcc := _a.RegisterFontsFromDirectory(opts.FontDirectory)
			if _afdcc != nil {
				_da.Log.Debug("\u0046\u0061\u0069l\u0020\u0074\u006f\u0020l\u006f\u0061\u0064\u0020\u0066\u006f\u006et\u0020\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u003a\u0020\u0025\u0076", _afdcc.Error())
			}
		}
	}
	var _ecfga *_bc.CT_PPrGeneral
	var _bgbbf *_bc.CT_RPr
	if _agge := d.Styles.X().DocDefaults; _agge != nil {
		if _abcd := _agge.PPrDefault; _abcd != nil {
			_ecfga = _abcd.PPr
		}
		if _bbdc := _agge.RPrDefault; _bbdc != nil {
			_bgbbf = _bbdc.RPr
		}
	}
	_gbdf, _ccfe := _ffcdc(210), _ffcdc(297)
	_ggdc := float64(_eag.Inch * 1)
	_dcgdg := _a.PointsFromTwips(720)
	_ecad, _bagc, _ddae, _cgaae := _ggdc, _ggdc, _ggdc, _ggdc
	var (
		_gdca, _fdbf float64
		_gdgc        []*headerFooterRef
		_efffe       float64
	)
	if _ggef := d.BodySection().X(); _ggef != nil {
		if _cbab := _ggef.PgSz; _cbab != nil {
			if _edfg := _ggef.PgMar; _edfg != nil {
				if _edfg.LeftAttr.ST_UnsignedDecimalNumber != nil {
					_ecad = _a.PointsFromTwips(int64(*_edfg.LeftAttr.ST_UnsignedDecimalNumber))
				}
				if _edfg.LeftAttr.ST_UnsignedDecimalNumber != nil {
					_bagc = _a.PointsFromTwips(int64(*_edfg.RightAttr.ST_UnsignedDecimalNumber))
				}
				if _edfg.TopAttr.Int64 != nil {
					_ddae = _a.PointsFromTwips(*_edfg.TopAttr.Int64)
				}
				if _edfg.BottomAttr.Int64 != nil {
					_cgaae = _a.PointsFromTwips(*_edfg.BottomAttr.Int64)
				}
				if _edfg.HeaderAttr.ST_UnsignedDecimalNumber != nil {
					_gdca = _a.PointsFromTwips(int64(*_edfg.HeaderAttr.ST_UnsignedDecimalNumber))
				}
				if _edfg.FooterAttr.ST_UnsignedDecimalNumber != nil {
					_fdbf = _cgaae - _a.PointsFromTwips(int64(*_edfg.FooterAttr.ST_UnsignedDecimalNumber))
					_efffe = _cgaae + _a.PointsFromTwips(int64(*_edfg.FooterAttr.ST_UnsignedDecimalNumber))
				}
			}
			if _cbab.WAttr != nil {
				_gbdf = _a.PointsFromTwips(int64(*_cbab.WAttr.ST_UnsignedDecimalNumber))
			}
			if _cbab.HAttr != nil {
				_ccfe = _a.PointsFromTwips(int64(*_cbab.HAttr.ST_UnsignedDecimalNumber))
			}
		}
		for _, _adeb := range _ggef.EG_HdrFtrReferences {
			if _cggg := _adeb.HeaderReference; _cggg != nil {
				_cbfe := &headerFooterRef{_eadd: true, _fgcc: _cggg.IdAttr, _gdae: _cggg.TypeAttr}
				_gdgc = append(_gdgc, _cbfe)
				if _gdca <= 0 {
					_gdca = _dcgdg
					_ddae = _ddae + _gdca
				}
			}
			if _dacg := _adeb.FooterReference; _dacg != nil {
				_dcbf := &headerFooterRef{_fcfcfa: true, _fgcc: _dacg.IdAttr, _gdae: _dacg.TypeAttr}
				_gdgc = append(_gdgc, _dcbf)
				if _fdbf <= 0 {
					_efffe = _dcgdg
				}
			}
		}
	}
	if d.Settings.X().DefaultTabStop == nil {
		_debg = _ffcdc(12.7)
	} else {
		_debg = _a.PointsFromTwips(int64(*d.Settings.X().DefaultTabStop.ValAttr.ST_UnsignedDecimalNumber))
	}
	_cgedg := _ea.New()
	_cgedg.SetPageSize(_ea.PageSize{_gbdf, _ccfe})
	_cgedg.SetPageMargins(_ecad, _bagc, _ddae, _efffe)
	_bacg := &convertContext{_fcgg: _cgedg, _cffa: d, _bcdegd: _ecfga, _fbef: _bgbbf, _eaebc: &_a.Rectangle{Top: _ddae, Bottom: _ccfe - _cgaae, Left: _ecad, Right: _gbdf - _bagc}, _agae: []note{}, _fbgf: map[int64]map[int64]int64{}, _cgaa: _cbc, _afff: opts, _dceg: _gdgc, _ggfa: _gdca, _cfbg: _fdbf, _bcce: _ecad}
	for _efec, _adfd := range d.X().Body.EG_BlockLevelElts {
		var _eacd []*_bc.EG_ContentBlockContent
		if _efec < len(d.X().Body.EG_BlockLevelElts)-1 {
			_edgc := d.X().Body.EG_BlockLevelElts[_efec+1]
			_eacd = _edgc.EG_ContentBlockContent
		}
		_bacg.addAbsoluteCBCs(_adfd.EG_ContentBlockContent, _eacd)
	}
	_bacg.addEndnotes()
	_bacg.alignSymbolsVertically()
	_bacg.drawPages()
	_bacg.drawHeaderFooter()
	return _cgedg
}

type tableWrapper struct {
	_cca *_ea.Table
	_fb  float64
}

type zoneToSkip struct {
	_gac *_a.Rectangle
	_aba *_bc.WdEG_WrapTypeChoice
}

func (_adb *convertContext) combinePPrWithStyles(_feca *_bc.CT_PPr) (*_bc.CT_PPr, *prefix) {
	if _feca == nil {
		return nil, nil
	}
	var _bbge *prefix
	if _feca != nil && _feca.NumPr != nil {
		if _aefd, _cddb := _feca.NumPr.Ilvl, _feca.NumPr.NumId; _aefd != nil && _cddb != nil {
			if _gacg := _adb._cffa.GetNumberingLevelByIds(_cddb.ValAttr, _aefd.ValAttr).X(); _gacg != nil {
				_feca = _cceb(_feca, _gacg.PPr, _gacg.RPr)
				if _bgdcd := _gacg.NumFmt; _bgdcd != nil {
					if _aebcg := _bgdcd.ValAttr; _aebcg != _bc.ST_NumberFormatNone && _aebcg != _bc.ST_NumberFormatCustom {
						var _fdcfa []float64
						if _eeac := _feca.Tabs; _eeac != nil && len(_eeac.Tab) != 0 {
							for _, _dagad := range _eeac.Tab {
								_fdcfa = append(_fdcfa, _a.PointsFromTwips(*_dagad.PosAttr.Int64))
							}
						}
						_bbge = &prefix{_aceb: _fdcfa, _eaee: true}
						if _aebcg == _bc.ST_NumberFormatBullet {
							if _afgde := _gacg.LvlText; _afgde != nil {
								if _defcg := _afgde.ValAttr; _defcg != nil && *_defcg != "" {
									_bbge._ebab = *_defcg
									_bbge._aedd = true
								}
							}
						} else {
							_bfec, _dbaf := _cddb.ValAttr, _aefd.ValAttr
							if _, _ccagb := _adb._fbgf[_bfec]; !_ccagb {
								_adb._fbgf[_bfec] = map[int64]int64{}
							}
							if _, _babg := _adb._fbgf[_bfec][_dbaf]; !_babg {
								_adb._fbgf[_bfec][_dbaf] = 1
								if _cddd := _gacg.Start; _cddd != nil {
									_adb._fbgf[_bfec][_dbaf] = _cddd.ValAttr
								}
							}
							if _, _abdee := _adb._fbgf[_bfec][_dbaf+1]; _abdee {
								_adb._fbgf[_bfec][_dbaf+1] = 1
							}
							_geeg := _adb._fbgf[_bfec][_dbaf]
							_agcb := _ge.FormatNumberingText(int64(_geeg), _gacg.IlvlAttr, *_gacg.LvlText.ValAttr, _gacg.NumFmt, _adb._fbgf[_bfec])
							_adb._fbgf[_bfec][_dbaf]++
							_bbge._ebab = _agcb
						}
					}
				}
			}
		}
	}
	if _feca != nil && _feca.PStyle != nil {
		_fedd, _dafc := _adb.getStyleProps(_feca.PStyle.ValAttr, _e.Style{})
		_feca = _cceb(_feca, _fedd, _dafc)
	}
	return _feca, _bbge
}

var (
	_ga = _ffcdc(6)
	_fg = _ffcdc(0.25)
	_fd = _ffcdc(1.9)
)

const (
	_ca = 0.67
	_db = 1.15
	_dd = 2.5
)

func (_bbfg *span) moveRight(_ebfb float64) {
	for _, _deba := range _bbfg._bda {
		_deba._cfe += _ebfb
	}
}

const (
	FontStyle_Regular    FontStyle = 0
	FontStyle_Bold       FontStyle = 1
	FontStyle_Italic     FontStyle = 2
	FontStyle_BoldItalic FontStyle = 3
)

var _gdff = []romanMatch{romanMatch{1000, "\u006d"}, romanMatch{900, "\u0063\u006d"}, romanMatch{500, "\u0064"}, romanMatch{400, "\u0063\u0064"}, romanMatch{100, "\u0063"}, romanMatch{90, "\u0078\u0063"}, romanMatch{50, "\u006c"}, romanMatch{40, "\u0078\u006c"}, romanMatch{10, "\u0078"}, romanMatch{9, "\u0069\u0078"}, romanMatch{5, "\u0076"}, romanMatch{4, "\u0069\u0076"}, romanMatch{1, "\u0069"}}

type symbol struct {
	_fge string
	_cdb float64
	_afa float64
	_cdf float64
	_cgb float64
	_ggc float64
	_ba  *_ea.TextStyle
	_bea *_ea.Image
	_cce *block
	_eae string
	_ffa bool
	_bcb bool
	_ad  *_ea.Color
	_fbb bool
	_fa  bool
}

func (_dbbb *convertContext) addCurrentWordToParagraph() {
	for {
		_cfeag := _dbbb._bbfa._dbb
		_agbg := _cfeag + _dbbb._gced._bfdc
		if _agbg > _dbbb._bbfa._fdc {
			_dbbb.newLine()
		}
		_addf := _dbbb._cecg._bgb + _dbbb._bbfa._ag
		_eeab := _addf + _dbbb._bbfa._bca
		_bcg := false
		_abaf := append(_dbbb._fbcd._gf, _dbbb._cecg._fcd...)
		for _, _eggf := range _abaf {
			_bgcaa := _eggf._gac
			if ((_addf > _bgcaa.Top && _addf < _bgcaa.Bottom) || (_eeab > _bgcaa.Top && _eeab < _bgcaa.Bottom)) && ((_cfeag > _bgcaa.Left && _cfeag < _bgcaa.Right) || (_agbg > _bgcaa.Left && _agbg < _bgcaa.Right)) {
				_bcg = true
				if _dbbb._bbfa._dbb < _bgcaa.Right {
					_dbbb._ecbg._cf = _bgcaa.Left
					_dbbb._bbfa._dbb = _bgcaa.Right
					_dbbb.newSpan()
				}
			}
		}
		if !_bcg {
			break
		}
	}
	if !_dbbb._gced._gfb || len(_dbbb._ecbg._bda) > 0 {
		_dbbb._gced._cfe = _dbbb._bbfa._dbb
		_dbbb._ecbg._bda = append(_dbbb._ecbg._bda, _dbbb._gced)
		_dbbb._bbfa._dbb += _dbbb._gced._bfdc
		for _, _bbg := range _dbbb._gced._gaa {
			_dbbb.adjustHeights(_bbg._cgb)
		}
	}
}

func (_ged *convertContext) alignSymbolsVertically() {
	for _, _dfb := range _ged._efceb {
		for _, _cgfd := range _dfb._eb {
			for _, _cced := range _cgfd._cc {
				_fcc := 0.0
				for _, _gea := range _cced._dag {
					for _, _dga := range _gea._bda {
						for _, _dfa := range _dga._gaa {
							if _dfa._cgb > _fcc {
								_fcc = _dfa._cgb
							}
						}
					}
				}
				for _, _gcb := range _cced._dag {
					for _, _dde := range _gcb._bda {
						for _, _ebf := range _dde._gaa {
							if _ebf._ggc < _fcc {
								_ebf._afa = _fcc - _ebf._ggc
							}
						}
					}
				}
			}
		}
	}
}

func (_fca *convertContext) addTextSymbol(_eef *symbol) {
	_gcfg := _ea.New()
	_bee := _gcfg.NewStyledParagraph()
	_bee.SetMargins(0, 0, 0, 0)
	_egg := _bee.Append(_eef._fge)
	_fae := 0.0
	if _eef._ba != nil {
		_egg.Style = *_eef._ba
		if _eef._ba.CharSpacing != 0 {
			_fae = _eef._ba.CharSpacing
		}
	}
	if _eef._cce == nil && _eef._bea == nil {
		_eef._cgb = _bee.Height() * _db
		_eef._ggc = _bee.Height()
	}
	if _eef._cdf == 0 {
		_eef._cdf = _bee.Width() + _fae
	}
	if _eef._cgb < _fca._cecg._ef {
		_eef._cgb = _fca._cecg._ef
	}
	if len(_fca._gced._gaa) > 0 {
		_gcac := _fca._gced._gaa[len(_fca._gced._gaa)-1]._fge
		if _a.IsNoSpaceLanguage(_gcac) || (_gcac == "\u0020") != (_eef._fge == "\u0020") {
			_fca.addCurrentWordToParagraph()
			_fca.newWord()
		}
	}
	_fca._gced._gaa = append(_fca._gced._gaa, _eef)
	_eef._cdb = _fca._gced._bfdc
	_fca._gced._bfdc += _eef._cdf
	if _eef._fge != "\u0020" {
		_fca._gced._gfb = false
	}
	if _eef._fge == "\u000d" {
		_fca.adjustHeights(_eef._cgb * 1.13)
		_fca.adjustHeights(_eef._cgb)
	}
}

func _cacg(_dbgc, _bagb *_bc.CT_TcPr) *_bc.CT_TcPr {
	if _dbgc == nil {
		return _bagb
	}
	if _bagb == nil {
		return _dbgc
	}
	if _dbgc.CnfStyle == nil {
		_dbgc.CnfStyle = _bagb.CnfStyle
	}
	if _dbgc.TcW == nil {
		_dbgc.TcW = _bagb.TcW
	}
	if _dbgc.GridSpan == nil {
		_dbgc.GridSpan = _bagb.GridSpan
	}
	if _dbgc.HMerge == nil {
		_dbgc.HMerge = _bagb.HMerge
	}
	if _dbgc.VMerge == nil {
		_dbgc.VMerge = _bagb.VMerge
	}
	if _dbgc.TcBorders == nil {
		_dbgc.TcBorders = _bagb.TcBorders
	}
	if _dbgc.Shd == nil {
		_dbgc.Shd = _bagb.Shd
	}
	if _dbgc.NoWrap == nil {
		_dbgc.NoWrap = _bagb.NoWrap
	}
	if _dbgc.TcMar == nil {
		_dbgc.TcMar = _bagb.TcMar
	}
	if _dbgc.TextDirection == nil {
		_dbgc.TextDirection = _bagb.TextDirection
	}
	if _dbgc.TcFitText == nil {
		_dbgc.TcFitText = _bagb.TcFitText
	}
	if _dbgc.VAlign == nil {
		_dbgc.VAlign = _bagb.VAlign
	}
	if _dbgc.HideMark == nil {
		_dbgc.HideMark = _bagb.HideMark
	}
	if _dbgc.Headers == nil {
		_dbgc.Headers = _bagb.Headers
	}
	if _dbgc.CellIns == nil {
		_dbgc.CellIns = _bagb.CellIns
	}
	if _dbgc.CellDel == nil {
		_dbgc.CellDel = _bagb.CellDel
	}
	if _dbgc.CellMerge == nil {
		_dbgc.CellMerge = _bagb.CellMerge
	}
	if _dbgc.TcPrChange == nil {
		_dbgc.TcPrChange = _bagb.TcPrChange
	}
	return _dbgc
}

func init() {
	_abac = _bf.MustCompile("\u0053E\u0054 \u0028\u002e\u002b\u0029\u0020\u0022\u0028\u002e\u002b\u0029\u0022")
	_bgdc = _bf.MustCompile("\u0052\u0045\u0046\u0020\u0028\u002e\u002b\u003f\u0029\u0020")
}

func (_ggbg *convertContext) drawHeaderFooter() {
	_ggbg._fcgg.DrawHeader(func(_ebde *_ea.Block, _bade _ea.HeaderFunctionArgs) {
		_gagb := _ggbg._efceb[_bade.PageNum-1]
		if len(_gagb._eba) < 1 && len(_ggbg._dceg) > 0 {
			_gagb._eba = _ggbg._dceg
		} else {
			for _dabca := _bade.PageNum - 2; _dabca >= 0; _dabca-- {
				_baccc := _ggbg._efceb[_dabca]
				if len(_baccc._eba) > 0 {
					_gagb._eba = _baccc._eba
					break
				}
			}
		}
		_ggbg._fbcd = _gagb
		_ggbg._fbcd._fga = nil
		_ggbg.assignHeaderFooterToPage(_gagb)
		_fcbcd(_ggbg._fcgg, _ebde, _ggbg._fbcd._fga, _ggbg._ggfa, _ggbg._fbcd._eeg.Bottom)
	})
	_ggbg._fcgg.DrawFooter(func(_gacdb *_ea.Block, _beaf _ea.FooterFunctionArgs) {
		_cefa := _ggbg._efceb[_beaf.PageNum-1]
		if len(_cefa._eba) < 1 && len(_ggbg._dceg) > 0 {
			_cefa._eba = _ggbg._dceg
		} else {
			for _bcac := _beaf.PageNum - 2; _bcac >= 0; _bcac-- {
				_ecef := _ggbg._efceb[_bcac]
				if len(_ecef._eba) > 0 {
					_cefa._eba = _ecef._eba
					break
				}
			}
		}
		_ggbg._fbcd = _cefa
		_ggbg._fbcd._cb = nil
		_ggbg.assignHeaderFooterToPage(_cefa)
		_fcbcd(_ggbg._fcgg, _gacdb, _ggbg._fbcd._cb, _ggbg._cfbg, _ggbg._fbcd._eeg.Bottom)
	})
}

func _ggbc(_adeba *_e.Document, _gbabf *_bc.CT_RPr, _cgfc *_bc.CT_PPr) *_bc.CT_RPr {
	var _ccgg, _fcbe *_bc.CT_RPr
	if _gbabf == nil {
		_gbabf = _bc.NewCT_RPr()
	}
	var _cbcf *_bc.CT_ParaRPr
	if _cgfc != nil && _cgfc.RPr != nil {
		_cbcf = _cgfc.RPr
	}
	if _cbcf == nil {
		_cbcf = _bc.NewCT_ParaRPr()
	}
	if _gbabf.RStyle != nil {
		_adfe := _adeba.GetStyleByID(_gbabf.RStyle.ValAttr)
		if _afdca := _adfe.X(); _afdca != nil {
			_ccgg = _afdca.RPr
		}
	}
	if _ccgg == nil {
		_ccgg = _bc.NewCT_RPr()
	}
	if _cbcf.RStyle != nil {
		_ccfc := _adeba.GetStyleByID(_cbcf.RStyle.ValAttr)
		if _acdbb := _ccfc.X(); _acdbb != nil {
			_fcbe = _acdbb.RPr
			if _acdbb.QFormat != nil && _dgfgag(_acdbb.QFormat) {
				return _fcbe
			}
		}
	}
	if _fcbe == nil {
		_fcbe = _bc.NewCT_RPr()
	}
	if _gbabf.Color == nil {
		if _ccgg.Color != nil {
			_gbabf.Color = _ccgg.Color
		} else if _cbcf.Color != nil {
			_gbabf.Color = _cbcf.Color
		} else if _fcbe.Color != nil {
			_gbabf.Color = _fcbe.Color
		}
	}
	if _gbabf.Spacing == nil {
		if _ccgg.Spacing != nil {
			_gbabf.Spacing = _ccgg.Spacing
		} else if _cbcf.Spacing != nil {
			_gbabf.Spacing = _cbcf.Spacing
		} else if _fcbe.Spacing != nil {
			_gbabf.Spacing = _fcbe.Spacing
		}
	}
	if _gbabf.Sz == nil {
		if _ccgg.Sz != nil {
			_gbabf.Sz = _ccgg.Sz
		} else if _cbcf.Sz != nil {
			_gbabf.Sz = _cbcf.Sz
		} else if _fcbe.Sz != nil {
			_gbabf.Sz = _fcbe.Sz
		}
	}
	if _gbabf.SzCs == nil {
		if _ccgg.SzCs != nil {
			_gbabf.SzCs = _ccgg.SzCs
		} else if _cbcf.SzCs != nil {
			_gbabf.SzCs = _cbcf.SzCs
		} else if _fcbe.SzCs != nil {
			_gbabf.SzCs = _fcbe.SzCs
		}
	}
	if _gbabf.B == nil {
		if _ccgg.B != nil {
			_gbabf.B = _ccgg.B
		} else if _cbcf.B != nil {
			_gbabf.B = _cbcf.B
		} else if _fcbe.B != nil {
			_gbabf.B = _fcbe.B
		}
	}
	if _gbabf.I == nil {
		if _ccgg.I != nil {
			_gbabf.I = _ccgg.I
		} else if _cbcf.I != nil {
			_gbabf.I = _cbcf.I
		} else if _fcbe.I != nil {
			_gbabf.I = _fcbe.I
		}
	}
	if _gbabf.U == nil {
		if _ccgg.U != nil {
			_gbabf.U = _ccgg.U
		} else if _cbcf.U != nil {
			_gbabf.U = _cbcf.U
		} else if _fcbe.U != nil {
			_gbabf.U = _fcbe.U
		}
	}
	if _gbabf.RFonts == nil {
		if _ccgg.RFonts != nil {
			_gbabf.RFonts = _ccgg.RFonts
		} else if _cbcf.RFonts != nil {
			_gbabf.RFonts = _cbcf.RFonts
		} else if _fcbe.RFonts != nil {
			_gbabf.RFonts = _fcbe.RFonts
		}
	}
	if _gbabf.VertAlign == nil {
		if _ccgg.VertAlign != nil {
			_gbabf.VertAlign = _ccgg.VertAlign
		} else if _cbcf.VertAlign != nil {
			_gbabf.VertAlign = _cbcf.VertAlign
		} else if _fcbe.VertAlign != nil {
			_gbabf.VertAlign = _fcbe.VertAlign
		}
	}
	if _gbabf.Caps == nil {
		if _ccgg.Caps != nil {
			_gbabf.Caps = _ccgg.Caps
		} else if _cbcf.Caps != nil {
			_gbabf.Caps = _cbcf.Caps
		} else if _fcbe.Caps != nil {
			_gbabf.Caps = _fcbe.Caps
		}
	}
	if _gbabf.SmallCaps == nil {
		if _ccgg.SmallCaps != nil {
			_gbabf.SmallCaps = _ccgg.SmallCaps
		} else if _cbcf.SmallCaps != nil {
			_gbabf.SmallCaps = _cbcf.SmallCaps
		} else if _fcbe.SmallCaps != nil {
			_gbabf.SmallCaps = _fcbe.SmallCaps
		}
	}
	if _gbabf.Bdr == nil {
		if _ccgg.Bdr != nil {
			_gbabf.Bdr = _ccgg.Bdr
		} else if _cbcf.Bdr != nil {
			_gbabf.Bdr = _cbcf.Bdr
		} else if _fcbe.Bdr != nil {
			_gbabf.Bdr = _fcbe.Bdr
		}
	}
	if _gbabf.Shd == nil {
		if _ccgg.Shd != nil {
			_gbabf.Shd = _ccgg.Shd
		} else if _cbcf.Shd != nil {
			_gbabf.Shd = _cbcf.Shd
		} else if _fcbe.Shd != nil {
			_gbabf.Shd = _fcbe.Shd
		}
	}
	return _gbabf
}

func _dbea(_fgae, _efd *_bc.CT_RPr) *_bc.CT_RPr {
	if _fgae == nil {
		return _efd
	}
	if _efd == nil {
		if _fgae.B != nil {
			_fgae.B = nil
		}
		if _fgae.BCs != nil {
			_fgae.BCs = nil
		}
		if _fgae.I != nil {
			_fgae.I = nil
		}
		if _fgae.ICs != nil {
			_fgae.ICs = nil
		}
		return _fgae
	}
	if _fgae.RStyle == nil {
		_fgae.RStyle = _efd.RStyle
	}
	if _fgae.RFonts == nil {
		_fgae.RFonts = _efd.RFonts
	}
	if _fgae.B == nil {
		_fgae.B = _efd.B
	}
	if _fgae.BCs == nil {
		_fgae.BCs = _efd.BCs
	}
	if _fgae.I == nil {
		_fgae.I = _efd.I
	}
	if _fgae.ICs == nil {
		_fgae.ICs = _efd.ICs
	}
	if _fgae.Caps == nil {
		_fgae.Caps = _efd.Caps
	}
	if _fgae.SmallCaps == nil {
		_fgae.SmallCaps = _efd.SmallCaps
	}
	if _fgae.Strike == nil {
		_fgae.Strike = _efd.Strike
	}
	if _fgae.Dstrike == nil {
		_fgae.Dstrike = _efd.Dstrike
	}
	if _fgae.Outline == nil {
		_fgae.Outline = _efd.Outline
	}
	if _fgae.Shadow == nil {
		_fgae.Shadow = _efd.Shadow
	}
	if _fgae.Emboss == nil {
		_fgae.Emboss = _efd.Emboss
	}
	if _fgae.Imprint == nil {
		_fgae.Imprint = _efd.Imprint
	}
	if _fgae.NoProof == nil {
		_fgae.NoProof = _efd.NoProof
	}
	if _fgae.SnapToGrid == nil {
		_fgae.SnapToGrid = _efd.SnapToGrid
	}
	if _fgae.Vanish == nil {
		_fgae.Vanish = _efd.Vanish
	}
	if _fgae.WebHidden == nil {
		_fgae.WebHidden = _efd.WebHidden
	}
	if _fgae.Color == nil {
		_fgae.Color = _efd.Color
	}
	if _fgae.Spacing == nil {
		_fgae.Spacing = _efd.Spacing
	}
	if _fgae.W == nil {
		_fgae.W = _efd.W
	}
	if _fgae.Kern == nil {
		_fgae.Kern = _efd.Kern
	}
	if _fgae.Position == nil {
		_fgae.Position = _efd.Position
	}
	if _fgae.Sz == nil {
		_fgae.Sz = _efd.Sz
	}
	if _fgae.SzCs == nil {
		_fgae.SzCs = _efd.SzCs
	}
	if _fgae.Highlight == nil {
		_fgae.Highlight = _efd.Highlight
	}
	if _fgae.U == nil {
		_fgae.U = _efd.U
	}
	if _fgae.Effect == nil {
		_fgae.Effect = _efd.Effect
	}
	if _fgae.Bdr == nil {
		_fgae.Bdr = _efd.Bdr
	}
	if _fgae.Shd == nil {
		_fgae.Shd = _efd.Shd
	}
	if _fgae.FitText == nil {
		_fgae.FitText = _efd.FitText
	}
	if _fgae.VertAlign == nil {
		_fgae.VertAlign = _efd.VertAlign
	}
	if _fgae.Rtl == nil {
		_fgae.Rtl = _efd.Rtl
	}
	if _fgae.Cs == nil {
		_fgae.Cs = _efd.Cs
	}
	if _fgae.Em == nil {
		_fgae.Em = _efd.Em
	}
	if _fgae.Lang == nil {
		_fgae.Lang = _efd.Lang
	}
	if _fgae.EastAsianLayout == nil {
		_fgae.EastAsianLayout = _efd.EastAsianLayout
	}
	if _fgae.SpecVanish == nil {
		_fgae.SpecVanish = _efd.SpecVanish
	}
	if _fgae.OMath == nil {
		_fgae.OMath = _efd.OMath
	}
	if _fgae.RPrChange == nil {
		_fgae.RPrChange = _efd.RPrChange
	}
	return _fgae
}

func _dccef(_dcde *_e.Document, _agba string, _gbba *_bc.CT_TblPr, _ebea *_bc.CT_PPrGeneral, _cddde *_bc.CT_RPr) (*_bc.CT_TblPr, *_bc.CT_PPrGeneral, *_bc.CT_RPr) {
	if _gbba.TblStyle != nil {
		_debfa := _dcde.GetStyleByID(_agba)
		if _dagf := _debfa.X(); _dagf != nil {
			if _feeg := _dagf.TblPr; _feeg != nil {
				_eggc := _gbba.TblBorders
				var _eebfc *_bc.CT_TblBorders
				if _feeg.TblBorders != nil {
					_eebfc = _feeg.TblBorders
				}
				if _eggc == nil {
					_eggc = _eebfc
				} else {
					if _eebfc != nil {
						if _eggc.Top == nil {
							_eggc.Top = _eebfc.Top
						}
						if _eggc.Bottom == nil {
							_eggc.Bottom = _eebfc.Bottom
						}
						if _eggc.Left == nil {
							_eggc.Left = _eebfc.Left
						}
						if _eggc.Right == nil {
							_eggc.Right = _eebfc.Right
						}
						if _eggc.InsideH == nil {
							_eggc.InsideH = _eebfc.InsideH
						}
						if _eggc.InsideV == nil {
							_eggc.InsideV = _eebfc.InsideV
						}
					}
				}
				_gbba.TblBorders = _eggc
				_aad := _gbba.Shd
				_fcaef := _feeg.Shd
				if _aad == nil {
					_aad = _fcaef
				} else {
					if _fcaef != nil && _aad.FillAttr == nil {
						_aad.FillAttr = _fcaef.FillAttr
					}
				}
				_gbba.Shd = _aad
				_fbaf := _gbba.TblCellMar
				_cagf := _feeg.TblCellMar
				if _fbaf == nil {
					_fbaf = _cagf
				} else {
					if _cagf != nil && _fbaf.Left == nil {
						_fbaf.Left = _cagf.Left
					}
				}
				_gbba.TblCellMar = _fbaf
				if _gbba.TblInd == nil {
					_gbba.TblInd = _feeg.TblInd
				}
				if _gbba.Jc == nil {
					_gbba.Jc = _feeg.Jc
				}
			}
			if _dagf.PPr != nil {
				_ebea = _bcda(_dagf.PPr, _ebea)
			}
			if _dagf.RPr != nil {
				_cddde = _gade(_dagf.RPr, _cddde)
			}
			if _dcbe := _dagf.BasedOn; _dcbe != nil {
				return _dccef(_dcde, _dcbe.ValAttr, _gbba, _ebea, _cddde)
			}
		}
	}
	return _gbba, _ebea, _cddde
}

func (_ddb *convertContext) drawPages() {
	for _, _gcfa := range _ddb._efceb {
		_ddb._fcgg.NewPage()
		_ddb.drawPage(_gcfa)
	}
}

func _gafdd(_agebe *_bc.CT_TblWidth, _gdgd, _faae float64) float64 {
	if _agebe != nil {
		if _gafdc := _agebe.WAttr; _gafdc != nil {
			if _cgc := _gafdc.ST_DecimalNumberOrPercent; _cgc != nil {
				if _fdgeg := _cgc.ST_UnqualifiedPercentage; _fdgeg != nil {
					switch _agebe.TypeAttr {
					case _bc.ST_TblWidthDxa:
						return float64(*_fdgeg) / 20
					case _bc.ST_TblWidthPct:
						return float64(*_fdgeg) / 100 * _gdgd
					default:
						return _faae
					}
				}
			}
		}
	}
	return _faae
}

type headerFooterRef struct {
	_eadd   bool
	_fcfcfa bool
	_fgcc   string
	_gdae   _bc.ST_HdrFtr
}

func _ebgb(_agbf int64, _dfbc _bc.ST_NumberFormat) string {
	_baec := int(_agbf)
	switch _dfbc {
	case _bc.ST_NumberFormatDecimal:
		return _b.Itoa(_baec)
	case _bc.ST_NumberFormatUpperRoman:
		return _ebddf(_baec, true)
	case _bc.ST_NumberFormatLowerRoman:
		return _ebddf(_baec, false)
	case _bc.ST_NumberFormatUpperLetter:
		return _eefad(_baec, true)
	case _bc.ST_NumberFormatLowerLetter:
		return _eefad(_baec, false)
	default:
		return _b.Itoa(_baec)
	}
}

func (_gfg *convertContext) newPage() {
	_ebag := &page{}
	_ebag._eeg = _gfg._eaebc
	_ebag._dbg = _ebag._eeg.Top
	if _gfg._dade {
		_ebag._de = true
		_ebag._dbg += _ga
	}
	_gfg._efceb = append(_gfg._efceb, _ebag)
	_gfg._fbcd = _ebag
}

func _gade(_efea, _fddaf *_bc.CT_RPr) *_bc.CT_RPr {
	if _efea == nil {
		return _fddaf
	}
	if _fddaf == nil {
		return _efea
	}
	if _efea.RStyle == nil {
		_efea.RStyle = _fddaf.RStyle
	}
	if _efea.RFonts == nil {
		_efea.RFonts = _fddaf.RFonts
	}
	if _efea.B == nil {
		_efea.B = _fddaf.B
	}
	if _efea.BCs == nil {
		_efea.BCs = _fddaf.BCs
	}
	if _efea.I == nil {
		_efea.I = _fddaf.I
	}
	if _efea.ICs == nil {
		_efea.ICs = _fddaf.ICs
	}
	if _efea.Caps == nil {
		_efea.Caps = _fddaf.Caps
	}
	if _efea.SmallCaps == nil {
		_efea.SmallCaps = _fddaf.SmallCaps
	}
	if _efea.Strike == nil {
		_efea.Strike = _fddaf.Strike
	}
	if _efea.Dstrike == nil {
		_efea.Dstrike = _fddaf.Dstrike
	}
	if _efea.Outline == nil {
		_efea.Outline = _fddaf.Outline
	}
	if _efea.Shadow == nil {
		_efea.Shadow = _fddaf.Shadow
	}
	if _efea.Emboss == nil {
		_efea.Emboss = _fddaf.Emboss
	}
	if _efea.Imprint == nil {
		_efea.Imprint = _fddaf.Imprint
	}
	if _efea.NoProof == nil {
		_efea.NoProof = _fddaf.NoProof
	}
	if _efea.SnapToGrid == nil {
		_efea.SnapToGrid = _fddaf.SnapToGrid
	}
	if _efea.Vanish == nil {
		_efea.Vanish = _fddaf.Vanish
	}
	if _efea.WebHidden == nil {
		_efea.WebHidden = _fddaf.WebHidden
	}
	if _efea.Color == nil {
		_efea.Color = _fddaf.Color
	}
	if _efea.Spacing == nil {
		_efea.Spacing = _fddaf.Spacing
	}
	if _efea.W == nil {
		_efea.W = _fddaf.W
	}
	if _efea.Kern == nil {
		_efea.Kern = _fddaf.Kern
	}
	if _efea.Position == nil {
		_efea.Position = _fddaf.Position
	}
	if _efea.Sz == nil {
		_efea.Sz = _fddaf.Sz
	}
	if _efea.SzCs == nil {
		_efea.SzCs = _fddaf.SzCs
	}
	if _efea.Highlight == nil {
		_efea.Highlight = _fddaf.Highlight
	}
	if _efea.U == nil {
		_efea.U = _fddaf.U
	}
	if _efea.Effect == nil {
		_efea.Effect = _fddaf.Effect
	}
	if _efea.Bdr == nil {
		_efea.Bdr = _fddaf.Bdr
	}
	if _efea.Shd == nil {
		_efea.Shd = _fddaf.Shd
	}
	if _efea.FitText == nil {
		_efea.FitText = _fddaf.FitText
	}
	if _efea.VertAlign == nil {
		_efea.VertAlign = _fddaf.VertAlign
	}
	if _efea.Rtl == nil {
		_efea.Rtl = _fddaf.Rtl
	}
	if _efea.Cs == nil {
		_efea.Cs = _fddaf.Cs
	}
	if _efea.Em == nil {
		_efea.Em = _fddaf.Em
	}
	if _efea.Lang == nil {
		_efea.Lang = _fddaf.Lang
	}
	if _efea.EastAsianLayout == nil {
		_efea.EastAsianLayout = _fddaf.EastAsianLayout
	}
	if _efea.SpecVanish == nil {
		_efea.SpecVanish = _fddaf.SpecVanish
	}
	if _efea.OMath == nil {
		_efea.OMath = _fddaf.OMath
	}
	if _efea.RPrChange == nil {
		_efea.RPrChange = _fddaf.RPrChange
	}
	return _efea
}

func _adbc(_dbfc *_bc.CT_TblPr, _befge *_bc.CT_TblPrEx, _cegc *_bc.CT_TcPr, _fbfe, _aabd, _aged, _dabbf int) *_bc.CT_TcPr {
	if _cegc == nil {
		_cegc = _bc.NewCT_TcPr()
	}
	if _befge == nil {
		_befge = _bc.NewCT_TblPrEx()
	}
	if _dbfc == nil {
		_dbfc = _bc.NewCT_TblPr()
	}
	if _cegc.TcBorders == nil {
		_cegc.TcBorders = _bc.NewCT_TcBorders()
	}
	if _befge.TblBorders == nil {
		_befge.TblBorders = _bc.NewCT_TblBorders()
	}
	if _dbfc.TblBorders == nil {
		_dbfc.TblBorders = _bc.NewCT_TblBorders()
	}
	if _cegc.TcBorders.Top == nil {
		if _befge.TblBorders.Top == nil {
			_cegc.TcBorders.Top = _befb(_dbfc.TblBorders.Top, _dbfc.TblBorders.InsideH, _fbfe == 0)
		} else {
			_cegc.TcBorders.Top = _befb(_befge.TblBorders.Top, _befge.TblBorders.InsideH, _fbfe == 0)
		}
	}
	if _cegc.TcBorders.Bottom == nil {
		if _befge.TblBorders.Bottom == nil {
			_cegc.TcBorders.Bottom = _befb(_dbfc.TblBorders.Bottom, _dbfc.TblBorders.InsideH, _fbfe == _aged-1)
		} else {
			_cegc.TcBorders.Bottom = _befb(_befge.TblBorders.Bottom, _befge.TblBorders.InsideH, _fbfe == _aged-1)
		}
	}
	if _cegc.TcBorders.Left == nil {
		if _befge.TblBorders.Left == nil {
			_cegc.TcBorders.Left = _befb(_dbfc.TblBorders.Left, _dbfc.TblBorders.InsideV, _aabd == 0)
		} else {
			_cegc.TcBorders.Left = _befb(_befge.TblBorders.Left, _befge.TblBorders.InsideV, _aabd == 0)
		}
	}
	if _cegc.TcBorders.Right == nil {
		if _befge.TblBorders.Right == nil {
			_cegc.TcBorders.Right = _befb(_dbfc.TblBorders.Right, _dbfc.TblBorders.InsideV, _aabd == _dabbf-1)
		} else {
			_cegc.TcBorders.Right = _befb(_befge.TblBorders.Right, _befge.TblBorders.InsideV, _aabd == _dabbf-1)
		}
	}
	if _cegc.Shd == nil {
		if _gefa := _dbfc.Shd; _gefa != nil {
			_cegc.Shd = _gefa
		}
	} else {
		if _dbfc.Shd != nil && _cegc.Shd.FillAttr == nil {
			_cegc.Shd.FillAttr = _dbfc.Shd.FillAttr
		}
	}
	if _cegc.TcMar == nil {
		if _bdac := _dbfc.TblCellMar; _bdac != nil {
			_cegc.TcMar = _bc.NewCT_TcMar()
			_cegc.TcMar.Left = _bdac.Left
		}
	} else {
		if _dbfc.TblCellMar != nil && _cegc.TcMar.Left == nil {
			_cegc.TcMar.Left = _dbfc.TblCellMar.Left
		}
	}
	return _cegc
}

type prefix struct {
	_ebab string
	_aceb []float64
	_eaee bool
	_aedd bool
}

func _febde(_ecagc string) uint16 {
	_bgac, _ccfa := _gefg[_ecagc]
	if !_ccfa {
		return 0
	}
	return _bgac
}

func _egbd(_acgfb string) bool {
	for _, _afec := range _acgfb {
		if _afec > 255 {
			return false
		}
	}
	return true
}

type romanMatch struct {
	_bcacd int
	_fbgb  string
}

func (_fcad *convertContext) makeRunStyle(_egefc *_bc.CT_RPr, _dbac, _eebaf, _ggcb, _bcabf, _abce bool) (_ea.TextStyle, bool, bool, *_ea.Color) {
	var _edcc *_ea.Color
	_cafe := _fcad._fcgg.NewTextStyle()
	if _egefc != nil {
		_afdf := _a.FontStyle_Regular
		_faaea := _dgfgag(_egefc.B)
		_ccee := _dgfgag(_egefc.I)
		if _faaea && _ccee {
			_afdf = _a.FontStyle_BoldItalic
		} else if _faaea {
			_afdf = _a.FontStyle_Bold
		} else if _ccee {
			_afdf = _a.FontStyle_Italic
		}
		if _bcabf {
			_cafe.Font = _a.AssignStdFontByName(_cafe, "\u0053\u0079\u006d\u0062\u006f\u006c")
		} else {
			_ebed := "\u0064e\u0066\u0061\u0075\u006c\u0074"
			if _edde := _egefc.RFonts; _edde != nil {
				if _gbafe := _edde.AsciiAttr; _gbafe != nil {
					_ebed = *_gbafe
				} else if _ffdf := _edde.HAnsiAttr; _ffdf != nil {
					_ebed = *_ffdf
				} else if _gccb := _edde.CsAttr; _gccb != nil {
					_ebed = *_gccb
				} else {
					_eab := _fcad._fbef
					if _eab != nil {
						_fegbf := ""
						if _feec := _eab.RFonts; _feec != nil {
							if _fded := _edde.HintAttr; _fded == _bc.ST_HintEastAsia {
								if _feec.EastAsiaAttr != nil {
									_ebed = *_eab.RFonts.EastAsiaAttr
								} else {
									if _feec.EastAsiaThemeAttr == _bc.ST_ThemeMajorEastAsia {
										_fegbf = _dgca
									}
									if _feec.EastAsiaThemeAttr == _bc.ST_ThemeMinorEastAsia {
										_fegbf = _acag
									}
								}
							} else {
								if _bbaa := _feec.AsciiAttr; _bbaa != nil {
									_ebed = *_bbaa
								} else if _caad := _feec.HAnsiAttr; _caad != nil {
									_ebed = *_caad
								}
							}
							if _ebed == "\u0064e\u0066\u0061\u0075\u006c\u0074" {
								if _fegbf == "" {
									if _feec.EastAsiaThemeAttr == _bc.ST_ThemeMajorEastAsia {
										_fegbf = _dgca
									} else if _feec.EastAsiaThemeAttr == _bc.ST_ThemeMinorEastAsia {
										_fegbf = _acag
									} else if _cgab := _feec.AsciiThemeAttr; _cgab == _bc.ST_ThemeMajorAscii || _cgab == _bc.ST_ThemeMajorHAnsi {
										_fegbf = _gfef
									} else if _dacea := _feec.AsciiThemeAttr; _dacea == _bc.ST_ThemeMinorAscii || _dacea == _bc.ST_ThemeMinorHAnsi {
										_fegbf = _daga
									} else {
										if _bcdg := _feec.HAnsiThemeAttr; _bcdg == _bc.ST_ThemeMajorAscii || _bcdg == _bc.ST_ThemeMajorHAnsi {
											_fegbf = _gfef
										} else if _fedgc := _feec.HAnsiThemeAttr; _fedgc == _bc.ST_ThemeMinorAscii || _fedgc == _bc.ST_ThemeMinorHAnsi {
											_fegbf = _daga
										}
									}
								}
								_gaed := ""
								if _ebed == "\u0064e\u0066\u0061\u0075\u006c\u0074" {
									if _fbdee := _fcad._cffa.Settings.X(); _fbdee != nil {
										_fgada := ""
										if _efdd := _fbdee.ThemeFontLang; _efdd != nil {
											if _efdd.ValAttr != nil {
												_fgada = *_efdd.ValAttr
											}
											if _efdd.EastAsiaAttr != nil {
												_fgada = *_efdd.EastAsiaAttr
											}
											if _efdd.BidiAttr != nil {
												_fgada = *_efdd.BidiAttr
											}
										}
										_gaed = _fccbba(_febde(_fgada))
									}
								}
								_becb := _fcad._cffa.Themes()
								for _, _dacee := range _becb {
									if _dacee.ThemeElements != nil {
										if _deef := _dacee.ThemeElements.FontScheme; _deef != nil {
											if _deef.MajorFont != nil && _fegbf == _dgca {
												if _deef.MajorFont.Ea != nil {
													_ebed = _deef.MajorFont.Ea.TypefaceAttr
													if _ebed == "" && _gaed != "" {
														for _, _cbfd := range _deef.MajorFont.Font {
															if _cbfd.ScriptAttr == _gaed {
																_ebed = _cbfd.TypefaceAttr
																break
															}
														}
													}
													break
												}
											} else if _deef.MinorFont != nil && _fegbf == _acag {
												if _deef.MinorFont.Ea != nil {
													_ebed = _deef.MinorFont.Ea.TypefaceAttr
													if _ebed == "" && _gaed != "" {
														for _, _dbef := range _deef.MinorFont.Font {
															if _dbef.ScriptAttr == _gaed {
																_ebed = _dbef.TypefaceAttr
																break
															}
														}
													}
													break
												}
											} else if _deef.MajorFont != nil && _fegbf == _gfef {
												if _deef.MajorFont.Latin != nil {
													_ebed = _deef.MajorFont.Latin.TypefaceAttr
													break
												}
											} else if _deef.MinorFont != nil && _fegbf == _daga {
												if _deef.MinorFont.Latin != nil {
													_ebed = _deef.MinorFont.Latin.TypefaceAttr
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
			}
			if _ebed != "\u0064e\u0066\u0061\u0075\u006c\u0074" && !_egbd(_ebed) {
				if _eaaa := _fcad._cffa.FontTable(); _eaaa != nil {
					for _, _aafb := range _eaaa.Font {
						if _aafb.NameAttr == _ebed && _aafb.AltName != nil && _egbd(_aafb.AltName.ValAttr) {
							_ebed = _aafb.AltName.ValAttr
							break
						}
						if _aafb.AltName != nil && !_egbd(_aafb.AltName.ValAttr) && _aafb.AltName.ValAttr == _ebed {
							_ebed = _aafb.NameAttr
							break
						}
					}
				}
			}
			if _bgbec, _agaf := _a.StdFontsMap[_ebed]; _agaf {
				_cafe.Font = _a.AssignStdFontByName(_cafe, _bgbec[_afdf])
			} else if _eeed := _a.GetRegisteredFont(_ebed, _afdf); _eeed != nil {
				_cafe.Font = _eeed
			} else {
				_da.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0077\u0069\u0074h\u0020\u0073\u0074\u0079\u006c\u0065\u0020\u0025s\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0072\u0065\u0073\u0065\u0074 \u0074\u006f\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u002e", _ebed, _afdf)
				_cafe.Font = _a.AssignStdFontByName(_cafe, _a.StdFontsMap["\u0064e\u0066\u0061\u0075\u006c\u0074"][_afdf])
			}
		}
		_daef := _gdeg(_egefc.Sz, _egefc.SzCs)
		if _ecag := _egefc.VertAlign; _ecag != nil {
			_ecddb := _ecag.ValAttr
			_dbac = _ecddb == _ee.ST_VerticalAlignRunSuperscript
			_eebaf = _ecddb == _ee.ST_VerticalAlignRunSubscript
		}
		if _daef > _fcad._fcfcf {
			_fcad._fcfcf = _daef
		}
		if _dbac || _eebaf {
			_daef *= 0.64
		}
		if _ggcb {
			if _dbac {
				_cafe.TextRise = 1.5
			}
			if _eebaf {
				_cafe.TextRise = -1.5
			}
		}
		_cafe.FontSize = _daef
		_bbgb := 0.0
		if _aebf := _egefc.Spacing; _aebf != nil {
			_bbgb = _a.PointsFromTwips(*_aebf.ValAttr.Int64)
			if _bbgb < 0.0 {
				_bbgb = 0.0
			}
		}
		_cafe.CharSpacing = _bbgb
		_cfce := 0.0
		if _dgbe := _egefc.Position; _dgbe != nil {
			_cfce = float64(*_dgbe.ValAttr.Int64) / 24 * _daef
		}
		_cafe.TextRise = _cfce
		_dgcc := _ea.ColorBlack
		if _egefc.Color != nil {
			_fddbc := _egefc.Color.ValAttr.ST_HexColorRGB
			if _fddbc != nil {
				_dgcc = _ea.ColorRGBFromHex("\u0023" + *_fddbc)
			}
		}
		if _ggcb {
			_gbfe, _bfaff, _gfca := _dgcc.ToRGB()
			_gbfe, _bfaff, _gfca = _a.Lighten(_gbfe), _a.Lighten(_bfaff), _a.Lighten(_gfca)
			_dgcc = _ea.ColorRGBFromArithmetic(_gbfe, _bfaff, _gfca)
		}
		_cafe.Color = _dgcc
		if _abce {
			_edcc = &_dgcc
		}
		if _egefc.U != nil && _egefc.U.ValAttr != _bc.ST_UnderlineNone && !_bcabf {
			_edcc = &_dgcc
			if _fddad := _egefc.U.ColorAttr; _fddad != nil {
				if _cabe := _fddad.ST_HexColorRGB; _cabe != nil {
					_afca := _ea.ColorRGBFromHex("\u0023" + *_cabe)
					_edcc = &_afca
				}
			}
		}
	}
	return _cafe, _dbac, _eebaf, _edcc
}

var _agddb, _abac, _bgdc *_bf.Regexp

// RegisterFontsFromDirectory registers all fonts from the given directory automatically detecting font families and styles. For composite fonts use RegisterCompositeFontsFromDirectory.
func RegisterFontsFromDirectory(dirName string) error {
	return _a.RegisterFontsFromDirectory(dirName)
}

func _gdeg(_dbce, _cfa *_bc.CT_HpsMeasure) float64 {
	var _aabb float64
	_gafe := _a.DefaultFontSize
	if _dbce != nil {
		_aabb = float64(*_dbce.ValAttr.ST_UnsignedDecimalNumber)
	} else if _cfa != nil {
		_aabb = float64(*_cfa.ValAttr.ST_UnsignedDecimalNumber)
	}
	if _aabb != 0 {
		_gafe = _aabb / 24 * _a.DefaultFontSize
	}
	return _gafe
}

func (_gadd *convertContext) makeBlockFromTextboxContent(_agac *_bc.TxbxContent, _fggg, _cdgf float64, _fbde *_a.Rectangle) (*block, error) {
	if _fbde == nil {
		_fbde = &_a.Rectangle{}
	}
	if _ebce := _agac.EG_ContentBlockContent; len(_ebce) > 0 {
		_aegb, _gcg := _gadd.makePdfBlockFromCBCs([][]*_bc.EG_ContentBlockContent{_ebce}, _fggg, _cdgf, _fbde, false, nil)
		if _gcg != nil {
			return nil, _gcg
		}
		_cbgc := &block{_egb: _aegb, _dgg: false, _ebg: 0, _dc: _ea.ColorBlack}
		return _cbgc, nil
	}
	return nil, nil
}

func (_eebdf *convertContext) addCurrentParagraphHeaderToCurrentPage() {
	_eebdf.alignParagraph()
	_eebdf._fbcd._fga = append(_eebdf._fbcd._fga, _eebdf._cecg)
}

type page struct {
	_eeg *_a.Rectangle
	_eb  []*paragraph
	_dbg float64
	_gf  []*zoneToSkip
	_fc  []*image
	_gec []*image
	_ceb []*block
	_ab  []*block
	_eaf []*note
	_de  bool
	_eba []*headerFooterRef
	_fga []*paragraph
	_cb  []*paragraph
}

func (_dfaf *convertContext) addAbsoluteHeaderFooterCBCs(_abafa []*_bc.EG_ContentBlockContent) {
	for _, _fcggb := range _abafa {
		for _, _adfb := range _fcggb.P {
			_dfaf.newParagraph()
			if _adfb.PPr != nil && _adfb.PPr.PStyle == nil {
				_gbfd := _dfaf._cffa.Styles.ParagraphStyles()
				for _, _faff := range _gbfd {
					if _effd := _faff.X().DefaultAttr; _effd != nil {
						if _eaed := _effd.Bool; _eaed != nil && *_eaed {
							_adfb.PPr = _cceb(_adfb.PPr, _faff.X().PPr, _faff.X().RPr)
						}
						if _abcg := _effd.ST_OnOff1; _abcg == _ee.ST_OnOff1On {
							_adfb.PPr = _cceb(_adfb.PPr, _faff.X().PPr, _faff.X().RPr)
						}
						break
					}
				}
			}
			_cgdb, _bgbc := _dfaf.combinePPrWithStyles(_adfb.PPr)
			if _bgbc != nil {
				_dfaf._eegda = _bgbc
			}
			_dfaf.assignPropsToAbsoluteParagraph(_cgdb, _dfaf._cecg)
			_dfaf.determineParagraphBounds()
			_dfaf.newLine()
			_dfaf.newWord()
			_dfdg := _adfb.EG_PContent
			if len(_dfdg) == 0 {
				_dfaf.addEmptyLine()
			} else {
				_dfaf.addAnchorBlocks(_dfdg)
				_dfaf.addAnchorExtra(_dfdg)
				_dfaf.addAbsoluteEGPC(_dfdg, _cgdb)
				_dfaf.addCurrentWordToParagraph()
			}
			if _dfaf._cbae {
				_dfaf.addCurrentParagraphHeaderToCurrentPage()
			} else {
				_dfaf.addCurrentParagraphFooterToCurrentPage()
			}
		}
		for _, _fdfd := range _fcggb.Tbl {
			if _dfaf._cecg == nil {
				_dfaf.newParagraph()
				_dfaf.determineParagraphBounds()
				_dfaf.newLine()
				_dfaf.newWord()
			}
			_dfaf.addAbsoluteHeaderFooterTable(_fdfd)
		}
	}
}

func (_affa *convertContext) makePdfBlockFromChart(_fefd *_gg.Chart, _eeege, _ceaf float64) (*_ea.Block, error) {
	_aed := _fefd.CT_RelId.IdAttr
	_eaeba := _affa._cffa.GetChartSpaceByRelId(_aed)
	if _eaeba == nil {
		return nil, _ce.New("\u004e\u006f\u0020\u0063\u0068\u0061\u0072\u0074\u0073\u0070\u0061\u0063\u0065")
	}
	var _ceef *_bb.Theme
	_aced := _affa._cffa.Themes()
	if len(_aced) > 0 {
		_ceef = _aced[0]
	}
	return _a.MakeBlockFromChartSpace(_eaeba, _eeege, _ceaf, _ceef)
}

func _eefad(_gadc int, _bdgg bool) string {
	_ccfg := (_gadc-1)/26 + 1
	_dbda := byte((_gadc - 1) % 26)
	if _bdgg {
		_dbda += byte(65)
	} else {
		_dbda += byte(97)
	}
	_bdec := _d.NewBuffer([]byte{})
	for _ffab := 0; _ffab < _ccfg; _ffab++ {
		_bdec.Write([]byte{_dbda})
	}
	return _bdec.String()
}

func (_dgfb *convertContext) newWord() { _dgfb._gced = &word{_gfb: true, _cfe: _dgfb._bbfa._dbb} }

func (_aafa *convertContext) newParagraph() {
	if _aafa._fbcd == nil {
		_aafa.newPage()
	}
	_ggcge := &paragraph{}
	_ggcge._ac = &_a.Rectangle{}
	_ggcge._bgb = _aafa._fbcd._dbg
	_aafa._cecg = _ggcge
}

type block struct {
	_egb *_ea.Block
	_aef float64
	_aga float64
	_dgg bool
	_ebg float64
	_dc  _ea.Color
}

func (_ege *convertContext) addAbsoluteTable(_gbde *_bc.CT_Tbl) {
	_dcfg := _gbde.TblGrid
	if _dcfg == nil {
		return
	}
	_gebf := len(_dcfg.GridCol)
	if _gebf == 0 {
		_da.Log.Debug("\u004d\u0069\u0073\u0073\u0069\u006e\u0067\u0020\u0067\u0072\u0069\u0064\u0043\u006f\u006c\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u002c\u0020\u0063r\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u0067\u0072\u0069\u0064C\u006f\u006c\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u002e")
		_cffb := _gbde.EG_ContentRowContent[0]
		if len(_cffb.Tr) < 1 {
			return
		}
		_fdfe := _cffb.Tr[0]
		if len(_fdfe.EG_ContentCellContent) < 1 {
			return
		}
		_fddag := 0
		if _gbde.TblPr != nil && _gbde.TblPr.TblW != nil {
			if _bcgf := _gbde.TblPr.TblW.WAttr; _bcgf != nil {
				switch _gbde.TblPr.TblW.TypeAttr {
				case _bc.ST_TblWidthPct, _bc.ST_TblWidthDxa:
					if _bcgf.ST_DecimalNumberOrPercent != nil {
						if _bcgf.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage != nil {
							_fddag = int(*_bcgf.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage)
						}
					}
				}
			}
		}
		var _bgdb []*_bc.CT_TblGridCol
		for _, _fffa := range _fdfe.EG_ContentCellContent {
			if _gebe := _fffa.Tc; len(_gebe) > 0 {
				if _efeb := _gebe[0]; _efeb != nil {
					_aeea := _bc.NewCT_TblGridCol()
					if _efeb.TcPr != nil {
						if _dbdc := _efeb.TcPr.TcW; _dbdc != nil {
							if _dbdc.WAttr != nil {
								if _cdaf := _dbdc.WAttr.ST_DecimalNumberOrPercent; _cdaf != nil {
									if _acge := _cdaf.ST_UnqualifiedPercentage; _acge != nil {
										_edc := uint64(*_acge)
										_bdfbd := &_ee.ST_TwipsMeasure{}
										_bdfbd.ST_UnsignedDecimalNumber = &_edc
										_aeea.WAttr = _bdfbd
									}
								}
							}
						}
						_bgdb = append(_bgdb, _aeea)
						if _efeb.TcPr.GridSpan != nil {
							for _bag := int(_efeb.TcPr.GridSpan.ValAttr) - 1; _bag > 0; _bag-- {
								_dfd := _bc.NewCT_TblGridCol()
								_bgdb = append(_bgdb, _dfd)
							}
						}
					} else {
						_bgdb = append(_bgdb, _aeea)
					}
				}
			}
		}
		_egbc := uint64(_fddag / len(_bgdb))
		for _, _abcf := range _bgdb {
			if _abcf.WAttr == nil {
				_eaa := &_ee.ST_TwipsMeasure{}
				_eaa.ST_UnsignedDecimalNumber = &_egbc
				_abcf.WAttr = _eaa
			}
		}
		_dcfg.GridCol = _bgdb
		_gebf = len(_bgdb)
	}
	_edgb := []float64{}
	_afce := []float64{}
	_dgcf := 0.0
	for _, _ffbe := range _dcfg.GridCol {
		_eefb := 0.0
		if _ffbe.WAttr.ST_UnsignedDecimalNumber != nil {
			_eefb = _a.PointsFromTwips(int64(*_ffbe.WAttr.ST_UnsignedDecimalNumber))
		}
		_edgb = append(_edgb, _eefb)
		_dgcf += _eefb
	}
	for _acga := 0; _acga < _gebf; _acga++ {
		_afce = append(_afce, _edgb[_acga]/_dgcf)
	}
	_cfb := _ege._fcgg.NewTable(_gebf)
	_cfb.SetColumnWidths(_afce...)
	_bffg := _ege._fcgg.NewTable(_gebf)
	_bffg.SetColumnWidths(_afce...)
	_dgdc, _ccea, _acgaf := _dcad(_ege._cffa, _gbde.TblPr)
	var _efcb []*_bc.CT_TblStylePr
	if _dgdc.TblStyle != nil {
		_efcb = _bcffb(_ege._cffa, _dgdc.TblStyle.ValAttr)
	}
	_bdfd := _gafdd(_dgdc.TblW, _ege._fbcd._eeg.Right-_ege._fbcd._eeg.Left, 0)
	_cged := _gafdd(_dgdc.TblInd, _ege._fbcd._eeg.Right-_ege._fbcd._eeg.Left, 0)
	_gdcb := _ege._fbcd._eeg.Bottom - _ege._cecg._bgb
	_beb := len(_gbde.EG_ContentRowContent)
	for _fgdd, _cea := range _gbde.EG_ContentRowContent {
		if _cea == nil {
			continue
		}
		_gafc := _ege._fcgg.NewTable(_gebf)
		_gafc.SetColumnWidths(_afce...)
		if _gdabd := _cea.Tr; len(_gdabd) > 0 {
			_ecfd := _gdabd[0]
			_fafb := _ecfd.TblPrEx
			var _bgfg float64
			if _acffb := _ecfd.TrPr; _acffb != nil {
				if len(_acffb.TrHeight) != 0 {
					_cgdg := _acffb.TrHeight[0]
					if _gdba := _cgdg.ValAttr; _gdba != nil {
						if _gdba.ST_UnsignedDecimalNumber != nil {
							_bgfg = _a.PointsFromTwips(int64(*_gdba.ST_UnsignedDecimalNumber))
						}
					}
				}
			}
			if _bgfg < _gafc.Height() {
				_bgfg = _gafc.Height()
			}
			if _bgfg < _ffcdc(4) {
				_bgfg = _ffcdc(4)
			}
			_bffg.SetRowHeight(_bffg.CurRow(), _bgfg)
			_gafc.SetRowHeight(_gafc.CurRow(), _bgfg)
			if _bdfd == 0 || _bdfd > _ege._fbcd._eeg.Right-_ege._fbcd._eeg.Left {
				_bdfd = _ege._fbcd._eeg.Right - _ege._fbcd._eeg.Left
			}
			for _fgefa, _fdcf := range _ecfd.EG_ContentCellContent {
				if _dcgd := _fdcf.Tc; len(_dcgd) > 0 {
					if _fdcb := _dcgd[0]; _fdcb != nil {
						_ade := _ege.addCellToTable(_bffg, _fdcb, _dgdc, _fafb, _fgdd, _fgefa, _beb, _gebf, _efcb, _ccea, _acgaf, false)
						_ege.addCellToTable(_gafc, _fdcb, _dgdc, _fafb, _fgdd, _fgefa, _beb, _gebf, _efcb, _ccea, _acgaf, false)
						if _ade {
							_gafdf := _a.MakeTempCreator(_bdfd, _ffcdc(1000))
							_gafdf.Draw(_bffg)
							*_bffg = *_gafc
							_bffg.SetRowHeight(_bffg.CurRow(), _bgfg)
							if _cfb != nil {
								_ege.addParagraphWithTable(*_cfb, _bdfd, _cged)
								_ege.newPage()
								_cfb = nil
							}
							_bffg = _ege._fcgg.NewTable(_gebf)
							_bffg.SetColumnWidths(_afce...)
							_gafc = _ege._fcgg.NewTable(_gebf)
							_gafc.SetColumnWidths(_afce...)
							_ege.addCellToTable(_bffg, _fdcb, _dgdc, _fafb, _fgdd, _fgefa, _beb, _gebf, _efcb, _ccea, _acgaf, true)
							_ege.addCellToTable(_gafc, _fdcb, _dgdc, _fafb, _fgdd, _fgefa, _beb, _gebf, _efcb, _ccea, _acgaf, true)
							continue
						}
					}
				}
			}
			_bbb := _a.MakeTempCreator(_bdfd, _ffcdc(1000))
			_bbb.Draw(_bffg)
			if _bffg.Height() >= _gdcb && _cfb != nil {
				_ege.addParagraphWithTable(*_cfb, _bdfd, _cged)
				_ege.newPage()
				*_bffg = *_gafc
				_bffg.SetRowHeight(_bffg.CurRow(), _bgfg)
				_gdcb = _ege._fbcd._eeg.Bottom - _ege._fbcd._eeg.Top
				_cfb = nil
			} else {
				if _cfb == nil {
					_cfb = _ege._fcgg.NewTable(_gebf)
					_cfb.SetColumnWidths(_afce...)
				}
				*_cfb = *_bffg
			}
		}
	}
	if _cfb != nil {
		_ege.addParagraphWithTable(*_cfb, _bdfd, _cged)
	}
}

type borderLine struct {
	_afaf _ea.Color
	_ecc  _a.BorderPosition
	_dca  float64
	_ed   float64
	_adc  float64
}

func _fccbba(_fegcg uint16) string {
	switch _fegcg {
	case 0x429, 0x401, 0x801, 0xc01, 0x1001, 0x1401, 0x1801, 0x1c01, 0x2001, 0x2401, 0x2801, 0x2c01, 0x3001, 0x3401, 0x3801, 0x3c01, 0x4001, 0x420, 0x846, 0x859, 0x45f, 0x460, 0x463, 0x48c:
		return "\u0041\u0072\u0061\u0062"
	case 0x42b:
		return "\u0041\u0072\u006d\u006e"
	case 0x445, 0x845, 0x44d, 0x458:
		return "\u0042\u0065\u006e\u0067"
	case 0x45d:
		return "\u0043\u0061\u006e\u0073"
	case 0x45c:
		return "\u0043\u0068\u0065\u0072"
	case 0x419, 0x402, 0x281a, 0x422, 0x819, 0xc1a, 0x1c1a, 0x201a, 0x301a, 0x423, 0x428, 0x82c, 0x42f, 0x43f, 0x440, 0x843, 0x444, 0x450, 0x46d, 0x485:
		return "\u0043\u0072\u0079\u006c"
	case 0x439, 0x44e, 0x44f, 0x457, 0x459, 0x860, 0x461, 0x861:
		return "\u0044\u0065\u0076\u0061"
	case 0x45e, 0x473, 0x873:
		return "\u0045\u0074\u0068\u0069"
	case 0x437:
		return "\u0047\u0065\u006f\u0072"
	case 0x408:
		return "\u0047\u0072\u0065\u006b"
	case 0x447:
		return "\u0047\u0075\u006a\u0072"
	case 0x446:
		return "\u0047\u0075\u0072\u0075"
	case 0x412:
		return "\u0048\u0061\u006e\u0067"
	case 0x804, 0x1004:
		return "\u0048\u0061\u006e\u0073"
	case 0x404, 0xc04, 0x1404:
		return "\u0048\u0061\u006e\u0074"
	case 0x40d, 0x43d:
		return "\u0048\u0065\u0062\u0072"
	case 0x411:
		return "\u004a\u0070\u0061\u006e"
	case 0x453:
		return "\u004b\u0068\u006d\u0072"
	case 0x44b:
		return "\u004b\u006e\u0064\u0061"
	case 0x454:
		return "\u004c\u0061\u006f\u006f"
	case 0x409, 0xc09, 0x809, 0x1009, 0x403, 0x406, 0x413, 0x813, 0x479, 0x40b, 0x40c, 0xc0c, 0x407, 0x807, 0xc07, 0x1007, 0x1407, 0x410, 0x414, 0x814, 0x416, 0x816, 0x40a, 0x41d, 0x405, 0x40e, 0x415, 0x41f, 0x42d, 0x424, 0x426, 0x427, 0x418, 0x818, 0x241a, 0x41a, 0x491, 0x83c, 0x430, 0x431, 0x432, 0x433, 0x434, 0x435, 0x436, 0x425, 0x456, 0x41b, 0x1409, 0x1809, 0x1c09, 0x2009, 0x2409, 0x2809, 0x2c09, 0x3009, 0x3409, 0x3809, 0x3c09, 0x4009, 0x4409, 0x4809, 0x80a, 0xc0a, 0x100a, 0x140a, 0x180a, 0x1c0a, 0x200a, 0x240a, 0x280a, 0x2c0a, 0x300a, 0x340a, 0x380a, 0x3c0a, 0x400a, 0x440a, 0x480a, 0x4c0a, 0x500a, 0x540a, 0x80c, 0x100c, 0x140c, 0x180c, 0x1c0c, 0x200c, 0x240c, 0x280c, 0x2c0c, 0x300c, 0x340c, 0x3c0c, 0x380c, 0x40f, 0x810, 0x417, 0x81a, 0x101a, 0x141a, 0x181a, 0x2c1a, 0x41c, 0x81d, 0x421, 0x42c, 0x42e, 0x82e, 0x438, 0x43a, 0x43b, 0x83b, 0xc3b, 0x103b, 0x143b, 0x183b, 0x1c3b, 0x203b, 0x243b, 0x43e, 0x83e, 0x441, 0x442, 0x443, 0x452, 0x85d, 0x85f, 0x462, 0x464, 0x466, 0x467, 0x468, 0x469, 0x46a, 0x46b, 0x86b, 0xc6b, 0x46c, 0x46e, 0x46f, 0x470, 0x471, 0x472, 0x474, 0x475, 0x476, 0x477, 0x47a, 0x47c, 0x47e, 0x481, 0x482, 0x483, 0x484, 0x486, 0x487, 0x488:
		return "\u004c\u0061\u0074\u006e"
	case 0x44c:
		return "\u004d\u006c\u0079\u006d"
	case 0x850:
		return "\u004d\u006f\u006e\u0067"
	case 0x455:
		return "\u004d\u0079\u006d\u0072"
	case 0x448:
		return "\u004f\u0072\u0079\u0061"
	case 0x45b:
		return "\u0053\u0069\u006e\u0068"
	case 0x45a:
		return "\u0053\u0079\u0072\u0063"
	case 0x449:
		return "\u0054\u0061\u006d\u006c"
	case 0x44a:
		return "\u0054\u0065\u006c\u0075"
	case 0x465:
		return "\u0054\u0068\u0061\u0061"
	case 0x41e:
		return "\u0054\u0068\u0061\u0069"
	case 0x851, 0x451:
		return "\u0054\u0069\u0062\u0074"
	case 0x480:
		return "\u0055\u0069\u0067\u0068"
	case 0x42a:
		return "\u0056\u0069\u0065\u0074"
	case 0x478:
		return "\u0059\u0069\u0069\u0069"
	}
	return ""
}

func (_geaf *convertContext) newLine() {
	if _geaf._cecg == nil {
		_geaf.newParagraph()
	}
	_cbgdd := _geaf._cecg._egd + _geaf._cecg._ac.Top
	_cbfa := &line{}
	if len(_geaf._cecg._cc) == 0 {
		_cbfa._bfd = _geaf._cecg._dg
	} else {
		_cbfa._bfd = _geaf._cecg._eg
	}
	_cbfa._fdc = _geaf._cecg._dee
	_cbfa._dbb = _cbfa._bfd
	_cbfa._ag = _cbgdd
	_geaf._cecg._cc = append(_geaf._cecg._cc, _cbfa)
	_geaf._bbfa = _cbfa
	_geaf.newSpan()
}

func _dbafe(_fgfe *_bc.CT_Border) (_ea.CellBorderStyle, *_ea.Color, float64) {
	if _fgfe == nil {
		return _ea.CellBorderStyleNone, nil, 0
	}
	var _gbag _ea.CellBorderStyle
	switch _fgfe.ValAttr {
	case _bc.ST_BorderSingle:
		_gbag = _ea.CellBorderStyleSingle
	case _bc.ST_BorderDouble:
		_gbag = _ea.CellBorderStyleDouble
	default:
		_gbag = _ea.CellBorderStyleNone
	}
	var _gceda _ea.Color
	if _dagec := _fgfe.ColorAttr; _dagec != nil {
		if _daea := _dagec.ST_HexColorRGB; _daea != nil {
			_gceda = _ea.ColorRGBFromHex("\u0023" + *_daea)
		}
	}
	_cdbe := 0.0
	if _fcdg := _fgfe.SzAttr; _fcdg != nil {
		_cdbe = float64(*_fcdg) / 8
	}
	return _gbag, &_gceda, _cdbe
}

func (_cade *convertContext) assignHeaderFooterToPage(_daad *page) {
	_bbgg := _cade._cffa.DocRels()
	_baaa := _bf.MustCompile("\u005b\u0030\u002d\u0039\u005d\u002b")
	if len(_daad._eba) > 0 {
		for _, _dccb := range _daad._eba {
			if _dccb != nil {
				if _dccb._eadd {
					if _abg := _bbgg.GetTargetByRelId(_dccb._fgcc); _abg != "" {
						_eebf, _ := _b.Atoi(_baaa.FindString(_abg))
						for _cbba, _eebag := range _cade._cffa.Headers() {
							if _eebf == (_cbba + 1) {
								_cade._cbae = true
								_cade._bcdd = false
								_cade.addAbsoluteHeaderFooterCBCs(_eebag.X().EG_ContentBlockContent)
							}
						}
					}
				}
				if _dccb._fcfcfa {
					if _bedf := _bbgg.GetTargetByRelId(_dccb._fgcc); _bedf != "" {
						_accb, _ := _b.Atoi(_baaa.FindString(_bedf))
						for _bacc, _bdag := range _cade._cffa.Footers() {
							if _accb == (_bacc + 1) {
								_cade._cbae = false
								_cade._bcdd = true
								_cade.addAbsoluteHeaderFooterCBCs(_bdag.X().EG_ContentBlockContent)
							}
						}
					}
				}
			}
		}
	}
}

type styleAttributes struct{}

func _fgef(_fgf, _bfc string, _gbc, _eff, _gbf bool) []*symbol {
	_decd := []*symbol{}
	for _, _aeg := range _fgf {
		_bfed := &symbol{_fge: string(_aeg), _ffa: _gbc, _bcb: _eff, _eae: _bfc, _fbb: _gbf}
		_decd = append(_decd, _bfed)
	}
	return _decd
}

func (_agfe *convertContext) adjustHeights(_ccg float64) {
	if _agfe._bbfa._bca < _ccg {
		_agfe._cecg._egd += (_ccg - _agfe._bbfa._bca)
		_agfe._bbfa._bca = _ccg
	}
}

type line struct {
	_ag  float64
	_bfd float64
	_fdc float64
	_dbb float64
	_bca float64
	_dag []*span
}

func _bcda(_fcdf, _bfdg *_bc.CT_PPrGeneral) *_bc.CT_PPrGeneral {
	if _fcdf == nil {
		return _fcdf
	}
	if _bfdg == nil {
		return _fcdf
	}
	if _fcdf.PStyle == nil {
		_fcdf.PStyle = _bfdg.PStyle
	}
	if _fcdf.KeepNext == nil {
		_fcdf.KeepNext = _bfdg.KeepNext
	}
	if _fcdf.KeepLines == nil {
		_fcdf.KeepLines = _bfdg.KeepLines
	}
	if _fcdf.PageBreakBefore == nil {
		_fcdf.PageBreakBefore = _bfdg.PageBreakBefore
	}
	if _fcdf.FramePr == nil {
		_fcdf.FramePr = _bfdg.FramePr
	}
	if _fcdf.WidowControl == nil {
		_fcdf.WidowControl = _bfdg.WidowControl
	}
	if _fcdf.NumPr == nil {
		_fcdf.NumPr = _bfdg.NumPr
	}
	if _fcdf.SuppressLineNumbers == nil {
		_fcdf.SuppressLineNumbers = _bfdg.SuppressLineNumbers
	}
	if _fcdf.PBdr == nil {
		_fcdf.PBdr = _bfdg.PBdr
	}
	if _fcdf.Shd == nil {
		_fcdf.Shd = _bfdg.Shd
	}
	if _fcdf.Tabs == nil {
		_fcdf.Tabs = _bfdg.Tabs
	}
	if _fcdf.SuppressAutoHyphens == nil {
		_fcdf.SuppressAutoHyphens = _bfdg.SuppressAutoHyphens
	}
	if _fcdf.Kinsoku == nil {
		_fcdf.Kinsoku = _bfdg.Kinsoku
	}
	if _fcdf.WordWrap == nil {
		_fcdf.WordWrap = _bfdg.WordWrap
	}
	if _fcdf.OverflowPunct == nil {
		_fcdf.OverflowPunct = _bfdg.OverflowPunct
	}
	if _fcdf.TopLinePunct == nil {
		_fcdf.TopLinePunct = _bfdg.TopLinePunct
	}
	if _fcdf.AutoSpaceDE == nil {
		_fcdf.AutoSpaceDE = _bfdg.AutoSpaceDE
	}
	if _fcdf.AutoSpaceDN == nil {
		_fcdf.AutoSpaceDN = _bfdg.AutoSpaceDN
	}
	if _fcdf.Bidi == nil {
		_fcdf.Bidi = _bfdg.Bidi
	}
	if _fcdf.AdjustRightInd == nil {
		_fcdf.AdjustRightInd = _bfdg.AdjustRightInd
	}
	if _fcdf.SnapToGrid == nil {
		_fcdf.SnapToGrid = _bfdg.SnapToGrid
	}
	if _fcdf.Spacing == nil {
		_fcdf.Spacing = _bfdg.Spacing
	}
	if _fcdf.Ind == nil {
		_fcdf.Ind = _bfdg.Ind
	}
	if _fcdf.ContextualSpacing == nil {
		_fcdf.ContextualSpacing = _bfdg.ContextualSpacing
	}
	if _fcdf.MirrorIndents == nil {
		_fcdf.MirrorIndents = _bfdg.MirrorIndents
	}
	if _fcdf.SuppressOverlap == nil {
		_fcdf.SuppressOverlap = _bfdg.SuppressOverlap
	}
	if _fcdf.Jc == nil {
		_fcdf.Jc = _bfdg.Jc
	}
	if _fcdf.TextDirection == nil {
		_fcdf.TextDirection = _bfdg.TextDirection
	}
	if _fcdf.TextAlignment == nil {
		_fcdf.TextAlignment = _bfdg.TextAlignment
	}
	if _fcdf.TextboxTightWrap == nil {
		_fcdf.TextboxTightWrap = _bfdg.TextboxTightWrap
	}
	if _fcdf.OutlineLvl == nil {
		_fcdf.OutlineLvl = _bfdg.OutlineLvl
	}
	if _fcdf.DivId == nil {
		_fcdf.DivId = _bfdg.DivId
	}
	if _fcdf.CnfStyle == nil {
		_fcdf.CnfStyle = _bfdg.CnfStyle
	}
	if _fcdf.PPrChange == nil {
		_fcdf.PPrChange = _bfdg.PPrChange
	}
	return _fcdf
}

func (_gafd *convertContext) addEmptyLine() {
	_gafd.addTextSymbol(&symbol{_fge: "\u000d", _cdf: 0, _cgb: _gafd._cecg._ef})
}

func (_abed *convertContext) alignParagraph() {
	_bgd := _abed._cecg
	if _bgd._ae == _ea.TextAlignmentLeft {
		return
	}
	_gaf := len(_bgd._cc) - 1
	for _dgd, _fgdf := range _bgd._cc {
		_bdeb := len(_fgdf._dag) - 1
		for _caff, _afg := range _fgdf._dag {
			_dbf := true
			_fbea := len(_afg._bda)
			_fdcc := 0.0
			for _fbead := len(_afg._bda) - 1; _fbead >= 0; _fbead-- {
				_cecf := _afg._bda[_fbead]
				if _dbf && _cecf._gfb {
					_fbea = _fbead
				} else {
					_dbf = false
					for _, _beac := range _cecf._gaa {
						_fdcc += _beac._cdf
					}
				}
			}
			_afg._bda = _afg._bda[:_fbea]
			_cga := _afg._cf - _afg._bd - _fdcc
			switch _bgd._ae {
			case _ea.TextAlignmentRight:
				_afg.moveRight(_cga)
			case _ea.TextAlignmentCenter:
				_afg.moveRight(_cga / 2)
			case _ea.TextAlignmentJustify:
				if _dgd != _gaf || _caff != _bdeb {
					_gagg := []*word{}
					for _, _deca := range _afg._bda {
						if _deca._gfb {
							_gagg = append(_gagg, _deca)
						}
					}
					_bdb := _cga / float64(len(_gagg))
					for _, _cebe := range _gagg {
						_cebe._bfdc += _bdb
					}
					var _fad *word
					for _, _eebd := range _afg._bda {
						if _fad != nil {
							_eebd._cfe = _fad._cfe + _fad._bfdc
						}
						_fad = _eebd
					}
				}
			}
		}
	}
}

func (_agbd *convertContext) currentParagraphOverflowsCurrentPage() bool {
	_fggb := _agbd._cecg._bgb + _agbd._cecg._ac.Top + _agbd._cecg._ac.Bottom
	_dcdgg := _agbd._fbcd._eeg.Bottom - _agbd._cecg._cd
	if len(_agbd._fbcd._eaf) == 0 && len(_agbd._cecg._gbg) > 0 {
		_dcdgg -= _ga
	}
	return _fggb+_agbd._cecg._egd > _dcdgg || _fggb+_agbd._cecg._cg > _dcdgg
}

func (_fcdaa *convertContext) addParagraphWithTable(_ggca _ea.Table, _gbcd, _bcgb float64) {
	_fcdaa.newParagraph()
	_fcdaa._cecg._ac = &_a.Rectangle{Top: _ffcdc(2), Bottom: _ffcdc(2), Left: 0, Right: 0}
	_fcdaa._cecg._fdg = &tableWrapper{_cca: &_ggca, _fb: _gbcd}
	_fcdaa._cecg._af = _bcgb
	_fcdaa._cecg._egd = _ggca.Height()
	_fcdaa.determineParagraphBounds()
	_fcdaa.addCurrentParagraphToCurrentPage()
}

func (_acgc *convertContext) makeBlockFromWdWsp(_fegf *_bc.WdWsp) (*block, error) {
	if _gdabc := _fegf.WChoice; _gdabc != nil {
		if _bcgd := _gdabc.Txbx; _bcgd != nil {
			if _ffag := _bcgd.TxbxContent; _ffag != nil {
				if _dfgc := _ffag.EG_ContentBlockContent; len(_dfgc) > 0 {
					if _fddae := _fegf.SpPr; _fddae != nil {
						if _ebcc := _fddae.Xfrm; _ebcc != nil {
							if _eccb := _ebcc.Ext; _eccb != nil {
								_fadb := _eag.FromEMU(_eccb.CxAttr)
								_ddfcg := _eag.FromEMU(_eccb.CyAttr)
								_eefadg := &_a.Rectangle{Top: _dadd, Bottom: _dadd, Left: _dadd, Right: _dadd}
								_bfda, _feba := _acgc.makePdfBlockFromCBCs([][]*_bc.EG_ContentBlockContent{_dfgc}, _fadb, _ddfcg, _eefadg, false, nil)
								if _feba != nil {
									return nil, _feba
								}
								var _ggcc bool
								var _eeadc float64
								var _eaedg _ea.Color
								if _ffec := _fddae.PrstGeom; _ffec != nil {
									if _ffec.PrstAttr == _bb.ST_ShapeTypeRect {
										if _edgfg := _fddae.Ln; _edgfg != nil {
											if _ddfea := _edgfg.WAttr; _ddfea != nil {
												_ggcc = true
												_eeadc = _eag.FromEMU(int64(*_ddfea))
												_eaedg = _ea.ColorBlack
												if _effe := _edgfg.SolidFill; _effe != nil {
													if _gddd := _effe.SrgbClr; _gddd != nil {
														_eaedg = _ea.ColorRGBFromHex("\u0023" + _gddd.ValAttr)
													}
												}
											}
										}
									}
								}
								_baaae := &block{_egb: _bfda, _dgg: _ggcc, _ebg: _eeadc, _dc: _eaedg}
								return _baaae, nil
							}
						}
					}
				}
			}
		}
	}
	return nil, nil
}

func (_fecc *convertContext) getStyleProps(_fddf string, _gdabf _e.Style) (*_bc.CT_PPrGeneral, *_bc.CT_RPr) {
	var _fegb *_bc.CT_PPrGeneral
	var _dbge *_bc.CT_RPr
	_acbg := _fecc._cffa.GetStyleByID(_fddf)
	_eece := int64(0)
	_edad := true
	if _cag := _acbg.X(); _cag != nil {
		_fegb = _cag.PPr
		_dbge = _cag.RPr
		if _cag.UiPriority != nil {
			_eece = _cag.UiPriority.ValAttr
		}
		if _bbff := _cag.BasedOn; _bbff != nil {
			_ffd, _cfbc := _fecc.getStyleProps(_bbff.ValAttr, _acbg)
			if _aab := _gdabf.X(); _aab != nil {
				if _aab.UiPriority != nil && _eece > 0 {
					if _cag.UiPriority.ValAttr > _eece {
						_edad = false
					}
				}
				if _aab.QFormat != nil && _cag.QFormat != nil && _dgfgag(_aab.QFormat) && _dgfgag(_cag.QFormat) {
					_edad = false
				}
			}
			if _edad {
				_fegb = _bcda(_fegb, _ffd)
				_dbge = _gade(_dbge, _cfbc)
			}
		}
	}
	return _fegb, _dbge
}

func (_edca *convertContext) addSeparator() {
	_edca.newParagraph()
	_edca._cecg._fgg = true
	_edca._cecg._egd = _ga
	if _edca.currentParagraphOverflowsCurrentPage() {
		_edca.moveCurrentParagraphToNewPage()
	}
	_edca.addCurrentParagraphToCurrentPage()
}

func (_gcfb *convertContext) addAbsoluteCRC(_bdcd []*_bc.EG_ContentRunContent, _acd *_bc.CT_PPr) bool {
	for _, _bbd := range _bdcd {
		if _ffe := _bbd.R; _ffe != nil {
			if _acd != nil && _acd.PStyle != nil {
				_cgg := _gcfb._cffa.GetStyleByID(_acd.PStyle.ValAttr)
				if _dae := _cgg.X(); _dae != nil {
					if _dae.QFormat != nil && _dgfgag(_dae.QFormat) {
						if _dae.RPr != nil && _acd.RPr != nil {
							_acd.RPr = _efdg(_acd.RPr, _dae.RPr)
						}
					}
					if _dae.RPr != nil {
						if _dae.UiPriority != nil && _dae.UiPriority.ValAttr > 0 && _ffe.RPr == nil {
							_acd.RPr = _efdg(_acd.RPr, _dae.RPr)
						}
						_ffe.RPr = _gade(_ffe.RPr, _dae.RPr)
					}
					if _gcfb._eegda != nil {
						_bgbb, _febc := _gcfb.getStyleProps(_acd.PStyle.ValAttr, _cgg)
						_acd = _cceb(_acd, _bgbb, _febc)
						_ffe.RPr = _gade(_ffe.RPr, _febc)
					}
				}
			}
			_cddc := _acd != nil || _ffe.RPr != nil
			if len(_ffe.EG_RunInnerContent) == 0 && _cddc {
				_gcfb.addEmptyLine()
			}
			_aaf := _ggbc(_gcfb._cffa, _ffe.RPr, _acd)
			if _gcfb._eegda != nil {
				_gcfb.addAbsoluteRIC(nil, _aaf)
				_gcfb._eegda = nil
				_gcfb._cecg._fcf = true
			}
			for _, _bdcg := range _ffe.EG_RunInnerContent {
				if _gcfb.addAbsoluteRIC(_bdcg, _aaf) {
					return true
				}
				_gcfb._cecg._fcf = false
			}
			for _, _eede := range _ffe.Extra {
				if _cff, _gcbf := _eede.(*_bc.AlternateContentRun); _gcbf {
					if _cdfc := _cff.Choice; _cdfc != nil {
						if _fccb := _cdfc.Drawing; _fccb != nil {
							if len(_fccb.Inline) > 0 {
								for _, _dad := range _fccb.Inline {
									_dded := _dad.Extent
									if _dded == nil {
										return false
									}
									_ffff := _eag.FromEMU(_dded.CxAttr)
									_fdgb := _eag.FromEMU(_dded.CyAttr)
									if _cge := _dad.Graphic; _cge != nil {
										if _bgbe := _cge.GraphicData; _bgbe != nil {
											for _, _gaae := range _bgbe.Any {
												if _daa, _gagd := _gaae.(*_bc.WdWsp); _gagd {
													_bbc, _gcef := _gcfb.makeBlockFromWdWsp(_daa)
													if _gcef != nil {
														_da.Log.Debug("C\u0061\u006e\u006e\u006ft \u0072e\u0061\u0064\u0020\u0062\u006co\u0063\u006b\u003a\u0020\u0025\u0073", _gcef)
													}
													if _bbc == nil {
														continue
													}
													_bbc._egb.Scale(_ffff/_bbc._egb.Width(), _fdgb/_bbc._egb.Height())
													_gcfb.addInlineSymbol(&symbol{_cgb: _fdgb, _cdf: _ffff, _cce: _bbc})
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func (_bgfe *convertContext) determineParagraphBounds() {
	_bgfe._cecg._eg = _bgfe._fbcd._eeg.Left + _bgfe._cecg._ac.Left
	_bgfe._cecg._dg = _bgfe._cecg._eg + _bgfe._cecg._af
	_bgfe._cecg._dee = _bgfe._fbcd._eeg.Right - _bgfe._cecg._ac.Right
}

func (_edda *convertContext) addEndnotes() {
	for _ebb, _bdea := range _edda._agae {
		if _ebb == 0 {
			_edda.addSeparator()
		}
		_edda._eegda = &prefix{_ebab: _bdea._bce}
		for _aagd, _bbgc := range _bdea._ec {
			if _ebb != 0 || _aagd != 0 {
				_edda._dade = true
			}
			_edda.addAbsoluteCBCs(_bbgc.EG_ContentBlockContent, nil)
		}
	}
	_edda._dade = false
}

func _baba(_fada string) (float64, float64) {
	_eeca := _g.SplitN(_fada, "\u002c", 2)
	_cfeaf := _g.ReplaceAll(_eeca[0], "\u0070\u0074", "")
	_edba := _g.ReplaceAll(_eeca[1], "\u0070\u0074", "")
	_bcea, _fcac := _b.ParseFloat(_cfeaf, 64)
	if _fcac != nil {
		_da.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020U\u006e\u0061\u0062\u006ce p\u0061rs\u0065\u0020\u0078\u003a\u0020\u0025\u0076 t\u006f\u0020\u0066\u006c\u006f\u0061\u00746\u0034", _fcac.Error())
	}
	_gege, _fcac := _b.ParseFloat(_edba, 64)
	if _fcac != nil {
		_da.Log.Debug("\u0045\u0052RO\u0052\u003a\u0020U\u006e\u0061\u0062\u006ce p\u0061rs\u0065\u0020\u0079\u003a\u0020\u0025\u0076 t\u006f\u0020\u0066\u006c\u006f\u0061\u00746\u0034", _fcac.Error())
	}
	return _bcea, _gege
}

func _bcffb(_fbfg *_e.Document, _ggae string) []*_bc.CT_TblStylePr {
	_gbac := _fbfg.GetStyleByID(_ggae)
	var _gfgc []*_bc.CT_TblStylePr
	if _ecbe := _gbac.X(); _ecbe != nil {
		if _badd := _ecbe.BasedOn; _badd != nil {
			_bcffb(_fbfg, _badd.ValAttr)
		}
		if len(_ecbe.TblStylePr) > 0 {
			_gfgc = _ecbe.TblStylePr
		}
	}
	return _gfgc
}

func (_ggdb *convertContext) addCellToTable(_efbb *_ea.Table, _efce *_bc.CT_Tc, _gge *_bc.CT_TblPr, _afgd *_bc.CT_TblPrEx, _dbbbc, _bebg, _efga, _cfgf int, _caafc []*_bc.CT_TblStylePr, _fefg *_bc.CT_PPrGeneral, _edcb *_bc.CT_RPr, _fba bool) bool {
	var _fccc *_ea.TableCell
	_bcdeg := 1
	_adf := _efce.TcPr
	_dfce := _bc.NewCT_RPr()
	for _, _cbd := range _caafc {
		if _dbbbc == 0 && _cbd.TypeAttr == _bc.ST_TblStyleOverrideTypeFirstRow {
			_adf = _cacg(_adf, _cbd.TcPr)
			_edcb = _gade(_dfce, _cbd.RPr)
			break
		}
		if _bebg == 0 && _cbd.TypeAttr == _bc.ST_TblStyleOverrideTypeFirstCol {
			_adf = _cacg(_adf, _cbd.TcPr)
			_edcb = _gade(_dfce, _cbd.RPr)
		}
		if _dbbbc == _efga-1 && _cbd.TypeAttr == _bc.ST_TblStyleOverrideTypeLastRow {
			_adf = _cacg(_adf, _cbd.TcPr)
			_edcb = _gade(_dfce, _cbd.RPr)
		}
		if _bebg == _cfgf-1 && _cbd.TypeAttr == _bc.ST_TblStyleOverrideTypeLastCol {
			_adf = _cacg(_adf, _cbd.TcPr)
			_edcb = _gade(_dfce, _cbd.RPr)
		}
		if _dbbbc%2 != 0 && _cbd.TypeAttr == _bc.ST_TblStyleOverrideTypeBand1Horz {
			_adf = _cacg(_adf, _cbd.TcPr)
			if _bebg == 0 {
				_edcb = _gade(_dfce, _cbd.RPr)
			} else {
				_edcb = _dbea(_edcb, _cbd.RPr)
			}
		}
		if _bebg%2 != 0 && _cbd.TypeAttr == _bc.ST_TblStyleOverrideTypeBand1Vert {
			_adf = _cacg(_adf, _cbd.TcPr)
			if _dbbbc == 0 {
				_edcb = _gade(_dfce, _cbd.RPr)
			} else {
				_edcb = _dbea(_dfce, _cbd.RPr)
			}
		}
		if _dbbbc%2 == 0 && _cbd.TypeAttr == _bc.ST_TblStyleOverrideTypeBand2Horz {
			_adf = _cacg(_adf, _cbd.TcPr)
			if _bebg == _cfgf-1 {
				_edcb = _gade(_dfce, _cbd.RPr)
			} else {
				_edcb = _dbea(_dfce, _cbd.RPr)
			}
		}
		if _bebg%2 == 0 && _cbd.TypeAttr == _bc.ST_TblStyleOverrideTypeBand2Vert {
			_adf = _cacg(_adf, _cbd.TcPr)
			if _dbbbc == _efga-1 {
				_edcb = _gade(_dfce, _cbd.RPr)
			} else {
				_edcb = _dbea(_dfce, _cbd.RPr)
			}
		}
	}
	_dfbg := _adbc(_gge, _afgd, _adf, _dbbbc, _bebg, _efga, _cfgf)
	_cedb := _fd
	_efag := _ea.CellVerticalAlignmentTop
	if _dfbg != nil {
		if _dfbg.GridSpan != nil {
			_bcdeg = int(_dfbg.GridSpan.ValAttr)
		}
		_fccc = _efbb.MultiColCell(_bcdeg)
		if _dada := _dfbg.TcBorders; _dada != nil {
			if _dadb := _dada.Left; _dadb != nil {
				_bge, _geca, _ebaac := _dbafe(_dadb)
				_fccc.SetBorder(_ea.CellBorderSideLeft, _bge, _ebaac)
				if _geca != nil && *_geca != nil {
					_fccc.SetSideBorderColor(_ea.CellBorderSideLeft, *_geca)
				}
			}
			if _ceg := _dada.Top; _ceg != nil {
				_cbee, _cceg, _cgedb := _dbafe(_ceg)
				_fccc.SetBorder(_ea.CellBorderSideTop, _cbee, _cgedb)
				if _cceg != nil && *_cceg != nil {
					_fccc.SetSideBorderColor(_ea.CellBorderSideTop, *_cceg)
				}
			}
			if _acec := _dada.Right; _acec != nil {
				_cgad, _gebff, _bceg := _dbafe(_acec)
				_fccc.SetBorder(_ea.CellBorderSideRight, _cgad, _bceg)
				if _gebff != nil && *_gebff != nil {
					_fccc.SetSideBorderColor(_ea.CellBorderSideRight, *_gebff)
				}
			}
			if _fcce := _dada.Bottom; _fcce != nil {
				_eeba, _eadfb, _agc := _dbafe(_fcce)
				_fccc.SetBorder(_ea.CellBorderSideBottom, _eeba, _agc)
				if _eadfb != nil && *_eadfb != nil {
					_fccc.SetSideBorderColor(_ea.CellBorderSideBottom, *_eadfb)
				}
			}
		} else {
			_fccc.SetBorder(_ea.CellBorderSideAll, _ea.CellBorderStyleSingle, _ffcdc(0.125))
			_fccc.SetBorderColor(_ea.ColorBlack)
		}
		if _cba := _dfbg.Shd; _cba != nil {
			if _bfgb := _cba.FillAttr; _bfgb != nil {
				if _gbe := _bfgb.ST_HexColorRGB; _gbe != nil {
					_aegg := _ea.ColorRGBFromHex("\u0023" + *_gbe)
					_fccc.SetBackgroundColor(_aegg)
				}
			}
		}
		if _gef := _dfbg.TcMar; _gef != nil {
			_cedb = _gafdd(_gef.Left, 0, _fd)
		}
		if _gecg := _dfbg.VAlign; _gecg != nil {
			switch _gecg.ValAttr {
			case _bc.ST_VerticalJcCenter:
				_efag = _ea.CellVerticalAlignmentMiddle
			case _bc.ST_VerticalJcBottom:
				_efag = _ea.CellVerticalAlignmentBottom
			}
		}
	}
	if _fccc == nil {
		_fccc = _efbb.NewCell()
	}
	_fccc.SetVerticalAlignment(_efag)
	_fccc.SetIndent(_cedb)
	_gfgd := _efce.EG_BlockLevelElts
	_bggd := _ggdb._fcgg.NewStyledParagraph()
	_acgf := false
	_eeeg := false
	for _, _edb := range _gfgd {
		for _, _cceag := range _edb.EG_ContentBlockContent {
			for _, _cbdd := range _cceag.P {
				if _acgf {
					_cgaff := _bggd.Append("\u000a")
					_efbc := _ggdb._fcgg.NewTextStyle()
					_efbc.FontSize = 0
					_cgaff.Style = _efbc
				}
				_gged := false
				_gbcdg, _ := _ggdb.combinePPrWithStyles(_cbdd.PPr)
				if _gbcdg != nil && _gbcdg.PStyle != nil {
					_dffd := _ggdb._cffa.GetStyleByID(_gbcdg.PStyle.ValAttr)
					if _gaccc := _dffd.X(); _gaccc != nil {
						if _gaccc.QFormat != nil && _dgfgag(_gaccc.QFormat) {
							if _gbcdg.RPr != nil && _gaccc.RPr != nil {
								_gbcdg.RPr = _efdg(_gbcdg.RPr, _gaccc.RPr)
							}
							_gbcdg = _cceb(_cbdd.PPr, _gaccc.PPr, _gaccc.RPr)
							_gged = true
						} else {
							_aeff, _dgbb := _ggdb.getStyleProps(_gbcdg.PStyle.ValAttr, _dffd)
							_gbcdg = _cceb(_cbdd.PPr, _aeff, _dgbb)
						}
					}
				} else {
					_aece := _ggdb._cffa.Styles.ParagraphStyles()
					for _, _bdcc := range _aece {
						if _edbf := _bdcc.X().DefaultAttr; _edbf != nil {
							if _bdcda := _edbf.Bool; _bdcda != nil && *_bdcda {
								_gbcdg = _cceb(_cbdd.PPr, _bdcc.X().PPr, _bdcc.X().RPr)
							}
							if _cdbf := _edbf.ST_OnOff1; _cdbf == _ee.ST_OnOff1On {
								_gbcdg = _cceb(_cbdd.PPr, _bdcc.X().PPr, _bdcc.X().RPr)
							}
							break
						}
					}
				}
				if !_gged {
					_gbcdg = _cceb(_cbdd.PPr, _fefg, _edcb)
				}
				for _, _fcdc := range _cbdd.EG_PContent {
					for _, _efage := range _fcdc.EG_ContentRunContent {
						if _fagg := _efage.Sdt; _fagg != nil {
							if _fagg.SdtContent != nil {
								for _, _fcb := range _fagg.SdtContent.EG_ContentRunContent {
									if _dbed := _fcb.R; _dbed != nil {
										_fdga := _ggbc(_ggdb._cffa, _dbed.RPr, _gbcdg)
										for _, _bcag := range _dbed.EG_RunInnerContent {
											var _dcb *_ea.TextChunk
											if _bcag.T != nil && _bcag.T.Content != "" {
												_dabf := _bcag.T.Content
												if _fdga != nil && _dgfgag(_fdga.Caps) {
													_dabf = _g.ToUpper(_dabf)
												}
												_acgf = true
												_dcb = _bggd.Append(_dabf)
												_dcb.Style, _, _, _ = _ggdb.makeRunStyle(_fdga, false, false, false, false, false)
											} else if _bcag.LastRenderedPageBreak != nil && !_fba {
												_eeeg = true
											} else if _bcag.Br != nil {
												_bggd.Append("\u000a\u0020")
												_acgf = true
											}
										}
									}
								}
							}
						}
						if _febd := _efage.R; _febd != nil {
							_fcff := _ggbc(_ggdb._cffa, _febd.RPr, _gbcdg)
							for _, _bcff := range _febd.EG_RunInnerContent {
								var _ggf *_ea.TextChunk
								if _bcff.T != nil && _bcff.T.Content != "" {
									_fea := _bcff.T.Content
									if _fcff != nil && _dgfgag(_fcff.Caps) {
										_fea = _g.ToUpper(_fea)
									}
									_acgf = true
									_ggf = _bggd.Append(_fea)
									_ggf.Style, _, _, _ = _ggdb.makeRunStyle(_fcff, false, false, false, false, false)
								} else if _bcff.LastRenderedPageBreak != nil && !_fba {
									_eeeg = true
								} else if _bcff.Br != nil {
									_bggd.Append("\u000a\u0020")
									_acgf = true
								}
							}
						}
					}
				}
				if !_acgf {
					_edec := _ggbc(_ggdb._cffa, _bc.NewCT_RPr(), _gbcdg)
					_cfee := _bggd.Append("\u0020")
					_cfee.Style, _, _, _ = _ggdb.makeRunStyle(_edec, false, false, false, false, false)
				}
				if _bggd != nil {
					if _efag == _ea.CellVerticalAlignmentTop {
						_gbcdg.TextAlignment = _bc.NewCT_TextAlignment()
						_gbcdg.TextAlignment.ValAttr = _bc.ST_TextAlignmentTop
					}
					_ggdb.assignPropsToRelativeParagraph(_gbcdg, _bggd)
				}
			}
		}
	}
	_fccc.SetContent(_bggd)
	return _eeeg
}

func _cceb(_eddd *_bc.CT_PPr, _bffd *_bc.CT_PPrGeneral, _aefe *_bc.CT_RPr) *_bc.CT_PPr {
	if _eddd == nil {
		_eddd = _bc.NewCT_PPr()
	}
	if _bffd != nil {
		if _eddd.Jc == nil && _bffd.Jc != nil {
			_eddd.Jc = _bffd.Jc
		}
		if _eddd.Spacing == nil {
			_eddd.Spacing = _bffd.Spacing
		} else if _dbgbc := _bffd.Spacing; _dbgbc != nil {
			if _eddd.Spacing.BeforeAttr == nil {
				_eddd.Spacing.BeforeAttr = _dbgbc.BeforeAttr
			}
			if _eddd.Spacing.AfterAttr == nil {
				_eddd.Spacing.AfterAttr = _dbgbc.AfterAttr
			}
			if _eddd.Spacing.LineAttr == nil {
				_eddd.Spacing.LineAttr = _dbgbc.LineAttr
			}
		}
		if _bffd.ContextualSpacing != nil {
			_eddd.ContextualSpacing = _bffd.ContextualSpacing
		}
		if _bffd.Ind != nil {
			if _eddd.Ind == nil {
				_eddd.Ind = _bffd.Ind
			} else {
				_gga, _edbd := _eddd.Ind.FirstLineAttr == nil, _eddd.Ind.HangingAttr == nil
				if _gga && _edbd && _bffd.Ind.FirstLineAttr != nil {
					_eddd.Ind.FirstLineAttr = _bffd.Ind.FirstLineAttr
					_gga = false
				}
				if _gga && _edbd && _bffd.Ind.HangingAttr != nil {
					_eddd.Ind.HangingAttr = _bffd.Ind.HangingAttr
				}
				if _eddd.Ind.LeftAttr == nil {
					_eddd.Ind.LeftAttr = _bffd.Ind.LeftAttr
				}
				if _eddd.Ind.RightAttr == nil {
					_eddd.Ind.RightAttr = _bffd.Ind.RightAttr
				}
			}
		}
		if _eddd.Tabs == nil && _bffd.Tabs != nil {
			_eddd.Tabs = _bffd.Tabs
		}
		if _bffd.PBdr != nil {
			_eddd.PBdr = _bffd.PBdr
		}
	}
	if _aefe != nil {
		var _badf _bc.CT_ParaRPr
		if _eddd.RPr == nil {
			_badf = *_bc.NewCT_ParaRPr()
		} else {
			_badf = *_eddd.RPr
		}
		if _badf.Color == nil && _aefe.Color != nil {
			_badf.Color = _aefe.Color
		}
		if _badf.Spacing == nil && _aefe.Spacing != nil {
			_badf.Spacing = _aefe.Spacing
		}
		if _badf.Sz == nil && _aefe.Sz != nil {
			_badf.Sz = _aefe.Sz
		}
		if _badf.SzCs == nil && _aefe.SzCs != nil {
			_badf.SzCs = _aefe.SzCs
		}
		if _badf.B == nil && _aefe.B != nil {
			_badf.B = _aefe.B
		}
		if _badf.I == nil && _aefe.I != nil {
			_badf.I = _aefe.I
		}
		if _badf.RFonts == nil && _aefe.RFonts != nil {
			_badf.RFonts = _aefe.RFonts
		}
		if _badf.VertAlign == nil && _aefe.VertAlign != nil {
			_badf.VertAlign = _aefe.VertAlign
		}
		if _badf.Bdr == nil && _aefe.Bdr != nil {
			_badf.Bdr = _aefe.Bdr
		}
		if _badf.Caps == nil && _aefe.Caps != nil {
			_badf.Caps = _aefe.Caps
		}
		if _badf.SmallCaps == nil && _aefe.SmallCaps != nil {
			_badf.SmallCaps = _aefe.SmallCaps
		}
		_eddd.RPr = &_badf
	}
	return _eddd
}

type paragraph struct {
	_af  float64
	_ac  *_a.Rectangle
	_dg  float64
	_eg  float64
	_dee float64
	_bgb float64
	_egd float64
	_ae  _ea.TextAlignment
	_ef  float64
	_cc  []*line
	_fdg *tableWrapper
	_ebd []*image
	_caf []*image
	_gb  []*block
	_bfb []*block
	_gbg []*note
	_cd  float64
	_fcd []*zoneToSkip
	_cg  float64
	_fgg bool
	_gcf []*borderLine
	_fcf bool
}

func (_bdee *convertContext) addAnchorExtra(_gca []*_bc.EG_PContent) {
	for _, _beab := range _gca {
		for _, _ecbc := range _beab.EG_ContentRunContent {
			if _cfg := _ecbc.R; _cfg != nil {
				for _, _bcbc := range _cfg.Extra {
					if _aebc, _agf := _bcbc.(*_bc.AlternateContentRun); _agf {
						if _fbdb := _aebc.Choice; _fbdb != nil {
							if _bgbd := _fbdb.Drawing; _bgbd != nil {
								for _, _cgaf := range _bgbd.Anchor {
									var _bac, _edd, _ede, _ebca float64
									_bacd, _gab := _cgaf.PositionH, _cgaf.PositionV
									if _fec := _bacd.Choice; _fec != nil {
										if _fec.PosOffset != nil {
											_bac = _eag.FromEMU(int64(*_fec.PosOffset))
										}
									}
									if _bdfb := _gab.Choice; _bdfb != nil {
										if _bdfb.PosOffset != nil {
											_edd = _eag.FromEMU(int64(*_bdfb.PosOffset))
										}
									}
									if _aee := _cgaf.Extent; _aee != nil {
										_ebca = _eag.FromEMU(_aee.CxAttr)
										_ede = _eag.FromEMU(_aee.CyAttr)
									}
									_ecf := _bdee._cecg._bgb + _edd
									_gcc := _ecf + _ede
									_geb := _bdee._cecg._eg + _bac
									_bbfd := _geb + _ebca
									_becg := _edd + _ede
									if _becg > _bdee._cecg._cg {
										_bdee._cecg._cg = _becg
									}
									if _cgaf.Choice != nil && _cgaf.Choice.WrapNone == nil {
										_bdee._cecg._fcd = append(_bdee._cecg._fcd, &zoneToSkip{_gac: &_a.Rectangle{Top: _ecf, Bottom: _gcc, Left: _geb, Right: _bbfd}, _aba: _cgaf.Choice})
									}
									if _geg := _cgaf.Graphic; _geg != nil {
										if _cadd := _geg.GraphicData; _cadd != nil {
											for _, _eebc := range _cadd.Any {
												if _dcce, _cdab := _eebc.(*_bc.WdWsp); _cdab {
													_aafe, _aea := _bdee.makeBlockFromWdWsp(_dcce)
													if _aea != nil {
														_da.Log.Debug("C\u0061\u006e\u006e\u006ft \u0072e\u0061\u0064\u0020\u0062\u006co\u0063\u006b\u003a\u0020\u0025\u0073", _aea)
													}
													if _aafe != nil {
														_aafe._egb.Scale(_ebca/_aafe._egb.Width(), _ede/_aafe._egb.Height())
														_aafe._aef = _geb
														_aafe._aga = _ecf
														if _cgaf.BehindDocAttr {
															_bdee._cecg._bfb = append(_bdee._cecg._bfb, _aafe)
														} else {
															_bdee._cecg._gb = append(_bdee._cecg._gb, _aafe)
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func (_gdb *convertContext) drawPage(_gee *page) {
	if _gee._de {
		_cae := _gee._eeg.Top + _ga*_ca
		_aeb := _gee._eeg.Left
		_bba := _gee._eeg.Right
		_a.DrawLine(_gdb._fcgg, _aeb, _cae, _bba, _cae, _fg, _ea.ColorBlack)
	}
	for _, _ffaf := range _gee._gec {
		_fe(_gdb._fcgg, _ffaf)
	}
	for _, _ggce := range _gee._ab {
		_dab(_gdb._fcgg, _ggce)
	}
	for _, _gad := range _gee._eb {
		if _gad._fgg {
			_fgad := _gad._bgb + _ga*_ca
			_cdc := _gee._eeg.Left
			_gag := _cdc + _ffcdc(50)
			_a.DrawLine(_gdb._fcgg, _cdc, _fgad, _gag, _fgad, _fg, _ea.ColorBlack)
		} else {
			for _, _abf := range _gad._cc {
				for _, _ded := range _abf._dag {
					for _, _acf := range _ded._bda {
						for _, _edf := range _acf._gaa {
							if _edf._bea != nil {
								_edf._bea.SetPos(_acf._cfe+_edf._cdb, _gad._bgb+_abf._ag)
								_gdb._fcgg.Draw(_edf._bea)
							} else if _edf._cce != nil {
								_edf._cce._aef = _acf._cfe + _edf._cdb
								_edf._cce._aga = _gad._bgb + _abf._ag
								_dab(_gdb._fcgg, _edf._cce)
							} else {
								_fdca := _gdb._fcgg.NewStyledParagraph()
								if _edf._ffa {
									_edf._afa = 0
								} else if _edf._bcb {
									_edf._afa = 1.2*_abf._bca - _edf._cgb
								}
								_abe := _acf._cfe + _edf._cdb
								_fbg := _gad._bgb + _abf._ag + _edf._afa
								_fdca.SetPos(_abe, _fbg)
								var _gff *_ea.TextChunk
								if _edf._eae != "" {
									_gff = _fdca.AddExternalLink(_edf._fge, _edf._eae)
								} else {
									_gff = _fdca.Append(_edf._fge)
								}
								if _edf._ba != nil {
									_gff.Style = *_edf._ba
								}
								_gdb._fcgg.Draw(_fdca)
								if _edf._ad != nil {
									_bab := _fbg + _edf._ggc + 2.0
									_a.DrawLine(_gdb._fcgg, _abe, _bab, _abe+_edf._cdf, _bab, 1, *_edf._ad)
								}
							}
						}
					}
				}
			}
			if _gad._fdg != nil {
				_ebc := _ea.NewBlock(_gad._fdg._fb, _gad._ac.Top+_gad._egd+_gad._ac.Bottom)
				_ebc.SetPos(_gad._dg, _gad._bgb+_gad._ac.Top)
				_ebc.Draw(_gad._fdg._cca)
				_gdb._fcgg.Draw(_ebc)
			}
			if _gad._gcf != nil {
				_babc := (_gee._eeg.Left/_a.DefaultFontSize - 1)
				_dage := 1.5
				for _, _bfbe := range _gad._gcf {
					switch _bfbe._ecc {
					case _a.BorderPositionTop:
						_daf := _gad._bgb + _bfbe._ed
						_a.DrawLine(_gdb._fcgg, _gad._eg-_babc, _daf, _gad._eg+_bfbe._dca+_babc, _daf, _bfbe._adc, _bfbe._afaf)
					case _a.BorderPositionLeft:
						_bbf := _gad._bgb + _gad._egd - _gad._ac.Top - _gad._ac.Bottom - _bfbe._ed - _dage
						_gbbg := _bbf + _gad._egd + _gad._ac.Top + _gad._ac.Bottom
						_feg := _gad._eg - _babc
						_a.DrawLine(_gdb._fcgg, _feg, _bbf, _feg, _gbbg, _bfbe._dca, _bfbe._afaf)
					case _a.BorderPositionBottom:
						_dgb := _gad._bgb + _bfbe._ed + _gad._ac.Top + _gad._egd + _gad._ac.Bottom
						_a.DrawLine(_gdb._fcgg, _gad._eg-_babc, _dgb, _gad._eg+_bfbe._dca+_babc, _dgb, _bfbe._adc, _bfbe._afaf)
					case _a.BorderPositionRight:
						_fbe := _gad._bgb + _gad._egd - _gad._ac.Top - _gad._ac.Bottom - _bfbe._ed - _dage
						_gce := _fbe + _gad._egd + _gad._ac.Top + _gad._ac.Bottom
						_efa := _gad._dee + _babc
						_a.DrawLine(_gdb._fcgg, _efa, _fbe, _efa, _gce, _bfbe._dca, _bfbe._afaf)
					}
				}
			}
		}
	}
	for _, _eed := range _gee._fc {
		_fe(_gdb._fcgg, _eed)
	}
	for _, _df := range _gee._ceb {
		_dab(_gdb._fcgg, _df)
	}
	if len(_gee._eaf) > 0 {
		_dgf := _gee._eeg.Bottom + _ga*_ca
		_bdc := _gee._eeg.Left
		_dce := _bdc + _ffcdc(50)
		_a.DrawLine(_gdb._fcgg, _bdc, _dgf, _dce, _dgf, _fg, _ea.ColorBlack)
		_dec := _gee._eeg.Bottom + _ga
		for _, _agb := range _gee._eaf {
			_agb._be.SetPos(_gee._eeg.Left, _dec)
			_gdb._fcgg.Draw(_agb._be)
			_dec += _agb._be.Height()
		}
	}
}
