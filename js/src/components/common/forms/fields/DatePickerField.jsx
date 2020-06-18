import React from 'react';
import { Form } from 'react-bootstrap';
import DatePicker from 'react-datepicker';
import { useField, useFormikContext } from 'formik';

export const DatePickerField = ({ ...props }) => {
  const { setFieldValue } = useFormikContext();
  const [field] = useField(props);
  const CustomDatePicker = (props) => (
    <DatePicker
      {...field}
      {...props}
      wrapperClassName='form-control'
      selected={(field.value && new Date(field.value)) || null}
      onChange={val => setFieldValue(field.name, val)}
    />
  );

  return (
    <Form.Control as={CustomDatePicker}/>
  );
};
