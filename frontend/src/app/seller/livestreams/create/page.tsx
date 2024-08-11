"use client";
import Box from "@mui/material/Box";
import Stepper from "@mui/material/Stepper";
import Step from "@mui/material/Step";
import StepLabel from "@mui/material/StepLabel";
import { Paper, Container, Grid, Button, Stack } from "@mui/material";
import { FormGroup, TextField, FormControl, FormLabel } from "@mui/material";
import { useState } from "react";
import LivestreamInfo from "./components/LivestreamInfo";
import LivestreamVariants from "./components/AddLivestreamProduct";
import LivestreamSchedule from "./components/LivestreamSchedule";

import ProductDetail from "./components/ProductDetail";
import { useAppSelector, useAppDispatch } from "@/store/store";

const steps = [
  {
    label: "Thêm sản phẩm",
    component: <LivestreamVariants />,
  },
  {
    label: "Thông tin livestream",
    component: <LivestreamInfo />,
  },
];

const Page = () => {
  const { currentStep } = useAppSelector((state) => state.livestreamCreate);

  return (
    <Paper sx={{ paddingY: "50px", paddingX: "24px" }}>
      <Stack gap={2}>
        <Stepper
          activeStep={currentStep}
          alternativeLabel
          sx={{ marginBottom: "20px" }}
        >
          {steps.map((step) => (
            <Step key={step.label}>
              <StepLabel>{step.label}</StepLabel>
            </Step>
          ))}
        </Stepper>
        {steps[currentStep].component}
      </Stack>
    </Paper>
  );
};

export default Page;
