import React from 'react';
import { NavDropdown } from 'react-bootstrap';
import { LinkItem } from './LinkItem';

export const Dropdown = ({ title, linkItems }) => (
  <NavDropdown title={title}>
    {linkItems.map(li => (
      <LinkItem
        key={`li-key-${li.display}`}
        link={li.link}
        display={li.display}
      />
    ))}
  </NavDropdown>
);