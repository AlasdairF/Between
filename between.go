package between

import (
 "bytes"
 "github.com/AlasdairF/Conv"
 "errors"
)

func Crop(b, at []byte) []byte {
	if i := bytes.Index(b, at); i > 0 {
		return b[i:]
	}
	return b
}

// Bytes returns the result as a slice of bytes.
func Bytes(b, from, to []byte, within int) ([]byte, error) {
	var pos int
	for {
		pos = bytes.Index(b, from)
		if pos == -1 {
			return nil, errors.New("From does not exist.")
		}
		b = b[pos + len(from):]
		pos = bytes.Index(b, to)
		if pos == -1 {
			return nil, errors.New("To does not exist.")
		}
		if pos < within || within < 0  {
			return b[0:pos], nil
			b = b[pos + len(to):]
		}
	}
}

// AllBytes returns all results between from & to.
func AllBytes(b, from, to []byte, within int) [][]byte {
	result := make([][]byte, 0)
	var pos int
	for {
		pos = bytes.Index(b, from)
		if pos == -1 {
			return result
		}
		b = b[pos + len(from):]
		pos = bytes.Index(b, to)
		if pos == -1 {
			return result
		}
		if pos < within || within < 0  {
			result = append(result, b[0:pos])
			b = b[pos + len(to):]
		}
	}
}

// Int returns the result as an int.
func Int(b, from, to []byte, within int) (int, error) {
	var pos int
	for {
		pos = bytes.Index(b, from)
		if pos == -1 {
			return 0, errors.New("From does not exist.")
		}
		b = b[pos + len(from):]
		pos = bytes.Index(b, to)
		if pos == -1 {
			return 0, errors.New("To does not exist.")
		}
		if pos < within || within < 0  {
			return conv.Int(b[0:pos]), nil
			b = b[pos + len(to):]
		}
	}
}

// AllInt returns all results between from & to as int.
func AllInt(b, from, to []byte, within int) []int {
	result := make([]int, 0)
	var pos int
	for {
		pos = bytes.Index(b, from)
		if pos == -1 {
			return result
		}
		b = b[pos + len(from):]
		pos = bytes.Index(b, to)
		if pos == -1 {
			return result
		}
		if pos < within || within < 0  {
			result = append(result, conv.Int(b[0:pos]))
			b = b[pos + len(to):]
		}
	}
}
