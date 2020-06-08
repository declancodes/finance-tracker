import React from 'react';
import DatePicker from 'react-datepicker';
import startCase from 'lodash.startcase';

export const LabeledDatePicker = ({
  name,
  initial,
  onChange
}) => (
  <div className='labeled-category-filter'>
    <label htmlFor={name}>{startCase(name)}</label>
    <DatePicker
      name={name}
      selected={initial}
      onChange={onChange}
    />
  </div>
);
