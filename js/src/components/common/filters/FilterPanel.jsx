import React from 'react';
import { LabeledDatePicker } from './LabeledDatePicker';
import { Options } from '../Options';
import './FilterPanel.css';
import '../../../../node_modules/react-datepicker/dist/react-datepicker.css';
import { helpers } from '../../../common/helpers';
import startCase from 'lodash.startcase';

export const FilterPanel = ({
  usesDates,
  start,
  end,
  filterCategory,
  filterCategoryName,
  filterCategoryOptions,
  setStart,
  setEnd,
  setFilterCategory
}) => {
  const filterCategoryDisplayName = startCase(filterCategoryName);
  return (
    <div className='filter-panel'>
      {usesDates &&
        <>
          <LabeledDatePicker
            name='from'
            initial={start}
            onChange={val => setStart(val)}
          />
          <LabeledDatePicker
            name='to'
            initial={end}
            onChange={val => setEnd(val)}
          />
        </>
      }
      <label htmlFor='categoryFilter'>{filterCategoryDisplayName}</label>
      <select
        name='categoryFilter'
        value={filterCategory}
        onChange={e => setFilterCategory(e.target.value)}
      >
        <Options
          entityName={filterCategoryDisplayName}
          options={helpers.getOptionsFromKey(filterCategoryOptions, filterCategoryName)}
          optionValue='name'
          optionDisplay='name'
        />
      </select>
    </div>
  );
};
