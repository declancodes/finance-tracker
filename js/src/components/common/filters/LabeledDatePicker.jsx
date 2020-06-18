import React from 'react';
import { Form, Col } from 'react-bootstrap';
import DatePicker from 'react-datepicker';
import startCase from 'lodash.startcase';

export const LabeledDatePicker = ({
  name,
  initial,
  onChange
}) => {
  const CustomDatePicker = (props) => (
    <DatePicker
      {...props}
      wrapperClassName='form-control'
      name={name}
      selected={initial}
      onChange={onChange}
    />
  );

  return (
    <Form.Group as={Col} xs='auto'>
      <Form.Label>{startCase(name)}</Form.Label>
      <Form.Control as={CustomDatePicker}/>
    </Form.Group>
  );
};
