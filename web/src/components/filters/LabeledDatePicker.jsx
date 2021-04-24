import React from 'react';
import { FormDatePicker } from '../common/FormDatePicker';
import { titleCase } from '../../common/helpers';

export const LabeledDatePicker = ({ ...props }) => (
  <div>
    <label>{titleCase(props.name)}</label>
    <FormDatePicker
      {...props}
    />
  </div>
);
