package presentation

import (
	_ea "archive/zip"
	_db "bytes"
	_gb "encoding/xml"
	_bc "errors"
	_fe "fmt"
	_gbg "github.com/arcpd/msword"
	_ec "github.com/arcpd/msword/common"
	_e "github.com/arcpd/msword/common/logger"
	_fgb "github.com/arcpd/msword/common/tempstorage"
	_gg "github.com/arcpd/msword/drawing"
	_fcf "github.com/arcpd/msword/internal/formatutils"
	_cb "github.com/arcpd/msword/measurement"
	_eef "github.com/arcpd/msword/schema/soo/dml"
	_f "github.com/arcpd/msword/schema/soo/dml/chart"
	_fc "github.com/arcpd/msword/schema/soo/ofc/sharedTypes"
	_fg "github.com/arcpd/msword/schema/soo/pkg/relationships"
	_a "github.com/arcpd/msword/schema/soo/pml"
	_cd "github.com/arcpd/msword/zippkg"
	_fd "image"
	_b "image/jpeg"
	_cdg "io"
	_eg "math"
	_gbc "os"
	_ee "path"
	_d "sort"
	_fa "strconv"
	_g "strings"
)

func _dfb() *Presentation {
	_gef := &Presentation{_dac: _a.NewPresentation()}
	_gef._dac.SldIdLst = _a.NewCT_SlideIdList()
	_gef._dac.ConformanceAttr = _fc.ST_ConformanceClassTransitional
	_gef.AppProperties = _ec.NewAppProperties()
	_gef.CoreProperties = _ec.NewCoreProperties()
	_gef._fead = _ec.NewTableStyles()
	_gef.ContentTypes = _ec.NewContentTypes()
	_gef.Rels = _ec.NewRelationships()
	_gef._bfe = _ec.NewRelationships()
	_gef._bfc = NewPresentationProperties()
	_gef._egb = NewViewProperties()
	_gef._bdeb = map[string]string{}
	return _gef
}

// SlideText is an array of extracted text items which has some methods for representing extracted text from a slide.
type SlideText struct{ Items []*TextItem }

// Themes returns an array of presentation themes.
func (_fddbe *Presentation) Themes() []*_eef.Theme { return _fddbe._eade }

