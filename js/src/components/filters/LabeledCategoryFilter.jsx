import React from 'react';
import { Form } from 'react-bootstrap';
import { DarklyReactSelect } from '../common/DarklyReactSelect';
import { helpers } from '../../common/helpers';
import startCase from 'lodash.startcase';

export const LabeledCategoryFilter = ({
  filterCategory,
  options,
  setFilterCategory
}) => {
  const displayName = startCase(filterCategory.name);
  const opts = () => helpers
    .getOptionsArrayFromKey(options, filterCategory.name)
    .map(o => {
      return {
        value: o[filterCategory.optionValue],
        label: o[filterCategory.optionDisplay]
      }
    });

  return (
    <>
      <Form.Label>{displayName}</Form.Label>
      <DarklyReactSelect
        isMulti
        placeholder='Filter by Category...'
        options={opts()}
        value={filterCategory.value}
        onChange={value => setFilterCategory(filterCategory.name, value)}
      />
    </>
  );
};
