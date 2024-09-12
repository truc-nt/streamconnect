import AddButton from "./component/AddButton";
import AddressInfoList from "./component/AddressInfoList";
import { Flex } from "antd";
const Page = () => {
  return (
    <Flex gap="middle" vertical>
      <AddButton />
      <AddressInfoList />
    </Flex>
  );
};

export default Page;
