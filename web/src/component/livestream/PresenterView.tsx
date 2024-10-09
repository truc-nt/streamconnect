import { useMeeting, useParticipant } from "@videosdk.live/react-sdk";
import { Fragment, useEffect, useMemo, useRef, useState } from "react";
import ReactPlayer from "react-player";
import React, { memo } from "react";
import { Flex } from "antd";

const PresenterView = () => {
  const mMeeting = useMeeting();
  const presenterId = mMeeting?.presenterId;

  const {
    micOn,
    isLocal,
    screenShareStream,
    screenShareAudioStream,
    screenShareOn,
  } = useParticipant(presenterId ?? "");
  console.log("presenterId", presenterId);

  const mediaStream = useMemo(() => {
    if (screenShareOn) {
      const mediaStream = new MediaStream();
      mediaStream.addTrack(screenShareStream.track);
      return mediaStream;
    }
  }, [screenShareStream, screenShareOn]);

  const audioPlayer = useRef<HTMLAudioElement>(null);

  useEffect(() => {
    if (
      !isLocal &&
      audioPlayer.current &&
      screenShareOn &&
      screenShareAudioStream
    ) {
      const mediaStream = new MediaStream();
      mediaStream.addTrack(screenShareAudioStream.track);

      audioPlayer.current.srcObject = mediaStream;
      audioPlayer.current.play().catch((err) => {
        if (
          err.message ===
          "play() failed because the user didn't interact with the document first. https://goo.gl/xX8pDD"
        ) {
          console.error("audio" + err.message);
        }
      });
    }
  }, [screenShareAudioStream, screenShareOn, isLocal]);

  return (
    <div className="h-full w-full bg-gray-500 relative overflow-hidden rounded-lg">
      <audio autoPlay playsInline controls={false} ref={audioPlayer} />
      <ReactPlayer
        playsinline
        playIcon={<></>}
        pip={false}
        light={false}
        controls={false}
        muted={false}
        playing={true}
        url={mediaStream}
        width={"100%"}
        height={"100%"}
        onError={(err) => {
          console.log(err, "participant video error");
        }}
        className="absolute object-cover"
        styles={{ position: "absolute", objectFit: "cover" }}
      />
    </div>
  );
};

const MemoizedPresenterView = memo(PresenterView, (prevProps, nextProps) => {
  return true;
});

export default MemoizedPresenterView;
