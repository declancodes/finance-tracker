import React from 'react';
import { Button } from './Button/Button';

const CustButton = ({
  type,
  display,
  onClick,
  variant
}) => (
  <Button
    type={type}
    onClick={onClick}
    className={variant}
  >
    {display}
  </Button>
);

export const ButtonPair = ({
  variant1,
  type1,
  display1,
  onClick1,
  variant2,
  type2,
  display2,
  onClick2
}) => (
  <div className='button-pair'>
    <CustButton
      type={type1}
      display={display1}
      onClick={onClick1}
      variant={variant1 || 'primary'}
    />
    {' '}
    <CustButton
      type={type2}
      display={display2}
      onClick={onClick2}
      variant={variant2 || 'secondary'}
    />
  </div>
);