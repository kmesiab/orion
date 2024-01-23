import { useState } from "react";
import "./App.scss";
import { ProcessPullRequest } from "../wailsjs/go/main/App.js";
import { GetCodeReviewFromAPI } from "../wailsjs/go/main/App.js";
import DiffViewer from "./components/DiffViewer/DiffViewer.jsx";
import AnalysisViewer from "./components/AnalysisViewer/AnalysisViewer.jsx";
import Stack from "@mui/material/Stack";
import ListItem from "@mui/material/ListItem";
import Divider from "@mui/material/Divider";

function App() {
  const [pullRequestText, setPullRequestText] = useState("");
  const [codeReviewText, setCodeReviewText] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [url, setUrl] = useState("");

  const updateUrl = (e) => setUrl(e.target.value);
  const processPullRequestText = (result) => {
    setPullRequestText(result);
    //Once we have pull request text we can call the API to get the code review text
    setIsLoading(true);
    GetCodeReviewFromAPI(result).then(proccessCodeReview);
  };
  const proccessCodeReview = (result) => {
    setCodeReviewText(result);
    setIsLoading(false);
  };

  function processPullRequest() {
    ProcessPullRequest(url).then(processPullRequestText);
  }

  return (
    <div id="App">
      <div id="input" className="input-box">
        <input
          id="url"
          className="input"
          onChange={updateUrl}
          name="input"
          type="text"
          placeholder={"https://github.com/owner/repo/pull/1234"}
        />
        <button className="btn" onClick={processPullRequest}>
          Review
        </button>
      </div>
      <Stack spacing={2}>
        <ListItem>
          <AnalysisViewer loading={isLoading} analysisReport={codeReviewText} />
        </ListItem>
        {codeReviewText && <Divider />}
        <ListItem>
          <DiffViewer text={pullRequestText} />
        </ListItem>
      </Stack>
    </div>
  );
}

export default App;
