import { useParams } from "next/navigation";
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
