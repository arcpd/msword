package convertutils

import (
	_d "bytes"
	_ef "errors"
	_e "fmt"
	_bc "github.com/arcpd/msword/common/logger"
	_aa "github.com/arcpd/msword/measurement"
	_fba "github.com/arcpd/msword/schema/soo/dml"
	_bd "github.com/arcpd/msword/schema/soo/dml/chart"
	_fb "github.com/arcpd/msword/spreadsheet/format"
	_gg "github.com/unidoc/unipdf/v3/creator"
	_ce "github.com/unidoc/unipdf/v3/model"
	_adc "github.com/unidoc/unipdf/v3/render"
	_bb "github.com/unidoc/unitype"
	_c "image"
	_ad "math"
	_g "os"
	_ba "sort"
	_ge "strconv"
	_fe "strings"
	_b "sync"
	_f "unicode"
)

func (_dac barSerByOrder) Swap(i, j int) { _dac[i], _dac[j] = _dac[j], _dac[i] }

func FromSTCoordinate(st _fba.ST_Coordinate) int64 {
	if _babg := st.ST_CoordinateUnqualified; _babg != nil {
		return *_babg
	}
	return 0
}

type legendItem struct {
	_acd string
	_afc *_fba.CT_ShapeProperties
}

func GetRegisteredFont(name string, style FontStyle) *_ce.PdfFont {
	_cdac._adec.Lock()
	defer _cdac._adec.Unlock()
	if _aaab, _fagb := _cdac._cgdg[name]; _fagb {
		if _gggg, _feeg := _aaab[style]; _feeg {
			return _gggg
		}
	}
	return nil
}

var _bg = _ebfa(1.5)

var _gfbe = _ebfa(1.5)

type serCategory struct {
	_ece  string
	_aeda []serValue
}

var _ggbc = _ebfa(5)

func _cbg(_gaag *_bd.CT_ValAx) (uint32, _bd.ST_AxPos, _bd.ST_TickMark, _bd.ST_TickLblPos, *_bd.CT_ChartLines, uint32, *_fba.CT_ShapeProperties, error) {
	var _bef, _bbae uint32
	var _befg _bd.ST_AxPos
	var _gdfc _bd.ST_TickMark
	var _ecb *_bd.CT_ChartLines
	var _dbbf _bd.ST_TickLblPos
	if _gaag.AxId == nil {
		return _bef, _befg, _gdfc, _dbbf, _ecb, _bbae, _gaag.SpPr, _ef.New("\u004e\u006f\u0020x\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044")
	} else {
		_bef = _gaag.AxId.ValAttr
	}
	if _gaag.AxPos == nil {
		return _bef, _befg, _gdfc, _dbbf, _ecb, _bbae, _gaag.SpPr, _ef.New("\u004eo\u0020x\u0020\u0061\u0078\u0069\u0073 \u0070\u006fs\u0069\u0074\u0069\u006f\u006e")
	} else {
		_befg = _gaag.AxPos.ValAttr
	}
	if _gaag.MajorTickMark != nil {
		_gdfc = _gaag.MajorTickMark.ValAttr
	}
	if _gaag.TickLblPos != nil {
		_dbbf = _gaag.TickLblPos.ValAttr
	}
	if _gaag.CrossAx == nil {
		return _bef, _befg, _gdfc, _dbbf, _ecb, _bbae, _gaag.SpPr, _ef.New("\u004e\u006f \u0063\u0072\u006fs\u0073\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044")
	} else {
		_bbae = _gaag.CrossAx.ValAttr
	}
	_ecb = _gaag.MajorGridlines
	return _bef, _befg, _gdfc, _dbbf, _ecb, _bbae, _gaag.SpPr, nil
}

func MakeImageFromChartSpace(cs *_bd.ChartSpace, width, height float64, theme *_fba.Theme) (_c.Image, error) {
	_dae, _fgbf := _ebf(cs, width, height, theme, true)
	if _fgbf != nil {
		return nil, _fgbf
	}
	_baf, _fgbf := GetPageFromCreator(_dae)
	if _fgbf != nil {
		return nil, _fgbf
	}
	return _adc.NewImageDevice().Render(_baf)
}

func RegisterFont(name string, style FontStyle, font *_ce.PdfFont) {
	_cdac._adec.Lock()
	if _cdac._cgdg[name] == nil {
		_cdac._cgdg[name] = map[FontStyle]*_ce.PdfFont{}
	}
	_cdac._cgdg[name][style] = font
	_cdac._adec.Unlock()
}

func FromSTCoordinate32(st _fba.ST_Coordinate32) int64 {
	if _edb := st.ST_Coordinate32Unqualified; _edb != nil {
		return int64(*_edb)
	}
	return 0
}

var _fbb = _ebfa(1.5)

func DrawRectangle(c *_gg.Creator, r *Rectangle, w float64, color _gg.Color) {
	if color == nil {
		return
	}
	DrawLine(c, r.Left, r.Top, r.Right, r.Top, w, color)
	DrawLine(c, r.Left, r.Top, r.Left, r.Bottom, w, color)
	DrawLine(c, r.Left, r.Bottom, r.Right, r.Bottom, w, color)
	DrawLine(c, r.Right, r.Top, r.Right, r.Bottom, w, color)
}

func (_bdbf *Rectangle) Translate(x, y float64) {
	_bdbf.Left += x
	_bdbf.Right += x
	_bdbf.Top += y
	_bdbf.Bottom += y
}

func (_dgac *Rectangle) scale(_afde float64) {
	_dgac.Top *= _afde
	_dgac.Bottom *= _afde
	_dgac.Left *= _afde
	_dgac.Right *= _afde
}

type FontStyle byte

const (
	FontStyle_Regular    FontStyle = 0
	FontStyle_Bold       FontStyle = 1
	FontStyle_Italic     FontStyle = 2
	FontStyle_BoldItalic FontStyle = 3
)

func _efe(_bbb *_bd.CT_DateAx) (uint32, _bd.ST_AxPos, _bd.ST_TickMark, _bd.ST_TickLblPos, *_bd.CT_ChartLines, uint32, *_fba.CT_ShapeProperties, error) {
	var _fbbcf, _gabae uint32
	var _dgb _bd.ST_AxPos
	var _ged _bd.ST_TickMark
	var _afbf *_bd.CT_ChartLines
	var _bfad _bd.ST_TickLblPos
	if _bbb.AxId == nil {
		return _fbbcf, _dgb, _ged, _bfad, _afbf, _gabae, _bbb.SpPr, _ef.New("\u004e\u006f\u0020x\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044")
	} else {
		_fbbcf = _bbb.AxId.ValAttr
	}
	if _bbb.AxPos == nil {
		return _fbbcf, _dgb, _ged, _bfad, _afbf, _gabae, _bbb.SpPr, _ef.New("\u004eo\u0020x\u0020\u0061\u0078\u0069\u0073 \u0070\u006fs\u0069\u0074\u0069\u006f\u006e")
	} else {
		_dgb = _bbb.AxPos.ValAttr
	}
	if _bbb.MajorTickMark != nil {
		_ged = _bbb.MajorTickMark.ValAttr
	}
	if _bbb.TickLblPos != nil {
		_bfad = _bbb.TickLblPos.ValAttr
	}
	if _bbb.CrossAx == nil {
		return _fbbcf, _dgb, _ged, _bfad, _afbf, _gabae, _bbb.SpPr, _ef.New("\u004e\u006f \u0063\u0072\u006fs\u0073\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044")
	} else {
		_gabae = _bbb.CrossAx.ValAttr
	}
	_afbf = _bbb.MajorGridlines
	return _fbbcf, _dgb, _ged, _bfad, _afbf, _gabae, _bbb.SpPr, nil
}

var _cbac = map[string]FontStyle{"\u0052e\u0067\u0075\u006c\u0061\u0072": FontStyle_Regular, "\u0042\u006f\u006c\u0064": FontStyle_Bold, "\u0049\u0074\u0061\u006c\u0069\u0063": FontStyle_Italic, "B\u006f\u006c\u0064\u0020\u0049\u0074\u0061\u006c\u0069\u0063": FontStyle_BoldItalic}

