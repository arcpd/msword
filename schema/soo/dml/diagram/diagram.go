package diagram

import (
	_a "encoding/xml"
	_f "fmt"
	_ae "strconv"

	_fd "github.com/arcpd/msword"
	_b "github.com/arcpd/msword/common/logger"
	_bg "github.com/arcpd/msword/schema/soo/dml"
)

func (_cbbgb ST_ChildOrderType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_efbee := _a.Attr{}
	_efbee.Name = name
	switch _cbbgb {
	case ST_ChildOrderTypeUnset:
		_efbee.Value = ""
	case ST_ChildOrderTypeB:
		_efbee.Value = "\u0062"
	case ST_ChildOrderTypeT:
		_efbee.Value = "\u0074"
	}
	return _efbee, nil
}

func (_bfbc *ST_OutputShapeType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ffge, _ggecf := d.Token()
	if _ggecf != nil {
		return _ggecf
	}
	if _efef, _dgba := _ffge.(_a.EndElement); _dgba && _efef.Name == start.Name {
		*_bfbc = 1
		return nil
	}
	if _gaaef, _agcec := _ffge.(_a.CharData); !_agcec {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ffge)
	} else {
		switch string(_gaaef) {
		case "":
			*_bfbc = 0
		case "\u006e\u006f\u006e\u0065":
			*_bfbc = 1
		case "\u0063\u006f\u006e\u006e":
			*_bfbc = 2
		}
	}
	_ffge, _ggecf = d.Token()
	if _ggecf != nil {
		return _ggecf
	}
	if _ggdgf, _baaa := _ffge.(_a.EndElement); _baaa && _ggdgf.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ffge)
}

// Validate validates the CT_CTName and its children
func (_fedd *CT_CTName) Validate() error {
	return _fedd.ValidateWithPath("\u0043T\u005f\u0043\u0054\u004e\u0061\u006de")
}

func (_dbagg ST_NodeHorizontalAlignment) String() string {
	switch _dbagg {
	case 0:
		return ""
	case 1:
		return "\u006c"
	case 2:
		return "\u0063\u0074\u0072"
	case 3:
		return "\u0072"
	}
	return ""
}

func (_adggfd ST_ArrowheadStyle) String() string {
	switch _adggfd {
	case 0:
		return ""
	case 1:
		return "\u0061\u0075\u0074\u006f"
	case 2:
		return "\u0061\u0072\u0072"
	case 3:
		return "\u006e\u006f\u0041r\u0072"
	}
	return ""
}

func (_gdddf ST_DiagramTextAlignment) Validate() error { return _gdddf.ValidateWithPath("") }

func (_fgag *LayoutDefHdr) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072"
	return _fgag.CT_DiagramDefinitionHeader.MarshalXML(e, start)
}

func (_gdga ST_PyramidAccentTextMargin) ValidateWithPath(path string) error {
	switch _gdga {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gdga))
	}
	return nil
}

func (_ddbf ST_HierarchyAlignment) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_facgfc := _a.Attr{}
	_facgfc.Name = name
	switch _ddbf {
	case ST_HierarchyAlignmentUnset:
		_facgfc.Value = ""
	case ST_HierarchyAlignmentTL:
		_facgfc.Value = "\u0074\u004c"
	case ST_HierarchyAlignmentTR:
		_facgfc.Value = "\u0074\u0052"
	case ST_HierarchyAlignmentTCtrCh:
		_facgfc.Value = "\u0074\u0043\u0074\u0072\u0043\u0068"
	case ST_HierarchyAlignmentTCtrDes:
		_facgfc.Value = "\u0074C\u0074\u0072\u0044\u0065\u0073"
	case ST_HierarchyAlignmentBL:
		_facgfc.Value = "\u0062\u004c"
	case ST_HierarchyAlignmentBR:
		_facgfc.Value = "\u0062\u0052"
	case ST_HierarchyAlignmentBCtrCh:
		_facgfc.Value = "\u0062\u0043\u0074\u0072\u0043\u0068"
	case ST_HierarchyAlignmentBCtrDes:
		_facgfc.Value = "\u0062C\u0074\u0072\u0044\u0065\u0073"
	case ST_HierarchyAlignmentLT:
		_facgfc.Value = "\u006c\u0054"
	case ST_HierarchyAlignmentLB:
		_facgfc.Value = "\u006c\u0042"
	case ST_HierarchyAlignmentLCtrCh:
		_facgfc.Value = "\u006c\u0043\u0074\u0072\u0043\u0068"
	case ST_HierarchyAlignmentLCtrDes:
		_facgfc.Value = "\u006cC\u0074\u0072\u0044\u0065\u0073"
	case ST_HierarchyAlignmentRT:
		_facgfc.Value = "\u0072\u0054"
	case ST_HierarchyAlignmentRB:
		_facgfc.Value = "\u0072\u0042"
	case ST_HierarchyAlignmentRCtrCh:
		_facgfc.Value = "\u0072\u0043\u0074\u0072\u0043\u0068"
	case ST_HierarchyAlignmentRCtrDes:
		_facgfc.Value = "\u0072C\u0074\u0072\u0044\u0065\u0073"
	}
	return _facgfc, nil
}

func (_ccbb *CT_SampleData) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _ccbb.UseDefAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0075\u0073\u0065\u0044\u0065\u0066"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_ccbb.UseDefAttr))})
	}
	e.EncodeToken(start)
	if _ccbb.DataModel != nil {
		_befea := _a.StartElement{Name: _a.Name{Local: "\u0064a\u0074\u0061\u004d\u006f\u0064\u0065l"}}
		e.EncodeElement(_ccbb.DataModel, _befea)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func ParseSliceST_Ints(s string) (ST_Ints, error) { return ST_Ints{}, nil }

func (_daeda *ST_HierBranchStyle) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_daeda = 0
	case "\u006c":
		*_daeda = 1
	case "\u0072":
		*_daeda = 2
	case "\u0068\u0061\u006e\u0067":
		*_daeda = 3
	case "\u0073\u0074\u0064":
		*_daeda = 4
	case "\u0069\u006e\u0069\u0074":
		*_daeda = 5
	}
	return nil
}

func (_decc ST_ContinueDirection) Validate() error { return _decc.ValidateWithPath("") }

func (_bfddf ST_TextBlockDirection) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_fcegg := _a.Attr{}
	_fcegg.Name = name
	switch _bfddf {
	case ST_TextBlockDirectionUnset:
		_fcegg.Value = ""
	case ST_TextBlockDirectionHorz:
		_fcegg.Value = "\u0068\u006f\u0072\u007a"
	case ST_TextBlockDirectionVert:
		_fcegg.Value = "\u0076\u0065\u0072\u0074"
	}
	return _fcegg, nil
}

func (_ba *CT_AdjLst) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _ba.Adj != nil {
		_ad := _a.StartElement{Name: _a.Name{Local: "\u0061\u0064\u006a"}}
		for _, _ccf := range _ba.Adj {
			e.EncodeElement(_ccf, _ad)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

const (
	ST_PyramidAccentTextMarginUnset ST_PyramidAccentTextMargin = 0
	ST_PyramidAccentTextMarginStep  ST_PyramidAccentTextMargin = 1
	ST_PyramidAccentTextMarginStack ST_PyramidAccentTextMargin = 2
)

func (_adeb *ST_GrowDirection) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_adeb = 0
	case "\u0074\u004c":
		*_adeb = 1
	case "\u0074\u0052":
		*_adeb = 2
	case "\u0062\u004c":
		*_adeb = 3
	case "\u0062\u0052":
		*_adeb = 4
	}
	return nil
}

// Validate validates the CT_NumericRule and its children
func (_bbaaa *CT_NumericRule) Validate() error {
	return _bbaaa.ValidateWithPath("\u0043\u0054\u005f\u004e\u0075\u006d\u0065\u0072\u0069c\u0052\u0075\u006c\u0065")
}

type StyleDef struct{ CT_StyleDefinition }

func (_aegbd ST_NodeHorizontalAlignment) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_aegbd.String(), start)
}

func (_bge *AG_ConstraintAttributes) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _d := range start.Attr {
		if _d.Name.Local == "\u0074\u0079\u0070\u0065" {
			_bge.TypeAttr.UnmarshalXMLAttr(_d)
			continue
		}
		if _d.Name.Local == "\u0066\u006f\u0072" {
			_bge.ForAttr.UnmarshalXMLAttr(_d)
			continue
		}
		if _d.Name.Local == "\u0066o\u0072\u004e\u0061\u006d\u0065" {
			_fa, _gcb := _d.Value, error(nil)
			if _gcb != nil {
				return _gcb
			}
			_bge.ForNameAttr = &_fa
			continue
		}
		if _d.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_bge.PtTypeAttr.UnmarshalXMLAttr(_d)
			continue
		}
	}
	for {
		_fe, _af := d.Token()
		if _af != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0041\u0047\u005f\u0043\u006f\u006es\u0074\u0072\u0061\u0069\u006e\u0074\u0041t\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0073\u003a\u0020%\u0073", _af)
		}
		if _eg, _aa := _fe.(_a.EndElement); _aa && _eg.Name == start.Name {
			break
		}
	}
	return nil
}

func (_efcc *ST_ArrowheadStyle) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_efcc = 0
	case "\u0061\u0075\u0074\u006f":
		*_efcc = 1
	case "\u0061\u0072\u0072":
		*_efcc = 2
	case "\u006e\u006f\u0041r\u0072":
		*_efcc = 3
	}
	return nil
}

const (
	ST_DiagramHorizontalAlignmentUnset ST_DiagramHorizontalAlignment = 0
	ST_DiagramHorizontalAlignmentL     ST_DiagramHorizontalAlignment = 1
	ST_DiagramHorizontalAlignmentCtr   ST_DiagramHorizontalAlignment = 2
	ST_DiagramHorizontalAlignmentR     ST_DiagramHorizontalAlignment = 3
	ST_DiagramHorizontalAlignmentNone  ST_DiagramHorizontalAlignment = 4
)

func (_caadc ST_ConnectorDimension) Validate() error { return _caadc.ValidateWithPath("") }

func (_dddgg ST_FunctionArgument) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _dddgg.ST_VariableType != ST_VariableTypeUnset {
		e.EncodeToken(_a.CharData(_dddgg.ST_VariableType.String()))
	}
	return e.EncodeToken(_a.EndElement{Name: start.Name})
}

func NewCT_TextProps() *CT_TextProps { _cedc := &CT_TextProps{}; return _cedc }

type CT_Parameter struct {
	TypeAttr ST_ParameterId
	ValAttr  ST_ParameterVal
}

func (_gcga *CT_Categories) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_eaa:
	for {
		_bcae, _gdg := d.Token()
		if _gdg != nil {
			return _gdg
		}
		switch _dfg := _bcae.(type) {
		case _a.StartElement:
			switch _dfg.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074"}:
				_gce := NewCT_Category()
				if _bgag := d.DecodeElement(_gce, &_dfg); _bgag != nil {
					return _bgag
				}
				_gcga.Cat = append(_gcga.Cat, _gce)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043a\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073 \u0025\u0076", _dfg.Name)
				if _cbg := d.Skip(); _cbg != nil {
					return _cbg
				}
			}
		case _a.EndElement:
			break _eaa
		case _a.CharData:
		}
	}
	return nil
}

func (_caaf *CT_ChildMax) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _caaf.ValAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", *_caaf.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_bebd ST_ParameterVal) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _bebd.ST_DiagramHorizontalAlignment != ST_DiagramHorizontalAlignmentUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_DiagramHorizontalAlignment.String()))
	}
	if _bebd.ST_VerticalAlignment != ST_VerticalAlignmentUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_VerticalAlignment.String()))
	}
	if _bebd.ST_ChildDirection != ST_ChildDirectionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_ChildDirection.String()))
	}
	if _bebd.ST_ChildAlignment != ST_ChildAlignmentUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_ChildAlignment.String()))
	}
	if _bebd.ST_SecondaryChildAlignment != ST_SecondaryChildAlignmentUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_SecondaryChildAlignment.String()))
	}
	if _bebd.ST_LinearDirection != ST_LinearDirectionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_LinearDirection.String()))
	}
	if _bebd.ST_SecondaryLinearDirection != ST_SecondaryLinearDirectionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_SecondaryLinearDirection.String()))
	}
	if _bebd.ST_StartingElement != ST_StartingElementUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_StartingElement.String()))
	}
	if _bebd.ST_BendPoint != ST_BendPointUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_BendPoint.String()))
	}
	if _bebd.ST_ConnectorRouting != ST_ConnectorRoutingUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_ConnectorRouting.String()))
	}
	if _bebd.ST_ArrowheadStyle != ST_ArrowheadStyleUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_ArrowheadStyle.String()))
	}
	if _bebd.ST_ConnectorDimension != ST_ConnectorDimensionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_ConnectorDimension.String()))
	}
	if _bebd.ST_RotationPath != ST_RotationPathUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_RotationPath.String()))
	}
	if _bebd.ST_CenterShapeMapping != ST_CenterShapeMappingUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_CenterShapeMapping.String()))
	}
	if _bebd.ST_NodeHorizontalAlignment != ST_NodeHorizontalAlignmentUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_NodeHorizontalAlignment.String()))
	}
	if _bebd.ST_NodeVerticalAlignment != ST_NodeVerticalAlignmentUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_NodeVerticalAlignment.String()))
	}
	if _bebd.ST_FallbackDimension != ST_FallbackDimensionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_FallbackDimension.String()))
	}
	if _bebd.ST_TextDirection != ST_TextDirectionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_TextDirection.String()))
	}
	if _bebd.ST_PyramidAccentPosition != ST_PyramidAccentPositionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_PyramidAccentPosition.String()))
	}
	if _bebd.ST_PyramidAccentTextMargin != ST_PyramidAccentTextMarginUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_PyramidAccentTextMargin.String()))
	}
	if _bebd.ST_TextBlockDirection != ST_TextBlockDirectionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_TextBlockDirection.String()))
	}
	if _bebd.ST_TextAnchorHorizontal != ST_TextAnchorHorizontalUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_TextAnchorHorizontal.String()))
	}
	if _bebd.ST_TextAnchorVertical != ST_TextAnchorVerticalUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_TextAnchorVertical.String()))
	}
	if _bebd.ST_DiagramTextAlignment != ST_DiagramTextAlignmentUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_DiagramTextAlignment.String()))
	}
	if _bebd.ST_AutoTextRotation != ST_AutoTextRotationUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_AutoTextRotation.String()))
	}
	if _bebd.ST_GrowDirection != ST_GrowDirectionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_GrowDirection.String()))
	}
	if _bebd.ST_FlowDirection != ST_FlowDirectionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_FlowDirection.String()))
	}
	if _bebd.ST_ContinueDirection != ST_ContinueDirectionUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_ContinueDirection.String()))
	}
	if _bebd.ST_Breakpoint != ST_BreakpointUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_Breakpoint.String()))
	}
	if _bebd.ST_Offset != ST_OffsetUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_Offset.String()))
	}
	if _bebd.ST_HierarchyAlignment != ST_HierarchyAlignmentUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_HierarchyAlignment.String()))
	}
	if _bebd.Int32 != nil {
		e.EncodeToken(_a.CharData(_f.Sprintf("\u0025\u0064", *_bebd.Int32)))
	}
	if _bebd.Float64 != nil {
		e.EncodeToken(_a.CharData(_f.Sprintf("\u0025\u0066", *_bebd.Float64)))
	}
	if _bebd.Bool != nil {
		e.EncodeToken(_a.CharData(_f.Sprintf("\u0025\u0064", _feae(*_bebd.Bool))))
	}
	if _bebd.StringVal != nil {
		e.EncodeToken(_a.CharData(*_bebd.StringVal))
	}
	if _bebd.ST_ConnectorPoint != ST_ConnectorPointUnset {
		e.EncodeToken(_a.CharData(_bebd.ST_ConnectorPoint.String()))
	}
	return e.EncodeToken(_a.EndElement{Name: start.Name})
}

func (_bbbbd ST_AnimOneStr) String() string {
	switch _bbbbd {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u006f\u006e\u0065"
	case 3:
		return "\u0062\u0072\u0061\u006e\u0063\u0068"
	}
	return ""
}

// Validate validates the CT_ElemPropSet and its children
func (_bba *CT_ElemPropSet) Validate() error {
	return _bba.ValidateWithPath("\u0043\u0054\u005f\u0045\u006c\u0065\u006d\u0050\u0072o\u0070\u0053\u0065\u0074")
}

func NewCT_AdjLst() *CT_AdjLst { _dcb := &CT_AdjLst{}; return _dcb }

func (_eedc *ST_VerticalAlignment) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ccbeg, _eageef := d.Token()
	if _eageef != nil {
		return _eageef
	}
	if _agff, _bfeg := _ccbeg.(_a.EndElement); _bfeg && _agff.Name == start.Name {
		*_eedc = 1
		return nil
	}
	if _edec, _bcaf := _ccbeg.(_a.CharData); !_bcaf {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ccbeg)
	} else {
		switch string(_edec) {
		case "":
			*_eedc = 0
		case "\u0074":
			*_eedc = 1
		case "\u006d\u0069\u0064":
			*_eedc = 2
		case "\u0062":
			*_eedc = 3
		case "\u006e\u006f\u006e\u0065":
			*_eedc = 4
		}
	}
	_ccbeg, _eageef = d.Token()
	if _eageef != nil {
		return _eageef
	}
	if _egebf, _deaf := _ccbeg.(_a.EndElement); _deaf && _egebf.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ccbeg)
}

func (_dafbc ST_ChildAlignment) ValidateWithPath(path string) error {
	switch _dafbc {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dafbc))
	}
	return nil
}

func (_cbfbf ST_ConnectorRouting) String() string {
	switch _cbfbf {
	case 0:
		return ""
	case 1:
		return "\u0073\u0074\u0072\u0061"
	case 2:
		return "\u0062\u0065\u006e\u0064"
	case 3:
		return "\u0063\u0075\u0072v\u0065"
	case 4:
		return "\u006co\u006e\u0067\u0043\u0075\u0072\u0076e"
	}
	return ""
}

func NewCT_Constraints() *CT_Constraints { _bdbba := &CT_Constraints{}; return _bdbba }

func (_cdabc *CT_HierBranchStyle) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _cdabc.ValAttr != ST_HierBranchStyleUnset {
		_cgcd, _ggbac := _cdabc.ValAttr.MarshalXMLAttr(_a.Name{Local: "\u0076\u0061\u006c"})
		if _ggbac != nil {
			return _ggbac
		}
		start.Attr = append(start.Attr, _cgcd)
	}
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_fabg *CT_Colors) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _fdba := range start.Attr {
		if _fdba.Name.Local == "\u006d\u0065\u0074\u0068" {
			_fabg.MethAttr.UnmarshalXMLAttr(_fdba)
			continue
		}
		if _fdba.Name.Local == "\u0068\u0075\u0065\u0044\u0069\u0072" {
			_fabg.HueDirAttr.UnmarshalXMLAttr(_fdba)
			continue
		}
	}
_daec:
	for {
		_abc, _dfae := d.Token()
		if _dfae != nil {
			return _dfae
		}
		switch _ecdc := _abc.(type) {
		case _a.StartElement:
			switch _ecdc.Name {
			default:
				_b.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0043\u006f\u006co\u0072\u0073 \u0025\u0076", _ecdc.Name)
				if _daae := d.Skip(); _daae != nil {
					return _daae
				}
			}
		case _a.EndElement:
			break _daec
		case _a.CharData:
		}
	}
	return nil
}

type CT_AnimLvl struct{ ValAttr ST_AnimLvlStr }

type CT_DiagramDefinitionHeaderLst struct{ LayoutDefHdr []*CT_DiagramDefinitionHeader }

func (_dabbg *ST_DiagramTextAlignment) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_dabbg = 0
	case "\u006c":
		*_dabbg = 1
	case "\u0063\u0074\u0072":
		*_dabbg = 2
	case "\u0072":
		*_dabbg = 3
	}
	return nil
}

func (_gfeb *CT_Pt) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _cadc := range start.Attr {
		if _cadc.Name.Local == "\u006do\u0064\u0065\u006c\u0049\u0064" {
			_dfff, _dcgf := ParseUnionST_ModelId(_cadc.Value)
			if _dcgf != nil {
				return _dcgf
			}
			_gfeb.ModelIdAttr = _dfff
			continue
		}
		if _cadc.Name.Local == "\u0074\u0079\u0070\u0065" {
			_gfeb.TypeAttr.UnmarshalXMLAttr(_cadc)
			continue
		}
		if _cadc.Name.Local == "\u0063\u0078\u006eI\u0064" {
			_adgb, _cfag := ParseUnionST_ModelId(_cadc.Value)
			if _cfag != nil {
				return _cfag
			}
			_gfeb.CxnIdAttr = &_adgb
			continue
		}
	}
_cdeg:
	for {
		_aefc, _ceeg := d.Token()
		if _ceeg != nil {
			return _ceeg
		}
		switch _gag := _aefc.(type) {
		case _a.StartElement:
			switch _gag.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0053e\u0074"}:
				_gfeb.PrSet = NewCT_ElemPropSet()
				if _dcbf := d.DecodeElement(_gfeb.PrSet, &_gag); _dcbf != nil {
					return _dcbf
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0070\u0050\u0072"}:
				_gfeb.SpPr = _bg.NewCT_ShapeProperties()
				if _aca := d.DecodeElement(_gfeb.SpPr, &_gag); _aca != nil {
					return _aca
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074"}:
				_gfeb.T = _bg.NewCT_TextBody()
				if _fgeaf := d.DecodeElement(_gfeb.T, &_gag); _fgeaf != nil {
					return _fgeaf
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_gfeb.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _ecbe := d.DecodeElement(_gfeb.ExtLst, &_gag); _ecbe != nil {
					return _ecbe
				}
			default:
				_b.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0043\u0054\u005fP\u0074 \u0025\u0076", _gag.Name)
				if _ggagf := d.Skip(); _ggagf != nil {
					return _ggagf
				}
			}
		case _a.EndElement:
			break _cdeg
		case _a.CharData:
		}
	}
	return nil
}

// Validate validates the CT_StyleDefinition and its children
func (_gfag *CT_StyleDefinition) Validate() error {
	return _gfag.ValidateWithPath("\u0043T\u005fS\u0074\u0079\u006c\u0065\u0044e\u0066\u0069n\u0069\u0074\u0069\u006f\u006e")
}

func NewCT_CTName() *CT_CTName { _bggb := &CT_CTName{}; return _bggb }

func (_geaaf ST_AxisType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_geaaf.String(), start)
}

func (_gdcg ST_ConnectorRouting) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_gdcg.String(), start)
}

func (_deeb ST_ChildAlignment) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_deeb.String(), start)
}

// Validate validates the CT_PtList and its children
func (_fgdd *CT_PtList) Validate() error {
	return _fgdd.ValidateWithPath("\u0043T\u005f\u0050\u0074\u004c\u0069\u0073t")
}

func NewCT_Categories() *CT_Categories { _fdd := &CT_Categories{}; return _fdd }

func (_gba *CT_AnimOne) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _dfc := range start.Attr {
		if _dfc.Name.Local == "\u0076\u0061\u006c" {
			_gba.ValAttr.UnmarshalXMLAttr(_dfc)
			continue
		}
	}
	for {
		_abdd, _ef := d.Token()
		if _ef != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004f\u006e\u0065\u003a\u0020%\u0073", _ef)
		}
		if _bac, _fca := _abdd.(_a.EndElement); _fca && _bac.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_ChildPref and its children, prefixing error messages with path
func (_baa *CT_ChildPref) ValidateWithPath(path string) error {
	if _baa.ValAttr != nil {
		if *_baa.ValAttr < -1 {
			return _f.Errorf("\u0025\u0073/m\u002e\u0056\u0061l\u0041\u0074\u0074\u0072 mu\u0073t \u0062\u0065\u0020\u003e\u003d\u0020\u002d1 \u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, *_baa.ValAttr)
		}
	}
	return nil
}

// ValidateWithPath validates the CT_LayoutNode and its children, prefixing error messages with path
func (_bece *CT_LayoutNode) ValidateWithPath(path string) error {
	if _gafee := _bece.ChOrderAttr.ValidateWithPath(path + "\u002f\u0043\u0068O\u0072\u0064\u0065\u0072\u0041\u0074\u0074\u0072"); _gafee != nil {
		return _gafee
	}
	for _ffagf, _edee := range _bece.Alg {
		if _ceded := _edee.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0041\u006c\u0067\u005b\u0025\u0064\u005d", path, _ffagf)); _ceded != nil {
			return _ceded
		}
	}
	for _gfad, _dabdb := range _bece.Shape {
		if _gcdd := _dabdb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fS\u0068\u0061\u0070\u0065\u005b\u0025\u0064\u005d", path, _gfad)); _gcdd != nil {
			return _gcdd
		}
	}
	for _fafca, _daadb := range _bece.PresOf {
		if _bcaa := _daadb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0050\u0072\u0065\u0073\u004f\u0066\u005b\u0025\u0064\u005d", path, _fafca)); _bcaa != nil {
			return _bcaa
		}
	}
	for _ebdd, _babg := range _bece.ConstrLst {
		if _bega := _babg.ValidateWithPath(_f.Sprintf("\u0025\u0073/\u0043\u006f\u006es\u0074\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _ebdd)); _bega != nil {
			return _bega
		}
	}
	for _agag, _dgcb := range _bece.RuleLst {
		if _bbbbb := _dgcb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0052\u0075\u006c\u0065\u004c\u0073t\u005b\u0025\u0064\u005d", path, _agag)); _bbbbb != nil {
			return _bbbbb
		}
	}
	for _afff, _bbdde := range _bece.VarLst {
		if _ceeb := _bbdde.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0056\u0061\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _afff)); _ceeb != nil {
			return _ceeb
		}
	}
	for _ffc, _fbfc := range _bece.ForEach {
		if _affd := _fbfc.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0046\u006f\u0072\u0045\u0061\u0063h\u005b\u0025\u0064\u005d", path, _ffc)); _affd != nil {
			return _affd
		}
	}
	for _cga, _afaaa := range _bece.LayoutNode {
		if _gdab := _afaaa.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064e\u005b\u0025\u0064\u005d", path, _cga)); _gdab != nil {
			return _gdab
		}
	}
	for _geed, _cdgg := range _bece.Choose {
		if _bdde := _cdgg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u006f\u0073\u0065\u005b\u0025\u0064\u005d", path, _geed)); _bdde != nil {
			return _bdde
		}
	}
	for _fecf, _afcag := range _bece.ExtLst {
		if _ddg := _afcag.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0045\u0078\u0074\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _fecf)); _ddg != nil {
			return _ddg
		}
	}
	return nil
}

func ParseSliceST_AxisTypes(s string) (ST_AxisTypes, error) { return ST_AxisTypes{}, nil }

type StyleDefHdr struct{ CT_StyleDefinitionHeader }

// Validate validates the CT_StyleDefinitionHeader and its children
func (_gfbc *CT_StyleDefinitionHeader) Validate() error {
	return _gfbc.ValidateWithPath("\u0043T\u005f\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0069\u006ei\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065\u0072")
}

func (_cgbgc *ST_PtType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_adbed, _efdcd := d.Token()
	if _efdcd != nil {
		return _efdcd
	}
	if _dabae, _fgfcg := _adbed.(_a.EndElement); _fgfcg && _dabae.Name == start.Name {
		*_cgbgc = 1
		return nil
	}
	if _ggae, _abaf := _adbed.(_a.CharData); !_abaf {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _adbed)
	} else {
		switch string(_ggae) {
		case "":
			*_cgbgc = 0
		case "\u006e\u006f\u0064\u0065":
			*_cgbgc = 1
		case "\u0061\u0073\u0073\u0074":
			*_cgbgc = 2
		case "\u0064\u006f\u0063":
			*_cgbgc = 3
		case "\u0070\u0072\u0065\u0073":
			*_cgbgc = 4
		case "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073":
			*_cgbgc = 5
		case "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073":
			*_cgbgc = 6
		}
	}
	_adbed, _efdcd = d.Token()
	if _efdcd != nil {
		return _efdcd
	}
	if _efgc, _cgaa := _adbed.(_a.EndElement); _cgaa && _efgc.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _adbed)
}

func (_egf *CT_BulletEnabled) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _gcbe := range start.Attr {
		if _gcbe.Name.Local == "\u0076\u0061\u006c" {
			_bfb, _gef := _ae.ParseBool(_gcbe.Value)
			if _gef != nil {
				return _gef
			}
			_egf.ValAttr = &_bfb
			continue
		}
	}
	for {
		_bbd, _cfb := d.Token()
		if _cfb != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0043\u0054\u005f\u0042\u0075\u006c\u006ce\u0074E\u006e\u0061\u0062\u006c\u0065\u0064\u003a \u0025\u0073", _cfb)
		}
		if _facg, _eab := _bbd.(_a.EndElement); _eab && _facg.Name == start.Name {
			break
		}
	}
	return nil
}

func ParseUnionST_ModelId(s string) (ST_ModelId, error) { return ST_ModelId{}, nil }

func (_ebbd *CT_StyleDefinitionHeaderLst) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _ebbd.StyleDefHdr != nil {
		_adggf := _a.StartElement{Name: _a.Name{Local: "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072"}}
		for _, _gcgg := range _ebbd.StyleDefHdr {
			e.EncodeElement(_gcgg, _adggf)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_fdef ST_ArrowheadStyle) ValidateWithPath(path string) error {
	switch _fdef {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fdef))
	}
	return nil
}

func (_feaf *CT_OrgChart) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _feaf.ValAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_feaf.ValAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

type CT_Direction struct{ ValAttr ST_Direction }

func (_fefag ST_CxnType) Validate() error { return _fefag.ValidateWithPath("") }

func (_cdb *CT_ColorTransform) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _cdb.UniqueIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_cdb.UniqueIdAttr)})
	}
	if _cdb.MinVerAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _f.Sprintf("\u0025\u0076", *_cdb.MinVerAttr)})
	}
	e.EncodeToken(start)
	if _cdb.Title != nil {
		_ada := _a.StartElement{Name: _a.Name{Local: "\u0074\u0069\u0074l\u0065"}}
		for _, _dfa := range _cdb.Title {
			e.EncodeElement(_dfa, _ada)
		}
	}
	if _cdb.Desc != nil {
		_baed := _a.StartElement{Name: _a.Name{Local: "\u0064\u0065\u0073\u0063"}}
		for _, _cgg := range _cdb.Desc {
			e.EncodeElement(_cgg, _baed)
		}
	}
	if _cdb.CatLst != nil {
		_aaab := _a.StartElement{Name: _a.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_cdb.CatLst, _aaab)
	}
	if _cdb.StyleLbl != nil {
		_cfa := _a.StartElement{Name: _a.Name{Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}}
		for _, _bedc := range _cdb.StyleLbl {
			e.EncodeElement(_bedc, _cfa)
		}
	}
	if _cdb.ExtLst != nil {
		_gdda := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_cdb.ExtLst, _gdda)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the StyleDefHdr and its children, prefixing error messages with path
func (_bcdgf *StyleDefHdr) ValidateWithPath(path string) error {
	if _eceaf := _bcdgf.CT_StyleDefinitionHeader.ValidateWithPath(path); _eceaf != nil {
		return _eceaf
	}
	return nil
}

func (_caec *CT_DiagramDefinition) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_caec.LayoutNode = NewCT_LayoutNode()
	for _, _beed := range start.Attr {
		if _beed.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_gfab, _bbfb := _beed.Value, error(nil)
			if _bbfb != nil {
				return _bbfb
			}
			_caec.UniqueIdAttr = &_gfab
			continue
		}
		if _beed.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_eddd, _gaf := _beed.Value, error(nil)
			if _gaf != nil {
				return _gaf
			}
			_caec.MinVerAttr = &_eddd
			continue
		}
		if _beed.Name.Local == "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065" {
			_fdaa, _eecgc := _beed.Value, error(nil)
			if _eecgc != nil {
				return _eecgc
			}
			_caec.DefStyleAttr = &_fdaa
			continue
		}
	}
_agfb:
	for {
		_gceg, _fedea := d.Token()
		if _fedea != nil {
			return _fedea
		}
		switch _bacc := _gceg.(type) {
		case _a.StartElement:
			switch _bacc.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_dfaec := NewCT_Name()
				if _faae := d.DecodeElement(_dfaec, &_bacc); _faae != nil {
					return _faae
				}
				_caec.Title = append(_caec.Title, _dfaec)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_beeg := NewCT_Description()
				if _abfe := d.DecodeElement(_beeg, &_bacc); _abfe != nil {
					return _abfe
				}
				_caec.Desc = append(_caec.Desc, _beeg)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_caec.CatLst = NewCT_Categories()
				if _fbcd := d.DecodeElement(_caec.CatLst, &_bacc); _fbcd != nil {
					return _fbcd
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0061\u006d\u0070\u0044\u0061\u0074\u0061"}:
				_caec.SampData = NewCT_SampleData()
				if _gca := d.DecodeElement(_caec.SampData, &_bacc); _gca != nil {
					return _gca
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073t\u0079\u006c\u0065\u0044\u0061\u0074a"}:
				_caec.StyleData = NewCT_SampleData()
				if _dfbd := d.DecodeElement(_caec.StyleData, &_bacc); _dfbd != nil {
					return _dfbd
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063l\u0072\u0044\u0061\u0074\u0061"}:
				_caec.ClrData = NewCT_SampleData()
				if _dge := d.DecodeElement(_caec.ClrData, &_bacc); _dge != nil {
					return _dge
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				if _cfd := d.DecodeElement(_caec.LayoutNode, &_bacc); _cfd != nil {
					return _cfd
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_caec.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _cgc := d.DecodeElement(_caec.ExtLst, &_bacc); _cgc != nil {
					return _cgc
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070o\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006de\u006et \u006f\u006e\u0020\u0043\u0054\u005f\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0020\u0025\u0076", _bacc.Name)
				if _dgdg := d.Skip(); _dgdg != nil {
					return _dgdg
				}
			}
		case _a.EndElement:
			break _agfb
		case _a.CharData:
		}
	}
	return nil
}

type CT_ColorTransform struct {
	UniqueIdAttr *string
	MinVerAttr   *string
	Title        []*CT_CTName
	Desc         []*CT_CTDescription
	CatLst       *CT_CTCategories
	StyleLbl     []*CT_CTStyleLabel
	ExtLst       *_bg.CT_OfficeArtExtensionList
}

func (_defcf ST_PtType) Validate() error { return _defcf.ValidateWithPath("") }

type CT_CTCategory struct {
	TypeAttr string
	PriAttr  uint32
}

// Validate validates the CT_Pt and its children
func (_fdfc *CT_Pt) Validate() error { return _fdfc.ValidateWithPath("\u0043\u0054\u005fP\u0074") }

func (_dbgf *LayoutDefHdrLst) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006ca\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074"
	return _dbgf.CT_DiagramDefinitionHeaderLst.MarshalXML(e, start)
}

func (_bfcg ST_ChildAlignment) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_abdgg := _a.Attr{}
	_abdgg.Name = name
	switch _bfcg {
	case ST_ChildAlignmentUnset:
		_abdgg.Value = ""
	case ST_ChildAlignmentT:
		_abdgg.Value = "\u0074"
	case ST_ChildAlignmentB:
		_abdgg.Value = "\u0062"
	case ST_ChildAlignmentL:
		_abdgg.Value = "\u006c"
	case ST_ChildAlignmentR:
		_abdgg.Value = "\u0072"
	}
	return _abdgg, nil
}

func (_dde *CT_Cxn) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _bgga := range start.Attr {
		if _bgga.Name.Local == "\u006do\u0064\u0065\u006c\u0049\u0064" {
			_bbcf, _cec := ParseUnionST_ModelId(_bgga.Value)
			if _cec != nil {
				return _cec
			}
			_dde.ModelIdAttr = _bbcf
			continue
		}
		if _bgga.Name.Local == "\u0074\u0079\u0070\u0065" {
			_dde.TypeAttr.UnmarshalXMLAttr(_bgga)
			continue
		}
		if _bgga.Name.Local == "\u0073\u0072\u0063I\u0064" {
			_aeb, _cbcb := ParseUnionST_ModelId(_bgga.Value)
			if _cbcb != nil {
				return _cbcb
			}
			_dde.SrcIdAttr = _aeb
			continue
		}
		if _bgga.Name.Local == "\u0064\u0065\u0073\u0074\u0049\u0064" {
			_agcfc, _afdd := ParseUnionST_ModelId(_bgga.Value)
			if _afdd != nil {
				return _afdd
			}
			_dde.DestIdAttr = _agcfc
			continue
		}
		if _bgga.Name.Local == "\u0073\u0072\u0063\u004f\u0072\u0064" {
			_ffad, _eaed := _ae.ParseUint(_bgga.Value, 10, 32)
			if _eaed != nil {
				return _eaed
			}
			_dde.SrcOrdAttr = uint32(_ffad)
			continue
		}
		if _bgga.Name.Local == "\u0064e\u0073\u0074\u004f\u0072\u0064" {
			_ggeg, _afae := _ae.ParseUint(_bgga.Value, 10, 32)
			if _afae != nil {
				return _afae
			}
			_dde.DestOrdAttr = uint32(_ggeg)
			continue
		}
		if _bgga.Name.Local == "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073\u0049\u0064" {
			_fbaa, _fede := ParseUnionST_ModelId(_bgga.Value)
			if _fede != nil {
				return _fede
			}
			_dde.ParTransIdAttr = &_fbaa
			continue
		}
		if _bgga.Name.Local == "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073\u0049\u0064" {
			_fddb, _cfebd := ParseUnionST_ModelId(_bgga.Value)
			if _cfebd != nil {
				return _cfebd
			}
			_dde.SibTransIdAttr = &_fddb
			continue
		}
		if _bgga.Name.Local == "\u0070\u0072\u0065\u0073\u0049\u0064" {
			_bda, _gdgc := _bgga.Value, error(nil)
			if _gdgc != nil {
				return _gdgc
			}
			_dde.PresIdAttr = &_bda
			continue
		}
	}
_baedg:
	for {
		_cbec, _ebfb := d.Token()
		if _ebfb != nil {
			return _ebfb
		}
		switch _gfa := _cbec.(type) {
		case _a.StartElement:
			switch _gfa.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_dde.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _bfbd := d.DecodeElement(_dde.ExtLst, &_gfa); _bfbd != nil {
					return _bfbd
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u0078\u006e\u0020\u0025\u0076", _gfa.Name)
				if _fbac := d.Skip(); _fbac != nil {
					return _fbac
				}
			}
		case _a.EndElement:
			break _baedg
		case _a.CharData:
		}
	}
	return nil
}

func (_bege *CT_Constraint) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _bege.OpAttr != ST_BoolOperatorUnset {
		_gfea, _gfdb := _bege.OpAttr.MarshalXMLAttr(_a.Name{Local: "\u006f\u0070"})
		if _gfdb != nil {
			return _gfdb
		}
		start.Attr = append(start.Attr, _gfea)
	}
	if _bege.ValAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", *_bege.ValAttr)})
	}
	if _bege.FactAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0066\u0061\u0063\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_bege.FactAttr)})
	}
	if _bege.TypeAttr != ST_ConstraintTypeUnset {
		_bfa, _dfab := _bege.TypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _dfab != nil {
			return _dfab
		}
		start.Attr = append(start.Attr, _bfa)
	}
	if _bege.ForAttr != ST_ConstraintRelationshipUnset {
		_edca, _bdbb := _bege.ForAttr.MarshalXMLAttr(_a.Name{Local: "\u0066\u006f\u0072"})
		if _bdbb != nil {
			return _bdbb
		}
		start.Attr = append(start.Attr, _edca)
	}
	if _bege.ForNameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0066o\u0072\u004e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_bege.ForNameAttr)})
	}
	if _bege.PtTypeAttr != ST_ElementTypeUnset {
		_caed, _dcgd := _bege.PtTypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"})
		if _dcgd != nil {
			return _dcgd
		}
		start.Attr = append(start.Attr, _caed)
	}
	if _bege.RefTypeAttr != ST_ConstraintTypeUnset {
		_adfc, _dbca := _bege.RefTypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0072e\u0066\u0054\u0079\u0070\u0065"})
		if _dbca != nil {
			return _dbca
		}
		start.Attr = append(start.Attr, _adfc)
	}
	if _bege.RefForAttr != ST_ConstraintRelationshipUnset {
		_begg, _aacg := _bege.RefForAttr.MarshalXMLAttr(_a.Name{Local: "\u0072\u0065\u0066\u0046\u006f\u0072"})
		if _aacg != nil {
			return _aacg
		}
		start.Attr = append(start.Attr, _begg)
	}
	if _bege.RefForNameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u0065\u0066\u0046\u006f\u0072\u004e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_bege.RefForNameAttr)})
	}
	if _bege.RefPtTypeAttr != ST_ElementTypeUnset {
		_afca, _efd := _bege.RefPtTypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0072e\u0066\u0050\u0074\u0054\u0079\u0070e"})
		if _efd != nil {
			return _efd
		}
		start.Attr = append(start.Attr, _afca)
	}
	e.EncodeToken(start)
	if _bege.ExtLst != nil {
		_afac := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_bege.ExtLst, _afac)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

type DataModel struct{ CT_DataModel }

func NewCT_StyleDefinitionHeaderLst() *CT_StyleDefinitionHeaderLst {
	_cebgc := &CT_StyleDefinitionHeaderLst{}
	return _cebgc
}

func (_gfgbd ST_AnimLvlStr) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_gfgbd.String(), start)
}

func (_eagee *ColorsDefHdrLst) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0063o\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074"
	return _eagee.CT_ColorTransformHeaderLst.MarshalXML(e, start)
}

type ST_ElementTypes []ST_ElementType

func (_dfacc *ST_PyramidAccentTextMargin) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ccfg, _bfba := d.Token()
	if _bfba != nil {
		return _bfba
	}
	if _bfgga, _gaggb := _ccfg.(_a.EndElement); _gaggb && _bfgga.Name == start.Name {
		*_dfacc = 1
		return nil
	}
	if _baccc, _ebeea := _ccfg.(_a.CharData); !_ebeea {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ccfg)
	} else {
		switch string(_baccc) {
		case "":
			*_dfacc = 0
		case "\u0073\u0074\u0065\u0070":
			*_dfacc = 1
		case "\u0073\u0074\u0061c\u006b":
			*_dfacc = 2
		}
	}
	_ccfg, _bfba = d.Token()
	if _bfba != nil {
		return _bfba
	}
	if _bebedg, _dcbb := _ccfg.(_a.EndElement); _dcbb && _bebedg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ccfg)
}

// Validate validates the RelIds and its children
func (_fbce *RelIds) Validate() error {
	return _fbce.ValidateWithPath("\u0052\u0065\u006c\u0049\u0064\u0073")
}

func (_afcd ST_StartingElement) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_afcd.String(), start)
}

func (_fgab *ST_ResizeHandlesStr) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_fgab = 0
	case "\u0065\u0078\u0061c\u0074":
		*_fgab = 1
	case "\u0072\u0065\u006c":
		*_fgab = 2
	}
	return nil
}

func (_bffgf ST_LinearDirection) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_bffgf.String(), start)
}

func (_ccgd *ST_BoolOperator) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ccgd = 0
	case "\u006e\u006f\u006e\u0065":
		*_ccgd = 1
	case "\u0065\u0071\u0075":
		*_ccgd = 2
	case "\u0067\u0074\u0065":
		*_ccgd = 3
	case "\u006c\u0074\u0065":
		*_ccgd = 4
	}
	return nil
}

// Validate validates the CT_Name and its children
func (_ecaf *CT_Name) Validate() error {
	return _ecaf.ValidateWithPath("\u0043T\u005f\u004e\u0061\u006d\u0065")
}

type LayoutDefHdr struct{ CT_DiagramDefinitionHeader }

func (_dddfa ST_ChildAlignment) Validate() error { return _dddfa.ValidateWithPath("") }

func (_agcfcc ST_DiagramHorizontalAlignment) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_agcfcc.String(), start)
}

func (_eeggb ST_FunctionType) ValidateWithPath(path string) error {
	switch _eeggb {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_eeggb))
	}
	return nil
}

func (_efadc ST_AlgorithmType) String() string {
	switch _efadc {
	case 0:
		return ""
	case 1:
		return "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e"
	case 2:
		return "\u0063\u006f\u006e\u006e"
	case 3:
		return "\u0063\u0079\u0063l\u0065"
	case 4:
		return "\u0068i\u0065\u0072\u0043\u0068\u0069\u006cd"
	case 5:
		return "\u0068\u0069\u0065\u0072\u0052\u006f\u006f\u0074"
	case 6:
		return "\u0070\u0079\u0072\u0061"
	case 7:
		return "\u006c\u0069\u006e"
	case 8:
		return "\u0073\u0070"
	case 9:
		return "\u0074\u0078"
	case 10:
		return "\u0073\u006e\u0061k\u0065"
	}
	return ""
}

const (
	ST_CxnTypeUnset               ST_CxnType = 0
	ST_CxnTypeParOf               ST_CxnType = 1
	ST_CxnTypePresOf              ST_CxnType = 2
	ST_CxnTypePresParOf           ST_CxnType = 3
	ST_CxnTypeUnknownRelationship ST_CxnType = 4
)

// ValidateWithPath validates the CT_PresentationOf and its children, prefixing error messages with path
func (_debfg *CT_PresentationOf) ValidateWithPath(path string) error {
	if _debfg.ExtLst != nil {
		if _bdgg := _debfg.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _bdgg != nil {
			return _bdgg
		}
	}
	return nil
}

func (_cebdg *CT_SDCategory) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0074\u0079\u0070\u0065"}, Value: _f.Sprintf("\u0025\u0076", _cebdg.TypeAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0072\u0069"}, Value: _f.Sprintf("\u0025\u0076", _cebdg.PriAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_gefd *StyleDefHdrLst) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0073\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048d\u0072\u004c\u0073\u0074"
	return _gefd.CT_StyleDefinitionHeaderLst.MarshalXML(e, start)
}

const (
	ST_FallbackDimensionUnset ST_FallbackDimension = 0
	ST_FallbackDimension1D    ST_FallbackDimension = 1
	ST_FallbackDimension2D    ST_FallbackDimension = 2
)

func (_defce ST_AutoTextRotation) Validate() error { return _defce.ValidateWithPath("") }

// Validate validates the CT_RelIds and its children
func (_aaaf *CT_RelIds) Validate() error {
	return _aaaf.ValidateWithPath("\u0043T\u005f\u0052\u0065\u006c\u0049\u0064s")
}

func NewCT_DiagramDefinition() *CT_DiagramDefinition {
	_eacbf := &CT_DiagramDefinition{}
	_eacbf.LayoutNode = NewCT_LayoutNode()
	return _eacbf
}

func (_bebdg ST_ConnectorPoint) String() string {
	switch _bebdg {
	case 0:
		return ""
	case 1:
		return "\u0061\u0075\u0074\u006f"
	case 2:
		return "\u0062\u0043\u0074\u0072"
	case 3:
		return "\u0063\u0074\u0072"
	case 4:
		return "\u006d\u0069\u0064\u004c"
	case 5:
		return "\u006d\u0069\u0064\u0052"
	case 6:
		return "\u0074\u0043\u0074\u0072"
	case 7:
		return "\u0062\u004c"
	case 8:
		return "\u0062\u0052"
	case 9:
		return "\u0074\u004c"
	case 10:
		return "\u0074\u0052"
	case 11:
		return "\u0072\u0061\u0064\u0069\u0061\u006c"
	}
	return ""
}

// Validate validates the StyleDefHdr and its children
func (_daccf *StyleDefHdr) Validate() error {
	return _daccf.ValidateWithPath("S\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072")
}

func NewCT_NumericRule() *CT_NumericRule { _cgea := &CT_NumericRule{}; return _cgea }

func (_df *AG_ConstraintRefAttributes) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _df.RefTypeAttr != ST_ConstraintTypeUnset {
		_ec, _eeg := _df.RefTypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0072e\u0066\u0054\u0079\u0070\u0065"})
		if _eeg != nil {
			return _eeg
		}
		start.Attr = append(start.Attr, _ec)
	}
	if _df.RefForAttr != ST_ConstraintRelationshipUnset {
		_cc, _dc := _df.RefForAttr.MarshalXMLAttr(_a.Name{Local: "\u0072\u0065\u0066\u0046\u006f\u0072"})
		if _dc != nil {
			return _dc
		}
		start.Attr = append(start.Attr, _cc)
	}
	if _df.RefForNameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u0065\u0066\u0046\u006f\u0072\u004e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_df.RefForNameAttr)})
	}
	if _df.RefPtTypeAttr != ST_ElementTypeUnset {
		_gb, _bgb := _df.RefPtTypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0072e\u0066\u0050\u0074\u0054\u0079\u0070e"})
		if _bgb != nil {
			return _bgb
		}
		start.Attr = append(start.Attr, _gb)
	}
	return nil
}

func (_aef *AG_ConstraintRefAttributes) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _aba := range start.Attr {
		if _aba.Name.Local == "\u0072e\u0066\u0054\u0079\u0070\u0065" {
			_aef.RefTypeAttr.UnmarshalXMLAttr(_aba)
			continue
		}
		if _aba.Name.Local == "\u0072\u0065\u0066\u0046\u006f\u0072" {
			_aef.RefForAttr.UnmarshalXMLAttr(_aba)
			continue
		}
		if _aba.Name.Local == "\u0072\u0065\u0066\u0046\u006f\u0072\u004e\u0061\u006d\u0065" {
			_abe, _ed := _aba.Value, error(nil)
			if _ed != nil {
				return _ed
			}
			_aef.RefForNameAttr = &_abe
			continue
		}
		if _aba.Name.Local == "\u0072e\u0066\u0050\u0074\u0054\u0079\u0070e" {
			_aef.RefPtTypeAttr.UnmarshalXMLAttr(_aba)
			continue
		}
	}
	for {
		_ga, _feg := d.Token()
		if _feg != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0041\u0047\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074\u0052\u0065\u0066A\u0074\u0074\u0072\u0069\u0062u\u0074\u0065s\u003a\u0020\u0025\u0073", _feg)
		}
		if _eed, _aad := _ga.(_a.EndElement); _aad && _eed.Name == start.Name {
			break
		}
	}
	return nil
}

func (_gbfd ST_FunctionValue) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _gbfd.Int32 != nil {
		e.EncodeToken(_a.CharData(_f.Sprintf("\u0025\u0064", *_gbfd.Int32)))
	}
	if _gbfd.Bool != nil {
		e.EncodeToken(_a.CharData(_f.Sprintf("\u0025\u0064", _feae(*_gbfd.Bool))))
	}
	if _gbfd.ST_Direction != ST_DirectionUnset {
		e.EncodeToken(_a.CharData(_gbfd.ST_Direction.String()))
	}
	if _gbfd.ST_HierBranchStyle != ST_HierBranchStyleUnset {
		e.EncodeToken(_a.CharData(_gbfd.ST_HierBranchStyle.String()))
	}
	if _gbfd.ST_AnimOneStr != ST_AnimOneStrUnset {
		e.EncodeToken(_a.CharData(_gbfd.ST_AnimOneStr.String()))
	}
	if _gbfd.ST_AnimLvlStr != ST_AnimLvlStrUnset {
		e.EncodeToken(_a.CharData(_gbfd.ST_AnimLvlStr.String()))
	}
	if _gbfd.ST_ResizeHandlesStr != ST_ResizeHandlesStrUnset {
		e.EncodeToken(_a.CharData(_gbfd.ST_ResizeHandlesStr.String()))
	}
	return e.EncodeToken(_a.EndElement{Name: start.Name})
}

func (_feaea ST_GrowDirection) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_eabbbc := _a.Attr{}
	_eabbbc.Name = name
	switch _feaea {
	case ST_GrowDirectionUnset:
		_eabbbc.Value = ""
	case ST_GrowDirectionTL:
		_eabbbc.Value = "\u0074\u004c"
	case ST_GrowDirectionTR:
		_eabbbc.Value = "\u0074\u0052"
	case ST_GrowDirectionBL:
		_eabbbc.Value = "\u0062\u004c"
	case ST_GrowDirectionBR:
		_eabbbc.Value = "\u0062\u0052"
	}
	return _eabbbc, nil
}

func (_aaee ST_CxnType) String() string {
	switch _aaee {
	case 0:
		return ""
	case 1:
		return "\u0070\u0061\u0072O\u0066"
	case 2:
		return "\u0070\u0072\u0065\u0073\u004f\u0066"
	case 3:
		return "\u0070r\u0065\u0073\u0050\u0061\u0072\u004ff"
	case 4:
		return "\u0075\u006e\u006b\u006eow\u006e\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"
	}
	return ""
}

func (_bafge ST_FlowDirection) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_faaef := _a.Attr{}
	_faaef.Name = name
	switch _bafge {
	case ST_FlowDirectionUnset:
		_faaef.Value = ""
	case ST_FlowDirectionRow:
		_faaef.Value = "\u0072\u006f\u0077"
	case ST_FlowDirectionCol:
		_faaef.Value = "\u0063\u006f\u006c"
	}
	return _faaef, nil
}

type ST_HierBranchStyle byte

func NewCT_BulletEnabled() *CT_BulletEnabled { _affc := &CT_BulletEnabled{}; return _affc }

type ST_TextAnchorHorizontal byte

func (_fffc *ST_ModelId) ValidateWithPath(path string) error {
	_cgcaa := []string{}
	if _fffc.Int32 != nil {
		_cgcaa = append(_cgcaa, "\u0049\u006e\u00743\u0032")
	}
	if _fffc.ST_Guid != nil {
		_cgcaa = append(_cgcaa, "\u0053T\u005f\u0047\u0075\u0069\u0064")
	}
	if len(_cgcaa) > 1 {
		return _f.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _cgcaa)
	}
	return nil
}

func (_g *AG_ConstraintAttributes) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _g.TypeAttr != ST_ConstraintTypeUnset {
		_ee, _cf := _g.TypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _cf != nil {
			return _cf
		}
		start.Attr = append(start.Attr, _ee)
	}
	if _g.ForAttr != ST_ConstraintRelationshipUnset {
		_ca, _gg := _g.ForAttr.MarshalXMLAttr(_a.Name{Local: "\u0066\u006f\u0072"})
		if _gg != nil {
			return _gg
		}
		start.Attr = append(start.Attr, _ca)
	}
	if _g.ForNameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0066o\u0072\u004e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_g.ForNameAttr)})
	}
	if _g.PtTypeAttr != ST_ElementTypeUnset {
		_bf, _eb := _g.PtTypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"})
		if _eb != nil {
			return _eb
		}
		start.Attr = append(start.Attr, _bf)
	}
	return nil
}

// ValidateWithPath validates the CT_AdjLst and its children, prefixing error messages with path
func (_be *CT_AdjLst) ValidateWithPath(path string) error {
	for _bae, _abba := range _be.Adj {
		if _eda := _abba.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0041\u0064\u006a\u005b\u0025\u0064\u005d", path, _bae)); _eda != nil {
			return _eda
		}
	}
	return nil
}

func (_bagg *ST_AnimLvlStr) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_bagg = 0
	case "\u006e\u006f\u006e\u0065":
		*_bagg = 1
	case "\u006c\u0076\u006c":
		*_bagg = 2
	case "\u0063\u0074\u0072":
		*_bagg = 3
	}
	return nil
}

func (_cafd ST_ElementType) ValidateWithPath(path string) error {
	switch _cafd {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cafd))
	}
	return nil
}

// ValidateWithPath validates the CT_ForEach and its children, prefixing error messages with path
func (_eace *CT_ForEach) ValidateWithPath(path string) error {
	for _gbe, _gcccg := range _eace.Alg {
		if _dcdc := _gcccg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0041\u006c\u0067\u005b\u0025\u0064\u005d", path, _gbe)); _dcdc != nil {
			return _dcdc
		}
	}
	for _dgdd, _dedg := range _eace.Shape {
		if _gcccb := _dedg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fS\u0068\u0061\u0070\u0065\u005b\u0025\u0064\u005d", path, _dgdd)); _gcccb != nil {
			return _gcccb
		}
	}
	for _eabea, _dgec := range _eace.PresOf {
		if _cgce := _dgec.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0050\u0072\u0065\u0073\u004f\u0066\u005b\u0025\u0064\u005d", path, _eabea)); _cgce != nil {
			return _cgce
		}
	}
	for _efcb, _gbaeb := range _eace.ConstrLst {
		if _bbcd := _gbaeb.ValidateWithPath(_f.Sprintf("\u0025\u0073/\u0043\u006f\u006es\u0074\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _efcb)); _bbcd != nil {
			return _bbcd
		}
	}
	for _ggege, _fcdg := range _eace.RuleLst {
		if _fcfe := _fcdg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0052\u0075\u006c\u0065\u004c\u0073t\u005b\u0025\u0064\u005d", path, _ggege)); _fcfe != nil {
			return _fcfe
		}
	}
	for _bfabca, _cfbb := range _eace.ForEach {
		if _dfd := _cfbb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0046\u006f\u0072\u0045\u0061\u0063h\u005b\u0025\u0064\u005d", path, _bfabca)); _dfd != nil {
			return _dfd
		}
	}
	for _cggg, _cef := range _eace.LayoutNode {
		if _edfea := _cef.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064e\u005b\u0025\u0064\u005d", path, _cggg)); _edfea != nil {
			return _edfea
		}
	}
	for _beff, _ebfcb := range _eace.Choose {
		if _eagd := _ebfcb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u006f\u0073\u0065\u005b\u0025\u0064\u005d", path, _beff)); _eagd != nil {
			return _eagd
		}
	}
	for _fbdd, _ceff := range _eace.ExtLst {
		if _cafc := _ceff.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0045\u0078\u0074\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _fbdd)); _cafc != nil {
			return _cafc
		}
	}
	return nil
}

func (_dbfdc ST_Direction) ValidateWithPath(path string) error {
	switch _dbfdc {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dbfdc))
	}
	return nil
}

func (_geccg ST_ElementType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_gafg := _a.Attr{}
	_gafg.Name = name
	switch _geccg {
	case ST_ElementTypeUnset:
		_gafg.Value = ""
	case ST_ElementTypeAll:
		_gafg.Value = "\u0061\u006c\u006c"
	case ST_ElementTypeDoc:
		_gafg.Value = "\u0064\u006f\u0063"
	case ST_ElementTypeNode:
		_gafg.Value = "\u006e\u006f\u0064\u0065"
	case ST_ElementTypeNorm:
		_gafg.Value = "\u006e\u006f\u0072\u006d"
	case ST_ElementTypeNonNorm:
		_gafg.Value = "\u006eo\u006e\u004e\u006f\u0072\u006d"
	case ST_ElementTypeAsst:
		_gafg.Value = "\u0061\u0073\u0073\u0074"
	case ST_ElementTypeNonAsst:
		_gafg.Value = "\u006eo\u006e\u0041\u0073\u0073\u0074"
	case ST_ElementTypeParTrans:
		_gafg.Value = "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073"
	case ST_ElementTypePres:
		_gafg.Value = "\u0070\u0072\u0065\u0073"
	case ST_ElementTypeSibTrans:
		_gafg.Value = "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073"
	}
	return _gafg, nil
}

func (_abafa ST_SecondaryLinearDirection) Validate() error { return _abafa.ValidateWithPath("") }

func (_fgbba ST_TextAnchorVertical) ValidateWithPath(path string) error {
	switch _fgbba {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fgbba))
	}
	return nil
}

func (_cdgee ST_VariableType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_cdgee.String(), start)
}

type ST_ClrAppMethod byte

func NewCT_Pt() *CT_Pt { _fcbbb := &CT_Pt{}; return _fcbbb }

type CT_OrgChart struct{ ValAttr *bool }

func (_cgdc *CT_Category) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0074\u0079\u0070\u0065"}, Value: _f.Sprintf("\u0025\u0076", _cgdc.TypeAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0072\u0069"}, Value: _f.Sprintf("\u0025\u0076", _cgdc.PriAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_ceda *CT_StyleDefinitionHeader) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", _ceda.UniqueIdAttr)})
	if _ceda.MinVerAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _f.Sprintf("\u0025\u0076", *_ceda.MinVerAttr)})
	}
	if _ceda.ResIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u0065\u0073I\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_ceda.ResIdAttr)})
	}
	e.EncodeToken(start)
	_dafec := _a.StartElement{Name: _a.Name{Local: "\u0074\u0069\u0074l\u0065"}}
	for _, _bcabd := range _ceda.Title {
		e.EncodeElement(_bcabd, _dafec)
	}
	_babcg := _a.StartElement{Name: _a.Name{Local: "\u0064\u0065\u0073\u0063"}}
	for _, _cbbfa := range _ceda.Desc {
		e.EncodeElement(_cbbfa, _babcg)
	}
	if _ceda.CatLst != nil {
		_gcgacc := _a.StartElement{Name: _a.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_ceda.CatLst, _gcgacc)
	}
	if _ceda.ExtLst != nil {
		_edcef := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_ceda.ExtLst, _edcef)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_ddcfe *ST_CenterShapeMapping) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ddcfe = 0
	case "\u006e\u006f\u006e\u0065":
		*_ddcfe = 1
	case "\u0066\u004e\u006fd\u0065":
		*_ddcfe = 2
	}
	return nil
}

func (_aeabc ST_ConstraintRelationship) ValidateWithPath(path string) error {
	switch _aeabc {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_aeabc))
	}
	return nil
}

func NewAG_ConstraintRefAttributes() *AG_ConstraintRefAttributes {
	_ebc := &AG_ConstraintRefAttributes{}
	return _ebc
}

func (_egeb ST_ClrAppMethod) Validate() error { return _egeb.ValidateWithPath("") }

func (_bfdd ST_BendPoint) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_afaed := _a.Attr{}
	_afaed.Name = name
	switch _bfdd {
	case ST_BendPointUnset:
		_afaed.Value = ""
	case ST_BendPointBeg:
		_afaed.Value = "\u0062\u0065\u0067"
	case ST_BendPointDef:
		_afaed.Value = "\u0064\u0065\u0066"
	case ST_BendPointEnd:
		_afaed.Value = "\u0065\u006e\u0064"
	}
	return _afaed, nil
}

// ValidateWithPath validates the CT_Adj and its children, prefixing error messages with path
func (_cg *CT_Adj) ValidateWithPath(path string) error {
	if _cg.IdxAttr < 1 {
		return _f.Errorf("%\u0073\u002f\u006d\u002e\u0049\u0064x\u0041\u0074\u0074\u0072\u0020\u006du\u0073\u0074\u0020\u0062\u0065\u0020\u003e=\u0020\u0031\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, _cg.IdxAttr)
	}
	return nil
}

func (_ggbf *CT_ChildPref) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _ggcd := range start.Attr {
		if _ggcd.Name.Local == "\u0076\u0061\u006c" {
			_adgf, _bdgf := _ae.ParseInt(_ggcd.Value, 10, 32)
			if _bdgf != nil {
				return _bdgf
			}
			_afaa := int32(_adgf)
			_ggbf.ValAttr = &_afaa
			continue
		}
	}
	for {
		_faga, _ffeg := d.Token()
		if _ffeg != nil {
			return _f.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005f\u0043\u0068i\u006c\u0064\u0050\u0072\u0065\u0066\u003a\u0020\u0025\u0073", _ffeg)
		}
		if _agfg, _eeb := _faga.(_a.EndElement); _eeb && _agfg.Name == start.Name {
			break
		}
	}
	return nil
}

func NewDataModel() *DataModel {
	_afcf := &DataModel{}
	_afcf.CT_DataModel = *NewCT_DataModel()
	return _afcf
}

func (_beae ST_AlgorithmType) ValidateWithPath(path string) error {
	switch _beae {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_beae))
	}
	return nil
}

func (_becee ST_PtType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_becee.String(), start)
}

func (_cafbbd ST_DiagramHorizontalAlignment) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_bage := _a.Attr{}
	_bage.Name = name
	switch _cafbbd {
	case ST_DiagramHorizontalAlignmentUnset:
		_bage.Value = ""
	case ST_DiagramHorizontalAlignmentL:
		_bage.Value = "\u006c"
	case ST_DiagramHorizontalAlignmentCtr:
		_bage.Value = "\u0063\u0074\u0072"
	case ST_DiagramHorizontalAlignmentR:
		_bage.Value = "\u0072"
	case ST_DiagramHorizontalAlignmentNone:
		_bage.Value = "\u006e\u006f\u006e\u0065"
	}
	return _bage, nil
}

func NewCT_ColorTransformHeader() *CT_ColorTransformHeader {
	_fgg := &CT_ColorTransformHeader{}
	return _fgg
}

func (_feeff ST_CenterShapeMapping) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_feeff.String(), start)
}

func (_gbdb ST_ResizeHandlesStr) String() string {
	switch _gbdb {
	case 0:
		return ""
	case 1:
		return "\u0065\u0078\u0061c\u0074"
	case 2:
		return "\u0072\u0065\u006c"
	}
	return ""
}

const (
	ST_ArrowheadStyleUnset ST_ArrowheadStyle = 0
	ST_ArrowheadStyleAuto  ST_ArrowheadStyle = 1
	ST_ArrowheadStyleArr   ST_ArrowheadStyle = 2
	ST_ArrowheadStyleNoArr ST_ArrowheadStyle = 3
)

func (_deef ST_Direction) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_deef.String(), start)
}

// Validate validates the CT_StyleLabel and its children
func (_ebad *CT_StyleLabel) Validate() error {
	return _ebad.ValidateWithPath("\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c")
}

// Validate validates the CT_CTStyleLabel and its children
func (_dgbb *CT_CTStyleLabel) Validate() error {
	return _dgbb.ValidateWithPath("\u0043T\u005fC\u0054\u0053\u0074\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c")
}

func (_ffec ST_ChildDirection) ValidateWithPath(path string) error {
	switch _ffec {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ffec))
	}
	return nil
}

func (_fbea ST_VerticalAlignment) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_fbea.String(), start)
}

func (_fgaf *CT_StyleDefinition) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _fgaf.UniqueIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_fgaf.UniqueIdAttr)})
	}
	if _fgaf.MinVerAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _f.Sprintf("\u0025\u0076", *_fgaf.MinVerAttr)})
	}
	e.EncodeToken(start)
	if _fgaf.Title != nil {
		_fgda := _a.StartElement{Name: _a.Name{Local: "\u0074\u0069\u0074l\u0065"}}
		for _, _fdegc := range _fgaf.Title {
			e.EncodeElement(_fdegc, _fgda)
		}
	}
	if _fgaf.Desc != nil {
		_fccb := _a.StartElement{Name: _a.Name{Local: "\u0064\u0065\u0073\u0063"}}
		for _, _edea := range _fgaf.Desc {
			e.EncodeElement(_edea, _fccb)
		}
	}
	if _fgaf.CatLst != nil {
		_cefd := _a.StartElement{Name: _a.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_fgaf.CatLst, _cefd)
	}
	if _fgaf.Scene3d != nil {
		_gfeee := _a.StartElement{Name: _a.Name{Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}}
		e.EncodeElement(_fgaf.Scene3d, _gfeee)
	}
	_bfabg := _a.StartElement{Name: _a.Name{Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}}
	for _, _baec := range _fgaf.StyleLbl {
		e.EncodeElement(_baec, _bfabg)
	}
	if _fgaf.ExtLst != nil {
		_feea := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_fgaf.ExtLst, _feea)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_bcbg ST_FunctionOperator) Validate() error { return _bcbg.ValidateWithPath("") }

type ST_FunctionOperator byte

func NewCT_StyleLabel() *CT_StyleLabel { _ggfc := &CT_StyleLabel{}; return _ggfc }

func (_agbeg ST_StartingElement) String() string {
	switch _agbeg {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u0064\u0065"
	case 2:
		return "\u0074\u0072\u0061n\u0073"
	}
	return ""
}

func (_aebe *CT_ForEach) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _gbfef := range start.Attr {
		if _gbfef.Name.Local == "\u0072\u0065\u0066" {
			_cea, _aaec := _gbfef.Value, error(nil)
			if _aaec != nil {
				return _aaec
			}
			_aebe.RefAttr = &_cea
			continue
		}
		if _gbfef.Name.Local == "\u006e\u0061\u006d\u0065" {
			_gdaf, _cbga := _gbfef.Value, error(nil)
			if _cbga != nil {
				return _cbga
			}
			_aebe.NameAttr = &_gdaf
			continue
		}
		if _gbfef.Name.Local == "\u0061\u0078\u0069\u0073" {
			_gedf, _gfdg := ParseSliceST_AxisTypes(_gbfef.Value)
			if _gfdg != nil {
				return _gfdg
			}
			_aebe.AxisAttr = &_gedf
			continue
		}
		if _gbfef.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_facc, _fgfd := ParseSliceST_ElementTypes(_gbfef.Value)
			if _fgfd != nil {
				return _fgfd
			}
			_aebe.PtTypeAttr = &_facc
			continue
		}
		if _gbfef.Name.Local == "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073" {
			_bbdd, _dbfd := ParseSliceST_Booleans(_gbfef.Value)
			if _dbfd != nil {
				return _dbfd
			}
			_aebe.HideLastTransAttr = &_bbdd
			continue
		}
		if _gbfef.Name.Local == "\u0073\u0074" {
			_ccge, _ffeea := ParseSliceST_Ints(_gbfef.Value)
			if _ffeea != nil {
				return _ffeea
			}
			_aebe.StAttr = &_ccge
			continue
		}
		if _gbfef.Name.Local == "\u0063\u006e\u0074" {
			_dbge, _aabd := ParseSliceST_UnsignedInts(_gbfef.Value)
			if _aabd != nil {
				return _aabd
			}
			_aebe.CntAttr = &_dbge
			continue
		}
		if _gbfef.Name.Local == "\u0073\u0074\u0065\u0070" {
			_ceaf, _fabdg := ParseSliceST_Ints(_gbfef.Value)
			if _fabdg != nil {
				return _fabdg
			}
			_aebe.StepAttr = &_ceaf
			continue
		}
	}
_fcag:
	for {
		_fcacg, _dacc := d.Token()
		if _dacc != nil {
			return _dacc
		}
		switch _afda := _fcacg.(type) {
		case _a.StartElement:
			switch _afda.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u006c\u0067"}:
				_eaab := NewCT_Algorithm()
				if _ebgef := d.DecodeElement(_eaab, &_afda); _ebgef != nil {
					return _ebgef
				}
				_aebe.Alg = append(_aebe.Alg, _eaab)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0068\u0061p\u0065"}:
				_cgcf := NewCT_Shape()
				if _aaga := d.DecodeElement(_cgcf, &_afda); _aaga != nil {
					return _aaga
				}
				_aebe.Shape = append(_aebe.Shape, _cgcf)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}:
				_cbbd := NewCT_PresentationOf()
				if _adca := d.DecodeElement(_cbbd, &_afda); _adca != nil {
					return _adca
				}
				_aebe.PresOf = append(_aebe.PresOf, _cbbd)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}:
				_ebag := NewCT_Constraints()
				if _cccff := d.DecodeElement(_ebag, &_afda); _cccff != nil {
					return _cccff
				}
				_aebe.ConstrLst = append(_aebe.ConstrLst, _ebag)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}:
				_bbfe := NewCT_Rules()
				if _bagd := d.DecodeElement(_bbfe, &_afda); _bagd != nil {
					return _bagd
				}
				_aebe.RuleLst = append(_aebe.RuleLst, _bbfe)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}:
				_caab := NewCT_ForEach()
				if _cfcd := d.DecodeElement(_caab, &_afda); _cfcd != nil {
					return _cfcd
				}
				_aebe.ForEach = append(_aebe.ForEach, _caab)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				_gfee := NewCT_LayoutNode()
				if _eddb := d.DecodeElement(_gfee, &_afda); _eddb != nil {
					return _eddb
				}
				_aebe.LayoutNode = append(_aebe.LayoutNode, _gfee)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}:
				_eabe := NewCT_Choose()
				if _fdf := d.DecodeElement(_eabe, &_afda); _fdf != nil {
					return _fdf
				}
				_aebe.Choose = append(_aebe.Choose, _eabe)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_ebggf := _bg.NewCT_OfficeArtExtensionList()
				if _cfcea := d.DecodeElement(_ebggf, &_afda); _cfcea != nil {
					return _cfcea
				}
				_aebe.ExtLst = append(_aebe.ExtLst, _ebggf)
			default:
				_b.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fF\u006f\u0072\u0045\u0061\u0063\u0068\u0020\u0025\u0076", _afda.Name)
				if _ffadg := d.Skip(); _ffadg != nil {
					return _ffadg
				}
			}
		case _a.EndElement:
			break _fcag
		case _a.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_ResizeHandles and its children, prefixing error messages with path
func (_gefc *CT_ResizeHandles) ValidateWithPath(path string) error {
	if _gggfa := _gefc.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _gggfa != nil {
		return _gggfa
	}
	return nil
}

func (_cdf *CT_CTCategory) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0074\u0079\u0070\u0065"}, Value: _f.Sprintf("\u0025\u0076", _cdf.TypeAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0072\u0069"}, Value: _f.Sprintf("\u0025\u0076", _cdf.PriAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_Constraints and its children, prefixing error messages with path
func (_afde *CT_Constraints) ValidateWithPath(path string) error {
	for _bbgf, _ggee := range _afde.Constr {
		if _aab := _ggee.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u006f\u006e\u0073\u0074\u0072\u005b\u0025\u0064\u005d", path, _bbgf)); _aab != nil {
			return _aab
		}
	}
	return nil
}

func (_gfagb *ST_TextBlockDirection) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_gfagb = 0
	case "\u0068\u006f\u0072\u007a":
		*_gfagb = 1
	case "\u0076\u0065\u0072\u0074":
		*_gfagb = 2
	}
	return nil
}

// Validate validates the CT_DiagramDefinitionHeader and its children
func (_cecg *CT_DiagramDefinitionHeader) Validate() error {
	return _cecg.ValidateWithPath("\u0043\u0054\u005f\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065a\u0064\u0065\u0072")
}

func (_gdgcf ST_AlgorithmType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_ggeb := _a.Attr{}
	_ggeb.Name = name
	switch _gdgcf {
	case ST_AlgorithmTypeUnset:
		_ggeb.Value = ""
	case ST_AlgorithmTypeComposite:
		_ggeb.Value = "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e"
	case ST_AlgorithmTypeConn:
		_ggeb.Value = "\u0063\u006f\u006e\u006e"
	case ST_AlgorithmTypeCycle:
		_ggeb.Value = "\u0063\u0079\u0063l\u0065"
	case ST_AlgorithmTypeHierChild:
		_ggeb.Value = "\u0068i\u0065\u0072\u0043\u0068\u0069\u006cd"
	case ST_AlgorithmTypeHierRoot:
		_ggeb.Value = "\u0068\u0069\u0065\u0072\u0052\u006f\u006f\u0074"
	case ST_AlgorithmTypePyra:
		_ggeb.Value = "\u0070\u0079\u0072\u0061"
	case ST_AlgorithmTypeLin:
		_ggeb.Value = "\u006c\u0069\u006e"
	case ST_AlgorithmTypeSp:
		_ggeb.Value = "\u0073\u0070"
	case ST_AlgorithmTypeTx:
		_ggeb.Value = "\u0074\u0078"
	case ST_AlgorithmTypeSnake:
		_ggeb.Value = "\u0073\u006e\u0061k\u0065"
	}
	return _ggeb, nil
}

func (_ecgd *CT_Parameter) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	_fbfga, _acdgd := _ecgd.TypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0074\u0079\u0070\u0065"})
	if _acdgd != nil {
		return _acdgd
	}
	start.Attr = append(start.Attr, _fbfga)
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", _ecgd.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_StyleDefinitionHeaderLst and its children, prefixing error messages with path
func (_bbca *CT_StyleDefinitionHeaderLst) ValidateWithPath(path string) error {
	for _afcg, _agae := range _bbca.StyleDefHdr {
		if _bgcf := _agae.ValidateWithPath(_f.Sprintf("\u0025s\u002fS\u0074\u0079\u006c\u0065\u0044e\u0066\u0048d\u0072\u005b\u0025\u0064\u005d", path, _afcg)); _bgcf != nil {
			return _bgcf
		}
	}
	return nil
}

func (_dfbga ST_TextAnchorVertical) String() string {
	switch _dfbga {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u006d\u0069\u0064"
	case 3:
		return "\u0062"
	}
	return ""
}

func (_fgdab ST_ChildOrderType) ValidateWithPath(path string) error {
	switch _fgdab {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fgdab))
	}
	return nil
}

func (_fgfbg *StyleDefHdrLst) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_fgfbg.CT_StyleDefinitionHeaderLst = *NewCT_StyleDefinitionHeaderLst()
_gacb:
	for {
		_eccfg, _feba := d.Token()
		if _feba != nil {
			return _feba
		}
		switch _gfdd := _eccfg.(type) {
		case _a.StartElement:
			switch _gfdd.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_egdda := NewCT_StyleDefinitionHeader()
				if _cfgf := d.DecodeElement(_egdda, &_gfdd); _cfgf != nil {
					return _cfgf
				}
				_fgfbg.StyleDefHdr = append(_fgfbg.StyleDefHdr, _egdda)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064r\u004cs\u0074\u0020\u0025\u0076", _gfdd.Name)
				if _gffb := d.Skip(); _gffb != nil {
					return _gffb
				}
			}
		case _a.EndElement:
			break _gacb
		case _a.CharData:
		}
	}
	return nil
}

func (_bgedf ST_BendPoint) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_bgedf.String(), start)
}

type ST_ArrowheadStyle byte

func (_adbe ST_ClrAppMethod) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_egfeg := _a.Attr{}
	_egfeg.Name = name
	switch _adbe {
	case ST_ClrAppMethodUnset:
		_egfeg.Value = ""
	case ST_ClrAppMethodSpan:
		_egfeg.Value = "\u0073\u0070\u0061\u006e"
	case ST_ClrAppMethodCycle:
		_egfeg.Value = "\u0063\u0079\u0063l\u0065"
	case ST_ClrAppMethodRepeat:
		_egfeg.Value = "\u0072\u0065\u0070\u0065\u0061\u0074"
	}
	return _egfeg, nil
}

func (_gacab ST_Offset) String() string {
	switch _gacab {
	case 0:
		return ""
	case 1:
		return "\u0063\u0074\u0072"
	case 2:
		return "\u006f\u0066\u0066"
	}
	return ""
}

func NewCT_OrgChart() *CT_OrgChart { _cddaf := &CT_OrgChart{}; return _cddaf }

func NewStyleDef() *StyleDef {
	_dgfa := &StyleDef{}
	_dgfa.CT_StyleDefinition = *NewCT_StyleDefinition()
	return _dgfa
}

// ValidateWithPath validates the ColorsDefHdrLst and its children, prefixing error messages with path
func (_dfbc *ColorsDefHdrLst) ValidateWithPath(path string) error {
	if _cfac := _dfbc.CT_ColorTransformHeaderLst.ValidateWithPath(path); _cfac != nil {
		return _cfac
	}
	return nil
}

// ValidateWithPath validates the CT_Direction and its children, prefixing error messages with path
func (_gggg *CT_Direction) ValidateWithPath(path string) error {
	if _bfabc := _gggg.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _bfabc != nil {
		return _bfabc
	}
	return nil
}

// Validate validates the CT_BulletEnabled and its children
func (_adba *CT_BulletEnabled) Validate() error {
	return _adba.ValidateWithPath("\u0043\u0054_\u0042\u0075\u006cl\u0065\u0074\u0045\u006e\u0061\u0062\u006c\u0065\u0064")
}

func (_ebbda ST_HueDir) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_ebbda.String(), start)
}

type ST_AnimOneStr byte

func (_accc *ST_FunctionArgument) ValidateWithPath(path string) error {
	_aefb := []string{}
	if _accc.ST_VariableType != ST_VariableTypeUnset {
		_aefb = append(_aefb, "\u0053T\u005fV\u0061\u0072\u0069\u0061\u0062\u006c\u0065\u0054\u0079\u0070\u0065")
	}
	if len(_aefb) > 1 {
		return _f.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _aefb)
	}
	return nil
}

type CT_CTDescription struct {
	LangAttr *string
	ValAttr  string
}

func (_fbec *ST_HierarchyAlignment) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_fbec = 0
	case "\u0074\u004c":
		*_fbec = 1
	case "\u0074\u0052":
		*_fbec = 2
	case "\u0074\u0043\u0074\u0072\u0043\u0068":
		*_fbec = 3
	case "\u0074C\u0074\u0072\u0044\u0065\u0073":
		*_fbec = 4
	case "\u0062\u004c":
		*_fbec = 5
	case "\u0062\u0052":
		*_fbec = 6
	case "\u0062\u0043\u0074\u0072\u0043\u0068":
		*_fbec = 7
	case "\u0062C\u0074\u0072\u0044\u0065\u0073":
		*_fbec = 8
	case "\u006c\u0054":
		*_fbec = 9
	case "\u006c\u0042":
		*_fbec = 10
	case "\u006c\u0043\u0074\u0072\u0043\u0068":
		*_fbec = 11
	case "\u006cC\u0074\u0072\u0044\u0065\u0073":
		*_fbec = 12
	case "\u0072\u0054":
		*_fbec = 13
	case "\u0072\u0042":
		*_fbec = 14
	case "\u0072\u0043\u0074\u0072\u0043\u0068":
		*_fbec = 15
	case "\u0072C\u0074\u0072\u0044\u0065\u0073":
		*_fbec = 16
	}
	return nil
}

const (
	ST_AxisTypeUnset       ST_AxisType = 0
	ST_AxisTypeSelf        ST_AxisType = 1
	ST_AxisTypeCh          ST_AxisType = 2
	ST_AxisTypeDes         ST_AxisType = 3
	ST_AxisTypeDesOrSelf   ST_AxisType = 4
	ST_AxisTypePar         ST_AxisType = 5
	ST_AxisTypeAncst       ST_AxisType = 6
	ST_AxisTypeAncstOrSelf ST_AxisType = 7
	ST_AxisTypeFollowSib   ST_AxisType = 8
	ST_AxisTypePrecedSib   ST_AxisType = 9
	ST_AxisTypeFollow      ST_AxisType = 10
	ST_AxisTypePreced      ST_AxisType = 11
	ST_AxisTypeRoot        ST_AxisType = 12
	ST_AxisTypeNone        ST_AxisType = 13
)

func (_ffeba *ST_SecondaryChildAlignment) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ebcdd, _ggabf := d.Token()
	if _ggabf != nil {
		return _ggabf
	}
	if _cgegg, _dfgeg := _ebcdd.(_a.EndElement); _dfgeg && _cgegg.Name == start.Name {
		*_ffeba = 1
		return nil
	}
	if _acda, _eagba := _ebcdd.(_a.CharData); !_eagba {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ebcdd)
	} else {
		switch string(_acda) {
		case "":
			*_ffeba = 0
		case "\u006e\u006f\u006e\u0065":
			*_ffeba = 1
		case "\u0074":
			*_ffeba = 2
		case "\u0062":
			*_ffeba = 3
		case "\u006c":
			*_ffeba = 4
		case "\u0072":
			*_ffeba = 5
		}
	}
	_ebcdd, _ggabf = d.Token()
	if _ggabf != nil {
		return _ggabf
	}
	if _ggggb, _cabga := _ebcdd.(_a.EndElement); _cabga && _ggggb.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ebcdd)
}

func (_fggec ST_GrowDirection) ValidateWithPath(path string) error {
	switch _fggec {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fggec))
	}
	return nil
}

func (_eeeba ST_TextDirection) Validate() error { return _eeeba.ValidateWithPath("") }

func (_cabd ST_ChildOrderType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_cabd.String(), start)
}

func (_bbbc ST_TextBlockDirection) Validate() error { return _bbbc.ValidateWithPath("") }

func (_dfgd *DataModel) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_dfgd.CT_DataModel = *NewCT_DataModel()
_efba:
	for {
		_eaad, _bead := d.Token()
		if _bead != nil {
			return _bead
		}
		switch _gegbbd := _eaad.(type) {
		case _a.StartElement:
			switch _gegbbd.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0074\u004cs\u0074"}:
				if _eggf := d.DecodeElement(_dfgd.PtLst, &_gegbbd); _eggf != nil {
					return _eggf
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0078\u006e\u004c\u0073\u0074"}:
				_dfgd.CxnLst = NewCT_CxnList()
				if _edcc := d.DecodeElement(_dfgd.CxnLst, &_gegbbd); _edcc != nil {
					return _edcc
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0062\u0067"}:
				_dfgd.Bg = _bg.NewCT_BackgroundFormatting()
				if _cgbf := d.DecodeElement(_dfgd.Bg, &_gegbbd); _cgbf != nil {
					return _cgbf
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0077\u0068\u006fl\u0065"}:
				_dfgd.Whole = _bg.NewCT_WholeE2oFormatting()
				if _affa := d.DecodeElement(_dfgd.Whole, &_gegbbd); _affa != nil {
					return _affa
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_dfgd.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _ffda := d.DecodeElement(_dfgd.ExtLst, &_gegbbd); _ffda != nil {
					return _ffda
				}
			default:
				_b.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0044\u0061\u0074\u0061\u004d\u006fd\u0065\u006c \u0025\u0076", _gegbbd.Name)
				if _gdeff := d.Skip(); _gdeff != nil {
					return _gdeff
				}
			}
		case _a.EndElement:
			break _efba
		case _a.CharData:
		}
	}
	return nil
}

func (_dgggf ST_DiagramHorizontalAlignment) Validate() error { return _dgggf.ValidateWithPath("") }

func (_fccfg *CT_LayoutNode) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _addg := range start.Attr {
		if _addg.Name.Local == "\u006e\u0061\u006d\u0065" {
			_egaa, _bbcg := _addg.Value, error(nil)
			if _bbcg != nil {
				return _bbcg
			}
			_fccfg.NameAttr = &_egaa
			continue
		}
		if _addg.Name.Local == "\u0063h\u004f\u0072\u0064\u0065\u0072" {
			_fccfg.ChOrderAttr.UnmarshalXMLAttr(_addg)
			continue
		}
		if _addg.Name.Local == "\u006d\u006f\u0076\u0065\u0057\u0069\u0074\u0068" {
			_cgcg, _gfedb := _addg.Value, error(nil)
			if _gfedb != nil {
				return _gfedb
			}
			_fccfg.MoveWithAttr = &_cgcg
			continue
		}
		if _addg.Name.Local == "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c" {
			_ebdf, _bffe := _addg.Value, error(nil)
			if _bffe != nil {
				return _bffe
			}
			_fccfg.StyleLblAttr = &_ebdf
			continue
		}
	}
_cdabe:
	for {
		_degg, _gfaf := d.Token()
		if _gfaf != nil {
			return _gfaf
		}
		switch _cadb := _degg.(type) {
		case _a.StartElement:
			switch _cadb.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u006c\u0067"}:
				_dbab := NewCT_Algorithm()
				if _aagd := d.DecodeElement(_dbab, &_cadb); _aagd != nil {
					return _aagd
				}
				_fccfg.Alg = append(_fccfg.Alg, _dbab)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0068\u0061p\u0065"}:
				_fec := NewCT_Shape()
				if _abgd := d.DecodeElement(_fec, &_cadb); _abgd != nil {
					return _abgd
				}
				_fccfg.Shape = append(_fccfg.Shape, _fec)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}:
				_feag := NewCT_PresentationOf()
				if _badc := d.DecodeElement(_feag, &_cadb); _badc != nil {
					return _badc
				}
				_fccfg.PresOf = append(_fccfg.PresOf, _feag)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}:
				_agacc := NewCT_Constraints()
				if _dada := d.DecodeElement(_agacc, &_cadb); _dada != nil {
					return _dada
				}
				_fccfg.ConstrLst = append(_fccfg.ConstrLst, _agacc)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}:
				_ede := NewCT_Rules()
				if _bfff := d.DecodeElement(_ede, &_cadb); _bfff != nil {
					return _bfff
				}
				_fccfg.RuleLst = append(_fccfg.RuleLst, _ede)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0076\u0061\u0072\u004c\u0073\u0074"}:
				_bgabc := NewCT_LayoutVariablePropertySet()
				if _gaea := d.DecodeElement(_bgabc, &_cadb); _gaea != nil {
					return _gaea
				}
				_fccfg.VarLst = append(_fccfg.VarLst, _bgabc)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}:
				_efffd := NewCT_ForEach()
				if _egbad := d.DecodeElement(_efffd, &_cadb); _egbad != nil {
					return _egbad
				}
				_fccfg.ForEach = append(_fccfg.ForEach, _efffd)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				_babe := NewCT_LayoutNode()
				if _bebf := d.DecodeElement(_babe, &_cadb); _bebf != nil {
					return _bebf
				}
				_fccfg.LayoutNode = append(_fccfg.LayoutNode, _babe)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}:
				_cfbc := NewCT_Choose()
				if _baab := d.DecodeElement(_cfbc, &_cadb); _baab != nil {
					return _baab
				}
				_fccfg.Choose = append(_fccfg.Choose, _cfbc)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_bffd := _bg.NewCT_OfficeArtExtensionList()
				if _fcbb := d.DecodeElement(_bffd, &_cadb); _fcbb != nil {
					return _fcbb
				}
				_fccfg.ExtLst = append(_fccfg.ExtLst, _bffd)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004ca\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065 \u0025\u0076", _cadb.Name)
				if _agebb := d.Skip(); _agebb != nil {
					return _agebb
				}
			}
		case _a.EndElement:
			break _cdabe
		case _a.CharData:
		}
	}
	return nil
}

type CT_ChildPref struct{ ValAttr *int32 }

type RelIds struct{ CT_RelIds }

func (_bbga *ST_StartingElement) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_bbga = 0
	case "\u006e\u006f\u0064\u0065":
		*_bbga = 1
	case "\u0074\u0072\u0061n\u0073":
		*_bbga = 2
	}
	return nil
}

func (_gdbd *ST_ConstraintType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_gdbd = 0
	case "\u006e\u006f\u006e\u0065":
		*_gdbd = 1
	case "\u0061\u006c\u0069\u0067\u006e\u004f\u0066\u0066":
		*_gdbd = 2
	case "\u0062e\u0067\u004d\u0061\u0072\u0067":
		*_gdbd = 3
	case "\u0062\u0065\u006e\u0064\u0044\u0069\u0073\u0074":
		*_gdbd = 4
	case "\u0062\u0065\u0067\u0050\u0061\u0064":
		*_gdbd = 5
	case "\u0062":
		*_gdbd = 6
	case "\u0062\u004d\u0061r\u0067":
		*_gdbd = 7
	case "\u0062\u004f\u0066\u0066":
		*_gdbd = 8
	case "\u0063\u0074\u0072\u0058":
		*_gdbd = 9
	case "\u0063t\u0072\u0058\u004f\u0066\u0066":
		*_gdbd = 10
	case "\u0063\u0074\u0072\u0059":
		*_gdbd = 11
	case "\u0063t\u0072\u0059\u004f\u0066\u0066":
		*_gdbd = 12
	case "\u0063\u006f\u006e\u006e\u0044\u0069\u0073\u0074":
		*_gdbd = 13
	case "\u0064\u0069\u0061\u006d":
		*_gdbd = 14
	case "\u0065n\u0064\u004d\u0061\u0072\u0067":
		*_gdbd = 15
	case "\u0065\u006e\u0064\u0050\u0061\u0064":
		*_gdbd = 16
	case "\u0068":
		*_gdbd = 17
	case "\u0068\u0041\u0072\u0048":
		*_gdbd = 18
	case "\u0068\u004f\u0066\u0066":
		*_gdbd = 19
	case "\u006c":
		*_gdbd = 20
	case "\u006c\u004d\u0061r\u0067":
		*_gdbd = 21
	case "\u006c\u004f\u0066\u0066":
		*_gdbd = 22
	case "\u0072":
		*_gdbd = 23
	case "\u0072\u004d\u0061r\u0067":
		*_gdbd = 24
	case "\u0072\u004f\u0066\u0066":
		*_gdbd = 25
	case "\u0070\u0072\u0069\u006d\u0046\u006f\u006e\u0074\u0053\u007a":
		*_gdbd = 26
	case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0052\u0061\u0074\u0069\u006f":
		*_gdbd = 27
	case "\u0073e\u0063\u0046\u006f\u006e\u0074\u0053z":
		*_gdbd = 28
	case "\u0073\u0069\u0062S\u0070":
		*_gdbd = 29
	case "\u0073\u0065\u0063\u0053\u0069\u0062\u0053\u0070":
		*_gdbd = 30
	case "\u0073\u0070":
		*_gdbd = 31
	case "\u0073t\u0065\u006d\u0054\u0068\u0069\u0063k":
		*_gdbd = 32
	case "\u0074":
		*_gdbd = 33
	case "\u0074\u004d\u0061r\u0067":
		*_gdbd = 34
	case "\u0074\u004f\u0066\u0066":
		*_gdbd = 35
	case "\u0075\u0073\u0065r\u0041":
		*_gdbd = 36
	case "\u0075\u0073\u0065r\u0042":
		*_gdbd = 37
	case "\u0075\u0073\u0065r\u0043":
		*_gdbd = 38
	case "\u0075\u0073\u0065r\u0044":
		*_gdbd = 39
	case "\u0075\u0073\u0065r\u0045":
		*_gdbd = 40
	case "\u0075\u0073\u0065r\u0046":
		*_gdbd = 41
	case "\u0075\u0073\u0065r\u0047":
		*_gdbd = 42
	case "\u0075\u0073\u0065r\u0048":
		*_gdbd = 43
	case "\u0075\u0073\u0065r\u0049":
		*_gdbd = 44
	case "\u0075\u0073\u0065r\u004a":
		*_gdbd = 45
	case "\u0075\u0073\u0065r\u004b":
		*_gdbd = 46
	case "\u0075\u0073\u0065r\u004c":
		*_gdbd = 47
	case "\u0075\u0073\u0065r\u004d":
		*_gdbd = 48
	case "\u0075\u0073\u0065r\u004e":
		*_gdbd = 49
	case "\u0075\u0073\u0065r\u004f":
		*_gdbd = 50
	case "\u0075\u0073\u0065r\u0050":
		*_gdbd = 51
	case "\u0075\u0073\u0065r\u0051":
		*_gdbd = 52
	case "\u0075\u0073\u0065r\u0052":
		*_gdbd = 53
	case "\u0075\u0073\u0065r\u0053":
		*_gdbd = 54
	case "\u0075\u0073\u0065r\u0054":
		*_gdbd = 55
	case "\u0075\u0073\u0065r\u0055":
		*_gdbd = 56
	case "\u0075\u0073\u0065r\u0056":
		*_gdbd = 57
	case "\u0075\u0073\u0065r\u0057":
		*_gdbd = 58
	case "\u0075\u0073\u0065r\u0058":
		*_gdbd = 59
	case "\u0075\u0073\u0065r\u0059":
		*_gdbd = 60
	case "\u0075\u0073\u0065r\u005a":
		*_gdbd = 61
	case "\u0077":
		*_gdbd = 62
	case "\u0077\u0041\u0072\u0048":
		*_gdbd = 63
	case "\u0077\u004f\u0066\u0066":
		*_gdbd = 64
	}
	return nil
}

type ST_Breakpoint byte

func (_bcbab *ST_SecondaryChildAlignment) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_bcbab = 0
	case "\u006e\u006f\u006e\u0065":
		*_bcbab = 1
	case "\u0074":
		*_bcbab = 2
	case "\u0062":
		*_bcbab = 3
	case "\u006c":
		*_bcbab = 4
	case "\u0072":
		*_bcbab = 5
	}
	return nil
}

type AG_ConstraintRefAttributes struct {
	RefTypeAttr    ST_ConstraintType
	RefForAttr     ST_ConstraintRelationship
	RefForNameAttr *string
	RefPtTypeAttr  ST_ElementType
}

func NewCT_ColorTransform() *CT_ColorTransform { _fcd := &CT_ColorTransform{}; return _fcd }

func (_cdggf *ST_TextBlockDirection) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ebeae, _cfddb := d.Token()
	if _cfddb != nil {
		return _cfddb
	}
	if _fcegb, _cfdg := _ebeae.(_a.EndElement); _cfdg && _fcegb.Name == start.Name {
		*_cdggf = 1
		return nil
	}
	if _caccg, _dcgc := _ebeae.(_a.CharData); !_dcgc {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ebeae)
	} else {
		switch string(_caccg) {
		case "":
			*_cdggf = 0
		case "\u0068\u006f\u0072\u007a":
			*_cdggf = 1
		case "\u0076\u0065\u0072\u0074":
			*_cdggf = 2
		}
	}
	_ebeae, _cfddb = d.Token()
	if _cfddb != nil {
		return _cfddb
	}
	if _fagae, _gece := _ebeae.(_a.EndElement); _gece && _fagae.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ebeae)
}

func (_eegg ST_AxisType) Validate() error { return _eegg.ValidateWithPath("") }

type ColorsDef struct{ CT_ColorTransform }

// ValidateWithPath validates the CT_Shape and its children, prefixing error messages with path
func (_cdea *CT_Shape) ValidateWithPath(path string) error {
	if _cdea.TypeAttr != nil {
		if _gfec := _cdea.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _gfec != nil {
			return _gfec
		}
	}
	if _cdea.AdjLst != nil {
		if _fgeb := _cdea.AdjLst.ValidateWithPath(path + "\u002fA\u0064\u006a\u004c\u0073\u0074"); _fgeb != nil {
			return _fgeb
		}
	}
	if _cdea.ExtLst != nil {
		if _gcbc := _cdea.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _gcbc != nil {
			return _gcbc
		}
	}
	return nil
}

func (_fdgc *CT_PresentationOf) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _fdgc.AxisAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0061\u0078\u0069\u0073"}, Value: _f.Sprintf("\u0025\u0076", *_fdgc.AxisAttr)})
	}
	if _fdgc.PtTypeAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_fdgc.PtTypeAttr)})
	}
	if _fdgc.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073"}, Value: _f.Sprintf("\u0025\u0076", *_fdgc.HideLastTransAttr)})
	}
	if _fdgc.StAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_fdgc.StAttr)})
	}
	if _fdgc.CntAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u006e\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_fdgc.CntAttr)})
	}
	if _fdgc.StepAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0074\u0065\u0070"}, Value: _f.Sprintf("\u0025\u0076", *_fdgc.StepAttr)})
	}
	e.EncodeToken(start)
	if _fdgc.ExtLst != nil {
		_ebeg := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_fdgc.ExtLst, _ebeg)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_fcfeg ST_VariableType) Validate() error { return _fcfeg.ValidateWithPath("") }

func (_ffcac *ST_ChildAlignment) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ffcac = 0
	case "\u0074":
		*_ffcac = 1
	case "\u0062":
		*_ffcac = 2
	case "\u006c":
		*_ffcac = 3
	case "\u0072":
		*_ffcac = 4
	}
	return nil
}

func (_bgagd ST_DiagramTextAlignment) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_gbga := _a.Attr{}
	_gbga.Name = name
	switch _bgagd {
	case ST_DiagramTextAlignmentUnset:
		_gbga.Value = ""
	case ST_DiagramTextAlignmentL:
		_gbga.Value = "\u006c"
	case ST_DiagramTextAlignmentCtr:
		_gbga.Value = "\u0063\u0074\u0072"
	case ST_DiagramTextAlignmentR:
		_gbga.Value = "\u0072"
	}
	return _gbga, nil
}

func (_gbgbb *ST_SecondaryLinearDirection) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_gbgbb = 0
	case "\u006e\u006f\u006e\u0065":
		*_gbgbb = 1
	case "\u0066\u0072\u006fm\u004c":
		*_gbgbb = 2
	case "\u0066\u0072\u006fm\u0052":
		*_gbgbb = 3
	case "\u0066\u0072\u006fm\u0054":
		*_gbgbb = 4
	case "\u0066\u0072\u006fm\u0042":
		*_gbgbb = 5
	}
	return nil
}

func (_cdffd ST_FallbackDimension) String() string {
	switch _cdffd {
	case 0:
		return ""
	case 1:
		return "\u0031\u0044"
	case 2:
		return "\u0032\u0044"
	}
	return ""
}

// ValidateWithPath validates the CT_NumericRule and its children, prefixing error messages with path
func (_afgc *CT_NumericRule) ValidateWithPath(path string) error {
	if _afgc.ExtLst != nil {
		if _gbbg := _afgc.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _gbbg != nil {
			return _gbbg
		}
	}
	if _ddddb := _afgc.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _ddddb != nil {
		return _ddddb
	}
	if _acbc := _afgc.ForAttr.ValidateWithPath(path + "\u002f\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _acbc != nil {
		return _acbc
	}
	if _deac := _afgc.PtTypeAttr.ValidateWithPath(path + "/\u0050\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _deac != nil {
		return _deac
	}
	return nil
}

func NewCT_Parameter() *CT_Parameter {
	_fdgb := &CT_Parameter{}
	_fdgb.TypeAttr = ST_ParameterId(1)
	return _fdgb
}

func (_bdae ST_Breakpoint) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_babbf := _a.Attr{}
	_babbf.Name = name
	switch _bdae {
	case ST_BreakpointUnset:
		_babbf.Value = ""
	case ST_BreakpointEndCnv:
		_babbf.Value = "\u0065\u006e\u0064\u0043\u006e\u0076"
	case ST_BreakpointBal:
		_babbf.Value = "\u0062\u0061\u006c"
	case ST_BreakpointFixed:
		_babbf.Value = "\u0066\u0069\u0078e\u0064"
	}
	return _babbf, nil
}

func (_fbdcd ST_LinearDirection) String() string {
	switch _fbdcd {
	case 0:
		return ""
	case 1:
		return "\u0066\u0072\u006fm\u004c"
	case 2:
		return "\u0066\u0072\u006fm\u0052"
	case 3:
		return "\u0066\u0072\u006fm\u0054"
	case 4:
		return "\u0066\u0072\u006fm\u0042"
	}
	return ""
}

func (_bged *CT_CTDescription) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _gdf := range start.Attr {
		if _gdf.Name.Local == "\u006c\u0061\u006e\u0067" {
			_edc, _ddac := _gdf.Value, error(nil)
			if _ddac != nil {
				return _ddac
			}
			_bged.LangAttr = &_edc
			continue
		}
		if _gdf.Name.Local == "\u0076\u0061\u006c" {
			_dcc, _bee := _gdf.Value, error(nil)
			if _bee != nil {
				return _bee
			}
			_bged.ValAttr = _dcc
			continue
		}
	}
	for {
		_ffa, _gbf := d.Token()
		if _gbf != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0043\u0054\u005f\u0043\u0054\u0044\u0065s\u0063r\u0069\u0070\u0074\u0069\u006f\u006e\u003a \u0025\u0073", _gbf)
		}
		if _gefb, _dab := _ffa.(_a.EndElement); _dab && _gefb.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ecff ST_AxisType) String() string {
	switch _ecff {
	case 0:
		return ""
	case 1:
		return "\u0073\u0065\u006c\u0066"
	case 2:
		return "\u0063\u0068"
	case 3:
		return "\u0064\u0065\u0073"
	case 4:
		return "\u0064e\u0073\u004f\u0072\u0053\u0065\u006cf"
	case 5:
		return "\u0070\u0061\u0072"
	case 6:
		return "\u0061\u006e\u0063s\u0074"
	case 7:
		return "a\u006e\u0063\u0073\u0074\u004f\u0072\u0053\u0065\u006c\u0066"
	case 8:
		return "\u0066o\u006c\u006c\u006f\u0077\u0053\u0069b"
	case 9:
		return "\u0070r\u0065\u0063\u0065\u0064\u0053\u0069b"
	case 10:
		return "\u0066\u006f\u006c\u006c\u006f\u0077"
	case 11:
		return "\u0070\u0072\u0065\u0063\u0065\u0064"
	case 12:
		return "\u0072\u006f\u006f\u0074"
	case 13:
		return "\u006e\u006f\u006e\u0065"
	}
	return ""
}

func (_gada ST_Breakpoint) ValidateWithPath(path string) error {
	switch _gada {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gada))
	}
	return nil
}

func (_gf *CT_Adj) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0069\u0064\u0078"}, Value: _f.Sprintf("\u0025\u0076", _gf.IdxAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", _gf.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

type CT_Description struct {
	LangAttr *string
	ValAttr  string
}

func (_gegdg ST_NodeVerticalAlignment) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_gegdg.String(), start)
}

func NewCT_SDName() *CT_SDName { _cbee := &CT_SDName{}; return _cbee }

func (_edf *CT_CTStyleLabel) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", _edf.NameAttr)})
	e.EncodeToken(start)
	if _edf.FillClrLst != nil {
		_bdc := _a.StartElement{Name: _a.Name{Local: "\u0066\u0069\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_edf.FillClrLst, _bdc)
	}
	if _edf.LinClrLst != nil {
		_adf := _a.StartElement{Name: _a.Name{Local: "\u006ci\u006e\u0043\u006c\u0072\u004c\u0073t"}}
		e.EncodeElement(_edf.LinClrLst, _adf)
	}
	if _edf.EffectClrLst != nil {
		_fbd := _a.StartElement{Name: _a.Name{Local: "\u0065\u0066\u0066e\u0063\u0074\u0043\u006c\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_edf.EffectClrLst, _fbd)
	}
	if _edf.TxLinClrLst != nil {
		_daed := _a.StartElement{Name: _a.Name{Local: "t\u0078\u004c\u0069\u006e\u0043\u006c\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_edf.TxLinClrLst, _daed)
	}
	if _edf.TxFillClrLst != nil {
		_gcda := _a.StartElement{Name: _a.Name{Local: "\u0074\u0078\u0046i\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_edf.TxFillClrLst, _gcda)
	}
	if _edf.TxEffectClrLst != nil {
		_geee := _a.StartElement{Name: _a.Name{Local: "\u0074\u0078\u0045\u0066\u0066\u0065\u0063\u0074\u0043l\u0072\u004c\u0073\u0074"}}
		e.EncodeElement(_edf.TxEffectClrLst, _geee)
	}
	if _edf.ExtLst != nil {
		_bggd := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_edf.ExtLst, _bggd)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_ccdb *ST_HierBranchStyle) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_effc, _bfgec := d.Token()
	if _bfgec != nil {
		return _bfgec
	}
	if _efcge, _agbca := _effc.(_a.EndElement); _agbca && _efcge.Name == start.Name {
		*_ccdb = 1
		return nil
	}
	if _aaeg, _cfcfe := _effc.(_a.CharData); !_cfcfe {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _effc)
	} else {
		switch string(_aaeg) {
		case "":
			*_ccdb = 0
		case "\u006c":
			*_ccdb = 1
		case "\u0072":
			*_ccdb = 2
		case "\u0068\u0061\u006e\u0067":
			*_ccdb = 3
		case "\u0073\u0074\u0064":
			*_ccdb = 4
		case "\u0069\u006e\u0069\u0074":
			*_ccdb = 5
		}
	}
	_effc, _bfgec = d.Token()
	if _bfgec != nil {
		return _bfgec
	}
	if _efeb, _geaa := _effc.(_a.EndElement); _geaa && _efeb.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _effc)
}

type CT_ResizeHandles struct{ ValAttr ST_ResizeHandlesStr }

func (_dgeea *ST_VariableType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_dgeea = 0
	case "\u006e\u006f\u006e\u0065":
		*_dgeea = 1
	case "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074":
		*_dgeea = 2
	case "\u0063\u0068\u004da\u0078":
		*_dgeea = 3
	case "\u0063\u0068\u0050\u0072\u0065\u0066":
		*_dgeea = 4
	case "\u0062\u0075\u006c\u0045\u006e\u0061\u0062\u006c\u0065\u0064":
		*_dgeea = 5
	case "\u0064\u0069\u0072":
		*_dgeea = 6
	case "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068":
		*_dgeea = 7
	case "\u0061n\u0069\u006d\u004f\u006e\u0065":
		*_dgeea = 8
	case "\u0061n\u0069\u006d\u004c\u0076\u006c":
		*_dgeea = 9
	case "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073":
		*_dgeea = 10
	}
	return nil
}

func (_aaagg ST_ResizeHandlesStr) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_cgggf := _a.Attr{}
	_cgggf.Name = name
	switch _aaagg {
	case ST_ResizeHandlesStrUnset:
		_cgggf.Value = ""
	case ST_ResizeHandlesStrExact:
		_cgggf.Value = "\u0065\u0078\u0061c\u0074"
	case ST_ResizeHandlesStrRel:
		_cgggf.Value = "\u0072\u0065\u006c"
	}
	return _cgggf, nil
}

type AG_ConstraintAttributes struct {
	TypeAttr    ST_ConstraintType
	ForAttr     ST_ConstraintRelationship
	ForNameAttr *string
	PtTypeAttr  ST_ElementType
}

type CT_PtList struct{ Pt []*CT_Pt }

func (_bfga ST_RotationPath) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_cefb := _a.Attr{}
	_cefb.Name = name
	switch _bfga {
	case ST_RotationPathUnset:
		_cefb.Value = ""
	case ST_RotationPathNone:
		_cefb.Value = "\u006e\u006f\u006e\u0065"
	case ST_RotationPathAlongPath:
		_cefb.Value = "\u0061l\u006f\u006e\u0067\u0050\u0061\u0074h"
	}
	return _cefb, nil
}

func (_cccf *CT_CTCategories) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _cccf.Cat != nil {
		_cbed := _a.StartElement{Name: _a.Name{Local: "\u0063\u0061\u0074"}}
		for _, _bec := range _cccf.Cat {
			e.EncodeElement(_bec, _cbed)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

type ST_TextAnchorVertical byte

func NewCT_Constraint() *CT_Constraint { _cbc := &CT_Constraint{}; return _cbc }

type LayoutDef struct{ CT_DiagramDefinition }

func (_ffgcd ST_PyramidAccentTextMargin) String() string {
	switch _ffgcd {
	case 0:
		return ""
	case 1:
		return "\u0073\u0074\u0065\u0070"
	case 2:
		return "\u0073\u0074\u0061c\u006b"
	}
	return ""
}

func (_aage ST_ContinueDirection) String() string {
	switch _aage {
	case 0:
		return ""
	case 1:
		return "\u0072\u0065\u0076\u0044\u0069\u0072"
	case 2:
		return "\u0073a\u006d\u0065\u0044\u0069\u0072"
	}
	return ""
}

type ST_SecondaryLinearDirection byte

func NewCT_StyleDefinitionHeader() *CT_StyleDefinitionHeader {
	_fcabf := &CT_StyleDefinitionHeader{}
	return _fcabf
}

func (_adede ST_CenterShapeMapping) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_egadbe := _a.Attr{}
	_egadbe.Name = name
	switch _adede {
	case ST_CenterShapeMappingUnset:
		_egadbe.Value = ""
	case ST_CenterShapeMappingNone:
		_egadbe.Value = "\u006e\u006f\u006e\u0065"
	case ST_CenterShapeMappingFNode:
		_egadbe.Value = "\u0066\u004e\u006fd\u0065"
	}
	return _egadbe, nil
}

const (
	ST_ContinueDirectionUnset   ST_ContinueDirection = 0
	ST_ContinueDirectionRevDir  ST_ContinueDirection = 1
	ST_ContinueDirectionSameDir ST_ContinueDirection = 2
)

func NewCT_CTDescription() *CT_CTDescription { _fedc := &CT_CTDescription{}; return _fedc }

func (_fafee ST_HierBranchStyle) String() string {
	switch _fafee {
	case 0:
		return ""
	case 1:
		return "\u006c"
	case 2:
		return "\u0072"
	case 3:
		return "\u0068\u0061\u006e\u0067"
	case 4:
		return "\u0073\u0074\u0064"
	case 5:
		return "\u0069\u006e\u0069\u0074"
	}
	return ""
}

func (_gbfa *ST_PrSetCustVal) Validate() error { return _gbfa.ValidateWithPath("") }

func (_bgab *CT_DataModel) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	_egba := _a.StartElement{Name: _a.Name{Local: "\u0070\u0074\u004cs\u0074"}}
	e.EncodeElement(_bgab.PtLst, _egba)
	if _bgab.CxnLst != nil {
		_fbff := _a.StartElement{Name: _a.Name{Local: "\u0063\u0078\u006e\u004c\u0073\u0074"}}
		e.EncodeElement(_bgab.CxnLst, _fbff)
	}
	if _bgab.Bg != nil {
		_eag := _a.StartElement{Name: _a.Name{Local: "\u0062\u0067"}}
		e.EncodeElement(_bgab.Bg, _eag)
	}
	if _bgab.Whole != nil {
		_cbde := _a.StartElement{Name: _a.Name{Local: "\u0077\u0068\u006fl\u0065"}}
		e.EncodeElement(_bgab.Whole, _cbde)
	}
	if _bgab.ExtLst != nil {
		_cdad := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_bgab.ExtLst, _cdad)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func NewCT_SDCategory() *CT_SDCategory { _caga := &CT_SDCategory{}; return _caga }

func (_dafb *CT_NumericRule) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _dafb.ValAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", *_dafb.ValAttr)})
	}
	if _dafb.FactAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0066\u0061\u0063\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_dafb.FactAttr)})
	}
	if _dafb.MaxAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006d\u0061\u0078"}, Value: _f.Sprintf("\u0025\u0076", *_dafb.MaxAttr)})
	}
	if _dafb.TypeAttr != ST_ConstraintTypeUnset {
		_eabf, _fgfb := _dafb.TypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _fgfb != nil {
			return _fgfb
		}
		start.Attr = append(start.Attr, _eabf)
	}
	if _dafb.ForAttr != ST_ConstraintRelationshipUnset {
		_cabb, _fdeb := _dafb.ForAttr.MarshalXMLAttr(_a.Name{Local: "\u0066\u006f\u0072"})
		if _fdeb != nil {
			return _fdeb
		}
		start.Attr = append(start.Attr, _cabb)
	}
	if _dafb.ForNameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0066o\u0072\u004e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_dafb.ForNameAttr)})
	}
	if _dafb.PtTypeAttr != ST_ElementTypeUnset {
		_bdeee, _bbbad := _dafb.PtTypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"})
		if _bbbad != nil {
			return _bbbad
		}
		start.Attr = append(start.Attr, _bdeee)
	}
	e.EncodeToken(start)
	if _dafb.ExtLst != nil {
		_addf := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_dafb.ExtLst, _addf)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ST_ModelId is a union type
type ST_ModelId struct {
	Int32   *int32
	ST_Guid *string
}

func (_gfgae ST_ClrAppMethod) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_gfgae.String(), start)
}

func (_bcadf ST_LayoutShapeType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _bcadf.ST_ShapeType != _bg.ST_ShapeTypeUnset {
		e.EncodeToken(_a.CharData(_bcadf.ST_ShapeType.String()))
	}
	if _bcadf.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		e.EncodeToken(_a.CharData(_bcadf.ST_OutputShapeType.String()))
	}
	return e.EncodeToken(_a.EndElement{Name: start.Name})
}

func (_aafg ST_AnimLvlStr) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_cfcfec := _a.Attr{}
	_cfcfec.Name = name
	switch _aafg {
	case ST_AnimLvlStrUnset:
		_cfcfec.Value = ""
	case ST_AnimLvlStrNone:
		_cfcfec.Value = "\u006e\u006f\u006e\u0065"
	case ST_AnimLvlStrLvl:
		_cfcfec.Value = "\u006c\u0076\u006c"
	case ST_AnimLvlStrCtr:
		_cfcfec.Value = "\u0063\u0074\u0072"
	}
	return _cfcfec, nil
}

func (_ffgc ST_AxisType) ValidateWithPath(path string) error {
	switch _ffgc {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ffgc))
	}
	return nil
}

const (
	ST_ElementTypeUnset    ST_ElementType = 0
	ST_ElementTypeAll      ST_ElementType = 1
	ST_ElementTypeDoc      ST_ElementType = 2
	ST_ElementTypeNode     ST_ElementType = 3
	ST_ElementTypeNorm     ST_ElementType = 4
	ST_ElementTypeNonNorm  ST_ElementType = 5
	ST_ElementTypeAsst     ST_ElementType = 6
	ST_ElementTypeNonAsst  ST_ElementType = 7
	ST_ElementTypeParTrans ST_ElementType = 8
	ST_ElementTypePres     ST_ElementType = 9
	ST_ElementTypeSibTrans ST_ElementType = 10
)

func (_eecge ST_TextAnchorHorizontal) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_fbgc := _a.Attr{}
	_fbgc.Name = name
	switch _eecge {
	case ST_TextAnchorHorizontalUnset:
		_fbgc.Value = ""
	case ST_TextAnchorHorizontalNone:
		_fbgc.Value = "\u006e\u006f\u006e\u0065"
	case ST_TextAnchorHorizontalCtr:
		_fbgc.Value = "\u0063\u0074\u0072"
	}
	return _fbgc, nil
}

func (_ggbe *ST_Offset) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ggbe = 0
	case "\u0063\u0074\u0072":
		*_ggbe = 1
	case "\u006f\u0066\u0066":
		*_ggbe = 2
	}
	return nil
}

func (_bcff *CT_DataModel) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_bcff.PtLst = NewCT_PtList()
_eeed:
	for {
		_efgf, _bfgg := d.Token()
		if _bfgg != nil {
			return _bfgg
		}
		switch _cdfg := _efgf.(type) {
		case _a.StartElement:
			switch _cdfg.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0074\u004cs\u0074"}:
				if _cbgf := d.DecodeElement(_bcff.PtLst, &_cdfg); _cbgf != nil {
					return _cbgf
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0078\u006e\u004c\u0073\u0074"}:
				_bcff.CxnLst = NewCT_CxnList()
				if _agec := d.DecodeElement(_bcff.CxnLst, &_cdfg); _agec != nil {
					return _agec
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0062\u0067"}:
				_bcff.Bg = _bg.NewCT_BackgroundFormatting()
				if _dcca := d.DecodeElement(_bcff.Bg, &_cdfg); _dcca != nil {
					return _dcca
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0077\u0068\u006fl\u0065"}:
				_bcff.Whole = _bg.NewCT_WholeE2oFormatting()
				if _ecec := d.DecodeElement(_bcff.Whole, &_cdfg); _ecec != nil {
					return _ecec
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_bcff.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _ebfg := d.DecodeElement(_bcff.ExtLst, &_cdfg); _ebfg != nil {
					return _ebfg
				}
			default:
				_b.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_D\u0061\u0074a\u004d\u006f\u0064\u0065\u006c\u0020\u0025\u0076", _cdfg.Name)
				if _bfef := d.Skip(); _bfef != nil {
					return _bfef
				}
			}
		case _a.EndElement:
			break _eeed
		case _a.CharData:
		}
	}
	return nil
}

func (_fae *CT_Choose) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _eaf := range start.Attr {
		if _eaf.Name.Local == "\u006e\u0061\u006d\u0065" {
			_cbf, _agda := _eaf.Value, error(nil)
			if _agda != nil {
				return _agda
			}
			_fae.NameAttr = &_cbf
			continue
		}
	}
_ageb:
	for {
		_dbdb, _aaa := d.Token()
		if _aaa != nil {
			return _aaa
		}
		switch _gdaa := _dbdb.(type) {
		case _a.StartElement:
			switch _gdaa.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0069\u0066"}:
				_bgd := NewCT_When()
				if _beec := d.DecodeElement(_bgd, &_gdaa); _beec != nil {
					return _beec
				}
				_fae.If = append(_fae.If, _bgd)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u006c\u0073\u0065"}:
				_fae.Else = NewCT_Otherwise()
				if _fcc := d.DecodeElement(_fae.Else, &_gdaa); _fcc != nil {
					return _fcc
				}
			default:
				_b.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0043\u0068\u006fo\u0073\u0065 \u0025\u0076", _gdaa.Name)
				if _adgc := d.Skip(); _adgc != nil {
					return _adgc
				}
			}
		case _a.EndElement:
			break _ageb
		case _a.CharData:
		}
	}
	return nil
}

func (_gecc ST_ClrAppMethod) ValidateWithPath(path string) error {
	switch _gecc {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gecc))
	}
	return nil
}

type ST_ConstraintRelationship byte

// ValidateWithPath validates the CT_Pt and its children, prefixing error messages with path
func (_cbad *CT_Pt) ValidateWithPath(path string) error {
	if _dabe := _cbad.ModelIdAttr.ValidateWithPath(path + "\u002f\u004d\u006fd\u0065\u006c\u0049\u0064\u0041\u0074\u0074\u0072"); _dabe != nil {
		return _dabe
	}
	if _gcecf := _cbad.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _gcecf != nil {
		return _gcecf
	}
	if _cbad.CxnIdAttr != nil {
		if _ebeeg := _cbad.CxnIdAttr.ValidateWithPath(path + "\u002f\u0043\u0078\u006e\u0049\u0064\u0041\u0074\u0074\u0072"); _ebeeg != nil {
			return _ebeeg
		}
	}
	if _cbad.PrSet != nil {
		if _befc := _cbad.PrSet.ValidateWithPath(path + "\u002f\u0050\u0072\u0053\u0065\u0074"); _befc != nil {
			return _befc
		}
	}
	if _cbad.SpPr != nil {
		if _ffaf := _cbad.SpPr.ValidateWithPath(path + "\u002f\u0053\u0070P\u0072"); _ffaf != nil {
			return _ffaf
		}
	}
	if _cbad.T != nil {
		if _fdcb := _cbad.T.ValidateWithPath(path + "\u002f\u0054"); _fdcb != nil {
			return _fdcb
		}
	}
	if _cbad.ExtLst != nil {
		if _gdfe := _cbad.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _gdfe != nil {
			return _gdfe
		}
	}
	return nil
}

func (_ddd *CT_CTCategories) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_cbb:
	for {
		_gbae, _ggfa := d.Token()
		if _ggfa != nil {
			return _ggfa
		}
		switch _dbc := _gbae.(type) {
		case _a.StartElement:
			switch _dbc.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074"}:
				_agbc := NewCT_CTCategory()
				if _ega := d.DecodeElement(_agbc, &_dbc); _ega != nil {
					return _ega
				}
				_ddd.Cat = append(_ddd.Cat, _agbc)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u0054\u0043a\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073\u0020\u0025\u0076", _dbc.Name)
				if _dca := d.Skip(); _dca != nil {
					return _dca
				}
			}
		case _a.EndElement:
			break _cbb
		case _a.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_AnimOne and its children, prefixing error messages with path
func (_aea *CT_AnimOne) ValidateWithPath(path string) error {
	if _gfge := _aea.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _gfge != nil {
		return _gfge
	}
	return nil
}

func (_ddde *ST_FunctionArgument) Validate() error { return _ddde.ValidateWithPath("") }

func (_dbbe *CT_Name) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _dbbe.LangAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _f.Sprintf("\u0025\u0076", *_dbbe.LangAttr)})
	}
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", _dbbe.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_fbaee *ST_ConnectorRouting) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_cgfae, _egfb := d.Token()
	if _egfb != nil {
		return _egfb
	}
	if _eebg, _ebaf := _cgfae.(_a.EndElement); _ebaf && _eebg.Name == start.Name {
		*_fbaee = 1
		return nil
	}
	if _eeddg, _deacb := _cgfae.(_a.CharData); !_deacb {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cgfae)
	} else {
		switch string(_eeddg) {
		case "":
			*_fbaee = 0
		case "\u0073\u0074\u0072\u0061":
			*_fbaee = 1
		case "\u0062\u0065\u006e\u0064":
			*_fbaee = 2
		case "\u0063\u0075\u0072v\u0065":
			*_fbaee = 3
		case "\u006co\u006e\u0067\u0043\u0075\u0072\u0076e":
			*_fbaee = 4
		}
	}
	_cgfae, _egfb = d.Token()
	if _egfb != nil {
		return _egfb
	}
	if _cfgeg, _cgab := _cgfae.(_a.EndElement); _cgab && _cfgeg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cgfae)
}

func (_aggg *CT_Shape) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _dbbec := range start.Attr {
		if _dbbec.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _dbbec.Name.Local == "\u0062\u006c\u0069\u0070" || _dbbec.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _dbbec.Name.Local == "\u0062\u006c\u0069\u0070" {
			_gege, _aeef := _dbbec.Value, error(nil)
			if _aeef != nil {
				return _aeef
			}
			_aggg.BlipAttr = &_gege
			continue
		}
		if _dbbec.Name.Local == "\u0072\u006f\u0074" {
			_gebg, _abdga := _ae.ParseFloat(_dbbec.Value, 64)
			if _abdga != nil {
				return _abdga
			}
			_aggg.RotAttr = &_gebg
			continue
		}
		if _dbbec.Name.Local == "\u0074\u0079\u0070\u0065" {
			_bcee, _fdag := ParseUnionST_LayoutShapeType(_dbbec.Value)
			if _fdag != nil {
				return _fdag
			}
			_aggg.TypeAttr = &_bcee
			continue
		}
		if _dbbec.Name.Local == "\u007aO\u0072\u0064\u0065\u0072\u004f\u0066f" {
			_eaeg, _cdfdd := _ae.ParseInt(_dbbec.Value, 10, 32)
			if _cdfdd != nil {
				return _cdfdd
			}
			_cebc := int32(_eaeg)
			_aggg.ZOrderOffAttr = &_cebc
			continue
		}
		if _dbbec.Name.Local == "\u0068\u0069\u0064\u0065\u0047\u0065\u006f\u006d" {
			_ggdf, _geab := _ae.ParseBool(_dbbec.Value)
			if _geab != nil {
				return _geab
			}
			_aggg.HideGeomAttr = &_ggdf
			continue
		}
		if _dbbec.Name.Local == "\u006ck\u0054\u0078\u0045\u006e\u0074\u0072y" {
			_acde, _fage := _ae.ParseBool(_dbbec.Value)
			if _fage != nil {
				return _fage
			}
			_aggg.LkTxEntryAttr = &_acde
			continue
		}
		if _dbbec.Name.Local == "\u0062l\u0069\u0070\u0050\u0068\u006c\u0064r" {
			_gadb, _bbdb := _ae.ParseBool(_dbbec.Value)
			if _bbdb != nil {
				return _bbdb
			}
			_aggg.BlipPhldrAttr = &_gadb
			continue
		}
	}
_bdgc:
	for {
		_fcec, _daecg := d.Token()
		if _daecg != nil {
			return _daecg
		}
		switch _edga := _fcec.(type) {
		case _a.StartElement:
			switch _edga.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u0064\u006a\u004c\u0073\u0074"}:
				_aggg.AdjLst = NewCT_AdjLst()
				if _gfcd := d.DecodeElement(_aggg.AdjLst, &_edga); _gfcd != nil {
					return _gfcd
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_aggg.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _efaa := d.DecodeElement(_aggg.ExtLst, &_edga); _efaa != nil {
					return _efaa
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065\u0020\u0025\u0076", _edga.Name)
				if _ggbcd := d.Skip(); _ggbcd != nil {
					return _ggbcd
				}
			}
		case _a.EndElement:
			break _bdgc
		case _a.CharData:
		}
	}
	return nil
}

func (_gbbaa ST_FlowDirection) ValidateWithPath(path string) error {
	switch _gbbaa {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gbbaa))
	}
	return nil
}

func NewRelIds() *RelIds { _fgdac := &RelIds{}; _fgdac.CT_RelIds = *NewCT_RelIds(); return _fgdac }

// Validate validates the CT_StyleDefinitionHeaderLst and its children
func (_daebf *CT_StyleDefinitionHeaderLst) Validate() error {
	return _daebf.ValidateWithPath("C\u0054\u005f\u0053\u0074\u0079\u006ce\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006fn\u0048\u0065\u0061d\u0065r\u004c\u0073\u0074")
}

func (_bgfa *StyleDef) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0073\u0074\u0079\u006c\u0065\u0044\u0065\u0066"
	return _bgfa.CT_StyleDefinition.MarshalXML(e, start)
}

func (_bcccb ST_StartingElement) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_addgf := _a.Attr{}
	_addgf.Name = name
	switch _bcccb {
	case ST_StartingElementUnset:
		_addgf.Value = ""
	case ST_StartingElementNode:
		_addgf.Value = "\u006e\u006f\u0064\u0065"
	case ST_StartingElementTrans:
		_addgf.Value = "\u0074\u0072\u0061n\u0073"
	}
	return _addgf, nil
}

// ValidateWithPath validates the CT_ColorTransform and its children, prefixing error messages with path
func (_eacb *CT_ColorTransform) ValidateWithPath(path string) error {
	for _ecf, _aacf := range _eacb.Title {
		if _bea := _aacf.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _ecf)); _bea != nil {
			return _bea
		}
	}
	for _ffdf, _gdfg := range _eacb.Desc {
		if _dcbc := _gdfg.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _ffdf)); _dcbc != nil {
			return _dcbc
		}
	}
	if _eacb.CatLst != nil {
		if _ecd := _eacb.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _ecd != nil {
			return _ecd
		}
	}
	for _ecbb, _bag := range _eacb.StyleLbl {
		if _bab := _bag.ValidateWithPath(_f.Sprintf("\u0025s\u002fS\u0074\u0079\u006c\u0065\u004c\u0062\u006c\u005b\u0025\u0064\u005d", path, _ecbb)); _bab != nil {
			return _bab
		}
	}
	if _eacb.ExtLst != nil {
		if _ece := _eacb.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _ece != nil {
			return _ece
		}
	}
	return nil
}

func (_bbeg *CT_PresentationOf) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _cdcc := range start.Attr {
		if _cdcc.Name.Local == "\u0061\u0078\u0069\u0073" {
			_fbdgf, _ebeb := ParseSliceST_AxisTypes(_cdcc.Value)
			if _ebeb != nil {
				return _ebeb
			}
			_bbeg.AxisAttr = &_fbdgf
			continue
		}
		if _cdcc.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_gggfc, _efcd := ParseSliceST_ElementTypes(_cdcc.Value)
			if _efcd != nil {
				return _efcd
			}
			_bbeg.PtTypeAttr = &_gggfc
			continue
		}
		if _cdcc.Name.Local == "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073" {
			_dfac, _aebf := ParseSliceST_Booleans(_cdcc.Value)
			if _aebf != nil {
				return _aebf
			}
			_bbeg.HideLastTransAttr = &_dfac
			continue
		}
		if _cdcc.Name.Local == "\u0073\u0074" {
			_cdfga, _ecad := ParseSliceST_Ints(_cdcc.Value)
			if _ecad != nil {
				return _ecad
			}
			_bbeg.StAttr = &_cdfga
			continue
		}
		if _cdcc.Name.Local == "\u0063\u006e\u0074" {
			_dfag, _ageg := ParseSliceST_UnsignedInts(_cdcc.Value)
			if _ageg != nil {
				return _ageg
			}
			_bbeg.CntAttr = &_dfag
			continue
		}
		if _cdcc.Name.Local == "\u0073\u0074\u0065\u0070" {
			_gadfc, _fgged := ParseSliceST_Ints(_cdcc.Value)
			if _fgged != nil {
				return _fgged
			}
			_bbeg.StepAttr = &_gadfc
			continue
		}
	}
_cabf:
	for {
		_aagb, _gecfb := d.Token()
		if _gecfb != nil {
			return _gecfb
		}
		switch _bbea := _aagb.(type) {
		case _a.StartElement:
			switch _bbea.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_bbeg.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _eeda := d.DecodeElement(_bbeg.ExtLst, &_bbea); _eeda != nil {
					return _eeda
				}
			default:
				_b.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u004f\u0066\u0020\u0025\u0076", _bbea.Name)
				if _ffdff := d.Skip(); _ffdff != nil {
					return _ffdff
				}
			}
		case _a.EndElement:
			break _cabf
		case _a.CharData:
		}
	}
	return nil
}

const (
	ST_HueDirUnset ST_HueDir = 0
	ST_HueDirCw    ST_HueDir = 1
	ST_HueDirCcw   ST_HueDir = 2
)

// ValidateWithPath validates the CT_StyleDefinition and its children, prefixing error messages with path
func (_ebff *CT_StyleDefinition) ValidateWithPath(path string) error {
	for _dbbb, _gabc := range _ebff.Title {
		if _dbfe := _gabc.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _dbbb)); _dbfe != nil {
			return _dbfe
		}
	}
	for _fgdb, _dcdd := range _ebff.Desc {
		if _ebcd := _dcdd.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _fgdb)); _ebcd != nil {
			return _ebcd
		}
	}
	if _ebff.CatLst != nil {
		if _aaaeb := _ebff.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _aaaeb != nil {
			return _aaaeb
		}
	}
	if _ebff.Scene3d != nil {
		if _dcfb := _ebff.Scene3d.ValidateWithPath(path + "\u002f\u0053\u0063\u0065\u006e\u0065\u0033\u0064"); _dcfb != nil {
			return _dcfb
		}
	}
	for _dcfg, _cbcf := range _ebff.StyleLbl {
		if _gfebg := _cbcf.ValidateWithPath(_f.Sprintf("\u0025s\u002fS\u0074\u0079\u006c\u0065\u004c\u0062\u006c\u005b\u0025\u0064\u005d", path, _dcfg)); _gfebg != nil {
			return _gfebg
		}
	}
	if _ebff.ExtLst != nil {
		if _dacec := _ebff.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _dacec != nil {
			return _dacec
		}
	}
	return nil
}

// ValidateWithPath validates the CT_StyleLabel and its children, prefixing error messages with path
func (_bdeeee *CT_StyleLabel) ValidateWithPath(path string) error {
	if _bdeeee.Scene3d != nil {
		if _eada := _bdeeee.Scene3d.ValidateWithPath(path + "\u002f\u0053\u0063\u0065\u006e\u0065\u0033\u0064"); _eada != nil {
			return _eada
		}
	}
	if _bdeeee.Sp3d != nil {
		if _gbgg := _bdeeee.Sp3d.ValidateWithPath(path + "\u002f\u0053\u00703\u0064"); _gbgg != nil {
			return _gbgg
		}
	}
	if _bdeeee.TxPr != nil {
		if _fbef := _bdeeee.TxPr.ValidateWithPath(path + "\u002f\u0054\u0078P\u0072"); _fbef != nil {
			return _fbef
		}
	}
	if _bdeeee.Style != nil {
		if _fddd := _bdeeee.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _fddd != nil {
			return _fddd
		}
	}
	if _bdeeee.ExtLst != nil {
		if _ceba := _bdeeee.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _ceba != nil {
			return _ceba
		}
	}
	return nil
}

func (_abcb ST_DiagramHorizontalAlignment) ValidateWithPath(path string) error {
	switch _abcb {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_abcb))
	}
	return nil
}

func (_edg *CT_ColorTransform) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _beeca := range start.Attr {
		if _beeca.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_adgfd, _ceb := _beeca.Value, error(nil)
			if _ceb != nil {
				return _ceb
			}
			_edg.UniqueIdAttr = &_adgfd
			continue
		}
		if _beeca.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_dfef, _daag := _beeca.Value, error(nil)
			if _daag != nil {
				return _daag
			}
			_edg.MinVerAttr = &_dfef
			continue
		}
	}
_ebee:
	for {
		_gdcf, _fedb := d.Token()
		if _fedb != nil {
			return _fedb
		}
		switch _gec := _gdcf.(type) {
		case _a.StartElement:
			switch _gec.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_ccb := NewCT_CTName()
				if _efcf := d.DecodeElement(_ccb, &_gec); _efcf != nil {
					return _efcf
				}
				_edg.Title = append(_edg.Title, _ccb)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_ffd := NewCT_CTDescription()
				if _edce := d.DecodeElement(_ffd, &_gec); _edce != nil {
					return _edce
				}
				_edg.Desc = append(_edg.Desc, _ffd)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_edg.CatLst = NewCT_CTCategories()
				if _abea := d.DecodeElement(_edg.CatLst, &_gec); _abea != nil {
					return _abea
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}:
				_fgf := NewCT_CTStyleLabel()
				if _baf := d.DecodeElement(_fgf, &_gec); _baf != nil {
					return _baf
				}
				_edg.StyleLbl = append(_edg.StyleLbl, _fgf)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_edg.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _bdfd := d.DecodeElement(_edg.ExtLst, &_gec); _bdfd != nil {
					return _bdfd
				}
			default:
				_b.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u0020\u0025\u0076", _gec.Name)
				if _adc := d.Skip(); _adc != nil {
					return _adc
				}
			}
		case _a.EndElement:
			break _ebee
		case _a.CharData:
		}
	}
	return nil
}

func (_ggdae ST_ContinueDirection) ValidateWithPath(path string) error {
	switch _ggdae {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ggdae))
	}
	return nil
}

type ColorsDefHdr struct{ CT_ColorTransformHeader }

// Validate validates the StyleDef and its children
func (_cgaf *StyleDef) Validate() error {
	return _cgaf.ValidateWithPath("\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066")
}

type CT_CTCategories struct{ Cat []*CT_CTCategory }

type ST_AlgorithmType byte

func (_ebgb *ST_ChildDirection) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ebgb = 0
	case "\u0068\u006f\u0072\u007a":
		*_ebgb = 1
	case "\u0076\u0065\u0072\u0074":
		*_ebgb = 2
	}
	return nil
}

func NewCT_SampleData() *CT_SampleData { _eggb := &CT_SampleData{}; return _eggb }

func (_dbae *ST_PyramidAccentPosition) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_dbae = 0
	case "\u0062\u0065\u0066":
		*_dbae = 1
	case "\u0061\u0066\u0074":
		*_dbae = 2
	}
	return nil
}

func (_cgag *ST_ChildOrderType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_eagfb, _daeff := d.Token()
	if _daeff != nil {
		return _daeff
	}
	if _agabb, _agaaf := _eagfb.(_a.EndElement); _agaaf && _agabb.Name == start.Name {
		*_cgag = 1
		return nil
	}
	if _agea, _cbae := _eagfb.(_a.CharData); !_cbae {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eagfb)
	} else {
		switch string(_agea) {
		case "":
			*_cgag = 0
		case "\u0062":
			*_cgag = 1
		case "\u0074":
			*_cgag = 2
		}
	}
	_eagfb, _daeff = d.Token()
	if _daeff != nil {
		return _daeff
	}
	if _fffe, _geegf := _eagfb.(_a.EndElement); _geegf && _fffe.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eagfb)
}

// ValidateWithPath validates the CT_SDCategories and its children, prefixing error messages with path
func (_cdcbg *CT_SDCategories) ValidateWithPath(path string) error {
	for _ebea, _fabb := range _cdcbg.Cat {
		if _dfbfg := _fabb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0061\u0074\u005b\u0025\u0064\u005d", path, _ebea)); _dfbfg != nil {
			return _dfbfg
		}
	}
	return nil
}

func NewCT_ColorTransformHeaderLst() *CT_ColorTransformHeaderLst {
	_ebed := &CT_ColorTransformHeaderLst{}
	return _ebed
}

type CT_Choose struct {
	NameAttr *string
	If       []*CT_When
	Else     *CT_Otherwise
}

func (_ebe *CT_Adj) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ebe.IdxAttr = 1
	for _, _cda := range start.Attr {
		if _cda.Name.Local == "\u0069\u0064\u0078" {
			_gab, _ddf := _ae.ParseUint(_cda.Value, 10, 32)
			if _ddf != nil {
				return _ddf
			}
			_ebe.IdxAttr = uint32(_gab)
			continue
		}
		if _cda.Name.Local == "\u0076\u0061\u006c" {
			_cfc, _afc := _ae.ParseFloat(_cda.Value, 64)
			if _afc != nil {
				return _afc
			}
			_ebe.ValAttr = _cfc
			continue
		}
	}
	for {
		_abb, _bb := d.Token()
		if _bb != nil {
			return _f.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0041d\u006a\u003a\u0020\u0025\u0073", _bb)
		}
		if _fc, _agc := _abb.(_a.EndElement); _agc && _fc.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_PtList and its children, prefixing error messages with path
func (_fbabc *CT_PtList) ValidateWithPath(path string) error {
	for _cfea, _fgcg := range _fbabc.Pt {
		if _ebbe := _fgcg.ValidateWithPath(_f.Sprintf("\u0025s\u002f\u0050\u0074\u005b\u0025\u0064]", path, _cfea)); _ebbe != nil {
			return _ebbe
		}
	}
	return nil
}

func (_dcbg *ST_FlowDirection) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_dcbg = 0
	case "\u0072\u006f\u0077":
		*_dcbg = 1
	case "\u0063\u006f\u006c":
		*_dcbg = 2
	}
	return nil
}

type CT_BulletEnabled struct{ ValAttr *bool }

func (_dbdgc *ST_TextAnchorHorizontal) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_dbdgc = 0
	case "\u006e\u006f\u006e\u0065":
		*_dbdgc = 1
	case "\u0063\u0074\u0072":
		*_dbdgc = 2
	}
	return nil
}

const (
	ST_ResizeHandlesStrUnset ST_ResizeHandlesStr = 0
	ST_ResizeHandlesStrExact ST_ResizeHandlesStr = 1
	ST_ResizeHandlesStrRel   ST_ResizeHandlesStr = 2
)

type CT_Category struct {
	TypeAttr string
	PriAttr  uint32
}

func (_abgc *CT_StyleLabel) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _bdbbaf := range start.Attr {
		if _bdbbaf.Name.Local == "\u006e\u0061\u006d\u0065" {
			_eacd, _bfag := _bdbbaf.Value, error(nil)
			if _bfag != nil {
				return _bfag
			}
			_abgc.NameAttr = _eacd
			continue
		}
	}
_aebfb:
	for {
		_gdge, _fbffd := d.Token()
		if _fbffd != nil {
			return _fbffd
		}
		switch _fade := _gdge.(type) {
		case _a.StartElement:
			switch _fade.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}:
				_abgc.Scene3d = _bg.NewCT_Scene3D()
				if _ggca := d.DecodeElement(_abgc.Scene3d, &_fade); _ggca != nil {
					return _ggca
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0070\u0033\u0064"}:
				_abgc.Sp3d = _bg.NewCT_Shape3D()
				if _acec := d.DecodeElement(_abgc.Sp3d, &_fade); _acec != nil {
					return _acec
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0078\u0050\u0072"}:
				_abgc.TxPr = NewCT_TextProps()
				if _cbgeb := d.DecodeElement(_abgc.TxPr, &_fade); _cbgeb != nil {
					return _cbgeb
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079l\u0065"}:
				_abgc.Style = _bg.NewCT_ShapeStyle()
				if _aecd := d.DecodeElement(_abgc.Style, &_fade); _aecd != nil {
					return _aecd
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_abgc.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _aaedf := d.DecodeElement(_abgc.ExtLst, &_fade); _aaedf != nil {
					return _aaedf
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053t\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c \u0025\u0076", _fade.Name)
				if _eagb := d.Skip(); _eagb != nil {
					return _eagb
				}
			}
		case _a.EndElement:
			break _aebfb
		case _a.CharData:
		}
	}
	return nil
}

func (_ddfg ST_SecondaryLinearDirection) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_ddfg.String(), start)
}

func (_gdea *ST_BendPoint) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_gdea = 0
	case "\u0062\u0065\u0067":
		*_gdea = 1
	case "\u0064\u0065\u0066":
		*_gdea = 2
	case "\u0065\u006e\u0064":
		*_gdea = 3
	}
	return nil
}

func (_fdec ST_TextAnchorHorizontal) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_fdec.String(), start)
}

// Validate validates the CT_Shape and its children
func (_bffbg *CT_Shape) Validate() error {
	return _bffbg.ValidateWithPath("\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065")
}

// ValidateWithPath validates the CT_SDCategory and its children, prefixing error messages with path
func (_dedfe *CT_SDCategory) ValidateWithPath(path string) error { return nil }

func (_eafaac ST_AutoTextRotation) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_eafaac.String(), start)
}

func (_eabbb ST_ConnectorPoint) Validate() error { return _eabbb.ValidateWithPath("") }

func NewLayoutDef() *LayoutDef {
	_cgbg := &LayoutDef{}
	_cgbg.CT_DiagramDefinition = *NewCT_DiagramDefinition()
	return _cgbg
}

func (_beag ST_ChildDirection) Validate() error { return _beag.ValidateWithPath("") }

// Validate validates the LayoutDefHdrLst and its children
func (_eagdf *LayoutDefHdrLst) Validate() error {
	return _eagdf.ValidateWithPath("\u004ca\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074")
}

func ParseSliceST_Booleans(s string) (ST_Booleans, error) { return ST_Booleans{}, nil }

func (_cdcdg ST_ConstraintType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_cbbgf := _a.Attr{}
	_cbbgf.Name = name
	switch _cdcdg {
	case ST_ConstraintTypeUnset:
		_cbbgf.Value = ""
	case ST_ConstraintTypeNone:
		_cbbgf.Value = "\u006e\u006f\u006e\u0065"
	case ST_ConstraintTypeAlignOff:
		_cbbgf.Value = "\u0061\u006c\u0069\u0067\u006e\u004f\u0066\u0066"
	case ST_ConstraintTypeBegMarg:
		_cbbgf.Value = "\u0062e\u0067\u004d\u0061\u0072\u0067"
	case ST_ConstraintTypeBendDist:
		_cbbgf.Value = "\u0062\u0065\u006e\u0064\u0044\u0069\u0073\u0074"
	case ST_ConstraintTypeBegPad:
		_cbbgf.Value = "\u0062\u0065\u0067\u0050\u0061\u0064"
	case ST_ConstraintTypeB:
		_cbbgf.Value = "\u0062"
	case ST_ConstraintTypeBMarg:
		_cbbgf.Value = "\u0062\u004d\u0061r\u0067"
	case ST_ConstraintTypeBOff:
		_cbbgf.Value = "\u0062\u004f\u0066\u0066"
	case ST_ConstraintTypeCtrX:
		_cbbgf.Value = "\u0063\u0074\u0072\u0058"
	case ST_ConstraintTypeCtrXOff:
		_cbbgf.Value = "\u0063t\u0072\u0058\u004f\u0066\u0066"
	case ST_ConstraintTypeCtrY:
		_cbbgf.Value = "\u0063\u0074\u0072\u0059"
	case ST_ConstraintTypeCtrYOff:
		_cbbgf.Value = "\u0063t\u0072\u0059\u004f\u0066\u0066"
	case ST_ConstraintTypeConnDist:
		_cbbgf.Value = "\u0063\u006f\u006e\u006e\u0044\u0069\u0073\u0074"
	case ST_ConstraintTypeDiam:
		_cbbgf.Value = "\u0064\u0069\u0061\u006d"
	case ST_ConstraintTypeEndMarg:
		_cbbgf.Value = "\u0065n\u0064\u004d\u0061\u0072\u0067"
	case ST_ConstraintTypeEndPad:
		_cbbgf.Value = "\u0065\u006e\u0064\u0050\u0061\u0064"
	case ST_ConstraintTypeH:
		_cbbgf.Value = "\u0068"
	case ST_ConstraintTypeHArH:
		_cbbgf.Value = "\u0068\u0041\u0072\u0048"
	case ST_ConstraintTypeHOff:
		_cbbgf.Value = "\u0068\u004f\u0066\u0066"
	case ST_ConstraintTypeL:
		_cbbgf.Value = "\u006c"
	case ST_ConstraintTypeLMarg:
		_cbbgf.Value = "\u006c\u004d\u0061r\u0067"
	case ST_ConstraintTypeLOff:
		_cbbgf.Value = "\u006c\u004f\u0066\u0066"
	case ST_ConstraintTypeR:
		_cbbgf.Value = "\u0072"
	case ST_ConstraintTypeRMarg:
		_cbbgf.Value = "\u0072\u004d\u0061r\u0067"
	case ST_ConstraintTypeROff:
		_cbbgf.Value = "\u0072\u004f\u0066\u0066"
	case ST_ConstraintTypePrimFontSz:
		_cbbgf.Value = "\u0070\u0072\u0069\u006d\u0046\u006f\u006e\u0074\u0053\u007a"
	case ST_ConstraintTypePyraAcctRatio:
		_cbbgf.Value = "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0052\u0061\u0074\u0069\u006f"
	case ST_ConstraintTypeSecFontSz:
		_cbbgf.Value = "\u0073e\u0063\u0046\u006f\u006e\u0074\u0053z"
	case ST_ConstraintTypeSibSp:
		_cbbgf.Value = "\u0073\u0069\u0062S\u0070"
	case ST_ConstraintTypeSecSibSp:
		_cbbgf.Value = "\u0073\u0065\u0063\u0053\u0069\u0062\u0053\u0070"
	case ST_ConstraintTypeSp:
		_cbbgf.Value = "\u0073\u0070"
	case ST_ConstraintTypeStemThick:
		_cbbgf.Value = "\u0073t\u0065\u006d\u0054\u0068\u0069\u0063k"
	case ST_ConstraintTypeT:
		_cbbgf.Value = "\u0074"
	case ST_ConstraintTypeTMarg:
		_cbbgf.Value = "\u0074\u004d\u0061r\u0067"
	case ST_ConstraintTypeTOff:
		_cbbgf.Value = "\u0074\u004f\u0066\u0066"
	case ST_ConstraintTypeUserA:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0041"
	case ST_ConstraintTypeUserB:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0042"
	case ST_ConstraintTypeUserC:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0043"
	case ST_ConstraintTypeUserD:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0044"
	case ST_ConstraintTypeUserE:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0045"
	case ST_ConstraintTypeUserF:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0046"
	case ST_ConstraintTypeUserG:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0047"
	case ST_ConstraintTypeUserH:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0048"
	case ST_ConstraintTypeUserI:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0049"
	case ST_ConstraintTypeUserJ:
		_cbbgf.Value = "\u0075\u0073\u0065r\u004a"
	case ST_ConstraintTypeUserK:
		_cbbgf.Value = "\u0075\u0073\u0065r\u004b"
	case ST_ConstraintTypeUserL:
		_cbbgf.Value = "\u0075\u0073\u0065r\u004c"
	case ST_ConstraintTypeUserM:
		_cbbgf.Value = "\u0075\u0073\u0065r\u004d"
	case ST_ConstraintTypeUserN:
		_cbbgf.Value = "\u0075\u0073\u0065r\u004e"
	case ST_ConstraintTypeUserO:
		_cbbgf.Value = "\u0075\u0073\u0065r\u004f"
	case ST_ConstraintTypeUserP:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0050"
	case ST_ConstraintTypeUserQ:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0051"
	case ST_ConstraintTypeUserR:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0052"
	case ST_ConstraintTypeUserS:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0053"
	case ST_ConstraintTypeUserT:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0054"
	case ST_ConstraintTypeUserU:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0055"
	case ST_ConstraintTypeUserV:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0056"
	case ST_ConstraintTypeUserW:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0057"
	case ST_ConstraintTypeUserX:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0058"
	case ST_ConstraintTypeUserY:
		_cbbgf.Value = "\u0075\u0073\u0065r\u0059"
	case ST_ConstraintTypeUserZ:
		_cbbgf.Value = "\u0075\u0073\u0065r\u005a"
	case ST_ConstraintTypeW:
		_cbbgf.Value = "\u0077"
	case ST_ConstraintTypeWArH:
		_cbbgf.Value = "\u0077\u0041\u0072\u0048"
	case ST_ConstraintTypeWOff:
		_cbbgf.Value = "\u0077\u004f\u0066\u0066"
	}
	return _cbbgf, nil
}

func (_dgab *CT_LayoutVariablePropertySet) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_bagf:
	for {
		_fgedb, _dgdcc := d.Token()
		if _dgdcc != nil {
			return _dgdcc
		}
		switch _cfdc := _fgedb.(type) {
		case _a.StartElement:
			switch _cfdc.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074"}:
				_dgab.OrgChart = NewCT_OrgChart()
				if _abddc := d.DecodeElement(_dgab.OrgChart, &_cfdc); _abddc != nil {
					return _abddc
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u004da\u0078"}:
				_dgab.ChMax = NewCT_ChildMax()
				if _dgdge := d.DecodeElement(_dgab.ChMax, &_cfdc); _dgdge != nil {
					return _dgdge
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u0050\u0072\u0065\u0066"}:
				_dgab.ChPref = NewCT_ChildPref()
				if _fagg := d.DecodeElement(_dgab.ChPref, &_cfdc); _fagg != nil {
					return _fagg
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0062\u0075\u006c\u006c\u0065\u0074\u0045\u006e\u0061\u0062\u006c\u0065\u0064"}:
				_dgab.BulletEnabled = NewCT_BulletEnabled()
				if _bgbg := d.DecodeElement(_dgab.BulletEnabled, &_cfdc); _bgbg != nil {
					return _bgbg
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0069\u0072"}:
				_dgab.Dir = NewCT_Direction()
				if _egab := d.DecodeElement(_dgab.Dir, &_cfdc); _egab != nil {
					return _egab
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"}:
				_dgab.HierBranch = NewCT_HierBranchStyle()
				if _fcdc := d.DecodeElement(_dgab.HierBranch, &_cfdc); _fcdc != nil {
					return _fcdc
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061n\u0069\u006d\u004f\u006e\u0065"}:
				_dgab.AnimOne = NewCT_AnimOne()
				if _edag := d.DecodeElement(_dgab.AnimOne, &_cfdc); _edag != nil {
					return _edag
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061n\u0069\u006d\u004c\u0076\u006c"}:
				_dgab.AnimLvl = NewCT_AnimLvl()
				if _gade := d.DecodeElement(_dgab.AnimLvl, &_cfdc); _gade != nil {
					return _gade
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073"}:
				_dgab.ResizeHandles = NewCT_ResizeHandles()
				if _fgbea := d.DecodeElement(_dgab.ResizeHandles, &_cfdc); _fgbea != nil {
					return _fgbea
				}
			default:
				_b.Log.Debug("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020C\u0054\u005f\u004c\u0061\u0079\u006f\u0075\u0074\u0056\u0061\u0072\u0069\u0061\u0062\u006c\u0065P\u0072\u006fpe\u0072\u0074\u0079S\u0065\u0074\u0020\u0025\u0076", _cfdc.Name)
				if _gdfb := d.Skip(); _gdfb != nil {
					return _gdfb
				}
			}
		case _a.EndElement:
			break _bagf
		case _a.CharData:
		}
	}
	return nil
}

type CT_CTStyleLabel struct {
	NameAttr       string
	FillClrLst     *CT_Colors
	LinClrLst      *CT_Colors
	EffectClrLst   *CT_Colors
	TxLinClrLst    *CT_Colors
	TxFillClrLst   *CT_Colors
	TxEffectClrLst *CT_Colors
	ExtLst         *_bg.CT_OfficeArtExtensionList
}

func (_fega *ST_ParameterId) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_fega = 0
	case "\u0068o\u0072\u007a\u0041\u006c\u0069\u0067n":
		*_fega = 1
	case "\u0076e\u0072\u0074\u0041\u006c\u0069\u0067n":
		*_fega = 2
	case "\u0063\u0068\u0044i\u0072":
		*_fega = 3
	case "\u0063h\u0041\u006c\u0069\u0067\u006e":
		*_fega = 4
	case "\u0073\u0065\u0063\u0043\u0068\u0041\u006c\u0069\u0067\u006e":
		*_fega = 5
	case "\u006c\u0069\u006e\u0044\u0069\u0072":
		*_fega = 6
	case "\u0073e\u0063\u004c\u0069\u006e\u0044\u0069r":
		*_fega = 7
	case "\u0073\u0074\u0045\u006c\u0065\u006d":
		*_fega = 8
	case "\u0062\u0065\u006e\u0064\u0050\u0074":
		*_fega = 9
	case "\u0063\u006f\u006e\u006e\u0052\u006f\u0075\u0074":
		*_fega = 10
	case "\u0062\u0065\u0067\u0053\u0074\u0079":
		*_fega = 11
	case "\u0065\u006e\u0064\u0053\u0074\u0079":
		*_fega = 12
	case "\u0064\u0069\u006d":
		*_fega = 13
	case "\u0072o\u0074\u0050\u0061\u0074\u0068":
		*_fega = 14
	case "\u0063t\u0072\u0053\u0068\u0070\u004d\u0061p":
		*_fega = 15
	case "\u006e\u006f\u0064\u0065\u0048\u006f\u0072\u007a\u0041\u006c\u0069\u0067\u006e":
		*_fega = 16
	case "\u006e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0041\u006c\u0069\u0067\u006e":
		*_fega = 17
	case "\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b":
		*_fega = 18
	case "\u0074\u0078\u0044i\u0072":
		*_fega = 19
	case "p\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0050\u006f\u0073":
		*_fega = 20
	case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054\u0078\u004d\u0061\u0072":
		*_fega = 21
	case "\u0074x\u0042\u006c\u0044\u0069\u0072":
		*_fega = 22
	case "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u007a":
		*_fega = 23
	case "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0056\u0065\u0072\u0074":
		*_fega = 24
	case "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0048o\u0072\u007a\u0043\u0068":
		*_fega = 25
	case "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0043\u0068":
		*_fega = 26
	case "\u0070\u0061\u0072\u0054\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e":
		*_fega = 27
	case "\u0070\u0061\u0072\u0054\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e":
		*_fega = 28
	case "\u0073h\u0070T\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e\u0043\u0068":
		*_fega = 29
	case "\u0073h\u0070T\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e\u0043\u0068":
		*_fega = 30
	case "\u0061u\u0074\u006f\u0054\u0078\u0052\u006ft":
		*_fega = 31
	case "\u0067\u0072\u0044i\u0072":
		*_fega = 32
	case "\u0066l\u006f\u0077\u0044\u0069\u0072":
		*_fega = 33
	case "\u0063o\u006e\u0074\u0044\u0069\u0072":
		*_fega = 34
	case "\u0062\u006b\u0070\u0074":
		*_fega = 35
	case "\u006f\u0066\u0066":
		*_fega = 36
	case "\u0068i\u0065\u0072\u0041\u006c\u0069\u0067n":
		*_fega = 37
	case "\u0062\u006b\u0050t\u0046\u0069\u0078\u0065\u0064\u0056\u0061\u006c":
		*_fega = 38
	case "s\u0074\u0042\u0075\u006c\u006c\u0065\u0074\u004c\u0076\u006c":
		*_fega = 39
	case "\u0073\u0074\u0041n\u0067":
		*_fega = 40
	case "\u0073p\u0061\u006e\u0041\u006e\u0067":
		*_fega = 41
	case "\u0061\u0072":
		*_fega = 42
	case "\u006cn\u0053\u0070\u0050\u0061\u0072":
		*_fega = 43
	case "\u006c\u006e\u0053\u0070\u0041\u0066\u0050\u0061\u0072\u0050":
		*_fega = 44
	case "\u006c\u006e\u0053\u0070\u0043\u0068":
		*_fega = 45
	case "\u006cn\u0053\u0070\u0041\u0066\u0043\u0068P":
		*_fega = 46
	case "r\u0074\u0053\u0068\u006f\u0072\u0074\u0044\u0069\u0073\u0074":
		*_fega = 47
	case "\u0061l\u0069\u0067\u006e\u0054\u0078":
		*_fega = 48
	case "p\u0079\u0072\u0061\u004c\u0076\u006c\u004e\u006f\u0064\u0065":
		*_fega = 49
	case "\u0070\u0079r\u0061\u0041\u0063c\u0074\u0042\u006b\u0067\u0064\u004e\u006f\u0064\u0065":
		*_fega = 50
	case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054x\u004e\u006f\u0064\u0065":
		*_fega = 51
	case "\u0073r\u0063\u004e\u006f\u0064\u0065":
		*_fega = 52
	case "\u0064s\u0074\u004e\u006f\u0064\u0065":
		*_fega = 53
	case "\u0062\u0065\u0067\u0050\u0074\u0073":
		*_fega = 54
	case "\u0065\u006e\u0064\u0050\u0074\u0073":
		*_fega = 55
	}
	return nil
}

// Validate validates the StyleDefHdrLst and its children
func (_gcaa *StyleDefHdrLst) Validate() error {
	return _gcaa.ValidateWithPath("\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048d\u0072\u004c\u0073\u0074")
}

// ValidateWithPath validates the CT_DataModel and its children, prefixing error messages with path
func (_fbfg *CT_DataModel) ValidateWithPath(path string) error {
	if _bceg := _fbfg.PtLst.ValidateWithPath(path + "\u002f\u0050\u0074\u004c\u0073\u0074"); _bceg != nil {
		return _bceg
	}
	if _fbfg.CxnLst != nil {
		if _bfc := _fbfg.CxnLst.ValidateWithPath(path + "\u002fC\u0078\u006e\u004c\u0073\u0074"); _bfc != nil {
			return _bfc
		}
	}
	if _fbfg.Bg != nil {
		if _bad := _fbfg.Bg.ValidateWithPath(path + "\u002f\u0042\u0067"); _bad != nil {
			return _bad
		}
	}
	if _fbfg.Whole != nil {
		if _dgge := _fbfg.Whole.ValidateWithPath(path + "\u002f\u0057\u0068\u006f\u006c\u0065"); _dgge != nil {
			return _dgge
		}
	}
	if _fbfg.ExtLst != nil {
		if _gfeg := _fbfg.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _gfeg != nil {
			return _gfeg
		}
	}
	return nil
}

type CT_TextProps struct {
	Sp3d   *_bg.CT_Shape3D
	FlatTx *_bg.CT_FlatText
}

func (_dfaf ST_StartingElement) ValidateWithPath(path string) error {
	switch _dfaf {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dfaf))
	}
	return nil
}

func (_ggcbbb *ST_ConstraintRelationship) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ggcbbb = 0
	case "\u0073\u0065\u006c\u0066":
		*_ggcbbb = 1
	case "\u0063\u0068":
		*_ggcbbb = 2
	case "\u0064\u0065\u0073":
		*_ggcbbb = 3
	}
	return nil
}

// Validate validates the LayoutDefHdr and its children
func (_cadgf *LayoutDefHdr) Validate() error {
	return _cadgf.ValidateWithPath("\u004c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072")
}

type CT_SDCategory struct {
	TypeAttr string
	PriAttr  uint32
}

// Validate validates the CT_Adj and its children
func (_deb *CT_Adj) Validate() error {
	return _deb.ValidateWithPath("\u0043\u0054\u005f\u0041\u0064\u006a")
}

func (_adbb *CT_Parameter) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_adbb.TypeAttr = ST_ParameterId(1)
	for _, _abfd := range start.Attr {
		if _abfd.Name.Local == "\u0074\u0079\u0070\u0065" {
			_adbb.TypeAttr.UnmarshalXMLAttr(_abfd)
			continue
		}
		if _abfd.Name.Local == "\u0076\u0061\u006c" {
			_begd, _caeb := ParseUnionST_ParameterVal(_abfd.Value)
			if _caeb != nil {
				return _caeb
			}
			_adbb.ValAttr = _begd
			continue
		}
	}
	for {
		_gffg, _beffb := d.Token()
		if _beffb != nil {
			return _f.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005f\u0050\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u003a\u0020\u0025\u0073", _beffb)
		}
		if _egge, _cdfa := _gffg.(_a.EndElement); _cdfa && _egge.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ddgf ST_ContinueDirection) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_ddgf.String(), start)
}

type CT_Algorithm struct {
	TypeAttr ST_AlgorithmType
	RevAttr  *uint32
	Param    []*CT_Parameter
	ExtLst   *_bg.CT_OfficeArtExtensionList
}

func (_gaab ST_ChildDirection) String() string {
	switch _gaab {
	case 0:
		return ""
	case 1:
		return "\u0068\u006f\u0072\u007a"
	case 2:
		return "\u0076\u0065\u0072\u0074"
	}
	return ""
}

func NewCT_Otherwise() *CT_Otherwise { _gegbb := &CT_Otherwise{}; return _gegbb }

func (_faeff ST_TextAnchorVertical) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_eeafg := _a.Attr{}
	_eeafg.Name = name
	switch _faeff {
	case ST_TextAnchorVerticalUnset:
		_eeafg.Value = ""
	case ST_TextAnchorVerticalT:
		_eeafg.Value = "\u0074"
	case ST_TextAnchorVerticalMid:
		_eeafg.Value = "\u006d\u0069\u0064"
	case ST_TextAnchorVerticalB:
		_eeafg.Value = "\u0062"
	}
	return _eeafg, nil
}

func (_bccbe ST_TextAnchorHorizontal) Validate() error { return _bccbe.ValidateWithPath("") }

type ST_LinearDirection byte

type ST_PyramidAccentPosition byte

type ST_FunctionType byte

func (_feddf ST_VariableType) ValidateWithPath(path string) error {
	switch _feddf {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_feddf))
	}
	return nil
}

// Validate validates the CT_DiagramDefinitionHeaderLst and its children
func (_bgda *CT_DiagramDefinitionHeaderLst) Validate() error {
	return _bgda.ValidateWithPath("\u0043\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065rL\u0073\u0074")
}

func (_afcfb ST_PrSetCustVal) String() string {
	if _afcfb.ST_Percentage != nil {
		return _f.Sprintf("\u0025\u0076", *_afcfb.ST_Percentage)
	}
	if _afcfb.Int32 != nil {
		return _f.Sprintf("\u0025\u0076", *_afcfb.Int32)
	}
	return ""
}

func NewCT_AnimLvl() *CT_AnimLvl { _feb := &CT_AnimLvl{}; return _feb }

type CT_Otherwise struct {
	NameAttr   *string
	Alg        []*CT_Algorithm
	Shape      []*CT_Shape
	PresOf     []*CT_PresentationOf
	ConstrLst  []*CT_Constraints
	RuleLst    []*CT_Rules
	ForEach    []*CT_ForEach
	LayoutNode []*CT_LayoutNode
	Choose     []*CT_Choose
	ExtLst     []*_bg.CT_OfficeArtExtensionList
}

const (
	ST_BoolOperatorUnset ST_BoolOperator = 0
	ST_BoolOperatorNone  ST_BoolOperator = 1
	ST_BoolOperatorEqu   ST_BoolOperator = 2
	ST_BoolOperatorGte   ST_BoolOperator = 3
	ST_BoolOperatorLte   ST_BoolOperator = 4
)

func (_ggdd ST_NodeHorizontalAlignment) ValidateWithPath(path string) error {
	switch _ggdd {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ggdd))
	}
	return nil
}

func (_gdbfb ST_FunctionValue) String() string {
	if _gdbfb.Int32 != nil {
		return _f.Sprintf("\u0025\u0076", *_gdbfb.Int32)
	}
	if _gdbfb.Bool != nil {
		return _f.Sprintf("\u0025\u0076", *_gdbfb.Bool)
	}
	if _gdbfb.ST_Direction != ST_DirectionUnset {
		return _gdbfb.ST_Direction.String()
	}
	if _gdbfb.ST_HierBranchStyle != ST_HierBranchStyleUnset {
		return _gdbfb.ST_HierBranchStyle.String()
	}
	if _gdbfb.ST_AnimOneStr != ST_AnimOneStrUnset {
		return _gdbfb.ST_AnimOneStr.String()
	}
	if _gdbfb.ST_AnimLvlStr != ST_AnimLvlStrUnset {
		return _gdbfb.ST_AnimLvlStr.String()
	}
	if _gdbfb.ST_ResizeHandlesStr != ST_ResizeHandlesStrUnset {
		return _gdbfb.ST_ResizeHandlesStr.String()
	}
	return ""
}

func (_fege *AG_IteratorAttributes) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _ea := range start.Attr {
		if _ea.Name.Local == "\u0061\u0078\u0069\u0073" {
			_fb, _db := ParseSliceST_AxisTypes(_ea.Value)
			if _db != nil {
				return _db
			}
			_fege.AxisAttr = &_fb
			continue
		}
		if _ea.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_bgc, _fg := ParseSliceST_ElementTypes(_ea.Value)
			if _fg != nil {
				return _fg
			}
			_fege.PtTypeAttr = &_bgc
			continue
		}
		if _ea.Name.Local == "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073" {
			_eac, _dbe := ParseSliceST_Booleans(_ea.Value)
			if _dbe != nil {
				return _dbe
			}
			_fege.HideLastTransAttr = &_eac
			continue
		}
		if _ea.Name.Local == "\u0073\u0074" {
			_faf, _cd := ParseSliceST_Ints(_ea.Value)
			if _cd != nil {
				return _cd
			}
			_fege.StAttr = &_faf
			continue
		}
		if _ea.Name.Local == "\u0063\u006e\u0074" {
			_dd, _eec := ParseSliceST_UnsignedInts(_ea.Value)
			if _eec != nil {
				return _eec
			}
			_fege.CntAttr = &_dd
			continue
		}
		if _ea.Name.Local == "\u0073\u0074\u0065\u0070" {
			_bgbf, _bd := ParseSliceST_Ints(_ea.Value)
			if _bd != nil {
				return _bd
			}
			_fege.StepAttr = &_bgbf
			continue
		}
	}
	for {
		_ebde, _fdb := d.Token()
		if _fdb != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073\u0069\u006eg\u0020\u0041\u0047\u005f\u0049\u0074\u0065\u0072\u0061\u0074\u006f\u0072\u0041t\u0074\u0072\u0069\u0062\u0075\u0074\u0065s\u003a\u0020\u0025\u0073", _fdb)
		}
		if _fde, _dfb := _ebde.(_a.EndElement); _dfb && _fde.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_DiagramDefinitionHeader and its children, prefixing error messages with path
func (_bced *CT_DiagramDefinitionHeader) ValidateWithPath(path string) error {
	for _dddd, _aded := range _bced.Title {
		if _gcec := _aded.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _dddd)); _gcec != nil {
			return _gcec
		}
	}
	for _cdce, _feda := range _bced.Desc {
		if _dcdf := _feda.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _cdce)); _dcdf != nil {
			return _dcdf
		}
	}
	if _bced.CatLst != nil {
		if _cgb := _bced.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _cgb != nil {
			return _cgb
		}
	}
	if _bced.ExtLst != nil {
		if _dba := _bced.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _dba != nil {
			return _dba
		}
	}
	return nil
}

func (_dcbd ST_ConstraintRelationship) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_eaga := _a.Attr{}
	_eaga.Name = name
	switch _dcbd {
	case ST_ConstraintRelationshipUnset:
		_eaga.Value = ""
	case ST_ConstraintRelationshipSelf:
		_eaga.Value = "\u0073\u0065\u006c\u0066"
	case ST_ConstraintRelationshipCh:
		_eaga.Value = "\u0063\u0068"
	case ST_ConstraintRelationshipDes:
		_eaga.Value = "\u0064\u0065\u0073"
	}
	return _eaga, nil
}

// Validate validates the CT_SDDescription and its children
func (_ecgad *CT_SDDescription) Validate() error {
	return _ecgad.ValidateWithPath("\u0043\u0054_\u0053\u0044\u0044e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e")
}

func (_edgf *ST_FunctionOperator) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_edgf = 0
	case "\u0065\u0071\u0075":
		*_edgf = 1
	case "\u006e\u0065\u0071":
		*_edgf = 2
	case "\u0067\u0074":
		*_edgf = 3
	case "\u006c\u0074":
		*_edgf = 4
	case "\u0067\u0074\u0065":
		*_edgf = 5
	case "\u006c\u0074\u0065":
		*_edgf = 6
	}
	return nil
}

func (_cgdag *ST_CenterShapeMapping) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_dbbbb, _baged := d.Token()
	if _baged != nil {
		return _baged
	}
	if _fbeba, _dege := _dbbbb.(_a.EndElement); _dege && _fbeba.Name == start.Name {
		*_cgdag = 1
		return nil
	}
	if _gbca, _egbg := _dbbbb.(_a.CharData); !_egbg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dbbbb)
	} else {
		switch string(_gbca) {
		case "":
			*_cgdag = 0
		case "\u006e\u006f\u006e\u0065":
			*_cgdag = 1
		case "\u0066\u004e\u006fd\u0065":
			*_cgdag = 2
		}
	}
	_dbbbb, _baged = d.Token()
	if _baged != nil {
		return _baged
	}
	if _bdabd, _cegab := _dbbbb.(_a.EndElement); _cegab && _bdabd.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dbbbb)
}

// Validate validates the ColorsDefHdr and its children
func (_ecce *ColorsDefHdr) Validate() error {
	return _ecce.ValidateWithPath("\u0043\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072")
}

func (_cdag ST_SecondaryLinearDirection) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_ddddf := _a.Attr{}
	_ddddf.Name = name
	switch _cdag {
	case ST_SecondaryLinearDirectionUnset:
		_ddddf.Value = ""
	case ST_SecondaryLinearDirectionNone:
		_ddddf.Value = "\u006e\u006f\u006e\u0065"
	case ST_SecondaryLinearDirectionFromL:
		_ddddf.Value = "\u0066\u0072\u006fm\u004c"
	case ST_SecondaryLinearDirectionFromR:
		_ddddf.Value = "\u0066\u0072\u006fm\u0052"
	case ST_SecondaryLinearDirectionFromT:
		_ddddf.Value = "\u0066\u0072\u006fm\u0054"
	case ST_SecondaryLinearDirectionFromB:
		_ddddf.Value = "\u0066\u0072\u006fm\u0042"
	}
	return _ddddf, nil
}

type ST_StartingElement byte

// ValidateWithPath validates the CT_CTStyleLabel and its children, prefixing error messages with path
func (_ffe *CT_CTStyleLabel) ValidateWithPath(path string) error {
	if _ffe.FillClrLst != nil {
		if _cfe := _ffe.FillClrLst.ValidateWithPath(path + "/\u0046\u0069\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"); _cfe != nil {
			return _cfe
		}
	}
	if _ffe.LinClrLst != nil {
		if _ebgg := _ffe.LinClrLst.ValidateWithPath(path + "\u002f\u004c\u0069\u006e\u0043\u006c\u0072\u004c\u0073\u0074"); _ebgg != nil {
			return _ebgg
		}
	}
	if _ffe.EffectClrLst != nil {
		if _ade := _ffe.EffectClrLst.ValidateWithPath(path + "\u002f\u0045\u0066\u0066\u0065\u0063\u0074\u0043\u006c\u0072\u004c\u0073\u0074"); _ade != nil {
			return _ade
		}
	}
	if _ffe.TxLinClrLst != nil {
		if _cgf := _ffe.TxLinClrLst.ValidateWithPath(path + "\u002f\u0054\u0078L\u0069\u006e\u0043\u006c\u0072\u004c\u0073\u0074"); _cgf != nil {
			return _cgf
		}
	}
	if _ffe.TxFillClrLst != nil {
		if _bbba := _ffe.TxFillClrLst.ValidateWithPath(path + "\u002f\u0054\u0078\u0046\u0069\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"); _bbba != nil {
			return _bbba
		}
	}
	if _ffe.TxEffectClrLst != nil {
		if _cae := _ffe.TxEffectClrLst.ValidateWithPath(path + "\u002fT\u0078E\u0066\u0066\u0065\u0063\u0074\u0043\u006c\u0072\u004c\u0073\u0074"); _cae != nil {
			return _cae
		}
	}
	if _ffe.ExtLst != nil {
		if _cgd := _ffe.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _cgd != nil {
			return _cgd
		}
	}
	return nil
}

// ST_ParameterVal is a union type
type ST_ParameterVal struct {
	ST_DiagramHorizontalAlignment ST_DiagramHorizontalAlignment
	ST_VerticalAlignment          ST_VerticalAlignment
	ST_ChildDirection             ST_ChildDirection
	ST_ChildAlignment             ST_ChildAlignment
	ST_SecondaryChildAlignment    ST_SecondaryChildAlignment
	ST_LinearDirection            ST_LinearDirection
	ST_SecondaryLinearDirection   ST_SecondaryLinearDirection
	ST_StartingElement            ST_StartingElement
	ST_BendPoint                  ST_BendPoint
	ST_ConnectorRouting           ST_ConnectorRouting
	ST_ArrowheadStyle             ST_ArrowheadStyle
	ST_ConnectorDimension         ST_ConnectorDimension
	ST_RotationPath               ST_RotationPath
	ST_CenterShapeMapping         ST_CenterShapeMapping
	ST_NodeHorizontalAlignment    ST_NodeHorizontalAlignment
	ST_NodeVerticalAlignment      ST_NodeVerticalAlignment
	ST_FallbackDimension          ST_FallbackDimension
	ST_TextDirection              ST_TextDirection
	ST_PyramidAccentPosition      ST_PyramidAccentPosition
	ST_PyramidAccentTextMargin    ST_PyramidAccentTextMargin
	ST_TextBlockDirection         ST_TextBlockDirection
	ST_TextAnchorHorizontal       ST_TextAnchorHorizontal
	ST_TextAnchorVertical         ST_TextAnchorVertical
	ST_DiagramTextAlignment       ST_DiagramTextAlignment
	ST_AutoTextRotation           ST_AutoTextRotation
	ST_GrowDirection              ST_GrowDirection
	ST_FlowDirection              ST_FlowDirection
	ST_ContinueDirection          ST_ContinueDirection
	ST_Breakpoint                 ST_Breakpoint
	ST_Offset                     ST_Offset
	ST_HierarchyAlignment         ST_HierarchyAlignment
	Int32                         *int32
	Float64                       *float64
	Bool                          *bool
	StringVal                     *string
	ST_ConnectorPoint             ST_ConnectorPoint
}

func (_dgaaa *ST_PtType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_dgaaa = 0
	case "\u006e\u006f\u0064\u0065":
		*_dgaaa = 1
	case "\u0061\u0073\u0073\u0074":
		*_dgaaa = 2
	case "\u0064\u006f\u0063":
		*_dgaaa = 3
	case "\u0070\u0072\u0065\u0073":
		*_dgaaa = 4
	case "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073":
		*_dgaaa = 5
	case "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073":
		*_dgaaa = 6
	}
	return nil
}

func (_degbf ST_ParameterId) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_degbf.String(), start)
}

func (_dgce *ST_VerticalAlignment) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_dgce = 0
	case "\u0074":
		*_dgce = 1
	case "\u006d\u0069\u0064":
		*_dgce = 2
	case "\u0062":
		*_dgce = 3
	case "\u006e\u006f\u006e\u0065":
		*_dgce = 4
	}
	return nil
}

func (_gabge *CT_When) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gabge.FuncAttr = ST_FunctionType(1)
	_gabge.OpAttr = ST_FunctionOperator(1)
	for _, _fdfd := range start.Attr {
		if _fdfd.Name.Local == "\u006e\u0061\u006d\u0065" {
			_cgcac, _bgbc := _fdfd.Value, error(nil)
			if _bgbc != nil {
				return _bgbc
			}
			_gabge.NameAttr = &_cgcac
			continue
		}
		if _fdfd.Name.Local == "\u0061\u0072\u0067" {
			_bbeef, _cagb := ParseUnionST_FunctionArgument(_fdfd.Value)
			if _cagb != nil {
				return _cagb
			}
			_gabge.ArgAttr = &_bbeef
			continue
		}
		if _fdfd.Name.Local == "\u0076\u0061\u006c" {
			_begb, _eegb := ParseUnionST_FunctionValue(_fdfd.Value)
			if _eegb != nil {
				return _eegb
			}
			_gabge.ValAttr = _begb
			continue
		}
		if _fdfd.Name.Local == "\u0066\u0075\u006e\u0063" {
			_gabge.FuncAttr.UnmarshalXMLAttr(_fdfd)
			continue
		}
		if _fdfd.Name.Local == "\u006f\u0070" {
			_gabge.OpAttr.UnmarshalXMLAttr(_fdfd)
			continue
		}
		if _fdfd.Name.Local == "\u0061\u0078\u0069\u0073" {
			_dcfac, _dbgb := ParseSliceST_AxisTypes(_fdfd.Value)
			if _dbgb != nil {
				return _dbgb
			}
			_gabge.AxisAttr = &_dcfac
			continue
		}
		if _fdfd.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_ccfea, _bddeg := ParseSliceST_ElementTypes(_fdfd.Value)
			if _bddeg != nil {
				return _bddeg
			}
			_gabge.PtTypeAttr = &_ccfea
			continue
		}
		if _fdfd.Name.Local == "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073" {
			_cefdf, _gegee := ParseSliceST_Booleans(_fdfd.Value)
			if _gegee != nil {
				return _gegee
			}
			_gabge.HideLastTransAttr = &_cefdf
			continue
		}
		if _fdfd.Name.Local == "\u0073\u0074" {
			_edcg, _dgece := ParseSliceST_Ints(_fdfd.Value)
			if _dgece != nil {
				return _dgece
			}
			_gabge.StAttr = &_edcg
			continue
		}
		if _fdfd.Name.Local == "\u0063\u006e\u0074" {
			_ddda, _gfcf := ParseSliceST_UnsignedInts(_fdfd.Value)
			if _gfcf != nil {
				return _gfcf
			}
			_gabge.CntAttr = &_ddda
			continue
		}
		if _fdfd.Name.Local == "\u0073\u0074\u0065\u0070" {
			_ffgf, _dbeg := ParseSliceST_Ints(_fdfd.Value)
			if _dbeg != nil {
				return _dbeg
			}
			_gabge.StepAttr = &_ffgf
			continue
		}
	}
_gdgfe:
	for {
		_dcgg, _gcdf := d.Token()
		if _gcdf != nil {
			return _gcdf
		}
		switch _ebdga := _dcgg.(type) {
		case _a.StartElement:
			switch _ebdga.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u006c\u0067"}:
				_cdaec := NewCT_Algorithm()
				if _fbda := d.DecodeElement(_cdaec, &_ebdga); _fbda != nil {
					return _fbda
				}
				_gabge.Alg = append(_gabge.Alg, _cdaec)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0068\u0061p\u0065"}:
				_cbgaa := NewCT_Shape()
				if _eceae := d.DecodeElement(_cbgaa, &_ebdga); _eceae != nil {
					return _eceae
				}
				_gabge.Shape = append(_gabge.Shape, _cbgaa)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}:
				_ggcbb := NewCT_PresentationOf()
				if _gdgcg := d.DecodeElement(_ggcbb, &_ebdga); _gdgcg != nil {
					return _gdgcg
				}
				_gabge.PresOf = append(_gabge.PresOf, _ggcbb)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}:
				_efdf := NewCT_Constraints()
				if _ebeege := d.DecodeElement(_efdf, &_ebdga); _ebeege != nil {
					return _ebeege
				}
				_gabge.ConstrLst = append(_gabge.ConstrLst, _efdf)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}:
				_gaeac := NewCT_Rules()
				if _gabag := d.DecodeElement(_gaeac, &_ebdga); _gabag != nil {
					return _gabag
				}
				_gabge.RuleLst = append(_gabge.RuleLst, _gaeac)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}:
				_fcfff := NewCT_ForEach()
				if _fgeae := d.DecodeElement(_fcfff, &_ebdga); _fgeae != nil {
					return _fgeae
				}
				_gabge.ForEach = append(_gabge.ForEach, _fcfff)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				_ffdbd := NewCT_LayoutNode()
				if _gcefa := d.DecodeElement(_ffdbd, &_ebdga); _gcefa != nil {
					return _gcefa
				}
				_gabge.LayoutNode = append(_gabge.LayoutNode, _ffdbd)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}:
				_ecbba := NewCT_Choose()
				if _edgb := d.DecodeElement(_ecbba, &_ebdga); _edgb != nil {
					return _edgb
				}
				_gabge.Choose = append(_gabge.Choose, _ecbba)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_begga := _bg.NewCT_OfficeArtExtensionList()
				if _acdfb := d.DecodeElement(_begga, &_ebdga); _acdfb != nil {
					return _acdfb
				}
				_gabge.ExtLst = append(_gabge.ExtLst, _begga)
			default:
				_b.Log.Debug("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0057\u0068\u0065\u006e\u0020\u0025\u0076", _ebdga.Name)
				if _dcgff := d.Skip(); _dcgff != nil {
					return _dcgff
				}
			}
		case _a.EndElement:
			break _gdgfe
		case _a.CharData:
		}
	}
	return nil
}

func (_agfdg *CT_PtList) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_efdc:
	for {
		_ecgc, _fccec := d.Token()
		if _fccec != nil {
			return _fccec
		}
		switch _gfdgd := _ecgc.(type) {
		case _a.StartElement:
			switch _gfdgd.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0074"}:
				_aaaa := NewCT_Pt()
				if _gbcg := d.DecodeElement(_aaaa, &_gfdgd); _gbcg != nil {
					return _gbcg
				}
				_agfdg.Pt = append(_agfdg.Pt, _aaaa)
			default:
				_b.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0050\u0074\u004ci\u0073\u0074 \u0025\u0076", _gfdgd.Name)
				if _edbe := d.Skip(); _edbe != nil {
					return _edbe
				}
			}
		case _a.EndElement:
			break _efdc
		case _a.CharData:
		}
	}
	return nil
}

func (_fac *CT_BulletEnabled) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _fac.ValAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_fac.ValAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

type CT_CTName struct {
	LangAttr *string
	ValAttr  string
}

const (
	ST_DirectionUnset ST_Direction = 0
	ST_DirectionNorm  ST_Direction = 1
	ST_DirectionRev   ST_Direction = 2
)

type CT_StyleDefinitionHeader struct {
	UniqueIdAttr string
	MinVerAttr   *string
	ResIdAttr    *int32
	Title        []*CT_SDName
	Desc         []*CT_SDDescription
	CatLst       *CT_SDCategories
	ExtLst       *_bg.CT_OfficeArtExtensionList
}

func (_gcgf *CT_Description) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _gcgf.LangAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _f.Sprintf("\u0025\u0076", *_gcgf.LangAttr)})
	}
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", _gcgf.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_feff ST_ModelId) String() string {
	if _feff.Int32 != nil {
		return _f.Sprintf("\u0025\u0076", *_feff.Int32)
	}
	if _feff.ST_Guid != nil {
		return _f.Sprintf("\u0025\u0076", *_feff.ST_Guid)
	}
	return ""
}

func (_daecf *ST_Offset) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gfgg, _cecgd := d.Token()
	if _cecgd != nil {
		return _cecgd
	}
	if _befa, _fggea := _gfgg.(_a.EndElement); _fggea && _befa.Name == start.Name {
		*_daecf = 1
		return nil
	}
	if _edgdd, _beaf := _gfgg.(_a.CharData); !_beaf {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gfgg)
	} else {
		switch string(_edgdd) {
		case "":
			*_daecf = 0
		case "\u0063\u0074\u0072":
			*_daecf = 1
		case "\u006f\u0066\u0066":
			*_daecf = 2
		}
	}
	_gfgg, _cecgd = d.Token()
	if _cecgd != nil {
		return _cecgd
	}
	if _fdgcc, _geba := _gfgg.(_a.EndElement); _geba && _fdgcc.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gfgg)
}

func (_dcff *ST_ChildAlignment) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_accga, _edcfc := d.Token()
	if _edcfc != nil {
		return _edcfc
	}
	if _fegcc, _ecgdb := _accga.(_a.EndElement); _ecgdb && _fegcc.Name == start.Name {
		*_dcff = 1
		return nil
	}
	if _dddeb, _cebf := _accga.(_a.CharData); !_cebf {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _accga)
	} else {
		switch string(_dddeb) {
		case "":
			*_dcff = 0
		case "\u0074":
			*_dcff = 1
		case "\u0062":
			*_dcff = 2
		case "\u006c":
			*_dcff = 3
		case "\u0072":
			*_dcff = 4
		}
	}
	_accga, _edcfc = d.Token()
	if _edcfc != nil {
		return _edcfc
	}
	if _bdgb, _dbag := _accga.(_a.EndElement); _dbag && _bdgb.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _accga)
}

func (_fdcgf ST_GrowDirection) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_fdcgf.String(), start)
}

type ST_PtType byte

// ST_FunctionValue is a union type
type ST_FunctionValue struct {
	Int32               *int32
	Bool                *bool
	ST_Direction        ST_Direction
	ST_HierBranchStyle  ST_HierBranchStyle
	ST_AnimOneStr       ST_AnimOneStr
	ST_AnimLvlStr       ST_AnimLvlStr
	ST_ResizeHandlesStr ST_ResizeHandlesStr
}

func (_cadd *CT_ResizeHandles) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _ebfd := range start.Attr {
		if _ebfd.Name.Local == "\u0076\u0061\u006c" {
			_cadd.ValAttr.UnmarshalXMLAttr(_ebfd)
			continue
		}
	}
	for {
		_gbcf, _addd := d.Token()
		if _addd != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0043\u0054\u005f\u0052\u0065\u0073\u0069z\u0065H\u0061\u006e\u0064\u006c\u0065\u0073\u003a \u0025\u0073", _addd)
		}
		if _efda, _dcad := _gbcf.(_a.EndElement); _dcad && _efda.Name == start.Name {
			break
		}
	}
	return nil
}

func (_cddg *CT_ColorTransformHeaderLst) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _cddg.ColorsDefHdr != nil {
		_agca := _a.StartElement{Name: _a.Name{Local: "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072"}}
		for _, _dadbb := range _cddg.ColorsDefHdr {
			e.EncodeElement(_dadbb, _agca)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_fbabdg ST_HierBranchStyle) ValidateWithPath(path string) error {
	switch _fbabdg {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fbabdg))
	}
	return nil
}

func (_gegb *CT_HierBranchStyle) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _bddg := range start.Attr {
		if _bddg.Name.Local == "\u0076\u0061\u006c" {
			_gegb.ValAttr.UnmarshalXMLAttr(_bddg)
			continue
		}
	}
	for {
		_abg, _fgeda := d.Token()
		if _fgeda != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054_\u0048\u0069\u0065\u0072\u0042\u0072\u0061n\u0063\u0068\u0053\u0074\u0079\u006c\u0065\u003a\u0020\u0025\u0073", _fgeda)
		}
		if _ffeeg, _bcce := _abg.(_a.EndElement); _bcce && _ffeeg.Name == start.Name {
			break
		}
	}
	return nil
}

type ST_ChildDirection byte

type ST_BoolOperator byte

func (_dddf ST_ModelId) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _dddf.Int32 != nil {
		e.EncodeToken(_a.CharData(_f.Sprintf("\u0025\u0064", *_dddf.Int32)))
	}
	if _dddf.ST_Guid != nil {
		e.EncodeToken(_a.CharData(*_dddf.ST_Guid))
	}
	return e.EncodeToken(_a.EndElement{Name: start.Name})
}

// Validate validates the CT_ChildPref and its children
func (_gcgb *CT_ChildPref) Validate() error {
	return _gcgb.ValidateWithPath("\u0043\u0054\u005fC\u0068\u0069\u006c\u0064\u0050\u0072\u0065\u0066")
}

func (_dgcae ST_PtType) String() string {
	switch _dgcae {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u0064\u0065"
	case 2:
		return "\u0061\u0073\u0073\u0074"
	case 3:
		return "\u0064\u006f\u0063"
	case 4:
		return "\u0070\u0072\u0065\u0073"
	case 5:
		return "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073"
	case 6:
		return "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073"
	}
	return ""
}

func ParseSliceST_ElementTypes(s string) (ST_ElementTypes, error) { return ST_ElementTypes{}, nil }

func (_ecab ST_PyramidAccentTextMargin) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_ecab.String(), start)
}

func (_efgfc ST_AutoTextRotation) ValidateWithPath(path string) error {
	switch _efgfc {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_efgfc))
	}
	return nil
}

// Validate validates the CT_AnimLvl and its children
func (_gge *CT_AnimLvl) Validate() error {
	return _gge.ValidateWithPath("\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004c\u0076\u006c")
}

// ValidateWithPath validates the CT_LayoutVariablePropertySet and its children, prefixing error messages with path
func (_eedd *CT_LayoutVariablePropertySet) ValidateWithPath(path string) error {
	if _eedd.OrgChart != nil {
		if _cebe := _eedd.OrgChart.ValidateWithPath(path + "\u002fO\u0072\u0067\u0043\u0068\u0061\u0072t"); _cebe != nil {
			return _cebe
		}
	}
	if _eedd.ChMax != nil {
		if _bcbb := _eedd.ChMax.ValidateWithPath(path + "\u002f\u0043\u0068\u004d\u0061\u0078"); _bcbb != nil {
			return _bcbb
		}
	}
	if _eedd.ChPref != nil {
		if _bcef := _eedd.ChPref.ValidateWithPath(path + "\u002fC\u0068\u0050\u0072\u0065\u0066"); _bcef != nil {
			return _bcef
		}
	}
	if _eedd.BulletEnabled != nil {
		if _eccbg := _eedd.BulletEnabled.ValidateWithPath(path + "\u002f\u0042\u0075\u006c\u006c\u0065\u0074\u0045\u006ea\u0062\u006c\u0065\u0064"); _eccbg != nil {
			return _eccbg
		}
	}
	if _eedd.Dir != nil {
		if _dffc := _eedd.Dir.ValidateWithPath(path + "\u002f\u0044\u0069\u0072"); _dffc != nil {
			return _dffc
		}
	}
	if _eedd.HierBranch != nil {
		if _afccd := _eedd.HierBranch.ValidateWithPath(path + "/\u0048\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"); _afccd != nil {
			return _afccd
		}
	}
	if _eedd.AnimOne != nil {
		if _cgba := _eedd.AnimOne.ValidateWithPath(path + "\u002f\u0041\u006e\u0069\u006d\u004f\u006e\u0065"); _cgba != nil {
			return _cgba
		}
	}
	if _eedd.AnimLvl != nil {
		if _fcabc := _eedd.AnimLvl.ValidateWithPath(path + "\u002f\u0041\u006e\u0069\u006d\u004c\u0076\u006c"); _fcabc != nil {
			return _fcabc
		}
	}
	if _eedd.ResizeHandles != nil {
		if _gbda := _eedd.ResizeHandles.ValidateWithPath(path + "\u002f\u0052\u0065\u0073\u0069\u007a\u0065\u0048\u0061n\u0064\u006c\u0065\u0073"); _gbda != nil {
			return _gbda
		}
	}
	return nil
}

func (_aaafa ST_FunctionType) String() string {
	switch _aaafa {
	case 0:
		return ""
	case 1:
		return "\u0063\u006e\u0074"
	case 2:
		return "\u0070\u006f\u0073"
	case 3:
		return "\u0072\u0065\u0076\u0050\u006f\u0073"
	case 4:
		return "\u0070o\u0073\u0045\u0076\u0065\u006e"
	case 5:
		return "\u0070\u006f\u0073\u004f\u0064\u0064"
	case 6:
		return "\u0076\u0061\u0072"
	case 7:
		return "\u0064\u0065\u0070t\u0068"
	case 8:
		return "\u006d\u0061\u0078\u0044\u0065\u0070\u0074\u0068"
	}
	return ""
}

func (_gcdfd ST_ParameterId) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_eegf := _a.Attr{}
	_eegf.Name = name
	switch _gcdfd {
	case ST_ParameterIdUnset:
		_eegf.Value = ""
	case ST_ParameterIdHorzAlign:
		_eegf.Value = "\u0068o\u0072\u007a\u0041\u006c\u0069\u0067n"
	case ST_ParameterIdVertAlign:
		_eegf.Value = "\u0076e\u0072\u0074\u0041\u006c\u0069\u0067n"
	case ST_ParameterIdChDir:
		_eegf.Value = "\u0063\u0068\u0044i\u0072"
	case ST_ParameterIdChAlign:
		_eegf.Value = "\u0063h\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdSecChAlign:
		_eegf.Value = "\u0073\u0065\u0063\u0043\u0068\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdLinDir:
		_eegf.Value = "\u006c\u0069\u006e\u0044\u0069\u0072"
	case ST_ParameterIdSecLinDir:
		_eegf.Value = "\u0073e\u0063\u004c\u0069\u006e\u0044\u0069r"
	case ST_ParameterIdStElem:
		_eegf.Value = "\u0073\u0074\u0045\u006c\u0065\u006d"
	case ST_ParameterIdBendPt:
		_eegf.Value = "\u0062\u0065\u006e\u0064\u0050\u0074"
	case ST_ParameterIdConnRout:
		_eegf.Value = "\u0063\u006f\u006e\u006e\u0052\u006f\u0075\u0074"
	case ST_ParameterIdBegSty:
		_eegf.Value = "\u0062\u0065\u0067\u0053\u0074\u0079"
	case ST_ParameterIdEndSty:
		_eegf.Value = "\u0065\u006e\u0064\u0053\u0074\u0079"
	case ST_ParameterIdDim:
		_eegf.Value = "\u0064\u0069\u006d"
	case ST_ParameterIdRotPath:
		_eegf.Value = "\u0072o\u0074\u0050\u0061\u0074\u0068"
	case ST_ParameterIdCtrShpMap:
		_eegf.Value = "\u0063t\u0072\u0053\u0068\u0070\u004d\u0061p"
	case ST_ParameterIdNodeHorzAlign:
		_eegf.Value = "\u006e\u006f\u0064\u0065\u0048\u006f\u0072\u007a\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdNodeVertAlign:
		_eegf.Value = "\u006e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdFallback:
		_eegf.Value = "\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b"
	case ST_ParameterIdTxDir:
		_eegf.Value = "\u0074\u0078\u0044i\u0072"
	case ST_ParameterIdPyraAcctPos:
		_eegf.Value = "p\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0050\u006f\u0073"
	case ST_ParameterIdPyraAcctTxMar:
		_eegf.Value = "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054\u0078\u004d\u0061\u0072"
	case ST_ParameterIdTxBlDir:
		_eegf.Value = "\u0074x\u0042\u006c\u0044\u0069\u0072"
	case ST_ParameterIdTxAnchorHorz:
		_eegf.Value = "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u007a"
	case ST_ParameterIdTxAnchorVert:
		_eegf.Value = "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0056\u0065\u0072\u0074"
	case ST_ParameterIdTxAnchorHorzCh:
		_eegf.Value = "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0048o\u0072\u007a\u0043\u0068"
	case ST_ParameterIdTxAnchorVertCh:
		_eegf.Value = "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0043\u0068"
	case ST_ParameterIdParTxLTRAlign:
		_eegf.Value = "\u0070\u0061\u0072\u0054\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdParTxRTLAlign:
		_eegf.Value = "\u0070\u0061\u0072\u0054\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e"
	case ST_ParameterIdShpTxLTRAlignCh:
		_eegf.Value = "\u0073h\u0070T\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e\u0043\u0068"
	case ST_ParameterIdShpTxRTLAlignCh:
		_eegf.Value = "\u0073h\u0070T\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e\u0043\u0068"
	case ST_ParameterIdAutoTxRot:
		_eegf.Value = "\u0061u\u0074\u006f\u0054\u0078\u0052\u006ft"
	case ST_ParameterIdGrDir:
		_eegf.Value = "\u0067\u0072\u0044i\u0072"
	case ST_ParameterIdFlowDir:
		_eegf.Value = "\u0066l\u006f\u0077\u0044\u0069\u0072"
	case ST_ParameterIdContDir:
		_eegf.Value = "\u0063o\u006e\u0074\u0044\u0069\u0072"
	case ST_ParameterIdBkpt:
		_eegf.Value = "\u0062\u006b\u0070\u0074"
	case ST_ParameterIdOff:
		_eegf.Value = "\u006f\u0066\u0066"
	case ST_ParameterIdHierAlign:
		_eegf.Value = "\u0068i\u0065\u0072\u0041\u006c\u0069\u0067n"
	case ST_ParameterIdBkPtFixedVal:
		_eegf.Value = "\u0062\u006b\u0050t\u0046\u0069\u0078\u0065\u0064\u0056\u0061\u006c"
	case ST_ParameterIdStBulletLvl:
		_eegf.Value = "s\u0074\u0042\u0075\u006c\u006c\u0065\u0074\u004c\u0076\u006c"
	case ST_ParameterIdStAng:
		_eegf.Value = "\u0073\u0074\u0041n\u0067"
	case ST_ParameterIdSpanAng:
		_eegf.Value = "\u0073p\u0061\u006e\u0041\u006e\u0067"
	case ST_ParameterIdAr:
		_eegf.Value = "\u0061\u0072"
	case ST_ParameterIdLnSpPar:
		_eegf.Value = "\u006cn\u0053\u0070\u0050\u0061\u0072"
	case ST_ParameterIdLnSpAfParP:
		_eegf.Value = "\u006c\u006e\u0053\u0070\u0041\u0066\u0050\u0061\u0072\u0050"
	case ST_ParameterIdLnSpCh:
		_eegf.Value = "\u006c\u006e\u0053\u0070\u0043\u0068"
	case ST_ParameterIdLnSpAfChP:
		_eegf.Value = "\u006cn\u0053\u0070\u0041\u0066\u0043\u0068P"
	case ST_ParameterIdRtShortDist:
		_eegf.Value = "r\u0074\u0053\u0068\u006f\u0072\u0074\u0044\u0069\u0073\u0074"
	case ST_ParameterIdAlignTx:
		_eegf.Value = "\u0061l\u0069\u0067\u006e\u0054\u0078"
	case ST_ParameterIdPyraLvlNode:
		_eegf.Value = "p\u0079\u0072\u0061\u004c\u0076\u006c\u004e\u006f\u0064\u0065"
	case ST_ParameterIdPyraAcctBkgdNode:
		_eegf.Value = "\u0070\u0079r\u0061\u0041\u0063c\u0074\u0042\u006b\u0067\u0064\u004e\u006f\u0064\u0065"
	case ST_ParameterIdPyraAcctTxNode:
		_eegf.Value = "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054x\u004e\u006f\u0064\u0065"
	case ST_ParameterIdSrcNode:
		_eegf.Value = "\u0073r\u0063\u004e\u006f\u0064\u0065"
	case ST_ParameterIdDstNode:
		_eegf.Value = "\u0064s\u0074\u004e\u006f\u0064\u0065"
	case ST_ParameterIdBegPts:
		_eegf.Value = "\u0062\u0065\u0067\u0050\u0074\u0073"
	case ST_ParameterIdEndPts:
		_eegf.Value = "\u0065\u006e\u0064\u0050\u0074\u0073"
	}
	return _eegf, nil
}

// Validate validates the CT_CTCategories and its children
func (_gcg *CT_CTCategories) Validate() error {
	return _gcg.ValidateWithPath("\u0043T\u005fC\u0054\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073")
}

// Validate validates the CT_SampleData and its children
func (_aacfb *CT_SampleData) Validate() error {
	return _aacfb.ValidateWithPath("\u0043\u0054\u005f\u0053\u0061\u006d\u0070\u006c\u0065\u0044\u0061\u0074\u0061")
}

func (_cfaa *CT_DiagramDefinitionHeader) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _ebfba := range start.Attr {
		if _ebfba.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_adaa, _afcec := _ebfba.Value, error(nil)
			if _afcec != nil {
				return _afcec
			}
			_cfaa.UniqueIdAttr = _adaa
			continue
		}
		if _ebfba.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_cfg, _facgg := _ebfba.Value, error(nil)
			if _facgg != nil {
				return _facgg
			}
			_cfaa.MinVerAttr = &_cfg
			continue
		}
		if _ebfba.Name.Local == "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065" {
			_eecc, _daba := _ebfba.Value, error(nil)
			if _daba != nil {
				return _daba
			}
			_cfaa.DefStyleAttr = &_eecc
			continue
		}
		if _ebfba.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_gcgac, _fbae := _ae.ParseInt(_ebfba.Value, 10, 32)
			if _fbae != nil {
				return _fbae
			}
			_ebfbg := int32(_gcgac)
			_cfaa.ResIdAttr = &_ebfbg
			continue
		}
	}
_dacd:
	for {
		_ggcb, _eeeb := d.Token()
		if _eeeb != nil {
			return _eeeb
		}
		switch _bbgfc := _ggcb.(type) {
		case _a.StartElement:
			switch _bbgfc.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_acdf := NewCT_Name()
				if _gbade := d.DecodeElement(_acdf, &_bbgfc); _gbade != nil {
					return _gbade
				}
				_cfaa.Title = append(_cfaa.Title, _acdf)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_fga := NewCT_Description()
				if _dafe := d.DecodeElement(_fga, &_bbgfc); _dafe != nil {
					return _dafe
				}
				_cfaa.Desc = append(_cfaa.Desc, _fga)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_cfaa.CatLst = NewCT_Categories()
				if _cba := d.DecodeElement(_cfaa.CatLst, &_bbgfc); _cba != nil {
					return _cba
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_cfaa.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _fgbb := d.DecodeElement(_cfaa.ExtLst, &_bbgfc); _fgbb != nil {
					return _fgbb
				}
			default:
				_b.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020o\u006e\u0020\u0043\u0054_\u0044\u0069a\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065\u0072\u0020\u0025\u0076", _bbgfc.Name)
				if _ceg := d.Skip(); _ceg != nil {
					return _ceg
				}
			}
		case _a.EndElement:
			break _dacd
		case _a.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Direction and its children
func (_agfa *CT_Direction) Validate() error {
	return _agfa.ValidateWithPath("\u0043\u0054\u005fD\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
}

// Validate validates the AG_ConstraintRefAttributes and its children
func (_gd *AG_ConstraintRefAttributes) Validate() error {
	return _gd.ValidateWithPath("\u0041\u0047\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069n\u0074\u0052\u0065\u0066\u0041\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0073")
}

// ValidateWithPath validates the CT_Cxn and its children, prefixing error messages with path
func (_cge *CT_Cxn) ValidateWithPath(path string) error {
	if _gbadc := _cge.ModelIdAttr.ValidateWithPath(path + "\u002f\u004d\u006fd\u0065\u006c\u0049\u0064\u0041\u0074\u0074\u0072"); _gbadc != nil {
		return _gbadc
	}
	if _ddcg := _cge.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _ddcg != nil {
		return _ddcg
	}
	if _effg := _cge.SrcIdAttr.ValidateWithPath(path + "\u002f\u0053\u0072\u0063\u0049\u0064\u0041\u0074\u0074\u0072"); _effg != nil {
		return _effg
	}
	if _ddff := _cge.DestIdAttr.ValidateWithPath(path + "/\u0044\u0065\u0073\u0074\u0049\u0064\u0041\u0074\u0074\u0072"); _ddff != nil {
		return _ddff
	}
	if _cge.ParTransIdAttr != nil {
		if _abef := _cge.ParTransIdAttr.ValidateWithPath(path + "\u002fP\u0061r\u0054\u0072\u0061\u006e\u0073\u0049\u0064\u0041\u0074\u0074\u0072"); _abef != nil {
			return _abef
		}
	}
	if _cge.SibTransIdAttr != nil {
		if _fcce := _cge.SibTransIdAttr.ValidateWithPath(path + "\u002fS\u0069b\u0054\u0072\u0061\u006e\u0073\u0049\u0064\u0041\u0074\u0074\u0072"); _fcce != nil {
			return _fcce
		}
	}
	if _cge.ExtLst != nil {
		if _bgdf := _cge.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _bgdf != nil {
			return _bgdf
		}
	}
	return nil
}

func (_cfcff *CT_Otherwise) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _cfcff.NameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_cfcff.NameAttr)})
	}
	e.EncodeToken(start)
	if _cfcff.Alg != nil {
		_fbad := _a.StartElement{Name: _a.Name{Local: "\u0061\u006c\u0067"}}
		for _, _gaee := range _cfcff.Alg {
			e.EncodeElement(_gaee, _fbad)
		}
	}
	if _cfcff.Shape != nil {
		_gdbe := _a.StartElement{Name: _a.Name{Local: "\u0073\u0068\u0061p\u0065"}}
		for _, _bgbb := range _cfcff.Shape {
			e.EncodeElement(_bgbb, _gdbe)
		}
	}
	if _cfcff.PresOf != nil {
		_abfc := _a.StartElement{Name: _a.Name{Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}}
		for _, _bada := range _cfcff.PresOf {
			e.EncodeElement(_bada, _abfc)
		}
	}
	if _cfcff.ConstrLst != nil {
		_cbbef := _a.StartElement{Name: _a.Name{Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}}
		for _, _gdcdg := range _cfcff.ConstrLst {
			e.EncodeElement(_gdcdg, _cbbef)
		}
	}
	if _cfcff.RuleLst != nil {
		_gegc := _a.StartElement{Name: _a.Name{Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}}
		for _, _eadc := range _cfcff.RuleLst {
			e.EncodeElement(_eadc, _gegc)
		}
	}
	if _cfcff.ForEach != nil {
		_adfd := _a.StartElement{Name: _a.Name{Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}}
		for _, _ecgaa := range _cfcff.ForEach {
			e.EncodeElement(_ecgaa, _adfd)
		}
	}
	if _cfcff.LayoutNode != nil {
		_badf := _a.StartElement{Name: _a.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
		for _, _gdbf := range _cfcff.LayoutNode {
			e.EncodeElement(_gdbf, _badf)
		}
	}
	if _cfcff.Choose != nil {
		_abde := _a.StartElement{Name: _a.Name{Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}}
		for _, _ffdb := range _cfcff.Choose {
			e.EncodeElement(_ffdb, _abde)
		}
	}
	if _cfcff.ExtLst != nil {
		_dace := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		for _, _gbfc := range _cfcff.ExtLst {
			e.EncodeElement(_gbfc, _dace)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

type CT_HierBranchStyle struct{ ValAttr ST_HierBranchStyle }

func (_acaga *ST_AlgorithmType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_acaga = 0
	case "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e":
		*_acaga = 1
	case "\u0063\u006f\u006e\u006e":
		*_acaga = 2
	case "\u0063\u0079\u0063l\u0065":
		*_acaga = 3
	case "\u0068i\u0065\u0072\u0043\u0068\u0069\u006cd":
		*_acaga = 4
	case "\u0068\u0069\u0065\u0072\u0052\u006f\u006f\u0074":
		*_acaga = 5
	case "\u0070\u0079\u0072\u0061":
		*_acaga = 6
	case "\u006c\u0069\u006e":
		*_acaga = 7
	case "\u0073\u0070":
		*_acaga = 8
	case "\u0074\u0078":
		*_acaga = 9
	case "\u0073\u006e\u0061k\u0065":
		*_acaga = 10
	}
	return nil
}

func (_ebgfe *ST_ConstraintType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_efcbc, _eccfa := d.Token()
	if _eccfa != nil {
		return _eccfa
	}
	if _fcga, _fffa := _efcbc.(_a.EndElement); _fffa && _fcga.Name == start.Name {
		*_ebgfe = 1
		return nil
	}
	if _cdedg, _gdagd := _efcbc.(_a.CharData); !_gdagd {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _efcbc)
	} else {
		switch string(_cdedg) {
		case "":
			*_ebgfe = 0
		case "\u006e\u006f\u006e\u0065":
			*_ebgfe = 1
		case "\u0061\u006c\u0069\u0067\u006e\u004f\u0066\u0066":
			*_ebgfe = 2
		case "\u0062e\u0067\u004d\u0061\u0072\u0067":
			*_ebgfe = 3
		case "\u0062\u0065\u006e\u0064\u0044\u0069\u0073\u0074":
			*_ebgfe = 4
		case "\u0062\u0065\u0067\u0050\u0061\u0064":
			*_ebgfe = 5
		case "\u0062":
			*_ebgfe = 6
		case "\u0062\u004d\u0061r\u0067":
			*_ebgfe = 7
		case "\u0062\u004f\u0066\u0066":
			*_ebgfe = 8
		case "\u0063\u0074\u0072\u0058":
			*_ebgfe = 9
		case "\u0063t\u0072\u0058\u004f\u0066\u0066":
			*_ebgfe = 10
		case "\u0063\u0074\u0072\u0059":
			*_ebgfe = 11
		case "\u0063t\u0072\u0059\u004f\u0066\u0066":
			*_ebgfe = 12
		case "\u0063\u006f\u006e\u006e\u0044\u0069\u0073\u0074":
			*_ebgfe = 13
		case "\u0064\u0069\u0061\u006d":
			*_ebgfe = 14
		case "\u0065n\u0064\u004d\u0061\u0072\u0067":
			*_ebgfe = 15
		case "\u0065\u006e\u0064\u0050\u0061\u0064":
			*_ebgfe = 16
		case "\u0068":
			*_ebgfe = 17
		case "\u0068\u0041\u0072\u0048":
			*_ebgfe = 18
		case "\u0068\u004f\u0066\u0066":
			*_ebgfe = 19
		case "\u006c":
			*_ebgfe = 20
		case "\u006c\u004d\u0061r\u0067":
			*_ebgfe = 21
		case "\u006c\u004f\u0066\u0066":
			*_ebgfe = 22
		case "\u0072":
			*_ebgfe = 23
		case "\u0072\u004d\u0061r\u0067":
			*_ebgfe = 24
		case "\u0072\u004f\u0066\u0066":
			*_ebgfe = 25
		case "\u0070\u0072\u0069\u006d\u0046\u006f\u006e\u0074\u0053\u007a":
			*_ebgfe = 26
		case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0052\u0061\u0074\u0069\u006f":
			*_ebgfe = 27
		case "\u0073e\u0063\u0046\u006f\u006e\u0074\u0053z":
			*_ebgfe = 28
		case "\u0073\u0069\u0062S\u0070":
			*_ebgfe = 29
		case "\u0073\u0065\u0063\u0053\u0069\u0062\u0053\u0070":
			*_ebgfe = 30
		case "\u0073\u0070":
			*_ebgfe = 31
		case "\u0073t\u0065\u006d\u0054\u0068\u0069\u0063k":
			*_ebgfe = 32
		case "\u0074":
			*_ebgfe = 33
		case "\u0074\u004d\u0061r\u0067":
			*_ebgfe = 34
		case "\u0074\u004f\u0066\u0066":
			*_ebgfe = 35
		case "\u0075\u0073\u0065r\u0041":
			*_ebgfe = 36
		case "\u0075\u0073\u0065r\u0042":
			*_ebgfe = 37
		case "\u0075\u0073\u0065r\u0043":
			*_ebgfe = 38
		case "\u0075\u0073\u0065r\u0044":
			*_ebgfe = 39
		case "\u0075\u0073\u0065r\u0045":
			*_ebgfe = 40
		case "\u0075\u0073\u0065r\u0046":
			*_ebgfe = 41
		case "\u0075\u0073\u0065r\u0047":
			*_ebgfe = 42
		case "\u0075\u0073\u0065r\u0048":
			*_ebgfe = 43
		case "\u0075\u0073\u0065r\u0049":
			*_ebgfe = 44
		case "\u0075\u0073\u0065r\u004a":
			*_ebgfe = 45
		case "\u0075\u0073\u0065r\u004b":
			*_ebgfe = 46
		case "\u0075\u0073\u0065r\u004c":
			*_ebgfe = 47
		case "\u0075\u0073\u0065r\u004d":
			*_ebgfe = 48
		case "\u0075\u0073\u0065r\u004e":
			*_ebgfe = 49
		case "\u0075\u0073\u0065r\u004f":
			*_ebgfe = 50
		case "\u0075\u0073\u0065r\u0050":
			*_ebgfe = 51
		case "\u0075\u0073\u0065r\u0051":
			*_ebgfe = 52
		case "\u0075\u0073\u0065r\u0052":
			*_ebgfe = 53
		case "\u0075\u0073\u0065r\u0053":
			*_ebgfe = 54
		case "\u0075\u0073\u0065r\u0054":
			*_ebgfe = 55
		case "\u0075\u0073\u0065r\u0055":
			*_ebgfe = 56
		case "\u0075\u0073\u0065r\u0056":
			*_ebgfe = 57
		case "\u0075\u0073\u0065r\u0057":
			*_ebgfe = 58
		case "\u0075\u0073\u0065r\u0058":
			*_ebgfe = 59
		case "\u0075\u0073\u0065r\u0059":
			*_ebgfe = 60
		case "\u0075\u0073\u0065r\u005a":
			*_ebgfe = 61
		case "\u0077":
			*_ebgfe = 62
		case "\u0077\u0041\u0072\u0048":
			*_ebgfe = 63
		case "\u0077\u004f\u0066\u0066":
			*_ebgfe = 64
		}
	}
	_efcbc, _eccfa = d.Token()
	if _eccfa != nil {
		return _eccfa
	}
	if _gcag, _fcdd := _efcbc.(_a.EndElement); _fcdd && _gcag.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _efcbc)
}

func (_baefd *ST_LayoutShapeType) Validate() error { return _baefd.ValidateWithPath("") }

// ValidateWithPath validates the CT_RelIds and its children, prefixing error messages with path
func (_edeb *CT_RelIds) ValidateWithPath(path string) error { return nil }

type ST_FallbackDimension byte

func (_adee ST_OutputShapeType) Validate() error { return _adee.ValidateWithPath("") }

// Validate validates the CT_PresentationOf and its children
func (_egad *CT_PresentationOf) Validate() error {
	return _egad.ValidateWithPath("\u0043\u0054\u005f\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u004f\u0066")
}

func (_gdag *ST_AlgorithmType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_fgcba, _eedea := d.Token()
	if _eedea != nil {
		return _eedea
	}
	if _aaedc, _bedcc := _fgcba.(_a.EndElement); _bedcc && _aaedc.Name == start.Name {
		*_gdag = 1
		return nil
	}
	if _feec, _gadfec := _fgcba.(_a.CharData); !_gadfec {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fgcba)
	} else {
		switch string(_feec) {
		case "":
			*_gdag = 0
		case "\u0063o\u006d\u0070\u006f\u0073\u0069\u0074e":
			*_gdag = 1
		case "\u0063\u006f\u006e\u006e":
			*_gdag = 2
		case "\u0063\u0079\u0063l\u0065":
			*_gdag = 3
		case "\u0068i\u0065\u0072\u0043\u0068\u0069\u006cd":
			*_gdag = 4
		case "\u0068\u0069\u0065\u0072\u0052\u006f\u006f\u0074":
			*_gdag = 5
		case "\u0070\u0079\u0072\u0061":
			*_gdag = 6
		case "\u006c\u0069\u006e":
			*_gdag = 7
		case "\u0073\u0070":
			*_gdag = 8
		case "\u0074\u0078":
			*_gdag = 9
		case "\u0073\u006e\u0061k\u0065":
			*_gdag = 10
		}
	}
	_fgcba, _eedea = d.Token()
	if _eedea != nil {
		return _eedea
	}
	if _edaeg, _adgbb := _fgcba.(_a.EndElement); _adgbb && _edaeg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fgcba)
}

func (_gdafg ST_ConstraintRelationship) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_gdafg.String(), start)
}

func (_ffdga ST_AnimLvlStr) Validate() error { return _ffdga.ValidateWithPath("") }

func (_eegea *RelIds) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_eegea.CT_RelIds = *NewCT_RelIds()
	for _, _becf := range start.Attr {
		if _becf.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _becf.Name.Local == "\u0064\u006d" || _becf.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _becf.Name.Local == "\u0064\u006d" {
			_bbfae, _badd := _becf.Value, error(nil)
			if _badd != nil {
				return _badd
			}
			_eegea.DmAttr = _bbfae
			continue
		}
		if _becf.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _becf.Name.Local == "\u006c\u006f" || _becf.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _becf.Name.Local == "\u006c\u006f" {
			_baefc, _ddaa := _becf.Value, error(nil)
			if _ddaa != nil {
				return _ddaa
			}
			_eegea.LoAttr = _baefc
			continue
		}
		if _becf.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _becf.Name.Local == "\u0071\u0073" || _becf.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _becf.Name.Local == "\u0071\u0073" {
			_febcb, _gbea := _becf.Value, error(nil)
			if _gbea != nil {
				return _gbea
			}
			_eegea.QsAttr = _febcb
			continue
		}
		if _becf.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _becf.Name.Local == "\u0063\u0073" || _becf.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _becf.Name.Local == "\u0063\u0073" {
			_ddef, _aegcd := _becf.Value, error(nil)
			if _aegcd != nil {
				return _aegcd
			}
			_eegea.CsAttr = _ddef
			continue
		}
	}
	for {
		_fbgf, _efbbe := d.Token()
		if _efbbe != nil {
			return _f.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0052e\u006c\u0049d\u0073\u003a\u0020\u0025\u0073", _efbbe)
		}
		if _caad, _bffa := _fbgf.(_a.EndElement); _bffa && _caad.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ccaa *ST_ConnectorDimension) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gfgad, _aegfd := d.Token()
	if _aegfd != nil {
		return _aegfd
	}
	if _aagbc, _dgbfc := _gfgad.(_a.EndElement); _dgbfc && _aagbc.Name == start.Name {
		*_ccaa = 1
		return nil
	}
	if _cbfbc, _ffae := _gfgad.(_a.CharData); !_ffae {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gfgad)
	} else {
		switch string(_cbfbc) {
		case "":
			*_ccaa = 0
		case "\u0031\u0044":
			*_ccaa = 1
		case "\u0032\u0044":
			*_ccaa = 2
		case "\u0063\u0075\u0073\u0074":
			*_ccaa = 3
		}
	}
	_gfgad, _aegfd = d.Token()
	if _aegfd != nil {
		return _aegfd
	}
	if _cedb, _fefagg := _gfgad.(_a.EndElement); _fefagg && _cedb.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gfgad)
}

func (_gdgce ST_PtType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_edgc := _a.Attr{}
	_edgc.Name = name
	switch _gdgce {
	case ST_PtTypeUnset:
		_edgc.Value = ""
	case ST_PtTypeNode:
		_edgc.Value = "\u006e\u006f\u0064\u0065"
	case ST_PtTypeAsst:
		_edgc.Value = "\u0061\u0073\u0073\u0074"
	case ST_PtTypeDoc:
		_edgc.Value = "\u0064\u006f\u0063"
	case ST_PtTypePres:
		_edgc.Value = "\u0070\u0072\u0065\u0073"
	case ST_PtTypeParTrans:
		_edgc.Value = "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073"
	case ST_PtTypeSibTrans:
		_edgc.Value = "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073"
	}
	return _edgc, nil
}

func (_gbddc *ST_AnimOneStr) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_gbddc = 0
	case "\u006e\u006f\u006e\u0065":
		*_gbddc = 1
	case "\u006f\u006e\u0065":
		*_gbddc = 2
	case "\u0062\u0072\u0061\u006e\u0063\u0068":
		*_gbddc = 3
	}
	return nil
}

func (_adea *CT_CxnList) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_aeab:
	for {
		_agdfd, _fccc := d.Token()
		if _fccc != nil {
			return _fccc
		}
		switch _dddg := _agdfd.(type) {
		case _a.StartElement:
			switch _dddg.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0078\u006e"}:
				_adgfe := NewCT_Cxn()
				if _eafaa := d.DecodeElement(_adgfe, &_dddg); _eafaa != nil {
					return _eafaa
				}
				_adea.Cxn = append(_adea.Cxn, _adgfe)
			default:
				_b.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fC\u0078\u006e\u004c\u0069\u0073\u0074\u0020\u0025\u0076", _dddg.Name)
				if _gggfe := d.Skip(); _gggfe != nil {
					return _gggfe
				}
			}
		case _a.EndElement:
			break _aeab
		case _a.CharData:
		}
	}
	return nil
}

type ST_AxisTypes []ST_AxisType

// Validate validates the CT_ColorTransformHeader and its children
func (_facgf *CT_ColorTransformHeader) Validate() error {
	return _facgf.ValidateWithPath("\u0043\u0054\u005fCo\u006c\u006f\u0072\u0054\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065\u0072")
}

func (_fceg *ST_ChildDirection) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_bfefe, _bbcb := d.Token()
	if _bbcb != nil {
		return _bbcb
	}
	if _dceg, _dccb := _bfefe.(_a.EndElement); _dccb && _dceg.Name == start.Name {
		*_fceg = 1
		return nil
	}
	if _fegca, _adcab := _bfefe.(_a.CharData); !_adcab {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bfefe)
	} else {
		switch string(_fegca) {
		case "":
			*_fceg = 0
		case "\u0068\u006f\u0072\u007a":
			*_fceg = 1
		case "\u0076\u0065\u0072\u0074":
			*_fceg = 2
		}
	}
	_bfefe, _bbcb = d.Token()
	if _bbcb != nil {
		return _bbcb
	}
	if _gcegg, _dcae := _bfefe.(_a.EndElement); _dcae && _gcegg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bfefe)
}

type ST_UnsignedInts []uint32

func (_edfee ST_BoolOperator) ValidateWithPath(path string) error {
	switch _edfee {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_edfee))
	}
	return nil
}

func (_ggdag ST_CxnType) ValidateWithPath(path string) error {
	switch _ggdag {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ggdag))
	}
	return nil
}

// ValidateWithPath validates the StyleDefHdrLst and its children, prefixing error messages with path
func (_cfdab *StyleDefHdrLst) ValidateWithPath(path string) error {
	if _ffafd := _cfdab.CT_StyleDefinitionHeaderLst.ValidateWithPath(path); _ffafd != nil {
		return _ffafd
	}
	return nil
}

type ST_ResizeHandlesStr byte

func (_adddg ST_AnimOneStr) Validate() error { return _adddg.ValidateWithPath("") }

type CT_Shape struct {
	RotAttr       *float64
	TypeAttr      *ST_LayoutShapeType
	BlipAttr      *string
	ZOrderOffAttr *int32
	HideGeomAttr  *bool
	LkTxEntryAttr *bool
	BlipPhldrAttr *bool
	AdjLst        *CT_AdjLst
	ExtLst        *_bg.CT_OfficeArtExtensionList
}

// ValidateWithPath validates the CT_Parameter and its children, prefixing error messages with path
func (_gebc *CT_Parameter) ValidateWithPath(path string) error {
	if _gebc.TypeAttr == ST_ParameterIdUnset {
		return _f.Errorf("\u0025\u0073\u002f\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u0069\u0073\u0020a\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _fefa := _gebc.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _fefa != nil {
		return _fefa
	}
	if _beac := _gebc.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _beac != nil {
		return _beac
	}
	return nil
}

func (_bfgag *ST_TextDirection) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_edgec, _cfdf := d.Token()
	if _cfdf != nil {
		return _cfdf
	}
	if _gdbgg, _gcgcc := _edgec.(_a.EndElement); _gcgcc && _gdbgg.Name == start.Name {
		*_bfgag = 1
		return nil
	}
	if _gdee, _fgefg := _edgec.(_a.CharData); !_fgefg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _edgec)
	} else {
		switch string(_gdee) {
		case "":
			*_bfgag = 0
		case "\u0066\u0072\u006fm\u0054":
			*_bfgag = 1
		case "\u0066\u0072\u006fm\u0042":
			*_bfgag = 2
		}
	}
	_edgec, _cfdf = d.Token()
	if _cfdf != nil {
		return _cfdf
	}
	if _eaaf, _cbgdb := _edgec.(_a.EndElement); _cbgdb && _eaaf.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _edgec)
}

const (
	ST_ConstraintTypeUnset         ST_ConstraintType = 0
	ST_ConstraintTypeNone          ST_ConstraintType = 1
	ST_ConstraintTypeAlignOff      ST_ConstraintType = 2
	ST_ConstraintTypeBegMarg       ST_ConstraintType = 3
	ST_ConstraintTypeBendDist      ST_ConstraintType = 4
	ST_ConstraintTypeBegPad        ST_ConstraintType = 5
	ST_ConstraintTypeB             ST_ConstraintType = 6
	ST_ConstraintTypeBMarg         ST_ConstraintType = 7
	ST_ConstraintTypeBOff          ST_ConstraintType = 8
	ST_ConstraintTypeCtrX          ST_ConstraintType = 9
	ST_ConstraintTypeCtrXOff       ST_ConstraintType = 10
	ST_ConstraintTypeCtrY          ST_ConstraintType = 11
	ST_ConstraintTypeCtrYOff       ST_ConstraintType = 12
	ST_ConstraintTypeConnDist      ST_ConstraintType = 13
	ST_ConstraintTypeDiam          ST_ConstraintType = 14
	ST_ConstraintTypeEndMarg       ST_ConstraintType = 15
	ST_ConstraintTypeEndPad        ST_ConstraintType = 16
	ST_ConstraintTypeH             ST_ConstraintType = 17
	ST_ConstraintTypeHArH          ST_ConstraintType = 18
	ST_ConstraintTypeHOff          ST_ConstraintType = 19
	ST_ConstraintTypeL             ST_ConstraintType = 20
	ST_ConstraintTypeLMarg         ST_ConstraintType = 21
	ST_ConstraintTypeLOff          ST_ConstraintType = 22
	ST_ConstraintTypeR             ST_ConstraintType = 23
	ST_ConstraintTypeRMarg         ST_ConstraintType = 24
	ST_ConstraintTypeROff          ST_ConstraintType = 25
	ST_ConstraintTypePrimFontSz    ST_ConstraintType = 26
	ST_ConstraintTypePyraAcctRatio ST_ConstraintType = 27
	ST_ConstraintTypeSecFontSz     ST_ConstraintType = 28
	ST_ConstraintTypeSibSp         ST_ConstraintType = 29
	ST_ConstraintTypeSecSibSp      ST_ConstraintType = 30
	ST_ConstraintTypeSp            ST_ConstraintType = 31
	ST_ConstraintTypeStemThick     ST_ConstraintType = 32
	ST_ConstraintTypeT             ST_ConstraintType = 33
	ST_ConstraintTypeTMarg         ST_ConstraintType = 34
	ST_ConstraintTypeTOff          ST_ConstraintType = 35
	ST_ConstraintTypeUserA         ST_ConstraintType = 36
	ST_ConstraintTypeUserB         ST_ConstraintType = 37
	ST_ConstraintTypeUserC         ST_ConstraintType = 38
	ST_ConstraintTypeUserD         ST_ConstraintType = 39
	ST_ConstraintTypeUserE         ST_ConstraintType = 40
	ST_ConstraintTypeUserF         ST_ConstraintType = 41
	ST_ConstraintTypeUserG         ST_ConstraintType = 42
	ST_ConstraintTypeUserH         ST_ConstraintType = 43
	ST_ConstraintTypeUserI         ST_ConstraintType = 44
	ST_ConstraintTypeUserJ         ST_ConstraintType = 45
	ST_ConstraintTypeUserK         ST_ConstraintType = 46
	ST_ConstraintTypeUserL         ST_ConstraintType = 47
	ST_ConstraintTypeUserM         ST_ConstraintType = 48
	ST_ConstraintTypeUserN         ST_ConstraintType = 49
	ST_ConstraintTypeUserO         ST_ConstraintType = 50
	ST_ConstraintTypeUserP         ST_ConstraintType = 51
	ST_ConstraintTypeUserQ         ST_ConstraintType = 52
	ST_ConstraintTypeUserR         ST_ConstraintType = 53
	ST_ConstraintTypeUserS         ST_ConstraintType = 54
	ST_ConstraintTypeUserT         ST_ConstraintType = 55
	ST_ConstraintTypeUserU         ST_ConstraintType = 56
	ST_ConstraintTypeUserV         ST_ConstraintType = 57
	ST_ConstraintTypeUserW         ST_ConstraintType = 58
	ST_ConstraintTypeUserX         ST_ConstraintType = 59
	ST_ConstraintTypeUserY         ST_ConstraintType = 60
	ST_ConstraintTypeUserZ         ST_ConstraintType = 61
	ST_ConstraintTypeW             ST_ConstraintType = 62
	ST_ConstraintTypeWArH          ST_ConstraintType = 63
	ST_ConstraintTypeWOff          ST_ConstraintType = 64
)

func (_gega ST_ResizeHandlesStr) Validate() error { return _gega.ValidateWithPath("") }

func (_geeg *CT_ElemPropSet) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _bcffb := range start.Attr {
		if _bcffb.Name.Local == "\u0063\u0075\u0073t\u0046\u006c\u0069\u0070\u0056\u0065\u0072\u0074" {
			_aacgb, _gbc := _ae.ParseBool(_bcffb.Value)
			if _gbc != nil {
				return _gbc
			}
			_geeg.CustFlipVertAttr = &_aacgb
			continue
		}
		if _bcffb.Name.Local == "p\u0072\u0065\u0073\u0041\u0073\u0073\u006f\u0063\u0049\u0044" {
			_bffg, _ecca := ParseUnionST_ModelId(_bcffb.Value)
			if _ecca != nil {
				return _ecca
			}
			_geeg.PresAssocIDAttr = &_bffg
			continue
		}
		if _bcffb.Name.Local == "c\u0075\u0073\u0074\u0046\u006c\u0069\u0070\u0048\u006f\u0072" {
			_afg, _bedg := _ae.ParseBool(_bcffb.Value)
			if _bedg != nil {
				return _bedg
			}
			_geeg.CustFlipHorAttr = &_afg
			continue
		}
		if _bcffb.Name.Local == "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u004c\u0062\u006c" {
			_defc, _befe := _bcffb.Value, error(nil)
			if _befe != nil {
				return _befe
			}
			_geeg.PresStyleLblAttr = &_defc
			continue
		}
		if _bcffb.Name.Local == "\u0063u\u0073\u0074\u0053\u007a\u0058" {
			_dbdf, _cede := _ae.ParseInt(_bcffb.Value, 10, 32)
			if _cede != nil {
				return _cede
			}
			_aeabf := int32(_dbdf)
			_geeg.CustSzXAttr = &_aeabf
			continue
		}
		if _bcffb.Name.Local == "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u0043\u006e\u0074" {
			_gfca, _cdfbe := _ae.ParseInt(_bcffb.Value, 10, 32)
			if _cdfbe != nil {
				return _cdfbe
			}
			_adcf := int32(_gfca)
			_geeg.PresStyleCntAttr = &_adcf
			continue
		}
		if _bcffb.Name.Local == "\u0063u\u0073\u0074\u0053\u007a\u0059" {
			_egbc, _gabg := _ae.ParseInt(_bcffb.Value, 10, 32)
			if _gabg != nil {
				return _gabg
			}
			_cada := int32(_egbc)
			_geeg.CustSzYAttr = &_cada
			continue
		}
		if _bcffb.Name.Local == "\u006co\u0043\u0061\u0074\u0049\u0064" {
			_fbca, _cff := _bcffb.Value, error(nil)
			if _cff != nil {
				return _cff
			}
			_geeg.LoCatIdAttr = &_fbca
			continue
		}
		if _bcffb.Name.Local == "\u0063\u0075\u0073\u0074\u0053\u0063\u0061\u006c\u0065\u0058" {
			_bdce, _fgac := ParseUnionST_PrSetCustVal(_bcffb.Value)
			if _fgac != nil {
				return _fgac
			}
			_geeg.CustScaleXAttr = &_bdce
			continue
		}
		if _bcffb.Name.Local == "\u0071s\u0043\u0061\u0074\u0049\u0064" {
			_dbf, _fdee := _bcffb.Value, error(nil)
			if _fdee != nil {
				return _fdee
			}
			_geeg.QsCatIdAttr = &_dbf
			continue
		}
		if _bcffb.Name.Local == "\u0063\u0075\u0073\u0074\u0053\u0063\u0061\u006c\u0065\u0059" {
			_cecd, _gggae := ParseUnionST_PrSetCustVal(_bcffb.Value)
			if _gggae != nil {
				return _gggae
			}
			_geeg.CustScaleYAttr = &_cecd
			continue
		}
		if _bcffb.Name.Local == "\u0063u\u0073\u0074\u0041\u006e\u0067" {
			_aaag, _dag := _ae.ParseInt(_bcffb.Value, 10, 32)
			if _dag != nil {
				return _dag
			}
			_dgcf := int32(_aaag)
			_geeg.CustAngAttr = &_dgcf
			continue
		}
		if _bcffb.Name.Local == "\u0063u\u0073t\u0052\u0061\u0064\u0053\u0063\u0061\u006c\u0065\u0052\u0061\u0064" {
			_ggbfa, _cbbf := ParseUnionST_PrSetCustVal(_bcffb.Value)
			if _cbbf != nil {
				return _cbbf
			}
			_geeg.CustRadScaleRadAttr = &_ggbfa
			continue
		}
		if _bcffb.Name.Local == "\u0063\u0075\u0073t\u004c\u0069\u006e\u0046\u0061\u0063\u0074\u0058" {
			_feddc, _cde := ParseUnionST_PrSetCustVal(_bcffb.Value)
			if _cde != nil {
				return _cde
			}
			_geeg.CustLinFactXAttr = &_feddc
			continue
		}
		if _bcffb.Name.Local == "\u0071\u0073\u0054\u0079\u0070\u0065\u0049\u0064" {
			_egfe, _fcb := _bcffb.Value, error(nil)
			if _fcb != nil {
				return _fcb
			}
			_geeg.QsTypeIdAttr = &_egfe
			continue
		}
		if _bcffb.Name.Local == "\u0063\u006f\u0068\u0065\u0072\u0065\u006e\u0074\u0033\u0044\u004f\u0066\u0066" {
			_fabe, _ggag := _ae.ParseBool(_bcffb.Value)
			if _ggag != nil {
				return _ggag
			}
			_geeg.Coherent3DOffAttr = &_fabe
			continue
		}
		if _bcffb.Name.Local == "\u0063\u0075\u0073t\u0054" {
			_cgdd, _cfcc := _ae.ParseBool(_bcffb.Value)
			if _cfcc != nil {
				return _cfcc
			}
			_geeg.CustTAttr = &_cgdd
			continue
		}
		if _bcffb.Name.Local == "\u0070\u0072\u0065\u0073\u004e\u0061\u006d\u0065" {
			_efee, _bdgd := _bcffb.Value, error(nil)
			if _bdgd != nil {
				return _bdgd
			}
			_geeg.PresNameAttr = &_efee
			continue
		}
		if _bcffb.Name.Local == "c\u0075s\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074N\u0065\u0069\u0067\u0068bo\u0072\u0059" {
			_cacf, _cdcf := ParseUnionST_PrSetCustVal(_bcffb.Value)
			if _cdcf != nil {
				return _cdcf
			}
			_geeg.CustLinFactNeighborYAttr = &_cacf
			continue
		}
		if _bcffb.Name.Local == "\u0063\u0075\u0073t\u004c\u0069\u006e\u0046\u0061\u0063\u0074\u0059" {
			_cbgd, _gff := ParseUnionST_PrSetCustVal(_bcffb.Value)
			if _gff != nil {
				return _gff
			}
			_geeg.CustLinFactYAttr = &_cbgd
			continue
		}
		if _bcffb.Name.Local == "\u006c\u006f\u0054\u0079\u0070\u0065\u0049\u0064" {
			_ecae, _ead := _bcffb.Value, error(nil)
			if _ead != nil {
				return _ead
			}
			_geeg.LoTypeIdAttr = &_ecae
			continue
		}
		if _bcffb.Name.Local == "\u0063u\u0073t\u0052\u0061\u0064\u0053\u0063\u0061\u006c\u0065\u0049\u006e\u0063" {
			_eccf, _dgfb := ParseUnionST_PrSetCustVal(_bcffb.Value)
			if _dgfb != nil {
				return _dgfb
			}
			_geeg.CustRadScaleIncAttr = &_eccf
			continue
		}
		if _bcffb.Name.Local == "\u0070\u0068\u006cd\u0072" {
			_gcbd, _ccab := _ae.ParseBool(_bcffb.Value)
			if _ccab != nil {
				return _ccab
			}
			_geeg.PhldrAttr = &_gcbd
			continue
		}
		if _bcffb.Name.Local == "\u0063\u0073\u0054\u0079\u0070\u0065\u0049\u0064" {
			_fedg, _dedd := _bcffb.Value, error(nil)
			if _dedd != nil {
				return _dedd
			}
			_geeg.CsTypeIdAttr = &_fedg
			continue
		}
		if _bcffb.Name.Local == "\u0063s\u0043\u0061\u0074\u0049\u0064" {
			_bfdc, _cacd := _bcffb.Value, error(nil)
			if _cacd != nil {
				return _cacd
			}
			_geeg.CsCatIdAttr = &_bfdc
			continue
		}
		if _bcffb.Name.Local == "\u0070\u0068\u006c\u0064\u0072\u0054" {
			_eccg, _dcfc := _bcffb.Value, error(nil)
			if _dcfc != nil {
				return _dcfc
			}
			_geeg.PhldrTAttr = &_eccg
			continue
		}
		if _bcffb.Name.Local == "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u0049\u0064\u0078" {
			_agfc, _cdeb := _ae.ParseInt(_bcffb.Value, 10, 32)
			if _cdeb != nil {
				return _cdeb
			}
			_afcc := int32(_agfc)
			_geeg.PresStyleIdxAttr = &_afcc
			continue
		}
		if _bcffb.Name.Local == "c\u0075s\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074N\u0065\u0069\u0067\u0068bo\u0072\u0058" {
			_gafd, _adff := ParseUnionST_PrSetCustVal(_bcffb.Value)
			if _adff != nil {
				return _adff
			}
			_geeg.CustLinFactNeighborXAttr = &_gafd
			continue
		}
	}
_fedae:
	for {
		_fbfgg, _fbabd := d.Token()
		if _fbabd != nil {
			return _fbabd
		}
		switch _gcdc := _fbfgg.(type) {
		case _a.StartElement:
			switch _gcdc.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004c\u0061\u0079\u006f\u0075t\u0056\u0061\u0072\u0073"}:
				_geeg.PresLayoutVars = NewCT_LayoutVariablePropertySet()
				if _fabc := d.DecodeElement(_geeg.PresLayoutVars, &_gcdc); _fabc != nil {
					return _fabc
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079l\u0065"}:
				_geeg.Style = _bg.NewCT_ShapeStyle()
				if _ggaa := d.DecodeElement(_geeg.Style, &_gcdc); _ggaa != nil {
					return _ggaa
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0045\u006c\u0065\u006d\u0050\u0072\u006fp\u0053e\u0074\u0020\u0025\u0076", _gcdc.Name)
				if _dfcd := d.Skip(); _dfcd != nil {
					return _dfcd
				}
			}
		case _a.EndElement:
			break _fedae
		case _a.CharData:
		}
	}
	return nil
}

func (_gebcd *ST_VariableType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_adeba, _dcggf := d.Token()
	if _dcggf != nil {
		return _dcggf
	}
	if _ggeec, _cafeg := _adeba.(_a.EndElement); _cafeg && _ggeec.Name == start.Name {
		*_gebcd = 1
		return nil
	}
	if _fcaa, _agaef := _adeba.(_a.CharData); !_agaef {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _adeba)
	} else {
		switch string(_fcaa) {
		case "":
			*_gebcd = 0
		case "\u006e\u006f\u006e\u0065":
			*_gebcd = 1
		case "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074":
			*_gebcd = 2
		case "\u0063\u0068\u004da\u0078":
			*_gebcd = 3
		case "\u0063\u0068\u0050\u0072\u0065\u0066":
			*_gebcd = 4
		case "\u0062\u0075\u006c\u0045\u006e\u0061\u0062\u006c\u0065\u0064":
			*_gebcd = 5
		case "\u0064\u0069\u0072":
			*_gebcd = 6
		case "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068":
			*_gebcd = 7
		case "\u0061n\u0069\u006d\u004f\u006e\u0065":
			*_gebcd = 8
		case "\u0061n\u0069\u006d\u004c\u0076\u006c":
			*_gebcd = 9
		case "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073":
			*_gebcd = 10
		}
	}
	_adeba, _dcggf = d.Token()
	if _dcggf != nil {
		return _dcggf
	}
	if _cgcacb, _eaedgf := _adeba.(_a.EndElement); _eaedgf && _cgcacb.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _adeba)
}

type CT_DiagramDefinitionHeader struct {
	UniqueIdAttr string
	MinVerAttr   *string
	DefStyleAttr *string
	ResIdAttr    *int32
	Title        []*CT_Name
	Desc         []*CT_Description
	CatLst       *CT_Categories
	ExtLst       *_bg.CT_OfficeArtExtensionList
}

type CT_ColorTransformHeaderLst struct{ ColorsDefHdr []*CT_ColorTransformHeader }

func NewCT_PtList() *CT_PtList { _cgge := &CT_PtList{}; return _cgge }

func (_edebe ST_OutputShapeType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_edebe.String(), start)
}

func (_eaac *StyleDefHdr) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_eaac.CT_StyleDefinitionHeader = *NewCT_StyleDefinitionHeader()
	for _, _fcaee := range start.Attr {
		if _fcaee.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_cgdcd, _fgcd := _fcaee.Value, error(nil)
			if _fgcd != nil {
				return _fgcd
			}
			_eaac.UniqueIdAttr = _cgdcd
			continue
		}
		if _fcaee.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_gbdg, _cbfg := _fcaee.Value, error(nil)
			if _cbfg != nil {
				return _cbfg
			}
			_eaac.MinVerAttr = &_gbdg
			continue
		}
		if _fcaee.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_bebed, _dbbc := _ae.ParseInt(_fcaee.Value, 10, 32)
			if _dbbc != nil {
				return _dbbc
			}
			_afe := int32(_bebed)
			_eaac.ResIdAttr = &_afe
			continue
		}
	}
_fdabb:
	for {
		_bcfe, _gfdgdf := d.Token()
		if _gfdgdf != nil {
			return _gfdgdf
		}
		switch _cagab := _bcfe.(type) {
		case _a.StartElement:
			switch _cagab.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_dffe := NewCT_SDName()
				if _cbeec := d.DecodeElement(_dffe, &_cagab); _cbeec != nil {
					return _cbeec
				}
				_eaac.Title = append(_eaac.Title, _dffe)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_fced := NewCT_SDDescription()
				if _aeeeb := d.DecodeElement(_fced, &_cagab); _aeeeb != nil {
					return _aeeeb
				}
				_eaac.Desc = append(_eaac.Desc, _fced)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_eaac.CatLst = NewCT_SDCategories()
				if _cdff := d.DecodeElement(_eaac.CatLst, &_cagab); _cdff != nil {
					return _cdff
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_eaac.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _aecbd := d.DecodeElement(_eaac.ExtLst, &_cagab); _aecbd != nil {
					return _aecbd
				}
			default:
				_b.Log.Debug("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072\u0020\u0025\u0076", _cagab.Name)
				if _efdca := d.Skip(); _efdca != nil {
					return _efdca
				}
			}
		case _a.EndElement:
			break _fdabb
		case _a.CharData:
		}
	}
	return nil
}

const (
	ST_FunctionTypeUnset    ST_FunctionType = 0
	ST_FunctionTypeCnt      ST_FunctionType = 1
	ST_FunctionTypePos      ST_FunctionType = 2
	ST_FunctionTypeRevPos   ST_FunctionType = 3
	ST_FunctionTypePosEven  ST_FunctionType = 4
	ST_FunctionTypePosOdd   ST_FunctionType = 5
	ST_FunctionTypeVar      ST_FunctionType = 6
	ST_FunctionTypeDepth    ST_FunctionType = 7
	ST_FunctionTypeMaxDepth ST_FunctionType = 8
)

func (_aegbb ST_SecondaryChildAlignment) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_aegbb.String(), start)
}

const (
	ST_ConstraintRelationshipUnset ST_ConstraintRelationship = 0
	ST_ConstraintRelationshipSelf  ST_ConstraintRelationship = 1
	ST_ConstraintRelationshipCh    ST_ConstraintRelationship = 2
	ST_ConstraintRelationshipDes   ST_ConstraintRelationship = 3
)

func NewCT_LayoutNode() *CT_LayoutNode { _edgd := &CT_LayoutNode{}; return _edgd }

func (_bfggd ST_HierBranchStyle) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_bfggd.String(), start)
}

func (_ddebg ST_PyramidAccentPosition) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_feedf := _a.Attr{}
	_feedf.Name = name
	switch _ddebg {
	case ST_PyramidAccentPositionUnset:
		_feedf.Value = ""
	case ST_PyramidAccentPositionBef:
		_feedf.Value = "\u0062\u0065\u0066"
	case ST_PyramidAccentPositionAft:
		_feedf.Value = "\u0061\u0066\u0074"
	}
	return _feedf, nil
}

func (_dgee ST_Breakpoint) Validate() error { return _dgee.ValidateWithPath("") }

func (_aeag *ST_CxnType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_aeag = 0
	case "\u0070\u0061\u0072O\u0066":
		*_aeag = 1
	case "\u0070\u0072\u0065\u0073\u004f\u0066":
		*_aeag = 2
	case "\u0070r\u0065\u0073\u0050\u0061\u0072\u004ff":
		*_aeag = 3
	case "\u0075\u006e\u006b\u006eow\u006e\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070":
		*_aeag = 4
	}
	return nil
}

func (_fbbd ST_AlgorithmType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_fbbd.String(), start)
}

func (_eecf ST_AnimOneStr) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_eecf.String(), start)
}

// Validate validates the CT_Parameter and its children
func (_abfa *CT_Parameter) Validate() error {
	return _abfa.ValidateWithPath("\u0043\u0054\u005fP\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072")
}

const (
	ST_AlgorithmTypeUnset     ST_AlgorithmType = 0
	ST_AlgorithmTypeComposite ST_AlgorithmType = 1
	ST_AlgorithmTypeConn      ST_AlgorithmType = 2
	ST_AlgorithmTypeCycle     ST_AlgorithmType = 3
	ST_AlgorithmTypeHierChild ST_AlgorithmType = 4
	ST_AlgorithmTypeHierRoot  ST_AlgorithmType = 5
	ST_AlgorithmTypePyra      ST_AlgorithmType = 6
	ST_AlgorithmTypeLin       ST_AlgorithmType = 7
	ST_AlgorithmTypeSp        ST_AlgorithmType = 8
	ST_AlgorithmTypeTx        ST_AlgorithmType = 9
	ST_AlgorithmTypeSnake     ST_AlgorithmType = 10
)

func (_bbebf *ST_NodeVerticalAlignment) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_bcafd, _fgbd := d.Token()
	if _fgbd != nil {
		return _fgbd
	}
	if _cege, _bdegd := _bcafd.(_a.EndElement); _bdegd && _cege.Name == start.Name {
		*_bbebf = 1
		return nil
	}
	if _eceac, _acdfe := _bcafd.(_a.CharData); !_acdfe {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bcafd)
	} else {
		switch string(_eceac) {
		case "":
			*_bbebf = 0
		case "\u0074":
			*_bbebf = 1
		case "\u006d\u0069\u0064":
			*_bbebf = 2
		case "\u0062":
			*_bbebf = 3
		}
	}
	_bcafd, _fgbd = d.Token()
	if _fgbd != nil {
		return _fgbd
	}
	if _face, _ddec := _bcafd.(_a.EndElement); _ddec && _face.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bcafd)
}

func (_gedag ST_OutputShapeType) ValidateWithPath(path string) error {
	switch _gedag {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gedag))
	}
	return nil
}

// ValidateWithPath validates the CT_StyleDefinitionHeader and its children, prefixing error messages with path
func (_bgdb *CT_StyleDefinitionHeader) ValidateWithPath(path string) error {
	for _cgcgb, _bcdc := range _bgdb.Title {
		if _afge := _bcdc.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _cgcgb)); _afge != nil {
			return _afge
		}
	}
	for _aagad, _fddaa := range _bgdb.Desc {
		if _dafff := _fddaa.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _aagad)); _dafff != nil {
			return _dafff
		}
	}
	if _bgdb.CatLst != nil {
		if _eaca := _bgdb.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _eaca != nil {
			return _eaca
		}
	}
	if _bgdb.ExtLst != nil {
		if _ggede := _bgdb.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _ggede != nil {
			return _ggede
		}
	}
	return nil
}

type CT_Pt struct {
	ModelIdAttr ST_ModelId
	TypeAttr    ST_PtType
	CxnIdAttr   *ST_ModelId
	PrSet       *CT_ElemPropSet
	SpPr        *_bg.CT_ShapeProperties
	T           *_bg.CT_TextBody
	ExtLst      *_bg.CT_OfficeArtExtensionList
}

func (_dedfee *CT_TextProps) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_eeff:
	for {
		_gdbb, _dbcfe := d.Token()
		if _dbcfe != nil {
			return _dbcfe
		}
		switch _agbf := _gdbb.(type) {
		case _a.StartElement:
			switch _agbf.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0070\u0033\u0064"}, _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072g\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u006d\u0061\u0069\u006e", Local: "\u0073\u0070\u0033\u0064"}:
				_dedfee.Sp3d = _bg.NewCT_Shape3D()
				if _gfce := d.DecodeElement(_dedfee.Sp3d, &_agbf); _gfce != nil {
					return _gfce
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e", Local: "\u0066\u006c\u0061\u0074\u0054\u0078"}, _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0063\u006c\u0063\u002e\u006f\u0072g\u002f\u006f\u006f\u0078\u006d\u006c\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u006d\u0061\u0069\u006e", Local: "\u0066\u006c\u0061\u0074\u0054\u0078"}:
				_dedfee.FlatTx = _bg.NewCT_FlatText()
				if _dadf := d.DecodeElement(_dedfee.FlatTx, &_agbf); _dadf != nil {
					return _dadf
				}
			default:
				_b.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_T\u0065\u0078t\u0050\u0072\u006f\u0070\u0073\u0020\u0025\u0076", _agbf.Name)
				if _fgcf := d.Skip(); _fgcf != nil {
					return _fgcf
				}
			}
		case _a.EndElement:
			break _eeff
		case _a.CharData:
		}
	}
	return nil
}

func (_bagef ST_PyramidAccentPosition) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_bagef.String(), start)
}

func (_gbcd ST_CxnType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_gbcd.String(), start)
}

// Validate validates the CT_LayoutVariablePropertySet and its children
func (_fedgd *CT_LayoutVariablePropertySet) Validate() error {
	return _fedgd.ValidateWithPath("\u0043\u0054\u005f\u004ca\u0079\u006f\u0075\u0074\u0056\u0061\u0072\u0069\u0061\u0062l\u0065P\u0072\u006f\u0070\u0065\u0072\u0074\u0079S\u0065\u0074")
}

// Validate validates the CT_Category and its children
func (_efc *CT_Category) Validate() error {
	return _efc.ValidateWithPath("C\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079")
}

func (_gdec ST_AutoTextRotation) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_gafdc := _a.Attr{}
	_gafdc.Name = name
	switch _gdec {
	case ST_AutoTextRotationUnset:
		_gafdc.Value = ""
	case ST_AutoTextRotationNone:
		_gafdc.Value = "\u006e\u006f\u006e\u0065"
	case ST_AutoTextRotationUpr:
		_gafdc.Value = "\u0075\u0070\u0072"
	case ST_AutoTextRotationGrav:
		_gafdc.Value = "\u0067\u0072\u0061\u0076"
	}
	return _gafdc, nil
}

func NewCT_HierBranchStyle() *CT_HierBranchStyle { _gefbb := &CT_HierBranchStyle{}; return _gefbb }

func (_gdgb ST_DiagramTextAlignment) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_gdgb.String(), start)
}

func NewCT_AnimOne() *CT_AnimOne { _gccc := &CT_AnimOne{}; return _gccc }

func NewCT_Description() *CT_Description { _fccef := &CT_Description{}; return _fccef }

func (_fdac ST_TextDirection) ValidateWithPath(path string) error {
	switch _fdac {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fdac))
	}
	return nil
}

// Validate validates the CT_Algorithm and its children
func (_agb *CT_Algorithm) Validate() error {
	return _agb.ValidateWithPath("\u0043\u0054\u005fA\u006c\u0067\u006f\u0072\u0069\u0074\u0068\u006d")
}

type ST_ConnectorPoint byte

func (_dcdaf ST_GrowDirection) Validate() error { return _dcdaf.ValidateWithPath("") }

func NewLayoutDefHdrLst() *LayoutDefHdrLst {
	_bfffc := &LayoutDefHdrLst{}
	_bfffc.CT_DiagramDefinitionHeaderLst = *NewCT_DiagramDefinitionHeaderLst()
	return _bfffc
}

func (_edbc ST_CenterShapeMapping) Validate() error { return _edbc.ValidateWithPath("") }

func (_dfacg *ST_ArrowheadStyle) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_cage, _gcaf := d.Token()
	if _gcaf != nil {
		return _gcaf
	}
	if _adggg, _dcfaa := _cage.(_a.EndElement); _dcfaa && _adggg.Name == start.Name {
		*_dfacg = 1
		return nil
	}
	if _cbdf, _ebgbf := _cage.(_a.CharData); !_ebgbf {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cage)
	} else {
		switch string(_cbdf) {
		case "":
			*_dfacg = 0
		case "\u0061\u0075\u0074\u006f":
			*_dfacg = 1
		case "\u0061\u0072\u0072":
			*_dfacg = 2
		case "\u006e\u006f\u0041r\u0072":
			*_dfacg = 3
		}
	}
	_cage, _gcaf = d.Token()
	if _gcaf != nil {
		return _gcaf
	}
	if _afgeg, _ggdce := _cage.(_a.EndElement); _ggdce && _afgeg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cage)
}

func (_gcafb ST_ConnectorDimension) ValidateWithPath(path string) error {
	switch _gcafb {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gcafb))
	}
	return nil
}

type ST_AutoTextRotation byte

func _feae(_dgbbf bool) uint8 {
	if _dgbbf {
		return 1
	}
	return 0
}

type CT_Colors struct {
	MethAttr       ST_ClrAppMethod
	HueDirAttr     ST_HueDir
	EG_ColorChoice []*_bg.EG_ColorChoice
}

func (_fddbec *DataModel) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0064a\u0074\u0061\u004d\u006f\u0064\u0065l"
	return _fddbec.CT_DataModel.MarshalXML(e, start)
}

type CT_SDDescription struct {
	LangAttr *string
	ValAttr  string
}

func (_ceffg *CT_RelIds) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u003a\u0064\u006d"}, Value: _f.Sprintf("\u0025\u0076", _ceffg.DmAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u003a\u006c\u006f"}, Value: _f.Sprintf("\u0025\u0076", _ceffg.LoAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u003a\u0071\u0073"}, Value: _f.Sprintf("\u0025\u0076", _ceffg.QsAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u003a\u0063\u0073"}, Value: _f.Sprintf("\u0025\u0076", _ceffg.CsAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

type ST_Offset byte

func (_ceecd *ST_FunctionOperator) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_dbdac, _cbfa := d.Token()
	if _cbfa != nil {
		return _cbfa
	}
	if _bbcag, _fggd := _dbdac.(_a.EndElement); _fggd && _bbcag.Name == start.Name {
		*_ceecd = 1
		return nil
	}
	if _bbade, _dbgc := _dbdac.(_a.CharData); !_dbgc {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dbdac)
	} else {
		switch string(_bbade) {
		case "":
			*_ceecd = 0
		case "\u0065\u0071\u0075":
			*_ceecd = 1
		case "\u006e\u0065\u0071":
			*_ceecd = 2
		case "\u0067\u0074":
			*_ceecd = 3
		case "\u006c\u0074":
			*_ceecd = 4
		case "\u0067\u0074\u0065":
			*_ceecd = 5
		case "\u006c\u0074\u0065":
			*_ceecd = 6
		}
	}
	_dbdac, _cbfa = d.Token()
	if _cbfa != nil {
		return _cbfa
	}
	if _dfgg, _dbdce := _dbdac.(_a.EndElement); _dbdce && _dfgg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dbdac)
}

func (_bgaec *ST_HierarchyAlignment) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gcbg, _effa := d.Token()
	if _effa != nil {
		return _effa
	}
	if _bdfda, _afbfa := _gcbg.(_a.EndElement); _afbfa && _bdfda.Name == start.Name {
		*_bgaec = 1
		return nil
	}
	if _cgedd, _ccefb := _gcbg.(_a.CharData); !_ccefb {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gcbg)
	} else {
		switch string(_cgedd) {
		case "":
			*_bgaec = 0
		case "\u0074\u004c":
			*_bgaec = 1
		case "\u0074\u0052":
			*_bgaec = 2
		case "\u0074\u0043\u0074\u0072\u0043\u0068":
			*_bgaec = 3
		case "\u0074C\u0074\u0072\u0044\u0065\u0073":
			*_bgaec = 4
		case "\u0062\u004c":
			*_bgaec = 5
		case "\u0062\u0052":
			*_bgaec = 6
		case "\u0062\u0043\u0074\u0072\u0043\u0068":
			*_bgaec = 7
		case "\u0062C\u0074\u0072\u0044\u0065\u0073":
			*_bgaec = 8
		case "\u006c\u0054":
			*_bgaec = 9
		case "\u006c\u0042":
			*_bgaec = 10
		case "\u006c\u0043\u0074\u0072\u0043\u0068":
			*_bgaec = 11
		case "\u006cC\u0074\u0072\u0044\u0065\u0073":
			*_bgaec = 12
		case "\u0072\u0054":
			*_bgaec = 13
		case "\u0072\u0042":
			*_bgaec = 14
		case "\u0072\u0043\u0074\u0072\u0043\u0068":
			*_bgaec = 15
		case "\u0072C\u0074\u0072\u0044\u0065\u0073":
			*_bgaec = 16
		}
	}
	_gcbg, _effa = d.Token()
	if _effa != nil {
		return _effa
	}
	if _cfdff, _fggde := _gcbg.(_a.EndElement); _fggde && _cfdff.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gcbg)
}

func (_fcece *ST_ConnectorPoint) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gbddg, _gfbgg := d.Token()
	if _gfbgg != nil {
		return _gfbgg
	}
	if _fbeff, _gecbg := _gbddg.(_a.EndElement); _gecbg && _fbeff.Name == start.Name {
		*_fcece = 1
		return nil
	}
	if _cfegd, _ddedg := _gbddg.(_a.CharData); !_ddedg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gbddg)
	} else {
		switch string(_cfegd) {
		case "":
			*_fcece = 0
		case "\u0061\u0075\u0074\u006f":
			*_fcece = 1
		case "\u0062\u0043\u0074\u0072":
			*_fcece = 2
		case "\u0063\u0074\u0072":
			*_fcece = 3
		case "\u006d\u0069\u0064\u004c":
			*_fcece = 4
		case "\u006d\u0069\u0064\u0052":
			*_fcece = 5
		case "\u0074\u0043\u0074\u0072":
			*_fcece = 6
		case "\u0062\u004c":
			*_fcece = 7
		case "\u0062\u0052":
			*_fcece = 8
		case "\u0074\u004c":
			*_fcece = 9
		case "\u0074\u0052":
			*_fcece = 10
		case "\u0072\u0061\u0064\u0069\u0061\u006c":
			*_fcece = 11
		}
	}
	_gbddg, _gfbgg = d.Token()
	if _gfbgg != nil {
		return _gfbgg
	}
	if _gface, _fegb := _gbddg.(_a.EndElement); _fegb && _gface.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gbddg)
}

// ST_FunctionArgument is a union type
type ST_FunctionArgument struct{ ST_VariableType ST_VariableType }

func (_fcdfc *LayoutDef) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_fcdfc.CT_DiagramDefinition = *NewCT_DiagramDefinition()
	for _, _ggbg := range start.Attr {
		if _ggbg.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_adcga, _fddbef := _ggbg.Value, error(nil)
			if _fddbef != nil {
				return _fddbef
			}
			_fcdfc.UniqueIdAttr = &_adcga
			continue
		}
		if _ggbg.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_efea, _eaff := _ggbg.Value, error(nil)
			if _eaff != nil {
				return _eaff
			}
			_fcdfc.MinVerAttr = &_efea
			continue
		}
		if _ggbg.Name.Local == "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065" {
			_ecgfb, _abaaf := _ggbg.Value, error(nil)
			if _abaaf != nil {
				return _abaaf
			}
			_fcdfc.DefStyleAttr = &_ecgfb
			continue
		}
	}
_adac:
	for {
		_fbeb, _eeae := d.Token()
		if _eeae != nil {
			return _eeae
		}
		switch _affde := _fbeb.(type) {
		case _a.StartElement:
			switch _affde.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_aeaa := NewCT_Name()
				if _ddccd := d.DecodeElement(_aeaa, &_affde); _ddccd != nil {
					return _ddccd
				}
				_fcdfc.Title = append(_fcdfc.Title, _aeaa)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_gdfgg := NewCT_Description()
				if _cgcc := d.DecodeElement(_gdfgg, &_affde); _cgcc != nil {
					return _cgcc
				}
				_fcdfc.Desc = append(_fcdfc.Desc, _gdfgg)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_fcdfc.CatLst = NewCT_Categories()
				if _ddcf := d.DecodeElement(_fcdfc.CatLst, &_affde); _ddcf != nil {
					return _ddcf
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0061\u006d\u0070\u0044\u0061\u0074\u0061"}:
				_fcdfc.SampData = NewCT_SampleData()
				if _bggad := d.DecodeElement(_fcdfc.SampData, &_affde); _bggad != nil {
					return _bggad
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073t\u0079\u006c\u0065\u0044\u0061\u0074a"}:
				_fcdfc.StyleData = NewCT_SampleData()
				if _ffca := d.DecodeElement(_fcdfc.StyleData, &_affde); _ffca != nil {
					return _ffca
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063l\u0072\u0044\u0061\u0074\u0061"}:
				_fcdfc.ClrData = NewCT_SampleData()
				if _ddbg := d.DecodeElement(_fcdfc.ClrData, &_affde); _ddbg != nil {
					return _ddbg
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				if _abbc := d.DecodeElement(_fcdfc.LayoutNode, &_affde); _abbc != nil {
					return _abbc
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_fcdfc.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _ffgga := d.DecodeElement(_fcdfc.ExtLst, &_affde); _ffgga != nil {
					return _ffgga
				}
			default:
				_b.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u004c\u0061\u0079\u006f\u0075\u0074D\u0065\u0066 \u0025\u0076", _affde.Name)
				if _daebe := d.Skip(); _daebe != nil {
					return _daebe
				}
			}
		case _a.EndElement:
			break _adac
		case _a.CharData:
		}
	}
	return nil
}

func (_ggbeb ST_Offset) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_ggbeb.String(), start)
}

func NewColorsDef() *ColorsDef {
	_gfef := &ColorsDef{}
	_gfef.CT_ColorTransform = *NewCT_ColorTransform()
	return _gfef
}

func (_babcf ST_ConstraintType) String() string {
	switch _babcf {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0061\u006c\u0069\u0067\u006e\u004f\u0066\u0066"
	case 3:
		return "\u0062e\u0067\u004d\u0061\u0072\u0067"
	case 4:
		return "\u0062\u0065\u006e\u0064\u0044\u0069\u0073\u0074"
	case 5:
		return "\u0062\u0065\u0067\u0050\u0061\u0064"
	case 6:
		return "\u0062"
	case 7:
		return "\u0062\u004d\u0061r\u0067"
	case 8:
		return "\u0062\u004f\u0066\u0066"
	case 9:
		return "\u0063\u0074\u0072\u0058"
	case 10:
		return "\u0063t\u0072\u0058\u004f\u0066\u0066"
	case 11:
		return "\u0063\u0074\u0072\u0059"
	case 12:
		return "\u0063t\u0072\u0059\u004f\u0066\u0066"
	case 13:
		return "\u0063\u006f\u006e\u006e\u0044\u0069\u0073\u0074"
	case 14:
		return "\u0064\u0069\u0061\u006d"
	case 15:
		return "\u0065n\u0064\u004d\u0061\u0072\u0067"
	case 16:
		return "\u0065\u006e\u0064\u0050\u0061\u0064"
	case 17:
		return "\u0068"
	case 18:
		return "\u0068\u0041\u0072\u0048"
	case 19:
		return "\u0068\u004f\u0066\u0066"
	case 20:
		return "\u006c"
	case 21:
		return "\u006c\u004d\u0061r\u0067"
	case 22:
		return "\u006c\u004f\u0066\u0066"
	case 23:
		return "\u0072"
	case 24:
		return "\u0072\u004d\u0061r\u0067"
	case 25:
		return "\u0072\u004f\u0066\u0066"
	case 26:
		return "\u0070\u0072\u0069\u006d\u0046\u006f\u006e\u0074\u0053\u007a"
	case 27:
		return "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0052\u0061\u0074\u0069\u006f"
	case 28:
		return "\u0073e\u0063\u0046\u006f\u006e\u0074\u0053z"
	case 29:
		return "\u0073\u0069\u0062S\u0070"
	case 30:
		return "\u0073\u0065\u0063\u0053\u0069\u0062\u0053\u0070"
	case 31:
		return "\u0073\u0070"
	case 32:
		return "\u0073t\u0065\u006d\u0054\u0068\u0069\u0063k"
	case 33:
		return "\u0074"
	case 34:
		return "\u0074\u004d\u0061r\u0067"
	case 35:
		return "\u0074\u004f\u0066\u0066"
	case 36:
		return "\u0075\u0073\u0065r\u0041"
	case 37:
		return "\u0075\u0073\u0065r\u0042"
	case 38:
		return "\u0075\u0073\u0065r\u0043"
	case 39:
		return "\u0075\u0073\u0065r\u0044"
	case 40:
		return "\u0075\u0073\u0065r\u0045"
	case 41:
		return "\u0075\u0073\u0065r\u0046"
	case 42:
		return "\u0075\u0073\u0065r\u0047"
	case 43:
		return "\u0075\u0073\u0065r\u0048"
	case 44:
		return "\u0075\u0073\u0065r\u0049"
	case 45:
		return "\u0075\u0073\u0065r\u004a"
	case 46:
		return "\u0075\u0073\u0065r\u004b"
	case 47:
		return "\u0075\u0073\u0065r\u004c"
	case 48:
		return "\u0075\u0073\u0065r\u004d"
	case 49:
		return "\u0075\u0073\u0065r\u004e"
	case 50:
		return "\u0075\u0073\u0065r\u004f"
	case 51:
		return "\u0075\u0073\u0065r\u0050"
	case 52:
		return "\u0075\u0073\u0065r\u0051"
	case 53:
		return "\u0075\u0073\u0065r\u0052"
	case 54:
		return "\u0075\u0073\u0065r\u0053"
	case 55:
		return "\u0075\u0073\u0065r\u0054"
	case 56:
		return "\u0075\u0073\u0065r\u0055"
	case 57:
		return "\u0075\u0073\u0065r\u0056"
	case 58:
		return "\u0075\u0073\u0065r\u0057"
	case 59:
		return "\u0075\u0073\u0065r\u0058"
	case 60:
		return "\u0075\u0073\u0065r\u0059"
	case 61:
		return "\u0075\u0073\u0065r\u005a"
	case 62:
		return "\u0077"
	case 63:
		return "\u0077\u0041\u0072\u0048"
	case 64:
		return "\u0077\u004f\u0066\u0066"
	}
	return ""
}

func NewCT_ChildPref() *CT_ChildPref { _dbd := &CT_ChildPref{}; return _dbd }

// ValidateWithPath validates the CT_CxnList and its children, prefixing error messages with path
func (_fcac *CT_CxnList) ValidateWithPath(path string) error {
	for _adaf, _gadg := range _fcac.Cxn {
		if _fccd := _gadg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0078\u006e\u005b\u0025\u0064\u005d", path, _adaf)); _fccd != nil {
			return _fccd
		}
	}
	return nil
}

func (_aecbcf ST_BoolOperator) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_aecbcf.String(), start)
}

func (_bdeg *ST_AnimOneStr) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gfgec, _dfffa := d.Token()
	if _dfffa != nil {
		return _dfffa
	}
	if _fgee, _begbd := _gfgec.(_a.EndElement); _begbd && _fgee.Name == start.Name {
		*_bdeg = 1
		return nil
	}
	if _cfed, _efaaf := _gfgec.(_a.CharData); !_efaaf {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gfgec)
	} else {
		switch string(_cfed) {
		case "":
			*_bdeg = 0
		case "\u006e\u006f\u006e\u0065":
			*_bdeg = 1
		case "\u006f\u006e\u0065":
			*_bdeg = 2
		case "\u0062\u0072\u0061\u006e\u0063\u0068":
			*_bdeg = 3
		}
	}
	_gfgec, _dfffa = d.Token()
	if _dfffa != nil {
		return _dfffa
	}
	if _gedc, _decb := _gfgec.(_a.EndElement); _decb && _gedc.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gfgec)
}

// ValidateWithPath validates the CT_Choose and its children, prefixing error messages with path
func (_fee *CT_Choose) ValidateWithPath(path string) error {
	for _bbgd, _ccd := range _fee.If {
		if _dgbe := _ccd.ValidateWithPath(_f.Sprintf("\u0025s\u002f\u0049\u0066\u005b\u0025\u0064]", path, _bbgd)); _dgbe != nil {
			return _dgbe
		}
	}
	if _fee.Else != nil {
		if _ddfb := _fee.Else.ValidateWithPath(path + "\u002f\u0045\u006cs\u0065"); _ddfb != nil {
			return _ddfb
		}
	}
	return nil
}

type CT_Constraints struct{ Constr []*CT_Constraint }

func (_gbddb ST_AnimOneStr) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_edbd := _a.Attr{}
	_edbd.Name = name
	switch _gbddb {
	case ST_AnimOneStrUnset:
		_edbd.Value = ""
	case ST_AnimOneStrNone:
		_edbd.Value = "\u006e\u006f\u006e\u0065"
	case ST_AnimOneStrOne:
		_edbd.Value = "\u006f\u006e\u0065"
	case ST_AnimOneStrBranch:
		_edbd.Value = "\u0062\u0072\u0061\u006e\u0063\u0068"
	}
	return _edbd, nil
}

const (
	ST_ConnectorDimensionUnset ST_ConnectorDimension = 0
	ST_ConnectorDimension1D    ST_ConnectorDimension = 1
	ST_ConnectorDimension2D    ST_ConnectorDimension = 2
	ST_ConnectorDimensionCust  ST_ConnectorDimension = 3
)

func (_gfdge ST_ConnectorPoint) ValidateWithPath(path string) error {
	switch _gfdge {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gfdge))
	}
	return nil
}

// Validate validates the CT_SDCategory and its children
func (_ffdfb *CT_SDCategory) Validate() error {
	return _ffdfb.ValidateWithPath("\u0043\u0054\u005f\u0053\u0044\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079")
}

func NewStyleDefHdrLst() *StyleDefHdrLst {
	_ceeed := &StyleDefHdrLst{}
	_ceeed.CT_StyleDefinitionHeaderLst = *NewCT_StyleDefinitionHeaderLst()
	return _ceeed
}

type ST_GrowDirection byte

// ValidateWithPath validates the CT_Name and its children, prefixing error messages with path
func (_efb *CT_Name) ValidateWithPath(path string) error { return nil }

func (_fged *CT_CTName) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _afce := range start.Attr {
		if _afce.Name.Local == "\u006c\u0061\u006e\u0067" {
			_dcg, _fbb := _afce.Value, error(nil)
			if _fbb != nil {
				return _fbb
			}
			_fged.LangAttr = &_dcg
			continue
		}
		if _afce.Name.Local == "\u0076\u0061\u006c" {
			_bdf, _eca := _afce.Value, error(nil)
			if _eca != nil {
				return _eca
			}
			_fged.ValAttr = _bdf
			continue
		}
	}
	for {
		_fgbf, _dfe := d.Token()
		if _dfe != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0043\u0054\u004ea\u006d\u0065\u003a\u0020\u0025\u0073", _dfe)
		}
		if _ebdg, _dac := _fgbf.(_a.EndElement); _dac && _ebdg.Name == start.Name {
			break
		}
	}
	return nil
}

func (_bbde ST_ParameterId) ValidateWithPath(path string) error {
	switch _bbde {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bbde))
	}
	return nil
}

func (_dgbeb ST_TextBlockDirection) String() string {
	switch _dgbeb {
	case 0:
		return ""
	case 1:
		return "\u0068\u006f\u0072\u007a"
	case 2:
		return "\u0076\u0065\u0072\u0074"
	}
	return ""
}

type ST_DiagramHorizontalAlignment byte

type ST_HierarchyAlignment byte

func (_bccb ST_Direction) String() string {
	switch _bccb {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u0072\u006d"
	case 2:
		return "\u0072\u0065\u0076"
	}
	return ""
}

// Validate validates the CT_CTCategory and its children
func (_afa *CT_CTCategory) Validate() error {
	return _afa.ValidateWithPath("\u0043\u0054\u005f\u0043\u0054\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079")
}

func (_dcef ST_ConstraintRelationship) String() string {
	switch _dcef {
	case 0:
		return ""
	case 1:
		return "\u0073\u0065\u006c\u0066"
	case 2:
		return "\u0063\u0068"
	case 3:
		return "\u0064\u0065\u0073"
	}
	return ""
}

func (_cdbd ST_ConstraintType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_cdbd.String(), start)
}

func (_effb *ST_ConnectorPoint) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_effb = 0
	case "\u0061\u0075\u0074\u006f":
		*_effb = 1
	case "\u0062\u0043\u0074\u0072":
		*_effb = 2
	case "\u0063\u0074\u0072":
		*_effb = 3
	case "\u006d\u0069\u0064\u004c":
		*_effb = 4
	case "\u006d\u0069\u0064\u0052":
		*_effb = 5
	case "\u0074\u0043\u0074\u0072":
		*_effb = 6
	case "\u0062\u004c":
		*_effb = 7
	case "\u0062\u0052":
		*_effb = 8
	case "\u0074\u004c":
		*_effb = 9
	case "\u0074\u0052":
		*_effb = 10
	case "\u0072\u0061\u0064\u0069\u0061\u006c":
		*_effb = 11
	}
	return nil
}

func NewCT_Colors() *CT_Colors { _egff := &CT_Colors{}; return _egff }

func (_cdgf ST_LinearDirection) ValidateWithPath(path string) error {
	switch _cdgf {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cdgf))
	}
	return nil
}

func (_aade ST_FallbackDimension) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_aade.String(), start)
}

// ValidateWithPath validates the CT_When and its children, prefixing error messages with path
func (_acac *CT_When) ValidateWithPath(path string) error {
	if _acac.FuncAttr == ST_FunctionTypeUnset {
		return _f.Errorf("\u0025\u0073\u002f\u0046\u0075\u006e\u0063\u0041\u0074\u0074\u0072\u0020\u0069\u0073\u0020a\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _aecf := _acac.FuncAttr.ValidateWithPath(path + "\u002fF\u0075\u006e\u0063\u0041\u0074\u0074r"); _aecf != nil {
		return _aecf
	}
	if _acac.ArgAttr != nil {
		if _gcedg := _acac.ArgAttr.ValidateWithPath(path + "\u002f\u0041\u0072\u0067\u0041\u0074\u0074\u0072"); _gcedg != nil {
			return _gcedg
		}
	}
	if _acac.OpAttr == ST_FunctionOperatorUnset {
		return _f.Errorf("\u0025\u0073\u002f\u004f\u0070\u0041\u0074\u0074\u0072\u0020i\u0073\u0020\u0061\u0020\u006d\u0061\u006ed\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _ffb := _acac.OpAttr.ValidateWithPath(path + "\u002fO\u0070\u0041\u0074\u0074\u0072"); _ffb != nil {
		return _ffb
	}
	if _abad := _acac.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _abad != nil {
		return _abad
	}
	for _ggcg, _ddbb := range _acac.Alg {
		if _bgeg := _ddbb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0041\u006c\u0067\u005b\u0025\u0064\u005d", path, _ggcg)); _bgeg != nil {
			return _bgeg
		}
	}
	for _aefcfe, _efdfb := range _acac.Shape {
		if _gfadg := _efdfb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fS\u0068\u0061\u0070\u0065\u005b\u0025\u0064\u005d", path, _aefcfe)); _gfadg != nil {
			return _gfadg
		}
	}
	for _fcfd, _dcea := range _acac.PresOf {
		if _fedcd := _dcea.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0050\u0072\u0065\u0073\u004f\u0066\u005b\u0025\u0064\u005d", path, _fcfd)); _fedcd != nil {
			return _fedcd
		}
	}
	for _dbdge, _agab := range _acac.ConstrLst {
		if _fgfce := _agab.ValidateWithPath(_f.Sprintf("\u0025\u0073/\u0043\u006f\u006es\u0074\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _dbdge)); _fgfce != nil {
			return _fgfce
		}
	}
	for _afaac, _bdfg := range _acac.RuleLst {
		if _ebgc := _bdfg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0052\u0075\u006c\u0065\u004c\u0073t\u005b\u0025\u0064\u005d", path, _afaac)); _ebgc != nil {
			return _ebgc
		}
	}
	for _cdfe, _gaaf := range _acac.ForEach {
		if _cacae := _gaaf.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0046\u006f\u0072\u0045\u0061\u0063h\u005b\u0025\u0064\u005d", path, _cdfe)); _cacae != nil {
			return _cacae
		}
	}
	for _bbaad, _cdcg := range _acac.LayoutNode {
		if _eaceg := _cdcg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064e\u005b\u0025\u0064\u005d", path, _bbaad)); _eaceg != nil {
			return _eaceg
		}
	}
	for _dcba, _bgebc := range _acac.Choose {
		if _efca := _bgebc.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u006f\u0073\u0065\u005b\u0025\u0064\u005d", path, _dcba)); _efca != nil {
			return _efca
		}
	}
	for _gbee, _eecgg := range _acac.ExtLst {
		if _aebb := _eecgg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0045\u0078\u0074\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _gbee)); _aebb != nil {
			return _aebb
		}
	}
	return nil
}

func (_fbgd ST_ConnectorRouting) ValidateWithPath(path string) error {
	switch _fbgd {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fbgd))
	}
	return nil
}

func (_fcccf ST_PyramidAccentTextMargin) Validate() error { return _fcccf.ValidateWithPath("") }

func NewCT_StyleDefinition() *CT_StyleDefinition { _cacfd := &CT_StyleDefinition{}; return _cacfd }

func (_gfbd ST_RotationPath) ValidateWithPath(path string) error {
	switch _gfbd {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gfbd))
	}
	return nil
}

const (
	ST_OffsetUnset ST_Offset = 0
	ST_OffsetCtr   ST_Offset = 1
	ST_OffsetOff   ST_Offset = 2
)

func (_dcag ST_NodeHorizontalAlignment) Validate() error { return _dcag.ValidateWithPath("") }

func (_gccda *ST_AutoTextRotation) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_edeg, _edgfa := d.Token()
	if _edgfa != nil {
		return _edgfa
	}
	if _deda, _bbgb := _edeg.(_a.EndElement); _bbgb && _deda.Name == start.Name {
		*_gccda = 1
		return nil
	}
	if _aadbb, _ccebcg := _edeg.(_a.CharData); !_ccebcg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _edeg)
	} else {
		switch string(_aadbb) {
		case "":
			*_gccda = 0
		case "\u006e\u006f\u006e\u0065":
			*_gccda = 1
		case "\u0075\u0070\u0072":
			*_gccda = 2
		case "\u0067\u0072\u0061\u0076":
			*_gccda = 3
		}
	}
	_edeg, _edgfa = d.Token()
	if _edgfa != nil {
		return _edgfa
	}
	if _dcac, _begad := _edeg.(_a.EndElement); _begad && _dcac.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _edeg)
}

func (_aeaba ST_HueDir) Validate() error { return _aeaba.ValidateWithPath("") }

func (_aafc *ST_FlowDirection) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_cfegdc, _bbdc := d.Token()
	if _bbdc != nil {
		return _bbdc
	}
	if _eafe, _fefd := _cfegdc.(_a.EndElement); _fefd && _eafe.Name == start.Name {
		*_aafc = 1
		return nil
	}
	if _abeeeg, _efdff := _cfegdc.(_a.CharData); !_efdff {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cfegdc)
	} else {
		switch string(_abeeeg) {
		case "":
			*_aafc = 0
		case "\u0072\u006f\u0077":
			*_aafc = 1
		case "\u0063\u006f\u006c":
			*_aafc = 2
		}
	}
	_cfegdc, _bbdc = d.Token()
	if _bbdc != nil {
		return _bbdc
	}
	if _fcbcg, _fdcdg := _cfegdc.(_a.EndElement); _fdcdg && _fcbcg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cfegdc)
}

func NewCT_Rules() *CT_Rules { _ggead := &CT_Rules{}; return _ggead }

func (_edcf ST_FunctionOperator) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_edcf.String(), start)
}

type CT_DiagramDefinition struct {
	UniqueIdAttr *string
	MinVerAttr   *string
	DefStyleAttr *string
	Title        []*CT_Name
	Desc         []*CT_Description
	CatLst       *CT_Categories
	SampData     *CT_SampleData
	StyleData    *CT_SampleData
	ClrData      *CT_SampleData
	LayoutNode   *CT_LayoutNode
	ExtLst       *_bg.CT_OfficeArtExtensionList
}

func NewCT_SDDescription() *CT_SDDescription { _febc := &CT_SDDescription{}; return _febc }

func (_ddcfg ST_AnimLvlStr) ValidateWithPath(path string) error {
	switch _ddcfg {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ddcfg))
	}
	return nil
}

func (_cgbb *ST_AnimLvlStr) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gdace, _gfbcf := d.Token()
	if _gfbcf != nil {
		return _gfbcf
	}
	if _gaeg, _eebd := _gdace.(_a.EndElement); _eebd && _gaeg.Name == start.Name {
		*_cgbb = 1
		return nil
	}
	if _fdfcb, _bbfd := _gdace.(_a.CharData); !_bbfd {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gdace)
	} else {
		switch string(_fdfcb) {
		case "":
			*_cgbb = 0
		case "\u006e\u006f\u006e\u0065":
			*_cgbb = 1
		case "\u006c\u0076\u006c":
			*_cgbb = 2
		case "\u0063\u0074\u0072":
			*_cgbb = 3
		}
	}
	_gdace, _gfbcf = d.Token()
	if _gfbcf != nil {
		return _gfbcf
	}
	if _eaea, _fgff := _gdace.(_a.EndElement); _fgff && _eaea.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gdace)
}

type CT_StyleDefinition struct {
	UniqueIdAttr *string
	MinVerAttr   *string
	Title        []*CT_SDName
	Desc         []*CT_SDDescription
	CatLst       *CT_SDCategories
	Scene3d      *_bg.CT_Scene3D
	StyleLbl     []*CT_StyleLabel
	ExtLst       *_bg.CT_OfficeArtExtensionList
}

func NewCT_DiagramDefinitionHeaderLst() *CT_DiagramDefinitionHeaderLst {
	_gfb := &CT_DiagramDefinitionHeaderLst{}
	return _gfb
}

func (_egfd *ST_NodeVerticalAlignment) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_egfd = 0
	case "\u0074":
		*_egfd = 1
	case "\u006d\u0069\u0064":
		*_egfd = 2
	case "\u0062":
		*_egfd = 3
	}
	return nil
}

func (_fdfa ST_VerticalAlignment) String() string {
	switch _fdfa {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u006d\u0069\u0064"
	case 3:
		return "\u0062"
	case 4:
		return "\u006e\u006f\u006e\u0065"
	}
	return ""
}

func (_afeg ST_SecondaryLinearDirection) ValidateWithPath(path string) error {
	switch _afeg {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_afeg))
	}
	return nil
}

// Validate validates the CT_ColorTransform and its children
func (_bbe *CT_ColorTransform) Validate() error {
	return _bbe.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061\u006es\u0066\u006f\u0072\u006d")
}

func (_bcec *CT_Direction) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _fcee := range start.Attr {
		if _fcee.Name.Local == "\u0076\u0061\u006c" {
			_bcec.ValAttr.UnmarshalXMLAttr(_fcee)
			continue
		}
	}
	for {
		_gbag, _fafc := d.Token()
		if _fafc != nil {
			return _f.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005f\u0044\u0069r\u0065\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0073", _fafc)
		}
		if _bdcg, _gdgf := _gbag.(_a.EndElement); _gdgf && _bdcg.Name == start.Name {
			break
		}
	}
	return nil
}

func NewCT_CTCategories() *CT_CTCategories { _cbe := &CT_CTCategories{}; return _cbe }

func (_gfage ST_ArrowheadStyle) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_bgedc := _a.Attr{}
	_bgedc.Name = name
	switch _gfage {
	case ST_ArrowheadStyleUnset:
		_bgedc.Value = ""
	case ST_ArrowheadStyleAuto:
		_bgedc.Value = "\u0061\u0075\u0074\u006f"
	case ST_ArrowheadStyleArr:
		_bgedc.Value = "\u0061\u0072\u0072"
	case ST_ArrowheadStyleNoArr:
		_bgedc.Value = "\u006e\u006f\u0041r\u0072"
	}
	return _bgedc, nil
}

func (_ccbd ST_NodeVerticalAlignment) String() string {
	switch _ccbd {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u006d\u0069\u0064"
	case 3:
		return "\u0062"
	}
	return ""
}

func (_fbdf ST_ArrowheadStyle) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_fbdf.String(), start)
}

func (_gaef ST_ParameterVal) String() string {
	if _gaef.ST_DiagramHorizontalAlignment != ST_DiagramHorizontalAlignmentUnset {
		return _gaef.ST_DiagramHorizontalAlignment.String()
	}
	if _gaef.ST_VerticalAlignment != ST_VerticalAlignmentUnset {
		return _gaef.ST_VerticalAlignment.String()
	}
	if _gaef.ST_ChildDirection != ST_ChildDirectionUnset {
		return _gaef.ST_ChildDirection.String()
	}
	if _gaef.ST_ChildAlignment != ST_ChildAlignmentUnset {
		return _gaef.ST_ChildAlignment.String()
	}
	if _gaef.ST_SecondaryChildAlignment != ST_SecondaryChildAlignmentUnset {
		return _gaef.ST_SecondaryChildAlignment.String()
	}
	if _gaef.ST_LinearDirection != ST_LinearDirectionUnset {
		return _gaef.ST_LinearDirection.String()
	}
	if _gaef.ST_SecondaryLinearDirection != ST_SecondaryLinearDirectionUnset {
		return _gaef.ST_SecondaryLinearDirection.String()
	}
	if _gaef.ST_StartingElement != ST_StartingElementUnset {
		return _gaef.ST_StartingElement.String()
	}
	if _gaef.ST_BendPoint != ST_BendPointUnset {
		return _gaef.ST_BendPoint.String()
	}
	if _gaef.ST_ConnectorRouting != ST_ConnectorRoutingUnset {
		return _gaef.ST_ConnectorRouting.String()
	}
	if _gaef.ST_ArrowheadStyle != ST_ArrowheadStyleUnset {
		return _gaef.ST_ArrowheadStyle.String()
	}
	if _gaef.ST_ConnectorDimension != ST_ConnectorDimensionUnset {
		return _gaef.ST_ConnectorDimension.String()
	}
	if _gaef.ST_RotationPath != ST_RotationPathUnset {
		return _gaef.ST_RotationPath.String()
	}
	if _gaef.ST_CenterShapeMapping != ST_CenterShapeMappingUnset {
		return _gaef.ST_CenterShapeMapping.String()
	}
	if _gaef.ST_NodeHorizontalAlignment != ST_NodeHorizontalAlignmentUnset {
		return _gaef.ST_NodeHorizontalAlignment.String()
	}
	if _gaef.ST_NodeVerticalAlignment != ST_NodeVerticalAlignmentUnset {
		return _gaef.ST_NodeVerticalAlignment.String()
	}
	if _gaef.ST_FallbackDimension != ST_FallbackDimensionUnset {
		return _gaef.ST_FallbackDimension.String()
	}
	if _gaef.ST_TextDirection != ST_TextDirectionUnset {
		return _gaef.ST_TextDirection.String()
	}
	if _gaef.ST_PyramidAccentPosition != ST_PyramidAccentPositionUnset {
		return _gaef.ST_PyramidAccentPosition.String()
	}
	if _gaef.ST_PyramidAccentTextMargin != ST_PyramidAccentTextMarginUnset {
		return _gaef.ST_PyramidAccentTextMargin.String()
	}
	if _gaef.ST_TextBlockDirection != ST_TextBlockDirectionUnset {
		return _gaef.ST_TextBlockDirection.String()
	}
	if _gaef.ST_TextAnchorHorizontal != ST_TextAnchorHorizontalUnset {
		return _gaef.ST_TextAnchorHorizontal.String()
	}
	if _gaef.ST_TextAnchorVertical != ST_TextAnchorVerticalUnset {
		return _gaef.ST_TextAnchorVertical.String()
	}
	if _gaef.ST_DiagramTextAlignment != ST_DiagramTextAlignmentUnset {
		return _gaef.ST_DiagramTextAlignment.String()
	}
	if _gaef.ST_AutoTextRotation != ST_AutoTextRotationUnset {
		return _gaef.ST_AutoTextRotation.String()
	}
	if _gaef.ST_GrowDirection != ST_GrowDirectionUnset {
		return _gaef.ST_GrowDirection.String()
	}
	if _gaef.ST_FlowDirection != ST_FlowDirectionUnset {
		return _gaef.ST_FlowDirection.String()
	}
	if _gaef.ST_ContinueDirection != ST_ContinueDirectionUnset {
		return _gaef.ST_ContinueDirection.String()
	}
	if _gaef.ST_Breakpoint != ST_BreakpointUnset {
		return _gaef.ST_Breakpoint.String()
	}
	if _gaef.ST_Offset != ST_OffsetUnset {
		return _gaef.ST_Offset.String()
	}
	if _gaef.ST_HierarchyAlignment != ST_HierarchyAlignmentUnset {
		return _gaef.ST_HierarchyAlignment.String()
	}
	if _gaef.Int32 != nil {
		return _f.Sprintf("\u0025\u0076", *_gaef.Int32)
	}
	if _gaef.Float64 != nil {
		return _f.Sprintf("\u0025\u0076", *_gaef.Float64)
	}
	if _gaef.Bool != nil {
		return _f.Sprintf("\u0025\u0076", *_gaef.Bool)
	}
	if _gaef.StringVal != nil {
		return _f.Sprintf("\u0025\u0076", *_gaef.StringVal)
	}
	if _gaef.ST_ConnectorPoint != ST_ConnectorPointUnset {
		return _gaef.ST_ConnectorPoint.String()
	}
	return ""
}

func (_aadg *CT_Constraints) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _aadg.Constr != nil {
		_bdff := _a.StartElement{Name: _a.Name{Local: "\u0063\u006f\u006e\u0073\u0074\u0072"}}
		for _, _fafe := range _aadg.Constr {
			e.EncodeElement(_fafe, _bdff)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_gcgbg ST_FunctionType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_gcgbg.String(), start)
}

func (_adeg *CT_ColorTransformHeader) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _gcef := range start.Attr {
		if _gcef.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_bbee, _cag := _gcef.Value, error(nil)
			if _cag != nil {
				return _cag
			}
			_adeg.UniqueIdAttr = _bbee
			continue
		}
		if _gcef.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_gegd, _eebb := _gcef.Value, error(nil)
			if _eebb != nil {
				return _eebb
			}
			_adeg.MinVerAttr = &_gegd
			continue
		}
		if _gcef.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_cab, _acc := _ae.ParseInt(_gcef.Value, 10, 32)
			if _acc != nil {
				return _acc
			}
			_gbad := int32(_cab)
			_adeg.ResIdAttr = &_gbad
			continue
		}
	}
_adab:
	for {
		_edfe, _def := d.Token()
		if _def != nil {
			return _def
		}
		switch _acd := _edfe.(type) {
		case _a.StartElement:
			switch _acd.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_dadb := NewCT_CTName()
				if _cdbg := d.DecodeElement(_dadb, &_acd); _cdbg != nil {
					return _cdbg
				}
				_adeg.Title = append(_adeg.Title, _dadb)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_ddc := NewCT_CTDescription()
				if _geec := d.DecodeElement(_ddc, &_acd); _geec != nil {
					return _geec
				}
				_adeg.Desc = append(_adeg.Desc, _ddc)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_adeg.CatLst = NewCT_CTCategories()
				if _dggb := d.DecodeElement(_adeg.CatLst, &_acd); _dggb != nil {
					return _dggb
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_adeg.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _egag := d.DecodeElement(_adeg.ExtLst, &_acd); _egag != nil {
					return _egag
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn\u0020\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061\u006es\u0066\u006f\u0072\u006d\u0048\u0065a\u0064\u0065\u0072 \u0025\u0076", _acd.Name)
				if _eea := d.Skip(); _eea != nil {
					return _eea
				}
			}
		case _a.EndElement:
			break _adab
		case _a.CharData:
		}
	}
	return nil
}

func (_dfbaa ST_HierarchyAlignment) ValidateWithPath(path string) error {
	switch _dfbaa {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dfbaa))
	}
	return nil
}

func (_bgfb ST_AutoTextRotation) String() string {
	switch _bgfb {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0075\u0070\u0072"
	case 3:
		return "\u0067\u0072\u0061\u0076"
	}
	return ""
}

func (_egea *StyleDefHdr) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072"
	return _egea.CT_StyleDefinitionHeader.MarshalXML(e, start)
}

func (_agcea ST_TextBlockDirection) ValidateWithPath(path string) error {
	switch _agcea {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_agcea))
	}
	return nil
}

func (_aedd ST_DiagramTextAlignment) String() string {
	switch _aedd {
	case 0:
		return ""
	case 1:
		return "\u006c"
	case 2:
		return "\u0063\u0074\u0072"
	case 3:
		return "\u0072"
	}
	return ""
}

func (_bcde ST_ElementType) Validate() error { return _bcde.ValidateWithPath("") }

func (_ddae *CT_Pt) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006do\u0064\u0065\u006c\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", _ddae.ModelIdAttr)})
	if _ddae.TypeAttr != ST_PtTypeUnset {
		_acfg, _deddf := _ddae.TypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _deddf != nil {
			return _deddf
		}
		start.Attr = append(start.Attr, _acfg)
	}
	if _ddae.CxnIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u0078\u006eI\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_ddae.CxnIdAttr)})
	}
	e.EncodeToken(start)
	if _ddae.PrSet != nil {
		_dffa := _a.StartElement{Name: _a.Name{Local: "\u0070\u0072\u0053e\u0074"}}
		e.EncodeElement(_ddae.PrSet, _dffa)
	}
	if _ddae.SpPr != nil {
		_gecb := _a.StartElement{Name: _a.Name{Local: "\u0073\u0070\u0050\u0072"}}
		e.EncodeElement(_ddae.SpPr, _gecb)
	}
	if _ddae.T != nil {
		_fadf := _a.StartElement{Name: _a.Name{Local: "\u0074"}}
		e.EncodeElement(_ddae.T, _fadf)
	}
	if _ddae.ExtLst != nil {
		_bgbd := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_ddae.ExtLst, _bgbd)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_dede ST_PyramidAccentPosition) Validate() error { return _dede.ValidateWithPath("") }

func (_cbbbf *CT_SDDescription) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _ebagd := range start.Attr {
		if _ebagd.Name.Local == "\u006c\u0061\u006e\u0067" {
			_fdge, _beca := _ebagd.Value, error(nil)
			if _beca != nil {
				return _beca
			}
			_cbbbf.LangAttr = &_fdge
			continue
		}
		if _ebagd.Name.Local == "\u0076\u0061\u006c" {
			_dbdg, _aecb := _ebagd.Value, error(nil)
			if _aecb != nil {
				return _aecb
			}
			_cbbbf.ValAttr = _dbdg
			continue
		}
	}
	for {
		_adbd, _fcbg := d.Token()
		if _fcbg != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0043\u0054\u005f\u0053\u0044\u0044\u0065s\u0063r\u0069\u0070\u0074\u0069\u006f\u006e\u003a \u0025\u0073", _fcbg)
		}
		if _adde, _affe := _adbd.(_a.EndElement); _affe && _adde.Name == start.Name {
			break
		}
	}
	return nil
}

func (_fegea *ST_FunctionType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_fegea = 0
	case "\u0063\u006e\u0074":
		*_fegea = 1
	case "\u0070\u006f\u0073":
		*_fegea = 2
	case "\u0072\u0065\u0076\u0050\u006f\u0073":
		*_fegea = 3
	case "\u0070o\u0073\u0045\u0076\u0065\u006e":
		*_fegea = 4
	case "\u0070\u006f\u0073\u004f\u0064\u0064":
		*_fegea = 5
	case "\u0076\u0061\u0072":
		*_fegea = 6
	case "\u0064\u0065\u0070t\u0068":
		*_fegea = 7
	case "\u006d\u0061\u0078\u0044\u0065\u0070\u0074\u0068":
		*_fegea = 8
	}
	return nil
}

type CT_Constraint struct {
	OpAttr         ST_BoolOperator
	ValAttr        *float64
	FactAttr       *float64
	ExtLst         *_bg.CT_OfficeArtExtensionList
	TypeAttr       ST_ConstraintType
	ForAttr        ST_ConstraintRelationship
	ForNameAttr    *string
	PtTypeAttr     ST_ElementType
	RefTypeAttr    ST_ConstraintType
	RefForAttr     ST_ConstraintRelationship
	RefForNameAttr *string
	RefPtTypeAttr  ST_ElementType
}

func (_egdd *ST_ParameterVal) ValidateWithPath(path string) error {
	_cffc := []string{}
	if _egdd.ST_DiagramHorizontalAlignment != ST_DiagramHorizontalAlignmentUnset {
		_cffc = append(_cffc, "\u0053\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0048\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0041\u006c\u0069\u0067\u006eme\u006e\u0074")
	}
	if _egdd.ST_VerticalAlignment != ST_VerticalAlignmentUnset {
		_cffc = append(_cffc, "S\u0054_\u0056\u0065\u0072\u0074\u0069\u0063\u0061\u006cA\u006c\u0069\u0067\u006eme\u006e\u0074")
	}
	if _egdd.ST_ChildDirection != ST_ChildDirectionUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u0044\u0069\u0072\u0065c\u0074\u0069\u006f\u006e")
	}
	if _egdd.ST_ChildAlignment != ST_ChildAlignmentUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u0041\u006c\u0069\u0067n\u006d\u0065\u006e\u0074")
	}
	if _egdd.ST_SecondaryChildAlignment != ST_SecondaryChildAlignmentUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u0053\u0065\u0063\u006f\u006e\u0064\u0061\u0072y\u0043\u0068\u0069\u006c\u0064\u0041\u006c\u0069\u0067\u006em\u0065\u006e\u0074")
	}
	if _egdd.ST_LinearDirection != ST_LinearDirectionUnset {
		_cffc = append(_cffc, "\u0053T\u005fL\u0069\u006e\u0065\u0061\u0072D\u0069\u0072e\u0063\u0074\u0069\u006f\u006e")
	}
	if _egdd.ST_SecondaryLinearDirection != ST_SecondaryLinearDirectionUnset {
		_cffc = append(_cffc, "S\u0054\u005f\u0053\u0065\u0063\u006fn\u0064\u0061\u0072\u0079\u004c\u0069\u006e\u0065\u0061r\u0044\u0069\u0072e\u0063t\u0069\u006f\u006e")
	}
	if _egdd.ST_StartingElement != ST_StartingElementUnset {
		_cffc = append(_cffc, "\u0053T\u005fS\u0074\u0061\u0072\u0074\u0069n\u0067\u0045l\u0065\u006d\u0065\u006e\u0074")
	}
	if _egdd.ST_BendPoint != ST_BendPointUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005fB\u0065\u006e\u0064\u0050\u006f\u0069\u006e\u0074")
	}
	if _egdd.ST_ConnectorRouting != ST_ConnectorRoutingUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u0043on\u006e\u0065\u0063\u0074\u006f\u0072\u0052\u006f\u0075\u0074\u0069\u006e\u0067")
	}
	if _egdd.ST_ArrowheadStyle != ST_ArrowheadStyleUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u0041\u0072\u0072\u006f\u0077\u0068\u0065\u0061\u0064S\u0074\u0079\u006c\u0065")
	}
	if _egdd.ST_ConnectorDimension != ST_ConnectorDimensionUnset {
		_cffc = append(_cffc, "S\u0054\u005f\u0043\u006fnn\u0065c\u0074\u006f\u0072\u0044\u0069m\u0065\u006e\u0073\u0069\u006f\u006e")
	}
	if _egdd.ST_RotationPath != ST_RotationPathUnset {
		_cffc = append(_cffc, "\u0053T\u005fR\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u0050\u0061\u0074\u0068")
	}
	if _egdd.ST_CenterShapeMapping != ST_CenterShapeMappingUnset {
		_cffc = append(_cffc, "S\u0054\u005f\u0043\u0065nt\u0065r\u0053\u0068\u0061\u0070\u0065M\u0061\u0070\u0070\u0069\u006e\u0067")
	}
	if _egdd.ST_NodeHorizontalAlignment != ST_NodeHorizontalAlignmentUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u004e\u006f\u0064\u0065\u0048\u006f\u0072\u0069z\u006f\u006e\u0074\u0061\u006c\u0041\u006c\u0069\u0067\u006em\u0065\u006e\u0074")
	}
	if _egdd.ST_NodeVerticalAlignment != ST_NodeVerticalAlignmentUnset {
		_cffc = append(_cffc, "\u0053T\u005f\u004e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0069\u0063a\u006c\u0041\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074")
	}
	if _egdd.ST_FallbackDimension != ST_FallbackDimensionUnset {
		_cffc = append(_cffc, "S\u0054_\u0046\u0061\u006c\u006c\u0062\u0061\u0063\u006bD\u0069\u006d\u0065\u006esi\u006f\u006e")
	}
	if _egdd.ST_TextDirection != ST_TextDirectionUnset {
		_cffc = append(_cffc, "\u0053\u0054_\u0054\u0065\u0078t\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _egdd.ST_PyramidAccentPosition != ST_PyramidAccentPositionUnset {
		_cffc = append(_cffc, "\u0053T\u005f\u0050\u0079\u0072\u0061\u006d\u0069\u0064\u0041\u0063\u0063e\u006e\u0074\u0050\u006f\u0073\u0069\u0074\u0069\u006f\u006e")
	}
	if _egdd.ST_PyramidAccentTextMargin != ST_PyramidAccentTextMarginUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u0050\u0079\u0072\u0061\u006d\u0069\u0064\u0041c\u0063\u0065\u006e\u0074\u0054\u0065\u0078\u0074\u004d\u0061r\u0067\u0069\u006e")
	}
	if _egdd.ST_TextBlockDirection != ST_TextBlockDirectionUnset {
		_cffc = append(_cffc, "S\u0054\u005f\u0054\u0065xt\u0042l\u006f\u0063\u006b\u0044\u0069r\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _egdd.ST_TextAnchorHorizontal != ST_TextAnchorHorizontalUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005fTe\u0078\u0074\u0041\u006e\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c")
	}
	if _egdd.ST_TextAnchorVertical != ST_TextAnchorVerticalUnset {
		_cffc = append(_cffc, "S\u0054\u005f\u0054\u0065xt\u0041n\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0069\u0063\u0061\u006c")
	}
	if _egdd.ST_DiagramTextAlignment != ST_DiagramTextAlignmentUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005fDi\u0061\u0067\u0072\u0061\u006d\u0054\u0065\u0078\u0074\u0041\u006c\u0069\u0067\u006e\u006d\u0065\u006e\u0074")
	}
	if _egdd.ST_AutoTextRotation != ST_AutoTextRotationUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u0041ut\u006f\u0054\u0065\u0078\u0074\u0052\u006f\u0074\u0061\u0074\u0069\u006f\u006e")
	}
	if _egdd.ST_GrowDirection != ST_GrowDirectionUnset {
		_cffc = append(_cffc, "\u0053\u0054_\u0047\u0072\u006fw\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _egdd.ST_FlowDirection != ST_FlowDirectionUnset {
		_cffc = append(_cffc, "\u0053\u0054_\u0046\u006c\u006fw\u0044\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _egdd.ST_ContinueDirection != ST_ContinueDirectionUnset {
		_cffc = append(_cffc, "S\u0054_\u0043\u006f\u006e\u0074\u0069\u006e\u0075\u0065D\u0069\u0072\u0065\u0063ti\u006f\u006e")
	}
	if _egdd.ST_Breakpoint != ST_BreakpointUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u0042\u0072\u0065\u0061\u006b\u0070\u006f\u0069\u006e\u0074")
	}
	if _egdd.ST_Offset != ST_OffsetUnset {
		_cffc = append(_cffc, "\u0053T\u005f\u004f\u0066\u0066\u0073\u0065t")
	}
	if _egdd.ST_HierarchyAlignment != ST_HierarchyAlignmentUnset {
		_cffc = append(_cffc, "S\u0054\u005f\u0048\u0069er\u0061r\u0063\u0068\u0079\u0041\u006ci\u0067\u006e\u006d\u0065\u006e\u0074")
	}
	if _egdd.Int32 != nil {
		_cffc = append(_cffc, "\u0049\u006e\u00743\u0032")
	}
	if _egdd.Float64 != nil {
		_cffc = append(_cffc, "\u0046l\u006f\u0061\u0074\u0036\u0034")
	}
	if _egdd.Bool != nil {
		_cffc = append(_cffc, "\u0042\u006f\u006f\u006c")
	}
	if _egdd.StringVal != nil {
		_cffc = append(_cffc, "\u0053t\u0072\u0069\u006e\u0067\u0056\u0061l")
	}
	if _egdd.ST_ConnectorPoint != ST_ConnectorPointUnset {
		_cffc = append(_cffc, "\u0053\u0054\u005f\u0043\u006f\u006e\u006e\u0065\u0063\u0074\u006f\u0072P\u006f\u0069\u006e\u0074")
	}
	if len(_cffc) > 1 {
		return _f.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _cffc)
	}
	return nil
}

func (_cacg *CT_ElemPropSet) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _cacg.PresAssocIDAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "p\u0072\u0065\u0073\u0041\u0073\u0073\u006f\u0063\u0049\u0044"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.PresAssocIDAttr)})
	}
	if _cacg.PresNameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0072\u0065\u0073\u004e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.PresNameAttr)})
	}
	if _cacg.PresStyleLblAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.PresStyleLblAttr)})
	}
	if _cacg.PresStyleIdxAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u0049\u0064\u0078"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.PresStyleIdxAttr)})
	}
	if _cacg.PresStyleCntAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0072\u0065s\u0053\u0074\u0079\u006c\u0065\u0043\u006e\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.PresStyleCntAttr)})
	}
	if _cacg.LoTypeIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006c\u006f\u0054\u0079\u0070\u0065\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.LoTypeIdAttr)})
	}
	if _cacg.LoCatIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006co\u0043\u0061\u0074\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.LoCatIdAttr)})
	}
	if _cacg.QsTypeIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0071\u0073\u0054\u0079\u0070\u0065\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.QsTypeIdAttr)})
	}
	if _cacg.QsCatIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0071s\u0043\u0061\u0074\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.QsCatIdAttr)})
	}
	if _cacg.CsTypeIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u0073\u0054\u0079\u0070\u0065\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CsTypeIdAttr)})
	}
	if _cacg.CsCatIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063s\u0043\u0061\u0074\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CsCatIdAttr)})
	}
	if _cacg.Coherent3DOffAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u006f\u0068\u0065\u0072\u0065\u006e\u0074\u0033\u0044\u004f\u0066\u0066"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_cacg.Coherent3DOffAttr))})
	}
	if _cacg.PhldrTAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0068\u006c\u0064\u0072\u0054"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.PhldrTAttr)})
	}
	if _cacg.PhldrAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0068\u006cd\u0072"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_cacg.PhldrAttr))})
	}
	if _cacg.CustAngAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063u\u0073\u0074\u0041\u006e\u0067"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustAngAttr)})
	}
	if _cacg.CustFlipVertAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u0075\u0073t\u0046\u006c\u0069\u0070\u0056\u0065\u0072\u0074"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_cacg.CustFlipVertAttr))})
	}
	if _cacg.CustFlipHorAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "c\u0075\u0073\u0074\u0046\u006c\u0069\u0070\u0048\u006f\u0072"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_cacg.CustFlipHorAttr))})
	}
	if _cacg.CustSzXAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063u\u0073\u0074\u0053\u007a\u0058"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustSzXAttr)})
	}
	if _cacg.CustSzYAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063u\u0073\u0074\u0053\u007a\u0059"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustSzYAttr)})
	}
	if _cacg.CustScaleXAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u0075\u0073\u0074\u0053\u0063\u0061\u006c\u0065\u0058"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustScaleXAttr)})
	}
	if _cacg.CustScaleYAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u0075\u0073\u0074\u0053\u0063\u0061\u006c\u0065\u0059"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustScaleYAttr)})
	}
	if _cacg.CustTAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u0075\u0073t\u0054"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_cacg.CustTAttr))})
	}
	if _cacg.CustLinFactXAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u0075\u0073t\u004c\u0069\u006e\u0046\u0061\u0063\u0074\u0058"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustLinFactXAttr)})
	}
	if _cacg.CustLinFactYAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u0075\u0073t\u004c\u0069\u006e\u0046\u0061\u0063\u0074\u0059"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustLinFactYAttr)})
	}
	if _cacg.CustLinFactNeighborXAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "c\u0075s\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074N\u0065\u0069\u0067\u0068bo\u0072\u0058"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustLinFactNeighborXAttr)})
	}
	if _cacg.CustLinFactNeighborYAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "c\u0075s\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074N\u0065\u0069\u0067\u0068bo\u0072\u0059"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustLinFactNeighborYAttr)})
	}
	if _cacg.CustRadScaleRadAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063u\u0073t\u0052\u0061\u0064\u0053\u0063\u0061\u006c\u0065\u0052\u0061\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustRadScaleRadAttr)})
	}
	if _cacg.CustRadScaleIncAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063u\u0073t\u0052\u0061\u0064\u0053\u0063\u0061\u006c\u0065\u0049\u006e\u0063"}, Value: _f.Sprintf("\u0025\u0076", *_cacg.CustRadScaleIncAttr)})
	}
	e.EncodeToken(start)
	if _cacg.PresLayoutVars != nil {
		_gaa := _a.StartElement{Name: _a.Name{Local: "\u0070\u0072\u0065\u0073\u004c\u0061\u0079\u006f\u0075t\u0056\u0061\u0072\u0073"}}
		e.EncodeElement(_cacg.PresLayoutVars, _gaa)
	}
	if _cacg.Style != nil {
		_beddc := _a.StartElement{Name: _a.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_cacg.Style, _beddc)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_fedaee ST_ResizeHandlesStr) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_fedaee.String(), start)
}

func (_dgbg ST_FlowDirection) String() string {
	switch _dgbg {
	case 0:
		return ""
	case 1:
		return "\u0072\u006f\u0077"
	case 2:
		return "\u0063\u006f\u006c"
	}
	return ""
}

// Validate validates the CT_Constraint and its children
func (_fgge *CT_Constraint) Validate() error {
	return _fgge.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074")
}

func (_gfe *CT_Choose) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _gfe.NameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_gfe.NameAttr)})
	}
	e.EncodeToken(start)
	_dggg := _a.StartElement{Name: _a.Name{Local: "\u0069\u0066"}}
	for _, _bcbd := range _gfe.If {
		e.EncodeElement(_bcbd, _dggg)
	}
	if _gfe.Else != nil {
		_bfgf := _a.StartElement{Name: _a.Name{Local: "\u0065\u006c\u0073\u0065"}}
		e.EncodeElement(_gfe.Else, _bfgf)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the DataModel and its children, prefixing error messages with path
func (_fedaeb *DataModel) ValidateWithPath(path string) error {
	if _cecfec := _fedaeb.CT_DataModel.ValidateWithPath(path); _cecfec != nil {
		return _cecfec
	}
	return nil
}

// ValidateWithPath validates the CT_SampleData and its children, prefixing error messages with path
func (_dbcfa *CT_SampleData) ValidateWithPath(path string) error {
	if _dbcfa.DataModel != nil {
		if _gegg := _dbcfa.DataModel.ValidateWithPath(path + "\u002f\u0044\u0061\u0074\u0061\u004d\u006f\u0064\u0065\u006c"); _gegg != nil {
			return _gegg
		}
	}
	return nil
}

func NewCT_Name() *CT_Name { _gead := &CT_Name{}; return _gead }

func (_aagfcf ST_BendPoint) String() string {
	switch _aagfcf {
	case 0:
		return ""
	case 1:
		return "\u0062\u0065\u0067"
	case 2:
		return "\u0064\u0065\u0066"
	case 3:
		return "\u0065\u006e\u0064"
	}
	return ""
}

type CT_ForEach struct {
	NameAttr          *string
	RefAttr           *string
	Alg               []*CT_Algorithm
	Shape             []*CT_Shape
	PresOf            []*CT_PresentationOf
	ConstrLst         []*CT_Constraints
	RuleLst           []*CT_Rules
	ForEach           []*CT_ForEach
	LayoutNode        []*CT_LayoutNode
	Choose            []*CT_Choose
	ExtLst            []*_bg.CT_OfficeArtExtensionList
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}

type CT_AdjLst struct{ Adj []*CT_Adj }

const (
	ST_PtTypeUnset    ST_PtType = 0
	ST_PtTypeNode     ST_PtType = 1
	ST_PtTypeAsst     ST_PtType = 2
	ST_PtTypeDoc      ST_PtType = 3
	ST_PtTypePres     ST_PtType = 4
	ST_PtTypeParTrans ST_PtType = 5
	ST_PtTypeSibTrans ST_PtType = 6
)

func (_bcggc ST_CenterShapeMapping) String() string {
	switch _bcggc {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0066\u004e\u006fd\u0065"
	}
	return ""
}

type ST_Booleans []bool

const (
	ST_CenterShapeMappingUnset ST_CenterShapeMapping = 0
	ST_CenterShapeMappingNone  ST_CenterShapeMapping = 1
	ST_CenterShapeMappingFNode ST_CenterShapeMapping = 2
)

func (_ceee *CT_StyleDefinition) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _fgfa := range start.Attr {
		if _fgfa.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_dgddg, _cdfde := _fgfa.Value, error(nil)
			if _cdfde != nil {
				return _cdfde
			}
			_ceee.UniqueIdAttr = &_dgddg
			continue
		}
		if _fgfa.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_aecbc, _cebb := _fgfa.Value, error(nil)
			if _cebb != nil {
				return _cebb
			}
			_ceee.MinVerAttr = &_aecbc
			continue
		}
	}
_fbfff:
	for {
		_babc, _ccbe := d.Token()
		if _ccbe != nil {
			return _ccbe
		}
		switch _bcbbg := _babc.(type) {
		case _a.StartElement:
			switch _bcbbg.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_acag := NewCT_SDName()
				if _eeca := d.DecodeElement(_acag, &_bcbbg); _eeca != nil {
					return _eeca
				}
				_ceee.Title = append(_ceee.Title, _acag)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_ccde := NewCT_SDDescription()
				if _gcce := d.DecodeElement(_ccde, &_bcbbg); _gcce != nil {
					return _gcce
				}
				_ceee.Desc = append(_ceee.Desc, _ccde)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_ceee.CatLst = NewCT_SDCategories()
				if _dcec := d.DecodeElement(_ceee.CatLst, &_bcbbg); _dcec != nil {
					return _dcec
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}:
				_ceee.Scene3d = _bg.NewCT_Scene3D()
				if _bcdd := d.DecodeElement(_ceee.Scene3d, &_bcbbg); _bcdd != nil {
					return _bcdd
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}:
				_bedb := NewCT_StyleLabel()
				if _aebee := d.DecodeElement(_bedb, &_bcbbg); _aebee != nil {
					return _aebee
				}
				_ceee.StyleLbl = append(_ceee.StyleLbl, _bedb)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_ceee.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _dddgc := d.DecodeElement(_ceee.ExtLst, &_bcbbg); _dddgc != nil {
					return _dddgc
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065l\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065\u0044e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0020\u0025\u0076", _bcbbg.Name)
				if _eagf := d.Skip(); _eagf != nil {
					return _eagf
				}
			}
		case _a.EndElement:
			break _fbfff
		case _a.CharData:
		}
	}
	return nil
}

func (_ecbce ST_RotationPath) Validate() error { return _ecbce.ValidateWithPath("") }

func (_cfgd *CT_SDCategories) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_ebccb:
	for {
		_cabg, _dbcf := d.Token()
		if _dbcf != nil {
			return _dbcf
		}
		switch _debbc := _cabg.(type) {
		case _a.StartElement:
			switch _debbc.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074"}:
				_dfbf := NewCT_SDCategory()
				if _gced := d.DecodeElement(_dfbf, &_debbc); _gced != nil {
					return _gced
				}
				_cfgd.Cat = append(_cfgd.Cat, _dfbf)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0044\u0043a\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073\u0020\u0025\u0076", _debbc.Name)
				if _gfda := d.Skip(); _gfda != nil {
					return _gfda
				}
			}
		case _a.EndElement:
			break _ebccb
		case _a.CharData:
		}
	}
	return nil
}

func NewColorsDefHdrLst() *ColorsDefHdrLst {
	_cgee := &ColorsDefHdrLst{}
	_cgee.CT_ColorTransformHeaderLst = *NewCT_ColorTransformHeaderLst()
	return _cgee
}

func (_bgfab ST_Direction) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_bcaed := _a.Attr{}
	_bcaed.Name = name
	switch _bgfab {
	case ST_DirectionUnset:
		_bcaed.Value = ""
	case ST_DirectionNorm:
		_bcaed.Value = "\u006e\u006f\u0072\u006d"
	case ST_DirectionRev:
		_bcaed.Value = "\u0072\u0065\u0076"
	}
	return _bcaed, nil
}

func (_eefa ST_ConstraintType) Validate() error { return _eefa.ValidateWithPath("") }

const (
	ST_BreakpointUnset  ST_Breakpoint = 0
	ST_BreakpointEndCnv ST_Breakpoint = 1
	ST_BreakpointBal    ST_Breakpoint = 2
	ST_BreakpointFixed  ST_Breakpoint = 3
)

type CT_DataModel struct {
	PtLst  *CT_PtList
	CxnLst *CT_CxnList
	Bg     *_bg.CT_BackgroundFormatting
	Whole  *_bg.CT_WholeE2oFormatting
	ExtLst *_bg.CT_OfficeArtExtensionList
}

func (_ffcg *ST_FallbackDimension) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ffcg = 0
	case "\u0031\u0044":
		*_ffcg = 1
	case "\u0032\u0044":
		*_ffcg = 2
	}
	return nil
}

// Validate validates the CT_Colors and its children
func (_ebge *CT_Colors) Validate() error {
	return _ebge.ValidateWithPath("\u0043T\u005f\u0043\u006f\u006c\u006f\u0072s")
}

func (_eagbac ST_Breakpoint) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_eagbac.String(), start)
}

func (_ebfgac *ST_DiagramHorizontalAlignment) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ebfgac = 0
	case "\u006c":
		*_ebfgac = 1
	case "\u0063\u0074\u0072":
		*_ebfgac = 2
	case "\u0072":
		*_ebfgac = 3
	case "\u006e\u006f\u006e\u0065":
		*_ebfgac = 4
	}
	return nil
}

func (_dbdfd ST_HierBranchStyle) Validate() error { return _dbdfd.ValidateWithPath("") }

type ST_ContinueDirection byte

func (_bagbg ST_ElementType) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_bagbg.String(), start)
}

func (_dgdee *ST_ClrAppMethod) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_eaede, _cecfb := d.Token()
	if _cecfb != nil {
		return _cecfb
	}
	if _bfce, _cagd := _eaede.(_a.EndElement); _cagd && _bfce.Name == start.Name {
		*_dgdee = 1
		return nil
	}
	if _dddae, _bfea := _eaede.(_a.CharData); !_bfea {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eaede)
	} else {
		switch string(_dddae) {
		case "":
			*_dgdee = 0
		case "\u0073\u0070\u0061\u006e":
			*_dgdee = 1
		case "\u0063\u0079\u0063l\u0065":
			*_dgdee = 2
		case "\u0072\u0065\u0070\u0065\u0061\u0074":
			*_dgdee = 3
		}
	}
	_eaede, _cecfb = d.Token()
	if _cecfb != nil {
		return _cecfb
	}
	if _cfdb, _dcfga := _eaede.(_a.EndElement); _dcfga && _cfdb.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eaede)
}

func (_egcb ST_NodeVerticalAlignment) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_eebbc := _a.Attr{}
	_eebbc.Name = name
	switch _egcb {
	case ST_NodeVerticalAlignmentUnset:
		_eebbc.Value = ""
	case ST_NodeVerticalAlignmentT:
		_eebbc.Value = "\u0074"
	case ST_NodeVerticalAlignmentMid:
		_eebbc.Value = "\u006d\u0069\u0064"
	case ST_NodeVerticalAlignmentB:
		_eebbc.Value = "\u0062"
	}
	return _eebbc, nil
}

func (_gafa ST_NodeVerticalAlignment) ValidateWithPath(path string) error {
	switch _gafa {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gafa))
	}
	return nil
}

func (_gddg ST_TextAnchorVertical) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_gddg.String(), start)
}

type ST_DiagramTextAlignment byte

const (
	ST_OutputShapeTypeUnset ST_OutputShapeType = 0
	ST_OutputShapeTypeNone  ST_OutputShapeType = 1
	ST_OutputShapeTypeConn  ST_OutputShapeType = 2
)

func (_gbeaf *ST_FunctionValue) Validate() error { return _gbeaf.ValidateWithPath("") }

func (_ggfeb *CT_DiagramDefinitionHeaderLst) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _ggfeb.LayoutDefHdr != nil {
		_edda := _a.StartElement{Name: _a.Name{Local: "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072"}}
		for _, _ffee := range _ggfeb.LayoutDefHdr {
			e.EncodeElement(_ffee, _edda)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

type ST_ParameterId byte

func (_abbd ST_ConnectorRouting) Validate() error { return _abbd.ValidateWithPath("") }

func (_eaafb *ST_Breakpoint) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_fgba, _dbfdd := d.Token()
	if _dbfdd != nil {
		return _dbfdd
	}
	if _baddf, _feaeb := _fgba.(_a.EndElement); _feaeb && _baddf.Name == start.Name {
		*_eaafb = 1
		return nil
	}
	if _ebada, _dfaga := _fgba.(_a.CharData); !_dfaga {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fgba)
	} else {
		switch string(_ebada) {
		case "":
			*_eaafb = 0
		case "\u0065\u006e\u0064\u0043\u006e\u0076":
			*_eaafb = 1
		case "\u0062\u0061\u006c":
			*_eaafb = 2
		case "\u0066\u0069\u0078e\u0064":
			*_eaafb = 3
		}
	}
	_fgba, _dbfdd = d.Token()
	if _dbfdd != nil {
		return _dbfdd
	}
	if _cbcfg, _fadag := _fgba.(_a.EndElement); _fadag && _cbcfg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fgba)
}

func (_dfgfb ST_TextDirection) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_fbcad := _a.Attr{}
	_fbcad.Name = name
	switch _dfgfb {
	case ST_TextDirectionUnset:
		_fbcad.Value = ""
	case ST_TextDirectionFromT:
		_fbcad.Value = "\u0066\u0072\u006fm\u0054"
	case ST_TextDirectionFromB:
		_fbcad.Value = "\u0066\u0072\u006fm\u0042"
	}
	return _fbcad, nil
}

func NewLayoutDefHdr() *LayoutDefHdr {
	_faaea := &LayoutDefHdr{}
	_faaea.CT_DiagramDefinitionHeader = *NewCT_DiagramDefinitionHeader()
	return _faaea
}

const (
	ST_AutoTextRotationUnset ST_AutoTextRotation = 0
	ST_AutoTextRotationNone  ST_AutoTextRotation = 1
	ST_AutoTextRotationUpr   ST_AutoTextRotation = 2
	ST_AutoTextRotationGrav  ST_AutoTextRotation = 3
)

func (_eefb ST_PtType) ValidateWithPath(path string) error {
	switch _eefb {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_eefb))
	}
	return nil
}

func (_cbfb *ST_LinearDirection) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_cbfb = 0
	case "\u0066\u0072\u006fm\u004c":
		*_cbfb = 1
	case "\u0066\u0072\u006fm\u0052":
		*_cbfb = 2
	case "\u0066\u0072\u006fm\u0054":
		*_cbfb = 3
	case "\u0066\u0072\u006fm\u0042":
		*_cbfb = 4
	}
	return nil
}

func (_bcgf *ST_CxnType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_cgfa, _fdbag := d.Token()
	if _fdbag != nil {
		return _fdbag
	}
	if _efbf, _gegcd := _cgfa.(_a.EndElement); _gegcd && _efbf.Name == start.Name {
		*_bcgf = 1
		return nil
	}
	if _ceeec, _cfdae := _cgfa.(_a.CharData); !_cfdae {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cgfa)
	} else {
		switch string(_ceeec) {
		case "":
			*_bcgf = 0
		case "\u0070\u0061\u0072O\u0066":
			*_bcgf = 1
		case "\u0070\u0072\u0065\u0073\u004f\u0066":
			*_bcgf = 2
		case "\u0070r\u0065\u0073\u0050\u0061\u0072\u004ff":
			*_bcgf = 3
		case "\u0075\u006e\u006b\u006eow\u006e\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070":
			*_bcgf = 4
		}
	}
	_cgfa, _fdbag = d.Token()
	if _fdbag != nil {
		return _fdbag
	}
	if _afed, _gebd := _cgfa.(_a.EndElement); _gebd && _afed.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cgfa)
}

// Validate validates the CT_Rules and its children
func (_cebd *CT_Rules) Validate() error {
	return _cebd.ValidateWithPath("\u0043\u0054\u005f\u0052\u0075\u006c\u0065\u0073")
}

func (_adg *CT_AdjLst) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_abd:
	for {
		_ge, _adb := d.Token()
		if _adb != nil {
			return _adb
		}
		switch _caf := _ge.(type) {
		case _a.StartElement:
			switch _caf.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u0064\u006a"}:
				_egd := NewCT_Adj()
				if _aagf := d.DecodeElement(_egd, &_caf); _aagf != nil {
					return _aagf
				}
				_adg.Adj = append(_adg.Adj, _egd)
			default:
				_b.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u0054\u005f\u0041\u0064\u006aL\u0073\u0074 \u0025\u0076", _caf.Name)
				if _aff := d.Skip(); _aff != nil {
					return _aff
				}
			}
		case _a.EndElement:
			break _abd
		case _a.CharData:
		}
	}
	return nil
}

const (
	ST_ChildAlignmentUnset ST_ChildAlignment = 0
	ST_ChildAlignmentT     ST_ChildAlignment = 1
	ST_ChildAlignmentB     ST_ChildAlignment = 2
	ST_ChildAlignmentL     ST_ChildAlignment = 3
	ST_ChildAlignmentR     ST_ChildAlignment = 4
)

func (_dadaa ST_TextDirection) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_dadaa.String(), start)
}

func (_eccgc ST_ArrowheadStyle) Validate() error { return _eccgc.ValidateWithPath("") }

func NewCT_Direction() *CT_Direction { _bbbb := &CT_Direction{}; return _bbbb }

const (
	ST_HierBranchStyleUnset ST_HierBranchStyle = 0
	ST_HierBranchStyleL     ST_HierBranchStyle = 1
	ST_HierBranchStyleR     ST_HierBranchStyle = 2
	ST_HierBranchStyleHang  ST_HierBranchStyle = 3
	ST_HierBranchStyleStd   ST_HierBranchStyle = 4
	ST_HierBranchStyleInit  ST_HierBranchStyle = 5
)

// ValidateWithPath validates the AG_ConstraintAttributes and its children, prefixing error messages with path
func (_da *AG_ConstraintAttributes) ValidateWithPath(path string) error {
	if _dae := _da.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _dae != nil {
		return _dae
	}
	if _aag := _da.ForAttr.ValidateWithPath(path + "\u002f\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _aag != nil {
		return _aag
	}
	if _ab := _da.PtTypeAttr.ValidateWithPath(path + "/\u0050\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _ab != nil {
		return _ab
	}
	return nil
}

func NewCT_Choose() *CT_Choose { _daff := &CT_Choose{}; return _daff }

type ST_RotationPath byte

type CT_Cxn struct {
	ModelIdAttr    ST_ModelId
	TypeAttr       ST_CxnType
	SrcIdAttr      ST_ModelId
	DestIdAttr     ST_ModelId
	SrcOrdAttr     uint32
	DestOrdAttr    uint32
	ParTransIdAttr *ST_ModelId
	SibTransIdAttr *ST_ModelId
	PresIdAttr     *string
	ExtLst         *_bg.CT_OfficeArtExtensionList
}

type ST_ConnectorDimension byte

// Validate validates the AG_IteratorAttributes and its children
func (_gcc *AG_IteratorAttributes) Validate() error {
	return _gcc.ValidateWithPath("A\u0047\u005f\u0049\u0074er\u0061t\u006f\u0072\u0041\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u0073")
}

type ST_CenterShapeMapping byte

func NewCT_ChildMax() *CT_ChildMax { _afcb := &CT_ChildMax{}; return _afcb }

func (_cdedc ST_NodeHorizontalAlignment) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_dagfg := _a.Attr{}
	_dagfg.Name = name
	switch _cdedc {
	case ST_NodeHorizontalAlignmentUnset:
		_dagfg.Value = ""
	case ST_NodeHorizontalAlignmentL:
		_dagfg.Value = "\u006c"
	case ST_NodeHorizontalAlignmentCtr:
		_dagfg.Value = "\u0063\u0074\u0072"
	case ST_NodeHorizontalAlignmentR:
		_dagfg.Value = "\u0072"
	}
	return _dagfg, nil
}

func (_edfead ST_HierBranchStyle) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_fgcff := _a.Attr{}
	_fgcff.Name = name
	switch _edfead {
	case ST_HierBranchStyleUnset:
		_fgcff.Value = ""
	case ST_HierBranchStyleL:
		_fgcff.Value = "\u006c"
	case ST_HierBranchStyleR:
		_fgcff.Value = "\u0072"
	case ST_HierBranchStyleHang:
		_fgcff.Value = "\u0068\u0061\u006e\u0067"
	case ST_HierBranchStyleStd:
		_fgcff.Value = "\u0073\u0074\u0064"
	case ST_HierBranchStyleInit:
		_fgcff.Value = "\u0069\u006e\u0069\u0074"
	}
	return _fgcff, nil
}

type CT_StyleLabel struct {
	NameAttr string
	Scene3d  *_bg.CT_Scene3D
	Sp3d     *_bg.CT_Shape3D
	TxPr     *CT_TextProps
	Style    *_bg.CT_ShapeStyle
	ExtLst   *_bg.CT_OfficeArtExtensionList
}

func (_gddaa *CT_StyleDefinitionHeader) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _abeee := range start.Attr {
		if _abeee.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_ccefa, _adfcg := _abeee.Value, error(nil)
			if _adfcg != nil {
				return _adfcg
			}
			_gddaa.UniqueIdAttr = _ccefa
			continue
		}
		if _abeee.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_dagc, _eeaf := _abeee.Value, error(nil)
			if _eeaf != nil {
				return _eeaf
			}
			_gddaa.MinVerAttr = &_dagc
			continue
		}
		if _abeee.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_cbab, _bdga := _ae.ParseInt(_abeee.Value, 10, 32)
			if _bdga != nil {
				return _bdga
			}
			_aeac := int32(_cbab)
			_gddaa.ResIdAttr = &_aeac
			continue
		}
	}
_ggdga:
	for {
		_adgaf, _aefcf := d.Token()
		if _aefcf != nil {
			return _aefcf
		}
		switch _ecgfg := _adgaf.(type) {
		case _a.StartElement:
			switch _ecgfg.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_gbge := NewCT_SDName()
				if _ebfcc := d.DecodeElement(_gbge, &_ecgfg); _ebfcc != nil {
					return _ebfcc
				}
				_gddaa.Title = append(_gddaa.Title, _gbge)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_gacdc := NewCT_SDDescription()
				if _dffg := d.DecodeElement(_gacdc, &_ecgfg); _dffg != nil {
					return _dffg
				}
				_gddaa.Desc = append(_gddaa.Desc, _gacdc)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_gddaa.CatLst = NewCT_SDCategories()
				if _efcg := d.DecodeElement(_gddaa.CatLst, &_ecgfg); _efcg != nil {
					return _efcg
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_gddaa.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _ecge := d.DecodeElement(_gddaa.ExtLst, &_ecgfg); _ecge != nil {
					return _ecge
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048e\u0061\u0064\u0065\u0072\u0020%\u0076", _ecgfg.Name)
				if _fgca := d.Skip(); _fgca != nil {
					return _fgca
				}
			}
		case _a.EndElement:
			break _ggdga
		case _a.CharData:
		}
	}
	return nil
}

const (
	ST_GrowDirectionUnset ST_GrowDirection = 0
	ST_GrowDirectionTL    ST_GrowDirection = 1
	ST_GrowDirectionTR    ST_GrowDirection = 2
	ST_GrowDirectionBL    ST_GrowDirection = 3
	ST_GrowDirectionBR    ST_GrowDirection = 4
)

const (
	ST_VerticalAlignmentUnset ST_VerticalAlignment = 0
	ST_VerticalAlignmentT     ST_VerticalAlignment = 1
	ST_VerticalAlignmentMid   ST_VerticalAlignment = 2
	ST_VerticalAlignmentB     ST_VerticalAlignment = 3
	ST_VerticalAlignmentNone  ST_VerticalAlignment = 4
)

func (_cgbd ST_AnimLvlStr) String() string {
	switch _cgbd {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u006c\u0076\u006c"
	case 3:
		return "\u0063\u0074\u0072"
	}
	return ""
}

// ValidateWithPath validates the CT_Categories and its children, prefixing error messages with path
func (_afb *CT_Categories) ValidateWithPath(path string) error {
	for _gabe, _gfc := range _afb.Cat {
		if _gggf := _gfc.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0061\u0074\u005b\u0025\u0064\u005d", path, _gabe)); _gggf != nil {
			return _gggf
		}
	}
	return nil
}

func ParseUnionST_FunctionArgument(s string) (ST_FunctionArgument, error) {
	return ST_FunctionArgument{}, nil
}

func (_ddbbc *ST_ParameterId) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ccgff, _adabe := d.Token()
	if _adabe != nil {
		return _adabe
	}
	if _fcedb, _eefc := _ccgff.(_a.EndElement); _eefc && _fcedb.Name == start.Name {
		*_ddbbc = 1
		return nil
	}
	if _dbcc, _bdfag := _ccgff.(_a.CharData); !_bdfag {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ccgff)
	} else {
		switch string(_dbcc) {
		case "":
			*_ddbbc = 0
		case "\u0068o\u0072\u007a\u0041\u006c\u0069\u0067n":
			*_ddbbc = 1
		case "\u0076e\u0072\u0074\u0041\u006c\u0069\u0067n":
			*_ddbbc = 2
		case "\u0063\u0068\u0044i\u0072":
			*_ddbbc = 3
		case "\u0063h\u0041\u006c\u0069\u0067\u006e":
			*_ddbbc = 4
		case "\u0073\u0065\u0063\u0043\u0068\u0041\u006c\u0069\u0067\u006e":
			*_ddbbc = 5
		case "\u006c\u0069\u006e\u0044\u0069\u0072":
			*_ddbbc = 6
		case "\u0073e\u0063\u004c\u0069\u006e\u0044\u0069r":
			*_ddbbc = 7
		case "\u0073\u0074\u0045\u006c\u0065\u006d":
			*_ddbbc = 8
		case "\u0062\u0065\u006e\u0064\u0050\u0074":
			*_ddbbc = 9
		case "\u0063\u006f\u006e\u006e\u0052\u006f\u0075\u0074":
			*_ddbbc = 10
		case "\u0062\u0065\u0067\u0053\u0074\u0079":
			*_ddbbc = 11
		case "\u0065\u006e\u0064\u0053\u0074\u0079":
			*_ddbbc = 12
		case "\u0064\u0069\u006d":
			*_ddbbc = 13
		case "\u0072o\u0074\u0050\u0061\u0074\u0068":
			*_ddbbc = 14
		case "\u0063t\u0072\u0053\u0068\u0070\u004d\u0061p":
			*_ddbbc = 15
		case "\u006e\u006f\u0064\u0065\u0048\u006f\u0072\u007a\u0041\u006c\u0069\u0067\u006e":
			*_ddbbc = 16
		case "\u006e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0041\u006c\u0069\u0067\u006e":
			*_ddbbc = 17
		case "\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b":
			*_ddbbc = 18
		case "\u0074\u0078\u0044i\u0072":
			*_ddbbc = 19
		case "p\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0050\u006f\u0073":
			*_ddbbc = 20
		case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054\u0078\u004d\u0061\u0072":
			*_ddbbc = 21
		case "\u0074x\u0042\u006c\u0044\u0069\u0072":
			*_ddbbc = 22
		case "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u007a":
			*_ddbbc = 23
		case "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0056\u0065\u0072\u0074":
			*_ddbbc = 24
		case "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0048o\u0072\u007a\u0043\u0068":
			*_ddbbc = 25
		case "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0043\u0068":
			*_ddbbc = 26
		case "\u0070\u0061\u0072\u0054\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e":
			*_ddbbc = 27
		case "\u0070\u0061\u0072\u0054\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e":
			*_ddbbc = 28
		case "\u0073h\u0070T\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e\u0043\u0068":
			*_ddbbc = 29
		case "\u0073h\u0070T\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e\u0043\u0068":
			*_ddbbc = 30
		case "\u0061u\u0074\u006f\u0054\u0078\u0052\u006ft":
			*_ddbbc = 31
		case "\u0067\u0072\u0044i\u0072":
			*_ddbbc = 32
		case "\u0066l\u006f\u0077\u0044\u0069\u0072":
			*_ddbbc = 33
		case "\u0063o\u006e\u0074\u0044\u0069\u0072":
			*_ddbbc = 34
		case "\u0062\u006b\u0070\u0074":
			*_ddbbc = 35
		case "\u006f\u0066\u0066":
			*_ddbbc = 36
		case "\u0068i\u0065\u0072\u0041\u006c\u0069\u0067n":
			*_ddbbc = 37
		case "\u0062\u006b\u0050t\u0046\u0069\u0078\u0065\u0064\u0056\u0061\u006c":
			*_ddbbc = 38
		case "s\u0074\u0042\u0075\u006c\u006c\u0065\u0074\u004c\u0076\u006c":
			*_ddbbc = 39
		case "\u0073\u0074\u0041n\u0067":
			*_ddbbc = 40
		case "\u0073p\u0061\u006e\u0041\u006e\u0067":
			*_ddbbc = 41
		case "\u0061\u0072":
			*_ddbbc = 42
		case "\u006cn\u0053\u0070\u0050\u0061\u0072":
			*_ddbbc = 43
		case "\u006c\u006e\u0053\u0070\u0041\u0066\u0050\u0061\u0072\u0050":
			*_ddbbc = 44
		case "\u006c\u006e\u0053\u0070\u0043\u0068":
			*_ddbbc = 45
		case "\u006cn\u0053\u0070\u0041\u0066\u0043\u0068P":
			*_ddbbc = 46
		case "r\u0074\u0053\u0068\u006f\u0072\u0074\u0044\u0069\u0073\u0074":
			*_ddbbc = 47
		case "\u0061l\u0069\u0067\u006e\u0054\u0078":
			*_ddbbc = 48
		case "p\u0079\u0072\u0061\u004c\u0076\u006c\u004e\u006f\u0064\u0065":
			*_ddbbc = 49
		case "\u0070\u0079r\u0061\u0041\u0063c\u0074\u0042\u006b\u0067\u0064\u004e\u006f\u0064\u0065":
			*_ddbbc = 50
		case "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054x\u004e\u006f\u0064\u0065":
			*_ddbbc = 51
		case "\u0073r\u0063\u004e\u006f\u0064\u0065":
			*_ddbbc = 52
		case "\u0064s\u0074\u004e\u006f\u0064\u0065":
			*_ddbbc = 53
		case "\u0062\u0065\u0067\u0050\u0074\u0073":
			*_ddbbc = 54
		case "\u0065\u006e\u0064\u0050\u0074\u0073":
			*_ddbbc = 55
		}
	}
	_ccgff, _adabe = d.Token()
	if _adabe != nil {
		return _adabe
	}
	if _aced, _cbade := _ccgff.(_a.EndElement); _cbade && _aced.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ccgff)
}

func (_cbcgb ST_SecondaryChildAlignment) String() string {
	switch _cbcgb {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0074"
	case 3:
		return "\u0062"
	case 4:
		return "\u006c"
	case 5:
		return "\u0072"
	}
	return ""
}

func (_aec *CT_LayoutVariablePropertySet) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _aec.OrgChart != nil {
		_fgbe := _a.StartElement{Name: _a.Name{Local: "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074"}}
		e.EncodeElement(_aec.OrgChart, _fgbe)
	}
	if _aec.ChMax != nil {
		_eebc := _a.StartElement{Name: _a.Name{Local: "\u0063\u0068\u004da\u0078"}}
		e.EncodeElement(_aec.ChMax, _eebc)
	}
	if _aec.ChPref != nil {
		_ecgf := _a.StartElement{Name: _a.Name{Local: "\u0063\u0068\u0050\u0072\u0065\u0066"}}
		e.EncodeElement(_aec.ChPref, _ecgf)
	}
	if _aec.BulletEnabled != nil {
		_debc := _a.StartElement{Name: _a.Name{Local: "\u0062\u0075\u006c\u006c\u0065\u0074\u0045\u006e\u0061\u0062\u006c\u0065\u0064"}}
		e.EncodeElement(_aec.BulletEnabled, _debc)
	}
	if _aec.Dir != nil {
		_ebec := _a.StartElement{Name: _a.Name{Local: "\u0064\u0069\u0072"}}
		e.EncodeElement(_aec.Dir, _ebec)
	}
	if _aec.HierBranch != nil {
		_eddf := _a.StartElement{Name: _a.Name{Local: "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"}}
		e.EncodeElement(_aec.HierBranch, _eddf)
	}
	if _aec.AnimOne != nil {
		_adcef := _a.StartElement{Name: _a.Name{Local: "\u0061n\u0069\u006d\u004f\u006e\u0065"}}
		e.EncodeElement(_aec.AnimOne, _adcef)
	}
	if _aec.AnimLvl != nil {
		_badcd := _a.StartElement{Name: _a.Name{Local: "\u0061n\u0069\u006d\u004c\u0076\u006c"}}
		e.EncodeElement(_aec.AnimLvl, _badcd)
	}
	if _aec.ResizeHandles != nil {
		_gefg := _a.StartElement{Name: _a.Name{Local: "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073"}}
		e.EncodeElement(_aec.ResizeHandles, _gefg)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_ggge *ST_BoolOperator) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_effd, _bace := d.Token()
	if _bace != nil {
		return _bace
	}
	if _ccgf, _gage := _effd.(_a.EndElement); _gage && _ccgf.Name == start.Name {
		*_ggge = 1
		return nil
	}
	if _cfga, _fcge := _effd.(_a.CharData); !_fcge {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _effd)
	} else {
		switch string(_cfga) {
		case "":
			*_ggge = 0
		case "\u006e\u006f\u006e\u0065":
			*_ggge = 1
		case "\u0065\u0071\u0075":
			*_ggge = 2
		case "\u0067\u0074\u0065":
			*_ggge = 3
		case "\u006c\u0074\u0065":
			*_ggge = 4
		}
	}
	_effd, _bace = d.Token()
	if _bace != nil {
		return _bace
	}
	if _ffce, _efae := _effd.(_a.EndElement); _efae && _ffce.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _effd)
}

func (_bgec *CT_Category) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _dfea := range start.Attr {
		if _dfea.Name.Local == "\u0074\u0079\u0070\u0065" {
			_bcf, _agcd := _dfea.Value, error(nil)
			if _agcd != nil {
				return _agcd
			}
			_bgec.TypeAttr = _bcf
			continue
		}
		if _dfea.Name.Local == "\u0070\u0072\u0069" {
			_edac, _ebdec := _ae.ParseUint(_dfea.Value, 10, 32)
			if _ebdec != nil {
				return _ebdec
			}
			_bgec.PriAttr = uint32(_edac)
			continue
		}
	}
	for {
		_bbc, _cddb := d.Token()
		if _cddb != nil {
			return _f.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079\u003a\u0020\u0025\u0073", _cddb)
		}
		if _eba, _bfd := _bbc.(_a.EndElement); _bfd && _eba.Name == start.Name {
			break
		}
	}
	return nil
}

type ST_TextBlockDirection byte

type CT_CxnList struct{ Cxn []*CT_Cxn }

func ParseUnionST_PrSetCustVal(s string) (ST_PrSetCustVal, error) { return ST_PrSetCustVal{}, nil }

func NewCT_Shape() *CT_Shape { _befec := &CT_Shape{}; return _befec }

type CT_StyleDefinitionHeaderLst struct{ StyleDefHdr []*CT_StyleDefinitionHeader }

type CT_RelIds struct {
	DmAttr string
	LoAttr string
	QsAttr string
	CsAttr string
}

func (_cdg *CT_ForEach) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _cdg.NameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_cdg.NameAttr)})
	}
	if _cdg.RefAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u0065\u0066"}, Value: _f.Sprintf("\u0025\u0076", *_cdg.RefAttr)})
	}
	if _cdg.AxisAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0061\u0078\u0069\u0073"}, Value: _f.Sprintf("\u0025\u0076", *_cdg.AxisAttr)})
	}
	if _cdg.PtTypeAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_cdg.PtTypeAttr)})
	}
	if _cdg.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073"}, Value: _f.Sprintf("\u0025\u0076", *_cdg.HideLastTransAttr)})
	}
	if _cdg.StAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_cdg.StAttr)})
	}
	if _cdg.CntAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u006e\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_cdg.CntAttr)})
	}
	if _cdg.StepAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0074\u0065\u0070"}, Value: _f.Sprintf("\u0025\u0076", *_cdg.StepAttr)})
	}
	e.EncodeToken(start)
	if _cdg.Alg != nil {
		_ffg := _a.StartElement{Name: _a.Name{Local: "\u0061\u006c\u0067"}}
		for _, _agg := range _cdg.Alg {
			e.EncodeElement(_agg, _ffg)
		}
	}
	if _cdg.Shape != nil {
		_ddee := _a.StartElement{Name: _a.Name{Local: "\u0073\u0068\u0061p\u0065"}}
		for _, _fcdf := range _cdg.Shape {
			e.EncodeElement(_fcdf, _ddee)
		}
	}
	if _cdg.PresOf != nil {
		_efff := _a.StartElement{Name: _a.Name{Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}}
		for _, _aeabg := range _cdg.PresOf {
			e.EncodeElement(_aeabg, _efff)
		}
	}
	if _cdg.ConstrLst != nil {
		_cfceb := _a.StartElement{Name: _a.Name{Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}}
		for _, _ace := range _cdg.ConstrLst {
			e.EncodeElement(_ace, _cfceb)
		}
	}
	if _cdg.RuleLst != nil {
		_gbdeb := _a.StartElement{Name: _a.Name{Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}}
		for _, _fedde := range _cdg.RuleLst {
			e.EncodeElement(_fedde, _gbdeb)
		}
	}
	if _cdg.ForEach != nil {
		_efgg := _a.StartElement{Name: _a.Name{Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}}
		for _, _gddb := range _cdg.ForEach {
			e.EncodeElement(_gddb, _efgg)
		}
	}
	if _cdg.LayoutNode != nil {
		_ggea := _a.StartElement{Name: _a.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
		for _, _bdge := range _cdg.LayoutNode {
			e.EncodeElement(_bdge, _ggea)
		}
	}
	if _cdg.Choose != nil {
		_baaga := _a.StartElement{Name: _a.Name{Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}}
		for _, _ffac := range _cdg.Choose {
			e.EncodeElement(_ffac, _baaga)
		}
	}
	if _cdg.ExtLst != nil {
		_fccf := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		for _, _fgga := range _cdg.ExtLst {
			e.EncodeElement(_fgga, _fccf)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_gdefe ST_PrSetCustVal) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _gdefe.ST_Percentage != nil {
		e.EncodeToken(_a.CharData(*_gdefe.ST_Percentage))
	}
	if _gdefe.Int32 != nil {
		e.EncodeToken(_a.CharData(_f.Sprintf("\u0025\u0064", *_gdefe.Int32)))
	}
	return e.EncodeToken(_a.EndElement{Name: start.Name})
}

func (_ccac *CT_DiagramDefinitionHeaderLst) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_ecda:
	for {
		_dedf, _fbab := d.Token()
		if _fbab != nil {
			return _fbab
		}
		switch _cddc := _dedf.(type) {
		case _a.StartElement:
			switch _cddc.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_bebb := NewCT_DiagramDefinitionHeader()
				if _eaeb := d.DecodeElement(_bebb, &_cddc); _eaeb != nil {
					return _eaeb
				}
				_ccac.LayoutDefHdr = append(_ccac.LayoutDefHdr, _bebb)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074 \u006f\u006e\u0020\u0043\u0054\u005f\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006ei\u0074\u0069\u006f\u006e\u0048\u0065a\u0064\u0065r\u004c\u0073t\u0020%\u0076", _cddc.Name)
				if _cdbbc := d.Skip(); _cdbbc != nil {
					return _cdbbc
				}
			}
		case _a.EndElement:
			break _ecda
		case _a.CharData:
		}
	}
	return nil
}

func (_fdcgg ST_RotationPath) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_fdcgg.String(), start)
}

type CT_When struct {
	NameAttr          *string
	FuncAttr          ST_FunctionType
	ArgAttr           *ST_FunctionArgument
	OpAttr            ST_FunctionOperator
	ValAttr           ST_FunctionValue
	Alg               []*CT_Algorithm
	Shape             []*CT_Shape
	PresOf            []*CT_PresentationOf
	ConstrLst         []*CT_Constraints
	RuleLst           []*CT_Rules
	ForEach           []*CT_ForEach
	LayoutNode        []*CT_LayoutNode
	Choose            []*CT_Choose
	ExtLst            []*_bg.CT_OfficeArtExtensionList
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}

func (_gdabd *ST_Direction) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_gdabd = 0
	case "\u006e\u006f\u0072\u006d":
		*_gdabd = 1
	case "\u0072\u0065\u0076":
		*_gdabd = 2
	}
	return nil
}

const (
	ST_FlowDirectionUnset ST_FlowDirection = 0
	ST_FlowDirectionRow   ST_FlowDirection = 1
	ST_FlowDirectionCol   ST_FlowDirection = 2
)

type ST_ConnectorRouting byte

// Validate validates the CT_DataModel and its children
func (_edded *CT_DataModel) Validate() error {
	return _edded.ValidateWithPath("\u0043\u0054\u005fD\u0061\u0074\u0061\u004d\u006f\u0064\u0065\u006c")
}

func (_daccd ST_DiagramTextAlignment) ValidateWithPath(path string) error {
	switch _daccd {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_daccd))
	}
	return nil
}

// ValidateWithPath validates the CT_ColorTransformHeaderLst and its children, prefixing error messages with path
func (_ffag *CT_ColorTransformHeaderLst) ValidateWithPath(path string) error {
	for _egdc, _fgea := range _ffag.ColorsDefHdr {
		if _efe := _fgea.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043ol\u006f\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072\u005b\u0025\u0064\u005d", path, _egdc)); _efe != nil {
			return _efe
		}
	}
	return nil
}

type CT_AnimOne struct{ ValAttr ST_AnimOneStr }

// ValidateWithPath validates the ColorsDefHdr and its children, prefixing error messages with path
func (_bfdcb *ColorsDefHdr) ValidateWithPath(path string) error {
	if _aebd := _bfdcb.CT_ColorTransformHeader.ValidateWithPath(path); _aebd != nil {
		return _aebd
	}
	return nil
}

type ST_TextDirection byte

func (_cacc ST_ChildAlignment) String() string {
	switch _cacc {
	case 0:
		return ""
	case 1:
		return "\u0074"
	case 2:
		return "\u0062"
	case 3:
		return "\u006c"
	case 4:
		return "\u0072"
	}
	return ""
}

func (_ddaca ST_FallbackDimension) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_gaca := _a.Attr{}
	_gaca.Name = name
	switch _ddaca {
	case ST_FallbackDimensionUnset:
		_gaca.Value = ""
	case ST_FallbackDimension1D:
		_gaca.Value = "\u0031\u0044"
	case ST_FallbackDimension2D:
		_gaca.Value = "\u0032\u0044"
	}
	return _gaca, nil
}

func (_feaadg ST_DiagramHorizontalAlignment) String() string {
	switch _feaadg {
	case 0:
		return ""
	case 1:
		return "\u006c"
	case 2:
		return "\u0063\u0074\u0072"
	case 3:
		return "\u0072"
	case 4:
		return "\u006e\u006f\u006e\u0065"
	}
	return ""
}

func (_dfgf *CT_SDName) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _dfgf.LangAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _f.Sprintf("\u0025\u0076", *_dfgf.LangAttr)})
	}
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", _dfgf.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_gdfcg ST_HueDir) String() string {
	switch _gdfcg {
	case 0:
		return ""
	case 1:
		return "\u0063\u0077"
	case 2:
		return "\u0063\u0063\u0077"
	}
	return ""
}

func (_fedga ST_ConnectorDimension) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_fedga.String(), start)
}

func (_fegcad ST_FlowDirection) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_fegcad.String(), start)
}

const (
	ST_TextAnchorHorizontalUnset ST_TextAnchorHorizontal = 0
	ST_TextAnchorHorizontalNone  ST_TextAnchorHorizontal = 1
	ST_TextAnchorHorizontalCtr   ST_TextAnchorHorizontal = 2
)

// ValidateWithPath validates the CT_SDDescription and its children, prefixing error messages with path
func (_gbfcd *CT_SDDescription) ValidateWithPath(path string) error { return nil }

// Validate validates the CT_ChildMax and its children
func (_abed *CT_ChildMax) Validate() error {
	return _abed.ValidateWithPath("C\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u004d\u0061\u0078")
}

const (
	ST_ParameterIdUnset            ST_ParameterId = 0
	ST_ParameterIdHorzAlign        ST_ParameterId = 1
	ST_ParameterIdVertAlign        ST_ParameterId = 2
	ST_ParameterIdChDir            ST_ParameterId = 3
	ST_ParameterIdChAlign          ST_ParameterId = 4
	ST_ParameterIdSecChAlign       ST_ParameterId = 5
	ST_ParameterIdLinDir           ST_ParameterId = 6
	ST_ParameterIdSecLinDir        ST_ParameterId = 7
	ST_ParameterIdStElem           ST_ParameterId = 8
	ST_ParameterIdBendPt           ST_ParameterId = 9
	ST_ParameterIdConnRout         ST_ParameterId = 10
	ST_ParameterIdBegSty           ST_ParameterId = 11
	ST_ParameterIdEndSty           ST_ParameterId = 12
	ST_ParameterIdDim              ST_ParameterId = 13
	ST_ParameterIdRotPath          ST_ParameterId = 14
	ST_ParameterIdCtrShpMap        ST_ParameterId = 15
	ST_ParameterIdNodeHorzAlign    ST_ParameterId = 16
	ST_ParameterIdNodeVertAlign    ST_ParameterId = 17
	ST_ParameterIdFallback         ST_ParameterId = 18
	ST_ParameterIdTxDir            ST_ParameterId = 19
	ST_ParameterIdPyraAcctPos      ST_ParameterId = 20
	ST_ParameterIdPyraAcctTxMar    ST_ParameterId = 21
	ST_ParameterIdTxBlDir          ST_ParameterId = 22
	ST_ParameterIdTxAnchorHorz     ST_ParameterId = 23
	ST_ParameterIdTxAnchorVert     ST_ParameterId = 24
	ST_ParameterIdTxAnchorHorzCh   ST_ParameterId = 25
	ST_ParameterIdTxAnchorVertCh   ST_ParameterId = 26
	ST_ParameterIdParTxLTRAlign    ST_ParameterId = 27
	ST_ParameterIdParTxRTLAlign    ST_ParameterId = 28
	ST_ParameterIdShpTxLTRAlignCh  ST_ParameterId = 29
	ST_ParameterIdShpTxRTLAlignCh  ST_ParameterId = 30
	ST_ParameterIdAutoTxRot        ST_ParameterId = 31
	ST_ParameterIdGrDir            ST_ParameterId = 32
	ST_ParameterIdFlowDir          ST_ParameterId = 33
	ST_ParameterIdContDir          ST_ParameterId = 34
	ST_ParameterIdBkpt             ST_ParameterId = 35
	ST_ParameterIdOff              ST_ParameterId = 36
	ST_ParameterIdHierAlign        ST_ParameterId = 37
	ST_ParameterIdBkPtFixedVal     ST_ParameterId = 38
	ST_ParameterIdStBulletLvl      ST_ParameterId = 39
	ST_ParameterIdStAng            ST_ParameterId = 40
	ST_ParameterIdSpanAng          ST_ParameterId = 41
	ST_ParameterIdAr               ST_ParameterId = 42
	ST_ParameterIdLnSpPar          ST_ParameterId = 43
	ST_ParameterIdLnSpAfParP       ST_ParameterId = 44
	ST_ParameterIdLnSpCh           ST_ParameterId = 45
	ST_ParameterIdLnSpAfChP        ST_ParameterId = 46
	ST_ParameterIdRtShortDist      ST_ParameterId = 47
	ST_ParameterIdAlignTx          ST_ParameterId = 48
	ST_ParameterIdPyraLvlNode      ST_ParameterId = 49
	ST_ParameterIdPyraAcctBkgdNode ST_ParameterId = 50
	ST_ParameterIdPyraAcctTxNode   ST_ParameterId = 51
	ST_ParameterIdSrcNode          ST_ParameterId = 52
	ST_ParameterIdDstNode          ST_ParameterId = 53
	ST_ParameterIdBegPts           ST_ParameterId = 54
	ST_ParameterIdEndPts           ST_ParameterId = 55
)

func (_gacd *CT_ResizeHandles) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _gacd.ValAttr != ST_ResizeHandlesStrUnset {
		_cfee, _dagde := _gacd.ValAttr.MarshalXMLAttr(_a.Name{Local: "\u0076\u0061\u006c"})
		if _dagde != nil {
			return _dagde
		}
		start.Attr = append(start.Attr, _cfee)
	}
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_dfba *LayoutDefHdrLst) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_dfba.CT_DiagramDefinitionHeaderLst = *NewCT_DiagramDefinitionHeaderLst()
_gbgb:
	for {
		_cadgc, _cccec := d.Token()
		if _cccec != nil {
			return _cccec
		}
		switch _ddedd := _cadgc.(type) {
		case _a.StartElement:
			switch _ddedd.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_bceee := NewCT_DiagramDefinitionHeader()
				if _fdabd := d.DecodeElement(_bceee, &_ddedd); _fdabd != nil {
					return _fdabd
				}
				_dfba.LayoutDefHdr = append(_dfba.LayoutDefHdr, _bceee)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u004c\u0061\u0079\u006f\u0075\u0074D\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074\u0020\u0025\u0076", _ddedd.Name)
				if _bafe := d.Skip(); _bafe != nil {
					return _bafe
				}
			}
		case _a.EndElement:
			break _gbgb
		case _a.CharData:
		}
	}
	return nil
}

const (
	ST_ConnectorRoutingUnset     ST_ConnectorRouting = 0
	ST_ConnectorRoutingStra      ST_ConnectorRouting = 1
	ST_ConnectorRoutingBend      ST_ConnectorRouting = 2
	ST_ConnectorRoutingCurve     ST_ConnectorRouting = 3
	ST_ConnectorRoutingLongCurve ST_ConnectorRouting = 4
)

func NewCT_CTCategory() *CT_CTCategory { _aga := &CT_CTCategory{}; return _aga }

func (_gcee ST_SecondaryChildAlignment) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_bbfc := _a.Attr{}
	_bbfc.Name = name
	switch _gcee {
	case ST_SecondaryChildAlignmentUnset:
		_bbfc.Value = ""
	case ST_SecondaryChildAlignmentNone:
		_bbfc.Value = "\u006e\u006f\u006e\u0065"
	case ST_SecondaryChildAlignmentT:
		_bbfc.Value = "\u0074"
	case ST_SecondaryChildAlignmentB:
		_bbfc.Value = "\u0062"
	case ST_SecondaryChildAlignmentL:
		_bbfc.Value = "\u006c"
	case ST_SecondaryChildAlignmentR:
		_bbfc.Value = "\u0072"
	}
	return _bbfc, nil
}

func (_fdc *CT_Direction) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _fdc.ValAttr != ST_DirectionUnset {
		_fdg, _cee := _fdc.ValAttr.MarshalXMLAttr(_a.Name{Local: "\u0076\u0061\u006c"})
		if _cee != nil {
			return _cee
		}
		start.Attr = append(start.Attr, _fdg)
	}
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_abf *CT_Cxn) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006do\u0064\u0065\u006c\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", _abf.ModelIdAttr)})
	if _abf.TypeAttr != ST_CxnTypeUnset {
		_eecg, _fba := _abf.TypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0074\u0079\u0070\u0065"})
		if _fba != nil {
			return _fba
		}
		start.Attr = append(start.Attr, _eecg)
	}
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0072\u0063I\u0064"}, Value: _f.Sprintf("\u0025\u0076", _abf.SrcIdAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0064\u0065\u0073\u0074\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", _abf.DestIdAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0072\u0063\u004f\u0072\u0064"}, Value: _f.Sprintf("\u0025\u0076", _abf.SrcOrdAttr)})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0064e\u0073\u0074\u004f\u0072\u0064"}, Value: _f.Sprintf("\u0025\u0076", _abf.DestOrdAttr)})
	if _abf.ParTransIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_abf.ParTransIdAttr)})
	}
	if _abf.SibTransIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_abf.SibTransIdAttr)})
	}
	if _abf.PresIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0072\u0065\u0073\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_abf.PresIdAttr)})
	}
	e.EncodeToken(start)
	if _abf.ExtLst != nil {
		_cad := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_abf.ExtLst, _cad)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_bcge *ST_BendPoint) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_defg, _bafg := d.Token()
	if _bafg != nil {
		return _bafg
	}
	if _fgeaee, _cadcb := _defg.(_a.EndElement); _cadcb && _fgeaee.Name == start.Name {
		*_bcge = 1
		return nil
	}
	if _cecgb, _baggg := _defg.(_a.CharData); !_baggg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _defg)
	} else {
		switch string(_cecgb) {
		case "":
			*_bcge = 0
		case "\u0062\u0065\u0067":
			*_bcge = 1
		case "\u0064\u0065\u0066":
			*_bcge = 2
		case "\u0065\u006e\u0064":
			*_bcge = 3
		}
	}
	_defg, _bafg = d.Token()
	if _bafg != nil {
		return _bafg
	}
	if _cfebca, _fdagc := _defg.(_a.EndElement); _fdagc && _cfebca.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _defg)
}

// ValidateWithPath validates the CT_Colors and its children, prefixing error messages with path
func (_eacbb *CT_Colors) ValidateWithPath(path string) error {
	if _dfbg := _eacbb.MethAttr.ValidateWithPath(path + "\u002fM\u0065\u0074\u0068\u0041\u0074\u0074r"); _dfbg != nil {
		return _dfbg
	}
	if _eeced := _eacbb.HueDirAttr.ValidateWithPath(path + "/\u0048\u0075\u0065\u0044\u0069\u0072\u0041\u0074\u0074\u0072"); _eeced != nil {
		return _eeced
	}
	for _dcbcg, _aae := range _eacbb.EG_ColorChoice {
		if _bbef := _aae.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0045\u0047_C\u006fl\u006f\u0072\u0043\u0068\u006fi\u0063\u0065\u005b\u0025\u0064\u005d", path, _dcbcg)); _bbef != nil {
			return _bbef
		}
	}
	return nil
}

// Validate validates the CT_TextProps and its children
func (_ccfef *CT_TextProps) Validate() error {
	return _ccfef.ValidateWithPath("\u0043\u0054\u005fT\u0065\u0078\u0074\u0050\u0072\u006f\u0070\u0073")
}

func (_aead ST_ChildDirection) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_aead.String(), start)
}

func NewCT_ForEach() *CT_ForEach { _gbfe := &CT_ForEach{}; return _gbfe }

func (_fedgb *ST_ConnectorRouting) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_fedgb = 0
	case "\u0073\u0074\u0072\u0061":
		*_fedgb = 1
	case "\u0062\u0065\u006e\u0064":
		*_fedgb = 2
	case "\u0063\u0075\u0072v\u0065":
		*_fedgb = 3
	case "\u006co\u006e\u0067\u0043\u0075\u0072\u0076e":
		*_fedgb = 4
	}
	return nil
}

// ValidateWithPath validates the CT_CTDescription and its children, prefixing error messages with path
func (_bcdg *CT_CTDescription) ValidateWithPath(path string) error { return nil }

func (_bddbf ST_FunctionType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_dfge := _a.Attr{}
	_dfge.Name = name
	switch _bddbf {
	case ST_FunctionTypeUnset:
		_dfge.Value = ""
	case ST_FunctionTypeCnt:
		_dfge.Value = "\u0063\u006e\u0074"
	case ST_FunctionTypePos:
		_dfge.Value = "\u0070\u006f\u0073"
	case ST_FunctionTypeRevPos:
		_dfge.Value = "\u0072\u0065\u0076\u0050\u006f\u0073"
	case ST_FunctionTypePosEven:
		_dfge.Value = "\u0070o\u0073\u0045\u0076\u0065\u006e"
	case ST_FunctionTypePosOdd:
		_dfge.Value = "\u0070\u006f\u0073\u004f\u0064\u0064"
	case ST_FunctionTypeVar:
		_dfge.Value = "\u0076\u0061\u0072"
	case ST_FunctionTypeDepth:
		_dfge.Value = "\u0064\u0065\u0070t\u0068"
	case ST_FunctionTypeMaxDepth:
		_dfge.Value = "\u006d\u0061\u0078\u0044\u0065\u0070\u0074\u0068"
	}
	return _dfge, nil
}

func (_egdb *ST_LinearDirection) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_bfefgf, _adfeb := d.Token()
	if _adfeb != nil {
		return _adfeb
	}
	if _ddfbg, _aagae := _bfefgf.(_a.EndElement); _aagae && _ddfbg.Name == start.Name {
		*_egdb = 1
		return nil
	}
	if _ddgdf, _fdbf := _bfefgf.(_a.CharData); !_fdbf {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _bfefgf)
	} else {
		switch string(_ddgdf) {
		case "":
			*_egdb = 0
		case "\u0066\u0072\u006fm\u004c":
			*_egdb = 1
		case "\u0066\u0072\u006fm\u0052":
			*_egdb = 2
		case "\u0066\u0072\u006fm\u0054":
			*_egdb = 3
		case "\u0066\u0072\u006fm\u0042":
			*_egdb = 4
		}
	}
	_bfefgf, _adfeb = d.Token()
	if _adfeb != nil {
		return _adfeb
	}
	if _dbeba, _agbe := _bfefgf.(_a.EndElement); _agbe && _dbeba.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _bfefgf)
}

const (
	ST_VariableTypeUnset         ST_VariableType = 0
	ST_VariableTypeNone          ST_VariableType = 1
	ST_VariableTypeOrgChart      ST_VariableType = 2
	ST_VariableTypeChMax         ST_VariableType = 3
	ST_VariableTypeChPref        ST_VariableType = 4
	ST_VariableTypeBulEnabled    ST_VariableType = 5
	ST_VariableTypeDir           ST_VariableType = 6
	ST_VariableTypeHierBranch    ST_VariableType = 7
	ST_VariableTypeAnimOne       ST_VariableType = 8
	ST_VariableTypeAnimLvl       ST_VariableType = 9
	ST_VariableTypeResizeHandles ST_VariableType = 10
)

func (_cgbab *ST_NodeHorizontalAlignment) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_baad, _eddea := d.Token()
	if _eddea != nil {
		return _eddea
	}
	if _gecgb, _fdgd := _baad.(_a.EndElement); _fdgd && _gecgb.Name == start.Name {
		*_cgbab = 1
		return nil
	}
	if _bcag, _ddcfd := _baad.(_a.CharData); !_ddcfd {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _baad)
	} else {
		switch string(_bcag) {
		case "":
			*_cgbab = 0
		case "\u006c":
			*_cgbab = 1
		case "\u0063\u0074\u0072":
			*_cgbab = 2
		case "\u0072":
			*_cgbab = 3
		}
	}
	_baad, _eddea = d.Token()
	if _eddea != nil {
		return _eddea
	}
	if _gedda, _eabad := _baad.(_a.EndElement); _eabad && _gedda.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _baad)
}

type ST_ElementType byte

func (_bddb *CT_LayoutNode) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _bddb.NameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_bddb.NameAttr)})
	}
	if _bddb.StyleLblAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}, Value: _f.Sprintf("\u0025\u0076", *_bddb.StyleLblAttr)})
	}
	if _bddb.ChOrderAttr != ST_ChildOrderTypeUnset {
		_affcb, _fdaaf := _bddb.ChOrderAttr.MarshalXMLAttr(_a.Name{Local: "\u0063h\u004f\u0072\u0064\u0065\u0072"})
		if _fdaaf != nil {
			return _fdaaf
		}
		start.Attr = append(start.Attr, _affcb)
	}
	if _bddb.MoveWithAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006d\u006f\u0076\u0065\u0057\u0069\u0074\u0068"}, Value: _f.Sprintf("\u0025\u0076", *_bddb.MoveWithAttr)})
	}
	e.EncodeToken(start)
	if _bddb.Alg != nil {
		_fcgd := _a.StartElement{Name: _a.Name{Local: "\u0061\u006c\u0067"}}
		for _, _dce := range _bddb.Alg {
			e.EncodeElement(_dce, _fcgd)
		}
	}
	if _bddb.Shape != nil {
		_cdabf := _a.StartElement{Name: _a.Name{Local: "\u0073\u0068\u0061p\u0065"}}
		for _, _agfdc := range _bddb.Shape {
			e.EncodeElement(_agfdc, _cdabf)
		}
	}
	if _bddb.PresOf != nil {
		_gfdgb := _a.StartElement{Name: _a.Name{Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}}
		for _, _fdad := range _bddb.PresOf {
			e.EncodeElement(_fdad, _gfdgb)
		}
	}
	if _bddb.ConstrLst != nil {
		_ecbc := _a.StartElement{Name: _a.Name{Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}}
		for _, _egdf := range _bddb.ConstrLst {
			e.EncodeElement(_egdf, _ecbc)
		}
	}
	if _bddb.RuleLst != nil {
		_dgeb := _a.StartElement{Name: _a.Name{Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}}
		for _, _eade := range _bddb.RuleLst {
			e.EncodeElement(_eade, _dgeb)
		}
	}
	if _bddb.VarLst != nil {
		_bgge := _a.StartElement{Name: _a.Name{Local: "\u0076\u0061\u0072\u004c\u0073\u0074"}}
		for _, _bgca := range _bddb.VarLst {
			e.EncodeElement(_bgca, _bgge)
		}
	}
	if _bddb.ForEach != nil {
		_cebg := _a.StartElement{Name: _a.Name{Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}}
		for _, _bbac := range _bddb.ForEach {
			e.EncodeElement(_bbac, _cebg)
		}
	}
	if _bddb.LayoutNode != nil {
		_acfc := _a.StartElement{Name: _a.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
		for _, _deab := range _bddb.LayoutNode {
			e.EncodeElement(_deab, _acfc)
		}
	}
	if _bddb.Choose != nil {
		_aaed := _a.StartElement{Name: _a.Name{Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}}
		for _, _gdac := range _bddb.Choose {
			e.EncodeElement(_gdac, _aaed)
		}
	}
	if _bddb.ExtLst != nil {
		_becb := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		for _, _egfg := range _bddb.ExtLst {
			e.EncodeElement(_egfg, _becb)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

type CT_ElemPropSet struct {
	PresAssocIDAttr          *ST_ModelId
	PresNameAttr             *string
	PresStyleLblAttr         *string
	PresStyleIdxAttr         *int32
	PresStyleCntAttr         *int32
	LoTypeIdAttr             *string
	LoCatIdAttr              *string
	QsTypeIdAttr             *string
	QsCatIdAttr              *string
	CsTypeIdAttr             *string
	CsCatIdAttr              *string
	Coherent3DOffAttr        *bool
	PhldrTAttr               *string
	PhldrAttr                *bool
	CustAngAttr              *int32
	CustFlipVertAttr         *bool
	CustFlipHorAttr          *bool
	CustSzXAttr              *int32
	CustSzYAttr              *int32
	CustScaleXAttr           *ST_PrSetCustVal
	CustScaleYAttr           *ST_PrSetCustVal
	CustTAttr                *bool
	CustLinFactXAttr         *ST_PrSetCustVal
	CustLinFactYAttr         *ST_PrSetCustVal
	CustLinFactNeighborXAttr *ST_PrSetCustVal
	CustLinFactNeighborYAttr *ST_PrSetCustVal
	CustRadScaleRadAttr      *ST_PrSetCustVal
	CustRadScaleIncAttr      *ST_PrSetCustVal
	PresLayoutVars           *CT_LayoutVariablePropertySet
	Style                    *_bg.CT_ShapeStyle
}

const (
	ST_BendPointUnset ST_BendPoint = 0
	ST_BendPointBeg   ST_BendPoint = 1
	ST_BendPointDef   ST_BendPoint = 2
	ST_BendPointEnd   ST_BendPoint = 3
)

func (_gfefb ST_CenterShapeMapping) ValidateWithPath(path string) error {
	switch _gfefb {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_gfefb))
	}
	return nil
}

func (_gadfe *CT_OrgChart) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _befd := range start.Attr {
		if _befd.Name.Local == "\u0076\u0061\u006c" {
			_ecga, _agef := _ae.ParseBool(_befd.Value)
			if _agef != nil {
				return _agef
			}
			_gadfe.ValAttr = &_ecga
			continue
		}
	}
	for {
		_dged, _cdbc := d.Token()
		if _cdbc != nil {
			return _f.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0072\u0067\u0043\u0068\u0061\u0072\u0074\u003a\u0020\u0025\u0073", _cdbc)
		}
		if _ecdd, _abaa := _dged.(_a.EndElement); _abaa && _ecdd.Name == start.Name {
			break
		}
	}
	return nil
}

func (_badag ST_LinearDirection) Validate() error { return _badag.ValidateWithPath("") }

func (_cged ST_TextBlockDirection) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_cged.String(), start)
}

const (
	ST_TextAnchorVerticalUnset ST_TextAnchorVertical = 0
	ST_TextAnchorVerticalT     ST_TextAnchorVertical = 1
	ST_TextAnchorVerticalMid   ST_TextAnchorVertical = 2
	ST_TextAnchorVerticalB     ST_TextAnchorVertical = 3
)

func (_dacf ST_GrowDirection) String() string {
	switch _dacf {
	case 0:
		return ""
	case 1:
		return "\u0074\u004c"
	case 2:
		return "\u0074\u0052"
	case 3:
		return "\u0062\u004c"
	case 4:
		return "\u0062\u0052"
	}
	return ""
}

func (_bafd *CT_SampleData) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _abga := range start.Attr {
		if _abga.Name.Local == "\u0075\u0073\u0065\u0044\u0065\u0066" {
			_deff, _daeb := _ae.ParseBool(_abga.Value)
			if _daeb != nil {
				return _daeb
			}
			_bafd.UseDefAttr = &_deff
			continue
		}
	}
_agbd:
	for {
		_cdef, _fdab := d.Token()
		if _fdab != nil {
			return _fdab
		}
		switch _dcdb := _cdef.(type) {
		case _a.StartElement:
			switch _dcdb.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064a\u0074\u0061\u004d\u006f\u0064\u0065l"}:
				_bafd.DataModel = NewCT_DataModel()
				if _gffc := d.DecodeElement(_bafd.DataModel, &_dcdb); _gffc != nil {
					return _gffc
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053a\u006d\u0070\u006c\u0065\u0044\u0061\u0074\u0061 \u0025\u0076", _dcdb.Name)
				if _ccfe := d.Skip(); _ccfe != nil {
					return _ccfe
				}
			}
		case _a.EndElement:
			break _agbd
		case _a.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Cxn and its children
func (_dafc *CT_Cxn) Validate() error {
	return _dafc.ValidateWithPath("\u0043\u0054\u005f\u0043\u0078\u006e")
}

const (
	ST_SecondaryChildAlignmentUnset ST_SecondaryChildAlignment = 0
	ST_SecondaryChildAlignmentNone  ST_SecondaryChildAlignment = 1
	ST_SecondaryChildAlignmentT     ST_SecondaryChildAlignment = 2
	ST_SecondaryChildAlignmentB     ST_SecondaryChildAlignment = 3
	ST_SecondaryChildAlignmentL     ST_SecondaryChildAlignment = 4
	ST_SecondaryChildAlignmentR     ST_SecondaryChildAlignment = 5
)

func (_egbe *CT_PtList) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _egbe.Pt != nil {
		_bbeed := _a.StartElement{Name: _a.Name{Local: "\u0070\u0074"}}
		for _, _gdef := range _egbe.Pt {
			e.EncodeElement(_gdef, _bbeed)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_ccdff *LayoutDefHdr) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ccdff.CT_DiagramDefinitionHeader = *NewCT_DiagramDefinitionHeader()
	for _, _cdcd := range start.Attr {
		if _cdcd.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_eedab, _gabda := _cdcd.Value, error(nil)
			if _gabda != nil {
				return _gabda
			}
			_ccdff.UniqueIdAttr = _eedab
			continue
		}
		if _cdcd.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_dbdc, _baba := _cdcd.Value, error(nil)
			if _baba != nil {
				return _baba
			}
			_ccdff.MinVerAttr = &_dbdc
			continue
		}
		if _cdcd.Name.Local == "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065" {
			_dcbfc, _acdfbc := _cdcd.Value, error(nil)
			if _acdfbc != nil {
				return _acdfbc
			}
			_ccdff.DefStyleAttr = &_dcbfc
			continue
		}
		if _cdcd.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_fcagd, _gafeb := _ae.ParseInt(_cdcd.Value, 10, 32)
			if _gafeb != nil {
				return _gafeb
			}
			_dccf := int32(_fcagd)
			_ccdff.ResIdAttr = &_dccf
			continue
		}
	}
_bcadg:
	for {
		_gbcfd, _adaab := d.Token()
		if _adaab != nil {
			return _adaab
		}
		switch _feab := _gbcfd.(type) {
		case _a.StartElement:
			switch _feab.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_ccbea := NewCT_Name()
				if _adcdc := d.DecodeElement(_ccbea, &_feab); _adcdc != nil {
					return _adcdc
				}
				_ccdff.Title = append(_ccdff.Title, _ccbea)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_eege := NewCT_Description()
				if _fdagg := d.DecodeElement(_eege, &_feab); _fdagg != nil {
					return _fdagg
				}
				_ccdff.Desc = append(_ccdff.Desc, _eege)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_ccdff.CatLst = NewCT_Categories()
				if _gcbeb := d.DecodeElement(_ccdff.CatLst, &_feab); _gcbeb != nil {
					return _gcbeb
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_ccdff.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _gded := d.DecodeElement(_ccdff.ExtLst, &_feab); _gded != nil {
					return _gded
				}
			default:
				_b.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u004c\u0061yo\u0075\u0074D\u0065\u0066\u0048\u0064\u0072\u0020\u0025\u0076", _feab.Name)
				if _gadd := d.Skip(); _gadd != nil {
					return _gadd
				}
			}
		case _a.EndElement:
			break _bcadg
		case _a.CharData:
		}
	}
	return nil
}

func (_ccbec ST_TextAnchorVertical) Validate() error { return _ccbec.ValidateWithPath("") }

type AG_IteratorAttributes struct {
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}

func (_ebga ST_BendPoint) Validate() error { return _ebga.ValidateWithPath("") }

func (_cbdec ST_ConstraintType) ValidateWithPath(path string) error {
	switch _cbdec {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_cbdec))
	}
	return nil
}

func (_dadac ST_ConnectorPoint) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_eebab := _a.Attr{}
	_eebab.Name = name
	switch _dadac {
	case ST_ConnectorPointUnset:
		_eebab.Value = ""
	case ST_ConnectorPointAuto:
		_eebab.Value = "\u0061\u0075\u0074\u006f"
	case ST_ConnectorPointBCtr:
		_eebab.Value = "\u0062\u0043\u0074\u0072"
	case ST_ConnectorPointCtr:
		_eebab.Value = "\u0063\u0074\u0072"
	case ST_ConnectorPointMidL:
		_eebab.Value = "\u006d\u0069\u0064\u004c"
	case ST_ConnectorPointMidR:
		_eebab.Value = "\u006d\u0069\u0064\u0052"
	case ST_ConnectorPointTCtr:
		_eebab.Value = "\u0074\u0043\u0074\u0072"
	case ST_ConnectorPointBL:
		_eebab.Value = "\u0062\u004c"
	case ST_ConnectorPointBR:
		_eebab.Value = "\u0062\u0052"
	case ST_ConnectorPointTL:
		_eebab.Value = "\u0074\u004c"
	case ST_ConnectorPointTR:
		_eebab.Value = "\u0074\u0052"
	case ST_ConnectorPointRadial:
		_eebab.Value = "\u0072\u0061\u0064\u0069\u0061\u006c"
	}
	return _eebab, nil
}

func (_cgcea ST_Offset) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_dbbd := _a.Attr{}
	_dbbd.Name = name
	switch _cgcea {
	case ST_OffsetUnset:
		_dbbd.Value = ""
	case ST_OffsetCtr:
		_dbbd.Value = "\u0063\u0074\u0072"
	case ST_OffsetOff:
		_dbbd.Value = "\u006f\u0066\u0066"
	}
	return _dbbd, nil
}

func (_cafg *CT_CxnList) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _cafg.Cxn != nil {
		_egg := _a.StartElement{Name: _a.Name{Local: "\u0063\u0078\u006e"}}
		for _, _ccef := range _cafg.Cxn {
			e.EncodeElement(_ccef, _egg)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_cbbc *CT_Colors) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _cbbc.MethAttr != ST_ClrAppMethodUnset {
		_eeba, _bbeb := _cbbc.MethAttr.MarshalXMLAttr(_a.Name{Local: "\u006d\u0065\u0074\u0068"})
		if _bbeb != nil {
			return _bbeb
		}
		start.Attr = append(start.Attr, _eeba)
	}
	if _cbbc.HueDirAttr != ST_HueDirUnset {
		_efg, _cgfe := _cbbc.HueDirAttr.MarshalXMLAttr(_a.Name{Local: "\u0068\u0075\u0065\u0044\u0069\u0072"})
		if _cgfe != nil {
			return _cgfe
		}
		start.Attr = append(start.Attr, _efg)
	}
	e.EncodeToken(start)
	if _cbbc.EG_ColorChoice != nil {
		for _, _acb := range _cbbc.EG_ColorChoice {
			_acb.MarshalXML(e, _a.StartElement{})
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_aabg *ST_PyramidAccentTextMargin) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_aabg = 0
	case "\u0073\u0074\u0065\u0070":
		*_aabg = 1
	case "\u0073\u0074\u0061c\u006b":
		*_aabg = 2
	}
	return nil
}

func (_gga *AG_IteratorAttributes) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _gga.AxisAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0061\u0078\u0069\u0073"}, Value: _f.Sprintf("\u0025\u0076", *_gga.AxisAttr)})
	}
	if _gga.PtTypeAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_gga.PtTypeAttr)})
	}
	if _gga.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073"}, Value: _f.Sprintf("\u0025\u0076", *_gga.HideLastTransAttr)})
	}
	if _gga.StAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_gga.StAttr)})
	}
	if _gga.CntAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u006e\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_gga.CntAttr)})
	}
	if _gga.StepAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0074\u0065\u0070"}, Value: _f.Sprintf("\u0025\u0076", *_gga.StepAttr)})
	}
	return nil
}

func (_fddg ST_ElementType) String() string {
	switch _fddg {
	case 0:
		return ""
	case 1:
		return "\u0061\u006c\u006c"
	case 2:
		return "\u0064\u006f\u0063"
	case 3:
		return "\u006e\u006f\u0064\u0065"
	case 4:
		return "\u006e\u006f\u0072\u006d"
	case 5:
		return "\u006eo\u006e\u004e\u006f\u0072\u006d"
	case 6:
		return "\u0061\u0073\u0073\u0074"
	case 7:
		return "\u006eo\u006e\u0041\u0073\u0073\u0074"
	case 8:
		return "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073"
	case 9:
		return "\u0070\u0072\u0065\u0073"
	case 10:
		return "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073"
	}
	return ""
}

func (_ecfcf ST_Breakpoint) String() string {
	switch _ecfcf {
	case 0:
		return ""
	case 1:
		return "\u0065\u006e\u0064\u0043\u006e\u0076"
	case 2:
		return "\u0062\u0061\u006c"
	case 3:
		return "\u0066\u0069\u0078e\u0064"
	}
	return ""
}

const (
	ST_NodeVerticalAlignmentUnset ST_NodeVerticalAlignment = 0
	ST_NodeVerticalAlignmentT     ST_NodeVerticalAlignment = 1
	ST_NodeVerticalAlignmentMid   ST_NodeVerticalAlignment = 2
	ST_NodeVerticalAlignmentB     ST_NodeVerticalAlignment = 3
)

func (_fdfcg *CT_SDCategories) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _fdfcg.Cat != nil {
		_fegc := _a.StartElement{Name: _a.Name{Local: "\u0063\u0061\u0074"}}
		for _, _eedf := range _fdfcg.Cat {
			e.EncodeElement(_eedf, _fegc)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_dgecf *CT_Name) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _accg := range start.Attr {
		if _accg.Name.Local == "\u006c\u0061\u006e\u0067" {
			_bgeb, _defd := _accg.Value, error(nil)
			if _defd != nil {
				return _defd
			}
			_dgecf.LangAttr = &_bgeb
			continue
		}
		if _accg.Name.Local == "\u0076\u0061\u006c" {
			_dcaf, _agaa := _accg.Value, error(nil)
			if _agaa != nil {
				return _agaa
			}
			_dgecf.ValAttr = _dcaf
			continue
		}
	}
	for {
		_dabaa, _egc := d.Token()
		if _egc != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004e\u0061\u006d\u0065\u003a\u0020\u0025\u0073", _egc)
		}
		if _dbda, _cfdd := _dabaa.(_a.EndElement); _cfdd && _dbda.Name == start.Name {
			break
		}
	}
	return nil
}

// ValidateWithPath validates the CT_Category and its children, prefixing error messages with path
func (_ged *CT_Category) ValidateWithPath(path string) error { return nil }

func (_babf ST_HueDir) ValidateWithPath(path string) error {
	switch _babf {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_babf))
	}
	return nil
}

func (_bcgdc *ST_ChildOrderType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_bcgdc = 0
	case "\u0062":
		*_bcgdc = 1
	case "\u0074":
		*_bcgdc = 2
	}
	return nil
}

func (_bdac ST_HierarchyAlignment) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_bdac.String(), start)
}

type ST_BendPoint byte

func ParseSliceST_UnsignedInts(s string) (ST_UnsignedInts, error) { return ST_UnsignedInts{}, nil }

type ST_CxnType byte

func (_cfaf *ST_AxisType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_cfaf = 0
	case "\u0073\u0065\u006c\u0066":
		*_cfaf = 1
	case "\u0063\u0068":
		*_cfaf = 2
	case "\u0064\u0065\u0073":
		*_cfaf = 3
	case "\u0064e\u0073\u004f\u0072\u0053\u0065\u006cf":
		*_cfaf = 4
	case "\u0070\u0061\u0072":
		*_cfaf = 5
	case "\u0061\u006e\u0063s\u0074":
		*_cfaf = 6
	case "a\u006e\u0063\u0073\u0074\u004f\u0072\u0053\u0065\u006c\u0066":
		*_cfaf = 7
	case "\u0066o\u006c\u006c\u006f\u0077\u0053\u0069b":
		*_cfaf = 8
	case "\u0070r\u0065\u0063\u0065\u0064\u0053\u0069b":
		*_cfaf = 9
	case "\u0066\u006f\u006c\u006c\u006f\u0077":
		*_cfaf = 10
	case "\u0070\u0072\u0065\u0063\u0065\u0064":
		*_cfaf = 11
	case "\u0072\u006f\u006f\u0074":
		*_cfaf = 12
	case "\u006e\u006f\u006e\u0065":
		*_cfaf = 13
	}
	return nil
}

func (_gffd ST_FunctionArgument) String() string {
	if _gffd.ST_VariableType != ST_VariableTypeUnset {
		return _gffd.ST_VariableType.String()
	}
	return ""
}

func NewCT_CxnList() *CT_CxnList { _bfab := &CT_CxnList{}; return _bfab }

func (_ggab *ST_FunctionValue) ValidateWithPath(path string) error {
	_eddbb := []string{}
	if _ggab.Int32 != nil {
		_eddbb = append(_eddbb, "\u0049\u006e\u00743\u0032")
	}
	if _ggab.Bool != nil {
		_eddbb = append(_eddbb, "\u0042\u006f\u006f\u006c")
	}
	if _ggab.ST_Direction != ST_DirectionUnset {
		_eddbb = append(_eddbb, "\u0053\u0054\u005fD\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e")
	}
	if _ggab.ST_HierBranchStyle != ST_HierBranchStyleUnset {
		_eddbb = append(_eddbb, "\u0053T\u005fH\u0069\u0065\u0072\u0042\u0072a\u006e\u0063h\u0053\u0074\u0079\u006c\u0065")
	}
	if _ggab.ST_AnimOneStr != ST_AnimOneStrUnset {
		_eddbb = append(_eddbb, "\u0053\u0054\u005f\u0041\u006e\u0069\u006d\u004f\u006e\u0065\u0053\u0074\u0072")
	}
	if _ggab.ST_AnimLvlStr != ST_AnimLvlStrUnset {
		_eddbb = append(_eddbb, "\u0053\u0054\u005f\u0041\u006e\u0069\u006d\u004c\u0076\u006c\u0053\u0074\u0072")
	}
	if _ggab.ST_ResizeHandlesStr != ST_ResizeHandlesStrUnset {
		_eddbb = append(_eddbb, "\u0053\u0054\u005f\u0052es\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073\u0053\u0074\u0072")
	}
	if len(_eddbb) > 1 {
		return _f.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _eddbb)
	}
	return nil
}

func (_fbbbc ST_TextAnchorHorizontal) String() string {
	switch _fbbbc {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0063\u0074\u0072"
	}
	return ""
}

func (_gccd ST_PyramidAccentPosition) String() string {
	switch _gccd {
	case 0:
		return ""
	case 1:
		return "\u0062\u0065\u0066"
	case 2:
		return "\u0061\u0066\u0074"
	}
	return ""
}

func (_ccba ST_BoolOperator) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_dcgga := _a.Attr{}
	_dcgga.Name = name
	switch _ccba {
	case ST_BoolOperatorUnset:
		_dcgga.Value = ""
	case ST_BoolOperatorNone:
		_dcgga.Value = "\u006e\u006f\u006e\u0065"
	case ST_BoolOperatorEqu:
		_dcgga.Value = "\u0065\u0071\u0075"
	case ST_BoolOperatorGte:
		_dcgga.Value = "\u0067\u0074\u0065"
	case ST_BoolOperatorLte:
		_dcgga.Value = "\u006c\u0074\u0065"
	}
	return _dcgga, nil
}

func (_bbbeb *ST_Breakpoint) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_bbbeb = 0
	case "\u0065\u006e\u0064\u0043\u006e\u0076":
		*_bbbeb = 1
	case "\u0062\u0061\u006c":
		*_bbbeb = 2
	case "\u0066\u0069\u0078e\u0064":
		*_bbbeb = 3
	}
	return nil
}

func (_efad *ST_HueDir) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_afad, _eaef := d.Token()
	if _eaef != nil {
		return _eaef
	}
	if _bagb, _ceec := _afad.(_a.EndElement); _ceec && _bagb.Name == start.Name {
		*_efad = 1
		return nil
	}
	if _eabab, _aagff := _afad.(_a.CharData); !_aagff {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _afad)
	} else {
		switch string(_eabab) {
		case "":
			*_efad = 0
		case "\u0063\u0077":
			*_efad = 1
		case "\u0063\u0063\u0077":
			*_efad = 2
		}
	}
	_afad, _eaef = d.Token()
	if _eaef != nil {
		return _eaef
	}
	if _gaddd, _agcc := _afad.(_a.EndElement); _agcc && _gaddd.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _afad)
}

// Validate validates the CT_AnimOne and its children
func (_fce *CT_AnimOne) Validate() error {
	return _fce.ValidateWithPath("\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004f\u006e\u0065")
}

const (
	ST_RotationPathUnset     ST_RotationPath = 0
	ST_RotationPathNone      ST_RotationPath = 1
	ST_RotationPathAlongPath ST_RotationPath = 2
)

type ST_PyramidAccentTextMargin byte

// Validate validates the CT_OrgChart and its children
func (_daecb *CT_OrgChart) Validate() error {
	return _daecb.ValidateWithPath("C\u0054\u005f\u004f\u0072\u0067\u0043\u0068\u0061\u0072\u0074")
}

// ValidateWithPath validates the CT_DiagramDefinition and its children, prefixing error messages with path
func (_fcgb *CT_DiagramDefinition) ValidateWithPath(path string) error {
	for _eccb, _fcff := range _fcgb.Title {
		if _aadc := _fcff.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _eccb)); _aadc != nil {
			return _aadc
		}
	}
	for _abdg, _cadf := range _fcgb.Desc {
		if _dgdc := _cadf.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _abdg)); _dgdc != nil {
			return _dgdc
		}
	}
	if _fcgb.CatLst != nil {
		if _bfaa := _fcgb.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _bfaa != nil {
			return _bfaa
		}
	}
	if _fcgb.SampData != nil {
		if _bfefg := _fcgb.SampData.ValidateWithPath(path + "\u002fS\u0061\u006d\u0070\u0044\u0061\u0074a"); _bfefg != nil {
			return _bfefg
		}
	}
	if _fcgb.StyleData != nil {
		if _ebb := _fcgb.StyleData.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065\u0044\u0061\u0074\u0061"); _ebb != nil {
			return _ebb
		}
	}
	if _fcgb.ClrData != nil {
		if _dee := _fcgb.ClrData.ValidateWithPath(path + "\u002f\u0043\u006c\u0072\u0044\u0061\u0074\u0061"); _dee != nil {
			return _dee
		}
	}
	if _fgd := _fcgb.LayoutNode.ValidateWithPath(path + "/\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"); _fgd != nil {
		return _fgd
	}
	if _fcgb.ExtLst != nil {
		if _cgeg := _fcgb.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _cgeg != nil {
			return _cgeg
		}
	}
	return nil
}

func NewCT_Algorithm() *CT_Algorithm {
	_gdc := &CT_Algorithm{}
	_gdc.TypeAttr = ST_AlgorithmType(1)
	return _gdc
}

type CT_Categories struct{ Cat []*CT_Category }

func (_acfgg *ST_GrowDirection) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_cacdd, _bddc := d.Token()
	if _bddc != nil {
		return _bddc
	}
	if _eeea, _gcfcd := _cacdd.(_a.EndElement); _gcfcd && _eeea.Name == start.Name {
		*_acfgg = 1
		return nil
	}
	if _edaf, _ebac := _cacdd.(_a.CharData); !_ebac {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _cacdd)
	} else {
		switch string(_edaf) {
		case "":
			*_acfgg = 0
		case "\u0074\u004c":
			*_acfgg = 1
		case "\u0074\u0052":
			*_acfgg = 2
		case "\u0062\u004c":
			*_acfgg = 3
		case "\u0062\u0052":
			*_acfgg = 4
		}
	}
	_cacdd, _bddc = d.Token()
	if _bddc != nil {
		return _bddc
	}
	if _bbbe, _ffggaa := _cacdd.(_a.EndElement); _ffggaa && _bbbe.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _cacdd)
}

type CT_PresentationOf struct {
	ExtLst            *_bg.CT_OfficeArtExtensionList
	AxisAttr          *ST_AxisTypes
	PtTypeAttr        *ST_ElementTypes
	HideLastTransAttr *ST_Booleans
	StAttr            *ST_Ints
	CntAttr           *ST_UnsignedInts
	StepAttr          *ST_Ints
}

// Validate validates the CT_Categories and its children
func (_age *CT_Categories) Validate() error {
	return _age.ValidateWithPath("\u0043\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073")
}

type CT_SDCategories struct{ Cat []*CT_SDCategory }

// Validate validates the CT_Otherwise and its children
func (_gged *CT_Otherwise) Validate() error {
	return _gged.ValidateWithPath("\u0043\u0054\u005fO\u0074\u0068\u0065\u0072\u0077\u0069\u0073\u0065")
}

type LayoutDefHdrLst struct{ CT_DiagramDefinitionHeaderLst }

func (_daee ST_Offset) ValidateWithPath(path string) error {
	switch _daee {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_daee))
	}
	return nil
}

func (_eaedg ST_LinearDirection) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_ffaad := _a.Attr{}
	_ffaad.Name = name
	switch _eaedg {
	case ST_LinearDirectionUnset:
		_ffaad.Value = ""
	case ST_LinearDirectionFromL:
		_ffaad.Value = "\u0066\u0072\u006fm\u004c"
	case ST_LinearDirectionFromR:
		_ffaad.Value = "\u0066\u0072\u006fm\u0052"
	case ST_LinearDirectionFromT:
		_ffaad.Value = "\u0066\u0072\u006fm\u0054"
	case ST_LinearDirectionFromB:
		_ffaad.Value = "\u0066\u0072\u006fm\u0042"
	}
	return _ffaad, nil
}

type ST_NodeVerticalAlignment byte

const (
	ST_HierarchyAlignmentUnset   ST_HierarchyAlignment = 0
	ST_HierarchyAlignmentTL      ST_HierarchyAlignment = 1
	ST_HierarchyAlignmentTR      ST_HierarchyAlignment = 2
	ST_HierarchyAlignmentTCtrCh  ST_HierarchyAlignment = 3
	ST_HierarchyAlignmentTCtrDes ST_HierarchyAlignment = 4
	ST_HierarchyAlignmentBL      ST_HierarchyAlignment = 5
	ST_HierarchyAlignmentBR      ST_HierarchyAlignment = 6
	ST_HierarchyAlignmentBCtrCh  ST_HierarchyAlignment = 7
	ST_HierarchyAlignmentBCtrDes ST_HierarchyAlignment = 8
	ST_HierarchyAlignmentLT      ST_HierarchyAlignment = 9
	ST_HierarchyAlignmentLB      ST_HierarchyAlignment = 10
	ST_HierarchyAlignmentLCtrCh  ST_HierarchyAlignment = 11
	ST_HierarchyAlignmentLCtrDes ST_HierarchyAlignment = 12
	ST_HierarchyAlignmentRT      ST_HierarchyAlignment = 13
	ST_HierarchyAlignmentRB      ST_HierarchyAlignment = 14
	ST_HierarchyAlignmentRCtrCh  ST_HierarchyAlignment = 15
	ST_HierarchyAlignmentRCtrDes ST_HierarchyAlignment = 16
)

func (_dbac *CT_StyleLabel) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", _dbac.NameAttr)})
	e.EncodeToken(start)
	if _dbac.Scene3d != nil {
		_dgae := _a.StartElement{Name: _a.Name{Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}}
		e.EncodeElement(_dbac.Scene3d, _dgae)
	}
	if _dbac.Sp3d != nil {
		_eced := _a.StartElement{Name: _a.Name{Local: "\u0073\u0070\u0033\u0064"}}
		e.EncodeElement(_dbac.Sp3d, _eced)
	}
	if _dbac.TxPr != nil {
		_adcaa := _a.StartElement{Name: _a.Name{Local: "\u0074\u0078\u0050\u0072"}}
		e.EncodeElement(_dbac.TxPr, _adcaa)
	}
	if _dbac.Style != nil {
		_fgef := _a.StartElement{Name: _a.Name{Local: "\u0073\u0074\u0079l\u0065"}}
		e.EncodeElement(_dbac.Style, _fgef)
	}
	if _dbac.ExtLst != nil {
		_fcfb := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_dbac.ExtLst, _fcfb)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_bfg *CT_Categories) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _bfg.Cat != nil {
		_dec := _a.StartElement{Name: _a.Name{Local: "\u0063\u0061\u0074"}}
		for _, _aac := range _bfg.Cat {
			e.EncodeElement(_aac, _dec)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_addb *LayoutDef) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006ca\u0079\u006f\u0075\u0074\u0044\u0065f"
	return _addb.CT_DiagramDefinition.MarshalXML(e, start)
}

func (_geda *ST_ContinueDirection) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_befg, _adbce := d.Token()
	if _adbce != nil {
		return _adbce
	}
	if _fecbg, _eceda := _befg.(_a.EndElement); _eceda && _fecbg.Name == start.Name {
		*_geda = 1
		return nil
	}
	if _ceae, _bgcc := _befg.(_a.CharData); !_bgcc {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _befg)
	} else {
		switch string(_ceae) {
		case "":
			*_geda = 0
		case "\u0072\u0065\u0076\u0044\u0069\u0072":
			*_geda = 1
		case "\u0073a\u006d\u0065\u0044\u0069\u0072":
			*_geda = 2
		}
	}
	_befg, _adbce = d.Token()
	if _adbce != nil {
		return _adbce
	}
	if _eegfe, _bbdf := _befg.(_a.EndElement); _bbdf && _eegfe.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _befg)
}

// ValidateWithPath validates the AG_IteratorAttributes and its children, prefixing error messages with path
func (_dgc *AG_IteratorAttributes) ValidateWithPath(path string) error { return nil }

// ValidateWithPath validates the CT_Otherwise and its children, prefixing error messages with path
func (_fdcg *CT_Otherwise) ValidateWithPath(path string) error {
	for _acfb, _ddfd := range _fdcg.Alg {
		if _fbde := _ddfd.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0041\u006c\u0067\u005b\u0025\u0064\u005d", path, _acfb)); _fbde != nil {
			return _fbde
		}
	}
	for _gbdd, _cceb := range _fdcg.Shape {
		if _ddeb := _cceb.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fS\u0068\u0061\u0070\u0065\u005b\u0025\u0064\u005d", path, _gbdd)); _ddeb != nil {
			return _ddeb
		}
	}
	for _adfb, _fbdg := range _fdcg.PresOf {
		if _ggbc := _fbdg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0050\u0072\u0065\u0073\u004f\u0066\u005b\u0025\u0064\u005d", path, _adfb)); _ggbc != nil {
			return _ggbc
		}
	}
	for _bbaac, _cfge := range _fdcg.ConstrLst {
		if _bgfe := _cfge.ValidateWithPath(_f.Sprintf("\u0025\u0073/\u0043\u006f\u006es\u0074\u0072\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _bbaac)); _bgfe != nil {
			return _bgfe
		}
	}
	for _cfda, _afgg := range _fdcg.RuleLst {
		if _bdag := _afgg.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0052\u0075\u006c\u0065\u004c\u0073t\u005b\u0025\u0064\u005d", path, _cfda)); _bdag != nil {
			return _bdag
		}
	}
	for _bcab, _geac := range _fdcg.ForEach {
		if _gfbg := _geac.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0046\u006f\u0072\u0045\u0061\u0063h\u005b\u0025\u0064\u005d", path, _bcab)); _gfbg != nil {
			return _gfbg
		}
	}
	for _fbdce, _ddad := range _fdcg.LayoutNode {
		if _bbad := _ddad.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064e\u005b\u0025\u0064\u005d", path, _fbdce)); _bbad != nil {
			return _bbad
		}
	}
	for _eacef, _cbge := range _fdcg.Choose {
		if _ecdg := _cbge.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0068\u006f\u006f\u0073\u0065\u005b\u0025\u0064\u005d", path, _eacef)); _ecdg != nil {
			return _ecdg
		}
	}
	for _agcbd, _accge := range _fdcg.ExtLst {
		if _cdee := _accge.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0045\u0078\u0074\u004c\u0073\u0074\u005b\u0025\u0064\u005d", path, _agcbd)); _cdee != nil {
			return _cdee
		}
	}
	return nil
}

func (_gaabe ST_OutputShapeType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_eggc := _a.Attr{}
	_eggc.Name = name
	switch _gaabe {
	case ST_OutputShapeTypeUnset:
		_eggc.Value = ""
	case ST_OutputShapeTypeNone:
		_eggc.Value = "\u006e\u006f\u006e\u0065"
	case ST_OutputShapeTypeConn:
		_eggc.Value = "\u0063\u006f\u006e\u006e"
	}
	return _eggc, nil
}

type CT_LayoutVariablePropertySet struct {
	OrgChart      *CT_OrgChart
	ChMax         *CT_ChildMax
	ChPref        *CT_ChildPref
	BulletEnabled *CT_BulletEnabled
	Dir           *CT_Direction
	HierBranch    *CT_HierBranchStyle
	AnimOne       *CT_AnimOne
	AnimLvl       *CT_AnimLvl
	ResizeHandles *CT_ResizeHandles
}

func (_cgfd *ST_LayoutShapeType) ValidateWithPath(path string) error {
	_fgaa := []string{}
	if _cgfd.ST_ShapeType != _bg.ST_ShapeTypeUnset {
		_fgaa = append(_fgaa, "\u0053\u0054\u005fS\u0068\u0061\u0070\u0065\u0054\u0079\u0070\u0065")
	}
	if _cgfd.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		_fgaa = append(_fgaa, "\u0053T\u005fO\u0075\u0074\u0070\u0075\u0074S\u0068\u0061p\u0065\u0054\u0079\u0070\u0065")
	}
	if len(_fgaa) > 1 {
		return _f.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _fgaa)
	}
	return nil
}

type ST_Ints []int32

func NewCT_CTStyleLabel() *CT_CTStyleLabel { _cdd := &CT_CTStyleLabel{}; return _cdd }

func (_ccbac *ST_PyramidAccentPosition) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_eadg, _cddgb := d.Token()
	if _cddgb != nil {
		return _cddgb
	}
	if _eebge, _gfgf := _eadg.(_a.EndElement); _gfgf && _eebge.Name == start.Name {
		*_ccbac = 1
		return nil
	}
	if _fbfe, _bcaeg := _eadg.(_a.CharData); !_bcaeg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _eadg)
	} else {
		switch string(_fbfe) {
		case "":
			*_ccbac = 0
		case "\u0062\u0065\u0066":
			*_ccbac = 1
		case "\u0061\u0066\u0074":
			*_ccbac = 2
		}
	}
	_eadg, _cddgb = d.Token()
	if _cddgb != nil {
		return _cddgb
	}
	if _efeg, _bbcda := _eadg.(_a.EndElement); _bbcda && _efeg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _eadg)
}

func (_bcga ST_FunctionOperator) ValidateWithPath(path string) error {
	switch _bcga {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bcga))
	}
	return nil
}

// Validate validates the CT_ForEach and its children
func (_ecea *CT_ForEach) Validate() error {
	return _ecea.ValidateWithPath("\u0043\u0054\u005f\u0046\u006f\u0072\u0045\u0061\u0063\u0068")
}

func (_cce *CT_Algorithm) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_cce.TypeAttr = ST_AlgorithmType(1)
	for _, _cfce := range start.Attr {
		if _cfce.Name.Local == "\u0074\u0079\u0070\u0065" {
			_cce.TypeAttr.UnmarshalXMLAttr(_cfce)
			continue
		}
		if _cfce.Name.Local == "\u0072\u0065\u0076" {
			_eae, _gee := _ae.ParseUint(_cfce.Value, 10, 32)
			if _gee != nil {
				return _gee
			}
			_ggd := uint32(_eae)
			_cce.RevAttr = &_ggd
			continue
		}
	}
_daf:
	for {
		_dgg, _agce := d.Token()
		if _agce != nil {
			return _agce
		}
		switch _bdg := _dgg.(type) {
		case _a.StartElement:
			switch _bdg.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0061\u0072a\u006d"}:
				_edd := NewCT_Parameter()
				if _bgg := d.DecodeElement(_edd, &_bdg); _bgg != nil {
					return _bgg
				}
				_cce.Param = append(_cce.Param, _edd)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_cce.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _geg := d.DecodeElement(_cce.ExtLst, &_bdg); _geg != nil {
					return _geg
				}
			default:
				_b.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_A\u006c\u0067o\u0072\u0069\u0074\u0068\u006d\u0020\u0025\u0076", _bdg.Name)
				if _fcgf := d.Skip(); _fcgf != nil {
					return _fcgf
				}
			}
		case _a.EndElement:
			break _daf
		case _a.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the LayoutDefHdr and its children, prefixing error messages with path
func (_gfacf *LayoutDefHdr) ValidateWithPath(path string) error {
	if _edgbe := _gfacf.CT_DiagramDefinitionHeader.ValidateWithPath(path); _edgbe != nil {
		return _edgbe
	}
	return nil
}

func (_dgcg *CT_CTCategory) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _cdaa := range start.Attr {
		if _cdaa.Name.Local == "\u0074\u0079\u0070\u0065" {
			_baea, _agbb := _cdaa.Value, error(nil)
			if _agbb != nil {
				return _agbb
			}
			_dgcg.TypeAttr = _baea
			continue
		}
		if _cdaa.Name.Local == "\u0070\u0072\u0069" {
			_ebg, _bgaf := _ae.ParseUint(_cdaa.Value, 10, 32)
			if _bgaf != nil {
				return _bgaf
			}
			_dgcg.PriAttr = uint32(_ebg)
			continue
		}
	}
	for {
		_bcd, _bdee := d.Token()
		if _bdee != nil {
			return _f.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0043\u0054\u005f\u0043T\u0043a\u0074e\u0067\u006f\u0072\u0079\u003a\u0020\u0025s", _bdee)
		}
		if _agdf, _caa := _bcd.(_a.EndElement); _caa && _agdf.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_ConstraintAttributes and its children
func (_fed *AG_ConstraintAttributes) Validate() error {
	return _fed.ValidateWithPath("\u0041\u0047\u005fCo\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074\u0041\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0073")
}

// ValidateWithPath validates the ColorsDef and its children, prefixing error messages with path
func (_eaabd *ColorsDef) ValidateWithPath(path string) error {
	if _fcbbbb := _eaabd.CT_ColorTransform.ValidateWithPath(path); _fcbbbb != nil {
		return _fcbbbb
	}
	return nil
}

func NewCT_ElemPropSet() *CT_ElemPropSet { _bedd := &CT_ElemPropSet{}; return _bedd }

type ColorsDefHdrLst struct{ CT_ColorTransformHeaderLst }

func (_ggfac ST_FunctionOperator) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_bcbbe := _a.Attr{}
	_bcbbe.Name = name
	switch _ggfac {
	case ST_FunctionOperatorUnset:
		_bcbbe.Value = ""
	case ST_FunctionOperatorEqu:
		_bcbbe.Value = "\u0065\u0071\u0075"
	case ST_FunctionOperatorNeq:
		_bcbbe.Value = "\u006e\u0065\u0071"
	case ST_FunctionOperatorGt:
		_bcbbe.Value = "\u0067\u0074"
	case ST_FunctionOperatorLt:
		_bcbbe.Value = "\u006c\u0074"
	case ST_FunctionOperatorGte:
		_bcbbe.Value = "\u0067\u0074\u0065"
	case ST_FunctionOperatorLte:
		_bcbbe.Value = "\u006c\u0074\u0065"
	}
	return _bcbbe, nil
}

func (_cedea ST_NodeVerticalAlignment) Validate() error { return _cedea.ValidateWithPath("") }

type CT_SDName struct {
	LangAttr *string
	ValAttr  string
}

func (_ddb *CT_Algorithm) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	_cdc, _ggf := _ddb.TypeAttr.MarshalXMLAttr(_a.Name{Local: "\u0074\u0079\u0070\u0065"})
	if _ggf != nil {
		return _ggf
	}
	start.Attr = append(start.Attr, _cdc)
	if _ddb.RevAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u0065\u0076"}, Value: _f.Sprintf("\u0025\u0076", *_ddb.RevAttr)})
	}
	e.EncodeToken(start)
	if _ddb.Param != nil {
		_edb := _a.StartElement{Name: _a.Name{Local: "\u0070\u0061\u0072a\u006d"}}
		for _, _daad := range _ddb.Param {
			e.EncodeElement(_daad, _edb)
		}
	}
	if _ddb.ExtLst != nil {
		_agd := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_ddb.ExtLst, _agd)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_cdfd *CT_ChildMax) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _bcg := range start.Attr {
		if _bcg.Name.Local == "\u0076\u0061\u006c" {
			_dccd, _decd := _ae.ParseInt(_bcg.Value, 10, 32)
			if _decd != nil {
				return _decd
			}
			_abee := int32(_dccd)
			_cdfd.ValAttr = &_abee
			continue
		}
	}
	for {
		_fea, _gcf := d.Token()
		if _gcf != nil {
			return _f.Errorf("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u004d\u0061\u0078\u003a\u0020\u0025\u0073", _gcf)
		}
		if _bbg, _agcg := _fea.(_a.EndElement); _agcg && _bbg.Name == start.Name {
			break
		}
	}
	return nil
}

func (_ccae *ST_ElementType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ccae = 0
	case "\u0061\u006c\u006c":
		*_ccae = 1
	case "\u0064\u006f\u0063":
		*_ccae = 2
	case "\u006e\u006f\u0064\u0065":
		*_ccae = 3
	case "\u006e\u006f\u0072\u006d":
		*_ccae = 4
	case "\u006eo\u006e\u004e\u006f\u0072\u006d":
		*_ccae = 5
	case "\u0061\u0073\u0073\u0074":
		*_ccae = 6
	case "\u006eo\u006e\u0041\u0073\u0073\u0074":
		*_ccae = 7
	case "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073":
		*_ccae = 8
	case "\u0070\u0072\u0065\u0073":
		*_ccae = 9
	case "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073":
		*_ccae = 10
	}
	return nil
}

// ValidateWithPath validates the AG_ConstraintRefAttributes and its children, prefixing error messages with path
func (_cb *AG_ConstraintRefAttributes) ValidateWithPath(path string) error {
	if _fag := _cb.RefTypeAttr.ValidateWithPath(path + "\u002f\u0052\u0065f\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _fag != nil {
		return _fag
	}
	if _daa := _cb.RefForAttr.ValidateWithPath(path + "/\u0052\u0065\u0066\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _daa != nil {
		return _daa
	}
	if _ag := _cb.RefPtTypeAttr.ValidateWithPath(path + "\u002f\u0052\u0065\u0066\u0050\u0074\u0054\u0079\u0070e\u0041\u0074\u0074\u0072"); _ag != nil {
		return _ag
	}
	return nil
}

// ValidateWithPath validates the CT_BulletEnabled and its children, prefixing error messages with path
func (_ac *CT_BulletEnabled) ValidateWithPath(path string) error { return nil }

// Validate validates the DataModel and its children
func (_ccgg *DataModel) Validate() error {
	return _ccgg.ValidateWithPath("\u0044a\u0074\u0061\u004d\u006f\u0064\u0065l")
}

func (_dfbe ST_SecondaryLinearDirection) String() string {
	switch _dfbe {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0066\u0072\u006fm\u004c"
	case 3:
		return "\u0066\u0072\u006fm\u0052"
	case 4:
		return "\u0066\u0072\u006fm\u0054"
	case 5:
		return "\u0066\u0072\u006fm\u0042"
	}
	return ""
}

func (_fbcdg ST_RotationPath) String() string {
	switch _fbcdg {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0061l\u006f\u006e\u0067\u0050\u0061\u0074h"
	}
	return ""
}

func (_gecf *CT_Constraints) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_gadf:
	for {
		_baag, _ccg := d.Token()
		if _ccg != nil {
			return _ccg
		}
		switch _dgd := _baag.(type) {
		case _a.StartElement:
			switch _dgd.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u006f\u006e\u0073\u0074\u0072"}:
				_cca := NewCT_Constraint()
				if _adbc := d.DecodeElement(_cca, &_dgd); _adbc != nil {
					return _adbc
				}
				_gecf.Constr = append(_gecf.Constr, _cca)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061i\u006et\u0073\u0020\u0025\u0076", _dgd.Name)
				if _fef := d.Skip(); _fef != nil {
					return _fef
				}
			}
		case _a.EndElement:
			break _gadf
		case _a.CharData:
		}
	}
	return nil
}

func NewCT_DataModel() *CT_DataModel {
	_gdbc := &CT_DataModel{}
	_gdbc.PtLst = NewCT_PtList()
	return _gdbc
}

const (
	ST_PyramidAccentPositionUnset ST_PyramidAccentPosition = 0
	ST_PyramidAccentPositionBef   ST_PyramidAccentPosition = 1
	ST_PyramidAccentPositionAft   ST_PyramidAccentPosition = 2
)

// Validate validates the ColorsDefHdrLst and its children
func (_acaa *ColorsDefHdrLst) Validate() error {
	return _acaa.ValidateWithPath("\u0043o\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074")
}

func NewCT_When() *CT_When {
	_dcfca := &CT_When{}
	_dcfca.FuncAttr = ST_FunctionType(1)
	_dcfca.OpAttr = ST_FunctionOperator(1)
	return _dcfca
}

// Validate validates the CT_Constraints and its children
func (_decg *CT_Constraints) Validate() error {
	return _decg.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072a\u0069\u006e\u0074\u0073")
}

func (_adgbe ST_ResizeHandlesStr) ValidateWithPath(path string) error {
	switch _adgbe {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_adgbe))
	}
	return nil
}

func (_dded *CT_DiagramDefinitionHeader) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", _dded.UniqueIdAttr)})
	if _dded.MinVerAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _f.Sprintf("\u0025\u0076", *_dded.MinVerAttr)})
	}
	if _dded.DefStyleAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_dded.DefStyleAttr)})
	}
	if _dded.ResIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u0065\u0073I\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_dded.ResIdAttr)})
	}
	e.EncodeToken(start)
	_cecf := _a.StartElement{Name: _a.Name{Local: "\u0074\u0069\u0074l\u0065"}}
	for _, _ebbg := range _dded.Title {
		e.EncodeElement(_ebbg, _cecf)
	}
	_dea := _a.StartElement{Name: _a.Name{Local: "\u0064\u0065\u0073\u0063"}}
	for _, _gaba := range _dded.Desc {
		e.EncodeElement(_gaba, _dea)
	}
	if _dded.CatLst != nil {
		_ebfc := _a.StartElement{Name: _a.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_dded.CatLst, _ebfc)
	}
	if _dded.ExtLst != nil {
		_dfeab := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_dded.ExtLst, _dfeab)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

const (
	ST_StartingElementUnset ST_StartingElement = 0
	ST_StartingElementNode  ST_StartingElement = 1
	ST_StartingElementTrans ST_StartingElement = 2
)

func NewStyleDefHdr() *StyleDefHdr {
	_fgec := &StyleDefHdr{}
	_fgec.CT_StyleDefinitionHeader = *NewCT_StyleDefinitionHeader()
	return _fgec
}

func (_gacdf *ST_HueDir) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_gacdf = 0
	case "\u0063\u0077":
		*_gacdf = 1
	case "\u0063\u0063\u0077":
		*_gacdf = 2
	}
	return nil
}

// ValidateWithPath validates the CT_SDName and its children, prefixing error messages with path
func (_cded *CT_SDName) ValidateWithPath(path string) error { return nil }

func (_dedc *ST_ResizeHandlesStr) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gafef, _abbeb := d.Token()
	if _abbeb != nil {
		return _abbeb
	}
	if _aeeb, _bcgfc := _gafef.(_a.EndElement); _bcgfc && _aeeb.Name == start.Name {
		*_dedc = 1
		return nil
	}
	if _bfgb, _acea := _gafef.(_a.CharData); !_acea {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gafef)
	} else {
		switch string(_bfgb) {
		case "":
			*_dedc = 0
		case "\u0065\u0078\u0061c\u0074":
			*_dedc = 1
		case "\u0072\u0065\u006c":
			*_dedc = 2
		}
	}
	_gafef, _abbeb = d.Token()
	if _abbeb != nil {
		return _abbeb
	}
	if _fgaag, _cbbff := _gafef.(_a.EndElement); _cbbff && _fgaag.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gafef)
}

type ST_OutputShapeType byte

func (_fgdacd *ST_TextAnchorVertical) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_fgdacd = 0
	case "\u0074":
		*_fgdacd = 1
	case "\u006d\u0069\u0064":
		*_fgdacd = 2
	case "\u0062":
		*_fgdacd = 3
	}
	return nil
}

// ValidateWithPath validates the CT_DiagramDefinitionHeaderLst and its children, prefixing error messages with path
func (_agfd *CT_DiagramDefinitionHeaderLst) ValidateWithPath(path string) error {
	for _cbbg, _cbbe := range _agfd.LayoutDefHdr {
		if _fagd := _cbbe.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u004cay\u006f\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072\u005b\u0025\u0064\u005d", path, _cbbg)); _fagd != nil {
			return _fagd
		}
	}
	return nil
}

func (_ffeead ST_VariableType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_efdcdd := _a.Attr{}
	_efdcdd.Name = name
	switch _ffeead {
	case ST_VariableTypeUnset:
		_efdcdd.Value = ""
	case ST_VariableTypeNone:
		_efdcdd.Value = "\u006e\u006f\u006e\u0065"
	case ST_VariableTypeOrgChart:
		_efdcdd.Value = "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074"
	case ST_VariableTypeChMax:
		_efdcdd.Value = "\u0063\u0068\u004da\u0078"
	case ST_VariableTypeChPref:
		_efdcdd.Value = "\u0063\u0068\u0050\u0072\u0065\u0066"
	case ST_VariableTypeBulEnabled:
		_efdcdd.Value = "\u0062\u0075\u006c\u0045\u006e\u0061\u0062\u006c\u0065\u0064"
	case ST_VariableTypeDir:
		_efdcdd.Value = "\u0064\u0069\u0072"
	case ST_VariableTypeHierBranch:
		_efdcdd.Value = "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"
	case ST_VariableTypeAnimOne:
		_efdcdd.Value = "\u0061n\u0069\u006d\u004f\u006e\u0065"
	case ST_VariableTypeAnimLvl:
		_efdcdd.Value = "\u0061n\u0069\u006d\u004c\u0076\u006c"
	case ST_VariableTypeResizeHandles:
		_efdcdd.Value = "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073"
	}
	return _efdcdd, nil
}

func (_dcde ST_ParameterId) Validate() error { return _dcde.ValidateWithPath("") }

type ST_NodeHorizontalAlignment byte

func (_fafb *ColorsDefHdrLst) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_fafb.CT_ColorTransformHeaderLst = *NewCT_ColorTransformHeaderLst()
_egadb:
	for {
		_abcd, _caac := d.Token()
		if _caac != nil {
			return _caac
		}
		switch _fcbc := _abcd.(type) {
		case _a.StartElement:
			switch _fcbc.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_cbff := NewCT_ColorTransformHeader()
				if _adffg := d.DecodeElement(_cbff, &_fcbc); _adffg != nil {
					return _adffg
				}
				_fafb.ColorsDefHdr = append(_fafb.ColorsDefHdr, _cbff)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u006f\u006c\u006f\u0072\u0073D\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074\u0020\u0025\u0076", _fcbc.Name)
				if _bcba := d.Skip(); _bcba != nil {
					return _bcba
				}
			}
		case _a.EndElement:
			break _egadb
		case _a.CharData:
		}
	}
	return nil
}

func (_eafef ST_HierarchyAlignment) Validate() error { return _eafef.ValidateWithPath("") }

const (
	ST_ChildOrderTypeUnset ST_ChildOrderType = 0
	ST_ChildOrderTypeB     ST_ChildOrderType = 1
	ST_ChildOrderTypeT     ST_ChildOrderType = 2
)

func (_gbd *CT_CTDescription) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _gbd.LangAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _f.Sprintf("\u0025\u0076", *_gbd.LangAttr)})
	}
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", _gbd.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_AnimLvl and its children, prefixing error messages with path
func (_bde *CT_AnimLvl) ValidateWithPath(path string) error {
	if _bgf := _bde.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _bgf != nil {
		return _bgf
	}
	return nil
}

func (_eeee ST_FallbackDimension) Validate() error { return _eeee.ValidateWithPath("") }

func NewCT_Category() *CT_Category { _gde := &CT_Category{}; return _gde }

func (_bfabf *ST_ParameterVal) Validate() error { return _bfabf.ValidateWithPath("") }

func (_gfacg ST_SecondaryChildAlignment) Validate() error { return _gfacg.ValidateWithPath("") }

const (
	ST_ClrAppMethodUnset  ST_ClrAppMethod = 0
	ST_ClrAppMethodSpan   ST_ClrAppMethod = 1
	ST_ClrAppMethodCycle  ST_ClrAppMethod = 2
	ST_ClrAppMethodRepeat ST_ClrAppMethod = 3
)

func (_cccbd ST_HierarchyAlignment) String() string {
	switch _cccbd {
	case 0:
		return ""
	case 1:
		return "\u0074\u004c"
	case 2:
		return "\u0074\u0052"
	case 3:
		return "\u0074\u0043\u0074\u0072\u0043\u0068"
	case 4:
		return "\u0074C\u0074\u0072\u0044\u0065\u0073"
	case 5:
		return "\u0062\u004c"
	case 6:
		return "\u0062\u0052"
	case 7:
		return "\u0062\u0043\u0074\u0072\u0043\u0068"
	case 8:
		return "\u0062C\u0074\u0072\u0044\u0065\u0073"
	case 9:
		return "\u006c\u0054"
	case 10:
		return "\u006c\u0042"
	case 11:
		return "\u006c\u0043\u0074\u0072\u0043\u0068"
	case 12:
		return "\u006cC\u0074\u0072\u0044\u0065\u0073"
	case 13:
		return "\u0072\u0054"
	case 14:
		return "\u0072\u0042"
	case 15:
		return "\u0072\u0043\u0074\u0072\u0043\u0068"
	case 16:
		return "\u0072C\u0074\u0072\u0044\u0065\u0073"
	}
	return ""
}

func (_faec ST_HueDir) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_gccf := _a.Attr{}
	_gccf.Name = name
	switch _faec {
	case ST_HueDirUnset:
		_gccf.Value = ""
	case ST_HueDirCw:
		_gccf.Value = "\u0063\u0077"
	case ST_HueDirCcw:
		_gccf.Value = "\u0063\u0063\u0077"
	}
	return _gccf, nil
}

type ST_ChildOrderType byte

func (_ccfd *ST_ConnectorDimension) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ccfd = 0
	case "\u0031\u0044":
		*_ccfd = 1
	case "\u0032\u0044":
		*_ccfd = 2
	case "\u0063\u0075\u0073\u0074":
		*_ccfd = 3
	}
	return nil
}

func (_gffgg *ST_FallbackDimension) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_fcgc, _gfbe := d.Token()
	if _gfbe != nil {
		return _gfbe
	}
	if _cefg, _cdgc := _fcgc.(_a.EndElement); _cdgc && _cefg.Name == start.Name {
		*_gffgg = 1
		return nil
	}
	if _acbag, _fdcd := _fcgc.(_a.CharData); !_fdcd {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fcgc)
	} else {
		switch string(_acbag) {
		case "":
			*_gffgg = 0
		case "\u0031\u0044":
			*_gffgg = 1
		case "\u0032\u0044":
			*_gffgg = 2
		}
	}
	_fcgc, _gfbe = d.Token()
	if _gfbe != nil {
		return _gfbe
	}
	if _ebgfa, _daga := _fcgc.(_a.EndElement); _daga && _ebgfa.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fcgc)
}

type ST_FlowDirection byte

func (_abecg *CT_SDCategory) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _ccdaf := range start.Attr {
		if _ccdaf.Name.Local == "\u0074\u0079\u0070\u0065" {
			_cbag, _afbf := _ccdaf.Value, error(nil)
			if _afbf != nil {
				return _afbf
			}
			_abecg.TypeAttr = _cbag
			continue
		}
		if _ccdaf.Name.Local == "\u0070\u0072\u0069" {
			_gcbdg, _gcdad := _ae.ParseUint(_ccdaf.Value, 10, 32)
			if _gcdad != nil {
				return _gcdad
			}
			_abecg.PriAttr = uint32(_gcbdg)
			continue
		}
	}
	for {
		_ceea, _fgfda := d.Token()
		if _fgfda != nil {
			return _f.Errorf("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0043\u0054\u005f\u0053D\u0043a\u0074e\u0067\u006f\u0072\u0079\u003a\u0020\u0025s", _fgfda)
		}
		if _edgdc, _bffb := _ceea.(_a.EndElement); _bffb && _edgdc.Name == start.Name {
			break
		}
	}
	return nil
}

func (_cfcffe *StyleDef) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_cfcffe.CT_StyleDefinition = *NewCT_StyleDefinition()
	for _, _gdde := range start.Attr {
		if _gdde.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_aecdb, _gaec := _gdde.Value, error(nil)
			if _gaec != nil {
				return _gaec
			}
			_cfcffe.UniqueIdAttr = &_aecdb
			continue
		}
		if _gdde.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_ddaf, _fdabc := _gdde.Value, error(nil)
			if _fdabc != nil {
				return _fdabc
			}
			_cfcffe.MinVerAttr = &_ddaf
			continue
		}
	}
_accb:
	for {
		_ffeb, _cfgg := d.Token()
		if _cfgg != nil {
			return _cfgg
		}
		switch _bbeaa := _ffeb.(type) {
		case _a.StartElement:
			switch _bbeaa.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_dgcga := NewCT_SDName()
				if _gagg := d.DecodeElement(_dgcga, &_bbeaa); _gagg != nil {
					return _gagg
				}
				_cfcffe.Title = append(_cfcffe.Title, _dgcga)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_bddd := NewCT_SDDescription()
				if _gcdcb := d.DecodeElement(_bddd, &_bbeaa); _gcdcb != nil {
					return _gcdcb
				}
				_cfcffe.Desc = append(_cfcffe.Desc, _bddd)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_cfcffe.CatLst = NewCT_SDCategories()
				if _abcdf := d.DecodeElement(_cfcffe.CatLst, &_bbeaa); _abcdf != nil {
					return _abcdf
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073c\u0065\u006e\u0065\u0033\u0064"}:
				_cfcffe.Scene3d = _bg.NewCT_Scene3D()
				if _aecca := d.DecodeElement(_cfcffe.Scene3d, &_bbeaa); _aecca != nil {
					return _aecca
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}:
				_gfgb := NewCT_StyleLabel()
				if _ecbdg := d.DecodeElement(_gfgb, &_bbeaa); _ecbdg != nil {
					return _ecbdg
				}
				_cfcffe.StyleLbl = append(_cfcffe.StyleLbl, _gfgb)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_cfcffe.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _aege := d.DecodeElement(_cfcffe.ExtLst, &_bbeaa); _aege != nil {
					return _aege
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0020\u0025\u0076", _bbeaa.Name)
				if _bgaa := d.Skip(); _bgaa != nil {
					return _bgaa
				}
			}
		case _a.EndElement:
			break _accb
		case _a.CharData:
		}
	}
	return nil
}

func (_cegabb ST_OutputShapeType) String() string {
	switch _cegabb {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0063\u006f\u006e\u006e"
	}
	return ""
}

func (_fcffg *CT_RelIds) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _abdda := range start.Attr {
		if _abdda.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _abdda.Name.Local == "\u0064\u006d" || _abdda.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _abdda.Name.Local == "\u0064\u006d" {
			_adga, _addc := _abdda.Value, error(nil)
			if _addc != nil {
				return _addc
			}
			_fcffg.DmAttr = _adga
			continue
		}
		if _abdda.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _abdda.Name.Local == "\u006c\u006f" || _abdda.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _abdda.Name.Local == "\u006c\u006f" {
			_fbeg, _dgbbg := _abdda.Value, error(nil)
			if _dgbbg != nil {
				return _dgbbg
			}
			_fcffg.LoAttr = _fbeg
			continue
		}
		if _abdda.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _abdda.Name.Local == "\u0071\u0073" || _abdda.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _abdda.Name.Local == "\u0071\u0073" {
			_bddf, _bebg := _abdda.Value, error(nil)
			if _bebg != nil {
				return _bebg
			}
			_fcffg.QsAttr = _bddf
			continue
		}
		if _abdda.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073" && _abdda.Name.Local == "\u0063\u0073" || _abdda.Name.Space == "\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073" && _abdda.Name.Local == "\u0063\u0073" {
			_cbaa, _ccce := _abdda.Value, error(nil)
			if _ccce != nil {
				return _ccce
			}
			_fcffg.CsAttr = _cbaa
			continue
		}
	}
	for {
		_ddcb, _abca := d.Token()
		if _abca != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0052\u0065\u006cI\u0064\u0073\u003a\u0020\u0025\u0073", _abca)
		}
		if _dcfe, _ddede := _ddcb.(_a.EndElement); _ddede && _dcfe.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CTDescription and its children
func (_dcd *CT_CTDescription) Validate() error {
	return _dcd.ValidateWithPath("\u0043\u0054_\u0043\u0054\u0044e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e")
}

const (
	ST_TextDirectionUnset ST_TextDirection = 0
	ST_TextDirectionFromT ST_TextDirection = 1
	ST_TextDirectionFromB ST_TextDirection = 2
)

func (_ccc *CT_AnimLvl) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _gfg := range start.Attr {
		if _gfg.Name.Local == "\u0076\u0061\u006c" {
			_ccc.ValAttr.UnmarshalXMLAttr(_gfg)
			continue
		}
	}
	for {
		_bdb, _bbb := d.Token()
		if _bbb != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004c\u0076\u006c\u003a\u0020%\u0073", _bbb)
		}
		if _dcf, _bga := _bdb.(_a.EndElement); _bga && _dcf.Name == start.Name {
			break
		}
	}
	return nil
}

func (_cgfg *ColorsDefHdr) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072"
	return _cgfg.CT_ColorTransformHeader.MarshalXML(e, start)
}

func NewAG_IteratorAttributes() *AG_IteratorAttributes { _ebd := &AG_IteratorAttributes{}; return _ebd }

func (_ebbee ST_TextAnchorHorizontal) ValidateWithPath(path string) error {
	switch _ebbee {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ebbee))
	}
	return nil
}

func (_bff *CT_CTStyleLabel) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _gbb := range start.Attr {
		if _gbb.Name.Local == "\u006e\u0061\u006d\u0065" {
			_fab, _ecb := _gbb.Value, error(nil)
			if _ecb != nil {
				return _ecb
			}
			_bff.NameAttr = _fab
			continue
		}
	}
_ce:
	for {
		_cbbb, _aeg := d.Token()
		if _aeg != nil {
			return _aeg
		}
		switch _aee := _cbbb.(type) {
		case _a.StartElement:
			switch _aee.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066\u0069\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"}:
				_bff.FillClrLst = NewCT_Colors()
				if _bca := d.DecodeElement(_bff.FillClrLst, &_aee); _bca != nil {
					return _bca
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006ci\u006e\u0043\u006c\u0072\u004c\u0073t"}:
				_bff.LinClrLst = NewCT_Colors()
				if _gdd := d.DecodeElement(_bff.LinClrLst, &_aee); _gdd != nil {
					return _gdd
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0066\u0066e\u0063\u0074\u0043\u006c\u0072\u004c\u0073\u0074"}:
				_bff.EffectClrLst = NewCT_Colors()
				if _beg := d.DecodeElement(_bff.EffectClrLst, &_aee); _beg != nil {
					return _beg
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "t\u0078\u004c\u0069\u006e\u0043\u006c\u0072\u004c\u0073\u0074"}:
				_bff.TxLinClrLst = NewCT_Colors()
				if _ggg := d.DecodeElement(_bff.TxLinClrLst, &_aee); _ggg != nil {
					return _ggg
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0078\u0046i\u006c\u006c\u0043\u006c\u0072\u004c\u0073\u0074"}:
				_bff.TxFillClrLst = NewCT_Colors()
				if _ced := d.DecodeElement(_bff.TxFillClrLst, &_aee); _ced != nil {
					return _ced
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0078\u0045\u0066\u0066\u0065\u0063\u0074\u0043l\u0072\u004c\u0073\u0074"}:
				_bff.TxEffectClrLst = NewCT_Colors()
				if _eee := d.DecodeElement(_bff.TxEffectClrLst, &_aee); _eee != nil {
					return _eee
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_bff.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _bggbg := d.DecodeElement(_bff.ExtLst, &_aee); _bggbg != nil {
					return _bggbg
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074e\u0064\u0020\u0065\u006c\u0065\u006d\u0065n\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u0054\u0053t\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c\u0020\u0025\u0076", _aee.Name)
				if _fad := d.Skip(); _fad != nil {
					return _fad
				}
			}
		case _a.EndElement:
			break _ce
		case _a.CharData:
		}
	}
	return nil
}

// ValidateWithPath validates the CT_CTName and its children, prefixing error messages with path
func (_eff *CT_CTName) ValidateWithPath(path string) error { return nil }

func (_gcccbb ST_BoolOperator) String() string {
	switch _gcccbb {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u0065\u0071\u0075"
	case 3:
		return "\u0067\u0074\u0065"
	case 4:
		return "\u006c\u0074\u0065"
	}
	return ""
}

func (_dccc ST_ConnectorPoint) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	return e.EncodeElement(_dccc.String(), start)
}

func (_cdcfa ST_ClrAppMethod) String() string {
	switch _cdcfa {
	case 0:
		return ""
	case 1:
		return "\u0073\u0070\u0061\u006e"
	case 2:
		return "\u0063\u0079\u0063l\u0065"
	case 3:
		return "\u0072\u0065\u0070\u0065\u0061\u0074"
	}
	return ""
}

type ST_HueDir byte

// ValidateWithPath validates the CT_Constraint and its children, prefixing error messages with path
func (_eedec *CT_Constraint) ValidateWithPath(path string) error {
	if _aeae := _eedec.OpAttr.ValidateWithPath(path + "\u002fO\u0070\u0041\u0074\u0074\u0072"); _aeae != nil {
		return _aeae
	}
	if _eedec.ExtLst != nil {
		if _faa := _eedec.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _faa != nil {
			return _faa
		}
	}
	if _eafa := _eedec.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _eafa != nil {
		return _eafa
	}
	if _dacb := _eedec.ForAttr.ValidateWithPath(path + "\u002f\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _dacb != nil {
		return _dacb
	}
	if _ggfe := _eedec.PtTypeAttr.ValidateWithPath(path + "/\u0050\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _ggfe != nil {
		return _ggfe
	}
	if _dga := _eedec.RefTypeAttr.ValidateWithPath(path + "\u002f\u0052\u0065f\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072"); _dga != nil {
		return _dga
	}
	if _dfgc := _eedec.RefForAttr.ValidateWithPath(path + "/\u0052\u0065\u0066\u0046\u006f\u0072\u0041\u0074\u0074\u0072"); _dfgc != nil {
		return _dfgc
	}
	if _ecbg := _eedec.RefPtTypeAttr.ValidateWithPath(path + "\u002f\u0052\u0065\u0066\u0050\u0074\u0054\u0079\u0070e\u0041\u0074\u0074\u0072"); _ecbg != nil {
		return _ecbg
	}
	return nil
}

// ValidateWithPath validates the CT_CTCategory and its children, prefixing error messages with path
func (_bdef *CT_CTCategory) ValidateWithPath(path string) error { return nil }

func (_acae *ST_SecondaryLinearDirection) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_fggf, _egcg := d.Token()
	if _egcg != nil {
		return _egcg
	}
	if _gabgec, _daea := _fggf.(_a.EndElement); _daea && _gabgec.Name == start.Name {
		*_acae = 1
		return nil
	}
	if _gfddg, _gcab := _fggf.(_a.CharData); !_gcab {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _fggf)
	} else {
		switch string(_gfddg) {
		case "":
			*_acae = 0
		case "\u006e\u006f\u006e\u0065":
			*_acae = 1
		case "\u0066\u0072\u006fm\u004c":
			*_acae = 2
		case "\u0066\u0072\u006fm\u0052":
			*_acae = 3
		case "\u0066\u0072\u006fm\u0054":
			*_acae = 4
		case "\u0066\u0072\u006fm\u0042":
			*_acae = 5
		}
	}
	_fggf, _egcg = d.Token()
	if _egcg != nil {
		return _egcg
	}
	if _decdf, _cbbfb := _fggf.(_a.EndElement); _cbbfb && _decdf.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _fggf)
}

func (_ggbcf ST_VerticalAlignment) Validate() error { return _ggbcf.ValidateWithPath("") }

func (_ebba ST_BendPoint) ValidateWithPath(path string) error {
	switch _ebba {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_ebba))
	}
	return nil
}

func (_cfcg *CT_AnimOne) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _cfcg.ValAttr != ST_AnimOneStrUnset {
		_ggb, _bef := _cfcg.ValAttr.MarshalXMLAttr(_a.Name{Local: "\u0076\u0061\u006c"})
		if _bef != nil {
			return _bef
		}
		start.Attr = append(start.Attr, _ggb)
	}
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_fgeaeb ST_FunctionOperator) String() string {
	switch _fgeaeb {
	case 0:
		return ""
	case 1:
		return "\u0065\u0071\u0075"
	case 2:
		return "\u006e\u0065\u0071"
	case 3:
		return "\u0067\u0074"
	case 4:
		return "\u006c\u0074"
	case 5:
		return "\u0067\u0074\u0065"
	case 6:
		return "\u006c\u0074\u0065"
	}
	return ""
}

func (_gbadb ST_Offset) Validate() error { return _gbadb.ValidateWithPath("") }

func (_dggf *ST_ConstraintRelationship) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_deffb, _gfcg := d.Token()
	if _gfcg != nil {
		return _gfcg
	}
	if _gaae, _cagbb := _deffb.(_a.EndElement); _cagbb && _gaae.Name == start.Name {
		*_dggf = 1
		return nil
	}
	if _acagc, _bccc := _deffb.(_a.CharData); !_bccc {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _deffb)
	} else {
		switch string(_acagc) {
		case "":
			*_dggf = 0
		case "\u0073\u0065\u006c\u0066":
			*_dggf = 1
		case "\u0063\u0068":
			*_dggf = 2
		case "\u0064\u0065\u0073":
			*_dggf = 3
		}
	}
	_deffb, _gfcg = d.Token()
	if _gfcg != nil {
		return _gfcg
	}
	if _accf, _agge := _deffb.(_a.EndElement); _agge && _accf.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _deffb)
}

type CT_ColorTransformHeader struct {
	UniqueIdAttr string
	MinVerAttr   *string
	ResIdAttr    *int32
	Title        []*CT_CTName
	Desc         []*CT_CTDescription
	CatLst       *CT_CTCategories
	ExtLst       *_bg.CT_OfficeArtExtensionList
}

// ValidateWithPath validates the CT_OrgChart and its children, prefixing error messages with path
func (_fecb *CT_OrgChart) ValidateWithPath(path string) error { return nil }

func (_adfce *CT_Description) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _gbde := range start.Attr {
		if _gbde.Name.Local == "\u006c\u0061\u006e\u0067" {
			_fcf, _gedd := _gbde.Value, error(nil)
			if _gedd != nil {
				return _gedd
			}
			_adfce.LangAttr = &_fcf
			continue
		}
		if _gbde.Name.Local == "\u0076\u0061\u006c" {
			_debf, _gbff := _gbde.Value, error(nil)
			if _gbff != nil {
				return _gbff
			}
			_adfce.ValAttr = _debf
			continue
		}
	}
	for {
		_cdab, _bcad := d.Token()
		if _bcad != nil {
			return _f.Errorf("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fD\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e:\u0020\u0025\u0073", _bcad)
		}
		if _dbdbg, _afbb := _cdab.(_a.EndElement); _afbb && _dbdbg.Name == start.Name {
			break
		}
	}
	return nil
}

func (_gcfe *ColorsDefHdr) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gcfe.CT_ColorTransformHeader = *NewCT_ColorTransformHeader()
	for _, _fegef := range start.Attr {
		if _fegef.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_caebg, _dedb := _fegef.Value, error(nil)
			if _dedb != nil {
				return _dedb
			}
			_gcfe.UniqueIdAttr = _caebg
			continue
		}
		if _fegef.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_aebg, _aacc := _fegef.Value, error(nil)
			if _aacc != nil {
				return _aacc
			}
			_gcfe.MinVerAttr = &_aebg
			continue
		}
		if _fegef.Name.Local == "\u0072\u0065\u0073I\u0064" {
			_gffga, _dabb := _ae.ParseInt(_fegef.Value, 10, 32)
			if _dabb != nil {
				return _dabb
			}
			_acba := int32(_gffga)
			_gcfe.ResIdAttr = &_acba
			continue
		}
	}
_gddd:
	for {
		_gbba, _ecdgf := d.Token()
		if _ecdgf != nil {
			return _ecdgf
		}
		switch _fbege := _gbba.(type) {
		case _a.StartElement:
			switch _fbege.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_egdcf := NewCT_CTName()
				if _fefad := d.DecodeElement(_egdcf, &_fbege); _fefad != nil {
					return _fefad
				}
				_gcfe.Title = append(_gcfe.Title, _egdcf)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_gcgc := NewCT_CTDescription()
				if _dbegc := d.DecodeElement(_gcgc, &_fbege); _dbegc != nil {
					return _dbegc
				}
				_gcfe.Desc = append(_gcfe.Desc, _gcgc)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_gcfe.CatLst = NewCT_CTCategories()
				if _ceac := d.DecodeElement(_gcfe.CatLst, &_fbege); _ceac != nil {
					return _ceac
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_gcfe.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _gcfc := d.DecodeElement(_gcfe.ExtLst, &_fbege); _gcfc != nil {
					return _gcfc
				}
			default:
				_b.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u006flo\u0072\u0073D\u0065\u0066\u0048\u0064\u0072\u0020\u0025\u0076", _fbege.Name)
				if _ddedc := d.Skip(); _ddedc != nil {
					return _ddedc
				}
			}
		case _a.EndElement:
			break _gddd
		case _a.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CxnList and its children
func (_fda *CT_CxnList) Validate() error {
	return _fda.ValidateWithPath("\u0043\u0054\u005f\u0043\u0078\u006e\u004c\u0069\u0073\u0074")
}

func (_eeaa ST_ConnectorRouting) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_daaa := _a.Attr{}
	_daaa.Name = name
	switch _eeaa {
	case ST_ConnectorRoutingUnset:
		_daaa.Value = ""
	case ST_ConnectorRoutingStra:
		_daaa.Value = "\u0073\u0074\u0072\u0061"
	case ST_ConnectorRoutingBend:
		_daaa.Value = "\u0062\u0065\u006e\u0064"
	case ST_ConnectorRoutingCurve:
		_daaa.Value = "\u0063\u0075\u0072v\u0065"
	case ST_ConnectorRoutingLongCurve:
		_daaa.Value = "\u006co\u006e\u0067\u0043\u0075\u0072\u0076e"
	}
	return _daaa, nil
}

// ValidateWithPath validates the LayoutDef and its children, prefixing error messages with path
func (_bcgd *LayoutDef) ValidateWithPath(path string) error {
	if _afaee := _bcgd.CT_DiagramDefinition.ValidateWithPath(path); _afaee != nil {
		return _afaee
	}
	return nil
}

type CT_LayoutNode struct {
	NameAttr     *string
	StyleLblAttr *string
	ChOrderAttr  ST_ChildOrderType
	MoveWithAttr *string
	Alg          []*CT_Algorithm
	Shape        []*CT_Shape
	PresOf       []*CT_PresentationOf
	ConstrLst    []*CT_Constraints
	RuleLst      []*CT_Rules
	VarLst       []*CT_LayoutVariablePropertySet
	ForEach      []*CT_ForEach
	LayoutNode   []*CT_LayoutNode
	Choose       []*CT_Choose
	ExtLst       []*_bg.CT_OfficeArtExtensionList
}

func (_fgefbe ST_VerticalAlignment) ValidateWithPath(path string) error {
	switch _fgefbe {
	case 0, 1, 2, 3, 4:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_fgefbe))
	}
	return nil
}

func (_daca ST_BoolOperator) Validate() error { return _daca.ValidateWithPath("") }

type StyleDefHdrLst struct{ CT_StyleDefinitionHeaderLst }

func (_dgcbd ST_ChildDirection) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_gcfa := _a.Attr{}
	_gcfa.Name = name
	switch _dgcbd {
	case ST_ChildDirectionUnset:
		_gcfa.Value = ""
	case ST_ChildDirectionHorz:
		_gcfa.Value = "\u0068\u006f\u0072\u007a"
	case ST_ChildDirectionVert:
		_gcfa.Value = "\u0076\u0065\u0072\u0074"
	}
	return _gcfa, nil
}

func (_abadd *ST_StartingElement) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_dfde, _dgag := d.Token()
	if _dgag != nil {
		return _dgag
	}
	if _cdcbb, _caegg := _dfde.(_a.EndElement); _caegg && _cdcbb.Name == start.Name {
		*_abadd = 1
		return nil
	}
	if _eeddc, _bbbg := _dfde.(_a.CharData); !_bbbg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dfde)
	} else {
		switch string(_eeddc) {
		case "":
			*_abadd = 0
		case "\u006e\u006f\u0064\u0065":
			*_abadd = 1
		case "\u0074\u0072\u0061n\u0073":
			*_abadd = 2
		}
	}
	_dfde, _dgag = d.Token()
	if _dgag != nil {
		return _dgag
	}
	if _egga, _aeggf := _dfde.(_a.EndElement); _aeggf && _egga.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dfde)
}

// Validate validates the CT_SDName and its children
func (_bafa *CT_SDName) Validate() error {
	return _bafa.ValidateWithPath("\u0043T\u005f\u0053\u0044\u004e\u0061\u006de")
}

// ValidateWithPath validates the LayoutDefHdrLst and its children, prefixing error messages with path
func (_cbbbb *LayoutDefHdrLst) ValidateWithPath(path string) error {
	if _cafbb := _cbbbb.CT_DiagramDefinitionHeaderLst.ValidateWithPath(path); _cafbb != nil {
		return _cafbb
	}
	return nil
}

func (_acaf ST_VariableType) String() string {
	switch _acaf {
	case 0:
		return ""
	case 1:
		return "\u006e\u006f\u006e\u0065"
	case 2:
		return "\u006f\u0072\u0067\u0043\u0068\u0061\u0072\u0074"
	case 3:
		return "\u0063\u0068\u004da\u0078"
	case 4:
		return "\u0063\u0068\u0050\u0072\u0065\u0066"
	case 5:
		return "\u0062\u0075\u006c\u0045\u006e\u0061\u0062\u006c\u0065\u0064"
	case 6:
		return "\u0064\u0069\u0072"
	case 7:
		return "\u0068\u0069\u0065\u0072\u0042\u0072\u0061\u006e\u0063\u0068"
	case 8:
		return "\u0061n\u0069\u006d\u004f\u006e\u0065"
	case 9:
		return "\u0061n\u0069\u006d\u004c\u0076\u006c"
	case 10:
		return "\u0072\u0065\u0073\u0069\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073"
	}
	return ""
}

func (_agcb *CT_Otherwise) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _gaeed := range start.Attr {
		if _gaeed.Name.Local == "\u006e\u0061\u006d\u0065" {
			_daceb, _gfeed := _gaeed.Value, error(nil)
			if _gfeed != nil {
				return _gfeed
			}
			_agcb.NameAttr = &_daceb
			continue
		}
	}
_dbeb:
	for {
		_bbebd, _ececc := d.Token()
		if _ececc != nil {
			return _ececc
		}
		switch _gccb := _bbebd.(type) {
		case _a.StartElement:
			switch _gccb.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0061\u006c\u0067"}:
				_agad := NewCT_Algorithm()
				if _fcae := d.DecodeElement(_agad, &_gccb); _fcae != nil {
					return _fcae
				}
				_agcb.Alg = append(_agcb.Alg, _agad)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0068\u0061p\u0065"}:
				_adaad := NewCT_Shape()
				if _ddacd := d.DecodeElement(_adaad, &_gccb); _ddacd != nil {
					return _ddacd
				}
				_agcb.Shape = append(_agcb.Shape, _adaad)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}:
				_ege := NewCT_PresentationOf()
				if _cadg := d.DecodeElement(_ege, &_gccb); _cadg != nil {
					return _cadg
				}
				_agcb.PresOf = append(_agcb.PresOf, _ege)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}:
				_cfebc := NewCT_Constraints()
				if _cecfe := d.DecodeElement(_cfebc, &_gccb); _cecfe != nil {
					return _cecfe
				}
				_agcb.ConstrLst = append(_agcb.ConstrLst, _cfebc)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}:
				_acg := NewCT_Rules()
				if _gcccc := d.DecodeElement(_acg, &_gccb); _gcccc != nil {
					return _gcccc
				}
				_agcb.RuleLst = append(_agcb.RuleLst, _acg)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}:
				_bfgc := NewCT_ForEach()
				if _ffdd := d.DecodeElement(_bfgc, &_gccb); _ffdd != nil {
					return _ffdd
				}
				_agcb.ForEach = append(_agcb.ForEach, _bfgc)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}:
				_aefd := NewCT_LayoutNode()
				if _faccf := d.DecodeElement(_aefd, &_gccb); _faccf != nil {
					return _faccf
				}
				_agcb.LayoutNode = append(_agcb.LayoutNode, _aefd)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}:
				_gcbdb := NewCT_Choose()
				if _gfga := d.DecodeElement(_gcbdb, &_gccb); _gfga != nil {
					return _gfga
				}
				_agcb.Choose = append(_agcb.Choose, _gcbdb)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_ddce := _bg.NewCT_OfficeArtExtensionList()
				if _aegb := d.DecodeElement(_ddce, &_gccb); _aegb != nil {
					return _aegb
				}
				_agcb.ExtLst = append(_agcb.ExtLst, _ddce)
			default:
				_b.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075n\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_O\u0074\u0068e\u0072\u0077\u0069\u0073\u0065\u0020\u0025\u0076", _gccb.Name)
				if _baef := d.Skip(); _baef != nil {
					return _baef
				}
			}
		case _a.EndElement:
			break _dbeb
		case _a.CharData:
		}
	}
	return nil
}

const (
	ST_ChildDirectionUnset ST_ChildDirection = 0
	ST_ChildDirectionHorz  ST_ChildDirection = 1
	ST_ChildDirectionVert  ST_ChildDirection = 2
)

func (_gffdg *ST_DiagramTextAlignment) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gcgbe, _gggd := d.Token()
	if _gggd != nil {
		return _gggd
	}
	if _daecd, _bbfbf := _gcgbe.(_a.EndElement); _bbfbf && _daecd.Name == start.Name {
		*_gffdg = 1
		return nil
	}
	if _fgcfa, _fegbe := _gcgbe.(_a.CharData); !_fegbe {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gcgbe)
	} else {
		switch string(_fgcfa) {
		case "":
			*_gffdg = 0
		case "\u006c":
			*_gffdg = 1
		case "\u0063\u0074\u0072":
			*_gffdg = 2
		case "\u0072":
			*_gffdg = 3
		}
	}
	_gcgbe, _gggd = d.Token()
	if _gggd != nil {
		return _gggd
	}
	if _afdef, _bdced := _gcgbe.(_a.EndElement); _bdced && _afdef.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gcgbe)
}

func (_bceb ST_AlgorithmType) Validate() error { return _bceb.ValidateWithPath("") }

// ValidateWithPath validates the CT_Description and its children, prefixing error messages with path
func (_bbed *CT_Description) ValidateWithPath(path string) error { return nil }

func (_efgfe *CT_DiagramDefinition) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _efgfe.UniqueIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_efgfe.UniqueIdAttr)})
	}
	if _efgfe.MinVerAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _f.Sprintf("\u0025\u0076", *_efgfe.MinVerAttr)})
	}
	if _efgfe.DefStyleAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0064\u0065\u0066\u0053\u0074\u0079\u006c\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_efgfe.DefStyleAttr)})
	}
	e.EncodeToken(start)
	if _efgfe.Title != nil {
		_ecdcb := _a.StartElement{Name: _a.Name{Local: "\u0074\u0069\u0074l\u0065"}}
		for _, _dbcd := range _efgfe.Title {
			e.EncodeElement(_dbcd, _ecdcb)
		}
	}
	if _efgfe.Desc != nil {
		_ebgec := _a.StartElement{Name: _a.Name{Local: "\u0064\u0065\u0073\u0063"}}
		for _, _edfg := range _efgfe.Desc {
			e.EncodeElement(_edfg, _ebgec)
		}
	}
	if _efgfe.CatLst != nil {
		_fcea := _a.StartElement{Name: _a.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_efgfe.CatLst, _fcea)
	}
	if _efgfe.SampData != nil {
		_bgcb := _a.StartElement{Name: _a.Name{Local: "\u0073\u0061\u006d\u0070\u0044\u0061\u0074\u0061"}}
		e.EncodeElement(_efgfe.SampData, _bgcb)
	}
	if _efgfe.StyleData != nil {
		_bgcbd := _a.StartElement{Name: _a.Name{Local: "\u0073t\u0079\u006c\u0065\u0044\u0061\u0074a"}}
		e.EncodeElement(_efgfe.StyleData, _bgcbd)
	}
	if _efgfe.ClrData != nil {
		_faef := _a.StartElement{Name: _a.Name{Local: "\u0063l\u0072\u0044\u0061\u0074\u0061"}}
		e.EncodeElement(_efgfe.ClrData, _faef)
	}
	_bagc := _a.StartElement{Name: _a.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
	e.EncodeElement(_efgfe.LayoutNode, _bagc)
	if _efgfe.ExtLst != nil {
		_bgce := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_efgfe.ExtLst, _bgce)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_ColorTransformHeaderLst and its children
func (_caaa *CT_ColorTransformHeaderLst) Validate() error {
	return _caaa.ValidateWithPath("\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061n\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065r\u004c\u0073\u0074")
}

func (_bdad ST_PyramidAccentPosition) ValidateWithPath(path string) error {
	switch _bdad {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_bdad))
	}
	return nil
}

type CT_Rules struct{ Rule []*CT_NumericRule }

// Validate validates the CT_HierBranchStyle and its children
func (_cddge *CT_HierBranchStyle) Validate() error {
	return _cddge.ValidateWithPath("\u0043T\u005fH\u0069\u0065\u0072\u0042\u0072a\u006e\u0063h\u0053\u0074\u0079\u006c\u0065")
}

func (_efbe ST_Direction) Validate() error { return _efbe.ValidateWithPath("") }

// ValidateWithPath validates the RelIds and its children, prefixing error messages with path
func (_agfcg *RelIds) ValidateWithPath(path string) error {
	if _ecag := _agfcg.CT_RelIds.ValidateWithPath(path); _ecag != nil {
		return _ecag
	}
	return nil
}

const (
	ST_LinearDirectionUnset ST_LinearDirection = 0
	ST_LinearDirectionFromL ST_LinearDirection = 1
	ST_LinearDirectionFromR ST_LinearDirection = 2
	ST_LinearDirectionFromT ST_LinearDirection = 3
	ST_LinearDirectionFromB ST_LinearDirection = 4
)

const (
	ST_ConnectorPointUnset  ST_ConnectorPoint = 0
	ST_ConnectorPointAuto   ST_ConnectorPoint = 1
	ST_ConnectorPointBCtr   ST_ConnectorPoint = 2
	ST_ConnectorPointCtr    ST_ConnectorPoint = 3
	ST_ConnectorPointMidL   ST_ConnectorPoint = 4
	ST_ConnectorPointMidR   ST_ConnectorPoint = 5
	ST_ConnectorPointTCtr   ST_ConnectorPoint = 6
	ST_ConnectorPointBL     ST_ConnectorPoint = 7
	ST_ConnectorPointBR     ST_ConnectorPoint = 8
	ST_ConnectorPointTL     ST_ConnectorPoint = 9
	ST_ConnectorPointTR     ST_ConnectorPoint = 10
	ST_ConnectorPointRadial ST_ConnectorPoint = 11
)

// Validate validates the CT_Choose and its children
func (_gabd *CT_Choose) Validate() error {
	return _gabd.ValidateWithPath("\u0043T\u005f\u0043\u0068\u006f\u006f\u0073e")
}

const (
	ST_FunctionOperatorUnset ST_FunctionOperator = 0
	ST_FunctionOperatorEqu   ST_FunctionOperator = 1
	ST_FunctionOperatorNeq   ST_FunctionOperator = 2
	ST_FunctionOperatorGt    ST_FunctionOperator = 3
	ST_FunctionOperatorLt    ST_FunctionOperator = 4
	ST_FunctionOperatorGte   ST_FunctionOperator = 5
	ST_FunctionOperatorLte   ST_FunctionOperator = 6
)

func (_fff *CT_Rules) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _fff.Rule != nil {
		_aaae := _a.StartElement{Name: _a.Name{Local: "\u0072\u0075\u006c\u0065"}}
		for _, _beef := range _fff.Rule {
			e.EncodeElement(_beef, _aaae)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

const (
	ST_SecondaryLinearDirectionUnset ST_SecondaryLinearDirection = 0
	ST_SecondaryLinearDirectionNone  ST_SecondaryLinearDirection = 1
	ST_SecondaryLinearDirectionFromL ST_SecondaryLinearDirection = 2
	ST_SecondaryLinearDirectionFromR ST_SecondaryLinearDirection = 3
	ST_SecondaryLinearDirectionFromT ST_SecondaryLinearDirection = 4
	ST_SecondaryLinearDirectionFromB ST_SecondaryLinearDirection = 5
)

func (_gedg *ST_ClrAppMethod) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_gedg = 0
	case "\u0073\u0070\u0061\u006e":
		*_gedg = 1
	case "\u0063\u0079\u0063l\u0065":
		*_gedg = 2
	case "\u0072\u0065\u0070\u0065\u0061\u0074":
		*_gedg = 3
	}
	return nil
}

type ST_VerticalAlignment byte

// ValidateWithPath validates the CT_TextProps and its children, prefixing error messages with path
func (_agcac *CT_TextProps) ValidateWithPath(path string) error {
	if _agcac.Sp3d != nil {
		if _cafab := _agcac.Sp3d.ValidateWithPath(path + "\u002f\u0053\u00703\u0064"); _cafab != nil {
			return _cafab
		}
	}
	if _agcac.FlatTx != nil {
		if _cgcdb := _agcac.FlatTx.ValidateWithPath(path + "\u002fF\u006c\u0061\u0074\u0054\u0078"); _cgcdb != nil {
			return _cgcdb
		}
	}
	return nil
}

func (_fdabf *ST_TextAnchorVertical) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_gbdc, _afgge := d.Token()
	if _afgge != nil {
		return _afgge
	}
	if _bbege, _dadfd := _gbdc.(_a.EndElement); _dadfd && _bbege.Name == start.Name {
		*_fdabf = 1
		return nil
	}
	if _eabd, _afceb := _gbdc.(_a.CharData); !_afceb {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _gbdc)
	} else {
		switch string(_eabd) {
		case "":
			*_fdabf = 0
		case "\u0074":
			*_fdabf = 1
		case "\u006d\u0069\u0064":
			*_fdabf = 2
		case "\u0062":
			*_fdabf = 3
		}
	}
	_gbdc, _afgge = d.Token()
	if _afgge != nil {
		return _afgge
	}
	if _gcea, _dgbbb := _gbdc.(_a.EndElement); _dgbbb && _gcea.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _gbdc)
}

func NewCT_SDCategories() *CT_SDCategories { _gbg := &CT_SDCategories{}; return _gbg }

func (_dcda *CT_TextProps) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	e.EncodeToken(start)
	if _dcda.Sp3d != nil {
		_gbfg := _a.StartElement{Name: _a.Name{Local: "\u0073\u0070\u0033\u0064"}}
		e.EncodeElement(_dcda.Sp3d, _gbfg)
	}
	if _dcda.FlatTx != nil {
		_aecc := _a.StartElement{Name: _a.Name{Local: "\u0066\u006c\u0061\u0074\u0054\u0078"}}
		e.EncodeElement(_dcda.FlatTx, _aecc)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_defda *ST_TextAnchorHorizontal) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ggaf, _cdcgb := d.Token()
	if _cdcgb != nil {
		return _cdcgb
	}
	if _bdbf, _egfed := _ggaf.(_a.EndElement); _egfed && _bdbf.Name == start.Name {
		*_defda = 1
		return nil
	}
	if _eggeb, _abfdg := _ggaf.(_a.CharData); !_abfdg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ggaf)
	} else {
		switch string(_eggeb) {
		case "":
			*_defda = 0
		case "\u006e\u006f\u006e\u0065":
			*_defda = 1
		case "\u0063\u0074\u0072":
			*_defda = 2
		}
	}
	_ggaf, _cdcgb = d.Token()
	if _cdcgb != nil {
		return _cdcgb
	}
	if _cagdg, _dbde := _ggaf.(_a.EndElement); _dbde && _cagdg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ggaf)
}

func (_efac *CT_NumericRule) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _ebab := range start.Attr {
		if _ebab.Name.Local == "\u0076\u0061\u006c" {
			_ffcd, _ddbe := _ae.ParseFloat(_ebab.Value, 64)
			if _ddbe != nil {
				return _ddbe
			}
			_efac.ValAttr = &_ffcd
			continue
		}
		if _ebab.Name.Local == "\u0066\u0061\u0063\u0074" {
			_bgac, _acdg := _ae.ParseFloat(_ebab.Value, 64)
			if _acdg != nil {
				return _acdg
			}
			_efac.FactAttr = &_bgac
			continue
		}
		if _ebab.Name.Local == "\u006d\u0061\u0078" {
			_ddccg, _fbge := _ae.ParseFloat(_ebab.Value, 64)
			if _fbge != nil {
				return _fbge
			}
			_efac.MaxAttr = &_ddccg
			continue
		}
		if _ebab.Name.Local == "\u0074\u0079\u0070\u0065" {
			_efac.TypeAttr.UnmarshalXMLAttr(_ebab)
			continue
		}
		if _ebab.Name.Local == "\u0066\u006f\u0072" {
			_efac.ForAttr.UnmarshalXMLAttr(_ebab)
			continue
		}
		if _ebab.Name.Local == "\u0066o\u0072\u004e\u0061\u006d\u0065" {
			_cafb, _dcbcd := _ebab.Value, error(nil)
			if _dcbcd != nil {
				return _dcbcd
			}
			_efac.ForNameAttr = &_cafb
			continue
		}
		if _ebab.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_efac.PtTypeAttr.UnmarshalXMLAttr(_ebab)
			continue
		}
	}
_bcdb:
	for {
		_baccg, _deabg := d.Token()
		if _deabg != nil {
			return _deabg
		}
		switch _bbaa := _baccg.(type) {
		case _a.StartElement:
			switch _bbaa.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_efac.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _eage := d.DecodeElement(_efac.ExtLst, &_bbaa); _eage != nil {
					return _eage
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004e\u0075\u006d\u0065\u0072\u0069\u0063R\u0075l\u0065\u0020\u0025\u0076", _bbaa.Name)
				if _eaged := d.Skip(); _eaged != nil {
					return _eaged
				}
			}
		case _a.EndElement:
			break _bcdb
		case _a.CharData:
		}
	}
	return nil
}

func (_ecfc *ST_AxisType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_babbc, _adace := d.Token()
	if _adace != nil {
		return _adace
	}
	if _aegf, _eacbe := _babbc.(_a.EndElement); _eacbe && _aegf.Name == start.Name {
		*_ecfc = 1
		return nil
	}
	if _cecge, _dccg := _babbc.(_a.CharData); !_dccg {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _babbc)
	} else {
		switch string(_cecge) {
		case "":
			*_ecfc = 0
		case "\u0073\u0065\u006c\u0066":
			*_ecfc = 1
		case "\u0063\u0068":
			*_ecfc = 2
		case "\u0064\u0065\u0073":
			*_ecfc = 3
		case "\u0064e\u0073\u004f\u0072\u0053\u0065\u006cf":
			*_ecfc = 4
		case "\u0070\u0061\u0072":
			*_ecfc = 5
		case "\u0061\u006e\u0063s\u0074":
			*_ecfc = 6
		case "a\u006e\u0063\u0073\u0074\u004f\u0072\u0053\u0065\u006c\u0066":
			*_ecfc = 7
		case "\u0066o\u006c\u006c\u006f\u0077\u0053\u0069b":
			*_ecfc = 8
		case "\u0070r\u0065\u0063\u0065\u0064\u0053\u0069b":
			*_ecfc = 9
		case "\u0066\u006f\u006c\u006c\u006f\u0077":
			*_ecfc = 10
		case "\u0070\u0072\u0065\u0063\u0065\u0064":
			*_ecfc = 11
		case "\u0072\u006f\u006f\u0074":
			*_ecfc = 12
		case "\u006e\u006f\u006e\u0065":
			*_ecfc = 13
		}
	}
	_babbc, _adace = d.Token()
	if _adace != nil {
		return _adace
	}
	if _feaad, _cbcba := _babbc.(_a.EndElement); _cbcba && _feaad.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _babbc)
}

// Validate validates the CT_DiagramDefinition and its children
func (_gea *CT_DiagramDefinition) Validate() error {
	return _gea.ValidateWithPath("C\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044e\u0066\u0069\u006e\u0069ti\u006f\u006e")
}

func (_aece *ColorsDef) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_aece.CT_ColorTransform = *NewCT_ColorTransform()
	for _, _daaga := range start.Attr {
		if _daaga.Name.Local == "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064" {
			_cbba, _fdbd := _daaga.Value, error(nil)
			if _fdbd != nil {
				return _fdbd
			}
			_aece.UniqueIdAttr = &_cbba
			continue
		}
		if _daaga.Name.Local == "\u006d\u0069\u006e\u0056\u0065\u0072" {
			_bgdc, _bfdf := _daaga.Value, error(nil)
			if _bfdf != nil {
				return _bfdf
			}
			_aece.MinVerAttr = &_bgdc
			continue
		}
	}
_gfecc:
	for {
		_ccfeg, _feef := d.Token()
		if _feef != nil {
			return _feef
		}
		switch _eabb := _ccfeg.(type) {
		case _a.StartElement:
			switch _eabb.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0074\u0069\u0074l\u0065"}:
				_cddbf := NewCT_CTName()
				if _dgdf := d.DecodeElement(_cddbf, &_eabb); _dgdf != nil {
					return _dgdf
				}
				_aece.Title = append(_aece.Title, _cddbf)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0064\u0065\u0073\u0063"}:
				_bbfa := NewCT_CTDescription()
				if _dgggb := d.DecodeElement(_bbfa, &_eabb); _dgggb != nil {
					return _dgggb
				}
				_aece.Desc = append(_aece.Desc, _bbfa)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}:
				_aece.CatLst = NewCT_CTCategories()
				if _bbede := d.DecodeElement(_aece.CatLst, &_eabb); _bbede != nil {
					return _bbede
				}
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0073\u0074\u0079\u006c\u0065\u004c\u0062\u006c"}:
				_aabb := NewCT_CTStyleLabel()
				if _beaa := d.DecodeElement(_aabb, &_eabb); _beaa != nil {
					return _beaa
				}
				_aece.StyleLbl = append(_aece.StyleLbl, _aabb)
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_aece.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _cggeb := d.DecodeElement(_aece.ExtLst, &_eabb); _cggeb != nil {
					return _cggeb
				}
			default:
				_b.Log.Debug("\u0073k\u0069\u0070p\u0069\u006e\u0067\u0020u\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006cem\u0065\u006e\u0074 \u006f\u006e \u0043\u006f\u006c\u006f\u0072\u0073D\u0065\u0066 \u0025\u0076", _eabb.Name)
				if _dagcf := d.Skip(); _dagcf != nil {
					return _dagcf
				}
			}
		case _a.EndElement:
			break _gfecc
		case _a.CharData:
		}
	}
	return nil
}

func (_caegd ST_ChildOrderType) String() string {
	switch _caegd {
	case 0:
		return ""
	case 1:
		return "\u0062"
	case 2:
		return "\u0074"
	}
	return ""
}

func (_dgaf ST_SecondaryChildAlignment) ValidateWithPath(path string) error {
	switch _dgaf {
	case 0, 1, 2, 3, 4, 5:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_dgaf))
	}
	return nil
}

func (_ffacb *ST_DiagramHorizontalAlignment) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ccdfb, _ffaa := d.Token()
	if _ffaa != nil {
		return _ffaa
	}
	if _afdg, _gadbc := _ccdfb.(_a.EndElement); _gadbc && _afdg.Name == start.Name {
		*_ffacb = 1
		return nil
	}
	if _fada, _dabc := _ccdfb.(_a.CharData); !_dabc {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ccdfb)
	} else {
		switch string(_fada) {
		case "":
			*_ffacb = 0
		case "\u006c":
			*_ffacb = 1
		case "\u0063\u0074\u0072":
			*_ffacb = 2
		case "\u0072":
			*_ffacb = 3
		case "\u006e\u006f\u006e\u0065":
			*_ffacb = 4
		}
	}
	_ccdfb, _ffaa = d.Token()
	if _ffaa != nil {
		return _ffaa
	}
	if _degdc, _ecfcb := _ccdfb.(_a.EndElement); _ecfcb && _degdc.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ccdfb)
}

func (_dgfe ST_AxisType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_bbefd := _a.Attr{}
	_bbefd.Name = name
	switch _dgfe {
	case ST_AxisTypeUnset:
		_bbefd.Value = ""
	case ST_AxisTypeSelf:
		_bbefd.Value = "\u0073\u0065\u006c\u0066"
	case ST_AxisTypeCh:
		_bbefd.Value = "\u0063\u0068"
	case ST_AxisTypeDes:
		_bbefd.Value = "\u0064\u0065\u0073"
	case ST_AxisTypeDesOrSelf:
		_bbefd.Value = "\u0064e\u0073\u004f\u0072\u0053\u0065\u006cf"
	case ST_AxisTypePar:
		_bbefd.Value = "\u0070\u0061\u0072"
	case ST_AxisTypeAncst:
		_bbefd.Value = "\u0061\u006e\u0063s\u0074"
	case ST_AxisTypeAncstOrSelf:
		_bbefd.Value = "a\u006e\u0063\u0073\u0074\u004f\u0072\u0053\u0065\u006c\u0066"
	case ST_AxisTypeFollowSib:
		_bbefd.Value = "\u0066o\u006c\u006c\u006f\u0077\u0053\u0069b"
	case ST_AxisTypePrecedSib:
		_bbefd.Value = "\u0070r\u0065\u0063\u0065\u0064\u0053\u0069b"
	case ST_AxisTypeFollow:
		_bbefd.Value = "\u0066\u006f\u006c\u006c\u006f\u0077"
	case ST_AxisTypePreced:
		_bbefd.Value = "\u0070\u0072\u0065\u0063\u0065\u0064"
	case ST_AxisTypeRoot:
		_bbefd.Value = "\u0072\u006f\u006f\u0074"
	case ST_AxisTypeNone:
		_bbefd.Value = "\u006e\u006f\u006e\u0065"
	}
	return _bbefd, nil
}

func NewAG_ConstraintAttributes() *AG_ConstraintAttributes {
	_c := &AG_ConstraintAttributes{}
	return _c
}

func (_eefe *RelIds) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0072\u0065\u006c\u0049\u0064\u0073"
	return _eefe.CT_RelIds.MarshalXML(e, start)
}

func (_ecbf *CT_SDName) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _aaf := range start.Attr {
		if _aaf.Name.Local == "\u006c\u0061\u006e\u0067" {
			_cdge, _fddad := _aaf.Value, error(nil)
			if _fddad != nil {
				return _fddad
			}
			_ecbf.LangAttr = &_cdge
			continue
		}
		if _aaf.Name.Local == "\u0076\u0061\u006c" {
			_agacf, _faaa := _aaf.Value, error(nil)
			if _faaa != nil {
				return _faaa
			}
			_ecbf.ValAttr = _agacf
			continue
		}
	}
	for {
		_cefa, _aadf := d.Token()
		if _aadf != nil {
			return _f.Errorf("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0044\u004ea\u006d\u0065\u003a\u0020\u0025\u0073", _aadf)
		}
		if _efbb, _fddbe := _cefa.(_a.EndElement); _fddbe && _efbb.Name == start.Name {
			break
		}
	}
	return nil
}

func (_egae *CT_Shape) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _egae.RotAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u006f\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_egae.RotAttr)})
	}
	if _egae.TypeAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0074\u0079\u0070\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_egae.TypeAttr)})
	}
	if _egae.BlipAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u003a\u0062\u006c\u0069\u0070"}, Value: _f.Sprintf("\u0025\u0076", *_egae.BlipAttr)})
	}
	if _egae.ZOrderOffAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u007aO\u0072\u0064\u0065\u0072\u004f\u0066f"}, Value: _f.Sprintf("\u0025\u0076", *_egae.ZOrderOffAttr)})
	}
	if _egae.HideGeomAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0068\u0069\u0064\u0065\u0047\u0065\u006f\u006d"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_egae.HideGeomAttr))})
	}
	if _egae.LkTxEntryAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006ck\u0054\u0078\u0045\u006e\u0074\u0072y"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_egae.LkTxEntryAttr))})
	}
	if _egae.BlipPhldrAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0062l\u0069\u0070\u0050\u0068\u006c\u0064r"}, Value: _f.Sprintf("\u0025\u0064", _feae(*_egae.BlipPhldrAttr))})
	}
	e.EncodeToken(start)
	if _egae.AdjLst != nil {
		_ebbb := _a.StartElement{Name: _a.Name{Local: "\u0061\u0064\u006a\u004c\u0073\u0074"}}
		e.EncodeElement(_egae.AdjLst, _ebbb)
	}
	if _egae.ExtLst != nil {
		_gecd := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_egae.ExtLst, _gecd)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_HierBranchStyle and its children, prefixing error messages with path
func (_ccdf *CT_HierBranchStyle) ValidateWithPath(path string) error {
	if _caca := _ccdf.ValAttr.ValidateWithPath(path + "\u002f\u0056\u0061\u006c\u0041\u0074\u0074\u0072"); _caca != nil {
		return _caca
	}
	return nil
}

type ST_ChildAlignment byte

func (_gfacef *ST_OutputShapeType) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_gfacef = 0
	case "\u006e\u006f\u006e\u0065":
		*_gfacef = 1
	case "\u0063\u006f\u006e\u006e":
		*_gfacef = 2
	}
	return nil
}

func (_gcdg ST_ConnectorDimension) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_cgfgf := _a.Attr{}
	_cgfgf.Name = name
	switch _gcdg {
	case ST_ConnectorDimensionUnset:
		_cgfgf.Value = ""
	case ST_ConnectorDimension1D:
		_cgfgf.Value = "\u0031\u0044"
	case ST_ConnectorDimension2D:
		_cgfgf.Value = "\u0032\u0044"
	case ST_ConnectorDimensionCust:
		_cgfgf.Value = "\u0063\u0075\u0073\u0074"
	}
	return _cgfgf, nil
}

type ST_SecondaryChildAlignment byte

type CT_Name struct {
	LangAttr *string
	ValAttr  string
}

func (_bdbaa ST_StartingElement) Validate() error { return _bdbaa.ValidateWithPath("") }

func (_cbece ST_TextDirection) String() string {
	switch _cbece {
	case 0:
		return ""
	case 1:
		return "\u0066\u0072\u006fm\u0054"
	case 2:
		return "\u0066\u0072\u006fm\u0042"
	}
	return ""
}

func NewCT_DiagramDefinitionHeader() *CT_DiagramDefinitionHeader {
	_cdbb := &CT_DiagramDefinitionHeader{}
	return _cdbb
}

func (_cbd *CT_AnimLvl) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _cbd.ValAttr != ST_AnimLvlStrUnset {
		_debb, _ded := _cbd.ValAttr.MarshalXMLAttr(_a.Name{Local: "\u0076\u0061\u006c"})
		if _ded != nil {
			return _ded
		}
		start.Attr = append(start.Attr, _debb)
	}
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_ccafae *ST_ContinueDirection) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_ccafae = 0
	case "\u0072\u0065\u0076\u0044\u0069\u0072":
		*_ccafae = 1
	case "\u0073a\u006d\u0065\u0044\u0069\u0072":
		*_ccafae = 2
	}
	return nil
}

type ST_ConstraintType byte

// Validate validates the LayoutDef and its children
func (_edge *LayoutDef) Validate() error {
	return _edge.ValidateWithPath("\u004ca\u0079\u006f\u0075\u0074\u0044\u0065f")
}

type ST_Direction byte

func (_bgbgd *ColorsDef) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0069"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u0063o\u006c\u006f\u0072\u0073\u0044\u0065f"
	return _bgbgd.CT_ColorTransform.MarshalXML(e, start)
}

func (_agf *CT_ChildPref) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _agf.ValAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", *_agf.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_bggaf *ST_AutoTextRotation) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_bggaf = 0
	case "\u006e\u006f\u006e\u0065":
		*_bggaf = 1
	case "\u0075\u0070\u0072":
		*_bggaf = 2
	case "\u0067\u0072\u0061\u0076":
		*_bggaf = 3
	}
	return nil
}

func (_eedad ST_FallbackDimension) ValidateWithPath(path string) error {
	switch _eedad {
	case 0, 1, 2:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_eedad))
	}
	return nil
}

func (_daab ST_ParameterId) String() string {
	switch _daab {
	case 0:
		return ""
	case 1:
		return "\u0068o\u0072\u007a\u0041\u006c\u0069\u0067n"
	case 2:
		return "\u0076e\u0072\u0074\u0041\u006c\u0069\u0067n"
	case 3:
		return "\u0063\u0068\u0044i\u0072"
	case 4:
		return "\u0063h\u0041\u006c\u0069\u0067\u006e"
	case 5:
		return "\u0073\u0065\u0063\u0043\u0068\u0041\u006c\u0069\u0067\u006e"
	case 6:
		return "\u006c\u0069\u006e\u0044\u0069\u0072"
	case 7:
		return "\u0073e\u0063\u004c\u0069\u006e\u0044\u0069r"
	case 8:
		return "\u0073\u0074\u0045\u006c\u0065\u006d"
	case 9:
		return "\u0062\u0065\u006e\u0064\u0050\u0074"
	case 10:
		return "\u0063\u006f\u006e\u006e\u0052\u006f\u0075\u0074"
	case 11:
		return "\u0062\u0065\u0067\u0053\u0074\u0079"
	case 12:
		return "\u0065\u006e\u0064\u0053\u0074\u0079"
	case 13:
		return "\u0064\u0069\u006d"
	case 14:
		return "\u0072o\u0074\u0050\u0061\u0074\u0068"
	case 15:
		return "\u0063t\u0072\u0053\u0068\u0070\u004d\u0061p"
	case 16:
		return "\u006e\u006f\u0064\u0065\u0048\u006f\u0072\u007a\u0041\u006c\u0069\u0067\u006e"
	case 17:
		return "\u006e\u006f\u0064\u0065\u0056\u0065\u0072\u0074\u0041\u006c\u0069\u0067\u006e"
	case 18:
		return "\u0066\u0061\u006c\u006c\u0062\u0061\u0063\u006b"
	case 19:
		return "\u0074\u0078\u0044i\u0072"
	case 20:
		return "p\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0050\u006f\u0073"
	case 21:
		return "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054\u0078\u004d\u0061\u0072"
	case 22:
		return "\u0074x\u0042\u006c\u0044\u0069\u0072"
	case 23:
		return "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0048\u006f\u0072\u007a"
	case 24:
		return "\u0074\u0078\u0041n\u0063\u0068\u006f\u0072\u0056\u0065\u0072\u0074"
	case 25:
		return "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0048o\u0072\u007a\u0043\u0068"
	case 26:
		return "\u0074\u0078\u0041\u006e\u0063\u0068\u006f\u0072\u0056e\u0072\u0074\u0043\u0068"
	case 27:
		return "\u0070\u0061\u0072\u0054\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e"
	case 28:
		return "\u0070\u0061\u0072\u0054\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e"
	case 29:
		return "\u0073h\u0070T\u0078\u004c\u0054\u0052\u0041\u006c\u0069\u0067\u006e\u0043\u0068"
	case 30:
		return "\u0073h\u0070T\u0078\u0052\u0054\u004c\u0041\u006c\u0069\u0067\u006e\u0043\u0068"
	case 31:
		return "\u0061u\u0074\u006f\u0054\u0078\u0052\u006ft"
	case 32:
		return "\u0067\u0072\u0044i\u0072"
	case 33:
		return "\u0066l\u006f\u0077\u0044\u0069\u0072"
	case 34:
		return "\u0063o\u006e\u0074\u0044\u0069\u0072"
	case 35:
		return "\u0062\u006b\u0070\u0074"
	case 36:
		return "\u006f\u0066\u0066"
	case 37:
		return "\u0068i\u0065\u0072\u0041\u006c\u0069\u0067n"
	case 38:
		return "\u0062\u006b\u0050t\u0046\u0069\u0078\u0065\u0064\u0056\u0061\u006c"
	case 39:
		return "s\u0074\u0042\u0075\u006c\u006c\u0065\u0074\u004c\u0076\u006c"
	case 40:
		return "\u0073\u0074\u0041n\u0067"
	case 41:
		return "\u0073p\u0061\u006e\u0041\u006e\u0067"
	case 42:
		return "\u0061\u0072"
	case 43:
		return "\u006cn\u0053\u0070\u0050\u0061\u0072"
	case 44:
		return "\u006c\u006e\u0053\u0070\u0041\u0066\u0050\u0061\u0072\u0050"
	case 45:
		return "\u006c\u006e\u0053\u0070\u0043\u0068"
	case 46:
		return "\u006cn\u0053\u0070\u0041\u0066\u0043\u0068P"
	case 47:
		return "r\u0074\u0053\u0068\u006f\u0072\u0074\u0044\u0069\u0073\u0074"
	case 48:
		return "\u0061l\u0069\u0067\u006e\u0054\u0078"
	case 49:
		return "p\u0079\u0072\u0061\u004c\u0076\u006c\u004e\u006f\u0064\u0065"
	case 50:
		return "\u0070\u0079r\u0061\u0041\u0063c\u0074\u0042\u006b\u0067\u0064\u004e\u006f\u0064\u0065"
	case 51:
		return "\u0070\u0079\u0072\u0061\u0041\u0063\u0063\u0074\u0054x\u004e\u006f\u0064\u0065"
	case 52:
		return "\u0073r\u0063\u004e\u006f\u0064\u0065"
	case 53:
		return "\u0064s\u0074\u004e\u006f\u0064\u0065"
	case 54:
		return "\u0062\u0065\u0067\u0050\u0074\u0073"
	case 55:
		return "\u0065\u006e\u0064\u0050\u0074\u0073"
	}
	return ""
}

// Validate validates the CT_ResizeHandles and its children
func (_feed *CT_ResizeHandles) Validate() error {
	return _feed.ValidateWithPath("\u0043\u0054_\u0052\u0065\u0073i\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073")
}

func (_caba *ST_Direction) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_abeg, _dggbg := d.Token()
	if _dggbg != nil {
		return _dggbg
	}
	if _fdebb, _febg := _abeg.(_a.EndElement); _febg && _fdebb.Name == start.Name {
		*_caba = 1
		return nil
	}
	if _ccebc, _fbag := _abeg.(_a.CharData); !_fbag {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _abeg)
	} else {
		switch string(_ccebc) {
		case "":
			*_caba = 0
		case "\u006e\u006f\u0072\u006d":
			*_caba = 1
		case "\u0072\u0065\u0076":
			*_caba = 2
		}
	}
	_abeg, _dggbg = d.Token()
	if _dggbg != nil {
		return _dggbg
	}
	if _eedg, _deggc := _abeg.(_a.EndElement); _deggc && _eedg.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _abeg)
}

// Validate validates the CT_SDCategories and its children
func (_aaad *CT_SDCategories) Validate() error {
	return _aaad.ValidateWithPath("\u0043T\u005fS\u0044\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073")
}

func (_ggc *CT_CTName) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _ggc.LangAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _f.Sprintf("\u0025\u0076", *_ggc.LangAttr)})
	}
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", _ggc.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

func (_aaea *CT_When) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _aaea.NameAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006e\u0061\u006d\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_aaea.NameAttr)})
	}
	_cgad, _aeba := _aaea.FuncAttr.MarshalXMLAttr(_a.Name{Local: "\u0066\u0075\u006e\u0063"})
	if _aeba != nil {
		return _aeba
	}
	start.Attr = append(start.Attr, _cgad)
	if _aaea.ArgAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0061\u0072\u0067"}, Value: _f.Sprintf("\u0025\u0076", *_aaea.ArgAttr)})
	}
	_cgad, _aeba = _aaea.OpAttr.MarshalXMLAttr(_a.Name{Local: "\u006f\u0070"})
	if _aeba != nil {
		return _aeba
	}
	start.Attr = append(start.Attr, _cgad)
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", _aaea.ValAttr)})
	if _aaea.AxisAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0061\u0078\u0069\u0073"}, Value: _f.Sprintf("\u0025\u0076", *_aaea.AxisAttr)})
	}
	if _aaea.PtTypeAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0070\u0074\u0054\u0079\u0070\u0065"}, Value: _f.Sprintf("\u0025\u0076", *_aaea.PtTypeAttr)})
	}
	if _aaea.HideLastTransAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0068\u0069\u0064\u0065\u004c\u0061\u0073\u0074\u0054\u0072\u0061\u006e\u0073"}, Value: _f.Sprintf("\u0025\u0076", *_aaea.HideLastTransAttr)})
	}
	if _aaea.StAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_aaea.StAttr)})
	}
	if _aaea.CntAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0063\u006e\u0074"}, Value: _f.Sprintf("\u0025\u0076", *_aaea.CntAttr)})
	}
	if _aaea.StepAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0073\u0074\u0065\u0070"}, Value: _f.Sprintf("\u0025\u0076", *_aaea.StepAttr)})
	}
	e.EncodeToken(start)
	if _aaea.Alg != nil {
		_bfabb := _a.StartElement{Name: _a.Name{Local: "\u0061\u006c\u0067"}}
		for _, _ggdb := range _aaea.Alg {
			e.EncodeElement(_ggdb, _bfabb)
		}
	}
	if _aaea.Shape != nil {
		_daagb := _a.StartElement{Name: _a.Name{Local: "\u0073\u0068\u0061p\u0065"}}
		for _, _eaba := range _aaea.Shape {
			e.EncodeElement(_eaba, _daagb)
		}
	}
	if _aaea.PresOf != nil {
		_bcabg := _a.StartElement{Name: _a.Name{Local: "\u0070\u0072\u0065\u0073\u004f\u0066"}}
		for _, _dcfa := range _aaea.PresOf {
			e.EncodeElement(_dcfa, _bcabg)
		}
	}
	if _aaea.ConstrLst != nil {
		_cdbgc := _a.StartElement{Name: _a.Name{Local: "\u0063o\u006e\u0073\u0074\u0072\u004c\u0073t"}}
		for _, _ccafa := range _aaea.ConstrLst {
			e.EncodeElement(_ccafa, _cdbgc)
		}
	}
	if _aaea.RuleLst != nil {
		_ccga := _a.StartElement{Name: _a.Name{Local: "\u0072u\u006c\u0065\u004c\u0073\u0074"}}
		for _, _baeg := range _aaea.RuleLst {
			e.EncodeElement(_baeg, _ccga)
		}
	}
	if _aaea.ForEach != nil {
		_afgcg := _a.StartElement{Name: _a.Name{Local: "\u0066o\u0072\u0045\u0061\u0063\u0068"}}
		for _, _caeg := range _aaea.ForEach {
			e.EncodeElement(_caeg, _afgcg)
		}
	}
	if _aaea.LayoutNode != nil {
		_fcffc := _a.StartElement{Name: _a.Name{Local: "\u006c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065"}}
		for _, _ddacf := range _aaea.LayoutNode {
			e.EncodeElement(_ddacf, _fcffc)
		}
	}
	if _aaea.Choose != nil {
		_dfcb := _a.StartElement{Name: _a.Name{Local: "\u0063\u0068\u006f\u006f\u0073\u0065"}}
		for _, _dggd := range _aaea.Choose {
			e.EncodeElement(_dggd, _dfcb)
		}
	}
	if _aaea.ExtLst != nil {
		_bddba := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		for _, _gfbf := range _aaea.ExtLst {
			e.EncodeElement(_gfbf, _bddba)
		}
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// ValidateWithPath validates the CT_CTCategories and its children, prefixing error messages with path
func (_ff *CT_CTCategories) ValidateWithPath(path string) error {
	for _bcc, _bed := range _ff.Cat {
		if _dda := _bed.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002f\u0043\u0061\u0074\u005b\u0025\u0064\u005d", path, _bcc)); _dda != nil {
			return _dda
		}
	}
	return nil
}

type CT_Adj struct {
	IdxAttr uint32
	ValAttr float64
}

func (_bgcba ST_PyramidAccentTextMargin) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_adcb := _a.Attr{}
	_adcb.Name = name
	switch _bgcba {
	case ST_PyramidAccentTextMarginUnset:
		_adcb.Value = ""
	case ST_PyramidAccentTextMarginStep:
		_adcb.Value = "\u0073\u0074\u0065\u0070"
	case ST_PyramidAccentTextMarginStack:
		_adcb.Value = "\u0073\u0074\u0061c\u006b"
	}
	return _adcb, nil
}

// Validate validates the CT_When and its children
func (_fcgbe *CT_When) Validate() error {
	return _fcgbe.ValidateWithPath("\u0043T\u005f\u0057\u0068\u0065\u006e")
}

func (_bccbf ST_FunctionType) Validate() error { return _bccbf.ValidateWithPath("") }

func (_agcf *CT_ColorTransformHeader) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0075\u006e\u0069\u0071\u0075\u0065\u0049\u0064"}, Value: _f.Sprintf("\u0025\u0076", _agcf.UniqueIdAttr)})
	if _agcf.MinVerAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006d\u0069\u006e\u0056\u0065\u0072"}, Value: _f.Sprintf("\u0025\u0076", *_agcf.MinVerAttr)})
	}
	if _agcf.ResIdAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0072\u0065\u0073I\u0064"}, Value: _f.Sprintf("\u0025\u0076", *_agcf.ResIdAttr)})
	}
	e.EncodeToken(start)
	_bdba := _a.StartElement{Name: _a.Name{Local: "\u0074\u0069\u0074l\u0065"}}
	for _, _dad := range _agcf.Title {
		e.EncodeElement(_dad, _bdba)
	}
	_fcab := _a.StartElement{Name: _a.Name{Local: "\u0064\u0065\u0073\u0063"}}
	for _, _gad := range _agcf.Desc {
		e.EncodeElement(_gad, _fcab)
	}
	if _agcf.CatLst != nil {
		_adgg := _a.StartElement{Name: _a.Name{Local: "\u0063\u0061\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_agcf.CatLst, _adgg)
	}
	if _agcf.ExtLst != nil {
		_ddfa := _a.StartElement{Name: _a.Name{Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}}
		e.EncodeElement(_agcf.ExtLst, _ddfa)
	}
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

const (
	ST_TextBlockDirectionUnset ST_TextBlockDirection = 0
	ST_TextBlockDirectionHorz  ST_TextBlockDirection = 1
	ST_TextBlockDirectionVert  ST_TextBlockDirection = 2
)

func NewCT_RelIds() *CT_RelIds { _cdga := &CT_RelIds{}; return _cdga }

func ParseUnionST_ParameterVal(s string) (ST_ParameterVal, error) { return ST_ParameterVal{}, nil }

// ValidateWithPath validates the CT_ChildMax and its children, prefixing error messages with path
func (_egb *CT_ChildMax) ValidateWithPath(path string) error {
	if _egb.ValAttr != nil {
		if *_egb.ValAttr < -1 {
			return _f.Errorf("\u0025\u0073/m\u002e\u0056\u0061l\u0041\u0074\u0074\u0072 mu\u0073t \u0062\u0065\u0020\u003e\u003d\u0020\u002d1 \u0028\u0068\u0061\u0076\u0065\u0020\u0025v\u0029", path, *_egb.ValAttr)
		}
	}
	return nil
}

type CT_ChildMax struct{ ValAttr *int32 }

func (_fabgg *CT_Rules) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_fead:
	for {
		_cdcea, _degd := d.Token()
		if _degd != nil {
			return _degd
		}
		switch _ffgg := _cdcea.(type) {
		case _a.StartElement:
			switch _ffgg.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0072\u0075\u006c\u0065"}:
				_eadb := NewCT_NumericRule()
				if _gdbg := d.DecodeElement(_eadb, &_ffgg); _gdbg != nil {
					return _gdbg
				}
				_fabgg.Rule = append(_fabgg.Rule, _eadb)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0075\u006c\u0065\u0073\u0020\u0025\u0076", _ffgg.Name)
				if _cddd := d.Skip(); _cddd != nil {
					return _cddd
				}
			}
		case _a.EndElement:
			break _fead
		case _a.CharData:
		}
	}
	return nil
}

const (
	ST_DiagramTextAlignmentUnset ST_DiagramTextAlignment = 0
	ST_DiagramTextAlignmentL     ST_DiagramTextAlignment = 1
	ST_DiagramTextAlignmentCtr   ST_DiagramTextAlignment = 2
	ST_DiagramTextAlignmentR     ST_DiagramTextAlignment = 3
)

func NewCT_ResizeHandles() *CT_ResizeHandles { _ddgd := &CT_ResizeHandles{}; return _ddgd }

func (_bagge *ST_NodeHorizontalAlignment) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_bagge = 0
	case "\u006c":
		*_bagge = 1
	case "\u0063\u0074\u0072":
		*_bagge = 2
	case "\u0072":
		*_bagge = 3
	}
	return nil
}

type CT_SampleData struct {
	UseDefAttr *bool
	DataModel  *CT_DataModel
}

func (_gfed *CT_Constraint) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	for _, _fbe := range start.Attr {
		if _fbe.Name.Local == "\u006f\u0070" {
			_gfed.OpAttr.UnmarshalXMLAttr(_fbe)
			continue
		}
		if _fbe.Name.Local == "\u0076\u0061\u006c" {
			_fbc, _ffegb := _ae.ParseFloat(_fbe.Value, 64)
			if _ffegb != nil {
				return _ffegb
			}
			_gfed.ValAttr = &_fbc
			continue
		}
		if _fbe.Name.Local == "\u0066\u0061\u0063\u0074" {
			_cdfb, _geb := _ae.ParseFloat(_fbe.Value, 64)
			if _geb != nil {
				return _geb
			}
			_gfed.FactAttr = &_cdfb
			continue
		}
		if _fbe.Name.Local == "\u0074\u0079\u0070\u0065" {
			_gfed.TypeAttr.UnmarshalXMLAttr(_fbe)
			continue
		}
		if _fbe.Name.Local == "\u0066\u006f\u0072" {
			_gfed.ForAttr.UnmarshalXMLAttr(_fbe)
			continue
		}
		if _fbe.Name.Local == "\u0066o\u0072\u004e\u0061\u006d\u0065" {
			_beb, _agac := _fbe.Value, error(nil)
			if _agac != nil {
				return _agac
			}
			_gfed.ForNameAttr = &_beb
			continue
		}
		if _fbe.Name.Local == "\u0070\u0074\u0054\u0079\u0070\u0065" {
			_gfed.PtTypeAttr.UnmarshalXMLAttr(_fbe)
			continue
		}
		if _fbe.Name.Local == "\u0072e\u0066\u0054\u0079\u0070\u0065" {
			_gfed.RefTypeAttr.UnmarshalXMLAttr(_fbe)
			continue
		}
		if _fbe.Name.Local == "\u0072\u0065\u0066\u0046\u006f\u0072" {
			_gfed.RefForAttr.UnmarshalXMLAttr(_fbe)
			continue
		}
		if _fbe.Name.Local == "\u0072\u0065\u0066\u0046\u006f\u0072\u004e\u0061\u006d\u0065" {
			_dcaa, _ecba := _fbe.Value, error(nil)
			if _ecba != nil {
				return _ecba
			}
			_gfed.RefForNameAttr = &_dcaa
			continue
		}
		if _fbe.Name.Local == "\u0072e\u0066\u0050\u0074\u0054\u0079\u0070e" {
			_gfed.RefPtTypeAttr.UnmarshalXMLAttr(_fbe)
			continue
		}
	}
_fabd:
	for {
		_fbdc, _ebf := d.Token()
		if _ebf != nil {
			return _ebf
		}
		switch _ggbd := _fbdc.(type) {
		case _a.StartElement:
			switch _ggbd.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0065\u0078\u0074\u004c\u0073\u0074"}:
				_gfed.ExtLst = _bg.NewCT_OfficeArtExtensionList()
				if _bgae := d.DecodeElement(_gfed.ExtLst, &_ggbd); _bgae != nil {
					return _bgae
				}
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043o\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074 \u0025\u0076", _ggbd.Name)
				if _dbb := d.Skip(); _dbb != nil {
					return _dbb
				}
			}
		case _a.EndElement:
			break _fabd
		case _a.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AdjLst and its children
func (_cafa *CT_AdjLst) Validate() error {
	return _cafa.ValidateWithPath("\u0043T\u005f\u0041\u0064\u006a\u004c\u0073t")
}

// ValidateWithPath validates the CT_Rules and its children, prefixing error messages with path
func (_dgca *CT_Rules) ValidateWithPath(path string) error {
	for _ccda, _gagd := range _dgca.Rule {
		if _gecg := _gagd.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0052\u0075\u006c\u0065\u005b\u0025\u0064\u005d", path, _ccda)); _gecg != nil {
			return _gecg
		}
	}
	return nil
}

func (_ggbfd *CT_SDDescription) MarshalXML(e *_a.Encoder, start _a.StartElement) error {
	if _ggbfd.LangAttr != nil {
		start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u006c\u0061\u006e\u0067"}, Value: _f.Sprintf("\u0025\u0076", *_ggbfd.LangAttr)})
	}
	start.Attr = append(start.Attr, _a.Attr{Name: _a.Name{Local: "\u0076\u0061\u006c"}, Value: _f.Sprintf("\u0025\u0076", _ggbfd.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(_a.EndElement{Name: start.Name})
	return nil
}

// Validate validates the CT_Description and its children
func (_cbef *CT_Description) Validate() error {
	return _cbef.ValidateWithPath("\u0043\u0054\u005f\u0044\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e")
}

func (_geff *CT_StyleDefinitionHeaderLst) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_fgbg:
	for {
		_ddba, _fgcb := d.Token()
		if _fgcb != nil {
			return _fgcb
		}
		switch _ccbc := _ddba.(type) {
		case _a.StartElement:
			switch _ccbc.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_fcabca := NewCT_StyleDefinitionHeader()
				if _gddf := d.DecodeElement(_fcabca, &_ccbc); _gddf != nil {
					return _gddf
				}
				_geff.StyleDefHdr = append(_geff.StyleDefHdr, _fcabca)
			default:
				_b.Log.Debug("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054_\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065a\u0064\u0065\u0072\u004c\u0073\u0074\u0020\u0025\u0076", _ccbc.Name)
				if _ggcbf := d.Skip(); _ggcbf != nil {
					return _ggcbf
				}
			}
		case _a.EndElement:
			break _fgbg
		case _a.CharData:
		}
	}
	return nil
}

func NewCT_Adj() *CT_Adj { _afd := &CT_Adj{}; _afd.IdxAttr = 1; return _afd }

func (_gdgea *ST_FunctionType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_dbebf, _daagf := d.Token()
	if _daagf != nil {
		return _daagf
	}
	if _cadcf, _fadd := _dbebf.(_a.EndElement); _fadd && _cadcf.Name == start.Name {
		*_gdgea = 1
		return nil
	}
	if _bdcf, _adcda := _dbebf.(_a.CharData); !_adcda {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _dbebf)
	} else {
		switch string(_bdcf) {
		case "":
			*_gdgea = 0
		case "\u0063\u006e\u0074":
			*_gdgea = 1
		case "\u0070\u006f\u0073":
			*_gdgea = 2
		case "\u0072\u0065\u0076\u0050\u006f\u0073":
			*_gdgea = 3
		case "\u0070o\u0073\u0045\u0076\u0065\u006e":
			*_gdgea = 4
		case "\u0070\u006f\u0073\u004f\u0064\u0064":
			*_gdgea = 5
		case "\u0076\u0061\u0072":
			*_gdgea = 6
		case "\u0064\u0065\u0070t\u0068":
			*_gdgea = 7
		case "\u006d\u0061\u0078\u0044\u0065\u0070\u0074\u0068":
			*_gdgea = 8
		}
	}
	_dbebf, _daagf = d.Token()
	if _daagf != nil {
		return _daagf
	}
	if _efbae, _fffd := _dbebf.(_a.EndElement); _fffd && _efbae.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _dbebf)
}

// ValidateWithPath validates the CT_ColorTransformHeader and its children, prefixing error messages with path
func (_dabd *CT_ColorTransformHeader) ValidateWithPath(path string) error {
	for _bce, _cac := range _dabd.Title {
		if _ddcc := _cac.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fT\u0069\u0074\u006c\u0065\u005b\u0025\u0064\u005d", path, _bce)); _ddcc != nil {
			return _ddcc
		}
	}
	for _ebcc, _aeee := range _dabd.Desc {
		if _ggga := _aeee.ValidateWithPath(_f.Sprintf("%\u0073\u002f\u0044\u0065\u0073\u0063\u005b\u0025\u0064\u005d", path, _ebcc)); _ggga != nil {
			return _ggga
		}
	}
	if _dabd.CatLst != nil {
		if _gfd := _dabd.CatLst.ValidateWithPath(path + "\u002fC\u0061\u0074\u004c\u0073\u0074"); _gfd != nil {
			return _gfd
		}
	}
	if _dabd.ExtLst != nil {
		if _bdd := _dabd.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _bdd != nil {
			return _bdd
		}
	}
	return nil
}

func (_febag ST_CxnType) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_efbbef := _a.Attr{}
	_efbbef.Name = name
	switch _febag {
	case ST_CxnTypeUnset:
		_efbbef.Value = ""
	case ST_CxnTypeParOf:
		_efbbef.Value = "\u0070\u0061\u0072O\u0066"
	case ST_CxnTypePresOf:
		_efbbef.Value = "\u0070\u0072\u0065\u0073\u004f\u0066"
	case ST_CxnTypePresParOf:
		_efbbef.Value = "\u0070r\u0065\u0073\u0050\u0061\u0072\u004ff"
	case ST_CxnTypeUnknownRelationship:
		_efbbef.Value = "\u0075\u006e\u006b\u006eow\u006e\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"
	}
	return _efbbef, nil
}

// ST_LayoutShapeType is a union type
type ST_LayoutShapeType struct {
	ST_ShapeType       _bg.ST_ShapeType
	ST_OutputShapeType ST_OutputShapeType
}

const (
	ST_AnimLvlStrUnset ST_AnimLvlStr = 0
	ST_AnimLvlStrNone  ST_AnimLvlStr = 1
	ST_AnimLvlStrLvl   ST_AnimLvlStr = 2
	ST_AnimLvlStrCtr   ST_AnimLvlStr = 3
)

func (_dgfc *ST_ModelId) Validate() error { return _dgfc.ValidateWithPath("") }

// Validate validates the ColorsDef and its children
func (_ffdg *ColorsDef) Validate() error {
	return _ffdg.ValidateWithPath("\u0043o\u006c\u006f\u0072\u0073\u0044\u0065f")
}

func NewCT_LayoutVariablePropertySet() *CT_LayoutVariablePropertySet {
	_degbg := &CT_LayoutVariablePropertySet{}
	return _degbg
}

func (_gbdbe ST_ConstraintRelationship) Validate() error { return _gbdbe.ValidateWithPath("") }

// ValidateWithPath validates the CT_Algorithm and its children, prefixing error messages with path
func (_eece *CT_Algorithm) ValidateWithPath(path string) error {
	if _eece.TypeAttr == ST_AlgorithmTypeUnset {
		return _f.Errorf("\u0025\u0073\u002f\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u0069\u0073\u0020a\u0020m\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020\u0066\u0069\u0065\u006c\u0064", path)
	}
	if _dgf := _eece.TypeAttr.ValidateWithPath(path + "\u002fT\u0079\u0070\u0065\u0041\u0074\u0074r"); _dgf != nil {
		return _dgf
	}
	for _bc, _dff := range _eece.Param {
		if _gda := _dff.ValidateWithPath(_f.Sprintf("\u0025\u0073\u002fP\u0061\u0072\u0061\u006d\u005b\u0025\u0064\u005d", path, _bc)); _gda != nil {
			return _gda
		}
	}
	if _eece.ExtLst != nil {
		if _abbb := _eece.ExtLst.ValidateWithPath(path + "\u002fE\u0078\u0074\u004c\u0073\u0074"); _abbb != nil {
			return _abbb
		}
	}
	return nil
}

// ValidateWithPath validates the StyleDef and its children, prefixing error messages with path
func (_eabc *StyleDef) ValidateWithPath(path string) error {
	if _cegf := _eabc.CT_StyleDefinition.ValidateWithPath(path); _cegf != nil {
		return _cegf
	}
	return nil
}

// Validate validates the CT_LayoutNode and its children
func (_cafe *CT_LayoutNode) Validate() error {
	return _cafe.ValidateWithPath("\u0043\u0054\u005f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065")
}

type ST_VariableType byte

type ST_AxisType byte

func ParseUnionST_FunctionValue(s string) (ST_FunctionValue, error) { return ST_FunctionValue{}, nil }

func (_debd ST_ContinueDirection) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_deba := _a.Attr{}
	_deba.Name = name
	switch _debd {
	case ST_ContinueDirectionUnset:
		_deba.Value = ""
	case ST_ContinueDirectionRevDir:
		_deba.Value = "\u0072\u0065\u0076\u0044\u0069\u0072"
	case ST_ContinueDirectionSameDir:
		_deba.Value = "\u0073a\u006d\u0065\u0044\u0069\u0072"
	}
	return _deba, nil
}

func (_fefb *ST_ElementType) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ffagb, _ccca := d.Token()
	if _ccca != nil {
		return _ccca
	}
	if _bgfd, _daefb := _ffagb.(_a.EndElement); _daefb && _bgfd.Name == start.Name {
		*_fefb = 1
		return nil
	}
	if _cccb, _caea := _ffagb.(_a.CharData); !_caea {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ffagb)
	} else {
		switch string(_cccb) {
		case "":
			*_fefb = 0
		case "\u0061\u006c\u006c":
			*_fefb = 1
		case "\u0064\u006f\u0063":
			*_fefb = 2
		case "\u006e\u006f\u0064\u0065":
			*_fefb = 3
		case "\u006e\u006f\u0072\u006d":
			*_fefb = 4
		case "\u006eo\u006e\u004e\u006f\u0072\u006d":
			*_fefb = 5
		case "\u0061\u0073\u0073\u0074":
			*_fefb = 6
		case "\u006eo\u006e\u0041\u0073\u0073\u0074":
			*_fefb = 7
		case "\u0070\u0061\u0072\u0054\u0072\u0061\u006e\u0073":
			*_fefb = 8
		case "\u0070\u0072\u0065\u0073":
			*_fefb = 9
		case "\u0073\u0069\u0062\u0054\u0072\u0061\u006e\u0073":
			*_fefb = 10
		}
	}
	_ffagb, _ccca = d.Token()
	if _ccca != nil {
		return _ccca
	}
	if _fcacc, _acdc := _ffagb.(_a.EndElement); _acdc && _fcacc.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ffagb)
}

const (
	ST_AnimOneStrUnset  ST_AnimOneStr = 0
	ST_AnimOneStrNone   ST_AnimOneStr = 1
	ST_AnimOneStrOne    ST_AnimOneStr = 2
	ST_AnimOneStrBranch ST_AnimOneStr = 3
)

type CT_NumericRule struct {
	ValAttr     *float64
	FactAttr    *float64
	MaxAttr     *float64
	ExtLst      *_bg.CT_OfficeArtExtensionList
	TypeAttr    ST_ConstraintType
	ForAttr     ST_ConstraintRelationship
	ForNameAttr *string
	PtTypeAttr  ST_ElementType
}

func (_cega *ST_RotationPath) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_cega = 0
	case "\u006e\u006f\u006e\u0065":
		*_cega = 1
	case "\u0061l\u006f\u006e\u0067\u0050\u0061\u0074h":
		*_cega = 2
	}
	return nil
}

func NewCT_PresentationOf() *CT_PresentationOf { _fgc := &CT_PresentationOf{}; return _fgc }

func (_eagea ST_AnimOneStr) ValidateWithPath(path string) error {
	switch _eagea {
	case 0, 1, 2, 3:
	default:
		return _f.Errorf("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d", path, int(_eagea))
	}
	return nil
}

func (_aeaae ST_VerticalAlignment) MarshalXMLAttr(name _a.Name) (_a.Attr, error) {
	_cegd := _a.Attr{}
	_cegd.Name = name
	switch _aeaae {
	case ST_VerticalAlignmentUnset:
		_cegd.Value = ""
	case ST_VerticalAlignmentT:
		_cegd.Value = "\u0074"
	case ST_VerticalAlignmentMid:
		_cegd.Value = "\u006d\u0069\u0064"
	case ST_VerticalAlignmentB:
		_cegd.Value = "\u0062"
	case ST_VerticalAlignmentNone:
		_cegd.Value = "\u006e\u006f\u006e\u0065"
	}
	return _cegd, nil
}

type ST_AnimLvlStr byte

func NewCT_Cxn() *CT_Cxn { _cfeb := &CT_Cxn{}; return _cfeb }

// ValidateWithPath validates the CT_ElemPropSet and its children, prefixing error messages with path
func (_cdfc *CT_ElemPropSet) ValidateWithPath(path string) error {
	if _cdfc.PresAssocIDAttr != nil {
		if _cdda := _cdfc.PresAssocIDAttr.ValidateWithPath(path + "\u002f\u0050r\u0065\u0073\u0041s\u0073\u006f\u0063\u0049\u0044\u0041\u0074\u0074\u0072"); _cdda != nil {
			return _cdda
		}
	}
	if _cdfc.CustScaleXAttr != nil {
		if _abbe := _cdfc.CustScaleXAttr.ValidateWithPath(path + "\u002fC\u0075s\u0074\u0053\u0063\u0061\u006c\u0065\u0058\u0041\u0074\u0074\u0072"); _abbe != nil {
			return _abbe
		}
	}
	if _cdfc.CustScaleYAttr != nil {
		if _fbg := _cdfc.CustScaleYAttr.ValidateWithPath(path + "\u002fC\u0075s\u0074\u0053\u0063\u0061\u006c\u0065\u0059\u0041\u0074\u0074\u0072"); _fbg != nil {
			return _fbg
		}
	}
	if _cdfc.CustLinFactXAttr != nil {
		if _cfef := _cdfc.CustLinFactXAttr.ValidateWithPath(path + "\u002f\u0043\u0075\u0073\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074X\u0041\u0074\u0074\u0072"); _cfef != nil {
			return _cfef
		}
	}
	if _cdfc.CustLinFactYAttr != nil {
		if _ggec := _cdfc.CustLinFactYAttr.ValidateWithPath(path + "\u002f\u0043\u0075\u0073\u0074\u004c\u0069\u006e\u0046\u0061\u0063\u0074Y\u0041\u0074\u0074\u0072"); _ggec != nil {
			return _ggec
		}
	}
	if _cdfc.CustLinFactNeighborXAttr != nil {
		if _bgaeb := _cdfc.CustLinFactNeighborXAttr.ValidateWithPath(path + "\u002fC\u0075\u0073\u0074\u004ci\u006e\u0046\u0061\u0063\u0074N\u0065i\u0067h\u0062\u006f\u0072\u0058\u0041\u0074\u0074r"); _bgaeb != nil {
			return _bgaeb
		}
	}
	if _cdfc.CustLinFactNeighborYAttr != nil {
		if _bgad := _cdfc.CustLinFactNeighborYAttr.ValidateWithPath(path + "\u002fC\u0075\u0073\u0074\u004ci\u006e\u0046\u0061\u0063\u0074N\u0065i\u0067h\u0062\u006f\u0072\u0059\u0041\u0074\u0074r"); _bgad != nil {
			return _bgad
		}
	}
	if _cdfc.CustRadScaleRadAttr != nil {
		if _gac := _cdfc.CustRadScaleRadAttr.ValidateWithPath(path + "/\u0043u\u0073\u0074\u0052\u0061\u0064\u0053\u0063\u0061l\u0065\u0052\u0061\u0064At\u0074\u0072"); _gac != nil {
			return _gac
		}
	}
	if _cdfc.CustRadScaleIncAttr != nil {
		if _ebca := _cdfc.CustRadScaleIncAttr.ValidateWithPath(path + "/\u0043u\u0073\u0074\u0052\u0061\u0064\u0053\u0063\u0061l\u0065\u0049\u006e\u0063At\u0074\u0072"); _ebca != nil {
			return _ebca
		}
	}
	if _cdfc.PresLayoutVars != nil {
		if _dffb := _cdfc.PresLayoutVars.ValidateWithPath(path + "\u002fP\u0072e\u0073\u004c\u0061\u0079\u006f\u0075\u0074\u0056\u0061\u0072\u0073"); _dffb != nil {
			return _dffb
		}
	}
	if _cdfc.Style != nil {
		if _babb := _cdfc.Style.ValidateWithPath(path + "\u002f\u0053\u0074\u0079\u006c\u0065"); _babb != nil {
			return _babb
		}
	}
	return nil
}

func (_dbgg *ST_RotationPath) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
	_ccbg, _eadbg := d.Token()
	if _eadbg != nil {
		return _eadbg
	}
	if _debcg, _ggde := _ccbg.(_a.EndElement); _ggde && _debcg.Name == start.Name {
		*_dbgg = 1
		return nil
	}
	if _eadcg, _aadda := _ccbg.(_a.CharData); !_aadda {
		return _f.Errorf("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054", _ccbg)
	} else {
		switch string(_eadcg) {
		case "":
			*_dbgg = 0
		case "\u006e\u006f\u006e\u0065":
			*_dbgg = 1
		case "\u0061l\u006f\u006e\u0067\u0050\u0061\u0074h":
			*_dbgg = 2
		}
	}
	_ccbg, _eadbg = d.Token()
	if _eadbg != nil {
		return _eadbg
	}
	if _fdeea, _dcaee := _ccbg.(_a.EndElement); _dcaee && _fdeea.Name == start.Name {
		return nil
	}
	return _f.Errorf("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076", _ccbg)
}

func (_acaaf ST_LayoutShapeType) String() string {
	if _acaaf.ST_ShapeType != _bg.ST_ShapeTypeUnset {
		return _acaaf.ST_ShapeType.String()
	}
	if _acaaf.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		return _acaaf.ST_OutputShapeType.String()
	}
	return ""
}

func (_efec *ST_PrSetCustVal) ValidateWithPath(path string) error {
	_acbbc := []string{}
	if _efec.ST_Percentage != nil {
		_acbbc = append(_acbbc, "\u0053\u0054\u005f\u0050\u0065\u0072\u0063\u0065\u006e\u0074\u0061\u0067\u0065")
	}
	if _efec.Int32 != nil {
		_acbbc = append(_acbbc, "\u0049\u006e\u00743\u0032")
	}
	if len(_acbbc) > 1 {
		return _f.Errorf("%\u0073\u0020\u0074\u006f\u006f\u0020m\u0061\u006e\u0079\u0020\u006d\u0065\u006d\u0062\u0065r\u0073\u0020\u0073e\u0074:\u0020\u0025\u0076", path, _acbbc)
	}
	return nil
}

func (_gae *CT_ColorTransformHeaderLst) UnmarshalXML(d *_a.Decoder, start _a.StartElement) error {
_cfcf:
	for {
		_ebgf, _bgea := d.Token()
		if _bgea != nil {
			return _bgea
		}
		switch _adce := _ebgf.(type) {
		case _a.StartElement:
			switch _adce.Name {
			case _a.Name{Space: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", Local: "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072"}:
				_cdae := NewCT_ColorTransformHeader()
				if _gdb := d.DecodeElement(_cdae, &_adce); _gdb != nil {
					return _gdb
				}
				_gae.ColorsDefHdr = append(_gae.ColorsDefHdr, _cdae)
			default:
				_b.Log.Debug("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020o\u006e\u0020\u0043\u0054_\u0043\u006fl\u006f\u0072\u0054\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065\u0072\u004c\u0073\u0074\u0020\u0025\u0076", _adce.Name)
				if _bbf := d.Skip(); _bbf != nil {
					return _bbf
				}
			}
		case _a.EndElement:
			break _cfcf
		case _a.CharData:
		}
	}
	return nil
}

func (_abgdb ST_ConnectorDimension) String() string {
	switch _abgdb {
	case 0:
		return ""
	case 1:
		return "\u0031\u0044"
	case 2:
		return "\u0032\u0044"
	case 3:
		return "\u0063\u0075\u0073\u0074"
	}
	return ""
}

func (_edba *ST_TextDirection) UnmarshalXMLAttr(attr _a.Attr) error {
	switch attr.Value {
	case "":
		*_edba = 0
	case "\u0066\u0072\u006fm\u0054":
		*_edba = 1
	case "\u0066\u0072\u006fm\u0042":
		*_edba = 2
	}
	return nil
}

// ST_PrSetCustVal is a union type
type ST_PrSetCustVal struct {
	ST_Percentage *string
	Int32         *int32
}

func init() {
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0043\u0054\u004e\u0061\u006de", NewCT_CTName)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0043\u0054\u0044e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e", NewCT_CTDescription)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u0054\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079", NewCT_CTCategory)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fC\u0054\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073", NewCT_CTCategories)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0043\u006f\u006c\u006f\u0072s", NewCT_Colors)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fC\u0054\u0053\u0074\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c", NewCT_CTStyleLabel)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061\u006es\u0066\u006f\u0072\u006d", NewCT_ColorTransform)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fCo\u006c\u006f\u0072\u0054\u0072\u0061\u006e\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065\u0072", NewCT_ColorTransformHeader)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u006f\u006c\u006f\u0072\u0054\u0072\u0061n\u0073\u0066\u006f\u0072\u006d\u0048\u0065\u0061\u0064\u0065r\u004c\u0073\u0074", NewCT_ColorTransformHeaderLst)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fP\u0074", NewCT_Pt)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0050\u0074\u004c\u0069\u0073t", NewCT_PtList)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u0078\u006e", NewCT_Cxn)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u0078\u006e\u004c\u0069\u0073\u0074", NewCT_CxnList)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fD\u0061\u0074\u0061\u004d\u006f\u0064\u0065\u006c", NewCT_DataModel)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074", NewCT_Constraint)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u006f\u006e\u0073\u0074\u0072a\u0069\u006e\u0074\u0073", NewCT_Constraints)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u004e\u0075\u006d\u0065\u0072\u0069c\u0052\u0075\u006c\u0065", NewCT_NumericRule)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0052\u0075\u006c\u0065\u0073", NewCT_Rules)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0050\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u004f\u0066", NewCT_PresentationOf)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0041\u0064\u006a", NewCT_Adj)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0041\u0064\u006a\u004c\u0073t", NewCT_AdjLst)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0053\u0068\u0061\u0070\u0065", NewCT_Shape)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fP\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072", NewCT_Parameter)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fA\u006c\u0067\u006f\u0072\u0069\u0074\u0068\u006d", NewCT_Algorithm)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u004c\u0061\u0079\u006f\u0075\u0074\u004e\u006f\u0064\u0065", NewCT_LayoutNode)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0046\u006f\u0072\u0045\u0061\u0063\u0068", NewCT_ForEach)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0057\u0068\u0065\u006e", NewCT_When)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fO\u0074\u0068\u0065\u0072\u0077\u0069\u0073\u0065", NewCT_Otherwise)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0043\u0068\u006f\u006f\u0073e", NewCT_Choose)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0053\u0061\u006d\u0070\u006c\u0065\u0044\u0061\u0074\u0061", NewCT_SampleData)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079", NewCT_Category)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073", NewCT_Categories)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u004e\u0061\u006d\u0065", NewCT_Name)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0044\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e", NewCT_Description)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044e\u0066\u0069\u006e\u0069ti\u006f\u006e", NewCT_DiagramDefinition)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044e\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065a\u0064\u0065\u0072", NewCT_DiagramDefinitionHeader)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0044\u0069\u0061\u0067\u0072\u0061\u006d\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065rL\u0073\u0074", NewCT_DiagramDefinitionHeaderLst)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0052\u0065\u006c\u0049\u0064s", NewCT_RelIds)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0045\u006c\u0065\u006d\u0050\u0072o\u0070\u0053\u0065\u0074", NewCT_ElemPropSet)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054\u005f\u004f\u0072\u0067\u0043\u0068\u0061\u0072\u0074", NewCT_OrgChart)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054\u005f\u0043\u0068\u0069\u006c\u0064\u004d\u0061\u0078", NewCT_ChildMax)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fC\u0068\u0069\u006c\u0064\u0050\u0072\u0065\u0066", NewCT_ChildPref)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0042\u0075\u006cl\u0065\u0074\u0045\u006e\u0061\u0062\u006c\u0065\u0064", NewCT_BulletEnabled)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fD\u0069\u0072\u0065\u0063\u0074\u0069\u006f\u006e", NewCT_Direction)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fH\u0069\u0065\u0072\u0042\u0072a\u006e\u0063h\u0053\u0074\u0079\u006c\u0065", NewCT_HierBranchStyle)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004f\u006e\u0065", NewCT_AnimOne)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0041\u006e\u0069\u006d\u004c\u0076\u006c", NewCT_AnimLvl)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0052\u0065\u0073i\u007a\u0065\u0048\u0061\u006e\u0064\u006c\u0065\u0073", NewCT_ResizeHandles)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u004ca\u0079\u006f\u0075\u0074\u0056\u0061\u0072\u0069\u0061\u0062l\u0065P\u0072\u006f\u0070\u0065\u0072\u0074\u0079S\u0065\u0074", NewCT_LayoutVariablePropertySet)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0053\u0044\u004e\u0061\u006de", NewCT_SDName)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054_\u0053\u0044\u0044e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e", NewCT_SDDescription)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0053\u0044\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0079", NewCT_SDCategory)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fS\u0044\u0043\u0061\u0074\u0065\u0067\u006f\u0072\u0069\u0065\u0073", NewCT_SDCategories)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005fT\u0065\u0078\u0074\u0050\u0072\u006f\u0070\u0073", NewCT_TextProps)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043\u0054\u005f\u0053\u0074\u0079\u006c\u0065\u004c\u0061\u0062\u0065\u006c", NewCT_StyleLabel)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005fS\u0074\u0079\u006c\u0065\u0044e\u0066\u0069n\u0069\u0074\u0069\u006f\u006e", NewCT_StyleDefinition)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0043T\u005f\u0053\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0069\u006ei\u0074\u0069\u006f\u006e\u0048\u0065\u0061\u0064\u0065\u0072", NewCT_StyleDefinitionHeader)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "C\u0054\u005f\u0053\u0074\u0079\u006ce\u0044\u0065\u0066\u0069\u006e\u0069\u0074\u0069\u006fn\u0048\u0065\u0061d\u0065r\u004c\u0073\u0074", NewCT_StyleDefinitionHeaderLst)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0063o\u006c\u006f\u0072\u0073\u0044\u0065f", NewColorsDef)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0063\u006f\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072", NewColorsDefHdr)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0063o\u006co\u0072\u0073\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074", NewColorsDefHdrLst)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0064a\u0074\u0061\u004d\u006f\u0064\u0065l", NewDataModel)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u006ca\u0079\u006f\u0075\u0074\u0044\u0065f", NewLayoutDef)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u006c\u0061\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072", NewLayoutDefHdr)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u006ca\u0079o\u0075\u0074\u0044\u0065\u0066\u0048\u0064\u0072\u004c\u0073\u0074", NewLayoutDefHdrLst)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0072\u0065\u006c\u0049\u0064\u0073", NewRelIds)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0073\u0074\u0079\u006c\u0065\u0044\u0065\u0066", NewStyleDef)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "s\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048\u0064\u0072", NewStyleDefHdr)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0073\u0074\u0079\u006c\u0065\u0044\u0065\u0066\u0048d\u0072\u004c\u0073\u0074", NewStyleDefHdrLst)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "A\u0047\u005f\u0049\u0074er\u0061t\u006f\u0072\u0041\u0074\u0074r\u0069\u0062\u0075\u0074\u0065\u0073", NewAG_IteratorAttributes)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0041\u0047\u005fCo\u006e\u0073\u0074\u0072\u0061\u0069\u006e\u0074\u0041\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0073", NewAG_ConstraintAttributes)
	_fd.RegisterConstructor("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0064\u0069\u0061\u0067\u0072\u0061\u006d", "\u0041\u0047\u005f\u0043\u006f\u006e\u0073\u0074\u0072\u0061\u0069n\u0074\u0052\u0065\u0066\u0041\u0074\u0074\u0072\u0069\u0062u\u0074\u0065\u0073", NewAG_ConstraintRefAttributes)
}

func (_feddg ST_ChildOrderType) Validate() error { return _feddg.ValidateWithPath("") }

func (_aggea ST_FlowDirection) Validate() error { return _aggea.ValidateWithPath("") }

func NewColorsDefHdr() *ColorsDefHdr {
	_gaac := &ColorsDefHdr{}
	_gaac.CT_ColorTransformHeader = *NewCT_ColorTransformHeader()
	return _gaac
}

func ParseUnionST_LayoutShapeType(s string) (ST_LayoutShapeType, error) {
	return ST_LayoutShapeType{}, nil
}

const (
	ST_NodeHorizontalAlignmentUnset ST_NodeHorizontalAlignment = 0
	ST_NodeHorizontalAlignmentL     ST_NodeHorizontalAlignment = 1
	ST_NodeHorizontalAlignmentCtr   ST_NodeHorizontalAlignment = 2
	ST_NodeHorizontalAlignmentR     ST_NodeHorizontalAlignment = 3
)
