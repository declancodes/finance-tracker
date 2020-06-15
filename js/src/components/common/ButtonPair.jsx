import React from 'react';
import { Button } from 'react-bootstrap';

const CustButton = ({
  type,
  display,
  onClick,
  variant,
  size,
}) => (
  <Button
    type={type}
    onClick={onClick}
    variant={variant}
    size={size}
  >
    {display}
  </Button>
);

export const ButtonPair = ({
  size,
  variant1,
  type1,
  display1,
  onClick1,
  type2,
  display2,
  onClick2,
  variant2
}) => (
  <div className='button-pair'>
    <CustButton
      type={type1}
      display={display1}
      onClick={onClick1}
      variant={variant1 || 'primary'}
      size={size}
    />
    {' '}
    <CustButton
      type={type2}
      display={display2}
      onClick={onClick2}
      variant={variant2 || 'secondary'}
      size={size}
    />
  </div>
);