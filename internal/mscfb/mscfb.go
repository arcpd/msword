package mscfb

import (
	_bb "bytes"
	_af "encoding/binary"
	_d "fmt"
	_gg "github.com/arcpd/msword/internal/mscfb/rw"
	_bd "github.com/richardlehane/msoleps/types"
	_ea "io"
	_be "os"
	_g "strconv"
	_e "time"
	_a "unicode"
	_gc "unicode/utf16"
)

func (_effe *Reader) readAt(_eag int64, _fff int) ([]byte, error) {
	if _effe._eaac {
		_dca, _gbb := _effe._gbab.(slicer).Slice(_eag, _fff)
		if _gbb != nil {
			return nil, Error{ErrRead, "\u0073\u006c\u0069\u0063er\u0020\u0072\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028" + _gbb.Error() + "\u0029", _eag}
		}
		return _dca, nil
	}
	if _fff > len(_effe._efg) {
		return nil, Error{ErrRead, "\u0072\u0065ad\u0020\u006c\u0065n\u0067\u0074\u0068\u0020gre\u0061te\u0072\u0020\u0074\u0068\u0061\u006e\u0020re\u0061\u0064\u0020\u0062\u0075\u0066\u0066e\u0072", int64(_fff)}
	}
	if _, _eage := _effe._gbab.ReadAt(_effe._efg[:_fff], _eag); _eage != nil {
		return nil, Error{ErrRead, _eage.Error(), _eag}
	}
	return _effe._efg[:_fff], nil
}

type Error struct {
	_dfce int
	_cdc  string
	_cgcb int64
}

func (_adc *Reader) traverse() error {
	_adc.File = make([]*File, 0, len(_adc._dgb))
	var (
		_bg  func(int, []string)
		_fd  error
		_bad int
	)
	_bg = func(_acb int, _da []string) {
		_bad++
		if _bad > len(_adc._dgb) {
			_fd = Error{ErrTraverse, "\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0063o\u0075\u006e\u0074\u0065\u0072\u0020\u006f\u0076\u0065\u0072f\u006c\u006f\u0077", int64(_acb)}
			return
		}
		if _acb < 0 || _acb >= len(_adc._dgb) {
			_fd = Error{ErrTraverse, "\u0069\u006c\u006ceg\u0061\u006c\u0020\u0074\u0072\u0061\u0076\u0065\u0072\u0073\u0061\u006c\u0020\u0069\u006e\u0064\u0065\u0078", int64(_acb)}
			return
		}
		_aa := _adc._dgb[_acb]
		if _aa._cg != _gafa {
			_bg(int(_aa._cg), _da)
		}
		_adc.File = append(_adc.File, _aa)
		_aa.Path = _da
		if _aa._ge != _gafa {
			if _acb > 0 {
				_bg(int(_aa._ge), append(_da, _aa.Name))
			} else {
				_bg(int(_aa._ge), _da)
			}
		}
		if _aa._edd != _gafa {
			_bg(int(_aa._edd), _da)
		}
		return
	}
	_bg(0, []string{})
	return _fd
}

type slicer interface {
	Slice(_adg int64, _fbg int) ([]byte, error)
}

func (_cgf *File) reset() { _cgf._adcc = 0; _cgf._gb = 0; _cgf._agg = _cgf._ga }

func (_gf *File) ReadAt(p []byte, off int64) (_caa int, _aggf error) {
	_dceg, _dbe, _fdc := _gf._adcc, _gf._gb, _gf._agg
	_, _aggf = _gf.Seek(off, 0)
	if _aggf == nil {
		_caa, _aggf = _gf.Read(p)
	}
	_gf._adcc, _gf._gb, _gf._agg = _dceg, _dbe, _fdc
	return _caa, _aggf
}

func _efa(_dbag, _dfbd uint32) int64 { return int64((_dfbd + 1) * _dbag) }

type fileInfo struct{ *File }

func (_dg *File) mode() _be.FileMode {
	if _dg._afa != _f {
		return _be.ModeDir | 0777
	}
	return 0666
}

func (_dc fileInfo) Name() string { return _dc.File.Name }

func (_fcb *Reader) findFatLocsOffset(_efaa bool) int64 {
	var _abb uint32
	if _efaa {
		_abb = _fcb._bbd._eceg[0]
	} else {
		_abb = _fcb._bbd._bac[0]
	}
	return _efa(_fcb._dege, _abb)
}

func (_gea fileInfo) ModTime() _e.Time { return _gea.Modified() }

func (_dfc *Reader) setDirEntries() error {
	_db := 20
	if _dfc._bbd._ccb > 0 {
		_db = int(_dfc._bbd._ccb)
	}
	_ef := make([]*File, 0, _db)
	_dfb := make(map[uint32]bool)
	_cff := int(_dfc._dege / _cba)
	_cfd := _dfc._bbd._agca
	for _cfd != _bagg {
		_ag, _fa := _dfc.readAt(_efa(_dfc._dege, _cfd), int(_dfc._dege))
		if _fa != nil {
			return Error{ErrRead, "\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020r\u0065\u0061\u0064\u0020\u0065\u0072\u0072\u006f\u0072\u0020\u0028" + _fa.Error() + "\u0029", _efa(_dfc._dege, _cfd)}
		}
		for _gd := 0; _gd < _cff; _gd++ {
			_gdc := &File{_bc: _dfc}
			_gdc.directoryEntryFields = _fe(_ag[_gd*int(_cba):])
			_bag(_dfc._bbd._daf, _gdc)
			_gdc._agg = _gdc._ga
			_ef = append(_ef, _gdc)
		}
		_dba, _fa := _dfc.findNext(_cfd, false)
		if _fa != nil {
			return Error{ErrRead, "d\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u0065r\u0072\u006f\u0072\u0020\u0066\u0069\u006e\u0064\u0069\u006eg \u0073\u0065\u0063t\u006fr\u0020\u0028" + _fa.Error() + "\u0029", int64(_dba)}
		}
		if _dba <= _cfd {
			if _dba == _cfd || _dfb[_dba] {
				return Error{ErrRead, "\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020e\u006e\u0074\u0072\u0069\u0065\u0073\u0020s\u0065\u0063\u0074\u006f\u0072\u0020\u0063\u0079\u0063\u006c\u0065", int64(_dba)}
			}
			_dfb[_dba] = true
		}
		_cfd = _dba
	}
	_dfc._dgb = _ef
	return nil
}

