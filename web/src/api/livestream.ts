import axios from "./axios";
import qs from "qs";

import { IBaseLivestream } from "@/model/livestream";

export const getLivestreams = async (queryParam: { [key: string]: any }) => {
  const res = await axios.get<IBaseLivestream[]>(
    `livestreams/?${decodeURIComponent(qs.stringify(queryParam, { arrayFormat: "brackets" }))}`,
  );
  return res.data;
};

export interface ILivestream {
  id_livestream: number;
  fk_shop: number;
  title: string;
  description: string;
  status: string;
  meeting_id: string;
  hls_url: string;
}

export interface ILivestreamExternalVariant {
  id_external_variant: number;
  quantity: number;
}

export interface ILivestreamProduct {
  id_product: number;
  priority: number;
  livestream_variants: {
    id_variant: number;
    livestream_external_variants: ILivestreamExternalVariant[];
  }[];
}

interface ILivestreamWithProducts extends ILivestream {
  livestream_products: ILivestreamProduct[];
}

export const createLivestream = async (
  shopId: number,
  data: ILivestreamWithProducts,
) => {
  return axios.post(`shops/${shopId}/livestreams/create`, data);
};

export interface ILivestreamProductInformation {
  id_livestream_product: number;
  id_product: number;
  name: string;
  min_price: number;
  max_price: number;
  image_url: string;
  priority: number;
}

export const getLivestreamProducts = async (livestreamId: number) => {
  const res = await axios.get<ILivestreamProductInformation[]>(
    `livestreams/${livestreamId}/livestream_products`,
  );
  return res.data;
};

interface IGetMeetingIdResponse {
  meeting_id: string;
  is_host: boolean;
  id_shop: number;
  shop_name: string;
}

export const getLivestreamInfo = async (livestreamId: number) => {
  const res = await axios.get<IGetMeetingIdResponse>(
    `livestreams/${livestreamId}/info`,
  );
  return res.data;
};

export const saveHls = async (livestreamId: number, hlsUrl: string) => {
  const res = await axios.post(`livestreams/${livestreamId}/save_hls`, {
    hls_url: hlsUrl,
  });
  return res.data;
};

export const addLivestreamProduct = async (
  livestreamId: number,
  data: ILivestreamProduct[],
) => {
  const res = await axios.post(
    `livestreams/${livestreamId}/add_livestream_product`,
    data,
  );
  return res.data;
};

export const startLivestream = async (livestreamId: number) => {
  const res = await axios.put(`livestreams/${livestreamId}/start`);
  return res.data;
};
