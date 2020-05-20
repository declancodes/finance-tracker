import React from 'react';
import { EntityFormik } from '../common/forms/EntityFormik';
import api from '../common/api';
import moment from 'moment';

class ContributionForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accounts: []
    };
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.doSubmit = this.doSubmit.bind(this);
  }

  doExtraModifications(values) {
    const aUuid = values.account;
    values.account = {
      uuid: aUuid
    };

    const dateToSubmit = moment(values.date).toISOString();
    values.date = dateToSubmit;
  }

  doSubmit(values) {
    this.props.doSubmit(values);
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
    const isCreating = this.props.isCreateMode;
    const c = this.props.contribution;
    const entity = {
      uuid: isCreating ? '' : c.uuid,
      name: isCreating ? '' : c.name,
      account: isCreating ? '' : c.account.uuid,
      description: isCreating ? '' : c.description,
      date: isCreating ? '' : moment(c.date).format('MM/DD/YYYY'),
      amount: isCreating ? 0 : c.amount
    };

    return (
      <EntityFormik
        entityName='Contribution'
        entity={entity}
        isCreateMode={isCreating}
        options={this.state.accounts}
        doExtraModifications={this.doExtraModifications}
        doSubmit={this.doSubmit}
      />
    );
  }
}

export default ContributionForm;
