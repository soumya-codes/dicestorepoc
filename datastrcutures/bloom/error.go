package bloom

import "github.com/soumya-codes/AlgoAndDS/generics/diceerrors"

var (
	errInvalidErrorRateType = diceerrors.NewErr("only float values can be provided for error rate")
	errInvalidErrorRate     = diceerrors.NewErr("invalid error rate value provided")
	errInvalidCapacityType  = diceerrors.NewErr("only integer values can be provided for Capacity")
	errInvalidCapacity      = diceerrors.NewErr("invalid Capacity value provided")

	errInvalidKey = diceerrors.NewErr("invalid key: no bloom filter found")

	errEmptyValue   = diceerrors.NewErr("empty value provided")
	errUnableToHash = diceerrors.NewErr("unable to hash given value")
)
