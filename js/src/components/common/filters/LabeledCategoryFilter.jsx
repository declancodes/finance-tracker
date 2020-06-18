import React from 'react';
import { Form } from 'react-bootstrap';
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

  return (
    <>
      <Form.Label>{displayName}</Form.Label>
      <Form.Control as='select'
        value={filterCategory.value}
        onChange={e => setFilterCategory(filterCategory.name, e.target.value)}
      >
        <Options
          defaultOptionText={`All ${pluralize(displayName)}`}
          options={helpers.getOptionsArrayFromKey(options, filterCategory.name)}
          optionValue={filterCategory.optionValue}
          optionDisplay={filterCategory.optionDisplay}
        />
      </Form.Control>
    </>
  );
};
