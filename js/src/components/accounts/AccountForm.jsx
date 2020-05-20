import React from 'react';
import { EntityFormik } from '../common/forms/EntityFormik';
import api from '../common/api';

class AccountForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      accountCategories: []
    };
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.doSubmit = this.doSubmit.bind(this);
  }

  doExtraModifications(values) {
    const acUuid = values.category;
    values.category = {
      uuid: acUuid
    };
  }

  doSubmit(values) {
    this.props.doSubmit(values);
  }

  componentDidMount() {
    api.getAccountCategories()
      .then(response => {
        let accountCategories = (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.name.localeCompare(b.name));
        this.setState({ accountCategories: accountCategories });
      })
  }

  render() {
    const isCreating = this.props.isCreateMode;
    const a = this.props.account;
    const entity = {
      uuid: isCreating ? '' : a.uuid,
      name: isCreating ? '' : a.name,
      category: isCreating ? '' : a.category.uuid,
      description: isCreating ? '' : a.description,
      amount: isCreating ? 0 : a.amount
    };

    return (
      <EntityFormik
        entityName='Account'
        entity={entity}
        isCreateMode={isCreating}
        options={this.state.accountCategories}
        doExtraModifications={this.doExtraModifications}
        doSubmit={this.doSubmit}
      />
    );
  }
}

export default AccountForm;