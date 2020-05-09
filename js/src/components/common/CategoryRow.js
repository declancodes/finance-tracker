import React from "react";
import DeleteButton from "./DeleteButton";
import EditCategoryForm from "./EditCategoryForm";

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
    this.setState({ isEditing: false });
  }

  render() {
    return (
      <tr>
        <td>{this.props.category.name}</td>
        <td>{this.props.category.description}</td>
        <td>
          <DeleteButton handleDelete={() => this.props.handleDelete()}/>
          {this.state.isEditing ? (
            <div>
              <EditCategoryForm
                categoryType={this.props.categoryType}
                category={this.props.category}
                doUpdate={this.handleUpdate}
              />
              <button onClick={() => this.setState({ isEditing: false })}>
                Cancel
              </button>
            </div>
          ) : (
            <button
              className="edit-button"
              onClick={() => this.setState({ isEditing: true })}>
                Edit
            </button>
          )}
        </td>
      </tr>
    );
  }
}

export default CategoryRow;