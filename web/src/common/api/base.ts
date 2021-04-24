import { StringifiableRecord, StringifyOptions, stringifyUrl, UrlObject } from "query-string";
import { TotalEntity } from "../types/entity";

export const create = async (url: string, values: any): Promise<Response> => {
  return await fetch(url, {
    headers: { "Content-Type": "application/json; charset=UTF-8" },
    method: "POST",
    body: JSON.stringify(values)
  });
};

export const get = async (baseUrl: string, filterParams: StringifiableRecord): Promise<Response | null> => {
  const parsedUrl: UrlObject = {
    url: baseUrl,
    query: filterParams
  };
  const options: StringifyOptions = {
    skipNull: true,
    skipEmptyString: true
  };

  const url = stringifyUrl(parsedUrl, options);
  const response = await fetch(url);

  return response.ok ? response : null;
};

export const update = async (url: string, values: any): Promise<Response> => {
  return await fetch(url, {
    headers: { "Content-Type": "application/json; charset=UTF-8" },
    method: "PUT",
    body: JSON.stringify(values)
  });
};

export const remove = async (url: string): Promise<Response> => {
  return await fetch(url, {
    headers: { "Content-Type": "application/json" },
    method: "DELETE"
  });
};

export const getEntities = async (promise: Promise<Response | null>): Promise<any[]> => {
  const response = await promise;
  if (!response || !response.ok) {
    return [];
  }

  return await response.json();
};

export const getTotalAmountEntity = async (promise: Promise<Response | null>, property: string): Promise<TotalEntity> => {
  const response = await promise;
  if (!response || !response.ok) {
    return { entities: [], total: 0 };
  }

  const data = await response.json();
  return {
    entities: data[property],
    total: data.total
  };
};
