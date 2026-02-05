package messagepro

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type messagepro struct {
	apikey             string
	specialPhoneNumber string
	endpoint           string
}

type MessagePro interface {
	MessageSend(phone, text string) (MessageProSendResponse, error)
}

func New(apiKey, specialPhoneNumber, endpoint string) MessagePro {
	return &messagepro{
		apikey:             apiKey,
		specialPhoneNumber: specialPhoneNumber,
		endpoint:           endpoint,
	}
}

func (m *messagepro) MessageSend(phone, text string) (response MessageProSendResponse, err error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	params := url.Values{}
	params.Add("key", m.apikey)
	params.Add("from", m.specialPhoneNumber)
	params.Add("to", phone)
	params.Add("text", text)
	fmt.Println(params.Encode())
	res, err := http.Get(m.endpoint + "/send?" + params.Encode())
	if err != nil {
		return
	}

	responseByte, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		err = errors.New(string(responseByte))
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	var _resp []MessageProSendResponse
	err = json.Unmarshal(responseByte, &_resp)
	if err != nil {
		fmt.Printf("Reading body failed: %s", err)
		return
	}
	response = _resp[0]
	return
}

func (m *messagepro) MessageGetstatus(id string) (response MessageProGetstatusResponse, err error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	params := url.Values{}
	params.Add("key", m.apikey)
	params.Add("id", id)
	res, err := http.Get(m.endpoint + "/getstatus?" + params.Encode())
	if err != nil {
		return
	}
	responseByte, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		err = errors.New(string(responseByte))
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	err = json.Unmarshal(responseByte, &response)
	if err != nil {
		fmt.Printf("Reading body failed: %s", err)
		return
	}
	return
}

func (m *messagepro) MessageFetch(phone, text string) (response string, err error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	params := url.Values{}
	params.Add("from", phone)
	params.Add("to", m.specialPhoneNumber)
	params.Add("text", text)
	res, err := http.Get(m.endpoint + "/fetch?" + params.Encode())
	if err != nil {
		return
	}
	responseByte, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		err = errors.New(string(responseByte))
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	response = string(responseByte)
	return
}

func (m *messagepro) OrderCampaign(body *OrderCampaignRequest) (response string, err error) {

	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}
	req, _ := http.NewRequest(http.MethodPost, m.endpoint+"/order-campaign", requestBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-key", m.apikey)

	res, err := http.DefaultClient.Do(req)
	responseByte, _ := io.ReadAll(res.Body)
	if res.StatusCode != 200 {
		err = errors.New(string(responseByte))
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	response = string(responseByte)
	return
}