func (_dbfe *Reader) setDifats() error {
	_dbfe._bbd._bac = _dbfe._bbd._ddd[:]
	if _dbfe._bbd._aceb == 0 {
		return nil
	}
	_egf := (_dbfe._dege / 4) - 1
	_afe := make([]uint32, 109, _dbfe._bbd._aceb*_egf+109)
	copy(_afe, _dbfe._bbd._bac)
	_dbfe._bbd._bac = _afe
	_fdd := _dbfe._bbd._cgfe
	for _fbef := 0; _fbef < int(_dbfe._bbd._aceb); _fbef++ {
		_baa, _gbfa := _dbfe.readAt(_efa(_dbfe._dege, _fdd), int(_dbfe._dege))
		if _gbfa != nil {
			return Error{ErrFormat, "e\u0072r\u006f\u0072\u0020\u0073\u0065\u0074\u0074\u0069n\u0067\u0020\u0044\u0049FA\u0054\u0028" + _gbfa.Error() + "\u0029", int64(_fdd)}
		}
		for _gga := 0; _gga < int(_egf); _gga++ {
			_dbfe._bbd._bac = append(_dbfe._bbd._bac, _af.LittleEndian.Uint32(_baa[_gga*4:_gga*4+4]))
		}
		_fdd = _af.LittleEndian.Uint32(_baa[len(_baa)-4:])
	}
	return nil
}

type File struct {
	Name    string
	Initial uint16
	Path    []string
	Size    int64
	_adcc   int64
	_agg    uint32
	_gb     int64
	*directoryEntryFields
	_bc *Reader
}

func (_ggf fileInfo) IsDir() bool { return _ggf.mode().IsDir() }

func (_ggb *Reader) ID() string { return _ggb.File[0].ID() }

func (_cge *File) FileInfo() _be.FileInfo { return fileInfo{_cge} }

func (_ceb *File) Created() _e.Time { return _ceb._ad.Time() }

func (_egca *Reader) Next() (*File, error) {
	_egca._fgbd++
	if _egca._fgbd >= len(_egca.File) {
		return nil, _ea.EOF
	}
	return _egca.File[_egca._fgbd], nil
}

func (_abg *File) changeSize(_aea int64) error {
	if _aea == _abg.Size {
		return nil
	}
	var _dfa *File
	for _, _aef := range _abg._bc._dgb {
		if _aef.Name == _abg.Name {
			_dfa = _aef
			break
		}
	}
	if _dfa == nil {
		return _d.Errorf("\u004e\u006f\u0020\u0064\u0069\u0072e\u0063\u0074\u006f\u0072\u0079\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0072\u0020\u0061\u0020\u0066\u0069\u006ce\u003a\u0020\u0025\u0073", _abg.Name)
	}
	_gae := _bb.NewBuffer([]byte{})
	if _eaf := _af.Write(_gae, _af.LittleEndian, _aea); _eaf != nil {
		return _eaf
	}
	for _gbc, _fb := range _gae.Bytes() {
		_dfa._cgb[_gbc] = _fb
	}
	var _cgg int64
	var _gad bool
	if _abg.Size < _gaeg {
		_gad = true
		_cgg = int64(_gfd)
	} else {
		_cgg = int64(_abg._bc._dege)
	}
	_degc := int((_abg.Size-1)/_cgg) + 1
	_gead := int((_aea-1)/_cgg) + 1
	if _degc < _gead {
		_dd, _dbc := _abg.findLast(_gad)
		if _dbc != nil {
			return _dbc
		}
		_cfa, _dbc := _abg._bc.findNextFreeSector(_gad)
		if _dbc != nil {
			return _dbc
		}
		for _dfe := _gead - _degc; _dfe > 0; _dfe-- {
			if _dec := _abg._bc.saveToFatLocs(_dd, _cfa, _gad); _dec != nil {
				return _dec
			}
			if _dfe > 1 {
				_dd = _cfa
				_cfa++
			} else if _beef := _abg._bc.saveToFatLocs(_cfa, _bagg, _gad); _beef != nil {
				return _beef
			}
		}
	} else if _degc > _gead {
		_gef := _abg._ga
		var _ffc error
		for _ebg := 0; _ebg < _gead-1; _ebg++ {
			_gef, _ffc = _abg._bc.findNext(_gef, _gad)
			if _ffc != nil {
				return _ffc
			}
		}
		if _bgb := _abg._bc.saveToFatLocs(_gef, _bagg, _gad); _bgb != nil {
			return _bgb
		}
	}
	_abg.Size = _aea
	return nil
}

func _fe(_eb []byte) *directoryEntryFields {
	_cea := &directoryEntryFields{}
	for _de := range _cea._ca {
		_cea._ca[_de] = _af.LittleEndian.Uint16(_eb[_de*2 : _de*2+2])
	}
	_cea._ed = _af.LittleEndian.Uint16(_eb[64:66])
	_cea._afa = uint8(_eb[66])
	_cea._ff = uint8(_eb[67])
	_cea._cg = _af.LittleEndian.Uint32(_eb[68:72])
	_cea._edd = _af.LittleEndian.Uint32(_eb[72:76])
	_cea._ge = _af.LittleEndian.Uint32(_eb[76:80])
	_cea._df = _bd.MustGuid(_eb[80:96])
	copy(_cea._cd[:], _eb[96:100])
	_cea._ad = _bd.MustFileTime(_eb[100:108])
	_cea._ce = _bd.MustFileTime(_eb[108:116])
	_cea._ga = _af.LittleEndian.Uint32(_eb[116:120])
	copy(_cea._cgb[:], _eb[120:128])
	return _cea
}

func (_fgbf *Reader) getOffset(_dcc uint32, _gagb bool) (int64, error) {
	if _gagb {
		_ceg := _fgbf._dege / 64
		_bff := int(_dcc / _ceg)
		if _bff >= len(_fgbf._bbd._cde) {
			return 0, Error{ErrRead, "\u006di\u006e\u0069s\u0065\u0063\u0074o\u0072\u0020\u006e\u0075\u006d\u0062\u0065r\u0020\u0069\u0073\u0020\u006f\u0075t\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0073\u0065c\u0074\u006f\u0072\u0020\u0072\u0061\u006e\u0067\u0065", int64(_bff)}
		}
		_gca := _dcc % _ceg
		return int64((_fgbf._bbd._cde[_bff]+1)*_fgbf._dege + _gca*64), nil
	}
	return _efa(_fgbf._dege, _dcc), nil
}

