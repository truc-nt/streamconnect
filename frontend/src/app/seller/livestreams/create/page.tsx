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

import LivestreamProductsInfo from "./components/LivestreamProductsInfo";
import ProductDetail from "./components/ProductDetail";
import { useAppSelector, useAppDispatch } from "@/store/store";
import { setPrevStep, setNextStep } from "@/store/livestream_create";
import { current } from "@reduxjs/toolkit";

const steps = [
  {
    label: "Thông tin livestream",
    component: <LivestreamInfo />,
  },
  {
    label: "Thêm sản phẩm",
    component: <LivestreamVariants />,
  },
  {
    label: "Hoàn thành",
    component: <ProductDetail />,
  },
];

const Page = () => {
  const dispatch = useAppDispatch();
  /*const { currentStep } = useSelector(
    (state: RootStore) => state.livestreamCreate,
  );*/

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
