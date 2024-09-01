import axios from "./axios";
import {axiosJava} from "./axios";

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
  console.log(data);
  return axios.post(`shops/${shopId}/livestreams/create`, data);
};

export const fetchMeeting = async (status: string, findAll: boolean) => {
  const params = new URLSearchParams();
  if (status) {
    params.append("status", status);
  }
  params.append("fetchAll", findAll.toString());

  try {
    const res = await axiosJava.get(`/livestream`, {params});

    return res.data;
  } catch (error) {
    console.log(error);
    return [];
  }
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
  const res = await axios.get<ILivestreamProduct[]>(`livestreams/${livestreamId}/products`)
  return res.data;
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