type Reader struct {
	_eaac bool
	_dege uint32
	_efg  []byte
	_bbd  *header
	File  []*File
	_dgb  []*File
	_fgbd int
	_gbab _ea.ReaderAt
	_afg  _ea.WriterAt
}

func _ega(_ddb *directoryEntryFields) (*_bb.Buffer, error) {
	_cfab := _bb.NewBuffer([]byte{})
	for _, _aad := range _ddb._ca {
		if _gcb := _af.Write(_cfab, _af.LittleEndian, &_aad); _gcb != nil {
			return nil, _gcb
		}
	}
	if _bbg := _af.Write(_cfab, _af.LittleEndian, &_ddb._ed); _bbg != nil {
		return nil, _bbg
	}
	if _abd := _af.Write(_cfab, _af.LittleEndian, &_ddb._afa); _abd != nil {
		return nil, _abd
	}
	if _dge := _af.Write(_cfab, _af.LittleEndian, &_ddb._ff); _dge != nil {
		return nil, _dge
	}
	if _gbfc := _af.Write(_cfab, _af.LittleEndian, &_ddb._cg); _gbfc != nil {
		return nil, _gbfc
	}
	if _fdcg := _af.Write(_cfab, _af.LittleEndian, &_ddb._edd); _fdcg != nil {
		return nil, _fdcg
	}
	if _cfgb := _af.Write(_cfab, _af.LittleEndian, &_ddb._ge); _cfgb != nil {
		return nil, _cfgb
	}
	if _ccad := _af.Write(_cfab, _af.LittleEndian, &_ddb._df.DataA); _ccad != nil {
		return nil, _ccad
	}
	if _dcg := _af.Write(_cfab, _af.LittleEndian, &_ddb._df.DataB); _dcg != nil {
		return nil, _dcg
	}
	if _dbge := _af.Write(_cfab, _af.LittleEndian, &_ddb._df.DataC); _dbge != nil {
		return nil, _dbge
	}
	if _, _adf := _cfab.Write(_ddb._df.DataD[:]); _adf != nil {
		return nil, _adf
	}
	if _, _dfaa := _cfab.Write(_ddb._cd[:]); _dfaa != nil {
		return nil, _dfaa
	}
	if _deca := _af.Write(_cfab, _af.LittleEndian, &_ddb._ad.Low); _deca != nil {
		return nil, _deca
	}
	if _gcf := _af.Write(_cfab, _af.LittleEndian, &_ddb._ad.High); _gcf != nil {
		return nil, _gcf
	}
	if _cce := _af.Write(_cfab, _af.LittleEndian, &_ddb._ce.Low); _cce != nil {
		return nil, _cce
	}
	if _beed := _af.Write(_cfab, _af.LittleEndian, &_ddb._ce.High); _beed != nil {
		return nil, _beed
	}
	if _eaa := _af.Write(_cfab, _af.LittleEndian, &_ddb._ga); _eaa != nil {
		return nil, _eaa
	}
	if _, _agb := _cfab.Write(_ddb._cgb[:]); _agb != nil {
		return nil, _agb
	}
	return _cfab, nil
}

func (_bee *File) Write(b []byte) (int, error) {
	if _bee.Size < 1 || _bee._adcc >= _bee.Size {
		return 0, _ea.EOF
	}
	if _dbf := _bee.ensureWriterAt(); _dbf != nil {
		return 0, _dbf
	}
	_dce := len(b)
	if int64(_dce) > _bee.Size-_bee._adcc {
		_dce = int(_bee.Size - _bee._adcc)
	}
	_fde, _dea := _bee.stream(_dce)
	if _dea != nil {
		return 0, _dea
	}
	var _ade, _geb int
	for _, _cac := range _fde {
		_cga := _ade + int(_cac[1])
		if _cga < _ade || _cga > _dce {
			return 0, Error{ErrWrite, "\u0062\u0061d\u0020\u0077\u0072i\u0074\u0065\u0020\u006c\u0065\u006e\u0067\u0074\u0068", int64(_cga)}
		}
		_fed, _cb := _bee._bc._afg.WriteAt(b[_ade:_cga], _cac[0])
		_geb = _geb + _fed
		if _cb != nil {
			_bee._adcc += int64(_geb)
			return _geb, Error{ErrWrite, "\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0077\u0072i\u0074\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028" + _cb.Error() + "\u0029", int64(_ade)}
		}
		_ade = _cga
	}
	_bee._adcc += int64(_geb)
	if _geb != _dce {
		_dea = Error{ErrWrite, "\u0062\u0079t\u0065\u0073\u0020\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0077\u0072\u0069\u0074\u0065\u0020\u0073\u0069\u007a\u0065", int64(_geb)}
	} else if _geb < len(b) {
		_dea = _ea.EOF
	}
	return _geb, _dea
}

const (
	_ba uint8 = 0x0
	_c  uint8 = 0x1
	_f  uint8 = 0x2
	_ac uint8 = 0x5
)

const (
	_bbf  uint64 = 0xE11AB1A1E011CFD0
	_gfd  uint32 = 64
	_gaeg int64  = 4096
	_cba  uint32 = 128
)

const _aceg int = 8 + 16 + 10 + 6 + 12 + 8 + 16 + 109*4

func (_dfaaf *Reader) Modified() _e.Time { return _dfaaf.File[0].Modified() }

func (_ggca *Reader) setMiniStream() error {
	if _ggca._dgb[0]._ga == _bagg || _ggca._bbd._efde == _bagg || _ggca._bbd._bdgg == 0 {
		return nil
	}
	_cggg := int(_ggca._bbd._bdgg)
	_ggca._bbd._eceg = make([]uint32, _cggg)
	_ggca._bbd._eceg[0] = _ggca._bbd._efde
	for _dda := 1; _dda < _cggg; _dda++ {
		_efaf, _acg := _ggca.findNext(_ggca._bbd._eceg[_dda-1], false)
		if _acg != nil {
			return Error{ErrFormat, "s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028" + _acg.Error() + "\u0029", int64(_ggca._bbd._eceg[_dda-1])}
		}
		_ggca._bbd._eceg[_dda] = _efaf
	}
	_cggg = int(_ggca._dege / 4 * _ggca._bbd._bdgg)
	_ggca._bbd._cde = make([]uint32, 0, _cggg)
	_acfa := _ggca._dgb[0]._ga
	var _aac error
	for _acfa != _bagg {
		_ggca._bbd._cde = append(_ggca._bbd._cde, _acfa)
		_acfa, _aac = _ggca.findNext(_acfa, false)
		if _aac != nil {
			return Error{ErrFormat, "s\u0065\u0074\u0074\u0069ng\u0020m\u0069\u006e\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u0028" + _aac.Error() + "\u0029", int64(_acfa)}
		}
	}
	return nil
}

