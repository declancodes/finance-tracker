import React from 'react';
import { ButtonPair } from '../common/ButtonPair';

export const ModifyRowPanel = ({ handleEdit, handleDelete }) => (
  <ButtonPair
    onClick1={handleEdit}
    display1='Edit'
    variant2='danger'
    onClick2={handleDelete}
    display2='Delete'
  />
);
