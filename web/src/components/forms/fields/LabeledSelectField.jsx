import React from 'react';
import Select from 'react-select';
import { useField, useFormikContext } from 'formik';
import { getOptionsArrayFromKey } from '../../../common/helpers';

export const LabeledSelectField = ({ ...props }) => {
  const { setFieldValue } = useFormikContext();
  const [field] = useField(props);

  const {
    options,
    name,
    optionDisplay,
    isMulti,
    displayName
  } = props;

  const opts = getOptionsArrayFromKey(options, name)
    .map(o => {
      return {
        value: o.uuid,
        label: optionDisplay(o)
      }
    });

  const getSelected = () => {
    if (!opts || !opts.length) {
      return isMulti ? [] : '';
    }

    return isMulti
      ? opts.filter(o => field.value.indexOf(o.value) >= 0)
      : opts.find(o => o.value === field.value);
  };

  const onChange = value => {
    const vals = isMulti
      ? (value ? value.map(o => o.value) : [])
      : value.value;

    setFieldValue(field.name, vals);
  };

  return (
    <Select
      name={field.name}
      isMulti={isMulti}
      placeholder={`Select ${displayName}...`}
      options={opts}
      value={getSelected()}
      onChange={onChange}
    />
  );
};
