import axios from "./axios";

interface IExternalProduct {
  id_external_product_shopify: number;
  fk_external_shop: number;
  fk_product: number;
  fk_variant: number;
  shopify_product_id: number;
  shopify_variant_id: number;
  name: string;
  sku: string;
  stock: number;
  option: Record<string, string>;
  price: number;
  image_url: string;
  created_at: string;
  updated_at: string;
}

export const getExternalProducts = async (externalShopId: number) => {
  return axios.get<IExternalProduct[]>(
    `/external_shops/${externalShopId}/external_products`,
  );
};
