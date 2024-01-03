package convert

import (
	_fe "bytes"
	_f "errors"
	_cd "github.com/arcpd/msword/common/logger"
	_b "github.com/arcpd/msword/common/tempstorage"
	_ge "github.com/arcpd/msword/internal/convertutils"
	_dg "github.com/arcpd/msword/measurement"
	_d "github.com/arcpd/msword/presentation"
	_bb "github.com/arcpd/msword/schema/soo/dml"
	_cbf "github.com/arcpd/msword/schema/soo/dml/chart"
	_dd "github.com/arcpd/msword/schema/soo/pml"
	_gf "github.com/unidoc/unipdf/v3/contentstream/draw"
	_be "github.com/unidoc/unipdf/v3/creator"
	_dc "github.com/unidoc/unipdf/v3/model"
	_ebf "github.com/unidoc/unipdf/v3/render"
	_c "image"
	_g "image/color"
	_ee "image/draw"
	_eb "strconv"
	_cb "strings"
)

func _ebge(_ffcb, _fcg *_bb.CT_TextBody) (*_bb.CT_TextBodyProperties, *_bb.CT_TextListStyle) {
	if _ffcb == nil && _fcg == nil {
		return nil, nil
	}
	if _ffcb == nil {
		return _fcg.BodyPr, _fcg.LstStyle
	}
	if _fcg == nil {
		return _ffcb.BodyPr, _ffcb.LstStyle
	}
	_fcbd, _fac := _ffcb.BodyPr, _ffcb.LstStyle
	_dfdd, _dedd := _fcg.BodyPr, _fcg.LstStyle
	_gcfb := _gded(_fcbd, _dfdd)
	_fca := _dbc(_fac, _dedd)
	return _gcfb, _fca
}
func (_gag *convertContext) drawSlide() {
	_gag._bcgf.NewPage()
	for _, _ded := range _gag._cbcb {
		if _ded != nil {
			_gag._bcgf.MoveTo(0, 0)
			_gag._bcgf.Draw(_ded)
		}
	}
}
func (_adeg *convertContext) makeStyleFromRPr(_egba *_bb.CT_TextCharacterProperties) (*_be.TextStyle, bool, bool, bool) {
	var _cfe, _dfee, _abfc bool
	_cbc := _adeg._bcgf.NewTextStyle()
	if _egba != nil {
		_bcb := _ge.FontStyle_Regular
		_dgbcd := _abed(_egba.BAttr)
		_cdae := _abed(_egba.IAttr)
		if _dgbcd && _cdae {
			_bcb = _ge.FontStyle_BoldItalic
		} else if _dgbcd {
			_bcb = _ge.FontStyle_Bold
		} else if _cdae {
			_bcb = _ge.FontStyle_Italic
		}
		_abfc = _egba.UAttr != _bb.ST_TextUnderlineTypeUnset && _egba.UAttr != _bb.ST_TextUnderlineTypeNone
		_efda := "\u0064e\u0066\u0061\u0075\u006c\u0074"
		if _dec := _egba.Latin; _dec != nil {
			_efda = _dec.TypefaceAttr
		} else if _bgfb := _egba.Ea; _bgfb != nil {
			_efda = _bgfb.TypefaceAttr
		} else if _abca := _egba.Cs; _abca != nil {
			_efda = _abca.TypefaceAttr
		} else if _ebee := _egba.Sym; _ebee != nil {
			_efda = _ebee.TypefaceAttr
		}
		if _bcg, _cad := _ge.StdFontsMap[_efda]; _cad {
			_cbc.Font = _ge.AssignStdFontByName(_cbc, _bcg[_bcb])
		} else if _gcge := _ge.GetRegisteredFont(_efda, _bcb); _gcge != nil {
			_cbc.Font = _gcge
		} else {
			_cd.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0077\u0069\u0074h\u0020\u0073\u0074\u0079\u006c\u0065\u0020\u0025s\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0072\u0065\u0073\u0065\u0074 \u0074\u006f\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u002e", _efda, _bcb)
			_cbc.Font = _ge.AssignStdFontByName(_cbc, _ge.StdFontsMap["\u0064e\u0066\u0061\u0075\u006c\u0074"][_bcb])
		}
		var _bfaa float64
		if _aafd := _egba.SzAttr; _aafd != nil {
			_bfaa = float64(*_aafd) / 100
		} else {
			_bfaa = _ge.DefaultFontSize
		}
		if _deb := _egba.BaselineAttr; _deb != nil {
			if _gbgc := _deb.ST_PercentageDecimal; _gbgc != nil {
				if *_gbgc > 0 {
					_cfe = true
				} else if *_gbgc < 0 {
					_dfee = true
				}
			}
		}
		if _cfe || _dfee {
			_bfaa *= 0.64
		}
		_cbc.FontSize = _bfaa
		_gcgd := 0.0
		if _fba := _egba.SpcAttr; _fba != nil {
			if _bcd := _fba.ST_TextPointUnqualified; _bcd != nil {
				_gcgd = float64(*_bcd) / 100
			}
		}
		_cbc.CharSpacing = _gcgd
	}
	return &_cbc, _cfe, _dfee, _abfc
}
func (_ffg *convertContext) getShapes(_cg *_dd.CT_Shape) []_be.Drawable {
	_geg := []_be.Drawable{}
	_edd := _cg.SpPr
	if _edd == nil {
		return _geg
	}
	var _gg bool
	if _acc := _cg.UseBgFillAttr; _acc != nil {
		_gg = *_acc
	}
	_ggg, _dfd, _agg, _edde, _eddb, _bdf, _eeda := _ffg.getShapesFromSpPr(_edd, _cg.Style, _gg)
	_geg = append(_geg, _ggg...)
	if _aaa := _cg.TxBody; _aaa != nil {
		_abf, _cdcf, _bff, _fffa, _baff := _ffg.getPhData(_cg)
		if _abf != nil && !_eeda {
			_dfd, _agg, _edde, _eddb = _ge.GetDataFromXfrm(_abf)
		}
		_cdgg, _dgbc := _ffg.makePdfBlockFromTxBody(_aaa, _cdcf, _bff, _edde, _eddb, _bdf, _fffa, _baff)
		if _dgbc != nil {
			_cd.Log.Debug("\u0043\u0061\u006e\u006e\u006f\u0074\u0020\u006d\u0061\u006b\u0065\u0020\u0050\u0044\u0046\u0020\u0062\u006c\u006f\u0063\u006b\u0020\u0066\u0072o\u006d\u0020\u0074\u0065\u0078t\u0062\u006fx\u003a\u0020\u0025\u0073", _dgbc)
		} else if _cdgg != nil {
			_cdgg.SetPos(_dfd, _agg)
			_geg = append(_geg, _cdgg)
		}
	}
	return _geg
}
func (_cffd *convertContext) getBorderStyle(_cce *_bb.CT_LineProperties) (_be.CellBorderStyle, *_be.Color, float64) {
	if _cce == nil || _cce.NoFill != nil {
		return _be.CellBorderStyleNone, nil, 0
	}
	var _cdca _be.Color
	if _gfgg := _cce.SolidFill; _gfgg != nil {
		_cdca, _ = _cffd.getColorFromSolidFill(_gfgg)
	}
	_afg := 0.0
	if _gcdd := _cce.WAttr; _gcdd != nil {
		_afg = _dg.FromEMU(int64(*_gcdd))
	}
	return _be.CellBorderStyleSingle, &_cdca, _afg
}

type convertContext struct {
	_bcgf  *_be.Creator
	_ddcg  *_ge.Rectangle
	_bbbb  *_d.Presentation
	_feb   *_d.Slide
	_daec  *_dd.SldMaster
	_cge   *_dd.SldLayout
	_abcea float64
	_aed   float64
	_cbcb  []_be.Drawable
	_eafd  *background
	_gabbe *_bb.CT_TextParagraphProperties
	_gef   *_bb.CT_TextCharacterProperties
	_cadf  *_bb.CT_TextParagraphProperties
	_gec   *_bb.CT_TextCharacterProperties
	_agag  *_bb.CT_TextParagraphProperties
	_dbf   *_bb.CT_TextCharacterProperties
	_cfc   []*_bb.CT_TextParagraphProperties
	_fdce  []*_bb.CT_TextParagraphProperties
	_ggdb  []*_bb.CT_TextParagraphProperties
	_dbg   *_bb.Theme
	_dfeb  *_bb.CT_ColorMappingOverride
}

var _fcbe = []romanMatch{romanMatch{1000, "\u006d"}, romanMatch{900, "\u0063\u006d"}, romanMatch{500, "\u0064"}, romanMatch{400, "\u0063\u0064"}, romanMatch{100, "\u0063"}, romanMatch{90, "\u0078\u0063"}, romanMatch{50, "\u006c"}, romanMatch{40, "\u0078\u006c"}, romanMatch{10, "\u0078"}, romanMatch{9, "\u0069\u0078"}, romanMatch{5, "\u0076"}, romanMatch{4, "\u0069\u0076"}, romanMatch{1, "\u0069"}}

func (_abe *convertContext) addCellToTable(_fgb *_be.Table, _bda *_bb.CT_TableCell, _ggcb *_bb.CT_TablePartStyle, _dccf float64, _gacf, _gafc, _cdcc, _ece bool) float64 {
	var _ecfd *_be.TableCell
	_fgaa := 1
	if _bda.GridSpanAttr != nil {
		_fgaa = int(*_bda.GridSpanAttr)
	}
	_ecfd = _fgb.MultiColCell(_fgaa)
	_bcga := _bda.TcPr
	var _cedf *_bb.CT_TableStyleTextStyle
	if _ggcb != nil {
		_bcga = _gffb(_bcga, _ggcb.TcStyle, _gacf, _gafc, _cdcc, _ece)
		_cedf = _ggcb.TcTxStyle
	}
	_dgdf := _gd
	_caa := _be.CellVerticalAlignmentMiddle
	if _bcga != nil {
		if _adedb := _bcga.LnL; _adedb != nil {
			_fgbc, _gbf, _ecec := _abe.getBorderStyle(_adedb)
			_ecfd.SetBorder(_be.CellBorderSideLeft, _fgbc, _ecec)
			if _gbf != nil && *_gbf != nil {
				_ecfd.SetSideBorderColor(_be.CellBorderSideLeft, *_gbf)
			}
		}
		if _gaa := _bcga.LnT; _gaa != nil {
			_ceed, _aeee, _cdce := _abe.getBorderStyle(_gaa)
			_ecfd.SetBorder(_be.CellBorderSideTop, _ceed, _cdce)
			if _aeee != nil && *_aeee != nil {
				_ecfd.SetSideBorderColor(_be.CellBorderSideTop, *_aeee)
			}
		}
		if _faag := _bcga.LnR; _faag != nil {
			_gbea, _bcba, _ccd := _abe.getBorderStyle(_faag)
			_ecfd.SetBorder(_be.CellBorderSideRight, _gbea, _ccd)
			if _bcba != nil && *_bcba != nil {
				_ecfd.SetSideBorderColor(_be.CellBorderSideRight, *_bcba)
			}
		}
		if _afcf := _bcga.LnB; _afcf != nil {
			_cefe, _gbge, _accg := _abe.getBorderStyle(_afcf)
			_ecfd.SetBorder(_be.CellBorderSideBottom, _cefe, _accg)
			if _gbge != nil && *_gbge != nil {
				_ecfd.SetSideBorderColor(_be.CellBorderSideBottom, *_gbge)
			}
		}
		if _ceee := _bcga.MarLAttr; _ceee != nil {
			_dgdf = float64(_ge.FromSTCoordinate32(*_ceee))
		}
		switch _bcga.AnchorAttr {
		case _bb.ST_TextAnchoringTypeT:
			_caa = _be.CellVerticalAlignmentTop
		case _bb.ST_TextAnchoringTypeB:
			_caa = _be.CellVerticalAlignmentBottom
		}
		if _bcga.NoFill == nil {
			if _caf := _bcga.SolidFill; _caf != nil {
				_fggg, _ := _abe.getColorFromSolidFill(_caf)
				_ecfd.SetBackgroundColor(_fggg)
			}
		}
	}
	_ecfd.SetVerticalAlignment(_caa)
	_ecfd.SetIndent(_dgdf)
	var _eag float64
	if _edcd := _bda.TxBody; _edcd != nil {
		_abce := _abe.makePdfDivisionFromTxBody(_edcd, _dccf, _eag, _cedf)
		_eag = _abce.Height()
		_ecfd.SetContent(_abce)
	}
	return _eag
}
func (_fabc *textboxContext) newLine() {
	if _fabc._ccdc == nil {
		_fabc.newParagraph()
	}
	_ggb := _fabc._ccdc._egf + _fabc._ccdc._eeb
	_gcda := &line{}
	_gcda._ebebb = _fabc._ccdc._eaac
	if len(_fabc._ccdc._ccag) == 0 {
		_gcda._ebebb += _fabc._ccdc._defg
	}
	_gcda._bedc = _fabc._ccdc._fbbcg
	_gcda._adf = _gcda._ebebb
	_gcda._cdac = _ggb
	_fabc._ccdc._ccag = append(_fabc._ccdc._ccag, _gcda)
	_fabc._cfff = _gcda
}

