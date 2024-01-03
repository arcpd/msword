package axcontrol

import (
	_g "bytes"
	_fe "encoding/binary"
	_be "errors"
	_c "fmt"
	_e "github.com/arcpd/msword/common/tempstorage"
	_ec "github.com/arcpd/msword/internal/mscfb"
	_d "github.com/arcpd/msword/internal/mscfb/rw"
	_ba "github.com/arcpd/msword/schema/schemas.microsoft.com/office/activeX"
	_b "io"
)

// GetForeColor gets a button text color value for a system palette from a commandButton control.
func (_fbg *CommandButtonControl) GetForeColor() uint32 { return _fbg._cbe._ggdd }

func (_becf *ImageControl) readExtraDataBlock(_dbgc *_d.Reader) error {
	_becf._aaa = &imageExtraDataBlock{}
	if _becf._fage._gbd {
		var _dcd uint64
		if _caac := _dbgc.ReadPairProperty(&_dcd); _caac != nil {
			return _caac
		}
		_becf._aaa._faf = uint32(_dcd)
		_becf._aaa._efg = uint32(_dcd >> 32)
	}
	return nil
}

// SetHeight sets height of the Label in HIMETRIC (0.01mm)
func (_dde *LabelControl) SetHeight(height uint32) {
	_dde._gcb._gad = true
	_dde._abeb._gdaa = height
}

type commandButtonExtraDataBlock struct {
	_dad  uint32
	_dfeg uint32
	_fad  string
}

// GetValue gets a value from a control which value can be represented as boolean (on/off).
func (_beed *morphDataControlBoolValue) GetValue() bool { return _beed.getValueBool() }

// GetHeight returns height of the Label in HIMETRIC (0.01mm)
func (_cfde *LabelControl) GetHeight() uint32 { return _cfde._abeb._gdaa }

// GetForeColor gets a button text color value for a system palette from a scrollBar control.
func (_efd *ScrollBarControl) GetForeColor() uint32 { return _efd._beeb._bdfb }

// ComboBoxDataControl is a representation of a combo box ActiveX form.
type ComboBoxDataControl struct{ *morphDataControlStringValue }

func (_deg *ImageControl) writePropMask(_cfe *_d.Writer) error {
	_dfed := uint32(0)
	_dfed >>= 2
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._ebe)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._gega)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._ggf)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._ffg)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._adf)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._bac)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._adb)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._gbd)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._cgg)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._de)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._dff)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._abc)
	_dfed = _d.PushLeftUI32(_dfed, _deg._fage._ddfc)
	_dfed >>= 17
	return _fe.Write(_cfe, _fe.LittleEndian, _dfed)
}

// SetForeColor sets a button text color value from a system palette for a label control.
func (_fgea *LabelControl) SetForeColor(foreColor uint32) {
	_fgea._gcb._ccf = true
	_fgea._ebc._beee = foreColor
}

// SetWidth sets width of the CommandButton in HIMETRIC (0.01mm)
func (_edc *CommandButtonControl) SetWidth(width uint32) {
	_edc._dg._cacd = true
	_edc._ag._dad = width
}

func (_gcbb *SpinButtonControl) writePropMask(_bdbg *_d.Writer) error {
	_ecgf := uint32(0)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._cede)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._aeec)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._fba)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._fbdg)
	_ecgf >>= 1
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._edbfd)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._fce)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._fgbc)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._gaab)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._ebee)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._bgee)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._efee)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._dged)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._cbgd)
	_ecgf = _d.PushLeftUI32(_ecgf, _gcbb._fdbac._fcca)
	_ecgf >>= 17
	return _fe.Write(_bdbg, _fe.LittleEndian, _ecgf)
}

func _gef(_ddf *_ec.Reader) (string, error) {
	_bf, _bag := _ddf.GetEntry("\u0043o\u006d\u0070\u004f\u0062\u006a")
	if _bag != nil {
		return "", _bag
	}
	_gfd := make([]byte, _bf.Size)
	_, _bag = _bf.Read(_gfd)
	if _bag != nil {
		return "", _bag
	}
	_gd := _g.Split(_gfd, []byte("\u002e"))
	if len(_gd) < 2 {
		return "", _be.New("\u0055\u006e\u006bn\u006f\u0077\u006e\u0020\u0066\u006f\u0072\u006d")
	}
	return string(_gd[len(_gd)-2]), nil
}

type morphDataColumnInfoDataBlock struct{}

// GetMax gets a button max value.
func (_ggde *ScrollBarControl) GetMax() int32 { return _ggde._beeb._agcg }

type spinButtonPropMask struct {
	_cede  bool
	_aeec  bool
	_fba   bool
	_fbdg  bool
	_edbfd bool
	_fce   bool
	_fgbc  bool
	_gaab  bool
	_ebee  bool
	_bgee  bool
	_efee  bool
	_dged  bool
	_cbgd  bool
	_fcca  bool
}

// GetWidth returns width of the morphDataControl in HIMETRIC (0.01mm)
func (_fcdg *morphDataControl) GetWidth() uint32 { return _fcdg._eff._ffdf }

// SetCaption sets a caption string for a commandButton control.
func (_ebb *CommandButtonControl) SetCaption(caption string) {
	_ebb._dg._dbcc = true
	_ebb._cbe._gdf = uint32(len(caption))
	_ebb._ag._fad = caption
}

// GetHeight returns height of the morphDataControl in HIMETRIC (0.01mm)
func (_fbf *morphDataControl) GetHeight() uint32 { return _fbf._eff._bcag }

func (_aadb *SpinButtonControl) readPropMask(_eega *_d.Reader) error {
	var _gfdda uint32
	if _ffba := _fe.Read(_eega, _fe.LittleEndian, &_gfdda); _ffba != nil {
		return _ffba
	}
	_aadb._fdbac = &spinButtonPropMask{}
	_aadb._fdbac._cede, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._aeec, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._fba, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._fbdg, _gfdda = _d.PopRightUI32(_gfdda)
	_gfdda >>= 1
	_aadb._fdbac._edbfd, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._fce, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._fgbc, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._gaab, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._ebee, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._bgee, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._efee, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._dged, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._cbgd, _gfdda = _d.PopRightUI32(_gfdda)
	_aadb._fdbac._fcca, _gfdda = _d.PopRightUI32(_gfdda)
	return nil
}

// SetMax sets a button max value.
func (_bada *SpinButtonControl) SetMax(max int32) {
	_bada._fdbac._fce = true
	_bada._deeda._fdfd = max
}

type controlBase struct {
	_fdd uint16
	_eee bool
	_efb *streamData
	_dbe []byte
}

func (_bgf *CommandButtonControl) readPropMask(_add *_d.Reader) error {
	var _gfc uint32
	if _cg := _fe.Read(_add, _fe.LittleEndian, &_gfc); _cg != nil {
		return _cg
	}
	_bgf._dg = &commandButtonPropMask{}
	_bgf._dg._cac, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._ddfd, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._ff, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._dbcc, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._caa, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._cacd, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._aef, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._ddb, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._acde, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._feg, _gfc = _d.PopRightUI32(_gfc)
	_bgf._dg._dcab, _gfc = _d.PopRightUI32(_gfc)
	return nil
}

func (_edb *LabelControl) writeExtraDataBlock(_aec *_d.Writer) error {
	if _edb._ebc._dfaf > 0 {
		if _cae := _aec.WriteStringProperty(_edb._abeb._eed); _cae != nil {
			return _cae
		}
	}
	if _edb._gcb._gad {
		_dbfg := uint64(_edb._abeb._dcf)<<32 | uint64(_edb._abeb._gdaa)
		if _fdgb := _aec.WritePropertyNoAlign(_dbfg); _fdgb != nil {
			return _fdgb
		}
	}
	return nil
}

type morphDataControlBoolValue struct{ *morphDataControl }

func (_cfa *CommandButtonControl) readExtraDataBlock(_eag *_d.Reader) error {
	_cfa._ag = &commandButtonExtraDataBlock{}
	if _cfa._cbe._gdf > 0 {
		_fac, _gacg := _eag.ReadStringProperty(_cfa._cbe._gdf)
		if _gacg != nil {
			return _gacg
		}
		_cfa._ag._fad = _fac
	}
	if _cfa._dg._cacd {
		var _bb uint64
		if _ggg := _eag.ReadPairProperty(&_bb); _ggg != nil {
			return _ggg
		}
		_cfa._ag._dfeg = uint32(_bb)
		_cfa._ag._dad = uint32(_bb >> 32)
	}
	return nil
}

// ControlChoice represents an ActiveX control inside a wrapper.
type ControlChoice struct {
	CheckBox      *CheckBoxDataControl
	TextBox       *TextBoxDataControl
	ListBox       *ListBoxDataControl
	ComboBox      *ComboBoxDataControl
	OptionButton  *OptionButtonDataControl
	ToggleButton  *ToggleButtonDataControl
	Label         *LabelControl
	Image         *ImageControl
	SpinButton    *SpinButtonControl
	CommandButton *CommandButtonControl
	ScrollBar     *ScrollBarControl
	_ggc          *controlBase
}

func (_ggcf *controlBase) getMouseIconBytes() []byte { return _ggcf._efb._egcg.getStdPictureBytes() }

// SetPosition sets a button position value.
func (_ddab *SpinButtonControl) SetPosition(position int32) {
	_ddab._fdbac._fgbc = true
	_ddab._deeda._ecb = position
}

type imageExtraDataBlock struct {
	_efg uint32
	_faf uint32
}

// SpinButtonControl is a representation of a spinButton ActiveX form.
type SpinButtonControl struct {
	controlBase
	_fdbac *spinButtonPropMask
	_deeda *spinButtonDataBlock
	_aeaa  *spinButtonExtraDataBlock
	_bcc   *streamData
}

type morphDataColumnInfoPropMask struct{}

