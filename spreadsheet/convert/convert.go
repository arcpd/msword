package convert

import (
	_a "github.com/arcpd/msword/common/logger"
	_cc "github.com/arcpd/msword/common/tempstorage"
	_bd "github.com/arcpd/msword/internal/convertutils"
	_g "github.com/arcpd/msword/measurement"
	_e "github.com/arcpd/msword/schema/soo/dml"
	_eb "github.com/arcpd/msword/schema/soo/dml/chart"
	_fa "github.com/arcpd/msword/schema/soo/ofc/sharedTypes"
	_cd "github.com/arcpd/msword/schema/soo/sml"
	_be "github.com/arcpd/msword/spreadsheet"
	_bg "github.com/arcpd/msword/spreadsheet/reference"
	_ag "github.com/unidoc/unipdf/v3/creator"
	_f "github.com/unidoc/unipdf/v3/model"
	_ce "image"
	_b "strconv"
)

func (_baaa *convertContext) drawSheet() {
	for _abfg, _eea := range _baaa._gcbg {
		_ebb := len(_eea._cdafb)
		if _abfg == len(_baaa._gcbg)-1 {
			for _eed := len(_eea._cdafb) - 1; _eed >= 0; _eed-- {
				if !_eea._cdafb[_eed]._bee {
					_ebb = _eed
				}
			}
		}
		_ccg := _eea._cdafb[:_ebb]
		for _, _egae := range _ccg {
			_baaa._ffdc.NewPage()
			_baaa.drawPage(_egae)
		}
	}
}
func (_baca *convertContext) getImage(_fggg _ce.Image, _gfb, _baba, _gacc, _aead, _dad, _cabg float64, _abbc _bd.ImgPart) *_ag.Image {
	_aead += _baca._gda
	_gacc += _baca._fgb
	_dff, _gcc := _bd.GetImage(_baca._ffdc, _fggg, _gfb, _baba, _gacc, _aead, _dad, _cabg, _abbc)
	if _gcc != nil {
		_a.Log.Debug("\u0043\u0061\u006eno\u0074\u0020\u0067\u0065\u0074\u0020\u0061\u006e\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _gcc)
		return nil
	}
	return _dff
}
func (_fea *convertContext) getSymbolsFromR(_fec []*_cd.CT_RElt, _ccec *style) []*symbol {
	_bgcd := []*symbol{}
	for _, _fdg := range _fec {
		_cfgf := _fea.combineCellStyleWithRPrElt(_ccec, _fdg.RPr)
		for _, _dfb := range _fdg.T {
			_bgcd = append(_bgcd, &symbol{_gbb: string(_dfb), _acbd: _fea.makeTextStyleFromCellStyle(_cfgf)})
		}
	}
	return _bgcd
}

type colWidthRange struct {
	_gfgb int
	_dcca int
	_fgd  float64
	_fade *style
}

func (_fee *convertContext) alignSymbolsVertically(_cbdf *cell, _gaac _cd.ST_VerticalAlignment) {
	var _ffeb float64
	switch _gaac {
	case _cd.ST_VerticalAlignmentTop:
		_ffeb = _ga
		if _cbdf._fbed {
			_ffeb -= _ed
		} else if _cbdf._efd {
			_ffeb += 4 * _ed
		}
		for _, _bgac := range _cbdf._fag {
			_ffeb += _bgac._dcae
			_bgac._cgc = _ffeb
			_ffeb += _cg
		}
	case _cd.ST_VerticalAlignmentCenter:
		_bgba := 0.0
		for _, _aggg := range _cbdf._fag {
			_bgba += _aggg._dcae + _faabg(1)
		}
		_ffeb = 0.5 * (_cbdf._ebbd - _bgba)
		if _cbdf._fbed {
			_ffeb -= 2 * _ed
		} else if _cbdf._efd {
			_ffeb += 2 * _ed
		}
		for _, _eag := range _cbdf._fag {
			_ffeb += _eag._dcae + 0.5*_cg
			_eag._cgc = _ffeb
			_ffeb += 0.5 * _cg
		}
	default:
		_ffeb = _cbdf._ebbd - _ga
		if _cbdf._fbed {
			_ffeb -= 4 * _ed
		} else if _cbdf._efd {
			_ffeb += _ed
		}
		for _gdegg := len(_cbdf._fag) - 1; _gdegg >= 0; _gdegg-- {
			_cbdf._fag[_gdegg]._cgc = _ffeb
			_ffeb -= _cbdf._fag[_gdegg]._dcae
			_ffeb -= _cg
		}
	}
}
func (_gec *convertContext) makeAnchors() {
	_bcc, _ba := _gec._eedf.GetDrawing()
	if _bcc != nil {
		for _, _dc := range _bcc.EG_Anchor {
			_ad := &anchor{}
			if _gea := _dc.TwoCellAnchor; _gea != nil {
				_ac, _dca := _gea.From, _gea.To
				if _ac == nil || _dca == nil {
					return
				}
				_ad._gcf = int(_ac.Row)
				_ad._fcf = _bd.FromSTCoordinate(_ac.RowOff)
				_ad._feg = int(_ac.Col)
				_ad._badd = _bd.FromSTCoordinate(_ac.ColOff)
				_ad._gddd = int(_dca.Row)
				_ad._ccf = _bd.FromSTCoordinate(_dca.RowOff)
				_ad._dfbe = int(_dca.Col)
				_ad._gga = _bd.FromSTCoordinate(_dca.ColOff)
				if _efb := _gea.Choice; _efb != nil {
					if _ege := _efb.Pic; _ege != nil {
						if _ae := _ege.BlipFill; _ae != nil {
							if _aca := _ae.Blip; _aca != nil {
								if _bb := _aca.EmbedAttr; _bb != nil {
									for _, _gad := range _ba.X().Relationship {
										if _gad.IdAttr == *_bb {
											for _, _gd := range _gec._cceg.Images {
												if _gd.Target() == _gad.TargetAttr {
													_bce, _fd := _cc.Open(_gd.Path())
													if _fd != nil {
														_a.Log.Debug("\u004fp\u0065\u006e\u0020\u0069m\u0061\u0067\u0065\u0020\u0066i\u006ce\u0020e\u0072\u0072\u006f\u0072\u003a\u0020\u0025s", _fd)
														continue
													}
													_de, _, _fd := _ce.Decode(_bce)
													if _fd != nil {
														_a.Log.Debug("\u0044\u0065\u0063\u006fde\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020%\u0073", _fd)
														continue
													}
													_ad._fdbb = _de
												}
											}
										}
									}
								}
							}
						}
					} else if _fb := _efb.GraphicFrame; _fb != nil {
						if _gbe := _fb.Graphic; _gbe != nil {
							if _dee := _gbe.GraphicData; _dee != nil {
								for _, _cag := range _dee.Any {
									if _bcd, _dec := _cag.(*_eb.Chart); _dec {
										for _, _ee := range _ba.X().Relationship {
											if _ee.IdAttr == _bcd.IdAttr {
												_gde := _gec._cceg.GetChartByTargetId(_ee.TargetAttr)
												if _gde != nil {
													_ad._dcgg = _gde
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
			if _ad._fdbb != nil || _ad._dcgg != nil {
				_gec._cfff = append(_gec._cfff, _ad)
			}
		}
	}
}

var _cg = _faabg(1)

func (_bebd *convertContext) makeTextStyleFromCellStyle(_ggac *style) *_ag.TextStyle {
	_dag := _bebd._ffdc.NewTextStyle()
	if _ggac == nil {
		_dag.FontSize = _bd.DefaultFontSize
		_dag.Font = _bd.AssignStdFontByName(_dag, _bd.StdFontsMap["\u0064e\u0066\u0061\u0075\u006c\u0074"][FontStyle_Regular])
		return &_dag
	}
	if _dae(_ggac._aeae) {
		_dag.Underline = true
		_dag.UnderlineStyle = _ag.TextDecorationLineStyle{Offset: 0.5, Thickness: _faabg(1 / 32)}
	}
	var _bfg FontStyle
	if _dae(_ggac._bcfb) && _dae(_ggac._fbd) {
		_bfg = FontStyle_BoldItalic
	} else if _dae(_ggac._bcfb) {
		_bfg = FontStyle_Bold
	} else if _dae(_ggac._fbd) {
		_bfg = FontStyle_Italic
	} else {
		_bfg = FontStyle_Regular
	}
	_faga := "\u0064e\u0066\u0061\u0075\u006c\u0074"
	if _ggac._acg != nil {
		_faga = *_ggac._acg
	}
	if _bbbbc, _dcaf := _bd.StdFontsMap[_faga]; _dcaf {
		_dag.Font = _bd.AssignStdFontByName(_dag, _bbbbc[_bfg])
	} else if _ffbe := _bd.GetRegisteredFont(_faga, _bfg); _ffbe != nil {
		_dag.Font = _ffbe
	} else {
		_a.Log.Debug("\u0046\u006f\u006e\u0074\u0020\u0025\u0073\u0020\u0077\u0069\u0074h\u0020\u0073\u0074\u0079\u006c\u0065\u0020\u0025s\u0020i\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002c\u0020\u0072\u0065\u0073\u0065\u0074 \u0074\u006f\u0020\u0064\u0065\u0066\u0061\u0075\u006c\u0074\u002e", _faga, _bfg)
		_dag.Font = _bd.AssignStdFontByName(_dag, _bd.StdFontsMap["\u0064e\u0066\u0061\u0075\u006c\u0074"][_bfg])
	}
	if _ggac._edgc != nil {
		_dag.FontSize = *_ggac._edgc
	}
	if _ggac._fdge != nil {
		_dag.Color = _ag.ColorRGBFromHex(*_ggac._fdge)
	}
	if _ggac._dcab != nil && *_ggac._dcab {
		_dag.FontSize *= _ab
	} else if _ggac._efgb != nil && *_ggac._efgb {
		_dag.FontSize *= _ab
	}
	return &_dag
}
func (_gaf *convertContext) makeCols() {
	_fbe := _gaf._eedf
	_bde := _fbe.X()
	_abd := []*colInfo{}
	_gaa := []colWidthRange{}
	if _caa := _bde.Cols; len(_caa) > 0 {
		for _, _cae := range _caa[0].Col {
			_ea := 65.0
			if _da := _cae.WidthAttr; _da != nil {
				if *_da > 0.83 {
					*_da -= 0.83
				}
				if *_da <= 1 {
					_ea = *_da * 11
				} else {
					_ea = 5 + *_da*6
				}
			}
			_aea := int(_cae.MinAttr - 1)
			_gag := int(_cae.MaxAttr - 1)
			_gaa = append(_gaa, colWidthRange{_gfgb: _aea, _dcca: _gag, _fgd: _ea, _fade: _gaf.getStyle(_cae.StyleAttr)})
		}
	}
	_gbef := 0
	for _fcc := 0; _fcc <= _gaf._fbgd; _fcc++ {
		var _bgb float64
		var _cea *style
		if _gbef >= len(_gaa) {
			_bgb = 65
		} else {
			_agg := _gaa[_gbef]
			if _fcc >= _agg._gfgb && _fcc <= _agg._dcca {
				_bgb = _agg._fgd
				_cea = _agg._fade
				if _fcc == _agg._dcca {
					_gbef++
				}
			} else {
				_bgb = 65
			}
		}
		_abd = append(_abd, &colInfo{_gfc: _bgb, _ebag: _cea})
	}
	_gaf._cdef = _abd
}
func (_cdcf *convertContext) alignSymbolsHorizontally(_bff *cell, _fdeg _cd.ST_HorizontalAlignment) {
	if _fdeg == _cd.ST_HorizontalAlignmentUnset {
		switch _bff._afg {
		case _cd.ST_CellTypeB:
			_fdeg = _cd.ST_HorizontalAlignmentCenter
		case _cd.ST_CellTypeN:
			_fdeg = _cd.ST_HorizontalAlignmentRight
		default:
			_fdeg = _cd.ST_HorizontalAlignmentLeft
		}
	}
	var _cfe float64
	for _, _cecb := range _bff._fag {
		switch _fdeg {
		case _cd.ST_HorizontalAlignmentLeft:
			_cfe = _ef
		case _cd.ST_HorizontalAlignmentRight:
			_aec := _ddbd(_cecb._acdd)
			_cfe = _bff._faedb - _ef - _aec
		case _cd.ST_HorizontalAlignmentCenter:
			_ceb := _ddbd(_cecb._acdd)
			_cfe = (_bff._faedb - _ceb) / 2
		}
		for _, _aag := range _cecb._acdd {
			_aag._agab += _cfe
		}
	}
}
func (_ebf *convertContext) makePagespans() {
	_ebf._gcbg = []*pagespan{}
	_edg := 0.0
	_fgc := 0
	for _fbg, _ebfg := range _ebf._cdef {
		_babd := _ebfg._gfc
		if _edg+_babd <= _ebf._gcg {
			_ebfg._faed = _edg
			_edg += _babd
		} else {
			_ebfg._faed = 0
			_ebf._gcbg = append(_ebf._gcbg, &pagespan{_ccdf: _edg, _agea: _fgc, _dged: _fbg})
			_edg = _babd
			_fgc = _fbg
		}
	}
	_ebf._gcbg = append(_ebf._gcbg, &pagespan{_ccdf: _edg, _agea: _fgc, _dged: len(_ebf._cdef)})
}
func (_edgf *convertContext) addRowToPage(_cca []*cell, _eaeb int) {
	_bagd := 0.0
	_ffef := _edgf._gcg
	for _, _ggf := range _cca {
		if len(_ggf._fag) != 0 {
			_ggf._bcae = _bagd
			_bagd = _ggf._badg + _ggf._faedb
		}
	}
	for _ccc := len(_cca) - 1; _ccc >= 0; _ccc-- {
		_afc := _cca[_ccc]
		if len(_afc._fag) != 0 {
			_afc._gbccb = _ffef
			_ffef = _afc._badg
		}
	}
	_edgf._cadb._ccab = append(_edgf._cadb._ccab, &pageRow{_ecca: _eaeb, _gcd: _cca})
}

type mergedCell struct {
	_dgb  uint32
	_eaf  uint32
	_geca uint32
	_ecd  uint32
	_cbe  float64
	_ebe  float64
}

func (_cf *convertContext) determineMaxIndexes() {
	var _bdd, _acd int
	_bdd = int(_cf._eedf.MaxColumnIdx())
	_bbf := _cf._eedf.Rows()
	if len(_bbf) > 0 {
		_acd = int(_bbf[len(_bbf)-1].RowNumber())
	}
	for _, _ega := range _cf._cfff {
		if _ega._gddd >= _acd {
			_acd = _ega._gddd + 1
		}
		if _ega._dfbe >= _bdd {
			_bdd = _ega._dfbe + 1
		}
	}
	_cf._cba = _acd
	_cf._fbgd = _bdd
}
func _dae(_ggd *bool) bool { return _ggd != nil && *_ggd }
func (_bbd *convertContext) distributeAnchors() {
	for _, _abg := range _bbd._cfff {
		_acb, _aff := _abg._gcf, _abg._fcf
		_fdf, _fbce := _abg._feg, _abg._badd
		_efe, _agfe := _abg._gddd, _abg._ccf
		_gg, _cfg := _abg._dfbe, _abg._gga
		var _dbac, _aaa, _daa, _cfa *page
		for _, _dce := range _bbd._gcbg {
			for _, _eecc := range _dce._cdafb {
				if _acb >= _eecc._bgbd._eegg && _acb < _eecc._bgbd._eadg {
					if _fdf >= _eecc._dfg._agea && _fdf < _eecc._dfg._dged {
						_eecc._bee = true
						_dbac = _eecc
					}
					if _gg >= _eecc._dfg._agea && _gg < _eecc._dfg._dged {
						_eecc._bee = true
						_aaa = _eecc
					}
				}
				if _efe >= _eecc._bgbd._eegg && _efe < _eecc._bgbd._eadg {
					if _fdf >= _eecc._dfg._agea && _fdf < _eecc._dfg._dged {
						_eecc._bee = true
						_cfa = _eecc
					}
					if _gg >= _eecc._dfg._agea && _gg < _eecc._dfg._dged {
						_eecc._bee = true
						_daa = _eecc
					}
				}
			}
		}
		_efbd := _dbac != _aaa
		_debg := _dbac != _cfa
		if _efbd && _debg {
			_faf := _bbd._cdef[_fdf]._faed + _g.FromEMU(_fbce)
			_adc := _dbac._dfg._ccdf
			_bbbb := _bbd._cdef[_gg]._faed + _g.FromEMU(_cfg)
			_cbb := _bbd._fadb[_acb]._dcda + _g.FromEMU(_aff)
			_gbed := float64(_dbac._bgbd._cbba)
			_ceg := _bbd._fadb[_efe]._dcda + _g.FromEMU(_agfe)
			_cfc := _bbbb + _adc - _faf
			_gfa := _ceg + _gbed - _cbb
			_bga := _bbd.imageFromAnchor(_abg, _cfc, _gfa)
			_dbac._fbfe = append(_dbac._fbfe, _bbd.getImage(_bga, _gfa, _cfc, _faf, _cbb, _adc-_faf, _gbed-_cbb, _bd.ImgPart_lt))
			_aaa._fbfe = append(_aaa._fbfe, _bbd.getImage(_bga, _gfa, _cfc, 0, _cbb, _adc-_faf, _gbed-_cbb, _bd.ImgPart_rt))
			_cfa._fbfe = append(_cfa._fbfe, _bbd.getImage(_bga, _gfa, _cfc, _faf, 0, _adc-_faf, _gbed-_cbb, _bd.ImgPart_lb))
			_daa._fbfe = append(_daa._fbfe, _bbd.getImage(_bga, _gfa, _cfc, 0, 0, _adc-_faf, _gbed-_cbb, _bd.ImgPart_rb))
		} else if _efbd {
			_afe := _bbd._fadb[_acb]._dcda + _g.FromEMU(_aff)
			_fgg := _bbd._fadb[_efe]._dcda + _g.FromEMU(_agfe)
			_dbab := _bbd._cdef[_fdf]._faed + _g.FromEMU(_fbce)
			_cdg := _dbac._dfg._ccdf
			_ccea := _bbd._cdef[_gg]._faed + _g.FromEMU(_cfg)
			_bcbf := _ccea + _cdg - _dbab
			_cbbb := _fgg - _afe
			_debga := _bbd.imageFromAnchor(_abg, _bcbf, _cbbb)
			_dbac._fbfe = append(_dbac._fbfe, _bbd.getImage(_debga, _cbbb, _bcbf, _dbab, _afe, _cdg-_dbab, 0, _bd.ImgPart_l))
			_aaa._fbfe = append(_aaa._fbfe, _bbd.getImage(_debga, _cbbb, _bcbf, 0, _afe, _cdg-_dbab, 0, _bd.ImgPart_r))
		} else if _debg {
			_dbff := _bbd._cdef[_fdf]._faed + _g.FromEMU(_fbce)
			_baa := _bbd._cdef[_gg]._faed + _g.FromEMU(_cfg)
			_abfb := _bbd._fadb[_acb]._dcda + _g.FromEMU(_aff)
			_egf := float64(_dbac._bgbd._cbba)
			_fad := _bbd._fadb[_efe]._dcda + _g.FromEMU(_agfe)
			_fggc := _baa - _dbff
			_dcec := _fad + _egf - _abfb
			_gdd := _bbd.imageFromAnchor(_abg, _fggc, _dcec)
			_dbac._fbfe = append(_dbac._fbfe, _bbd.getImage(_gdd, _dcec, _fggc, _dbff, _abfb, 0, _egf-_abfb, _bd.ImgPart_t))
			_cfa._fbfe = append(_cfa._fbfe, _bbd.getImage(_gdd, _dcec, _fggc, _dbff, 0, 0, _egf-_abfb, _bd.ImgPart_b))
		} else {
			_add := _bbd._cdef[_fdf]._faed + _g.FromEMU(_fbce)
			_cdb := _bbd._cdef[_gg]._faed + _g.FromEMU(_cfg)
			_geb := _bbd._fadb[_acb]._dcda + _g.FromEMU(_aff)
			_fdfd := _bbd._fadb[_efe]._dcda + _g.FromEMU(_agfe)
			_cfd := _cdb - _add
			_bbe := _fdfd - _geb
			_baag := _bbd.imageFromAnchor(_abg, _cfd, _bbe)
			_dbac._fbfe = append(_dbac._fbfe, _bbd.getImage(_baag, _bbe, _cfd, _add, _geb, 0, 0, _bd.ImgPart_whole))
		}
	}
}

var _cce = 3.025 / _faabg(1)

func (_bea *convertContext) getBorder(_fgggc *_cd.CT_BorderPr) *border {
	_fab := &border{}
	switch _fgggc.StyleAttr {
	case _cd.ST_BorderStyleThin:
		_fab._ccdd = _eg
	case _cd.ST_BorderStyleMedium:
		_fab._ccdd = _eg * 2
	case _cd.ST_BorderStyleThick:
		_fab._ccdd = _eg * 4
	}
	if _fab._ccdd == 0.0 {
		return nil
	}
	if _edbf := _fgggc.Color; _edbf != nil {
		_bfc := _bea.getColorStringFromSmlColor(_edbf)
		if _bfc != nil {
			_fab._beg = _ag.ColorRGBFromHex(*_bfc)
		} else {
			_fab._beg = _ag.ColorBlack
		}
	}
	return _fab
}
func (_efda *convertContext) getStyle(_dafd *uint32) *style {
	_acad := &style{}
	_ffdg := false
	if _dafd != nil {
		_eage := _efda._gab.GetCellStyle(*_dafd)
		_aefb := _eage.GetFont()
		for _, _egga := range _aefb.Name {
			if _egga != nil {
				_acad._acg = &_egga.ValAttr
				_ffdg = true
				break
			}
		}
		for _, _ede := range _aefb.B {
			if _ede != nil {
				_gca := _ede.ValAttr == nil || *_ede.ValAttr
				_acad._bcfb = &_gca
				_ffdg = true
				break
			}
		}
		for _, _gabb := range _aefb.I {
			if _gabb != nil {
				_deec := _gabb.ValAttr == nil || *_gabb.ValAttr
				_acad._fbd = &_deec
				_ffdg = true
				break
			}
		}
		for _, _cdgc := range _aefb.U {
			if _cdgc != nil {
				_gfcc := _cdgc.ValAttr == _cd.ST_UnderlineValuesSingle || _cdgc.ValAttr == _cd.ST_UnderlineValuesUnset
				_acad._aeae = &_gfcc
				_ffdg = true
				break
			}
		}
		for _, _ebc := range _aefb.Sz {
			if _ebc != nil {
				_beee := _ebc.ValAttr / 12 * _bd.DefaultFontSize
				_acad._edgc = &_beee
				_ffdg = true
				break
			}
		}
		for _, _gegc := range _aefb.VertAlign {
			if _gegc != nil {
				_ggff := _gegc.ValAttr == _fa.ST_VerticalAlignRunSuperscript
				_acad._dcab = &_ggff
				_aagc := _gegc.ValAttr == _fa.ST_VerticalAlignRunSubscript
				_acad._efgb = &_aagc
				_ffdg = true
				break
			}
		}
		for _, _gdec := range _aefb.Color {
			if _gdec != nil {
				_acad._fdge = _efda.getColorStringFromSmlColor(_gdec)
				_ffdg = true
				break
			}
		}
		_fdde := _eage.GetBorder()
		if _fdde.Top != nil {
			_acad._afef = _efda.getBorder(_fdde.Top)
			_ffdg = true
		}
		if _fdde.Bottom != nil {
			_acad._fddg = _efda.getBorder(_fdde.Bottom)
			_ffdg = true
		}
		if _fdde.Left != nil {
			_acad._dfe = _efda.getBorder(_fdde.Left)
			_ffdg = true
		}
		if _fdde.Right != nil {
			_acad._fbfc = _efda.getBorder(_fdde.Right)
			_ffdg = true
		}
		if _eage.Wrapped() {
			_acad._cdad = true
			_ffdg = true
		}
		if _dea := _eage.GetVerticalAlignment(); _dea != _cd.ST_VerticalAlignmentUnset {
			_acad._dabg = _dea
			_ffdg = true
		}
		if _gcda := _eage.GetHorizontalAlignment(); _gcda != _cd.ST_HorizontalAlignmentUnset {
			_acad._aae = _gcda
			_ffdg = true
		}
	}
	if _ffdg {
		return _acad
	}
	return nil
}

type cell struct {
	_afg   _cd.ST_CellType
	_eaef  int
	_badg  float64
	_fag   []*line
	_faedb float64
	_bfa   float64
	_ebbd  float64
	_bcae  float64
	_gbccb float64
	_abcd  *_ag.TextStyle
	_ffg   *border
	_baab  *border
	_dfdb  *border
	_bda   *border
	_fbed  bool
	_efd   bool
}

const _cdc = 0.25

type pagespan struct {
	_ccdf  float64
	_cdafb []*page
	_agea  int
	_dged  int
}

func (_bgg *convertContext) makePages() {
	for _, _af := range _bgg._gcbg {
		for _, _dbad := range _bgg._adf {
			_af._cdafb = append(_af._cdafb, &page{_ccab: []*pageRow{}, _dfg: _af, _bgbd: _dbad})
		}
	}
}

type rowInfo struct {
	_dcda float64
	_bgdc bool
	_bae  float64
	_eagd *style
	_gfg  []*cell
	_egg  float64
}
type convertContext struct {
	_ffdc  *_ag.Creator
	_cceg  *_be.Workbook
	_adba  *_e.Theme
	_eedf  *_be.Sheet
	_gab   *_be.StyleSheet
	_cba   int
	_fbgd  int
	_gcbg  []*pagespan
	_cadb  *page
	_cdef  []*colInfo
	_fadb  []*rowInfo
	_adf   []*rowspan
	_gda   float64
	_fgb   float64
	_eeccb float64
	_gcg   float64
	_aac   []*mergedCell
	_cfff  []*anchor
}

func _cgbb(_aded *symbol) {
	_bcf := _ag.New()
	_dced := _bcf.NewStyledParagraph()
	_dced.SetMargins(0, 0, 0, 0)
	_gceb := _dced.Append(_aded._gbb)
	if _aded._acbd != nil {
		_gceb.Style = *_aded._acbd
	}
	_aded._acaa = _dced.Height()
	if _aded._dgaa == 0 {
		_aded._dgaa = _dced.Width()
	}
}
func (_dgad *convertContext) getSymbolsFromString(_abeg string, _daf *style) []*symbol {
	_gecga := []*symbol{}
	_gbcc := _dgad.makeTextStyleFromCellStyle(_daf)
	for _, _aaad := range _abeg {
		_gecga = append(_gecga, &symbol{_gbb: string(_aaad), _acbd: _gbcc})
	}
	return _gecga
}

type style struct {
	_fdge *string
	_edgc *float64
	_acg  *string
	_bcfb *bool
	_fbd  *bool
	_aeae *bool
	_dcab *bool
	_efgb *bool
	_afef *border
	_fddg *border
	_dfe  *border
	_fbfc *border
	_cdad bool
	_dabg _cd.ST_VerticalAlignment
	_aae  _cd.ST_HorizontalAlignment
}

func (_dacd *convertContext) imageFromAnchor(_fdd *anchor, _age, _aee float64) _ce.Image {
	if _fdd._fdbb != nil {
		return _fdd._fdbb
	}
	if _fdd._dcgg != nil {
		_fbb, _ecc := _bd.MakeImageFromChartSpace(_fdd._dcgg, _age, _aee, _dacd._adba)
		if _ecc != nil {
			_a.Log.Debug("C\u0061\u006e\u006e\u006f\u0074\u0020\u006d\u0061\u006b\u0065\u0020\u0061\u006e\u0020\u0069\u006d\u0061\u0067e\u0020\u0066\u0072\u006f\u006d\u0020\u0063\u0068\u0061\u0072tS\u0070\u0061\u0063e\u003a \u0025\u0073", _ecc)
			return nil
		}
		return _fbb
	}
	return nil
}

type colInfo struct {
	_faed float64
	_gfc  float64
	_ebag *style
}

// RegisterFont makes a PdfFont accessible for using in converting to PDF.
func RegisterFont(name string, style FontStyle, font *_f.PdfFont) {
	_bd.RegisterFont(name, style, font)
}

type anchor struct {
	_fdbb _ce.Image
	_dcgg *_eb.ChartSpace
	_gcf  int
	_fcf  int64
	_feg  int
	_badd int64
	_gddd int
	_ccf  int64
	_dfbe int
	_gga  int64
}
type rowspan struct {
	_cbba float64
	_eegg int
	_eadg int
}

func (_bbb *convertContext) makeRowspans() {
	var _bca float64
	_agc := 0
	for _dcg, _dba := range _bbb._fadb {
		_ffcc := _dba._bae + _dba._egg
		if _bca+_ffcc <= _bbb._eeccb {
			_dba._dcda = _bca
			_bca += _ffcc
		} else {
			_bbb._adf = append(_bbb._adf, &rowspan{_cbba: _bca, _eegg: _agc, _eadg: _dcg})
			_agc = _dcg
			_dba._dcda = 0
			_bca = _ffcc
		}
	}
	_bbb._adf = append(_bbb._adf, &rowspan{_cbba: _bca, _eegg: _agc, _eadg: len(_bbb._fadb)})
}
func (_acde *convertContext) getColorStringFromSmlColor(_fbgc *_cd.CT_Color) *string {
	var _begg string
	if _fbgc.RgbAttr != nil {
		_begg = *_fbgc.RgbAttr
	} else if _fbgc.IndexedAttr != nil && *_fbgc.IndexedAttr < 64 {
		_begg = _ead[*_fbgc.IndexedAttr]
	} else if _fbgc.ThemeAttr != nil {
		_bgcb := *_fbgc.ThemeAttr
		_begg = _acde.getColorFromTheme(_bgcb)
	}
	if _begg == "" {
		return nil
	}
	if len(_begg) > 6 {
		_begg = _begg[(len(_begg) - 6):]
	}
	if _fbgc.TintAttr != nil {
		_abge := *_fbgc.TintAttr
		_begg = _bd.AdjustColorByTint(_begg, _abge)
	}
	_begg = "\u0023" + _begg
	return &_begg
}

const _ef = 3

type page struct {
	_ccab []*pageRow
	_bee  bool
	_fbfe []*_ag.Image
	_dfg  *pagespan
	_bgbd *rowspan
}

func (_dcc *convertContext) getContentFromCell(_gba _be.Cell, _gbd *style, _faae float64, _ecbe bool) ([]*line, _cd.ST_CellType) {
	_ccb := _gba.X()
	var _aef []*symbol
	switch _ccb.TAttr {
	case _cd.ST_CellTypeS:
		_cfcb := _ccb.V
		if _cfcb != nil {
			_ebfge, _agd := _b.Atoi(*_cfcb)
			if _agd == nil {
				_bgbe := _dcc._cceg.SharedStrings.X().Si[_ebfge]
				if _bgbe.T != nil {
					_aef = _dcc.getSymbolsFromString(*_bgbe.T, _gbd)
				} else if _bgbe.R != nil {
					_aef = _dcc.getSymbolsFromR(_bgbe.R, _gbd)
				}
			}
		}
	case _cd.ST_CellTypeB:
		_ffcg := _ccb.V
		if _ffcg != nil {
			if *_ffcg == "\u0030" {
				_aef = _dcc.getSymbolsFromString("\u0046\u0041\u004cS\u0045", _gbd)
			} else {
				_aef = _dcc.getSymbolsFromString("\u0054\u0052\u0055\u0045", _gbd)
			}
		}
	default:
		_aef = _dcc.getSymbolsFromString(_gba.GetFormattedValue(), _gbd)
	}
	_dab := 0.0
	_cfga := 0.0
	var _ffd []*line
	var _dga bool
	if _gbd != nil {
		if _gbd._dcab != nil {
			if *_gbd._dcab {
				_dga = true
			}
		}
		if _gbd._efgb != nil {
			if *_gbd._efgb {
				_dga = true
			}
		}
	}
	if _ecbe {
		_ffd = []*line{}
		_fdbd := _faae - 2*_ef
		_ecbb := []*symbol{}
		for _, _addf := range _aef {
			_cgbb(_addf)
			if _dab+_addf._dgaa >= _fdbd {
				_cbf := _adgf(_ecbb)
				if _dga {
					_cbf /= _ab
				}
				_ffd = append(_ffd, &line{_cgc: _cfga, _acdd: _ecbb, _dcae: _cbf})
				_ecbb = []*symbol{_addf}
				_dab = _addf._dgaa
				_cfga += _cbf
			} else {
				_addf._agab = _dab
				_dab += _addf._dgaa
				_ecbb = append(_ecbb, _addf)
			}
		}
		_abgf := _adgf(_ecbb)
		if _dga {
			_abgf /= _ab
		}
		if len(_ecbb) > 0 {
			_ffd = append(_ffd, &line{_cgc: _cfga, _acdd: _ecbb, _dcae: _abgf})
		}
	} else {
		for _, _ccd := range _aef {
			_cgbb(_ccd)
			_ccd._agab = _dab
			_dab += _ccd._dgaa
		}
		if len(_aef) > 0 {
			_ffd = []*line{&line{_acdd: _aef, _dcae: _adgf(_aef)}}
		}
	}
	_dcd := _ccb.TAttr
	if _dcd == _cd.ST_CellTypeUnset {
		_dcd = _cd.ST_CellTypeN
	}
	return _ffd, _dcd
}
func _fecg(_adea, _dbbc *style) {
	if _dbbc == nil {
		return
	}
	if _adea == nil {
		_adea = _dbbc
		return
	}
	if _adea._acg == nil {
		_adea._acg = _dbbc._acg
	}
	if _adea._fdge == nil {
		_adea._fdge = _dbbc._fdge
	}
	if _adea._edgc == nil {
		_adea._edgc = _dbbc._edgc
	}
	if _adea._bcfb == nil {
		_adea._bcfb = _dbbc._bcfb
	}
	if _adea._fbd == nil {
		_adea._fbd = _dbbc._fbd
	}
	if _adea._aeae == nil {
		_adea._aeae = _dbbc._aeae
	}
	if _adea._dcab == nil {
		_adea._dcab = _dbbc._dcab
	}
	if _adea._efgb == nil {
		_adea._efgb = _dbbc._efgb
	}
	if _adea._afef == nil {
		_adea._afef = _dbbc._afef
	}
	if _adea._fddg == nil {
		_adea._fddg = _dbbc._fddg
	}
	if _adea._dfe == nil {
		_adea._dfe = _dbbc._dfe
	}
	if _adea._fbfc == nil {
		_adea._fbfc = _dbbc._fbfc
	}
	if _adea._dabg == _cd.ST_VerticalAlignmentUnset {
		_adea._dabg = _dbbc._dabg
	}
	if _adea._aae == _cd.ST_HorizontalAlignmentUnset {
		_adea._aae = _dbbc._aae
	}
}

const _ed = 1.5

type pageRow struct {
	_ecca int
	_gcd  []*cell
}

func (_adbg *convertContext) getColorFromTheme(_bdb uint32) string {
	_fcfg := _adbg._cceg.Themes()
	if len(_fcfg) != 0 {
		_cdae := _fcfg[0]
		if _aaf := _cdae.ThemeElements; _aaf != nil {
			if _ccgf := _aaf.ClrScheme; _ccgf != nil {
				switch _bdb {
				case 0:
					return _bd.GetColorStringFromDmlColor(_ccgf.Lt1)
				case 1:
					return _bd.GetColorStringFromDmlColor(_ccgf.Dk1)
				case 2:
					return _bd.GetColorStringFromDmlColor(_ccgf.Lt2)
				case 3:
					return _bd.GetColorStringFromDmlColor(_ccgf.Dk2)
				case 4:
					return _bd.GetColorStringFromDmlColor(_ccgf.Accent1)
				case 5:
					return _bd.GetColorStringFromDmlColor(_ccgf.Accent2)
				case 6:
					return _bd.GetColorStringFromDmlColor(_ccgf.Accent3)
				case 7:
					return _bd.GetColorStringFromDmlColor(_ccgf.Accent4)
				case 8:
					return _bd.GetColorStringFromDmlColor(_ccgf.Accent5)
				case 9:
					return _bd.GetColorStringFromDmlColor(_ccgf.Accent6)
				}
			}
		}
	}
	return ""
}
func (_df *convertContext) makeCells() {
	_faea := _df._eedf
	_abf := _faea.Rows()
	_dac := 0
	for _, _gac := range _df._fadb {
		_gac._gfg = []*cell{}
		_geg := 0.0
		_ff := _gac._eagd
		if _gac._bgdc {
			_bac := _abf[_dac]
			_dac++
			_adgc := _gac._bae
			for _, _gf := range _bac.Cells() {
				_gae, _cef := _bg.ParseCellReference(_gf.Reference())
				if _cef != nil {
					_a.Log.Debug("\u0043\u0061\u006e\u006eo\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0061\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u003a \u0025\u0073", _cef)
					continue
				}
				_agf := _df._cdef[_gae.ColumnIdx]
				_eac := _agf._gfc
				_cbc := _eac
				_abb := _agf._ebag
				var _cdaf, _ecf, _bgc, _ecb bool
				for _, _gee := range _df._aac {
					if _gae.RowIdx >= _gee._dgb && _gae.RowIdx <= _gee._geca && _gae.ColumnIdx >= _gee._eaf && _gae.ColumnIdx <= _gee._ecd {
						if _gae.ColumnIdx == _gee._eaf && _gae.RowIdx == _gee._dgb {
							_eac = _gee._cbe
							_adgc = _gee._ebe
						}
						_cdaf = _gae.RowIdx != _gee._dgb
						_ecf = _gae.RowIdx != _gee._geca
						_bgc = _gae.ColumnIdx != _gee._eaf
						_ecb = _gae.ColumnIdx != _gee._ecd
					}
				}
				_cbd := _df.getStyleFromCell(_gf, _ff, _abb)
				var _abe, _bf, _cac bool
				var _fga, _beb, _eecf, _fbc *border
				var _dfa _cd.ST_VerticalAlignment
				var _ffc _cd.ST_HorizontalAlignment
				if _cbd != nil {
					if !_cdaf {
						_fga = _cbd._afef
					}
					if !_ecf {
						_beb = _cbd._fddg
					}
					if !_bgc {
						_eecf = _cbd._dfe
					}
					if !_ecb {
						_fbc = _cbd._fbfc
					}
					if _beb != nil && _beb._ccdd > _geg {
						_geg = _beb._ccdd
					}
					_dfa = _cbd._dabg
					_ffc = _cbd._aae
					if _cbd._dcab != nil {
						_abe = *_cbd._dcab
					}
					if _cbd._efgb != nil {
						_bf = *_cbd._efgb
					}
					_cac = _cbd._cdad
				}
				_cab, _bfd := _df.getContentFromCell(_gf, _cbd, _eac, _cac)
				_ddd := &cell{_afg: _bfd, _faedb: _eac, _bfa: _cbc, _ebbd: _adgc, _fag: _cab, _ffg: _fga, _baab: _beb, _dfdb: _eecf, _bda: _fbc, _fbed: _abe, _efd: _bf}
				_df.alignSymbolsHorizontally(_ddd, _ffc)
				_df.alignSymbolsVertically(_ddd, _dfa)
				_gac._gfg = append(_gac._gfg, _ddd)
			}
		}
		_gac._egg = _geg
	}
}
func _ddbd(_bad []*symbol) float64 {
	_bbc := 0.0
	for _, _cgbc := range _bad {
		_bbc += _cgbc._dgaa
	}
	return _bbc
}
func (_fbac *convertContext) getStyleFromRPrElt(_dbg *_cd.CT_RPrElt) *style {
	if _dbg == nil {
		return nil
	}
	_bbbd := &style{}
	_bbbd._acg = &_dbg.RFont.ValAttr
	if _dfdc := _dbg.B; _dfdc != nil {
		_egee := _dfdc.ValAttr == nil || *_dfdc.ValAttr
		_bbbd._bcfb = &_egee
	}
	if _caaf := _dbg.I; _caaf != nil {
		_ebg := _caaf.ValAttr == nil || *_caaf.ValAttr
		_bbbd._fbd = &_ebg
	}
	if _abbb := _dbg.U; _abbb != nil {
		_gafd := _abbb.ValAttr == _cd.ST_UnderlineValuesSingle || _abbb.ValAttr == _cd.ST_UnderlineValuesUnset
		_bbbd._aeae = &_gafd
	}
	if _egaf := _dbg.VertAlign; _egaf != nil {
		_debf := _egaf.ValAttr == _fa.ST_VerticalAlignRunSuperscript
		_bbbd._dcab = &_debf
		_gagb := _egaf.ValAttr == _fa.ST_VerticalAlignRunSubscript
		_bbbd._efgb = &_gagb
	}
	if _cdfb := _dbg.Sz; _cdfb != nil {
		_fge := _cdfb.ValAttr / 12 * _bd.DefaultFontSize
		_bbbd._edgc = &_fge
	}
	if _bdc := _dbg.Color; _bdc != nil {
		_bbbd._fdge = _fbac.getColorStringFromSmlColor(_bdc)
	}
	return _bbbd
}

var _ead = []string{"\u0030\u0030\u0030\u0030\u0030\u0030", "\u0066\u0066\u0066\u0066\u0066\u0066", "\u0066\u0066\u0030\u0030\u0030\u0030", "\u0030\u0030\u0066\u0066\u0030\u0030", "\u0030\u0030\u0030\u0030\u0066\u0066", "\u0066\u0066\u0066\u0066\u0030\u0030", "\u0066\u0066\u0030\u0030\u0066\u0066", "\u0030\u0030\u0066\u0066\u0066\u0066", "\u0030\u0030\u0030\u0030\u0030\u0030", "\u0066\u0066\u0066\u0066\u0066\u0066", "\u0066\u0066\u0030\u0030\u0030\u0030", "\u0030\u0030\u0066\u0066\u0030\u0030", "\u0030\u0030\u0030\u0030\u0066\u0066", "\u0066\u0066\u0066\u0066\u0030\u0030", "\u0066\u0066\u0030\u0030\u0066\u0066", "\u0030\u0030\u0066\u0066\u0066\u0066", "\u0038\u0030\u0030\u0030\u0030\u0030", "\u0030\u0030\u0038\u0030\u0030\u0030", "\u0030\u0030\u0030\u0030\u0038\u0030", "\u0038\u0030\u0038\u0030\u0030\u0030", "\u0038\u0030\u0030\u0030\u0038\u0030", "\u0030\u0030\u0038\u0030\u0038\u0030", "\u0063\u0030\u0063\u0030\u0063\u0030", "\u0038\u0030\u0038\u0030\u0038\u0030", "\u0039\u0039\u0039\u0039\u0066\u0066", "\u0039\u0039\u0033\u0033\u0036\u0036", "\u0066\u0066\u0066\u0066\u0063\u0063", "\u0063\u0063\u0066\u0066\u0066\u0066", "\u0036\u0036\u0030\u0030\u0036\u0036", "\u0066\u0066\u0038\u0030\u0038\u0030", "\u0030\u0030\u0036\u0036\u0063\u0063", "\u0063\u0063\u0063\u0063\u0066\u0066", "\u0030\u0030\u0030\u0030\u0038\u0030", "\u0066\u0066\u0030\u0030\u0066\u0066", "\u0066\u0066\u0066\u0066\u0030\u0030", "\u0030\u0030\u0066\u0066\u0066\u0066", "\u0038\u0030\u0030\u0030\u0038\u0030", "\u0038\u0030\u0030\u0030\u0030\u0030", "\u0030\u0030\u0038\u0030\u0038\u0030", "\u0030\u0030\u0030\u0030\u0066\u0066", "\u0030\u0030\u0063\u0063\u0066\u0066", "\u0063\u0063\u0066\u0066\u0066\u0066", "\u0063\u0063\u0066\u0066\u0063\u0063", "\u0066\u0066\u0066\u0066\u0039\u0039", "\u0039\u0039\u0063\u0063\u0066\u0066", "\u0066\u0066\u0039\u0039\u0063\u0063", "\u0063\u0063\u0039\u0039\u0066\u0066", "\u0066\u0066\u0063\u0063\u0039\u0039", "\u0033\u0033\u0036\u0036\u0066\u0066", "\u0033\u0033\u0063\u0063\u0063\u0063", "\u0039\u0039\u0063\u0063\u0030\u0030", "\u0066\u0066\u0063\u0063\u0030\u0030", "\u0066\u0066\u0039\u0039\u0030\u0030", "\u0066\u0066\u0036\u0036\u0030\u0030", "\u0036\u0036\u0036\u0036\u0039\u0039", "\u0039\u0036\u0039\u0036\u0039\u0036", "\u0030\u0030\u0033\u0033\u0036\u0036", "\u0033\u0033\u0039\u0039\u0036\u0036", "\u0030\u0030\u0033\u0033\u0030\u0030", "\u0033\u0033\u0033\u0033\u0030\u0030", "\u0039\u0039\u0033\u0033\u0030\u0030", "\u0039\u0039\u0033\u0033\u0036\u0036", "\u0033\u0033\u0033\u0033\u0039\u0039", "\u0033\u0033\u0033\u0033\u0033\u0033"}

func _adgf(_gcb []*symbol) float64 {
	_egd := 0.0
	for _, _fbf := range _gcb {
		if _fbf._acaa > _egd {
			_egd = _fbf._acaa
		}
	}
	return _egd
}

// RegisterFontsFromDirectory registers all fonts from the given directory automatically detecting font families and styles.
func RegisterFontsFromDirectory(dirName string) error { return _bd.RegisterFontsFromDirectory(dirName) }
func (_cdca *convertContext) getStyleFromCell(_cdafe _be.Cell, _fegg, _aeg *style) *style {
	_fcgf := _cdafe.X()
	_cdag := _cdca.getStyle(_fcgf.SAttr)
	_fecg(_cdag, _fegg)
	_fecg(_cdag, _aeg)
	return _cdag
}

var _eg = _faabg(0.0625)

const _ga = 2

func (_fbceg *convertContext) combineCellStyleWithRPrElt(_dffc *style, _ccbb *_cd.CT_RPrElt) *style {
	_bgce := *_dffc
	_deda := _fbceg.getStyleFromRPrElt(_ccbb)
	if _deda == nil {
		return &_bgce
	}
	if _deda._fdge != nil {
		_bgce._fdge = _deda._fdge
	}
	if _deda._edgc != nil {
		_bgce._edgc = _deda._edgc
	}
	if _deda._acg != nil {
		_bgce._acg = _deda._acg
	}
	if _deda._bcfb != nil {
		_bgce._bcfb = _deda._bcfb
	}
	if _deda._fbd != nil {
		_bgce._fbd = _deda._fbd
	}
	if _deda._aeae != nil {
		_bgce._aeae = _deda._aeae
	}
	if _deda._dcab != nil {
		_bgce._dcab = _deda._dcab
	}
	if _deda._efgb != nil {
		_bgce._efgb = _deda._efgb
	}
	return &_bgce
}

const _ab = 0.64

func (_eec *convertContext) makeMergedCells() {
	_gbf := []*mergedCell{}
	for _, _adg := range _eec._eedf.MergedCells() {
		_deg, _gbg, _cad := _bg.ParseRangeReference(_adg.Reference())
		if _cad != nil {
			_a.Log.Debug("\u0065\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006c\u006c: \u0025\u0073", _cad)
			continue
		}
		_gecg := mergedCell{_dgb: _deg.RowIdx, _eaf: _deg.ColumnIdx, _geca: _gbg.RowIdx, _ecd: _gbg.ColumnIdx}
		for _fcg := _gecg._eaf - 1; _fcg < _gecg._ecd; _fcg++ {
			_gecg._cbe += _eec._cdef[_fcg]._gfc
		}
		for _eae := _gecg._dgb - 1; _eae < _gecg._geca; _eae++ {
			_gecg._ebe += _eec._fadb[_eae]._bae
		}
		_gbf = append(_gbf, &_gecg)
	}
	_eec._aac = _gbf
}
func (_ece *convertContext) fillPages() {
	for _gfd, _adb := range _ece._adf {
		_bfb := _ece._fadb[_adb._eegg:_adb._eadg]
		for _edf, _gbea := range _bfb {
			_abc := 0
			_faab := 0.0
			_deef := []*cell{}
			if _gbea._bgdc {
				for _, _deb := range _gbea._gfg {
					_dbf := _ece._gcbg[_abc]
					_ece._cadb = _dbf._cdafb[_gfd]
					_ece._cadb._bee = true
					_fffd := _deb._bfa
					if _faab+_fffd > _dbf._ccdf {
						_ece.addRowToPage(_deef, _edf)
						_deef = []*cell{_deb}
						_faab = _fffd
						_abc++
					} else {
						_deb._badg = _faab
						_deef = append(_deef, _deb)
						_faab += _fffd
					}
				}
				if len(_deef) > 0 {
					_gdeg := _ece._gcbg[_abc]
					_ece._cadb = _gdeg._cdafb[_gfd]
					_ece._cadb._bee = true
					_ece.addRowToPage(_deef, _edf)
				}
			}
		}
	}
}

const (
	FontStyle_Regular    FontStyle = 0
	FontStyle_Bold       FontStyle = 1
	FontStyle_Italic     FontStyle = 2
	FontStyle_BoldItalic FontStyle = 3
)

func (_dega *convertContext) drawPage(_agcc *page) {
	_ggc := _dega._gda
	_dace := _dega._fgb
	for _, _dbb := range _agcc._ccab {
		_ffb := _dega._fadb[_dbb._ecca]
		for _, _eeac := range _dbb._gcd {
			_cfdb := _eeac._bcae < _eeac._badg
			_ffe := _eeac._gbccb > _eeac._badg+_eeac._faedb
			var _efbe, _aaaf bool
			for _, _efc := range _eeac._fag {
				for _, _cbcb := range _efc._acdd {
					if _cfdb && !_efbe {
						_efbe = _cbcb._agab < 0
					}
					if _ffe && !_aaaf {
						_aaaf = _eeac._faedb < _cbcb._agab+_cbcb._dgaa
					}
					if _eeac._badg+_cbcb._agab >= _eeac._bcae && _eeac._badg+_cbcb._agab+_cbcb._dgaa <= _eeac._gbccb {
						_aggb := _dega._ffdc.NewStyledParagraph()
						_gce := _dace + _eeac._badg + _cbcb._agab
						_dcgf := _ggc + _ffb._dcda + _efc._cgc - _cbcb._acaa - _faabg(0.5)
						_aggb.SetPos(_gce, _dcgf)
						var _eeg *_ag.TextChunk
						if _cbcb._feaa != "" {
							_eeg = _aggb.AddExternalLink(_cbcb._gbb, _cbcb._feaa)
						} else {
							_eeg = _aggb.Append(_cbcb._gbb)
						}
						if _cbcb._acbd != nil {
							_eeg.Style = *_cbcb._acbd
						}
						_dega._ffdc.Draw(_aggb)
					}
				}
			}
			var _fba, _fe, _cbdd, _fde, _ddc, _efef float64
			var _egff, _fgce, _eaca, _bag _ag.Color
			if _gaca := _eeac._ffg; _gaca != nil {
				_fba = _gaca._ccdd
				_egff = _gaca._beg
			}
			if _abgg := _eeac._baab; _abgg != nil {
				_fe = _abgg._ccdd
				_fgce = _abgg._beg
			}
			if _fada := _eeac._dfdb; _fada != nil {
				_cbdd = _fada._ccdd
				_ddc = _cbdd / 2
				_eaca = _fada._beg
			}
			if _cgb := _eeac._bda; _cgb != nil {
				_fde = _cgb._ccdd
				_efef = _fde / 2
				_bag = _cgb._beg
			}
			var _ddb float64
			if _dbb._ecca > 1 {
				_ddb = _dega._fadb[_dbb._ecca-1]._egg
			}
			_dgc := _ggc + _ffb._dcda - 0.5*(_ddb-_fba)
			_bede := _ggc + _ffb._dcda + _ffb._bae + 0.5*(_ffb._egg+_fe)
			_egfa := _dace + _eeac._badg
			_dbc := _egfa + _eeac._bfa
			_bd.DrawLine(_dega._ffdc, _egfa, _dgc, _dbc, _dgc, _fba, _egff)
			_bd.DrawLine(_dega._ffdc, _egfa, _bede, _dbc, _bede, _fe, _fgce)
			if !_efbe {
				_bd.DrawLine(_dega._ffdc, _egfa-_ddc, _dgc, _egfa-_ddc, _bede, _cbdd, _eaca)
			}
			if !_aaaf {
				_bd.DrawLine(_dega._ffdc, _dbc-_efef, _dgc, _dbc-_efef, _bede, _fde, _bag)
			}
		}
	}
	for _, _efg := range _agcc._fbfe {
		if _efg != nil {
			_dega._ffdc.Draw(_efg)
		}
	}
}

type line struct {
	_cgc  float64
	_acdd []*symbol
	_dcae float64
}

// FontStyle represents a kind of font styling. It can be FontStyle_Regular, FontStyle_Bold, FontStyle_Italic and FontStyle_BoldItalic.
type FontStyle = _bd.FontStyle
type symbol struct {
	_gbb  string
	_agab float64
	_acaa float64
	_dgaa float64
	_acbd *_ag.TextStyle
	_feaa string
}

var _fgad = map[uint32]_ag.PageSize{1: _ag.PageSize{8.5 * _g.Inch, 11 * _g.Inch}, 2: _ag.PageSize{8.5 * _g.Inch, 11 * _g.Inch}, 3: _ag.PageSize{11 * _g.Inch, 17 * _g.Inch}, 4: _ag.PageSize{17 * _g.Inch, 11 * _g.Inch}, 5: _ag.PageSize{8.5 * _g.Inch, 14 * _g.Inch}, 6: _ag.PageSize{5.5 * _g.Inch, 8.5 * _g.Inch}, 7: _ag.PageSize{7.5 * _g.Inch, 10 * _g.Inch}, 8: _ag.PageSize{_faabg(297), _faabg(420)}, 9: _ag.PageSize{_faabg(210), _faabg(297)}, 10: _ag.PageSize{_faabg(210), _faabg(297)}, 11: _ag.PageSize{_faabg(148), _faabg(210)}, 70: _ag.PageSize{_faabg(105), _faabg(148)}, 12: _ag.PageSize{_faabg(250), _faabg(354)}, 13: _ag.PageSize{_faabg(182), _faabg(257)}, 14: _ag.PageSize{8.5 * _g.Inch, 13 * _g.Inch}, 20: _ag.PageSize{4.125 * _g.Inch, 9.5 * _g.Inch}, 27: _ag.PageSize{_faabg(110), _faabg(220)}, 28: _ag.PageSize{_faabg(162), _faabg(229)}, 34: _ag.PageSize{_faabg(250), _faabg(176)}, 29: _ag.PageSize{_faabg(324), _faabg(458)}, 30: _ag.PageSize{_faabg(229), _faabg(324)}, 31: _ag.PageSize{_faabg(114), _faabg(162)}, 37: _ag.PageSize{3.88 * _g.Inch, 7.5 * _g.Inch}, 43: _ag.PageSize{_faabg(100), _faabg(148)}, 69: _ag.PageSize{_faabg(200), _faabg(148)}}

func _faabg(_dgcb float64) float64 { return _dgcb * _g.Millimeter }

type border struct {
	_ccdd float64
	_beg  _ag.Color
}

// ConvertToPdf converts a sheet to a PDF file. This package is beta, breaking changes can take place.
func ConvertToPdf(s *_be.Sheet) *_ag.Creator {
	_d := s.X()
	if _d == nil {
		return nil
	}
	var _dg _ag.PageSize
	var _bed bool
	if _fg := _d.PageSetup; _fg != nil {
		if _ca := _fg.PaperSizeAttr; _ca != nil {
			_dg = _fgad[*_ca]
		}
		_bed = _fg.OrientationAttr == _cd.ST_OrientationLandscape
	}
	if (_dg == _ag.PageSize{}) {
		_dg = _fgad[1]
	}
	if _bed {
		_dg[0], _dg[1] = _dg[1], _dg[0]
	}
	_bgd := _ag.New()
	_bgd.SetPageSize(_dg)
	var _dge, _fc, _ge, _gb float64
	if _cb := _d.PageMargins; _cb != nil {
		_ge = _cb.LeftAttr
		_gb = _cb.RightAttr
		_dge = _cb.TopAttr
		_fc = _cb.BottomAttr
	}
	if _ge < 0.25 {
		_ge = 0.25
	}
	if _gb < 0.25 {
		_gb = 0.25
	}
	_dge *= _g.Inch
	_fc *= _g.Inch
	_ge *= _g.Inch
	_gb *= _g.Inch
	_bgd.SetPageMargins(_ge, _gb, _dge, _fc)
	_edb := s.Workbook()
	var _cde *_e.Theme
	if len(_edb.Themes()) > 0 {
		_cde = _edb.Themes()[0]
	}
	_bc := &convertContext{_ffdc: _bgd, _eedf: s, _cceg: s.Workbook(), _adba: _cde, _gab: &s.Workbook().StyleSheet, _gda: _dge, _fgb: _ge, _eeccb: _dg[1] - _fc - _dge, _gcg: _dg[0] - _gb - _ge}
	_bc.makeAnchors()
	_bc.determineMaxIndexes()
	if _bc._cba == 0 && _bc._fbgd == 0 {
		_bgd.NewPage()
		return _bgd
	}
	_bc.makeCols()
	_bc.makeRows()
	_bc.makeMergedCells()
	_bc.makeCells()
	_bc.makePagespans()
	_bc.makeRowspans()
	_bc.makePages()
	_bc.fillPages()
	_bc.distributeAnchors()
	_bc.drawSheet()
	return _bgd
}
func (_egaa *convertContext) makeRows() {
	_cff := []*rowInfo{}
	_bab := _egaa._eedf.Rows()
	_aa := 0
	for _, _cec := range _bab {
		_aa++
		_ec := int(_cec.RowNumber())
		if _ec > _aa {
			for _fae := _aa; _fae < _ec; _fae++ {
				_cff = append(_cff, &rowInfo{_bae: 16 / _cce})
			}
			_aa = _ec
		}
		var _fgf float64
		if _cec.X().HtAttr == nil {
			_fgf = 16
		} else {
			_fgf = *_cec.X().HtAttr
		}
		_cff = append(_cff, &rowInfo{_bae: _fgf / _cce, _bgdc: true, _eagd: _egaa.getStyle(_cec.X().SAttr)})
	}
	for _aga := len(_cff); _aga < _egaa._cba; _aga++ {
		_cff = append(_cff, &rowInfo{_bae: 16 / _cce})
	}
	_egaa._fadb = _cff
}
