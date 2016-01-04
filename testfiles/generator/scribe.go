// This file is automatically generated. Do not modify.

package gentest

import (
	"fmt"
	"strconv"
)

var _ = fmt.Sprintf

type ResultCode int32

const (
	ResultCodeOk       ResultCode = 0
	ResultCodeTryLater ResultCode = 1
)

var (
	ResultCodeByName = map[string]ResultCode{
		"ResultCode.OK":        ResultCodeOk,
		"ResultCode.TRY_LATER": ResultCodeTryLater,
	}
	ResultCodeByValue = map[ResultCode]string{
		ResultCodeOk:       "ResultCode.OK",
		ResultCodeTryLater: "ResultCode.TRY_LATER",
	}
)

func (e ResultCode) String() string {
	name := ResultCodeByValue[e]
	if name == "" {
		name = fmt.Sprintf("Unknown enum value ResultCode(%d)", e)
	}
	return name
}

func (e ResultCode) MarshalJSON() ([]byte, error) {
	name := ResultCodeByValue[e]
	if name == "" {
		name = strconv.Itoa(int(e))
	}
	return []byte("\"" + name + "\""), nil
}

func (e *ResultCode) UnmarshalJSON(b []byte) error {
	st := string(b)
	if st[0] == '"' {
		*e = ResultCode(ResultCodeByName[st[1:len(st)-1]])
		return nil
	}
	i, err := strconv.Atoi(st)
	*e = ResultCode(i)
	return err
}

func (e ResultCode) Ptr() *ResultCode {
	return &e
}

type LogEntry struct {
	Category *string `thrift:"1,required" json:"category"`
	Message  *string `thrift:"2,required" json:"message"`
}

func (l *LogEntry) GetCategory() (val string) {
	if l != nil && l.Category != nil {
		return *l.Category
	}

	return
}

func (l *LogEntry) GetMessage() (val string) {
	if l != nil && l.Message != nil {
		return *l.Message
	}

	return
}

type FailedException struct {
	Reason *string `thrift:"1,required" json:"reason"`
}

func (f *FailedException) GetReason() (val string) {
	if f != nil && f.Reason != nil {
		return *f.Reason
	}

	return
}

func (e FailedException) Error() string {
	return fmt.Sprintf("FailedException{Reason: %+v}", e.Reason)
}

type Scribe interface {
	Echo(messages *LogEntry) (*LogEntry, error)
	Log(messages []LogEntry) (*ResultCode, error)
}

type ScribeServer struct {
	Implementation Scribe
}

func (s *ScribeServer) Echo(req *ScribeEchoRequest, res *ScribeEchoResponse) error {
	val, err := s.Implementation.Echo(req.Messages)
	switch e := err.(type) {
	case *FailedException:
		res.F = e
		err = nil
	}
	res.Value = val
	return err
}

func (s *ScribeServer) Log(req *ScribeLogRequest, res *ScribeLogResponse) error {
	val, err := s.Implementation.Log(req.Messages)
	switch e := err.(type) {
	case *FailedException:
		res.F = e
		err = nil
	}
	res.Value = val
	return err
}

type ScribeEchoRequest struct {
	Messages *LogEntry `thrift:"1,required" json:"messages"`
}

func (s *ScribeEchoRequest) GetMessages() (val LogEntry) {
	if s != nil && s.Messages != nil {
		return *s.Messages
	}

	return
}

type ScribeEchoResponse struct {
	Value *LogEntry        `thrift:"0" json:"value,omitempty"`
	F     *FailedException `thrift:"1" json:"f,omitempty"`
}

func (s *ScribeEchoResponse) GetValue() (val LogEntry) {
	if s != nil && s.Value != nil {
		return *s.Value
	}

	return
}

func (s *ScribeEchoResponse) GetF() (val FailedException) {
	if s != nil && s.F != nil {
		return *s.F
	}

	return
}

type ScribeLogRequest struct {
	Messages []LogEntry `thrift:"1,required" json:"messages"`
}

func (s *ScribeLogRequest) GetMessages() (val []LogEntry) {
	if s != nil {
		return s.Messages
	}

	return
}

type ScribeLogResponse struct {
	Value *ResultCode      `thrift:"0" json:"value,omitempty"`
	F     *FailedException `thrift:"1" json:"f,omitempty"`
}

func (s *ScribeLogResponse) GetValue() (val ResultCode) {
	if s != nil && s.Value != nil {
		return *s.Value
	}

	return
}

func (s *ScribeLogResponse) GetF() (val FailedException) {
	if s != nil && s.F != nil {
		return *s.F
	}

	return
}

type ScribeClient struct {
	Client RPCClient
}

func (s *ScribeClient) Echo(messages *LogEntry) (ret *LogEntry, err error) {
	req := &ScribeEchoRequest{
		Messages: messages,
	}
	res := &ScribeEchoResponse{}
	err = s.Client.Call("Echo", req, res)
	if err == nil {
		switch {
		case res.F != nil:
			err = res.F
		}
	}
	if err == nil {
		ret = res.Value
	}
	return
}

func (s *ScribeClient) Log(messages []LogEntry) (ret *ResultCode, err error) {
	req := &ScribeLogRequest{
		Messages: messages,
	}
	res := &ScribeLogResponse{}
	err = s.Client.Call("Log", req, res)
	if err == nil {
		switch {
		case res.F != nil:
			err = res.F
		}
	}
	if err == nil {
		ret = res.Value
	}
	return
}