var _ccc = _ebfa(7.5)

func _afdg(_deae *_bd.CT_SerAx) (uint32, _bd.ST_AxPos, _bd.ST_TickMark, _bd.ST_TickLblPos, *_bd.CT_ChartLines, uint32, *_fba.CT_ShapeProperties, error) {
	var _egfb, _ede uint32
	var _cgg _bd.ST_AxPos
	var _bddg _bd.ST_TickMark
	var _egd *_bd.CT_ChartLines
	var _fgc _bd.ST_TickLblPos
	if _deae.AxId == nil {
		return _egfb, _cgg, _bddg, _fgc, _egd, _ede, _deae.SpPr, _ef.New("\u004e\u006f\u0020x\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044")
	} else {
		_egfb = _deae.AxId.ValAttr
	}
	if _deae.AxPos == nil {
		return _egfb, _cgg, _bddg, _fgc, _egd, _ede, _deae.SpPr, _ef.New("\u004eo\u0020x\u0020\u0061\u0078\u0069\u0073 \u0070\u006fs\u0069\u0074\u0069\u006f\u006e")
	} else {
		_cgg = _deae.AxPos.ValAttr
	}
	if _deae.MajorTickMark != nil {
		_bddg = _deae.MajorTickMark.ValAttr
	}
	if _deae.TickLblPos != nil {
		_fgc = _deae.TickLblPos.ValAttr
	}
	if _deae.CrossAx == nil {
		return _egfb, _cgg, _bddg, _fgc, _egd, _ede, _deae.SpPr, _ef.New("\u004e\u006f \u0063\u0072\u006fs\u0073\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044")
	} else {
		_ede = _deae.CrossAx.ValAttr
	}
	_egd = _deae.MajorGridlines
	return _egfb, _cgg, _bddg, _fgc, _egd, _ede, _deae.SpPr, nil
}

func AssignStdFontByName(style _gg.TextStyle, fontName string) *_ce.PdfFont {
	_ccd := _ce.StdFontName(fontName)
	return _ce.NewStandard14FontMustCompile(_ccd)
}

func AdjustColorByLumMod(colorStr string, lum float64) string {
	var _dada, _cdd, _cfgd uint8
	_dbff, _ := _e.Sscanf(colorStr, "\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078", &_dada, &_cdd, &_cfgd)
	if _dbff != 3 {
		return ""
	}
	_baa, _cfba, _afg := _agf(_dada, _cdd, _cfgd)
	_afg = lum * _afg
	_dada, _cdd, _cfgd = _bcac(_baa, _cfba, _afg)
	return _e.Sprintf("\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078", _dada, _cdd, _cfgd)
}

func (_cfgb *creatorContext) drawBorderWithProps(_dded *_fba.CT_ShapeProperties, _cgfa *Rectangle, _gea float64) {
	if _cgfa != nil && _dded != nil && _dded.Ln != nil && _dded.Ln.SolidFill != nil {
		_ccbd := _cfgb.getPdfColorFromSolidFill(_dded.Ln.SolidFill)
		DrawRectangle(_cfgb._dcabc, _cgfa, _gea, _ccbd)
	}
}

func (_aab *creatorContext) drawLegend(_aga *Rectangle, _aedb []*legendItem, _dafe bool) {
	_gdb := _aab._bdda
	_beb := _ebfa(2.5) * _gdb
	_cda := _bg * _gdb
	_afdf := (_beb - _cda) / 2
	_bga := float64(len(_aedb))
	if _dafe {
		_egf := &Rectangle{Top: _aga.Top + _ebfa(1)*_gdb, Bottom: _aga.Bottom - _ebfa(1)*_gdb, Left: _aga.Left + _ebfa(2.5)*_gdb, Right: _aga.Right - _ebfa(2.5)*_gdb}
		var _dffb float64
		if _bga > 1 {
			_dffb = (_egf.Right - _egf.Left) / _bga
		}
		_bfgb := _egf.Left
		_cdgb := _egf.Top
		for _, _cb := range _aedb {
			if _aded := _cb._afc; _aded != nil {
				_aab.drawRectangleWithProps(_aded, _bfgb, _cdgb+_afdf, _cda, _cda, false)
				_bad := _bfgb + _cda*2
				_caf := _aab._dcabc.NewStyledParagraph()
				_caf.SetPos(_bad, _cdgb)
				_gdd := _caf.Append(_cb._acd)
				_gcc, _bae := _ce.NewStandard14Font(_ce.HelveticaName)
				if _bae == nil {
					_gdd.Style = _gg.TextStyle{FontSize: _beb, Font: _gcc, TextRise: 0.4}
					_aab._dcabc.Draw(_caf)
				}
			}
			_bfgb += _dffb
		}
	} else {
		_acf := &Rectangle{Top: _aga.Top + _ebfa(2.5)*_gdb, Bottom: _aga.Bottom - _ebfa(2.5)*_gdb, Left: _aga.Left + _ebfa(2.5)*_gdb, Right: _aga.Right - _ebfa(2.5)*_gdb}
		var _gfd float64
		if _bga > 1 {
			_gfd = (_acf.Bottom - _acf.Top - _beb) / (_bga - 1)
		}
		_gdc := _acf.Top
		_adfe := _acf.Left
		_ddc := _adfe + _cda*2
		for _, _gfc := range _aedb {
			if _dgf := _gfc._afc; _dgf != nil {
				_aab.drawRectangleWithProps(_dgf, _adfe, _gdc+_afdf, _cda, _cda, false)
				_fggf := _aab._dcabc.NewStyledParagraph()
				_fggf.SetPos(_ddc, _gdc)
				_gfa := _fggf.Append(_gfc._acd)
				_gag, _fdd := _ce.NewStandard14Font(_ce.HelveticaName)
				if _fdd == nil {
					_gfa.Style = _gg.TextStyle{FontSize: _beb, Font: _gag, TextRise: 0.4}
					_aab._dcabc.Draw(_fggf)
				}
			}
			_gdc += _gfd
		}
	}
}

func GetColorStringFromDmlColor(dmlColor *_fba.CT_Color) string {
	var _dfa string
	if _bff := dmlColor.SrgbClr; _bff != nil {
		_dfa = _bff.ValAttr
	} else if _bfe := dmlColor.SysClr; _bfe != nil {
		return "\u0030\u0030\u0030\u0030\u0030\u0030"
	}
	return _dfa
}

func RegisterFontsFromDirectory(dirName string) error {
	_affg, _ccgf := _g.Open(dirName)
	if _ccgf != nil {
		return _ccgf
	}
	defer _affg.Close()
	_dfgd, _ccgf := _affg.Readdirnames(0)
	if _ccgf != nil {
		return _ccgf
	}
	for _, _bgac := range _dfgd {
		if _fe.HasSuffix(_bgac, "\u002e\u0074\u0074\u0066") {
			_facb := dirName + "\u002f" + _bgac
			_daccd := _ebfg(_facb)
			if _daccd != nil {
				_bc.Log.Debug("\u0075\u006ea\u0062\u006c\u0065\u0020\u0074o\u0020\u0070\u0072\u006f\u0063e\u0073\u0073\u0020\u0061\u006e\u0064\u0020\u0072\u0065\u0067\u0069\u0073\u0074\u0065\u0072\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020\u0054\u0054\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073", _daccd)
				continue
			}
		}
	}
	return nil
}

func _aebf(_ebbff, _edd, _bfab float64) float64 {
	if _ebbff*6 < 1 {
		return _bfab + (_edd-_bfab)*6*_ebbff
	} else if _ebbff*2 < 1 {
		return _edd
	} else if _ebbff*3 < 2 {
		return _bfab + (_edd-_bfab)*(2.0/3.0-_ebbff)*6
	} else {
		return _bfab
	}
}

func _dbbg(_fegg uint8, _fdg float64) string {
	_cbb := float64(_fegg)
	return _e.Sprintf("\u0025\u0030\u0032\u0078", int(_cbb*_fdg))
}

var _cdac = fontsMap{_adec: &_b.Mutex{}, _cgdg: map[string]map[FontStyle]*_ce.PdfFont{}}

