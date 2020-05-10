import React from "react";
import Button from "../common/Button";
import ContributionForm from "./ContributionForm";
import ModifyRowPanel from "../common/ModifyRowPanel";

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
            <ModifyRowPanel
              handleEdit={() => this.setEditing(true)}
              handleDelete={() => this.props.handleDelete(c.uuid)}
            />
          )}
        </td>
      </tr>
    );
  }
}

export default ContributionRow;
