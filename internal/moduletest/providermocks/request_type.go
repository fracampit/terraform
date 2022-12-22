package providermocks

type requestType rune

const (
	readRequest  requestType = 'R'
	planRequest  requestType = 'P'
	applyRequest requestType = 'A'
)

//go:generate go run golang.org/x/tools/cmd/stringer -type requestType
