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
      <Formik
        initialValues={{ name: '', description: '' }}
        onSubmit={(values, { setSubmitting }) => {
          this.createAccountCategory(values);
          setSubmitting(false)
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
    );
  }
}

export default CreateAccountCategoryForm;