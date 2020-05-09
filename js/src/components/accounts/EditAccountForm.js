import React from 'react';
import { Formik, Field, Form } from 'formik';

class EditAccountForm extends React.Component {
  render() {
    const a = this.props.account;

    return (
      <div>
        <h2>Edit Account</h2>
        <Formik
          initialValues={{
            uuid: a.uuid,
            name: a.name,
            category: a.category.name,
            description: a.description,
            amount: a.amount
          }}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            this.props.doUpdate(values);
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
            <button type="submit">Update</button>
          </Form>
        </Formik>
      </div>
    );
  }
}

export default EditAccountForm;