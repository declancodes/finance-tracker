import React from 'react';
import { Button } from 'react-bootstrap';
import { Form, Formik } from 'formik';
import { LabeledField } from './fields/LabeledField';
import { LabeledSelectField } from './fields/LabeledSelectField';

export const EntityForm = ({
  entityName,
  entity,
  getInitialValues,
  isCreateMode,
  options,
  doExtraModifications,
  doSubmit,
  doFinalState
}) => {
  const typeDisplay = isCreateMode ? 'Create' : 'Update';
  const initialValues = getInitialValues === undefined ?
    entity :
    getInitialValues(entity);

  return (
    <div className='entity-form'>
      <h2>
        {`${typeDisplay} ${entityName}`}
      </h2>
      <Formik
        initialValues={initialValues}
        onSubmit={(values, { setSubmitting, resetForm }) => {
          if (isCreateMode) {
            delete values.uuid;
          }

          if (doExtraModifications !== undefined) {
            doExtraModifications(values);
          }

          doSubmit(values);
          setSubmitting(false);
          resetForm();

          if (doFinalState !== undefined) {
            doFinalState();
          }
        }}
      >
        <Form>
          {entity.hasOwnProperty('name') && 
            <LabeledField name='name' fieldType='text'/>
          }
          {entity.hasOwnProperty('category') &&
            <LabeledSelectField name='category' options={options} optionDisplay='name'/>
          }
          {entity.hasOwnProperty('account') &&
            <LabeledSelectField name='account' options={options} optionDisplay='name'/>
          }
          {entity.hasOwnProperty('fund') &&
            <LabeledSelectField name='fund' options={options} optionDisplay='tickerSymbol'/>
          }
          {entity.hasOwnProperty('description') &&
            <LabeledField name='description' fieldType='text'/>
          }
          {entity.hasOwnProperty('tickerSymbol') &&
            <LabeledField name='tickerSymbol' fieldType='text'/>
          }
          {entity.hasOwnProperty('date') &&
            <LabeledField name='date' fieldType='date'/>
          }
          {entity.hasOwnProperty('amount') &&
            <LabeledField name='amount' fieldType='number'/>
          }
          {entity.hasOwnProperty('sharePrice') &&
            <LabeledField name='sharePrice' fieldType='number'/>
          }
          {entity.hasOwnProperty('shares') &&
            <LabeledField name='shares' fieldType='number'/>
          }
          {entity.hasOwnProperty('expenseRatio') &&
            <LabeledField name='expenseRatio' fieldType='number'/>
          }
          <Button variant='success' type='submit'>
            {typeDisplay}
          </Button>
        </Form>
      </Formik>
    </div>
  );
};