// RegisterFont makes a PdfFont accessible for using in converting to PDF.
func RegisterFont(name string, style FontStyle, font *_dc.PdfFont) {
	_ge.RegisterFont(name, style, font)
}
func (_gggd *convertContext) makePdfBlockFromTxBody(_ggd *_bb.CT_TextBody, _ged *_bb.CT_TextBodyProperties, _aded *_bb.CT_TextListStyle, _eaef, _gff float64, _acgg _be.Color, _fdca, _gcg bool) (*_be.Block, error) {
	var _aaf *_bb.CT_TextParagraphProperties
	if _fccb := _ggd.LstStyle; _fccb != nil {
		var _deg *_bb.CT_TextParagraphProperties
		if _fccb.Lvl1pPr != nil {
			_deg = _fccb.Lvl1pPr
		}
		_aaf = _dabf(_deg, _fccb.DefPPr)
	}
	var _gdc *_bb.CT_TextParagraphProperties
	if _aaf != nil {
		if _fdca {
			_gdc = _gggd._fdce[0]
		} else if _gcg {
			_gdc = _gggd._ggdb[0]
		} else {
			_gdc = _gggd._cfc[0]
		}
		if _aded != nil {
			_gdc = _dabf(_dabf(_aded.Lvl1pPr, _aded.DefPPr), _gdc)
		}
		_gdc = _dabf(_aaf, _gdc)
	} else {
		if _fdca {
			_gdc = _gggd._cadf
		} else if _gcg {
			_gdc = _gggd._agag
		} else {
			_gdc = _gggd._gabbe
		}
	}
	_bfbg, _fggc := _eeaa(2.5), _eeaa(2.5)
	_agf, _bge := _eeaa(1.3), _eeaa(1.3)
	_fde := true
	_bded := _bb.ST_TextAnchoringTypeT
	if _ged != nil {
		if _afd := _ged.AnchorAttr; _afd != _bb.ST_TextAnchoringTypeUnset {
			_bded = _afd
		}
	}
	if _gcc := _ggd.BodyPr; _gcc != nil {
		if _fbc := _gcc.LInsAttr; _fbc != nil {
			_bfbg = _dg.FromEMU(_ge.FromSTCoordinate32(*_fbc))
		}
		if _fdgd := _gcc.TInsAttr; _fdgd != nil {
			_agf = _dg.FromEMU(_ge.FromSTCoordinate32(*_fdgd))
		}
		if _bfaf := _gcc.RInsAttr; _bfaf != nil {
			_fggc = _dg.FromEMU(_ge.FromSTCoordinate32(*_bfaf))
		}
		if _dfdb := _gcc.BInsAttr; _dfdb != nil {
			_bge = _dg.FromEMU(_ge.FromSTCoordinate32(*_dfdb))
		}
		_fde = _gcc.WrapAttr != _bb.ST_TextWrappingTypeNone
		if _ddaf := _gcc.AnchorAttr; _ddaf != _bb.ST_TextAnchoringTypeUnset {
			_bded = _gcc.AnchorAttr
		}
	}
	_eac := _ge.MakeTempCreator(_eaef, _gff)
	_eac.SetPageMargins(_bfbg, _fggc, _agf, _bge)
	_ddcd := &textboxContext{_bbga: _gggd, _efe: _fde, _gggc: _eaef - _bfbg - _fggc, _gbfa: _gff - _agf - _bge, _dcce: _eac, _edb: []*paragraph{}}
	_gga := 1
	for _, _egb := range _ggd.P {
		if _egb != nil {
			_gcf := _egb.PPr
			var _cdb *prefixData
			if _gcf != nil && _gcf.BuNone == nil {
				var _cgb string
				var _cde bool
				if _gcgc := _gcf.BuAutoNum; _gcgc != nil {
					var _def string
					if _fecf := _gcgc.StartAtAttr; _fecf != nil {
						_gga = int(*_fecf)
					}
					var _ddfc string
					switch _gcgc.TypeAttr {
					case _bb.ST_TextAutonumberSchemeAlphaUcParenBoth, _bb.ST_TextAutonumberSchemeAlphaUcParenR, _bb.ST_TextAutonumberSchemeAlphaUcPeriod:
						_ddfc = _aba(_gga, true)
					case _bb.ST_TextAutonumberSchemeAlphaLcParenBoth, _bb.ST_TextAutonumberSchemeAlphaLcParenR, _bb.ST_TextAutonumberSchemeAlphaLcPeriod:
						_ddfc = _aba(_gga, false)
					case _bb.ST_TextAutonumberSchemeRomanUcParenBoth, _bb.ST_TextAutonumberSchemeRomanUcParenR, _bb.ST_TextAutonumberSchemeRomanUcPeriod:
						_ddfc = _fecg(_gga, true)
					case _bb.ST_TextAutonumberSchemeRomanLcParenBoth, _bb.ST_TextAutonumberSchemeRomanLcParenR, _bb.ST_TextAutonumberSchemeRomanLcPeriod:
						_ddfc = _fecg(_gga, false)
					default:
						_ddfc = _eb.Itoa(_gga)
					}
					switch _gcgc.TypeAttr {
					case _bb.ST_TextAutonumberSchemeAlphaLcPeriod, _bb.ST_TextAutonumberSchemeAlphaUcPeriod, _bb.ST_TextAutonumberSchemeArabicPeriod, _bb.ST_TextAutonumberSchemeRomanLcPeriod, _bb.ST_TextAutonumberSchemeRomanUcPeriod, _bb.ST_TextAutonumberSchemeArabicDbPeriod, _bb.ST_TextAutonumberSchemeEa1ChsPeriod, _bb.ST_TextAutonumberSchemeEa1ChtPeriod, _bb.ST_TextAutonumberSchemeEa1JpnChsDbPeriod, _bb.ST_TextAutonumberSchemeEa1JpnKorPeriod, _bb.ST_TextAutonumberSchemeThaiAlphaPeriod, _bb.ST_TextAutonumberSchemeThaiNumPeriod, _bb.ST_TextAutonumberSchemeHindiAlphaPeriod, _bb.ST_TextAutonumberSchemeHindiNumPeriod, _bb.ST_TextAutonumberSchemeHindiAlpha1Period:
						_def = "\u002e"
					case _bb.ST_TextAutonumberSchemeAlphaLcParenR, _bb.ST_TextAutonumberSchemeAlphaUcParenR, _bb.ST_TextAutonumberSchemeArabicParenR, _bb.ST_TextAutonumberSchemeRomanLcParenR, _bb.ST_TextAutonumberSchemeRomanUcParenR, _bb.ST_TextAutonumberSchemeThaiAlphaParenR, _bb.ST_TextAutonumberSchemeThaiNumParenR, _bb.ST_TextAutonumberSchemeHindiNumParenR:
						_def = "\u0029"
					}
					_cgb = _ddfc + _def
					_gga++
				} else if _bfad := _gcf.BuChar; _bfad != nil {
					_cea := _bfad.CharAttr
					if _bfbc, _dgc := _ggcd[_cea]; _dgc {
						_cea = string(rune(_bfbc))
					} else {
						_cea = "\u2022"
					}
					_cgb = _cea
					_cde = true
				}
				if _cgb != "" {
					var _aggf, _fdec float64
					if _gcf.MarLAttr != nil {
						_aggf = _dg.FromEMU(int64(*_gcf.MarLAttr))
					}
					if _gcf.IndentAttr != nil {
						_fdec = _dg.FromEMU(int64(*_gcf.IndentAttr))
					}
					_cdb = &prefixData{_bdcb: _cgb, _bcag: _cde, _eea: _aggf, _gggg: _fdec}
				}
			}
			_gcf = _dabf(_gcf, _gdc)
			_eeec := _cagb(_egb.EndParaRPr, _gcf.DefRPr)
			_ddcd.newParagraph()
			_ddcd.assignPropsToCurrentParagraph(_gcf)
			_ddcd.newLine()
			_ddcd.newWord()
			for _gfc, _cff := range _egb.EG_TextRun {
				_ddcd.addTextRun(_cff, _eeec, _acgg, _cdb)
				if _gfc > 0 {
					_cdb = nil
				}
			}
			_ddcd.addCurrentWordToParagraph()
		}
		_ddcd.addCurrentParagraph()
	}
	_ddcd.alignVertically(_bded)
	_ddcd.drawParagraphs()
	return _ge.MakeBlockFromCreator(_eac)
}
func _dbc(_gedf, _geaa *_bb.CT_TextListStyle) *_bb.CT_TextListStyle {
	_efgf := _bb.NewCT_TextListStyle()
	if _gedf != nil {
		*_efgf = *_gedf
	}
	if _geaa == nil {
		return _efgf
	}
	_efgf.DefPPr = _dabf(_efgf.DefPPr, _geaa.DefPPr)
	_efgf.Lvl1pPr = _dabf(_efgf.Lvl1pPr, _geaa.Lvl1pPr)
	_efgf.Lvl2pPr = _dabf(_efgf.Lvl2pPr, _geaa.Lvl2pPr)
	_efgf.Lvl3pPr = _dabf(_efgf.Lvl3pPr, _geaa.Lvl3pPr)
	_efgf.Lvl4pPr = _dabf(_efgf.Lvl4pPr, _geaa.Lvl4pPr)
	_efgf.Lvl5pPr = _dabf(_efgf.Lvl5pPr, _geaa.Lvl5pPr)
	_efgf.Lvl6pPr = _dabf(_efgf.Lvl6pPr, _geaa.Lvl6pPr)
	_efgf.Lvl7pPr = _dabf(_efgf.Lvl7pPr, _geaa.Lvl7pPr)
	_efgf.Lvl8pPr = _dabf(_efgf.Lvl8pPr, _geaa.Lvl8pPr)
	_efgf.Lvl9pPr = _dabf(_efgf.Lvl9pPr, _geaa.Lvl9pPr)
	return _efgf
}
func (_bfba *convertContext) makePdfDivisionFromTxBody(_daad *_bb.CT_TextBody, _fdda, _gdb float64, _abcg *_bb.CT_TableStyleTextStyle) *_be.Division {
	_fab := _bfba._bcgf.NewDivision()
	_aeba := _bfba._gabbe
	_gabb := _bb.ST_TextAnchoringTypeT
	if _gfad := _daad.BodyPr; _gfad != nil {
		if _gbbe := _gfad.AnchorAttr; _gbbe != _bb.ST_TextAnchoringTypeUnset {
			_gabb = _gfad.AnchorAttr
		}
	}
	if _gfb := _daad.LstStyle; _gfb != nil {
		var _afcd *_bb.CT_TextParagraphProperties
		if _gfb.Lvl1pPr != nil {
			_afcd = _gfb.Lvl1pPr
		} else {
			_afcd = _bfba._cfc[0]
		}
		_aeba = _dabf(_afcd, _dabf(_gfb.DefPPr, _aeba))
	}
	for _, _dgbf := range _daad.P {
		if _dgbf != nil {
			_gge := _bfba._bcgf.NewStyledParagraph()
			_bde := _dabf(_dgbf.PPr, _aeba)
			_ebae := _cagb(_dgbf.EndParaRPr, _bde.DefRPr)
			if len(_dgbf.EG_TextRun) == 0 {
				_gge.Append("\u000a")
				_fab.Add(_gge)
				continue
			}
			for _, _ggc := range _dgbf.EG_TextRun {
				if _egd := _ggc.Br; _egd != nil {
					_gge.Append("\u000a")
				} else if _bgggc := _ggc.R; _bgggc != nil {
					_gegg := _fcae(_bgggc.RPr, _abcg)
					_gegg = _cagb(_gegg, _ebae)
					var _cefc _be.Color
					if _gegg.SolidFill != nil {
						_cefc, _ = _bfba.getColorFromSolidFill(_gegg.SolidFill)
					} else {
						_cefc = _be.ColorBlack
					}
					_aga, _dda, _fdc, _ := _bfba.makeStyleFromRPr(_gegg)
					_aga.Color = _cefc
					if _dda {
						_aga.TextRise = 0.5
					} else if _fdc {
						_aga.TextRise = -0.5
					}
					_eeg := _bgggc.T
					if _gegg.CapAttr == _bb.ST_TextCapsTypeAll {
						for _, _cee := range _eeg {
							_cee = []rune(_cb.ToUpper(string(_cee)))[0]
						}
					}
					_efc := _gge.Append(_eeg)
					_efc.Style = *_aga
				}
			}
			_ = _gabb
			_fab.Add(_gge)
		}
	}
	return _fab
}
func (_gffc *textboxContext) newWord() { _gffc._ffgdb = &word{_fdfb: true, _efdd: _gffc._cfff._adf} }
func (_dace *convertContext) getColorFromFontReference(_fagg *_bb.CT_FontReference) _be.Color {
	var _cgc _be.Color
	var _abac string
	if _dcfd := _fagg.SrgbClr; _dcfd != nil {
		_abac = _dcfd.ValAttr
	} else if _ggde := _fagg.SchemeClr; _ggde != nil {
		_abac = _ge.GetColorStringFromDmlColor(_dace._bbbb.GetColorBySchemeColor(_ggde.ValAttr))
		_abac = _ge.AdjustColor(_abac, _ggde.EG_ColorTransform)
	}
	if _abac != "" {
		_cgc = _be.ColorRGBFromHex("\u0023" + _abac)
	}
	return _cgc
}
func _eeaa(_agad float64) float64 { return _agad * _dg.Millimeter }
func _efcg(_beea string) []*symbol {
	_afcde := []*symbol{}
	for _, _degc := range _beea {
		_afcde = append(_afcde, &symbol{_ffd: string(_degc)})
	}
	return _afcde
}
func _fecg(_ffaf int, _aabb bool) string {
	_eedf := _fe.NewBuffer([]byte{})
	for _, _fccfa := range _fcbe {
		for {
			if _ffaf < _fccfa._dedg {
				break
			}
			_eedf.WriteString(_fccfa._bbc)
			_ffaf -= _fccfa._dedg
		}
	}
	_ffdd := _eedf.String()
	if _aabb {
		_ffdd = _cb.ToUpper(_ffdd)
	}
	return _ffdd
}
func (_gfaed *convertContext) getColorFromSolidFill(_cdgc *_bb.CT_SolidColorFillProperties) (_be.Color, float64) {
	if _cdgc == nil {
		return nil, 1
	}
	var _cdbg string
	_aff := 1.0
	if _bfec := _cdgc.SrgbClr; _bfec != nil {
		_cdbg = _bfec.ValAttr
		_aff = _ge.GetOpacityFromColorTransform(_bfec.EG_ColorTransform)
	} else if _acd := _cdgc.SchemeClr; _acd != nil {
		_cdbg = _ge.GetColorStringFromDmlColor(_gfaed._feb.GetColorBySchemeColor(_acd.ValAttr))
		_cdbg = _ge.AdjustColor(_cdbg, _acd.EG_ColorTransform)
		_aff = _ge.GetOpacityFromColorTransform(_acd.EG_ColorTransform)
	}
	if _cdbg != "" {
		_dggc := _be.ColorRGBFromHex("\u0023" + _cdbg)
		return _dggc, _aff
	}
	return nil, 1
}
func (_decd *textboxContext) addTextSymbol(_fgec *symbol) {
	_cebb := _be.New()
	_ddcb := _cebb.NewStyledParagraph()
	_ddcb.SetMargins(0, 0, 0, 0)
	_cfeg := _ddcb.Append(_fgec._ffd)
	_bcaa := 0.0
	if _fgec._cdccd != nil {
		_cfeg.Style = *_fgec._cdccd
		if _fgec._cdccd.CharSpacing != 0 {
			_bcaa = _fgec._cdccd.CharSpacing
		}
	}
	_fgec._bbgg = _ddcb.Height()
	_fgec._bcac = _ddcb.Height() * 1.2
	if _fgec._bcac < _decd._ccdc._caff {
		_fgec._bcac = _decd._ccdc._caff
	}
	if _fgec._fbe == 0 {
		_fgec._fbe = _ddcb.Width() + _bcaa
	}
	if len(_decd._ffgdb._dacc) > 0 {
		_cbfa := _decd._ffgdb._dacc[len(_decd._ffgdb._dacc)-1]._ffd
		if _decd._ccdc._geda || _ge.IsNoSpaceLanguage(_cbfa) || (_cbfa == "\u0020") != (_fgec._ffd == "\u0020") {
			_decd.addCurrentWordToParagraph()
			_decd.newWord()
		}
	}
	_decd._ffgdb._dacc = append(_decd._ffgdb._dacc, _fgec)
	_fgec._gde = _decd._ffgdb._aef
	_decd._ffgdb._aef += _fgec._fbe
	if _fgec._ffd != "\u0020" {
		_decd._ffgdb._fdfb = false
	}
}

