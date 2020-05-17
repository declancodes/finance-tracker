import React from "react";
import api from "../common/api";
import CategoryForm from "./CategoryForm";
import CategoryRow from "./CategoryRow";

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
        var categories = (response.data === null || response.data === undefined)
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
    return this.props.categoryType === "Account"
  }

  render() {
    return (
      <div>
        <h1>{this.props.categoryType} Categories</h1>
        <table>
          <thead>
            <tr>
              <th>Name</th>
              <th>Description</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {this.state.categories.length > 0 ? (
              this.state.categories.map(category => (
                (
                  <CategoryRow
                    key={category.uuid}
                    categoryType={this.props.categoryType}
                    category={category}
                    handleUpdate={this.handleUpdate}
                    handleDelete={this.handleDelete}
                  />
                )
              ))
            ) : (
              <tr>
                <td colSpan={3}>No {this.props.categoryType} Categories</td>
              </tr>
            )}
          </tbody>
        </table>
        <CategoryForm
          isEditMode={false}
          categoryType={this.props.categoryType}
          doSubmit={this.handleCreate}
        />
      </div>
    );
  }
}

export default CategoryPage;