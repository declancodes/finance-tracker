import React from 'react';
import { EntityForm } from './forms/EntityForm';
import { EntityHeader } from './tables/EntityHeader';
import { EntityRow } from './tables/EntityRow';
import { FilterPanel } from './filters/FilterPanel';
import moment from 'moment';
import size from 'lodash.size';

class EntityPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      entities: [],
      options: [],
      start: moment().startOf('month').toDate(),
      end: moment().endOf('month').toDate(),
      filterCategory: ''
    };
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
    this.handleStartDateSet = this.handleStartDateSet.bind(this);
    this.handleEndDateSet = this.handleEndDateSet.bind(this);
    this.handleFilterCategorySet = this.handleFilterCategorySet.bind(this);
  }

  handleCreate(values) {
    this.handlePromise(this.props.createEntity(values));
  }

  handleUpdate(values) {
    this.handlePromise(this.props.updateEntity(values));
  }

  handleDelete(uuid) {
    this.handlePromise(this.props.deleteEntity(uuid));
  }

  handleStartDateSet(value) {
    this.handleFilterFieldSet({ start: value });
  }

  handleEndDateSet(value) {
    this.handleFilterFieldSet({ end: value });
  }

  handleFilterCategorySet(value) {
    this.handleFilterFieldSet({ filterCategory: value });
  }

  handlePromise(promise) {
    promise.then(() => this.setEntities());
  }

  handleFilterFieldSet(filterField) {
    this.setState(
      filterField,
      () => this.setEntities());
  }

  componentDidMount() {
    this.setEntities();
    this.setOptions();
  }

  setOptions() {
    if (this.props.getOptions !== undefined) {
      this.props.getOptions()
        .then(response => this.setState({ options: response }));
    }
  }

  setEntities() {
    (this.props.usesFilters
      ? this.props.getEntities(
          this.state.start.toISOString(),
          this.state.end.toISOString(),
          this.state.filterCategory
        )
      : this.props.getEntities()
    ).then(response => this.setState({ entities: response }));
  }

  render() {
    return (
      <div>
        <h1>{this.props.entityPlural}</h1>
        {this.props.usesFilters &&
          <FilterPanel
            start={this.state.start}
            end={this.state.end}
            filterCategory={this.state.filterCategory}
            filterCategoryOptions={this.state.options}
            filterCategoryName={this.props.filterCategoryName}
            setStart={this.handleStartDateSet}
            setEnd={this.handleEndDateSet}
            setFilterCategory={this.handleFilterCategorySet}
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
                  options={this.state.options}
                  doExtraModifications={this.props.doExtraModifications}
                  handleUpdate={this.handleUpdate}
                  handleDelete={this.handleDelete}
                />
              ))
            ) : (
              <tr>
                <td colSpan={size(this.props.blankEntity)}>
                  No {this.props.entityPlural}
                </td>
              </tr>
            )}
          </tbody>
        </table>
        <EntityForm
          entityName={this.props.entityName}
          entity={this.props.blankEntity}
          isCreateMode={true}
          options={this.state.options}
          doExtraModifications={this.props.doExtraModifications}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default EntityPage;