// AddSlideWithLayout adds a new slide with content copied from a layout.  Normally you should
// use AddDefaultSlideWithLayout as it will do some post processing similar to PowerPoint to
// clear place holder text, etc.
func (_ade *Presentation) AddSlideWithLayout(l SlideLayout) (Slide, error) {
	_gdgb := _a.NewCT_SlideIdListEntry()
	_gdgb.IdAttr = 256
	for _, _efeb := range _ade._dac.SldIdLst.SldId {
		if _efeb.IdAttr >= _gdgb.IdAttr {
			_gdgb.IdAttr = _efeb.IdAttr + 1
		}
	}
	_ade._dac.SldIdLst.SldId = append(_ade._dac.SldIdLst.SldId, _gdgb)
	_cab := _a.NewSld()
	_gbcd := _db.Buffer{}
	_adf := _gb.NewEncoder(&_gbcd)
	_fca := _gb.StartElement{Name: _gb.Name{Local: "\u0073\u006c\u0069d\u0065"}}
	_fca.Attr = append(_fca.Attr, _gb.Attr{Name: _gb.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069o\u006e\u006d\u006c\u002f\u0032\u00300\u0036\u002f\u006da\u0069\u006e"})
	_fca.Attr = append(_fca.Attr, _gb.Attr{Name: _gb.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0061"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065m\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006cf\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067m\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"})
	_fca.Attr = append(_fca.Attr, _gb.Attr{Name: _gb.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0070"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078m\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002eo\u0072\u0067\u002f\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069o\u006e\u006d\u006c\u002f\u0032\u00300\u0036\u002f\u006da\u0069\u006e"})
	_fca.Attr = append(_fca.Attr, _gb.Attr{Name: _gb.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0072"}, Value: "\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"})
	_fca.Attr = append(_fca.Attr, _gb.Attr{Name: _gb.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0073\u0068"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0073\u0063\u0068\u0065m\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0073\u0068\u0061\u0072e\u0064\u0054\u0079\u0070\u0065\u0073"})
	_fca.Attr = append(_fca.Attr, _gb.Attr{Name: _gb.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	if _acf := l._dcbe.CSld.MarshalXML(_adf, _fca); _acf != nil {
		return Slide{}, _acf
	}
	_adf.Flush()
	_fbe := _gb.NewDecoder(&_gbcd)
	_cab.CSld = _a.NewCT_CommonSlideData()
	if _gcdc := _fbe.Decode(_cab.CSld); _gcdc != nil {
		return Slide{}, _gcdc
	}
	_cab.CSld.NameAttr = nil
	_cab.CSld.SpTree.Choice = _dba(_cab.CSld.SpTree.Choice)
	_ade._eda = append(_ade._eda, _cab)
	_bbb := _ade._bfe.AddAutoRelationship(_gbg.DocTypePresentation, _gbg.OfficeDocumentType, len(_ade._eda), _gbg.SlideType)
	_gdgb.RIdAttr = _bbb.ID()
	_ccc := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.SlideType, len(_ade._eda))
	_ade.ContentTypes.AddOverride(_ccc, _gbg.SlideContentType)
	_ecea := _ec.NewRelationships()
	_ade._gfd = append(_ade._gfd, _ecea)
	_ebac := len(_ade._gfd) - 1
	for _bfda, _ccgg := range _ade._gff {
		if _ccgg == l.X() {
			_cbb := _ade._dfa[_bfda]
			for _, _adbc := range _cbb.X().Relationship {
				if _adbc.TypeAttr != _gbg.SlideMasterType {
					_ade._gfd[_ebac].X().Relationship = append(_ade._gfd[_ebac].X().Relationship, _adbc)
				}
			}
			_ecea.AddAutoRelationship(_gbg.DocTypePresentation, _gbg.SlideType, _bfda+1, _gbg.SlideLayoutType)
		}
	}
	_fff := Slide{_gdgb, _cab, _ade, nil}
	return _fff, nil
}

// NewSlideScreenSize returns slide screen size with default MS PowerPoint slide screen size 16x9.
func NewSlideScreenSize() SlideScreenSize {
	return NewSlideScreenSizeWithValue(SlideScreenSize16x9[0], SlideScreenSize16x9[1])
}

// TextItem is used for keeping text with references to a paragraph and run, a shape or a table, a row and a cell where it is located.
type TextItem struct {
	Text         string
	Presentation *Presentation
	Shape        *_a.CT_Shape
	GraphicFrame *_a.CT_GraphicalObjectFrame
	Paragraph    *_eef.CT_TextParagraph
	Run          *_eef.CT_RegularTextRun
	TableInfo    *TableInfo
	_dg          []rectangle
	_ca          int
	_ecf         int
}

// Height returns slide screen size height in EMU units.
func (_cdfbg *SlideScreenSize) Height() int32 { return _cdfbg[1] }

// AddDefaultSlideWithLayout tries to replicate what PowerPoint does when
// inserting a slide with a new style by clearing placeholder content and removing
// some placeholders.  Use AddSlideWithLayout if you need more control.
func (_fbb *Presentation) AddDefaultSlideWithLayout(l SlideLayout) (Slide, error) {
	_dbd, _gbga := _fbb.AddSlideWithLayout(l)
	for _, _ddac := range _dbd.PlaceHolders() {
		_ddac.Clear()
		switch _ddac.Type() {
		case _a.ST_PlaceholderTypeFtr, _a.ST_PlaceholderTypeDt, _a.ST_PlaceholderTypeSldNum:
			_ddac.Remove()
		}
	}
	return _dbd, _gbga
}
func _egg(_gde *Presentation, _eed *_a.CT_Shape, _feb *_a.CT_GraphicalObjectFrame, _facc *TableInfo, _ac *_eef.CT_Transform2D, _eedb int, _dfde []rectangle, _gdge []*_eef.CT_TextParagraph) []*TextItem {
	_cg := []*TextItem{}
	var _fgbg, _gag, _ddg, _gbgf, _ccg, _ffc int64
	_fgbb := _ac == nil
	_ae := 0
	for _, _gcde := range _gdge {
		for _, _efe := range _gcde.EG_TextRun {
			if _bfb := _efe.R; _bfb != nil {
				if !_fgbb {
					if _ac.Off != nil {
						if _eedf := _ac.Ext; _eedf != nil {
							_ccg, _ffc = _eedf.CxAttr, _eedf.CyAttr
						}
						if _cba := _ac.Off.XAttr.ST_CoordinateUnqualified; _cba != nil {
							_fgbg = *_cba
							_gag = _fgbg + _ccg
							_fgbb = true
						}
						if _adaf := _ac.Off.YAttr.ST_CoordinateUnqualified; _adaf != nil {
							_ddg = *_adaf
							_gbgf = _ddg + _ffc
							_fgbb = true
						}
					}
				}
				_gca := append([]rectangle{}, _dfde...)
				_gca = append(_gca, rectangle{_cfc: _fgbg, _cdga: _gag, _dgb: _ddg, _cbg: _gbgf})
				_cg = append(_cg, &TextItem{Presentation: _gde, Shape: _eed, GraphicFrame: _feb, TableInfo: _facc, Paragraph: _gcde, Run: _bfb, Text: _bfb.T, _dg: _gca, _ca: _eedb, _ecf: _ae})
				_ae++
			}
		}
	}
	return _cg
}

// NewViewProperties constructs a new ViewProperties.
func NewViewProperties() ViewProperties { return ViewProperties{_ebacf: _a.NewViewPr()} }

// SlideLayouts returns a slice of all layouts in SlideMaster.
func (_bbda SlideMaster) SlideLayouts() []SlideLayout {
	_ecae := map[string]int{}
	_agc := []SlideLayout{}
	for _, _abab := range _bbda._gcba.Relationships() {
		_aeac := _g.Replace(_abab.Target(), ".\u002e\u002f\u0073\u006c\u0069\u0064e\u004c\u0061\u0079\u006f\u0075\u0074\u0073\u002f\u0073l\u0069\u0064\u0065L\u0061y\u006f\u0075\u0074", "", -1)
		_aeac = _g.Replace(_aeac, "\u002e\u0078\u006d\u006c", "", -1)
		if _bcdgf, _cfcbd := _fa.ParseInt(_aeac, 10, 32); _cfcbd == nil {
			_ecae[_abab.ID()] = int(_bcdgf)
		}
	}
	for _, _adafaf := range _bbda._fec.SldLayoutIdLst.SldLayoutId {
		if _dece, _feae := _ecae[_adafaf.RIdAttr]; _feae {
			_cefb := _bbda._dfab._gff[_dece-1]
			_agc = append(_agc, SlideLayout{_cefb})
		}
	}
	return _agc
}

// GetPlaceholder returns a placeholder given its type.  If there are multiplace
// placeholders of the same type, this method returns the first one.  You must use the
// PlaceHolders() method to access the others.
func (_dcb Slide) GetPlaceholder(t _a.ST_PlaceholderType) (PlaceHolder, error) {
	for _, _bffg := range _dcb._gfgd.CSld.SpTree.Choice {
		for _, _gaac := range _bffg.Sp {
			if _gaac.NvSpPr != nil && _gaac.NvSpPr.NvPr != nil && _gaac.NvSpPr.NvPr.Ph != nil {
				if _gaac.NvSpPr.NvPr.Ph.TypeAttr == t {
					return PlaceHolder{_gaac, _dcb._gfgd}, nil
				}
			}
		}
	}
	return PlaceHolder{}, _bc.New("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066i\u006e\u0064\u0020\u0070\u006c\u0061\u0063\u0065\u0068\u006fl\u0064\u0065\u0072")
}

// SetSize sets the slide size, take argument of SlideScreenSize.
func (_fcae *SlideSize) SetSize(sz SlideScreenSize) {
	_fcae._deg.CxAttr = sz[0]
	_fcae._deg.CyAttr = sz[1]
}

// Close closes the presentation, removing any temporary files that might have been
// created when opening a document.
func (_fdg *Presentation) Close() error {
	if _fdg.TmpPath != "" {
		return _fgb.RemoveAll(_fdg.TmpPath)
	}
	return nil
}

// RemoveSlide removes a slide from a presentation.
func (_fgbf *Presentation) RemoveSlide(s Slide) error {
	_egee := false
	_gcaf := 0
	for _eeg, _bafc := range _fgbf._eda {
		if _bafc == s._gfgd {
			if _fgbf._dac.SldIdLst.SldId[_eeg] != s._gddb {
				return _bc.New("i\u006e\u0063\u006f\u006e\u0073\u0069s\u0074\u0065\u006e\u0063\u0079\u0020i\u006e\u0020\u0073\u006c\u0069\u0064\u0065s\u0020\u0061\u006e\u0064\u0020\u0049\u0044\u0020\u006c\u0069s\u0074")
			}
			copy(_fgbf._eda[_eeg:], _fgbf._eda[_eeg+1:])
			_fgbf._eda = _fgbf._eda[0 : len(_fgbf._eda)-1]
			copy(_fgbf._gfd[_eeg:], _fgbf._gfd[_eeg+1:])
			_fgbf._gfd = _fgbf._gfd[0 : len(_fgbf._gfd)-1]
			copy(_fgbf._dac.SldIdLst.SldId[_eeg:], _fgbf._dac.SldIdLst.SldId[_eeg+1:])
			_fgbf._dac.SldIdLst.SldId = _fgbf._dac.SldIdLst.SldId[0 : len(_fgbf._dac.SldIdLst.SldId)-1]
			_egee = true
			_gcaf = _eeg
		}
	}
	if !_egee {
		return _bc.New("u\u006ea\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066i\u006e\u0064\u0020\u0073li\u0064\u0065")
	}
	_cbe := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.SlideType, 0)
	return _fgbf.ContentTypes.RemoveOverrideByIndex(_cbe, _gcaf)
}

// Presentation returns a slide's presentation.
func (_bea Slide) Presentation() *Presentation { return _bea._gac }

// X returns the inner wrapped XML type.
func (_fefa *SlideSize) X() *_a.CT_SlideSize { return _fefa._deg }

type sort2d []*TextItem
type rectangle struct {
	_cfc  int64
	_dgb  int64
	_cdga int64
	_cbg  int64
}

// SlideScreenSize represents the slide screen size as a 2 element array
// representing the width and height in EMU units.
type SlideScreenSize [2]int32

// X returns the inner wrapped XML type.
func (_caa *Presentation) X() *_a.Presentation { return _caa._dac }

// PlaceHolders returns all of the content place holders within a given slide.
func (_gfb Slide) PlaceHolders() []PlaceHolder {
	_bfgf := []PlaceHolder{}
	for _, _fceb := range _gfb._gfgd.CSld.SpTree.Choice {
		for _, _dgac := range _fceb.Sp {
			if _dgac.NvSpPr != nil && _dgac.NvSpPr.NvPr != nil && _dgac.NvSpPr.NvPr.Ph != nil {
				_bfgf = append(_bfgf, PlaceHolder{_dgac, _gfb._gfgd})
			}
		}
	}
	return _bfgf
}

// NormalViewPr returns the NormalViewPr property.
func (_debf ViewProperties) NormalViewPr() *_a.CT_NormalViewProperties {
	return _debf._ebacf.NormalViewPr
}

// TextBox is a text box within a slide.
type TextBox struct{ _bbcd *_a.CT_Shape }

// AddTable adds an empty table to a slide.
func (_dca Slide) AddTable() *_ec.Table {
	_agdf := _a.NewCT_GroupShapeChoice()
	_dca._gfgd.CSld.SpTree.Choice = append(_dca._gfgd.CSld.SpTree.Choice, _agdf)
	_eabc := _a.NewCT_GraphicalObjectFrame()
	_agdf.GraphicFrame = append(_agdf.GraphicFrame, _eabc)
	_eabc.Xfrm.Off = _eef.NewCT_Point2D()
	_bad := int64(1)
	_eabc.Xfrm.Off.XAttr = _eef.ST_Coordinate{ST_CoordinateUnqualified: &_bad}
	_eabc.Xfrm.Off.YAttr = _eef.ST_Coordinate{ST_CoordinateUnqualified: &_bad}
	_acfc := _eabc.Graphic.CT_GraphicalObject.GraphicData
	_acfc.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002eo\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073.\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006dl/\u0032\u0030\u0030\u0036\u002f\u0074\u0061\u0062\u006c\u0065"
	_eff := _ec.NewTableWithXfrm(_eabc.Xfrm)
	_acfc.Any = append(_acfc.Any, _eff.X())
	return _eff
}

// AddImage adds an image textbox to a slide.
func (_gfda Slide) AddImage(img _ec.ImageRef) Image {
	_fgaf := _a.NewCT_GroupShapeChoice()
	_gfda._gfgd.CSld.SpTree.Choice = append(_gfda._gfgd.CSld.SpTree.Choice, _fgaf)
	_fgcgf := _a.NewCT_Picture()
	_fgaf.Pic = append(_fgaf.Pic, _fgcgf)
	_fgcgf.NvPicPr.CNvPicPr = _eef.NewCT_NonVisualPictureProperties()
	_fgcgf.NvPicPr.CNvPicPr.PicLocks = _eef.NewCT_PictureLocking()
	_fgcgf.NvPicPr.CNvPicPr.PicLocks.NoChangeAspectAttr = _gbg.Bool(true)
	_fgcgf.BlipFill = _eef.NewCT_BlipFillProperties()
	_fgcgf.BlipFill.Blip = _eef.NewCT_Blip()
	_dbae := _gfda.AddImageToRels(img)
	_fgcgf.BlipFill.Blip.EmbedAttr = _gbg.String(_dbae)
	_fgcgf.BlipFill.Stretch = _eef.NewCT_StretchInfoProperties()
	_fgcgf.BlipFill.Stretch.FillRect = _eef.NewCT_RelativeRect()
	_fgcgf.SpPr = _eef.NewCT_ShapeProperties()
	_fgcgf.SpPr.PrstGeom = _eef.NewCT_PresetGeometry2D()
	_fgcgf.SpPr.PrstGeom.PrstAttr = _eef.ST_ShapeTypeRect
	_cdbb := Image{_fgcgf}
	_fcee := img.Size()
	_cdbb.Properties().SetWidth(_cb.Distance(_fcee.X) * _cb.Pixel72)
	_cdbb.Properties().SetHeight(_cb.Distance(_fcee.Y) * _cb.Pixel72)
	_cdbb.Properties().SetPosition(0, 0)
	return _cdbb
}

// SlideMasters returns the slide masters defined in the presentation.
func (_cccf *Presentation) SlideMasters() []SlideMaster {
	_ggbc := []SlideMaster{}
	for _gaf, _dge := range _cccf._bda {
		_ggbc = append(_ggbc, SlideMaster{_cccf, _cccf._gdbf[_gaf], _dge})
	}
	return _ggbc
}
func (_baebe *Presentation) onNewRelationship(_baf *_cd.DecodeMap, _fcd, _dae string, _cdea []*_ea.File, _eeaa *_fg.Relationship, _ggbb _cd.Target) error {
	_cebc := _gbg.DocTypePresentation
	switch _dae {
	case _gbg.OfficeDocumentType:
		_baebe._dac = _a.NewPresentation()
		_baf.AddTarget(_fcd, _baebe._dac, _dae, 0)
		_baf.AddTarget(_cd.RelationsPathFor(_fcd), _baebe._bfe.X(), _dae, 0)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, 0)
	case _gbg.CorePropertiesType:
		_baf.AddTarget(_fcd, _baebe.CoreProperties.X(), _dae, 0)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, 0)
	case _gbg.CustomPropertiesType:
		_baf.AddTarget(_fcd, _baebe.CustomProperties.X(), _dae, 0)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, 0)
	case _gbg.PresentationPropertiesType:
		_baf.AddTarget(_fcd, _baebe._bfc.X(), _dae, 0)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, 0)
	case _gbg.ViewPropertiesType:
		_baf.AddTarget(_fcd, _baebe._egb.X(), _dae, 0)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, 0)
	case _gbg.TableStylesType:
		_baf.AddTarget(_fcd, _baebe._fead.X(), _dae, 0)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, 0)
	case _gbg.HyperLinkType:
		_cfeb := _eef.NewCT_Hyperlink()
		_ddc := uint32(len(_baebe._cef))
		_baf.AddTarget(_fcd, _cfeb, _dae, _ddc)
		_baebe._cef = append(_baebe._cef, _cfeb)
	case _gbg.CustomXMLType:
		_egba := &_gbg.XSDAny{}
		_deb := uint32(len(_baebe._cge))
		_baf.AddTarget(_fcd, _egba, _dae, _deb)
		_baebe._cge = append(_baebe._cge, _egba)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, len(_baebe._cge))
	case _gbg.ChartType:
		_cbdc := chart{_gd: _f.NewChartSpace()}
		_fffb := uint32(len(_baebe._bfg))
		_baf.AddTarget(_fcd, _cbdc._gd, _dae, _fffb)
		_baebe._bfg = append(_baebe._bfg, &_cbdc)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, len(_baebe._bfg))
		_cbdc._gc = _eeaa.TargetAttr
	case _gbg.HandoutMasterType:
		_aeae := _a.NewHandoutMaster()
		_gegb := uint32(len(_baebe._ffac))
		_baf.AddTarget(_fcd, _aeae, _dae, _gegb)
		_baebe._ffac = append(_baebe._ffac, _aeae)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, len(_baebe._ffac))
	case _gbg.NotesMasterType:
		_adg := _a.NewNotesMaster()
		_aaag := uint32(len(_baebe._fbd))
		_baf.AddTarget(_fcd, _adg, _dae, _aaag)
		_baebe._fbd = append(_baebe._fbd, _adg)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, len(_baebe._fbd))
	case _gbg.ExtendedPropertiesType:
		_baf.AddTarget(_fcd, _baebe.AppProperties.X(), _dae, 0)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, 0)
	case _gbg.SlideType:
		if _edc, _gfg := _fcf.StringToNumbers(_fcd); _gfg {
			if len(_baebe._eda) < _edc {
				_fgf := _a.NewSld()
				_baebe._eda = append(_baebe._eda, _fgf)
				_baf.AddTarget(_fcd, _fgf, _dae, uint32(_edc))
				_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, _edc)
				_ggc := _ec.NewRelationships()
				_baf.AddTarget(_cd.RelationsPathFor(_fcd), _ggc.X(), _dae, 0)
				if len(_baebe._gfd) >= _edc {
					_baebe._gfd[_edc-1] = _ggc
				} else {
					_baebe._gfd = append(_baebe._gfd, _ggc)
				}
			}
		}
	case _gbg.SlideMasterType:
		_dgg := _a.NewSldMaster()
		if !_baf.AddTarget(_fcd, _dgg, _dae, uint32(len(_baebe._bda)+1)) {
			return nil
		}
		_baebe._bda = append(_baebe._bda, _dgg)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, len(_baebe._bda))
		_abf := _ec.NewRelationships()
		_baf.AddTarget(_cd.RelationsPathFor(_fcd), _abf.X(), _dae, 0)
		_baebe._gdbf = append(_baebe._gdbf, _abf)
	case _gbg.SlideLayoutType:
		_fedb := _a.NewSldLayout()
		if !_baf.AddTarget(_fcd, _fedb, _dae, uint32(len(_baebe._gff)+1)) {
			return nil
		}
		_baebe._gff = append(_baebe._gff, _fedb)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, len(_baebe._gff))
		_aba := _ec.NewRelationships()
		_baf.AddTarget(_cd.RelationsPathFor(_fcd), _aba.X(), _dae, 0)
		_baebe._dfa = append(_baebe._dfa, _aba)
	case _gbg.ThumbnailType:
		for _bdf, _bff := range _cdea {
			if _bff == nil {
				continue
			}
			if _bff.Name == _fcd {
				_bcdf, _bdd := _bff.Open()
				if _bdd != nil {
					return _fe.Errorf("e\u0072\u0072\u006f\u0072\u0020\u0072e\u0061\u0064\u0069\u006e\u0067\u0020\u0074\u0068\u0075m\u0062\u006e\u0061i\u006c:\u0020\u0025\u0073", _bdd)
				}
				_baebe.Thumbnail, _, _bdd = _fd.Decode(_bcdf)
				_bcdf.Close()
				if _bdd != nil {
					return _fe.Errorf("\u0065\u0072\u0072\u006fr\u0020\u0064\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020t\u0068u\u006d\u0062\u006e\u0061\u0069\u006c\u003a \u0025\u0073", _bdd)
				}
				_cdea[_bdf] = nil
			}
		}
	case _gbg.ThemeType:
		_dfe := _eef.NewTheme()
		if !_baf.AddTarget(_fcd, _dfe, _dae, uint32(len(_baebe._eade)+1)) {
			return nil
		}
		_baebe._eade = append(_baebe._eade, _dfe)
		_eeaa.TargetAttr = _gbg.RelativeFilename(_cebc, _ggbb.Typ, _dae, len(_baebe._eade))
		_cgfc := _ec.NewRelationships()
		_baf.AddTarget(_cd.RelationsPathFor(_fcd), _cgfc.X(), _dae, 0)
		_baebe._cdf = append(_baebe._cdf, _cgfc)
	case _gbg.ImageType:
		_fcd = _ee.Clean(_fcd)
		if _edgb, _bgf := _baebe._bdeb[_fcd]; _bgf {
			_eeaa.TargetAttr = _edgb
			return nil
		}
		_bebe := ""
		for _eeff, _edca := range _cdea {
			if _edca == nil {
				continue
			}
			if _edca.Name == _fcd {
				_cfcg, _dcf := _cd.ExtractToDiskTmp(_edca, _baebe.TmpPath)
				if _dcf != nil {
					return _dcf
				}
				_eebf, _dcf := _ec.ImageFromStorage(_cfcg)
				if _dcf != nil {
					return _dcf
				}
				_bebe = _eebf.Format
				_edcd := _ec.MakeImageRef(_eebf, &_baebe.DocBase, _baebe._bfe)
				_edcd.SetTarget("\u002e\u002e\u002f" + _fcd[4:])
				_baebe.Images = append(_baebe.Images, _edcd)
				_cdea[_eeff] = nil
				_gccf := len(_baebe.Images)
				if _gcb, _cdgb := _fcf.StringToNumbers(_fcd); _cdgb {
					_gccf = _gcb
				}
				_baf.RecordIndex(_fcd, _gccf)
				break
			}
		}
		_fbdf := _baf.IndexFor(_fcd)
		_eeaa.TargetAttr = _gbg.RelativeImageFilename(_cebc, _ggbb.Typ, _dae, _fbdf, _bebe)
		_baebe._bdeb[_fcd] = _eeaa.TargetAttr
	default:
		_e.Log.Debug("\u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073\u0068\u0069\u0070\u0020\u0074\u0079\u0070\u0065\u003a\u0020\u0025\u0073\u0020\u0074\u0067\u0074\u003a\u0020\u0025\u0073", _dae, _fcd)
	}
	return nil
}

