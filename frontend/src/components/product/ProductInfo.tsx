import { Paper, Stack, Box, Typography, Divider, Button } from "@mui/material";

interface IProductInfo {
  name: string;
  option: Record<string, string[]>;
  chosenOption: Record<string, string>;
  handleChangeOption: (key: string, value: string) => void;
}

const ProductInfo = ({
  name,
  option,
  chosenOption,
  handleChangeOption,
}: IProductInfo) => {
  return (
    <Stack spacing={2}>
      <Stack direction="row" spacing={2}>
        <Box sx={{ width: 100, height: 100 }}>
          <Box
            component="img"
            sx={{
              width: "100%",
              height: "100%",
              objectFit: "cover",
            }}
            alt="The house from the offer."
            src="https://images.unsplash.com/photo-1512917774080-9991f1c4c750?auto=format&w=350&dpr=2"
          />
        </Box>
        <Typography variant="h6">{name}</Typography>
      </Stack>
      <Divider />
      {Object.entries(option).map(([key, values]) => (
        <Box key={key}>
          <Typography variant="body2">{key}</Typography>
          <Stack spacing={2} direction="row">
            {values.map((value, index) => (
              <Button
                key={index}
                variant="contained"
                color={chosenOption[key] === value ? "primary" : "secondary"}
                onClick={() => handleChangeOption(key, value)}
              >
                {value}
              </Button>
            ))}
          </Stack>
        </Box>
      ))}
    </Stack>
  );
};

export default ProductInfo;
