import React, { useState } from 'react';
import { Button } from '../Button';
import { EntityForm } from '../forms/EntityForm';
import { ModifyRowPanel } from './ModifyRowPanel';
import { helpers } from '../../../common/helpers';

export const EntityRow = ({
  entityName,
  entity,
  getInitialValues,
  options1,
  options2,
  doExtraModifications,
  handleUpdate,
  handleDelete
}) => {
  const [isEditing, setIsEditing] = useState(false);
  return (
    <tr>
      {entity.hasOwnProperty('name') && <td>{entity.name}</td>}
      {entity.hasOwnProperty('account') && <td>{entity.account.name}</td>}
      {entity.hasOwnProperty('category') && <td>{entity.category.name}</td>}
      {entity.hasOwnProperty('fund') && <td>{entity.fund.tickerSymbol}</td>}
      {entity.hasOwnProperty('description') && <td>{entity.description}</td>}
      {entity.hasOwnProperty('tickerSymbol') && <td>{entity.tickerSymbol}</td>}
      {entity.hasOwnProperty('date') && <td>{helpers.displayDate(entity.date)}</td>}
      {entity.hasOwnProperty('amount') && <td>{`$${entity.amount}`}</td>}
      {entity.hasOwnProperty('sharePrice') && <td>{`$${entity.sharePrice}`}</td>}
      {entity.hasOwnProperty('shares') && <td>{entity.shares}</td>}
      <td>
        {isEditing ? (
          <div>
            <EntityForm
              entityName={entityName}
              entity={entity}
              getInitialValues={getInitialValues}
              isCreateMode={false}
              options1={options1}
              options2={options2}
              doExtraModifications={doExtraModifications}
              doSubmit={handleUpdate}
              doFinalState={() => setIsEditing(false)}
            />
            <Button
              name='Cancel'
              handleFunc={() => setIsEditing(false)}
            />
          </div>
        ) : (
          <ModifyRowPanel
            handleEdit={() => setIsEditing(true)}
            handleDelete={() => handleDelete(entity.uuid)}
          />
        )}
      </td>
    </tr>
  );
};
