"use client";
import { useParams } from "next/navigation";

import {
  Constants,
  createCameraVideoTrack,
  useMeeting,
  usePubSub,
} from "@videosdk.live/react-sdk";
import React, { Fragment, useEffect, useMemo, useRef, useState } from "react";
import { Button, Flex, Dropdown, Space, theme, ButtonProps } from "antd";
import {
  AudioMutedOutlined,
  AudioOutlined,
  VideoCameraOutlined,
  DownOutlined,
  TeamOutlined,
  MessageOutlined,
  PlayCircleOutlined,
  PauseCircleOutlined,
  BorderOutlined,
  SkinFilled,
  WalletOutlined,
  DesktopOutlined,
} from "@ant-design/icons";
import { useIsHls } from "@/hook/hls";
import { useIsRecording } from "@/hook/recording";
import useLoading from "@/hook/loading";
import { saveHls, updateLivestream } from "@/api/livestream";
import { LivestreamStatus } from "@/constant/livestream";
import { useRouter } from "next/navigation";

const EndButton = ({ livestreamId }: { livestreamId: number }) => {
  const { end } = useMeeting();
  const executeUpdateLivestream = useLoading(updateLivestream);
  const router = useRouter();

  return (
    <Button
      size="large"
      type="primary"
      onClick={async () => {
        end();
        await executeUpdateLivestream(livestreamId, {
          status: LivestreamStatus.ENDED,
        });
        router.push(`/seller/livestreams`);
      }}
    >
      END
    </Button>
  );
};

const HlsButton = ({ livestreamId }: { livestreamId: number }) => {
  const { startHls, stopHls, hlsState, hlsUrls } = useMeeting({});

  const isHls = useIsHls();

  const isHlsRef = useRef(isHls);

  const executeUpdateLivestream = useLoading(updateLivestream);

  useEffect(() => {
    isHlsRef.current = isHls;
  }, [isHls]);

  useEffect(() => {
    const startHls = async () => {
      if (hlsUrls.livestreamUrl) {
        try {
          await saveHls(livestreamId, hlsUrls.playbackHlsUrl);
        } catch (e) {}
        await executeUpdateLivestream(livestreamId, {
          status: LivestreamStatus.PLAYED,
          hls_url: hlsUrls.livestreamUrl,
        });
      } else {
        await executeUpdateLivestream(livestreamId, {
          status: LivestreamStatus.STARTED,
          hls_url: hlsUrls.livestreamUrl,
        });
      }
    };
    startHls();
  }, [hlsUrls]);

  const handleClick = () => {
    const isHls = isHlsRef.current;

    if (isHls) {
      stopHls();
    } else {
      startHls();
    }
  };

  return (
    <Button
      size="large"
      onClick={handleClick}
      type={isHls ? "primary" : "default"}
      icon={isHls ? <PauseCircleOutlined /> : <PlayCircleOutlined />}
      loading={
        hlsState === Constants.hlsEvents.HLS_STARTING ||
        hlsState === Constants.hlsEvents.HLS_STOPPING
      }
    />
  );
};

const RecordingButton = () => {
  const { startRecording, stopRecording, recordingState } = useMeeting();

  const isRecording = useIsRecording();

  const isRecordingRef = useRef(isRecording);

  useEffect(() => {
    isRecordingRef.current = isRecording;
  }, [isRecording]);

  const handleClick = () => {
    const isRecording = isRecordingRef.current;

    if (isRecording) {
      stopRecording();
    } else {
      startRecording();
    }
  };

  return (
    <Button
      size="large"
      onClick={handleClick}
      type={isRecording ? "primary" : "default"}
      loading={
        recordingState === Constants.recordingEvents.RECORDING_STARTING ||
        recordingState === Constants.recordingEvents.RECORDING_STOPPING
      }
    >
      REC
    </Button>
  );
};

const ShareScreenButton = () => {
  const { localScreenShareOn, toggleScreenShare, presenterId } = useMeeting();
  return (
    <Button
      size="large"
      icon={<DesktopOutlined />}
      type={localScreenShareOn ? "primary" : "default"}
      onClick={() => {
        toggleScreenShare();
      }}
    />
  );
};

const MicButton = () => {
  const mMeeting = useMeeting();
  const [mics, setMics] = useState<
    {
      deviceId: string;
      label: string;
    }[]
  >([]);
  const [selectedMicDeviceId, setSelectedMicDeviceId] = useState("");
  const localMicOn = mMeeting?.localMicOn;
  const changeMic = mMeeting?.changeMic;
  const { token } = theme.useToken();

  const getMics = async () => {
    const mics = await mMeeting?.getMics();

    mics && mics?.length && setMics(mics);
    if (selectedMicDeviceId === "" && mics?.length) {
      setSelectedMicDeviceId(mics[0]?.deviceId);
    }
  };

  useEffect(() => {
    getMics();
  }, []);

  return (
    <Space.Compact>
      <Button
        size="large"
        icon={<AudioOutlined />}
        type={localMicOn ? "primary" : "default"}
        onClick={() => {
          mMeeting.toggleMic();
        }}
      />
      <Dropdown
        trigger={["click"]}
        menu={{
          items: mics?.map((mic) => ({
            key: mic.deviceId,
            value: mic.deviceId,
            label: mic.label,
            onClick: async () => {
              setSelectedMicDeviceId(mic.deviceId);
              changeMic(mic.deviceId);
            },
            style: {
              color:
                selectedMicDeviceId == mic.deviceId
                  ? token.colorPrimary
                  : "black",
            },
          })),
        }}
      >
        <Button size="large" icon={<DownOutlined />} onClick={getMics} />
      </Dropdown>
    </Space.Compact>
  );
};