var _gd = _eeaa(1.9)

func _cagb(_dbde, _dffd *_bb.CT_TextCharacterProperties) *_bb.CT_TextCharacterProperties {
	_fbgcg := _bb.NewCT_TextCharacterProperties()
	if _dbde != nil {
		*_fbgcg = *_dbde
	}
	if _dffd == nil {
		return _fbgcg
	}
	if _fbgcg.KumimojiAttr == nil {
		_fbgcg.KumimojiAttr = _dffd.KumimojiAttr
	}
	if _fbgcg.LangAttr == nil {
		_fbgcg.LangAttr = _dffd.LangAttr
	}
	if _fbgcg.AltLangAttr == nil {
		_fbgcg.AltLangAttr = _dffd.AltLangAttr
	}
	if _fbgcg.SzAttr == nil {
		_fbgcg.SzAttr = _dffd.SzAttr
	}
	if _fbgcg.BAttr == nil {
		_fbgcg.BAttr = _dffd.BAttr
	}
	if _fbgcg.IAttr == nil {
		_fbgcg.IAttr = _dffd.IAttr
	}
	if _fbgcg.UAttr == _bb.ST_TextUnderlineTypeUnset {
		_fbgcg.UAttr = _dffd.UAttr
	}
	if _fbgcg.StrikeAttr == _bb.ST_TextStrikeTypeUnset {
		_fbgcg.StrikeAttr = _dffd.StrikeAttr
	}
	if _fbgcg.KernAttr == nil {
		_fbgcg.KernAttr = _dffd.KernAttr
	}
	if _fbgcg.CapAttr == _bb.ST_TextCapsTypeUnset {
		_fbgcg.CapAttr = _dffd.CapAttr
	}
	if _fbgcg.SpcAttr == nil {
		_fbgcg.SpcAttr = _dffd.SpcAttr
	}
	if _fbgcg.NormalizeHAttr == nil {
		_fbgcg.NormalizeHAttr = _dffd.NormalizeHAttr
	}
	if _fbgcg.BaselineAttr == nil {
		_fbgcg.BaselineAttr = _dffd.BaselineAttr
	}
	if _fbgcg.NoProofAttr == nil {
		_fbgcg.NoProofAttr = _dffd.NoProofAttr
	}
	if _fbgcg.DirtyAttr == nil {
		_fbgcg.DirtyAttr = _dffd.DirtyAttr
	}
	if _fbgcg.ErrAttr == nil {
		_fbgcg.ErrAttr = _dffd.ErrAttr
	}
	if _fbgcg.SmtCleanAttr == nil {
		_fbgcg.SmtCleanAttr = _dffd.SmtCleanAttr
	}
	if _fbgcg.SmtIdAttr == nil {
		_fbgcg.SmtIdAttr = _dffd.SmtIdAttr
	}
	if _fbgcg.BmkAttr == nil {
		_fbgcg.BmkAttr = _dffd.BmkAttr
	}
	if _fbgcg.Ln == nil {
		_fbgcg.Ln = _dffd.Ln
	}
	if _fbgcg.NoFill == nil {
		_fbgcg.NoFill = _dffd.NoFill
	}
	if _fbgcg.SolidFill == nil {
		_fbgcg.SolidFill = _dffd.SolidFill
	}
	if _fbgcg.BlipFill == nil {
		_fbgcg.BlipFill = _dffd.BlipFill
	}
	if _fbgcg.EffectLst == nil {
		_fbgcg.EffectLst = _dffd.EffectLst
	}
	if _fbgcg.EffectDag == nil {
		_fbgcg.EffectDag = _dffd.EffectDag
	}
	if _fbgcg.Highlight == nil {
		_fbgcg.Highlight = _dffd.Highlight
	}
	if _fbgcg.ULnTx == nil {
		_fbgcg.ULnTx = _dffd.ULnTx
	}
	if _fbgcg.ULn == nil {
		_fbgcg.ULn = _dffd.ULn
	}
	if _fbgcg.UFillTx == nil {
		_fbgcg.UFillTx = _dffd.UFillTx
	}
	if _fbgcg.UFill == nil {
		_fbgcg.UFill = _dffd.UFill
	}
	if _fbgcg.Latin == nil {
		_fbgcg.Latin = _dffd.Latin
	}
	if _fbgcg.Ea == nil {
		_fbgcg.Ea = _dffd.Ea
	}
	if _fbgcg.Cs == nil {
		_fbgcg.Cs = _dffd.Cs
	}
	if _fbgcg.Sym == nil {
		_fbgcg.Sym = _dffd.Sym
	}
	if _fbgcg.Rtl == nil {
		_fbgcg.Rtl = _dffd.Rtl
	}
	return _fbgcg
}
func (_caec *textboxContext) addCurrentWordToParagraph() {
	if _caec._efe && _caec._cfff._adf+_caec._ffgdb._aef > _caec._cfff._bedc {
		_caec.newLine()
	}
	if !_caec._ffgdb._fdfb || len(_caec._cfff._gcdg) > 0 {
		_caec._ffgdb._efdd = _caec._cfff._adf
		_caec._cfff._gcdg = append(_caec._cfff._gcdg, _caec._ffgdb)
		_caec._cfff._adf += _caec._ffgdb._aef
		for _, _abgb := range _caec._ffgdb._dacc {
			_caec.adjustHeights(_abgb._bcac)
		}
	}
}
func _afbd(_bdfa *_dd.CT_CommonSlideData, _ebaf _dd.ST_PlaceholderType, _dba *uint32, _affc bool) (*_bb.CT_Transform2D, *_bb.CT_TextBody, *bool, *bool) {
	if _bdfa != nil && (_ebaf != _dd.ST_PlaceholderTypeUnset || !_affc) {
		if _bdaf := _bdfa.SpTree; _bdaf != nil {
			for _, _ddad := range _bdaf.Choice {
				if _ddad != nil {
					for _, _ggaa := range _ddad.Sp {
						if _ggaa != nil {
							_bdec, _efa := _abgba(_ggaa)
							if _ebaf == _bdec {
								if (_affc && _efa == nil) || (!_affc && _efa != nil && *_efa == *_dba) {
									var _edgb *_bb.CT_Transform2D
									if _ggaa.SpPr != nil {
										_edgb = _ggaa.SpPr.Xfrm
									}
									_dfdf := _ebaf == _dd.ST_PlaceholderTypeTitle || _ebaf == _dd.ST_PlaceholderTypeCtrTitle
									_dbac := !_dfdf && _ebaf != _dd.ST_PlaceholderTypeUnset
									return _edgb, _ggaa.TxBody, &_dfdf, &_dbac
								}
							}
						}
					}
				}
			}
		}
	}
	return nil, nil, nil, nil
}

var _ggcd = map[string]int32{"\u0076": 9830, "\u00d8": 8594, "\u00fc": 8730}