func (_dbffg *Reader) saveToFatLocs(_bgd uint32, _efb interface{}, _aagd bool) error {
	_adea := _bb.NewBuffer([]byte{})
	if _bce := _af.Write(_adea, _af.LittleEndian, _efb); _bce != nil {
		return _bce
	}
	_ebga := _dbffg.findFatLocsOffset(_aagd) + int64(_bgd*4)
	_, _gcd := _dbffg._afg.WriteAt(_adea.Bytes(), _ebga)
	return _gcd
}

func (_ggc *File) WriteAt(p []byte, off int64) (_cca int, _aaf error) {
	_gda, _bdc, _add := _ggc._adcc, _ggc._gb, _ggc._agg
	_, _aaf = _ggc.Seek(off, 0)
	if _aaf == nil {
		_cca, _aaf = _ggc.Write(p)
	}
	_ggc._adcc, _ggc._gb, _ggc._agg = _gda, _bdc, _add
	return _cca, _aaf
}

type headerFields struct {
	_fbe  uint64
	_     [16]byte
	_bbee uint16
	_daf  uint16
	_     [2]byte
	_bba  uint16
	_     [2]byte
	_     [6]byte
	_ccb  uint32
	_gdg  uint32
	_agca uint32
	_     [4]byte
	_     [4]byte
	_efde uint32
	_bdgg uint32
	_cgfe uint32
	_aceb uint32
	_ddd  [109]uint32
}

func (_bde *Reader) Created() _e.Time { return _bde.File[0].Created() }

func (_feg *File) seek(_addf int64) error {
	var _edbc bool
	var _bf int64
	if _feg.Size < _gaeg {
		_edbc = true
		_bf = 64
	} else {
		_bf = int64(_feg._bc._dege)
	}
	var _fdcb int64
	var _agc error
	if _feg._gb > 0 {
		if _bf-_feg._gb <= _addf {
			_feg._agg, _agc = _feg._bc.findNext(_feg._agg, _edbc)
			if _agc != nil {
				return _agc
			}
			_fdcb += _bf - _feg._gb
			_feg._gb = 0
			if _fdcb == _addf {
				return nil
			}
		} else {
			_feg._gb += _addf
			return nil
		}
		if _feg._agg == _bagg {
			return Error{ErrRead, "\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e", int64(_feg._agg)}
		}
	}
	for {
		if _addf-_fdcb < _bf {
			_feg._gb = _addf - _fdcb
			return nil
		} else {
			_fdcb += _bf
			_feg._agg, _agc = _feg._bc.findNext(_feg._agg, _edbc)
			if _agc != nil {
				return _agc
			}
			if _fdcb == _addf {
				return nil
			}
		}
	}
}

const (
	_dde  uint32 = 0xFFFFFFFA
	_gcg  uint32 = 0xFFFFFFFC
	_eafb uint32 = 0xFFFFFFFD
	_bagg uint32 = 0xFFFFFFFE
	_ded  uint32 = 0xFFFFFFFF
	_egd  byte   = 0xFF
	_dcf  uint32 = 0xFFFFFFFA
	_gafa uint32 = 0xFFFFFFFF
)

func (_aee *Reader) exportFAT(_cfc *_gg.Writer, _gcgf []uint32) error {
	if _aee._bbd._gdg == 0 {
		return nil
	}
	_dfcc := _bb.NewBuffer([]byte{})
	if _ddc := _af.Write(_dfcc, _af.LittleEndian, _eafb); _ddc != nil {
		return _ddc
	}
	for _ddba := 0; _ddba < len(_gcgf)-1; _ddba++ {
		for _cfb := _gcgf[_ddba]; _cfb < _gcgf[_ddba+1]-1; _cfb++ {
			if _bcf := _af.Write(_dfcc, _af.LittleEndian, _cfb); _bcf != nil {
				return _bcf
			}
		}
		if _edgg := _af.Write(_dfcc, _af.LittleEndian, _bagg); _edgg != nil {
			return _edgg
		}
	}
	_abbc := 512
	for _, _bef := range _dfcc.Bytes() {
		if _gaef := _cfc.WriteByteAt(_bef, _abbc); _gaef != nil {
			return _gaef
		}
		_abbc++
	}
	return nil
}

func (_agda *Reader) Debug() map[string][]uint32 {
	_gfg := map[string][]uint32{"s\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065": []uint32{_agda._dege}, "\u006d\u0069\u006e\u0069\u0020\u0066\u0061\u0074\u0020\u006c\u006f\u0063\u0073": _agda._bbd._eceg, "\u006d\u0069n\u0069\u0020\u0073t\u0072\u0065\u0061\u006d\u0020\u006c\u006f\u0063\u0073": _agda._bbd._cde, "\u0064\u0069r\u0065\u0063\u0074o\u0072\u0079\u0020\u0073\u0065\u0063\u0074\u006f\u0072": []uint32{_agda._bbd._agca}, "\u006d\u0069\u006e\u0069 s\u0074\u0072\u0065\u0061\u006d\u0020\u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069z\u0065": []uint32{_agda.File[0]._ga, _af.LittleEndian.Uint32(_agda.File[0]._cgb[:])}}
	for _cbc, _fegb := _agda.Next(); _fegb == nil; _cbc, _fegb = _agda.Next() {
		_gfg[_cbc.Name+" \u0073\u0074\u0061\u0072\u0074\u002f\u0073\u0069\u007a\u0065"] = []uint32{_cbc._ga, _af.LittleEndian.Uint32(_cbc._cgb[:])}
	}
	return _gfg
}

func (_egdf *Reader) Read(b []byte) (_fffa int, _becb error) {
	if _egdf._fgbd >= len(_egdf.File) {
		return 0, _ea.EOF
	}
	return _egdf.File[_egdf._fgbd].Read(b)
}

