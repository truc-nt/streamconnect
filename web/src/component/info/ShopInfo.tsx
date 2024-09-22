"use client";
import { useRouter } from "next/navigation";
import { Card, Flex, Typography, Button } from "antd";

import { IShopGetRequest } from "@/api/shop";

const ShopInfo = ({ id_shop, name, is_following }: IShopGetRequest) => {
  const router = useRouter();
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
      {is_following ? (
        <Button type="primary">Theo dõi</Button>
      ) : (
        <Button type="default">Đang theo dõi</Button>
      )}
    </Flex>
  );
};

export default ShopInfo;
