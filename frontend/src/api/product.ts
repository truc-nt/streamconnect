import axios from "./axios";

export interface IVariant {
  id_variant: number;
  sku: string;
  status: string;
  option: Record<string, string>;
  external_products: {
    id_ecommerce: number;
    id_external_product: number;
    ecommerce: string;
    price: number;
    stock: number;
  }[];
}

export const getVariants = async (productId: number) => {
  return axios.get<IVariant[]>(`/products/${productId}/variants`);
};
