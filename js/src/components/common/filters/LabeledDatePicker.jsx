import React from 'react';
import DatePicker from 'react-datepicker';
import capitalize from 'lodash.capitalize';

export const LabeledDatePicker = ({
  name,
  initial,
  onChange
}) => (
  <>
    <label htmlFor={name}>{capitalize(name)}</label>
    <DatePicker
      name={name}
      selected={initial}
      onChange={onChange}
    />
  </>
);
