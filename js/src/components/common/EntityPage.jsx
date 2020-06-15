import React from 'react';
import { Button, Table } from 'react-bootstrap';
import moment from 'moment';
import pluralize from 'pluralize';
import { DisplayForm } from './forms/DisplayForm';
import { EntityHeader } from './tables/EntityHeader';
import { EntityRow } from './tables/EntityRow';
import { FilterPanel } from './filters/FilterPanel';
import { helpers } from '../../common/helpers';

class EntityPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      entities: [],
      options: [],
      start: moment().startOf('month').toDate(),
      end: moment().endOf('month').toDate(),
      filterCategories: [],
      isCreating: false
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

  handleFilterCategorySet(name, value) {
    const filterCategories = this.state.filterCategories.map(fc => {
      if (fc.name === name) {
        fc.value = value;
      }
      return fc;
    });

    this.handleFilterFieldSet({ filterCategories: filterCategories });
  }

  handlePromise(promise) {
    promise.then(() => this.setEntities());
  }

  handleFilterFieldSet(filterField) {
    this.setState(
      filterField,
      () => this.setEntities());
  }

  getFilterCategoryValue(name) {
    return helpers.getValueFromKey(this.state.filterCategories, name);
  }

  componentDidMount() {
    this.setEntities();
    this.setOptions();
    this.setState({ filterCategories: this.props.filterCategories });
  }

  setIsCreating(val) {
    this.setState({ isCreating: val });
  }

  setOptions() {
    if (this.props.getOptions === undefined) {
      return;
    }

    const opts = this.props.getOptions.map(getOpts => 
      getOpts.value().then(response => {
        return {
          name: getOpts.name,
          value: response
        }
      })
    );

    Promise.all(opts)
      .then(opt => this.setState({ options: opt}))
  }

  setEntities() {
    this.props.getEntities({
      start: this.props.usesDates ? this.state.start.toISOString() : null,
      end: this.props.usesDates ? this.state.end.toISOString() : null,
      category: this.getFilterCategoryValue('category'),
      account: this.getFilterCategoryValue('account'),
      fund: this.getFilterCategoryValue('fund')
    })
    .then(response => this.setState({ entities: response }));
  }

  render() {
    const entityPluralName = pluralize(this.props.entityName);
    return (
      <div>
        <h1>{entityPluralName}</h1>
        {this.props.usesFilters &&
          <FilterPanel
            usesDates={this.props.usesDates}
            start={this.state.start}
            end={this.state.end}
            filterCategories={this.state.filterCategories}
            options={this.state.options}
            setStart={this.handleStartDateSet}
            setEnd={this.handleEndDateSet}
            setFilterCategory={this.handleFilterCategorySet}
          />
        }
        {this.props.children}
        <Table
          variant='dark'
          striped
          bordered
          hover
        >
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
                <td colSpan={Object.keys(this.props.blankEntity).length}>
                  No {entityPluralName}
                </td>
              </tr>
            )}
          </tbody>
        </Table>
        {this.state.isCreating ? (
          <DisplayForm
            entityName={this.props.entityName}
            entity={this.props.blankEntity}
            isCreateMode={true}
            options={this.state.options}
            doExtraModifications={this.props.doExtraModifications}
            doSubmit={this.handleCreate}
            setNotUsing={() => this.setIsCreating(false)}
          />
        ) : (
          <Button
            variant='dark'
            onClick={() => this.setIsCreating(true)}
          >
            {`Create ${this.props.entityName}`}
          </Button>
        )}
      </div>
    );
  }
}

export default EntityPage;
