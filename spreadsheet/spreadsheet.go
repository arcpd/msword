package spreadsheet

import (
	_bdg "archive/zip"
	_gb "bytes"
	_fb "errors"
	_cg "fmt"
	_b "github.com/arcpd/msword"
	_ec "github.com/arcpd/msword/chart"
	_fbf "github.com/arcpd/msword/color"
	_dcb "github.com/arcpd/msword/common"
	_ca "github.com/arcpd/msword/common/logger"
	_ae "github.com/arcpd/msword/common/tempstorage"
	// _ac "github.com/arcpd/msword/internal/license"
	_bg "github.com/arcpd/msword/measurement"
	_acg "github.com/arcpd/msword/schema/soo/dml"
	_fe "github.com/arcpd/msword/schema/soo/dml/chart"
	_ag "github.com/arcpd/msword/schema/soo/dml/spreadsheetDrawing"
	_dfc "github.com/arcpd/msword/schema/soo/pkg/relationships"
	_bbg "github.com/arcpd/msword/schema/soo/sml"
	_ce "github.com/arcpd/msword/spreadsheet/format"
	_eg "github.com/arcpd/msword/spreadsheet/formula"
	_df "github.com/arcpd/msword/spreadsheet/reference"
	_gd "github.com/arcpd/msword/spreadsheet/update"
	_bb "github.com/arcpd/msword/vmldrawing"
	_dd "github.com/arcpd/msword/zippkg"
	_bf "image"
	_d "image/jpeg"
	_ee "io"
	_f "math"
	_gbg "math/big"
	_bd "os"
	_a "path"
	_g "path/filepath"
	_ge "reflect"
	_dc "regexp"
	_c "sort"
	_dg "strconv"
	_ba "strings"
	_be "time"
)

// AddString adds a string to the shared string cache.
func (_ffbb SharedStrings) AddString(v string) int {
	if _fdfe, _cafg := _ffbb._ffcg[v]; _cafg {
		return _fdfe
	}
	_bcfge := _bbg.NewCT_Rst()
	_bcfge.T = _b.String(v)
	_ffbb._ddec.Si = append(_ffbb._ddec.Si, _bcfge)
	_gfga := len(_ffbb._ddec.Si) - 1
	_ffbb._ffcg[v] = _gfga
	_ffbb._ddec.CountAttr = _b.Uint32(uint32(len(_ffbb._ddec.Si)))
	_ffbb._ddec.UniqueCountAttr = _ffbb._ddec.CountAttr
	return _gfga
}

// BottomRight is a no-op.
func (_ea AbsoluteAnchor) BottomRight() CellMarker { return CellMarker{} }

// X returns the inner wrapped XML type.
func (_caf Comments) X() *_bbg.Comments { return _caf._fcg }
func (_cdea Font) SetColor(c _fbf.Color) {
	_abdf := _bbg.NewCT_Color()
	_cfff := "\u0066\u0066" + *c.AsRGBString()
	_abdf.RgbAttr = &_cfff
	_cdea._bdbb.Color = []*_bbg.CT_Color{_abdf}
}

// SetHidden marks the defined name as hidden.
func (_dgg DefinedName) SetLocalSheetID(id uint32) { _dgg._dfced.LocalSheetIdAttr = _b.Uint32(id) }

// Column returns the cell column
func (_fdb Cell) Column() (string, error) {
	_aeb, _ffc := _df.ParseCellReference(_fdb.Reference())
	if _ffc != nil {
		return "", _ffc
	}
	return _aeb.Column, nil
}

// SetValues sets the possible values. This is incompatible with SetRange.
func (_bdec DataValidationList) SetValues(values []string) {
	_bdec._cfg.Formula1 = _b.String("\u0022" + _ba.Join(values, "\u002c") + "\u0022")
	_bdec._cfg.Formula2 = _b.String("\u0030")
}
func (_cagc StandardFormat) String() string {
	switch {
	case 0 <= _cagc && _cagc <= 4:
		return _dfdca[_abgb[_cagc]:_abgb[_cagc+1]]
	case 9 <= _cagc && _cagc <= 22:
		_cagc -= 9
		return _bacb[_fecaa[_cagc]:_fecaa[_cagc+1]]
	case 37 <= _cagc && _cagc <= 40:
		_cagc -= 37
		return _ggca[_gadf[_cagc]:_gadf[_cagc+1]]
	case 45 <= _cagc && _cagc <= 49:
		_cagc -= 45
		return _fcdf[_eaag[_cagc]:_eaag[_cagc+1]]
	default:
		return _cg.Sprintf("\u0053t\u0061n\u0064\u0061\u0072\u0064\u0046o\u0072\u006da\u0074\u0028\u0025\u0064\u0029", _cagc)
	}
}

// SetUnderline controls if the run is underlined.
func (_bfdb RichTextRun) SetUnderline(u _bbg.ST_UnderlineValues) {
	_bfdb.ensureRpr()
	_bfdb._ffae.RPr.U = _bbg.NewCT_UnderlineProperty()
	_bfdb._ffae.RPr.U.ValAttr = u
}

// DataBarScale is a colored scale that fills the cell with a background
// gradeint depending on the value.
type DataBarScale struct{ _gdb *_bbg.CT_DataBar }

// HasFormula returns true if the cell has an asoociated formula.
func (_fdgd Cell) HasFormula() bool { return _fdgd._age.F != nil }

// IsError returns true if the cell is an error type cell.
func (_gefb Cell) IsError() bool { return _gefb._age.TAttr == _bbg.ST_CellTypeE }

// AddRow adds a new row to a sheet.  You can mix this with numbered rows,
// however it will get confusing. You should prefer to use either automatically
// numbered rows with AddRow or manually numbered rows with Row/AddNumberedRow
func (_dfdg *Sheet) AddRow() Row {
	_gcbb := uint32(0)
	_gdbc := uint32(len(_dfdg._ecgc.SheetData.Row))
	if _gdbc > 0 && _dfdg._ecgc.SheetData.Row[_gdbc-1].RAttr != nil && *_dfdg._ecgc.SheetData.Row[_gdbc-1].RAttr == _gdbc {
		return _dfdg.addNumberedRowFast(_gdbc + 1)
	}
	for _, _aefe := range _dfdg._ecgc.SheetData.Row {
		if _aefe.RAttr != nil && *_aefe.RAttr > _gcbb {
			_gcbb = *_aefe.RAttr
		}
	}
	return _dfdg.AddNumberedRow(_gcbb + 1)
}

// SetStringByID sets the cell type to string, and the value a string in the
// shared strings table.
func (_ege Cell) SetStringByID(id int) {
	_ege._cea.ensureSharedStringsRelationships()
	_ege.clearValue()
	_ege._age.V = _b.String(_dg.Itoa(id))
	_ege._age.TAttr = _bbg.ST_CellTypeS
}

type Table struct{ _febfb *_bbg.Table }

func (_fcgf *Workbook) onNewRelationship(_agec *_dd.DecodeMap, _ceab, _cbdgb string, _ebaf []*_bdg.File, _fbea *_dfc.Relationship, _ccca _dd.Target) error {
	_cdceb := _b.DocTypeSpreadsheet
	switch _cbdgb {
	case _b.OfficeDocumentType:
		_fcgf._bdff = _bbg.NewWorkbook()
		_agec.AddTarget(_ceab, _fcgf._bdff, _cbdgb, 0)
		_fcgf._fdae = _dcb.NewRelationships()
		_agec.AddTarget(_dd.RelationsPathFor(_ceab), _fcgf._fdae.X(), _cbdgb, 0)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, 0)
	case _b.CorePropertiesType:
		_agec.AddTarget(_ceab, _fcgf.CoreProperties.X(), _cbdgb, 0)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, 0)
	case _b.CustomPropertiesType:
		_agec.AddTarget(_ceab, _fcgf.CustomProperties.X(), _cbdgb, 0)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, 0)
	case _b.ExtendedPropertiesType:
		_agec.AddTarget(_ceab, _fcgf.AppProperties.X(), _cbdgb, 0)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, 0)
	case _b.WorksheetType:
		_fgeb := _bbg.NewWorksheet()
		_bdbdg := uint32(len(_fcgf._cfgf))
		_fcgf._cfgf = append(_fcgf._cfgf, _fgeb)
		_agec.AddTarget(_ceab, _fgeb, _cbdgb, _bdbdg)
		_fcgac := _dcb.NewRelationships()
		_agec.AddTarget(_dd.RelationsPathFor(_ceab), _fcgac.X(), _cbdgb, 0)
		_fcgf._fgdcb = append(_fcgf._fgdcb, _fcgac)
		_fcgf._cffde = append(_fcgf._cffde, nil)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, len(_fcgf._cfgf))
	case _b.StylesType:
		_fcgf.StyleSheet = NewStyleSheet(_fcgf)
		_agec.AddTarget(_ceab, _fcgf.StyleSheet.X(), _cbdgb, 0)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, 0)
	case _b.ThemeType:
		_ccecf := _acg.NewTheme()
		_fcgf._cbae = append(_fcgf._cbae, _ccecf)
		_agec.AddTarget(_ceab, _ccecf, _cbdgb, 0)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, len(_fcgf._cbae))
	case _b.SharedStringsType:
		_fcgf.SharedStrings = NewSharedStrings()
		_agec.AddTarget(_ceab, _fcgf.SharedStrings.X(), _cbdgb, 0)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, 0)
	case _b.ThumbnailType:
		for _cfgd, _cccg := range _ebaf {
			if _cccg == nil {
				continue
			}
			if _cccg.Name == _ceab {
				_afeb, _bcacg := _cccg.Open()
				if _bcacg != nil {
					return _cg.Errorf("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0068\u0075m\u0062\u006e\u0061i\u006c:\u0020\u0025\u0073", _bcacg)
				}
				_fcgf.Thumbnail, _, _bcacg = _bf.Decode(_afeb)
				_afeb.Close()
				if _bcacg != nil {
					return _cg.Errorf("\u0065\u0072\u0072\u006fr\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020t\u0068u\u006d\u0062\u006e\u0061\u0069\u006c\u003a \u0025\u0073", _bcacg)
				}
				_ebaf[_cfgd] = nil
			}
		}
	case _b.ImageType:
		for _efbec, _ffdce := range _fcgf._bcae {
			_ecgfa := _a.Clean(_ceab)
			if _ecgfa == _efbec {
				_fbea.TargetAttr = _ffdce
				return nil
			}
		}
		_ebge := _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, len(_fcgf.Images)+1)
		for _efbeg, _gcfeca := range _ebaf {
			if _gcfeca == nil {
				continue
			}
			if _gcfeca.Name == _a.Clean(_ceab) {
				_dbbd, _fgda := _dd.ExtractToDiskTmp(_gcfeca, _fcgf.TmpPath)
				if _fgda != nil {
					return _fgda
				}
				_cdgcd, _fgda := _dcb.ImageFromStorage(_dbbd)
				if _fgda != nil {
					return _fgda
				}
				_eafg := _dcb.MakeImageRef(_cdgcd, &_fcgf.DocBase, _fcgf._fdae)
				_eafg.SetTarget(_ebge)
				_fcgf._bcae[_gcfeca.Name] = _ebge
				_fcgf.Images = append(_fcgf.Images, _eafg)
				_ebaf[_efbeg] = nil
			}
		}
		_fbea.TargetAttr = _ebge
	case _b.DrawingType:
		_ecfca := _ag.NewWsDr()
		_edbc := uint32(len(_fcgf._eded))
		_agec.AddTarget(_ceab, _ecfca, _cbdgb, _edbc)
		_fcgf._eded = append(_fcgf._eded, _ecfca)
		_fbca := _dcb.NewRelationships()
		_agec.AddTarget(_dd.RelationsPathFor(_ceab), _fbca.X(), _cbdgb, _edbc)
		_fcgf._ccag = append(_fcgf._ccag, _fbca)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, len(_fcgf._eded))
	case _b.VMLDrawingType:
		_fgfd := _bb.NewContainer()
		_ddca := uint32(len(_fcgf._fdbb))
		_agec.AddTarget(_ceab, _fgfd, _cbdgb, _ddca)
		_fcgf._fdbb = append(_fcgf._fdbb, _fgfd)
	case _b.CommentsType:
		_fcgf._cffde[_ccca.Index] = _bbg.NewComments()
		_agec.AddTarget(_ceab, _fcgf._cffde[_ccca.Index], _cbdgb, _ccca.Index)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, len(_fcgf._cffde))
	case _b.ChartType:
		_acfe := _fe.NewChartSpace()
		_adffd := uint32(len(_fcgf._ccbc))
		_agec.AddTarget(_ceab, _acfe, _cbdgb, _adffd)
		_fcgf._ccbc = append(_fcgf._ccbc, _acfe)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, len(_fcgf._ccbc))
		_fcgf._aabe[_fbea.TargetAttr] = _acfe
	case _b.TableType:
		_dcffe := _bbg.NewTable()
		_eeab := uint32(len(_fcgf._cgdd))
		_agec.AddTarget(_ceab, _dcffe, _cbdgb, _eeab)
		_fcgf._cgdd = append(_fcgf._cgdd, _dcffe)
		_fbea.TargetAttr = _b.RelativeFilename(_cdceb, _ccca.Typ, _cbdgb, len(_fcgf._cgdd))
	default:
		_ca.Log.Debug("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0072\u0065\u006c\u0061\u0074\u0069o\u006e\u0073\u0068\u0069\u0070\u0020\u0025\u0073\u0020\u0025\u0073", _ceab, _cbdgb)
	}
	return nil
}

// DefinedName is a named range, formula, etc.
type DefinedName struct{ _dfced *_bbg.CT_DefinedName }

// RemoveFont removes a font from the style sheet.  It *does not* update styles that refer
// to this font.
func (_ecggb StyleSheet) RemoveFont(f Font) error {
	for _eegg, _gcbba := range _ecggb._fgbg.Fonts.Font {
		if _gcbba == f.X() {
			_ecggb._fgbg.Fonts.Font = append(_ecggb._fgbg.Fonts.Font[:_eegg], _ecggb._fgbg.Fonts.Font[_eegg+1:]...)
			return nil
		}
	}
	return _fb.New("\u0066\u006f\u006e\u0074\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064")
}

// Save writes the workbook out to a writer in the zipped xlsx format.
func (_ecfcg *Workbook) Save(w _ee.Writer) error {
	const _dcgf = "\u0073\u0070\u0072\u0065ad\u0073\u0068\u0065\u0065\u0074\u003a\u0077\u0062\u002e\u0053\u0061\u0076\u0065"
	// if !_ac.GetLicenseKey().IsLicensed() && !_cggde {
	// 	_cg.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
	// 	_cg.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
	// 	return _fb.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	// }
	// if len(_ecfcg._edac) == 0 {
	// 	_edab, _gebb := _ac.GenRefId("\u0073\u0077")
	// 	if _gebb != nil {
	// 		_ca.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gebb)
	// 		return _gebb
	// 	}
	// 	_ecfcg._edac = _edab
	// }
	// if _cfbb := _ac.Track(_ecfcg._edac, _dcgf); _cfbb != nil {
	// 	_ca.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cfbb)
	// 	return _cfbb
	// }
	_deef := _bdg.NewWriter(w)
	defer _deef.Close()
	_dddcd := _b.DocTypeSpreadsheet
	if _fdgb := _dd.MarshalXML(_deef, _b.BaseRelsFilename, _ecfcg.Rels.X()); _fdgb != nil {
		return _fdgb
	}
	if _edaf := _dd.MarshalXMLByType(_deef, _dddcd, _b.ExtendedPropertiesType, _ecfcg.AppProperties.X()); _edaf != nil {
		return _edaf
	}
	if _gcaac := _dd.MarshalXMLByType(_deef, _dddcd, _b.CorePropertiesType, _ecfcg.CoreProperties.X()); _gcaac != nil {
		return _gcaac
	}
	_abagf := _b.AbsoluteFilename(_dddcd, _b.OfficeDocumentType, 0)
	if _abdbd := _dd.MarshalXML(_deef, _abagf, _ecfcg._bdff); _abdbd != nil {
		return _abdbd
	}
	if _dggd := _dd.MarshalXML(_deef, _dd.RelationsPathFor(_abagf), _ecfcg._fdae.X()); _dggd != nil {
		return _dggd
	}
	if _caee := _dd.MarshalXMLByType(_deef, _dddcd, _b.StylesType, _ecfcg.StyleSheet.X()); _caee != nil {
		return _caee
	}
	for _adee, _aecc := range _ecfcg._cbae {
		if _ggdf := _dd.MarshalXMLByTypeIndex(_deef, _dddcd, _b.ThemeType, _adee+1, _aecc); _ggdf != nil {
			return _ggdf
		}
	}
	for _cdab, _fgcb := range _ecfcg._cfgf {
		_fgcb.Dimension.RefAttr = Sheet{_ecfcg, nil, _fgcb}.Extents()
		_cgaf := _b.AbsoluteFilename(_dddcd, _b.WorksheetType, _cdab+1)
		_dd.MarshalXML(_deef, _cgaf, _fgcb)
		_dd.MarshalXML(_deef, _dd.RelationsPathFor(_cgaf), _ecfcg._fgdcb[_cdab].X())
	}
	if _dbgb := _dd.MarshalXMLByType(_deef, _dddcd, _b.SharedStringsType, _ecfcg.SharedStrings.X()); _dbgb != nil {
		return _dbgb
	}
	if _ecfcg.CustomProperties.X() != nil {
		if _fgff := _dd.MarshalXMLByType(_deef, _dddcd, _b.CustomPropertiesType, _ecfcg.CustomProperties.X()); _fgff != nil {
			return _fgff
		}
	}
	if _ecfcg.Thumbnail != nil {
		_bgcde := _b.AbsoluteFilename(_dddcd, _b.ThumbnailType, 0)
		_cdccf, _bdcf := _deef.Create(_bgcde)
		if _bdcf != nil {
			return _bdcf
		}
		if _gebcf := _d.Encode(_cdccf, _ecfcg.Thumbnail, nil); _gebcf != nil {
			return _gebcf
		}
	}
	for _cbdg, _afff := range _ecfcg._ccbc {
		_dbaf := _b.AbsoluteFilename(_dddcd, _b.ChartType, _cbdg+1)
		_dd.MarshalXML(_deef, _dbaf, _afff)
	}
	for _egff, _efgb := range _ecfcg._cgdd {
		_dcded := _b.AbsoluteFilename(_dddcd, _b.TableType, _egff+1)
		_dd.MarshalXML(_deef, _dcded, _efgb)
	}
	for _egeg, _dafg := range _ecfcg._eded {
		_ceaf := _b.AbsoluteFilename(_dddcd, _b.DrawingType, _egeg+1)
		_dd.MarshalXML(_deef, _ceaf, _dafg)
		if !_ecfcg._ccag[_egeg].IsEmpty() {
			_dd.MarshalXML(_deef, _dd.RelationsPathFor(_ceaf), _ecfcg._ccag[_egeg].X())
		}
	}
	for _bgce, _cacb := range _ecfcg._fdbb {
		_dd.MarshalXML(_deef, _b.AbsoluteFilename(_dddcd, _b.VMLDrawingType, _bgce+1), _cacb)
	}
	for _acgbg, _gcef := range _ecfcg.Images {
		if _abadb := _dcb.AddImageToZip(_deef, _gcef, _acgbg+1, _b.DocTypeSpreadsheet); _abadb != nil {
			return _abadb
		}
	}
	if _fccea := _dd.MarshalXML(_deef, _b.ContentTypesFilename, _ecfcg.ContentTypes.X()); _fccea != nil {
		return _fccea
	}
	for _befc, _gfbde := range _ecfcg._cffde {
		if _gfbde == nil {
			continue
		}
		_dd.MarshalXML(_deef, _b.AbsoluteFilename(_dddcd, _b.CommentsType, _befc+1), _gfbde)
	}
	if _dcfa := _ecfcg.WriteExtraFiles(_deef); _dcfa != nil {
		return _dcfa
	}
	return _deef.Close()
}

const (
	_dfdca = "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061tGe\u006e\u0065\u0072\u0061\u006cS\u0074a\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0057\u0068\u006f\u006ce\u004e\u0075\u006d\u0062\u0065\u0072\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0032\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006da\u0074\u0033\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064F\u006f\u0072\u006d\u0061\u0074\u0034"
	_bacb  = "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074P\u0065\u0072\u0063\u0065\u006e\u0074\u0053\u0074\u0061nd\u0061r\u0064F\u006fr\u006d\u0061\u0074\u0031\u0030\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061t\u0031\u0031\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064F\u006f\u0072\u006d\u0061\u0074\u0031\u0032\u0053\u0074a\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0033\u0053t\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0044\u0061\u0074\u0065\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046o\u0072\u006d\u0061\u0074\u00315\u0053\u0074\u0061\u006e\u0064a\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0036\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0037S\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0031\u0038\u0053\u0074\u0061n\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0054\u0069\u006d\u0065\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u00320\u0053\u0074a\u006e\u0064a\u0072\u0064\u0046\u006f\u0072\u006d\u0061t\u0032\u0031\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0044\u0061t\u0065\u0054\u0069\u006d\u0065"
	_ggca  = "\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0033\u0037\u0053t\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006da\u0074\u0033\u0038\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u00339\u0053\u0074\u0061\u006e\u0064\u0061r\u0064\u0046o\u0072\u006da\u00744\u0030"
	_fcdf  = "\u0053t\u0061\u006e\u0064a\u0072\u0064\u0046o\u0072ma\u0074\u0034\u0035\u0053\u0074\u0061\u006ed\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0034\u0036\u0053\u0074\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061\u0074\u0034\u0037\u0053ta\u006ed\u0061\u0072\u0064\u0046\u006f\u0072m\u0061\u0074\u0034\u0038\u0053t\u0061\u006e\u0064\u0061\u0072\u0064\u0046\u006f\u0072\u006d\u0061t\u0034\u0039"
)

// AddConditionalFormatting adds conditional formatting to the sheet.
func (_fedg *Sheet) AddConditionalFormatting(cellRanges []string) ConditionalFormatting {
	_ffdc := _bbg.NewCT_ConditionalFormatting()
	_fedg._ecgc.ConditionalFormatting = append(_fedg._ecgc.ConditionalFormatting, _ffdc)
	_bdf := make(_bbg.ST_Sqref, 0, 0)
	_ffdc.SqrefAttr = &_bdf
	for _, _gged := range cellRanges {
		*_ffdc.SqrefAttr = append(*_ffdc.SqrefAttr, _gged)
	}
	return ConditionalFormatting{_ffdc}
}

// LastColumn returns the name of last column which contains data in range of context sheet's given rows.
func (_bga *evalContext) LastColumn(rowFrom, rowTo int) string {
	_ffgd := _bga._befg
	_fge := 1
	for _aaae := rowFrom; _aaae <= rowTo; _aaae++ {
		_bbda := len(_ffgd.Row(uint32(_aaae)).Cells())
		if _bbda > _fge {
			_fge = _bbda
		}
	}
	return _df.IndexToColumn(uint32(_fge - 1))
}

// SetYSplit sets the row split point
func (_afgf SheetView) SetYSplit(v float64) {
	_afgf.ensurePane()
	_afgf._abfc.Pane.YSplitAttr = _b.Float64(v)
}

// SetTopLeft sets the top left visible cell after the split.
func (_cca SheetView) SetTopLeft(cellRef string) {
	_cca.ensurePane()
	_cca._abfc.Pane.TopLeftCellAttr = &cellRef
}

// LockObject controls the locking of the sheet objects.
func (_bdgc SheetProtection) LockObject(b bool) {
	if !b {
		_bdgc._ggbb.ObjectsAttr = nil
	} else {
		_bdgc._ggbb.ObjectsAttr = _b.Bool(true)
	}
}

// Font allows editing fonts within a spreadsheet stylesheet.
type Font struct {
	_bdbb *_bbg.CT_Font
	_gabf *_bbg.StyleSheet
}

// Tables returns a slice of all defined tables in the workbook.
func (_gfcdf *Workbook) Tables() []Table {
	if _gfcdf._cgdd == nil {
		return nil
	}
	_ebede := []Table{}
	for _, _cfcb := range _gfcdf._cgdd {
		_ebede = append(_ebede, Table{_cfcb})
	}
	return _ebede
}

// X returns the inner wrapped XML type.
func (_dcbf DefinedName) X() *_bbg.CT_DefinedName { return _dcbf._dfced }
func CreateDefaultNumberFormat(id StandardFormat) NumberFormat {
	_gde := NumberFormat{_cbd: _bbg.NewCT_NumFmt()}
	_gde._cbd.NumFmtIdAttr = uint32(id)
	_gde._cbd.FormatCodeAttr = "\u0047e\u006e\u0065\u0072\u0061\u006c"
	switch id {
	case StandardFormat0:
		_gde._cbd.FormatCodeAttr = "\u0047e\u006e\u0065\u0072\u0061\u006c"
	case StandardFormat1:
		_gde._cbd.FormatCodeAttr = "\u0030"
	case StandardFormat2:
		_gde._cbd.FormatCodeAttr = "\u0030\u002e\u0030\u0030"
	case StandardFormat3:
		_gde._cbd.FormatCodeAttr = "\u0023\u002c\u0023#\u0030"
	case StandardFormat4:
		_gde._cbd.FormatCodeAttr = "\u0023\u002c\u0023\u0023\u0030\u002e\u0030\u0030"
	case StandardFormat9:
		_gde._cbd.FormatCodeAttr = "\u0030\u0025"
	case StandardFormat10:
		_gde._cbd.FormatCodeAttr = "\u0030\u002e\u00300\u0025"
	case StandardFormat11:
		_gde._cbd.FormatCodeAttr = "\u0030\u002e\u0030\u0030\u0045\u002b\u0030\u0030"
	case StandardFormat12:
		_gde._cbd.FormatCodeAttr = "\u0023\u0020\u003f/\u003f"
	case StandardFormat13:
		_gde._cbd.FormatCodeAttr = "\u0023 \u003f\u003f\u002f\u003f\u003f"
	case StandardFormat14:
		_gde._cbd.FormatCodeAttr = "\u006d\u002f\u0064\u002f\u0079\u0079"
	case StandardFormat15:
		_gde._cbd.FormatCodeAttr = "\u0064\u002d\u006d\u006d\u006d\u002d\u0079\u0079"
	case StandardFormat16:
		_gde._cbd.FormatCodeAttr = "\u0064\u002d\u006dm\u006d"
	case StandardFormat17:
		_gde._cbd.FormatCodeAttr = "\u006d\u006d\u006d\u002d\u0079\u0079"
	case StandardFormat18:
		_gde._cbd.FormatCodeAttr = "\u0068\u003a\u006d\u006d\u0020\u0041\u004d\u002f\u0050\u004d"
	case StandardFormat19:
		_gde._cbd.FormatCodeAttr = "\u0068\u003a\u006d\u006d\u003a\u0073\u0073\u0020\u0041\u004d\u002f\u0050\u004d"
	case StandardFormat20:
		_gde._cbd.FormatCodeAttr = "\u0068\u003a\u006d\u006d"
	case StandardFormat21:
		_gde._cbd.FormatCodeAttr = "\u0068:\u006d\u006d\u003a\u0073\u0073"
	case StandardFormat22:
		_gde._cbd.FormatCodeAttr = "m\u002f\u0064\u002f\u0079\u0079\u0020\u0068\u003a\u006d\u006d"
	case StandardFormat37:
		_gde._cbd.FormatCodeAttr = "\u0023\u002c\u0023\u0023\u0030\u0020\u003b\u0028\u0023,\u0023\u0023\u0030\u0029"
	case StandardFormat38:
		_gde._cbd.FormatCodeAttr = "\u0023\u002c\u0023\u00230 \u003b\u005b\u0052\u0065\u0064\u005d\u0028\u0023\u002c\u0023\u0023\u0030\u0029"
	case StandardFormat39:
		_gde._cbd.FormatCodeAttr = "\u0023\u002c\u0023\u00230.\u0030\u0030\u003b\u0028\u0023\u002c\u0023\u0023\u0030\u002e\u0030\u0030\u0029"
	case StandardFormat40:
		_gde._cbd.FormatCodeAttr = "\u0023,\u0023\u0023\u0030\u002e\u0030\u0030\u003b\u005b\u0052\u0065\u0064]\u0028\u0023\u002c\u0023\u0023\u0030\u002e\u0030\u0030\u0029"
	case StandardFormat45:
		_gde._cbd.FormatCodeAttr = "\u006d\u006d\u003as\u0073"
	case StandardFormat46:
		_gde._cbd.FormatCodeAttr = "\u005bh\u005d\u003a\u006d\u006d\u003a\u0073s"
	case StandardFormat47:
		_gde._cbd.FormatCodeAttr = "\u006dm\u003a\u0073\u0073\u002e\u0030"
	case StandardFormat48:
		_gde._cbd.FormatCodeAttr = "\u0023\u0023\u0030\u002e\u0030\u0045\u002b\u0030"
	case StandardFormat49:
		_gde._cbd.FormatCodeAttr = "\u0040"
	}
	return _gde
}
func (_cabe Sheet) validateSheetNames() error {
	_decfc := len([]rune(_cabe.Name()))
	if _decfc > 31 {
		return _cg.Errorf("\u0073\u0068\u0065\u0065\u0074 \u006e\u0061\u006d\u0065\u0020\u0027\u0025\u0073\u0027\u0020\u0068\u0061\u0073 \u0025\u0064\u0020\u0063\u0068\u0061\u0072\u0061\u0063\u0074\u0065\u0072\u0073\u002c\u0020\u006d\u0061\u0078\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0069\u0073\u0020\u00331", _cabe.Name(), _decfc)
	}
	return nil
}

