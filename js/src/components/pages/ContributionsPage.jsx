import React from 'react';
import { DateRangePanel } from '../common/DateRangePanel';
import { EmptyEntityRow } from '../common/tables/EmptyEntityRow';
import EntityFormik from '../common/forms/EntityFormik';
import { EntityHeader } from '../common/tables/EntityHeader';
import EntityRow from '../common/tables/EntityRow';
import api from '../../api';
import moment from 'moment';

class ContributionsPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      contributions: [],
      start: moment().startOf('month').toDate(),
      end: moment().endOf('month').toDate()
    };
    this.getOptions = this.getOptions.bind(this);
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
    this.handleStartDateSet = this.handleStartDateSet.bind(this);
    this.handleEndDateSet = this.handleEndDateSet.bind(this);
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

  handleCreate(values) {
    api.createContribution(values)
      .then(() => this.setContributions())
  }

  handleDelete(uuid) {
    api.deleteContribution(uuid)
      .then(() => this.setContributions())
  }

  handleUpdate(values) {
    api.updateContribution(values)
      .then(() => this.setContributions())
  }

  handleStartDateSet(value) {
    this.setState(
      { start: value },
      () => this.setContributions());
  }

  handleEndDateSet(value) {
    this.setState(
      { end: value },
      () => this.setContributions());
  }

  componentDidMount() {
    this.setContributions()
  }

  setContributions() {
    api.getContributions(this.state.start.toISOString(), this.state.end.toISOString())
      .then(response => {
        let contributions = (response.data === null || response.data === undefined)
          ? []
          : response.data
            .sort((a, b) => a.date.localeCompare(b.date));
        this.setState({ contributions: contributions });
      });
  }

  render() {
    const entityName = 'Contribution';
    const entityPlural = `${entityName}s`;
    const blankEntity = {
      uuid: '',
      name: '',
      account: '',
      description: '',
      date: '',
      amount: 0
    };

    return (
      <div>
        <h1>{entityPlural}</h1>
        <DateRangePanel
          start={this.state.start}
          end={this.state.end}
          setStart={this.handleStartDateSet}
          setEnd={this.handleEndDateSet}
        />
        <table>
          <EntityHeader entity={blankEntity}/>
          <tbody>
            {this.state.contributions.length > 0 ? (
              this.state.contributions.map(contribution => {
                let initialVals = JSON.parse(JSON.stringify(contribution));
                initialVals.account = contribution.account.uuid;
                initialVals.date = moment(contribution.date).format('MM/DD/YYYY')

                return (
                  <EntityRow
                    key={contribution.uuid}
                    entityName={entityName}
                    entity={contribution}
                    initialValues={initialVals}
                    getOptions={this.getOptions}
                    doExtraModifications={this.doExtraModifications}
                    handleUpdate={this.handleUpdate}
                    handleDelete={this.handleDelete}
                  />
                );
              })
            ) : (
              <EmptyEntityRow columnLength={6} entityPlural={entityPlural}/>
            )}
          </tbody>
        </table>
        <EntityFormik
          entityName={entityName}
          entity={blankEntity}
          isCreateMode={true}
          getOptions={this.getOptions}
          doExtraModifications={this.doExtraModifications}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default ContributionsPage;