package termsreview

import (
	"errors"

	"github.com/redhat-developer/app-services-cli/internal/build"
	termsreviewapi "github.com/redhat-developer/app-services-cli/pkg/api/termsreview"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
)

// CheckTermsAccepted checks whether the current user has accepted the terms and conditions
func CheckTermsAccepted(conn connection.Connection) (accepted bool, redirectURI string, err error) {
	termsParams := termsreviewapi.TermsReviewInput{
		EventCode: build.TermsReviewEventCode,
		SiteCode:  build.TermsReviewSiteCode,
	}

	termsReview, _, err := conn.API().TermsReview().
		SelfTermsReview(&termsParams)

	if err != nil {
		return false, "", err
	}

	if !termsReview.TermsAvailable() && !termsReview.TermsRequired() {
		return true, "", nil
	}

	if redirectURL, ok := termsReview.GetRedirectUrl(); ok {
		return false, redirectURL, nil
	}

	return false, "", errors.New("terms must be signed, but there is no terms URL")
}
