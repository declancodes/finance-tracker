import React from 'react';
import { Field } from 'formik';
import { DatePickerField } from './DatePickerField';
import { Options } from '../../Options';
import startCase from 'lodash.startcase';

export const LabeledField = ({
  name,
  fieldType,
  options,
  optionDisplay
}) => {
  const displayName = startCase(name);
  return (
    <div className='labeled-field'>
      <label htmlFor={name}>{displayName}</label>
      {options !== undefined ? (
        <Field name={name} as='select'>
          <Options
            defaultOptionText={`Select ${displayName}`}
            options={options}
            optionValue='uuid'
            optionDisplay={optionDisplay}
          />
        </Field>
      ) : fieldType === 'date' ? (
        <DatePickerField name={name}/>
      ) : (
        <Field name={name} type={fieldType}/>
      )}
    </div>
  );
};