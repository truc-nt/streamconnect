import {
  getProductsByShopId,
  getProductById,
  getVariantsByProductId,
} from "@/api/product";

import useSWR from "swr";

export const useGetProductsByShopId = (shopId: number) => {
  return useSWR(
    `/api/shops/${shopId}/products`,
    () => getProductsByShopId(shopId),
    {
      revalidateOnFocus: false,
    },
  );
};

export const useGetProductById = (productId: number) => {
  return useSWR(`/api/products/${productId}`, () => getProductById(productId), {
    revalidateOnFocus: false,
  });
};

export const useGetVariantsByProductId = (productId: number) => {
  return useSWR(
    `/api/products/${productId}/variants`,
    () => getVariantsByProductId(productId),
    {
      revalidateOnFocus: false,
    },
  );
};
