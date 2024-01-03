package spreadsheetDrawing

import (
	_ae "encoding/xml"
	_gg "fmt"
	_g "strconv"

	_ac "github.com/arcpd/msword"
	_f "github.com/arcpd/msword/common/logger"
	_c "github.com/arcpd/msword/schema/soo/dml"
)

func (_bag ST_EditAs) ValidateWithPath(path string) error {
	switch _bag {
	case 0, 1, 2, 3:
	default:
		return _gg.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bag))
	}
	return nil
}

type CT_Rel struct{ IdAttr string }

type CT_TwoCellAnchor struct {
	EditAsAttr ST_EditAs
	From       *CT_Marker
	To         *CT_Marker
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

// Validate validates the CT_Shape and its children
func (_dafa *CT_Shape) Validate() error {
	return _dafa.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065")
}

// Validate validates the CT_GraphicalObjectFrame and its children
func (_bfc *CT_GraphicalObjectFrame) Validate() error {
	return _bfc.ValidateWithPath("\u0043\u0054\u005fGr\u0061\u0070\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006a\u0065\u0063\u0074\u0046\u0072\u0061\u006d\u0065")
}

func NewFrom() *From { _ggb := &From{}; _ggb.CT_Marker = *NewCT_Marker(); return _ggb }

type EG_Anchor struct {
	TwoCellAnchor  *CT_TwoCellAnchor
	OneCellAnchor  *CT_OneCellAnchor
	AbsoluteAnchor *CT_AbsoluteAnchor
}

func (_ag *CT_AnchorClientData) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _ag.FLocksWithSheetAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0066L\u006fc\u006b\u0073\u0057\u0069\u0074\u0068\u0053\u0068\u0065\u0065\u0074"}, Value: _gg.Sprintf("\u0025\u0064", _dggcd(*_ag.FLocksWithSheetAttr))})
	}
	if _ag.FPrintsWithSheetAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0066\u0050r\u0069\u006e\u0074s\u0057\u0069\u0074\u0068\u0053\u0068\u0065\u0065\u0074"}, Value: _gg.Sprintf("\u0025\u0064", _dggcd(*_ag.FPrintsWithSheetAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

func (_dgbf *CT_Marker) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_dgbf.Col = 0
	_dgbf.Row = 0
_aeff:
	for {
		_cgcb, _fae := d.Token()
		if _fae != nil {
			return _fae
		}
		switch _gace := _cgcb.(type) {
		case _ae.StartElement:
			switch _gace.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}:
				if _fbe := d.DecodeElement(&_dgbf.Col, &_gace); _fbe != nil {
					return _fbe
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}:
				_ecd, _ecee := d.Token()
				if _ecee != nil {
					return _ecee
				}
				switch _edc := _ecd.(type) {
				case _ae.CharData:
					_agf := string(_edc)
					_gacea, _cegd := _c.ParseUnionST_Coordinate(_agf)
					if _cegd != nil {
						return nil
					}
					_dgbf.ColOff = _gacea
					d.Skip()
				case _ae.EndElement:
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}:
				if _afb := d.DecodeElement(&_dgbf.Row, &_gace); _afb != nil {
					return _afb
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}:
				_gce, _cccg := d.Token()
				if _cccg != nil {
					return _cccg
				}
				switch _bcbb := _gce.(type) {
				case _ae.CharData:
					_dcdc := string(_bcbb)
					_ade, _eeac := _c.ParseUnionST_Coordinate(_dcdc)
					if _eeac != nil {
						return nil
					}
					_dgbf.RowOff = _ade
					d.Skip()
				case _ae.EndElement:
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u004d\u0061\u0072k\u0065\u0072 \u0025\u0076", _gace.Name)
				if _aadeg := d.Skip(); _aadeg != nil {
					return _aadeg
				}
			}
		case _ae.EndElement:
			break _aeff
		case _ae.CharData:
		}
	}
	return nil
}

func (_gb *CT_ConnectorNonVisual) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	_fa := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_gb.CNvPr, _fa)
	_eb := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u004e\u0076\u0043\u0078n\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_gb.CNvCxnSpPr, _eb)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

type CT_Picture struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvPicPr        *CT_PictureNonVisual
	BlipFill       *_c.CT_BlipFillProperties
	SpPr           *_c.CT_ShapeProperties
	Style          *_c.CT_ShapeStyle
}

type CT_ConnectorNonVisual struct {
	CNvPr      *_c.CT_NonVisualDrawingProps
	CNvCxnSpPr *_c.CT_NonVisualConnectorProperties
}

func (_baa *EG_ObjectChoices) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
_cgef:
	for {
		_fcce, _gabc := d.Token()
		if _gabc != nil {
			return _gabc
		}
		switch _beff := _fcce.(type) {
		case _ae.StartElement:
			switch _beff.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_baa.Choice = NewEG_ObjectChoicesChoice()
				if _geg := d.DecodeElement(&_baa.Choice.Sp, &_beff); _geg != nil {
					return _geg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_baa.Choice = NewEG_ObjectChoicesChoice()
				if _cfaf := d.DecodeElement(&_baa.Choice.GrpSp, &_beff); _cfaf != nil {
					return _cfaf
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_baa.Choice = NewEG_ObjectChoicesChoice()
				if _fbde := d.DecodeElement(&_baa.Choice.GraphicFrame, &_beff); _fbde != nil {
					return _fbde
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_baa.Choice = NewEG_ObjectChoicesChoice()
				if _gcg := d.DecodeElement(&_baa.Choice.CxnSp, &_beff); _gcg != nil {
					return _gcg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_baa.Choice = NewEG_ObjectChoicesChoice()
				if _gfdd := d.DecodeElement(&_baa.Choice.Pic, &_beff); _gfdd != nil {
					return _gfdd
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_baa.Choice = NewEG_ObjectChoicesChoice()
				if _bcgd := d.DecodeElement(&_baa.Choice.ContentPart, &_beff); _bcgd != nil {
					return _bcgd
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u0047\u005f\u004f\u0062\u006a\u0065\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0020\u0025v", _beff.Name)
				if _ead := d.Skip(); _ead != nil {
					return _ead
				}
			}
		case _ae.EndElement:
			break _cgef
		case _ae.CharData:
		}
	}
	return nil
}

func (_fga *CT_GraphicalObjectFrameNonVisual) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_fga.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_fga.CNvGraphicFramePr = _c.NewCT_NonVisualGraphicFrameProperties()
_ccc:
	for {
		_ecc, _eagb := d.Token()
		if _eagb != nil {
			return _eagb
		}
		switch _cag := _ecc.(type) {
		case _ae.StartElement:
			switch _cag.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _dade := d.DecodeElement(_fga.CNvPr, &_cag); _dade != nil {
					return _dade
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"}:
				if _bce := d.DecodeElement(_fga.CNvGraphicFramePr, &_cag); _bce != nil {
					return _bce
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073u\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u0061p\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006a\u0065\u0063\u0074\u0046\u0072\u0061\u006de\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061l\u0020\u0025\u0076", _cag.Name)
				if _bgd := d.Skip(); _bgd != nil {
					return _bgd
				}
			}
		case _ae.EndElement:
			break _ccc
		case _ae.CharData:
		}
	}
	return nil
}

func _dggcd(_bbcg bool) uint8 {
	if _bbcg {
		return 1
	}
	return 0
}

type CT_GraphicalObjectFrameNonVisual struct {
	CNvPr             *_c.CT_NonVisualDrawingProps
	CNvGraphicFramePr *_c.CT_NonVisualGraphicFrameProperties
}

type CT_PictureNonVisual struct {
	CNvPr    *_c.CT_NonVisualDrawingProps
	CNvPicPr *_c.CT_NonVisualPictureProperties
}

func (_dad *CT_GraphicalObjectFrame) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_dad.NvGraphicFramePr = NewCT_GraphicalObjectFrameNonVisual()
	_dad.Xfrm = _c.NewCT_Transform2D()
	_dad.Graphic = _c.NewGraphic()
	for _, _beea := range start.Attr {
		if _beea.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_fgc, _bcb := _beea.Value, error(nil)
			if _bcb != nil {
				return _bcb
			}
			_dad.MacroAttr = &_fgc
			continue
		}
		if _beea.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_ded, _gcd := _g.ParseBool(_beea.Value)
			if _gcd != nil {
				return _gcd
			}
			_dad.FPublishedAttr = &_ded
			continue
		}
	}
_gcae:
	for {
		_gbc, _cbf := d.Token()
		if _cbf != nil {
			return _cbf
		}
		switch _fcd := _gbc.(type) {
		case _ae.StartElement:
			switch _fcd.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006e\u0076G\u0072\u0061\u0070h\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006e\u0076G\u0072\u0061\u0070h\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u0050\u0072"}:
				if _ddg := d.DecodeElement(_dad.NvGraphicFramePr, &_fcd); _ddg != nil {
					return _ddg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0078\u0066\u0072\u006d"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0078\u0066\u0072\u006d"}:
				if _bbb := d.DecodeElement(_dad.Xfrm, &_fcd); _bbb != nil {
					return _bbb
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0067r\u0061\u0070\u0068\u0069\u0063"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072g\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u006d\u0061\u0069\u006e", Local: "\u0067r\u0061\u0070\u0068\u0069\u0063"}:
				if _eag := d.DecodeElement(_dad.Graphic, &_fcd); _eag != nil {
					return _eag
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn\u0020\u0043\u0054\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0061\u006cO\u0062\u006a\u0065\u0063\u0074\u0046r\u0061\u006d\u0065 \u0025\u0076", _fcd.Name)
				if _cf := d.Skip(); _cf != nil {
					return _cf
				}
			}
		case _ae.EndElement:
			break _gcae
		case _ae.CharData:
		}
	}
	return nil
}

type CT_GroupShapeNonVisual struct {
	CNvPr      *_c.CT_NonVisualDrawingProps
	CNvGrpSpPr *_c.CT_NonVisualGroupDrawingShapeProps
}

// Validate validates the CT_GroupShape and its children
func (_eff *CT_GroupShape) Validate() error {
	return _eff.ValidateWithPath("\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065")
}

// Validate validates the CT_Drawing and its children
func (_gfc *CT_Drawing) Validate() error {
	return _gfc.ValidateWithPath("\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067")
}

func NewCT_Shape() *CT_Shape {
	_fcb := &CT_Shape{}
	_fcb.NvSpPr = NewCT_ShapeNonVisual()
	_fcb.SpPr = _c.NewCT_ShapeProperties()
	return _fcb
}

type To struct{ CT_Marker }

type CT_GroupShapeChoice struct {
	Sp           []*CT_Shape
	GrpSp        []*CT_GroupShape
	GraphicFrame []*CT_GraphicalObjectFrame
	CxnSp        []*CT_Connector
	Pic          []*CT_Picture
}

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (_dfg *CT_PictureNonVisual) ValidateWithPath(path string) error {
	if _bbg := _dfg.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _bbg != nil {
		return _bbg
	}
	if _dccg := _dfg.CNvPicPr.ValidateWithPath(path + "\u002fC\u004e\u0076\u0050\u0069\u0063\u0050r"); _dccg != nil {
		return _dccg
	}
	return nil
}

// Validate validates the EG_ObjectChoices and its children
func (_gdd *EG_ObjectChoices) Validate() error {
	return _gdd.ValidateWithPath("\u0045\u0047_\u004f\u0062\u006ae\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073")
}

func (_ffga *CT_Drawing) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
_fca:
	for {
		_ffa, _cafe := d.Token()
		if _cafe != nil {
			return _cafe
		}
		switch _bed := _ffa.(type) {
		case _ae.StartElement:
			switch _bed.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_dc := NewEG_Anchor()
				_dc.TwoCellAnchor = NewCT_TwoCellAnchor()
				if _geee := d.DecodeElement(_dc.TwoCellAnchor, &_bed); _geee != nil {
					return _geee
				}
				_ffga.EG_Anchor = append(_ffga.EG_Anchor, _dc)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_cef := NewEG_Anchor()
				_cef.OneCellAnchor = NewCT_OneCellAnchor()
				if _cdf := d.DecodeElement(_cef.OneCellAnchor, &_bed); _cdf != nil {
					return _cdf
				}
				_ffga.EG_Anchor = append(_ffga.EG_Anchor, _cef)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}:
				_fda := NewEG_Anchor()
				_fda.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if _bee := d.DecodeElement(_fda.AbsoluteAnchor, &_bed); _bee != nil {
					return _bee
				}
				_ffga.EG_Anchor = append(_ffga.EG_Anchor, _fda)
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fD\u0072\u0061\u0077\u0069\u006e\u0067\u0020\u0025\u0076", _bed.Name)
				if _eea := d.Skip(); _eea != nil {
					return _eea
				}
			}
		case _ae.EndElement:
			break _fca
		case _ae.CharData:
		}
	}
	return nil
}

