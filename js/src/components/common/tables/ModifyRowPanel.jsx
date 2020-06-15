import React from 'react';
import { Button } from 'react-bootstrap';

export const ModifyRowPanel = ({ handleEdit, handleDelete}) => (
  <div>
    <Button
      variant='dark'
      onClick={handleEdit}
    >
      Edit
    </Button>
    <Button
      variant='danger'
      onClick={handleDelete}
    >
      Delete
    </Button>
  </div>
);
