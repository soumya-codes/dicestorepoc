package ops

type Response struct {
	requestID string   // The ID of the request
	result    []string // The result of the command
}

type Error struct {
	ShardID  int    // The ID of the shard where the error occurred
	ErrorMsg string // A description of the error
}

type Operation struct {
	Cmd  string
	Args []string
}

type Request struct {
	RequestID    string       // The request ID
	OperationSet []*Operation // Put of commands to be executed atomically
	ShardID      int          // The ID of the shard on which the store commands will be executed
	WorkerID     string       // The ID of the worker that sent this store operation
}
