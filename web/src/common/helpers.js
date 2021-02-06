import { format, formatISO, parseISO } from 'date-fns';

const dateDisplayFormat = 'MM/dd/yyyy';

const currencyFormatter = new Intl.NumberFormat(undefined, {
  style: 'currency',
  currency: 'USD',
});

export const displayDate = (dateString) => {
  return format(parseISO(dateString), dateDisplayFormat);
}

export const consumeDate = (dateString) => {
  return formatISO(dateString);
}

export const displayCurrency = (currencyString) => {
  return currencyFormatter.format(currencyString);
}

export const displayDecimals = (num, places) => {
  return (Math.round(num * 1000) / 1000).toFixed(places);
}

export const displayPercentage = (num, places) => {
  const percentageStr = (Math.round(num * 100000) / 1000).toFixed(places);
  return parseFloat(percentageStr);
}

export const titleCase = (str) => {
  return str.split(' ')
    .map(w => w[0].toUpperCase() + w.substr(1).toLowerCase())
    .join(' ');
}

export const getOptionsArrayFromKey = (options, name) => {
  return getOptionsFromKey(options, name, options);
}

export const getValueFromKey = (options, name) => {
  return getOptionsFromKey(options, name, []);
}

const getOptionsFromKey = (options, name, defaultValue) => {
  if (!isNonEmptyArray(options)) {
    return defaultValue;
  }

  const opt = options.find(o => o.name === name);
  return opt && isNonEmptyArray(opt.value)
    ? opt.value
    : defaultValue;
}

const isNonEmptyArray = (obj) => {
  return Array.isArray(obj) && obj.length > 0;
}
