import React from 'react';
import { LabeledFilter } from '../LabeledFilter/LabeledFilter';
import { FormDatePicker } from '../../common/FormDatePicker';
import { titleCase } from '../../../common/helpers';
import './LabeledDatePicker.scss';

export const LabeledDatePicker = ({ ...props }) => (
  <LabeledFilter
    label={titleCase(props.name)}
  >
    <FormDatePicker
      className='datepicker'
      {...props}
    />
  </LabeledFilter>
);
