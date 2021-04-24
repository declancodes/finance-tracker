import React from 'react';

export const Dropdown = ({ linkItems }) => (
  <div>
    {linkItems.map(li => (
      <a
        key={`li-key-${li.display}`}
        href={li.link}
      >
        {li.display}
      </a>
    ))}
  </div>
);