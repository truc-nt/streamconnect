"use client";

import React from "react";
import { useRouter } from "next/navigation";
import {
  Box,
  Avatar,
  Typography,
  Button,
  Badge,
  IconButton,
} from "@mui/material";
import ReactPlayer from "react-player";
import { Fullscreen, Visibility, VolumeOff } from "@mui/icons-material";

const formatViews = (views: number) => {
  return views.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
};

const LivestreamFrame: React.FC = () => {
  const router = useRouter();

  const handleContainerClick = (event: React.MouseEvent) => {
    router.push("/livestream");
  };

  const handleButtonClick = (event: React.MouseEvent) => {
    event.stopPropagation();
  };

  const viewCount = 12345885;

  return (
    <Box
      sx={{
        width: "100%",
        height: "100%",
      }}
      onClick={handleContainerClick}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          mb: 2,
        }}
      >
        <Box sx={{ display: "flex", alignItems: "center" }}>
          <Avatar alt="Username" src="/assets/your-avatar-image.jpg" />
          <Typography
            variant="subtitle1"
            sx={{ fontWeight: "bold", color: "white", ml: 2, fontSize: "18px" }}
          >
            Username
          </Typography>
        </Box>
        <Button
          variant="contained"
          sx={{
            backgroundColor: "#08D2ED",
            color: "white",
            textTransform: "none",
          }}
          onClick={handleButtonClick}
        >
          Follow
        </Button>
      </Box>

      <Box
        sx={{
          position: "relative",
          width: "100%",
          paddingTop: "56.25%", // 16:9 aspect ratio
          borderRadius: "8px",
          overflow: "hidden",
        }}
        onClick={(event) => event.stopPropagation()}
      >
        <Box
          sx={{
            position: "absolute",
            top: 0,
            left: 0,
            width: "100%",
            height: "100%",
          }}
        >
          <ReactPlayer
            url="https://www.youtube.com/watch?v=VBKNoLcj8jA"
            width="100%"
            height="100%"
            controls
            playing
            muted
          />
        </Box>

        {/* Badge Live */}
        <Badge
          badgeContent="Live"
          color="error"
          sx={{ position: "absolute", top: 24, left: 32, zIndex: 3 }}
        >
          <Box />
        </Badge>

        {/* View-count */}
        <Box
          sx={{
            position: "absolute",
            top: 16,
            right: 16,
            display: "flex",
            alignItems: "center",
            zIndex: 3,
            gap: 1,
          }}
        >
          <Visibility sx={{ fontSize: "16px", color: "#D8D8D8" }} />
          <Typography variant="body2" sx={{ color: "#D8D8D8" }}>
            {formatViews(viewCount)}
          </Typography>
        </Box>

        <Box
          sx={{
            position: "absolute",
            bottom: 16,
            right: 16,
            display: "flex",
            gap: 1,
            zIndex: 3,
          }}
        >
          <IconButton sx={{ color: "white" }} onClick={handleButtonClick}>
            <Fullscreen />
          </IconButton>
          <IconButton sx={{ color: "white" }} onClick={handleButtonClick}>
            <VolumeOff />
          </IconButton>
        </Box>
      </Box>
    </Box>
  );
};

export default LivestreamFrame;
