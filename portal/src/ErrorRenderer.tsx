import React from "react";
import { FormattedMessage } from "@oursky/react-messageformat";
import { ParsedAPIError } from "./error/parse";

export interface ErrorRendererProps {
  error?: ParsedAPIError;
  errors?: readonly ParsedAPIError[];
}

const ErrorRenderer: React.FC<ErrorRendererProps> = function ErrorRenderer(
  props: ErrorRendererProps
) {
  const { error, errors } = props;

  let errorArray: ParsedAPIError[] = [];
  if (error != null) {
    errorArray.push(error);
  }
  if (errors != null) {
    errorArray = [...errorArray, ...errors];
  }

  const children = [];
  for (let i = 0; i < errorArray.length; i++) {
    const e = errorArray[i];
    children.push(
      <FormattedMessage key={i} id={e.messageID ?? ""} values={e.arguments} />
    );
  }

  if (children.length === 0) {
    return null;
  }

  return <>{children}</>;
};

export default ErrorRenderer;
