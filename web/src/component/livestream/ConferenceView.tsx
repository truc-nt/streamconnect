import { useParticipant } from "@videosdk.live/react-sdk";
import { Fragment, useEffect, useMemo, useRef, useState } from "react";
import ReactPlayer from "react-player";
import React, { memo } from "react";
import { Flex } from "antd";

const ConferenceView = ({ participantId }: { participantId: string }) => {
  const micRef = useRef<HTMLAudioElement>(null);
  const { webcamStream, micStream, webcamOn, micOn, isLocal, displayName } =
    useParticipant(participantId);

  const videoStream = useMemo(() => {
    if (webcamOn && webcamStream) {
      const mediaStream = new MediaStream();
      mediaStream.addTrack(webcamStream.track);
      return mediaStream;
    }
  }, [webcamStream, webcamOn]);

  useEffect(() => {
    if (micRef.current) {
      if (micOn && micStream) {
        const mediaStream = new MediaStream();
        mediaStream.addTrack(micStream.track);

        micRef.current.srcObject = mediaStream;
        micRef.current
          .play()
          .catch((error) =>
            console.error("videoElem.current.play() failed", error),
          );
      } else {
        micRef.current.srcObject = null;
      }
    }
  }, [micStream, micOn]);

  return (
    <div className="h-full w-full bg-gray-500 relative overflow-hidden rounded-lg">
      <audio ref={micRef} autoPlay playsInline muted={isLocal} />
      {webcamOn ? (
        <ReactPlayer
          playsinline
          playIcon={<></>}
          pip={false}
          light={false}
          controls={false}
          muted={false}
          playing={true}
          url={videoStream}
          width={"100%"}
          height={"100%"}
          onError={(err) => {
            console.log(err, "participant video error");
          }}
          style={
            isLocal
              ? { transform: "scaleX(-1)", WebkitTransform: "scaleX(-1)" }
              : {}
          }
          className="absolute object-cover"
          styles={{ position: "absolute", objectFit: "cover" }}
        />
      ) : (
        <div className="h-full w-full flex items-center justify-center">
          <div
            className={`z-10 flex items-center justify-center rounded-full bg-gray-800 2xl:h-[92px] h-[52px] 2xl:w-[92px] w-[52px]`}
          >
            <p className="text-2xl text-white">
              {String(displayName).charAt(0).toUpperCase()}
            </p>
          </div>
        </div>
      )}
    </div>
  );
};

const MemoizedConferenceView = memo(ConferenceView, (prevProps, nextProps) => {
  return prevProps.participantId === nextProps.participantId;
});

export default MemoizedConferenceView;