func (_cada *convertContext) getShapeFromBlipFill(_dceg *_bb.CT_BlipFillProperties, _agfc, _cafg, _adbd, _ccgb float64) _be.Drawable {
	_becb, _fefe, _bggcg := _cada.makePdfImageFromBlipFill(_dceg)
	if _bggcg != nil {
		_cd.Log.Debug("\u0043\u0061\u006e\u006e\u006f\u0074\u0020\u006d\u0061\u006b\u0065\u0020\u0050D\u0046\u0020\u0069\u006d\u0061\u0067e\u0020\u0066\u0072\u006f\u006d\u0020\u0042\u006c\u0069\u0070\u0046\u0069\u006cl\u003a\u0020\u0025\u0073", _bggcg)
		return nil
	}
	if _becb == nil {
		return nil
	}
	if _cfb := _dceg.Tile; _cfb != nil {
		_becb = _cada.tileImage(_becb, _dceg.Tile, _adbd, _ccgb)
	}
	if _acf := _dceg.Stretch; _acf != nil {
		_becb, _agfc, _cafg = _cada.stretchImage(_becb, _dceg.Stretch, _agfc, _cafg, _adbd, _ccgb)
	}
	if len(_fefe) == 0 {
		_becb.SetPos(_agfc, _cafg)
		return _becb
	}
	_becb = _cada.applyBlipEffectsOnImg(_becb, _agfc, _cafg, _fefe)
	_becb.SetPos(_agfc, _cafg)
	return _becb
}

type prefixData struct {
	_bdcb string
	_bcag bool
	_eea  float64
	_gggg float64
}
type background struct {
	_daecf _be.Color
	_fdb   float64
	_fbbbb *_bb.CT_BlipFillProperties
}

func (_dagf *line) moveRight(_cfef float64) {
	for _, _bbd := range _dagf._gcdg {
		_bbd._efdd += _cfef
	}
}
func (_abee *convertContext) getStyleColors(_eef *_bb.CT_ShapeStyle) (_be.Color, _be.Color, _be.Color) {
	var _fafg, _bffa, _bgbb _be.Color
	if _caabc := _eef.LnRef; _caabc != nil {
		_bffa = _abee.getColorFromMatrixReference(_caabc)
	}
	if _bcdd := _eef.FillRef; _bcdd != nil {
		_bgbb = _abee.getColorFromMatrixReference(_bcdd)
	}
	if _gedc := _eef.FontRef; _gedc != nil {
		_fafg = _abee.getColorFromFontReference(_gedc)
	}
	return _fafg, _bgbb, _bffa
}
func _abed(_gcee *bool) bool { return _gcee != nil && *_gcee }
func (_fd *convertContext) extractDefaultProperties() {
	_ad := _fd._bbbb.X()
	_ef := _ad.DefaultTextStyle
	var _bg, _ag, _geb, _ae, _db, _dgb, _cc, _fg, _ca, _faa *_bb.CT_TextParagraphProperties
	if _ef != nil {
		_bg = _ef.DefPPr
		_ag = _dabf(_ef.Lvl1pPr, _bg)
		_geb = _dabf(_ef.Lvl2pPr, _bg)
		_ae = _dabf(_ef.Lvl3pPr, _bg)
		_db = _dabf(_ef.Lvl4pPr, _bg)
		_dgb = _dabf(_ef.Lvl5pPr, _bg)
		_cc = _dabf(_ef.Lvl6pPr, _bg)
		_fg = _dabf(_ef.Lvl7pPr, _bg)
		_ca = _dabf(_ef.Lvl8pPr, _bg)
		_faa = _dabf(_ef.Lvl9pPr, _bg)
		_fd._gabbe = _bg
		_fd._gef = _bg.DefRPr
	}
	_fd._cfc = make([]*_bb.CT_TextParagraphProperties, 9)
	_fd._cfc[0] = _ag
	_fd._cfc[1] = _geb
	_fd._cfc[2] = _ae
	_fd._cfc[3] = _db
	_fd._cfc[4] = _dgb
	_fd._cfc[5] = _cc
	_fd._cfc[6] = _fg
	_fd._cfc[7] = _ca
	_fd._cfc[8] = _faa
	_cdc := _fd._bbbb.SlideMasters()[0].X()
	_aa := _cdc.TxStyles
	_aag := _aa.TitleStyle
	_fd._cadf = _dabf(_aag.DefPPr, _bg)
	_fd._gec = _fd._cadf.DefRPr
	_fd._fdce = make([]*_bb.CT_TextParagraphProperties, 9)
	_fd._fdce[0] = _dabf(_aag.Lvl1pPr, _ag)
	_fd._fdce[1] = _dabf(_aag.Lvl2pPr, _geb)
	_fd._fdce[2] = _dabf(_aag.Lvl3pPr, _ae)
	_fd._fdce[3] = _dabf(_aag.Lvl4pPr, _db)
	_fd._fdce[4] = _dabf(_aag.Lvl5pPr, _dgb)
	_fd._fdce[5] = _dabf(_aag.Lvl6pPr, _cc)
	_fd._fdce[6] = _dabf(_aag.Lvl7pPr, _fg)
	_fd._fdce[7] = _dabf(_aag.Lvl8pPr, _ca)
	_fd._fdce[8] = _dabf(_aag.Lvl9pPr, _faa)
	_gcd := _aa.BodyStyle
	_fd._agag = _dabf(_gcd.DefPPr, _bg)
	_fd._dbf = _fd._agag.DefRPr
	_fd._ggdb = make([]*_bb.CT_TextParagraphProperties, 9)
	_fd._ggdb[0] = _dabf(_gcd.Lvl1pPr, _ag)
	_fd._ggdb[1] = _dabf(_gcd.Lvl2pPr, _geb)
	_fd._ggdb[2] = _dabf(_gcd.Lvl3pPr, _ae)
	_fd._ggdb[3] = _dabf(_gcd.Lvl4pPr, _db)
	_fd._ggdb[4] = _dabf(_gcd.Lvl5pPr, _dgb)
	_fd._ggdb[5] = _dabf(_gcd.Lvl6pPr, _cc)
	_fd._ggdb[6] = _dabf(_gcd.Lvl7pPr, _fg)
	_fd._ggdb[7] = _dabf(_gcd.Lvl8pPr, _ca)
	_fd._ggdb[8] = _dabf(_gcd.Lvl9pPr, _faa)
}
func (_daaa *textboxContext) alignParagraphsVertically(_gege _bb.ST_TextAnchoringType) {
	if _gege == _bb.ST_TextAnchoringTypeT {
		return
	}
	_daag := 0.0
	for _, _dcff := range _daaa._edb {
		_daag += _dcff._eeb + _dcff._egf + _dcff._ggf
	}
	var _fdgdg float64
	switch _gege {
	case _bb.ST_TextAnchoringTypeCtr:
		_fdgdg = (_daaa._gbfa - _daag) / 2
	case _bb.ST_TextAnchoringTypeB:
		_fdgdg = _daaa._gbfa - _daag
	}
	for _, _aedg := range _daaa._edb {
		_aedg._cfed += _fdgdg
	}
}

type romanMatch struct {
	_dedg int
	_bbc  string
}

func (_bad *textboxContext) alignParagraph() {
	_cffe := _bad._ccdc
	if _cffe._dcca == _bb.ST_TextAlignTypeL {
		return
	}
	_ddb := len(_cffe._ccag) - 1
	for _ceg, _dbgb := range _cffe._ccag {
		_dgagf := true
		_cfee := len(_dbgb._gcdg)
		_dfa := 0.0
		for _cbe := len(_dbgb._gcdg) - 1; _cbe >= 0; _cbe-- {
			_agfef := _dbgb._gcdg[_cbe]
			if _dgagf && _agfef._fdfb {
				_cfee = _cbe
			} else {
				_dgagf = false
				for _, _cafe := range _agfef._dacc {
					_dfa += _cafe._fbe
				}
			}
		}
		_dbgb._gcdg = _dbgb._gcdg[:_cfee]
		_agea := _dbgb._bedc - _dbgb._ebebb - _dfa
		switch _cffe._dcca {
		case _bb.ST_TextAlignTypeR:
			_dbgb.moveRight(_agea)
		case _bb.ST_TextAlignTypeCtr:
			_dbgb.moveRight(_agea / 2)
		case _bb.ST_TextAlignTypeJust:
			if _ceg != _ddb {
				_dfc := []*word{}
				for _, _ggfc := range _dbgb._gcdg {
					if _ggfc._fdfb {
						_dfc = append(_dfc, _ggfc)
					}
				}
				_fdea := _agea / float64(len(_dfc))
				for _, _edeb := range _dfc {
					_edeb._aef += _fdea
				}
				var _agc *word
				for _, _eeca := range _dbgb._gcdg {
					if _agc != nil {
						_eeca._efdd = _agc._efdd + _agc._aef
					}
					_agc = _eeca
				}
			}
		}
	}
}
func _fbbd(_bacf, _acdc *_bb.CT_TableStyleCellStyle) *_bb.CT_TableStyleCellStyle {
	_bgcb := _bb.NewCT_TableStyleCellStyle()
	if _bacf != nil {
		*_bgcb = *_bacf
	}
	if _acdc == nil {
		return _bgcb
	}
	if _bgcb.TcBdr == nil {
		_bgcb.TcBdr = _acdc.TcBdr
	}
	if _bgcb.Fill == nil {
		_bgcb.Fill = _acdc.Fill
	}
	if _bgcb.FillRef == nil {
		_bgcb.FillRef = _acdc.FillRef
	}
	return _bgcb
}
func (_edaf *textboxContext) addTextRun(_efbb *_bb.EG_TextRun, _abad *_bb.CT_TextCharacterProperties, _fdfa _be.Color, _dfda *prefixData) {
	if _cbcd := _efbb.Br; _cbcd != nil {
		_edaf.addCurrentWordToParagraph()
		_edaf.newLine()
		_edaf.newWord()
	} else if _bfc := _efbb.R; _bfc != nil {
		var _dcgb _be.Color
		if _bfc.RPr.SolidFill != nil {
			_dcgb, _ = _edaf._bbga.getColorFromSolidFill(_bfc.RPr.SolidFill)
		} else if _fdfa != nil {
			_dcgb = _fdfa
		} else if _abad.SolidFill != nil {
			_dcgb, _ = _edaf._bbga.getColorFromSolidFill(_abad.SolidFill)
		} else {
			_dcgb = _be.ColorBlack
		}
		_aedf := _cagb(_bfc.RPr, _abad)
		_gdf, _fgfa, _dgag, _eddd := _edaf._bbga.makeStyleFromRPr(_aedf)
		_gdf.Color = _dcgb
		if _dfda != nil {
			_edaf.addPrefix(_dfda, _gdf)
		}
		_fege := _efcg(_bfc.T)
		for _, _bag := range _fege {
			_bag._cdccd = _gdf
			_bag._aaae = _fgfa
			_bag._fccf = _dgag
			_bag._dfb = _eddd
			if _aedf.CapAttr == _bb.ST_TextCapsTypeAll {
				_bag._ffd = _cb.ToUpper(_bag._ffd)
			}
			_edaf.addTextSymbol(_bag)
		}
	}
}
func (_eec *convertContext) stretchImage(_aac *_be.Image, _dga *_bb.CT_StretchInfoProperties, _ddd, _acg, _bbfe, _bfb float64) (*_be.Image, float64, float64) {
	_eaa := _dga.FillRect
	if _eaa == nil {
		_bgf := _aac.Width() / _aac.Height()
		_gagd := _bbfe / _bfb
		var _cca, _fddb float64
		if _bgf > _gagd {
			_fddb = _bfb
			_cca = _bfb * _bgf
		} else {
			_cca = _bbfe
			_fddb = _bbfe / _bgf
		}
		_aac.Scale(_cca/_aac.Width(), _fddb/_aac.Height())
		return _aac, _ddd, _acg
	}
	var _eca, _df, _afcb, _baf float64
	if _eba := _eaa.LAttr; _eba != nil {
		_afcb = _ge.FromSTPercentage(_eba)
	}
	if _da := _eaa.TAttr; _da != nil {
		_eca = _ge.FromSTPercentage(_da)
	}
	if _fgg := _eaa.RAttr; _fgg != nil {
		_baf = _ge.FromSTPercentage(_fgg)
	}
	if _gbc := _eaa.BAttr; _gbc != nil {
		_df = _ge.FromSTPercentage(_gbc)
	}
	_eecd := _bbfe * (1 - _afcb - _baf)
	_bgg := _bfb * (1 - _eca - _df)
	_aac.Scale(_eecd/_aac.Width(), _bgg/_aac.Height())
	return _aac, _ddd + _afcb*_bbfe, _acg + _eca*_bfb
}
func (_ffef *textboxContext) addPrefix(_cdf *prefixData, _cedc *_be.TextStyle) {
	_defgf := _efcg(_cdf._bdcb)
	_bdgf := *_cedc
	if _cdf._bcag {
		_bdgf.Font = _ge.AssignStdFontByName(_bdgf, "\u0053\u0079\u006d\u0062\u006f\u006c")
	}
	for _, _aaee := range _defgf {
		_aaee._cdccd = &_bdgf
		_ffef.addTextSymbol(_aaee)
	}
	_gba := -(_cdf._gggg + _ffef._ffgdb._aef)
	if _gba < 0 {
		_gba = 0
	}
	_gebb := &symbol{_ffd: "\u0020", _fbe: _gba}
	_ffef.addTextSymbol(_gebb)
	_ffef._ffgdb._efdd += (_cdf._gggg + _cdf._eea)
}
func (_cfeeg *convertContext) getInfoFromLn(_fagf *_bb.CT_LineProperties) (_be.Color, float64, float64) {
	if _fagf == nil || _fagf.NoFill != nil {
		return nil, 0, 0
	}
	var _cebf float64
	_gbade, _efef := _cfeeg.getColorFromSolidFill(_fagf.SolidFill)
	if _eeea := _fagf.WAttr; _eeea != nil {
		_cebf = _dg.FromEMU(int64(*_eeea))
	} else {
		_cebf = 1
	}
	return _gbade, _cebf, _efef
}

