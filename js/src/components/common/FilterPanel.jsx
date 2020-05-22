import React from 'react';
import DatePicker from 'react-datepicker';
import { Options } from './Options';
import '../../../node_modules/react-datepicker/dist/react-datepicker.css';

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
    <label htmlFor='start'>From</label>
    <DatePicker
      name='start'
      selected={start}
      onChange={val => setStart(val)}
    />
    <label htmlFor='end'>To</label>
    <DatePicker
      name='end'
      selected={end}
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