// ExportToByteArray makes a byte array from a control as it is stored in .bin files.
func (_cdf *Control) ExportToByteArray() ([]byte, error) {
	_fb, _bdg := _e.Open(_cdf._gg)
	if _bdg != nil {
		return nil, _bdg
	}
	defer _fb.Close()
	_gab, _bdg := _ec.New(_fb)
	if _bdg != nil {
		return nil, _bdg
	}
	_abd := _d.NewWriter()
	if _gfa := _fe.Write(_abd, _fe.LittleEndian, &_cdf._ed); _gfa != nil {
		return nil, _gfa
	}
	if _dbc := _fe.Write(_abd, _fe.LittleEndian, &_cdf._db); _dbc != nil {
		return nil, _dbc
	}
	if _cde := _fe.Write(_abd, _fe.LittleEndian, uint16(0)); _cde != nil {
		return nil, _cde
	}
	var _ca error
	if _gb := _cdf.Choice; _gb != nil {
		if _gb.CheckBox != nil {
			_ca = _gb.CheckBox.export(_abd)
		} else if _gb.TextBox != nil {
			_ca = _gb.TextBox.export(_abd)
		} else if _gb.ComboBox != nil {
			_ca = _gb.ComboBox.export(_abd)
		} else if _gb.ListBox != nil {
			_ca = _gb.ListBox.export(_abd)
		} else if _gb.OptionButton != nil {
			_ca = _gb.OptionButton.export(_abd)
		} else if _gb.ToggleButton != nil {
			_ca = _gb.ToggleButton.export(_abd)
		} else if _gb.Label != nil {
			_ca = _gb.Label.export(_abd)
		} else if _gb.SpinButton != nil {
			_ca = _gb.SpinButton.export(_abd)
		} else if _gb.CommandButton != nil {
			_ca = _gb.CommandButton.export(_abd)
		} else if _gb.ScrollBar != nil {
			_ca = _gb.ScrollBar.export(_abd)
		} else {
			_ca = _gb._ggc.writeTheRest(_abd)
		}
	}
	if _ca != nil {
		return nil, _ca
	}
	if _acg := _abd.WriteByteAt(byte(_cdf._a), 2); _acg != nil {
		return nil, _acg
	}
	if _dd := _abd.WriteByteAt(byte(_cdf._a>>8), 3); _dd != nil {
		return nil, _dd
	}
	_ea, _bdg := _gab.GetEntry("\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
	if _bdg != nil {
		return nil, _bdg
	}
	if _ggb := _ea.SetEntryContent(_abd.Bytes()); _ggb != nil {
		return nil, _ggb
	}
	return _gab.Export()
}

// ImportFromFile makes a Control from a file in a storage.
func ImportFromFile(storagePath string) (*Control, error) {
	_bec, _gc := _e.Open(storagePath)
	if _gc != nil {
		return nil, _gc
	}
	defer _bec.Close()
	_gcd, _gc := _ec.New(_bec)
	if _gc != nil {
		return nil, _gc
	}
	_ggcc, _gc := _gef(_gcd)
	if _gc != nil {
		return nil, _gc
	}
	_ad, _gc := _gcd.GetEntry("\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0073")
	if _gc != nil {
		return nil, _gc
	}
	_ge := make([]byte, _ad.Size)
	_, _gc = _ad.Read(_ge)
	if _gc != nil {
		return nil, _gc
	}
	_bc := &Control{_gg: storagePath}
	_da, _gc := _d.NewReader(_ge)
	if _gc != nil {
		return nil, _gc
	}
	if _cc := _fe.Read(_da, _fe.LittleEndian, &_bc._ed); _cc != nil {
		return nil, _cc
	}
	if _gf := _fe.Read(_da, _fe.LittleEndian, &_bc._db); _gf != nil {
		return nil, _gf
	}
	if _bd := _fe.Read(_da, _fe.LittleEndian, &_bc._a); _bd != nil {
		return nil, _bd
	}
	switch _ggcc {
	case "\u0043\u0068\u0065\u0063\u006b\u0042\u006f\u0078":
		_fa, _ga := _cbcd(_da)
		if _ga != nil {
			return nil, _ga
		}
		_fa._eee = true
		_bc.Choice = &ControlChoice{CheckBox: &CheckBoxDataControl{&morphDataControlBoolValue{_fa}}}
	case "\u0054e\u0078\u0074\u0042\u006f\u0078":
		_fg, _cf := _cbcd(_da)
		if _cf != nil {
			return nil, _cf
		}
		_fg._eee = true
		_bc.Choice = &ControlChoice{TextBox: &TextBoxDataControl{&morphDataControlStringValue{_fg}}}
	case "\u0043\u006f\u006d\u0062\u006f\u0042\u006f\u0078":
		_bee, _gag := _cbcd(_da)
		if _gag != nil {
			return nil, _gag
		}
		_bee._eee = true
		_bc.Choice = &ControlChoice{ComboBox: &ComboBoxDataControl{&morphDataControlStringValue{_bee}}}
	case "\u004ci\u0073\u0074\u0042\u006f\u0078":
		_ee, _bef := _cbcd(_da)
		if _bef != nil {
			return nil, _bef
		}
		_ee._eee = true
		_bc.Choice = &ControlChoice{ListBox: &ListBoxDataControl{&morphDataControlStringValue{_ee}}}
	case "\u004f\u0070\u0074i\u006f\u006e\u0042\u0075\u0074\u0074\u006f\u006e":
		_ac, _cd := _cbcd(_da)
		if _cd != nil {
			return nil, _cd
		}
		_ac._eee = true
		_bc.Choice = &ControlChoice{OptionButton: &OptionButtonDataControl{&morphDataControlBoolValue{_ac}}}
	case "\u0054\u006f\u0067g\u006c\u0065\u0042\u0075\u0074\u0074\u006f\u006e":
		_eb, _gff := _cbcd(_da)
		if _gff != nil {
			return nil, _gff
		}
		_eb._eee = true
		_bc.Choice = &ControlChoice{ToggleButton: &ToggleButtonDataControl{&morphDataControlBoolValue{_eb}}}
	case "\u004c\u0061\u0062e\u006c":
		_befd, _eg := _cdfb(_da)
		if _eg != nil {
			return nil, _eg
		}
		_bc.Choice = &ControlChoice{Label: _befd}
	case "\u0053\u0070\u0069\u006e\u0042\u0075\u0074\u0074\u006f\u006e":
		_bg, _daf := _bgad(_da)
		if _daf != nil {
			return nil, _daf
		}
		_bc.Choice = &ControlChoice{SpinButton: _bg}
	case "\u0043\u006f\u006d\u006d\u0061\u006e\u0064\u0042\u0075\u0074\u0074\u006f\u006e":
		_fc, _ab := _ggd(_da)
		if _ab != nil {
			return nil, _ab
		}
		_bc.Choice = &ControlChoice{CommandButton: _fc}
	case "\u0053c\u0072\u006f\u006c\u006c\u0042\u0061r":
		_gac, _dbg := _acff(_da)
		if _dbg != nil {
			return nil, _dbg
		}
		_bc.Choice = &ControlChoice{ScrollBar: _gac}
	default:
		_cce := &controlBase{}
		if _fgg := _cce.readTheRest(_da); _fgg != nil {
			return nil, _fgg
		}
		_bc.Choice = &ControlChoice{_ggc: _cce}
	}
	return _bc, nil
}

// GetCaption gets a caption string from a label control.
func (_ffbc *LabelControl) GetCaption() string { return _ffbc._abeb._eed }

func (_cdff *ScrollBarControl) writeExtraDataBlock(_ded *_d.Writer) error {
	if _cdff._babc._afac {
		_eadc := uint64(_cdff._fdfg._cfeg)<<32 | uint64(_cdff._fdfg._gceb)
		if _caebg := _ded.WritePropertyNoAlign(_eadc); _caebg != nil {
			return _caebg
		}
	}
	return nil
}

// TextBoxDataControl is a representation of a text box ActiveX form.
type TextBoxDataControl struct{ *morphDataControlStringValue }

func (_ggff *morphDataControl) writeColumnInfoPropMask(_ddbb *_d.Writer) error { return nil }

// GetCaption gets a caption string from a commandButton control.
func (_fd *CommandButtonControl) GetCaption() string { return _fd._ag._fad }

// CommandButtonControl is a representation of a commandButton ActiveX form.
type CommandButtonControl struct {
	controlBase
	_dg  *commandButtonPropMask
	_cbe *commandButtonDataBlock
	_ag  *commandButtonExtraDataBlock
	_dca *streamData
}

type morphDataControl struct {
	controlBase
	_feb  *morphDataPropMask
	_efa  *morphDataDataBlock
	_eff  *morphDataExtraDataBlock
	_fff  *streamData
	_ccac *morphDataColumnInfo
	_fdf  *morphDataColumnInfoPropMask
	_bggg *morphDataColumnInfoDataBlock
}

// SetValue sets a value for a control which value can be represented as boolean (on/off).
func (_caeb *morphDataControlBoolValue) SetValue(value bool) { _caeb.setValueBool(value) }

// GetWidth returns width of the SpinButton in HIMETRIC (0.01mm)
func (_gegd *SpinButtonControl) GetWidth() uint32 { return _gegd._aeaa._eede }

// GetWidth returns width of the CommandButton in HIMETRIC (0.01mm)
func (_ddc *CommandButtonControl) GetWidth() uint32 { return _ddc._ag._dad }

// SetMax sets a button max value.
func (_eccg *ScrollBarControl) SetMax(max int32) {
	_eccg._babc._edcef = true
	_eccg._beeb._agcg = max
}

func (_gfac *morphDataControl) readPropMask(_cgabg *_d.Reader) error {
	var _bfbdb uint64
	if _bagc := _fe.Read(_cgabg, _fe.LittleEndian, &_bfbdb); _bagc != nil {
		return _bagc
	}
	_gfac._feb = &morphDataPropMask{}
	_gfac._feb._gcce, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._afdd, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._def, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._gged, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._geec, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._bfc, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._bfbd, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._bbc, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._aeef, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._fca, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._agbc, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._dbff, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._gfdg, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._bdgf, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._ffga, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._cggb, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._babf, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._eggg, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._beeeb, _bfbdb = _d.PopRightUI64(_bfbdb)
	_bfbdb >>= 1
	_gfac._feb._ffag, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._adbb, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._fcc, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._bcec, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._dbbc, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._addg, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._dda, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._cceg, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._bdca, _bfbdb = _d.PopRightUI64(_bfbdb)
	_gfac._feb._effb, _bfbdb = _d.PopRightUI64(_bfbdb)
	_bfbdb >>= 1
	_bfbdb >>= 1
	_gfac._feb._edbg, _bfbdb = _d.PopRightUI64(_bfbdb)
	return nil
}

func (_cdc *morphDataControl) readDataBlock(_aeg *_d.Reader) error {
	_cdc._efa = &morphDataDataBlock{}
	if _cdc._feb._gcce {
		if _deed := _aeg.ReadProperty(&_cdc._efa._ece); _deed != nil {
			return _deed
		}
	}
	if _cdc._feb._afdd {
		if _dab := _aeg.ReadProperty(&_cdc._efa._aefc); _dab != nil {
			return _dab
		}
	}
	if _cdc._feb._def {
		if _agae := _aeg.ReadProperty(&_cdc._efa._fdbab); _agae != nil {
			return _agae
		}
	}
	if _cdc._feb._gged {
		if _cdaa := _aeg.ReadProperty(&_cdc._efa._cbf); _cdaa != nil {
			return _cdaa
		}
	}
	if _cdc._feb._geec {
		if _bdff := _aeg.ReadProperty(&_cdc._efa._dcc); _bdff != nil {
			return _bdff
		}
	}
	if _cdc._feb._bfc {
		if _fagf := _aeg.ReadProperty(&_cdc._efa._adbe); _fagf != nil {
			return _fagf
		}
	}
	if _cdc._feb._bfbd {
		if _ade := _aeg.ReadProperty(&_cdc._efa._cadg); _ade != nil {
			return _ade
		}
	}
	if _cdc._feb._bbc {
		if _cbcg := _aeg.ReadProperty(&_cdc._efa._aed); _cbcg != nil {
			return _cbcg
		}
	}
	if _cdc._feb._fca {
		if _eae := _aeg.ReadProperty(&_cdc._efa._adffb); _eae != nil {
			return _eae
		}
	}
	if _cdc._feb._agbc {
		if _fcge := _aeg.ReadProperty(&_cdc._efa._fcac); _fcge != nil {
			return _fcge
		}
	}
	if _cdc._feb._dbff {
		if _egggb := _aeg.ReadProperty(&_cdc._efa._dag); _egggb != nil {
			return _egggb
		}
	}
	if _cdc._feb._gfdg {
		if _cag := _aeg.ReadProperty(&_cdc._efa._cfb); _cag != nil {
			return _cag
		}
	}
	if _cdc._feb._bdgf {
		if _aeca := _aeg.ReadProperty(&_cdc._efa._fgb); _aeca != nil {
			return _aeca
		}
	}
	if _cdc._feb._ffga {
		if _agfg := _aeg.ReadProperty(&_cdc._efa._febe); _agfg != nil {
			return _agfg
		}
	}
	if _cdc._feb._cggb {
		if _gcdeb := _aeg.ReadProperty(&_cdc._efa._egb); _gcdeb != nil {
			return _gcdeb
		}
	}
	if _cdc._feb._babf {
		if _gbba := _aeg.ReadProperty(&_cdc._efa._gfde); _gbba != nil {
			return _gbba
		}
	}
	if _cdc._feb._eggg {
		if _dfea := _aeg.ReadProperty(&_cdc._efa._dbgcb); _dfea != nil {
			return _dfea
		}
	}
	if _cdc._feb._beeeb {
		if _cdce := _aeg.ReadProperty(&_cdc._efa._bbb); _cdce != nil {
			return _cdce
		}
	}
	if _cdc._feb._ffag {
		if _bced := _aeg.ReadProperty(&_cdc._efa._eadab); _bced != nil {
			return _bced
		}
	}
	if _cdc._feb._adbb {
		if _edd := _aeg.ReadProperty(&_cdc._efa._bcdb); _edd != nil {
			return _edd
		}
	}
	if _cdc._feb._fcc {
		var _cegb uint32
		if _bdd := _aeg.ReadProperty(&_cegb); _bdd != nil {
			return _bdd
		}
		_cdc._efa._dbee, _cdc._efa._cfed = _ae(_cegb)
	}
	if _cdc._feb._bcec {
		var _gggc uint32
		if _decg := _aeg.ReadProperty(&_gggc); _decg != nil {
			return _decg
		}
		_cdc._efa._deeg, _cdc._efa._gba = _ae(_gggc)
	}
	if _cdc._feb._dbbc {
		if _eac := _aeg.ReadProperty(&_cdc._efa._fab); _eac != nil {
			return _eac
		}
	}
	if _cdc._feb._addg {
		if _gcba := _aeg.ReadProperty(&_cdc._efa._gcfb); _gcba != nil {
			return _gcba
		}
	}
	if _cdc._feb._dda {
		if _cfc := _aeg.ReadProperty(&_cdc._efa._ceg); _cfc != nil {
			return _cfc
		}
	}
	if _cdc._feb._cceg {
		if _gbfe := _aeg.ReadProperty(&_cdc._efa._dadc); _gbfe != nil {
			return _gbfe
		}
	}
	if _cdc._feb._bdca {
		if _bfbf := _aeg.ReadProperty(&_cdc._efa._cebg); _bfbf != nil {
			return _bfbf
		}
	}
	if _cdc._feb._effb {
		if _ace := _aeg.ReadProperty(&_cdc._efa._gdgf); _ace != nil {
			return _ace
		}
	}
	if _cdc._feb._edbg {
		var _cded uint32
		if _dbeb := _aeg.ReadProperty(&_cded); _dbeb != nil {
			return _dbeb
		}
		_cdc._efa._beeda, _cdc._efa._fcacf = _ae(_cded)
	}
	return nil
}

// GetBackColor gets a button text color value for a system palette from a commandButton control.
func (_bcd *CommandButtonControl) GetBackColor() uint32 { return _bcd._cbe._adc }

// GetBackColor gets a button text color value for a system palette from a scrollBar control.
func (_agg *ScrollBarControl) GetBackColor() uint32 { return _agg._beeb._daag }

func _cb(_gfe uint32, _gbf bool) uint32 {
	if _gfe == 0 {
		return 0
	}
	if _gbf {
		_gfe |= 1 << 31
	}
	return _gfe
}

func (_ebfe *controlBase) readStreamDataM(_baag *_d.Reader, _ccbg bool) error {
	_ebfe._efb = &streamData{}
	if _ccbg {
		_ebfe._efb._egcg = &guidAndPicture{}
		if _caad := _ebfe._efb._egcg.importFromReader(_baag); _caad != nil {
			return _caad
		}
	}
	return nil
}

type streamData struct {
	_egcg *guidAndPicture
	_egef *guidAndPicture
}

func (_bcaac *controlBase) writeStreamDataPM(_cfda *_d.Writer) error {
	if _bcaac._efb != nil {
		if _bcaac._efb._egef != nil {
			if _cgfb := _bcaac._efb._egef.export(_cfda); _cgfb != nil {
				return _cgfb
			}
		}
		if _bcaac._efb._egcg != nil {
			if _gegg := _bcaac._efb._egcg.export(_cfda); _gegg != nil {
				return _gegg
			}
		}
	}
	return nil
}

func (_gcc *CommandButtonControl) export(_aee *_d.Writer) error {
	if _bdb := _gcc.writePropMask(_aee); _bdb != nil {
		return _bdb
	}
	if _acd := _gcc.writeDataBlock(_aee); _acd != nil {
		return _acd
	}
	if _ef := _gcc.writeExtraDataBlock(_aee); _ef != nil {
		return _ef
	}
	_gcc._fdd = uint16(_aee.Len() - 4)
	if _fag := _gcc.writeStreamDataPM(_aee); _fag != nil {
		return _fag
	}
	return _gcc.writeTheRest(_aee)
}

type labelExtraDataBlock struct {
	_dcf  uint32
	_gdaa uint32
	_eed  string
}

func (_bed *guidAndPicture) importFromReader(_fgfd *_d.Reader) error {
	_agff := _g.NewBuffer([]byte{})
	if _, _cbc := _b.CopyN(_agff, _fgfd, int64(_bfb)); _cbc != nil {
		return _cbc
	}
	_cfac := _agff.Bytes()
	for _dgbc := 0; _dgbc < _bfb; _dgbc++ {
		if _cfac[_dgbc] != _eeef[_dgbc] {
			return _c.Errorf("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0047\u0055\u0049\u0044\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0025v\u0020\u0061\u0074\u0020\u0069\u006e\u0064e\u0078\u0020\u0025\u0076", _cfac[_dgbc], _dgbc)
		}
	}
	_bed._gcf = _cfac
	_bed._fcg = &stdPicture{}
	if _cced := _fgfd.ReadProperty(&_bed._fcg._gdc); _cced != nil {
		return _cced
	}
	if _bed._fcg._gdc != _bff {
		return _c.Errorf("\u0049\u006e\u0076\u0061li\u0064\u0020\u0070\u0072\u0065\u0061\u006d\u0062\u006c\u0065\u0020\u0025\u0076", _bed._fcg._gdc)
	}
	if _gfdb := _fgfd.ReadProperty(&_bed._fcg._ada); _gfdb != nil {
		return _gfdb
	}
	if _bed._fcg._ada != 0 {
		_abed := _g.NewBuffer([]byte{})
		if _, _gdd := _b.Copy(_abed, _fgfd); _gdd != nil {
			return _gdd
		}
		_bed._fcg._gagc = _abed
	}
	return nil
}

// SetWidth sets width of the morphDataControl in HIMETRIC (0.01mm)
func (_ddbe *morphDataControl) SetWidth(width uint32) {
	_ddbe._feb._aeef = true
	_ddbe._eff._ffdf = width
}

// GetPosition gets a button position value.
func (_faef *ScrollBarControl) GetPosition() int32 { return _faef._beeb._cggd }

func (_gdaf *morphDataControl) setValueBool(_ccg bool) {
	_gdaf._feb._fcc = true
	_gdaf._efa._cfed = 1
	_gdaf._efa._dbee = true
	_bca := "\u0030"
	if _ccg {
		_bca = "\u0031"
	}
	_gdaf._eff._befa = _bca
}

func (_adff *ImageControl) writeDataBlock(_eca *_d.Writer) error {
	if _adff._fage._gega {
		if _ecf := _eca.WriteProperty(_adff._bffg._fddf); _ecf != nil {
			return _ecf
		}
	}
	if _adff._fage._ggf {
		if _bgg := _eca.WriteProperty(_adff._bffg._geaf); _bgg != nil {
			return _bgg
		}
	}
	if _adff._fage._ffg {
		if _fdbc := _eca.WriteProperty(_adff._bffg._bbd); _fdbc != nil {
			return _fdbc
		}
	}
	if _adff._fage._adf {
		if _geea := _eca.WriteProperty(_adff._bffg._adfe); _geea != nil {
			return _geea
		}
	}
	if _adff._fage._bac {
		if _cfdd := _eca.WriteProperty(_adff._bffg._gfcf); _cfdd != nil {
			return _cfdd
		}
	}
	if _adff._fage._adb {
		if _geac := _eca.WriteProperty(_adff._bffg._bagf); _geac != nil {
			return _geac
		}
	}
	if _adff._fage._cgg {
		if _gbe := _eca.WriteProperty(_adff._bffg._ddfe); _gbe != nil {
			return _gbe
		}
	}
	if _adff._fage._de {
		if _daa := _eca.WriteProperty(_adff._bffg._eba); _daa != nil {
			return _daa
		}
	}
	if _adff._fage._abc {
		if _adadd := _eca.WriteProperty(_adff._bffg._dfa); _adadd != nil {
			return _adadd
		}
	}
	if _adff._fage._ddfc {
		if _bab := _eca.WriteProperty(_adff._bffg._dec); _bab != nil {
			return _bab
		}
	}
	return _eca.AlignLength(4)
}

func _acff(_fbca *_d.Reader) (*ScrollBarControl, error) {
	_eaca := &ScrollBarControl{}
	if _eafb := _eaca.readPropMask(_fbca); _eafb != nil {
		return nil, _eafb
	}
	if _bbbd := _eaca.readDataBlock(_fbca); _bbbd != nil {
		return nil, _bbbd
	}
	if _ggdeb := _eaca.readExtraDataBlock(_fbca); _ggdeb != nil {
		return nil, _ggdeb
	}
	if _cegd := _eaca.readStreamDataM(_fbca, _eaca._babc._gafe); _cegd != nil {
		return nil, _cegd
	}
	if _agac := _eaca.readTheRest(_fbca); _agac != nil {
		return nil, _agac
	}
	return _eaca, nil
}

// ToggleButtonDataControl is a representation of a toggle button ActiveX form.
type ToggleButtonDataControl struct{ *morphDataControlBoolValue }

const _bbgb = 6

// GetMin gets a button min value.
func (_caf *SpinButtonControl) GetMin() int32 { return _caf._deeda._dbea }

func (_gdb *ScrollBarControl) readExtraDataBlock(_effg *_d.Reader) error {
	_gdb._fdfg = &scrollBarExtraDataBlock{}
	if _gdb._babc._afac {
		var _dea uint64
		if _feda := _effg.ReadPairProperty(&_dea); _feda != nil {
			return _feda
		}
		_gdb._fdfg._gceb = uint32(_dea)
		_gdb._fdfg._cfeg = uint32(_dea >> 32)
	}
	return nil
}

// SetHeight sets height of the ScrollBar in HIMETRIC (0.01mm)
func (_aadd *ScrollBarControl) SetHeight(height uint32) {
	_aadd._babc._afac = true
	_aadd._fdfg._gceb = height
}

func _bgad(_egeb *_d.Reader) (*SpinButtonControl, error) {
	_gcab := &SpinButtonControl{}
	if _ccae := _gcab.readPropMask(_egeb); _ccae != nil {
		return nil, _ccae
	}
	if _gbg := _gcab.readDataBlock(_egeb); _gbg != nil {
		return nil, _gbg
	}
	if _agbg := _gcab.readExtraDataBlock(_egeb); _agbg != nil {
		return nil, _agbg
	}
	if _bdaf := _gcab.readStreamDataM(_egeb, _gcab._fdbac._cbgd); _bdaf != nil {
		return nil, _bdaf
	}
	if _cab := _gcab.readTheRest(_egeb); _cab != nil {
		return nil, _cab
	}
	return _gcab, nil
}

// GetBackColor gets a button text color value for a system palette from a label control.
func (_feaa *LabelControl) GetBackColor() uint32 { return _feaa._ebc._dga }

type morphDataPropMask struct {
	_gcce  bool
	_afdd  bool
	_def   bool
	_gged  bool
	_geec  bool
	_bfc   bool
	_bfbd  bool
	_bbc   bool
	_aeef  bool
	_fca   bool
	_agbc  bool
	_dbff  bool
	_gfdg  bool
	_bdgf  bool
	_ffga  bool
	_cggb  bool
	_babf  bool
	_eggg  bool
	_beeeb bool
	_ffag  bool
	_adbb  bool
	_fcc   bool
	_bcec  bool
	_dbbc  bool
	_addg  bool
	_dda   bool
	_cceg  bool
	_bdca  bool
	_effb  bool
	_edbg  bool
}

func (_acf *morphDataControl) readExtraDataBlock(_effba *_d.Reader) error {
	_acf._eff = &morphDataExtraDataBlock{}
	if _acf._feb._aeef {
		var _ddgdf uint64
		if _dacg := _effba.ReadPairProperty(&_ddgdf); _dacg != nil {
			return _dacg
		}
		_acf._eff._bcag = uint32(_ddgdf)
		_acf._eff._ffdf = uint32(_ddgdf >> 32)
	}
	if _acf._efa._cfed > 0 {
		_dcbf, _fadbb := _effba.ReadStringProperty(_acf._efa._cfed)
		if _fadbb != nil {
			return _fadbb
		}
		_acf._eff._befa = _dcbf
	}
	if _acf._efa._gba > 0 {
		_baf, _abebc := _effba.ReadStringProperty(_acf._efa._gba)
		if _abebc != nil {
			return _abebc
		}
		_acf._eff._cdea = _baf
	}
	if _acf._efa._fcacf > 0 {
		_dcdc, _ced := _effba.ReadStringProperty(_acf._efa._fcacf)
		if _ced != nil {
			return _ced
		}
		_acf._eff._cgae = _dcdc
	}
	return nil
}

const _cfaf = 4

var _eeef = []byte{0x04, 0x52, 0xE3, 0x0B, 0x91, 0x8F, 0xCE, 0x11, 0x9D, 0xE3, 0, 0xAA, 0, 0x4B, 0xB8, 0x51}

type morphDataColumnInfo struct{}

// SetCaption sets a caption string for a label control.
func (_cbed *LabelControl) SetCaption(caption string) {
	_cbed._gcb._cgga = true
	_cbed._ebc._dfaf = uint32(len(caption))
	_cbed._abeb._eed = caption
}

func (_afad *morphDataControl) getValueBool() bool {
	return _afad._feb._fcc && _afad._eff._befa == "\u0031"
}

// GetHeight returns height of the ScrollBar in HIMETRIC (0.01mm)
func (_cedf *ScrollBarControl) GetHeight() uint32 { return _cedf._fdfg._gceb }

// OptionButtonDataControl is a representation of an option button ActiveX form.
type OptionButtonDataControl struct{ *morphDataControlBoolValue }

func (_ccad *SpinButtonControl) readDataBlock(_defa *_d.Reader) error {
	_ccad._deeda = &spinButtonDataBlock{}
	if _ccad._fdbac._cede {
		if _fcdb := _defa.ReadProperty(&_ccad._deeda._fagff); _fcdb != nil {
			return _fcdb
		}
	}
	if _ccad._fdbac._aeec {
		if _gcabb := _defa.ReadProperty(&_ccad._deeda._abfba); _gcabb != nil {
			return _gcabb
		}
	}
	if _ccad._fdbac._fba {
		if _cbee := _defa.ReadProperty(&_ccad._deeda._dbbe); _cbee != nil {
			return _cbee
		}
	}
	if _ccad._fdbac._edbfd {
		if _eadd := _defa.ReadProperty(&_ccad._deeda._dbea); _eadd != nil {
			return _eadd
		}
	}
	if _ccad._fdbac._fce {
		if _geag := _defa.ReadProperty(&_ccad._deeda._fdfd); _geag != nil {
			return _geag
		}
	}
	if _ccad._fdbac._fgbc {
		if _cef := _defa.ReadProperty(&_ccad._deeda._ecb); _cef != nil {
			return _cef
		}
	}
	if _ccad._fdbac._gaab {
		if _dacc := _defa.ReadProperty(&_ccad._deeda._cegg); _dacc != nil {
			return _dacc
		}
	}
	if _ccad._fdbac._ebee {
		if _bgb := _defa.ReadProperty(&_ccad._deeda._gede); _bgb != nil {
			return _bgb
		}
	}
	if _ccad._fdbac._bgee {
		if _ffce := _defa.ReadProperty(&_ccad._deeda._agaeg); _ffce != nil {
			return _ffce
		}
	}
	if _ccad._fdbac._efee {
		if _bage := _defa.ReadProperty(&_ccad._deeda._adfed); _bage != nil {
			return _bage
		}
	}
	if _ccad._fdbac._dged {
		if _fccb := _defa.ReadProperty(&_ccad._deeda._fbcf); _fccb != nil {
			return _fccb
		}
	}
	if _ccad._fdbac._cbgd {
		if _feee := _defa.ReadProperty(&_ccad._deeda._bcac); _feee != nil {
			return _feee
		}
	}
	if _ccad._fdbac._fcca {
		if _cfce := _defa.ReadProperty(&_ccad._deeda._bcaf); _cfce != nil {
			return _cfce
		}
	}
	return nil
}

func (_fcf *LabelControl) readDataBlock(_gge *_d.Reader) error {
	_fcf._ebc = &labelDataBlock{}
	if _fcf._gcb._ccf {
		if _bdgd := _gge.ReadProperty(&_fcf._ebc._beee); _bdgd != nil {
			return _bdgd
		}
	}
	if _fcf._gcb._edcc {
		if _bae := _gge.ReadProperty(&_fcf._ebc._dga); _bae != nil {
			return _bae
		}
	}
	if _fcf._gcb._bad {
		if _gcda := _gge.ReadProperty(&_fcf._ebc._ebf); _gcda != nil {
			return _gcda
		}
	}
	if _fcf._gcb._cgga {
		var _adg uint32
		if _cdgd := _gge.ReadProperty(&_adg); _cdgd != nil {
			return _cdgd
		}
		_fcf._ebc._eeae, _fcf._ebc._dfaf = _ae(_adg)
	}
	if _fcf._gcb._aca {
		if _fgc := _gge.ReadProperty(&_fcf._ebc._daff); _fgc != nil {
			return _fgc
		}
	}
	if _fcf._gcb._begb {
		if _gfdd := _gge.ReadProperty(&_fcf._ebc._gagb); _gfdd != nil {
			return _gfdd
		}
	}
	if _fcf._gcb._egg {
		if _bgfg := _gge.ReadProperty(&_fcf._ebc._gecf); _bgfg != nil {
			return _bgfg
		}
	}
	if _fcf._gcb._gfdf {
		if _bdad := _gge.ReadProperty(&_fcf._ebc._dee); _bdad != nil {
			return _bdad
		}
	}
	if _fcf._gcb._cgec {
		if _babg := _gge.ReadProperty(&_fcf._ebc._bea); _babg != nil {
			return _babg
		}
	}
	if _fcf._gcb._caacd {
		if _dbb := _gge.ReadProperty(&_fcf._ebc._bfd); _dbb != nil {
			return _dbb
		}
	}
	if _fcf._gcb._fgdf {
		if _dge := _gge.ReadProperty(&_fcf._ebc._fadb); _dge != nil {
			return _dge
		}
	}
	if _fcf._gcb._aad {
		if _fbbd := _gge.ReadProperty(&_fcf._ebc._agc); _fbbd != nil {
			return _fbbd
		}
	}
	return nil
}

var _bfb = len(_eeef)

func (_adfc *morphDataControl) readColumnInfo(_fadda *_d.Reader) error { return nil }

// FmPictureSizeMode represents one of the three picture size modes according to MS-OFORMS document.
type FmPictureSizeMode byte

// SetPosition sets a button position value.
func (_dgbb *ScrollBarControl) SetPosition(position int32) {
	_dgbb._babc._aefa = true
	_dgbb._beeb._cggd = position
}

func (_gfce *controlBase) readStreamDataPM(_cdbf *_d.Reader, _bbaf, _afcdg bool) error {
	_gfce._efb = &streamData{}
	if _bbaf {
		_gfce._efb._egef = &guidAndPicture{}
		if _abab := _gfce._efb._egef.importFromReader(_cdbf); _abab != nil {
			return _abab
		}
	}
	if _afcdg {
		_gfce._efb._egcg = &guidAndPicture{}
		if _ecbf := _gfce._efb._egcg.importFromReader(_cdbf); _ecbf != nil {
			return _ecbf
		}
	}
	return nil
}

func _ae(_cad uint32) (bool, uint32) {
	if _cad == 0 {
		return false, 0
	}
	_cdg := _cad >= 1<<31
	if _cdg {
		_cad -= 1 << 31
	}
	return _cdg, _cad
}

// SetBackColor sets a button text color value from a system palette for a commandButton control.
func (_fef *CommandButtonControl) SetBackColor(backColor uint32) {
	_fef._dg._ddfd = true
	_fef._cbe._adc = backColor
}

func (_daca *SpinButtonControl) export(_dgg *_d.Writer) error {
	if _ddcf := _daca.writePropMask(_dgg); _ddcf != nil {
		return _ddcf
	}
	if _gccf := _daca.writeDataBlock(_dgg); _gccf != nil {
		return _gccf
	}
	if _dcfg := _daca.writeExtraDataBlock(_dgg); _dcfg != nil {
		return _dcfg
	}
	_daca._fdd = uint16(_dgg.Len() - 4)
	if _dce := _daca.writeStreamDataM(_dgg); _dce != nil {
		return _dce
	}
	return _daca.writeTheRest(_dgg)
}

func (_dbfb *ImageControl) export(_ega *_d.Writer) error {
	if _ffb := _dbfb.writePropMask(_ega); _ffb != nil {
		return _ffb
	}
	if _dcb := _dbfb.writeDataBlock(_ega); _dcb != nil {
		return _dcb
	}
	if _befg := _dbfb.writeExtraDataBlock(_ega); _befg != nil {
		return _befg
	}
	_dbfb._fdd = uint16(_ega.Len() - 4)
	if _ddcda := _dbfb.writeStreamDataPM(_ega); _ddcda != nil {
		return _ddcda
	}
	return _dbfb.writeTheRest(_ega)
}

func (_fge *CommandButtonControl) readDataBlock(_edcg *_d.Reader) error {
	_fge._cbe = &commandButtonDataBlock{}
	if _fge._dg._cac {
		if _dfe := _edcg.ReadProperty(&_fge._cbe._ggdd); _dfe != nil {
			return _dfe
		}
	}
	if _fge._dg._ddfd {
		if _ccd := _edcg.ReadProperty(&_fge._cbe._adc); _ccd != nil {
			return _ccd
		}
	}
	if _fge._dg._ff {
		if _gec := _edcg.ReadProperty(&_fge._cbe._ebg); _gec != nil {
			return _gec
		}
	}
	if _fge._dg._dbcc {
		var _edce uint32
		if _dbga := _edcg.ReadProperty(&_edce); _dbga != nil {
			return _dbga
		}
		_fge._cbe._acc, _fge._cbe._gdf = _ae(_edce)
	}
	if _fge._dg._caa {
		if _bagd := _edcg.ReadProperty(&_fge._cbe._fae); _bagd != nil {
			return _bagd
		}
	}
	if _fge._dg._aef {
		if _bge := _edcg.ReadProperty(&_fge._cbe._bgff); _bge != nil {
			return _bge
		}
	}
	if _fge._dg._ddb {
		if _dbgd := _edcg.ReadProperty(&_fge._cbe._gfeb); _dbgd != nil {
			return _dbgd
		}
	}
	if _fge._dg._acde {
		if _cgd := _edcg.ReadProperty(&_fge._cbe._agf); _cgd != nil {
			return _cgd
		}
	}
	if _fge._dg._dcab {
		if _fea := _edcg.ReadProperty(&_fge._cbe._df); _fea != nil {
			return _fea
		}
	}
	return nil
}

// SetValue sets a value for a control which value can be represented as a string.
func (_ecfb *morphDataControlStringValue) SetValue(value string) { _ecfb.setValueString(value) }

type morphDataControlStringValue struct{ *morphDataControl }

// SetForeColor sets a button text color value from a system palette for a scrollBar control.
func (_fecg *ScrollBarControl) SetForeColor(foreColor uint32) {
	_fecg._babc._face = true
	_fecg._beeb._bdfb = foreColor
}

func (_abde *LabelControl) writeDataBlock(_bdcd *_d.Writer) error {
	if _abde._gcb._ccf {
		if _ccee := _bdcd.WriteProperty(_abde._ebc._beee); _ccee != nil {
			return _ccee
		}
	}
	if _abde._gcb._edcc {
		if _aeae := _bdcd.WriteProperty(_abde._ebc._dga); _aeae != nil {
			return _aeae
		}
	}
	if _abde._gcb._bad {
		if _cdgdb := _bdcd.WriteProperty(_abde._ebc._ebf); _cdgdb != nil {
			return _cdgdb
		}
	}
	if _abde._gcb._cgga {
		_bfdg := _cb(_abde._ebc._dfaf, _abde._ebc._eeae)
		if _gdg := _bdcd.WriteProperty(_bfdg); _gdg != nil {
			return _gdg
		}
	}
	if _abde._gcb._aca {
		if _adbf := _bdcd.WriteProperty(_abde._ebc._daff); _adbf != nil {
			return _adbf
		}
	}
	if _abde._gcb._begb {
		if _dbbd := _bdcd.WriteProperty(_abde._ebc._gagb); _dbbd != nil {
			return _dbbd
		}
	}
	if _abde._gcb._egg {
		if _dgfe := _bdcd.WriteProperty(_abde._ebc._gecf); _dgfe != nil {
			return _dgfe
		}
	}
	if _abde._gcb._gfdf {
		if _fcd := _bdcd.WriteProperty(_abde._ebc._dee); _fcd != nil {
			return _fcd
		}
	}
	if _abde._gcb._cgec {
		if _bdbf := _bdcd.WriteProperty(_abde._ebc._bea); _bdbf != nil {
			return _bdbf
		}
	}
	if _abde._gcb._caacd {
		if _ffa := _bdcd.WriteProperty(_abde._ebc._bfd); _ffa != nil {
			return _ffa
		}
	}
	if _abde._gcb._fgdf {
		if _cbaa := _bdcd.WriteProperty(_abde._ebc._fadb); _cbaa != nil {
			return _cbaa
		}
	}
	if _abde._gcb._aad {
		if _bfde := _bdcd.WriteProperty(_abde._ebc._agc); _bfde != nil {
			return _bfde
		}
	}
	return _bdcd.AlignLength(4)
}

const _ffbf = 3

// SetWidth sets width of the ScrollBar in HIMETRIC (0.01mm)
func (_adgf *ScrollBarControl) SetWidth(width uint32) {
	_adgf._babc._afac = true
	_adgf._fdfg._cfeg = width
}

type commandButtonPropMask struct {
	_cac  bool
	_ddfd bool
	_ff   bool
	_dbcc bool
	_caa  bool
	_cacd bool
	_aef  bool
	_ddb  bool
	_acde bool
	_feg  bool
	_dcab bool
}

func (_fdg *CommandButtonControl) writeDataBlock(_dcac *_d.Writer) error {
	if _fdg._dg._cac {
		if _cgdd := _dcac.WriteProperty(_fdg._cbe._ggdd); _cgdd != nil {
			return _cgdd
		}
	}
	if _fdg._dg._ddfd {
		if _fbb := _dcac.WriteProperty(_fdg._cbe._adc); _fbb != nil {
			return _fbb
		}
	}
	if _fdg._dg._ff {
		if _fgge := _dcac.WriteProperty(_fdg._cbe._ebg); _fgge != nil {
			return _fgge
		}
	}
	if _fdg._dg._dbcc {
		_dgf := _cb(_fdg._cbe._gdf, _fdg._cbe._acc)
		if _cgf := _dcac.WriteProperty(_dgf); _cgf != nil {
			return _cgf
		}
	}
	if _fdg._dg._caa {
		if _gfab := _dcac.WriteProperty(_fdg._cbe._fae); _gfab != nil {
			return _gfab
		}
	}
	if _fdg._dg._aef {
		if _fgd := _dcac.WriteProperty(_fdg._cbe._bgff); _fgd != nil {
			return _fgd
		}
	}
	if _fdg._dg._ddb {
		if _faea := _dcac.WriteProperty(_fdg._cbe._gfeb); _faea != nil {
			return _faea
		}
	}
	if _fdg._dg._acde {
		if _bgef := _dcac.WriteProperty(_fdg._cbe._agf); _bgef != nil {
			return _bgef
		}
	}
	if _fdg._dg._dcab {
		if _ffd := _dcac.WriteProperty(_fdg._cbe._df); _ffd != nil {
			return _ffd
		}
	}
	return _dcac.AlignLength(4)
}

// SetWidth sets width of the SpinButton in HIMETRIC (0.01mm)
func (_dbca *SpinButtonControl) SetWidth(width uint32) {
	_dbca._fdbac._fbdg = true
	_dbca._aeaa._eede = width
}

func (_fbcc *morphDataControl) writeColumnInfo(_bdag *_d.Writer) error { return nil }

func (_gcff *morphDataControl) readColumnInfoPropMask(_cec *_d.Reader) error { return nil }

func (_dffa *morphDataControl) export(_gae *_d.Writer) error {
	if _gdag := _dffa.writePropMask(_gae); _gdag != nil {
		return _gdag
	}
	if _baad := _dffa.writeDataBlock(_gae); _baad != nil {
		return _baad
	}
	if _afdf := _dffa.writeExtraDataBlock(_gae); _afdf != nil {
		return _afdf
	}
	_dffa._fdd = uint16(_gae.Len() - 4)
	if _gca := _dffa.writeStreamDataMP(_gae); _gca != nil {
		return _gca
	}
	return _dffa.writeTheRest(_gae)
}

func (_aeb *controlBase) readStreamDataMP(_ceba *_d.Reader, _abgd, _afaf bool) error {
	_aeb._efb = &streamData{}
	if _abgd {
		_aeb._efb._egcg = &guidAndPicture{}
		if _faaa := _aeb._efb._egcg.importFromReader(_ceba); _faaa != nil {
			return _faaa
		}
	}
	if _afaf {
		_aeb._efb._egef = &guidAndPicture{}
		if _ccgcc := _aeb._efb._egef.importFromReader(_ceba); _ccgcc != nil {
			return _ccgcc
		}
	}
	return nil
}

func (_acce *controlBase) setPictureBytes(_bged []byte) {
	if _acce._efb._egef == nil {
		_acce._efb._egef = &guidAndPicture{}
	}
	_acce._efb._egef.setStdPictureBytes(_bged)
}

// GetForeColor gets a button text color value for a system palette from a spinButton control.
func (_begf *SpinButtonControl) GetForeColor() uint32 { return _begf._deeda._fagff }

// SetForeColor sets a button text color value from a system palette for a spinButton control.
func (_eegfe *SpinButtonControl) SetForeColor(foreColor uint32) {
	_eegfe._fdbac._cede = true
	_eegfe._deeda._fagff = foreColor
}

// SetHeight sets height of the morphDataControl in HIMETRIC (0.01mm)
func (_efbe *morphDataControl) SetHeight(height uint32) {
	_efbe._feb._aeef = true
	_efbe._eff._bcag = height
}

func (_bdfc *morphDataControl) getValueString() string {
	if _bdfc._feb._fcc {
		return _bdfc._eff._befa
	}
	return ""
}

// SetBackColor sets a button text color value from a system palette for a label control.
func (_ccc *LabelControl) SetBackColor(backColor uint32) {
	_ccc._gcb._edcc = true
	_ccc._ebc._dga = backColor
}

const (
	FmPictureSizeModeClip FmPictureSizeMode = iota
	FmPictureSizeModeStretch
	_
	FmPictureSizeModeZoom
)

type spinButtonExtraDataBlock struct {
	_eede uint32
	_badg uint32
}

// GetBackColor gets a button text color value for a system palette from a spinButton control.
func (_dede *SpinButtonControl) GetBackColor() uint32 { return _dede._deeda._abfba }

func (_gfg *CommandButtonControl) writeExtraDataBlock(_bdf *_d.Writer) error {
	if _gfg._cbe._gdf > 0 {
		if _dfb := _bdf.WriteStringProperty(_gfg._ag._fad); _dfb != nil {
			return _dfb
		}
	}
	if _gfg._dg._cacd {
		_fagc := uint64(_gfg._ag._dad)<<32 | uint64(_gfg._ag._dfeg)
		if _faa := _bdf.WritePropertyNoAlign(_fagc); _faa != nil {
			return _faa
		}
	}
	return nil
}

// GetWidth returns width of the Label in HIMETRIC (0.01mm)
func (_dbaa *LabelControl) GetWidth() uint32 { return _dbaa._abeb._dcf }

// GetHeight returns height of the CommandButton in HIMETRIC (0.01mm)
func (_dbd *CommandButtonControl) GetHeight() uint32 { return _dbd._ag._dfeg }

func (_acfeg *controlBase) writeStreamDataMP(_feddd *_d.Writer) error {
	if _acfeg._efb != nil {
		if _acfeg._efb._egcg != nil {
			if _bcfg := _acfeg._efb._egcg.export(_feddd); _bcfg != nil {
				return _bcfg
			}
		}
		if _acfeg._efb._egef != nil {
			if _abga := _acfeg._efb._egef.export(_feddd); _abga != nil {
				return _abga
			}
		}
	}
	return nil
}

// ListBoxDataControl is a representation of a list box ActiveX form.
type ListBoxDataControl struct{ *morphDataControlStringValue }

func (_ebae *LabelControl) writePropMask(_gce *_d.Writer) error {
	_deb := uint32(0)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._ccf)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._edcc)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._bad)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._cgga)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._aca)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._gad)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._begb)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._egg)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._gfdf)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._cgec)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._caacd)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._fgdf)
	_deb = _d.PushLeftUI32(_deb, _ebae._gcb._aad)
	_deb >>= 19
	return _fe.Write(_gce, _fe.LittleEndian, _deb)
}

