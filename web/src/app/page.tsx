import Image from "next/image";
import CategorySlider from "@/component/livestream/CategorySlider";
import { Flex, Carousel } from "antd";
import LivestreamPreviewGrid from "@/app/component/LivestreamPreviewGrid";

const Page = () => {
  return (
    <Flex vertical gap="large">
      <CategorySlider />
      <Carousel arrows infinite></Carousel>
      <LivestreamPreviewGrid />
    </Flex>
  );
};

export default Page;
