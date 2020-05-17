import React from "react";
import moment from "moment";
import api from "../common/api"
import ContributionForm from "./ContributionForm";
import ContributionRow from "./ContributionRow";
import DateRangePanel from "../common/DateRangePanel";

class ContributionsPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      contributions: [],
      start: moment().startOf("month").toDate(),
      end: moment().endOf("month").toDate()
    };
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
    this.handleStartDateSet = this.handleStartDateSet.bind(this);
    this.handleEndDateSet = this.handleEndDateSet.bind(this);
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
        var contributions = (response.data === null || response.data === undefined)
          ? []
          : response.data
            .sort((a, b) => a.date.localeCompare(b.date));
        this.setState({ contributions: contributions });
      });
  }

  render() {
    return (
      <div>
        <h1>Contributions</h1>
        <DateRangePanel
          start={this.state.start}
          end={this.state.end}
          setStart={this.handleStartDateSet}
          setEnd={this.handleEndDateSet}
        />
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Account</th>
              <th>Description</th>
              <th>Date</th>
              <th>Amount</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {this.state.contributions.length > 0 ? (
              this.state.contributions.map(contribution => (
                (
                  <ContributionRow
                    key={contribution.uuid}
                    contribution={contribution}
                    handleUpdate={this.handleUpdate}
                    handleDelete={this.handleDelete}
                  />
                )
              ))
            ) : (
              <tr>
                <td colSpan={6}>No Contributions</td>
              </tr>
            )}
          </tbody>
        </table>
        <ContributionForm
          isEditMode={false}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default ContributionsPage;