// WebPr returns the WebPr property.
func (_gagg PresentationProperties) WebPr() *_a.CT_WebProperties { return _gagg._ddad.WebPr }

// NotesTextViewPr returns the NotesTextViewPr property.
func (_aff ViewProperties) NotesTextViewPr() *_a.CT_NotesTextViewProperties {
	return _aff._ebacf.NotesTextViewPr
}

// AddTable adds a new table to a placeholder.
func (_ge PlaceHolder) AddTable() *_ec.Table {
	_ge.Clear()
	_gdb := _a.NewCT_GroupShapeChoice()
	_ge._cebf.CSld.SpTree.Choice = append(_ge._cebf.CSld.SpTree.Choice, _gdb)
	_cca := _a.NewCT_GraphicalObjectFrame()
	_gdb.GraphicFrame = append(_gdb.GraphicFrame, _cca)
	_cca.Xfrm.Off = _eef.NewCT_Point2D()
	_ceef := int64(1)
	_cca.Xfrm.Off.XAttr = _eef.ST_Coordinate{ST_CoordinateUnqualified: &_ceef}
	_cca.Xfrm.Off.YAttr = _eef.ST_Coordinate{ST_CoordinateUnqualified: &_ceef}
	_aca := _cca.Graphic.CT_GraphicalObject.GraphicData
	_aca.UriAttr = "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002eo\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073.\u006f\u0072\u0067\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006dl/\u0032\u0030\u0030\u0036\u002f\u0074\u0061\u0062\u006c\u0065"
	_bg := _ec.NewTableWithXfrm(_cca.Xfrm)
	_aca.Any = append(_aca.Any, _bg.X())
	return _bg
}

// OpenTemplate opens a template file.
func OpenTemplate(fn string) (*Presentation, error) {
	_gf, _fdd := Open(fn)
	if _fdd != nil {
		return nil, _fdd
	}
	return _gf, nil
}
func (_ggb sort2d) Len() int { return len(_ggb) }

// Properties returns the properties of the TextBox.
func (_eegf TextBox) Properties() _gg.ShapeProperties {
	if _eegf._bbcd.SpPr == nil {
		_eegf._bbcd.SpPr = _eef.NewCT_ShapeProperties()
	}
	return _gg.MakeShapeProperties(_eegf._bbcd.SpPr)
}

// NewSlideScreenSizeWithValue returns slide screen size with given width and height.
// Width and Height value is in EMU units, use our measurement.ToEMU to convert the -
// width and height value.
func NewSlideScreenSizeWithValue(width, height int32) SlideScreenSize {
	return SlideScreenSize{width, height}
}

// GetPlaceholderByIndex returns a placeholder given its index.  If there are multiplace
// placeholders of the same index, this method returns the first one.  You must use the
// PlaceHolders() method to access the others.
func (_adgf Slide) GetPlaceholderByIndex(idx uint32) (PlaceHolder, error) {
	for _, _daa := range _adgf._gfgd.CSld.SpTree.Choice {
		for _, _fgd := range _daa.Sp {
			if _fgd.NvSpPr != nil && _fgd.NvSpPr.NvPr != nil && _fgd.NvSpPr.NvPr.Ph != nil {
				if (idx == 0 && _fgd.NvSpPr.NvPr.Ph.IdxAttr == nil) || (_fgd.NvSpPr.NvPr.Ph.IdxAttr != nil && *_fgd.NvSpPr.NvPr.Ph.IdxAttr == idx) {
					return PlaceHolder{_fgd, _adgf._gfgd}, nil
				}
			}
		}
	}
	return PlaceHolder{}, _bc.New("\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066i\u006e\u0064\u0020\u0070\u006c\u0061\u0063\u0065\u0068\u006fl\u0064\u0065\u0072")
}

// Remove removes a placeholder from a presentation.
func (_cbf PlaceHolder) Remove() error {
	for _cgf, _gba := range _cbf._cebf.CSld.SpTree.Choice {
		for _, _abe := range _gba.Sp {
			if _abe == _cbf._beb {
				copy(_cbf._cebf.CSld.SpTree.Choice[_cgf:], _cbf._cebf.CSld.SpTree.Choice[_cgf+1:])
				_cbf._cebf.CSld.SpTree.Choice = _cbf._cebf.CSld.SpTree.Choice[0 : len(_cbf._cebf.CSld.SpTree.Choice)-1]
				return nil
			}
		}
	}
	return _bc.New("\u0070\u006c\u0061\u0063\u0065\u0068\u006f\u006c\u0064\u0065r\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u006c\u0069\u0064\u0065")
}
func (_eeb sort2d) Swap(i, j int) { _eeb[i], _eeb[j] = _eeb[j], _eeb[i] }

// SlideMaster is the slide master for a presentation.
type SlideMaster struct {
	_dfab *Presentation
	_gcba _ec.Relationships
	_fec  *_a.SldMaster
}

// ExtLst returns the ExtLst property.
func (_cfff PresentationProperties) ExtLst() *_a.CT_ExtensionList { return _cfff._ddad.ExtLst }
func (_ccb TextBox) getOff() *_eef.CT_Point2D {
	if _ccb._bbcd.SpPr == nil {
		_ccb._bbcd.SpPr = _eef.NewCT_ShapeProperties()
	}
	if _ccb._bbcd.SpPr.Xfrm == nil {
		_ccb._bbcd.SpPr.Xfrm = _eef.NewCT_Transform2D()
	}
	if _ccb._bbcd.SpPr.Xfrm.Off == nil {
		_ccb._bbcd.SpPr.Xfrm.Off = _eef.NewCT_Point2D()
	}
	return _ccb._bbcd.SpPr.Xfrm.Off
}

var (
	SlideScreenSize16x9 = SlideScreenSize{12192000, 6858000}
	SlideScreenSize4x3  = SlideScreenSize{9144000, 6858000}
	SlideScreenSizeA4   = SlideScreenSize{9906000, 6858000}
)

// Presentation is the a presentation base document.
type Presentation struct {
	_ec.DocBase
	_dac  *_a.Presentation
	_bfe  _ec.Relationships
	_eda  []*_a.Sld
	_gfd  []_ec.Relationships
	_bda  []*_a.SldMaster
	_gdbf []_ec.Relationships
	_gff  []*_a.SldLayout
	_dfa  []_ec.Relationships
	_eade []*_eef.Theme
	_cdf  []_ec.Relationships
	_fead _ec.TableStyles
	_bfc  PresentationProperties
	_egb  ViewProperties
	_cef  []*_eef.CT_Hyperlink
	_bfg  []*chart
	_ffac []*_a.HandoutMaster
	_fbd  []*_a.NotesMaster
	_cge  []*_gbg.XSDAny
	_bdeb map[string]string
	_bef  string
}

// SaveToFile writes the Presentation out to a file.
func (_bfce *Presentation) SaveToFile(path string) error { return _bfce.saveToFile(path, false) }

// SetHeight sets height of slide screen size with given value in EMU units.
func (_cafb *SlideScreenSize) SetHeight(val int32) { _cafb[1] = val }

// Image is an image within a slide.
type Image struct{ _cbc *_a.CT_Picture }

// Clear clears the placeholder contents and adds a single empty paragraph.  The
// empty paragrah is required by PowerPoint or it will report the file as being
// invalid.
func (_gbb PlaceHolder) Clear() {
	_gbb.ClearAll()
	_fb := _eef.NewCT_TextParagraph()
	_gbb._beb.TxBody.P = []*_eef.CT_TextParagraph{_fb}
	_fb.EndParaRPr = _eef.NewCT_TextCharacterProperties()
	_fb.EndParaRPr.LangAttr = _gbg.String("\u0065\u006e\u002dU\u0053")
}

// Save writes the presentation out to a writer in the Zip package format
func (_dff *Presentation) Save(w _cdg.Writer) error { return _dff.save(w, false) }

// PresentationProperties contains document specific properties.
type PresentationProperties struct{ _ddad *_a.PresentationPr }

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// document.
func (_cbcb *Presentation) GetImageByRelID(relID string) (_ec.ImageRef, bool) {
	for _, _fddb := range _cbcb.Images {
		if _fddb.RelID() == relID {
			return _fddb, true
		}
	}
	return _ec.ImageRef{}, false
}
func (_aaa *Presentation) Validate() error {
	if _dfcg := _aaa._dac.Validate(); _dfcg != nil {
		return _dfcg
	}
	for _gdd, _geg := range _aaa.Slides() {
		if _eeab := _geg.ValidateWithPath(_fe.Sprintf("\u0053l\u0069\u0064\u0065\u005b\u0025\u0064]", _gdd)); _eeab != nil {
			return _eeab
		}
	}
	for _fbgd, _gbba := range _aaa._bda {
		if _ddda := _gbba.ValidateWithPath(_fe.Sprintf("\u0053l\u0069d\u0065\u004d\u0061\u0073\u0074\u0065\u0072\u005b\u0025\u0064\u005d", _fbgd)); _ddda != nil {
			return _ddda
		}
	}
	for _fgeb, _gcce := range _aaa._gff {
		if _gdc := _gcce.ValidateWithPath(_fe.Sprintf("\u0053l\u0069d\u0065\u004c\u0061\u0079\u006f\u0075\u0074\u005b\u0025\u0064\u005d", _fgeb)); _gdc != nil {
			return _gdc
		}
	}
	return nil
}

// ExtractText returns text from a presentation as a PresentationText object.
func (_eb *Presentation) ExtractText() *PresentationText {
	_ef := []*SlideText{}
	for _, _gcc := range _eb.Slides() {
		_cdb := _gcc.ExtractText()
		if _cdb != nil {
			_ef = append(_ef, _cdb)
		}
	}
	return &PresentationText{Slides: _ef}
}

// Open opens and reads a document from a file (.pptx).
func Open(filename string) (*Presentation, error) {
	_dbea, _ebc := _gbc.Open(filename)
	if _ebc != nil {
		return nil, _fe.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _ebc)
	}
	defer _dbea.Close()
	_ggg, _ebc := _gbc.Stat(filename)
	if _ebc != nil {
		return nil, _fe.Errorf("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073", filename, _ebc)
	}
	_ = _ggg
	return Read(_dbea, _ggg.Size())
}

