import React from "react";
import "./DeleteButton.css"

class DeleteButton extends React.Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this)
  }

  handleClick() {
    this.props.handleDelete();
  }

  render() {
    return (
      <button
        className="delete-button"
        onClick={this.handleClick}>
          Delete
      </button>
    );
  }
}

export default DeleteButton;