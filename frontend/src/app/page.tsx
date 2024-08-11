import CategorySlider from "@/components/CategorySlider";
import LivestreamPreview from "@/components/LivestreamPreview";
import LivestreamsList from "@/components/LivestreamsList";
import { Typography } from "@mui/material";
import Alert from "@/components/core/Alert";
import { useAppSelector } from "@/store/store";

const Page = () => {
  return (
    <>
      <CategorySlider />
      <LivestreamPreview />
      <Typography
        variant="h6"
        sx={{ fontSize: "20px", fontWeight: "bold", color: "white" }}
      >
        Livestreams Được Đề Xuất
      </Typography>
      <LivestreamsList />
    </>
  );
};

export default Page;
