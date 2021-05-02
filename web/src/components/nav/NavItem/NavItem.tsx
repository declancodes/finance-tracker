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
  <li className='nav-item'>
    <a href={to}>
      {title}
    </a>
  </li>
);
