import {
  getLivestreamProduct,
  getFollowLivestreamProductsInLivestream,
} from "@/api/livestream_product";
import useSWR from "swr";

export const useGetLivestreamProduct = (livestreamProductId: number) => {
  return useSWR(
    `/api/livestream_products/${livestreamProductId}`,
    () => getLivestreamProduct(livestreamProductId),
    {
      revalidateOnFocus: false,
    },
  );
};

export const useGetFollowLivestreamProductsInLivestream = (
  livestreamId: number,
) => {
  return useSWR(
    `/api/livestreams/${livestreamId}/livestream_products/follow`,
    () => getFollowLivestreamProductsInLivestream(livestreamId),
    {
      revalidateOnFocus: false,
    },
  );
};
