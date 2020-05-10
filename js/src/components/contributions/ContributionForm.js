import React from 'react';
import { Formik, Field, Form } from 'formik';
import api from "../common/api";

class ContributionForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accounts: []
    };
  }

  componentDidMount() {
    api.getAccounts()
      .then(response => response.data)
      .then(data => this.setState({ accounts: data }))
  }

  render() {
    const c = this.props.contribution;

    return (
      <div>
        <h2>
          {this.props.isEditMode ? "Edit" : "Create"} Contribution
        </h2>
        <Formik
          initialValues={{
            uuid: c ? c.uuid : '',
            name: c ? c.name : '',
            account: c ? c.account.uuid : '',
            description: c ? c.description : '',
            date: c ? c.date : '',
            amount: c ? c.amount : 0
          }}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            if (!this.props.isEditMode) {
              delete values.uuid;
            }

            var aUuid = values.account;
            values.account = {
              uuid: aUuid
            };

            this.props.doSubmit(values);
            setSubmitting(false);
            resetForm();
          }}
        >
          <Form>
            <label htmlFor="name">Name</label>
            <Field name="name" type="text" />
            <label htmlFor="account">Account</label>
            <Field name="account" as="select">
              <option defaultValue="">Select Account</option>
              {this.state.accounts.map(account => (
                <option
                  key={account.uuid}
                  value={account.uuid}
                >
                  {account.name}
                </option>
              ))}
            </Field>
            <label htmlFor="description">Description</label>
            <Field name="description" type="text" />
            <label htmlFor="date">Date</label>
            <Field name="date" type="text" />
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

export default ContributionForm;