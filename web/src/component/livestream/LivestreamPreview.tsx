"use client";

import React, {useEffect, useMemo, useRef} from "react";
// import { Box, Avatar, Typography, Badge, IconButton } from "@mui/material";
// import dynamic from "next/dynamic";
// import {
//   FeaturedVideo,
//   Fullscreen,
//   Visibility,
//   VolumeOff,
// } from "@mui/icons-material";
import { keyframes } from "@emotion/react";
import { useRouter } from "next/navigation";
import {useMeeting, Constants} from "@videosdk.live/react-sdk";
import Hls from "hls.js";
import {Avatar, Typography} from "antd";
import {EyeOutlined} from "@ant-design/icons";
import styles from "./LivestreamPreview.module.css";

const gradientAnimation = keyframes`
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
`;

const LivestreamPreview: React.FC<any> = () => {
  const router = useRouter();

  const handleVideoClick = () => {
    router.push("/livestreams/1");
  };
  const { hlsUrls, hlsState } = useMeeting();

  const playHls = useMemo(() => {
    return (
        hlsUrls.livestreamUrl &&
        (hlsState == Constants.hlsEvents.HLS_PLAYABLE ||
            hlsState == Constants.hlsEvents.HLS_STOPPING)
    );
  }, [hlsUrls, hlsState]);

  useEffect(() => {
    if (playHls) {
      if (Hls.isSupported()) {
        const hls = new Hls({
          maxLoadingDelay: 1, // max video loading delay used in automatic start level selection
          defaultAudioCodec: "mp4a.40.2", // default audio codec
          maxBufferLength: 0, // If buffer length is/become less than this value, a new fragment will be loaded
          maxMaxBufferLength: 1, // Hls.js will never exceed this value
          startLevel: 0, // Start playback at the lowest quality level
          startPosition: -1, // set -1 playback will start from intialtime = 0
          maxBufferHole: 0.001, // 'Maximum' inter-fragment buffer hole tolerance that hls.js can cope with when searching for the next fragment to load.
          highBufferWatchdogPeriod: 0, // if media element is expected to play and if currentTime has not moved for more than highBufferWatchdogPeriod and if there are more than maxBufferHole seconds buffered upfront, hls.js will jump buffer gaps, or try to nudge playhead to recover playback.
          nudgeOffset: 0.05, // In case playback continues to stall after first playhead nudging, currentTime will be nudged evenmore following nudgeOffset to try to restore playback. media.currentTime += (nb nudge retry -1)*nudgeOffset
          nudgeMaxRetry: 1, // Max nb of nudge retries before hls.js raise a fatal BUFFER_STALLED_ERROR
          maxFragLookUpTolerance: 0.1, // This tolerance factor is used during fragment lookup.
          liveSyncDurationCount: 1, // if set to 3, playback will start from fragment N-3, N being the last fragment of the live playlist
          abrEwmaFastLive: 1, // Fast bitrate Exponential moving average half-life, used to compute average bitrate for Live streams.
          abrEwmaSlowLive: 3, // Slow bitrate Exponential moving average half-life, used to compute average bitrate for Live streams.
          abrEwmaFastVoD: 1, // Fast bitrate Exponential moving average half-life, used to compute average bitrate for VoD streams
          abrEwmaSlowVoD: 3, // Slow bitrate Exponential moving average half-life, used to compute average bitrate for VoD streams
          maxStarvationDelay: 1, // ABR algorithm will always try to choose a quality level that should avoid rebuffering
        });

        let player = document.querySelector("#hlsPlayer");
        hls.loadSource(hlsUrls.livestreamUrl);
        hls.attachMedia(player);
        hls.on(Hls.Events.MANIFEST_PARSED, function () {});
        hls.on(Hls.Events.ERROR, function (err) {
          console.log(err);
        });
      } else {
        console.error("HLS is not supported");
      }
    }
  }, [playHls]);

  return (
      <div className={styles.livestreamPreview} onClick={handleVideoClick}>
        <div className={styles.videoContainer}>
          <div className={styles.videoWrapper}>
            <video
                id="hlsPlayer"
                autoPlay
                width="100%"
                height="100%"
                controls
                muted
                style={{ pointerEvents: "none" }}
                onError={(err) => {
                  console.log(err, "hls video error");
                }}
            />
          </div>
          {/*<div className={styles.avatarContainer}>*/}
          {/*  <Avatar alt="Username" src="/assets/your-avatar-image.jpg" />*/}
          {/*  <div className={styles.avatarDetails}>*/}
          {/*    <Typography.Text strong style={{ color: "white" }}>*/}
          {/*      Username*/}
          {/*    </Typography.Text>*/}
          {/*    <div className={styles.livestreamInfo}>*/}
          {/*      <Typography.Text style={{ color: "#D8D8D8" }}>*/}
          {/*        Livestream Name*/}
          {/*      </Typography.Text>*/}
          {/*      <EyeOutlined style={{ fontSize: "16px", color: "#D8D8D8" }} />*/}
          {/*      <Typography.Text style={{ color: "#D8D8D8" }}>*/}
          {/*        1234*/}
          {/*      </Typography.Text>*/}
          {/*    </div>*/}
          {/*  </div>*/}
          {/*</div>*/}
        </div>
      </div>
  );
};

export default LivestreamPreview;
