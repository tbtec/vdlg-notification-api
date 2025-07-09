package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

var (
	orderPayload map[string]interface{}
	response     *http.Response
	responseBody []byte
)

func TestFeatures(t *testing.T) {
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"features"},
	}
	status := godog.TestSuite{
		Name:                "order-bdd",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()
	if status != 0 {
		t.Fail()
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have a valid order payload$`, iHaveAValidOrderPayload)
	ctx.Step(`^I send a POST request to "([^"]*)"$`, iSendAPOSTRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, func(code string) error {
		c, _ := strconv.Atoi(code)
		return theResponseCodeShouldBe(c)
	})
	ctx.Step(`^the response should contain the order ID$`, theResponseShouldContainTheOrderID)
}

func iHaveAValidOrderPayload() error {
	orderPayload = map[string]interface{}{
		"customerId": "0196e597-f075-4041-2489-e737f7674898",
	}
	return nil
}

func iSendAPOSTRequestTo(path string) error {

	serverURL := "http://localhost:8083"

	body, err := json.Marshal(orderPayload)
	if err != nil {
		return err
	}

	resp, err := http.Post(serverURL+path, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	response = resp
	responseBody, _ = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return nil
}

func theResponseCodeShouldBe(code int) error {
	if response == nil {
		return fmt.Errorf("response is nil")
	}
	if response.StatusCode != code {
		return fmt.Errorf("expected status %d, got %d", code, response.StatusCode)
	}
	return nil
}

func theResponseShouldContainTheOrderID() error {
	var respMap map[string]interface{}
	err := json.Unmarshal(responseBody, &respMap)
	if err != nil {
		return err
	}
	_, ok := respMap["id"]
	assert.True(nil, ok, "response should contain 'id'")
	return nil
}
