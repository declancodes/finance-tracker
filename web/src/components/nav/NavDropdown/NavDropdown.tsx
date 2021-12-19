import React, { useState } from 'react';
import { NavItem, NavItemProps } from '../NavItem/NavItem';
import './NavDropdown.scss';

interface NavDropdownProps {
  title: string,
  navItems: NavItemProps[]
}

export const NavDropdown = ({
  title,
  navItems
}: NavDropdownProps) => {
  const [isOpen, setIsOpen] = useState<boolean>(false);

  const onClick = (_: any) => {
    setIsOpen(!isOpen);
  };

  return (
    <div
      className='nav-dropdown-container'
      tabIndex={0}
      onClick={onClick}
    >
      <NavItem
        to='#'
        className='dropdown'
      >
        <div className='title'>{title}</div>
        <div className='triangle'/>
      </NavItem>
      {isOpen &&
        <ul className='nav-items'>
          {navItems.map(ni => (
            <NavItem
              key={`nav-dropdown-item-${ni.title}`}
              to={ni.to}
            >
              {ni.title}
            </NavItem>
          ))}
        </ul>
      }
    </div>
  );
};
