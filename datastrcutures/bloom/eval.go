package bloom

import (
	ds "github.com/soumya-codes/AlgoAndDS/generics/datastrcutures"
	"github.com/soumya-codes/AlgoAndDS/generics/diceerrors"
	"github.com/soumya-codes/AlgoAndDS/generics/ops"
	"github.com/soumya-codes/AlgoAndDS/generics/store"
)

const (
	defaultExpiry = -1
)

type Eval struct {
	store *store.Store[ds.DSInterface]
	op    *ops.Operation
}

func NewEval(store *store.Store[ds.DSInterface], op *ops.Operation) *Eval {
	return &Eval{
		store: store,
		op:    op,
	}
}

func (e *Eval) Evaluate() []byte {
	switch e.op.Cmd {
	case "BF.ADD":
		// Put the value in the store with the most memory efficient precision type
		return e.BFADD(e.op.Args)
	}

	return nil
}

func (e *Eval) BFADD(args []string) []byte {
	if len(args) != 2 {
		return diceerrors.NewErrArity("BFADD")
	}

	opts, _ := NewBloomOpts(args[1:], true)

	bloom := NewBloomFilter(opts)

	resp, err := bloom.Add(args[1])
	if err != nil {
		return diceerrors.NewErrWithFormattedMessage("%w for 'BFADD' command", err)
	}

	return resp
}

//func (e *Eval) get() []byte {
//	args := e.op.Args
//	if len(args) < 1 {
//		return diceerrors.NewErrArity("GET")
//	}
//
//	key := args[0]
//	val, ok := e.store.get(key)
//	if !ok {
//		return []byte("-1")
//	}
//
//	var sVal string
//	switch v := val.(type) {
//	case *String[uint8], *String[uint16], *String[uint32], *String[uint64],
//		*String[int8], *String[int16], *String[int32], *String[int64],
//		*String[float32], *String[float64], *String[[]byte], *String[string]:
//		sVal = v.get()
//	default:
//		sVal = ""
//	}
//
//	return []byte(sVal)
//}
