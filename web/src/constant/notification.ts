export enum NotificationType {
  MESSAGE = "MESSAGE",
  ORDER = "ORDER",
}

export enum NotificationStatus {
  NEW = "NEW",
  SEND = "SEND",
  READ = "READ",
}

export const NotificationLogoMap: Record<NotificationType, string> = {
  [NotificationType.MESSAGE]: "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png",
  [NotificationType.ORDER]: "https://gw.alipayobjects.com/zos/rmsportal/OKJXDXrmkNshAMvwtvhu.png",
}