type imageDataBlock struct {
	_fddf uint32
	_geaf uint32
	_bbd  byte
	_adfe uint8
	_gfcf byte
	_bagf byte
	_ddfe uint16
	_eba  byte
	_dfa  uint32
	_dec  uint16
}

func (_afc *morphDataControl) readColumnInfoDataBlock(_babb *_d.Reader) error { return nil }

func (_cgcd *SpinButtonControl) readExtraDataBlock(_ccdc *_d.Reader) error {
	_cgcd._aeaa = &spinButtonExtraDataBlock{}
	if _cgcd._fdbac._fbdg {
		var _ccgd uint64
		if _ggdg := _ccdc.ReadPairProperty(&_ccgd); _ggdg != nil {
			return _ggdg
		}
		_cgcd._aeaa._badg = uint32(_ccgd)
		_cgcd._aeaa._eede = uint32(_ccgd >> 32)
	}
	return nil
}

func (_edg *controlBase) writeStreamDataM(_gdbg *_d.Writer) error {
	if _edg._efb != nil {
		if _edg._efb._egcg != nil {
			if _agd := _edg._efb._egcg.export(_gdbg); _agd != nil {
				return _agd
			}
		}
	}
	return nil
}

// SetBackColor sets a button text color value from a system palette for a spinButton control.
func (_acag *SpinButtonControl) SetBackColor(backColor uint32) {
	_acag._fdbac._aeec = true
	_acag._deeda._abfba = backColor
}

