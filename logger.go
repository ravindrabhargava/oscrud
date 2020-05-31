package oscrud

// Logger :
type Logger interface {
	StartRequest(Context)
	Log(operation string, content string)
	EndRequest(Context)
}
