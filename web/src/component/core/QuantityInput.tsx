import { Input, Space, Button } from "antd";
import { PlusOutlined, MinusOutlined } from "@ant-design/icons";
const QuantityInput = ({
  quantity,
  onDecrease,
  onIncrease,
}: {
  quantity?: number;
  onDecrease: () => void;
  onIncrease: () => void;
}) => {
  return (
    <Space.Compact block>
      <Button icon={<MinusOutlined />} size="small" onClick={onDecrease} />
      <Input
        size="small"
        defaultValue={quantity ?? 1}
        style={{ width: "30px" }}
      />
      <Button icon={<PlusOutlined />} size="small" onClick={onIncrease} />
    </Space.Compact>
  );
};

export default QuantityInput;
