import React from 'react';
import { Formik } from 'formik';
import { EntityForm } from '../common/forms/EntityForm';
import api from "../common/api";

class AccountForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accountCategories: []
    };
  }

  componentDidMount() {
    api.getAccountCategories()
      .then(response => {
        let accountCategories = (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.name.localeCompare(b.name));
        this.setState({ accountCategories: accountCategories });
      })
  }

  render() {
    const a = this.props.account;
    const initialAccountValues = {
      uuid: a ? a.uuid : '',
      name: a ? a.name : '',
      category: a ? a.category.uuid : '',
      description: a ? a.description : '',
      amount: a ? a.amount : 0
    };

    return (
      <div>
        <h2>
          {this.props.isEditMode ? "Edit" : "Create"} Account
        </h2>
        <Formik
          initialValues={initialAccountValues}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            if (!this.props.isEditMode) {
              delete values.uuid;
            }

            let acUuid = values.category;
            values.category = {
              uuid: acUuid
            };

            this.props.doSubmit(values);
            setSubmitting(false);
            resetForm();
          }}
        >
          <EntityForm
            entity={initialAccountValues}
            options={this.state.accountCategories}
            isEditMode={this.props.isEditMode}
          />
        </Formik>
      </div>
    );
  }
}

export default AccountForm;