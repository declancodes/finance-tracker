import React from "react";
import Button from "../common/Button";
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
          <Button
            name="Delete"
            handleFunc={() => this.props.handleDelete(c.uuid)}
          />
          {this.state.isEditing ? (
            <div>
              <CategoryForm
                isEditMode={true}
                categoryType={this.props.categoryType}
                category={c}
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

export default CategoryRow;