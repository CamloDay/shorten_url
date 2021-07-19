package step_definitions

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/rdumont/assistdog"
	"net/http"
	"net/http/httptest"
	"reflect"
	"shorten_url/api"
	"shorten_url/api/handlers"
	"strings"
)

type Feature struct {
	service *api.Service
	resp    *httptest.ResponseRecorder
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	f := &Feature{}

	ctx.BeforeScenario(f.beforeScenario)

	ctx.Step(`^this url exists "([^"]*)" for short "([^"]*)"$`, f.thisUrlExistsForShort)
	ctx.Step(`^I send a "([^"]*)" to "([^"]*)"$`, f.iSendATo)
	ctx.Step(`^I send a "([^"]*)" to "([^"]*)" with the following:$`, f.iSendAToWithTheFollowing)
	ctx.Step(`^I should get a (\d+) HTTP response code`, f.iShouldGetAHTTPResponseCode)
	ctx.Step(`^I should get a (\d+) HTTP response containing the following "([^"]*)":$`, f.iShouldGetAHTTPResponseContainingTheFollowing)
}

func (f *Feature) beforeScenario(*godog.Scenario) {
	f.service = api.NewService()

	f.resp = httptest.NewRecorder()
	api.CreateRoutes(f.service, "localhost:8080")
}

func (f *Feature) thisUrlExistsForShort(original, short string) error {
	f.service.Urls[short] = &original
	return nil
}

func (f *Feature) iSendATo(method, path string) error {
	return f.sendAPIRequest(method, path, "")
}

func (f *Feature) iSendAToWithTheFollowing(method, path string, body *messages.PickleStepArgument_PickleDocString) error {
	return f.sendAPIRequest(method, path, body.Content)
}

func (f *Feature) sendAPIRequest(method, path string, body string) error {
	var req *http.Request
	if body == "" {
		if path == "http://localhost:8080/def" { //map testing constant to random short
			keys := reflect.ValueOf(f.service.Urls).MapKeys()
			if len(keys) > 0 {
				path = "http://localhost:80080/" + keys[0].String()
			}
		}
		req = httptest.NewRequest(method, path, nil)
	} else {
		req = httptest.NewRequest(method, path, strings.NewReader(body))
	}
	req.Header.Set("Content-Type", "application/json")

	f.resp = httptest.NewRecorder()
	f.service.Router.ServeHTTP(f.resp, req)
	return nil
}

func (f *Feature) iShouldGetAHTTPResponseCode(code int) error {
	if f.resp.Code != code {
		return fmt.Errorf("expected a HTTP status code of %d, got %d - %s\n", code, f.resp.Code, f.resp.Body.String())
	}
	return nil
}

func (f *Feature) iShouldGetAHTTPResponseContainingTheFollowing(code int, typ string, tbl *godog.Table) error {
	if f.resp.Code != code {
		return fmt.Errorf("expected a HTTP status code of %d, got %d - %s\n", code, f.resp.Code, f.resp.Body.String())
	}
	var act interface{}
	//var err error
	switch typ {
	case "StandardResponse":
		act = new(handlers.StandardResponse)
	case "ShortResponse":
		act = new(handlers.ShortResponse)
	default:
		fmt.Printf("iShouldGetAHTTPResponseContainingTheFollowing %s not implemented\n", typ)
		return godog.ErrPending
	}

	assist := newAssist()
	exp, err := assist.CreateInstance(act, tbl)
	if err != nil {
		fmt.Printf("iShouldGetAHTTPResponseContainingTheFollowing err: %v\n", err)
		return err
	}

	if err = json.Unmarshal(f.resp.Body.Bytes(), act); err != nil {
		fmt.Printf("iShouldGetAHTTPResponseContainingTheFollowing - invalid Body: %s\n", f.resp.Body.String())
		return err
	}

	switch typ {
	case "ShortResponse":
		sr := act.(*handlers.ShortResponse)
		keys := reflect.ValueOf(f.service.Urls).MapKeys()
		if len(keys) < 1 {
			return fmt.Errorf("iShouldGetAHTTPResponseContainingTheFollowing - no short urls in map\n")
		}
		sr.Short = "http://localhost:8080/def" //map random short to testing constant
		act = sr
	default:
	}

	if !reflect.DeepEqual(exp, act) {
		return fmt.Errorf("\nexp: %+v\nact: %+v\n", exp, act)
	}
	return nil
}

func newAssist() *assistdog.Assist {
	assist := assistdog.NewDefault()
	return assist
}
