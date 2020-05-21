import React from 'react';
import { EmptyEntityRow } from '../common/tables/EmptyEntityRow';
import EntityForm from '../common/forms/EntityForm';
import { EntityHeader } from '../common/tables/EntityHeader';
import EntityRow from '../common/tables/EntityRow';
import api from '../../api';

class CategoryPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      categories: []
    };
    this.handleCreate = this.handleCreate.bind(this);
    this.handleUpdate = this.handleUpdate.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
  }

  handleCreate(values) {
    const isAccountCategory = this.isAccountCategory()
    const p = isAccountCategory
      ? api.createAccountCategory(values)
      : api.createExpenseCategory(values)

    p.then(() => this.setCategories(isAccountCategory))
  }

  handleDelete(uuid) {
    const isAccountCategory = this.isAccountCategory()
    const p = isAccountCategory
      ? api.deleteAccountCategory(uuid)
      : api.deleteExpenseCategory(uuid)

    p.then(() => this.setCategories(isAccountCategory))
  }

  handleUpdate(values) {
    const isAccountCategory = this.isAccountCategory()
    const p = isAccountCategory
      ? api.updateAccountCategory(values)
      : api.updateExpenseCategory(values)

    p.then(() => this.setCategories(isAccountCategory))
  }

  componentDidMount() {
    this.setCategories(this.isAccountCategory())
  }

  setCategories(isAccountCategory) {
    this.getCategories(isAccountCategory)
      .then(response => {
        let categories = (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.name.localeCompare(b.name))
        this.setState({ categories: categories })
      })
  }

  getCategories(isAccountCategory) {
    return isAccountCategory
        ? api.getAccountCategories()
        : api.getExpenseCategories()
  }

  isAccountCategory() {
    return this.props.categoryType === 'Account'
  }

  render() {
    const entityName = `${this.props.categoryType} Category`;
    const entityPlural = `${this.props.categoryType} Categories`;
    const blankEntity = {
      uuid: '',
      name: '',
      description: ''
    };

    return (
      <div>
        <h1>{entityPlural}</h1>
        <table>
          <EntityHeader entity={blankEntity}/>
          <tbody>
            {this.state.categories.length > 0 ? (
              this.state.categories.map(category => (
                (
                  <EntityRow
                    key={category.uuid}
                    entityName={entityName}
                    entity={category}
                    handleUpdate={this.handleUpdate}
                    handleDelete={this.handleDelete}
                  />
                )
              ))
            ) : (
              <EmptyEntityRow columnLength={3} entityPlural={entityPlural}/>
            )}
          </tbody>
        </table>
        <EntityForm
          entityName={entityName}
          entity={blankEntity}
          isCreateMode={true}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default CategoryPage;