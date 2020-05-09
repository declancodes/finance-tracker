import React from 'react';
import { Formik, Field, Form } from 'formik';

class CategoryForm extends React.Component {
  render() {
    const c = this.props.category;

    return (
      <div>
        <h2>
          {this.props.isEditMode ? "Edit" : "Create"} {this.props.categoryType} Category
        </h2>
        <Formik
          initialValues={{
            uuid: c ? c.uuid : '',
            name: c ? c.name : '',
            description: c ? c.description : ''
          }}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            if (!c) {
              delete values.uuid;
            }
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
            <button type="submit">
              {this.props.isEditMode ? "Update" : "Create"}
            </button>
          </Form>
        </Formik>
      </div>
    );
  }
}

export default CategoryForm;