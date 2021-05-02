import React from 'react';
import './LabeledControl.scss';

interface LabeledControlProps {
  label: string,
  children: React.ReactNode
}

export const LabeledControl = ({
  label,
  children
}: LabeledControlProps) => {
  return (
    <div className='labeled-control'>
      <label>{label}</label>
      {children}
    </div>
  );
};