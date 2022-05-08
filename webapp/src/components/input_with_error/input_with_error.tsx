import React, { useState } from "react";
import clsx from "clsx";
// Local modules
import styles from "./input-with-error.module.css";
import Icon from "@components/icon";
import { Icons } from "@components/icon/constants";
import AnimateHeight from "@components/animate-height";

export type InputWithErrorProps = {
  type?: string;
  id?: string;
  label: string;
  className?: string;
  required?: boolean;
  as?: string;
  containerClassName?: string;
  hasError?: boolean;
  errorMessage?: string;
  [key: string]: any;
};

const InputWithError: React.FC<InputWithErrorProps> = React.forwardRef(
  (props, ref) => {
    const {
      type = undefined,
      as = "input",
      id,
      label,
      className = "",
      containerClassName = "",
      required,
      errorMessage,
      hasError,
      ...rest
    } = props;
    const [typeToUse, setTypeToUse] = useState<undefined | string>(type);
    const isPassword = type === "password";
    // Classes
    const typeStyles =
      typeToUse !== "radio" ? styles.textInput : styles.radioInput;
    const errorMessageInput = "border-error text-error font-bold";
    const inputClass = clsx({
      "border-r-0": isPassword,
      [errorMessageInput]: errorMessage || hasError,
    });
    /* =========
     Functions
     ========= */
    const handlePasswordEyeClick = (): void => {
      if (type !== "password") {
        console.error('Input type is not "password"');
        return;
      }
      setTypeToUse((oldType) => (oldType === "text" ? "password" : "text"));
    };
    /* =========
     Render
     ========= */
    return (
      <div className={`form-control ${containerClassName}`}>
        <div className={`${isPassword ? "input-group" : ""} ${typeStyles}`}>
          {React.createElement(as as any, {
            type: typeToUse,
            className: `input input-bordered bg-white w-full transition disabled:bg-gray-100 disabled:border-[rgba(45,55,72,0.2)] ${inputClass} ${className}`,
            id,
            required,
            placeholder: " ",
            ...rest,
            ref,
          })}

          <label
            htmlFor={id}
            className={clsx(
              "label whitespace-pre-line",
              (errorMessage || hasError) && "!text-error"
            )}
          >
            {label}
            {required ? (
              <span className="text-red-500 p-0 bg-transparent">*</span>
            ) : null}
          </label>

          {isPassword ? (
            <span
              className={`password bg-transparent transition border border-l-0 ${
                errorMessage || hasError
                  ? errorMessageInput
                  : "border-base-content"
              }`}
              onClick={handlePasswordEyeClick}
            >
              <button
                className={clsx(
                  "btn btn-plain text-2xl p-0 group",
                  (errorMessage || hasError) && "text-error"
                )}
                type="button"
              >
                <Icon
                  icon={Icons.VisibilityOff}
                  className={clsx(
                    "transition",
                    typeToUse === "password" && "text-grey-100",
                    errorMessage || hasError
                      ? "group-focus:drop-shadow-error"
                      : "group-focus:drop-shadow"
                  )}
                />
              </button>
            </span>
          ) : null}
        </div>

        <AnimateHeight
          className="text-error text-left overflow-hidden translate-y-1"
          isVisible={!!errorMessage}
          visible={{
            marginBottom: 0,
          }}
        >
          {errorMessage}
        </AnimateHeight>
      </div>
    );
  }
);

InputWithError.displayName = "InputWithError";

export default InputWithError;
