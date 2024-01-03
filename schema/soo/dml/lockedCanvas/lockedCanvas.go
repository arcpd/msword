package lockedCanvas

import (
	_ga "encoding/xml"
	_f "fmt"

	_fe "github.com/arcpd/msword"
	_c "github.com/arcpd/msword/schema/soo/dml"
)

func (_e *LockedCanvas) UnmarshalXML(d *_ga.Decoder, start _ga.StartElement) error {
	_e.CT_GvmlGroupShape = *_c.NewCT_GvmlGroupShape()
	for {
		_b, _ef := d.Token()
		if _ef != nil {
			return _f.Errorf("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u004c\u006f\u0063\u006b\u0065d\u0043\u0061\u006e\u0076\u0061\u0073\u003a\u0020\u0025\u0073", _ef)
		}
		if _fa, _cc := _b.(_ga.EndElement); _cc && _fa.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the LockedCanvas and its children
func (_gf *LockedCanvas) Validate() error {
	return _gf.ValidateWithPath("\u004c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073")
}

func (_gc *LockedCanvas) MarshalXML(e *_ga.Encoder, start _ga.StartElement) error {
	start.Attr = append(start.Attr, _ga.Attr{Name: _ga.Name{Local: "\u0078\u006d\u006cn\u0073"}, Value: "h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006c\u006f\u0063\u006b\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073"})
	start.Attr = append(start.Attr, _ga.Attr{Name: _ga.Name{Local: "\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"}, Value: "\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"})
	start.Name.Local = "\u006c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073"
	return _gc.CT_GvmlGroupShape.MarshalXML(e, start)
}

func NewLockedCanvas() *LockedCanvas {
	_ce := &LockedCanvas{}
	_ce.CT_GvmlGroupShape = *_c.NewCT_GvmlGroupShape()
	return _ce
}

type LockedCanvas struct{ _c.CT_GvmlGroupShape }

// ValidateWithPath validates the LockedCanvas and its children, prefixing error messages with path
func (_gb *LockedCanvas) ValidateWithPath(path string) error {
	if _d := _gb.CT_GvmlGroupShape.ValidateWithPath(path); _d != nil {
		return _d
	}
	return nil
}

func init() {
	_fe.RegisterConstructor("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006c\u006f\u0063\u006b\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073", "\u006c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073", NewLockedCanvas)
}