// DVCompareOp is a comparison operator for a data validation rule.
type DVCompareOp byte

// SetBorder applies a border to a cell style avoiding redundancy. The function checks if the given border
// already exists in the saved borders. If found, the existing border is reused; otherwise,
// the new border is added to the saved borders collection. The border is then applied to the cell style,
// affecting all styles that reference it by index.
func (_dcce CellStyle) SetBorder(b Border) {
	_eebf := b._de.Border
	for _, _gaa := range _eebf {
		if _ge.DeepEqual(_gaa, b._ef) {
			b._ef = _gaa
			_dcce._gcg.BorderIdAttr = _b.Uint32(b.Index())
			_dcce._gcg.ApplyBorderAttr = _b.Bool(true)
			return
		}
	}
	b._de.Border = append(b._de.Border, b._ef)
	b._de.CountAttr = _b.Uint32(uint32(len(b._de.Border)))
	_dcce._gcg.BorderIdAttr = _b.Uint32(b.Index())
	_dcce._gcg.ApplyBorderAttr = _b.Bool(true)
}

// Extents returns the sheet extents in the form "A1:B15". This requires
// scanning the entire sheet.
func (_ggac Sheet) Extents() string {
	_aedd, _defe, _egag, _gfddb := _ggac.ExtentsIndex()
	return _cg.Sprintf("\u0025s\u0025\u0064\u003a\u0025\u0073\u0025d", _aedd, _defe, _egag, _gfddb)
}
func (_bdbf StyleSheet) GetCellStyle(id uint32) CellStyle {
	for _bcad, _fgaff := range _bdbf._fgbg.CellXfs.Xf {
		if uint32(_bcad) == id {
			return CellStyle{_bdbf._fccg, _fgaff, _bdbf._fgbg.CellXfs}
		}
	}
	return CellStyle{}
}
func (_eead Row) renumberAs(_dddcb uint32) {
	_eead._fddd.RAttr = _b.Uint32(_dddcb)
	for _, _bfc := range _eead.Cells() {
		_bba, _fgdc := _df.ParseCellReference(_bfc.Reference())
		if _fgdc == nil {
			_gac := _cg.Sprintf("\u0025\u0073\u0025\u0064", _bba.Column, _dddcb)
			_bfc._age.RAttr = _b.String(_gac)
		}
	}
}

const (
	DVCompareOpEqual        = DVCompareOp(_bbg.ST_DataValidationOperatorEqual)
	DVCompareOpBetween      = DVCompareOp(_bbg.ST_DataValidationOperatorBetween)
	DVCompareOpNotBetween   = DVCompareOp(_bbg.ST_DataValidationOperatorNotBetween)
	DVCompareOpNotEqual     = DVCompareOp(_bbg.ST_DataValidationOperatorNotEqual)
	DVCompareOpGreater      = DVCompareOp(_bbg.ST_DataValidationOperatorGreaterThan)
	DVCompareOpGreaterEqual = DVCompareOp(_bbg.ST_DataValidationOperatorGreaterThanOrEqual)
	DVCompareOpLess         = DVCompareOp(_bbg.ST_DataValidationOperatorLessThan)
	DVCompareOpLessEqual    = DVCompareOp(_bbg.ST_DataValidationOperatorLessThanOrEqual)
)

// Type returns the type of the rule
func (_begg ConditionalFormattingRule) Type() _bbg.ST_CfType { return _begg._daf.TypeAttr }

// X returns the inner wrapped XML type.
func (_bbge NumberFormat) X() *_bbg.CT_NumFmt { return _bbge._cbd }

// SetWidth controls the width of a column.
func (_fgd Column) SetWidth(w _bg.Distance) {
	_fgd._ggba.WidthAttr = _b.Float64(float64(w / _bg.Character))
}

// X returns the inner XML entity for a stylesheet.
func (_eeacf StyleSheet) X() *_bbg.StyleSheet { return _eeacf._fgbg }

// GetFilename returns the filename of the context's workbook.
func (_fecf *evalContext) GetFilename() string { return _fecf._befg._afga.GetFilename() }

// SetItalic causes the text to be displayed in italic.
func (_faa RichTextRun) SetItalic(b bool) {
	_faa.ensureRpr()
	_faa._ffae.RPr.I = _bbg.NewCT_BooleanProperty()
	_faa._ffae.RPr.I.ValAttr = _b.Bool(b)
}

// ClearFont clears any font configuration from the cell style.
func (_dec CellStyle) ClearFont() { _dec._gcg.FontIdAttr = nil; _dec._gcg.ApplyFontAttr = nil }

const (
	StandardFormatGeneral     StandardFormat = 0
	StandardFormat0           StandardFormat = 0
	StandardFormatWholeNumber StandardFormat = 1
	StandardFormat1           StandardFormat = 1
	StandardFormat2           StandardFormat = 2
	StandardFormat3           StandardFormat = 3
	StandardFormat4           StandardFormat = 4
	StandardFormatPercent     StandardFormat = 9
	StandardFormat9           StandardFormat = 9
	StandardFormat10          StandardFormat = 10
	StandardFormat11          StandardFormat = 11
	StandardFormat12          StandardFormat = 12
	StandardFormat13          StandardFormat = 13
	StandardFormatDate        StandardFormat = 14
	StandardFormat14          StandardFormat = 14
	StandardFormat15          StandardFormat = 15
	StandardFormat16          StandardFormat = 16
	StandardFormat17          StandardFormat = 17
	StandardFormat18          StandardFormat = 18
	StandardFormatTime        StandardFormat = 19
	StandardFormat19          StandardFormat = 19
	StandardFormat20          StandardFormat = 20
	StandardFormat21          StandardFormat = 21
	StandardFormatDateTime    StandardFormat = 22
	StandardFormat22          StandardFormat = 22
	StandardFormat37          StandardFormat = 37
	StandardFormat38          StandardFormat = 38
	StandardFormat39          StandardFormat = 39
	StandardFormat40          StandardFormat = 40
	StandardFormat45          StandardFormat = 45
	StandardFormat46          StandardFormat = 46
	StandardFormat47          StandardFormat = 47
	StandardFormat48          StandardFormat = 48
	StandardFormat49          StandardFormat = 49
)

// PasswordHash returns the hash of the workbook password.
func (_gddb SheetProtection) PasswordHash() string {
	if _gddb._ggbb.PasswordAttr == nil {
		return ""
	}
	return *_gddb._ggbb.PasswordAttr
}

// Workbook is the top level container item for a set of spreadsheets.
type Workbook struct {
	_dcb.DocBase
	_bdff         *_bbg.Workbook
	StyleSheet    StyleSheet
	SharedStrings SharedStrings
	_cffde        []*_bbg.Comments
	_cfgf         []*_bbg.Worksheet
	_fgdcb        []_dcb.Relationships
	_fdae         _dcb.Relationships
	_cbae         []*_acg.Theme
	_eded         []*_ag.WsDr
	_ccag         []_dcb.Relationships
	_fdbb         []*_bb.Container
	_ccbc         []*_fe.ChartSpace
	_cgdd         []*_bbg.Table
	_aafb         string
	_bcae         map[string]string
	_aabe         map[string]*_fe.ChartSpace
	_edac         string
}

func (_gdg Font) SetSize(size float64) { _gdg._bdbb.Sz = []*_bbg.CT_FontSize{{ValAttr: size}} }
func _gaag() *_ag.CT_TwoCellAnchor {
	_aba := _ag.NewCT_TwoCellAnchor()
	_aba.EditAsAttr = _ag.ST_EditAsOneCell
	_aba.From.Col = 5
	_aba.From.Row = 0
	_aba.From.ColOff.ST_CoordinateUnqualified = _b.Int64(0)
	_aba.From.RowOff.ST_CoordinateUnqualified = _b.Int64(0)
	_aba.To.Col = 10
	_aba.To.Row = 20
	_aba.To.ColOff.ST_CoordinateUnqualified = _b.Int64(0)
	_aba.To.RowOff.ST_CoordinateUnqualified = _b.Int64(0)
	return _aba
}

type SheetProtection struct{ _ggbb *_bbg.CT_SheetProtection }

// AddRule adds and returns a new rule that can be configured.
func (_ffge ConditionalFormatting) AddRule() ConditionalFormattingRule {
	_aef := _bbg.NewCT_CfRule()
	_ffge._cff.CfRule = append(_ffge._cff.CfRule, _aef)
	_fgafa := ConditionalFormattingRule{_aef}
	_fgafa.InitializeDefaults()
	_fgafa.SetPriority(int32(len(_ffge._cff.CfRule) + 1))
	return _fgafa
}

// SetNumberFormat applies a number format to a cell style avoiding redundancy. The function checks if the given string
// already exists in the saved number formats. If found, the existing number format is reused; otherwise,
// the new number format is added to the saved number formats collection. The number format is then applied to the cell style,
// affecting all styles that reference it by index.
func (_dgc CellStyle) SetNumberFormat(s string) {
	var _ced NumberFormat
	if _dgc._ffgf.StyleSheet._fgbg.NumFmts == nil {
		_dgc._ffgf.StyleSheet._fgbg.NumFmts = _bbg.NewCT_NumFmts()
	}
	_bbdc := _dgc._ffgf.StyleSheet._fgbg.NumFmts.NumFmt
	for _, _gcbc := range _bbdc {
		if _ge.DeepEqual(_gcbc.FormatCodeAttr, s) {
			_ced = NumberFormat{_dgc._ffgf, _gcbc}
			_dgc._gcg.ApplyNumberFormatAttr = _b.Bool(true)
			_dgc._gcg.NumFmtIdAttr = _b.Uint32(_ced.ID())
			return
		}
	}
	_ega := _bbg.NewCT_NumFmt()
	_ega.NumFmtIdAttr = uint32(200 + len(_dgc._ffgf.StyleSheet._fgbg.NumFmts.NumFmt))
	_dgc._ffgf.StyleSheet._fgbg.NumFmts.NumFmt = append(_dgc._ffgf.StyleSheet._fgbg.NumFmts.NumFmt, _ega)
	_dgc._ffgf.StyleSheet._fgbg.NumFmts.CountAttr = _b.Uint32(uint32(len(_dgc._ffgf.StyleSheet._fgbg.NumFmts.NumFmt)))
	_ced = NumberFormat{_dgc._ffgf, _ega}
	_ced._cbd.FormatCodeAttr = s
	_dgc._gcg.ApplyNumberFormatAttr = _b.Bool(true)
	_dgc._gcg.NumFmtIdAttr = _b.Uint32(_ced.ID())
}

// Cell retrieves or adds a new cell to a row. Col is the column (e.g. 'A', 'B')
func (_fdac Row) Cell(col string) Cell {
	_dfca := _cg.Sprintf("\u0025\u0073\u0025\u0064", col, _fdac.RowNumber())
	for _, _gdef := range _fdac._fddd.C {
		if _gdef.RAttr != nil && *_gdef.RAttr == _dfca {
			return Cell{_fdac._fadd, _fdac._bafc, _fdac._fddd, _gdef}
		}
	}
	return _fdac.AddNamedCell(col)
}

// SetActiveSheet sets the active sheet which will be the tab displayed when the
// spreadsheet is initially opened.
func (_gbfff *Workbook) SetActiveSheet(s Sheet) {
	for _dacg, _ffaa := range _gbfff._cfgf {
		if s._ecgc == _ffaa {
			_gbfff.SetActiveSheetIndex(uint32(_dacg))
		}
	}
}

// GetVerticalAlignment sets the vertical alignment of a cell style.
func (_egc CellStyle) GetVerticalAlignment() _bbg.ST_VerticalAlignment {
	if _egc._gcg.Alignment == nil {
		return _bbg.ST_VerticalAlignmentUnset
	}
	return _egc._gcg.Alignment.VerticalAttr
}

// GetFormat returns a cell data format.
func (_fbb *evalContext) GetFormat(cellRef string) string {
	return _fbb._befg.Cell(cellRef).getFormat()
}

// SetRowOffset sets a column offset in absolute distance.
func (_cdca CellMarker) SetRowOffset(m _bg.Distance) {
	_cdca._bef.RowOff.ST_CoordinateUnqualified = _b.Int64(int64(m / _bg.EMU))
}

// MoveTo is a no-op.
func (_dfd AbsoluteAnchor) MoveTo(x, y int32) {}

// X returns the inner wrapped XML type.
func (_bfdf SheetProtection) X() *_bbg.CT_SheetProtection { return _bfdf._ggbb }

// CopySheetByName copies the existing sheet with the name `name` and puts its copy with the name `copiedSheetName`.
func (_ebcc *Workbook) CopySheetByName(name, copiedSheetName string) (Sheet, error) {
	_ffdg := -1
	for _bfce, _gdcb := range _ebcc.Sheets() {
		if name == _gdcb.Name() {
			_ffdg = _bfce
			break
		}
	}
	if _ffdg == -1 {
		return Sheet{}, ErrorNotFound
	}
	return _ebcc.CopySheet(_ffdg, copiedSheetName)
}

// SetStyle applies a style to a cell avoiding redundancy. The function checks if the given style
// already exists in the saved styles. If found, the existing style is reused; otherwise,
// the new style is added to the saved styles collection. The style is then applied to the cell.
// This style is referenced in the generated XML via CellStyle.Index().
func (_dfee Cell) SetStyle(cs CellStyle) {
	_agfc := cs._efb.Xf
	for _, _dcfe := range _agfc {
		if _ge.DeepEqual(_dcfe, cs._gcg) {
			cs._gcg = _dcfe
			_dfee.SetStyleIndex(cs.Index())
			return
		}
	}
	cs._efb.Xf = append(cs._efb.Xf, cs._gcg)
	cs._efb.CountAttr = _b.Uint32(uint32(len(cs._efb.Xf)))
	_dfee.SetStyleIndex(cs.Index())
}

// SetPasswordHash sets the password hash to the input.
func (_dedg WorkbookProtection) SetPasswordHash(pwHash string) {
	_dedg._fefe.WorkbookPasswordAttr = _b.String(pwHash)
}
func (_dfdaag *Sheet) removeColumnFromNamedRanges(_agg uint32) error {
	for _, _abdb := range _dfdaag._afga.DefinedNames() {
		_gcfa := _abdb.Name()
		_dgbe := _abdb.Content()
		_egeb := _ba.Split(_dgbe, "\u0021")
		if len(_egeb) != 2 {
			return _fb.New("\u0049\u006e\u0063\u006frr\u0065\u0063\u0074\u0020\u006e\u0061\u006d\u0065\u0064\u0020\u0072\u0061\u006e\u0067e\u003a" + _dgbe)
		}
		_dbgf := _egeb[0]
		if _dfdaag.Name() == _dbgf {
			_cedcd := _dfdaag._afga.RemoveDefinedName(_abdb)
			if _cedcd != nil {
				return _cedcd
			}
			_dcff := _ffaf(_egeb[1], _agg, true)
			if _dcff != "" {
				_cabb := _dbgf + "\u0021" + _dcff
				_dfdaag._afga.AddDefinedName(_gcfa, _cabb)
			}
		}
	}
	_caeg := 0
	if _dfdaag._ecgc.TableParts != nil && _dfdaag._ecgc.TableParts.TablePart != nil {
		_caeg = len(_dfdaag._ecgc.TableParts.TablePart)
	}
	if _caeg != 0 {
		_agga := 0
		for _, _ecaa := range _dfdaag._afga.Sheets() {
			if _ecaa.Name() == _dfdaag.Name() {
				break
			} else {
				if _ecaa._ecgc.TableParts != nil && _ecaa._ecgc.TableParts.TablePart != nil {
					_agga += len(_ecaa._ecgc.TableParts.TablePart)
				}
			}
		}
		_eeca := _dfdaag._afga._cgdd[_agga : _agga+_caeg]
		for _fab, _gedf := range _eeca {
			_bcfb := _gedf
			_bcfb.RefAttr = _ffaf(_bcfb.RefAttr, _agg, false)
			_dfdaag._afga._cgdd[_agga+_fab] = _bcfb
		}
	}
	return nil
}

// ColorScale colors a cell background based off of the cell value.
type ColorScale struct{ _gfa *_bbg.CT_ColorScale }

// Priority returns the rule priority
func (_fgdb ConditionalFormattingRule) Priority() int32 { return _fgdb._daf.PriorityAttr }

// X returns the inner wrapped XML type.
func (_bddg ColorScale) X() *_bbg.CT_ColorScale { return _bddg._gfa }

// BottomRight is a no-op.
func (_cdf OneCellAnchor) BottomRight() CellMarker { return CellMarker{} }

// SetStyle sets the style to be used for conditional rules
func (_dgca ConditionalFormattingRule) SetStyle(d DifferentialStyle) {
	_dgca._daf.DxfIdAttr = _b.Uint32(d.Index())
}

// GetCachedFormulaResult returns the cached formula result if it exists. If the
// cell type is not a formula cell, the result will be the cell value if it's a
// string/number/bool cell.
func (_fafc Cell) GetCachedFormulaResult() string {
	if _fafc._age.V != nil {
		return *_fafc._age.V
	}
	return ""
}

// GetSheet returns a sheet by name, or an error if a sheet by the given name
// was not found.
func (_dbde *Workbook) GetSheet(name string) (Sheet, error) {
	for _, _bcaeg := range _dbde.Sheets() {
		if _bcaeg.Name() == name {
			return _bcaeg, nil
		}
	}
	return Sheet{}, ErrorNotFound
}
func (_bace *Sheet) removeColumnFromMergedCells(_fffcd uint32) error {
	if _bace._ecgc.MergeCells == nil || _bace._ecgc.MergeCells.MergeCell == nil {
		return nil
	}
	_beegf := []*_bbg.CT_MergeCell{}
	for _, _cbbb := range _bace.MergedCells() {
		_gbdc := _ffaf(_cbbb.Reference(), _fffcd, true)
		if _gbdc != "" {
			_cbbb.SetReference(_gbdc)
			_beegf = append(_beegf, _cbbb.X())
		}
	}
	_bace._ecgc.MergeCells.MergeCell = _beegf
	return nil
}

// LessRows compares two rows based off of a column. If the column doesn't exist
// in one row, that row is 'less'.
func (_ccc Comparer) LessRows(column string, lhs, rhs Row) bool {
	var _cgfg, _bfd Cell
	for _, _gcf := range lhs.Cells() {
		_eff, _ := _df.ParseCellReference(_gcf.Reference())
		if _eff.Column == column {
			_cgfg = _gcf
			break
		}
	}
	for _, _bfg := range rhs.Cells() {
		_dab, _ := _df.ParseCellReference(_bfg.Reference())
		if _dab.Column == column {
			_bfd = _bfg
			break
		}
	}
	return _ccc.LessCells(_cgfg, _bfd)
}

// AnchorType is the type of anchor.
type AnchorType byte

// AddDifferentialStyle adds a new empty differential cell style to the stylesheet.
func (_fbgg StyleSheet) AddDifferentialStyle() DifferentialStyle {
	if _fbgg._fgbg.Dxfs == nil {
		_fbgg._fgbg.Dxfs = _bbg.NewCT_Dxfs()
	}
	_dgdc := _bbg.NewCT_Dxf()
	_fbgg._fgbg.Dxfs.Dxf = append(_fbgg._fgbg.Dxfs.Dxf, _dgdc)
	_fbgg._fgbg.Dxfs.CountAttr = _b.Uint32(uint32(len(_fbgg._fgbg.Dxfs.Dxf)))
	return DifferentialStyle{_dgdc, _fbgg._fccg, _fbgg._fgbg.Dxfs}
}

// SortOrder is a column sort order.
//
//go:generate stringer -type=SortOrder
type SortOrder byte

const (
	DVCompareTypeWholeNumber = DVCompareType(_bbg.ST_DataValidationTypeWhole)
	DVCompareTypeDecimal     = DVCompareType(_bbg.ST_DataValidationTypeDecimal)
	DVCompareTypeDate        = DVCompareType(_bbg.ST_DataValidationTypeDate)
	DVCompareTypeTime        = DVCompareType(_bbg.ST_DataValidationTypeTime)
	DVompareTypeTextLength   = DVCompareType(_bbg.ST_DataValidationTypeTextLength)
)

// IsNumber returns true if the cell is a number type cell.
func (_fgg Cell) IsNumber() bool {
	switch _fgg._age.TAttr {
	case _bbg.ST_CellTypeN:
		return true
	case _bbg.ST_CellTypeS, _bbg.ST_CellTypeB:
		return false
	}
	return _fgg._age.V != nil && _ce.IsNumber(*_fgg._age.V)
}
func (_fffa DataValidation) SetComparison(t DVCompareType, op DVCompareOp) DataValidationCompare {
	_fffa.clear()
	_fffa._bge.TypeAttr = _bbg.ST_DataValidationType(t)
	_fffa._bge.OperatorAttr = _bbg.ST_DataValidationOperator(op)
	return DataValidationCompare{_fffa._bge}
}

// GetFormula returns the formula for a cell.
func (_gdd Cell) GetFormula() string {
	if _gdd._age.F != nil {
		return _gdd._age.F.Content
	}
	return ""
}
func _fca(_bgbd bool) int {
	if _bgbd {
		return 1
	}
	return 0
}

var _eac = [...]uint8{0, 18, 37}

// RemoveCalcChain removes the cached caculation chain. This is sometimes needed
// as we don't update it when rows are added/removed.
func (_gdfb *Workbook) RemoveCalcChain() {
	var _cgag string
	for _, _bbad := range _gdfb._fdae.Relationships() {
		if _bbad.Type() == "ht\u0074\u0070\u003a\u002f\u002f\u0073\u0063he\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006et\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068i\u0070s\u002f\u0063\u0061\u006c\u0063\u0043\u0068\u0061\u0069\u006e" {
			_cgag = "\u0078\u006c\u002f" + _bbad.Target()
			_gdfb._fdae.Remove(_bbad)
			break
		}
	}
	if _cgag == "" {
		return
	}
	_gdfb.ContentTypes.RemoveOverride(_cgag)
	for _gfcb, _cfbee := range _gdfb.ExtraFiles {
		if _cfbee.ZipPath == _cgag {
			_gdfb.ExtraFiles[_gfcb] = _gdfb.ExtraFiles[len(_gdfb.ExtraFiles)-1]
			_gdfb.ExtraFiles = _gdfb.ExtraFiles[:len(_gdfb.ExtraFiles)-1]
			return
		}
	}
}

// DataValidationCompare is a view on a data validation rule that is oriented
// towards value comparisons.
type DataValidationCompare struct{ _fad *_bbg.CT_DataValidation }

// ClearFill clears any fill configuration from the cell style.
func (_eag CellStyle) ClearFill() { _eag._gcg.FillIdAttr = nil; _eag._gcg.ApplyFillAttr = nil }

const _cd = "\u00320\u0030\u0036\u002d\u00301\u002d\u0030\u0032\u0054\u00315\u003a0\u0034:\u0030\u0035\u005a\u0030\u0037\u003a\u00300"

// GetFont gets a Font from a cell style.
func (_gbc CellStyle) GetFont() *_bbg.CT_Font {
	if _fggc := _gbc._gcg.FontIdAttr; _fggc != nil {
		_ebab := _gbc._ffgf.StyleSheet.Fonts()
		if int(*_fggc) < len(_ebab) {
			return _ebab[int(*_fggc)].X()
		}
	}
	return nil
}

// SheetViews returns the sheet views defined.  This is where splits and frozen
// rows/cols are configured.  Multiple sheet views are allowed, but I'm not
// aware of there being a use for more than a single sheet view.
func (_eec *Sheet) SheetViews() []SheetView {
	if _eec._ecgc.SheetViews == nil {
		return nil
	}
	_ffbbd := []SheetView{}
	for _, _geeg := range _eec._ecgc.SheetViews.SheetView {
		_ffbbd = append(_ffbbd, SheetView{_geeg})
	}
	return _ffbbd
}

// SetProtectedAndHidden sets protected and hidden for given cellStyle
func (_dfce CellStyle) SetProtection(protected bool, hidden bool) {
	_dfce._gcg.Protection = &_bbg.CT_CellProtection{LockedAttr: &protected, HiddenAttr: &hidden}
}

// Operator returns the operator for the rule
func (_cgb ConditionalFormattingRule) Operator() _bbg.ST_ConditionalFormattingOperator {
	return _cgb._daf.OperatorAttr
}

// SetColorScale configures the rule as a color scale, removing existing
// configuration.
func (_fdcf ConditionalFormattingRule) SetColorScale() ColorScale {
	_fdcf.clear()
	_fdcf.SetType(_bbg.ST_CfTypeColorScale)
	_fdcf._daf.ColorScale = _bbg.NewCT_ColorScale()
	return ColorScale{_fdcf._daf.ColorScale}
}

// SetVerticalAlignment sets the vertical alignment of a cell style.
func (_egf CellStyle) SetVerticalAlignment(a _bbg.ST_VerticalAlignment) {
	if _egf._gcg.Alignment == nil {
		_egf._gcg.Alignment = _bbg.NewCT_CellAlignment()
	}
	_egf._gcg.ApplyAlignmentAttr = _b.Bool(true)
	_egf._gcg.Alignment.VerticalAttr = a
}

// SetFill applies a fill to a cell style avoiding redundancy. The function checks if the given fill
// already exists in the saved fills. If found, the existing fill is reused; otherwise,
// the new fill is added to the saved fills collection. The fill is then applied to the cell style,
// affecting all styles that reference it by index.
func (_dccd CellStyle) SetFill(f Fill) {
	_gbe := f._fgc.Fill
	for _, _gge := range _gbe {
		if _ge.DeepEqual(_gge, f._aeag) {
			f._aeag = _gge
			_dccd._gcg.FillIdAttr = _b.Uint32(f.Index())
			_dccd._gcg.ApplyFillAttr = _b.Bool(true)
			return
		}
	}
	f._fgc.Fill = append(f._fgc.Fill, f._aeag)
	f._fgc.CountAttr = _b.Uint32(uint32(len(f._fgc.Fill)))
	_dccd._gcg.FillIdAttr = _b.Uint32(f.Index())
	_dccd._gcg.ApplyFillAttr = _b.Bool(true)
}

// GetFormat sets the number format code.
func (_cfgc NumberFormat) GetFormat() string { return _cfgc._cbd.FormatCodeAttr }

// X returns the inner wrapped XML type.
func (_fdcac DataBarScale) X() *_bbg.CT_DataBar { return _fdcac._gdb }

