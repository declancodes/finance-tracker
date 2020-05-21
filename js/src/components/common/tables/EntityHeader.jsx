import React from 'react';

export const EntityHeader = ({ entity }) => (
  <thead>
    <tr>
      <th>Name</th>
      {entity.hasOwnProperty('category') && <th>Category</th>}
      {entity.hasOwnProperty('account') && <th>Account</th>}
      <th>Description</th>
      {entity.hasOwnProperty('date') && <th>Date</th>}
      {entity.hasOwnProperty('amount') && <th>Amount</th>}
      <th>Actions</th>
    </tr>
  </thead>
);