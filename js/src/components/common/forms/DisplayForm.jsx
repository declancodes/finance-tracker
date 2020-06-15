import React from 'react';
import { Button } from 'react-bootstrap';
import { EntityForm } from './EntityForm';

export const DisplayForm = ({
  entityName,
  entity,
  getInitialValues,
  isCreateMode,
  options,
  doExtraModifications,
  doSubmit,
  setNotUsing
}) => (
  <div className='display-form'>
    <EntityForm
      entityName={entityName}
      entity={entity}
      getInitialValues={getInitialValues}
      isCreateMode={isCreateMode}
      options={options}
      doExtraModifications={doExtraModifications}
      doSubmit={doSubmit}
      doFinalState={setNotUsing}
    />
    <Button
      variant='dark'
      onClick={setNotUsing}
    >
      Cancel
    </Button>
  </div>
);