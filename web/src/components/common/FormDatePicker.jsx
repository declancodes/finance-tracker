import React from 'react';
import { Form } from 'react-bootstrap';
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
    <Form.Control as={CustomDatePicker}/>
  );
};