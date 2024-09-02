package sdsold

import ds "github.com/soumya-codes/AlgoAndDS/generics/datastrcutures"

type SDSType interface {
	ds.DSInterface
	Set(val string) error
	Get() string
}

func GetSDS(ds ds.DSInterface) (SDSType, bool) {
	sds, ok := ds.(SDSType)
	return sds, ok
}