func (_ddc *CT_PictureNonVisual) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	_gfcf := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_ddc.CNvPr, _gfcf)
	_aaff := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072:\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_ddc.CNvPicPr, _aaff)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

func (_cbe *CT_GraphicalObjectFrameNonVisual) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	_acb := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_cbe.CNvPr, _acb)
	_gd := _ae.StartElement{Name: _ae.Name{Local: "x\u0064\u0072\u003a\u0063Nv\u0047r\u0061\u0070\u0068\u0069\u0063F\u0072\u0061\u006d\u0065\u0050\u0072"}}
	e.EncodeElement(_cbe.CNvGraphicFramePr, _gd)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_ShapeNonVisual and its children, prefixing error messages with path
func (_efg *CT_ShapeNonVisual) ValidateWithPath(path string) error {
	if _geaag := _efg.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _geaag != nil {
		return _geaag
	}
	if _gdg := _efg.CNvSpPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0053\u0070\u0050\u0072"); _gdg != nil {
		return _gdg
	}
	return nil
}

type EG_ObjectChoices struct{ Choice *EG_ObjectChoicesChoice }

// ValidateWithPath validates the CT_Connector and its children, prefixing error messages with path
func (_gff *CT_Connector) ValidateWithPath(path string) error {
	if _cb := _gff.NvCxnSpPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"); _cb != nil {
		return _cb
	}
	if _ecg := _gff.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _ecg != nil {
		return _ecg
	}
	if _gff.Style != nil {
		if _bb := _gff.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _bb != nil {
			return _bb
		}
	}
	return nil
}

func NewCT_Marker() *CT_Marker { _aade := &CT_Marker{}; _aade.Col = 0; _aade.Row = 0; return _aade }

func (_afc *CT_GraphicalObjectFrame) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _afc.MacroAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _gg.Sprintf("\u0025\u0076", *_afc.MacroAttr)})
	}
	if _afc.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _gg.Sprintf("\u0025\u0064", _dggcd(*_afc.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_bff := _ae.StartElement{Name: _ae.Name{Local: "x\u0064r\u003a\u006e\u0076\u0047\u0072\u0061\u0070\u0068i\u0063\u0046\u0072\u0061me\u0050\u0072"}}
	e.EncodeElement(_afc.NvGraphicFramePr, _bff)
	_abcd := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0078\u0066\u0072\u006d"}}
	e.EncodeElement(_afc.Xfrm, _abcd)
	_bg := _ae.StartElement{Name: _ae.Name{Local: "\u0061:\u0067\u0072\u0061\u0070\u0068\u0069c"}}
	_bg.Attr = append(_bg.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	e.EncodeElement(_afc.Graphic, _bg)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Marker and its children, prefixing error messages with path
func (_abff *CT_Marker) ValidateWithPath(path string) error {
	if _abff.Col < 0 {
		return _gg.Errorf("\u0025\u0073\u002fm\u002e\u0043\u006f\u006c \u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u003e\u003d\u0020\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _abff.Col)
	}
	if _cage := _abff.ColOff.ValidateWithPath(path + "\u002fC\u006f\u006c\u004f\u0066\u0066"); _cage != nil {
		return _cage
	}
	if _abff.Row < 0 {
		return _gg.Errorf("\u0025\u0073\u002fm\u002e\u0052\u006f\u0077 \u006d\u0075\u0073\u0074\u0020\u0062\u0065 \u003e\u003d\u0020\u0030\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029", path, _abff.Row)
	}
	if _gbe := _abff.RowOff.ValidateWithPath(path + "\u002fR\u006f\u0077\u004f\u0066\u0066"); _gbe != nil {
		return _gbe
	}
	return nil
}

func (_cdd *CT_Connector) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_cdd.NvCxnSpPr = NewCT_ConnectorNonVisual()
	_cdd.SpPr = _c.NewCT_ShapeProperties()
	for _, _eced := range start.Attr {
		if _eced.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_ffg, _efe := _eced.Value, error(nil)
			if _efe != nil {
				return _efe
			}
			_cdd.MacroAttr = &_ffg
			continue
		}
		if _eced.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_eca, _aeef := _g.ParseBool(_eced.Value)
			if _aeef != nil {
				return _aeef
			}
			_cdd.FPublishedAttr = &_eca
			continue
		}
	}
_fdb:
	for {
		_db, _gee := d.Token()
		if _gee != nil {
			return _gee
		}
		switch _ad := _db.(type) {
		case _ae.StartElement:
			switch _ad.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006ev\u0043\u0078\u006e\u0053\u0070\u0050r"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0043\u0078\u006e\u0053\u0070\u0050r"}:
				if _fc := d.DecodeElement(_cdd.NvCxnSpPr, &_ad); _fc != nil {
					return _fc
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _bc := d.DecodeElement(_cdd.SpPr, &_ad); _bc != nil {
					return _bc
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_cdd.Style = _c.NewCT_ShapeStyle()
				if _caf := d.DecodeElement(_cdd.Style, &_ad); _caf != nil {
					return _caf
				}
			default:
				_f.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_C\u006f\u006en\u0065\u0063\u0074\u006f\u0072\u0020\u0025\u0076", _ad.Name)
				if _dgd := d.Skip(); _dgd != nil {
					return _dgd
				}
			}
		case _ae.EndElement:
			break _fdb
		case _ae.CharData:
		}
	}
	return nil
}

func NewCT_GroupShape() *CT_GroupShape {
	_bffe := &CT_GroupShape{}
	_bffe.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	_bffe.GrpSpPr = _c.NewCT_GroupShapeProperties()
	return _bffe
}

type CT_ShapeNonVisual struct {
	CNvPr   *_c.CT_NonVisualDrawingProps
	CNvSpPr *_c.CT_NonVisualDrawingShapeProps
}

// ValidateWithPath validates the From and its children, prefixing error messages with path
func (_dgf *From) ValidateWithPath(path string) error {
	if _dfad := _dgf.CT_Marker.ValidateWithPath(path); _dfad != nil {
		return _dfad
	}
	return nil
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_bde *CT_Picture) ValidateWithPath(path string) error {
	if _abae := _bde.NvPicPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0050\u0069\u0063\u0050\u0072"); _abae != nil {
		return _abae
	}
	if _fgf := _bde.BlipFill.ValidateWithPath(path + "\u002fB\u006c\u0069\u0070\u0046\u0069\u006cl"); _fgf != nil {
		return _fgf
	}
	if _acfb := _bde.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _acfb != nil {
		return _acfb
	}
	if _bde.Style != nil {
		if _egfa := _bde.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _egfa != nil {
			return _egfa
		}
	}
	return nil
}

func (_ebg *CT_Picture) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _ebg.MacroAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _gg.Sprintf("\u0025\u0076", *_ebg.MacroAttr)})
	}
	if _ebg.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _gg.Sprintf("\u0025\u0064", _dggcd(*_ebg.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_fce := _ae.StartElement{Name: _ae.Name{Local: "x\u0064\u0072\u003a\u006e\u0076\u0050\u0069\u0063\u0050\u0072"}}
	e.EncodeElement(_ebg.NvPicPr, _fce)
	_edd := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072:\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}}
	e.EncodeElement(_ebg.BlipFill, _edd)
	_cedd := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_ebg.SpPr, _cedd)
	if _ebg.Style != nil {
		_acca := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0073\u0074\u0079\u006ce"}}
		e.EncodeElement(_ebg.Style, _acca)
	}
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

func NewCT_OneCellAnchor() *CT_OneCellAnchor {
	_dgef := &CT_OneCellAnchor{}
	_dgef.From = NewCT_Marker()
	_dgef.Ext = _c.NewCT_PositiveSize2D()
	_dgef.ClientData = NewCT_AnchorClientData()
	return _dgef
}

func (_dbab *CT_Rel) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0072\u003a\u0069\u0064"}, Value: _gg.Sprintf("\u0025\u0076", _dbab.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

func NewCT_PictureNonVisual() *CT_PictureNonVisual {
	_ceee := &CT_PictureNonVisual{}
	_ceee.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_ceee.CNvPicPr = _c.NewCT_NonVisualPictureProperties()
	return _ceee
}

func NewCT_GraphicalObjectFrameNonVisual() *CT_GraphicalObjectFrameNonVisual {
	_ba := &CT_GraphicalObjectFrameNonVisual{}
	_ba.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_ba.CNvGraphicFramePr = _c.NewCT_NonVisualGraphicFrameProperties()
	return _ba
}

// ValidateWithPath validates the EG_Anchor and its children, prefixing error messages with path
func (_edg *EG_Anchor) ValidateWithPath(path string) error {
	if _edg.TwoCellAnchor != nil {
		if _aefg := _edg.TwoCellAnchor.ValidateWithPath(path + "\u002f\u0054\u0077\u006f\u0043\u0065\u006c\u006c\u0041n\u0063\u0068\u006f\u0072"); _aefg != nil {
			return _aefg
		}
	}
	if _edg.OneCellAnchor != nil {
		if _gffe := _edg.OneCellAnchor.ValidateWithPath(path + "\u002f\u004f\u006e\u0065\u0043\u0065\u006c\u006c\u0041n\u0063\u0068\u006f\u0072"); _gffe != nil {
			return _gffe
		}
	}
	if _edg.AbsoluteAnchor != nil {
		if _fbee := _edg.AbsoluteAnchor.ValidateWithPath(path + "\u002fA\u0062s\u006f\u006c\u0075\u0074\u0065\u0041\u006e\u0063\u0068\u006f\u0072"); _fbee != nil {
			return _fbee
		}
	}
	return nil
}

