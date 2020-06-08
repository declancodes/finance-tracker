import moment from 'moment';

const dateDisplayFormat = 'MM/DD/YYYY';

const currencyFormatter = new Intl.NumberFormat(undefined, {
  style: 'currency',
  currency: 'USD',
});

export const helpers = {
  displayDate(dateString) {
    return moment(dateString).local().format(dateDisplayFormat);
  },

  consumeDate(dateString) {
    return moment(dateString).toISOString();
  },

  displayCurrency(currencyString) {
    return currencyFormatter.format(currencyString);
  },

  displayDecimals(num, places) {
    return (Math.round(num * 1000) / 1000).toFixed(places);
  },

  displayPercentage(num, places) {
    return (Math.round(num * 100000) / 1000).toFixed(places);
  },

  getOptionsFromKey(options, name, defaultValue) {
    if (!this.isNonEmptyArray(options)) {
      return defaultValue;
    }

    const opt = options.filter(o => o.name === name);
    return this.isNonEmptyArray(opt) ? opt[0].value : defaultValue;
  },

  getOptionsArrayFromKey(options, name) {
    return this.getOptionsFromKey(options, name, options);
  },

  getValueFromKey(options, name) {
    return this.getOptionsFromKey(options, name, '');
  },

  isNonEmptyArray(obj) {
    return Array.isArray(obj) && obj.length > 0;
  }
};
