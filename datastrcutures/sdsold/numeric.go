package sdsold

import (
	"errors"
	ds "github.com/soumya-codes/AlgoAndDS/generics/datastrcutures"
	"golang.org/x/exp/constraints"
	"strconv"
)

// Numeric constraint only for numeric types
type numericConstraint interface {
	constraints.Integer | constraints.Float
}

// NumericSDS is a specific implementation of DSInterface that holds numeric values
type NumericSDS[T numericConstraint] struct {
	ds.BaseDataStructure[ds.DSInterface]
	value T
}

func NewNumericSDS[T numericConstraint](val T) *NumericSDS[T] {
	return &NumericSDS[T]{value: val}
}

// Get method optimized for each numeric type
func (s *NumericSDS[T]) Get() string {
	switch v := any(s.value).(type) {
	case int:
		return strconv.Itoa(v)
	case int8, int16, int32:
		return strconv.FormatInt(int64(v.(int)), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8, uint16, uint32:
		return strconv.FormatUint(uint64(v.(uint)), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return ""
	}
}

func (s *NumericSDS[T]) Set(val string) error {
	var newValue T
	u, err := strconv.ParseUint(val, 10, 64)
	if err == nil {
		switch any(newValue).(type) {
		case uint8:
			s.value = T(uint8(u))
		case uint16:
			s.value = T(uint16(u))
		case uint32:
			s.value = T(uint32(u))
		case uint64:
			s.value = T(u)
		default:
			return errors.New("unsupported type")
		}
		return nil
	}

	i, err := strconv.ParseInt(val, 10, 64)
	if err == nil {
		switch any(newValue).(type) {
		case int8:
			s.value = T(int8(i))
		case int16:
			s.value = T(int16(i))
		case int32:
			s.value = T(int32(i))
		case int64:
			s.value = T(i)
		default:
			return errors.New("unsupported type")
		}
		return nil
	}

	f, err := strconv.ParseFloat(val, 64)
	if err == nil {
		switch any(newValue).(type) {
		case float32:
			s.value = T(float32(f))
		case float64:
			s.value = T(f)
		default:
			return errors.New("unsupported type")
		}
		return nil
	}

	return errors.New("invalid value")
}

func (s *NumericSDS[T]) Incr() {
	s.value++
}

func (s *NumericSDS[T]) IncrBy(i T) {
	s.value += i
}

func (s *NumericSDS[T]) Decr() {
	s.value--
}

func (s *NumericSDS[T]) DecrBy(i T) {
	s.value -= i
}
