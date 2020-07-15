import React from 'react';
import { useField, useFormikContext } from 'formik';
import { FormDatePicker } from '../../common/FormDatePicker';

export const DatePickerField = ({ ...props }) => {
  const { setFieldValue } = useFormikContext();
  const [field] = useField(props);

  return (
    <FormDatePicker
      {...field}
      selected={(field.value && new Date(field.value)) || null}
      onChange={val => setFieldValue(field.name, val)}
    />
  );
};
