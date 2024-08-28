import { Tag as AntdTag } from "antd";
const Tag = ({ label }: { label: string }) => {
  const colors = [
    "magenta",
    "red",
    "volcano",
    "orange",
    "gold",
    "lime",
    "green",
    "cyan",
    "blue",
    "geekblue",
    "purple",
  ];
  return (
    <AntdTag color={colors[Math.floor(Math.random() * colors.length)]}>
      {label}
    </AntdTag>
  );
};

export default Tag;
