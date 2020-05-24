import moment from 'moment';

export const helpers = {
  DateDisplayFormat: 'MM/DD/YYYY',

  displayDate(dateString) {
    return moment(dateString).local().format(this.DateDisplayFormat);
  },

  consumeDate(dateString) {
    return moment(dateString).toISOString();
  }
};