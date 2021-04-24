import React, { MouseEventHandler, ReactChildren } from 'react';
import './Button.css';

interface ButtonProps {
  className?: string,
  type: 'button' | 'submit' | 'reset',
  onClick?: MouseEventHandler<HTMLButtonElement>,
  children?: React.ReactNode
}

export const Button = ({
  className,
  type,
  onClick,
  children
} : ButtonProps) => {
  const buttonClass = `button ${className}`;

  return (
    <button
      className={buttonClass}
      type={type}
      onClick={onClick}
    >
      {children}
    </button>
  );
};