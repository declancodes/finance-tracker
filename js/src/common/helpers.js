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
  }
};
