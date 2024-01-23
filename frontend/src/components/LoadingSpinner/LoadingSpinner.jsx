import * as React from "react";
import Box from "@mui/material/Box";
import LinearProgress from "@mui/material/LinearProgress";
import Typography from "@mui/material/Typography";

const LoadingSpinner = ({ text }) => {
  return (
    <Box sx={{ width: "100%", textAlign: "center", marginTop: 2 }}>
      {text && <Typography variant="body2">{text}</Typography>}
      <LinearProgress />
    </Box>
  );
};

export default LoadingSpinner;
