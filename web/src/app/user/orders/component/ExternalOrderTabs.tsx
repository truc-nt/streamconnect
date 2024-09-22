"use client";
import { useParams } from "next/navigation";

import { Tabs } from "antd";

import ExternalOrderTable from "@/component/table/ExternalOrderTable";
import { useGetOrder } from "@/hook/order";
import { ECOMMERCE_PLATFORMS } from "@/constant/ecommerce";

const ExternalOrderTabs = () => {
  const { id: orderId } = useParams();
  const { data: orders } = useGetOrder(Number(orderId));
  return (
    <Tabs
      defaultActiveKey="1"
      centered
      items={orders?.map((order) => {
        return {
          label: ECOMMERCE_PLATFORMS[order.id_ecommerce],
          key: order.id_ecommerce.toString(),
          children: (
            <ExternalOrderTable {...order} id_ecommerce={order.id_ecommerce} />
          ),
        };
      })}
    />
  );
};

export default ExternalOrderTabs;
