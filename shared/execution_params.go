package shared

type ExecutionParams struct {
	EntryPoint string `json:"entryPoint"`
	Verbose    bool   `json:"verbose"`
}

type ExecutionResult struct {
	Error       string `json:"error"`
	ElapsedTime int64  `json:"elapsedTime"`
}
