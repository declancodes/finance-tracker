import React from 'react';
import './Button.css';

export const Button = ({ name, handleFunc }) => (
  <button
    className={`${name.toLowerCase()}-button`}
    onClick={handleFunc}>
      {name}
  </button>
);
