package common

import (
	_bd "archive/zip"
	_f "bytes"
	_e "encoding/xml"
	_d "errors"
	_bde "fmt"
	_ea "github.com/arcpd/msword"
	_ge "github.com/arcpd/msword/common/logger"
	_eg "github.com/arcpd/msword/common/tempstorage"
	_gg "github.com/arcpd/msword/common/tempstorage/diskstore"
	_gf "github.com/arcpd/msword/measurement"
	_bdd "github.com/arcpd/msword/schema/soo/dml"
	_ga "github.com/arcpd/msword/schema/soo/ofc/custom_properties"
	_de "github.com/arcpd/msword/schema/soo/ofc/docPropsVTypes"
	_ad "github.com/arcpd/msword/schema/soo/ofc/extended_properties"
	_fc "github.com/arcpd/msword/schema/soo/pkg/content_types"
	_ag "github.com/arcpd/msword/schema/soo/pkg/metadata/core_properties"
	_ac "github.com/arcpd/msword/schema/soo/pkg/relationships"
	_dg "github.com/arcpd/msword/zippkg"
	_cee "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	_ce "os"
	_a "reflect"
	_dc "regexp"
	_fe "strconv"
	_ca "strings"
	_gb "time"
)

// AddExtraFileFromZip is used when reading an unsupported file from an OOXML
// file. This ensures that unsupported file content will at least round-trip
// correctly.
func (_cda *DocBase) AddExtraFileFromZip(f *_bd.File) error {
	_gace, _agff := _dg.ExtractToDiskTmp(f, _cda.TmpPath)
	if _agff != nil {
		return _bde.Errorf("\u0065\u0072r\u006f\u0072\u0020\u0065x\u0074\u0072a\u0063\u0074\u0069\u006e\u0067\u0020\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u0069\u006ce\u003a\u0020\u0025\u0073", _agff)
	}
	_cda.ExtraFiles = append(_cda.ExtraFiles, ExtraFile{ZipPath: f.Name, DiskPath: _gace})
	return nil
}

// X returns the underlying raw XML data.
func (_cdge Relationships) X() *_ac.Relationships { return _cdge._dafb }

// ImageRef is a reference to an image within a document.
type ImageRef struct {
	_aed *DocBase
	_bae Relationships
	_fcb Image
	_gfd string
	_aae string
}

// NewRelationshipsCopy creates a new relationships wrapper as a copy of passed in instance.
func NewRelationshipsCopy(rels Relationships) Relationships {
	_ddcc := *rels._dafb
	return Relationships{_dafb: &_ddcc}
}

func (_fgg CustomProperties) SetPropertyAsError(name string, error string) {
	_deec := _fgg.getNewProperty(name)
	_deec.Error = &error
	_fgg.setOrReplaceProperty(_deec)
}

// NewCoreProperties constructs a new CoreProperties.
func NewCoreProperties() CoreProperties { return CoreProperties{_ae: _ag.NewCoreProperties()} }

// AddImageToZip adds an image (either from bytes or from disk) and adds it to the zip file.
func AddImageToZip(z *_bd.Writer, img ImageRef, imageNum int, dt _ea.DocType) error {
	_cbd := _ea.AbsoluteImageFilename(dt, imageNum, _ca.ToLower(img.Format()))
	if img.Data() != nil && len(*img.Data()) > 0 {
		if _ffed := _dg.AddFileFromBytes(z, _cbd, *img.Data()); _ffed != nil {
			return _ffed
		}
	} else if img.Path() != "" {
		if _dbgf := _dg.AddFileFromDisk(z, _cbd, img.Path()); _dbgf != nil {
			return _dbgf
		}
	} else {
		return _bde.Errorf("\u0075\u006es\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0073\u006f\u0075\u0072\u0063\u0065\u003a %\u002b\u0076", img)
	}
	return nil
}

// Properties returns table properties.
func (_fccd Table) Properties() *_bdd.CT_TableProperties { return _fccd._bef.TblPr }

// X returns the inner wrapped XML type.
func (_dfef CustomProperties) X() *_ga.Properties { return _dfef._gbg }

func (_ebg CustomProperties) SetPropertyAsFiletime(name string, filetime _gb.Time) {
	_gda := _ebg.getNewProperty(name)
	_gda.Filetime = &filetime
	_ebg.setOrReplaceProperty(_gda)
}

// ImageFromBytes returns an Image struct for an in-memory image. You can also
// construct an Image directly if the file and size are known.
func ImageFromBytes(data []byte) (Image, error) {
	_dcdd := Image{}
	_agge, _egcc, _beae := _cee.Decode(_f.NewReader(data))
	if _beae != nil {
		return _dcdd, _bde.Errorf("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s", _beae)
	}
	_dcdd.Data = &data
	_dcdd.Format = _egcc
	_dcdd.Size = _agge.Bounds().Size()
	return _dcdd, nil
}

func (_cdd CustomProperties) SetPropertyAsLpstr(name string, lpstr string) {
	_eae := _cdd.getNewProperty(name)
	_eae.Lpstr = &lpstr
	_cdd.setOrReplaceProperty(_eae)
}

