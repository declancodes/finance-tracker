import React from 'react';
import { DatePickerField } from './DatePickerField';
import { LabeledControl } from '../../common/LabeledControl/LabeledControl';
import { SelectField } from './SelectField';
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
    <LabeledControl
      label={displayName}
    >
      {options !== undefined ? (
        <SelectField
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
    </LabeledControl>
  );
};