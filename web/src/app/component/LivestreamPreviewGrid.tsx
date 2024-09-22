"use client";
import { List } from "antd";
import LivestreamPreview from "@/component/livestream/LivestreamPreview";
import { useGetAllLivestreamsInStartedAndStreamingStatus } from "@/hook/livestream";
const LivestreamPreviewGrid = () => {
  const { data } = useGetAllLivestreamsInStartedAndStreamingStatus();
  return (
    <List
      grid={{ gutter: 16, column: 4 }}
      dataSource={data ?? []}
      renderItem={(item) => (
        <List.Item>
          <LivestreamPreview {...item} />
        </List.Item>
      )}
    />
  );
};

export default LivestreamPreviewGrid;