func CropImageByRect(sourceImg _c.Image, rect _c.Rectangle) _c.Image {
	_ccgc, _cff, _ddd, _caae := rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y
	_gefb := _c.NewNRGBA(_c.Rect(0, 0, _ddd-_ccgc, _caae-_cff))
	for _efaa := _ccgc; _efaa < _ddd; _efaa++ {
		for _cdba := _cff; _cdba < _caae; _cdba++ {
			_gefb.Set(_efaa-_ccgc, _cdba-_cff, sourceImg.At(_efaa, _cdba))
		}
	}
	return _gefb
}

func _cec(_ead _fba.ST_SchemeColorVal, _fga *_fba.Theme) string {
	if _afe := _fga.ThemeElements; _afe != nil {
		if _gfbab := _afe.ClrScheme; _gfbab != nil {
			switch _ead {
			case _fba.ST_SchemeColorValLt1:
				return GetColorStringFromDmlColor(_gfbab.Lt1)
			case _fba.ST_SchemeColorValDk1, _fba.ST_SchemeColorValTx1:
				return GetColorStringFromDmlColor(_gfbab.Dk1)
			case _fba.ST_SchemeColorValLt2:
				return GetColorStringFromDmlColor(_gfbab.Lt2)
			case _fba.ST_SchemeColorValDk2:
				return GetColorStringFromDmlColor(_gfbab.Dk2)
			case _fba.ST_SchemeColorValAccent1:
				return GetColorStringFromDmlColor(_gfbab.Accent1)
			case _fba.ST_SchemeColorValAccent2:
				return GetColorStringFromDmlColor(_gfbab.Accent2)
			case _fba.ST_SchemeColorValAccent3:
				return GetColorStringFromDmlColor(_gfbab.Accent3)
			case _fba.ST_SchemeColorValAccent4:
				return GetColorStringFromDmlColor(_gfbab.Accent4)
			case _fba.ST_SchemeColorValAccent5:
				return GetColorStringFromDmlColor(_gfbab.Accent5)
			case _fba.ST_SchemeColorValAccent6:
				return GetColorStringFromDmlColor(_gfbab.Accent6)
			}
		}
	}
	return ""
}

func (_ggf *creatorContext) drawAxes(_cbf *_bd.CT_PlotAreaChoice1, _gce, _dcd, _efb float64, _eca []string, _agc *Rectangle, _aeg bool) error {
	_cdc := _ggf._dcabc
	_fbbc := _ggf._bdda
	if _cbf == nil {
		return _ef.New("\u004e\u006f\u0020\u0061xi\u0073\u0020\u0069\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e")
	}
	if len(_cbf.ValAx) == 0 || (len(_cbf.CatAx) == 0 && len(_cbf.DateAx) == 0 && len(_cbf.SerAx) == 0) {
		return _ef.New("\u004e\u006f\u0020\u0078\u0020\u006f\u0072\u0020\u0079 \u0061\u0078\u0069\u0073")
	}
	var _eecc, _eega, _beea, _dcca uint32
	var _dacc, _eef _bd.ST_AxPos
	var _bec, _dfc _bd.ST_TickMark
	var _cg, _cdgd *_bd.CT_ChartLines
	var _bfb, _cba _bd.ST_TickLblPos
	var _gae, _gda *_fba.CT_ShapeProperties
	var _ffd error
	_gbc := _agc.Right - _agc.Left
	_fcf := _agc.Bottom - _agc.Top
	if len(_cbf.ValAx) > 0 {
		_eega, _eef, _dfc, _cba, _cdgd, _dcca, _gda, _ffd = _cbg(_cbf.ValAx[0])
	}
	if _eef != _bd.ST_AxPosL && _eef != _bd.ST_AxPosB {
		return _ef.New("\u004f\u006e\u006c\u0079\u0020l\u0065\u0066\u0074\u0020\u006f\u0072\u0020\u0062\u006f\u0074\u0074\u006f\u006d \u0079\u0020\u0061\u0078\u0069\u0073\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0073\u006f\u0020\u0066\u0061\u0072")
	}
	_fbe := _dcd - _gce
	_ddg := int(_fbe/_efb) + 1
	var _gfba, _dea float64
	switch _dfc {
	case _bd.ST_TickMarkIn:
		_gfba, _dea = _ebb, 0
	case _bd.ST_TickMarkOut:
		_gfba, _dea = 0, _ebb
	case _bd.ST_TickMarkCross:
		_gfba, _dea = _ebb, _ebb
	}
	_gfba = _gfba * _fbbc
	_dea = _dea * _fbbc
	var _gdaa *_fba.CT_ShapeProperties
	if _cdgd != nil {
		_gdaa = _cdgd.SpPr
	}
	_eagb, _bebg := _bfb != _bd.ST_TickLblPosNone, _cba != _bd.ST_TickLblPosNone
	_gdbd := _gce
	if len(_cbf.CatAx) > 0 {
		_eecc, _dacc, _bec, _bfb, _cg, _beea, _gae, _ffd = _dgae(_cbf.CatAx[0])
	} else if len(_cbf.DateAx) > 0 {
		_eecc, _dacc, _bec, _bfb, _cg, _beea, _gae, _ffd = _efe(_cbf.DateAx[0])
	} else if len(_cbf.SerAx) > 0 {
		_eecc, _dacc, _bec, _bfb, _cg, _beea, _gae, _ffd = _afdg(_cbf.SerAx[0])
	}
	if _ffd != nil {
		return _ffd
	}
	if _dacc != _bd.ST_AxPosL && _dacc != _bd.ST_AxPosB {
		return _ef.New("\u004f\u006e\u006c\u0079\u0020l\u0065\u0066\u0074\u0020\u006f\u0072\u0020\u0062\u006f\u0074\u0074\u006f\u006d \u0078\u0020\u0061\u0078\u0069\u0073\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0073\u006f\u0020\u0066\u0061\u0072")
	}
	if _eecc != _dcca || _eega != _beea {
		return _ef.New("a\u0078i\u0073\u0020\u0069\u0064\u0073\u0020\u0064\u006fn\u0027\u0074\u0020\u006dat\u0063\u0068")
	}
	_fac := len(_eca) + 1
	var _gdae, _gab float64
	switch _bec {
	case _bd.ST_TickMarkIn:
		_gdae, _gab = _ebb, 0
	case _bd.ST_TickMarkOut:
		_gdae, _gab = 0, _ebb
	case _bd.ST_TickMarkCross:
		_gdae, _gab = _ebb, _ebb
	}
	_gdae = _gdae * _fbbc
	_gab = _gab * _fbbc
	var _bbdg *_fba.CT_ShapeProperties
	if _cg != nil {
		_bbdg = _cg.SpPr
	}
	if _aeg {
		_eac := _fcf / float64(len(_eca))
		_bba := _agc.Left - _gce*_gbc/_fbe
		_abe := _bba - _ccc*_fbbc
		if _eagb {
			var _gfbf float64
			for _bagd := 0; _bagd < _fac; _bagd++ {
				_cgd := _agc.Bottom - float64(_bagd)*_eac
				if _bagd < _fac-1 {
					_beba := _cdc.NewParagraph(_eca[_bagd])
					_beba.SetFontSize(_cde * _fbbc)
					_beba.SetPos(_abe, _cgd-_eac/2-_gfbe*_fbbc)
					_cdc.Draw(_beba)
					_bfac := (_beba.Width()/1000 - _ccc) * _fbbc
					if _bfac > 0 && _bfac > _gfbf {
						_gfbf = _bfac
					}
				}
			}
			if _gfbf > 0 {
				_agc.Left += _gfbf + _fgf
				_bba = _agc.Left - _gce*_gbc/_fbe
				_gbc = _agc.Right - _agc.Left
			}
		}
		_bfgd := _bba - _gab
		_gaba := _bba + _gdae
		_gfda := _agc.Left
		_aea := _agc.Right
		for _bge := 0; _bge < _fac; _bge++ {
			_ace := _agc.Bottom - float64(_bge)*_eac
			_ggf.drawLineWithProps(_gae, _bfgd, _ace, _gaba, _ace, true)
			_ggf.drawLineWithProps(_bbdg, _gfda, _ace, _aea, _ace, true)
		}
		_cdfd := _gbc / _fbe
		_degc := _agc.Bottom - _gfba
		_ffbd := _agc.Bottom + _dea
		_ddb := _agc.Top
		_gdf := _agc.Bottom
		for _cea := 0; _cea < _ddg; _cea++ {
			_dge := _agc.Left + (_gdbd-_gce)*_cdfd
			_ggf.drawLineWithProps(_gda, _dge, _degc, _dge, _ffbd, true)
			_ggf.drawLineWithProps(_gdaa, _dge, _ddb, _dge, _gdf, true)
			if _bebg {
				_abb := _cdc.NewParagraph(_ge.FormatFloat(_gdbd, 'g', -1, 64))
				_abb.SetFontSize(_cde * _fbbc)
				_abb.SetPos(_dge-_ffc*_fbbc, _gdf+_fbb*_fbbc)
				_cdc.Draw(_abb)
			}
			_gdbd += _efb
		}
	} else {
		_dafd := _fcf / _fbe
		_gec := _agc.Left
		if _bebg {
			var _dga float64
			for _cdfdc := 0; _cdfdc < _ddg; _cdfdc++ {
				_gac := _agc.Bottom - (_gdbd-_gce)*_dafd
				_aae := _cdc.NewParagraph(_ge.FormatFloat(_gdbd, 'g', -1, 64))
				_aae.SetFontSize(_cde * _fbbc)
				_aae.SetPos(_gec-_ccc*_fbbc, _gac-_gfbe*_fbbc)
				_cdc.Draw(_aae)
				_gfdg := (_aae.Width()/1000 - _ccc) * _fbbc
				if _gfdg > 0 && _gfdg > _dga {
					_dga = _gfdg
				}
				_gdbd += _efb
			}
			if _dga > 0 {
				_agc.Left += _dga + _fgf
				_gbc = _agc.Right - _agc.Left
			}
		}
		_gdbd = _gce
		_cbfa := _agc.Left - _dea
		_fbbb := _agc.Left + _gfba
		_gec = _agc.Left
		_fbag := _agc.Right
		for _fab := 0; _fab < _ddg; _fab++ {
			_eda := _agc.Bottom - (_gdbd-_gce)*_dafd
			_ggf.drawLineWithProps(_gda, _cbfa, _eda, _fbbb, _eda, true)
			_ggf.drawLineWithProps(_gdaa, _gec, _eda, _fbag, _eda, true)
			_gdbd += _efb
		}
		_fgd := _gbc / float64(len(_eca))
		_dcab := _agc.Bottom + _gce*_fcf/_fbe
		_fbcd := _dcab - _gdae
		_aaa := _dcab + _gab
		_dfd := _agc.Top
		_abd := _agc.Bottom
		_bab := _dcab + _fbb*_fbbc
		for _aba := 0; _aba < _fac; _aba++ {
			_bcfa := _agc.Left + float64(_aba)*_fgd
			_ggf.drawLineWithProps(_gae, _bcfa, _fbcd, _bcfa, _aaa, true)
			_ggf.drawLineWithProps(_bbdg, _bcfa, _dfd, _bcfa, _abd, true)
			if _eagb && _aba < _fac-1 {
				_agcg := _cdc.NewParagraph(_eca[_aba])
				_agcg.SetFontSize(_cde * _fbbc)
				_agcg.SetPos(_bcfa+_ggbc*_fbbc, _bab)
				_cdc.Draw(_agcg)
			}
		}
	}
	return nil
}

