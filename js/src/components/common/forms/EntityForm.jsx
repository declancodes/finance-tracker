import React from 'react';
import { Form } from 'formik';
import { DatePickerField } from './DatePickerField';
import { LabeledField } from './LabeledField';

export const EntityForm = ({ entity, options, isCreateMode}) => (
  <Form>
    {entity.hasOwnProperty('name') && <LabeledField name='name' fieldType='text'/>}
    {entity.hasOwnProperty('category') && <LabeledField name='category' options={options}/>}
    {entity.hasOwnProperty('account') && <LabeledField name='account' options={options}/>}
    {entity.hasOwnProperty('description') && <LabeledField name='description' fieldType='text'/>}
    {entity.hasOwnProperty('date') && <DatePickerField name='date'/>}
    {entity.hasOwnProperty('amount') && <LabeledField name='amount' fieldType='number'/>}
    <button type='submit'>
      {isCreateMode ? 'Create' : 'Update'}
    </button>
  </Form>
);