func NewCT_GroupShapeNonVisual() *CT_GroupShapeNonVisual {
	_cgc := &CT_GroupShapeNonVisual{}
	_cgc.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_cgc.CNvGrpSpPr = _c.NewCT_NonVisualGroupDrawingShapeProps()
	return _cgc
}

func (_fabg *CT_PictureNonVisual) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_fabg.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_fabg.CNvPicPr = _c.NewCT_NonVisualPictureProperties()
_cbb:
	for {
		_ged, _abaeg := d.Token()
		if _abaeg != nil {
			return _abaeg
		}
		switch _egcb := _ged.(type) {
		case _ae.StartElement:
			switch _egcb.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _bgb := d.DecodeElement(_fabg.CNvPr, &_egcb); _bgb != nil {
					return _bgb
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}:
				if _efda := d.DecodeElement(_fabg.CNvPicPr, &_egcb); _efda != nil {
					return _efda
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065No\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _egcb.Name)
				if _cbac := d.Skip(); _cbac != nil {
					return _cbac
				}
			}
		case _ae.EndElement:
			break _cbb
		case _ae.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_GroupShapeChoice and its children, prefixing error messages with path
func (_cgfa *CT_GroupShapeChoice) ValidateWithPath(path string) error {
	for _eec, _cbd := range _cgfa.Sp {
		if _ddf := _cbd.ValidateWithPath(_gg.Sprintf("\u0025s\u002f\u0053\u0070\u005b\u0025\u0064]", path, _eec)); _ddf != nil {
			return _ddf
		}
	}
	for _bbc, _cec := range _cgfa.GrpSp {
		if _cfa := _cec.ValidateWithPath(_gg.Sprintf("\u0025\u0073\u002fG\u0072\u0070\u0053\u0070\u005b\u0025\u0064\u005d", path, _bbc)); _cfa != nil {
			return _cfa
		}
	}
	for _edbb, _afg := range _cgfa.GraphicFrame {
		if _adf := _afg.ValidateWithPath(_gg.Sprintf("\u0025\u0073\u002f\u0047ra\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065\u005b\u0025\u0064\u005d", path, _edbb)); _adf != nil {
			return _adf
		}
	}
	for _dfa, _gdf := range _cgfa.CxnSp {
		if _fdad := _gdf.ValidateWithPath(_gg.Sprintf("\u0025\u0073\u002fC\u0078\u006e\u0053\u0070\u005b\u0025\u0064\u005d", path, _dfa)); _fdad != nil {
			return _fdad
		}
	}
	for _aad, _acea := range _cgfa.Pic {
		if _fdab := _acea.ValidateWithPath(_gg.Sprintf("\u0025\u0073\u002f\u0050\u0069\u0063\u005b\u0025\u0064\u005d", path, _aad)); _fdab != nil {
			return _fdab
		}
	}
	return nil
}

// Validate validates the EG_ObjectChoicesChoice and its children
func (_dbd *EG_ObjectChoicesChoice) Validate() error {
	return _dbd.ValidateWithPath("\u0045\u0047\u005f\u004fbj\u0065\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0043\u0068\u006f\u0069c\u0065")
}

// Validate validates the From and its children
func (_aec *From) Validate() error { return _aec.ValidateWithPath("\u0046\u0072\u006f\u006d") }

func NewEG_ObjectChoices() *EG_ObjectChoices { _fcg := &EG_ObjectChoices{}; return _fcg }

type CT_OneCellAnchor struct {
	From       *CT_Marker
	Ext        *_c.CT_PositiveSize2D
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

// ValidateWithPath validates the CT_ConnectorNonVisual and its children, prefixing error messages with path
func (_acfd *CT_ConnectorNonVisual) ValidateWithPath(path string) error {
	if _efbfb := _acfd.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _efbfb != nil {
		return _efbfb
	}
	if _ccb := _acfd.CNvCxnSpPr.ValidateWithPath(path + "/\u0043\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"); _ccb != nil {
		return _ccb
	}
	return nil
}

func (_gde *CT_ShapeNonVisual) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_gde.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_gde.CNvSpPr = _c.NewCT_NonVisualDrawingShapeProps()
_abb:
	for {
		_bedc, _ggcg := d.Token()
		if _ggcg != nil {
			return _ggcg
		}
		switch _bdfb := _bedc.(type) {
		case _ae.StartElement:
			switch _bdfb.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _geaa := d.DecodeElement(_gde.CNvPr, &_bdfb); _geaa != nil {
					return _geaa
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063N\u0076\u0053\u0070\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063N\u0076\u0053\u0070\u0050\u0072"}:
				if _cgd := d.DecodeElement(_gde.CNvSpPr, &_bdfb); _cgd != nil {
					return _cgd
				}
			default:
				_f.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _bdfb.Name)
				if _gad := d.Skip(); _gad != nil {
					return _gad
				}
			}
		case _ae.EndElement:
			break _abb
		case _ae.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the To and its children, prefixing error messages with path
func (_beee *To) ValidateWithPath(path string) error {
	if _adad := _beee.CT_Marker.ValidateWithPath(path); _adad != nil {
		return _adad
	}
	return nil
}

func NewCT_AbsoluteAnchor() *CT_AbsoluteAnchor {
	_d := &CT_AbsoluteAnchor{}
	_d.Pos = _c.NewCT_Point2D()
	_d.Ext = _c.NewCT_PositiveSize2D()
	_d.ClientData = NewCT_AnchorClientData()
	return _d
}

func NewCT_Picture() *CT_Picture {
	_dagg := &CT_Picture{}
	_dagg.NvPicPr = NewCT_PictureNonVisual()
	_dagg.BlipFill = _c.NewCT_BlipFillProperties()
	_dagg.SpPr = _c.NewCT_ShapeProperties()
	return _dagg
}

func (_eceda *CT_Shape) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _eceda.MacroAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _gg.Sprintf("\u0025\u0076", *_eceda.MacroAttr)})
	}
	if _eceda.TextlinkAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0074\u0065\u0078\u0074\u006c\u0069\u006e\u006b"}, Value: _gg.Sprintf("\u0025\u0076", *_eceda.TextlinkAttr)})
	}
	if _eceda.FLocksTextAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0066\u004c\u006f\u0063\u006b\u0073\u0054\u0065\u0078\u0074"}, Value: _gg.Sprintf("\u0025\u0064", _dggcd(*_eceda.FLocksTextAttr))})
	}
	if _eceda.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _gg.Sprintf("\u0025\u0064", _dggcd(*_eceda.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_ddfg := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u006e\u0076\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_eceda.NvSpPr, _ddfg)
	_gfe := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_eceda.SpPr, _gfe)
	if _eceda.Style != nil {
		_agb := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0073\u0074\u0079\u006ce"}}
		e.EncodeElement(_eceda.Style, _agb)
	}
	if _eceda.TxBody != nil {
		_bcbf := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0074\u0078\u0042\u006f\u0064\u0079"}}
		e.EncodeElement(_eceda.TxBody, _bcbf)
	}
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_GraphicalObjectFrameNonVisual and its children
func (_dba *CT_GraphicalObjectFrameNonVisual) Validate() error {
	return _dba.ValidateWithPath("\u0043\u0054\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006ae\u0063t\u0046\u0072\u0061\u006d\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}

// Validate validates the CT_ShapeNonVisual and its children
func (_effc *CT_ShapeNonVisual) Validate() error {
	return _effc.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c")
}

func (_ecdb *CT_Rel) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	for _, _ggg := range start.Attr {
		if _ggg.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _ggg.Name.Local == "\u0069\u0064" || _ggg.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _ggg.Name.Local == "\u0069\u0064" {
			_bcg, _cbc := _ggg.Value, error(nil)
			if _cbc != nil {
				return _cbc
			}
			_ecdb.IdAttr = _bcg
			continue
		}
	}
	for {
		_feec, _dfea := d.Token()
		if _dfea != nil {
			return _gg.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073", _dfea)
		}
		if _gae, _afcb := _feec.(_ae.EndElement); _afcb && _gae.Name == start.Name {
			break
		}
	}
	return nil
}

func (_gacf *From) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u0064r"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0078\u0064\u0072\u003a\u0066\u0072\u006f\u006d"
	return _gacf.CT_Marker.MarshalXML(e, start)
}

func (_dce *From) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_dce.CT_Marker = *NewCT_Marker()
_efa:
	for {
		_cfge, _adag := d.Token()
		if _adag != nil {
			return _adag
		}
		switch _efcc := _cfge.(type) {
		case _ae.StartElement:
			switch _efcc.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}:
				if _eaf := d.DecodeElement(&_dce.Col, &_efcc); _eaf != nil {
					return _eaf
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}:
				_cbea, _bdfd := d.Token()
				if _bdfd != nil {
					return _bdfd
				}
				switch _fcab := _cbea.(type) {
				case _ae.CharData:
					_cbae := string(_fcab)
					_fecfa, _ecba := _c.ParseUnionST_Coordinate(_cbae)
					if _ecba != nil {
						return nil
					}
					_dce.ColOff = _fecfa
					d.Skip()
				case _ae.EndElement:
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}:
				if _efag := d.DecodeElement(&_dce.Row, &_efcc); _efag != nil {
					return _efag
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}:
				_gfea, _ccf := d.Token()
				if _ccf != nil {
					return _ccf
				}
				switch _ffgc := _gfea.(type) {
				case _ae.CharData:
					_aggc := string(_ffgc)
					_bfdf, _fbfb := _c.ParseUnionST_Coordinate(_aggc)
					if _fbfb != nil {
						return nil
					}
					_dce.RowOff = _bfdf
					d.Skip()
				case _ae.EndElement:
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0046\u0072o\u006d\u0020\u0025\u0076", _efcc.Name)
				if _eaeb := d.Skip(); _eaeb != nil {
					return _eaeb
				}
			}
		case _ae.EndElement:
			break _efa
		case _ae.CharData:
		}
	}
	return nil
}

// Validate validates the WsDr and its children
func (_dbg *WsDr) Validate() error { return _dbg.ValidateWithPath("\u0057\u0073\u0044\u0072") }

