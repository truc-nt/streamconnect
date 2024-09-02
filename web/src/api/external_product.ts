import axios from "./axios";

export interface IExternalVariant {
  id_external_variant: number;
  fk_variant: number | null;
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
  const res = await axios.get<IExternalVariant[]>(
    `/external_products/${externalProductIdMapping}`,
  );
  return res.data;
};

export const syncExternalVariants = async (externalShopId: number) => {
  return axios.post(`/external_shops/${externalShopId}/sync_external_variants`);
};

export interface IVariant {
  id_variant: number;
  fk_product: number;
  sku: string;
  status: string;
  option: { [key: string]: string };
  created_at: string;
  updated_at: string;
  image_url: string;
}

export const getVariantsByExternalProductIdMapping = async (
  externalProductIdMapping: string,
) => {
  const res = await axios.get<IVariant[]>(
    `/external_products/${externalProductIdMapping}/variants`,
  );
  return res.data;
};

export const connectVariants = async (
  connectVariantsRequest: {
    id_variant: number;
    id_external_variant: number;
  }[],
) => {
  const res = await axios.post(
    `/external_variants/connect`,
    connectVariantsRequest,
  );
  return res.data;
};
