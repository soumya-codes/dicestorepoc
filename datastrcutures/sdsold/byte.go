package sdsold

import ds "github.com/soumya-codes/AlgoAndDS/generics/datastrcutures"

// ByteSDS is a specific implementation of DSInterface that holds a byte slice
type ByteSDS struct {
	ds.BaseDataStructure[ds.DSInterface]
	value []byte
}

func NewByteSDS(s []byte) *ByteSDS {
	return &ByteSDS{value: s}
}

func (s *ByteSDS) Get() string {
	return string(s.value)
}

func (s *ByteSDS) Set(val string) error {
	s.value = []byte(val)
	return nil
}