// SetLocked sets cell locked or not.
func (_cgd *evalContext) SetLocked(cellRef string, locked bool) {
	_cgd._befg.Cell(cellRef).setLocked(locked)
}

// Column returns or creates a column that with a given index (1-N).  Columns
// can span multiple column indices, this method will return the column that
// applies to a column index if it exists or create a new column that only
// applies to the index passed in otherwise.
func (_dddcf *Sheet) Column(idx uint32) Column {
	for _, _eedb := range _dddcf._ecgc.Cols {
		for _, _bfaf := range _eedb.Col {
			if idx >= _bfaf.MinAttr && idx <= _bfaf.MaxAttr {
				return Column{_bfaf}
			}
		}
	}
	var _efcaf *_bbg.CT_Cols
	if len(_dddcf._ecgc.Cols) == 0 {
		_efcaf = _bbg.NewCT_Cols()
		_dddcf._ecgc.Cols = append(_dddcf._ecgc.Cols, _efcaf)
	} else {
		_efcaf = _dddcf._ecgc.Cols[0]
	}
	_bcfa := _bbg.NewCT_Col()
	_bcfa.MinAttr = idx
	_bcfa.MaxAttr = idx
	_efcaf.Col = append(_efcaf.Col, _bcfa)
	return Column{_bcfa}
}
func (_gecc *Sheet) addNumberedRowFast(_cafab uint32) Row {
	_gcec := _bbg.NewCT_Row()
	_gcec.RAttr = _b.Uint32(_cafab)
	_gecc._ecgc.SheetData.Row = append(_gecc._ecgc.SheetData.Row, _gcec)
	return Row{_gecc._afga, _gecc, _gcec}
}

type Fills struct{ _egaa *_bbg.CT_Fills }

func (_eed CellStyle) SetShrinkToFit(b bool) {
	if _eed._gcg.Alignment == nil {
		_eed._gcg.Alignment = _bbg.NewCT_CellAlignment()
	}
	_eed._gcg.ApplyAlignmentAttr = _b.Bool(true)
	if !b {
		_eed._gcg.Alignment.ShrinkToFitAttr = nil
	} else {
		_eed._gcg.Alignment.ShrinkToFitAttr = _b.Bool(b)
	}
}

// TopLeft is a no-op.
func (_acc AbsoluteAnchor) TopLeft() CellMarker { return CellMarker{} }

// SetColOffset sets the column offset of the top-left anchor.
func (_ebaa OneCellAnchor) SetColOffset(m _bg.Distance) { _ebaa.TopLeft().SetColOffset(m) }

// LastRow returns the name of last row which contains data in range of context sheet's given columns.
func (_gbga *evalContext) LastRow(col string) int {
	_dfdd := _gbga._befg
	_fcdg := int(_df.ColumnToIndex(col))
	_becf := 1
	for _, _dadb := range _dfdd._ecgc.SheetData.Row {
		if _dadb.RAttr != nil {
			_dfec := Row{_dfdd._afga, _dfdd, _dadb}
			_fee := len(_dfec.Cells())
			if _fee > _fcdg {
				_becf = int(_dfec.RowNumber())
			}
		}
	}
	return _becf
}

// SetRange sets the cell or range of cells that the validation should apply to.
// It can be a single cell (e.g. "A1") or a range of cells (e.g. "A1:B5")
func (_dfba DataValidation) SetRange(cellRange string) {
	_dfba._bge.SqrefAttr = _bbg.ST_Sqref{cellRange}
}

// TopLeft returns the CellMaker for the top left corner of the anchor.
func (_dgde TwoCellAnchor) TopLeft() CellMarker { return CellMarker{_dgde._ccec.From} }

// SetNumberFormatStandard sets the format based off of the ECMA 376 standard formats.  These
// formats are standardized and don't need to be defined in the styles.
func (_cbaf CellStyle) SetNumberFormatStandard(s StandardFormat) {
	_cbaf._gcg.NumFmtIdAttr = _b.Uint32(uint32(s))
	_cbaf._gcg.ApplyNumberFormatAttr = _b.Bool(true)
}

// New constructs a new workbook.
func New() *Workbook {
	_gbbg := &Workbook{}
	_gbbg._bdff = _bbg.NewWorkbook()
	_gbbg.AppProperties = _dcb.NewAppProperties()
	_gbbg.CoreProperties = _dcb.NewCoreProperties()
	_gbbg.StyleSheet = NewStyleSheet(_gbbg)
	_gbbg.Rels = _dcb.NewRelationships()
	_gbbg._fdae = _dcb.NewRelationships()
	_gbbg.Rels.AddRelationship(_b.RelativeFilename(_b.DocTypeSpreadsheet, "", _b.ExtendedPropertiesType, 0), _b.ExtendedPropertiesType)
	_gbbg.Rels.AddRelationship(_b.RelativeFilename(_b.DocTypeSpreadsheet, "", _b.CorePropertiesType, 0), _b.CorePropertiesType)
	_gbbg.Rels.AddRelationship(_b.RelativeFilename(_b.DocTypeSpreadsheet, "", _b.OfficeDocumentType, 0), _b.OfficeDocumentType)
	_gbbg._fdae.AddRelationship(_b.RelativeFilename(_b.DocTypeSpreadsheet, _b.OfficeDocumentType, _b.StylesType, 0), _b.StylesType)
	_gbbg.ContentTypes = _dcb.NewContentTypes()
	_gbbg.ContentTypes.AddDefault("\u0076\u006d\u006c", _b.VMLDrawingContentType)
	_gbbg.ContentTypes.AddOverride(_b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.OfficeDocumentType, 0), "\u0061\u0070\u0070\u006c\u0069c\u0061\u0074\u0069\u006f\u006e\u002fv\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066o\u0072\u006d\u0061\u0074s\u002d\u006f\u0066\u0066\u0069\u0063e\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068e\u0065\u0074\u006d\u006c\u002e\u0073\u0068\u0065\u0065\u0074\u002e\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
	_gbbg.ContentTypes.AddOverride(_b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.StylesType, 0), _b.SMLStyleSheetContentType)
	_gbbg.SharedStrings = NewSharedStrings()
	_gbbg.ContentTypes.AddOverride(_b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.SharedStringsType, 0), _b.SharedStringsContentType)
	_gbbg._fdae.AddRelationship(_b.RelativeFilename(_b.DocTypeSpreadsheet, _b.OfficeDocumentType, _b.SharedStringsType, 0), _b.SharedStringsType)
	_gbbg._bcae = map[string]string{}
	return _gbbg
}

// SetColor sets the text color.
func (_egaf RichTextRun) SetColor(c _fbf.Color) {
	_egaf.ensureRpr()
	_egaf._ffae.RPr.Color = _bbg.NewCT_Color()
	_gfee := "\u0066\u0066" + *c.AsRGBString()
	_egaf._ffae.RPr.Color.RgbAttr = &_gfee
}

// SetPassword sets the password hash to a hash of the input password.
func (_dceb WorkbookProtection) SetPassword(pw string) { _dceb.SetPasswordHash(PasswordHash(pw)) }

// WorkbookText is an array of extracted text items which has some methods for representing extracted text from a workbook.
type WorkbookText struct{ Sheets []*SheetText }

func (_gf Border) SetLeft(style _bbg.ST_BorderStyle, c _fbf.Color) {
	if _gf._ef.Left == nil {
		_gf._ef.Left = _bbg.NewCT_BorderPr()
	}
	_gf._ef.Left.Color = _bbg.NewCT_Color()
	_gf._ef.Left.Color.RgbAttr = c.AsRGBAString()
	_gf._ef.Left.StyleAttr = style
}

// SetHeightCells is a no-op.
func (_cf AbsoluteAnchor) SetHeightCells(int32) {}

// Read reads a workbook from an io.Reader(.xlsx).
func Read(r _ee.ReaderAt, size int64) (*Workbook, error) {
	const _gbbgc = "\u0073\u0070r\u0065\u0061\u0064s\u0068\u0065\u0065\u0074\u003a\u0052\u0065\u0061\u0064"
	// if !_ac.GetLicenseKey().IsLicensed() && !_cggde {
	// 	_cg.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
	// 	_cg.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
	// 	return nil, _fb.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	// }
	_gdea := New()
	// _agdd, _fbd := _ac.GenRefId("\u0073\u0072")
	// if _fbd != nil {
	// 	_ca.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _fbd)
	// 	return nil, _fbd
	// }
	// _gdea._edac = _agdd
	// if _gbdg := _ac.Track(_gdea._edac, _gbbgc); _gbdg != nil {
	// 	_ca.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gbdg)
	// 	return nil, _gbdg
	// }
	_bafd, _fbd := _ae.TempDir("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065-\u0078\u006c\u0073\u0078")
	if _fbd != nil {
		return nil, _fbd
	}
	_gdea.TmpPath = _bafd
	_dfgd, _fbd := _bdg.NewReader(r, size)
	if _fbd != nil {
		return nil, _cg.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u007a\u0069\u0070\u003a\u0020\u0025\u0073", _fbd)
	}
	_dgfg := []*_bdg.File{}
	_dgfg = append(_dgfg, _dfgd.File...)
	_dgec := false
	for _, _bag := range _dgfg {
		if _bag.FileHeader.Name == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
			_dgec = true
			break
		}
	}
	if _dgec {
		_gdea.CreateCustomProperties()
	}
	_aaab := _dd.DecodeMap{}
	_aaab.SetOnNewRelationshipFunc(_gdea.onNewRelationship)
	_aaab.AddTarget(_b.ContentTypesFilename, _gdea.ContentTypes.X(), "", 0)
	_aaab.AddTarget(_b.BaseRelsFilename, _gdea.Rels.X(), "", 0)
	if _bffb := _aaab.Decode(_dgfg); _bffb != nil {
		return nil, _bffb
	}
	for _, _eab := range _dgfg {
		if _eab == nil {
			continue
		}
		if _effb := _gdea.AddExtraFileFromZip(_eab); _effb != nil {
			return nil, _effb
		}
	}
	if _dgec {
		_gadc := false
		for _, _ged := range _gdea.Rels.X().Relationship {
			if _ged.TargetAttr == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
				_gadc = true
				break
			}
		}
		if !_gadc {
			_gdea.AddCustomRelationships()
		}
	}
	return _gdea, nil
}

// MoveTo moves the top-left of the anchored object.
func (_aff OneCellAnchor) MoveTo(col, row int32) {
	_aff.TopLeft().SetCol(col)
	_aff.TopLeft().SetRow(row)
}

// MaxColumnIdx returns the max used column of the sheet.
func (_agae Sheet) MaxColumnIdx() uint32 {
	_ceeg := uint32(0)
	for _, _cfad := range _agae.Rows() {
		_gea := _cfad._fddd.C
		if len(_gea) > 0 {
			_agege := _gea[len(_gea)-1]
			_cded, _ := _df.ParseCellReference(*_agege.RAttr)
			if _ceeg < _cded.ColumnIdx {
				_ceeg = _cded.ColumnIdx
			}
		}
	}
	return _ceeg
}

// Fonts returns the list of fonts defined in the stylesheet.
func (_ddeca StyleSheet) Fonts() []Font {
	_eda := []Font{}
	for _, _bdab := range _ddeca._fgbg.Fonts.Font {
		_eda = append(_eda, Font{_bdab, _ddeca._fgbg})
	}
	return _eda
}

// SetRange sets the range that contains the possible values. This is incompatible with SetValues.
func (_bfbe DataValidationList) SetRange(cellRange string) {
	_bfbe._cfg.Formula1 = _b.String(cellRange)
	_bfbe._cfg.Formula2 = _b.String("\u0030")
}
func (_gbcdf Fill) SetPatternFill() PatternFill {
	_gbcdf._aeag.GradientFill = nil
	_gbcdf._aeag.PatternFill = _bbg.NewCT_PatternFill()
	_gbcdf._aeag.PatternFill.PatternTypeAttr = _bbg.ST_PatternTypeSolid
	return PatternFill{_gbcdf._aeag.PatternFill, _gbcdf._aeag}
}

// AddHyperlink adds a hyperlink to a sheet. Adding the hyperlink to the sheet
// and setting it on a cell is more efficient than setting hyperlinks directly
// on a cell.
func (_ggcb *Sheet) AddHyperlink(url string) _dcb.Hyperlink {
	for _afbc, _gdgg := range _ggcb._afga._cfgf {
		if _gdgg == _ggcb._ecgc {
			return _ggcb._afga._fgdcb[_afbc].AddHyperlink(url)
		}
	}
	return _dcb.Hyperlink{}
}

type PatternFill struct {
	_aad  *_bbg.CT_PatternFill
	_acgd *_bbg.CT_Fill
}

// Name returns the name of the table
func (_cega Table) Name() string {
	if _cega._febfb.NameAttr != nil {
		return *_cega._febfb.NameAttr
	}
	return ""
}

// X returns the inner wrapped XML type.
func (_gdc Row) X() *_bbg.CT_Row { return _gdc._fddd }
func (_bdd Cell) getRawSortValue() (string, bool) {
	if _bdd.HasFormula() {
		_dda := _bdd.GetCachedFormulaResult()
		return _dda, _ce.IsNumber(_dda)
	}
	_aaa, _ := _bdd.GetRawValue()
	return _aaa, _ce.IsNumber(_aaa)
}

// SetAutoFilter creates autofilters on the sheet. These are the automatic
// filters that are common for a header row.  The RangeRef should be of the form
// "A1:C5" and cover the entire range of cells to be filtered, not just the
// header. SetAutoFilter replaces any existing auto filter on the sheet.
func (_dgfb *Sheet) SetAutoFilter(rangeRef string) {
	rangeRef = _ba.Replace(rangeRef, "\u0024", "", -1)
	_dgfb._ecgc.AutoFilter = _bbg.NewCT_AutoFilter()
	_dgfb._ecgc.AutoFilter.RefAttr = _b.String(rangeRef)
	_bdcg := "\u0027" + _dgfb.Name() + "\u0027\u0021"
	var _gccg DefinedName
	for _, _edg := range _dgfb._afga.DefinedNames() {
		if _edg.Name() == _afcb {
			if _ba.HasPrefix(_edg.Content(), _bdcg) {
				_gccg = _edg
				_gccg.SetContent(_dgfb.RangeReference(rangeRef))
				break
			}
		}
	}
	if _gccg.X() == nil {
		_gccg = _dgfb._afga.AddDefinedName(_afcb, _dgfb.RangeReference(rangeRef))
	}
	for _dged, _bgfg := range _dgfb._afga._cfgf {
		if _bgfg == _dgfb._ecgc {
			_gccg.SetLocalSheetID(uint32(_dged))
		}
	}
}

// SetCachedFormulaResult sets the cached result of a formula. This is normally
// not needed but is used internally when expanding an array formula.
func (_cda Cell) SetCachedFormulaResult(s string) { _cda._age.V = &s }

const (
	DVOpGreater = _bbg.ST_DataValidationOperatorGreaterThanOrEqual
)

// SetColOffset sets the column offset of the top-left of the image in fixed units.
func (_gc AbsoluteAnchor) SetColOffset(m _bg.Distance) {
	_gc._da.Pos.XAttr.ST_CoordinateUnqualified = _b.Int64(int64(m / _bg.EMU))
}
func (_ffgec Sheet) ExtentsIndex() (string, uint32, string, uint32) {
	var _gage, _caef, _afd, _eebe uint32 = 1, 1, 0, 0
	for _, _gba := range _ffgec.Rows() {
		if _gba.RowNumber() < _gage {
			_gage = _gba.RowNumber()
		} else if _gba.RowNumber() > _caef {
			_caef = _gba.RowNumber()
		}
		for _, _babde := range _gba.Cells() {
			_cgccd, _bcgg := _df.ParseCellReference(_babde.Reference())
			if _bcgg == nil {
				if _cgccd.ColumnIdx < _afd {
					_afd = _cgccd.ColumnIdx
				} else if _cgccd.ColumnIdx > _eebe {
					_eebe = _cgccd.ColumnIdx
				}
			}
		}
	}
	return _df.IndexToColumn(_afd), _gage, _df.IndexToColumn(_eebe), _caef
}

// DVCompareType is a comparison type for a data validation rule. This restricts
// the input format of the cell.
type DVCompareType byte

func (_baeb ConditionalFormattingRule) InitializeDefaults() {
	_baeb.SetType(_bbg.ST_CfTypeCellIs)
	_baeb.SetOperator(_bbg.ST_ConditionalFormattingOperatorGreaterThan)
	_baeb.SetPriority(1)
}
func (_ecda *Sheet) setArray(_gedb string, _dddd _eg.Result) error {
	_cefc, _dfage := _df.ParseCellReference(_gedb)
	if _dfage != nil {
		return _dfage
	}
	for _dddfc, _agebf := range _dddd.ValueArray {
		_aaff := _ecda.Row(_cefc.RowIdx + uint32(_dddfc))
		for _dggg, _bgaf := range _agebf {
			_cad := _aaff.Cell(_df.IndexToColumn(_cefc.ColumnIdx + uint32(_dggg)))
			if _bgaf.Type != _eg.ResultTypeEmpty {
				if _bgaf.IsBoolean {
					_cad.SetBool(_bgaf.ValueNumber != 0)
				} else {
					_cad.SetCachedFormulaResult(_bgaf.String())
				}
			}
		}
	}
	return nil
}

// IsHidden returns whether the row is hidden or not.
func (_gbef Row) IsHidden() bool { return _gbef._fddd.HiddenAttr != nil && *_gbef._fddd.HiddenAttr }

// ClearCachedFormulaResults clears any computed formula values that are stored
// in the sheet. This may be required if you modify cells that are used as a
// formula input to force the formulas to be recomputed the next time the sheet
// is opened in Excel.
func (_cfbg *Workbook) ClearCachedFormulaResults() {
	for _, _cfbc := range _cfbg.Sheets() {
		_cfbc.ClearCachedFormulaResults()
	}
}

// Row is a row within a spreadsheet.
type Row struct {
	_fadd *Workbook
	_bafc *Sheet
	_fddd *_bbg.CT_Row
}

func (_gabg SortOrder) String() string {
	if _gabg >= SortOrder(len(_eac)-1) {
		return _cg.Sprintf("\u0053\u006f\u0072\u0074\u004f\u0072\u0064\u0065\u0072\u0028\u0025\u0064\u0029", _gabg)
	}
	return _feeb[_eac[_gabg]:_eac[_gabg+1]]
}
func (_fd Border) SetBottom(style _bbg.ST_BorderStyle, c _fbf.Color) {
	if _fd._ef.Bottom == nil {
		_fd._ef.Bottom = _bbg.NewCT_BorderPr()
	}
	_fd._ef.Bottom.Color = _bbg.NewCT_Color()
	_fd._ef.Bottom.Color.RgbAttr = c.AsRGBAString()
	_fd._ef.Bottom.StyleAttr = style
}

// AddMergedCells merges cells within a sheet.
func (_gagc *Sheet) AddMergedCells(fromRef, toRef string) MergedCell {
	if _gagc._ecgc.MergeCells == nil {
		_gagc._ecgc.MergeCells = _bbg.NewCT_MergeCells()
	}
	_ceef := _bbg.NewCT_MergeCell()
	_ceef.RefAttr = _cg.Sprintf("\u0025\u0073\u003a%\u0073", fromRef, toRef)
	_gagc._ecgc.MergeCells.MergeCell = append(_gagc._ecgc.MergeCells.MergeCell, _ceef)
	_gagc._ecgc.MergeCells.CountAttr = _b.Uint32(uint32(len(_gagc._ecgc.MergeCells.MergeCell)))
	return MergedCell{_gagc._afga, _gagc, _ceef}
}

// GetLocked returns true if the cell is locked.
func (_cebe *evalContext) GetLocked(cellRef string) bool {
	return _cebe._befg.Cell(cellRef).getLocked()
}

// Type returns the type of anchor
func (_cfaag OneCellAnchor) Type() AnchorType { return AnchorTypeOneCell }

// ClearSheetViews clears the list of sheet views.  This will clear the results
// of AddView() or SetFrozen.
func (_ggbd *Sheet) ClearSheetViews() { _ggbd._ecgc.SheetViews = nil }

// HasFormula returns true if the cell contains formula.
func (_eega *evalContext) HasFormula(cellRef string) bool {
	return _eega._befg.Cell(cellRef).HasFormula()
}

// SetDate sets the cell value to a date. It's stored as the number of days past
// th sheet epoch. When we support v5 strict, we can store an ISO 8601 date
// string directly, however that's not allowed with v5 transitional  (even
// though it works in Excel). The cell is not styled via this method, so it will
// display as a number. SetDateWithStyle should normally be used instead.
func (_bee Cell) SetDate(d _be.Time) {
	_bee.clearValue()
	d = _dcc(d)
	_fec := _bee._cea.Epoch()
	if d.Before(_fec) {
		_ca.Log.Debug("d\u0061\u0074\u0065\u0073\u0020\u0062e\u0066\u006f\u0072\u0065\u0020\u00319\u0030\u0030\u0020\u0061\u0072\u0065\u0020n\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064")
		return
	}
	_dfg := d.Sub(_fec)
	_bfbb := new(_gbg.Float)
	_daa := new(_gbg.Float)
	_daa.SetPrec(128)
	_daa.SetUint64(uint64(_dfg))
	_badb := new(_gbg.Float)
	_badb.SetUint64(24 * 60 * 60 * 1e9)
	_bfbb.Quo(_daa, _badb)
	_bccc, _ := _bfbb.Uint64()
	_bee._age.V = _b.Stringf("\u0025\u0064", _bccc)
}

// SetRowOffset sets the row offset of the top-left anchor.
func (_afa OneCellAnchor) SetRowOffset(m _bg.Distance) { _afa.TopLeft().SetRowOffset(m) }

// SetNumber sets the cell type to number, and the value to the given number
func (_dcfd Cell) SetNumber(v float64) {
	_dcfd.clearValue()
	if _f.IsNaN(v) || _f.IsInf(v, 0) {
		_dcfd._age.TAttr = _bbg.ST_CellTypeE
		_dcfd._age.V = _b.String("\u0023\u004e\u0055M\u0021")
		return
	}
	_dcfd._age.TAttr = _bbg.ST_CellTypeN
	_dcfd._age.V = _b.String(_dg.FormatFloat(v, 'f', -1, 64))
}
func (_dbf RichTextRun) ensureRpr() {
	if _dbf._ffae.RPr == nil {
		_dbf._ffae.RPr = _bbg.NewCT_RPrElt()
	}
}

// Validate attempts to validate the structure of a workbook.
func (_gfab *Workbook) Validate() error {
	if _gfab == nil || _gfab._bdff == nil {
		return _fb.New("\u0077o\u0072\u006bb\u006f\u006f\u006b\u0020n\u006f\u0074\u0020i\u006e\u0069\u0074\u0069\u0061\u006c\u0069\u007a\u0065d \u0063\u006f\u0072r\u0065\u0063t\u006c\u0079\u002c\u0020\u006e\u0069l\u0020\u0062a\u0073\u0065")
	}
	_cace := uint32(0)
	for _, _fgbb := range _gfab._bdff.Sheets.Sheet {
		if _fgbb.SheetIdAttr > _cace {
			_cace = _fgbb.SheetIdAttr
		}
	}
	if _cace != uint32(len(_gfab._cfgf)) {
		return _cg.Errorf("\u0066\u006f\u0075\u006e\u0064\u0020%\u0064\u0020\u0077\u006f\u0072\u006b\u0073\u0068\u0065\u0065\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074i\u006f\u006e\u0073\u0020\u0061\u006e\u0064\u0020\u0025\u0064\u0020\u0077\u006f\u0072k\u0073h\u0065\u0065\u0074\u0073", _cace, len(_gfab._cfgf))
	}
	_defeg := map[string]struct{}{}
	for _fade, _abagg := range _gfab._bdff.Sheets.Sheet {
		_ebbaf := Sheet{_gfab, _abagg, _gfab._cfgf[_fade]}
		if _, _ggdb := _defeg[_ebbaf.Name()]; _ggdb {
			return _cg.Errorf("\u0077\u006f\u0072k\u0062\u006f\u006f\u006b\u002f\u0053\u0068\u0065\u0065\u0074\u005b\u0025\u0064\u005d\u0020\u0068\u0061\u0073\u0020\u0064\u0075\u0070\u006c\u0069\u0063\u0061\u0074\u0065\u0020n\u0061\u006d\u0065\u0020\u0027\u0025\u0073\u0027", _fade, _ebbaf.Name())
		}
		_defeg[_ebbaf.Name()] = struct{}{}
		if _egdf := _ebbaf.ValidateWithPath(_cg.Sprintf("\u0077o\u0072k\u0062\u006f\u006f\u006b\u002fS\u0068\u0065e\u0074\u005b\u0025\u0064\u005d", _fade)); _egdf != nil {
			return _egdf
		}
		if _bgae := _ebbaf.Validate(); _bgae != nil {
			return _bgae
		}
	}
	return nil
}

// LessCells returns true if the lhs value is less than the rhs value. If the
// cells contain numeric values, their value interpreted as a floating point is
// compared. Otherwise their string contents are compared.
func (_ffb Comparer) LessCells(lhs, rhs Cell) bool {
	if _ffb.Order == SortOrderDescending {
		lhs, rhs = rhs, lhs
	}
	if lhs.X() == nil {
		if rhs.X() == nil {
			return false
		}
		return true
	}
	if rhs.X() == nil {
		return false
	}
	_cdce, _feda := lhs.getRawSortValue()
	_bbga, _babd := rhs.getRawSortValue()
	switch {
	case _feda && _babd:
		_fgge, _ := _dg.ParseFloat(_cdce, 64)
		_dbgg, _ := _dg.ParseFloat(_bbga, 64)
		return _fgge < _dbgg
	case _feda:
		return true
	case _babd:
		return false
	}
	_cdce = lhs.GetFormattedValue()
	_bbga = rhs.GetFormattedValue()
	return _cdce < _bbga
}

// MakeComments constructs a new Comments wrapper.
func MakeComments(w *Workbook, x *_bbg.Comments) Comments { return Comments{w, x} }

// AddSheet adds a new sheet to a workbook.
func (_dffa *Workbook) AddSheet() Sheet {
	_cgbe := _bbg.NewCT_Sheet()
	_cgbe.SheetIdAttr = 1
	for _, _abffa := range _dffa._bdff.Sheets.Sheet {
		if _cgbe.SheetIdAttr <= _abffa.SheetIdAttr {
			_cgbe.SheetIdAttr = _abffa.SheetIdAttr + 1
		}
	}
	_dffa._bdff.Sheets.Sheet = append(_dffa._bdff.Sheets.Sheet, _cgbe)
	_cgbe.NameAttr = _cg.Sprintf("\u0053\u0068\u0065\u0065\u0074\u0020\u0025\u0064", _cgbe.SheetIdAttr)
	_bgba := _bbg.NewWorksheet()
	_bgba.Dimension = _bbg.NewCT_SheetDimension()
	_bgba.Dimension.RefAttr = "\u0041\u0031"
	_dffa._cfgf = append(_dffa._cfgf, _bgba)
	_gfaag := _dcb.NewRelationships()
	_dffa._fgdcb = append(_dffa._fgdcb, _gfaag)
	_bgba.SheetData = _bbg.NewCT_SheetData()
	_dffa._cffde = append(_dffa._cffde, nil)
	_ceadd := _b.DocTypeSpreadsheet
	_ecdgg := _dffa._fdae.AddAutoRelationship(_ceadd, _b.OfficeDocumentType, len(_dffa._bdff.Sheets.Sheet), _b.WorksheetType)
	_cgbe.IdAttr = _ecdgg.ID()
	_dffa.ContentTypes.AddOverride(_b.AbsoluteFilename(_ceadd, _b.WorksheetContentType, len(_dffa._bdff.Sheets.Sheet)), _b.WorksheetContentType)
	return Sheet{_dffa, _cgbe, _bgba}
}

