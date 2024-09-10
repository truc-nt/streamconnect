import { Constants, useMeeting } from "@videosdk.live/react-sdk";
import { useMemo } from "react";

export const useIsHls = () => {
  const { hlsState } = useMeeting();
  const isHls = useMemo(
    () =>
      hlsState === Constants.hlsEvents.HLS_STARTED ||
      hlsState === Constants.hlsEvents.HLS_PLAYABLE ||
      hlsState === Constants.hlsEvents.HLS_STOPPING,
    [hlsState],
  );

  return isHls;
};
