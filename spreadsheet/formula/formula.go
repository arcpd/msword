package formula

import (
	_bd "bytes"
	_c "errors"
	_e "fmt"
	_cf "github.com/arcpd/msword/common/logger"
	_ad "github.com/arcpd/msword/internal/mergesort"
	_cd "github.com/arcpd/msword/internal/wildcard"
	_fg "github.com/arcpd/msword/spreadsheet/format"
	_fb "github.com/arcpd/msword/spreadsheet/reference"
	_eg "github.com/arcpd/msword/spreadsheet/update"
	_ee "io"
	_cc "math"
	_eeg "math/big"
	_d "math/rand"
	_f "regexp"
	_b "sort"
	_afc "strconv"
	_ga "strings"
	_af "sync"
	_eb "time"
	_a "unicode"
)

// Eval evaluates and returns an expression with prefix.
func (_eafgc PrefixExpr) Eval(ctx Context, ev Evaluator) Result {
	_gfbb := _eafgc._dfgab.Reference(ctx, ev)
	switch _gfbb.Type {
	case ReferenceTypeSheet:
		if _dcbb(_gfbb, ctx) {
			return MakeErrorResultType(ErrorTypeName, _e.Sprintf("\u0053h\u0065e\u0074\u0020\u0025\u0073\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064", _gfbb.Value))
		}
		_egcdb := ctx.Sheet(_gfbb.Value)
		return _eafgc._bcada.Eval(_egcdb, ev)
	default:
		return MakeErrorResult(_e.Sprintf("\u006e\u006f\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0066\u006f\u0072\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _gfbb.Type))
	}
}
func _dddg(_gde float64) bool { return _gde == 1 || _gde == 2 || _gde == 4 }
func NewLexer() *Lexer        { return &Lexer{_cgdbfa: make(chan *node)} }

// NewPrefixVerticalRange constructs a new full columns range with prefix.
func NewPrefixVerticalRange(pfx Expression, v string) Expression {
	_bacg := _ga.Split(v, "\u003a")
	if len(_bacg) != 2 {
		return nil
	}
	if _bacg[0] > _bacg[1] {
		_bacg[0], _bacg[1] = _bacg[1], _bacg[0]
	}
	return PrefixVerticalRange{_fcga: pfx, _agfae: _bacg[0], _fefae: _bacg[1]}
}

// LookupFunction looks up and returns a standard function or nil.
func LookupFunction(name string) Function {
	_agde.Lock()
	defer _agde.Unlock()
	if _dedb, _fbab := _bdgeg[name]; _fbab {
		return _dedb
	}
	return nil
}

const _ebeae = 2

func _eec(_dbg Result, _faa, _gag string) (float64, Result) {
	var _fdbf float64
	switch _dbg.Type {
	case ResultTypeNumber:
		_fdbf = float64(int(_dbg.ValueNumber))
	case ResultTypeString:
		_aaaf := DateValue([]Result{_dbg})
		if _aaaf.Type == ResultTypeError {
			return 0, MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020" + _faa + "\u0020\u0064\u0061\u0074\u0065\u0020\u0066\u006f\u0072\u0020" + _gag)
		}
		_fdbf = _aaaf.ValueNumber
	default:
		return 0, MakeErrorResult("\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020" + _gag)
	}
	if _fdbf < 0 {
		return 0, MakeErrorResultType(ErrorTypeNum, _faa+"\u0020\u0073\u0068ou\u006c\u0064\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	return _fdbf, _feb
}

var _bdgeg = map[string]Function{}

// Duration implements the Excel DURATION function.
func Duration(args []Result) Result {
	_ebbc, _aeeed := _cgbf(args, "\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e")
	if _aeeed.Type == ResultTypeError {
		return _aeeed
	}
	_dbga := _ebbc._bdce
	_ecgf := _ebbc._fbda
	_bbeca := _ebbc._agd
	_dccc := _ebbc._cacc
	_eded := _ebbc._efaf
	_fdcfd := _ebbc._bebb
	return _fbcc(_dbga, _ecgf, _bbeca, _dccc, _eded, _fdcfd)
}

// Yieldmat implements the Excel YIELDMAT function.
func Yieldmat(args []Result) Result {
	_fdee := len(args)
	if _fdee != 5 && _fdee != 6 {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0066\u0069v\u0065\u0020\u006f\u0072\u0020\u0073\u0069\u0078\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_dbefe, _bbed, _beca := _gfcg(args[0], args[1], "\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054")
	if _beca.Type == ResultTypeError {
		return _beca
	}
	_ebaf, _beca := _eec(args[2], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054")
	if _beca.Type == ResultTypeError {
		return _beca
	}
	if _ebaf >= _dbefe {
		return MakeErrorResult("\u0059\u0049\u0045\u004cD\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062e\u0020\u0062\u0065\u0066\u006fr\u0065\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049E\u004c\u0044\u004d\u0041T\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072a\u0074\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_gcde := args[3].ValueNumber
	if _gcde < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0049\u0045\u004c\u0044M\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072a\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0079\u0069\u0065\u006c\u0064\u0020o\u0066\u0020\u0074\u0079\u0070e\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_ededf := args[4].ValueNumber
	if _ededf <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "Y\u0049\u0045\u004c\u0044\u004d\u0041T\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0070\u0072\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069v\u0065")
	}
	_feda := 0
	if _fdee == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("\u0059I\u0045\u004cD\u004d\u0041\u0054 \u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_feda = int(args[5].ValueNumber)
		if !_gfe(_feda) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0059\u0049\u0045L\u0044\u004d\u0041\u0054")
		}
	}
	_defb, _beca := _fcfa(_ebaf, _bbed, _feda)
	if _beca.Type == ResultTypeError {
		return _beca
	}
	_becf, _beca := _fcfa(_ebaf, _dbefe, _feda)
	if _beca.Type == ResultTypeError {
		return _beca
	}
	_ecbd, _beca := _fcfa(_dbefe, _bbed, _feda)
	if _beca.Type == ResultTypeError {
		return _beca
	}
	_gacc := 1 + _defb*_gcde
	_gacc /= _ededf/100 + _becf*_gcde
	_gacc--
	_gacc /= _ecbd
	return MakeNumberResult(_gacc)
}

// Reference returns a string reference value to a vertical range.
func (_ddcbc VerticalRange) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeVerticalRange, Value: _ddcbc.verticalRangeReference()}
}
func _feac(_eacd, _aaga []float64, _beag float64) float64 {
	_dcebb := _beag + 1
	_cecc := 0.0
	_begb := len(_eacd)
	_fdfd := _aaga[0]
	for _begd := 1; _begd < _begb; _begd++ {
		_eced := (_aaga[_begd] - _fdfd) / 365
		_cecc -= _eced * _eacd[_begd] / _cc.Pow(_dcebb, _eced+1)
	}
	return _cecc
}

// TextJoin is an implementation of the Excel TEXTJOIN function.
func TextJoin(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("\u0054\u0045\u0058\u0054\u004aO\u0049\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074h\u0072\u0065\u0065\u0020\u006f\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeString {
		return MakeErrorResult("\u0054\u0045\u0058T\u004a\u004f\u0049\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0064\u0065\u006c\u0069\u006d\u0069\u0074\u0065\u0072\u0020\u0074\u006f\u0020\u0062\u0065 \u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	_agdc := args[0].ValueString
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0054\u0045\u0058\u0054\u004a\u004f\u0049\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065c\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0072 \u0062\u006f\u006f\u006c\u0065a\u006e")
	}
	_edde := args[1].ValueNumber != 0
	_ebga := _aefe(args[2:], []string{}, _edde)
	return MakeStringResult(_ga.Join(_ebga, _agdc))
}

// Dollarfr implements the Excel DOLLARFR function.
func Dollarfr(args []Result) Result {
	_ddee, _fceg, _bac := _gcga(args, "\u0044\u004f\u004c\u004c\u0041\u0052\u0046\u0052")
	if _bac.Type == ResultTypeError {
		return _bac
	}
	if _fceg == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0044\u004f\u004c\u004c\u0041R\u0046\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066r\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if _ddee == 0 {
		return MakeNumberResult(0)
	}
	_gcfg := _ddee < 0
	if _gcfg {
		_ddee = -_ddee
	}
	_bdgd := float64(int(_ddee))
	_dge := args[0].Value()
	_bgff := _ga.Split(_dge, "\u002e")
	_dcdb := 0.0
	if len(_bgff) > 1 {
		var _ddfd error
		_aefb := _bgff[1]
		_dcdb, _ddfd = _afc.ParseFloat(_aefb, 64)
		if _ddfd != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006fr \u0044\u004f\u004cL\u0041R\u0046\u0052")
		}
		_cbcfa := float64(len(_aefb))
		_dcdb /= _cc.Pow(10, _cbcfa)
	}
	_gdba := _dcdb*_fceg/_cc.Pow(10, float64(int(_cc.Log10(_fceg)))+1) + _bdgd
	if _gcfg {
		_gdba = -_gdba
	}
	return MakeNumberResult(_gdba)
}

var _gdgdf = [...]uint8{0, 20, 37, 60, 78, 96}

func _bgga(_cfgaa [][]Result, _beead int) [][]Result {
	_dgfd := [][]Result{}
	for _ceda := range _cfgaa {
		if _ceda == 0 {
			continue
		}
		_dacg := []Result{}
		for _gacbe := range _cfgaa {
			if _gacbe == _beead {
				continue
			}
			_dacg = append(_dacg, _cfgaa[_ceda][_gacbe])
		}
		_dgfd = append(_dgfd, _dacg)
	}
	return _dgfd
}
func _ccce(_gff string) (int, int, float64, bool, bool, Result) {
	_egc := ""
	_bfb := []string{}
	for _gbbb, _cbf := range _fbgc {
		_bfb = _cbf.FindStringSubmatch(_gff)
		if len(_bfb) > 1 {
			_egc = _gbbb
			break
		}
	}
	if _egc == "" {
		return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
	}
	_fee := _bfb[1] == ""
	_bfb = _bfb[49:]
	_cddf := len(_bfb)
	_aabc := _bfb[_cddf-1]
	_daec := _aabc == "\u0061\u006d"
	_dffg := _aabc == "\u0070\u006d"
	var _cae, _acdg int
	var _fbef float64
	var _cdf error
	switch _egc {
	case "\u0068\u0068":
		_cae, _cdf = _afc.Atoi(_bfb[0])
		if _cdf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
		}
		_acdg = 0
		_fbef = 0
	case "\u0068\u0068\u003am\u006d":
		_cae, _cdf = _afc.Atoi(_bfb[0])
		if _cdf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
		}
		_acdg, _cdf = _afc.Atoi(_bfb[2])
		if _cdf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
		}
		_fbef = 0
	case "\u006d\u006d\u003as\u0073":
		_cae = 0
		_acdg, _cdf = _afc.Atoi(_bfb[0])
		if _cdf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
		}
		_fbef, _cdf = _afc.ParseFloat(_bfb[2], 64)
		if _cdf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
		}
	case "\u0068\u0068\u003a\u006d\u006d\u003a\u0073\u0073":
		_cae, _cdf = _afc.Atoi(_bfb[0])
		if _cdf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
		}
		_acdg, _cdf = _afc.Atoi(_bfb[2])
		if _cdf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
		}
		_fbef, _cdf = _afc.ParseFloat(_bfb[4], 64)
		if _cdf != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
		}
	}
	if _acdg >= 60 {
		return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
	}
	if _daec || _dffg {
		if _cae > 12 || _fbef >= 60 {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
		} else if _cae == 12 {
			_cae = 0
		}
	} else if _cae >= 24 || _fbef >= 10000 {
		return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, _eadb)
	}
	return _cae, _acdg, _fbef, _dffg, _fee, _feb
}

// Xirr implements the Excel XIRR function.
func Xirr(args []Result) Result {
	_ebda := len(args)
	if _ebda != 2 && _ebda != 3 {
		return MakeErrorResult("\u0058\u0049RR\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s t\u0077o \u006f\u0072\u0020\u0074\u0068\u0072\u0065e \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_bcce, _dga := _gcfgb(args[0], args[1], "\u0058\u0049\u0052\u0052")
	if _dga.Type == ResultTypeError {
		return _dga
	}
	_gdfe := _bcce._faaf
	_cgaf := _bcce._ceee
	_feaa := 0.1
	if _ebda == 3 && args[2].Type != ResultTypeEmpty {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("\u0058\u0049\u0052\u0052\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0067\u0075\u0065\u0073\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_feaa = args[2].ValueNumber
		if _feaa <= -1 {
			return MakeErrorResult("\u0058\u0049\u0052\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0067\u0075\u0065\u0073\u0073\u0020\u0074\u006f\u0020\u0062e\u0020\u006d\u006f\u0072\u0065 \u0074\u0068a\u006e\u0020\u002d\u0031")
		}
	}
	return _dec(_gdfe, _cgaf, _feaa)
}

// Npv implements the Excel NPV function.
func Npv(args []Result) Result {
	_fefa := len(args)
	if _fefa < 2 {
		return MakeErrorResult("\u004e\u0050\u0056 r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f \u006fr\u0020m\u006f\u0072\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u004e\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020r\u0061\u0074\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_ecadf := args[0].ValueNumber
	if _ecadf == -1 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "")
	}
	_bcg := []float64{}
	for _, _ecfcb := range args[1:] {
		switch _ecfcb.Type {
		case ResultTypeNumber:
			_bcg = append(_bcg, _ecfcb.ValueNumber)
		case ResultTypeArray, ResultTypeList:
			_becg := _aadag(_ecfcb)
			for _, _cdac := range _becg {
				for _, _fcfge := range _cdac {
					if _fcfge.Type == ResultTypeNumber && !_fcfge.IsBoolean {
						_bcg = append(_bcg, _fcfge.ValueNumber)
					}
				}
			}
		}
	}
	_edgg := 0.0
	for _dbgc, _aaeb := range _bcg {
		_edgg += _aaeb / _cc.Pow(1+_ecadf, float64(_dbgc)+1)
	}
	return MakeNumberResult(_edgg)
}
func _bddd(_cfcg, _ecef int64) float64 { return float64(int(0.5 + float64((_ecef-_cfcg)/86400))) }

// PrefixVerticalRange is a range expression that when evaluated returns a list of Results from references like Sheet1!AA:IJ (all cells from columns AA to IJ of sheet 'Sheet1').
type PrefixVerticalRange struct {
	_fcga          Expression
	_agfae, _fefae string
}

// Reference returns a string reference value to a horizontal range.
func (_fefe HorizontalRange) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeHorizontalRange, Value: _fefe.horizontalRangeReference()}
}

// Median implements the MEDIAN function that returns the median of a range of
// values.
func Median(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u004d\u0045D\u0049\u0041\u004e\u0020r\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020l\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_gcada := _efba(args)
	_b.Float64s(_gcada)
	var _fdcb float64
	if len(_gcada)%2 == 0 {
		_fdcb = (_gcada[len(_gcada)/2-1] + _gcada[len(_gcada)/2]) / 2
	} else {
		_fdcb = _gcada[len(_gcada)/2]
	}
	return MakeNumberResult(_fdcb)
}

var _cfgcg, _aded, _bdfd, _caed, _bgda, _ddae *_f.Regexp

const _gbae = 57350

// MakeEmptyResult is ued when parsing an empty argument.
func MakeEmptyResult() Result { return Result{Type: ResultTypeEmpty} }

// Update updates references in the Range after removing a row/column.
func (_gfebf Range) Update(q *_eg.UpdateQuery) Expression {
	_bebfa := _gfebf
	if q.UpdateCurrentSheet {
		_bebfa._fbdb = _gfebf._fbdb.Update(q)
		_bebfa._agedb = _gfebf._agedb.Update(q)
	}
	return _bebfa
}
func _eage(_ccb, _egeb _eb.Time, _eaceg int) _eb.Time {
	_ebfc := _eb.Date(_ccb.Year(), _egeb.Month(), _egeb.Day(), 0, 0, 0, 0, _eb.UTC)
	if _ebfc.After(_ccb) {
		_ebfc = _ebfc.AddDate(-1, 0, 0)
	}
	for !_ebfc.After(_ccb) {
		_ebfc = _ebfc.AddDate(0, 12/_eaceg, 0)
	}
	return _ebfc
}
func _fcfa(_fda, _aed float64, _efbf int) (float64, Result) {
	_fdag, _cddcf := _bgee(_fda), _bgee(_aed)
	_ebbd := _fdag.Unix()
	_gggd := _cddcf.Unix()
	if _ebbd == _gggd {
		return 0, _feb
	}
	_fdbd, _gfb, _dcea := _fdag.Date()
	_bbbf, _fgfb, _gef := _cddcf.Date()
	_dege, _cddce := int(_gfb), int(_fgfb)
	var _eee, _cge float64
	switch _efbf {
	case 0:
		if _dcea == 31 {
			_dcea--
		}
		if _dcea == 30 && _gef == 31 {
			_gef--
		} else if _cgaa := _bee(_fdbd); _dege == 2 && ((_cgaa && _dcea == 29) || (!_cgaa && _dcea == 28)) {
			_dcea = 30
			if _fbee := _bee(_bbbf); _cddce == 2 && ((_fbee && _gef == 29) || (!_fbee && _gef == 28)) {
				_gef = 30
			}
		}
		_eee = float64((_bbbf-_fdbd)*360 + (_cddce-_dege)*30 + (_gef - _dcea))
		_cge = 360
	case 1:
		_eee = _aed - _fda
		_caa := _fdbd != _bbbf
		if _caa && (_bbbf != _fdbd+1 || _dege < _cddce || (_dege == _cddce && _dcea < _gef)) {
			_fgc := 0
			for _bebe := _fdbd; _bebe <= _bbbf; _bebe++ {
				_fgc += _fbeb(_bebe, 1)
			}
			_cge = float64(_fgc) / float64(_bbbf-_fdbd+1)
		} else {
			if !_caa && _bee(_fdbd) {
				_cge = 366
			} else {
				if _caa && ((_bee(_fdbd) && (_dege < 2 || (_dege == 2 && _dcea <= 29))) || (_bee(_bbbf) && (_cddce > 2 || (_cddce == 2 && _gef == 29)))) {
					_cge = 366
				} else {
					_cge = 365
				}
			}
		}
	case 2:
		_eee = _aed - _fda
		_cge = 360
	case 3:
		_eee = _aed - _fda
		_cge = 365
	case 4:
		if _dcea == 31 {
			_dcea--
		}
		if _gef == 31 {
			_gef--
		}
		_eee = float64((_bbbf-_fdbd)*360 + (_cddce-_dege)*30 + (_gef - _dcea))
		_cge = 360
	default:
		return 0, MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073 \u0066o\u0072\u0020\u0059\u0065\u0061\u0072\u0046r\u0061\u0063")
	}
	return _eee / _cge, _feb
}

// Update returns the same object as updating sheet references does not affect Bool.
func (_fbc Bool) Update(q *_eg.UpdateQuery) Expression { return _fbc }

// LookupFunctionComplex looks up and returns a complex function or nil.
func LookupFunctionComplex(name string) FunctionComplex {
	_agde.Lock()
	defer _agde.Unlock()
	if _fbbce, _dgga := _aaebee[name]; _dgga {
		return _fbbce
	}
	return nil
}

// Coupdaysnc implements the Excel COUPDAYSNC function.
func Coupdaysnc(args []Result) Result {
	_aaae, _bcee := _cebf(args, "\u0043\u004f\u0055\u0050\u0044\u0041\u0059\u0053\u004e\u0043")
	if _bcee.Type == ResultTypeError {
		return _bcee
	}
	return MakeNumberResult(_ddg(_aaae._bbaa, _aaae._aee, _aaae._gbf, _aaae._ceea))
}
func (_eaf BinOpType) String() string {
	if _eaf >= BinOpType(len(_ebc)-1) {
		return _e.Sprintf("\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065\u0028\u0025\u0064\u0029", _eaf)
	}
	return _dcf[_ebc[_eaf]:_ebc[_eaf+1]]
}

// String returns a string representation of CellRef.
func (_ebb CellRef) String() string { return _ebb._ab }
func _ggbd(_ccdfc, _abd, _dbfb, _dcaa float64, _egab int) float64 {
	var _cdcd float64
	if _ccdfc == 0 {
		_cdcd = _dcaa + _dbfb*_abd
	} else {
		_ebge := _cc.Pow(1+_ccdfc, _abd)
		if _egab == 1 {
			_cdcd = _dcaa*_ebge + _dbfb*(1+_ccdfc)*(_ebge-1)/_ccdfc
		} else {
			_cdcd = _dcaa*_ebge + _dbfb*(_ebge-1)/_ccdfc
		}
	}
	return -_cdcd
}
func _ebg(_fd BinOpType, _ce, _dg []Result) Result {
	_acg := []Result{}
	for _ebd := range _ce {
		_fe := _ce[_ebd].AsNumber()
		_eca := _dg[_ebd].AsNumber()
		if _fe.Type != ResultTypeNumber || _eca.Type != ResultTypeNumber {
			return MakeErrorResult("\u006e\u006f\u006e\u002d\u006e\u0075\u006e\u006d\u0065\u0072\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0062\u0069n\u0061\u0072\u0079\u0020\u006fp\u0065\u0072a\u0074\u0069\u006f\u006e")
		}
		switch _fd {
		case BinOpTypePlus:
			_acg = append(_acg, MakeNumberResult(_fe.ValueNumber+_eca.ValueNumber))
		case BinOpTypeMinus:
			_acg = append(_acg, MakeNumberResult(_fe.ValueNumber-_eca.ValueNumber))
		case BinOpTypeMult:
			_acg = append(_acg, MakeNumberResult(_fe.ValueNumber*_eca.ValueNumber))
		case BinOpTypeDiv:
			if _eca.ValueNumber == 0 {
				return MakeErrorResultType(ErrorTypeDivideByZero, "")
			}
			_acg = append(_acg, MakeNumberResult(_fe.ValueNumber/_eca.ValueNumber))
		case BinOpTypeExp:
			_acg = append(_acg, MakeNumberResult(_cc.Pow(_fe.ValueNumber, _eca.ValueNumber)))
		case BinOpTypeLT:
			_acg = append(_acg, MakeBoolResult(_fe.ValueNumber < _eca.ValueNumber))
		case BinOpTypeGT:
			_acg = append(_acg, MakeBoolResult(_fe.ValueNumber > _eca.ValueNumber))
		case BinOpTypeEQ:
			_acg = append(_acg, MakeBoolResult(_fe.ValueNumber == _eca.ValueNumber))
		case BinOpTypeLEQ:
			_acg = append(_acg, MakeBoolResult(_fe.ValueNumber <= _eca.ValueNumber))
		case BinOpTypeGEQ:
			_acg = append(_acg, MakeBoolResult(_fe.ValueNumber >= _eca.ValueNumber))
		case BinOpTypeNE:
			_acg = append(_acg, MakeBoolResult(_fe.ValueNumber != _eca.ValueNumber))
		default:
			return MakeErrorResult(_e.Sprintf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006c\u0069\u0073\u0074\u0020\u0062\u0069\u006e\u0061\u0072\u0079\u0020\u006fp \u0025\u0073", _fd))
		}
	}
	return MakeListResult(_acg)
}
func _cbgbd(_aead string, _dfbbg func(_dfde float64) float64) Function {
	return func(_cade []Result) Result {
		if len(_cade) != 1 {
			return MakeErrorResult(_aead + "\u0020\u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
		}
		_bbgdg := _cade[0].AsNumber()
		switch _bbgdg.Type {
		case ResultTypeNumber:
			_gabe := _dfbbg(_bbgdg.ValueNumber)
			if _cc.IsNaN(_gabe) {
				return MakeErrorResult(_aead + "\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0065\u0064\u0020\u004e\u0061\u004e")
			}
			if _cc.IsInf(_gabe, 0) {
				return MakeErrorResult(_aead + "\u0020r\u0065t\u0075\u0072\u006e\u0065\u0064 \u0069\u006ef\u0069\u006e\u0069\u0074\u0079")
			}
			if _gabe == 0 {
				return MakeErrorResultType(ErrorTypeDivideByZero, _aead+"\u0020d\u0069v\u0069\u0064\u0065\u0020\u0062\u0079\u0020\u007a\u0065\u0072\u006f")
			}
			return MakeNumberResult(1 / _gabe)
		case ResultTypeList, ResultTypeString:
			return MakeErrorResult(_aead + "\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u006e\u0075\u006de\u0072i\u0063\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		case ResultTypeError:
			return _bbgdg
		default:
			return MakeErrorResult(_e.Sprintf("\u0075\u006e\u0068a\u006e\u0064\u006c\u0065d\u0020\u0025\u0073\u0028\u0029\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _aead, _bbgdg.Type))
		}
	}
}

// Eval evaluates and returns the result of an empty expression.
func (_dae EmptyExpr) Eval(ctx Context, ev Evaluator) Result { return MakeEmptyResult() }

// AsNumber attempts to intepret a string cell value as a number. Upon success,
// it returns a new number result, upon  failure it returns the original result.
// This is used as functions return strings that can then act like number (e.g.
// LEFT(1.2345,3) + LEFT(1.2345,3) = 2.4)
func (_ggfa Result) AsNumber() Result {
	if _ggfa.Type == ResultTypeString {
		_aefeb, _fbdga := _afc.ParseFloat(_ggfa.ValueString, 64)
		if _fbdga == nil {
			return MakeNumberResult(_aefeb)
		}
	}
	if _ggfa.Type == ResultTypeEmpty {
		return MakeNumberResult(0)
	}
	return _ggfa
}

// Function is a standard function whose result only depends on its arguments.
type Function func(_fcbba []Result) Result

// Upper is an implementation of the Excel UPPER function that returns a upper
// case version of a string.
func Upper(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0055\u0050\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_deagf := args[0].AsString()
	if _deagf.Type != ResultTypeString {
		return MakeErrorResult("\u0055\u0050\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeStringResult(_ga.ToUpper(_deagf.ValueString))
}

// LastRow returns 0 for the invalid reference context.
func (_febaa *ivr) LastRow(colFrom string) int { return 0 }
func _daecb(_ccbdc string) *criteriaRegex {
	_dadab := &criteriaRegex{}
	if _ccbdc == "" {
		return _dadab
	}
	if _bbgf := _cfgcg.FindStringSubmatch(_ccbdc); len(_bbgf) > 1 {
		_dadab._afag = _bfcd
		_dadab._gbcfa = _bbgf[1]
	} else if _dggc := _aded.FindStringSubmatch(_ccbdc); len(_dggc) > 1 {
		_dadab._afag = _bfcd
		_dadab._gbcfa = _dggc[1]
	} else if _afadg := _ddae.FindStringSubmatch(_ccbdc); len(_afadg) > 1 {
		_dadab._afag = _eggf
		_dadab._gbcfa = _afadg[1]
	} else if _ggege := _bgda.FindStringSubmatch(_ccbdc); len(_ggege) > 1 {
		_dadab._afag = _edfe
		_dadab._gbcfa = _ggege[1]
	} else if _dfbg := _caed.FindStringSubmatch(_ccbdc); len(_dfbg) > 1 {
		_dadab._afag = _cdbac
		_dadab._gbcfa = _dfbg[1]
	} else if _bfdfe := _bdfd.FindStringSubmatch(_ccbdc); len(_bfdfe) > 1 {
		_dadab._afag = _dfef
		_dadab._gbcfa = _bfdfe[1]
	}
	return _dadab
}

var _aagf *_d.Rand

// Days is an implementation of the Excel DAYS() function.
func Days(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("D\u0041\u0059\u0053\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	var _ccef, _daed float64
	switch args[0].Type {
	case ResultTypeNumber:
		_daed = args[0].ValueNumber
	case ResultTypeString:
		_bfe := DateValue([]Result{args[0]})
		if _bfe.Type == ResultTypeError {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0065\u006e\u0064\u0020\u0064\u0061\u0074e\u0020\u0066\u006fr\u0020D\u0041\u0059\u0053")
		}
		_daed = _bfe.ValueNumber
	default:
		return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0066\u006fr\u0020D\u0041\u0059\u0053")
	}
	switch args[1].Type {
	case ResultTypeNumber:
		_ccef = args[1].ValueNumber
		if _ccef < 62 && _daed >= 62 {
			_ccef--
		}
	case ResultTypeString:
		_bae := DateValue([]Result{args[1]})
		if _bae.Type == ResultTypeError {
			return MakeErrorResult("\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0064\u0061\u0074\u0065\u0020\u0066\u006f\u0072\u0020DA\u0059\u0053")
		}
		_ccef = _bae.ValueNumber
	default:
		return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0066\u006fr\u0020D\u0041\u0059\u0053")
	}
	_fbb := float64(int(_daed - _ccef))
	return MakeNumberResult(_fbb)
}

// BinOpType is the binary operation operator type
//
//go:generate stringer -type=BinOpType
type BinOpType byte

// NewString constructs a new string expression.
func NewString(v string) Expression {
	v = _ga.Replace(v, "\u0022\u0022", "\u0022", -1)
	return String{_bedb: v}
}

// Reference returns a string reference value to a range.
func (_befc Range) Reference(ctx Context, ev Evaluator) Reference {
	_fgbdb := _befc._fbdb.Reference(ctx, ev)
	_gbdef := _befc._agedb.Reference(ctx, ev)
	if _fgbdb.Type == ReferenceTypeCell && _gbdef.Type == ReferenceTypeCell {
		return MakeRangeReference(_ccaeb(_fgbdb, _gbdef))
	}
	return ReferenceInvalid
}

// PrefixHorizontalRange is a range expression that when evaluated returns a list of Results from references like Sheet1!1:4 (all cells from rows 1 to 4 of sheet 'Sheet1').
type PrefixHorizontalRange struct {
	_dggf         Expression
	_egbcc, _bcaa int
}

// Search is an implementation of the Excel SEARCH().
func Search(args []Result) Result {
	_aagdf, _geddf := _faafg("\u0046\u0049\u004e\u0044", args)
	if _geddf.Type != ResultTypeEmpty {
		return _geddf
	}
	_bgdd := _ga.ToLower(_aagdf._eccf)
	if _bgdd == "" {
		return MakeNumberResult(1.0)
	}
	_eaaf := _ga.ToLower(_aagdf._bccec)
	_acbg := _aagdf._aeag
	_agfc := 1
	for _cgcd := range _eaaf {
		if _agfc < _acbg {
			_agfc++
			continue
		}
		_fbgda := _cd.Index(_bgdd, _eaaf[_cgcd:])
		if _fbgda == 0 {
			return MakeNumberResult(float64(_agfc))
		}
		_agfc++
	}
	return MakeErrorResultType(ErrorTypeValue, "\u004eo\u0074\u0020\u0066\u006f\u0075\u006ed")
}
func _gdbdb(_cfbf yyLexer) int { return _aabcbf().Parse(_cfbf) }

// ResultType is the type of the result
//
//go:generate stringer -type=ResultType
type ResultType byte

// Mround is an implementation of the Excel MROUND function.  It is not a
// generic rounding function and has some oddities to match Excel's behavior.
func Mround(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u004d\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0074\u0077o\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_cbcg := args[0].AsNumber()
	if _cbcg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072\u0073\u0074\u0020\u0061r\u0067\u0075\u006de\u006e\u0074\u0020\u0074o\u0020\u004d\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_faggg := float64(1)
	_cdcf := args[1].AsNumber()
	if _cdcf.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073e\u0063\u006fn\u0064\u0020\u0061\u0072g\u0075\u006d\u0065n\u0074\u0020\u0074\u006f\u0020\u004d\u0052\u004f\u0055ND\u0028\u0029\u0020m\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020n\u0075\u006db\u0065\u0072")
	}
	_faggg = _cdcf.ValueNumber
	if _faggg < 0 && _cbcg.ValueNumber > 0 || _faggg > 0 && _cbcg.ValueNumber < 0 {
		return MakeErrorResult("\u004d\u0052\u004fUN\u0044\u0028\u0029\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020s\u0069g\u006e\u0073\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068")
	}
	_faac := _cbcg.ValueNumber
	_faac, _fdcd := _cc.Modf(_faac / _faggg)
	if _cc.Trunc(_fdcd+0.5) > 0 {
		_faac++
	}
	return MakeNumberResult(_faac * _faggg)
}

// SumIf implements the SUMIF function.
func SumIf(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("\u0053\u0055\u004d\u0049\u0046\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0074\u0068\u0072e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_bcdbe := args[0]
	if _bcdbe.Type != ResultTypeArray && _bcdbe.Type != ResultTypeList {
		return MakeErrorResult("\u0053\u0055\u004d\u0049\u0046\u0020\u0072e\u0071\u0075\u0069r\u0065\u0073\u0020\u0066i\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_fffff := _aadag(_bcdbe)
	_bgaf := args[2]
	if _bgaf.Type != ResultTypeArray && _bgaf.Type != ResultTypeList {
		return MakeErrorResult("\u0053\u0055\u004dI\u0046\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0061\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074y\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_cbgcc := _aadag(_bgaf)
	_ggda := _cafa(args[1])
	_bbgee := 0.0
	for _cfdd, _fafgbb := range _fffff {
		for _ecde, _gfcde := range _fafgbb {
			if _addg(_gfcde, _ggda) {
				_bbgee += _cbgcc[_cfdd][_ecde].ValueNumber
			}
		}
	}
	return MakeNumberResult(_bbgee)
}

const _debac = 57378
const _dd = "\u0028\u0028\u005b0\u002d\u0039\u005d\u0029\u002b\u0029\u003a\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u005c\u002e\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u0028\u0020(\u0061\u006d\u007c\u0070\u006d\u0029\u0029\u003f"

func _cgbf(_cebfd []Result, _babe string) (*durationArgs, Result) {
	_abea := len(_cebfd)
	if _abea != 5 && _abea != 6 {
		return nil, MakeErrorResult(_babe + "\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s\u0020\u0066\u0069\u0076\u0065\u0020\u006fr\u0020\u0073\u0069\u0078\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_bddg, _fdcf, _gcgd := _gfcg(_cebfd[0], _cebfd[1], _babe)
	if _gcgd.Type == ResultTypeError {
		return nil, _gcgd
	}
	_gfbff := _cebfd[2]
	if _gfbff.Type != ResultTypeNumber {
		return nil, MakeErrorResult(_babe + "\u0020\u0072eq\u0075\u0069\u0072e\u0073\u0020\u0063\u006fupo\u006e r\u0061\u0074\u0065\u0020\u006f\u0066\u0020ty\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072")
	}
	_cgcb := _gfbff.ValueNumber
	if _cgcb < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, "\u0043\u006f\u0075po\u006e\u0020\u0072\u0061\u0074\u0065\u0020\u0073\u0068o\u0075l\u0064 \u006eo\u0074\u0020\u0062\u0065\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	_afbf := _cebfd[3]
	if _afbf.Type != ResultTypeNumber {
		return nil, MakeErrorResult(_babe + " \u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0079\u0069\u0065\u006cd\u0020\u0072\u0061\u0074\u0065\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072")
	}
	_gac := _afbf.ValueNumber
	if _gac < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006c\u0064\u0020r\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u006ft\u0020\u0062\u0065\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	_bdab := _cebfd[4]
	if _bdab.Type != ResultTypeNumber {
		return nil, MakeErrorResult(_babe + "\u0020\u0072\u0065qu\u0069\u0072\u0065\u0073\u0020\u0066\u0072\u0065\u0071u\u0065n\u0063y\u0020o\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_egfb := float64(int(_bdab.ValueNumber))
	if !_dddg(_egfb) {
		return nil, MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_dage := 0
	if _abea == 6 && _cebfd[5].Type != ResultTypeEmpty {
		_gbea := _cebfd[5]
		if _gbea.Type != ResultTypeNumber {
			return nil, MakeErrorResult(_babe + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020b\u0061\u0073\u0069\u0073\u0020\u006f\u0066 \u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_dage = int(_gbea.ValueNumber)
		if !_gfe(_dage) {
			return nil, MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062a\u0073\u0069\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020f\u006f\u0072\u0020"+_babe)
		}
	}
	return &durationArgs{_bddg, _fdcf, _cgcb, _gac, _egfb, _dage}, _feb
}
func _gdbf(_cde string) bool {
	for _, _dafd := range _dag {
		_acfb := _dafd.FindStringSubmatch(_cde)
		if len(_acfb) > 1 {
			return true
		}
	}
	return false
}

const _gbcfb = 57373

func _ffcg(_egba, _ecdb int) int {
	if _ecdb == 2 && _bee(_egba) {
		return 29
	} else {
		return _gdg[_ecdb-1]
	}
}

// Number is a nubmer expression.
type Number struct{ _fbfe float64 }

// Concat is an implementation of the Excel CONCAT() and deprecated CONCATENATE() function.
func Concat(args []Result) Result {
	_ceadc := _bd.Buffer{}
	for _, _dbdd := range args {
		switch _dbdd.Type {
		case ResultTypeString:
			_ceadc.WriteString(_dbdd.ValueString)
		case ResultTypeNumber:
			var _defa string
			if _dbdd.IsBoolean {
				if _dbdd.ValueNumber == 0 {
					_defa = "\u0046\u0041\u004cS\u0045"
				} else {
					_defa = "\u0054\u0052\u0055\u0045"
				}
			} else {
				_defa = _dbdd.AsString().ValueString
			}
			_ceadc.WriteString(_defa)
		default:
			return MakeErrorResult("\u0043\u004f\u004e\u0043\u0041T\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073")
		}
	}
	return MakeStringResult(_ceadc.String())
}
func init() {
	_cgc()
	RegisterFunction("\u0044\u0041\u0054\u0045", Date)
	RegisterFunction("\u0044A\u0054\u0045\u0044\u0049\u0046", DateDif)
	RegisterFunction("\u0044A\u0054\u0045\u0056\u0041\u004c\u0055E", DateValue)
	RegisterFunction("\u0044\u0041\u0059", Day)
	RegisterFunction("\u0044\u0041\u0059\u0053", Days)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0044\u0041\u0059\u0053", Days)
	RegisterFunction("\u0045\u0044\u0041T\u0045", Edate)
	RegisterFunction("\u0045O\u004d\u004f\u004e\u0054\u0048", Eomonth)
	RegisterFunction("\u004d\u0049\u004e\u0055\u0054\u0045", Minute)
	RegisterFunction("\u004d\u004f\u004eT\u0048", Month)
	RegisterFunction("\u004e\u004f\u0057", Now)
	RegisterFunction("\u0054\u0049\u004d\u0045", Time)
	RegisterFunction("\u0054I\u004d\u0045\u0056\u0041\u004c\u0055E", TimeValue)
	RegisterFunction("\u0054\u004f\u0044A\u0059", Today)
	RegisterFunctionComplex("\u0059\u0045\u0041\u0052", Year)
	RegisterFunction("\u0059\u0045\u0041\u0052\u0046\u0052\u0041\u0043", YearFrac)
}

// Error is an error expression.
type Error struct{ _gbb string }

// Right implements the Excel RIGHT(string,[n]) function which returns the
// rightmost n characters.
func Right(args []Result) Result {
	_caca := 1
	switch len(args) {
	case 1:
	case 2:
		if args[1].Type != ResultTypeNumber {
			return MakeErrorResult("\u0052\u0049\u0047\u0048\u0054\u0020\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_caca = int(args[1].ValueNumber)
		if _caca < 0 {
			return MakeErrorResult("R\u0049\u0047\u0048\u0054\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064 \u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u003e\u003d \u0030")
		}
		if _caca == 0 {
			return MakeStringResult("")
		}
	default:
		return MakeErrorResult("\u0052\u0049\u0047HT\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020o\u006ee\u0020o\u0072 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type == ResultTypeList {
		return MakeErrorResult("\u0052\u0049\u0047\u0048\u0054\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020c\u0061l\u006c\u0065\u0064\u0020\u006f\u006e\u0020\u0061\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_bbedd := args[0].Value()
	_bcdg := len(_bbedd)
	if _caca > _bcdg {
		return MakeStringResult(_bbedd)
	}
	return MakeStringResult(_bbedd[_bcdg-_caca : _bcdg])
}
func (_ddfa *yyParserImpl) Parse(yylex yyLexer) int {
	_bdebd := _eb.Now()
	var _fbfg int
	var _gedgg yySymType
	var _dagfe []yySymType
	_ = _dagfe
	_bega := _ddfa._bbbbg[:]
	Nerrs := 0
	Errflag := 0
	_fbbb := 0
	_ddfa._gaeec = -1
	_gedb := -1
	defer func() { _fbbb = -1; _ddfa._gaeec = -1; _gedb = -1 }()
	_gbfca := -1
	goto _ebbba
_acdc:
	return 0
_gbcbd:
	return 1
_ebbba:
	if _adgeff(_bdebd) {
		_cf.Log.Error("\u0050\u0061\u0072\u0073\u0065\u0020\u0074\u0069\u006d\u0065\u006f\u0075\u0074")
		goto _gbcbd
	}
	if _cegc >= 4 {
		_e.Printf("\u0063\u0068\u0061\u0072\u0020\u0025\u0076\u0020\u0069n\u0020\u0025\u0076\u000a", _dgda(_gedb), _fcafc(_fbbb))
	}
	_gbfca++
	if _gbfca >= len(_bega) {
		_bddfe := make([]yySymType, len(_bega)*2)
		copy(_bddfe, _bega)
		_bega = _bddfe
	}
	_bega[_gbfca] = _gedgg
	_bega[_gbfca]._bddde = _fbbb
_bdba:
	if _adgeff(_bdebd) {
		_cf.Log.Error("\u0050\u0061\u0072\u0073\u0065\u0020\u0074\u0069\u006d\u0065\u006f\u0075\u0074")
		goto _gbcbd
	}
	_fbfg = _gded[_fbbb]
	if _fbfg <= _agfb {
		goto _ddef
	}
	if _ddfa._gaeec < 0 {
		_ddfa._gaeec, _gedb = _feab(yylex, &_ddfa._bggb)
	}
	_fbfg += _gedb
	if _fbfg < 0 || _fbfg >= _gebf {
		goto _ddef
	}
	_fbfg = _gbeb[_fbfg]
	if _cgec[_fbfg] == _gedb {
		_ddfa._gaeec = -1
		_gedb = -1
		_gedgg = _ddfa._bggb
		_fbbb = _fbfg
		if Errflag > 0 {
			Errflag--
		}
		goto _ebbba
	}
_ddef:
	if _adgeff(_bdebd) {
		_cf.Log.Error("\u0050\u0061\u0072\u0073\u0065\u0020\u0074\u0069\u006d\u0065\u006f\u0075\u0074")
		goto _gbcbd
	}
	_fbfg = _ebgdg[_fbbb]
	if _fbfg == -2 {
		if _ddfa._gaeec < 0 {
			_ddfa._gaeec, _gedb = _feab(yylex, &_ddfa._bggb)
		}
		_bcfec := 0
		for {
			if _bcab[_bcfec+0] == -1 && _bcab[_bcfec+1] == _fbbb {
				break
			}
			_bcfec += 2
		}
		for _bcfec += 2; ; _bcfec += 2 {
			_fbfg = _bcab[_bcfec+0]
			if _fbfg < 0 || _fbfg == _gedb {
				break
			}
		}
		_fbfg = _bcab[_bcfec+1]
		if _fbfg < 0 {
			goto _acdc
		}
	}
	if _fbfg == 0 {
		switch Errflag {
		case 0:
			yylex.Error(_gaacb(_fbbb, _gedb))
			Nerrs++
			if _cegc >= 1 {
				_e.Printf("\u0025\u0073", _fcafc(_fbbb))
				_e.Printf("\u0020\u0073\u0061\u0077\u0020\u0025\u0073\u000a", _dgda(_gedb))
			}
			fallthrough
		case 1, 2:
			Errflag = 3
			for _gbfca >= 0 {
				_fbfg = _gded[_bega[_gbfca]._bddde] + _ebeae
				if _fbfg >= 0 && _fbfg < _gebf {
					_fbbb = _gbeb[_fbfg]
					if _cgec[_fbbb] == _ebeae {
						goto _ebbba
					}
				}
				if _cegc >= 2 {
					_e.Printf("\u0065\u0072r\u006f\u0072\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u0070\u006f\u0070\u0073\u0020\u0073\u0074\u0061\u0074\u0065 %\u0064\u000a", _bega[_gbfca]._bddde)
				}
				_gbfca--
			}
			goto _gbcbd
		case 3:
			if _cegc >= 2 {
				_e.Printf("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u0064\u0069s\u0063\u0061\u0072d\u0073 \u0025\u0073\u000a", _dgda(_gedb))
			}
			if _gedb == _dfdcd {
				goto _gbcbd
			}
			_ddfa._gaeec = -1
			_gedb = -1
			goto _bdba
		}
	}
	if _cegc >= 2 {
		_e.Printf("\u0072e\u0064u\u0063\u0065\u0020\u0025\u0076 \u0069\u006e:\u000a\u0009\u0025\u0076\u000a", _fbfg, _fcafc(_fbbb))
	}
	_aedc := _fbfg
	_efbe := _gbfca
	_ = _efbe
	_gbfca -= _fdcdf[_fbfg]
	if _gbfca+1 >= len(_bega) {
		_dcggd := make([]yySymType, len(_bega)*2)
		copy(_dcggd, _bega)
		_bega = _dcggd
	}
	_gedgg = _bega[_gbfca+1]
	_fbfg = _bdcb[_fbfg]
	_adce := _eefb[_fbfg]
	_gfgad := _adce + _bega[_gbfca]._bddde + 1
	if _gfgad >= _gebf {
		_fbbb = _gbeb[_adce]
	} else {
		_fbbb = _gbeb[_gfgad]
		if _cgec[_fbbb] != -_fbfg {
			_fbbb = _gbeb[_adce]
		}
	}
	switch _aedc {
	case 1:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			yylex.(*plex)._ffccb = _gedgg._ecac
		}
	case 3:
		_dagfe = _bega[_efbe-2 : _efbe+1]
		{
			_gedgg._ecac = _dagfe[2]._ecac
		}
	case 4:
		_dagfe = _bega[_efbe-4 : _efbe+1]
		{
		}
	case 5:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecac = NewBool(_dagfe[1]._ecdgg._eedag)
		}
	case 6:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecac = NewNumber(_dagfe[1]._ecdgg._eedag)
		}
	case 7:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecac = NewString(_dagfe[1]._ecdgg._eedag)
		}
	case 8:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecac = NewError(_dagfe[1]._ecdgg._eedag)
		}
	case 9:
		_dagfe = _bega[_efbe-2 : _efbe+1]
		{
			_gedgg._ecac = _dagfe[2]._ecac
		}
	case 10:
		_dagfe = _bega[_efbe-2 : _efbe+1]
		{
			_gedgg._ecac = NewNegate(_dagfe[2]._ecac)
		}
	case 15:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = _dagfe[2]._ecac
		}
	case 17:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewConstArrayExpr(_dagfe[2]._ecda)
		}
	case 18:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecda = append(_gedgg._ecda, _dagfe[1]._adeda)
		}
	case 19:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecda = append(_dagfe[1]._ecda, _dagfe[3]._adeda)
		}
	case 20:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._adeda = append(_gedgg._adeda, _dagfe[1]._ecac)
		}
	case 21:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._adeda = append(_dagfe[1]._adeda, _dagfe[3]._ecac)
		}
	case 23:
		_dagfe = _bega[_efbe-2 : _efbe+1]
		{
			_gedgg._ecac = NewPrefixExpr(_dagfe[1]._ecac, _dagfe[2]._ecac)
		}
	case 25:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecac = NewSheetPrefixExpr(_dagfe[1]._ecdgg._eedag)
		}
	case 26:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecac = NewCellRef(_dagfe[1]._ecdgg._eedag)
		}
	case 27:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewRange(_dagfe[1]._ecac, _dagfe[3]._ecac)
		}
	case 28:
		_dagfe = _bega[_efbe-4 : _efbe+1]
		{
			_gedgg._ecac = NewPrefixRangeExpr(_dagfe[1]._ecac, _dagfe[2]._ecac, _dagfe[4]._ecac)
		}
	case 29:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecac = NewNamedRangeRef(_dagfe[1]._ecdgg._eedag)
		}
	case 30:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecac = NewHorizontalRange(_dagfe[1]._ecdgg._eedag)
		}
	case 31:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._ecac = NewVerticalRange(_dagfe[1]._ecdgg._eedag)
		}
	case 32:
		_dagfe = _bega[_efbe-2 : _efbe+1]
		{
			_gedgg._ecac = NewPrefixHorizontalRange(_dagfe[1]._ecac, _dagfe[2]._ecdgg._eedag)
		}
	case 33:
		_dagfe = _bega[_efbe-2 : _efbe+1]
		{
			_gedgg._ecac = NewPrefixVerticalRange(_dagfe[1]._ecac, _dagfe[2]._ecdgg._eedag)
		}
	case 34:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypePlus, _dagfe[3]._ecac)
		}
	case 35:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeMinus, _dagfe[3]._ecac)
		}
	case 36:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeMult, _dagfe[3]._ecac)
		}
	case 37:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeDiv, _dagfe[3]._ecac)
		}
	case 38:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeExp, _dagfe[3]._ecac)
		}
	case 39:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeLT, _dagfe[3]._ecac)
		}
	case 40:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeGT, _dagfe[3]._ecac)
		}
	case 41:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeLEQ, _dagfe[3]._ecac)
		}
	case 42:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeGEQ, _dagfe[3]._ecac)
		}
	case 43:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeEQ, _dagfe[3]._ecac)
		}
	case 44:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeNE, _dagfe[3]._ecac)
		}
	case 45:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewBinaryExpr(_dagfe[1]._ecac, BinOpTypeConcat, _dagfe[3]._ecac)
		}
	case 47:
		_dagfe = _bega[_efbe-2 : _efbe+1]
		{
			_gedgg._ecac = NewFunction(_dagfe[1]._ecdgg._eedag, nil)
		}
	case 48:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._ecac = NewFunction(_dagfe[1]._ecdgg._eedag, _dagfe[2]._adeda)
		}
	case 49:
		_dagfe = _bega[_efbe-1 : _efbe+1]
		{
			_gedgg._adeda = append(_gedgg._adeda, _dagfe[1]._ecac)
		}
	case 50:
		_dagfe = _bega[_efbe-3 : _efbe+1]
		{
			_gedgg._adeda = append(_dagfe[1]._adeda, _dagfe[3]._ecac)
		}
	case 53:
		_dagfe = _bega[_efbe-0 : _efbe+1]
		{
			_gedgg._ecac = NewEmptyExpr()
		}
	}
	goto _ebbba
}
func _fgfe(_fgg BinOpType, _edd []Result, _fdg Result) Result {
	_dc := []Result{}
	switch _fdg.Type {
	case ResultTypeNumber:
		_bc := _fdg.ValueNumber
		for _be := range _edd {
			_bg := _edd[_be].AsNumber()
			if _bg.Type != ResultTypeNumber {
				return MakeErrorResult("\u006e\u006f\u006e\u002d\u006e\u0075\u006e\u006d\u0065\u0072\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0062\u0069n\u0061\u0072\u0079\u0020\u006fp\u0065\u0072a\u0074\u0069\u006f\u006e")
			}
			switch _fgg {
			case BinOpTypePlus:
				_dc = append(_dc, MakeNumberResult(_bg.ValueNumber+_bc))
			case BinOpTypeMinus:
				_dc = append(_dc, MakeNumberResult(_bg.ValueNumber-_bc))
			case BinOpTypeMult:
				_dc = append(_dc, MakeNumberResult(_bg.ValueNumber*_bc))
			case BinOpTypeDiv:
				if _bc == 0 {
					return MakeErrorResultType(ErrorTypeDivideByZero, "")
				}
				_dc = append(_dc, MakeNumberResult(_bg.ValueNumber/_bc))
			case BinOpTypeExp:
				_dc = append(_dc, MakeNumberResult(_cc.Pow(_bg.ValueNumber, _bc)))
			case BinOpTypeLT:
				_dc = append(_dc, MakeBoolResult(_bg.ValueNumber < _bc))
			case BinOpTypeGT:
				_dc = append(_dc, MakeBoolResult(_bg.ValueNumber > _bc))
			case BinOpTypeEQ:
				_dc = append(_dc, MakeBoolResult(_bg.ValueNumber == _bc))
			case BinOpTypeLEQ:
				_dc = append(_dc, MakeBoolResult(_bg.ValueNumber <= _bc))
			case BinOpTypeGEQ:
				_dc = append(_dc, MakeBoolResult(_bg.ValueNumber >= _bc))
			case BinOpTypeNE:
				_dc = append(_dc, MakeBoolResult(_bg.ValueNumber != _bc))
			default:
				return MakeErrorResult(_e.Sprintf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006c\u0069\u0073\u0074\u0020\u0062\u0069\u006e\u0061\u0072\u0079\u0020\u006fp \u0025\u0073", _fgg))
			}
		}
	case ResultTypeString:
		_cg := _fdg.ValueString
		for _ae := range _edd {
			_fbg := _edd[_ae].AsString()
			if _fbg.Type != ResultTypeString {
				return MakeErrorResult("\u006e\u006f\u006e\u002d\u006e\u0075\u006e\u006d\u0065\u0072\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0062\u0069n\u0061\u0072\u0079\u0020\u006fp\u0065\u0072a\u0074\u0069\u006f\u006e")
			}
			switch _fgg {
			case BinOpTypeLT:
				_dc = append(_dc, MakeBoolResult(_fbg.ValueString < _cg))
			case BinOpTypeGT:
				_dc = append(_dc, MakeBoolResult(_fbg.ValueString > _cg))
			case BinOpTypeEQ:
				_dc = append(_dc, MakeBoolResult(_fbg.ValueString == _cg))
			case BinOpTypeLEQ:
				_dc = append(_dc, MakeBoolResult(_fbg.ValueString <= _cg))
			case BinOpTypeGEQ:
				_dc = append(_dc, MakeBoolResult(_fbg.ValueString >= _cg))
			case BinOpTypeNE:
				_dc = append(_dc, MakeBoolResult(_fbg.ValueString != _cg))
			default:
				return MakeErrorResult(_e.Sprintf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006c\u0069\u0073\u0074\u0020\u0062\u0069\u006e\u0061\u0072\u0079\u0020\u006fp \u0025\u0073", _fgg))
			}
		}
	default:
		return MakeErrorResult("\u006e\u006f\u006e\u002d\u006e\u0075\u006e\u006d\u0065\u0072\u0069c\u0020\u0061\u006e\u0064\u0020\u006e\u006f\u006e-\u0073t\u0072\u0069\u006e\u0067\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u006e\u0061r\u0079\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u0069\u006f\u006e")
	}
	return MakeListResult(_dc)
}

// Lower is an implementation of the Excel LOWER function that returns a lower
// case version of a string.
func Lower(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004c\u004f\u0057\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_bfggc := args[0]
	switch _bfggc.Type {
	case ResultTypeError:
		return _bfggc
	case ResultTypeNumber, ResultTypeString:
		return _gbde(args[0])
	case ResultTypeList:
		_geea := _bfggc.ValueList
		_bacc := []Result{}
		for _, _gebe := range _geea {
			_efffg := _gbde(_gebe)
			if _efffg.Type == ResultTypeError {
				return _efffg
			}
			_bacc = append(_bacc, _efffg)
		}
		return MakeListResult(_bacc)
	case ResultTypeArray:
		_edcf := _bfggc.ValueArray
		_gccc := [][]Result{}
		for _, _gebg := range _edcf {
			_daga := []Result{}
			for _, _adedf := range _gebg {
				_acdbf := _gbde(_adedf)
				if _acdbf.Type == ResultTypeError {
					return _acdbf
				}
				_daga = append(_daga, _acdbf)
			}
			_gccc = append(_gccc, _daga)
		}
		return MakeArrayResult(_gccc)
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u004c\u004fW\u0045\u0052")
	}
}
func (_cbab ResultType) String() string {
	if _cbab >= ResultType(len(_gcafc)-1) {
		return _e.Sprintf("\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070e\u0028\u0025\u0064\u0029", _cbab)
	}
	return _abcd[_gcafc[_cbab]:_gcafc[_cbab+1]]
}

// EmptyExpr is an empty expression.
type EmptyExpr struct{}
type parsedSearchObject struct {
	_eccf  string
	_bccec string
	_aeag  int
}

func _deg(_ffc string, _aefd uint32) string {
	_aae := _fb.ColumnToIndex(_ffc)
	if _aae == _aefd {
		return "\u0023\u0052\u0045F\u0021"
	} else if _aae > _aefd {
		return _fb.IndexToColumn(_aae - 1)
	} else {
		return _ffc
	}
}

var _bgdbd = [...]string{}

func Unicode(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0055\u004e\u0049\u0043\u004fD\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_afgab := args[0].AsString()
	if _afgab.Type != ResultTypeString {
		return MakeErrorResult("\u0055\u004e\u0049\u0043\u004fD\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(_afgab.ValueString) == 0 {
		return MakeErrorResult("\u0055\u004e\u0049\u0043\u004f\u0044\u0045 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073 \u0061\u0020\u006e\u006f\u006e\u002d\u007a\u0065\u0072\u006f\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(float64(_afgab.ValueString[0]))
}

// Date is an implementation of the Excel DATE() function.
func Date(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0041TE\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s t\u0068re\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_gcef := int(args[0].ValueNumber)
	if _gcef < 0 || _gcef >= 10000 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074 \u0064\u0061\u0074\u0065")
	} else if _gcef <= 1899 {
		_gcef += 1900
	}
	_cgfc := _eb.Month(args[1].ValueNumber)
	_ebdd := int(args[2].ValueNumber)
	_bbgb := _gadd(_gcef, _cgfc, _ebdd)
	_fcaa := _bddd(_afcc, _bbgb) + 1
	if _fcaa < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074 \u0064\u0061\u0074\u0065")
	}
	return MakeNumberResult(_fcaa)
}

// Mid is an implementation of the Excel MID function that returns a copy
// of the string with each word capitalized.
func Mid(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u004d\u0049\u0044\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065e\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	_deefd := args[0]
	if _deefd.Type == ResultTypeError {
		return _deefd
	}
	if _deefd.Type != ResultTypeString && _deefd.Type != ResultTypeNumber && _deefd.Type != ResultTypeEmpty {
		return MakeErrorResult("\u004d\u0049\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0065x\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	_eegb := args[0].Value()
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0049D\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u0073\u0074\u0061\u0072\u0074\u005fn\u0075\u006d\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_gegd := int(args[1].ValueNumber)
	if _gegd < 1 {
		return MakeErrorResult("M\u0049\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0074\u0061\u0072\u0074\u005fn\u0075\u006d\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006dor\u0065\u0020\u0074h\u0061n\u0020\u0030")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0049D\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u005f\u0063\u0068a\u0072\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_afaf := int(args[2].ValueNumber)
	if _afaf < 0 {
		return MakeErrorResult("\u004d\u0049\u0044\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u005f\u0063\u0068a\u0072\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	_cfecb := len(_eegb)
	if _gegd > _cfecb {
		return MakeStringResult("")
	}
	_gegd--
	_fbcba := _gegd + _afaf
	if _fbcba > _cfecb {
		return MakeStringResult(_eegb[_gegd:])
	} else {
		return MakeStringResult(_eegb[_gegd:_fbcba])
	}
}

// Update returns the same object as updating sheet references does not affect EmptyExpr.
func (_cbe EmptyExpr) Update(q *_eg.UpdateQuery) Expression { return _cbe }

// Indirect is an implementation of the Excel INDIRECT function that returns the
// contents of a cell.
func Indirect(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 && len(args) != 2 {
		return MakeErrorResult("\u0049\u004e\u0044\u0049\u0052\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u006f\u0072 \u0074\u0077\u006f\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_debaa := args[0].AsString()
	if _debaa.Type != ResultTypeString {
		return MakeErrorResult("\u0049\u004e\u0044\u0049\u0052\u0045\u0043\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069r\u0073t\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066 \u0074\u0079\u0070\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	return ctx.Cell(_debaa.ValueString, ev)
}

// Disc implements the Excel DISC function.
func Disc(args []Result) Result {
	_eabf := len(args)
	if _eabf != 4 && _eabf != 5 {
		return MakeErrorResult("\u0044\u0049SC\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s f\u006fur\u0020\u006f\u0072\u0020\u0066\u0069\u0076e \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_aebba, _dgcg, _cgg := _gfcg(args[0], args[1], "\u0044\u0049\u0053\u0043")
	if _cgg.Type == ResultTypeError {
		return _cgg
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_cdeb := args[2].ValueNumber
	if _cdeb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "D\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0070\u0072\u0020\u0074o \u0062\u0065\u0020\u0070o\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0049S\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_acffb := args[3].ValueNumber
	if _acffb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0049\u0053\u0043\u0020\u0072\u0065q\u0075\u0069\u0072e\u0073\u0020\u0072e\u0064\u0065m\u0070\u0074\u0069\u006f\u006e\u0020t\u006f b\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fgfa := 0
	if _eabf == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0044\u0049\u0053\u0043\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_fgfa = int(args[4].ValueNumber)
		if !_gfe(_fgfa) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0066\u006f\u0072 \u0044\u0049\u0053\u0043")
		}
	}
	_eefee, _cgg := _fcfa(_aebba, _dgcg, _fgfa)
	if _cgg.Type == ResultTypeError {
		return _cgg
	}
	return MakeNumberResult((_acffb - _cdeb) / _acffb / _eefee)
}

// Price implements the Excel PRICE function.
func Price(args []Result) Result {
	_face := len(args)
	if _face != 6 && _face != 7 {
		return MakeErrorResult("\u0050\u0052I\u0043\u0045\u0020\u0072e\u0071\u0075i\u0072\u0065\u0073\u0020\u0073\u0069\u0078\u0020o\u0072\u0020\u0073\u0065\u0076\u0065\u006e\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_fggd, _fbgcg, _cecd := _gfcg(args[0], args[1], "\u0050\u0052\u0049C\u0045")
	if _cecd.Type == ResultTypeError {
		return _cecd
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052\u0049CE\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0072a\u0074e\u0020o\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_gfcd := args[2].ValueNumber
	if _gfcd < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u006eo\u0074\u0020\u0062\u0065\u0020n\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("P\u0052\u0049\u0043\u0045\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u0079\u0069\u0065\u006c\u0064\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072")
	}
	_egbg := args[3].ValueNumber
	if _egbg < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0079\u0069\u0065\u006c\u0064 \u0074\u006f\u0020\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("P\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0072\u0065\u0064em\u0070\u0074\u0069\u006fn\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_edaed := args[4].ValueNumber
	if _edaed <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049\u0043\u0045\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073 \u0072\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e \u0074\u006f\u0020\u0062\u0065 p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_egef := args[5]
	if _egef.Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0066\u0072\u0065\u0071\u0075e\u006e\u0063\u0079\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_cbfb := _egef.ValueNumber
	if !_dddg(_cbfb) {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_abce := 0
	if _face == 7 && args[6].Type != ResultTypeEmpty {
		if args[6].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0052\u0049C\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_abce = int(args[6].ValueNumber)
		if !_gfe(_abce) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069s\u0020a\u0072g\u0075m\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0050\u0052\u0049\u0043\u0045")
		}
	}
	_eedbd, _cecd := _cgce(_fggd, _fbgcg, _gfcd, _egbg, _edaed, _cbfb, _abce)
	if _cecd.Type == ResultTypeError {
		return _cecd
	}
	return MakeNumberResult(_eedbd)
}

// ISNUMBER is an implementation of the Excel ISNUMBER() function.
func IsNumber(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053NU\u004d\u0042\u0045R\u0028\u0029\u0020\u0061cce\u0070ts\u0020\u0061\u0020\u0073\u0069\u006e\u0067le\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeNumber)
}
func _gebc(_gecfd Context, _becb Evaluator, _ccdcg, _ggee string) Result {
	_fecf, _cfdb := _fb.ParseCellReference(_ccdcg)
	if _cfdb != nil {
		return MakeErrorResult(_e.Sprintf("\u0075\u006e\u0061bl\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073e\u0020r\u0061n\u0067e\u0020\u0025\u0073\u003a\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _ccdcg, _cfdb.Error()))
	}
	_bbbe, _efee := _fecf.ColumnIdx, _fecf.RowIdx
	_bdfec, _bbaf := _fb.ParseCellReference(_ggee)
	if _bbaf != nil {
		return MakeErrorResult(_e.Sprintf("\u0075\u006e\u0061bl\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073e\u0020r\u0061n\u0067e\u0020\u0025\u0073\u003a\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0025\u0073", _ggee, _bbaf.Error()))
	}
	_caeda, _gegf := _bdfec.ColumnIdx, _bdfec.RowIdx
	_abgcc := [][]Result{}
	for _fgcfe := _efee; _fgcfe <= _gegf; _fgcfe++ {
		_cdfb := []Result{}
		for _ecagc := _bbbe; _ecagc <= _caeda; _ecagc++ {
			_gbgbd := _gecfd.Cell(_e.Sprintf("\u0025\u0073\u0025\u0064", _fb.IndexToColumn(_ecagc), _fgcfe), _becb)
			_cdfb = append(_cdfb, _gbgbd)
		}
		_abgcc = append(_abgcc, _cdfb)
	}
	if len(_abgcc) == 1 {
		if len(_abgcc[0]) == 1 {
			return _abgcc[0][0]
		}
		return MakeListResult(_abgcc[0])
	}
	return MakeArrayResult(_abgcc)
}

const (
	ErrorTypeValue ErrorType = iota
	ErrorTypeNull
	ErrorTypeRef
	ErrorTypeName
	ErrorTypeNum
	ErrorTypeSpill
	ErrorTypeNA
	ErrorTypeDivideByZero
)
const _gggab = 57355

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

const _acbad int = 30

func _dad(_cgda, _ccfe, _eaba, _ggga int) int {
	if _ccfe > _eaba {
		return 0
	}
	if _gbgg(_ggga) {
		return (_eaba - _ccfe + 1) * 30
	}
	_bcdb := 0
	for _fade := _ccfe; _fade <= _eaba; _fade++ {
		_bcdb += _ffcg(_cgda, _fade)
	}
	return _bcdb
}

// Trim is an implementation of the Excel TRIM function that removes leading,
// trailing and consecutive spaces.
func Trim(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0054\u0052\u0049\u004d\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_fdba := args[0].AsString()
	if _fdba.Type != ResultTypeString {
		return MakeErrorResult("\u0054\u0052\u0049\u004d\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_ceeg := _bd.Buffer{}
	_aeefb := false
	_acbde := false
	_fdeag := 0
	for _, _aadd := range _fdba.ValueString {
		_eggb := _aadd == ' '
		if _eggb {
			if !_aeefb {
				continue
			}
			if !_acbde {
				_fdeag++
				_ceeg.WriteRune(_aadd)
			}
		} else {
			_fdeag = 0
			_aeefb = true
			_ceeg.WriteRune(_aadd)
		}
		_acbde = _eggb
	}
	_ceeg.Truncate(_ceeg.Len() - _fdeag)
	return MakeStringResult(_ceeg.String())
}

// CountIfs implements the COUNTIFS function.
func CountIfs(args []Result) Result {
	_cbeef := _eged(args, false, "\u0043\u004f\u0055\u004e\u0054\u0049\u0046\u0053")
	if _cbeef.Type != ResultTypeEmpty {
		return _cbeef
	}
	_ffabb := _ecbg(args)
	return MakeNumberResult(float64(len(_ffabb)))
}

// Update updates the FunctionCall references after removing a row/column.
func (_ccbc FunctionCall) Update(q *_eg.UpdateQuery) Expression {
	_cgdbf := []Expression{}
	for _, _gfebb := range _ccbc._dgce {
		_bddbg := _gfebb.Update(q)
		_cgdbf = append(_cgdbf, _bddbg)
	}
	return FunctionCall{_eecfd: _ccbc._eecfd, _dgce: _cgdbf}
}

type amorArgs struct {
	_dffd  float64
	_dca   float64
	_geg   float64
	_abge  float64
	_dfaaf int
	_afed  float64
	_cbdf  int
}

func _gcfa(_bfgad, _fcfac, _gbacd, _faaa, _dfbab, _febe float64) float64 {
	var _bcfc, _agcfb float64
	_ccfgb := 0.0
	_fbdfc := _cc.Ceil(_dfbab)
	_decc := _bfgad - _fcfac
	_bebg := false
	_dgdd := 0.0
	for _dbfc := 1.0; _dbfc <= _fbdfc; _dbfc++ {
		if !_bebg {
			_bcfc = _gege(_bfgad, _fcfac, _gbacd, _dbfc, _febe)
			_dgdd = _decc / (_gbacd - _dbfc + 1)
			if _dgdd > _bcfc {
				_agcfb = _dgdd
				_bebg = true
			} else {
				_agcfb = _bcfc
				_decc -= _bcfc
			}
		} else {
			_agcfb = _dgdd
		}
		if _dbfc == _fbdfc {
			_agcfb *= _dfbab + 1 - _fbdfc
		}
		_ccfgb += _agcfb
	}
	return _ccfgb
}

// Minute is an implementation of the Excel MINUTE() function.
func Minute(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004d\u0049\u004e\u0055T\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_cdc := args[0]
	switch _cdc.Type {
	case ResultTypeEmpty:
		return MakeNumberResult(0)
	case ResultTypeNumber:
		_fga := _bgee(_cdc.ValueNumber)
		return MakeNumberResult(float64(_fga.Minute()))
	case ResultTypeString:
		_fafa := _ga.ToLower(_cdc.ValueString)
		if !_gdbf(_fafa) {
			_, _, _, _gccg, _fgd := _ddd(_fafa)
			if _fgd.Type == ResultTypeError {
				_fgd.ErrorMessage = "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074s\u0020\u0066\u006f\u0072\u0020\u004d\u0049\u004e\u0055\u0054\u0045"
				return _fgd
			}
			if _gccg {
				return MakeNumberResult(0)
			}
		}
		_, _eefga, _, _, _, _dbbg := _ccce(_fafa)
		if _dbbg.Type == ResultTypeError {
			return _dbbg
		}
		return MakeNumberResult(float64(_eefga))
	default:
		return MakeErrorResult("\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u004d\u0049NU\u0054\u0045")
	}
}

// True is an implementation of the Excel TRUE() function.  It takes no
// arguments.
func True(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("\u0054\u0052\u0055E \u0074\u0061\u006b\u0065\u0073\u0020\u006e\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	return MakeBoolResult(true)
}

// Update returns the same object as updating sheet references does not affect Number.
func (_bddfec Number) Update(q *_eg.UpdateQuery) Expression { return _bddfec }

// MaxA is an implementation of the Excel MAXA() function.
func MaxA(args []Result) Result { return _bceb(args, true) }

// Mirr implements the Excel MIRR function.
func Mirr(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u004d\u0049R\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	if args[0].Type != ResultTypeList && args[0].Type != ResultTypeArray {
		return MakeErrorResult("M\u0049\u0052\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0076\u0061\u006c\u0075\u0065s\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020ar\u0072\u0061\u0079 \u0074y\u0070\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0049\u0052\u0052\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0069\u006e\u0061\u006e\u0063e\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_gccd := args[1].ValueNumber + 1
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0049\u0052\u0052\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0069\u006e\u0076\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_bcdf := args[2].ValueNumber + 1
	if _bcdf == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "")
	}
	_deab := _aadag(args[0])
	_baffd := float64(len(_deab))
	_bdad, _agaec := 0.0, 0.0
	_cgge, _abec := 1.0, 1.0
	_ebdde, _fbgb := false, false
	for _, _fafb := range _deab {
		for _, _gdcg := range _fafb {
			if _gdcg.Type == ResultTypeNumber && !_gdcg.IsBoolean {
				_ggce := _gdcg.ValueNumber
				if _ggce == 0 {
					continue
				} else {
					if _ggce > 0 {
						_ebdde = true
						_agaec += _gdcg.ValueNumber * _abec
					} else {
						_fbgb = true
						_bdad += _gdcg.ValueNumber * _cgge
					}
					_cgge /= _gccd
					_abec /= _bcdf
				}
			}
		}
	}
	if !_ebdde || !_fbgb {
		return MakeErrorResultType(ErrorTypeDivideByZero, "")
	}
	_agad := -_agaec / _bdad
	_agad *= _cc.Pow(_bcdf, _baffd-1)
	_agad = _cc.Pow(_agad, 1/(_baffd-1))
	return MakeNumberResult(_agad - 1)
}

const _ceb = "\u0028(\u005b0\u002d\u0039\u005d\u0029\u002b)\u0020\u0028a\u006d\u007c\u0070\u006d\u0029"

// Update updates references in the PrefixRangeExpr after removing a row/column.
func (_dbfgc PrefixRangeExpr) Update(q *_eg.UpdateQuery) Expression {
	_agaad := _dbfgc
	_cadaaf := _dbfgc._beee.String()
	if _cadaaf == q.SheetToUpdate {
		_gcac := *q
		_gcac.UpdateCurrentSheet = true
		_agaad._bbbbge = _dbfgc._bbbbge.Update(&_gcac)
		_agaad._affc = _dbfgc._affc.Update(&_gcac)
	}
	return _agaad
}

// TimeValue is an implementation of the Excel TIMEVALUE() function.
func TimeValue(args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeString {
		return MakeErrorResult("\u0054I\u004d\u0045V\u0041\u004c\u0055\u0045 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069ng\u006c\u0065\u0020s\u0074\u0072i\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_edae := _ga.ToLower(args[0].ValueString)
	if !_gdbf(_edae) {
		_, _, _, _abef, _gdb := _ddd(_edae)
		if _gdb.Type == ResultTypeError {
			_gdb.ErrorMessage = "\u0049\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020f\u006f\u0072\u0020\u0054\u0049\u004d\u0045V\u0041\u004c\u0055\u0045"
			return _gdb
		}
		if _abef {
			return MakeNumberResult(0)
		}
	}
	_ggag, _acc, _dbbc, _bbd, _, _bbb := _ccce(_edae)
	if _bbb.Type == ResultTypeError {
		return _bbb
	}
	_acf := _bcd(float64(_ggag), float64(_acc), _dbbc)
	if _bbd {
		_acf += 0.5
	} else if _acf >= 1 {
		_acf -= float64(int(_acf))
	}
	return MakeNumberResult(_acf)
}

// GetFilename returns an empty string for the invalid reference context.
func (_bbde *ivr) GetFilename() string { return "" }
func _gcfgb(_eafc, _fdab Result, _abfc string) (*xargs, Result) {
	if _eafc.Type != ResultTypeList && _eafc.Type != ResultTypeArray {
		return nil, MakeErrorResult(_abfc + "\u0020\u0072eq\u0075\u0069\u0072e\u0073\u0020\u0076\u0061lue\u0073 t\u006f\u0020\u0062\u0065\u0020\u006f\u0066 a\u0072\u0072\u0061\u0079\u0020\u0074\u0079p\u0065")
	}
	_dfae := _aadag(_eafc)
	_bfec := []float64{}
	for _, _babea := range _dfae {
		for _, _affe := range _babea {
			if _affe.Type == ResultTypeNumber && !_affe.IsBoolean {
				_bfec = append(_bfec, _affe.ValueNumber)
			} else {
				return nil, MakeErrorResult(_abfc + "\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006dbe\u0072\u0073")
			}
		}
	}
	_efad := len(_bfec)
	if len(_bfec) < 2 {
		return nil, MakeErrorResultType(ErrorTypeNum, "")
	}
	if _fdab.Type != ResultTypeList && _fdab.Type != ResultTypeArray {
		return nil, MakeErrorResult(_abfc + " \u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0064\u0061\u0074\u0065s\u0020\u0074\u006f\u0020\u0062\u0065\u0020o\u0066\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0074\u0079p\u0065")
	}
	_dbfg := _aadag(_fdab)
	_efddg := []float64{}
	_gbcdb := 0.0
	for _, _ecae := range _dbfg {
		for _, _bfdb := range _ecae {
			if _bfdb.Type == ResultTypeNumber && !_bfdb.IsBoolean {
				_cdecb := float64(int(_bfdb.ValueNumber))
				if _cdecb < _gbcdb {
					return nil, MakeErrorResultType(ErrorTypeNum, _abfc+" \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0064\u0061\u0074\u0065\u0073\u0020\u0074\u006f\u0020b\u0065\u0020\u0069\u006e\u0020\u0061\u0073\u0063\u0065\u006edi\u006e\u0067\u0020o\u0072d\u0065\u0072")
				}
				_efddg = append(_efddg, _cdecb)
				_gbcdb = _cdecb
			} else {
				return nil, MakeErrorResult(_abfc + "\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0064\u0061\u0074\u0065\u0073\u0020t\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0073")
			}
		}
	}
	if len(_efddg) != _efad {
		return nil, MakeErrorResultType(ErrorTypeNum, "")
	}
	return &xargs{_bfec, _efddg}, MakeEmptyResult()
}

// Nper implements the Excel NPER function.
func Nper(args []Result) Result {
	_abefd := len(args)
	if _abefd < 3 || _abefd > 5 {
		return MakeErrorResult("\u004e\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067e\u0020\u006f\u0066\u0020\u0033\u0020\u0061\u006e\u0064\u0020\u0035")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("N\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_bfcb := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u004e\u0050\u0045\u0052\u0020\u0072\u0065q\u0075\u0069\u0072e\u0073\u0020\u0070\u0061y\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fcfab := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u004e\u0050\u0045\u0052\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_efgf := args[2].ValueNumber
	_ggaa := 0.0
	if _abefd >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("\u004e\u0050\u0045\u0052\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0075\u0074\u0075\u0072\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_ggaa = args[3].ValueNumber
	}
	_bdfa := 0.0
	if _abefd == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("N\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_bdfa = args[4].ValueNumber
		if _bdfa != 0 {
			_bdfa = 1
		}
	}
	_cddb := _fcfab*(1+_bfcb*_bdfa) - _ggaa*_bfcb
	_fbea := (_efgf*_bfcb + _fcfab*(1+_bfcb*_bdfa))
	return MakeNumberResult(_cc.Log(_cddb/_fbea) / _cc.Log(1+_bfcb))
}

// Large implements the Excel LARGE function.
func Large(args []Result) Result { return _gceb(args, true) }

// Offset is an implementation of the Excel OFFSET function.
func Offset(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 3 && len(args) != 5 {
		return MakeErrorResult("\u004f\u0046\u0046\u0053\u0045\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0074\u0068\u0072\u0065e\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_fdfb := args[0].Ref
	for _fdfb.Type == ReferenceTypeNamedRange {
		_fdfb = ctx.NamedRange(_fdfb.Value)
	}
	_dcae := ""
	switch _fdfb.Type {
	case ReferenceTypeCell:
		_dcae = _fdfb.Value
	case ReferenceTypeRange:
		_gafb := _ga.Split(_fdfb.Value, "\u003a")
		if len(_gafb) == 2 {
			_dcae = _gafb[0]
		}
	default:
		return MakeErrorResult(_e.Sprintf("\u0049\u006ev\u0061\u006c\u0069\u0064\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0069\u006e\u0020\u004f\u0046\u0046\u0053\u0045\u0054\u0028\u0029: \u0025\u0073", _fdfb.Type))
	}
	_ebgd, _gedf := _fb.ParseCellReference(_dcae)
	if _gedf != nil {
		return MakeErrorResult(_e.Sprintf("\u0070\u0061\u0072s\u0065\u0020\u006f\u0072i\u0067\u0069\u006e\u0020\u0065\u0072\u0072o\u0072\u0020\u004f\u0046\u0046\u0053\u0045\u0054\u0028\u0029\u003a\u0020\u0025\u0073", _gedf.Error()))
	}
	_dead, _abff, _dedc := _ebgd.Column, _ebgd.RowIdx, _ebgd.SheetName
	_ddgg := args[1].AsNumber()
	if _ddgg.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0046\u0046SE\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020n\u0075m\u0065r\u0069\u0063\u0020\u0072\u006f\u0077\u0020\u006f\u0066\u0066\u0073\u0065\u0074")
	}
	_bfce := args[2].AsNumber()
	if _bfce.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0046\u0046SE\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020n\u0075m\u0065r\u0069\u0063\u0020\u0063\u006f\u006c\u0020\u006f\u0066\u0066\u0073\u0065\u0074")
	}
	var _cbba, _egda Result
	if len(args) == 3 {
		_cbba = MakeNumberResult(1)
		_egda = MakeNumberResult(1)
	} else {
		_cbba = args[3].AsNumber()
		if _cbba.Type != ResultTypeNumber {
			return MakeErrorResult("\u004f\u0046\u0046\u0053\u0045\u0054\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u006e\u0075m\u0065\u0072\u0069\u0063\u0020\u0068\u0065\u0069\u0067\u0068\u0074")
		}
		if _cbba.ValueNumber == 0 {
			return MakeErrorResultType(ErrorTypeRef, "")
		}
		_egda = args[4].AsNumber()
		if _egda.Type != ResultTypeNumber {
			return MakeErrorResult("\u004f\u0046F\u0053\u0045\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0077id\u0074\u0068")
		}
		if _egda.ValueNumber == 0 {
			return MakeErrorResultType(ErrorTypeRef, "")
		}
	}
	_afba := _fb.ColumnToIndex(_dead)
	_baag := _abff + uint32(_ddgg.ValueNumber)
	_afgf := _afba + uint32(_bfce.ValueNumber)
	_edce := _baag + uint32(_cbba.ValueNumber)
	_dggbf := _afgf + uint32(_egda.ValueNumber)
	if _cbba.ValueNumber > 0 {
		_edce--
	} else {
		_edce++
		_baag, _edce = _edce, _baag
	}
	if _egda.ValueNumber > 0 {
		_dggbf--
	} else {
		_dggbf++
		_afgf, _dggbf = _dggbf, _afgf
	}
	_adec := _e.Sprintf("\u0025\u0073\u0025\u0064", _fb.IndexToColumn(_afgf), _baag)
	_eead := _e.Sprintf("\u0025\u0073\u0025\u0064", _fb.IndexToColumn(_dggbf), _edce)
	if _dedc == "" {
		return _gebc(ctx, ev, _adec, _eead)
	} else {
		return _gebc(ctx.Sheet(_dedc), ev, _adec, _eead)
	}
}
func (_dgb *evCache) SetCache(key string, value Result) {
	_dgb._ef.Lock()
	_dgb._bgg[key] = value
	_dgb._ef.Unlock()
}
func _ccefc(_fbdf, _baeb _eb.Time, _dbge, _daecf int) _eb.Time {
	_bcc := _baeb
	_dgd := _fbdf.Year() - _baeb.Year()
	_bcc = _bcc.AddDate(_dgd, 0, 0)
	if _fbdf.After(_bcc) {
		_bcc = _bcc.AddDate(1, 0, 0)
	}
	_dafc := -12 / _dbge
	for _bcc.After(_fbdf) {
		_bcc = _bcc.AddDate(0, _dafc, 0)
	}
	return _bcc
}
func _efba(_ffbb []Result) []float64 {
	_gbfeg := make([]float64, 0)
	for _, _adcga := range _ffbb {
		if _adcga.Type == ResultTypeEmpty {
			continue
		}
		_adcga = _adcga.AsNumber()
		switch _adcga.Type {
		case ResultTypeNumber:
			if !_adcga.IsBoolean {
				_gbfeg = append(_gbfeg, _adcga.ValueNumber)
			}
		case ResultTypeList, ResultTypeArray:
			_gbfeg = append(_gbfeg, _efba(_adcga.ListValues())...)
		case ResultTypeString:
		default:
			_cf.Log.Debug("\u0075\u006e\u0068\u0061\u006ed\u006c\u0065\u0064\u0020\u0065\u0078\u0074\u0072\u0061\u0063\u0074\u004e\u0075m\u0062\u0065\u0072\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _adcga.Type)
		}
	}
	return _gbfeg
}

// CellRef is a reference to a single cell
type CellRef struct{ _ab string }

var _bdcb = [...]int{0, 7, 3, 3, 3, 8, 8, 8, 8, 1, 1, 1, 2, 2, 2, 2, 2, 14, 15, 15, 17, 17, 4, 4, 4, 13, 5, 6, 6, 6, 6, 6, 6, 6, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 9, 9, 9, 16, 16, 11, 10, 10}

// Counta implements the COUNTA function.
func Counta(args []Result) Result { return MakeNumberResult(_aabce(args, _cbfbc)) }

// Cumipmt implements the Excel CUMIPMT function.
func Cumipmt(args []Result) Result {
	_dcef, _dfb := _aabcb(args, "\u0043U\u004d\u0049\u0050\u004d\u0054")
	if _dfb.Type == ResultTypeError {
		return _dfb
	}
	_ddc := _dcef._ffa
	_cdfc := _dcef._dbcdb
	_efd := _dcef._bdeb
	_ccee := _dcef._gggb
	_dfba := _dcef._gffb
	_ffab := _dcef._cfgc
	_dgg := _ecc(_ddc, _cdfc, _efd, 0, _ffab)
	_bedf := 0.0
	if _ccee == 1 {
		if _ffab == 0 {
			_bedf = -_efd
			_ccee++
		}
	}
	for _ffb := _ccee; _ffb <= _dfba; _ffb++ {
		if _ffab == 1 {
			_bedf += _ggbd(_ddc, _ffb-2, _dgg, _efd, 1) - _dgg
		} else {
			_bedf += _ggbd(_ddc, _ffb-1, _dgg, _efd, 0)
		}
	}
	_bedf *= _ddc
	return MakeNumberResult(_bedf)
}

var _accgc = [...]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}

// Parse parses a string to get an Expression.
func ParseString(s string) Expression {
	if s == "" {
		return NewEmptyExpr()
	}
	return Parse(_ga.NewReader(s))
}

// FunctionComplex is a function whose result  depends on its arguments and the
// context that it's in.  As an example, INDIRECT is a complex function so that
// INDIRECT("A1") which returns the value of the "A1" cell in a sheet can use
// the context to reach into the sheet and pull out required values.
type FunctionComplex func(_bdge Context, _gcaf Evaluator, _dcfd []Result) Result

func (_bgc *evCache) GetFromCache(key string) (Result, bool) {
	_bgc._ef.Lock()
	_gd, _caf := _bgc._bgg[key]
	_bgc._ef.Unlock()
	return _gd, _caf
}
func _bfbe(_abba, _gace []float64, _cafd float64) float64 {
	_gggg := _cafd + 1
	_bdgb := _abba[0]
	_edec := len(_abba)
	_ggeg := _gace[0]
	for _dda := 1; _dda < _edec; _dda++ {
		_bdgb += _abba[_dda] / _cc.Pow(_gggg, (_gace[_dda]-_ggeg)/365)
	}
	return _bdgb
}
func (_ecbc PrefixHorizontalRange) horizontalRangeReference(_cdeff string) string {
	return _e.Sprintf("\u0025\u0073\u0021\u0025\u0064\u003a\u0025\u0064", _cdeff, _ecbc._egbcc, _ecbc._bcaa)
}

// String returns a string of a range.
func (_abdd Range) String() string {
	return _e.Sprintf("\u0025\u0073\u003a%\u0073", _abdd._fbdb.String(), _abdd._agedb.String())
}

type plex struct {
	_cegce chan *node
	_ffccb Expression
	_ddcd  string
}
type rmode byte

// Find is an implementation of the Excel FIND().
func Find(args []Result) Result {
	_gbfac, _aeeaa := _faafg("\u0046\u0049\u004e\u0044", args)
	if _aeeaa.Type != ResultTypeEmpty {
		return _aeeaa
	}
	_dcbe := _gbfac._eccf
	if _dcbe == "" {
		return MakeNumberResult(1.0)
	}
	_egbc := _gbfac._bccec
	_dedf := _gbfac._aeag
	_aedgb := 1
	for _gfce := range _egbc {
		if _aedgb < _dedf {
			_aedgb++
			continue
		}
		_dfaf := _ga.Index(_egbc[_gfce:], _dcbe)
		if _dfaf == 0 {
			return MakeNumberResult(float64(_aedgb))
		}
		_aedgb++
	}
	return MakeErrorResultType(ErrorTypeValue, "\u004eo\u0074\u0020\u0066\u006f\u0075\u006ed")
}

// Ppmt implements the Excel PPPMT function.
func Ppmt(args []Result) Result {
	_cfd := len(args)
	if _cfd < 4 || _cfd > 6 {
		return MakeErrorResult("\u0050\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006ff\u0020\u0066\u006f\u0075\u0072\u0020a\u006e\u0064\u0020s\u0069\u0078")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("P\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_cffe := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0050\u004dT\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_eege := args[1].ValueNumber
	if _eege <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "P\u0050\u004d\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020p\u0065\u0072\u0069\u006f\u0064\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069v\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cdec := args[2].ValueNumber
	if _cdec < _eege {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064s\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u0074\u0020\u006c\u0065s\u0073\u0020\u0074\u0068\u0061\u006e \u0070\u0065\u0072i\u006f\u0064")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0050\u004d\u0054\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_aacc := args[3].ValueNumber
	_dcb := 0.0
	if _cfd >= 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0050\u004d\u0054\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0075\u0074\u0075\u0072\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_dcb = args[4].ValueNumber
	}
	_babfb := 0
	if _cfd == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("P\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_babfb = int(args[5].ValueNumber)
		if _babfb != 0 {
			_babfb = 1
		}
	}
	return MakeNumberResult(_ecc(_cffe, _cdec, _aacc, _dcb, _babfb) - _dagb(_cffe, _eege, _cdec, _aacc, _dcb, _babfb))
}

// Exact is an implementation of the Excel EXACT() which compares two strings.
func Exact(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0043\u004f\u004e\u0043\u0041\u0054\u0045N\u0041\u0054\u0045(\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ggegb := args[0].AsString()
	_edbbe := args[1].AsString()
	if _ggegb.Type != ResultTypeString || _edbbe.Type != ResultTypeString {
		return MakeErrorResult("\u0043\u004f\u004e\u0043\u0041\u0054\u0045N\u0041\u0054\u0045(\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	return MakeBoolResult(_ggegb.ValueString == _edbbe.ValueString)
}
func init() {
	RegisterFunction("\u0041\u0043\u0043\u0052\u0049\u004e\u0054\u004d", Accrintm)
	RegisterFunction("\u0041M\u004f\u0052\u0044\u0045\u0047\u0052C", Amordegrc)
	RegisterFunction("\u0041\u004d\u004f\u0052\u004c\u0049\u004e\u0043", Amorlinc)
	RegisterFunction("\u0043O\u0055\u0050\u0044\u0041\u0059\u0042S", Coupdaybs)
	RegisterFunction("\u0043\u004f\u0055\u0050\u0044\u0041\u0059\u0053", Coupdays)
	RegisterFunction("\u0043\u004f\u0055\u0050\u0044\u0041\u0059\u0053\u004e\u0043", Coupdaysnc)
	RegisterFunction("\u0043O\u0055\u0050\u004e\u0055\u004d", Coupnum)
	RegisterFunction("\u0043O\u0055\u0050\u004e\u0043\u0044", Coupncd)
	RegisterFunction("\u0043O\u0055\u0050\u0050\u0043\u0044", Couppcd)
	RegisterFunction("\u0043U\u004d\u0049\u0050\u004d\u0054", Cumipmt)
	RegisterFunction("\u0043\u0055\u004d\u0050\u0052\u0049\u004e\u0043", Cumprinc)
	RegisterFunction("\u0044\u0042", Db)
	RegisterFunction("\u0044\u0044\u0042", Ddb)
	RegisterFunction("\u0044\u0049\u0053\u0043", Disc)
	RegisterFunction("\u0044\u004f\u004c\u004c\u0041\u0052\u0044\u0045", Dollarde)
	RegisterFunction("\u0044\u004f\u004c\u004c\u0041\u0052\u0046\u0052", Dollarfr)
	RegisterFunction("\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e", Duration)
	RegisterFunction("\u0045\u0046\u0046\u0045\u0043\u0054", Effect)
	RegisterFunction("\u0046\u0056", Fv)
	RegisterFunction("\u0046\u0056\u0053\u0043\u0048\u0045\u0044\u0055\u004c\u0045", Fvschedule)
	RegisterFunction("\u0049N\u0054\u0052\u0041\u0054\u0045", Intrate)
	RegisterFunction("\u0049\u0050\u004d\u0054", Ipmt)
	RegisterFunction("\u0049\u0052\u0052", Irr)
	RegisterFunction("\u0049\u0053\u0050M\u0054", Ispmt)
	RegisterFunction("\u004dD\u0055\u0052\u0041\u0054\u0049\u004fN", Mduration)
	RegisterFunction("\u004d\u0049\u0052\u0052", Mirr)
	RegisterFunction("\u004eO\u004d\u0049\u004e\u0041\u004c", Nominal)
	RegisterFunction("\u004e\u0050\u0045\u0052", Nper)
	RegisterFunction("\u004e\u0050\u0056", Npv)
	RegisterFunction("\u004fD\u0044\u004c\u0050\u0052\u0049\u0043E", Oddlprice)
	RegisterFunction("\u004fD\u0044\u004c\u0059\u0049\u0045\u004cD", Oddlyield)
	RegisterFunction("\u0050D\u0055\u0052\u0041\u0054\u0049\u004fN", Pduration)
	RegisterFunction("\u005fx\u006cf\u006e\u002e\u0050\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e", Pduration)
	RegisterFunction("\u0050\u004d\u0054", Pmt)
	RegisterFunction("\u0050\u0050\u004d\u0054", Ppmt)
	RegisterFunction("\u0050\u0052\u0049C\u0045", Price)
	RegisterFunction("\u0050R\u0049\u0043\u0045\u0044\u0049\u0053C", Pricedisc)
	RegisterFunction("\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054", Pricemat)
	RegisterFunction("\u0050\u0056", Pv)
	RegisterFunction("\u0052\u0041\u0054\u0045", Rate)
	RegisterFunction("\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044", Received)
	RegisterFunction("\u0052\u0052\u0049", Rri)
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0052\u0052I", Rri)
	RegisterFunction("\u0053\u004c\u004e", Sln)
	RegisterFunction("\u0053\u0059\u0044", Syd)
	RegisterFunction("\u0054B\u0049\u004c\u004c\u0045\u0051", Tbilleq)
	RegisterFunction("\u0054\u0042\u0049\u004c\u004c\u0050\u0052\u0049\u0043\u0045", Tbillprice)
	RegisterFunction("\u0054\u0042\u0049\u004c\u004c\u0059\u0049\u0045\u004c\u0044", Tbillyield)
	RegisterFunction("\u0056\u0044\u0042", Vdb)
	RegisterFunction("\u0058\u0049\u0052\u0052", Xirr)
	RegisterFunction("\u0058\u004e\u0050\u0056", Xnpv)
	RegisterFunction("\u0059\u0049\u0045L\u0044", Yield)
	RegisterFunction("\u0059I\u0045\u004c\u0044\u0044\u0049\u0053C", Yielddisc)
	RegisterFunction("\u0059\u0049\u0045\u004c\u0044\u004d\u0041\u0054", Yieldmat)
}

// IsNA is an implementation of the Excel ISNA() function.
func IsNA(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u004e\u0041\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeError && args[0].ValueString == "\u0023\u004e\u002f\u0041")
}
func _cafa(_geeb Result) *criteriaParsed {
	_gceff := _geeb.Type == ResultTypeNumber
	_dadg := _geeb.ValueNumber
	_ggcga := _ga.ToLower(_geeb.ValueString)
	_effc := _daecb(_ggcga)
	return &criteriaParsed{_gceff, _dadg, _ggcga, _effc}
}

// Choose implements the Excel CHOOSE function.
func Choose(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u0043\u0048O\u004f\u0053\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	_fcgbd := args[0]
	if _fcgbd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0043H\u004f\u004fS\u0045\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020f\u0069\u0072\u0073\u0074\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074y\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_fabb := int(_fcgbd.ValueNumber)
	if _fabb < 1 {
		return MakeErrorResult("\u0049\u006e\u0064\u0065\u0078\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u0061 \u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0076\u0061\u006c\u0075\u0065")
	}
	if len(args) <= _fabb {
		return MakeErrorResult("\u0049\u006e\u0064\u0065\u0078\u0020\u0073\u0068\u006f\u0075\u006cd\u0020\u0062\u0065\u0020\u006c\u0065\u0073\u0073 \u006fr\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0076\u0061\u006c\u0075\u0065\u0073")
	}
	return args[_fabb]
}

// Yield implements the Excel YIELD function.
func Yield(args []Result) Result {
	_dfeb := len(args)
	if _dfeb != 6 && _dfeb != 7 {
		return MakeErrorResult("\u0059\u0049E\u004c\u0044\u0020\u0072e\u0071\u0075i\u0072\u0065\u0073\u0020\u0073\u0069\u0078\u0020o\u0072\u0020\u0073\u0065\u0076\u0065\u006e\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_adfe, _ddce, _cadac := _gfcg(args[0], args[1], "\u0059\u0049\u0045L\u0044")
	if _cadac.Type == ResultTypeError {
		return _cadac
	}
	_egdfd := args[2]
	if _egdfd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045LD\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0072a\u0074e\u0020o\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_fged := _egdfd.ValueNumber
	if _fged < 0 {
		return MakeErrorResultType(ErrorTypeNum, "R\u0061\u0074\u0065\u0020\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006e\u006fn\u0020\u006e\u0065g\u0061t\u0069\u0076\u0065")
	}
	_agcd := args[3]
	if _agcd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020p\u0072 \u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ggbcb := _agcd.ValueNumber
	if _ggbcb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "p\u0072\u0020\u0073\u0068ou\u006cd\u0020\u0062\u0065\u0020\u0070o\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	_agcda := args[4]
	if _agcda.Type != ResultTypeNumber {
		return MakeErrorResult("Y\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065m\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0074yp\u0065\u0020\u006eu\u006db\u0065\u0072")
	}
	_becd := _agcda.ValueNumber
	if _becd < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006cd\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u006eo\u006e\u0020\u006e\u0065\u0067\u0061\u0074i\u0076\u0065")
	}
	_fcgg := args[5]
	if _fcgg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0066\u0072\u0065\u0071\u0075e\u006e\u0063\u0079\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_ecb := float64(int(_fcgg.ValueNumber))
	if !_dddg(_ecb) {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_edee := 0
	if _dfeb == 7 && args[6].Type != ResultTypeEmpty {
		_agb := args[6]
		if _agb.Type != ResultTypeNumber {
			return MakeErrorResult("Y\u0049\u0045\u004c\u0044\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u0062\u0061\u0073\u0069\u0073\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062e\u0072")
		}
		_edee = int(_agb.ValueNumber)
		if !_gfe(_edee) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063o\u0072\u0072\u0065\u0063t\u0020\u0062\u0061\u0073\u0069\u0073\u0020v\u0061\u006c\u0075\u0065\u0020\u0066\u006f\u0072\u0020\u0059\u0049\u0045\u004c\u0044")
		}
	}
	_ceffb := 0.0
	_bbab := 0.0
	_gdce := 1.0
	_egbdb, _cadac := _cgce(_adfe, _ddce, _fged, _bbab, _becd, _ecb, _edee)
	if _cadac.Type == ResultTypeError {
		return _cadac
	}
	_aefc, _cadac := _cgce(_adfe, _ddce, _fged, _gdce, _becd, _ecb, _edee)
	if _cadac.Type == ResultTypeError {
		return _cadac
	}
	_feae := (_gdce - _bbab) * 0.5
	for _daedfe := 0; _daedfe < 100 && _ceffb != _ggbcb; _daedfe++ {
		_ceffb, _cadac = _cgce(_adfe, _ddce, _fged, _feae, _becd, _ecb, _edee)
		if _cadac.Type == ResultTypeError {
			return _cadac
		}
		if _ggbcb == _egbdb {
			return MakeNumberResult(_bbab)
		} else if _ggbcb == _aefc {
			return MakeNumberResult(_gdce)
		} else if _ggbcb == _ceffb {
			return MakeNumberResult(_feae)
		} else if _ggbcb < _aefc {
			_gdce *= 2.0
			_aefc, _cadac = _cgce(_adfe, _ddce, _fged, _gdce, _becd, _ecb, _edee)
			if _cadac.Type == ResultTypeError {
				return _cadac
			}
			_feae = (_gdce - _bbab) * 0.5
		} else {
			if _ggbcb < _ceffb {
				_bbab = _feae
				_egbdb = _ceffb
			} else {
				_gdce = _feae
				_aefc = _ceffb
			}
			_feae = _gdce - (_gdce-_bbab)*((_ggbcb-_aefc)/(_egbdb-_aefc))
		}
	}
	return MakeNumberResult(_feae)
}
func _def(_gdf, _aaab float64, _acagd, _fcg int) float64 {
	_ggb := _bgee(_gdf)
	_facg := _bgee(_aaab)
	if _fcg == 1 {
		_fgff := _ccefc(_ggb, _facg, _acagd, 1)
		_cfaf := _fgff.AddDate(0, 12/_acagd, 0)
		return _cbea(_fgff, _cfaf, _fcg)
	}
	return float64(_fbeb(0, _fcg)) / float64(_acagd)
}
func (_bebfe *yyParserImpl) Lookahead() int { return _bebfe._gaeec }

const _cdfcd = 57348
const _ccag = 57344

// Ipmt implements the Excel IPMT function.
func Ipmt(args []Result) Result {
	_afeb := len(args)
	if _afeb < 4 || _afeb > 6 {
		return MakeErrorResult("\u0049P\u004d\u0054\u0020\u0072\u0065\u0071\u0075ir\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074s\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065n\u0020\u0066ou\u0072\u0020\u0061n\u0064\u0020\u0073\u0069\u0078")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("I\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_dfeg := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0050\u004dT\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_acdf := args[1].ValueNumber
	if _acdf <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u0050\u004d\u0054\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006fd\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_dbac := args[2].ValueNumber
	if _dbac <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062er\u0020o\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f \u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0050\u004d\u0054\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_fcdf := args[3].ValueNumber
	_gdaa := 0.0
	if _afeb > 4 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0049\u0050\u004d\u0054\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0075\u0074\u0075\u0072\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_gdaa = args[4].ValueNumber
	}
	_ada := 0
	if _afeb == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("I\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_ada = int(args[5].ValueNumber)
		if _ada != 0 {
			_ada = 1
		}
	}
	return MakeNumberResult(_dagb(_dfeg, _acdf, _dbac, _fcdf, _gdaa, _ada))
}
func _afdcc(_fcgef []Result, _fcfdf rmode) Result {
	if len(_fcgef) != 2 {
		return MakeErrorResult("\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_abdc := _fcgef[0].AsNumber()
	if _abdc.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_adda := _fcgef[1].AsNumber()
	if _adda.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020a\u0072\u0067\u0075m\u0065\u006e\u0074\u0020t\u006f\u0020\u0052\u004f\u0055\u004e\u0044\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_deef := _adda.ValueNumber
	_afdccb := _abdc.ValueNumber
	_fgded := 1.0
	if _deef > 0 {
		_fgded = _cc.Pow(1/10.0, _deef)
	} else {
		_fgded = _cc.Pow(10.0, -_deef)
	}
	_afdccb, _ggdg := _cc.Modf(_afdccb / _fgded)
	switch _fcfdf {
	case _afbac:
		const _fddc = 0.499999999
		if _ggdg >= _fddc {
			_afdccb++
		} else if _ggdg <= -_fddc {
			_afdccb--
		}
	case _accd:
	case _gcfgg:
		if _ggdg > 0 {
			_afdccb++
		} else if _ggdg < 0 {
			_afdccb--
		}
	}
	return MakeNumberResult(_afdccb * _fgded)
}

// FactDouble is an implementation of the excel FACTDOUBLE function which
// returns the double factorial of a positive numeric input.
func FactDouble(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0046\u0041C\u0054\u0044\u004f\u0055\u0042\u004c\u0045\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cbecd := args[0].AsNumber()
	if _cbecd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0041C\u0054\u0044\u004f\u0055\u0042\u004c\u0045\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _cbecd.ValueNumber < 0 {
		return MakeErrorResult("\u0046A\u0043\u0054D\u004f\u0055\u0042\u004cE\u0028\u0029\u0020a\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u006f\u006ely\u0020\u0070\u006fs\u0069\u0074i\u0076\u0065\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_defbf := float64(1)
	_aada := _cc.Trunc(_cbecd.ValueNumber)
	for _gdegb := _aada; _gdegb > 1; _gdegb -= 2 {
		_defbf *= _gdegb
	}
	return MakeNumberResult(_defbf)
}

// MinIfs implements the MINIFS function.
func MinIfs(args []Result) Result {
	_fbcd := _eged(args, true, "\u004d\u0049\u004e\u0049\u0046\u0053")
	if _fbcd.Type != ResultTypeEmpty {
		return _fbcd
	}
	_afda := _ecbg(args[1:])
	_bbdg := _cc.MaxFloat64
	_adbdf := _aadag(args[0])
	for _, _adefb := range _afda {
		_cfbc := _adbdf[_adefb._aaba][_adefb._gbfb].ValueNumber
		if _bbdg > _cfbc {
			_bbdg = _cfbc
		}
	}
	if _bbdg == _cc.MaxFloat64 {
		_bbdg = 0
	}
	return MakeNumberResult(float64(_bbdg))
}

const _cbcfe = 57370

// NewBool constructs a new boolean expression.
func NewBool(v string) Expression {
	_beb, _dff := _afc.ParseBool(v)
	if _dff != nil {
		_cf.Log.Debug("\u0065\u0072\u0072\u006f\u0072\u0020p\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u0075\u006ca\u0020\u0062\u006f\u006f\u006c\u0020\u0025s\u003a\u0020\u0025\u0076", v, _dff)
	}
	return Bool{_aef: _beb}
}

const (
	_fagga countMode = iota
	_cbfbc
	_cdgce
)

func _deece(_dbea Result, _fbgg, _abcaf string) (string, Result) {
	switch _dbea.Type {
	case ResultTypeString, ResultTypeNumber, ResultTypeEmpty:
		return _dbea.Value(), _feb
	default:
		return "", MakeErrorResult(_fbgg + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020" + _abcaf + "\u0020t\u006f\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006db\u0065r\u0020o\u0072\u0020\u0073\u0074\u0072\u0069\u006eg")
	}
}

// Reference returns a string reference value to an expression with prefix.
func (_ddcf PrefixExpr) Reference(ctx Context, ev Evaluator) Reference {
	_bfcf := _ddcf._dfgab.Reference(ctx, ev)
	_cbgdf := _ddcf._bcada.Reference(ctx, ev)
	if _bfcf.Type == ReferenceTypeSheet && _cbgdf.Type == ReferenceTypeCell {
		return Reference{Type: ReferenceTypeCell, Value: _bfcf.Value + "\u0021" + _cbgdf.Value}
	}
	return ReferenceInvalid
}
func _gfe(_cdb int) bool { return _cdb >= 0 && _cdb <= 4 }

// NewPrefixRangeExpr constructs a new range with prefix.
func NewPrefixRangeExpr(pfx, from, to Expression) Expression {
	_cbag, _baga, _bgbf := _beacf(from, to)
	if _bgbf != nil {
		_cf.Log.Debug(_bgbf.Error())
		return NewError(_bgbf.Error())
	}
	return PrefixRangeExpr{_beee: pfx, _bbbbge: _cbag, _affc: _baga}
}

// PrefixRangeExpr is a range expression that when evaluated returns a list of Results from a given sheet like Sheet1!A1:B4 (all cells from A1 to B4 from a sheet 'Sheet1').
type PrefixRangeExpr struct{ _beee, _bbbbge, _affc Expression }

// NA is an implementation of the Excel NA() function that just returns the #N/A! error.
func NA(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("\u004eA\u0028\u0029\u0020\u0061c\u0063\u0065\u0070\u0074\u0073 \u006eo\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074s")
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}
func _gbgg(_edac int) bool { return _edac == 0 || _edac == 4 }
func _dgge(_feadb Result, _beed *criteriaParsed) bool {
	_caffc := _ga.ToLower(_feadb.ValueString)
	_gaba := _beed._bgafc._afag
	_gcfdc := _beed._bgafc._gbcfa
	if _gaba == _bfcd {
		return _caffc == _gcfdc || _cd.Match(_gcfdc, _caffc)
	}
	if _feadb.Type != ResultTypeEmpty {
		if _caffc == _beed._abgb || _cd.Match(_beed._abgb, _caffc) {
			return true
		}
		if _, _ccgf := _afc.ParseFloat(_gcfdc, 64); _ccgf == nil {
			return false
		}
		switch _gaba {
		case _eggf:
			return _caffc <= _gcfdc
		case _edfe:
			return _caffc >= _gcfdc
		case _cdbac:
			return _caffc < _gcfdc
		case _dfef:
			return _caffc > _gcfdc
		}
	}
	return false
}

// Min is an implementation of the Excel MIN() function.
func Min(args []Result) Result { return _bdcee(args, false) }

// Reference returns an invalid reference for ConstArrayExpr.
func (_cef ConstArrayExpr) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// GetLocked returns FALSE for the invalid reference context.
func (_cbbbd *ivr) GetLocked(cellRef string) bool { return false }

var _afcc int64 = _gadd(1900, _eb.January, 1)

// SheetPrefixExpr is a reference to a sheet like Sheet1! (reference to sheet 'Sheet1').
type SheetPrefixExpr struct{ _bgfda string }

// Nominal implements the Excel NOMINAL function.
func Nominal(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u004e\u004f\u004d\u0049\u004e\u0041\u004c\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("NO\u004d\u0049N\u0041\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u006f\u006d\u0069\u006e\u0061\u006c\u0020\u0069\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062e\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072g\u0075m\u0065\u006et")
	}
	_bbecg := args[0].ValueNumber
	if _bbecg <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u004e\u004fM\u0049\u004e\u0041\u004c\u0020r\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0065\u0066\u0066\u0065\u0063\u0074\u0020\u0069\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u004e\u004f\u004d\u0049\u004e\u0041\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006d\u0070\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0070\u0065\u0072i\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_deaad := float64(int(args[1].ValueNumber))
	if _deaad < 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u004e\u004f\u004d\u0049\u004e\u0041\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006eum\u0062e\u0072\u0020\u006f\u0066\u0020\u0063\u006f\u006d\u0070\u006f\u0075\u006ed\u0069\u006e\u0067\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065 \u0031\u0020\u006f\u0072\u0020\u006d\u006f\u0072\u0065")
	}
	return MakeNumberResult((_cc.Pow(_bbecg+1, 1/_deaad) - 1) * _deaad)
}

type Reference struct {
	Type  ReferenceType
	Value string
}

// Negate is a negate expression like -A1.
type Negate struct{ _effb Expression }

const _cecbf = 57368
const _agbc = 57366

func _geffd(_effee []Result) (bool, Result) {
	for _, _dafg := range _effee {
		if _dafg.Type == ResultTypeError {
			return true, _dafg
		}
	}
	return false, MakeEmptyResult()
}

// Degrees is an implementation of the Excel function DEGREES() that converts
// radians to degrees.
func Degrees(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0044\u0045\u0047R\u0045\u0045\u0053\u0028)\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ddfde := args[0].AsNumber()
	if _ddfde.Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0045\u0047RE\u0045\u0053\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072e\u0073 \u006eu\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(180.0 / _cc.Pi * _ddfde.ValueNumber)
}

var _gcafc = [...]uint8{0, 17, 33, 49, 63, 78, 93, 108}

// Round is an implementation of the Excel ROUND function that rounds a number
// to a specified number of digits.
func Round(args []Result) Result { return _afdcc(args, _afbac) }

var _ebee []byte = []byte{0, 1, 2, 1, 11, 1, 12, 1, 13, 1, 14, 1, 15, 1, 16, 1, 17, 1, 18, 1, 19, 1, 20, 1, 21, 1, 22, 1, 23, 1, 24, 1, 25, 1, 26, 1, 27, 1, 28, 1, 29, 1, 30, 1, 31, 1, 32, 1, 33, 1, 34, 1, 35, 1, 36, 1, 37, 1, 38, 1, 39, 1, 40, 1, 41, 1, 42, 1, 43, 2, 0, 1, 2, 3, 4, 2, 3, 5, 2, 3, 6, 2, 3, 7, 2, 3, 8, 2, 3, 9, 2, 3, 10}

const (
	_afbac rmode = iota
	_accd
	_gcfgg
)

var _eaag = []ri{{1000, "\u004d"}, {995, "\u0056\u004d"}, {990, "\u0058\u004d"}, {950, "\u004c\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {495, "\u0056\u0044"}, {490, "\u0058\u0044"}, {450, "\u004c\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {99, "\u0049\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {45, "\u0056\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}

// Match implements the MATCH function.
func Match(args []Result) Result {
	_ccgd := len(args)
	if _ccgd != 2 && _ccgd != 3 {
		return MakeErrorResult("\u004d\u0041T\u0043\u0048\u0020\u0072e\u0071\u0075i\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020o\u0072\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_ddedc := 1
	if _ccgd == 3 && args[2].Type != ResultTypeEmpty {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("\u004d\u0041\u0054\u0043\u0048\u0020\u0072\u0065q\u0075\u0069\u0072es\u0020\u0074\u0068\u0065\u0020\u0074h\u0069\u0072\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006db\u0065\u0072")
		}
		_edeg := args[2].ValueNumber
		if _edeg == -1 || _edeg == 0 {
			_ddedc = int(_edeg)
		}
	}
	_dbcgb := args[1]
	var _ceaae []Result
	switch _dbcgb.Type {
	case ResultTypeList:
		_ceaae = _dbcgb.ValueList
	case ResultTypeArray:
		_cafc := _dbcgb.ValueArray
		for _, _gbge := range _cafc {
			if len(_gbge) != 1 {
				return MakeErrorResult("\u004d\u0041\u0054\u0043\u0048\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068e\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006f\u006e\u0065\u002dd\u0069\u006d\u0065\u006e\u0073\u0069o\u006e\u0061l\u0020\u0072a\u006eg\u0065")
			}
			_ceaae = append(_ceaae, _gbge[0])
		}
	default:
		return MakeErrorResult("\u004d\u0041\u0054\u0043\u0048\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068e\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006f\u006e\u0065\u002dd\u0069\u006d\u0065\u006e\u0073\u0069o\u006e\u0061l\u0020\u0072a\u006eg\u0065")
	}
	_agff := _cafa(args[0])
	switch _ddedc {
	case 0:
		for _fbge, _egag := range _ceaae {
			if _fedaf(_egag, _agff) {
				return MakeNumberResult(float64(_fbge + 1))
			}
		}
	case -1:
		for _efdg := 0; _efdg < len(_ceaae); _efdg++ {
			if _fedaf(_ceaae[_efdg], _agff) {
				return MakeNumberResult(float64(_efdg + 1))
			}
			if _agff._ceeeb && (_ceaae[_efdg].ValueNumber < _agff._adag) {
				if _efdg == 0 {
					return MakeErrorResultType(ErrorTypeNA, "")
				}
				return MakeNumberResult(float64(_efdg))
			}
		}
	case 1:
		for _bddb := 0; _bddb < len(_ceaae); _bddb++ {
			if _fedaf(_ceaae[_bddb], _agff) {
				return MakeNumberResult(float64(_bddb + 1))
			}
			if _agff._ceeeb && (_ceaae[_bddb].ValueNumber > _agff._adag) {
				if _bddb == 0 {
					return MakeErrorResultType(ErrorTypeNA, "")
				}
				return MakeNumberResult(float64(_bddb))
			}
		}
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}

// Reference returns a string reference value to a range with prefix.
func (_efag PrefixRangeExpr) Reference(ctx Context, ev Evaluator) Reference {
	_aege := _efag._beee.Reference(ctx, ev)
	_ebdac := _efag._bbbbge.Reference(ctx, ev)
	_ecec := _efag._affc.Reference(ctx, ev)
	if _aege.Type == ReferenceTypeSheet && _ebdac.Type == ReferenceTypeCell && _ecec.Type == ReferenceTypeCell {
		return MakeRangeReference(_cdbdg(_aege, _ebdac, _ecec))
	}
	return ReferenceInvalid
}
func _ecbg(_adee []Result) []rangeIndex {
	_bbaac := []rangeIndex{}
	_cgde := len(_adee)
	for _fddd := 0; _fddd < _cgde-1; _fddd += 2 {
		_gbfe := []rangeIndex{}
		_bade := _aadag(_adee[_fddd])
		_cfda := _cafa(_adee[_fddd+1])
		if _fddd == 0 {
			for _gagb, _caea := range _bade {
				for _dfbc, _aadc := range _caea {
					if _addg(_aadc, _cfda) {
						_gbfe = append(_gbfe, rangeIndex{_gagb, _dfbc})
					}
				}
			}
		} else {
			for _, _fdega := range _bbaac {
				_feacd := _bade[_fdega._aaba][_fdega._gbfb]
				if _addg(_feacd, _cfda) {
					_gbfe = append(_gbfe, _fdega)
				}
			}
		}
		if len(_gbfe) == 0 {
			return []rangeIndex{}
		}
		_bbaac = _gbfe[:]
	}
	return _bbaac
}

// Value returns a string version of the result.
func (_gadcf Result) Value() string {
	switch _gadcf.Type {
	case ResultTypeNumber:
		_ebcc := _afc.FormatFloat(_gadcf.ValueNumber, 'f', -1, 64)
		if len(_ebcc) > 12 {
			_ebbef := 12
			for _fgegc := _ebbef; _fgegc > 0 && _ebcc[_fgegc] == '0'; _fgegc-- {
				_ebbef--
			}
			_ebcc = _ebcc[0 : _ebbef+1]
		}
		return _ebcc
	case ResultTypeError:
		return _gadcf.ValueString
	case ResultTypeString:
		return _gadcf.ValueString
	case ResultTypeList:
		if len(_gadcf.ValueList) == 0 {
			return ""
		}
		return _gadcf.ValueList[0].Value()
	case ResultTypeArray:
		if len(_gadcf.ValueArray) == 0 || len(_gadcf.ValueArray[0]) == 0 {
			return ""
		}
		return _gadcf.ValueArray[0][0].Value()
	case ResultTypeEmpty:
		return ""
	default:
		return "\u0075\u006e\u0068\u0061nd\u006c\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0076\u0061\u006cu\u0065"
	}
}

const _bcfe = _eb.Millisecond * 1000

func _fbeb(_edfd, _aad int) int {
	switch _aad {
	case 1:
		if _bee(_edfd) {
			return 366
		} else {
			return 365
		}
	case 3:
		return 365
	default:
		return 360
	}
}

const _bggae = 57346

// Amordegrc implements the Excel AMORDEGRC function.
func Amordegrc(args []Result) Result {
	_egfg, _bec := _ffee(args, "\u0041M\u004f\u0052\u0044\u0045\u0047\u0052C")
	if _bec.Type == ResultTypeError {
		return _bec
	}
	_ggc := _egfg._dffd
	_ebgf := _egfg._dca
	_gaddd := _egfg._geg
	_acaa := _egfg._abge
	_debb := _egfg._dfaaf
	_gfga := _egfg._afed
	if _gfga >= 0.5 {
		return MakeErrorResultType(ErrorTypeNum, "\u0041\u004d\u004f\u0052\u0044\u0045\u0047R\u0043\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u0030\u002e\u0035")
	}
	_egg := _egfg._cbdf
	_bdcea := 1.0 / _gfga
	_fdeg := 2.5
	if _bdcea < 3 {
		_fdeg = 1
	} else if _bdcea < 5 {
		_fdeg = 1.5
	} else if _bdcea <= 6 {
		_fdeg = 2
	}
	_gfga *= _fdeg
	_gbaf, _adgd := _fcfa(_ebgf, _gaddd, _egg)
	if _adgd.Type == ResultTypeError {
		return MakeErrorResult("\u0069\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0064\u0061\u0074\u0065\u0073\u0020\u0066\u006f\u0072\u0020\u0041\u004d\u004f\u0052\u0044EG\u0052\u0043")
	}
	_aaad := _fcca(_gbaf * _gfga * _ggc)
	_ggc -= _aaad
	_bde := _ggc - _acaa
	for _ace := 0; _ace < _debb; _ace++ {
		_aaad = _fcca(_gfga * _ggc)
		_bde -= _aaad
		if _bde < 0 {
			switch _debb - _ace {
			case 0:
			case 1:
				return MakeNumberResult(_fcca(_ggc * 0.5))
			default:
				return MakeNumberResult(0)
			}
		}
		_ggc -= _aaad
	}
	return MakeNumberResult(_aaad)
}
func _gadd(_dffb int, _eab _eb.Month, _gafe int) int64 {
	if _dffb == 1900 && int(_eab) <= 2 {
		_gafe--
	}
	_dcc := _eb.Date(_dffb, _eab, _gafe, 0, 0, 0, 0, _eb.UTC)
	return _dcc.Unix()
}
func MakeRangeReference(ref string) Reference { return Reference{Type: ReferenceTypeRange, Value: ref} }
func _bceb(_abbbd []Result, _ggea bool) Result {
	_ebbfg := "\u004d\u0041\u0058"
	if _ggea {
		_ebbfg = "\u004d\u0041\u0058\u0041"
	}
	if len(_abbbd) == 0 {
		return MakeErrorResult(_ebbfg + "\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s\u0020\u0061\u0074\u0020\u006c\u0065\u0061s\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_feded := -_cc.MaxFloat64
	for _, _becdd := range _abbbd {
		switch _becdd.Type {
		case ResultTypeNumber:
			if (_ggea || !_becdd.IsBoolean) && _becdd.ValueNumber > _feded {
				_feded = _becdd.ValueNumber
			}
		case ResultTypeList, ResultTypeArray:
			_ffbe := _bceb(_becdd.ListValues(), _ggea)
			if _ffbe.ValueNumber > _feded {
				_feded = _ffbe.ValueNumber
			}
		case ResultTypeEmpty:
		case ResultTypeString:
			_aaebe := 0.0
			if _ggea {
				_aaebe = _becdd.AsNumber().ValueNumber
			}
			if _aaebe > _feded {
				_feded = _aaebe
			}
		default:
			_cf.Log.Debug("\u0075\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020"+_ebbfg+"\u0028\u0029\u0020\u0061rg\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _becdd.Type)
		}
	}
	if _feded == -_cc.MaxFloat64 {
		_feded = 0
	}
	return MakeNumberResult(_feded)
}

// Eval evaluates and returns the result of a constant array expression.
func (_fff ConstArrayExpr) Eval(ctx Context, ev Evaluator) Result {
	_gcb := [][]Result{}
	for _, _adg := range _fff._ba {
		_cag := []Result{}
		for _, _abe := range _adg {
			_cag = append(_cag, _abe.Eval(ctx, ev))
		}
		_gcb = append(_gcb, _cag)
	}
	return MakeArrayResult(_gcb)
}

type criteriaRegex struct {
	_afag  byte
	_gbcfa string
}

const _ggded = 57364

func Sign(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0053\u0049\u0047\u004e(\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_aadf := args[0].AsNumber()
	if _aadf.Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0049\u0047N(\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020a\u0020n\u0075m\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _aadf.ValueNumber < 0 {
		return MakeNumberResult(-1)
	} else if _aadf.ValueNumber > 0 {
		return MakeNumberResult(1)
	}
	return MakeNumberResult(0)
}

// Radians is an implementation of the Excel function RADIANS() that converts
// degrees to radians.
func Radians(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0052\u0041\u0044I\u0041\u004e\u0053\u0028)\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cgafe := args[0].AsNumber()
	if _cgafe.Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0041\u0044IA\u004e\u0053\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072e\u0073 \u006eu\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(_cc.Pi / 180.0 * _cgafe.ValueNumber)
}

// Combina is an implementation of the Excel COMBINA function whic returns the
// number of combinations with repetitions.
func Combina(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0043\u004f\u004dB\u0049\u004e\u0041\u0028)\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gaddc := args[0].AsNumber()
	_feca := args[1].AsNumber()
	if _gaddc.Type != ResultTypeNumber || _feca.Type != ResultTypeNumber {
		return MakeErrorResult("\u0043\u004fMB\u0049\u004e\u0041(\u0029\u0020\u0072\u0065qui\u0072es\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_gcebd := _cc.Trunc(_gaddc.ValueNumber)
	_fgbba := _cc.Trunc(_feca.ValueNumber)
	if _gcebd < _fgbba {
		return MakeErrorResult("\u0043O\u004d\u0042\u0049\u004e\u0041\u0028\u0029\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u006e\u0020\u003e\u0020\u006b")
	}
	if _gcebd == 0 {
		return MakeNumberResult(0)
	}
	args[0] = MakeNumberResult(_gcebd + _fgbba - 1)
	args[1] = MakeNumberResult(_gcebd - 1)
	return Combin(args)
}

// LastEvalIsRef returns if last evaluation with the evaluator was a reference.
func (_bfd *defEval) LastEvalIsRef() bool { return _bfd._fdc }

var (
	_cegc = 0
	_bdff = false
)

func _feab(_bfdd yyLexer, _ggcec *yySymType) (_dfbgf, _eeegf int) {
	_eeegf = 0
	_dfbgf = _bfdd.Lex(_ggcec)
	if _dfbgf <= 0 {
		_eeegf = _bcdab[0]
		goto _cedgc
	}
	if _dfbgf < len(_bcdab) {
		_eeegf = _bcdab[_dfbgf]
		goto _cedgc
	}
	if _dfbgf >= _ccag {
		if _dfbgf < _ccag+len(_accgc) {
			_eeegf = _accgc[_dfbgf-_ccag]
			goto _cedgc
		}
	}
	for _aaafb := 0; _aaafb < len(_efgc); _aaafb += 2 {
		_eeegf = _efgc[_aaafb+0]
		if _eeegf == _dfbgf {
			_eeegf = _efgc[_aaafb+1]
			goto _cedgc
		}
	}
_cedgc:
	if _eeegf == 0 {
		_eeegf = _accgc[1]
	}
	if _cegc >= 3 {
		_e.Printf("l\u0065\u0078\u0020\u0025\u0073\u0028\u0025\u0064\u0029\u000a", _dgda(_eeegf), uint(_dfbgf))
	}
	return _dfbgf, _eeegf
}

// Product is an implementation of the Excel PRODUCT() function.
func Product(args []Result) Result {
	_efafe := 1.0
	for _, _gabb := range args {
		_gabb = _gabb.AsNumber()
		switch _gabb.Type {
		case ResultTypeNumber:
			_efafe *= _gabb.ValueNumber
		case ResultTypeList, ResultTypeArray:
			_bfbea := Product(_gabb.ListValues())
			if _bfbea.Type != ResultTypeNumber {
				return _bfbea
			}
			_efafe *= _bfbea.ValueNumber
		case ResultTypeString:
		case ResultTypeError:
			return _gabb
		case ResultTypeEmpty:
		default:
			return MakeErrorResult(_e.Sprintf("\u0075\u006eha\u006e\u0064\u006ce\u0064\u0020\u0050\u0052ODU\u0043T(\u0029\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0020\u0074\u0079\u0070\u0065\u0020%\u0073", _gabb.Type))
		}
	}
	return MakeNumberResult(_efafe)
}
func (_fgbc *Lexer) emit(_ebbaa tokenType, _gfgfb []byte) {
	if _gafa {
		_e.Println("\u0065\u006d\u0069\u0074", _ebbaa, _cabg(string(_gfgfb)))
	}
	_fgbc._cgdbfa <- &node{_ebbaa, string(_gfgfb)}
}
func _dfa(_fea string, _acb *_eg.UpdateQuery) string {
	_fa, _cgf := _fb.ParseCellReference(_fea)
	if _cgf != nil {
		return "\u0023\u0052\u0045F\u0021"
	}
	if _acb.UpdateType == _eg.UpdateActionRemoveColumn {
		_aba := _acb.ColumnIdx
		_agg := _fa.ColumnIdx
		if _agg < _aba {
			return _fea
		} else if _agg == _aba {
			return "\u0023\u0052\u0045F\u0021"
		} else {
			return _fa.Update(_eg.UpdateActionRemoveColumn).String()
		}
	}
	return _fea
}

var _efgc = [...]int{0}

const _efaag = 57352

var _gbeb = [...]int{45, 3, 44, 32, 18, 40, 72, 46, 47, 30, 31, 32, 39, 48, 28, 29, 30, 31, 32, 75, 39, 49, 32, 56, 50, 70, 23, 39, 76, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 77, 71, 69, 54, 43, 13, 19, 21, 55, 82, 11, 78, 9, 74, 28, 29, 30, 31, 32, 37, 33, 34, 35, 36, 38, 1, 20, 39, 10, 2, 8, 0, 80, 79, 0, 0, 0, 83, 0, 81, 73, 28, 29, 30, 31, 32, 37, 33, 34, 35, 36, 38, 0, 0, 39, 28, 29, 30, 31, 32, 37, 33, 34, 35, 36, 38, 26, 27, 39, 51, 52, 25, 14, 15, 16, 17, 0, 24, 23, 22, 41, 23, 12, 0, 6, 7, 26, 27, 0, 42, 0, 25, 14, 15, 16, 17, 0, 24, 23, 22, 5, 0, 12, 0, 6, 7, 26, 27, 0, 4, 0, 25, 14, 15, 16, 17, 0, 24, 23, 22, 41, 0, 12, 53, 6, 7, 26, 27, 0, 0, 0, 25, 14, 15, 16, 17, 0, 24, 23, 22, 41, 0, 12, 0, 6, 7}

// Fv implements the Excel FV function.
func Fv(args []Result) Result {
	_aagd := len(args)
	if _aagd < 3 || _aagd > 5 {
		return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0033\u0020\u0061\u006e\u0064\u00205")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_fbdg := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_edfdg := args[1].ValueNumber
	if _edfdg != float64(int(_edfdg)) {
		return MakeErrorResultType(ErrorTypeNum, "\u0046\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0061\u0079\u006d\u0065\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_beg := args[2].ValueNumber
	_dcge := 0.0
	if _aagd >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("F\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0070\u0072\u0065\u0073\u0065\u006et \u0076\u0061\u006c\u0075e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_dcge = args[3].ValueNumber
	}
	_cgdf := 0
	if _aagd == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0046\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_cgdf = int(args[4].ValueNumber)
		if _cgdf != 0 {
			_cgdf = 1
		}
	}
	return MakeNumberResult(_ggbd(_fbdg, _edfdg, _beg, _dcge, _cgdf))
}

// Fact is an implementation of the excel FACT function which returns the
// factorial of a positive numeric input.
func Fact(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("F\u0041\u0043\u0054\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_beae := args[0].AsNumber()
	if _beae.Type != ResultTypeNumber {
		return MakeErrorResult("F\u0041\u0043\u0054\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	if _beae.ValueNumber < 0 {
		return MakeErrorResult("\u0046\u0041\u0043\u0054\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u006f\u006e\u006c\u0079\u0020\u0070\u006f\u0073\u0069t\u0069\u0076\u0065\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	return MakeNumberResult(_cfga(_beae.ValueNumber))
}

// GCD implements the Excel GCD() function which returns the greatest common
// divisor of a range of numbers.
func GCD(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0047\u0043D(\u0029\u0020\u0072e\u0071\u0075\u0069\u0072es \u0061t \u006c\u0065\u0061\u0073\u0074\u0020\u006fne\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_abfce := []float64{}
	for _, _agee := range args {
		switch _agee.Type {
		case ResultTypeString:
			_geed := _agee.AsNumber()
			if _geed.Type != ResultTypeNumber {
				return MakeErrorResult("\u0047\u0043D(\u0029\u0020\u006fn\u006c\u0079\u0020\u0061cce\u0070ts\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
			}
			_abfce = append(_abfce, _geed.ValueNumber)
		case ResultTypeList, ResultTypeArray:
			_fffe := GCD(_agee.ListValues())
			if _fffe.Type != ResultTypeNumber {
				return _fffe
			}
			_abfce = append(_abfce, _fffe.ValueNumber)
		case ResultTypeNumber:
			_abfce = append(_abfce, _agee.ValueNumber)
		case ResultTypeError:
			return _agee
		default:
			return MakeErrorResult(_e.Sprintf("\u0047\u0043\u0044()\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065d\u0020a\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _agee.Type))
		}
	}
	if _abfce[0] < 0 {
		return MakeErrorResult("\u0047\u0043D\u0028\u0029\u0020\u006fn\u006c\u0079 \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if len(_abfce) == 1 {
		return MakeNumberResult(_abfce[0])
	}
	_agggbf := _abfce[0]
	for _gedd := 1; _gedd < len(_abfce); _gedd++ {
		if _abfce[_gedd] < 0 {
			return MakeErrorResult("\u0047\u0043D\u0028\u0029\u0020\u006fn\u006c\u0079 \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
		}
		_agggbf = _ebggd(_agggbf, _abfce[_gedd])
	}
	return MakeNumberResult(_agggbf)
}

// Ispmt implements the Excel ISPMT function.
func Ispmt(args []Result) Result {
	if len(args) != 4 {
		return MakeErrorResult("\u0049\u0053P\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0050\u004d\u0054 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bebff := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0050\u004d\u0054\u0020\u0072e\u0071\u0075\u0069r\u0065\u0073\u0020\u0070e\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_dcac := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072\u0069o\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006dbe\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_faefe := args[2].ValueNumber
	if _faefe <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049S\u0050\u004d\u0054\u0020\u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020p\u0065\u0072i\u006f\u0064\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006eu\u006d\u0062er\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065s\u0065\u006e\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_caecd := args[3].ValueNumber
	return MakeNumberResult(_caecd * _bebff * (_dcac/_faefe - 1))
}

var _ebgdg = [...]int{0, -2, 1, 2, 0, 0, 0, 0, 11, 12, 13, 14, 0, 16, 5, 6, 7, 8, 22, 0, 24, 46, 0, 26, 25, 29, 30, 31, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 18, 20, 9, 10, 0, 0, 23, 32, 33, 47, 0, 49, 51, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 0, 17, 0, 0, 15, 27, 0, 48, 53, 4, 19, 21, 28, 50, 52}

// Small implements the Excel SMALL function.
func Small(args []Result) Result { return _gceb(args, false) }
func _ccced(_dabf Result) []Result {
	_ceg := _dabf.ValueList
	if _dabf.Type == ResultTypeArray {
		_ceg = nil
		for _, _fcba := range _dabf.ValueArray {
			if len(_fcba) > 0 {
				_ceg = append(_ceg, _fcba[0])
			} else {
				_ceg = append(_ceg, _feb)
			}
		}
	}
	return _ceg
}

// Arabic implements the Excel ARABIC function which parses roman numerals.  It
// accepts one numeric argument.
func Arabic(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0041\u0052\u0041\u0042I\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_feaab := args[0]
	switch _feaab.Type {
	case ResultTypeNumber, ResultTypeList, ResultTypeEmpty:
		return MakeErrorResult("\u0041\u0052\u0041B\u0049\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	case ResultTypeString:
		_gafdc := 0.0
		_dffcg := 0.0
		for _, _ebfde := range _feaab.ValueString {
			_defd := 0.0
			switch _ebfde {
			case 'I':
				_defd = 1
			case 'V':
				_defd = 5
			case 'X':
				_defd = 10
			case 'L':
				_defd = 50
			case 'C':
				_defd = 100
			case 'D':
				_defd = 500
			case 'M':
				_defd = 1000
			}
			_gafdc += _defd
			switch {
			case _dffcg == _defd && (_dffcg == 5 || _dffcg == 50 || _dffcg == 500):
				return MakeErrorResult("i\u006e\u0076\u0061\u006cid\u0020A\u0052\u0041\u0042\u0049\u0043 \u0066\u006f\u0072\u006d\u0061\u0074")
			case 2*_dffcg == _defd:
				return MakeErrorResult("i\u006e\u0076\u0061\u006cid\u0020A\u0052\u0041\u0042\u0049\u0043 \u0066\u006f\u0072\u006d\u0061\u0074")
			}
			if _dffcg < _defd {
				_gafdc -= 2 * _dffcg
			}
			_dffcg = _defd
		}
		return MakeNumberResult(_gafdc)
	case ResultTypeError:
		return _feaab
	default:
		return MakeErrorResult(_e.Sprintf("\u0075\u006e\u0068an\u0064\u006c\u0065\u0064\u0020\u0041\u0043\u004f\u0053H\u0028)\u0020a\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _feaab.Type))
	}
}

// Pricedisc implements the Excel PRICEDISC function.
func Pricedisc(args []Result) Result {
	_deabg := len(args)
	if _deabg != 4 && _deabg != 5 {
		return MakeErrorResult("\u0050\u0052\u0049\u0043\u0045D\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020f\u006f\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_geb, _cdbb, _eddd := _gfcg(args[0], args[1], "\u0050R\u0049\u0043\u0045\u0044\u0049\u0053C")
	if _eddd.Type == ResultTypeError {
		return _eddd
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052\u0049C\u0045\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0064\u0069\u0073\u0063\u006f\u0075\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079p\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_cab := args[2].ValueNumber
	if _cab <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049C\u0045\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0064\u0069\u0073\u0063\u006f\u0075\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065 \u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050R\u0049\u0043E\u0044\u0049\u0053\u0043 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065mp\u0074\u0069\u006fn\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020n\u0075\u006db\u0065\u0072")
	}
	_bfad := args[3].ValueNumber
	if _bfad <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050R\u0049\u0043E\u0044\u0049\u0053\u0043 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065mp\u0074\u0069\u006fn\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006fs\u0069\u0074i\u0076\u0065")
	}
	_fcdc := 0
	if _deabg == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0052I\u0043\u0045\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_fcdc = int(args[4].ValueNumber)
		if !_gfe(_fcdc) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0050R\u0049\u0043E\u0044\u0049\u0053\u0043")
		}
	}
	_feea, _eddd := _fcfa(_geb, _cdbb, _fcdc)
	if _eddd.Type == ResultTypeError {
		return _eddd
	}
	return MakeNumberResult(_bfad * (1 - _cab*_feea))
}

var _ggde, _fgbb, _cgfbe, _aebbb, _cgbg, _geabe, _bfggb, _gbcbc, _fbf, _aaed, _gaab, _ceef, _befb, _fffd, _cdbe *_f.Regexp

// Eval evaluates and returns the result of a Negate expression.
func (_begeb Negate) Eval(ctx Context, ev Evaluator) Result {
	_ababd := _begeb._effb.Eval(ctx, ev)
	if _ababd.Type == ResultTypeNumber {
		return MakeNumberResult(-_ababd.ValueNumber)
	}
	return MakeErrorResult("\u004e\u0045\u0047A\u0054\u0045\u0020\u0065x\u0070\u0065\u0063\u0074\u0065\u0064\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
}

// Reference returns an invalid reference for FunctionCall.
func (_adeg FunctionCall) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// Coupnum implements the Excel COUPNUM function.
func Coupnum(args []Result) Result {
	_egfeb, _acfc := _cebf(args, "\u0043O\u0055\u0050\u004e\u0055\u004d")
	if _acfc.Type == ResultTypeError {
		return _acfc
	}
	_bdgcb := _egfeb._gbf
	_gbcg := _egfeb._ceea
	_gefe, _acfc := _cefd(_egfeb._bbaa, _egfeb._aee, _bdgcb, _gbcg)
	if _acfc.Type == ResultTypeError {
		return _acfc
	}
	return MakeNumberResult(_gefe)
}

// IsDBCS returns false for the invalid reference context.
func (_gdeef *ivr) IsDBCS() bool { return false }

// Cell is an implementation of the Excel CELL function that returns information
// about the formatting, location, or contents of a cell.
func Cell(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 && len(args) != 2 {
		return MakeErrorResult("\u0043\u0045\u004cL \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020o\u006ee\u0020o\u0072 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ebfcc := args[0].AsString()
	if _ebfcc.Type != ResultTypeString {
		return MakeErrorResult("\u0043\u0045\u004c\u004c\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065 \u0073t\u0072\u0069\u006e\u0067")
	}
	_cdgg := "\u0041\u0031"
	if len(args) == 2 {
		_cdba := args[1].Ref
		if _cdba.Type != ReferenceTypeCell {
			return MakeErrorResult("\u0043\u0045\u004c\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064 \u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079p\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
		}
		_cdgg = _cdba.Value
	}
	switch _ebfcc.ValueString {
	case "\u0061d\u0064\u0072\u0065\u0073\u0073":
		_baae, _adbee := _fb.ParseCellReference(_cdgg)
		if _adbee != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _cdgg)
		}
		_cffd := "\u0024" + _baae.Column + "\u0024" + _afc.Itoa(int(_baae.RowIdx))
		if _baae.SheetName != "" {
			_cffd = _baae.SheetName + "\u0021" + _cffd
		}
		return MakeStringResult(_cffd)
	case "\u0063\u006f\u006c":
		_ceeee, _dfgf := _fb.ParseCellReference(_cdgg)
		if _dfgf != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _cdgg)
		}
		return MakeNumberResult(float64(_ceeee.ColumnIdx + 1))
	case "\u0063\u006f\u006co\u0072":
		_ggcd := _ga.Contains(ctx.GetFormat(_cdgg), "\u005b\u0052\u0045D\u005d")
		return MakeBoolResult(_ggcd)
	case "\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073":
		return args[1]
	case "\u0066\u0069\u006c\u0065\u006e\u0061\u006d\u0065":
		return MakeStringResult(ctx.GetFilename())
	case "\u0066\u006f\u0072\u006d\u0061\u0074":
		_dbca := "\u0047"
		_dfbb := ctx.GetFormat(_cdgg)
		if _dfbb == "\u0047e\u006e\u0065\u0072\u0061\u006c" || _ggde.MatchString(_dfbb) {
			_dbca = "\u0046\u0030"
		} else if _dfbb == "\u0030\u0025" {
			_dbca = "\u0050\u0030"
		} else if _dfbb == "\u004d\u004d\u004d\u0020\u0044\u0044" {
			_dbca = "\u0044\u0032"
		} else if _dfbb == "\u004d\u004d\u002fY\u0059" {
			_dbca = "\u0044\u0033"
		} else if _dfbb == "\u004d\u004d\u002f\u0044D/\u0059\u0059\u005c\u0020\u0048\u0048\u003a\u004d\u004d\u005c\u0020\u0041\u004d\u002fP\u004d" || _dfbb == "M\u004d/\u0044\u0044\u002f\u0059\u0059\u0059\u0059\u005c \u0048\u0048\u003a\u004dM:\u0053\u0053" {
			_dbca = "\u0044\u0034"
		} else if _dfbb == "\u004d\u004d\u005c\u002d\u0044\u0044" {
			_dbca = "\u0044\u0035"
		} else if _dfbb == "\u0048H\u003aM\u004d\u003a\u0053\u0053\u005c\u0020\u0041\u004d\u002f\u0050\u004d" {
			_dbca = "\u0044\u0036"
		} else if _dfbb == "\u0048\u0048\u003aM\u004d\u005c\u0020\u0041\u004d\u002f\u0050\u004d" {
			_dbca = "\u0044\u0037"
		} else if _dfbb == "\u0048\u0048\u003a\u004d\u004d\u003a\u0053\u0053" {
			_dbca = "\u0044\u0038"
		} else if _dfbb == "\u0048\u0048\u003aM\u004d" {
			_dbca = "\u0044\u0039"
		} else if _cgfbe.MatchString(_dfbb) {
			_dbca = "\u002e\u0030"
		} else if _aebbb.MatchString(_dfbb) {
			_dbca = "\u002e\u0030\u0028\u0029"
		} else if _aaed.MatchString(_dfbb) {
			_dbca = "\u0043\u0030"
		} else if _befb.MatchString(_dfbb) || _fffd.MatchString(_dfbb) {
			_dbca = "\u0044\u0031"
		} else if _bccd := _fgbb.FindStringSubmatch(_dfbb); len(_bccd) > 1 {
			_dbca = "\u0046" + _afc.Itoa(len(_bccd[1]))
		} else if _ddde := _geabe.FindStringSubmatch(_dfbb); len(_ddde) > 1 {
			_dbca = "\u002e" + _afc.Itoa(len(_ddde[2]))
		} else if _bagf := _fbf.FindStringSubmatch(_dfbb); len(_bagf) > 1 {
			_dbca = "\u0050" + _afc.Itoa(len(_bagf[2]))
		} else if _bgce := _gaab.FindStringSubmatch(_dfbb); len(_bgce) > 1 {
			_dbca = "\u0043" + _cccgb(_bgce, 1)
		} else if _eebef := _ceef.FindStringSubmatch(_dfbb); len(_eebef) > 1 {
			_dbca = "\u0043" + _cccgb(_eebef, 1)
		} else if _gbdf := _bfggb.FindStringSubmatch(_dfbb); len(_gbdf) > 1 {
			_dbca = "\u002e" + _cccgb(_gbdf, 1) + "\u0028\u0029"
		} else if _cdgc := _cgbg.FindStringSubmatch(_dfbb); len(_cdgc) > 1 {
			_dbca = "\u002e" + _cccgb(_cdgc, 1)
		} else if _dada := _cdbe.FindStringSubmatch(_dfbb); len(_dada) > 1 {
			_dbca = "\u0053" + _cccgb(_dada, 3)
		}
		if _dbca != "\u0047" && _ga.Contains(_dfbb, "\u005b\u0052\u0045D\u005d") {
			_dbca += "\u002d"
		}
		return MakeStringResult(_dbca)
	case "p\u0061\u0072\u0065\u006e\u0074\u0068\u0065\u0073\u0065\u0073":
		_bege := ctx.GetFormat(_cdgg)
		if _gbcbc.MatchString(_bege) {
			return MakeNumberResult(1)
		} else {
			return MakeNumberResult(0)
		}
	case "\u0070\u0072\u0065\u0066\u0069\u0078":
		return MakeStringResult(ctx.GetLabelPrefix(_cdgg))
	case "\u0070r\u006f\u0074\u0065\u0063\u0074":
		_dbfd := 0.0
		if ctx.GetLocked(_cdgg) {
			_dbfd = 1.0
		}
		return MakeNumberResult(_dbfd)
	case "\u0072\u006f\u0077":
		_acee, _agac := _fb.ParseCellReference(_cdgg)
		if _agac != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _cdgg)
		}
		return MakeNumberResult(float64(_acee.RowIdx))
	case "\u0074\u0079\u0070\u0065":
		switch args[1].Type {
		case ResultTypeEmpty:
			return MakeStringResult("\u0062")
		case ResultTypeString:
			return MakeStringResult("\u006c")
		default:
			return MakeStringResult("\u0076")
		}
	case "\u0077\u0069\u0064t\u0068":
		_dbedd, _daedd := _fb.ParseCellReference(_cdgg)
		if _daedd != nil {
			return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _cdgg)
		}
		if _dbedd.SheetName == "" {
			return MakeNumberResult(ctx.GetWidth(int(_dbedd.ColumnIdx)))
		} else {
			return MakeNumberResult(ctx.Sheet(_dbedd.SheetName).GetWidth(int(_dbedd.ColumnIdx)))
		}
	}
	return MakeErrorResult("\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0066\u0069\u0072\u0073t\u0020a\u0072g\u0075m\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0043\u0045\u004c\u004c\u003a\u0020" + _ebfcc.ValueString)
}

// DateDif is an implementation of the Excel DATEDIF() function.
func DateDif(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeString {
		return MakeErrorResult("\u0044\u0041\u0054\u0045\u0044I\u0046\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077o\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u006e\u0064\u0020\u006f\u006e\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_fba := args[0].ValueNumber
	_cfaa := args[1].ValueNumber
	if _cfaa < _fba {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0068\u0065\u0020\u0073\u0074\u0061r\u0074\u0020\u0064a\u0074\u0065\u0020\u0069s\u0020\u0067\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0074\u0068\u0065\u0020\u0065\u006e\u0064\u0020\u0064\u0061\u0074\u0065")
	}
	if _cfaa == _fba {
		return MakeNumberResult(0)
	}
	_aga := _ga.ToLower(args[2].ValueString)
	if _aga == "\u0064" {
		return MakeNumberResult(_cfaa - _fba)
	}
	_cec := _bgee(_fba)
	_gfc := _bgee(_cfaa)
	_dde, _ecg, _bdga := _cec.Date()
	_fggc, _dfe, _agc := _gfc.Date()
	_cga := int(_ecg)
	_ecd := int(_dfe)
	var _eeda float64
	switch _aga {
	case "\u0079":
		_eeda = float64(_fggc - _dde)
		if _ecd < _cga || (_ecd == _cga && _agc < _bdga) {
			_eeda--
		}
	case "\u006d":
		_fab := _fggc - _dde
		_edc := _ecd - _cga
		if _agc < _bdga {
			_edc--
		}
		if _edc < 0 {
			_fab--
			_edc += 12
		}
		_eeda = float64(_fab*12 + _edc)
	case "\u006d\u0064":
		_eae := _ecd
		if _agc < _bdga {
			_eae--
		}
		_eeda = float64(int(_cfaa - _bad(_fggc, _eae, _bdga)))
	case "\u0079\u006d":
		_eeda = float64(_ecd - _cga)
		if _agc < _bdga {
			_eeda--
		}
		if _eeda < 0 {
			_eeda += 12
		}
	case "\u0079\u0064":
		_eag := _fggc
		if _ecd < _cga || (_ecd == _cga && _agc < _bdga) {
			_eag--
		}
		_eeda = float64(int(_cfaa - _bad(_eag, _cga, _bdga)))
	default:
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0069\u006e\u0074e\u0072\u0076\u0061\u006c\u0020\u0076\u0061\u006c\u0075\u0065")
	}
	return MakeNumberResult(_eeda)
}

const (
	ReferenceTypeInvalid ReferenceType = iota
	ReferenceTypeCell
	ReferenceTypeHorizontalRange
	ReferenceTypeVerticalRange
	ReferenceTypeNamedRange
	ReferenceTypeRange
	ReferenceTypeSheet
)

func _fbcad(_aedab string, _gggc []Result) (*parsedReplaceObject, Result) {
	if len(_gggc) != 4 {
		return nil, MakeErrorResult(_aedab + "\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _gggc[0].Type != ResultTypeString {
		return nil, MakeErrorResult(_aedab + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u0073t\u0072\u0069\u006e\u0067")
	}
	_gdege := _gggc[0].ValueString
	if _gggc[1].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_aedab + " \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062e \u0061\u0020\u006eu\u006db\u0065\u0072")
	}
	_eagb := int(_gggc[1].ValueNumber) - 1
	if _gggc[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_aedab + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0069r\u0064\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_fgae := int(_gggc[2].ValueNumber)
	if _gggc[3].Type != ResultTypeString {
		return nil, MakeErrorResult(_aedab + " \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072\u0074\u0068\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062e \u0061\u0020\u0073t\u0072i\u006e\u0067")
	}
	_cadgb := _gggc[3].ValueString
	return &parsedReplaceObject{_gdege, _eagb, _fgae, _cadgb}, _feb
}
func init() {
	RegisterFunction("\u0043\u0048\u004f\u004f\u0053\u0045", Choose)
	RegisterFunction("\u0043\u004f\u004c\u0055\u004d\u004e", Column)
	RegisterFunction("\u0043O\u004c\u0055\u004d\u004e\u0053", Columns)
	RegisterFunction("\u0049\u004e\u0044E\u0058", Index)
	RegisterFunctionComplex("\u0049\u004e\u0044\u0049\u0052\u0045\u0043\u0054", Indirect)
	RegisterFunctionComplex("\u004f\u0046\u0046\u0053\u0045\u0054", Offset)
	RegisterFunction("\u004d\u0041\u0054C\u0048", Match)
	RegisterFunction("\u0048L\u004f\u004f\u004b\u0055\u0050", HLookup)
	RegisterFunction("\u004c\u0041\u0052G\u0045", Large)
	RegisterFunction("\u004c\u004f\u004f\u004b\u0055\u0050", Lookup)
	RegisterFunction("\u0052\u004f\u0057", Row)
	RegisterFunction("\u0052\u004f\u0057\u0053", Rows)
	RegisterFunction("\u0053\u004d\u0041L\u004c", Small)
	RegisterFunction("\u0056L\u004f\u004f\u004b\u0055\u0050", VLookup)
	RegisterFunction("\u0054R\u0041\u004e\u0053\u0050\u004f\u0053E", Transpose)
}

// String returns a string representation of a horizontal range.
func (_dgcd HorizontalRange) String() string { return _dgcd.horizontalRangeReference() }

// Tbilleq implements the Excel TBILLEQ function.
func Tbilleq(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0054\u0042\u0049\u004c\u004c\u0045\u0051\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020t\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_fcbe, _dbdb, _dccd := _gfcg(args[0], args[1], "\u0054B\u0049\u004c\u004c\u0045\u0051")
	if _dccd.Type == ResultTypeError {
		return _dccd
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("T\u0042\u0049\u004c\u004c\u0045\u0051\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0064is\u0063\u006f\u0075\u006et\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bffe := _dbdb - _fcbe
	if _bffe > 365 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004c\u0045\u0051\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006d\u0061\u0074\u0075\u0072\u0069\u0074\u0079\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eo\u0074\u0020m\u006f\u0072e\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065\u0020\u0079\u0065\u0061r \u0061\u0066\u0074\u0065\u0072\u0020\u0073\u0065\u0074t\u006c\u0065\u006d\u0065\u006e\u0074")
	}
	_ddbc := args[2].ValueNumber
	if _ddbc <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004c\u0045Q\u0020\u0072\u0065q\u0075\u0069\u0072e\u0073\u0020\u0064\u0069\u0073\u0063\u006f\u0075\u006e\u0074 \u0074\u006f\u0020\u0062\u0065 p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult((365 * _ddbc) / (360 - _ddbc*_bffe))
}

type criteriaParsed struct {
	_ceeeb bool
	_adag  float64
	_abgb  string
	_bgafc *criteriaRegex
}

// NewPrefixHorizontalRange constructs a new full rows range with prefix.
func NewPrefixHorizontalRange(pfx Expression, v string) Expression {
	_gfbad := _ga.Split(v, "\u003a")
	if len(_gfbad) != 2 {
		return nil
	}
	_gacgc, _ := _afc.Atoi(_gfbad[0])
	_cgdeg, _ := _afc.Atoi(_gfbad[1])
	if _gacgc > _cgdeg {
		_gacgc, _cgdeg = _cgdeg, _gacgc
	}
	return PrefixHorizontalRange{_dggf: pfx, _egbcc: _gacgc, _bcaa: _cgdeg}
}

type yySymType struct {
	_bddde int
	_ecdgg *node
	_ecac  Expression
	_adeda []Expression
	_ecda  [][]Expression
}

var _febb = [...]struct {
	_afac  int
	_ebae  int
	_agbgd string
}{}

func _ebggd(_eadde, _agfab float64) float64 {
	_eadde = _cc.Trunc(_eadde)
	_agfab = _cc.Trunc(_agfab)
	if _eadde == 0 {
		return _agfab
	}
	if _agfab == 0 {
		return _eadde
	}
	for _eadde != _agfab {
		if _eadde > _agfab {
			_eadde = _eadde - _agfab
		} else {
			_agfab = _agfab - _eadde
		}
	}
	return _eadde
}
func _gefg(_cfaaf Result, _ededb int) []Result {
	_caba := []Result{}
	switch _cfaaf.Type {
	case ResultTypeList:
		_badga := _cfaaf.ValueList
		_fcde := len(_badga)
		for _fbaa := 0; _fbaa < _ededb; _fbaa++ {
			if _fbaa < _fcde {
				_caba = append(_caba, _badga[_fbaa])
			} else {
				_caba = append(_caba, MakeErrorResultType(ErrorTypeNA, ""))
			}
		}
	case ResultTypeNumber, ResultTypeString, ResultTypeError, ResultTypeEmpty:
		for _efddc := 0; _efddc < _ededb; _efddc++ {
			_caba = append(_caba, _cfaaf)
		}
	}
	return _caba
}

// Reference returns a string reference value to a horizontal range with prefix.
func (_ddgf PrefixHorizontalRange) Reference(ctx Context, ev Evaluator) Reference {
	_ccfa := _ddgf._dggf.Reference(ctx, ev)
	return Reference{Type: ReferenceTypeHorizontalRange, Value: _ddgf.horizontalRangeReference(_ccfa.Value)}
}

// SupportedFunctions returns a list of supported functions.
func SupportedFunctions() []string {
	_fffbd := []string{}
	for _bcabb := range _bdgeg {
		_fffbd = append(_fffbd, _bcabb)
	}
	for _geefe := range _aaebee {
		_fffbd = append(_fffbd, _geefe)
	}
	_b.Strings(_fffbd)
	return _fffbd
}

// Proper is an implementation of the Excel PROPER function that returns a copy
// of the string with each word capitalized.
func Proper(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("P\u0052\u004f\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006eg \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_gaccb := args[0].AsString()
	if _gaccb.Type != ResultTypeString {
		return MakeErrorResult("P\u0052\u004f\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006eg \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_gbbg := _bd.Buffer{}
	_cedaf := false
	for _, _gegec := range _gaccb.ValueString {
		if !_cedaf && _a.IsLetter(_gegec) {
			_gbbg.WriteRune(_a.ToUpper(_gegec))
		} else {
			_gbbg.WriteRune(_a.ToLower(_gegec))
		}
		_cedaf = _a.IsLetter(_gegec)
	}
	return MakeStringResult(_gbbg.String())
}

type yyLexer interface {
	Lex(_deade *yySymType) int
	Error(_aceag string)
}

// SumProduct is an implementation of the Excel SUMPRODUCT() function.
func SumProduct(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0053\u0055\u004d\u0050\u0052\u004f\u0044U\u0043\u0054\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_aggd := args[0].Type
	for _, _aadfg := range args {
		if _aadfg.Type != _aggd {
			return MakeErrorResult("\u0053\u0055M\u0050\u0052\u004f\u0044\u0055C\u0054\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u006c\u006c\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0074\u0079\u0070\u0065")
		}
	}
	switch _aggd {
	case ResultTypeNumber:
		return Product(args)
	case ResultTypeList, ResultTypeArray:
		_gbbd := len(args[0].ListValues())
		_dcbc := make([]float64, _gbbd)
		for _gbefg := range _dcbc {
			_dcbc[_gbefg] = 1.0
		}
		for _, _daece := range args {
			if len(_daece.ListValues()) != _gbbd {
				return MakeErrorResult("\u0053\u0055\u004d\u0050\u0052\u004f\u0044\u0055\u0043\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069re\u0073 \u0061\u006c\u006c\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074s\u0020\u0074\u006f\u0020\u0068\u0061\u0076\u0065\u0020\u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065 \u0064\u0069\u006d\u0065\u006e\u0073\u0069\u006f\u006e")
			}
			for _acagb, _dfdc := range _daece.ListValues() {
				_dfdc = _dfdc.AsNumber()
				if _dfdc.Type != ResultTypeNumber {
					return MakeErrorResult("\u0053\u0055\u004d\u0050\u0052\u004fD\u0055\u0043\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u006c\u006c\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020n\u0075m\u0065\u0072\u0069\u0063")
				}
				_dcbc[_acagb] = _dcbc[_acagb] * _dfdc.ValueNumber
			}
		}
		_cbbb := 0.0
		for _, _efaadb := range _dcbc {
			_cbbb += _efaadb
		}
		return MakeNumberResult(_cbbb)
	}
	return MakeNumberResult(1.0)
}

// ErrorType is a formula evaluation error type.
type ErrorType byte

// Coupncd implements the Excel COUPNCD function.
func Coupncd(args []Result) Result {
	_fgbg, _ceaa := _cebf(args, "\u0043O\u0055\u0050\u004e\u0043\u0044")
	if _ceaa.Type == ResultTypeError {
		return _ceaa
	}
	_efaa := _bgee(_fgbg._bbaa)
	_afdc := _bgee(_fgbg._aee)
	_cddg := _fgbg._gbf
	_fcc := _eage(_efaa, _afdc, _cddg)
	_gbcb, _agae, _geab := _fcc.Date()
	return MakeNumberResult(_bad(_gbcb, int(_agae), _geab))
}

// ISODD is an implementation of the Excel ISODD() function.
func IsOdd(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u004f\u0044\u0044\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u004f\u0044\u0044\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061 \u006eu\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fedc := int(args[0].ValueNumber)
	return MakeBoolResult(_fedc != _fedc/2*2)
}

// Columns implements the Excel COLUMNS function.
func Columns(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0043\u004fL\u0055\u004d\u004e\u0053\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075me\u006e\u0074")
	}
	_ggec := args[0]
	if _ggec.Type != ResultTypeArray && _ggec.Type != ResultTypeList {
		return MakeErrorResult("\u0043O\u004c\u0055M\u004e\u0053\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0066\u0069\u0072\u0073\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020t\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_adbc := _ggec.ValueArray
	if len(_adbc) == 0 {
		return MakeErrorResult("\u0043\u004f\u004c\u0055\u004d\u004e\u0053\u0020r\u0065\u0071\u0075ir\u0065\u0073\u0020\u0061\u0072\u0072a\u0079\u0020\u0074\u006f\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020\u0031\u0020r\u006f\u0077")
	}
	return MakeNumberResult(float64(len(_adbc[0])))
}

type evCache struct {
	_bgg map[string]Result
	_ef  *_af.Mutex
}

// Len is an implementation of the Excel LEN function that returns length of a string
func Len(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004c\u0045N\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067l\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_gdbd := args[0].AsString()
	if _gdbd.Type != ResultTypeString {
		return MakeErrorResult("\u004c\u0045N\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067l\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(float64(len(_gdbd.ValueString)))
}

// GetFormat returns an empty string for the invalid reference context.
func (_fccf *ivr) GetFormat(cellRef string) string { return "" }

const _gbg = "\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u002d\u0028\u0028\u005b\u0030-\u0039]\u0029\u002b\u0029\u002d\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029"

func init() {
	_cccg()
	RegisterFunction("\u004e\u0041", NA)
	RegisterFunction("\u0049S\u0042\u004c\u0041\u004e\u004b", IsBlank)
	RegisterFunction("\u0049\u0053\u0045R\u0052", IsErr)
	RegisterFunction("\u0049S\u0045\u0052\u0052\u004f\u0052", IsError)
	RegisterFunction("\u0049\u0053\u0045\u0056\u0045\u004e", IsEven)
	RegisterFunctionComplex("\u005fx\u006cf\u006e\u002e\u0049\u0053\u0046\u004f\u0052\u004d\u0055\u004c\u0041", IsFormula)
	RegisterFunctionComplex("\u004fR\u0047\u002e\u004f\u0050E\u004e\u004f\u0046\u0046\u0049C\u0045.\u0049S\u004c\u0045\u0041\u0050\u0059\u0045\u0041R", IsLeapYear)
	RegisterFunctionComplex("\u0049S\u004c\u004f\u0047\u0049\u0043\u0041L", IsLogical)
	RegisterFunction("\u0049\u0053\u004e\u0041", IsNA)
	RegisterFunction("\u0049S\u004e\u004f\u004e\u0054\u0045\u0058T", IsNonText)
	RegisterFunction("\u0049\u0053\u004e\u0055\u004d\u0042\u0045\u0052", IsNumber)
	RegisterFunction("\u0049\u0053\u004fD\u0044", IsOdd)
	RegisterFunctionComplex("\u0049\u0053\u0052E\u0046", IsRef)
	RegisterFunction("\u0049\u0053\u0054\u0045\u0058\u0054", IsText)
	RegisterFunctionComplex("\u0043\u0045\u004c\u004c", Cell)
}

// Average implements the AVERAGE function. It differs slightly from Excel (and
// agrees with LibreOffice) in that boolean values are counted. As an example,
// AVERAGE of two cells containing TRUE & FALSE is 0.5 in LibreOffice and
// #DIV/0! in Excel. github.com/arcpd/msword will return 0.5 in this case.
func Average(args []Result) Result {
	_dgded, _ccead := _ebbcf(args, false)
	if _ccead == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0041\u0056\u0045\u0052AG\u0045\u0020\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0062\u0079\u0020\u007a\u0065r\u006f")
	}
	return MakeNumberResult(_dgded / _ccead)
}

// Clean is an implementation of the Excel CLEAN function that removes
// unprintable characters.
func Clean(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0043\u004c\u0045\u0041\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_bdfbc := args[0].AsString()
	if _bdfbc.Type != ResultTypeString {
		return MakeErrorResult("\u0043\u0048\u0041\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_eeadb := _bd.Buffer{}
	for _, _dggcf := range _bdfbc.ValueString {
		if _a.IsPrint(_dggcf) {
			_eeadb.WriteRune(_dggcf)
		}
	}
	return MakeStringResult(_eeadb.String())
}

// Findb is an implementation of the Excel FINDB().
func Findb(ctx Context, ev Evaluator, args []Result) Result {
	if !ctx.IsDBCS() {
		return Find(args)
	}
	_bgbeb, _dfgfd := _faafg("\u0046\u0049\u004e\u0044", args)
	if _dfgfd.Type != ResultTypeEmpty {
		return _dfgfd
	}
	_gaaff := _bgbeb._eccf
	if _gaaff == "" {
		return MakeNumberResult(1.0)
	}
	_ggcc := _bgbeb._bccec
	_fdea := _bgbeb._aeag - 1
	_acfd := 1
	_egfgb := 0
	for _adab := range _ggcc {
		if _adab != 0 {
			_fcdd := 1
			if _adab-_egfgb > 1 {
				_fcdd = 2
			}
			_acfd += _fcdd
		}
		if _acfd > _fdea {
			_gegg := _ga.Index(_ggcc[_adab:], _gaaff)
			if _gegg == 0 {
				return MakeNumberResult(float64(_acfd))
			}
		}
		_egfgb = _adab
	}
	return MakeErrorResultType(ErrorTypeValue, "\u004eo\u0074\u0020\u0066\u006f\u0075\u006ed")
}

const _eeffc = 57354

// Bool is a boolean expression.
type Bool struct{ _aef bool }

const _dcf = "\u0042\u0069\u006e\u004f\u0070\u0054y\u0070\u0065\u0055\u006e\u006bn\u006fw\u006e\u0042\u0069\u006eO\u0070\u0054\u0079\u0070\u0065\u0050\u006c\u0075\u0073\u0042\u0069\u006eO\u0070\u0054\u0079\u0070\u0065\u004d\u0069\u006e\u0075\u0073\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065M\u0075lt\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065\u0044\u0069\u0076\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065\u0045\u0078\u0070\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065\u004c\u0054\u0042\u0069\u006eO\u0070\u0054\u0079\u0070\u0065G\u0054B\u0069\u006eO\u0070\u0054\u0079\u0070\u0065\u0045\u0051\u0042\u0069nO\u0070\u0054\u0079\u0070\u0065\u004c\u0045\u0051\u0042i\u006eO\u0070\u0054\u0079\u0070\u0065\u0047\u0045\u0051\u0042\u0069\u006e\u004f\u0070\u0054\u0079\u0070\u0065N\u0045\u0042\u0069\u006eO\u0070\u0054\u0079\u0070\u0065\u0043\u006f\u006e\u0063\u0061\u0074"

// Multinomial implements the excel MULTINOMIAL function.
func Multinomial(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u004d\u0055\u004c\u0054\u0049\u004eO\u004d\u0049\u0041\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006ce\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c\u0020i\u006e\u0070\u0075\u0074")
	}
	_accb, _cbbed, _gcab := _bgdb(args)
	if _gcab.Type == ResultTypeError {
		return _gcab
	}
	return MakeNumberResult(_cfga(_accb) / _cbbed)
}

// NewNamedRangeRef constructs a new named range reference.
func NewNamedRangeRef(v string) Expression { return NamedRangeRef{_ceaf: v} }

// Day is an implementation of the Excel DAY() function.
func Day(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0044A\u0059\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065 \u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_acbf := args[0]
	switch _acbf.Type {
	case ResultTypeEmpty:
		return MakeNumberResult(0)
	case ResultTypeNumber:
		_abc := _bgee(_acbf.ValueNumber)
		return MakeNumberResult(float64(_abc.Day()))
	case ResultTypeString:
		_ecag := _ga.ToLower(_acbf.ValueString)
		if !_cddc(_ecag) {
			_, _, _, _, _egd, _afcf := _ccce(_ecag)
			if _afcf.Type == ResultTypeError {
				_afcf.ErrorMessage = "I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073\u0020\u0066o\u0072 \u0044\u0041\u0059"
				return _afcf
			}
			if _egd {
				return MakeNumberResult(0)
			}
		}
		_, _, _cbb, _, _gca := _ddd(_ecag)
		if _gca.Type == ResultTypeError {
			return _gca
		}
		return MakeNumberResult(float64(_cbb))
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072 \u0044\u0041\u0059")
	}
}

// String returns a string representation of PrefixExpr.
func (_ededbb PrefixExpr) String() string {
	return _e.Sprintf("\u0025\u0073\u0021%\u0073", _ededbb._dfgab.String(), _ededbb._bcada.String())
}

var _fae = map[string]int{"\u006aa\u006e\u0075\u0061\u0072\u0079": 1, "\u0066\u0065\u0062\u0072\u0075\u0061\u0072\u0079": 2, "\u006d\u0061\u0072c\u0068": 3, "\u0061\u0070\u0072i\u006c": 4, "\u006d\u0061\u0079": 5, "\u006a\u0075\u006e\u0065": 6, "\u006a\u0075\u006c\u0079": 7, "\u0061\u0075\u0067\u0075\u0073\u0074": 8, "\u0073e\u0070\u0074\u0065\u006d\u0070\u0065r": 9, "\u006fc\u0074\u006f\u0062\u0065\u0072": 10, "\u006e\u006f\u0076\u0065\u006d\u0062\u0065\u0072": 11, "\u0064\u0065\u0063\u0065\u006d\u0062\u0065\u0072": 12, "\u006a\u0061\u006e": 1, "\u0066\u0065\u0062": 2, "\u006d\u0061\u0072": 3, "\u0061\u0070\u0072": 4, "\u006a\u0075\u006e": 6, "\u006a\u0075\u006c": 7, "\u0061\u0075\u0067": 8, "\u0073\u0065\u0070": 9, "\u006f\u0063\u0074": 10, "\u006e\u006f\u0076": 11, "\u0064\u0065\u0063": 12}

// Eval evaluates and returns a boolean.
func (_bge Bool) Eval(ctx Context, ev Evaluator) Result { return MakeBoolResult(_bge._aef) }
func _aefe(_aaec []Result, _fefg []string, _gdgeg bool) []string {
	for _, _ebaag := range _aaec {
		switch _ebaag.Type {
		case ResultTypeEmpty:
			if !_gdgeg {
				_fefg = append(_fefg, "")
			}
		case ResultTypeString:
			if _ebaag.ValueString != "" || !_gdgeg {
				_fefg = append(_fefg, _ebaag.ValueString)
			}
		case ResultTypeNumber:
			_fefg = append(_fefg, _ebaag.Value())
		case ResultTypeList:
			_fefg = _bgdeg(_fefg, _aefe(_ebaag.ValueList, []string{}, _gdgeg))
		case ResultTypeArray:
			for _, _gbcdd := range _ebaag.ValueArray {
				_fefg = _bgdeg(_fefg, _aefe(_gbcdd, []string{}, _gdgeg))
			}
		}
	}
	return _fefg
}

// MakeBoolResult constructs a boolean result (internally a number).
func MakeBoolResult(b bool) Result {
	if b {
		return Result{Type: ResultTypeNumber, ValueNumber: 1, IsBoolean: true}
	}
	return Result{Type: ResultTypeNumber, ValueNumber: 0, IsBoolean: true}
}

// Log implements the Excel LOG function which returns the log of a number. By
// default the result is base 10, however the second argument to the function
// can specify a different base.
func Log(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 o\u006e\u0065\u0020n\u0075\u006de\u0072\u0069\u0063\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("L\u004f\u0047\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u006d\u0061\u0078i\u006d\u0075\u006d\u0020\u006f\u0066\u0020\u0074\u0077\u006f a\u0072\u0067\u0075m\u0065n\u0074\u0073")
	}
	_bcba := args[0].AsNumber()
	if _bcba.Type != ResultTypeNumber {
		return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 o\u006e\u0065\u0020n\u0075\u006de\u0072\u0069\u0063\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_agdf := 10.0
	if len(args) > 1 {
		_fabbff := args[1].AsNumber()
		if _fabbff.Type != ResultTypeNumber {
			return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061rg\u0075\u006d\u0065n\u0074\u0020t\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0065r\u0069\u0063")
		}
		_agdf = args[1].ValueNumber
	}
	if _bcba.ValueNumber == 0 {
		return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072gu\u006d\u0065\u006et\u0020\u0074o\u0020\u0062\u0065\u0020\u006e\u006fn\u002d\u007ae\u0072\u006f")
	}
	if _agdf == 0 {
		return MakeErrorResult("\u004cO\u0047\u0028)\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0073e\u0063\u006f\u006e\u0064\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062e\u0020\u006e\u006f\u006e\u002d\u007a\u0065\u0072\u006f")
	}
	return MakeNumberResult(_cc.Log(_bcba.ValueNumber) / _cc.Log(_agdf))
}
func _gadc(_ffff Result) bool {
	_fcacc := _ffff.Type
	return _fcacc != ResultTypeArray && _fcacc != ResultTypeList
}

// Searchb is an implementation of the Excel SEARCHB().
func Searchb(ctx Context, ev Evaluator, args []Result) Result {
	if !ctx.IsDBCS() {
		return Search(args)
	}
	_cdff, _gcfc := _faafg("\u0046\u0049\u004e\u0044", args)
	if _gcfc.Type != ResultTypeEmpty {
		return _gcfc
	}
	_eecad := _ga.ToLower(_cdff._eccf)
	_facbf := _ga.ToLower(_cdff._bccec)
	if _eecad == "" {
		return MakeNumberResult(1.0)
	}
	_gdegg := _cdff._aeag - 1
	_ceadg := 1
	_cecab := 0
	for _bggc := range _facbf {
		if _bggc != 0 {
			_gcffb := 1
			if _bggc-_cecab > 1 {
				_gcffb = 2
			}
			_ceadg += _gcffb
		}
		if _ceadg > _gdegg {
			_gebb := _cd.Index(_eecad, _facbf[_bggc:])
			if _gebb == 0 {
				return MakeNumberResult(float64(_ceadg))
			}
		}
		_cecab = _bggc
	}
	return MakeErrorResultType(ErrorTypeValue, "\u004eo\u0074\u0020\u0066\u006f\u0075\u006ed")
}

const _cccde = 57372

// CountIf implements the COUNTIF function.
func CountIf(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u0043\u004f\u0055N\u0054\u0049\u0046\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0073")
	}
	_dcgc := args[0]
	if _dcgc.Type != ResultTypeArray && _dcgc.Type != ResultTypeList {
		return MakeErrorResult("\u0043O\u0055\u004eT\u0049\u0046\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0066\u0069\u0072\u0073\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020t\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_cfcb := _cafa(args[1])
	_dcgg := 0
	for _, _bbfaa := range _aadag(_dcgc) {
		for _, _cecbg := range _bbfaa {
			if _addg(_cecbg, _cfcb) {
				_dcgg++
			}
		}
	}
	return MakeNumberResult(float64(_dcgg))
}

// Transpose implements the TRANSPOSE function that transposes a cell range.
func Transpose(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0054\u0052AN\u0053\u0050\u004fS\u0045\u0020\u0072\u0065qui\u0072es\u0020\u0061\u0020\u0073\u0069\u006e\u0067le\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	if args[0].Type != ResultTypeArray && args[0].Type != ResultTypeList {
		return MakeErrorResult("T\u0052\u0041\u004e\u0053\u0050\u004fS\u0045\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0061\u0020\u0072a\u006e\u0067\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_caee := args[0]
	if _caee.Type == ResultTypeList {
		_fcgf := [][]Result{}
		for _, _aaea := range _caee.ValueList {
			_fcgf = append(_fcgf, []Result{_aaea})
		}
		return MakeArrayResult(_fcgf)
	}
	_faee := make([][]Result, len(_caee.ValueArray[0]))
	for _, _ffea := range _caee.ValueArray {
		for _febc, _ggaab := range _ffea {
			_faee[_febc] = append(_faee[_febc], _ggaab)
		}
	}
	return MakeArrayResult(_faee)
}

type Expression interface {
	Eval(_gfgg Context, _gga Evaluator) Result
	Reference(_ggg Context, _efg Evaluator) Reference
	String() string
	Update(_fcaf *_eg.UpdateQuery) Expression
}

func init() {
	RegisterFunction("\u0041\u004e\u0044", And)
	RegisterFunction("\u0046\u0041\u004cS\u0045", False)
	RegisterFunction("\u0049\u0046", If)
	RegisterFunction("\u0049F\u0045\u0052\u0052\u004f\u0052", IfError)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0049\u0046\u004e\u0041", IfNA)
	RegisterFunction("\u0049\u0046\u0053", Ifs)
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0049\u0046S", Ifs)
	RegisterFunction("\u004e\u004f\u0054", Not)
	RegisterFunction("\u004f\u0052", Or)
	RegisterFunction("\u0054\u0052\u0055\u0045", True)
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0058\u004fR", Xor)
}

var _dag = []*_f.Regexp{}

func _bggf(_ebgb, _aecc float64) float64 {
	_ebgb = _cc.Trunc(_ebgb)
	_aecc = _cc.Trunc(_aecc)
	if _ebgb == 0 && _aecc == 0 {
		return 0
	}
	return _ebgb * _aecc / _ebggd(_ebgb, _aecc)
}

// MDeterm is an implementation of the Excel MDETERM which finds the determinant
// of a matrix.
func MDeterm(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004d\u0044\u0045T\u0045\u0052\u004d\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0072\u0061\u0079 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bbcb := args[0]
	if _bbcb.Type != ResultTypeArray {
		return MakeErrorResult("\u004d\u0044\u0045T\u0045\u0052\u004d\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0072\u0061\u0079 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fgfeg := len(_bbcb.ValueArray)
	for _, _dcefd := range _bbcb.ValueArray {
		if len(_dcefd) != _fgfeg {
			return MakeErrorResult("\u004d\u0044\u0045TE\u0052\u004d\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072e\u0073 \u0061 \u0073\u0071\u0075\u0061\u0072\u0065\u0020\u006d\u0061\u0074\u0072\u0069\u0078")
		}
	}
	return MakeNumberResult(_beagb(_bbcb.ValueArray))
}
func _eedb(_gbef, _gfbf _eb.Time) bool {
	_eefe := _gbef.Unix()
	_gdbg := _gfbf.Unix()
	_dbc := _gbef.Year()
	_egcd := _gadd(_dbc, _eb.March, 1)
	if _bee(_dbc) && _eefe < _egcd && _gdbg >= _egcd {
		return true
	}
	var _cgdc = _gfbf.Year()
	var _aea = _gadd(_cgdc, _eb.March, 1)
	return (_bee(_cgdc) && _gdbg >= _aea && _eefe < _aea)
}
func (_egbbbf *Lexer) lex(_bacd _ee.Reader) {
	_ccege, _fddfd, _egdaa := 0, 0, 0
	_cbfef := -1
	_dfdf, _bfcbg, _gebed := 0, 0, 0
	_ = _gebed
	_ffac := 1
	_ = _ffac
	_aaag := make([]byte, 4096)
	_aadg := false
	for !_aadg {
		_ccfec := 0
		if _dfdf > 0 {
			_ccfec = _fddfd - _dfdf
		}
		_fddfd = 0
		_bcbea, _gcdf := _bacd.Read(_aaag[_ccfec:])
		if _bcbea == 0 || _gcdf != nil {
			_aadg = true
		}
		_egdaa = _bcbea + _ccfec
		if _egdaa < len(_aaag) {
			_cbfef = _egdaa
		}
		{
			_ccege = _egefd
			_dfdf = 0
			_bfcbg = 0
			_gebed = 0
		}
		{
			var _gfgb int
			var _dbad uint
			if _fddfd == _egdaa {
				goto _eddec
			}
			if _ccege == 0 {
				goto _fcfcf
			}
		_bgbcc:
			_gfgb = int(_bgfb[_ccege])
			_dbad = uint(_ebee[_gfgb])
			_gfgb++
			for ; _dbad > 0; _dbad-- {
				_gfgb++
				switch _ebee[_gfgb-1] {
				case 2:
					_dfdf = _fddfd
				}
			}
			switch _ccege {
			case 30:
				switch _aaag[_fddfd] {
				case 34:
					goto _fgdc
				case 35:
					goto _dfga
				case 36:
					goto _egga
				case 38:
					goto _fbad
				case 39:
					goto _bgab
				case 40:
					goto _deca
				case 41:
					goto _fcefd
				case 42:
					goto _febcb
				case 43:
					goto _ecdc
				case 44:
					goto _cbbad
				case 45:
					goto _decag
				case 47:
					goto _egaed
				case 58:
					goto _fgbcc
				case 59:
					goto _afgc
				case 60:
					goto _ggccg
				case 61:
					goto _ebagc
				case 62:
					goto _agcfc
				case 63:
					goto _geee
				case 70:
					goto _gcca
				case 84:
					goto _gdad
				case 92:
					goto _gecfg
				case 94:
					goto _ffdae
				case 95:
					goto _ddedg
				case 123:
					goto _aaaec
				case 125:
					goto _bbgff
				}
				switch {
				case _aaag[_fddfd] < 65:
					switch {
					case _aaag[_fddfd] > 37:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbadd
						}
					case _aaag[_fddfd] >= 33:
						goto _geee
					}
				case _aaag[_fddfd] > 90:
					switch {
					case _aaag[_fddfd] > 93:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _cabcf
						}
					case _aaag[_fddfd] >= 91:
						goto _geee
					}
				default:
					goto _adgca
				}
				goto _gbeaa
			case 1:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 47:
					goto _efbd
				case 123:
					goto _efbd
				case 125:
					goto _efbd
				}
				switch {
				case _aaag[_fddfd] < 37:
					if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
						goto _efbd
					}
				case _aaag[_fddfd] > 45:
					switch {
					case _aaag[_fddfd] > 63:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _efbd
						}
					case _aaag[_fddfd] >= 58:
						goto _efbd
					}
				default:
					goto _efbd
				}
				goto _gbeaa
			case 0:
				goto _fcfcf
			case 2:
				if _aaag[_fddfd] == 34 {
					goto _ebcfa
				}
				goto _fgdc
			case 31:
				if _aaag[_fddfd] == 34 {
					goto _fgdc
				}
				goto _gafc
			case 3:
				switch _aaag[_fddfd] {
				case 78:
					goto _dgdg
				case 82:
					goto _bcaee
				}
				goto _geee
			case 4:
				switch _aaag[_fddfd] {
				case 47:
					goto _ecfgb
				case 85:
					goto _cfdf
				}
				goto _geee
			case 5:
				if _aaag[_fddfd] == 65 {
					goto _cbdce
				}
				goto _geee
			case 6:
				switch _aaag[_fddfd] {
				case 76:
					goto _eacb
				case 77:
					goto _bbgac
				}
				goto _geee
			case 7:
				if _aaag[_fddfd] == 76 {
					goto _bbgac
				}
				goto _geee
			case 8:
				if _aaag[_fddfd] == 33 {
					goto _cbdce
				}
				goto _geee
			case 9:
				if _aaag[_fddfd] == 69 {
					goto _begbc
				}
				goto _geee
			case 10:
				if _aaag[_fddfd] == 70 {
					goto _cfdgf
				}
				goto _geee
			case 11:
				if _aaag[_fddfd] == 33 {
					goto _bgea
				}
				goto _geee
			case 12:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 47:
					goto _geee
				case 123:
					goto _geee
				case 125:
					goto _geee
				}
				switch {
				case _aaag[_fddfd] < 48:
					switch {
					case _aaag[_fddfd] > 35:
						if 37 <= _aaag[_fddfd] && _aaag[_fddfd] <= 45 {
							goto _geee
						}
					case _aaag[_fddfd] >= 34:
						goto _geee
					}
				case _aaag[_fddfd] > 57:
					switch {
					case _aaag[_fddfd] < 65:
						if 58 <= _aaag[_fddfd] && _aaag[_fddfd] <= 63 {
							goto _geee
						}
					case _aaag[_fddfd] > 90:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _geee
						}
					default:
						goto _faceg
					}
				default:
					goto _fafd
				}
				goto _gbeaa
			case 13:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 47:
					goto _geee
				case 58:
					goto _ggdge
				case 123:
					goto _geee
				case 125:
					goto _geee
				}
				switch {
				case _aaag[_fddfd] < 48:
					switch {
					case _aaag[_fddfd] > 35:
						if 37 <= _aaag[_fddfd] && _aaag[_fddfd] <= 45 {
							goto _geee
						}
					case _aaag[_fddfd] >= 34:
						goto _geee
					}
				case _aaag[_fddfd] > 57:
					switch {
					case _aaag[_fddfd] > 63:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _geee
						}
					case _aaag[_fddfd] >= 59:
						goto _geee
					}
				default:
					goto _fafd
				}
				goto _gbeaa
			case 14:
				if _aaag[_fddfd] == 36 {
					goto _cccb
				}
				if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
					goto _fabbc
				}
				goto _efbd
			case 15:
				if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
					goto _fabbc
				}
				goto _efbd
			case 32:
				if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
					goto _fabbc
				}
				goto _fgada
			case 16:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 47:
					goto _geee
				case 58:
					goto _afgfg
				case 123:
					goto _geee
				case 125:
					goto _geee
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 45:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _affba
						}
					case _aaag[_fddfd] >= 34:
						goto _geee
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] > 90:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _geee
						}
					case _aaag[_fddfd] >= 65:
						goto _faceg
					}
				default:
					goto _geee
				}
				goto _gbeaa
			case 17:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 47:
					goto _efbd
				case 123:
					goto _efbd
				case 125:
					goto _efbd
				}
				switch {
				case _aaag[_fddfd] < 48:
					switch {
					case _aaag[_fddfd] > 35:
						if 37 <= _aaag[_fddfd] && _aaag[_fddfd] <= 45 {
							goto _efbd
						}
					case _aaag[_fddfd] >= 34:
						goto _efbd
					}
				case _aaag[_fddfd] > 57:
					switch {
					case _aaag[_fddfd] > 63:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _efbd
						}
					case _aaag[_fddfd] >= 58:
						goto _efbd
					}
				default:
					goto _affba
				}
				goto _gbeaa
			case 33:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 47:
					goto _cebag
				case 123:
					goto _cebag
				case 125:
					goto _cebag
				}
				switch {
				case _aaag[_fddfd] < 48:
					switch {
					case _aaag[_fddfd] > 35:
						if 37 <= _aaag[_fddfd] && _aaag[_fddfd] <= 45 {
							goto _cebag
						}
					case _aaag[_fddfd] >= 34:
						goto _cebag
					}
				case _aaag[_fddfd] > 57:
					switch {
					case _aaag[_fddfd] > 63:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _cebag
						}
					case _aaag[_fddfd] >= 58:
						goto _cebag
					}
				default:
					goto _affba
				}
				goto _gbeaa
			case 18:
				if _aaag[_fddfd] == 36 {
					goto _gebbg
				}
				if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
					goto _gceac
				}
				goto _efbd
			case 19:
				if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
					goto _gceac
				}
				goto _efbd
			case 34:
				if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
					goto _gceac
				}
				goto _dbfgb
			case 20:
				switch _aaag[_fddfd] {
				case 39:
					goto _geee
				case 42:
					goto _geee
				case 47:
					goto _geee
				case 58:
					goto _geee
				case 63:
					goto _geee
				}
				if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 93 {
					goto _geee
				}
				goto _acffe
			case 21:
				switch _aaag[_fddfd] {
				case 39:
					goto _fafc
				case 42:
					goto _geee
				case 47:
					goto _geee
				case 58:
					goto _geee
				case 63:
					goto _geee
				}
				if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 93 {
					goto _geee
				}
				goto _acffe
			case 22:
				if _aaag[_fddfd] == 33 {
					goto _eebg
				}
				goto _geee
			case 35:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _gbfdee
				case 58:
					goto _ggdge
				case 101:
					goto _gedad
				case 123:
					goto _afec
				case 125:
					goto _afec
				}
				switch {
				case _aaag[_fddfd] < 48:
					switch {
					case _aaag[_fddfd] > 35:
						if 37 <= _aaag[_fddfd] && _aaag[_fddfd] <= 47 {
							goto _afec
						}
					case _aaag[_fddfd] >= 34:
						goto _afec
					}
				case _aaag[_fddfd] > 57:
					switch {
					case _aaag[_fddfd] > 63:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _afec
						}
					case _aaag[_fddfd] >= 59:
						goto _afec
					}
				default:
					goto _bbadd
				}
				goto _gbeaa
			case 36:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 47:
					goto _afec
				case 101:
					goto _gedad
				case 123:
					goto _afec
				case 125:
					goto _afec
				}
				switch {
				case _aaag[_fddfd] < 48:
					switch {
					case _aaag[_fddfd] > 35:
						if 37 <= _aaag[_fddfd] && _aaag[_fddfd] <= 45 {
							goto _afec
						}
					case _aaag[_fddfd] >= 34:
						goto _afec
					}
				case _aaag[_fddfd] > 57:
					switch {
					case _aaag[_fddfd] > 63:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _afec
						}
					case _aaag[_fddfd] >= 58:
						goto _afec
					}
				default:
					goto _gbfdee
				}
				goto _gbeaa
			case 23:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 47:
					goto _cggfc
				case 123:
					goto _cggfc
				case 125:
					goto _cggfc
				}
				switch {
				case _aaag[_fddfd] < 48:
					switch {
					case _aaag[_fddfd] > 35:
						if 37 <= _aaag[_fddfd] && _aaag[_fddfd] <= 45 {
							goto _cggfc
						}
					case _aaag[_fddfd] >= 34:
						goto _cggfc
					}
				case _aaag[_fddfd] > 57:
					switch {
					case _aaag[_fddfd] > 63:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _cggfc
						}
					case _aaag[_fddfd] >= 58:
						goto _cggfc
					}
				default:
					goto _edbg
				}
				goto _gbeaa
			case 37:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 47:
					goto _afec
				case 123:
					goto _afec
				case 125:
					goto _afec
				}
				switch {
				case _aaag[_fddfd] < 48:
					switch {
					case _aaag[_fddfd] > 35:
						if 37 <= _aaag[_fddfd] && _aaag[_fddfd] <= 45 {
							goto _afec
						}
					case _aaag[_fddfd] >= 34:
						goto _afec
					}
				case _aaag[_fddfd] > 57:
					switch {
					case _aaag[_fddfd] > 63:
						if 91 <= _aaag[_fddfd] && _aaag[_fddfd] <= 94 {
							goto _afec
						}
					case _aaag[_fddfd] >= 58:
						goto _afec
					}
				default:
					goto _edbg
				}
				goto _gbeaa
			case 38:
				switch _aaag[_fddfd] {
				case 61:
					goto _gfgd
				case 62:
					goto _gbfde
				}
				goto _egcb
			case 39:
				if _aaag[_fddfd] == 61 {
					goto _cccee
				}
				goto _adgb
			case 24:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _geee
				case 125:
					goto _geee
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _geee
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _geee
					}
				default:
					goto _geee
				}
				goto _gbeaa
			case 40:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _aebf
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _aebf
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 41:
				switch _aaag[_fddfd] {
				case 46:
					goto _gdbgd
				case 92:
					goto _gdbgd
				case 95:
					goto _gdbgd
				}
				switch {
				case _aaag[_fddfd] < 65:
					if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
						goto _gdbgd
					}
				case _aaag[_fddfd] > 90:
					if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
						goto _gdbgd
					}
				default:
					goto _gdbgd
				}
				goto _ccga
			case 42:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _bbagd
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _bbagd
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 43:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _cebag
				case 125:
					goto _cebag
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _cebag
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					default:
						goto _cebag
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _aebf
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _cebag
					}
				default:
					goto _cebag
				}
				goto _gbeaa
			case 44:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _efbd
				case 125:
					goto _efbd
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _efbd
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _efbd
					}
				default:
					goto _efbd
				}
				goto _gbeaa
			case 25:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 65:
					goto _cgdcg
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _geee
				case 125:
					goto _geee
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _geee
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 66 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _geee
					}
				default:
					goto _geee
				}
				goto _gbeaa
			case 45:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 76:
					goto _gacd
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 46:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 83:
					goto _bcfg
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 47:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 69:
					goto _cdadc
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 26:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 79:
					goto _bbeb
				case 82:
					goto _cdce
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _geee
				case 125:
					goto _geee
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _geee
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _geee
					}
				default:
					goto _geee
				}
				goto _gbeaa
			case 48:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 68:
					goto _bcge
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 49:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 79:
					goto _fgcff
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 50:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 36:
					goto _gccgd
				case 40:
					goto _efgcg
				case 46:
					goto _aebf
				case 58:
					goto _afgfg
				case 85:
					goto _bcfg
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 59:
					switch {
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _ceacc
						}
					case _aaag[_fddfd] >= 34:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _feec
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 27:
				switch _aaag[_fddfd] {
				case 46:
					goto _gdbgd
				case 92:
					goto _gdbgd
				case 95:
					goto _gdbgd
				}
				switch {
				case _aaag[_fddfd] < 65:
					if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
						goto _gdbgd
					}
				case _aaag[_fddfd] > 90:
					if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
						goto _gdbgd
					}
				default:
					goto _gdbgd
				}
				goto _geee
			case 28:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _bbagd
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 120:
					goto _cegb
				case 123:
					goto _geee
				case 125:
					goto _geee
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _geee
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _geee
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _bbagd
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _geee
					}
				default:
					goto _geee
				}
				goto _gbeaa
			case 51:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _bbagd
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 108:
					goto _aebfe
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _bbagd
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 52:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _bbagd
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 102:
					goto _cebc
				case 110:
					goto _gdcgd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _bbagd
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 53:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _bbagd
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 110:
					goto _ebeac
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _bbagd
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 54:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _dcaed
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _bbagd
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 55:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _bbagd
				case 92:
					goto _gdbgd
				case 95:
					goto _gddb
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _gddb
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 56:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 40:
					goto _beagf
				case 46:
					goto _gddb
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _gddb
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _gddb
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 57:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _bbagd
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 109:
					goto _ecdce
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _bbagd
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 58:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _ggcdc
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _ccga
				case 125:
					goto _ccga
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _ccga
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _bbagd
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _ccga
					}
				default:
					goto _ccga
				}
				goto _gbeaa
			case 59:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _bbagd
				case 92:
					goto _gdbgd
				case 95:
					goto _gcfcc
				case 123:
					goto _efbd
				case 125:
					goto _efbd
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _efbd
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _efbd
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _gcfcc
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _efbd
					}
				default:
					goto _efbd
				}
				goto _gbeaa
			case 29:
				switch _aaag[_fddfd] {
				case 33:
					goto _gaca
				case 46:
					goto _bbagd
				case 92:
					goto _gdbgd
				case 95:
					goto _bbagd
				case 123:
					goto _geee
				case 125:
					goto _geee
				}
				switch {
				case _aaag[_fddfd] < 58:
					switch {
					case _aaag[_fddfd] < 37:
						if 34 <= _aaag[_fddfd] && _aaag[_fddfd] <= 35 {
							goto _geee
						}
					case _aaag[_fddfd] > 47:
						if 48 <= _aaag[_fddfd] && _aaag[_fddfd] <= 57 {
							goto _bbagd
						}
					default:
						goto _geee
					}
				case _aaag[_fddfd] > 63:
					switch {
					case _aaag[_fddfd] < 91:
						if 65 <= _aaag[_fddfd] && _aaag[_fddfd] <= 90 {
							goto _bbagd
						}
					case _aaag[_fddfd] > 94:
						if 97 <= _aaag[_fddfd] && _aaag[_fddfd] <= 122 {
							goto _bbagd
						}
					default:
						goto _geee
					}
				default:
					goto _geee
				}
				goto _gbeaa
			}
		_geee:
			_ccege = 0
			goto _gdgfa
		_gbeaa:
			_ccege = 1
			goto _gdgfa
		_fgdc:
			_ccege = 2
			goto _gdgfa
		_dfga:
			_ccege = 3
			goto _gdgfa
		_dgdg:
			_ccege = 4
			goto _gdgfa
		_ecfgb:
			_ccege = 5
			goto _gdgfa
		_cfdf:
			_ccege = 6
			goto _gdgfa
		_eacb:
			_ccege = 7
			goto _gdgfa
		_bbgac:
			_ccege = 8
			goto _gdgfa
		_bcaee:
			_ccege = 9
			goto _gdgfa
		_begbc:
			_ccege = 10
			goto _gdgfa
		_cfdgf:
			_ccege = 11
			goto _gdgfa
		_egga:
			_ccege = 12
			goto _gdgfa
		_fafd:
			_ccege = 13
			goto _gdgfa
		_ggdge:
			_ccege = 14
			goto _gdgfa
		_cccb:
			_ccege = 15
			goto _gdgfa
		_faceg:
			_ccege = 16
			goto _gdgfa
		_gccgd:
			_ccege = 17
			goto _gdgfa
		_afgfg:
			_ccege = 18
			goto _gdgfa
		_gebbg:
			_ccege = 19
			goto _gdgfa
		_bgab:
			_ccege = 20
			goto _gdgfa
		_acffe:
			_ccege = 21
			goto _gdgfa
		_fafc:
			_ccege = 22
			goto _gdgfa
		_gedad:
			_ccege = 23
			goto _gdgfa
		_adgca:
			_ccege = 24
			goto _gdgfa
		_gcca:
			_ccege = 25
			goto _gdgfa
		_gdad:
			_ccege = 26
			goto _gdgfa
		_gecfg:
			_ccege = 27
			goto _gdgfa
		_ddedg:
			_ccege = 28
			goto _gdgfa
		_cabcf:
			_ccege = 29
			goto _gdgfa
		_efbd:
			_ccege = 30
			goto _egfgc
		_gaca:
			_ccege = 30
			goto _bcgf
		_cbdce:
			_ccege = 30
			goto _fcbb
		_bgea:
			_ccege = 30
			goto _feefg
		_eebg:
			_ccege = 30
			goto _fgbcb
		_cggfc:
			_ccege = 30
			goto _fddg
		_efgcg:
			_ccege = 30
			goto _dcefdb
		_fbad:
			_ccege = 30
			goto _cadaa
		_deca:
			_ccege = 30
			goto _aecg
		_fcefd:
			_ccege = 30
			goto _ddced
		_febcb:
			_ccege = 30
			goto _ddabd
		_ecdc:
			_ccege = 30
			goto _fbeaa
		_cbbad:
			_ccege = 30
			goto _debd
		_decag:
			_ccege = 30
			goto _efcad
		_egaed:
			_ccege = 30
			goto _bbabc
		_fgbcc:
			_ccege = 30
			goto _cgea
		_afgc:
			_ccege = 30
			goto _dbde
		_ebagc:
			_ccege = 30
			goto _bgbab
		_ffdae:
			_ccege = 30
			goto _aggbd
		_aaaec:
			_ccege = 30
			goto _dgbaa
		_bbgff:
			_ccege = 30
			goto _eecg
		_gafc:
			_ccege = 30
			goto _ecee
		_fgada:
			_ccege = 30
			goto _deee
		_cebag:
			_ccege = 30
			goto _fdeb
		_dbfgb:
			_ccege = 30
			goto _cdfa
		_afec:
			_ccege = 30
			goto _fffb
		_egcb:
			_ccege = 30
			goto _dede
		_gfgd:
			_ccege = 30
			goto _gcadc
		_gbfde:
			_ccege = 30
			goto _gfde
		_adgb:
			_ccege = 30
			goto _ggecb
		_cccee:
			_ccege = 30
			goto _fabac
		_ccga:
			_ccege = 30
			goto _baee
		_beagf:
			_ccege = 30
			goto _bgffd
		_ebcfa:
			_ccege = 31
			goto _aegc
		_fabbc:
			_ccege = 32
			goto _gdgfa
		_affba:
			_ccege = 33
			goto _dbec
		_gceac:
			_ccege = 34
			goto _gdgfa
		_bbadd:
			_ccege = 35
			goto _caedg
		_gbfdee:
			_ccege = 36
			goto _caedg
		_edbg:
			_ccege = 37
			goto _caedg
		_ggccg:
			_ccege = 38
			goto _gdgfa
		_agcfc:
			_ccege = 39
			goto _gdgfa
		_aebf:
			_ccege = 40
			goto _cbgaa
		_gdbgd:
			_ccege = 41
			goto _gdgfa
		_bbagd:
			_ccege = 42
			goto _cbgaa
		_ceacc:
			_ccege = 43
			goto _dbec
		_feec:
			_ccege = 44
			goto _cbgaa
		_cdadc:
			_ccege = 44
			goto _bedfb
		_fgcff:
			_ccege = 44
			goto _abefdd
		_cgdcg:
			_ccege = 45
			goto _cbgaa
		_gacd:
			_ccege = 46
			goto _cbgaa
		_bcfg:
			_ccege = 47
			goto _cbgaa
		_bbeb:
			_ccege = 48
			goto _cbgaa
		_bcge:
			_ccege = 49
			goto _cbgaa
		_cdce:
			_ccege = 50
			goto _cbgaa
		_cegb:
			_ccege = 51
			goto _cbgaa
		_aebfe:
			_ccege = 52
			goto _cbgaa
		_cebc:
			_ccege = 53
			goto _cbgaa
		_ebeac:
			_ccege = 54
			goto _cbgaa
		_dcaed:
			_ccege = 55
			goto _cbgaa
		_gddb:
			_ccege = 56
			goto _cbgaa
		_gdcgd:
			_ccege = 57
			goto _cbgaa
		_ecdce:
			_ccege = 58
			goto _cbgaa
		_ggcdc:
			_ccege = 59
			goto _cbgaa
		_gcfcc:
			_ccege = 59
			goto _edecb
		_fcbb:
			_gfgb = 3
			goto _gcgf
		_feefg:
			_gfgb = 5
			goto _gcgf
		_bcgf:
			_gfgb = 7
			goto _gcgf
		_fgbcb:
			_gfgb = 9
			goto _gcgf
		_dcefdb:
			_gfgb = 11
			goto _gcgf
		_bgffd:
			_gfgb = 13
			goto _gcgf
		_cadaa:
			_gfgb = 15
			goto _gcgf
		_dgbaa:
			_gfgb = 17
			goto _gcgf
		_eecg:
			_gfgb = 19
			goto _gcgf
		_aecg:
			_gfgb = 21
			goto _gcgf
		_ddced:
			_gfgb = 23
			goto _gcgf
		_fbeaa:
			_gfgb = 25
			goto _gcgf
		_efcad:
			_gfgb = 27
			goto _gcgf
		_ddabd:
			_gfgb = 29
			goto _gcgf
		_bbabc:
			_gfgb = 31
			goto _gcgf
		_aggbd:
			_gfgb = 33
			goto _gcgf
		_bgbab:
			_gfgb = 35
			goto _gcgf
		_gcadc:
			_gfgb = 37
			goto _gcgf
		_fabac:
			_gfgb = 39
			goto _gcgf
		_gfde:
			_gfgb = 41
			goto _gcgf
		_cgea:
			_gfgb = 43
			goto _gcgf
		_dbde:
			_gfgb = 45
			goto _gcgf
		_debd:
			_gfgb = 47
			goto _gcgf
		_fffb:
			_gfgb = 49
			goto _gcgf
		_fdeb:
			_gfgb = 51
			goto _gcgf
		_deee:
			_gfgb = 53
			goto _gcgf
		_cdfa:
			_gfgb = 55
			goto _gcgf
		_baee:
			_gfgb = 57
			goto _gcgf
		_ecee:
			_gfgb = 59
			goto _gcgf
		_dede:
			_gfgb = 61
			goto _gcgf
		_ggecb:
			_gfgb = 63
			goto _gcgf
		_fddg:
			_gfgb = 65
			goto _gcgf
		_egfgc:
			_gfgb = 67
			goto _gcgf
		_bedfb:
			_gfgb = 72
			goto _gcgf
		_caedg:
			_gfgb = 75
			goto _gcgf
		_dbec:
			_gfgb = 78
			goto _gcgf
		_abefdd:
			_gfgb = 81
			goto _gcgf
		_edecb:
			_gfgb = 84
			goto _gcgf
		_cbgaa:
			_gfgb = 87
			goto _gcgf
		_aegc:
			_gfgb = 90
			goto _gcgf
		_gcgf:
			_dbad = uint(_ebee[_gfgb])
			_gfgb++
			for ; _dbad > 0; _dbad-- {
				_gfgb++
				switch _ebee[_gfgb-1] {
				case 3:
					_bfcbg = _fddfd + 1
				case 4:
					_gebed = 1
				case 5:
					_gebed = 2
				case 6:
					_gebed = 3
				case 7:
					_gebed = 4
				case 8:
					_gebed = 11
				case 9:
					_gebed = 14
				case 10:
					_gebed = 15
				case 11:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_gggab, _aaag[_dfdf:_bfcbg])
					}
				case 12:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_adaag, _aaag[_dfdf:_bfcbg])
					}
				case 13:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_cfdg, _aaag[_dfdf:_bfcbg-1])
					}
				case 14:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_cfdg, _aaag[_dfdf+1:_bfcbg-2])
					}
				case 15:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_eeeg, _aaag[_dfdf:_bfcbg-1])
					}
				case 16:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_eeeg, _aaag[_dfdf:_bfcbg-1])
					}
				case 17:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_cdae, _aaag[_dfdf:_bfcbg])
					}
				case 18:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_egfcf, _aaag[_dfdf:_bfcbg])
					}
				case 19:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_gdac, _aaag[_dfdf:_bfcbg])
					}
				case 20:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_eabc, _aaag[_dfdf:_bfcbg])
					}
				case 21:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_gcgg, _aaag[_dfdf:_bfcbg])
					}
				case 22:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_ggded, _aaag[_dfdf:_bfcbg])
					}
				case 23:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_acec, _aaag[_dfdf:_bfcbg])
					}
				case 24:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_agbc, _aaag[_dfdf:_bfcbg])
					}
				case 25:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_cebaa, _aaag[_dfdf:_bfcbg])
					}
				case 26:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_cecbf, _aaag[_dfdf:_bfcbg])
					}
				case 27:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_gecg, _aaag[_dfdf:_bfcbg])
					}
				case 28:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_cccde, _aaag[_dfdf:_bfcbg])
					}
				case 29:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_gbcfb, _aaag[_dfdf:_bfcbg])
					}
				case 30:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_dgfg, _aaag[_dfdf:_bfcbg])
					}
				case 31:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_ebac, _aaag[_dfdf:_bfcbg])
					}
				case 32:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_debac, _aaag[_dfdf:_bfcbg])
					}
				case 33:
					_bfcbg = _fddfd + 1
					{
						_egbbbf.emit(_agbgc, _aaag[_dfdf:_bfcbg])
					}
				case 34:
					_bfcbg = _fddfd
					_fddfd--
					{
						_egbbbf.emit(_gaeff, _aaag[_dfdf:_bfcbg])
					}
				case 35:
					_bfcbg = _fddfd
					_fddfd--
					{
						_egbbbf.emit(_eafe, _aaag[_dfdf:_bfcbg])
					}
				case 36:
					_bfcbg = _fddfd
					_fddfd--
					{
						_egbbbf.emit(_bggae, _aaag[_dfdf:_bfcbg])
					}
				case 37:
					_bfcbg = _fddfd
					_fddfd--
					{
						_egbbbf.emit(_degcb, _aaag[_dfdf:_bfcbg])
					}
				case 38:
					_bfcbg = _fddfd
					_fddfd--
					{
						_egbbbf.emit(_ecbdf, _aaag[_dfdf:_bfcbg])
					}
				case 39:
					_bfcbg = _fddfd
					_fddfd--
					{
						_egbbbf.emit(_eeffc, _aaag[_dfdf+1:_bfcbg-1])
					}
				case 40:
					_bfcbg = _fddfd
					_fddfd--
					{
						_egbbbf.emit(_cbcfe, _aaag[_dfdf:_bfcbg])
					}
				case 41:
					_bfcbg = _fddfd
					_fddfd--
					{
						_egbbbf.emit(_dgbf, _aaag[_dfdf:_bfcbg])
					}
				case 42:
					_fddfd = (_bfcbg) - 1
					{
						_egbbbf.emit(_gaeff, _aaag[_dfdf:_bfcbg])
					}
				case 43:
					switch _gebed {
					case 0:
						{
							_ccege = 0
							goto _gdgfa
						}
					case 1:
						{
							_fddfd = (_bfcbg) - 1
							_egbbbf.emit(_efaag, _aaag[_dfdf:_bfcbg])
						}
					case 2:
						{
							_fddfd = (_bfcbg) - 1
							_egbbbf.emit(_gaeff, _aaag[_dfdf:_bfcbg])
						}
					case 3:
						{
							_fddfd = (_bfcbg) - 1
							_egbbbf.emit(_eafe, _aaag[_dfdf:_bfcbg])
						}
					case 4:
						{
							_fddfd = (_bfcbg) - 1
							_egbbbf.emit(_ebfg, _aaag[_dfdf:_bfcbg])
						}
					case 11:
						{
							_fddfd = (_bfcbg) - 1
							_egbbbf.emit(_cdfcd, _aaag[_dfdf:_bfcbg])
						}
					case 14:
						{
							_fddfd = (_bfcbg) - 1
							_egbbbf.emit(_ecbdf, _aaag[_dfdf:_bfcbg])
						}
					case 15:
						{
							_fddfd = (_bfcbg) - 1
							_egbbbf.emit(_eeffc, _aaag[_dfdf+1:_bfcbg-1])
						}
					}
				}
			}
			goto _gdgfa
		_gdgfa:
			_gfgb = int(_egaa[_ccege])
			_dbad = uint(_ebee[_gfgb])
			_gfgb++
			for ; _dbad > 0; _dbad-- {
				_gfgb++
				switch _ebee[_gfgb-1] {
				case 0:
					_dfdf = 0
				case 1:
					_gebed = 0
				}
			}
			if _ccege == 0 {
				goto _fcfcf
			}
			if _fddfd++; _fddfd != _egdaa {
				goto _bgbcc
			}
		_eddec:
			{
			}
			if _fddfd == _cbfef {
				switch _ccege {
				case 1:
					goto _efbd
				case 2:
					goto _efbd
				case 31:
					goto _gafc
				case 14:
					goto _efbd
				case 15:
					goto _efbd
				case 32:
					goto _fgada
				case 17:
					goto _efbd
				case 33:
					goto _cebag
				case 18:
					goto _efbd
				case 19:
					goto _efbd
				case 34:
					goto _dbfgb
				case 35:
					goto _afec
				case 36:
					goto _afec
				case 23:
					goto _cggfc
				case 37:
					goto _afec
				case 38:
					goto _egcb
				case 39:
					goto _adgb
				case 40:
					goto _ccga
				case 41:
					goto _ccga
				case 42:
					goto _ccga
				case 43:
					goto _cebag
				case 44:
					goto _efbd
				case 45:
					goto _ccga
				case 46:
					goto _ccga
				case 47:
					goto _ccga
				case 48:
					goto _ccga
				case 49:
					goto _ccga
				case 50:
					goto _ccga
				case 51:
					goto _ccga
				case 52:
					goto _ccga
				case 53:
					goto _ccga
				case 54:
					goto _ccga
				case 55:
					goto _ccga
				case 56:
					goto _ccga
				case 57:
					goto _ccga
				case 58:
					goto _ccga
				case 59:
					goto _efbd
				}
			}
		_fcfcf:
			{
			}
		}
		if _dfdf > 0 {
			copy(_aaag[0:], _aaag[_dfdf:])
		}
	}
	_ = _cbfef
	if _ccege == _cbfa {
		_egbbbf.emit(_gbae, nil)
	}
	close(_egbbbf._cgdbfa)
}

// String returns a string representation of a named range.
func (_gdbab NamedRangeRef) String() string { return _gdbab._ceaf }

// Db implements the Excel DB function.
func Db(args []Result) Result {
	_ccaa := len(args)
	if _ccaa != 4 && _ccaa != 5 {
		return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072\u0020\u006f\u0072 \u0066\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0063\u006f\u0073\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_efdd := args[0].ValueNumber
	if _efdd < 0 {
		return MakeErrorResultType(ErrorTypeNum, "D\u0042\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0063\u006fs\u0074\u0020\u0074\u006f\u0020\u0062\u0065 \u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069v\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_acagdg := args[1].ValueNumber
	if _acagdg < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062e\u0020\u006e\u006f\u006e\u0020n\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069\u0066\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_deaa := args[2].ValueNumber
	if _deaa <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0042\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006c\u0069\u0066\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("D\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_efgb := args[3].ValueNumber
	if _efgb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0042\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u0076\u0065")
	}
	if _efgb-_deaa > 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u0042")
	}
	_afg := 12.0
	if _ccaa == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006do\u006e\u0074\u0068\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_afg = args[4].ValueNumber
		if _afg < 1 || _afg > 12 {
			return MakeErrorResultType(ErrorTypeNum, "\u0044B\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u006do\u006e\u0074\u0068\u0020\u0074\u006f\u0020\u0062\u0065 i\u006e\u0020\u0072a\u006e\u0067e\u0020\u006f\u0066\u0020\u0031\u0020a\u006e\u0064 \u0031\u0032")
		}
	}
	if _afg == 12 && _efgb > _deaa {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063or\u0072\u0065\u0063\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u0042")
	}
	if _acagdg >= _efdd {
		return MakeNumberResult(0)
	}
	_gccgg := 1 - _cc.Pow(_acagdg/_efdd, 1/_deaa)
	_gccgg = float64(int(_gccgg*1000+0.5)) / 1000
	_gdfc := _efdd * _gccgg * _afg / 12
	if _efgb == 1 {
		return MakeNumberResult(_gdfc)
	}
	_gdgg := _gdfc
	_edea := 0.0
	_beac := _deaa
	if _beac > _efgb {
		_beac = _efgb
	}
	for _ebaa := 2.0; _ebaa <= _beac; _ebaa++ {
		_edea = (_efdd - _gdgg) * _gccgg
		_gdgg += _edea
	}
	if _efgb > _deaa {
		return MakeNumberResult((_efdd - _gdgg) * _gccgg * (12 - _afg) / 12)
	}
	return MakeNumberResult(_edea)
}

// Reference returns a string reference value to a sheet.
func (_bbgfe SheetPrefixExpr) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeSheet, Value: _bbgfe._bgfda}
}

// SumIfs implements the SUMIFS function.
func SumIfs(args []Result) Result {
	_afdf := _eged(args, true, "\u0053\u0055\u004d\u0049\u0046\u0053")
	if _afdf.Type != ResultTypeEmpty {
		return _afdf
	}
	_aeea := _ecbg(args[1:])
	_ebgbe := 0.0
	_gbdc := _aadag(args[0])
	for _, _bagd := range _aeea {
		_ebgbe += _gbdc[_bagd._aaba][_bagd._gbfb].ValueNumber
	}
	return MakeNumberResult(float64(_ebgbe))
}
func _fcafc(_bcbg int) string {
	if _bcbg >= 0 && _bcbg < len(_bgdbd) {
		if _bgdbd[_bcbg] != "" {
			return _bgdbd[_bcbg]
		}
	}
	return _e.Sprintf("\u0073\u0074\u0061\u0074\u0065\u002d\u0025\u0076", _bcbg)
}

// Edate is an implementation of the Excel EDATE() function.
func Edate(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0045\u0044\u0041\u0054E\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0077o\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
	}
	_fcad := args[1].ValueNumber
	_bbec := args[0]
	var _dba float64
	switch _bbec.Type {
	case ResultTypeEmpty:
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
	case ResultTypeNumber:
		_dba = _bbec.ValueNumber
	case ResultTypeString:
		_fcd := DateValue([]Result{args[0]})
		if _fcd.Type == ResultTypeError {
			return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
		}
		_dba = _fcd.ValueNumber
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
	}
	_bgf := _bgee(_dba)
	_gaf := _bgf.AddDate(0, int(_fcad), 0)
	_ecad, _cbc, _dfaa := _gaf.Date()
	_aebb := _bad(_ecad, int(_cbc), _dfaa)
	if _aebb < 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0045\u0044A\u0054\u0045")
	}
	return MakeNumberResult(_aebb)
}

// String returns an empty string for Error.
func (_bb Error) String() string { return "" }

// Dollarde implements the Excel DOLLARDE function.
func Dollarde(args []Result) Result {
	_bdcd, _aeee, _degb := _gcga(args, "\u0044\u004f\u004c\u004c\u0041\u0052\u0044\u0045")
	if _degb.Type == ResultTypeError {
		return _degb
	}
	if _aeee < 1 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0044\u004f\u004c\u004c\u0041\u0052\u0044\u0045\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0072a\u0063t\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u006f\u0072 \u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u0031")
	}
	if _bdcd == 0 {
		return MakeNumberResult(0)
	}
	_cecba := _bdcd < 0
	if _cecba {
		_bdcd = -_bdcd
	}
	_gbdb := args[0].Value()
	_accg := _ga.Split(_gbdb, "\u002e")
	_cace := float64(int(_bdcd))
	_bdac := _accg[1]
	_eeacg := len(_bdac)
	_aggad := int(_cc.Log10(_aeee)) + 1
	_cfce := float64(_aggad - _eeacg)
	_ceac, _abaf := _afc.ParseFloat(_bdac, 64)
	if _abaf != nil {
		return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u006fr \u0044\u004f\u004cL\u0041R\u0044\u0045")
	}
	_ceac *= _cc.Pow(10, _cfce)
	_cda := _cace + _ceac/_aeee
	if _cecba {
		_cda = -_cda
	}
	return MakeNumberResult(_cda)
}

const _aadce int = 30

// Eval evaluates a vertical range with prefix returning a list of results or an error.
func (_dgag PrefixVerticalRange) Eval(ctx Context, ev Evaluator) Result {
	_bebec := _dgag._fcga.Reference(ctx, ev)
	switch _bebec.Type {
	case ReferenceTypeSheet:
		if _dcbb(_bebec, ctx) {
			return MakeErrorResultType(ErrorTypeName, _e.Sprintf("\u0053h\u0065e\u0074\u0020\u0025\u0073\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064", _bebec.Value))
		}
		_fbedd := _dgag.verticalRangeReference(_bebec.Value)
		if _cgaade, _becdc := ev.GetFromCache(_fbedd); _becdc {
			return _cgaade
		}
		_cadae := ctx.Sheet(_bebec.Value)
		_edaa, _gcdb := _cfaeb(_cadae, _dgag._agfae, _dgag._fefae)
		_eeea := _gebc(_cadae, ev, _edaa, _gcdb)
		ev.SetCache(_fbedd, _eeea)
		return _eeea
	default:
		return MakeErrorResult(_e.Sprintf("\u006e\u006f\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0066\u006f\u0072\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _bebec.Type))
	}
}

var _egaa []byte = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 69, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

func init() {
	_aagf = _d.New(_d.NewSource(_eb.Now().UnixNano()))
	RegisterFunction("\u0041\u0042\u0053", _gfcf("\u0041\u0053\u0049\u004e", _cc.Abs))
	RegisterFunction("\u0041\u0043\u004f\u0053", _gfcf("\u0041\u0053\u0049\u004e", _cc.Acos))
	RegisterFunction("\u0041\u0043\u004fS\u0048", _gfcf("\u0041\u0053\u0049\u004e", _cc.Acosh))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0041\u0043\u004f\u0054", _gfcf("\u0041\u0043\u004f\u0054", func(_dcaf float64) float64 { return _cc.Pi/2 - _cc.Atan(_dcaf) }))
	RegisterFunction("_\u0078\u006c\u0066\u006e\u002e\u0041\u0043\u004f\u0054\u0048", _gfcf("\u0041\u0043\u004fT\u0048", func(_eeba float64) float64 { return _cc.Atanh(1 / _eeba) }))
	RegisterFunction("\u005f\u0078\u006cf\u006e\u002e\u0041\u0052\u0041\u0042\u0049\u0043", Arabic)
	RegisterFunction("\u0041\u0053\u0049\u004e", _gfcf("\u0041\u0053\u0049\u004e", _cc.Asin))
	RegisterFunction("\u0041\u0053\u0049N\u0048", _gfcf("\u0041\u0053\u0049N\u0048", _cc.Asinh))
	RegisterFunction("\u0041\u0054\u0041\u004e", _gfcf("\u0041\u0054\u0041\u004e", _cc.Atan))
	RegisterFunction("\u0041\u0054\u0041N\u0048", _gfcf("\u0041\u0054\u0041N\u0048", _cc.Atanh))
	RegisterFunction("\u0041\u0054\u0041N\u0032", Atan2)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0042\u0041\u0053\u0045", Base)
	RegisterFunction("\u0043E\u0049\u004c\u0049\u004e\u0047", Ceiling)
	RegisterFunction("\u005fx\u006cf\u006e\u002e\u0043\u0045\u0049L\u0049\u004eG\u002e\u004d\u0041\u0054\u0048", CeilingMath)
	RegisterFunction("_\u0078\u006c\u0066\u006e.C\u0045I\u004c\u0049\u004e\u0047\u002eP\u0052\u0045\u0043\u0049\u0053\u0045", CeilingPrecise)
	RegisterFunction("\u0043\u004f\u004d\u0042\u0049\u004e", Combin)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0043\u004f\u004d\u0042\u0049\u004e\u0041", Combina)
	RegisterFunction("\u0043\u004f\u0053", _gfcf("\u0043\u004f\u0053", _cc.Cos))
	RegisterFunction("\u0043\u004f\u0053\u0048", _gfcf("\u0043\u004f\u0053\u0048", _cc.Cosh))
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0043\u004fT", _cbgbd("\u0043\u004f\u0054", _cc.Tan))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0043\u004f\u0054\u0048", _cbgbd("\u0043\u004f\u0054\u0048", _cc.Tanh))
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0043\u0053C", _cbgbd("\u0043\u0053\u0043", _cc.Sin))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0043\u0053\u0043\u0048", _cbgbd("\u0043\u0053\u0043", _cc.Sinh))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0044\u0045\u0043\u0049\u004d\u0041\u004c", Decimal)
	RegisterFunction("\u0044E\u0047\u0052\u0045\u0045\u0053", Degrees)
	RegisterFunction("\u0045\u0056\u0045\u004e", Even)
	RegisterFunction("\u0045\u0058\u0050", _gfcf("\u0045\u0058\u0050", _cc.Exp))
	RegisterFunction("\u0046\u0041\u0043\u0054", Fact)
	RegisterFunction("\u0046\u0041\u0043\u0054\u0044\u004f\u0055\u0042\u004c\u0045", FactDouble)
	RegisterFunction("\u0046\u004c\u004fO\u0052", Floor)
	RegisterFunction("\u005f\u0078l\u0066\u006e\u002eF\u004c\u004f\u004f\u0052\u002e\u004d\u0041\u0054\u0048", FloorMath)
	RegisterFunction("\u005f\u0078\u006c\u0066n.\u0046\u004c\u004f\u004f\u0052\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045", FloorPrecise)
	RegisterFunction("\u0047\u0043\u0044", GCD)
	RegisterFunction("\u0049\u004e\u0054", Int)
	RegisterFunction("I\u0053\u004f\u002e\u0043\u0045\u0049\u004c\u0049\u004e\u0047", CeilingPrecise)
	RegisterFunction("\u004c\u0043\u004d", LCM)
	RegisterFunction("\u004c\u004e", _gfcf("\u004c\u004e", _cc.Log))
	RegisterFunction("\u004c\u004f\u0047", Log)
	RegisterFunction("\u004c\u004f\u00471\u0030", _gfcf("\u004c\u004f\u00471\u0030", _cc.Log10))
	RegisterFunction("\u004dD\u0045\u0054\u0045\u0052\u004d", MDeterm)
	RegisterFunction("\u004d\u004f\u0044", Mod)
	RegisterFunction("\u004d\u0052\u004f\u0055\u004e\u0044", Mround)
	RegisterFunction("M\u0055\u004c\u0054\u0049\u004e\u004f\u004d\u0049\u0041\u004c", Multinomial)
	RegisterFunction("_\u0078\u006c\u0066\u006e\u002e\u004d\u0055\u004e\u0049\u0054", Munit)
	RegisterFunction("\u004f\u0044\u0044", Odd)
	RegisterFunction("\u0050\u0049", Pi)
	RegisterFunction("\u0050\u004f\u0057E\u0052", Power)
	RegisterFunction("\u0050R\u004f\u0044\u0055\u0043\u0054", Product)
	RegisterFunction("\u0051\u0055\u004f\u0054\u0049\u0045\u004e\u0054", Quotient)
	RegisterFunction("\u0052A\u0044\u0049\u0041\u004e\u0053", Radians)
	RegisterFunction("\u0052\u0041\u004e\u0044", Rand)
	RegisterFunction("R\u0041\u004e\u0044\u0042\u0045\u0054\u0057\u0045\u0045\u004e", RandBetween)
	RegisterFunction("\u0052\u004f\u004dA\u004e", Roman)
	RegisterFunction("\u0052\u004f\u0055N\u0044", Round)
	RegisterFunction("\u0052O\u0055\u004e\u0044\u0044\u004f\u0057N", RoundDown)
	RegisterFunction("\u0052O\u0055\u004e\u0044\u0055\u0050", RoundUp)
	RegisterFunction("\u005fx\u006c\u0066\u006e\u002e\u0053\u0045C", _cbgbd("\u0053\u0045\u0043", _cc.Cos))
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0053\u0045\u0043\u0048", _cbgbd("\u0053\u0045\u0043\u0048", _cc.Cosh))
	RegisterFunction("\u0053E\u0052\u0049\u0045\u0053\u0053\u0055M", SeriesSum)
	RegisterFunction("\u0053\u0049\u0047\u004e", Sign)
	RegisterFunction("\u0053\u0049\u004e", _gfcf("\u0053\u0049\u004e", _cc.Sin))
	RegisterFunction("\u0053\u0049\u004e\u0048", _gfcf("\u0053\u0049\u004e\u0048", _cc.Sinh))
	RegisterFunction("\u0053\u0051\u0052\u0054", _gfcf("\u0053\u0051\u0052\u0054", _cc.Sqrt))
	RegisterFunction("\u0053\u0051\u0052\u0054\u0050\u0049", _gfcf("\u0053\u0051\u0052\u0054\u0050\u0049", func(_adeab float64) float64 { return _cc.Sqrt(_adeab * _cc.Pi) }))
	RegisterFunction("\u0053\u0055\u004d", Sum)
	RegisterFunction("\u0053\u0055\u004dI\u0046", SumIf)
	RegisterFunction("\u0053\u0055\u004d\u0049\u0046\u0053", SumIfs)
	RegisterFunction("\u0053\u0055\u004d\u0050\u0052\u004f\u0044\u0055\u0043\u0054", SumProduct)
	RegisterFunction("\u0053\u0055\u004dS\u0051", SumSquares)
	RegisterFunction("\u0054\u0041\u004e", _gfcf("\u0054\u0041\u004e", _cc.Tan))
	RegisterFunction("\u0054\u0041\u004e\u0048", _gfcf("\u0054\u0041\u004e\u0048", _cc.Tanh))
	RegisterFunction("\u0054\u0052\u0055N\u0043", Trunc)
}

// Reference returns an invalid reference for Bool.
func (_ecfc Bool) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// String returns a string representation of ConstArrayExpr.
func (_fcf ConstArrayExpr) String() string { return "" }

const _edeae = 16
const _agbgc = 57376

func _cfaeb(_accda Context, _bbea, _fbddg string) (string, string) {
	_abbg := _bbea + "\u0031"
	_gcbg := _accda.LastRow(_bbea)
	_edgb := _fbddg + _afc.Itoa(_gcbg)
	return _abbg, _edgb
}

// RandBetween is an implementation of the Excel RANDBETWEEN() function that returns a random
// integer in the range specified.
func RandBetween(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0052A\u004e\u0044B\u0045\u0054\u0057\u0045E\u004e\u0028\u0029 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020tw\u006f\u0020\u006eu\u006d\u0065r\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_cgab := args[0].AsNumber()
	_eaee := args[1].AsNumber()
	if _cgab.Type != ResultTypeNumber || _eaee.Type != ResultTypeNumber {
		return MakeErrorResult("\u0052A\u004e\u0044B\u0045\u0054\u0057\u0045E\u004e\u0028\u0029 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020tw\u006f\u0020\u006eu\u006d\u0065r\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	if _eaee.ValueNumber < _cgab.ValueNumber {
		return MakeErrorResult("\u0052\u0041\u004e\u0044\u0042E\u0054\u0057\u0045\u0045\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065r")
	}
	_edbd := int64(_cgab.ValueNumber)
	_bbag := int64(_eaee.ValueNumber)
	return MakeNumberResult(float64(_aagf.Int63n(_bbag-_edbd+1) + _edbd))
}

// Reference returns an invalid reference for Negate.
func (_gdcd Negate) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// FloorMath implements _xlfn.FLOOR.MATH which rounds numbers down to the
// nearest multiple of the second argument, toward or away from zero as
// specified by the third argument.
func FloorMath(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0046\u004c\u004f\u004f\u0052\u002e\u004dA\u0054\u0048\u0028)\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u0046\u004c\u004f\u004f\u0052\u002e\u004dA\u0054\u0048\u0028)\u0020\u0061\u006c\u006co\u0077\u0073\u0020\u0061\u0074\u0020\u006d\u006f\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ggdb := args[0].AsNumber()
	if _ggdb.Type != ResultTypeNumber {
		return MakeErrorResult("f\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020FL\u004f\u004f\u0052\u002eM\u0041\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073t \u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_fgga := float64(1)
	if _ggdb.ValueNumber < 0 {
		_fgga = -1
	}
	if len(args) > 1 {
		_fbdd := args[1].AsNumber()
		if _fbdd.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061r\u0067\u0075\u006den\u0074\u0020\u0074\u006f\u0020\u0046L\u004f\u004f\u0052\u002e\u004d\u0041\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006db\u0065\u0072")
		}
		_fgga = _fbdd.ValueNumber
	}
	_fceff := float64(1)
	if len(args) > 2 {
		_baaea := args[2].AsNumber()
		if _baaea.Type != ResultTypeNumber {
			return MakeErrorResult("t\u0068\u0069\u0072\u0064\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020FL\u004f\u004f\u0052\u002eM\u0041\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073t \u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_fceff = _baaea.ValueNumber
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Floor(_ggdb.ValueNumber))
	}
	_cgee := _ggdb.ValueNumber
	_cgee, _bbga := _cc.Modf(_cgee / _fgga)
	if _bbga != 0 && _ggdb.ValueNumber < 0 && _fceff > 0 {
		_cgee++
	}
	return MakeNumberResult(_cgee * _fgga)
}

// MakeStringResult constructs a string result.
func MakeStringResult(s string) Result { return Result{Type: ResultTypeString, ValueString: s} }

// Update updates references in the Negate after removing a row/column.
func (_ffeb Negate) Update(q *_eg.UpdateQuery) Expression {
	return Negate{_effb: _ffeb._effb.Update(q)}
}

// Rand is an implementation of the Excel RAND() function that returns random
// numbers in the range [0,1).
func Rand(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("R\u0041\u004e\u0044\u0028\u0029\u0020a\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u006e\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	return MakeNumberResult(_aagf.Float64())
}

// SumSquares is an implementation of the Excel SUMSQ() function.
func SumSquares(args []Result) Result {
	_afdbc := MakeNumberResult(0)
	for _, _dggdg := range args {
		_dggdg = _dggdg.AsNumber()
		switch _dggdg.Type {
		case ResultTypeNumber:
			_afdbc.ValueNumber += _dggdg.ValueNumber * _dggdg.ValueNumber
		case ResultTypeList, ResultTypeArray:
			_cbfd := SumSquares(_dggdg.ListValues())
			if _cbfd.Type != ResultTypeNumber {
				return _cbfd
			}
			_afdbc.ValueNumber += _cbfd.ValueNumber
		case ResultTypeString:
		case ResultTypeError:
			return _dggdg
		case ResultTypeEmpty:
		default:
			return MakeErrorResult(_e.Sprintf("\u0075\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u0053\u0055\u004dS\u0051\u0055\u0041\u0052\u0045\u0053(\u0029\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079p\u0065\u0020\u0025\u0073", _dggdg.Type))
		}
	}
	return _afdbc
}
func _ecc(_bbdf, _eddf, _cadg, _fdadc float64, _ede int) float64 {
	var _dggd float64
	if _bbdf == 0 {
		_dggd = (_cadg + _fdadc) / _eddf
	} else {
		_cca := _cc.Pow(1+_bbdf, _eddf)
		if _ede == 1 {
			_dggd = (_fdadc*_bbdf/(_cca-1) + _cadg*_bbdf/(1-1/_cca)) / (1 + _bbdf)
		} else {
			_dggd = _fdadc*_bbdf/(_cca-1) + _cadg*_bbdf/(1-1/_cca)
		}
	}
	return -_dggd
}

// Reference returns a string reference value to a named range.
func (_ebef NamedRangeRef) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeNamedRange, Value: _ebef._ceaf}
}

// NewEmptyExpr constructs a new empty expression.
func NewEmptyExpr() Expression { return EmptyExpr{} }

// Eval evaluates and returns the result of an error expression.
func (_gf Error) Eval(ctx Context, ev Evaluator) Result { return MakeErrorResult(_gf._gbb) }

// NewHorizontalRange constructs a new full rows range.
func NewHorizontalRange(v string) Expression {
	_cbga := _ga.Split(v, "\u003a")
	if len(_cbga) != 2 {
		return nil
	}
	_ddedd, _ := _afc.Atoi(_cbga[0])
	_cadde, _ := _afc.Atoi(_cbga[1])
	if _ddedd > _cadde {
		_ddedd, _cadde = _cadde, _ddedd
	}
	return HorizontalRange{_fcbd: _ddedd, _fbcg: _cadde}
}

var ReferenceInvalid = Reference{Type: ReferenceTypeInvalid}

const _bed = _bddc + "\u0020\u0028\u0028[0\u002d\u0039\u005d\u0029\u002b\u0029\u002c\u0020\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029"

func LexReader(r _ee.Reader) chan *node {
	_bfcbf := NewLexer()
	go _bfcbf.lex(r)
	return _bfcbf._cgdbfa
}
func _aabce(_beace []Result, _gfcad countMode) float64 {
	_gec := 0.0
	for _, _gaafc := range _beace {
		switch _gaafc.Type {
		case ResultTypeNumber:
			if _gfcad == _cbfbc || (_gfcad == _fagga && !_gaafc.IsBoolean) {
				_gec++
			}
		case ResultTypeList, ResultTypeArray:
			_gec += _aabce(_gaafc.ListValues(), _gfcad)
		case ResultTypeString:
			if _gfcad == _cbfbc {
				_gec++
			}
		case ResultTypeEmpty:
			if _gfcad == _cdgce {
				_gec++
			}
		}
	}
	return _gec
}

const _egad = "\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0054\u0079\u0070\u0065\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0054\u0079\u0070\u0065\u0043\u0065\u006c\u006c\u0052\u0065\u0066\u0065r\u0065\u006ec\u0065\u0054\u0079\u0070e\u004e\u0061\u006d\u0065\u0064\u0052\u0061\u006e\u0067\u0065R\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0054y\u0070\u0065\u0052\u0061\u006e\u0067\u0065\u0052\u0065\u0066e\u0072\u0065\u006ec\u0065\u0054\u0079\u0070\u0065\u0053\u0068e\u0065\u0074"

func _fbcc(_bgba, _abg, _bceef, _cfcge, _bbdb float64, _cecb int) Result {
	_abca, _acbc := _fcfa(_bgba, _abg, _cecb)
	if _acbc.Type == ResultTypeError {
		return _acbc
	}
	_eeed, _ccba := _cefd(_bgba, _abg, int(_bbdb), _cecb)
	if _ccba.Type == ResultTypeError {
		return _ccba
	}
	_ggbe := 0.0
	_fdbc := 0.0
	_bceef *= 100 / _bbdb
	_cfcge /= _bbdb
	_cfcge++
	_fdf := _abca*_bbdb - _eeed
	for _ccfg := 1.0; _ccfg < _eeed; _ccfg++ {
		_cbdc := _ccfg + _fdf
		_gead := _bceef / _cc.Pow(_cfcge, _cbdc)
		_fdbc += _gead
		_ggbe += _cbdc * _gead
	}
	_ffdg := (_bceef + 100) / _cc.Pow(_cfcge, _eeed+_fdf)
	_fdbc += _ffdg
	_ggbe += (_eeed + _fdf) * _ffdg
	_ggbe /= _fdbc
	_ggbe /= _bbdb
	return MakeNumberResult(_ggbe)
}

// Xor is an implementation of the Excel XOR() function and takes a variable
// number of arguments. It's odd to say the least.  If any argument is numeric,
// it returns true if the number of non-zero numeric arguments is odd and false
// otherwise.  If no argument is numeric, it returns an error.
func Xor(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0058\u004f\u0052 r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061t\u0020l\u0065a\u0073t\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gdee := 0
	_bace := false
	for _, _gdeg := range args {
		switch _gdeg.Type {
		case ResultTypeList, ResultTypeArray:
			_dbfdbf := Xor(_gdeg.ListValues())
			if _dbfdbf.Type == ResultTypeError {
				return _dbfdbf
			}
			if _dbfdbf.ValueNumber != 0 {
				_gdee++
			}
			_bace = true
		case ResultTypeNumber:
			if _gdeg.ValueNumber != 0 {
				_gdee++
			}
			_bace = true
		case ResultTypeString:
		case ResultTypeError:
			return _gdeg
		default:
			return MakeErrorResult("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0061\u0072\u0067u\u006de\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u0058\u004f\u0052")
		}
	}
	if !_bace {
		return MakeErrorResult("\u0058\u004f\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0069n\u0070\u0075\u0074")
	}
	return MakeBoolResult(_gdee%2 != 0)
}

// Update returns the same object as updating sheet references does not affect SheetPrefixExpr.
func (_dbcc SheetPrefixExpr) Update(q *_eg.UpdateQuery) Expression { return _dbcc }

// Context is a formula execution context.  Formula evaluation uses the context
// to retreive information from sheets.
type Context interface {

	// Cell returns the result of evaluating a cell.
	Cell(_gg string, _eac Evaluator) Result

	// Sheet returns an evaluation context for a given sheet name.  This is used
	// when evaluating cells that pull data from other sheets (e.g. ='Sheet 2'!A1).
	Sheet(_bab string) Context

	// GetEpoch returns the time epoch of the context's Workbook.
	GetEpoch() _eb.Time

	// GetFilename returns the full filename of the context's Workbook.
	GetFilename() string

	// GetWidth returns a worksheet's column width.
	GetWidth(_aab int) float64

	// GetFormat returns a cell's format.
	GetFormat(_fcfd string) string

	// GetLabelPrefix returns cell's label prefix dependent on cell horizontal alignment.
	GetLabelPrefix(_cfa string) string

	// GetFormat returns if cell is protected.
	GetLocked(_dgf string) bool

	// HasFormula returns if cell contains formula.
	HasFormula(_fcfb string) bool

	// IsBool returns if cell contains boolean value.
	IsBool(_fgb string) bool

	// IsDBCS returns if workbook default language is among DBCS.
	IsDBCS() bool

	// LastColumn returns the name of last column which contains data in range of context sheet's given rows.
	LastColumn(_acd, _ecfb int) string

	// LastRow returns the name of last row which contains data in range of context sheet's given columns.
	LastRow(_cbgg string) int

	// SetLocked returns sets cell's protected attribute.
	SetLocked(_dffc string, _fffc bool)

	// NamedRange returns a named range.
	NamedRange(_gcc string) Reference

	// SetOffset is used so that the Context can evaluate cell references
	// differently when they are not absolute (e.g. not like '$A$5').  See the
	// shared formula support in Cell for usage.
	SetOffset(_fad, _afb uint32)
}

// Odd is an implementation of the Excel ODD() that rounds a number to the
// nearest odd integer.
func Odd(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("O\u0044\u0044\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006ee\u0020\u0061\u0072g\u0075m\u0065\u006e\u0074")
	}
	_adba := args[0].AsNumber()
	if _adba.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_geda := _cc.Signbit(_adba.ValueNumber)
	_gaag, _eaaa := _cc.Modf((_adba.ValueNumber - 1) / 2)
	_dgbaf := _gaag*2 + 1
	if _eaaa != 0 {
		if !_geda {
			_dgbaf += 2
		} else {
			_dgbaf -= 2
		}
	}
	return MakeNumberResult(_dgbaf)
}

// Update returns the same object as updating sheet references does not affect String.
func (_ecbf String) Update(q *_eg.UpdateQuery) Expression { return _ecbf }

type durationArgs struct {
	_bdce float64
	_fbda float64
	_agd  float64
	_cacc float64
	_efaf float64
	_bebb int
}

// BinaryExpr is a binary expression.
type BinaryExpr struct {
	_gb, _gba Expression
	_fc       BinOpType
}

// Mduration implements the Excel MDURATION function.
func Mduration(args []Result) Result {
	_ddcb, _bebd := _cgbf(args, "\u004dD\u0055\u0052\u0041\u0054\u0049\u004fN")
	if _bebd.Type == ResultTypeError {
		return _bebd
	}
	_gdc := _ddcb._bdce
	_deebd := _ddcb._fbda
	_bbgde := _ddcb._agd
	_geec := _ddcb._cacc
	_fead := _ddcb._efaf
	_ebca := _ddcb._bebb
	_ebbf := _fbcc(_gdc, _deebd, _bbgde, _geec, _fead, _ebca)
	if _ebbf.Type == ResultTypeError {
		return _ebbf
	}
	_cdab := _ebbf.ValueNumber / (1.0 + _geec/_fead)
	return MakeNumberResult(_cdab)
}

// Eval evaluates and returns a number.
func (_aeagg Number) Eval(ctx Context, ev Evaluator) Result { return MakeNumberResult(_aeagg._fbfe) }

// Atan2 implements the Excel ATAN2 function.  It accepts two numeric arguments,
// and the arguments are (x,y), reversed from normal to match Excel's behaviour.
func Atan2(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0041\u0054\u0041\u004e2\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0077o\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	_dbdg := args[0].AsNumber()
	_bfbd := args[1].AsNumber()
	if _dbdg.Type == ResultTypeNumber && _bfbd.Type == ResultTypeNumber {
		_eecbc := _cc.Atan2(_bfbd.ValueNumber, _dbdg.ValueNumber)
		if _eecbc != _eecbc {
			return MakeErrorResult("\u0041T\u0041N\u0032\u0020\u0072\u0065\u0074u\u0072\u006ee\u0064\u0020\u004e\u0061\u004e")
		}
		return MakeNumberResult(_eecbc)
	}
	for _, _cdda := range []ResultType{_dbdg.Type, _bfbd.Type} {
		switch _cdda {
		case ResultTypeList, ResultTypeString:
			return MakeErrorResult("\u0041\u0054\u0041\u004e\u0032\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		case ResultTypeError:
			return _dbdg
		default:
			return MakeErrorResult(_e.Sprintf("\u0075\u006e\u0068an\u0064\u006c\u0065\u0064\u0020\u0041\u0054\u0041\u004e2\u0028)\u0020a\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _cdda))
		}
	}
	return MakeErrorResult("u\u006e\u0068\u0061\u006e\u0064\u006ce\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0066o\u0072\u0020\u0041T\u0041N\u0032\u0028\u0029")
}

// Munit is an implementation of the Excel MUNIT function that returns an
// identity matrix.
func Munit(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004d\u0055\u004eIT\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073 \u006fn\u0065 \u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0069\u006e\u0070\u0075\u0074")
	}
	_gfbd := args[0].AsNumber()
	if _gfbd.Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u0055\u004eIT\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073 \u006fn\u0065 \u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0069\u006e\u0070\u0075\u0074")
	}
	_fbbc := int(_gfbd.ValueNumber)
	_cbdef := make([][]Result, 0, _fbbc)
	for _fagc := 0; _fagc < _fbbc; _fagc++ {
		_cdbag := make([]Result, _fbbc)
		for _cdbc := 0; _cdbc < _fbbc; _cdbc++ {
			if _fagc == _cdbc {
				_cdbag[_cdbc] = MakeNumberResult(1.0)
			} else {
				_cdbag[_cdbc] = MakeNumberResult(0.0)
			}
		}
		_cbdef = append(_cbdef, _cdbag)
	}
	return MakeArrayResult(_cbdef)
}

// Vdb implements the Excel VDB function.
func Vdb(args []Result) Result {
	_bcad := len(args)
	if _bcad < 5 || _bcad > 7 {
		return MakeErrorResult("\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020b\u0065\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u0062\u0065\u0074\u0077\u0065\u0065\u006e\u0020\u0066\u0069\u0076\u0065\u0020a\u006e\u0064\u0020\u0073\u0065v\u0065\u006e")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020c\u006f\u0073\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_eegg := args[0].ValueNumber
	if _eegg < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044B \u0072\u0065\u0071u\u0069\u0072\u0065\u0073 co\u0073t \u0074\u006f\u0020\u0062\u0065\u0020\u006eon\u0020\u006e\u0065\u0067\u0061\u0074\u0069v\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0056\u0044\u0042 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_dbfeb := args[1].ValueNumber
	if _dbfeb < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0061\u006c\u0076\u0061\u0067\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020l\u0069\u0066\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_eedf := args[2].ValueNumber
	if _eedf == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if _eedf < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("V\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0073\u0074\u0061\u0072\u0074 p\u0065\u0072\u0069\u006fd\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cfgf := args[3].ValueNumber
	if _cfgf < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044\u0042\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074o\u0020\u0062\u0065\u0020\u006e\u006f\u0074\u0020\u006c\u0065\u0073\u0073\u0020\u0074h\u0061n\u0020\u006f\u006e\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("\u0056D\u0042\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0065\u006e\u0064 \u0070\u0065\u0072\u0069\u006f\u0064 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gffba := args[4].ValueNumber
	if _cfgf > _gffba {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020s\u0074\u0061\u0072\u0074\u0020\u0070\u0065r\u0069\u006f\u0064\u0020\u0066\u006f\u0072\u0020\u0056\u0044\u0042")
	}
	if _gffba > _eedf {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0065\u006e\u0064\u0020\u0070e\u0072i\u006f\u0064\u0020\u0066\u006f\u0072\u0020V\u0044\u0042")
	}
	_dfff := 2.0
	if _bcad > 5 {
		if args[5].Type == ResultTypeEmpty {
			_dfff = 0.0
		} else {
			if args[5].Type != ResultTypeNumber {
				return MakeErrorResult("\u0056\u0044\u0042\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0061\u0063\u0074\u006f\u0072 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
			}
			_dfff = args[5].ValueNumber
			if _dfff < 0 {
				return MakeErrorResultType(ErrorTypeNum, "\u0056\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u0074\u006f\u0020\u0062e\u0020\u006e\u006f\u006e\u0020n\u0065\u0067a\u0074\u0069\u0076\u0065")
			}
		}
	}
	_abgd := false
	if _bcad > 6 && args[6].Type != ResultTypeEmpty {
		if args[6].Type != ResultTypeNumber {
			return MakeErrorResult("\u0056D\u0042\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020n\u006f\u005f\u0073\u0077\u0069\u0074\u0063\u0068\u0020to\u0020\u0062\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
		}
		_abgd = args[6].ValueNumber != 0
	}
	_abab := 0.0
	_fafgb := _cc.Floor(_cfgf)
	_fdgg := _cc.Ceil(_gffba)
	if _abgd {
		for _gbbbd := _fafgb + 1; _gbbbd <= _fdgg; _gbbbd++ {
			_bfga := _gege(_eegg, _dbfeb, _eedf, _gbbbd, _dfff)
			if _gbbbd == _fafgb+1 {
				_bfga *= _cc.Min(_gffba, _fafgb+1) - _cfgf
			} else if _gbbbd == _fdgg {
				_bfga *= _gffba + 1 - _fdgg
			}
			_abab += _bfga
		}
	} else {
		_ebe := _eedf
		var _cbdcc float64
		if !_gacg(_cfgf, _cc.Floor(_cfgf)) {
			if _dfff == 1 {
				_ccgb := _eedf / 2
				if _cfgf > _ccgb || _gacg(_cfgf, _ccgb) {
					_cbdcc = _cfgf - _ccgb
					_cfgf = _ccgb
					_gffba -= _cbdcc
					_ebe++
				}
			}
		}
		if _dfff != 0 {
			_eegg -= _gcfa(_eegg, _dbfeb, _eedf, _ebe, _cfgf, _dfff)
		}
		_abab = _gcfa(_eegg, _dbfeb, _eedf, _eedf-_cfgf, _gffba-_cfgf, _dfff)
	}
	return MakeNumberResult(_abab)
}

// MakeErrorResult constructs a #VALUE! error with a given extra error message.
// The error message is for debugging formula evaluation only and is not stored
// in the sheet.
func MakeErrorResult(msg string) Result { return MakeErrorResultType(ErrorTypeValue, msg) }

// Value is an implementation of the Excel VALUE function.
func Value(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0056\u0041\u004c\u0055\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020a\u0020s\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gdbgg := args[0]
	if _gdbgg.Type == ResultTypeNumber {
		return _gdbgg
	}
	if _gdbgg.Type == ResultTypeString {
		_acbe, _bebc := _afc.ParseFloat(_gdbgg.Value(), 64)
		if _bebc == nil {
			return MakeNumberResult(_acbe)
		}
	}
	return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0056\u0041L\u0055\u0045")
}

// Month is an implementation of the Excel MONTH() function.
func Month(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("M\u004f\u004e\u0054\u0048\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006ee\u0020\u0061\u0072g\u0075m\u0065\u006e\u0074")
	}
	_bcae := args[0]
	switch _bcae.Type {
	case ResultTypeEmpty:
		return MakeNumberResult(1)
	case ResultTypeNumber:
		_babbg := _bgee(_bcae.ValueNumber)
		return MakeNumberResult(float64(_babbg.Month()))
	case ResultTypeString:
		_egfe := _ga.ToLower(_bcae.ValueString)
		if !_cddc(_egfe) {
			_, _, _, _, _fdd, _gge := _ccce(_egfe)
			if _gge.Type == ResultTypeError {
				_gge.ErrorMessage = "\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0072\u0020\u004dON\u0054\u0048"
				return _gge
			}
			if _fdd {
				return MakeNumberResult(1)
			}
		}
		_, _fdca, _, _, _gcec := _ddd(_egfe)
		if _gcec.Type == ResultTypeError {
			return _gcec
		}
		return MakeNumberResult(float64(_fdca))
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u004d\u004fN\u0054\u0048")
	}
}
func _cgce(_caac, _fgfg, _agcf, _bbf, _ceed, _ggbde float64, _fege int) (float64, Result) {
	_bcafg := int(_ggbde)
	_befd := _def(_caac, _fgfg, _bcafg, _fege)
	_fag := _ddg(_caac, _fgfg, _bcafg, _fege) / _befd
	_gade, _eeaga := _cefd(_caac, _fgfg, _bcafg, _fege)
	if _eeaga.Type == ResultTypeError {
		return 0, _eeaga
	}
	_bdca := _fbebf(_caac, _fgfg, _bcafg, _fege)
	_geef := _ceed / _cc.Pow(1+_bbf/_ggbde, _gade-1+_fag)
	_geef -= 100 * _agcf / _ggbde * _bdca / _befd
	_gdea := 100 * _agcf / _ggbde
	_gbfd := 1 + _bbf/_ggbde
	for _cdde := 0.0; _cdde < _gade; _cdde++ {
		_geef += _gdea / _cc.Pow(_gbfd, _cdde+_fag)
	}
	return _geef, MakeEmptyResult()
}

var _bcab = [...]int{-1, 1, 1, -1, -2, 0}

// IsLeapYear is an implementation of the Excel ISLEAPYEAR() function.
func IsLeapYear(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049S\u004c\u0045A\u0050\u0059\u0045\u0041R\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073in\u0067\u006c\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_fedg := ctx.GetEpoch()
	_bagb, _ggcac := _fgcbc(args[0].Value(), _fedg)
	if _ggcac != nil {
		return MakeErrorResult("\u0049S\u004c\u0045A\u0050\u0059\u0045\u0041R\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073in\u0067\u006c\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_ffbc := _bagb.Year()
	return MakeBoolResult(_bee(_ffbc))
}
func _adgg(_cbec, _efga, _bbgd, _gfd int) int {
	if !_gbgg(_gfd) {
		return _bbgd
	}
	_cbeg := _bbgd
	_bag := _ffcg(_cbec, _efga)
	if _cbeg > 30 || _bbgd >= _bag || _cbeg >= _bag {
		_cbeg = 30
	}
	return _cbeg
}

// ISERR is an implementation of the Excel ISERR() function.
func IsErr(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u0045\u0052\u0052\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeError && args[0].ValueString != "\u0023\u004e\u002f\u0041")
}
func _ddg(_eacf, _gbfc float64, _abf, _bda int) float64 {
	_fdad := _bgee(_eacf)
	_gaed := _bgee(_gbfc)
	_gfdb := _eage(_fdad, _gaed, _abf)
	return _cbea(_fdad, _gfdb, _bda)
}

var _fce = map[string]*_f.Regexp{}

// RoundUp is an implementation of the Excel ROUNDUP function that rounds a number
// up to a specified number of digits.
func RoundUp(args []Result) Result { return _afdcc(args, _gcfgg) }

// String returns a string representation of a vertical range.
func (_eagg VerticalRange) String() string { return _eagg.verticalRangeReference() }
func _aabcbf() yyParser                    { return &yyParserImpl{} }
func _ccaeb(_egfbd, _eaeec Reference) string {
	return _e.Sprintf("\u0025\u0073\u003a%\u0073", _egfbd.Value, _eaeec.Value)
}

var _gafa = false

const (
	ResultTypeUnknown ResultType = iota
	ResultTypeNumber
	ResultTypeString
	ResultTypeList
	ResultTypeArray
	ResultTypeError
	ResultTypeEmpty
)

// Reference returns an invalid reference for EmptyExpr.
func (_cdd EmptyExpr) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

const _cfc = "\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u003a\u0028\u0028\u005b\u0030-\u0039]\u0029\u002b\u0029\u0028\u0020\u0028\u0061\u006d\u007c\u0070\u006d\u0029\u0029\u003f"

// Cumprinc implements the Excel CUMPRINC function.
func Cumprinc(args []Result) Result {
	_fcef, _adcb := _aabcb(args, "\u0043\u0055\u004d\u0050\u0052\u0049\u004e\u0043")
	if _adcb.Type == ResultTypeError {
		return _adcb
	}
	_bfgf := _fcef._ffa
	_bgec := _fcef._dbcdb
	_bcda := _fcef._bdeb
	_ebcf := _fcef._gggb
	_dadc := _fcef._gffb
	_aacd := _fcef._cfgc
	_egec := _ecc(_bfgf, _bgec, _bcda, 0, _aacd)
	_bgbe := 0.0
	if _ebcf == 1 {
		if _aacd == 0 {
			_bgbe = _egec + _bcda*_bfgf
		} else {
			_bgbe = _egec
		}
		_ebcf++
	}
	for _babbe := _ebcf; _babbe <= _dadc; _babbe++ {
		if _aacd == 1 {
			_bgbe += _egec - (_ggbd(_bfgf, _babbe-2, _egec, _bcda, 1)-_egec)*_bfgf
		} else {
			_bgbe += _egec - _ggbd(_bfgf, _babbe-1, _egec, _bcda, 0)*_bfgf
		}
	}
	return MakeNumberResult(_bgbe)
}

type ivr struct{}

// Amorlinc implements the Excel AMORLINC function.
func Amorlinc(args []Result) Result {
	_ggac, _ecgc := _ffee(args, "\u0041\u004d\u004f\u0052\u004c\u0049\u004e\u0043")
	if _ecgc.Type == ResultTypeError {
		return _ecgc
	}
	_afe := _ggac._dffd
	_egdf := _ggac._dca
	_dgcb := _ggac._geg
	_egea := _ggac._abge
	_fgde := _ggac._dfaaf
	_eagc := _ggac._afed
	_cddcb := _ggac._cbdf
	_acbb, _ebdg := _fcfa(_egdf, _dgcb, _cddcb)
	if _ebdg.Type == ResultTypeError {
		return MakeErrorResult("\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0064\u0061\u0074\u0065\u0073 \u0066o\u0072\u0020\u0041\u004d\u004f\u0052\u004cI\u004e\u0043")
	}
	_ccg := _acbb * _eagc * _afe
	if _fgde == 0 {
		return MakeNumberResult(_ccg)
	}
	_edfb := _afe * _eagc
	_bff := _afe - _egea
	_abbb := int((_bff - _ccg) / _edfb)
	if _fgde <= _abbb {
		return MakeNumberResult(_edfb)
	} else if _fgde == _abbb+1 {
		return MakeNumberResult(_bff - _edfb*float64(_abbb) - _ccg)
	} else {
		return MakeNumberResult(0)
	}
}

// Reference returns an invalid reference for String.
func (_dcagbe String) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// Update returns the same object as updating sheet references does not affect ConstArrayExpr.
func (_cbg ConstArrayExpr) Update(q *_eg.UpdateQuery) Expression { return _cbg }

// Evaluator is the interface for a formula evaluator.  This is needed so we can
// pass it to the spreadsheet to let it evaluate formula cells before returning
// the results.
// NOTE: in order to implement Evaluator without cache embed noCache in it.
type Evaluator interface {
	Eval(_cee Context, formula string) Result
	SetCache(_agf string, _gcf Result)
	GetFromCache(_gfg string) (Result, bool)
	LastEvalIsRef() bool
}

func Trunc(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("T\u0052\u0055\u004e\u0043\u0028\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0061t \u006c\u0065\u0061\u0073t\u0020\u006f\u006e\u0065\u0020\u006e\u0075\u006d\u0065ri\u0063\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_cceae := args[0].AsNumber()
	if _cceae.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0054\u0052\u0055\u004e\u0043\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ddcgb := float64(0)
	if len(args) > 1 {
		_bbgad := args[1].AsNumber()
		if _bbgad.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020a\u0072\u0067\u0075m\u0065\u006e\u0074\u0020t\u006f\u0020\u0054\u0052\u0055\u004e\u0043\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_ddcgb = _bbgad.ValueNumber
	}
	_cggef := _cceae.ValueNumber
	_aade := 1.0
	if _ddcgb >= 0 {
		_aade = _cc.Pow(1/10.0, _ddcgb)
	} else {
		return MakeNumberResult(0)
	}
	_cggef, _bdaff := _cc.Modf(_cggef / _aade)
	_aaeab := 0.99999
	if _bdaff > _aaeab {
		_cggef++
	} else if _bdaff < -_aaeab {
		_cggef--
	}
	_ = _bdaff
	return MakeNumberResult(_cggef * _aade)
}

// Accrintm implements the Excel ACCRINTM function.
func Accrintm(args []Result) Result {
	_cfec := len(args)
	if _cfec != 4 && _cfec != 5 {
		return MakeErrorResult("A\u0043\u0043\u0052\u0049\u004e\u0054\u004d\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066o\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065 a\u0072\u0067\u0075m\u0065n\u0074\u0073")
	}
	_edaf, _bea := _eec(args[0], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u0041\u0043\u0043\u0052\u0049\u004e\u0054\u004d")
	if _bea.Type == ResultTypeError {
		return _bea
	}
	_dgba, _bea := _eec(args[1], "\u0073e\u0074t\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065", "\u0041\u0043\u0043\u0052\u0049\u004e\u0054\u004d")
	if _bea.Type == ResultTypeError {
		return _bea
	}
	if _edaf >= _dgba {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u0073\u0073\u0075\u0065\u0020d\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0065\u0061\u0072\u006c\u0069\u0065r\u0020\u0074\u0068\u0061\u006e\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065n\u0074 \u0064\u0061\u0074\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0041C\u0043\u0052I\u004e\u0054\u004d\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020to\u0020\u0062\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_efeg := args[2].ValueNumber
	if _efeg <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0041\u0043\u0043\u0052\u0049\u004e\u0054\u004d\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061t\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0041\u0043\u0043\u0052\u0049\u004e\u0054M\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s\u0020\u0070\u0061\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fcgb := args[3].ValueNumber
	if _fcgb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0041\u0043C\u0052\u0049\u004e\u0054\u004d \u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0070\u0061\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_egee := 0
	if _cfec == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0041C\u0043\u0052I\u004e\u0054\u004d \u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_egee = int(args[4].ValueNumber)
		if !_gfe(_egee) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0041\u0043\u0043R\u0049\u004e\u0054\u004d")
		}
	}
	_faae, _bea := _fcfa(_edaf, _dgba, _egee)
	if _bea.Type == ResultTypeError {
		return _bea
	}
	return MakeNumberResult(_fcgb * _efeg * _faae)
}

const _acec = 57365

func init() {
	_aeaf()
	RegisterFunction("\u0041V\u0045\u0052\u0041\u0047\u0045", Average)
	RegisterFunction("\u0041\u0056\u0045\u0052\u0041\u0047\u0045\u0041", Averagea)
	RegisterFunction("\u0043\u004f\u0055N\u0054", Count)
	RegisterFunction("\u0043\u004f\u0055\u004e\u0054\u0041", Counta)
	RegisterFunction("\u0043O\u0055\u004e\u0054\u0049\u0046", CountIf)
	RegisterFunction("\u0043\u004f\u0055\u004e\u0054\u0049\u0046\u0053", CountIfs)
	RegisterFunction("\u0043\u004f\u0055\u004e\u0054\u0042\u004c\u0041\u004e\u004b", CountBlank)
	RegisterFunction("\u004d\u0041\u0058", Max)
	RegisterFunction("\u004d\u0041\u0058\u0041", MaxA)
	RegisterFunction("\u004d\u0041\u0058\u0049\u0046\u0053", MaxIfs)
	RegisterFunction("\u005f\u0078\u006cf\u006e\u002e\u004d\u0041\u0058\u0049\u0046\u0053", MaxIfs)
	RegisterFunction("\u004d\u0045\u0044\u0049\u0041\u004e", Median)
	RegisterFunction("\u004d\u0049\u004e", Min)
	RegisterFunction("\u004d\u0049\u004e\u0041", MinA)
	RegisterFunction("\u004d\u0049\u004e\u0049\u0046\u0053", MinIfs)
	RegisterFunction("\u005f\u0078\u006cf\u006e\u002e\u004d\u0049\u004e\u0049\u0046\u0053", MinIfs)
}

const (
	_ byte = iota
	_bfcd
	_eggf
	_edfe
	_cdbac
	_dfef
)
const _agfb = -1000
const _abcd = "\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070\u0065U\u006e\u006b\u006e\u006f\u0077\u006e\u0052\u0065\u0073u\u006c\u0074\u0054y\u0070\u0065\u004e\u0075\u006d\u0062\u0065\u0072\u0052\u0065s\u0075\u006c\u0074\u0054\u0079\u0070\u0065\u0053\u0074\u0072\u0069\u006e\u0067\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070\u0065\u004c\u0069\u0073\u0074\u0052\u0065\u0073\u0075lt\u0054\u0079p\u0065\u0041r\u0072\u0061\u0079\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070\u0065\u0045\u0072\u0072\u006f\u0072\u0052\u0065\u0073\u0075\u006c\u0074\u0054\u0079\u0070\u0065\u0045\u006d\u0070\u0074\u0079"

func (_fcae ReferenceType) String() string {
	if _fcae >= ReferenceType(len(_gdgdf)-1) {
		return _e.Sprintf("\u0052\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0054\u0079\u0070e\u0028\u0025\u0064\u0029", _fcae)
	}
	return _egad[_gdgdf[_fcae]:_gdgdf[_fcae+1]]
}

// Eval evaluates and returns the result of the cell reference.
func (_cbd CellRef) Eval(ctx Context, ev Evaluator) Result { return ctx.Cell(_cbd._ab, ev) }

// MinA is an implementation of the Excel MINA() function.
func MinA(args []Result) Result { return _bdcee(args, true) }

// Reference returns an invalid reference for Error.
func (_fbd Error) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }

// GetWidth returns 0 for the invalid reference context.
func (_gbafg *ivr) GetWidth(colIdx int) float64 { return float64(0) }
func _aadag(_dfffg Result) [][]Result {
	switch _dfffg.Type {
	case ResultTypeArray:
		return _dfffg.ValueArray
	case ResultTypeList:
		return [][]Result{_dfffg.ValueList}
	default:
		return [][]Result{}
	}
}

// String is a string expression.
type String struct{ _bedb string }

// Eval evaluates the binary expression using the context given.
func (_eea BinaryExpr) Eval(ctx Context, ev Evaluator) Result {
	_ac := _eea._gb.Eval(ctx, ev)
	if _ac.Type == ResultTypeError {
		return _ac
	}
	_fgf := _eea._gba.Eval(ctx, ev)
	if _fgf.Type == ResultTypeError {
		return _fgf
	}
	if _ac.Type == _fgf.Type {
		if _ac.Type == ResultTypeArray {
			if !_gc(_ac.ValueArray, _fgf.ValueArray) {
				return MakeErrorResult("l\u0068\u0073\u002f\u0072\u0068\u0073 \u0073\u0068\u006f\u0075\u006c\u0064 \u0068\u0061\u0076\u0065\u0020\u0073\u0061m\u0065\u0020\u0064\u0069\u006d\u0065\u006e\u0073\u0069\u006fn\u0073")
			}
			return _bdg(_eea._fc, _ac.ValueArray, _fgf.ValueArray)
		} else if _ac.Type == ResultTypeList {
			if len(_ac.ValueList) != len(_fgf.ValueList) {
				return MakeErrorResult("l\u0068\u0073\u002f\u0072\u0068\u0073 \u0073\u0068\u006f\u0075\u006c\u0064 \u0068\u0061\u0076\u0065\u0020\u0073\u0061m\u0065\u0020\u0064\u0069\u006d\u0065\u006e\u0073\u0069\u006fn\u0073")
			}
			return _ebg(_eea._fc, _ac.ValueList, _fgf.ValueList)
		}
	} else if _ac.Type == ResultTypeArray && (_fgf.Type == ResultTypeNumber || _fgf.Type == ResultTypeString) {
		return _de(_eea._fc, _ac.ValueArray, _fgf)
	} else if _ac.Type == ResultTypeList && (_fgf.Type == ResultTypeNumber || _fgf.Type == ResultTypeString) {
		return _fgfe(_eea._fc, _ac.ValueList, _fgf)
	}
	switch _eea._fc {
	case BinOpTypePlus:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeNumberResult(_ac.ValueNumber + _fgf.ValueNumber)
			}
		}
	case BinOpTypeMinus:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeNumberResult(_ac.ValueNumber - _fgf.ValueNumber)
			}
		}
	case BinOpTypeMult:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeNumberResult(_ac.ValueNumber * _fgf.ValueNumber)
			}
		}
	case BinOpTypeDiv:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				if _fgf.ValueNumber == 0 {
					return MakeErrorResultType(ErrorTypeDivideByZero, "\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0062\u0079 \u007a\u0065\u0072\u006f")
				}
				return MakeNumberResult(_ac.ValueNumber / _fgf.ValueNumber)
			}
		}
	case BinOpTypeExp:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeNumberResult(_cc.Pow(_ac.ValueNumber, _fgf.ValueNumber))
			}
		}
	case BinOpTypeLT:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeBoolResult(_ac.ValueNumber < _fgf.ValueNumber)
			}
			if _ac.Type == ResultTypeString {
				return MakeBoolResult(_ac.ValueString < _fgf.ValueString)
			}
			if _ac.Type == ResultTypeEmpty {
				return MakeBoolResult(false)
			}
		} else if _ac.Type == ResultTypeString && _fgf.Type == ResultTypeNumber {
			return MakeBoolResult(false)
		} else if _ac.Type == ResultTypeNumber && _fgf.Type == ResultTypeString {
			return MakeBoolResult(true)
		} else if _ac.Type == ResultTypeEmpty && (_fgf.Type == ResultTypeNumber || _fgf.Type == ResultTypeString) {
			return MakeBoolResult(true)
		} else if (_ac.Type == ResultTypeNumber || _ac.Type == ResultTypeString) && _fgf.Type == ResultTypeEmpty {
			return MakeBoolResult(false)
		}
	case BinOpTypeGT:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeBoolResult(_ac.ValueNumber > _fgf.ValueNumber)
			}
			if _ac.Type == ResultTypeString {
				return MakeBoolResult(_ac.ValueString > _fgf.ValueString)
			}
			if _ac.Type == ResultTypeEmpty {
				return MakeBoolResult(false)
			}
		} else if _ac.Type == ResultTypeString && _fgf.Type == ResultTypeNumber {
			return MakeBoolResult(true)
		} else if _ac.Type == ResultTypeNumber && _fgf.Type == ResultTypeString {
			return MakeBoolResult(false)
		} else if _ac.Type == ResultTypeEmpty && (_fgf.Type == ResultTypeNumber || _fgf.Type == ResultTypeString) {
			return MakeBoolResult(false)
		} else if (_ac.Type == ResultTypeNumber || _ac.Type == ResultTypeString) && _fgf.Type == ResultTypeEmpty {
			return MakeBoolResult(true)
		}
	case BinOpTypeEQ:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeBoolResult(_ac.ValueNumber == _fgf.ValueNumber)
			}
			if _ac.Type == ResultTypeString {
				return MakeBoolResult(_ac.ValueString == _fgf.ValueString)
			}
			if _ac.Type == ResultTypeEmpty {
				return MakeBoolResult(true)
			}
		} else if (_ac.Type == ResultTypeString && _fgf.Type == ResultTypeNumber) || (_ac.Type == ResultTypeNumber && _fgf.Type == ResultTypeString) {
			return MakeBoolResult(false)
		} else if _ac.Type == ResultTypeEmpty && (_fgf.Type == ResultTypeNumber || _fgf.Type == ResultTypeString) {
			return MakeBoolResult(_ag(_fgf))
		} else if (_ac.Type == ResultTypeNumber || _ac.Type == ResultTypeString) && _fgf.Type == ResultTypeEmpty {
			return MakeBoolResult(_ag(_ac))
		}
	case BinOpTypeNE:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeBoolResult(_ac.ValueNumber != _fgf.ValueNumber)
			}
			if _ac.Type == ResultTypeString {
				return MakeBoolResult(_ac.ValueString != _fgf.ValueString)
			}
			if _ac.Type == ResultTypeEmpty {
				return MakeBoolResult(false)
			}
		} else if (_ac.Type == ResultTypeString && _fgf.Type == ResultTypeNumber) || (_ac.Type == ResultTypeNumber && _fgf.Type == ResultTypeString) {
			return MakeBoolResult(true)
		} else if _ac.Type == ResultTypeEmpty && (_fgf.Type == ResultTypeNumber || _fgf.Type == ResultTypeString) {
			return MakeBoolResult(!_ag(_fgf))
		} else if (_ac.Type == ResultTypeNumber || _ac.Type == ResultTypeString) && _fgf.Type == ResultTypeEmpty {
			return MakeBoolResult(!_ag(_ac))
		}
	case BinOpTypeLEQ:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeBoolResult(_ac.ValueNumber <= _fgf.ValueNumber)
			}
			if _ac.Type == ResultTypeString {
				return MakeBoolResult(_ac.ValueString <= _fgf.ValueString)
			}
			if _ac.Type == ResultTypeEmpty {
				return MakeBoolResult(true)
			}
		} else if _ac.Type == ResultTypeString && _fgf.Type == ResultTypeNumber {
			return MakeBoolResult(false)
		} else if _ac.Type == ResultTypeNumber && _fgf.Type == ResultTypeString {
			return MakeBoolResult(true)
		} else if _ac.Type == ResultTypeEmpty && (_fgf.Type == ResultTypeNumber || _fgf.Type == ResultTypeString) {
			return MakeBoolResult(_ag(_fgf))
		} else if (_ac.Type == ResultTypeNumber || _ac.Type == ResultTypeString) && _fgf.Type == ResultTypeEmpty {
			return MakeBoolResult(_ag(_ac))
		}
	case BinOpTypeGEQ:
		if _ac.Type == _fgf.Type {
			if _ac.Type == ResultTypeNumber {
				return MakeBoolResult(_ac.ValueNumber >= _fgf.ValueNumber)
			}
			if _ac.Type == ResultTypeString {
				return MakeBoolResult(_ac.ValueString >= _fgf.ValueString)
			}
			if _ac.Type == ResultTypeEmpty {
				return MakeBoolResult(true)
			}
		} else if _ac.Type == ResultTypeString && _fgf.Type == ResultTypeNumber {
			return MakeBoolResult(true)
		} else if _ac.Type == ResultTypeNumber && _fgf.Type == ResultTypeString {
			return MakeBoolResult(false)
		} else if _ac.Type == ResultTypeEmpty && (_fgf.Type == ResultTypeNumber || _fgf.Type == ResultTypeString) {
			return MakeBoolResult(_ag(_fgf))
		} else if (_ac.Type == ResultTypeNumber || _ac.Type == ResultTypeString) && _fgf.Type == ResultTypeEmpty {
			return MakeBoolResult(_ag(_ac))
		}
	case BinOpTypeConcat:
		return MakeStringResult(_ac.Value() + _fgf.Value())
	}
	return MakeErrorResult("u\u006e\u0073\u0075\u0070po\u0072t\u0065\u0064\u0020\u0062\u0069n\u0061\u0072\u0079\u0020\u006f\u0070")
}

// Rate implements the Excel RATE function.
func Rate(args []Result) Result {
	_ggca := len(args)
	if _ggca < 3 || _ggca > 6 {
		return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0072\u0065\u0065 \u0061\u006e\u0064\u0020\u0073i\u0078")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ccdc := args[0].ValueNumber
	if _ccdc != float64(int(_ccdc)) {
		return MakeErrorResultType(ErrorTypeNum, "R\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072i\u006fd\u0073\u0020\u0074\u006f \u0062\u0065 \u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072\u0065q\u0075\u0069\u0072e\u0073\u0020\u0070\u0061y\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gab := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_ddfdg := args[2].ValueNumber
	_dbgb := 0.0
	if _ggca >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s \u0066\u0075\u0074\u0075\u0072\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		}
		_dbgb = args[3].ValueNumber
	}
	_efef := 0.0
	if _ggca >= 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("R\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_efef = args[4].ValueNumber
		if _efef != 0 {
			_efef = 1
		}
	}
	_bddgg := 0.1
	if _ggca >= 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("\u0052\u0041\u0054\u0045\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0067\u0075\u0065\u0073\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_bddgg = args[5].ValueNumber
	}
	_gcbe := 100
	_effe := 0
	_agcg := false
	_eega := 1e-6
	_bgecd := _bddgg
	for _effe < _gcbe && !_agcg {
		_cebd := _cc.Pow(_bgecd+1, _ccdc)
		_bdcf := _cc.Pow(_bgecd+1, _ccdc-1)
		_cbee := _bgecd*_efef + 1
		_age := _gab * (_cebd - 1)
		_bbfg := _dbgb + _cebd*_ddfdg + _age*_cbee/_bgecd
		_facb := _ccdc*_bdcf*_ddfdg - _age*_cbee/_cc.Pow(_bgecd, 2)
		_fadbc := (_ccdc*_gab*_bdcf*_cbee + _age*_efef) / _bgecd
		_badbc := _bbfg / (_facb + _fadbc)
		if _cc.Abs(_badbc) < _eega {
			_agcg = true
		}
		_effe++
		_bgecd -= _badbc
	}
	return MakeNumberResult(_bgecd)
}

const _cfdg = 57357

func _bgdb(_cbbf []Result) (float64, float64, Result) {
	_ccdg := 0.0
	_fegb := 1.0
	for _, _geecg := range _cbbf {
		switch _geecg.Type {
		case ResultTypeNumber:
			_ccdg += _geecg.ValueNumber
			_fegb *= _cfga(_geecg.ValueNumber)
		case ResultTypeList, ResultTypeArray:
			_gefa, _fbca, _eecf := _bgdb(_geecg.ListValues())
			_ccdg += _gefa
			_fegb *= _cfga(_fbca)
			if _eecf.Type == ResultTypeError {
				return 0, 0, _eecf
			}
		case ResultTypeString:
			return 0, 0, MakeErrorResult("M\u0055\u004c\u0054\u0049\u004e\u004f\u004d\u0049\u0041\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063 a\u0072\u0067\u0075m\u0065n\u0074\u0073")
		case ResultTypeError:
			return 0, 0, _geecg
		}
	}
	return _ccdg, _fegb, _feb
}

// Power is an implementation of the Excel POWER function that raises a number
// to a power. It requires two numeric arguments.
func Power(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0050\u004f\u0057\u0045\u0052\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_bfece := args[0].AsNumber()
	if _bfece.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0050\u004f\u0057\u0045\u0052\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_cgef := args[1].AsNumber()
	if _cgef.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020a\u0072\u0067\u0075m\u0065\u006e\u0074\u0020t\u006f\u0020\u0050\u004f\u0057\u0045\u0052\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	return MakeNumberResult(_cc.Pow(_bfece.ValueNumber, _cgef.ValueNumber))
}

// ISEVEN is an implementation of the Excel ISEVEN() function.
func IsEven(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u0045VE\u004e\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070t\u0073 \u0061 \u0073i\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u0053\u0045\u0056\u0045\u004e \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_cfde := int(args[0].ValueNumber)
	return MakeBoolResult(_cfde == _cfde/2*2)
}
func _bdg(_ed BinOpType, _ge, _ecf [][]Result) Result {
	_bdd := [][]Result{}
	for _edb := range _ge {
		_cac := _ebg(_ed, _ge[_edb], _ecf[_edb])
		if _cac.Type == ResultTypeError {
			return _cac
		}
		_bdd = append(_bdd, _cac.ValueList)
	}
	return MakeArrayResult(_bdd)
}

var _feb Result = MakeEmptyResult()

// Update updates the horizontal range references after removing a row/column.
func (_bece HorizontalRange) Update(q *_eg.UpdateQuery) Expression { return _bece }

var _aaebee = map[string]FunctionComplex{}

// Reference returns an invalid reference for Number.
func (_eabae Number) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }
func (_gcgc *plex) Lex(lval *yySymType) int {
	_bdff = true
	_adgeb := <-_gcgc._cegce
	if _adgeb != nil {
		lval._ecdgg = _adgeb
		return int(lval._ecdgg._dggg)
	}
	return 0
}

// Or is an implementation of the Excel OR() function and takes a variable
// number of arguments.
func Or(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u004f\u0052\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 \u006f\u006e\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_eeaa := false
	for _, _geff := range args {
		switch _geff.Type {
		case ResultTypeList, ResultTypeArray:
			_fegg := Or(_geff.ListValues())
			if _fegg.Type == ResultTypeError {
				return _fegg
			}
			if _fegg.ValueNumber != 0 {
				_eeaa = true
			}
		case ResultTypeNumber:
			if _geff.ValueNumber != 0 {
				_eeaa = true
			}
		case ResultTypeString:
			return MakeErrorResult("\u004f\u0052 \u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u006f\u0070\u0065\u0072\u0061\u0074\u0065\u0020\u006f\u006e\u0020\u0073\u0074\u0072in\u0067\u0073")
		case ResultTypeError:
			return _geff
		default:
			return MakeErrorResult("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u004f\u0052")
		}
	}
	return MakeBoolResult(_eeaa)
}

// Update updates references in the BinaryExpr after removing a row/column.
func (_gbc BinaryExpr) Update(q *_eg.UpdateQuery) Expression {
	_eef := _gbc
	_eef._gb = _gbc._gb.Update(q)
	_eef._gba = _gbc._gba.Update(q)
	return _eef
}

// NewPrefixExpr constructs an expression with prefix.
func NewPrefixExpr(pfx, exp Expression) Expression { return &PrefixExpr{_dfgab: pfx, _bcada: exp} }
func _gacg(_bga, _eaa float64) bool                { return _cc.Abs(_bga-_eaa) < 1.0e-6 }

// Update returns the same object as updating sheet references does not affect Error.
func (_cgd Error) Update(q *_eg.UpdateQuery) Expression { return _cgd }

// Eval evaluates and returns the result of a formula.
func (_bfg *defEval) Eval(ctx Context, formula string) Result {
	_eed := ParseString(formula)
	_fca := make(chan Result)
	go func() {
		if _eed == nil {
			_fca <- MakeErrorResult(_e.Sprintf("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070a\u0072\u0073\u0065\u0020\u0066\u006f\u0072\u006d\u0075\u006ca\u0020\u0025\u0073", formula))
		} else {
			_bfg.checkLastEvalIsRef(ctx, _eed)
			_fca <- _eed.Eval(ctx, _bfg)
		}
	}()
	select {
	case _dcd := <-_fca:
		return _dcd
	case <-_eb.After(_ffdf):
		_cf.Log.Debug("\u0055\u006e\u0069\u004ff\u0066\u0069\u0063\u0065\u0020\u0065\u0076\u0061\u006c\u0075a\u0074i\u006f\u006e\u0020\u0074\u0069\u006d\u0065o\u0075\u0074")
		return MakeNumberResult(0)
	}
}

var _gegbg = map[string]bool{"\u0049F\u0045\u0052\u0052\u004f\u0052": true, "\u0049\u0046\u004e\u0041": true, "\u005f\u0078\u006c\u0066\u006e\u002e\u0049\u0046\u004e\u0041": true, "\u0049\u0053\u0045R\u0052": true, "\u0049S\u0045\u0052\u0052\u004f\u0052": true, "\u0049\u0053\u004e\u0041": true, "\u0049\u0053\u0052E\u0046": true}
var _agde _af.Mutex
var _afd = []*_f.Regexp{}

func _bgeeg(_gagc []Result) Result {
	_gaeb := _gagc[0].ValueArray
	if len(_gagc) == 1 {
		_efff := [][]Result{}
		for _, _ggacg := range _gaeb {
			_efff = append(_efff, _faca([]Result{MakeListResult(_ggacg)}).ValueList)
		}
		return MakeArrayResult(_efff)
	} else if len(_gagc) == 2 {
		_gdeab := len(_gaeb)
		_dfddg := len(_gaeb[0])
		_ggcgd := _efce(_gagc[1], _gdeab, _dfddg)
		_edbc := len(_ggcgd)
		_gcfb := [][]Result{}
		var _badgdc []Result
		for _gefc, _cedg := range _gaeb {
			if _gefc < _edbc {
				_badgdc = _ggcgd[_gefc]
			} else {
				_badgdc = _gefg(MakeErrorResultType(ErrorTypeNA, ""), _dfddg)
			}
			_gcfb = append(_gcfb, _faca([]Result{MakeListResult(_cedg), MakeListResult(_badgdc)}).ValueList)
		}
		return MakeArrayResult(_gcfb)
	} else if len(_gagc) == 3 {
		_gfdbe := len(_gaeb)
		_fbcb := len(_gaeb[0])
		_gaac := _efce(_gagc[1], _gfdbe, _fbcb)
		_eeef := _efce(_gagc[2], _gfdbe, _fbcb)
		_cgbd := len(_gaac)
		_gfbcf := len(_eeef)
		_daef := [][]Result{}
		var _cbgc, _dcab []Result
		for _cbae, _gfca := range _gaeb {
			if _cbae < _cgbd {
				_cbgc = _gaac[_cbae]
			} else {
				_cbgc = _gefg(MakeErrorResultType(ErrorTypeNA, ""), _fbcb)
			}
			if _cbae < _gfbcf {
				_dcab = _eeef[_cbae]
			} else {
				_dcab = _gefg(MakeErrorResultType(ErrorTypeNA, ""), _fbcb)
			}
			_daef = append(_daef, _faca([]Result{MakeListResult(_gfca), MakeListResult(_cbgc), MakeListResult(_dcab)}).ValueList)
		}
		return MakeArrayResult(_daef)
	}
	return MakeErrorResultType(ErrorTypeValue, "")
}

const (
	BinOpTypeUnknown BinOpType = iota
	BinOpTypePlus
	BinOpTypeMinus
	BinOpTypeMult
	BinOpTypeDiv
	BinOpTypeExp
	BinOpTypeLT
	BinOpTypeGT
	BinOpTypeEQ
	BinOpTypeLEQ
	BinOpTypeGEQ
	BinOpTypeNE
	BinOpTypeConcat
)

func (_efadg node) String() string {
	return _e.Sprintf("\u007b%\u0073\u0020\u0025\u0073\u007d", _efadg._dggg, _cabg(string(_efadg._eedag)))
}

// NamedRangeRef is a reference to a named range.
type NamedRangeRef struct{ _ceaf string }

func init() {
	RegisterFunction("\u0043\u0048\u0041\u0052", Char)
	RegisterFunction("\u0043\u004c\u0045A\u004e", Clean)
	RegisterFunction("\u0043\u004f\u0044\u0045", Code)
	RegisterFunction("C\u004f\u004e\u0043\u0041\u0054\u0045\u004e\u0041\u0054\u0045", Concat)
	RegisterFunction("\u0043\u004f\u004e\u0043\u0041\u0054", Concat)
	RegisterFunction("\u005f\u0078\u006cf\u006e\u002e\u0043\u004f\u004e\u0043\u0041\u0054", Concat)
	RegisterFunction("\u0045\u0058\u0041C\u0054", Exact)
	RegisterFunction("\u0046\u0049\u004e\u0044", Find)
	RegisterFunctionComplex("\u0046\u0049\u004eD\u0042", Findb)
	RegisterFunction("\u004c\u0045\u0046\u0054", Left)
	RegisterFunction("\u004c\u0045\u0046T\u0042", Left)
	RegisterFunction("\u004c\u0045\u004e", Len)
	RegisterFunction("\u004c\u0045\u004e\u0042", Len)
	RegisterFunction("\u004c\u004f\u0057E\u0052", Lower)
	RegisterFunction("\u004d\u0049\u0044", Mid)
	RegisterFunction("\u0050\u0052\u004f\u0050\u0045\u0052", Proper)
	RegisterFunction("\u0052E\u0050\u004c\u0041\u0043\u0045", Replace)
	RegisterFunction("\u0052\u0045\u0050\u0054", Rept)
	RegisterFunction("\u0052\u0049\u0047H\u0054", Right)
	RegisterFunction("\u0052\u0049\u0047\u0048\u0054\u0042", Right)
	RegisterFunction("\u0053\u0045\u0041\u0052\u0043\u0048", Search)
	RegisterFunctionComplex("\u0053E\u0041\u0052\u0043\u0048\u0042", Searchb)
	RegisterFunction("\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", Substitute)
	RegisterFunction("\u0054", T)
	RegisterFunction("\u0054\u0045\u0058\u0054", Text)
	RegisterFunction("\u0054\u0045\u0058\u0054\u004a\u004f\u0049\u004e", TextJoin)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0054\u0045\u0058T\u004a\u004f\u0049\u004e", TextJoin)
	RegisterFunction("\u0054\u0052\u0049\u004d", Trim)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0055\u004e\u0049\u0043\u0048\u0041\u0052", Char)
	RegisterFunction("\u005f\u0078\u006c\u0066\u006e\u002e\u0055\u004e\u0049\u0043\u004f\u0044\u0045", Unicode)
	RegisterFunction("\u0055\u0050\u0050E\u0052", Upper)
	RegisterFunction("\u0056\u0041\u004cU\u0045", Value)
}

const _cbfa int = 0

// Char is an implementation of the Excel CHAR function that takes an integer in
// the range [0,255] and returns the corresponding ASCII character.
func Char(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0043\u0048\u0041\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0061\u0020\u0073\u0069\u006e\u0067l\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_abfb := args[0].AsNumber()
	if _abfb.Type != ResultTypeNumber {
		return MakeErrorResult("\u0043\u0048\u0041\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0061\u0020\u0073\u0069\u006e\u0067l\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_dcdce := int(_abfb.ValueNumber)
	if _dcdce < 0 || _dcdce > 255 {
		return MakeErrorResult("\u0043H\u0041\u0052 \u0072\u0065\u0071\u0075i\u0072\u0065\u0073 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073 i\u006e\u0020\u0074h\u0065\u0020r\u0061\u006e\u0067\u0065\u0020\u005b0\u002c\u00325\u0035\u005d")
	}
	return MakeStringResult(_e.Sprintf("\u0025\u0063", _dcdce))
}

// NewSheetPrefixExpr constructs a new prefix expression.
func NewSheetPrefixExpr(s string) Expression { return &SheetPrefixExpr{_bgfda: s} }

type node struct {
	_dggg  tokenType
	_eedag string
}

func _bad(_bbe, _cebe, _ccf int) float64 {
	return float64(_gadd(_bbe, _eb.Month(_cebe), _ccf)/86400) + _acbd
}

var _fdcdf = [...]int{0, 1, 1, 2, 4, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 3, 1, 3, 1, 3, 1, 3, 1, 2, 1, 1, 1, 3, 4, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 3, 1, 3, 1, 1, 0}

const _bbgc = 86400000000000
const _gecg = 57369

// Pv implements the Excel PV function.
func Pv(args []Result) Result {
	_bfee := len(args)
	if _bfee < 3 || _bfee > 5 {
		return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0033\u0020\u0061\u006e\u0064\u00205")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_dgdc := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020o\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_ffcb := args[1].ValueNumber
	if _ffcb != float64(int(_ffcb)) {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006ff\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0061\u0079\u006d\u0065\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bcgg := args[2].ValueNumber
	_gaa := 0.0
	if _bfee >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0056 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0075\u0074\u0075\u0072\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_gaa = args[3].ValueNumber
	}
	_fddf := 0.0
	if _bfee == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_fddf = args[4].ValueNumber
		if _fddf != 0 {
			_fddf = 1
		}
	}
	if _dgdc == 0 {
		return MakeNumberResult(-_bcgg*_ffcb - _gaa)
	} else {
		return MakeNumberResult((((1-_cc.Pow(1+_dgdc, _ffcb))/_dgdc)*_bcgg*(1+_dgdc*_fddf) - _gaa) / _cc.Pow(1+_dgdc, _ffcb))
	}
}

// IfError is an implementation of the Excel IFERROR() function. It takes two arguments.
func IfError(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0049\u0046\u0045\u0052\u0052\u004f\u0052\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeError {
		if args[0].Type == ResultTypeEmpty {
			return MakeNumberResult(0)
		}
		return args[0]
	}
	return args[1]
}

// Floor is an implementation of the FlOOR function.
func Floor(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0046\u004c\u004f\u004f\u0052\u0028\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_gdge := args[0].AsNumber()
	if _gdge.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072s\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0046\u004c\u004f\u004f\u0052\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	var _ddge float64
	_gbeg := args[1].AsNumber()
	if _gbeg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020a\u0072\u0067\u0075m\u0065\u006e\u0074\u0020t\u006f\u0020\u0046\u004c\u004f\u004f\u0052\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ddge = _gbeg.ValueNumber
	if _ddge < 0 && _gdge.ValueNumber >= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0046L\u004f\u004f\u0052")
	}
	_bbce := _gdge.ValueNumber
	_bbce, _fcaca := _cc.Modf(_bbce / _ddge)
	if _fcaca != 0 {
		if _gdge.ValueNumber < 0 && _fcaca < 0 {
			_bbce--
		}
	}
	return MakeNumberResult(_bbce * _ddge)
}

// CountBlank implements the COUNTBLANK function.
func CountBlank(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0043\u004f\u0055N\u0054\u0042\u004c\u0041N\u004b\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u006e\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeNumberResult(_aabce(args, _cdgce))
}

// Intrate implements the Excel INTRATE function.
func Intrate(args []Result) Result {
	_gcea := len(args)
	if _gcea != 4 && _gcea != 5 {
		return MakeErrorResult("\u0049\u004e\u0054\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0066\u006f\u0075r\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_caff, _edfbc, _aabf := _gfcg(args[0], args[1], "\u0049N\u0054\u0052\u0041\u0054\u0045")
	if _aabf.Type == ResultTypeError {
		return _aabf
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u004e\u0054\u0052\u0041\u0054E\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0069\u006e\u0076\u0065\u0073\u0074\u006d\u0065\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_deeb := args[2].ValueNumber
	if _deeb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u004e\u0054\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0069\u006e\u0076e\u0073\u0074\u006d\u0065\u006e\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u0061r\u0067\u0075\u006de\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u004e\u0054\u0052\u0041\u0054E\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_acea := args[3].ValueNumber
	if _acea <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u004e\u0054\u0052\u0041\u0054\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064e\u006d\u0070\u0074\u0069\u006f\u006e\u0020\u0074\u006f \u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u0061r\u0067\u0075\u006de\u006e\u0074")
	}
	_adbf := 0
	if _gcea == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0049N\u0054\u0052A\u0054\u0045\u0020\u0072e\u0071\u0075\u0069r\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073\u0020to\u0020\u0062\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
		}
		_adbf = int(args[4].ValueNumber)
		if !_gfe(_adbf) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006eco\u0072\u0072\u0065c\u0074\u0020\u0062\u0061sis\u0020ar\u0067\u0075\u006d\u0065\u006e\u0074\u0020fo\u0072\u0020\u0049\u004e\u0054\u0052\u0041T\u0045")
		}
	}
	_fgbd, _aabf := _fcfa(_caff, _edfbc, _adbf)
	if _aabf.Type == ResultTypeError {
		return _aabf
	}
	return MakeNumberResult((_acea - _deeb) / _deeb / _fgbd)
}

// Ddb implements the Excel DDB function.
func Ddb(args []Result) Result {
	_efeb := len(args)
	if _efeb != 4 && _efeb != 5 {
		return MakeErrorResult("\u0044\u0044\u0042 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020c\u006f\u0073\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_adfg := args[0].ValueNumber
	if _adfg < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044B \u0072\u0065\u0071u\u0069\u0072\u0065\u0073 co\u0073t \u0074\u006f\u0020\u0062\u0065\u0020\u006eon\u0020\u006e\u0065\u0067\u0061\u0074\u0069v\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0044\u0042 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cbff := args[1].ValueNumber
	if _cbff < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0061\u006c\u0076\u0061\u0067\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020l\u0069\u0066\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_cdbd := args[2].ValueNumber
	if _cdbd <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0044\u0042\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ecaf := args[3].ValueNumber
	if _ecaf < 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044\u0042\u0020\u0072\u0065\u0071u\u0069\u0072\u0065s\u0020\u0070\u0065\u0072i\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u0074\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065")
	}
	if _ecaf > _cdbd {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0070\u0065\u0072i\u006f\u0064\u0020\u0066\u006f\u0072\u0020\u0044\u0044\u0042")
	}
	_bddda := 2.0
	if _efeb == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0044\u0044\u0042\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0061\u0063\u0074\u006f\u0072 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_bddda = args[4].ValueNumber
		if _bddda < 0 {
			return MakeErrorResultType(ErrorTypeNum, "\u0044\u0044\u0042\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u0074\u006f\u0020\u0062e\u0020\u006e\u006f\u006e\u0020n\u0065\u0067a\u0074\u0069\u0076\u0065")
		}
	}
	return MakeNumberResult(_gege(_adfg, _cbff, _cdbd, _ecaf, _bddda))
}

// Pduration implements the Excel PDURATION function.
func Pduration(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0050\u0044\u0055RA\u0054\u0049\u004f\u004e\u0020\u0072\u0065\u0071\u0075i\u0072e\u0073 \u0074h\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050D\u0055\u0052A\u0054\u0049\u004fN\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0072\u0061\u0074\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_adbb := args[0].ValueNumber
	if _adbb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020p\u006f\u0073i\u0074\u0069\u0076\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0063\u0075\u0072\u0072\u0065\u006e\u0074\u0020\u0076\u0061l\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006dbe\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ecgd := args[1].ValueNumber
	if _ecgd <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "P\u0044\u0055\u0052\u0041\u0054\u0049\u004f\u004e\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 c\u0075\u0072\u0072\u0065n\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074o \u0062\u0065 \u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0044\u0055\u0052\u0041\u0054I\u004f\u004e\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0073\u0070\u0065\u0063\u0069\u0066i\u0065\u0064\u0020\u0076\u0061lu\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cgbfb := args[2].ValueNumber
	if _cgbfb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0044\u0055\u0052\u0041\u0054I\u004f\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0065d\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070o\u0073i\u0074\u0069\u0076\u0065")
	}
	return MakeNumberResult((_cc.Log10(_cgbfb) - _cc.Log10(_ecgd)) / _cc.Log10(1+_adbb))
}

const _dgbf = 57371

// MakeListResult constructs a list result.
func MakeListResult(list []Result) Result { return Result{Type: ResultTypeList, ValueList: list} }

// False is an implementation of the Excel FALSE() function. It takes no
// arguments.
func False(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("\u0046A\u004c\u0053\u0045\u0020\u0074\u0061\u006b\u0065\u0073\u0020\u006eo\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	return MakeBoolResult(false)
}

const _eadb = "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0054\u0049\u004d\u0045\u0056\u0041\u004c\u0055\u0045"
const _adb = "\u005e\u0028\u0028" + _acde + "\u007c" + _bed + "\u007c" + _gbg + "\u007c" + _faef + "\u0029\u0020\u0029\u003f"

// Received implements the Excel RECEIVED function.
func Received(args []Result) Result {
	_bcb := len(args)
	if _bcb != 4 && _bcb != 5 {
		return MakeErrorResult("R\u0045\u0043\u0045\u0049\u0056\u0045\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066o\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065 a\u0072\u0067\u0075m\u0065n\u0074\u0073")
	}
	_gedg, _bdceaa, _fafg := _gfcg(args[0], args[1], "\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044")
	if _fafg.Type == ResultTypeError {
		return _fafg
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020i\u006e\u0076\u0065\u0073\u0074\u006d\u0065n\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_deed := args[2].ValueNumber
	if _deed <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0069\u006ev\u0065\u0073\u0074\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0045\u0043\u0045\u0049\u0056\u0045\u0044 \u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u0064\u0069s\u0063\u006f\u0075\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_gcaa := args[3].ValueNumber
	if _gcaa <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052\u0045\u0043\u0045I\u0056\u0045\u0044\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0064\u0069\u0073\u0063\u006f\u0075\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020p\u006f\u0073\u0069\u0074\u0069v\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_adaa := 0
	if _bcb == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0052E\u0043\u0045I\u0056\u0045\u0044 \u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_adaa = int(args[4].ValueNumber)
		if !_gfe(_adaa) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0052\u0045\u0043E\u0049\u0056\u0045\u0044")
		}
	}
	_ccfdc, _fafg := _fcfa(_gedg, _bdceaa, _adaa)
	if _fafg.Type == ResultTypeError {
		return _fafg
	}
	return MakeNumberResult(_deed / (1 - _gcaa*_ccfdc))
}

// Update updates references in the PrefixExpr after removing a row/column.
func (_decbd PrefixExpr) Update(q *_eg.UpdateQuery) Expression {
	_aaebg := _decbd
	_fbcge := _decbd._dfgab.String()
	if _fbcge == q.SheetToUpdate {
		_gbdgc := *q
		_gbdgc.UpdateCurrentSheet = true
		_aaebg._bcada = _decbd._bcada.Update(&_gbdgc)
	}
	return _aaebg
}
func (_cad *defEval) checkLastEvalIsRef(_dbb Context, _dcg Expression) {
	switch _dcg.(type) {
	case FunctionCall:
		switch _dcg.(FunctionCall)._eecfd {
		case "\u0049\u0053\u0052E\u0046":
			for _, _cada := range _dcg.(FunctionCall)._dgce {
				switch _cada.(type) {
				case CellRef, Range, HorizontalRange, VerticalRange, NamedRangeRef, PrefixExpr, PrefixRangeExpr, PrefixHorizontalRange, PrefixVerticalRange:
					_dgff := _cada.Eval(_dbb, _cad)
					_cad._fdc = !(_dgff.Type == ResultTypeError && _dgff.ValueString == "\u0023\u004e\u0041\u004d\u0045\u003f")
				default:
					_cad._fdc = false
				}
			}
		}
	}
}
func _cebf(_agfa []Result, _cbde string) (*couponArgs, Result) {
	_aabb := len(_agfa)
	if _aabb != 3 && _aabb != 4 {
		return nil, MakeErrorResult(_cbde + "\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u006f\u0072\u0020\u0066o\u0075\u0072\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_cccd, _adgeg, _gdbb := _gfcg(_agfa[0], _agfa[1], _cbde)
	if _gdbb.Type == ResultTypeError {
		return nil, _gdbb
	}
	if _agfa[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_cbde + "\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0066\u0072\u0065\u0071\u0075\u0065\u006e\u0063\u0079 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cddgd := _agfa[2].ValueNumber
	if !_dddg(_cddgd) {
		return nil, MakeErrorResult("\u0049n\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0066\u0072\u0065q\u0075\u0065\u006e\u0063\u0079\u0020\u0066\u006f\u0072\u0020" + _cbde)
	}
	_dgcc := 0
	if _aabb == 4 && _agfa[3].Type != ResultTypeEmpty {
		if _agfa[3].Type != ResultTypeNumber {
			return nil, MakeErrorResult(_cbde + "\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020b\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020b\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_dgcc = int(_agfa[3].ValueNumber)
		if !_gfe(_dgcc) {
			return nil, MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020fo\u0072\u0020"+_cbde)
		}
	}
	return &couponArgs{_cccd, _adgeg, int(_cddgd), _dgcc}, _feb
}

// And is an implementation of the Excel AND() function.
func And(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0041\u004e\u0044 r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061t\u0020l\u0065a\u0073t\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gaec := true
	for _, _cbca := range args {
		_cbca = _cbca.AsNumber()
		switch _cbca.Type {
		case ResultTypeList, ResultTypeArray:
			_caag := And(_cbca.ListValues())
			if _caag.Type == ResultTypeError {
				return _caag
			}
			if _caag.ValueNumber == 0 {
				_gaec = false
			}
		case ResultTypeNumber:
			if _cbca.ValueNumber == 0 {
				_gaec = false
			}
		case ResultTypeString:
			return MakeErrorResult("\u0041\u004e\u0044\u0020\u0064\u006f\u0065\u0073\u006e\u0027t\u0020\u006f\u0070\u0065\u0072\u0061\u0074e\u0020\u006f\u006e\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0073")
		case ResultTypeError:
			return _cbca
		default:
			return MakeErrorResult("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0061\u0072\u0067u\u006de\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u0041\u004e\u0044")
		}
	}
	return MakeBoolResult(_gaec)
}
func _addg(_gaad Result, _agfff *criteriaParsed) bool {
	if _gaad.IsBoolean {
		return false
	}
	_dcec := _gaad.Type
	if _agfff._ceeeb {
		return _dcec == ResultTypeNumber && _gaad.ValueNumber == _agfff._adag
	} else if _dcec == ResultTypeNumber {
		return _fgfcf(_gaad.ValueNumber, _agfff._bgafc)
	}
	return _dgge(_gaad, _agfff)
}

// Even is an implementation of the Excel EVEN() that rounds a number to the
// nearest even integer.
func Even(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0045\u0056\u0045\u004e(\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_gfdd := args[0].AsNumber()
	if _gfdd.Type != ResultTypeNumber {
		return MakeErrorResult("\u0045\u0056\u0045N\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ddfe := _cc.Signbit(_gfdd.ValueNumber)
	_degcg, _ffbg := _cc.Modf(_gfdd.ValueNumber / 2)
	_egbbb := _degcg * 2
	if _ffbg != 0 {
		if !_ddfe {
			_egbbb += 2
		} else {
			_egbbb -= 2
		}
	}
	return MakeNumberResult(_egbbb)
}

// Text is an implementation of the Excel TEXT function.
func Text(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("T\u0045\u0058\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	_ebfdec := args[0]
	if _ebfdec.Type != ResultTypeNumber && _ebfdec.Type != ResultTypeString && _ebfdec.Type != ResultTypeEmpty {
		return MakeErrorResult("\u0054\u0045\u0058\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020a\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	if args[1].Type != ResultTypeString {
		return MakeErrorResult("\u0054E\u0058\u0054 \u0072\u0065\u0071\u0075i\u0072\u0065\u0073 \u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072gu\u006d\u0065\u006et\u0020\u0074o\u0020\u0062\u0065\u0020\u0061\u0020s\u0074\u0072i\u006e\u0067")
	}
	_eeae := args[1].ValueString
	switch _ebfdec.Type {
	case ResultTypeNumber:
		return MakeStringResult(_fg.Number(_ebfdec.ValueNumber, _eeae))
	case ResultTypeString:
		return MakeStringResult(_fg.String(_ebfdec.ValueString, _eeae))
	case ResultTypeEmpty:
		return MakeStringResult(_fg.Number(0, _eeae))
	case ResultTypeArray, ResultTypeList:
		return MakeErrorResultType(ErrorTypeSpill, "\u0054\u0045X\u0054\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0077\u006f\u0072\u006b\u0020\u0077\u0069\u0074\u0068\u0020\u0061\u0072ra\u0079\u0073")
	default:
		return MakeErrorResult("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0020\u0066\u006fr\u0020T\u0045\u0058\u0054")
	}
}

// IsBool returns false for the invalid reference context.
func (_eeadd *ivr) IsBool(cellRef string) bool { return false }

const _gebf = 187

func (_fcfdg VerticalRange) verticalRangeReference() string {
	return _e.Sprintf("\u0025\u0073\u003a%\u0073", _fcfdg._ggcb, _fcfdg._acdd)
}
func _cfga(_cgaad float64) float64 {
	_gfee := float64(1)
	for _edda := float64(2); _edda <= _cgaad; _edda++ {
		_gfee *= _edda
	}
	return _gfee
}
func _cccgb(_eaea []string, _ffeaa int) string { return _afc.Itoa(len(_eaea[len(_eaea)-1-_ffeaa])) }

const _gcgg = 57363

// RegisterFunctionComplex registers a standard function.
func RegisterFunctionComplex(name string, fn FunctionComplex) {
	_agde.Lock()
	defer _agde.Unlock()
	if _, _cgbdc := _aaebee[name]; _cgbdc {
		_cf.Log.Debug("\u0064\u0075p\u006c\u0069\u0063\u0061t\u0065\u0020r\u0065\u0067\u0069\u0073\u0074\u0072\u0061\u0074i\u006f\u006e\u0020\u006f\u0066\u0020\u0066\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0025\u0073", name)
	}
	_aaebee[name] = fn
}

// PrefixExpr is an expression containing reference to another sheet like Sheet1!A1 (the value of the cell A1 from sheet 'Sheet1').
type PrefixExpr struct {
	_dfgab Expression
	_bcada Expression
}

const _eafe = 57358

// FunctionCall is a function call expression.
type FunctionCall struct {
	_eecfd string
	_dgce  []Expression
}

func (_eegae *noCache) GetFromCache(key string) (Result, bool) { return _feb, false }

const _eabc = 57362
const _cdae = 57377

func (_fgeb tokenType) String() string { return _dgda(int(_fgeb)) }

// MakeErrorResultType makes an error result of a given type with a specified
// debug message
func MakeErrorResultType(t ErrorType, msg string) Result {
	switch t {
	case ErrorTypeNull:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u004e\u0055\u004c\u004c\u0021", ErrorMessage: msg}
	case ErrorTypeValue:
		return Result{Type: ResultTypeError, ValueString: "\u0023V\u0041\u004c\u0055\u0045\u0021", ErrorMessage: msg}
	case ErrorTypeRef:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u0052\u0045F\u0021", ErrorMessage: msg}
	case ErrorTypeName:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u004e\u0041\u004d\u0045\u003f", ErrorMessage: msg}
	case ErrorTypeNum:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u004e\u0055M\u0021", ErrorMessage: msg}
	case ErrorTypeSpill:
		return Result{Type: ResultTypeError, ValueString: "\u0023S\u0050\u0049\u004c\u004c\u0021", ErrorMessage: msg}
	case ErrorTypeNA:
		return Result{Type: ResultTypeError, ValueString: "\u0023\u004e\u002f\u0041", ErrorMessage: msg}
	case ErrorTypeDivideByZero:
		return Result{Type: ResultTypeError, ValueString: "\u0023D\u0049\u0056\u002f\u0030\u0021", ErrorMessage: msg}
	default:
		return Result{Type: ResultTypeError, ValueString: "\u0023V\u0041\u004c\u0055\u0045\u0021", ErrorMessage: msg}
	}
}

// Eval evaluates a horizontal range with prefix returning a list of results or an error.
func (_bdcdd PrefixHorizontalRange) Eval(ctx Context, ev Evaluator) Result {
	_fadd := _bdcdd._dggf.Reference(ctx, ev)
	switch _fadd.Type {
	case ReferenceTypeSheet:
		if _dcbb(_fadd, ctx) {
			return MakeErrorResultType(ErrorTypeName, _e.Sprintf("\u0053h\u0065e\u0074\u0020\u0025\u0073\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064", _fadd.Value))
		}
		_fgbdd := _bdcdd.horizontalRangeReference(_fadd.Value)
		if _cffee, _gfa := ev.GetFromCache(_fgbdd); _gfa {
			return _cffee
		}
		_dgbff := ctx.Sheet(_fadd.Value)
		_ceaag, _dbfbg := _gdbdc(_dgbff, _bdcdd._egbcc, _bdcdd._bcaa)
		_ecab := _gebc(_dgbff, ev, _ceaag, _dbfbg)
		ev.SetCache(_fgbdd, _ecab)
		return _ecab
	default:
		return MakeErrorResult(_e.Sprintf("\u006e\u006f\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0066\u006f\u0072\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _fadd.Type))
	}
}

// CeilingMath implements _xlfn.CEILING.MATH which rounds numbers to the nearest
// multiple of the second argument, toward or away from zero as specified by the
// third argument.
func CeilingMath(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0043E\u0049\u004cI\u004e\u0047\u002eM\u0041\u0054\u0048\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006ee\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u0043E\u0049\u004cI\u004e\u0047\u002eM\u0041\u0054\u0048\u0028\u0029\u0020\u0061l\u006c\u006f\u0077\u0073\u0020\u0061t\u0020\u006d\u006f\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_daace := args[0].AsNumber()
	if _daace.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072\u0073\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u002e\u004dA\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061 \u006eu\u006d\u0062\u0065\u0072")
	}
	_cded := float64(1)
	if _daace.ValueNumber < 0 {
		_cded = -1
	}
	if len(args) > 1 {
		_cabag := args[1].AsNumber()
		if _cabag.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f \u0043\u0045\u0049\u004c\u0049\u004e\u0047.\u004d\u0041\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_cded = _cabag.ValueNumber
	}
	_cadc := float64(1)
	if len(args) > 2 {
		_gaaf := args[2].AsNumber()
		if _gaaf.Type != ResultTypeNumber {
			return MakeErrorResult("\u0074\u0068\u0069\u0072\u0064\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u002e\u004dA\u0054\u0048\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061 \u006eu\u006d\u0062\u0065\u0072")
		}
		_cadc = _gaaf.ValueNumber
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Ceil(_daace.ValueNumber))
	}
	_ggaf := _daace.ValueNumber
	_ggaf, _bbcfb := _cc.Modf(_ggaf / _cded)
	if _bbcfb != 0 {
		if _daace.ValueNumber > 0 {
			_ggaf++
		} else if _cadc < 0 {
			_ggaf--
		}
	}
	return MakeNumberResult(_ggaf * _cded)
}

var _gcfd = []ri{{1000, "\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}

type yyParserImpl struct {
	_bggb  yySymType
	_bbbbg [_edeae]yySymType
	_gaeec int
}

var _fbgc = map[string]*_f.Regexp{}

const _gaeff = 57353

func _dgda(_bbgcf int) string {
	if _bbgcf >= 1 && _bbgcf-1 < len(_fceb) {
		if _fceb[_bbgcf-1] != "" {
			return _fceb[_bbgcf-1]
		}
	}
	return _e.Sprintf("\u0074\u006f\u006b\u002d\u0025\u0076", _bbgcf)
}

// Tbillyield implements the Excel TBILLYIELD function.
func Tbillyield(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("T\u0042\u0049\u004c\u004c\u0059\u0049E\u004c\u0044\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_abbf, _ecfge, _bceee := _gfcg(args[0], args[1], "\u0054\u0042\u0049\u004c\u004c\u0059\u0049\u0045\u004c\u0044")
	if _bceee.Type == ResultTypeError {
		return _bceee
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0054\u0042\u0049\u004c\u004c\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0064\u0069\u0073\u0063\u006f\u0075n\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fadeb := _ecfge - _abbf
	if _fadeb > 365 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004cY\u0049\u0045\u004c\u0044\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020m\u0061\u0074\u0075r\u0069\u0074\u0079\u0020t\u006f\u0020\u0062\u0065\u0020\u006eo\u0074\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065\u0020\u0079e\u0061\u0072\u0020\u0061\u0066\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074")
	}
	_abgc := args[2].ValueNumber
	if _abgc <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004c\u0059\u0049\u0045\u004c\u0044\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020p\u0072 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fgad := (100 - _abgc) / _abgc
	_bgfc := 360 / _fadeb
	return MakeNumberResult(_fgad * _bgfc)
}

const _ebac = 57375

// Combin is an implementation of the Excel COMBINA function whic returns the
// number of combinations.
func Combin(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0043\u004f\u004d\u0042\u0049\u004e\u0028\u0029\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_dgad := args[0].AsNumber()
	_adca := args[1].AsNumber()
	if _dgad.Type != ResultTypeNumber || _adca.Type != ResultTypeNumber {
		return MakeErrorResult("C\u004f\u004d\u0042\u0049\u004e\u0028)\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006d\u0065r\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_aagfg := _cc.Trunc(_dgad.ValueNumber)
	_gcecg := _cc.Trunc(_adca.ValueNumber)
	if _gcecg > _aagfg {
		return MakeErrorResult("\u0043O\u004d\u0042\u0049\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006b\u0020\u003c\u003d\u0020\u006e")
	}
	if _gcecg == _aagfg || _gcecg == 0 {
		return MakeNumberResult(1)
	}
	_gfdg := float64(1)
	for _bfbf := float64(1); _bfbf <= _gcecg; _bfbf++ {
		_gfdg *= (_aagfg + 1 - _bfbf) / _bfbf
	}
	return MakeNumberResult(_gfdg)
}

// Count implements the COUNT function.
func Count(args []Result) Result { return MakeNumberResult(_aabce(args, _fagga)) }

// Couppcd implements the Excel COUPPCD function.
func Couppcd(args []Result) Result {
	_egfa, _dedg := _cebf(args, "\u0043O\u0055\u0050\u0050\u0043\u0044")
	if _dedg.Type == ResultTypeError {
		return _dedg
	}
	_fgcb := _bgee(_egfa._bbaa)
	_daede := _bgee(_egfa._aee)
	_cbcd := _egfa._gbf
	_bfc := _egfa._ceea
	_aff := _ccefc(_fgcb, _daede, _cbcd, _bfc)
	_fadb, _ddgc, _egae := _aff.Date()
	return MakeNumberResult(_bad(_fadb, int(_ddgc), _egae))
}

// String returns a string representation of SheetPrefixExpr.
func (_facd SheetPrefixExpr) String() string { return _facd._bgfda }

// GetEpoch returns a null time object for the invalid reference context.
func (_ddag *ivr) GetEpoch() _eb.Time { return _eb.Time{} }
func _gaacb(_egge, _bdea int) string {
	const TOKSTART = 4
	if !_bdff {
		return "\u0073\u0079\u006et\u0061\u0078\u0020\u0065\u0072\u0072\u006f\u0072"
	}
	for _, _eacdg := range _febb {
		if _eacdg._afac == _egge && _eacdg._ebae == _bdea {
			return "\u0073\u0079\u006e\u0074\u0061\u0078\u0020\u0065\u0072r\u006f\u0072\u003a\u0020" + _eacdg._agbgd
		}
	}
	_ggfc := "\u0073y\u006e\u0074\u0061\u0078 \u0065\u0072\u0072\u006f\u0072:\u0020u\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064 " + _dgda(_bdea)
	_cfaagc := make([]int, 0, 4)
	_acfde := _gded[_egge]
	for _bfefg := TOKSTART; _bfefg-1 < len(_fceb); _bfefg++ {
		if _efgg := _acfde + _bfefg; _efgg >= 0 && _efgg < _gebf && _cgec[_gbeb[_efgg]] == _bfefg {
			if len(_cfaagc) == cap(_cfaagc) {
				return _ggfc
			}
			_cfaagc = append(_cfaagc, _bfefg)
		}
	}
	if _ebgdg[_egge] == -2 {
		_aaf := 0
		for _bcab[_aaf] != -1 || _bcab[_aaf+1] != _egge {
			_aaf += 2
		}
		for _aaf += 2; _bcab[_aaf] >= 0; _aaf += 2 {
			_fdfc := _bcab[_aaf]
			if _fdfc < TOKSTART || _bcab[_aaf+1] == 0 {
				continue
			}
			if len(_cfaagc) == cap(_cfaagc) {
				return _ggfc
			}
			_cfaagc = append(_cfaagc, _fdfc)
		}
		if _bcab[_aaf+1] != 0 {
			return _ggfc
		}
	}
	for _cdgcec, _daeg := range _cfaagc {
		if _cdgcec == 0 {
			_ggfc += "\u002c\u0020\u0065x\u0070\u0065\u0063\u0074\u0069\u006e\u0067\u0020"
		} else {
			_ggfc += "\u0020\u006f\u0072\u0020"
		}
		_ggfc += _dgda(_daeg)
	}
	return _ggfc
}

// Replace is an implementation of the Excel REPLACE().
func Replace(args []Result) Result {
	_cece, _ebbb := _fbcad("\u0052E\u0050\u004c\u0041\u0043\u0045", args)
	if _ebbb.Type != ResultTypeEmpty {
		return _ebbb
	}
	_cbcfad := _cece._bdfaf
	_cgdad := _cece._bebed
	_cdaa := _cece._ecadb
	_egebc := _cece._cbegb
	_gaafd := len(_cbcfad)
	if _cgdad > _gaafd {
		_cgdad = _gaafd
	}
	_agbg := _cgdad + _cdaa
	if _agbg > _gaafd {
		_agbg = _gaafd
	}
	_bffc := _cbcfad[0:_cgdad] + _egebc + _cbcfad[_agbg:]
	return MakeStringResult(_bffc)
}

const _egefd int = 30

// CeilingPrecise is an implementation of the CEILING.PRECISE function which
// returns the ceiling of a number.
func CeilingPrecise(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u002eP\u0052\u0045\u0043IS\u0045\u0028\u0029\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020o\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("\u0043\u0045I\u004c\u0049\u004e\u0047\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0074\u0020\u006d\u006f\u0073\u0074\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ebbce := args[0].AsNumber()
	if _ebbce.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069r\u0073\u0074\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074 \u0074\u006f\u0020\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_cbac := float64(1)
	if _ebbce.ValueNumber < 0 {
		_cbac = -1
	}
	if len(args) > 1 {
		_cafe := args[1].AsNumber()
		if _cafe.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0043E\u0049L\u0049\u004e\u0047\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_cbac = _cc.Abs(_cafe.ValueNumber)
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Ceil(_ebbce.ValueNumber))
	}
	_gage := _ebbce.ValueNumber
	_gage, _abfd := _cc.Modf(_gage / _cbac)
	if _abfd != 0 {
		if _ebbce.ValueNumber > 0 {
			_gage++
		}
	}
	return MakeNumberResult(_gage * _cbac)
}
func _gcga(_dfdd []Result, _efdb string) (float64, float64, Result) {
	if len(_dfdd) != 2 {
		return 0, 0, MakeErrorResult(_efdb + "\u0020\u0072\u0065qu\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _dfdd[0].Type != ResultTypeNumber {
		return 0, 0, MakeErrorResult(_efdb + "\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0064\u006f\u006c\u006c\u0061\u0072 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_fcgd := _dfdd[0].ValueNumber
	if _dfdd[1].Type != ResultTypeNumber {
		return 0, 0, MakeErrorResult(_efdb + " \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_acae := float64(int(_dfdd[1].ValueNumber))
	if _acae < 0 {
		return 0, 0, MakeErrorResultType(ErrorTypeNum, _efdb+"\u0020r\u0065\u0071u\u0069\u0072\u0065\u0073 \u0066\u0072\u0061c\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062e \u006e\u006f\u006e \u006e\u0065g\u0061\u0074\u0069\u0076\u0065\u0020n\u0075\u006db\u0065\u0072")
	}
	return _fcgd, _acae, _feb
}
func _fbebf(_bddf, _cfe float64, _cgcf, _cbegc int) float64 {
	_eadd := _bgee(_bddf)
	_eecb := _bgee(_cfe)
	_gfcc := _ccefc(_eadd, _eecb, _cgcf, _cbegc)
	return _cbea(_gfcc, _eadd, _cbegc)
}

const _ffdf = _eb.Second * 1

type rangeIndex struct {
	_aaba int
	_gbfb int
}

const _dgfg = 57374

// If is an implementation of the Excel IF() function. It takes one, two or
// three arguments.
func If(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u0049\u0046\u0020re\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074 \u006ce\u0061s\u0074 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u0049\u0046\u0020ac\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0074\u0020m\u006fs\u0074 \u0074h\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_aggbg := args[0]
	switch _aggbg.Type {
	case ResultTypeError:
		return _aggbg
	case ResultTypeNumber:
		if len(args) == 1 {
			return MakeBoolResult(_aggbg.ValueNumber != 0)
		}
		if _aggbg.ValueNumber != 0 {
			return args[1]
		}
		if len(args) == 3 {
			return args[2]
		} else {
			return MakeBoolResult(false)
		}
	case ResultTypeList:
		return _faca(args)
	case ResultTypeArray:
		return _bgeeg(args)
	default:
		return MakeErrorResult("\u0049F\u0020\u0069n\u0069\u0074\u0069\u0061l\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u006d\u0075\u0073t \u0062\u0065\u0020n\u0075\u006de\u0072\u0069\u0063\u0020\u006f\u0072 \u0061\u0072r\u0061\u0079")
	}
}

const _degcb = 57347

// NewBinaryExpr constructs a new binary expression with a given operator.
func NewBinaryExpr(lhs Expression, op BinOpType, rhs Expression) Expression {
	return BinaryExpr{_gb: lhs, _gba: rhs, _fc: op}
}

var _acbd float64 = 25569.0

// FloorPrecise is an implementation of the FlOOR.PRECISE function.
func FloorPrecise(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0046\u004cO\u004f\u0052\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("\u0046L\u004f\u004fR\u002e\u0050\u0052\u0045C\u0049\u0053\u0045(\u0029\u0020\u0061\u006c\u006c\u006f\u0077\u0073\u0020at\u0020\u006d\u006fs\u0074\u0020t\u0077\u006f\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_gdfcc := args[0].AsNumber()
	if _gdfcc.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020F\u004c\u004f\u004f\u0052\u002e\u0050\u0052E\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_gbcea := float64(1)
	if _gdfcc.ValueNumber < 0 {
		_gbcea = -1
	}
	if len(args) > 1 {
		_dfee := args[1].AsNumber()
		if _dfee.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073\u0065\u0063\u006f\u006ed\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020F\u004c\u004f\u004f\u0052\u002e\u0050\u0052\u0045\u0043\u0049\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065r")
		}
		_gbcea = _cc.Abs(_dfee.ValueNumber)
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Floor(_gdfcc.ValueNumber))
	}
	_edef := _gdfcc.ValueNumber
	_edef, _bgcf := _cc.Modf(_edef / _gbcea)
	if _bgcf != 0 {
		if _gdfcc.ValueNumber < 0 {
			_edef--
		}
	}
	return MakeNumberResult(_edef * _gbcea)
}

// LastColumn returns empty string for the invalid reference context.
func (_dagcc *ivr) LastColumn(rowFrom, rowTo int) string { return "" }
func _fgfcf(_acgb float64, _bccg *criteriaRegex) bool {
	_deagb, _caaa := _afc.ParseFloat(_bccg._gbcfa, 64)
	if _caaa != nil {
		return false
	}
	switch _bccg._afag {
	case _bfcd:
		return _acgb == _deagb
	case _eggf:
		return _acgb <= _deagb
	case _edfe:
		return _acgb >= _deagb
	case _cdbac:
		return _acgb < _deagb
	case _dfef:
		return _acgb > _deagb
	}
	return false
}

// ISREF is an implementation of the Excel ISREF() function.
func IsRef(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u0052\u0045\u0046\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeBoolResult(ev.LastEvalIsRef())
}
func _cefd(_gbd, _deag float64, _caef, _cgeg int) (float64, Result) {
	_adbd, _cfgd := _bgee(_gbd), _bgee(_deag)
	if _cfgd.After(_adbd) {
		_gdd := _ccefc(_adbd, _cfgd, _caef, _cgeg)
		_eefc := (_cfgd.Year()-_gdd.Year())*12 + int(_cfgd.Month()) - int(_gdd.Month())
		return float64(_eefc*_caef) / 12.0, _feb
	}
	return 0, MakeErrorResultType(ErrorTypeNum, "\u0053\u0065t\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0062\u0065\u0066\u006f\u0072\u0065\u0020\u006d\u0061\u0074\u0075\u0072\u0069\u0074\u0079\u0020\u0064\u0061\u0074\u0065")
}

// String returns a string representation of Number.
func (_dddd Number) String() string { return _afc.FormatFloat(_dddd._fbfe, 'f', -1, 64) }

// String returns a string representation of a horizontal range with prefix.
func (_fafe PrefixHorizontalRange) String() string {
	return _e.Sprintf("\u0025\u0073\u0021\u0025\u0064\u003a\u0025\u0064", _fafe._dggf.String(), _fafe._egbcc, _fafe._bcaa)
}
func _bcd(_dceb, _acag, _egfc float64) float64 { return (_dceb*3600 + _acag*60 + _egfc) / 86400 }

// RoundDown is an implementation of the Excel ROUNDDOWN function that rounds a number
// down to a specified number of digits.
func RoundDown(args []Result) Result { return _afdcc(args, _accd) }
func _cdbdg(_egaee, _daced, _acef Reference) string {
	return _e.Sprintf("\u0025\u0073\u0021\u0025\u0073\u003a\u0025\u0073", _egaee.Value, _daced.Value, _acef.Value)
}

var _gdg = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// Row implements the Excel ROW function.
func Row(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0052O\u0057\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065 \u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_daad := args[0].Ref
	if _daad.Type != ReferenceTypeCell {
		return MakeErrorResult("\u0052\u004f\u0057\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073 a\u006e\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079p\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065n\u0063\u0065")
	}
	_cgfa, _daged := _fb.ParseCellReference(_daad.Value)
	if _daged != nil {
		return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _daad.Value)
	}
	return MakeNumberResult(float64(_cgfa.RowIdx))
}

const _bbg = "\u0028\u0020\u0028" + _ceb + "\u007c" + _cfc + "\u007c" + _dd + "\u007c" + _fed + "\u0029\u0029\u003f\u0024"

// Coupdays implements the Excel COUPDAYS function.
func Coupdays(args []Result) Result {
	_fde, _ega := _cebf(args, "\u0043\u004f\u0055\u0050\u0044\u0041\u0059\u0053")
	if _ega.Type == ResultTypeError {
		return _ega
	}
	return MakeNumberResult(_def(_fde._bbaa, _fde._aee, _fde._gbf, _fde._ceea))
}

const _dfdcd = 1

func _adgeff(_gaeba _eb.Time) bool { return _eb.Now().Sub(_gaeba) >= _bcfe }

// Base is an implementation of the Excel BASE function that returns a string
// form of an integer in a specified base and of a minimum length with padded
// zeros.
func Base(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u0042\u0041\u0053\u0045\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 \u0074\u0077\u006f\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u0042\u0041S\u0045\u0028\u0029\u0020a\u006c\u006co\u0077\u0073\u0020\u0061\u0074\u0020\u006d\u006fs\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_acba := args[0].AsNumber()
	if _acba.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066\u0069\u0072\u0073\u0074 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0042A\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_fcfc := args[1].AsNumber()
	if _fcfc.Type != ResultTypeNumber {
		return MakeErrorResult("\u0073\u0065\u0063o\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0042\u0041\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_abeea := int(_fcfc.ValueNumber)
	if _abeea < 0 || _abeea > 36 {
		return MakeErrorResult("\u0072\u0061\u0064\u0069\u0078\u0020m\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0074\u0068\u0065 \u0072\u0061\u006e\u0067\u0065\u0020\u005b0\u002c\u0033\u0036\u005d")
	}
	_dgfb := 0
	if len(args) > 2 {
		_aecb := args[2].AsNumber()
		if _aecb.Type != ResultTypeNumber {
			return MakeErrorResult("\u0074\u0068\u0069\u0072\u0064 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0042A\u0053\u0045\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_dgfb = int(_aecb.ValueNumber)
	}
	_gfcb := _afc.FormatInt(int64(_acba.ValueNumber), _abeea)
	if len(_gfcb) < _dgfb {
		_gfcb = _ga.Repeat("\u0030", _dgfb-len(_gfcb)) + _gfcb
	}
	return MakeStringResult(_gfcb)
}

// Substitute is an implementation of the Excel SUBSTITUTE function.
func Substitute(args []Result) Result {
	_dcdcc := len(args)
	if _dcdcc != 3 && _dcdcc != 4 {
		return MakeErrorResult("\u0053\u0055\u0042\u0053\u0054\u0049\u0054U\u0054\u0045\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u006f\u0072\u0020\u0066\u006f\u0075\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_fgag, _debf := _deece(args[0], "\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", "\u0074\u0065\u0078\u0074")
	if _debf.Type == ResultTypeError {
		return _debf
	}
	_fgage, _debf := _deece(args[1], "\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", "\u006f\u006c\u0064\u0020\u0074\u0065\u0078\u0074")
	if _debf.Type == ResultTypeError {
		return _debf
	}
	_bfade, _debf := _deece(args[2], "\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", "\u006e\u0065\u0077\u0020\u0074\u0065\u0078\u0074")
	if _debf.Type == ResultTypeError {
		return _debf
	}
	_bbad := 0
	if _dcdcc == 3 {
		return MakeStringResult(_ga.Replace(_fgag, _fgage, _bfade, -1))
	} else {
		_fabbg, _edgd := _bfeef(args[3], "\u0053\u0055\u0042\u0053\u0054\u0049\u0054\u0055\u0054\u0045", "\u0069\u006e\u0073t\u0061\u006e\u0063\u0065\u005f\u006e\u0075\u006d")
		if _edgd.Type == ResultTypeError {
			return _edgd
		}
		_bbad = int(_fabbg)
		if _bbad < 1 {
			return MakeErrorResult("\u0069\u006es\u0074\u0061\u006e\u0063e\u005f\u006eu\u006d\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e \u007a\u0065\u0072\u006f")
		}
		_daea := _fgag
		_debaf := _bbad
		_aeg := -1
		_ecdf := len(_fgage)
		_adecg := 0
		for {
			_debaf--
			_deda := _ga.Index(_daea, _fgage)
			if _deda == -1 {
				_aeg = -1
				break
			} else {
				_aeg = _deda + _adecg
				if _debaf == 0 {
					break
				}
				_geac := _ecdf + _deda
				_adecg += _geac
				_daea = _daea[_geac:]
			}
		}
		if _aeg == -1 {
			return MakeStringResult(_fgag)
		} else {
			_ebbg := _fgag[:_aeg]
			_agea := _fgag[_aeg+_ecdf:]
			return MakeStringResult(_ebbg + _bfade + _agea)
		}
	}
}
func _gdbdc(_bgcc Context, _eafff, _gefef int) (string, string) {
	_begee := "\u0041" + _afc.Itoa(_eafff)
	_ccega := _bgcc.LastColumn(_eafff, _gefef)
	_gdgd := _ccega + _afc.Itoa(_gefef)
	return _begee, _gdgd
}
func _gceb(_gfbc []Result, _fgfbf bool) Result {
	var _dcfa string
	if _fgfbf {
		_dcfa = "\u004c\u0041\u0052G\u0045"
	} else {
		_dcfa = "\u0053\u004d\u0041L\u004c"
	}
	if len(_gfbc) != 2 {
		return MakeErrorResult(_dcfa + "\u0020\u0072\u0065qu\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_deec := _gfbc[0]
	var _adbdc [][]Result
	switch _deec.Type {
	case ResultTypeArray:
		_adbdc = _deec.ValueArray
	case ResultTypeList:
		_adbdc = [][]Result{_deec.ValueList}
	default:
		return MakeErrorResult(_dcfa + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u006f\u0066\u0020\u0074\u0079p\u0065\u0020a\u0072\u0072\u0061\u0079")
	}
	if len(_adbdc) == 0 {
		return MakeErrorResult(_dcfa + "\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0061\u0072\u0072\u0061\u0079\u0020\u0074\u006f\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u0031\u0020\u0072\u006f\u0077")
	}
	if _gfbc[1].Type != ResultTypeNumber {
		return MakeErrorResult(_dcfa + " \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072g\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074yp\u0065\u0020\u006eu\u006db\u0065\u0072")
	}
	_gbce := _gfbc[1].ValueNumber
	if _gbce < 1 {
		return MakeErrorResultType(ErrorTypeNum, _dcfa+"\u0020\u0072e\u0071\u0075\u0069\u0072\u0065s\u0020\u0073\u0065\u0063\u006fn\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
	}
	_ffce := int(_gbce)
	if float64(_ffce) != _gbce {
		return MakeErrorResultType(ErrorTypeNum, _dcfa+"\u0020\u0072e\u0071\u0075\u0069\u0072\u0065s\u0020\u0073\u0065\u0063\u006fn\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u0030")
	}
	_beeg := []float64{}
	for _, _ddgb := range _adbdc {
		for _, _ddfcf := range _ddgb {
			if _ddfcf.Type == ResultTypeNumber {
				_beeg = append(_beeg, _ddfcf.ValueNumber)
			}
		}
	}
	if _ffce > len(_beeg) {
		return MakeErrorResultType(ErrorTypeNum, _dcfa+" \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u006c\u0065s\u0073\u0020\u006f\u0072\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u0068\u0061\u006e\u0020t\u0068\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u006e\u0075\u006d\u0062\u0065\u0072s\u0020\u0069\u006e\u0020t\u0068\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_badgd := _ad.MergeSort(_beeg)
	if _fgfbf {
		return MakeNumberResult(_badgd[len(_badgd)-_ffce])
	} else {
		return MakeNumberResult(_badgd[_ffce-1])
	}
}

// String returns a string representation of String.
func (_dfab String) String() string { return "\u0022" + _dfab._bedb + "\u0022" }

var _cgec = [...]int{-1000, -7, -3, -1, 27, 18, 22, 23, -2, -8, -4, -9, 20, -14, 10, 11, 12, 13, -5, -13, -6, -12, 17, 16, 15, 9, 4, 5, 22, 23, 24, 25, 26, 28, 29, 30, 31, 27, 32, 35, -1, 18, 27, -15, -17, -1, -1, -1, -1, 33, -5, 4, 5, 21, -16, -11, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 19, 36, 34, 21, -5, 33, 21, 34, 19, -17, -1, -5, -10, -1}

// Syd implements the Excel SYD function.
func Syd(args []Result) Result {
	if len(args) != 4 {
		return MakeErrorResult("S\u0059\u0044\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0066\u006f\u0075\u0072 \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0059\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020c\u006f\u0073\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_cbaf := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0059\u0044 \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_bfdf := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0059\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020l\u0069\u0066\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_cfge := args[2].ValueNumber
	if _cfge <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0053\u0059\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0059\u0044\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_afcgf := args[3].ValueNumber
	if _afcgf <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0053\u0059\u0044 r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070e\u0072i\u006fd\u0020t\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if _afcgf > _cfge {
		return MakeErrorResultType(ErrorTypeNum, "\u0053\u0059\u0044\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0065q\u0075\u0061\u006c\u0020\u006f\u0072\u0020\u006c\u0065\u0073\u0073\u0020\u0074\u0068a\u006e \u006c\u0069\u0066\u0065")
	}
	_egdb := (_cbaf - _bfdf) * (_cfge - _afcgf + 1) * 2
	_ggcg := _cfge * (_cfge + 1)
	return MakeNumberResult(_egdb / _ggcg)
}

// Fvschedule implements the Excel FVSCHEDULE function.
func Fvschedule(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0046\u0056\u0053\u0043\u0048\u0045D\u0055\u004c\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0077\u006f\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0046\u0056\u0053\u0043\u0048E\u0044\u0055\u004c\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0070\u0072\u0069\u006e\u0063\u0069\u0070\u0061\u006c\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_badg := args[0].ValueNumber
	switch args[1].Type {
	case ResultTypeNumber:
		return MakeNumberResult(_badg * (args[1].ValueNumber + 1))
	case ResultTypeList, ResultTypeArray:
		_gbdg := _aadag(args[1])
		for _, _bafd := range _gbdg {
			for _, _aeef := range _bafd {
				if _aeef.Type != ResultTypeNumber || _aeef.IsBoolean {
					return MakeErrorResult("\u0046\u0056\u0053\u0043\u0048\u0045\u0044\u0055\u004c\u0045\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020r\u0061\u0074\u0065\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0073")
				}
				_badg *= 1.0 + _aeef.ValueNumber
			}
		}
		return MakeNumberResult(_badg)
	default:
		return MakeErrorResult("\u0046\u0056\u0053\u0043\u0048\u0045\u0044\u0055\u004c\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0063\u0068\u0065\u0064\u0075\u006c\u0065\u0020\u0074o\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0072\u0020\u0061\u0072\u0072a\u0079\u0020\u0074y\u0070\u0065")
	}
}

var _bdebe = []ri{{1000, "\u004d"}, {950, "\u004c\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {450, "\u004c\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {95, "\u0056\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {45, "\u0056\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}

// String returns an empty string for EmptyExpr.
func (_ggf EmptyExpr) String() string { return "" }

// Eval evaluates a range returning a list of results or an error.
func (_dbgba Range) Eval(ctx Context, ev Evaluator) Result {
	_cgece := _dbgba._fbdb.Reference(ctx, ev)
	_debfe := _dbgba._agedb.Reference(ctx, ev)
	_agbd := _ccaeb(_cgece, _debfe)
	if _cgece.Type == ReferenceTypeCell && _debfe.Type == ReferenceTypeCell {
		if _dgec, _cgefc := ev.GetFromCache(_agbd); _cgefc {
			return _dgec
		} else {
			_daae := _gebc(ctx, ev, _cgece.Value, _debfe.Value)
			ev.SetCache(_agbd, _daae)
			return _daae
		}
	}
	return MakeErrorResult("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072a\u006e\u0067\u0065\u0020" + _agbd)
}

type noCache struct{}

// Rows implements the Excel ROWS function.
func Rows(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0052\u004f\u0057\u0053\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074")
	}
	_dbbgc := args[0]
	if _dbbgc.Type != ResultTypeArray && _dbbgc.Type != ResultTypeList {
		return MakeErrorResult("\u0052\u004f\u0057S\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074y\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_aebd := _dbbgc.ValueArray
	if len(_aebd) == 0 {
		return MakeErrorResult("\u0052O\u0057\u0053 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0072r\u0061\u0079\u0020\u0074\u006f\u0020c\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0061\u0074\u0020\u006ce\u0061\u0073\u0074\u0020\u0031\u0020\u0072\u006f\u0077")
	}
	return MakeNumberResult(float64(len(_aebd)))
}

// Time is an implementation of the Excel TIME() function.
func Time(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0054\u0049ME\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s t\u0068re\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_gccf := args[0].ValueNumber
	_afcd := args[1].ValueNumber
	_adcg := args[2].ValueNumber
	_dcgb := _bcd(_gccf, _afcd, _adcg)
	if _dcgb >= 0 {
		return MakeNumberResult(_dcgb)
	} else {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
}

var _gded = [...]int{123, -1000, -1000, 74, 163, 103, 163, 163, -1000, -1000, -1000, -1000, 163, -1000, -1000, -1000, -1000, -1000, -12, 106, -1000, -1000, 143, -1000, -1000, -1000, -1000, -1000, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 163, 74, 163, 163, 6, -28, 74, -15, -15, 60, 10, -14, -1000, -1000, -1000, 7, -1000, 74, -15, -15, -23, -23, -1000, -8, -8, -8, -8, -8, -8, -4, 33, -1000, 163, 163, -1000, -1000, 10, -1000, 163, -1000, -28, 74, -1000, -1000, 74}

func _bgee(_bebf float64) _eb.Time {
	_eefg := int64((_bebf - _acbd) * _bbgc)
	return _eb.Unix(0, _eefg).UTC()
}

// ReferenceType is a type of reference
//
//go:generate stringer -type=ReferenceType
type ReferenceType byte

func (_decg HorizontalRange) horizontalRangeReference() string {
	return _e.Sprintf("\u0025\u0064\u003a%\u0064", _decg._fcbd, _decg._fbcg)
}

// NewEvaluator constructs a new defEval object which is the default formula evaluator.
func NewEvaluator() Evaluator { _aag := &defEval{}; _aag.evCache = _ece(); return _aag }
func _eged(_bdgde []Result, _bgfd bool, _deea string) Result {
	var _ebcfc, _ebffdb string
	if _bgfd {
		_ebcfc = "\u0074\u0068\u0072e\u0065"
		_ebffdb = "\u006f\u0064\u0064"
	} else {
		_ebcfc = "\u0074\u0077\u006f"
		_ebffdb = "\u0065\u0076\u0065\u006e"
	}
	_gfeb := len(_bdgde)
	if (_bgfd && _gfeb < 3) || (!_bgfd && _gfeb < 2) {
		return MakeErrorResult(_deea + "\u0020\u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020" + _ebcfc + " \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0073")
	}
	if (_gfeb/2*2 == _gfeb) == _bgfd {
		return MakeErrorResult(_deea + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020" + _ebffdb + " \u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020a\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	_defe := -1
	_dgbe := -1
	for _aefdg := 0; _aefdg < _gfeb; _aefdg += 2 {
		_eaec := _bdgde[_aefdg]
		if _eaec.Type != ResultTypeArray && _eaec.Type != ResultTypeList {
			return MakeErrorResult(_deea + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u006e\u0067\u0065\u0073\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065 \u006c\u0069\u0073\u0074\u0020o\u0072\u0020a\u0072\u0072\u0061\u0079")
		}
		_fcgc := _aadag(_eaec)
		if _dgbe == -1 {
			_dgbe = len(_fcgc)
			_defe = len(_fcgc[0])
		} else if len(_fcgc) != _dgbe || len(_fcgc[0]) != _defe {
			return MakeErrorResult(_deea + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0061l\u006c\u0020\u0072\u0061n\u0067\u0065\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0073\u0061\u006d\u0065\u0020\u0073\u0069\u007a\u0065")
		}
		if _bgfd && _aefdg == 0 {
			_aefdg--
		}
	}
	return _feb
}

type tokenType int

const (
	_cgae  cmpResult = 0
	_gcefb cmpResult = -1
	_acbcc cmpResult = 1
	_bfaf  cmpResult = 2
)

// Rri implements the Excel RRI function.
func Rri(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0052\u0052\u0049\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065e\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0052I\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_adbe := args[0].ValueNumber
	if _adbe <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052R\u0049\u0020r\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0066\u0020p\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062e\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u0052\u0049\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073 p\u0072\u0065\u0073\u0065\u006e\u0074 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_fcge := args[1].ValueNumber
	if _fcge <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052\u0052\u0049\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065\u006et\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("R\u0052\u0049\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0075\u0074\u0075\u0072e \u0076\u0061\u006c\u0075e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_febf := args[2].ValueNumber
	if _febf < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0052R\u0049\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020f\u0075\u0074\u0075\u0072\u0065\u0020\u0076\u0061\u006cue\u0020\u0074\u006f \u0062\u0065 \u006e\u006f\u006e\u0020\u006e\u0065g\u0061\u0074i\u0076\u0065")
	}
	return MakeNumberResult(_cc.Pow(_febf/_fcge, 1/_adbe) - 1)
}

const _bddc = "\u0028\u0028\u006a\u0061\u006e|\u006a\u0061\u006e\u0075\u0061\u0072\u0079\u0029\u007c\u0028\u0066\u0065\u0062\u007c\u0066\u0065\u0062\u0072\u0075a\u0072\u0079\u0029\u007c\u0028\u006da\u0072\u007c\u006da\u0072\u0063\u0068\u0029\u007c\u0028\u0061\u0070\u0072\u007c\u0061\u0070\u0072\u0069\u006c\u0029\u007c\u0028\u006d\u0061\u0079\u0029\u007c\u0028j\u0075\u006e\u007cj\u0075\u006e\u0065\u0029\u007c\u0028\u006a\u0075\u006c\u007c\u006a\u0075\u006c\u0079\u0029\u007c\u0028a\u0075\u0067\u007c\u0061\u0075\u0067\u0075\u0073t\u0029\u007c\u0028\u0073\u0065\u0070\u007c\u0073\u0065\u0070\u0074\u0065\u006d\u0062\u0065\u0072\u0029\u007c\u0028o\u0063\u0074\u007c\u006f\u0063\u0074\u006f\u0062\u0065\u0072\u0029\u007c\u0028\u006e\u006f\u0076\u007c\u006e\u006f\u0076\u0065\u006d\u0062e\u0072\u0029\u007c\u0028\u0064\u0065\u0063\u007c\u0064\u0065\u0063\u0065\u006d\u0062\u0065\u0072\u0029\u0029"

// DateValue is an implementation of the Excel DATEVALUE() function.
func DateValue(args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeString {
		return MakeErrorResult("\u0044A\u0054\u0045V\u0041\u004c\u0055\u0045 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069ng\u006c\u0065\u0020s\u0074\u0072i\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065n\u0074\u0073")
	}
	_bdf := _ga.ToLower(args[0].ValueString)
	if !_cddc(_bdf) {
		_, _, _, _, _aeb, _fec := _ccce(_bdf)
		if _fec.Type == ResultTypeError {
			_fec.ErrorMessage = "\u0049\u006e\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020f\u006f\u0072\u0020\u0044\u0041\u0054\u0045V\u0041\u004c\u0055\u0045"
			return _fec
		}
		if _aeb {
			return MakeNumberResult(0)
		}
	}
	_dcfc, _fggf, _eda, _, _bggd := _ddd(_bdf)
	if _bggd.Type == ResultTypeError {
		return _bggd
	}
	return MakeNumberResult(_bad(_dcfc, _fggf, _eda))
}

// NewVerticalRange constructs a new full columns range.
func NewVerticalRange(v string) Expression {
	_agdeb := _ga.Split(v, "\u003a")
	if len(_agdeb) != 2 {
		return nil
	}
	if _agdeb[0] > _agdeb[1] {
		_agdeb[0], _agdeb[1] = _agdeb[1], _agdeb[0]
	}
	return VerticalRange{_ggcb: _agdeb[0], _acdd: _agdeb[1]}
}
func (_fedgd *ivr) SetOffset(col, row uint32) {}

// HasFormula returns FALSE for the invalid reference context.
func (_gfgag *ivr) HasFormula(cellRef string) bool { return false }

// Index implements the Excel INDEX function.
func Index(args []Result) Result {
	_dggb := len(args)
	if _dggb < 2 || _dggb > 3 {
		return MakeErrorResult("\u0049\u004e\u0044E\u0058\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0072\u006f\u006d\u0020\u006f\u006e\u0065\u0020\u0074\u006f\u0020\u0074\u0068\u0072\u0065\u0065\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_cbdcf := args[0]
	if _cbdcf.Type != ResultTypeArray && _cbdcf.Type != ResultTypeList {
		return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0072e\u0071\u0075\u0069r\u0065\u0073\u0020\u0066i\u0072\u0073\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_afca := args[1].AsNumber()
	if _afca.Type != ResultTypeNumber {
		return MakeErrorResult("I\u004e\u0044\u0045\u0058\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u006e\u0075\u006d\u0065\u0072\u0069\u0063 \u0072\u006f\u0077\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_bgac := int(_afca.ValueNumber) - 1
	_adcba := -1
	if _dggb == 3 && args[2].Type != ResultTypeEmpty {
		_cbbe := args[2].AsNumber()
		if _cbbe.Type != ResultTypeNumber {
			return MakeErrorResult("I\u004e\u0044\u0045\u0058\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073 \u006e\u0075\u006d\u0065\u0072\u0069\u0063 \u0063\u006f\u006c\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
		}
		_adcba = int(_cbbe.ValueNumber) - 1
	}
	if _bgac == -1 && _adcba == -1 {
		return MakeErrorResult("\u0049\u004e\u0044EX\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0072o\u0077 \u006fr\u0020\u0063\u006f\u006c\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	var _ffg []Result
	if _cbdcf.Type == ResultTypeArray {
		_befe := _cbdcf.ValueArray
		if _bgac < -1 || _bgac >= len(_befe) {
			return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0072o\u0077\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
		}
		if _bgac == -1 {
			if _adcba >= len(_befe[0]) {
				return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0063o\u006c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
			}
			_faea := [][]Result{}
			for _, _ebag := range _befe {
				_cgcg := _ebag[_adcba]
				if _cgcg.Type == ResultTypeEmpty {
					_cgcg = MakeNumberResult(0)
				}
				_faea = append(_faea, []Result{_cgcg})
			}
			return MakeArrayResult(_faea)
		}
		_ffg = _befe[_bgac]
	} else {
		_bgd := _cbdcf.ValueList
		if _bgac < -1 || _bgac >= 1 {
			return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0072o\u0077\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
		}
		if _bgac == -1 {
			if _adcba >= len(_bgd) {
				return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0063o\u006c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
			}
			_cdg := _bgd[_adcba]
			if _cdg.Type == ResultTypeEmpty {
				_cdg = MakeNumberResult(0)
			}
			return _cdg
		}
		_ffg = _bgd
	}
	if _adcba < -1 || _adcba > len(_ffg) {
		return MakeErrorResult("\u0049\u004e\u0044\u0045\u0058\u0020\u0068\u0061\u0073\u0020\u0063o\u006c\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072a\u006e\u0067\u0065")
	}
	if _adcba == -1 {
		_acab := []Result{}
		for _, _ccffg := range _ffg {
			if _ccffg.Type == ResultTypeEmpty {
				_acab = append(_acab, MakeNumberResult(0))
			} else {
				_acab = append(_acab, _ccffg)
			}
		}
		return MakeArrayResult([][]Result{_acab})
	}
	_ceba := _ffg[_adcba]
	if _ceba.Type == ResultTypeEmpty {
		return MakeNumberResult(0)
	}
	return _ceba
}

// String returns a string representation for Bool.
func (_ffd Bool) String() string {
	if _ffd._aef {
		return "\u0054\u0052\u0055\u0045"
	} else {
		return "\u0046\u0041\u004cS\u0045"
	}
}
func (_agbb PrefixVerticalRange) verticalRangeReference(_dcgee string) string {
	return _e.Sprintf("\u0025\u0073\u0021\u0025\u0073\u003a\u0025\u0073", _dcgee, _agbb._agfae, _agbb._fefae)
}
func _eebbg(_ffba _eb.Time) _eb.Time {
	_ffba = _ffba.UTC()
	return _eb.Date(_ffba.Year(), _ffba.Month(), _ffba.Day(), _ffba.Hour(), _ffba.Minute(), _ffba.Second(), _ffba.Nanosecond(), _eb.Local)
}

// Left implements the Excel LEFT(string,[n]) function which returns the
// leftmost n characters.
func Left(args []Result) Result {
	_aecf := 1
	switch len(args) {
	case 1:
	case 2:
		if args[1].Type != ResultTypeNumber {
			return MakeErrorResult("\u004c\u0045F\u0054\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075me\u006e\u0074")
		}
		_aecf = int(args[1].ValueNumber)
		if _aecf < 0 {
			return MakeErrorResult("\u004c\u0045\u0046T \u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020n\u0075m\u0062e\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u003e\u003d\u0020\u0030")
		}
		if _aecf == 0 {
			return MakeStringResult("")
		}
	default:
		return MakeErrorResult("\u004c\u0045\u0046T \u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020o\u006ee\u0020o\u0072 \u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[0].Type == ResultTypeList {
		return MakeErrorResult("\u004c\u0045\u0046T\u0020\u0063\u0061\u006e'\u0074\u0020\u0062\u0065\u0020\u0063\u0061l\u006c\u0065\u0064\u0020\u006f\u006e\u0020\u0061\u0020\u0072\u0061\u006e\u0067\u0065")
	}
	_accef := args[0].Value()
	if _aecf > len(_accef) {
		return MakeStringResult(_accef)
	}
	return MakeStringResult(_accef[0:_aecf])
}
func _ece() evCache {
	_gce := evCache{}
	_gce._bgg = make(map[string]Result)
	_gce._ef = &_af.Mutex{}
	return _gce
}
func _ebbcf(_eecc []Result, _gfgf bool) (float64, float64) {
	_bfbed := 0.0
	_dece := 0.0
	for _, _facga := range _eecc {
		switch _facga.Type {
		case ResultTypeNumber:
			if _gfgf || !_facga.IsBoolean {
				_dece += _facga.ValueNumber
				_bfbed++
			}
		case ResultTypeList, ResultTypeArray:
			_fffab, _afcaa := _ebbcf(_facga.ListValues(), _gfgf)
			_dece += _fffab
			_bfbed += _afcaa
		case ResultTypeString:
			if _gfgf {
				_bfbed++
			}
		case ResultTypeEmpty:
		}
	}
	return _dece, _bfbed
}
func _ag(_ca Result) bool {
	if _ca.Type == ResultTypeString {
		return _ca.ValueString == ""
	}
	return _ca.ValueNumber == 0
}

// Tbillprice implements the Excel TBILLPRICE function.
func Tbillprice(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("T\u0042\u0049\u004c\u004c\u0050\u0052I\u0043\u0045\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_bada, _dacf, _dafcf := _gfcg(args[0], args[1], "\u0054\u0042\u0049\u004c\u004c\u0050\u0052\u0049\u0043\u0045")
	if _dafcf.Type == ResultTypeError {
		return _dafcf
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0054\u0042\u0049\u004c\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0064\u0069\u0073\u0063\u006f\u0075n\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_eaed := _dacf - _bada
	if _eaed > 365 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004cP\u0052\u0049\u0043\u0045\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020m\u0061\u0074\u0075r\u0069\u0074\u0079\u0020t\u006f\u0020\u0062\u0065\u0020\u006eo\u0074\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u006e\u0065\u0020\u0079e\u0061\u0072\u0020\u0061\u0066\u0074\u0065\u0072\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074")
	}
	_gdec := args[2].ValueNumber
	if _gdec <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0054\u0042\u0049\u004c\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020d\u0069\u0073\u0063\u006f\u0075\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeNumberResult(100 * (1 - _gdec*_eaed/360))
}

// Eval evaluates the binary expression using the context given.
func (_cb BinaryExpr) String() string {
	_egf := ""
	switch _cb._fc {
	case BinOpTypePlus:
		_egf = "\u002b"
	case BinOpTypeMinus:
		_egf = "\u002d"
	case BinOpTypeMult:
		_egf = "\u002a"
	case BinOpTypeDiv:
		_egf = "\u002f"
	case BinOpTypeExp:
		_egf = "\u005e"
	case BinOpTypeLT:
		_egf = "\u003c"
	case BinOpTypeGT:
		_egf = "\u003e"
	case BinOpTypeEQ:
		_egf = "\u003d"
	case BinOpTypeLEQ:
		_egf = "\u003c\u003d"
	case BinOpTypeGEQ:
		_egf = "\u003e\u003d"
	case BinOpTypeNE:
		_egf = "\u003c\u003e"
	case BinOpTypeConcat:
		_egf = "\u0026"
	}
	return _cb._gb.String() + _egf + _cb._gba.String()
}

// Update updates references in the VerticalRange after removing a row/column.
func (_cafef VerticalRange) Update(q *_eg.UpdateQuery) Expression {
	if q.UpdateType == _eg.UpdateActionRemoveColumn {
		_ecbaf := _cafef
		if q.UpdateCurrentSheet {
			_bccf := q.ColumnIdx
			_ecbaf._ggcb = _deg(_cafef._ggcb, _bccf)
			_ecbaf._acdd = _deg(_cafef._acdd, _bccf)
		}
		return _ecbaf
	}
	return _cafef
}
func (_fgdcg *noCache) SetCache(key string, value Result) {}

// NewNumber constructs a new number expression.
func NewNumber(v string) Expression {
	_dfbaa, _gbgab := _afc.ParseFloat(v, 64)
	if _gbgab != nil {
		_cf.Log.Debug("e\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069\u006e\u0067\u0020f\u006f\u0072\u006d\u0075\u006c\u0061\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0025\u0073\u003a\u0020%\u0076", v, _gbgab)
	}
	return Number{_fbfe: _dfbaa}
}
func (_beeadb *ivr) Sheet(name string) Context { return _beeadb }
func _bca(_dfc, _bdgc, _gbee int) bool {
	if _bdgc < 1 || _bdgc > 12 {
		return false
	}
	if _gbee < 1 {
		return false
	}
	return _gbee <= _ffcg(_dfc, _bdgc)
}
func _beacf(_fgfed, _geeef Expression) (Expression, Expression, error) {
	_egbccb, _bbecgfd := _fgfed.(CellRef)
	if !_bbecgfd {
		return nil, nil, _c.New(_e.Sprintf("\u0049\u006e\u0063\u006frr\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020%\u0073", _fgfed.String()))
	}
	_eefdc, _bbecgfd := _geeef.(CellRef)
	if !_bbecgfd {
		return nil, nil, _c.New(_e.Sprintf("\u0049\u006e\u0063\u006frr\u0065\u0063\u0074\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020%\u0073", _geeef.String()))
	}
	_aafb, _adaf := _fb.ParseCellReference(_egbccb._ab)
	if _adaf != nil {
		return nil, nil, _adaf
	}
	_acdef, _bbbad := _fb.ParseCellReference(_eefdc._ab)
	if _bbbad != nil {
		return nil, nil, _bbbad
	}
	_cddee := false
	if _aafb.RowIdx > _acdef.RowIdx {
		_cddee = true
		_aafb.RowIdx, _acdef.RowIdx = _acdef.RowIdx, _aafb.RowIdx
	}
	if _aafb.ColumnIdx > _acdef.ColumnIdx {
		_cddee = true
		_aafb.ColumnIdx, _acdef.ColumnIdx = _acdef.ColumnIdx, _aafb.ColumnIdx
		_aafb.Column, _acdef.Column = _acdef.Column, _aafb.Column
	}
	if _cddee {
		return NewCellRef(_aafb.String()), NewCellRef(_acdef.String()), nil
	}
	return _fgfed, _geeef, nil
}

// YearFrac is an implementation of the Excel YEARFRAC() function.
func YearFrac(args []Result) Result {
	_dee := len(args)
	if (_dee != 2 && _dee != 3) || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber {
		return MakeErrorResult("Y\u0045\u0041\u0052\u0046\u0052\u0041\u0043\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020tw\u006f\u0020\u006f\u0072 \u0074\u0068\u0072\u0065\u0065\u0020\u006e\u0075\u006dbe\u0072\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_dbf := 0
	if _dee == 3 && args[2].Type != ResultTypeEmpty {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("Y\u0045\u0041\u0052\u0046\u0052\u0041\u0043\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020ba\u0073\u0069\u0073\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074o \u0062\u0065 \u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_dbf = int(args[2].ValueNumber)
		if !_gfe(_dbf) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0059\u0045\u0041R\u0046\u0052\u0041\u0043")
		}
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0045\u0041\u0052\u0046\u0052\u0041\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0074\u0061\u0072\u0074\u0020\u0064\u0061t\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cfaag := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0045\u0041\u0052\u0046\u0052\u0041\u0043 \u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u0065\u006ed\u0020\u0064\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_bce := args[1].ValueNumber
	_bfgg, _baaf := _fcfa(_cfaag, _bce, _dbf)
	if _baaf.Type == ResultTypeError {
		return _baaf
	}
	return MakeNumberResult(_bfgg)
}

// Ifs is an implementation of the Excel IFS() function.
func Ifs(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("I\u0046\u0053\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061t\u0020\u006c\u0065\u0061\u0073\u0074\u0020t\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	for _ebffd := 0; _ebffd < len(args)-1; _ebffd += 2 {
		if args[_ebffd].ValueNumber == 1 {
			return args[_ebffd+1]
		}
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}

// Int is an implementation of the Excel INT() function that rounds a number
// down to an integer.
func Int(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("I\u004e\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_cbda := args[0].AsNumber()
	if _cbda.Type != ResultTypeNumber {
		return MakeErrorResult("I\u004e\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_gbdgd, _acbbd := _cc.Modf(_cbda.ValueNumber)
	if _acbbd < 0 {
		_gbdgd--
	}
	return MakeNumberResult(_gbdgd)
}

// Now is an implementation of the Excel NOW() function.
func Now(args []Result) Result {
	if len(args) > 0 {
		return MakeErrorResult("\u004e\u004fW\u0020\u0064\u006f\u0065\u0073\u006e\u0027\u0074\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	_ded := _eb.Now()
	_, _adc := _ded.Zone()
	_dea := _acbd + float64(_ded.Unix()+int64(_adc))/86400
	return MakeNumberResult(_dea)
}

const _eeeg = 57359

func (_dcgba Result) String() string { return _dcgba.Value() }

// ISNONTEXT is an implementation of the Excel ISNONTEXT() function.
func IsNonText(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053N\u004f\u004e\u0054\u0045X\u0054\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073 \u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	return MakeBoolResult(args[0].Type != ResultTypeString)
}

// Effect implements the Excel EFFECT function.
func Effect(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0045\u0046F\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0045\u0046\u0046\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u006f\u006d\u0069n\u0061\u006c\u0020\u0069\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020a\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_afbd := args[0].ValueNumber
	if _afbd <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0045\u0046\u0046\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u006f\u006d\u0069n\u0061\u006c\u0020\u0069\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0072\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062e\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0045\u0046\u0046\u0045\u0043\u0054 \u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u006f\u0066 \u0063\u006f\u006d\u0070\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020p\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074")
	}
	_dbef := float64(int(args[1].ValueNumber))
	if _dbef < 1 {
		return MakeErrorResultType(ErrorTypeNum, "E\u0046\u0046\u0045\u0043\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0063o\u006dp\u006f\u0075\u006e\u0064i\u006e\u0067 \u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0031\u0020\u006f\u0072\u0020\u006d\u006f\u0072\u0065")
	}
	return MakeNumberResult(_cc.Pow((1+_afbd/_dbef), _dbef) - 1)
}

// String returns a string representation of a range with prefix.
func (_egece PrefixRangeExpr) String() string {
	return _e.Sprintf("\u0025\u0073\u0021\u0025\u0073\u003a\u0025\u0073", _egece._beee.String(), _egece._bbbbge.String(), _egece._affc.String())
}

// ISERROR is an implementation of the Excel ISERROR() function.
func IsError(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("I\u0053\u0045\u0052\u0052\u004f\u0052(\u0029\u0020\u0061\u0063\u0063\u0065p\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeError)
}

// NewRange constructs a new range.
func NewRange(from, to Expression) Expression {
	_gdgda, _dbacf, _aeade := _beacf(from, to)
	if _aeade != nil {
		_cf.Log.Debug(_aeade.Error())
		return NewError(_aeade.Error())
	}
	return Range{_fbdb: _gdgda, _agedb: _dbacf}
}

// Today is an implementation of the Excel TODAY() function.
func Today(args []Result) Result {
	if len(args) > 0 {
		return MakeErrorResult("\u0054\u004f\u0044A\u0059\u0020\u0064\u006fe\u0073\u006e\u0027\u0074\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ccc := _eb.Now()
	_, _eebe := _ccc.Zone()
	_bbge := _bddd(_afcc, _ccc.Unix()+int64(_eebe)) + 1
	return MakeNumberResult(_bbge)
}

// String returns a string representation of a vertical range with prefix.
func (_gefag PrefixVerticalRange) String() string {
	return _e.Sprintf("\u0025\u0073\u0021\u0025\u0073\u003a\u0025\u0073", _gefag._fcga.String(), _gefag._agfae, _gefag._fefae)
}

// ISTEXT is an implementation of the Excel ISTEXT() function.
func IsText(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u0054EX\u0054\u0028\u0029\u0020\u0061\u0063\u0063\u0065\u0070t\u0073 \u0061 \u0073i\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeString)
}

// Result is the result of a formula or cell evaluation .
type Result struct {
	ValueNumber  float64
	ValueString  string
	ValueList    []Result
	ValueArray   [][]Result
	IsBoolean    bool
	ErrorMessage string
	Type         ResultType
	Ref          Reference
}

var _bgfb []byte = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// VLookup implements the VLOOKUP function that returns a matching value from a
// column in an array.
func VLookup(args []Result) Result {
	_ecgcg := len(args)
	if _ecgcg < 3 {
		return MakeErrorResult("\u0056\u004c\u004f\u004f\u004bU\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _ecgcg > 4 {
		return MakeErrorResult("\u0056\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0074\u0020m\u006f\u0073\u0074\u0020\u0066\u006f\u0075\u0072\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_fagd := args[0]
	_gacb := args[1]
	if _gacb.Type != ResultTypeArray {
		return MakeErrorResult("\u0056\u004cO\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_ccea := args[2].AsNumber()
	if _ccea.Type != ResultTypeNumber {
		return MakeErrorResult("\u0056\u004cO\u004f\u004b\u0055\u0050 \u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075m\u0065\u0072\u0069\u0063\u0020\u0063\u006f\u006c\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_fecg := false
	if _ecgcg == 4 && args[3].Type != ResultTypeEmpty {
		_fggfe := args[3].AsNumber()
		if _fggfe.Type != ResultTypeNumber {
			return MakeErrorResult("\u0056\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		if _fggfe.ValueNumber == 0 {
			_fecg = true
		}
	}
	_afbde := int(_ccea.ValueNumber) - 1
	_gfcgc := -1
	_ccffa := false
_beea:
	for _add, _fbbe := range _gacb.ValueArray {
		if len(_fbbe) == 0 {
			continue
		}
		_ccbb := _fbbe[0]
		switch _edeef(_ccbb, _fagd, false, _fecg) {
		case _gcefb:
			_gfcgc = _add
		case _cgae:
			_gfcgc = _add
			_ccffa = true
			break _beea
		}
	}
	if _gfcgc == -1 {
		return MakeErrorResultType(ErrorTypeNA, "\u0056\u004c\u004fOK\u0055\u0050\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_aaada := _gacb.ValueArray[_gfcgc]
	if _afbde < 0 || _afbde >= len(_aaada) {
		return MakeErrorResult("\u0056\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0068\u0061\u0073\u0020\u0069\u006e\u0076a\u006ci\u0064\u0020\u0063\u006f\u006c\u0075\u006d\u006e\u0020\u0069\u006e\u0064\u0065\u0078")
	}
	if _ccffa || !_fecg {
		return _aaada[_afbde]
	}
	return MakeErrorResultType(ErrorTypeNA, "\u0056\u004c\u004fOK\u0055\u0050\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

// Eval evaluates a range with prefix returning a list of results or an error.
func (_edge PrefixRangeExpr) Eval(ctx Context, ev Evaluator) Result {
	_ecedg := _edge._beee.Reference(ctx, ev)
	_becde := _edge._bbbbge.Reference(ctx, ev)
	_dcda := _edge._affc.Reference(ctx, ev)
	switch _ecedg.Type {
	case ReferenceTypeSheet:
		if _dcbb(_ecedg, ctx) {
			return MakeErrorResultType(ErrorTypeName, _e.Sprintf("\u0053h\u0065e\u0074\u0020\u0025\u0073\u0020n\u006f\u0074 \u0066\u006f\u0075\u006e\u0064", _ecedg.Value))
		}
		_ebdc := _cdbdg(_ecedg, _becde, _dcda)
		if _becde.Type == ReferenceTypeCell && _dcda.Type == ReferenceTypeCell {
			if _cceb, _ggece := ev.GetFromCache(_ebdc); _ggece {
				return _cceb
			} else {
				_gbfbc := _gebc(ctx.Sheet(_ecedg.Value), ev, _becde.Value, _dcda.Value)
				ev.SetCache(_ebdc, _gbfbc)
				return _gbfbc
			}
		}
		return MakeErrorResult("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072a\u006e\u0067\u0065\u0020" + _ebdc)
	default:
		return MakeErrorResult(_e.Sprintf("\u006e\u006f\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0020\u0066\u006f\u0072\u0020r\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _ecedg.Type))
	}
}

// Eval evaluates and returns a string.
func (_ecba String) Eval(ctx Context, ev Evaluator) Result { return MakeStringResult(_ecba._bedb) }

type xargs struct {
	_faaf []float64
	_ceee []float64
}

func (_ebafa *Lexer) Next() *node {
	_ebafa._gegda.Lock()
	defer _ebafa._gegda.Unlock()
	if len(_ebafa._cdcgc) > 0 {
		_ebeb := _ebafa._cdcgc[0]
		_ebafa._cdcgc = _ebafa._cdcgc[1:]
		return _ebeb
	}
	return _ebafa.nextRaw()
}

// NewError constructs a new error expression from a string.
func NewError(v string) Expression { return Error{_gbb: v} }
func _fedaf(_dfag Result, _ggd *criteriaParsed) bool {
	if _dfag.Type == ResultTypeEmpty {
		return false
	}
	if _ggd._ceeeb {
		return _dfag.ValueNumber == _ggd._adag
	} else {
		_gegb := _ga.ToLower(_dfag.ValueString)
		return _ggd._abgb == _gegb || _cd.Match(_ggd._abgb, _gegb)
	}
}

// Coupdaybs implements the Excel COUPDAYBS function.
func Coupdaybs(args []Result) Result {
	_fcab, _dcfca := _cebf(args, "\u0043O\u0055\u0050\u0044\u0041\u0059\u0042S")
	if _dcfca.Type == ResultTypeError {
		return _dcfca
	}
	return MakeNumberResult(_fbebf(_fcab._bbaa, _fcab._aee, _fcab._gbf, _fcab._ceea))
}

// Pmt implements the Excel PMT function.
func Pmt(args []Result) Result {
	_fgcf := len(args)
	if _fgcf < 3 || _fgcf > 5 {
		return MakeErrorResult("\u0050\u004dT\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u006f\u0066\u0020\u0033\u0020\u0061\u006e\u0064\u0020\u0035")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020r\u0061\u0074\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_egbd := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u004dT\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_cggf := args[1].ValueNumber
	if _cggf == 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u004d\u0054\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u0074\u0020\u0065\u0071\u0075\u0061\u006c\u0020\u0074\u006f\u00200")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073 p\u0072\u0065\u0073\u0065\u006e\u0074 \u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_dadb := args[2].ValueNumber
	_ggacc := 0.0
	if _fgcf >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("P\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0075\u0074\u0075\u0072e \u0076\u0061\u006c\u0075e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075mb\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_ggacc = args[3].ValueNumber
	}
	_cgbe := 0.0
	if _fgcf == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050\u004d\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020t\u0079\u0070\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_cgbe = args[4].ValueNumber
		if _cgbe != 0 {
			_cgbe = 1
		}
	}
	var _dbfe float64
	if _egbd == 0 {
		_dbfe = (_dadb + _ggacc) / _cggf
	} else {
		_gbfa := _cc.Pow(1+_egbd, _cggf)
		if _cgbe == 1 {
			_dbfe = (_ggacc*_egbd/(_gbfa-1) + _dadb*_egbd/(1-1/_gbfa)) / (1 + _egbd)
		} else {
			_dbfe = _ggacc*_egbd/(_gbfa-1) + _dadb*_egbd/(1-1/_gbfa)
		}
	}
	return MakeNumberResult(-_dbfe)
}

// Pi is an implementation of the Excel Pi() function that just returns the Pi
// constant.
func Pi(args []Result) Result {
	if len(args) != 0 {
		return MakeErrorResult("\u0050I\u0028\u0029\u0020\u0061c\u0063\u0065\u0070\u0074\u0073 \u006eo\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074s")
	}
	return MakeNumberResult(_cc.Pi)
}

// Max is an implementation of the Excel MAX() function.
func Max(args []Result) Result { return _bceb(args, false) }

// SeriesSum implements the Excel SERIESSUM function.
func SeriesSum(args []Result) Result {
	if len(args) != 4 {
		return MakeErrorResult("\u0053\u0045\u0052\u0049\u0045\u0053\u0053\u0055\u004d\u0028\u0029\u0020\u0072\u0065\u0071u\u0069r\u0065\u0073\u0020\u0034\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_abcg := args[0].AsNumber()
	_fffa := args[1].AsNumber()
	_ggfe := args[2].AsNumber()
	_ffgg := args[3].ListValues()
	if _abcg.Type != ResultTypeNumber || _fffa.Type != ResultTypeNumber || _ggfe.Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u0045\u0052\u0049\u0045\u0053S\u0055\u004d\u0028)\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073t\u0020\u0074\u0068\u0072\u0065e \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063")
	}
	_ffge := float64(0)
	for _dcccg, _dbab := range _ffgg {
		_ffge += _dbab.ValueNumber * _cc.Pow(_abcg.ValueNumber, _fffa.ValueNumber+float64(_dcccg)*_ggfe.ValueNumber)
	}
	return MakeNumberResult(_ffge)
}

// Error is called in the case of parsing error and saves an error to a plex.
func (_eefd *plex) Error(s string) {
	_cf.Log.Debug("\u0070a\u0072s\u0065\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0073", s)
	_eefd._ddcd = s
}

// Eval evaluates and returns the result of the NamedRangeRef reference.
func (_bbdgf NamedRangeRef) Eval(ctx Context, ev Evaluator) Result {
	_cegcd := ctx.NamedRange(_bbdgf._ceaf)
	_bbadb := _cegcd.Value
	if _aged, _cfgb := ev.GetFromCache(_bbadb); _cfgb {
		return _aged
	}
	_afdcd := _ga.Split(_bbadb, "\u0021")
	if len(_afdcd) != 2 {
		return MakeErrorResult(_e.Sprintf("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006e\u0061\u006de\u0064 \u0072\u0061\u006e\u0067\u0065\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0025\u0073", _bbadb))
	}
	_afdec := ctx.Sheet(_afdcd[0])
	_gafbd := _ga.Split(_afdcd[1], "\u003a")
	switch len(_gafbd) {
	case 1:
		_ebgdd := ev.Eval(_afdec, _gafbd[0])
		ev.SetCache(_bbadb, _ebgdd)
		return _ebgdd
	case 2:
		_bbcfc := _gebc(_afdec, ev, _gafbd[0], _gafbd[1])
		ev.SetCache(_bbadb, _bbcfc)
		return _bbcfc
	}
	return MakeErrorResult(_e.Sprintf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u0079\u0070e \u0025\u0073", _cegcd.Type))
}
func _ddd(_faf string) (int, int, int, bool, Result) {
	_gea := ""
	_fac := []string{}
	for _dded, _cea := range _fce {
		_fac = _cea.FindStringSubmatch(_faf)
		if len(_fac) > 1 {
			_gea = _dded
			break
		}
	}
	if _gea == "" {
		return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
	}
	_fbgf := false
	var _cceg, _fcb, _abb int
	var _edf error
	switch _gea {
	case "\u006d\u006d\u002f\u0064\u0064\u002f\u0079\u0079":
		_fcb, _edf = _afc.Atoi(_fac[1])
		if _edf != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_abb, _edf = _afc.Atoi(_fac[3])
		if _edf != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_cceg, _edf = _afc.Atoi(_fac[5])
		if _edf != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		if _cceg < 0 || _cceg > 9999 || (_cceg > 99 && _cceg < 1900) {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_cceg = _edg(_cceg)
		_fbgf = _fac[8] == ""
	case "\u006dm\u0020\u0064\u0064\u002c\u0020\u0079y":
		_fcb = _fae[_fac[1]]
		_abb, _edf = _afc.Atoi(_fac[14])
		if _edf != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_cceg, _edf = _afc.Atoi(_fac[16])
		if _edf != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		if _cceg < 0 || _cceg > 9999 || (_cceg > 99 && _cceg < 1900) {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_cceg = _edg(_cceg)
		_fbgf = _fac[19] == ""
	case "\u0079\u0079\u002d\u006d\u006d\u002d\u0064\u0064":
		_eeac, _bdc := _afc.Atoi(_fac[1])
		if _bdc != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_babb, _bdc := _afc.Atoi(_fac[3])
		if _bdc != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_ead, _bdc := _afc.Atoi(_fac[5])
		if _bdc != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		if _eeac >= 1900 && _eeac < 10000 {
			_cceg = _eeac
			_fcb = _babb
			_abb = _ead
		} else if _eeac > 0 && _eeac < 13 {
			_fcb = _eeac
			_abb = _babb
			_cceg = _ead
		} else {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_fbgf = _fac[8] == ""
	case "y\u0079\u002d\u006d\u006d\u0053\u0074\u0072\u002d\u0064\u0064":
		_cceg, _edf = _afc.Atoi(_fac[16])
		if _edf != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_fcb = _fae[_fac[3]]
		_abb, _edf = _afc.Atoi(_fac[1])
		if _edf != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
		}
		_fbgf = _fac[19] == ""
	}
	if !_bca(_cceg, _fcb, _abb) {
		return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, _gbe)
	}
	return _cceg, _fcb, _abb, _fbgf, _feb
}
func _cddc(_ged string) bool {
	for _, _fcag := range _afd {
		_gcg := _fcag.FindStringSubmatch(_ged)
		if len(_gcg) > 1 {
			return true
		}
	}
	return false
}

const _egfcf = 57360

func _dec(_caefe, _ggbc []float64, _dbcg float64) Result {
	_aaeg := false
	_bef := false
	for _aggg := 0; _aggg < len(_caefe); _aggg++ {
		if _caefe[_aggg] > 0 {
			_aaeg = true
		}
		if _caefe[_aggg] < 0 {
			_bef = true
		}
	}
	if !_aaeg || !_bef {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
	_ccfdf := _dbcg
	_ggge := 1e-10
	_gbec := 0
	_fadea := 50
	_cdca := false
	for {
		_dbaf := _bfbe(_caefe, _ggbc, _ccfdf)
		_dfg := _ccfdf - _dbaf/_feac(_caefe, _ggbc, _ccfdf)
		_caec := _cc.Abs(_dfg - _ccfdf)
		_ccfdf = _dfg
		_gbec++
		if _caec <= _ggge || _cc.Abs(_dbaf) <= _ggge {
			break
		}
		if _gbec > _fadea {
			_cdca = true
			break
		}
	}
	if _cdca || _cc.IsNaN(_ccfdf) || _cc.IsInf(_ccfdf, 0) {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
	return MakeNumberResult(_ccfdf)
}

const _gdac = 57361

var _aedg = []ri{{1000, "\u004d"}, {999, "\u0049\u004d"}, {995, "\u0056\u004d"}, {990, "\u0058\u004d"}, {950, "\u004c\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {499, "\u0049\u0044"}, {495, "\u0056\u0044"}, {490, "\u0058\u0044"}, {450, "\u004c\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {99, "\u0049\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {45, "\u0056\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}

type defEval struct {
	evCache
	_fdc bool
}

func _gbde(_faceb Result) Result {
	if _faceb.Type == ResultTypeEmpty {
		return _faceb
	}
	_eaad := _faceb.AsString()
	if _eaad.Type != ResultTypeString {
		return MakeErrorResult("\u004c\u004f\u0057\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0061\u0020\u0073\u0069\u006eg\u006c\u0065\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	if _faceb.IsBoolean {
		if _eaad.ValueString == "\u0031" {
			return MakeStringResult("\u0074\u0072\u0075\u0065")
		} else if _eaad.ValueString == "\u0030" {
			return MakeStringResult("\u0066\u0061\u006cs\u0065")
		} else {
			return MakeErrorResult("\u0049\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u004c\u004fW\u0045\u0052")
		}
	} else {
		return MakeStringResult(_ga.ToLower(_eaad.ValueString))
	}
}

// ListValues converts an array to a list or returns a lists values. This is used
// for functions that can accept an array, but don't care about ordering to
// reuse the list function logic.
func (_eaaac Result) ListValues() []Result {
	if _eaaac.Type == ResultTypeArray {
		_febd := []Result{}
		for _, _aggc := range _eaaac.ValueArray {
			for _, _aeead := range _aggc {
				_febd = append(_febd, _aeead)
			}
		}
		return _febd
	}
	if _eaaac.Type == ResultTypeList {
		return _eaaac.ValueList
	}
	return nil
}
func _bee(_deb int) bool {
	if _deb == _deb/400*400 {
		return true
	}
	if _deb == _deb/100*100 {
		return false
	}
	return _deb == _deb/4*4
}
func _bba(_ceff, _bafe, _fgac int) int {
	if _ceff > _bafe {
		return 0
	}
	if _gbgg(_fgac) {
		return (_bafe - _ceff + 1) * 360
	}
	_adf := 0
	for _eebd := _ceff; _eebd <= _bafe; _eebd++ {
		_debc := 365
		if _bee(_eebd) {
			_debc = 366
		}
		_adf += _debc
	}
	return _adf
}

type cmpResult int8

// Quotient is an implementation of the Excel QUOTIENT function that returns the
// integer portion of division.
func Quotient(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0051\u0055\u004f\u0054\u0049E\u004e\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0074\u0077\u006f\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_adbaf := args[0].AsNumber()
	_aabfb := args[1].AsNumber()
	if _adbaf.Type != ResultTypeNumber || _aabfb.Type != ResultTypeNumber {
		return MakeErrorResult("\u0051\u0055\u004f\u0054\u0049E\u004e\u0054\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u0074\u0077\u006f\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _aabfb.ValueNumber == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0051U\u004f\u0054\u0049\u0045N\u0054\u0028\u0029\u0020\u0064i\u0076i\u0064e\u0020\u0062\u0079\u0020\u007a\u0065\u0072o")
	}
	return MakeNumberResult(_cc.Trunc(_adbaf.ValueNumber / _aabfb.ValueNumber))
}

var _dfge string = string([]byte{92})

func _de(_db BinOpType, _ea [][]Result, _fdb Result) Result {
	_bf := [][]Result{}
	for _da := range _ea {
		_ff := _fgfe(_db, _ea[_da], _fdb)
		if _ff.Type == ResultTypeError {
			return _ff
		}
		_bf = append(_bf, _ff.ValueList)
	}
	return MakeArrayResult(_bf)
}

type couponArgs struct {
	_bbaa float64
	_aee  float64
	_gbf  int
	_ceea int
}

// IsLogical is an implementation of the Excel ISLOGICAL() function.
func IsLogical(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053\u004c\u004f\u0047\u0049\u0043A\u004c\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fcac := args[0].Ref
	if _fcac.Type != ReferenceTypeCell {
		return MakeErrorResult("I\u0053\u004c\u004f\u0047\u0049\u0043\u0041\u004c\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u0072\u0073t\u0020a\u0072\u0067\u0075\u006de\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
	}
	return MakeBoolResult(ctx.Cell(_fcac.Value, ev).IsBoolean)
}

// MaxIfs implements the MAXIFS function.
func MaxIfs(args []Result) Result {
	_decb := _eged(args, true, "\u004d\u0041\u0058\u0049\u0046\u0053")
	if _decb.Type != ResultTypeEmpty {
		return _decb
	}
	_bafb := _ecbg(args[1:])
	_egfba := -_cc.MaxFloat64
	_eceb := _aadag(args[0])
	for _, _abfe := range _bafb {
		_eaebd := _eceb[_abfe._aaba][_abfe._gbfb].ValueNumber
		if _egfba < _eaebd {
			_egfba = _eaebd
		}
	}
	if _egfba == -_cc.MaxFloat64 {
		_egfba = 0
	}
	return MakeNumberResult(float64(_egfba))
}

// Roman is an implementation of the Excel ROMAN function that convers numbers
// to roman numerals in one of 5 formats.
func Roman(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u0052\u004fM\u0041\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("\u0052\u004fM\u0041\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006d\u006f\u0073\u0074\u0020\u0074\u0077\u006f\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_cceec := args[0].AsNumber()
	if _cceec.Type != ResultTypeNumber {
		return MakeErrorResult("\u0052\u004fM\u0041\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006e\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gbgb := 0
	if len(args) > 1 {
		_bgfa := args[1]
		if _bgfa.Type != ResultTypeNumber {
			return MakeErrorResult("\u0052\u004fM\u0041\u004e\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0065\u0072\u0069\u0063")
		}
		_gbgb = int(_bgfa.ValueNumber)
		if _gbgb < 0 {
			_gbgb = 0
		} else if _gbgb > 4 {
			_gbgb = 4
		}
	}
	_gbcf := _gcfd
	switch _gbgb {
	case 1:
		_gbcf = _bdebe
	case 2:
		_gbcf = _bdcac
	case 3:
		_gbcf = _eaag
	case 4:
		_gbcf = _aedg
	}
	_bgdef := _cc.Trunc(_cceec.ValueNumber)
	_dagea := _bd.Buffer{}
	for _, _gbcc := range _gbcf {
		for _bgdef >= _gbcc._faec {
			_dagea.WriteString(_gbcc._ddcg)
			_bgdef -= _gbcc._faec
		}
	}
	return MakeStringResult(_dagea.String())
}

// T is an implementation of the Excel T function that returns whether the
// argument is text.
func T(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("T\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0020\u0073i\u006e\u0067\u006c\u0065\u0020\u0073\u0074r\u0069\u006e\u0067\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_ddgd := args[0]
	if _ddgd.Type == ResultTypeError || _ddgd.Type == ResultTypeString {
		return _ddgd
	}
	return _feb
}

type ri struct {
	_faec float64
	_ddcg string
}

// Ceiling is an implementation of the CEILING function which
// returns the ceiling of a number.
func Ceiling(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("C\u0045\u0049\u004c\u0049\u004e\u0047\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006f\u006ee \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	if len(args) > 2 {
		return MakeErrorResult("\u0043\u0045\u0049\u004c\u0049\u004e\u0047\u0028\u0029\u0020\u0061\u006c\u006c\u006f\u0077\u0073\u0020\u0061\u0074\u0020\u006d\u006f\u0073\u0074 \u0074\u0077\u006f\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_efgad := args[0].AsNumber()
	if _efgad.Type != ResultTypeNumber {
		return MakeErrorResult("\u0066i\u0072\u0073t\u0020\u0061\u0072\u0067u\u006d\u0065\u006et\u0020\u0074\u006f\u0020\u0043\u0045\u0049\u004c\u0049NG\u0028\u0029\u0020m\u0075\u0073t\u0020\u0062\u0065\u0020\u0061\u0020n\u0075\u006db\u0065\u0072")
	}
	_geffg := float64(1)
	if _efgad.ValueNumber < 0 {
		_geffg = -1
	}
	if len(args) > 1 {
		_aec := args[1].AsNumber()
		if _aec.Type != ResultTypeNumber {
			return MakeErrorResult("\u0073e\u0063\u006fn\u0064\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020t\u006f\u0020\u0043\u0045\u0049\u004cI\u004e\u0047\u0028\u0029\u0020\u006d\u0075\u0073\u0074\u0020\u0062e\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
		}
		_geffg = _aec.ValueNumber
	}
	if _geffg < 0 && _efgad.ValueNumber > 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u006e\u0065\u0067\u0061\u0074\u0069v\u0065\u0020\u0073\u0069\u0067\u0020\u0074\u006f\u0020\u0043\u0045\u0049\u004cI\u004e\u0047\u0028\u0029\u0020\u0069\u006ev\u0061\u006c\u0069\u0064")
	}
	if len(args) == 1 {
		return MakeNumberResult(_cc.Ceil(_efgad.ValueNumber))
	}
	_bagg := _efgad.ValueNumber
	_bagg, _bbef := _cc.Modf(_bagg / _geffg)
	if _bbef > 0 {
		_bagg++
	}
	return MakeNumberResult(_bagg * _geffg)
}

var InvalidReferenceContext = &ivr{}

// Update updates references in the PrefixVerticalRange after removing a row/column.
func (_adcbg PrefixVerticalRange) Update(q *_eg.UpdateQuery) Expression {
	if q.UpdateType == _eg.UpdateActionRemoveColumn {
		_aaadc := _adcbg
		_gaebaa := _adcbg._fcga.String()
		if _gaebaa == q.SheetToUpdate {
			_abac := q.ColumnIdx
			_aaadc._agfae = _deg(_adcbg._agfae, _abac)
			_aaadc._fefae = _deg(_adcbg._fefae, _abac)
		}
		return _aaadc
	}
	return _adcbg
}

const _ecbdf = 57351

// Oddlyield implements the Excel ODDLYIELD function.
func Oddlyield(args []Result) Result {
	if len(args) != 7 && len(args) != 8 {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0059\u0049\u0045L\u0044\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0073\u0065\u0076\u0065\u006e\u0020\u006f\u0072\u0020\u0065\u0069\u0067\u0068\u0074\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ddbeg, _acdb, _aggb := _gfcg(args[0], args[1], "\u004fD\u0044\u004c\u0059\u0049\u0045\u004cD")
	if _aggb.Type == ResultTypeError {
		return _aggb
	}
	_acce, _aggb := _eec(args[2], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u004fD\u0044\u004c\u0050\u0052\u0049\u0043E")
	if _aggb.Type == ResultTypeError {
		return _aggb
	}
	if _acce >= _ddbeg {
		return MakeErrorResultType(ErrorTypeNum, "\u004c\u0061\u0073\u0074\u0020i\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0064\u0061\u0074\u0065\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0062\u0065\u0066\u006f\u0072\u0065\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074e")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020o\u0066\u0020\u0074\u0079\u0070e\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_cddbb := args[3].ValueNumber
	if _cddbb < 0 {
		return MakeErrorResultType(ErrorTypeNum, "R\u0061\u0074\u0065\u0020\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006e\u006fn\u0020\u006e\u0065g\u0061t\u0069\u0076\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("O\u0044\u0044\u004c\u0059\u0049\u0045\u004c\u0044\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 p\u0072\u0065\u0073\u0065n\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006ff \u0074\u0079p\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_eff := args[4].ValueNumber
	if _eff <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0073h\u006fu\u006c\u0064\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065")
	}
	if args[5].Type != ResultTypeNumber {
		return MakeErrorResult("\u004fD\u0044\u004cY\u0049\u0045\u004c\u0044 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065mp\u0074\u0069\u006fn\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020n\u0075\u006db\u0065\u0072")
	}
	_daa := args[5].ValueNumber
	if _daa < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006cd\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u006eo\u006e\u0020\u006e\u0065\u0067\u0061\u0074i\u0076\u0065")
	}
	if args[6].Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0059\u0049\u0045L\u0044\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0072\u0065\u0071\u0075\u0065\u006e\u0063\u0079\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_fefc := float64(int(args[6].ValueNumber))
	if !_dddg(_fefc) {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_fgeg := 0
	if len(args) == 8 && args[7].Type != ResultTypeEmpty {
		if args[7].Type != ResultTypeNumber {
			return MakeErrorResult("\u004f\u0044\u0044\u004c\u0059\u0049\u0045\u004c\u0044\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0062a\u0073\u0069\u0073\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
		}
		_fgeg = int(args[7].ValueNumber)
		if !_gfe(_fgeg) {
			return MakeErrorResultType(ErrorTypeNum, "I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0062\u0061\u0073\u0069s\u0020\u0076\u0061\u006c\u0075\u0065\u0020f\u006f\u0072\u0020\u004f\u0044\u0044\u004c\u0059\u0049\u0045L\u0044")
		}
	}
	_gafd, _aggb := _fcfa(_acce, _acdb, _fgeg)
	if _aggb.Type == ResultTypeError {
		return _aggb
	}
	_gafd *= _fefc
	_geadf, _aggb := _fcfa(_ddbeg, _acdb, _fgeg)
	if _aggb.Type == ResultTypeError {
		return _aggb
	}
	_geadf *= _fefc
	_aeefc, _aggb := _fcfa(_acce, _ddbeg, _fgeg)
	if _aggb.Type == ResultTypeError {
		return _aggb
	}
	_aeefc *= _fefc
	_adcf := _daa + _gafd*100*_cddbb/_fefc
	_adcf /= _eff + _aeefc*100*_cddbb/_fefc
	_adcf--
	_adcf *= _fefc / _geadf
	return MakeNumberResult(_adcf)
}
func _bfeef(_bcaea Result, _fdgc, _ggdef string) (float64, Result) {
	switch _bcaea.Type {
	case ResultTypeEmpty:
		return 0, _feb
	case ResultTypeNumber:
		return _bcaea.ValueNumber, _feb
	case ResultTypeString:
		_dbee, _eaca := _afc.ParseFloat(_bcaea.ValueString, 64)
		if _eaca != nil {
			return 0, MakeErrorResult(_ggdef + "\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0066\u006f\u0072\u0020" + _fdgc)
		}
		return _dbee, _feb
	default:
		return 0, MakeErrorResult(_fdgc + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020" + _ggdef + "\u0020t\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0065\u006d\u0070\u0074\u0079")
	}
}

// Range is a range expression that when evaluated returns a list of Results.
type Range struct{ _fbdb, _agedb Expression }

// Pricemat implements the Excel PRICEMAT function.
func Pricemat(args []Result) Result {
	_abde := len(args)
	if _abde != 5 && _abde != 6 {
		return MakeErrorResult("\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0066\u0069v\u0065\u0020\u006f\u0072\u0020\u0073\u0069\u0078\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_afa, _ccbd, _eedg := _gfcg(args[0], args[1], "\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054")
	if _eedg.Type == ResultTypeError {
		return _eedg
	}
	_bdef, _eedg := _eec(args[2], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054")
	if _eedg.Type == ResultTypeError {
		return _eedg
	}
	if _bdef >= _afa {
		return MakeErrorResult("\u0050\u0052\u0049\u0043E\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065\u0020\u0074\u006f\u0020\u0062e\u0020\u0062\u0065\u0066\u006fr\u0065\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052I\u0043\u0045\u004d\u0041T\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072a\u0074\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_bafde := args[3].ValueNumber
	if _bafde < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049\u0043\u0045M\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072a\u0074\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052\u0049\u0043\u0045\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0079\u0069\u0065\u006c\u0064\u0020o\u0066\u0020\u0074\u0079\u0070e\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_ade := args[4].ValueNumber
	if _ade < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0050\u0052\u0049C\u0045\u004d\u0041\u0054\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0079\u0069\u0065\u006c\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065")
	}
	_cgaab := 0
	if _abde == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("\u0050R\u0049\u0043E\u004d\u0041\u0054 \u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020\u0062\u0061\u0073\u0069\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_cgaab = int(args[5].ValueNumber)
		if !_gfe(_cgaab) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065c\u0074\u0020b\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074\u0020\u0066\u006f\u0072\u0020\u0050\u0052\u0049C\u0045\u004d\u0041\u0054")
		}
	}
	_daac, _eedg := _fcfa(_afa, _ccbd, _cgaab)
	if _eedg.Type == ResultTypeError {
		return _eedg
	}
	_fbba, _eedg := _fcfa(_bdef, _ccbd, _cgaab)
	if _eedg.Type == ResultTypeError {
		return _eedg
	}
	_eebb, _eedg := _fcfa(_bdef, _afa, _cgaab)
	if _eedg.Type == ResultTypeError {
		return _eedg
	}
	_fadg := 1 + _fbba*_bafde
	_gbff := 1 + _daac*_ade
	return MakeNumberResult((_fadg/_gbff - _eebb*_bafde) * 100)
}

// MakeArrayResult constructs an array result (matrix).
func MakeArrayResult(arr [][]Result) Result { return Result{Type: ResultTypeArray, ValueArray: arr} }

// Eomonth is an implementation of the Excel EOMONTH() function.
func Eomonth(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0045\u004f\u004d\u004f\u004e\u0054\u0048\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074w\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0045\u004f\u004d\u004f\u004e\u0054\u0048")
	}
	_efb := args[1].ValueNumber
	_adge := args[0]
	var _dac float64
	switch _adge.Type {
	case ResultTypeEmpty:
		_dac = 0
	case ResultTypeNumber:
		_dac = _adge.ValueNumber
	case ResultTypeString:
		_aaca := DateValue([]Result{args[0]})
		if _aaca.Type == ResultTypeError {
			return MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0045\u004f\u004d\u004f\u004e\u0054\u0048")
		}
		_dac = _aaca.ValueNumber
	default:
		return MakeErrorResult("\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0045\u004f\u004d\u004f\u004e\u0054\u0048")
	}
	_cefa := _bgee(_dac)
	_aaa := _cefa.AddDate(0, int(_efb+1), 0)
	_afcg, _dgc, _ := _aaa.Date()
	_dbd := _bad(_afcg, int(_dgc), 0)
	if _dbd < 1 {
		return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0045\u004f\u004d\u004f\u004e\u0054\u0048")
	}
	if _afcg == 1900 && _dgc == 3 {
		_dbd--
	}
	return MakeNumberResult(_dbd)
}

const _fed = "\u0028\u0028\u005b\u0030\u002d\u0039]\u0029\u002b\u0029:\u0028\u0028\u005b0\u002d\u0039\u005d\u0029\u002b\u0029\u003a\u0028\u0028\u005b0\u002d\u0039\u005d\u0029\u002b(\\\u002e\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u003f\u0029\u0028\u0020\u0028\u0061\u006d\u007c\u0070\u006d\u0029\u0029\u003f"

// Decimal is an implementation of the Excel function DECIMAL() that parses a string
// in a given base and returns the numeric result.
func Decimal(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u0044\u0045\u0043\u0049\u004d\u0041\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069r\u0065s\u0020\u0074\u0077\u006f\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ecgca := args[0].AsString()
	if _ecgca.Type != ResultTypeString {
		return MakeErrorResult("D\u0045\u0043\u0049\u004d\u0041\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0066\u0069\u0072\u0073t \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_dabg := args[1].AsNumber()
	if _dabg.Type != ResultTypeNumber {
		return MakeErrorResult("\u0044\u0045\u0043\u0049\u004dA\u004c\u0028\u0029\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gdca := _ecgca.ValueString
	if len(_gdca) > 2 && (_ga.HasPrefix(_gdca, "\u0030\u0078") || _ga.HasPrefix(_gdca, "\u0030\u0058")) {
		_gdca = _gdca[2:]
	}
	_gcdc, _fgcbf := _afc.ParseInt(_gdca, int(_dabg.ValueNumber), 64)
	if _fgcbf != nil {
		return MakeErrorResult("\u0044\u0045C\u0049\u004d\u0041\u004c\u0028\u0029\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0069\u006e\u0020\u0063\u006f\u006e\u0076\u0065\u0072si\u006f\u006e")
	}
	return MakeNumberResult(float64(_gcdc))
}

// VerticalRange is a range expression that when evaluated returns a list of Results from references like AA:IJ (all cells from columns AA to IJ).
type VerticalRange struct{ _ggcb, _acdd string }

// Parse parses an io.Reader to get an Expression. If expression is parsed with an error, nil is returned
func Parse(r _ee.Reader) Expression {
	_cdef := &plex{_cegce: LexReader(r)}
	_gdbdb(_cdef)
	if _cdef._ddcd != "" {
		return nil
	}
	return _cdef._ffccb
}

var _ebc = [...]uint8{0, 16, 29, 43, 56, 68, 80, 91, 102, 113, 125, 137, 148, 163}

func _fcca(_dbgg float64) float64 { return float64(int(_dbgg + 0.5)) }

// Eval evaluates and returns the result of a sheet expression.
func (_dfgb SheetPrefixExpr) Eval(ctx Context, ev Evaluator) Result {
	return MakeErrorResult("\u0073\u0068\u0065\u0065\u0074\u0020\u0070\u0072\u0065\u0066\u0069\u0078\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u006e\u0065\u0076\u0065r\u0020\u0062\u0065\u0020\u0065v\u0061\u006cu\u0061\u0074\u0065\u0064")
}

// Reference returns a string reference value to a cell.
func (_egb CellRef) Reference(ctx Context, ev Evaluator) Reference {
	return Reference{Type: ReferenceTypeCell, Value: _egb._ab}
}
func _gfcf(_fegc string, _cfab func(_dgdeg float64) float64) Function {
	return func(_fdac []Result) Result {
		if len(_fdac) != 1 {
			return MakeErrorResult(_fegc + "\u0020\u0072\u0065\u0071ui\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
		}
		_ddgbc := _fdac[0].AsNumber()
		switch _ddgbc.Type {
		case ResultTypeNumber:
			_gabg := _cfab(_ddgbc.ValueNumber)
			if _cc.IsNaN(_gabg) {
				return MakeErrorResult(_fegc + "\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0065\u0064\u0020\u004e\u0061\u004e")
			}
			if _cc.IsInf(_gabg, 0) {
				return MakeErrorResult(_fegc + "\u0020r\u0065t\u0075\u0072\u006e\u0065\u0064 \u0069\u006ef\u0069\u006e\u0069\u0074\u0079")
			}
			return MakeNumberResult(_gabg)
		case ResultTypeList, ResultTypeString:
			return MakeErrorResult(_fegc + "\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u006e\u0075\u006de\u0072i\u0063\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
		case ResultTypeError:
			return _ddgbc
		default:
			return MakeErrorResult(_e.Sprintf("\u0075\u006e\u0068a\u006e\u0064\u006c\u0065d\u0020\u0025\u0073\u0028\u0029\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _fegc, _ddgbc.Type))
		}
	}
}

// Rept is an implementation of the Excel REPT function that returns n copies of
// a string.
func Rept(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("R\u0045\u0050\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	_faad := args[0].AsString()
	if _faad.Type != ResultTypeString {
		return MakeErrorResult("\u0050R\u004f\u0050E\u0052\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020f\u0069\u0072\u0073\u0074\u0020\u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062e\u0020\u0061\u0020\u0073\u0074\u0072\u0069\u006e\u0067")
	}
	_ebad := args[1].AsNumber()
	if _ebad.Type != ResultTypeNumber {
		return MakeErrorResult("\u0050\u0052O\u0050\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0061\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	if _ebad.ValueNumber < 0 {
		return MakeErrorResult("\u0050\u0052\u004fP\u0045\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074o\u0020\u0062\u0065\u0020\u003e\u003d\u0020\u0030")
	}
	if _ebad.ValueNumber == 0 {
		return MakeStringResult("")
	}
	_ggafa := _bd.Buffer{}
	for _ccda := 0; _ccda < int(_ebad.ValueNumber); _ccda++ {
		_ggafa.WriteString(_faad.ValueString)
	}
	return MakeStringResult(_ggafa.String())
}
func _fgcbc(_bgbc string, _ebec _eb.Time) (_eb.Time, error) {
	_eedff, _, _faag := _eeg.ParseFloat(_bgbc, 10, 128, _eeg.ToNearestEven)
	if _faag != nil {
		return _eb.Time{}, _faag
	}
	_adeaa := new(_eeg.Float)
	_adeaa.SetUint64(uint64(24 * _eb.Hour))
	_eedff.Mul(_eedff, _adeaa)
	_fbff, _ := _eedff.Uint64()
	_babeb := _ebec.Add(_eb.Duration(_fbff))
	return _eebbg(_babeb), nil
}
func _cabg(_dbda string) string {
	_dbda = _ga.Replace(_dbda, "\u000a", "\u005c\u006e", -1)
	_dbda = _ga.Replace(_dbda, "\u000d", "\u005c\u0072", -1)
	_dbda = _ga.Replace(_dbda, "\u0009", "\u005c\u0074", -1)
	return _dbda
}

// Yielddisc implements the Excel YIELDDISC function.
func Yielddisc(args []Result) Result {
	_efaad := len(args)
	if _efaad != 4 && _efaad != 5 {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044D\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020f\u006f\u0075\u0072\u0020\u006f\u0072\u0020\u0066\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ebgg, _egbga, _ebbda := _gfcg(args[0], args[1], "\u0059I\u0045\u004c\u0044\u0044\u0049\u0053C")
	if _ebbda.Type == ResultTypeError {
		return _ebbda
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044\u0044\u0049S\u0043\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0070\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fecb := args[2].ValueNumber
	if _fecb <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0049E\u004c\u0044\u0044\u0049\u0053C\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0070\u0072\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0049\u0045\u004c\u0044D\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020r\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	_fcbf := args[3].ValueNumber
	if _fcbf <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "YI\u0045\u004cD\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065\u006d\u0070\u0074\u0069\u006f\u006e\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076e\u0020n\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072g\u0075m\u0065\u006et")
	}
	_ggff := 0
	if _efaad == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("\u0059\u0049E\u004c\u0044\u0044\u0049\u0053\u0043\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
		}
		_ggff = int(args[4].ValueNumber)
		if !_gfe(_ggff) {
			return MakeErrorResultType(ErrorTypeNum, "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074 \u0066\u006f\u0072\u0020\u0059I\u0045\u004cD\u0044\u0049\u0053\u0043")
		}
	}
	_ffda, _ebbda := _fcfa(_ebgg, _egbga, _ggff)
	if _ebbda.Type == ResultTypeError {
		return _ebbda
	}
	return MakeNumberResult((_fcbf/_fecb - 1) / _ffda)
}

// Oddlprice implements the Excel ODDLPRICE function.
func Oddlprice(args []Result) Result {
	if len(args) != 8 && len(args) != 9 {
		return MakeErrorResult("\u004f\u0044\u0044L\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0065\u0069\u0067\u0068\u0074\u0020\u006f\u0072\u0020\u006e\u0069\u006e\u0065\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_cbbc, _acfff, _efc := _gfcg(args[0], args[1], "\u004fD\u0044\u004c\u0050\u0052\u0049\u0043E")
	if _efc.Type == ResultTypeError {
		return _efc
	}
	_acbda, _efc := _eec(args[2], "\u0069\u0073\u0073\u0075\u0065\u0020\u0064\u0061\u0074\u0065", "\u004fD\u0044\u004c\u0050\u0052\u0049\u0043E")
	if _efc.Type == ResultTypeError {
		return _efc
	}
	if _acbda >= _cbbc {
		return MakeErrorResultType(ErrorTypeNum, "\u004c\u0061\u0073\u0074\u0020i\u006e\u0074\u0065\u0072\u0065\u0073\u0074\u0020\u0064\u0061\u0074\u0065\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0062\u0065\u0066\u006f\u0072\u0065\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074e")
	}
	_fece := args[3]
	if _fece.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020o\u0066\u0020\u0074\u0079\u0070e\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_adcc := _fece.ValueNumber
	if _adcc < 0 {
		return MakeErrorResultType(ErrorTypeNum, "R\u0061\u0074\u0065\u0020\u0073\u0068o\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u006e\u006fn\u0020\u006e\u0065g\u0061t\u0069\u0076\u0065")
	}
	_feg := args[4]
	if _feg.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0079i\u0065\u006c\u0064\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
	}
	_ccdd := _feg.ValueNumber
	if _ccdd < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006cd\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u006eo\u006e\u0020\u006e\u0065\u0067\u0061\u0074i\u0076\u0065")
	}
	_eafa := args[5]
	if _eafa.Type != ResultTypeNumber {
		return MakeErrorResult("\u004fD\u0044\u004cP\u0052\u0049\u0043\u0045 \u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0072\u0065\u0064\u0065mp\u0074\u0069\u006fn\u0020\u006ff\u0020\u0074\u0079\u0070\u0065\u0020n\u0075\u006db\u0065\u0072")
	}
	_gbac := _eafa.ValueNumber
	if _gbac < 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0059\u0069\u0065\u006cd\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u006eo\u006e\u0020\u006e\u0065\u0067\u0061\u0074i\u0076\u0065")
	}
	_gcag := args[6]
	if _gcag.Type != ResultTypeNumber {
		return MakeErrorResult("\u004f\u0044\u0044\u004c\u0050\u0052\u0049C\u0045\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020\u0066\u0072\u0065\u0071\u0075\u0065\u006e\u0063\u0079\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072")
	}
	_ddfc := float64(int(_gcag.ValueNumber))
	if !_dddg(_ddfc) {
		return MakeErrorResultType(ErrorTypeNum, "\u0049n\u0063\u006f\u0072\u0072e\u0063\u0074\u0020\u0066\u0072e\u0071u\u0065n\u0063\u0065\u0020\u0076\u0061\u006c\u0075e")
	}
	_cba := 0
	if len(args) == 8 && args[7].Type != ResultTypeEmpty {
		_ccff := args[7]
		if _ccff.Type != ResultTypeNumber {
			return MakeErrorResult("\u004f\u0044\u0044\u004c\u0050\u0052\u0049\u0043\u0045\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0062a\u0073\u0069\u0073\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u006eu\u006d\u0062\u0065\u0072")
		}
		_cba = int(_ccff.ValueNumber)
		if !_gfe(_cba) {
			return MakeErrorResultType(ErrorTypeNum, "I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0062\u0061\u0073\u0069s\u0020\u0076\u0061\u006c\u0075\u0065\u0020f\u006f\u0072\u0020\u004f\u0044\u0044\u004c\u0050\u0052\u0049C\u0045")
		}
	}
	_dega, _efc := _fcfa(_acbda, _acfff, _cba)
	if _efc.Type == ResultTypeError {
		return _efc
	}
	_dega *= _ddfc
	_bfde, _efc := _fcfa(_cbbc, _acfff, _cba)
	if _efc.Type == ResultTypeError {
		return _efc
	}
	_bfde *= _ddfc
	_agggb, _efc := _fcfa(_acbda, _cbbc, _cba)
	if _efc.Type == ResultTypeError {
		return _efc
	}
	_agggb *= _ddfc
	_bfba := _gbac + _dega*100*_adcc/_ddfc
	_bfba /= _bfde*_ccdd/_ddfc + 1
	_bfba -= _agggb * 100 * _adcc / _ddfc
	return MakeNumberResult(_bfba)
}

//go:generate ragel -G2 -Z lexer.rl
//go:generate goimports -w lexer.go
type Lexer struct {
	_cgdbfa chan *node
	_gegda  _af.Mutex
	_abeff  []chan *node
	_cdcgc  []*node
}

const _ebfg = 57349

func _efce(_ebbe Result, _efed, _dggbd int) [][]Result {
	_bbfa := [][]Result{}
	switch _ebbe.Type {
	case ResultTypeArray:
		for _dabc, _ddgba := range _ebbe.ValueArray {
			if _dabc < _efed {
				_bbfa = append(_bbfa, _gefg(MakeListResult(_ddgba), _dggbd))
			} else {
				_bbfa = append(_bbfa, _gefg(MakeErrorResultType(ErrorTypeNA, ""), _dggbd))
			}
		}
	case ResultTypeList:
		_edcg := _gefg(_ebbe, _dggbd)
		for _ebafe := 0; _ebafe < _efed; _ebafe++ {
			_bbfa = append(_bbfa, _edcg)
		}
	case ResultTypeNumber, ResultTypeString, ResultTypeError, ResultTypeEmpty:
		for _dbfdb := 0; _dbfdb < _efed; _dbfdb++ {
			_gbbf := _gefg(_ebbe, _dggbd)
			_bbfa = append(_bbfa, _gbbf)
		}
	}
	return _bbfa
}

// Irr implements the Excel IRR function.
func Irr(args []Result) Result {
	_bdaf := len(args)
	if _bdaf == 0 || _bdaf > 2 {
		return MakeErrorResult("\u0049\u0052\u0052\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006f\u006e\u0065\u0020\u006f\u0072\u0020t\u0077\u006f\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if args[0].Type != ResultTypeList && args[0].Type != ResultTypeArray {
		return MakeErrorResult("\u0049\u0052\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020v\u0061\u006c\u0075\u0065\u0073\u0020t\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0061\u0072\u0072\u0061\u0079 \u0074\u0079\u0070\u0065")
	}
	_afde := _aadag(args[0])
	_abed := []float64{}
	for _, _fadf := range _afde {
		for _, _cfeb := range _fadf {
			if _cfeb.Type == ResultTypeNumber && !_cfeb.IsBoolean {
				_abed = append(_abed, _cfeb.ValueNumber)
			}
		}
	}
	_bcaf := len(_abed)
	if len(_abed) < 2 {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
	_fgfc := 0.1
	if _bdaf == 2 && args[1].Type != ResultTypeEmpty {
		if args[1].Type != ResultTypeNumber {
			return MakeErrorResult("I\u0052\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0067\u0075\u0065\u0073\u0073\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
		}
		_fgfc = args[1].ValueNumber
		if _fgfc <= -1 {
			return MakeErrorResult("\u0049\u0052R\u0020\u0072\u0065\u0071u\u0069\u0072e\u0073\u0020\u0067\u0075\u0065\u0073\u0073\u0020t\u006f\u0020\u0062\u0065\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068a\u006e\u0020\u002d\u0031")
		}
	}
	_bcac := []float64{}
	for _ebdgf := 0; _ebdgf < _bcaf; _ebdgf++ {
		if _ebdgf == 0 {
			_bcac = append(_bcac, 0)
		} else {
			_bcac = append(_bcac, _bcac[_ebdgf-1]+365)
		}
	}
	return _dec(_abed, _bcac, _fgfc)
}

// Eval evaluates and returns the result of a function call.
func (_dbgge FunctionCall) Eval(ctx Context, ev Evaluator) Result {
	_gaef := LookupFunction(_dbgge._eecfd)
	if _gaef != nil {
		_gcccf := make([]Result, len(_dbgge._dgce))
		for _bedfc, _fdddd := range _dbgge._dgce {
			_gcccf[_bedfc] = _fdddd.Eval(ctx, ev)
			_gcccf[_bedfc].Ref = _fdddd.Reference(ctx, ev)
		}
		if _, _bgecb := _gegbg[_dbgge._eecfd]; !_bgecb {
			if _cdcff, _fbgea := _geffd(_gcccf); _cdcff {
				return _fbgea
			}
		}
		return _gaef(_gcccf)
	}
	_acdff := LookupFunctionComplex(_dbgge._eecfd)
	if _acdff != nil {
		_bcbaf := make([]Result, len(_dbgge._dgce))
		for _eggbe, _dfaef := range _dbgge._dgce {
			_bcbaf[_eggbe] = _dfaef.Eval(ctx, ev)
			_bcbaf[_eggbe].Ref = _dfaef.Reference(ctx, ev)
		}
		if _, _eafg := _gegbg[_dbgge._eecfd]; !_eafg {
			if _eggg, _abfcd := _geffd(_bcbaf); _eggg {
				return _abfcd
			}
		}
		return _acdff(ctx, ev, _bcbaf)
	}
	return MakeErrorResult("\u0075\u006e\u006b\u006e\u006f\u0077\u006e\u0020\u0066\u0075\u006e\u0063t\u0069\u006f\u006e\u0020" + _dbgge._eecfd)
}

// Year is an implementation of the Excel YEAR() function.
func Year(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0059\u0045\u0041\u0052\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_ddb := ctx.GetEpoch()
	_badc, _cff := _fgcbc(args[0].Value(), _ddb)
	if _cff != nil {
		return MakeErrorResult("\u0059\u0045AR\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s a\u0020si\u006e\u0067\u006c\u0065\u0020\u0064\u0061te\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeNumberResult(float64(_badc.Year()))
}

// HorizontalRange is a range expression that when evaluated returns a list of Results from references like 1:4 (all cells from rows 1 to 4).
type HorizontalRange struct{ _fcbd, _fbcg int }

// Eval evaluates a vertical range returning a list of results or an error.
func (_cccge VerticalRange) Eval(ctx Context, ev Evaluator) Result {
	_eafcb := _cccge.verticalRangeReference()
	if _gbfbg, _ddfb := ev.GetFromCache(_eafcb); _ddfb {
		return _gbfbg
	}
	_fddfdg, _dbff := _cfaeb(ctx, _cccge._ggcb, _cccge._acdd)
	_aceg := _gebc(ctx, ev, _fddfdg, _dbff)
	ev.SetCache(_eafcb, _aceg)
	return _aceg
}
func _faafg(_degbe string, _gdgf []Result) (*parsedSearchObject, Result) {
	_cbeae := len(_gdgf)
	if _cbeae != 2 && _cbeae != 3 {
		return nil, MakeErrorResult(_degbe + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f\u0020\u006fr\u0020t\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	_ebde := _gdgf[0]
	if _ebde.Type == ResultTypeError {
		return nil, _ebde
	}
	if _ebde.Type != ResultTypeString && _ebde.Type != ResultTypeNumber {
		return nil, MakeErrorResult("\u0054\u0068e\u0020\u0066\u0069\u0072s\u0074\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020s\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065\u0020\u0061\u0020\u0073t\u0072\u0069\u006e\u0067")
	}
	_gadca := _gdgf[1]
	if _gadca.Type == ResultTypeError {
		return nil, _gadca
	}
	if _gadca.Type != ResultTypeString && _gadca.Type != ResultTypeNumber {
		return nil, MakeErrorResult("\u0054\u0068\u0065\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020a\u0020\u0073t\u0072\u0069\u006e\u0067")
	}
	_bbba := _gadca.Value()
	_cccea := _ebde.Value()
	_efcgg := 1
	if _cbeae == 3 && _gdgf[2].Type != ResultTypeEmpty {
		_edeea := _gdgf[2]
		if _edeea.Type != ResultTypeNumber {
			return nil, MakeErrorResult("P\u006f\u0073\u0069\u0074\u0069\u006fn\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062e\u0020\u0061\u0020n\u0075m\u0062\u0065\u0072")
		}
		_efcgg = int(_edeea.ValueNumber)
		if _efcgg < 1 {
			return nil, MakeErrorResultType(ErrorTypeValue, "\u0050\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0061 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0074h\u0061\u006e\u0020\u0030")
		}
		if _efcgg > len(_bbba) {
			return nil, MakeErrorResultType(ErrorTypeValue, "\u0050\u006f\u0073\u0069\u0074\u0069\u006f\u006e\u0020\u0073\u0068\u006f\u0075l\u0064\u0020\u0062\u0065\u0020\u0061 \u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006d\u006f\u0072\u0065\u0020\u0074h\u0061\u006e\u0020\u0030")
		}
	}
	return &parsedSearchObject{_cccea, _bbba, _efcgg}, _feb
}

var _bdcac = []ri{{1000, "\u004d"}, {990, "\u0058\u004d"}, {950, "\u004c\u004d"}, {900, "\u0043\u004d"}, {500, "\u0044"}, {490, "\u0058\u0044"}, {450, "\u004c\u0044"}, {400, "\u0043\u0044"}, {100, "\u0043"}, {99, "\u0049\u0043"}, {90, "\u0058\u0043"}, {50, "\u004c"}, {45, "\u0056\u004c"}, {40, "\u0058\u004c"}, {10, "\u0058"}, {9, "\u0049\u0058"}, {5, "\u0056"}, {4, "\u0049\u0056"}, {1, "\u0049"}}

// NewConstArrayExpr constructs a new constant array expression with a given data.
func NewConstArrayExpr(data [][]Expression) Expression { return &ConstArrayExpr{_ba: data} }
func _dcbb(_dgfe Reference, _ffed Context) bool {
	return _ffed.Sheet(_dgfe.Value) == InvalidReferenceContext
}
func _beagb(_febeb [][]Result) float64 {
	if len(_febeb) == 2 {
		_bede := _febeb[0][0].AsNumber()
		_abeg := _febeb[0][1].AsNumber()
		_dgdb := _febeb[1][0].AsNumber()
		_ccfdd := _febeb[1][1].AsNumber()
		if _bede.Type != ResultTypeNumber || _abeg.Type != ResultTypeNumber || _dgdb.Type != ResultTypeNumber || _ccfdd.Type != ResultTypeNumber {
			return _cc.NaN()
		}
		return _bede.ValueNumber*_ccfdd.ValueNumber - _dgdb.ValueNumber*_abeg.ValueNumber
	}
	_bgde := float64(0)
	_gfcdf := float64(1)
	for _eaeb := range _febeb {
		_bgde += _gfcdf * _febeb[0][_eaeb].ValueNumber * _beagb(_bgga(_febeb, _eaeb))
		_gfcdf *= -1
	}
	return _bgde
}

// Update makes a reference to point to one of the neighboring cells after removing a row/column with respect to the update type.
func (_fbe CellRef) Update(q *_eg.UpdateQuery) Expression {
	if q.UpdateCurrentSheet {
		_fbe._ab = _dfa(_fbe._ab, q)
	}
	return _fbe
}

// Reference returns an invalid reference for BinaryExpr.
func (_df BinaryExpr) Reference(ctx Context, ev Evaluator) Reference { return ReferenceInvalid }
func _bdcee(_bgae []Result, _eaffb bool) Result {
	_gefgc := "\u004d\u0049\u004e"
	if _eaffb {
		_gefgc = "\u004d\u0049\u004e\u0041"
	}
	if len(_bgae) == 0 {
		return MakeErrorResult(_gefgc + "\u0020\u0072\u0065q\u0075\u0069\u0072\u0065s\u0020\u0061\u0074\u0020\u006c\u0065\u0061s\u0074\u0020\u006f\u006e\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_debcg := _cc.MaxFloat64
	for _, _gbbfc := range _bgae {
		switch _gbbfc.Type {
		case ResultTypeNumber:
			if (_eaffb || !_gbbfc.IsBoolean) && _gbbfc.ValueNumber < _debcg {
				_debcg = _gbbfc.ValueNumber
			}
		case ResultTypeList, ResultTypeArray:
			_gfccc := _bdcee(_gbbfc.ListValues(), _eaffb)
			if _gfccc.ValueNumber < _debcg {
				_debcg = _gfccc.ValueNumber
			}
		case ResultTypeEmpty:
		case ResultTypeString:
			_gfcdd := 0.0
			if _eaffb {
				_gfcdd = _gbbfc.AsNumber().ValueNumber
			}
			if _gfcdd < _debcg {
				_debcg = _gfcdd
			}
		default:
			_cf.Log.Debug("\u0075\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020"+_gefgc+"\u0028\u0029\u0020\u0061rg\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _gbbfc.Type)
		}
	}
	if _debcg == _cc.MaxFloat64 {
		_debcg = 0
	}
	return MakeNumberResult(_debcg)
}

// NewFunction constructs a new function call expression.
func NewFunction(name string, args []Expression) Expression {
	return FunctionCall{_eecfd: name, _dgce: args}
}
func _ffee(_bdfe []Result, _bfa string) (*amorArgs, Result) {
	_degc := len(_bdfe)
	if _degc != 6 && _degc != 7 {
		return nil, MakeErrorResult(_bfa + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0069\u0078\u0020\u006fr\u0020s\u0065\u0076\u0065\u006e\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _bdfe[0].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bfa + "\u0020\u0072eq\u0075\u0069\u0072e\u0073\u0020\u0063\u006fst \u0074o \u0062\u0065\u0020\u006e\u0075\u006d\u0062er\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_baff := _bdfe[0].ValueNumber
	if _baff < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bfa+"\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0063\u006f\u0073\u0074\u0020\u0074\u006f\u0020\u0062\u0065 \u006e\u006f\u006e\u0020\u006e\u0065\u0067a\u0074\u0069\u0076\u0065")
	}
	_ddf, _cead := _eec(_bdfe[1], "\u0064\u0061\u0074\u0065\u0020\u0070\u0075\u0072\u0063h\u0061\u0073\u0065\u0064", _bfa)
	if _cead.Type == ResultTypeError {
		return nil, _cead
	}
	_dbcd, _cead := _eec(_bdfe[2], "\u0066\u0069\u0072s\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064", _bfa)
	if _cead.Type == ResultTypeError {
		return nil, _cead
	}
	if _dbcd < _ddf {
		return nil, MakeErrorResultType(ErrorTypeNum, _bfa+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0066\u0069\u0072\u0073\u0074 \u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020l\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0064\u0061te\u0020\u0070\u0075\u0072\u0063\u0068\u0061\u0073\u0065\u0064")
	}
	if _bdfe[3].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bfa + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006cv\u0061\u0067\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_gadb := _bdfe[3].ValueNumber
	if _gadb < 0 || _gadb > _baff {
		return nil, MakeErrorResultType(ErrorTypeNum, _bfa+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061g\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0062\u0065\u0074\u0077\u0065e\u006e\u0020\u0030\u0020\u0061\u006e\u0064\u0020\u0074\u0068\u0065\u0020in\u0069\u0074\u0069\u0061\u006c\u0020\u0063\u006f\u0073\u0074")
	}
	if _bdfe[4].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bfa + "\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_cede := int(_bdfe[4].ValueNumber)
	if _cede < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bfa+" \u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0070\u0065\u0072\u0069o\u0064\u0020\u0074\u006f\u0020\u0062\u0065 \u006e\u006f\u006e\u002d\u006e\u0065\u0067\u0061\u0074\u0069v\u0065")
	}
	if _bdfe[5].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bfa + "\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0064\u0065\u0070\u0072\u0065\u0063\u0069\u0061\u0074\u0069\u006f\u006e\u0020\u0072\u0061\u0074\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_fcfg := _bdfe[5].ValueNumber
	if _fcfg < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bfa+"\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073 d\u0065\u0070\u0072\u0065\u0063\u0069\u0061\u0074\u0069\u006f\u006e\u0020\u0072\u0061t\u0065\u0020t\u006f\u0020\u0062e\u0020\u006d\u006f\u0072\u0065\u0020\u0074\u0068\u0061\u006e\u0020\u0030\u0020\u0061\u006e\u0064 \u006c\u0065ss\u0020\u0074\u0068a\u006e\u0020\u0030\u002e\u0035")
	}
	_gcd := 0
	if _degc == 7 && _bdfe[6].Type != ResultTypeEmpty {
		if _bdfe[6].Type != ResultTypeNumber {
			return nil, MakeErrorResult(_bfa + "\u0020\u0072e\u0071\u0075\u0069\u0072e\u0073\u0020b\u0061\u0073\u0069\u0073\u0020\u0074\u006f\u0020b\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		_gcd = int(_bdfe[6].ValueNumber)
		if !_gfe(_gcd) || _gcd == 2 {
			return nil, MakeErrorResultType(ErrorTypeNum, "\u0049\u006ec\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0062\u0061\u0073\u0069\u0073\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020fo\u0072\u0020"+_bfa)
		}
	}
	return &amorArgs{_baff, _ddf, _dbcd, _gadb, _cede, _fcfg, _gcd}, _feb
}

// HLookup implements the HLOOKUP function that returns a matching value from a
// row in an array.
func HLookup(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("\u0048\u004c\u004f\u004f\u004bU\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074 \u006c\u0065\u0061\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if len(args) > 4 {
		return MakeErrorResult("\u0048\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u0061\u0074\u0020m\u006f\u0073\u0074\u0020\u0066\u006f\u0075\u0072\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_gccb := args[0]
	_geba := args[1]
	if _geba.Type != ResultTypeArray {
		return MakeErrorResult("\u0048\u004cO\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_eafag := args[2].AsNumber()
	if _eafag.Type != ResultTypeNumber {
		return MakeErrorResult("\u0048\u004cO\u004f\u004b\u0055\u0050 \u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075m\u0065\u0072\u0069\u0063\u0020\u0072\u006f\u0077\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_dcee := false
	if len(args) == 4 {
		_faafe := args[3].AsNumber()
		if _faafe.Type != ResultTypeNumber {
			return MakeErrorResult("\u0048\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006de\u0072\u0069\u0063\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
		}
		if _faafe.ValueNumber == 0 {
			_dcee = true
		}
	}
	_bedd := -1
	_dagc := false
	if len(_geba.ValueArray) == 0 {
		return MakeErrorResult("\u0048\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065s\u0020n\u006f\u006e\u002d\u0065\u006d\u0070\u0074\u0079\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_cabc := _geba.ValueArray[0]
_gaaa:
	for _eccg, _cbcb := range _cabc {
		switch _edeef(_cbcb, _gccb, false, _dcee) {
		case _gcefb:
			_bedd = _eccg
		case _cgae:
			_bedd = _eccg
			_dagc = true
			break _gaaa
		}
	}
	if _bedd == -1 {
		return MakeErrorResultType(ErrorTypeNA, "\u0048\u004c\u004fOK\u0055\u0050\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
	}
	_fdaga := int(_eafag.ValueNumber) - 1
	if _fdaga < 0 || _fdaga > len(_geba.ValueArray) {
		return MakeErrorResult("\u0048L\u004f\u004f\u004b\u0055P\u0020\u0068\u0061\u0064\u0020i\u006ev\u0061l\u0069\u0064\u0020\u0069\u006e\u0064\u0065x")
	}
	_cabc = _geba.ValueArray[_fdaga]
	if _bedd < 0 || _bedd >= len(_cabc) {
		return MakeErrorResult("\u0056\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0068\u0061\u0073\u0020\u0069\u006e\u0076a\u006ci\u0064\u0020\u0063\u006f\u006c\u0075\u006d\u006e\u0020\u0069\u006e\u0064\u0065\u0078")
	}
	if _dagc || !_dcee {
		return _cabc[_bedd]
	}
	return MakeErrorResultType(ErrorTypeNA, "\u0056\u004c\u004fOK\u0055\u0050\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075\u006e\u0064")
}

const _gbe = "\u0049\u006e\u0063\u006f\u0072\u0072\u0065\u0063\u0074\u0020\u0061\u0072\u0067\u0075\u006de\u006et\u0020\u0066\u006f\u0072\u0020\u0044\u0041\u0054\u0045\u0056\u0041\u004c\u0055\u0045"

// NewNegate constructs a new negate expression.
func NewNegate(e Expression) Expression { return Negate{_effb: e} }

var _bcdab = [...]int{1}

const _faef = "(\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u002d" + _bddc + "-\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029"

var _fceb = [...]string{"\u0024\u0065\u006e\u0064", "\u0065\u0072\u0072o\u0072", "\u0024\u0075\u006e\u006b", "t\u006fk\u0065\u006e\u0048\u006f\u0072\u0069\u007a\u006fn\u0074\u0061\u006c\u0052an\u0067\u0065", "\u0074o\u006be\u006e\u0056\u0065\u0072\u0074i\u0063\u0061l\u0052\u0061\u006e\u0067\u0065", "\u0074\u006f\u006b\u0065\u006e\u0052\u0065\u0073\u0065\u0072\u0076\u0065d\u004e\u0061\u006d\u0065", "\u0074\u006f\u006be\u006e\u0044\u0044\u0045\u0043\u0061\u006c\u006c", "\u0074\u006f\u006b\u0065\u006e\u004c\u0065\u0078\u0045\u0072\u0072\u006f\u0072", "\u0074o\u006be\u006e\u004e\u0061\u006d\u0065\u0064\u0052\u0061\u006e\u0067\u0065", "\u0074o\u006b\u0065\u006e\u0042\u006f\u006fl", "t\u006f\u006b\u0065\u006e\u004e\u0075\u006d\u0062\u0065\u0072", "t\u006f\u006b\u0065\u006e\u0053\u0074\u0072\u0069\u006e\u0067", "\u0074\u006f\u006b\u0065\u006e\u0045\u0072\u0072\u006f\u0072", "\u0074\u006f\u006b\u0065\u006e\u0045\u0072\u0072\u006f\u0072\u0052\u0065\u0066", "\u0074\u006f\u006b\u0065\u006e\u0053\u0068\u0065\u0065\u0074", "\u0074o\u006b\u0065\u006e\u0043\u0065\u006cl", "t\u006fk\u0065\u006e\u0046\u0075\u006e\u0063\u0074\u0069o\u006e\u0042\u0075\u0069lt\u0069\u006e", "t\u006f\u006b\u0065\u006e\u004c\u0042\u0072\u0061\u0063\u0065", "t\u006f\u006b\u0065\u006e\u0052\u0042\u0072\u0061\u0063\u0065", "t\u006f\u006b\u0065\u006e\u004c\u0050\u0061\u0072\u0065\u006e", "t\u006f\u006b\u0065\u006e\u0052\u0050\u0061\u0072\u0065\u006e", "\u0074o\u006b\u0065\u006e\u0050\u006c\u0075s", "\u0074\u006f\u006b\u0065\u006e\u004d\u0069\u006e\u0075\u0073", "\u0074o\u006b\u0065\u006e\u004d\u0075\u006ct", "\u0074\u006f\u006b\u0065\u006e\u0044\u0069\u0076", "\u0074\u006f\u006b\u0065\u006e\u0045\u0078\u0070", "\u0074o\u006b\u0065\u006e\u0045\u0051", "\u0074o\u006b\u0065\u006e\u004c\u0054", "\u0074o\u006b\u0065\u006e\u0047\u0054", "\u0074\u006f\u006b\u0065\u006e\u004c\u0045\u0051", "\u0074\u006f\u006b\u0065\u006e\u0047\u0045\u0051", "\u0074o\u006b\u0065\u006e\u004e\u0045", "\u0074\u006f\u006b\u0065\u006e\u0043\u006f\u006c\u006f\u006e", "\u0074\u006f\u006b\u0065\u006e\u0043\u006f\u006d\u006d\u0061", "\u0074\u006f\u006b\u0065\u006e\u0041\u006d\u0070\u0065r\u0073\u0061\u006e\u0064", "\u0074o\u006b\u0065\u006e\u0053\u0065\u006di"}

// SetLocked does nothing for the invalid reference context.
func (_agfee *ivr) SetLocked(cellRef string, locked bool) {}
func (_effcb Result) AsString() Result {
	switch _effcb.Type {
	case ResultTypeNumber:
		return MakeStringResult(_effcb.Value())
	default:
		return _effcb
	}
}

// Code is an implementation of the Excel CODE function that returns the first
// character of the string as a number.
func Code(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0043\u004f\u0044\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_adedb := args[0].AsString()
	if _adedb.Type != ResultTypeString {
		return MakeErrorResult("\u0043\u004f\u0044\u0045\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0073t\u0072\u0069\u006e\u0067\u0020a\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	if len(_adedb.ValueString) == 0 {
		return MakeNumberResult(0)
	}
	return MakeNumberResult(float64(_adedb.ValueString[0]))
}

const _adaag = 57356

func _dagb(_ggcf, _cbfe, _cbed, _cadd, _eebed float64, _eefa int) float64 {
	_aaadg := _ecc(_ggcf, _cbed, _cadd, _eebed, _eefa)
	var _efac float64
	if _cbfe == 1 {
		if _eefa == 1 {
			_efac = 0
		} else {
			_efac = -_cadd
		}
	} else {
		if _eefa == 1 {
			_efac = _ggbd(_ggcf, _cbfe-2, _aaadg, _cadd, 1) - _aaadg
		} else {
			_efac = _ggbd(_ggcf, _cbfe-1, _aaadg, _cadd, 0)
		}
	}
	return _efac * _ggcf
}

// Averagea implements the AVERAGEA function, AVERAGEA counts cells that contain
// text as a zero where AVERAGE ignores them entirely.
func Averagea(args []Result) Result {
	_ddfg, _cbgga := _ebbcf(args, true)
	if _cbgga == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0041\u0056\u0045\u0052AG\u0045\u0020\u0064\u0069\u0076\u0069\u0064\u0065\u0020\u0062\u0079\u0020\u007a\u0065r\u006f")
	}
	return MakeNumberResult(_ddfg / _cbgga)
}
func _gege(_gbead, _bbdba, _egbb, _ccbg, _aagdd float64) float64 {
	var _dcfcd float64
	_fbcf := _aagdd / _egbb
	if _fbcf >= 1 {
		_fbcf = 1
		if _ccbg == 1 {
			_dcfcd = _gbead
		} else {
			_dcfcd = 0
		}
	} else {
		_dcfcd = _gbead * _cc.Pow(1-_fbcf, _ccbg-1)
	}
	_bbecgf := _gbead * _cc.Pow(1-_fbcf, _ccbg)
	var _dbbd float64
	if _bbecgf < _bbdba {
		_dbbd = _dcfcd - _bbdba
	} else {
		_dbbd = _dcfcd - _bbecgf
	}
	if _dbbd < 0 {
		_dbbd = 0
	}
	return _dbbd
}

// ISBLANK is an implementation of the Excel ISBLANK() function.
func IsBlank(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("I\u0053\u0042\u004c\u0041\u004e\u004b(\u0029\u0020\u0061\u0063\u0063\u0065p\u0074\u0073\u0020\u0061\u0020\u0073\u0069n\u0067\u006c\u0065\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	return MakeBoolResult(args[0].Type == ResultTypeEmpty)
}

// String returns a string representation of FunctionCall expression.
func (_bgebd FunctionCall) String() string {
	_cadca := _bd.Buffer{}
	_cadca.WriteString(_bgebd._eecfd)
	_cadca.WriteString("\u0028")
	_ddccc := len(_bgebd._dgce) - 1
	for _feaag, _dgab := range _bgebd._dgce {
		_cadca.WriteString(_dgab.String())
		if _feaag != _ddccc {
			_cadca.WriteString("\u002c")
		}
	}
	_cadca.WriteString("\u0029")
	return _cadca.String()
}

// Update returns the same object as updating sheet references does not affect named ranges.
func (_bdae NamedRangeRef) Update(q *_eg.UpdateQuery) Expression { return _bdae }
func _aeaf() {
	_cfgcg = _f.MustCompile("\u005e\u0028\u005b\u0030\u002d\u0039\u005d\u002b\u0029\u0024")
	_aded = _f.MustCompile("\u005e=\u0028\u002e\u002a\u0029\u0024")
	_caed = _f.MustCompile("\u005e<\u0028\u002e\u002a\u0029\u0024")
	_bdfd = _f.MustCompile("\u005e>\u0028\u002e\u002a\u0029\u0024")
	_ddae = _f.MustCompile("\u005e\u003c\u003d\u0028\u002e\u002a\u0029\u0024")
	_bgda = _f.MustCompile("\u005e\u003e\u003d\u0028\u002e\u002a\u0029\u0024")
}

type cumulArgs struct {
	_ffa   float64
	_dbcdb float64
	_bdeb  float64
	_gggb  float64
	_gffb  float64
	_cfgc  int
}

func _aabcb(_afbff []Result, _bbaaf string) (*cumulArgs, Result) {
	if len(_afbff) != 6 {
		return nil, MakeErrorResult(_bbaaf + "\u0020\u0072\u0065qu\u0069\u0072\u0065\u0073\u0020\u0073\u0069\u0078\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0073")
	}
	if _afbff[0].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bbaaf + "\u0020\u0072eq\u0075\u0069\u0072e\u0073\u0020\u0072\u0061te \u0074o \u0062\u0065\u0020\u006e\u0075\u006d\u0062er\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_ccdf := _afbff[0].ValueNumber
	if _ccdf <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bbaaf+"\u0020r\u0065\u0071u\u0069\u0072\u0065s\u0020\u0072\u0061\u0074\u0065\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006fs\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _afbff[1].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bbaaf + "\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061r\u0067u\u006d\u0065\u006e\u0074")
	}
	_caga := _afbff[1].ValueNumber
	if _caga <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bbaaf+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u006f\u0066\u0020p\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074\u006f \u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006eu\u006d\u0062\u0065\u0072\u0020\u0061r\u0067\u0075\u006de\u006e\u0074")
	}
	if _afbff[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bbaaf + "\u0020r\u0065\u0071u\u0069\u0072\u0065s\u0020\u0070\u0072\u0065\u0073\u0065\u006et\u0020\u0076\u0061\u006c\u0075\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_gcff := _afbff[2].ValueNumber
	if _gcff <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bbaaf+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0072\u0065\u0073\u0065n\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074\u006f\u0020\u0062\u0065 \u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006dbe\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _afbff[3].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bbaaf + "\u0020r\u0065\u0071u\u0069\u0072\u0065\u0073 \u0073\u0074\u0061r\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020to\u0020\u0062\u0065 \u006e\u0075m\u0062\u0065\u0072\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	_ddbg := _afbff[3].ValueNumber
	if _ddbg <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bbaaf+"\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073t\u0061\u0072\u0074\u0020\u0070\u0065\u0072\u0069o\u0064 \u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	if _afbff[4].Type != ResultTypeNumber {
		return nil, MakeErrorResult(_bbaaf + "\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0065\u006e\u0064\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_ddbe := _afbff[4].ValueNumber
	if _ddbe <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bbaaf+"\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0073\u0020\u0065\u006e\u0064\u0020\u0070\u0065\u0072\u0069\u006fd\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u0070\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	if _ddbe < _ddbg {
		return nil, MakeErrorResultType(ErrorTypeNum, _bbaaf+"\u0020\u0072\u0065\u0071u\u0069\u0072\u0065\u0073\u0020\u0065\u006e\u0064\u0020p\u0065\u0072\u0069\u006f\u0064\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006c\u0061\u0074\u0065\u0072\u0020o\u0072\u0020\u0065\u0071\u0075a\u006c\u0020\u0074\u006f\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0070\u0065\u0072\u0069\u006f\u0064")
	}
	if _ddbe > _caga {
		return nil, MakeErrorResultType(ErrorTypeNum, _bbaaf+" \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0070\u0065\u0072\u0069\u006f\u0064\u0073\u0020\u0074o\u0020\u0062\u0065\u0020\u0069\u006e\u0020\u0070\u0065\u0072io\u0064\u0073\u0020r\u0061n\u0067\u0065")
	}
	_dcdf := int(_afbff[5].ValueNumber)
	if _dcdf != 0 && _dcdf != 1 {
		return nil, MakeErrorResultType(ErrorTypeNum, _bbaaf+" \u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0074\u0079\u0070\u0065\u0020\u0074\u006f \u0062\u0065\u00200\u0020o\u0072\u0020\u0031")
	}
	return &cumulArgs{_ccdf, _caga, _gcff, _ddbg, _ddbe, _dcdf}, _feb
}

// NewCellRef constructs a new cell reference.
func NewCellRef(v string) Expression { return CellRef{_ab: v} }
func _edg(_dfec int) int {
	if _dfec < 1900 {
		if _dfec < 30 {
			_dfec += 2000
		} else {
			_dfec += 1900
		}
	}
	return _dfec
}

type countMode byte

// Update updates references in the PrefixHorizontalRange after removing a row/column.
func (_deeaf PrefixHorizontalRange) Update(q *_eg.UpdateQuery) Expression { return _deeaf }
func _faca(_bcgd []Result) Result {
	_ffaf := _bcgd[0].ValueList
	_gedgb := len(_ffaf)
	switch len(_bcgd) {
	case 1:
		_egde := []Result{}
		for _, _bbfab := range _ffaf {
			_egde = append(_egde, MakeBoolResult(_bbfab.ValueNumber != 0))
		}
		return MakeListResult(_egde)
	case 2:
		_cafcc := _bcgd[1]
		switch _cafcc.Type {
		case ResultTypeNumber, ResultTypeString, ResultTypeEmpty:
			_fbde := []Result{}
			for _, _baaab := range _ffaf {
				var _afga Result
				if _baaab.ValueNumber == 0 {
					_afga = MakeBoolResult(false)
				} else {
					_afga = _cafcc
				}
				_fbde = append(_fbde, _afga)
			}
			return MakeListResult(_fbde)
		case ResultTypeList:
			_gbacg := _gefg(_cafcc, _gedgb)
			_dcfb := []Result{}
			for _afce, _fbec := range _ffaf {
				var _eebdd Result
				if _fbec.ValueNumber == 0 {
					_eebdd = MakeBoolResult(false)
				} else {
					_eebdd = _gbacg[_afce]
				}
				_dcfb = append(_dcfb, _eebdd)
			}
			return MakeListResult(_dcfb)
		case ResultTypeArray:
			_ddcc := _efce(_cafcc, len(_cafcc.ValueArray), _gedgb)
			_adgc := [][]Result{}
			for _, _faeg := range _ddcc {
				_cgbec := []Result{}
				for _ebea, _eaef := range _ffaf {
					var _fdfda Result
					if _eaef.ValueNumber == 0 {
						_fdfda = MakeBoolResult(false)
					} else {
						_fdfda = _faeg[_ebea]
					}
					_cgbec = append(_cgbec, _fdfda)
				}
				_adgc = append(_adgc, _cgbec)
			}
			return MakeArrayResult(_adgc)
		}
	case 3:
		_cgdbc := _bcgd[1]
		_ebeg := _bcgd[2]
		_fecee := _gadc(_cgdbc)
		_bbc := _gadc(_ebeg)
		if _fecee && _bbc {
			_gfff := []Result{}
			for _, _eegef := range _ffaf {
				var _gbga Result
				if _eegef.ValueNumber == 0 {
					_gbga = _ebeg
				} else {
					_gbga = _cgdbc
				}
				_gfff = append(_gfff, _gbga)
			}
			return MakeListResult(_gfff)
		}
		if _cgdbc.Type != ResultTypeArray && _ebeg.Type != ResultTypeArray {
			_aacb := _gefg(_cgdbc, _gedgb)
			_addc := _gefg(_ebeg, _gedgb)
			_caeg := []Result{}
			for _cfff, _ebff := range _ffaf {
				var _faba Result
				if _ebff.ValueNumber == 0 {
					_faba = _addc[_cfff]
				} else {
					_faba = _aacb[_cfff]
				}
				_caeg = append(_caeg, _faba)
			}
			return MakeListResult(_caeg)
		}
		_dccg, _cfb := len(_cgdbc.ValueArray), len(_ebeg.ValueArray)
		_gcda, _bdddc := _dccg, _cfb
		if _cfb > _gcda {
			_gcda, _bdddc = _bdddc, _gcda
		}
		_ggcag := _efce(_cgdbc, _gcda, _gedgb)
		_bbcf := _efce(_ebeg, _gcda, _gedgb)
		_cggec := [][]Result{}
		for _bgbce := 0; _bgbce < _gcda; _bgbce++ {
			_efgae := []Result{}
			for _gaeea, _ggfd := range _ffaf {
				var _egebg Result
				if _ggfd.ValueNumber == 0 {
					if _bgbce < _cfb {
						_egebg = _bbcf[_bgbce][_gaeea]
					} else {
						_egebg = MakeErrorResultType(ErrorTypeNA, "")
					}
				} else {
					if _bgbce < _dccg {
						_egebg = _ggcag[_bgbce][_gaeea]
					} else {
						_egebg = MakeErrorResultType(ErrorTypeNA, "")
					}
				}
				_efgae = append(_efgae, _egebg)
			}
			_cggec = append(_cggec, _efgae)
		}
		return MakeArrayResult(_cggec)
	}
	return MakeErrorResult("")
}
func _cbea(_dfad, _dcfcg _eb.Time, _agfe int) float64 {
	if _dfad.After(_dcfcg) {
		_dfad, _dcfcg = _dcfcg, _dfad
	}
	_ddea := 0
	_ccfd, _cgb, _cbgb := _dfad.Date()
	_efe, _cbcf, _gbefe := _dcfcg.Date()
	_gaee, _eace := int(_cgb), int(_cbcf)
	_gbcd, _acff := _adgg(_ccfd, _gaee, _cbgb, _agfe), _adgg(_efe, _eace, _gbefe, _agfe)
	if !_gbgg(_agfe) {
		return _bad(_efe, _eace, _acff) - _bad(_ccfd, _gaee, _gbcd)
	}
	if _agfe == 0 {
		if (_gaee == 2 || _gbcd < 30) && _gbefe == 31 {
			_acff = 31
		} else if _eace == 2 && _acff == _ffcg(_efe, _eace) {
			_acff = _ffcg(_efe, 2)
		}
	} else {
		if _gaee == 2 && _gbcd == 30 {
			_gbcd = _ffcg(_ccfd, 2)
		}
		if _eace == 2 && _acff == 30 {
			_acff = _ffcg(_efe, 2)
		}
	}
	if _ccfd < _efe || (_ccfd == _efe && _gaee < _eace) {
		_ddea = 30 - _gbcd + 1
		_cbgb = 1
		_gbcd = 1
		_eba := _eb.Date(_ccfd, _eb.Month(_gaee), _cbgb, 0, 0, 0, 0, _eb.UTC).AddDate(0, 1, 0)
		if _eba.Year() < _efe {
			_ddea += _dad(_eba.Year(), int(_eba.Month()), 12, _agfe)
			_eba = _eba.AddDate(0, 13-int(_eba.Month()), 0)
			_ddea += _bba(_eba.Year(), _efe-1, _agfe)
		}
		_ddea += _dad(_efe, int(_eba.Month()), _eace-1, _agfe)
		_eba = _eba.AddDate(0, _eace-int(_eba.Month()), 0)
		_gaee = _eba.Day()
	}
	_ddea += _acff - _gbcd
	if _ddea > 0 {
		return float64(_ddea)
	} else {
		return 0
	}
}

// LCM implements the Excel LCM() function which returns the least common
// multiple of a range of numbers.
func LCM(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("\u004c\u0043M(\u0029\u0020\u0072e\u0071\u0075\u0069\u0072es \u0061t \u006c\u0065\u0061\u0073\u0074\u0020\u006fne\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074")
	}
	_affb := []float64{}
	for _, _dcfe := range args {
		switch _dcfe.Type {
		case ResultTypeString:
			_deedg := _dcfe.AsNumber()
			if _deedg.Type != ResultTypeNumber {
				return MakeErrorResult("\u004c\u0043M(\u0029\u0020\u006fn\u006c\u0079\u0020\u0061cce\u0070ts\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
			}
			_affb = append(_affb, _deedg.ValueNumber)
		case ResultTypeList:
			_fagg := LCM(_dcfe.ValueList)
			if _fagg.Type != ResultTypeNumber {
				return _fagg
			}
			_affb = append(_affb, _fagg.ValueNumber)
		case ResultTypeNumber:
			_affb = append(_affb, _dcfe.ValueNumber)
		case ResultTypeEmpty:
		case ResultTypeError:
			return _dcfe
		}
	}
	if len(_affb) == 0 {
		return MakeErrorResult("\u004cC\u004d\u0020r\u0065\u0071\u0075\u0069r\u0065\u0073\u0020a\u0074\u0020\u006c\u0065\u0061\u0073\u0074\u0020\u006fne\u0020\u006e\u006fn\u002d\u0065m\u0070\u0074\u0079\u0020\u0061\u0072g\u0075\u006de\u006e\u0074")
	}
	if _affb[0] < 0 {
		return MakeErrorResult("\u004c\u0043M\u0028\u0029\u0020\u006fn\u006c\u0079 \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if len(_affb) == 1 {
		return MakeNumberResult(_affb[0])
	}
	_gdfg := _affb[0]
	for _cgca := 1; _cgca < len(_affb); _cgca++ {
		if _affb[_cgca] < 0 {
			return MakeErrorResult("\u004c\u0043M\u0028\u0029\u0020\u006fn\u006c\u0079 \u0061\u0063\u0063\u0065\u0070\u0074\u0073\u0020p\u006f\u0073\u0069\u0074\u0069\u0076\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
		}
		_gdfg = _bggf(_gdfg, _affb[_cgca])
	}
	return MakeNumberResult(_gdfg)
}

// Not is an implementation of the Excel NOT() function and takes a single
// argument.
func Not(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u004eO\u0054\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065 \u0061\u0072\u0067\u0075\u006d\u0065\u006et")
	}
	switch args[0].Type {
	case ResultTypeError:
		return args[0]
	case ResultTypeString, ResultTypeList:
		return MakeErrorResult("\u004e\u004f\u0054\u0020\u0065\u0078\u0070\u0065\u0063\u0074s\u0020\u0061\u0020\u006e\u0075\u006d\u0065r\u0069\u0063\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	case ResultTypeNumber:
		return MakeBoolResult(!(args[0].ValueNumber != 0))
	default:
		return MakeErrorResult("u\u006e\u0068\u0061\u006e\u0064\u006ce\u0064\u0020\u004e\u004f\u0054\u0020\u0061\u0072\u0067u\u006d\u0065\u006et\u0020t\u0079\u0070\u0065")
	}
}

// Xnpv implements the Excel XNPV function.
func Xnpv(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0058\u004eP\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075\u006den\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("X\u004e\u0050\u0056\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020t\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065r \u0061\u0072\u0067u\u006de\u006e\u0074")
	}
	_dagf := args[0].ValueNumber
	if _dagf <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "\u0058\u004e\u0050\u0056\u0020\u0072e\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0072\u0061\u0074\u0065\u0020\u0074o\u0020\u0062\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u0076\u0065")
	}
	_cddge, _fede := _gcfgb(args[1], args[2], "\u0058\u004e\u0050\u0056")
	if _fede.Type == ResultTypeError {
		return _fede
	}
	_daedf := _cddge._faaf
	_fcdg := _cddge._ceee
	_ffad := 0.0
	_eeff := _fcdg[0]
	for _beec, _bfda := range _daedf {
		_ffad += _bfda / _cc.Pow(1+_dagf, (_fcdg[_beec]-_eeff)/365)
	}
	return MakeNumberResult(_ffad)
}
func _cgc() {
	_fce["\u006d\u006d\u002f\u0064\u0064\u002f\u0079\u0079"] = _f.MustCompile("\u005e" + _acde + _bbg)
	_fce["\u006dm\u0020\u0064\u0064\u002c\u0020\u0079y"] = _f.MustCompile("\u005e" + _bed + _bbg)
	_fce["\u0079\u0079\u002d\u006d\u006d\u002d\u0064\u0064"] = _f.MustCompile("\u005e" + _gbg + _bbg)
	_fce["y\u0079\u002d\u006d\u006d\u0053\u0074\u0072\u002d\u0064\u0064"] = _f.MustCompile("\u005e" + _faef + _bbg)
	_fbgc["\u0068\u0068"] = _f.MustCompile(_adb + _ceb + "\u0024")
	_fbgc["\u0068\u0068\u003am\u006d"] = _f.MustCompile(_adb + _cfc + "\u0024")
	_fbgc["\u006d\u006d\u003as\u0073"] = _f.MustCompile(_adb + _dd + "\u0024")
	_fbgc["\u0068\u0068\u003a\u006d\u006d\u003a\u0073\u0073"] = _f.MustCompile(_adb + _fed + "\u0024")
	_afd = []*_f.Regexp{_f.MustCompile("\u005e" + _acde + "\u0024"), _f.MustCompile("\u005e" + _bed + "\u0024"), _f.MustCompile("\u005e" + _gbg + "\u0024"), _f.MustCompile("\u005e" + _faef + "\u0024")}
	_dag = []*_f.Regexp{_f.MustCompile("\u005e" + _ceb + "\u0024"), _f.MustCompile("\u005e" + _cfc + "\u0024"), _f.MustCompile("\u005e" + _dd + "\u0024"), _f.MustCompile("\u005e" + _fed + "\u0024")}
}
func (_ffeg *Lexer) nextRaw() *node {
	for len(_ffeg._abeff) != 0 {
		_dbdc := <-_ffeg._abeff[len(_ffeg._abeff)-1]
		if _dbdc != nil {
			return _dbdc
		}
		_ffeg._abeff = _ffeg._abeff[0 : len(_ffeg._abeff)-1]
	}
	return <-_ffeg._cgdbfa
}

// Sum is an implementation of the Excel SUM() function.
func Sum(args []Result) Result {
	_afad := MakeNumberResult(0)
	for _, _babd := range args {
		_babd = _babd.AsNumber()
		switch _babd.Type {
		case ResultTypeNumber:
			_afad.ValueNumber += _babd.ValueNumber
		case ResultTypeList, ResultTypeArray:
			_fecaf := Sum(_babd.ListValues())
			if _fecaf.Type != ResultTypeNumber {
				return _fecaf
			}
			_afad.ValueNumber += _fecaf.ValueNumber
		case ResultTypeString:
		case ResultTypeError:
			return _babd
		case ResultTypeEmpty:
		default:
			return MakeErrorResult(_e.Sprintf("\u0075\u006e\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u0053\u0055\u004d\u0028\u0029 \u0061r\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0073", _babd.Type))
		}
	}
	return _afad
}

// Eval evaluates a horizontal range returning a list of results or an error.
func (_gced HorizontalRange) Eval(ctx Context, ev Evaluator) Result {
	_baea := _gced.horizontalRangeReference()
	if _degaa, _cfddc := ev.GetFromCache(_baea); _cfddc {
		return _degaa
	}
	_feaf, _gged := _gdbdc(ctx, _gced._fcbd, _gced._fbcg)
	_dcfg := _gebc(ctx, ev, _feaf, _gged)
	ev.SetCache(_baea, _dcfg)
	return _dcfg
}

// ConstArrayExpr is a constant array expression.
type ConstArrayExpr struct{ _ba [][]Expression }

// Reference returns a string reference value to a vertical range with prefix.
func (_bgcb PrefixVerticalRange) Reference(ctx Context, ev Evaluator) Reference {
	_fdbdd := _bgcb._fcga.Reference(ctx, ev)
	return Reference{Type: ReferenceTypeVerticalRange, Value: _bgcb.verticalRangeReference(_fdbdd.Value)}
}
func _gfcg(_ecfg, _egbf Result, _feba string) (float64, float64, Result) {
	_dace, _ecfbb := _eec(_ecfg, "\u0073e\u0074t\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065", _feba)
	if _ecfbb.Type == ResultTypeError {
		return 0, 0, _ecfbb
	}
	_beda, _ecfbb := _eec(_egbf, "\u006d\u0061\u0074\u0075\u0072\u0069\u0074\u0079\u0020\u0064\u0061\u0074\u0065", _feba)
	if _ecfbb.Type == ResultTypeError {
		return 0, 0, _ecfbb
	}
	if _dace >= _beda {
		return 0, 0, MakeErrorResultType(ErrorTypeNum, _feba+"\u0020\u0072\u0065\u0071\u0075\u0069r\u0065\u0073\u0020m\u0061\u0074\u0075r\u0069\u0074\u0079\u0020\u0064\u0061\u0074\u0065\u0020\u0074o\u0020\u0062\u0065\u0020\u006cat\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0073\u0065\u0074\u0074\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u0064\u0061\u0074\u0065")
	}
	return _dace, _beda, _feb
}
func (_cdbacd *ivr) NamedRange(ref string) Reference { return ReferenceInvalid }
func _edeef(_bcdff, _ecaed Result, _afbb, _dgaa bool) cmpResult {
	_bcdff = _bcdff.AsNumber()
	_ecaed = _ecaed.AsNumber()
	if _bcdff.Type != _ecaed.Type {
		return _bfaf
	}
	if _bcdff.Type == ResultTypeNumber {
		if _bcdff.ValueNumber == _ecaed.ValueNumber {
			return _cgae
		}
		if _bcdff.ValueNumber < _ecaed.ValueNumber {
			return _gcefb
		}
		return _acbcc
	}
	if _bcdff.Type == ResultTypeString {
		_bgdg := _bcdff.ValueString
		_abbag := _ecaed.ValueString
		if !_afbb {
			_bgdg = _ga.ToLower(_bgdg)
			_abbag = _ga.ToLower(_abbag)
		}
		if _dgaa {
			_eacg := _cd.Match(_abbag, _bgdg)
			if _eacg {
				return _cgae
			} else {
				return _acbcc
			}
		}
		return cmpResult(_ga.Compare(_bgdg, _abbag))
	}
	if _bcdff.Type == ResultTypeEmpty {
		return _cgae
	}
	if _bcdff.Type == ResultTypeList {
		if len(_bcdff.ValueList) < len(_ecaed.ValueList) {
			return _gcefb
		}
		if len(_bcdff.ValueList) > len(_ecaed.ValueList) {
			return _acbcc
		}
		for _eaff := range _bcdff.ValueList {
			_cfad := _edeef(_bcdff.ValueList[_eaff], _ecaed.ValueList[_eaff], _afbb, _dgaa)
			if _cfad != _cgae {
				return _cfad
			}
		}
		return _cgae
	}
	if _bcdff.Type == ResultTypeList {
		if len(_bcdff.ValueArray) < len(_ecaed.ValueArray) {
			return _gcefb
		}
		if len(_bcdff.ValueArray) > len(_ecaed.ValueArray) {
			return _acbcc
		}
		for _fabbf := range _bcdff.ValueArray {
			_bdcfg := _bcdff.ValueArray[_fabbf]
			_fcfe := _bcdff.ValueArray[_fabbf]
			if len(_bdcfg) < len(_fcfe) {
				return _gcefb
			}
			if len(_bdcfg) > len(_fcfe) {
				return _acbcc
			}
			for _gcce := range _bdcfg {
				_ebfd := _edeef(_bdcfg[_gcce], _fcfe[_gcce], _afbb, _dgaa)
				if _ebfd != _cgae {
					return _ebfd
				}
			}
		}
		return _cgae
	}
	return _bfaf
}

// String returns a string representation for Negate.
func (_aagda Negate) String() string { return "\u002d" + _aagda._effb.String() }

const _cebaa = 57367

// Mod is an implementation of the Excel MOD function which returns the
// remainder after division. It requires two numeric argumnts.
func Mod(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("\u004d\u004fD(\u0029\u0020\u0072e\u0071\u0075\u0069\u0072es \u0074wo\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	_eaeg := args[0].AsNumber()
	_bfef := args[1].AsNumber()
	if _eaeg.Type != ResultTypeNumber || _bfef.Type != ResultTypeNumber {
		return MakeErrorResult("\u004d\u004fD(\u0029\u0020\u0072e\u0071\u0075\u0069\u0072es \u0074wo\u0020\u006e\u0075\u006d\u0065\u0072\u0069c \u0061\u0072\u0067\u0075\u006d\u0065\u006et\u0073")
	}
	if _bfef.ValueNumber == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "M\u004fD\u0028\u0029\u0020\u0064\u0069\u0076\u0069\u0064e\u0020\u0062\u0079\u0020ze\u0072\u006f")
	}
	_bgdf, _baac := _cc.Modf(_eaeg.ValueNumber / _bfef.ValueNumber)
	if _baac < 0 {
		_bgdf--
	}
	return MakeNumberResult(_eaeg.ValueNumber - _bfef.ValueNumber*_bgdf)
}

// RegisterFunction registers a standard function.
func RegisterFunction(name string, fn Function) {
	_agde.Lock()
	defer _agde.Unlock()
	if _, _aagdg := _bdgeg[name]; _aagdg {
		_cf.Log.Debug("\u0064\u0075p\u006c\u0069\u0063\u0061t\u0065\u0020r\u0065\u0067\u0069\u0073\u0074\u0072\u0061\u0074i\u006f\u006e\u0020\u006f\u0066\u0020\u0066\u0075\u006e\u0063\u0074\u0069o\u006e\u0020\u0025\u0073", name)
	}
	_bdgeg[name] = fn
}
func _cccg() {
	_ggde = _f.MustCompile("\u005e\u0030\u002b\u0024")
	_cgfbe = _f.MustCompile("\u005e\u0028\u0028\u0023|0\u0029\u002b\u002c\u0029\u002b\u0028\u0023\u007c\u0030\u0029\u002b\u0028\u003b\u007c$\u0029")
	_aebbb = _f.MustCompile("\u005e\u0028\u0023\u007c\u0030\u007c\u002c\u0029\u002a\u005f\u005c\u0029\u003b")
	_fgbb = _f.MustCompile("\u005e\u0030\u002b\u005c\u002e\u0028\u0030\u002b\u0029\u0024")
	_cgbg = _f.MustCompile("\u005e\u0028\u0028\u0023\u007c\u0030\u0029\u002b\u002c\u0029+\u0028\u0023\u007c\u0030\u0029\u002b\u005c.\u0028\u0030\u002b\u0029\u002e\u002a\u0028\u003b\u007c\u0024\u0029")
	_geabe = _f.MustCompile("^\u0028\u005f\u007c\u002d\u007c\u0020)\u002b\u005c\u002a\u0020\u0023\u002b\u002c\u0023\u002b0\u005c\u002e\u00280\u002b)\u002e\u002a\u003b")
	_bfggb = _f.MustCompile("\u005e\u0028\u0028\u0023\u007c\u0030)\u002b\u002c\u0029\u002b\u0028\u0023\u007c\u0030\u0029\u002b\u005c\u002e\u0028(\u0023\u007c\u0030\u0029\u002b\u0029\u005f\\\u0029\u002e\u002a\u003b")
	_fbf = _f.MustCompile("\u005e\u0028\u0023\u007c0)\u002b\u005c\u002e\u0028\u0028\u0023\u007c\u0030\u0029\u002b\u0029\u0025\u0024")
	_aaed = _f.MustCompile("\u005c\u005b\u005c$\u005c\u0024\u002d\u002e+\u005c\u005d\u0028\u005c\u002a\u0020\u0029?\u0028\u0023\u007c\u0030\u0029\u002b\u002c\u0028\u0023\u007c\u0030\u0029\u002b\u003b")
	_gaab = _f.MustCompile("\u005c[\u005c\u0024\\\u0024\u002d\u002e+\u005c\u005d\u0028\u005c\u002a\u0020\u0029?\u0028\u0023\u007c\u0030\u0029\u002b,\u0028\u0023\u007c\u0030\u0029\u002b\u005c\u002e\u0028\u0028\u0023|\u0030\u007c\u002d\u0029\u002b\u0029\u002e\u002a\u003b")
	_ceef = _f.MustCompile("\u005e(\u0028\u0023|\u0030\u0029\u002b,\u0029\u002b\u0028\u0023\u007c\u0030\u0029+\u0028\u005c\u002e\u0028\u0028\u0023|\u0030\u007c\u002d\u0029\u002b\u0029\u0029\u003f\u002e\u002b\u005c[\u005c\u0024\u002e\u002b\u005c\u005d\u002e\u002a\u003b")
	_befb = _f.MustCompile("\u005e\u004d\u002b(\u002f\u007c\u0020\u007c\u002c\u007c\u0022\u007c" + _dfge + _dfge + "\u0029\u002b\u0044\u002b\u0028\u002f\u007c\u0020\u007c\u002c\u007c\u0022\u007c" + _dfge + _dfge + "\u0029\u002b\u0059+\u0024")
	_fffd = _f.MustCompile("\u005e\u0044\u002b\u0028\u002f\u007c\u0020\u007c\u005c\u002e\u007c\u0022\u007c" + _dfge + _dfge + "\u0029\u002b\u004d\u002b\u0028\u002f\u007c\u0020\u007c\\\u002e\u007c\u0022\u007c" + _dfge + _dfge + "\u0029\u002b\u0059+\u0024")
	_cdbe = _f.MustCompile("\u005e\u0028\u0023|\u0030\u0029\u002b\u005c.\u0028\u0028\u0023\u007c\u0030\u0029\u002a)\u0045\u005c\u002b\u0028\u0023\u007c\u0030\u0029\u002b\u0028\u003b\u007c\u0024\u0029")
	_gbcbc = _f.MustCompile("\u005e.\u002a\u005f\u005c\u0029\u002e\u002a;")
}
func _bgdeg(_feef, _bfceb []string) []string {
	for _, _agab := range _bfceb {
		_feef = append(_feef, _agab)
	}
	return _feef
}

// Column implements the Excel COLUMN function.
func Column(args []Result) Result {
	if len(args) < 1 {
		return MakeErrorResult("\u0043\u004f\u004c\u0055M\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073 \u006fn\u0065\u0020\u0061\u0072\u0067\u0075\u006de\u006e\u0074")
	}
	_ffeec := args[0].Ref
	if _ffeec.Type != ReferenceTypeCell {
		return MakeErrorResult("\u0043\u004f\u004c\u0055\u004dN\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u006e\u0020a\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063e")
	}
	_bdfb, _cffed := _fb.ParseCellReference(_ffeec.Value)
	if _cffed != nil {
		return MakeErrorResult("I\u006e\u0063\u006f\u0072re\u0063t\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065\u003a\u0020" + _ffeec.Value)
	}
	return MakeNumberResult(float64(_bdfb.ColumnIdx + 1))
}

// ISFORMULA is an implementation of the Excel ISFORMULA() function.
func IsFormula(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("\u0049\u0053F\u004f\u0052\u004d\u0055L\u0041\u0028)\u0020\u0061\u0063\u0063\u0065\u0070\u0074\u0073 \u0061\u0020\u0073\u0069\u006e\u0067\u006c\u0065\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_afdb := args[0].Ref
	if _afdb.Type != ReferenceTypeCell {
		return MakeErrorResult("I\u0053\u0046\u004f\u0052\u004d\u0055\u004c\u0041\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u0072\u0073t\u0020a\u0072\u0067\u0075\u006de\u006e\u0074 \u0074\u006f\u0020\u0062\u0065\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065")
	}
	return MakeBoolResult(ctx.HasFormula(_afdb.Value))
}

const _acde = "\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029\u002f\u0028\u0028\u005b\u0030-\u0039]\u0029\u002b\u0029\u002f\u0028\u0028\u005b\u0030\u002d\u0039\u005d\u0029\u002b\u0029"

func _gc(_cfg, _ec [][]Result) bool {
	if len(_cfg) != len(_ec) {
		return false
	}
	for _aa := range _cfg {
		if len(_cfg[_aa]) != len(_ec[_aa]) {
			return false
		}
	}
	return true
}

var _eefb = [...]int{0, 0, 71, 70, 69, 4, 67, 66, 53, 51, 50, 49, 48, 47, 46, 45, 44, 2}

// MakeNumberResult constructs a number result.
func MakeNumberResult(v float64) Result {
	if v == _cc.Copysign(0, -1) {
		v = 0
	}
	return Result{Type: ResultTypeNumber, ValueNumber: v}
}

// Sln implements the Excel SLN function.
func Sln(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("\u0053\u004c\u004e\u0020r\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0068r\u0065e\u0020\u0061\u0072\u0067\u0075\u006d\u0065n\u0074\u0073")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u004c\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020c\u006f\u0073\u0074\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_deba := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u004c\u004e \u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0061\u006c\u0076\u0061\u0067\u0065\u0020\u0074\u006f\u0020\u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072 \u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074")
	}
	_fbgd := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("\u0053\u004c\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020l\u0069\u0066\u0065\u0020\u0074\u006f \u0062\u0065\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0061\u0072\u0067u\u006d\u0065\u006e\u0074")
	}
	_abee := args[2].ValueNumber
	if _abee == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "\u0053\u004c\u004e\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u006c\u0069f\u0065 \u0074\u006f\u0020\u0062\u0065\u0020\u006e\u006f\u006e\u0020\u007a\u0065\u0072\u006f")
	}
	return MakeNumberResult((_deba - _fbgd) / _abee)
}

// Lookup implements the LOOKUP function that returns a matching value from a
// column, or from the same index in a second column.
func Lookup(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0061\u0074\u0020\u006c\u0065\u0061\u0073\u0074 \u0074\u0077\u006f\u0020\u0061r\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	if len(args) > 3 {
		return MakeErrorResult("\u004c\u004f\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072e\u0073\u0020\u0061\u0074\u0020\u006do\u0073\u0074\u0020\u0074\u0068\u0072\u0065\u0065\u0020\u0061\u0072\u0067\u0075m\u0065\u006e\u0074\u0073")
	}
	_fdff := args[0]
	_cfae := args[1]
	if _cfae.Type != ResultTypeArray && _cfae.Type != ResultTypeList {
		return MakeErrorResult("\u0056\u004cO\u004f\u004b\u0055\u0050\u0020\u0072\u0065\u0071\u0075\u0069\u0072\u0065\u0073\u0020\u0073\u0065\u0063\u006f\u006e\u0064\u0020\u0061\u0072\u0067\u0075\u006d\u0065\u006e\u0074\u0020\u006f\u0066\u0020\u0074\u0079\u0070\u0065\u0020\u0061\u0072\u0072\u0061\u0079")
	}
	_ebdf := _ccced(_cfae)
	_efbc := -1
	for _aeda, _efcg := range _ebdf {
		if _edeef(_fdff, _efcg, false, false) == _cgae {
			_efbc = _aeda
		}
	}
	if _efbc == -1 {
		return MakeErrorResultType(ErrorTypeNA, "\u004c\u004f\u004f\u004bUP\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	_bgeb := _ebdf
	if len(args) == 3 {
		_bgeb = _ccced(args[2])
	}
	if _efbc < 0 || _efbc >= len(_bgeb) {
		return MakeErrorResultType(ErrorTypeNA, "\u004c\u004f\u004f\u004bUP\u0020\u006e\u006f\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0066\u006f\u0075n\u0064")
	}
	return _bgeb[_efbc]
}
func (_ebdaf *ivr) Cell(ref string, ev Evaluator) Result {
	return MakeErrorResult("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0066\u0065r\u0065\u006e\u0063\u0065")
}

// GetLabelPrefix returns an empty string for the invalid reference context.
func (_aaac *ivr) GetLabelPrefix(cellRef string) string { return "" }

// IfNA is an implementation of the Excel IFNA() function. It takes two arguments.
func IfNA(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("I\u0046\u004e\u0041\u0020\u0072\u0065q\u0075\u0069\u0072\u0065\u0073\u0020\u0074\u0077\u006f \u0061\u0072\u0067u\u006de\u006e\u0074\u0073")
	}
	if args[0].Type == ResultTypeError && args[0].ValueString == "\u0023\u004e\u002f\u0041" {
		return args[1]
	}
	return args[0]
}

type parsedReplaceObject struct {
	_bdfaf string
	_bebed int
	_ecadb int
	_cbegb string
}
