package rw

import (
	_g "bytes"
	_ed "encoding/binary"
	_c "errors"
	_ag "fmt"
	_ac "io"
	_cf "io/ioutil"
	_e "reflect"
)

var _gbg = _c.New("r\u0077.\u0057\u0072\u0069\u0074\u0065\u0072\u003a\u0020t\u006f\u006f\u0020\u006car\u0067\u0065")

func NewWriter() *Writer { return &Writer{_gb: []byte{}} }

func (_ade *Writer) align(_eb int) error { return _ade.Skip((_eb - (_ade.Len())%_eb) % _eb) }

const _fc = 64

func (_b *Reader) skip(_cb int) error {
	_, _cbd := _ac.CopyN(_cf.Discard, _b, int64(_cb))
	if _cbd != nil {
		return _cbd
	}
	return nil
}

func _edfg(_aed int) []byte {
	defer func() {
		if recover() != nil {
			panic(_gbg)
		}
	}()
	return make([]byte, _aed)
}

func (_ad *Reader) ReadStringProperty(n uint32) (string, error) {
	if _dc := _ad.align(4); _dc != nil {
		return "", _dc
	}
	_bc := make([]byte, n)
	if _edfd := _ed.Read(_ad, _ed.LittleEndian, &_bc); _edfd != nil {
		return "", _edfd
	}
	return string(_bc), nil
}

func PushLeftUI32(v uint32, flag bool) uint32 {
	v >>= 1
	if flag {
		v |= 1 << 31
	}
	return v
}

type Writer struct {
	_gb []byte
	_de int
}

func (_fg *Writer) Cap() int { return cap(_fg._gb) }

func (_cg *Writer) WriteByteAt(b byte, off int) error {
	if off >= len(_cg._gb) {
		return _c.New("\u004f\u0075\u0074\u0020\u006f\u0066\u0020\u0062\u006f\u0075\u006e\u0064\u0073")
	}
	_cg._gb[off] = b
	return nil
}

func (_edg *Writer) curPos() int { return int(_edg.Cap()) - _edg.Len() }

func (_bb *Writer) Write(p []byte) (_ab int, _ca error) {
	_fb, _cfg := _bb.tryGrowByReslice(len(p))
	if !_cfg {
		var _bba error
		_fb, _bba = _bb.grow(len(p))
		if _bba != nil {
			return 0, _bba
		}
	}
	return copy(_bb._gb[_fb:], p), nil
}

func (_bdc *Writer) grow(_cdc int) (int, error) {
	_gf := _bdc.Len()
	if _gf == 0 && _bdc._de != 0 {
		_bdc.reset()
	}
	if _dba, _fa := _bdc.tryGrowByReslice(_cdc); _fa {
		return _dba, nil
	}
	if _bdc._gb == nil && _cdc <= _fc {
		_bdc._gb = make([]byte, _cdc, _fc)
		return 0, nil
	}
	_ded := cap(_bdc._gb)
	if _cdc <= _ded/2-_gf {
		copy(_bdc._gb, _bdc._gb[_bdc._de:])
	} else if _ded > _db-_ded-_cdc {
		return 0, _gbg
	} else {
		_fe := _edfg(2*_ded + _cdc)
		copy(_fe, _bdc._gb[_bdc._de:])
		_bdc._gb = _fe
	}
	_bdc._de = 0
	_bdc._gb = _bdc._gb[:_gf+_cdc]
	return _gf, nil
}

const _db = int(^uint(0) >> 1)

func (_eba *Writer) Len() int { return len(_eba._gb) - _eba._de }

func (_gg *Reader) ReadPairProperty(p interface{}) error {
	if _ged := _gg.align(4); _ged != nil {
		return _ged
	}
	_edf := _e.ValueOf(p)
	for _edf.Kind() == _e.Ptr {
		_edf = _edf.Elem()
	}
	if !_edf.IsValid() {
		return _ag.Errorf("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064")
	}
	if _cbda := _ed.Read(_gg, _ed.LittleEndian, p); _cbda != nil {
		return _cbda
	}
	return nil
}

