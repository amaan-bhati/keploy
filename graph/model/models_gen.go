// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type App struct {
	ID string `json:"id"`
}

type BodyResult struct {
	Normal   bool         `json:"normal"`
	Type     BodyType     `json:"type"`
	Expected string       `json:"expected"`
	Actual   string       `json:"actual"`
	Errors   []*JSONError `json:"errors"`
}

type DepMetaResult struct {
	Normal   *bool   `json:"normal"`
	Key      *string `json:"key"`
	Expected *string `json:"expected"`
	Actual   *string `json:"actual"`
}

type DepResult struct {
	Name string           `json:"name"`
	Type DependencyType   `json:"type"`
	Meta []*DepMetaResult `json:"meta"`
}

type Dependency struct {
	Name string         `json:"name"`
	Type DependencyType `json:"type"`
	Meta []*Kv          `json:"meta"`
}

type DependencyInput struct {
	Name string         `json:"name"`
	Type DependencyType `json:"type"`
	Meta []*KVInput     `json:"meta"`
}

type HTTPReq struct {
	ProtoMajor int       `json:"protoMajor"`
	ProtoMinor int       `json:"protoMinor"`
	URL        *string   `json:"url"`
	URLParam   []*Kv     `json:"urlParam"`
	Header     []*Header `json:"header"`
	Method     Method    `json:"method"`
	Body       string    `json:"body"`
}

type HTTPReqInput struct {
	ProtoMajor *int           `json:"protoMajor"`
	ProtoMinor *int           `json:"protoMinor"`
	URL        *string        `json:"url"`
	URLParam   []*KVInput     `json:"urlParam"`
	Header     []*HeaderInput `json:"header"`
	Method     *Method        `json:"method"`
	Body       *string        `json:"body"`
}

type HTTPResp struct {
	StatusCode int       `json:"statusCode"`
	Header     []*Header `json:"header"`
	Body       string    `json:"body"`
}

type HTTPRespInput struct {
	StatusCode *int           `json:"statusCode"`
	Header     []*HeaderInput `json:"header"`
	Body       *string        `json:"body"`
}

type Header struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

type HeaderInput struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

type HeaderResult struct {
	Normal   *bool   `json:"normal"`
	Key      string  `json:"key"`
	Expected *Header `json:"expected"`
	Actual   *Header `json:"actual"`
}

type IntResult struct {
	Normal   *bool `json:"normal"`
	Expected int   `json:"expected"`
	Actual   int   `json:"actual"`
}

type JSONError struct {
	Key               string `json:"key"`
	MissingInExpected bool   `json:"missingInExpected"`
	MissingInActual   bool   `json:"missingInActual"`
}

type KVInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Result struct {
	StatusCode    *IntResult      `json:"statusCode"`
	HeadersResult []*HeaderResult `json:"headersResult"`
	BodyResult    *BodyResult     `json:"bodyResult"`
	DepResult     []*DepResult    `json:"depResult"`
}

type Test struct {
	ID         string        `json:"id"`
	Status     TestStatus    `json:"status"`
	Started    time.Time     `json:"started"`
	Completed  *time.Time    `json:"completed"`
	Result     *Result       `json:"result"`
	TestCaseID string        `json:"testCaseID"`
	URI        *string       `json:"uri"`
	Req        *HTTPReq      `json:"req"`
	Deps       []*Dependency `json:"deps"`
	Noise      []string      `json:"noise"`
}

type TestCase struct {
	ID       string        `json:"id"`
	Created  time.Time     `json:"created"`
	Updated  time.Time     `json:"updated"`
	Captured time.Time     `json:"captured"`
	Cid      string        `json:"cid"`
	App      string        `json:"app"`
	URI      string        `json:"uri"`
	HTTPReq  *HTTPReq      `json:"httpReq"`
	HTTPResp *HTTPResp     `json:"httpResp"`
	Deps     []*Dependency `json:"deps"`
	Anchors  []string      `json:"anchors"`
	Noise    []string      `json:"noise"`
}

