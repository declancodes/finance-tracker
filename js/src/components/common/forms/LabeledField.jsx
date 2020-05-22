import React from 'react';
import { Field } from 'formik';
import { Options } from '../Options';
import capitalize from 'lodash.capitalize';

export const LabeledField = ({ name, fieldType, options }) => {
  const displayName = capitalize(name);
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