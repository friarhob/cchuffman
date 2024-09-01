package exitcodes

type ExitCodes int

const (
	OK ExitCodes = iota
	ErrorReadingFile
	UsageError
)

var names = [...]string{
	"OK",
	"ErrorReadingFile",
	"UsageError",
}

func (e ExitCodes) String() string {
	return names[e]
}
