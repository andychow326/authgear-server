import React, { useCallback, useContext, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { PrimaryButton } from "@fluentui/react";
import { FormattedMessage, Context } from "@oursky/react-messageformat";
import WizardScreenLayout from "../../WizardScreenLayout";
import WizardContentLayout from "../../WizardContentLayout";
import FormTextField from "../../FormTextField";
import ShowError from "../../ShowError";
import ShowLoading from "../../ShowLoading";
import { FormErrorMessageBar } from "../../FormErrorMessageBar";
import { FormProvider } from "../../form";
import { useSystemConfig } from "../../context/SystemConfigContext";
import { useCreateAppMutation } from "./mutations/createAppMutation";
import { useAppListQuery } from "./query/appListQuery";
import { ErrorParseRule, makeReasonErrorParseRule } from "../../error/parse";
import { useSimpleForm } from "../../hook/useSimpleForm";

interface FormState {
  appID: string;
}

const DEFAULT_STATE: FormState = {
  appID: "",
};

const FORM_TEXT_FIELD_STYLES = {
  description: {
    display: "block",
    marginTop: "10px",
    fontSize: "12px",
  },
};

const APP_ID_SCHEME = "https://";

const errorRules: ErrorParseRule[] = [
  makeReasonErrorParseRule(
    "DuplicatedAppID",
    "CreateProjectScreen.error.duplicated-app-id"
  ),
  makeReasonErrorParseRule(
    "AppIDReserved",
    "CreateProjectScreen.error.reserved-app-id"
  ),
  makeReasonErrorParseRule(
    "InvalidAppID",
    "CreateProjectScreen.error.invalid-app-id"
  ),
];

interface CreateProjectScreenContentProps {
  numberOfApps: number;
}

function CreateProjectScreenContent(props: CreateProjectScreenContentProps) {
  const { numberOfApps } = props;
  const navigate = useNavigate();
  const { appHostSuffix } = useSystemConfig();
  const { createApp } = useCreateAppMutation();
  const { renderToString } = useContext(Context);

  const submit = useCallback(
    async (state: FormState) => {
      return createApp(state.appID);
    },
    [createApp]
  );

  const form = useSimpleForm({
    stateMode:
      "ConstantInitialStateAndResetCurrentStatetoInitialStateAfterSave",
    defaultState: DEFAULT_STATE,
    submit,
  });

  const {
    updateError,
    save,
    isUpdating,
    state: { appID },
    setState,
  } = form;

  const onChangeAppID = useCallback(
    (_e, newValue) => {
      if (newValue != null) {
        setState((prev) => ({ ...prev, appID: newValue }));
      }
    },
    [setState]
  );

  const onSubmitForm = useCallback(
    (e) => {
      e.preventDefault();
      e.stopPropagation();
      save().catch(() => {});
    },
    [save]
  );

  useEffect(() => {
    if (form.submissionResult) {
      const appID = form.submissionResult;
      navigate(`/project/${encodeURIComponent(appID)}/wizard`);
    }
  }, [form, navigate]);

  return (
    <FormProvider loading={isUpdating} error={updateError}>
      <WizardScreenLayout>
        <FormErrorMessageBar />
        <WizardContentLayout
          title={
            <FormattedMessage
              id="CreateProjectScreen.title"
              values={{
                apps: numberOfApps,
              }}
            />
          }
          backButtonDisabled={true}
          primaryButton={
            <PrimaryButton onClick={onSubmitForm}>
              <FormattedMessage id="CreateProjectScreen.create-project.label" />
            </PrimaryButton>
          }
        >
          <form onSubmit={onSubmitForm}>
            <FormTextField
              styles={FORM_TEXT_FIELD_STYLES}
              parentJSONPointer=""
              fieldName="app_id"
              value={appID}
              onChange={onChangeAppID}
              errorRules={errorRules}
              prefix={APP_ID_SCHEME}
              suffix={appHostSuffix}
              label={renderToString("CreateProjectScreen.app-id.label")}
              description={renderToString(
                "CreateProjectScreen.app-id.description"
              )}
            />
          </form>
        </WizardContentLayout>
      </WizardScreenLayout>
    </FormProvider>
  );
}

const CreateProjectScreen: React.FC = function CreateProjectScreen() {
  const { loading, error, apps, refetch } = useAppListQuery();

  if (loading) {
    return <ShowLoading />;
  }

  if (error != null) {
    return <ShowError error={error} onRetry={refetch} />;
  }

  return <CreateProjectScreenContent numberOfApps={apps?.length ?? 0} />;
};

export default CreateProjectScreen;
