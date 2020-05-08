import React from "react";
import DeleteButton from "./DeleteButton";

class Category extends React.Component {
  render() {
    return (
      <div className="category">
        <div className="category-body">
          <h4>{this.props.category.name}</h4>
          <h5>{this.props.category.description}</h5>
        </div>
        <DeleteButton handleDelete={() => this.props.handleDelete()}/>
      </div>
    );
  }
}

export default Category;