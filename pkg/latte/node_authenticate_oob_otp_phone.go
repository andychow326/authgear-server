package latte

import (
	"context"
	"errors"
	"time"

	"github.com/authgear/authgear-server/pkg/api"
	"github.com/authgear/authgear-server/pkg/api/model"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator"
	"github.com/authgear/authgear-server/pkg/lib/workflow"
	"github.com/authgear/authgear-server/pkg/util/phone"
)

func init() {
	workflow.RegisterNode(&NodeAuthenticateOOBOTPPhone{})
}

type NodeAuthenticateOOBOTPPhone struct {
	Authenticator *authenticator.Info `json:"authenticator,omitempty"`
}

func (n *NodeAuthenticateOOBOTPPhone) Kind() string {
	return "latte.NodeAuthenticateOOBOTPPhone"
}

func (n *NodeAuthenticateOOBOTPPhone) GetEffects(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) (effs []workflow.Effect, err error) {
	return nil, nil
}

func (n *NodeAuthenticateOOBOTPPhone) CanReactTo(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) ([]workflow.Input, error) {
	return []workflow.Input{
		&InputTakeOOBOTPCode{},
		&InputResendOOBOTPCode{},
	}, nil
}

func (n *NodeAuthenticateOOBOTPPhone) ReactTo(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow, input workflow.Input) (*workflow.Node, error) {
	var inputTakeOOBOTPCode inputTakeOOBOTPCode
	var inputResendOOBOTPCode inputResendOOBOTPCode
	switch {
	case workflow.AsInput(input, &inputResendOOBOTPCode):
		info := n.Authenticator
		_, err := (&SendOOBCode{
			WorkflowID:        workflow.GetWorkflowID(ctx),
			Deps:              deps,
			Stage:             authenticatorKindToStage(info.Kind),
			IsAuthenticating:  true,
			AuthenticatorInfo: info,
		}).Do()
		if err != nil {
			return nil, err
		}
		return nil, workflow.ErrSameNode
	case workflow.AsInput(input, &inputTakeOOBOTPCode):
		info := n.Authenticator
		_, err := deps.Authenticators.VerifyWithSpec(info, &authenticator.Spec{
			OOBOTP: &authenticator.OOBOTPSpec{
				Code: inputTakeOOBOTPCode.GetCode(),
			},
		})
		if errors.Is(err, authenticator.ErrInvalidCredentials) {
			if err := DispatchAuthenticationFailedEvent(deps.Events, info); err != nil {
				return nil, err
			}
			return nil, api.ErrInvalidCredentials
		} else if err != nil {
			return nil, err
		}
		return workflow.NewNodeSimple(&NodeVerifiedAuthenticator{
			Authenticator: info,
		}), nil
	}
	return nil, workflow.ErrIncompatibleInput
}

func (n *NodeAuthenticateOOBOTPPhone) OutputData(ctx context.Context, deps *workflow.Dependencies, w *workflow.Workflow) (interface{}, error) {
	bucket := deps.AntiSpamOTPCodeBucket.MakeBucket(model.AuthenticatorOOBChannelSMS, n.Authenticator.OOBOTP.Phone)
	_, resetDuration, err := deps.RateLimiter.CheckToken(bucket)
	if err != nil {
		return nil, err
	}
	now := deps.Clock.NowUTC()
	canResendAt := now.Add(resetDuration)

	type NodeAuthenticateOOBOTPPhoneOutput struct {
		MaskedPhoneNumber string    `json:"masked_phone_number"`
		CanResendAt       time.Time `json:"can_resend_at"`
	}

	return NodeAuthenticateOOBOTPPhoneOutput{
		MaskedPhoneNumber: phone.Mask(n.Authenticator.OOBOTP.Phone),
		CanResendAt:       canResendAt,
	}, nil
}
