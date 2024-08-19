import { useGetLivestreamProducts } from "@/hook/livestream";
import { message } from "@/store/antd_app";
import ProductCard from "@/component/product/Card";
import { List } from "antd";
import LivestreamProductInfo from "./LivestreamProductInfo";
import { useState } from "react";

const LivestreamProductSegmented = ({
  livestreamId,
}: {
  livestreamId: number;
}) => {
  const { data } = useGetLivestreamProducts(livestreamId);
  const [livestreamProductId, setLivestreamProductId] = useState<number | null>(
    null,
  );

  /*if (!data) {
    return null;
  }*/
  return (
    <>
      <List
        grid={{ gutter: [2, 2], column: 1 }}
        dataSource={data?.data || []}
        renderItem={(item) => (
          <List.Item style={{ padding: 0 }}>
            <ProductCard
              {...item}
              onClick={() => {
                setLivestreamProductId(item.id_livestream_product);
              }}
            />
          </List.Item>
        )}
        className="overflow-y-scroll overflow-x-visible p-1"
      />
      {livestreamProductId !== null && (
        <LivestreamProductInfo
          livestreamProductId={livestreamProductId}
          setLivestreamProductId={setLivestreamProductId}
        />
      )}
    </>
  );
};

export default LivestreamProductSegmented;
