import React from 'react';
import { LabeledControl } from '../../common/LabeledControl/LabeledControl';
import { FormDatePicker } from '../../common/FormDatePicker';
import { titleCase } from '../../../common/helpers';
import './LabeledDatePicker.scss';

export const LabeledDatePicker = ({ ...props }) => (
  <LabeledControl
    label={titleCase(props.name)}
  >
    <FormDatePicker
      className='datepicker'
      {...props}
    />
  </LabeledControl>
);
