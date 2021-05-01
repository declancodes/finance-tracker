import React from 'react';
import './NavItem.scss';

export interface NavItemProps {
  to: string,
  title: string
}

export const NavItem = ({
  to,
  title
}: NavItemProps) => (
  <div className='nav-item-container'>
    <a href={to}>
      {title}
    </a>
  </div>
);
