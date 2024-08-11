import { Grid } from "@mui/material";
import LivestreamFrame from "./components/LivestreamFrame";
import ChatSection from "./components/ChatSection";
import ProductList from "./components/ProductList";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <Grid container spacing={2}>
        <Grid item xs={12} lg={8}>
          <LivestreamFrame />
        </Grid>
        <Grid item xs={12} lg={4}>
          <ChatSection />
        </Grid>
      </Grid>
      <ProductList />
    </>
  );
}