// PresentationText is an array of extracted text items which has some methods for representing extracted text.
type PresentationText struct{ Slides []*SlideText }

// HtmlPubPr returns the HtmlPubPr property.
func (_fgc PresentationProperties) HtmlPubPr() *_a.CT_HtmlPublishProperties {
	return _fgc._ddad.HtmlPubPr
}

// GridSpacing returns the GridSpacing property.
func (_efbc ViewProperties) GridSpacing() *_eef.CT_PositiveSize2D { return _efbc._ebacf.GridSpacing }

// X returns the inner wrapped XML type.
func (_cdca SlideMaster) X() *_a.SldMaster { return _cdca._fec }

// SetWidth sets width of slide screen size with given value in EMU units.
func (_bbfa *SlideScreenSize) SetWidth(val int32) { _bbfa[0] = val }

// Properties returns the properties of the TextBox.
func (_cec Image) Properties() _gg.ShapeProperties {
	if _cec._cbc.SpPr == nil {
		_cec._cbc.SpPr = _eef.NewCT_ShapeProperties()
	}
	return _gg.MakeShapeProperties(_cec._cbc.SpPr)
}

// ShowPr returns the ShowPr property.
func (_ag PresentationProperties) ShowPr() *_a.CT_ShowProperties { return _ag._ddad.ShowPr }

// Read reads a document from an io.Reader.
func Read(r _cdg.ReaderAt, size int64) (*Presentation, error) {
	const _geaf = "\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e:\u0052\u0065\u0061\u0064"
	// if !_cf.GetLicenseKey().IsLicensed() && !_bfdd {
	// 	_fe.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
	// 	_fe.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
	// 	return nil, _bc.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	// }
	_fcfd := _dfb()
	// _cbef, _gaea := _cf.GenRefId("\u0070\u0072")
	// if _gaea != nil {
	// 	_e.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _gaea)
	// 	return nil, _gaea
	// }
	// _fcfd._bef = _cbef
	// if _cbgd := _cf.Track(_fcfd._bef, _geaf); _cbgd != nil {
	// 	_e.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cbgd)
	// 	return nil, _cbgd
	// }
	_dggf, _gaea := _fgb.TempDir("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065-\u0070\u0070\u0074\u0078")
	if _gaea != nil {
		return nil, _gaea
	}
	_fcfd.TmpPath = _dggf
	_cfef, _gaea := _ea.NewReader(r, size)
	if _gaea != nil {
		return nil, _fe.Errorf("\u0070a\u0072s\u0069\u006e\u0067\u0020\u007a\u0069\u0070\u003a\u0020\u0025\u0073", _gaea)
	}
	_fdf := []*_ea.File{}
	_fdf = append(_fdf, _cfef.File...)
	_ggf := false
	for _, _ddgg := range _fdf {
		if _ddgg.FileHeader.Name == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
			_ggf = true
			break
		}
	}
	if _ggf {
		_fcfd.CreateCustomProperties()
	}
	_fdb := _cd.DecodeMap{}
	_fdb.SetOnNewRelationshipFunc(_fcfd.onNewRelationship)
	_fdb.AddTarget(_gbg.ContentTypesFilename, _fcfd.ContentTypes.X(), "", 0)
	_fdb.AddTarget(_gbg.BaseRelsFilename, _fcfd.Rels.X(), "", 0)
	if _ccge := _fdb.Decode(_fdf); _ccge != nil {
		return nil, _ccge
	}
	for _, _ffbg := range _fdf {
		if _ffbg == nil {
			continue
		}
		if _bceg := _fcfd.AddExtraFileFromZip(_ffbg); _bceg != nil {
			return nil, _bceg
		}
	}
	if _ggf {
		_cgec := false
		for _, _gdgea := range _fcfd.Rels.X().Relationship {
			if _gdgea.TargetAttr == "\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c" {
				_cgec = true
				break
			}
		}
		if !_cgec {
			_fcfd.AddCustomRelationships()
		}
	}
	return _fcfd, nil
}

// GetColorBySchemeColor returns *dml.CT_Color mapped to scheme colors like dk1, lt1 etc. depending on what theme is used in the presentation.
func (_bcea *Presentation) GetColorBySchemeColor(schClr _eef.ST_SchemeColorVal) *_eef.CT_Color {
	if len(_bcea._bda) == 0 || len(_bcea._eade) == 0 {
		return nil
	}
	var _daf _eef.ST_ColorSchemeIndex
	_facb := _bcea._bda[0]
	_eecc := _facb.ClrMap
	switch schClr.String() {
	case "\u0062\u0067\u0031":
		_daf = _eecc.Bg1Attr
	case "\u0062\u0067\u0032":
		_daf = _eecc.Bg2Attr
	case "\u0074\u0078\u0031":
		_daf = _eecc.Tx1Attr
	case "\u0074\u0078\u0032":
		_daf = _eecc.Tx2Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0031":
		_daf = _eecc.Accent1Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0032":
		_daf = _eecc.Accent2Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0033":
		_daf = _eecc.Accent3Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0034":
		_daf = _eecc.Accent4Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0035":
		_daf = _eecc.Accent5Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0036":
		_daf = _eecc.Accent6Attr
	case "\u0068\u006c\u0069n\u006b":
		_daf = _eecc.HlinkAttr
	case "\u0066\u006f\u006c\u0048\u006c\u0069\u006e\u006b":
		_daf = _eecc.FolHlinkAttr
	case "\u0064\u006b\u0031":
		_daf = _eef.ST_ColorSchemeIndexDk1
	case "\u0064\u006b\u0032":
		_daf = _eef.ST_ColorSchemeIndexDk2
	case "\u006c\u0074\u0031":
		_daf = _eef.ST_ColorSchemeIndexLt1
	case "\u006c\u0074\u0032":
		_daf = _eef.ST_ColorSchemeIndexLt2
	default:
		_daf = _eef.ST_ColorSchemeIndexUnset
	}
	_ceeg := _bcea._eade[0]
	_ecd := _ceeg.ThemeElements
	if _ecd == nil {
		return nil
	}
	var _acb *_eef.CT_Color
	_bcf := _ecd.ClrScheme
	switch _daf.String() {
	case "\u0064\u006b\u0031":
		_acb = _bcf.Dk1
	case "\u0064\u006b\u0032":
		_acb = _bcf.Dk2
	case "\u006c\u0074\u0031":
		_acb = _bcf.Lt1
	case "\u006c\u0074\u0032":
		_acb = _bcf.Lt2
	case "\u0061c\u0063\u0065\u006e\u0074\u0031":
		_acb = _bcf.Accent1
	case "\u0061c\u0063\u0065\u006e\u0074\u0032":
		_acb = _bcf.Accent2
	case "\u0061c\u0063\u0065\u006e\u0074\u0033":
		_acb = _bcf.Accent3
	case "\u0061c\u0063\u0065\u006e\u0074\u0034":
		_acb = _bcf.Accent4
	case "\u0061c\u0063\u0065\u006e\u0074\u0035":
		_acb = _bcf.Accent5
	case "\u0061c\u0063\u0065\u006e\u0074\u0036":
		_acb = _bcf.Accent6
	case "\u0068\u006c\u0069n\u006b":
		_acb = _bcf.Hlink
	case "\u0066\u006f\u006c\u0048\u006c\u0069\u006e\u006b":
		_acb = _bcf.FolHlink
	default:
		return nil
	}
	return _acb
}

// Size returns slide size value as SlideScreenSize.
func (_gddf *SlideSize) Size() SlideScreenSize {
	return SlideScreenSize{_gddf._deg.CxAttr, _gddf._deg.CyAttr}
}

// ShowCommentsAttr returns the WebPr property.
func (_beff ViewProperties) ShowCommentsAttr() *bool { return _beff._ebacf.ShowCommentsAttr }

// GetColorBySchemeColor returns *dml.CT_Color mapped to scheme colors like dk1, lt1 etc. depending on what theme is used in the presentation.
func (_caad *Slide) GetColorBySchemeColor(schClr _eef.ST_SchemeColorVal) *_eef.CT_Color {
	_caad.ensureClrMap()
	_gcbe := _caad._fag
	if _gcbe == nil {
		return nil
	}
	var _aaf _eef.ST_ColorSchemeIndex
	switch schClr.String() {
	case "\u0062\u0067\u0031":
		_aaf = _gcbe.Bg1Attr
	case "\u0062\u0067\u0032":
		_aaf = _gcbe.Bg2Attr
	case "\u0074\u0078\u0031":
		_aaf = _gcbe.Tx1Attr
	case "\u0074\u0078\u0032":
		_aaf = _gcbe.Tx2Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0031":
		_aaf = _gcbe.Accent1Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0032":
		_aaf = _gcbe.Accent2Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0033":
		_aaf = _gcbe.Accent3Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0034":
		_aaf = _gcbe.Accent4Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0035":
		_aaf = _gcbe.Accent5Attr
	case "\u0061c\u0063\u0065\u006e\u0074\u0036":
		_aaf = _gcbe.Accent6Attr
	case "\u0068\u006c\u0069n\u006b":
		_aaf = _gcbe.HlinkAttr
	case "\u0066\u006f\u006c\u0048\u006c\u0069\u006e\u006b":
		_aaf = _gcbe.FolHlinkAttr
	case "\u0064\u006b\u0031":
		_aaf = _eef.ST_ColorSchemeIndexDk1
	case "\u0064\u006b\u0032":
		_aaf = _eef.ST_ColorSchemeIndexDk2
	case "\u006c\u0074\u0031":
		_aaf = _eef.ST_ColorSchemeIndexLt1
	case "\u006c\u0074\u0032":
		_aaf = _eef.ST_ColorSchemeIndexLt2
	default:
		_aaf = _eef.ST_ColorSchemeIndexUnset
	}
	_cebb := _caad._gac._eade[0]
	_febb := _cebb.ThemeElements
	if _febb == nil {
		return nil
	}
	var _gade *_eef.CT_Color
	_bafa := _febb.ClrScheme
	switch _aaf.String() {
	case "\u0064\u006b\u0031":
		_gade = _bafa.Dk1
	case "\u0064\u006b\u0032":
		_gade = _bafa.Dk2
	case "\u006c\u0074\u0031":
		_gade = _bafa.Lt1
	case "\u006c\u0074\u0032":
		_gade = _bafa.Lt2
	case "\u0061c\u0063\u0065\u006e\u0074\u0031":
		_gade = _bafa.Accent1
	case "\u0061c\u0063\u0065\u006e\u0074\u0032":
		_gade = _bafa.Accent2
	case "\u0061c\u0063\u0065\u006e\u0074\u0033":
		_gade = _bafa.Accent3
	case "\u0061c\u0063\u0065\u006e\u0074\u0034":
		_gade = _bafa.Accent4
	case "\u0061c\u0063\u0065\u006e\u0074\u0035":
		_gade = _bafa.Accent5
	case "\u0061c\u0063\u0065\u006e\u0074\u0036":
		_gade = _bafa.Accent6
	case "\u0068\u006c\u0069n\u006b":
		_gade = _bafa.Hlink
	case "\u0066\u006f\u006c\u0048\u006c\u0069\u006e\u006b":
		_gade = _bafa.FolHlink
	default:
		return nil
	}
	return _gade
}

// Text returns text from a presentation as one string separated with line breaks.
func (_gae *PresentationText) Text() string {
	_dfc := _db.NewBuffer([]byte{})
	for _, _dgbd := range _gae.Slides {
		_dfc.WriteString(_dgbd.Text())
	}
	return _dfc.String()
}

// AddTextBox adds an empty textbox to a slide.
func (_gce Slide) AddTextBox() TextBox {
	_fcebb := _a.NewCT_GroupShapeChoice()
	_gce._gfgd.CSld.SpTree.Choice = append(_gce._gfgd.CSld.SpTree.Choice, _fcebb)
	_dcce := _a.NewCT_Shape()
	_fcebb.Sp = append(_fcebb.Sp, _dcce)
	_dcce.SpPr = _eef.NewCT_ShapeProperties()
	_dcce.SpPr.Xfrm = _eef.NewCT_Transform2D()
	_dcce.SpPr.PrstGeom = _eef.NewCT_PresetGeometry2D()
	_dcce.SpPr.PrstGeom.PrstAttr = _eef.ST_ShapeTypeRect
	_dcce.NvSpPr = _a.NewCT_ShapeNonVisual()
	_dcce.NvSpPr.CNvSpPr = _eef.NewCT_NonVisualDrawingShapeProps()
	_bdc := true
	_dcce.NvSpPr.CNvSpPr.TxBoxAttr = &_bdc
	_dcce.TxBody = _eef.NewCT_TextBody()
	_dcce.TxBody.BodyPr = _eef.NewCT_TextBodyProperties()
	_dcce.TxBody.BodyPr.WrapAttr = _eef.ST_TextWrappingTypeSquare
	_dcce.TxBody.BodyPr.SpAutoFit = _eef.NewCT_TextShapeAutofit()
	_afb := TextBox{_dcce}
	_afb.Properties().SetWidth(3 * _cb.Inch)
	_afb.Properties().SetHeight(1 * _cb.Inch)
	_afb.Properties().SetPosition(0, 0)
	return _afb
}

