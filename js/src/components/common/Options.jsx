import React from 'react';

export const Options = ({
  defaultOptionText,
  options,
  optionValue,
  optionDisplay
}) => (
  <>
    <option value='' defaultValue=''>{defaultOptionText}</option>
    {options.length > 0 && (
      options.map(option => (
        <option key={option.uuid} value={option[optionValue]}>
          {option[optionDisplay]}
        </option>
      ))
    )}
  </>
);
