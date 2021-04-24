import { format, formatISO, isValid, parse, parseISO } from "date-fns";

const dateDisplayFormat: string = "MM/dd/yyyy";

const currencyFormatter: Intl.NumberFormat = new Intl.NumberFormat(undefined, {
  style: "currency",
  currency: "USD",
});

export const displayDate = (date: string): string => {
  return format(parseISO(date), dateDisplayFormat);
}

export const consumeDate = (date: any): string => {
  const parsedDate = isValid(date)
    ? date
    : parse(date, dateDisplayFormat, new Date());

  return formatISO(parsedDate);
}

export const displayCurrency = (currency: number): string => {
  return currencyFormatter.format(currency);
}

export const displayDecimals = (num: number, places: number): string => {
  return (Math.round(num * 1000) / 1000).toFixed(places);
}

export const displayPercentage = (num: number, places: number): number => {
  const percentageStr = (Math.round(num * 100000) / 1000).toFixed(places);
  return parseFloat(percentageStr);
}

export const titleCase = (str: string): string => {
  return str.split(" ")
    .map(w => w[0].toUpperCase() + w.substr(1).toLowerCase())
    .join(" ");
}

export const getOptionsArrayFromKey = (options: any, name: string) => {
  return getOptionsFromKey(options, name, options);
}

export const getValueFromKey = (options: any, name: string) => {
  return getOptionsFromKey(options, name, []);
}

const getOptionsFromKey = (options: any, name: string, defaultValue: any) => {
  if (!isNonEmptyArray(options)) {
    return defaultValue;
  }

  const opt = options.find((o: any) => o.name === name);
  return opt && isNonEmptyArray(opt.value)
    ? opt.value
    : defaultValue;
}

const isNonEmptyArray = (obj: any) => {
  return Array.isArray(obj) && obj.length > 0;
}
