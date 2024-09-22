import { usePubSub, useMeeting } from "@videosdk.live/react-sdk";
import { useState, useEffect, useRef } from "react";
import ChatMessage from "@/component/core/ChatMessage";
import { Flex, List, Space, Input, Button } from "antd";
import { SendOutlined } from "@ant-design/icons";
import { useAppSelector } from "@/store/store";

const ChatPanel = () => {
  const { publish, messages } = usePubSub("CHAT");
  const mMeeting = useMeeting();
  const { username } = useAppSelector((state) => state.authReducer);

  const [message, setMessage] = useState("");

  const handleSendMessage = () => {
    message.trim();
    publish(message, { persist: true });
    console.log("why");
    setMessage("");
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      e.preventDefault();
      handleSendMessage();
    }
  };

  const chatContainerRef = useRef<HTMLDivElement>(null);
  useEffect(() => {
    if (chatContainerRef.current) {
      chatContainerRef.current.scrollTop =
        chatContainerRef.current.scrollHeight;
    }
  }, [messages]);

  return (
    <>
      <div className="flex-1 overflow-y-scroll p-2" ref={chatContainerRef}>
        <List
          dataSource={messages}
          renderItem={(item, index) => {
            const { senderName, message, timestamp } = item;
            return (
              <List.Item style={{ padding: 0, border: 0, marginBottom: "5px" }}>
                <ChatMessage
                  key={index}
                  sender={senderName}
                  isLocalSender={senderName === username}
                  message={message}
                  createdAt={timestamp}
                />
              </List.Item>
            );
          }}
          className="p-1"
        />
      </div>
      <Space.Compact>
        <Input
          placeholder="Bình luận"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={handleKeyDown}
        />
        <Button
          type="primary"
          onClick={handleSendMessage}
          icon={<SendOutlined />}
        />
      </Space.Compact>
    </>
  );
};

export default ChatPanel;
