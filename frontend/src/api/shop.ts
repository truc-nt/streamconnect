import axios from "./axios";

export interface IExternalShop {
  id_external_shop: number;
  name: string;
  ecommerce: string;
  status: string;
  created_at: string;
  updated_at: string;
}

export const getExternalShops = async (shopId: number) => {
  const res = await axios.get<IExternalShop[]>(
    `shops/${shopId}/external_shops`,
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
  option_titles: Record<string, string[]>;
  created_at: string;
  updated_at: string;
}

export const getProducts = async (shopId: number) => {
  const res = await axios.get<IProduct[]>(`shops/${shopId}/products`);
  return res.data;
};