func (_dc *creatorContext) drawBarChart(_ca *_bd.CT_BarChart, _ee *Rectangle, _fc *_bd.CT_PlotAreaChoice1) ([]*legendItem, error) {
	var _fg bool
	if _bca := _ca.BarDir; _bca != nil {
		_fg = _bca.ValAttr == _bd.ST_BarDirBar
	}
	_af := _ca.Ser
	_ba.Sort(barSerByOrder(_af))
	_dd := map[string]serCategory{}
	_be := []string{}
	_db := []*legendItem{}
	_bf := _ad.Inf(1)
	_fgb := _ad.Inf(-1)
	for _, _eec := range _af {
		var _ag string
		if _ff := _eec.Tx; _ff != nil {
			if _bfg := _ff.Choice; _bfg != nil {
				if _bfg.V != nil {
					_ag = *_bfg.V
				} else if _ggb := _bfg.StrRef; _ggb != nil {
					if _dbg := _ggb.StrCache; _dbg != nil {
						for _, _bbd := range _dbg.Pt {
							_ag = _bbd.V
						}
					}
				}
			}
		}
		if _ga := _eec.Cat; _ga != nil {
			if _bdg := _ga.Choice; _bdg != nil {
				if _gee := _bdg.StrRef; _gee != nil {
					if _eg := _gee.StrCache; _eg != nil {
						for _, _gb := range _eg.Pt {
							_dbe := _gb.V
							if _, _ea := _dd[_dbe]; !_ea {
								_dd[_dbe] = serCategory{_ece: _dbe, _aeda: []serValue{}}
								_be = append(_be, _dbe)
							}
						}
					}
				} else if _de := _bdg.NumRef; _de != nil {
					if _eb := _de.NumCache; _eb != nil {
						var _dbc string
						if _eb.FormatCode != nil {
							_dbc = *_eb.FormatCode
						}
						for _, _bag := range _eb.Pt {
							var _df string
							if _bag.FormatCodeAttr == nil {
								_df = _dbc
							} else {
								_df = *_bag.FormatCodeAttr
							}
							var _dfe string
							_bce, _eab := _ge.ParseFloat(_bag.V, 64)
							if _eab != nil {
								_dfe = _bag.V
							} else {
								_dfe = _fb.Number(_bce, _df)
							}
							if _, _dbf := _dd[_dfe]; !_dbf {
								_dd[_dfe] = serCategory{_ece: _dfe, _aeda: []serValue{}}
								_be = append(_be, _dfe)
							}
						}
					}
				}
			}
		}
		if _cfc := _eec.Val; _cfc != nil {
			if _dff := _cfc.Choice; _dff != nil {
				if _afd := _dff.NumRef; _afd != nil {
					if _geb := _afd.NumCache; _geb != nil {
						for _gef, _eeg := range _geb.Pt {
							_fbf, _eba := _ge.ParseFloat(_eeg.V, 64)
							if _eba != nil {
								_fbf = 0
								_bc.Log.Debug("\u0070a\u0072s\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", _eba)
							}
							if _fbf > _fgb {
								_fgb = _fbf
							}
							if _fbf < _bf {
								_bf = _fbf
							}
							_afda := _dd[_be[_gef]]
							_afda._aeda = append(_afda._aeda, serValue{_fbd: _ag, _aeb: _fbf, _cage: _eec.SpPr})
							_dd[_be[_gef]] = _afda
						}
					}
				}
			}
		}
		_db = append(_db, &legendItem{_acd: _ag, _afc: _eec.SpPr})
	}
	var _aff float64
	var _bee, _cd float64
	if _fgb == 0 && _bf == 0 {
		_aff = 0.2
		_cd = 0
		_bee = 1
	} else {
		var _ac float64
		if _fcg := _ad.Abs(_bf); _fgb < _fcg {
			_ac = _fcg
		} else {
			_ac = _fgb
		}
		_cfg := _ad.Pow(10, _ad.Floor(_ad.Log10(_ac)))
		_gd := _ac / _cfg
		if _gd >= 1.715 && _gd < 4.29 {
			_aff = 0.5
		} else if _gd >= 4.29 && _gd < 8.58 {
			_aff = 1
		} else {
			_aff = 2
		}
		_aff *= _cfg
		if _fgb <= 0 {
			_bee = 0
		} else {
			_bee = (_ad.Ceil(_fgb/_aff) + 1) * _aff
		}
		if _bf >= 0 {
			_cd = 0
		} else {
			_cd = (_ad.Floor(_bf/_aff) - 1) * _aff
		}
	}
	_fgg := _dc.drawAxes(_fc, _cd, _bee, _aff, _be, _ee, _fg)
	if _fgg != nil {
		return nil, _fgg
	}
	_afa := 0.0
	if _ca.GapWidth != nil {
		if _egb := _ca.GapWidth.ValAttr; _egb != nil {
			if _fd := _egb.ST_GapAmountUShort; _fd != nil {
				_afa = float64(*_fd) / 100.0
			}
		}
	}
	_dcg := _ee.Right - _ee.Left
	_dbb := _ee.Bottom - _ee.Top
	_fbfd := float64(len(_be))
	if _fg {
		_dfb := _bee / (_bee - _cd) * _dcg
		_ec := -_cd / (_bee - _cd) * _dcg
		_ceb := _ee.Left + _ec
		_eag := _dbb / _fbfd
		for _cag, _fcge := range _be {
			_dcc := _dd[_fcge]
			_aed := float64(len(_dcc._aeda)) + _afa
			_afb := _eag / _aed
			_cfge := _afb * _afa
			_cdg := _ee.Bottom - float64(_cag)*_eag - _cfge/2 - _afb
			for _, _fbc := range _dcc._aeda {
				if _fbc._aeb == 0 {
					continue
				}
				var _bfa, _ebg float64
				if _fbc._aeb > 0 {
					_ebg = _fbc._aeb / _bee * _dfb
					_dc.drawRectangleWithProps(_fbc._cage, _ceb, _cdg, _ebg, _afb, false)
				} else {
					_ebg = _fbc._aeb / _cd * _ec
					_bfa = _ceb - _ebg
					_dc.drawRectangleWithProps(_fbc._cage, _bfa, _cdg, _ebg, _afb, false)
				}
				_cdg -= _afb
			}
		}
	} else {
		_gc := _bee / (_bee - _cd) * _dbb
		_ab := -_cd / (_bee - _cd) * _dbb
		_da := _ee.Top + _gc
		_fbg := _dcg / _fbfd
		for _caga, _dccd := range _be {
			_bdd := _dd[_dccd]
			_gbg := float64(len(_bdd._aeda)) + _afa
			_ffb := _fbg / _gbg
			_fdf := _ffb * _afa
			_dffe := _ee.Left + float64(_caga)*_fbg + _fdf/2
			for _, _gad := range _bdd._aeda {
				var _dde, _ffa float64
				if _gad._aeb > 0 {
					_ffa = _gad._aeb / _bee * _gc
					_dde = _da - _ffa
					_dc.drawRectangleWithProps(_gad._cage, _dffe, _dde, _ffb, _ffa, false)
				} else {
					_ffa = _gad._aeb / _cd * _ab
					_dc.drawRectangleWithProps(_gad._cage, _dffe, _da, _ffb, _ffa, false)
				}
				_dffe += _ffb
			}
		}
	}
	return _db, nil
}

