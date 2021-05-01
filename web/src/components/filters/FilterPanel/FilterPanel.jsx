import React from 'react';
import { LabeledCategoryFilter } from '../LabeledCategoryFilter/LabeledCategoryFilter';
import { LabeledDatePicker } from '../LabeledDatePicker/LabeledDatePicker';
import './FilterPanel.scss';

export const FilterPanel = ({
  usesDates,
  start,
  end,
  filterCategories,
  options,
  setStart,
  setEnd,
  setFilterCategory
}) => (
  <div className='filter-panel'>
    {usesDates &&
      <>
        <LabeledDatePicker
          name='from'
          selected={start}
          onChange={val => setStart(val)}
        />
        <LabeledDatePicker
          name='to'
          selected={end}
          onChange={val => setEnd(val)}
        />
      </>
    }
    {filterCategories && filterCategories.length > 0 &&
      filterCategories.map(fc => (
        <LabeledCategoryFilter
          key={`lcf-${fc.name}`}
          filterCategory={fc}
          options={options}
          setFilterCategory={setFilterCategory}
        />
      )
    )}
  </div>
);