// Slide represents a slide of a presentation.
type Slide struct {
	_gddb *_a.CT_SlideIdListEntry
	_gfgd *_a.Sld
	_gac  *Presentation
	_fag  *_eef.CT_ColorMapping
}

func _dd(_be *Presentation, _fad []*_a.CT_GroupShapeChoice, _gdf []rectangle, _ead []*TextItem) []*TextItem {
	for _, _ceb := range _fad {
		_bb := append([]rectangle{}, _gdf...)
		for _, _ada := range _ceb.Sp {
			_ead = append(_ead, _egg(_be, _ada, nil, nil, _ada.SpPr.Xfrm, 0, _gdf, _ada.TxBody.P)...)
		}
		for _, _cee := range _ceb.GraphicFrame {
			if _cee != nil && _cee.Graphic != nil && _cee.Graphic.GraphicData != nil {
				_efd := _cee.Xfrm
				for _, _gcg := range _cee.Graphic.GraphicData.Any {
					if _de, _ab := _gcg.(*_eef.Tbl); _ab {
						_cc := &_de.CT_Table
						_dfd := 0
						for _ed, _ffa := range _de.Tr {
							for _gcd, _cfa := range _ffa.Tc {
								_eba := &TableInfo{Table: _cc, Row: _ffa, Cell: _cfa, RowIndex: _ed, ColIndex: _gcd}
								_ead = append(_ead, _egg(_be, nil, _cee, _eba, _efd, _dfd, _gdf, _cfa.TxBody.P)...)
								_dfd++
							}
						}
					}
				}
			}
		}
		for _, _bf := range _ceb.GrpSp {
			if _bf.GrpSpPr != nil {
				_fac := _bf.GrpSpPr.Xfrm
				var _bcd, _da int64
				if _fac.Off != nil {
					_bfd, _ga := _fac.Off.XAttr.ST_CoordinateUnqualified, _fac.Off.YAttr.ST_CoordinateUnqualified
					if _bfd != nil && _ga != nil {
						if _gcgg := _fac.Ext; _gcgg != nil {
							_bcd, _da = _gcgg.CxAttr, _gcgg.CyAttr
						}
						_bb = append(_bb, rectangle{_cfc: *_bfd, _dgb: *_ga, _cdga: *_bfd + _bcd, _cbg: *_ga + _da})
					}
				}
			}
			_ead = _dd(_be, _bf.Choice, _bb, _ead)
		}
	}
	return _ead
}

// X returns the inner wrapped XML type.
func (_cgc PresentationProperties) X() *_a.PresentationPr { return _cgc._ddad }

// Type returns the type of the slide layout.
func (_acde SlideLayout) Type() _a.ST_SlideLayoutType { return _acde._dcbe.TypeAttr }

// NewPresentationProperties constructs a new PresentationProperties.
func NewPresentationProperties() PresentationProperties {
	return PresentationProperties{_ddad: _a.NewPresentationPr()}
}

// GetTextBoxes returns a list of all text boxes from a slide.
func (_ebbd Slide) GetTextBoxes() []*TextBox {
	_ebcc := []*TextBox{}
	_dfda := _ebbd._gfgd.CSld.SpTree.Choice
	for _, _gec := range _dfda {
		for _, _dce := range _gec.Sp {
			if _dce.NvSpPr.CNvSpPr.TxBoxAttr != nil && *_dce.NvSpPr.CNvSpPr.TxBoxAttr {
				_ebcc = append(_ebcc, &TextBox{_dce})
			}
		}
	}
	return _ebcc
}

// X returns TextBox's underlying *pml.CT_Shape.
func (_fdebe TextBox) X() *_a.CT_Shape { return _fdebe._bbcd }

// GetLayoutByName retrieves a slide layout given a layout name.
func (_cea *Presentation) GetLayoutByName(name string) (SlideLayout, error) {
	for _, _dgc := range _cea._gff {
		if _dgc.CSld.NameAttr != nil && name == *_dgc.CSld.NameAttr {
			return SlideLayout{_dgc}, nil
		}
	}
	return SlideLayout{}, _bc.New("\u0075\u006eab\u006c\u0065\u0020t\u006f\u0020\u0066\u0069nd \u006cay\u006f\u0075\u0074\u0020\u0077\u0069\u0074h \u0074\u0068\u0061\u0074\u0020\u006e\u0061m\u0065")
}

// Width returns slide screen size width in EMU units.
func (_aabb *SlideScreenSize) Width() int32 { return _aabb[0] }

// GetTableStyleById returns *dml.CT_TableStyle by its style id.
func (_dcgg *Presentation) GetTableStyleById(id string) *_eef.CT_TableStyle {
	_fffbe := _dcgg._fead.TblStyle()
	for _, _abfd := range _fffbe {
		if _abfd.StyleIdAttr == id {
			return _abfd
		}
	}
	return nil
}

var _bfdd = false

func (_dccb *Presentation) saveToFile(_dcg string, _dag bool) error {
	_efbd, _eaf := _gbc.Create(_dcg)
	if _eaf != nil {
		return _eaf
	}
	defer _efbd.Close()
	return _dccb.save(_efbd, _dag)
}

// TableInfo is used for keep information about a table, a row and a cell where the text is located.
type TableInfo struct {
	Table    *_eef.CT_Table
	Row      *_eef.CT_TableRow
	Cell     *_eef.CT_TableCell
	RowIndex int
	ColIndex int
}

// Name returns the name of the slide layout.
func (_fgde SlideLayout) Name() string {
	if _fgde._dcbe.CSld != nil && _fgde._dcbe.CSld.NameAttr != nil {
		return *_fgde._dcbe.CSld.NameAttr
	}
	return ""
}
func (_bec *Presentation) save(_ccfg _cdg.Writer, _aeag bool) error {
	const _faa = "\u0050\u0072\u0065\u0073en\u0074\u0061\u0074\u0069\u006f\u006e\u003a\u0070\u002e\u0053\u0061\u0076\u0065"
	if _ffef := _bec._dac.Validate(); _ffef != nil {
		_e.Log.Debug("\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u0020i\u006e\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u003a\u0020\u0025\u0073", _ffef)
	}
	// if !_cf.GetLicenseKey().IsLicensed() && !_bfdd {
	// 	_fe.Println("\u0055\u006e\u006ci\u0063\u0065\u006e\u0073e\u0064\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u006f\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065")
	// 	_fe.Println("\u002d\u0020\u0047e\u0074\u0020\u0061\u0020\u0074\u0072\u0069\u0061\u006c\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u006f\u006e\u0020\u0068\u0074\u0074\u0070\u0073\u003a\u002f\u002fu\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f")
	// 	return _bc.New("\u0075\u006e\u0069\u006f\u0066\u0066\u0069\u0063\u0065\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0072\u0065\u0071\u0075i\u0072\u0065\u0064")
	// }
	// if len(_bec._bef) == 0 {
	// 	_decce, _cad := _cf.GenRefId("\u0070\u0077")
	// 	if _cad != nil {
	// 		_e.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _cad)
	// 		return _cad
	// 	}
	// 	_bec._bef = _decce
	// }
	// if _acg := _cf.Track(_bec._bef, _faa); _acg != nil {
	// 	_e.Log.Error("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v", _acg)
	// 	return _acg
	// }
	if _aeag {
		_bec.ContentTypes.RemoveOverride("\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006ff\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u006d\u006c\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u002e\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
		_bec.ContentTypes.EnsureOverride("/\u0070\u0070\u0074\u002fpr\u0065s\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u002e\u0078\u006d\u006c", "\u0061\u0070pl\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074s\u002d\u006f\u0066\u0066ic\u0065\u0064o\u0063u\u006d\u0065\u006e\u0074\u002e\u0070r\u0065\u0073\u0065n\u0074\u0061t\u0069\u006f\u006e\u006d\u006c\u002e\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065.\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
	} else {
		_bec.ContentTypes.RemoveOverride("\u0061\u0070pl\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074s\u002d\u006f\u0066\u0066ic\u0065\u0064o\u0063u\u006d\u0065\u006e\u0074\u002e\u0070r\u0065\u0073\u0065n\u0074\u0061t\u0069\u006f\u006e\u006d\u006c\u002e\u0074\u0065\u006d\u0070\u006c\u0061\u0074\u0065.\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
		_bec.ContentTypes.EnsureOverride("/\u0070\u0070\u0074\u002fpr\u0065s\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006ff\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u006d\u006c\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u002e\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
	}
	_ffb := _gbg.DocTypePresentation
	_edf := _ea.NewWriter(_ccfg)
	defer _edf.Close()
	if _egc := _cd.MarshalXML(_edf, _gbg.BaseRelsFilename, _bec.Rels.X()); _egc != nil {
		return _egc
	}
	if _gea := _cd.MarshalXMLByType(_edf, _ffb, _gbg.ExtendedPropertiesType, _bec.AppProperties.X()); _gea != nil {
		return _gea
	}
	if _cfda := _cd.MarshalXMLByType(_edf, _ffb, _gbg.CorePropertiesType, _bec.CoreProperties.X()); _cfda != nil {
		return _cfda
	}
	if _fffg := _cd.MarshalXMLByType(_edf, _ffb, _gbg.PresentationPropertiesType, _bec._bfc.X()); _fffg != nil {
		return _fffg
	}
	if _ege := _cd.MarshalXMLByType(_edf, _ffb, _gbg.ViewPropertiesType, _bec._egb.X()); _ege != nil {
		return _ege
	}
	if _fba := _cd.MarshalXMLByType(_edf, _ffb, _gbg.TableStylesType, _bec._fead.X()); _fba != nil {
		return _fba
	}
	if _bec.CustomProperties.X() != nil {
		if _bfdg := _cd.MarshalXMLByType(_edf, _ffb, _gbg.CustomPropertiesType, _bec.CustomProperties.X()); _bfdg != nil {
			return _bfdg
		}
	}
	if _bec.Thumbnail != nil {
		_dgd, _dcc := _edf.Create("\u0064\u006f\u0063Pr\u006f\u0070\u0073\u002f\u0074\u0068\u0075\u006d\u0062\u006e\u0061\u0069\u006c\u002e\u006a\u0070\u0065\u0067")
		if _dcc != nil {
			return _dcc
		}
		if _eea := _b.Encode(_dgd, _bec.Thumbnail, nil); _eea != nil {
			return _eea
		}
	}
	_afc := _gbg.AbsoluteFilename(_ffb, _gbg.OfficeDocumentType, 0)
	if _ceba := _cd.MarshalXML(_edf, _afc, _bec._dac); _ceba != nil {
		return _ceba
	}
	if _cfab := _cd.MarshalXML(_edf, _cd.RelationsPathFor(_afc), _bec._bfe.X()); _cfab != nil {
		return _cfab
	}
	for _gfdc, _gdbfd := range _bec._eda {
		_dad := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.SlideType, _gfdc+1)
		_cd.MarshalXML(_edf, _dad, _gdbfd)
		if !_bec._gfd[_gfdc].IsEmpty() {
			_bac := _cd.RelationsPathFor(_dad)
			_cd.MarshalXML(_edf, _bac, _bec._gfd[_gfdc].X())
		}
	}
	for _cgg, _cbbf := range _bec._bda {
		_fbg := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.SlideMasterType, _cgg+1)
		_cd.MarshalXML(_edf, _fbg, _cbbf)
		if !_bec._gdbf[_cgg].IsEmpty() {
			_cade := _cd.RelationsPathFor(_fbg)
			_cd.MarshalXML(_edf, _cade, _bec._gdbf[_cgg].X())
		}
	}
	for _fab, _bcc := range _bec._gff {
		_bag := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.SlideLayoutType, _fab+1)
		_cd.MarshalXML(_edf, _bag, _bcc)
		if !_bec._dfa[_fab].IsEmpty() {
			_befg := _cd.RelationsPathFor(_bag)
			_cd.MarshalXML(_edf, _befg, _bec._dfa[_fab].X())
		}
	}
	for _ffee, _cgga := range _bec._eade {
		_cgd := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.ThemeType, _ffee+1)
		_cd.MarshalXML(_edf, _cgd, _cgga)
		if !_bec._cdf[_ffee].IsEmpty() {
			_fbf := _cd.RelationsPathFor(_cgd)
			_cd.MarshalXML(_edf, _fbf, _bec._cdf[_ffee].X())
		}
	}
	for _eabb, _ebab := range _bec._bfg {
		_cbda := _gbg.AbsoluteFilename(_ffb, _gbg.ChartType, _eabb+1)
		_cd.MarshalXML(_edf, _cbda, _ebab)
	}
	for _egbf, _eec := range _bec._ffac {
		_aab := _gbg.AbsoluteFilename(_ffb, _gbg.HandoutMasterType, _egbf+1)
		_cd.MarshalXML(_edf, _aab, _eec)
	}
	for _cga, _gaec := range _bec._fbd {
		_gaa := _gbg.AbsoluteFilename(_ffb, _gbg.NotesMasterType, _cga+1)
		_cd.MarshalXML(_edf, _gaa, _gaec)
	}
	for _cfabc, _edfc := range _bec._cge {
		_efdc := _gbg.AbsoluteFilename(_ffb, _gbg.CustomXMLType, _cfabc+1)
		_cd.MarshalXML(_edf, _efdc, _edfc)
	}
	for _bfac, _cdfb := range _bec.Images {
		if _fga := _ec.AddImageToZip(_edf, _cdfb, _bfac+1, _gbg.DocTypePresentation); _fga != nil {
			return _fga
		}
	}
	_bec.ContentTypes.EnsureDefault("\u0070\u006e\u0067", "\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg")
	_bec.ContentTypes.EnsureDefault("\u006a\u0070\u0065\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_bec.ContentTypes.EnsureDefault("\u006a\u0070\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_bec.ContentTypes.EnsureDefault("\u0077\u006d\u0066", "i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066")
	if _bfed := _cd.MarshalXML(_edf, _gbg.ContentTypesFilename, _bec.ContentTypes.X()); _bfed != nil {
		return _bfed
	}
	if _eag := _bec.WriteExtraFiles(_edf); _eag != nil {
		return _eag
	}
	return nil
}
func (_afa *Slide) ensureClrMap() {
	if len(_afa._gac._bda) == 0 || len(_afa._gac._eade) == 0 {
		return
	}
	_gadf := _afa._gac._bda[0]
	_fcbd := _gadf.ClrMap
	if _eae := _afa._gfgd.ClrMapOvr; _eae != nil {
		if _ebfd := _eae.Choice; _ebfd != nil {
			if _ebfd.MasterClrMapping == nil {
				if _acd := _ebfd.OverrideClrMapping; _acd != nil {
					if _acd.Bg1Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Bg1Attr = _acd.Bg1Attr
					}
					if _acd.Tx1Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Tx1Attr = _acd.Tx1Attr
					}
					if _acd.Bg2Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Bg2Attr = _acd.Bg2Attr
					}
					if _acd.Tx2Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Tx2Attr = _acd.Tx2Attr
					}
					if _acd.Accent1Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Accent1Attr = _acd.Accent1Attr
					}
					if _acd.Accent2Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Accent2Attr = _acd.Accent2Attr
					}
					if _acd.Accent3Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Accent3Attr = _acd.Accent3Attr
					}
					if _acd.Accent4Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Accent4Attr = _acd.Accent4Attr
					}
					if _acd.Accent5Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Accent5Attr = _acd.Accent5Attr
					}
					if _acd.Accent6Attr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.Accent6Attr = _acd.Accent6Attr
					}
					if _acd.HlinkAttr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.HlinkAttr = _acd.HlinkAttr
					}
					if _acd.FolHlinkAttr != _eef.ST_ColorSchemeIndexUnset {
						_fcbd.FolHlinkAttr = _acd.FolHlinkAttr
					}
				}
			}
		}
	}
	_afa._fag = _fcbd
}

