package zippkg

import (
	_d "archive/zip"
	_gb "bytes"
	_g "encoding/xml"
	_e "fmt"
	_ac "github.com/arcpd/msword"
	_ga "github.com/arcpd/msword/algo"
	_ea "github.com/arcpd/msword/common/tempstorage"
	_gbg "github.com/arcpd/msword/schema/soo/pkg/relationships"
	_dg "io"
	_dc "path"
	_ba "path/filepath"
	_da "sort"
	_c "strings"
	_a "time"
)

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML(z *_d.Writer, filename string, v interface{}) error {
	_cbf := &_d.FileHeader{}
	_cbf.Method = _d.Deflate
	_cbf.Name = filename
	_cbf.SetModTime(_a.Now())
	_ee, _fde := z.CreateHeader(_cbf)
	if _fde != nil {
		return _e.Errorf("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073", filename, _fde)
	}
	_, _fde = _ee.Write([]byte(XMLHeader))
	if _fde != nil {
		return _e.Errorf("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073", filename, _fde)
	}
	if _fde = _g.NewEncoder(SelfClosingWriter{_ee}).Encode(v); _fde != nil {
		return _e.Errorf("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073", filename, _fde)
	}
	_, _fde = _ee.Write(_afc)
	return _fde
}

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where github.com/arcpd/msword will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func(_f *DecodeMap, _ed, _df string, _fa []*_d.File, _fc *_gbg.Relationship, _db Target) error

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct {
	_dce map[string]Target
	_gbe map[*_gbg.Relationships]string
	_cf  []Target
	_be  OnNewRelationshipFunc
	_cfa map[string]struct{}
	_ff  map[string]int
}

var _ecb = []byte{'/', '>'}

func MarshalXMLByType(z *_d.Writer, dt _ac.DocType, typ string, v interface{}) error {
	_gaf := _ac.AbsoluteFilename(dt, typ, 0)
	return MarshalXML(z, _gaf, v)
}

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk(z *_d.Writer, zipPath, storagePath string) error {
	_gg, _ca := z.Create(zipPath)
	if _ca != nil {
		return _e.Errorf("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073", zipPath, _ca)
	}
	_gc, _ca := _ea.Open(storagePath)
	if _ca != nil {
		return _e.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", storagePath, _ca)
	}
	defer _gc.Close()
	_, _ca = _dg.Copy(_gg, _gc)
	return _ca
}

var _afc = []byte{'\r', '\n'}

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode(f *_d.File, dest interface{}) error {
	_cb, _fe := f.Open()
	if _fe != nil {
		return _e.Errorf("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", f.Name, _fe)
	}
	defer _cb.Close()
	_ag := _g.NewDecoder(_cb)
	if _ge := _ag.Decode(dest); _ge != nil {
		return _e.Errorf("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073", f.Name, _ge)
	}
	if _edg, _daf := dest.(*_gbg.Relationships); _daf {
		for _eg, _ab := range _edg.Relationship {
			switch _ab.TypeAttr {
			case _ac.OfficeDocumentTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.OfficeDocumentType
			case _ac.StylesTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.StylesType
			case _ac.ThemeTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.ThemeType
			case _ac.ControlTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.ControlType
			case _ac.SettingsTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.SettingsType
			case _ac.ImageTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.ImageType
			case _ac.CommentsTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.CommentsType
			case _ac.ThumbnailTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.ThumbnailType
			case _ac.DrawingTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.DrawingType
			case _ac.ChartTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.ChartType
			case _ac.ExtendedPropertiesTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.ExtendedPropertiesType
			case _ac.CustomXMLTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.CustomXMLType
			case _ac.WorksheetTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.WorksheetType
			case _ac.SharedStringsTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.SharedStringsType
			case _ac.TableTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.TableType
			case _ac.HeaderTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.HeaderType
			case _ac.FooterTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.FooterType
			case _ac.NumberingTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.NumberingType
			case _ac.FontTableTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.FontTableType
			case _ac.WebSettingsTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.WebSettingsType
			case _ac.FootNotesTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.FootNotesType
			case _ac.EndNotesTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.EndNotesType
			case _ac.SlideTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.SlideType
			case _ac.VMLDrawingTypeStrict:
				_edg.Relationship[_eg].TypeAttr = _ac.VMLDrawingType
			}
		}
		_da.Slice(_edg.Relationship, func(_eaa, _gfg int) bool {
			_gde := _edg.Relationship[_eaa]
			_gbb := _edg.Relationship[_gfg]
			return _ga.NaturalLess(_gde.IdAttr, _gbb.IdAttr)
		})
	}
	return nil
}

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp(f *_d.File, path string) (string, error) {
	_ffd, _dbc := _ea.TempFile(path, "\u007a\u007a")
	if _dbc != nil {
		return "", _dbc
	}
	defer _ffd.Close()
	_ec, _dbc := f.Open()
	if _dbc != nil {
		return "", _dbc
	}
	defer _ec.Close()
	_, _dbc = _dg.Copy(_ffd, _ec)
	if _dbc != nil {
		return "", _dbc
	}
	return _ffd.Name(), nil
}

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes(z *_d.Writer, zipPath string, data []byte) error {
	_ggb, _fdg := z.Create(zipPath)
	if _fdg != nil {
		return _e.Errorf("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073", zipPath, _fdg)
	}
	_, _fdg = _dg.Copy(_ggb, _gb.NewReader(data))
	return _fdg
}

