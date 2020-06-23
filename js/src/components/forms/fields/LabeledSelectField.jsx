import React from 'react';
import { useField, useFormikContext } from 'formik';
import { helpers } from '../../../common/helpers';
import { DarklyReactSelect } from '../../common/DarklyReactSelect';

export const LabeledSelectField = ({ ...props }) => {
  const { setFieldValue } = useFormikContext();
  const [field] = useField(props);
  const opts = helpers
    .getOptionsArrayFromKey(props.options, props.name)
    .map(o => {
      return {
        value: o['uuid'],
        label: o[props.optionDisplay]
      }
    });
  const selected = opts.length > 0 ?
    opts.filter(o => o.value === field.value)[0] :
    [];

  return (
    <DarklyReactSelect
      placeholder={`Select ${props.displayName}...`}
      options={opts}
      value={selected}
      onChange={value => setFieldValue(field.name, value)}
    />
  );
};