func (_cgfe *creatorContext) getPdfColorFromSolidFill(_acec *_fba.CT_SolidColorFillProperties) _gg.Color {
	if _acec == nil {
		return nil
	}
	_gddf := ""
	if _cbaf := _acec.SrgbClr; _cbaf != nil {
		_gddf = _cbaf.ValAttr
	} else if _gff := _acec.SchemeClr; _gff != nil {
		_gddf = _cec(_gff.ValAttr, _cgfe._beeag)
	}
	if _gddf == "" {
		return nil
	}
	return _gg.ColorRGBFromHex("\u0023" + _gddf)
}

func GetPageFromCreator(c *_gg.Creator) (*_ce.PdfPage, error) {
	_dfbg := _d.NewBuffer([]byte{})
	_beeg := c.Write(_dfbg)
	if _beeg != nil {
		return nil, _beeg
	}
	_gabaec := _d.NewReader(_dfbg.Bytes())
	_fed, _beeg := _ce.NewPdfReader(_gabaec)
	if _beeg != nil {
		return nil, _beeg
	}
	return _fed.GetPage(1)
}

func AdjustColorByShade(colorStr string, shade float64) string {
	var _egg, _bega, _ecec uint8
	_fbef, _ := _e.Sscanf(colorStr, "\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078", &_egg, &_bega, &_ecec)
	if _fbef != 3 {
		return ""
	}
	return _dbbg(_egg, shade) + _dbbg(_bega, shade) + _dbbg(_ecec, shade)
}

func AdjustColorByTint(colorStr string, tint float64) string {
	var _gba, _ffab, _dgc uint8
	_bafb, _ := _e.Sscanf(colorStr, "\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078", &_gba, &_ffab, &_dgc)
	if _bafb != 3 {
		return ""
	}
	return _agbb(_gba, tint) + _agbb(_ffab, tint) + _agbb(_dgc, tint)
}

var _beg = _ebfa(0.125)

func Lighten(clr float64) float64 { return 0.6 + 0.4*clr }

type ImgPart byte

type serValue struct {
	_fbd  string
	_aeb  float64
	_cage *_fba.CT_ShapeProperties
}

type barSerByOrder []*_bd.CT_BarSer

func _agf(_caad, _geee, _bgd uint8) (float64, float64, float64) {
	_aedd, _fbcb, _fcfg := float64(_caad)/255, float64(_geee)/255, float64(_bgd)/255
	_acfb := _aedd
	if _fbcb < _acfb {
		_acfb = _fbcb
	}
	if _fcfg < _acfb {
		_acfb = _fcfg
	}
	var _cdfg, _eadd bool
	_gabg := _aedd
	if _fbcb > _gabg {
		_gabg = _fbcb
		_cdfg = true
	}
	if _fcfg > _gabg {
		_gabg = _fcfg
		_cdfg = false
		_eadd = true
	}
	_fee := (_acfb + _gabg) / 2
	var _ggd float64
	if _acfb != _gabg {
		if _fee <= 0.5 {
			_ggd = (_gabg - _acfb) / (_gabg + _acfb)
		} else {
			_ggd = (_gabg - _acfb) / (2.0 - _gabg - _acfb)
		}
	}
	var _ccg float64
	if _acfb != _gabg {
		if _cdfg {
			_ccg = 2.0 + (_fcfg-_aedd)/(_gabg-_acfb)
		} else if _eadd {
			_ccg = 4.0 + (_aedd-_fbcb)/(_gabg-_acfb)
		} else {
			_ccg = (_fbcb - _fcfg) / (_gabg - _acfb)
		}
		_ccg *= 60
		if _ccg < 0 {
			_ccg += 360
		}
	}
	return _ccg, _ggd, _fee
}

func MakeBlockFromCreator(c *_gg.Creator) (*_gg.Block, error) {
	_agba, _ecg := GetPageFromCreator(c)
	if _ecg != nil {
		return nil, _ecg
	}
	_baaa, _ecg := _gg.NewBlockFromPage(_agba)
	if _ecg != nil {
		return nil, _ecg
	}
	return _baaa, nil
}

