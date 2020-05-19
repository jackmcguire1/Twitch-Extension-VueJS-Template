package pinpoint

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/pinpoint"
)

type UpdateSegment struct {
	pinpoint.WriteSegmentRequest
}

type Segment struct {
	pinpoint.SegmentResponse
}

func (p *Pinpoint) UpdateUserAttrSegmentDimension(
	segmentId string,
	userAttrKey string,
	items []string,
) (
	err error,
) {
	segment, err := p.GetSegment(segmentId)
	if err != nil {
		return
	}

	userAttributes := segment.Dimensions.UserAttributes

	userAttrDimension, ok := userAttributes[userAttrKey]
	if !ok {
		userAttrDimension = &pinpoint.AttributeDimension{
			AttributeType: aws.String(pinpoint.AttributeTypeInclusive),
			Values:        []*string{},
		}
	}
	userAttrDimension.Values = append(userAttrDimension.Values, aws.StringSlice(items)...)
	userAttributes[userAttrKey] = userAttrDimension

	segment.Dimensions.UserAttributes = userAttributes

	_, err = p.updateSegment(*segment.Id, &pinpoint.WriteSegmentRequest{
		Name:          segment.Name,
		Dimensions:    segment.Dimensions,
		SegmentGroups: segment.SegmentGroups,
	})

	return
}

func (p *Pinpoint) GetSegment(
	segmentId string,
) (
	resp *Segment,
	err error,
) {
	segment, err := svc.GetSegment(&pinpoint.GetSegmentInput{
		ApplicationId: &p.ApplicationId,
		SegmentId:     &segmentId,
	})
	if err != nil {
		return
	}

	raw, err := json.Marshal(segment.SegmentResponse)
	if err != nil {
		return
	}

	resp = &Segment{}
	err = json.Unmarshal(raw, resp)
	if err != nil {
		return
	}

	return
}

func (p *Pinpoint) CreateSegment(
	name string,
) (
	resp *Segment,
	err error,
) {
	segmentOutput, err := svc.CreateSegment(&pinpoint.CreateSegmentInput{
		ApplicationId: &p.ApplicationId,
		WriteSegmentRequest: &pinpoint.WriteSegmentRequest{
			Name: &name,
			Dimensions: &pinpoint.SegmentDimensions{
				UserAttributes: map[string]*pinpoint.AttributeDimension{},
			},
		},
	})
	if err != nil {
		return
	}

	raw, err := json.Marshal(segmentOutput.SegmentResponse)
	if err != nil {
		return
	}

	resp = &Segment{}
	err = json.Unmarshal(raw, resp)
	if err != nil {
		return
	}

	return
}

func (p *Pinpoint) updateSegment(
	segmentId string,
	segmentReq *pinpoint.WriteSegmentRequest,
) (
	resp *pinpoint.SegmentResponse,
	err error,
) {
	updateOutput, err := svc.UpdateSegment(&pinpoint.UpdateSegmentInput{
		SegmentId:           &segmentId,
		ApplicationId:       &p.ApplicationId,
		WriteSegmentRequest: segmentReq,
	})
	if err != nil {
		return
	}
	resp = updateOutput.SegmentResponse

	return
}
