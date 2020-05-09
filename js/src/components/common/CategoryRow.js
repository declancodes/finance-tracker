import React from "react";
import DeleteButton from "./DeleteButton";
import EditCategoryForm from "./EditCategoryForm";

class CategoryRow extends React.Component {
  render() {
    return (
      <tr>
        <td>{this.props.category.name}</td>
        <td>{this.props.category.description}</td>
        <td>
          <DeleteButton handleDelete={() => this.props.handleDelete()}/>
          {this.props.isEditing ? (
            <div>
              <EditCategoryForm
                categoryType={this.props.categoryType}
                category={this.props.category}
                doUpdate={this.props.performUpdate}
              />
              <button onClick={() => this.props.setIsEditing(false)}>
                Cancel
              </button>
            </div>
          ) : (
            <button
              className="edit-button"
              onClick={() => this.props.setIsEditing(true)}>
                Edit
            </button>
          )}
        </td>
      </tr>
    );
  }
}

export default CategoryRow;