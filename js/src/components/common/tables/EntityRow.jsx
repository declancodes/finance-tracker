import React from 'react';
import { Button } from '../Button';
import EntityForm from '../forms/EntityForm';
import { ModifyRowPanel } from './ModifyRowPanel';
import moment from 'moment';

class EntityRow extends React.Component {
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
    const e = this.props.entity;

    return (
      <tr>
        {e.hasOwnProperty('name') && <td>{e.name}</td>}
        {e.hasOwnProperty('account') && <td>{e.account.name}</td>}
        {e.hasOwnProperty('category') && <td>{e.category.name}</td>}
        {e.hasOwnProperty('description') && <td>{e.description}</td>}
        {e.hasOwnProperty('date') && <td>{moment(e.date).format('MM/DD/YYYY')}</td>}
        {e.hasOwnProperty('amount') && <td>${e.amount}</td>}
        <td>
          {this.state.isEditing ? (
            <div>
              <EntityForm
                entityName={this.props.entityName}
                entity={e}
                getInitialValues={this.props.getInitialValues}
                isCreateMode={false}
                getOptions={this.props.getOptions}
                doExtraModifications={this.props.doExtraModifications}
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
              handleDelete={() => this.props.handleDelete(e.uuid)}
            />
          )}
        </td>
      </tr>
    );
  }
}

export default EntityRow;