func (_ceb *LabelControl) readPropMask(_cfdef *_d.Reader) error {
	var _caacdg uint32
	if _fdc := _fe.Read(_cfdef, _fe.LittleEndian, &_caacdg); _fdc != nil {
		return _fdc
	}
	_ceb._gcb = &labelPropMask{}
	_ceb._gcb._ccf, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._edcc, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._bad, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._cgga, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._aca, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._gad, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._begb, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._egg, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._gfdf, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._cgec, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._caacd, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._fgdf, _caacdg = _d.PopRightUI32(_caacdg)
	_ceb._gcb._aad, _caacdg = _d.PopRightUI32(_caacdg)
	return nil
}

type scrollBarExtraDataBlock struct {
	_cfeg uint32
	_gceb uint32
}

type stdPicture struct {
	_gdc  uint32
	_ada  uint32
	_gagc *_g.Buffer
}

type morphDataExtraDataBlock struct {
	_ffdf uint32
	_bcag uint32
	_befa string
	_cdea string
	_cgae string
}

func (_cga *morphDataControl) setValueString(_eda string) {
	_cga._feb._fcc = true
	_cga._efa._cfed = uint32(len(_eda))
	_cga._eff._befa = _eda
}

func (_dba *controlBase) setMouseIconBytes(_dgb []byte) {
	if _dba._efb._egcg == nil {
		_dba._efb._egcg = &guidAndPicture{}
	}
	_dba._efb._egcg.setStdPictureBytes(_dgb)
}

