package recaptcha

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	recaptchaenterprise "cloud.google.com/go/recaptchaenterprise/apiv1beta1"
	recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1beta1"
)

var siteKey, projectNumber string

func Init() (err error) {
	projectNumber = os.Getenv("PROJECT_NUMBER")
	siteKey = os.Getenv("SITE_KEY")

	if siteKey == "" || projectNumber == "" {
		return fmt.Errorf("project_number and site_key are mandatory")
	}
	return nil
}

func GetAssessment(token string) (siteAssessment []byte, err error) {

	ctx := context.Background()
	c, err := recaptchaenterprise.NewRecaptchaEnterpriseServiceV1Beta1Client(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	event := recaptchaenterprisepb.Event{
		Token:   token,
		SiteKey: siteKey,
	}
	assessment := recaptchaenterprisepb.Assessment{
		Event: &event,
	}

	req := &recaptchaenterprisepb.CreateAssessmentRequest{
		Parent:     "projects/" + projectNumber,
		Assessment: &assessment,
	}

	resp, err := c.CreateAssessment(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonResp, _ := json.Marshal(resp)

	return jsonResp, nil
}
