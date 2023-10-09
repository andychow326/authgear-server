package webapp

import "github.com/google/wire"

var DependencySet = wire.NewSet(
	wire.Struct(new(AuthEntryPointMiddleware), "*"),

	wire.Struct(new(ResponseRenderer), "*"),
	wire.Struct(new(FormPrefiller), "*"),
	wire.Bind(new(Renderer), new(*ResponseRenderer)),

	wire.Struct(new(ControllerDeps), "*"),
	wire.Struct(new(ControllerFactory), "*"),

	wire.Struct(new(AuthflowController), "*"),
	NewAuthflowControllerLogger,

	NewPublisher,
	wire.Struct(new(GlobalSessionServiceFactory), "*"),

	NewPanicMiddlewareLogger,
	wire.Struct(new(PanicMiddleware), "*"),

	wire.Struct(new(AppStaticAssetsHandler), "*"),

	wire.Struct(new(RootHandler), "*"),
	wire.Struct(new(OAuthEntrypointHandler), "*"),
	wire.Struct(new(LoginHandler), "*"),
	wire.Struct(new(SignupHandler), "*"),
	wire.Struct(new(PromoteHandler), "*"),
	wire.Struct(new(SelectAccountHandler), "*"),
	wire.Struct(new(SSOCallbackHandler), "*"),
	wire.Struct(new(EnterLoginIDHandler), "*"),
	wire.Struct(new(EnterPasswordHandler), "*"),
	wire.Struct(new(ConfirmTerminateOtherSessionsHandler), "*"),
	wire.Struct(new(UsePasskeyHandler), "*"),
	wire.Struct(new(CreatePasswordHandler), "*"),
	wire.Struct(new(CreatePasskeyHandler), "*"),
	wire.Struct(new(PromptCreatePasskeyHandler), "*"),
	wire.Struct(new(SetupTOTPHandler), "*"),
	wire.Struct(new(EnterTOTPHandler), "*"),
	wire.Struct(new(SetupOOBOTPHandler), "*"),
	wire.Struct(new(EnterOOBOTPHandler), "*"),
	wire.Struct(new(SetupWhatsappOTPHandler), "*"),
	wire.Struct(new(WhatsappOTPHandler), "*"),
	wire.Struct(new(SetupLoginLinkOTPHandler), "*"),
	wire.Struct(new(LoginLinkOTPHandler), "*"),
	wire.Struct(new(VerifyLoginLinkOTPHandler), "*"),
	wire.Struct(new(EnterRecoveryCodeHandler), "*"),
	wire.Struct(new(SetupRecoveryCodeHandler), "*"),
	wire.Struct(new(VerifyIdentityHandler), "*"),
	wire.Struct(new(VerifyIdentitySuccessHandler), "*"),
	wire.Struct(new(ForgotPasswordHandler), "*"),
	wire.Struct(new(ForgotPasswordSuccessHandler), "*"),
	wire.Struct(new(ResetPasswordHandler), "*"),
	wire.Struct(new(ResetPasswordSuccessHandler), "*"),
	wire.Struct(new(SettingsHandler), "*"),
	wire.Struct(new(TesterHandler), "*"),
	wire.Struct(new(SettingsProfileHandler), "*"),
	wire.Struct(new(SettingsProfileEditHandler), "*"),
	wire.Struct(new(SettingsIdentityHandler), "*"),
	wire.Struct(new(SettingsBiometricHandler), "*"),
	wire.Struct(new(SettingsMFAHandler), "*"),
	wire.Struct(new(SettingsTOTPHandler), "*"),
	wire.Struct(new(SettingsOOBOTPHandler), "*"),
	wire.Struct(new(SettingsRecoveryCodeHandler), "*"),
	wire.Struct(new(SettingsSessionsHandler), "*"),
	wire.Struct(new(ForceChangePasswordHandler), "*"),
	wire.Struct(new(SettingsChangePasswordHandler), "*"),
	wire.Struct(new(ForceChangeSecondaryPasswordHandler), "*"),
	wire.Struct(new(SettingsChangeSecondaryPasswordHandler), "*"),
	wire.Struct(new(SettingsDeleteAccountHandler), "*"),
	wire.Struct(new(SettingsDeleteAccountSuccessHandler), "*"),
	wire.Struct(new(SettingsPasskeyHandler), "*"),
	wire.Struct(new(AccountStatusHandler), "*"),
	wire.Struct(new(LogoutHandler), "*"),
	wire.Struct(new(ReturnHandler), "*"),
	wire.Struct(new(ErrorHandler), "*"),
	wire.Struct(new(NotFoundHandler), "*"),
	wire.Struct(new(WebsocketHandler), "*"),
	wire.Struct(new(WechatAuthHandler), "*"),
	wire.Struct(new(WechatCallbackHandler), "*"),
	wire.Struct(new(PasskeyCreationOptionsHandler), "*"),
	wire.Struct(new(PasskeyRequestOptionsHandler), "*"),
	wire.Struct(new(ConnectWeb3AccountHandler), "*"),
	wire.Struct(new(MissingWeb3WalletHandler), "*"),
	wire.Struct(new(FeatureDisabledHandler), "*"),

	wire.Struct(new(AuthflowLoginHandler), "*"),
	wire.Struct(new(AuthflowSignupHandler), "*"),
	wire.Struct(new(AuthflowEnterPasswordHandler), "*"),
	wire.Struct(new(AuthflowEnterOOBOTPHandler), "*"),
	wire.Struct(new(AuthflowCreatePasswordHandler), "*"),
	wire.Struct(new(AuthflowEnterTOTPHandler), "*"),
	wire.Struct(new(AuthflowSetupTOTPHandler), "*"),
	wire.Struct(new(AuthflowViewRecoveryCodeHandler), "*"),
	wire.Struct(new(AuthflowWhatsappOTPHandler), "*"),
	wire.Struct(new(AuthflowOOBOTPLinkHandler), "*"),

	wire.Struct(new(ResponseWriter), "*"),
)
