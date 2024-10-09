import axios from "./axios";
import { IBaseLivestreamProduct } from "@/model/livestream";

export interface ILivestreamExternalVariant {
  id_livestream_external_variant: number;
  id_ecommerce: number;
  quantity: number;
  stock: number;
  price: number;
  image_url: string;
}

interface ILivestreamVariant {
  id_variant: number;
  option: { [key: string]: string };
  livestream_external_variants: ILivestreamExternalVariant[];
}

export interface ILivestreamProduct extends IBaseLivestreamProduct {
  /*id_product: number;
  name: string;
  description: string;
  image_url: string;*/
  livestream_variants: ILivestreamVariant[];
}

export const getLivestreamProduct = async (idLivestreamProduct: number) => {
  const res = await axios.get<ILivestreamProduct>(
    `livestream_products/${idLivestreamProduct}`,
  );
  return res.data;
};

export interface IUpdateLivestreamExternalVariantQuantity {
  id_livestream_external_variant: number;
  quantity: number;
}

export const updateLivestreamExternalVariantQuantity = async (
  updateLivestreamExternalVariantQuantity: IUpdateLivestreamExternalVariantQuantity[],
) => {
  const res = await axios.post<IUpdateLivestreamExternalVariantQuantity[]>(
    `livestream_external_variants/update_quantity`,
    updateLivestreamExternalVariantQuantity,
  );
  return res.data;
};

export interface IPinLivestreamProduct {
  id_livestream_product: number;
  priority: number;
}

export const updateLivestreamProductPriority = async (
  pinLivestreamProduct: IPinLivestreamProduct[],
) => {
  const res = await axios.post<IPinLivestreamProduct[]>(
    `livestream_products/priority`,
    pinLivestreamProduct,
  );
  return res.data;
};

export interface IUpdateLivestreamProductRequest {
  id_livestream_product: number;
  priority?: number;
  is_livestreamed?: boolean;
}
export const updateLivestreamProduct = async (
  livestreamId: number,
  updateLivestreamProduct: IUpdateLivestreamProductRequest[],
) => {
  const res = await axios.patch(
    `livestreams/${livestreamId}/update_livestream_products`,
    updateLivestreamProduct,
  );
  return res.data;
};

export const registerLivestreamProductFollower = async (
  livestreamId: number,
  livestreamProductIds: number[],
) => {
  const res = await axios.post(
    `livestreams/${livestreamId}/livestream_products/follow`,
    livestreamProductIds,
  );
  return res.data;
};

export const getFollowLivestreamProductsInLivestream = async (
  livestreamId: number,
) => {
  const res = await axios.get<
    {
      fk_user: number;
      fk_livestream_product: number;
    }[]
  >(`livestreams/${livestreamId}/livestream_products/follow`);
  return res.data;
};

export const deleteFollowLivestreamProduct = async (
  livestreamProductId: number,
) => {
  const res = await axios.delete(
    `livestream_products/${livestreamProductId}/follow`,
  );
  return res.data;
};
