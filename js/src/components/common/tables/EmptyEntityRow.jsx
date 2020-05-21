import React from 'react';

export const EmptyEntityRow = ({ columnLength, entityPlural }) => (
  <tr>
    <td colSpan={columnLength}>No {entityPlural}</td>
  </tr>
);