func (_fcc *File) Read(b []byte) (int, error) {
	if _fcc.Size < 1 || _fcc._adcc >= _fcc.Size {
		return 0, _ea.EOF
	}
	_fce := len(b)
	if int64(_fce) > _fcc.Size-_fcc._adcc {
		_fce = int(_fcc.Size - _fcc._adcc)
	}
	_aga, _ae := _fcc.stream(_fce)
	if _ae != nil {
		return 0, _ae
	}
	var _adb, _efd int
	for _, _bbe := range _aga {
		_bcg := _adb + int(_bbe[1])
		if _bcg < _adb || _bcg > _fce {
			return 0, Error{ErrRead, "\u0062a\u0064 \u0072\u0065\u0061\u0064\u0020\u006c\u0065\u006e\u0067\u0074\u0068", int64(_bcg)}
		}
		_bgc, _eddc := _fcc._bc._gbab.ReadAt(b[_adb:_bcg], _bbe[0])
		_efd = _efd + _bgc
		if _eddc != nil {
			_fcc._adcc += int64(_efd)
			return _efd, Error{ErrRead, "\u0075n\u0064\u0065\u0072\u006c\u0079\u0069\u006e\u0067\u0020\u0072\u0065a\u0064\u0065\u0072\u0020\u0066\u0061\u0069\u006c\u0020\u0028" + _eddc.Error() + "\u0029", int64(_adb)}
		}
		_adb = _bcg
	}
	_fcc._adcc += int64(_efd)
	if _efd != _fce {
		_ae = Error{ErrRead, "\u0062\u0079\u0074e\u0073\u0020\u0072\u0065\u0061\u0064\u0020\u0064\u006f\u0020\u006e\u006f\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020r\u0065\u0061\u0064\u0020\u0073\u0069\u007a\u0065", int64(_efd)}
	} else if _efd < len(b) {
		_ae = _ea.EOF
	}
	return _efd, _ae
}

func (_aba *Reader) exportDirEntries(_cbd *_gg.Writer) error {
	if int64(_cbd.Len()) != _efa(_aba._dege, _aba._bbd._agca) {
		return Error{ErrWrite, _d.Sprintf("I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u0077\u0072\u0069\u0074\u0065\u0072\u0020l\u0065\u006e\u0067t\u0068:\u0020\u0025\u0076", _cbd.Len()), 0}
	}
	for _, _fbd := range _aba._dgb {
		_fgb, _gdbe := _ega(_fbd.directoryEntryFields)
		if _gdbe != nil {
			return _gdbe
		}
		if _, _ace := _ea.Copy(_cbd, _fgb); _ace != nil {
			return _ace
		}
	}
	return nil
}

func (_ec fileInfo) Size() int64 {
	if _ec._afa != _f {
		return 0
	}
	return _ec.File.Size
}

func _edg(_dcec []byte) *headerFields {
	_dad := &headerFields{}
	_dad._fbe = _af.LittleEndian.Uint64(_dcec[:8])
	_dad._bbee = _af.LittleEndian.Uint16(_dcec[24:26])
	_dad._daf = _af.LittleEndian.Uint16(_dcec[26:28])
	_dad._bba = _af.LittleEndian.Uint16(_dcec[30:32])
	_dad._ccb = _af.LittleEndian.Uint32(_dcec[40:44])
	_dad._gdg = _af.LittleEndian.Uint32(_dcec[44:48])
	_dad._agca = _af.LittleEndian.Uint32(_dcec[48:52])
	_dad._efde = _af.LittleEndian.Uint32(_dcec[60:64])
	_dad._bdgg = _af.LittleEndian.Uint32(_dcec[64:68])
	_dad._cgfe = _af.LittleEndian.Uint32(_dcec[68:72])
	_dad._aceb = _af.LittleEndian.Uint32(_dcec[72:76])
	var _cgbdd int
	for _agbd := 76; _agbd < 512; _agbd = _agbd + 4 {
		_dad._ddd[_cgbdd] = _af.LittleEndian.Uint32(_dcec[_agbd : _agbd+4])
		_cgbdd++
	}
	return _dad
}

const (
	ErrFormat = iota
	ErrRead
	ErrSeek
	ErrWrite
	ErrTraverse
)

const _cc int = 64 + 4*4 + 16 + 4 + 8*2 + 4 + 8

func New(ra _ea.ReaderAt) (*Reader, error) {
	_ggd := &Reader{_gbab: ra}
	if _, _daag := ra.(slicer); _daag {
		_ggd._eaac = true
	} else {
		_ggd._efg = make([]byte, _aceg)
	}
	if _gfe := _ggd.setHeader(); _gfe != nil {
		return nil, _gfe
	}
	if !_ggd._eaac && int(_ggd._dege) > len(_ggd._efg) {
		_ggd._efg = make([]byte, _ggd._dege)
	}
	if _fdf := _ggd.setDifats(); _fdf != nil {
		return nil, _fdf
	}
	if _bagb := _ggd.setDirEntries(); _bagb != nil {
		return nil, _bagb
	}
	if _dbff := _ggd.setMiniStream(); _dbff != nil {
		return nil, _dbff
	}
	if _dgcb := _ggd.traverse(); _dgcb != nil {
		return nil, _dgcb
	}
	return _ggd, nil
}

func (_gcc *Reader) exportDifats(_ccef *_gg.Writer) error {
	if _gcc._bbd._aceb == 0 {
		return nil
	}
	return nil
}

func (_fc *File) Modified() _e.Time { return _fc._ce.Time() }

func (_gdac Error) Typ() int { return _gdac._dfce }