func (_aea *controlBase) readTheRest(_eccc *_d.Reader) error {
	_acdc := _g.NewBuffer([]byte{})
	_, _eea := _b.Copy(_acdc, _eccc)
	if _eea != nil {
		return _eea
	}
	_aea._dbe = _acdc.Bytes()
	return nil
}

type labelPropMask struct {
	_ccf   bool
	_edcc  bool
	_bad   bool
	_cgga  bool
	_aca   bool
	_gad   bool
	_begb  bool
	_egg   bool
	_gfdf  bool
	_cgec  bool
	_caacd bool
	_fgdf  bool
	_aad   bool
}

func (_ddgd *LabelControl) export(_abfd *_d.Writer) error {
	if _cbae := _ddgd.writePropMask(_abfd); _cbae != nil {
		return _cbae
	}
	if _bdgg := _ddgd.writeDataBlock(_abfd); _bdgg != nil {
		return _bdgg
	}
	if _cfef := _ddgd.writeExtraDataBlock(_abfd); _cfef != nil {
		return _cfef
	}
	_ddgd._fdd = uint16(_abfd.Len() - 4)
	if _afa := _ddgd.writeStreamDataPM(_abfd); _afa != nil {
		return _afa
	}
	return _ddgd.writeTheRest(_abfd)
}

