import React from 'react';
import './Table.scss';

interface TableProps {
  children: React.ReactNode
}

export const Table = ({ children }: TableProps) => (
  <table className='table-container'>
    {children}
  </table>
)