func (_bgbc *Reader) findNext(_fcf uint32, _ebe bool) (uint32, error) {
	_ead := _bgbc._dege / 4
	_dbce := int(_fcf / _ead)
	var _abc uint32
	if _ebe {
		if _dbce < 0 || _dbce >= len(_bgbc._bbd._eceg) {
			return 0, Error{ErrRead, "\u006d\u0069\u006e\u0069\u0073e\u0063\u0074\u006f\u0072\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073 \u006f\u0075\u0074\u0073\u0069\u0064\u0065\u0020\u006d\u0069\u006e\u0069\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065", int64(_dbce)}
		}
		_abc = _bgbc._bbd._eceg[_dbce]
	} else {
		if _dbce < 0 || _dbce >= len(_bgbc._bbd._bac) {
			return 0, Error{ErrRead, "\u0046\u0041\u0054\u0020\u0069\u006e\u0064\u0065\u0078\u0020\u0069\u0073\u0020\u006f\u0075t\u0073i\u0064\u0065\u0020\u0044\u0049\u0046\u0041\u0054\u0020\u0072\u0061\u006e\u0067\u0065", int64(_dbce)}
		}
		_abc = _bgbc._bbd._bac[_dbce]
	}
	_dgg := _fcf % _ead
	_cgdb := _efa(_bgbc._dege, _abc) + int64(_dgg*4)
	_egc, _dcad := _bgbc.readAt(_cgdb, 4)
	if _dcad != nil {
		return 0, Error{ErrRead, "\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028" + _dcad.Error() + "\u0029", _cgdb}
	}
	_cfga := _af.LittleEndian.Uint32(_egc)
	return _cfga, nil
}

func _deb(_deg *File) {
	if _deg._ed < 4 || _deg._ed > 64 {
		return
	}
	_acf := int(_deg._ed/2 - 1)
	_deg.Initial = _deg._ca[0]
	var _ada int
	if !_a.IsPrint(rune(_deg.Initial)) {
		_ada = 1
	}
	_deg.Name = string(_gc.Decode(_deg._ca[_ada:_acf]))
}

func (_def *Reader) setHeader() error {
	_afd, _aaa := _def.readAt(0, _aceg)
	if _aaa != nil {
		return _aaa
	}
	_def._bbd = &header{headerFields: _edg(_afd)}
	if _def._bbd._fbe != _bbf {
		return Error{ErrFormat, "\u0062\u0061\u0064\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065", int64(_def._bbd._fbe)}
	}
	if _def._bbd._bba == 0x0009 || _def._bbd._bba == 0x000c {
		_def._dege = uint32(1 << _def._bbd._bba)
	} else {
		return Error{ErrFormat, "\u0069\u006c\u006c\u0065ga\u006c\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0073\u0069\u007a\u0065", int64(_def._bbd._bba)}
	}
	if _def._bbd._aceb > 0 {
		_ebgf := (_def._dege / 4) - 1
		if int(_def._bbd._aceb*_ebgf+109) < 0 {
			return Error{ErrFormat, "\u0044I\u0046A\u0054\u0020\u0069\u006e\u0074 \u006f\u0076e\u0072\u0066\u006c\u006f\u0077", int64(_def._bbd._aceb)}
		}
		if _def._bbd._aceb*_ebgf+109 > _def._bbd._gdg+_ebgf {
			return Error{ErrFormat, "\u006e\u0075\u006d\u0020\u0044\u0049\u0046\u0041\u0054\u0073 \u0065\u0078\u0063\u0065\u0065\u0064\u0073 \u0046\u0041\u0054\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0073", int64(_def._bbd._aceb)}
		}
	}
	if _def._bbd._bdgg > 0 {
		if int(_def._dege/4*_def._bbd._bdgg) < 0 {
			return Error{ErrFormat, "m\u0069\u006e\u0069\u0020FA\u0054 \u0069\u006e\u0074\u0020\u006fv\u0065\u0072\u0066\u006c\u006f\u0077", int64(_def._bbd._bdgg)}
		}
		if _def._bbd._bdgg > _def._bbd._gdg*(_def._dege/_gfd) {
			return Error{ErrFormat, "\u006e\u0075\u006d\u0020\u006d\u0069n\u0069\u0020\u0046\u0041\u0054\u0073\u0020\u0065\u0078\u0063\u0065\u0065\u0064s\u0020\u0046\u0041\u0054\u0020\u0073\u0065c\u0074\u006f\u0072\u0073", int64(_def._bbd._gdg)}
		}
	}
	return nil
}

func _dbga(_dbd [][2]int64) [][2]int64 {
	_gba := len(_dbd)
	for _ede, _ffcb := 0, 0; _ede < _gba && _ffcb+1 < len(_dbd); _ede++ {
		if _dbd[_ffcb][0]+_dbd[_ffcb][1] == _dbd[_ffcb+1][0] {
			_dbd[_ffcb][1] = _dbd[_ffcb][1] + _dbd[_ffcb+1][1]
			for _ceae := range _dbd[_ffcb+1 : len(_dbd)-1] {
				_dbd[_ffcb+1+_ceae] = _dbd[_ceae+_ffcb+2]
			}
			_dbd = _dbd[:len(_dbd)-1]
		} else {
			_ffcb += 1
		}
	}
	return _dbd
}

func (_gdb fileInfo) Mode() _be.FileMode { return _gdb.File.mode() }

func (_ebb *File) ID() string { return _ebb._df.String() }

func (_ab *File) ensureWriterAt() error {
	if _ab._bc._afg == nil {
		_bcc, _ece := _ab._bc._gbab.(_ea.WriterAt)
		if !_ece {
			return Error{ErrWrite, "\u006d\u0073\u0063\u0066\u0062\u002e\u004ee\u0077\u0020\u006d\u0075\u0073\u0074\u0020\u0062\u0065\u0020\u0067\u0069\u0076\u0065n\u0020R\u0065\u0061\u0064\u0065\u0072\u0041t\u0020\u0063\u006f\u006e\u0076\u0065\u0072t\u0069\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0061\u0020\u0069\u006f\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u0041\u0074\u0020\u0069n\u0020\u006f\u0072\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0077\u0072\u0069t\u0065", 0}
		}
		_ab._bc._afg = _bcc
	}
	return nil
}

