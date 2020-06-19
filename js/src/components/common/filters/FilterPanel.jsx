import React from 'react';
import { Col, Form } from 'react-bootstrap';
import { LabeledCategoryFilter } from './LabeledCategoryFilter';
import { LabeledDatePicker } from './LabeledDatePicker';

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
  <Form className='filter-panel'>
    <div className='container-fluid'>
      <Form.Row>
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
            <Form.Group key={`lcf-${fc.name}`} as={Col} xs='auto'>
              <LabeledCategoryFilter
                filterCategory={fc}
                options={options}
                setFilterCategory={setFilterCategory}
              />
            </Form.Group>
          )
        )}
      </Form.Row>
    </div>
  </Form>
);
