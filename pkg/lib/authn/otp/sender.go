package otp

import (
	"context"
	"net/url"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/infra/mail"
	"github.com/authgear/authgear-server/pkg/lib/infra/sms"
	"github.com/authgear/authgear-server/pkg/lib/infra/task"
	"github.com/authgear/authgear-server/pkg/lib/tasks"
	"github.com/authgear/authgear-server/pkg/util/intl"
	"github.com/authgear/authgear-server/pkg/util/template"
)

type EndpointsProvider interface {
	BaseURL() *url.URL
}

type MessageSender struct {
	Context        context.Context
	ServerConfig   *config.ServerConfig
	Localization   *config.LocalizationConfig
	AppMetadata    config.AppMetadata
	Messaging      *config.MessagingConfig
	TemplateEngine *template.Engine
	Endpoints      EndpointsProvider
	TaskQueue      task.Queue
}

type SendOptions struct {
	OTP         string
	URL         string
	MessageType MessageType
}

func (s *MessageSender) makeData(opts SendOptions) *MessageTemplateContext {
	preferredLanguageTags := intl.GetPreferredLanguageTags(s.Context)
	appName := intl.LocalizeJSONObject(preferredLanguageTags, intl.Fallback(s.Localization.FallbackLanguage), s.AppMetadata, "app_name")

	ctx := &MessageTemplateContext{
		AppName: appName,
		// To be filled by caller
		Email:                "",
		Phone:                "",
		Code:                 opts.OTP,
		URL:                  opts.URL,
		Host:                 s.Endpoints.BaseURL().Host,
		StaticAssetURLPrefix: s.ServerConfig.StaticAsset.URLPrefix,
	}

	return ctx
}

func (s *MessageSender) SendEmail(email string, opts SendOptions, message config.EmailMessageConfig) (err error) {
	data := s.makeData(opts)
	data.Email = email

	var textTemplate, htmlTemplate string
	switch opts.MessageType {
	case MessageTypeVerification:
		textTemplate = TemplateItemTypeVerificationEmailTXT
		htmlTemplate = TemplateItemTypeVerificationEmailHTML
	case MessageTypeSetupPrimaryOOB:
		textTemplate = TemplateItemTypeSetupPrimaryOOBEmailTXT
		htmlTemplate = TemplateItemTypeSetupPrimaryOOBEmailHTML
	case MessageTypeSetupSecondaryOOB:
		textTemplate = TemplateItemTypeSetupSecondaryOOBEmailTXT
		htmlTemplate = TemplateItemTypeSetupSecondaryOOBEmailHTML
	case MessageTypeAuthenticatePrimaryOOB:
		textTemplate = TemplateItemTypeAuthenticatePrimaryOOBEmailTXT
		htmlTemplate = TemplateItemTypeAuthenticatePrimaryOOBEmailHTML
	case MessageTypeAuthenticateSecondaryOOB:
		textTemplate = TemplateItemTypeAuthenticateSecondaryOOBEmailTXT
		htmlTemplate = TemplateItemTypeAuthenticateSecondaryOOBEmailHTML
	default:
		panic("otp: unknown message type: " + opts.MessageType)
	}

	preferredLanguageTags := intl.GetPreferredLanguageTags(s.Context)
	renderCtx := &template.RenderContext{
		PreferredLanguageTags: preferredLanguageTags,
	}

	textBody, err := s.TemplateEngine.Render(renderCtx, textTemplate, data)
	if err != nil {
		return
	}

	htmlBody, err := s.TemplateEngine.Render(renderCtx, htmlTemplate, data)
	if err != nil {
		return
	}

	s.TaskQueue.Enqueue(&tasks.SendMessagesParam{
		EmailMessages: []mail.SendOptions{
			{
				MessageConfig: config.NewEmailMessageConfig(
					s.Messaging.DefaultEmailMessage,
					message,
				),
				Recipient: data.Email,
				TextBody:  textBody,
				HTMLBody:  htmlBody,
			},
		},
	})

	return
}

func (s *MessageSender) SendSMS(phone string, opts SendOptions, message config.SMSMessageConfig) (err error) {
	data := s.makeData(opts)
	data.Phone = phone

	var templateType string
	switch opts.MessageType {
	case MessageTypeVerification:
		templateType = TemplateItemTypeVerificationSMSTXT
	case MessageTypeSetupPrimaryOOB:
		templateType = TemplateItemTypeSetupPrimaryOOBSMSTXT
	case MessageTypeSetupSecondaryOOB:
		templateType = TemplateItemTypeSetupSecondaryOOBSMSTXT
	case MessageTypeAuthenticatePrimaryOOB:
		templateType = TemplateItemTypeAuthenticatePrimaryOOBSMSTXT
	case MessageTypeAuthenticateSecondaryOOB:
		templateType = TemplateItemTypeAuthenticateSecondaryOOBSMSTXT
	default:
		panic("otp: unknown message type: " + opts.MessageType)
	}

	preferredLanguageTags := intl.GetPreferredLanguageTags(s.Context)
	renderCtx := &template.RenderContext{
		PreferredLanguageTags: preferredLanguageTags,
	}

	body, err := s.TemplateEngine.Render(renderCtx, templateType, data)
	if err != nil {
		return
	}

	s.TaskQueue.Enqueue(&tasks.SendMessagesParam{
		SMSMessages: []sms.SendOptions{
			{
				MessageConfig: config.NewSMSMessageConfig(
					s.Messaging.DefaultSMSMessage,
					message,
				),
				To:   data.Phone,
				Body: body,
			},
		},
	})

	return
}
