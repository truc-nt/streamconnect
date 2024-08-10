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

const LivestreamInfo = () => {
  const dispatch = useAppDispatch();

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);
    dispatch(
      setLivestreamInformation({
        title: formData.get("title"),
        description: formData.get("description"),
        image: formData.get("image"),
      }),
    );
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
          </Grid>
          <FormControl>
            <FormLabel htmlFor="image">Ảnh bìa</FormLabel>
            <TextField type="file"></TextField>
          </FormControl>
        </Grid>
        <Button variant="contained" type="submit" sx={{ marginLeft: "auto" }}>
          Tiếp theo
        </Button>
      </FormGroup>
    </form>
  );
};

export default LivestreamInfo;
