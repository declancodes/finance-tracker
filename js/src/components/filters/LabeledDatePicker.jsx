import React from 'react';
import { Col, Form } from 'react-bootstrap';
import startCase from 'lodash.startcase';
import { FormDatePicker } from '../common/FormDatePicker';

export const LabeledDatePicker = ({ ...props }) => (
  <Form.Group as={Col} xs='auto'>
    <Form.Label>{startCase(props.name)}</Form.Label>
    <FormDatePicker
      {...props}
    />
  </Form.Group>
);
