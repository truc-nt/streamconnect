import { Button, Modal, Form, ModalProps } from "antd";
import VoucherForm from "@/component/form/VoucherForm";

import { createVoucher } from "@/api/voucher";
import useLoading from "@/hook/loading";

interface IVoucherModalProps extends ModalProps {
  successfullySubmitPostAction?: () => void;
}
const VoucherModal = (props: IVoucherModalProps) => {
  const [form] = Form.useForm();
  const handleCreateVoucher = useLoading(
    createVoucher,
    "Tạo voucher thành công",
    "Tạo voucher thất bại",
  );
  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();

      const request = {
        ...values,
      };

      await handleCreateVoucher(request);
      props.successfullySubmitPostAction?.();
    } catch (e) {}
  };
  return (
    <Modal
      footer={null}
      centered
      title="Thông tin voucher"
      width="60%"
      {...props}
    >
      <VoucherForm form={form} />
      <Button type="primary" className="w-full" onClick={handleSubmit}>
        Thêm
      </Button>
    </Modal>
  );
};

export default VoucherModal;
