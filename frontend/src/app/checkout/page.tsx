"use client";

import React, { useState } from "react";
import {
  Box,
  Button,
  Typography,
  Stepper,
  Step,
  StepLabel,
} from "@mui/material";
import { styled } from "@mui/material/styles";
import DoneIcon from "@mui/icons-material/Done";
import { useRouter } from "next/navigation";

const steps = [
  "Kiểm tra sản phẩm đặt hàng",
  "Xác nhận thông tin đơn hàng",
  "Thanh toán",
];

const CustomStepIconRoot = styled("div")<{
  ownerState: { active?: boolean; completed?: boolean };
}>(({ ownerState }) => ({
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  width: 28,
  height: 28,
  borderRadius: "50%",
  backgroundColor:
    ownerState.active || ownerState.completed ? "#08d1ed" : "transparent",
  border: `1.5px solid ${ownerState.active || ownerState.completed ? "#08d1ed" : "white"}`,
  color: "white",
  fontSize: "12px",
  cursor: "pointer",
  fontFamily: "Roboto",
}));

const CustomStepIcon: React.FC<{
  active?: boolean;
  completed?: boolean;
  icon?: React.ReactNode;
  className?: string;
}> = ({ active = false, completed = false, icon, className }) => {
  return (
    <CustomStepIconRoot
      ownerState={{ active, completed }}
      className={className}
    >
      <span>{icon}</span>
    </CustomStepIconRoot>
  );
};

const Checkout: React.FC = () => {
  const router = useRouter();

  const [activeStep, setActiveStep] = useState(0);

  const handleHomeClick = () => {
    router.push("/");
  };

  const handleNext = () => {
    setActiveStep((prevActiveStep) => prevActiveStep + 1);
  };

  const handleBack = () => {
    setActiveStep((prevActiveStep) => prevActiveStep - 1);
  };

  const handleStepClick = (step: number) => {
    setActiveStep(step);
  };

  return (
    <>
      <Typography
        variant="h6"
        sx={{
          color: "white",
          fontSize: "20px",
          fontWeight: "bold",
          textAlign: "center",
          marginBottom: 2,
        }}
      >
        Thanh toán
      </Typography>
      <Stepper activeStep={activeStep} alternativeLabel>
        {steps.map((label, index) => (
          <Step key={label} onClick={() => handleStepClick(index)}>
            <StepLabel
              StepIconComponent={(props) => (
                <CustomStepIcon {...props} icon={index + 1} />
              )}
              sx={{
                "& .MuiStepLabel-label": {
                  color: "white",
                  "&.Mui-active": {
                    color: "white",
                  },
                  "&.Mui-completed": {
                    color: "white",
                  },
                },
              }}
            >
              {label}
            </StepLabel>
          </Step>
        ))}
      </Stepper>
      {activeStep === steps.length ? (
        <Box
          sx={{
            padding: 3,
            backgroundColor: "#282a39",
            mx: 14,
            my: 2,
            textAlign: "center",
            height: "300px",
            borderRadius: 4,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
          }}
        >
          <Box sx={{ display: "flex", alignItems: "center", marginBottom: 2 }}>
            <Box
              sx={{
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                width: 40,
                height: 40,
                borderRadius: "50%",
                backgroundColor: "transparent",
                border: `1.5px solid #08d2ed`,
                marginRight: 2,
              }}
            >
              <DoneIcon sx={{ color: "#08d2ed" }} />
            </Box>
            <Typography variant="h6" sx={{ color: "white" }}>
              Đặt hàng thành công
            </Typography>
          </Box>
          <Box
            sx={{ display: "flex", justifyContent: "center", gap: 2, mt: 2 }}
          >
            <Button
              variant="contained"
              sx={{
                backgroundColor: "#08d1ed",
                color: "white",
                "&:hover": {
                  backgroundColor: "#08d1ed",
                  borderColor: "transparent",
                },
              }}
              onClick={handleHomeClick}
            >
              Quay về trang chủ
            </Button>
            <Button
              variant="contained"
              sx={{
                backgroundColor: "#08d1ed",
                color: "white",
                "&:hover": {
                  backgroundColor: "#08d1ed",
                  borderColor: "transparent",
                },
              }}
            >
              Xem lại đơn hàng
            </Button>
          </Box>
        </Box>
      ) : (
        <>
          {/* Board Content */}
          <Box
            sx={{
              padding: 3,
              backgroundColor: "#282a39",
              marginTop: 2,
              height: "300px",
              borderRadius: 4,
              mx: 2,
            }}
          >
            {activeStep === 0 && (
              <Box sx={{ marginBottom: 2 }}>
                {/* Kiểm tra sản phẩm đặt hàng */}
              </Box>
            )}
            {activeStep === 1 && (
              <Box sx={{ marginBottom: 2 }}>
                {/* Xác nhận thông tin đơn hàng */}
              </Box>
            )}
            {activeStep === 2 && (
              <Box sx={{ marginBottom: 2 }}>{/* Thanh toán */}</Box>
            )}
          </Box>

          {/* Navigate Buttons */}
          <Box
            sx={{
              display: "flex",
              justifyContent: "flex-end",
              marginTop: 2,
              mx: 2,
            }}
          >
            <Button
              disabled={activeStep === 0}
              onClick={handleBack}
              variant="contained"
              sx={{
                backgroundColor: "#e2e2e2",
                color: "black",
                marginRight: 2,
                "&:hover": {
                  backgroundColor: "#e2e2e2",
                  borderColor: "transparent",
                },
              }}
            >
              Quay lại
            </Button>
            <Button
              variant="contained"
              sx={{
                backgroundColor: "#08d1ed",
                color: "white",
                "&:hover": {
                  backgroundColor: "#08d1ed",
                  borderColor: "transparent",
                },
              }}
              onClick={handleNext}
            >
              {activeStep === steps.length - 1 ? "Hoàn tất" : "Tiếp theo"}
            </Button>
          </Box>
        </>
      )}
    </>
  );
};

export default Checkout;
