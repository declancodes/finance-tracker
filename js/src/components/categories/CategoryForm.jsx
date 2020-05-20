import React from 'react';
import { Formik } from 'formik';
import { EntityForm } from '../common/forms/EntityForm';

class CategoryForm extends React.Component {
  render() {
    const c = this.props.category;
    const initialCategoryValues = {
      uuid: c ? c.uuid : '',
      name: c ? c.name : '',
      description: c ? c.description : ''
    };

    return (
      <div>
        <h2>
          {this.props.isEditMode ? "Edit" : "Create"} {this.props.categoryType} Category
        </h2>
        <Formik
          initialValues={initialCategoryValues}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            if (!c) {
              delete values.uuid;
            }
            this.props.doSubmit(values);
            setSubmitting(false);
            resetForm();
          }}
        >
          <EntityForm
            entity={initialCategoryValues}
            isEditMode={this.props.isEditMode}
          />
        </Formik>
      </div>
    );
  }
}

export default CategoryForm;