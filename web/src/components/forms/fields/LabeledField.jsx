import React from 'react';
import { DatePickerField } from './DatePickerField';
import { LabeledSelectField } from './LabeledSelectField';
import { titleCase } from '../../../common/helpers';

export const LabeledField = ({
  name,
  fieldType,
  options,
  optionDisplay,
  props,
  isMulti
}) => {
  const displayName = titleCase(name);

  return (
    <div className='container-fluid'>
      <label>{displayName}</label>
      {options !== undefined ? (
        <LabeledSelectField
          name={name}
          displayName={displayName}
          options={options}
          optionDisplay={optionDisplay}
          isMulti={isMulti}
        />
      ) : fieldType === 'date' ? (
        <DatePickerField name={name}/>
      ) : (
        <input
          name={name}
          type={fieldType}
          value={props.values[name]}
          onChange={props.handleChange}
        />
      )}
    </div>
  );
};