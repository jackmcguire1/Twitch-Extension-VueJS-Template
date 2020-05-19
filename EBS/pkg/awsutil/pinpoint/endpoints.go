package pinpoint

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/pinpoint"
)

type EndpointReq pinpoint.EndpointRequest
type EndpointBatchItem pinpoint.EndpointBatchItem
type EndpointsResp pinpoint.EndpointsResponse
type EndpointUser pinpoint.EndpointUser

type ResponseMessage struct {
	Message   string
	RequestId string
}

func (p *Pinpoint) GetEndpointByEndpointId(
	endpointId string,
) (
	resp *Endpoint,
	err error,
) {
	output, err := svc.GetEndpoint(&pinpoint.GetEndpointInput{
		ApplicationId: &p.ApplicationId,
		EndpointId:    &endpointId,
	})
	if err != nil {
		return
	}

	raw, err := json.Marshal(output.EndpointResponse)
	if err != nil {
		return
	}

	resp = &Endpoint{}
	err = json.Unmarshal(raw, resp)
	if err != nil {
		return
	}

	return
}

func (p *Pinpoint) GetEndpointsByUserId(
	userId string,
) (
	resp []*Endpoint,
	err error,
) {
	output, err := svc.GetUserEndpoints(&pinpoint.GetUserEndpointsInput{
		ApplicationId: &p.ApplicationId,
		UserId:        &userId,
	})
	if err != nil {
		return
	}

	raw, err := json.Marshal(output.EndpointsResponse.Item)
	if err != nil {
		return
	}

	resp = []*Endpoint{}
	err = json.Unmarshal(raw, &resp)
	if err != nil {
		return
	}

	return
}

func (p *Pinpoint) UpdateEndpointsUserAttributes(
	userAttrKey string,
	items []string,
	endpointIds []string,
) (
	resp *ResponseMessage,
	err error,
) {

	endpointsToUpdate := []*pinpoint.EndpointBatchItem{}
	for _, endpointId := range endpointIds {
		item := &pinpoint.EndpointBatchItem{
			Id: aws.String(endpointId),
			User: &pinpoint.EndpointUser{
				UserAttributes: map[string][]*string{userAttrKey: aws.StringSlice(items)},
			},
		}

		endpointsToUpdate = append(endpointsToUpdate, item)
	}

	raw, err := json.Marshal(endpointsToUpdate)
	if err != nil {
		return
	}

	batchItems := []*EndpointBatchItem{}
	err = json.Unmarshal(raw, &batchItems)
	if err != nil {
		return
	}

	resp, err = p.BatchUpdateEndpoints(batchItems)
	if err != nil {
		return
	}

	return
}

func (p *Pinpoint) UpdateEndpoint(
	endpointId string,
	req *EndpointReq,
) (
	resp *ResponseMessage,
	err error,
) {
	raw, err := json.Marshal(req)
	if err != nil {
		return
	}

	endpointReq := &pinpoint.EndpointRequest{}
	err = json.Unmarshal(raw, endpointReq)
	if err != nil {
		return
	}

	output, err := svc.UpdateEndpoint(&pinpoint.UpdateEndpointInput{
		ApplicationId:   &p.ApplicationId,
		EndpointId:      &endpointId,
		EndpointRequest: endpointReq,
	})
	if err != nil {
		return
	}

	resp = &ResponseMessage{
		Message:   *output.MessageBody.Message,
		RequestId: *output.MessageBody.Message,
	}

	return
}

func (p *Pinpoint) BatchUpdateEndpoints(
	endpoints []*EndpointBatchItem,
) (
	resp *ResponseMessage,
	err error,
) {

	raw, err := json.Marshal(endpoints)
	if err != nil {
		return
	}

	batchItems := []*pinpoint.EndpointBatchItem{}
	err = json.Unmarshal(raw, &batchItems)
	if err != nil {
		return
	}

	output, err := svc.UpdateEndpointsBatch(&pinpoint.UpdateEndpointsBatchInput{
		ApplicationId: &p.ApplicationId,
		EndpointBatchRequest: &pinpoint.EndpointBatchRequest{
			Item: batchItems,
		},
	})
	if err != nil {
		return
	}

	resp = &ResponseMessage{
		Message:   *output.MessageBody.Message,
		RequestId: *output.MessageBody.Message,
	}

	return
}

type Endpoint struct {
	Id             string              `json:"id"`
	Address        string              `json:"address,omitempty"`
	Metrics        map[string]float64  `json:"metrics,omitempty"`
	ChannelType    string              `json:"channelType"`
	EffectiveDate  string              `json:"effectiveDate"`
	Demographic    Demographic         `json:"demographic"`
	OptOut         string              `json:"optOut"`
	EndpointStatus string              `json:"endpointStatus"`
	Attributes     map[string][]string `json:"attributes"`
	Location       *Location           `json:"location,omitempty"`
	User           User                `json:"user"`
}

type Demographic struct {
	AppVersion      string `json:"appVersion"`
	Locale          string `json:"locale"`
	Make            string `json:"make"`
	Model           string `json:"model"`
	ModelVersion    string `json:"modelVersion"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
	Timezone        string `json:"timezone"`
}

type Location struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	PostalCode string  `json:"postalCode"`
	City       string  `json:"city"`
	Region     string  `json:"region"`
	Country    string  `json:"country"`
}

type User struct {
	UserId         string              `json:"userId"`
	UserAttributes map[string][]string `json:"userAttributes"`
}
