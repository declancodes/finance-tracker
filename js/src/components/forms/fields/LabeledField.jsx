import React from 'react';
import { Col, Form } from 'react-bootstrap';
import { DatePickerField } from './DatePickerField';
import { LabeledSelectField } from './LabeledSelectField';
import { helpers } from '../../../common/helpers';

export const LabeledField = ({
  name,
  fieldType,
  options,
  optionDisplay,
  props
}) => {
  const displayName = helpers.titleCase(name);

  return (
    <div className='container-fluid'>
      <Form.Row>
        <Form.Label
          column
          sm={1}
        >
          {displayName}
        </Form.Label>
        <Col sm={6}>
          {options !== undefined ? (
            <LabeledSelectField
              name={name}
              displayName={displayName}
              options={options}
              optionDisplay={optionDisplay}
            />
          ) : fieldType === 'date' ? (
            <DatePickerField name={name}/>
          ) : (
            <Form.Control
              name={name}
              type={fieldType}
              value={props.values[name]}
              onChange={props.handleChange}
            />
          )}
        </Col>
      </Form.Row>
    </div>
  );
};