func _ebf(_dca *_bd.ChartSpace, _abg, _eagf float64, _ed *_fba.Theme, _bed bool) (*_gg.Creator, error) {
	_afaf := 1.0
	if _bed {
		_afaf = 8.0
	}
	_age := &Rectangle{}
	_adf := &Rectangle{Top: _age.Top, Bottom: _eagf - _age.Bottom, Left: _age.Left, Right: _abg - _age.Right}
	_dfbd := MakeTempCreator(_abg*_afaf+1, _eagf*_afaf+1)
	_daf := &creatorContext{_dcabc: _dfbd, _beeag: _ed, _bdda: _afaf}
	var _gaa bool
	if _eed := _dca.Chart; _eed != nil {
		_cdf := _eed.PlotArea
		if _cdf == nil {
			return nil, _ef.New("\u004e\u006f\u0020p\u006c\u006f\u0074\u0020\u0061\u0072\u0065\u0061")
		}
		_ded := &Rectangle{Top: _ebfa(10), Bottom: _adf.Bottom - _ebfa(15), Left: _ebfa(10), Right: _adf.Right - _ebfa(10)}
		var _dacb *Rectangle
		_gge := _eed.Legend
		if _gge != nil {
			_bcf := _gge.Overlay != nil && _gge.Overlay.ValAttr != nil && *_gge.Overlay.ValAttr
			if _agb := _gge.LegendPos; _agb != nil {
				switch _agb.ValAttr {
				case _bd.ST_LegendPosTr:
					if !_bcf {
						_ded = &Rectangle{Top: _ebfa(25), Bottom: _adf.Bottom - _ebfa(10), Left: _ebfa(10), Right: _adf.Right - _ebfa(25)}
					}
					_dacb = &Rectangle{Top: _ebfa(2.5), Bottom: _ebfa(22.5), Left: _adf.Right - _ebfa(22.5), Right: _adf.Right - _ebfa(2.5)}
				case _bd.ST_LegendPosT:
					_dacb = &Rectangle{Top: _ebfa(2.5), Bottom: _ebfa(7.5), Left: (_adf.Right - _adf.Left) * 0.25, Right: (_adf.Right - _adf.Left) * 0.75}
					if !_bcf {
						_ded = &Rectangle{Top: _ebfa(12.5), Bottom: _adf.Bottom - _ebfa(15), Left: _ebfa(10), Right: _adf.Right - _ebfa(5)}
					}
					_gaa = true
				case _bd.ST_LegendPosB:
					_dacb = &Rectangle{Top: _adf.Bottom - _ebfa(7.5), Bottom: _adf.Bottom - _ebfa(2.5), Left: (_adf.Right - _adf.Left) * 0.25, Right: (_adf.Right - _adf.Left) * 0.75}
					if !_bcf {
						_ded = &Rectangle{Top: _ebfa(5), Bottom: _adf.Bottom - _ebfa(15), Left: _ebfa(10), Right: _adf.Right - _ebfa(5)}
					}
					_gaa = true
				case _bd.ST_LegendPosR:
					_dacb = &Rectangle{Top: (_adf.Bottom-_adf.Top)/2 - _ebfa(10), Bottom: (_adf.Bottom-_adf.Top)/2 + _ebfa(10), Left: _adf.Right - _ebfa(22.5), Right: _adf.Right - _ebfa(2.5)}
					if !_bcf {
						_ded = &Rectangle{Top: _ebfa(5), Bottom: _adf.Bottom - _ebfa(12.5), Left: _ebfa(10), Right: _adf.Right - _ebfa(25)}
					}
				case _bd.ST_LegendPosL:
					_dacb = &Rectangle{Top: (_adf.Bottom-_adf.Top)/2 - _ebfa(10), Bottom: (_adf.Bottom-_adf.Top)/2 + _ebfa(10), Left: _ebfa(2.5), Right: _ebfa(22.5)}
					if !_bcf {
						_ded = &Rectangle{Top: _ebfa(5), Bottom: _adf.Bottom - _ebfa(12.5), Left: _ebfa(30), Right: _adf.Right - _ebfa(5)}
					}
				default:
					_dacb = &Rectangle{Top: (_adf.Bottom-_adf.Top)/2 - _ebfa(10), Bottom: (_adf.Bottom-_adf.Top)/2 + _ebfa(10), Left: _adf.Right - _ebfa(25), Right: _adf.Right - _ebfa(5)}
					if !_bcf {
						_ded = &Rectangle{Top: _ebfa(5), Bottom: _adf.Bottom - _ebfa(12.5), Left: _ebfa(100), Right: _adf.Right - _ebfa(25)}
					}
				}
			}
		}
		_ded.scale(_afaf)
		_daf.drawBorderWithProps(_cdf.SpPr, _ded, _beg)
		_ade := []*legendItem{}
		var _bcc error
		_ccf := _cdf.CChoice
		for _, _bcfc := range _cdf.Choice {
			if _dad := _bcfc.BarChart; _dad != nil {
				_ade, _bcc = _daf.drawBarChart(_dad, _ded, _ccf)
				if _bcc != nil {
					return nil, _bcc
				}
			}
		}
		if _gge != nil {
			_dacb.scale(_afaf)
			_daf.drawBorderWithProps(_gge.SpPr, _dacb, _beg)
			if len(_ade) != 0 {
				_daf.drawLegend(_dacb, _ade, _gaa)
			}
		}
	}
	_adf.scale(_afaf)
	_daf.drawBorderWithProps(_dca.SpPr, _adf, _beg)
	return _dfbd, nil
}

func RegisterFontsFromFiles(files []string) error {
	for _, _cgf := range files {
		if _fe.HasSuffix(_cgf, "\u002e\u0074\u0074\u0066") {
			_afab := _ebfg(_cgf)
			if _afab != nil {
				_bc.Log.Debug("\u0075\u006ea\u0062\u006c\u0065\u0020\u0074o\u0020\u0070\u0072\u006f\u0063e\u0073\u0073\u0020\u0061\u006e\u0064\u0020\u0072\u0065\u0067\u0069\u0073\u0074\u0065\u0072\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020\u0054\u0054\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073", _afab)
				continue
			}
		}
	}
	return nil
}

var _ffc = _ebfa(0.5)

func MakeBlockFromChartSpace(cs *_bd.ChartSpace, width, height float64, theme *_fba.Theme) (*_gg.Block, error) {
	_abf, _gfb := _ebf(cs, width, height, theme, false)
	if _gfb != nil {
		return nil, _gfb
	}
	_cc, _gfb := GetPageFromCreator(_abf)
	if _gfb != nil {
		return nil, _gfb
	}
	_fca, _gfb := _gg.NewBlockFromPage(_cc)
	if _gfb != nil {
		return nil, _gfb
	}
	return _fca, nil
}

type BorderPosition byte

func (_cfa *creatorContext) drawLineWithProps(_edg *_fba.CT_ShapeProperties, _dfbf, _cbc, _fegf, _cafg float64, _cdb bool) {
	if _edg != nil {
		if _cagd := _edg.Ln; _cagd != nil {
			_badd := _cfa.getPdfColorFromSolidFill(_cagd.SolidFill)
			if _badd == nil && _cdb {
				_badd = _gg.ColorBlack
			}
			if _badd != nil {
				var _cca float64
				if _eaa := _cagd.WAttr; _eaa != nil {
					_cca = _aa.FromEMU(int64(*_eaa))
				} else {
					_cca = _beg
				}
				DrawLine(_cfa._dcabc, _dfbf, _cbc, _fegf, _cafg, _cca, _badd)
			}
		}
	}
}

func _dgae(_dfg *_bd.CT_CatAx) (uint32, _bd.ST_AxPos, _bd.ST_TickMark, _bd.ST_TickLblPos, *_bd.CT_ChartLines, uint32, *_fba.CT_ShapeProperties, error) {
	var _fcgb, _gaf uint32
	var _dfbb _bd.ST_AxPos
	var _badc _bd.ST_TickMark
	var _dcdf *_bd.CT_ChartLines
	var _faa _bd.ST_TickLblPos
	if _dfg.AxId == nil {
		return _fcgb, _dfbb, _badc, _faa, _dcdf, _gaf, _dfg.SpPr, _ef.New("\u004e\u006f\u0020x\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044")
	} else {
		_fcgb = _dfg.AxId.ValAttr
	}
	if _dfg.AxPos == nil {
		return _fcgb, _dfbb, _badc, _faa, _dcdf, _gaf, _dfg.SpPr, _ef.New("\u004eo\u0020x\u0020\u0061\u0078\u0069\u0073 \u0070\u006fs\u0069\u0074\u0069\u006f\u006e")
	} else {
		_dfbb = _dfg.AxPos.ValAttr
	}
	if _dfg.MajorTickMark != nil {
		_badc = _dfg.MajorTickMark.ValAttr
	}
	if _dfg.TickLblPos != nil {
		_faa = _dfg.TickLblPos.ValAttr
	}
	if _dfg.CrossAx == nil {
		return _fcgb, _dfbb, _badc, _faa, _dcdf, _gaf, _dfg.SpPr, _ef.New("\u004e\u006f \u0063\u0072\u006fs\u0073\u0020\u0061\u0078\u0069\u0073\u0020\u0049\u0044")
	} else {
		_gaf = _dfg.CrossAx.ValAttr
	}
	_dcdf = _dfg.MajorGridlines
	return _fcgb, _dfbb, _badc, _faa, _dcdf, _gaf, _dfg.SpPr, nil
}

