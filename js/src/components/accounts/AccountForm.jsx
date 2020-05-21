import React from 'react';
import EntityFormik from '../common/forms/EntityFormik';
import api from '../common/api';

class AccountForm extends React.Component {
  constructor(props) {
    super(props);
    this.getOptions = this.getOptions.bind(this);
    this.doExtraModifications = this.doExtraModifications.bind(this);
    this.doSubmit = this.doSubmit.bind(this);
  }

  getOptions() {
    return api.getAccountCategories()
      .then(response => {
        return (response.data === null || response.data === undefined)
          ? []
          : response.data.sort((a, b) => a.name.localeCompare(b.name));
      });
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

  render() {
    const entity = {
      uuid: this.props.isCreateMode ? '' : this.props.account.uuid,
      name: this.props.isCreateMode ? '' : this.props.account.name,
      category: this.props.isCreateMode ? '' : this.props.account.category.uuid,
      description: this.props.isCreateMode ? '' : this.props.account.description,
      amount: this.props.isCreateMode ? 0 : this.props.account.amount
    };

    return (
      <EntityFormik
        entityName='Account'
        entity={entity}
        isCreateMode={this.props.isCreateMode}
        getOptions={this.getOptions}
        doExtraModifications={this.doExtraModifications}
        doSubmit={this.doSubmit}
      />
    );
  }
}

export default AccountForm;
