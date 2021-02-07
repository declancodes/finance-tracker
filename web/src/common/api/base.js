import querystring from 'query-string';
import sortBy from 'lodash.sortby';

export const create = (url, values) => {
  return fetch(url, {
    headers: { 'Content-Type': 'application/json; charset=UTF-8' },
    method: 'POST',
    body: JSON.stringify(values)
  });
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

  return fetch(url)
    .then(response => response.json())
    .then(json => json);
};

export const update = (url, values) => {
  return fetch(url, {
    headers: { 'Content-Type': 'application/json; charset=UTF-8' },
    method: 'PUT',
    body: JSON.stringify(values)
  });
};

export const remove = (url) => {
  return fetch(url, {
    headers: { 'Content-Type': 'application/json' },
    method: 'DELETE'
  });
};

export const sort = (promise, order) => {
  return promise
    .then(response => {
      return hasNoData(response)
        ? []
        : sortBy(response, order)
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
          entities: sortBy(response[property], order),
          total: response.total
        };
    })
    .catch(error => {
      if (notFound(error)) {
        return emptyTotal;
      }
    });
};

const hasNoData = (response) => {
  return response === undefined || response === null;
};

const notFound = (error) => {
  return error.response && error.response.status === 404
};
