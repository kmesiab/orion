import React from "react";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Markdown from "react-markdown";
import LoadingSpinner from "../LoadingSpinner/LoadingSpinner";

const AnalysisViewer = ({ loading, analysisReport }) => {
  if (loading) {
    return <LoadingSpinner text="Generating your report... 🤖" />;
  }
  if (!analysisReport) {
    return null;
  }
  const analysisStyles = {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    width: "100%",
  };

  return (
    <Box sx={analysisStyles}>
      <Typography variant="h4" gutterBottom>
        Analysis Report
      </Typography>
      <Markdown>{analysisReport}</Markdown>
    </Box>
  );
};

export default AnalysisViewer;
