import Image from "next/image";
import CategorySlider from "@/component/livestream/CategorySlider";
import { Flex, Carousel } from "antd";

const Page = () => {
  return (
    <Flex vertical gap="large">
      <CategorySlider />
      <Carousel arrows infinite></Carousel>
    </Flex>
  );
};

export default Page;
