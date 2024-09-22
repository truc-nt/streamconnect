import axios from "./axios";

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

export interface ILivestreamProduct {
  id_product: number;
  name: string;
  description: string;
  image_url: string;
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

export const pinLivestreamProduct = async (
  pinLivestreamProduct: IPinLivestreamProduct[],
) => {
  const res = await axios.post<IPinLivestreamProduct[]>(
    `livestream_products/pin`,
    pinLivestreamProduct,
  );
  return res.data;
};
