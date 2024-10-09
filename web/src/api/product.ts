import axios from "./axios";

import {
  IBaseProduct,
  IBaseVariant,
  IBaseExternalVariant,
} from "@/model/product";

export const getProductsByShopId = async (shopId: number) => {
  const res = await axios.get<IBaseProduct[]>(`/shops/${shopId}/products`);
  return res.data;
};

export const getProductById = async (productId: number) => {
  type IGetProductByIdResponse = IBaseProduct & {
    variants: (IBaseVariant & {
      external_variants: (IBaseExternalVariant & {
        shop_name: string;
      })[];
    })[];
  };
  const res = await axios.get<IGetProductByIdResponse>(
    `/products/${productId}`,
  );
  return res.data;
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
  external_variants: {
    id_external_variant: number;
    id_ecommerce: number;
    price: number;
    stock: number;
  }[];
}

export const getVariantsByProductId = async (productId: number) => {
  const res = await axios.get<IVariant[]>(`/products/${productId}/variants`);
  return res.data;
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

export interface IUpdateProductRequest {
  name?: string;
  description?: string;
}

export const updateProduct = async (
  productId: number,
  data: IUpdateProductRequest,
) => {
  return axios.patch(`/products/${productId}`, data);
};

export const getExternalVariantsByVariantId = async (variantId: number) => {
  const res = await axios.get<IBaseExternalVariant[]>(`/variants/${variantId}`);
  return res.data;
};
