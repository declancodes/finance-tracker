import React from 'react';
import EntityPage from '../common/EntityPage';
import api from '../../api';

class CategoryPage extends React.Component {
  constructor(props) {
    super(props);
    this.createCategory = this.createCategory.bind(this);
    this.getCategories = this.getCategories.bind(this);
    this.updateCategory = this.updateCategory.bind(this);
    this.deleteCategory = this.deleteCategory.bind(this);
  }

  createCategory(values) {
    return this.isAccountCategory()
      ? api.createAccountCategory(values)
      : api.createExpenseCategory(values);
  }

  getCategories() {
    return (this.isAccountCategory()
      ? api.getAccountCategories()
      : api.getExpenseCategories()
    ).then(response => {
      return (response.data === null || response.data === undefined)
        ? []
        : response.data.sort((a, b) => a.name.localeCompare(b.name));
    });
  }

  updateCategory(values) {
    return this.isAccountCategory()
      ? api.updateAccountCategory(values)
      : api.updateExpenseCategory(values);
  }

  deleteCategory(uuid) {
    return this.isAccountCategory()
      ? api.deleteAccountCategory(uuid)
      : api.deleteExpenseCategory(uuid);
  }

  isAccountCategory() {
    return this.props.categoryType === 'Account';
  }

  render() {
    return (
      <EntityPage
        entityName={`${this.props.categoryType} Category`}
        entityPlural={`${this.props.categoryType} Categories`}
        columnLength={3}
        blankEntity={{
          uuid: '',
          name: '',
          description: ''
        }}
        usesDates={false}
        createEntity={this.createCategory}
        getEntities={this.getCategories}
        updateEntity={this.updateCategory}
        deleteEntity={this.deleteCategory}
      />
    );
  }
}

export default CategoryPage;
