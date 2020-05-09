import React from 'react';
import { Formik, Field, Form } from 'formik';
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
      .then(response => response.data)
      .then(data => this.setState({ accountCategories: data }))
  }

  render() {
    const a = this.props.account;

    return (
      <div>
        <h2>
          {this.props.isEditMode ? "Edit" : "Create"} Account
        </h2>
        <Formik
          initialValues={{
            uuid: a ? a.uuid : '',
            name: a ? a.name : '',
            category: a ? a.category.uuid : '',
            description: a ? a.description : '',
            amount: a ? a.amount : 0
          }}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            if (!this.props.isEditMode) {
              delete values.uuid;
            }

            var acUuid = values.category;
            values.category = {
              uuid: acUuid
            };

            this.props.doSubmit(values);
            setSubmitting(false);
            resetForm();
          }}
        >
          <Form>
            <label htmlFor="name">Name</label>
            <Field name="name" type="text" />
            <label htmlFor="category">Category</label>
            <Field name="category" as="select">
              <option defaultValue="">Select Category</option>
              {this.state.accountCategories.map(accountCategory => (
                <option
                  key={accountCategory.uuid}
                  value={accountCategory.uuid}
                >
                  {accountCategory.name}
                </option>
              ))}
            </Field>
            <label htmlFor="description">Description</label>
            <Field name="description" type="text" />
            <label htmlFor="amount">Amount</label>
            <Field name="amount" type="number" />
            <button type="submit">
              {this.props.isEditMode ? "Update" : "Create"}
            </button>
          </Form>
        </Formik>
      </div>
    );
  }
}

export default AccountForm;