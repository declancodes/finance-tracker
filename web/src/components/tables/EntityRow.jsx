import React, { useState } from 'react';
import { EntityForm } from '../forms/EntityForm/EntityForm';
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
      {entity.hasOwnProperty('amount') && <td className='number'>{displayCurrency(entity.amount)}</td>}
      {entity.hasOwnProperty('sharePrice') && <td className='number'>{displayCurrency(entity.sharePrice)}</td>}
      {entity.hasOwnProperty('shares') && <td className='number'>{displayDecimals(entity.shares, 3)}</td>}
      {entity.hasOwnProperty('expenseRatio') && <td className='number'>{`${displayPercentage(entity.expenseRatio, 3)}%`}</td>}
      {entity.hasOwnProperty('effectiveExpense') && <td className='number'>{displayCurrency(entity.effectiveExpense)}</td>}
      {entity.hasOwnProperty('isPrivate') && <td>{entity.isPrivate.toString()}</td>}
      {entity.hasOwnProperty('holdings') &&
        <td>
          <ul>
            {entity.holdings.length > 0 ? (
              entity.holdings.map(h => (
                <li key={h.holding.uuid}>
                  {`${h.holding.account.name}: ${h.holding.fund.name}`}
                </li>
              ))
            ) : (
              <li>None</li>
            )}
          </ul>
        </td>
      }
      {entity.hasOwnProperty('assetAllocation') &&
        <td>
          <ul>
            {entity.assetAllocation.length > 0 ? (
              entity.assetAllocation.map(aa => (
                <li key={aa.category.uuid}>
                  {`${aa.category.name}: ${aa.percentage}`}
                </li>
              ))
            ) : (
              <li>None</li>
            )}
          </ul>
        </td>
      }
      {entity.hasOwnProperty('value') && <td className='number'>{displayCurrency(entity.value)}</td>}
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
