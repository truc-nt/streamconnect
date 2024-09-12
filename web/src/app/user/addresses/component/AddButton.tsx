"use client";
import { theme } from "antd";
const AddButton = () => {
  const { token } = theme.useToken();
  return (
    <div
      style={{
        display: "flex",
        height: "80px",
        alignItems: "center",
        justifyContent: "center",
        color: token.colorTextTertiary,
        backgroundColor: token.colorFillAlter,
        borderRadius: token.borderRadiusLG,
        border: `1px dashed ${token.colorBorder}`,
        cursor: "pointer",
      }}
    >
      Thêm địa chỉ
    </div>
  );
};

export default AddButton;
