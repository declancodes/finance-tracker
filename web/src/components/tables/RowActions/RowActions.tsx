import React from 'react';
import { Button } from '../../common/Button/Button';
import { EditIcon } from '../../common/icons/EditIcon/EditIcon';
import { TrashIcon } from '../../common/icons/TrashIcon/TrashIcon';
import './RowActions.scss';

interface RowActionsProps {
  handleEdit: React.MouseEventHandler,
  handleDelete: React.MouseEventHandler
}

const iconFill = "#adbac7";

export const RowActions = ({ handleEdit, handleDelete }: RowActionsProps) => (
  <div className="button-panel">
    <Button
      type="button"
      title="Edit"
      className="primary"
      onClick={handleEdit}
    >
      <EditIcon fill={iconFill} />
    </Button>
    <span>|</span>
    <Button
      type="button"
      title="Delete"
      className="danger"
      onClick={handleDelete}
    >
      <TrashIcon fill={iconFill} />
    </Button>
  </div>
);