// IsWindowLocked returns whether the workbook windows are locked.
func (_caeb WorkbookProtection) IsWindowLocked() bool {
	return _caeb._fefe.LockWindowsAttr != nil && *_caeb._fefe.LockWindowsAttr
}

// X returns the inner wrapped XML type.
func (_daaa ConditionalFormattingRule) X() *_bbg.CT_CfRule { return _daaa._daf }
func (_efbe ConditionalFormattingRule) clear() {
	_efbe._daf.OperatorAttr = _bbg.ST_ConditionalFormattingOperatorUnset
	_efbe._daf.ColorScale = nil
	_efbe._daf.IconSet = nil
	_efbe._daf.Formula = nil
}
func _ffaf(_cgdc string, _eebga uint32, _bcce bool) string {
	_dcee, _dgad, _dbdc := _df.ParseRangeReference(_cgdc)
	if _dbdc == nil {
		_dgdf, _fedc := _dcee.ColumnIdx, _dgad.ColumnIdx
		if _eebga >= _dgdf && _eebga <= _fedc {
			if _dgdf == _fedc {
				if _bcce {
					return ""
				} else {
					return _cgdc
				}
			} else {
				_fceg := _dgad.Update(_gd.UpdateActionRemoveColumn)
				return _cg.Sprintf("\u0025\u0073\u003a%\u0073", _dcee.String(), _fceg.String())
			}
		} else if _eebga < _dgdf {
			_ccda := _dcee.Update(_gd.UpdateActionRemoveColumn)
			_efcg := _dgad.Update(_gd.UpdateActionRemoveColumn)
			return _cg.Sprintf("\u0025\u0073\u003a%\u0073", _ccda.String(), _efcg.String())
		}
	} else {
		_bebd, _ddfa, _dfge := _df.ParseColumnRangeReference(_cgdc)
		if _dfge != nil {
			return ""
		}
		_bfeee, _dgccb := _bebd.ColumnIdx, _ddfa.ColumnIdx
		if _eebga >= _bfeee && _eebga <= _dgccb {
			if _bfeee == _dgccb {
				if _bcce {
					return ""
				} else {
					return _cgdc
				}
			} else {
				_eegb := _ddfa.Update(_gd.UpdateActionRemoveColumn)
				return _cg.Sprintf("\u0025\u0073\u003a%\u0073", _bebd.String(), _eegb.String())
			}
		} else if _eebga < _bfeee {
			_edcc := _bebd.Update(_gd.UpdateActionRemoveColumn)
			_dfgbc := _ddfa.Update(_gd.UpdateActionRemoveColumn)
			return _cg.Sprintf("\u0025\u0073\u003a%\u0073", _edcc.String(), _dfgbc.String())
		}
	}
	return ""
}

// Epoch returns the point at which the dates/times in the workbook are relative to.
func (_gced *Workbook) Epoch() _be.Time {
	if _gced.Uses1904Dates() {
		_be.Date(1904, 1, 1, 0, 0, 0, 0, _be.UTC)
	}
	return _be.Date(1899, 12, 30, 0, 0, 0, 0, _be.UTC)
}

// SetRowOffset sets the row offset of the two cell anchor
func (_adfb TwoCellAnchor) SetRowOffset(m _bg.Distance) {
	_dbefb := m - _adfb.TopLeft().RowOffset()
	_adfb.TopLeft().SetRowOffset(m)
	_adfb.BottomRight().SetRowOffset(_adfb.BottomRight().RowOffset() + _dbefb)
}

// Drawing is a drawing overlay on a sheet.  Only a single drawing is allowed
// per sheet, so to display multiple charts and images on a single sheet, they
// must be added to the same drawing.
type Drawing struct {
	_eafd *Workbook
	_abd  *_ag.WsDr
}

// X returns the inner wrapped XML type.
func (_bcfg Column) X() *_bbg.CT_Col { return _bcfg._ggba }

// SetWidthCells is a no-op.
func (_fea AbsoluteAnchor) SetWidthCells(int32) {}

// RecalculateFormulas re-computes any computed formula values that are stored
// in the sheet. As github.com/arcpd/msword formula support is still new and not all functins are
// supported, if formula execution fails either due to a parse error or missing
// function, or erorr in the result (even if expected) the cached value will be
// left empty allowing Excel to recompute it on load.
func (_dfbg *Workbook) RecalculateFormulas() {
	for _, _cbgd := range _dfbg.Sheets() {
		_cbgd.RecalculateFormulas()
	}
}

// SetRichTextString sets the cell to rich string mode and returns a struct that
// can be used to add formatted text to the cell.
func (_cdc Cell) SetRichTextString() RichText {
	_cdc.clearValue()
	_cdc._age.Is = _bbg.NewCT_Rst()
	_cdc._age.TAttr = _bbg.ST_CellTypeInlineStr
	return RichText{_cdc._age.Is}
}

// Sort sorts all of the rows within a sheet by the contents of a column. As the
// file format doesn't suppot indicating that a column should be sorted by the
// viewing/editing program, we actually need to reorder rows and change cell
// references during a sort. If the sheet contains formulas, you should call
// RecalculateFormulas() prior to sorting.  The column is in the form "C" and
// specifies the column to sort by. The firstRow is a 1-based index and
// specifies the firstRow to include in the sort, allowing skipping over a
// header row.
func (_gae *Sheet) Sort(column string, firstRow uint32, order SortOrder) {
	_beea := _gae._ecgc.SheetData.Row
	_ceac := _gae.Rows()
	for _gfaac, _eae := range _ceac {
		if _eae.RowNumber() == firstRow {
			_beea = _gae._ecgc.SheetData.Row[_gfaac:]
			break
		}
	}
	_fbac := Comparer{Order: order}
	_c.Slice(_beea, func(_gcbbc, _fcgag int) bool {
		return _fbac.LessRows(column, Row{_gae._afga, _gae, _beea[_gcbbc]}, Row{_gae._afga, _gae, _beea[_fcgag]})
	})
	for _eefg, _gdee := range _gae.Rows() {
		_eaed := uint32(_eefg + 1)
		if _gdee.RowNumber() != _eaed {
			_gdee.renumberAs(_eaed)
		}
	}
}

const _feeb = "\u0053\u006fr\u0074\u004f\u0072\u0064e\u0072\u0041s\u0063\u0065\u006e\u0064\u0069\u006e\u0067\u0053o\u0072\u0074\u004f\u0072\u0064\u0065\u0072\u0044\u0065\u0073\u0063\u0065n\u0064\u0069\u006e\u0067"

// IsDBCS returns if a workbook's default language is among DBCS.
func (_edf *evalContext) IsDBCS() bool {
	_eafc := _edf._befg._afga.CoreProperties.X().Language
	if _eafc == nil {
		return false
	}
	_abc := string(_eafc.Data)
	for _, _ccbd := range _gbcd {
		if _abc == _ccbd {
			return true
		}
	}
	return false
}

// SetFormat sets the number format code.
func (_decf NumberFormat) SetFormat(f string) { _decf._cbd.FormatCodeAttr = f }

// GetOrCreateStandardNumberFormat gets or creates a cell style with a given
// standard format. This should only be used when you want to perform
// number/date/time formatting only.  Manipulating the style returned will cause
// all cells using style returned from this for a given format to be formatted.
func (_adef StyleSheet) GetOrCreateStandardNumberFormat(f StandardFormat) CellStyle {
	for _, _fcdd := range _adef.CellStyles() {
		if _fcdd.HasNumberFormat() && _fcdd.NumberFormat() == uint32(f) {
			return _fcdd
		}
	}
	_egba := _adef.AddCellStyle()
	_egba.SetNumberFormatStandard(f)
	return _egba
}

// SetError sets the cell type to error and the value to the given error message.
func (_gga Cell) SetError(msg string) {
	_gga.clearValue()
	_gga._age.V = _b.String(msg)
	_gga._age.TAttr = _bbg.ST_CellTypeE
}

// SetHeightCells sets the height the anchored object by moving the bottom.  It
// is not compatible with SetHeight.
func (_bbfb TwoCellAnchor) SetHeightCells(h int32) {
	_bbfb.SetHeight(0)
	_bddf := _bbfb.TopLeft()
	_geefa := _bbfb.BottomRight()
	_geefa.SetRow(_bddf.Row() + h)
}

// GetFill gets a Fill from a cell style.
func (_feb CellStyle) GetFill() *_bbg.CT_Fill {
	if _gca := _feb._gcg.FillIdAttr; _gca != nil {
		_feg := _feb._ffgf.StyleSheet.Fills().X().Fill
		if int(*_gca) < len(_feg) {
			return _feg[int(*_gca)]
		}
	}
	return nil
}

// SetPattern sets the pattern of the fill.
func (_cfafa PatternFill) SetPattern(p _bbg.ST_PatternType) { _cfafa._aad.PatternTypeAttr = p }

var (
	_abgb  = [...]uint8{0, 21, 46, 61, 76, 91}
	_fecaa = [...]uint8{0, 21, 37, 53, 69, 85, 103, 119, 135, 151, 167, 185, 201, 217, 239}
	_gadf  = [...]uint8{0, 16, 32, 48, 64}
	_eaag  = [...]uint8{0, 16, 32, 48, 64, 80}
)

// Content returns the content of the defined range (the range in most cases)/
func (_cgc DefinedName) Content() string { return _cgc._dfced.Content }

// Borders returns the list of borders defined in the stylesheet.
func (_baga StyleSheet) Borders() []Border {
	_dbdg := []Border{}
	for _, _bdgf := range _baga._fgbg.Borders.Border {
		_dbdg = append(_dbdg, Border{_ef: _bdgf})
	}
	return _dbdg
}
func (_dcae Cell) GetRawValue() (string, error) {
	switch _dcae._age.TAttr {
	case _bbg.ST_CellTypeInlineStr:
		if _dcae._age.Is == nil || _dcae._age.Is.T == nil {
			return "", nil
		}
		return *_dcae._age.Is.T, nil
	case _bbg.ST_CellTypeS:
		if _dcae._age.V == nil {
			return "", nil
		}
		_gfg, _dfb := _dg.Atoi(*_dcae._age.V)
		if _dfb != nil {
			return "", _dfb
		}
		return _dcae._cea.SharedStrings.GetString(_gfg)
	case _bbg.ST_CellTypeStr:
		if _dcae._age.F != nil {
			return _dcae._age.F.Content, nil
		}
	}
	if _dcae._age.V == nil {
		return "", nil
	}
	return *_dcae._age.V, nil
}

// SetHeightCells is a no-op.
func (_aga OneCellAnchor) SetHeightCells(int32) {}

// RemoveDefinedName removes an existing defined name.
func (_abddg *Workbook) RemoveDefinedName(dn DefinedName) error {
	if dn.X() == nil {
		return _fb.New("\u0061\u0074\u0074\u0065\u006d\u0070t\u0020\u0074\u006f\u0020\u0072\u0065\u006d\u006f\u0076\u0065\u0020\u006e\u0069l\u0020\u0044\u0065\u0066\u0069\u006e\u0065d\u004e\u0061\u006d\u0065")
	}
	for _aded, _aeaf := range _abddg._bdff.DefinedNames.DefinedName {
		if _aeaf == dn.X() {
			copy(_abddg._bdff.DefinedNames.DefinedName[_aded:], _abddg._bdff.DefinedNames.DefinedName[_aded+1:])
			_abddg._bdff.DefinedNames.DefinedName[len(_abddg._bdff.DefinedNames.DefinedName)-1] = nil
			_abddg._bdff.DefinedNames.DefinedName = _abddg._bdff.DefinedNames.DefinedName[:len(_abddg._bdff.DefinedNames.DefinedName)-1]
			return nil
		}
	}
	return _fb.New("\u0064\u0065\u0066\u0069ne\u0064\u0020\u006e\u0061\u006d\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064")
}

// SetBold causes the text to be displayed in bold.
func (_gfce RichTextRun) SetBold(b bool) {
	_gfce.ensureRpr()
	_gfce._ffae.RPr.B = _bbg.NewCT_BooleanProperty()
	_gfce._ffae.RPr.B.ValAttr = _b.Bool(b)
}

// X returns the inner wrapped XML type.
func (_fgeg Font) X() *_bbg.CT_Font { return _fgeg._bdbb }

// DataValidationList is just a view on a DataValidation configured as a list.
// It presents a drop-down combo box for spreadsheet users to select values. The
// contents of the dropdown can either pull from a rang eof cells (SetRange) or
// specified directly (SetValues).
type DataValidationList struct{ _cfg *_bbg.CT_DataValidation }

const (
	AnchorTypeAbsolute AnchorType = iota
	AnchorTypeOneCell
	AnchorTypeTwoCell
)

// GetFormattedValue returns the formatted cell value as it would appear in
// Excel. This involves determining the format string to apply, parsing it, and
// then formatting the value according to the format string.  This should only
// be used if you care about replicating what Excel would show, otherwise
// GetValueAsNumber()/GetValueAsTime
func (_gad Cell) GetFormattedValue() string {
	_egg := _gad.getFormat()
	switch _gad._age.TAttr {
	case _bbg.ST_CellTypeB:
		_cfe, _ := _gad.GetValueAsBool()
		if _cfe {
			return "\u0054\u0052\u0055\u0045"
		}
		return "\u0046\u0041\u004cS\u0045"
	case _bbg.ST_CellTypeN:
		_fed, _ := _gad.GetValueAsNumber()
		return _ce.Number(_fed, _egg)
	case _bbg.ST_CellTypeE:
		if _gad._age.V != nil {
			return *_gad._age.V
		}
		return ""
	case _bbg.ST_CellTypeS, _bbg.ST_CellTypeInlineStr:
		return _ce.String(_gad.GetString(), _egg)
	case _bbg.ST_CellTypeStr:
		_dga := _gad.GetString()
		if _ce.IsNumber(_dga) {
			_ebc, _ := _dg.ParseFloat(_dga, 64)
			return _ce.Number(_ebc, _egg)
		}
		return _ce.String(_dga, _egg)
	case _bbg.ST_CellTypeUnset:
		fallthrough
	default:
		_bde, _ := _gad.GetRawValue()
		if len(_bde) == 0 {
			return ""
		}
		_fagf, _dfcd := _gad.GetValueAsNumber()
		if _dfcd == nil {
			return _ce.Number(_fagf, _egg)
		}
		return _ce.String(_bde, _egg)
	}
}

// StyleSheet is a document style sheet.
type StyleSheet struct {
	_fccg *Workbook
	_fgbg *_bbg.StyleSheet
}

// AddComment adds a new comment and returns a RichText which will contain the
// styled comment text.
func (_efdg Comments) AddComment(cellRef string, author string) RichText {
	_bcda := _bbg.NewCT_Comment()
	_efdg._fcg.CommentList.Comment = append(_efdg._fcg.CommentList.Comment, _bcda)
	_bcda.RefAttr = cellRef
	_bcda.AuthorIdAttr = _efdg.getOrCreateAuthor(author)
	_bcda.Text = _bbg.NewCT_Rst()
	return RichText{_bcda.Text}
}
func (_babe Cell) getLocked() bool {
	if _babe._age.SAttr == nil {
		return false
	}
	_faf := *_babe._age.SAttr
	_gebc := _babe._cea.StyleSheet.GetCellStyle(_faf)
	return *_gebc._gcg.Protection.LockedAttr
}

// RowOffset returns the offset from the row cell.
func (_ad CellMarker) RowOffset() _bg.Distance {
	if _ad._bef.RowOff.ST_CoordinateUnqualified == nil {
		return 0
	}
	return _bg.Distance(float64(*_ad._bef.RowOff.ST_CoordinateUnqualified) * _bg.EMU)
}

// SetActiveSheetIndex sets the index of the active sheet (0-n) which will be
// the tab displayed when the spreadsheet is initially opened.
func (_ebgf *Workbook) SetActiveSheetIndex(idx uint32) {
	if _ebgf._bdff.BookViews == nil {
		_ebgf._bdff.BookViews = _bbg.NewCT_BookViews()
	}
	if len(_ebgf._bdff.BookViews.WorkbookView) == 0 {
		_ebgf._bdff.BookViews.WorkbookView = append(_ebgf._bdff.BookViews.WorkbookView, _bbg.NewCT_BookView())
	}
	_ebgf._bdff.BookViews.WorkbookView[0].ActiveTabAttr = _b.Uint32(idx)
}

// AbsoluteAnchor has a fixed top-left corner in distance units as well as a
// fixed height/width.
type AbsoluteAnchor struct{ _da *_ag.CT_AbsoluteAnchor }

// SetRow set the row of the cell marker.
func (_bcf CellMarker) SetRow(row int32) { _bcf._bef.Row = row }

// SetHeight is a nop-op.
func (_aegc TwoCellAnchor) SetHeight(h _bg.Distance) {}

var _cggde = false

// Author returns the author of the comment
func (_cag Comment) Author() string {
	if _cag._acb.AuthorIdAttr < uint32(len(_cag._fffd.Authors.Author)) {
		return _cag._fffd.Authors.Author[_cag._acb.AuthorIdAttr]
	}
	return ""
}
func (_bae CellStyle) Index() uint32 {
	for _fcf, _edb := range _bae._efb.Xf {
		if _bae._gcg == _edb {
			return uint32(_fcf)
		}
	}
	return 0
}

// X returns the inner wrapped XML type.
func (_cdg DataValidation) X() *_bbg.CT_DataValidation { return _cdg._bge }

// Reference returns the region of cells that are merged.
func (_ggad MergedCell) Reference() string { return _ggad._gfeg.RefAttr }

// SetFormulaRaw sets the cell type to formula, and the raw formula to the given string
func (_bad Cell) SetFormulaRaw(s string) {
	_fg := _eg.ParseString(s)
	if _fg == nil {
		return
	}
	_bad.clearValue()
	_bad._age.TAttr = _bbg.ST_CellTypeStr
	_bad._age.F = _bbg.NewCT_CellFormula()
	_bad._age.F.Content = s
}

// ClearAutoFilter removes the autofilters from the sheet.
func (_dfbac *Sheet) ClearAutoFilter() {
	_dfbac._ecgc.AutoFilter = nil
	_cggbd := "\u0027" + _dfbac.Name() + "\u0027\u0021"
	for _, _bafg := range _dfbac._afga.DefinedNames() {
		if _bafg.Name() == _afcb {
			if _ba.HasPrefix(_bafg.Content(), _cggbd) {
				_dfbac._afga.RemoveDefinedName(_bafg)
				break
			}
		}
	}
}
func (_aabd Fills) X() *_bbg.CT_Fills { return _aabd._egaa }
func (_ed Cell) getFormat() string {
	if _ed._age.SAttr == nil {
		return "\u0047e\u006e\u0065\u0072\u0061\u006c"
	}
	_bfba := *_ed._age.SAttr
	_fag := _ed._cea.StyleSheet.GetCellStyle(_bfba)
	_badg := _ed._cea.StyleSheet.GetNumberFormat(_fag.NumberFormat())
	return _badg.GetFormat()
}

// CellText is used for keeping text with references to a cell where it is located.
type CellText struct {
	Text string
	Cell Cell
}

// Type returns the type of anchor
func (_bc AbsoluteAnchor) Type() AnchorType { return AnchorTypeAbsolute }

// SetHidden marks the defined name as hidden.
func (_ade DefinedName) SetHidden(b bool) { _ade._dfced.HiddenAttr = _b.Bool(b) }

// ExtractText returns text from the workbook as a WorkbookText object.
func (_deaa *Workbook) ExtractText() *WorkbookText {
	_adc := []*SheetText{}
	for _, _ebbg := range _deaa.Sheets() {
		_adc = append(_adc, &SheetText{Cells: _ebbg.ExtractText().Cells})
	}
	return &WorkbookText{Sheets: _adc}
}

// SetPriority sets the rule priority
func (_fac ConditionalFormattingRule) SetPriority(p int32) { _fac._daf.PriorityAttr = p }

// RangeReference converts a range reference of the form 'A1:A5' to 'Sheet
// 1'!$A$1:$A$5 . Renaming a sheet after calculating a range reference will
// invalidate the reference.
func (_ebdb Sheet) RangeReference(n string) string {
	_fcce := _ba.Split(n, "\u003a")
	_gbff, _ := _df.ParseCellReference(_fcce[0])
	_abba := _cg.Sprintf("\u0024\u0025\u0073\u0024\u0025\u0064", _gbff.Column, _gbff.RowIdx)
	if len(_fcce) == 1 {
		return _cg.Sprintf("\u0027%\u0073\u0027\u0021\u0025\u0073", _ebdb.Name(), _abba)
	}
	_ceec, _ := _df.ParseCellReference(_fcce[1])
	_abg := _cg.Sprintf("\u0024\u0025\u0073\u0024\u0025\u0064", _ceec.Column, _ceec.RowIdx)
	return _cg.Sprintf("\u0027\u0025\u0073\u0027\u0021\u0025\u0073\u003a\u0025\u0073", _ebdb.Name(), _abba, _abg)
}

// IsBool returns true if the cell is a boolean type cell.
func (_dbbe Cell) IsBool() bool { return _dbbe._age.TAttr == _bbg.ST_CellTypeB }

// ValidateWithPath validates the sheet passing path informaton for a better
// error message
func (_dbeb Sheet) ValidateWithPath(path string) error { return _dbeb._ecgc.ValidateWithPath(path) }

// Wrapped returns true if the cell will wrap text.
func (_bbde CellStyle) Wrapped() bool {
	if _bbde._gcg.Alignment == nil {
		return false
	}
	if _bbde._gcg.Alignment.WrapTextAttr == nil {
		return false
	}
	return *_bbde._gcg.Alignment.WrapTextAttr
}

var ErrorNotFound = _fb.New("\u006eo\u0074\u0020\u0066\u006f\u0075\u006ed")

// SetStyle sets the cell style for an entire column.
func (_fede Column) SetStyle(cs CellStyle) { _fede._ggba.StyleAttr = _b.Uint32(cs.Index()) }
func (_bfgb DataValidation) clear() {
	_bfgb._bge.Formula1 = _b.String("\u0030")
	_bfgb._bge.Formula2 = _b.String("\u0030")
}
func NewPatternFill(fills *_bbg.CT_Fills) PatternFill {
	_edeb := _bbg.NewCT_Fill()
	_edeb.PatternFill = _bbg.NewCT_PatternFill()
	return PatternFill{_edeb.PatternFill, _edeb}
}

// StandardFormat is a standard ECMA 376 number format.
//
//go:generate stringer -type=StandardFormat
type StandardFormat uint32

// Cell creates or returns a cell given a cell reference of the form 'A10'
func (_eca *Sheet) Cell(cellRef string) Cell {
	_fffe, _ecfe := _df.ParseCellReference(cellRef)
	if _ecfe != nil {
		_ca.Log.Debug("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0063e\u006cl\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u003a\u0020\u0025\u0073", _ecfe)
		return _eca.AddRow().AddCell()
	}
	return _eca.Row(_fffe.RowIdx).Cell(_fffe.Column)
}

// AddDefinedName adds a name for a cell or range reference that can be used in
// formulas and charts.
func (_fdcfa *Workbook) AddDefinedName(name, ref string) DefinedName {
	if _fdcfa._bdff.DefinedNames == nil {
		_fdcfa._bdff.DefinedNames = _bbg.NewCT_DefinedNames()
	}
	_gdde := _bbg.NewCT_DefinedName()
	_gdde.Content = ref
	_gdde.NameAttr = name
	_fdcfa._bdff.DefinedNames.DefinedName = append(_fdcfa._bdff.DefinedNames.DefinedName, _gdde)
	return DefinedName{_gdde}
}

var _gbcd []string = []string{"\u007a\u0068\u002dH\u004b", "\u007a\u0068\u002dM\u004f", "\u007a\u0068\u002dC\u004e", "\u007a\u0068\u002dS\u0047", "\u007a\u0068\u002dT\u0057", "\u006a\u0061\u002dJ\u0050", "\u006b\u006f\u002dK\u0052"}

func (_ecee Fill) Index() uint32 {
	if _ecee._fgc == nil {
		return 0
	}
	for _dcg, _cfaf := range _ecee._fgc.Fill {
		if _ecee._aeag == _cfaf {
			return uint32(_dcg)
		}
	}
	return 0
}
func (_ccdg PatternFill) X() *_bbg.CT_PatternFill { return _ccdg._aad }

// GetValueAsBool retrieves the cell's value as a boolean
func (_gee Cell) GetValueAsBool() (bool, error) {
	if _gee._age.TAttr != _bbg.ST_CellTypeB {
		return false, _fb.New("\u0063e\u006c\u006c\u0020\u0069\u0073\u0020\u006e\u006f\u0074\u0020\u006ff\u0020\u0062\u006f\u006f\u006c\u0020\u0074\u0079\u0070\u0065")
	}
	if _gee._age.V == nil {
		return false, _fb.New("\u0063\u0065\u006c\u006c\u0020\u0068\u0061\u0073\u0020\u006e\u006f\u0020v\u0061\u006c\u0075\u0065")
	}
	return _dg.ParseBool(*_gee._age.V)
}
func (_fcd Comments) getOrCreateAuthor(_gebd string) uint32 {
	for _bcg, _ddd := range _fcd._fcg.Authors.Author {
		if _ddd == _gebd {
			return uint32(_bcg)
		}
	}
	_fgf := uint32(len(_fcd._fcg.Authors.Author))
	_fcd._fcg.Authors.Author = append(_fcd._fcg.Authors.Author, _gebd)
	return _fgf
}

// Cell returns the actual cell behind the merged region
func (_ebcg MergedCell) Cell() Cell {
	_dfga := _ebcg.Reference()
	if _dac := _ba.Index(_ebcg.Reference(), "\u003a"); _dac != -1 {
		_dfga = _dfga[0:_dac]
		return _ebcg._eccd.Cell(_dfga)
	}
	return Cell{}
}

// AddView adds a sheet view.
func (_egad *Sheet) AddView() SheetView {
	if _egad._ecgc.SheetViews == nil {
		_egad._ecgc.SheetViews = _bbg.NewCT_SheetViews()
	}
	_gfda := _bbg.NewCT_SheetView()
	_egad._ecgc.SheetViews.SheetView = append(_egad._ecgc.SheetViews.SheetView, _gfda)
	return SheetView{_gfda}
}

// X returns the inner wrapped XML type.
func (_bdgg ConditionalFormatting) X() *_bbg.CT_ConditionalFormatting { return _bdgg._cff }

// IsSheetLocked returns whether the sheet objects are locked.
func (_becfb SheetProtection) IsObjectLocked() bool {
	return _becfb._ggbb.ObjectsAttr != nil && *_becfb._ggbb.ObjectsAttr
}

// NumberFormat returns the number format that the cell style uses, or zero if
// it is not set.
func (_gcce CellStyle) NumberFormat() uint32 {
	if _gcce._gcg.NumFmtIdAttr == nil {
		return 0
	}
	return *_gcce._gcg.NumFmtIdAttr
}

var _gade *_dc.Regexp = _dc.MustCompile("\u005e(\u005ba\u002d\u007a\u005d\u002b\u0029(\u005b\u0030-\u0039\u005d\u002b\u0029\u0024")

// ID returns the number format ID.  This is not an index as there are some
// predefined number formats which can be used in cell styles and don't need a
// corresponding NumberFormat.
func (_ecfd NumberFormat) ID() uint32 { return _ecfd._cbd.NumFmtIdAttr }
func NewFills() Fills                 { return Fills{_bbg.NewCT_Fills()} }

// NewSharedStrings constructs a new Shared Strings table.
func NewSharedStrings() SharedStrings {
	return SharedStrings{_ddec: _bbg.NewSst(), _ffcg: make(map[string]int)}
}

