import React from 'react';
import './NavBar.css';

interface NavBarProps {
  children?: React.ReactNode
}

export const NavBar = ({ children }: NavBarProps) => (
  <nav className='nav-bar'>
    {children}
  </nav>
);