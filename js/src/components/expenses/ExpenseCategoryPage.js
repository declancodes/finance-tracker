import React from "react";
import axios from "axios";
import Category from "../common/Category";
import CreateCategoryForm from "../common/CreateCategoryForm";

const API_URL = "http://localhost:8080/expensecategories"

class ExpenseCategoryPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      expenseCategories: []
    };
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(values) {
    axios.post(API_URL, values)
      .then(response => {
        console.log(response.data);
        return axios.get(API_URL);
      })
      .then(response => {
        this.setState({ expenseCategories: response.data })
      })
  }

  handleDelete(uuid) {
    const url = `${API_URL}/${uuid}`

    axios.delete(url)
      .then(() => axios.get(API_URL))
      .then(response => {
        this.setState({ expenseCategories: response.data })
      })
  }

  componentDidMount() {
    axios.get(API_URL).then(response => response.data)
      .then((data) => {
        this.setState({ expenseCategories: data })
      })
  }

  render() {
    return (
      <div>
        <h1>Expense Categories</h1>
        <div className="expenseCategories">
          {this.state.expenseCategories.map(expenseCategory =>
            (
              (
                <Category
                  key={expenseCategory.expenseCategoryUuid}
                  category={expenseCategory}
                  handleDelete={() => this.handleDelete(expenseCategory.expenseCategoryUuid)}
                />
              )
            )
          )}
        </div>
        <CreateCategoryForm categoryType="Expense" doSubmit={this.handleSubmit}/>
      </div>
    );
  }
}

export default ExpenseCategoryPage;