"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { Card, Flex, Typography, Button } from "antd";

import { IGetShopResponse, followShop } from "@/api/shop";
import useLoading from "@/hook/loading";

const ShopInfo = ({ id_shop, name, is_following }: IGetShopResponse) => {
  const router = useRouter();
  const executeFollowShop = useLoading(followShop, "", "", () => {
    setIsFollowing(true);
  });
  const [isFollowing, setIsFollowing] = useState(is_following);

  return (
    <Flex
      justify="space-between"
      align="center"
      className="bg-white p-4 rounded-lg"
    >
      <Typography.Title
        level={3}
        className="cursor-pointer"
        style={{ margin: 0 }}
        onClick={() => router.push(`/shops/${id_shop}`)}
      >
        {name}
      </Typography.Title>
      {!isFollowing ? (
        <Button
          type="primary"
          onClick={async () => {
            await executeFollowShop(id_shop);
          }}
        >
          Theo dõi
        </Button>
      ) : (
        <Button type="default">Đang theo dõi</Button>
      )}
    </Flex>
  );
};

export default ShopInfo;
