import React from 'react';
import './NavBar.scss';

interface NavBarProps {
  children?: React.ReactNode
}

export const NavBar = ({ children }: NavBarProps) => (
  <nav className='nav-bar'>
    {children}
  </nav>
);