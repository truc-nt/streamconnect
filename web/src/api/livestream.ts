import axios from "./axios";
import qs from "qs";

import { IBaseLivestream, IBaseLivestreamProduct } from "@/model/livestream";

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

export const getLivestreamProducts = async (livestreamId: number) => {
  const res = await axios.get<IBaseLivestreamProduct[]>(
    `livestreams/${livestreamId}/livestream_products`,
  );
  return res.data;
};

interface IGetLivestreamResponse extends IBaseLivestream {
  is_host: boolean;
  shop_name: string;
}

export const getLivestream = async (livestreamId: number) => {
  const res = await axios.get<IGetLivestreamResponse>(
    `livestreams/${livestreamId}`,
  );
  return res.data;
};

interface IUpdateLivestreamRequest {
  title: string;
  description: string;
  status: string;
  meeting_id: string;
  hls_url: string;
  start_time: string;
  end_time: string;
}
export const updateLivestream = async (
  livestreamId: number,
  request: IUpdateLivestreamRequest,
) => {
  const res = await axios.patch(`livestreams/${livestreamId}`, request);
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
