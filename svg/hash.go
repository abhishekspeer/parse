package svg

// generated by hasher -type=Hash -file=hash.go; DO NOT EDIT, except for adding more constants to the list and rerun go generate

// uses github.com/tdewolff/hasher
//go:generate hasher -type=Hash -file=hash.go
type Hash uint32

const (
	A                            Hash = 0x1d01
	Alignment_Baseline           Hash = 0x2912
	Baseline_Shift               Hash = 0x330e
	Buffered_Rendering           Hash = 0x12
	Clip                         Hash = 0x4d04
	Clip_Path                    Hash = 0x4d09
	Clip_Rule                    Hash = 0x6909
	Color                        Hash = 0x9505
	Color_Interpolation          Hash = 0x9513
	Color_Interpolation_Filters  Hash = 0x951b
	Color_Profile                Hash = 0x1bc0d
	Color_Rendering              Hash = 0x1f50f
	Cursor                       Hash = 0x7e06
	D                            Hash = 0x701
	Defs                         Hash = 0x30b04
	Direction                    Hash = 0x2b109
	Display                      Hash = 0x16007
	Dominant_Baseline            Hash = 0x14011
	Enable_Background            Hash = 0x15011
	Fill                         Hash = 0x8604
	Fill_Opacity                 Hash = 0x2df0c
	Fill_Rule                    Hash = 0x8609
	Filter                       Hash = 0xa906
	Flood_Color                  Hash = 0x8f0b
	Flood_Opacity                Hash = 0xbe0d
	Font                         Hash = 0xcb04
	Font_Family                  Hash = 0xcb0b
	Font_Size                    Hash = 0xd609
	Font_Size_Adjust             Hash = 0xd610
	Font_Stretch                 Hash = 0xee0c
	Font_Style                   Hash = 0xfa0a
	Font_Variant                 Hash = 0x1040c
	Font_Weight                  Hash = 0x11d0b
	G                            Hash = 0x1101
	Glyph_Orientation_Horizontal Hash = 0x1981c
	Glyph_Orientation_Vertical   Hash = 0x111a
	Height                       Hash = 0x5506
	Image_Rendering              Hash = 0x1280f
	Kerning                      Hash = 0x18107
	Letter_Spacing               Hash = 0x700e
	Lighting_Color               Hash = 0x1b30e
	Line                         Hash = 0x3704
	Marker                       Hash = 0x13706
	Marker_End                   Hash = 0x1370a
	Marker_Mid                   Hash = 0x1680a
	Marker_Start                 Hash = 0x1720c
	Mask                         Hash = 0x17e04
	Metadata                     Hash = 0x18808
	Missing_Glyph                Hash = 0x1900d
	Opacity                      Hash = 0xc407
	Overflow                     Hash = 0x1c908
	Paint_Order                  Hash = 0x2530b
	Path                         Hash = 0x5204
	Pattern                      Hash = 0x1db07
	Pointer_Events               Hash = 0x1e20e
	Points                       Hash = 0x20406
	Polygon                      Hash = 0x21607
	Polyline                     Hash = 0x21d08
	Rect                         Hash = 0x2b304
	Rx                           Hash = 0x4a02
	Ry                           Hash = 0x8302
	Shape_Rendering              Hash = 0xaf0f
	Solid_Color                  Hash = 0x1ef0b
	Solid_Opacity                Hash = 0x2090d
	Stop_Color                   Hash = 0xe40a
	Stop_Opacity                 Hash = 0x30e0c
	Stroke                       Hash = 0x22506
	Stroke_Dasharray             Hash = 0x22510
	Stroke_Dashoffset            Hash = 0x23511
	Stroke_Linecap               Hash = 0x2460e
	Stroke_Linejoin              Hash = 0x25e0f
	Stroke_Miterlimit            Hash = 0x26d11
	Stroke_Opacity               Hash = 0x27e0e
	Stroke_Width                 Hash = 0x28c0c
	Style                        Hash = 0xff05
	Svg                          Hash = 0x29803
	Switch                       Hash = 0x29b06
	Symbol                       Hash = 0x2a106
	Text_Anchor                  Hash = 0x400b
	Text_Decoration              Hash = 0x5a0f
	Text_Rendering               Hash = 0x10f0e
	Unicode_Bidi                 Hash = 0x2a70c
	Vector_Effect                Hash = 0x2ba0d
	Version                      Hash = 0x2c707
	ViewBox                      Hash = 0x2ce07
	Viewport_Fill                Hash = 0x2d60d
	Viewport_Fill_Opacity        Hash = 0x2d615
	Visibility                   Hash = 0x2eb0a
	White_Space                  Hash = 0x1d00b
	Width                        Hash = 0x29305
	Word_Spacing                 Hash = 0x2f50c
	Writing_Mode                 Hash = 0x3010c
	X                            Hash = 0x4201
	X1                           Hash = 0x4b02
	X2                           Hash = 0x2d402
	Y                            Hash = 0x1301
	Y1                           Hash = 0x8402
	Y2                           Hash = 0x16602
)

