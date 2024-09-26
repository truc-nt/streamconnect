import { useGetLivestreamProducts } from "@/hook/livestream";
import ProductItem from "@/component/list_item/ProductItem";
import { List } from "antd";
import LivestreamProductSellingModal from "@/app/livestreams/[id]/component/LivestreamProductSellingModal";
import EditLivestreamProductModal from "@/app/livestreams/[id]/component/EditLivestreamProductModal";
import { useState } from "react";
import { useParams } from "next/navigation";
import { Button, Space, Modal, Flex, Checkbox } from "antd";
import {
  PushpinOutlined,
  PlusCircleOutlined,
  ArrowRightOutlined,
  HeartOutlined,
} from "@ant-design/icons";
import { CheckboxChangeEvent } from "antd/lib/checkbox";
import {
  updateLivestreamProductPriority,
  IPinLivestreamProduct,
} from "@/api/livestream_product";
import useLoading from "@/hook/loading";
import ChosenLivestreamVariant from "@/component/livestream_variant/ChosenLivestreamVariant";
import { addLivestreamProduct, ILivestreamProduct } from "@/api/livestream";
import { IChosenLivestreamVariant } from "@/app/seller/livestreams/create/component/LivestreamCreate";
import { useAppSelector } from "@/store/store";
import {
  Constants,
  useMeeting,
  useParticipant,
} from "@videosdk.live/react-sdk";
import { useMeetingAppContext } from "@/component/livestream/MeetingProvider";
import LivestreamProductList from "@/component/list/LivestreamProductList";
import { LivestreamStatus } from "@/constant/livestream";

const ProductPanel = () => {
  const { id: livestreamId } = useParams();
  const { shopId, livestreamStatus } = useMeetingAppContext();

  const { localParticipant } = useMeeting();
  const { mode } = useParticipant(localParticipant?.id ?? "");

  return (
    <LivestreamProductList
      shopId={shopId}
      livestreamId={Number(livestreamId)}
      mode={mode}
      status={
        LivestreamStatus[
          livestreamStatus.toUpperCase() as keyof typeof LivestreamStatus
        ]
      }
    />
  );
};

export default ProductPanel;
