import {
  getLivestreamProducts,
  getLivestreams,
  getLivestreamInfo,
  saveHls,
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

export const useGetLivstreamInfo = (livestreamId: number) => {
  return useSWR(
    `/api/livestreams/${livestreamId}/info`,
    () => getLivestreamInfo(livestreamId),
    {
      //shouldRetryOnError: false,
      revalidateOnFocus: false,
    },
  );
};
