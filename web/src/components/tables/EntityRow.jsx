import React, { useState } from 'react';
import { EntityForm } from '../forms/EntityForm';
import { ModifyRowPanel } from './ModifyRowPanel';
import {
  displayCurrency,
  displayDate,
  displayDecimals,
  displayPercentage
} from '../../common/helpers';

export const EntityRow = ({
  entityName,
  entity,
  getInitialValues,
  options,
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
      {entity.hasOwnProperty('date') && <td>{displayDate(entity.date)}</td>}
      {entity.hasOwnProperty('amount') && <td>{displayCurrency(entity.amount)}</td>}
      {entity.hasOwnProperty('sharePrice') && <td>{displayCurrency(entity.sharePrice)}</td>}
      {entity.hasOwnProperty('shares') && <td>{displayDecimals(entity.shares, 3)}</td>}
      {entity.hasOwnProperty('expenseRatio') && <td>{`${displayPercentage(entity.expenseRatio, 3)}%`}</td>}
      {entity.hasOwnProperty('effectiveExpense') && <td>{displayCurrency(entity.effectiveExpense)}</td>}
      {entity.hasOwnProperty('value') && <td>{displayCurrency(entity.value)}</td>}
      <td>
        {isEditing ? (
          <EntityForm
            entityName={entityName}
            entity={entity}
            getInitialValues={getInitialValues}
            options={options}
            doExtraModifications={doExtraModifications}
            doSubmit={handleUpdate}
            doFinalState={() => setIsEditing(false)}
          />
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
