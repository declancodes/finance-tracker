import React from 'react';
import { Formik, Field, Form } from 'formik';

class EditCategoryForm extends React.Component {
  render() {
    return (
      <div>
        <h2>Edit {this.props.categoryType} Category</h2>
        <Formik
          initialValues={{
            uuid: this.props.category.uuid,
            name: this.props.category.name,
            description: this.props.category.description
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
            <label htmlFor="description">Description</label>
            <Field name="description" type="text" />
            <button type="submit">Update</button>
          </Form>
        </Formik>
      </div>
    );
  }
}

export default EditCategoryForm;