"use client";
import { useEffect, useMemo, useState, useRef } from "react";
import { Constants, useMeeting, usePubSub } from "@videosdk.live/react-sdk";
import dynamic from "next/dynamic";
import { Flex, Typography, Button } from "antd";

const ReactHlsPlayer = dynamic(() => import("react-hls-player"), {
  ssr: false,
});

const HlsPlayer = ({
  hlsUrl,
  playerRef,
}: {
  hlsUrl: string;
  playerRef: any;
}) => {
  return (
    <>
      {hlsUrl ? (
        <ReactHlsPlayer
          playsInline
          autoPlay
          controls={true}
          src={hlsUrl}
          height="100%"
          width="100%"
          playerRef={playerRef}
          hlsConfig={{
            maxLoadingDelay: 60,
            minAutoBitrate: 0,
            lowLatencyMode: true,
          }}
          muted={false}
          onError={(e) => console.error("Player error:", e)}
          className="rounded-lg max-h-full"
        />
      ) : (
        <Flex
          vertical
          justify="center"
          align="center"
          className="h-full bg-gray-800 rounded-lg"
        >
          <Typography.Title level={3} style={{ color: "white" }}>
            Người bán tạm thời chưa phát sóng
          </Typography.Title>
        </Flex>
      )}
    </>
  );
};

export default HlsPlayer;
