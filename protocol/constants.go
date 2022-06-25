package protocol

import (
	"time"
)

// Response codes sent as DNS answers.
const (
	SuccessDNSResponse = "1.1.1.1"
	FailureDNSResponse = "1.1.1.2"
)

// TXT record default responses
var (
	NoCmdTxtResponse  = "v=B2B3FE1C"
	ErrorTxtResponse  = "v=D31CFAA4"
	CmdTxtResponse    = "v=A9F466E8"
	UploadTxtResponse = "v=H34FERKL"
)

// MaxLabelSize is the maximum size a DNS hostname label may be.
const MaxLabelSize = 63

// Protocols understood
const (
	FileProtocol = iota
	CmdProtocol
	UploadProtocol
	CobaltStrikeProtocol
)

// PollTypes to expect from agents.
const (
	PollTypeUndefined = iota
	PollTypeCheckin
	PollTypeUpload
)

// Request stream status
const (
	StreamStart = 0xbe
	StreamData  = 0xef
	StreamEnd   = 0xca
)

// IncomingDNSBuffer represents a pending incoming DNS conversation
type IncomingDNSBuffer struct {
	Identifier string
	Data       []byte
	Seq        int
	Started    bool
	Finished   bool
	Protocol   int
}

// Agent represents an agent connected to this C2
type Agent struct {
	Identifier   string
	FirstCheckin time.Time
	LastCheckin  time.Time
}
