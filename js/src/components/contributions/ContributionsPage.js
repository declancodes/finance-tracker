import React from "react";
import api from "../common/api"
import ContributionForm from "./ContributionForm";
import ContributionRow from "./ContributionRow";

class ContributionsPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      contributions: []
    };
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
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

  componentDidMount() {
    this.setContributions()
  }

  setContributions() {
    api.getContributions()
      .then(response => {
        var contributions = (response.data === null || response.data === undefined)
          ? []
          : response.data
        this.setState({ contributions: contributions })
      })
  }

  render() {
    return (
      <div>
        <h1>Contributions</h1>
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