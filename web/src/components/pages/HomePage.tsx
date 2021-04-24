import React from 'react';
import { Button } from '../common/Button/Button';

export const HomePage = () => (
  <div>
    <h1>Welcome to finance-tracker</h1>
    <p>Track all of your finances in one easy location.</p>
    <p>
      <Button className='primary' type='button'>
        <a href='/accounts'>Get Started</a>
      </Button>
    </p>
  </div>
);