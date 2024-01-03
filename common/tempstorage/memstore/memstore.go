package memstore

import (
	_gff "encoding/hex"
	_db "errors"
	_gf "fmt"
	_de "github.com/arcpd/msword/common/tempstorage"
	_f "io"
	_a "io/ioutil"
	_dg "math/rand"
	_g "sync"
)

// Open returns tempstorage File object by name
func (_cfb *memStorage) Open(path string) (_de.File, error) {
	_fdc, _bg := _cfb._ffd.Load(path)
	if !_bg {
		return nil, _db.New(_gf.Sprintf("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073", path))
	}
	return &memFile{_aa: _fdc.(*memDataCell)}, nil
}

// ReadAt reads from the underlying memDataCell at an offset provided in order to implement ReaderAt interface.
// It does not affect f.readOffset.
func (_cf *memFile) ReadAt(p []byte, readOffset int64) (int, error) {
	_ffa := _cf._aa._def
	_ac := int64(len(p))
	if _ac > _ffa {
		_ac = _ffa
		p = p[:_ac]
	}
	if readOffset >= _ffa {
		return 0, _f.EOF
	}
	_df := readOffset + _ac
	if _df >= _ffa {
		_df = _ffa
	}
	_cfg := copy(p, _cf._aa._gfa[readOffset:_df])
	return _cfg, nil
}

type memStorage struct{ _ffd _g.Map }

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage() { _fd := memStorage{_ffd: _g.Map{}}; _de.SetAsStorage(&_fd) }

func _ba(_ce string) string { _ddb, _ := _afg(6); return _ce + _ddb }

type memDataCell struct {
	_fe  string
	_gfa []byte
	_def int64
}

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_gc *memFile) Read(p []byte) (int, error) {
	_cg := _gc._c
	_af := _gc._aa._def
	_ff := int64(len(p))
	if _ff > _af {
		_ff = _af
		p = p[:_ff]
	}
	if _cg >= _af {
		return 0, _f.EOF
	}
	_e := _cg + _ff
	if _e >= _af {
		_e = _af
	}
	_afc := copy(p, _gc._aa._gfa[_cg:_e])
	_gc._c = _e
	return _afc, nil
}

// Close is not applicable in this implementation
func (_cb *memFile) Close() error { return nil }

// Name returns the filename of the underlying memDataCell
func (_b *memFile) Name() string { return _b._aa._fe }

// Add reads a file from a disk and adds it to the storage
func (_fba *memStorage) Add(path string) error {
	_, _gfad := _fba._ffd.Load(path)
	if _gfad {
		return nil
	}
	_cff, _dee := _a.ReadFile(path)
	if _dee != nil {
		return _dee
	}
	_fba._ffd.Store(path, &memDataCell{_fe: path, _gfa: _cff, _def: int64(len(_cff))})
	return nil
}

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_fb *memFile) Write(p []byte) (int, error) {
	_fb._aa._gfa = append(_fb._aa._gfa, p...)
	_fb._aa._def += int64(len(p))
	return len(p), nil
}

func _afg(_gg int) (string, error) {
	_ed := make([]byte, _gg)
	if _, _bc := _dg.Read(_ed); _bc != nil {
		return "", _bc
	}
	return _gff.EncodeToString(_ed), nil
}

// TempFile creates a new empty file in the storage and returns it
func (_ee *memStorage) TempFile(dir, pattern string) (_de.File, error) {
	_ag := dir + "\u002f" + _ba(pattern)
	_ga := &memDataCell{_fe: _ag, _gfa: []byte{}}
	_cbd := &memFile{_aa: _ga}
	_ee._ffd.Store(_ag, _ga)
	return _cbd, nil
}

type memFile struct {
	_aa *memDataCell
	_c  int64
}

// RemoveAll removes all files according to the dir argument prefix
func (_dd *memStorage) RemoveAll(dir string) error {
	_dd._ffd.Range(func(_cbdd, _ca interface{}) bool { _dd._ffd.Delete(_cbdd); return true })
	return nil
}

// TempDir creates a name for a new temp directory using a pattern argument
func (_dfc *memStorage) TempDir(pattern string) (string, error) { return _ba(pattern), nil }