func GetOpacityFromColorTransform(trs []*_fba.EG_ColorTransform) float64 {
	for _, _eegf := range trs {
		if _eegf != nil {
			if _gbge := _eegf.Alpha; _gbge != nil {
				if _dfdb := _gbge.ValAttr.ST_PositiveFixedPercentageDecimal; _dfdb != nil {
					return float64(*_dfdb) / 100000
				}
			}
		}
	}
	return 1.0
}

func DrawLine(c *_gg.Creator, x0, y0, x1, y1, width float64, color _gg.Color) {
	if color == nil {
		return
	}
	_aebb := c.NewLine(x0, y0, x1, y1)
	_aebb.SetLineWidth(width)
	_aebb.SetColor(color)
	c.Draw(_aebb)
}

var StdFontsMap = map[string][]string{"\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a": []string{"\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a", "\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0042\u006f\u006c\u0064", "\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u004f\u0062l\u0069\u0071\u0075\u0065", "H\u0065\u006c\u0076\u0065ti\u0063a\u002d\u0042\u006f\u006c\u0064O\u0062\u006c\u0069\u0071\u0075\u0065"}, "\u0043o\u0075\u0072\u0069\u0065\u0072": []string{"\u0043o\u0075\u0072\u0069\u0065\u0072", "\u0043\u006f\u0075r\u0069\u0065\u0072\u002d\u0042\u006f\u006c\u0064", "\u0043o\u0075r\u0069\u0065\u0072\u002d\u004f\u0062\u006c\u0069\u0071\u0075\u0065", "\u0043\u006f\u0075\u0072ie\u0072\u002d\u0042\u006f\u006c\u0064\u004f\u0062\u006c\u0069\u0071\u0075\u0065"}, "\u0054i\u006de\u0073\u0020\u004e\u0065\u0077\u0020\u0052\u006f\u006d\u0061\u006e": []string{"T\u0069\u006d\u0065\u0073\u002d\u0052\u006f\u006d\u0061\u006e", "\u0054\u0069\u006d\u0065\u0073\u002d\u0042\u006f\u006c\u0064", "\u0054\u0069\u006de\u0073\u002d\u0049\u0074\u0061\u006c\u0069\u0063", "\u0054\u0069m\u0065\u0073\u002dB\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063"}, "\u0064e\u0066\u0061\u0075\u006c\u0074": []string{"\u0048e\u006c\u0076\u0065\u0074\u0069\u0063a", "\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061-\u0042\u006f\u006c\u0064", "\u0048\u0065\u006c\u0076\u0065\u0074\u0069\u0063\u0061\u002d\u004f\u0062l\u0069\u0071\u0075\u0065", "H\u0065\u006c\u0076\u0065ti\u0063a\u002d\u0042\u006f\u006c\u0064O\u0062\u006c\u0069\u0071\u0075\u0065"}}

func _gfaf(_dbce _fba.ST_SchemeColorVal, _fgae *_fba.Theme) string {
	if _cdad := _fgae.ThemeElements; _cdad != nil {
		if _ecc := _cdad.ClrScheme; _ecc != nil {
			switch _dbce {
			case _fba.ST_SchemeColorValLt1:
				return GetColorStringFromDmlColor(_ecc.Lt1)
			case _fba.ST_SchemeColorValDk1, _fba.ST_SchemeColorValTx1:
				return GetColorStringFromDmlColor(_ecc.Dk1)
			case _fba.ST_SchemeColorValLt2:
				return GetColorStringFromDmlColor(_ecc.Lt2)
			case _fba.ST_SchemeColorValDk2:
				return GetColorStringFromDmlColor(_ecc.Dk2)
			case _fba.ST_SchemeColorValAccent1:
				return GetColorStringFromDmlColor(_ecc.Accent1)
			case _fba.ST_SchemeColorValAccent2:
				return GetColorStringFromDmlColor(_ecc.Accent2)
			case _fba.ST_SchemeColorValAccent3:
				return GetColorStringFromDmlColor(_ecc.Accent3)
			case _fba.ST_SchemeColorValAccent4:
				return GetColorStringFromDmlColor(_ecc.Accent4)
			case _fba.ST_SchemeColorValAccent5:
				return GetColorStringFromDmlColor(_ecc.Accent5)
			case _fba.ST_SchemeColorValAccent6:
				return GetColorStringFromDmlColor(_ecc.Accent6)
			}
		}
	}
	return ""
}

type creatorContext struct {
	_dcabc *_gg.Creator
	_beeag *_fba.Theme
	_bdda  float64
}

const (
	BorderPositionTop    BorderPosition = 0
	BorderPositionLeft   BorderPosition = 1
	BorderPositionBottom BorderPosition = 2
	BorderPositionRight  BorderPosition = 3
)

func TwipsFromPoints(points float64) float64 { return points / _aa.Twips }

var _fgf = _ebfa(2)

func _agbb(_bea uint8, _bda float64) string {
	_ggg := float64(_bea)
	var _aag float64
	if _bda < 0 {
		_aag = _ggg * (1 + _bda)
	} else {
		_aag = _ggg + (255-_ggg)*_bda
	}
	return _e.Sprintf("\u0025\u0030\u0032\u0078", int(_aag))
}

type fontsMap struct {
	_adec *_b.Mutex
	_cgdg map[string]map[FontStyle]*_ce.PdfFont
}

func FromSTPercentage(st *_fba.ST_Percentage) float64 {
	if _cgce := st.ST_PercentageDecimal; _cgce != nil {
		return float64(*_cgce) / 100000
	}
	return 0
}

var _ebb = _ebfa(1)

func GetImage(c *_gg.Creator, goImg _c.Image, imgHeight, imgWidth, left, top, dividerX, dividerY float64, part ImgPart) (*_gg.Image, error) {
	if goImg == nil {
		return nil, nil
	}
	_dda := goImg.Bounds().Size()
	_afba := _dda.X
	_ebe := _dda.Y
	if dividerX != 0 {
		dividerX = dividerX / imgWidth * float64(_afba)
	}
	if dividerY != 0 {
		dividerY = dividerY / imgHeight * float64(_ebe)
	}
	var _ddbf _c.Rectangle
	switch part {
	case ImgPart_t:
		_ddbf = _c.Rect(0, 0, _afba, int(dividerY))
	case ImgPart_b:
		_ddbf = _c.Rect(0, int(dividerY), _afba, _ebe)
	case ImgPart_l:
		_ddbf = _c.Rect(0, 0, int(dividerX), _ebe)
	case ImgPart_r:
		_ddbf = _c.Rect(int(dividerX), 0, _afba, _ebe)
	case ImgPart_lt:
		_ddbf = _c.Rect(0, 0, int(dividerX), int(dividerY))
	case ImgPart_rt:
		_ddbf = _c.Rect(int(dividerX), 0, _afba, int(dividerY))
	case ImgPart_lb:
		_ddbf = _c.Rect(0, int(dividerY), int(dividerX), _ebe)
	case ImgPart_rb:
		_ddbf = _c.Rect(int(dividerX), int(dividerY), _afba, _ebe)
	default:
		_ddbf = _c.Rect(0, 0, _afba, _ebe)
	}
	_ecd := CropImageByRect(goImg, _ddbf)
	_cgfc, _cac := c.NewImageFromGoImage(_ecd)
	if _cac != nil {
		return nil, _cac
	}
	_cgfc.Scale(imgWidth/float64(_afba), imgHeight/float64(_ebe))
	_cgfc.SetPos(left, top)
	return _cgfc, nil
}

func IsNoSpaceLanguage(symbol string) bool {
	for _, _cdab := range symbol {
		if _f.Is(_f.Han, _cdab) {
			return true
		}
	}
	return false
}

func (_gcg FontStyle) String() string {
	return []string{"\u0052e\u0067\u0075\u006c\u0061\u0072", "\u0042\u006f\u006c\u0064", "\u0049\u0074\u0061\u006c\u0069\u0063", "\u0042\u006f\u006c\u0064\u0049\u0074\u0061\u006c\u0069\u0063"}[int(_gcg)]
}

func (_cfe barSerByOrder) Less(i, j int) bool {
	return _cfe[i].Order.ValAttr < _cfe[j].Order.ValAttr
}

