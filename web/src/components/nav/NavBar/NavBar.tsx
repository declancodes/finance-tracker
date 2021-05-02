import React from 'react';
import './NavBar.scss';

interface NavBarProps {
  children?: React.ReactNode
}

export const NavBar = ({ children }: NavBarProps) => (
  <nav role='navigation' className='nav-bar'>
    <ul>
      {children}
    </ul>
  </nav>
);