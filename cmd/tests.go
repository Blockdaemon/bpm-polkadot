package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	baseUrl       = "http://127.0.0.1:9933"
	sysChain      = "Alexander"
	sysName       = "parity-polkadot"
	tokenDecimals = 15
	tokenSymbol   = "DOT"
	sysVersion    = "0.4.4"
)

type testRunner struct {
	failed    int
	succeeded int
}

func (t *testRunner) test(testFunc func() error) {
	if err := testFunc(); err != nil {
		t.failed++
	} else {
		t.succeeded++
	}
}

func runAllTests() error {
	runner := testRunner{}

	runner.test(chainGetBlock)
	runner.test(chainGetBlockHash)
	runner.test(chainGetFinalizedHead)
	runner.test(chainGetHeader)
	runner.test(systemChain)
	runner.test(systemHealth)
	runner.test(systemName)
	runner.test(systemNetworkState)
	runner.test(systemPeers)
	runner.test(systemProperties)
	runner.test(systemVersion)

	if runner.failed > 0 {
		return fmt.Errorf("%d out of %d tests failed", runner.failed, runner.failed+runner.succeeded)
	}

	return nil
}

func chainGetBlock() error {
	// hash parameter can be set
	// only checking for no error 200 right now
	method, statusCode, expectedID, data := polkadotPost("chain_getBlock", "")

	errorValue, errorKey := data["error"]
	result, _ := data["result"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage := fmt.Sprintf("PASSED: %s %d", method, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure %v", method, result)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func chainGetBlockHash() error {
	// blockNumber parameter can be set
	// only checking for no error 200 right now
	method, statusCode, expectedID, data := polkadotPost("chain_getBlockHash", "")

	errorValue, errorKey := data["error"]
	result, _ := data["result"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage := fmt.Sprintf("PASSED: %s %d", method, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure %v", method, result)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func chainGetFinalizedHead() error {
	// only checking for no error 200 right now
	method, statusCode, expectedID, data := polkadotPost("chain_getFinalizedHead", "")

	errorValue, errorKey := data["error"]
	result, _ := data["result"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage := fmt.Sprintf("PASSED: %s %d", method, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure %v", method, result)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func chainGetHeader() error {
	// hash parameter can be set
	// only checking for no error 200 right now
	method, statusCode, expectedID, data := polkadotPost("chain_getHeader", "")

	errorValue, errorKey := data["error"]
	result, _ := data["result"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if statusCode == 200 {
		logMessage := fmt.Sprintf("PASSED: %s %d", method, statusCode)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure %v", method, result)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func systemChain() error {
	name := sysChain
	method, statusCode, expectedID, data := polkadotPost("system_chain", "")

	errorValue, errorKey := data["error"]
	result, _ := data["result"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if result == name {
		logMessage := fmt.Sprintf("PASSED: %s %s", method, name)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure %v", method, result)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func systemHealth() error {
	method, statusCode, expectedID, data := polkadotPost("system_health", "")

	errorValue, errorKey := data["error"]
	resultMap, _ := data["result"].(map[string]interface{})
	resultIsSyncing, _ := resultMap["isSyncing"].(bool)
	resultPeers, _ := resultMap["peers"].(float64)
	resultShouldHavePeers, _ := resultMap["shouldHavePeers"].(bool)
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if resultIsSyncing && resultPeers > 0 && resultShouldHavePeers {
		//might need to change this if we ever expect not to have peers
		logMessage := fmt.Sprintf("PASSED: %s Syncing: %v Peers: %d Should Have Peers: %v", method, resultIsSyncing, int(resultPeers), resultShouldHavePeers)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure %v", method, resultMap)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func systemName() error {
	name := sysName
	method, statusCode, expectedID, data := polkadotPost("system_name", "")

	errorValue, errorKey := data["error"]
	result, _ := data["result"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if result == name {
		logMessage := fmt.Sprintf("PASSED: %s %s", method, result)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure %v", method, result)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func systemNetworkState() error {
	method, statusCode, expectedID, data := polkadotPost("system_networkState", "")

	errorValue, errorKey := data["error"]
	resultMap, _ := data["result"].(map[string]interface{})
	resultAverageDownloadPerSec, _ := resultMap["averageDownloadPerSec"].(float64)
	resultAverageUploadPerSec, _ := resultMap["averageUploadPerSec"].(float64)
	//resultConnectedPeers, _ := resultMap["connectedPeers"]
	//resultExternalAddresses, _ := resultMap["externalAddresses"].([]string)
	//resultListenedAddresses, _ := resultMap["listenedAddresses"].([]string)
	//resultNotConnectedPeers, _ := resultMap["notConnectedPeers"].(map[string]interface{})
	//resultPeerID, _ := resultMap["peerId"]
	//resultpeerset, _ := resultMap["averageUploadPerSec"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if resultAverageDownloadPerSec > 0 && resultAverageUploadPerSec > 0 {
		logMessage := fmt.Sprintf("PASSED: %s average DL/UL %d/%d", method, int(resultAverageDownloadPerSec), int(resultAverageUploadPerSec))
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure %v", method, resultMap)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func systemPeers() error {
	method, statusCode, expectedID, data := polkadotPost("system_peers", "")

	errorValue, errorKey := data["error"]
	result, _ := data["result"]
	//firstPeerMap := result.([]interface{})[0].(map[string]interface{})
	//bestHash, _ := firstPeerMap["bestHash"]
	//bestNumber, _ := firstPeerMap["bestNumber"]
	//peerID, _ := firstPeerMap["peerId"]
	//protocolVersion, _ := firstPeerMap["protocolVersion"]
	//roles, _ := firstPeerMap["roles"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if len(result.([]interface{})) > 0 {
		logMessage := fmt.Sprintf("PASSED: %s %d peer(s)", method, len(result.([]interface{})))
		//logMessage := fmt.Sprintf("PASSED: %s returns peers %v", method, result)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure %v", method, result)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func systemProperties() error {
	method, statusCode, expectedID, data := polkadotPost("system_properties", "")

	errorValue, errorKey := data["error"]
	resultMap, _ := data["result"].(map[string]interface{})
	resultTokenDecimals, _ := resultMap["tokenDecimals"]
	resultTokenSymbol, _ := resultMap["tokenSymbol"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if resultTokenDecimals != float64(tokenDecimals) {
		logMessage := fmt.Sprintf("FAILED: %s result decimals: %v but expected: %d", method, resultTokenDecimals, tokenDecimals)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if resultTokenSymbol != tokenSymbol {
		logMessage := fmt.Sprintf("FAILED: %s result symbol: %v but expected: %s", method, resultTokenSymbol, tokenSymbol)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if resultTokenDecimals == float64(tokenDecimals) && resultTokenSymbol == tokenSymbol {
		logMessage := fmt.Sprintf("PASSED: %s token decimals: %v token symbol: %v", method, resultTokenDecimals, resultTokenSymbol)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure", method)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func systemVersion() error {
	expectedVersion := sysVersion
	method, statusCode, expectedID, data := polkadotPost("system_version", "")

	errorValue, errorKey := data["error"]
	resultVersion, _ := data["result"]
	messageIDValue, _ := data["id"]

	if statusCode != 200 {
		logMessage := fmt.Sprintf("FAILED: %s status code: %d", method, statusCode)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if errorKey {
		logMessage := fmt.Sprintf("FAILED: %s error in body: %v", method, errorValue)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if messageIDValue != float64(expectedID) {
		logMessage := fmt.Sprintf("FAILED: %s message id: %f but expected %d", method, messageIDValue, expectedID)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if resultVersion != expectedVersion {
		logMessage := fmt.Sprintf("FAILED: %s version: %s but expected %v", method, expectedVersion, resultVersion)
		log.Print(logMessage)
		return errors.New(logMessage)
	} else if resultVersion == expectedVersion {
		logMessage := fmt.Sprintf("PASSED: %s %v", method, resultVersion)
		log.Print(logMessage)
		return nil
	} else {
		logMessage := fmt.Sprintf("FAILED: %s unknown failure", method)
		log.Print(logMessage)
		return errors.New(logMessage)
	}
}

func polkadotPost(method string, params string) (string, int, int, map[string]interface{}) {
	rand.Seed(time.Now().UnixNano())
	var messageID = rand.Int()
	requestBody, err := json.Marshal(map[string]interface{}{
		"method":  method,
		"id":      messageID,
		"jsonrpc": "2.0",
	})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(baseUrl+params, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}

	return method, resp.StatusCode, messageID, data
}
