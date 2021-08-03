package termsreview

import (
	authorizationsv1 "github.com/openshift-online/ocm-sdk-go/authorizations/v1"
	"golang.org/x/oauth2"
)

type TermsReviewAPI interface {
	SelfTermsReview(terms *TermsReviewInput) (*TermsReviewResponse, int, error)
}

type TermsReviewInput struct {
	EventCode string
	SiteCode  string
}

type TermsReviewResponse struct {
	*authorizationsv1.TermsReviewResponse
}

type Client struct {
	tc *authorizationsv1.TermsReviewClient
}

func NewClient(tr *oauth2.Transport, apiBaseURL string) *Client {
	client := Client{
		tc: authorizationsv1.NewTermsReviewClient(tr, apiBaseURL),
	}

	return &client
}

func (c *Client) SelfTermsReview(terms *TermsReviewInput) (*TermsReviewResponse, int, error) {
	request, _ := authorizationsv1.NewTermsReviewRequest().
		EventCode(terms.EventCode).
		SiteCode(terms.SiteCode).
		Build()

	termsReviewResp, err := c.tc.Post().Request(request).
		Parameter("event_code", terms.EventCode).
		Parameter("site_code", terms.SiteCode).
		Send()

	httpStatus := termsReviewResp.Status()

	if err != nil {
		return nil, httpStatus, err
	}

	termsReview := TermsReviewResponse{
		termsReviewResp.Response(),
	}

	return &termsReview, httpStatus, nil
}
