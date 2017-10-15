package id

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Println

func TestID(t *testing.T) {
	id := NewFromString("foobar")
	//fmt.Println("id:", id.String())

	// Test round-trip
	var (
		m   []byte
		u   ID
		err error
	)

	m, err = id.MarshalText()
	assert.NoError(t, err)

	err = u.UnmarshalText(m)
	assert.NoError(t, err)

	assert.Equal(t, id[:], u[:])
}

// Test that we can unmarshal improperly-formatted things.
func TestIDFormatting(t *testing.T) {
	expected := "YOVY74J-XEDUK3F-ECH3U4U-M2Z4RFR-2OLEWC7-I4D2SRH-ZMBYUZL-XQYTZA7"
	testCases := []string{
		// Original
		"YOVY74J-XEDUK3F-ECH3U4U-M2Z4RFR-2OLEWC7-I4D2SRH-ZMBYUZL-XQYTZA7",

		// Mis-entered dashes
		"YOVY74JXE-DUK3F-ECH3U4UM2-Z4RFR2O-LEWC7I4D2S-RH-ZMBYUZLXQY-TZA7",

		// Spaces instead of dashes
		"YOVY74J XEDUK3F ECH3U4U M2Z4RFR 2OLEWC7 I4D2SRH ZMBYUZL XQYTZA7",

		// Mis-entered spaces
		"YOVY74JXE DUK3F ECH3U4UM2 Z4RFR2O LEWC7I4D2S RH ZMBYUZLXQY TZA7",

		// No spaces OR dashes
		"YOVY74JXEDUK3FECH3U4UM2Z4RFR2OLEWC7I4D2SRHZMBYUZLXQYTZA7",

		// Lowercase
		"yovy74jxeduk3fech3u4um2z4rfr2olewc7i4d2srhzmbyuzlxqytza7",

		// Mis-spelling - uses '0' instead of 'O'
		"Y0VY74JXEDUK3FECH3U4UM2Z4RFR20LEWC7I4D2SRHZMBYUZLXQYTZA7",

		// Mis-spelling - uses '0' instead of 'O', lowercase
		"y0vy74jxeduk3fech3u4um2z4rfr20lewc7i4d2srhzmbyuzlxqytza7",

		// Mis-spelling - uses '1' instead of 'I'
		"YOVY74JXEDUK3FECH3U4UM2Z4RFR2OLEWC714D2SRHZMBYUZLXQYTZA7",

		// Mis-spelling - uses '1' instead of 'I', lowercase
		"yovy74jxeduk3fech3u4um2z4rfr2olewc714d2srhzmbyuzlxqytza7",

		// TODO: another ID with 'B' / '8' misspelling?
	}

	for i, tcase := range testCases {
		var id ID

		err := id.UnmarshalText([]byte(tcase))
		if assert.NoError(t, err, "#%d error unmarshalling: %s", i, err) {
			assert.Equal(t, expected, id.String(),
				"#%d formatting incorrect", i)
		}
	}
}

// Test that we catch invalid IDs
func TestIDValidity(t *testing.T) {
	testCases := []struct {
		s     string
		valid bool
	}{
		// Original
		{"YOVY74J-XEDUK3F-ECH3U4U-M2Z4RFR-2OLEWC7-I4D2SRH-ZMBYUZL-XQYTZA7", true},

		// Empty
		{"", false},

		// Spaces in place of dashes
		{"YOVY74J XEDUK3F ECH3U4U M2Z4RFR 2OLEWC7 I4D2SRH ZMBYUZL XQYTZA7", true},

		// No spaces or dashes
		{"YOVY74JXEDUK3FECH3U4UM2Z4RFR2OLEWC7I4D2SRHZMBYUZLXQYTZA7", true},

		// Missing a character from the middle (it's a space here')
		{"YOVY74JXEDUK3FECH3U4UM2Z4RF 2OLEWC7I4D2SRHZMBYUZLXQYTZA7", false},

		// Extra bits at the end
		{"YOVY74JXEDUK3FECH3U4UM2Z4RFR2OLEWC7I4D2SRHZMBYUZLXQYTZA7cccc", false},
	}

	for i, tcase := range testCases {
		var id ID

		err := id.UnmarshalText([]byte(tcase.s))
		isValid := err == nil
		assert.Equal(t, tcase.valid, isValid, "#%d improper validity", i)
	}
}

// Test comparison functions
func TestComparison(t *testing.T) {
	id := NewFromString("foobar")

	assert.True(t, id.Equals(id))
	assert.Equal(t, 0, id.Compare(id))
}
