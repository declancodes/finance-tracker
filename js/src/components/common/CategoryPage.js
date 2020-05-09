import React from "react";
import axios from "axios";
import CreateCategoryForm from "../common/CreateCategoryForm";
import CategoryRow from "./CategoryRow";

class CategoryPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      categories: [],
      isEditing: false
    };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.performUpdate = this.performUpdate.bind(this);
    this.apiUrl = `http://localhost:8080/${this.props.categoryType.toLowerCase()}categories`;
  }

  handleSubmit(values) {
    axios.post(this.apiUrl, values)
      .then(response => {
        console.log(response.data);
        return axios.get(this.apiUrl);
      })
      .then(response => {
        this.setState({ categories: response.data })
      })
  }

  handleDelete(uuid) {
    const url = `${this.apiUrl}/${uuid}`

    axios.delete(url)
      .then(() => axios.get(this.apiUrl))
      .then(response => {
        this.setState({ categories: response.data })
      })
  }

  performUpdate(values) {
    const url = `${this.apiUrl}/${values.uuid}`

    axios.put(url, values)
      .then(response => {
        console.log(response.data);
        return axios.get(this.apiUrl);
      })
      .then(response => {
        this.setState({
          categories: response.data,
          isEditing: false
        })
      })
  }

  componentDidMount() {
    axios.get(this.apiUrl).then(response => response.data)
      .then((data) => {
        this.setState({ categories: data })
      })
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
                    isEditing={this.state.isEditing}
                    categoryType={this.props.categoryType}
                    category={category}
                    performUpdate={this.performUpdate}
                    handleDelete={() => this.handleDelete(category.uuid)}
                    setIsEditing={(editing) => this.setState({ isEditing: editing })}
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
        <CreateCategoryForm
          categoryType={this.props.categoryType}
          doSubmit={this.handleSubmit}
        />
      </div>
    );
  }
}

export default CategoryPage;