func (_cddg *Reader) exportMiniStream() (*_gg.Writer, *_gg.Writer, error) {
	_cace, _fbdb := _gg.NewWriter(), _gg.NewWriter()
	_cebd := uint32(0)
	for _, _ebf := range _cddg.File {
		if _ebf.Size == 0 {
			continue
		}
		_ebf.reset()
		_ebf._ga = _cebd
		_eddg := int(_ebf.Size) / int(_gfd)
		if int(_ebf.Size)%int(_gfd) != 0 {
			_eddg++
		}
		for _dddf := 1; _dddf < _eddg; _dddf++ {
			_cebd++
			if _dedg := _af.Write(_cace, _af.LittleEndian, _cebd); _dedg != nil {
				return nil, nil, _dedg
			}
		}
		if _bdfa := _af.Write(_cace, _af.LittleEndian, _bagg); _bdfa != nil {
			return nil, nil, _bdfa
		}
		_cebd++
		if _, _dee := _ea.Copy(_fbdb, _ebf); _dee != nil {
			return nil, nil, _dee
		}
		if _geg := _fbdb.AlignLength(64); _geg != nil {
			return nil, nil, _geg
		}
	}
	if _fgbfa := _cace.FillWithByte(int(_cddg._dege)-_cace.Len(), _egd); _fgbfa != nil {
		return nil, nil, _fgbfa
	}
	if _gdae := _fbdb.AlignLength(int(_cddg._dege)); _gdae != nil {
		return nil, nil, _gdae
	}
	return _cace, _fbdb, nil
}

func (_cgfd Error) Error() string {
	return "\u006ds\u0063\u0066\u0062\u003a\u0020" + _cgfd._cdc + "\u003b\u0020" + _g.FormatInt(_cgfd._cgcb, 10)
}

func (_dgc *File) findLast(_dac bool) (uint32, error) {
	_aag := _dgc._ga
	for {
		_dbb, _gbf := _dgc._bc.findNext(_aag, _dac)
		if _gbf != nil {
			return 0, Error{ErrRead, "\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028" + _gbf.Error() + "\u0029", 0}
		}
		if _dbb == _bagg {
			break
		}
		_aag = _dbb
	}
	return _aag, nil
}

type header struct {
	*headerFields
	_bac  []uint32
	_eceg []uint32
	_cde  []uint32
}

const (
	_cf  uint8 = 0x0
	_acd uint8 = 0x1
)

func (_bec *File) Seek(offset int64, whence int) (int64, error) {
	var _dbg int64
	switch whence {
	default:
		return 0, Error{ErrSeek, "\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077h\u0065\u006e\u0063\u0065", int64(whence)}
	case 0:
		_dbg = offset
	case 1:
		_dbg = _bec._adcc + offset
	case 2:
		_dbg = _bec.Size - offset
	}
	switch {
	case _dbg < 0:
		return _bec._adcc, Error{ErrSeek, "\u0063\u0061\u006e'\u0074\u0020\u0073\u0065e\u006b\u0020\u0062\u0065\u0066\u006f\u0072e\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u006f\u0066\u0020\u0046\u0069\u006c\u0065", _dbg}
	case _dbg >= _bec.Size:
		return _bec._adcc, Error{ErrSeek, "c\u0061\u006e\u0027\u0074\u0020\u0073e\u0065\u006b\u0020\u0070\u0061\u0073\u0074\u0020\u0046i\u006c\u0065\u0020l\u0065n\u0067\u0074\u0068", _dbg}
	case _dbg == _bec._adcc:
		return _dbg, nil
	case _dbg > _bec._adcc:
		_cfg := _bec._adcc
		_bec._adcc = _dbg
		return _bec._adcc, _bec.seek(_dbg - _cfg)
	}
	if _bec._gb >= _bec._adcc-_dbg {
		_bec._gb = _bec._gb - (_bec._adcc - _dbg)
		_bec._adcc = _dbg
		return _bec._adcc, nil
	}
	_bec._gb = 0
	_bec._agg = _bec._ga
	_bec._adcc = _dbg
	return _bec._adcc, _bec.seek(_dbg)
}

func (_gaf fileInfo) Sys() interface{} { return nil }

func (_dgd *Reader) GetEntry(entryName string) (*File, error) {
	for _abaa, _bdac := _dgd.Next(); _bdac == nil; _abaa, _bdac = _dgd.Next() {
		if _abaa.Name == entryName {
			return _abaa, nil
		}
	}
	return nil, Error{ErrTraverse, "\u004e\u006f\u0020\u0065\u006e\u0074\u0072\u0079\u0020\u0066o\u0075\u006e\u0064\u0020\u0066\u006f\u0072 \u0067\u0069\u0076\u0065\u006e\u0020\u006e\u0061\u006d\u0065\u002e", 0}
}

func (_aed *File) SetEntryContent(b []byte) error {
	if _ffe := _aed.ensureWriterAt(); _ffe != nil {
		return _ffe
	}
	_aed.reset()
	if _ccf := _aed.changeSize(int64(len(b))); _ccf != nil {
		return _ccf
	}
	_, _bdg := _aed.Write(b)
	return _bdg
}

func (_cebf *File) stream(_gbg int) ([][2]int64, error) {
	var _cdf bool
	var _cgd int
	var _gab int64
	if _cebf.Size < _gaeg {
		_cdf = true
		_cgd = _gbg/int(_gfd) + 2
		_gab = int64(_gfd)
	} else {
		_cgd = _gbg/int(_cebf._bc._dege) + 2
		_gab = int64(_cebf._bc._dege)
	}
	_aab := make([][2]int64, 0, _cgd)
	var _dcb, _eg int
	if _cebf._gb > 0 {
		_bda, _dcd := _cebf._bc.getOffset(_cebf._agg, _cdf)
		if _dcd != nil {
			return nil, _dcd
		}
		if _gab-_cebf._gb >= int64(_gbg) {
			_aab = append(_aab, [2]int64{_bda + _cebf._gb, int64(_gbg)})
		} else {
			_aab = append(_aab, [2]int64{_bda + _cebf._gb, _gab - _cebf._gb})
		}
		if _gab-_cebf._gb <= int64(_gbg) {
			_cebf._agg, _dcd = _cebf._bc.findNext(_cebf._agg, _cdf)
			if _dcd != nil {
				return nil, _dcd
			}
			_eg += int(_gab - _cebf._gb)
			_cebf._gb = 0
		} else {
			_cebf._gb += int64(_gbg)
		}
		if _aab[0][1] == int64(_gbg) {
			return _aab, nil
		}
		if _cebf._agg == _bagg {
			return nil, Error{ErrRead, "\u0075\u006ee\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0065\u0061\u0072\u006c\u0079\u0020\u0065\u006e\u0064\u0020\u006f\u0066\u0020\u0063ha\u0069\u006e", int64(_cebf._agg)}
		}
		_dcb++
	}
	for {
		if _dcb >= cap(_aab) {
			return nil, Error{ErrRead, "\u0069\u006e\u0064\u0065x\u0020\u006f\u0076\u0065\u0072\u0072\u0075\u006e\u0073\u0020s\u0065c\u0074\u006f\u0072\u0020\u006c\u0065\u006eg\u0074\u0068", int64(_dcb)}
		}
		_fegc, _bafc := _cebf._bc.getOffset(_cebf._agg, _cdf)
		if _bafc != nil {
			return nil, _bafc
		}
		if _gbg-_eg < int(_gab) {
			_aab = append(_aab, [2]int64{_fegc, int64(_gbg - _eg)})
			_cebf._gb = int64(_gbg - _eg)
			return _dbga(_aab), nil
		} else {
			_aab = append(_aab, [2]int64{_fegc, _gab})
			_eg += int(_gab)
			_cebf._agg, _bafc = _cebf._bc.findNext(_cebf._agg, _cdf)
			if _bafc != nil {
				return nil, _bafc
			}
			if _eg == _gbg {
				return _dbga(_aab), nil
			}
		}
		_dcb++
	}
}

