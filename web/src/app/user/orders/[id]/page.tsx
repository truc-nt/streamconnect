import { Flex } from "antd";
import OrderAdditionInfo from "@/component/order/OrderAdditionInfo";
import ExternalOrderTabs from "../component/ExternalOrderTabs";

const Page = ({ params }: { params: { id: string } }) => {
  const test = {
    name: "test",
    phone: "test",
    address: "test",
    city: "test",
    id_shipping_method: 1,
    id_payment_method: 1,
  };
  return (
    <Flex vertical gap="large">
      <OrderAdditionInfo {...test} />
      <ExternalOrderTabs />
    </Flex>
  );
};

export default Page;
