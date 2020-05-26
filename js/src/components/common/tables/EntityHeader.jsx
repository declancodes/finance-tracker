import React from 'react';

export const EntityHeader = ({ entity }) => (
  <thead>
    <tr>
      {entity.hasOwnProperty('name') && <th>Name</th>}
      {entity.hasOwnProperty('account') && <th>Account</th>}
      {entity.hasOwnProperty('category') && <th>Category</th>}
      {entity.hasOwnProperty('fund') && <th>Fund</th>}
      {entity.hasOwnProperty('description') && <th>Description</th>}
      {entity.hasOwnProperty('tickerSymbol') && <th>Ticker Symbol</th>}
      {entity.hasOwnProperty('date') && <th>Date</th>}
      {entity.hasOwnProperty('amount') && <th>Amount</th>}
      {entity.hasOwnProperty('sharePrice') && <th>Share Price</th>}
      {entity.hasOwnProperty('shares') && <th>Shares</th>}
      <th>Actions</th>
    </tr>
  </thead>
);