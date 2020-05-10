import React from "react";
import "./Button.css";

class Button extends React.Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this)
  }

  handleClick() {
    this.props.handleFunc();
  }

  render() {
    return (
      <button
        className={`${this.props.name.toLowerCase()}-button`}
        onClick={this.handleClick}>
          {this.props.name}
      </button>
    );
  }
}

export default Button;