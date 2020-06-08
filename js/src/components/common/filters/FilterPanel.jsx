import React from 'react';
import { LabeledCategoryFilter } from './LabeledCategoryFilter';
import { LabeledDatePicker } from './LabeledDatePicker';
import './FilterPanel.css';
import '../../../../node_modules/react-datepicker/dist/react-datepicker.css';

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
    {filterCategories.length > 0 &&
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