const WebCamButton = () => {
  const mMeeting = useMeeting();
  const [webcams, setWebcams] = useState<
    {
      deviceId: string;
      label: string;
      facingMode: "environment" | "front";
    }[]
  >([]);
  const [selectWebcamDeviceId, setSelectWebcamDeviceId] = useState("");

  const localWebcamOn = mMeeting?.localWebcamOn;
  const changeWebcam = mMeeting?.changeWebcam;
  const { token } = theme.useToken();

  const getWebcams = async () => {
    const webcams = await mMeeting?.getWebcams();
    webcams && webcams?.length && setWebcams(webcams);
    if (selectWebcamDeviceId === "" && webcams?.length) {
      setSelectWebcamDeviceId(webcams[0]?.deviceId);
    }
  };

  return (
    <Space.Compact>
      <Button
        size="large"
        icon={<VideoCameraOutlined />}
        type={localWebcamOn ? "primary" : "default"}
        onClick={async () => {
          let track;
          if (!localWebcamOn) {
            track = await createCameraVideoTrack({
              optimizationMode: "motion",
              encoderConfig: "h540p_w960p",
              facingMode: "environment",
              multiStream: false,
              cameraId: selectWebcamDeviceId,
            });
          }
          mMeeting.toggleWebcam(track);
        }}
      />
      <Dropdown
        trigger={["click"]}
        menu={{
          items: webcams?.map((webcam) => ({
            key: webcam.deviceId,
            value: webcam.deviceId,
            label: webcam.label,
            onClick: async () => {
              setSelectWebcamDeviceId(webcam.deviceId);
              const track = await createCameraVideoTrack({
                optimizationMode: "motion",
                encoderConfig: "h540p_w960p",
                facingMode: "environment",
                multiStream: false,
                cameraId: webcam.deviceId,
              });
              changeWebcam(track);
              console.log(selectWebcamDeviceId, webcam.deviceId);
            },
            style: {
              color:
                selectWebcamDeviceId == webcam.deviceId
                  ? token.colorPrimary
                  : "black",
            },
          })),
        }}
      >
        <Button size="large" icon={<DownOutlined />} onClick={getWebcams} />
      </Dropdown>
    </Space.Compact>
  );
};

const ParticipantButton = (props: ButtonProps) => {
  const { participants } = useMeeting();

  return (
    <Button size="large" icon={<TeamOutlined />} {...props}>
      {new Map(participants)?.size}
    </Button>
  );
};

const ChatButton = (props: ButtonProps) => {
  return <Button size="large" icon={<MessageOutlined />} {...props} />;
};

const ProductButton = (props: ButtonProps) => {
  return <Button size="large" icon={<SkinFilled />} {...props} />;
};

const VoucherButton = (props: ButtonProps) => {
  return <Button size="large" icon={<WalletOutlined />} {...props} />;
};

const BottomBar = ({
  activePanel,
  setActivePanel,
}: {
  activePanel: string;
  setActivePanel: React.Dispatch<React.SetStateAction<string>>;
}) => {
  const { localParticipant } = useMeeting();

  const { id: livestreamId } = useParams();

  return (
    <Flex
      gap="small"
      justify={
        localParticipant.mode === Constants.modes.CONFERENCE
          ? "space-between"
          : "flex-end"
      }
      align="center"
    >
      {localParticipant.mode === Constants.modes.CONFERENCE && (
        <Flex gap="small" justify="center" align="center">
          <RecordingButton />
          <HlsButton livestreamId={Number(livestreamId)} />
          <ShareScreenButton />
        </Flex>
      )}
      {localParticipant.mode === Constants.modes.CONFERENCE && (
        <Flex gap="small" justify="center" align="center">
          <MicButton />
          <WebCamButton />
          <EndButton livestreamId={Number(livestreamId)} />
        </Flex>
      )}
      <Flex gap="small" justify="center" align="center">
        <VoucherButton
          type={activePanel === "voucher" ? "primary" : "default"}
          onClick={() =>
            setActivePanel((prevState) =>
              prevState === "voucher" ? "" : "voucher",
            )
          }
        />
        <ProductButton
          type={activePanel === "product" ? "primary" : "default"}
          onClick={() =>
            setActivePanel((prevState) =>
              prevState === "product" ? "" : "product",
            )
          }
        />
        <ChatButton
          type={activePanel === "chat" ? "primary" : "default"}
          onClick={() =>
            setActivePanel((prevState) => (prevState === "chat" ? "" : "chat"))
          }
        />
        {localParticipant.mode === Constants.modes.CONFERENCE && (
          <ParticipantButton
            type={activePanel === "participant" ? "primary" : "default"}
            onClick={() =>
              setActivePanel((prevState) =>
                prevState === "participant" ? "" : "participant",
              )
            }
          />
        )}
      </Flex>
    </Flex>
  );
};

export default BottomBar;
