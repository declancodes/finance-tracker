import React from 'react';
import { Field } from 'formik';
import { Options } from '../Options';
import startCase from 'lodash.startcase';

export const LabeledField = ({
  name,
  fieldType,
  options,
  optionDisplay
}) => {
  const displayName = startCase(name);
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
            optionDisplay={optionDisplay}
          />
        </Field>
      )}
    </div>
  );
};