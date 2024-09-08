import { useMemo } from "react";
import { Constants, useMeeting } from "@videosdk.live/react-sdk";

export const useIsRecording = () => {
  const { recordingState } = useMeeting();

  const isRecording = useMemo(
    () =>
      recordingState === Constants.recordingEvents.RECORDING_STARTED ||
      recordingState === Constants.recordingEvents.RECORDING_STOPPING,
    [recordingState],
  );

  return isRecording;
};