func (_gcdf *guidAndPicture) setStdPictureBytes(_fdba []byte) {
	if _fdba == nil {
		return
	}
	if _gcdf._gcf == nil {
		_gcdf._gcf = _eeef
	}
	if _gcdf._fcg == nil {
		_gcdf._fcg = &stdPicture{_gdc: _bff}
	}
	_gcdf._fcg._ada = uint32(len(_fdba))
	_gcdf._fcg._gagc = _g.NewBuffer(_fdba)
}

func (_af *controlBase) writeTheRest(_efe *_d.Writer) error {
	_, _geg := _efe.Write(_af._dbe)
	return _geg
}

const (
	FmPictureAlignmentTopLeft FmPictureAlignment = iota
	FmPictureAlignmentTopRight
	FmPictureAlignmentCenter
	FmPictureAlignmentBottomLeft
	FmPictureAlignmentBottomRight
)

type morphDataDataBlock struct {
	_ece   uint32
	_aefc  uint32
	_fdbab uint32
	_cbf   uint32
	_dcc   byte
	_adbe  byte
	_cadg  byte
	_aed   byte
	_adffb rune
	_fcac  uint32
	_dag   uint16
	_cfb   int16
	_fgb   int16
	_febe  uint16
	_egb   uint16
	_gfde  uint8
	_dbgcb uint8
	_bbb   uint8
	_eadab uint8
	_bcdb  uint8
	_cfed  uint32
	_dbee  bool
	_gba   uint32
	_deeg  bool
	_fab   uint32
	_gcfb  uint32
	_ceg   uint32
	_dadc  uint16
	_cebg  uint16
	_gdgf  rune
	_fcacf uint32
	_beeda bool
}

func (_dccf *SpinButtonControl) writeExtraDataBlock(_dgdf *_d.Writer) error {
	if _dccf._fdbac._fbdg {
		_cfdc := uint64(_dccf._aeaa._eede)<<32 | uint64(_dccf._aeaa._badg)
		if _feaf := _dgdf.WritePropertyNoAlign(_cfdc); _feaf != nil {
			return _feaf
		}
	}
	return nil
}

// GetPosition gets a button position value.
func (_gfaf *SpinButtonControl) GetPosition() int32 { return _gfaf._deeda._ecb }

func (_eegf *morphDataControl) writeDataBlock(_fadd *_d.Writer) error {
	if _eegf._feb._gcce {
		if _fegb := _fadd.WriteProperty(_eegf._efa._ece); _fegb != nil {
			return _fegb
		}
	}
	if _eegf._feb._afdd {
		if _cccc := _fadd.WriteProperty(_eegf._efa._aefc); _cccc != nil {
			return _cccc
		}
	}
	if _eegf._feb._def {
		if _fagfb := _fadd.WriteProperty(_eegf._efa._fdbab); _fagfb != nil {
			return _fagfb
		}
	}
	if _eegf._feb._gged {
		if _abca := _fadd.WriteProperty(_eegf._efa._cbf); _abca != nil {
			return _abca
		}
	}
	if _eegf._feb._geec {
		if _agef := _fadd.WriteProperty(_eegf._efa._dcc); _agef != nil {
			return _agef
		}
	}
	if _eegf._feb._bfc {
		if _bfe := _fadd.WriteProperty(_eegf._efa._adbe); _bfe != nil {
			return _bfe
		}
	}
	if _eegf._feb._bfbd {
		if _bga := _fadd.WriteProperty(_eegf._efa._cadg); _bga != nil {
			return _bga
		}
	}
	if _eegf._feb._bbc {
		if _eagc := _fadd.WriteProperty(_eegf._efa._aed); _eagc != nil {
			return _eagc
		}
	}
	if _eegf._feb._fca {
		if _cgbg := _fadd.WriteProperty(_eegf._efa._adffb); _cgbg != nil {
			return _cgbg
		}
	}
	if _eegf._feb._agbc {
		if _gfgb := _fadd.WriteProperty(_eegf._efa._fcac); _gfgb != nil {
			return _gfgb
		}
	}
	if _eegf._feb._dbff {
		if _daba := _fadd.WriteProperty(_eegf._efa._dag); _daba != nil {
			return _daba
		}
	}
	if _eegf._feb._gfdg {
		if _egab := _fadd.WriteProperty(_eegf._efa._cfb); _egab != nil {
			return _egab
		}
	}
	if _eegf._feb._bdgf {
		if _dagg := _fadd.WriteProperty(_eegf._efa._fgb); _dagg != nil {
			return _dagg
		}
	}
	if _eegf._feb._ffga {
		if _bcf := _fadd.WriteProperty(_eegf._efa._febe); _bcf != nil {
			return _bcf
		}
	}
	if _eegf._feb._cggb {
		if _fgba := _fadd.WriteProperty(_eegf._efa._egb); _fgba != nil {
			return _fgba
		}
	}
	if _eegf._feb._babf {
		if _gfed := _fadd.WriteProperty(_eegf._efa._gfde); _gfed != nil {
			return _gfed
		}
	}
	if _eegf._feb._eggg {
		if _gdce := _fadd.WriteProperty(_eegf._efa._dbgcb); _gdce != nil {
			return _gdce
		}
	}
	if _eegf._feb._beeeb {
		if _acae := _fadd.WriteProperty(_eegf._efa._bbb); _acae != nil {
			return _acae
		}
	}
	if _eegf._feb._ffag {
		if _cbedb := _fadd.WriteProperty(_eegf._efa._eadab); _cbedb != nil {
			return _cbedb
		}
	}
	if _eegf._feb._adbb {
		if _aded := _fadd.WriteProperty(_eegf._efa._bcdb); _aded != nil {
			return _aded
		}
	}
	if _eegf._feb._fcc {
		_fffg := _cb(_eegf._efa._cfed, _eegf._efa._dbee)
		if _ccaf := _fadd.WriteProperty(_fffg); _ccaf != nil {
			return _ccaf
		}
	}
	if _eegf._feb._bcec {
		_ggce := _cb(_eegf._efa._gba, _eegf._efa._deeg)
		if _ggda := _fadd.WriteProperty(_ggce); _ggda != nil {
			return _ggda
		}
	}
	if _eegf._feb._dbbc {
		if _gbbe := _fadd.WriteProperty(_eegf._efa._fab); _gbbe != nil {
			return _gbbe
		}
	}
	if _eegf._feb._addg {
		if _cgc := _fadd.WriteProperty(_eegf._efa._gcfb); _cgc != nil {
			return _cgc
		}
	}
	if _eegf._feb._dda {
		if _acb := _fadd.WriteProperty(_eegf._efa._ceg); _acb != nil {
			return _acb
		}
	}
	if _eegf._feb._cceg {
		if _cgag := _fadd.WriteProperty(_eegf._efa._dadc); _cgag != nil {
			return _cgag
		}
	}
	if _eegf._feb._bdca {
		if _adac := _fadd.WriteProperty(_eegf._efa._cebg); _adac != nil {
			return _adac
		}
	}
	if _eegf._feb._effb {
		if _gabg := _fadd.WriteProperty(_eegf._efa._gdgf); _gabg != nil {
			return _gabg
		}
	}
	if _eegf._feb._edbg {
		_abb := _cb(_eegf._efa._fcacf, _eegf._efa._beeda)
		if _fbc := _fadd.WriteProperty(_abb); _fbc != nil {
			return _fbc
		}
	}
	return _fadd.AlignLength(4)
}

// SetCaption sets a caption string for a morph control.
func (_bce *morphDataControl) SetCaption(caption string) {
	if _bce._eee {
		_bce._feb._bcec = true
		_bce._efa._gba = uint32(len(caption))
		_bce._eff._cdea = caption
	}
}

func (_cfba *morphDataControl) writeExtraDataBlock(_fbbb *_d.Writer) error {
	if _cfba._feb._aeef {
		_efgf := uint64(_cfba._eff._ffdf)<<32 | uint64(_cfba._eff._bcag)
		if _aff := _fbbb.WritePropertyNoAlign(_efgf); _aff != nil {
			return _aff
		}
	}
	if _cfba._efa._cfed > 0 {
		if _bdde := _fbbb.WriteStringProperty(_cfba._eff._befa); _bdde != nil {
			return _bdde
		}
	}
	if _cfba._efa._gba > 0 {
		if _bcaa := _fbbb.WriteStringProperty(_cfba._eff._cdea); _bcaa != nil {
			return _bcaa
		}
	}
	if _cfba._efa._fcacf > 0 {
		if _ggec := _fbbb.WriteStringProperty(_cfba._eff._cgae); _ggec != nil {
			return _ggec
		}
	}
	return nil
}

// GetForeColor gets a button text color value for a system palette from a label control.
func (_dega *LabelControl) GetForeColor() uint32 { return _dega._ebc._beee }

type scrollBarDataBlock struct {
	_bdfb uint32
	_daag uint32
	_ceda uint32
	_dgcg uint8
	_ecg  int32
	_agcg int32
	_cggd int32
	_fda  int32
	_abac int32
	_cfag int32
	_dgfb int32
	_ddba uint32
	_gcg  uint32
	_dae  uint32
	_bbgd uint16
}

func _bbf(_gaa *_d.Reader) (*ImageControl, error) {
	_ce := &ImageControl{}
	if _gea := _ce.readPropMask(_gaa); _gea != nil {
		return nil, _gea
	}
	if _eaf := _ce.readDataBlock(_gaa); _eaf != nil {
		return nil, _eaf
	}
	if _bda := _ce.readExtraDataBlock(_gaa); _bda != nil {
		return nil, _bda
	}
	if _cfd := _ce.readStreamDataPM(_gaa, _ce._fage._cgg, _ce._fage._ddfc); _cfd != nil {
		return nil, _cfd
	}
	if _bgd := _ce.readTheRest(_gaa); _bgd != nil {
		return nil, _bgd
	}
	return _ce, nil
}

type guidAndPicture struct {
	_gcf []byte
	_fcg *stdPicture
}

// ScrollBarControl is a representation of a scrollBar ActiveX form.
type ScrollBarControl struct {
	controlBase
	_babc *scrollBarPropMask
	_beeb *scrollBarDataBlock
	_fdfg *scrollBarExtraDataBlock
	_cedb *streamData
}

