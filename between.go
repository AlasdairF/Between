package between

import (
 "bytes"
 "strconv"
 "errors"
)

// Bytes returns the result as a slice of bytes.
func Bytes(b, from, to []byte) ([]byte, error) {
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

// AllInts returns all results between from & to.
func AllBytes(b, from, to []byte) [][]byte {
	result := make([][]byte, 0)
	l := len(to)
	for {
		s := bytes.Index(b, from)
		if s==-1 {
			return result
		}
		s += len(from)
		e := bytes.Index(b[s:], to)
		if e==-1 {
				return result
			}
		f := s + e
		result = append(result,b[s:f])
		b = b[f+l:]
	}
}

// Int returns the result as an int.
func Int(b, from, to []byte) (int, error) {
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

// AllInt returns all results between from & to as int.
func AllInt(b, from, to []byte) []int {
	result := make([]int, 0)
	l := len(to)
	for {
		s := bytes.Index(b, from)
		if s==-1 {
			return result
		}
		s += len(from)
		e := bytes.Index(b[s:], to)
		if e==-1 {
				return result
			}
		f := s + e
		v, _ := strconv.ParseInt(string(b[s:f]),10,64)
		result = append(result,int(v))
		b = b[f+l:]
	}
}

// Float returns the result as a float64.
func Float(b, from, to []byte) (float64, error) {
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

// AllFloat returns all results between from & to as float64.
func AllFloat(b, from, to []byte) []float64 {
	result := make([]float64, 0)
	l := len(to)
	for {
		s := bytes.Index(b, from)
		if s==-1 {
			return result
		}
		s += len(from)
		e := bytes.Index(b[s:], to)
		if e==-1 {
				return result
			}
		f := s + e
		v, _ := strconv.ParseFloat(string(b[s:s+e]),64)
		result = append(result,v)
		b = b[f+l:]
	}
}
