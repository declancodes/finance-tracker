import React from 'react';
import { DateRangePanel } from './DateRangePanel';
import { EmptyEntityRow } from './tables/EmptyEntityRow';
import EntityForm from './forms/EntityForm';
import { EntityHeader } from './tables/EntityHeader';
import EntityRow from './tables/EntityRow';
import moment from 'moment';

class EntityPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      entities: [],
      start: moment().startOf('month').toDate(),
      end: moment().endOf('month').toDate()
    };
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
    this.handleStartDateSet = this.handleStartDateSet.bind(this);
    this.handleEndDateSet = this.handleEndDateSet.bind(this);
  }

  handleCreate(values) {
    this.props.createEntity(values)
      .then(() => this.setEntities());
  }

  handleUpdate(values) {
    this.props.updateEntity(values)
      .then(() => this.setEntities());
  }

  handleDelete(uuid) {
    this.props.deleteEntity(uuid)
      .then(() => this.setEntities());
  }

  handleStartDateSet(value) {
    this.setState(
      { start: value },
      () => this.setEntities());
  }

  handleEndDateSet(value) {
    this.setState(
      { end: value },
      () => this.setEntities());
  }

  componentDidMount() {
    this.setEntities();
  }

  setEntities() {
    (this.props.usesDates
      ? this.props.getEntities(this.state.start.toISOString(), this.state.end.toISOString())
      : this.props.getEntities()
    ).then(response => this.setState({ entities: response }));
  }

  render() {
    return (
      <div>
        <h1>{this.props.entityPlural}</h1>
        {this.props.usesDates &&
          <DateRangePanel
            start={this.state.start}
            end={this.state.end}
            setStart={this.handleStartDateSet}
            setEnd={this.handleEndDateSet}
          />
        }
        <table>
          <EntityHeader entity={this.props.blankEntity}/>
          <tbody>
            {this.state.entities.length > 0 ? (
              this.state.entities.map(entity => (
                <EntityRow
                  key={entity.uuid}
                  entityName={this.props.entityName}
                  entity={entity}
                  getInitialValues={this.props.getInitialValues}
                  getOptions={this.props.getOptions}
                  doExtraModifications={this.props.doExtraModifications}
                  handleUpdate={this.handleUpdate}
                  handleDelete={this.handleDelete}
                />
              ))
            ) : (
              <EmptyEntityRow
                columnLength={this.props.columnLength}
                entityPlural={this.props.entityPlural}
              />
            )}
          </tbody>
        </table>
        <EntityForm
          entityName={this.props.entityName}
          entity={this.props.blankEntity}
          isCreateMode={true}
          getOptions={this.props.getOptions}
          doExtraModifications={this.props.doExtraModifications}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default EntityPage;
