import { ModalProps, Modal } from "antd";
import { useGetLivestreamProduct } from "@/hook/livestream_product";
import LivestreamProductInformation from "./LivestreamProductInformation";

interface ILivestreamProductProps extends ModalProps {
  livestreamProductId: number;
}

const LivestreamProductModal = ({
  livestreamProductId,
  ...props
}: ILivestreamProductProps) => {
  const { data } = useGetLivestreamProduct(livestreamProductId);
  return (
    <Modal
      {...props}
      centered
      title="Thông tin sản phẩm"
      open={true}
      footer={null}
      width="60%"
    >
      {data && <LivestreamProductInformation livestreamProduct={data!} />}
    </Modal>
  );
};

export default LivestreamProductModal;
