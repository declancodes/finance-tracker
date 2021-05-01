import React from 'react';
import './LabeledFilter.scss';

interface LabeledFilterProps {
  label: string,
  children: React.ReactNode
}

export const LabeledFilter = ({
  label,
  children
}: LabeledFilterProps) => {
  return (
    <div className='labeled-filter'>
      <label>{label}</label>
      {children}
    </div>
  );
};
