//go:build go1.22

// Package codeHelpAdminGen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version (devel) DO NOT EDIT.
package codeHelpAdminGen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for ContestStatus.
const (
	CLOSED  ContestStatus = "CLOSED"
	OPEN    ContestStatus = "OPEN"
	STARTED ContestStatus = "STARTED"
)

// Defines values for Difficulty.
const (
	EASY   Difficulty = "EASY"
	HARD   Difficulty = "HARD"
	MEDIUM Difficulty = "MEDIUM"
)

// Category defines model for Category.
type Category struct {
	Name string `json:"name"`
}

// CategoryResponse defines model for CategoryResponse.
type CategoryResponse struct {
	Categories []Category `json:"categories"`
}

// Code defines model for Code.
type Code = string

// Contest defines model for Contest.
type Contest struct {
	Duration string        `json:"duration"`
	Id       int64         `json:"id"`
	Name     string        `json:"name"`
	StartsOn Date          `json:"startsOn"`
	Status   ContestStatus `json:"status"`
}

// ContestDetail defines model for ContestDetail.
type ContestDetail struct {
	Duration string           `json:"duration"`
	Id       int64            `json:"id"`
	Name     string           `json:"name"`
	Problems []ContestProblem `json:"problems"`
	StartsOn Date             `json:"startsOn"`
	Status   ContestStatus    `json:"status"`
}

// ContestEditRequest defines model for ContestEditRequest.
type ContestEditRequest struct {
	Duration string `json:"duration"`
	Name     string `json:"name"`
	Problems *[]struct {
		ContestProblemId *int64 `json:"contestProblemId,omitempty"`
		Score            *int   `json:"score,omitempty"`
	} `json:"problems,omitempty"`
	StartsOn Date           `json:"startsOn"`
	Status   *ContestStatus `json:"status,omitempty"`
}

// ContestProblem defines model for ContestProblem.
type ContestProblem struct {
	Category         *Category   `json:"category,omitempty"`
	ContestProblemId int64       `json:"contestProblemId"`
	Difficulty       *Difficulty `json:"difficulty,omitempty"`
	Id               *int64      `json:"id,omitempty"`
	Score            int         `json:"score"`
	Title            *string     `json:"title,omitempty"`
}

// ContestRequest defines model for ContestRequest.
type ContestRequest struct {
	Duration string         `json:"duration"`
	Name     string         `json:"name"`
	StartsOn Date           `json:"startsOn"`
	Status   *ContestStatus `json:"status,omitempty"`
}

// ContestResponse defines model for ContestResponse.
type ContestResponse struct {
	Contests []Contest `json:"contests"`
}

// ContestStatus defines model for ContestStatus.
type ContestStatus string

// Date defines model for Date.
type Date = openapi_types.Date

// Difficulty defines model for Difficulty.
type Difficulty string

// Problem defines model for Problem.
type Problem struct {
	Category   *Category   `json:"category,omitempty"`
	Difficulty *Difficulty `json:"difficulty,omitempty"`
	Id         *int64      `json:"id,omitempty"`
	Title      *string     `json:"title,omitempty"`
}

// ProblemDetail defines model for ProblemDetail.
type ProblemDetail struct {
	Category    *Category   `json:"category,omitempty"`
	Difficulty  *Difficulty `json:"difficulty,omitempty"`
	Id          *int64      `json:"id,omitempty"`
	Markdown    string      `json:"markdown"`
	RunnerCode  Code        `json:"runnerCode"`
	StarterCode Code        `json:"starterCode"`
	TestCases   TestCases   `json:"testCases"`
	Title       *string     `json:"title,omitempty"`
}

// ProblemRequest defines model for ProblemRequest.
type ProblemRequest struct {
	Category    *Category  `json:"category,omitempty"`
	Difficulty  Difficulty `json:"difficulty"`
	Markdown    string     `json:"markdown"`
	RunnerCode  Code       `json:"runnerCode"`
	StarterCode Code       `json:"starterCode"`
	TestCases   TestCases  `json:"testCases"`
	Title       string     `json:"title"`
}

// ProblemResponse defines model for ProblemResponse.
type ProblemResponse struct {
	Problems []Problem `json:"problems"`
}

// TestCase defines model for TestCase.
type TestCase = string

// TestCases defines model for TestCases.
type TestCases = []struct {
	Id  *int64    `json:"id,omitempty"`
	In  *TestCase `json:"in,omitempty"`
	Out *TestCase `json:"out,omitempty"`
}

// CategoryName defines model for CategoryName.
type CategoryName = string

// ContestId defines model for ContestId.
type ContestId = int64

// ProblemId defines model for ProblemId.
type ProblemId = int64

// CreateProblemParams defines parameters for CreateProblem.
type CreateProblemParams struct {
	ContestId *int `form:"contestId,omitempty" json:"contestId,omitempty"`
}

