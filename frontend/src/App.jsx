import {useState} from 'react';
import './App.scss';
import {ProcessPullRequest} from "../wailsjs/go/main/App.js";
import DiffViewer from "./components/DiffViewer/DiffViewer.jsx";

function App() {
    const [resultText, setResultText] = useState('');
    const [url, setName] = useState('');

    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);

    function processPullRequest() {
        ProcessPullRequest(url).then(updateResultText);
    }

    return (
        <div id="App">
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} name="input" type="text" placeholder={"https://github.com/owner/repo/pull/1234"}/>
                <button className="btn" onClick={processPullRequest}>Review</button>
            </div>
            <DiffViewer text={resultText} />
        </div>
    )
}

export default App
