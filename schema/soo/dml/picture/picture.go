package picture

import (
	_f "encoding/xml"

	_e "github.com/arcpd/msword"
	_d "github.com/arcpd/msword/common/logger"
	_fa "github.com/arcpd/msword/schema/soo/dml"
)

type CT_PictureNonVisual struct {
	CNvPr    *_fa.CT_NonVisualDrawingProps
	CNvPicPr *_fa.CT_NonVisualPictureProperties
}

func (_cea *Pic) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_cea.CT_Picture = *NewCT_Picture()
_gg:
	for {
		_gad, _ffb := d.Token()
		if _ffb != nil {
			return _ffb
		}
		switch _ba := _gad.(type) {
		case _f.StartElement:
			switch _ba.Name {
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}, _f.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _aggd := d.DecodeElement(_cea.NvPicPr, &_ba); _aggd != nil {
					return _aggd
				}
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}, _f.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _fb := d.DecodeElement(_cea.BlipFill, &_ba); _fb != nil {
					return _fb
				}
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}, _f.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}:
				if _baa := d.DecodeElement(_cea.SpPr, &_ba); _baa != nil {
					return _baa
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006fn\u0020\u0050i\u0063\u0020\u0025\u0076", _ba.Name)
				if _abe := d.Skip(); _abe != nil {
					return _abe
				}
			}
		case _f.EndElement:
			break _gg
		case _f.CharData:
		}
	}
	return nil
}

// Validate validates the Pic and its children
func (_bd *Pic) Validate() error { return _bd.ValidateWithPath("\u0050\u0069\u0063") }

