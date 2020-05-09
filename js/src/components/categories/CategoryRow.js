import React from "react";
import DeleteButton from "../common/DeleteButton";
import CategoryForm from "./CategoryForm";

class CategoryRow extends React.Component {
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
    const c = this.props.category;

    return (
      <tr>
        <td>{c.name}</td>
        <td>{c.description}</td>
        <td>
          <DeleteButton handleDelete={() => this.props.handleDelete(c.uuid)}/>
          {this.state.isEditing ? (
            <div>
              <CategoryForm
                isEditMode={true}
                categoryType={this.props.categoryType}
                category={c}
                doSubmit={this.handleUpdate}
              />
              <button onClick={() => this.setEditing(false)}>
                Cancel
              </button>
            </div>
          ) : (
            <button
              className="edit-button"
              onClick={() => this.setEditing(true)}>
                Edit
            </button>
          )}
        </td>
      </tr>
    );
  }
}

export default CategoryRow;