// Append appends DocBase part of an office document to another DocBase.
func (_dga DocBase) Append(docBase1 DocBase) DocBase {
	_fb := docBase1.ContentTypes.X()
	for _, _ec := range _fb.Default {
		_dga.ContentTypes.AddDefault(_ec.ExtensionAttr, _ec.ContentTypeAttr)
	}
	for _, _ggb := range _fb.Override {
		_dga.ContentTypes.AddOverride(_ggb.PartNameAttr, _ggb.ContentTypeAttr)
	}
	_cc := _dga.AppProperties.X()
	_acd := docBase1.AppProperties.X()
	if _cc.Pages != nil {
		if _acd.Pages != nil {
			*_cc.Pages += *_acd.Pages
		}
	} else if _acd.Pages != nil {
		_cc.Pages = _acd.Pages
	}
	if _cc.Words != nil {
		if _acd.Words != nil {
			*_cc.Words += *_acd.Words
		}
	} else if _acd.Words != nil {
		_cc.Words = _acd.Words
	}
	if _cc.Characters != nil {
		if _acd.Characters != nil {
			*_cc.Characters += *_acd.Characters
		}
	} else if _acd.Characters != nil {
		_cc.Characters = _acd.Characters
	}
	if _cc.Lines != nil {
		if _acd.Lines != nil {
			*_cc.Lines += *_acd.Lines
		}
	} else if _acd.Lines != nil {
		_cc.Lines = _acd.Lines
	}
	if _cc.Paragraphs != nil {
		if _acd.Paragraphs != nil {
			*_cc.Paragraphs += *_acd.Paragraphs
		}
	} else if _acd.Paragraphs != nil {
		_cc.Paragraphs = _acd.Paragraphs
	}
	if _cc.Notes != nil {
		if _acd.Notes != nil {
			*_cc.Notes += *_acd.Notes
		}
	} else if _acd.Notes != nil {
		_cc.Notes = _acd.Notes
	}
	if _cc.HiddenSlides != nil {
		if _acd.HiddenSlides != nil {
			*_cc.HiddenSlides += *_acd.HiddenSlides
		}
	} else if _acd.HiddenSlides != nil {
		_cc.HiddenSlides = _acd.HiddenSlides
	}
	if _cc.MMClips != nil {
		if _acd.MMClips != nil {
			*_cc.MMClips += *_acd.MMClips
		}
	} else if _acd.MMClips != nil {
		_cc.MMClips = _acd.MMClips
	}
	if _cc.LinksUpToDate != nil {
		if _acd.LinksUpToDate != nil {
			*_cc.LinksUpToDate = *_cc.LinksUpToDate && *_acd.LinksUpToDate
		}
	} else if _acd.LinksUpToDate != nil {
		_cc.LinksUpToDate = _acd.LinksUpToDate
	}
	if _cc.CharactersWithSpaces != nil {
		if _acd.CharactersWithSpaces != nil {
			*_cc.CharactersWithSpaces += *_acd.CharactersWithSpaces
		}
	} else if _acd.CharactersWithSpaces != nil {
		_cc.CharactersWithSpaces = _acd.CharactersWithSpaces
	}
	if _cc.SharedDoc != nil {
		if _acd.SharedDoc != nil {
			*_cc.SharedDoc = *_cc.SharedDoc || *_acd.SharedDoc
		}
	} else if _acd.SharedDoc != nil {
		_cc.SharedDoc = _acd.SharedDoc
	}
	if _cc.HyperlinksChanged != nil {
		if _acd.HyperlinksChanged != nil {
			*_cc.HyperlinksChanged = *_cc.HyperlinksChanged || *_acd.HyperlinksChanged
		}
	} else if _acd.HyperlinksChanged != nil {
		_cc.HyperlinksChanged = _acd.HyperlinksChanged
	}
	_cc.DigSig = nil
	if _cc.TitlesOfParts == nil && _acd.TitlesOfParts != nil {
		_cc.TitlesOfParts = _acd.TitlesOfParts
	}
	if _cc.HeadingPairs != nil {
		if _acd.HeadingPairs != nil {
			_fa := _cc.HeadingPairs.Vector
			_cd := _acd.HeadingPairs.Vector
			_cad := _fa.Variant
			_aa := _cd.Variant
			_ggd := []*_de.Variant{}
			for _gdb := 0; _gdb < len(_aa); _gdb += 2 {
				_bb := _aa[_gdb].Lpstr
				_fce := false
				for _ded := 0; _ded < len(_cad); _ded += 2 {
					_ff := _cad[_ded].Lpstr
					if _ff != nil && _bb != nil && *_ff == *_bb {
						*_cad[_ded+1].I4 = *_cad[_ded+1].I4 + *_aa[_gdb+1].I4
						_fce = true
						break
					}
				}
				if !_fce {
					_ggd = append(_ggd, &_de.Variant{CT_Variant: _de.CT_Variant{Lpstr: _aa[_gdb].Lpstr}})
					_ggd = append(_ggd, &_de.Variant{CT_Variant: _de.CT_Variant{I4: _aa[_gdb].I4}})
				}
			}
			_cad = append(_cad, _ggd...)
			_fa.SizeAttr = uint32(len(_cad))
		}
	} else if _acd.HeadingPairs != nil {
		_cc.HeadingPairs = _acd.HeadingPairs
	}
	if _cc.HLinks != nil {
		if _acd.HLinks != nil {
			_bbb := _cc.HLinks.Vector
			_aac := _acd.HLinks.Vector
			_agf := _bbb.Variant
			_da := _aac.Variant
			for _, _gba := range _da {
				_bac := true
				for _, _df := range _agf {
					if _a.DeepEqual(_df, _gba) {
						_bac = false
						break
					}
				}
				if _bac {
					_agf = append(_agf, _gba)
					_bbb.SizeAttr++
				}
			}
		}
	} else if _acd.HLinks != nil {
		_cc.HLinks = _acd.HLinks
	}
	_fef := _dga.GetOrCreateCustomProperties()
	_ecc := docBase1.GetOrCreateCustomProperties()
	for _, _cb := range _ecc.PropertiesList() {
		_fef.setProperty(_cb)
	}
	_dga.CustomProperties = _fef
	_ffc := _dga.Rels.X().Relationship
	for _, _bcc := range docBase1.Rels.X().Relationship {
		_acf := true
		for _, _ccb := range _ffc {
			if _ccb.TargetAttr == _bcc.TargetAttr && _ccb.TypeAttr == _bcc.TypeAttr {
				_acf = false
				break
			}
		}
		if _acf {
			_dga.Rels.AddRelationship(_bcc.TargetAttr, _bcc.TypeAttr)
		}
	}
	for _, _gde := range docBase1.ExtraFiles {
		_cdg := _gde.ZipPath
		_cf := true
		for _, _ab := range _dga.ExtraFiles {
			if _ab.ZipPath == _cdg {
				_cf = false
				break
			}
		}
		if _cf {
			_dga.ExtraFiles = append(_dga.ExtraFiles, _gde)
		}
	}
	return _dga
}

// RemoveOverride removes an override given a path.
func (_dfd ContentTypes) RemoveOverride(path string) {
	if !_ca.HasPrefix(path, "\u002f") {
		path = "\u002f" + path
	}
	for _egg, _aaa := range _dfd._abd.Override {
		if _aaa.PartNameAttr == path {
			copy(_dfd._abd.Override[_egg:], _dfd._abd.Override[_egg+1:])
			_dfd._abd.Override = _dfd._abd.Override[0 : len(_dfd._abd.Override)-1]
		}
	}
}

// GetTargetByRelId returns a target path with the associated relation ID.
func (_edf Relationships) GetTargetByRelId(idAttr string) string {
	for _, _gdab := range _edf._dafb.Relationship {
		if _gdab.IdAttr == idAttr {
			return _gdab.TargetAttr
		}
	}
	return ""
}

// CreateCustomProperties creates the custom properties of the document.
func (_fed *DocBase) CreateCustomProperties() {
	_fed.CustomProperties = NewCustomProperties()
	_fed.AddCustomRelationships()
}

// NewTableWithXfrm makes a new table with a pointer to its parent Xfrm for changing its offset and size.
func NewTableWithXfrm(xfrm *_bdd.CT_Transform2D) *Table {
	_cef := _bdd.NewTbl()
	_cef.TblPr = _bdd.NewCT_TableProperties()
	return &Table{_bef: _cef, _dgcc: xfrm}
}

func (_gga CustomProperties) SetPropertyAsR8(name string, r8 float64) {
	_bcd := _gga.getNewProperty(name)
	_bcd.R8 = &r8
	_gga.setOrReplaceProperty(_bcd)
}

// X returns the inner wrapped XML type.
func (_gce TableStyles) X() *_bdd.TblStyleLst { return _gce._debd }

// ID returns the ID of a relationship.
func (_bcdf Relationship) ID() string { return _bcdf._fdg.IdAttr }

// SetCreated sets the time that the document was created.
func (_aff CoreProperties) SetCreated(t _gb.Time) {
	_aff._ae.Created = _agca(t, "\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064")
}

// GetByRelId returns a relationship with the associated relation ID.
func (_dge Relationships) GetByRelId(idAttr string) Relationship {
	for _, _dabe := range _dge._dafb.Relationship {
		if _dabe.IdAttr == idAttr {
			return Relationship{_fdg: _dabe}
		}
	}
	return Relationship{}
}

// Clear removes any existing relationships.
func (_bffc Relationships) Clear() { _bffc._dafb.Relationship = nil }

