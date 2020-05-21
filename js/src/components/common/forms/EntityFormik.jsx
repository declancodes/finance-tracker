import React from 'react';
import { Form, Formik } from 'formik';
import { DatePickerField } from './DatePickerField';
import { LabeledField } from './LabeledField';

class EntityFormik extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      options: []
    };
  }

  componentDidMount() {
    if (this.props.getOptions !== undefined) {
      this.props.getOptions()
        .then(response => this.setState({ options: response }));
    }
  }

  render() {
    const e = this.props.entity;
    return (
      <div>
        <h2>
          {this.props.isCreateMode ? 'Create' : 'Edit'} {this.props.entityName}
        </h2>
        <Formik
          initialValues={e}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            if (this.props.isCreateMode) {
              delete values.uuid;
            }
  
            if (this.props.doExtraModifications !== undefined) {
              this.props.doExtraModifications(values);
            }
  
            this.props.doSubmit(values);
            setSubmitting(false);
            resetForm();
          }}
        >
          <Form>
            {e.hasOwnProperty('name') && <LabeledField name='name' fieldType='text'/>}
            {e.hasOwnProperty('category') && <LabeledField name='category' options={this.state.options}/>}
            {e.hasOwnProperty('account') && <LabeledField name='account' options={this.state.options}/>}
            {e.hasOwnProperty('description') && <LabeledField name='description' fieldType='text'/>}
            {e.hasOwnProperty('date') && <DatePickerField name='date'/>}
            {e.hasOwnProperty('amount') && <LabeledField name='amount' fieldType='number'/>}
            <button type='submit'>
              {this.props.isCreateMode ? 'Create' : 'Update'}
            </button>
          </Form>
        </Formik>
      </div>
    );
  }
}

export default EntityFormik;