// Less is for implementing sorting of two locations. Symbols share the same location if they are in the same paragraph or table. One location is 'less' than another first by y coordinate, if y coordinates are equal or differ by less than yEpsilon, then x coordinates are compared, then if they are also equal, indexes of locations in the table are compared, then positions of locations in a paragraph.
func (_ced sort2d) Less(i, j int) bool {
	_bde, _dde := _ced[i], _ced[j]
	_aa, _bce := _bde._dg, _dde._dg
	_dee, _cfd := len(_aa)-1, len(_bce)-1
	_fea, _gbe := 0, 0
	for {
		_gga, _bca, _dfg, _ebb, _fce, _ccf, _ece, _ddgd := _aa[_fea]._dgb, _bce[_gbe]._dgb, _aa[_fea]._cbg, _bce[_gbe]._cbg, _aa[_fea]._cfc, _bce[_gbe]._cfc, _aa[_fea]._cdga, _bce[_gbe]._cdga
		if _gga == _bca || ((_eg.Abs(float64(_gga)-float64(_bca)) < _cdd) && ((_gga >= _bca && _gga <= _ebb) || (_bca >= _gga && _bca <= _dfg)) && (_ece < _ccf || _fce > _ddgd)) {
			if _fce == _ccf {
				if _fea < _dee && _gbe < _cfd {
					_fea++
					_gbe++
					continue
				}
				if _fea >= _dee && _gbe >= _cfd {
					break
				}
				return _fea >= _dee
			} else {
				return _fce < _ccf
			}
		} else {
			return _gga < _bca
		}
	}
	_egd, _cfdd, _dda, _dec := _bde._ca, _dde._ca, _bde._ecf, _dde._ecf
	if _egd == _cfdd {
		return _dda <= _dec
	}
	return _egd < _cfdd
}

// ValidateWithPath validates the slide passing path informaton for a better
// error message.
func (_bbd Slide) ValidateWithPath(path string) error {
	if _egbg := _bbd._gfgd.ValidateWithPath(path); _egbg != nil {
		return _egbg
	}
	for _, _ggge := range _bbd._gfgd.CSld.SpTree.Choice {
		for _, _daec := range _ggge.Sp {
			if _daec.TxBody != nil {
				if len(_daec.TxBody.P) == 0 {
					return _bc.New(path + "\u0020\u003a \u0073\u006c\u0069\u0064\u0065 \u0073\u0068\u0061\u0070\u0065 \u0077\u0069\u0074\u0068\u0020\u0061\u0020\u0074\u0078\u0062\u006f\u0064\u0079\u0020\u006d\u0075\u0073\u0074\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0020\u0070\u0061\u0072\u0061\u0067\u0072\u0061\u0070\u0068\u0073")
				}
			}
		}
	}
	return nil
}
func (_ba *chart) Target() string { return _ba._gc }

// Paragraphs returns the paragraphs defined in the placeholder.
func (_eadc PlaceHolder) Paragraphs() []_gg.Paragraph {
	_feab := []_gg.Paragraph{}
	for _, _cff := range _eadc._beb.TxBody.P {
		_feab = append(_feab, _gg.MakeParagraph(_cff))
	}
	return _feab
}

// OutlineViewPr returns the OutlineViewPr property.
func (_bbg ViewProperties) OutlineViewPr() *_a.CT_OutlineViewProperties {
	return _bbg._ebacf.OutlineViewPr
}
func (_gdg *chart) X() *_f.ChartSpace { return _gdg._gd }

