"use client";

import React from "react";
import styles from "./LiveStreamList.module.css";
import {EyeOutlined} from "@ant-design/icons";
import {Avatar, Badge, Button, Col, Row, Typography} from "antd";

interface Livestream {
  imageUrl: string;
  avatarUrl: string;
  username: string;
  livestreamName: string;
  views: number;
}

// Mock data for livestreams
const mockLivestreams: Livestream[] = [
  {
    imageUrl:
      "https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png",
    avatarUrl:
      "https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png",
    username: "User1",
    livestreamName: "Livestream 1",
    views: 1000,
  },
  {
    imageUrl:
      "https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png",
    avatarUrl:
      "https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png",
    username: "User2",
    livestreamName: "Livestream 2",
    views: 800,
  },
  {
    imageUrl:
      "https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png",
    avatarUrl:
      "https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png",
    username: "User3",
    livestreamName: "Livestream 3",
    views: 600,
  },
  {
    imageUrl:
      "https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png",
    avatarUrl:
      "https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png",
    username: "User3",
    livestreamName: "Livestream 3",
    views: 600,
  },
  {
    imageUrl:
      "https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png",
    avatarUrl:
      "https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png",
    username: "User3",
    livestreamName: "Livestream 3",
    views: 600,
  },
  {
    imageUrl:
      "https://s-aicmscdn.nss.vn/nss-media/23/7/3/livestream-shopping-da-phat-trien-thanh-thi-truong-512-ty-usd_64a24b998d9e0.png",
    avatarUrl:
      "https://www.ibm.com/brand/experience-guides/developer/static/28e7b09c59aad7dce5e93296bc425b90/a5df1/Red-Hat_mobile.png",
    username: "User3",
    livestreamName: "Livestream 3",
    views: 600,
  },
];

const formatViews = (views: number) => {
  return views.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
};

const LivestreamsList: React.FC<any> = ({meetingIds}: {meetingIds: string[]}) => {
  const handleLivestreamClick = (livestream: Livestream) => {
    console.log(`Clicked on ${livestream.livestreamName}`);
  };

  return (
      <div className={styles.mt3}>
        <Row gutter={[16, 16]}>
          {mockLivestreams.map((livestream, index) => (
              <Col key={index} xs={24} sm={12} md={8}>
                <Button
                    onClick={() => handleLivestreamClick(livestream)}
                    className={styles.fullWidthButton}
                >
                  <div className={styles.relativeContainer}>
                    <img
                        src={livestream.imageUrl}
                        alt="Livestream Thumbnail"
                        className={styles.thumbnail}
                    />
                    <div className={styles.viewerCount}>
                      <EyeOutlined className={styles.icon} />
                      <Typography.Text className={styles.whiteText}>
                        {formatViews(livestream.views)}
                      </Typography.Text>
                    </div>
                    <div className={styles.badgeContainer}>
                      <Badge count="Live" color="red" />
                    </div>
                    <div className={styles.avatarInfoContainer}>
                      <Avatar
                          alt={livestream.username}
                          src={livestream.avatarUrl}
                          className={styles.avatar}
                      />
                      <div className={styles.avatarDetails}>
                        <Typography.Text strong className={styles.whiteText}>
                          {livestream.username}
                        </Typography.Text>
                        <Typography.Text className={styles.whiteText}>
                          {livestream.livestreamName}
                        </Typography.Text>
                      </div>
                    </div>
                  </div>
                </Button>
              </Col>
          ))}
        </Row>
      </div>
  );
};

export default LivestreamsList;
