import axios from "./axios";

export interface ILivestreamProduct {
  name: string;
  description: string;
  image_url: string;
  livestream_variants: {
    option: { [key: string]: string };
    livestream_external_variants: {
      id_livestream_external_variant: number;
      id_ecommerce: number;
      quantity: number;
      price: number;
      image_url: string;
    }[];
  }[];
}
export const getLivestreamProduct = async (idLivestreamProduct: number) => {
  const res = await axios.get<ILivestreamProduct>(
    `livestream_products/${idLivestreamProduct}`,
  );
  return res.data;
};
