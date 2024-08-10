import { Button, Stack } from "@mui/material";
import DataGrid from "@/components/core/DataGrid";

const LivestreamProductsInfo = () => {
  return (
    <Stack gap={2}>
      <Stack direction="row" spacing={2}>
        <Button variant="contained">Thêm sản phẩm</Button>
      </Stack>
      <DataGrid rows={[]} columns={[]} getRowId={(row) => row.id_product} />
    </Stack>
  );
};

export default LivestreamProductsInfo;
