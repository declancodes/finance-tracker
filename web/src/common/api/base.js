import axios from 'axios';
import querystring from 'query-string';
import sortBy from 'lodash.sortby';

export const create = (url, values) => {
  return axios.post(url, values)
};

export const get = (baseUrl, filterParams) => {
  const parsedUrl = {
    url: baseUrl,
    query: filterParams
  };
  const options = {
    skipNull: true,
    skipEmptyString: true
  };
  const url = querystring.stringifyUrl(parsedUrl, options);

  return axios.get(url);
};

export const update = (url, values) => {
  return axios.put(url, values)
};

export const remove = (url) => {
  return axios.delete(url)
};

export const sort = (promise, order) => {
  return promise
    .then(response =>
      response.data === undefined || response.data === null
          ? []
          : sortBy(response.data, order)
    );
};

export const sortTotal = (promise, property, order) => {
  return promise
    .then(response => {
      const hasNoData = response.data === undefined || response.data === null;
      return {
        entities: hasNoData ? [] : sortBy(response.data[property], order),
        total: hasNoData ? 0 : response.data.total
      };
    });
};
