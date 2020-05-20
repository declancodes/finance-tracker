import React from 'react';
import { Formik } from 'formik';
import { EntityForm } from './EntityForm';

export const EntityFormik = ({
  entityName,
  entity,
  options,
  isCreateMode,
  doExtraModifications,
  doSubmit
}) => (
  <div>
    <h2>
      {isCreateMode ? 'Create' : 'Edit'} {entityName}
    </h2>
    <Formik
      initialValues={entity}
      onSubmit={(values, { setSubmitting, resetForm }) => {
        if (isCreateMode) {
          delete values.uuid;
        }

        if (typeof doExtraModifications !== 'undefined') {
          doExtraModifications(values);
        }

        doSubmit(values);
        setSubmitting(false);
        resetForm();
      }}
    >
      <EntityForm
        entity={entity}
        options={options}
        isCreateMode={isCreateMode}
      />
    </Formik>
  </div>
);
