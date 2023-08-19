package constants

const traceId string = "trace-id"

type Headers struct {
	TraceId string
}

var GetHeaders = Headers{
	TraceId: traceId,
}