func _bag(_eac uint16, _baf *File) {
	_deb(_baf)
	if _baf._afa != _f {
		return
	}
	if _eac > 3 {
		_baf.Size = int64(_af.LittleEndian.Uint64(_baf._cgb[:]))
	} else {
		_baf.Size = int64(_af.LittleEndian.Uint32(_baf._cgb[:4]))
	}
}

func (_geaf *Reader) findNextFreeSector(_abgc bool) (uint32, error) {
	_aacg := _geaf.findFatLocsOffset(_abgc)
	_aada := uint32(0)
	_dfbc := _geaf._dege / 4
	for {
		_afc, _bdf := _geaf.readAt(_aacg, 4)
		if _bdf != nil {
			return 0, Error{ErrRead, "\u0062\u0061\u0064\u0020\u0072\u0065\u0061\u0064\u0020\u0066i\u006e\u0064\u0069\u006e\u0067\u0020\u006ee\u0078\u0074\u0020\u0073\u0065\u0063\u0074\u006f\u0072\u0020\u0028" + _bdf.Error() + "\u0029", _aacg}
		}
		_cgc := _af.LittleEndian.Uint32(_afc)
		if _cgc == _ded {
			break
		}
		if _aada >= _dfbc {
			return 0, Error{ErrRead, "\u0065\u006e\u0064\u0020of\u0020\u006d\u0069\u006e\u0069\u0046\u0061\u0074\u0020\u0072\u0065\u0061\u0063\u0068e\u0064", _aacg}
		}
		_aada++
		_aacg += 4
	}
	return _aada, nil
}

func (_fbdc *Reader) exportHeader(_gac *_gg.Writer) error {
	if _agbe := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._fbe); _agbe != nil {
		return _agbe
	}
	if _abcf := _gac.Skip(16); _abcf != nil {
		return _abcf
	}
	if _ecf := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._bbee); _ecf != nil {
		return _ecf
	}
	if _adbe := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._daf); _adbe != nil {
		return _adbe
	}
	if _bfc := _af.Write(_gac, _af.LittleEndian, uint16(0xfffe)); _bfc != nil {
		return _bfc
	}
	if _adbb := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._bba); _adbb != nil {
		return _adbb
	}
	if _ced := _af.Write(_gac, _af.LittleEndian, uint16(0x0006)); _ced != nil {
		return _ced
	}
	if _aff := _gac.Skip(6); _aff != nil {
		return _aff
	}
	if _ffb := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._ccb); _ffb != nil {
		return _ffb
	}
	if _faa := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._gdg); _faa != nil {
		return _faa
	}
	if _cbb := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._agca); _cbb != nil {
		return _cbb
	}
	if _adge := _gac.Skip(4); _adge != nil {
		return _adge
	}
	if _cdee := _af.Write(_gac, _af.LittleEndian, uint32(0x00001000)); _cdee != nil {
		return _cdee
	}
	if _abba := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._efde); _abba != nil {
		return _abba
	}
	if _egb := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._bdgg); _egb != nil {
		return _egb
	}
	if _dgdb := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._cgfe); _dgdb != nil {
		return _dgdb
	}
	if _cddc := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._aceb); _cddc != nil {
		return _cddc
	}
	for _bca := 0; _bca < 109; _bca++ {
		if _egfa := _af.Write(_gac, _af.LittleEndian, &_fbdc._bbd._ddd[_bca]); _egfa != nil {
			return _egfa
		}
	}
	return nil
}

type directoryEntryFields struct {
	_ca  [32]uint16
	_ed  uint16
	_afa uint8
	_ff  uint8
	_cg  uint32
	_edd uint32
	_ge  uint32
	_df  _bd.Guid
	_cd  [4]byte
	_ad  _bd.FileTime
	_ce  _bd.FileTime
	_ga  uint32
	_cgb [8]byte
}

func (_cbad *Reader) GetHeader() *header { return _cbad._bbd }

func (_gfc *Reader) Export() ([]byte, error) {
	_degcb := _gg.NewWriter()
	if _bagf := _gfc.exportHeader(_degcb); _bagf != nil {
		return nil, _bagf
	}
	if _cdd := _degcb.FillWithByte(512, _egd); _cdd != nil {
		return nil, _cdd
	}
	_cfgag := []uint32{}
	if _ggbf := _gfc.exportDifats(_degcb); _ggbf != nil {
		return nil, _ggbf
	}
	_dcegc, _fedc, _bcgf := _gfc.exportMiniStream()
	if _bcgf != nil {
		return nil, _bcgf
	}
	_cfgag = append(_cfgag, uint32(_degcb.Len())/_gfc._dege)
	if _gee := _gfc.exportDirEntries(_degcb); _gee != nil {
		return nil, _gee
	}
	_cfgag = append(_cfgag, uint32(_degcb.Len())/_gfc._dege)
	if _, _fbb := _dcegc.WriteTo(_degcb); _fbb != nil {
		return nil, _fbb
	}
	_cfgag = append(_cfgag, uint32(_degcb.Len())/_gfc._dege)
	if _, _bbdg := _fedc.WriteTo(_degcb); _bbdg != nil {
		return nil, _bbdg
	}
	_cfgag = append(_cfgag, uint32(_degcb.Len())/_gfc._dege)
	if _dae := _gfc.exportFAT(_degcb, _cfgag); _dae != nil {
		return nil, _dae
	}
	return _degcb.Bytes(), nil
}
