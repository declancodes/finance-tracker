import React from 'react';
import { Form, Formik } from 'formik';
import { LabeledField } from './fields/LabeledField';

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
            <LabeledField name='category' options={options[0]} optionDisplay='name'/>
          }
          {entity.hasOwnProperty('account') &&
            <LabeledField name='account' options={options[0]} optionDisplay='name'/>
          }
          {entity.hasOwnProperty('fund') &&
            <LabeledField name='fund' options={options[1]} optionDisplay='tickerSymbol'/>
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
          <button type='submit'>
            {typeDisplay}
          </button>
        </Form>
      </Formik>
    </div>
  );
};
