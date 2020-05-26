import React from 'react';

export const Options = ({
  entityName,
  options,
  optionValue,
  optionDisplay
}) => (
  <>
    <option value='' defaultValue=''>Select {entityName}</option>
    {options.length > 0 && (
      options.map(option => (
        <option key={option.uuid} value={option[optionValue]}>
          {option[optionDisplay]}
        </option>
      ))
    )}
  </>
);