// AddImage adds an image to the document package, returning a reference that
// can be used to add the image to a run and place it in the document contents.
func (_cfcb *Presentation) AddImage(i _ec.Image) (_ec.ImageRef, error) {
	_dged := _ec.MakeImageRef(i, &_cfcb.DocBase, _cfcb._bfe)
	if i.Data == nil && i.Path == "" {
		return _dged, _bc.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0064\u0061t\u0061\u0020\u006f\u0072\u0020\u0061\u0020\u0070\u0061\u0074\u0068")
	}
	if i.Format == "" {
		return _dged, _bc.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006d\u0075\u0073\u0074 \u0068\u0061\u0076\u0065\u0020\u0061\u0020v\u0061\u006c\u0069\u0064\u0020\u0066\u006f\u0072\u006d\u0061\u0074")
	}
	if i.Size.X == 0 || i.Size.Y == 0 {
		return _dged, _bc.New("\u0069\u006d\u0061\u0067e\u0020\u006d\u0075\u0073\u0074\u0020\u0068\u0061\u0076\u0065 \u0061 \u0076\u0061\u006c\u0069\u0064\u0020\u0073i\u007a\u0065")
	}
	if i.Path != "" {
		_deee := _fgb.Add(i.Path)
		if _deee != nil {
			return _dged, _deee
		}
	}
	_cfcb.Images = append(_cfcb.Images, _dged)
	_cfcb.ContentTypes.EnsureDefault("\u0070\u006e\u0067", "\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg")
	_cfcb.ContentTypes.EnsureDefault("\u006a\u0070\u0065\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_cfcb.ContentTypes.EnsureDefault("\u006a\u0070\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_cfcb.ContentTypes.EnsureDefault("\u0077\u006d\u0066", "i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066")
	_cfcb.ContentTypes.EnsureDefault(i.Format, "\u0069\u006d\u0061\u0067\u0065\u002f"+i.Format)
	return _dged, nil
}

// AddParagraph adds a paragraph to the text box
func (_bfgff TextBox) AddParagraph() _gg.Paragraph {
	_acgg := _eef.NewCT_TextParagraph()
	_bfgff._bbcd.TxBody.P = append(_bfgff._bbcd.TxBody.P, _acgg)
	return _gg.MakeParagraph(_acgg)
}

// SetTextAnchor controls the text anchoring
func (_cgce TextBox) SetTextAnchor(a _eef.ST_TextAnchoringType) {
	_cgce._bbcd.TxBody.BodyPr = _eef.NewCT_TextBodyProperties()
	_cgce._bbcd.TxBody.BodyPr.AnchorAttr = a
}

// X returns the inner wrapped XML type.
func (_aag PlaceHolder) X() *_a.CT_Shape { return _aag._beb }

// X returns the inner wrapped XML type.
func (_effa ViewProperties) X() *_a.ViewPr { return _effa._ebacf }

// X returns the inner wrapped XML type.
func (_fdebb Slide) X() *_a.Sld { return _fdebb._gfgd }

// PrnPr returns the PrnPr property.
func (_eee PresentationProperties) PrnPr() *_a.CT_PrintProperties { return _eee._ddad.PrnPr }

// SaveToFileAsTemplate writes the Presentation out to a file as a template.
func (_cgea *Presentation) SaveToFileAsTemplate(path string) error {
	return _cgea.saveToFile(path, true)
}

// LastViewAttr returns the LastViewAttr property.
func (_feg ViewProperties) LastViewAttr() _a.ST_ViewType { return _feg._ebacf.LastViewAttr }
func _dba(_eab []*_a.CT_GroupShapeChoice) []*_a.CT_GroupShapeChoice {
	var _acc []*_a.CT_GroupShapeChoice
	for _, _fde := range _eab {
		if len(_fde.Pic) == 0 {
			_acc = append(_acc, _fde)
		}
	}
	return _acc
}

// SetText sets the text of a placeholder for the initial paragraph. This is a
// shortcut method that is useful for things like titles which only contain a
// single paragraph.
func (_fcb PlaceHolder) SetText(text string) {
	_fcb.Clear()
	_cbcc := _eef.NewEG_TextRun()
	_cbcc.R = _eef.NewCT_RegularTextRun()
	_cbcc.R.T = text
	if len(_fcb._beb.TxBody.P) == 0 {
		_fcb._beb.TxBody.P = append(_fcb._beb.TxBody.P, _eef.NewCT_TextParagraph())
	}
	_fcb._beb.TxBody.P[0].EG_TextRun = nil
	_fcb._beb.TxBody.P[0].EG_TextRun = append(_fcb._beb.TxBody.P[0].EG_TextRun, _cbcc)
}
func (_ggga *Slide) getSlideRels() _ec.Relationships {
	_cda := _ggga._gac
	for _fbc, _ggfc := range _cda.Slides() {
		if *_ggga._gfgd == *_ggfc._gfgd {
			return _cda._gfd[_fbc]
		}
	}
	return _ec.Relationships{}
}

// GetSlideLayout returns a slide layout related to the slide.
func (_gffb *Slide) GetSlideLayout() *_a.SldLayout {
	for _ecdb := range _gffb._gac.Slides() {
		_caf := _gffb._gac._gfd[_ecdb]
		for _, _fda := range _caf.Relationships() {
			if _fda.Type() == _gbg.SlideLayoutType {
				if _dbf, _abeg := _fcf.StringToNumbers(_fda.Target()); _abeg {
					return _gffb._gac._gff[_dbf-1]
				}
				return nil
			}
		}
	}
	return nil
}

// Index returns the placeholder index
func (_dea PlaceHolder) Index() uint32 {
	if _dea._beb.NvSpPr.NvPr.Ph.IdxAttr == nil {
		return 0
	}
	return *_dea._beb.NvSpPr.NvPr.Ph.IdxAttr
}

// ExtLst returns the ExtLst property.
func (_abd ViewProperties) ExtLst() *_a.CT_ExtensionList { return _abd._ebacf.ExtLst }

// ClearAll completely clears a placeholder. To be useable, at least one
// paragraph must be added after ClearAll via AddParagraph.
func (_bae PlaceHolder) ClearAll() {
	_bae._beb.SpPr = _eef.NewCT_ShapeProperties()
	_bae._beb.TxBody = _eef.NewCT_TextBody()
	_bae._beb.TxBody.LstStyle = _eef.NewCT_TextListStyle()
}

// AddImageToRels adds an image relationship to a slide without putting image on the slide.
func (_fbfb Slide) AddImageToRels(img _ec.ImageRef) string {
	_gbed := 0
	for _ffg, _eebc := range _fbfb._gac.Images {
		if _eebc == img {
			_gbed = _ffg + 1
			break
		}
	}
	var _gge string
	for _efc, _aabg := range _fbfb._gac.Slides() {
		if _aabg._gfgd == _fbfb._gfgd {
			_bgd := _fe.Sprintf("\u002e\u002e\u002f\u006ded\u0069\u0061\u002f\u0069\u006d\u0061\u0067\u0065\u0025\u0064\u002e\u0025\u0073", _gbed, img.Format())
			_fgg := _fbfb._gac._gfd[_efc].AddRelationship(_bgd, _gbg.ImageType)
			_gge = _fgg.ID()
		}
	}
	return _gge
}

// X returns the inner wrapped XML type.
func (_cafbb SlideLayout) X() *_a.SldLayout { return _cafbb._dcbe }

// SetOffsetY sets vertical offset of text box in distance units (see measurement package).
func (_gbad TextBox) SetOffsetY(offY float64) {
	_edad := _gbad.getOff()
	_dgcf := _cb.ToEMU(offY)
	_edad.YAttr = _eef.ST_Coordinate{ST_CoordinateUnqualified: &_dgcf}
}

// AddSlide adds a new slide to the presentation.
func (_dga *Presentation) AddSlide() Slide {
	_baa := _a.NewCT_SlideIdListEntry()
	_baa.IdAttr = _dga.nextSlideID()
	_dga._dac.SldIdLst.SldId = append(_dga._dac.SldIdLst.SldId, _baa)
	_dfge := _a.NewSld()
	_dfge.CSld.SpTree.NvGrpSpPr.CNvPr.IdAttr = 1
	_dfge.CSld.SpTree.GrpSpPr.Xfrm = _eef.NewCT_GroupTransform2D()
	_dfge.CSld.SpTree.GrpSpPr.Xfrm.Off = _eef.NewCT_Point2D()
	_dfge.CSld.SpTree.GrpSpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = _gbg.Int64(0)
	_dfge.CSld.SpTree.GrpSpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = _gbg.Int64(0)
	_dfge.CSld.SpTree.GrpSpPr.Xfrm.Ext = _eef.NewCT_PositiveSize2D()
	_dfge.CSld.SpTree.GrpSpPr.Xfrm.Ext.CxAttr = int64(0 * _cb.Point)
	_dfge.CSld.SpTree.GrpSpPr.Xfrm.Ext.CyAttr = int64(0 * _cb.Point)
	_dfge.CSld.SpTree.GrpSpPr.Xfrm.ChOff = _dfge.CSld.SpTree.GrpSpPr.Xfrm.Off
	_dfge.CSld.SpTree.GrpSpPr.Xfrm.ChExt = _dfge.CSld.SpTree.GrpSpPr.Xfrm.Ext
	_dga._eda = append(_dga._eda, _dfge)
	_ddd := _dga._bfe.AddAutoRelationship(_gbg.DocTypePresentation, _gbg.OfficeDocumentType, len(_dga._eda), _gbg.SlideType)
	_baa.RIdAttr = _ddd.ID()
	_bgg := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.SlideType, len(_dga._eda))
	_dga.ContentTypes.AddOverride(_bgg, _gbg.SlideContentType)
	_gad := _ec.NewRelationships()
	_dga._gfd = append(_dga._gfd, _gad)
	_gad.AddAutoRelationship(_gbg.DocTypePresentation, _gbg.SlideType, len(_dga._gff), _gbg.SlideLayoutType)
	return Slide{_baa, _dfge, _dga, nil}
}

// ExtractText returns text from a slide as a SlideText object.
func (_adb *Slide) ExtractText() *SlideText {
	_ce := _dd(_adb._gac, _adb._gfgd.CSld.SpTree.Choice, []rectangle{}, []*TextItem{})
	_d.Sort(sort2d(_ce))
	return &SlideText{Items: _ce}
}

// SetOffsetX sets horizontal offset of text box in distance units (see measurement package).
func (_faab TextBox) SetOffsetX(offX float64) {
	_ggbe := _faab.getOff()
	_ebfda := _cb.ToEMU(offX)
	_ggbe.XAttr = _eef.ST_Coordinate{ST_CoordinateUnqualified: &_ebfda}
}

// SorterViewPr returns the SorterViewPr property.
func (_adc ViewProperties) SorterViewPr() *_a.CT_SlideSorterViewProperties {
	return _adc._ebacf.SorterViewPr
}

// SlideLayouts returns the slide layouts defined in the presentation.
func (_ebf *Presentation) SlideLayouts() []SlideLayout {
	_ccd := []SlideLayout{}
	for _, _gdbg := range _ebf._gff {
		_ccd = append(_ccd, SlideLayout{_gdbg})
	}
	return _ccd
}

// SaveAsTemplate writes the presentation out to a writer in the Zip package format as a template
func (_bdg *Presentation) SaveAsTemplate(w _cdg.Writer) error { return _bdg.save(w, true) }

// Text returns text from a slide as one string separated with line breaks.
func (_bd *SlideText) Text() string {
	_fcc := _db.NewBuffer([]byte{})
	for _, _adafa := range _bd.Items {
		if _adafa.Text != "" {
			_fcc.WriteString(_adafa.Text)
			_fcc.WriteString("\u000a")
		}
	}
	return _fcc.String()
}

// SlideSize returns presentation slide size.
func (_becf *Presentation) SlideSize() SlideSize {
	if _becf._dac.SldSz == nil {
		_becf._dac.SldSz = _a.NewCT_SlideSize()
	}
	return SlideSize{_becf._dac.SldSz, _becf}
}

// SlideSize represents a slide size of a presentation.
type SlideSize struct {
	_deg  *_a.CT_SlideSize
	_dbdb *Presentation
}

func (_fef *Presentation) nextSlideID() uint32 {
	_cbd := uint32(256)
	for _, _bcdg := range _fef._dac.SldIdLst.SldId {
		if _bcdg.IdAttr >= _cbd {
			_cbd = _bcdg.IdAttr + 1
		}
	}
	return _cbd
}

// Type returns the placeholder type
func (_efb PlaceHolder) Type() _a.ST_PlaceholderType { return _efb._beb.NvSpPr.NvPr.Ph.TypeAttr }

