import React from 'react';
import { Jumbotron } from 'react-bootstrap';
import { Button } from '../common/Button/Button';

export const HomePage = () => (
  <Jumbotron>
    <h1>Welcome to finance-tracker</h1>
    <p>Track all of your finances in one easy location.</p>
    <p>
      <Button className='primary' href='/accounts'>
        <a href='/accounts'>Get Started</a>
      </Button>
    </p>
  </Jumbotron>
);