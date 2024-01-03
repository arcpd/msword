package testutils

import (
	_ea "crypto/md5"
	_ag "encoding/hex"
	_c "encoding/json"
	_bf "errors"
	_ad "fmt"
	_fd "github.com/stretchr/testify/require"
	_de "golang.org/x/image/font"
	_agc "golang.org/x/image/font/opentype"
	_fe "golang.org/x/image/math/fixed"
	_gc "image"
	_a "image/color"
	_ca "image/draw"
	_ab "image/png"
	_gg "io"
	_e "io/ioutil"
	_fg "math"
	_ee "os"
	_f "os/exec"
	_ed "path/filepath"
	_d "strings"
	_age "sync"
	_b "testing"
	_ge "time"
)

func (_dac *ReferenceMap) Keys() (_cd []string) {
	_cd = make([]string, len(_dac._bff))
	var _gf int
	for _ff := range _dac._bff {
		_cd[_gf] = _ff
		_gf++
	}
	return _cd
}

func CopyFile(src, dst string) error {
	_agd, _aee := _ee.Open(src)
	if _aee != nil {
		return _aee
	}
	defer _agd.Close()
	_gfe, _aee := _ee.Create(dst)
	if _aee != nil {
		return _aee
	}
	defer _gfe.Close()
	_, _aee = _gg.Copy(_gfe, _agd)
	return _aee
}

func (_dc *ReferenceMap) Len() int { return len(_dc._bff) }

func (_be *ReferenceMap) UnmarshalJSON(data []byte) error { return _c.Unmarshal(data, &_be._bff) }

func CompareImages(img1, img2 _gc.Image) (bool, error) {
	_bb := img1.Bounds()
	_edd := 0
	for _cde := 0; _cde < _bb.Size().X; _cde++ {
		for _dab := 0; _dab < _bb.Size().Y; _dab++ {
			_gcg, _gdc, _aaa, _ := img1.At(_cde, _dab).RGBA()
			_fb, _ebc, _gbd, _ := img2.At(_cde, _dab).RGBA()
			if _gcg != _fb || _gdc != _ebc || _aaa != _gbd {
				_edd++
			}
		}
	}
	_fcc := float64(_edd) / float64(_bb.Dx()*_bb.Dy())
	if _fcc > 0.0001 {
		_ad.Printf("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a", _fcc, _edd)
		return false, nil
	}
	return true, nil
}