// CreateCategoryJSONRequestBody defines body for CreateCategory for application/json ContentType.
type CreateCategoryJSONRequestBody = Category

// UpdateCategoryJSONRequestBody defines body for UpdateCategory for application/json ContentType.
type UpdateCategoryJSONRequestBody = Category

// CreateContestJSONRequestBody defines body for CreateContest for application/json ContentType.
type CreateContestJSONRequestBody = ContestRequest

// UpdateContestJSONRequestBody defines body for UpdateContest for application/json ContentType.
type UpdateContestJSONRequestBody = ContestEditRequest

// CreateProblemJSONRequestBody defines body for CreateProblem for application/json ContentType.
type CreateProblemJSONRequestBody = ProblemRequest

// UpdateProblemJSONRequestBody defines body for UpdateProblem for application/json ContentType.
type UpdateProblemJSONRequestBody = ProblemRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /categories)
	GetAllCategories(w http.ResponseWriter, r *http.Request)

	// (POST /categories)
	CreateCategory(w http.ResponseWriter, r *http.Request)

	// (PUT /categories/{name})
	UpdateCategory(w http.ResponseWriter, r *http.Request, name CategoryName)

	// (GET /contests)
	GetAllContests(w http.ResponseWriter, r *http.Request)

	// (POST /contests)
	CreateContest(w http.ResponseWriter, r *http.Request)

	// (DELETE /contests/{id})
	DeleteContest(w http.ResponseWriter, r *http.Request, id ContestId)

	// (GET /contests/{id})
	GetContest(w http.ResponseWriter, r *http.Request, id ContestId)

	// (PUT /contests/{id})
	UpdateContest(w http.ResponseWriter, r *http.Request, id ContestId)

	// (GET /problems)
	GetAllProblems(w http.ResponseWriter, r *http.Request)

	// (POST /problems)
	CreateProblem(w http.ResponseWriter, r *http.Request, params CreateProblemParams)

	// (DELETE /problems/{id})
	DeleteProblem(w http.ResponseWriter, r *http.Request, id ProblemId)

	// (GET /problems/{id})
	GetProblem(w http.ResponseWriter, r *http.Request, id ProblemId)

	// (PUT /problems/{id})
	UpdateProblem(w http.ResponseWriter, r *http.Request, id ProblemId)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetAllCategories operation middleware
