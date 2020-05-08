import React from "react";
import axios from "axios";
import Category from "../common/Category";
import CreateCategoryForm from "../common/CreateCategoryForm";

class CategoryPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      categories: []
    };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.apiUrl = `http://localhost:8080/${this.props.categoryType.toLowerCase()}categories`
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
        <div className="categories">
          {this.state.categories.map(category =>
            (
              (
                <Category
                  key={category.uuid}
                  category={category}
                  handleDelete={() => this.handleDelete(category.uuid)}
                />
              )
            )
          )}
        </div>
        <CreateCategoryForm
          categoryType={this.props.categoryType}
          doSubmit={this.handleSubmit}
        />
      </div>
    );
  }
}

export default CategoryPage;