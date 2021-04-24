import React from 'react';
import DatePicker from 'react-datepicker';

export const FormDatePicker = ({ ...props }) => {
  const CustomDatePicker = (formProps) => (
    <DatePicker
      {...props}
      {...formProps}
      wrapperClassName='form-control'
    />
  );

  return (
    <CustomDatePicker/>
  );
};