import { getNotifications } from "@/api/notification";

import useSWR from "swr";

export const useGetNotification = () => {
  return useSWR(`/notification`, () => getNotifications(), {
    revalidateOnFocus: false,
  });
};