func (_ec *Writer) WritePropertyNoAlign(a interface{}) error {
	if _cc := _ed.Write(_ec, _ed.LittleEndian, a); _cc != nil {
		return _cc
	}
	return nil
}

func NewReader(b []byte) (*Reader, error) { return &Reader{_g.NewReader(b)}, nil }

func (_adef *Writer) tryGrowByReslice(_gbgd int) (int, bool) {
	if _ba := len(_adef._gb); _gbgd <= cap(_adef._gb)-_ba {
		_adef._gb = _adef._gb[:_ba+_gbgd]
		return _ba, true
	}
	return 0, false
}

func (_ce *Writer) WriteProperty(a interface{}) error {
	if _gga := _ce.align(int(_e.TypeOf(a).Size())); _gga != nil {
		return _gga
	}
	return _ce.WritePropertyNoAlign(a)
}

func (_cba *Writer) AlignLength(alignTo int) error {
	_ef := _cba.Len() % alignTo
	if _ef > 0 {
		_, _ccb := _cba.Write(make([]byte, alignTo-_ef))
		if _ccb != nil {
			return _ccb
		}
	}
	return nil
}

func (_dga *Writer) WriteStringProperty(s string) error {
	_dga.align(4)
	_ggb := []byte(s)
	if _gc := _ed.Write(_dga, _ed.LittleEndian, &_ggb); _gc != nil {
		return _gc
	}
	return nil
}

func (_ae *Writer) Bytes() []byte { return _ae._gb }

func (_eff *Writer) reset() { _eff._gb = _eff._gb[:0]; _eff._de = 0 }

func PopRightUI32(v uint32) (bool, uint32) { return (v & uint32(1)) == 1, v >> 1 }

func (_ge *Reader) ReadProperty(a interface{}) error {
	_bd := _e.ValueOf(a)
	for _bd.Kind() == _e.Ptr {
		_bd = _bd.Elem()
	}
	if !_bd.IsValid() {
		return _ag.Errorf("\u0076a\u006cu\u0065\u0020\u0069\u0073\u0020n\u006f\u0074 \u0076\u0061\u006c\u0069\u0064")
	}
	if _dg := _ge.align(int(_bd.Type().Size())); _dg != nil {
		return _dg
	}
	if _af := _ed.Read(_ge, _ed.LittleEndian, a); _af != nil {
		return _af
	}
	return nil
}

func PopRightUI64(v uint64) (bool, uint64) { return (v & uint64(1)) == 1, v >> 1 }

func (_d *Reader) curPos() int { return int(_d.Size()) - _d.Len() }

func PushLeftUI64(v uint64, flag bool) uint64 {
	v >>= 1
	if flag {
		v |= 1 << 63
	}
	return v
}

func (_be *Reader) align(_cd int) error { return _be.skip((_cd - _be.curPos()%_cd) % _cd) }

func (_bf *Writer) Skip(n int) error {
	if n == 0 {
		return nil
	}
	_, _gbb := _bf.Write(make([]byte, n))
	return _gbb
}

func (_ccbc *Writer) FillWithByte(fillSize int, b byte) error {
	for _agc := 0; _agc < fillSize; _agc++ {
		if _ccg := _ccbc.WritePropertyNoAlign(b); _ccg != nil {
			return _ccg
		}
	}
	return nil
}

type Reader struct{ *_g.Reader }

func (_ggf *Writer) WriteTo(wTo _ac.Writer) (_fca int64, _cfa error) {
	if _ga := _ggf.Len(); _ga > 0 {
		_gce, _ada := wTo.Write(_ggf._gb[_ggf._de:])
		if _gce > _ga {
			return 0, _c.New("\u0072\u0077\u002e\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f\u003a\u0020\u0069\u006e\u0076\u0061\u006ci\u0064\u0020\u0057\u0072\u0069t\u0065\u0020c\u006f\u0075\u006e\u0074")
		}
		_ggf._de += _gce
		_fca = int64(_gce)
		if _ada != nil {
			return _fca, _ada
		}
		if _gce != _ga {
			return _fca, _ac.ErrShortWrite
		}
	}
	_ggf.reset()
	return _fca, nil
}
