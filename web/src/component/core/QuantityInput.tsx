import { Input, Space, Button, InputProps, InputNumberProps } from "antd";
import { PlusOutlined, MinusOutlined } from "@ant-design/icons";

interface IQuantityInputProps extends InputNumberProps {
  onDecrease: () => void;
  onIncrease: () => void;
}

const QuantityInput = ({
  onDecrease,
  onIncrease,
  ...props
}: IQuantityInputProps) => {
  return (
    <Space.Compact block>
      <Button icon={<MinusOutlined />} size="small" onClick={onDecrease} />
      <Input size="small" style={{ width: "30px" }} {...props} />
      <Button icon={<PlusOutlined />} size="small" onClick={onIncrease} />
    </Space.Compact>
  );
};

export default QuantityInput;
