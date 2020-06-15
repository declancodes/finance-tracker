import React from 'react';
import { Button, Jumbotron } from 'react-bootstrap'

export const HomePage = () => (
  <Jumbotron>
    <h1>Welcome to finance-tracker</h1>
    <p>Track all of your finances in one easy location.</p>
    <p>
      <Button
        variant='dark'
        href='/accounts'
      >
        Get Started
      </Button>
    </p>
  </Jumbotron>
);