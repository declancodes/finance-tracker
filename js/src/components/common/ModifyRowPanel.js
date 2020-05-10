import React from "react";
import Button from "../common/Button";

class ModifyRowPanel extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <Button
          name="Edit"
          handleFunc={this.props.handleEdit}
        />
        <Button
          name="Delete"
          handleFunc={this.props.handleDelete}
        />
      </div>
    );
  }
}

export default ModifyRowPanel;