type TestCaseInput struct {
	ID       string             `json:"id"`
	Created  *time.Time         `json:"created"`
	Updated  *time.Time         `json:"updated"`
	Captured *time.Time         `json:"captured"`
	Cid      *string            `json:"cid"`
	App      *string            `json:"app"`
	URI      *string            `json:"uri"`
	HTTPReq  *HTTPReqInput      `json:"httpReq"`
	HTTPResp *HTTPRespInput     `json:"httpResp"`
	Deps     []*DependencyInput `json:"deps"`
	Anchors  []string           `json:"anchors"`
	Noise    []string           `json:"noise"`
}

type TestCases struct {
	Tc    []*TestCase `json:"tc"`
	Count int         `json:"count"`
}

type TestRun struct {
	ID      string        `json:"id"`
	Created time.Time     `json:"created"`
	Updated time.Time     `json:"updated"`
	Status  TestRunStatus `json:"status"`
	App     string        `json:"app"`
	User    string        `json:"user"`
	Success int           `json:"success"`
	Failure int           `json:"failure"`
	Total   int           `json:"total"`
	Tests   []*Test       `json:"tests"`
}

type BodyType string

const (
	BodyTypePlain BodyType = "PLAIN"
	BodyTypeJSON  BodyType = "JSON"
)

var AllBodyType = []BodyType{
	BodyTypePlain,
	BodyTypeJSON,
}

func (e BodyType) IsValid() bool {
	switch e {
	case BodyTypePlain, BodyTypeJSON:
		return true
	}
	return false
}

func (e BodyType) String() string {
	return string(e)
}

func (e *BodyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BodyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BodyType", str)
	}
	return nil
}

func (e BodyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DependencyType string

const (
	DependencyTypeNoSQLDb DependencyType = "NO_SQL_DB"
	DependencyTypeSQLDb   DependencyType = "SQL_DB"
)

var AllDependencyType = []DependencyType{
	DependencyTypeNoSQLDb,
	DependencyTypeSQLDb,
}

func (e DependencyType) IsValid() bool {
	switch e {
	case DependencyTypeNoSQLDb, DependencyTypeSQLDb:
		return true
	}
	return false
}

func (e DependencyType) String() string {
	return string(e)
}

func (e *DependencyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DependencyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DependencyType", str)
	}
	return nil
}

func (e DependencyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Method string

const (
	MethodGet     Method = "GET"
	MethodPut     Method = "PUT"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPatch   Method = "PATCH"
	MethodDelete  Method = "DELETE"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
)

var AllMethod = []Method{
	MethodGet,
	MethodPut,
	MethodHead,
	MethodPost,
	MethodPatch,
	MethodDelete,
	MethodOptions,
	MethodTrace,
}

func (e Method) IsValid() bool {
	switch e {
	case MethodGet, MethodPut, MethodHead, MethodPost, MethodPatch, MethodDelete, MethodOptions, MethodTrace:
		return true
	}
	return false
}

func (e Method) String() string {
	return string(e)
}

func (e *Method) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Method(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Method", str)
	}
	return nil
}

func (e Method) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TestRunStatus string

const (
	TestRunStatusRunning TestRunStatus = "RUNNING"
	TestRunStatusFailed  TestRunStatus = "FAILED"
	TestRunStatusPassed  TestRunStatus = "PASSED"
)

var AllTestRunStatus = []TestRunStatus{
	TestRunStatusRunning,
	TestRunStatusFailed,
	TestRunStatusPassed,
}

func (e TestRunStatus) IsValid() bool {
	switch e {
	case TestRunStatusRunning, TestRunStatusFailed, TestRunStatusPassed:
		return true
	}
	return false
}

func (e TestRunStatus) String() string {
	return string(e)
}

func (e *TestRunStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TestRunStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TestRunStatus", str)
	}
	return nil
}

func (e TestRunStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TestStatus string

const (
	TestStatusPending TestStatus = "PENDING"
	TestStatusRunning TestStatus = "RUNNING"
	TestStatusFailed  TestStatus = "FAILED"
	TestStatusPassed  TestStatus = "PASSED"
)

var AllTestStatus = []TestStatus{
	TestStatusPending,
	TestStatusRunning,
	TestStatusFailed,
	TestStatusPassed,
}

func (e TestStatus) IsValid() bool {
	switch e {
	case TestStatusPending, TestStatusRunning, TestStatusFailed, TestStatusPassed:
		return true
	}
	return false
}

func (e TestStatus) String() string {
	return string(e)
}

func (e *TestStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TestStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TestStatus", str)
	}
	return nil
}

func (e TestStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
