"use client";
import { DatePicker } from "@mui/x-date-pickers";
import { Stack } from "@mui/material";

const LivestreamSchedule = () => {
  return (
    <form>
      <Stack gap={2}>
        <DatePicker />
      </Stack>
    </form>
  );
};

export default LivestreamSchedule;
