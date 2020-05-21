import React from 'react';
import EntityFormik from '../common/forms/EntityFormik';
import api from '../common/api';
import moment from 'moment';

class ContributionForm extends React.Component {
  constructor(props) {
    super(props);
    this.getOptions = this.getOptions.bind(this);
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.doSubmit = this.doSubmit.bind(this);
  }

  getOptions() {
    return api.getAccounts()
      .then(response => {
        return (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.name.localeCompare(b.name));
      });
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
        getOptions={this.getOptions}
        doExtraModifications={this.doExtraModifications}
        doSubmit={this.doSubmit}
      />
    );
  }
}

export default ContributionForm;
