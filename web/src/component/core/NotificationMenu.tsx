import React, { ReactNode, useEffect, useState } from "react";
import { Badge, Dropdown, List, Avatar, Button } from "antd";
import styles from "./NotificationMenu.module.css";
import {
  batchUpdateNotificationStatus,
  Notification,
} from "@/api/notification";
import {
  NotificationLogoMap,
  NotificationStatus,
  NotificationType,
} from "@/constant/notification";

const NotificationDropdown = ({
  items,
  newItemCount,
  children,
}: {
  items: Notification[];
  newItemCount: number;
  children: ReactNode;
}) => {
  const renderNotification = (item: Notification) => {
    return (
      <List.Item
        actions={[
          item.status != NotificationStatus.READ && (
            <span className={styles.unreadDot} />
          ),
        ]}
      >
        <List.Item.Meta
          avatar={
            <Avatar src={NotificationLogoMap[item.type as NotificationType]} />
          }
          title={item.title}
          description={item.message}
        />
      </List.Item>
    );
  };

  const menu = (
    <div className={styles.notificationMenuContainer}>
      <List
        itemLayout="horizontal"
        dataSource={items}
        renderItem={renderNotification}
      />
    </div>
  );

  return (
    <Dropdown
      dropdownRender={() => menu}
      trigger={["click"]}
      placement="bottomRight"
    >
      <Badge count={newItemCount}>{children}</Badge>
    </Dropdown>
  );
};

export default NotificationDropdown;
