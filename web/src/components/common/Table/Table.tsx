import React from 'react';
import './Table.css';

interface TableProps {
  children: React.ReactNode
}

export const Table = ({ children }: TableProps) => (
  <table className='table-container'>
    {children}
  </table>
)