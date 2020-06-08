import React from 'react';
import { LabeledField } from './LabeledField';
import { helpers } from '../../../../common/helpers';

export const LabeledSelectField = ({
  name,
  options,
  optionDisplay
}) => (
  <LabeledField
    name={name}
    options={helpers.getOptionsArrayFromKey(options, name)}
    optionDisplay={optionDisplay}
  />
);