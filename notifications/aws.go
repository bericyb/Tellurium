package notifications

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sns"
)

func getSNSClient() *sns.SNS {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sns.New(sess)
	return svc
}

func getSESClient() *ses.SES {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ses.New(sess)
	return svc
}

func SendAWSTextAlert(phoneNumber, command string, time time.Duration) {

	params := &sns.PublishInput{

		Message:     aws.String("Execution of \"" + command + "\" finished in: " + time.String()),
		PhoneNumber: aws.String("+1" + phoneNumber),
	}

	_, err := getSNSClient().Publish(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Text sent!")
}
func SendAWSEmailAlert(emailAddress, command string, time time.Duration) {

	params := &ses.SendEmailInput{
		Source: aws.String("bericb@gmail.com"),
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(emailAddress),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String("Execution of \"" + command + "\" finished in: " + time.String()),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Tellurium Alert"),
			},
		},
	}

	_, err := getSESClient().SendEmail(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Email sent!")
}
