import { Flex, Button } from "antd";
import Link from "next/link";
const Page = () => {
  return (
    <Flex>
      <Link href="/seller/livestreams/create">
        <Button type="primary">Tạo livestream mới</Button>
      </Link>
    </Flex>
  );
};

export default Page;
