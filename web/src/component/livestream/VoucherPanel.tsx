import { useState } from "react";
import { Modal, Button, Space, List, Typography } from "antd";
import { PlusCircleOutlined } from "@ant-design/icons";
import VoucherItem from "@/component/list_item/VoucherItem";
import VoucherModal from "@/component/modal/VoucherModal";
import { useGetShopVouchers } from "@/hook/voucher";
import {
  Constants,
  useMeeting,
  useParticipant,
} from "@videosdk.live/react-sdk";
import { useMeetingAppContext } from "@/component/livestream/MeetingProvider";
import useLoading from "@/hook/loading";
import { addUserVoucher } from "@/api/voucher";

const VoucherPanel = () => {
  const [openAddModal, setOpenAddModal] = useState(false);
  const { localParticipant } = useMeeting();
  const { mode } = useParticipant(localParticipant?.id ?? "");
  const { shopId } = useMeetingAppContext();
  const { data: shopVouchers, mutate: mutateGetShopVouchers } =
    useGetShopVouchers(shopId);

  const handleAddVoucher = useLoading(
    addUserVoucher,
    "Thêm voucher thành công",
    "Thêm voucher thất bại",
  );

  return (
    <>
      <div className="flex-1 overflow-y-scroll p-2">
        <List
          grid={{ gutter: [2, 2], column: 1 }}
          dataSource={shopVouchers ?? []}
          renderItem={(item) => (
            <List.Item key={item?.id_voucher} style={{ padding: 0 }}>
              <VoucherItem
                {...item}
                button={
                  mode === Constants.modes.VIEWER &&
                  (item?.is_saved ? (
                    <Typography.Text>Đã lưu</Typography.Text>
                  ) : (
                    <Button
                      type="primary"
                      size="small"
                      onClick={async () => {
                        try {
                          await handleAddVoucher(item?.id_voucher);
                          mutateGetShopVouchers();
                        } catch (e) {
                          console.error(e);
                          // Optionally show a notification to the user
                        }
                      }}
                    >
                      Lưu
                    </Button>
                  ))
                }
              />
            </List.Item>
          )}
        />
      </div>
      {mode === Constants.modes.CONFERENCE && (
        <Space.Compact className="m-auto">
          <Button
            icon={<PlusCircleOutlined />}
            onClick={() => setOpenAddModal(true)}
            aria-label="Add Voucher"
          />
        </Space.Compact>
      )}
      {mode === Constants.modes.CONFERENCE && (
        <VoucherModal
          open={openAddModal}
          onCancel={() => setOpenAddModal(false)}
          successfullySubmitPostAction={() => mutateGetShopVouchers()}
        />
      )}
    </>
  );
};

export default VoucherPanel;
