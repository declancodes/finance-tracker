import React from "react";
import Button from "../common/Button";
import ContributionForm from "./ContributionForm";

class ContributionRow extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isEditing: false
    };
    this.handleUpdate = this.handleUpdate.bind(this);
  }

  handleUpdate(values) {
    this.props.handleUpdate(values);
    this.setEditing(false);
  }

  setEditing(val) {
    this.setState({ isEditing: val })
  }

  render() {
    const c = this.props.contribution;

    return (
      <tr>
        <td>{c.name}</td>
        <td>{c.account.name}</td>
        <td>{c.description}</td>
        <td>{c.date}</td>
        <td>${c.amount}</td>
        <td>
          <Button
            name="Delete"
            handleFunc={() => this.props.handleDelete(c.uuid)}
          />
          {this.state.isEditing ? (
            <div>
              <ContributionForm
                isEditMode={true}
                contribution={c}
                doSubmit={this.handleUpdate}
              />
              <Button
                name="Cancel"
                handleFunc={() => this.setEditing(false)}
              />
            </div>
          ) : (
            <Button
              name="Edit"
              handleFunc={() => this.setEditing(true)}
            />
          )}
        </td>
      </tr>
    );
  }
}

export default ContributionRow;
