package shard

import (
	"github.com/dolthub/swiss"
	ds "github.com/soumya-codes/AlgoAndDS/generics/datastrcutures"
	"github.com/soumya-codes/AlgoAndDS/generics/datastrcutures/sds"
	"github.com/soumya-codes/AlgoAndDS/generics/ops"
	"github.com/soumya-codes/AlgoAndDS/generics/store"
)

type Shard struct {
	id        int                           // Unique identifier for the shard
	store     *store.Store[ds.DSInterface]  // Store to hold data
	reqChan   chan *ops.Request             // Common channel to receive requests from workers
	workerMap map[string]chan *ops.Response // Map of WorkerID to their unique response channel
	errorChan chan *ops.Error               // Channel to send system-level errors to workers or server
}

func NewShard(id int) *Shard {
	return &Shard{
		id: id,
		store: &store.Store[ds.DSInterface]{
			Store:   swiss.NewMap[string, ds.DSInterface](42),
			Expires: swiss.NewMap[*ds.DSInterface, uint64](42),
		},
		reqChan:   make(chan *ops.Request),
		workerMap: make(map[string]chan *ops.Response),
		errorChan: make(chan *ops.Error),
	}
}

func (s *Shard) EvaluateRequest(req *ops.Request) []byte {
	// Evaluate the request and send the response back to the worker
	response := make([]byte, 0)
	oSet := req.OperationSet
	for _, op := range oSet {
		switch op.Cmd {
		// All the operations for the SDS data-structure
		case "SET", "GET":
			response = append(response, sds.NewEval(s.store, op).Evaluate()...)
		// All the operations for the BloomFilter data-structure
		case "BF.ADD", "BF.EXISTS":
			response = append(response, sds.NewEval(s.store, op).Evaluate()...)
		// All generic operations
		case "DEL":
		}
	}

	return response
}
