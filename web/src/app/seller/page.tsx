import { Card } from "antd";

import ShopForm from "@/component/form/ShopForm";

const Page = () => {
  return (
    <Card title="Thông tin cửa hàng" className="flex-1">
      <ShopForm />
    </Card>
  );
};

export default Page;
