import { Button, Modal, Form, ModalProps } from "antd";
import AddressForm from "@/component/form/AddressForm";

import { createAddress } from "@/api/user";
import useLoading from "@/hook/loading";

interface IAddressModalProps extends ModalProps {
  successfullySubmitPostAction?: () => void;
}
const AddressModal = (props: IAddressModalProps) => {
  const [form] = Form.useForm();
  const handleCreateAddress = useLoading(
    createAddress,
    "Tạo địa chỉ thành công",
    "Tạo địa chỉ thất bại",
  );

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields();

      const request = {
        ...values,
      };

      console.log(request);
      await handleCreateAddress(request);
      props.successfullySubmitPostAction?.();
    } catch (e) {}
  };
  return (
    <Modal
      footer={null}
      centered
      title="Thông tin địa chỉ"
      width="60%"
      {...props}
    >
      <AddressForm form={form} />
      <Button type="primary" className="w-full" onClick={handleSubmit}>
        Thêm
      </Button>
    </Modal>
  );
};

export default AddressModal;
