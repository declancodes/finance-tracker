import React from 'react';
import { Formik } from 'formik';
import { EntityForm } from '../common/forms/EntityForm';

export const CategoryForm = ({ category, categoryType, doSubmit, isEditMode}) => {
  const initialCategoryValues = {
    uuid: category ? category.uuid : '',
    name: category ? category.name : '',
    description: category ? category.description : ''
  };
  return (
    <div>
      <h2>
        {isEditMode ? 'Edit' : 'Create'} {categoryType} Category
      </h2>
      <Formik
        initialValues={initialCategoryValues}
        onSubmit={(values, { setSubmitting, resetForm }) => {
          if (!category) {
            delete values.uuid;
          }
          doSubmit(values);
          setSubmitting(false);
          resetForm();
        }}
      >
        <EntityForm
          entity={initialCategoryValues}
          isEditMode={isEditMode}
        />
      </Formik>
    </div>
  );
};