// GetValueAsTime retrieves the cell's value as a time.  There is no difference
// in SpreadsheetML between a time/date cell other than formatting, and that
// typically a date cell won't have a fractional component. GetValueAsTime will
// work for date cells as well.
func (_ecf Cell) GetValueAsTime() (_be.Time, error) {
	if _ecf._age.TAttr != _bbg.ST_CellTypeUnset {
		return _be.Time{}, _fb.New("\u0063e\u006c\u006c\u0020\u0074y\u0070\u0065\u0020\u0073\u0068o\u0075l\u0064 \u0062\u0065\u0020\u0075\u006e\u0073\u0065t")
	}
	if _ecf._age.V == nil {
		return _be.Time{}, _fb.New("\u0063\u0065\u006c\u006c\u0020\u0068\u0061\u0073\u0020\u006e\u006f\u0020v\u0061\u006c\u0075\u0065")
	}
	_fcb, _, _gceb := _gbg.ParseFloat(*_ecf._age.V, 10, 128, _gbg.ToNearestEven)
	if _gceb != nil {
		return _be.Time{}, _gceb
	}
	_gcca := new(_gbg.Float)
	_gcca.SetUint64(uint64(24 * _be.Hour))
	_fcb.Mul(_fcb, _gcca)
	_cead, _ := _fcb.Uint64()
	_dbc := _ecf._cea.Epoch().Add(_be.Duration(_cead))
	return _fdc(_dbc), nil
}

// RemoveMergedCell removes merging from a cell range within a sheet.  The cells
// that made up the merged cell remain, but are no lon merged.
func (_adff *Sheet) RemoveMergedCell(mc MergedCell) {
	for _bggg, _ebba := range _adff._ecgc.MergeCells.MergeCell {
		if _ebba == mc.X() {
			copy(_adff._ecgc.MergeCells.MergeCell[_bggg:], _adff._ecgc.MergeCells.MergeCell[_bggg+1:])
			_adff._ecgc.MergeCells.MergeCell[len(_adff._ecgc.MergeCells.MergeCell)-1] = nil
			_adff._ecgc.MergeCells.MergeCell = _adff._ecgc.MergeCells.MergeCell[:len(_adff._ecgc.MergeCells.MergeCell)-1]
		}
	}
}

// Border is a cell border configuraton.
type Border struct {
	_ef *_bbg.CT_Border
	_de *_bbg.CT_Borders
}

// SetDateWithStyle sets a date with the default date style applied.
func (_efd Cell) SetDateWithStyle(d _be.Time) {
	_efd.SetDate(d)
	for _, _fae := range _efd._cea.StyleSheet.CellStyles() {
		if _fae.HasNumberFormat() && _fae.NumberFormat() == uint32(StandardFormatDate) {
			_efd.SetStyle(_fae)
			return
		}
	}
	_ffg := _efd._cea.StyleSheet.AddCellStyle()
	_ffg.SetNumberFormatStandard(StandardFormatDate)
	_efd.SetStyle(_ffg)
}

const _afcb = "_\u0078\u006c\u006e\u006d._\u0046i\u006c\u0074\u0065\u0072\u0044a\u0074\u0061\u0062\u0061\u0073\u0065"

// Comments returns the comments for a sheet.
func (_cbff *Sheet) Comments() Comments {
	for _dbce, _gaca := range _cbff._afga._cfgf {
		if _gaca == _cbff._ecgc {
			if _cbff._afga._cffde[_dbce] == nil {
				_cbff._afga._cffde[_dbce] = _bbg.NewComments()
				_cbff._afga._fgdcb[_dbce].AddAutoRelationship(_b.DocTypeSpreadsheet, _b.WorksheetType, _dbce+1, _b.CommentsType)
				_cbff._afga.ContentTypes.AddOverride(_b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.CommentsType, _dbce+1), _b.CommentsContentType)
			}
			if len(_cbff._afga._fdbb) == 0 {
				_cbff._afga._fdbb = append(_cbff._afga._fdbb, _bb.NewCommentDrawing())
				_baea := _cbff._afga._fgdcb[_dbce].AddAutoRelationship(_b.DocTypeSpreadsheet, _b.WorksheetType, 1, _b.VMLDrawingType)
				if _cbff._ecgc.LegacyDrawing == nil {
					_cbff._ecgc.LegacyDrawing = _bbg.NewCT_LegacyDrawing()
				}
				_cbff._ecgc.LegacyDrawing.IdAttr = _baea.ID()
			}
			return Comments{_cbff._afga, _cbff._afga._cffde[_dbce]}
		}
	}
	_ca.Log.Debug("\u0061\u0074\u0074\u0065\u006dp\u0074\u0065\u0064\u0020\u0074\u006f\u0020\u0061\u0063\u0063\u0065\u0073\u0073 \u0063\u006f\u006d\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u006e\u006f\u006e\u002d\u0065\u0078\u0069\u0073\u0074\u0065\u006e\u0074\u0020\u0073\u0068\u0065\u0065t")
	return Comments{}
}

// SetMaxLength sets the maximum bar length in percent.
func (_febb DataBarScale) SetMaxLength(l uint32) { _febb._gdb.MaxLengthAttr = _b.Uint32(l) }
func (_dced *Workbook) ensureSharedStringsRelationships() {
	_fdde := false
	for _, _gdfe := range _dced.ContentTypes.X().Override {
		if _gdfe.ContentTypeAttr == _b.SharedStringsContentType {
			_fdde = true
			break
		}
	}
	if !_fdde {
		_dced.ContentTypes.AddOverride(_ffec, _b.SharedStringsContentType)
	}
	_bgbe := false
	for _, _ead := range _dced._fdae.Relationships() {
		if _ead.X().TargetAttr == _ggcg {
			_bgbe = true
			break
		}
	}
	if !_bgbe {
		_dced._fdae.AddRelationship(_ggcg, _b.SharedStringsType)
	}
}

// CellMarker represents a cell position
type CellMarker struct{ _bef *_ag.CT_Marker }

// CopySheet copies the existing sheet at index `ind` and puts its copy with the name `copiedSheetName`.
func (_gcfec *Workbook) CopySheet(ind int, copiedSheetName string) (Sheet, error) {
	if _gcfec.SheetCount() <= ind {
		return Sheet{}, ErrorNotFound
	}
	var _egbcb _dcb.Relationship
	for _, _cdb := range _gcfec._fdae.Relationships() {
		if _cdb.ID() == _gcfec._bdff.Sheets.Sheet[ind].IdAttr {
			var _ggfeg bool
			if _egbcb, _ggfeg = _gcfec._fdae.CopyRelationship(_cdb.ID()); !_ggfeg {
				return Sheet{}, ErrorNotFound
			}
			break
		}
	}
	_gcfec.ContentTypes.CopyOverride(_b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.WorksheetContentType, ind+1), _b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.WorksheetContentType, len(_gcfec.ContentTypes.X().Override)))
	_ebgaa := *_gcfec._cfgf[ind]
	_gcfec._cfgf = append(_gcfec._cfgf, &_ebgaa)
	var _bdedc uint32 = 0
	for _, _bbbdd := range _gcfec._bdff.Sheets.Sheet {
		if _bbbdd.SheetIdAttr > _bdedc {
			_bdedc = _bbbdd.SheetIdAttr
		}
	}
	_bdedc++
	_ggaf := *_gcfec._bdff.Sheets.Sheet[ind]
	_ggaf.IdAttr = _egbcb.ID()
	_ggaf.NameAttr = copiedSheetName
	_ggaf.SheetIdAttr = _bdedc
	_gcfec._bdff.Sheets.Sheet = append(_gcfec._bdff.Sheets.Sheet, &_ggaf)
	_gdcg := _dcb.NewRelationshipsCopy(_gcfec._fgdcb[ind])
	_gcfec._fgdcb = append(_gcfec._fgdcb, _gdcg)
	_cdgf := _gcfec._cffde[ind]
	if _cdgf == nil {
		_gcfec._cffde = append(_gcfec._cffde, nil)
	} else {
		_cgde := *_cdgf
		_gcfec._cffde = append(_gcfec._cffde, &_cgde)
	}
	_cbfa := Sheet{_gcfec, &_ggaf, &_ebgaa}
	return _cbfa, nil
}

// Reference returns the cell reference (e.g. "A4"). This is not required,
// however both github.com/arcpd/msword and Excel will always set it.
func (_ga Cell) Reference() string {
	if _ga._age.RAttr != nil {
		return *_ga._age.RAttr
	}
	return ""
}

// GetHorizontalAlignment sets the horizontal alignment of a cell style.
func (_ebg CellStyle) GetHorizontalAlignment() _bbg.ST_HorizontalAlignment {
	if _ebg._gcg.Alignment == nil {
		return _bbg.ST_HorizontalAlignmentUnset
	}
	return _ebg._gcg.Alignment.HorizontalAttr
}

// MergedCells returns the merged cell regions within the sheet.
func (_gcfe *Sheet) MergedCells() []MergedCell {
	if _gcfe._ecgc.MergeCells == nil {
		return nil
	}
	_ccgf := []MergedCell{}
	for _, _ggfea := range _gcfe._ecgc.MergeCells.MergeCell {
		_ccgf = append(_ccgf, MergedCell{_gcfe._afga, _gcfe, _ggfea})
	}
	return _ccgf
}

// SetXSplit sets the column split point
func (_ffdd SheetView) SetXSplit(v float64) {
	_ffdd.ensurePane()
	_ffdd._abfc.Pane.XSplitAttr = _b.Float64(v)
}

// SetRowOffset sets the row offset of the top-left of the image in fixed units.
func (_ff AbsoluteAnchor) SetRowOffset(m _bg.Distance) {
	_ff._da.Pos.YAttr.ST_CoordinateUnqualified = _b.Int64(int64(m / _bg.EMU))
}

// SetString sets the cell type to string, and the value to the given string,
// returning an ID from the shared strings table. To reuse a string, call
// SetStringByID with the ID returned.
func (_beg Cell) SetString(s string) int {
	_beg._cea.ensureSharedStringsRelationships()
	_beg.clearValue()
	_bcc := _beg._cea.SharedStrings.AddString(s)
	_beg._age.V = _b.String(_dg.Itoa(_bcc))
	_beg._age.TAttr = _bbg.ST_CellTypeS
	return _bcc
}

// SetFormulaShared sets the cell type to formula shared, and the raw formula to
// the given string. The range is the range of cells that the formula applies
// to, and is used to conserve disk space.
func (_dcf Cell) SetFormulaShared(formulaStr string, rows, cols uint32) error {
	_gcc := _eg.ParseString(formulaStr)
	if _gcc == nil {
		return _fb.New(_cg.Sprintf("\u0043a\u006en\u006f\u0074\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0025\u0073", formulaStr))
	}
	_dcf.clearValue()
	_dcf._age.TAttr = _bbg.ST_CellTypeStr
	_dcf._age.F = _bbg.NewCT_CellFormula()
	_dcf._age.F.TAttr = _bbg.ST_CellFormulaTypeShared
	_dcf._age.F.Content = formulaStr
	_cbe, _fdg := _df.ParseCellReference(_dcf.Reference())
	if _fdg != nil {
		return _fdg
	}
	_bab := uint32(0)
	for _, _bcd := range _dcf._acf.Rows() {
		for _, _bfa := range _bcd._fddd.C {
			if _bfa.F != nil && _bfa.F.SiAttr != nil && *_bfa.F.SiAttr >= _bab {
				_bab = *_bfa.F.SiAttr
			}
		}
	}
	_bab++
	_ffe := _cg.Sprintf("\u0025s\u0025\u0064\u003a\u0025\u0073\u0025d", _cbe.Column, _cbe.RowIdx, _df.IndexToColumn(_cbe.ColumnIdx+cols), _cbe.RowIdx+rows)
	_dcf._age.F.RefAttr = _b.String(_ffe)
	_dcf._age.F.SiAttr = _b.Uint32(_bab)
	_bfad := Sheet{_dcf._cea, _dcf._acf._adgb, _dcf._acf._ecgc}
	for _bdb := _cbe.RowIdx; _bdb <= _cbe.RowIdx+rows; _bdb++ {
		for _aee := _cbe.ColumnIdx; _aee <= _cbe.ColumnIdx+cols; _aee++ {
			if _bdb == _cbe.RowIdx && _aee == _cbe.ColumnIdx {
				continue
			}
			_fef := _cg.Sprintf("\u0025\u0073\u0025\u0064", _df.IndexToColumn(_aee), _bdb)
			_bfad.Cell(_fef).Clear()
			_bfad.Cell(_fef).X().F = _bbg.NewCT_CellFormula()
			_bfad.Cell(_fef).X().F.TAttr = _bbg.ST_CellFormulaTypeShared
			_bfad.Cell(_fef).X().F.SiAttr = _b.Uint32(_bab)
		}
	}
	return nil
}

type DifferentialStyle struct {
	_fdga *_bbg.CT_Dxf
	_fega *Workbook
	_babf *_bbg.CT_Dxfs
}

// X returns the inner wrapped XML type.
func (_dagc Table) X() *_bbg.Table { return _dagc._febfb }

// Protection allows control over the workbook protections.
func (_ceea *Workbook) Protection() WorkbookProtection {
	if _ceea._bdff.WorkbookProtection == nil {
		_ceea._bdff.WorkbookProtection = _bbg.NewCT_WorkbookProtection()
	}
	return WorkbookProtection{_ceea._bdff.WorkbookProtection}
}

// SetPassword sets the password hash to a hash of the input password.
func (_cdddg SheetProtection) SetPassword(pw string) { _cdddg.SetPasswordHash(PasswordHash(pw)) }
func (_bfcf *Sheet) setShared(_edbdc string, _bgbc, _dbebd _df.CellReference, _cebg string) {
	_bbded := _bfcf.FormulaContext()
	_bbgec := _eg.NewEvaluator()
	for _baac := _bgbc.RowIdx; _baac <= _dbebd.RowIdx; _baac++ {
		for _ecgg := _bgbc.ColumnIdx; _ecgg <= _dbebd.ColumnIdx; _ecgg++ {
			_eeea := _baac - _bgbc.RowIdx
			_ffgbd := _ecgg - _bgbc.ColumnIdx
			_bbded.SetOffset(_ffgbd, _eeea)
			_bfee := _bbgec.Eval(_bbded, _cebg)
			_fafb := _cg.Sprintf("\u0025\u0073\u0025\u0064", _df.IndexToColumn(_ecgg), _baac)
			_dggf := _bfcf.Cell(_fafb)
			if _bfee.Type == _eg.ResultTypeNumber {
				_dggf.X().TAttr = _bbg.ST_CellTypeN
			} else {
				_dggf.X().TAttr = _bbg.ST_CellTypeInlineStr
			}
			_dggf.X().V = _b.String(_bfee.Value())
		}
	}
	_ = _bbgec
	_ = _bbded
}

// CellStyles returns the list of defined cell styles
func (_edgb StyleSheet) CellStyles() []CellStyle {
	_abcd := []CellStyle{}
	for _, _ecfc := range _edgb._fgbg.CellXfs.Xf {
		_abcd = append(_abcd, CellStyle{_edgb._fccg, _ecfc, _edgb._fgbg.CellXfs})
	}
	return _abcd
}

// Index returns the index of the border for use with a cell style.
func (_geg Border) Index() uint32 {
	for _bfb, _cb := range _geg._de.Border {
		if _cb == _geg._ef {
			return uint32(_bfb)
		}
	}
	return 0
}

// RemoveSheet removes the sheet with the given index from the workbook.
func (_bgag *Workbook) RemoveSheet(ind int) error {
	if _bgag.SheetCount() <= ind {
		return ErrorNotFound
	}
	for _, _cfega := range _bgag._fdae.Relationships() {
		if _cfega.ID() == _bgag._bdff.Sheets.Sheet[ind].IdAttr {
			_bgag._fdae.Remove(_cfega)
			break
		}
	}
	_bgag.ContentTypes.RemoveOverride(_b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.WorksheetContentType, ind+1))
	copy(_bgag._cfgf[ind:], _bgag._cfgf[ind+1:])
	_bgag._cfgf = _bgag._cfgf[:len(_bgag._cfgf)-1]
	_acfc := _bgag._bdff.Sheets.Sheet[ind]
	copy(_bgag._bdff.Sheets.Sheet[ind:], _bgag._bdff.Sheets.Sheet[ind+1:])
	_bgag._bdff.Sheets.Sheet = _bgag._bdff.Sheets.Sheet[:len(_bgag._bdff.Sheets.Sheet)-1]
	for _gabb := range _bgag._bdff.Sheets.Sheet {
		if _bgag._bdff.Sheets.Sheet[_gabb].SheetIdAttr > _acfc.SheetIdAttr {
			_bgag._bdff.Sheets.Sheet[_gabb].SheetIdAttr--
		}
	}
	copy(_bgag._fgdcb[ind:], _bgag._fgdcb[ind+1:])
	_bgag._fgdcb = _bgag._fgdcb[:len(_bgag._fgdcb)-1]
	copy(_bgag._cffde[ind:], _bgag._cffde[ind+1:])
	_bgag._cffde = _bgag._cffde[:len(_bgag._cffde)-1]
	return nil
}

// AddNumberFormat adds a new blank number format to the stylesheet.
func (_cbda StyleSheet) AddNumberFormat() NumberFormat {
	if _cbda._fgbg.NumFmts == nil {
		_cbda._fgbg.NumFmts = _bbg.NewCT_NumFmts()
	}
	_cdgc := _bbg.NewCT_NumFmt()
	_cdgc.NumFmtIdAttr = uint32(200 + len(_cbda._fgbg.NumFmts.NumFmt))
	_cbda._fgbg.NumFmts.NumFmt = append(_cbda._fgbg.NumFmts.NumFmt, _cdgc)
	_cbda._fgbg.NumFmts.CountAttr = _b.Uint32(uint32(len(_cbda._fgbg.NumFmts.NumFmt)))
	return NumberFormat{_cbda._fccg, _cdgc}
}

// OneCellAnchor is anchored to a top-left cell with a fixed with/height
// in distance.
type OneCellAnchor struct{ _cdcd *_ag.CT_OneCellAnchor }

// GetValueAsNumber retrieves the cell's value as a number
func (_gbf Cell) GetValueAsNumber() (float64, error) {
	if _gbf._age.V == nil && _gbf._age.Is == nil {
		return 0, nil
	}
	if _gbf._age.TAttr == _bbg.ST_CellTypeS || !_ce.IsNumber(*_gbf._age.V) {
		return _f.NaN(), _fb.New("\u0063\u0065\u006c\u006c\u0020\u0069\u0073\u0020\u006e\u006f\u0074 \u006f\u0066\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020t\u0079\u0070\u0065")
	}
	return _dg.ParseFloat(*_gbf._age.V, 64)
}

// SetName sets the sheet name.
func (_gaga *Sheet) SetName(name string) { _gaga._adgb.NameAttr = name }

// ClearBorder clears any border configuration from the cell style.
func (_fbfg CellStyle) ClearBorder() { _fbfg._gcg.BorderIdAttr = nil; _fbfg._gcg.ApplyBorderAttr = nil }

// SaveToFile writes the workbook out to a file.
func (_agdc *Workbook) SaveToFile(path string) error {
	_dcdef, _gdae := _bd.Create(path)
	if _gdae != nil {
		return _gdae
	}
	defer _dcdef.Close()
	return _agdc.Save(_dcdef)
}

// FormulaContext returns a formula evaluation context that can be used to
// evaluate formaulas.
func (_cefg *Sheet) FormulaContext() _eg.Context { return _agd(_cefg) }

// SetCellReference sets the cell reference within a sheet that a comment refers
// to (e.g. "A1")
func (_gbgb Comment) SetCellReference(cellRef string) { _gbgb._acb.RefAttr = cellRef }

// SetColor sets teh color of the databar.
func (_ecb DataBarScale) SetColor(c _fbf.Color) {
	_ecb._gdb.Color = _bbg.NewCT_Color()
	_ecb._gdb.Color.RgbAttr = c.AsRGBAString()
}

// ClearProtection clears all workbook protections.
func (_dagcd *Workbook) ClearProtection() { _dagcd._bdff.WorkbookProtection = nil }

// Text returns text from the sheet as one string separated with line breaks.
func (_fffb *SheetText) Text() string {
	_gafb := _gb.NewBuffer([]byte{})
	for _, _ccbg := range _fffb.Cells {
		if _ccbg.Text != "" {
			_gafb.WriteString(_ccbg.Text)
			_gafb.WriteString("\u000a")
		}
	}
	return _gafb.String()
}

// AddFill creates a new empty Fill style.
func (_aec Fills) AddFill() Fill { _effg := _bbg.NewCT_Fill(); return Fill{_effg, _aec._egaa} }

// AddFont creates a new empty Font style.
func (_ebef StyleSheet) AddFont() Font { _gfge := _bbg.NewCT_Font(); return Font{_gfge, _ebef._fgbg} }

// DataValidation controls cell validation
type DataValidation struct{ _bge *_bbg.CT_DataValidation }

// AddRun adds a new run of text to the cell.
func (_cdga RichText) AddRun() RichTextRun {
	_fbde := _bbg.NewCT_RElt()
	_cdga._agac.R = append(_cdga._agac.R, _fbde)
	return RichTextRun{_fbde}
}
func (_gce Border) SetDiagonal(style _bbg.ST_BorderStyle, c _fbf.Color, up, down bool) {
	if _gce._ef.Diagonal == nil {
		_gce._ef.Diagonal = _bbg.NewCT_BorderPr()
	}
	_gce._ef.Diagonal.Color = _bbg.NewCT_Color()
	_gce._ef.Diagonal.Color.RgbAttr = c.AsRGBAString()
	_gce._ef.Diagonal.StyleAttr = style
	if up {
		_gce._ef.DiagonalUpAttr = _b.Bool(true)
	}
	if down {
		_gce._ef.DiagonalDownAttr = _b.Bool(true)
	}
}

// SetAllowBlank controls if blank values are accepted.
func (_add DataValidation) SetAllowBlank(b bool) {
	if !b {
		_add._bge.AllowBlankAttr = nil
	} else {
		_add._bge.AllowBlankAttr = _b.Bool(true)
	}
}

// SetReference sets the regin of cells that the merged cell applies to.
func (_afbb MergedCell) SetReference(ref string) { _afbb._gfeg.RefAttr = ref }

// RemoveSheetByName removes the sheet with the given name from the workbook.
func (_gdbd *Workbook) RemoveSheetByName(name string) error {
	_cdgd := -1
	for _ecad, _dfef := range _gdbd.Sheets() {
		if name == _dfef.Name() {
			_cdgd = _ecad
			break
		}
	}
	if _cdgd == -1 {
		return ErrorNotFound
	}
	return _gdbd.RemoveSheet(_cdgd)
}

// Comment is a single comment within a sheet.
type Comment struct {
	_ecd  *Workbook
	_acb  *_bbg.CT_Comment
	_fffd *_bbg.Comments
}

// Themes returns the array of workbook dml.Theme.
func (_gcbcf *Workbook) Themes() []*_acg.Theme { return _gcbcf._cbae }
func (_bgfgb *Sheet) setList(_ecaf string, _edc _eg.Result) error {
	_ebgag, _fecfg := _df.ParseCellReference(_ecaf)
	if _fecfg != nil {
		return _fecfg
	}
	_cbc := _bgfgb.Row(_ebgag.RowIdx)
	for _bgea, _eeadg := range _edc.ValueList {
		_abbag := _cbc.Cell(_df.IndexToColumn(_ebgag.ColumnIdx + uint32(_bgea)))
		if _eeadg.Type != _eg.ResultTypeEmpty {
			if _eeadg.IsBoolean {
				_abbag.SetBool(_eeadg.ValueNumber != 0)
			} else {
				_abbag.SetCachedFormulaResult(_eeadg.String())
			}
		}
	}
	return nil
}
func (_bbea *evalContext) SetOffset(col, row uint32) { _bbea._cdec = col; _bbea._cbg = row }

// X returns the inner wrapped XML type.
func (_ddg Cell) X() *_bbg.CT_Cell { return _ddg._age }

// Name returns the name of the defined name.
func (_bdbdf DefinedName) Name() string { return _bdbdf._dfced.NameAttr }
func (_ceb Border) SetTop(style _bbg.ST_BorderStyle, c _fbf.Color) {
	if _ceb._ef.Top == nil {
		_ceb._ef.Top = _bbg.NewCT_BorderPr()
	}
	_ceb._ef.Top.Color = _bbg.NewCT_Color()
	_ceb._ef.Top.Color.RgbAttr = c.AsRGBAString()
	_ceb._ef.Top.StyleAttr = style
}
func (_aeed *Sheet) getAllCellsInFormulaArrays(_gbbf bool) (map[string]bool, error) {
	_ebed := _eg.NewEvaluator()
	_bbc := _aeed.FormulaContext()
	_eafa := map[string]bool{}
	for _, _dddg := range _aeed.Rows() {
		for _, _cdeb := range _dddg.Cells() {
			if _cdeb.X().F != nil {
				_bfadg := _cdeb.X().F.Content
				if _cdeb.X().F.TAttr == _bbg.ST_CellFormulaTypeArray {
					_bgc := _ebed.Eval(_bbc, _bfadg).AsString()
					if _bgc.Type == _eg.ResultTypeError {
						_ca.Log.Debug("\u0065\u0072\u0072o\u0072\u0020\u0065\u0076a\u0075\u006c\u0061\u0074\u0069\u006e\u0067 \u0066\u006f\u0072\u006d\u0075\u006c\u0061\u0020\u0025\u0073\u003a\u0020\u0025\u0073", _bfadg, _bgc.ErrorMessage)
						_cdeb.X().V = nil
					}
					if _bgc.Type == _eg.ResultTypeArray {
						_gbab, _fbcf := _df.ParseCellReference(_cdeb.Reference())
						if _fbcf != nil {
							return map[string]bool{}, _fbcf
						}
						if (_gbbf && len(_bgc.ValueArray) == 1) || (!_gbbf && len(_bgc.ValueArray[0]) == 1) {
							continue
						}
						for _ecdg, _gabe := range _bgc.ValueArray {
							_becg := _gbab.RowIdx + uint32(_ecdg)
							for _fde := range _gabe {
								_cgbb := _df.IndexToColumn(_gbab.ColumnIdx + uint32(_fde))
								_eafa[_cg.Sprintf("\u0025\u0073\u0025\u0064", _cgbb, _becg)] = true
							}
						}
					} else if _bgc.Type == _eg.ResultTypeList {
						_dcde, _bddc := _df.ParseCellReference(_cdeb.Reference())
						if _bddc != nil {
							return map[string]bool{}, _bddc
						}
						if _gbbf || len(_bgc.ValueList) == 1 {
							continue
						}
						_cbad := _dcde.RowIdx
						for _fceb := range _bgc.ValueList {
							_caea := _df.IndexToColumn(_dcde.ColumnIdx + uint32(_fceb))
							_eafa[_cg.Sprintf("\u0025\u0073\u0025\u0064", _caea, _cbad)] = true
						}
					}
				}
			}
		}
	}
	return _eafa, nil
}

// SetWrapped configures the cell to wrap text.
func (_fbe CellStyle) SetWrapped(b bool) {
	if _fbe._gcg.Alignment == nil {
		_fbe._gcg.Alignment = _bbg.NewCT_CellAlignment()
	}
	if !b {
		_fbe._gcg.Alignment.WrapTextAttr = nil
	} else {
		_fbe._gcg.Alignment.WrapTextAttr = _b.Bool(true)
		_fbe._gcg.ApplyAlignmentAttr = _b.Bool(true)
	}
}

// SetHeight sets the height of the anchored object.
func (_cee AbsoluteAnchor) SetHeight(h _bg.Distance) { _cee._da.Ext.CyAttr = int64(h / _bg.EMU) }

// SetColOffset sets the column offset of the two cell anchor.
func (_ddcgb TwoCellAnchor) SetColOffset(m _bg.Distance) {
	_dfdce := m - _ddcgb.TopLeft().ColOffset()
	_ddcgb.TopLeft().SetColOffset(m)
	_ddcgb.BottomRight().SetColOffset(_ddcgb.BottomRight().ColOffset() + _dfdce)
}

