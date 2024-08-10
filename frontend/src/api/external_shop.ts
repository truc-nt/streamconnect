import axios from "./axios";

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
    `/external_shops/${externalShopId}/external_products`,
  );
};

export const syncExternalProducts = async (externalShopId: number) => {
  return axios.get(`/external_shops/${externalShopId}/sync_external_products`);
};
