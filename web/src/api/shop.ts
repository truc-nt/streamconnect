import axios from "./axios";

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