// SetFont applies a font to a cell style avoiding redundancy. The function checks if the given font
// already exists in the saved fonts. If found, the existing font is reused; otherwise,
// the new font is added to the saved fonts collection. The font is then applied to the cell style,
// affecting all styles that reference it by index.
func (_fga CellStyle) SetFont(f Font) {
	_gbd := f._gabf.Fonts.Font
	for _, _fdgc := range _gbd {
		if _ge.DeepEqual(_fdgc, f._bdbb) {
			f._bdbb = _fdgc
			_fga._gcg.FontIdAttr = _b.Uint32(f.Index())
			_fga._gcg.ApplyFontAttr = _b.Bool(true)
			return
		}
	}
	f._gabf.Fonts.Font = append(f._gabf.Fonts.Font, f._bdbb)
	f._gabf.Fonts.CountAttr = _b.Uint32(uint32(len(f._gabf.Fonts.Font)))
	_fga._gcg.FontIdAttr = _b.Uint32(f.Index())
	_fga._gcg.ApplyFontAttr = _b.Bool(true)
}

// AddHyperlink creates and sets a hyperlink on a cell.
func (_cc Cell) AddHyperlink(url string) {
	for _dfgb, _ab := range _cc._cea._cfgf {
		if _ab == _cc._acf._ecgc {
			_cc.SetHyperlink(_cc._cea._fgdcb[_dfgb].AddHyperlink(url))
			return
		}
	}
}
func (_dfa *evalContext) Sheet(name string) _eg.Context {
	for _, _gbb := range _dfa._befg._afga.Sheets() {
		if _gbb.Name() == name {
			return _gbb.FormulaContext()
		}
	}
	return _eg.InvalidReferenceContext
}
func _agd(_gaf *Sheet) *evalContext {
	return &evalContext{_befg: _gaf, _bbb: make(map[string]struct{})}
}

// AddCellStyle creates a new empty cell style.
func (_babb StyleSheet) AddCellStyle() CellStyle {
	_cbac := _bbg.NewCT_Xf()
	return CellStyle{_babb._fccg, _cbac, _babb._fgbg.CellXfs}
}

// SetState sets the sheet view state (frozen/split/frozen-split)
func (_ecgb SheetView) SetState(st _bbg.ST_PaneState) {
	_ecgb.ensurePane()
	_ecgb._abfc.Pane.StateAttr = st
}

// SetWidth is a no-op.
func (_fgbf TwoCellAnchor) SetWidth(w _bg.Distance) {}
func (_fcab *Sheet) updateAfterRemove(_cga uint32, _cfae _gd.UpdateAction) error {
	_ecdc := _fcab.Name()
	_facd := &_gd.UpdateQuery{UpdateType: _cfae, ColumnIdx: _cga, SheetToUpdate: _ecdc}
	for _, _ecgf := range _fcab._afga.Sheets() {
		_facd.UpdateCurrentSheet = _ecdc == _ecgf.Name()
		for _, _babee := range _ecgf.Rows() {
			for _, _cbb := range _babee.Cells() {
				if _cbb.X().F != nil {
					_eedf := _cbb.X().F.Content
					_fbec := _eg.ParseString(_eedf)
					if _fbec == nil {
						_cbb.SetError("\u0023\u0052\u0045F\u0021")
					} else {
						_fgfb := _fbec.Update(_facd)
						_cbb.X().F.Content = _cg.Sprintf("\u003d\u0025\u0073", _fgfb.String())
					}
				}
			}
		}
	}
	return nil
}

// GetString returns the string in a cell if it's an inline or string table
// string. Otherwise it returns an empty string.
func (_cde Cell) GetString() string {
	switch _cde._age.TAttr {
	case _bbg.ST_CellTypeInlineStr:
		if _cde._age.Is != nil && _cde._age.Is.T != nil {
			return *_cde._age.Is.T
		}
		if _cde._age.V != nil {
			return *_cde._age.V
		}
	case _bbg.ST_CellTypeS:
		if _cde._age.V == nil {
			return ""
		}
		_cgg, _cfdd := _dg.Atoi(*_cde._age.V)
		if _cfdd != nil {
			return ""
		}
		_eggf, _cfdd := _cde._cea.SharedStrings.GetString(_cgg)
		if _cfdd != nil {
			return ""
		}
		return _eggf
	}
	if _cde._age.V == nil {
		return ""
	}
	return *_cde._age.V
}

// IsBool returns true if the cell boolean value.
func (_afb *evalContext) IsBool(cellRef string) bool { return _afb._befg.Cell(cellRef).IsBool() }
func (_aebf Sheet) IsValid() bool                    { return _aebf._ecgc != nil }

// Validate validates the sheet, returning an error if it is found to be invalid.
func (_bdag Sheet) Validate() error {
	_aadb := []func() error{_bdag.validateRowCellNumbers, _bdag.validateMergedCells, _bdag.validateSheetNames}
	for _, _fcfga := range _aadb {
		if _egd := _fcfga(); _egd != nil {
			return _egd
		}
	}
	if _bdca := _bdag._ecgc.Validate(); _bdca != nil {
		return _bdca
	}
	return _bdag._ecgc.Validate()
}
func _dcc(_ebb _be.Time) _be.Time {
	_ebb = _ebb.Local()
	return _be.Date(_ebb.Year(), _ebb.Month(), _ebb.Day(), _ebb.Hour(), _ebb.Minute(), _ebb.Second(), _ebb.Nanosecond(), _be.UTC)
}

// Row will return a row with a given row number, creating a new row if
// necessary.
func (_bdbe *Sheet) Row(rowNum uint32) Row {
	for _, _bdga := range _bdbe._ecgc.SheetData.Row {
		if _bdga.RAttr != nil && *_bdga.RAttr == rowNum {
			return Row{_bdbe._afga, _bdbe, _bdga}
		}
	}
	return _bdbe.AddNumberedRow(rowNum)
}
func (_dddbb PatternFill) ClearBgColor() { _dddbb._aad.BgColor = nil }

// GetEpoch returns a workbook's time epoch.
func (_bfbf *evalContext) GetEpoch() _be.Time { return _bfbf._befg._afga.Epoch() }

// SetIcons sets the icon set to use for display.
func (_cdeca IconScale) SetIcons(t _bbg.ST_IconSetType) { _cdeca._bfdc.IconSetAttr = t }

// SetContent sets the defined name content.
func (_dgcc DefinedName) SetContent(s string) { _dgcc._dfced.Content = s }

// Cells returns a slice of cells.  The cells can be manipulated, but appending
// to the slice will have no effect.
func (_acgc Row) Cells() []Cell {
	_cbf := []Cell{}
	_ebe := -1
	_cccf := append([]*_bbg.CT_Cell{}, _acgc._fddd.C...)
	for _, _cggd := range _cccf {
		if _cggd.RAttr == nil {
			_ca.Log.Debug("\u0052\u0041\u0074tr\u0020\u0069\u0073\u0020\u006e\u0069\u006c\u0020\u0066o\u0072 \u0061 \u0063e\u006c\u006c\u002c\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e")
			continue
		}
		_fadbe, _ceed := _df.ParseCellReference(*_cggd.RAttr)
		if _ceed != nil {
			_ca.Log.Debug("\u0052\u0041\u0074t\u0072\u0020\u0069\u0073 \u0069\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0066\u006f\u0072\u0020\u0061\u0020\u0063\u0065\u006c\u006c\u003a\u0020" + *_cggd.RAttr + ",\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u002e")
			continue
		}
		_ccbf := int(_fadbe.ColumnIdx)
		if _ccbf-_ebe > 1 {
			for _cebf := _ebe + 1; _cebf < _ccbf; _cebf++ {
				_cbf = append(_cbf, _acgc.Cell(_df.IndexToColumn(uint32(_cebf))))
			}
		}
		_ebe = _ccbf
		_cbf = append(_cbf, Cell{_acgc._fadd, _acgc._bafc, _acgc._fddd, _cggd})
	}
	return _cbf
}
func (_bcfc StyleSheet) GetNumberFormat(id uint32) NumberFormat {
	if id >= 0 && id < 50 {
		return CreateDefaultNumberFormat(StandardFormat(id))
	}
	for _, _cddg := range _bcfc._fgbg.NumFmts.NumFmt {
		if _cddg.NumFmtIdAttr == id {
			return NumberFormat{_bcfc._fccg, _cddg}
		}
	}
	return NumberFormat{}
}

// RecalculateFormulas re-computes any computed formula values that are stored
// in the sheet. As github.com/arcpd/msword formula support is still new and not all functins are
// supported,  if formula execution fails either due to a parse error or missing
// function, or erorr in the result (even if expected) the cached value will be
// left empty allowing Excel to recompute it on load.
func (_ddeb *Sheet) RecalculateFormulas() {
	_gfed := _eg.NewEvaluator()
	_dee := _ddeb.FormulaContext()
	for _, _abag := range _ddeb.Rows() {
		for _, _bbdf := range _abag.Cells() {
			if _bbdf.X().F != nil {
				_ffbd := _bbdf.X().F.Content
				if _bbdf.X().F.TAttr == _bbg.ST_CellFormulaTypeShared && len(_ffbd) == 0 {
					continue
				}
				_gbea := _gfed.Eval(_dee, _ffbd).AsString()
				if _gbea.Type == _eg.ResultTypeError {
					_ca.Log.Debug("\u0065\u0072\u0072o\u0072\u0020\u0065\u0076a\u0075\u006c\u0061\u0074\u0069\u006e\u0067 \u0066\u006f\u0072\u006d\u0075\u006c\u0061\u0020\u0025\u0073\u003a\u0020\u0025\u0073", _ffbd, _gbea.ErrorMessage)
					_bbdf.X().V = nil
				} else {
					if _gbea.Type == _eg.ResultTypeNumber {
						_bbdf.X().TAttr = _bbg.ST_CellTypeN
					} else {
						_bbdf.X().TAttr = _bbg.ST_CellTypeInlineStr
					}
					_bbdf.X().V = _b.String(_gbea.Value())
					if _bbdf.X().F.TAttr == _bbg.ST_CellFormulaTypeArray {
						if _gbea.Type == _eg.ResultTypeArray {
							_ddeb.setArray(_bbdf.Reference(), _gbea)
						} else if _gbea.Type == _eg.ResultTypeList {
							_ddeb.setList(_bbdf.Reference(), _gbea)
						}
					} else if _bbdf.X().F.TAttr == _bbg.ST_CellFormulaTypeShared && _bbdf.X().F.RefAttr != nil {
						_bbfe, _fceac, _dgee := _df.ParseRangeReference(*_bbdf.X().F.RefAttr)
						if _dgee != nil {
							_ca.Log.Debug("\u0065\u0072r\u006f\u0072\u0020\u0069n\u0020\u0073h\u0061\u0072\u0065\u0064\u0020\u0066\u006f\u0072m\u0075\u006c\u0061\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u003a\u0020\u0025\u0073", _dgee)
							continue
						}
						_ddeb.setShared(_bbdf.Reference(), _bbfe, _fceac, _ffbd)
					}
				}
			}
		}
	}
}

// SheetText is an array of extracted text items which has some methods for representing extracted text from a sheet.
type SheetText struct{ Cells []CellText }

// Comparer is used to compare rows based off a column and cells based off of
// their value.
type Comparer struct{ Order SortOrder }

// X returns the inner wrapped XML type.
func (_ceaa SheetView) X() *_bbg.CT_SheetView { return _ceaa._abfc }

// SharedStrings is a shared strings table, where string data can be placed
// outside of the sheet contents and referenced from a sheet.
type SharedStrings struct {
	_ddec *_bbg.Sst
	_ffcg map[string]int
}

// AddFormatValue adds a format value to be used to determine the cell background.
func (_baf ColorScale) AddFormatValue(t _bbg.ST_CfvoType, val string) {
	_fdca := _bbg.NewCT_Cfvo()
	_fdca.TypeAttr = t
	_fdca.ValAttr = _b.String(val)
	_baf._gfa.Cfvo = append(_baf._gfa.Cfvo, _fdca)
}

// X returns the inner wrapped XML type.
func (_dddc RichTextRun) X() *_bbg.CT_RElt { return _dddc._ffae }

// Comments is the container for comments for a single sheet.
type Comments struct {
	_fgb *Workbook
	_fcg *_bbg.Comments
}

// PasswordHash returns the hash of the workbook password.
func (_egde WorkbookProtection) PasswordHash() string {
	if _egde._fefe.WorkbookPasswordAttr == nil {
		return ""
	}
	return *_egde._fefe.WorkbookPasswordAttr
}

// ExtractText returns text from the sheet as a SheetText object.
func (_cegb *Sheet) ExtractText() *SheetText {
	_gdeb := []CellText{}
	for _, _gcgf := range _cegb.Rows() {
		for _, _agfe := range _gcgf.Cells() {
			if !_agfe.IsEmpty() {
				if _feea := _agfe.GetFormattedValue(); _feea != "" {
					_gdeb = append(_gdeb, CellText{Text: _feea, Cell: _agfe})
				}
			}
		}
	}
	return &SheetText{Cells: _gdeb}
}
func _egfef() *_ag.CT_AbsoluteAnchor { _gcfg := _ag.NewCT_AbsoluteAnchor(); return _gcfg }

// IsEmpty returns true if the cell is empty.
func (_bbe Cell) IsEmpty() bool {
	return _bbe._age.TAttr == _bbg.ST_CellTypeUnset && _bbe._age.V == nil && _bbe._age.F == nil
}

// NumberFormat is a number formatting string that can be applied to a cell
// style.
type NumberFormat struct {
	_abf *Workbook
	_cbd *_bbg.CT_NumFmt
}
type evalContext struct {
	_befg       *Sheet
	_cdec, _cbg uint32
	_bbb        map[string]struct{}
}

// X returns the inner wrapped XML type.
func (_aced WorkbookProtection) X() *_bbg.CT_WorkbookProtection { return _aced._fefe }

// SetText sets the text to be displayed.
func (_fcbd RichTextRun) SetText(s string) { _fcbd._ffae.T = s }
func _adec() *_ag.CT_OneCellAnchor         { _dfag := _ag.NewCT_OneCellAnchor(); return _dfag }

// SetOperator sets the operator for the rule.
func (_ecc ConditionalFormattingRule) SetOperator(t _bbg.ST_ConditionalFormattingOperator) {
	_ecc._daf.OperatorAttr = t
}

// SetFont sets the font name for a rich text run.
func (_dbef RichTextRun) SetFont(s string) {
	_dbef.ensureRpr()
	_dbef._ffae.RPr.RFont = _bbg.NewCT_FontName()
	_dbef._ffae.RPr.RFont.ValAttr = s
}

// X returns the inner wrapped XML type.
func (_cddd Sheet) X() *_bbg.Worksheet { return _cddd._ecgc }

// X returns the inner wrapped XML type.
func (_eeb Border) X() *_bbg.CT_Border                 { return _eeb._ef }
func (_ebga DataValidationCompare) SetValue2(v string) { _ebga._fad.Formula2 = &v }

// HasNumberFormat returns true if the cell style has a number format applied.
func (_efe CellStyle) HasNumberFormat() bool {
	return _efe._gcg.NumFmtIdAttr != nil && _efe._gcg.ApplyNumberFormatAttr != nil && *_efe._gcg.ApplyNumberFormatAttr
}
func (_dgff PatternFill) ClearFgColor() { _dgff._aad.FgColor = nil }

// Reference returns the table reference (the cells within the table)
func (_aaba Table) Reference() string { return _aaba._febfb.RefAttr }
func (_bbf Font) Index() uint32 {
	for _fgcd, _abcc := range _bbf._gabf.Fonts.Font {
		if _bbf._bdbb == _abcc {
			return uint32(_fgcd)
		}
	}
	return 0
}
func (_eeaa *Sheet) getAllCellsInFormulaArraysForColumn() (map[string]bool, error) {
	return _eeaa.getAllCellsInFormulaArrays(false)
}

// CellStyle is a formatting style for a cell. CellStyles are spreadsheet global
// and can be applied to cells across sheets.
type CellStyle struct {
	_ffgf *Workbook
	_gcg  *_bbg.CT_Xf
	_efb  *_bbg.CT_CellXfs
}

// Sheet is a single sheet within a workbook.
type Sheet struct {
	_afga *Workbook
	_adgb *_bbg.CT_Sheet
	_ecgc *_bbg.Worksheet
}

// AddChart adds an chart to a drawing, returning the chart and an anchor that
// can be used to position the chart within the sheet.
func (_cafa Drawing) AddChart(at AnchorType) (_ec.Chart, Anchor) {
	_edea := _fe.NewChartSpace()
	_cafa._eafd._ccbc = append(_cafa._eafd._ccbc, _edea)
	_geef := _b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.ChartContentType, len(_cafa._eafd._ccbc))
	_cafa._eafd.ContentTypes.AddOverride(_geef, _b.ChartContentType)
	var _baa string
	for _egb, _aeee := range _cafa._eafd._eded {
		if _aeee == _cafa._abd {
			_cef := _b.RelativeFilename(_b.DocTypeSpreadsheet, _b.DrawingType, _b.ChartType, len(_cafa._eafd._ccbc))
			_dbe := _cafa._eafd._ccag[_egb].AddRelationship(_cef, _b.ChartType)
			_baa = _dbe.ID()
			break
		}
	}
	var _fadb Anchor
	var _ecff *_ag.CT_GraphicalObjectFrame
	switch at {
	case AnchorTypeAbsolute:
		_bcab := _egfef()
		_cafa._abd.EG_Anchor = append(_cafa._abd.EG_Anchor, &_ag.EG_Anchor{AbsoluteAnchor: _bcab})
		_bcab.Choice = &_ag.EG_ObjectChoicesChoice{}
		_bcab.Choice.GraphicFrame = _ag.NewCT_GraphicalObjectFrame()
		_ecff = _bcab.Choice.GraphicFrame
		_fadb = AbsoluteAnchor{_bcab}
	case AnchorTypeOneCell:
		_egcd := _adec()
		_cafa._abd.EG_Anchor = append(_cafa._abd.EG_Anchor, &_ag.EG_Anchor{OneCellAnchor: _egcd})
		_egcd.Choice = &_ag.EG_ObjectChoicesChoice{}
		_egcd.Choice.GraphicFrame = _ag.NewCT_GraphicalObjectFrame()
		_ecff = _egcd.Choice.GraphicFrame
		_fadb = OneCellAnchor{_egcd}
	case AnchorTypeTwoCell:
		_bded := _gaag()
		_cafa._abd.EG_Anchor = append(_cafa._abd.EG_Anchor, &_ag.EG_Anchor{TwoCellAnchor: _bded})
		_bded.Choice = &_ag.EG_ObjectChoicesChoice{}
		_bded.Choice.GraphicFrame = _ag.NewCT_GraphicalObjectFrame()
		_ecff = _bded.Choice.GraphicFrame
		_fadb = TwoCellAnchor{_bded}
	}
	_ecff.NvGraphicFramePr = _ag.NewCT_GraphicalObjectFrameNonVisual()
	_ecff.NvGraphicFramePr.CNvPr.IdAttr = uint32(len(_cafa._abd.EG_Anchor))
	_ecff.NvGraphicFramePr.CNvPr.NameAttr = "\u0043\u0068\u0061r\u0074"
	_ecff.Graphic = _acg.NewGraphic()
	_ecff.Graphic.GraphicData.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002eo\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073.\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006dl/\u0032\u0030\u0030\u0036\u002f\u0063\u0068\u0061\u0072\u0074"
	_eegac := _fe.NewChart()
	_eegac.IdAttr = _baa
	_ecff.Graphic.GraphicData.Any = []_b.Any{_eegac}
	_gfbd := _ec.MakeChart(_edea)
	_gfbd.Properties().SetSolidFill(_fbf.White)
	_gfbd.SetDisplayBlanksAs(_fe.ST_DispBlanksAsGap)
	return _gfbd, _fadb
}

const (
	SortOrderAscending SortOrder = iota
	SortOrderDescending
)

// SetHidden hides or unhides the row
func (_aedc Row) SetHidden(hidden bool) {
	if !hidden {
		_aedc._fddd.HiddenAttr = nil
	} else {
		_aedc._fddd.HiddenAttr = _b.Bool(true)
	}
}

// AddFormatValue adds a format value to be used in determining which icons to display.
func (_ffcb IconScale) AddFormatValue(t _bbg.ST_CfvoType, val string) {
	_fcae := _bbg.NewCT_Cfvo()
	_fcae.TypeAttr = t
	_fcae.ValAttr = _b.String(val)
	_ffcb._bfdc.Cfvo = append(_ffcb._bfdc.Cfvo, _fcae)
}

// SetWidthCells sets the height the anchored object by moving the right hand
// side. It is not compatible with SetWidth.
func (_bgcd TwoCellAnchor) SetWidthCells(w int32) {
	_gdf := _bgcd.TopLeft()
	_dfbag := _bgcd.BottomRight()
	_dfbag.SetCol(_gdf.Col() + w)
}
func (_fcea PatternFill) SetBgColor(c _fbf.Color) {
	_fcea._aad.BgColor = _bbg.NewCT_Color()
	_fcea._aad.BgColor.RgbAttr = c.AsRGBAString()
}

var _ffec = _b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.SharedStringsType, 0)

// SetValue sets the first value to be used in the comparison.  For comparisons
// that need only one value, this is the only value used.  For comparisons like
// 'between' that require two values, SetValue2 must also be used.
func (_cac DataValidationCompare) SetValue(v string) { _cac._fad.Formula1 = &v }

// SetMinLength sets the minimum bar length in percent.
func (_befgg DataBarScale) SetMinLength(l uint32) { _befgg._gdb.MinLengthAttr = _b.Uint32(l) }

// SetFormulaArray sets the cell type to formula array, and the raw formula to
// the given string. This is equivlent to entering a formula and pressing
// Ctrl+Shift+Enter in Excel.
func (_fa Cell) SetFormulaArray(s string) {
	_afc := _eg.ParseString(s)
	if _afc == nil {
		return
	}
	_fa.clearValue()
	_fa._age.TAttr = _bbg.ST_CellTypeStr
	_fa._age.F = _bbg.NewCT_CellFormula()
	_fa._age.F.TAttr = _bbg.ST_CellFormulaTypeArray
	_fa._age.F.Content = s
}

// Column represents a column within a sheet. It's only used for formatting
// purposes, so it's possible to construct a sheet without configuring columns.
type Column struct{ _ggba *_bbg.CT_Col }

func (_fdd Font) SetBold(b bool) {
	if b {
		_fdd._bdbb.B = []*_bbg.CT_BooleanProperty{{}}
	} else {
		_fdd._bdbb.B = nil
	}
}

// RowNumber returns the row number (1-N), or zero if it is unset.
func (_bbdg Row) RowNumber() uint32 {
	if _bbdg._fddd.RAttr != nil {
		return *_bbdg._fddd.RAttr
	}
	return 0
}

// Anchor is the interface implemented by anchors. It's modeled after the most
// common anchor (Two cell variant with a from/to position), but will also be
// used for one-cell anchors.  In that case the only non-noop methods are
// TopLeft/MoveTo/SetColOffset/SetRowOffset.
type Anchor interface {

	// BottomRight returns the CellMaker for the bottom right corner of the
	// anchor.
	BottomRight() CellMarker

	// TopLeft returns the CellMaker for the top left corner of the anchor.
	TopLeft() CellMarker

	// MoveTo repositions the anchor without changing the objects size.
	MoveTo(_agf, _bdge int32)

	// SetWidth sets the width of the anchored object. It is not compatible with
	// SetWidthCells.
	SetWidth(_eea _bg.Distance)

	// SetWidthCells sets the height the anchored object by moving the right
	// hand side. It is not compatible with SetWidth.
	SetWidthCells(_cfa int32)

	// SetHeight sets the height of the anchored object. It is not compatible
	// with SetHeightCells.
	SetHeight(_gg _bg.Distance)

	// SetHeightCells sets the height the anchored object by moving the bottom.
	// It is not compatible with SetHeight.
	SetHeightCells(_eaa int32)

	// SetColOffset sets the column offset of the top-left anchor.
	SetColOffset(_bca _bg.Distance)

	// SetRowOffset sets the row offset of the top-left anchor.
	SetRowOffset(_dfe _bg.Distance)

	// Type returns the type of anchor
	Type() AnchorType
}

// SetCol set the column of the cell marker.
func (_ddc CellMarker) SetCol(col int32) { _ddc._bef.Col = col }

// X returns the inner wrapped XML type.
func (_caa *Workbook) X() *_bbg.Workbook { return _caa._bdff }

// ColOffset returns the offset from the row cell.
func (_dad CellMarker) ColOffset() _bg.Distance {
	if _dad._bef.RowOff.ST_CoordinateUnqualified == nil {
		return 0
	}
	return _bg.Distance(float64(*_dad._bef.ColOff.ST_CoordinateUnqualified) * _bg.EMU)
}

// X returns the inner wrapped XML type.
func (_gdaf MergedCell) X() *_bbg.CT_MergeCell { return _gdaf._gfeg }

// AddNumberedRow adds a row with a given row number.  If you reuse a row number
// the resulting file will fail validation and fail to open in Office programs. Use
// Row instead which creates a new row or returns an existing row.
func (_dffd *Sheet) AddNumberedRow(rowNum uint32) Row {
	_bbgdc := _bbg.NewCT_Row()
	_bbgdc.RAttr = _b.Uint32(rowNum)
	_dffd._ecgc.SheetData.Row = append(_dffd._ecgc.SheetData.Row, _bbgdc)
	_c.Slice(_dffd._ecgc.SheetData.Row, func(_dcdg, _ceae int) bool {
		_fbg := _dffd._ecgc.SheetData.Row[_dcdg].RAttr
		_dbd := _dffd._ecgc.SheetData.Row[_ceae].RAttr
		if _fbg == nil {
			return true
		}
		if _dbd == nil {
			return true
		}
		return *_fbg < *_dbd
	})
	return Row{_dffd._afga, _dffd, _bbgdc}
}

// ClearProtection removes any protections applied to teh sheet.
func (_bgac *Sheet) ClearProtection() { _bgac._ecgc.SheetProtection = nil }

// CellReference returns the cell reference within a sheet that a comment refers
// to (e.g. "A1")
func (_ggeb Comment) CellReference() string { return _ggeb._acb.RefAttr }
func (_ceadf Cell) setLocked(_fff bool) {
	_bbd := _ceadf._age.SAttr
	if _bbd != nil {
		_cbee := _ceadf._cea.StyleSheet.GetCellStyle(*_bbd)
		if _cbee._gcg.Protection == nil {
			_cbee._gcg.Protection = _bbg.NewCT_CellProtection()
		}
		_cbee._gcg.Protection.LockedAttr = &_fff
	}
}

// Close closes the workbook, removing any temporary files that might have been
// created when opening a document.
func (_fdfd *Workbook) Close() error {
	if _fdfd.TmpPath != "" {
		return _ae.RemoveAll(_fdfd.TmpPath)
	}
	return nil
}

// Type returns the type of anchor
func (_afac TwoCellAnchor) Type() AnchorType { return AnchorTypeTwoCell }

// RichText is a container for the rich text within a cell. It's similar to a
// paragaraph for a document, except a cell can only contain one rich text item.
type RichText struct{ _agac *_bbg.CT_Rst }

// AddFormatValue adds a format value (databars require two).
func (_dae DataBarScale) AddFormatValue(t _bbg.ST_CfvoType, val string) {
	_dgf := _bbg.NewCT_Cfvo()
	_dgf.TypeAttr = t
	_dgf.ValAttr = _b.String(val)
	_dae._gdb.Cfvo = append(_dae._gdb.Cfvo, _dgf)
}
func (_af Border) SetRight(style _bbg.ST_BorderStyle, c _fbf.Color) {
	if _af._ef.Right == nil {
		_af._ef.Right = _bbg.NewCT_BorderPr()
	}
	_af._ef.Right.Color = _bbg.NewCT_Color()
	_af._ef.Right.Color.RgbAttr = c.AsRGBAString()
	_af._ef.Right.StyleAttr = style
}

