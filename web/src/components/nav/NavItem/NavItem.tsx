import React from 'react';
import './NavItem.scss';

export interface NavItemProps {
  to: string;
  className?: string;
  children?: React.ReactNode;
}

export const NavItem = ({
  to,
  className,
  children,
}: NavItemProps) => {
  const baseButtonClass = 'nav-item';

  const fullClassName = className === undefined
    ? baseButtonClass
    : `${baseButtonClass} ${className}`;

  return (
    <li className={fullClassName}>
      <a href={to}>
        {children}
      </a>
    </li>
  );
}
