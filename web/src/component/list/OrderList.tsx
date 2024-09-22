"use client";
import { useRouter, usePathname } from "next/navigation";

import { List } from "antd";

import { useGetBuyOrders } from "@/hook/order";
import OrderTable from "@/component/table/OrderTable";

const OrderList = () => {
  const { data: buyOrders } = useGetBuyOrders();
  const router = useRouter();
  const pathname = usePathname();

  return (
    <List
      grid={{ gutter: 16, column: 1 }}
      dataSource={buyOrders}
      renderItem={(item) => (
        <List.Item onClick={() => router.push(`${pathname}/${item.id_order}`)}>
          <OrderTable {...item} />
        </List.Item>
      )}
    />
  );
};

export default OrderList;
