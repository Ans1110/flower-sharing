import { z } from "zod";

const customErrorMap: z.ZodErrorMap = (issue, ctx) => {
  switch (issue.code) {
    case z.ZodIssueCode.invalid_type:
      if (issue.received === "undefined" || issue.received === "null") {
        return { message: "This field is required" };
      }
      if (issue.expected === "string") {
        return { message: "Please enter text" };
      }
      if (issue.expected === "number") {
        return { message: "Please enter a number" };
      }
      return { message: `Invalid value type` };

    case z.ZodIssueCode.too_small:
      if (issue.type === "string") {
        return { message: `Minimum ${issue.minimum} characters required` };
      }
      if (issue.type === "number") {
        return {
          message: `Number must be greater than or equal to ${issue.minimum}`,
        };
      }
      return { message: `Value is too small` };

    case z.ZodIssueCode.too_big:
      if (issue.type === "string") {
        return { message: `Maximum ${issue.maximum} characters allowed` };
      }
      if (issue.type === "number") {
        return {
          message: `Number must be less than or equal to ${issue.maximum}`,
        };
      }
      return { message: `Value is too large` };

    case z.ZodIssueCode.invalid_string:
      if (issue.validation === "email") {
        return { message: "Please enter a valid email address" };
      }
      if (issue.validation === "url") {
        return { message: "Please enter a valid URL" };
      }
      return { message: "Invalid text format" };

    case z.ZodIssueCode.invalid_date:
      return { message: "Please enter a valid date" };

    case z.ZodIssueCode.custom:
      return { message: issue.message || "Invalid value" };

    default:
      return { message: ctx.defaultError };
  }
};

const oauthErrorMessages: Record<string, string> = {
  invalid_state: "Invalid state token. Please try again.",
  no_code: "No authorization code received. Please try again.",
  token_exchange_failed: "Failed to exchange token. Please try again.",
  user_info_failed: "Failed to get user information. Please try again.",
  parse_failed: "Failed to parse user information. Please try again.",
  user_creation_failed: "Failed to create user account. Please try again.",
  token_generation_failed:
    "Failed to generate authentication token. Please try again.",
  token_save_failed: "Failed to save authentication token. Please try again.",
  no_email:
    "No email found in your account. Please make sure your email is public or try another method.",
};

export { customErrorMap, oauthErrorMessages };