func RunRenderTest(t *_b.T, pdfPath, outputDir, baselineRenderPath string, saveBaseline bool, currentHashMap *ReferenceMap, refFile *ReferenceFile) {
	_eeec := _d.TrimSuffix(_ed.Base(pdfPath), _ed.Ext(pdfPath))
	t.Run("\u0072\u0065\u006e\u0064\u0065\u0072", func(_fcff *_b.T) {
		_ebe := _ed.Join(outputDir, _eeec)
		_beb := _ebe + "\u002d%\u0064\u002e\u0070\u006e\u0067"
		if _bbf := RenderPDFToPNGs(pdfPath, 0, _beb); _bbf != nil {
			_fcff.Skip(_bbf)
		}
		_fdb := _eeec + "\u002em\u0073\u0077\u006f\u0072\u0064"
		_fda := _ed.Join(outputDir, _fdb)
		_dfa := _fda + "\u002d%\u0064\u002e\u0070\u006e\u0067"
		_dedbf := false
		if saveBaseline {
			_aec := _ed.Dir(pdfPath)
			_fdf := _ed.Join(_aec, _fdb+"\u002e\u0070\u0064\u0066")
			if _, _ccbc := _ee.Stat(_fdf); _ccbc == nil {
				_fcff.Logf("\u0052\u0065\u006e\u0064er\u0020\u004d\u0053\u0020\u0057\u006f\u0072\u0064\u0020\u0050\u0044\u0046\u003a\u0020%\u0076", _fdf)
				if _aac := RenderPDFToPNGs(_fdf, 0, _dfa); _aac != nil {
					_fcff.Skip(_aac)
				}
				_dedbf = true
			}
		}
		for _edab := 1; true; _edab++ {
			_bdb := _ad.Sprintf("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg", _ebe, _edab)
			_eeb := _ed.Join(baselineRenderPath, _ad.Sprintf("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067", _eeec, _edab))
			if _, _abf := _ee.Stat(_bdb); _abf != nil {
				break
			}
			_fcff.Logf("\u0025\u0073", _eeb)
			if saveBaseline {
				_fcff.Logf("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073", _bdb, _eeb)
				_bbfc := CopyFile(_bdb, _eeb)
				if _bbfc != nil {
					_fcff.Fatalf("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _eeb, _bbfc)
				}
				if _dedbf {
					_fgf := _ad.Sprintf("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg", _fda, _edab)
					_bccg := _ed.Join(baselineRenderPath, _ad.Sprintf("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067", _fdb, _edab))
					_fcff.Logf("\u0043\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u004d\u0053\u0020\u0057\u006f\u0072\u0064 \u0072e\u0073\u0075\u006c\u0074\u0073\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073", _fgf, _bccg)
					_afe := CopyFile(_fgf, _bccg)
					if _afe != nil {
						_fcff.Logf("\u0045\u0052RO\u0052\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076\u002c \u0068\u0061\u0076\u0069n\u0067\u0020\u0064\u0069ff\u0065\u0072e\u006et\u0020\u0070\u0061\u0067\u0065\u0020s\u0069\u007a\u0065 \u0072\u0065s\u0075\u006c\u0074\u0073\u002c\u0020\u0075\u0073\u0065\u0020\u0070\u0072\u0065\u0076i\u006f\u0075\u0073\u0020\u0070\u0061\u0067\u0065", _bccg, _afe)
						_fgf = _ad.Sprintf("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg", _fda, _edab-1)
						_bccg = _ed.Join(baselineRenderPath, _ad.Sprintf("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067", _fdb, _edab-1))
						if _fge := CopyFile(_fgf, _bccg); _fge != nil {
							_fcff.Fatalf("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076", _bccg, _fge)
						}
					}
					_fcff.Logf("\u0043\u006f\u006d\u0062\u0069\u006e\u0069\u006e\u0067\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063e\u0020\u0072\u0065\u0073\u0075\u006c\u0074s\u0020\u0077\u0069\u0074\u0068\u0020\u004d\u0053\u0020\u0057\u006fr\u0064\u0020\u0025\u0073\u0020\u0061\u006e\u0064\u0020\u0025\u0073", _eeb, _bccg)
					_cdf, _afe := CombinePNGFiles(_eeb, _bccg)
					if _ee.IsNotExist(_afe) {
						_fcff.Fatal("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
					} else if !_cdf {
						_fcff.Fatal("\u0055n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063\u006f\u006db\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0073")
					}
					_fcff.Logf("Cr\u0065\u0061t\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065 \u0064\u0069\u0066\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0073\u0020\u0077\u0069\u0074\u0068\u0020M\u0053 \u0057\u006f\u0072\u0064\u0020\u0025\u0073\u0020a\u006ed\u0020\u0025s", _eeb, _bccg)
					_cdf, _fag, _dagc, _beba, _afe := CreatePNGDiff(_eeb, _bccg)
					if _afe != nil && _afe != ErrImageSizeNotMatch {
						_fcff.Fatalf("\u0045\u0072\u0072\u006fr\u0020\u006f\u006e\u0020\u0063\u0072\u0065\u0061\u0074\u0065 \u0050N\u0047\u0020\u0044\u0069\u0066\u0066\u003a \u0025\u0076", _afe)
					}
					if _cdf {
						_fcff.Logf("\u0049\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073\u000a", _fag)
						_fcff.Logf("D\u0069\u0066\u0066\u0020Pe\u0072c\u0065\u006e\u0074\u003a\u0020%\u0032\u002e\u0066\u0025\u0025\u000a", _dagc)
						_fcff.Logf("\u0044i\u0066f\u0020\u0054\u006f\u0074\u0061\u006c\u003a\u0020\u0025\u0066\u000a", _beba)
						_bac := _ed.Base(_fag)
						if _ecb, _ggc := currentHashMap.Read(_bac); _ggc {
							if _ecb.DiffPercent < _dagc || _ecb.DiffTotal < _beba {
								_fcff.Fatalf("\u004e\u0065\u0077\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0073\u0020\u0068\u0061\u0076\u0069\u006e\u0067 h\u0069g\u0068\u0065\u0072\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0070\u0065\u0072\u0063\u0065\u006e\u0074\u003a\u0020\u0025\u0066\u0020\u0061\u006e\u0064 \u0074\u006f\u0074\u0061\u006c\u0020\u0025\u0066\u000a", _dagc, _beba)
							}
						}
						_ccc, _eaaa := HashFile(_fag)
						_fd.NoError(_fcff, _eaaa)
						_daga := ReferenceEntry{Timestamp: _ge.Now().UTC().Unix(), Value: _ccc, ResultSize: _aace(_fcff, _fag), DiffPercent: _dagc, DiffTotal: _beba}
						currentHashMap.Write(_bac, _daga)
						if _eaaa = refFile.Update(currentHashMap); _eaaa != nil {
							_fcff.Logf("\u0055\u0070\u0064\u0061\u0074\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u0020f\u0069\u006c\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076", _eaaa)
						}
					}
				}
				continue
			}
			_fcff.Run(_ad.Sprintf("\u0070\u0061\u0067\u0065\u0025\u0064", _edab), func(_ffbf *_b.T) {
				_ffbf.Logf("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073", _bdb, _eeb)
				_cef, _eadf := ComparePNGFiles(_bdb, _eeb)
				if _ee.IsNotExist(_eadf) {
					_ffbf.Fatal("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067")
				} else if !_cef {
					_ffbf.Fatal("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064")
				}
			})
		}
	})
}

func (_fcb *ReferenceMap) Read(key string) (ReferenceEntry, bool) {
	_fcb.Lock()
	defer _fcb.Unlock()
	if _fcb._bff == nil {
		return ReferenceEntry{}, false
	}
	_cc, _bgd := _fcb._bff[key]
	return _cc, _bgd
}

func (_add *ReferenceMap) Copy() *ReferenceMap {
	_fgg := ReferenceMap{_bff: make(map[string]ReferenceEntry, len(_add._bff))}
	for _af, _ddf := range _add._bff {
		_fgg._bff[_af] = _ddf
	}
	return &_fgg
}

func ReadFile(dirPath, testName string, createEmpty bool) (*ReferenceFile, error) {
	if dirPath == "" && createEmpty {
		return &ReferenceFile{Map: &ReferenceMap{}}, nil
	}
	if dirPath == "" {
		return nil, _ee.ErrNotExist
	}
	_dae := _ed.Join(dirPath, testName+"\u005fr\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u002e\u006a\u0073\u006f\u006e")
	_def := &ReferenceFile{_abe: _dae}
	_aa, _bd := _ee.Open(_dae)
	if _bf.Is(_bd, _ee.ErrNotExist) && createEmpty {
		_def.Timestamp = _ge.Now().UTC()
		_def.Map = &ReferenceMap{}
		return _def, nil
	}
	if _bd != nil {
		return nil, _bd
	}
	defer _aa.Close()
	if _ccb := _c.NewDecoder(_aa).Decode(_def); _ccb != nil {
		if _ccb.Error() == "i\u006c\u006c\u0065\u0067\u0061\u006c \u0062\u0061\u0073\u0065\u0036\u0034 \u0064\u0061\u0074\u0061\u0020\u0061\u0074 \u0069\u006e\u0070\u0075\u0074\u0020\u0062\u0079\u0074\u0065 \u0030" && createEmpty {
			return _def, nil
		}
		return nil, _ccb
	}
	return _def, nil
}

func (_eaa *ReferenceFile) updateMap(_bg *ReferenceMap) int {
	var _cf int
	if _eaa.Map._bff == nil {
		_eaa.Map._bff = map[string]ReferenceEntry{}
	}
	for _dd, _ce := range _bg._bff {
		_dba, _fc := _eaa.Map._bff[_dd]
		if !_fc {
			_eaa.Map._bff[_dd] = _ce
			_cf++
			continue
		}
		if string(_dba.Value) != string(_ce.Value) {
			_eaa.Map._bff[_dd] = _ce
			_cf++
		}
	}
	for _ae := range _eaa.Map._bff {
		if _, _ga := _bg._bff[_ae]; !_ga {
			delete(_eaa.Map._bff, _ae)
			_cf++
		}
	}
	return _cf
}

type ReferenceFile struct {
	Timestamp _ge.Time      `json:"timestamp"`
	Map       *ReferenceMap `json:"map,omitempty"`
	_abe      string
}

func CombinePNGFiles(file1, file2 string) (bool, error) {
	_gedb, _eg := ReadPNG(file1)
	if _eg != nil {
		return false, _eg
	}
	_cg, _eg := ReadPNG(file2)
	if _eg != nil {
		return false, _eg
	}
	_adc := _gc.Point{_gedb.Bounds().Dx(), 0}
	_cgd := _gc.Rectangle{_adc, _adc.Add(_cg.Bounds().Size())}
	_cdb := _gc.Rectangle{_gc.Point{0, 0}, _cgd.Max}
	_cdbe := _gc.NewRGBA(_cdb)
	_ca.Draw(_cdbe, _gedb.Bounds(), _gedb, _gc.Point{0, 0}, _ca.Src)
	_ca.Draw(_cdbe, _cgd, _cg, _gc.Point{0, 0}, _ca.Src)
	_bge := _ed.Dir(file1)
	_fgc := _d.TrimSuffix(_ed.Base(file1), _ed.Ext(file1))
	_dace, _eg := _ee.Create(_bge + "\u002f" + _fgc + "\u002d\u0063\u006f\u006d\u0062\u0069\u006e\u0065\u0064\u002e\u0070\u006e\u0067")
	if _eg != nil {
		return false, _eg
	}
	defer _dace.Close()
	if _fga := _ab.Encode(_dace, _cdbe); _fga != nil {
		return false, _fga
	}
	return true, nil
}

func _aace(_gff *_b.T, _dadb string) int64 {
	_bab, _ecg := _ee.Stat(_dadb)
	_fd.NoError(_gff, _ecg)
	return _bab.Size()
}

func CreatePNGDiff(src, dst string) (bool, string, float64, float64, error) {
	_ef, _ccd := ReadPNG(src)
	if _ccd != nil {
		return false, "", 0, 0, _ccd
	}
	_cfg, _ccd := ReadPNG(dst)
	if _ccd != nil {
		return false, "", 0, 0, _ccd
	}
	_agg := _ef.Bounds()
	_bfg := _cfg.Bounds()
	if !_eda(_agg, _bfg) {
		return false, "", 0, 0, ErrImageSizeNotMatch
	}
	_ffb := _gc.NewRGBA(_gc.Rect(0, 0, _agg.Max.X, _agg.Max.Y))
	var (
		_dbd float64
		_ac  float64
	)
	for _ffbe := _agg.Min.Y; _ffbe < _agg.Max.Y; _ffbe++ {
		for _fa := _agg.Min.X; _fa < _agg.Max.X; _fa++ {
			_dcc, _ddg, _ace, _dedb := _cfg.At(_fa, _ffbe).RGBA()
			_ffb.Set(_fa, _ffbe, _a.RGBA{uint8(_dcc), uint8(_ddg), uint8(_ace), 64})
			_ddb, _cff, _feb, _gfd := _ef.At(_fa, _ffbe).RGBA()
			if !_ebf(_ef.At(_fa, _ffbe), _cfg.At(_fa, _ffbe)) {
				_ffb.Set(_fa, _ffbe, _a.RGBA{uint8(_ddb), uint8(_cff), uint8(_feb), uint8(_gfd)})
				_eee := float64(_ddb) + float64(_cff) + float64(_feb) + float64(_gfd) - float64(_dcc) + float64(_ddg) + float64(_ace) + float64(_dedb)
				_adf := _fg.Sqrt(_fg.Pow(_eee/float64(_agg.Dx()*_agg.Dy()), 2))
				_ac += _adf
				_dbd++
			}
		}
	}
	_cffa := _dbd / float64(_agg.Dx()*_agg.Dy()) * 100
	_ffc := _ed.Dir(src)
	_dccb := _d.TrimSuffix(_ed.Base(src), _ed.Ext(src))
	_edc, _ccd := _ee.Create(_ffc + "\u002f" + _dccb + "\u002dd\u0069\u0066\u0066\u002e\u0070\u006eg")
	if _ccd != nil {
		return false, "", 0, 0, _ccd
	}
	defer _edc.Close()
	_gfeg := _d.Replace(_ffc, "\u0072\u0065\u006e\u0064\u0065\u0072", "\u0066\u006f\u006et\u0073", 1) + "\u002f\u0043\u0061l\u0069\u0062\u0072\u0069\u002e\u0074\u0074\u0066"
	_cec := _ad.Sprintf("\u0044\u0069f\u0066\u0065\u0072e\u006e\u0063\u0065\u003a\u0020\u0025\u0066\u0025\u0025", _cffa)
	_ccd = _bfe(_ffb, _gfeg, _cec, 15, 22)
	if _ccd != nil {
		return false, "", 0, 0, _ccd
	}
	_cec = _ad.Sprintf("T\u006ft\u0061\u006c\u0020\u0044\u0069\u0066\u0066\u0065r\u0065\u006e\u0063\u0065: \u0025\u0066", _ac)
	_ccd = _bfe(_ffb, _gfeg, _cec, 15, 44)
	if _ccd != nil {
		return false, "", 0, 0, _ccd
	}
	if _eag := _ab.Encode(_edc, _ffb); _eag != nil {
		return false, "", 0, 0, _eag
	}
	return true, _edc.Name(), _cffa, _ac, nil
}

func (_bc *ReferenceMap) Write(key string, entry ReferenceEntry) {
	_bc.Lock()
	defer _bc.Unlock()
	if _bc._bff == nil {
		_bc._bff = map[string]ReferenceEntry{}
	}
	_bc._bff[key] = entry
}

func _bfe(_bgf *_gc.RGBA, _ba string, _df string, _adg, _adca int) error {
	_dade, _faa := _e.ReadFile(_ba)
	if _faa != nil {
		return _faa
	}
	_dag, _faa := _agc.Parse(_dade)
	if _faa != nil {
		return _faa
	}
	_ead, _faa := _agc.NewFace(_dag, &_agc.FaceOptions{Size: 18, DPI: 72, Hinting: _de.HintingNone})
	if _faa != nil {
		return _faa
	}
	_bgee := &_de.Drawer{Dst: _bgf, Src: _gc.NewUniform(_a.RGBA{200, 100, 0, 255}), Face: _ead, Dot: _fe.P(_adg, _adca)}
	_bgee.DrawString(_df)
	return nil
}

func ComparePNGFiles(file1, file2 string) (bool, error) {
	_dad, _cae := HashFile(file1)
	if _cae != nil {
		return false, _cae
	}
	_ffa, _cae := HashFile(file2)
	if _cae != nil {
		return false, _cae
	}
	if _dad == _ffa {
		return true, nil
	}
	_bgdb, _cae := ReadPNG(file1)
	if _cae != nil {
		return false, _cae
	}
	_aga, _cae := ReadPNG(file2)
	if _cae != nil {
		return false, _cae
	}
	if _bgdb.Bounds() != _aga.Bounds() {
		return false, nil
	}
	return CompareImages(_bgdb, _aga)
}

func HashFile(file string) (string, error) {
	_ged, _cb := _ee.Open(file)
	if _cb != nil {
		return "", _cb
	}
	defer _ged.Close()
	_gb := _ea.New()
	if _, _cb = _gg.Copy(_gb, _ged); _cb != nil {
		return "", _cb
	}
	return _ag.EncodeToString(_gb.Sum(nil)), nil
}

func _eda(_bcc, _gfa _gc.Rectangle) bool {
	return _bcc.Min.X == _gfa.Min.X && _bcc.Min.Y == _gfa.Min.Y && _bcc.Max.X == _gfa.Max.X && _bcc.Max.Y == _gfa.Max.Y
}

func (_gcb *ReferenceMap) MarshalJSON() ([]byte, error) { return _c.Marshal(_gcb._bff) }

func RenderPDFToPNGs(pdfPath string, dpi int, outpathTpl string) error {
	if dpi <= 0 {
		dpi = 100
	}
	if _, _aca := _f.LookPath("\u0067\u0073"); _aca != nil {
		return ErrRenderNotSupported
	}
	return _f.Command("\u0067\u0073", "\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061", "\u002d\u006f", outpathTpl, _ad.Sprintf("\u002d\u0072\u0025\u0064", dpi), pdfPath).Run()
}

func _ebf(_fcbe, _eaag _a.Color) bool {
	_fcf, _gfc, _dacg, _ec := _fcbe.RGBA()
	_bbb, _gfb, _aad, _eae := _eaag.RGBA()
	return _fcf == _bbb && _gfc == _gfb && _dacg == _aad && _ec == _eae
}

var (
	ErrRenderNotSupported = _bf.New("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m")
	ErrImageSizeNotMatch  = _bf.New("\u0069\u006d\u0061ge\u0020\u0073\u0069\u007a\u0065\u0073\u0020\u0064\u006f\u006e\u0027\u0074\u0020\u006d\u0061\u0074\u0063\u0068")
)

func (_ded *ReferenceFile) Update(currentMap *ReferenceMap) error {
	if _ded._abe == "" {
		return nil
	}
	_fdc := _ded.updateMap(currentMap)
	if _fdc == 0 {
		return nil
	}
	_eb, _db := _ee.OpenFile(_ded._abe, _ee.O_CREATE|_ee.O_TRUNC|_ee.O_WRONLY, 0664)
	if _db != nil {
		return _db
	}
	defer _eb.Close()
	_ded.Timestamp = _ge.Now().UTC()
	_da := _c.NewEncoder(_eb)
	_da.SetIndent("", "\u0009")
	return _da.Encode(_ded)
}

type ReferenceMap struct {
	_age.Mutex
	_bff map[string]ReferenceEntry
}

type ReferenceEntry struct {
	Timestamp   int64   `json:"timestamp"`
	Value       string  `json:"value"`
	ResultSize  int64   `json:"resultSize,omitempty"`
	DiffPercent float64 `json:"diffPercent,omitempty"`
	DiffTotal   float64 `json:"diffValue,omitempty"`
	Invalid     bool    `json:"markedInvalid,omitempty"`
}

func ReadPNG(file string) (_gc.Image, error) {
	_eba, _edf := _ee.Open(file)
	if _edf != nil {
		return nil, _edf
	}
	defer _eba.Close()
	return _ab.Decode(_eba)
}