func (_cgb *CT_GroupShapeNonVisual) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_cgb.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_cgb.CNvGrpSpPr = _c.NewCT_NonVisualGroupDrawingShapeProps()
_bede:
	for {
		_fad, _cae := d.Token()
		if _cae != nil {
			return _cae
		}
		switch _feef := _fad.(type) {
		case _ae.StartElement:
			switch _feef.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _cfg := d.DecodeElement(_cgb.CNvPr, &_feef); _cfg != nil {
					return _cfg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}:
				if _faa := d.DecodeElement(_cgb.CNvGrpSpPr, &_feef); _faa != nil {
					return _faa
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070p\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0047\u0072\u006f\u0075p\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _feef.Name)
				if _fac := d.Skip(); _fac != nil {
					return _fac
				}
			}
		case _ae.EndElement:
			break _bede
		case _ae.CharData:
		}
	}
	return nil
}

func NewCT_ConnectorNonVisual() *CT_ConnectorNonVisual {
	_de := &CT_ConnectorNonVisual{}
	_de.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_de.CNvCxnSpPr = _c.NewCT_NonVisualConnectorProperties()
	return _de
}

func (_ceg *CT_Drawing) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	if _ceg.EG_Anchor != nil {
		for _, _cada := range _ceg.EG_Anchor {
			_cada.MarshalXML(e, _ae.StartElement{})
		}
	}
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

func (_gegc *WsDr) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u0064r"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0078\u0064\u0072\u003a\u0077\u0073\u0044\u0072"
	return _gegc.CT_Drawing.MarshalXML(e, start)
}

// ValidateWithPath validates the CT_GroupShape and its children, prefixing error messages with path
func (_bffa *CT_GroupShape) ValidateWithPath(path string) error {
	if _efef := _bffa.NvGrpSpPr.ValidateWithPath(path + "\u002f\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _efef != nil {
		return _efef
	}
	if _gac := _bffa.GrpSpPr.ValidateWithPath(path + "\u002f\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _gac != nil {
		return _gac
	}
	for _abf, _fegf := range _bffa.Choice {
		if _bgab := _fegf.ValidateWithPath(_gg.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d", path, _abf)); _bgab != nil {
			return _bgab
		}
	}
	return nil
}

func (_bdb *ST_EditAs) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_gfed, _fbc := d.Token()
	if _fbc != nil {
		return _fbc
	}
	if _eggf, _afee := _gfed.(_ae.EndElement); _afee && _eggf.Name == start.Name {
		*_bdb = 1
		return nil
	}
	if _cga, _cccc := _gfed.(_ae.CharData); !_cccc {
		return _gg.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gfed)
	} else {
		switch string(_cga) {
		case "":
			*_bdb = 0
		case "\u0074w\u006f\u0043\u0065\u006c\u006c":
			*_bdb = 1
		case "\u006fn\u0065\u0043\u0065\u006c\u006c":
			*_bdb = 2
		case "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065":
			*_bdb = 3
		}
	}
	_gfed, _fbc = d.Token()
	if _fbc != nil {
		return _fbc
	}
	if _gdca, _caca := _gfed.(_ae.EndElement); _caca && _gdca.Name == start.Name {
		return nil
	}
	return _gg.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gfed)
}

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_bbgg *CT_Rel) ValidateWithPath(path string) error { return nil }

func (_cge *CT_OneCellAnchor) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_cge.From = NewCT_Marker()
	_cge.Ext = _c.NewCT_PositiveSize2D()
	_cge.ClientData = NewCT_AnchorClientData()
_dab:
	for {
		_bbe, _bdaf := d.Token()
		if _bdaf != nil {
			return _bdaf
		}
		switch _eeb := _bbe.(type) {
		case _ae.StartElement:
			switch _eeb.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}:
				if _cdg := d.DecodeElement(_cge.From, &_eeb); _cdg != nil {
					return _cdg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}:
				if _egca := d.DecodeElement(_cge.Ext, &_eeb); _egca != nil {
					return _egca
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_cge.Choice = NewEG_ObjectChoicesChoice()
				if _dga := d.DecodeElement(&_cge.Choice.Sp, &_eeb); _dga != nil {
					return _dga
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_cge.Choice = NewEG_ObjectChoicesChoice()
				if _fab := d.DecodeElement(&_cge.Choice.GrpSp, &_eeb); _fab != nil {
					return _fab
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_cge.Choice = NewEG_ObjectChoicesChoice()
				if _def := d.DecodeElement(&_cge.Choice.GraphicFrame, &_eeb); _def != nil {
					return _def
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_cge.Choice = NewEG_ObjectChoicesChoice()
				if _bfe := d.DecodeElement(&_cge.Choice.CxnSp, &_eeb); _bfe != nil {
					return _bfe
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_cge.Choice = NewEG_ObjectChoicesChoice()
				if _egdc := d.DecodeElement(&_cge.Choice.Pic, &_eeb); _egdc != nil {
					return _egdc
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_cge.Choice = NewEG_ObjectChoicesChoice()
				if _gced := d.DecodeElement(&_cge.Choice.ContentPart, &_eeb); _gced != nil {
					return _gced
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}:
				if _cdce := d.DecodeElement(_cge.ClientData, &_eeb); _cdce != nil {
					return _cdce
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025v", _eeb.Name)
				if _ffag := d.Skip(); _ffag != nil {
					return _ffag
				}
			}
		case _ae.EndElement:
			break _dab
		case _ae.CharData:
		}
	}
	return nil
}

type CT_GraphicalObjectFrame struct {
	MacroAttr        *string
	FPublishedAttr   *bool
	NvGraphicFramePr *CT_GraphicalObjectFrameNonVisual
	Xfrm             *_c.CT_Transform2D
	Graphic          *_c.Graphic
}

// ValidateWithPath validates the CT_OneCellAnchor and its children, prefixing error messages with path
func (_acba *CT_OneCellAnchor) ValidateWithPath(path string) error {
	if _cbfg := _acba.From.ValidateWithPath(path + "\u002f\u0046\u0072o\u006d"); _cbfg != nil {
		return _cbfg
	}
	if _fdbed := _acba.Ext.ValidateWithPath(path + "\u002f\u0045\u0078\u0074"); _fdbed != nil {
		return _fdbed
	}
	if _acba.Choice != nil {
		if _daec := _acba.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _daec != nil {
			return _daec
		}
	}
	if _caa := _acba.ClientData.ValidateWithPath(path + "/\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"); _caa != nil {
		return _caa
	}
	return nil
}

func (_edbf *CT_OneCellAnchor) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	_fgac := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0066\u0072\u006f\u006d"}}
	e.EncodeElement(_edbf.From, _fgac)
	_bea := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0065\u0078\u0074"}}
	e.EncodeElement(_edbf.Ext, _bea)
	if _edbf.Choice != nil {
		_edbf.Choice.MarshalXML(e, _ae.StartElement{})
	}
	_ffab := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u006c\u0069\u0065\u006et\u0044\u0061\u0074\u0061"}}
	e.EncodeElement(_edbf.ClientData, _ffab)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_AbsoluteAnchor and its children, prefixing error messages with path
func (_eg *CT_AbsoluteAnchor) ValidateWithPath(path string) error {
	if _b := _eg.Pos.ValidateWithPath(path + "\u002f\u0050\u006f\u0073"); _b != nil {
		return _b
	}
	if _efb := _eg.Ext.ValidateWithPath(path + "\u002f\u0045\u0078\u0074"); _efb != nil {
		return _efb
	}
	if _eg.Choice != nil {
		if _ca := _eg.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _ca != nil {
			return _ca
		}
	}
	if _ee := _eg.ClientData.ValidateWithPath(path + "/\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"); _ee != nil {
		return _ee
	}
	return nil
}

func (_cab *CT_GroupShapeChoice) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _cab.Sp != nil {
		_bcf := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070"}}
		for _, _gbb := range _cab.Sp {
			e.EncodeElement(_gbb, _bcf)
		}
	}
	if _cab.GrpSp != nil {
		_fbd := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0067\u0072\u0070\u0053p"}}
		for _, _gcb := range _cab.GrpSp {
			e.EncodeElement(_gcb, _fbd)
		}
	}
	if _cab.GraphicFrame != nil {
		_abgb := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064r\u003a\u0067\u0072a\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}}
		for _, _bdf := range _cab.GraphicFrame {
			e.EncodeElement(_bdf, _abgb)
		}
	}
	if _cab.CxnSp != nil {
		_ecb := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0063\u0078\u006e\u0053p"}}
		for _, _cdc := range _cab.CxnSp {
			e.EncodeElement(_cdc, _ecb)
		}
	}
	if _cab.Pic != nil {
		_dcd := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0070\u0069\u0063"}}
		for _, _acg := range _cab.Pic {
			e.EncodeElement(_acg, _dcd)
		}
	}
	return nil
}

func (_egd *CT_AnchorClientData) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	for _, _ff := range start.Attr {
		if _ff.Name.Local == "\u0066L\u006fc\u006b\u0073\u0057\u0069\u0074\u0068\u0053\u0068\u0065\u0065\u0074" {
			_fd, _be := _g.ParseBool(_ff.Value)
			if _be != nil {
				return _be
			}
			_egd.FLocksWithSheetAttr = &_fd
			continue
		}
		if _ff.Name.Local == "\u0066\u0050r\u0069\u006e\u0074s\u0057\u0069\u0074\u0068\u0053\u0068\u0065\u0065\u0074" {
			_fe, _gf := _g.ParseBool(_ff.Value)
			if _gf != nil {
				return _gf
			}
			_egd.FPrintsWithSheetAttr = &_fe
			continue
		}
	}
	for {
		_acf, _aa := d.Token()
		if _aa != nil {
			return _gg.Errorf("\u0070\u0061\u0072s\u0069\u006e\u0067\u0020C\u0054\u005f\u0041\u006e\u0063\u0068\u006fr\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061\u003a\u0020\u0025\u0073", _aa)
		}
		if _efd, _dgc := _acf.(_ae.EndElement); _dgc && _efd.Name == start.Name {
			break
		}
	}
	return nil
}

type CT_Drawing struct{ EG_Anchor []*EG_Anchor }

type CT_AnchorClientData struct {
	FLocksWithSheetAttr  *bool
	FPrintsWithSheetAttr *bool
}

func (_adfa *CT_Marker) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	_aff := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0063\u006f\u006c"}}
	e.EncodeElement(_adfa.Col, _aff)
	_egc := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u006f\u006c\u004f\u0066\u0066"}}
	e.EncodeElement(_adfa.ColOff, _egc)
	_agg := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0072\u006f\u0077"}}
	e.EncodeElement(_adfa.Row, _agg)
	_aef := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0072\u006f\u0077\u004f\u0066\u0066"}}
	e.EncodeElement(_adfa.RowOff, _aef)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

func NewCT_Drawing() *CT_Drawing { _dac := &CT_Drawing{}; return _dac }

func NewCT_Rel() *CT_Rel { _bffd := &CT_Rel{}; return _bffd }

type CT_Marker struct {
	Col    int32
	ColOff _c.ST_Coordinate
	Row    int32
	RowOff _c.ST_Coordinate
}

type CT_Connector struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvCxnSpPr      *CT_ConnectorNonVisual
	SpPr           *_c.CT_ShapeProperties
	Style          *_c.CT_ShapeStyle
}

