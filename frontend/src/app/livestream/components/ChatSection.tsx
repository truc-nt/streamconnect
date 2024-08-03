"use client";

import React, { useRef, useEffect, useMemo } from "react";
import { Box, Avatar, Typography, TextField, IconButton } from "@mui/material";
import SendIcon from "@mui/icons-material/Send";

const ChatSection: React.FC = () => {
  const messages = useMemo(
    () => [
      {
        id: 1,
        username: "User1",
        avatar: "/assets/user1-avatar.jpg",
        message: "Lo",
        isCurrentUser: false,
      },
      {
        id: 2,
        username: "User2",
        avatar: "/assets/user2-avatar.jpg",
        message: "Chốt đơn",
        isCurrentUser: false,
      },
      {
        id: 3,
        username: "User",
        avatar: "/assets/current-user-avatar.jpg",
        message: "Ahihi",
        isCurrentUser: true,
      },
      {
        id: 4,
        username: "User3",
        avatar: "/assets/user2-avatar.jpg",
        message: "Chốt đơn",
        isCurrentUser: false,
      },
      {
        id: 5,
        username: "User4",
        avatar: "/assets/user2-avatar.jpg",
        message: "Chốt đơn",
        isCurrentUser: false,
      },
      {
        id: 6,
        username: "User5",
        avatar: "/assets/user2-avatar.jpg",
        message: "Xin giá",
        isCurrentUser: false,
      },
      {
        id: 3,
        username: "User",
        avatar: "/assets/current-user-avatar.jpg",
        message: "Inbox bạn ơi",
        isCurrentUser: true,
      },
    ],
    [],
  );

  const chatContainerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    // Scroll to bottom when component mounts or messages change
    if (chatContainerRef.current) {
      chatContainerRef.current.scrollTop =
        chatContainerRef.current.scrollHeight;
    }
  }, [messages]);

  return (
    <Box
      sx={{
        backgroundColor: "#282a39",
        padding: 2,
        borderRadius: 2,
        display: "flex",
        flexDirection: "column",
      }}
    >
      <Typography
        variant="h6"
        sx={{
          color: "white",
          fontSize: "20px",
          fontWeight: "bold",
          marginBottom: 2,
          textAlign: "center",
        }}
      >
        Trò chuyện Live
      </Typography>

      <Box
        ref={chatContainerRef}
        sx={{
          flexGrow: 1,
          height: "250px",
          overflowY: "auto",
          marginBottom: 2,
          "::-webkit-scrollbar": { display: "none" },
          msOverflowStyle: "none",
          scrollbarWidth: "none",
        }}
      >
        {messages.map((msg) => (
          <Box
            key={msg.id}
            sx={{ display: "flex", alignItems: "center", marginBottom: 1.5 }}
          >
            <Avatar
              alt={msg.username}
              src={msg.avatar}
              sx={{ marginRight: 1 }}
            />
            <Box>
              <Typography
                sx={{ color: "white", fontWeight: "normal", marginBottom: 0.5 }}
              >
                {msg.username}
              </Typography>
              <Box
                sx={{
                  backgroundColor: msg.isCurrentUser ? "#08D2ED" : "#d3d3d7",
                  color: msg.isCurrentUser ? "white" : "#1C1C1C",
                  fontSize: "15px",
                  borderRadius: 2,
                  padding: 1.2,
                  maxWidth: "200px",
                  wordBreak: "break-word",
                  fontWeight: "520",
                  fontFamily: "sans-serif",
                }}
              >
                {msg.message}
              </Box>
            </Box>
          </Box>
        ))}
      </Box>

      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          backgroundColor: "white",
          borderRadius: 8,
          padding: "0 4px",
        }}
      >
        <TextField
          variant="outlined"
          placeholder="Viết gì đó..."
          sx={{
            flexGrow: 1,
            "& .MuiOutlinedInput-root": {
              padding: "4px 8px",
              "& fieldset": { border: "none" },
              "& input": {
                padding: "4px 0",
                fontSize: "14px",
                height: "20px",
                color: "#3D3D3D",
              },
            },
          }}
        />
        <IconButton sx={{ padding: "8px" }}>
          <SendIcon sx={{ color: "#3D3D3D" }} />
        </IconButton>
      </Box>
    </Box>
  );
};

export default ChatSection;
