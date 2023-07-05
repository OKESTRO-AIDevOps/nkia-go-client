package goclient

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
	"github.com/OKESTRO-AIDevOps/npia-server/src/controller"
	"github.com/OKESTRO-AIDevOps/npia-server/src/modules"

	goya "github.com/goccy/go-yaml"
)

func LoadConfigYaml_Test() {

	var config_yaml map[string]string

	file_byte, err := os.ReadFile("config.yaml")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = goya.Unmarshal(file_byte, &config_yaml)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	base_url := config_yaml["BASE_URL"]

	COMM_URL = base_url + COMM_URL

	COMM_URL_AUTH = base_url + COMM_URL_AUTH

	COMM_URL_MULTIMODE = base_url + COMM_URL_MULTIMODE

	fmt.Println(COMM_URL)
	fmt.Println(COMM_URL_AUTH)
	fmt.Println(COMM_URL_MULTIMODE)
}

func BaseFlow_API_Test() {

	var err error

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Jar: jar,
	}

	err = ClientAuthChallenge(client)

	if err != nil {
		fmt.Println(err)
		return
	}

	var req_body controller.APIMessageRequest
	var resp_body controller.APIMessageResponse

	query_plain := "hello npia"

	query_enc, err := modules.EncryptWithSymmetricKey([]byte(SESSION_SYM_KEY), []byte(query_plain))

	if err != nil {

		fmt.Println(err)
		return
	}

	req_body.Query = hex.EncodeToString(query_enc)

	req_b, err := json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err := client.Post(COMM_URL, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body_bytes))
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_enc := resp_body.QueryResult

	result_enc_b, err := hex.DecodeString(result_enc)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_b, err := modules.DecryptWithSymmetricKey([]byte(SESSION_SYM_KEY), result_enc_b)

	if err != nil {

		fmt.Println(err)
		return
	}

	var api_out apistandard.API_OUTPUT

	err = json.Unmarshal(result_b, &api_out)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(resp_body.ServerMessage)
	fmt.Println(api_out)

}

func BaseFlow_APIThenMultiMode_Test() {

	var err error

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Jar: jar,
	}

	err = ClientAuthChallenge(client)

	if err != nil {
		fmt.Println(err)
		return
	}

	var req_body controller.APIMessageRequest
	var resp_body controller.APIMessageResponse

	query_plain := "hello npia"

	query_enc, err := modules.EncryptWithSymmetricKey([]byte(SESSION_SYM_KEY), []byte(query_plain))

	if err != nil {

		fmt.Println(err)
		return
	}

	req_body.Query = hex.EncodeToString(query_enc)

	req_b, err := json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err := client.Post(COMM_URL, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body_bytes))
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_enc := resp_body.QueryResult

	result_enc_b, err := hex.DecodeString(result_enc)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_b, err := modules.DecryptWithSymmetricKey([]byte(SESSION_SYM_KEY), result_enc_b)

	if err != nil {

		fmt.Println(err)
		return
	}

	var api_out apistandard.API_OUTPUT

	err = json.Unmarshal(result_b, &api_out)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("----------API----------")
	fmt.Println(resp_body.ServerMessage)
	fmt.Println(api_out)
	fmt.Println("-----------------------")
	fmt.Println(" ")

	query_plain = "INIT:"

	query_enc, err = modules.EncryptWithSymmetricKey([]byte(SESSION_SYM_KEY), []byte(query_plain))

	if err != nil {

		fmt.Println(err)
		return
	}

	req_body.Query = hex.EncodeToString(query_enc)

	req_b, err = json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err = client.Post(COMM_URL_MULTIMODE, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body_bytes))
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(resp_body.ServerMessage)

	result_enc = resp_body.QueryResult

	result_enc_b, err = hex.DecodeString(result_enc)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_b, err = modules.DecryptWithSymmetricKey([]byte(SESSION_SYM_KEY), result_enc_b)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("----------MULTIMODE----------")
	fmt.Println(string(result_b))

	query_plain = "SWITCH:kind-kindcluster2"

	query_enc, err = modules.EncryptWithSymmetricKey([]byte(SESSION_SYM_KEY), []byte(query_plain))

	if err != nil {

		fmt.Println(err)
		return
	}

	req_body.Query = hex.EncodeToString(query_enc)

	req_b, err = json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err = client.Post(COMM_URL_MULTIMODE, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body_bytes))
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(resp_body.ServerMessage)

	result_enc = resp_body.QueryResult

	result_enc_b, err = hex.DecodeString(result_enc)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_b, err = modules.DecryptWithSymmetricKey([]byte(SESSION_SYM_KEY), result_enc_b)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(string(result_b))

}
