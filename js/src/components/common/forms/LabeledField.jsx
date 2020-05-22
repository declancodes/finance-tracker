import React from 'react';
import { Field } from 'formik';
import { Options } from '../Options';

export const LabeledField = ({ name, fieldType, options }) => {
  const displayName = name.charAt(0).toUpperCase() + name.slice(1);
  return (
    <div>
      <label htmlFor={name}>{displayName}</label>
      {options === undefined ? (
        <Field name={name} type={fieldType} />
      ) : (
        <Field name={name} as='select'>
          <Options
            entityName={displayName}
            options={options}
            optionValue='uuid'
          />
        </Field>
      )}
    </div>
  );
};