import React from 'react';
import { Formik, Field, Form } from 'formik';
import axios from 'axios'

const API_URL = "http://localhost:8080/accountcategories"

class CreateAccountCategoryForm extends React.Component {
  createAccountCategory(accountCategory) {
    axios.post(API_URL, accountCategory).then(response => response.data)
    .then((data) => {
      console.log(data)
    })
  }

  render() {
    return (
      <div>
        <h2>Create New Account Category</h2>
        <Formik
          initialValues={{ name: '', description: '' }}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            this.props.doSubmit(values);
            setSubmitting(false);
            resetForm();
          }}
        >
          <Form>
            <label htmlFor="name">Name</label>
            <Field name="name" type="text" />
            <label htmlFor="description">Description</label>
            <Field name="description" type="text" />
            <button type="submit">Create</button>
          </Form>
        </Formik>
      </div>
    );
  }
}

export default CreateAccountCategoryForm;