import React from "react";

class DeleteButton extends React.Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this)
  }

  handleClick() {
    this.props.doClick();
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