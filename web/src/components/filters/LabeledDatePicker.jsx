import React from 'react';
import { Col, Form } from 'react-bootstrap';
import { FormDatePicker } from '../common/FormDatePicker';
import { helpers } from '../../common/helpers';

export const LabeledDatePicker = ({ ...props }) => (
  <Form.Group as={Col} xs='auto'>
    <Form.Label>{helpers.titleCase(props.name)}</Form.Label>
    <FormDatePicker
      {...props}
    />
  </Form.Group>
);