// ValidateWithPath validates the CT_Drawing and its children, prefixing error messages with path
func (_feea *CT_Drawing) ValidateWithPath(path string) error {
	for _accb, _gge := range _feea.EG_Anchor {
		if _abc := _gge.ValidateWithPath(_gg.Sprintf("\u0025\u0073/\u0045\u0047\u005fA\u006e\u0063\u0068\u006f\u0072\u005b\u0025\u0064\u005d", path, _accb)); _abc != nil {
			return _abc
		}
	}
	return nil
}

// Validate validates the CT_ConnectorNonVisual and its children
func (_ffe *CT_ConnectorNonVisual) Validate() error {
	return _ffe.ValidateWithPath("C\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u004e\u006fn\u0056\u0069\u0073\u0075\u0061\u006c")
}

type CT_Shape struct {
	MacroAttr      *string
	TextlinkAttr   *string
	FLocksTextAttr *bool
	FPublishedAttr *bool
	NvSpPr         *CT_ShapeNonVisual
	SpPr           *_c.CT_ShapeProperties
	Style          *_c.CT_ShapeStyle
	TxBody         *_c.CT_TextBody
}

func NewEG_Anchor() *EG_Anchor { _agca := &EG_Anchor{}; return _agca }

func (_dgdc *To) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u0064r"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067"})
	start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0078\u0064\u0072\u003a\u0074\u006f"
	return _dgdc.CT_Marker.MarshalXML(e, start)
}

func (_bcbd *EG_ObjectChoicesChoice) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _bcbd.Sp != nil {
		_fgb := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070"}}
		e.EncodeElement(_bcbd.Sp, _fgb)
	}
	if _bcbd.GrpSp != nil {
		_gcgd := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0067\u0072\u0070\u0053p"}}
		e.EncodeElement(_bcbd.GrpSp, _gcgd)
	}
	if _bcbd.GraphicFrame != nil {
		_cdcf := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064r\u003a\u0067\u0072a\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}}
		e.EncodeElement(_bcbd.GraphicFrame, _cdcf)
	}
	if _bcbd.CxnSp != nil {
		_facd := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0063\u0078\u006e\u0053p"}}
		e.EncodeElement(_bcbd.CxnSp, _facd)
	}
	if _bcbd.Pic != nil {
		_cdcc := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0070\u0069\u0063"}}
		e.EncodeElement(_bcbd.Pic, _cdcc)
	}
	if _bcbd.ContentPart != nil {
		_gaa := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072:\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}}
		e.EncodeElement(_bcbd.ContentPart, _gaa)
	}
	return nil
}

func (_abd ST_EditAs) String() string {
	switch _abd {
	case 0:
		return ""
	case 1:
		return "\u0074w\u006f\u0043\u0065\u006c\u006c"
	case 2:
		return "\u006fn\u0065\u0043\u0065\u006c\u006c"
	case 3:
		return "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065"
	}
	return ""
}

// Validate validates the CT_GroupShapeChoice and its children
func (_bdd *CT_GroupShapeChoice) Validate() error {
	return _bdd.ValidateWithPath("\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u0043\u0068\u006f\u0069\u0063\u0065")
}

// Validate validates the CT_PictureNonVisual and its children
func (_afga *CT_PictureNonVisual) Validate() error {
	return _afga.ValidateWithPath("\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c")
}

// ValidateWithPath validates the CT_Shape and its children, prefixing error messages with path
func (_bad *CT_Shape) ValidateWithPath(path string) error {
	if _cgba := _bad.NvSpPr.ValidateWithPath(path + "\u002fN\u0076\u0053\u0070\u0050\u0072"); _cgba != nil {
		return _cgba
	}
	if _cfab := _bad.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _cfab != nil {
		return _cfab
	}
	if _bad.Style != nil {
		if _bgg := _bad.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _bgg != nil {
			return _bgg
		}
	}
	if _bad.TxBody != nil {
		if _deb := _bad.TxBody.ValidateWithPath(path + "\u002fT\u0078\u0042\u006f\u0064\u0079"); _deb != nil {
			return _deb
		}
	}
	return nil
}

// ValidateWithPath validates the EG_ObjectChoicesChoice and its children, prefixing error messages with path
func (_edec *EG_ObjectChoicesChoice) ValidateWithPath(path string) error {
	if _edec.Sp != nil {
		if _gcca := _edec.Sp.ValidateWithPath(path + "\u002f\u0053\u0070"); _gcca != nil {
			return _gcca
		}
	}
	if _edec.GrpSp != nil {
		if _cdfe := _edec.GrpSp.ValidateWithPath(path + "\u002f\u0047\u0072\u0070\u0053\u0070"); _cdfe != nil {
			return _cdfe
		}
	}
	if _edec.GraphicFrame != nil {
		if _eaga := _edec.GraphicFrame.ValidateWithPath(path + "\u002f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"); _eaga != nil {
			return _eaga
		}
	}
	if _edec.CxnSp != nil {
		if _fdc := _edec.CxnSp.ValidateWithPath(path + "\u002f\u0043\u0078\u006e\u0053\u0070"); _fdc != nil {
			return _fdc
		}
	}
	if _edec.Pic != nil {
		if _abe := _edec.Pic.ValidateWithPath(path + "\u002f\u0050\u0069\u0063"); _abe != nil {
			return _abe
		}
	}
	if _edec.ContentPart != nil {
		if _adce := _edec.ContentPart.ValidateWithPath(path + "\u002f\u0043\u006fn\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"); _adce != nil {
			return _adce
		}
	}
	return nil
}

func NewCT_TwoCellAnchor() *CT_TwoCellAnchor {
	_fddg := &CT_TwoCellAnchor{}
	_fddg.From = NewCT_Marker()
	_fddg.To = NewCT_Marker()
	_fddg.ClientData = NewCT_AnchorClientData()
	return _fddg
}

type EG_ObjectChoicesChoice struct {
	Sp           *CT_Shape
	GrpSp        *CT_GroupShape
	GraphicFrame *CT_GraphicalObjectFrame
	CxnSp        *CT_Connector
	Pic          *CT_Picture
	ContentPart  *CT_Rel
}

