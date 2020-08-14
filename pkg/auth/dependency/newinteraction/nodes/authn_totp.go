package nodes

import (
	"errors"

	"github.com/authgear/authgear-server/pkg/auth/dependency/newinteraction"
	"github.com/authgear/authgear-server/pkg/lib/authn"
	"github.com/authgear/authgear-server/pkg/lib/authn/authenticator"
)

func init() {
	newinteraction.RegisterNode(&NodeAuthenticationTOTP{})
}

type InputAuthenticationTOTP interface {
	GetTOTP() string
}

type EdgeAuthenticationTOTP struct {
	Stage          newinteraction.AuthenticationStage
	Authenticators []*authenticator.Info
}

func (e *EdgeAuthenticationTOTP) AuthenticatorType() authn.AuthenticatorType {
	return authn.AuthenticatorTypeTOTP
}

func (e *EdgeAuthenticationTOTP) HasDefaultTag() bool {
	filtered := filterAuthenticators(e.Authenticators, authenticator.KeepTag(authenticator.TagDefaultAuthenticator))
	return len(filtered) > 0
}

func (e *EdgeAuthenticationTOTP) Instantiate(ctx *newinteraction.Context, graph *newinteraction.Graph, rawInput interface{}) (newinteraction.Node, error) {
	input, ok := rawInput.(InputAuthenticationTOTP)
	if !ok {
		return nil, newinteraction.ErrIncompatibleInput
	}

	inputTOTP := input.GetTOTP()

	var info *authenticator.Info
	for _, a := range e.Authenticators {
		err := ctx.Authenticators.VerifySecret(a, nil, inputTOTP)
		if errors.Is(err, authenticator.ErrInvalidCredentials) {
			continue
		} else if err != nil {
			return nil, err
		} else {
			aa := a
			info = aa
		}
	}

	return &NodeAuthenticationTOTP{Stage: e.Stage, Authenticator: info}, nil
}

type NodeAuthenticationTOTP struct {
	Stage         newinteraction.AuthenticationStage `json:"stage"`
	Authenticator *authenticator.Info                `json:"authenticator"`
}

func (n *NodeAuthenticationTOTP) Prepare(ctx *newinteraction.Context, graph *newinteraction.Graph) error {
	return nil
}

func (n *NodeAuthenticationTOTP) Apply(perform func(eff newinteraction.Effect) error, graph *newinteraction.Graph) error {
	return nil
}

func (n *NodeAuthenticationTOTP) DeriveEdges(graph *newinteraction.Graph) ([]newinteraction.Edge, error) {
	return []newinteraction.Edge{
		&EdgeAuthenticationEnd{
			Stage:                 n.Stage,
			VerifiedAuthenticator: n.Authenticator,
		},
	}, nil
}
