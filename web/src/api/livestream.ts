import axios from "./axios";

interface ILivestream {
  title: string;
  description: string;
  startTime: string;
  livestream_external_variants: {
    id_product: number;
    id_variant: number;
    id_external_variant: number;
    quantity: number;
  }[];
}

export const createLivestream = async (shopId: number, data: ILivestream) => {
  return axios.post(`shops/${shopId}/livestreams/create`, data);
};

export interface ILivestreamProduct {
  id_livestream_product: number;
  id_product: number;
  name: string;
  min_price: number;
  max_price: number;
  priority: number;
}

export const getLivestreamProducts = async (livestreamId: number) => {
  return axios.get<ILivestreamProduct[]>(
    `livestreams/${livestreamId}/products`,
  );
};

export interface ILivestreamExternalVariant {
  option: Record<string, string>;
  variants: {
    option: Record<string, string>;
    livestream_external_variants: {
      id_external_variant: number;
      ecommerce: string;
      quantity: number;
      price: number;
    }[];
  }[];
}
export const getLivstreamExternalVariants = async (
  idLivestreamProduct: number,
) => {
  return axios.get(`livestream_products/${idLivestreamProduct}`);
};
