import React from 'react';
import { Nav, NavDropdown } from 'react-bootstrap';
import { LinkContainer } from 'react-router-bootstrap';

export const LinkItem = ({ type, link, display }) => (
  <LinkContainer to={link}>
    {type === 'link' ? (
      <Nav.Link>{display}</Nav.Link>
    ) : (
      <NavDropdown.Item>{display}</NavDropdown.Item>
    )}
  </LinkContainer>
);