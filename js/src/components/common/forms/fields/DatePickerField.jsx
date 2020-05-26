import React from 'react';
import DatePicker from 'react-datepicker';
import { useField, useFormikContext } from 'formik';
import '../../../../../node_modules/react-datepicker/dist/react-datepicker.css';

export const DatePickerField = ({ ...props }) => {
  const { setFieldValue } = useFormikContext();
  const [field] = useField(props);
  return (
    <DatePicker
      {...field}
      {...props}
      selected={(field.value && new Date(field.value)) || null}
      onChange={val => setFieldValue(field.name, val)}
    />
  );
};
