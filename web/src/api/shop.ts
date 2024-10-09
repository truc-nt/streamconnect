import qs from "qs";
import axios from "./axios";

import { IBaseShop } from "@/model/shop";

export const connectExternalShop = async (
  ecommerce: string,
  queryParams: { [key: string]: any },
) => {
  const res = await axios.get(
    `${ecommerce}/connect?${decodeURIComponent(qs.stringify(queryParams, { arrayFormat: "brackets" }))}`,
  );
  return res.data;
};

export interface IProduct {
  id_product: number;
  name: string;
  description: string;
  status: string;
  min_price: number;
  max_price: number;
  total_stock: number;
  option: Record<string, string[]>;
  created_at: string;
  updated_at: string;
}

export const getProducts = async (shopId: number) => {
  const res = await axios.get<IProduct[]>(`shops/${shopId}/products`);
  return res.data;
};

export interface IGetShopResponse extends IBaseShop {
  is_following: boolean;
}

export const getShop = async (shopId: number) => {
  const res = await axios.get<IGetShopResponse>(`shops/${shopId}`);
  return res.data;
};

interface IUpdateShopRequest {
  name: string;
  description: string;
}

export const updateShop = async (shopId: number, data: IUpdateShopRequest) => {
  const res = await axios.patch(`shops/${shopId}`, data);
  return res;
};

export const followShop = async (shopId: number) => {
  const res = await axios.post(`shops/${shopId}/follow`);
  return res;
};