func (siw *ServerInterfaceWrapper) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAllCategories(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateCategory operation middleware
func (siw *ServerInterfaceWrapper) CreateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateCategory(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateCategory operation middleware
func (siw *ServerInterfaceWrapper) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name CategoryName

	err = runtime.BindStyledParameterWithOptions("simple", "name", r.PathValue("name"), &name, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateCategory(w, r, name)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetAllContests operation middleware
func (siw *ServerInterfaceWrapper) GetAllContests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAllContests(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateContest operation middleware
func (siw *ServerInterfaceWrapper) CreateContest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateContest(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteContest operation middleware
func (siw *ServerInterfaceWrapper) DeleteContest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id ContestId

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteContest(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetContest operation middleware
func (siw *ServerInterfaceWrapper) GetContest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id ContestId

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetContest(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateContest operation middleware
func (siw *ServerInterfaceWrapper) UpdateContest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id ContestId

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateContest(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetAllProblems operation middleware
func (siw *ServerInterfaceWrapper) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetAllProblems(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateProblem operation middleware
func (siw *ServerInterfaceWrapper) CreateProblem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateProblemParams

	// ------------- Optional query parameter "contestId" -------------

	err = runtime.BindQueryParameter("form", true, false, "contestId", r.URL.Query(), &params.ContestId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "contestId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateProblem(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteProblem operation middleware
func (siw *ServerInterfaceWrapper) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id ProblemId

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteProblem(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetProblem operation middleware
func (siw *ServerInterfaceWrapper) GetProblem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id ProblemId

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProblem(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateProblem operation middleware
func (siw *ServerInterfaceWrapper) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id ProblemId

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateProblem(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       *http.ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m *http.ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m *http.ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("GET "+options.BaseURL+"/categories", wrapper.GetAllCategories)
	m.HandleFunc("POST "+options.BaseURL+"/categories", wrapper.CreateCategory)
	m.HandleFunc("PUT "+options.BaseURL+"/categories/{name}", wrapper.UpdateCategory)
	m.HandleFunc("GET "+options.BaseURL+"/contests", wrapper.GetAllContests)
	m.HandleFunc("POST "+options.BaseURL+"/contests", wrapper.CreateContest)
	m.HandleFunc("DELETE "+options.BaseURL+"/contests/{id}", wrapper.DeleteContest)
	m.HandleFunc("GET "+options.BaseURL+"/contests/{id}", wrapper.GetContest)
	m.HandleFunc("PUT "+options.BaseURL+"/contests/{id}", wrapper.UpdateContest)
	m.HandleFunc("GET "+options.BaseURL+"/problems", wrapper.GetAllProblems)
	m.HandleFunc("POST "+options.BaseURL+"/problems", wrapper.CreateProblem)
	m.HandleFunc("DELETE "+options.BaseURL+"/problems/{id}", wrapper.DeleteProblem)
	m.HandleFunc("GET "+options.BaseURL+"/problems/{id}", wrapper.GetProblem)
	m.HandleFunc("PUT "+options.BaseURL+"/problems/{id}", wrapper.UpdateProblem)

	return m
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RYT2/buBP9KgT7O/wWUC05TXvwzbHd1thtHMTpYdHkwEhjm61EKhSVrhHouy9E0hL1",
	"z5YXdorFnlo5w+HMvPeGQ75gn0cxZ8BkgkcvOCaCRCBBqK8JkbDmYntNIsi/KcMjHBO5wQ5m6jf9j4MF",
	"PKVUQIBHUqTg4MTfQETyNXIb53aJFJStcZY5eMKZhETOgw6XNNjrcMVFRGRux+SHS+zsdqBMwhqE2uJG",
	"8McQovNtke3MK3VSFRQ8BiEpqL8wUzn4i0RxmPu4ooyILZICICk9W+Upw/qmlz8UVvzxO/gS5zU0O95C",
	"EnOWQHNnX1uYLyohUv/5n4AVHuE3bom7azJxizSyYkciBNk2wrJ8twbHg1rSq5T5knKG/k/Eevgberln",
	"CLku8nkAaM0hQRsQcM+ye9asScGYZpJBKkjut7rbcPPOi7BTwliYtfimQS/EnRYoTVjIaKDhOpFEyGTB",
	"DtV9SiQYe5keRknvutTGdWgUs01Axp9j51/E1A6ccj0FSWiYx0HCcLHCo2+9IsKZU8cn1jo8goLaldHv",
	"QSIW/pvZPJT5zAIqb+EpNRQ6Kqndup651TRYyWbel2mJzwVYvdPubA3MGvXprsOuqr1rUMLwCplVGkzd",
	"/W7pXpwtjE/eJv4d6jexHKn3PYeINjhav4dPkJ3jPWEti4IAS6N81eJmdo0dvLwb397NptjBkz8Wy9nU",
	"clKWX5W1AtiFd3Hx1nv/dvihAnpu17aerlbUT0O5tSOYjZd/Ygd/mU3nX79gB38e36rty12MRcOfJb7W",
	"g3p7zOkcVILby67S8pjDTlIZQvv41oDLZHbsqdHdXCIifgT8Z02ub96gu58cLdNoMBgg031RAIkvaNwl",
	"W5EyBmI3kOxnbgCFbI9bkpN1QhI4qI67wrCuhyLlagCVBOyN2vugKWlnH3xNqv0XQLR0UqaY55ek0cHZ",
	"Xi+tFNg5AQ8sFnR19aOnshOMYw7e1a1arfHVZIo+zq+n6NPink0mV9ZXGxfubJg6xq7eXY6yvmDn1jyV",
	"/c17jGoqgBVXXVbTSN2c0GcIYzQOIpoX4BlEoiYXPBx4A0/FEQMjMcUj/G7gDYbYUZdblblbvfWtQUWc",
	"10bNA/mchj+BHIfhpDTMMdRUUYsuPK84+pkemOM4pL7y4H5P9BhVXpb7NJKCiyprW+4jvPhdV4usE+tu",
	"ucV5Q4t50pLBRACRUDQpzUFI5BUPticPPQ85a5TonRorK3noqIJBVzaZY8PjvuSzWqbIm7bk+DUOqjna",
	"LzMdJ2tp4lZebrKHX1Giy2aJrjkyGw9yIl9qplVtrkiATLDGqNWRRCuesgPVtqbXfVLYmZ1TCLVJu1UH",
	"AzQOQ1REXcnMzNUHRWHszgR47U7cUxkl7B052WC5LzTItI8Q9AhfTXSqfi8TPVIYxeOjVkV/ynbi0UWt",
	"s4R4ckqasb2DkMV9rouLe3rXidI/G43tZ6FWKnv/kMr2fLWn79zszM4Icn0WPHz+muArnabtoEOkGN8p",
	"Q3ID6JH4P4AF2GntTLsZskEH9Tz/lII65cz7vF9wwH6WL6bGi5YX+TNxpXalauXJ8Kjd6reCBiBmkOhA",
	"xSZYz17ZWfoDSiyf33r3ykmLOGxGdcnhLCGeXEb7euXeWeVjPqvkFu/bmsqcSRCMhCgB8QwCgRBcdIuy",
	"u+Weqoq/RkavCJcuV6fEMgdrKHT1alDHwND4Zo6SGHy6MhFiB6cixCO8kTIeuW7IfRJueCJHl57nuSSm",
	"7vMQZw/Z3wEAAP//sODFGnIdAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
