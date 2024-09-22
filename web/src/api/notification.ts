import {axiosJava} from "@/api/axios";
export interface Notification {
  id: number;
  userId: number; // Assuming user ID is a number
  title: string;
  message: string;
  type: string;
  status: string;
  redirectUrl?: string;
  createdAt: string; // ISO 8601 date string
}

export const getNotifications = async (userId?: number, status?: string) => {
  // Assuming the API endpoint is /notifications
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

export const batchUpdateNotificationStatus = async (ids: number[], status: string) => {
  // Assuming the API endpoint is /notifications
  const url = `/notification/batch-status-update`;
  await axiosJava.post<NotificationStatusUpdateRequest>(url, {
    notificationIds: ids,
    status,
  });
}