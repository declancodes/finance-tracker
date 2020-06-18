import React from 'react';
import { Form } from 'react-bootstrap';
import { Formik } from 'formik';
import { LabeledField } from './fields/LabeledField';
import { ButtonPair } from '../ButtonPair';

const LabeledFieldOrNull = (entity, field, fieldType, props, options, optionDisplay) => (
  entity.hasOwnProperty(field) &&
    <LabeledField
      name={field}
      fieldType={fieldType}
      props={props}
      options={options}
      optionDisplay={optionDisplay}
    />
);

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

          doFinalState();
        }}
      >
        {props => (
          <Form noValidate onSubmit={props.handleSubmit}>
            {LabeledFieldOrNull(entity, 'name', 'text', props)}
            {LabeledFieldOrNull(entity, 'category', null, props, options, 'name')}
            {LabeledFieldOrNull(entity, 'account', null, props, options, 'name')}
            {LabeledFieldOrNull(entity, 'fund', null, props, options, 'tickerSymbol')}
            {LabeledFieldOrNull(entity, 'description', 'text', props)}
            {LabeledFieldOrNull(entity, 'tickerSymbol', 'text', props)}
            {LabeledFieldOrNull(entity, 'date', 'date', props)}
            {LabeledFieldOrNull(entity, 'amount', 'number', props)}
            {LabeledFieldOrNull(entity, 'sharePrice', 'number', props)}
            {LabeledFieldOrNull(entity, 'shares', 'number', props)}
            {LabeledFieldOrNull(entity, 'expenseRatio', 'number', props)}
            <ButtonPair
              type1='submit'
              display1={typeDisplay}
              onClick2={doFinalState}
              display2='Cancel'
            />
          </Form>
        )}
      </Formik>
    </div>
  );
};
