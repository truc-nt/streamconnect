import axios from "./axios";

export interface ILivestreamProduct {
  name: string;
  description: string;
  option: Record<string, string[]>;
  livestream_variants: {
    option: Record<string, string>;
    livestream_external_variants: {
      id_livestream_external_variant: number;
      id_ecommerce: number;
      quantity: number;
      price: number;
    }[];
  }[];
}
export const getLivestreamProduct = async (idLivestreamProduct: number) => {
  return axios.get<ILivestreamProduct>(
    `livestream_products/${idLivestreamProduct}`,
  );
};
