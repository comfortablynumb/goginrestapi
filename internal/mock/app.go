package mock

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	"github.com/comfortablynumb/goginrestapi/internal/app"
	"github.com/comfortablynumb/goginrestapi/internal/config"
	jsoniter "github.com/json-iterator/go"
)

// Structs

// MockAppOptions

type MockAppOptions struct {
	Headers          http.Header
	Body             interface{}
	ExpectedResponse interface{}
}

func (m *MockAppOptions) WithHeader(header string, value string) *MockAppOptions {
	m.Headers.Add(header, value)

	return m
}

func (m *MockAppOptions) WithBody(body interface{}) *MockAppOptions {
	m.Body = body

	return m
}

func (m *MockAppOptions) WithExpectedResponse(expectedResponse interface{}) *MockAppOptions {
	m.ExpectedResponse = expectedResponse

	return m
}

// MockApp

type MockApp struct {
	App app.App
}

func (m *MockApp) NewGetRequest(uri string, options *MockAppOptions) (*httptest.ResponseRecorder, error) {
	return m.NewRequest(http.MethodGet, uri, options)
}

func (m *MockApp) NewPostRequest(uri string, options *MockAppOptions) (*httptest.ResponseRecorder, error) {
	return m.NewRequest(http.MethodPost, uri, options)
}

func (m *MockApp) NewPutRequest(uri string, options *MockAppOptions) (*httptest.ResponseRecorder, error) {
	return m.NewRequest(http.MethodPut, uri, options)
}

func (m *MockApp) NewDeleteRequest(uri string, options *MockAppOptions) (*httptest.ResponseRecorder, error) {
	return m.NewRequest(http.MethodDelete, uri, options)
}

func (m *MockApp) NewRequest(method string, uri string, options *MockAppOptions) (*httptest.ResponseRecorder, error) {
	w := httptest.NewRecorder()
	var bodyReader io.Reader
	var headers http.Header

	if options == nil {
		options = NewMockAppOptions()
	}

	// Headers

	if method != http.MethodGet {
		contentType := options.Headers.Get("Content-Type")

		if contentType == "" {
			contentType = "application/json"
		}

		headers = make(map[string][]string)

		headers.Add("Content-Type", contentType)
	}

	// Body

	body := options.Body

	switch body.(type) {
	case nil:
		bodyReader = nil
	case string:
		bodyReader = strings.NewReader(body.(string))
	default:
		jsonBytes, err := jsoniter.Marshal(body)

		if err != nil {
			panic(err)
		}

		bodyReader = strings.NewReader(string(jsonBytes))
	}

	// Request!

	req, err := http.NewRequest(method, uri, bodyReader)

	if err != nil {
		return nil, err
	}

	req.Header = headers

	m.App.GetRouter().ServeHTTP(w, req)

	// Response

	if options.ExpectedResponse != nil {
		err := jsoniter.Unmarshal(w.Body.Bytes(), options.ExpectedResponse)

		if err != nil {
			return nil, err
		}
	}

	return w, nil
}

// Static functions

func NewMockApp(config *config.AppConfig) *MockApp {
	mockApp := &MockApp{
		App: app.NewApp(config),
	}

	mockApp.App.SetUp()

	return mockApp
}

func NewMockAppWithDefaultConfig() *MockApp {
	return NewMockApp(
		&config.AppConfig{
			Port:             8080,
			LogLevel:         "DEBUG",
			DbUri:            "file:test.db?cache=shared&mode=memory",
			DbMigrationsPath: fmt.Sprintf("file://%s", GetMigrationsAbsolutePath()),
			DefaultLocale:    "en",
			DefaultLimit:     50,
		},
	)
}

// GetMigrationsAbsolutePath We need to determine the migrations path. For this, the only way to do it from the tests
// is to get the current directory from os.Getwd() and go up until we find the folder which contains "database/migrations"
// directory.
func GetMigrationsAbsolutePath() string {
	workingDirectory, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	lastDir := workingDirectory
	migrationsRelPath := "database/migrations"

	for {
		currentPath := fmt.Sprintf("%s/%s", lastDir, migrationsRelPath)

		fi, err := os.Stat(currentPath)

		if err == nil {
			switch mode := fi.Mode(); {
			case mode.IsDir():
				return currentPath
			}
		}

		newDir := filepath.Dir(lastDir)

		if newDir == lastDir {
			return ""
		}

		lastDir = newDir
	}
}

func NewMockAppOptions() *MockAppOptions {
	return &MockAppOptions{
		Headers: make(map[string][]string),
		Body:    nil,
	}
}
