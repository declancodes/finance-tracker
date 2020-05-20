import React from 'react';
import { Button } from '../common/Button';
import { CategoryForm } from './CategoryForm';
import { ModifyRowPanel } from '../common/ModifyRowPanel';

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
          {this.state.isEditing ? (
            <div>
              <CategoryForm
                category={c}
                categoryType={this.props.categoryType}
                doSubmit={this.handleUpdate}
              />
              <Button
                name='Cancel'
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

export default CategoryRow;