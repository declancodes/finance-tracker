import React from 'react';
import { LabeledDatePicker } from './LabeledDatePicker';
import { Options } from '../Options';
import '../../../../node_modules/react-datepicker/dist/react-datepicker.css';

export const FilterPanel = ({
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
    <label htmlFor='categoryFilter'>{filterCategoryName}</label>
    <select
      name='categoryFilter'
      value={filterCategory}
      onChange={e => setFilterCategory(e.target.value)}
    >
      <Options
        entityName={filterCategoryName}
        options={filterCategoryOptions}
        optionValue='name'
      />
    </select>
  </div>
);
