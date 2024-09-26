import { axiosJava } from "@/api/axios";
export interface Notification {
  id: number;
  userId: number;
  title: string;
  message: string;
  type: string;
  status: string;
  redirectUrl?: string;
  createdAt: string;
}

export const getNotifications = async (userId?: number, status?: string) => {
  let url = `/notification?`;
  if (!!userId) {
    url += `userId=${userId}`;
  }
  if (!!status) {
    url += `&status=${status}`;
  }
  const response = await axiosJava.get<Notification[]>(url);
  return response.data;
};

export interface NotificationStatusUpdateRequest {
  notificationIds: number[];
  status: string;
}

export const batchUpdateNotificationStatus = async (
  ids: number[],
  status: string,
) => {
  const url = `/notification/batch-status-update`;
  await axiosJava.post<NotificationStatusUpdateRequest>(url, {
    notificationIds: ids,
    status,
  });
};

export const notifyLivestreamProductFollowers = async (
  livestreamProductId: number,
) => {
  return axiosJava.get(
    `/ecommerce/notify-livestream-product-follower?productId=${livestreamProductId}`,
  );
};