// New initializes and reurns a new presentation
func New() *Presentation {
	_aea := _dfb()
	_aea.ContentTypes.AddOverride("/\u0070\u0070\u0074\u002fpr\u0065s\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002d\u006ff\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u006d\u006c\u002e\u0070\u0072\u0065\u0073\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e\u002e\u006d\u0061\u0069\u006e\u002b\u0078\u006d\u006c")
	_aea.Rels.AddRelationship("\u0064\u006f\u0063\u0050\u0072\u006f\u0070\u0073\u002f\u0063\u006f\u0072e\u002e\u0078\u006d\u006c", "\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061s\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006ba\u0067\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u002f\u006d\u0065\u0074\u0061\u0064\u0061\u0074\u0061/\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006f\u0070e\u0072\u0074i\u0065\u0073")
	_aea.Rels.AddRelationship("\u0064\u006fc\u0050\u0072\u006fp\u0073\u002f\u0061\u0070\u0070\u002e\u0078\u006d\u006c", "\u0068t\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002eo\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073.\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0073\u0068\u0069p\u0073\u002f\u0065x\u0074\u0065\u006e\u0064\u0065d\u002d\u0070\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073")
	_aea.Rels.AddRelationship("p\u0070t\u002f\u0070\u0072\u0065\u0073\u0065\u006e\u0074a\u0074\u0069\u006f\u006e.x\u006d\u006c", "\u0068\u0074\u0074\u0070\u003a\u002f\u002fs\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073\u002e\u006f\u0072g\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006fc\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074")
	_aea.Rels.AddRelationship("\u0070\u0070\u0074\u002f\u0070\u0072\u0065\u0073\u0050\u0072\u006f\u0070s\u002e\u0078\u006d\u006c", "ht\u0074\u0070\u003a\u002f\u002f\u0073\u0063he\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006et\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068i\u0070s\u002f\u0070\u0072\u0065\u0073\u0050\u0072\u006f\u0070\u0073")
	_aea.Rels.AddRelationship("\u0070\u0070\u0074\u002f\u0076\u0069\u0065\u0077\u0050\u0072\u006f\u0070s\u002e\u0078\u006d\u006c", "ht\u0074\u0070\u003a\u002f\u002f\u0073\u0063he\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006da\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006et\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068i\u0070s\u002f\u0076\u0069\u0065\u0077\u0050\u0072\u006f\u0070\u0073")
	_aea.Rels.AddRelationship("\u0070\u0070\u0074\u002fta\u0062\u006c\u0065\u0053\u0074\u0079\u006c\u0065\u0073\u002e\u0078\u006d\u006c", "\u0068\u0074\u0074\u0070\u003a\u002f\u002fs\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006cf\u006fr\u006d\u0061\u0074\u0073\u002e\u006fr\u0067\u002f\u006f\u0066\u0066\u0069\u0063e\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073\u002f\u0074\u0061\u0062\u006c\u0065\u0053\u0074\u0079\u006ce\u0073")
	_aea._dac.SldMasterIdLst = _a.NewCT_SlideMasterIdList()
	_edb := _a.NewSldMaster()
	_edb.ClrMap.Bg1Attr = _eef.ST_ColorSchemeIndexLt1
	_edb.ClrMap.Bg2Attr = _eef.ST_ColorSchemeIndexLt2
	_edb.ClrMap.Tx1Attr = _eef.ST_ColorSchemeIndexDk1
	_edb.ClrMap.Tx2Attr = _eef.ST_ColorSchemeIndexDk2
	_edb.ClrMap.Accent1Attr = _eef.ST_ColorSchemeIndexAccent1
	_edb.ClrMap.Accent2Attr = _eef.ST_ColorSchemeIndexAccent2
	_edb.ClrMap.Accent3Attr = _eef.ST_ColorSchemeIndexAccent3
	_edb.ClrMap.Accent4Attr = _eef.ST_ColorSchemeIndexAccent4
	_edb.ClrMap.Accent5Attr = _eef.ST_ColorSchemeIndexAccent5
	_edb.ClrMap.Accent6Attr = _eef.ST_ColorSchemeIndexAccent6
	_edb.ClrMap.HlinkAttr = _eef.ST_ColorSchemeIndexHlink
	_edb.ClrMap.FolHlinkAttr = _eef.ST_ColorSchemeIndexFolHlink
	_aea._bda = append(_aea._bda, _edb)
	_cce := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.SlideMasterType, 1)
	_aea.ContentTypes.AddOverride(_cce, _gbg.SlideMasterContentType)
	_af := _aea._bfe.AddAutoRelationship(_gbg.DocTypePresentation, _gbg.OfficeDocumentType, 1, _gbg.SlideMasterType)
	_fcg := _a.NewCT_SlideMasterIdListEntry()
	_fcg.IdAttr = _gbg.Uint32(2147483648)
	_fcg.RIdAttr = _af.ID()
	_aea._dac.SldMasterIdLst.SldMasterId = append(_aea._dac.SldMasterIdLst.SldMasterId, _fcg)
	_ffe := _ec.NewRelationships()
	_aea._gdbf = append(_aea._gdbf, _ffe)
	_decc := _a.NewSldLayout()
	_aaga := _ffe.AddAutoRelationship(_gbg.DocTypePresentation, _gbg.SlideMasterType, 1, _gbg.SlideLayoutType)
	_cde := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.SlideLayoutType, 1)
	_aea.ContentTypes.AddOverride(_cde, _gbg.SlideLayoutContentType)
	_ffe.AddAutoRelationship(_gbg.DocTypePresentation, _gbg.SlideMasterType, 1, _gbg.ThemeType)
	_aea._gff = append(_aea._gff, _decc)
	_edb.SldLayoutIdLst = _a.NewCT_SlideLayoutIdList()
	_aef := _a.NewCT_SlideLayoutIdListEntry()
	_aef.IdAttr = _gbg.Uint32(2147483649)
	_aef.RIdAttr = _aaga.ID()
	_edb.SldLayoutIdLst.SldLayoutId = append(_edb.SldLayoutIdLst.SldLayoutId, _aef)
	_cag := _ec.NewRelationships()
	_aea._dfa = append(_aea._dfa, _cag)
	_cag.AddAutoRelationship(_gbg.DocTypePresentation, _gbg.SlideType, 1, _gbg.SlideMasterType)
	_aea._dac.NotesSz.CxAttr = 6858000
	_aea._dac.NotesSz.CyAttr = 9144000
	_bbc := _eef.NewTheme()
	_bbc.NameAttr = _gbg.String("\u0075n\u0069o\u0066\u0066\u0069\u0063\u0065\u0020\u0054\u0068\u0065\u006d\u0065")
	_bbc.ThemeElements.ClrScheme.NameAttr = "\u004f\u0066\u0066\u0069\u0063\u0065"
	_bbc.ThemeElements.ClrScheme.Dk1.SysClr = _eef.NewCT_SystemColor()
	_bbc.ThemeElements.ClrScheme.Dk1.SysClr.LastClrAttr = _gbg.String("\u0030\u0030\u0030\u0030\u0030\u0030")
	_bbc.ThemeElements.ClrScheme.Dk1.SysClr.ValAttr = _eef.ST_SystemColorValWindowText
	_bbc.ThemeElements.ClrScheme.Lt1.SysClr = _eef.NewCT_SystemColor()
	_bbc.ThemeElements.ClrScheme.Lt1.SysClr.LastClrAttr = _gbg.String("\u0066\u0066\u0066\u0066\u0066\u0066")
	_bbc.ThemeElements.ClrScheme.Lt1.SysClr.ValAttr = _eef.ST_SystemColorValWindow
	_bbc.ThemeElements.ClrScheme.Dk2.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.Dk2.SrgbClr.ValAttr = "\u0034\u0034\u0035\u0034\u0036\u0061"
	_bbc.ThemeElements.ClrScheme.Lt2.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.Lt2.SrgbClr.ValAttr = "\u0065\u0037\u0065\u0037\u0065\u0036"
	_bbc.ThemeElements.ClrScheme.Accent1.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.Accent1.SrgbClr.ValAttr = "\u0034\u0034\u0037\u0032\u0063\u0034"
	_bbc.ThemeElements.ClrScheme.Accent2.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.Accent2.SrgbClr.ValAttr = "\u0065\u0064\u0037\u0064\u0033\u0031"
	_bbc.ThemeElements.ClrScheme.Accent3.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.Accent3.SrgbClr.ValAttr = "\u0061\u0035\u0061\u0035\u0061\u0035"
	_bbc.ThemeElements.ClrScheme.Accent4.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.Accent4.SrgbClr.ValAttr = "\u0066\u0066\u0063\u0030\u0030\u0030"
	_bbc.ThemeElements.ClrScheme.Accent5.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.Accent5.SrgbClr.ValAttr = "\u0035\u0062\u0039\u0062\u0064\u0035"
	_bbc.ThemeElements.ClrScheme.Accent6.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.Accent6.SrgbClr.ValAttr = "\u0037\u0030\u0061\u0064\u0034\u0037"
	_bbc.ThemeElements.ClrScheme.Hlink.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.Hlink.SrgbClr.ValAttr = "\u0030\u0035\u0036\u0033\u0063\u0031"
	_bbc.ThemeElements.ClrScheme.FolHlink.SrgbClr = _eef.NewCT_SRgbColor()
	_bbc.ThemeElements.ClrScheme.FolHlink.SrgbClr.ValAttr = "\u0039\u0035\u0034\u0066\u0037\u0032"
	_bbc.ThemeElements.FontScheme.NameAttr = "\u004f\u0066\u0066\u0069\u0063\u0065"
	_bbc.ThemeElements.FontScheme.MajorFont.Latin.TypefaceAttr = "\u0043\u0061\u006c\u0069\u0062\u0072\u0069\u0020\u004c\u0069\u0067\u0068\u0074"
	_bbc.ThemeElements.FontScheme.MinorFont.Latin.TypefaceAttr = "\u0043a\u006c\u0069\u0062\u0072\u0069"
	_bbc.ThemeElements.FmtScheme.NameAttr = _gbg.String("\u004f\u0066\u0066\u0069\u0063\u0065")
	_gefa := _eef.NewEG_FillProperties()
	_bbc.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties = append(_bbc.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties, _gefa)
	_gefa.SolidFill = &_eef.CT_SolidColorFillProperties{SchemeClr: &_eef.CT_SchemeColor{ValAttr: _eef.ST_SchemeColorValPhClr}}
	_gefa = _eef.NewEG_FillProperties()
	_bbc.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties = append(_bbc.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties, _gefa)
	_bbc.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties = append(_bbc.ThemeElements.FmtScheme.FillStyleLst.EG_FillProperties, _gefa)
	_gefa.GradFill = &_eef.CT_GradientFillProperties{RotWithShapeAttr: _gbg.Bool(true), GsLst: &_eef.CT_GradientStopList{}, Lin: &_eef.CT_LinearShadeProperties{}}
	_gefa.GradFill.Lin.AngAttr = _gbg.Int32(5400000)
	_gefa.GradFill.Lin.ScaledAttr = _gbg.Bool(false)
	_edg := _eef.NewCT_GradientStop()
	_edg.PosAttr.ST_PositiveFixedPercentageDecimal = _gbg.Int32(0)
	_edg.SchemeClr = &_eef.CT_SchemeColor{ValAttr: _eef.ST_SchemeColorValPhClr}
	_gefa.GradFill.GsLst.Gs = append(_gefa.GradFill.GsLst.Gs, _edg)
	_edg = _eef.NewCT_GradientStop()
	_edg.PosAttr.ST_PositiveFixedPercentageDecimal = _gbg.Int32(50000)
	_edg.SchemeClr = &_eef.CT_SchemeColor{ValAttr: _eef.ST_SchemeColorValPhClr}
	_gefa.GradFill.GsLst.Gs = append(_gefa.GradFill.GsLst.Gs, _edg)
	_bbc.ThemeElements.FmtScheme.LnStyleLst = _eef.NewCT_LineStyleList()
	for _bfa := 0; _bfa < 3; _bfa++ {
		_fddf := _eef.NewCT_LineProperties()
		_fddf.WAttr = _gbg.Int32(int32(6350 * (_bfa + 1)))
		_fddf.CapAttr = _eef.ST_LineCapFlat
		_fddf.CmpdAttr = _eef.ST_CompoundLineSng
		_fddf.AlgnAttr = _eef.ST_PenAlignmentCtr
		_bbc.ThemeElements.FmtScheme.LnStyleLst.Ln = append(_bbc.ThemeElements.FmtScheme.LnStyleLst.Ln, _fddf)
	}
	_bbc.ThemeElements.FmtScheme.EffectStyleLst = _eef.NewCT_EffectStyleList()
	for _efdd := 0; _efdd < 3; _efdd++ {
		_cbae := _eef.NewCT_EffectStyleItem()
		_cbae.EffectLst = _eef.NewCT_EffectList()
		_bbc.ThemeElements.FmtScheme.EffectStyleLst.EffectStyle = append(_bbc.ThemeElements.FmtScheme.EffectStyleLst.EffectStyle, _cbae)
	}
	_baeb := _eef.NewEG_FillProperties()
	_baeb.SolidFill = &_eef.CT_SolidColorFillProperties{SchemeClr: &_eef.CT_SchemeColor{ValAttr: _eef.ST_SchemeColorValPhClr}}
	_bbc.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties = append(_bbc.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties, _baeb)
	_bbc.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties = append(_bbc.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties, _baeb)
	_bbc.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties = append(_bbc.ThemeElements.FmtScheme.BgFillStyleLst.EG_FillProperties, _gefa)
	_aea._eade = append(_aea._eade, _bbc)
	_fcec := _gbg.AbsoluteFilename(_gbg.DocTypePresentation, _gbg.ThemeType, 1)
	_aea.ContentTypes.AddOverride(_fcec, _gbg.ThemeContentType)
	_aea._bfe.AddAutoRelationship(_gbg.DocTypePresentation, _gbg.OfficeDocumentType, 1, _gbg.ThemeType)
	_ceg := _ec.NewRelationships()
	_aea._cdf = append(_aea._cdf, _ceg)
	return _aea
}

const _cdd float64 = 500000

// ClrMru returns the ClrMru property.
func (_fed PresentationProperties) ClrMru() *_eef.CT_ColorMRU { return _fed._ddad.ClrMru }

// SlideViewPr returns the SlideViewPr property.
func (_adab ViewProperties) SlideViewPr() *_a.CT_SlideViewProperties { return _adab._ebacf.SlideViewPr }

// GetChartSpaceByRelId returns a *crt.ChartSpace with the associated relation ID in the
// slide.
func (_cecd *Slide) GetChartSpaceByRelId(relId string) *_f.ChartSpace {
	_cac := _cecd.getSlideRels()
	if (_cac == _ec.Relationships{}) {
		return nil
	}
	_fbda := _cac.GetTargetByRelId(relId)
	for _, _abbc := range _cecd._gac._bfg {
		if _fbda == _abbc.Target() {
			return _abbc._gd
		}
	}
	return nil
}

// NotesViewPr returns the NotesViewPr property.
func (_agb ViewProperties) NotesViewPr() *_a.CT_NotesViewProperties { return _agb._ebacf.NotesViewPr }

type chart struct {
	_gd  *_f.ChartSpace
	_dbe string
	_gc  string
}

// GetImageByRelID returns an ImageRef with the associated relation ID in the
// slide.
func (_cfbg *Slide) GetImageByRelID(relID string) (_ec.ImageRef, bool) {
	_gdgd := _cfbg.getSlideRels()
	if (_gdgd == _ec.Relationships{}) {
		return _ec.ImageRef{}, false
	}
	var _fcag string
	for _, _abc := range _gdgd.X().Relationship {
		if _abc.IdAttr == relID {
			_fcag = _abc.TargetAttr
			break
		}
	}
	for _, _fdfd := range _cfbg._gac.Images {
		if _fdfd.Target() == _fcag {
			return _fdfd, true
		}
	}
	return _ec.ImageRef{}, false
}
func (_ad *chart) RelId() string { return _ad._dbe }

// PlaceHolder is a place holder from a slide.
type PlaceHolder struct {
	_beb  *_a.CT_Shape
	_cebf *_a.Sld
}

// ViewProperties contains presentation specific properties.
type ViewProperties struct{ _ebacf *_a.ViewPr }

// SlideLayout is a layout from which slides can be created.
type SlideLayout struct{ _dcbe *_a.SldLayout }

// AddParagraph adds a new paragraph to a placeholder.
func (_dgfb PlaceHolder) AddParagraph() _gg.Paragraph {
	_cdc := _gg.MakeParagraph(_eef.NewCT_TextParagraph())
	_dgfb._beb.TxBody.P = append(_dgfb._beb.TxBody.P, _cdc.X())
	return _cdc
}

// Slides returns the slides in the presentation.
func (_edd *Presentation) Slides() []Slide {
	_deea := []Slide{}
	for _cfcf, _afcc := range _edd._eda {
		_deea = append(_deea, Slide{_edd._dac.SldIdLst.SldId[_cfcf], _afcc, _edd, nil})
	}
	return _deea
}
