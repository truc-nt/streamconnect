import axios from "./axios";

import { IBaseShop } from "@/model/shop";

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

export interface IShopGetRequest extends IBaseShop {
  is_following: boolean;
}

export const getShop = async (shopId: number) => {
  const res = await axios.get<IShopGetRequest>(`shops/${shopId}`);
  return res.data;
};