func (_ggca *LabelControl) readExtraDataBlock(_baa *_d.Reader) error {
	_ggca._abeb = &labelExtraDataBlock{}
	if _ggca._ebc._dfaf > 0 {
		_dcbd, _ecd := _baa.ReadStringProperty(_ggca._ebc._dfaf)
		if _ecd != nil {
			return _ecd
		}
		_ggca._abeb._eed = _dcbd
	}
	if _ggca._gcb._gad {
		var _bgde uint64
		if _ecfa := _baa.ReadPairProperty(&_bgde); _ecfa != nil {
			return _ecfa
		}
		_ggca._abeb._gdaa = uint32(_bgde)
		_ggca._abeb._dcf = uint32(_bgde >> 32)
	}
	return nil
}

type labelDataBlock struct {
	_beee uint32
	_dga  uint32
	_ebf  uint32
	_dfaf uint32
	_eeae bool
	_daff uint32
	_gagb uint8
	_gecf uint32
	_dee  uint16
	_bea  uint16
	_bfd  uint16
	_fadb rune
	_agc  uint16
}

const _abda = 5

// FmPictureAlignment represents one of the five picture aignments according to MS-OFORMS document.
type FmPictureAlignment byte

func (_ddff *ImageControl) readPropMask(_fga *_d.Reader) error {
	var _cgdda uint32
	if _afg := _fe.Read(_fga, _fe.LittleEndian, &_cgdda); _afg != nil {
		return _afg
	}
	_ddff._fage = &imagePropMask{}
	_cgdda >>= 2
	_ddff._fage._ebe, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._gega, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._ggf, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._ffg, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._adf, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._bac, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._adb, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._gbd, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._cgg, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._de, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._dff, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._abc, _cgdda = _d.PopRightUI32(_cgdda)
	_ddff._fage._ddfc, _cgdda = _d.PopRightUI32(_cgdda)
	return nil
}

func (_fbd *guidAndPicture) export(_adad *_d.Writer) error {
	if _, _cca := _b.Copy(_adad, _g.NewBuffer(_fbd._gcf)); _cca != nil {
		return _cca
	}
	if _fec := _adad.WriteProperty(_fbd._fcg._gdc); _fec != nil {
		return _fec
	}
	if _dbf := _adad.WriteProperty(_fbd._fcg._ada); _dbf != nil {
		return _dbf
	}
	_, _fdb := _b.Copy(_adad, _fbd._fcg._gagc)
	if _fdb != nil {
		return _fdb
	}
	return nil
}

const _cgab = 1

type commandButtonDataBlock struct {
	_ggdd uint32
	_adc  uint32
	_ebg  uint32
	_gdf  uint32
	_acc  bool
	_fae  uint32
	_bgff uint8
	_gfeb uint16
	_agf  rune
	_df   uint16
}

func (_ddg *ImageControl) writeExtraDataBlock(_aga *_d.Writer) error {
	if _ddg._fage._gbd {
		_cfad := uint64(_ddg._aaa._efg)<<32 | uint64(_ddg._aaa._faf)
		if _geeg := _aga.WritePropertyNoAlign(_cfad); _geeg != nil {
			return _geeg
		}
	}
	return nil
}

// ImageControl is a representation of an image ActiveX form.
type ImageControl struct {
	controlBase
	_fage *imagePropMask
	_bffg *imageDataBlock
	_aaa  *imageExtraDataBlock
	_fdgg *streamData
}

type spinButtonDataBlock struct {
	_fagff uint32
	_abfba uint32
	_dbbe  uint32
	_dbea  int32
	_fdfd  int32
	_ecb   int32
	_cegg  int32
	_gede  int32
	_agaeg int32
	_adfed uint32
	_fbcf  uint32
	_bcac  uint16
	_bcaf  uint8
}

// SetBackColor sets a button text color value from a system palette for a scrollBar control.
func (_fcdgg *ScrollBarControl) SetBackColor(backColor uint32) {
	_fcdgg._babc._eeea = true
	_fcdgg._beeb._daag = backColor
}

const _bff uint32 = 0x0000746C

// GetCaption gets a caption string from a morph control.
func (_dcbe *morphDataControl) GetCaption() string {
	if _dcbe._eee && _dcbe._feb._bcec {
		return _dcbe._eff._cdea
	}
	return ""
}

// GetMax gets a button max value.
func (_acceg *SpinButtonControl) GetMax() int32 { return _acceg._deeda._fdfd }

func (_gaga *guidAndPicture) getStdPictureBytes() []byte {
	if _gaga != nil && _gaga._fcg != nil && _gaga._fcg._gagc != nil {
		return _gaga._fcg._gagc.Bytes()
	}
	return nil
}

// GetWidth returns width of the ScrollBar in HIMETRIC (0.01mm)
func (_caacb *ScrollBarControl) GetWidth() uint32 { return _caacb._fdfg._cfeg }

// GetHeight returns height of the SpinButton in HIMETRIC (0.01mm)
func (_ddbad *SpinButtonControl) GetHeight() uint32 { return _ddbad._aeaa._badg }

type imagePropMask struct {
	_ebe  bool
	_gega bool
	_ggf  bool
	_ffg  bool
	_adf  bool
	_bac  bool
	_adb  bool
	_gbd  bool
	_cgg  bool
	_de   bool
	_dff  bool
	_abc  bool
	_ddfc bool
}

func (_bbg *ImageControl) readDataBlock(_cda *_d.Reader) error {
	_bbg._bffg = &imageDataBlock{}
	if _bbg._fage._gega {
		if _bbdb := _cda.ReadProperty(&_bbg._bffg._fddf); _bbdb != nil {
			return _bbdb
		}
	}
	if _bbg._fage._ggf {
		if _dcaa := _cda.ReadProperty(&_bbg._bffg._geaf); _dcaa != nil {
			return _dcaa
		}
	}
	if _bbg._fage._ffg {
		if _cge := _cda.ReadProperty(&_bbg._bffg._bbd); _cge != nil {
			return _cge
		}
	}
	if _bbg._fage._adf {
		if _efc := _cda.ReadProperty(&_bbg._bffg._adfe); _efc != nil {
			return _efc
		}
	}
	if _bbg._fage._bac {
		if _bdc := _cda.ReadProperty(&_bbg._bffg._gfcf); _bdc != nil {
			return _bdc
		}
	}
	if _bbg._fage._adb {
		if _fefc := _cda.ReadProperty(&_bbg._bffg._bagf); _fefc != nil {
			return _fefc
		}
	}
	if _bbg._fage._cgg {
		if _egf := _cda.ReadProperty(&_bbg._bffg._ddfe); _egf != nil {
			return _egf
		}
	}
	if _bbg._fage._de {
		if _cfg := _cda.ReadProperty(&_bbg._bffg._eba); _cfg != nil {
			return _cfg
		}
	}
	if _bbg._fage._abc {
		if _cba := _cda.ReadProperty(&_bbg._bffg._dfa); _cba != nil {
			return _cba
		}
	}
	if _bbg._fage._ddfc {
		if _ddfg := _cda.ReadProperty(&_bbg._bffg._dec); _ddfg != nil {
			return _ddfg
		}
	}
	return nil
}

func _cbcd(_beeg *_d.Reader) (*morphDataControl, error) {
	_adgg := &morphDataControl{}
	if _egc := _adgg.readPropMask(_beeg); _egc != nil {
		return nil, _egc
	}
	if _fdcg := _adgg.readDataBlock(_beeg); _fdcg != nil {
		return nil, _fdcg
	}
	if _cgb := _adgg.readExtraDataBlock(_beeg); _cgb != nil {
		return nil, _cgb
	}
	if _ege := _adgg.readStreamDataMP(_beeg, _adgg._feb._cceg, _adgg._feb._bdca); _ege != nil {
		return nil, _ege
	}
	switch _adgg._efa._cadg {
	case _bega, _ffbf:
		if _eada := _adgg.readColumnInfo(_beeg); _eada != nil {
			return nil, _eada
		}
		if _bfa := _adgg.readColumnInfoPropMask(_beeg); _bfa != nil {
			return nil, _bfa
		}
		if _age := _adgg.readColumnInfoDataBlock(_beeg); _age != nil {
			return nil, _age
		}
	}
	if _afd := _adgg.readTheRest(_beeg); _afd != nil {
		return nil, _afd
	}
	return _adgg, nil
}

// SetWidth sets width of the Label in HIMETRIC (0.01mm)
func (_gaf *LabelControl) SetWidth(width uint32) { _gaf._gcb._gad = true; _gaf._abeb._dcf = width }

func (_geeb *ScrollBarControl) writeDataBlock(_aaddb *_d.Writer) error {
	if _geeb._babc._face {
		if _geb := _aaddb.WriteProperty(_geeb._beeb._bdfb); _geb != nil {
			return _geb
		}
	}
	if _geeb._babc._eeea {
		if _dadf := _aaddb.WriteProperty(_geeb._beeb._daag); _dadf != nil {
			return _dadf
		}
	}
	if _geeb._babc._fcad {
		if _decb := _aaddb.WriteProperty(_geeb._beeb._ceda); _decb != nil {
			return _decb
		}
	}
	if _geeb._babc._faga {
		if _cdge := _aaddb.WriteProperty(_geeb._beeb._dgcg); _cdge != nil {
			return _cdge
		}
	}
	if _geeb._babc._cee {
		if _abg := _aaddb.WriteProperty(_geeb._beeb._ecg); _abg != nil {
			return _abg
		}
	}
	if _geeb._babc._edcef {
		if _edbf := _aaddb.WriteProperty(_geeb._beeb._agcg); _edbf != nil {
			return _edbf
		}
	}
	if _geeb._babc._aefa {
		if _egd := _aaddb.WriteProperty(_geeb._beeb._cggd); _egd != nil {
			return _egd
		}
	}
	if _geeb._babc._cff {
		if _dcff := _aaddb.WriteProperty(_geeb._beeb._fda); _dcff != nil {
			return _dcff
		}
	}
	if _geeb._babc._fedd {
		if _ecec := _aaddb.WriteProperty(_geeb._beeb._abac); _ecec != nil {
			return _ecec
		}
	}
	if _geeb._babc._ddac {
		if _ggcb := _aaddb.WriteProperty(_geeb._beeb._cfag); _ggcb != nil {
			return _ggcb
		}
	}
	if _geeb._babc._abfb {
		if _ebbb := _aaddb.WriteProperty(_geeb._beeb._dgfb); _ebbb != nil {
			return _ebbb
		}
	}
	if _geeb._babc._dcg {
		if _efcf := _aaddb.WriteProperty(_geeb._beeb._ddba); _efcf != nil {
			return _efcf
		}
	}
	if _geeb._babc._bcecg {
		if _eafe := _aaddb.WriteProperty(_geeb._beeb._gcg); _eafe != nil {
			return _eafe
		}
	}
	if _geeb._babc._bceg {
		if _agad := _aaddb.WriteProperty(_geeb._beeb._dae); _agad != nil {
			return _agad
		}
	}
	if _geeb._babc._gafe {
		if _gdcd := _aaddb.WriteProperty(_geeb._beeb._bbgd); _gdcd != nil {
			return _gdcd
		}
	}
	return _aaddb.AlignLength(4)
}

func (_dbda *CommandButtonControl) writePropMask(_ddcd *_d.Writer) error {
	_cadf := uint32(0)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._cac)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._ddfd)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._ff)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._dbcc)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._caa)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._cacd)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._aef)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._ddb)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._acde)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._feg)
	_cadf = _d.PushLeftUI32(_cadf, _dbda._dg._dcab)
	_cadf >>= 21
	return _fe.Write(_ddcd, _fe.LittleEndian, _cadf)
}

// Control represents an ActiveX control wrapper.
type Control struct {
	TargetAttr string
	Ocx        *_ba.Ocx
	Choice     *ControlChoice
	_gg        string
	_ed        uint8
	_db        uint8
	_a         uint16
}

