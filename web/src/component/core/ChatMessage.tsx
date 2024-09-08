import { Flex, Typography, theme, Divider } from "antd";
interface IChatMessageProps {
  sender: string;
  isLocalSender: boolean;
  message: string;
  createdAt: string;
}

const ChatMessage = ({
  sender,
  isLocalSender,
  message,
  createdAt,
}: IChatMessageProps) => {
  const { token } = theme.useToken();

  return (
    <Flex
      vertical
      wrap
      align={isLocalSender ? "end" : "start"}
      style={{
        backgroundColor: isLocalSender
          ? token.colorPrimary
          : token.colorTextSecondary,
        marginLeft: isLocalSender ? "auto" : "0",
      }}
      className="p-2 rounded-md w-fit max-w-full"
    >
      <Flex align="center">
        <Typography.Text className="text-xs text-[#ffffff80]">
          {isLocalSender ? "You" : sender}
        </Typography.Text>
        <Divider type="vertical" className="border-[#ffffff80] mx-1" />
        <Typography.Text className="text-[#ffffff80] text-xs italic">
          {new Date(createdAt).toLocaleTimeString()}
        </Typography.Text>
      </Flex>
      <div style={{ overflowWrap: "break-word", wordWrap: "break-word" }}>
        <p className="inline-block whitespace-pre-wrap break-words text-right text-white">
          {message}
        </p>
      </div>
    </Flex>
  );
};

export default ChatMessage;
