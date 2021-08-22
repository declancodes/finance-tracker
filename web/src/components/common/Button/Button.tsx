import React, { MouseEventHandler } from 'react';
import './Button.scss';

interface ButtonProps {
  className?: string,
  type: 'button' | 'submit' | 'reset',
  title?: string,
  onClick?: MouseEventHandler<HTMLButtonElement>,
  children?: React.ReactNode
}

export const Button = ({
  className,
  type,
  title,
  onClick,
  children
} : ButtonProps) => {
  const buttonClass = getButtonClass(className);

  return (
    <button
      className={buttonClass}
      type={type}
      title={title}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

const getButtonClass = (className: string | undefined): string => {
  const baseButtonClass = 'button';

  return className === undefined
    ? baseButtonClass
    : `${baseButtonClass} ${className}`;
}