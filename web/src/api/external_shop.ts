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

interface IExternalProduct {
  external_product_external_id: number;
  external_product_name: string;
  fk_product: number;
  product_name: string;
  total_stock: number;
  min_price: number;
  max_price: number;
  image_url: string;
  updated_at: string;
}

export const getExternalProducts = async (externalShopId: number) => {
  return axios.get<IExternalProduct[]>(
    `/external_shops/${externalShopId}/external_variants`,
  );
};

export const syncExternalVariants = async (externalShopId: number) => {
  return axios.get(`/external_shops/${externalShopId}/sync_external_variants`);
};
