import React from "react";
import { Field } from "formik";

export const LabeledField = ({ name, fieldType, options }) => {
  const displayName = name.charAt(0).toUpperCase() + name.slice(1);
  return (
    <div>
      <label htmlFor={name}>{displayName}</label>
      {typeof options === "undefined" ? (
        <Field name={name} type={fieldType} />
      ) : (
        <Field name={name} as="select">
          <option defaultValue="">Select {displayName}</option>
          {options.length > 0 && (
            options.map(option => (
              <option key={option.uuid} value={option.uuid}>
                {option.name}
              </option>
            ))
          )}
        </Field>
      )}
    </div>
  );
};