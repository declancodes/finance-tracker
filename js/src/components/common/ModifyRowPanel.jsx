import React from 'react';
import { Button } from './Button';

export const ModifyRowPanel = ({ handleEdit, handleDelete}) => (
  <div>
    <Button
      name='Edit'
      handleFunc={handleEdit}
    />
    <Button
      name='Delete'
      handleFunc={handleDelete}
    />
  </div>
);