// IsStructureLocked returns whether the workbook structure is locked.
func (_cege WorkbookProtection) IsStructureLocked() bool {
	return _cege._fefe.LockStructureAttr != nil && *_cege._fefe.LockStructureAttr
}
func (_gebcc Font) SetItalic(b bool) {
	if b {
		_gebcc._bdbb.I = []*_bbg.CT_BooleanProperty{{}}
	} else {
		_gebcc._bdbb.I = nil
	}
}
func _fdc(_cab _be.Time) _be.Time {
	_cab = _cab.UTC()
	return _be.Date(_cab.Year(), _cab.Month(), _cab.Day(), _cab.Hour(), _cab.Minute(), _cab.Second(), _cab.Nanosecond(), _be.Local)
}
func _ccd(_fcgb string) bool {
	_fcgb = _ba.Replace(_fcgb, "\u0024", "", -1)
	if _ebgb := _gade.FindStringSubmatch(_ba.ToLower(_fcgb)); len(_ebgb) > 2 {
		_dfda := _ebgb[1]
		_eagf, _bdbd := _dg.Atoi(_ebgb[2])
		if _bdbd != nil {
			return false
		}
		return _eagf <= 1048576 && _dfda <= "\u007a\u007a"
	}
	return false
}

// SetHidden controls the visibility of a column.
func (_dag Column) SetHidden(b bool) {
	if !b {
		_dag._ggba.HiddenAttr = nil
	} else {
		_dag._ggba.HiddenAttr = _b.Bool(true)
	}
}

// IsEmpty checks if the cell style contains nothing.
func (_bbgd CellStyle) IsEmpty() bool {
	return _bbgd._ffgf == nil || _bbgd._gcg == nil || _bbgd._efb == nil || _bbgd._efb.Xf == nil
}

// MoveTo repositions the anchor without changing the objects size.
func (_ggbf TwoCellAnchor) MoveTo(col, row int32) {
	_cacae := _ggbf.TopLeft()
	_bcfd := _ggbf.BottomRight()
	_gbfb := _bcfd.Col() - _cacae.Col()
	_dddgd := _bcfd.Row() - _cacae.Row()
	_cacae.SetCol(col)
	_cacae.SetRow(row)
	_bcfd.SetCol(col + _gbfb)
	_bcfd.SetRow(row + _dddgd)
}

// X returns the inner wrapped XML type.
func (_dfdc SharedStrings) X() *_bbg.Sst { return _dfdc._ddec }

// Rows returns all of the rows in a sheet.
func (_fefac *Sheet) Rows() []Row {
	_aeac := []Row{}
	for _, _fcbe := range _fefac._ecgc.SheetData.Row {
		_aeac = append(_aeac, Row{_fefac._afga, _fefac, _fcbe})
	}
	return _aeac
}
func (_addd Sheet) validateMergedCells() error {
	_efca := map[uint64]struct{}{}
	for _, _gcac := range _addd.MergedCells() {
		_efdb, _beff, _ebf := _df.ParseRangeReference(_gcac.Reference())
		if _ebf != nil {
			return _cg.Errorf("\u0073\u0068e\u0065\u0074\u0020\u006e\u0061m\u0065\u0020\u0027\u0025\u0073'\u0020\u0068\u0061\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006c\u006c\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0025\u0073", _addd.Name(), _gcac.Reference())
		}
		for _daad := _efdb.RowIdx; _daad <= _beff.RowIdx; _daad++ {
			for _ffd := _efdb.ColumnIdx; _ffd <= _beff.ColumnIdx; _ffd++ {
				_bbgb := uint64(_daad)<<32 | uint64(_ffd)
				if _, _fbeb := _efca[_bbgb]; _fbeb {
					return _cg.Errorf("\u0073\u0068\u0065\u0065\u0074\u0020n\u0061\u006d\u0065\u0020\u0027\u0025\u0073\u0027\u0020\u0068\u0061\u0073\u0020\u006f\u0076\u0065\u0072\u006c\u0061\u0070p\u0069\u006e\u0067\u0020\u006d\u0065\u0072\u0067\u0065\u0064\u0020\u0063\u0065\u006cl\u0020r\u0061\u006e\u0067\u0065", _addd.Name())
				}
				_efca[_bbgb] = struct{}{}
			}
		}
	}
	return nil
}

// AddCommentWithStyle adds a new comment styled in a default way
func (_dbg Comments) AddCommentWithStyle(cellRef string, author string, comment string) error {
	_ceeb := _dbg.AddComment(cellRef, author)
	_dddf := _ceeb.AddRun()
	_dddf.SetBold(true)
	_dddf.SetSize(10)
	_dddf.SetColor(_fbf.Black)
	_dddf.SetFont("\u0043a\u006c\u0069\u0062\u0072\u0069")
	_dddf.SetText(author + "\u003a")
	_dddf = _ceeb.AddRun()
	_dddf.SetSize(10)
	_dddf.SetFont("\u0043a\u006c\u0069\u0062\u0072\u0069")
	_dddf.SetColor(_fbf.Black)
	_dddf.SetText("\u000d\u000a" + comment + "\u000d\u000a")
	_fgaf, _aea := _df.ParseCellReference(cellRef)
	if _aea != nil {
		return _aea
	}
	_dbg._fgb._fdbb[0].Shape = append(_dbg._fgb._fdbb[0].Shape, _bb.NewCommentShape(int64(_fgaf.ColumnIdx), int64(_fgaf.RowIdx-1)))
	return nil
}

// ConditionalFormatting controls the formatting styles and rules for a range of
// cells with the same conditional formatting.
type ConditionalFormatting struct {
	_cff *_bbg.CT_ConditionalFormatting
}

// LockWindow controls the locking of the workbook windows.
func (_cgdb WorkbookProtection) LockWindow(b bool) {
	if !b {
		_cgdb._fefe.LockWindowsAttr = nil
	} else {
		_cgdb._fefe.LockWindowsAttr = _b.Bool(true)
	}
}
func (_bagc Sheet) validateRowCellNumbers() error {
	_edeac := map[uint32]struct{}{}
	for _, _gegc := range _bagc._ecgc.SheetData.Row {
		if _gegc.RAttr != nil {
			if _, _bdbc := _edeac[*_gegc.RAttr]; _bdbc {
				return _cg.Errorf("\u0027%\u0073'\u0020\u0072\u0065\u0075\u0073e\u0064\u0020r\u006f\u0077\u0020\u0025\u0064", _bagc.Name(), *_gegc.RAttr)
			}
			_edeac[*_gegc.RAttr] = struct{}{}
		}
		_bgg := map[string]struct{}{}
		for _, _bfe := range _gegc.C {
			if _bfe.RAttr == nil {
				continue
			}
			if _, _faac := _bgg[*_bfe.RAttr]; _faac {
				return _cg.Errorf("\u0027\u0025\u0073\u0027 r\u0065\u0075\u0073\u0065\u0064\u0020\u0063\u0065\u006c\u006c\u0020\u0025\u0073", _bagc.Name(), *_bfe.RAttr)
			}
			_bgg[*_bfe.RAttr] = struct{}{}
		}
	}
	return nil
}

// GetFilename returns the name of file from which workbook was opened with full path to it
func (_cgbd *Workbook) GetFilename() string { return _cgbd._aafb }

// SetDataBar configures the rule as a data bar, removing existing
// configuration.
func (_efc ConditionalFormattingRule) SetDataBar() DataBarScale {
	_efc.clear()
	_efc.SetType(_bbg.ST_CfTypeDataBar)
	_efc._daf.DataBar = _bbg.NewCT_DataBar()
	_aab := DataBarScale{_efc._daf.DataBar}
	_aab.SetShowValue(true)
	_aab.SetMinLength(10)
	_aab.SetMaxLength(90)
	return _aab
}

// SetConditionValue sets the condition value to be used for style applicaton.
func (_ede ConditionalFormattingRule) SetConditionValue(v string) { _ede._daf.Formula = []string{v} }

// SetAuthor sets the author of the comment. If the comment body contains the
// author's name (as is the case with Excel and Comments.AddCommentWithStyle, it
// will not be changed).  This method only changes the metadata author of the
// comment.
func (_cgf Comment) SetAuthor(author string) {
	_cgf._acb.AuthorIdAttr = Comments{_cgf._ecd, _cgf._fffd}.getOrCreateAuthor(author)
}
func (_dff Font) SetName(name string) { _dff._bdbb.Name = []*_bbg.CT_FontName{{ValAttr: name}} }

// BottomRight returns the CellMaker for the bottom right corner of the anchor.
func (_bdcae TwoCellAnchor) BottomRight() CellMarker { return CellMarker{_bdcae._ccec.To} }

// LockStructure controls the locking of the workbook structure.
func (_bdbfb WorkbookProtection) LockStructure(b bool) {
	if !b {
		_bdbfb._fefe.LockStructureAttr = nil
	} else {
		_bdbfb._fefe.LockStructureAttr = _b.Bool(true)
	}
}

// Fills returns a Fills object that can be used to add/create/edit fills.
func (_acfd StyleSheet) Fills() Fills { return Fills{_acfd._fgbg.Fills} }

// PasswordHash returns the password hash for a workbook using the modified
// spreadsheetML password hash that is compatible with Excel.
func PasswordHash(s string) string {
	_bfgbc := uint16(0)
	if len(s) > 0 {
		for _fage := len(s) - 1; _fage >= 0; _fage-- {
			_ccbad := s[_fage]
			_bfgbc = ((_bfgbc >> 14) & 0x01) | ((_bfgbc << 1) & 0x7fff)
			_bfgbc ^= uint16(_ccbad)
		}
		_bfgbc = ((_bfgbc >> 14) & 0x01) | ((_bfgbc << 1) & 0x7fff)
		_bfgbc ^= uint16(len(s))
		_bfgbc ^= (0x8000 | ('N' << 8) | 'K')
	}
	return _cg.Sprintf("\u0025\u0030\u0034\u0058", uint64(_bfgbc))
}

type ConditionalFormattingRule struct{ _daf *_bbg.CT_CfRule }

var _ggcg = _b.RelativeFilename(_b.DocTypeSpreadsheet, _b.OfficeDocumentType, _b.SharedStringsType, 0)

// SetHeight sets the height of the anchored object.
func (_fce OneCellAnchor) SetHeight(h _bg.Distance) { _fce._cdcd.Ext.CyAttr = int64(h / _bg.EMU) }

// GetDrawing return the worksheet drawing and its relationships if exists.
func (_ecea *Sheet) GetDrawing() (*_ag.WsDr, _dcb.Relationships) {
	if _cceg := _ecea._ecgc.Drawing; _cceg != nil {
		_bfac := 0
		for _, _dfde := range _ecea._afga._cfgf {
			if _gfbb := _dfde.Drawing; _gfbb != nil {
				if _dfde == _ecea._ecgc {
					return _ecea._afga._eded[_bfac], _ecea._afga._ccag[_bfac]
				}
				_bfac++
			}
		}
	}
	return nil, _dcb.Relationships{}
}

// SheetCount returns the number of sheets in the workbook.
func (_gacf Workbook) SheetCount() int { return len(_gacf._cfgf) }

// Name returns the sheet name
func (_debe Sheet) Name() string { return _debe._adgb.NameAttr }

// SetFrozen removes any existing sheet views and creates a new single view with
// either the first row, first column or both frozen.
func (_acbe *Sheet) SetFrozen(firstRow, firstCol bool) {
	_acbe._ecgc.SheetViews = nil
	_gddf := _acbe.AddView()
	_gddf.SetState(_bbg.ST_PaneStateFrozen)
	switch {
	case firstRow && firstCol:
		_gddf.SetYSplit(1)
		_gddf.SetXSplit(1)
		_gddf.SetTopLeft("\u0042\u0032")
	case firstRow:
		_gddf.SetYSplit(1)
		_gddf.SetTopLeft("\u0041\u0032")
	case firstCol:
		_gddf.SetXSplit(1)
		_gddf.SetTopLeft("\u0042\u0031")
	}
}

// AddGradientStop adds a color gradient stop.
func (_dea ColorScale) AddGradientStop(color _fbf.Color) {
	_gcge := _bbg.NewCT_Color()
	_gcge.RgbAttr = color.AsRGBAString()
	_dea._gfa.Color = append(_dea._gfa.Color, _gcge)
}

// TwoCellAnchor is an anchor that is attached to a top-left cell with a fixed
// width/height in cells.
type TwoCellAnchor struct{ _ccec *_ag.CT_TwoCellAnchor }

// SetZoom controls the zoom level of the sheet and is measured in percent. The
// default value is 100.
func (_begf SheetView) SetZoom(pct uint32) { _begf._abfc.ZoomScaleAttr = &pct }

// SetSize sets the text size for a rich text run.
func (_gcaa RichTextRun) SetSize(m _bg.Distance) {
	_gcaa.ensureRpr()
	_gcaa._ffae.RPr.Sz = _bbg.NewCT_FontSize()
	_gcaa._ffae.RPr.Sz.ValAttr = float64(m / _bg.Point)
}

// Uses1904Dates returns true if the the workbook uses dates relative to
// 1 Jan 1904. This is uncommon.
func (_dcbfe *Workbook) Uses1904Dates() bool {
	if _dcbfe._bdff.WorkbookPr == nil || _dcbfe._bdff.WorkbookPr.Date1904Attr == nil {
		return false
	}
	return *_dcbfe._bdff.WorkbookPr.Date1904Attr
}

// AddBorder creates a new empty Border style.
func (_cgec StyleSheet) AddBorder() Border {
	_abdg := _bbg.NewCT_Border()
	return Border{_abdg, _cgec._fgbg.Borders}
}

// SetColOffset sets a column offset in absolute distance.
func (_ccg CellMarker) SetColOffset(m _bg.Distance) {
	_ccg._bef.ColOff.ST_CoordinateUnqualified = _b.Int64(int64(m / _bg.EMU))
}

// InitializeDefaults initializes a border to its defaulte empty values.
func (_def Border) InitializeDefaults() {
	_def._ef.Left = _bbg.NewCT_BorderPr()
	_def._ef.Bottom = _bbg.NewCT_BorderPr()
	_def._ef.Right = _bbg.NewCT_BorderPr()
	_def._ef.Top = _bbg.NewCT_BorderPr()
	_def._ef.Diagonal = _bbg.NewCT_BorderPr()
}

type MergedCell struct {
	_febf *Workbook
	_eccd *Sheet
	_gfeg *_bbg.CT_MergeCell
}

// X returns the inner wrapped XML type.
func (_aege Comment) X() *_bbg.CT_Comment { return _aege._acb }

// SetHeight sets the row height in points.
func (_cce Row) SetHeight(d _bg.Distance) {
	_cce._fddd.HtAttr = _b.Float64(float64(d))
	_cce._fddd.CustomHeightAttr = _b.Bool(true)
}

// SetHeightAuto sets the row height to be automatically determined.
func (_deb Row) SetHeightAuto() { _deb._fddd.HtAttr = nil; _deb._fddd.CustomHeightAttr = nil }

// InsertRow inserts a new row into a spreadsheet at a particular row number.  This
// row will now be the row number specified, and any rows after it will be renumbed.
func (_aca *Sheet) InsertRow(rowNum int) Row {
	_ceee := uint32(rowNum)
	for _, _gcbbf := range _aca.Rows() {
		if _gcbbf._fddd.RAttr != nil && *_gcbbf._fddd.RAttr >= _ceee {
			*_gcbbf._fddd.RAttr++
			for _, _deag := range _gcbbf.Cells() {
				_bce, _bbdd := _df.ParseCellReference(_deag.Reference())
				if _bbdd != nil {
					continue
				}
				_bce.RowIdx++
				_deag._age.RAttr = _b.String(_bce.String())
			}
		}
	}
	for _, _fcgd := range _aca.MergedCells() {
		_ggbe, _dfdaa, _cfdc := _df.ParseRangeReference(_fcgd.Reference())
		if _cfdc != nil {
			continue
		}
		if int(_ggbe.RowIdx) >= rowNum {
			_ggbe.RowIdx++
		}
		if int(_dfdaa.RowIdx) >= rowNum {
			_dfdaa.RowIdx++
		}
		_agee := _cg.Sprintf("\u0025\u0073\u003a%\u0073", _ggbe, _dfdaa)
		_fcgd.SetReference(_agee)
	}
	return _aca.AddNumberedRow(_ceee)
}

// Open opens and reads a workbook from a file (.xlsx).
func Open(filename string) (*Workbook, error) {
	_eeed, _bda := _bd.Open(filename)
	if _bda != nil {
		return nil, _cg.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _bda)
	}
	defer _eeed.Close()
	_edfe, _bda := _bd.Stat(filename)
	if _bda != nil {
		return nil, _cg.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _bda)
	}
	_cgcc, _bda := Read(_eeed, _edfe.Size())
	if _bda != nil {
		return nil, _bda
	}
	_gcgfa, _ := _g.Abs(_g.Dir(filename))
	_cgcc._aafb = _g.Join(_gcgfa, filename)
	return _cgcc, nil
}

// AddCell adds a cell to a spreadsheet.
func (_adb Row) AddCell() Cell {
	_edd := uint32(len(_adb._fddd.C))
	var _fgcde *string
	if _edd > 0 {
		_debf := _b.Stringf("\u0025\u0073\u0025\u0064", _df.IndexToColumn(_edd-1), _adb.RowNumber())
		if _adb._fddd.C[_edd-1].RAttr != nil && *_adb._fddd.C[_edd-1].RAttr == *_debf {
			_fgcde = _b.Stringf("\u0025\u0073\u0025\u0064", _df.IndexToColumn(_edd), _adb.RowNumber())
		}
	}
	_cgdf := _bbg.NewCT_Cell()
	_adb._fddd.C = append(_adb._fddd.C, _cgdf)
	if _fgcde == nil {
		_dgecf := uint32(0)
		for _, _ffbe := range _adb._fddd.C {
			if _ffbe.RAttr != nil {
				_ecg, _ := _df.ParseCellReference(*_ffbe.RAttr)
				if _ecg.ColumnIdx >= _dgecf {
					_dgecf = _ecg.ColumnIdx + 1
				}
			}
		}
		_fgcde = _b.Stringf("\u0025\u0073\u0025\u0064", _df.IndexToColumn(_dgecf), _adb.RowNumber())
	}
	_cgdf.RAttr = _fgcde
	return Cell{_adb._fadd, _adb._bafc, _adb._fddd, _cgdf}
}

type WorkbookProtection struct{ _fefe *_bbg.CT_WorkbookProtection }

// GetLabelPrefix returns label prefix which depends on the cell's horizontal alignment.
func (_gfaf *evalContext) GetLabelPrefix(cellRef string) string {
	return _gfaf._befg.Cell(cellRef).getLabelPrefix()
}
func (_afec Fills) appendFill() Fill {
	_cbcg := _bbg.NewCT_Fill()
	_afec._egaa.Fill = append(_afec._egaa.Fill, _cbcg)
	_afec._egaa.CountAttr = _b.Uint32(uint32(len(_afec._egaa.Fill)))
	return Fill{_cbcg, _afec._egaa}
}

// ClearCachedFormulaResults clears any computed formula values that are stored
// in the sheet. This may be required if you modify cells that are used as a
// formula input to force the formulas to be recomputed the next time the sheet
// is opened in Excel.
func (_bagd *Sheet) ClearCachedFormulaResults() {
	for _, _bcabg := range _bagd.Rows() {
		for _, _dddcg := range _bcabg.Cells() {
			if _dddcg.X().F != nil {
				_dddcg.X().V = nil
			}
		}
	}
}

// AddNamedCell adds a new named cell to a row and returns it. You should
// normally prefer Cell() as it will return the existing cell if the cell
// already exists, while AddNamedCell will duplicate the cell creating an
// invaild spreadsheet.
func (_adbb Row) AddNamedCell(col string) Cell {
	_adg := _bbg.NewCT_Cell()
	_adg.RAttr = _b.Stringf("\u0025\u0073\u0025\u0064", col, _adbb.RowNumber())
	_ffeb := -1
	_fddf := _df.ColumnToIndex(col)
	for _gffg, _bfbbf := range _adbb._fddd.C {
		_ddgd, _cae := _df.ParseCellReference(*_bfbbf.RAttr)
		if _cae != nil {
			return Cell{}
		}
		if _fddf < _ddgd.ColumnIdx {
			_ffeb = _gffg
			break
		}
	}
	if _ffeb == -1 {
		_adbb._fddd.C = append(_adbb._fddd.C, _adg)
	} else {
		_adbb._fddd.C = append(_adbb._fddd.C[:_ffeb], append([]*_bbg.CT_Cell{_adg}, _adbb._fddd.C[_ffeb:]...)...)
	}
	return Cell{_adbb._fadd, _adbb._bafc, _adbb._fddd, _adg}
}

// SetBorder is a helper function for creating borders across multiple cells. In
// the OOXML spreadsheet format, a border applies to a single cell.  To draw a
// 'boxed' border around multiple cells, you need to apply different styles to
// the cells on the top,left,right,bottom and four corners.  This function
// breaks apart a single border into its components and applies it to cells as
// needed to give the effect of a border applying to multiple cells.
func (_bfgc *Sheet) SetBorder(cellRange string, border Border) error {
	_edge, _gege, _cagef := _df.ParseRangeReference(cellRange)
	if _cagef != nil {
		return _cagef
	}
	_cfbe := _bfgc._afga.StyleSheet.AddCellStyle()
	_fbed := _bfgc._afga.StyleSheet.AddBorder()
	_cfbe.SetBorder(_fbed)
	_fbed._ef.Top = border._ef.Top
	_fbed._ef.Left = border._ef.Left
	_cfeg := _bfgc._afga.StyleSheet.AddCellStyle()
	_efdc := _bfgc._afga.StyleSheet.AddBorder()
	_cfeg.SetBorder(_efdc)
	_efdc._ef.Top = border._ef.Top
	_efdc._ef.Right = border._ef.Right
	_fafa := _bfgc._afga.StyleSheet.AddCellStyle()
	_abad := _bfgc._afga.StyleSheet.AddBorder()
	_fafa.SetBorder(_abad)
	_abad._ef.Top = border._ef.Top
	_bgggg := _bfgc._afga.StyleSheet.AddCellStyle()
	_bdggb := _bfgc._afga.StyleSheet.AddBorder()
	_bgggg.SetBorder(_bdggb)
	_bdggb._ef.Left = border._ef.Left
	_cdcce := _bfgc._afga.StyleSheet.AddCellStyle()
	_gfcd := _bfgc._afga.StyleSheet.AddBorder()
	_cdcce.SetBorder(_gfcd)
	_gfcd._ef.Right = border._ef.Right
	_fbced := _bfgc._afga.StyleSheet.AddCellStyle()
	_ceadff := _bfgc._afga.StyleSheet.AddBorder()
	_fbced.SetBorder(_ceadff)
	_ceadff._ef.Bottom = border._ef.Bottom
	_dead := _bfgc._afga.StyleSheet.AddCellStyle()
	_fced := _bfgc._afga.StyleSheet.AddBorder()
	_dead.SetBorder(_fced)
	_fced._ef.Bottom = border._ef.Bottom
	_fced._ef.Left = border._ef.Left
	_beeb := _bfgc._afga.StyleSheet.AddCellStyle()
	_dcgb := _bfgc._afga.StyleSheet.AddBorder()
	_beeb.SetBorder(_dcgb)
	_dcgb._ef.Bottom = border._ef.Bottom
	_dcgb._ef.Right = border._ef.Right
	_affc := _edge.RowIdx
	_aebd := _edge.ColumnIdx
	_feca := _gege.RowIdx
	_efaa := _gege.ColumnIdx
	for _aaf := _affc; _aaf <= _feca; _aaf++ {
		for _ageb := _aebd; _ageb <= _efaa; _ageb++ {
			_egbc := _cg.Sprintf("\u0025\u0073\u0025\u0064", _df.IndexToColumn(_ageb), _aaf)
			switch {
			case _aaf == _affc && _ageb == _aebd:
				_bfgc.Cell(_egbc).SetStyle(_cfbe)
			case _aaf == _affc && _ageb == _efaa:
				_bfgc.Cell(_egbc).SetStyle(_cfeg)
			case _aaf == _feca && _ageb == _aebd:
				_bfgc.Cell(_egbc).SetStyle(_dead)
			case _aaf == _feca && _ageb == _efaa:
				_bfgc.Cell(_egbc).SetStyle(_beeb)
			case _aaf == _affc:
				_bfgc.Cell(_egbc).SetStyle(_fafa)
			case _aaf == _feca:
				_bfgc.Cell(_egbc).SetStyle(_fbced)
			case _ageb == _aebd:
				_bfgc.Cell(_egbc).SetStyle(_bgggg)
			case _ageb == _efaa:
				_bfgc.Cell(_egbc).SetStyle(_cdcce)
			}
		}
	}
	return nil
}

// SetType sets the type of the rule.
func (_ageg ConditionalFormattingRule) SetType(t _bbg.ST_CfType) { _ageg._daf.TypeAttr = t }

// SetWidth sets the width of the anchored object.
func (_db AbsoluteAnchor) SetWidth(w _bg.Distance) { _db._da.Ext.CxAttr = int64(w / _bg.EMU) }

// IsSheetLocked returns whether the sheet is locked.
func (_fafbe SheetProtection) IsSheetLocked() bool {
	return _fafbe._ggbb.SheetAttr != nil && *_fafbe._ggbb.SheetAttr
}

// Cell is a single cell within a sheet.
type Cell struct {
	_cea *Workbook
	_acf *Sheet
	_geb *_bbg.CT_Row
	_age *_bbg.CT_Cell
}

// SetIcons configures the rule as an icon scale, removing existing
// configuration.
func (_gfdd ConditionalFormattingRule) SetIcons() IconScale {
	_gfdd.clear()
	_gfdd.SetType(_bbg.ST_CfTypeIconSet)
	_gfdd._daf.IconSet = _bbg.NewCT_IconSet()
	_ded := IconScale{_gfdd._daf.IconSet}
	_ded.SetIcons(_bbg.ST_IconSetType3TrafficLights1)
	return _ded
}
func (_cdcg StyleSheet) appendFont() Font {
	_bgab := _bbg.NewCT_Font()
	_cdcg._fgbg.Fonts.Font = append(_cdcg._fgbg.Fonts.Font, _bgab)
	_cdcg._fgbg.Fonts.CountAttr = _b.Uint32(uint32(len(_cdcg._fgbg.Fonts.Font)))
	return Font{_bgab, _cdcg._fgbg}
}

// SetInlineString adds a string inline instead of in the shared strings table.
func (_cba Cell) SetInlineString(s string) {
	_cba.clearValue()
	_cba._age.Is = _bbg.NewCT_Rst()
	_cba._age.Is.T = _b.String(s)
	_cba._age.TAttr = _bbg.ST_CellTypeInlineStr
}

// Text returns text from the workbook as one string separated with line breaks.
func (_cage *WorkbookText) Text() string {
	_ddea := _gb.NewBuffer([]byte{})
	for _, _beeg := range _cage.Sheets {
		_ddea.WriteString(_beeg.Text())
	}
	return _ddea.String()
}

type Fill struct {
	_aeag *_bbg.CT_Fill
	_fgc  *_bbg.CT_Fills
}

