import axios from "./axios";

export interface ILivestreamExternalVariant {
  id_external_variant: number;
  quantity: number;
}

interface ILivestream {
  title: string;
  description: string;
  start_time: string;
  livestream_products: {
    id_product: number;
    priority: number;
    livestream_variants: {
      id_variant: number;
      livestream_external_variants: ILivestreamExternalVariant[];
    };
  };
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
  image_url: string;
  priority: number;
}

export const getLivestreamProducts = async (livestreamId: number) => {
  return axios.get<ILivestreamProduct[]>(
    `livestreams/${livestreamId}/livestream_products`,
  );
};
