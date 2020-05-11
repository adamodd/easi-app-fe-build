package email

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/cmsgov/easi-app/pkg/appconfig"
	"github.com/cmsgov/easi-app/pkg/testhelpers"
)

type SESTestSuite struct {
	suite.Suite
	logger *zap.Logger
	client Client
}

func TestSESTestSuite(t *testing.T) {
	logger := zap.NewNop()
	config := testhelpers.NewConfig()

	emailConfig := Config{
		GRTEmail:          config.GetString(appconfig.GRTEmailKey),
		URLHost:           config.GetString(appconfig.ApplicationHostKey),
		URLScheme:         config.GetString(appconfig.ApplicationProtocolKey),
		TemplateDirectory: config.GetString(appconfig.EmailTemplateDirectoryKey),
	}

	sesConfig := SESConfig{
		SourceARN: config.GetString(appconfig.AWSSESSourceARNKey),
		Source:    config.GetString(appconfig.AWSSESSourceKey),
	}

	sesSession := session.Must(session.NewSession())
	sesClient := ses.New(sesSession)
	sesSender := SESSender{
		sesClient,
		sesConfig,
	}

	client, err := NewClient(emailConfig, sesSender)
	if err != nil {
		t.Errorf("Failed to create email client %v", err)
	}
	sesTestSuite := &SESTestSuite{
		Suite:  suite.Suite{},
		logger: logger,
		client: client,
	}

	suite.Run(t, sesTestSuite)
}