// String returns the hash' name.
func (i Hash) String() string {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_Hash_text)) {
		return ""
	}
	return _Hash_text[start : start+n]
}

// Hash returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func ToHash(s []byte) Hash {
	if len(s) == 0 || len(s) > _Hash_maxLen {
		return 0
	}
	h := _Hash_fnv(s)
	if i := _Hash_table[h&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	if i := _Hash_table[(h>>16)&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	return 0
}

// _Hash_fnv computes the FNV hash with an arbitrary starting value h.
func _Hash_fnv(s []byte) uint32 {
	h := uint32(_Hash_hash0)
	for i := range s {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func _Hash_match(s string, t []byte) bool {
	for i, c := range t {
		if s[i] != c {
			return false
		}
	}
	return true
}

func _Hash_string(i Hash) string {
	return _Hash_text[i>>8 : i>>8+i&0xff]
}

const _Hash_hash0 = 0xf7ffa4d4
const _Hash_maxLen = 28
const _Hash_text = "buffered-renderinglyph-orientation-verticalignment-baseline-" +
	"shiftext-anchorx1clip-patheightext-decorationclip-ruletter-s" +
	"pacingcursory1fill-ruleflood-color-interpolation-filtershape" +
	"-renderingflood-opacityfont-familyfont-size-adjustop-colorfo" +
	"nt-stretchfont-stylefont-variantext-renderingfont-weightimag" +
	"e-renderingmarker-endominant-baselinenable-backgroundisplay2" +
	"marker-midmarker-startmaskerningmetadatamissing-glyph-orient" +
	"ation-horizontalighting-color-profileoverflowhite-spacepatte" +
	"rnpointer-eventsolid-color-renderingpointsolid-opacitypolygo" +
	"npolylinestroke-dasharraystroke-dashoffsetstroke-linecapaint" +
	"-orderstroke-linejoinstroke-miterlimitstroke-opacitystroke-w" +
	"idthsvgswitchsymbolunicode-bidirectionvector-effectversionvi" +
	"ewBox2viewport-fill-opacityvisibilityword-spacingwriting-mod" +
	"efstop-opacity"

var _Hash_table = [1 << 7]Hash{
	0x0:  0x14011, // dominant-baseline
	0x1:  0x18808, // metadata
	0x2:  0xd610,  // font-size-adjust
	0x3:  0xcb04,  // font
	0x4:  0x29305, // width
	0x5:  0x30b04, // defs
	0x6:  0x400b,  // text-anchor
	0x7:  0x2ba0d, // vector-effect
	0x8:  0x5204,  // path
	0x9:  0xfa0a,  // font-style
	0xa:  0x29803, // svg
	0xb:  0x26d11, // stroke-miterlimit
	0xc:  0x1981c, // glyph-orientation-horizontal
	0xd:  0x13706, // marker
	0xe:  0x1370a, // marker-end
	0xf:  0x1040c, // font-variant
	0x10: 0x2a106, // symbol
	0x11: 0x1e20e, // pointer-events
	0x12: 0x8402,  // y1
	0x13: 0x1900d, // missing-glyph
	0x14: 0xbe0d,  // flood-opacity
	0x16: 0x3010c, // writing-mode
	0x18: 0x2f50c, // word-spacing
	0x1b: 0x7e06,  // cursor
	0x1d: 0x4a02,  // rx
	0x1e: 0x30e0c, // stop-opacity
	0x21: 0xa906,  // filter
	0x22: 0x2b304, // rect
	0x25: 0x2a70c, // unicode-bidi
	0x26: 0x15011, // enable-background
	0x27: 0x5506,  // height
	0x28: 0x22510, // stroke-dasharray
	0x2a: 0x9505,  // color
	0x2c: 0x2d402, // x2
	0x2e: 0x4d09,  // clip-path
	0x2f: 0x4b02,  // x1
	0x30: 0x21d08, // polyline
	0x31: 0x8302,  // ry
	0x32: 0x1c908, // overflow
	0x33: 0x17e04, // mask
	0x34: 0x2d60d, // viewport-fill
	0x35: 0x2df0c, // fill-opacity
	0x36: 0x12,    // buffered-rendering
	0x38: 0x951b,  // color-interpolation-filters
	0x39: 0x10f0e, // text-rendering
	0x3a: 0x21607, // polygon
	0x3b: 0x5a0f,  // text-decoration
	0x3d: 0xd609,  // font-size
	0x3e: 0x1720c, // marker-start
	0x3f: 0x1bc0d, // color-profile
	0x40: 0x20406, // points
	0x41: 0xee0c,  // font-stretch
	0x42: 0xe40a,  // stop-color
	0x43: 0x16007, // display
	0x44: 0x4201,  // x
	0x45: 0x1ef0b, // solid-color
	0x46: 0x3704,  // line
	0x48: 0x2c707, // version
	0x49: 0x1101,  // g
	0x4b: 0x6909,  // clip-rule
	0x4d: 0x29b06, // switch
	0x4f: 0xc407,  // opacity
	0x50: 0x9513,  // color-interpolation
	0x51: 0x28c0c, // stroke-width
	0x52: 0x4d04,  // clip
	0x53: 0x1f50f, // color-rendering
	0x54: 0x22506, // stroke
	0x55: 0x2b109, // direction
	0x56: 0x1b30e, // lighting-color
	0x57: 0x1301,  // y
	0x59: 0xff05,  // style
	0x5a: 0x330e,  // baseline-shift
	0x5c: 0x8604,  // fill
	0x5d: 0xaf0f,  // shape-rendering
	0x5f: 0x2460e, // stroke-linecap
	0x60: 0x23511, // stroke-dashoffset
	0x61: 0x2530b, // paint-order
	0x63: 0x2090d, // solid-opacity
	0x64: 0x2912,  // alignment-baseline
	0x65: 0x1d00b, // white-space
	0x66: 0x1db07, // pattern
	0x67: 0x27e0e, // stroke-opacity
	0x68: 0x2eb0a, // visibility
	0x6a: 0x2ce07, // viewBox
	0x6b: 0x700e,  // letter-spacing
	0x6c: 0xcb0b,  // font-family
	0x6f: 0x1d01,  // a
	0x70: 0x701,   // d
	0x71: 0x25e0f, // stroke-linejoin
	0x72: 0x2d615, // viewport-fill-opacity
	0x74: 0x1280f, // image-rendering
	0x76: 0x8f0b,  // flood-color
	0x77: 0x1680a, // marker-mid
	0x78: 0x18107, // kerning
	0x7a: 0x11d0b, // font-weight
	0x7b: 0x111a,  // glyph-orientation-vertical
	0x7e: 0x8609,  // fill-rule
	0x7f: 0x16602, // y2
}
