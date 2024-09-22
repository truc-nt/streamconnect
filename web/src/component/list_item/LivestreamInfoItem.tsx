import { Card } from "antd";

import { IBaseLivestream } from "@/model/livestream";

const LivestreamInfoItem = ({ title, description }: IBaseLivestream) => {
  return (
    <Card bordered={false}>
      <Card.Meta title={title} description={description} />
    </Card>
  );
};

export default LivestreamInfoItem;
