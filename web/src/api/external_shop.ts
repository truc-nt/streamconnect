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
