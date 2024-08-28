import axios from "./axios";

export interface IProduct {
  id_product: number;
  name: string;
  description: string;
  status: string;
  image_url: string;
  created_at: string;
  updated_at: string;
}

export const getProductsByShopId = async (shopId: number) => {
  return axios.get<IProduct[]>(`/shops/${shopId}/products`);
};

export interface IVariant {
  id_variant: number;
  sku: string;
  status: string;
  option: { [key: string]: string };
  external_variants: {
    id_external_variant: number;
    id_ecommerce: number;
    price: number;
    stock: number;
  }[];
}

export const getVariantsByProductId = async (productId: number) => {
  return axios.get<IVariant[]>(`/products/${productId}/variants`);
};

export interface IProductWithVariants {
  external_product_id_mapping: string;
}

export const createProductWithVariants = async (
  shopId: number,
  createProductWithVariants: IProductWithVariants[],
) => {
  return axios.post(`/shops/${shopId}/products/`, createProductWithVariants);
};
