import React from 'react';
import { Options } from '../Options';
import { helpers } from '../../../common/helpers';
import pluralize from 'pluralize';
import startCase from 'lodash.startcase';

export const LabeledCategoryFilter = ({
  filterCategory,
  options,
  setFilterCategory
}) => {
  const displayName = startCase(filterCategory.name);
  const htmlFor = `${filterCategory.name}-filter`;

  return (
    <div className='labeled-category-filter'>
      <label htmlFor={htmlFor}>{displayName}</label>
      <select
        name={htmlFor}
        value={filterCategory.value}
        onChange={e => setFilterCategory(filterCategory.name, e.target.value)}
      >
        <Options
          defaultOptionText={`All ${pluralize(displayName)}`}
          options={helpers.getOptionsArrayFromKey(options, filterCategory.name)}
          optionValue={filterCategory.optionValue}
          optionDisplay={filterCategory.optionDisplay}
        />
      </select>
    </div>
  );
};
