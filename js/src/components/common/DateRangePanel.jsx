import React from 'react';
import DatePicker from 'react-datepicker';
import '../../../node_modules/react-datepicker/dist/react-datepicker.css';

export const DateRangePanel = ({
  start,
  end,
  setStart,
  setEnd
}) => (
  <div className='date-range-panel'>
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
  </div>
);
