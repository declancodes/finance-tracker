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
    .then(response => {
      return hasNoData(response)
        ? []
        : sortBy(response.data, order)
    })
    .catch(error => {
      if (notFound(error)) {
        return [];
      }
    });
};

export const sortTotal = (promise, property, order) => {
  const emptyTotal = { entities: [], total: 0 };

  return promise
    .then(response => {
      return hasNoData(response)
        ? emptyTotal
        : {
          entities: sortBy(response.data[property], order),
          total: response.data.total
        };
    })
    .catch(error => {
      if (notFound(error)) {
        return emptyTotal;
      }
    });
};

const hasNoData = (response) => {
  return response.data === undefined || response.data === null;
};

const notFound = (error) => {
  return error.response && error.response.status === 404
};