func (_dfc *DecodeMap) IndexFor(path string) int { return _dfc._ff[path] }

func (_aeg SelfClosingWriter) Write(b []byte) (int, error) {
	_dfb := 0
	_gdc := 0
	for _bg := 0; _bg < len(b)-2; _bg++ {
		if b[_bg] == '>' && b[_bg+1] == '<' && b[_bg+2] == '/' {
			_deg := []byte{}
			_aba := _bg
			for _dgcb := _bg; _dgcb >= 0; _dgcb-- {
				if b[_dgcb] == ' ' {
					_aba = _dgcb
				} else if b[_dgcb] == '<' {
					_deg = b[_dgcb+1 : _aba]
					break
				}
			}
			_abg := []byte{}
			for _fae := _bg + 3; _fae < len(b); _fae++ {
				if b[_fae] == '>' {
					_abg = b[_bg+3 : _fae]
					break
				}
			}
			if !_gb.Equal(_deg, _abg) {
				continue
			}
			_bfe, _aege := _aeg.W.Write(b[_dfb:_bg])
			if _aege != nil {
				return _gdc + _bfe, _aege
			}
			_gdc += _bfe
			_, _aege = _aeg.W.Write(_ecb)
			if _aege != nil {
				return _gdc, _aege
			}
			_gdc += 3
			for _dfg := _bg + 2; _dfg < len(b) && b[_dfg] != '>'; _dfg++ {
				_gdc++
				_dfb = _dfg + 2
				_bg = _dfb
			}
		}
	}
	_cab, _cfe := _aeg.W.Write(b[_dfb:])
	return _cab + _gdc, _cfe
}

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{ W _dg.Writer }

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_af *DecodeMap) SetOnNewRelationshipFunc(fn OnNewRelationshipFunc) { _af._be = fn }

// Decode loops decoding targets registered with AddTarget and calling th
func (_de *DecodeMap) Decode(files []*_d.File) error {
	_cd := 1
	for _cd > 0 {
		for len(_de._cf) > 0 {
			_ce := _de._cf[0]
			_de._cf = _de._cf[1:]
			_gf := _ce.Ifc.(*_gbg.Relationships)
			for _, _gd := range _gf.Relationship {
				_ceg := _de._gbe[_gf]
				if _ba.IsAbs(_gd.TargetAttr) {
					_gd.TargetAttr = _c.TrimPrefix(_gd.TargetAttr, "\u002f")
					if _c.HasPrefix(_gd.TargetAttr, _ceg) {
						_ceg = ""
					}
				}
				_de._be(_de, _ceg+_gd.TargetAttr, _gd.TypeAttr, files, _gd, _ce)
			}
		}
		for _cc, _fgc := range files {
			if _fgc == nil {
				continue
			}
			if _dfa, _fcd := _de._dce[_fgc.Name]; _fcd {
				delete(_de._dce, _fgc.Name)
				if _gbec := Decode(_fgc, _dfa.Ifc); _gbec != nil {
					return _gbec
				}
				files[_cc] = nil
				if _gac, _dgd := _dfa.Ifc.(*_gbg.Relationships); _dgd {
					_de._cf = append(_de._cf, _dfa)
					_gfc, _ := _dc.Split(_dc.Clean(_fgc.Name + "\u002f\u002e\u002e\u002f"))
					_de._gbe[_gac] = _gfc
					_cd++
				}
			}
		}
		_cd--
	}
	return nil
}

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor(path string) string {
	_bf := _c.Split(path, "\u002f")
	_fd := _c.Join(_bf[0:len(_bf)-1], "\u002f")
	_ae := _bf[len(_bf)-1]
	_fd += "\u002f_\u0072\u0065\u006c\u0073\u002f"
	_ae += "\u002e\u0072\u0065l\u0073"
	return _fd + _ae
}

func MarshalXMLByTypeIndex(z *_d.Writer, dt _ac.DocType, typ string, idx int, v interface{}) error {
	_dd := _ac.AbsoluteFilename(dt, typ, idx)
	return MarshalXML(z, _dd, v)
}

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_ef *DecodeMap) AddTarget(filePath string, ifc interface{}, sourceFileType string, idx uint32) bool {
	if _ef._dce == nil {
		_ef._dce = make(map[string]Target)
		_ef._gbe = make(map[*_gbg.Relationships]string)
		_ef._cfa = make(map[string]struct{})
		_ef._ff = make(map[string]int)
	}
	if _dc.IsAbs(filePath) {
		filePath = _c.TrimPrefix(filePath, "\u002f")
	}
	_edf := _dc.Clean(filePath)
	if _, _bd := _ef._cfa[_edf]; _bd {
		return false
	}
	_ef._cfa[_edf] = struct{}{}
	_ef._dce[_edf] = Target{Path: _edf, Typ: sourceFileType, Ifc: ifc, Index: idx}
	return true
}

const XMLHeader = "\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e" + "\u000a"

type Target struct {
	Path  string
	Typ   string
	Ifc   interface{}
	Index uint32
}

func (_dgc *DecodeMap) RecordIndex(path string, idx int) { _dgc._ff[path] = idx }
