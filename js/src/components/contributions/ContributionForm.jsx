import React from 'react';
import { Formik } from 'formik';
import { EntityForm } from '../common/forms/EntityForm'
import moment from 'moment';
import api from '../common/api';

class ContributionForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accounts: []
    };
  }

  componentDidMount() {
    api.getAccounts()
      .then(response => {
        let accounts = (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.name.localeCompare(b.name))
        this.setState({ accounts: accounts })
      })
  }

  render() {
    const c = this.props.contribution;
    const initialContributionValues = {
      uuid: c ? c.uuid : '',
      name: c ? c.name : '',
      account: c ? c.account.uuid : '',
      description: c ? c.description : '',
      date: c ? moment(c.date).format('MM/DD/YYYY') : '',
      amount: c ? c.amount : 0
    };

    return (
      <div>
        <h2>
          {this.props.isEditMode ? 'Edit' : 'Create'} Contribution
        </h2>
        <Formik
          initialValues={initialContributionValues}
          onSubmit={(values, { setSubmitting, resetForm }) => {
            if (!this.props.isEditMode) {
              delete values.uuid;
            }

            let aUuid = values.account;
            values.account = {
              uuid: aUuid
            };

            let dateToSubmit = moment(values.date).toISOString();
            values.date = dateToSubmit;

            this.props.doSubmit(values);
            setSubmitting(false);
            resetForm();
          }}
        >
          <EntityForm
            entity={initialContributionValues}
            options={this.state.accounts}
            isEditMode={this.props.isEditMode}
          />
        </Formik>
      </div>
    );
  }
}

export default ContributionForm;