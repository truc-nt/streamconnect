import { ILivestream } from "@/api/livestream";
import { Card } from "antd";
import dynamic from "next/dynamic";
import { useRef } from "react";
import HlsPlayer from "@/component/livestream/HlsPlayer";
import { useRouter } from "next/navigation";

const LivestreamPreview = ({
  title,
  description,
  hls_url,
  id_livestream,
}: ILivestream) => {
  const playerRef = useRef(null);
  const router = useRouter();
  return (
    <Card
      cover={<HlsPlayer hlsUrl={hls_url} playerRef={playerRef} />}
      onClick={() => router.push(`/livestreams/${id_livestream}`)}
    >
      <Card.Meta title={title} description={description} />
    </Card>
  );
};

export default LivestreamPreview;