// SetDescription records the description of the document.
func (_ffe CoreProperties) SetDescription(s string) {
	if _ffe._ae.Description == nil {
		_ffe._ae.Description = &_ea.XSDAny{XMLName: _e.Name{Local: "\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}}
	}
	_ffe._ae.Description.Data = []byte(s)
}

func (_fbbd CustomProperties) SetPropertyAsDecimal(name string, decimal float64) {
	_dbc := _fbbd.getNewProperty(name)
	_dbc.Decimal = &decimal
	_fbbd.setOrReplaceProperty(_dbc)
}

const _cgfd = 30

func (_gbf CustomProperties) SetPropertyAsI1(name string, i1 int8) {
	_ddac := _gbf.getNewProperty(name)
	_ddac.I1 = &i1
	_gbf.setOrReplaceProperty(_ddac)
}

// MakeImageRef constructs an image reference which is a reference to a
// particular image file inside a document.  The same image can be used multiple
// times in a document by re-use the ImageRef.
func MakeImageRef(img Image, d *DocBase, rels Relationships) ImageRef {
	return ImageRef{_fcb: img, _aed: d, _bae: rels}
}

// Type returns the type of a relationship.
func (_edd Relationship) Type() string { return _edd._fdg.TypeAttr }

func (_fga CustomProperties) getNewProperty(_fda string) *_ga.CT_Property {
	_gec := _fga._gbg.Property
	_dacg := int32(1)
	for _, _dbf := range _gec {
		if _dbf.PidAttr > _dacg {
			_dacg = _dbf.PidAttr
		}
	}
	_aab := _ga.NewCT_Property()
	_aab.NameAttr = &_fda
	_aab.PidAttr = _dacg + 1
	_aab.FmtidAttr = "\u007b\u0044\u0035\u0043\u0044\u0044\u0035\u0030\u0035\u002d\u0032\u0045\u0039\u0043\u002d\u0031\u0030\u0031\u0042\u002d\u0039\u0033\u0039\u0037-\u0030\u0038\u0030\u0030\u0032B\u0032\u0043F\u0039\u0041\u0045\u007d"
	return _aab
}

// DocBase is the type embedded in in the Document/Workbook/Presentation types
// that contains members common to all.
type DocBase struct {
	ContentTypes     ContentTypes
	AppProperties    AppProperties
	Rels             Relationships
	CoreProperties   CoreProperties
	CustomProperties CustomProperties
	Thumbnail        _cee.Image
	Images           []ImageRef
	ExtraFiles       []ExtraFile
	TmpPath          string
}

// X returns the inner wrapped XML type.
func (_gdc Table) X() *_bdd.Tbl { return _gdc._bef }

// RelID returns the relationship ID.
func (_gab ImageRef) RelID() string { return _gab._gfd }

// LastModifiedBy returns the name of the last person to modify the document
func (_dbb CoreProperties) LastModifiedBy() string {
	if _dbb._ae.LastModifiedBy != nil {
		return *_dbb._ae.LastModifiedBy
	}
	return ""
}

// GetImageBytesByTarget returns Image object with Data bytes read from its target.
func (_cdf *DocBase) GetImageBytesByTarget(target string) (Image, error) {
	if target != "" {
		target = "\u0077\u006f\u0072d\u002f" + target
		for _, _bcde := range _cdf.Images {
			if _bcde.Target() == target {
				return ImageFromStorage(_bcde.Path())
			}
		}
	}
	return Image{}, _eeba
}

// Title returns the Title of the document
func (_efb CoreProperties) Title() string {
	if _efb._ae.Title != nil {
		return string(_efb._ae.Title.Data)
	}
	return ""
}

func (_cba CustomProperties) SetPropertyAsCy(name string, cy string) {
	_cbga := _cba.getNewProperty(name)
	_cbga.Cy = &cy
	_cba.setOrReplaceProperty(_cbga)
}

func (_fffc CustomProperties) SetPropertyAsArray(name string, array *_de.Array) {
	_faa := _fffc.getNewProperty(name)
	_faa.Array = array
	_fffc.setOrReplaceProperty(_faa)
}

func (_bfc CustomProperties) SetPropertyAsBstr(name string, bstr string) {
	_gdd := _bfc.getNewProperty(name)
	_gdd.Bstr = &bstr
	_bfc.setOrReplaceProperty(_gdd)
}

// Path returns the path to an image file, if any.
func (_abdd ImageRef) Path() string { return _abdd._fcb.Path }

// Created returns the time that the document was created.
func (_bf CoreProperties) Created() _gb.Time { return _dffc(_bf._ae.Created) }

// AddAutoRelationship adds a relationship with an automatically generated
// filename based off of the type. It should be preferred over AddRelationship
// to ensure consistent filenames are maintained.
func (_afe Relationships) AddAutoRelationship(dt _ea.DocType, src string, idx int, ctype string) Relationship {
	return _afe.AddRelationship(_ea.RelativeFilename(dt, src, ctype, idx), ctype)
}

func (_bdc CustomProperties) SetPropertyAsOstream(name string, ostream string) {
	_fcc := _bdc.getNewProperty(name)
	_fcc.Ostream = &ostream
	_bdc.setOrReplaceProperty(_fcc)
}

// X returns the inner raw content types.
func (_gcc ContentTypes) X() *_fc.Types { return _gcc._abd }

// AddCol adds a column to a table.
func (_gadd Table) AddCol() *TableCol {
	_gaf := _bdd.NewCT_TableCol()
	_gadd._bef.TblGrid.GridCol = append(_gadd._bef.TblGrid.GridCol, _gaf)
	for _, _ccbe := range _gadd._bef.Tr {
		_abgc := _bdd.NewCT_TableCell()
		_ccbe.Tc = append(_ccbe.Tc, _abgc)
	}
	return &TableCol{_eaba: _gaf}
}

func (_eba CustomProperties) SetPropertyAsLpwstr(name string, lpwstr string) {
	_ddc := _eba.getNewProperty(name)
	_ddc.Lpwstr = &lpwstr
	_eba.setOrReplaceProperty(_ddc)
}

// RelativeWidth returns the relative width of an image given a fixed height.
// This is used when setting image to a fixed height to calculate the width
// required to keep the same image aspect ratio.
func (_ddaa ImageRef) RelativeWidth(h _gf.Distance) _gf.Distance {
	_bbgc := float64(_ddaa.Size().X) / float64(_ddaa.Size().Y)
	return h * _gf.Distance(_bbgc)
}

// Category returns the category of the document
func (_dbd CoreProperties) Category() string {
	if _dbd._ae.Category != nil {
		return *_dbd._ae.Category
	}
	return ""
}

// AddRow adds a row to a table.
func (_ffcf Table) AddRow() *TableRow {
	_efc := _bdd.NewCT_TableRow()
	for _fbf := 0; _fbf < len(_ffcf._bef.TblGrid.GridCol); _fbf++ {
		_efc.Tc = append(_efc.Tc, _bdd.NewCT_TableCell())
	}
	_ffcf._bef.Tr = append(_ffcf._bef.Tr, _efc)
	return &TableRow{_cbde: _efc}
}

// SetDocSecurity sets the document security flag.
func (_fg AppProperties) SetDocSecurity(v int32) { _fg._agfd.DocSecurity = _ea.Int32(v) }

// SetCategory records the category of the document.
func (_cbg CoreProperties) SetCategory(s string) { _cbg._ae.Category = &s }

// Target returns the target attrubute of the image reference (a path where the image file is located in the document structure).
func (_fccf *ImageRef) Target() string { return _fccf._aae }

// Theme is a drawingml theme.
type Theme struct{ _acda *_bdd.Theme }

// CoreProperties contains document specific properties.
type CoreProperties struct{ _ae *_ag.CoreProperties }

// TableRow represents a row in a table.
type TableRow struct{ _cbde *_bdd.CT_TableRow }

const _face = 15

// CustomProperties contains document specific properties.
type CustomProperties struct{ _gbg *_ga.Properties }

// SetOffsetY sets vertical offset of a table in distance units (see measurement package).
func (_aadb Table) SetOffsetY(offY float64) {
	if _aadb._dgcc.Off == nil {
		_aadb._dgcc.Off = _bdd.NewCT_Point2D()
		_geag := int64(0)
		_aadb._dgcc.Off.XAttr = _bdd.ST_Coordinate{ST_CoordinateUnqualified: &_geag}
	}
	_eeea := _gf.ToEMU(offY)
	_aadb._dgcc.Off.YAttr = _bdd.ST_Coordinate{ST_CoordinateUnqualified: &_eeea}
}

// ContentTypes is the top level "[Content_Types].xml" in a zip package.
type ContentTypes struct{ _abd *_fc.Types }

// AddHyperlink adds an external hyperlink relationship.
func (_agfb Relationships) AddHyperlink(target string) Hyperlink {
	_ebf := _agfb.AddRelationship(target, _ea.HyperLinkType)
	_ebf._fdg.TargetModeAttr = _ac.ST_TargetModeExternal
	return Hyperlink(_ebf)
}

func (_cac CustomProperties) SetPropertyAsI8(name string, i8 int64) {
	_gded := _cac.getNewProperty(name)
	_gded.I8 = &i8
	_cac.setOrReplaceProperty(_gded)
}

func (_afa CustomProperties) SetPropertyAsUi1(name string, ui1 uint8) {
	_agce := _afa.getNewProperty(name)
	_agce.Ui1 = &ui1
	_afa.setOrReplaceProperty(_agce)
}

// SetModified sets the time that the document was modified.
func (_dfe CoreProperties) SetModified(t _gb.Time) {
	_dfe._ae.Modified = _agca(t, "\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064")
}

// GetPropertyByName returns a custom property selected by it's name.
func (_dda CustomProperties) GetPropertyByName(name string) CustomProperty {
	_deb := _dda._gbg.Property
	for _, _dee := range _deb {
		if *_dee.NameAttr == name {
			return CustomProperty{_dba: _dee}
		}
	}
	return CustomProperty{}
}

// SetTarget changes the target attribute of the image reference (e.g. in the case of the creation of the reference or if the image which the reference is related to was moved from one location to another).
func (_acb *ImageRef) SetTarget(target string) { _acb._aae = target }

// AddDefault registers a default content type for a given file extension.
func (_dff ContentTypes) AddDefault(fileExtension string, contentType string) {
	fileExtension = _ca.ToLower(fileExtension)
	for _, _dgaf := range _dff._abd.Default {
		if _dgaf.ExtensionAttr == fileExtension {
			return
		}
	}
	_gdea := _fc.NewDefault()
	_gdea.ExtensionAttr = fileExtension
	_gdea.ContentTypeAttr = contentType
	_dff._abd.Default = append(_dff._abd.Default, _gdea)
}

// Data returns the data of an image file, if any.
func (_bdg ImageRef) Data() *[]byte { return _bdg._fcb.Data }

// AddOverride adds an override content type for a given path name.
func (_ece ContentTypes) AddOverride(path, contentType string) {
	if !_ca.HasPrefix(path, "\u002f") {
		path = "\u002f" + path
	}
	if _ca.HasPrefix(contentType, "\u0068\u0074\u0074\u0070") {
		_ge.Log.Debug("\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0074\u0079p\u0065\u0020\u0027%\u0073\u0027\u0020\u0069\u0073\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002c m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0068\u0074\u0074\u0070", contentType)
	}
	for _, _dbg := range _ece._abd.Override {
		if _dbg.PartNameAttr == path {
			return
		}
	}
	_be := _fc.NewOverride()
	_be.PartNameAttr = path
	_be.ContentTypeAttr = contentType
	_ece._abd.Override = append(_ece._abd.Override, _be)
}

// SetTarget set the target (path) of a relationship.
func (_ege Relationship) SetTarget(s string) { _ege._fdg.TargetAttr = s }

// SetOffsetX sets horizontal offset of a table in distance units (see measurement package).
func (_fac Table) SetOffsetX(offX float64) {
	if _fac._dgcc.Off == nil {
		_fac._dgcc.Off = _bdd.NewCT_Point2D()
		_dad := int64(0)
		_fac._dgcc.Off.YAttr = _bdd.ST_Coordinate{ST_CoordinateUnqualified: &_dad}
	}
	_add := _gf.ToEMU(offX)
	_fac._dgcc.Off.XAttr = _bdd.ST_Coordinate{ST_CoordinateUnqualified: &_add}
}

const _bbg = "2\u00300\u0036\u002d\u0030\u0031\u002d\u0030\u0032\u00541\u0035\u003a\u0030\u0034:0\u0035\u005a"

func (_feaf CustomProperties) SetPropertyAsStream(name string, stream string) {
	_bff := _feaf.getNewProperty(name)
	_bff.Stream = &stream
	_feaf.setOrReplaceProperty(_bff)
}

// SetContentStatus records the content status of the document.
func (_dac CoreProperties) SetContentStatus(s string) { _dac._ae.ContentStatus = &s }

// RelativeHeight returns the relative height of an image given a fixed width.
// This is used when setting image to a fixed width to calculate the height
// required to keep the same image aspect ratio.
func (_ffce ImageRef) RelativeHeight(w _gf.Distance) _gf.Distance {
	_fag := float64(_ffce.Size().Y) / float64(_ffce.Size().X)
	return w * _gf.Distance(_fag)
}

// SetApplication sets the name of the application that created the document.
func (_eed AppProperties) SetApplication(s string) { _eed._agfd.Application = &s }

// FindRIDForN returns the relationship ID for the i'th relationship of type t.
func (_ggfa Relationships) FindRIDForN(i int, t string) string {
	for _, _fgb := range _ggfa._dafb.CT_Relationships.Relationship {
		if _fgb.TypeAttr == t {
			if i == 0 {
				return _fgb.IdAttr
			}
			i--
		}
	}
	return ""
}

var ReleasedAt = _gb.Date(_baeb, _ccc, _ggdg, _face, _cgfd, 0, 0, _gb.UTC)

// EnsureDefault esnures that an extension and default content type exist,
// adding it if necessary.
func (_gdbg ContentTypes) EnsureDefault(ext, contentType string) {
	ext = _ca.ToLower(ext)
	for _, _ed := range _gdbg._abd.Default {
		if _ed.ExtensionAttr == ext {
			_ed.ContentTypeAttr = contentType
			return
		}
	}
	_ggf := &_fc.Default{}
	_ggf.ContentTypeAttr = contentType
	_ggf.ExtensionAttr = ext
	_gdbg._abd.Default = append(_gdbg._abd.Default, _ggf)
}

// Remove removes an existing relationship.
func (_dgb Relationships) Remove(rel Relationship) bool {
	for _eee, _aad := range _dgb._dafb.Relationship {
		if _aad == rel._fdg {
			copy(_dgb._dafb.Relationship[_eee:], _dgb._dafb.Relationship[_eee+1:])
			_dgb._dafb.Relationship = _dgb._dafb.Relationship[0 : len(_dgb._dafb.Relationship)-1]
			return true
		}
	}
	return false
}

// Rows returns all table rows.
func (_bdab Table) Rows() []*TableRow {
	_dgeb := _bdab._bef.Tr
	_cgd := []*TableRow{}
	for _, _bgb := range _dgeb {
		_cgd = append(_cgd, &TableRow{_cbde: _bgb})
	}
	return _cgd
}

// RemoveOverrideByIndex removes an override given a path and override index.
func (_gdec ContentTypes) RemoveOverrideByIndex(path string, indexToFind int) error {
	_bea := path[0 : len(path)-5]
	if !_ca.HasPrefix(_bea, "\u002f") {
		_bea = "\u002f" + _bea
	}
	_fd, _baf := _dc.Compile(_bea + "\u0028\u005b\u0030-\u0039\u005d\u002b\u0029\u002e\u0078\u006d\u006c")
	if _baf != nil {
		return _baf
	}
	_afd := 0
	_ffg := -1
	for _dca, _bg := range _gdec._abd.Override {
		if _feb := _fd.FindStringSubmatch(_bg.PartNameAttr); len(_feb) > 1 {
			if _afd == indexToFind {
				_ffg = _dca
			} else if _afd > indexToFind {
				_dcc, _ := _fe.Atoi(_feb[1])
				_dcc--
				_bg.PartNameAttr = _bde.Sprintf("\u0025\u0073\u0025\u0064\u002e\u0078\u006d\u006c", _bea, _dcc)
			}
			_afd++
		}
	}
	if _ffg > -1 {
		copy(_gdec._abd.Override[_ffg:], _gdec._abd.Override[_ffg+1:])
		_gdec._abd.Override = _gdec._abd.Override[0 : len(_gdec._abd.Override)-1]
	}
	return nil
}

func (_edaf CustomProperties) SetPropertyAsI2(name string, i2 int16) {
	_cabd := _edaf.getNewProperty(name)
	_cabd.I2 = &i2
	_edaf.setOrReplaceProperty(_cabd)
}

// X returns the inner wrapped XML type.
func (_edb CoreProperties) X() *_ag.CoreProperties { return _edb._ae }

func (_geab Relationship) String() string {
	return _bde.Sprintf("\u007b\u0049\u0044\u003a \u0025\u0073\u0020\u0054\u0061\u0072\u0067\u0065\u0074\u003a \u0025s\u0020\u0054\u0079\u0070\u0065\u003a\u0020%\u0073\u007d", _geab.ID(), _geab.Target(), _geab.Type())
}

// Cells returns an array of row cells.
func (_gfed TableRow) Cells() []*_bdd.CT_TableCell { return _gfed._cbde.Tc }

// SetCompany sets the name of the company that created the document.
func (_bca AppProperties) SetCompany(s string) { _bca._agfd.Company = &s }

func (_ccf CustomProperties) SetPropertyAsVstream(name string, vstream *_de.Vstream) {
	_fcdg := _ccf.getNewProperty(name)
	_fcdg.Vstream = vstream
	_ccf.setOrReplaceProperty(_fcdg)
}

const Version = "\u0031\u002e\u0032\u0038\u002e\u0030"

// AddCustomRelationships adds relationships related to custom properties to the document.
func (_aabd *DocBase) AddCustomRelationships() {
	_aabd.ContentTypes.AddOverride("/\u0064o\u0063\u0050\u0072\u006f\u0070\u0073\u002f\u0063u\u0073\u0074\u006f\u006d.x\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002fv\u006e\u0064\u002e\u006f\u0070\u0065n\u0078\u006d\u006c\u0066\u006fr\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u002e\u0063\u0075\u0073\u0074\u006f\u006d\u002d\u0070r\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073+\u0078\u006d\u006c")
	_aabd.Rels.AddRelationship("\u0064\u006f\u0063\u0050ro\u0070\u0073\u002f\u0063\u0075\u0073\u0074\u006f\u006d\u002e\u0078\u006d\u006c", _ea.CustomPropertiesType)
}

func (_dfg CustomProperties) SetPropertyAsOblob(name, oblob string) {
	_edc := _dfg.getNewProperty(name)
	_edc.Oblob = &oblob
	_dfg.setOrReplaceProperty(_edc)
}

func UtcTimeFormat(t _gb.Time) string { return t.Format(_befb) + "\u0020\u0055\u0054\u0043" }

// Modified returns the time that the document was modified.
func (_gcd CoreProperties) Modified() _gb.Time { return _dffc(_gcd._ae.Modified) }

// ContentStatus returns the content status of the document (e.g. "Final", "Draft")
func (_gee CoreProperties) ContentStatus() string {
	if _gee._ae.ContentStatus != nil {
		return *_gee._ae.ContentStatus
	}
	return ""
}

// CustomProperty contains document specific property.
// Using of this type is deprecated.
type CustomProperty struct{ _dba *_ga.CT_Property }

// SetID set the ID of a relationship.
func (_gbcg Relationship) SetID(ID string) { _gbcg._fdg.IdAttr = ID }

// SetLinksUpToDate sets the links up to date flag.
func (_bbe AppProperties) SetLinksUpToDate(v bool) { _bbe._agfd.LinksUpToDate = _ea.Bool(v) }

func _dffc(_aggg *_ea.XSDAny) _gb.Time {
	if _aggg == nil {
		return _gb.Time{}
	}
	_bee, _fea := _gb.Parse(_bbg, string(_aggg.Data))
	if _fea != nil {
		_ge.Log.Debug("\u0065\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0074\u0069\u006d\u0065 \u0066\u0072\u006f\u006d\u0020\u0025\u0073\u003a\u0020\u0025\u0073", string(_aggg.Data), _fea)
	}
	return _bee
}

// Pages returns total number of pages which are saved by the text editor which produced the document.
// For github.com/arcpd/msword created documents, it is 0.
func (_gc AppProperties) Pages() int32 {
	if _gc._agfd.Pages != nil {
		return *_gc._agfd.Pages
	}
	return 0
}

// NewTheme constructs a new theme.
func NewTheme() Theme { return Theme{_bdd.NewTheme()} }

func (_gdeb CustomProperties) SetPropertyAsBool(name string, b bool) {
	_gbb := _gdeb.getNewProperty(name)
	_gbb.Bool = &b
	_gdeb.setOrReplaceProperty(_gbb)
}

// GetTargetByRelIdAndType returns a target path with the associated relation ID.
func (_gdg Relationships) GetTargetByRelIdAndType(idAttr string, typeAttr string) string {
	for _, _bccg := range _gdg._dafb.Relationship {
		if _bccg.IdAttr == idAttr && _bccg.TypeAttr == typeAttr {
			return _bccg.TargetAttr
		}
	}
	return ""
}

// NewCustomProperties constructs a new CustomProperties.
func NewCustomProperties() CustomProperties { return CustomProperties{_gbg: _ga.NewProperties()} }

// X returns the inner wrapped XML type.
func (_fab AppProperties) X() *_ad.Properties { return _fab._agfd }

func (_fdd CustomProperties) SetPropertyAsVector(name string, vector *_de.Vector) {
	_bacg := _fdd.getNewProperty(name)
	_bacg.Vector = vector
	_fdd.setOrReplaceProperty(_bacg)
}

func (_eef CustomProperties) setOrReplaceProperty(_ead *_ga.CT_Property) {
	_eef.setPropertyHelper(_ead, true)
}

// X returns the inner wrapped XML type.
func (_gad Relationship) X() *_ac.Relationship { return _gad._fdg }

// ExtraFile is an unsupported file type extracted from, or to be written to a
// zip package
type ExtraFile struct {
	ZipPath  string
	DiskPath string
}

const _baeb = 2023

// IsEmpty returns true if there are no relationships.
func (_afc Relationships) IsEmpty() bool {
	return _afc._dafb == nil || len(_afc._dafb.Relationship) == 0
}

// SetAuthor records the author of the document.
func (_eda CoreProperties) SetAuthor(s string) {
	if _eda._ae.Creator == nil {
		_eda._ae.Creator = &_ea.XSDAny{XMLName: _e.Name{Local: "\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}}
	}
	_eda._ae.Creator.Data = []byte(s)
}

// Relationship is a relationship within a .rels file.
type Relationship struct{ _fdg *_ac.Relationship }

func (_egc CustomProperties) SetPropertyAsUint(name string, ui uint) {
	_gggf := _egc.getNewProperty(name)
	_edg := uint32(ui)
	_gggf.Uint = &_edg
	_egc.setOrReplaceProperty(_gggf)
}

// Target returns the target (path) of a relationship.
func (_ebad Relationship) Target() string { return _ebad._fdg.TargetAttr }

func (_bba CustomProperties) SetPropertyAsNull(name string) {
	_abg := _bba.getNewProperty(name)
	_abg.Null = _de.NewNull()
	_bba.setOrReplaceProperty(_abg)
}

// Properties returns table properties.
func (_egb Table) Grid() *_bdd.CT_TableGrid { return _egb._bef.TblGrid }

// NewTable makes a new table.
func NewTable() *Table {
	_cbe := _bdd.NewTbl()
	_cbe.TblPr = _bdd.NewCT_TableProperties()
	return &Table{_bef: _cbe}
}

func (_gecb CustomProperties) setProperty(_faf *_ga.CT_Property) {
	_gecb.setPropertyHelper(_faf, false)
}

func (_bdcb CustomProperties) SetPropertyAsOstorage(name string, ostorage string) {
	_gaa := _bdcb.getNewProperty(name)
	_gaa.Ostorage = &ostorage
	_bdcb.setOrReplaceProperty(_gaa)
}

// AddRelationship adds a relationship.
func (_ebag Relationships) AddRelationship(target, ctype string) Relationship {
	if !_ca.HasPrefix(ctype, "\u0068t\u0074\u0070\u003a\u002f\u002f") {
		_ge.Log.Debug("\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006es\u0068\u0069\u0070 t\u0079\u0070\u0065\u0020\u0025\u0073 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069t\u0068\u0020\u0027\u0068\u0074\u0074\u0070\u003a/\u002f\u0027", ctype)
	}
	_cdab := _ac.NewRelationship()
	_geec := len(_ebag._dafb.Relationship) + 1
	_fgcd := map[string]struct{}{}
	for _, _eebd := range _ebag._dafb.Relationship {
		_fgcd[_eebd.IdAttr] = struct{}{}
	}
	for _, _gfg := _fgcd[_bde.Sprintf("\u0072\u0049\u0064%\u0064", _geec)]; _gfg; _, _gfg = _fgcd[_bde.Sprintf("\u0072\u0049\u0064%\u0064", _geec)] {
		_geec++
	}
	_cdab.IdAttr = _bde.Sprintf("\u0072\u0049\u0064%\u0064", _geec)
	_cdab.TargetAttr = target
	_cdab.TypeAttr = ctype
	_ebag._dafb.Relationship = append(_ebag._dafb.Relationship, _cdab)
	return Relationship{_fdg: _cdab}
}

// Relationships represents a .rels file.
type Relationships struct{ _dafb *_ac.Relationships }

// PropertiesList returns the list of all custom properties of the document.
func (_fdf CustomProperties) PropertiesList() []*_ga.CT_Property { return _fdf._gbg.Property }

const _ccc = 9

// Image is a container for image information. It's used as we need format and
// and size information to use images.
// It contains either the filesystem path to the image, or the image itself.
type Image struct {
	Size   _cee.Point
	Format string
	Path   string
	Data   *[]byte
}

func (_gea CustomProperties) SetPropertyAsI4(name string, i4 int32) {
	_dgab := _gea.getNewProperty(name)
	_dgab.I4 = &i4
	_gea.setOrReplaceProperty(_dgab)
}

// Size returns the size of an image
func (_aeg ImageRef) Size() _cee.Point { return _aeg._fcb.Size }

// ApplicationVersion returns the version of the application that created the
// document.
func (_fff AppProperties) ApplicationVersion() string {
	if _fff._agfd.AppVersion != nil {
		return *_fff._agfd.AppVersion
	}
	return ""
}

// TableCol represents a column in a table.
type TableCol struct{ _eaba *_bdd.CT_TableCol }

// DefAttr returns the DefAttr property.
func (_ccec TableStyles) DefAttr() string { return _ccec._debd.DefAttr }

// SetLastModifiedBy records the last person to modify the document.
func (_fae CoreProperties) SetLastModifiedBy(s string) { _fae._ae.LastModifiedBy = &s }

// X returns the inner wrapped XML type of CustomProperty.
func (_ecf CustomProperty) X() *_ga.CT_Property { return _ecf._dba }

// NewAppProperties constructs a new AppProperties.
func NewAppProperties() AppProperties {
	_af := AppProperties{_agfd: _ad.NewProperties()}
	_af.SetCompany("\u0046\u006f\u0078\u0079\u0055\u0074\u0069\u006c\u0073\u0020\u0065\u0068\u0066")
	_af.SetApplication("g\u0069\u0074\u0068\u0075\u0062\u002ec\u006f\u006d\u002f\u0075\u006e\u0069\u0064\u006f\u0063/\u0075\u006e\u0069o\u0066f\u0069\u0063\u0065")
	_af.SetDocSecurity(0)
	_af.SetLinksUpToDate(false)
	var _abb, _afg, _ee int64
	_bde.Sscanf(Version, "\u0025\u0064\u002e\u0025\u0064\u002e\u0025\u0064", &_abb, &_afg, &_ee)
	_db := float64(_abb) + float64(_afg)/10000.0
	_af.SetApplicationVersion(_bde.Sprintf("\u0025\u0030\u0037\u002e\u0034\u0066", _db))
	return _af
}

// Table represents a table in the document.
type Table struct {
	_bef  *_bdd.Tbl
	_dgcc *_bdd.CT_Transform2D
}

func (_eeb CustomProperties) SetPropertyAsStorage(name string, storage string) {
	_daf := _eeb.getNewProperty(name)
	_daf.Storage = &storage
	_eeb.setOrReplaceProperty(_daf)
}

// ImageFromStorage reads an image using the currently set
// temporary storage mechanism (see tempstorage). You can also
// construct an Image directly if the file and size are known.
func ImageFromStorage(path string) (Image, error) {
	_acdf := Image{}
	_afga, _bcac := _eg.Open(path)
	if _bcac != nil {
		return _acdf, _bde.Errorf("\u0065\u0072\u0072or\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _bcac)
	}
	defer _afga.Close()
	_aagd, _dbe, _bcac := _cee.Decode(_afga)
	if _bcac != nil {
		return _acdf, _bde.Errorf("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s", _bcac)
	}
	_acdf.Path = path
	_acdf.Format = _dbe
	_acdf.Size = _aagd.Bounds().Size()
	return _acdf, nil
}

func _agca(_cff _gb.Time, _aag string) *_ea.XSDAny {
	_fgc := &_ea.XSDAny{XMLName: _e.Name{Local: _aag}}
	_fgc.Attrs = append(_fgc.Attrs, _e.Attr{Name: _e.Name{Local: "\u0078\u0073\u0069\u003a\u0074\u0079\u0070\u0065"}, Value: "\u0064\u0063\u0074\u0065\u0072\u006d\u0073\u003a\u00573\u0043\u0044\u0054\u0046"})
	_fgc.Attrs = append(_fgc.Attrs, _e.Attr{Name: _e.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u0073i"}, Value: "\u0068\u0074\u0074\u0070\u003a/\u002f\u0077\u0077\u0077\u002e\u0077\u0033\u002e\u006f\u0072\u0067\u002f\u00320\u0030\u0031\u002f\u0058\u004d\u004c\u0053\u0063\u0068\u0065\u006d\u0061\u002d\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065"})
	_fgc.Attrs = append(_fgc.Attrs, _e.Attr{Name: _e.Name{Local: "\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"}, Value: "\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"})
	_fgc.Data = []byte(_cff.Format(_bbg))
	return _fgc
}

// Author returns the author of the document
func (_agc CoreProperties) Author() string {
	if _agc._ae.Creator != nil {
		return string(_agc._ae.Creator.Data)
	}
	return ""
}

// NewTableStyles constructs a new TableStyles.
func NewTableStyles() TableStyles { return TableStyles{_debd: _bdd.NewTblStyleLst()} }

// SetApplicationVersion sets the version of the application that created the
// document.  Per MS, the verison string mut be in the form 'XX.YYYY'.
func (_ada AppProperties) SetApplicationVersion(s string) { _ada._agfd.AppVersion = &s }

// SetStyle assigns TableStyle to a table.
func (_dbge Table) SetStyle(style *_bdd.CT_TableStyle) {
	if _dbge._bef.TblPr == nil {
		_dbge._bef.TblPr = _bdd.NewCT_TableProperties()
	}
	if _dbge._bef.TblPr.Choice == nil {
		_dbge._bef.TblPr.Choice = _bdd.NewCT_TablePropertiesChoice()
	}
	_dbge._bef.TblPr.Choice.TableStyle = style
}

// TblStyle returns the TblStyle property.
func (_bdgc TableStyles) TblStyle() []*_bdd.CT_TableStyle { return _bdgc._debd.TblStyle }

// Application returns the name of the application that created the document.
// For github.com/arcpd/msword created documents, it defaults to github.com/arcpd/msword
func (_cce AppProperties) Application() string {
	if _cce._agfd.Application != nil {
		return *_cce._agfd.Application
	}
	return ""
}

func (_aabf CustomProperties) SetPropertyAsDate(name string, date _gb.Time) {
	date = date.UTC()
	_dgcg, _dcce, _dacc := date.Date()
	_fdfe, _fcg, _fggf := date.Clock()
	_gfe := _gb.Date(_dgcg, _dcce, _dacc, _fdfe, _fcg, _fggf, 0, _gb.UTC)
	_gge := _aabf.getNewProperty(name)
	_gge.Filetime = &_gfe
	_aabf.setOrReplaceProperty(_gge)
}

func init() { _gg.SetAsStorage() }

// SetLanguage records the language of the document.
func (_ggg CoreProperties) SetLanguage(s string) {
	_ggg._ae.Language = &_ea.XSDAny{XMLName: _e.Name{Local: "d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}}
	_ggg._ae.Language.Data = []byte(s)
}

// AppProperties contains properties specific to the document and the
// application that created it.
type AppProperties struct{ _agfd *_ad.Properties }

func (_eeg CustomProperties) SetPropertyAsR4(name string, r4 float32) {
	_cge := _eeg.getNewProperty(name)
	_cge.R4 = &r4
	_eeg.setOrReplaceProperty(_cge)
}

// NewContentTypes returns a wrapper around a newly constructed content-types.
func NewContentTypes() ContentTypes {
	_afb := ContentTypes{_abd: _fc.NewTypes()}
	_afb.AddDefault("\u0078\u006d\u006c", "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c")
	_afb.AddDefault("\u0072\u0065\u006c\u0073", "\u0061\u0070\u0070\u006c\u0069\u0063a\u0074\u0069\u006fn\u002f\u0076\u006ed\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006fr\u006d\u0061\u0074\u0073\u002dpa\u0063\u006b\u0061\u0067\u0065\u002e\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u002b\u0078\u006d\u006c")
	_afb.AddDefault("\u0070\u006e\u0067", "\u0069m\u0061\u0067\u0065\u002f\u0070\u006eg")
	_afb.AddDefault("\u006a\u0070\u0065\u0067", "\u0069\u006d\u0061\u0067\u0065\u002f\u006a\u0070\u0065\u0067")
	_afb.AddDefault("\u006a\u0070\u0067", "\u0069m\u0061\u0067\u0065\u002f\u006a\u0070g")
	_afb.AddDefault("\u0077\u006d\u0066", "i\u006d\u0061\u0067\u0065\u002f\u0078\u002d\u0077\u006d\u0066")
	_afb.AddOverride("\u002fd\u006fc\u0050\u0072\u006f\u0070\u0073/\u0063\u006fr\u0065\u002e\u0078\u006d\u006c", "\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073-\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002e\u0063\u006f\u0072\u0065\u002dp\u0072\u006f\u0070\u0065\u0072\u0074i\u0065\u0073\u002bx\u006d\u006c")
	_afb.AddOverride("\u002f\u0064\u006f\u0063\u0050\u0072\u006f\u0070\u0073\u002f\u0061\u0070p\u002e\u0078\u006d\u006c", "a\u0070\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0076\u006e\u0064\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066o\u0072\u006d\u0061\u0074\u0073\u002d\u006f\u0066\u0066\u0069\u0063\u0065\u0064\u006f\u0063\u0075m\u0065\u006e\u0074\u002e\u0065\u0078\u0074\u0065\u006e\u0064\u0065\u0064\u002dp\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u002b\u0078m\u006c")
	return _afb
}

// SetTitle records the title of the document.
func (_fbb CoreProperties) SetTitle(s string) {
	if _fbb._ae.Title == nil {
		_fbb._ae.Title = &_ea.XSDAny{XMLName: _e.Name{Local: "\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}}
	}
	_fbb._ae.Title.Data = []byte(s)
}

// GetOrCreateCustomProperties returns the custom properties of the document (and if they not exist yet, creating them first).
func (_dcaf *DocBase) GetOrCreateCustomProperties() CustomProperties {
	if _dcaf.CustomProperties.X() == nil {
		_dcaf.CreateCustomProperties()
	}
	return _dcaf.CustomProperties
}

// SetWidth sets column width, see measurement package.
func (_dedc TableCol) SetWidth(m _gf.Distance) {
	_cddc := _gf.ToEMU(float64(m))
	_dedc._eaba.WAttr = _bdd.ST_Coordinate{ST_CoordinateUnqualified: &_cddc}
}

// WriteExtraFiles writes the extra files to the zip package.
func (_egf *DocBase) WriteExtraFiles(z *_bd.Writer) error {
	for _, _bfd := range _egf.ExtraFiles {
		if _fccb := _dg.AddFileFromDisk(z, _bfd.ZipPath, _bfd.DiskPath); _fccb != nil {
			return _fccb
		}
	}
	return nil
}

func (_dce CustomProperties) SetPropertyAsUi8(name string, ui8 uint64) {
	_aaba := _dce.getNewProperty(name)
	_aaba.Ui8 = &ui8
	_dce.setOrReplaceProperty(_aaba)
}

// NewRelationships creates a new relationship wrapper.
func NewRelationships() Relationships { return Relationships{_dafb: _ac.NewRelationships()} }

// SetHeight sets row height, see measurement package.
func (_bed TableRow) SetHeight(m _gf.Distance) {
	_dgba := _gf.ToEMU(float64(m))
	_bed._cbde.HAttr = _bdd.ST_Coordinate{ST_CoordinateUnqualified: &_dgba}
}

// CopyRelationship copies the relationship.
func (_afbe Relationships) CopyRelationship(idAttr string) (Relationship, bool) {
	for _deeg := range _afbe._dafb.Relationship {
		if _afbe._dafb.Relationship[_deeg].IdAttr == idAttr {
			_fbe := *_afbe._dafb.Relationship[_deeg]
			_ccef := len(_afbe._dafb.Relationship) + 1
			_bafc := map[string]struct{}{}
			for _, _dffg := range _afbe._dafb.Relationship {
				_bafc[_dffg.IdAttr] = struct{}{}
			}
			for _, _bdeb := _bafc[_bde.Sprintf("\u0072\u0049\u0064%\u0064", _ccef)]; _bdeb; _, _bdeb = _bafc[_bde.Sprintf("\u0072\u0049\u0064%\u0064", _ccef)] {
				_ccef++
			}
			_fbe.IdAttr = _bde.Sprintf("\u0072\u0049\u0064%\u0064", _ccef)
			_afbe._dafb.Relationship = append(_afbe._dafb.Relationship, &_fbe)
			return Relationship{_fdg: &_fbe}, true
		}
	}
	return Relationship{}, false
}

// NewRelationship constructs a new relationship.
func NewRelationship() Relationship { return Relationship{_fdg: _ac.NewRelationship()} }

func (_abgd *ImageRef) SetRelID(id string) { _abgd._gfd = id }

func (_bccd CustomProperties) SetPropertyAsClsid(name string, clsid string) {
	_dcf := _bccd.getNewProperty(name)
	_dcf.Clsid = &clsid
	_bccd.setOrReplaceProperty(_dcf)
}

// TableStyles contains document specific properties.
type TableStyles struct{ _debd *_bdd.TblStyleLst }

func (_gac CustomProperties) setPropertyHelper(_gccg *_ga.CT_Property, _dcd bool) {
	_ccee := _gac.GetPropertyByName(*_gccg.NameAttr)
	if (_ccee == CustomProperty{}) {
		_gac._gbg.Property = append(_gac._gbg.Property, _gccg)
	} else if _dcd {
		_gccg.FmtidAttr = _ccee._dba.FmtidAttr
		if _ccee._dba.PidAttr == 0 {
			_gccg.PidAttr = _ccee._dba.PidAttr
		}
		_gccg.LinkTargetAttr = _ccee._dba.LinkTargetAttr
		*_ccee._dba = *_gccg
	}
}

// ImageFromFile reads an image from a file on disk. It doesn't keep the image
// in memory and only reads it to determine the format and size. You can also
// construct an Image directly if the file and size are known.
// NOTE: See also ImageFromStorage.
func ImageFromFile(path string) (Image, error) {
	_abf, _bddf := _ce.Open(path)
	_aef := Image{}
	if _bddf != nil {
		return _aef, _bde.Errorf("\u0065\u0072\u0072or\u0020\u0072\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073", _bddf)
	}
	defer _abf.Close()
	_ffec, _gbc, _bddf := _cee.Decode(_abf)
	if _bddf != nil {
		return _aef, _bde.Errorf("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020i\u006d\u0061\u0067\u0065\u003a\u0020\u0025s", _bddf)
	}
	_aef.Path = path
	_aef.Format = _gbc
	_aef.Size = _ffec.Bounds().Size()
	return _aef, nil
}

// Company returns the name of the company that created the document.
// For github.com/arcpd/msword created documents, it defaults to github.com/arcpd/msword
func (_ef AppProperties) Company() string {
	if _ef._agfd.Company != nil {
		return *_ef._agfd.Company
	}
	return ""
}

// CopyOverride copies override content type for a given `path` and puts it with a path `newPath`.
func (_dgc ContentTypes) CopyOverride(path, newPath string) {
	if !_ca.HasPrefix(path, "\u002f") {
		path = "\u002f" + path
	}
	if !_ca.HasPrefix(newPath, "\u002f") {
		newPath = "\u002f" + newPath
	}
	for _ggfb := range _dgc._abd.Override {
		if _dgc._abd.Override[_ggfb].PartNameAttr == path {
			_adb := *_dgc._abd.Override[_ggfb]
			_adb.PartNameAttr = newPath
			_dgc._abd.Override = append(_dgc._abd.Override, &_adb)
		}
	}
}

func (_bdf CustomProperties) SetPropertyAsInt(name string, i int) {
	_fbc := _bdf.getNewProperty(name)
	_bda := int32(i)
	_fbc.Int = &_bda
	_bdf.setOrReplaceProperty(_fbc)
}

func (_fba CustomProperties) SetPropertyAsUi4(name string, ui4 uint32) {
	_fdc := _fba.getNewProperty(name)
	_fdc.Ui4 = &ui4
	_fba.setOrReplaceProperty(_fdc)
}

// Hyperlink is just an appropriately configured relationship.
type Hyperlink Relationship

func (_gef TableRow) addCell() *_bdd.CT_TableCell {
	_cgec := _bdd.NewCT_TableCell()
	_gef._cbde.Tc = append(_gef._cbde.Tc, _cgec)
	return _cgec
}

var _eeba = _d.New("\u0069\u006d\u0061\u0067\u0065\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u0069\u006e\u0020\u0073\u0074\u006fr\u0061\u0067\u0065")

// Description returns the description of the document
func (_fca CoreProperties) Description() string {
	if _fca._ae.Description != nil {
		return string(_fca._ae.Description.Data)
	}
	return ""
}

// Format returns the format of the underlying image
func (_egd ImageRef) Format() string { return _egd._fcb.Format }

// EnsureOverride ensures that an override for the given path exists, adding it if necessary
func (_fec ContentTypes) EnsureOverride(path, contentType string) {
	for _, _aace := range _fec._abd.Override {
		if _aace.PartNameAttr == path {
			if _ca.HasPrefix(contentType, "\u0068\u0074\u0074\u0070") {
				_ge.Log.Debug("\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u0074\u0079p\u0065\u0020\u0027%\u0073\u0027\u0020\u0069\u0073\u0020\u0069\u006e\u0063\u006fr\u0072\u0065\u0063\u0074\u002c m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0073\u0074\u0061\u0072\u0074\u0020\u0077\u0069\u0074\u0068\u0020\u0068\u0074\u0074\u0070", contentType)
			}
			_aace.ContentTypeAttr = contentType
			return
		}
	}
	_fec.AddOverride(path, contentType)
}

func (_geg CustomProperties) SetPropertyAsUi2(name string, ui2 uint16) {
	_cgf := _geg.getNewProperty(name)
	_cgf.Ui2 = &ui2
	_geg.setOrReplaceProperty(_cgf)
}

const _befb = "\u0032\u0020\u004aan\u0075\u0061\u0072\u0079\u0020\u0032\u0030\u0030\u0036\u0020\u0061\u0074\u0020\u0031\u0035\u003a\u0030\u0034"

// Relationships returns a slice of all of the relationships.
func (_dace Relationships) Relationships() []Relationship {
	_dab := []Relationship{}
	for _, _eca := range _dace._dafb.Relationship {
		_dab = append(_dab, Relationship{_fdg: _eca})
	}
	return _dab
}

func (_gcb CustomProperties) SetPropertyAsBlob(name, blob string) {
	_fbd := _gcb.getNewProperty(name)
	_fbd.Blob = &blob
	_gcb.setOrReplaceProperty(_fbd)
}

const _ggdg = 17

// X returns the inner wrapped XML type.
func (_ebc Theme) X() *_bdd.Theme { return _ebc._acda }

func (_edcd CustomProperties) SetPropertyAsEmpty(name string) {
	_dffca := _edcd.getNewProperty(name)
	_dffca.Empty = _de.NewEmpty()
	_edcd.setOrReplaceProperty(_dffca)
}
