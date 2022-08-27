package enums

type GafamType int64

const (
	Google    = 0
	Amazon    = 1
	Facebook  = 2
	Apple     = 3
	Microsoft = 4
	Other     = 5
)

func (s GafamType) String() string {
	switch s {
	case Google:
		return "Google"
	case Amazon:
		return "Amazon"
	case Facebook:
		return "Facebook"
	case Apple:
		return "Apple"
	case Microsoft:
		return "Microsoft"
	case Other:
		return "Other"
	}
	return "Unknown"
}