func (_edfe *CT_ShapeNonVisual) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	_feeb := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_edfe.CNvPr, _feeb)
	_effa := _ae.StartElement{Name: _ae.Name{Local: "x\u0064\u0072\u003a\u0063\u004e\u0076\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_edfe.CNvSpPr, _effa)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

type CT_AbsoluteAnchor struct {
	Pos        *_c.CT_Point2D
	Ext        *_c.CT_PositiveSize2D
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

// ValidateWithPath validates the CT_AnchorClientData and its children, prefixing error messages with path
func (_cee *CT_AnchorClientData) ValidateWithPath(path string) error { return nil }

// Validate validates the CT_AbsoluteAnchor and its children
func (_ef *CT_AbsoluteAnchor) Validate() error {
	return _ef.ValidateWithPath("\u0043\u0054\u005f\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072")
}

func (_fcbf *CT_Shape) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_fcbf.NvSpPr = NewCT_ShapeNonVisual()
	_fcbf.SpPr = _c.NewCT_ShapeProperties()
	for _, _ebb := range start.Attr {
		if _ebb.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_bdaa, _dgga := _ebb.Value, error(nil)
			if _dgga != nil {
				return _dgga
			}
			_fcbf.MacroAttr = &_bdaa
			continue
		}
		if _ebb.Name.Local == "\u0074\u0065\u0078\u0074\u006c\u0069\u006e\u006b" {
			_ggd, _ecea := _ebb.Value, error(nil)
			if _ecea != nil {
				return _ecea
			}
			_fcbf.TextlinkAttr = &_ggd
			continue
		}
		if _ebb.Name.Local == "\u0066\u004c\u006f\u0063\u006b\u0073\u0054\u0065\u0078\u0074" {
			_beb, _ecad := _g.ParseBool(_ebb.Value)
			if _ecad != nil {
				return _ecad
			}
			_fcbf.FLocksTextAttr = &_beb
			continue
		}
		if _ebb.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_acfe, _gaf := _g.ParseBool(_ebb.Value)
			if _gaf != nil {
				return _gaf
			}
			_fcbf.FPublishedAttr = &_acfe
			continue
		}
	}
_bcfa:
	for {
		_acfdg, _dgag := d.Token()
		if _dgag != nil {
			return _dgag
		}
		switch _dfb := _acfdg.(type) {
		case _ae.StartElement:
			switch _dfb.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006e\u0076\u0053\u0070\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006e\u0076\u0053\u0070\u0050\u0072"}:
				if _adee := d.DecodeElement(_fcbf.NvSpPr, &_dfb); _adee != nil {
					return _adee
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _cgga := d.DecodeElement(_fcbf.SpPr, &_dfb); _cgga != nil {
					return _cgga
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_fcbf.Style = _c.NewCT_ShapeStyle()
				if _edf := d.DecodeElement(_fcbf.Style, &_dfb); _edf != nil {
					return _edf
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u0078\u0042\u006f\u0064\u0079"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0078\u0042\u006f\u0064\u0079"}:
				_fcbf.TxBody = _c.NewCT_TextBody()
				if _adc := d.DecodeElement(_fcbf.TxBody, &_dfb); _adc != nil {
					return _adc
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u0020\u0025\u0076", _dfb.Name)
				if _eaec := d.Skip(); _eaec != nil {
					return _eaec
				}
			}
		case _ae.EndElement:
			break _bcfa
		case _ae.CharData:
		}
	}
	return nil
}

func (_egce *EG_ObjectChoicesChoice) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
_efdd:
	for {
		_beed, _eedg := d.Token()
		if _eedg != nil {
			return _eedg
		}
		switch _gbfc := _beed.(type) {
		case _ae.StartElement:
			switch _gbfc.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_egce.Sp = NewCT_Shape()
				if _cccga := d.DecodeElement(_egce.Sp, &_gbfc); _cccga != nil {
					return _cccga
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_egce.GrpSp = NewCT_GroupShape()
				if _cgcc := d.DecodeElement(_egce.GrpSp, &_gbfc); _cgcc != nil {
					return _cgcc
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_egce.GraphicFrame = NewCT_GraphicalObjectFrame()
				if _gfgg := d.DecodeElement(_egce.GraphicFrame, &_gbfc); _gfgg != nil {
					return _gfgg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_egce.CxnSp = NewCT_Connector()
				if _fgg := d.DecodeElement(_egce.CxnSp, &_gbfc); _fgg != nil {
					return _fgg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_egce.Pic = NewCT_Picture()
				if _cggd := d.DecodeElement(_egce.Pic, &_gbfc); _cggd != nil {
					return _cggd
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_egce.ContentPart = NewCT_Rel()
				if _agfa := d.DecodeElement(_egce.ContentPart, &_gbfc); _agfa != nil {
					return _agfa
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070p\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045G\u005f\u004f\u0062\u006a\u0065c\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _gbfc.Name)
				if _egbb := d.Skip(); _egbb != nil {
					return _egbb
				}
			}
		case _ae.EndElement:
			break _efdd
		case _ae.CharData:
		}
	}
	return nil
}

func (_dge *CT_GroupShape) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_dge.NvGrpSpPr = NewCT_GroupShapeNonVisual()
	_dge.GrpSpPr = _c.NewCT_GroupShapeProperties()
_aab:
	for {
		_efca, _ced := d.Token()
		if _ced != nil {
			return _ced
		}
		switch _dbb := _efca.(type) {
		case _ae.StartElement:
			switch _dbb.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006ev\u0047\u0072\u0070\u0053\u0070\u0050r"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0047\u0072\u0070\u0053\u0070\u0050r"}:
				if _aaf := d.DecodeElement(_dge.NvGrpSpPr, &_dbb); _aaf != nil {
					return _aaf
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067r\u0070\u0053\u0070\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067r\u0070\u0053\u0070\u0050\u0072"}:
				if _gab := d.DecodeElement(_dge.GrpSpPr, &_dbb); _gab != nil {
					return _gab
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_fge := NewCT_GroupShapeChoice()
				if _bga := d.DecodeElement(&_fge.Sp, &_dbb); _bga != nil {
					return _bga
				}
				_dge.Choice = append(_dge.Choice, _fge)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_dcgb := NewCT_GroupShapeChoice()
				if _dca := d.DecodeElement(&_dcgb.GrpSp, &_dbb); _dca != nil {
					return _dca
				}
				_dge.Choice = append(_dge.Choice, _dcgb)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_acbg := NewCT_GroupShapeChoice()
				if _gfd := d.DecodeElement(&_acbg.GraphicFrame, &_dbb); _gfd != nil {
					return _gfd
				}
				_dge.Choice = append(_dge.Choice, _acbg)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_cac := NewCT_GroupShapeChoice()
				if _cgg := d.DecodeElement(&_cac.CxnSp, &_dbb); _cgg != nil {
					return _cgg
				}
				_dge.Choice = append(_dge.Choice, _cac)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_edb := NewCT_GroupShapeChoice()
				if _dag := d.DecodeElement(&_edb.Pic, &_dbb); _dag != nil {
					return _dag
				}
				_dge.Choice = append(_dge.Choice, _edb)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0047r\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065 \u0025\u0076", _dbb.Name)
				if _dcc := d.Skip(); _dcc != nil {
					return _dcc
				}
			}
		case _ae.EndElement:
			break _aab
		case _ae.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Marker and its children
func (_fdbe *CT_Marker) Validate() error {
	return _fdbe.ValidateWithPath("\u0043T\u005f\u004d\u0061\u0072\u006b\u0065r")
}

const (
	ST_EditAsUnset    ST_EditAs = 0
	ST_EditAsTwoCell  ST_EditAs = 1
	ST_EditAsOneCell  ST_EditAs = 2
	ST_EditAsAbsolute ST_EditAs = 3
)

func (_ge *CT_AbsoluteAnchor) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_ge.Pos = _c.NewCT_Point2D()
	_ge.Ext = _c.NewCT_PositiveSize2D()
	_ge.ClientData = NewCT_AnchorClientData()
_ggc:
	for {
		_abg, _ec := d.Token()
		if _ec != nil {
			return _ec
		}
		switch _ace := _abg.(type) {
		case _ae.StartElement:
			switch _ace.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u006f\u0073"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u006f\u0073"}:
				if _gc := d.DecodeElement(_ge.Pos, &_ace); _gc != nil {
					return _gc
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0065\u0078\u0074"}:
				if _cg := d.DecodeElement(_ge.Ext, &_ace); _cg != nil {
					return _cg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_ge.Choice = NewEG_ObjectChoicesChoice()
				if _ea := d.DecodeElement(&_ge.Choice.Sp, &_ace); _ea != nil {
					return _ea
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_ge.Choice = NewEG_ObjectChoicesChoice()
				if _da := d.DecodeElement(&_ge.Choice.GrpSp, &_ace); _da != nil {
					return _da
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_ge.Choice = NewEG_ObjectChoicesChoice()
				if _fg := d.DecodeElement(&_ge.Choice.GraphicFrame, &_ace); _fg != nil {
					return _fg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_ge.Choice = NewEG_ObjectChoicesChoice()
				if _cc := d.DecodeElement(&_ge.Choice.CxnSp, &_ace); _cc != nil {
					return _cc
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_ge.Choice = NewEG_ObjectChoicesChoice()
				if _ga := d.DecodeElement(&_ge.Choice.Pic, &_ace); _ga != nil {
					return _ga
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_ge.Choice = NewEG_ObjectChoicesChoice()
				if _dd := d.DecodeElement(&_ge.Choice.ContentPart, &_ace); _dd != nil {
					return _dd
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}:
				if _af := d.DecodeElement(_ge.ClientData, &_ace); _af != nil {
					return _af
				}
			default:
				_f.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025\u0076", _ace.Name)
				if _gca := d.Skip(); _gca != nil {
					return _gca
				}
			}
		case _ae.EndElement:
			break _ggc
		case _ae.CharData:
		}
	}
	return nil
}

func NewCT_AnchorClientData() *CT_AnchorClientData { _dg := &CT_AnchorClientData{}; return _dg }

func NewCT_ShapeNonVisual() *CT_ShapeNonVisual {
	_gceg := &CT_ShapeNonVisual{}
	_gceg.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_gceg.CNvSpPr = _c.NewCT_NonVisualDrawingShapeProps()
	return _gceg
}

func (_cdda *CT_GroupShapeNonVisual) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	_dggg := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0063\u004e\u0076\u0050r"}}
	e.EncodeElement(_cdda.CNvPr, _dggg)
	_eae := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u004e\u0076\u0047\u0072p\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_cdda.CNvGrpSpPr, _eae)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

func NewCT_GroupShapeChoice() *CT_GroupShapeChoice { _effe := &CT_GroupShapeChoice{}; return _effe }

func NewEG_ObjectChoicesChoice() *EG_ObjectChoicesChoice {
	_bdcf := &EG_ObjectChoicesChoice{}
	return _bdcf
}

// Validate validates the CT_Connector and its children
func (_bf *CT_Connector) Validate() error {
	return _bf.ValidateWithPath("\u0043\u0054\u005fC\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072")
}

func (_dacc *WsDr) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_dacc.CT_Drawing = *NewCT_Drawing()
_cbaf:
	for {
		_fdf, _fef := d.Token()
		if _fef != nil {
			return _fef
		}
		switch _gcad := _fdf.(type) {
		case _ae.StartElement:
			switch _gcad.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_afeg := NewEG_Anchor()
				_afeg.TwoCellAnchor = NewCT_TwoCellAnchor()
				if _bca := d.DecodeElement(_afeg.TwoCellAnchor, &_gcad); _bca != nil {
					return _bca
				}
				_dacc.EG_Anchor = append(_dacc.EG_Anchor, _afeg)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_eeg := NewEG_Anchor()
				_eeg.OneCellAnchor = NewCT_OneCellAnchor()
				if _dgfe := d.DecodeElement(_eeg.OneCellAnchor, &_gcad); _dgfe != nil {
					return _dgfe
				}
				_dacc.EG_Anchor = append(_dacc.EG_Anchor, _eeg)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}:
				_eddc := NewEG_Anchor()
				_eddc.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if _acee := d.DecodeElement(_eddc.AbsoluteAnchor, &_gcad); _acee != nil {
					return _acee
				}
				_dacc.EG_Anchor = append(_dacc.EG_Anchor, _eddc)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0057\u0073D\u0072\u0020\u0025\u0076", _gcad.Name)
				if _ddga := d.Skip(); _ddga != nil {
					return _ddga
				}
			}
		case _ae.EndElement:
			break _cbaf
		case _ae.CharData:
		}
	}
	return nil
}

type From struct{ CT_Marker }

// ValidateWithPath validates the CT_TwoCellAnchor and its children, prefixing error messages with path
func (_gbf *CT_TwoCellAnchor) ValidateWithPath(path string) error {
	if _efefa := _gbf.EditAsAttr.ValidateWithPath(path + "/\u0045\u0064\u0069\u0074\u0041\u0073\u0041\u0074\u0074\u0072"); _efefa != nil {
		return _efefa
	}
	if _dea := _gbf.From.ValidateWithPath(path + "\u002f\u0046\u0072o\u006d"); _dea != nil {
		return _dea
	}
	if _aafc := _gbf.To.ValidateWithPath(path + "\u002f\u0054\u006f"); _aafc != nil {
		return _aafc
	}
	if _gbf.Choice != nil {
		if _eedb := _gbf.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _eedb != nil {
			return _eedb
		}
	}
	if _agfc := _gbf.ClientData.ValidateWithPath(path + "/\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"); _agfc != nil {
		return _agfc
	}
	return nil
}

func NewTo() *To { _aggf := &To{}; _aggf.CT_Marker = *NewCT_Marker(); return _aggf }

func (_e *CT_AbsoluteAnchor) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	_aee := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0070\u006f\u0073"}}
	e.EncodeElement(_e.Pos, _aee)
	_ab := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0065\u0078\u0074"}}
	e.EncodeElement(_e.Ext, _ab)
	if _e.Choice != nil {
		_e.Choice.MarshalXML(e, _ae.StartElement{})
	}
	_ce := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u006c\u0069\u0065\u006et\u0044\u0061\u0074\u0061"}}
	e.EncodeElement(_e.ClientData, _ce)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

func NewWsDr() *WsDr { _gabce := &WsDr{}; _gabce.CT_Drawing = *NewCT_Drawing(); return _gabce }

func (_geb *CT_GroupShape) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	e.EncodeToken(start)
	_dae := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u006e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_geb.NvGrpSpPr, _dae)
	_fea := _ae.StartElement{Name: _ae.Name{Local: "x\u0064\u0072\u003a\u0067\u0072\u0070\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_geb.GrpSpPr, _fea)
	if _geb.Choice != nil {
		for _, _fgd := range _geb.Choice {
			_fgd.MarshalXML(e, _ae.StartElement{})
		}
	}
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

// Validate validates the EG_Anchor and its children
func (_gcce *EG_Anchor) Validate() error {
	return _gcce.ValidateWithPath("\u0045G\u005f\u0041\u006e\u0063\u0068\u006fr")
}

// ValidateWithPath validates the EG_ObjectChoices and its children, prefixing error messages with path
func (_cbg *EG_ObjectChoices) ValidateWithPath(path string) error {
	if _cbg.Choice != nil {
		if _ede := _cbg.Choice.ValidateWithPath(path + "\u002fC\u0068\u006f\u0069\u0063\u0065"); _ede != nil {
			return _ede
		}
	}
	return nil
}

func (_gdfd *CT_TwoCellAnchor) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _gdfd.EditAsAttr != ST_EditAsUnset {
		_dfbe, _cda := _gdfd.EditAsAttr.MarshalXMLAttr(_ae.Name{Local: "\u0065\u0064\u0069\u0074\u0041\u0073"})
		if _cda != nil {
			return _cda
		}
		start.Attr = append(start.Attr, _dfbe)
	}
	e.EncodeToken(start)
	_bfcf := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0066\u0072\u006f\u006d"}}
	e.EncodeElement(_gdfd.From, _bfcf)
	_acef := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0074\u006f"}}
	e.EncodeElement(_gdfd.To, _acef)
	if _gdfd.Choice != nil {
		_gdfd.Choice.MarshalXML(e, _ae.StartElement{})
	}
	_dfga := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0063\u006c\u0069\u0065\u006et\u0044\u0061\u0074\u0061"}}
	e.EncodeElement(_gdfd.ClientData, _dfga)
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_GroupShapeNonVisual and its children
func (_egf *CT_GroupShapeNonVisual) Validate() error {
	return _egf.ValidateWithPath("\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075a\u006c")
}

func (_fee *CT_Connector) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _fee.MacroAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u006d\u0061\u0063r\u006f"}, Value: _gg.Sprintf("\u0025\u0076", *_fee.MacroAttr)})
	}
	if _fee.FPublishedAttr != nil {
		start.Attr = append(start.Attr, _ae.Attr{Name: _ae.Name{Local: "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064"}, Value: _gg.Sprintf("\u0025\u0064", _dggcd(*_fee.FPublishedAttr))})
	}
	e.EncodeToken(start)
	_ece := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u006e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}}
	e.EncodeElement(_fee.NvCxnSpPr, _ece)
	_efc := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0073\u0070\u0050\u0072"}}
	e.EncodeElement(_fee.SpPr, _efc)
	if _fee.Style != nil {
		_efbf := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072\u003a\u0073\u0074\u0079\u006ce"}}
		e.EncodeElement(_fee.Style, _efbf)
	}
	e.EncodeToken(_ae.EndElement{Name: start.Name})
	return nil
}

type ST_EditAs byte

func (_gddg *ST_EditAs) UnmarshalXMLAttr(attr _ae.Attr) error {
	switch attr.Value {
	case "":
		*_gddg = 0
	case "\u0074w\u006f\u0043\u0065\u006c\u006c":
		*_gddg = 1
	case "\u006fn\u0065\u0043\u0065\u006c\u006c":
		*_gddg = 2
	case "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065":
		*_gddg = 3
	}
	return nil
}

type CT_GroupShape struct {
	NvGrpSpPr *CT_GroupShapeNonVisual
	GrpSpPr   *_c.CT_GroupShapeProperties
	Choice    []*CT_GroupShapeChoice
}

func (_edce *EG_Anchor) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
_egg:
	for {
		_gcbg, _cfgc := d.Token()
		if _cfgc != nil {
			return _cfgc
		}
		switch _dgeb := _gcbg.(type) {
		case _ae.StartElement:
			switch _dgeb.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_edce.TwoCellAnchor = NewCT_TwoCellAnchor()
				if _dfed := d.DecodeElement(_edce.TwoCellAnchor, &_dgeb); _dfed != nil {
					return _dfed
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072"}:
				_edce.OneCellAnchor = NewCT_OneCellAnchor()
				if _cddb := d.DecodeElement(_edce.OneCellAnchor, &_dgeb); _cddb != nil {
					return _cddb
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072"}:
				_edce.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if _bae := d.DecodeElement(_edce.AbsoluteAnchor, &_dgeb); _bae != nil {
					return _bae
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0045\u0047\u005f\u0041\u006e\u0063h\u006f\u0072 \u0025\u0076", _dgeb.Name)
				if _adaa := d.Skip(); _adaa != nil {
					return _adaa
				}
			}
		case _ae.EndElement:
			break _egg
		case _ae.CharData:
		}
	}
	return nil
}

func NewCT_Connector() *CT_Connector {
	_cd := &CT_Connector{}
	_cd.NvCxnSpPr = NewCT_ConnectorNonVisual()
	_cd.SpPr = _c.NewCT_ShapeProperties()
	return _cd
}

// ValidateWithPath validates the CT_GraphicalObjectFrame and its children, prefixing error messages with path
func (_afca *CT_GraphicalObjectFrame) ValidateWithPath(path string) error {
	if _ecgc := _afca.NvGraphicFramePr.ValidateWithPath(path + "\u002f\u004e\u0076\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0046\u0072a\u006d\u0065\u0050\u0072"); _ecgc != nil {
		return _ecgc
	}
	if _dcg := _afca.Xfrm.ValidateWithPath(path + "\u002f\u0058\u0066r\u006d"); _dcg != nil {
		return _dcg
	}
	if _fbf := _afca.Graphic.ValidateWithPath(path + "\u002f\u0047\u0072\u0061\u0070\u0068\u0069\u0063"); _fbf != nil {
		return _fbf
	}
	return nil
}

func NewCT_GraphicalObjectFrame() *CT_GraphicalObjectFrame {
	_fb := &CT_GraphicalObjectFrame{}
	_fb.NvGraphicFramePr = NewCT_GraphicalObjectFrameNonVisual()
	_fb.Xfrm = _c.NewCT_Transform2D()
	_fb.Graphic = _c.NewGraphic()
	return _fb
}

// Validate validates the CT_TwoCellAnchor and its children
func (_aadb *CT_TwoCellAnchor) Validate() error {
	return _aadb.ValidateWithPath("\u0043\u0054_\u0054\u0077\u006fC\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072")
}

// ValidateWithPath validates the CT_GroupShapeNonVisual and its children, prefixing error messages with path
func (_fdd *CT_GroupShapeNonVisual) ValidateWithPath(path string) error {
	if _ege := _fdd.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _ege != nil {
		return _ege
	}
	if _cgce := _fdd.CNvGrpSpPr.ValidateWithPath(path + "/\u0043\u004e\u0076\u0047\u0072\u0070\u0053\u0070\u0050\u0072"); _cgce != nil {
		return _cgce
	}
	return nil
}

// Validate validates the To and its children
func (_cbdd *To) Validate() error { return _cbdd.ValidateWithPath("\u0054\u006f") }

// Validate validates the CT_Rel and its children
func (_gga *CT_Rel) Validate() error {
	return _gga.ValidateWithPath("\u0043\u0054\u005f\u0052\u0065\u006c")
}

type WsDr struct{ CT_Drawing }

func (_eac *CT_ConnectorNonVisual) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_eac.CNvPr = _c.NewCT_NonVisualDrawingProps()
	_eac.CNvCxnSpPr = _c.NewCT_NonVisualConnectorProperties()
_gba:
	for {
		_feg, _bd := d.Token()
		if _bd != nil {
			return _bd
		}
		switch _acc := _feg.(type) {
		case _ae.StartElement:
			switch _acc.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076P\u0072"}:
				if _gcc := d.DecodeElement(_eac.CNvPr, &_acc); _gcc != nil {
					return _gcc
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u004e\u0076\u0043\u0078\u006e\u0053\u0070\u0050\u0072"}:
				if _fcc := d.DecodeElement(_eac.CNvCxnSpPr, &_acc); _fcc != nil {
					return _fcc
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075n\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u006f\u006e C\u0054\u005f\u0043\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076", _acc.Name)
				if _bdc := d.Skip(); _bdc != nil {
					return _bdc
				}
			}
		case _ae.EndElement:
			break _gba
		case _ae.CharData:
		}
	}
	return nil
}

func (_eaef ST_EditAs) MarshalXMLAttr(name _ae.Name) (_ae.Attr, error) {
	_affc := _ae.Attr{}
	_affc.Name = name
	switch _eaef {
	case ST_EditAsUnset:
		_affc.Value = ""
	case ST_EditAsTwoCell:
		_affc.Value = "\u0074w\u006f\u0043\u0065\u006c\u006c"
	case ST_EditAsOneCell:
		_affc.Value = "\u006fn\u0065\u0043\u0065\u006c\u006c"
	case ST_EditAsAbsolute:
		_affc.Value = "\u0061\u0062\u0073\u006f\u006c\u0075\u0074\u0065"
	}
	return _affc, nil
}

func (_bgabe ST_EditAs) Validate() error { return _bgabe.ValidateWithPath("") }

// Validate validates the CT_OneCellAnchor and its children
func (_gdc *CT_OneCellAnchor) Validate() error {
	return _gdc.ValidateWithPath("\u0043\u0054_\u004f\u006e\u0065C\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072")
}

// ValidateWithPath validates the WsDr and its children, prefixing error messages with path
func (_daa *WsDr) ValidateWithPath(path string) error {
	if _gggd := _daa.CT_Drawing.ValidateWithPath(path); _gggd != nil {
		return _gggd
	}
	return nil
}

// Validate validates the CT_Picture and its children
func (_cgfc *CT_Picture) Validate() error {
	return _cgfc.ValidateWithPath("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065")
}

func (_cadg *CT_Picture) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_cadg.NvPicPr = NewCT_PictureNonVisual()
	_cadg.BlipFill = _c.NewCT_BlipFillProperties()
	_cadg.SpPr = _c.NewCT_ShapeProperties()
	for _, _gea := range start.Attr {
		if _gea.Name.Local == "\u006d\u0061\u0063r\u006f" {
			_gcaee, _adea := _gea.Value, error(nil)
			if _adea != nil {
				return _adea
			}
			_cadg.MacroAttr = &_gcaee
			continue
		}
		if _gea.Name.Local == "\u0066\u0050\u0075\u0062\u006c\u0069\u0073\u0068\u0065\u0064" {
			_faae, _feac := _g.ParseBool(_gea.Value)
			if _feac != nil {
				return _feac
			}
			_cadg.FPublishedAttr = &_faae
			continue
		}
	}
_cdcg:
	for {
		_deg, _ceeg := d.Token()
		if _ceeg != nil {
			return _ceeg
		}
		switch _afd := _deg.(type) {
		case _ae.StartElement:
			switch _afd.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u006ev\u0050\u0069\u0063\u0050\u0072"}:
				if _dfe := d.DecodeElement(_cadg.NvPicPr, &_afd); _dfe != nil {
					return _dfe
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:
				if _agff := d.DecodeElement(_cadg.BlipFill, &_afd); _agff != nil {
					return _agff
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070\u0050\u0072"}:
				if _cba := d.DecodeElement(_cadg.SpPr, &_afd); _cba != nil {
					return _cba
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0074\u0079l\u0065"}:
				_cadg.Style = _c.NewCT_ShapeStyle()
				if _gabb := d.DecodeElement(_cadg.Style, &_afd); _gabb != nil {
					return _gabb
				}
			default:
				_f.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0069\u0063\u0074\u0075\u0072\u0065\u0020\u0025\u0076", _afd.Name)
				if _baf := d.Skip(); _baf != nil {
					return _baf
				}
			}
		case _ae.EndElement:
			break _cdcg
		case _ae.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AnchorClientData and its children
func (_dgb *CT_AnchorClientData) Validate() error {
	return _dgb.ValidateWithPath("\u0043\u0054\u005f\u0041nc\u0068\u006f\u0072\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061")
}

func (_cafef *EG_ObjectChoices) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _cafef.Choice != nil {
		_cafef.Choice.MarshalXML(e, _ae.StartElement{})
	}
	return nil
}

// ValidateWithPath validates the CT_GraphicalObjectFrameNonVisual and its children, prefixing error messages with path
func (_daf *CT_GraphicalObjectFrameNonVisual) ValidateWithPath(path string) error {
	if _ecac := _daf.CNvPr.ValidateWithPath(path + "\u002f\u0043\u004e\u0076\u0050\u0072"); _ecac != nil {
		return _ecac
	}
	if _gffb := _daf.CNvGraphicFramePr.ValidateWithPath(path + "\u002fC\u004ev\u0047\u0072\u0061\u0070\u0068i\u0063\u0046r\u0061\u006d\u0065\u0050\u0072"); _gffb != nil {
		return _gffb
	}
	return nil
}

func (_dec *EG_Anchor) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	if _dec.TwoCellAnchor != nil {
		_faba := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u0074\u0077\u006f\u0043\u0065\u006c\u006c\u0041n\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_dec.TwoCellAnchor, _faba)
	}
	if _dec.OneCellAnchor != nil {
		_fecf := _ae.StartElement{Name: _ae.Name{Local: "\u0078\u0064\u0072\u003a\u006f\u006e\u0065\u0043\u0065\u006c\u006c\u0041n\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_dec.OneCellAnchor, _fecf)
	}
	if _dec.AbsoluteAnchor != nil {
		_dgbfe := _ae.StartElement{Name: _ae.Name{Local: "\u0078d\u0072:\u0061\u0062\u0073\u006f\u006cu\u0074\u0065A\u006e\u0063\u0068\u006f\u0072"}}
		e.EncodeElement(_dec.AbsoluteAnchor, _dgbfe)
	}
	return nil
}

func (_ddfgd *To) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_ddfgd.CT_Marker = *NewCT_Marker()
_aadee:
	for {
		_gaed, _gbg := d.Token()
		if _gbg != nil {
			return _gbg
		}
		switch _dggc := _gaed.(type) {
		case _ae.StartElement:
			switch _dggc.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c"}:
				if _gggf := d.DecodeElement(&_ddfgd.Col, &_dggc); _gggf != nil {
					return _gggf
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006f\u006c\u004f\u0066\u0066"}:
				_dccb, _faaa := d.Token()
				if _faaa != nil {
					return _faaa
				}
				switch _fcge := _dccb.(type) {
				case _ae.CharData:
					_bfa := string(_fcge)
					_cdb, _eggb := _c.ParseUnionST_Coordinate(_bfa)
					if _eggb != nil {
						return nil
					}
					_ddfgd.ColOff = _cdb
					d.Skip()
				case _ae.EndElement:
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077"}:
				if _egdcb := d.DecodeElement(&_ddfgd.Row, &_dggc); _egdcb != nil {
					return _egdcb
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0072\u006f\u0077\u004f\u0066\u0066"}:
				_aga, _fgff := d.Token()
				if _fgff != nil {
					return _fgff
				}
				switch _gdga := _aga.(type) {
				case _ae.CharData:
					_eee := string(_gdga)
					_fdg, _caba := _c.ParseUnionST_Coordinate(_eee)
					if _caba != nil {
						return nil
					}
					_ddfgd.RowOff = _fdg
					d.Skip()
				case _ae.EndElement:
				}
			default:
				_f.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006eg\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020T\u006f\u0020\u0025\u0076", _dggc.Name)
				if _gfgc := d.Skip(); _gfgc != nil {
					return _gfgc
				}
			}
		case _ae.EndElement:
			break _aadee
		case _ae.CharData:
		}
	}
	return nil
}

func (_egb *CT_GroupShapeChoice) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
_bab:
	for {
		_gabd, _bfd := d.Token()
		if _bfd != nil {
			return _bfd
		}
		switch _dgg := _gabd.(type) {
		case _ae.StartElement:
			switch _dgg.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_ccbg := NewCT_Shape()
				if _add := d.DecodeElement(_ccbg, &_dgg); _add != nil {
					return _add
				}
				_egb.Sp = append(_egb.Sp, _ccbg)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_bdca := NewCT_GroupShape()
				if _eed := d.DecodeElement(_bdca, &_dgg); _eed != nil {
					return _eed
				}
				_egb.GrpSp = append(_egb.GrpSp, _bdca)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_adg := NewCT_GraphicalObjectFrame()
				if _gfg := d.DecodeElement(_adg, &_dgg); _gfg != nil {
					return _gfg
				}
				_egb.GraphicFrame = append(_egb.GraphicFrame, _adg)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_agc := NewCT_Connector()
				if _aba := d.DecodeElement(_agc, &_dgg); _aba != nil {
					return _aba
				}
				_egb.CxnSp = append(_egb.CxnSp, _agc)
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_ffb := NewCT_Picture()
				if _bef := d.DecodeElement(_ffb, &_dgg); _bef != nil {
					return _bef
				}
				_egb.Pic = append(_egb.Pic, _ffb)
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068ap\u0065\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076", _dgg.Name)
				if _dacd := d.Skip(); _dacd != nil {
					return _dacd
				}
			}
		case _ae.EndElement:
			break _bab
		case _ae.CharData:
		}
	}
	return nil
}

func (_fbec ST_EditAs) MarshalXML(e *_ae.Encoder, start _ae.StartElement) error {
	return e.EncodeElement(_fbec.String(), start)
}

func (_gef *CT_TwoCellAnchor) UnmarshalXML(d *_ae.Decoder, start _ae.StartElement) error {
	_gef.From = NewCT_Marker()
	_gef.To = NewCT_Marker()
	_gef.ClientData = NewCT_AnchorClientData()
	for _, _abca := range start.Attr {
		if _abca.Name.Local == "\u0065\u0064\u0069\u0074\u0041\u0073" {
			_gef.EditAsAttr.UnmarshalXMLAttr(_abca)
			continue
		}
	}
_eef:
	for {
		_cagc, _fabf := d.Token()
		if _fabf != nil {
			return _fabf
		}
		switch _babg := _cagc.(type) {
		case _ae.StartElement:
			switch _babg.Name {
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0066\u0072\u006f\u006d"}:
				if _gabg := d.DecodeElement(_gef.From, &_babg); _gabg != nil {
					return _gabg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0074\u006f"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0074\u006f"}:
				if _abaea := d.DecodeElement(_gef.To, &_babg); _abaea != nil {
					return _abaea
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0073\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0073\u0070"}:
				_gef.Choice = NewEG_ObjectChoicesChoice()
				if _daba := d.DecodeElement(&_gef.Choice.Sp, &_babg); _daba != nil {
					return _daba
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0070S\u0070"}:
				_gef.Choice = NewEG_ObjectChoicesChoice()
				if _fec := d.DecodeElement(&_gef.Choice.GrpSp, &_babg); _fec != nil {
					return _fec
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0067\u0072\u0061p\u0068\u0069\u0063\u0046\u0072\u0061\u006d\u0065"}:
				_gef.Choice = NewEG_ObjectChoicesChoice()
				if _cdgd := d.DecodeElement(&_gef.Choice.GraphicFrame, &_babg); _cdgd != nil {
					return _cdgd
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u0078\u006eS\u0070"}:
				_gef.Choice = NewEG_ObjectChoicesChoice()
				if _cbcc := d.DecodeElement(&_gef.Choice.CxnSp, &_babg); _cbcc != nil {
					return _cbcc
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0070\u0069\u0063"}:
				_gef.Choice = NewEG_ObjectChoicesChoice()
				if _bec := d.DecodeElement(&_gef.Choice.Pic, &_babg); _bec != nil {
					return _bec
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "c\u006f\u006e\u0074\u0065\u006e\u0074\u0050\u0061\u0072\u0074"}:
				_gef.Choice = NewEG_ObjectChoicesChoice()
				if _bdg := d.DecodeElement(&_gef.Choice.ContentPart, &_babg); _bdg != nil {
					return _bdg
				}
			case _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}, _ae.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fdr\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0073\u0070\u0072\u0065\u0061\u0064\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061\u0077\u0069\u006e\u0067", Local: "\u0063\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061"}:
				if _cgceg := d.DecodeElement(_gef.ClientData, &_babg); _cgceg != nil {
					return _cgceg
				}
			default:
				_f.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0077\u006f\u0043\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072\u0020\u0025v", _babg.Name)
				if _acbb := d.Skip(); _acbb != nil {
					return _acbb
				}
			}
		case _ae.EndElement:
			break _eef
		case _ae.CharData:
		}
	}
	return nil
}

func init() {
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0041nc\u0068\u006f\u0072\u0043\u006c\u0069\u0065\u006e\u0074\u0044\u0061\u0074\u0061", NewCT_AnchorClientData)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056i\u0073\u0075\u0061\u006c", NewCT_ShapeNonVisual)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065", NewCT_Shape)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "C\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u004e\u006fn\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_ConnectorNonVisual)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005fC\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072", NewCT_Connector)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_PictureNonVisual)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065", NewCT_Picture)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047\u0072\u0061\u0070\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006ae\u0063t\u0046\u0072\u0061\u006d\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c", NewCT_GraphicalObjectFrameNonVisual)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005fGr\u0061\u0070\u0068\u0069\u0063\u0061\u006c\u004f\u0062\u006a\u0065\u0063\u0074\u0046\u0072\u0061\u006d\u0065", NewCT_GraphicalObjectFrame)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047ro\u0075\u0070\u0053\u0068\u0061\u0070\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075a\u006c", NewCT_GroupShapeNonVisual)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0047\u0072\u006f\u0075\u0070\u0053\u0068\u0061\u0070\u0065", NewCT_GroupShape)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0052\u0065\u006c", NewCT_Rel)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043T\u005f\u004d\u0061\u0072\u006b\u0065r", NewCT_Marker)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054_\u0054\u0077\u006fC\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072", NewCT_TwoCellAnchor)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054_\u004f\u006e\u0065C\u0065\u006c\u006c\u0041\u006e\u0063\u0068\u006f\u0072", NewCT_OneCellAnchor)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0041\u0062\u0073\u006f\u006c\u0075\u0074\u0065\u0041n\u0063\u0068\u006f\u0072", NewCT_AbsoluteAnchor)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0043\u0054\u005f\u0044\u0072\u0061\u0077\u0069\u006e\u0067", NewCT_Drawing)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0066\u0072\u006f\u006d", NewFrom)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0074\u006f", NewTo)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0077\u0073\u0044\u0072", NewWsDr)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0045\u0047_\u004f\u0062\u006ae\u0063\u0074\u0043\u0068\u006f\u0069\u0063\u0065\u0073", NewEG_ObjectChoices)
	_ac.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0070\u0072\u0065\u0061d\u0073\u0068\u0065\u0065\u0074\u0044\u0072\u0061w\u0069\u006e\u0067", "\u0045G\u005f\u0041\u006e\u0063\u0068\u006fr", NewEG_Anchor)
}
