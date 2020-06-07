import React from 'react';
import { LabeledDatePicker } from './LabeledDatePicker';
import { Options } from '../Options';
import './FilterPanel.css';
import '../../../../node_modules/react-datepicker/dist/react-datepicker.css';

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
}) => (
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
    <label htmlFor='categoryFilter'>{filterCategoryName}</label>
    <select
      name='categoryFilter'
      value={filterCategory}
      onChange={e => setFilterCategory(e.target.value)}
    >
      <Options
        entityName={filterCategoryName}
        options={filterCategoryOptions[0]}
        optionValue='name'
        optionDisplay='name'
      />
    </select>
  </div>
);
