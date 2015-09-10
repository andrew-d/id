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
	expected := "YOVY74J-XEDUK3J-ECH3U4U-M2Z4RFQ-2OLEWC7-I4D2SR4-ZMBYUZL-XQYTZAV"
	testCases := []string{
		// Original
		"YOVY74J-XEDUK3J-ECH3U4U-M2Z4RFQ-2OLEWC7-I4D2SR4-ZMBYUZL-XQYTZAV",

		// Mis-entered dashes
		"YOVY74JXE-DUK3J-ECH3U4UM2-Z4RFQ2O-LEWC7I4D2S-R4-ZMBYUZLXQY-TZAV",

		// Spaces instead of dashes
		"YOVY74J XEDUK3J ECH3U4U M2Z4RFQ 2OLEWC7 I4D2SR4 ZMBYUZL XQYTZAV",

		// Mis-entered spaces
		"YOVY74JXE DUK3J ECH3U4UM2 Z4RFQ2O LEWC7I4D2S R4 ZMBYUZLXQY TZAV",

		// No spaces OR dashes
		"YOVY74JXEDUK3JECH3U4UM2Z4RFQ2OLEWC7I4D2SR4ZMBYUZLXQYTZAV",

		// Lowercase
		"yovy74jxeduk3jech3u4um2z4rfq2olewc7i4d2sr4zmbyuzlxqytzav",

		// Mis-spelling - uses '0' instead of 'O'
		"Y0VY74JXEDUK3JECH3U4UM2Z4RFQ20LEWC7I4D2SR4ZMBYUZLXQYTZAV",

		// Mis-spelling - uses '0' instead of 'O', lowercase
		"y0vy74jxeduk3jech3u4um2z4rfq20lewc7i4d2sr4zmbyuzlxqytzav",

		// Mis-spelling - uses '1' instead of 'I'
		"YOVY74JXEDUK3JECH3U4UM2Z4RFQ2OLEWC714D2SR4ZMBYUZLXQYTZAV",

		// Mis-spelling - uses '1' instead of 'I', lowercase
		"yovy74jxeduk3jech3u4um2z4rfq2olewc714d2sr4zmbyuzlxqytzav",

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
		{"YOVY74J-XEDUK3J-ECH3U4U-M2Z4RFQ-2OLEWC7-I4D2SR4-ZMBYUZL-XQYTZAV", true},

		// Empty
		{"", false},

		// Spaces in place of dashes
		{"YOVY74J XEDUK3J ECH3U4U M2Z4RFQ 2OLEWC7 I4D2SR4 ZMBYUZL XQYTZAV", true},

		// No spaces or dashes
		{"YOVY74JXEDUK3JECH3U4UM2Z4RFQ2OLEWC7I4D2SR4ZMBYUZLXQYTZAV", true},

		// Missing a character from the middle (it's a space here')
		{"YOVY74JXEDUK3JECH3U4UM2Z4RF 2OLEWC7I4D2SR4ZMBYUZLXQYTZAV", false},

		// Extra bits at the end
		{"YOVY74JXEDUK3JECH3U4UM2Z4RFQ2OLEWC7I4D2SR4ZMBYUZLXQYTZAVcccc", false},
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