func (_cbde *SpinButtonControl) writeDataBlock(_ccgc *_d.Writer) error {
	if _cbde._fdbac._cede {
		if _afcd := _ccgc.WriteProperty(_cbde._deeda._fagff); _afcd != nil {
			return _afcd
		}
	}
	if _cbde._fdbac._aeec {
		if _gebc := _ccgc.WriteProperty(_cbde._deeda._abfba); _gebc != nil {
			return _gebc
		}
	}
	if _cbde._fdbac._fba {
		if _bfef := _ccgc.WriteProperty(_cbde._deeda._dbbe); _bfef != nil {
			return _bfef
		}
	}
	if _cbde._fdbac._edbfd {
		if _bffc := _ccgc.WriteProperty(_cbde._deeda._dbea); _bffc != nil {
			return _bffc
		}
	}
	if _cbde._fdbac._fce {
		if _efdd := _ccgc.WriteProperty(_cbde._deeda._fdfd); _efdd != nil {
			return _efdd
		}
	}
	if _cbde._fdbac._fgbc {
		if _adef := _ccgc.WriteProperty(_cbde._deeda._ecb); _adef != nil {
			return _adef
		}
	}
	if _cbde._fdbac._gaab {
		if _ccb := _ccgc.WriteProperty(_cbde._deeda._cegg); _ccb != nil {
			return _ccb
		}
	}
	if _cbde._fdbac._ebee {
		if _fdfe := _ccgc.WriteProperty(_cbde._deeda._gede); _fdfe != nil {
			return _fdfe
		}
	}
	if _cbde._fdbac._bgee {
		if _edaf := _ccgc.WriteProperty(_cbde._deeda._agaeg); _edaf != nil {
			return _edaf
		}
	}
	if _cbde._fdbac._efee {
		if _bafe := _ccgc.WriteProperty(_cbde._deeda._adfed); _bafe != nil {
			return _bafe
		}
	}
	if _cbde._fdbac._dged {
		if _fefa := _ccgc.WriteProperty(_cbde._deeda._fbcf); _fefa != nil {
			return _fefa
		}
	}
	if _cbde._fdbac._cbgd {
		if _eeag := _ccgc.WriteProperty(_cbde._deeda._bcac); _eeag != nil {
			return _eeag
		}
	}
	if _cbde._fdbac._fcca {
		if _ffdc := _ccgc.WriteProperty(_cbde._deeda._bcaf); _ffdc != nil {
			return _ffdc
		}
	}
	return _ccgc.AlignLength(4)
}

func (_ffc *ScrollBarControl) export(_fcacb *_d.Writer) error {
	if _bcad := _ffc.writePropMask(_fcacb); _bcad != nil {
		return _bcad
	}
	if _acfe := _ffc.writeDataBlock(_fcacb); _acfe != nil {
		return _acfe
	}
	if _ebcf := _ffc.writeExtraDataBlock(_fcacb); _ebcf != nil {
		return _ebcf
	}
	_ffc._fdd = uint16(_fcacb.Len() - 4)
	if _dgd := _ffc.writeStreamDataM(_fcacb); _dgd != nil {
		return _dgd
	}
	return _ffc.writeTheRest(_fcacb)
}

func (_dgdd *ScrollBarControl) readDataBlock(_cedbd *_d.Reader) error {
	_dgdd._beeb = &scrollBarDataBlock{}
	if _dgdd._babc._face {
		if _gddc := _cedbd.ReadProperty(&_dgdd._beeb._bdfb); _gddc != nil {
			return _gddc
		}
	}
	if _dgdd._babc._eeea {
		if _aadg := _cedbd.ReadProperty(&_dgdd._beeb._daag); _aadg != nil {
			return _aadg
		}
	}
	if _dgdd._babc._fcad {
		if _cecd := _cedbd.ReadProperty(&_dgdd._beeb._ceda); _cecd != nil {
			return _cecd
		}
	}
	if _dgdd._babc._faga {
		if _cgda := _cedbd.ReadProperty(&_dgdd._beeb._dgcg); _cgda != nil {
			return _cgda
		}
	}
	if _dgdd._babc._cee {
		if _cgdg := _cedbd.ReadProperty(&_dgdd._beeb._ecg); _cgdg != nil {
			return _cgdg
		}
	}
	if _dgdd._babc._edcef {
		if _fggb := _cedbd.ReadProperty(&_dgdd._beeb._agcg); _fggb != nil {
			return _fggb
		}
	}
	if _dgdd._babc._aefa {
		if _bba := _cedbd.ReadProperty(&_dgdd._beeb._cggd); _bba != nil {
			return _bba
		}
	}
	if _dgdd._babc._cff {
		if _dafe := _cedbd.ReadProperty(&_dgdd._beeb._fda); _dafe != nil {
			return _dafe
		}
	}
	if _dgdd._babc._fedd {
		if _fabc := _cedbd.ReadProperty(&_dgdd._beeb._abac); _fabc != nil {
			return _fabc
		}
	}
	if _dgdd._babc._ddac {
		if _ffcb := _cedbd.ReadProperty(&_dgdd._beeb._cfag); _ffcb != nil {
			return _ffcb
		}
	}
	if _dgdd._babc._abfb {
		if _ddfda := _cedbd.ReadProperty(&_dgdd._beeb._dgfb); _ddfda != nil {
			return _ddfda
		}
	}
	if _dgdd._babc._dcg {
		if _eef := _cedbd.ReadProperty(&_dgdd._beeb._ddba); _eef != nil {
			return _eef
		}
	}
	if _dgdd._babc._bcecg {
		if _ecge := _cedbd.ReadProperty(&_dgdd._beeb._gcg); _ecge != nil {
			return _ecge
		}
	}
	if _dgdd._babc._bceg {
		if _fbgc := _cedbd.ReadProperty(&_dgdd._beeb._dae); _fbgc != nil {
			return _fbgc
		}
	}
	if _dgdd._babc._gafe {
		if _bfab := _cedbd.ReadProperty(&_dgdd._beeb._bbgd); _bfab != nil {
			return _bfab
		}
	}
	return nil
}

func (_fed *controlBase) getPictureBytes() []byte { return _fed._efb._egef.getStdPictureBytes() }

// CheckBoxDataControl is a representation of a check box ActiveX form.
type CheckBoxDataControl struct{ *morphDataControlBoolValue }

// SetForeColor sets a button text color value from a system palette for a commandButton control.
func (_gbb *CommandButtonControl) SetForeColor(foreColor uint32) {
	_gbb._dg._cac = true
	_gbb._cbe._ggdd = foreColor
}

type scrollBarPropMask struct {
	_face  bool
	_eeea  bool
	_fcad  bool
	_faga  bool
	_afac  bool
	_cee   bool
	_edcef bool
	_aefa  bool
	_cff   bool
	_fedd  bool
	_ddac  bool
	_abfb  bool
	_dcg   bool
	_bcecg bool
	_bceg  bool
	_gafe  bool
}

// GetMin gets a button min value.
func (_gccea *ScrollBarControl) GetMin() int32 { return _gccea._beeb._ecg }

const _bega = 2

func (_cggg *morphDataControl) writePropMask(_ged *_d.Writer) error {
	_gbda := uint64(0)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._gcce)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._afdd)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._def)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._gged)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._geec)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._bfc)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._bfbd)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._bbc)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._aeef)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._fca)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._agbc)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._dbff)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._gfdg)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._bdgf)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._ffga)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._cggb)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._babf)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._eggg)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._beeeb)
	_gbda >>= 1
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._ffag)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._adbb)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._fcc)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._bcec)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._dbbc)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._addg)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._dda)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._cceg)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._bdca)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._effb)
	_gbda >>= 1
	_gbda = _d.PushLeftUI64(_gbda, true)
	_gbda = _d.PushLeftUI64(_gbda, _cggg._feb._edbg)
	_gbda >>= 31
	return _fe.Write(_ged, _fe.LittleEndian, _gbda)
}

// LabelControl is a representation of a label ActiveX form.
type LabelControl struct {
	controlBase
	_gcb  *labelPropMask
	_ebc  *labelDataBlock
	_abeb *labelExtraDataBlock
	_cbd  *streamData
}

// SetHeight sets height of the SpinButton in HIMETRIC (0.01mm)
func (_egge *SpinButtonControl) SetHeight(height uint32) {
	_egge._fdbac._fbdg = true
	_egge._aeaa._badg = height
}

func (_fee *morphDataControl) writeColumnInfoDataBlock(_ecee *_d.Writer) error { return nil }

// SetMin sets a button min value.
func (_cedd *ScrollBarControl) SetMin(min int32) { _cedd._babc._cee = true; _cedd._beeb._ecg = min }

func _ggd(_bdge *_d.Reader) (*CommandButtonControl, error) {
	_fggd := &CommandButtonControl{}
	if _ead := _fggd.readPropMask(_bdge); _ead != nil {
		return nil, _ead
	}
	if _aa := _fggd.readDataBlock(_bdge); _aa != nil {
		return nil, _aa
	}
	if _beg := _fggd.readExtraDataBlock(_bdge); _beg != nil {
		return nil, _beg
	}
	if _dac := _fggd.readStreamDataPM(_bdge, _fggd._dg._ddb, _fggd._dg._dcab); _dac != nil {
		return nil, _dac
	}
	if _fgf := _fggd.readTheRest(_bdge); _fgf != nil {
		return nil, _fgf
	}
	return _fggd, nil
}

func (_adga *ScrollBarControl) writePropMask(_cbg *_d.Writer) error {
	_cegdf := uint32(0)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._face)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._eeea)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._fcad)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._afac)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._faga)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._cee)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._edcef)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._aefa)
	_cegdf >>= 1
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._cff)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._fedd)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._ddac)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._abfb)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._dcg)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._bcecg)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._bceg)
	_cegdf = _d.PushLeftUI32(_cegdf, _adga._babc._gafe)
	_cegdf >>= 15
	return _fe.Write(_cbg, _fe.LittleEndian, _cegdf)
}

// SetHeight sets height of the CommandButton in HIMETRIC (0.01mm)
func (_ecc *CommandButtonControl) SetHeight(height uint32) {
	_ecc._dg._cacd = true
	_ecc._ag._dfeg = height
}

func (_afb *ScrollBarControl) readPropMask(_fcfa *_d.Reader) error {
	var _cbfc uint32
	if _ecfba := _fe.Read(_fcfa, _fe.LittleEndian, &_cbfc); _ecfba != nil {
		return _ecfba
	}
	_afb._babc = &scrollBarPropMask{}
	_afb._babc._face, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._eeea, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._fcad, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._afac, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._faga, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._cee, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._edcef, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._aefa, _cbfc = _d.PopRightUI32(_cbfc)
	_cbfc >>= 1
	_afb._babc._cff, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._fedd, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._ddac, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._abfb, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._dcg, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._bcecg, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._bceg, _cbfc = _d.PopRightUI32(_cbfc)
	_afb._babc._gafe, _cbfc = _d.PopRightUI32(_cbfc)
	return nil
}

func _cdfb(_bedg *_d.Reader) (*LabelControl, error) {
	_gda := &LabelControl{}
	if _dbaf := _gda.readPropMask(_bedg); _dbaf != nil {
		return nil, _dbaf
	}
	if _gcde := _gda.readDataBlock(_bedg); _gcde != nil {
		return nil, _gcde
	}
	if _efef := _gda.readExtraDataBlock(_bedg); _efef != nil {
		return nil, _efef
	}
	if _cdb := _gda.readStreamDataPM(_bedg, _gda._gcb._caacd, _gda._gcb._aad); _cdb != nil {
		return nil, _cdb
	}
	if _abf := _gda.readTheRest(_bedg); _abf != nil {
		return nil, _abf
	}
	return _gda, nil
}

// SetMin sets a button min value.
func (_gbdc *SpinButtonControl) SetMin(min int32) {
	_gbdc._fdbac._edbfd = true
	_gbdc._deeda._dbea = min
}

// GetValue gets a value from a control which value can be represented as a string.
func (_agb *morphDataControlStringValue) GetValue() string { return _agb.getValueString() }
