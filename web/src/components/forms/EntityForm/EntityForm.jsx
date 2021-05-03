import React from 'react';
import { FieldArray, Formik, Field } from 'formik';
import { LabeledField } from '../fields/LabeledField/LabeledField';
import { Button } from '../../common/Button/Button';
import { ButtonPair } from '../../common/ButtonPair';
import { getOptionsArrayFromKey } from '../../../common/helpers';
import './EntityForm.scss';

const LabeledFieldOrNull = (
  entity,
  field,
  fieldType,
  props,
  options,
  optionDisplay,
  isMulti
) => (
  entity.hasOwnProperty(field) &&
    <LabeledField
      name={field}
      fieldType={fieldType}
      props={props}
      options={options}
      optionDisplay={optionDisplay}
      isMulti={isMulti}
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
      <h4>
        {`${typeDisplay} ${entityName}`}
      </h4>
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
          <form
            className='entity-form'
            noValidate
            onSubmit={props.handleSubmit}
          >
            {LabeledFieldOrNull(entity, 'name', 'text', props)}
            {LabeledFieldOrNull(entity, 'category', null, props, options, o => o.name, false)}
            {LabeledFieldOrNull(entity, 'account', null, props, options, o => o.name, false)}
            {LabeledFieldOrNull(entity, 'fund', null, props, options, o => o.tickerSymbol, false)}
            {LabeledFieldOrNull(entity, 'description', 'text', props)}
            {LabeledFieldOrNull(entity, 'tickerSymbol', 'text', props)}
            {LabeledFieldOrNull(entity, 'date', 'date', props)}
            {LabeledFieldOrNull(entity, 'amount', 'number', props)}
            {LabeledFieldOrNull(entity, 'sharePrice', 'number', props)}
            {LabeledFieldOrNull(entity, 'shares', 'number', props)}
            {LabeledFieldOrNull(entity, 'expenseRatio', 'number', props)}
            {LabeledFieldOrNull(entity, 'holdings', null, props, options, o => `${o.account.name}: ${o.fund.tickerSymbol}`, true)}
            {entity.hasOwnProperty('assetAllocation') &&
              <div>
                <FieldArray name='assetAllocation'>
                  {innerProps => (
                    <div>
                      {props.values.assetAllocation.map((aa, i) => (
                        <div key={i}>
                          <label column sm={1}>
                            Asset Category
                          </label>
                          <div sm={3}>
                            <Field
                              name={`assetAllocation.${i}.category`}
                              component='select'
                            >
                              {getOptionsArrayFromKey(options, 'category').map(c => (
                                <option key={c.uuid} value={c.uuid}>
                                  {c.name}
                                </option>
                              ))}
                            </Field>
                          </div>
                          <label column sm={1}>
                            Percentage
                          </label>
                          <div sm={3}>
                            <Field
                              name={`assetAllocation.${i}.percentage`}
                              type='number'
                            />
                          </div>
                          <div sm={3}>
                            <Button onClick={() => { innerProps.remove(i); }}>
                              Remove
                            </Button>
                            <Button onClick={() => { innerProps.insert(i, aa); console.log(props.values); }}>
                              Insert
                            </Button>
                          </div>
                        </div>
                      ))}
                      <Button onClick={() => { innerProps.push({ category: '', percentage: 0 }); }}>
                        Add Asset Category
                      </Button>
                    </div>
                  )}
                </FieldArray>
              </div>
            }
            <ButtonPair
              type1='submit'
              display1={typeDisplay}
              onClick2={doFinalState}
              display2='Cancel'
            />
          </form>
        )}
      </Formik>
    </div>
  );
};
