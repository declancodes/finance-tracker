import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api';
import moment from 'moment';

class ContributionsPage extends React.Component {
  constructor(props) {
    super(props);
    this.createContribution = this.createContribution.bind(this);
    this.getContributions = this.getContributions.bind(this);
    this.updateContribution = this.updateContribution.bind(this);
    this.deleteContribution = this.deleteContribution.bind(this);
    this.getOptions = this.getOptions.bind(this);
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.getInitialValues = this.getInitialValues.bind(this);
  }

  createContribution(values) {
    return api.createContribution(values);
  }

  getContributions(start, end) {
    return api.getContributions(start, end)
      .then(response => {
        return (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.date.localeCompare(b.date));
      });
  }

  updateContribution(values) {
    return api.updateContribution(values);
  }

  deleteContribution(uuid) {
    return api.deleteContribution(uuid);
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

  getInitialValues(contribution) {
    let initialValues = JSON.parse(JSON.stringify(contribution));
    initialValues.account = contribution.account.uuid;
    initialValues.date = moment(contribution.date).format('MM/DD/YYYY')

    return initialValues;
  }

  render() {
    return (
      <EntityPage
        entityName='Contribution'
        entityPlural='Contributions'
        columnLength={6}
        blankEntity={{
          uuid: '',
          name: '',
          account: '',
          description: '',
          date: '',
          amount: 0
        }}
        usesDates={true}
        createEntity={this.createContribution}
        getEntities={this.getContributions}
        updateEntity={this.updateContribution}
        deleteEntity={this.deleteContribution}
        getOptions={this.getOptions}
        doExtraModifications={this.doExtraModifications}
        getInitialValues={this.getInitialValues}
      />
    );
  }
}

export default ContributionsPage;