// ConvertToPdf converts a presentation to a PDF file. This package is beta, breaking changes can take place.
func ConvertToPdf(pr *_d.Presentation) *_be.Creator {
	_de := pr.X().SldSz
	_ab := _dg.FromEMU(int64(_de.CxAttr))
	_ed := _dg.FromEMU(int64(_de.CyAttr))
	_ac := _be.PageSize{_ab, _ed}
	_fc := _be.New()
	_fc.SetPageSize(_ac)
	var _egc *_bb.Theme
	if len(pr.Themes()) > 0 {
		_egc = pr.Themes()[0]
	}
	for _, _fb := range pr.Slides() {
		if _fb.X() == nil {
			continue
		}
		_gfa := &convertContext{_bcgf: _fc, _feb: &_fb, _cge: _fb.GetSlideLayout(), _daec: pr.SlideMasters()[0].X(), _bbbb: pr, _dbg: _egc, _dfeb: _fb.X().ClrMapOvr, _abcea: _ac[1], _aed: _ac[0]}
		_gfa._cge = _fb.GetSlideLayout()
		_gfa.extractDefaultProperties()
		_gfa.makeSlide()
		_gfa.drawSlide()
	}
	return _fc
}
func (_bggc *textboxContext) newParagraph() {
	_egfc := &paragraph{}
	_egfc._cfed = _bggc._fcea
	_bggc._ccdc = _egfc
}
func (_bbff *textboxContext) alignSymbolsVertically() {
	for _, _gegda := range _bbff._edb {
		for _, _feef := range _gegda._ccag {
			_bgff := 0.0
			for _, _gbca := range _feef._gcdg {
				for _, _fega := range _gbca._dacc {
					if _fega._bcac > _bgff {
						_bgff = _fega._bcac
					}
				}
			}
			for _, _gage := range _feef._gcdg {
				for _, _gacd := range _gage._dacc {
					if _gacd._bbgg < _bgff {
						_gacd._gcff = _bgff - _gacd._bbgg
					}
				}
			}
		}
	}
}
func (_ddcdg *convertContext) applyBlipEffectsOnImg(_fffb *_be.Image, _acec, _fcgg float64, _abd []*_bb.CT_BlipChoice) *_be.Image {
	if len(_abd) == 0 {
		return _fffb
	}
	_fffb.SetPos(_acec, _fcgg)
	_dddc, _ecfc := _ddcdg.renderPageWithDrawableToGoImage(_fffb)
	if _ecfc != nil {
		_cd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0020a\u006e\u0020\u0069\u006d\u0061\u0067e\u0020\u0074\u006f\u0020\u0061\u0020\u0047\u006f\u0020\u0069\u006d\u0061\u0067e\u003a\u0020\u0025\u0073", _ecfc)
		return _fffb
	}
	_fabd, _ecfc := _ddcdg.renderCurrentStateToGoImage()
	if _ecfc != nil {
		_cd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0072\u0065n\u0064\u0065\u0072 t\u0068\u0065\u0020\u0063\u0075\u0072r\u0065\u006e\u0074\u0020\u0073\u0074\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0061\u0020G\u006f\u0020\u0069\u006d\u0061\u0067\u0065\u003a \u0025\u0073", _ecfc)
		return _fffb
	}
	_bfge := _dddc.Bounds()
	_cadad := _c.NewRGBA(_bfge)
	_bbcd, _bega := _fffb.Width(), _fffb.Height()
	for _, _cgfa := range _abd {
		for _, _cceb := range _cgfa.AlphaModFix {
			if _dffe := _cceb.AmtAttr; _dffe != nil {
				if _gegef := _dffe.ST_PositivePercentageDecimal; _gegef != nil {
					_deec := uint8(255 * (*_gegef) / 100000)
					_eccg := _c.NewUniform(_g.Alpha{_deec})
					_ee.Draw(_cadad, _bfge, _fabd, _c.Point{0, 0}, _ee.Src)
					_ee.DrawMask(_cadad, _bfge, _dddc, _c.Point{0, 0}, _eccg, _c.Point{0, 0}, _ee.Over)
				}
			}
		}
	}
	_abff := _c.Rect(int(_acec), int(_fcgg), int(_acec+_bbcd)+1, int(_fcgg+_bega)+1)
	_dfaf := _ge.CropImageByRect(_cadad, _abff)
	_ecdd, _ecfc := _ddcdg._bcgf.NewImageFromGoImage(_dfaf)
	if _ecfc != nil {
		_cd.Log.Debug("\u0045\u0072\u0072\u006f\u0072\u0020\u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065\u0020t\u006f \u0061\u0020\u0047\u006f\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _ecfc)
		return _fffb
	}
	return _ecdd
}
func (_gcba *convertContext) renderCurrentStateToGoImage() (_c.Image, error) {
	_fffad := _ge.MakeTempCreator(_gcba._aed, _gcba._abcea)
	_fffad.NewPage()
	for _, _aege := range _gcba._cbcb {
		if _aege != nil {
			_fffad.MoveTo(0, 0)
			_fffad.Draw(_aege)
		}
	}
	_eebaa, _fcbf := _ge.GetPageFromCreator(_fffad)
	if _fcbf != nil {
		return nil, _fcbf
	}
	return _ebf.NewImageDevice().Render(_eebaa)
}
func _abgba(_eecb *_dd.CT_Shape) (_dd.ST_PlaceholderType, *uint32) {
	if _edba := _eecb.NvSpPr; _edba != nil {
		if _gbgf := _edba.NvPr; _gbgf != nil {
			if _cdgb := _gbgf.Ph; _cdgb != nil {
				return _cdgb.TypeAttr, _cdgb.IdxAttr
			}
		}
	}
	return _dd.ST_PlaceholderTypeUnset, nil
}
func (_aab *convertContext) tileImage(_fda *_be.Image, _bae *_bb.CT_TileInfoProperties, _geac, _ccg float64) *_be.Image {
	_aee, _bbg := 1.0, 1.0
	if _eee := _bae.SxAttr; _eee != nil {
		_aee = _ge.FromSTPercentage(_eee)
	}
	if _bbf := _bae.SyAttr; _bbf != nil {
		_bbg = _ge.FromSTPercentage(_bbf)
	}
	_ffc := _ge.MakeTempCreator(_geac, _ccg)
	_fda.Scale(_aee, _bbg)
	_bbb, _ccgc := _fda.Width(), _fda.Height()
	var _ddc, _dce float64
	if _cae := _bae.TxAttr; _cae != nil {
		_ddc = _dg.FromEMU(_ge.FromSTCoordinate(*_cae))
	}
	if _dee := _bae.TyAttr; _dee != nil {
		_dce = _dg.FromEMU(_ge.FromSTCoordinate(*_dee))
	}
	if _ddc > 0 {
		_ddc -= _bbb
	}
	if _dce > 0 {
		_dce -= _ccgc
	}
	_cdg := _aab._aed/_bbb + 1
	_ace := _aab._abcea/_ccgc + 1
	for _gfac := 0.0; _gfac <= _cdg; _gfac++ {
		_fff := _gfac * _bbb
		for _ea := 0.0; _ea <= _ace; _ea++ {
			_ce := _ea * _ccgc
			_fda.SetPos(_fff+_ddc, _ce+_dce)
			_ffc.Draw(_fda)
		}
	}
	_ebc, _baee := _ge.GetPageFromCreator(_ffc)
	if _baee != nil {
		_cd.Log.Debug("\u0043\u0061\u006e\u006e\u006f\u0074 \u0067\u0065\u0074\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0066\u0072\u006fm\u0020\u0063\u0072\u0065\u0061\u0074\u006fr\u003a\u0020\u0025\u0073", _baee)
		return nil
	}
	_aabf, _baee := _ebf.NewImageDevice().Render(_ebc)
	if _baee != nil {
		_cd.Log.Debug("\u0043\u0061\u006eno\u0074\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _baee)
		return nil
	}
	_age, _baee := _aab._bcgf.NewImageFromGoImage(_aabf)
	if _baee != nil {
		_cd.Log.Debug("\u0043\u0061nn\u006f\u0074\u0020c\u0072\u0065\u0061\u0074e i\u006dag\u0065\u0020\u0066\u0072\u006f\u006d\u0020Go\u0049\u006d\u0061\u0067\u0065\u003a\u0020%\u0073", _baee)
		return nil
	}
	return _age
}

type textboxContext struct {
	_bbga  *convertContext
	_gggc  float64
	_gbfa  float64
	_dcce  *_be.Creator
	_fcea  float64
	_edb   []*paragraph
	_ccdc  *paragraph
	_cfff  *line
	_ffgdb *word
	_efe   bool
}

func (_cdccg *textboxContext) drawParagraphs() {
	_cdccg._dcce.NewPage()
	for _, _gdeb := range _cdccg._edb {
		for _, _dcfc := range _gdeb._ccag {
			for _, _ceeed := range _dcfc._gcdg {
				for _, _bbeb := range _ceeed._dacc {
					_aeg := _cdccg._dcce.NewStyledParagraph()
					if _bbeb._aaae {
						_bbeb._gcff = 0
					} else if _bbeb._fccf {
						_bbeb._gcff = 1.2*_dcfc._gaab - _bbeb._bbgg
					}
					_dcbd := _ceeed._efdd + _bbeb._gde
					_daaf := _gdeb._cfed + _dcfc._cdac + _bbeb._gcff
					_aeg.SetPos(_dcbd, _daaf)
					_gcgdb := _aeg.Append(_bbeb._ffd)
					if _bbeb._cdccd != nil {
						_gcgdb.Style = *_bbeb._cdccd
					}
					_cdccg._dcce.Draw(_aeg)
					if _bbeb._dfb {
						_egcg := _daaf + _bbeb._bbgg + 2
						_ge.DrawLine(_cdccg._dcce, _dcbd, _egcg, _dcbd+_bbeb._fbe, _egcg, 1, _bbeb._cdccd.Color)
					}
				}
			}
		}
	}
}
func _dabf(_cadd, _eab *_bb.CT_TextParagraphProperties) *_bb.CT_TextParagraphProperties {
	_cfedf := _bb.NewCT_TextParagraphProperties()
	if _cadd != nil {
		*_cfedf = *_cadd
	}
	if _eab == nil {
		return _cfedf
	}
	if _cfedf.MarLAttr == nil {
		_cfedf.MarLAttr = _eab.MarLAttr
	}
	if _cfedf.MarRAttr == nil {
		_cfedf.MarRAttr = _eab.MarRAttr
	}
	if _cfedf.LvlAttr == nil {
		_cfedf.LvlAttr = _eab.LvlAttr
	}
	if _cfedf.IndentAttr == nil {
		_cfedf.IndentAttr = _eab.IndentAttr
	}
	if _cfedf.AlgnAttr == _bb.ST_TextAlignTypeUnset {
		_cfedf.AlgnAttr = _eab.AlgnAttr
	}
	if _cfedf.DefTabSzAttr == nil {
		_cfedf.DefTabSzAttr = _eab.DefTabSzAttr
	}
	if _cfedf.RtlAttr == nil {
		_cfedf.RtlAttr = _eab.RtlAttr
	}
	if _cfedf.EaLnBrkAttr == nil {
		_cfedf.EaLnBrkAttr = _eab.EaLnBrkAttr
	}
	if _cfedf.FontAlgnAttr == _bb.ST_TextFontAlignTypeUnset {
		_cfedf.FontAlgnAttr = _eab.FontAlgnAttr
	}
	if _cfedf.LatinLnBrkAttr == nil {
		_cfedf.LatinLnBrkAttr = _eab.LatinLnBrkAttr
	}
	if _cfedf.HangingPunctAttr == nil {
		_cfedf.HangingPunctAttr = _eab.HangingPunctAttr
	}
	if _cfedf.LnSpc == nil {
		_cfedf.LnSpc = _eab.LnSpc
	}
	if _cfedf.SpcBef == nil {
		_cfedf.SpcBef = _eab.SpcBef
	}
	if _cfedf.SpcAft == nil {
		_cfedf.SpcAft = _eab.SpcAft
	}
	if _cfedf.BuClrTx == nil {
		_cfedf.BuClrTx = _eab.BuClrTx
	}
	if _cfedf.BuClr == nil {
		_cfedf.BuClr = _eab.BuClr
	}
	if _cfedf.BuSzTx == nil {
		_cfedf.BuSzTx = _eab.BuSzTx
	}
	if _cfedf.BuSzPct == nil {
		_cfedf.BuSzPct = _eab.BuSzPct
	}
	if _cfedf.BuSzPts == nil {
		_cfedf.BuSzPts = _eab.BuSzPts
	}
	if _cfedf.BuFontTx == nil {
		_cfedf.BuFontTx = _eab.BuFontTx
	}
	if _cfedf.BuFont == nil {
		_cfedf.BuFont = _eab.BuFont
	}
	if _cfedf.BuNone == nil {
		_cfedf.BuNone = _eab.BuNone
	}
	if _cfedf.BuAutoNum == nil {
		_cfedf.BuAutoNum = _eab.BuAutoNum
	}
	if _cfedf.BuChar == nil {
		_cfedf.BuChar = _eab.BuChar
	}
	if _cfedf.BuBlip == nil {
		_cfedf.BuBlip = _eab.BuBlip
	}
	if _cfedf.TabLst == nil {
		_cfedf.TabLst = _eab.TabLst
	}
	if _cfedf.ExtLst == nil {
		_cfedf.ExtLst = _eab.ExtLst
	}
	_cfedf.DefRPr = _cagb(_cfedf.DefRPr, _eab.DefRPr)
	return _cfedf
}

// RegisterFontsFromDirectory registers all fonts from the given directory automatically detecting font families and styles.
func RegisterFontsFromDirectory(dirName string) error { return _ge.RegisterFontsFromDirectory(dirName) }
func (_bfff *convertContext) makePdfImageFromBlipFill(_add *_bb.CT_BlipFillProperties) (*_be.Image, []*_bb.CT_BlipChoice, error) {
	if _cgfd := _add.Blip; _cgfd != nil {
		if _fbba := _cgfd.EmbedAttr; _fbba != nil {
			_bcbad, _cbba := _bfff._feb.GetImageByRelID(*_fbba)
			if _cbba {
				_cbfb, _faee := _b.Open(_bcbad.Path())
				if _faee != nil {
					_cd.Log.Debug("\u0046\u0069\u006c\u0065 o\u0070\u0065\u006e\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", _faee)
					return nil, nil, _faee
				}
				defer _cbfb.Close()
				_abfb, _, _faee := _c.Decode(_cbfb)
				if _faee != nil {
					_cd.Log.Debug("\u0044\u0065\u0063\u006fde\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020%\u0073", _faee)
					return nil, nil, _faee
				}
				if _ffda := _add.SrcRect; _ffda != nil {
					_gbad := _abfb.Bounds().Size()
					_efg := _gbad.X
					_ecc := _gbad.Y
					var _gabg, _aca, _gdcd, _dab int
					var _fdcag bool
					if _aaea := _ffda.LAttr; _aaea != nil {
						_gabg = int(float64(_efg) * _ge.FromSTPercentage(_aaea))
						_fdcag = true
					} else {
						_gabg = 0
					}
					if _bccc := _ffda.TAttr; _bccc != nil {
						_gdcd = int(float64(_ecc) * _ge.FromSTPercentage(_bccc))
						_fdcag = true
					} else {
						_gdcd = 0
					}
					if _cdcd := _ffda.RAttr; _cdcd != nil {
						_aca = int(float64(_efg) * (1 - _ge.FromSTPercentage(_cdcd)))
						_fdcag = true
					} else {
						_aca = _efg
					}
					if _fafb := _ffda.BAttr; _fafb != nil {
						_dab = int(float64(_ecc) * (1 - _ge.FromSTPercentage(_fafb)))
						_fdcag = true
					} else {
						_dab = _ecc
					}
					if _fdcag {
						_abfb = _ge.CropImageByRect(_abfb, _c.Rect(_gabg, _gdcd, _aca+1, _dab+1))
					}
				}
				_caab, _faee := _bfff._bcgf.NewImageFromGoImage(_abfb)
				if _faee != nil {
					_cd.Log.Debug("\u0043\u0061\u006e\u006e\u006ft\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0050\u0044\u0046\u0020\u0069m\u0061\u0067\u0065\u0020\u0066\u0072\u006f\u006d\u0020\u0047\u006f\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _faee)
					return nil, nil, _faee
				}
				return _caab, _cgfd.Choice, nil
			}
		}
	}
	return nil, nil, nil
}

type word struct {
	_dacc []*symbol
	_efdd float64
	_aef  float64
	_fdfb bool
}

func (_eafg *textboxContext) addCurrentParagraph() {
	_eafg._fcea = _eafg._ccdc._cfed + _eafg._ccdc._eeb + _eafg._ccdc._egf + _eafg._ccdc._ggf
	_eafg._edb = append(_eafg._edb, _eafg._ccdc)
	_eafg.alignParagraph()
}
func _gagea(_adc *_bb.CT_AdjPoint2D) (float64, float64) {
	var _ebad, _fdbc float64
	_fef, _gcde := _adc.XAttr, _adc.YAttr
	if _fccg := _fef.ST_Coordinate; _fccg != nil {
		_ebad = _dg.FromEMU(_ge.FromSTCoordinate(*_fccg))
	}
	if _deea := _gcde.ST_Coordinate; _deea != nil {
		_fdbc = _dg.FromEMU(_ge.FromSTCoordinate(*_deea))
	}
	return _ebad, _fdbc
}
func _bbbf(_cbee, _cac *_bb.CT_TablePartStyle) *_bb.CT_TablePartStyle {
	_dge := _bb.NewCT_TablePartStyle()
	if _cbee != nil {
		*_dge = *_cbee
	}
	if _cac == nil {
		return _dge
	}
	if _dge.TcTxStyle == nil {
		_dge.TcTxStyle = _cac.TcTxStyle
	} else {
		_dge.TcTxStyle = _ffgg(_dge.TcTxStyle, _cac.TcTxStyle)
	}
	if _dge.TcStyle == nil {
		_dge.TcStyle = _cac.TcStyle
	} else {
		_dge.TcStyle = _fbbd(_dge.TcStyle, _cac.TcStyle)
	}
	return _dge
}
func _aba(_gbfc int, _efddf bool) string {
	_faff := (_gbfc-1)/26 + 1
	_ceb := byte((_gbfc - 1) % 26)
	if _efddf {
		_ceb += byte(65)
	} else {
		_ceb += byte(97)
	}
	_beg := _fe.NewBuffer([]byte{})
	for _cbgf := 0; _cbgf < _faff; _cbgf++ {
		_beg.Write([]byte{_ceb})
	}
	return _beg.String()
}
func _gffb(_dccg *_bb.CT_TableCellProperties, _bdfaa *_bb.CT_TableStyleCellStyle, _acde, _afgc, _bcaag, _dacg bool) *_bb.CT_TableCellProperties {
	_bgea := _bb.NewCT_TableCellProperties()
	if _dccg != nil {
		*_bgea = *_dccg
	}
	if _bdfaa == nil {
		return _bgea
	}
	if _eece := _bdfaa.FillRef; _eece != nil {
		_cba := _bb.NewCT_SolidColorFillProperties()
		_cba.ScrgbClr = _eece.ScrgbClr
		_cba.SrgbClr = _eece.SrgbClr
		_cba.HslClr = _eece.HslClr
		_cba.SysClr = _eece.SysClr
		_cba.SchemeClr = _eece.SchemeClr
		_cba.PrstClr = _eece.PrstClr
		_bgea.SolidFill = _cba
	}
	if _bgea.NoFill == nil && _bgea.SolidFill == nil {
		if _edgc := _bdfaa.Fill; _edgc != nil {
			if _bgea.NoFill == nil {
				_bgea.NoFill = _edgc.NoFill
			}
			if _bgea.SolidFill == nil {
				_bgea.SolidFill = _edgc.SolidFill
			}
		}
	}
	if _geba := _bdfaa.TcBdr; _geba != nil {
		if _bgea.LnL == nil {
			var _eff *_bb.CT_ThemeableLineStyle
			if _bcaag {
				_eff = _geba.Left
			} else {
				_eff = _geba.InsideV
			}
			if _eff != nil {
				_bgea.LnL = _eff.Ln
			}
		}
		if _bgea.LnR == nil {
			var _fccd *_bb.CT_ThemeableLineStyle
			if _dacg {
				_fccd = _geba.Right
			} else {
				_fccd = _geba.InsideV
			}
			if _fccd != nil {
				_bgea.LnR = _fccd.Ln
			}
		}
		if _bgea.LnT == nil {
			var _caee *_bb.CT_ThemeableLineStyle
			if _acde {
				_caee = _geba.Top
			} else {
				_caee = _geba.InsideH
			}
			if _caee != nil {
				_bgea.LnT = _caee.Ln
			}
		}
		if _bgea.LnB == nil {
			var _bbfd *_bb.CT_ThemeableLineStyle
			if _afgc {
				_bbfd = _geba.Bottom
			} else {
				_bbfd = _geba.InsideH
			}
			if _bbfd != nil {
				_bgea.LnB = _bbfd.Ln
			}
		}
	}
	return _bgea
}

type line struct {
	_cdac  float64
	_ebebb float64
	_bedc  float64
	_adf   float64
	_gaab  float64
	_gcdg  []*word
}
type symbolStyle struct {
	_gfaf  *string
	_fbfb  *float64
	_bce   *string
	_dgdef *bool
	_adb   *bool
	_ebd   *bool
	_eddbd *bool
	_gcca  *bool
}

func (_bggg *convertContext) getShapesFromSpPr(_gdd *_bb.CT_ShapeProperties, _eaf *_bb.CT_ShapeStyle, _gddb bool) ([]_be.Drawable, float64, float64, float64, float64, _be.Color, bool) {
	_fbgc := []_be.Drawable{}
	var _bdb, _gce, _gaf, _fdf, _dac float64
	var _cef, _dag, _dedb, _cbg _be.Color
	var _daa *_bb.CT_BlipFillProperties
	_ade, _fbf := 1.0, 1.0
	if _eaf != nil {
		_cef, _dag, _cbg = _bggg.getStyleColors(_eaf)
	}
	if _ffgb := _gdd.Ln; _ffgb != nil {
		if _ffgb.NoFill != nil {
			_dedb, _dac = nil, 0
		} else {
			_dedb, _dac, _ade = _bggg.getInfoFromLn(_ffgb)
			if _dedb == nil {
				_dedb = _cbg
			}
		}
	}
	if _gdd.NoFill != nil {
		_dag, _fbf = nil, 0
	} else if _gddb {
		_dag = _bggg._eafd._daecf
		_fbf = _bggg._eafd._fdb
		_daa = _bggg._eafd._fbbbb
	} else if _gegd := _gdd.SolidFill; _gegd != nil {
		_dag, _fbf = _bggg.getColorFromSolidFill(_gegd)
	}
	var _gbg bool
	if _eae := _gdd.Xfrm; _eae != nil {
		_bdb, _gce, _gaf, _fdf = _ge.GetDataFromXfrm(_eae)
		_gbg = true
	}
	if _ddf := _gdd.CustGeom; _ddf != nil {
		_dfde := []_gf.Point{}
		_dff, _dae := 1.0, 1.0
		if _baa := _ddf.PathLst; _baa != nil {
			for _, _efd := range _baa.Path {
				if _efd != nil {
					if _fgga := _efd.WAttr; _fgga != nil {
						_dff = _gaf / _dg.FromEMU(*_fgga)
					}
					if _agd := _efd.HAttr; _agd != nil {
						_dae = _fdf / _dg.FromEMU(*_agd)
					}
					for _, _gee := range _efd.Close {
						if _gee != nil {
						}
					}
					for _, _deee := range _efd.MoveTo {
						if _deee != nil && _deee.Pt != nil {
							_bc, _fge := _gagea(_deee.Pt)
							_dfde = append(_dfde, _gf.Point{X: _bc*_dff + _bdb, Y: _fge*_dae + _gce})
						}
					}
					for _, _bee := range _efd.LnTo {
						if _bee != nil && _bee.Pt != nil {
							_fcc, _bed := _gagea(_bee.Pt)
							_dfde = append(_dfde, _gf.Point{X: _fcc*_dff + _bdb, Y: _bed*_dae + _gce})
						}
					}
					_dcg := _bggg._bcgf.NewPolygon([][]_gf.Point{_dfde})
					_dcg.SetFillColor(_dag)
					_dcg.SetFillOpacity(_fbf)
					_dcg.SetBorderWidth(_dac)
					if _dedb != nil {
						_dcg.SetBorderColor(_dedb)
						_dcg.SetBorderOpacity(_ade)
					}
					_fbgc = append(_fbgc, _dcg)
				}
			}
		}
	}
	if _gbe := _gdd.PrstGeom; _gbe != nil {
		switch _gbe.PrstAttr {
		case _bb.ST_ShapeTypeRect:
			if _daa != nil {
				_ffa := _bggg.getShapeFromBlipFill(_daa, _bdb, _gce, _gaf, _fdf)
				_fbgc = append(_fbgc, _ffa)
			} else {
				_fdaf := _bggg._bcgf.NewRectangle(_bdb, _gce, _gaf, _fdf)
				_fdaf.SetFillColor(_dag)
				_fdaf.SetFillOpacity(_fbf)
				_fdaf.SetBorderWidth(_dac)
				if _dedb != nil {
					_fdaf.SetBorderColor(_dedb)
					_fdaf.SetBorderOpacity(_ade)
				}
				_fbgc = append(_fbgc, _fdaf)
			}
		case _bb.ST_ShapeTypeLine:
			_ffe := _bggg._bcgf.NewLine(_bdb, _gce, _bdb+_gaf, _gce+_fdf)
			_ffe.SetLineWidth(_dac)
			if _dedb != nil {
				_ffe.SetColor(_dedb)
			}
			_fbgc = append(_fbgc, _ffe)
		}
	}
	return _fbgc, _bdb, _gce, _gaf, _fdf, _cef, _gbg
}
func _gded(_baeb, _abfcg *_bb.CT_TextBodyProperties) *_bb.CT_TextBodyProperties {
	_cdbe := _bb.NewCT_TextBodyProperties()
	if _baeb != nil {
		*_cdbe = *_baeb
	}
	if _abfcg == nil {
		return _cdbe
	}
	if _cdbe.RotAttr == nil {
		_cdbe.RotAttr = _abfcg.RotAttr
	}
	if _cdbe.SpcFirstLastParaAttr == nil {
		_cdbe.SpcFirstLastParaAttr = _abfcg.SpcFirstLastParaAttr
	}
	if _cdbe.VertOverflowAttr == _bb.ST_TextVertOverflowTypeUnset {
		_cdbe.VertOverflowAttr = _abfcg.VertOverflowAttr
	}
	if _cdbe.HorzOverflowAttr == _bb.ST_TextHorzOverflowTypeUnset {
		_cdbe.HorzOverflowAttr = _abfcg.HorzOverflowAttr
	}
	if _cdbe.VertAttr == _bb.ST_TextVerticalTypeUnset {
		_cdbe.VertAttr = _abfcg.VertAttr
	}
	if _cdbe.WrapAttr == _bb.ST_TextWrappingTypeUnset {
		_cdbe.WrapAttr = _abfcg.WrapAttr
	}
	if _cdbe.LInsAttr == nil {
		_cdbe.LInsAttr = _abfcg.LInsAttr
	}
	if _cdbe.TInsAttr == nil {
		_cdbe.TInsAttr = _abfcg.TInsAttr
	}
	if _cdbe.RInsAttr == nil {
		_cdbe.RInsAttr = _abfcg.RInsAttr
	}
	if _cdbe.BInsAttr == nil {
		_cdbe.BInsAttr = _abfcg.BInsAttr
	}
	if _cdbe.NumColAttr == nil {
		_cdbe.NumColAttr = _abfcg.NumColAttr
	}
	if _cdbe.SpcColAttr == nil {
		_cdbe.SpcColAttr = _abfcg.SpcColAttr
	}
	if _cdbe.RtlColAttr == nil {
		_cdbe.RtlColAttr = _abfcg.RtlColAttr
	}
	if _cdbe.AnchorAttr == _bb.ST_TextAnchoringTypeUnset {
		_cdbe.AnchorAttr = _abfcg.AnchorAttr
	}
	if _cdbe.AnchorCtrAttr == nil {
		_cdbe.AnchorCtrAttr = _abfcg.AnchorCtrAttr
	}
	if _cdbe.ForceAAAttr == nil {
		_cdbe.ForceAAAttr = _abfcg.ForceAAAttr
	}
	if _cdbe.UprightAttr == nil {
		_cdbe.UprightAttr = _abfcg.UprightAttr
	}
	if _cdbe.CompatLnSpcAttr == nil {
		_cdbe.CompatLnSpcAttr = _abfcg.CompatLnSpcAttr
	}
	if _cdbe.PrstTxWarp == nil {
		_cdbe.PrstTxWarp = _abfcg.PrstTxWarp
	}
	if _cdbe.NoAutofit == nil {
		_cdbe.NoAutofit = _abfcg.NoAutofit
	}
	if _cdbe.NormAutofit == nil {
		_cdbe.NormAutofit = _abfcg.NormAutofit
	}
	if _cdbe.SpAutoFit == nil {
		_cdbe.SpAutoFit = _abfcg.SpAutoFit
	}
	if _cdbe.Scene3d == nil {
		_cdbe.Scene3d = _abfcg.Scene3d
	}
	if _cdbe.Sp3d == nil {
		_cdbe.Sp3d = _abfcg.Sp3d
	}
	if _cdbe.FlatTx == nil {
		_cdbe.FlatTx = _abfcg.FlatTx
	}
	if _cdbe.ExtLst == nil {
		_cdbe.ExtLst = _abfcg.ExtLst
	}
	return _cdbe
}

// FontStyle represents a kind of font styling. It can be FontStyle_Regular, FontStyle_Bold, FontStyle_Italic and FontStyle_BoldItalic.
type FontStyle = _ge.FontStyle
type paragraph struct {
	_fabf  float64
	_eeb   float64
	_ggf   float64
	_defg  float64
	_eaac  float64
	_fbbcg float64
	_cfed  float64
	_egf   float64
	_dcca  _bb.ST_TextAlignType
	_caff  float64
	_geda  bool
	_ccag  []*line
}

func _fcae(_cafc *_bb.CT_TextCharacterProperties, _cfba *_bb.CT_TableStyleTextStyle) *_bb.CT_TextCharacterProperties {
	_ddfe := _bb.NewCT_TextCharacterProperties()
	if _cafc != nil {
		*_ddfe = *_cafc
	}
	if _cfba == nil {
		return _ddfe
	}
	if _ddfe.BAttr == nil && _cfba.BAttr != _bb.ST_OnOffStyleTypeUnset {
		_efdaf := _cfba.BAttr == _bb.ST_OnOffStyleTypeOn
		_ddfe.BAttr = &_efdaf
	}
	if _ddfe.IAttr == nil && _cfba.IAttr != _bb.ST_OnOffStyleTypeUnset {
		_aaga := _cfba.IAttr == _bb.ST_OnOffStyleTypeOn
		_ddfe.IAttr = &_aaga
	}
	if _ddfe.NoFill == nil && _ddfe.SolidFill == nil {
		_ddfe.SolidFill = _bb.NewCT_SolidColorFillProperties()
		_ddfe.SolidFill.ScrgbClr = _cfba.ScrgbClr
		_ddfe.SolidFill.SrgbClr = _cfba.SrgbClr
		_ddfe.SolidFill.HslClr = _cfba.HslClr
		_ddfe.SolidFill.SysClr = _cfba.SysClr
		_ddfe.SolidFill.SchemeClr = _cfba.SchemeClr
		_ddfe.SolidFill.PrstClr = _cfba.PrstClr
	}
	if _egbd := _cfba.Font; _egbd != nil && _ddfe.Latin == nil && _ddfe.Ea == nil && _ddfe.Cs == nil {
		_ddfe.Latin = _egbd.Latin
		_ddfe.Ea = _egbd.Ea
		_ddfe.Cs = _egbd.Cs
	}
	return _ddfe
}
func (_cga *textboxContext) adjustHeights(_edce float64) {
	if _cga._cfff._gaab < _edce {
		_cga._ccdc._egf += (_edce - _cga._cfff._gaab)
		_cga._cfff._gaab = _edce
	}
}
func (_bcad *convertContext) getColorFromMatrixReference(_eaaf *_bb.CT_StyleMatrixReference) _be.Color {
	if _eaaf == nil {
		return nil
	}
	var _fgc _be.Color
	var _cdfe string
	if _beca := _eaaf.SrgbClr; _beca != nil {
		_cdfe = _beca.ValAttr
	} else if _fcfa := _eaaf.SchemeClr; _fcfa != nil {
		_cdfe = _ge.GetColorStringFromDmlColor(_bcad._bbbb.GetColorBySchemeColor(_fcfa.ValAttr))
		_cdfe = _ge.AdjustColor(_cdfe, _fcfa.EG_ColorTransform)
	}
	if _cdfe != "" {
		_fgc = _be.ColorRGBFromHex("\u0023" + _cdfe)
	}
	return _fgc
}
func (_af *convertContext) makeSlide() {
	_dcc := _af._feb.X().CSld
	if _dcc == nil {
		return
	}
	_egcd := &background{}
	if _ebe := _dcc.Bg; _ebe != nil {
		if _fga := _ebe.BgPr; _fga != nil {
			if _fga.NoFill == nil {
				if _ecd := _fga.SolidFill; _ecd != nil {
					_faf, _bd := _af.getColorFromSolidFill(_ecd)
					if _faf != nil {
						_egcd._daecf = _faf
						_egcd._fdb = _bd
					}
				} else if _cbb := _fga.BlipFill; _cbb != nil {
					_egcd._fbbbb = _cbb
				}
			}
		}
	}
	_af._eafd = _egcd
	if _gb := _dcc.SpTree; _gb != nil {
		for _, _ga := range _gb.Choice {
			if _ga != nil {
				for _, _aeb := range _ga.Sp {
					if _aeb != nil {
						_eed := _af.getShapes(_aeb)
						_af._cbcb = append(_af._cbcb, _eed...)
					}
				}
				for _, _bf := range _ga.GraphicFrame {
					if _bf != nil {
						var _dgd, _aagd, _ega, _cda float64
						if _abc := _bf.Xfrm; _abc != nil {
							_dgd, _aagd, _ega, _cda = _ge.GetDataFromXfrm(_abc)
						}
						if _cab := _bf.Graphic; _cab != nil {
							if _fbd := _cab.GraphicData; _fbd != nil {
								for _, _bdg := range _fbd.Any {
									if _fce, _gfg := _bdg.(*_cbf.Chart); _gfg {
										_aagg, _gea := _af.makePdfBlockFromChart(_fce, _ega, _cda)
										if _gea != nil {
											_cd.Log.Debug("C\u0061\u006e\u006e\u006ft \u0072e\u0061\u0064\u0020\u0062\u006co\u0063\u006b\u003a\u0020\u0025\u0073", _gea)
										}
										if _aagg != nil {
											_aagg.SetPos(_dgd, _aagd)
											_af._cbcb = append(_af._cbcb, _aagg)
										}
									} else if _fdg, _bdc := _bdg.(*_bb.Tbl); _bdc {
										_gcb := _af.makePdfBlockFromTable(_fdg, _ega)
										if _gcb != nil {
											_efb := _be.NewBlock(_ega, _cda)
											_efb.SetPos(_dgd, _aagd)
											_efb.Draw(_gcb)
											_af._cbcb = append(_af._cbcb, _efb)
										}
									}
								}
							}
						}
					}
				}
				for _, _fdd := range _ga.CxnSp {
					if _fdd != nil {
						_gfae := _af.getConnectors(_fdd)
						_af._cbcb = append(_af._cbcb, _gfae...)
					}
				}
				for _, _edg := range _ga.Pic {
					if _edg != nil {
						var _bgb, _gab, _fcb, _fec float64
						if _dgde := _edg.SpPr; _dgde != nil {
							if _afc := _dgde.Xfrm; _afc != nil {
								_bgb, _gab, _fcb, _fec = _ge.GetDataFromXfrm(_afc)
							}
						}
						if _bba := _edg.BlipFill; _bba != nil {
							_fbb := _af.getShapeFromBlipFill(_bba, _bgb, _gab, _fcb, _fec)
							_af._cbcb = append(_af._cbcb, _fbb)
						}
					}
				}
			}
		}
	}
}
func (_feg *convertContext) getConnectors(_afe *_dd.CT_Connector) []_be.Drawable {
	_ebea, _, _, _, _, _, _ := _feg.getShapesFromSpPr(_afe.SpPr, _afe.Style, false)
	return _ebea
}
func _ffgg(_gdcc, _gbfe *_bb.CT_TableStyleTextStyle) *_bb.CT_TableStyleTextStyle {
	_bfd := _bb.NewCT_TableStyleTextStyle()
	if _gdcc != nil {
		*_bfd = *_gdcc
	}
	if _gbfe == nil {
		return _bfd
	}
	if _bfd.BAttr == _bb.ST_OnOffStyleTypeUnset {
		_bfd.BAttr = _gbfe.BAttr
	}
	if _bfd.IAttr == _bb.ST_OnOffStyleTypeUnset {
		_bfd.IAttr = _gbfe.IAttr
	}
	if _bfd.Font == nil {
		_bfd.Font = _gbfe.Font
	}
	if _bfd.FontRef == nil {
		_bfd.FontRef = _gbfe.FontRef
	}
	if _bfd.ScrgbClr == nil {
		_bfd.ScrgbClr = _gbfe.ScrgbClr
	}
	if _bfd.SrgbClr == nil {
		_bfd.SrgbClr = _gbfe.SrgbClr
	}
	if _bfd.HslClr == nil {
		_bfd.HslClr = _gbfe.HslClr
	}
	if _bfd.SysClr == nil {
		_bfd.SysClr = _gbfe.SysClr
	}
	if _bfd.SchemeClr == nil {
		_bfd.SchemeClr = _gbfe.SchemeClr
	}
	if _bfd.PrstClr == nil {
		_bfd.PrstClr = _gbfe.PrstClr
	}
	return _bfd
}
func (_efgb *convertContext) renderPageWithDrawableToGoImage(_fgd _be.Drawable) (_c.Image, error) {
	_efcb := _ge.MakeTempCreator(_efgb._aed, _efgb._abcea)
	_efcb.NewPage()
	_efcb.Draw(_fgd)
	_feea, _bbed := _ge.GetPageFromCreator(_efcb)
	if _bbed != nil {
		return nil, _bbed
	}
	return _ebf.NewImageDevice().Render(_feea)
}

type symbol struct {
	_ffd   string
	_gde   float64
	_bcac  float64
	_bbgg  float64
	_gcff  float64
	_fbe   float64
	_cdccd *_be.TextStyle
	_fbag  string
	_aaae  bool
	_fccf  bool
	_dfb   bool
}

func (_bcgb *convertContext) makePdfBlockFromTable(_aggg *_bb.Tbl, _gffa float64) *_be.Table {
	_afb := _aggg.TblGrid
	if _afb == nil {
		return nil
	}
	_ege := len(_afb.GridCol)
	if _ege == 0 {
		return nil
	}
	_edc := []float64{}
	_fgf := 0.0
	for _, _ecb := range _afb.GridCol {
		_abg := _dg.FromEMU(_ge.FromSTCoordinate(_ecb.WAttr))
		_edc = append(_edc, _abg)
		_fgf += _abg
	}
	_ede := []float64{}
	for _ggag := 0; _ggag < _ege; _ggag++ {
		_ede = append(_ede, _edc[_ggag]/_fgf)
	}
	_efdg := _bcgb._bcgf.NewTable(_ege)
	_efdg.SetColumnWidths(_ede...)
	_ecbf := _aggg.TblPr
	var _ccgf *_bb.CT_TableStyle
	if _gcbe := _ecbf.Choice; _gcbe != nil {
		if _gcbe.TableStyle != nil {
			_ccgf = _gcbe.TableStyle
		} else if _gcbe.TableStyleId != nil {
			_ccgf = _bcgb._bbbb.GetTableStyleById(*_gcbe.TableStyleId)
		}
	}
	_dedbe := _bb.NewCT_TablePartStyle()
	_dedbe.TcStyle = _bb.NewCT_TableStyleCellStyle()
	_dedbe.TcTxStyle = _bb.NewCT_TableStyleTextStyle()
	if _ccgf != nil {
		if _ccgf.WholeTbl != nil {
			*_dedbe = *_ccgf.WholeTbl
		}
		if _ccgf.TblBg != nil {
			if _dedbe.TcStyle.Fill == nil {
				_dedbe.TcStyle.Fill = _ccgf.TblBg.Fill
			}
		}
	}
	if _dedbe.TcStyle.Fill == nil {
		_dedbe.TcStyle.Fill = _bb.NewCT_FillProperties()
		_dedbe.TcStyle.Fill.NoFill = _ecbf.NoFill
		_dedbe.TcStyle.Fill.SolidFill = _ecbf.SolidFill
	}
	_bfe := len(_aggg.Tr)
	for _ebbb, _bac := range _aggg.Tr {
		_gbd := _ebbb == 0
		_dcf := _ebbb == _bfe-1
		_acbe := _ebbb%2 == 0
		_ebef := len(_bac.Tc)
		var _ccac *_bb.CT_TablePartStyle
		if _gbd {
			_ccac = _ccgf.FirstRow
		} else if _acbe {
			_ccac = _ccgf.Band2H
		} else {
			_ccac = _ccgf.Band1H
		}
		var _ffgd float64
		for _ead, _ffee := range _bac.Tc {
			_bca := _ead == 0
			_bgc := _ead == _ebef-1
			_aae := _ead%2 == 0
			var _fcf *_bb.CT_TablePartStyle
			if _bca {
				_fcf = _ccgf.FirstCol
			} else if _aae {
				_fcf = _ccgf.Band2V
			} else {
				_fcf = _ccgf.Band1V
			}
			_ebeb := _bbbf(_bbbf(_fcf, _ccac), _dedbe)
			_fdgf := _bcgb.addCellToTable(_efdg, _ffee, _ebeb, _gffa*_ede[_ead], _gbd, _dcf, _bca, _bgc)
			if _fdgf > _ffgd {
				_ffgd = _fdgf
			}
		}
		_gac := _dg.FromEMU(_ge.FromSTCoordinate(_bac.HAttr))
		if _gac < _ffgd {
			_gac = _ffgd
		}
		if _gac < _eeaa(4) {
			_gac = _eeaa(4)
		}
		_efdg.SetRowHeight(_efdg.CurRow(), _gac)
	}
	return _efdg
}
func (_cdaa *convertContext) getPhData(_eceb *_dd.CT_Shape) (*_bb.CT_Transform2D, *_bb.CT_TextBodyProperties, *_bb.CT_TextListStyle, bool, bool) {
	_fdfd, _egbc := _abgba(_eceb)
	_cged := _egbc == nil
	_eegd, _gfgga, _dcee, _bfca := _afbd(_cdaa._daec.CSld, _fdfd, _egbc, _cged)
	_dcac, _cfcf, _fbdd, _ecdg := _afbd(_cdaa._cge.CSld, _fdfd, _egbc, _cged)
	if _dcac == nil {
		_dcac = _eegd
	}
	_eeba, _cbge := _ebge(_cfcf, _gfgga)
	var _fegaa, _ebadc bool
	if _fbdd == nil {
		if _dcee != nil {
			_fegaa = *_dcee
		}
	} else {
		_fegaa = *_fbdd
	}
	if _ecdg == nil {
		if _bfca != nil {
			_ebadc = *_bfca
		}
	} else {
		_ebadc = *_ecdg
	}
	return _dcac, _eeba, _cbge, _fegaa, _ebadc
}
func (_ebb *convertContext) makePdfBlockFromChart(_dde *_cbf.Chart, _gc, _a float64) (*_be.Block, error) {
	_ec := _dde.CT_RelId.IdAttr
	_bec := _ebb._feb.GetChartSpaceByRelId(_ec)
	if _bec == nil {
		return nil, _f.New("\u004e\u006f\u0020\u0063\u0068\u0061\u0072\u0074\u0073\u0070\u0061\u0063\u0065")
	}
	var _eg *_bb.Theme
	_ff := _ebb._bbbb.Themes()
	if len(_ff) > 0 {
		_eg = _ff[0]
	}
	return _ge.MakeBlockFromChartSpace(_bec, _gc, _a, _eg)
}
func (_edgg *textboxContext) alignVertically(_cebg _bb.ST_TextAnchoringType) {
	_edgg.alignParagraphsVertically(_cebg)
	_edgg.alignSymbolsVertically()
}
func (_gbfab *textboxContext) assignPropsToCurrentParagraph(_feeb *_bb.CT_TextParagraphProperties) {
	_cdcab := 12.4
	if _feeb == nil {
		_gbfab._ccdc._caff = _cdcab
		return
	}
	if _aafe := _feeb.DefRPr; _aafe != nil {
		_agfe := _aafe.SzAttr
		if _agfe != nil {
			_bdga := float64(*_agfe) / 1200
			if _cdcab <= _bdga {
				_cdcab = _bdga
			}
		}
	}
	if _bgfe := _feeb.MarLAttr; _bgfe != nil {
		_gbfab._ccdc._eaac = _dg.FromEMU(int64(*_bgfe))
	}
	_gbfab._ccdc._fbbcg = _gbfab._gggc
	if _ecg := _feeb.MarRAttr; _ecg != nil {
		_gbfab._ccdc._fbbcg -= _dg.FromEMU(int64(*_ecg))
	}
	if _fbff := _feeb.IndentAttr; _fbff != nil {
		_gbfab._ccdc._defg = _dg.FromEMU(int64(*_fbff))
	}
	if _gccd := _feeb.LatinLnBrkAttr; _gccd != nil {
		_gbfab._ccdc._geda = *_gccd
	}
	if _gffae := _feeb.LnSpc; _gffae != nil {
		if _cagg := _gffae.SpcPct; _cagg != nil {
			if _adbf := _cagg.ValAttr.ST_TextSpacingPercent; _adbf != nil {
				_cdcab = float64(*_adbf) / 5000
			}
		}
	}
	var _bcf float64
	if _gefc := _feeb.SpcBef; _gefc != nil {
		if _caag := _gefc.SpcPts; _caag != nil {
			_bcf = float64(_caag.ValAttr) / 100
		}
	}
	_dcb := _gbfab._edb
	if len(_dcb) > 0 {
		_bcf -= _dcb[len(_dcb)-1]._ggf
		if _bcf < 0 {
			_bcf = 0
		}
	}
	_gbfab._ccdc._eeb = _bcf
	if _cbfc := _feeb.SpcAft; _cbfc != nil {
		if _cgg := _cbfc.SpcPts; _cgg != nil {
			_gbfab._ccdc._ggf = float64(_cgg.ValAttr) / 100
		}
	}
	_gbfab._ccdc._caff = _cdcab
	_gbfab._ccdc._dcca = _feeb.AlgnAttr
}

const (
	FontStyle_Regular    FontStyle = 0
	FontStyle_Bold       FontStyle = 1
	FontStyle_Italic     FontStyle = 2
	FontStyle_BoldItalic FontStyle = 3
)
