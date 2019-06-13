package arg

import (
	"encoding/json"
)

var (
	Database  string
	MySQL     string
	Out       string
	SshTunnel string
	Table     string
	Module    string
	TmplDir   string
)

type CmdDt struct {
	Data json.RawMessage `json:"data"`
}
