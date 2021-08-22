import React from 'react';
import { Button } from '../../common/Button/Button';
import { EditIcon } from '../../common/icons/EditIcon/EditIcon';
import { TrashIcon } from '../../common/icons/TrashIcon/TrashIcon';
import './RowActions.scss';

interface RowActionsProps {
  handleEdit: React.MouseEventHandler,
  handleDelete: React.MouseEventHandler
}

export const RowActions = ({ handleEdit, handleDelete }: RowActionsProps) => (
  <div className="button-panel">
    <Button
      type="button"
      title="Edit"
      onClick={handleEdit}
    >
      <EditIcon/>
    </Button>
    <span>|</span>
    <Button
      type="button"
      title="Delete"
      onClick={handleDelete}
    >
      <TrashIcon/>
    </Button>
  </div>
);