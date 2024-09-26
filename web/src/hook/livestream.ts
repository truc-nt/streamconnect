import {
  getLivestreamProducts,
  getLivestreams,
  getLivestream,
} from "@/api/livestream";
import useSWR from "swr";

export const useGetAllLivestreams = (shopId: number) => {
  return useSWR(
    "/api/livestreams",
    () =>
      getLivestreams({
        shop_id: shopId,
      }),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};
export const useGetAllLivestreamsInStartedAndStreamingStatus = () => {
  return useSWR(
    "/api/livestreams",
    () =>
      getLivestreams({
        status: ["started", "streaming"],
      }),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};

export const useGetLivestreamProducts = (livestreamId: number) => {
  return useSWR(
    `/api/livestreams/${livestreamId}/livestream_products`,
    () => getLivestreamProducts(livestreamId),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};

export const useGetLivestream = (livestreamId: number) => {
  return useSWR(
    `/api/livestreams/${livestreamId}`,
    () => getLivestream(livestreamId),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};
