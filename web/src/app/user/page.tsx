import UserInfoForm from "@/component/form/UserInfoForm";
import { Flex, Card } from "antd";

const Page = () => {
  return (
    <Flex>
      <Card title="Thông tin cá nhân" className="flex-1">
        <UserInfoForm />
      </Card>
    </Flex>
  );
};

export default Page;