func (_ed *CT_PictureNonVisual) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	e.EncodeToken(start)
	_fae := _f.StartElement{Name: _f.Name{Local: "\u0070i\u0063\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_ed.CNvPr, _fae)
	_dab := _f.StartElement{Name: _f.Name{Local: "\u0070\u0069\u0063:\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_ed.CNvPicPr, _dab)
	e.EncodeToken(_f.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_agb *CT_Picture) ValidateWithPath(path string) error {
	if _ee := _agb.NvPicPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0050\u0069\u0063\u0050\u0072"); _ee != nil {
		return _ee
	}
	if _fc := _agb.BlipFill.ValidateWithPath(path + "\u002fB\u006c\u0069\u0070\u0046\u0069\u006cl"); _fc != nil {
		return _fc
	}
	if _dg := _agb.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _dg != nil {
		return _dg
	}
	return nil
}

// Validate validates the CT_PictureNonVisual and its children
func (_eeb *CT_PictureNonVisual) Validate() error {
	return _eeb.ValidateWithPath("\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}

type CT_Picture struct {
	NvPicPr  *CT_PictureNonVisual
	BlipFill *_fa.CT_BlipFillProperties
	SpPr     *_fa.CT_ShapeProperties
}

type Pic struct{ CT_Picture }

// ValidateWithPath validates the Pic and its children, prefixing error messages with path
func (_cgf *Pic) ValidateWithPath(path string) error {
	if _eb := _cgf.CT_Picture.ValidateWithPath(path); _eb != nil {
		return _eb
	}
	return nil
}

func NewCT_PictureNonVisual() *CT_PictureNonVisual {
	_agg := &CT_PictureNonVisual{}
	_agg.CNvPr = _fa.NewCT_NonVisualDrawingProps()
	_agg.CNvPicPr = _fa.NewCT_NonVisualPictureProperties()
	return _agg
}

func NewPic() *Pic { _dgf := &Pic{}; _dgf.CT_Picture = *NewCT_Picture(); return _dgf }

func (_ff *CT_PictureNonVisual) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_ff.CNvPr = _fa.NewCT_NonVisualDrawingProps()
	_ff.CNvPicPr = _fa.NewCT_NonVisualPictureProperties()
_ge:
	for {
		_ef, _dff := d.Token()
		if _dff != nil {
			return _dff
		}
		switch _agf := _ef.(type) {
		case _f.StartElement:
			switch _agf.Name {
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0063\u004e\u0076P\u0072"}, _f.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0063\u004e\u0076P\u0072"}:
				if _fga := d.DecodeElement(_ff.CNvPr, &_agf); _fga != nil {
					return _fga
				}
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}, _f.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}:
				if _bg := d.DecodeElement(_ff.CNvPicPr, &_agf); _bg != nil {
					return _bg
				}
			default:
				_d.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065No\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _agf.Name)
				if _cc := d.Skip(); _cc != nil {
					return _cc
				}
			}
		case _f.EndElement:
			break _ge
		case _f.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (_ga *CT_PictureNonVisual) ValidateWithPath(path string) error {
	if _ceb := _ga.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _ceb != nil {
		return _ceb
	}
	if _fee := _ga.CNvPicPr.ValidateWithPath(path + "\u002fC\u004e\u0076\u0050\u0069\u0063\u0050r"); _fee != nil {
		return _fee
	}
	return nil
}

// Validate validates the CT_Picture and its children
func (_fe *CT_Picture) Validate() error {
	return _fe.ValidateWithPath("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065")
}

func (_df *CT_Picture) UnmarshalXML(d *_f.Decoder, start _f.StartElement) error {
	_df.NvPicPr = NewCT_PictureNonVisual()
	_df.BlipFill = _fa.NewCT_BlipFillProperties()
	_df.SpPr = _fa.NewCT_ShapeProperties()
_b:
	for {
		_cg, _cgd := d.Token()
		if _cgd != nil {
			return _cgd
		}
		switch _cge := _cg.(type) {
		case _f.StartElement:
			switch _cge.Name {
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}, _f.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _ag := d.DecodeElement(_df.NvPicPr, &_cge); _ag != nil {
					return _ag
				}
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}, _f.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _fg := d.DecodeElement(_df.BlipFill, &_cge); _fg != nil {
					return _fg
				}
			case _f.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}, _f.Name{Space: "\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065", Local: "\u0073\u0070\u0050\u0072"}:
				if _ce := d.DecodeElement(_df.SpPr, &_cge); _ce != nil {
					return _ce
				}
			default:
				_d.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0069\u0063\u0074\u0075\u0072\u0065\u0020\u0025\u0076", _cge.Name)
				if _g := d.Skip(); _g != nil {
					return _g
				}
			}
		case _f.EndElement:
			break _b
		case _f.CharData:
		}
	}
	return nil
}

func NewCT_Picture() *CT_Picture {
	_dd := &CT_Picture{}
	_dd.NvPicPr = NewCT_PictureNonVisual()
	_dd.BlipFill = _fa.NewCT_BlipFillProperties()
	_dd.SpPr = _fa.NewCT_ShapeProperties()
	return _dd
}

func (_c *CT_Picture) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	e.EncodeToken(start)
	_da := _f.StartElement{Name: _f.Name{Local: "p\u0069\u0063\u003a\u006e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_c.NvPicPr, _da)
	_eg := _f.StartElement{Name: _f.Name{Local: "\u0070\u0069\u0063:\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}}
	e.EncodeElement(_c.BlipFill, _eg)
	_ab := _f.StartElement{Name: _f.Name{Local: "\u0070\u0069\u0063\u003a\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_c.SpPr, _ab)
	e.EncodeToken(_f.EndElement{Name: start.Name})
	return nil
}

func (_dfg *Pic) MarshalXML(e *_f.Encoder, start _f.StartElement) error {
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"})
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0070\u0069c"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"})
	start.Attr = append(start.Attr, _f.Attr{Name: _f.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0070i\u0063\u003a\u0070\u0069\u0063"
	return _dfg.CT_Picture.MarshalXML(e, start)
}

func init() {
	_e.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", "\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_PictureNonVisual)
	_e.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", "\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065", NewCT_Picture)
	_e.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065", "\u0070\u0069\u0063", NewPic)
}
