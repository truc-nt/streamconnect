import axios from "./axios";

export interface IExternalVariant {
  id_external_variant: number;
  id_variant: number | null;
  sku: string;
  option: { [key: string]: string };
  image_url: string;
  price: number;
}

export interface IExternalProduct {
  external_product_id_mapping: string;
  name: string;
  image_url: string;
  status: string;
  id_product: number;
  product_name: string;
  shop_name: string;
  external_variants: IExternalVariant[];
}

export const getExternalProducts = async () => {
  return axios.get<IExternalProduct[]>(`/external_products/`);
};

export const getExternalVariants = async (externalProductIdMapping: string) => {
  return axios.get<IExternalVariant[]>(
    `/external_products/${externalProductIdMapping}`,
  );
};

export const syncExternalVariants = async (externalShopId: number) => {
  return axios.post(`/external_shops/${externalShopId}/sync_external_variants`);
};
