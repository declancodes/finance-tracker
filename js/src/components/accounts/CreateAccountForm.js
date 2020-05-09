import React from 'react';
import { Formik, Field, Form } from 'formik';

class CreateAccountForm extends React.Component {
  render() {
    return (
      <div>
        <h2>Create New Account</h2>
        <Formik
          initialValues={{
            name: '',
            category: '',
            description: '',
            amount: 0
          }}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            this.props.doSubmit(values);
            setSubmitting(false);
            resetForm();
          }}
        >
          <Form>
            <label htmlFor="name">Name</label>
            <Field name="name" type="text" />
            <label htmlFor="category">Category</label>
            <Field name="category" type="text" />
            <label htmlFor="description">Description</label>
            <Field name="description" type="text" />
            <label htmlFor="amount">Amount</label>
            <Field name="amount" type="number" />
            <button type="submit">Create</button>
          </Form>
        </Formik>
      </div>
    );
  }
}

export default CreateAccountForm;