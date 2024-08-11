import {
  FormGroup,
  FormControl,
  FormLabel,
  TextField,
  Button,
} from "@mui/material";
import { Grid } from "@mui/material";
import { useState } from "react";
import { useAppSelector, useAppDispatch } from "@/store/store";
import { setLivestreamInformation } from "@/store/livestream_create";
import { DateTimePicker } from "@mui/x-date-pickers";
import { createLivestream } from "@/api/livestream";
import { setOpen } from "@/store/alert";
import { reset } from "@/store/livestream_create";

const LivestreamInfo = () => {
  const dispatch = useAppDispatch();
  const { livestreamExternalVariants } = useAppSelector(
    (state: any) => state.livestreamCreate,
  );

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);

    dispatch(
      setLivestreamInformation({
        title: formData.get("title"),
        description: formData.get("description"),
        startTime: formData.get("date"),
      }),
    );

    const livestreamCreateData = {
      title: formData.get("title"),
      description: formData.get("description"),
      start_time: formData.get("date"),
      livestream_external_variants: livestreamExternalVariants.reduce(
        (acc, variant) =>
          acc.concat(
            variant.externalVariants.map((externalVariant: any) => ({
              id_product: variant.idProduct,
              id_variant: variant.idVariant,
              id_external_variant: externalVariant.idExternalVariant,
              quantity: externalVariant.quantity,
            })),
          ),
        [],
      ),
    };

    try {
      await createLivestream(1, livestreamCreateData);
      dispatch(
        setOpen({
          message: "Tạo livestream thành công",
          type: "success",
        }),
      );
      dispatch(reset());
    } catch (error) {
      dispatch(
        setOpen({
          message: "Tạo livestream không thành công",
          type: "error",
        }),
      );
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <FormGroup sx={{ gap: 2 }}>
        <Grid
          container
          gap={2}
          sx={{ display: "flex", justifyContent: "space-between" }}
        >
          <Grid
            item
            xs={8}
            sx={{ display: "flex", flexDirection: "column", gap: 2 }}
          >
            <FormControl>
              <FormLabel htmlFor="title" required>
                Tiêu đề
              </FormLabel>
              <TextField
                id="title"
                multiline
                rows={2}
                required
                variant="outlined"
                name="title"
              />
            </FormControl>
            <FormControl>
              <FormLabel htmlFor="description">Mô tả</FormLabel>
              <TextField
                id="description"
                placeholder="Giới thiệu về livestream của bạn cho người xem"
                multiline
                rows={4}
                name="description"
              />
            </FormControl>
            <FormControl>
              <FormLabel htmlFor="date">Ngày công chiếu</FormLabel>
              <DateTimePicker
                name="date"
                ampm={false}
                //value={selectedDate}
                //onChange={(newDate) => setSelectedDate(newDate)}
                //renderInput={(params) => <TextField {...params} />}
              />
            </FormControl>
          </Grid>
          <FormControl>
            <FormLabel htmlFor="image">Ảnh bìa</FormLabel>
            <TextField type="file"></TextField>
          </FormControl>
        </Grid>
        <Button variant="contained" type="submit" sx={{ marginLeft: "auto" }}>
          Hoàn thành
        </Button>
      </FormGroup>
    </form>
  );
};

export default LivestreamInfo;