func AdjustColor(colorStr string, EG_ColorTransform []*_fba.EG_ColorTransform) string {
	for _, _cfb := range EG_ColorTransform {
		if _caa := _cfb.Tint; _caa != nil {
			if _ddf := _caa.ValAttr.ST_PositiveFixedPercentageDecimal; _ddf != nil {
				colorStr = AdjustColorByTint(colorStr, float64(*_ddf)/100000)
			}
		}
		if _ffag := _cfb.Shade; _ffag != nil {
			if _ebbf := _ffag.ValAttr.ST_PositiveFixedPercentageDecimal; _ebbf != nil {
				colorStr = AdjustColorByShade(colorStr, float64(*_ebbf)/100000)
			}
		}
		if _bgc := _cfb.LumMod; _bgc != nil {
			if _aabb := _bgc.ValAttr.ST_PercentageDecimal; _aabb != nil {
				colorStr = AdjustColorByLumMod(colorStr, float64(*_aabb)/100000)
			}
		}
	}
	return colorStr
}

func (_dg barSerByOrder) Len() int { return len(_dg) }

func PointsFromTwips(twips int64) float64 {
	return float64(int64(float64(twips)*_aa.Twips*10+0.5)) / 10
}

const _cde = 6.0

func _ebfa(_geea float64) float64 { return _geea * _aa.Millimeter }

func MakeTempCreator(width, height float64) *_gg.Creator {
	_cgc := _gg.New()
	_cgc.SetPageSize(_gg.PageSize{width, height})
	_cgc.SetPageMargins(0, 0, 0, 0)
	return _cgc
}

func GetDataFromXfrm(xfrm *_fba.CT_Transform2D) (float64, float64, float64, float64) {
	var _cfbb, _dgfb, _cdcc, _fgaa float64
	if _abgd := xfrm.Off; _abgd != nil {
		_cfbb = _aa.FromEMU(FromSTCoordinate(_abgd.XAttr))
		_dgfb = _aa.FromEMU(FromSTCoordinate(_abgd.YAttr))
	}
	if _dafb := xfrm.Ext; _dafb != nil {
		_cdcc = _aa.FromEMU(_dafb.CxAttr)
		_fgaa = _aa.FromEMU(_dafb.CyAttr)
	}
	return _cfbb, _dgfb, _cdcc, _fgaa
}

func (_dbbc *creatorContext) drawRectangleWithProps(_dba *_fba.CT_ShapeProperties, _bdgf, _bbbf, _ccb, _gdg float64, _gdcg bool) {
	_dbcd := _dbbc._dcabc.NewRectangle(_bdgf, _bbbf, _ccb, _gdg)
	if _dba == nil {
		if _gdcg {
			_dbcd.SetBorderWidth(_beg)
		} else {
			return
		}
	} else {
		_gfe := _dbbc.getPdfColorFromSolidFill(_dba.SolidFill)
		if _gfe != nil {
			_dbcd.SetFillColor(_gfe)
		}
		if _ggfg := _dba.Ln; _ggfg != nil {
			if _cbbe := _ggfg.WAttr; _cbbe != nil {
				_cfac := _aa.FromEMU(int64(*_cbbe))
				_dbcd.SetBorderWidth(_cfac)
				if _aac := _ggfg.SolidFill; _aac != nil {
					_efag := _dbbc.getPdfColorFromSolidFill(_aac)
					if _efag != nil {
						_dbcd.SetBorderColor(_efag)
					}
				}
			} else {
				_dbcd.SetBorderWidth(0)
			}
		}
	}
	_dbbc._dcabc.Draw(_dbcd)
}

func _ebfg(_fbcdf string) error {
	if !_fe.HasSuffix(_fbcdf, "\u002e\u0074\u0074\u0066") {
		_bc.Log.Debug("\u0055\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u0066\u0069\u006c\u0065\u0020\u0066\u006f\u0072ma\u0074\u002e")
		return _e.Errorf("\u0055\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020f\u006f\u006e\u0074\u0020\u0066\u0069l\u0065\u0020\u0066\u006f\u0072m\u0061\u0074\u002c\u0020\u0063\u0075\u0072\u0072\u0065\u006e\u0074\u006cy\u0020\u006f\u006e\u006c\u0079\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0054T\u0046\u0020\u0066\u006f\u006e\u0074\u0020\u0066i\u006c\u0065\u002e")
	}
	_cad, _gdbf := _bb.ParseFile(_fbcdf)
	if _gdbf != nil {
		_bc.Log.Debug("\u0043a\u006e\u006e\u006f\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020T\u0054\u0046\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073", _gdbf)
		return _gdbf
	}
	_bddge, _gdbf := _ce.NewCompositePdfFontFromTTFFile(_fbcdf)
	if _gdbf != nil {
		return _gdbf
	}
	_efa := _cad.GetNameRecords()
	for _, _aaaa := range _efa {
		_dec := _aaaa[1]
		if _dec == "" {
			return _e.Errorf("\u004e\u006f\u0020\u0066\u006fn\u0074\u0020\u0066\u0061\u006d\u0069\u006c\u0079\u0020\u0069\u006e\u0066\u006fr\u006d\u0061\u0074\u0069\u006f\u006e\u0020\u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073", _fbcdf)
		}
		_add := make([]byte, 0)
		for _feeb := 0; _feeb < len(_dec); _feeb++ {
			if _dec[_feeb] == 39 || _dec[_feeb] == 92 {
				continue
			}
			_fcd := 4
			if _feeb+_fcd < len(_dec) {
				if _dec[_feeb:_feeb+_fcd] == "\u0000" {
					_feeb = _feeb + _fcd + 1
					continue
				}
			}
			_add = append(_add, _dec[_feeb])
		}
		_dec = _fe.Replace(string(_add), "\u0078\u0030\u0030", "", -1)
		_babc := _aaaa[2]
		if _babc == "" {
			return _e.Errorf("N\u006f\u0020\u0073\u0074\u0079\u006ce\u0020\u0069\u006e\u0066\u006f\u0072m\u0061\u0074\u0069\u006f\u006e\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020%\u0073", _fbcdf)
		}
		_add = make([]byte, 0)
		for _acef := 0; _acef < len(_babc); _acef++ {
			if _babc[_acef] == 39 || _babc[_acef] == 92 {
				continue
			}
			_fbbg := 4
			if _acef+_fbbg < len(_babc) {
				if _babc[_acef:_acef+_fbbg] == "\u0000" {
					_acef = _acef + _fbbg + 1
					continue
				}
			}
			_add = append(_add, _babc[_acef])
		}
		_babc = _fe.Replace(string(_add), "\u0078\u0030\u0030", "", -1)
		RegisterFont(_dec, _cbac[_babc], _bddge)
	}
	return nil
}

const (
	ImgPart_whole ImgPart = 0
	ImgPart_t     ImgPart = 1
	ImgPart_b     ImgPart = 2
	ImgPart_l     ImgPart = 3
	ImgPart_r     ImgPart = 4
	ImgPart_lt    ImgPart = 5
	ImgPart_rt    ImgPart = 6
	ImgPart_lb    ImgPart = 7
	ImgPart_rb    ImgPart = 8
)

func _bcac(_afeb, _bgaf, _fea float64) (uint8, uint8, uint8) {
	var _edad float64
	if _fea < 0.5 {
		_edad = _fea * (1 + _bgaf)
	} else {
		_edad = _fea + _bgaf - _fea*_bgaf
	}
	_efc := _fea*2 - _edad
	_afeb /= 360.0
	_fgge := _gggf(_afeb + 1.0/3.0)
	_gcd := _gggf(_afeb)
	_abaf := _gggf(_afeb - 1.0/3.0)
	_bdb := _aebf(_fgge, _edad, _efc)
	_acc := _aebf(_gcd, _edad, _efc)
	_egdc := _aebf(_abaf, _edad, _efc)
	return uint8(255 * _bdb), uint8(255 * _acc), uint8(255 * _egdc)
}

const DefaultFontSize = 12.0

func _gggf(_eadc float64) float64 {
	if _eadc < 0 {
		_eadc += float64(-int(_eadc) + 1)
	} else if _eadc > 1 {
		_eadc -= float64(int(_eadc))
	}
	return _eadc
}

type Rectangle struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
}
