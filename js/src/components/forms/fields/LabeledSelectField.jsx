import React from 'react';
import { Form } from 'react-bootstrap';
import { useField, useFormikContext } from 'formik';
import { helpers } from '../../../common/helpers';
import { Options } from '../../common/Options';

export const LabeledSelectField = ({ ...props }) => {
  const { setFieldValue } = useFormikContext();
  const [field] = useField(props);

  return (
    <Form.Control
      {...field}
      as='select'
      onChange={e => setFieldValue(field.name, e.target.value)}
    >
      <Options
        defaultOptionText={`Select ${props.displayName}`}
        options={helpers.getOptionsArrayFromKey(props.options, props.name)}
        optionValue='uuid'
        optionDisplay={props.optionDisplay}
      />
    </Form.Control>
  );
};