// AddImage adds an image to the workbook package, returning a reference that
// can be used to add the image to a drawing.
func (_ggee *Workbook) AddImage(i _dcb.Image) (_dcb.ImageRef, error) {
	_dabd := _dcb.MakeImageRef(i, &_ggee.DocBase, _ggee._fdae)
	if i.Data == nil && i.Path == "" {
		return _dabd, _fb.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _dabd, _fb.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _dabd, _fb.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	if i.Path != "" {
		_ffgfc := _ae.Add(i.Path)
		if _ffgfc != nil {
			return _dabd, _ffgfc
		}
	}
	_ggee.Images = append(_ggee.Images, _dabd)
	return _dabd, nil
}

// Workbook returns sheet's parent workbook.
func (_ecbd *Sheet) Workbook() *Workbook { return _ecbd._afga }
func (_bcac Cell) clearValue() {
	_bcac._age.F = nil
	_bcac._age.Is = nil
	_bcac._age.V = nil
	_bcac._age.TAttr = _bbg.ST_CellTypeUnset
}

// SetFgColor sets the *fill* foreground color.  As an example, the solid pattern foreground color becomes the
// background color of the cell when applied.
func (_cggb PatternFill) SetFgColor(c _fbf.Color) {
	_cggb._aad.FgColor = _bbg.NewCT_Color()
	_cggb._aad.FgColor.RgbAttr = c.AsRGBAString()
}

// X returns the inner wrapped XML type.
func (_abce IconScale) X() *_bbg.CT_IconSet { return _abce._bfdc }

// SetPasswordHash sets the password hash to the input.
func (_abfe SheetProtection) SetPasswordHash(pwHash string) {
	_abfe._ggbb.PasswordAttr = _b.String(pwHash)
}

// SetWidthCells is a no-op.
func (_dcd OneCellAnchor) SetWidthCells(int32) {}

// GetBorder gets a Border from a cell style.
func (_gccab CellStyle) GetBorder() *_bbg.CT_Border {
	if _gfc := _gccab._gcg.BorderIdAttr; _gfc != nil {
		_fafdf := _gccab._ffgf.StyleSheet.Borders()
		if int(*_gfc) < len(_fafdf) {
			return _fafdf[int(*_gfc)].X()
		}
	}
	return nil
}

// NewStyleSheet constructs a new default stylesheet.
func NewStyleSheet(wb *Workbook) StyleSheet {
	_eddf := _bbg.NewStyleSheet()
	_eddf.CellStyleXfs = _bbg.NewCT_CellStyleXfs()
	_eddf.CellXfs = _bbg.NewCT_CellXfs()
	_eddf.CellStyles = _bbg.NewCT_CellStyles()
	_ddb := _bbg.NewCT_CellStyle()
	_ddb.NameAttr = _b.String("\u004e\u006f\u0072\u006d\u0061\u006c")
	_ddb.XfIdAttr = 0
	_ddb.BuiltinIdAttr = _b.Uint32(0)
	_eddf.CellStyles.CellStyle = append(_eddf.CellStyles.CellStyle, _ddb)
	_eddf.CellStyles.CountAttr = _b.Uint32(uint32(len(_eddf.CellStyles.CellStyle)))
	_aegg := _bbg.NewCT_Xf()
	_aegg.NumFmtIdAttr = _b.Uint32(0)
	_aegg.FontIdAttr = _b.Uint32(0)
	_aegg.FillIdAttr = _b.Uint32(0)
	_aegg.BorderIdAttr = _b.Uint32(0)
	_eddf.CellStyleXfs.Xf = append(_eddf.CellStyleXfs.Xf, _aegg)
	_eddf.CellStyleXfs.CountAttr = _b.Uint32(uint32(len(_eddf.CellStyleXfs.Xf)))
	_dgaa := NewFills()
	_eddf.Fills = _dgaa.X()
	_cfga := _dgaa.appendFill().SetPatternFill()
	_cfga.SetPattern(_bbg.ST_PatternTypeNone)
	_cfga = _dgaa.appendFill().SetPatternFill()
	_cfga.SetPattern(_bbg.ST_PatternTypeGray125)
	_eddf.Fonts = _bbg.NewCT_Fonts()
	_eddf.Borders = _bbg.NewCT_Borders()
	_fgfa := StyleSheet{wb, _eddf}
	_fgfa.appendBorder().InitializeDefaults()
	_fgdba := _fgfa.appendFont()
	_fgdba.SetName("\u0043a\u006c\u0069\u0062\u0072\u0069")
	_fgdba.SetSize(11)
	_faacc := _bbg.NewCT_Xf()
	*_faacc = *_aegg
	_faacc.XfIdAttr = _b.Uint32(0)
	_eddf.CellXfs.Xf = append(_eddf.CellXfs.Xf, _faacc)
	_eddf.CellXfs.CountAttr = _b.Uint32(uint32(len(_eddf.CellXfs.Xf)))
	return _fgfa
}

// TopLeft returns the top-left corner of the anchored object.
func (_fcc OneCellAnchor) TopLeft() CellMarker { return CellMarker{_fcc._cdcd.From} }

// SetWidth sets the width of the anchored object.
func (_eafda OneCellAnchor) SetWidth(w _bg.Distance) { _eafda._cdcd.Ext.CxAttr = int64(w / _bg.EMU) }

// Clear clears the cell's value and type.
func (_aeg Cell) Clear() { _aeg.clearValue(); _aeg._age.TAttr = _bbg.ST_CellTypeUnset }

// RichTextRun is a segment of text within a cell that is directly formatted.
type RichTextRun struct{ _ffae *_bbg.CT_RElt }

// GetString retrieves a string from the shared strings table by index.
func (_eedg SharedStrings) GetString(id int) (string, error) {
	if id < 0 {
		return "", _cg.Errorf("\u0069\u006eva\u006c\u0069\u0064 \u0073\u0074\u0072\u0069ng \u0069nd\u0065\u0078\u0020\u0025\u0064\u002c\u0020mu\u0073\u0074\u0020\u0062\u0065\u0020\u003e \u0030", id)
	}
	if id > len(_eedg._ddec.Si)-1 {
		return "", _cg.Errorf("\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0025\u0064\u002c\u0020\u0074\u0061b\u006c\u0065\u0020\u006f\u006e\u006c\u0079\u0020\u0068\u0061\u0073\u0020\u0025\u0064 \u0076a\u006c\u0075\u0065\u0073", id, len(_eedg._ddec.Si))
	}
	_fbdd := _eedg._ddec.Si[id]
	if _fbdd.T != nil {
		return *_fbdd.T, nil
	}
	_dgb := ""
	for _, _cbgf := range _fbdd.R {
		if _cbgf.T != "" {
			_dgb += _cbgf.T
		}
	}
	return _dgb, nil
}

// SetBool sets the cell type to boolean and the value to the given boolean
// value.
func (_dba Cell) SetBool(v bool) {
	_dba.clearValue()
	_dba._age.V = _b.String(_dg.Itoa(_fca(v)))
	_dba._age.TAttr = _bbg.ST_CellTypeB
}

// LockSheet controls the locking of the sheet.
func (_ffga SheetProtection) LockSheet(b bool) {
	if !b {
		_ffga._ggbb.SheetAttr = nil
	} else {
		_ffga._ggbb.SheetAttr = _b.Bool(true)
	}
}
func (_cfaa DifferentialStyle) Fill() Fill {
	if _cfaa._fdga.Fill == nil {
		_cfaa._fdga.Fill = _bbg.NewCT_Fill()
	}
	return Fill{_cfaa._fdga.Fill, nil}
}

// AddImage adds an image with a paricular anchor type, returning an anchor to
// allow adusting the image size/position.
func (_bffd Drawing) AddImage(img _dcb.ImageRef, at AnchorType) Anchor {
	_cgcb := 0
	for _gda, _adf := range _bffd._eafd.Images {
		if _adf == img {
			_cgcb = _gda + 1
			break
		}
	}
	var _bbdb string
	for _bfde, _cfc := range _bffd._eafd._eded {
		if _cfc == _bffd._abd {
			_cge := _cg.Sprintf("\u002e\u002e\u002f\u006ded\u0069\u0061\u002f\u0069\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", _cgcb, img.Format())
			_fbce := _bffd._eafd._ccag[_bfde].AddRelationship(_cge, _b.ImageType)
			_bbdb = _fbce.ID()
			break
		}
	}
	var _bffa Anchor
	var _deda *_ag.CT_Picture
	switch at {
	case AnchorTypeAbsolute:
		_fffc := _egfef()
		_bffd._abd.EG_Anchor = append(_bffd._abd.EG_Anchor, &_ag.EG_Anchor{AbsoluteAnchor: _fffc})
		_fffc.Choice = &_ag.EG_ObjectChoicesChoice{}
		_fffc.Choice.Pic = _ag.NewCT_Picture()
		_fffc.Pos.XAttr.ST_CoordinateUnqualified = _b.Int64(0)
		_fffc.Pos.YAttr.ST_CoordinateUnqualified = _b.Int64(0)
		_deda = _fffc.Choice.Pic
		_bffa = AbsoluteAnchor{_fffc}
	case AnchorTypeOneCell:
		_ebd := _adec()
		_bffd._abd.EG_Anchor = append(_bffd._abd.EG_Anchor, &_ag.EG_Anchor{OneCellAnchor: _ebd})
		_ebd.Choice = &_ag.EG_ObjectChoicesChoice{}
		_ebd.Choice.Pic = _ag.NewCT_Picture()
		_deda = _ebd.Choice.Pic
		_bffa = OneCellAnchor{_ebd}
	case AnchorTypeTwoCell:
		_fcfg := _gaag()
		_bffd._abd.EG_Anchor = append(_bffd._abd.EG_Anchor, &_ag.EG_Anchor{TwoCellAnchor: _fcfg})
		_fcfg.Choice = &_ag.EG_ObjectChoicesChoice{}
		_fcfg.Choice.Pic = _ag.NewCT_Picture()
		_deda = _fcfg.Choice.Pic
		_bffa = TwoCellAnchor{_fcfg}
	}
	_deda.NvPicPr.CNvPr.IdAttr = uint32(len(_bffd._abd.EG_Anchor))
	_deda.NvPicPr.CNvPr.NameAttr = "\u0049\u006d\u0061g\u0065"
	_deda.BlipFill.Blip = _acg.NewCT_Blip()
	_deda.BlipFill.Blip.EmbedAttr = _b.String(_bbdb)
	_deda.BlipFill.Stretch = _acg.NewCT_StretchInfoProperties()
	_deda.SpPr = _acg.NewCT_ShapeProperties()
	_deda.SpPr.Xfrm = _acg.NewCT_Transform2D()
	_deda.SpPr.Xfrm.Off = _acg.NewCT_Point2D()
	_deda.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _b.Int64(0)
	_deda.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _b.Int64(0)
	_deda.SpPr.Xfrm.Ext = _acg.NewCT_PositiveSize2D()
	_deda.SpPr.Xfrm.Ext.CxAttr = int64(float64(img.Size().X*_bg.Pixel72) / _bg.EMU)
	_deda.SpPr.Xfrm.Ext.CyAttr = int64(float64(img.Size().Y*_bg.Pixel72) / _bg.EMU)
	_deda.SpPr.PrstGeom = _acg.NewCT_PresetGeometry2D()
	_deda.SpPr.PrstGeom.PrstAttr = _acg.ST_ShapeTypeRect
	_deda.SpPr.Ln = _acg.NewCT_LineProperties()
	_deda.SpPr.Ln.NoFill = _acg.NewCT_NoFillProperties()
	return _bffa
}
func (_dbge *Sheet) slideCellsLeft(_gbbgcc []*_bbg.CT_Cell) []*_bbg.CT_Cell {
	for _, _dccg := range _gbbgcc {
		_baacb, _dbac := _df.ParseCellReference(*_dccg.RAttr)
		if _dbac != nil {
			return _gbbgcc
		}
		_caca := _baacb.ColumnIdx - 1
		_baeg := _df.IndexToColumn(_caca) + _cg.Sprintf("\u0025\u0064", _baacb.RowIdx)
		_dccg.RAttr = &_baeg
	}
	return _gbbgcc
}

// Sheets returns the sheets from the workbook.
func (_daeg *Workbook) Sheets() []Sheet {
	_ecdgb := []Sheet{}
	for _fcfgc, _ffbbf := range _daeg._cfgf {
		_eceaa := _daeg._bdff.Sheets.Sheet[_fcfgc]
		_cfdf := Sheet{_daeg, _eceaa, _ffbbf}
		_ecdgb = append(_ecdgb, _cfdf)
	}
	return _ecdgb
}

// SetRotation configures the cell to be rotated.
func (_gab CellStyle) SetRotation(deg uint8) {
	if _gab._gcg.Alignment == nil {
		_gab._gcg.Alignment = _bbg.NewCT_CellAlignment()
	}
	_gab._gcg.ApplyAlignmentAttr = _b.Bool(true)
	_gab._gcg.Alignment.TextRotationAttr = _b.Uint8(deg)
}

// GetChartByTargetId returns the array of workbook crt.ChartSpace.
func (_fedgc *Workbook) GetChartByTargetId(targetAttr string) *_fe.ChartSpace {
	return _fedgc._aabe[targetAttr]
}

// Index returns the index of the differential style.
func (_cgba DifferentialStyle) Index() uint32 {
	for _efg, _dcad := range _cgba._babf.Dxf {
		if _cgba._fdga == _dcad {
			return uint32(_efg)
		}
	}
	return 0
}

// SetShowValue controls if the cell value is displayed.
func (_aag DataBarScale) SetShowValue(b bool) { _aag._gdb.ShowValueAttr = _b.Bool(b) }

// Col returns the column of the cell marker.
func (_eba CellMarker) Col() int32 { return _eba._bef.Col }

// SetTime sets the cell value to a date. It's stored as the number of days past
// th sheet epoch. When we support v5 strict, we can store an ISO 8601 date
// string directly, however that's not allowed with v5 transitional  (even
// though it works in Excel).
func (_gef Cell) SetTime(d _be.Time) {
	_gef.clearValue()
	d = _dcc(d)
	_ffee := _gef._cea.Epoch()
	if d.Before(_ffee) {
		_ca.Log.Debug("t\u0069\u006d\u0065\u0073\u0020\u0062e\u0066\u006f\u0072\u0065\u0020\u00319\u0030\u0030\u0020\u0061\u0072\u0065\u0020n\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064")
		return
	}
	_eeg := d.Sub(_ffee)
	_bac := new(_gbg.Float)
	_fdf := new(_gbg.Float)
	_fdf.SetPrec(128)
	_fdf.SetUint64(uint64(_eeg))
	_ggf := new(_gbg.Float)
	_ggf.SetUint64(24 * 60 * 60 * 1e9)
	_bac.Quo(_fdf, _ggf)
	_gef._age.V = _b.String(_bac.Text('g', 20))
}
func (_fefa Cell) getLabelPrefix() string {
	if _fefa._age.SAttr == nil {
		return ""
	}
	_aed := *_fefa._age.SAttr
	_fc := _fefa._cea.StyleSheet.GetCellStyle(_aed)
	switch _fc._gcg.Alignment.HorizontalAttr {
	case _bbg.ST_HorizontalAlignmentLeft:
		return "\u0027"
	case _bbg.ST_HorizontalAlignmentRight:
		return "\u0022"
	case _bbg.ST_HorizontalAlignmentCenter:
		return "\u005e"
	case _bbg.ST_HorizontalAlignmentFill:
		return "\u005c"
	default:
		return ""
	}
}

// SetDrawing sets the worksheet drawing.  A worksheet can have a reference to a
// single drawing, but the drawing can have many charts.
func (_cddc *Sheet) SetDrawing(d Drawing) {
	var _dbae _dcb.Relationships
	for _abdd, _cccb := range _cddc._afga._cfgf {
		if _cccb == _cddc._ecgc {
			_dbae = _cddc._afga._fgdcb[_abdd]
			break
		}
	}
	var _feff string
	for _fdddd, _gedd := range d._eafd._eded {
		if _gedd == d._abd {
			_cegf := _dbae.AddAutoRelationship(_b.DocTypeSpreadsheet, _b.WorksheetType, _fdddd+1, _b.DrawingType)
			_feff = _cegf.ID()
			break
		}
	}
	_cddc._ecgc.Drawing = _bbg.NewCT_Drawing()
	_cddc._ecgc.Drawing.IdAttr = _feff
}

// Comments returns the list of comments for this sheet
func (_ffgb Comments) Comments() []Comment {
	_bccb := []Comment{}
	for _, _abb := range _ffgb._fcg.CommentList.Comment {
		_bccb = append(_bccb, Comment{_ffgb._fgb, _abb, _ffgb._fcg})
	}
	return _bccb
}

// X returns the inner wrapped XML type.
func (_bff Drawing) X() *_ag.WsDr { return _bff._abd }

// RemoveColumn removes column from the sheet and moves all columns to the right of the removed column one step left.
func (_badgf *Sheet) RemoveColumn(column string) error {
	_cdfa, _aafd := _badgf.getAllCellsInFormulaArraysForColumn()
	if _aafd != nil {
		return _aafd
	}
	_dfcab := _df.ColumnToIndex(column)
	for _, _ebeb := range _badgf.Rows() {
		_cfcc := _cg.Sprintf("\u0025\u0073\u0025\u0064", column, *_ebeb.X().RAttr)
		if _, _fdda := _cdfa[_cfcc]; _fdda {
			return nil
		}
	}
	for _, _dgcb := range _badgf.Rows() {
		_aafa := _dgcb._fddd.C
		for _bceb, _debee := range _aafa {
			_acbf, _egcb := _df.ParseCellReference(*_debee.RAttr)
			if _egcb != nil {
				return _egcb
			}
			if _acbf.ColumnIdx == _dfcab {
				_dgcb._fddd.C = append(_aafa[:_bceb], _badgf.slideCellsLeft(_aafa[_bceb+1:])...)
				break
			} else if _acbf.ColumnIdx > _dfcab {
				_dgcb._fddd.C = append(_aafa[:_bceb], _badgf.slideCellsLeft(_aafa[_bceb:])...)
				break
			}
		}
	}
	_aafd = _badgf.updateAfterRemove(_dfcab, _gd.UpdateActionRemoveColumn)
	if _aafd != nil {
		return _aafd
	}
	_aafd = _badgf.removeColumnFromNamedRanges(_dfcab)
	if _aafd != nil {
		return _aafd
	}
	_aafd = _badgf.removeColumnFromMergedCells(_dfcab)
	if _aafd != nil {
		return _aafd
	}
	for _, _cffd := range _badgf._afga.Sheets() {
		_cffd.RecalculateFormulas()
	}
	return nil
}

// ClearNumberFormat removes any number formatting from the style.
func (_gcb CellStyle) ClearNumberFormat() {
	_gcb._gcg.NumFmtIdAttr = nil
	_gcb._gcg.ApplyNumberFormatAttr = nil
}

// Protection controls the protection on an individual sheet.
func (_efec *Sheet) Protection() SheetProtection {
	if _efec._ecgc.SheetProtection == nil {
		_efec._ecgc.SheetProtection = _bbg.NewCT_SheetProtection()
	}
	return SheetProtection{_efec._ecgc.SheetProtection}
}

// CellsWithEmpty returns a slice of cells including empty ones from the first column to the last one used in the sheet.
// The cells can be manipulated, but appending to the slice will have no effect.
func (_bece Row) CellsWithEmpty(lastColIdx uint32) []Cell {
	_dggc := []Cell{}
	for _ace := uint32(0); _ace <= lastColIdx; _ace++ {
		_abff := _bece.Cell(_df.IndexToColumn(_ace))
		_dggc = append(_dggc, _abff)
	}
	return _dggc
}

// SetNumberWithStyle sets a number and applies a standard format to the cell.
func (_gfd Cell) SetNumberWithStyle(v float64, f StandardFormat) {
	_gfd.SetNumber(v)
	_gfd.SetStyle(_gfd._cea.StyleSheet.GetOrCreateStandardNumberFormat(f))
}

// SetStyleIndex directly sets a style index to the cell.  This should only be
// called with an index retrieved from CellStyle.Index()
func (_bgf Cell) SetStyleIndex(idx uint32) { _bgf._age.SAttr = _b.Uint32(idx) }

// GetWidth returns a worksheet's column width.
func (_afe *evalContext) GetWidth(colIdx int) float64 {
	colIdx++
	for _, _ddf := range _afe._befg.X().Cols[0].Col {
		if int(_ddf.MinAttr) <= colIdx && colIdx <= int(_ddf.MaxAttr) {
			return float64(int(*_ddf.WidthAttr))
		}
	}
	return 0
}

// SetShowRuler controls the visibility of the ruler
func (_gaba SheetView) SetShowRuler(b bool) {
	if !b {
		_gaba._abfc.ShowRulerAttr = _b.Bool(false)
	} else {
		_gaba._abfc.ShowRulerAttr = nil
	}
}

// AddDataValidation adds a data validation rule to a sheet.
func (_ccf *Sheet) AddDataValidation() DataValidation {
	if _ccf._ecgc.DataValidations == nil {
		_ccf._ecgc.DataValidations = _bbg.NewCT_DataValidations()
	}
	_bddb := _bbg.NewCT_DataValidation()
	_bddb.ShowErrorMessageAttr = _b.Bool(true)
	_ccf._ecgc.DataValidations.DataValidation = append(_ccf._ecgc.DataValidations.DataValidation, _bddb)
	_ccf._ecgc.DataValidations.CountAttr = _b.Uint32(uint32(len(_ccf._ecgc.DataValidations.DataValidation)))
	return DataValidation{_bddb}
}

// X returns the inner wrapped XML type.
func (_gbgc CellMarker) X() *_ag.CT_Marker { return _gbgc._bef }
func (_bcdg SheetView) ensurePane() {
	if _bcdg._abfc.Pane == nil {
		_bcdg._abfc.Pane = _bbg.NewCT_Pane()
		_bcdg._abfc.Pane.ActivePaneAttr = _bbg.ST_PaneBottomLeft
	}
}

// SetHyperlink sets a hyperlink on a cell.
func (_fafd Cell) SetHyperlink(hl _dcb.Hyperlink) {
	_eee := _fafd._acf._ecgc
	if _eee.Hyperlinks == nil {
		_eee.Hyperlinks = _bbg.NewCT_Hyperlinks()
	}
	_aa := _dcb.Relationship(hl)
	_ggb := _bbg.NewCT_Hyperlink()
	_ggb.RefAttr = _fafd.Reference()
	_ggb.IdAttr = _b.String(_aa.ID())
	_eee.Hyperlinks.Hyperlink = append(_eee.Hyperlinks.Hyperlink, _ggb)
}

// SheetView is a view of a sheet. There is typically one per sheet, though more
// are supported.
type SheetView struct{ _abfc *_bbg.CT_SheetView }

// InitialView returns the first defined sheet view. If there are no views, one
// is created and returned.
func (_bbbd *Sheet) InitialView() SheetView {
	if _bbbd._ecgc.SheetViews == nil || len(_bbbd._ecgc.SheetViews.SheetView) == 0 {
		return _bbbd.AddView()
	}
	return SheetView{_bbbd._ecgc.SheetViews.SheetView[0]}
}

// AddDrawing adds a drawing to a workbook.  However the drawing is not actually
// displayed or used until it's set on a sheet.
func (_daca *Workbook) AddDrawing() Drawing {
	_geba := _ag.NewWsDr()
	_daca._eded = append(_daca._eded, _geba)
	_dcfed := _b.AbsoluteFilename(_b.DocTypeSpreadsheet, _b.DrawingType, len(_daca._eded))
	_daca.ContentTypes.AddOverride(_dcfed, _b.DrawingContentType)
	_daca._ccag = append(_daca._ccag, _dcb.NewRelationships())
	return Drawing{_daca, _geba}
}
func (_ada StyleSheet) appendBorder() Border {
	_badba := _bbg.NewCT_Border()
	_ada._fgbg.Borders.Border = append(_ada._fgbg.Borders.Border, _badba)
	_ada._fgbg.Borders.CountAttr = _b.Uint32(uint32(len(_ada._fgbg.Borders.Border)))
	return Border{_badba, _ada._fgbg.Borders}
}
func (_bec *evalContext) Cell(ref string, ev _eg.Evaluator) _eg.Result {
	if !_ccd(ref) {
		return _eg.MakeErrorResultType(_eg.ErrorTypeName, "")
	}
	_dgd := _bec._befg.Name() + "\u0021" + ref
	if _egfe, _gcba := ev.GetFromCache(_dgd); _gcba {
		return _egfe
	}
	_ccb, _dddb := _df.ParseCellReference(ref)
	if _dddb != nil {
		return _eg.MakeErrorResult(_cg.Sprintf("e\u0072r\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", ref, _dddb))
	}
	if _bec._cdec != 0 && !_ccb.AbsoluteColumn {
		_ccb.ColumnIdx += _bec._cdec
		_ccb.Column = _df.IndexToColumn(_ccb.ColumnIdx)
	}
	if _bec._cbg != 0 && !_ccb.AbsoluteRow {
		_ccb.RowIdx += _bec._cbg
	}
	_aeab := _bec._befg.Cell(_ccb.String())
	if _aeab.HasFormula() {
		if _, _dbca := _bec._bbb[ref]; _dbca {
			return _eg.MakeErrorResult("r\u0065\u0063\u0075\u0072\u0073\u0069\u006f\u006e\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0065\u0064\u0020d\u0075\u0072\u0069\u006e\u0067\u0020\u0065\u0076\u0061\u006cua\u0074\u0069\u006fn\u0020o\u0066\u0020" + ref)
		}
		_bec._bbb[ref] = struct{}{}
		_ebag := ev.Eval(_bec, _aeab.GetFormula())
		delete(_bec._bbb, ref)
		ev.SetCache(_dgd, _ebag)
		return _ebag
	}
	if _aeab.IsEmpty() {
		_dde := _eg.MakeEmptyResult()
		ev.SetCache(_dgd, _dde)
		return _dde
	} else if _aeab.IsNumber() {
		_cfb, _ := _aeab.GetValueAsNumber()
		_efad := _eg.MakeNumberResult(_cfb)
		ev.SetCache(_dgd, _efad)
		return _efad
	} else if _aeab.IsBool() {
		_gfcc, _ := _aeab.GetValueAsBool()
		_dge := _eg.MakeBoolResult(_gfcc)
		ev.SetCache(_dgd, _dge)
		return _dge
	}
	_ceg, _ := _aeab.GetRawValue()
	if _aeab.IsError() {
		_ebbc := _eg.MakeErrorResult("")
		_ebbc.ValueString = _ceg
		ev.SetCache(_dgd, _ebbc)
		return _ebbc
	}
	_bcga := _eg.MakeStringResult(_ceg)
	ev.SetCache(_dgd, _bcga)
	return _bcga
}

// IconScale maps values to icons.
type IconScale struct{ _bfdc *_bbg.CT_IconSet }

func (_cabc DataValidation) SetList() DataValidationList {
	_cabc.clear()
	_cabc._bge.TypeAttr = _bbg.ST_DataValidationTypeList
	_cabc._bge.OperatorAttr = _bbg.ST_DataValidationOperatorEqual
	return DataValidationList{_cabc._bge}
}

// X returns the inner wrapped XML type.
func (_caga RichText) X() *_bbg.CT_Rst { return _caga._agac }

// X returns the inner wrapped XML type.
func (_edbd DifferentialStyle) X() *_bbg.CT_Dxf { return _edbd._fdga }

// DefinedNames returns a slice of all defined names in the workbook.
func (_ggebg *Workbook) DefinedNames() []DefinedName {
	if _ggebg._bdff.DefinedNames == nil {
		return nil
	}
	_efda := []DefinedName{}
	for _, _fdbf := range _ggebg._bdff.DefinedNames.DefinedName {
		_efda = append(_efda, DefinedName{_fdbf})
	}
	return _efda
}

// Row returns the row of the cell marker.
func (_eebg CellMarker) Row() int32 { return _eebg._bef.Row }

// SetHorizontalAlignment sets the horizontal alignment of a cell style.
func (_bcff CellStyle) SetHorizontalAlignment(a _bbg.ST_HorizontalAlignment) {
	if _bcff._gcg.Alignment == nil {
		_bcff._gcg.Alignment = _bbg.NewCT_CellAlignment()
	}
	_bcff._gcg.Alignment.HorizontalAttr = a
	_bcff._gcg.ApplyAlignmentAttr = _b.Bool(true)
}
func (_fcdc *evalContext) NamedRange(ref string) _eg.Reference {
	for _, _cdcac := range _fcdc._befg._afga.DefinedNames() {
		if _cdcac.Name() == ref {
			return _eg.MakeRangeReference(_cdcac.Content())
		}
	}
	for _, _fba := range _fcdc._befg._afga.Tables() {
		if _fba.Name() == ref {
			return _eg.MakeRangeReference(_cg.Sprintf("\u0025\u0073\u0021%\u0073", _fcdc._befg.Name(), _fba.Reference()))
		}
	}
	return _eg.ReferenceInvalid
}
