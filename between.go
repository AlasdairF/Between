package between

import (
 "bytes"
 "strconv"
 "errors"
)

// Bytes returns the result as a slice of bytes.
func Bytes(from, to, b []byte) ([]byte, error) {
	s := bytes.Index(b, from)
	if s==-1 {
		return nil, errors.New("From does not exist.")
	}
	s += len(from)
	e := bytes.Index(b[s:], to)
	if e==-1 {
			return nil, errors.New("To does not exist.")
		}
	return b[s:s+e], nil
}

// Int returns the result as an int.
func Int(from, to, b []byte) (int, error) {
	s := bytes.Index(b, from)
	if s==-1 {
		return -1, errors.New("From does not exist.")
	}
	s += len(from)
	e := bytes.Index(b[s:], to)
	if e==-1 {
			return -1, errors.New("To does not exist.")
		}
	v, err := strconv.ParseInt(string(b[s:s+e]),10,64)
	if err!=nil {
		return -1, errors.New("Result not an int.")
	}
	return int(v), nil
}

// Float returns the result as a float64.
func Float(from, to, b []byte) (float64, error) {
	s := bytes.Index(b, from)
	if s==-1 {
		return -1, errors.New("From does not exist.")
	}
	s += len(from)
	e := bytes.Index(b[s:], to)
	if e==-1 {
			return -1, errors.New("To does not exist.")
		}
	v, err := strconv.ParseFloat(string(b[s:s+e]),64)
	if err!=nil {
		return -1, errors.New("Result not a float.")
	}
	return v, nil
}
